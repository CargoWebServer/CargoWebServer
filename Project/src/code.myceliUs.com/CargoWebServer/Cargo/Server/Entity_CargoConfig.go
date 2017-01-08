package Server

import (
	"encoding/json"

	"code.google.com/p/go-uuid/uuid"
	"code.myceliUs.com/CargoWebServer/Cargo/Config/CargoConfig"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
	//	"log"
	"strings"
)

/** Entity Prototype creation **/
func (this *EntityManager) Create_CargoConfig_ConfigurationEntityPrototype() {

	var configurationEntityProto EntityPrototype
	configurationEntityProto.TypeName = "CargoConfig.Configuration"
	configurationEntityProto.IsAbstract = true
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "CargoConfig.SmtpConfiguration")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "CargoConfig.DataStoreConfiguration")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "CargoConfig.LdapConfiguration")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "CargoConfig.ApplicationConfiguration")
	configurationEntityProto.SubstitutionGroup = append(configurationEntityProto.SubstitutionGroup, "CargoConfig.ServerConfiguration")
	configurationEntityProto.Ids = append(configurationEntityProto.Ids, "uuid")
	configurationEntityProto.Fields = append(configurationEntityProto.Fields, "uuid")
	configurationEntityProto.FieldsType = append(configurationEntityProto.FieldsType, "xs.string")
	configurationEntityProto.FieldsOrder = append(configurationEntityProto.FieldsOrder, 0)
	configurationEntityProto.FieldsVisibility = append(configurationEntityProto.FieldsVisibility, false)
	configurationEntityProto.Indexs = append(configurationEntityProto.Indexs, "parentUuid")
	configurationEntityProto.Fields = append(configurationEntityProto.Fields, "parentUuid")
	configurationEntityProto.FieldsType = append(configurationEntityProto.FieldsType, "xs.string")
	configurationEntityProto.FieldsOrder = append(configurationEntityProto.FieldsOrder, 1)
	configurationEntityProto.FieldsVisibility = append(configurationEntityProto.FieldsVisibility, false)

	/** members of Configuration **/
	configurationEntityProto.Ids = append(configurationEntityProto.Ids, "M_id")
	configurationEntityProto.FieldsOrder = append(configurationEntityProto.FieldsOrder, 2)
	configurationEntityProto.FieldsVisibility = append(configurationEntityProto.FieldsVisibility, true)
	configurationEntityProto.Fields = append(configurationEntityProto.Fields, "M_id")
	configurationEntityProto.FieldsType = append(configurationEntityProto.FieldsType, "xs.ID")
	configurationEntityProto.Fields = append(configurationEntityProto.Fields, "childsUuid")
	configurationEntityProto.FieldsType = append(configurationEntityProto.FieldsType, "[]xs.string")
	configurationEntityProto.FieldsOrder = append(configurationEntityProto.FieldsOrder, 3)
	configurationEntityProto.FieldsVisibility = append(configurationEntityProto.FieldsVisibility, false)

	configurationEntityProto.Fields = append(configurationEntityProto.Fields, "referenced")
	configurationEntityProto.FieldsType = append(configurationEntityProto.FieldsType, "[]EntityRef")
	configurationEntityProto.FieldsOrder = append(configurationEntityProto.FieldsOrder, 4)
	configurationEntityProto.FieldsVisibility = append(configurationEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&configurationEntityProto)

}

////////////////////////////////////////////////////////////////////////////////
//              			SmtpConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoConfig_SmtpConfigurationEntity struct {
	/** not the object id, except for the definition **/
	uuid           string
	parentPtr      Entity
	parentUuid     string
	childsPtr      []Entity
	childsUuid     []string
	referencesUuid []string
	referencesPtr  []Entity
	prototype      *EntityPrototype
	referenced     []EntityRef
	object         *CargoConfig.SmtpConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewCargoConfigSmtpConfigurationEntity(objectId string, object interface{}) *CargoConfig_SmtpConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoConfigSmtpConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoConfig.SmtpConfiguration).TYPENAME = "CargoConfig.SmtpConfiguration"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoConfig.SmtpConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*CargoConfig_SmtpConfigurationEntity)
		}
	} else {
		uuidStr = "CargoConfig.SmtpConfiguration%" + uuid.NewRandom().String()
	}
	entity := new(CargoConfig_SmtpConfigurationEntity)
	if object == nil {
		entity.object = new(CargoConfig.SmtpConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoConfig.SmtpConfiguration)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoConfig.SmtpConfiguration"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoConfig.SmtpConfiguration", "CargoConfig")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoConfig_SmtpConfigurationEntity) GetTypeName() string {
	return "CargoConfig.SmtpConfiguration"
}
func (this *CargoConfig_SmtpConfigurationEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoConfig_SmtpConfigurationEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoConfig_SmtpConfigurationEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoConfig_SmtpConfigurationEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoConfig_SmtpConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoConfig_SmtpConfigurationEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoConfig_SmtpConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	refsPtr := make([]Entity, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsPtr = append(refsPtr, reference)
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	this.SetReferencesPtr(refsPtr)

	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_SmtpConfigurationEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoConfig_SmtpConfigurationEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoConfig_SmtpConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoConfig_SmtpConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoConfig_SmtpConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		}
	}

	this.childsUuid = childsUuid
	params := make([]interface{}, 1)
	childsPtr := make([]Entity, 0)
	for i := 0; i < len(this.GetChildsPtr()); i++ {
		if this.GetChildsPtr()[i].GetUuid() != uuid {
			childsPtr = append(childsPtr, this.GetChildsPtr()[i])
		} else {
			params[0] = this.GetChildsPtr()[i].GetObject()
		}
	}
	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_SmtpConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoConfig_SmtpConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoConfig_SmtpConfigurationEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoConfig_SmtpConfigurationEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoConfig_SmtpConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoConfig_SmtpConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoConfig_SmtpConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoConfig_SmtpConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoConfig_SmtpConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoConfig_SmtpConfigurationEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *CargoConfig_SmtpConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoConfig.SmtpConfiguration"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoConfig_SmtpConfigurationEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) Create_CargoConfig_SmtpConfigurationEntityPrototype() {

	var smtpConfigurationEntityProto EntityPrototype
	smtpConfigurationEntityProto.TypeName = "CargoConfig.SmtpConfiguration"
	smtpConfigurationEntityProto.SuperTypeNames = append(smtpConfigurationEntityProto.SuperTypeNames, "CargoConfig.Configuration")
	smtpConfigurationEntityProto.Ids = append(smtpConfigurationEntityProto.Ids, "uuid")
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "uuid")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 0)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, false)
	smtpConfigurationEntityProto.Indexs = append(smtpConfigurationEntityProto.Indexs, "parentUuid")
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "parentUuid")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 1)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, false)

	/** members of Configuration **/
	smtpConfigurationEntityProto.Ids = append(smtpConfigurationEntityProto.Ids, "M_id")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 2)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_id")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.ID")

	/** members of SmtpConfiguration **/
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 3)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_hostName")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 4)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_ipv4")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 5)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_port")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.int")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 6)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_user")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 7)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, true)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_pwd")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "xs.string")

	/** associations of SmtpConfiguration **/
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 8)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, false)
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "M_parentPtr")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "CargoConfig.Configurations:Ref")
	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "childsUuid")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "[]xs.string")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 9)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, false)

	smtpConfigurationEntityProto.Fields = append(smtpConfigurationEntityProto.Fields, "referenced")
	smtpConfigurationEntityProto.FieldsType = append(smtpConfigurationEntityProto.FieldsType, "[]EntityRef")
	smtpConfigurationEntityProto.FieldsOrder = append(smtpConfigurationEntityProto.FieldsOrder, 10)
	smtpConfigurationEntityProto.FieldsVisibility = append(smtpConfigurationEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&smtpConfigurationEntityProto)

}

/** Create **/
func (this *CargoConfig_SmtpConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoConfig.SmtpConfiguration"

	var query EntityQuery
	query.TypeName = "CargoConfig.SmtpConfiguration"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of SmtpConfiguration **/
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_port")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")

	/** associations of SmtpConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var SmtpConfigurationInfo []interface{}

	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.GetUuid())
	if this.parentPtr != nil {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.parentPtr.GetUuid())
	} else {
		SmtpConfigurationInfo = append(SmtpConfigurationInfo, "")
	}

	/** members of Configuration **/
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_id)

	/** members of SmtpConfiguration **/
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_hostName)
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_ipv4)
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_port)
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_user)
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_pwd)

	/** associations of SmtpConfiguration **/

	/** Save parent type Configurations **/
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, this.object.M_parentPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	SmtpConfigurationInfo = append(SmtpConfigurationInfo, string(referencedStr))
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "uuid="+this.uuid)
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoConfigDB, string(queryStr), SmtpConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoConfigDB, string(queryStr), SmtpConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoConfig_SmtpConfigurationEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoConfig_SmtpConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoConfig.SmtpConfiguration"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of SmtpConfiguration **/
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_port")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")

	/** associations of SmtpConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of SmtpConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoConfig.SmtpConfiguration)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoConfig.SmtpConfiguration"

		this.parentUuid = results[0][1].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of SmtpConfiguration **/

		/** hostName **/
		if results[0][3] != nil {
			this.object.M_hostName = results[0][3].(string)
		}

		/** ipv4 **/
		if results[0][4] != nil {
			this.object.M_ipv4 = results[0][4].(string)
		}

		/** port **/
		if results[0][5] != nil {
			this.object.M_port = results[0][5].(int)
		}

		/** user **/
		if results[0][6] != nil {
			this.object.M_user = results[0][6].(string)
		}

		/** pwd **/
		if results[0][7] != nil {
			this.object.M_pwd = results[0][7].(string)
		}

		/** associations of SmtpConfiguration **/

		/** parentPtr **/
		if results[0][8] != nil {
			id := results[0][8].(string)
			if len(id) > 0 {
				refTypeName := "CargoConfig.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][9].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][10].(string)
		this.referenced = make([]EntityRef, 0)
		err = json.Unmarshal([]byte(referencedStr), &this.referenced)
		if err != nil {
			return err
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoConfigSmtpConfigurationEntityFromObject(object *CargoConfig.SmtpConfiguration) *CargoConfig_SmtpConfigurationEntity {
	return this.NewCargoConfigSmtpConfigurationEntity(object.UUID, object)
}

/** Delete **/
func (this *CargoConfig_SmtpConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoConfigSmtpConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoConfig.SmtpConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoConfig_SmtpConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
		this.childsPtr = append(this.childsPtr, child)
	} else {
		childsPtr := make([]Entity, 0)
		for i := 0; i < len(this.childsPtr); i++ {
			if this.childsPtr[i].GetUuid() != child.GetUuid() {
				childsPtr = append(childsPtr, this.childsPtr[i])
			}
		}
		childsPtr = append(childsPtr, child)
		this.SetChildsPtr(childsPtr)
	}
	// Set this as parent in the child
	child.SetParentPtr(this)

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
func (this *CargoConfig_SmtpConfigurationEntity) AppendReference(reference Entity) {

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
		this.referencesPtr = append(this.referencesPtr, reference)
	} else {
		// The reference must be update in that case.
		this.referencesPtr[index] = reference
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			DataStoreConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoConfig_DataStoreConfigurationEntity struct {
	/** not the object id, except for the definition **/
	uuid           string
	parentPtr      Entity
	parentUuid     string
	childsPtr      []Entity
	childsUuid     []string
	referencesUuid []string
	referencesPtr  []Entity
	prototype      *EntityPrototype
	referenced     []EntityRef
	object         *CargoConfig.DataStoreConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewCargoConfigDataStoreConfigurationEntity(objectId string, object interface{}) *CargoConfig_DataStoreConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoConfigDataStoreConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoConfig.DataStoreConfiguration).TYPENAME = "CargoConfig.DataStoreConfiguration"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoConfig.DataStoreConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*CargoConfig_DataStoreConfigurationEntity)
		}
	} else {
		uuidStr = "CargoConfig.DataStoreConfiguration%" + uuid.NewRandom().String()
	}
	entity := new(CargoConfig_DataStoreConfigurationEntity)
	if object == nil {
		entity.object = new(CargoConfig.DataStoreConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoConfig.DataStoreConfiguration)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoConfig.DataStoreConfiguration"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoConfig.DataStoreConfiguration", "CargoConfig")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoConfig_DataStoreConfigurationEntity) GetTypeName() string {
	return "CargoConfig.DataStoreConfiguration"
}
func (this *CargoConfig_DataStoreConfigurationEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoConfig_DataStoreConfigurationEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoConfig_DataStoreConfigurationEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoConfig_DataStoreConfigurationEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoConfig_DataStoreConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoConfig_DataStoreConfigurationEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoConfig_DataStoreConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	refsPtr := make([]Entity, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsPtr = append(refsPtr, reference)
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	this.SetReferencesPtr(refsPtr)

	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_DataStoreConfigurationEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoConfig_DataStoreConfigurationEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoConfig_DataStoreConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoConfig_DataStoreConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoConfig_DataStoreConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		}
	}

	this.childsUuid = childsUuid
	params := make([]interface{}, 1)
	childsPtr := make([]Entity, 0)
	for i := 0; i < len(this.GetChildsPtr()); i++ {
		if this.GetChildsPtr()[i].GetUuid() != uuid {
			childsPtr = append(childsPtr, this.GetChildsPtr()[i])
		} else {
			params[0] = this.GetChildsPtr()[i].GetObject()
		}
	}
	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_DataStoreConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoConfig_DataStoreConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoConfig_DataStoreConfigurationEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoConfig_DataStoreConfigurationEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoConfig_DataStoreConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoConfig_DataStoreConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoConfig_DataStoreConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoConfig_DataStoreConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoConfig_DataStoreConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoConfig_DataStoreConfigurationEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *CargoConfig_DataStoreConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoConfig.DataStoreConfiguration"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoConfig_DataStoreConfigurationEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) Create_CargoConfig_DataStoreConfigurationEntityPrototype() {

	var dataStoreConfigurationEntityProto EntityPrototype
	dataStoreConfigurationEntityProto.TypeName = "CargoConfig.DataStoreConfiguration"
	dataStoreConfigurationEntityProto.SuperTypeNames = append(dataStoreConfigurationEntityProto.SuperTypeNames, "CargoConfig.Configuration")
	dataStoreConfigurationEntityProto.Ids = append(dataStoreConfigurationEntityProto.Ids, "uuid")
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "uuid")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 0)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, false)
	dataStoreConfigurationEntityProto.Indexs = append(dataStoreConfigurationEntityProto.Indexs, "parentUuid")
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "parentUuid")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 1)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, false)

	/** members of Configuration **/
	dataStoreConfigurationEntityProto.Ids = append(dataStoreConfigurationEntityProto.Ids, "M_id")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 2)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_id")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.ID")

	/** members of DataStoreConfiguration **/
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 3)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_dataStoreType")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "enum:DataStoreType_SQL_STORE:DataStoreType_MEMORY_STORE:DataStoreType_KEY_VALUE_STORE:DataStoreType_GRAPH_STORE")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 4)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_dataStoreVendor")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "enum:DataStoreVendor_MYCELIUS:DataStoreVendor_MYSQL:DataStoreVendor_MSSQL:DataStoreVendor_ODBC:DataStoreVendor_KNOWLEDGEBASE")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 5)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_hostName")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 6)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_ipv4")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 7)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_user")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 8)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_pwd")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.string")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 9)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, true)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_port")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "xs.int")

	/** associations of DataStoreConfiguration **/
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 10)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, false)
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "M_parentPtr")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "CargoConfig.Configurations:Ref")
	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "childsUuid")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "[]xs.string")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 11)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, false)

	dataStoreConfigurationEntityProto.Fields = append(dataStoreConfigurationEntityProto.Fields, "referenced")
	dataStoreConfigurationEntityProto.FieldsType = append(dataStoreConfigurationEntityProto.FieldsType, "[]EntityRef")
	dataStoreConfigurationEntityProto.FieldsOrder = append(dataStoreConfigurationEntityProto.FieldsOrder, 12)
	dataStoreConfigurationEntityProto.FieldsVisibility = append(dataStoreConfigurationEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&dataStoreConfigurationEntityProto)

}

/** Create **/
func (this *CargoConfig_DataStoreConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoConfig.DataStoreConfiguration"

	var query EntityQuery
	query.TypeName = "CargoConfig.DataStoreConfiguration"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of DataStoreConfiguration **/
	query.Fields = append(query.Fields, "M_dataStoreType")
	query.Fields = append(query.Fields, "M_dataStoreVendor")
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")
	query.Fields = append(query.Fields, "M_port")

	/** associations of DataStoreConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var DataStoreConfigurationInfo []interface{}

	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.GetUuid())
	if this.parentPtr != nil {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.parentPtr.GetUuid())
	} else {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, "")
	}

	/** members of Configuration **/
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_id)

	/** members of DataStoreConfiguration **/

	/** Save dataStoreType type DataStoreType **/
	if this.object.M_dataStoreType == CargoConfig.DataStoreType_SQL_STORE {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 0)
	} else if this.object.M_dataStoreType == CargoConfig.DataStoreType_MEMORY_STORE {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 1)
	} else if this.object.M_dataStoreType == CargoConfig.DataStoreType_KEY_VALUE_STORE {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 2)
	} else if this.object.M_dataStoreType == CargoConfig.DataStoreType_GRAPH_STORE {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 3)
	} else {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 0)
	}

	/** Save dataStoreVendor type DataStoreVendor **/
	if this.object.M_dataStoreVendor == CargoConfig.DataStoreVendor_MYCELIUS {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 0)
	} else if this.object.M_dataStoreVendor == CargoConfig.DataStoreVendor_MYSQL {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 1)
	} else if this.object.M_dataStoreVendor == CargoConfig.DataStoreVendor_MSSQL {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 2)
	} else if this.object.M_dataStoreVendor == CargoConfig.DataStoreVendor_ODBC {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 3)
	} else if this.object.M_dataStoreVendor == CargoConfig.DataStoreVendor_KNOWLEDGEBASE {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 4)
	} else {
		DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, 0)
	}
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_hostName)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_ipv4)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_user)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_pwd)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_port)

	/** associations of DataStoreConfiguration **/

	/** Save parent type Configurations **/
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, this.object.M_parentPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	DataStoreConfigurationInfo = append(DataStoreConfigurationInfo, string(referencedStr))
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "uuid="+this.uuid)
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoConfigDB, string(queryStr), DataStoreConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoConfigDB, string(queryStr), DataStoreConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoConfig_DataStoreConfigurationEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoConfig_DataStoreConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoConfig.DataStoreConfiguration"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of DataStoreConfiguration **/
	query.Fields = append(query.Fields, "M_dataStoreType")
	query.Fields = append(query.Fields, "M_dataStoreVendor")
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_user")
	query.Fields = append(query.Fields, "M_pwd")
	query.Fields = append(query.Fields, "M_port")

	/** associations of DataStoreConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of DataStoreConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoConfig.DataStoreConfiguration)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoConfig.DataStoreConfiguration"

		this.parentUuid = results[0][1].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of DataStoreConfiguration **/

		/** dataStoreType **/
		if results[0][3] != nil {
			enumIndex := results[0][3].(int)
			if enumIndex == 0 {
				this.object.M_dataStoreType = CargoConfig.DataStoreType_SQL_STORE
			} else if enumIndex == 1 {
				this.object.M_dataStoreType = CargoConfig.DataStoreType_MEMORY_STORE
			} else if enumIndex == 2 {
				this.object.M_dataStoreType = CargoConfig.DataStoreType_KEY_VALUE_STORE
			} else if enumIndex == 3 {
				this.object.M_dataStoreType = CargoConfig.DataStoreType_GRAPH_STORE
			}
		}

		/** dataStoreVendor **/
		if results[0][4] != nil {
			enumIndex := results[0][4].(int)
			if enumIndex == 0 {
				this.object.M_dataStoreVendor = CargoConfig.DataStoreVendor_MYCELIUS
			} else if enumIndex == 1 {
				this.object.M_dataStoreVendor = CargoConfig.DataStoreVendor_MYSQL
			} else if enumIndex == 2 {
				this.object.M_dataStoreVendor = CargoConfig.DataStoreVendor_MSSQL
			} else if enumIndex == 3 {
				this.object.M_dataStoreVendor = CargoConfig.DataStoreVendor_ODBC
			} else if enumIndex == 4 {
				this.object.M_dataStoreVendor = CargoConfig.DataStoreVendor_KNOWLEDGEBASE
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

		/** user **/
		if results[0][7] != nil {
			this.object.M_user = results[0][7].(string)
		}

		/** pwd **/
		if results[0][8] != nil {
			this.object.M_pwd = results[0][8].(string)
		}

		/** port **/
		if results[0][9] != nil {
			this.object.M_port = results[0][9].(int)
		}

		/** associations of DataStoreConfiguration **/

		/** parentPtr **/
		if results[0][10] != nil {
			id := results[0][10].(string)
			if len(id) > 0 {
				refTypeName := "CargoConfig.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][11].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][12].(string)
		this.referenced = make([]EntityRef, 0)
		err = json.Unmarshal([]byte(referencedStr), &this.referenced)
		if err != nil {
			return err
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoConfigDataStoreConfigurationEntityFromObject(object *CargoConfig.DataStoreConfiguration) *CargoConfig_DataStoreConfigurationEntity {
	return this.NewCargoConfigDataStoreConfigurationEntity(object.UUID, object)
}

/** Delete **/
func (this *CargoConfig_DataStoreConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoConfigDataStoreConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoConfig.DataStoreConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoConfig_DataStoreConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
		this.childsPtr = append(this.childsPtr, child)
	} else {
		childsPtr := make([]Entity, 0)
		for i := 0; i < len(this.childsPtr); i++ {
			if this.childsPtr[i].GetUuid() != child.GetUuid() {
				childsPtr = append(childsPtr, this.childsPtr[i])
			}
		}
		childsPtr = append(childsPtr, child)
		this.SetChildsPtr(childsPtr)
	}
	// Set this as parent in the child
	child.SetParentPtr(this)

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
func (this *CargoConfig_DataStoreConfigurationEntity) AppendReference(reference Entity) {

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
		this.referencesPtr = append(this.referencesPtr, reference)
	} else {
		// The reference must be update in that case.
		this.referencesPtr[index] = reference
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			LdapConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoConfig_LdapConfigurationEntity struct {
	/** not the object id, except for the definition **/
	uuid           string
	parentPtr      Entity
	parentUuid     string
	childsPtr      []Entity
	childsUuid     []string
	referencesUuid []string
	referencesPtr  []Entity
	prototype      *EntityPrototype
	referenced     []EntityRef
	object         *CargoConfig.LdapConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewCargoConfigLdapConfigurationEntity(objectId string, object interface{}) *CargoConfig_LdapConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoConfigLdapConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoConfig.LdapConfiguration).TYPENAME = "CargoConfig.LdapConfiguration"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoConfig.LdapConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*CargoConfig_LdapConfigurationEntity)
		}
	} else {
		uuidStr = "CargoConfig.LdapConfiguration%" + uuid.NewRandom().String()
	}
	entity := new(CargoConfig_LdapConfigurationEntity)
	if object == nil {
		entity.object = new(CargoConfig.LdapConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoConfig.LdapConfiguration)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoConfig.LdapConfiguration"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoConfig.LdapConfiguration", "CargoConfig")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoConfig_LdapConfigurationEntity) GetTypeName() string {
	return "CargoConfig.LdapConfiguration"
}
func (this *CargoConfig_LdapConfigurationEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoConfig_LdapConfigurationEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoConfig_LdapConfigurationEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoConfig_LdapConfigurationEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoConfig_LdapConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoConfig_LdapConfigurationEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoConfig_LdapConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	refsPtr := make([]Entity, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsPtr = append(refsPtr, reference)
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	this.SetReferencesPtr(refsPtr)

	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_LdapConfigurationEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoConfig_LdapConfigurationEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoConfig_LdapConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoConfig_LdapConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoConfig_LdapConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		}
	}

	this.childsUuid = childsUuid
	params := make([]interface{}, 1)
	childsPtr := make([]Entity, 0)
	for i := 0; i < len(this.GetChildsPtr()); i++ {
		if this.GetChildsPtr()[i].GetUuid() != uuid {
			childsPtr = append(childsPtr, this.GetChildsPtr()[i])
		} else {
			params[0] = this.GetChildsPtr()[i].GetObject()
		}
	}
	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_LdapConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoConfig_LdapConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoConfig_LdapConfigurationEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoConfig_LdapConfigurationEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoConfig_LdapConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoConfig_LdapConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoConfig_LdapConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoConfig_LdapConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoConfig_LdapConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoConfig_LdapConfigurationEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *CargoConfig_LdapConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoConfig.LdapConfiguration"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoConfig_LdapConfigurationEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) Create_CargoConfig_LdapConfigurationEntityPrototype() {

	var ldapConfigurationEntityProto EntityPrototype
	ldapConfigurationEntityProto.TypeName = "CargoConfig.LdapConfiguration"
	ldapConfigurationEntityProto.SuperTypeNames = append(ldapConfigurationEntityProto.SuperTypeNames, "CargoConfig.Configuration")
	ldapConfigurationEntityProto.Ids = append(ldapConfigurationEntityProto.Ids, "uuid")
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "uuid")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 0)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, false)
	ldapConfigurationEntityProto.Indexs = append(ldapConfigurationEntityProto.Indexs, "parentUuid")
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "parentUuid")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 1)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, false)

	/** members of Configuration **/
	ldapConfigurationEntityProto.Ids = append(ldapConfigurationEntityProto.Ids, "M_id")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 2)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_id")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.ID")

	/** members of LdapConfiguration **/
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 3)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_hostName")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 4)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_ipv4")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 5)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_port")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.int")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 6)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_user")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 7)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_pwd")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 8)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_domain")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 9)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, true)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_searchBase")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "xs.string")

	/** associations of LdapConfiguration **/
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 10)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, false)
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "M_parentPtr")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "CargoConfig.Configurations:Ref")
	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "childsUuid")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "[]xs.string")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 11)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, false)

	ldapConfigurationEntityProto.Fields = append(ldapConfigurationEntityProto.Fields, "referenced")
	ldapConfigurationEntityProto.FieldsType = append(ldapConfigurationEntityProto.FieldsType, "[]EntityRef")
	ldapConfigurationEntityProto.FieldsOrder = append(ldapConfigurationEntityProto.FieldsOrder, 12)
	ldapConfigurationEntityProto.FieldsVisibility = append(ldapConfigurationEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&ldapConfigurationEntityProto)

}

/** Create **/
func (this *CargoConfig_LdapConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoConfig.LdapConfiguration"

	var query EntityQuery
	query.TypeName = "CargoConfig.LdapConfiguration"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

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

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var LdapConfigurationInfo []interface{}

	LdapConfigurationInfo = append(LdapConfigurationInfo, this.GetUuid())
	if this.parentPtr != nil {
		LdapConfigurationInfo = append(LdapConfigurationInfo, this.parentPtr.GetUuid())
	} else {
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
	LdapConfigurationInfo = append(LdapConfigurationInfo, this.object.M_parentPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	LdapConfigurationInfo = append(LdapConfigurationInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	LdapConfigurationInfo = append(LdapConfigurationInfo, string(referencedStr))
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "uuid="+this.uuid)
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoConfigDB, string(queryStr), LdapConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoConfigDB, string(queryStr), LdapConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoConfig_LdapConfigurationEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoConfig_LdapConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoConfig.LdapConfiguration"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

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

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of LdapConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoConfig.LdapConfiguration)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoConfig.LdapConfiguration"

		this.parentUuid = results[0][1].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of LdapConfiguration **/

		/** hostName **/
		if results[0][3] != nil {
			this.object.M_hostName = results[0][3].(string)
		}

		/** ipv4 **/
		if results[0][4] != nil {
			this.object.M_ipv4 = results[0][4].(string)
		}

		/** port **/
		if results[0][5] != nil {
			this.object.M_port = results[0][5].(int)
		}

		/** user **/
		if results[0][6] != nil {
			this.object.M_user = results[0][6].(string)
		}

		/** pwd **/
		if results[0][7] != nil {
			this.object.M_pwd = results[0][7].(string)
		}

		/** domain **/
		if results[0][8] != nil {
			this.object.M_domain = results[0][8].(string)
		}

		/** searchBase **/
		if results[0][9] != nil {
			this.object.M_searchBase = results[0][9].(string)
		}

		/** associations of LdapConfiguration **/

		/** parentPtr **/
		if results[0][10] != nil {
			id := results[0][10].(string)
			if len(id) > 0 {
				refTypeName := "CargoConfig.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][11].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][12].(string)
		this.referenced = make([]EntityRef, 0)
		err = json.Unmarshal([]byte(referencedStr), &this.referenced)
		if err != nil {
			return err
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoConfigLdapConfigurationEntityFromObject(object *CargoConfig.LdapConfiguration) *CargoConfig_LdapConfigurationEntity {
	return this.NewCargoConfigLdapConfigurationEntity(object.UUID, object)
}

/** Delete **/
func (this *CargoConfig_LdapConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoConfigLdapConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoConfig.LdapConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoConfig_LdapConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
		this.childsPtr = append(this.childsPtr, child)
	} else {
		childsPtr := make([]Entity, 0)
		for i := 0; i < len(this.childsPtr); i++ {
			if this.childsPtr[i].GetUuid() != child.GetUuid() {
				childsPtr = append(childsPtr, this.childsPtr[i])
			}
		}
		childsPtr = append(childsPtr, child)
		this.SetChildsPtr(childsPtr)
	}
	// Set this as parent in the child
	child.SetParentPtr(this)

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
func (this *CargoConfig_LdapConfigurationEntity) AppendReference(reference Entity) {

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
		this.referencesPtr = append(this.referencesPtr, reference)
	} else {
		// The reference must be update in that case.
		this.referencesPtr[index] = reference
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			ApplicationConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoConfig_ApplicationConfigurationEntity struct {
	/** not the object id, except for the definition **/
	uuid           string
	parentPtr      Entity
	parentUuid     string
	childsPtr      []Entity
	childsUuid     []string
	referencesUuid []string
	referencesPtr  []Entity
	prototype      *EntityPrototype
	referenced     []EntityRef
	object         *CargoConfig.ApplicationConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewCargoConfigApplicationConfigurationEntity(objectId string, object interface{}) *CargoConfig_ApplicationConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoConfigApplicationConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoConfig.ApplicationConfiguration).TYPENAME = "CargoConfig.ApplicationConfiguration"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoConfig.ApplicationConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*CargoConfig_ApplicationConfigurationEntity)
		}
	} else {
		uuidStr = "CargoConfig.ApplicationConfiguration%" + uuid.NewRandom().String()
	}
	entity := new(CargoConfig_ApplicationConfigurationEntity)
	if object == nil {
		entity.object = new(CargoConfig.ApplicationConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoConfig.ApplicationConfiguration)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoConfig.ApplicationConfiguration"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoConfig.ApplicationConfiguration", "CargoConfig")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoConfig_ApplicationConfigurationEntity) GetTypeName() string {
	return "CargoConfig.ApplicationConfiguration"
}
func (this *CargoConfig_ApplicationConfigurationEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoConfig_ApplicationConfigurationEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoConfig_ApplicationConfigurationEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoConfig_ApplicationConfigurationEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoConfig_ApplicationConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoConfig_ApplicationConfigurationEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoConfig_ApplicationConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	refsPtr := make([]Entity, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsPtr = append(refsPtr, reference)
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	this.SetReferencesPtr(refsPtr)

	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_ApplicationConfigurationEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoConfig_ApplicationConfigurationEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoConfig_ApplicationConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoConfig_ApplicationConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoConfig_ApplicationConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		}
	}

	this.childsUuid = childsUuid
	params := make([]interface{}, 1)
	childsPtr := make([]Entity, 0)
	for i := 0; i < len(this.GetChildsPtr()); i++ {
		if this.GetChildsPtr()[i].GetUuid() != uuid {
			childsPtr = append(childsPtr, this.GetChildsPtr()[i])
		} else {
			params[0] = this.GetChildsPtr()[i].GetObject()
		}
	}
	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_ApplicationConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoConfig_ApplicationConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoConfig_ApplicationConfigurationEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoConfig_ApplicationConfigurationEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoConfig_ApplicationConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoConfig_ApplicationConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoConfig_ApplicationConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoConfig_ApplicationConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoConfig_ApplicationConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoConfig_ApplicationConfigurationEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *CargoConfig_ApplicationConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoConfig.ApplicationConfiguration"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoConfig_ApplicationConfigurationEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) Create_CargoConfig_ApplicationConfigurationEntityPrototype() {

	var applicationConfigurationEntityProto EntityPrototype
	applicationConfigurationEntityProto.TypeName = "CargoConfig.ApplicationConfiguration"
	applicationConfigurationEntityProto.SuperTypeNames = append(applicationConfigurationEntityProto.SuperTypeNames, "CargoConfig.Configuration")
	applicationConfigurationEntityProto.Ids = append(applicationConfigurationEntityProto.Ids, "uuid")
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "uuid")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "xs.string")
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 0)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, false)
	applicationConfigurationEntityProto.Indexs = append(applicationConfigurationEntityProto.Indexs, "parentUuid")
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "parentUuid")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "xs.string")
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 1)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, false)

	/** members of Configuration **/
	applicationConfigurationEntityProto.Ids = append(applicationConfigurationEntityProto.Ids, "M_id")
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 2)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, true)
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "M_id")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "xs.ID")

	/** members of ApplicationConfiguration **/
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 3)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, true)
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "M_indexPage")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "xs.string")

	/** associations of ApplicationConfiguration **/
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 4)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, false)
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "M_parentPtr")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "CargoConfig.Configurations:Ref")
	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "childsUuid")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "[]xs.string")
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 5)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, false)

	applicationConfigurationEntityProto.Fields = append(applicationConfigurationEntityProto.Fields, "referenced")
	applicationConfigurationEntityProto.FieldsType = append(applicationConfigurationEntityProto.FieldsType, "[]EntityRef")
	applicationConfigurationEntityProto.FieldsOrder = append(applicationConfigurationEntityProto.FieldsOrder, 6)
	applicationConfigurationEntityProto.FieldsVisibility = append(applicationConfigurationEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&applicationConfigurationEntityProto)

}

/** Create **/
func (this *CargoConfig_ApplicationConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoConfig.ApplicationConfiguration"

	var query EntityQuery
	query.TypeName = "CargoConfig.ApplicationConfiguration"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ApplicationConfiguration **/
	query.Fields = append(query.Fields, "M_indexPage")

	/** associations of ApplicationConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ApplicationConfigurationInfo []interface{}

	ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.GetUuid())
	if this.parentPtr != nil {
		ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.parentPtr.GetUuid())
	} else {
		ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, "")
	}

	/** members of Configuration **/
	ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.object.M_id)

	/** members of ApplicationConfiguration **/
	ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.object.M_indexPage)

	/** associations of ApplicationConfiguration **/

	/** Save parent type Configurations **/
	ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, this.object.M_parentPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ApplicationConfigurationInfo = append(ApplicationConfigurationInfo, string(referencedStr))
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "uuid="+this.uuid)
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoConfigDB, string(queryStr), ApplicationConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoConfigDB, string(queryStr), ApplicationConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoConfig_ApplicationConfigurationEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoConfig_ApplicationConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoConfig.ApplicationConfiguration"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ApplicationConfiguration **/
	query.Fields = append(query.Fields, "M_indexPage")

	/** associations of ApplicationConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ApplicationConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoConfig.ApplicationConfiguration)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoConfig.ApplicationConfiguration"

		this.parentUuid = results[0][1].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of ApplicationConfiguration **/

		/** indexPage **/
		if results[0][3] != nil {
			this.object.M_indexPage = results[0][3].(string)
		}

		/** associations of ApplicationConfiguration **/

		/** parentPtr **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "CargoConfig.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][5].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][6].(string)
		this.referenced = make([]EntityRef, 0)
		err = json.Unmarshal([]byte(referencedStr), &this.referenced)
		if err != nil {
			return err
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoConfigApplicationConfigurationEntityFromObject(object *CargoConfig.ApplicationConfiguration) *CargoConfig_ApplicationConfigurationEntity {
	return this.NewCargoConfigApplicationConfigurationEntity(object.UUID, object)
}

/** Delete **/
func (this *CargoConfig_ApplicationConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoConfigApplicationConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoConfig.ApplicationConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoConfig_ApplicationConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
		this.childsPtr = append(this.childsPtr, child)
	} else {
		childsPtr := make([]Entity, 0)
		for i := 0; i < len(this.childsPtr); i++ {
			if this.childsPtr[i].GetUuid() != child.GetUuid() {
				childsPtr = append(childsPtr, this.childsPtr[i])
			}
		}
		childsPtr = append(childsPtr, child)
		this.SetChildsPtr(childsPtr)
	}
	// Set this as parent in the child
	child.SetParentPtr(this)

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
func (this *CargoConfig_ApplicationConfigurationEntity) AppendReference(reference Entity) {

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
		this.referencesPtr = append(this.referencesPtr, reference)
	} else {
		// The reference must be update in that case.
		this.referencesPtr[index] = reference
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			ServerConfiguration
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoConfig_ServerConfigurationEntity struct {
	/** not the object id, except for the definition **/
	uuid           string
	parentPtr      Entity
	parentUuid     string
	childsPtr      []Entity
	childsUuid     []string
	referencesUuid []string
	referencesPtr  []Entity
	prototype      *EntityPrototype
	referenced     []EntityRef
	object         *CargoConfig.ServerConfiguration
}

/** Constructor function **/
func (this *EntityManager) NewCargoConfigServerConfigurationEntity(objectId string, object interface{}) *CargoConfig_ServerConfigurationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoConfigServerConfigurationExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoConfig.ServerConfiguration).TYPENAME = "CargoConfig.ServerConfiguration"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoConfig.ServerConfiguration).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*CargoConfig_ServerConfigurationEntity)
		}
	} else {
		uuidStr = "CargoConfig.ServerConfiguration%" + uuid.NewRandom().String()
	}
	entity := new(CargoConfig_ServerConfigurationEntity)
	if object == nil {
		entity.object = new(CargoConfig.ServerConfiguration)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoConfig.ServerConfiguration)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoConfig.ServerConfiguration"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoConfig.ServerConfiguration", "CargoConfig")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoConfig_ServerConfigurationEntity) GetTypeName() string {
	return "CargoConfig.ServerConfiguration"
}
func (this *CargoConfig_ServerConfigurationEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoConfig_ServerConfigurationEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoConfig_ServerConfigurationEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoConfig_ServerConfigurationEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoConfig_ServerConfigurationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoConfig_ServerConfigurationEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoConfig_ServerConfigurationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	refsPtr := make([]Entity, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsPtr = append(refsPtr, reference)
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	this.SetReferencesPtr(refsPtr)

	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_ServerConfigurationEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoConfig_ServerConfigurationEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoConfig_ServerConfigurationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoConfig_ServerConfigurationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoConfig_ServerConfigurationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		}
	}

	this.childsUuid = childsUuid
	params := make([]interface{}, 1)
	childsPtr := make([]Entity, 0)
	for i := 0; i < len(this.GetChildsPtr()); i++ {
		if this.GetChildsPtr()[i].GetUuid() != uuid {
			childsPtr = append(childsPtr, this.GetChildsPtr()[i])
		} else {
			params[0] = this.GetChildsPtr()[i].GetObject()
		}
	}
	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_ServerConfigurationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoConfig_ServerConfigurationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoConfig_ServerConfigurationEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoConfig_ServerConfigurationEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoConfig_ServerConfigurationEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoConfig_ServerConfigurationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoConfig_ServerConfigurationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoConfig_ServerConfigurationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoConfig_ServerConfigurationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoConfig_ServerConfigurationEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *CargoConfig_ServerConfigurationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoConfig.ServerConfiguration"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoConfig_ServerConfigurationEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) Create_CargoConfig_ServerConfigurationEntityPrototype() {

	var serverConfigurationEntityProto EntityPrototype
	serverConfigurationEntityProto.TypeName = "CargoConfig.ServerConfiguration"
	serverConfigurationEntityProto.SuperTypeNames = append(serverConfigurationEntityProto.SuperTypeNames, "CargoConfig.Configuration")
	serverConfigurationEntityProto.Ids = append(serverConfigurationEntityProto.Ids, "uuid")
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "uuid")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 0)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, false)
	serverConfigurationEntityProto.Indexs = append(serverConfigurationEntityProto.Indexs, "parentUuid")
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "parentUuid")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 1)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, false)

	/** members of Configuration **/
	serverConfigurationEntityProto.Ids = append(serverConfigurationEntityProto.Ids, "M_id")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 2)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_id")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.ID")

	/** members of ServerConfiguration **/
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 3)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_hostName")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 4)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_ipv4")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 5)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_port")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.int")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 6)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_applicationsPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 7)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_dataPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 8)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_scriptsPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 9)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_definitionsPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 10)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_schemasPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 11)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_tmpPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 12)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, true)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_binPath")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "xs.string")

	/** associations of ServerConfiguration **/
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 13)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, false)
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "M_parentPtr")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "CargoConfig.Configurations:Ref")
	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "childsUuid")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "[]xs.string")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 14)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, false)

	serverConfigurationEntityProto.Fields = append(serverConfigurationEntityProto.Fields, "referenced")
	serverConfigurationEntityProto.FieldsType = append(serverConfigurationEntityProto.FieldsType, "[]EntityRef")
	serverConfigurationEntityProto.FieldsOrder = append(serverConfigurationEntityProto.FieldsOrder, 15)
	serverConfigurationEntityProto.FieldsVisibility = append(serverConfigurationEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&serverConfigurationEntityProto)

}

/** Create **/
func (this *CargoConfig_ServerConfigurationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoConfig.ServerConfiguration"

	var query EntityQuery
	query.TypeName = "CargoConfig.ServerConfiguration"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ServerConfiguration **/
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_port")
	query.Fields = append(query.Fields, "M_applicationsPath")
	query.Fields = append(query.Fields, "M_dataPath")
	query.Fields = append(query.Fields, "M_scriptsPath")
	query.Fields = append(query.Fields, "M_definitionsPath")
	query.Fields = append(query.Fields, "M_schemasPath")
	query.Fields = append(query.Fields, "M_tmpPath")
	query.Fields = append(query.Fields, "M_binPath")

	/** associations of ServerConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ServerConfigurationInfo []interface{}

	ServerConfigurationInfo = append(ServerConfigurationInfo, this.GetUuid())
	if this.parentPtr != nil {
		ServerConfigurationInfo = append(ServerConfigurationInfo, this.parentPtr.GetUuid())
	} else {
		ServerConfigurationInfo = append(ServerConfigurationInfo, "")
	}

	/** members of Configuration **/
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_id)

	/** members of ServerConfiguration **/
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_hostName)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_ipv4)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_port)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_applicationsPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_dataPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_scriptsPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_definitionsPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_schemasPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_tmpPath)
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_binPath)

	/** associations of ServerConfiguration **/

	/** Save parent type Configurations **/
	ServerConfigurationInfo = append(ServerConfigurationInfo, this.object.M_parentPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ServerConfigurationInfo = append(ServerConfigurationInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ServerConfigurationInfo = append(ServerConfigurationInfo, string(referencedStr))
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "uuid="+this.uuid)
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoConfigDB, string(queryStr), ServerConfigurationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoConfigDB, string(queryStr), ServerConfigurationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoConfig_ServerConfigurationEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoConfig_ServerConfigurationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoConfig.ServerConfiguration"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Configuration **/
	query.Fields = append(query.Fields, "M_id")

	/** members of ServerConfiguration **/
	query.Fields = append(query.Fields, "M_hostName")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_port")
	query.Fields = append(query.Fields, "M_applicationsPath")
	query.Fields = append(query.Fields, "M_dataPath")
	query.Fields = append(query.Fields, "M_scriptsPath")
	query.Fields = append(query.Fields, "M_definitionsPath")
	query.Fields = append(query.Fields, "M_schemasPath")
	query.Fields = append(query.Fields, "M_tmpPath")
	query.Fields = append(query.Fields, "M_binPath")

	/** associations of ServerConfiguration **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ServerConfiguration...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoConfig.ServerConfiguration)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoConfig.ServerConfiguration"

		this.parentUuid = results[0][1].(string)

		/** members of Configuration **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of ServerConfiguration **/

		/** hostName **/
		if results[0][3] != nil {
			this.object.M_hostName = results[0][3].(string)
		}

		/** ipv4 **/
		if results[0][4] != nil {
			this.object.M_ipv4 = results[0][4].(string)
		}

		/** port **/
		if results[0][5] != nil {
			this.object.M_port = results[0][5].(int)
		}

		/** applicationsPath **/
		if results[0][6] != nil {
			this.object.M_applicationsPath = results[0][6].(string)
		}

		/** dataPath **/
		if results[0][7] != nil {
			this.object.M_dataPath = results[0][7].(string)
		}

		/** scriptsPath **/
		if results[0][8] != nil {
			this.object.M_scriptsPath = results[0][8].(string)
		}

		/** definitionsPath **/
		if results[0][9] != nil {
			this.object.M_definitionsPath = results[0][9].(string)
		}

		/** schemasPath **/
		if results[0][10] != nil {
			this.object.M_schemasPath = results[0][10].(string)
		}

		/** tmpPath **/
		if results[0][11] != nil {
			this.object.M_tmpPath = results[0][11].(string)
		}

		/** binPath **/
		if results[0][12] != nil {
			this.object.M_binPath = results[0][12].(string)
		}

		/** associations of ServerConfiguration **/

		/** parentPtr **/
		if results[0][13] != nil {
			id := results[0][13].(string)
			if len(id) > 0 {
				refTypeName := "CargoConfig.Configurations"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][14].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][15].(string)
		this.referenced = make([]EntityRef, 0)
		err = json.Unmarshal([]byte(referencedStr), &this.referenced)
		if err != nil {
			return err
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoConfigServerConfigurationEntityFromObject(object *CargoConfig.ServerConfiguration) *CargoConfig_ServerConfigurationEntity {
	return this.NewCargoConfigServerConfigurationEntity(object.UUID, object)
}

/** Delete **/
func (this *CargoConfig_ServerConfigurationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoConfigServerConfigurationExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoConfig.ServerConfiguration"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoConfig_ServerConfigurationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
		this.childsPtr = append(this.childsPtr, child)
	} else {
		childsPtr := make([]Entity, 0)
		for i := 0; i < len(this.childsPtr); i++ {
			if this.childsPtr[i].GetUuid() != child.GetUuid() {
				childsPtr = append(childsPtr, this.childsPtr[i])
			}
		}
		childsPtr = append(childsPtr, child)
		this.SetChildsPtr(childsPtr)
	}
	// Set this as parent in the child
	child.SetParentPtr(this)

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
func (this *CargoConfig_ServerConfigurationEntity) AppendReference(reference Entity) {

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
		this.referencesPtr = append(this.referencesPtr, reference)
	} else {
		// The reference must be update in that case.
		this.referencesPtr[index] = reference
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Configurations
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoConfig_ConfigurationsEntity struct {
	/** not the object id, except for the definition **/
	uuid           string
	parentPtr      Entity
	parentUuid     string
	childsPtr      []Entity
	childsUuid     []string
	referencesUuid []string
	referencesPtr  []Entity
	prototype      *EntityPrototype
	referenced     []EntityRef
	object         *CargoConfig.Configurations
}

/** Constructor function **/
func (this *EntityManager) NewCargoConfigConfigurationsEntity(objectId string, object interface{}) *CargoConfig_ConfigurationsEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoConfigConfigurationsExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoConfig.Configurations).TYPENAME = "CargoConfig.Configurations"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoConfig.Configurations).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*CargoConfig_ConfigurationsEntity)
		}
	} else {
		uuidStr = "CargoConfig.Configurations%" + uuid.NewRandom().String()
	}
	entity := new(CargoConfig_ConfigurationsEntity)
	if object == nil {
		entity.object = new(CargoConfig.Configurations)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoConfig.Configurations)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoConfig.Configurations"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoConfig.Configurations", "CargoConfig")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoConfig_ConfigurationsEntity) GetTypeName() string {
	return "CargoConfig.Configurations"
}
func (this *CargoConfig_ConfigurationsEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoConfig_ConfigurationsEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoConfig_ConfigurationsEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoConfig_ConfigurationsEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoConfig_ConfigurationsEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoConfig_ConfigurationsEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoConfig_ConfigurationsEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	refsPtr := make([]Entity, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsPtr = append(refsPtr, reference)
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	this.SetReferencesPtr(refsPtr)

	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_ConfigurationsEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoConfig_ConfigurationsEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoConfig_ConfigurationsEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoConfig_ConfigurationsEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoConfig_ConfigurationsEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		}
	}

	this.childsUuid = childsUuid
	params := make([]interface{}, 1)
	childsPtr := make([]Entity, 0)
	for i := 0; i < len(this.GetChildsPtr()); i++ {
		if this.GetChildsPtr()[i].GetUuid() != uuid {
			childsPtr = append(childsPtr, this.GetChildsPtr()[i])
		} else {
			params[0] = this.GetChildsPtr()[i].GetObject()
		}
	}
	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoConfig_ConfigurationsEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoConfig_ConfigurationsEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoConfig_ConfigurationsEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoConfig_ConfigurationsEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoConfig_ConfigurationsEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoConfig_ConfigurationsEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoConfig_ConfigurationsEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoConfig_ConfigurationsEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoConfig_ConfigurationsEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoConfig_ConfigurationsEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *CargoConfig_ConfigurationsEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoConfig.Configurations"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoConfig_ConfigurationsEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) Create_CargoConfig_ConfigurationsEntityPrototype() {

	var configurationsEntityProto EntityPrototype
	configurationsEntityProto.TypeName = "CargoConfig.Configurations"
	configurationsEntityProto.Ids = append(configurationsEntityProto.Ids, "uuid")
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "uuid")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.string")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 0)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, false)
	configurationsEntityProto.Indexs = append(configurationsEntityProto.Indexs, "parentUuid")
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "parentUuid")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.string")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 1)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, false)

	/** members of Configurations **/
	configurationsEntityProto.Ids = append(configurationsEntityProto.Ids, "M_id")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 2)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_id")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.ID")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 3)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_name")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.string")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 4)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_version")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.string")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 5)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_filePath")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "xs.string")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 6)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_serverConfig")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "CargoConfig.ServerConfiguration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 7)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_dataStoreConfigs")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]CargoConfig.DataStoreConfiguration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 8)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_smtpConfigs")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]CargoConfig.SmtpConfiguration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 9)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_ldapConfigs")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]CargoConfig.LdapConfiguration")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 10)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, true)
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "M_applicationConfigs")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]CargoConfig.ApplicationConfiguration")
	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "childsUuid")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]xs.string")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 11)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, false)

	configurationsEntityProto.Fields = append(configurationsEntityProto.Fields, "referenced")
	configurationsEntityProto.FieldsType = append(configurationsEntityProto.FieldsType, "[]EntityRef")
	configurationsEntityProto.FieldsOrder = append(configurationsEntityProto.FieldsOrder, 12)
	configurationsEntityProto.FieldsVisibility = append(configurationsEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoConfigDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&configurationsEntityProto)

}

/** Create **/
func (this *CargoConfig_ConfigurationsEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoConfig.Configurations"

	var query EntityQuery
	query.TypeName = "CargoConfig.Configurations"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Configurations **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_version")
	query.Fields = append(query.Fields, "M_filePath")
	query.Fields = append(query.Fields, "M_serverConfig")
	query.Fields = append(query.Fields, "M_dataStoreConfigs")
	query.Fields = append(query.Fields, "M_smtpConfigs")
	query.Fields = append(query.Fields, "M_ldapConfigs")
	query.Fields = append(query.Fields, "M_applicationConfigs")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ConfigurationsInfo []interface{}

	ConfigurationsInfo = append(ConfigurationsInfo, this.GetUuid())
	if this.parentPtr != nil {
		ConfigurationsInfo = append(ConfigurationsInfo, this.parentPtr.GetUuid())
	} else {
		ConfigurationsInfo = append(ConfigurationsInfo, "")
	}

	/** members of Configurations **/
	ConfigurationsInfo = append(ConfigurationsInfo, this.object.M_id)
	ConfigurationsInfo = append(ConfigurationsInfo, this.object.M_name)
	ConfigurationsInfo = append(ConfigurationsInfo, this.object.M_version)
	ConfigurationsInfo = append(ConfigurationsInfo, this.object.M_filePath)

	/** Save serverConfig type ServerConfiguration **/
	if this.object.M_serverConfig != nil {
		serverConfigEntity := GetServer().GetEntityManager().NewCargoConfigServerConfigurationEntity(this.object.M_serverConfig.UUID, this.object.M_serverConfig)
		ConfigurationsInfo = append(ConfigurationsInfo, serverConfigEntity.uuid)
		serverConfigEntity.AppendReferenced("serverConfig", this)
		this.AppendChild("serverConfig", serverConfigEntity)
		if serverConfigEntity.NeedSave() {
			serverConfigEntity.SaveEntity()
		}
	} else {
		ConfigurationsInfo = append(ConfigurationsInfo, "")
	}

	/** Save dataStoreConfigs type DataStoreConfiguration **/
	dataStoreConfigsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_dataStoreConfigs); i++ {
		dataStoreConfigsEntity := GetServer().GetEntityManager().NewCargoConfigDataStoreConfigurationEntity(this.object.M_dataStoreConfigs[i].UUID, this.object.M_dataStoreConfigs[i])
		dataStoreConfigsIds = append(dataStoreConfigsIds, dataStoreConfigsEntity.uuid)
		dataStoreConfigsEntity.AppendReferenced("dataStoreConfigs", this)
		this.AppendChild("dataStoreConfigs", dataStoreConfigsEntity)
		if dataStoreConfigsEntity.NeedSave() {
			dataStoreConfigsEntity.SaveEntity()
		}
	}
	dataStoreConfigsStr, _ := json.Marshal(dataStoreConfigsIds)
	ConfigurationsInfo = append(ConfigurationsInfo, string(dataStoreConfigsStr))

	/** Save smtpConfigs type SmtpConfiguration **/
	smtpConfigsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_smtpConfigs); i++ {
		smtpConfigsEntity := GetServer().GetEntityManager().NewCargoConfigSmtpConfigurationEntity(this.object.M_smtpConfigs[i].UUID, this.object.M_smtpConfigs[i])
		smtpConfigsIds = append(smtpConfigsIds, smtpConfigsEntity.uuid)
		smtpConfigsEntity.AppendReferenced("smtpConfigs", this)
		this.AppendChild("smtpConfigs", smtpConfigsEntity)
		if smtpConfigsEntity.NeedSave() {
			smtpConfigsEntity.SaveEntity()
		}
	}
	smtpConfigsStr, _ := json.Marshal(smtpConfigsIds)
	ConfigurationsInfo = append(ConfigurationsInfo, string(smtpConfigsStr))

	/** Save ldapConfigs type LdapConfiguration **/
	ldapConfigsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_ldapConfigs); i++ {
		ldapConfigsEntity := GetServer().GetEntityManager().NewCargoConfigLdapConfigurationEntity(this.object.M_ldapConfigs[i].UUID, this.object.M_ldapConfigs[i])
		ldapConfigsIds = append(ldapConfigsIds, ldapConfigsEntity.uuid)
		ldapConfigsEntity.AppendReferenced("ldapConfigs", this)
		this.AppendChild("ldapConfigs", ldapConfigsEntity)
		if ldapConfigsEntity.NeedSave() {
			ldapConfigsEntity.SaveEntity()
		}
	}
	ldapConfigsStr, _ := json.Marshal(ldapConfigsIds)
	ConfigurationsInfo = append(ConfigurationsInfo, string(ldapConfigsStr))

	/** Save applicationConfigs type ApplicationConfiguration **/
	applicationConfigsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_applicationConfigs); i++ {
		applicationConfigsEntity := GetServer().GetEntityManager().NewCargoConfigApplicationConfigurationEntity(this.object.M_applicationConfigs[i].UUID, this.object.M_applicationConfigs[i])
		applicationConfigsIds = append(applicationConfigsIds, applicationConfigsEntity.uuid)
		applicationConfigsEntity.AppendReferenced("applicationConfigs", this)
		this.AppendChild("applicationConfigs", applicationConfigsEntity)
		if applicationConfigsEntity.NeedSave() {
			applicationConfigsEntity.SaveEntity()
		}
	}
	applicationConfigsStr, _ := json.Marshal(applicationConfigsIds)
	ConfigurationsInfo = append(ConfigurationsInfo, string(applicationConfigsStr))
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ConfigurationsInfo = append(ConfigurationsInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ConfigurationsInfo = append(ConfigurationsInfo, string(referencedStr))
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "uuid="+this.uuid)
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoConfigDB, string(queryStr), ConfigurationsInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoConfigDB, string(queryStr), ConfigurationsInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoConfig_ConfigurationsEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoConfig_ConfigurationsEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoConfig.Configurations"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Configurations **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_version")
	query.Fields = append(query.Fields, "M_filePath")
	query.Fields = append(query.Fields, "M_serverConfig")
	query.Fields = append(query.Fields, "M_dataStoreConfigs")
	query.Fields = append(query.Fields, "M_smtpConfigs")
	query.Fields = append(query.Fields, "M_ldapConfigs")
	query.Fields = append(query.Fields, "M_applicationConfigs")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Configurations...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoConfig.Configurations)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoConfig.Configurations"

		this.parentUuid = results[0][1].(string)

		/** members of Configurations **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** name **/
		if results[0][3] != nil {
			this.object.M_name = results[0][3].(string)
		}

		/** version **/
		if results[0][4] != nil {
			this.object.M_version = results[0][4].(string)
		}

		/** filePath **/
		if results[0][5] != nil {
			this.object.M_filePath = results[0][5].(string)
		}

		/** serverConfig **/
		if results[0][6] != nil {
			uuid := results[0][6].(string)
			if len(uuid) > 0 {
				var serverConfigEntity *CargoConfig_ServerConfigurationEntity
				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
					serverConfigEntity = instance.(*CargoConfig_ServerConfigurationEntity)
				} else {
					serverConfigEntity = GetServer().GetEntityManager().NewCargoConfigServerConfigurationEntity(uuid, nil)
					serverConfigEntity.InitEntity(uuid)
					GetServer().GetEntityManager().insert(serverConfigEntity)
				}
				serverConfigEntity.AppendReferenced("serverConfig", this)
				this.AppendChild("serverConfig", serverConfigEntity)
			}
		}

		/** dataStoreConfigs **/
		if results[0][7] != nil {
			uuidsStr := results[0][7].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var dataStoreConfigsEntity *CargoConfig_DataStoreConfigurationEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataStoreConfigsEntity = instance.(*CargoConfig_DataStoreConfigurationEntity)
					} else {
						dataStoreConfigsEntity = GetServer().GetEntityManager().NewCargoConfigDataStoreConfigurationEntity(uuids[i], nil)
						dataStoreConfigsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataStoreConfigsEntity)
					}
					dataStoreConfigsEntity.AppendReferenced("dataStoreConfigs", this)
					this.AppendChild("dataStoreConfigs", dataStoreConfigsEntity)
				}
			}
		}

		/** smtpConfigs **/
		if results[0][8] != nil {
			uuidsStr := results[0][8].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var smtpConfigsEntity *CargoConfig_SmtpConfigurationEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						smtpConfigsEntity = instance.(*CargoConfig_SmtpConfigurationEntity)
					} else {
						smtpConfigsEntity = GetServer().GetEntityManager().NewCargoConfigSmtpConfigurationEntity(uuids[i], nil)
						smtpConfigsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(smtpConfigsEntity)
					}
					smtpConfigsEntity.AppendReferenced("smtpConfigs", this)
					this.AppendChild("smtpConfigs", smtpConfigsEntity)
				}
			}
		}

		/** ldapConfigs **/
		if results[0][9] != nil {
			uuidsStr := results[0][9].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var ldapConfigsEntity *CargoConfig_LdapConfigurationEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						ldapConfigsEntity = instance.(*CargoConfig_LdapConfigurationEntity)
					} else {
						ldapConfigsEntity = GetServer().GetEntityManager().NewCargoConfigLdapConfigurationEntity(uuids[i], nil)
						ldapConfigsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(ldapConfigsEntity)
					}
					ldapConfigsEntity.AppendReferenced("ldapConfigs", this)
					this.AppendChild("ldapConfigs", ldapConfigsEntity)
				}
			}
		}

		/** applicationConfigs **/
		if results[0][10] != nil {
			uuidsStr := results[0][10].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var applicationConfigsEntity *CargoConfig_ApplicationConfigurationEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						applicationConfigsEntity = instance.(*CargoConfig_ApplicationConfigurationEntity)
					} else {
						applicationConfigsEntity = GetServer().GetEntityManager().NewCargoConfigApplicationConfigurationEntity(uuids[i], nil)
						applicationConfigsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(applicationConfigsEntity)
					}
					applicationConfigsEntity.AppendReferenced("applicationConfigs", this)
					this.AppendChild("applicationConfigs", applicationConfigsEntity)
				}
			}
		}
		childsUuidStr := results[0][11].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][12].(string)
		this.referenced = make([]EntityRef, 0)
		err = json.Unmarshal([]byte(referencedStr), &this.referenced)
		if err != nil {
			return err
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoConfigConfigurationsEntityFromObject(object *CargoConfig.Configurations) *CargoConfig_ConfigurationsEntity {
	return this.NewCargoConfigConfigurationsEntity(object.UUID, object)
}

/** Delete **/
func (this *CargoConfig_ConfigurationsEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoConfigConfigurationsExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoConfig.Configurations"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoConfigDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoConfig_ConfigurationsEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
		this.childsPtr = append(this.childsPtr, child)
	} else {
		childsPtr := make([]Entity, 0)
		for i := 0; i < len(this.childsPtr); i++ {
			if this.childsPtr[i].GetUuid() != child.GetUuid() {
				childsPtr = append(childsPtr, this.childsPtr[i])
			}
		}
		childsPtr = append(childsPtr, child)
		this.SetChildsPtr(childsPtr)
	}
	// Set this as parent in the child
	child.SetParentPtr(this)

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
func (this *CargoConfig_ConfigurationsEntity) AppendReference(reference Entity) {

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
		this.referencesPtr = append(this.referencesPtr, reference)
	} else {
		// The reference must be update in that case.
		this.referencesPtr[index] = reference
	}
}

/** Register the entity to the dynamic typing system. **/
func (this *EntityManager) RegisterCargoConfigObjects() {
	Utility.RegisterType((*CargoConfig.SmtpConfiguration)(nil))
	Utility.RegisterType((*CargoConfig.DataStoreConfiguration)(nil))
	Utility.RegisterType((*CargoConfig.LdapConfiguration)(nil))
	Utility.RegisterType((*CargoConfig.ApplicationConfiguration)(nil))
	Utility.RegisterType((*CargoConfig.ServerConfiguration)(nil))
	Utility.RegisterType((*CargoConfig.Configurations)(nil))
}

/** Create entity prototypes contain in a package **/
func (this *EntityManager) CreateCargoConfigPrototypes() {
	this.Create_CargoConfig_ConfigurationEntityPrototype()
	this.Create_CargoConfig_SmtpConfigurationEntityPrototype()
	this.Create_CargoConfig_DataStoreConfigurationEntityPrototype()
	this.Create_CargoConfig_LdapConfigurationEntityPrototype()
	this.Create_CargoConfig_ApplicationConfigurationEntityPrototype()
	this.Create_CargoConfig_ServerConfigurationEntityPrototype()
	this.Create_CargoConfig_ConfigurationsEntityPrototype()
}
