package Server

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	var activeConfigurations *Config.Configurations
	if len(configsUuid) > 0 {
		entity, _ := GetServer().GetEntityManager().getEntityByUuid(configsUuid, false)
		this.m_activeConfigurationsEntity = entity.(*Config_ConfigurationsEntity)
		activeConfigurations = this.m_activeConfigurationsEntity.GetObject().(*Config.Configurations)
	} else {
		activeConfigurations = new(Config.Configurations)
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

		activeConfigurations.M_serviceConfigs = this.m_servicesConfiguration

		// Where queries are store by default...
		activeConfigurations.NeedSave = true

		// Create the configuration entity from the configuration and save it.
		this.m_activeConfigurationsEntity = GetServer().GetEntityManager().NewConfigConfigurationsEntity("", "", activeConfigurations)
		this.m_activeConfigurationsEntity.SaveEntity()
	}

	// Set the TCP | WS
	this.setServiceConfiguration("CargoServiceContainer_TCP", activeConfigurations.M_serverConfig.M_tcp_serviceContainerPort)
	this.setServiceConfiguration("CargoServiceContainer_WS", activeConfigurations.M_serverConfig.M_ws_serviceContainerPort)
}

func (this *ConfigurationManager) getId() string {
	return "ConfigurationManager"
}

func (this *ConfigurationManager) start() {
	log.Println("--> Start ConfigurationManager")

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
}

func (this *ConfigurationManager) stop() {
	log.Println("--> Stop ConfigurationManager")
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
	config.NeedSave = true

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
	GetTaskManager().scheduleTask(task)
}

// @api 1.0
// Cancel the next task instance.
// @param {string} uuid The instance to cancel.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *ConfigurationManager) CancelTask(uuid string, messageId string, sessionId string) {

	var errObj *CargoEntities.Error
	errObj = GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	// Cancel the task.
	GetTaskManager().m_cancelScheduledTasksChan <- uuid

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

// @api 1.0
// Return the list of task intance in the server memory.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @return {*Config.Configurations}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
// ConfigurationManager.prototype.getTaskInstancesInfos = function (successCallback, errorCallback, caller) {
//    var params = []
//    // Call it on the server.
//    server.executeJsFunction(
//        "ConfigurationManagerGetTaskInstancesInfos", // The function to execute remotely on server
//        params, // The parameters to pass to that function
//        function (index, total, caller) { // The progress callback
//            // Nothing special to do here.
//        },
//        function (results, caller) {
//        	if (caller.successCallback != undefined) {
//        		caller.successCallback(results[0], caller.caller)
//          	caller.successCallback = undefined
//          }
//        },
//        function (errMsg, caller) {
//          server.errorManager.onError(errMsg)
//         	if( caller.errorCallback != undefined){
//          	caller.errorCallback(errMsg, caller.caller)
//				caller.errorCallback = undefined
//			}
//        }, // Error callback
//        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
//    )
//}
func (this *ConfigurationManager) GetTaskInstancesInfos(messageId string, sessionId string) []*TaskInstanceInfo {

	return GetTaskManager().getTaskInstancesInfos()
}
