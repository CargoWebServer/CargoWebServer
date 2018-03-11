// +build CargoEntities

package Entities

import (
	"encoding/xml"
	"log"
	"os"
	"path/filepath"
	//	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
	"golang.org/x/net/html/charset"
)

type CargoEntitiesXmlFactory struct {
	m_references      map[string]interface{}
	m_object          map[string]map[string][]string
	m_getEntityByUuid func(string) (interface{}, error)
	m_setEntity       func(interface{})
	m_generateUuid    func(interface{}) string
}

/** Get entity by uuid function **/
func (this *CargoEntitiesXmlFactory) SetEntityGetter(fct func(string) (interface{}, error)) {
	this.m_getEntityByUuid = fct
}

/** Use to put the entity in the cache **/
func (this *CargoEntitiesXmlFactory) SetEntitySetter(fct func(interface{})) {
	this.m_setEntity = fct
}

/** Set the uuid genarator function. **/
func (this *CargoEntitiesXmlFactory) SetUuidGenerator(fct func(entity interface{}) string) {
	this.m_generateUuid = fct
}

/** Initialization function from xml file **/
func (this *CargoEntitiesXmlFactory) InitXml(inputPath string, object *CargoEntities.Entities) error {
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
	var xmlElement *CargoEntities.XsdEntities
	xmlElement = new(CargoEntities.XsdEntities)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(xmlElement); err != nil {
		return err
	}
	this.m_references = make(map[string]interface{}, 0)
	this.m_object = make(map[string]map[string][]string, 0)
	this.InitEntities("", xmlElement, object)
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
func (this *CargoEntitiesXmlFactory) SerializeXml(outputPath string, toSerialize *CargoEntities.Entities) error {
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

	var xmlElement *CargoEntities.XsdEntities
	xmlElement = new(CargoEntities.XsdEntities)

	this.SerialyzeEntities(xmlElement, toSerialize)
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

/** inititialisation of Account **/
func (this *CargoEntitiesXmlFactory) InitAccount(parentUuid string, xmlElement *CargoEntities.XsdAccount, object *CargoEntities.Account) {
	log.Println("Initialize Account")
	object.TYPENAME = "CargoEntities.Account"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Entity **/
	object.M_id = xmlElement.M_id

	/** Account **/
	object.M_name = xmlElement.M_name

	/** Account **/
	object.M_password = xmlElement.M_password

	/** Account **/
	object.M_email = xmlElement.M_email

	/** Init ref userRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_userRef != nil {
		if _, ok := this.m_object[object.M_id]["SetUserRef"]; !ok {
			this.m_object[object.M_id]["SetUserRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["SetUserRef"] = append(this.m_object[object.M_id]["SetUserRef"], *xmlElement.M_userRef)
	}

	/** Init ref rolesRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_rolesRef); i++ {
		if _, ok := this.m_object[object.M_id]["AppendRolesRef"]; !ok {
			this.m_object[object.M_id]["AppendRolesRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["AppendRolesRef"] = append(this.m_object[object.M_id]["AppendRolesRef"], xmlElement.M_rolesRef[i])
	}

	/** Init ref permissionsRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_permissionsRef); i++ {
		if _, ok := this.m_object[object.M_id]["AppendPermissionsRef"]; !ok {
			this.m_object[object.M_id]["AppendPermissionsRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["AppendPermissionsRef"] = append(this.m_object[object.M_id]["AppendPermissionsRef"], xmlElement.M_permissionsRef[i])
	}

	/** Init session **/
	for i := 0; i < len(xmlElement.M_sessions); i++ {
		if object.M_sessions == nil {
			object.M_sessions = make([]string, 0)
		}
		val := new(CargoEntities.Session)
		this.InitSession(object.GetUuid(), xmlElement.M_sessions[i], val)
		object.AppendSessions(val)

		/** association initialisation **/
		val.SetAccountPtr(object)
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Action **/
func (this *CargoEntitiesXmlFactory) InitAction(parentUuid string, xmlElement *CargoEntities.XsdAction, object *CargoEntities.Action) {
	log.Println("Initialize Action")
	object.TYPENAME = "CargoEntities.Action"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Action **/
	object.M_name = xmlElement.M_name

	/** Action **/
	object.M_doc = xmlElement.M_doc

	if xmlElement.M_accessType == "##Hidden" {
		object.M_accessType = CargoEntities.AccessType_Hidden
	} else if xmlElement.M_accessType == "##Public" {
		object.M_accessType = CargoEntities.AccessType_Public
	} else if xmlElement.M_accessType == "##Restricted" {
		object.M_accessType = CargoEntities.AccessType_Restricted
	} else {
		object.M_accessType = CargoEntities.AccessType_Hidden
	}
	if len(object.M_name) > 0 {
		this.m_references[object.M_name] = object
	}

	/** Init parameter **/
	for i := 0; i < len(xmlElement.M_parameters); i++ {
		if object.M_parameters == nil {
			object.M_parameters = make([]string, 0)
		}
		val := new(CargoEntities.Parameter)
		this.InitParameter(object.GetUuid(), xmlElement.M_parameters[i], val)
		object.AppendParameters(val)

		/** association initialisation **/
	}

	/** Init parameter **/
	for i := 0; i < len(xmlElement.M_results); i++ {
		if object.M_results == nil {
			object.M_results = make([]string, 0)
		}
		val := new(CargoEntities.Parameter)
		this.InitParameter(object.GetUuid(), xmlElement.M_results[i], val)
		object.AppendResults(val)

		/** association initialisation **/
	}
}

/** inititialisation of Session **/
func (this *CargoEntitiesXmlFactory) InitSession(parentUuid string, xmlElement *CargoEntities.XsdSession, object *CargoEntities.Session) {
	log.Println("Initialize Session")
	object.TYPENAME = "CargoEntities.Session"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Session **/
	object.M_id = xmlElement.M_id

	if xmlElement.M_sessionState == "##Online" {
		object.M_sessionState = CargoEntities.SessionState_Online
	} else if xmlElement.M_sessionState == "##Away" {
		object.M_sessionState = CargoEntities.SessionState_Away
	} else if xmlElement.M_sessionState == "##Offline" {
		object.M_sessionState = CargoEntities.SessionState_Offline
	} else {
		object.M_sessionState = CargoEntities.SessionState_Online
	}

	/** Session **/
	object.M_startTime = xmlElement.M_startTime

	/** Session **/
	object.M_endTime = xmlElement.M_endTime
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref computerRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_computerRef != nil {
		if _, ok := this.m_object[object.M_id]["SetComputerRef"]; !ok {
			this.m_object[object.M_id]["SetComputerRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["SetComputerRef"] = append(this.m_object[object.M_id]["SetComputerRef"], *xmlElement.M_computerRef)
	}

}

/** inititialisation of Computer **/
func (this *CargoEntitiesXmlFactory) InitComputer(parentUuid string, xmlElement *CargoEntities.XsdComputer, object *CargoEntities.Computer) {
	log.Println("Initialize Computer")
	object.TYPENAME = "CargoEntities.Computer"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Entity **/
	object.M_id = xmlElement.M_id

	if xmlElement.M_osType == "##Unknown" {
		object.M_osType = CargoEntities.OsType_Unknown
	} else if xmlElement.M_osType == "##Linux" {
		object.M_osType = CargoEntities.OsType_Linux
	} else if xmlElement.M_osType == "##Windows7" {
		object.M_osType = CargoEntities.OsType_Windows7
	} else if xmlElement.M_osType == "##Windows8" {
		object.M_osType = CargoEntities.OsType_Windows8
	} else if xmlElement.M_osType == "##Windows10" {
		object.M_osType = CargoEntities.OsType_Windows10
	} else if xmlElement.M_osType == "##OSX" {
		object.M_osType = CargoEntities.OsType_OSX
	} else if xmlElement.M_osType == "##IOS" {
		object.M_osType = CargoEntities.OsType_IOS
	} else {
		object.M_osType = CargoEntities.OsType_Unknown
	}

	if xmlElement.M_platformType == "##Unknown" {
		object.M_platformType = CargoEntities.PlatformType_Unknown
	} else if xmlElement.M_platformType == "##Tablet" {
		object.M_platformType = CargoEntities.PlatformType_Tablet
	} else if xmlElement.M_platformType == "##Phone" {
		object.M_platformType = CargoEntities.PlatformType_Phone
	} else if xmlElement.M_platformType == "##Desktop" {
		object.M_platformType = CargoEntities.PlatformType_Desktop
	} else if xmlElement.M_platformType == "##Laptop" {
		object.M_platformType = CargoEntities.PlatformType_Laptop
	} else {
		object.M_platformType = CargoEntities.PlatformType_Unknown
	}

	/** Computer **/
	object.M_name = xmlElement.M_name

	/** Computer **/
	object.M_ipv4 = xmlElement.M_ipv4
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Log **/
func (this *CargoEntitiesXmlFactory) InitLog(parentUuid string, xmlElement *CargoEntities.XsdLog, object *CargoEntities.Log) {
	log.Println("Initialize Log")
	object.TYPENAME = "CargoEntities.Log"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init logEntry **/
	for i := 0; i < len(xmlElement.M_entries); i++ {
		if object.M_entries == nil {
			object.M_entries = make([]string, 0)
		}
		val := new(CargoEntities.LogEntry)
		this.InitLogEntry(object.GetUuid(), xmlElement.M_entries[i], val)
		object.AppendEntries(val)

		/** association initialisation **/
	}
}

/** inititialisation of Entities **/
func (this *CargoEntitiesXmlFactory) InitEntities(parentUuid string, xmlElement *CargoEntities.XsdEntities, object *CargoEntities.Entities) {
	log.Println("Initialize Entities")
	object.TYPENAME = "CargoEntities.Entities"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Entities **/
	object.M_id = xmlElement.M_id

	/** Entities **/
	object.M_name = xmlElement.M_name

	/** Entities **/
	object.M_version = xmlElement.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init accountsRef **/
	for i := 0; i < len(xmlElement.M_entities_0); i++ {
		if object.M_entities == nil {
			object.M_entities = make([]string, 0)
		}
		val := new(CargoEntities.Account)
		this.InitAccount(object.GetUuid(), xmlElement.M_entities_0[i], val)
		object.AppendEntities(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init filesRef **/
	for i := 0; i < len(xmlElement.M_entities_1); i++ {
		if object.M_entities == nil {
			object.M_entities = make([]string, 0)
		}
		val := new(CargoEntities.File)
		this.InitFile(object.GetUuid(), xmlElement.M_entities_1[i], val)
		object.AppendEntities(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init project **/
	for i := 0; i < len(xmlElement.M_entities_2); i++ {
		if object.M_entities == nil {
			object.M_entities = make([]string, 0)
		}
		val := new(CargoEntities.Project)
		this.InitProject(object.GetUuid(), xmlElement.M_entities_2[i], val)
		object.AppendEntities(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init computerRef **/
	for i := 0; i < len(xmlElement.M_entities_3); i++ {
		if object.M_entities == nil {
			object.M_entities = make([]string, 0)
		}
		val := new(CargoEntities.Computer)
		this.InitComputer(object.GetUuid(), xmlElement.M_entities_3[i], val)
		object.AppendEntities(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init membersRef **/
	for i := 0; i < len(xmlElement.M_entities_4); i++ {
		if object.M_entities == nil {
			object.M_entities = make([]string, 0)
		}
		val := new(CargoEntities.User)
		this.InitUser(object.GetUuid(), xmlElement.M_entities_4[i], val)
		object.AppendEntities(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init memberOfRef **/
	for i := 0; i < len(xmlElement.M_entities_5); i++ {
		if object.M_entities == nil {
			object.M_entities = make([]string, 0)
		}
		val := new(CargoEntities.Group)
		this.InitGroup(object.GetUuid(), xmlElement.M_entities_5[i], val)
		object.AppendEntities(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init logEntry **/
	for i := 0; i < len(xmlElement.M_entities_6); i++ {
		if object.M_entities == nil {
			object.M_entities = make([]string, 0)
		}
		val := new(CargoEntities.LogEntry)
		this.InitLogEntry(object.GetUuid(), xmlElement.M_entities_6[i], val)
		object.AppendEntities(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init log **/
	for i := 0; i < len(xmlElement.M_entities_7); i++ {
		if object.M_entities == nil {
			object.M_entities = make([]string, 0)
		}
		val := new(CargoEntities.Log)
		this.InitLog(object.GetUuid(), xmlElement.M_entities_7[i], val)
		object.AppendEntities(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init rolesRef **/
	for i := 0; i < len(xmlElement.M_roles); i++ {
		if object.M_roles == nil {
			object.M_roles = make([]string, 0)
		}
		val := new(CargoEntities.Role)
		this.InitRole(object.GetUuid(), xmlElement.M_roles[i], val)
		object.AppendRoles(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init actionsRef **/
	for i := 0; i < len(xmlElement.M_actions); i++ {
		if object.M_actions == nil {
			object.M_actions = make([]string, 0)
		}
		val := new(CargoEntities.Action)
		this.InitAction(object.GetUuid(), xmlElement.M_actions[i], val)
		object.AppendActions(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init permissionsRef **/
	for i := 0; i < len(xmlElement.M_permissions); i++ {
		if object.M_permissions == nil {
			object.M_permissions = make([]string, 0)
		}
		val := new(CargoEntities.Permission)
		this.InitPermission(object.GetUuid(), xmlElement.M_permissions[i], val)
		object.AppendPermissions(val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}
}

/** inititialisation of Parameter **/
func (this *CargoEntitiesXmlFactory) InitParameter(parentUuid string, xmlElement *CargoEntities.XsdParameter, object *CargoEntities.Parameter) {
	log.Println("Initialize Parameter")
	object.TYPENAME = "CargoEntities.Parameter"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Parameter **/
	object.M_name = xmlElement.M_name

	/** Parameter **/
	object.M_type = xmlElement.M_type

	/** Parameter **/
	object.M_isArray = xmlElement.M_isArray
	if len(object.M_name) > 0 {
		this.m_references[object.M_name] = object
	}
}

/** inititialisation of Permission **/
func (this *CargoEntitiesXmlFactory) InitPermission(parentUuid string, xmlElement *CargoEntities.XsdPermission, object *CargoEntities.Permission) {
	log.Println("Initialize Permission")
	object.TYPENAME = "CargoEntities.Permission"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Permission **/
	object.M_id = xmlElement.M_id

	/** Permission **/
	object.M_types = xmlElement.M_types
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref accountsRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_accountsRef); i++ {
		if _, ok := this.m_object[object.M_id]["AppendAccountsRef"]; !ok {
			this.m_object[object.M_id]["AppendAccountsRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["AppendAccountsRef"] = append(this.m_object[object.M_id]["AppendAccountsRef"], xmlElement.M_accountsRef[i])
	}

}

/** inititialisation of File **/
func (this *CargoEntitiesXmlFactory) InitFile(parentUuid string, xmlElement *CargoEntities.XsdFile, object *CargoEntities.File) {
	log.Println("Initialize File")
	object.TYPENAME = "CargoEntities.File"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Entity **/
	object.M_id = xmlElement.M_id

	/** File **/
	object.M_name = xmlElement.M_name

	/** File **/
	object.M_path = xmlElement.M_path

	/** File **/
	object.M_size = xmlElement.M_size

	/** File **/
	object.M_modeTime = xmlElement.M_modeTime

	/** File **/
	object.M_isDir = xmlElement.M_isDir

	/** File **/
	object.M_checksum = xmlElement.M_checksum

	/** File **/
	object.M_data = xmlElement.M_data

	/** File **/
	object.M_thumbnail = xmlElement.M_thumbnail

	/** File **/
	object.M_mime = xmlElement.M_mime

	if xmlElement.M_fileType == "##dbFile" {
		object.M_fileType = CargoEntities.FileType_DbFile
	} else if xmlElement.M_fileType == "##diskFile" {
		object.M_fileType = CargoEntities.FileType_DiskFile
	} else {
		object.M_fileType = CargoEntities.FileType_DbFile
	}

	/** Init filesRef **/
	for i := 0; i < len(xmlElement.M_files); i++ {
		if object.M_files == nil {
			object.M_files = make([]string, 0)
		}
		val := new(CargoEntities.File)
		this.InitFile(object.GetUuid(), xmlElement.M_files[i], val)
		object.AppendFiles(val)

		/** association initialisation **/
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Project **/
func (this *CargoEntitiesXmlFactory) InitProject(parentUuid string, xmlElement *CargoEntities.XsdProject, object *CargoEntities.Project) {
	log.Println("Initialize Project")
	object.TYPENAME = "CargoEntities.Project"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Entity **/
	object.M_id = xmlElement.M_id

	/** Project **/
	object.M_name = xmlElement.M_name

	/** Init ref filesRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_filesRef); i++ {
		if _, ok := this.m_object[object.M_id]["AppendFilesRef"]; !ok {
			this.m_object[object.M_id]["AppendFilesRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["AppendFilesRef"] = append(this.m_object[object.M_id]["AppendFilesRef"], xmlElement.M_filesRef[i])
	}

	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of User **/
func (this *CargoEntitiesXmlFactory) InitUser(parentUuid string, xmlElement *CargoEntities.XsdUser, object *CargoEntities.User) {
	log.Println("Initialize User")
	object.TYPENAME = "CargoEntities.User"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Entity **/
	object.M_id = xmlElement.M_id

	/** User **/
	object.M_firstName = xmlElement.M_firstName

	/** User **/
	object.M_lastName = xmlElement.M_lastName

	/** User **/
	object.M_middle = xmlElement.M_middle

	/** User **/
	object.M_email = xmlElement.M_email

	/** User **/
	object.M_phone = xmlElement.M_phone

	/** Init ref memberOfRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_memberOfRef); i++ {
		if _, ok := this.m_object[object.M_id]["AppendMemberOfRef"]; !ok {
			this.m_object[object.M_id]["AppendMemberOfRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["AppendMemberOfRef"] = append(this.m_object[object.M_id]["AppendMemberOfRef"], xmlElement.M_memberOfRef[i])
	}

	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Group **/
func (this *CargoEntitiesXmlFactory) InitGroup(parentUuid string, xmlElement *CargoEntities.XsdGroup, object *CargoEntities.Group) {
	log.Println("Initialize Group")
	object.TYPENAME = "CargoEntities.Group"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Entity **/
	object.M_id = xmlElement.M_id

	/** Group **/
	object.M_name = xmlElement.M_name

	/** Init ref membersRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_membersRef); i++ {
		if _, ok := this.m_object[object.M_id]["AppendMembersRef"]; !ok {
			this.m_object[object.M_id]["AppendMembersRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["AppendMembersRef"] = append(this.m_object[object.M_id]["AppendMembersRef"], xmlElement.M_membersRef[i])
	}

	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Role **/
func (this *CargoEntitiesXmlFactory) InitRole(parentUuid string, xmlElement *CargoEntities.XsdRole, object *CargoEntities.Role) {
	log.Println("Initialize Role")
	object.TYPENAME = "CargoEntities.Role"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** Role **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref accountsRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_accountsRef); i++ {
		if _, ok := this.m_object[object.M_id]["AppendAccountsRef"]; !ok {
			this.m_object[object.M_id]["AppendAccountsRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["AppendAccountsRef"] = append(this.m_object[object.M_id]["AppendAccountsRef"], xmlElement.M_accountsRef[i])
	}

	/** Init ref actionsRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_actionsRef); i++ {
		if _, ok := this.m_object[object.M_id]["AppendActionsRef"]; !ok {
			this.m_object[object.M_id]["AppendActionsRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["AppendActionsRef"] = append(this.m_object[object.M_id]["AppendActionsRef"], xmlElement.M_actionsRef[i])
	}

}

/** inititialisation of LogEntry **/
func (this *CargoEntitiesXmlFactory) InitLogEntry(parentUuid string, xmlElement *CargoEntities.XsdLogEntry, object *CargoEntities.LogEntry) {
	log.Println("Initialize LogEntry")
	object.TYPENAME = "CargoEntities.LogEntry"
	object.SetParentUuid(parentUuid)
	object.SetUuidGenerator(this.m_generateUuid)
	object.SetEntitySetter(this.m_setEntity)
	object.SetEntityGetter(this.m_getEntityByUuid)

	/** LogEntry **/
	object.M_creationTime = xmlElement.M_creationTime
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref entityRef **/
	if len(object.M_id) == 0 {
		this.m_references[object.GetUuid()] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_entityRef) > 0 {
		if _, ok := this.m_object[object.M_id]["SetEntityRef"]; !ok {
			this.m_object[object.M_id]["SetEntityRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["SetEntityRef"] = append(this.m_object[object.M_id]["SetEntityRef"], xmlElement.M_entityRef)
	}

}

/** serialysation of User **/
func (this *CargoEntitiesXmlFactory) SerialyzeUser(xmlElement *CargoEntities.XsdUser, object *CargoEntities.User) {
	if xmlElement == nil {
		return
	}

	/** Entity **/
	xmlElement.M_id = object.M_id

	/** User **/
	xmlElement.M_firstName = object.M_firstName

	/** User **/
	xmlElement.M_lastName = object.M_lastName

	/** User **/
	xmlElement.M_middle = object.M_middle

	/** User **/
	xmlElement.M_email = object.M_email

	/** User **/
	xmlElement.M_phone = object.M_phone

	/** Serialyze ref memberOfRef **/
	xmlElement.M_memberOfRef = object.M_memberOfRef

	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Role **/
func (this *CargoEntitiesXmlFactory) SerialyzeRole(xmlElement *CargoEntities.XsdRole, object *CargoEntities.Role) {
	if xmlElement == nil {
		return
	}

	/** Role **/
	xmlElement.M_id = object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref accountsRef **/
	xmlElement.M_accountsRef = object.M_accountsRef

	/** Serialyze ref actionsRef **/
	xmlElement.M_actionsRef = object.M_actionsRef

}

/** serialysation of Permission **/
func (this *CargoEntitiesXmlFactory) SerialyzePermission(xmlElement *CargoEntities.XsdPermission, object *CargoEntities.Permission) {
	if xmlElement == nil {
		return
	}

	/** Permission **/
	xmlElement.M_id = object.M_id

	/** Permission **/
	xmlElement.M_types = object.M_types
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref accountsRef **/
	xmlElement.M_accountsRef = object.M_accountsRef

}

/** serialysation of Computer **/
func (this *CargoEntitiesXmlFactory) SerialyzeComputer(xmlElement *CargoEntities.XsdComputer, object *CargoEntities.Computer) {
	if xmlElement == nil {
		return
	}

	/** Entity **/
	xmlElement.M_id = object.M_id

	if object.M_osType == CargoEntities.OsType_Unknown {
		xmlElement.M_osType = "##Unknown"
	} else if object.M_osType == CargoEntities.OsType_Linux {
		xmlElement.M_osType = "##Linux"
	} else if object.M_osType == CargoEntities.OsType_Windows7 {
		xmlElement.M_osType = "##Windows7"
	} else if object.M_osType == CargoEntities.OsType_Windows8 {
		xmlElement.M_osType = "##Windows8"
	} else if object.M_osType == CargoEntities.OsType_Windows10 {
		xmlElement.M_osType = "##Windows10"
	} else if object.M_osType == CargoEntities.OsType_OSX {
		xmlElement.M_osType = "##OSX"
	} else if object.M_osType == CargoEntities.OsType_IOS {
		xmlElement.M_osType = "##IOS"
	} else {
		xmlElement.M_osType = "##Unknown"
	}

	if object.M_platformType == CargoEntities.PlatformType_Unknown {
		xmlElement.M_platformType = "##Unknown"
	} else if object.M_platformType == CargoEntities.PlatformType_Tablet {
		xmlElement.M_platformType = "##Tablet"
	} else if object.M_platformType == CargoEntities.PlatformType_Phone {
		xmlElement.M_platformType = "##Phone"
	} else if object.M_platformType == CargoEntities.PlatformType_Desktop {
		xmlElement.M_platformType = "##Desktop"
	} else if object.M_platformType == CargoEntities.PlatformType_Laptop {
		xmlElement.M_platformType = "##Laptop"
	} else {
		xmlElement.M_platformType = "##Unknown"
	}

	/** Computer **/
	xmlElement.M_name = object.M_name

	/** Computer **/
	xmlElement.M_ipv4 = object.M_ipv4
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Log **/
func (this *CargoEntitiesXmlFactory) SerialyzeLog(xmlElement *CargoEntities.XsdLog, object *CargoEntities.Log) {
	if xmlElement == nil {
		return
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze LogEntry **/
	if len(object.M_entries) > 0 {
		xmlElement.M_entries = make([]*CargoEntities.XsdLogEntry, 0)
	}

	/** Now I will save the value of entries **/
	for i := 0; i < len(object.M_entries); i++ {
		xmlElement.M_entries = append(xmlElement.M_entries, new(CargoEntities.XsdLogEntry))
		this.SerialyzeLogEntry(xmlElement.M_entries[i], object.GetEntries()[i])
	}
}

/** serialysation of Parameter **/
func (this *CargoEntitiesXmlFactory) SerialyzeParameter(xmlElement *CargoEntities.XsdParameter, object *CargoEntities.Parameter) {
	if xmlElement == nil {
		return
	}

	/** Parameter **/
	xmlElement.M_name = object.M_name

	/** Parameter **/
	xmlElement.M_type = object.M_type

	/** Parameter **/
	xmlElement.M_isArray = object.M_isArray
	if len(object.M_name) > 0 {
		this.m_references[object.M_name] = object
	}
}

/** serialysation of Entities **/
func (this *CargoEntitiesXmlFactory) SerialyzeEntities(xmlElement *CargoEntities.XsdEntities, object *CargoEntities.Entities) {
	if xmlElement == nil {
		return
	}

	/** Entities **/
	xmlElement.M_id = object.M_id

	/** Entities **/
	xmlElement.M_name = object.M_name

	/** Entities **/
	xmlElement.M_version = object.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze Entity **/
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_0 = make([]*CargoEntities.XsdAccount, 0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_1 = make([]*CargoEntities.XsdFile, 0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_2 = make([]*CargoEntities.XsdProject, 0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_3 = make([]*CargoEntities.XsdComputer, 0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_4 = make([]*CargoEntities.XsdUser, 0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_5 = make([]*CargoEntities.XsdGroup, 0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_6 = make([]*CargoEntities.XsdLogEntry, 0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_7 = make([]*CargoEntities.XsdLog, 0)
	}

	/** Now I will save the value of entities **/
	for i := 0; i < len(object.M_entities); i++ {
		switch v := object.GetEntities()[i].(type) {
		case *CargoEntities.Account:
			xmlElement.M_entities_0 = append(xmlElement.M_entities_0, new(CargoEntities.XsdAccount))
			this.SerialyzeAccount(xmlElement.M_entities_0[len(xmlElement.M_entities_0)-1], v)
		case *CargoEntities.File:
			xmlElement.M_entities_1 = append(xmlElement.M_entities_1, new(CargoEntities.XsdFile))
			this.SerialyzeFile(xmlElement.M_entities_1[len(xmlElement.M_entities_1)-1], v)
		case *CargoEntities.Project:
			xmlElement.M_entities_2 = append(xmlElement.M_entities_2, new(CargoEntities.XsdProject))
			this.SerialyzeProject(xmlElement.M_entities_2[len(xmlElement.M_entities_2)-1], v)
		case *CargoEntities.Computer:
			xmlElement.M_entities_3 = append(xmlElement.M_entities_3, new(CargoEntities.XsdComputer))
			this.SerialyzeComputer(xmlElement.M_entities_3[len(xmlElement.M_entities_3)-1], v)
		case *CargoEntities.User:
			xmlElement.M_entities_4 = append(xmlElement.M_entities_4, new(CargoEntities.XsdUser))
			this.SerialyzeUser(xmlElement.M_entities_4[len(xmlElement.M_entities_4)-1], v)
		case *CargoEntities.Group:
			xmlElement.M_entities_5 = append(xmlElement.M_entities_5, new(CargoEntities.XsdGroup))
			this.SerialyzeGroup(xmlElement.M_entities_5[len(xmlElement.M_entities_5)-1], v)
		case *CargoEntities.LogEntry:
			xmlElement.M_entities_6 = append(xmlElement.M_entities_6, new(CargoEntities.XsdLogEntry))
			this.SerialyzeLogEntry(xmlElement.M_entities_6[len(xmlElement.M_entities_6)-1], v)
		case *CargoEntities.Log:
			xmlElement.M_entities_7 = append(xmlElement.M_entities_7, new(CargoEntities.XsdLog))
			this.SerialyzeLog(xmlElement.M_entities_7[len(xmlElement.M_entities_7)-1], v)
		}
	}

	/** Serialyze Role **/
	if len(object.M_roles) > 0 {
		xmlElement.M_roles = make([]*CargoEntities.XsdRole, 0)
	}

	/** Now I will save the value of roles **/
	for i := 0; i < len(object.M_roles); i++ {
		xmlElement.M_roles = append(xmlElement.M_roles, new(CargoEntities.XsdRole))
		this.SerialyzeRole(xmlElement.M_roles[i], object.GetRoles()[i])
	}

	/** Serialyze Action **/
	if len(object.M_actions) > 0 {
		xmlElement.M_actions = make([]*CargoEntities.XsdAction, 0)
	}

	/** Now I will save the value of actions **/
	for i := 0; i < len(object.M_actions); i++ {
		xmlElement.M_actions = append(xmlElement.M_actions, new(CargoEntities.XsdAction))
		this.SerialyzeAction(xmlElement.M_actions[i], object.GetActions()[i])
	}

	/** Serialyze Permission **/
	if len(object.M_permissions) > 0 {
		xmlElement.M_permissions = make([]*CargoEntities.XsdPermission, 0)
	}

	/** Now I will save the value of permissions **/
	for i := 0; i < len(object.M_permissions); i++ {
		xmlElement.M_permissions = append(xmlElement.M_permissions, new(CargoEntities.XsdPermission))
		this.SerialyzePermission(xmlElement.M_permissions[i], object.GetPermissions()[i])
	}
}

/** serialysation of Project **/
func (this *CargoEntitiesXmlFactory) SerialyzeProject(xmlElement *CargoEntities.XsdProject, object *CargoEntities.Project) {
	if xmlElement == nil {
		return
	}

	/** Entity **/
	xmlElement.M_id = object.M_id

	/** Project **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref filesRef **/
	xmlElement.M_filesRef = object.M_filesRef

	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Group **/
func (this *CargoEntitiesXmlFactory) SerialyzeGroup(xmlElement *CargoEntities.XsdGroup, object *CargoEntities.Group) {
	if xmlElement == nil {
		return
	}

	/** Entity **/
	xmlElement.M_id = object.M_id

	/** Group **/
	xmlElement.M_name = object.M_name

	/** Serialyze ref membersRef **/
	xmlElement.M_membersRef = object.M_membersRef

	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Account **/
func (this *CargoEntitiesXmlFactory) SerialyzeAccount(xmlElement *CargoEntities.XsdAccount, object *CargoEntities.Account) {
	if xmlElement == nil {
		return
	}

	/** Entity **/
	xmlElement.M_id = object.M_id

	/** Account **/
	xmlElement.M_name = object.M_name

	/** Account **/
	xmlElement.M_password = object.M_password

	/** Account **/
	xmlElement.M_email = object.M_email

	/** Serialyze ref userRef **/
	xmlElement.M_userRef = &object.M_userRef

	/** Serialyze ref rolesRef **/
	xmlElement.M_rolesRef = object.M_rolesRef

	/** Serialyze ref permissionsRef **/
	xmlElement.M_permissionsRef = object.M_permissionsRef

	/** Serialyze Session **/
	if len(object.M_sessions) > 0 {
		xmlElement.M_sessions = make([]*CargoEntities.XsdSession, 0)
	}

	/** Now I will save the value of sessions **/
	for i := 0; i < len(object.M_sessions); i++ {
		xmlElement.M_sessions = append(xmlElement.M_sessions, new(CargoEntities.XsdSession))
		this.SerialyzeSession(xmlElement.M_sessions[i], object.GetSessions()[i])
	}

	/** Serialyze Message **/

	/** Now I will save the value of messages **/
	for i := 0; i < len(object.M_messages); i++ {
		switch v := object.GetMessages()[i].(type) {
		}
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Session **/
func (this *CargoEntitiesXmlFactory) SerialyzeSession(xmlElement *CargoEntities.XsdSession, object *CargoEntities.Session) {
	if xmlElement == nil {
		return
	}

	/** Session **/
	xmlElement.M_id = object.M_id

	if object.M_sessionState == CargoEntities.SessionState_Online {
		xmlElement.M_sessionState = "##Online"
	} else if object.M_sessionState == CargoEntities.SessionState_Away {
		xmlElement.M_sessionState = "##Away"
	} else if object.M_sessionState == CargoEntities.SessionState_Offline {
		xmlElement.M_sessionState = "##Offline"
	} else {
		xmlElement.M_sessionState = "##Online"
	}

	/** Session **/
	xmlElement.M_startTime = object.M_startTime

	/** Session **/
	xmlElement.M_endTime = object.M_endTime
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref computerRef **/
	xmlElement.M_computerRef = &object.M_computerRef

}

/** serialysation of File **/
func (this *CargoEntitiesXmlFactory) SerialyzeFile(xmlElement *CargoEntities.XsdFile, object *CargoEntities.File) {
	if xmlElement == nil {
		return
	}

	/** Entity **/
	xmlElement.M_id = object.M_id

	/** File **/
	xmlElement.M_name = object.M_name

	/** File **/
	xmlElement.M_path = object.M_path

	/** File **/
	xmlElement.M_size = object.M_size

	/** File **/
	xmlElement.M_modeTime = object.M_modeTime

	/** File **/
	xmlElement.M_isDir = object.M_isDir

	/** File **/
	xmlElement.M_checksum = object.M_checksum

	/** File **/
	xmlElement.M_data = object.M_data

	/** File **/
	xmlElement.M_thumbnail = object.M_thumbnail

	/** File **/
	xmlElement.M_mime = object.M_mime

	if object.M_fileType == CargoEntities.FileType_DbFile {
		xmlElement.M_fileType = "##dbFile"
	} else if object.M_fileType == CargoEntities.FileType_DiskFile {
		xmlElement.M_fileType = "##diskFile"
	} else {
		xmlElement.M_fileType = "##dbFile"
	}

	/** Serialyze File **/
	if len(object.M_files) > 0 {
		xmlElement.M_files = make([]*CargoEntities.XsdFile, 0)
	}

	/** Now I will save the value of files **/
	for i := 0; i < len(object.M_files); i++ {
		xmlElement.M_files = append(xmlElement.M_files, new(CargoEntities.XsdFile))
		this.SerialyzeFile(xmlElement.M_files[i], object.GetFiles()[i])
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of LogEntry **/
func (this *CargoEntitiesXmlFactory) SerialyzeLogEntry(xmlElement *CargoEntities.XsdLogEntry, object *CargoEntities.LogEntry) {
	if xmlElement == nil {
		return
	}

	/** LogEntry **/
	xmlElement.M_creationTime = object.M_creationTime
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref entityRef **/
	xmlElement.M_entityRef = object.M_entityRef

}

/** serialysation of Action **/
func (this *CargoEntitiesXmlFactory) SerialyzeAction(xmlElement *CargoEntities.XsdAction, object *CargoEntities.Action) {
	if xmlElement == nil {
		return
	}

	/** Action **/
	xmlElement.M_name = object.M_name

	/** Action **/
	xmlElement.M_doc = object.M_doc

	if object.M_accessType == CargoEntities.AccessType_Hidden {
		xmlElement.M_accessType = "##Hidden"
	} else if object.M_accessType == CargoEntities.AccessType_Public {
		xmlElement.M_accessType = "##Public"
	} else if object.M_accessType == CargoEntities.AccessType_Restricted {
		xmlElement.M_accessType = "##Restricted"
	} else {
		xmlElement.M_accessType = "##Hidden"
	}
	if len(object.M_name) > 0 {
		this.m_references[object.M_name] = object
	}

	/** Serialyze Parameter **/
	if len(object.M_parameters) > 0 {
		xmlElement.M_parameters = make([]*CargoEntities.XsdParameter, 0)
	}

	/** Now I will save the value of parameters **/
	for i := 0; i < len(object.M_parameters); i++ {
		xmlElement.M_parameters = append(xmlElement.M_parameters, new(CargoEntities.XsdParameter))
		this.SerialyzeParameter(xmlElement.M_parameters[i], object.GetParameters()[i])
	}

	/** Serialyze Parameter **/
	if len(object.M_results) > 0 {
		xmlElement.M_results = make([]*CargoEntities.XsdParameter, 0)
	}

	/** Now I will save the value of results **/
	for i := 0; i < len(object.M_results); i++ {
		xmlElement.M_results = append(xmlElement.M_results, new(CargoEntities.XsdParameter))
		this.SerialyzeParameter(xmlElement.M_results[i], object.GetResults()[i])
	}
}
