// +build CargoEntities

package Server

import (
	"encoding/json"
	"log"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_EntityEntityPrototype() {

	var entityEntityProto EntityPrototype
	entityEntityProto.TypeName = "CargoEntities.Entity"
	entityEntityProto.IsAbstract = true
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Log")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Project")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Notification")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.TextMessage")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Account")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Computer")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.User")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Error")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.LogEntry")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.File")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Group")
	entityEntityProto.Ids = append(entityEntityProto.Ids, "uuid")
	entityEntityProto.Fields = append(entityEntityProto.Fields, "uuid")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.string")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 0)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)
	entityEntityProto.Indexs = append(entityEntityProto.Indexs, "parentUuid")
	entityEntityProto.Fields = append(entityEntityProto.Fields, "parentUuid")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.string")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 1)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	entityEntityProto.Ids = append(entityEntityProto.Ids, "M_id")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 2)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, true)
	entityEntityProto.Fields = append(entityEntityProto.Fields, "M_id")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.ID")

	/** associations of Entity **/
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 3)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)
	entityEntityProto.Fields = append(entityEntityProto.Fields, "M_entitiesPtr")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	entityEntityProto.Fields = append(entityEntityProto.Fields, "childsUuid")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "[]xs.string")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 4)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)

	entityEntityProto.Fields = append(entityEntityProto.Fields, "referenced")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "[]EntityRef")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 5)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&entityEntityProto)

}

////////////////////////////////////////////////////////////////////////////////
//              			Parameter
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_ParameterEntity struct {
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
	object         *CargoEntities.Parameter
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesParameterEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_ParameterEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesParameterExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Parameter).TYPENAME = "CargoEntities.Parameter"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Parameter", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Parameter).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Parameter).UUID
			}
			return val.(*CargoEntities_ParameterEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_ParameterEntity)
	if object == nil {
		entity.object = new(CargoEntities.Parameter)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Parameter)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Parameter"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_ParameterEntity) GetTypeName() string {
	return "CargoEntities.Parameter"
}
func (this *CargoEntities_ParameterEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_ParameterEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_ParameterEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_ParameterEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_ParameterEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_ParameterEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_ParameterEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_ParameterEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_ParameterEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_ParameterEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_ParameterEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_ParameterEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_ParameterEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_ParameterEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_ParameterEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_ParameterEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_ParameterEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_ParameterEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_ParameterEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_ParameterEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_ParameterEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_ParameterEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_ParameterEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Parameter"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_ParameterEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ParameterEntityPrototype() {

	var parameterEntityProto EntityPrototype
	parameterEntityProto.TypeName = "CargoEntities.Parameter"
	parameterEntityProto.Ids = append(parameterEntityProto.Ids, "uuid")
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "uuid")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 0)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)
	parameterEntityProto.Indexs = append(parameterEntityProto.Indexs, "parentUuid")
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "parentUuid")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 1)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)

	/** members of Parameter **/
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 2)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, true)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_name")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 3)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, true)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_type")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 4)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, true)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_isArray")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.boolean")

	/** associations of Parameter **/
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 5)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_parametersPtr")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "CargoEntities.Parameter:Ref")
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "childsUuid")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "[]xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 6)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)

	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "referenced")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "[]EntityRef")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 7)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&parameterEntityProto)

}

/** Create **/
func (this *CargoEntities_ParameterEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Parameter"

	var query EntityQuery
	query.TypeName = "CargoEntities.Parameter"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Parameter **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_type")
	query.Fields = append(query.Fields, "M_isArray")

	/** associations of Parameter **/
	query.Fields = append(query.Fields, "M_parametersPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ParameterInfo []interface{}

	ParameterInfo = append(ParameterInfo, this.GetUuid())
	if this.parentPtr != nil {
		ParameterInfo = append(ParameterInfo, this.parentPtr.GetUuid())
	} else {
		ParameterInfo = append(ParameterInfo, "")
	}

	/** members of Parameter **/
	ParameterInfo = append(ParameterInfo, this.object.M_name)
	ParameterInfo = append(ParameterInfo, this.object.M_type)
	ParameterInfo = append(ParameterInfo, this.object.M_isArray)

	/** associations of Parameter **/

	/** Save parameters type Parameter **/
	/** attribute Parameter has no method GetId, must be an error here...*/
	ParameterInfo = append(ParameterInfo, "")
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ParameterInfo = append(ParameterInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ParameterInfo = append(ParameterInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), ParameterInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), ParameterInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_ParameterEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_ParameterEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Parameter"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Parameter **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_type")
	query.Fields = append(query.Fields, "M_isArray")

	/** associations of Parameter **/
	query.Fields = append(query.Fields, "M_parametersPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Parameter...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Parameter)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Parameter"

		this.parentUuid = results[0][1].(string)

		/** members of Parameter **/

		/** name **/
		if results[0][2] != nil {
			this.object.M_name = results[0][2].(string)
		}

		/** type **/
		if results[0][3] != nil {
			this.object.M_type = results[0][3].(string)
		}

		/** isArray **/
		if results[0][4] != nil {
			this.object.M_isArray = results[0][4].(bool)
		}

		/** associations of Parameter **/

		/** parametersPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Parameter"
				id_ := refTypeName + "$$" + id
				this.object.M_parametersPtr = id
				GetServer().GetEntityManager().appendReference("parametersPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][6].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][7].(string)
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
func (this *EntityManager) NewCargoEntitiesParameterEntityFromObject(object *CargoEntities.Parameter) *CargoEntities_ParameterEntity {
	return this.NewCargoEntitiesParameterEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_ParameterEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesParameterExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Parameter"
	query.Indexs = append(query.Indexs, "M_name="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_ParameterEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_ParameterEntity) AppendReference(reference Entity) {

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
//              			Action
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_ActionEntity struct {
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
	object         *CargoEntities.Action
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesActionEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_ActionEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesActionExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Action).TYPENAME = "CargoEntities.Action"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Action", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Action).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Action).UUID
			}
			return val.(*CargoEntities_ActionEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_ActionEntity)
	if object == nil {
		entity.object = new(CargoEntities.Action)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Action)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Action"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_ActionEntity) GetTypeName() string {
	return "CargoEntities.Action"
}
func (this *CargoEntities_ActionEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_ActionEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_ActionEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_ActionEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_ActionEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_ActionEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_ActionEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_ActionEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_ActionEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_ActionEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_ActionEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_ActionEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_ActionEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_ActionEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_ActionEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_ActionEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_ActionEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_ActionEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_ActionEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_ActionEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_ActionEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_ActionEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_ActionEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Action"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_ActionEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ActionEntityPrototype() {

	var actionEntityProto EntityPrototype
	actionEntityProto.TypeName = "CargoEntities.Action"
	actionEntityProto.Ids = append(actionEntityProto.Ids, "uuid")
	actionEntityProto.Fields = append(actionEntityProto.Fields, "uuid")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.string")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 0)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)
	actionEntityProto.Indexs = append(actionEntityProto.Indexs, "parentUuid")
	actionEntityProto.Fields = append(actionEntityProto.Fields, "parentUuid")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.string")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 1)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)

	/** members of Action **/
	actionEntityProto.Ids = append(actionEntityProto.Ids, "M_name")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 2)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_name")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.ID")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 3)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_parameters")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "[]CargoEntities.Parameter")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 4)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_results")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "[]CargoEntities.Parameter")

	/** associations of Action **/
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 5)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_entitiesPtr")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	actionEntityProto.Fields = append(actionEntityProto.Fields, "childsUuid")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "[]xs.string")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 6)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)

	actionEntityProto.Fields = append(actionEntityProto.Fields, "referenced")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "[]EntityRef")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 7)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&actionEntityProto)

}

/** Create **/
func (this *CargoEntities_ActionEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Action"

	var query EntityQuery
	query.TypeName = "CargoEntities.Action"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Action **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_parameters")
	query.Fields = append(query.Fields, "M_results")

	/** associations of Action **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ActionInfo []interface{}

	ActionInfo = append(ActionInfo, this.GetUuid())
	if this.parentPtr != nil {
		ActionInfo = append(ActionInfo, this.parentPtr.GetUuid())
	} else {
		ActionInfo = append(ActionInfo, "")
	}

	/** members of Action **/
	ActionInfo = append(ActionInfo, this.object.M_name)

	/** Save parameters type Parameter **/
	parametersIds := make([]string, 0)
	for i := 0; i < len(this.object.M_parameters); i++ {
		parametersEntity := GetServer().GetEntityManager().NewCargoEntitiesParameterEntity(this.GetUuid(), this.object.M_parameters[i].UUID, this.object.M_parameters[i])
		parametersIds = append(parametersIds, parametersEntity.uuid)
		parametersEntity.AppendReferenced("parameters", this)
		this.AppendChild("parameters", parametersEntity)
		if parametersEntity.NeedSave() {
			parametersEntity.SaveEntity()
		}
	}
	parametersStr, _ := json.Marshal(parametersIds)
	ActionInfo = append(ActionInfo, string(parametersStr))

	/** Save results type Parameter **/
	resultsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_results); i++ {
		resultsEntity := GetServer().GetEntityManager().NewCargoEntitiesParameterEntity(this.GetUuid(), this.object.M_results[i].UUID, this.object.M_results[i])
		resultsIds = append(resultsIds, resultsEntity.uuid)
		resultsEntity.AppendReferenced("results", this)
		this.AppendChild("results", resultsEntity)
		if resultsEntity.NeedSave() {
			resultsEntity.SaveEntity()
		}
	}
	resultsStr, _ := json.Marshal(resultsIds)
	ActionInfo = append(ActionInfo, string(resultsStr))

	/** associations of Action **/

	/** Save entities type Entities **/
	ActionInfo = append(ActionInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ActionInfo = append(ActionInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ActionInfo = append(ActionInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), ActionInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), ActionInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_ActionEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_ActionEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Action"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Action **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_parameters")
	query.Fields = append(query.Fields, "M_results")

	/** associations of Action **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Action...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Action)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Action"

		this.parentUuid = results[0][1].(string)

		/** members of Action **/

		/** name **/
		if results[0][2] != nil {
			this.object.M_name = results[0][2].(string)
		}

		/** parameters **/
		if results[0][3] != nil {
			uuidsStr := results[0][3].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var parametersEntity *CargoEntities_ParameterEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						parametersEntity = instance.(*CargoEntities_ParameterEntity)
					} else {
						parametersEntity = GetServer().GetEntityManager().NewCargoEntitiesParameterEntity(this.GetUuid(), uuids[i], nil)
						parametersEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(parametersEntity)
					}
					parametersEntity.AppendReferenced("parameters", this)
					this.AppendChild("parameters", parametersEntity)
				}
			}
		}

		/** results **/
		if results[0][4] != nil {
			uuidsStr := results[0][4].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var resultsEntity *CargoEntities_ParameterEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						resultsEntity = instance.(*CargoEntities_ParameterEntity)
					} else {
						resultsEntity = GetServer().GetEntityManager().NewCargoEntitiesParameterEntity(this.GetUuid(), uuids[i], nil)
						resultsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(resultsEntity)
					}
					resultsEntity.AppendReferenced("results", this)
					this.AppendChild("results", resultsEntity)
				}
			}
		}

		/** associations of Action **/

		/** entitiesPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][6].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][7].(string)
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
func (this *EntityManager) NewCargoEntitiesActionEntityFromObject(object *CargoEntities.Action) *CargoEntities_ActionEntity {
	return this.NewCargoEntitiesActionEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_ActionEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesActionExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Action"
	query.Indexs = append(query.Indexs, "M_name="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_ActionEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_ActionEntity) AppendReference(reference Entity) {

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
//              			Error
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_ErrorEntity struct {
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
	object         *CargoEntities.Error
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesErrorEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_ErrorEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesErrorExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Error).TYPENAME = "CargoEntities.Error"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Error", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Error).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Error).UUID
			}
			return val.(*CargoEntities_ErrorEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_ErrorEntity)
	if object == nil {
		entity.object = new(CargoEntities.Error)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Error)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Error"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_ErrorEntity) GetTypeName() string {
	return "CargoEntities.Error"
}
func (this *CargoEntities_ErrorEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_ErrorEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_ErrorEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_ErrorEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_ErrorEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_ErrorEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_ErrorEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_ErrorEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_ErrorEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_ErrorEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_ErrorEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_ErrorEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_ErrorEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_ErrorEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_ErrorEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_ErrorEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_ErrorEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_ErrorEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_ErrorEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_ErrorEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_ErrorEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_ErrorEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_ErrorEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Error"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_ErrorEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ErrorEntityPrototype() {

	var errorEntityProto EntityPrototype
	errorEntityProto.TypeName = "CargoEntities.Error"
	errorEntityProto.SuperTypeNames = append(errorEntityProto.SuperTypeNames, "CargoEntities.Entity")
	errorEntityProto.SuperTypeNames = append(errorEntityProto.SuperTypeNames, "CargoEntities.Message")
	errorEntityProto.Ids = append(errorEntityProto.Ids, "uuid")
	errorEntityProto.Fields = append(errorEntityProto.Fields, "uuid")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 0)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)
	errorEntityProto.Indexs = append(errorEntityProto.Indexs, "parentUuid")
	errorEntityProto.Fields = append(errorEntityProto.Fields, "parentUuid")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 1)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	errorEntityProto.Ids = append(errorEntityProto.Ids, "M_id")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 2)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_id")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.ID")

	/** members of Message **/
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 3)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_body")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")

	/** members of Error **/
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 4)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_errorPath")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 5)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_code")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.int")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 6)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_accountRef")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "CargoEntities.Account:Ref")

	/** associations of Error **/
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 7)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_entitiesPtr")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	errorEntityProto.Fields = append(errorEntityProto.Fields, "childsUuid")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "[]xs.string")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 8)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)

	errorEntityProto.Fields = append(errorEntityProto.Fields, "referenced")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "[]EntityRef")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 9)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&errorEntityProto)

}

/** Create **/
func (this *CargoEntities_ErrorEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Error"

	var query EntityQuery
	query.TypeName = "CargoEntities.Error"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of Error **/
	query.Fields = append(query.Fields, "M_errorPath")
	query.Fields = append(query.Fields, "M_code")
	query.Fields = append(query.Fields, "M_accountRef")

	/** associations of Error **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ErrorInfo []interface{}

	ErrorInfo = append(ErrorInfo, this.GetUuid())
	if this.parentPtr != nil {
		ErrorInfo = append(ErrorInfo, this.parentPtr.GetUuid())
	} else {
		ErrorInfo = append(ErrorInfo, "")
	}

	/** members of Entity **/
	ErrorInfo = append(ErrorInfo, this.object.M_id)

	/** members of Message **/
	ErrorInfo = append(ErrorInfo, this.object.M_body)

	/** members of Error **/
	ErrorInfo = append(ErrorInfo, this.object.M_errorPath)
	ErrorInfo = append(ErrorInfo, this.object.M_code)

	/** Save accountRef type Account **/
	ErrorInfo = append(ErrorInfo, this.object.M_accountRef)

	/** associations of Error **/

	/** Save entities type Entities **/
	ErrorInfo = append(ErrorInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ErrorInfo = append(ErrorInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ErrorInfo = append(ErrorInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), ErrorInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), ErrorInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_ErrorEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_ErrorEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Error"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of Error **/
	query.Fields = append(query.Fields, "M_errorPath")
	query.Fields = append(query.Fields, "M_code")
	query.Fields = append(query.Fields, "M_accountRef")

	/** associations of Error **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Error...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Error)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Error"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of Message **/

		/** body **/
		if results[0][3] != nil {
			this.object.M_body = results[0][3].(string)
		}

		/** members of Error **/

		/** errorPath **/
		if results[0][4] != nil {
			this.object.M_errorPath = results[0][4].(string)
		}

		/** code **/
		if results[0][5] != nil {
			this.object.M_code = results[0][5].(int)
		}

		/** accountRef **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_accountRef = id
				GetServer().GetEntityManager().appendReference("accountRef", this.object.UUID, id_)
			}
		}

		/** associations of Error **/

		/** entitiesPtr **/
		if results[0][7] != nil {
			id := results[0][7].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][8].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][9].(string)
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
func (this *EntityManager) NewCargoEntitiesErrorEntityFromObject(object *CargoEntities.Error) *CargoEntities_ErrorEntity {
	return this.NewCargoEntitiesErrorEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_ErrorEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesErrorExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Error"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_ErrorEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_ErrorEntity) AppendReference(reference Entity) {

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
//              			LogEntry
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_LogEntryEntity struct {
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
	object         *CargoEntities.LogEntry
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesLogEntryEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_LogEntryEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesLogEntryExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.LogEntry).TYPENAME = "CargoEntities.LogEntry"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.LogEntry", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.LogEntry).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.LogEntry).UUID
			}
			return val.(*CargoEntities_LogEntryEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_LogEntryEntity)
	if object == nil {
		entity.object = new(CargoEntities.LogEntry)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.LogEntry)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.LogEntry"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_LogEntryEntity) GetTypeName() string {
	return "CargoEntities.LogEntry"
}
func (this *CargoEntities_LogEntryEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_LogEntryEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_LogEntryEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_LogEntryEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_LogEntryEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_LogEntryEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_LogEntryEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_LogEntryEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_LogEntryEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_LogEntryEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_LogEntryEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_LogEntryEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_LogEntryEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_LogEntryEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_LogEntryEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_LogEntryEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_LogEntryEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_LogEntryEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_LogEntryEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_LogEntryEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_LogEntryEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_LogEntryEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_LogEntryEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.LogEntry"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_LogEntryEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_LogEntryEntityPrototype() {

	var logEntryEntityProto EntityPrototype
	logEntryEntityProto.TypeName = "CargoEntities.LogEntry"
	logEntryEntityProto.SuperTypeNames = append(logEntryEntityProto.SuperTypeNames, "CargoEntities.Entity")
	logEntryEntityProto.Ids = append(logEntryEntityProto.Ids, "uuid")
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "uuid")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.string")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 0)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.Indexs = append(logEntryEntityProto.Indexs, "parentUuid")
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "parentUuid")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.string")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 1)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	logEntryEntityProto.Ids = append(logEntryEntityProto.Ids, "M_id")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 2)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, true)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_id")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.ID")

	/** members of LogEntry **/
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 3)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, true)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_creationTime")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.long")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 4)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, true)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_entityRef")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "CargoEntities.Entity:Ref")

	/** associations of LogEntry **/
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 5)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_loggerPtr")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "CargoEntities.Log:Ref")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 6)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_entitiesPtr")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "childsUuid")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "[]xs.string")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 7)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)

	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "referenced")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "[]EntityRef")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 8)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&logEntryEntityProto)

}

/** Create **/
func (this *CargoEntities_LogEntryEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.LogEntry"

	var query EntityQuery
	query.TypeName = "CargoEntities.LogEntry"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of LogEntry **/
	query.Fields = append(query.Fields, "M_creationTime")
	query.Fields = append(query.Fields, "M_entityRef")

	/** associations of LogEntry **/
	query.Fields = append(query.Fields, "M_loggerPtr")
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var LogEntryInfo []interface{}

	LogEntryInfo = append(LogEntryInfo, this.GetUuid())
	if this.parentPtr != nil {
		LogEntryInfo = append(LogEntryInfo, this.parentPtr.GetUuid())
	} else {
		LogEntryInfo = append(LogEntryInfo, "")
	}

	/** members of Entity **/
	LogEntryInfo = append(LogEntryInfo, this.object.M_id)

	/** members of LogEntry **/
	LogEntryInfo = append(LogEntryInfo, this.object.M_creationTime)

	/** Save entityRef type Entity **/
	LogEntryInfo = append(LogEntryInfo, this.object.M_entityRef)

	/** associations of LogEntry **/

	/** Save logger type Log **/
	LogEntryInfo = append(LogEntryInfo, this.object.M_loggerPtr)

	/** Save entities type Entities **/
	LogEntryInfo = append(LogEntryInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	LogEntryInfo = append(LogEntryInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	LogEntryInfo = append(LogEntryInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), LogEntryInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), LogEntryInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_LogEntryEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_LogEntryEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.LogEntry"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of LogEntry **/
	query.Fields = append(query.Fields, "M_creationTime")
	query.Fields = append(query.Fields, "M_entityRef")

	/** associations of LogEntry **/
	query.Fields = append(query.Fields, "M_loggerPtr")
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of LogEntry...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.LogEntry)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.LogEntry"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of LogEntry **/

		/** creationTime **/
		if results[0][3] != nil {
			this.object.M_creationTime = results[0][3].(int64)
		}

		/** entityRef **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entity"
				id_ := refTypeName + "$$" + id
				this.object.M_entityRef = id
				GetServer().GetEntityManager().appendReference("entityRef", this.object.UUID, id_)
			}
		}

		/** associations of LogEntry **/

		/** loggerPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Log"
				id_ := refTypeName + "$$" + id
				this.object.M_loggerPtr = id
				GetServer().GetEntityManager().appendReference("loggerPtr", this.object.UUID, id_)
			}
		}

		/** entitiesPtr **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][7].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][8].(string)
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
func (this *EntityManager) NewCargoEntitiesLogEntryEntityFromObject(object *CargoEntities.LogEntry) *CargoEntities_LogEntryEntity {
	return this.NewCargoEntitiesLogEntryEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_LogEntryEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesLogEntryExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.LogEntry"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_LogEntryEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_LogEntryEntity) AppendReference(reference Entity) {

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
//              			Log
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_LogEntity struct {
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
	object         *CargoEntities.Log
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesLogEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_LogEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesLogExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Log).TYPENAME = "CargoEntities.Log"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Log", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Log).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Log).UUID
			}
			return val.(*CargoEntities_LogEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_LogEntity)
	if object == nil {
		entity.object = new(CargoEntities.Log)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Log)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Log"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_LogEntity) GetTypeName() string {
	return "CargoEntities.Log"
}
func (this *CargoEntities_LogEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_LogEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_LogEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_LogEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_LogEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_LogEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_LogEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_LogEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_LogEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_LogEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_LogEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_LogEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_LogEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_LogEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_LogEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_LogEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_LogEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_LogEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_LogEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_LogEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_LogEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_LogEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_LogEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Log"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_LogEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_LogEntityPrototype() {

	var logEntityProto EntityPrototype
	logEntityProto.TypeName = "CargoEntities.Log"
	logEntityProto.SuperTypeNames = append(logEntityProto.SuperTypeNames, "CargoEntities.Entity")
	logEntityProto.Ids = append(logEntityProto.Ids, "uuid")
	logEntityProto.Fields = append(logEntityProto.Fields, "uuid")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.string")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 0)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)
	logEntityProto.Indexs = append(logEntityProto.Indexs, "parentUuid")
	logEntityProto.Fields = append(logEntityProto.Fields, "parentUuid")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.string")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 1)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	logEntityProto.Ids = append(logEntityProto.Ids, "M_id")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 2)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, true)
	logEntityProto.Fields = append(logEntityProto.Fields, "M_id")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.ID")

	/** members of Log **/
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 3)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, true)
	logEntityProto.Fields = append(logEntityProto.Fields, "M_entries")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "[]CargoEntities.LogEntry")

	/** associations of Log **/
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 4)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)
	logEntityProto.Fields = append(logEntityProto.Fields, "M_entitiesPtr")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	logEntityProto.Fields = append(logEntityProto.Fields, "childsUuid")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "[]xs.string")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 5)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)

	logEntityProto.Fields = append(logEntityProto.Fields, "referenced")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "[]EntityRef")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 6)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&logEntityProto)

}

/** Create **/
func (this *CargoEntities_LogEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Log"

	var query EntityQuery
	query.TypeName = "CargoEntities.Log"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Log **/
	query.Fields = append(query.Fields, "M_entries")

	/** associations of Log **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var LogInfo []interface{}

	LogInfo = append(LogInfo, this.GetUuid())
	if this.parentPtr != nil {
		LogInfo = append(LogInfo, this.parentPtr.GetUuid())
	} else {
		LogInfo = append(LogInfo, "")
	}

	/** members of Entity **/
	LogInfo = append(LogInfo, this.object.M_id)

	/** members of Log **/

	/** Save entries type LogEntry **/
	entriesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_entries); i++ {
		entriesEntity := GetServer().GetEntityManager().NewCargoEntitiesLogEntryEntity(this.GetUuid(), this.object.M_entries[i].UUID, this.object.M_entries[i])
		entriesIds = append(entriesIds, entriesEntity.uuid)
		entriesEntity.AppendReferenced("entries", this)
		this.AppendChild("entries", entriesEntity)
		if entriesEntity.NeedSave() {
			entriesEntity.SaveEntity()
		}
	}
	entriesStr, _ := json.Marshal(entriesIds)
	LogInfo = append(LogInfo, string(entriesStr))

	/** associations of Log **/

	/** Save entities type Entities **/
	LogInfo = append(LogInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	LogInfo = append(LogInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	LogInfo = append(LogInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), LogInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), LogInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_LogEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_LogEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Log"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Log **/
	query.Fields = append(query.Fields, "M_entries")

	/** associations of Log **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Log...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Log)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Log"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of Log **/

		/** entries **/
		if results[0][3] != nil {
			uuidsStr := results[0][3].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var entriesEntity *CargoEntities_LogEntryEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						entriesEntity = instance.(*CargoEntities_LogEntryEntity)
					} else {
						entriesEntity = GetServer().GetEntityManager().NewCargoEntitiesLogEntryEntity(this.GetUuid(), uuids[i], nil)
						entriesEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(entriesEntity)
					}
					entriesEntity.AppendReferenced("entries", this)
					this.AppendChild("entries", entriesEntity)
				}
			}
		}

		/** associations of Log **/

		/** entitiesPtr **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
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
func (this *EntityManager) NewCargoEntitiesLogEntityFromObject(object *CargoEntities.Log) *CargoEntities_LogEntity {
	return this.NewCargoEntitiesLogEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_LogEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesLogExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Log"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_LogEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_LogEntity) AppendReference(reference Entity) {

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
//              			Project
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_ProjectEntity struct {
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
	object         *CargoEntities.Project
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesProjectEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_ProjectEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesProjectExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Project).TYPENAME = "CargoEntities.Project"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Project", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Project).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Project).UUID
			}
			return val.(*CargoEntities_ProjectEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_ProjectEntity)
	if object == nil {
		entity.object = new(CargoEntities.Project)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Project)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Project"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_ProjectEntity) GetTypeName() string {
	return "CargoEntities.Project"
}
func (this *CargoEntities_ProjectEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_ProjectEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_ProjectEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_ProjectEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_ProjectEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_ProjectEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_ProjectEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_ProjectEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_ProjectEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_ProjectEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_ProjectEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_ProjectEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_ProjectEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_ProjectEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_ProjectEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_ProjectEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_ProjectEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_ProjectEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_ProjectEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_ProjectEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_ProjectEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_ProjectEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_ProjectEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Project"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_ProjectEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ProjectEntityPrototype() {

	var projectEntityProto EntityPrototype
	projectEntityProto.TypeName = "CargoEntities.Project"
	projectEntityProto.SuperTypeNames = append(projectEntityProto.SuperTypeNames, "CargoEntities.Entity")
	projectEntityProto.Ids = append(projectEntityProto.Ids, "uuid")
	projectEntityProto.Fields = append(projectEntityProto.Fields, "uuid")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 0)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)
	projectEntityProto.Indexs = append(projectEntityProto.Indexs, "parentUuid")
	projectEntityProto.Fields = append(projectEntityProto.Fields, "parentUuid")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 1)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	projectEntityProto.Ids = append(projectEntityProto.Ids, "M_id")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 2)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, true)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_id")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.ID")

	/** members of Project **/
	projectEntityProto.Indexs = append(projectEntityProto.Indexs, "M_name")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 3)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, true)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_name")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 4)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, true)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_filesRef")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "[]CargoEntities.File:Ref")

	/** associations of Project **/
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 5)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_entitiesPtr")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	projectEntityProto.Fields = append(projectEntityProto.Fields, "childsUuid")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "[]xs.string")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 6)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)

	projectEntityProto.Fields = append(projectEntityProto.Fields, "referenced")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "[]EntityRef")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 7)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&projectEntityProto)

}

/** Create **/
func (this *CargoEntities_ProjectEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Project"

	var query EntityQuery
	query.TypeName = "CargoEntities.Project"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Project **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_filesRef")

	/** associations of Project **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ProjectInfo []interface{}

	ProjectInfo = append(ProjectInfo, this.GetUuid())
	if this.parentPtr != nil {
		ProjectInfo = append(ProjectInfo, this.parentPtr.GetUuid())
	} else {
		ProjectInfo = append(ProjectInfo, "")
	}

	/** members of Entity **/
	ProjectInfo = append(ProjectInfo, this.object.M_id)

	/** members of Project **/
	ProjectInfo = append(ProjectInfo, this.object.M_name)

	/** Save filesRef type File **/
	filesRefStr, _ := json.Marshal(this.object.M_filesRef)
	ProjectInfo = append(ProjectInfo, string(filesRefStr))

	/** associations of Project **/

	/** Save entities type Entities **/
	ProjectInfo = append(ProjectInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ProjectInfo = append(ProjectInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ProjectInfo = append(ProjectInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), ProjectInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), ProjectInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_ProjectEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_ProjectEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Project"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Project **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_filesRef")

	/** associations of Project **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Project...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Project)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Project"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of Project **/

		/** name **/
		if results[0][3] != nil {
			this.object.M_name = results[0][3].(string)
		}

		/** filesRef **/
		if results[0][4] != nil {
			idsStr := results[0][4].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.File"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_filesRef = append(this.object.M_filesRef, ids[i])
					GetServer().GetEntityManager().appendReference("filesRef", this.object.UUID, id_)
				}
			}
		}

		/** associations of Project **/

		/** entitiesPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][6].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][7].(string)
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
func (this *EntityManager) NewCargoEntitiesProjectEntityFromObject(object *CargoEntities.Project) *CargoEntities_ProjectEntity {
	return this.NewCargoEntitiesProjectEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_ProjectEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesProjectExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Project"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_ProjectEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_ProjectEntity) AppendReference(reference Entity) {

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

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_MessageEntityPrototype() {

	var messageEntityProto EntityPrototype
	messageEntityProto.TypeName = "CargoEntities.Message"
	messageEntityProto.IsAbstract = true
	messageEntityProto.SuperTypeNames = append(messageEntityProto.SuperTypeNames, "CargoEntities.Entity")
	messageEntityProto.SubstitutionGroup = append(messageEntityProto.SubstitutionGroup, "CargoEntities.Error")
	messageEntityProto.SubstitutionGroup = append(messageEntityProto.SubstitutionGroup, "CargoEntities.Notification")
	messageEntityProto.SubstitutionGroup = append(messageEntityProto.SubstitutionGroup, "CargoEntities.TextMessage")
	messageEntityProto.Ids = append(messageEntityProto.Ids, "uuid")
	messageEntityProto.Fields = append(messageEntityProto.Fields, "uuid")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 0)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)
	messageEntityProto.Indexs = append(messageEntityProto.Indexs, "parentUuid")
	messageEntityProto.Fields = append(messageEntityProto.Fields, "parentUuid")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 1)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	messageEntityProto.Ids = append(messageEntityProto.Ids, "M_id")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 2)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, true)
	messageEntityProto.Fields = append(messageEntityProto.Fields, "M_id")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.ID")

	/** members of Message **/
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 3)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, true)
	messageEntityProto.Fields = append(messageEntityProto.Fields, "M_body")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")

	/** associations of Message **/
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 4)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)
	messageEntityProto.Fields = append(messageEntityProto.Fields, "M_entitiesPtr")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	messageEntityProto.Fields = append(messageEntityProto.Fields, "childsUuid")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "[]xs.string")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 5)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)

	messageEntityProto.Fields = append(messageEntityProto.Fields, "referenced")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "[]EntityRef")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 6)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&messageEntityProto)

}

////////////////////////////////////////////////////////////////////////////////
//              			Notification
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_NotificationEntity struct {
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
	object         *CargoEntities.Notification
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesNotificationEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_NotificationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesNotificationExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Notification).TYPENAME = "CargoEntities.Notification"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Notification", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Notification).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Notification).UUID
			}
			return val.(*CargoEntities_NotificationEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_NotificationEntity)
	if object == nil {
		entity.object = new(CargoEntities.Notification)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Notification)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Notification"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_NotificationEntity) GetTypeName() string {
	return "CargoEntities.Notification"
}
func (this *CargoEntities_NotificationEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_NotificationEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_NotificationEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_NotificationEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_NotificationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_NotificationEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_NotificationEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_NotificationEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_NotificationEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_NotificationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_NotificationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_NotificationEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_NotificationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_NotificationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_NotificationEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_NotificationEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_NotificationEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_NotificationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_NotificationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_NotificationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_NotificationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_NotificationEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_NotificationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Notification"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_NotificationEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_NotificationEntityPrototype() {

	var notificationEntityProto EntityPrototype
	notificationEntityProto.TypeName = "CargoEntities.Notification"
	notificationEntityProto.SuperTypeNames = append(notificationEntityProto.SuperTypeNames, "CargoEntities.Entity")
	notificationEntityProto.SuperTypeNames = append(notificationEntityProto.SuperTypeNames, "CargoEntities.Message")
	notificationEntityProto.Ids = append(notificationEntityProto.Ids, "uuid")
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "uuid")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 0)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)
	notificationEntityProto.Indexs = append(notificationEntityProto.Indexs, "parentUuid")
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "parentUuid")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 1)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	notificationEntityProto.Ids = append(notificationEntityProto.Ids, "M_id")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 2)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_id")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.ID")

	/** members of Message **/
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 3)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_body")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")

	/** members of Notification **/
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 4)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_fromRef")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "CargoEntities.Account:Ref")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 5)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_toRef")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "CargoEntities.Account:Ref")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 6)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_type")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 7)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_code")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.int")

	/** associations of Notification **/
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 8)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_entitiesPtr")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "childsUuid")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "[]xs.string")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 9)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)

	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "referenced")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "[]EntityRef")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 10)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&notificationEntityProto)

}

/** Create **/
func (this *CargoEntities_NotificationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Notification"

	var query EntityQuery
	query.TypeName = "CargoEntities.Notification"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of Notification **/
	query.Fields = append(query.Fields, "M_fromRef")
	query.Fields = append(query.Fields, "M_toRef")
	query.Fields = append(query.Fields, "M_type")
	query.Fields = append(query.Fields, "M_code")

	/** associations of Notification **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var NotificationInfo []interface{}

	NotificationInfo = append(NotificationInfo, this.GetUuid())
	if this.parentPtr != nil {
		NotificationInfo = append(NotificationInfo, this.parentPtr.GetUuid())
	} else {
		NotificationInfo = append(NotificationInfo, "")
	}

	/** members of Entity **/
	NotificationInfo = append(NotificationInfo, this.object.M_id)

	/** members of Message **/
	NotificationInfo = append(NotificationInfo, this.object.M_body)

	/** members of Notification **/

	/** Save fromRef type Account **/
	NotificationInfo = append(NotificationInfo, this.object.M_fromRef)

	/** Save toRef type Account **/
	NotificationInfo = append(NotificationInfo, this.object.M_toRef)
	NotificationInfo = append(NotificationInfo, this.object.M_type)
	NotificationInfo = append(NotificationInfo, this.object.M_code)

	/** associations of Notification **/

	/** Save entities type Entities **/
	NotificationInfo = append(NotificationInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	NotificationInfo = append(NotificationInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	NotificationInfo = append(NotificationInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), NotificationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), NotificationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_NotificationEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_NotificationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Notification"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of Notification **/
	query.Fields = append(query.Fields, "M_fromRef")
	query.Fields = append(query.Fields, "M_toRef")
	query.Fields = append(query.Fields, "M_type")
	query.Fields = append(query.Fields, "M_code")

	/** associations of Notification **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Notification...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Notification)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Notification"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of Message **/

		/** body **/
		if results[0][3] != nil {
			this.object.M_body = results[0][3].(string)
		}

		/** members of Notification **/

		/** fromRef **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_fromRef = id
				GetServer().GetEntityManager().appendReference("fromRef", this.object.UUID, id_)
			}
		}

		/** toRef **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_toRef = id
				GetServer().GetEntityManager().appendReference("toRef", this.object.UUID, id_)
			}
		}

		/** type **/
		if results[0][6] != nil {
			this.object.M_type = results[0][6].(string)
		}

		/** code **/
		if results[0][7] != nil {
			this.object.M_code = results[0][7].(int)
		}

		/** associations of Notification **/

		/** entitiesPtr **/
		if results[0][8] != nil {
			id := results[0][8].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
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
func (this *EntityManager) NewCargoEntitiesNotificationEntityFromObject(object *CargoEntities.Notification) *CargoEntities_NotificationEntity {
	return this.NewCargoEntitiesNotificationEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_NotificationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesNotificationExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Notification"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_NotificationEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_NotificationEntity) AppendReference(reference Entity) {

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
//              			TextMessage
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_TextMessageEntity struct {
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
	object         *CargoEntities.TextMessage
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesTextMessageEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_TextMessageEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesTextMessageExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.TextMessage).TYPENAME = "CargoEntities.TextMessage"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.TextMessage", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.TextMessage).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.TextMessage).UUID
			}
			return val.(*CargoEntities_TextMessageEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_TextMessageEntity)
	if object == nil {
		entity.object = new(CargoEntities.TextMessage)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.TextMessage)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.TextMessage"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_TextMessageEntity) GetTypeName() string {
	return "CargoEntities.TextMessage"
}
func (this *CargoEntities_TextMessageEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_TextMessageEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_TextMessageEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_TextMessageEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_TextMessageEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_TextMessageEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_TextMessageEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_TextMessageEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_TextMessageEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_TextMessageEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_TextMessageEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_TextMessageEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_TextMessageEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_TextMessageEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_TextMessageEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_TextMessageEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_TextMessageEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_TextMessageEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_TextMessageEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_TextMessageEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_TextMessageEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_TextMessageEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_TextMessageEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.TextMessage"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_TextMessageEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_TextMessageEntityPrototype() {

	var textMessageEntityProto EntityPrototype
	textMessageEntityProto.TypeName = "CargoEntities.TextMessage"
	textMessageEntityProto.SuperTypeNames = append(textMessageEntityProto.SuperTypeNames, "CargoEntities.Entity")
	textMessageEntityProto.SuperTypeNames = append(textMessageEntityProto.SuperTypeNames, "CargoEntities.Message")
	textMessageEntityProto.Ids = append(textMessageEntityProto.Ids, "uuid")
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "uuid")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 0)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)
	textMessageEntityProto.Indexs = append(textMessageEntityProto.Indexs, "parentUuid")
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "parentUuid")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 1)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	textMessageEntityProto.Ids = append(textMessageEntityProto.Ids, "M_id")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 2)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_id")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.ID")

	/** members of Message **/
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 3)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_body")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")

	/** members of TextMessage **/
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 4)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_creationTime")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.long")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 5)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_fromRef")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "CargoEntities.Account:Ref")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 6)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_toRef")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "CargoEntities.Account:Ref")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 7)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_title")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")

	/** associations of TextMessage **/
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 8)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_entitiesPtr")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "childsUuid")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "[]xs.string")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 9)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)

	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "referenced")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "[]EntityRef")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 10)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&textMessageEntityProto)

}

/** Create **/
func (this *CargoEntities_TextMessageEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.TextMessage"

	var query EntityQuery
	query.TypeName = "CargoEntities.TextMessage"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of TextMessage **/
	query.Fields = append(query.Fields, "M_creationTime")
	query.Fields = append(query.Fields, "M_fromRef")
	query.Fields = append(query.Fields, "M_toRef")
	query.Fields = append(query.Fields, "M_title")

	/** associations of TextMessage **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var TextMessageInfo []interface{}

	TextMessageInfo = append(TextMessageInfo, this.GetUuid())
	if this.parentPtr != nil {
		TextMessageInfo = append(TextMessageInfo, this.parentPtr.GetUuid())
	} else {
		TextMessageInfo = append(TextMessageInfo, "")
	}

	/** members of Entity **/
	TextMessageInfo = append(TextMessageInfo, this.object.M_id)

	/** members of Message **/
	TextMessageInfo = append(TextMessageInfo, this.object.M_body)

	/** members of TextMessage **/
	TextMessageInfo = append(TextMessageInfo, this.object.M_creationTime)

	/** Save fromRef type Account **/
	TextMessageInfo = append(TextMessageInfo, this.object.M_fromRef)

	/** Save toRef type Account **/
	TextMessageInfo = append(TextMessageInfo, this.object.M_toRef)
	TextMessageInfo = append(TextMessageInfo, this.object.M_title)

	/** associations of TextMessage **/

	/** Save entities type Entities **/
	TextMessageInfo = append(TextMessageInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	TextMessageInfo = append(TextMessageInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	TextMessageInfo = append(TextMessageInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), TextMessageInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), TextMessageInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_TextMessageEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_TextMessageEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.TextMessage"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of TextMessage **/
	query.Fields = append(query.Fields, "M_creationTime")
	query.Fields = append(query.Fields, "M_fromRef")
	query.Fields = append(query.Fields, "M_toRef")
	query.Fields = append(query.Fields, "M_title")

	/** associations of TextMessage **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of TextMessage...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.TextMessage)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.TextMessage"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of Message **/

		/** body **/
		if results[0][3] != nil {
			this.object.M_body = results[0][3].(string)
		}

		/** members of TextMessage **/

		/** creationTime **/
		if results[0][4] != nil {
			this.object.M_creationTime = results[0][4].(int64)
		}

		/** fromRef **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_fromRef = id
				GetServer().GetEntityManager().appendReference("fromRef", this.object.UUID, id_)
			}
		}

		/** toRef **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_toRef = id
				GetServer().GetEntityManager().appendReference("toRef", this.object.UUID, id_)
			}
		}

		/** title **/
		if results[0][7] != nil {
			this.object.M_title = results[0][7].(string)
		}

		/** associations of TextMessage **/

		/** entitiesPtr **/
		if results[0][8] != nil {
			id := results[0][8].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
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
func (this *EntityManager) NewCargoEntitiesTextMessageEntityFromObject(object *CargoEntities.TextMessage) *CargoEntities_TextMessageEntity {
	return this.NewCargoEntitiesTextMessageEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_TextMessageEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesTextMessageExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.TextMessage"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_TextMessageEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_TextMessageEntity) AppendReference(reference Entity) {

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
//              			Session
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_SessionEntity struct {
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
	object         *CargoEntities.Session
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesSessionEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_SessionEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesSessionExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Session).TYPENAME = "CargoEntities.Session"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Session", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Session).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Session).UUID
			}
			return val.(*CargoEntities_SessionEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_SessionEntity)
	if object == nil {
		entity.object = new(CargoEntities.Session)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Session)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Session"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_SessionEntity) GetTypeName() string {
	return "CargoEntities.Session"
}
func (this *CargoEntities_SessionEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_SessionEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_SessionEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_SessionEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_SessionEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_SessionEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_SessionEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_SessionEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_SessionEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_SessionEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_SessionEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_SessionEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_SessionEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_SessionEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_SessionEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_SessionEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_SessionEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_SessionEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_SessionEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_SessionEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_SessionEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_SessionEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_SessionEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Session"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_SessionEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_SessionEntityPrototype() {

	var sessionEntityProto EntityPrototype
	sessionEntityProto.TypeName = "CargoEntities.Session"
	sessionEntityProto.Ids = append(sessionEntityProto.Ids, "uuid")
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "uuid")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.string")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 0)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)
	sessionEntityProto.Indexs = append(sessionEntityProto.Indexs, "parentUuid")
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "parentUuid")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.string")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 1)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)

	/** members of Session **/
	sessionEntityProto.Ids = append(sessionEntityProto.Ids, "M_id")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 2)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_id")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.ID")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 3)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_startTime")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.long")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 4)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_endTime")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.long")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 5)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_statusTime")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.long")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 6)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_sessionState")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "enum:SessionState_Online:SessionState_Away:SessionState_Offline")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 7)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_computerRef")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "CargoEntities.Computer:Ref")

	/** associations of Session **/
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 8)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_accountPtr")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "CargoEntities.Account:Ref")
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "childsUuid")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "[]xs.string")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 9)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)

	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "referenced")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "[]EntityRef")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 10)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&sessionEntityProto)

}

/** Create **/
func (this *CargoEntities_SessionEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Session"

	var query EntityQuery
	query.TypeName = "CargoEntities.Session"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Session **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_startTime")
	query.Fields = append(query.Fields, "M_endTime")
	query.Fields = append(query.Fields, "M_statusTime")
	query.Fields = append(query.Fields, "M_sessionState")
	query.Fields = append(query.Fields, "M_computerRef")

	/** associations of Session **/
	query.Fields = append(query.Fields, "M_accountPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var SessionInfo []interface{}

	SessionInfo = append(SessionInfo, this.GetUuid())
	if this.parentPtr != nil {
		SessionInfo = append(SessionInfo, this.parentPtr.GetUuid())
	} else {
		SessionInfo = append(SessionInfo, "")
	}

	/** members of Session **/
	SessionInfo = append(SessionInfo, this.object.M_id)
	SessionInfo = append(SessionInfo, this.object.M_startTime)
	SessionInfo = append(SessionInfo, this.object.M_endTime)
	SessionInfo = append(SessionInfo, this.object.M_statusTime)

	/** Save sessionState type SessionState **/
	if this.object.M_sessionState == CargoEntities.SessionState_Online {
		SessionInfo = append(SessionInfo, 0)
	} else if this.object.M_sessionState == CargoEntities.SessionState_Away {
		SessionInfo = append(SessionInfo, 1)
	} else if this.object.M_sessionState == CargoEntities.SessionState_Offline {
		SessionInfo = append(SessionInfo, 2)
	} else {
		SessionInfo = append(SessionInfo, 0)
	}

	/** Save computerRef type Computer **/
	SessionInfo = append(SessionInfo, this.object.M_computerRef)

	/** associations of Session **/

	/** Save account type Account **/
	SessionInfo = append(SessionInfo, this.object.M_accountPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	SessionInfo = append(SessionInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	SessionInfo = append(SessionInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), SessionInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), SessionInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_SessionEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_SessionEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Session"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Session **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_startTime")
	query.Fields = append(query.Fields, "M_endTime")
	query.Fields = append(query.Fields, "M_statusTime")
	query.Fields = append(query.Fields, "M_sessionState")
	query.Fields = append(query.Fields, "M_computerRef")

	/** associations of Session **/
	query.Fields = append(query.Fields, "M_accountPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Session...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Session)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Session"

		this.parentUuid = results[0][1].(string)

		/** members of Session **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** startTime **/
		if results[0][3] != nil {
			this.object.M_startTime = results[0][3].(int64)
		}

		/** endTime **/
		if results[0][4] != nil {
			this.object.M_endTime = results[0][4].(int64)
		}

		/** statusTime **/
		if results[0][5] != nil {
			this.object.M_statusTime = results[0][5].(int64)
		}

		/** sessionState **/
		if results[0][6] != nil {
			enumIndex := results[0][6].(int)
			if enumIndex == 0 {
				this.object.M_sessionState = CargoEntities.SessionState_Online
			} else if enumIndex == 1 {
				this.object.M_sessionState = CargoEntities.SessionState_Away
			} else if enumIndex == 2 {
				this.object.M_sessionState = CargoEntities.SessionState_Offline
			}
		}

		/** computerRef **/
		if results[0][7] != nil {
			id := results[0][7].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Computer"
				id_ := refTypeName + "$$" + id
				this.object.M_computerRef = id
				GetServer().GetEntityManager().appendReference("computerRef", this.object.UUID, id_)
			}
		}

		/** associations of Session **/

		/** accountPtr **/
		if results[0][8] != nil {
			id := results[0][8].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_accountPtr = id
				GetServer().GetEntityManager().appendReference("accountPtr", this.object.UUID, id_)
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
func (this *EntityManager) NewCargoEntitiesSessionEntityFromObject(object *CargoEntities.Session) *CargoEntities_SessionEntity {
	return this.NewCargoEntitiesSessionEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_SessionEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesSessionExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Session"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_SessionEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_SessionEntity) AppendReference(reference Entity) {

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
//              			Role
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_RoleEntity struct {
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
	object         *CargoEntities.Role
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesRoleEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_RoleEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesRoleExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Role).TYPENAME = "CargoEntities.Role"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Role", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Role).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Role).UUID
			}
			return val.(*CargoEntities_RoleEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_RoleEntity)
	if object == nil {
		entity.object = new(CargoEntities.Role)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Role)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Role"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_RoleEntity) GetTypeName() string {
	return "CargoEntities.Role"
}
func (this *CargoEntities_RoleEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_RoleEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_RoleEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_RoleEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_RoleEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_RoleEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_RoleEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_RoleEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_RoleEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_RoleEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_RoleEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_RoleEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_RoleEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_RoleEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_RoleEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_RoleEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_RoleEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_RoleEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_RoleEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_RoleEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_RoleEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_RoleEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_RoleEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Role"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_RoleEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_RoleEntityPrototype() {

	var roleEntityProto EntityPrototype
	roleEntityProto.TypeName = "CargoEntities.Role"
	roleEntityProto.Ids = append(roleEntityProto.Ids, "uuid")
	roleEntityProto.Fields = append(roleEntityProto.Fields, "uuid")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.string")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 0)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)
	roleEntityProto.Indexs = append(roleEntityProto.Indexs, "parentUuid")
	roleEntityProto.Fields = append(roleEntityProto.Fields, "parentUuid")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.string")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 1)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)

	/** members of Role **/
	roleEntityProto.Ids = append(roleEntityProto.Ids, "M_id")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 2)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, true)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_id")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.ID")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 3)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, true)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_accounts")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "[]CargoEntities.Account:Ref")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 4)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, true)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_actions")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "[]CargoEntities.Action:Ref")

	/** associations of Role **/
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 5)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_entitiesPtr")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	roleEntityProto.Fields = append(roleEntityProto.Fields, "childsUuid")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "[]xs.string")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 6)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)

	roleEntityProto.Fields = append(roleEntityProto.Fields, "referenced")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "[]EntityRef")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 7)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&roleEntityProto)

}

/** Create **/
func (this *CargoEntities_RoleEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Role"

	var query EntityQuery
	query.TypeName = "CargoEntities.Role"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Role **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_accounts")
	query.Fields = append(query.Fields, "M_actions")

	/** associations of Role **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var RoleInfo []interface{}

	RoleInfo = append(RoleInfo, this.GetUuid())
	if this.parentPtr != nil {
		RoleInfo = append(RoleInfo, this.parentPtr.GetUuid())
	} else {
		RoleInfo = append(RoleInfo, "")
	}

	/** members of Role **/
	RoleInfo = append(RoleInfo, this.object.M_id)

	/** Save accounts type Account **/
	accountsStr, _ := json.Marshal(this.object.M_accounts)
	RoleInfo = append(RoleInfo, string(accountsStr))

	/** Save actions type Action **/
	actionsStr, _ := json.Marshal(this.object.M_actions)
	RoleInfo = append(RoleInfo, string(actionsStr))

	/** associations of Role **/

	/** Save entities type Entities **/
	RoleInfo = append(RoleInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	RoleInfo = append(RoleInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	RoleInfo = append(RoleInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), RoleInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), RoleInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_RoleEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_RoleEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Role"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Role **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_accounts")
	query.Fields = append(query.Fields, "M_actions")

	/** associations of Role **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Role...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Role)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Role"

		this.parentUuid = results[0][1].(string)

		/** members of Role **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** accounts **/
		if results[0][3] != nil {
			idsStr := results[0][3].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Account"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_accounts = append(this.object.M_accounts, ids[i])
					GetServer().GetEntityManager().appendReference("accounts", this.object.UUID, id_)
				}
			}
		}

		/** actions **/
		if results[0][4] != nil {
			idsStr := results[0][4].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Action"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_actions = append(this.object.M_actions, ids[i])
					GetServer().GetEntityManager().appendReference("actions", this.object.UUID, id_)
				}
			}
		}

		/** associations of Role **/

		/** entitiesPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][6].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][7].(string)
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
func (this *EntityManager) NewCargoEntitiesRoleEntityFromObject(object *CargoEntities.Role) *CargoEntities_RoleEntity {
	return this.NewCargoEntitiesRoleEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_RoleEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesRoleExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Role"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_RoleEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_RoleEntity) AppendReference(reference Entity) {

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
//              			Account
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_AccountEntity struct {
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
	object         *CargoEntities.Account
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesAccountEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_AccountEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesAccountExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Account).TYPENAME = "CargoEntities.Account"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Account", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Account).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Account).UUID
			}
			return val.(*CargoEntities_AccountEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_AccountEntity)
	if object == nil {
		entity.object = new(CargoEntities.Account)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Account)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Account"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_AccountEntity) GetTypeName() string {
	return "CargoEntities.Account"
}
func (this *CargoEntities_AccountEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_AccountEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_AccountEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_AccountEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_AccountEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_AccountEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_AccountEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_AccountEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_AccountEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_AccountEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_AccountEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_AccountEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_AccountEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_AccountEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_AccountEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_AccountEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_AccountEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_AccountEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_AccountEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_AccountEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_AccountEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_AccountEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_AccountEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Account"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_AccountEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_AccountEntityPrototype() {

	var accountEntityProto EntityPrototype
	accountEntityProto.TypeName = "CargoEntities.Account"
	accountEntityProto.SuperTypeNames = append(accountEntityProto.SuperTypeNames, "CargoEntities.Entity")
	accountEntityProto.Ids = append(accountEntityProto.Ids, "uuid")
	accountEntityProto.Fields = append(accountEntityProto.Fields, "uuid")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 0)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)
	accountEntityProto.Indexs = append(accountEntityProto.Indexs, "parentUuid")
	accountEntityProto.Fields = append(accountEntityProto.Fields, "parentUuid")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 1)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	accountEntityProto.Ids = append(accountEntityProto.Ids, "M_id")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 2)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_id")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.ID")

	/** members of Account **/
	accountEntityProto.Indexs = append(accountEntityProto.Indexs, "M_name")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 3)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_name")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 4)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_password")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 5)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_email")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 6)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_sessions")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Session")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 7)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_permissions")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Permission")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 8)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_messages")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Message")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 9)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_userRef")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "CargoEntities.User:Ref")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 10)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_rolesRef")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Role:Ref")

	/** associations of Account **/
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 11)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_entitiesPtr")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	accountEntityProto.Fields = append(accountEntityProto.Fields, "childsUuid")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 12)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)

	accountEntityProto.Fields = append(accountEntityProto.Fields, "referenced")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]EntityRef")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 13)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&accountEntityProto)

}

/** Create **/
func (this *CargoEntities_AccountEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Account"

	var query EntityQuery
	query.TypeName = "CargoEntities.Account"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Account **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_password")
	query.Fields = append(query.Fields, "M_email")
	query.Fields = append(query.Fields, "M_sessions")
	query.Fields = append(query.Fields, "M_permissions")
	query.Fields = append(query.Fields, "M_messages")
	query.Fields = append(query.Fields, "M_userRef")
	query.Fields = append(query.Fields, "M_rolesRef")

	/** associations of Account **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var AccountInfo []interface{}

	AccountInfo = append(AccountInfo, this.GetUuid())
	if this.parentPtr != nil {
		AccountInfo = append(AccountInfo, this.parentPtr.GetUuid())
	} else {
		AccountInfo = append(AccountInfo, "")
	}

	/** members of Entity **/
	AccountInfo = append(AccountInfo, this.object.M_id)

	/** members of Account **/
	AccountInfo = append(AccountInfo, this.object.M_name)
	AccountInfo = append(AccountInfo, this.object.M_password)
	AccountInfo = append(AccountInfo, this.object.M_email)

	/** Save sessions type Session **/
	sessionsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_sessions); i++ {
		sessionsEntity := GetServer().GetEntityManager().NewCargoEntitiesSessionEntity(this.GetUuid(), this.object.M_sessions[i].UUID, this.object.M_sessions[i])
		sessionsIds = append(sessionsIds, sessionsEntity.uuid)
		sessionsEntity.AppendReferenced("sessions", this)
		this.AppendChild("sessions", sessionsEntity)
		if sessionsEntity.NeedSave() {
			sessionsEntity.SaveEntity()
		}
	}
	sessionsStr, _ := json.Marshal(sessionsIds)
	AccountInfo = append(AccountInfo, string(sessionsStr))

	/** Save permissions type Permission **/
	permissionsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_permissions); i++ {
		permissionsEntity := GetServer().GetEntityManager().NewCargoEntitiesPermissionEntity(this.GetUuid(), this.object.M_permissions[i].UUID, this.object.M_permissions[i])
		permissionsIds = append(permissionsIds, permissionsEntity.uuid)
		permissionsEntity.AppendReferenced("permissions", this)
		this.AppendChild("permissions", permissionsEntity)
		if permissionsEntity.NeedSave() {
			permissionsEntity.SaveEntity()
		}
	}
	permissionsStr, _ := json.Marshal(permissionsIds)
	AccountInfo = append(AccountInfo, string(permissionsStr))

	/** Save messages type Message **/
	messagesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_messages); i++ {
		switch v := this.object.M_messages[i].(type) {
		case *CargoEntities.Notification:
			messagesEntity := GetServer().GetEntityManager().NewCargoEntitiesNotificationEntity(this.GetUuid(), v.UUID, v)
			messagesIds = append(messagesIds, messagesEntity.uuid)
			messagesEntity.AppendReferenced("messages", this)
			this.AppendChild("messages", messagesEntity)
			if messagesEntity.NeedSave() {
				messagesEntity.SaveEntity()
			}
		case *CargoEntities.TextMessage:
			messagesEntity := GetServer().GetEntityManager().NewCargoEntitiesTextMessageEntity(this.GetUuid(), v.UUID, v)
			messagesIds = append(messagesIds, messagesEntity.uuid)
			messagesEntity.AppendReferenced("messages", this)
			this.AppendChild("messages", messagesEntity)
			if messagesEntity.NeedSave() {
				messagesEntity.SaveEntity()
			}
		case *CargoEntities.Error:
			messagesEntity := GetServer().GetEntityManager().NewCargoEntitiesErrorEntity(this.GetUuid(), v.UUID, v)
			messagesIds = append(messagesIds, messagesEntity.uuid)
			messagesEntity.AppendReferenced("messages", this)
			this.AppendChild("messages", messagesEntity)
			if messagesEntity.NeedSave() {
				messagesEntity.SaveEntity()
			}
		}
	}
	messagesStr, _ := json.Marshal(messagesIds)
	AccountInfo = append(AccountInfo, string(messagesStr))

	/** Save userRef type User **/
	AccountInfo = append(AccountInfo, this.object.M_userRef)

	/** Save rolesRef type Role **/
	rolesRefStr, _ := json.Marshal(this.object.M_rolesRef)
	AccountInfo = append(AccountInfo, string(rolesRefStr))

	/** associations of Account **/

	/** Save entities type Entities **/
	AccountInfo = append(AccountInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	AccountInfo = append(AccountInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	AccountInfo = append(AccountInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), AccountInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), AccountInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_AccountEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_AccountEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Account"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Account **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_password")
	query.Fields = append(query.Fields, "M_email")
	query.Fields = append(query.Fields, "M_sessions")
	query.Fields = append(query.Fields, "M_permissions")
	query.Fields = append(query.Fields, "M_messages")
	query.Fields = append(query.Fields, "M_userRef")
	query.Fields = append(query.Fields, "M_rolesRef")

	/** associations of Account **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Account...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Account)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Account"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of Account **/

		/** name **/
		if results[0][3] != nil {
			this.object.M_name = results[0][3].(string)
		}

		/** password **/
		if results[0][4] != nil {
			this.object.M_password = results[0][4].(string)
		}

		/** email **/
		if results[0][5] != nil {
			this.object.M_email = results[0][5].(string)
		}

		/** sessions **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var sessionsEntity *CargoEntities_SessionEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						sessionsEntity = instance.(*CargoEntities_SessionEntity)
					} else {
						sessionsEntity = GetServer().GetEntityManager().NewCargoEntitiesSessionEntity(this.GetUuid(), uuids[i], nil)
						sessionsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(sessionsEntity)
					}
					sessionsEntity.AppendReferenced("sessions", this)
					this.AppendChild("sessions", sessionsEntity)
				}
			}
		}

		/** permissions **/
		if results[0][7] != nil {
			uuidsStr := results[0][7].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var permissionsEntity *CargoEntities_PermissionEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						permissionsEntity = instance.(*CargoEntities_PermissionEntity)
					} else {
						permissionsEntity = GetServer().GetEntityManager().NewCargoEntitiesPermissionEntity(this.GetUuid(), uuids[i], nil)
						permissionsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(permissionsEntity)
					}
					permissionsEntity.AppendReferenced("permissions", this)
					this.AppendChild("permissions", permissionsEntity)
				}
			}
		}

		/** messages **/
		if results[0][8] != nil {
			uuidsStr := results[0][8].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				typeName := uuids[i][0:strings.Index(uuids[i], "%")]
				if err != nil {
					log.Println("type ", typeName, " not found!")
					return err
				}
				if typeName == "CargoEntities.Error" {
					if len(uuids[i]) > 0 {
						var messagesEntity *CargoEntities_ErrorEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							messagesEntity = instance.(*CargoEntities_ErrorEntity)
						} else {
							messagesEntity = GetServer().GetEntityManager().NewCargoEntitiesErrorEntity(this.GetUuid(), uuids[i], nil)
							messagesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(messagesEntity)
						}
						messagesEntity.AppendReferenced("messages", this)
						this.AppendChild("messages", messagesEntity)
					}
				} else if typeName == "CargoEntities.Notification" {
					if len(uuids[i]) > 0 {
						var messagesEntity *CargoEntities_NotificationEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							messagesEntity = instance.(*CargoEntities_NotificationEntity)
						} else {
							messagesEntity = GetServer().GetEntityManager().NewCargoEntitiesNotificationEntity(this.GetUuid(), uuids[i], nil)
							messagesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(messagesEntity)
						}
						messagesEntity.AppendReferenced("messages", this)
						this.AppendChild("messages", messagesEntity)
					}
				} else if typeName == "CargoEntities.TextMessage" {
					if len(uuids[i]) > 0 {
						var messagesEntity *CargoEntities_TextMessageEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							messagesEntity = instance.(*CargoEntities_TextMessageEntity)
						} else {
							messagesEntity = GetServer().GetEntityManager().NewCargoEntitiesTextMessageEntity(this.GetUuid(), uuids[i], nil)
							messagesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(messagesEntity)
						}
						messagesEntity.AppendReferenced("messages", this)
						this.AppendChild("messages", messagesEntity)
					}
				}
			}
		}

		/** userRef **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.User"
				id_ := refTypeName + "$$" + id
				this.object.M_userRef = id
				GetServer().GetEntityManager().appendReference("userRef", this.object.UUID, id_)
			}
		}

		/** rolesRef **/
		if results[0][10] != nil {
			idsStr := results[0][10].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Role"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_rolesRef = append(this.object.M_rolesRef, ids[i])
					GetServer().GetEntityManager().appendReference("rolesRef", this.object.UUID, id_)
				}
			}
		}

		/** associations of Account **/

		/** entitiesPtr **/
		if results[0][11] != nil {
			id := results[0][11].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][12].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][13].(string)
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
func (this *EntityManager) NewCargoEntitiesAccountEntityFromObject(object *CargoEntities.Account) *CargoEntities_AccountEntity {
	return this.NewCargoEntitiesAccountEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_AccountEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesAccountExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Account"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_AccountEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_AccountEntity) AppendReference(reference Entity) {

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
//              			Computer
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_ComputerEntity struct {
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
	object         *CargoEntities.Computer
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesComputerEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_ComputerEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesComputerExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Computer).TYPENAME = "CargoEntities.Computer"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Computer", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Computer).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Computer).UUID
			}
			return val.(*CargoEntities_ComputerEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_ComputerEntity)
	if object == nil {
		entity.object = new(CargoEntities.Computer)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Computer)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Computer"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_ComputerEntity) GetTypeName() string {
	return "CargoEntities.Computer"
}
func (this *CargoEntities_ComputerEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_ComputerEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_ComputerEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_ComputerEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_ComputerEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_ComputerEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_ComputerEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_ComputerEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_ComputerEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_ComputerEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_ComputerEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_ComputerEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_ComputerEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_ComputerEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_ComputerEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_ComputerEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_ComputerEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_ComputerEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_ComputerEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_ComputerEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_ComputerEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_ComputerEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_ComputerEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Computer"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_ComputerEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ComputerEntityPrototype() {

	var computerEntityProto EntityPrototype
	computerEntityProto.TypeName = "CargoEntities.Computer"
	computerEntityProto.SuperTypeNames = append(computerEntityProto.SuperTypeNames, "CargoEntities.Entity")
	computerEntityProto.Ids = append(computerEntityProto.Ids, "uuid")
	computerEntityProto.Fields = append(computerEntityProto.Fields, "uuid")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 0)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)
	computerEntityProto.Indexs = append(computerEntityProto.Indexs, "parentUuid")
	computerEntityProto.Fields = append(computerEntityProto.Fields, "parentUuid")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 1)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	computerEntityProto.Ids = append(computerEntityProto.Ids, "M_id")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 2)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_id")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.ID")

	/** members of Computer **/
	computerEntityProto.Indexs = append(computerEntityProto.Indexs, "M_name")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 3)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_name")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 4)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_ipv4")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 5)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_osType")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "enum:OsType_Unknown:OsType_Linux:OsType_Windows7:OsType_Windows8:OsType_Windows10:OsType_OSX:OsType_IOS")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 6)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_platformType")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "enum:PlatformType_Unknown:PlatformType_Tablet:PlatformType_Phone:PlatformType_Desktop:PlatformType_Laptop")

	/** associations of Computer **/
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 7)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_entitiesPtr")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	computerEntityProto.Fields = append(computerEntityProto.Fields, "childsUuid")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "[]xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 8)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)

	computerEntityProto.Fields = append(computerEntityProto.Fields, "referenced")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "[]EntityRef")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 9)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&computerEntityProto)

}

/** Create **/
func (this *CargoEntities_ComputerEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Computer"

	var query EntityQuery
	query.TypeName = "CargoEntities.Computer"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Computer **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_osType")
	query.Fields = append(query.Fields, "M_platformType")

	/** associations of Computer **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ComputerInfo []interface{}

	ComputerInfo = append(ComputerInfo, this.GetUuid())
	if this.parentPtr != nil {
		ComputerInfo = append(ComputerInfo, this.parentPtr.GetUuid())
	} else {
		ComputerInfo = append(ComputerInfo, "")
	}

	/** members of Entity **/
	ComputerInfo = append(ComputerInfo, this.object.M_id)

	/** members of Computer **/
	ComputerInfo = append(ComputerInfo, this.object.M_name)
	ComputerInfo = append(ComputerInfo, this.object.M_ipv4)

	/** Save osType type OsType **/
	if this.object.M_osType == CargoEntities.OsType_Unknown {
		ComputerInfo = append(ComputerInfo, 0)
	} else if this.object.M_osType == CargoEntities.OsType_Linux {
		ComputerInfo = append(ComputerInfo, 1)
	} else if this.object.M_osType == CargoEntities.OsType_Windows7 {
		ComputerInfo = append(ComputerInfo, 2)
	} else if this.object.M_osType == CargoEntities.OsType_Windows8 {
		ComputerInfo = append(ComputerInfo, 3)
	} else if this.object.M_osType == CargoEntities.OsType_Windows10 {
		ComputerInfo = append(ComputerInfo, 4)
	} else if this.object.M_osType == CargoEntities.OsType_OSX {
		ComputerInfo = append(ComputerInfo, 5)
	} else if this.object.M_osType == CargoEntities.OsType_IOS {
		ComputerInfo = append(ComputerInfo, 6)
	} else {
		ComputerInfo = append(ComputerInfo, 0)
	}

	/** Save platformType type PlatformType **/
	if this.object.M_platformType == CargoEntities.PlatformType_Unknown {
		ComputerInfo = append(ComputerInfo, 0)
	} else if this.object.M_platformType == CargoEntities.PlatformType_Tablet {
		ComputerInfo = append(ComputerInfo, 1)
	} else if this.object.M_platformType == CargoEntities.PlatformType_Phone {
		ComputerInfo = append(ComputerInfo, 2)
	} else if this.object.M_platformType == CargoEntities.PlatformType_Desktop {
		ComputerInfo = append(ComputerInfo, 3)
	} else if this.object.M_platformType == CargoEntities.PlatformType_Laptop {
		ComputerInfo = append(ComputerInfo, 4)
	} else {
		ComputerInfo = append(ComputerInfo, 0)
	}

	/** associations of Computer **/

	/** Save entities type Entities **/
	ComputerInfo = append(ComputerInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ComputerInfo = append(ComputerInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ComputerInfo = append(ComputerInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), ComputerInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), ComputerInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_ComputerEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_ComputerEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Computer"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Computer **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_osType")
	query.Fields = append(query.Fields, "M_platformType")

	/** associations of Computer **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Computer...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Computer)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Computer"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of Computer **/

		/** name **/
		if results[0][3] != nil {
			this.object.M_name = results[0][3].(string)
		}

		/** ipv4 **/
		if results[0][4] != nil {
			this.object.M_ipv4 = results[0][4].(string)
		}

		/** osType **/
		if results[0][5] != nil {
			enumIndex := results[0][5].(int)
			if enumIndex == 0 {
				this.object.M_osType = CargoEntities.OsType_Unknown
			} else if enumIndex == 1 {
				this.object.M_osType = CargoEntities.OsType_Linux
			} else if enumIndex == 2 {
				this.object.M_osType = CargoEntities.OsType_Windows7
			} else if enumIndex == 3 {
				this.object.M_osType = CargoEntities.OsType_Windows8
			} else if enumIndex == 4 {
				this.object.M_osType = CargoEntities.OsType_Windows10
			} else if enumIndex == 5 {
				this.object.M_osType = CargoEntities.OsType_OSX
			} else if enumIndex == 6 {
				this.object.M_osType = CargoEntities.OsType_IOS
			}
		}

		/** platformType **/
		if results[0][6] != nil {
			enumIndex := results[0][6].(int)
			if enumIndex == 0 {
				this.object.M_platformType = CargoEntities.PlatformType_Unknown
			} else if enumIndex == 1 {
				this.object.M_platformType = CargoEntities.PlatformType_Tablet
			} else if enumIndex == 2 {
				this.object.M_platformType = CargoEntities.PlatformType_Phone
			} else if enumIndex == 3 {
				this.object.M_platformType = CargoEntities.PlatformType_Desktop
			} else if enumIndex == 4 {
				this.object.M_platformType = CargoEntities.PlatformType_Laptop
			}
		}

		/** associations of Computer **/

		/** entitiesPtr **/
		if results[0][7] != nil {
			id := results[0][7].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][8].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][9].(string)
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
func (this *EntityManager) NewCargoEntitiesComputerEntityFromObject(object *CargoEntities.Computer) *CargoEntities_ComputerEntity {
	return this.NewCargoEntitiesComputerEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_ComputerEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesComputerExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Computer"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_ComputerEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_ComputerEntity) AppendReference(reference Entity) {

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
//              			Permission
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_PermissionEntity struct {
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
	object         *CargoEntities.Permission
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesPermissionEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_PermissionEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesPermissionExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Permission).TYPENAME = "CargoEntities.Permission"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Permission", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Permission).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Permission).UUID
			}
			return val.(*CargoEntities_PermissionEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_PermissionEntity)
	if object == nil {
		entity.object = new(CargoEntities.Permission)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Permission)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Permission"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_PermissionEntity) GetTypeName() string {
	return "CargoEntities.Permission"
}
func (this *CargoEntities_PermissionEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_PermissionEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_PermissionEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_PermissionEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_PermissionEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_PermissionEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_PermissionEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_PermissionEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_PermissionEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_PermissionEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_PermissionEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_PermissionEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_PermissionEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_PermissionEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_PermissionEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_PermissionEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_PermissionEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_PermissionEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_PermissionEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_PermissionEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_PermissionEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_PermissionEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_PermissionEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Permission"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_PermissionEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_PermissionEntityPrototype() {

	var permissionEntityProto EntityPrototype
	permissionEntityProto.TypeName = "CargoEntities.Permission"
	permissionEntityProto.Ids = append(permissionEntityProto.Ids, "uuid")
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "uuid")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.string")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 0)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)
	permissionEntityProto.Indexs = append(permissionEntityProto.Indexs, "parentUuid")
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "parentUuid")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.string")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 1)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)

	/** members of Permission **/
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 2)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, true)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_pattern")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.string")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 3)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, true)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_type")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "enum:PermissionType_Create:PermissionType_Read:PermissionType_Update:PermissionType_Delete")

	/** associations of Permission **/
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 4)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_accountPtr")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "CargoEntities.Account:Ref")
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "childsUuid")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "[]xs.string")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 5)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)

	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "referenced")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "[]EntityRef")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 6)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&permissionEntityProto)

}

/** Create **/
func (this *CargoEntities_PermissionEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Permission"

	var query EntityQuery
	query.TypeName = "CargoEntities.Permission"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Permission **/
	query.Fields = append(query.Fields, "M_pattern")
	query.Fields = append(query.Fields, "M_type")

	/** associations of Permission **/
	query.Fields = append(query.Fields, "M_accountPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var PermissionInfo []interface{}

	PermissionInfo = append(PermissionInfo, this.GetUuid())
	if this.parentPtr != nil {
		PermissionInfo = append(PermissionInfo, this.parentPtr.GetUuid())
	} else {
		PermissionInfo = append(PermissionInfo, "")
	}

	/** members of Permission **/
	PermissionInfo = append(PermissionInfo, this.object.M_pattern)

	/** Save type type PermissionType **/
	if this.object.M_type == CargoEntities.PermissionType_Create {
		PermissionInfo = append(PermissionInfo, 0)
	} else if this.object.M_type == CargoEntities.PermissionType_Read {
		PermissionInfo = append(PermissionInfo, 1)
	} else if this.object.M_type == CargoEntities.PermissionType_Update {
		PermissionInfo = append(PermissionInfo, 2)
	} else if this.object.M_type == CargoEntities.PermissionType_Delete {
		PermissionInfo = append(PermissionInfo, 3)
	} else {
		PermissionInfo = append(PermissionInfo, 0)
	}

	/** associations of Permission **/

	/** Save account type Account **/
	PermissionInfo = append(PermissionInfo, this.object.M_accountPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	PermissionInfo = append(PermissionInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	PermissionInfo = append(PermissionInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), PermissionInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), PermissionInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_PermissionEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_PermissionEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Permission"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Permission **/
	query.Fields = append(query.Fields, "M_pattern")
	query.Fields = append(query.Fields, "M_type")

	/** associations of Permission **/
	query.Fields = append(query.Fields, "M_accountPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Permission...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Permission)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Permission"

		this.parentUuid = results[0][1].(string)

		/** members of Permission **/

		/** pattern **/
		if results[0][2] != nil {
			this.object.M_pattern = results[0][2].(string)
		}

		/** type **/
		if results[0][3] != nil {
			enumIndex := results[0][3].(int)
			if enumIndex == 0 {
				this.object.M_type = CargoEntities.PermissionType_Create
			} else if enumIndex == 1 {
				this.object.M_type = CargoEntities.PermissionType_Read
			} else if enumIndex == 2 {
				this.object.M_type = CargoEntities.PermissionType_Update
			} else if enumIndex == 3 {
				this.object.M_type = CargoEntities.PermissionType_Delete
			}
		}

		/** associations of Permission **/

		/** accountPtr **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_accountPtr = id
				GetServer().GetEntityManager().appendReference("accountPtr", this.object.UUID, id_)
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
func (this *EntityManager) NewCargoEntitiesPermissionEntityFromObject(object *CargoEntities.Permission) *CargoEntities_PermissionEntity {
	return this.NewCargoEntitiesPermissionEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_PermissionEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesPermissionExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Permission"
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_PermissionEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_PermissionEntity) AppendReference(reference Entity) {

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
//              			File
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_FileEntity struct {
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
	object         *CargoEntities.File
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesFileEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_FileEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesFileExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.File).TYPENAME = "CargoEntities.File"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.File", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.File).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.File).UUID
			}
			return val.(*CargoEntities_FileEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_FileEntity)
	if object == nil {
		entity.object = new(CargoEntities.File)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.File)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.File"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_FileEntity) GetTypeName() string {
	return "CargoEntities.File"
}
func (this *CargoEntities_FileEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_FileEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_FileEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_FileEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_FileEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_FileEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_FileEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_FileEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_FileEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_FileEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_FileEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_FileEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_FileEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_FileEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_FileEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_FileEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_FileEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_FileEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_FileEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_FileEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_FileEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_FileEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_FileEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.File"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_FileEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_FileEntityPrototype() {

	var fileEntityProto EntityPrototype
	fileEntityProto.TypeName = "CargoEntities.File"
	fileEntityProto.SuperTypeNames = append(fileEntityProto.SuperTypeNames, "CargoEntities.Entity")
	fileEntityProto.Ids = append(fileEntityProto.Ids, "uuid")
	fileEntityProto.Fields = append(fileEntityProto.Fields, "uuid")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 0)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.Indexs = append(fileEntityProto.Indexs, "parentUuid")
	fileEntityProto.Fields = append(fileEntityProto.Fields, "parentUuid")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 1)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	fileEntityProto.Ids = append(fileEntityProto.Ids, "M_id")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 2)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_id")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.ID")

	/** members of File **/
	fileEntityProto.Indexs = append(fileEntityProto.Indexs, "M_name")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 3)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_name")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 4)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_path")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 5)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_size")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.int")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 6)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_modeTime")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.long")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 7)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_isDir")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.boolean")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 8)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_checksum")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 9)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_data")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 10)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_thumbnail")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 11)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_mime")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 12)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_files")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "[]CargoEntities.File")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 13)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_fileType")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "enum:FileType_DbFile:FileType_DiskFile")

	/** associations of File **/
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 14)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_parentDirPtr")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "CargoEntities.File:Ref")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 15)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_entitiesPtr")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	fileEntityProto.Fields = append(fileEntityProto.Fields, "childsUuid")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "[]xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 16)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)

	fileEntityProto.Fields = append(fileEntityProto.Fields, "referenced")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "[]EntityRef")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 17)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&fileEntityProto)

}

/** Create **/
func (this *CargoEntities_FileEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.File"

	var query EntityQuery
	query.TypeName = "CargoEntities.File"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of File **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_path")
	query.Fields = append(query.Fields, "M_size")
	query.Fields = append(query.Fields, "M_modeTime")
	query.Fields = append(query.Fields, "M_isDir")
	query.Fields = append(query.Fields, "M_checksum")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_thumbnail")
	query.Fields = append(query.Fields, "M_mime")
	query.Fields = append(query.Fields, "M_files")
	query.Fields = append(query.Fields, "M_fileType")

	/** associations of File **/
	query.Fields = append(query.Fields, "M_parentDirPtr")
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var FileInfo []interface{}

	FileInfo = append(FileInfo, this.GetUuid())
	if this.parentPtr != nil {
		FileInfo = append(FileInfo, this.parentPtr.GetUuid())
	} else {
		FileInfo = append(FileInfo, "")
	}

	/** members of Entity **/
	FileInfo = append(FileInfo, this.object.M_id)

	/** members of File **/
	FileInfo = append(FileInfo, this.object.M_name)
	FileInfo = append(FileInfo, this.object.M_path)
	FileInfo = append(FileInfo, this.object.M_size)
	FileInfo = append(FileInfo, this.object.M_modeTime)
	FileInfo = append(FileInfo, this.object.M_isDir)
	FileInfo = append(FileInfo, this.object.M_checksum)
	FileInfo = append(FileInfo, this.object.M_data)
	FileInfo = append(FileInfo, this.object.M_thumbnail)
	FileInfo = append(FileInfo, this.object.M_mime)

	/** Save files type File **/
	filesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_files); i++ {
		filesEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity(this.GetUuid(), this.object.M_files[i].UUID, this.object.M_files[i])
		filesIds = append(filesIds, filesEntity.uuid)
		filesEntity.AppendReferenced("files", this)
		this.AppendChild("files", filesEntity)
		if filesEntity.NeedSave() {
			filesEntity.SaveEntity()
		}
	}
	filesStr, _ := json.Marshal(filesIds)
	FileInfo = append(FileInfo, string(filesStr))

	/** Save fileType type FileType **/
	if this.object.M_fileType == CargoEntities.FileType_DbFile {
		FileInfo = append(FileInfo, 0)
	} else if this.object.M_fileType == CargoEntities.FileType_DiskFile {
		FileInfo = append(FileInfo, 1)
	} else {
		FileInfo = append(FileInfo, 0)
	}

	/** associations of File **/

	/** Save parentDir type File **/
	FileInfo = append(FileInfo, this.object.M_parentDirPtr)

	/** Save entities type Entities **/
	FileInfo = append(FileInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	FileInfo = append(FileInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	FileInfo = append(FileInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), FileInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), FileInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_FileEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_FileEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.File"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of File **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_path")
	query.Fields = append(query.Fields, "M_size")
	query.Fields = append(query.Fields, "M_modeTime")
	query.Fields = append(query.Fields, "M_isDir")
	query.Fields = append(query.Fields, "M_checksum")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_thumbnail")
	query.Fields = append(query.Fields, "M_mime")
	query.Fields = append(query.Fields, "M_files")
	query.Fields = append(query.Fields, "M_fileType")

	/** associations of File **/
	query.Fields = append(query.Fields, "M_parentDirPtr")
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of File...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.File)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.File"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of File **/

		/** name **/
		if results[0][3] != nil {
			this.object.M_name = results[0][3].(string)
		}

		/** path **/
		if results[0][4] != nil {
			this.object.M_path = results[0][4].(string)
		}

		/** size **/
		if results[0][5] != nil {
			this.object.M_size = results[0][5].(int)
		}

		/** modeTime **/
		if results[0][6] != nil {
			this.object.M_modeTime = results[0][6].(int64)
		}

		/** isDir **/
		if results[0][7] != nil {
			this.object.M_isDir = results[0][7].(bool)
		}

		/** checksum **/
		if results[0][8] != nil {
			this.object.M_checksum = results[0][8].(string)
		}

		/** data **/
		if results[0][9] != nil {
			this.object.M_data = results[0][9].(string)
		}

		/** thumbnail **/
		if results[0][10] != nil {
			this.object.M_thumbnail = results[0][10].(string)
		}

		/** mime **/
		if results[0][11] != nil {
			this.object.M_mime = results[0][11].(string)
		}

		/** files **/
		if results[0][12] != nil {
			uuidsStr := results[0][12].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var filesEntity *CargoEntities_FileEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						filesEntity = instance.(*CargoEntities_FileEntity)
					} else {
						filesEntity = GetServer().GetEntityManager().NewCargoEntitiesFileEntity(this.GetUuid(), uuids[i], nil)
						filesEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(filesEntity)
					}
					filesEntity.AppendReferenced("files", this)
					this.AppendChild("files", filesEntity)
				}
			}
		}

		/** fileType **/
		if results[0][13] != nil {
			enumIndex := results[0][13].(int)
			if enumIndex == 0 {
				this.object.M_fileType = CargoEntities.FileType_DbFile
			} else if enumIndex == 1 {
				this.object.M_fileType = CargoEntities.FileType_DiskFile
			}
		}

		/** associations of File **/

		/** parentDirPtr **/
		if results[0][14] != nil {
			id := results[0][14].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.File"
				id_ := refTypeName + "$$" + id
				this.object.M_parentDirPtr = id
				GetServer().GetEntityManager().appendReference("parentDirPtr", this.object.UUID, id_)
			}
		}

		/** entitiesPtr **/
		if results[0][15] != nil {
			id := results[0][15].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][16].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][17].(string)
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
func (this *EntityManager) NewCargoEntitiesFileEntityFromObject(object *CargoEntities.File) *CargoEntities_FileEntity {
	return this.NewCargoEntitiesFileEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_FileEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesFileExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.File"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_FileEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_FileEntity) AppendReference(reference Entity) {

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
//              			User
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_UserEntity struct {
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
	object         *CargoEntities.User
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesUserEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_UserEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesUserExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.User).TYPENAME = "CargoEntities.User"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.User", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.User).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.User).UUID
			}
			return val.(*CargoEntities_UserEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_UserEntity)
	if object == nil {
		entity.object = new(CargoEntities.User)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.User)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.User"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_UserEntity) GetTypeName() string {
	return "CargoEntities.User"
}
func (this *CargoEntities_UserEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_UserEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_UserEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_UserEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_UserEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_UserEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_UserEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_UserEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_UserEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_UserEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_UserEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_UserEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_UserEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_UserEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_UserEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_UserEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_UserEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_UserEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_UserEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_UserEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_UserEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_UserEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_UserEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.User"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_UserEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_UserEntityPrototype() {

	var userEntityProto EntityPrototype
	userEntityProto.TypeName = "CargoEntities.User"
	userEntityProto.SuperTypeNames = append(userEntityProto.SuperTypeNames, "CargoEntities.Entity")
	userEntityProto.Ids = append(userEntityProto.Ids, "uuid")
	userEntityProto.Fields = append(userEntityProto.Fields, "uuid")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 0)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)
	userEntityProto.Indexs = append(userEntityProto.Indexs, "parentUuid")
	userEntityProto.Fields = append(userEntityProto.Fields, "parentUuid")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 1)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	userEntityProto.Ids = append(userEntityProto.Ids, "M_id")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 2)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_id")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.ID")

	/** members of User **/
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 3)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_firstName")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 4)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_lastName")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 5)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_middle")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 6)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_phone")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 7)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_email")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 8)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_memberOfRef")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "[]CargoEntities.Group:Ref")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 9)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_accounts")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "[]CargoEntities.Account:Ref")

	/** associations of User **/
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 10)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_entitiesPtr")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	userEntityProto.Fields = append(userEntityProto.Fields, "childsUuid")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "[]xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 11)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)

	userEntityProto.Fields = append(userEntityProto.Fields, "referenced")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "[]EntityRef")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 12)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&userEntityProto)

}

/** Create **/
func (this *CargoEntities_UserEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.User"

	var query EntityQuery
	query.TypeName = "CargoEntities.User"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of User **/
	query.Fields = append(query.Fields, "M_firstName")
	query.Fields = append(query.Fields, "M_lastName")
	query.Fields = append(query.Fields, "M_middle")
	query.Fields = append(query.Fields, "M_phone")
	query.Fields = append(query.Fields, "M_email")
	query.Fields = append(query.Fields, "M_memberOfRef")
	query.Fields = append(query.Fields, "M_accounts")

	/** associations of User **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var UserInfo []interface{}

	UserInfo = append(UserInfo, this.GetUuid())
	if this.parentPtr != nil {
		UserInfo = append(UserInfo, this.parentPtr.GetUuid())
	} else {
		UserInfo = append(UserInfo, "")
	}

	/** members of Entity **/
	UserInfo = append(UserInfo, this.object.M_id)

	/** members of User **/
	UserInfo = append(UserInfo, this.object.M_firstName)
	UserInfo = append(UserInfo, this.object.M_lastName)
	UserInfo = append(UserInfo, this.object.M_middle)
	UserInfo = append(UserInfo, this.object.M_phone)
	UserInfo = append(UserInfo, this.object.M_email)

	/** Save memberOfRef type Group **/
	memberOfRefStr, _ := json.Marshal(this.object.M_memberOfRef)
	UserInfo = append(UserInfo, string(memberOfRefStr))

	/** Save accounts type Account **/
	accountsStr, _ := json.Marshal(this.object.M_accounts)
	UserInfo = append(UserInfo, string(accountsStr))

	/** associations of User **/

	/** Save entities type Entities **/
	UserInfo = append(UserInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	UserInfo = append(UserInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	UserInfo = append(UserInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), UserInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), UserInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_UserEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_UserEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.User"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of User **/
	query.Fields = append(query.Fields, "M_firstName")
	query.Fields = append(query.Fields, "M_lastName")
	query.Fields = append(query.Fields, "M_middle")
	query.Fields = append(query.Fields, "M_phone")
	query.Fields = append(query.Fields, "M_email")
	query.Fields = append(query.Fields, "M_memberOfRef")
	query.Fields = append(query.Fields, "M_accounts")

	/** associations of User **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of User...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.User)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.User"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of User **/

		/** firstName **/
		if results[0][3] != nil {
			this.object.M_firstName = results[0][3].(string)
		}

		/** lastName **/
		if results[0][4] != nil {
			this.object.M_lastName = results[0][4].(string)
		}

		/** middle **/
		if results[0][5] != nil {
			this.object.M_middle = results[0][5].(string)
		}

		/** phone **/
		if results[0][6] != nil {
			this.object.M_phone = results[0][6].(string)
		}

		/** email **/
		if results[0][7] != nil {
			this.object.M_email = results[0][7].(string)
		}

		/** memberOfRef **/
		if results[0][8] != nil {
			idsStr := results[0][8].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Group"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_memberOfRef = append(this.object.M_memberOfRef, ids[i])
					GetServer().GetEntityManager().appendReference("memberOfRef", this.object.UUID, id_)
				}
			}
		}

		/** accounts **/
		if results[0][9] != nil {
			idsStr := results[0][9].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Account"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_accounts = append(this.object.M_accounts, ids[i])
					GetServer().GetEntityManager().appendReference("accounts", this.object.UUID, id_)
				}
			}
		}

		/** associations of User **/

		/** entitiesPtr **/
		if results[0][10] != nil {
			id := results[0][10].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
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
func (this *EntityManager) NewCargoEntitiesUserEntityFromObject(object *CargoEntities.User) *CargoEntities_UserEntity {
	return this.NewCargoEntitiesUserEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_UserEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesUserExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.User"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_UserEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_UserEntity) AppendReference(reference Entity) {

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
//              			Group
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_GroupEntity struct {
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
	object         *CargoEntities.Group
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesGroupEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_GroupEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesGroupExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Group).TYPENAME = "CargoEntities.Group"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Group", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Group).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Group).UUID
			}
			return val.(*CargoEntities_GroupEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_GroupEntity)
	if object == nil {
		entity.object = new(CargoEntities.Group)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Group)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Group"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_GroupEntity) GetTypeName() string {
	return "CargoEntities.Group"
}
func (this *CargoEntities_GroupEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_GroupEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_GroupEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_GroupEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_GroupEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_GroupEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_GroupEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_GroupEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_GroupEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_GroupEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_GroupEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_GroupEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_GroupEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_GroupEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_GroupEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_GroupEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_GroupEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_GroupEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_GroupEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_GroupEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_GroupEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_GroupEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_GroupEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Group"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_GroupEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_GroupEntityPrototype() {

	var groupEntityProto EntityPrototype
	groupEntityProto.TypeName = "CargoEntities.Group"
	groupEntityProto.SuperTypeNames = append(groupEntityProto.SuperTypeNames, "CargoEntities.Entity")
	groupEntityProto.Ids = append(groupEntityProto.Ids, "uuid")
	groupEntityProto.Fields = append(groupEntityProto.Fields, "uuid")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 0)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)
	groupEntityProto.Indexs = append(groupEntityProto.Indexs, "parentUuid")
	groupEntityProto.Fields = append(groupEntityProto.Fields, "parentUuid")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 1)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)

	/** members of Entity **/
	groupEntityProto.Ids = append(groupEntityProto.Ids, "M_id")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 2)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, true)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_id")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.ID")

	/** members of Group **/
	groupEntityProto.Indexs = append(groupEntityProto.Indexs, "M_name")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 3)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, true)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_name")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 4)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, true)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_membersRef")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "[]CargoEntities.User:Ref")

	/** associations of Group **/
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 5)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_entitiesPtr")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "CargoEntities.Entities:Ref")
	groupEntityProto.Fields = append(groupEntityProto.Fields, "childsUuid")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "[]xs.string")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 6)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)

	groupEntityProto.Fields = append(groupEntityProto.Fields, "referenced")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "[]EntityRef")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 7)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&groupEntityProto)

}

/** Create **/
func (this *CargoEntities_GroupEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Group"

	var query EntityQuery
	query.TypeName = "CargoEntities.Group"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Group **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_membersRef")

	/** associations of Group **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var GroupInfo []interface{}

	GroupInfo = append(GroupInfo, this.GetUuid())
	if this.parentPtr != nil {
		GroupInfo = append(GroupInfo, this.parentPtr.GetUuid())
	} else {
		GroupInfo = append(GroupInfo, "")
	}

	/** members of Entity **/
	GroupInfo = append(GroupInfo, this.object.M_id)

	/** members of Group **/
	GroupInfo = append(GroupInfo, this.object.M_name)

	/** Save membersRef type User **/
	membersRefStr, _ := json.Marshal(this.object.M_membersRef)
	GroupInfo = append(GroupInfo, string(membersRefStr))

	/** associations of Group **/

	/** Save entities type Entities **/
	GroupInfo = append(GroupInfo, this.object.M_entitiesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	GroupInfo = append(GroupInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	GroupInfo = append(GroupInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), GroupInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), GroupInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_GroupEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_GroupEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Group"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Group **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_membersRef")

	/** associations of Group **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Group...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Group)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Group"

		this.parentUuid = results[0][1].(string)

		/** members of Entity **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** members of Group **/

		/** name **/
		if results[0][3] != nil {
			this.object.M_name = results[0][3].(string)
		}

		/** membersRef **/
		if results[0][4] != nil {
			idsStr := results[0][4].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.User"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_membersRef = append(this.object.M_membersRef, ids[i])
					GetServer().GetEntityManager().appendReference("membersRef", this.object.UUID, id_)
				}
			}
		}

		/** associations of Group **/

		/** entitiesPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][6].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][7].(string)
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
func (this *EntityManager) NewCargoEntitiesGroupEntityFromObject(object *CargoEntities.Group) *CargoEntities_GroupEntity {
	return this.NewCargoEntitiesGroupEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_GroupEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesGroupExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Group"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_GroupEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_GroupEntity) AppendReference(reference Entity) {

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
//              			Entities
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_EntitiesEntity struct {
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
	object         *CargoEntities.Entities
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesEntitiesEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_EntitiesEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesEntitiesExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Entities).TYPENAME = "CargoEntities.Entities"
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Entities", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Entities).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Entities).UUID
			}
			return val.(*CargoEntities_EntitiesEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = "Config.Configurations%" + Utility.RandomUUID()
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
	entity := new(CargoEntities_EntitiesEntity)
	if object == nil {
		entity.object = new(CargoEntities.Entities)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Entities)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "CargoEntities.Entities"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *CargoEntities_EntitiesEntity) GetTypeName() string {
	return "CargoEntities.Entities"
}
func (this *CargoEntities_EntitiesEntity) GetUuid() string {
	return this.uuid
}
func (this *CargoEntities_EntitiesEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *CargoEntities_EntitiesEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *CargoEntities_EntitiesEntity) AppendReferenced(name string, owner Entity) {
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

func (this *CargoEntities_EntitiesEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_EntitiesEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *CargoEntities_EntitiesEntity) RemoveReference(name string, reference Entity) {
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

func (this *CargoEntities_EntitiesEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *CargoEntities_EntitiesEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *CargoEntities_EntitiesEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_EntitiesEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *CargoEntities_EntitiesEntity) RemoveChild(name string, uuid string) {
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

func (this *CargoEntities_EntitiesEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_EntitiesEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_EntitiesEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *CargoEntities_EntitiesEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *CargoEntities_EntitiesEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_EntitiesEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_EntitiesEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_EntitiesEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_EntitiesEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_EntitiesEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_EntitiesEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Entities"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_EntitiesEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_EntitiesEntityPrototype() {

	var entitiesEntityProto EntityPrototype
	entitiesEntityProto.TypeName = "CargoEntities.Entities"
	entitiesEntityProto.Ids = append(entitiesEntityProto.Ids, "uuid")
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "uuid")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 0)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, false)
	entitiesEntityProto.Indexs = append(entitiesEntityProto.Indexs, "parentUuid")
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "parentUuid")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 1)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, false)

	/** members of Entities **/
	entitiesEntityProto.Ids = append(entitiesEntityProto.Ids, "M_id")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 2)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_id")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.ID")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 3)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_name")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 4)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_version")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 5)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_entities")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Entity")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 6)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_roles")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Role")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 7)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_actions")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Action")
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "childsUuid")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 8)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, false)

	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "referenced")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]EntityRef")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 9)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&entitiesEntityProto)

}

/** Create **/
func (this *CargoEntities_EntitiesEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "CargoEntities.Entities"

	var query EntityQuery
	query.TypeName = "CargoEntities.Entities"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entities **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_version")
	query.Fields = append(query.Fields, "M_entities")
	query.Fields = append(query.Fields, "M_roles")
	query.Fields = append(query.Fields, "M_actions")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var EntitiesInfo []interface{}

	EntitiesInfo = append(EntitiesInfo, this.GetUuid())
	if this.parentPtr != nil {
		EntitiesInfo = append(EntitiesInfo, this.parentPtr.GetUuid())
	} else {
		EntitiesInfo = append(EntitiesInfo, "")
	}

	/** members of Entities **/
	EntitiesInfo = append(EntitiesInfo, this.object.M_id)
	EntitiesInfo = append(EntitiesInfo, this.object.M_name)
	EntitiesInfo = append(EntitiesInfo, this.object.M_version)

	/** Save entities type Entity **/
	entitiesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_entities); i++ {
		switch v := this.object.M_entities[i].(type) {
		case *CargoEntities.Error:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesErrorEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		case *CargoEntities.LogEntry:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesLogEntryEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		case *CargoEntities.Notification:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesNotificationEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		case *CargoEntities.TextMessage:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesTextMessageEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		case *CargoEntities.File:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		case *CargoEntities.Group:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesGroupEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		case *CargoEntities.Log:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesLogEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		case *CargoEntities.Project:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesProjectEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		case *CargoEntities.Account:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesAccountEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		case *CargoEntities.Computer:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesComputerEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		case *CargoEntities.User:
			entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesUserEntity(this.GetUuid(), v.UUID, v)
			entitiesIds = append(entitiesIds, entitiesEntity.uuid)
			entitiesEntity.AppendReferenced("entities", this)
			this.AppendChild("entities", entitiesEntity)
			if entitiesEntity.NeedSave() {
				entitiesEntity.SaveEntity()
			}
		}
	}
	entitiesStr, _ := json.Marshal(entitiesIds)
	EntitiesInfo = append(EntitiesInfo, string(entitiesStr))

	/** Save roles type Role **/
	rolesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_roles); i++ {
		rolesEntity := GetServer().GetEntityManager().NewCargoEntitiesRoleEntity(this.GetUuid(), this.object.M_roles[i].UUID, this.object.M_roles[i])
		rolesIds = append(rolesIds, rolesEntity.uuid)
		rolesEntity.AppendReferenced("roles", this)
		this.AppendChild("roles", rolesEntity)
		if rolesEntity.NeedSave() {
			rolesEntity.SaveEntity()
		}
	}
	rolesStr, _ := json.Marshal(rolesIds)
	EntitiesInfo = append(EntitiesInfo, string(rolesStr))

	/** Save actions type Action **/
	actionsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_actions); i++ {
		actionsEntity := GetServer().GetEntityManager().NewCargoEntitiesActionEntity(this.GetUuid(), this.object.M_actions[i].UUID, this.object.M_actions[i])
		actionsIds = append(actionsIds, actionsEntity.uuid)
		actionsEntity.AppendReferenced("actions", this)
		this.AppendChild("actions", actionsEntity)
		if actionsEntity.NeedSave() {
			actionsEntity.SaveEntity()
		}
	}
	actionsStr, _ := json.Marshal(actionsIds)
	EntitiesInfo = append(EntitiesInfo, string(actionsStr))
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	EntitiesInfo = append(EntitiesInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	EntitiesInfo = append(EntitiesInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), EntitiesInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), EntitiesInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_EntitiesEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_EntitiesEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Entities"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Entities **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_version")
	query.Fields = append(query.Fields, "M_entities")
	query.Fields = append(query.Fields, "M_roles")
	query.Fields = append(query.Fields, "M_actions")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Entities...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Entities)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "CargoEntities.Entities"

		this.parentUuid = results[0][1].(string)

		/** members of Entities **/

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

		/** entities **/
		if results[0][5] != nil {
			uuidsStr := results[0][5].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				typeName := uuids[i][0:strings.Index(uuids[i], "%")]
				if err != nil {
					log.Println("type ", typeName, " not found!")
					return err
				}
				if typeName == "CargoEntities.Error" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_ErrorEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_ErrorEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesErrorEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				} else if typeName == "CargoEntities.LogEntry" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_LogEntryEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_LogEntryEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesLogEntryEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				} else if typeName == "CargoEntities.Notification" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_NotificationEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_NotificationEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesNotificationEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				} else if typeName == "CargoEntities.TextMessage" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_TextMessageEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_TextMessageEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesTextMessageEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				} else if typeName == "CargoEntities.File" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_FileEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_FileEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesFileEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				} else if typeName == "CargoEntities.Group" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_GroupEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_GroupEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesGroupEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				} else if typeName == "CargoEntities.User" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_UserEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_UserEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesUserEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				} else if typeName == "CargoEntities.Log" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_LogEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_LogEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesLogEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				} else if typeName == "CargoEntities.Project" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_ProjectEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_ProjectEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesProjectEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				} else if typeName == "CargoEntities.Account" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_AccountEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_AccountEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesAccountEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				} else if typeName == "CargoEntities.Computer" {
					if len(uuids[i]) > 0 {
						var entitiesEntity *CargoEntities_ComputerEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entitiesEntity = instance.(*CargoEntities_ComputerEntity)
						} else {
							entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesComputerEntity(this.GetUuid(), uuids[i], nil)
							entitiesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(entitiesEntity)
						}
						entitiesEntity.AppendReferenced("entities", this)
						this.AppendChild("entities", entitiesEntity)
					}
				}
			}
		}

		/** roles **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var rolesEntity *CargoEntities_RoleEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						rolesEntity = instance.(*CargoEntities_RoleEntity)
					} else {
						rolesEntity = GetServer().GetEntityManager().NewCargoEntitiesRoleEntity(this.GetUuid(), uuids[i], nil)
						rolesEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(rolesEntity)
					}
					rolesEntity.AppendReferenced("roles", this)
					this.AppendChild("roles", rolesEntity)
				}
			}
		}

		/** actions **/
		if results[0][7] != nil {
			uuidsStr := results[0][7].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var actionsEntity *CargoEntities_ActionEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						actionsEntity = instance.(*CargoEntities_ActionEntity)
					} else {
						actionsEntity = GetServer().GetEntityManager().NewCargoEntitiesActionEntity(this.GetUuid(), uuids[i], nil)
						actionsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(actionsEntity)
					}
					actionsEntity.AppendReferenced("actions", this)
					this.AppendChild("actions", actionsEntity)
				}
			}
		}
		childsUuidStr := results[0][8].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][9].(string)
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
func (this *EntityManager) NewCargoEntitiesEntitiesEntityFromObject(object *CargoEntities.Entities) *CargoEntities_EntitiesEntity {
	return this.NewCargoEntitiesEntitiesEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_EntitiesEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesEntitiesExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Entities"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_EntitiesEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *CargoEntities_EntitiesEntity) AppendReference(reference Entity) {

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
func (this *EntityManager) registerCargoEntitiesObjects() {
	Utility.RegisterType((*CargoEntities.Parameter)(nil))
	Utility.RegisterType((*CargoEntities.Action)(nil))
	Utility.RegisterType((*CargoEntities.Error)(nil))
	Utility.RegisterType((*CargoEntities.LogEntry)(nil))
	Utility.RegisterType((*CargoEntities.Log)(nil))
	Utility.RegisterType((*CargoEntities.Project)(nil))
	Utility.RegisterType((*CargoEntities.Notification)(nil))
	Utility.RegisterType((*CargoEntities.TextMessage)(nil))
	Utility.RegisterType((*CargoEntities.Session)(nil))
	Utility.RegisterType((*CargoEntities.Role)(nil))
	Utility.RegisterType((*CargoEntities.Account)(nil))
	Utility.RegisterType((*CargoEntities.Computer)(nil))
	Utility.RegisterType((*CargoEntities.Permission)(nil))
	Utility.RegisterType((*CargoEntities.File)(nil))
	Utility.RegisterType((*CargoEntities.User)(nil))
	Utility.RegisterType((*CargoEntities.Group)(nil))
	Utility.RegisterType((*CargoEntities.Entities)(nil))
}

/** Create entity prototypes contain in a package **/
func (this *EntityManager) createCargoEntitiesPrototypes() {
	this.create_CargoEntities_EntityEntityPrototype()
	this.create_CargoEntities_ParameterEntityPrototype()
	this.create_CargoEntities_ActionEntityPrototype()
	this.create_CargoEntities_ErrorEntityPrototype()
	this.create_CargoEntities_LogEntryEntityPrototype()
	this.create_CargoEntities_LogEntityPrototype()
	this.create_CargoEntities_ProjectEntityPrototype()
	this.create_CargoEntities_MessageEntityPrototype()
	this.create_CargoEntities_NotificationEntityPrototype()
	this.create_CargoEntities_TextMessageEntityPrototype()
	this.create_CargoEntities_SessionEntityPrototype()
	this.create_CargoEntities_RoleEntityPrototype()
	this.create_CargoEntities_AccountEntityPrototype()
	this.create_CargoEntities_ComputerEntityPrototype()
	this.create_CargoEntities_PermissionEntityPrototype()
	this.create_CargoEntities_FileEntityPrototype()
	this.create_CargoEntities_UserEntityPrototype()
	this.create_CargoEntities_GroupEntityPrototype()
	this.create_CargoEntities_EntitiesEntityPrototype()
}
