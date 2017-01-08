package Server

import (
	"encoding/json"

	"code.google.com/p/go-uuid/uuid"
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/DC"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
	//	"log"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////
//              			Font
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type DC_FontEntity struct {
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
	object         *DC.Font
}

/** Constructor function **/
func (this *EntityManager) NewDCFontEntity(objectId string, object interface{}) *DC_FontEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = DCFontExists(objectId)
		}
	}
	if object != nil {
		object.(*DC.Font).TYPENAME = "DC.Font"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*DC.Font).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*DC_FontEntity)
		}
	} else {
		uuidStr = "DC.Font%" + uuid.NewRandom().String()
	}
	entity := new(DC_FontEntity)
	if object == nil {
		entity.object = new(DC.Font)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*DC.Font)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "DC.Font"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("DC.Font", "BPMN20")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *DC_FontEntity) GetTypeName() string {
	return "DC.Font"
}
func (this *DC_FontEntity) GetUuid() string {
	return this.uuid
}
func (this *DC_FontEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *DC_FontEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *DC_FontEntity) AppendReferenced(name string, owner Entity) {
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

func (this *DC_FontEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *DC_FontEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *DC_FontEntity) RemoveReference(name string, reference Entity) {
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

func (this *DC_FontEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *DC_FontEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *DC_FontEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *DC_FontEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *DC_FontEntity) RemoveChild(name string, uuid string) {
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

func (this *DC_FontEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *DC_FontEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *DC_FontEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *DC_FontEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *DC_FontEntity) GetObject() interface{} {
	return this.object
}

func (this *DC_FontEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *DC_FontEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *DC_FontEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *DC_FontEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *DC_FontEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *DC_FontEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "DC.Font"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(DCDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *DC_FontEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) Create_DC_FontEntityPrototype() {

	var fontEntityProto EntityPrototype
	fontEntityProto.TypeName = "DC.Font"
	fontEntityProto.Ids = append(fontEntityProto.Ids, "uuid")
	fontEntityProto.Fields = append(fontEntityProto.Fields, "uuid")
	fontEntityProto.FieldsType = append(fontEntityProto.FieldsType, "xs.string")
	fontEntityProto.FieldsOrder = append(fontEntityProto.FieldsOrder, 0)
	fontEntityProto.FieldsVisibility = append(fontEntityProto.FieldsVisibility, false)
	fontEntityProto.Indexs = append(fontEntityProto.Indexs, "parentUuid")
	fontEntityProto.Fields = append(fontEntityProto.Fields, "parentUuid")
	fontEntityProto.FieldsType = append(fontEntityProto.FieldsType, "xs.string")
	fontEntityProto.FieldsOrder = append(fontEntityProto.FieldsOrder, 1)
	fontEntityProto.FieldsVisibility = append(fontEntityProto.FieldsVisibility, false)

	/** members of Font **/
	fontEntityProto.FieldsOrder = append(fontEntityProto.FieldsOrder, 2)
	fontEntityProto.FieldsVisibility = append(fontEntityProto.FieldsVisibility, true)
	fontEntityProto.Fields = append(fontEntityProto.Fields, "M_name")
	fontEntityProto.FieldsType = append(fontEntityProto.FieldsType, "xs.string")
	fontEntityProto.FieldsOrder = append(fontEntityProto.FieldsOrder, 3)
	fontEntityProto.FieldsVisibility = append(fontEntityProto.FieldsVisibility, true)
	fontEntityProto.Fields = append(fontEntityProto.Fields, "M_size")
	fontEntityProto.FieldsType = append(fontEntityProto.FieldsType, "xs.float64")
	fontEntityProto.FieldsOrder = append(fontEntityProto.FieldsOrder, 4)
	fontEntityProto.FieldsVisibility = append(fontEntityProto.FieldsVisibility, true)
	fontEntityProto.Fields = append(fontEntityProto.Fields, "M_isBold")
	fontEntityProto.FieldsType = append(fontEntityProto.FieldsType, "xs.bool")
	fontEntityProto.FieldsOrder = append(fontEntityProto.FieldsOrder, 5)
	fontEntityProto.FieldsVisibility = append(fontEntityProto.FieldsVisibility, true)
	fontEntityProto.Fields = append(fontEntityProto.Fields, "M_isItalic")
	fontEntityProto.FieldsType = append(fontEntityProto.FieldsType, "xs.bool")
	fontEntityProto.FieldsOrder = append(fontEntityProto.FieldsOrder, 6)
	fontEntityProto.FieldsVisibility = append(fontEntityProto.FieldsVisibility, true)
	fontEntityProto.Fields = append(fontEntityProto.Fields, "M_isUnderline")
	fontEntityProto.FieldsType = append(fontEntityProto.FieldsType, "xs.bool")
	fontEntityProto.FieldsOrder = append(fontEntityProto.FieldsOrder, 7)
	fontEntityProto.FieldsVisibility = append(fontEntityProto.FieldsVisibility, true)
	fontEntityProto.Fields = append(fontEntityProto.Fields, "M_isStrikeThrough")
	fontEntityProto.FieldsType = append(fontEntityProto.FieldsType, "xs.bool")
	fontEntityProto.Fields = append(fontEntityProto.Fields, "childsUuid")
	fontEntityProto.FieldsType = append(fontEntityProto.FieldsType, "[]xs.string")
	fontEntityProto.FieldsOrder = append(fontEntityProto.FieldsOrder, 8)
	fontEntityProto.FieldsVisibility = append(fontEntityProto.FieldsVisibility, false)

	fontEntityProto.Fields = append(fontEntityProto.Fields, "referenced")
	fontEntityProto.FieldsType = append(fontEntityProto.FieldsType, "[]EntityRef")
	fontEntityProto.FieldsOrder = append(fontEntityProto.FieldsOrder, 9)
	fontEntityProto.FieldsVisibility = append(fontEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DCDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&fontEntityProto)

}

/** Create **/
func (this *DC_FontEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "DC.Font"

	var query EntityQuery
	query.TypeName = "DC.Font"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Font **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_size")
	query.Fields = append(query.Fields, "M_isBold")
	query.Fields = append(query.Fields, "M_isItalic")
	query.Fields = append(query.Fields, "M_isUnderline")
	query.Fields = append(query.Fields, "M_isStrikeThrough")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var FontInfo []interface{}

	FontInfo = append(FontInfo, this.GetUuid())
	if this.parentPtr != nil {
		FontInfo = append(FontInfo, this.parentPtr.GetUuid())
	} else {
		FontInfo = append(FontInfo, "")
	}

	/** members of Font **/
	FontInfo = append(FontInfo, this.object.M_name)
	FontInfo = append(FontInfo, this.object.M_size)
	FontInfo = append(FontInfo, this.object.M_isBold)
	FontInfo = append(FontInfo, this.object.M_isItalic)
	FontInfo = append(FontInfo, this.object.M_isUnderline)
	FontInfo = append(FontInfo, this.object.M_isStrikeThrough)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	FontInfo = append(FontInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	FontInfo = append(FontInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(DCDB, string(queryStr), FontInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(DCDB, string(queryStr), FontInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *DC_FontEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*DC_FontEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "DC.Font"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Font **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_size")
	query.Fields = append(query.Fields, "M_isBold")
	query.Fields = append(query.Fields, "M_isItalic")
	query.Fields = append(query.Fields, "M_isUnderline")
	query.Fields = append(query.Fields, "M_isStrikeThrough")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(DCDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Font...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(DC.Font)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "DC.Font"

		this.parentUuid = results[0][1].(string)

		/** members of Font **/

		/** name **/
		if results[0][2] != nil {
			this.object.M_name = results[0][2].(string)
		}

		/** size **/
		if results[0][3] != nil {
			this.object.M_size = results[0][3].(float64)
		}

		/** isBold **/
		if results[0][4] != nil {
			this.object.M_isBold = results[0][4].(bool)
		}

		/** isItalic **/
		if results[0][5] != nil {
			this.object.M_isItalic = results[0][5].(bool)
		}

		/** isUnderline **/
		if results[0][6] != nil {
			this.object.M_isUnderline = results[0][6].(bool)
		}

		/** isStrikeThrough **/
		if results[0][7] != nil {
			this.object.M_isStrikeThrough = results[0][7].(bool)
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
func (this *EntityManager) NewDCFontEntityFromObject(object *DC.Font) *DC_FontEntity {
	return this.NewDCFontEntity(object.UUID, object)
}

/** Delete **/
func (this *DC_FontEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func DCFontExists(val string) string {
	var query EntityQuery
	query.TypeName = "DC.Font"
	query.Indexs = append(query.Indexs, "M_name="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(DCDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *DC_FontEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *DC_FontEntity) AppendReference(reference Entity) {

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
//              			Point
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type DC_PointEntity struct {
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
	object         *DC.Point
}

/** Constructor function **/
func (this *EntityManager) NewDCPointEntity(objectId string, object interface{}) *DC_PointEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = DCPointExists(objectId)
		}
	}
	if object != nil {
		object.(*DC.Point).TYPENAME = "DC.Point"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*DC.Point).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*DC_PointEntity)
		}
	} else {
		uuidStr = "DC.Point%" + uuid.NewRandom().String()
	}
	entity := new(DC_PointEntity)
	if object == nil {
		entity.object = new(DC.Point)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*DC.Point)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "DC.Point"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("DC.Point", "BPMN20")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *DC_PointEntity) GetTypeName() string {
	return "DC.Point"
}
func (this *DC_PointEntity) GetUuid() string {
	return this.uuid
}
func (this *DC_PointEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *DC_PointEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *DC_PointEntity) AppendReferenced(name string, owner Entity) {
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

func (this *DC_PointEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *DC_PointEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *DC_PointEntity) RemoveReference(name string, reference Entity) {
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

func (this *DC_PointEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *DC_PointEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *DC_PointEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *DC_PointEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *DC_PointEntity) RemoveChild(name string, uuid string) {
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

func (this *DC_PointEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *DC_PointEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *DC_PointEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *DC_PointEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *DC_PointEntity) GetObject() interface{} {
	return this.object
}

func (this *DC_PointEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *DC_PointEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *DC_PointEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *DC_PointEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *DC_PointEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *DC_PointEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "DC.Point"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(DCDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *DC_PointEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) Create_DC_PointEntityPrototype() {

	var pointEntityProto EntityPrototype
	pointEntityProto.TypeName = "DC.Point"
	pointEntityProto.Ids = append(pointEntityProto.Ids, "uuid")
	pointEntityProto.Fields = append(pointEntityProto.Fields, "uuid")
	pointEntityProto.FieldsType = append(pointEntityProto.FieldsType, "xs.string")
	pointEntityProto.FieldsOrder = append(pointEntityProto.FieldsOrder, 0)
	pointEntityProto.FieldsVisibility = append(pointEntityProto.FieldsVisibility, false)
	pointEntityProto.Indexs = append(pointEntityProto.Indexs, "parentUuid")
	pointEntityProto.Fields = append(pointEntityProto.Fields, "parentUuid")
	pointEntityProto.FieldsType = append(pointEntityProto.FieldsType, "xs.string")
	pointEntityProto.FieldsOrder = append(pointEntityProto.FieldsOrder, 1)
	pointEntityProto.FieldsVisibility = append(pointEntityProto.FieldsVisibility, false)

	/** members of Point **/
	pointEntityProto.FieldsOrder = append(pointEntityProto.FieldsOrder, 2)
	pointEntityProto.FieldsVisibility = append(pointEntityProto.FieldsVisibility, true)
	pointEntityProto.Fields = append(pointEntityProto.Fields, "M_id")
	pointEntityProto.FieldsType = append(pointEntityProto.FieldsType, "xs.string")
	pointEntityProto.FieldsOrder = append(pointEntityProto.FieldsOrder, 3)
	pointEntityProto.FieldsVisibility = append(pointEntityProto.FieldsVisibility, true)
	pointEntityProto.Fields = append(pointEntityProto.Fields, "M_x")
	pointEntityProto.FieldsType = append(pointEntityProto.FieldsType, "xs.float64")
	pointEntityProto.FieldsOrder = append(pointEntityProto.FieldsOrder, 4)
	pointEntityProto.FieldsVisibility = append(pointEntityProto.FieldsVisibility, true)
	pointEntityProto.Fields = append(pointEntityProto.Fields, "M_y")
	pointEntityProto.FieldsType = append(pointEntityProto.FieldsType, "xs.float64")
	pointEntityProto.Fields = append(pointEntityProto.Fields, "childsUuid")
	pointEntityProto.FieldsType = append(pointEntityProto.FieldsType, "[]xs.string")
	pointEntityProto.FieldsOrder = append(pointEntityProto.FieldsOrder, 5)
	pointEntityProto.FieldsVisibility = append(pointEntityProto.FieldsVisibility, false)

	pointEntityProto.Fields = append(pointEntityProto.Fields, "referenced")
	pointEntityProto.FieldsType = append(pointEntityProto.FieldsType, "[]EntityRef")
	pointEntityProto.FieldsOrder = append(pointEntityProto.FieldsOrder, 6)
	pointEntityProto.FieldsVisibility = append(pointEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DCDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&pointEntityProto)

}

/** Create **/
func (this *DC_PointEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "DC.Point"

	var query EntityQuery
	query.TypeName = "DC.Point"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Point **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_x")
	query.Fields = append(query.Fields, "M_y")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var PointInfo []interface{}

	PointInfo = append(PointInfo, this.GetUuid())
	if this.parentPtr != nil {
		PointInfo = append(PointInfo, this.parentPtr.GetUuid())
	} else {
		PointInfo = append(PointInfo, "")
	}

	/** members of Point **/
	PointInfo = append(PointInfo, this.object.M_id)
	PointInfo = append(PointInfo, this.object.M_x)
	PointInfo = append(PointInfo, this.object.M_y)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	PointInfo = append(PointInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	PointInfo = append(PointInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(DCDB, string(queryStr), PointInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(DCDB, string(queryStr), PointInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *DC_PointEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*DC_PointEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "DC.Point"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Point **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_x")
	query.Fields = append(query.Fields, "M_y")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(DCDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Point...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(DC.Point)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "DC.Point"

		this.parentUuid = results[0][1].(string)

		/** members of Point **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** x **/
		if results[0][3] != nil {
			this.object.M_x = results[0][3].(float64)
		}

		/** y **/
		if results[0][4] != nil {
			this.object.M_y = results[0][4].(float64)
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
func (this *EntityManager) NewDCPointEntityFromObject(object *DC.Point) *DC_PointEntity {
	return this.NewDCPointEntity(object.UUID, object)
}

/** Delete **/
func (this *DC_PointEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func DCPointExists(val string) string {
	var query EntityQuery
	query.TypeName = "DC.Point"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(DCDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *DC_PointEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *DC_PointEntity) AppendReference(reference Entity) {

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
//              			Bounds
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type DC_BoundsEntity struct {
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
	object         *DC.Bounds
}

/** Constructor function **/
func (this *EntityManager) NewDCBoundsEntity(objectId string, object interface{}) *DC_BoundsEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = DCBoundsExists(objectId)
		}
	}
	if object != nil {
		object.(*DC.Bounds).TYPENAME = "DC.Bounds"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*DC.Bounds).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*DC_BoundsEntity)
		}
	} else {
		uuidStr = "DC.Bounds%" + uuid.NewRandom().String()
	}
	entity := new(DC_BoundsEntity)
	if object == nil {
		entity.object = new(DC.Bounds)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*DC.Bounds)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "DC.Bounds"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("DC.Bounds", "BPMN20")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *DC_BoundsEntity) GetTypeName() string {
	return "DC.Bounds"
}
func (this *DC_BoundsEntity) GetUuid() string {
	return this.uuid
}
func (this *DC_BoundsEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *DC_BoundsEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *DC_BoundsEntity) AppendReferenced(name string, owner Entity) {
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

func (this *DC_BoundsEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *DC_BoundsEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *DC_BoundsEntity) RemoveReference(name string, reference Entity) {
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

func (this *DC_BoundsEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *DC_BoundsEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *DC_BoundsEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *DC_BoundsEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *DC_BoundsEntity) RemoveChild(name string, uuid string) {
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

func (this *DC_BoundsEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *DC_BoundsEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *DC_BoundsEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *DC_BoundsEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *DC_BoundsEntity) GetObject() interface{} {
	return this.object
}

func (this *DC_BoundsEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *DC_BoundsEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *DC_BoundsEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *DC_BoundsEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *DC_BoundsEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *DC_BoundsEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "DC.Bounds"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(DCDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *DC_BoundsEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) Create_DC_BoundsEntityPrototype() {

	var boundsEntityProto EntityPrototype
	boundsEntityProto.TypeName = "DC.Bounds"
	boundsEntityProto.Ids = append(boundsEntityProto.Ids, "uuid")
	boundsEntityProto.Fields = append(boundsEntityProto.Fields, "uuid")
	boundsEntityProto.FieldsType = append(boundsEntityProto.FieldsType, "xs.string")
	boundsEntityProto.FieldsOrder = append(boundsEntityProto.FieldsOrder, 0)
	boundsEntityProto.FieldsVisibility = append(boundsEntityProto.FieldsVisibility, false)
	boundsEntityProto.Indexs = append(boundsEntityProto.Indexs, "parentUuid")
	boundsEntityProto.Fields = append(boundsEntityProto.Fields, "parentUuid")
	boundsEntityProto.FieldsType = append(boundsEntityProto.FieldsType, "xs.string")
	boundsEntityProto.FieldsOrder = append(boundsEntityProto.FieldsOrder, 1)
	boundsEntityProto.FieldsVisibility = append(boundsEntityProto.FieldsVisibility, false)

	/** members of Bounds **/
	boundsEntityProto.FieldsOrder = append(boundsEntityProto.FieldsOrder, 2)
	boundsEntityProto.FieldsVisibility = append(boundsEntityProto.FieldsVisibility, true)
	boundsEntityProto.Fields = append(boundsEntityProto.Fields, "M_id")
	boundsEntityProto.FieldsType = append(boundsEntityProto.FieldsType, "xs.string")
	boundsEntityProto.FieldsOrder = append(boundsEntityProto.FieldsOrder, 3)
	boundsEntityProto.FieldsVisibility = append(boundsEntityProto.FieldsVisibility, true)
	boundsEntityProto.Fields = append(boundsEntityProto.Fields, "M_x")
	boundsEntityProto.FieldsType = append(boundsEntityProto.FieldsType, "xs.float64")
	boundsEntityProto.FieldsOrder = append(boundsEntityProto.FieldsOrder, 4)
	boundsEntityProto.FieldsVisibility = append(boundsEntityProto.FieldsVisibility, true)
	boundsEntityProto.Fields = append(boundsEntityProto.Fields, "M_y")
	boundsEntityProto.FieldsType = append(boundsEntityProto.FieldsType, "xs.float64")
	boundsEntityProto.FieldsOrder = append(boundsEntityProto.FieldsOrder, 5)
	boundsEntityProto.FieldsVisibility = append(boundsEntityProto.FieldsVisibility, true)
	boundsEntityProto.Fields = append(boundsEntityProto.Fields, "M_width")
	boundsEntityProto.FieldsType = append(boundsEntityProto.FieldsType, "xs.float64")
	boundsEntityProto.FieldsOrder = append(boundsEntityProto.FieldsOrder, 6)
	boundsEntityProto.FieldsVisibility = append(boundsEntityProto.FieldsVisibility, true)
	boundsEntityProto.Fields = append(boundsEntityProto.Fields, "M_height")
	boundsEntityProto.FieldsType = append(boundsEntityProto.FieldsType, "xs.float64")
	boundsEntityProto.Fields = append(boundsEntityProto.Fields, "childsUuid")
	boundsEntityProto.FieldsType = append(boundsEntityProto.FieldsType, "[]xs.string")
	boundsEntityProto.FieldsOrder = append(boundsEntityProto.FieldsOrder, 7)
	boundsEntityProto.FieldsVisibility = append(boundsEntityProto.FieldsVisibility, false)

	boundsEntityProto.Fields = append(boundsEntityProto.Fields, "referenced")
	boundsEntityProto.FieldsType = append(boundsEntityProto.FieldsType, "[]EntityRef")
	boundsEntityProto.FieldsOrder = append(boundsEntityProto.FieldsOrder, 8)
	boundsEntityProto.FieldsVisibility = append(boundsEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DCDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&boundsEntityProto)

}

/** Create **/
func (this *DC_BoundsEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "DC.Bounds"

	var query EntityQuery
	query.TypeName = "DC.Bounds"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Bounds **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_x")
	query.Fields = append(query.Fields, "M_y")
	query.Fields = append(query.Fields, "M_width")
	query.Fields = append(query.Fields, "M_height")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var BoundsInfo []interface{}

	BoundsInfo = append(BoundsInfo, this.GetUuid())
	if this.parentPtr != nil {
		BoundsInfo = append(BoundsInfo, this.parentPtr.GetUuid())
	} else {
		BoundsInfo = append(BoundsInfo, "")
	}

	/** members of Bounds **/
	BoundsInfo = append(BoundsInfo, this.object.M_id)
	BoundsInfo = append(BoundsInfo, this.object.M_x)
	BoundsInfo = append(BoundsInfo, this.object.M_y)
	BoundsInfo = append(BoundsInfo, this.object.M_width)
	BoundsInfo = append(BoundsInfo, this.object.M_height)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	BoundsInfo = append(BoundsInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	BoundsInfo = append(BoundsInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(DCDB, string(queryStr), BoundsInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(DCDB, string(queryStr), BoundsInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *DC_BoundsEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*DC_BoundsEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "DC.Bounds"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Bounds **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_x")
	query.Fields = append(query.Fields, "M_y")
	query.Fields = append(query.Fields, "M_width")
	query.Fields = append(query.Fields, "M_height")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(DCDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Bounds...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(DC.Bounds)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "DC.Bounds"

		this.parentUuid = results[0][1].(string)

		/** members of Bounds **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** x **/
		if results[0][3] != nil {
			this.object.M_x = results[0][3].(float64)
		}

		/** y **/
		if results[0][4] != nil {
			this.object.M_y = results[0][4].(float64)
		}

		/** width **/
		if results[0][5] != nil {
			this.object.M_width = results[0][5].(float64)
		}

		/** height **/
		if results[0][6] != nil {
			this.object.M_height = results[0][6].(float64)
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
func (this *EntityManager) NewDCBoundsEntityFromObject(object *DC.Bounds) *DC_BoundsEntity {
	return this.NewDCBoundsEntity(object.UUID, object)
}

/** Delete **/
func (this *DC_BoundsEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func DCBoundsExists(val string) string {
	var query EntityQuery
	query.TypeName = "DC.Bounds"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(DCDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *DC_BoundsEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *DC_BoundsEntity) AppendReference(reference Entity) {

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
func (this *EntityManager) RegisterDCObjects() {
	Utility.RegisterType((*DC.Font)(nil))
	Utility.RegisterType((*DC.Point)(nil))
	Utility.RegisterType((*DC.Bounds)(nil))
}

/** Create entity prototypes contain in a package **/
func (this *EntityManager) CreateDCPrototypes() {
	this.Create_DC_FontEntityPrototype()
	this.Create_DC_PointEntityPrototype()
	this.Create_DC_BoundsEntityPrototype()
}
