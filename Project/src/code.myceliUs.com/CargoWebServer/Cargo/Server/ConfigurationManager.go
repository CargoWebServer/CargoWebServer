package Server

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	b64 "encoding/base64"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/CargoWebServer/Cargo/JS"
	"code.myceliUs.com/Utility"
)

const (
	// The configuration db
	ConfigDB = "Config"
)

/**
 * The configuration manager is use to keep all
 * internal settings and all external services settings
 * used by applications served.
 */
type ConfigurationManager struct {

	// This is the root of the server.
	// must be set at server starting time and never be saved!
	m_filePath string

	// the active configurations...
	m_activeConfigurationsEntity *Config_ConfigurationsEntity

	// The list of service configurations...
	m_servicesConfiguration []*Config.ServiceConfiguration

	// The list of data configurations...
	m_datastoreConfiguration []*Config.DataStoreConfiguration

	// synchronize task...
	m_setScheduledTasksChan    chan *Config.ScheduledTask
	m_cancelScheduledTasksChan chan *Config.ScheduledTask
}

var configurationManager *ConfigurationManager

func (this *Server) GetConfigurationManager() *ConfigurationManager {
	if configurationManager == nil {
		configurationManager = newConfigurationManager()
	}
	return configurationManager
}

func newConfigurationManager() *ConfigurationManager {

	configurationManager := new(ConfigurationManager)

	// can be set or not...
	cargoRoot := os.Getenv("CARGOROOT")

	if len(cargoRoot) == 0 {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

		if err != nil {
			panic(err)
		}

		if stat, err := os.Stat(dir + "/WebApp"); err == nil && stat.IsDir() {
			// path is a directory
			cargoRoot = dir
		} else if strings.Index(dir, "CargoWebServer") != -1 {
			cargoRoot = dir[0:strings.Index(dir, "CargoWebServer")]
		}
	}

	// Now I will load the configurations...
	// Development...
	dir, err := filepath.Abs(cargoRoot)

	dir = strings.Replace(dir, "\\", "/", -1)
	if err != nil {
		log.Println(err)
	}

	if strings.HasSuffix(dir, "/") == false {
		dir += "/"
	}

	// Here I will set the root...
	configurationManager.m_filePath = dir + "WebApp/Cargo"
	log.Println("--> server file path is " + configurationManager.m_filePath)
	JS.NewJsRuntimeManager(configurationManager.m_filePath + "/Script")

	// The list of registered services config
	configurationManager.m_servicesConfiguration = make([]*Config.ServiceConfiguration, 0)

	// The list of default datastore.
	configurationManager.m_datastoreConfiguration = make([]*Config.DataStoreConfiguration, 0)
	//serverConfig := configurationManager.m_activeConfigurationsEntity.GetObject().(*Config.Configurations).GetServerConfig()
	hostName := "localhost" //serverConfig.GetHostName()
	ipv4 := "127.0.0.1"     //serverConfig.GetIpv4()
	port := 9393            //serverConfig.GetServerPort()

	// Configuration db itself.
	cargoConfigDB := new(Config.DataStoreConfiguration)
	cargoConfigDB.M_id = ConfigDB
	cargoConfigDB.M_storeName = ConfigDB
	cargoConfigDB.M_hostName = hostName
	cargoConfigDB.M_ipv4 = ipv4
	cargoConfigDB.M_port = port
	cargoConfigDB.M_dataStoreVendor = Config.DataStoreVendor_MYCELIUS
	cargoConfigDB.M_dataStoreType = Config.DataStoreType_KEY_VALUE_STORE
	cargoConfigDB.NeedSave = true
	configurationManager.appendDefaultDataStoreConfiguration(cargoConfigDB)

	// The cargo entities store config
	cargoEntitiesDB := new(Config.DataStoreConfiguration)
	cargoEntitiesDB.M_id = CargoEntitiesDB
	cargoEntitiesDB.M_storeName = CargoEntitiesDB
	cargoEntitiesDB.M_hostName = hostName
	cargoEntitiesDB.M_ipv4 = ipv4
	cargoEntitiesDB.M_port = port
	cargoEntitiesDB.M_dataStoreVendor = Config.DataStoreVendor_MYCELIUS
	cargoEntitiesDB.M_dataStoreType = Config.DataStoreType_KEY_VALUE_STORE
	cargoEntitiesDB.NeedSave = true
	configurationManager.appendDefaultDataStoreConfiguration(cargoEntitiesDB)

	// The sql info data store.
	sqlInfoDB := new(Config.DataStoreConfiguration)
	sqlInfoDB.M_id = "sql_info"
	sqlInfoDB.M_storeName = "sql_info"
	sqlInfoDB.M_hostName = hostName
	sqlInfoDB.M_ipv4 = ipv4
	sqlInfoDB.M_port = port
	sqlInfoDB.M_dataStoreVendor = Config.DataStoreVendor_MYCELIUS
	sqlInfoDB.M_dataStoreType = Config.DataStoreType_KEY_VALUE_STORE
	sqlInfoDB.NeedSave = true
	configurationManager.appendDefaultDataStoreConfiguration(sqlInfoDB)

	// Create the default configurations
	configurationManager.setServiceConfiguration(configurationManager.getId(), -1)

	// Open the channel.
	configurationManager.m_setScheduledTasksChan = make(chan *Config.ScheduledTask)
	configurationManager.m_cancelScheduledTasksChan = make(chan *Config.ScheduledTask)

	return configurationManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Do intialysation stuff here.
 */
func (this *ConfigurationManager) initialize() {
	log.Println("--> initialyze ConfigurationManager")
	// So here if there is no configuration...
	configsUuid := ConfigConfigurationsExists("CARGO_DEFAULT_CONFIGURATIONS")
	if len(configsUuid) > 0 {
		entity, _ := GetServer().GetEntityManager().getEntityByUuid(configsUuid, false)
		this.m_activeConfigurationsEntity = entity.(*Config_ConfigurationsEntity)
	} else {
		activeConfigurations := new(Config.Configurations)
		activeConfigurations.M_id = "CARGO_DEFAULT_CONFIGURATIONS"
		activeConfigurations.M_name = "Cargo Default Configurations"
		activeConfigurations.M_version = "1.0"

		// Now the default server configuration...
		// Sever default values...
		activeConfigurations.M_serverConfig = new(Config.ServerConfiguration)
		activeConfigurations.M_serverConfig.NeedSave = true
		activeConfigurations.M_serverConfig.M_id = "CARGO_DEFAULT_SERVER"
		activeConfigurations.M_serverConfig.M_serverPort = 9393
		activeConfigurations.M_serverConfig.M_ws_serviceContainerPort = 9494
		activeConfigurations.M_serverConfig.M_tcp_serviceContainerPort = 9595
		activeConfigurations.M_serverConfig.M_hostName = "localhost"
		activeConfigurations.M_serverConfig.M_ipv4 = "127.0.0.1"

		// Here I will create the C++ service container TCP | WS
		this.setServiceConfiguration("CargoServiceContainer_TCP", activeConfigurations.M_serverConfig.M_tcp_serviceContainerPort)
		this.setServiceConfiguration("CargoServiceContainer_WS", activeConfigurations.M_serverConfig.M_ws_serviceContainerPort)

		// Scrpit to start the service container.
		tcpServiceContainerStart := new(Config.ScheduledTask)
		tcpServiceContainerStart.M_isActive = true
		tcpServiceContainerStart.M_frequencyType = Config.FrequencyType_ONCE
		tcpServiceContainerStart.M_id = "tcpServiceContainerStart"
		tcpServiceContainerStart.M_script = "tcpServiceContainerStart"

		wsServiceContainerStart := new(Config.ScheduledTask)
		wsServiceContainerStart.M_isActive = true
		wsServiceContainerStart.M_frequencyType = Config.FrequencyType_ONCE
		wsServiceContainerStart.M_id = "wsServiceContainerStart"
		wsServiceContainerStart.M_script = "wsServiceContainerStart"

		// Append the newly create task into the cargo entities
		activeConfigurations.M_scheduledTasks = append(activeConfigurations.M_scheduledTasks, tcpServiceContainerStart)
		activeConfigurations.M_scheduledTasks = append(activeConfigurations.M_scheduledTasks, wsServiceContainerStart)

		// Server folders...
		activeConfigurations.M_serverConfig.M_applicationsPath = "/Apps"
		os.MkdirAll(this.GetApplicationDirectoryPath(), 0777)
		activeConfigurations.M_serverConfig.M_dataPath = "/Data"
		os.MkdirAll(this.GetDataPath(), 0777)
		activeConfigurations.M_serverConfig.M_definitionsPath = "/Definitions"
		os.MkdirAll(this.GetDefinitionsPath(), 0777)
		activeConfigurations.M_serverConfig.M_scriptsPath = "/Script"
		os.MkdirAll(this.GetScriptPath(), 0777)
		activeConfigurations.M_serverConfig.M_schemasPath = "/Schemas"
		os.MkdirAll(this.GetSchemasPath(), 0777)
		activeConfigurations.M_serverConfig.M_tmpPath = "/tmp"
		os.MkdirAll(this.GetTmpPath(), 0777)
		activeConfigurations.M_serverConfig.M_binPath = "/bin"
		os.MkdirAll(this.GetBinPath(), 0777)

		// Where queries are store by default...
		activeConfigurations.NeedSave = true

		// Create the configuration entity from the configuration and save it.
		this.m_activeConfigurationsEntity = GetServer().GetEntityManager().NewConfigConfigurationsEntity("", "", activeConfigurations)
		this.m_activeConfigurationsEntity.SaveEntity()
	}

}

func (this *ConfigurationManager) getId() string {
	return "ConfigurationManager"
}

func (this *ConfigurationManager) start() {
	log.Println("--> Start ConfigurationManager")

	// First of all i will create task if they not exist...
	for i := 0; i < len(this.m_activeConfigurationsEntity.object.M_scheduledTasks); i++ {
		task := this.m_activeConfigurationsEntity.object.M_scheduledTasks[i]
		if task.M_id == "tcpServiceContainerStart" {
			if len(CargoEntitiesFileExists("tcpServiceContainerStart")) == 0 {
				var script string
				script = "function tcpServiceContainerStart(){\n"
				script += `	GetServer().RunCmd("CargoServiceContainer_TCP", ["` + strconv.Itoa(this.GetTcpConfigurationServicePort()) + `"])` + "\n"
				script += `	setInterval(function(){ GetServer().RunCmd("CargoServiceContainer_TCP", ["` + strconv.Itoa(this.GetTcpConfigurationServicePort()) + `"])}, 1000)` + "\n"
				script += "}\n"
				script += "tcpServiceContainerStart()\n"
				GetServer().GetFileManager().createDbFile("tcpServiceContainerStart", "tcpServiceContainerStart.js", "application/javascript", script)
			}
		} else if task.M_id == "wsServiceContainerStart" {
			if len(CargoEntitiesFileExists("wsServiceContainerStart")) == 0 {
				var script string
				script = "function wsServiceContainerStart(){\n"
				script += `	GetServer().RunCmd("CargoServiceContainer_WS", ["` + strconv.Itoa(this.GetWsConfigurationServicePort()) + `"])` + "\n"
				script += `	setInterval(function(){GetServer().RunCmd("CargoServiceContainer_WS", ["` + strconv.Itoa(this.GetWsConfigurationServicePort()) + `"])}, 1000)` + "\n"
				script += "}\n"
				script += "wsServiceContainerStart()\n"
				GetServer().GetFileManager().createDbFile("wsServiceContainerStart", "tcpServiceContainerStart.js", "application/javascript", script)
			}
		}
	}

	// Set services configurations...
	for i := 0; i < len(this.m_servicesConfiguration); i++ {
		serviceUuid := ConfigServiceConfigurationExists(this.m_servicesConfiguration[i].GetId())
		if len(serviceUuid) == 0 {
			// Set the new config...
			activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
			activeConfiguration := activeConfigurationEntity.GetObject().(*Config.Configurations)
			if err == nil {
				activeConfiguration.SetServiceConfigs(this.m_servicesConfiguration[i])
				this.m_activeConfigurationsEntity.SaveEntity()
			} else {
				log.Panicln(err)
			}
		}
	}

	// Set datastores configuration.
	for i := 0; i < len(this.m_datastoreConfiguration); i++ {
		storeUuid := ConfigDataStoreConfigurationExists(this.m_datastoreConfiguration[i].GetId())
		if len(storeUuid) == 0 {
			// Set the new config...
			activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
			activeConfiguration := activeConfigurationEntity.GetObject().(*Config.Configurations)
			if err == nil {
				activeConfiguration.SetDataStoreConfigs(this.m_datastoreConfiguration[i])
				this.m_activeConfigurationsEntity.SaveEntity()
			} else {
				log.Panicln(err)
			}
		}
	}

	////////////////////////////////////////////////////////////////////////////
	// Scheduled Task.
	////////////////////////////////////////////////////////////////////////////
	// Task processing function.
	go func() {
		tasks := make(map[string]*time.Timer, 0)
		for {
			select {
			// set the task.
			case task := <-GetServer().GetConfigurationManager().m_setScheduledTasksChan:
				// Set the timer.
				if tasks[task.GetId()] != nil {
					// Stop previous instance if there is one.
					tasks[task.GetId()].Stop()
				}
				// Plan the next instance.
				startTime := time.Unix(task.GetStartTime(), 0)
				delay := startTime.Sub(time.Now())
				timer := time.NewTimer(delay)
				tasks[task.GetId()] = timer

				go func(task *Config.ScheduledTask, timer *time.Timer) {
					<-timer.C                                           // wait util the delay expire...
					GetServer().GetConfigurationManager().runTask(task) // run the task.
				}(task, timer)

			// reset the task.
			case task := <-GetServer().GetConfigurationManager().m_cancelScheduledTasksChan:
				tasks[task.GetId()].Stop()
				delete(tasks, task.GetId())

			default:
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()

}

func (this *ConfigurationManager) stop() {
	log.Println("--> Stop ConfigurationManager")
}

/**
 * Run a task.
 */
func (this *ConfigurationManager) runTask(task *Config.ScheduledTask) {
	// first of all I will test if the task is active.
	if task.IsActive() == false {
		return // Nothing to do here.
	}

	dbFile, err := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.File", []interface{}{task.M_script}, false)
	if err == nil {
		script, err := b64.StdEncoding.DecodeString(dbFile.GetObject().(*CargoEntities.File).GetData())
		// Now I will run the script...
		if err == nil {
			go func(task *Config.ScheduledTask, script string) {
				_, err := JS.GetJsRuntimeManager().RunScript("", script)
				if err != nil {
					log.Println("---> fail to run task ", task.GetId(), " with error: ", err)
				} else {
					log.Println("--> task ", task.GetId(), "run successfully!")
				}
			}(task, string(script))
		} else {
			log.Println("---> fail to get script for task ", task.GetId(), err)
		}

		if task.GetFrequencyType() != Config.FrequencyType_ONCE {
			// So here I will re-schedule the task, it will not be schedule if is it
			// expired or it must run once.
			GetServer().GetConfigurationManager().scheduleTask(task)
		} else {
			if task.GetStartTime() > 0 {
				GetServer().GetConfigurationManager().scheduleTask(task)
			}
		}

	} else {
		log.Println("---> fail to get db file for task ", task.GetId(), err)
	}

}

// daysIn returns the number of days in a month for a given year.
func daysIn(m time.Month, year int) int {
	// This is equivalent to time.daysIn(m, year).
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

/**
 * That function is use to schedule a task.
 */
func (this *ConfigurationManager) scheduleTask(task *Config.ScheduledTask) {

	// Here I will get the entity for the task.
	entity := GetServer().GetEntityManager().NewConfigScheduledTaskEntityFromObject(task)

	// first of all I will test if the task is active.
	if task.IsActive() == false {
		return // Nothing to do here.
	}

	// If the task is expired
	if task.M_expirationTime > 0 {
		if task.M_expirationTime < time.Now().Unix() {
			// The task has expire!
			entity.GetObject().(*Config.ScheduledTask).SetIsActive(false)
			return
		}
	}

	var nextTime time.Time
	if task.GetFrequencyType() != Config.FrequencyType_ONCE {
		startTime := time.Unix(task.M_startTime, 0)

		// Now I will get the next time when the task must be executed.
		nextTime = startTime
		var previous time.Time
		frequency := task.GetFrequency()
		if frequency == 0 {
			frequency = 1 // Must be one by default if not specify.
		}

		for nextTime.Sub(time.Now()) < 0 {
			previous = nextTime
			// I will append
			if task.GetFrequencyType() == Config.FrequencyType_DAILY {
				f := time.Duration((24 * 60 * 60 * 1000) / frequency)
				nextTime = nextTime.Add(f * time.Millisecond)
			} else if task.GetFrequencyType() == Config.FrequencyType_WEEKELY {
				f := time.Duration((7 * 24 * 60 * 60 * 1000) / frequency)
				nextTime = nextTime.Add(f * time.Millisecond)
			} else if task.GetFrequencyType() == Config.FrequencyType_MONTHLY {
				numberOfDay := daysIn(nextTime.Month(), nextTime.Year())
				f := time.Duration((numberOfDay * 24 * 60 * 60 * 1000) / frequency)
				nextTime = nextTime.Add(f * time.Millisecond)
			}
		}

		// Here I will test if the previous time combine with offset value can
		// be use as nextTime.
		for i := 0; i < len(task.M_offsets); i++ {
			if task.GetFrequencyType() == Config.FrequencyType_WEEKELY || task.GetFrequencyType() == Config.FrequencyType_MONTHLY {
				offset := time.Duration(task.M_offsets[i] * 24 * 60 * 60 * 1000) // in hours
				nextTime_ := previous.Add(offset * time.Millisecond)
				if nextTime_.Sub(time.Now()) > 0 {
					nextTime = nextTime_
					break
				}
			} else if task.GetFrequencyType() == Config.FrequencyType_DAILY {
				// Here the offset represent hours and not days.
				offset := time.Duration(task.M_offsets[i] * 60 * 60 * 1000) // in hours
				nextTime_ := previous.Add(offset * time.Millisecond)
				if nextTime_.Sub(time.Now()) > 0 {
					nextTime = nextTime_
					break
				}
			}
		}

		// Set it start time...
		entity.GetObject().(*Config.ScheduledTask).SetStartTime(nextTime.Unix())

	} else {
		if entity.GetObject().(*Config.ScheduledTask).GetStartTime() == 0 {
			// Run the task directly in that case.
			this.runTask(task)
			return // nothing more to do...
		} else {
			startTime := time.Unix(entity.GetObject().(*Config.ScheduledTask).GetStartTime(), 0)
			if startTime.Sub(time.Now()) < 0 {
				task.NeedSave = true
				// innactivate the task in that case.
				entity.GetObject().(*Config.ScheduledTask).SetIsActive(false)
			}
		}
	}

	// Save modification.
	entity.SaveEntity()

	// Process the task.
	startTime := time.Unix(entity.GetObject().(*Config.ScheduledTask).GetStartTime(), 0)
	if startTime.Sub(time.Now()) >= 0 {
		this.m_setScheduledTasksChan <- entity.GetObject().(*Config.ScheduledTask)
	}
}

/**
 * Return the active configuration.
 */
func (this *ConfigurationManager) getActiveConfigurationsEntity() (*Config_ConfigurationsEntity, error) {
	if this.m_activeConfigurationsEntity != nil {
		activeConfigurationsEntity, err := GetServer().GetEntityManager().getEntityByUuid(this.m_activeConfigurationsEntity.GetUuid(), false)
		if err != nil {
			return nil, errors.New(err.GetBody())
		}
		return activeConfigurationsEntity.(*Config_ConfigurationsEntity), nil
	}

	return nil, errors.New("no active configuration found!")
}

/**
 * Return the OAuth2 configuration entity.
 */
func (this *ConfigurationManager) getOAuthConfigurationEntity() *Config_OAuth2ConfigurationEntity {
	configurationsEntity, err := this.getActiveConfigurationsEntity()
	if err == nil {
		configurations := configurationsEntity.GetObject().(*Config.Configurations)
		oauthConfigurationEntity, err := GetServer().GetEntityManager().getEntityByUuid(configurations.GetOauth2Configuration().GetUUID(), false)
		if err == nil {
			return oauthConfigurationEntity.(*Config_OAuth2ConfigurationEntity)
		}
	}

	// Panic here.
	log.Panicln(err)

	return nil
}

/**
 * Server configuration values...
 */
func (this *ConfigurationManager) GetApplicationDirectoryPath() string {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return this.m_filePath + "/Apps"
	}

	return this.m_filePath + activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_applicationsPath
}

func (this *ConfigurationManager) GetDataPath() string {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return this.m_filePath + "/Data"
	}
	return this.m_filePath + activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_dataPath
}

func (this *ConfigurationManager) GetScriptPath() string {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return this.m_filePath + "/Script"
	}
	return this.m_filePath + activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_scriptsPath
}

func (this *ConfigurationManager) GetDefinitionsPath() string {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return this.m_filePath + "/Definitions"
	}
	return this.m_filePath + activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_definitionsPath
}

func (this *ConfigurationManager) GetSchemasPath() string {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return this.m_filePath + "/Schemas"
	}
	return this.m_filePath + activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_schemasPath
}

func (this *ConfigurationManager) GetTmpPath() string {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return this.m_filePath + "/tmp"
	}
	return this.m_filePath + activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_tmpPath
}

func (this *ConfigurationManager) GetBinPath() string {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return this.m_filePath + "/bin"
	}
	return this.m_filePath + activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_binPath
}

func (this *ConfigurationManager) GetHostName() string {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return "localhost"
	}
	// Default port...
	return activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_hostName
}

func (this *ConfigurationManager) GetIpv4() string {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return "127.0.0.1"
	}
	// Default port...
	return activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_ipv4
}

/**
 * Cargo server port.
 **/
func (this *ConfigurationManager) GetServerPort() int {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return 9393
	}
	return activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_serverPort
}

/**
 * Cargo service container port.
 */
func (this *ConfigurationManager) GetWsConfigurationServicePort() int {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return 9494
	}
	return activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_ws_serviceContainerPort
}

func (this *ConfigurationManager) GetTcpConfigurationServicePort() int {
	activeConfigurationEntity, err := this.getActiveConfigurationsEntity()
	if err != nil {
		return 9595
	}
	return activeConfigurationEntity.GetObject().(*Config.Configurations).M_serverConfig.M_tcp_serviceContainerPort
}

/**
 * Append configuration to the list.
 */
func (this *ConfigurationManager) setServiceConfiguration(id string, port int) {
	// Create the default service configurations
	config := new(Config.ServiceConfiguration)
	config.M_id = id
	config.M_ipv4 = this.GetIpv4()
	config.M_start = true
	if port == -1 {
		config.M_port = this.GetServerPort()
	} else {
		config.M_port = port
	}

	config.M_hostName = this.GetHostName()
	this.m_servicesConfiguration = append(this.m_servicesConfiguration, config)
	return
}

/**
 * Append a default store configurations.
 */
func (this *ConfigurationManager) appendDefaultDataStoreConfiguration(config *Config.DataStoreConfiguration) {
	this.m_datastoreConfiguration = append(this.m_datastoreConfiguration, config)
}

/**
 * Append a datastore config.
 */
func (this *ConfigurationManager) appendDataStoreConfiguration(config *Config.DataStoreConfiguration) {
	// Save the data store.
	activeConfigurationsEntity, err := this.getActiveConfigurationsEntity()
	if err == nil {
		activeConfigurationsEntity.GetObject().(*Config.Configurations).SetDataStoreConfigs(config)
		activeConfigurationsEntity.SaveEntity()
	} else {
		// append in the list of configuration store and save it latter...
		this.m_datastoreConfiguration = append(this.m_datastoreConfiguration, config)
	}
}

/**
 * Return the list of default datastore configurations.
 */
func (this *ConfigurationManager) getDefaultDataStoreConfigurations() []*Config.DataStoreConfiguration {
	return this.m_datastoreConfiguration
}

/**
 * Get the configuration of a given service.
 */
func (this *ConfigurationManager) getServiceConfigurationById(id string) *Config.ServiceConfiguration {

	// Here I will get a look in the list...
	for i := 0; i < len(this.m_servicesConfiguration); i++ {
		if this.m_servicesConfiguration[i].GetId() == id {
			return this.m_servicesConfiguration[i]
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

// @api 1.0
// Schedule a given task.
// @param {string} task The task to schedule.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *ConfigurationManager) ScheduleTask(task *Config.ScheduledTask, messageId string, sessionId string) {
	var errObj *CargoEntities.Error
	errObj = GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}
	this.scheduleTask(task)
}

// @api 1.0
// Cancel the next task instance.
// @param {string} task The task to cancel.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *ConfigurationManager) CancelTask(task *Config.ScheduledTask, messageId string, sessionId string) {
	var errObj *CargoEntities.Error
	errObj = GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	// Cancel the task.
	this.m_cancelScheduledTasksChan <- task

}

// @api 1.0
// Return the current configurations object.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @return {*Config.Configurations}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *ConfigurationManager) GetActiveConfigurations(messageId string, sessionId string) *Config.Configurations {
	return this.m_activeConfigurationsEntity.GetObject().(*Config.Configurations)
}
