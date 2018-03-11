// +build Config

package Entities

import (
	"encoding/xml"
	"log"
	"os"
	"path/filepath"
	//	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
	"golang.org/x/net/html/charset"
)

type ConfigXmlFactory struct {
	m_references      map[string]interface{}
	m_object          map[string]map[string][]string
	m_getEntityByUuid func(string) (interface{}, error)
	m_setEntity       func(interface{})
	m_generateUuid    func(interface{}) string
}

/** Get entity by uuid function **/
func (this *ConfigXmlFactory) SetEntityGetter(fct func(string) (interface{}, error)) {
	this.m_getEntityByUuid = fct
}

/** Use to put the entity in the cache **/
func (this *ConfigXmlFactory) SetEntitySetter(fct func(interface{})) {
	this.m_setEntity = fct
}

/** Set the uuid genarator function. **/
func (this *ConfigXmlFactory) SetUuidGenerator(fct func(entity interface{}) string) {
	this.m_generateUuid = fct
}

/** Initialization function from xml file **/
func (this *ConfigXmlFactory) InitXml(inputPath string, object *Config.Configurations) error {
	xmlFilePath, err := filepath.Abs(inputPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	reader, err := os.Open(xmlFilePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	var xmlElement *Config.XsdConfigurations
	xmlElement = new(Config.XsdConfigurations)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(xmlElement); err != nil {
		return err
	}
	this.m_references = make(map[string]interface{}, 0)
	this.m_object = make(map[string]map[string][]string, 0)
	this.InitConfigurations("", xmlElement, object)
	for ref0, refMap := range this.m_object {
		refOwner := this.m_references[ref0]
		if refOwner != nil {
			for ref1, _ := range refMap {
				refs := refMap[ref1]
				for i := 0; i < len(refs); i++ {
					ref := this.m_references[refs[i]]
					if ref != nil {
						params := make([]interface{}, 0)
						params = append(params, ref)
						Utility.CallMethod(refOwner, ref1, params)
					} else {
						params := make([]interface{}, 0)
						params = append(params, refs[i])
						Utility.CallMethod(refOwner, ref1, params)
					}
				}
			}
		}
	}
	return nil
}

/** Serialization to xml file **/
func (this *ConfigXmlFactory) SerializeXml(outputPath string, toSerialize *Config.Configurations) error {
	xmlFilePath, err := filepath.Abs(outputPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fo, err := os.Create(xmlFilePath)
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	var xmlElement *Config.XsdConfigurations
	xmlElement = new(Config.XsdConfigurations)

	this.SerialyzeConfigurations(xmlElement, toSerialize)
	output, err := xml.MarshalIndent(xmlElement, "  ", "    ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fileContent := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n")
	fileContent = append(fileContent, output...)
	_, err = fo.Write(fileContent)
	return nil
}

/** inititialisation of ServiceConfiguration **/
func (this *ConfigXmlFactory) InitServiceConfiguration(parentUuid string, xmlElement *Config.XsdServiceConfiguration, object *Config.ServiceConfiguration) {
	log.Println("Initialize ServiceConfiguration")
	object.TYPENAME = "Config.ServiceConfiguration"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Configuration **/
	object.M_id = xmlElement.M_id

	/** ServiceConfiguration **/
	object.M_hostName = xmlElement.M_hostName

	/** ServiceConfiguration **/
	object.M_ipv4 = xmlElement.M_ipv4

	/** ServiceConfiguration **/
	object.M_port = xmlElement.M_port

	/** ServiceConfiguration **/
	object.M_user = xmlElement.M_user

	/** ServiceConfiguration **/
	object.M_pwd = xmlElement.M_pwd

	/** ServiceConfiguration **/
	object.M_start = xmlElement.M_start
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ScheduledTask **/
func (this *ConfigXmlFactory) InitScheduledTask(parentUuid string, xmlElement *Config.XsdScheduledTask, object *Config.ScheduledTask) {
	log.Println("Initialize ScheduledTask")
	object.TYPENAME = "Config.ScheduledTask"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Configuration **/
	object.M_id = xmlElement.M_id

	/** ScheduledTask **/
	object.M_isActive = xmlElement.M_isActive

	/** ScheduledTask **/
	object.M_script = xmlElement.M_script

	/** ScheduledTask **/
	object.M_startTime = xmlElement.M_startTime

	/** ScheduledTask **/
	object.M_expirationTime = xmlElement.M_expirationTime

	/** ScheduledTask **/
	object.M_frequency = xmlElement.M_frequency

	if xmlElement.M_frequencyType == "##ONCE" {
		object.M_frequencyType = Config.FrequencyType_ONCE
	} else if xmlElement.M_frequencyType == "##DAILY" {
		object.M_frequencyType = Config.FrequencyType_DAILY
	} else if xmlElement.M_frequencyType == "##WEEKELY" {
		object.M_frequencyType = Config.FrequencyType_WEEKELY
	} else if xmlElement.M_frequencyType == "##MONTHLY" {
		object.M_frequencyType = Config.FrequencyType_MONTHLY
	}

	/** ScheduledTask **/
	object.M_offsets = xmlElement.M_offsets
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ServerConfiguration **/
func (this *ConfigXmlFactory) InitServerConfiguration(parentUuid string, xmlElement *Config.XsdServerConfiguration, object *Config.ServerConfiguration) {
	log.Println("Initialize ServerConfiguration")
	object.TYPENAME = "Config.ServerConfiguration"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Configuration **/
	object.M_id = xmlElement.M_id

	/** ServerConfiguration **/
	object.M_ipv4 = xmlElement.M_ipv4

	/** ServerConfiguration **/
	object.M_hostName = xmlElement.M_hostName

	/** ServerConfiguration **/
	object.M_serverPort = xmlElement.M_serverPort

	/** ServerConfiguration **/
	object.M_serviceContainerPort = xmlElement.M_serviceContainerPort

	/** ServerConfiguration **/
	object.M_applicationsPath = xmlElement.M_applicationsPath

	/** ServerConfiguration **/
	object.M_dataPath = xmlElement.M_dataPath

	/** ServerConfiguration **/
	object.M_scriptsPath = xmlElement.M_scriptsPath

	/** ServerConfiguration **/
	object.M_definitionsPath = xmlElement.M_definitionsPath

	/** ServerConfiguration **/
	object.M_schemasPath = xmlElement.M_schemasPath

	/** ServerConfiguration **/
	object.M_tmpPath = xmlElement.M_tmpPath

	/** ServerConfiguration **/
	object.M_binPath = xmlElement.M_binPath

	/** ServerConfiguration **/
	object.M_shards = xmlElement.M_shards

	/** ServerConfiguration **/
	object.M_lifeWindow = xmlElement.M_lifeWindow

	/** ServerConfiguration **/
	object.M_maxEntriesInWindow = xmlElement.M_maxEntriesInWindow

	/** ServerConfiguration **/
	object.M_maxEntrySize = xmlElement.M_maxEntrySize

	/** ServerConfiguration **/
	object.M_verbose = xmlElement.M_verbose

	/** ServerConfiguration **/
	object.M_hardMaxCacheSize = xmlElement.M_hardMaxCacheSize
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ApplicationConfiguration **/
func (this *ConfigXmlFactory) InitApplicationConfiguration(parentUuid string, xmlElement *Config.XsdApplicationConfiguration, object *Config.ApplicationConfiguration) {
	log.Println("Initialize ApplicationConfiguration")
	object.TYPENAME = "Config.ApplicationConfiguration"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Configuration **/
	object.M_id = xmlElement.M_id

	/** ApplicationConfiguration **/
	object.M_indexPage = xmlElement.M_indexPage
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of LdapConfiguration **/
func (this *ConfigXmlFactory) InitLdapConfiguration(parentUuid string, xmlElement *Config.XsdLdapConfiguration, object *Config.LdapConfiguration) {
	log.Println("Initialize LdapConfiguration")
	object.TYPENAME = "Config.LdapConfiguration"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Configuration **/
	object.M_id = xmlElement.M_id

	/** LdapConfiguration **/
	object.M_hostName = xmlElement.M_hostName

	/** LdapConfiguration **/
	object.M_ipv4 = xmlElement.M_ipv4

	/** LdapConfiguration **/
	object.M_port = xmlElement.M_port

	/** LdapConfiguration **/
	object.M_user = xmlElement.M_user

	/** LdapConfiguration **/
	object.M_pwd = xmlElement.M_pwd

	/** LdapConfiguration **/
	object.M_domain = xmlElement.M_domain

	/** LdapConfiguration **/
	object.M_searchBase = xmlElement.M_searchBase
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of OAuth2Configuration **/
func (this *ConfigXmlFactory) InitOAuth2Configuration(parentUuid string, xmlElement *Config.XsdOAuth2Configuration, object *Config.OAuth2Configuration) {
	log.Println("Initialize OAuth2Configuration")
	object.TYPENAME = "Config.OAuth2Configuration"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Configuration **/
	object.M_id = xmlElement.M_id

	/** OAuth2Configuration **/
	object.M_authorizationExpiration = xmlElement.M_authorizationExpiration

	/** OAuth2Configuration **/
	object.M_accessExpiration = xmlElement.M_accessExpiration

	/** OAuth2Configuration **/
	object.M_tokenType = xmlElement.M_tokenType

	/** OAuth2Configuration **/
	object.M_errorStatusCode = xmlElement.M_errorStatusCode

	/** OAuth2Configuration **/
	object.M_allowClientSecretInParams = xmlElement.M_allowClientSecretInParams

	/** OAuth2Configuration **/
	object.M_allowGetAccessRequest = xmlElement.M_allowGetAccessRequest

	/** OAuth2Configuration **/
	object.M_redirectUriSeparator = xmlElement.M_redirectUriSeparator

	/** OAuth2Configuration **/
	object.M_privateKey = xmlElement.M_privateKey
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Configurations **/
func (this *ConfigXmlFactory) InitConfigurations(parentUuid string, xmlElement *Config.XsdConfigurations, object *Config.Configurations) {
	log.Println("Initialize Configurations")
	object.TYPENAME = "Config.Configurations"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Configurations **/
	object.M_id = xmlElement.M_id

	/** Configurations **/
	object.M_name = xmlElement.M_name

	/** Configurations **/
	object.M_version = xmlElement.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init serverConfiguration **/
	val := new(Config.ServerConfiguration)
	this.InitServerConfiguration(object.GetUuid(), &xmlElement.M_serverConfig, val)
	object.SetServerConfig(val)

	/** association initialisation **/

	/** Init applicationConfiguration **/
	for i := 0; i < len(xmlElement.M_applicationConfigs); i++ {
		if object.M_applicationConfigs == nil {
			object.M_applicationConfigs = make([]string, 0)
		}
		val := new(Config.ApplicationConfiguration)
		this.InitApplicationConfiguration(object.GetUuid(), xmlElement.M_applicationConfigs[i], val)
		object.AppendApplicationConfigs(val)

		/** association initialisation **/
	}

	/** Init smtpConfiguration **/
	for i := 0; i < len(xmlElement.M_smtpConfigs); i++ {
		if object.M_smtpConfigs == nil {
			object.M_smtpConfigs = make([]string, 0)
		}
		val := new(Config.SmtpConfiguration)
		this.InitSmtpConfiguration(object.GetUuid(), xmlElement.M_smtpConfigs[i], val)
		object.AppendSmtpConfigs(val)

		/** association initialisation **/
	}

	/** Init ldapConfiguration **/
	for i := 0; i < len(xmlElement.M_ldapConfigs); i++ {
		if object.M_ldapConfigs == nil {
			object.M_ldapConfigs = make([]string, 0)
		}
		val := new(Config.LdapConfiguration)
		this.InitLdapConfiguration(object.GetUuid(), xmlElement.M_ldapConfigs[i], val)
		object.AppendLdapConfigs(val)

		/** association initialisation **/
	}

	/** Init dataStoreConfiguration **/
	for i := 0; i < len(xmlElement.M_dataStoreConfigs); i++ {
		if object.M_dataStoreConfigs == nil {
			object.M_dataStoreConfigs = make([]string, 0)
		}
		val := new(Config.DataStoreConfiguration)
		this.InitDataStoreConfiguration(object.GetUuid(), xmlElement.M_dataStoreConfigs[i], val)
		object.AppendDataStoreConfigs(val)

		/** association initialisation **/
	}

	/** Init serviceConfiguration **/
	for i := 0; i < len(xmlElement.M_serviceConfigs); i++ {
		if object.M_serviceConfigs == nil {
			object.M_serviceConfigs = make([]string, 0)
		}
		val := new(Config.ServiceConfiguration)
		this.InitServiceConfiguration(object.GetUuid(), xmlElement.M_serviceConfigs[i], val)
		object.AppendServiceConfigs(val)

		/** association initialisation **/
	}

	/** Init oauth2Configuration **/
	if xmlElement.M_oauth2Configuration != nil {
		val := new(Config.OAuth2Configuration)
		this.InitOAuth2Configuration(object.GetUuid(), xmlElement.M_oauth2Configuration, val)
		object.SetOauth2Configuration(val)

		/** association initialisation **/
	}

	/** Init scheduledTask **/
	for i := 0; i < len(xmlElement.M_scheduledTasks); i++ {
		if object.M_scheduledTasks == nil {
			object.M_scheduledTasks = make([]string, 0)
		}
		val := new(Config.ScheduledTask)
		this.InitScheduledTask(object.GetUuid(), xmlElement.M_scheduledTasks[i], val)
		object.AppendScheduledTasks(val)

		/** association initialisation **/
	}
}

/** inititialisation of SmtpConfiguration **/
func (this *ConfigXmlFactory) InitSmtpConfiguration(parentUuid string, xmlElement *Config.XsdSmtpConfiguration, object *Config.SmtpConfiguration) {
	log.Println("Initialize SmtpConfiguration")
	object.TYPENAME = "Config.SmtpConfiguration"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Configuration **/
	object.M_id = xmlElement.M_id

	/** SmtpConfiguration **/
	object.M_hostName = xmlElement.M_hostName

	/** SmtpConfiguration **/
	object.M_ipv4 = xmlElement.M_ipv4

	/** SmtpConfiguration **/
	object.M_port = xmlElement.M_port

	/** SmtpConfiguration **/
	object.M_user = xmlElement.M_user

	/** SmtpConfiguration **/
	object.M_pwd = xmlElement.M_pwd

	if xmlElement.M_textEncoding == "##UTF8" {
		object.M_textEncoding = Config.Encoding_UTF8
	} else if xmlElement.M_textEncoding == "##WINDOWS_1250" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1250
	} else if xmlElement.M_textEncoding == "##WINDOWS_1251" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1251
	} else if xmlElement.M_textEncoding == "##WINDOWS_1252" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1252
	} else if xmlElement.M_textEncoding == "##WINDOWS_1253" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1253
	} else if xmlElement.M_textEncoding == "##WINDOWS_1254" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1254
	} else if xmlElement.M_textEncoding == "##WINDOWS_1255" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1255
	} else if xmlElement.M_textEncoding == "##WINDOWS_1256" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1256
	} else if xmlElement.M_textEncoding == "##WINDOWS_1257" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1257
	} else if xmlElement.M_textEncoding == "##WINDOWS_1258" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1258
	} else if xmlElement.M_textEncoding == "##ISO8859_1" {
		object.M_textEncoding = Config.Encoding_ISO8859_1
	} else if xmlElement.M_textEncoding == "##ISO8859_2" {
		object.M_textEncoding = Config.Encoding_ISO8859_2
	} else if xmlElement.M_textEncoding == "##ISO8859_3" {
		object.M_textEncoding = Config.Encoding_ISO8859_3
	} else if xmlElement.M_textEncoding == "##ISO8859_4" {
		object.M_textEncoding = Config.Encoding_ISO8859_4
	} else if xmlElement.M_textEncoding == "##ISO8859_5" {
		object.M_textEncoding = Config.Encoding_ISO8859_5
	} else if xmlElement.M_textEncoding == "##ISO8859_6" {
		object.M_textEncoding = Config.Encoding_ISO8859_6
	} else if xmlElement.M_textEncoding == "##ISO8859_7" {
		object.M_textEncoding = Config.Encoding_ISO8859_7
	} else if xmlElement.M_textEncoding == "##ISO8859_8" {
		object.M_textEncoding = Config.Encoding_ISO8859_8
	} else if xmlElement.M_textEncoding == "##ISO8859_9" {
		object.M_textEncoding = Config.Encoding_ISO8859_9
	} else if xmlElement.M_textEncoding == "##ISO8859_10" {
		object.M_textEncoding = Config.Encoding_ISO8859_10
	} else if xmlElement.M_textEncoding == "##ISO8859_13" {
		object.M_textEncoding = Config.Encoding_ISO8859_13
	} else if xmlElement.M_textEncoding == "##ISO8859_14" {
		object.M_textEncoding = Config.Encoding_ISO8859_14
	} else if xmlElement.M_textEncoding == "##ISO8859_15" {
		object.M_textEncoding = Config.Encoding_ISO8859_15
	} else if xmlElement.M_textEncoding == "##ISO8859_16" {
		object.M_textEncoding = Config.Encoding_ISO8859_16
	} else if xmlElement.M_textEncoding == "##KOI8R" {
		object.M_textEncoding = Config.Encoding_KOI8R
	} else if xmlElement.M_textEncoding == "##KOI8U" {
		object.M_textEncoding = Config.Encoding_KOI8U
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of DataStoreConfiguration **/
func (this *ConfigXmlFactory) InitDataStoreConfiguration(parentUuid string, xmlElement *Config.XsdDataStoreConfiguration, object *Config.DataStoreConfiguration) {
	log.Println("Initialize DataStoreConfiguration")
	object.TYPENAME = "Config.DataStoreConfiguration"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Configuration **/
	object.M_id = xmlElement.M_id

	/** DataStoreConfiguration **/
	object.M_storeName = xmlElement.M_storeName

	/** DataStoreConfiguration **/
	object.M_hostName = xmlElement.M_hostName

	/** DataStoreConfiguration **/
	object.M_ipv4 = xmlElement.M_ipv4

	/** DataStoreConfiguration **/
	object.M_port = xmlElement.M_port

	/** DataStoreConfiguration **/
	object.M_user = xmlElement.M_user

	/** DataStoreConfiguration **/
	object.M_pwd = xmlElement.M_pwd

	if xmlElement.M_dataStoreType == "##SQL_STORE" {
		object.M_dataStoreType = Config.DataStoreType_SQL_STORE
	} else if xmlElement.M_dataStoreType == "##GRAPH_STORE" {
		object.M_dataStoreType = Config.DataStoreType_GRAPH_STORE
	}

	if xmlElement.M_dataStoreVendor == "##CARGO" {
		object.M_dataStoreVendor = Config.DataStoreVendor_CARGO
	} else if xmlElement.M_dataStoreVendor == "##MYSQL" {
		object.M_dataStoreVendor = Config.DataStoreVendor_MYSQL
	} else if xmlElement.M_dataStoreVendor == "##MSSQL" {
		object.M_dataStoreVendor = Config.DataStoreVendor_MSSQL
	}

	if xmlElement.M_textEncoding == "##UTF8" {
		object.M_textEncoding = Config.Encoding_UTF8
	} else if xmlElement.M_textEncoding == "##WINDOWS_1250" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1250
	} else if xmlElement.M_textEncoding == "##WINDOWS_1251" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1251
	} else if xmlElement.M_textEncoding == "##WINDOWS_1252" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1252
	} else if xmlElement.M_textEncoding == "##WINDOWS_1253" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1253
	} else if xmlElement.M_textEncoding == "##WINDOWS_1254" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1254
	} else if xmlElement.M_textEncoding == "##WINDOWS_1255" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1255
	} else if xmlElement.M_textEncoding == "##WINDOWS_1256" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1256
	} else if xmlElement.M_textEncoding == "##WINDOWS_1257" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1257
	} else if xmlElement.M_textEncoding == "##WINDOWS_1258" {
		object.M_textEncoding = Config.Encoding_WINDOWS_1258
	} else if xmlElement.M_textEncoding == "##ISO8859_1" {
		object.M_textEncoding = Config.Encoding_ISO8859_1
	} else if xmlElement.M_textEncoding == "##ISO8859_2" {
		object.M_textEncoding = Config.Encoding_ISO8859_2
	} else if xmlElement.M_textEncoding == "##ISO8859_3" {
		object.M_textEncoding = Config.Encoding_ISO8859_3
	} else if xmlElement.M_textEncoding == "##ISO8859_4" {
		object.M_textEncoding = Config.Encoding_ISO8859_4
	} else if xmlElement.M_textEncoding == "##ISO8859_5" {
		object.M_textEncoding = Config.Encoding_ISO8859_5
	} else if xmlElement.M_textEncoding == "##ISO8859_6" {
		object.M_textEncoding = Config.Encoding_ISO8859_6
	} else if xmlElement.M_textEncoding == "##ISO8859_7" {
		object.M_textEncoding = Config.Encoding_ISO8859_7
	} else if xmlElement.M_textEncoding == "##ISO8859_8" {
		object.M_textEncoding = Config.Encoding_ISO8859_8
	} else if xmlElement.M_textEncoding == "##ISO8859_9" {
		object.M_textEncoding = Config.Encoding_ISO8859_9
	} else if xmlElement.M_textEncoding == "##ISO8859_10" {
		object.M_textEncoding = Config.Encoding_ISO8859_10
	} else if xmlElement.M_textEncoding == "##ISO8859_13" {
		object.M_textEncoding = Config.Encoding_ISO8859_13
	} else if xmlElement.M_textEncoding == "##ISO8859_14" {
		object.M_textEncoding = Config.Encoding_ISO8859_14
	} else if xmlElement.M_textEncoding == "##ISO8859_15" {
		object.M_textEncoding = Config.Encoding_ISO8859_15
	} else if xmlElement.M_textEncoding == "##ISO8859_16" {
		object.M_textEncoding = Config.Encoding_ISO8859_16
	} else if xmlElement.M_textEncoding == "##KOI8R" {
		object.M_textEncoding = Config.Encoding_KOI8R
	} else if xmlElement.M_textEncoding == "##KOI8U" {
		object.M_textEncoding = Config.Encoding_KOI8U
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of OAuth2Configuration **/
func (this *ConfigXmlFactory) SerialyzeOAuth2Configuration(xmlElement *Config.XsdOAuth2Configuration, object *Config.OAuth2Configuration) {
	if xmlElement == nil {
		return
	}

	/** Configuration **/
	xmlElement.M_id = object.M_id

	/** OAuth2Configuration **/
	xmlElement.M_authorizationExpiration = object.M_authorizationExpiration

	/** OAuth2Configuration **/
	xmlElement.M_accessExpiration = object.M_accessExpiration

	/** OAuth2Configuration **/
	xmlElement.M_tokenType = object.M_tokenType

	/** OAuth2Configuration **/
	xmlElement.M_errorStatusCode = object.M_errorStatusCode

	/** OAuth2Configuration **/
	xmlElement.M_allowClientSecretInParams = object.M_allowClientSecretInParams

	/** OAuth2Configuration **/
	xmlElement.M_allowGetAccessRequest = object.M_allowGetAccessRequest

	/** OAuth2Configuration **/
	xmlElement.M_redirectUriSeparator = object.M_redirectUriSeparator

	/** OAuth2Configuration **/
	xmlElement.M_privateKey = object.M_privateKey
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ScheduledTask **/
func (this *ConfigXmlFactory) SerialyzeScheduledTask(xmlElement *Config.XsdScheduledTask, object *Config.ScheduledTask) {
	if xmlElement == nil {
		return
	}

	/** Configuration **/
	xmlElement.M_id = object.M_id

	/** ScheduledTask **/
	xmlElement.M_isActive = object.M_isActive

	/** ScheduledTask **/
	xmlElement.M_script = object.M_script

	/** ScheduledTask **/
	xmlElement.M_startTime = object.M_startTime

	/** ScheduledTask **/
	xmlElement.M_expirationTime = object.M_expirationTime

	/** ScheduledTask **/
	xmlElement.M_frequency = object.M_frequency

	if object.M_frequencyType == Config.FrequencyType_ONCE {
		xmlElement.M_frequencyType = "##ONCE"
	} else if object.M_frequencyType == Config.FrequencyType_DAILY {
		xmlElement.M_frequencyType = "##DAILY"
	} else if object.M_frequencyType == Config.FrequencyType_WEEKELY {
		xmlElement.M_frequencyType = "##WEEKELY"
	} else if object.M_frequencyType == Config.FrequencyType_MONTHLY {
		xmlElement.M_frequencyType = "##MONTHLY"
	}

	/** ScheduledTask **/
	xmlElement.M_offsets = object.M_offsets
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ServerConfiguration **/
func (this *ConfigXmlFactory) SerialyzeServerConfiguration(xmlElement *Config.XsdServerConfiguration, object *Config.ServerConfiguration) {
	if xmlElement == nil {
		return
	}

	/** Configuration **/
	xmlElement.M_id = object.M_id

	/** ServerConfiguration **/
	xmlElement.M_ipv4 = object.M_ipv4

	/** ServerConfiguration **/
	xmlElement.M_hostName = object.M_hostName

	/** ServerConfiguration **/
	xmlElement.M_serverPort = object.M_serverPort

	/** ServerConfiguration **/
	xmlElement.M_serviceContainerPort = object.M_serviceContainerPort

	/** ServerConfiguration **/
	xmlElement.M_applicationsPath = object.M_applicationsPath

	/** ServerConfiguration **/
	xmlElement.M_dataPath = object.M_dataPath

	/** ServerConfiguration **/
	xmlElement.M_scriptsPath = object.M_scriptsPath

	/** ServerConfiguration **/
	xmlElement.M_definitionsPath = object.M_definitionsPath

	/** ServerConfiguration **/
	xmlElement.M_schemasPath = object.M_schemasPath

	/** ServerConfiguration **/
	xmlElement.M_tmpPath = object.M_tmpPath

	/** ServerConfiguration **/
	xmlElement.M_binPath = object.M_binPath

	/** ServerConfiguration **/
	xmlElement.M_shards = object.M_shards

	/** ServerConfiguration **/
	xmlElement.M_lifeWindow = object.M_lifeWindow

	/** ServerConfiguration **/
	xmlElement.M_maxEntriesInWindow = object.M_maxEntriesInWindow

	/** ServerConfiguration **/
	xmlElement.M_maxEntrySize = object.M_maxEntrySize

	/** ServerConfiguration **/
	xmlElement.M_verbose = object.M_verbose

	/** ServerConfiguration **/
	xmlElement.M_hardMaxCacheSize = object.M_hardMaxCacheSize
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ApplicationConfiguration **/
func (this *ConfigXmlFactory) SerialyzeApplicationConfiguration(xmlElement *Config.XsdApplicationConfiguration, object *Config.ApplicationConfiguration) {
	if xmlElement == nil {
		return
	}

	/** Configuration **/
	xmlElement.M_id = object.M_id

	/** ApplicationConfiguration **/
	xmlElement.M_indexPage = object.M_indexPage
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of SmtpConfiguration **/
func (this *ConfigXmlFactory) SerialyzeSmtpConfiguration(xmlElement *Config.XsdSmtpConfiguration, object *Config.SmtpConfiguration) {
	if xmlElement == nil {
		return
	}

	/** Configuration **/
	xmlElement.M_id = object.M_id

	/** SmtpConfiguration **/
	xmlElement.M_hostName = object.M_hostName

	/** SmtpConfiguration **/
	xmlElement.M_ipv4 = object.M_ipv4

	/** SmtpConfiguration **/
	xmlElement.M_port = object.M_port

	/** SmtpConfiguration **/
	xmlElement.M_user = object.M_user

	/** SmtpConfiguration **/
	xmlElement.M_pwd = object.M_pwd

	if object.M_textEncoding == Config.Encoding_UTF8 {
		xmlElement.M_textEncoding = "##UTF8"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1250 {
		xmlElement.M_textEncoding = "##WINDOWS_1250"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1251 {
		xmlElement.M_textEncoding = "##WINDOWS_1251"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1252 {
		xmlElement.M_textEncoding = "##WINDOWS_1252"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1253 {
		xmlElement.M_textEncoding = "##WINDOWS_1253"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1254 {
		xmlElement.M_textEncoding = "##WINDOWS_1254"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1255 {
		xmlElement.M_textEncoding = "##WINDOWS_1255"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1256 {
		xmlElement.M_textEncoding = "##WINDOWS_1256"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1257 {
		xmlElement.M_textEncoding = "##WINDOWS_1257"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1258 {
		xmlElement.M_textEncoding = "##WINDOWS_1258"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_1 {
		xmlElement.M_textEncoding = "##ISO8859_1"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_2 {
		xmlElement.M_textEncoding = "##ISO8859_2"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_3 {
		xmlElement.M_textEncoding = "##ISO8859_3"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_4 {
		xmlElement.M_textEncoding = "##ISO8859_4"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_5 {
		xmlElement.M_textEncoding = "##ISO8859_5"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_6 {
		xmlElement.M_textEncoding = "##ISO8859_6"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_7 {
		xmlElement.M_textEncoding = "##ISO8859_7"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_8 {
		xmlElement.M_textEncoding = "##ISO8859_8"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_9 {
		xmlElement.M_textEncoding = "##ISO8859_9"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_10 {
		xmlElement.M_textEncoding = "##ISO8859_10"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_13 {
		xmlElement.M_textEncoding = "##ISO8859_13"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_14 {
		xmlElement.M_textEncoding = "##ISO8859_14"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_15 {
		xmlElement.M_textEncoding = "##ISO8859_15"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_16 {
		xmlElement.M_textEncoding = "##ISO8859_16"
	} else if object.M_textEncoding == Config.Encoding_KOI8R {
		xmlElement.M_textEncoding = "##KOI8R"
	} else if object.M_textEncoding == Config.Encoding_KOI8U {
		xmlElement.M_textEncoding = "##KOI8U"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of LdapConfiguration **/
func (this *ConfigXmlFactory) SerialyzeLdapConfiguration(xmlElement *Config.XsdLdapConfiguration, object *Config.LdapConfiguration) {
	if xmlElement == nil {
		return
	}

	/** Configuration **/
	xmlElement.M_id = object.M_id

	/** LdapConfiguration **/
	xmlElement.M_hostName = object.M_hostName

	/** LdapConfiguration **/
	xmlElement.M_ipv4 = object.M_ipv4

	/** LdapConfiguration **/
	xmlElement.M_port = object.M_port

	/** LdapConfiguration **/
	xmlElement.M_user = object.M_user

	/** LdapConfiguration **/
	xmlElement.M_pwd = object.M_pwd

	/** LdapConfiguration **/
	xmlElement.M_domain = object.M_domain

	/** LdapConfiguration **/
	xmlElement.M_searchBase = object.M_searchBase
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of DataStoreConfiguration **/
func (this *ConfigXmlFactory) SerialyzeDataStoreConfiguration(xmlElement *Config.XsdDataStoreConfiguration, object *Config.DataStoreConfiguration) {
	if xmlElement == nil {
		return
	}

	/** Configuration **/
	xmlElement.M_id = object.M_id

	/** DataStoreConfiguration **/
	xmlElement.M_storeName = object.M_storeName

	/** DataStoreConfiguration **/
	xmlElement.M_hostName = object.M_hostName

	/** DataStoreConfiguration **/
	xmlElement.M_ipv4 = object.M_ipv4

	/** DataStoreConfiguration **/
	xmlElement.M_port = object.M_port

	/** DataStoreConfiguration **/
	xmlElement.M_user = object.M_user

	/** DataStoreConfiguration **/
	xmlElement.M_pwd = object.M_pwd

	if object.M_dataStoreType == Config.DataStoreType_SQL_STORE {
		xmlElement.M_dataStoreType = "##SQL_STORE"
	} else if object.M_dataStoreType == Config.DataStoreType_GRAPH_STORE {
		xmlElement.M_dataStoreType = "##GRAPH_STORE"
	}

	if object.M_dataStoreVendor == Config.DataStoreVendor_CARGO {
		xmlElement.M_dataStoreVendor = "##CARGO"
	} else if object.M_dataStoreVendor == Config.DataStoreVendor_MYSQL {
		xmlElement.M_dataStoreVendor = "##MYSQL"
	} else if object.M_dataStoreVendor == Config.DataStoreVendor_MSSQL {
		xmlElement.M_dataStoreVendor = "##MSSQL"
	}

	if object.M_textEncoding == Config.Encoding_UTF8 {
		xmlElement.M_textEncoding = "##UTF8"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1250 {
		xmlElement.M_textEncoding = "##WINDOWS_1250"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1251 {
		xmlElement.M_textEncoding = "##WINDOWS_1251"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1252 {
		xmlElement.M_textEncoding = "##WINDOWS_1252"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1253 {
		xmlElement.M_textEncoding = "##WINDOWS_1253"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1254 {
		xmlElement.M_textEncoding = "##WINDOWS_1254"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1255 {
		xmlElement.M_textEncoding = "##WINDOWS_1255"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1256 {
		xmlElement.M_textEncoding = "##WINDOWS_1256"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1257 {
		xmlElement.M_textEncoding = "##WINDOWS_1257"
	} else if object.M_textEncoding == Config.Encoding_WINDOWS_1258 {
		xmlElement.M_textEncoding = "##WINDOWS_1258"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_1 {
		xmlElement.M_textEncoding = "##ISO8859_1"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_2 {
		xmlElement.M_textEncoding = "##ISO8859_2"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_3 {
		xmlElement.M_textEncoding = "##ISO8859_3"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_4 {
		xmlElement.M_textEncoding = "##ISO8859_4"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_5 {
		xmlElement.M_textEncoding = "##ISO8859_5"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_6 {
		xmlElement.M_textEncoding = "##ISO8859_6"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_7 {
		xmlElement.M_textEncoding = "##ISO8859_7"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_8 {
		xmlElement.M_textEncoding = "##ISO8859_8"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_9 {
		xmlElement.M_textEncoding = "##ISO8859_9"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_10 {
		xmlElement.M_textEncoding = "##ISO8859_10"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_13 {
		xmlElement.M_textEncoding = "##ISO8859_13"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_14 {
		xmlElement.M_textEncoding = "##ISO8859_14"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_15 {
		xmlElement.M_textEncoding = "##ISO8859_15"
	} else if object.M_textEncoding == Config.Encoding_ISO8859_16 {
		xmlElement.M_textEncoding = "##ISO8859_16"
	} else if object.M_textEncoding == Config.Encoding_KOI8R {
		xmlElement.M_textEncoding = "##KOI8R"
	} else if object.M_textEncoding == Config.Encoding_KOI8U {
		xmlElement.M_textEncoding = "##KOI8U"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Configurations **/
func (this *ConfigXmlFactory) SerialyzeConfigurations(xmlElement *Config.XsdConfigurations, object *Config.Configurations) {
	if xmlElement == nil {
		return
	}

	/** Configurations **/
	xmlElement.M_id = object.M_id

	/** Configurations **/
	xmlElement.M_name = object.M_name

	/** Configurations **/
	xmlElement.M_version = object.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ServerConfiguration **/

	/** Now I will save the value of serverConfig **/
	if object.GetServerConfig() != nil {
		this.SerialyzeServerConfiguration(&xmlElement.M_serverConfig, object.GetServerConfig())
	}

	/** Serialyze ApplicationConfiguration **/
	if len(object.M_applicationConfigs) > 0 {
		xmlElement.M_applicationConfigs = make([]*Config.XsdApplicationConfiguration, 0)
	}

	/** Now I will save the value of applicationConfigs **/
	for i := 0; i < len(object.M_applicationConfigs); i++ {
		xmlElement.M_applicationConfigs = append(xmlElement.M_applicationConfigs, new(Config.XsdApplicationConfiguration))
		this.SerialyzeApplicationConfiguration(xmlElement.M_applicationConfigs[i], object.GetApplicationConfigs()[i])
	}

	/** Serialyze SmtpConfiguration **/
	if len(object.M_smtpConfigs) > 0 {
		xmlElement.M_smtpConfigs = make([]*Config.XsdSmtpConfiguration, 0)
	}

	/** Now I will save the value of smtpConfigs **/
	for i := 0; i < len(object.M_smtpConfigs); i++ {
		xmlElement.M_smtpConfigs = append(xmlElement.M_smtpConfigs, new(Config.XsdSmtpConfiguration))
		this.SerialyzeSmtpConfiguration(xmlElement.M_smtpConfigs[i], object.GetSmtpConfigs()[i])
	}

	/** Serialyze LdapConfiguration **/
	if len(object.M_ldapConfigs) > 0 {
		xmlElement.M_ldapConfigs = make([]*Config.XsdLdapConfiguration, 0)
	}

	/** Now I will save the value of ldapConfigs **/
	for i := 0; i < len(object.M_ldapConfigs); i++ {
		xmlElement.M_ldapConfigs = append(xmlElement.M_ldapConfigs, new(Config.XsdLdapConfiguration))
		this.SerialyzeLdapConfiguration(xmlElement.M_ldapConfigs[i], object.GetLdapConfigs()[i])
	}

	/** Serialyze DataStoreConfiguration **/
	if len(object.M_dataStoreConfigs) > 0 {
		xmlElement.M_dataStoreConfigs = make([]*Config.XsdDataStoreConfiguration, 0)
	}

	/** Now I will save the value of dataStoreConfigs **/
	for i := 0; i < len(object.M_dataStoreConfigs); i++ {
		xmlElement.M_dataStoreConfigs = append(xmlElement.M_dataStoreConfigs, new(Config.XsdDataStoreConfiguration))
		this.SerialyzeDataStoreConfiguration(xmlElement.M_dataStoreConfigs[i], object.GetDataStoreConfigs()[i])
	}

	/** Serialyze ServiceConfiguration **/
	if len(object.M_serviceConfigs) > 0 {
		xmlElement.M_serviceConfigs = make([]*Config.XsdServiceConfiguration, 0)
	}

	/** Now I will save the value of serviceConfigs **/
	for i := 0; i < len(object.M_serviceConfigs); i++ {
		xmlElement.M_serviceConfigs = append(xmlElement.M_serviceConfigs, new(Config.XsdServiceConfiguration))
		this.SerialyzeServiceConfiguration(xmlElement.M_serviceConfigs[i], object.GetServiceConfigs()[i])
	}

	/** Serialyze OAuth2Configuration **/
	if object.GetOauth2Configuration() != nil {
		xmlElement.M_oauth2Configuration = new(Config.XsdOAuth2Configuration)
	}

	/** Now I will save the value of oauth2Configuration **/
	if object.GetOauth2Configuration() != nil {
		this.SerialyzeOAuth2Configuration(xmlElement.M_oauth2Configuration, object.GetOauth2Configuration())
	}

	/** Serialyze ScheduledTask **/
	if len(object.M_scheduledTasks) > 0 {
		xmlElement.M_scheduledTasks = make([]*Config.XsdScheduledTask, 0)
	}

	/** Now I will save the value of scheduledTasks **/
	for i := 0; i < len(object.M_scheduledTasks); i++ {
		xmlElement.M_scheduledTasks = append(xmlElement.M_scheduledTasks, new(Config.XsdScheduledTask))
		this.SerialyzeScheduledTask(xmlElement.M_scheduledTasks[i], object.GetScheduledTasks()[i])
	}
}

/** serialysation of ServiceConfiguration **/
func (this *ConfigXmlFactory) SerialyzeServiceConfiguration(xmlElement *Config.XsdServiceConfiguration, object *Config.ServiceConfiguration) {
	if xmlElement == nil {
		return
	}

	/** Configuration **/
	xmlElement.M_id = object.M_id

	/** ServiceConfiguration **/
	xmlElement.M_hostName = object.M_hostName

	/** ServiceConfiguration **/
	xmlElement.M_ipv4 = object.M_ipv4

	/** ServiceConfiguration **/
	xmlElement.M_port = object.M_port

	/** ServiceConfiguration **/
	xmlElement.M_user = object.M_user

	/** ServiceConfiguration **/
	xmlElement.M_pwd = object.M_pwd

	/** ServiceConfiguration **/
	xmlElement.M_start = object.M_start
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}
