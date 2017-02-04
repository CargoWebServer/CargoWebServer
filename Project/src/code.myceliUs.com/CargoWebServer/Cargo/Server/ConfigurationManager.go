package Server

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Config/CargoConfig"
	"code.myceliUs.com/CargoWebServer/Cargo/JS"
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
	m_activeConfigurations *CargoConfig.Configurations

	// the configuration entity...
	m_configurationEntity *CargoConfig_ConfigurationsEntity
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

	log.Println("root dir is: ", cargoRoot)

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

	return configurationManager
}

/**
 * Do intialysation stuff here.
 */
func (this *ConfigurationManager) Initialize() {
	// So here if there is no configuration...
	cargoConfigsUuid := CargoConfigConfigurationsExists("CARGO_CONFIGURATIONS")
	if len(cargoConfigsUuid) > 0 {
		this.m_configurationEntity = GetServer().GetEntityManager().NewCargoConfigConfigurationsEntity(cargoConfigsUuid, nil)
		this.m_configurationEntity.InitEntity(cargoConfigsUuid)
		this.m_activeConfigurations = this.m_configurationEntity.GetObject().(*CargoConfig.Configurations)
	} else {
		this.m_configurationEntity = GetServer().GetEntityManager().NewCargoConfigConfigurationsEntity(cargoConfigsUuid, nil)
		this.m_activeConfigurations = this.m_configurationEntity.GetObject().(*CargoConfig.Configurations)
		this.m_activeConfigurations.M_id = "CARGO_CONFIGURATIONS"
		this.m_activeConfigurations.M_name = "Default"
		this.m_activeConfigurations.M_version = "1.0"
		this.m_activeConfigurations.M_filePath = this.m_filePath

		// Now the default server configuration...
		// Sever default values...
		this.m_activeConfigurations.M_serverConfig = new(CargoConfig.ServerConfiguration)
		this.m_activeConfigurations.M_serverConfig.NeedSave = true

		this.m_activeConfigurations.M_serverConfig.M_id = "CARGO_DEFAULT_SERVER"
		this.m_activeConfigurations.M_serverConfig.M_serverPort = 9393
		this.m_activeConfigurations.M_serverConfig.M_servicePort = 9494
		this.m_activeConfigurations.M_serverConfig.M_hostName = "localhost"
		this.m_activeConfigurations.M_serverConfig.M_ipv4 = "127.0.0.1"

		// Server folders...
		this.m_activeConfigurations.M_serverConfig.M_applicationsPath = "/Apps"
		this.m_activeConfigurations.M_serverConfig.M_dataPath = "/Data"
		this.m_activeConfigurations.M_serverConfig.M_definitionsPath = "/Definitions"
		this.m_activeConfigurations.M_serverConfig.M_scriptsPath = "/Script"
		this.m_activeConfigurations.M_serverConfig.M_schemasPath = "/Schemas"
		this.m_activeConfigurations.M_serverConfig.M_tmpPath = "/tmp"
		this.m_activeConfigurations.M_serverConfig.M_binPath = "/bin"

		this.m_activeConfigurations.NeedSave = true
		this.m_configurationEntity.SaveEntity()
	}
}

func (this *ConfigurationManager) GetId() string {
	return "ConfigurationManager"
}

func (this *ConfigurationManager) Start() {
	log.Println("--> Start ConfigurationManager")
}

func (this *ConfigurationManager) Stop() {
	log.Println("--> Stop ConfigurationManager")
}

/**
 * Server configuration values...
 */
func (this *ConfigurationManager) GetApplicationDirectoryPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/Apps"
	}
	return this.m_filePath + this.m_activeConfigurations.M_serverConfig.M_applicationsPath
}

func (this *ConfigurationManager) GetDataPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/Data"
	}
	return this.m_filePath + this.m_activeConfigurations.M_serverConfig.M_dataPath
}

func (this *ConfigurationManager) GetScriptPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/Script"
	}
	return this.m_filePath + this.m_activeConfigurations.M_serverConfig.M_scriptsPath
}

func (this *ConfigurationManager) GetDefinitionsPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/Definitions"
	}
	return this.m_filePath + this.m_activeConfigurations.M_serverConfig.M_definitionsPath
}

func (this *ConfigurationManager) GetSchemasPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/Schemas"
	}
	return this.m_filePath + this.m_activeConfigurations.M_serverConfig.M_schemasPath
}

func (this *ConfigurationManager) GetTmpPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/tmp"
	}
	return this.m_filePath + this.m_activeConfigurations.M_serverConfig.M_tmpPath
}

func (this *ConfigurationManager) GetBinPath() string {
	if this.m_activeConfigurations == nil {
		return this.m_filePath + "/bin"
	}
	return this.m_filePath + this.m_activeConfigurations.M_serverConfig.M_binPath
}

func (this *ConfigurationManager) GetHostName() string {
	if this.m_activeConfigurations == nil {
		return "localhost"
	}
	// Default port...
	return this.m_activeConfigurations.M_serverConfig.M_hostName
}

func (this *ConfigurationManager) GetIpv4() string {
	if this.m_activeConfigurations == nil {
		return "127.0.0.1"
	}
	// Default port...
	return this.m_activeConfigurations.M_serverConfig.M_ipv4
}

/**
 * Cargo server port.
 **/
func (this *ConfigurationManager) GetServerPort() int {
	// Default port...
	if this.m_activeConfigurations == nil {
		return 9393
	}
	return this.m_activeConfigurations.M_serverConfig.M_serverPort
}

/**
 * Cargo service container port.
 */
func (this *ConfigurationManager) GetServicePort() int {
	// Default port...
	if this.m_activeConfigurations == nil {
		return 9494
	}
	return this.m_activeConfigurations.M_serverConfig.M_servicePort
}

/**
 * Return the list of local services configurations.
 */
func (this *ConfigurationManager) GetLocalServiceConfigurations() []CargoConfig.ServiceConfiguration {
	services := make([]CargoConfig.ServiceConfiguration, 7)

	// Each service that can be start and stop as needed...
	var emailManager CargoConfig.ServiceConfiguration
	emailManager.M_id = "EmailManager"
	emailManager.M_ipv4 = "127.0.0.1"
	emailManager.M_start = true
	emailManager.M_port = 9393
	emailManager.M_hostName = "localhost"
	services[0] = emailManager

	var schemaManager CargoConfig.ServiceConfiguration
	schemaManager.M_id = "SchemaManager"
	schemaManager.M_ipv4 = "127.0.0.1"
	schemaManager.M_start = true
	schemaManager.M_port = 9393
	schemaManager.M_hostName = "localhost"
	services[1] = schemaManager

	var ldapManager CargoConfig.ServiceConfiguration
	ldapManager.M_id = "LdapManager"
	ldapManager.M_ipv4 = "127.0.0.1"
	ldapManager.M_start = false
	ldapManager.M_port = 9393
	ldapManager.M_hostName = "localhost"
	services[2] = ldapManager

	var projectManager CargoConfig.ServiceConfiguration
	projectManager.M_id = "ProjectManager"
	projectManager.M_ipv4 = "127.0.0.1"
	projectManager.M_start = false
	projectManager.M_port = 9393
	projectManager.M_hostName = "localhost"
	services[3] = projectManager

	var fileManager CargoConfig.ServiceConfiguration
	fileManager.M_id = "FileManager"
	fileManager.M_ipv4 = "127.0.0.1"
	fileManager.M_start = false
	fileManager.M_port = 9393
	fileManager.M_hostName = "localhost"
	services[4] = fileManager

	var workflowManager CargoConfig.ServiceConfiguration
	workflowManager.M_id = "WorkflowManager"
	workflowManager.M_ipv4 = "127.0.0.1"
	workflowManager.M_start = true
	workflowManager.M_port = 9393
	workflowManager.M_hostName = "localhost"
	services[5] = workflowManager

	var workflowProcessor CargoConfig.ServiceConfiguration
	workflowProcessor.M_id = "WorkflowProcessor"
	workflowProcessor.M_ipv4 = "127.0.0.1"
	workflowProcessor.M_start = true
	workflowProcessor.M_port = 9393
	workflowProcessor.M_hostName = "localhost"
	services[6] = workflowProcessor

	return services
}

func (this *ConfigurationManager) GetDefaultDataStoreConfigurations() []CargoConfig.DataStoreConfiguration {

	stores := make([]CargoConfig.DataStoreConfiguration, 4)

	// Various persistent entities, account, user, group, file etc...
	var cargoEntitiesDB CargoConfig.DataStoreConfiguration
	cargoEntitiesDB.M_id = CargoEntitiesDB
	cargoEntitiesDB.M_dataStoreVendor = CargoConfig.DataStoreVendor_MYCELIUS
	cargoEntitiesDB.M_dataStoreType = CargoConfig.DataStoreType_KEY_VALUE_STORE
	cargoEntitiesDB.NeedSave = true
	stores[0] = cargoEntitiesDB

	// Configuration entities.
	var cargoConfigDB CargoConfig.DataStoreConfiguration
	cargoConfigDB.M_id = CargoConfigDB
	cargoConfigDB.M_dataStoreVendor = CargoConfig.DataStoreVendor_MYCELIUS
	cargoConfigDB.M_dataStoreType = CargoConfig.DataStoreType_KEY_VALUE_STORE
	cargoConfigDB.NeedSave = true
	stores[1] = cargoConfigDB

	// The BPMN 2.0 entities
	var bpmn20DB CargoConfig.DataStoreConfiguration
	bpmn20DB.M_id = BPMN20DB
	bpmn20DB.M_dataStoreVendor = CargoConfig.DataStoreVendor_MYCELIUS
	bpmn20DB.M_dataStoreType = CargoConfig.DataStoreType_KEY_VALUE_STORE
	bpmn20DB.NeedSave = true
	stores[2] = bpmn20DB

	// The workflow manager runtime entities
	var bpmnRuntimeDB CargoConfig.DataStoreConfiguration
	bpmnRuntimeDB.M_id = BPMS_RuntimeDB
	bpmnRuntimeDB.M_dataStoreVendor = CargoConfig.DataStoreVendor_MYCELIUS
	bpmnRuntimeDB.M_dataStoreType = CargoConfig.DataStoreType_KEY_VALUE_STORE
	bpmnRuntimeDB.NeedSave = true
	stores[3] = bpmnRuntimeDB

	return stores
}

/**
 * Tha function retreive the store data configuration.
 */
func (this *ConfigurationManager) GetDataStoreConfigurations() []CargoConfig.DataStoreConfiguration {
	var configurations []CargoConfig.DataStoreConfiguration

	entities, err := GetServer().GetEntityManager().getEntitiesByType("CargoConfig.DataStoreConfiguration", "", "CargoConfig")
	if err != nil {
		return configurations
	}

	for i := 0; i < len(entities); i++ {
		storeConfiguration := entities[i].GetObject().(*CargoConfig.DataStoreConfiguration)
		configurations = append(configurations, *storeConfiguration)
	}

	return configurations
}

/**
 * Tha function retreive the ldap configuration.
 */
func (this *ConfigurationManager) GetLdapConfigurations() []CargoConfig.LdapConfiguration {
	var configurations []CargoConfig.LdapConfiguration

	entities, err := GetServer().GetEntityManager().getEntitiesByType("CargoConfig.LdapConfiguration", "", "CargoConfig")
	if err != nil {
		return configurations
	}

	for i := 0; i < len(entities); i++ {
		ldapConfiguration := entities[i].GetObject().(*CargoConfig.LdapConfiguration)
		configurations = append(configurations, *ldapConfiguration)
	}

	return configurations
}

/**
 * Tha function retreive the smtp configuration.
 */
func (this *ConfigurationManager) GetSmtpConfigurations() []CargoConfig.SmtpConfiguration {
	var configurations []CargoConfig.SmtpConfiguration

	entities, err := GetServer().GetEntityManager().getEntitiesByType("CargoConfig.SmtpConfiguration", "", "CargoConfig")
	if err != nil {
		return configurations
	}

	for i := 0; i < len(entities); i++ {
		smtpConfiguration := entities[i].GetObject().(*CargoConfig.SmtpConfiguration)
		configurations = append(configurations, *smtpConfiguration)
	}

	return configurations
}

/**
 * Tha function retreive the services configuration.
 */
func (this *ConfigurationManager) GetServiceConfigurations() []CargoConfig.ServiceConfiguration {
	var configurations []CargoConfig.ServiceConfiguration

	entities, err := GetServer().GetEntityManager().getEntitiesByType("CargoConfig.ServiceConfiguration", "", "CargoConfig")
	if err != nil {
		return configurations
	}

	for i := 0; i < len(entities); i++ {
		serviceConfiguration := entities[i].GetObject().(*CargoConfig.ServiceConfiguration)
		configurations = append(configurations, *serviceConfiguration)
	}

	return configurations
}
