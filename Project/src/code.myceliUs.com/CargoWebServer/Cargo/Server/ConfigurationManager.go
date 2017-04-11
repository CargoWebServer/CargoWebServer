package Server

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/CargoWebServer/Cargo/JS"
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

	// The variable CARGOROOT must be set at first...
	cargoRoot := os.Getenv("CARGOROOT")

	// Now I will load the configurations...
	// Development...
	dir, err := filepath.Abs(cargoRoot)

	dir = strings.Replace(dir, "\\", "/", -1)
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasSuffix(dir, "/") == false {
		dir += "/"
	}

	// Here I will set the root...
	configurationManager.m_filePath = dir + "WebApp/Cargo"
	JS.NewJsRuntimeManager(configurationManager.m_filePath + "/Script")

	// The list of registered services config
	configurationManager.m_servicesConfiguration = make([]*Config.ServiceConfiguration, 0)

	// The list of default datastore.
	configurationManager.m_datastoreConfiguration = make([]*Config.DataStoreConfiguration, 0)

	// Configuration db itself.
	cargoConfigDB := new(Config.DataStoreConfiguration)
	cargoConfigDB.M_id = ConfigDB
	cargoConfigDB.M_dataStoreVendor = Config.DataStoreVendor_MYCELIUS
	cargoConfigDB.M_dataStoreType = Config.DataStoreType_KEY_VALUE_STORE
	cargoConfigDB.NeedSave = true
	configurationManager.appendDefaultDataStoreConfiguration(cargoConfigDB)

	// The cargo entities store config
	cargoEntitiesDB := new(Config.DataStoreConfiguration)
	cargoEntitiesDB.M_id = CargoEntitiesDB
	cargoEntitiesDB.M_dataStoreVendor = Config.DataStoreVendor_MYCELIUS
	cargoEntitiesDB.M_dataStoreType = Config.DataStoreType_KEY_VALUE_STORE
	cargoEntitiesDB.NeedSave = true
	configurationManager.appendDefaultDataStoreConfiguration(cargoEntitiesDB)

	// The sql info data store.
	sqlInfoDB := new(Config.DataStoreConfiguration)
	sqlInfoDB.M_id = "sql_info"
	sqlInfoDB.M_dataStoreVendor = Config.DataStoreVendor_MYCELIUS
	sqlInfoDB.M_dataStoreType = Config.DataStoreType_KEY_VALUE_STORE
	sqlInfoDB.NeedSave = true
	configurationManager.appendDefaultDataStoreConfiguration(sqlInfoDB)

	// Create the default configurations
	configurationManager.setServiceConfiguration(configurationManager.getId())

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
	configsUuid := ConfigConfigurationsExists("CARGO_CONFIGURATIONS")
	if len(configsUuid) > 0 {
		entity, _ := GetServer().GetEntityManager().getEntityByUuid(configsUuid)
		this.m_activeConfigurationsEntity = entity.(*Config_ConfigurationsEntity)
	} else {
		activeConfigurations := new(Config.Configurations)
		activeConfigurations.M_id = "CARGO_CONFIGURATIONS"
		activeConfigurations.M_name = "Default"
		activeConfigurations.M_version = "1.0"
		activeConfigurations.M_filePath = this.m_filePath

		// Now the default server configuration...
		// Sever default values...
		activeConfigurations.M_serverConfig = new(Config.ServerConfiguration)
		activeConfigurations.M_serverConfig.NeedSave = true

		activeConfigurations.M_serverConfig.M_id = "CARGO_DEFAULT_SERVER"
		activeConfigurations.M_serverConfig.M_serverPort = 9393
		activeConfigurations.M_serverConfig.M_servicePort = 9494
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

		// Where queries are store by default...
		activeConfigurations.M_serverConfig.M_queriesPath = activeConfigurations.M_serverConfig.M_applicationsPath + "/queries"
		os.MkdirAll(this.GetQueriesPath(), 0777)

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

	// Set services configurations...
	for i := 0; i < len(this.m_servicesConfiguration); i++ {
		serviceUuid := ConfigServiceConfigurationExists(this.m_servicesConfiguration[i].GetId())
		if len(serviceUuid) == 0 {
			// Set the new config...
			this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).SetServiceConfigs(this.m_servicesConfiguration[i])
			this.m_activeConfigurationsEntity.SaveEntity()
		}
	}

	// Set datastores configuration.
	for i := 0; i < len(this.m_datastoreConfiguration); i++ {
		storeUuid := ConfigDataStoreConfigurationExists(this.m_datastoreConfiguration[i].GetId())
		if len(storeUuid) == 0 {
			// Set the new config...
			this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).SetDataStoreConfigs(this.m_datastoreConfiguration[i])
			this.getActiveConfigurationsEntity().SaveEntity()
		}
	}

}

func (this *ConfigurationManager) stop() {
	log.Println("--> Stop ConfigurationManager")
}

/**
 * Return the active configuration.
 */
func (this *ConfigurationManager) getActiveConfigurationsEntity() *Config_ConfigurationsEntity {
	if this.m_activeConfigurationsEntity != nil {
		activeConfigurationsEntity, err := GetServer().GetEntityManager().getEntityByUuid(this.m_activeConfigurationsEntity.GetUuid())
		if err != nil {
			return nil
		}
		return activeConfigurationsEntity.(*Config_ConfigurationsEntity)
	}
	return nil
}

/**
 * Return the OAuth2 configuration entity.
 */
func (this *ConfigurationManager) getOAuthConfigurationEntity() *Config_OAuth2ConfigurationEntity {
	configurationsEntity := this.getActiveConfigurationsEntity()
	if configurationsEntity != nil {
		configurations := configurationsEntity.GetObject().(*Config.Configurations)
		oauthConfigurationEntity, err := GetServer().GetEntityManager().getEntityByUuid(configurations.GetOauth2Configuration().GetUUID())
		if err == nil {
			return oauthConfigurationEntity.(*Config_OAuth2ConfigurationEntity)
		}
	}
	return nil
}

/**
 * Server configuration values...
 */
func (this *ConfigurationManager) GetApplicationDirectoryPath() string {
	if this.getActiveConfigurationsEntity() == nil {
		return this.m_filePath + "/Apps"
	}
	return this.m_filePath + this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_applicationsPath
}

func (this *ConfigurationManager) GetDataPath() string {
	if this.getActiveConfigurationsEntity() == nil {
		return this.m_filePath + "/Data"
	}
	return this.m_filePath + this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_dataPath
}

func (this *ConfigurationManager) GetScriptPath() string {
	if this.getActiveConfigurationsEntity() == nil {
		return this.m_filePath + "/Script"
	}
	return this.m_filePath + this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_scriptsPath
}

func (this *ConfigurationManager) GetDefinitionsPath() string {
	if this.getActiveConfigurationsEntity() == nil {
		return this.m_filePath + "/Definitions"
	}
	return this.m_filePath + this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_definitionsPath
}

func (this *ConfigurationManager) GetSchemasPath() string {
	if this.getActiveConfigurationsEntity() == nil {
		return this.m_filePath + "/Schemas"
	}
	return this.m_filePath + this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_schemasPath
}

func (this *ConfigurationManager) GetTmpPath() string {
	if this.getActiveConfigurationsEntity() == nil {
		return this.m_filePath + "/tmp"
	}
	return this.m_filePath + this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_tmpPath
}

func (this *ConfigurationManager) GetBinPath() string {
	if this.getActiveConfigurationsEntity() == nil {
		return this.m_filePath + "/bin"
	}
	return this.m_filePath + this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_binPath
}

func (this *ConfigurationManager) GetQueriesPath() string {
	if this.getActiveConfigurationsEntity() == nil {
		return this.m_filePath + "/queries"
	}
	return this.m_filePath + this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_queriesPath
}

func (this *ConfigurationManager) GetHostName() string {
	if this.getActiveConfigurationsEntity() == nil {
		return "localhost"
	}
	// Default port...
	return this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_hostName
}

func (this *ConfigurationManager) GetIpv4() string {
	if this.getActiveConfigurationsEntity() == nil {
		return "127.0.0.1"
	}
	// Default port...
	return this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_ipv4
}

/**
 * Cargo server port.
 **/
func (this *ConfigurationManager) GetServerPort() int {
	// Default port...
	if this.getActiveConfigurationsEntity() == nil {
		return 9393
	}
	return this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_serverPort
}

/**
 * Cargo service container port.
 */
func (this *ConfigurationManager) GetServicePort() int {
	// Default port...
	if this.getActiveConfigurationsEntity() == nil {
		return 9494
	}
	return this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations).M_serverConfig.M_servicePort
}

/**
 * Append configuration to the list.
 */
func (this *ConfigurationManager) setServiceConfiguration(id string) {
	// Create the default service configurations
	config := new(Config.ServiceConfiguration)
	config.M_id = id
	config.M_ipv4 = this.GetIpv4()
	config.M_start = true
	config.M_port = this.GetServerPort()
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
	configurations := this.getActiveConfigurationsEntity().GetObject().(*Config.Configurations)
	if configurations != nil {
		configurations.SetDataStoreConfigs(config)
		this.m_activeConfigurationsEntity.SaveEntity()
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
