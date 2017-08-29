// +build Config

package Entities

import (
	"encoding/xml"
	"log"
	"os"
	"path/filepath"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
	"golang.org/x/net/html/charset"
)

type ConfigXmlFactory struct {
	m_references map[string]interface{}
	m_object     map[string]map[string][]string
}

/** Initialization function from xml file **/
func (this *ConfigXmlFactory) InitXml(inputPath string, object *Config.Configurations) error {
	this.m_references = make(map[string]interface{})
	this.m_object = make(map[string]map[string][]string)
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
						methodName := "Set" + strings.ToUpper(ref1[0:1]) + ref1[1:]
						Utility.CallMethod(refOwner, methodName, params)
					} else {
						params := make([]interface{}, 0)
						params = append(params, refs[i])
						methodName := "Set" + strings.ToUpper(ref1[0:1]) + ref1[1:]
						Utility.CallMethod(refOwner, methodName, params)
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

/** inititialisation of OAuth2Configuration **/
func (this *ConfigXmlFactory) InitOAuth2Configuration(parentUuid string, xmlElement *Config.XsdOAuth2Configuration, object *Config.OAuth2Configuration) {
	log.Println("Initialize OAuth2Configuration")

	/** OAuth2Configuration **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
	}
	var uuidStr string
	if len(parentUuid) > 0 {
		uuidStr += parentUuid + ":"
	}
	uuidStr += "Config.OAuth2Configuration:" + object.M_id
	object.UUID = "Config.OAuth2Configuration%" + Utility.GenerateUUID(uuidStr)

	/** Configuration **/
	object.M_authorizationExpiration = xmlElement.M_authorizationExpiration

	/** Configuration **/
	object.M_accessExpiration = xmlElement.M_accessExpiration

	/** Configuration **/
	object.M_tokenType = xmlElement.M_tokenType

	/** Configuration **/
	object.M_errorStatusCode = xmlElement.M_errorStatusCode

	/** Configuration **/
	object.M_allowClientSecretInParams = xmlElement.M_allowClientSecretInParams

	/** Configuration **/
	object.M_allowGetAccessRequest = xmlElement.M_allowGetAccessRequest

	/** Configuration **/
	object.M_redirectUriSeparator = xmlElement.M_redirectUriSeparator

	/** Configuration **/
	object.M_privateKey = xmlElement.M_privateKey
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ServerConfiguration **/
func (this *ConfigXmlFactory) InitServerConfiguration(parentUuid string, xmlElement *Config.XsdServerConfiguration, object *Config.ServerConfiguration) {
	log.Println("Initialize ServerConfiguration")

	/** ServerConfiguration **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
	}
	var uuidStr string
	if len(parentUuid) > 0 {
		uuidStr += parentUuid + ":"
	}
	uuidStr += "Config.ServerConfiguration:" + object.M_id
	object.UUID = "Config.ServerConfiguration%" + Utility.GenerateUUID(uuidStr)

	/** Configuration **/
	object.M_ipv4 = xmlElement.M_ipv4

	/** Configuration **/
	object.M_hostName = xmlElement.M_hostName

	/** Configuration **/
	object.M_serverPort = xmlElement.M_serverPort

	/** Configuration **/
	object.M_ws_serviceContainerPort = xmlElement.M_ws_serviceContainerPort

	/** Configuration **/
	object.M_tcp_serviceContainerPort = xmlElement.M_tcp_serviceContainerPort

	/** Configuration **/
	object.M_applicationsPath = xmlElement.M_applicationsPath

	/** Configuration **/
	object.M_dataPath = xmlElement.M_dataPath

	/** Configuration **/
	object.M_scriptsPath = xmlElement.M_scriptsPath

	/** Configuration **/
	object.M_definitionsPath = xmlElement.M_definitionsPath

	/** Configuration **/
	object.M_schemasPath = xmlElement.M_schemasPath

	/** Configuration **/
	object.M_tmpPath = xmlElement.M_tmpPath

	/** Configuration **/
	object.M_binPath = xmlElement.M_binPath

	/** Configuration **/
	object.M_queriesPath = xmlElement.M_queriesPath
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ApplicationConfiguration **/
func (this *ConfigXmlFactory) InitApplicationConfiguration(parentUuid string, xmlElement *Config.XsdApplicationConfiguration, object *Config.ApplicationConfiguration) {
	log.Println("Initialize ApplicationConfiguration")

	/** ApplicationConfiguration **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
	}
	var uuidStr string
	if len(parentUuid) > 0 {
		uuidStr += parentUuid + ":"
	}
	uuidStr += "Config.ApplicationConfiguration:" + object.M_id
	object.UUID = "Config.ApplicationConfiguration%" + Utility.GenerateUUID(uuidStr)

	/** Configuration **/
	object.M_indexPage = xmlElement.M_indexPage
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of DataStoreConfiguration **/
func (this *ConfigXmlFactory) InitDataStoreConfiguration(parentUuid string, xmlElement *Config.XsdDataStoreConfiguration, object *Config.DataStoreConfiguration) {
	log.Println("Initialize DataStoreConfiguration")

	/** DataStoreConfiguration **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
	}
	var uuidStr string
	if len(parentUuid) > 0 {
		uuidStr += parentUuid + ":"
	}
	uuidStr += "Config.DataStoreConfiguration:" + object.M_id
	object.UUID = "Config.DataStoreConfiguration%" + Utility.GenerateUUID(uuidStr)

	/** Configuration **/
	object.M_hostName = xmlElement.M_hostName

	/** Configuration **/
	object.M_ipv4 = xmlElement.M_ipv4

	/** Configuration **/
	object.M_port = xmlElement.M_port

	/** Configuration **/
	object.M_user = xmlElement.M_user

	/** Configuration **/
	object.M_pwd = xmlElement.M_pwd

	/** DataStoreType **/
	if xmlElement.M_dataStoreType == "##SQL_STORE" {
		object.M_dataStoreType = Config.DataStoreType_SQL_STORE
	} else if xmlElement.M_dataStoreType == "##KEY_VALUE_STORE" {
		object.M_dataStoreType = Config.DataStoreType_KEY_VALUE_STORE
	}

	/** DataStoreVendor **/
	if xmlElement.M_dataStoreVendor == "##MYCELIUS" {
		object.M_dataStoreVendor = Config.DataStoreVendor_MYCELIUS
	} else if xmlElement.M_dataStoreVendor == "##MYSQL" {
		object.M_dataStoreVendor = Config.DataStoreVendor_MYSQL
	} else if xmlElement.M_dataStoreVendor == "##MSSQL" {
		object.M_dataStoreVendor = Config.DataStoreVendor_MSSQL
	}

	/** Encoding **/
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

/** inititialisation of ServiceConfiguration **/
func (this *ConfigXmlFactory) InitServiceConfiguration(parentUuid string, xmlElement *Config.XsdServiceConfiguration, object *Config.ServiceConfiguration) {
	log.Println("Initialize ServiceConfiguration")

	/** ServiceConfiguration **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
	}
	var uuidStr string
	if len(parentUuid) > 0 {
		uuidStr += parentUuid + ":"
	}
	uuidStr += "Config.ServiceConfiguration:" + object.M_id
	object.UUID = "Config.ServiceConfiguration%" + Utility.GenerateUUID(uuidStr)

	/** Configuration **/
	object.M_hostName = xmlElement.M_hostName

	/** Configuration **/
	object.M_ipv4 = xmlElement.M_ipv4

	/** Configuration **/
	object.M_port = xmlElement.M_port

	/** Configuration **/
	object.M_user = xmlElement.M_user

	/** Configuration **/
	object.M_pwd = xmlElement.M_pwd

	/** Configuration **/
	object.M_start = xmlElement.M_start
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Configurations **/
func (this *ConfigXmlFactory) InitConfigurations(parentUuid string, xmlElement *Config.XsdConfigurations, object *Config.Configurations) {
	log.Println("Initialize Configurations")

	/** Init serverConfiguration **/
	if object.M_serverConfig == nil {
		object.M_serverConfig = new(Config.ServerConfiguration)
	}
	this.InitServerConfiguration(object.UUID, &xmlElement.M_serverConfig, object.M_serverConfig)

	/** association initialisation **/

	/** Init applicationConfiguration **/
	object.M_applicationConfigs = make([]*Config.ApplicationConfiguration, 0)
	for i := 0; i < len(xmlElement.M_applicationConfigs); i++ {
		val := new(Config.ApplicationConfiguration)
		this.InitApplicationConfiguration(object.UUID, xmlElement.M_applicationConfigs[i], val)
		object.M_applicationConfigs = append(object.M_applicationConfigs, val)

		/** association initialisation **/
	}

	/** Init smtpConfiguration **/
	object.M_smtpConfigs = make([]*Config.SmtpConfiguration, 0)
	for i := 0; i < len(xmlElement.M_smtpConfigs); i++ {
		val := new(Config.SmtpConfiguration)
		this.InitSmtpConfiguration(object.UUID, xmlElement.M_smtpConfigs[i], val)
		object.M_smtpConfigs = append(object.M_smtpConfigs, val)

		/** association initialisation **/
	}

	/** Init ldapConfiguration **/
	object.M_ldapConfigs = make([]*Config.LdapConfiguration, 0)
	for i := 0; i < len(xmlElement.M_ldapConfigs); i++ {
		val := new(Config.LdapConfiguration)
		this.InitLdapConfiguration(object.UUID, xmlElement.M_ldapConfigs[i], val)
		object.M_ldapConfigs = append(object.M_ldapConfigs, val)

		/** association initialisation **/
	}

	/** Init dataStoreConfiguration **/
	object.M_dataStoreConfigs = make([]*Config.DataStoreConfiguration, 0)
	for i := 0; i < len(xmlElement.M_dataStoreConfigs); i++ {
		val := new(Config.DataStoreConfiguration)
		this.InitDataStoreConfiguration(object.UUID, xmlElement.M_dataStoreConfigs[i], val)
		object.M_dataStoreConfigs = append(object.M_dataStoreConfigs, val)

		/** association initialisation **/
	}

	/** Init serviceConfiguration **/
	object.M_serviceConfigs = make([]*Config.ServiceConfiguration, 0)
	for i := 0; i < len(xmlElement.M_serviceConfigs); i++ {
		val := new(Config.ServiceConfiguration)
		this.InitServiceConfiguration(object.UUID, xmlElement.M_serviceConfigs[i], val)
		object.M_serviceConfigs = append(object.M_serviceConfigs, val)

		/** association initialisation **/
	}

	/** Init oauth2Configuration **/
	if xmlElement.M_oauth2Configuration != nil {
		object.M_oauth2Configuration = new(Config.OAuth2Configuration)
		this.InitOAuth2Configuration(object.UUID, xmlElement.M_oauth2Configuration, object.M_oauth2Configuration)

		/** association initialisation **/
	}

	/** Init scheduledTask **/
	object.M_scheduledTasks = make([]*Config.ScheduledTask, 0)
	for i := 0; i < len(xmlElement.M_scheduledTasks); i++ {
		val := new(Config.ScheduledTask)
		this.InitScheduledTask(object.UUID, xmlElement.M_scheduledTasks[i], val)
		object.M_scheduledTasks = append(object.M_scheduledTasks, val)

		/** association initialisation **/
	}

	/** Configurations **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
	}
	var uuidStr string
	if len(parentUuid) > 0 {
		uuidStr += parentUuid + ":"
	}
	uuidStr += "Config.Configurations:" + object.M_id
	object.UUID = "Config.Configurations%" + Utility.GenerateUUID(uuidStr)

	/** Configurations **/
	object.M_name = xmlElement.M_name

	/** Configurations **/
	object.M_version = xmlElement.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of SmtpConfiguration **/
func (this *ConfigXmlFactory) InitSmtpConfiguration(parentUuid string, xmlElement *Config.XsdSmtpConfiguration, object *Config.SmtpConfiguration) {
	log.Println("Initialize SmtpConfiguration")

	/** SmtpConfiguration **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
	}
	var uuidStr string
	if len(parentUuid) > 0 {
		uuidStr += parentUuid + ":"
	}
	uuidStr += "Config.SmtpConfiguration:" + object.M_id
	object.UUID = "Config.SmtpConfiguration%" + Utility.GenerateUUID(uuidStr)

	/** Configuration **/
	object.M_hostName = xmlElement.M_hostName

	/** Configuration **/
	object.M_ipv4 = xmlElement.M_ipv4

	/** Configuration **/
	object.M_port = xmlElement.M_port

	/** Configuration **/
	object.M_user = xmlElement.M_user

	/** Configuration **/
	object.M_pwd = xmlElement.M_pwd

	/** Encoding **/
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

/** inititialisation of LdapConfiguration **/
func (this *ConfigXmlFactory) InitLdapConfiguration(parentUuid string, xmlElement *Config.XsdLdapConfiguration, object *Config.LdapConfiguration) {
	log.Println("Initialize LdapConfiguration")

	/** LdapConfiguration **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
	}
	var uuidStr string
	if len(parentUuid) > 0 {
		uuidStr += parentUuid + ":"
	}
	uuidStr += "Config.LdapConfiguration:" + object.M_id
	object.UUID = "Config.LdapConfiguration%" + Utility.GenerateUUID(uuidStr)

	/** Configuration **/
	object.M_hostName = xmlElement.M_hostName

	/** Configuration **/
	object.M_ipv4 = xmlElement.M_ipv4

	/** Configuration **/
	object.M_port = xmlElement.M_port

	/** Configuration **/
	object.M_user = xmlElement.M_user

	/** Configuration **/
	object.M_pwd = xmlElement.M_pwd

	/** Configuration **/
	object.M_domain = xmlElement.M_domain

	/** Configuration **/
	object.M_searchBase = xmlElement.M_searchBase
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ScheduledTask **/
func (this *ConfigXmlFactory) InitScheduledTask(parentUuid string, xmlElement *Config.XsdScheduledTask, object *Config.ScheduledTask) {
	log.Println("Initialize ScheduledTask")

	/** ScheduledTask **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
	}
	var uuidStr string
	if len(parentUuid) > 0 {
		uuidStr += parentUuid + ":"
	}
	uuidStr += "Config.ScheduledTask:" + object.M_id
	object.UUID = "Config.ScheduledTask%" + Utility.GenerateUUID(uuidStr)

	/** Configuration **/
	object.M_isActive = xmlElement.M_isActive

	/** Configuration **/
	object.M_script = xmlElement.M_script

	/** Configuration **/
	object.M_startTime = xmlElement.M_startTime

	/** Configuration **/
	object.M_expirationTime = xmlElement.M_expirationTime

	/** Configuration **/
	object.M_frequency = xmlElement.M_frequency

	/** FrequencyType **/
	if xmlElement.M_frequencyType == "##ONCE" {
		object.M_frequencyType = Config.FrequencyType_ONCE
	} else if xmlElement.M_frequencyType == "##DAILY" {
		object.M_frequencyType = Config.FrequencyType_DAILY
	} else if xmlElement.M_frequencyType == "##WEEKELY" {
		object.M_frequencyType = Config.FrequencyType_WEEKELY
	} else if xmlElement.M_frequencyType == "##MONTHLY" {
		object.M_frequencyType = Config.FrequencyType_MONTHLY
	}

	/** Configuration **/
	object.M_offsets = xmlElement.M_offsets

	/** Configuration **/
	object.M_iterations = xmlElement.M_iterations

	/** Configuration **/
	object.M_keepAlive = xmlElement.M_keepAlive
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Configurations **/
func (this *ConfigXmlFactory) SerialyzeConfigurations(xmlElement *Config.XsdConfigurations, object *Config.Configurations) {
	if xmlElement == nil {
		return
	}

	/** Serialyze ServerConfiguration **/

	/** Now I will save the value of serverConfig **/
	if object.M_serverConfig != nil {
		this.SerialyzeServerConfiguration(&xmlElement.M_serverConfig, object.M_serverConfig)
	}

	/** Serialyze ApplicationConfiguration **/
	if len(object.M_applicationConfigs) > 0 {
		xmlElement.M_applicationConfigs = make([]*Config.XsdApplicationConfiguration, 0)
	}

	/** Now I will save the value of applicationConfigs **/
	for i := 0; i < len(object.M_applicationConfigs); i++ {
		xmlElement.M_applicationConfigs = append(xmlElement.M_applicationConfigs, new(Config.XsdApplicationConfiguration))
		this.SerialyzeApplicationConfiguration(xmlElement.M_applicationConfigs[i], object.M_applicationConfigs[i])
	}

	/** Serialyze SmtpConfiguration **/
	if len(object.M_smtpConfigs) > 0 {
		xmlElement.M_smtpConfigs = make([]*Config.XsdSmtpConfiguration, 0)
	}

	/** Now I will save the value of smtpConfigs **/
	for i := 0; i < len(object.M_smtpConfigs); i++ {
		xmlElement.M_smtpConfigs = append(xmlElement.M_smtpConfigs, new(Config.XsdSmtpConfiguration))
		this.SerialyzeSmtpConfiguration(xmlElement.M_smtpConfigs[i], object.M_smtpConfigs[i])
	}

	/** Serialyze LdapConfiguration **/
	if len(object.M_ldapConfigs) > 0 {
		xmlElement.M_ldapConfigs = make([]*Config.XsdLdapConfiguration, 0)
	}

	/** Now I will save the value of ldapConfigs **/
	for i := 0; i < len(object.M_ldapConfigs); i++ {
		xmlElement.M_ldapConfigs = append(xmlElement.M_ldapConfigs, new(Config.XsdLdapConfiguration))
		this.SerialyzeLdapConfiguration(xmlElement.M_ldapConfigs[i], object.M_ldapConfigs[i])
	}

	/** Serialyze DataStoreConfiguration **/
	if len(object.M_dataStoreConfigs) > 0 {
		xmlElement.M_dataStoreConfigs = make([]*Config.XsdDataStoreConfiguration, 0)
	}

	/** Now I will save the value of dataStoreConfigs **/
	for i := 0; i < len(object.M_dataStoreConfigs); i++ {
		xmlElement.M_dataStoreConfigs = append(xmlElement.M_dataStoreConfigs, new(Config.XsdDataStoreConfiguration))
		this.SerialyzeDataStoreConfiguration(xmlElement.M_dataStoreConfigs[i], object.M_dataStoreConfigs[i])
	}

	/** Serialyze ServiceConfiguration **/
	if len(object.M_serviceConfigs) > 0 {
		xmlElement.M_serviceConfigs = make([]*Config.XsdServiceConfiguration, 0)
	}

	/** Now I will save the value of serviceConfigs **/
	for i := 0; i < len(object.M_serviceConfigs); i++ {
		xmlElement.M_serviceConfigs = append(xmlElement.M_serviceConfigs, new(Config.XsdServiceConfiguration))
		this.SerialyzeServiceConfiguration(xmlElement.M_serviceConfigs[i], object.M_serviceConfigs[i])
	}

	/** Serialyze OAuth2Configuration **/
	if object.M_oauth2Configuration != nil {
		xmlElement.M_oauth2Configuration = new(Config.XsdOAuth2Configuration)
	}

	/** Now I will save the value of oauth2Configuration **/
	if object.M_oauth2Configuration != nil {
		this.SerialyzeOAuth2Configuration(xmlElement.M_oauth2Configuration, object.M_oauth2Configuration)
	}

	/** Serialyze ScheduledTask **/
	if len(object.M_scheduledTasks) > 0 {
		xmlElement.M_scheduledTasks = make([]*Config.XsdScheduledTask, 0)
	}

	/** Now I will save the value of scheduledTasks **/
	for i := 0; i < len(object.M_scheduledTasks); i++ {
		xmlElement.M_scheduledTasks = append(xmlElement.M_scheduledTasks, new(Config.XsdScheduledTask))
		this.SerialyzeScheduledTask(xmlElement.M_scheduledTasks[i], object.M_scheduledTasks[i])
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
}

/** serialysation of ServerConfiguration **/
func (this *ConfigXmlFactory) SerialyzeServerConfiguration(xmlElement *Config.XsdServerConfiguration, object *Config.ServerConfiguration) {
	if xmlElement == nil {
		return
	}

	/** ServerConfiguration **/
	xmlElement.M_id = object.M_id

	/** Configuration **/
	xmlElement.M_ipv4 = object.M_ipv4

	/** Configuration **/
	xmlElement.M_hostName = object.M_hostName

	/** Configuration **/
	xmlElement.M_serverPort = object.M_serverPort

	/** Configuration **/
	xmlElement.M_ws_serviceContainerPort = object.M_ws_serviceContainerPort

	/** Configuration **/
	xmlElement.M_tcp_serviceContainerPort = object.M_tcp_serviceContainerPort

	/** Configuration **/
	xmlElement.M_applicationsPath = object.M_applicationsPath

	/** Configuration **/
	xmlElement.M_dataPath = object.M_dataPath

	/** Configuration **/
	xmlElement.M_scriptsPath = object.M_scriptsPath

	/** Configuration **/
	xmlElement.M_definitionsPath = object.M_definitionsPath

	/** Configuration **/
	xmlElement.M_schemasPath = object.M_schemasPath

	/** Configuration **/
	xmlElement.M_tmpPath = object.M_tmpPath

	/** Configuration **/
	xmlElement.M_binPath = object.M_binPath

	/** Configuration **/
	xmlElement.M_queriesPath = object.M_queriesPath
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of LdapConfiguration **/
func (this *ConfigXmlFactory) SerialyzeLdapConfiguration(xmlElement *Config.XsdLdapConfiguration, object *Config.LdapConfiguration) {
	if xmlElement == nil {
		return
	}

	/** LdapConfiguration **/
	xmlElement.M_id = object.M_id

	/** Configuration **/
	xmlElement.M_hostName = object.M_hostName

	/** Configuration **/
	xmlElement.M_ipv4 = object.M_ipv4

	/** Configuration **/
	xmlElement.M_port = object.M_port

	/** Configuration **/
	xmlElement.M_user = object.M_user

	/** Configuration **/
	xmlElement.M_pwd = object.M_pwd

	/** Configuration **/
	xmlElement.M_domain = object.M_domain

	/** Configuration **/
	xmlElement.M_searchBase = object.M_searchBase
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of OAuth2Configuration **/
func (this *ConfigXmlFactory) SerialyzeOAuth2Configuration(xmlElement *Config.XsdOAuth2Configuration, object *Config.OAuth2Configuration) {
	if xmlElement == nil {
		return
	}

	/** OAuth2Configuration **/
	xmlElement.M_id = object.M_id

	/** Configuration **/
	xmlElement.M_authorizationExpiration = object.M_authorizationExpiration

	/** Configuration **/
	xmlElement.M_accessExpiration = object.M_accessExpiration

	/** Configuration **/
	xmlElement.M_tokenType = object.M_tokenType

	/** Configuration **/
	xmlElement.M_errorStatusCode = object.M_errorStatusCode

	/** Configuration **/
	xmlElement.M_allowClientSecretInParams = object.M_allowClientSecretInParams

	/** Configuration **/
	xmlElement.M_allowGetAccessRequest = object.M_allowGetAccessRequest

	/** Configuration **/
	xmlElement.M_redirectUriSeparator = object.M_redirectUriSeparator

	/** Configuration **/
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

	/** ScheduledTask **/
	xmlElement.M_id = object.M_id

	/** Configuration **/
	xmlElement.M_isActive = object.M_isActive

	/** Configuration **/
	xmlElement.M_script = object.M_script

	/** Configuration **/
	xmlElement.M_startTime = object.M_startTime

	/** Configuration **/
	xmlElement.M_expirationTime = object.M_expirationTime

	/** Configuration **/
	xmlElement.M_frequency = object.M_frequency

	/** FrequencyType **/
	if object.M_frequencyType == Config.FrequencyType_ONCE {
		xmlElement.M_frequencyType = "##ONCE"
	} else if object.M_frequencyType == Config.FrequencyType_DAILY {
		xmlElement.M_frequencyType = "##DAILY"
	} else if object.M_frequencyType == Config.FrequencyType_WEEKELY {
		xmlElement.M_frequencyType = "##WEEKELY"
	} else if object.M_frequencyType == Config.FrequencyType_MONTHLY {
		xmlElement.M_frequencyType = "##MONTHLY"
	}

	/** Configuration **/
	xmlElement.M_offsets = object.M_offsets

	/** Configuration **/
	xmlElement.M_iterations = object.M_iterations

	/** Configuration **/
	xmlElement.M_keepAlive = object.M_keepAlive
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ApplicationConfiguration **/
func (this *ConfigXmlFactory) SerialyzeApplicationConfiguration(xmlElement *Config.XsdApplicationConfiguration, object *Config.ApplicationConfiguration) {
	if xmlElement == nil {
		return
	}

	/** ApplicationConfiguration **/
	xmlElement.M_id = object.M_id

	/** Configuration **/
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

	/** SmtpConfiguration **/
	xmlElement.M_id = object.M_id

	/** Configuration **/
	xmlElement.M_hostName = object.M_hostName

	/** Configuration **/
	xmlElement.M_ipv4 = object.M_ipv4

	/** Configuration **/
	xmlElement.M_port = object.M_port

	/** Configuration **/
	xmlElement.M_user = object.M_user

	/** Configuration **/
	xmlElement.M_pwd = object.M_pwd

	/** Encoding **/
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

/** serialysation of DataStoreConfiguration **/
func (this *ConfigXmlFactory) SerialyzeDataStoreConfiguration(xmlElement *Config.XsdDataStoreConfiguration, object *Config.DataStoreConfiguration) {
	if xmlElement == nil {
		return
	}

	/** DataStoreConfiguration **/
	xmlElement.M_id = object.M_id

	/** Configuration **/
	xmlElement.M_hostName = object.M_hostName

	/** Configuration **/
	xmlElement.M_ipv4 = object.M_ipv4

	/** Configuration **/
	xmlElement.M_port = object.M_port

	/** Configuration **/
	xmlElement.M_user = object.M_user

	/** Configuration **/
	xmlElement.M_pwd = object.M_pwd

	/** DataStoreType **/
	if object.M_dataStoreType == Config.DataStoreType_SQL_STORE {
		xmlElement.M_dataStoreType = "##SQL_STORE"
	} else if object.M_dataStoreType == Config.DataStoreType_KEY_VALUE_STORE {
		xmlElement.M_dataStoreType = "##KEY_VALUE_STORE"
	}

	/** DataStoreVendor **/
	if object.M_dataStoreVendor == Config.DataStoreVendor_MYCELIUS {
		xmlElement.M_dataStoreVendor = "##MYCELIUS"
	} else if object.M_dataStoreVendor == Config.DataStoreVendor_MYSQL {
		xmlElement.M_dataStoreVendor = "##MYSQL"
	} else if object.M_dataStoreVendor == Config.DataStoreVendor_MSSQL {
		xmlElement.M_dataStoreVendor = "##MSSQL"
	}

	/** Encoding **/
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

/** serialysation of ServiceConfiguration **/
func (this *ConfigXmlFactory) SerialyzeServiceConfiguration(xmlElement *Config.XsdServiceConfiguration, object *Config.ServiceConfiguration) {
	if xmlElement == nil {
		return
	}

	/** ServiceConfiguration **/
	xmlElement.M_id = object.M_id

	/** Configuration **/
	xmlElement.M_hostName = object.M_hostName

	/** Configuration **/
	xmlElement.M_ipv4 = object.M_ipv4

	/** Configuration **/
	xmlElement.M_port = object.M_port

	/** Configuration **/
	xmlElement.M_user = object.M_user

	/** Configuration **/
	xmlElement.M_pwd = object.M_pwd

	/** Configuration **/
	xmlElement.M_start = object.M_start
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}
