package Server

import (
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
	//ConfigDB The configuration db
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
	m_activeConfigurations *Config.Configurations

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

	// Default parameters...
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
	cargoConfigDB.M_dataStoreVendor = Config.DataStoreVendor_CAYLEY
	cargoConfigDB.M_dataStoreType = Config.DataStoreType_GRAPH_STORE
	cargoConfigDB.NeedSave = true
	configurationManager.appendDefaultDataStoreConfiguration(cargoConfigDB)

	// The cargo entities store config
	cargoEntitiesDB := new(Config.DataStoreConfiguration)
	cargoEntitiesDB.M_id = CargoEntitiesDB
	cargoEntitiesDB.M_storeName = CargoEntitiesDB
	cargoEntitiesDB.M_hostName = hostName
	cargoEntitiesDB.M_ipv4 = ipv4
	cargoEntitiesDB.M_port = port
	cargoEntitiesDB.M_dataStoreVendor = Config.DataStoreVendor_CAYLEY
	cargoEntitiesDB.M_dataStoreType = Config.DataStoreType_GRAPH_STORE
	cargoEntitiesDB.NeedSave = true
	configurationManager.appendDefaultDataStoreConfiguration(cargoEntitiesDB)

	// The sql info data store.
	sqlInfoDB := new(Config.DataStoreConfiguration)
	sqlInfoDB.M_id = "sql_info"
	sqlInfoDB.M_storeName = "sql_info"
	sqlInfoDB.M_hostName = hostName
	sqlInfoDB.M_ipv4 = ipv4
	sqlInfoDB.M_port = port
	sqlInfoDB.M_dataStoreVendor = Config.DataStoreVendor_CAYLEY
	sqlInfoDB.M_dataStoreType = Config.DataStoreType_GRAPH_STORE
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
	entity, cargoError := GetServer().GetEntityManager().getEntityById("Config.Configurations", "Config", []interface{}{"CARGO_DEFAULT_CONFIGURATIONS"})
	if cargoError == nil {
		this.m_activeConfigurations = entity.(*Config.Configurations)
	} else {
		// Create directory if there are not already there.
		os.MkdirAll(this.GetApplicationDirectoryPath(), 0777)
		os.MkdirAll(this.GetDataPath(), 0777)
		os.MkdirAll(this.GetDefinitionsPath(), 0777)
		os.MkdirAll(this.GetScriptPath(), 0777)
		os.MkdirAll(this.GetSchemasPath(), 0777)
		os.MkdirAll(this.GetTmpPath(), 0777)
		os.MkdirAll(this.GetBinPath(), 0777)

		this.m_activeConfigurations = new(Config.Configurations)
		this.m_activeConfigurations.M_id = "CARGO_DEFAULT_CONFIGURATIONS"
		this.m_activeConfigurations.M_name = "Cargo Default Configurations"
		this.m_activeConfigurations.M_version = "1.0"

		// Where queries are store by default...
		this.m_activeConfigurations.NeedSave = true

		// Create the configuration entity from the configuration and save it.
		GetServer().GetEntityManager().saveEntity(this.m_activeConfigurations)

		this.m_activeConfigurations.SetServiceConfigs(this.m_servicesConfiguration)

		// Now the default server configuration...
		// Sever default values...
		var serverConfig = new(Config.ServerConfiguration)
		serverConfig = new(Config.ServerConfiguration)
		serverConfig.NeedSave = true
		serverConfig.M_id = "CARGO_DEFAULT_SERVER"
		serverConfig.M_serverPort = 9393
		serverConfig.M_serviceContainerPort = 9494
		serverConfig.M_hostName = "localhost"
		serverConfig.M_ipv4 = "127.0.0.1"

		// Server folders...
		serverConfig.M_applicationsPath = "/Apps"
		serverConfig.M_dataPath = "/Data"
		serverConfig.M_definitionsPath = "/Definitions"
		serverConfig.M_scriptsPath = "/Script"
		serverConfig.M_schemasPath = "/Schemas"
		serverConfig.M_tmpPath = "/tmp"
		serverConfig.M_binPath = "/bin"

		GetServer().GetEntityManager().createEntity(this.m_activeConfigurations.GetUuid(), "M_serverConfig", "Config.ServerConfiguration", serverConfig.GetId(), serverConfig)
		this.m_activeConfigurations.SetServerConfig(serverConfig)
	}

	// Set the service container configuration
	this.setServiceConfiguration("CargoServiceContainer", this.m_activeConfigurations.GetServerConfig().GetServiceContainerPort())
}

func (this *ConfigurationManager) getId() string {
	return "ConfigurationManager"
}

func (this *ConfigurationManager) start() {
	log.Println("--> Start ConfigurationManager")
	// Set services configurations...
	for i := 0; i < len(this.m_servicesConfiguration); i++ {
		_, err := GetServer().GetEntityManager().getEntityById("Config.ServiceConfiguration", "Config", []interface{}{this.m_servicesConfiguration[i].GetId()})
		if err != nil {
			// Set the new config...
			GetServer().GetEntityManager().createEntity(this.m_activeConfigurations.GetUuid(), "M_serviceConfigs", "Config.ServiceConfiguration", this.m_servicesConfiguration[i].GetId(), this.m_servicesConfiguration[i])
		}
	}

	// Set datastores configuration.
	for i := 0; i < len(this.m_datastoreConfiguration); i++ {
		_, err := GetServer().GetEntityManager().getEntityById("Config.DataStoreConfiguration", "Config", []interface{}{this.m_datastoreConfiguration[i].GetId()})
		if err != nil {
			// Set the new config...
			GetServer().GetEntityManager().createEntity(this.m_activeConfigurations.GetUuid(), "M_dataStoreConfigs", "Config.DataStoreConfiguration", this.m_datastoreConfiguration[i].GetId(), this.m_datastoreConfiguration[i])
		}
	}
}

func (this *ConfigurationManager) stop() {
	log.Println("--> Stop ConfigurationManager")
}

/**
 * Return the OAuth2 configuration entity.
 */
func (this *ConfigurationManager) getOAuthConfigurationEntity() *Config.OAuth2Configuration {

	oauthConfiguration, err := GetServer().GetEntityManager().getEntityByUuid(this.m_activeConfigurations.GetOauth2Configuration().GetUuid())
	if err == nil {
		return oauthConfiguration.(*Config.OAuth2Configuration)
	}

	return nil
}

/**
 * Server configuration values...
 */
func (this *ConfigurationManager) GetApplicationDirectoryPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/Apps"
	}

	return this.m_filePath + this.m_activeConfigurations.GetServerConfig().GetApplicationsPath()
}

func (this *ConfigurationManager) GetDataPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/Data"
	}
	return this.m_filePath + this.m_activeConfigurations.GetServerConfig().GetDataPath()
}

func (this *ConfigurationManager) GetScriptPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/Script"
	}
	return this.m_filePath + this.m_activeConfigurations.GetServerConfig().GetScriptsPath()
}

func (this *ConfigurationManager) GetDefinitionsPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/Definitions"
	}
	return this.m_filePath + this.m_activeConfigurations.GetServerConfig().GetDefinitionsPath()
}

func (this *ConfigurationManager) GetSchemasPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/Schemas"
	}
	return this.m_filePath + this.m_activeConfigurations.GetServerConfig().GetSchemasPath()
}

func (this *ConfigurationManager) GetTmpPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/tmp"
	}
	return this.m_filePath + this.m_activeConfigurations.GetServerConfig().GetTmpPath()
}

func (this *ConfigurationManager) GetBinPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/bin"
	}
	return this.m_filePath + this.m_activeConfigurations.GetServerConfig().GetBinPath()
}

func (this *ConfigurationManager) GetHostName() string {
	if this.m_activeConfigurations == nil {
		return "localhost"
	}
	// Default port...
	return this.m_activeConfigurations.GetServerConfig().GetHostName()
}

func (this *ConfigurationManager) GetIpv4() string {
	if this.m_activeConfigurations == nil {
		return "127.0.0.1"
	}
	// Default port...
	return this.m_activeConfigurations.GetServerConfig().GetIpv4()
}

/**
 * Cargo server port.
 **/
func (this *ConfigurationManager) GetServerPort() int {
	if this.m_activeConfigurations == nil {
		return 9393
	}
	return this.m_activeConfigurations.GetServerConfig().GetServerPort()
}

/**
 * Cargo service container port.
 */
func (this *ConfigurationManager) GetConfigurationServicePort() int {
	if this.m_activeConfigurations == nil {
		return 9494
	}
	return this.m_activeConfigurations.GetServerConfig().GetServiceContainerPort()
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
	if this.m_activeConfigurations != nil {
		this.m_activeConfigurations.AppendDataStoreConfigs(config)
		GetServer().GetEntityManager().saveEntity(this.m_activeConfigurations)
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
	return this.m_activeConfigurations
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
