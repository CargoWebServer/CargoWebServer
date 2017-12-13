// +build Config

package Server

import (
	"encoding/json"
	//	"log"
	"strings"
	"unsafe"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
)

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_ConfigurationEntityPrototype() {

	var configurationEntityProto EntityPrototype
	configurationEntityProto.TypeName = "Config.Configuration"
	configurationEntityProto.IsAbstract = true
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "Config.SmtpConfiguration")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "Config.DataStoreConfiguration")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "Config.LdapConfiguration")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "Config.OAuth2Configuration")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "Config.ServiceConfiguration")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "Config.ScheduledTask")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "Config.ApplicationConfiguration")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "Config.ServerConfiguration")
	configurationEntityProto.Ids = append(configurationEntityProto.Ids, "UUID")
	configurationEntityProto.Fields = append(configurationEntityProto.Fields, "UUID")
	configurationEntityProto.FieldsType = append(configurationEntityProto.FieldsType, "xs.string")
	configurationEntityProto.FieldsOrder = append(configurationEntityProto.FieldsOrder, 0)
	configurationEntityProto.FieldsVisibility = append(configurationEntityProto.FieldsVisibility, false)
	configurationEntityProto.FieldsDefaultValue = append(configurationEntityProto.FieldsDefaultValue, "")
	configurationEntityProto.Indexs = append(configurationEntityProto.Indexs, "ParentUuid")
	configurationEntityProto.Fields = append(configurationEntityProto.Fields, "ParentUuid")
	configurationEntityProto.FieldsType = append(configurationEntityProto.FieldsType, "xs.string")
	configurationEntityProto.FieldsOrder = append(configurationEntityProto.FieldsOrder, 1)
	configurationEntityProto.FieldsVisibility = append(configurationEntityProto.FieldsVisibility, false)
	configurationEntityProto.FieldsDefaultValue = append(configurationEntityProto.FieldsDefaultValue, "")
	configurationEntityProto.Fields = append(configurationEntityProto.Fields, "ParentLnk")
	configurationEntityProto.FieldsType = append(configurationEntityProto.FieldsType, "xs.string")
	configurationEntityProto.FieldsOrder = append(configurationEntityProto.FieldsOrder, 2)
	configurationEntityProto.FieldsVisibility = append(configurationEntityProto.FieldsVisibility, false)
	configurationEntityProto.FieldsDefaultValue = append(configurationEntityProto.FieldsDefaultValue, "")

	/** members of Configuration **/
	configurationEntityProto.Ids = append(configurationEntityProto.Ids, "M_id")
	configurationEntityProto.FieldsOrder = append(configurationEntityProto.FieldsOrder, 3)
	configurationEntityProto.FieldsVisibility = append(configurationEntityProto.FieldsVisibility, true)
	configurationEntityProto.Fields = append(configurationEntityProto.Fields, "M_id")
	configurationEntityProto.FieldsType = append(configurationEntityProto.FieldsType, "xs.ID")
	configurationEntityProto.FieldsDefaultValue = append(configurationEntityProto.FieldsDefaultValue, "")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&configurationEntityProto)

}

////////////////////////////////////////////////////////////////////////////////
//              			SmtpConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_SmtpConfigurationEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.SmtpConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewConfigSmtpConfigurationEntity(parentUuid string, objectId string, object interface{}) *Config_SmtpConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigSmtpConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.SmtpConfiguration).TYPENAME = "Config.SmtpConfiguration"
		object.(*Config.SmtpConfiguration).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.SmtpConfiguration", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.SmtpConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.SmtpConfiguration).UUID
			}
			return val.(*Config_SmtpConfigurationEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_SmtpConfigurationEntity)
	if object == nil {
		entity.object = new(Config.SmtpConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.SmtpConfiguration)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.SmtpConfiguration"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_SmtpConfigurationEntity) GetTypeName() string {
	return "Config.SmtpConfiguration"
}
func (this *Config_SmtpConfigurationEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_SmtpConfigurationEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_SmtpConfigurationEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_SmtpConfigurationEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_SmtpConfigurationEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_SmtpConfigurationEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_SmtpConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_SmtpConfigurationEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_SmtpConfigurationEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_SmtpConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_SmtpConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_SmtpConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_SmtpConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_SmtpConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_SmtpConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_SmtpConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_SmtpConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_SmtpConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_SmtpConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_SmtpConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_SmtpConfigurationEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_SmtpConfigurationEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_SmtpConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.SmtpConfiguration"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_SmtpConfigurationEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_SmtpConfigurationEntityPrototype() {

	var smtpConfigurationEntityProto EntityPrototype
	smtpConfigurationEntityProto.TypeName = "Config.SmtpConfiguration"
	smtpConfigurationEntityProto.SuperTypeNames = append(smtpConfigurationEntityProto.SuperTypeNames, "Config.Configuration")
	smtpConfigurationEntityProto.Ids = append(smtpConfigurationEntityProto.Ids, "UUID")
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "UUID")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 0)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, false)
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "")
	smtpConfigurationEntityProto.Indexs = append(smtpConfigurationEntityProto.Indexs, "ParentUuid")
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "ParentUuid")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 1)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, false)
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "")
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "ParentLnk")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 2)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, false)
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of Configuration **/
	smtpConfigurationEntityProto.Ids = append(smtpConfigurationEntityProto.Ids, "M_id")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 3)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_id")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.ID")
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of SmtpConfiguration **/
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 4)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_textEncoding")
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "1")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "enum:Encoding_UTF8:Encoding_WINDOWS_1250:Encoding_WINDOWS_1251:Encoding_WINDOWS_1252:Encoding_WINDOWS_1253:Encoding_WINDOWS_1254:Encoding_WINDOWS_1255:Encoding_WINDOWS_1256:Encoding_WINDOWS_1257:Encoding_WINDOWS_1258:Encoding_ISO8859_1:Encoding_ISO8859_2:Encoding_ISO8859_3:Encoding_ISO8859_4:Encoding_ISO8859_5:Encoding_ISO8859_6:Encoding_ISO8859_7:Encoding_ISO8859_8:Encoding_ISO8859_9:Encoding_ISO8859_10:Encoding_ISO8859_13:Encoding_ISO8859_14:Encoding_ISO8859_15:Encoding_ISO8859_16:Encoding_KOI8R:Encoding_KOI8U")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 5)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_hostName")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 6)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_ipv4")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 7)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_port")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.int")
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "0")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 8)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_user")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 9)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_pwd")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "")

	/** associations of SmtpConfiguration **/
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 10)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, false)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_parentPtr")
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "undefined")
	smtpConfigurationEntityProto.FieldsDefaultValue = append(smtpConfigurationEntityProto.FieldsDefaultValue, "undefined")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "Config.Configurations:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&smtpConfigurationEntityProto)

}

/** Create **/
func (this *Config_SmtpConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.SmtpConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of SmtpConfiguration **/
	query.Fields = append(query.Fields, "M_textEncoding")
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_port")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")

	/** associations of SmtpConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var SmtpConfigurationInfo []interface{}

	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.GetParentPtr().GetUuid())
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.GetParentLnk())
	} else {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, "")
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, "")
	}

	/** members of Configuration **/
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_id)

	/** members of SmtpConfiguration **/

	/** Save textEncoding type Encoding **/
	if this.object.M_textEncoding == Config.Encoding_UTF8 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 0)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1250 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 1)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1251 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 2)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1252 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 3)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1253 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 4)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1254 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 5)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1255 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 6)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1256 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 7)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1257 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 8)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1258 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 9)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_1 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 10)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_2 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 11)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_3 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 12)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_4 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 13)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_5 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 14)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_6 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 15)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_7 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 16)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_8 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 17)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_9 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 18)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_10 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 19)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_13 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 20)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_14 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 21)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_15 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 22)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_16 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 23)
	} else if this.object.M_textEncoding == Config.Encoding_KOI8R {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 24)
	} else if this.object.M_textEncoding == Config.Encoding_KOI8U {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 25)
	} else {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, 0)
	}
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_hostName)
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_ipv4)
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_port)
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_user)
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_pwd)

	/** associations of SmtpConfiguration **/

	/** Save parent type Configurations **/
	if len(this.object.M_parentPtr) > 0 {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_parentPtr)
	} else {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), SmtpConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), SmtpConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_SmtpConfigurationEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_SmtpConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.SmtpConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of SmtpConfiguration **/
	query.Fields = append(query.Fields, "M_textEncoding")
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_port")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")

	/** associations of SmtpConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of SmtpConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.SmtpConfiguration)
		this.object.TYPENAME = "Config.SmtpConfiguration"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of SmtpConfiguration **/

		/** textEncoding **/
		if results[0][4] != nil {
			enumIndex := results[0][4].(int)
			if enumIndex == 0 {
				this.object.M_textEncoding = Config.Encoding_UTF8
			} else if enumIndex == 1 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1250
			} else if enumIndex == 2 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1251
			} else if enumIndex == 3 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1252
			} else if enumIndex == 4 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1253
			} else if enumIndex == 5 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1254
			} else if enumIndex == 6 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1255
			} else if enumIndex == 7 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1256
			} else if enumIndex == 8 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1257
			} else if enumIndex == 9 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1258
			} else if enumIndex == 10 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_1
			} else if enumIndex == 11 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_2
			} else if enumIndex == 12 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_3
			} else if enumIndex == 13 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_4
			} else if enumIndex == 14 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_5
			} else if enumIndex == 15 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_6
			} else if enumIndex == 16 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_7
			} else if enumIndex == 17 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_8
			} else if enumIndex == 18 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_9
			} else if enumIndex == 19 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_10
			} else if enumIndex == 20 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_13
			} else if enumIndex == 21 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_14
			} else if enumIndex == 22 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_15
			} else if enumIndex == 23 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_16
			} else if enumIndex == 24 {
				this.object.M_textEncoding = Config.Encoding_KOI8R
			} else if enumIndex == 25 {
				this.object.M_textEncoding = Config.Encoding_KOI8U
			}
		}

		/** hostName **/
		if results[0][5] != nil {
			this.object.M_hostName = results[0][5].(string)
		}

		/** ipv4 **/
		if results[0][6] != nil {
			this.object.M_ipv4 = results[0][6].(string)
		}

		/** port **/
		if results[0][7] != nil {
			this.object.M_port = results[0][7].(int)
		}

		/** user **/
		if results[0][8] != nil {
			this.object.M_user = results[0][8].(string)
		}

		/** pwd **/
		if results[0][9] != nil {
			this.object.M_pwd = results[0][9].(string)
		}

		/** associations of SmtpConfiguration **/

		/** parentPtr **/
		if results[0][10] != nil {
			id := results[0][10].(string)
			if len(id) > 0 {
				refTypeName := "Config.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigSmtpConfigurationEntityFromObject(object *Config.SmtpConfiguration) *Config_SmtpConfigurationEntity {
	return this.NewConfigSmtpConfigurationEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_SmtpConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigSmtpConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.SmtpConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_SmtpConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_SmtpConfigurationEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			DataStoreConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_DataStoreConfigurationEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.DataStoreConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewConfigDataStoreConfigurationEntity(parentUuid string, objectId string, object interface{}) *Config_DataStoreConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigDataStoreConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.DataStoreConfiguration).TYPENAME = "Config.DataStoreConfiguration"
		object.(*Config.DataStoreConfiguration).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.DataStoreConfiguration", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.DataStoreConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.DataStoreConfiguration).UUID
			}
			return val.(*Config_DataStoreConfigurationEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_DataStoreConfigurationEntity)
	if object == nil {
		entity.object = new(Config.DataStoreConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.DataStoreConfiguration)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.DataStoreConfiguration"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_DataStoreConfigurationEntity) GetTypeName() string {
	return "Config.DataStoreConfiguration"
}
func (this *Config_DataStoreConfigurationEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_DataStoreConfigurationEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_DataStoreConfigurationEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_DataStoreConfigurationEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_DataStoreConfigurationEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_DataStoreConfigurationEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_DataStoreConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_DataStoreConfigurationEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_DataStoreConfigurationEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_DataStoreConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_DataStoreConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_DataStoreConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_DataStoreConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_DataStoreConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_DataStoreConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_DataStoreConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_DataStoreConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_DataStoreConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_DataStoreConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_DataStoreConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_DataStoreConfigurationEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_DataStoreConfigurationEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_DataStoreConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.DataStoreConfiguration"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_DataStoreConfigurationEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_DataStoreConfigurationEntityPrototype() {

	var dataStoreConfigurationEntityProto EntityPrototype
	dataStoreConfigurationEntityProto.TypeName = "Config.DataStoreConfiguration"
	dataStoreConfigurationEntityProto.SuperTypeNames = append(dataStoreConfigurationEntityProto.SuperTypeNames, "Config.Configuration")
	dataStoreConfigurationEntityProto.Ids = append(dataStoreConfigurationEntityProto.Ids, "UUID")
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "UUID")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 0)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, false)
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "")
	dataStoreConfigurationEntityProto.Indexs = append(dataStoreConfigurationEntityProto.Indexs, "ParentUuid")
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "ParentUuid")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 1)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, false)
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "")
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "ParentLnk")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 2)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, false)
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of Configuration **/
	dataStoreConfigurationEntityProto.Ids = append(dataStoreConfigurationEntityProto.Ids, "M_id")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 3)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_id")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.ID")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of DataStoreConfiguration **/
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 4)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_dataStoreType")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "1")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "enum:DataStoreType_SQL_STORE:DataStoreType_KEY_VALUE_STORE")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 5)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_dataStoreVendor")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "1")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "enum:DataStoreVendor_MYCELIUS:DataStoreVendor_MYSQL:DataStoreVendor_MSSQL:DataStoreVendor_ODBC:DataStoreVendor_KNOWLEDGEBASE")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 6)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_textEncoding")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "1")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "enum:Encoding_UTF8:Encoding_WINDOWS_1250:Encoding_WINDOWS_1251:Encoding_WINDOWS_1252:Encoding_WINDOWS_1253:Encoding_WINDOWS_1254:Encoding_WINDOWS_1255:Encoding_WINDOWS_1256:Encoding_WINDOWS_1257:Encoding_WINDOWS_1258:Encoding_ISO8859_1:Encoding_ISO8859_2:Encoding_ISO8859_3:Encoding_ISO8859_4:Encoding_ISO8859_5:Encoding_ISO8859_6:Encoding_ISO8859_7:Encoding_ISO8859_8:Encoding_ISO8859_9:Encoding_ISO8859_10:Encoding_ISO8859_13:Encoding_ISO8859_14:Encoding_ISO8859_15:Encoding_ISO8859_16:Encoding_KOI8R:Encoding_KOI8U")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 7)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_storeName")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 8)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_hostName")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 9)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_ipv4")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 10)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_user")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 11)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_pwd")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 12)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_port")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.int")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "0")

	/** associations of DataStoreConfiguration **/
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 13)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, false)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_parentPtr")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "undefined")
	dataStoreConfigurationEntityProto.FieldsDefaultValue = append(dataStoreConfigurationEntityProto.FieldsDefaultValue, "undefined")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "Config.Configurations:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&dataStoreConfigurationEntityProto)

}

/** Create **/
func (this *Config_DataStoreConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.DataStoreConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of DataStoreConfiguration **/
	query.Fields = append(query.Fields, "M_dataStoreType")
	query.Fields = append(query.Fields, "M_dataStoreVendor")
	query.Fields = append(query.Fields, "M_textEncoding")
	query.Fields = append(query.Fields, "M_storeName")
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")
	query.Fields = append(query.Fields, "M_port")

	/** associations of DataStoreConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var DataStoreConfigurationInfo []interface{}

	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.GetParentPtr().GetUuid())
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.GetParentLnk())
	} else {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, "")
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, "")
	}

	/** members of Configuration **/
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_id)

	/** members of DataStoreConfiguration **/

	/** Save dataStoreType type DataStoreType **/
	if this.object.M_dataStoreType == Config.DataStoreType_SQL_STORE {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 0)
	} else if this.object.M_dataStoreType == Config.DataStoreType_KEY_VALUE_STORE {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 1)
	} else {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 0)
	}

	/** Save dataStoreVendor type DataStoreVendor **/
	if this.object.M_dataStoreVendor == Config.DataStoreVendor_MYCELIUS {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 0)
	} else if this.object.M_dataStoreVendor == Config.DataStoreVendor_MYSQL {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 1)
	} else if this.object.M_dataStoreVendor == Config.DataStoreVendor_MSSQL {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 2)
	} else if this.object.M_dataStoreVendor == Config.DataStoreVendor_ODBC {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 3)
	} else if this.object.M_dataStoreVendor == Config.DataStoreVendor_KNOWLEDGEBASE {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 4)
	} else {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 0)
	}

	/** Save textEncoding type Encoding **/
	if this.object.M_textEncoding == Config.Encoding_UTF8 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 0)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1250 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 1)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1251 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 2)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1252 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 3)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1253 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 4)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1254 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 5)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1255 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 6)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1256 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 7)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1257 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 8)
	} else if this.object.M_textEncoding == Config.Encoding_WINDOWS_1258 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 9)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_1 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 10)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_2 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 11)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_3 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 12)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_4 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 13)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_5 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 14)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_6 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 15)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_7 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 16)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_8 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 17)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_9 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 18)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_10 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 19)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_13 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 20)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_14 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 21)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_15 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 22)
	} else if this.object.M_textEncoding == Config.Encoding_ISO8859_16 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 23)
	} else if this.object.M_textEncoding == Config.Encoding_KOI8R {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 24)
	} else if this.object.M_textEncoding == Config.Encoding_KOI8U {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 25)
	} else {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 0)
	}
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_storeName)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_hostName)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_ipv4)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_user)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_pwd)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_port)

	/** associations of DataStoreConfiguration **/

	/** Save parent type Configurations **/
	if len(this.object.M_parentPtr) > 0 {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_parentPtr)
	} else {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), DataStoreConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), DataStoreConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_DataStoreConfigurationEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_DataStoreConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.DataStoreConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of DataStoreConfiguration **/
	query.Fields = append(query.Fields, "M_dataStoreType")
	query.Fields = append(query.Fields, "M_dataStoreVendor")
	query.Fields = append(query.Fields, "M_textEncoding")
	query.Fields = append(query.Fields, "M_storeName")
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")
	query.Fields = append(query.Fields, "M_port")

	/** associations of DataStoreConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of DataStoreConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.DataStoreConfiguration)
		this.object.TYPENAME = "Config.DataStoreConfiguration"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of DataStoreConfiguration **/

		/** dataStoreType **/
		if results[0][4] != nil {
			enumIndex := results[0][4].(int)
			if enumIndex == 0 {
				this.object.M_dataStoreType = Config.DataStoreType_SQL_STORE
			} else if enumIndex == 1 {
				this.object.M_dataStoreType = Config.DataStoreType_KEY_VALUE_STORE
			}
		}

		/** dataStoreVendor **/
		if results[0][5] != nil {
			enumIndex := results[0][5].(int)
			if enumIndex == 0 {
				this.object.M_dataStoreVendor = Config.DataStoreVendor_MYCELIUS
			} else if enumIndex == 1 {
				this.object.M_dataStoreVendor = Config.DataStoreVendor_MYSQL
			} else if enumIndex == 2 {
				this.object.M_dataStoreVendor = Config.DataStoreVendor_MSSQL
			} else if enumIndex == 3 {
				this.object.M_dataStoreVendor = Config.DataStoreVendor_ODBC
			} else if enumIndex == 4 {
				this.object.M_dataStoreVendor = Config.DataStoreVendor_KNOWLEDGEBASE
			}
		}

		/** textEncoding **/
		if results[0][6] != nil {
			enumIndex := results[0][6].(int)
			if enumIndex == 0 {
				this.object.M_textEncoding = Config.Encoding_UTF8
			} else if enumIndex == 1 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1250
			} else if enumIndex == 2 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1251
			} else if enumIndex == 3 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1252
			} else if enumIndex == 4 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1253
			} else if enumIndex == 5 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1254
			} else if enumIndex == 6 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1255
			} else if enumIndex == 7 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1256
			} else if enumIndex == 8 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1257
			} else if enumIndex == 9 {
				this.object.M_textEncoding = Config.Encoding_WINDOWS_1258
			} else if enumIndex == 10 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_1
			} else if enumIndex == 11 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_2
			} else if enumIndex == 12 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_3
			} else if enumIndex == 13 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_4
			} else if enumIndex == 14 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_5
			} else if enumIndex == 15 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_6
			} else if enumIndex == 16 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_7
			} else if enumIndex == 17 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_8
			} else if enumIndex == 18 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_9
			} else if enumIndex == 19 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_10
			} else if enumIndex == 20 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_13
			} else if enumIndex == 21 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_14
			} else if enumIndex == 22 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_15
			} else if enumIndex == 23 {
				this.object.M_textEncoding = Config.Encoding_ISO8859_16
			} else if enumIndex == 24 {
				this.object.M_textEncoding = Config.Encoding_KOI8R
			} else if enumIndex == 25 {
				this.object.M_textEncoding = Config.Encoding_KOI8U
			}
		}

		/** storeName **/
		if results[0][7] != nil {
			this.object.M_storeName = results[0][7].(string)
		}

		/** hostName **/
		if results[0][8] != nil {
			this.object.M_hostName = results[0][8].(string)
		}

		/** ipv4 **/
		if results[0][9] != nil {
			this.object.M_ipv4 = results[0][9].(string)
		}

		/** user **/
		if results[0][10] != nil {
			this.object.M_user = results[0][10].(string)
		}

		/** pwd **/
		if results[0][11] != nil {
			this.object.M_pwd = results[0][11].(string)
		}

		/** port **/
		if results[0][12] != nil {
			this.object.M_port = results[0][12].(int)
		}

		/** associations of DataStoreConfiguration **/

		/** parentPtr **/
		if results[0][13] != nil {
			id := results[0][13].(string)
			if len(id) > 0 {
				refTypeName := "Config.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigDataStoreConfigurationEntityFromObject(object *Config.DataStoreConfiguration) *Config_DataStoreConfigurationEntity {
	return this.NewConfigDataStoreConfigurationEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_DataStoreConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigDataStoreConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.DataStoreConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_DataStoreConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_DataStoreConfigurationEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			LdapConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_LdapConfigurationEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.LdapConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewConfigLdapConfigurationEntity(parentUuid string, objectId string, object interface{}) *Config_LdapConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigLdapConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.LdapConfiguration).TYPENAME = "Config.LdapConfiguration"
		object.(*Config.LdapConfiguration).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.LdapConfiguration", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.LdapConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.LdapConfiguration).UUID
			}
			return val.(*Config_LdapConfigurationEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_LdapConfigurationEntity)
	if object == nil {
		entity.object = new(Config.LdapConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.LdapConfiguration)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.LdapConfiguration"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_LdapConfigurationEntity) GetTypeName() string {
	return "Config.LdapConfiguration"
}
func (this *Config_LdapConfigurationEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_LdapConfigurationEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_LdapConfigurationEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_LdapConfigurationEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_LdapConfigurationEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_LdapConfigurationEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_LdapConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_LdapConfigurationEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_LdapConfigurationEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_LdapConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_LdapConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_LdapConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_LdapConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_LdapConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_LdapConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_LdapConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_LdapConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_LdapConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_LdapConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_LdapConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_LdapConfigurationEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_LdapConfigurationEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_LdapConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.LdapConfiguration"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_LdapConfigurationEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_LdapConfigurationEntityPrototype() {

	var ldapConfigurationEntityProto EntityPrototype
	ldapConfigurationEntityProto.TypeName = "Config.LdapConfiguration"
	ldapConfigurationEntityProto.SuperTypeNames = append(ldapConfigurationEntityProto.SuperTypeNames, "Config.Configuration")
	ldapConfigurationEntityProto.Ids = append(ldapConfigurationEntityProto.Ids, "UUID")
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "UUID")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 0)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, false)
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "")
	ldapConfigurationEntityProto.Indexs = append(ldapConfigurationEntityProto.Indexs, "ParentUuid")
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "ParentUuid")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 1)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, false)
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "")
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "ParentLnk")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 2)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, false)
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of Configuration **/
	ldapConfigurationEntityProto.Ids = append(ldapConfigurationEntityProto.Ids, "M_id")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 3)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_id")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.ID")
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of LdapConfiguration **/
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 4)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_hostName")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 5)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_ipv4")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 6)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_port")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.int")
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "0")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 7)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_user")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 8)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_pwd")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 9)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_domain")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 10)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_searchBase")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "")

	/** associations of LdapConfiguration **/
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 11)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, false)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_parentPtr")
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "undefined")
	ldapConfigurationEntityProto.FieldsDefaultValue = append(ldapConfigurationEntityProto.FieldsDefaultValue, "undefined")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "Config.Configurations:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&ldapConfigurationEntityProto)

}

/** Create **/
func (this *Config_LdapConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.LdapConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of LdapConfiguration **/
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_port")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")
	query.Fields = append(query.Fields, "M_domain")
	query.Fields = append(query.Fields, "M_searchBase")

	/** associations of LdapConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var LdapConfigurationInfo []interface{}

	LdapConfigurationInfo = append(LdapConfigurationInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		LdapConfigurationInfo = append(LdapConfigurationInfo, this.GetParentPtr().GetUuid())
		LdapConfigurationInfo = append(LdapConfigurationInfo, this.GetParentLnk())
	} else {
		LdapConfigurationInfo = append(LdapConfigurationInfo, "")
		LdapConfigurationInfo = append(LdapConfigurationInfo, "")
	}

	/** members of Configuration **/
	LdapConfigurationInfo = append(LdapConfigurationInfo, this.object.M_id)

	/** members of LdapConfiguration **/
	LdapConfigurationInfo = append(LdapConfigurationInfo, this.object.M_hostName)
	LdapConfigurationInfo = append(LdapConfigurationInfo, this.object.M_ipv4)
	LdapConfigurationInfo = append(LdapConfigurationInfo, this.object.M_port)
	LdapConfigurationInfo = append(LdapConfigurationInfo, this.object.M_user)
	LdapConfigurationInfo = append(LdapConfigurationInfo, this.object.M_pwd)
	LdapConfigurationInfo = append(LdapConfigurationInfo, this.object.M_domain)
	LdapConfigurationInfo = append(LdapConfigurationInfo, this.object.M_searchBase)

	/** associations of LdapConfiguration **/

	/** Save parent type Configurations **/
	if len(this.object.M_parentPtr) > 0 {
		LdapConfigurationInfo = append(LdapConfigurationInfo, this.object.M_parentPtr)
	} else {
		LdapConfigurationInfo = append(LdapConfigurationInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), LdapConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), LdapConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_LdapConfigurationEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_LdapConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.LdapConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of LdapConfiguration **/
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_port")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")
	query.Fields = append(query.Fields, "M_domain")
	query.Fields = append(query.Fields, "M_searchBase")

	/** associations of LdapConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of LdapConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.LdapConfiguration)
		this.object.TYPENAME = "Config.LdapConfiguration"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of LdapConfiguration **/

		/** hostName **/
		if results[0][4] != nil {
			this.object.M_hostName = results[0][4].(string)
		}

		/** ipv4 **/
		if results[0][5] != nil {
			this.object.M_ipv4 = results[0][5].(string)
		}

		/** port **/
		if results[0][6] != nil {
			this.object.M_port = results[0][6].(int)
		}

		/** user **/
		if results[0][7] != nil {
			this.object.M_user = results[0][7].(string)
		}

		/** pwd **/
		if results[0][8] != nil {
			this.object.M_pwd = results[0][8].(string)
		}

		/** domain **/
		if results[0][9] != nil {
			this.object.M_domain = results[0][9].(string)
		}

		/** searchBase **/
		if results[0][10] != nil {
			this.object.M_searchBase = results[0][10].(string)
		}

		/** associations of LdapConfiguration **/

		/** parentPtr **/
		if results[0][11] != nil {
			id := results[0][11].(string)
			if len(id) > 0 {
				refTypeName := "Config.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigLdapConfigurationEntityFromObject(object *Config.LdapConfiguration) *Config_LdapConfigurationEntity {
	return this.NewConfigLdapConfigurationEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_LdapConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigLdapConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.LdapConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_LdapConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_LdapConfigurationEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			OAuth2Client
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_OAuth2ClientEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.OAuth2Client
}

/** Constructor function **/
func (this *EntityManager) NewConfigOAuth2ClientEntity(parentUuid string, objectId string, object interface{}) *Config_OAuth2ClientEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigOAuth2ClientExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.OAuth2Client).TYPENAME = "Config.OAuth2Client"
		object.(*Config.OAuth2Client).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.OAuth2Client", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.OAuth2Client).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.OAuth2Client).UUID
			}
			return val.(*Config_OAuth2ClientEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_OAuth2ClientEntity)
	if object == nil {
		entity.object = new(Config.OAuth2Client)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.OAuth2Client)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.OAuth2Client"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_OAuth2ClientEntity) GetTypeName() string {
	return "Config.OAuth2Client"
}
func (this *Config_OAuth2ClientEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_OAuth2ClientEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_OAuth2ClientEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_OAuth2ClientEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_OAuth2ClientEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_OAuth2ClientEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_OAuth2ClientEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_OAuth2ClientEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_OAuth2ClientEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_OAuth2ClientEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_OAuth2ClientEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_OAuth2ClientEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_OAuth2ClientEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_OAuth2ClientEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_OAuth2ClientEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_OAuth2ClientEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_OAuth2ClientEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_OAuth2ClientEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_OAuth2ClientEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_OAuth2ClientEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_OAuth2ClientEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_OAuth2ClientEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_OAuth2ClientEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Client"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_OAuth2ClientEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_OAuth2ClientEntityPrototype() {

	var oAuth2ClientEntityProto EntityPrototype
	oAuth2ClientEntityProto.TypeName = "Config.OAuth2Client"
	oAuth2ClientEntityProto.Ids = append(oAuth2ClientEntityProto.Ids, "UUID")
	oAuth2ClientEntityProto.Fields = append(oAuth2ClientEntityProto.Fields, "UUID")
	oAuth2ClientEntityProto.FieldsType = append(oAuth2ClientEntityProto.FieldsType, "xs.string")
	oAuth2ClientEntityProto.FieldsOrder = append(oAuth2ClientEntityProto.FieldsOrder, 0)
	oAuth2ClientEntityProto.FieldsVisibility = append(oAuth2ClientEntityProto.FieldsVisibility, false)
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "")
	oAuth2ClientEntityProto.Indexs = append(oAuth2ClientEntityProto.Indexs, "ParentUuid")
	oAuth2ClientEntityProto.Fields = append(oAuth2ClientEntityProto.Fields, "ParentUuid")
	oAuth2ClientEntityProto.FieldsType = append(oAuth2ClientEntityProto.FieldsType, "xs.string")
	oAuth2ClientEntityProto.FieldsOrder = append(oAuth2ClientEntityProto.FieldsOrder, 1)
	oAuth2ClientEntityProto.FieldsVisibility = append(oAuth2ClientEntityProto.FieldsVisibility, false)
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "")
	oAuth2ClientEntityProto.Fields = append(oAuth2ClientEntityProto.Fields, "ParentLnk")
	oAuth2ClientEntityProto.FieldsType = append(oAuth2ClientEntityProto.FieldsType, "xs.string")
	oAuth2ClientEntityProto.FieldsOrder = append(oAuth2ClientEntityProto.FieldsOrder, 2)
	oAuth2ClientEntityProto.FieldsVisibility = append(oAuth2ClientEntityProto.FieldsVisibility, false)
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "")

	/** members of OAuth2Client **/
	oAuth2ClientEntityProto.Ids = append(oAuth2ClientEntityProto.Ids, "M_id")
	oAuth2ClientEntityProto.FieldsOrder = append(oAuth2ClientEntityProto.FieldsOrder, 3)
	oAuth2ClientEntityProto.FieldsVisibility = append(oAuth2ClientEntityProto.FieldsVisibility, true)
	oAuth2ClientEntityProto.Fields = append(oAuth2ClientEntityProto.Fields, "M_id")
	oAuth2ClientEntityProto.FieldsType = append(oAuth2ClientEntityProto.FieldsType, "xs.ID")
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "")
	oAuth2ClientEntityProto.FieldsOrder = append(oAuth2ClientEntityProto.FieldsOrder, 4)
	oAuth2ClientEntityProto.FieldsVisibility = append(oAuth2ClientEntityProto.FieldsVisibility, true)
	oAuth2ClientEntityProto.Fields = append(oAuth2ClientEntityProto.Fields, "M_secret")
	oAuth2ClientEntityProto.FieldsType = append(oAuth2ClientEntityProto.FieldsType, "xs.string")
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "")
	oAuth2ClientEntityProto.FieldsOrder = append(oAuth2ClientEntityProto.FieldsOrder, 5)
	oAuth2ClientEntityProto.FieldsVisibility = append(oAuth2ClientEntityProto.FieldsVisibility, true)
	oAuth2ClientEntityProto.Fields = append(oAuth2ClientEntityProto.Fields, "M_redirectUri")
	oAuth2ClientEntityProto.FieldsType = append(oAuth2ClientEntityProto.FieldsType, "xs.string")
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "")
	oAuth2ClientEntityProto.FieldsOrder = append(oAuth2ClientEntityProto.FieldsOrder, 6)
	oAuth2ClientEntityProto.FieldsVisibility = append(oAuth2ClientEntityProto.FieldsVisibility, true)
	oAuth2ClientEntityProto.Fields = append(oAuth2ClientEntityProto.Fields, "M_tokenUri")
	oAuth2ClientEntityProto.FieldsType = append(oAuth2ClientEntityProto.FieldsType, "xs.string")
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "")
	oAuth2ClientEntityProto.FieldsOrder = append(oAuth2ClientEntityProto.FieldsOrder, 7)
	oAuth2ClientEntityProto.FieldsVisibility = append(oAuth2ClientEntityProto.FieldsVisibility, true)
	oAuth2ClientEntityProto.Fields = append(oAuth2ClientEntityProto.Fields, "M_authorizationUri")
	oAuth2ClientEntityProto.FieldsType = append(oAuth2ClientEntityProto.FieldsType, "xs.string")
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "")
	oAuth2ClientEntityProto.FieldsOrder = append(oAuth2ClientEntityProto.FieldsOrder, 8)
	oAuth2ClientEntityProto.FieldsVisibility = append(oAuth2ClientEntityProto.FieldsVisibility, true)
	oAuth2ClientEntityProto.Fields = append(oAuth2ClientEntityProto.Fields, "M_extra")
	oAuth2ClientEntityProto.FieldsType = append(oAuth2ClientEntityProto.FieldsType, "xs.[]uint8")
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "undefined")

	/** associations of OAuth2Client **/
	oAuth2ClientEntityProto.FieldsOrder = append(oAuth2ClientEntityProto.FieldsOrder, 9)
	oAuth2ClientEntityProto.FieldsVisibility = append(oAuth2ClientEntityProto.FieldsVisibility, false)
	oAuth2ClientEntityProto.Fields = append(oAuth2ClientEntityProto.Fields, "M_parentPtr")
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "undefined")
	oAuth2ClientEntityProto.FieldsDefaultValue = append(oAuth2ClientEntityProto.FieldsDefaultValue, "undefined")
	oAuth2ClientEntityProto.FieldsType = append(oAuth2ClientEntityProto.FieldsType, "Config.OAuth2Configuration:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&oAuth2ClientEntityProto)

}

/** Create **/
func (this *Config_OAuth2ClientEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.OAuth2Client"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2Client **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_secret")
	query.Fields = append(query.Fields, "M_redirectUri")
	query.Fields = append(query.Fields, "M_tokenUri")
	query.Fields = append(query.Fields, "M_authorizationUri")
	query.Fields = append(query.Fields, "M_extra")

	/** associations of OAuth2Client **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var OAuth2ClientInfo []interface{}

	OAuth2ClientInfo = append(OAuth2ClientInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		OAuth2ClientInfo = append(OAuth2ClientInfo, this.GetParentPtr().GetUuid())
		OAuth2ClientInfo = append(OAuth2ClientInfo, this.GetParentLnk())
	} else {
		OAuth2ClientInfo = append(OAuth2ClientInfo, "")
		OAuth2ClientInfo = append(OAuth2ClientInfo, "")
	}

	/** members of OAuth2Client **/
	OAuth2ClientInfo = append(OAuth2ClientInfo, this.object.M_id)
	OAuth2ClientInfo = append(OAuth2ClientInfo, this.object.M_secret)
	OAuth2ClientInfo = append(OAuth2ClientInfo, this.object.M_redirectUri)
	OAuth2ClientInfo = append(OAuth2ClientInfo, this.object.M_tokenUri)
	OAuth2ClientInfo = append(OAuth2ClientInfo, this.object.M_authorizationUri)
	OAuth2ClientInfo = append(OAuth2ClientInfo, this.object.M_extra)

	/** associations of OAuth2Client **/

	/** Save parent type OAuth2Configuration **/
	if len(this.object.M_parentPtr) > 0 {
		OAuth2ClientInfo = append(OAuth2ClientInfo, this.object.M_parentPtr)
	} else {
		OAuth2ClientInfo = append(OAuth2ClientInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), OAuth2ClientInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), OAuth2ClientInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_OAuth2ClientEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_OAuth2ClientEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.OAuth2Client"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2Client **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_secret")
	query.Fields = append(query.Fields, "M_redirectUri")
	query.Fields = append(query.Fields, "M_tokenUri")
	query.Fields = append(query.Fields, "M_authorizationUri")
	query.Fields = append(query.Fields, "M_extra")

	/** associations of OAuth2Client **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of OAuth2Client...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.OAuth2Client)
		this.object.TYPENAME = "Config.OAuth2Client"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of OAuth2Client **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** secret **/
		if results[0][4] != nil {
			this.object.M_secret = results[0][4].(string)
		}

		/** redirectUri **/
		if results[0][5] != nil {
			this.object.M_redirectUri = results[0][5].(string)
		}

		/** tokenUri **/
		if results[0][6] != nil {
			this.object.M_tokenUri = results[0][6].(string)
		}

		/** authorizationUri **/
		if results[0][7] != nil {
			this.object.M_authorizationUri = results[0][7].(string)
		}

		/** extra **/
		if results[0][8] != nil {
			this.object.M_extra = results[0][8].([]uint8)
		}

		/** associations of OAuth2Client **/

		/** parentPtr **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2Configuration"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigOAuth2ClientEntityFromObject(object *Config.OAuth2Client) *Config_OAuth2ClientEntity {
	return this.NewConfigOAuth2ClientEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_OAuth2ClientEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigOAuth2ClientExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Client"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_OAuth2ClientEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_OAuth2ClientEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			OAuth2Authorize
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_OAuth2AuthorizeEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.OAuth2Authorize
}

/** Constructor function **/
func (this *EntityManager) NewConfigOAuth2AuthorizeEntity(parentUuid string, objectId string, object interface{}) *Config_OAuth2AuthorizeEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigOAuth2AuthorizeExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.OAuth2Authorize).TYPENAME = "Config.OAuth2Authorize"
		object.(*Config.OAuth2Authorize).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.OAuth2Authorize", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.OAuth2Authorize).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.OAuth2Authorize).UUID
			}
			return val.(*Config_OAuth2AuthorizeEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_OAuth2AuthorizeEntity)
	if object == nil {
		entity.object = new(Config.OAuth2Authorize)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.OAuth2Authorize)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.OAuth2Authorize"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_OAuth2AuthorizeEntity) GetTypeName() string {
	return "Config.OAuth2Authorize"
}
func (this *Config_OAuth2AuthorizeEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_OAuth2AuthorizeEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_OAuth2AuthorizeEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_OAuth2AuthorizeEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_OAuth2AuthorizeEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_OAuth2AuthorizeEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_OAuth2AuthorizeEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_OAuth2AuthorizeEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_OAuth2AuthorizeEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_OAuth2AuthorizeEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_OAuth2AuthorizeEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_OAuth2AuthorizeEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_OAuth2AuthorizeEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_OAuth2AuthorizeEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_OAuth2AuthorizeEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_OAuth2AuthorizeEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_OAuth2AuthorizeEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_OAuth2AuthorizeEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_OAuth2AuthorizeEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_OAuth2AuthorizeEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_OAuth2AuthorizeEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_OAuth2AuthorizeEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_OAuth2AuthorizeEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Authorize"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_OAuth2AuthorizeEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_OAuth2AuthorizeEntityPrototype() {

	var oAuth2AuthorizeEntityProto EntityPrototype
	oAuth2AuthorizeEntityProto.TypeName = "Config.OAuth2Authorize"
	oAuth2AuthorizeEntityProto.Ids = append(oAuth2AuthorizeEntityProto.Ids, "UUID")
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "UUID")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "xs.string")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 0)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, false)
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "")
	oAuth2AuthorizeEntityProto.Indexs = append(oAuth2AuthorizeEntityProto.Indexs, "ParentUuid")
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "ParentUuid")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "xs.string")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 1)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, false)
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "")
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "ParentLnk")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "xs.string")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 2)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, false)
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "")

	/** members of OAuth2Authorize **/
	oAuth2AuthorizeEntityProto.Ids = append(oAuth2AuthorizeEntityProto.Ids, "M_id")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 3)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, true)
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "M_id")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "xs.ID")
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 4)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, true)
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "M_client")
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "Config.OAuth2Client:Ref")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 5)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, true)
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "M_expiresIn")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "xs.time")
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "0")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 6)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, true)
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "M_scope")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "xs.string")
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 7)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, true)
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "M_redirectUri")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "xs.string")
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 8)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, true)
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "M_state")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "xs.string")
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 9)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, true)
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "M_userData")
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "Config.OAuth2IdToken:Ref")
	oAuth2AuthorizeEntityProto.FieldsOrder = append(oAuth2AuthorizeEntityProto.FieldsOrder, 10)
	oAuth2AuthorizeEntityProto.FieldsVisibility = append(oAuth2AuthorizeEntityProto.FieldsVisibility, true)
	oAuth2AuthorizeEntityProto.Fields = append(oAuth2AuthorizeEntityProto.Fields, "M_createdAt")
	oAuth2AuthorizeEntityProto.FieldsType = append(oAuth2AuthorizeEntityProto.FieldsType, "xs.date")
	oAuth2AuthorizeEntityProto.FieldsDefaultValue = append(oAuth2AuthorizeEntityProto.FieldsDefaultValue, "new Date()")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&oAuth2AuthorizeEntityProto)

}

/** Create **/
func (this *Config_OAuth2AuthorizeEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.OAuth2Authorize"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2Authorize **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_client")
	query.Fields = append(query.Fields, "M_expiresIn")
	query.Fields = append(query.Fields, "M_scope")
	query.Fields = append(query.Fields, "M_redirectUri")
	query.Fields = append(query.Fields, "M_state")
	query.Fields = append(query.Fields, "M_userData")
	query.Fields = append(query.Fields, "M_createdAt")

	var OAuth2AuthorizeInfo []interface{}

	OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.GetParentPtr().GetUuid())
		OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.GetParentLnk())
	} else {
		OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, "")
		OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, "")
	}

	/** members of OAuth2Authorize **/
	OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.object.M_id)

	/** Save client type OAuth2Client **/
	if len(this.object.M_client) > 0 {
		OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.object.M_client)
	} else {
		OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, "")
	}
	OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.object.M_expiresIn)
	OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.object.M_scope)
	OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.object.M_redirectUri)
	OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.object.M_state)

	/** Save userData type OAuth2IdToken **/
	if len(this.object.M_userData) > 0 {
		OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.object.M_userData)
	} else {
		OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, "")
	}
	OAuth2AuthorizeInfo = append(OAuth2AuthorizeInfo, this.object.M_createdAt)
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), OAuth2AuthorizeInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), OAuth2AuthorizeInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_OAuth2AuthorizeEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_OAuth2AuthorizeEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.OAuth2Authorize"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2Authorize **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_client")
	query.Fields = append(query.Fields, "M_expiresIn")
	query.Fields = append(query.Fields, "M_scope")
	query.Fields = append(query.Fields, "M_redirectUri")
	query.Fields = append(query.Fields, "M_state")
	query.Fields = append(query.Fields, "M_userData")
	query.Fields = append(query.Fields, "M_createdAt")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of OAuth2Authorize...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.OAuth2Authorize)
		this.object.TYPENAME = "Config.OAuth2Authorize"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of OAuth2Authorize **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** client **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2Client"
				id_ := refTypeName + "$$" + id
				this.object.M_client = id
				GetServer().GetEntityManager().appendReference("client", this.object.UUID, id_)
			}
		}

		/** expiresIn **/
		if results[0][5] != nil {
			this.object.M_expiresIn = results[0][5].(int64)
		}

		/** scope **/
		if results[0][6] != nil {
			this.object.M_scope = results[0][6].(string)
		}

		/** redirectUri **/
		if results[0][7] != nil {
			this.object.M_redirectUri = results[0][7].(string)
		}

		/** state **/
		if results[0][8] != nil {
			this.object.M_state = results[0][8].(string)
		}

		/** userData **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2IdToken"
				id_ := refTypeName + "$$" + id
				this.object.M_userData = id
				GetServer().GetEntityManager().appendReference("userData", this.object.UUID, id_)
			}
		}

		/** createdAt **/
		if results[0][10] != nil {
			this.object.M_createdAt = results[0][10].(int64)
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigOAuth2AuthorizeEntityFromObject(object *Config.OAuth2Authorize) *Config_OAuth2AuthorizeEntity {
	return this.NewConfigOAuth2AuthorizeEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_OAuth2AuthorizeEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigOAuth2AuthorizeExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Authorize"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_OAuth2AuthorizeEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_OAuth2AuthorizeEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			OAuth2IdToken
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_OAuth2IdTokenEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.OAuth2IdToken
}

/** Constructor function **/
func (this *EntityManager) NewConfigOAuth2IdTokenEntity(parentUuid string, objectId string, object interface{}) *Config_OAuth2IdTokenEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigOAuth2IdTokenExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.OAuth2IdToken).TYPENAME = "Config.OAuth2IdToken"
		object.(*Config.OAuth2IdToken).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.OAuth2IdToken", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.OAuth2IdToken).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.OAuth2IdToken).UUID
			}
			return val.(*Config_OAuth2IdTokenEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_OAuth2IdTokenEntity)
	if object == nil {
		entity.object = new(Config.OAuth2IdToken)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.OAuth2IdToken)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.OAuth2IdToken"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_OAuth2IdTokenEntity) GetTypeName() string {
	return "Config.OAuth2IdToken"
}
func (this *Config_OAuth2IdTokenEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_OAuth2IdTokenEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_OAuth2IdTokenEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_OAuth2IdTokenEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_OAuth2IdTokenEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_OAuth2IdTokenEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_OAuth2IdTokenEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_OAuth2IdTokenEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_OAuth2IdTokenEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_OAuth2IdTokenEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_OAuth2IdTokenEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_OAuth2IdTokenEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_OAuth2IdTokenEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_OAuth2IdTokenEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_OAuth2IdTokenEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_OAuth2IdTokenEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_OAuth2IdTokenEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_OAuth2IdTokenEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_OAuth2IdTokenEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_OAuth2IdTokenEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_OAuth2IdTokenEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_OAuth2IdTokenEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_OAuth2IdTokenEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.OAuth2IdToken"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_OAuth2IdTokenEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_OAuth2IdTokenEntityPrototype() {

	var oAuth2IdTokenEntityProto EntityPrototype
	oAuth2IdTokenEntityProto.TypeName = "Config.OAuth2IdToken"
	oAuth2IdTokenEntityProto.Ids = append(oAuth2IdTokenEntityProto.Ids, "UUID")
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "UUID")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.string")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 0)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, false)
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")
	oAuth2IdTokenEntityProto.Indexs = append(oAuth2IdTokenEntityProto.Indexs, "ParentUuid")
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "ParentUuid")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.string")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 1)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, false)
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "ParentLnk")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.string")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 2)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, false)
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")

	/** members of OAuth2IdToken **/
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 3)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_issuer")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.string")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")
	oAuth2IdTokenEntityProto.Ids = append(oAuth2IdTokenEntityProto.Ids, "M_id")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 4)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_id")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.ID")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 5)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_client")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "undefined")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "undefined")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "Config.OAuth2Client:Ref")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 6)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_expiration")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.date")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "new Date()")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 7)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_issuedAt")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.date")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "new Date()")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 8)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_nonce")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.string")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 9)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_email")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.string")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 10)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_emailVerified")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.boolean")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "false")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 11)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_name")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.string")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 12)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_familyName")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.string")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 13)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_givenName")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.string")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 14)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, true)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_local")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "xs.string")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "")

	/** associations of OAuth2IdToken **/
	oAuth2IdTokenEntityProto.FieldsOrder = append(oAuth2IdTokenEntityProto.FieldsOrder, 15)
	oAuth2IdTokenEntityProto.FieldsVisibility = append(oAuth2IdTokenEntityProto.FieldsVisibility, false)
	oAuth2IdTokenEntityProto.Fields = append(oAuth2IdTokenEntityProto.Fields, "M_parentPtr")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "undefined")
	oAuth2IdTokenEntityProto.FieldsDefaultValue = append(oAuth2IdTokenEntityProto.FieldsDefaultValue, "undefined")
	oAuth2IdTokenEntityProto.FieldsType = append(oAuth2IdTokenEntityProto.FieldsType, "Config.OAuth2Configuration:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&oAuth2IdTokenEntityProto)

}

/** Create **/
func (this *Config_OAuth2IdTokenEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.OAuth2IdToken"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2IdToken **/
	query.Fields = append(query.Fields, "M_issuer")
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_client")
	query.Fields = append(query.Fields, "M_expiration")
	query.Fields = append(query.Fields, "M_issuedAt")
	query.Fields = append(query.Fields, "M_nonce")
	query.Fields = append(query.Fields, "M_email")
	query.Fields = append(query.Fields, "M_emailVerified")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_familyName")
	query.Fields = append(query.Fields, "M_givenName")
	query.Fields = append(query.Fields, "M_local")

	/** associations of OAuth2IdToken **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var OAuth2IdTokenInfo []interface{}

	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.GetParentPtr().GetUuid())
		OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.GetParentLnk())
	} else {
		OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, "")
		OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, "")
	}

	/** members of OAuth2IdToken **/
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_issuer)
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_id)

	/** Save client type OAuth2Client **/
	if len(this.object.M_client) > 0 {
		OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_client)
	} else {
		OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, "")
	}
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_expiration)
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_issuedAt)
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_nonce)
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_email)
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_emailVerified)
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_name)
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_familyName)
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_givenName)
	OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_local)

	/** associations of OAuth2IdToken **/

	/** Save parent type OAuth2Configuration **/
	if len(this.object.M_parentPtr) > 0 {
		OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, this.object.M_parentPtr)
	} else {
		OAuth2IdTokenInfo = append(OAuth2IdTokenInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), OAuth2IdTokenInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), OAuth2IdTokenInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_OAuth2IdTokenEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_OAuth2IdTokenEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.OAuth2IdToken"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2IdToken **/
	query.Fields = append(query.Fields, "M_issuer")
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_client")
	query.Fields = append(query.Fields, "M_expiration")
	query.Fields = append(query.Fields, "M_issuedAt")
	query.Fields = append(query.Fields, "M_nonce")
	query.Fields = append(query.Fields, "M_email")
	query.Fields = append(query.Fields, "M_emailVerified")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_familyName")
	query.Fields = append(query.Fields, "M_givenName")
	query.Fields = append(query.Fields, "M_local")

	/** associations of OAuth2IdToken **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of OAuth2IdToken...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.OAuth2IdToken)
		this.object.TYPENAME = "Config.OAuth2IdToken"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of OAuth2IdToken **/

		/** issuer **/
		if results[0][3] != nil {
			this.object.M_issuer = results[0][3].(string)
		}

		/** id **/
		if results[0][4] != nil {
			this.object.M_id = results[0][4].(string)
		}

		/** client **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2Client"
				id_ := refTypeName + "$$" + id
				this.object.M_client = id
				GetServer().GetEntityManager().appendReference("client", this.object.UUID, id_)
			}
		}

		/** expiration **/
		if results[0][6] != nil {
			this.object.M_expiration = results[0][6].(int64)
		}

		/** issuedAt **/
		if results[0][7] != nil {
			this.object.M_issuedAt = results[0][7].(int64)
		}

		/** nonce **/
		if results[0][8] != nil {
			this.object.M_nonce = results[0][8].(string)
		}

		/** email **/
		if results[0][9] != nil {
			this.object.M_email = results[0][9].(string)
		}

		/** emailVerified **/
		if results[0][10] != nil {
			this.object.M_emailVerified = results[0][10].(bool)
		}

		/** name **/
		if results[0][11] != nil {
			this.object.M_name = results[0][11].(string)
		}

		/** familyName **/
		if results[0][12] != nil {
			this.object.M_familyName = results[0][12].(string)
		}

		/** givenName **/
		if results[0][13] != nil {
			this.object.M_givenName = results[0][13].(string)
		}

		/** local **/
		if results[0][14] != nil {
			this.object.M_local = results[0][14].(string)
		}

		/** associations of OAuth2IdToken **/

		/** parentPtr **/
		if results[0][15] != nil {
			id := results[0][15].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2Configuration"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigOAuth2IdTokenEntityFromObject(object *Config.OAuth2IdToken) *Config_OAuth2IdTokenEntity {
	return this.NewConfigOAuth2IdTokenEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_OAuth2IdTokenEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigOAuth2IdTokenExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.OAuth2IdToken"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_OAuth2IdTokenEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_OAuth2IdTokenEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			OAuth2Access
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_OAuth2AccessEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.OAuth2Access
}

/** Constructor function **/
func (this *EntityManager) NewConfigOAuth2AccessEntity(parentUuid string, objectId string, object interface{}) *Config_OAuth2AccessEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigOAuth2AccessExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.OAuth2Access).TYPENAME = "Config.OAuth2Access"
		object.(*Config.OAuth2Access).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.OAuth2Access", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.OAuth2Access).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.OAuth2Access).UUID
			}
			return val.(*Config_OAuth2AccessEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_OAuth2AccessEntity)
	if object == nil {
		entity.object = new(Config.OAuth2Access)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.OAuth2Access)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.OAuth2Access"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_OAuth2AccessEntity) GetTypeName() string {
	return "Config.OAuth2Access"
}
func (this *Config_OAuth2AccessEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_OAuth2AccessEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_OAuth2AccessEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_OAuth2AccessEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_OAuth2AccessEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_OAuth2AccessEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_OAuth2AccessEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_OAuth2AccessEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_OAuth2AccessEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_OAuth2AccessEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_OAuth2AccessEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_OAuth2AccessEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_OAuth2AccessEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_OAuth2AccessEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_OAuth2AccessEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_OAuth2AccessEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_OAuth2AccessEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_OAuth2AccessEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_OAuth2AccessEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_OAuth2AccessEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_OAuth2AccessEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_OAuth2AccessEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_OAuth2AccessEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Access"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_OAuth2AccessEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_OAuth2AccessEntityPrototype() {

	var oAuth2AccessEntityProto EntityPrototype
	oAuth2AccessEntityProto.TypeName = "Config.OAuth2Access"
	oAuth2AccessEntityProto.Ids = append(oAuth2AccessEntityProto.Ids, "UUID")
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "UUID")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "xs.string")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 0)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, false)
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "")
	oAuth2AccessEntityProto.Indexs = append(oAuth2AccessEntityProto.Indexs, "ParentUuid")
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "ParentUuid")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "xs.string")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 1)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, false)
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "")
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "ParentLnk")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "xs.string")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 2)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, false)
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "")

	/** members of OAuth2Access **/
	oAuth2AccessEntityProto.Ids = append(oAuth2AccessEntityProto.Ids, "M_id")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 3)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, true)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_id")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "xs.ID")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 4)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, true)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_client")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "Config.OAuth2Client:Ref")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 5)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, true)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_authorize")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "xs.string")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 6)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, true)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_previous")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "xs.string")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 7)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, true)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_refreshToken")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "Config.OAuth2Refresh:Ref")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 8)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, true)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_expiresIn")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "xs.time")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "0")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 9)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, true)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_scope")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "xs.string")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 10)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, true)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_redirectUri")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "xs.string")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 11)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, true)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_userData")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "Config.OAuth2IdToken:Ref")
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 12)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, true)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_createdAt")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "xs.date")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "new Date()")

	/** associations of OAuth2Access **/
	oAuth2AccessEntityProto.FieldsOrder = append(oAuth2AccessEntityProto.FieldsOrder, 13)
	oAuth2AccessEntityProto.FieldsVisibility = append(oAuth2AccessEntityProto.FieldsVisibility, false)
	oAuth2AccessEntityProto.Fields = append(oAuth2AccessEntityProto.Fields, "M_parentPtr")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AccessEntityProto.FieldsDefaultValue = append(oAuth2AccessEntityProto.FieldsDefaultValue, "undefined")
	oAuth2AccessEntityProto.FieldsType = append(oAuth2AccessEntityProto.FieldsType, "Config.OAuth2Configuration:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&oAuth2AccessEntityProto)

}

/** Create **/
func (this *Config_OAuth2AccessEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.OAuth2Access"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2Access **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_client")
	query.Fields = append(query.Fields, "M_authorize")
	query.Fields = append(query.Fields, "M_previous")
	query.Fields = append(query.Fields, "M_refreshToken")
	query.Fields = append(query.Fields, "M_expiresIn")
	query.Fields = append(query.Fields, "M_scope")
	query.Fields = append(query.Fields, "M_redirectUri")
	query.Fields = append(query.Fields, "M_userData")
	query.Fields = append(query.Fields, "M_createdAt")

	/** associations of OAuth2Access **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var OAuth2AccessInfo []interface{}

	OAuth2AccessInfo = append(OAuth2AccessInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		OAuth2AccessInfo = append(OAuth2AccessInfo, this.GetParentPtr().GetUuid())
		OAuth2AccessInfo = append(OAuth2AccessInfo, this.GetParentLnk())
	} else {
		OAuth2AccessInfo = append(OAuth2AccessInfo, "")
		OAuth2AccessInfo = append(OAuth2AccessInfo, "")
	}

	/** members of OAuth2Access **/
	OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_id)

	/** Save client type OAuth2Client **/
	if len(this.object.M_client) > 0 {
		OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_client)
	} else {
		OAuth2AccessInfo = append(OAuth2AccessInfo, "")
	}
	OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_authorize)
	OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_previous)

	/** Save refreshToken type OAuth2Refresh **/
	if len(this.object.M_refreshToken) > 0 {
		OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_refreshToken)
	} else {
		OAuth2AccessInfo = append(OAuth2AccessInfo, "")
	}
	OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_expiresIn)
	OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_scope)
	OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_redirectUri)

	/** Save userData type OAuth2IdToken **/
	if len(this.object.M_userData) > 0 {
		OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_userData)
	} else {
		OAuth2AccessInfo = append(OAuth2AccessInfo, "")
	}
	OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_createdAt)

	/** associations of OAuth2Access **/

	/** Save parent type OAuth2Configuration **/
	if len(this.object.M_parentPtr) > 0 {
		OAuth2AccessInfo = append(OAuth2AccessInfo, this.object.M_parentPtr)
	} else {
		OAuth2AccessInfo = append(OAuth2AccessInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), OAuth2AccessInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), OAuth2AccessInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_OAuth2AccessEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_OAuth2AccessEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.OAuth2Access"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2Access **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_client")
	query.Fields = append(query.Fields, "M_authorize")
	query.Fields = append(query.Fields, "M_previous")
	query.Fields = append(query.Fields, "M_refreshToken")
	query.Fields = append(query.Fields, "M_expiresIn")
	query.Fields = append(query.Fields, "M_scope")
	query.Fields = append(query.Fields, "M_redirectUri")
	query.Fields = append(query.Fields, "M_userData")
	query.Fields = append(query.Fields, "M_createdAt")

	/** associations of OAuth2Access **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of OAuth2Access...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.OAuth2Access)
		this.object.TYPENAME = "Config.OAuth2Access"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of OAuth2Access **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** client **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2Client"
				id_ := refTypeName + "$$" + id
				this.object.M_client = id
				GetServer().GetEntityManager().appendReference("client", this.object.UUID, id_)
			}
		}

		/** authorize **/
		if results[0][5] != nil {
			this.object.M_authorize = results[0][5].(string)
		}

		/** previous **/
		if results[0][6] != nil {
			this.object.M_previous = results[0][6].(string)
		}

		/** refreshToken **/
		if results[0][7] != nil {
			id := results[0][7].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2Refresh"
				id_ := refTypeName + "$$" + id
				this.object.M_refreshToken = id
				GetServer().GetEntityManager().appendReference("refreshToken", this.object.UUID, id_)
			}
		}

		/** expiresIn **/
		if results[0][8] != nil {
			this.object.M_expiresIn = results[0][8].(int64)
		}

		/** scope **/
		if results[0][9] != nil {
			this.object.M_scope = results[0][9].(string)
		}

		/** redirectUri **/
		if results[0][10] != nil {
			this.object.M_redirectUri = results[0][10].(string)
		}

		/** userData **/
		if results[0][11] != nil {
			id := results[0][11].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2IdToken"
				id_ := refTypeName + "$$" + id
				this.object.M_userData = id
				GetServer().GetEntityManager().appendReference("userData", this.object.UUID, id_)
			}
		}

		/** createdAt **/
		if results[0][12] != nil {
			this.object.M_createdAt = results[0][12].(int64)
		}

		/** associations of OAuth2Access **/

		/** parentPtr **/
		if results[0][13] != nil {
			id := results[0][13].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2Configuration"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigOAuth2AccessEntityFromObject(object *Config.OAuth2Access) *Config_OAuth2AccessEntity {
	return this.NewConfigOAuth2AccessEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_OAuth2AccessEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigOAuth2AccessExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Access"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_OAuth2AccessEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_OAuth2AccessEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			OAuth2Refresh
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_OAuth2RefreshEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.OAuth2Refresh
}

/** Constructor function **/
func (this *EntityManager) NewConfigOAuth2RefreshEntity(parentUuid string, objectId string, object interface{}) *Config_OAuth2RefreshEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigOAuth2RefreshExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.OAuth2Refresh).TYPENAME = "Config.OAuth2Refresh"
		object.(*Config.OAuth2Refresh).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.OAuth2Refresh", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.OAuth2Refresh).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.OAuth2Refresh).UUID
			}
			return val.(*Config_OAuth2RefreshEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_OAuth2RefreshEntity)
	if object == nil {
		entity.object = new(Config.OAuth2Refresh)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.OAuth2Refresh)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.OAuth2Refresh"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_OAuth2RefreshEntity) GetTypeName() string {
	return "Config.OAuth2Refresh"
}
func (this *Config_OAuth2RefreshEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_OAuth2RefreshEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_OAuth2RefreshEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_OAuth2RefreshEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_OAuth2RefreshEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_OAuth2RefreshEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_OAuth2RefreshEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_OAuth2RefreshEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_OAuth2RefreshEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_OAuth2RefreshEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_OAuth2RefreshEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_OAuth2RefreshEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_OAuth2RefreshEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_OAuth2RefreshEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_OAuth2RefreshEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_OAuth2RefreshEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_OAuth2RefreshEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_OAuth2RefreshEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_OAuth2RefreshEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_OAuth2RefreshEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_OAuth2RefreshEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_OAuth2RefreshEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_OAuth2RefreshEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Refresh"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_OAuth2RefreshEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_OAuth2RefreshEntityPrototype() {

	var oAuth2RefreshEntityProto EntityPrototype
	oAuth2RefreshEntityProto.TypeName = "Config.OAuth2Refresh"
	oAuth2RefreshEntityProto.Ids = append(oAuth2RefreshEntityProto.Ids, "UUID")
	oAuth2RefreshEntityProto.Fields = append(oAuth2RefreshEntityProto.Fields, "UUID")
	oAuth2RefreshEntityProto.FieldsType = append(oAuth2RefreshEntityProto.FieldsType, "xs.string")
	oAuth2RefreshEntityProto.FieldsOrder = append(oAuth2RefreshEntityProto.FieldsOrder, 0)
	oAuth2RefreshEntityProto.FieldsVisibility = append(oAuth2RefreshEntityProto.FieldsVisibility, false)
	oAuth2RefreshEntityProto.FieldsDefaultValue = append(oAuth2RefreshEntityProto.FieldsDefaultValue, "")
	oAuth2RefreshEntityProto.Indexs = append(oAuth2RefreshEntityProto.Indexs, "ParentUuid")
	oAuth2RefreshEntityProto.Fields = append(oAuth2RefreshEntityProto.Fields, "ParentUuid")
	oAuth2RefreshEntityProto.FieldsType = append(oAuth2RefreshEntityProto.FieldsType, "xs.string")
	oAuth2RefreshEntityProto.FieldsOrder = append(oAuth2RefreshEntityProto.FieldsOrder, 1)
	oAuth2RefreshEntityProto.FieldsVisibility = append(oAuth2RefreshEntityProto.FieldsVisibility, false)
	oAuth2RefreshEntityProto.FieldsDefaultValue = append(oAuth2RefreshEntityProto.FieldsDefaultValue, "")
	oAuth2RefreshEntityProto.Fields = append(oAuth2RefreshEntityProto.Fields, "ParentLnk")
	oAuth2RefreshEntityProto.FieldsType = append(oAuth2RefreshEntityProto.FieldsType, "xs.string")
	oAuth2RefreshEntityProto.FieldsOrder = append(oAuth2RefreshEntityProto.FieldsOrder, 2)
	oAuth2RefreshEntityProto.FieldsVisibility = append(oAuth2RefreshEntityProto.FieldsVisibility, false)
	oAuth2RefreshEntityProto.FieldsDefaultValue = append(oAuth2RefreshEntityProto.FieldsDefaultValue, "")

	/** members of OAuth2Refresh **/
	oAuth2RefreshEntityProto.Ids = append(oAuth2RefreshEntityProto.Ids, "M_id")
	oAuth2RefreshEntityProto.FieldsOrder = append(oAuth2RefreshEntityProto.FieldsOrder, 3)
	oAuth2RefreshEntityProto.FieldsVisibility = append(oAuth2RefreshEntityProto.FieldsVisibility, true)
	oAuth2RefreshEntityProto.Fields = append(oAuth2RefreshEntityProto.Fields, "M_id")
	oAuth2RefreshEntityProto.FieldsType = append(oAuth2RefreshEntityProto.FieldsType, "xs.ID")
	oAuth2RefreshEntityProto.FieldsDefaultValue = append(oAuth2RefreshEntityProto.FieldsDefaultValue, "")
	oAuth2RefreshEntityProto.FieldsOrder = append(oAuth2RefreshEntityProto.FieldsOrder, 4)
	oAuth2RefreshEntityProto.FieldsVisibility = append(oAuth2RefreshEntityProto.FieldsVisibility, true)
	oAuth2RefreshEntityProto.Fields = append(oAuth2RefreshEntityProto.Fields, "M_access")
	oAuth2RefreshEntityProto.FieldsDefaultValue = append(oAuth2RefreshEntityProto.FieldsDefaultValue, "undefined")
	oAuth2RefreshEntityProto.FieldsDefaultValue = append(oAuth2RefreshEntityProto.FieldsDefaultValue, "undefined")
	oAuth2RefreshEntityProto.FieldsType = append(oAuth2RefreshEntityProto.FieldsType, "Config.OAuth2Access:Ref")

	/** associations of OAuth2Refresh **/
	oAuth2RefreshEntityProto.FieldsOrder = append(oAuth2RefreshEntityProto.FieldsOrder, 5)
	oAuth2RefreshEntityProto.FieldsVisibility = append(oAuth2RefreshEntityProto.FieldsVisibility, false)
	oAuth2RefreshEntityProto.Fields = append(oAuth2RefreshEntityProto.Fields, "M_parentPtr")
	oAuth2RefreshEntityProto.FieldsDefaultValue = append(oAuth2RefreshEntityProto.FieldsDefaultValue, "undefined")
	oAuth2RefreshEntityProto.FieldsDefaultValue = append(oAuth2RefreshEntityProto.FieldsDefaultValue, "undefined")
	oAuth2RefreshEntityProto.FieldsType = append(oAuth2RefreshEntityProto.FieldsType, "Config.OAuth2Configuration:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&oAuth2RefreshEntityProto)

}

/** Create **/
func (this *Config_OAuth2RefreshEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.OAuth2Refresh"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2Refresh **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_access")

	/** associations of OAuth2Refresh **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var OAuth2RefreshInfo []interface{}

	OAuth2RefreshInfo = append(OAuth2RefreshInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		OAuth2RefreshInfo = append(OAuth2RefreshInfo, this.GetParentPtr().GetUuid())
		OAuth2RefreshInfo = append(OAuth2RefreshInfo, this.GetParentLnk())
	} else {
		OAuth2RefreshInfo = append(OAuth2RefreshInfo, "")
		OAuth2RefreshInfo = append(OAuth2RefreshInfo, "")
	}

	/** members of OAuth2Refresh **/
	OAuth2RefreshInfo = append(OAuth2RefreshInfo, this.object.M_id)

	/** Save access type OAuth2Access **/
	if len(this.object.M_access) > 0 {
		OAuth2RefreshInfo = append(OAuth2RefreshInfo, this.object.M_access)
	} else {
		OAuth2RefreshInfo = append(OAuth2RefreshInfo, "")
	}

	/** associations of OAuth2Refresh **/

	/** Save parent type OAuth2Configuration **/
	if len(this.object.M_parentPtr) > 0 {
		OAuth2RefreshInfo = append(OAuth2RefreshInfo, this.object.M_parentPtr)
	} else {
		OAuth2RefreshInfo = append(OAuth2RefreshInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), OAuth2RefreshInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), OAuth2RefreshInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_OAuth2RefreshEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_OAuth2RefreshEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.OAuth2Refresh"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2Refresh **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_access")

	/** associations of OAuth2Refresh **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of OAuth2Refresh...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.OAuth2Refresh)
		this.object.TYPENAME = "Config.OAuth2Refresh"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of OAuth2Refresh **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** access **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2Access"
				id_ := refTypeName + "$$" + id
				this.object.M_access = id
				GetServer().GetEntityManager().appendReference("access", this.object.UUID, id_)
			}
		}

		/** associations of OAuth2Refresh **/

		/** parentPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2Configuration"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigOAuth2RefreshEntityFromObject(object *Config.OAuth2Refresh) *Config_OAuth2RefreshEntity {
	return this.NewConfigOAuth2RefreshEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_OAuth2RefreshEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigOAuth2RefreshExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Refresh"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_OAuth2RefreshEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_OAuth2RefreshEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			OAuth2Expires
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_OAuth2ExpiresEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.OAuth2Expires
}

/** Constructor function **/
func (this *EntityManager) NewConfigOAuth2ExpiresEntity(parentUuid string, objectId string, object interface{}) *Config_OAuth2ExpiresEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigOAuth2ExpiresExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.OAuth2Expires).TYPENAME = "Config.OAuth2Expires"
		object.(*Config.OAuth2Expires).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.OAuth2Expires", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.OAuth2Expires).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.OAuth2Expires).UUID
			}
			return val.(*Config_OAuth2ExpiresEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_OAuth2ExpiresEntity)
	if object == nil {
		entity.object = new(Config.OAuth2Expires)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.OAuth2Expires)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.OAuth2Expires"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_OAuth2ExpiresEntity) GetTypeName() string {
	return "Config.OAuth2Expires"
}
func (this *Config_OAuth2ExpiresEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_OAuth2ExpiresEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_OAuth2ExpiresEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_OAuth2ExpiresEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_OAuth2ExpiresEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_OAuth2ExpiresEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_OAuth2ExpiresEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_OAuth2ExpiresEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_OAuth2ExpiresEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_OAuth2ExpiresEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_OAuth2ExpiresEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_OAuth2ExpiresEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_OAuth2ExpiresEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_OAuth2ExpiresEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_OAuth2ExpiresEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_OAuth2ExpiresEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_OAuth2ExpiresEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_OAuth2ExpiresEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_OAuth2ExpiresEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_OAuth2ExpiresEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_OAuth2ExpiresEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_OAuth2ExpiresEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_OAuth2ExpiresEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Expires"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_OAuth2ExpiresEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_OAuth2ExpiresEntityPrototype() {

	var oAuth2ExpiresEntityProto EntityPrototype
	oAuth2ExpiresEntityProto.TypeName = "Config.OAuth2Expires"
	oAuth2ExpiresEntityProto.Ids = append(oAuth2ExpiresEntityProto.Ids, "UUID")
	oAuth2ExpiresEntityProto.Fields = append(oAuth2ExpiresEntityProto.Fields, "UUID")
	oAuth2ExpiresEntityProto.FieldsType = append(oAuth2ExpiresEntityProto.FieldsType, "xs.string")
	oAuth2ExpiresEntityProto.FieldsOrder = append(oAuth2ExpiresEntityProto.FieldsOrder, 0)
	oAuth2ExpiresEntityProto.FieldsVisibility = append(oAuth2ExpiresEntityProto.FieldsVisibility, false)
	oAuth2ExpiresEntityProto.FieldsDefaultValue = append(oAuth2ExpiresEntityProto.FieldsDefaultValue, "")
	oAuth2ExpiresEntityProto.Indexs = append(oAuth2ExpiresEntityProto.Indexs, "ParentUuid")
	oAuth2ExpiresEntityProto.Fields = append(oAuth2ExpiresEntityProto.Fields, "ParentUuid")
	oAuth2ExpiresEntityProto.FieldsType = append(oAuth2ExpiresEntityProto.FieldsType, "xs.string")
	oAuth2ExpiresEntityProto.FieldsOrder = append(oAuth2ExpiresEntityProto.FieldsOrder, 1)
	oAuth2ExpiresEntityProto.FieldsVisibility = append(oAuth2ExpiresEntityProto.FieldsVisibility, false)
	oAuth2ExpiresEntityProto.FieldsDefaultValue = append(oAuth2ExpiresEntityProto.FieldsDefaultValue, "")
	oAuth2ExpiresEntityProto.Fields = append(oAuth2ExpiresEntityProto.Fields, "ParentLnk")
	oAuth2ExpiresEntityProto.FieldsType = append(oAuth2ExpiresEntityProto.FieldsType, "xs.string")
	oAuth2ExpiresEntityProto.FieldsOrder = append(oAuth2ExpiresEntityProto.FieldsOrder, 2)
	oAuth2ExpiresEntityProto.FieldsVisibility = append(oAuth2ExpiresEntityProto.FieldsVisibility, false)
	oAuth2ExpiresEntityProto.FieldsDefaultValue = append(oAuth2ExpiresEntityProto.FieldsDefaultValue, "")

	/** members of OAuth2Expires **/
	oAuth2ExpiresEntityProto.Ids = append(oAuth2ExpiresEntityProto.Ids, "M_id")
	oAuth2ExpiresEntityProto.FieldsOrder = append(oAuth2ExpiresEntityProto.FieldsOrder, 3)
	oAuth2ExpiresEntityProto.FieldsVisibility = append(oAuth2ExpiresEntityProto.FieldsVisibility, true)
	oAuth2ExpiresEntityProto.Fields = append(oAuth2ExpiresEntityProto.Fields, "M_id")
	oAuth2ExpiresEntityProto.FieldsType = append(oAuth2ExpiresEntityProto.FieldsType, "xs.ID")
	oAuth2ExpiresEntityProto.FieldsDefaultValue = append(oAuth2ExpiresEntityProto.FieldsDefaultValue, "")
	oAuth2ExpiresEntityProto.FieldsOrder = append(oAuth2ExpiresEntityProto.FieldsOrder, 4)
	oAuth2ExpiresEntityProto.FieldsVisibility = append(oAuth2ExpiresEntityProto.FieldsVisibility, true)
	oAuth2ExpiresEntityProto.Fields = append(oAuth2ExpiresEntityProto.Fields, "M_expiresAt")
	oAuth2ExpiresEntityProto.FieldsType = append(oAuth2ExpiresEntityProto.FieldsType, "xs.date")
	oAuth2ExpiresEntityProto.FieldsDefaultValue = append(oAuth2ExpiresEntityProto.FieldsDefaultValue, "new Date()")

	/** associations of OAuth2Expires **/
	oAuth2ExpiresEntityProto.FieldsOrder = append(oAuth2ExpiresEntityProto.FieldsOrder, 5)
	oAuth2ExpiresEntityProto.FieldsVisibility = append(oAuth2ExpiresEntityProto.FieldsVisibility, false)
	oAuth2ExpiresEntityProto.Fields = append(oAuth2ExpiresEntityProto.Fields, "M_parentPtr")
	oAuth2ExpiresEntityProto.FieldsDefaultValue = append(oAuth2ExpiresEntityProto.FieldsDefaultValue, "undefined")
	oAuth2ExpiresEntityProto.FieldsDefaultValue = append(oAuth2ExpiresEntityProto.FieldsDefaultValue, "undefined")
	oAuth2ExpiresEntityProto.FieldsType = append(oAuth2ExpiresEntityProto.FieldsType, "Config.OAuth2Configuration:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&oAuth2ExpiresEntityProto)

}

/** Create **/
func (this *Config_OAuth2ExpiresEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.OAuth2Expires"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2Expires **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_expiresAt")

	/** associations of OAuth2Expires **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var OAuth2ExpiresInfo []interface{}

	OAuth2ExpiresInfo = append(OAuth2ExpiresInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		OAuth2ExpiresInfo = append(OAuth2ExpiresInfo, this.GetParentPtr().GetUuid())
		OAuth2ExpiresInfo = append(OAuth2ExpiresInfo, this.GetParentLnk())
	} else {
		OAuth2ExpiresInfo = append(OAuth2ExpiresInfo, "")
		OAuth2ExpiresInfo = append(OAuth2ExpiresInfo, "")
	}

	/** members of OAuth2Expires **/
	OAuth2ExpiresInfo = append(OAuth2ExpiresInfo, this.object.M_id)
	OAuth2ExpiresInfo = append(OAuth2ExpiresInfo, this.object.M_expiresAt)

	/** associations of OAuth2Expires **/

	/** Save parent type OAuth2Configuration **/
	if len(this.object.M_parentPtr) > 0 {
		OAuth2ExpiresInfo = append(OAuth2ExpiresInfo, this.object.M_parentPtr)
	} else {
		OAuth2ExpiresInfo = append(OAuth2ExpiresInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), OAuth2ExpiresInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), OAuth2ExpiresInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_OAuth2ExpiresEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_OAuth2ExpiresEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.OAuth2Expires"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of OAuth2Expires **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_expiresAt")

	/** associations of OAuth2Expires **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of OAuth2Expires...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.OAuth2Expires)
		this.object.TYPENAME = "Config.OAuth2Expires"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of OAuth2Expires **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** expiresAt **/
		if results[0][4] != nil {
			this.object.M_expiresAt = results[0][4].(int64)
		}

		/** associations of OAuth2Expires **/

		/** parentPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "Config.OAuth2Configuration"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigOAuth2ExpiresEntityFromObject(object *Config.OAuth2Expires) *Config_OAuth2ExpiresEntity {
	return this.NewConfigOAuth2ExpiresEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_OAuth2ExpiresEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigOAuth2ExpiresExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Expires"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_OAuth2ExpiresEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_OAuth2ExpiresEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			OAuth2Configuration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_OAuth2ConfigurationEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.OAuth2Configuration
}

/** Constructor function **/
func (this *EntityManager) NewConfigOAuth2ConfigurationEntity(parentUuid string, objectId string, object interface{}) *Config_OAuth2ConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigOAuth2ConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.OAuth2Configuration).TYPENAME = "Config.OAuth2Configuration"
		object.(*Config.OAuth2Configuration).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.OAuth2Configuration", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.OAuth2Configuration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.OAuth2Configuration).UUID
			}
			return val.(*Config_OAuth2ConfigurationEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_OAuth2ConfigurationEntity)
	if object == nil {
		entity.object = new(Config.OAuth2Configuration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.OAuth2Configuration)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.OAuth2Configuration"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_OAuth2ConfigurationEntity) GetTypeName() string {
	return "Config.OAuth2Configuration"
}
func (this *Config_OAuth2ConfigurationEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_OAuth2ConfigurationEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_OAuth2ConfigurationEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_OAuth2ConfigurationEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_OAuth2ConfigurationEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_OAuth2ConfigurationEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_OAuth2ConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_OAuth2ConfigurationEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_OAuth2ConfigurationEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_OAuth2ConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_OAuth2ConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_OAuth2ConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_OAuth2ConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_OAuth2ConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_OAuth2ConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_OAuth2ConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_OAuth2ConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_OAuth2ConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_OAuth2ConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_OAuth2ConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_OAuth2ConfigurationEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_OAuth2ConfigurationEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_OAuth2ConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Configuration"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_OAuth2ConfigurationEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_OAuth2ConfigurationEntityPrototype() {

	var oAuth2ConfigurationEntityProto EntityPrototype
	oAuth2ConfigurationEntityProto.TypeName = "Config.OAuth2Configuration"
	oAuth2ConfigurationEntityProto.SuperTypeNames = append(oAuth2ConfigurationEntityProto.SuperTypeNames, "Config.Configuration")
	oAuth2ConfigurationEntityProto.Ids = append(oAuth2ConfigurationEntityProto.Ids, "UUID")
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "UUID")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.string")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 0)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, false)
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "")
	oAuth2ConfigurationEntityProto.Indexs = append(oAuth2ConfigurationEntityProto.Indexs, "ParentUuid")
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "ParentUuid")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.string")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 1)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, false)
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "")
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "ParentLnk")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.string")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 2)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, false)
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of Configuration **/
	oAuth2ConfigurationEntityProto.Ids = append(oAuth2ConfigurationEntityProto.Ids, "M_id")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 3)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_id")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.ID")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of OAuth2Configuration **/
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 4)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_authorizationExpiration")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.int")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "0")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 5)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_accessExpiration")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.time")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "0")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 6)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_tokenType")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.string")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 7)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_errorStatusCode")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.int")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "0")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 8)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_allowClientSecretInParams")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.boolean")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "false")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 9)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_allowGetAccessRequest")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.boolean")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "false")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 10)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_redirectUriSeparator")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.string")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 11)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_privateKey")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "xs.string")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 12)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_allowedAuthorizeTypes")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "[]xs.string")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "[]")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 13)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_allowedAccessTypes")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "[]xs.string")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "[]")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 14)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_clients")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "[]")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "[]Config.OAuth2Client")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 15)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_authorize")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "[]")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "[]Config.OAuth2Authorize")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 16)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_access")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "[]")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "[]Config.OAuth2Access")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 17)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_ids")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "[]")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "[]Config.OAuth2IdToken")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 18)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_refresh")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "[]")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "[]Config.OAuth2Refresh")
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 19)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, true)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_expire")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "[]")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "[]Config.OAuth2Expires")

	/** associations of OAuth2Configuration **/
	oAuth2ConfigurationEntityProto.FieldsOrder = append(oAuth2ConfigurationEntityProto.FieldsOrder, 20)
	oAuth2ConfigurationEntityProto.FieldsVisibility = append(oAuth2ConfigurationEntityProto.FieldsVisibility, false)
	oAuth2ConfigurationEntityProto.Fields = append(oAuth2ConfigurationEntityProto.Fields, "M_parentPtr")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "undefined")
	oAuth2ConfigurationEntityProto.FieldsDefaultValue = append(oAuth2ConfigurationEntityProto.FieldsDefaultValue, "undefined")
	oAuth2ConfigurationEntityProto.FieldsType = append(oAuth2ConfigurationEntityProto.FieldsType, "Config.Configurations:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&oAuth2ConfigurationEntityProto)

}

/** Create **/
func (this *Config_OAuth2ConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.OAuth2Configuration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of OAuth2Configuration **/
	query.Fields = append(query.Fields, "M_authorizationExpiration")
	query.Fields = append(query.Fields, "M_accessExpiration")
	query.Fields = append(query.Fields, "M_tokenType")
	query.Fields = append(query.Fields, "M_errorStatusCode")
	query.Fields = append(query.Fields, "M_allowClientSecretInParams")
	query.Fields = append(query.Fields, "M_allowGetAccessRequest")
	query.Fields = append(query.Fields, "M_redirectUriSeparator")
	query.Fields = append(query.Fields, "M_privateKey")
	query.Fields = append(query.Fields, "M_allowedAuthorizeTypes")
	query.Fields = append(query.Fields, "M_allowedAccessTypes")
	query.Fields = append(query.Fields, "M_clients")
	query.Fields = append(query.Fields, "M_authorize")
	query.Fields = append(query.Fields, "M_access")
	query.Fields = append(query.Fields, "M_ids")
	query.Fields = append(query.Fields, "M_refresh")
	query.Fields = append(query.Fields, "M_expire")

	/** associations of OAuth2Configuration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var OAuth2ConfigurationInfo []interface{}

	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.GetParentPtr().GetUuid())
		OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.GetParentLnk())
	} else {
		OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, "")
		OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, "")
	}

	/** members of Configuration **/
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_id)

	/** members of OAuth2Configuration **/
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_authorizationExpiration)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_accessExpiration)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_tokenType)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_errorStatusCode)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_allowClientSecretInParams)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_allowGetAccessRequest)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_redirectUriSeparator)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_privateKey)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_allowedAuthorizeTypes)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_allowedAccessTypes)

	/** Save clients type OAuth2Client **/
	clientsIds := make([]string, 0)
	lazy_clients := this.lazyMap["M_clients"] != nil && len(this.object.M_clients) == 0
	if !lazy_clients {
		for i := 0; i < len(this.object.M_clients); i++ {
			clientsEntity := GetServer().GetEntityManager().NewConfigOAuth2ClientEntity(this.GetUuid(), this.object.M_clients[i].UUID, this.object.M_clients[i])
			clientsIds = append(clientsIds, clientsEntity.GetUuid())
			clientsEntity.AppendReferenced("clients", this)
			this.AppendChild("clients", clientsEntity)
			if clientsEntity.NeedSave() {
				clientsEntity.SaveEntity()
			}
		}
	} else {
		clientsIds = this.lazyMap["M_clients"].([]string)
	}
	clientsStr, _ := json.Marshal(clientsIds)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, string(clientsStr))

	/** Save authorize type OAuth2Authorize **/
	authorizeIds := make([]string, 0)
	lazy_authorize := this.lazyMap["M_authorize"] != nil && len(this.object.M_authorize) == 0
	if !lazy_authorize {
		for i := 0; i < len(this.object.M_authorize); i++ {
			authorizeEntity := GetServer().GetEntityManager().NewConfigOAuth2AuthorizeEntity(this.GetUuid(), this.object.M_authorize[i].UUID, this.object.M_authorize[i])
			authorizeIds = append(authorizeIds, authorizeEntity.GetUuid())
			authorizeEntity.AppendReferenced("authorize", this)
			this.AppendChild("authorize", authorizeEntity)
			if authorizeEntity.NeedSave() {
				authorizeEntity.SaveEntity()
			}
		}
	} else {
		authorizeIds = this.lazyMap["M_authorize"].([]string)
	}
	authorizeStr, _ := json.Marshal(authorizeIds)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, string(authorizeStr))

	/** Save access type OAuth2Access **/
	accessIds := make([]string, 0)
	lazy_access := this.lazyMap["M_access"] != nil && len(this.object.M_access) == 0
	if !lazy_access {
		for i := 0; i < len(this.object.M_access); i++ {
			accessEntity := GetServer().GetEntityManager().NewConfigOAuth2AccessEntity(this.GetUuid(), this.object.M_access[i].UUID, this.object.M_access[i])
			accessIds = append(accessIds, accessEntity.GetUuid())
			accessEntity.AppendReferenced("access", this)
			this.AppendChild("access", accessEntity)
			if accessEntity.NeedSave() {
				accessEntity.SaveEntity()
			}
		}
	} else {
		accessIds = this.lazyMap["M_access"].([]string)
	}
	accessStr, _ := json.Marshal(accessIds)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, string(accessStr))

	/** Save ids type OAuth2IdToken **/
	idsIds := make([]string, 0)
	lazy_ids := this.lazyMap["M_ids"] != nil && len(this.object.M_ids) == 0
	if !lazy_ids {
		for i := 0; i < len(this.object.M_ids); i++ {
			idsEntity := GetServer().GetEntityManager().NewConfigOAuth2IdTokenEntity(this.GetUuid(), this.object.M_ids[i].UUID, this.object.M_ids[i])
			idsIds = append(idsIds, idsEntity.GetUuid())
			idsEntity.AppendReferenced("ids", this)
			this.AppendChild("ids", idsEntity)
			if idsEntity.NeedSave() {
				idsEntity.SaveEntity()
			}
		}
	} else {
		idsIds = this.lazyMap["M_ids"].([]string)
	}
	idsStr, _ := json.Marshal(idsIds)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, string(idsStr))

	/** Save refresh type OAuth2Refresh **/
	refreshIds := make([]string, 0)
	lazy_refresh := this.lazyMap["M_refresh"] != nil && len(this.object.M_refresh) == 0
	if !lazy_refresh {
		for i := 0; i < len(this.object.M_refresh); i++ {
			refreshEntity := GetServer().GetEntityManager().NewConfigOAuth2RefreshEntity(this.GetUuid(), this.object.M_refresh[i].UUID, this.object.M_refresh[i])
			refreshIds = append(refreshIds, refreshEntity.GetUuid())
			refreshEntity.AppendReferenced("refresh", this)
			this.AppendChild("refresh", refreshEntity)
			if refreshEntity.NeedSave() {
				refreshEntity.SaveEntity()
			}
		}
	} else {
		refreshIds = this.lazyMap["M_refresh"].([]string)
	}
	refreshStr, _ := json.Marshal(refreshIds)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, string(refreshStr))

	/** Save expire type OAuth2Expires **/
	expireIds := make([]string, 0)
	lazy_expire := this.lazyMap["M_expire"] != nil && len(this.object.M_expire) == 0
	if !lazy_expire {
		for i := 0; i < len(this.object.M_expire); i++ {
			expireEntity := GetServer().GetEntityManager().NewConfigOAuth2ExpiresEntity(this.GetUuid(), this.object.M_expire[i].UUID, this.object.M_expire[i])
			expireIds = append(expireIds, expireEntity.GetUuid())
			expireEntity.AppendReferenced("expire", this)
			this.AppendChild("expire", expireEntity)
			if expireEntity.NeedSave() {
				expireEntity.SaveEntity()
			}
		}
	} else {
		expireIds = this.lazyMap["M_expire"].([]string)
	}
	expireStr, _ := json.Marshal(expireIds)
	OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, string(expireStr))

	/** associations of OAuth2Configuration **/

	/** Save parent type Configurations **/
	if len(this.object.M_parentPtr) > 0 {
		OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, this.object.M_parentPtr)
	} else {
		OAuth2ConfigurationInfo = append(OAuth2ConfigurationInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), OAuth2ConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), OAuth2ConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_OAuth2ConfigurationEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_OAuth2ConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.OAuth2Configuration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of OAuth2Configuration **/
	query.Fields = append(query.Fields, "M_authorizationExpiration")
	query.Fields = append(query.Fields, "M_accessExpiration")
	query.Fields = append(query.Fields, "M_tokenType")
	query.Fields = append(query.Fields, "M_errorStatusCode")
	query.Fields = append(query.Fields, "M_allowClientSecretInParams")
	query.Fields = append(query.Fields, "M_allowGetAccessRequest")
	query.Fields = append(query.Fields, "M_redirectUriSeparator")
	query.Fields = append(query.Fields, "M_privateKey")
	query.Fields = append(query.Fields, "M_allowedAuthorizeTypes")
	query.Fields = append(query.Fields, "M_allowedAccessTypes")
	query.Fields = append(query.Fields, "M_clients")
	query.Fields = append(query.Fields, "M_authorize")
	query.Fields = append(query.Fields, "M_access")
	query.Fields = append(query.Fields, "M_ids")
	query.Fields = append(query.Fields, "M_refresh")
	query.Fields = append(query.Fields, "M_expire")

	/** associations of OAuth2Configuration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of OAuth2Configuration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.OAuth2Configuration)
		this.object.TYPENAME = "Config.OAuth2Configuration"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of OAuth2Configuration **/

		/** authorizationExpiration **/
		if results[0][4] != nil {
			this.object.M_authorizationExpiration = results[0][4].(int)
		}

		/** accessExpiration **/
		if results[0][5] != nil {
			this.object.M_accessExpiration = results[0][5].(int64)
		}

		/** tokenType **/
		if results[0][6] != nil {
			this.object.M_tokenType = results[0][6].(string)
		}

		/** errorStatusCode **/
		if results[0][7] != nil {
			this.object.M_errorStatusCode = results[0][7].(int)
		}

		/** allowClientSecretInParams **/
		if results[0][8] != nil {
			this.object.M_allowClientSecretInParams = results[0][8].(bool)
		}

		/** allowGetAccessRequest **/
		if results[0][9] != nil {
			this.object.M_allowGetAccessRequest = results[0][9].(bool)
		}

		/** redirectUriSeparator **/
		if results[0][10] != nil {
			this.object.M_redirectUriSeparator = results[0][10].(string)
		}

		/** privateKey **/
		if results[0][11] != nil {
			this.object.M_privateKey = results[0][11].(string)
		}

		/** allowedAuthorizeTypes **/
		if results[0][12] != nil {
			this.object.M_allowedAuthorizeTypes = append(this.object.M_allowedAuthorizeTypes, results[0][12].([]string)...)
		}

		/** allowedAccessTypes **/
		if results[0][13] != nil {
			this.object.M_allowedAccessTypes = append(this.object.M_allowedAccessTypes, results[0][13].([]string)...)
		}

		/** clients **/
		if results[0][14] != nil {
			uuidsStr := results[0][14].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var clientsEntity *Config_OAuth2ClientEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							clientsEntity = instance.(*Config_OAuth2ClientEntity)
						} else {
							clientsEntity = GetServer().GetEntityManager().NewConfigOAuth2ClientEntity(this.GetUuid(), uuids[i], nil)
							clientsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(clientsEntity)
						}
						clientsEntity.AppendReferenced("clients", this)
						this.AppendChild("clients", clientsEntity)
					}
				} else {
					this.lazyMap["M_clients"] = uuids
				}
			}
		}

		/** authorize **/
		if results[0][15] != nil {
			uuidsStr := results[0][15].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var authorizeEntity *Config_OAuth2AuthorizeEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							authorizeEntity = instance.(*Config_OAuth2AuthorizeEntity)
						} else {
							authorizeEntity = GetServer().GetEntityManager().NewConfigOAuth2AuthorizeEntity(this.GetUuid(), uuids[i], nil)
							authorizeEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(authorizeEntity)
						}
						authorizeEntity.AppendReferenced("authorize", this)
						this.AppendChild("authorize", authorizeEntity)
					}
				} else {
					this.lazyMap["M_authorize"] = uuids
				}
			}
		}

		/** access **/
		if results[0][16] != nil {
			uuidsStr := results[0][16].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var accessEntity *Config_OAuth2AccessEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							accessEntity = instance.(*Config_OAuth2AccessEntity)
						} else {
							accessEntity = GetServer().GetEntityManager().NewConfigOAuth2AccessEntity(this.GetUuid(), uuids[i], nil)
							accessEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(accessEntity)
						}
						accessEntity.AppendReferenced("access", this)
						this.AppendChild("access", accessEntity)
					}
				} else {
					this.lazyMap["M_access"] = uuids
				}
			}
		}

		/** ids **/
		if results[0][17] != nil {
			uuidsStr := results[0][17].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var idsEntity *Config_OAuth2IdTokenEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							idsEntity = instance.(*Config_OAuth2IdTokenEntity)
						} else {
							idsEntity = GetServer().GetEntityManager().NewConfigOAuth2IdTokenEntity(this.GetUuid(), uuids[i], nil)
							idsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(idsEntity)
						}
						idsEntity.AppendReferenced("ids", this)
						this.AppendChild("ids", idsEntity)
					}
				} else {
					this.lazyMap["M_ids"] = uuids
				}
			}
		}

		/** refresh **/
		if results[0][18] != nil {
			uuidsStr := results[0][18].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var refreshEntity *Config_OAuth2RefreshEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							refreshEntity = instance.(*Config_OAuth2RefreshEntity)
						} else {
							refreshEntity = GetServer().GetEntityManager().NewConfigOAuth2RefreshEntity(this.GetUuid(), uuids[i], nil)
							refreshEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(refreshEntity)
						}
						refreshEntity.AppendReferenced("refresh", this)
						this.AppendChild("refresh", refreshEntity)
					}
				} else {
					this.lazyMap["M_refresh"] = uuids
				}
			}
		}

		/** expire **/
		if results[0][19] != nil {
			uuidsStr := results[0][19].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var expireEntity *Config_OAuth2ExpiresEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							expireEntity = instance.(*Config_OAuth2ExpiresEntity)
						} else {
							expireEntity = GetServer().GetEntityManager().NewConfigOAuth2ExpiresEntity(this.GetUuid(), uuids[i], nil)
							expireEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(expireEntity)
						}
						expireEntity.AppendReferenced("expire", this)
						this.AppendChild("expire", expireEntity)
					}
				} else {
					this.lazyMap["M_expire"] = uuids
				}
			}
		}

		/** associations of OAuth2Configuration **/

		/** parentPtr **/
		if results[0][20] != nil {
			id := results[0][20].(string)
			if len(id) > 0 {
				refTypeName := "Config.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigOAuth2ConfigurationEntityFromObject(object *Config.OAuth2Configuration) *Config_OAuth2ConfigurationEntity {
	return this.NewConfigOAuth2ConfigurationEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_OAuth2ConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigOAuth2ConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.OAuth2Configuration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_OAuth2ConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_OAuth2ConfigurationEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			ServiceConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_ServiceConfigurationEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.ServiceConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewConfigServiceConfigurationEntity(parentUuid string, objectId string, object interface{}) *Config_ServiceConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigServiceConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.ServiceConfiguration).TYPENAME = "Config.ServiceConfiguration"
		object.(*Config.ServiceConfiguration).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.ServiceConfiguration", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.ServiceConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.ServiceConfiguration).UUID
			}
			return val.(*Config_ServiceConfigurationEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_ServiceConfigurationEntity)
	if object == nil {
		entity.object = new(Config.ServiceConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.ServiceConfiguration)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.ServiceConfiguration"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_ServiceConfigurationEntity) GetTypeName() string {
	return "Config.ServiceConfiguration"
}
func (this *Config_ServiceConfigurationEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_ServiceConfigurationEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_ServiceConfigurationEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_ServiceConfigurationEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_ServiceConfigurationEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_ServiceConfigurationEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_ServiceConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_ServiceConfigurationEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_ServiceConfigurationEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_ServiceConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_ServiceConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_ServiceConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_ServiceConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_ServiceConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_ServiceConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_ServiceConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_ServiceConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_ServiceConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_ServiceConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_ServiceConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_ServiceConfigurationEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_ServiceConfigurationEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_ServiceConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.ServiceConfiguration"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_ServiceConfigurationEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_ServiceConfigurationEntityPrototype() {

	var serviceConfigurationEntityProto EntityPrototype
	serviceConfigurationEntityProto.TypeName = "Config.ServiceConfiguration"
	serviceConfigurationEntityProto.SuperTypeNames = append(serviceConfigurationEntityProto.SuperTypeNames, "Config.Configuration")
	serviceConfigurationEntityProto.Ids = append(serviceConfigurationEntityProto.Ids, "UUID")
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "UUID")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "xs.string")
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 0)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, false)
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "")
	serviceConfigurationEntityProto.Indexs = append(serviceConfigurationEntityProto.Indexs, "ParentUuid")
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "ParentUuid")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "xs.string")
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 1)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, false)
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "")
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "ParentLnk")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "xs.string")
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 2)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, false)
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of Configuration **/
	serviceConfigurationEntityProto.Ids = append(serviceConfigurationEntityProto.Ids, "M_id")
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 3)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, true)
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "M_id")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "xs.ID")
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of ServiceConfiguration **/
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 4)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, true)
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "M_hostName")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "xs.string")
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "")
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 5)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, true)
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "M_ipv4")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "xs.string")
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "")
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 6)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, true)
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "M_port")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "xs.int")
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "0")
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 7)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, true)
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "M_user")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "xs.string")
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "")
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 8)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, true)
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "M_pwd")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "xs.string")
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "")
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 9)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, true)
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "M_start")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "xs.boolean")
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "false")

	/** associations of ServiceConfiguration **/
	serviceConfigurationEntityProto.FieldsOrder = append(serviceConfigurationEntityProto.FieldsOrder, 10)
	serviceConfigurationEntityProto.FieldsVisibility = append(serviceConfigurationEntityProto.FieldsVisibility, false)
	serviceConfigurationEntityProto.Fields = append(serviceConfigurationEntityProto.Fields, "M_parentPtr")
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "undefined")
	serviceConfigurationEntityProto.FieldsDefaultValue = append(serviceConfigurationEntityProto.FieldsDefaultValue, "undefined")
	serviceConfigurationEntityProto.FieldsType = append(serviceConfigurationEntityProto.FieldsType, "Config.Configurations:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&serviceConfigurationEntityProto)

}

/** Create **/
func (this *Config_ServiceConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.ServiceConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ServiceConfiguration **/
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_port")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")
	query.Fields = append(query.Fields, "M_start")

	/** associations of ServiceConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var ServiceConfigurationInfo []interface{}

	ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.GetParentPtr().GetUuid())
		ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.GetParentLnk())
	} else {
		ServiceConfigurationInfo = append(ServiceConfigurationInfo, "")
		ServiceConfigurationInfo = append(ServiceConfigurationInfo, "")
	}

	/** members of Configuration **/
	ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.object.M_id)

	/** members of ServiceConfiguration **/
	ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.object.M_hostName)
	ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.object.M_ipv4)
	ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.object.M_port)
	ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.object.M_user)
	ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.object.M_pwd)
	ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.object.M_start)

	/** associations of ServiceConfiguration **/

	/** Save parent type Configurations **/
	if len(this.object.M_parentPtr) > 0 {
		ServiceConfigurationInfo = append(ServiceConfigurationInfo, this.object.M_parentPtr)
	} else {
		ServiceConfigurationInfo = append(ServiceConfigurationInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), ServiceConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), ServiceConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_ServiceConfigurationEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_ServiceConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.ServiceConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ServiceConfiguration **/
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_port")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")
	query.Fields = append(query.Fields, "M_start")

	/** associations of ServiceConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ServiceConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.ServiceConfiguration)
		this.object.TYPENAME = "Config.ServiceConfiguration"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of ServiceConfiguration **/

		/** hostName **/
		if results[0][4] != nil {
			this.object.M_hostName = results[0][4].(string)
		}

		/** ipv4 **/
		if results[0][5] != nil {
			this.object.M_ipv4 = results[0][5].(string)
		}

		/** port **/
		if results[0][6] != nil {
			this.object.M_port = results[0][6].(int)
		}

		/** user **/
		if results[0][7] != nil {
			this.object.M_user = results[0][7].(string)
		}

		/** pwd **/
		if results[0][8] != nil {
			this.object.M_pwd = results[0][8].(string)
		}

		/** start **/
		if results[0][9] != nil {
			this.object.M_start = results[0][9].(bool)
		}

		/** associations of ServiceConfiguration **/

		/** parentPtr **/
		if results[0][10] != nil {
			id := results[0][10].(string)
			if len(id) > 0 {
				refTypeName := "Config.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigServiceConfigurationEntityFromObject(object *Config.ServiceConfiguration) *Config_ServiceConfigurationEntity {
	return this.NewConfigServiceConfigurationEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_ServiceConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigServiceConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.ServiceConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_ServiceConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_ServiceConfigurationEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			ScheduledTask
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_ScheduledTaskEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.ScheduledTask
}

/** Constructor function **/
func (this *EntityManager) NewConfigScheduledTaskEntity(parentUuid string, objectId string, object interface{}) *Config_ScheduledTaskEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigScheduledTaskExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.ScheduledTask).TYPENAME = "Config.ScheduledTask"
		object.(*Config.ScheduledTask).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.ScheduledTask", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.ScheduledTask).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.ScheduledTask).UUID
			}
			return val.(*Config_ScheduledTaskEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_ScheduledTaskEntity)
	if object == nil {
		entity.object = new(Config.ScheduledTask)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.ScheduledTask)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.ScheduledTask"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_ScheduledTaskEntity) GetTypeName() string {
	return "Config.ScheduledTask"
}
func (this *Config_ScheduledTaskEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_ScheduledTaskEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_ScheduledTaskEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_ScheduledTaskEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_ScheduledTaskEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_ScheduledTaskEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_ScheduledTaskEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_ScheduledTaskEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_ScheduledTaskEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_ScheduledTaskEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_ScheduledTaskEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_ScheduledTaskEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_ScheduledTaskEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_ScheduledTaskEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_ScheduledTaskEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_ScheduledTaskEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_ScheduledTaskEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_ScheduledTaskEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_ScheduledTaskEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_ScheduledTaskEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_ScheduledTaskEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_ScheduledTaskEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_ScheduledTaskEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.ScheduledTask"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_ScheduledTaskEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_ScheduledTaskEntityPrototype() {

	var scheduledTaskEntityProto EntityPrototype
	scheduledTaskEntityProto.TypeName = "Config.ScheduledTask"
	scheduledTaskEntityProto.SuperTypeNames = append(scheduledTaskEntityProto.SuperTypeNames, "Config.Configuration")
	scheduledTaskEntityProto.Ids = append(scheduledTaskEntityProto.Ids, "UUID")
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "UUID")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "xs.string")
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 0)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, false)
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "")
	scheduledTaskEntityProto.Indexs = append(scheduledTaskEntityProto.Indexs, "ParentUuid")
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "ParentUuid")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "xs.string")
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 1)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, false)
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "")
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "ParentLnk")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "xs.string")
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 2)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, false)
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "")

	/** members of Configuration **/
	scheduledTaskEntityProto.Ids = append(scheduledTaskEntityProto.Ids, "M_id")
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 3)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, true)
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "M_id")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "xs.ID")
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "")

	/** members of ScheduledTask **/
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 4)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, true)
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "M_isActive")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "xs.boolean")
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "false")
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 5)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, true)
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "M_script")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "xs.string")
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "")
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 6)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, true)
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "M_startTime")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "xs.time")
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "0")
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 7)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, true)
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "M_expirationTime")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "xs.time")
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "0")
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 8)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, true)
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "M_frequency")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "xs.int")
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "0")
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 9)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, true)
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "M_frequencyType")
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "1")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "enum:FrequencyType_ONCE:FrequencyType_DAILY:FrequencyType_WEEKELY:FrequencyType_MONTHLY")
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 10)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, true)
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "M_offsets")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "[]xs.int")
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "[]")

	/** associations of ScheduledTask **/
	scheduledTaskEntityProto.FieldsOrder = append(scheduledTaskEntityProto.FieldsOrder, 11)
	scheduledTaskEntityProto.FieldsVisibility = append(scheduledTaskEntityProto.FieldsVisibility, false)
	scheduledTaskEntityProto.Fields = append(scheduledTaskEntityProto.Fields, "M_parentPtr")
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "undefined")
	scheduledTaskEntityProto.FieldsDefaultValue = append(scheduledTaskEntityProto.FieldsDefaultValue, "undefined")
	scheduledTaskEntityProto.FieldsType = append(scheduledTaskEntityProto.FieldsType, "Config.Configurations:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&scheduledTaskEntityProto)

}

/** Create **/
func (this *Config_ScheduledTaskEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.ScheduledTask"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ScheduledTask **/
	query.Fields = append(query.Fields, "M_isActive")
	query.Fields = append(query.Fields, "M_script")
	query.Fields = append(query.Fields, "M_startTime")
	query.Fields = append(query.Fields, "M_expirationTime")
	query.Fields = append(query.Fields, "M_frequency")
	query.Fields = append(query.Fields, "M_frequencyType")
	query.Fields = append(query.Fields, "M_offsets")

	/** associations of ScheduledTask **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var ScheduledTaskInfo []interface{}

	ScheduledTaskInfo = append(ScheduledTaskInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		ScheduledTaskInfo = append(ScheduledTaskInfo, this.GetParentPtr().GetUuid())
		ScheduledTaskInfo = append(ScheduledTaskInfo, this.GetParentLnk())
	} else {
		ScheduledTaskInfo = append(ScheduledTaskInfo, "")
		ScheduledTaskInfo = append(ScheduledTaskInfo, "")
	}

	/** members of Configuration **/
	ScheduledTaskInfo = append(ScheduledTaskInfo, this.object.M_id)

	/** members of ScheduledTask **/
	ScheduledTaskInfo = append(ScheduledTaskInfo, this.object.M_isActive)
	ScheduledTaskInfo = append(ScheduledTaskInfo, this.object.M_script)
	ScheduledTaskInfo = append(ScheduledTaskInfo, this.object.M_startTime)
	ScheduledTaskInfo = append(ScheduledTaskInfo, this.object.M_expirationTime)
	ScheduledTaskInfo = append(ScheduledTaskInfo, this.object.M_frequency)

	/** Save frequencyType type FrequencyType **/
	if this.object.M_frequencyType == Config.FrequencyType_ONCE {
		ScheduledTaskInfo = append(ScheduledTaskInfo, 0)
	} else if this.object.M_frequencyType == Config.FrequencyType_DAILY {
		ScheduledTaskInfo = append(ScheduledTaskInfo, 1)
	} else if this.object.M_frequencyType == Config.FrequencyType_WEEKELY {
		ScheduledTaskInfo = append(ScheduledTaskInfo, 2)
	} else if this.object.M_frequencyType == Config.FrequencyType_MONTHLY {
		ScheduledTaskInfo = append(ScheduledTaskInfo, 3)
	} else {
		ScheduledTaskInfo = append(ScheduledTaskInfo, 0)
	}
	ScheduledTaskInfo = append(ScheduledTaskInfo, this.object.M_offsets)

	/** associations of ScheduledTask **/

	/** Save parent type Configurations **/
	if len(this.object.M_parentPtr) > 0 {
		ScheduledTaskInfo = append(ScheduledTaskInfo, this.object.M_parentPtr)
	} else {
		ScheduledTaskInfo = append(ScheduledTaskInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), ScheduledTaskInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), ScheduledTaskInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_ScheduledTaskEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_ScheduledTaskEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.ScheduledTask"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ScheduledTask **/
	query.Fields = append(query.Fields, "M_isActive")
	query.Fields = append(query.Fields, "M_script")
	query.Fields = append(query.Fields, "M_startTime")
	query.Fields = append(query.Fields, "M_expirationTime")
	query.Fields = append(query.Fields, "M_frequency")
	query.Fields = append(query.Fields, "M_frequencyType")
	query.Fields = append(query.Fields, "M_offsets")

	/** associations of ScheduledTask **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ScheduledTask...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.ScheduledTask)
		this.object.TYPENAME = "Config.ScheduledTask"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of ScheduledTask **/

		/** isActive **/
		if results[0][4] != nil {
			this.object.M_isActive = results[0][4].(bool)
		}

		/** script **/
		if results[0][5] != nil {
			this.object.M_script = results[0][5].(string)
		}

		/** startTime **/
		if results[0][6] != nil {
			this.object.M_startTime = results[0][6].(int64)
		}

		/** expirationTime **/
		if results[0][7] != nil {
			this.object.M_expirationTime = results[0][7].(int64)
		}

		/** frequency **/
		if results[0][8] != nil {
			this.object.M_frequency = results[0][8].(int)
		}

		/** frequencyType **/
		if results[0][9] != nil {
			enumIndex := results[0][9].(int)
			if enumIndex == 0 {
				this.object.M_frequencyType = Config.FrequencyType_ONCE
			} else if enumIndex == 1 {
				this.object.M_frequencyType = Config.FrequencyType_DAILY
			} else if enumIndex == 2 {
				this.object.M_frequencyType = Config.FrequencyType_WEEKELY
			} else if enumIndex == 3 {
				this.object.M_frequencyType = Config.FrequencyType_MONTHLY
			}
		}

		/** offsets **/
		if results[0][10] != nil {
			this.object.M_offsets = append(this.object.M_offsets, results[0][10].([]int)...)
		}

		/** associations of ScheduledTask **/

		/** parentPtr **/
		if results[0][11] != nil {
			id := results[0][11].(string)
			if len(id) > 0 {
				refTypeName := "Config.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigScheduledTaskEntityFromObject(object *Config.ScheduledTask) *Config_ScheduledTaskEntity {
	return this.NewConfigScheduledTaskEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_ScheduledTaskEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigScheduledTaskExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.ScheduledTask"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_ScheduledTaskEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_ScheduledTaskEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			ApplicationConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_ApplicationConfigurationEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.ApplicationConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewConfigApplicationConfigurationEntity(parentUuid string, objectId string, object interface{}) *Config_ApplicationConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigApplicationConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.ApplicationConfiguration).TYPENAME = "Config.ApplicationConfiguration"
		object.(*Config.ApplicationConfiguration).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.ApplicationConfiguration", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.ApplicationConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.ApplicationConfiguration).UUID
			}
			return val.(*Config_ApplicationConfigurationEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_ApplicationConfigurationEntity)
	if object == nil {
		entity.object = new(Config.ApplicationConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.ApplicationConfiguration)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.ApplicationConfiguration"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_ApplicationConfigurationEntity) GetTypeName() string {
	return "Config.ApplicationConfiguration"
}
func (this *Config_ApplicationConfigurationEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_ApplicationConfigurationEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_ApplicationConfigurationEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_ApplicationConfigurationEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_ApplicationConfigurationEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_ApplicationConfigurationEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_ApplicationConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_ApplicationConfigurationEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_ApplicationConfigurationEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_ApplicationConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_ApplicationConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_ApplicationConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_ApplicationConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_ApplicationConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_ApplicationConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_ApplicationConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_ApplicationConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_ApplicationConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_ApplicationConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_ApplicationConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_ApplicationConfigurationEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_ApplicationConfigurationEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_ApplicationConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.ApplicationConfiguration"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_ApplicationConfigurationEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_ApplicationConfigurationEntityPrototype() {

	var applicationConfigurationEntityProto EntityPrototype
	applicationConfigurationEntityProto.TypeName = "Config.ApplicationConfiguration"
	applicationConfigurationEntityProto.SuperTypeNames = append(applicationConfigurationEntityProto.SuperTypeNames, "Config.Configuration")
	applicationConfigurationEntityProto.Ids = append(applicationConfigurationEntityProto.Ids, "UUID")
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "UUID")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "xs.string")
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 0)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, false)
	applicationConfigurationEntityProto.FieldsDefaultValue = append(applicationConfigurationEntityProto.FieldsDefaultValue, "")
	applicationConfigurationEntityProto.Indexs = append(applicationConfigurationEntityProto.Indexs, "ParentUuid")
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "ParentUuid")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "xs.string")
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 1)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, false)
	applicationConfigurationEntityProto.FieldsDefaultValue = append(applicationConfigurationEntityProto.FieldsDefaultValue, "")
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "ParentLnk")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "xs.string")
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 2)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, false)
	applicationConfigurationEntityProto.FieldsDefaultValue = append(applicationConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of Configuration **/
	applicationConfigurationEntityProto.Ids = append(applicationConfigurationEntityProto.Ids, "M_id")
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 3)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, true)
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "M_id")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "xs.ID")
	applicationConfigurationEntityProto.FieldsDefaultValue = append(applicationConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of ApplicationConfiguration **/
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 4)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, true)
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "M_indexPage")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "xs.string")
	applicationConfigurationEntityProto.FieldsDefaultValue = append(applicationConfigurationEntityProto.FieldsDefaultValue, "")

	/** associations of ApplicationConfiguration **/
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 5)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, false)
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "M_parentPtr")
	applicationConfigurationEntityProto.FieldsDefaultValue = append(applicationConfigurationEntityProto.FieldsDefaultValue, "undefined")
	applicationConfigurationEntityProto.FieldsDefaultValue = append(applicationConfigurationEntityProto.FieldsDefaultValue, "undefined")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "Config.Configurations:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&applicationConfigurationEntityProto)

}

/** Create **/
func (this *Config_ApplicationConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.ApplicationConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ApplicationConfiguration **/
	query.Fields = append(query.Fields, "M_indexPage")

	/** associations of ApplicationConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var ApplicationConfigurationInfo []interface{}

	ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.GetParentPtr().GetUuid())
		ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.GetParentLnk())
	} else {
		ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, "")
		ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, "")
	}

	/** members of Configuration **/
	ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.object.M_id)

	/** members of ApplicationConfiguration **/
	ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.object.M_indexPage)

	/** associations of ApplicationConfiguration **/

	/** Save parent type Configurations **/
	if len(this.object.M_parentPtr) > 0 {
		ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.object.M_parentPtr)
	} else {
		ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), ApplicationConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), ApplicationConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_ApplicationConfigurationEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_ApplicationConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.ApplicationConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ApplicationConfiguration **/
	query.Fields = append(query.Fields, "M_indexPage")

	/** associations of ApplicationConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ApplicationConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.ApplicationConfiguration)
		this.object.TYPENAME = "Config.ApplicationConfiguration"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of ApplicationConfiguration **/

		/** indexPage **/
		if results[0][4] != nil {
			this.object.M_indexPage = results[0][4].(string)
		}

		/** associations of ApplicationConfiguration **/

		/** parentPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "Config.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigApplicationConfigurationEntityFromObject(object *Config.ApplicationConfiguration) *Config_ApplicationConfigurationEntity {
	return this.NewConfigApplicationConfigurationEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_ApplicationConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigApplicationConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.ApplicationConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_ApplicationConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_ApplicationConfigurationEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			ServerConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_ServerConfigurationEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.ServerConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewConfigServerConfigurationEntity(parentUuid string, objectId string, object interface{}) *Config_ServerConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigServerConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.ServerConfiguration).TYPENAME = "Config.ServerConfiguration"
		object.(*Config.ServerConfiguration).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.ServerConfiguration", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.ServerConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.ServerConfiguration).UUID
			}
			return val.(*Config_ServerConfigurationEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_ServerConfigurationEntity)
	if object == nil {
		entity.object = new(Config.ServerConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.ServerConfiguration)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.ServerConfiguration"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_ServerConfigurationEntity) GetTypeName() string {
	return "Config.ServerConfiguration"
}
func (this *Config_ServerConfigurationEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_ServerConfigurationEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_ServerConfigurationEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_ServerConfigurationEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_ServerConfigurationEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_ServerConfigurationEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_ServerConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_ServerConfigurationEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_ServerConfigurationEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_ServerConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_ServerConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_ServerConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_ServerConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_ServerConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_ServerConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_ServerConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_ServerConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_ServerConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_ServerConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_ServerConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_ServerConfigurationEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_ServerConfigurationEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_ServerConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.ServerConfiguration"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_ServerConfigurationEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_ServerConfigurationEntityPrototype() {

	var serverConfigurationEntityProto EntityPrototype
	serverConfigurationEntityProto.TypeName = "Config.ServerConfiguration"
	serverConfigurationEntityProto.SuperTypeNames = append(serverConfigurationEntityProto.SuperTypeNames, "Config.Configuration")
	serverConfigurationEntityProto.Ids = append(serverConfigurationEntityProto.Ids, "UUID")
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "UUID")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 0)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, false)
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")
	serverConfigurationEntityProto.Indexs = append(serverConfigurationEntityProto.Indexs, "ParentUuid")
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "ParentUuid")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 1)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, false)
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "ParentLnk")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 2)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, false)
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of Configuration **/
	serverConfigurationEntityProto.Ids = append(serverConfigurationEntityProto.Ids, "M_id")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 3)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_id")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.ID")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")

	/** members of ServerConfiguration **/
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 4)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_hostName")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 5)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_ipv4")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 6)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_serverPort")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.int")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "0")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 7)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_ws_serviceContainerPort")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.int")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "0")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 8)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_tcp_serviceContainerPort")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.int")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "0")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 9)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_applicationsPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 10)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_dataPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 11)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_scriptsPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 12)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_definitionsPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 13)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_schemasPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 14)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_tmpPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 15)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_binPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "")

	/** associations of ServerConfiguration **/
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 16)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, false)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_parentPtr")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "undefined")
	serverConfigurationEntityProto.FieldsDefaultValue = append(serverConfigurationEntityProto.FieldsDefaultValue, "undefined")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "Config.Configurations:Ref")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&serverConfigurationEntityProto)

}

/** Create **/
func (this *Config_ServerConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.ServerConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ServerConfiguration **/
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_serverPort")
	query.Fields = append(query.Fields, "M_ws_serviceContainerPort")
	query.Fields = append(query.Fields, "M_tcp_serviceContainerPort")
	query.Fields = append(query.Fields, "M_applicationsPath")
	query.Fields = append(query.Fields, "M_dataPath")
	query.Fields = append(query.Fields, "M_scriptsPath")
	query.Fields = append(query.Fields, "M_definitionsPath")
	query.Fields = append(query.Fields, "M_schemasPath")
	query.Fields = append(query.Fields, "M_tmpPath")
	query.Fields = append(query.Fields, "M_binPath")

	/** associations of ServerConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	var ServerConfigurationInfo []interface{}

	ServerConfigurationInfo = append(ServerConfigurationInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		ServerConfigurationInfo = append(ServerConfigurationInfo, this.GetParentPtr().GetUuid())
		ServerConfigurationInfo = append(ServerConfigurationInfo, this.GetParentLnk())
	} else {
		ServerConfigurationInfo = append(ServerConfigurationInfo, "")
		ServerConfigurationInfo = append(ServerConfigurationInfo, "")
	}

	/** members of Configuration **/
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_id)

	/** members of ServerConfiguration **/
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_hostName)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_ipv4)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_serverPort)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_ws_serviceContainerPort)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_tcp_serviceContainerPort)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_applicationsPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_dataPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_scriptsPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_definitionsPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_schemasPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_tmpPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_binPath)

	/** associations of ServerConfiguration **/

	/** Save parent type Configurations **/
	if len(this.object.M_parentPtr) > 0 {
		ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_parentPtr)
	} else {
		ServerConfigurationInfo = append(ServerConfigurationInfo, "")
	}
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), ServerConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), ServerConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_ServerConfigurationEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_ServerConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.ServerConfiguration"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ServerConfiguration **/
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_serverPort")
	query.Fields = append(query.Fields, "M_ws_serviceContainerPort")
	query.Fields = append(query.Fields, "M_tcp_serviceContainerPort")
	query.Fields = append(query.Fields, "M_applicationsPath")
	query.Fields = append(query.Fields, "M_dataPath")
	query.Fields = append(query.Fields, "M_scriptsPath")
	query.Fields = append(query.Fields, "M_definitionsPath")
	query.Fields = append(query.Fields, "M_schemasPath")
	query.Fields = append(query.Fields, "M_tmpPath")
	query.Fields = append(query.Fields, "M_binPath")

	/** associations of ServerConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ServerConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.ServerConfiguration)
		this.object.TYPENAME = "Config.ServerConfiguration"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of ServerConfiguration **/

		/** hostName **/
		if results[0][4] != nil {
			this.object.M_hostName = results[0][4].(string)
		}

		/** ipv4 **/
		if results[0][5] != nil {
			this.object.M_ipv4 = results[0][5].(string)
		}

		/** serverPort **/
		if results[0][6] != nil {
			this.object.M_serverPort = results[0][6].(int)
		}

		/** ws_serviceContainerPort **/
		if results[0][7] != nil {
			this.object.M_ws_serviceContainerPort = results[0][7].(int)
		}

		/** tcp_serviceContainerPort **/
		if results[0][8] != nil {
			this.object.M_tcp_serviceContainerPort = results[0][8].(int)
		}

		/** applicationsPath **/
		if results[0][9] != nil {
			this.object.M_applicationsPath = results[0][9].(string)
		}

		/** dataPath **/
		if results[0][10] != nil {
			this.object.M_dataPath = results[0][10].(string)
		}

		/** scriptsPath **/
		if results[0][11] != nil {
			this.object.M_scriptsPath = results[0][11].(string)
		}

		/** definitionsPath **/
		if results[0][12] != nil {
			this.object.M_definitionsPath = results[0][12].(string)
		}

		/** schemasPath **/
		if results[0][13] != nil {
			this.object.M_schemasPath = results[0][13].(string)
		}

		/** tmpPath **/
		if results[0][14] != nil {
			this.object.M_tmpPath = results[0][14].(string)
		}

		/** binPath **/
		if results[0][15] != nil {
			this.object.M_binPath = results[0][15].(string)
		}

		/** associations of ServerConfiguration **/

		/** parentPtr **/
		if results[0][16] != nil {
			id := results[0][16].(string)
			if len(id) > 0 {
				refTypeName := "Config.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigServerConfigurationEntityFromObject(object *Config.ServerConfiguration) *Config_ServerConfigurationEntity {
	return this.NewConfigServerConfigurationEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_ServerConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigServerConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.ServerConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_ServerConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_ServerConfigurationEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Configurations
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type Config_ConfigurationsEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *Config.Configurations
}

/** Constructor function **/
func (this *EntityManager) NewConfigConfigurationsEntity(parentUuid string, objectId string, object interface{}) *Config_ConfigurationsEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = ConfigConfigurationsExists(objectId)
		}
	}
	if object != nil {
		object.(*Config.Configurations).TYPENAME = "Config.Configurations"
		object.(*Config.Configurations).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("Config.Configurations", "Config")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*Config.Configurations).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*Config.Configurations).UUID
			}
			return val.(*Config_ConfigurationsEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(Config_ConfigurationsEntity)
	if object == nil {
		entity.object = new(Config.Configurations)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*Config.Configurations)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "Config.Configurations"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *Config_ConfigurationsEntity) GetTypeName() string {
	return "Config.Configurations"
}
func (this *Config_ConfigurationsEntity) GetUuid() string {
	return this.object.UUID
}
func (this *Config_ConfigurationsEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *Config_ConfigurationsEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *Config_ConfigurationsEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *Config_ConfigurationsEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *Config_ConfigurationsEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *Config_ConfigurationsEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *Config_ConfigurationsEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *Config_ConfigurationsEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *Config_ConfigurationsEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *Config_ConfigurationsEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *Config_ConfigurationsEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *Config_ConfigurationsEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *Config_ConfigurationsEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *Config_ConfigurationsEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *Config_ConfigurationsEntity) GetObject() interface{} {
	return this.object
}

func (this *Config_ConfigurationsEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *Config_ConfigurationsEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *Config_ConfigurationsEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *Config_ConfigurationsEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *Config_ConfigurationsEntity) IsLazy() bool {
	return this.lazy
}

func (this *Config_ConfigurationsEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *Config_ConfigurationsEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "Config.Configurations"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *Config_ConfigurationsEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_Config_ConfigurationsEntityPrototype() {

	var configurationsEntityProto EntityPrototype
	configurationsEntityProto.TypeName = "Config.Configurations"
	configurationsEntityProto.Ids = append(configurationsEntityProto.Ids, "UUID")
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "UUID")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.string")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 0)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, false)
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "")
	configurationsEntityProto.Indexs = append(configurationsEntityProto.Indexs, "ParentUuid")
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "ParentUuid")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.string")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 1)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, false)
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "")
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "ParentLnk")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.string")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 2)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, false)
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "")

	/** members of Configurations **/
	configurationsEntityProto.Ids = append(configurationsEntityProto.Ids, "M_id")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 3)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_id")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.ID")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 4)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_name")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.string")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 5)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_version")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.string")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 6)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_serverConfig")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "undefined")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "Config.ServerConfiguration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 7)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_oauth2Configuration")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "undefined")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "Config.OAuth2Configuration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 8)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_serviceConfigs")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "[]")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]Config.ServiceConfiguration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 9)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_dataStoreConfigs")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "[]")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]Config.DataStoreConfiguration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 10)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_smtpConfigs")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "[]")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]Config.SmtpConfiguration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 11)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_ldapConfigs")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "[]")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]Config.LdapConfiguration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 12)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_applicationConfigs")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "[]")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]Config.ApplicationConfiguration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 13)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_scheduledTasks")
	configurationsEntityProto.FieldsDefaultValue = append(configurationsEntityProto.FieldsDefaultValue, "[]")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]Config.ScheduledTask")

	store := GetServer().GetDataManager().getDataStore(ConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&configurationsEntityProto)

}

/** Create **/
func (this *Config_ConfigurationsEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "Config.Configurations"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configurations **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_version")
	query.Fields = append(query.Fields, "M_serverConfig")
	query.Fields = append(query.Fields, "M_oauth2Configuration")
	query.Fields = append(query.Fields, "M_serviceConfigs")
	query.Fields = append(query.Fields, "M_dataStoreConfigs")
	query.Fields = append(query.Fields, "M_smtpConfigs")
	query.Fields = append(query.Fields, "M_ldapConfigs")
	query.Fields = append(query.Fields, "M_applicationConfigs")
	query.Fields = append(query.Fields, "M_scheduledTasks")

	var ConfigurationsInfo []interface{}

	ConfigurationsInfo = append(ConfigurationsInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		ConfigurationsInfo = append(ConfigurationsInfo, this.GetParentPtr().GetUuid())
		ConfigurationsInfo = append(ConfigurationsInfo, this.GetParentLnk())
	} else {
		ConfigurationsInfo = append(ConfigurationsInfo, "")
		ConfigurationsInfo = append(ConfigurationsInfo, "")
	}

	/** members of Configurations **/
	ConfigurationsInfo = append(ConfigurationsInfo, this.object.M_id)
	ConfigurationsInfo = append(ConfigurationsInfo, this.object.M_name)
	ConfigurationsInfo = append(ConfigurationsInfo, this.object.M_version)

	/** Save serverConfig type ServerConfiguration **/
	lazy_serverConfig := this.lazyMap["M_serverConfig"] != nil && this.object.M_serverConfig == nil
	if this.object.M_serverConfig != nil && !lazy_serverConfig {
		serverConfigEntity := GetServer().GetEntityManager().NewConfigServerConfigurationEntity(this.GetUuid(), this.object.M_serverConfig.UUID, this.object.M_serverConfig)
		ConfigurationsInfo = append(ConfigurationsInfo, serverConfigEntity.GetUuid())
		serverConfigEntity.AppendReferenced("serverConfig", this)
		this.AppendChild("serverConfig", serverConfigEntity)
		if serverConfigEntity.NeedSave() {
			serverConfigEntity.SaveEntity()
		}
	} else if lazy_serverConfig {
		ConfigurationsInfo = append(ConfigurationsInfo, this.lazyMap["M_serverConfig"].(string))
	} else if this.object.M_serverConfig == nil {
		ConfigurationsInfo = append(ConfigurationsInfo, "")
	}

	/** Save oauth2Configuration type OAuth2Configuration **/
	lazy_oauth2Configuration := this.lazyMap["M_oauth2Configuration"] != nil && this.object.M_oauth2Configuration == nil
	if this.object.M_oauth2Configuration != nil && !lazy_oauth2Configuration {
		oauth2ConfigurationEntity := GetServer().GetEntityManager().NewConfigOAuth2ConfigurationEntity(this.GetUuid(), this.object.M_oauth2Configuration.UUID, this.object.M_oauth2Configuration)
		ConfigurationsInfo = append(ConfigurationsInfo, oauth2ConfigurationEntity.GetUuid())
		oauth2ConfigurationEntity.AppendReferenced("oauth2Configuration", this)
		this.AppendChild("oauth2Configuration", oauth2ConfigurationEntity)
		if oauth2ConfigurationEntity.NeedSave() {
			oauth2ConfigurationEntity.SaveEntity()
		}
	} else if lazy_oauth2Configuration {
		ConfigurationsInfo = append(ConfigurationsInfo, this.lazyMap["M_oauth2Configuration"].(string))
	} else if this.object.M_oauth2Configuration == nil {
		ConfigurationsInfo = append(ConfigurationsInfo, "")
	}

	/** Save serviceConfigs type ServiceConfiguration **/
	serviceConfigsIds := make([]string, 0)
	lazy_serviceConfigs := this.lazyMap["M_serviceConfigs"] != nil && len(this.object.M_serviceConfigs) == 0
	if !lazy_serviceConfigs {
		for i := 0; i < len(this.object.M_serviceConfigs); i++ {
			serviceConfigsEntity := GetServer().GetEntityManager().NewConfigServiceConfigurationEntity(this.GetUuid(), this.object.M_serviceConfigs[i].UUID, this.object.M_serviceConfigs[i])
			serviceConfigsIds = append(serviceConfigsIds, serviceConfigsEntity.GetUuid())
			serviceConfigsEntity.AppendReferenced("serviceConfigs", this)
			this.AppendChild("serviceConfigs", serviceConfigsEntity)
			if serviceConfigsEntity.NeedSave() {
				serviceConfigsEntity.SaveEntity()
			}
		}
	} else {
		serviceConfigsIds = this.lazyMap["M_serviceConfigs"].([]string)
	}
	serviceConfigsStr, _ := json.Marshal(serviceConfigsIds)
	ConfigurationsInfo = append(ConfigurationsInfo, string(serviceConfigsStr))

	/** Save dataStoreConfigs type DataStoreConfiguration **/
	dataStoreConfigsIds := make([]string, 0)
	lazy_dataStoreConfigs := this.lazyMap["M_dataStoreConfigs"] != nil && len(this.object.M_dataStoreConfigs) == 0
	if !lazy_dataStoreConfigs {
		for i := 0; i < len(this.object.M_dataStoreConfigs); i++ {
			dataStoreConfigsEntity := GetServer().GetEntityManager().NewConfigDataStoreConfigurationEntity(this.GetUuid(), this.object.M_dataStoreConfigs[i].UUID, this.object.M_dataStoreConfigs[i])
			dataStoreConfigsIds = append(dataStoreConfigsIds, dataStoreConfigsEntity.GetUuid())
			dataStoreConfigsEntity.AppendReferenced("dataStoreConfigs", this)
			this.AppendChild("dataStoreConfigs", dataStoreConfigsEntity)
			if dataStoreConfigsEntity.NeedSave() {
				dataStoreConfigsEntity.SaveEntity()
			}
		}
	} else {
		dataStoreConfigsIds = this.lazyMap["M_dataStoreConfigs"].([]string)
	}
	dataStoreConfigsStr, _ := json.Marshal(dataStoreConfigsIds)
	ConfigurationsInfo = append(ConfigurationsInfo, string(dataStoreConfigsStr))

	/** Save smtpConfigs type SmtpConfiguration **/
	smtpConfigsIds := make([]string, 0)
	lazy_smtpConfigs := this.lazyMap["M_smtpConfigs"] != nil && len(this.object.M_smtpConfigs) == 0
	if !lazy_smtpConfigs {
		for i := 0; i < len(this.object.M_smtpConfigs); i++ {
			smtpConfigsEntity := GetServer().GetEntityManager().NewConfigSmtpConfigurationEntity(this.GetUuid(), this.object.M_smtpConfigs[i].UUID, this.object.M_smtpConfigs[i])
			smtpConfigsIds = append(smtpConfigsIds, smtpConfigsEntity.GetUuid())
			smtpConfigsEntity.AppendReferenced("smtpConfigs", this)
			this.AppendChild("smtpConfigs", smtpConfigsEntity)
			if smtpConfigsEntity.NeedSave() {
				smtpConfigsEntity.SaveEntity()
			}
		}
	} else {
		smtpConfigsIds = this.lazyMap["M_smtpConfigs"].([]string)
	}
	smtpConfigsStr, _ := json.Marshal(smtpConfigsIds)
	ConfigurationsInfo = append(ConfigurationsInfo, string(smtpConfigsStr))

	/** Save ldapConfigs type LdapConfiguration **/
	ldapConfigsIds := make([]string, 0)
	lazy_ldapConfigs := this.lazyMap["M_ldapConfigs"] != nil && len(this.object.M_ldapConfigs) == 0
	if !lazy_ldapConfigs {
		for i := 0; i < len(this.object.M_ldapConfigs); i++ {
			ldapConfigsEntity := GetServer().GetEntityManager().NewConfigLdapConfigurationEntity(this.GetUuid(), this.object.M_ldapConfigs[i].UUID, this.object.M_ldapConfigs[i])
			ldapConfigsIds = append(ldapConfigsIds, ldapConfigsEntity.GetUuid())
			ldapConfigsEntity.AppendReferenced("ldapConfigs", this)
			this.AppendChild("ldapConfigs", ldapConfigsEntity)
			if ldapConfigsEntity.NeedSave() {
				ldapConfigsEntity.SaveEntity()
			}
		}
	} else {
		ldapConfigsIds = this.lazyMap["M_ldapConfigs"].([]string)
	}
	ldapConfigsStr, _ := json.Marshal(ldapConfigsIds)
	ConfigurationsInfo = append(ConfigurationsInfo, string(ldapConfigsStr))

	/** Save applicationConfigs type ApplicationConfiguration **/
	applicationConfigsIds := make([]string, 0)
	lazy_applicationConfigs := this.lazyMap["M_applicationConfigs"] != nil && len(this.object.M_applicationConfigs) == 0
	if !lazy_applicationConfigs {
		for i := 0; i < len(this.object.M_applicationConfigs); i++ {
			applicationConfigsEntity := GetServer().GetEntityManager().NewConfigApplicationConfigurationEntity(this.GetUuid(), this.object.M_applicationConfigs[i].UUID, this.object.M_applicationConfigs[i])
			applicationConfigsIds = append(applicationConfigsIds, applicationConfigsEntity.GetUuid())
			applicationConfigsEntity.AppendReferenced("applicationConfigs", this)
			this.AppendChild("applicationConfigs", applicationConfigsEntity)
			if applicationConfigsEntity.NeedSave() {
				applicationConfigsEntity.SaveEntity()
			}
		}
	} else {
		applicationConfigsIds = this.lazyMap["M_applicationConfigs"].([]string)
	}
	applicationConfigsStr, _ := json.Marshal(applicationConfigsIds)
	ConfigurationsInfo = append(ConfigurationsInfo, string(applicationConfigsStr))

	/** Save scheduledTasks type ScheduledTask **/
	scheduledTasksIds := make([]string, 0)
	lazy_scheduledTasks := this.lazyMap["M_scheduledTasks"] != nil && len(this.object.M_scheduledTasks) == 0
	if !lazy_scheduledTasks {
		for i := 0; i < len(this.object.M_scheduledTasks); i++ {
			scheduledTasksEntity := GetServer().GetEntityManager().NewConfigScheduledTaskEntity(this.GetUuid(), this.object.M_scheduledTasks[i].UUID, this.object.M_scheduledTasks[i])
			scheduledTasksIds = append(scheduledTasksIds, scheduledTasksEntity.GetUuid())
			scheduledTasksEntity.AppendReferenced("scheduledTasks", this)
			this.AppendChild("scheduledTasks", scheduledTasksEntity)
			if scheduledTasksEntity.NeedSave() {
				scheduledTasksEntity.SaveEntity()
			}
		}
	} else {
		scheduledTasksIds = this.lazyMap["M_scheduledTasks"].([]string)
	}
	scheduledTasksStr, _ := json.Marshal(scheduledTasksIds)
	ConfigurationsInfo = append(ConfigurationsInfo, string(scheduledTasksStr))
	eventData := make([]*MessageData, 2)

	msgData0 := new(MessageData)
	msgData0.TYPENAME = "Server.MessageData"
	msgData0.Name = "entity"
	msgData0.Value = this.GetObject()
	eventData[0] = msgData0
	msgData1 := new(MessageData)
	msgData1.TYPENAME = "Server.MessageData"
	msgData1.Name = "prototype"
	msgData1.Value = this.GetPrototype()
	eventData[1] = msgData1
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(ConfigDB, string(queryStr), ConfigurationsInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(ConfigDB, string(queryStr), ConfigurationsInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *Config_ConfigurationsEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*Config_ConfigurationsEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "Config.Configurations"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Configurations **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_version")
	query.Fields = append(query.Fields, "M_serverConfig")
	query.Fields = append(query.Fields, "M_oauth2Configuration")
	query.Fields = append(query.Fields, "M_serviceConfigs")
	query.Fields = append(query.Fields, "M_dataStoreConfigs")
	query.Fields = append(query.Fields, "M_smtpConfigs")
	query.Fields = append(query.Fields, "M_ldapConfigs")
	query.Fields = append(query.Fields, "M_applicationConfigs")
	query.Fields = append(query.Fields, "M_scheduledTasks")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Configurations...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(Config.Configurations)
		this.object.TYPENAME = "Config.Configurations"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Configurations **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** name **/
		if results[0][4] != nil {
			this.object.M_name = results[0][4].(string)
		}

		/** version **/
		if results[0][5] != nil {
			this.object.M_version = results[0][5].(string)
		}

		/** serverConfig **/
		if results[0][6] != nil {
			uuid := results[0][6].(string)
			if len(uuid) > 0 {
				var serverConfigEntity *Config_ServerConfigurationEntity
				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
					serverConfigEntity = instance.(*Config_ServerConfigurationEntity)
				} else {
					serverConfigEntity = GetServer().GetEntityManager().NewConfigServerConfigurationEntity(this.GetUuid(), uuid, nil)
					serverConfigEntity.InitEntity(uuid, lazy)
					GetServer().GetEntityManager().insert(serverConfigEntity)
				}
				serverConfigEntity.AppendReferenced("serverConfig", this)
				this.AppendChild("serverConfig", serverConfigEntity)
			}
		}

		/** oauth2Configuration **/
		if results[0][7] != nil {
			uuid := results[0][7].(string)
			if len(uuid) > 0 {
				var oauth2ConfigurationEntity *Config_OAuth2ConfigurationEntity
				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
					oauth2ConfigurationEntity = instance.(*Config_OAuth2ConfigurationEntity)
				} else {
					oauth2ConfigurationEntity = GetServer().GetEntityManager().NewConfigOAuth2ConfigurationEntity(this.GetUuid(), uuid, nil)
					oauth2ConfigurationEntity.InitEntity(uuid, lazy)
					GetServer().GetEntityManager().insert(oauth2ConfigurationEntity)
				}
				oauth2ConfigurationEntity.AppendReferenced("oauth2Configuration", this)
				this.AppendChild("oauth2Configuration", oauth2ConfigurationEntity)
			}
		}

		/** serviceConfigs **/
		if results[0][8] != nil {
			uuidsStr := results[0][8].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var serviceConfigsEntity *Config_ServiceConfigurationEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							serviceConfigsEntity = instance.(*Config_ServiceConfigurationEntity)
						} else {
							serviceConfigsEntity = GetServer().GetEntityManager().NewConfigServiceConfigurationEntity(this.GetUuid(), uuids[i], nil)
							serviceConfigsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(serviceConfigsEntity)
						}
						serviceConfigsEntity.AppendReferenced("serviceConfigs", this)
						this.AppendChild("serviceConfigs", serviceConfigsEntity)
					}
				} else {
					this.lazyMap["M_serviceConfigs"] = uuids
				}
			}
		}

		/** dataStoreConfigs **/
		if results[0][9] != nil {
			uuidsStr := results[0][9].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var dataStoreConfigsEntity *Config_DataStoreConfigurationEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							dataStoreConfigsEntity = instance.(*Config_DataStoreConfigurationEntity)
						} else {
							dataStoreConfigsEntity = GetServer().GetEntityManager().NewConfigDataStoreConfigurationEntity(this.GetUuid(), uuids[i], nil)
							dataStoreConfigsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(dataStoreConfigsEntity)
						}
						dataStoreConfigsEntity.AppendReferenced("dataStoreConfigs", this)
						this.AppendChild("dataStoreConfigs", dataStoreConfigsEntity)
					}
				} else {
					this.lazyMap["M_dataStoreConfigs"] = uuids
				}
			}
		}

		/** smtpConfigs **/
		if results[0][10] != nil {
			uuidsStr := results[0][10].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var smtpConfigsEntity *Config_SmtpConfigurationEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							smtpConfigsEntity = instance.(*Config_SmtpConfigurationEntity)
						} else {
							smtpConfigsEntity = GetServer().GetEntityManager().NewConfigSmtpConfigurationEntity(this.GetUuid(), uuids[i], nil)
							smtpConfigsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(smtpConfigsEntity)
						}
						smtpConfigsEntity.AppendReferenced("smtpConfigs", this)
						this.AppendChild("smtpConfigs", smtpConfigsEntity)
					}
				} else {
					this.lazyMap["M_smtpConfigs"] = uuids
				}
			}
		}

		/** ldapConfigs **/
		if results[0][11] != nil {
			uuidsStr := results[0][11].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var ldapConfigsEntity *Config_LdapConfigurationEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							ldapConfigsEntity = instance.(*Config_LdapConfigurationEntity)
						} else {
							ldapConfigsEntity = GetServer().GetEntityManager().NewConfigLdapConfigurationEntity(this.GetUuid(), uuids[i], nil)
							ldapConfigsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(ldapConfigsEntity)
						}
						ldapConfigsEntity.AppendReferenced("ldapConfigs", this)
						this.AppendChild("ldapConfigs", ldapConfigsEntity)
					}
				} else {
					this.lazyMap["M_ldapConfigs"] = uuids
				}
			}
		}

		/** applicationConfigs **/
		if results[0][12] != nil {
			uuidsStr := results[0][12].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var applicationConfigsEntity *Config_ApplicationConfigurationEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							applicationConfigsEntity = instance.(*Config_ApplicationConfigurationEntity)
						} else {
							applicationConfigsEntity = GetServer().GetEntityManager().NewConfigApplicationConfigurationEntity(this.GetUuid(), uuids[i], nil)
							applicationConfigsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(applicationConfigsEntity)
						}
						applicationConfigsEntity.AppendReferenced("applicationConfigs", this)
						this.AppendChild("applicationConfigs", applicationConfigsEntity)
					}
				} else {
					this.lazyMap["M_applicationConfigs"] = uuids
				}
			}
		}

		/** scheduledTasks **/
		if results[0][13] != nil {
			uuidsStr := results[0][13].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var scheduledTasksEntity *Config_ScheduledTaskEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							scheduledTasksEntity = instance.(*Config_ScheduledTaskEntity)
						} else {
							scheduledTasksEntity = GetServer().GetEntityManager().NewConfigScheduledTaskEntity(this.GetUuid(), uuids[i], nil)
							scheduledTasksEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(scheduledTasksEntity)
						}
						scheduledTasksEntity.AppendReferenced("scheduledTasks", this)
						this.AppendChild("scheduledTasks", scheduledTasksEntity)
					}
				} else {
					this.lazyMap["M_scheduledTasks"] = uuids
				}
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewConfigConfigurationsEntityFromObject(object *Config.Configurations) *Config_ConfigurationsEntity {
	return this.NewConfigConfigurationsEntity("", object.UUID, object)
}

/** Delete **/
func (this *Config_ConfigurationsEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func ConfigConfigurationsExists(val string) string {
	var query EntityQuery
	query.TypeName = "Config.Configurations"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(ConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *Config_ConfigurationsEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *Config_ConfigurationsEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

/** Register the entity to the dynamic typing system. **/
func (this *EntityManager) registerConfigObjects() {
	Utility.RegisterType((*Config.SmtpConfiguration)(nil))
	Utility.RegisterType((*Config.DataStoreConfiguration)(nil))
	Utility.RegisterType((*Config.LdapConfiguration)(nil))
	Utility.RegisterType((*Config.OAuth2Client)(nil))
	Utility.RegisterType((*Config.OAuth2Authorize)(nil))
	Utility.RegisterType((*Config.OAuth2IdToken)(nil))
	Utility.RegisterType((*Config.OAuth2Access)(nil))
	Utility.RegisterType((*Config.OAuth2Refresh)(nil))
	Utility.RegisterType((*Config.OAuth2Expires)(nil))
	Utility.RegisterType((*Config.OAuth2Configuration)(nil))
	Utility.RegisterType((*Config.ServiceConfiguration)(nil))
	Utility.RegisterType((*Config.ScheduledTask)(nil))
	Utility.RegisterType((*Config.ApplicationConfiguration)(nil))
	Utility.RegisterType((*Config.ServerConfiguration)(nil))
	Utility.RegisterType((*Config.Configurations)(nil))
}

/** Create entity prototypes contain in a package **/
func (this *EntityManager) createConfigPrototypes() {
	this.create_Config_ConfigurationEntityPrototype()
	this.create_Config_SmtpConfigurationEntityPrototype()
	this.create_Config_DataStoreConfigurationEntityPrototype()
	this.create_Config_LdapConfigurationEntityPrototype()
	this.create_Config_OAuth2ClientEntityPrototype()
	this.create_Config_OAuth2AuthorizeEntityPrototype()
	this.create_Config_OAuth2IdTokenEntityPrototype()
	this.create_Config_OAuth2AccessEntityPrototype()
	this.create_Config_OAuth2RefreshEntityPrototype()
	this.create_Config_OAuth2ExpiresEntityPrototype()
	this.create_Config_OAuth2ConfigurationEntityPrototype()
	this.create_Config_ServiceConfigurationEntityPrototype()
	this.create_Config_ScheduledTaskEntityPrototype()
	this.create_Config_ApplicationConfigurationEntityPrototype()
	this.create_Config_ServerConfigurationEntityPrototype()
	this.create_Config_ConfigurationsEntityPrototype()
}
