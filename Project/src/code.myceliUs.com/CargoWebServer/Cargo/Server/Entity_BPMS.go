// +build BPMS

package Server

import (
	"encoding/json"
	"log"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/BPMS"
	"code.myceliUs.com/Utility"
)

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_InstanceEntityPrototype() {

	var instanceEntityProto EntityPrototype
	instanceEntityProto.TypeName = "BPMS.Instance"
	instanceEntityProto.IsAbstract = true
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS.DefinitionsInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS.EventDefinitionInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS.ConnectingObject")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS.ActivityInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS.GatewayInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS.ProcessInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS.SubprocessInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS.EventInstance")
	instanceEntityProto.Ids = append(instanceEntityProto.Ids, "uuid")
	instanceEntityProto.Fields = append(instanceEntityProto.Fields, "uuid")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType, "xs.string")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder, 0)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility, false)
	instanceEntityProto.Indexs = append(instanceEntityProto.Indexs, "parentUuid")
	instanceEntityProto.Fields = append(instanceEntityProto.Fields, "parentUuid")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType, "xs.string")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder, 1)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility, false)

	/** members of Instance **/
	instanceEntityProto.Ids = append(instanceEntityProto.Ids, "M_id")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder, 2)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility, true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields, "M_id")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType, "xs.ID")
	instanceEntityProto.Indexs = append(instanceEntityProto.Indexs, "M_bpmnElementId")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder, 3)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility, true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields, "M_bpmnElementId")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType, "xs.string")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder, 4)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility, true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields, "M_participants")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType, "[]xs.string")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder, 5)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility, true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields, "M_dataRef")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder, 6)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility, true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields, "M_data")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder, 7)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility, true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields, "M_logInfoRef")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType, "[]BPMS.LogInfo:Ref")
	instanceEntityProto.Fields = append(instanceEntityProto.Fields, "childsUuid")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType, "[]xs.string")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder, 8)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility, false)

	instanceEntityProto.Fields = append(instanceEntityProto.Fields, "referenced")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType, "[]EntityRef")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder, 9)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&instanceEntityProto)

}

////////////////////////////////////////////////////////////////////////////////
//              			ConnectingObject
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_ConnectingObjectEntity struct {
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
	object         *BPMS.ConnectingObject
}

/** Constructor function **/
func (this *EntityManager) NewBPMSConnectingObjectEntity(objectId string, object interface{}) *BPMS_ConnectingObjectEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSConnectingObjectExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.ConnectingObject).TYPENAME = "BPMS.ConnectingObject"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.ConnectingObject).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_ConnectingObjectEntity)
		}
	} else {
		uuidStr = "BPMS.ConnectingObject%" + Utility.RandomUUID()
	}
	entity := new(BPMS_ConnectingObjectEntity)
	if object == nil {
		entity.object = new(BPMS.ConnectingObject)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.ConnectingObject)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.ConnectingObject"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.ConnectingObject", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_ConnectingObjectEntity) GetTypeName() string {
	return "BPMS.ConnectingObject"
}
func (this *BPMS_ConnectingObjectEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_ConnectingObjectEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_ConnectingObjectEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_ConnectingObjectEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_ConnectingObjectEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_ConnectingObjectEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_ConnectingObjectEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_ConnectingObjectEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_ConnectingObjectEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_ConnectingObjectEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_ConnectingObjectEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_ConnectingObjectEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_ConnectingObjectEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_ConnectingObjectEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_ConnectingObjectEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_ConnectingObjectEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_ConnectingObjectEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_ConnectingObjectEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_ConnectingObjectEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_ConnectingObjectEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_ConnectingObjectEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_ConnectingObjectEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_ConnectingObjectEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.ConnectingObject"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_ConnectingObjectEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_ConnectingObjectEntityPrototype() {

	var connectingObjectEntityProto EntityPrototype
	connectingObjectEntityProto.TypeName = "BPMS.ConnectingObject"
	connectingObjectEntityProto.SuperTypeNames = append(connectingObjectEntityProto.SuperTypeNames, "BPMS.Instance")
	connectingObjectEntityProto.Ids = append(connectingObjectEntityProto.Ids, "uuid")
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "uuid")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "xs.string")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 0)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, false)
	connectingObjectEntityProto.Indexs = append(connectingObjectEntityProto.Indexs, "parentUuid")
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "parentUuid")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "xs.string")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 1)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, false)

	/** members of Instance **/
	connectingObjectEntityProto.Ids = append(connectingObjectEntityProto.Ids, "M_id")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 2)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_id")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "xs.ID")
	connectingObjectEntityProto.Indexs = append(connectingObjectEntityProto.Indexs, "M_bpmnElementId")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 3)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_bpmnElementId")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "xs.string")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 4)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_participants")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "[]xs.string")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 5)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_dataRef")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 6)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_data")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 7)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_logInfoRef")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "[]BPMS.LogInfo:Ref")

	/** members of ConnectingObject **/
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 8)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_connectingObjectType")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "enum:ConnectingObjectType_SequenceFlow:ConnectingObjectType_MessageFlow:ConnectingObjectType_Association:ConnectingObjectType_DataAssociation")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 9)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_sourceRef")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "BPMS.FlowNodeInstance:Ref")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 10)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_targetRef")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "BPMS.FlowNodeInstance:Ref")

	/** associations of ConnectingObject **/
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 11)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, false)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_SubprocessInstancePtr")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "BPMS.SubprocessInstance:Ref")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 12)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, false)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "M_processInstancePtr")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "BPMS.ProcessInstance:Ref")
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "childsUuid")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "[]xs.string")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 13)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, false)

	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields, "referenced")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType, "[]EntityRef")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder, 14)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&connectingObjectEntityProto)

}

/** Create **/
func (this *BPMS_ConnectingObjectEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.ConnectingObject"

	var query EntityQuery
	query.TypeName = "BPMS.ConnectingObject"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of ConnectingObject **/
	query.Fields = append(query.Fields, "M_connectingObjectType")
	query.Fields = append(query.Fields, "M_sourceRef")
	query.Fields = append(query.Fields, "M_targetRef")

	/** associations of ConnectingObject **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ConnectingObjectInfo []interface{}

	ConnectingObjectInfo = append(ConnectingObjectInfo, this.GetUuid())
	if this.parentPtr != nil {
		ConnectingObjectInfo = append(ConnectingObjectInfo, this.parentPtr.GetUuid())
	} else {
		ConnectingObjectInfo = append(ConnectingObjectInfo, "")
	}

	/** members of Instance **/
	ConnectingObjectInfo = append(ConnectingObjectInfo, this.object.M_id)
	ConnectingObjectInfo = append(ConnectingObjectInfo, this.object.M_bpmnElementId)
	ConnectingObjectInfo = append(ConnectingObjectInfo, this.object.M_participants)

	/** Save dataRef type ItemAwareElementInstance **/
	dataRefStr, _ := json.Marshal(this.object.M_dataRef)
	ConnectingObjectInfo = append(ConnectingObjectInfo, string(dataRefStr))

	/** Save data type ItemAwareElementInstance **/
	dataIds := make([]string, 0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity := GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(this.object.M_data[i].UUID, this.object.M_data[i])
		dataIds = append(dataIds, dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data", dataEntity)
		if dataEntity.NeedSave() {
			dataEntity.SaveEntity()
		}
	}
	dataStr, _ := json.Marshal(dataIds)
	ConnectingObjectInfo = append(ConnectingObjectInfo, string(dataStr))

	/** Save logInfoRef type LogInfo **/
	logInfoRefStr, _ := json.Marshal(this.object.M_logInfoRef)
	ConnectingObjectInfo = append(ConnectingObjectInfo, string(logInfoRefStr))

	/** members of ConnectingObject **/

	/** Save connectingObjectType type ConnectingObjectType **/
	if this.object.M_connectingObjectType == BPMS.ConnectingObjectType_SequenceFlow {
		ConnectingObjectInfo = append(ConnectingObjectInfo, 0)
	} else if this.object.M_connectingObjectType == BPMS.ConnectingObjectType_MessageFlow {
		ConnectingObjectInfo = append(ConnectingObjectInfo, 1)
	} else if this.object.M_connectingObjectType == BPMS.ConnectingObjectType_Association {
		ConnectingObjectInfo = append(ConnectingObjectInfo, 2)
	} else if this.object.M_connectingObjectType == BPMS.ConnectingObjectType_DataAssociation {
		ConnectingObjectInfo = append(ConnectingObjectInfo, 3)
	} else {
		ConnectingObjectInfo = append(ConnectingObjectInfo, 0)
	}

	/** Save sourceRef type FlowNodeInstance **/
	ConnectingObjectInfo = append(ConnectingObjectInfo, this.object.M_sourceRef)

	/** Save targetRef type FlowNodeInstance **/
	ConnectingObjectInfo = append(ConnectingObjectInfo, this.object.M_targetRef)

	/** associations of ConnectingObject **/

	/** Save SubprocessInstance type SubprocessInstance **/
	ConnectingObjectInfo = append(ConnectingObjectInfo, this.object.M_SubprocessInstancePtr)

	/** Save processInstance type ProcessInstance **/
	ConnectingObjectInfo = append(ConnectingObjectInfo, this.object.M_processInstancePtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ConnectingObjectInfo = append(ConnectingObjectInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ConnectingObjectInfo = append(ConnectingObjectInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), ConnectingObjectInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), ConnectingObjectInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_ConnectingObjectEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_ConnectingObjectEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.ConnectingObject"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of ConnectingObject **/
	query.Fields = append(query.Fields, "M_connectingObjectType")
	query.Fields = append(query.Fields, "M_sourceRef")
	query.Fields = append(query.Fields, "M_targetRef")

	/** associations of ConnectingObject **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ConnectingObject...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.ConnectingObject)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.ConnectingObject"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** bpmnElementId **/
		if results[0][3] != nil {
			this.object.M_bpmnElementId = results[0][3].(string)
		}

		/** participants **/
		if results[0][4] != nil {
			this.object.M_participants = append(this.object.M_participants, results[0][4].([]string)...)
		}

		/** dataRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ItemAwareElementInstance"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef, ids[i])
					GetServer().GetEntityManager().appendReference("dataRef", this.object.UUID, id_)
				}
			}
		}

		/** data **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_ItemAwareElementInstanceEntity)
					} else {
						dataEntity = GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data", dataEntity)
				}
			}
		}

		/** logInfoRef **/
		if results[0][7] != nil {
			idsStr := results[0][7].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.LogInfo"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef, ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef", this.object.UUID, id_)
				}
			}
		}

		/** members of ConnectingObject **/

		/** connectingObjectType **/
		if results[0][8] != nil {
			enumIndex := results[0][8].(int)
			if enumIndex == 0 {
				this.object.M_connectingObjectType = BPMS.ConnectingObjectType_SequenceFlow
			} else if enumIndex == 1 {
				this.object.M_connectingObjectType = BPMS.ConnectingObjectType_MessageFlow
			} else if enumIndex == 2 {
				this.object.M_connectingObjectType = BPMS.ConnectingObjectType_Association
			} else if enumIndex == 3 {
				this.object.M_connectingObjectType = BPMS.ConnectingObjectType_DataAssociation
			}
		}

		/** sourceRef **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.FlowNodeInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_sourceRef = id
				GetServer().GetEntityManager().appendReference("sourceRef", this.object.UUID, id_)
			}
		}

		/** targetRef **/
		if results[0][10] != nil {
			id := results[0][10].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.FlowNodeInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_targetRef = id
				GetServer().GetEntityManager().appendReference("targetRef", this.object.UUID, id_)
			}
		}

		/** associations of ConnectingObject **/

		/** SubprocessInstancePtr **/
		if results[0][11] != nil {
			id := results[0][11].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.SubprocessInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_SubprocessInstancePtr = id
				GetServer().GetEntityManager().appendReference("SubprocessInstancePtr", this.object.UUID, id_)
			}
		}

		/** processInstancePtr **/
		if results[0][12] != nil {
			id := results[0][12].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.ProcessInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_processInstancePtr = id
				GetServer().GetEntityManager().appendReference("processInstancePtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][13].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][14].(string)
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
func (this *EntityManager) NewBPMSConnectingObjectEntityFromObject(object *BPMS.ConnectingObject) *BPMS_ConnectingObjectEntity {
	return this.NewBPMSConnectingObjectEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_ConnectingObjectEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSConnectingObjectExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.ConnectingObject"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_ConnectingObjectEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_ConnectingObjectEntity) AppendReference(reference Entity) {

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
func (this *EntityManager) create_BPMS_FlowNodeInstanceEntityPrototype() {

	var flowNodeInstanceEntityProto EntityPrototype
	flowNodeInstanceEntityProto.TypeName = "BPMS.FlowNodeInstance"
	flowNodeInstanceEntityProto.IsAbstract = true
	flowNodeInstanceEntityProto.SuperTypeNames = append(flowNodeInstanceEntityProto.SuperTypeNames, "BPMS.Instance")
	flowNodeInstanceEntityProto.SubstitutionGroup = append(flowNodeInstanceEntityProto.SubstitutionGroup, "BPMS.ActivityInstance")
	flowNodeInstanceEntityProto.SubstitutionGroup = append(flowNodeInstanceEntityProto.SubstitutionGroup, "BPMS.SubprocessInstance")
	flowNodeInstanceEntityProto.SubstitutionGroup = append(flowNodeInstanceEntityProto.SubstitutionGroup, "BPMS.GatewayInstance")
	flowNodeInstanceEntityProto.SubstitutionGroup = append(flowNodeInstanceEntityProto.SubstitutionGroup, "BPMS.EventInstance")
	flowNodeInstanceEntityProto.Ids = append(flowNodeInstanceEntityProto.Ids, "uuid")
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "uuid")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "xs.string")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 0)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, false)
	flowNodeInstanceEntityProto.Indexs = append(flowNodeInstanceEntityProto.Indexs, "parentUuid")
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "parentUuid")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "xs.string")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 1)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, false)

	/** members of Instance **/
	flowNodeInstanceEntityProto.Ids = append(flowNodeInstanceEntityProto.Ids, "M_id")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 2)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_id")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "xs.ID")
	flowNodeInstanceEntityProto.Indexs = append(flowNodeInstanceEntityProto.Indexs, "M_bpmnElementId")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 3)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_bpmnElementId")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "xs.string")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 4)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_participants")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "[]xs.string")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 5)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_dataRef")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 6)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_data")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 7)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_logInfoRef")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "[]BPMS.LogInfo:Ref")

	/** members of FlowNodeInstance **/
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 8)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_flowNodeType")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "enum:FlowNodeType_AbstractTask:FlowNodeType_ServiceTask:FlowNodeType_UserTask:FlowNodeType_ManualTask:FlowNodeType_BusinessRuleTask:FlowNodeType_ScriptTask:FlowNodeType_EmbeddedSubprocess:FlowNodeType_EventSubprocess:FlowNodeType_AdHocSubprocess:FlowNodeType_Transaction:FlowNodeType_CallActivity:FlowNodeType_ParallelGateway:FlowNodeType_ExclusiveGateway:FlowNodeType_InclusiveGateway:FlowNodeType_EventBasedGateway:FlowNodeType_ComplexGateway:FlowNodeType_StartEvent:FlowNodeType_IntermediateCatchEvent:FlowNodeType_BoundaryEvent:FlowNodeType_EndEvent:FlowNodeType_IntermediateThrowEvent")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 9)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_lifecycleState")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "enum:LifecycleState_Completed:LifecycleState_Compensated:LifecycleState_Failed:LifecycleState_Terminated:LifecycleState_Completing:LifecycleState_Compensating:LifecycleState_Failing:LifecycleState_Terminating:LifecycleState_Ready:LifecycleState_Active")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 10)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_inputRef")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject:Ref")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 11)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_outputRef")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject:Ref")

	/** associations of FlowNodeInstance **/
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 12)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, false)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_SubprocessInstancePtr")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "BPMS.SubprocessInstance:Ref")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 13)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, false)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "M_processInstancePtr")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "BPMS.ProcessInstance:Ref")
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "childsUuid")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "[]xs.string")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 14)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, false)

	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields, "referenced")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType, "[]EntityRef")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder, 15)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&flowNodeInstanceEntityProto)

}

////////////////////////////////////////////////////////////////////////////////
//              			ActivityInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_ActivityInstanceEntity struct {
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
	object         *BPMS.ActivityInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMSActivityInstanceEntity(objectId string, object interface{}) *BPMS_ActivityInstanceEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSActivityInstanceExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.ActivityInstance).TYPENAME = "BPMS.ActivityInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.ActivityInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_ActivityInstanceEntity)
		}
	} else {
		uuidStr = "BPMS.ActivityInstance%" + Utility.RandomUUID()
	}
	entity := new(BPMS_ActivityInstanceEntity)
	if object == nil {
		entity.object = new(BPMS.ActivityInstance)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.ActivityInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.ActivityInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.ActivityInstance", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_ActivityInstanceEntity) GetTypeName() string {
	return "BPMS.ActivityInstance"
}
func (this *BPMS_ActivityInstanceEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_ActivityInstanceEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_ActivityInstanceEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_ActivityInstanceEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_ActivityInstanceEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_ActivityInstanceEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_ActivityInstanceEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_ActivityInstanceEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_ActivityInstanceEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_ActivityInstanceEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_ActivityInstanceEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_ActivityInstanceEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_ActivityInstanceEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_ActivityInstanceEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_ActivityInstanceEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_ActivityInstanceEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_ActivityInstanceEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_ActivityInstanceEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_ActivityInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_ActivityInstanceEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_ActivityInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_ActivityInstanceEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_ActivityInstanceEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.ActivityInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_ActivityInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_ActivityInstanceEntityPrototype() {

	var activityInstanceEntityProto EntityPrototype
	activityInstanceEntityProto.TypeName = "BPMS.ActivityInstance"
	activityInstanceEntityProto.SuperTypeNames = append(activityInstanceEntityProto.SuperTypeNames, "BPMS.Instance")
	activityInstanceEntityProto.SuperTypeNames = append(activityInstanceEntityProto.SuperTypeNames, "BPMS.FlowNodeInstance")
	activityInstanceEntityProto.Ids = append(activityInstanceEntityProto.Ids, "uuid")
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "uuid")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "xs.string")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 0)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, false)
	activityInstanceEntityProto.Indexs = append(activityInstanceEntityProto.Indexs, "parentUuid")
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "parentUuid")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "xs.string")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 1)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, false)

	/** members of Instance **/
	activityInstanceEntityProto.Ids = append(activityInstanceEntityProto.Ids, "M_id")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 2)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_id")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "xs.ID")
	activityInstanceEntityProto.Indexs = append(activityInstanceEntityProto.Indexs, "M_bpmnElementId")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 3)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_bpmnElementId")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "xs.string")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 4)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_participants")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "[]xs.string")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 5)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_dataRef")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 6)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_data")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 7)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_logInfoRef")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "[]BPMS.LogInfo:Ref")

	/** members of FlowNodeInstance **/
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 8)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_flowNodeType")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "enum:FlowNodeType_AbstractTask:FlowNodeType_ServiceTask:FlowNodeType_UserTask:FlowNodeType_ManualTask:FlowNodeType_BusinessRuleTask:FlowNodeType_ScriptTask:FlowNodeType_EmbeddedSubprocess:FlowNodeType_EventSubprocess:FlowNodeType_AdHocSubprocess:FlowNodeType_Transaction:FlowNodeType_CallActivity:FlowNodeType_ParallelGateway:FlowNodeType_ExclusiveGateway:FlowNodeType_InclusiveGateway:FlowNodeType_EventBasedGateway:FlowNodeType_ComplexGateway:FlowNodeType_StartEvent:FlowNodeType_IntermediateCatchEvent:FlowNodeType_BoundaryEvent:FlowNodeType_EndEvent:FlowNodeType_IntermediateThrowEvent")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 9)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_lifecycleState")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "enum:LifecycleState_Completed:LifecycleState_Compensated:LifecycleState_Failed:LifecycleState_Terminated:LifecycleState_Completing:LifecycleState_Compensating:LifecycleState_Failing:LifecycleState_Terminating:LifecycleState_Ready:LifecycleState_Active")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 10)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_inputRef")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject:Ref")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 11)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_outputRef")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject:Ref")

	/** members of ActivityInstance **/
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 12)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_activityType")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "enum:ActivityType_AbstractTask:ActivityType_ServiceTask:ActivityType_UserTask:ActivityType_ManualTask:ActivityType_BusinessRuleTask:ActivityType_ScriptTask:ActivityType_EmbeddedSubprocess:ActivityType_EventSubprocess:ActivityType_AdHocSubprocess:ActivityType_Transaction:ActivityType_CallActivity")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 13)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_multiInstanceBehaviorType")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "enum:MultiInstanceBehaviorType_None:MultiInstanceBehaviorType_One:MultiInstanceBehaviorType_All:MultiInstanceBehaviorType_Complex")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 14)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_loopCharacteristicType")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "enum:LoopCharacteristicType_StandardLoopCharacteristics:LoopCharacteristicType_MultiInstanceLoopCharacteristics")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 15)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_tokenCount")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "xs.int")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 16)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_ressources")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "[]BPMS.RessourceInstance")

	/** associations of ActivityInstance **/
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 17)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, false)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_SubprocessInstancePtr")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "BPMS.SubprocessInstance:Ref")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 18)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, false)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "M_processInstancePtr")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "BPMS.ProcessInstance:Ref")
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "childsUuid")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "[]xs.string")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 19)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, false)

	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields, "referenced")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType, "[]EntityRef")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder, 20)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&activityInstanceEntityProto)

}

/** Create **/
func (this *BPMS_ActivityInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.ActivityInstance"

	var query EntityQuery
	query.TypeName = "BPMS.ActivityInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of FlowNodeInstance **/
	query.Fields = append(query.Fields, "M_flowNodeType")
	query.Fields = append(query.Fields, "M_lifecycleState")
	query.Fields = append(query.Fields, "M_inputRef")
	query.Fields = append(query.Fields, "M_outputRef")

	/** members of ActivityInstance **/
	query.Fields = append(query.Fields, "M_activityType")
	query.Fields = append(query.Fields, "M_multiInstanceBehaviorType")
	query.Fields = append(query.Fields, "M_loopCharacteristicType")
	query.Fields = append(query.Fields, "M_tokenCount")
	query.Fields = append(query.Fields, "M_ressources")

	/** associations of ActivityInstance **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ActivityInstanceInfo []interface{}

	ActivityInstanceInfo = append(ActivityInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		ActivityInstanceInfo = append(ActivityInstanceInfo, this.parentPtr.GetUuid())
	} else {
		ActivityInstanceInfo = append(ActivityInstanceInfo, "")
	}

	/** members of Instance **/
	ActivityInstanceInfo = append(ActivityInstanceInfo, this.object.M_id)
	ActivityInstanceInfo = append(ActivityInstanceInfo, this.object.M_bpmnElementId)
	ActivityInstanceInfo = append(ActivityInstanceInfo, this.object.M_participants)

	/** Save dataRef type ItemAwareElementInstance **/
	dataRefStr, _ := json.Marshal(this.object.M_dataRef)
	ActivityInstanceInfo = append(ActivityInstanceInfo, string(dataRefStr))

	/** Save data type ItemAwareElementInstance **/
	dataIds := make([]string, 0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity := GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(this.object.M_data[i].UUID, this.object.M_data[i])
		dataIds = append(dataIds, dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data", dataEntity)
		if dataEntity.NeedSave() {
			dataEntity.SaveEntity()
		}
	}
	dataStr, _ := json.Marshal(dataIds)
	ActivityInstanceInfo = append(ActivityInstanceInfo, string(dataStr))

	/** Save logInfoRef type LogInfo **/
	logInfoRefStr, _ := json.Marshal(this.object.M_logInfoRef)
	ActivityInstanceInfo = append(ActivityInstanceInfo, string(logInfoRefStr))

	/** members of FlowNodeInstance **/

	/** Save flowNodeType type FlowNodeType **/
	if this.object.M_flowNodeType == BPMS.FlowNodeType_AbstractTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ServiceTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 1)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_UserTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 2)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ManualTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 3)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_BusinessRuleTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 4)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ScriptTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 5)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EmbeddedSubprocess {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 6)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EventSubprocess {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 7)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_AdHocSubprocess {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 8)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_Transaction {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 9)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_CallActivity {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 10)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ParallelGateway {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 11)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ExclusiveGateway {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 12)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_InclusiveGateway {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 13)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EventBasedGateway {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 14)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ComplexGateway {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 15)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_StartEvent {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 16)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_IntermediateCatchEvent {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 17)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_BoundaryEvent {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 18)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EndEvent {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 19)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_IntermediateThrowEvent {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 20)
	} else {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	}

	/** Save lifecycleState type LifecycleState **/
	if this.object.M_lifecycleState == BPMS.LifecycleState_Completed {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Compensated {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 1)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Failed {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 2)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Terminated {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 3)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Completing {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 4)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Compensating {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 5)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Failing {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 6)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Terminating {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 7)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Ready {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 8)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Active {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 9)
	} else {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	}

	/** Save inputRef type ConnectingObject **/
	inputRefStr, _ := json.Marshal(this.object.M_inputRef)
	ActivityInstanceInfo = append(ActivityInstanceInfo, string(inputRefStr))

	/** Save outputRef type ConnectingObject **/
	outputRefStr, _ := json.Marshal(this.object.M_outputRef)
	ActivityInstanceInfo = append(ActivityInstanceInfo, string(outputRefStr))

	/** members of ActivityInstance **/

	/** Save activityType type ActivityType **/
	if this.object.M_activityType == BPMS.ActivityType_AbstractTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	} else if this.object.M_activityType == BPMS.ActivityType_ServiceTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 1)
	} else if this.object.M_activityType == BPMS.ActivityType_UserTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 2)
	} else if this.object.M_activityType == BPMS.ActivityType_ManualTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 3)
	} else if this.object.M_activityType == BPMS.ActivityType_BusinessRuleTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 4)
	} else if this.object.M_activityType == BPMS.ActivityType_ScriptTask {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 5)
	} else if this.object.M_activityType == BPMS.ActivityType_EmbeddedSubprocess {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 6)
	} else if this.object.M_activityType == BPMS.ActivityType_EventSubprocess {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 7)
	} else if this.object.M_activityType == BPMS.ActivityType_AdHocSubprocess {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 8)
	} else if this.object.M_activityType == BPMS.ActivityType_Transaction {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 9)
	} else if this.object.M_activityType == BPMS.ActivityType_CallActivity {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 10)
	} else {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	}

	/** Save multiInstanceBehaviorType type MultiInstanceBehaviorType **/
	if this.object.M_multiInstanceBehaviorType == BPMS.MultiInstanceBehaviorType_None {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	} else if this.object.M_multiInstanceBehaviorType == BPMS.MultiInstanceBehaviorType_One {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 1)
	} else if this.object.M_multiInstanceBehaviorType == BPMS.MultiInstanceBehaviorType_All {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 2)
	} else if this.object.M_multiInstanceBehaviorType == BPMS.MultiInstanceBehaviorType_Complex {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 3)
	} else {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	}

	/** Save loopCharacteristicType type LoopCharacteristicType **/
	if this.object.M_loopCharacteristicType == BPMS.LoopCharacteristicType_StandardLoopCharacteristics {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	} else if this.object.M_loopCharacteristicType == BPMS.LoopCharacteristicType_MultiInstanceLoopCharacteristics {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 1)
	} else {
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	}
	ActivityInstanceInfo = append(ActivityInstanceInfo, this.object.M_tokenCount)

	/** Save ressources type RessourceInstance **/
	ressourcesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_ressources); i++ {
		ressourcesEntity := GetServer().GetEntityManager().NewBPMSRessourceInstanceEntity(this.object.M_ressources[i].UUID, this.object.M_ressources[i])
		ressourcesIds = append(ressourcesIds, ressourcesEntity.uuid)
		ressourcesEntity.AppendReferenced("ressources", this)
		this.AppendChild("ressources", ressourcesEntity)
		if ressourcesEntity.NeedSave() {
			ressourcesEntity.SaveEntity()
		}
	}
	ressourcesStr, _ := json.Marshal(ressourcesIds)
	ActivityInstanceInfo = append(ActivityInstanceInfo, string(ressourcesStr))

	/** associations of ActivityInstance **/

	/** Save SubprocessInstance type SubprocessInstance **/
	ActivityInstanceInfo = append(ActivityInstanceInfo, this.object.M_SubprocessInstancePtr)

	/** Save processInstance type ProcessInstance **/
	ActivityInstanceInfo = append(ActivityInstanceInfo, this.object.M_processInstancePtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ActivityInstanceInfo = append(ActivityInstanceInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ActivityInstanceInfo = append(ActivityInstanceInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), ActivityInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), ActivityInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_ActivityInstanceEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_ActivityInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.ActivityInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of FlowNodeInstance **/
	query.Fields = append(query.Fields, "M_flowNodeType")
	query.Fields = append(query.Fields, "M_lifecycleState")
	query.Fields = append(query.Fields, "M_inputRef")
	query.Fields = append(query.Fields, "M_outputRef")

	/** members of ActivityInstance **/
	query.Fields = append(query.Fields, "M_activityType")
	query.Fields = append(query.Fields, "M_multiInstanceBehaviorType")
	query.Fields = append(query.Fields, "M_loopCharacteristicType")
	query.Fields = append(query.Fields, "M_tokenCount")
	query.Fields = append(query.Fields, "M_ressources")

	/** associations of ActivityInstance **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ActivityInstance...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.ActivityInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.ActivityInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** bpmnElementId **/
		if results[0][3] != nil {
			this.object.M_bpmnElementId = results[0][3].(string)
		}

		/** participants **/
		if results[0][4] != nil {
			this.object.M_participants = append(this.object.M_participants, results[0][4].([]string)...)
		}

		/** dataRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ItemAwareElementInstance"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef, ids[i])
					GetServer().GetEntityManager().appendReference("dataRef", this.object.UUID, id_)
				}
			}
		}

		/** data **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_ItemAwareElementInstanceEntity)
					} else {
						dataEntity = GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data", dataEntity)
				}
			}
		}

		/** logInfoRef **/
		if results[0][7] != nil {
			idsStr := results[0][7].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.LogInfo"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef, ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef", this.object.UUID, id_)
				}
			}
		}

		/** members of FlowNodeInstance **/

		/** flowNodeType **/
		if results[0][8] != nil {
			enumIndex := results[0][8].(int)
			if enumIndex == 0 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_AbstractTask
			} else if enumIndex == 1 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ServiceTask
			} else if enumIndex == 2 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_UserTask
			} else if enumIndex == 3 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ManualTask
			} else if enumIndex == 4 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_BusinessRuleTask
			} else if enumIndex == 5 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ScriptTask
			} else if enumIndex == 6 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EmbeddedSubprocess
			} else if enumIndex == 7 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EventSubprocess
			} else if enumIndex == 8 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_AdHocSubprocess
			} else if enumIndex == 9 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_Transaction
			} else if enumIndex == 10 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_CallActivity
			} else if enumIndex == 11 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ParallelGateway
			} else if enumIndex == 12 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ExclusiveGateway
			} else if enumIndex == 13 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_InclusiveGateway
			} else if enumIndex == 14 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EventBasedGateway
			} else if enumIndex == 15 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ComplexGateway
			} else if enumIndex == 16 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_StartEvent
			} else if enumIndex == 17 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_IntermediateCatchEvent
			} else if enumIndex == 18 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_BoundaryEvent
			} else if enumIndex == 19 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EndEvent
			} else if enumIndex == 20 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_IntermediateThrowEvent
			}
		}

		/** lifecycleState **/
		if results[0][9] != nil {
			enumIndex := results[0][9].(int)
			if enumIndex == 0 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Completed
			} else if enumIndex == 1 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Compensated
			} else if enumIndex == 2 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Failed
			} else if enumIndex == 3 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Terminated
			} else if enumIndex == 4 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Completing
			} else if enumIndex == 5 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Compensating
			} else if enumIndex == 6 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Failing
			} else if enumIndex == 7 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Terminating
			} else if enumIndex == 8 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Ready
			} else if enumIndex == 9 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Active
			}
		}

		/** inputRef **/
		if results[0][10] != nil {
			idsStr := results[0][10].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ConnectingObject"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_inputRef = append(this.object.M_inputRef, ids[i])
					GetServer().GetEntityManager().appendReference("inputRef", this.object.UUID, id_)
				}
			}
		}

		/** outputRef **/
		if results[0][11] != nil {
			idsStr := results[0][11].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ConnectingObject"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_outputRef = append(this.object.M_outputRef, ids[i])
					GetServer().GetEntityManager().appendReference("outputRef", this.object.UUID, id_)
				}
			}
		}

		/** members of ActivityInstance **/

		/** activityType **/
		if results[0][12] != nil {
			enumIndex := results[0][12].(int)
			if enumIndex == 0 {
				this.object.M_activityType = BPMS.ActivityType_AbstractTask
			} else if enumIndex == 1 {
				this.object.M_activityType = BPMS.ActivityType_ServiceTask
			} else if enumIndex == 2 {
				this.object.M_activityType = BPMS.ActivityType_UserTask
			} else if enumIndex == 3 {
				this.object.M_activityType = BPMS.ActivityType_ManualTask
			} else if enumIndex == 4 {
				this.object.M_activityType = BPMS.ActivityType_BusinessRuleTask
			} else if enumIndex == 5 {
				this.object.M_activityType = BPMS.ActivityType_ScriptTask
			} else if enumIndex == 6 {
				this.object.M_activityType = BPMS.ActivityType_EmbeddedSubprocess
			} else if enumIndex == 7 {
				this.object.M_activityType = BPMS.ActivityType_EventSubprocess
			} else if enumIndex == 8 {
				this.object.M_activityType = BPMS.ActivityType_AdHocSubprocess
			} else if enumIndex == 9 {
				this.object.M_activityType = BPMS.ActivityType_Transaction
			} else if enumIndex == 10 {
				this.object.M_activityType = BPMS.ActivityType_CallActivity
			}
		}

		/** multiInstanceBehaviorType **/
		if results[0][13] != nil {
			enumIndex := results[0][13].(int)
			if enumIndex == 0 {
				this.object.M_multiInstanceBehaviorType = BPMS.MultiInstanceBehaviorType_None
			} else if enumIndex == 1 {
				this.object.M_multiInstanceBehaviorType = BPMS.MultiInstanceBehaviorType_One
			} else if enumIndex == 2 {
				this.object.M_multiInstanceBehaviorType = BPMS.MultiInstanceBehaviorType_All
			} else if enumIndex == 3 {
				this.object.M_multiInstanceBehaviorType = BPMS.MultiInstanceBehaviorType_Complex
			}
		}

		/** loopCharacteristicType **/
		if results[0][14] != nil {
			enumIndex := results[0][14].(int)
			if enumIndex == 0 {
				this.object.M_loopCharacteristicType = BPMS.LoopCharacteristicType_StandardLoopCharacteristics
			} else if enumIndex == 1 {
				this.object.M_loopCharacteristicType = BPMS.LoopCharacteristicType_MultiInstanceLoopCharacteristics
			}
		}

		/** tokenCount **/
		if results[0][15] != nil {
			this.object.M_tokenCount = results[0][15].(int)
		}

		/** ressources **/
		if results[0][16] != nil {
			uuidsStr := results[0][16].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var ressourcesEntity *BPMS_RessourceInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						ressourcesEntity = instance.(*BPMS_RessourceInstanceEntity)
					} else {
						ressourcesEntity = GetServer().GetEntityManager().NewBPMSRessourceInstanceEntity(uuids[i], nil)
						ressourcesEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(ressourcesEntity)
					}
					ressourcesEntity.AppendReferenced("ressources", this)
					this.AppendChild("ressources", ressourcesEntity)
				}
			}
		}

		/** associations of ActivityInstance **/

		/** SubprocessInstancePtr **/
		if results[0][17] != nil {
			id := results[0][17].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.SubprocessInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_SubprocessInstancePtr = id
				GetServer().GetEntityManager().appendReference("SubprocessInstancePtr", this.object.UUID, id_)
			}
		}

		/** processInstancePtr **/
		if results[0][18] != nil {
			id := results[0][18].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.ProcessInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_processInstancePtr = id
				GetServer().GetEntityManager().appendReference("processInstancePtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][19].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][20].(string)
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
func (this *EntityManager) NewBPMSActivityInstanceEntityFromObject(object *BPMS.ActivityInstance) *BPMS_ActivityInstanceEntity {
	return this.NewBPMSActivityInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_ActivityInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSActivityInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.ActivityInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_ActivityInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_ActivityInstanceEntity) AppendReference(reference Entity) {

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
//              			SubprocessInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_SubprocessInstanceEntity struct {
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
	object         *BPMS.SubprocessInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMSSubprocessInstanceEntity(objectId string, object interface{}) *BPMS_SubprocessInstanceEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSSubprocessInstanceExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.SubprocessInstance).TYPENAME = "BPMS.SubprocessInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.SubprocessInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_SubprocessInstanceEntity)
		}
	} else {
		uuidStr = "BPMS.SubprocessInstance%" + Utility.RandomUUID()
	}
	entity := new(BPMS_SubprocessInstanceEntity)
	if object == nil {
		entity.object = new(BPMS.SubprocessInstance)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.SubprocessInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.SubprocessInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.SubprocessInstance", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_SubprocessInstanceEntity) GetTypeName() string {
	return "BPMS.SubprocessInstance"
}
func (this *BPMS_SubprocessInstanceEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_SubprocessInstanceEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_SubprocessInstanceEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_SubprocessInstanceEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_SubprocessInstanceEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_SubprocessInstanceEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_SubprocessInstanceEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_SubprocessInstanceEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_SubprocessInstanceEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_SubprocessInstanceEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_SubprocessInstanceEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_SubprocessInstanceEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_SubprocessInstanceEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_SubprocessInstanceEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_SubprocessInstanceEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_SubprocessInstanceEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_SubprocessInstanceEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_SubprocessInstanceEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_SubprocessInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_SubprocessInstanceEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_SubprocessInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_SubprocessInstanceEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_SubprocessInstanceEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.SubprocessInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_SubprocessInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_SubprocessInstanceEntityPrototype() {

	var subprocessInstanceEntityProto EntityPrototype
	subprocessInstanceEntityProto.TypeName = "BPMS.SubprocessInstance"
	subprocessInstanceEntityProto.SuperTypeNames = append(subprocessInstanceEntityProto.SuperTypeNames, "BPMS.Instance")
	subprocessInstanceEntityProto.SuperTypeNames = append(subprocessInstanceEntityProto.SuperTypeNames, "BPMS.FlowNodeInstance")
	subprocessInstanceEntityProto.Ids = append(subprocessInstanceEntityProto.Ids, "uuid")
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "uuid")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "xs.string")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 0)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, false)
	subprocessInstanceEntityProto.Indexs = append(subprocessInstanceEntityProto.Indexs, "parentUuid")
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "parentUuid")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "xs.string")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 1)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, false)

	/** members of Instance **/
	subprocessInstanceEntityProto.Ids = append(subprocessInstanceEntityProto.Ids, "M_id")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 2)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_id")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "xs.ID")
	subprocessInstanceEntityProto.Indexs = append(subprocessInstanceEntityProto.Indexs, "M_bpmnElementId")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 3)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_bpmnElementId")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "xs.string")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 4)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_participants")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]xs.string")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 5)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_dataRef")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 6)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_data")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 7)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_logInfoRef")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]BPMS.LogInfo:Ref")

	/** members of FlowNodeInstance **/
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 8)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_flowNodeType")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "enum:FlowNodeType_AbstractTask:FlowNodeType_ServiceTask:FlowNodeType_UserTask:FlowNodeType_ManualTask:FlowNodeType_BusinessRuleTask:FlowNodeType_ScriptTask:FlowNodeType_EmbeddedSubprocess:FlowNodeType_EventSubprocess:FlowNodeType_AdHocSubprocess:FlowNodeType_Transaction:FlowNodeType_CallActivity:FlowNodeType_ParallelGateway:FlowNodeType_ExclusiveGateway:FlowNodeType_InclusiveGateway:FlowNodeType_EventBasedGateway:FlowNodeType_ComplexGateway:FlowNodeType_StartEvent:FlowNodeType_IntermediateCatchEvent:FlowNodeType_BoundaryEvent:FlowNodeType_EndEvent:FlowNodeType_IntermediateThrowEvent")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 9)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_lifecycleState")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "enum:LifecycleState_Completed:LifecycleState_Compensated:LifecycleState_Failed:LifecycleState_Terminated:LifecycleState_Completing:LifecycleState_Compensating:LifecycleState_Failing:LifecycleState_Terminating:LifecycleState_Ready:LifecycleState_Active")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 10)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_inputRef")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject:Ref")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 11)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_outputRef")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject:Ref")

	/** members of SubprocessInstance **/
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 12)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_SubprocessType")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "enum:SubprocessType_EmbeddedSubprocess:SubprocessType_EventSubprocess:SubprocessType_AdHocSubprocess:SubprocessType_Transaction")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 13)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_flowNodeInstances")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]BPMS.FlowNodeInstance")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 14)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_connectingObjects")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 15)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_ressources")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]BPMS.RessourceInstance")

	/** associations of SubprocessInstance **/
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 16)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, false)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_SubprocessInstancePtr")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "BPMS.SubprocessInstance:Ref")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 17)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, false)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "M_processInstancePtr")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "BPMS.ProcessInstance:Ref")
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "childsUuid")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]xs.string")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 18)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, false)

	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields, "referenced")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType, "[]EntityRef")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder, 19)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&subprocessInstanceEntityProto)

}

/** Create **/
func (this *BPMS_SubprocessInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.SubprocessInstance"

	var query EntityQuery
	query.TypeName = "BPMS.SubprocessInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of FlowNodeInstance **/
	query.Fields = append(query.Fields, "M_flowNodeType")
	query.Fields = append(query.Fields, "M_lifecycleState")
	query.Fields = append(query.Fields, "M_inputRef")
	query.Fields = append(query.Fields, "M_outputRef")

	/** members of SubprocessInstance **/
	query.Fields = append(query.Fields, "M_SubprocessType")
	query.Fields = append(query.Fields, "M_flowNodeInstances")
	query.Fields = append(query.Fields, "M_connectingObjects")
	query.Fields = append(query.Fields, "M_ressources")

	/** associations of SubprocessInstance **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var SubprocessInstanceInfo []interface{}

	SubprocessInstanceInfo = append(SubprocessInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, this.parentPtr.GetUuid())
	} else {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, "")
	}

	/** members of Instance **/
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, this.object.M_id)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, this.object.M_bpmnElementId)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, this.object.M_participants)

	/** Save dataRef type ItemAwareElementInstance **/
	dataRefStr, _ := json.Marshal(this.object.M_dataRef)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(dataRefStr))

	/** Save data type ItemAwareElementInstance **/
	dataIds := make([]string, 0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity := GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(this.object.M_data[i].UUID, this.object.M_data[i])
		dataIds = append(dataIds, dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data", dataEntity)
		if dataEntity.NeedSave() {
			dataEntity.SaveEntity()
		}
	}
	dataStr, _ := json.Marshal(dataIds)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(dataStr))

	/** Save logInfoRef type LogInfo **/
	logInfoRefStr, _ := json.Marshal(this.object.M_logInfoRef)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(logInfoRefStr))

	/** members of FlowNodeInstance **/

	/** Save flowNodeType type FlowNodeType **/
	if this.object.M_flowNodeType == BPMS.FlowNodeType_AbstractTask {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ServiceTask {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 1)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_UserTask {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 2)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ManualTask {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 3)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_BusinessRuleTask {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 4)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ScriptTask {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 5)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EmbeddedSubprocess {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 6)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EventSubprocess {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 7)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_AdHocSubprocess {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 8)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_Transaction {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 9)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_CallActivity {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 10)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ParallelGateway {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 11)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ExclusiveGateway {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 12)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_InclusiveGateway {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 13)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EventBasedGateway {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 14)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ComplexGateway {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 15)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_StartEvent {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 16)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_IntermediateCatchEvent {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 17)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_BoundaryEvent {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 18)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EndEvent {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 19)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_IntermediateThrowEvent {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 20)
	} else {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	}

	/** Save lifecycleState type LifecycleState **/
	if this.object.M_lifecycleState == BPMS.LifecycleState_Completed {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Compensated {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 1)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Failed {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 2)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Terminated {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 3)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Completing {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 4)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Compensating {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 5)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Failing {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 6)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Terminating {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 7)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Ready {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 8)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Active {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 9)
	} else {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	}

	/** Save inputRef type ConnectingObject **/
	inputRefStr, _ := json.Marshal(this.object.M_inputRef)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(inputRefStr))

	/** Save outputRef type ConnectingObject **/
	outputRefStr, _ := json.Marshal(this.object.M_outputRef)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(outputRefStr))

	/** members of SubprocessInstance **/

	/** Save SubprocessType type SubprocessType **/
	if this.object.M_SubprocessType == BPMS.SubprocessType_EmbeddedSubprocess {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	} else if this.object.M_SubprocessType == BPMS.SubprocessType_EventSubprocess {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 1)
	} else if this.object.M_SubprocessType == BPMS.SubprocessType_AdHocSubprocess {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 2)
	} else if this.object.M_SubprocessType == BPMS.SubprocessType_Transaction {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 3)
	} else {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	}

	/** Save flowNodeInstances type FlowNodeInstance **/
	flowNodeInstancesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_flowNodeInstances); i++ {
		switch v := this.object.M_flowNodeInstances[i].(type) {
		case *BPMS.ActivityInstance:
			flowNodeInstancesEntity := GetServer().GetEntityManager().NewBPMSActivityInstanceEntity(v.UUID, v)
			flowNodeInstancesIds = append(flowNodeInstancesIds, flowNodeInstancesEntity.uuid)
			flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
			this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
			if flowNodeInstancesEntity.NeedSave() {
				flowNodeInstancesEntity.SaveEntity()
			}
		case *BPMS.SubprocessInstance:
			flowNodeInstancesEntity := GetServer().GetEntityManager().NewBPMSSubprocessInstanceEntity(v.UUID, v)
			flowNodeInstancesIds = append(flowNodeInstancesIds, flowNodeInstancesEntity.uuid)
			flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
			this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
			if flowNodeInstancesEntity.NeedSave() {
				flowNodeInstancesEntity.SaveEntity()
			}
		case *BPMS.GatewayInstance:
			flowNodeInstancesEntity := GetServer().GetEntityManager().NewBPMSGatewayInstanceEntity(v.UUID, v)
			flowNodeInstancesIds = append(flowNodeInstancesIds, flowNodeInstancesEntity.uuid)
			flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
			this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
			if flowNodeInstancesEntity.NeedSave() {
				flowNodeInstancesEntity.SaveEntity()
			}
		case *BPMS.EventInstance:
			flowNodeInstancesEntity := GetServer().GetEntityManager().NewBPMSEventInstanceEntity(v.UUID, v)
			flowNodeInstancesIds = append(flowNodeInstancesIds, flowNodeInstancesEntity.uuid)
			flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
			this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
			if flowNodeInstancesEntity.NeedSave() {
				flowNodeInstancesEntity.SaveEntity()
			}
		}
	}
	flowNodeInstancesStr, _ := json.Marshal(flowNodeInstancesIds)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(flowNodeInstancesStr))

	/** Save connectingObjects type ConnectingObject **/
	connectingObjectsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_connectingObjects); i++ {
		connectingObjectsEntity := GetServer().GetEntityManager().NewBPMSConnectingObjectEntity(this.object.M_connectingObjects[i].UUID, this.object.M_connectingObjects[i])
		connectingObjectsIds = append(connectingObjectsIds, connectingObjectsEntity.uuid)
		connectingObjectsEntity.AppendReferenced("connectingObjects", this)
		this.AppendChild("connectingObjects", connectingObjectsEntity)
		if connectingObjectsEntity.NeedSave() {
			connectingObjectsEntity.SaveEntity()
		}
	}
	connectingObjectsStr, _ := json.Marshal(connectingObjectsIds)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(connectingObjectsStr))

	/** Save ressources type RessourceInstance **/
	ressourcesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_ressources); i++ {
		ressourcesEntity := GetServer().GetEntityManager().NewBPMSRessourceInstanceEntity(this.object.M_ressources[i].UUID, this.object.M_ressources[i])
		ressourcesIds = append(ressourcesIds, ressourcesEntity.uuid)
		ressourcesEntity.AppendReferenced("ressources", this)
		this.AppendChild("ressources", ressourcesEntity)
		if ressourcesEntity.NeedSave() {
			ressourcesEntity.SaveEntity()
		}
	}
	ressourcesStr, _ := json.Marshal(ressourcesIds)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(ressourcesStr))

	/** associations of SubprocessInstance **/

	/** Save SubprocessInstance type SubprocessInstance **/
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, this.object.M_SubprocessInstancePtr)

	/** Save processInstance type ProcessInstance **/
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, this.object.M_processInstancePtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), SubprocessInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), SubprocessInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_SubprocessInstanceEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_SubprocessInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.SubprocessInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of FlowNodeInstance **/
	query.Fields = append(query.Fields, "M_flowNodeType")
	query.Fields = append(query.Fields, "M_lifecycleState")
	query.Fields = append(query.Fields, "M_inputRef")
	query.Fields = append(query.Fields, "M_outputRef")

	/** members of SubprocessInstance **/
	query.Fields = append(query.Fields, "M_SubprocessType")
	query.Fields = append(query.Fields, "M_flowNodeInstances")
	query.Fields = append(query.Fields, "M_connectingObjects")
	query.Fields = append(query.Fields, "M_ressources")

	/** associations of SubprocessInstance **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of SubprocessInstance...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.SubprocessInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.SubprocessInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** bpmnElementId **/
		if results[0][3] != nil {
			this.object.M_bpmnElementId = results[0][3].(string)
		}

		/** participants **/
		if results[0][4] != nil {
			this.object.M_participants = append(this.object.M_participants, results[0][4].([]string)...)
		}

		/** dataRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ItemAwareElementInstance"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef, ids[i])
					GetServer().GetEntityManager().appendReference("dataRef", this.object.UUID, id_)
				}
			}
		}

		/** data **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_ItemAwareElementInstanceEntity)
					} else {
						dataEntity = GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data", dataEntity)
				}
			}
		}

		/** logInfoRef **/
		if results[0][7] != nil {
			idsStr := results[0][7].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.LogInfo"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef, ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef", this.object.UUID, id_)
				}
			}
		}

		/** members of FlowNodeInstance **/

		/** flowNodeType **/
		if results[0][8] != nil {
			enumIndex := results[0][8].(int)
			if enumIndex == 0 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_AbstractTask
			} else if enumIndex == 1 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ServiceTask
			} else if enumIndex == 2 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_UserTask
			} else if enumIndex == 3 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ManualTask
			} else if enumIndex == 4 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_BusinessRuleTask
			} else if enumIndex == 5 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ScriptTask
			} else if enumIndex == 6 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EmbeddedSubprocess
			} else if enumIndex == 7 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EventSubprocess
			} else if enumIndex == 8 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_AdHocSubprocess
			} else if enumIndex == 9 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_Transaction
			} else if enumIndex == 10 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_CallActivity
			} else if enumIndex == 11 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ParallelGateway
			} else if enumIndex == 12 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ExclusiveGateway
			} else if enumIndex == 13 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_InclusiveGateway
			} else if enumIndex == 14 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EventBasedGateway
			} else if enumIndex == 15 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ComplexGateway
			} else if enumIndex == 16 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_StartEvent
			} else if enumIndex == 17 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_IntermediateCatchEvent
			} else if enumIndex == 18 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_BoundaryEvent
			} else if enumIndex == 19 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EndEvent
			} else if enumIndex == 20 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_IntermediateThrowEvent
			}
		}

		/** lifecycleState **/
		if results[0][9] != nil {
			enumIndex := results[0][9].(int)
			if enumIndex == 0 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Completed
			} else if enumIndex == 1 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Compensated
			} else if enumIndex == 2 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Failed
			} else if enumIndex == 3 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Terminated
			} else if enumIndex == 4 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Completing
			} else if enumIndex == 5 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Compensating
			} else if enumIndex == 6 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Failing
			} else if enumIndex == 7 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Terminating
			} else if enumIndex == 8 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Ready
			} else if enumIndex == 9 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Active
			}
		}

		/** inputRef **/
		if results[0][10] != nil {
			idsStr := results[0][10].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ConnectingObject"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_inputRef = append(this.object.M_inputRef, ids[i])
					GetServer().GetEntityManager().appendReference("inputRef", this.object.UUID, id_)
				}
			}
		}

		/** outputRef **/
		if results[0][11] != nil {
			idsStr := results[0][11].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ConnectingObject"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_outputRef = append(this.object.M_outputRef, ids[i])
					GetServer().GetEntityManager().appendReference("outputRef", this.object.UUID, id_)
				}
			}
		}

		/** members of SubprocessInstance **/

		/** SubprocessType **/
		if results[0][12] != nil {
			enumIndex := results[0][12].(int)
			if enumIndex == 0 {
				this.object.M_SubprocessType = BPMS.SubprocessType_EmbeddedSubprocess
			} else if enumIndex == 1 {
				this.object.M_SubprocessType = BPMS.SubprocessType_EventSubprocess
			} else if enumIndex == 2 {
				this.object.M_SubprocessType = BPMS.SubprocessType_AdHocSubprocess
			} else if enumIndex == 3 {
				this.object.M_SubprocessType = BPMS.SubprocessType_Transaction
			}
		}

		/** flowNodeInstances **/
		if results[0][13] != nil {
			uuidsStr := results[0][13].(string)
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
				if typeName == "BPMS.GatewayInstance" {
					if len(uuids[i]) > 0 {
						var flowNodeInstancesEntity *BPMS_GatewayInstanceEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							flowNodeInstancesEntity = instance.(*BPMS_GatewayInstanceEntity)
						} else {
							flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMSGatewayInstanceEntity(uuids[i], nil)
							flowNodeInstancesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
						}
						flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
						this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
					}
				} else if typeName == "BPMS.EventInstance" {
					if len(uuids[i]) > 0 {
						var flowNodeInstancesEntity *BPMS_EventInstanceEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							flowNodeInstancesEntity = instance.(*BPMS_EventInstanceEntity)
						} else {
							flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMSEventInstanceEntity(uuids[i], nil)
							flowNodeInstancesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
						}
						flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
						this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
					}
				} else if typeName == "BPMS.ActivityInstance" {
					if len(uuids[i]) > 0 {
						var flowNodeInstancesEntity *BPMS_ActivityInstanceEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							flowNodeInstancesEntity = instance.(*BPMS_ActivityInstanceEntity)
						} else {
							flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMSActivityInstanceEntity(uuids[i], nil)
							flowNodeInstancesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
						}
						flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
						this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
					}
				} else if typeName == "BPMS.SubprocessInstance" {
					if len(uuids[i]) > 0 {
						var flowNodeInstancesEntity *BPMS_SubprocessInstanceEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							flowNodeInstancesEntity = instance.(*BPMS_SubprocessInstanceEntity)
						} else {
							flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMSSubprocessInstanceEntity(uuids[i], nil)
							flowNodeInstancesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
						}
						flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
						this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
					}
				}
			}
		}

		/** connectingObjects **/
		if results[0][14] != nil {
			uuidsStr := results[0][14].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var connectingObjectsEntity *BPMS_ConnectingObjectEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						connectingObjectsEntity = instance.(*BPMS_ConnectingObjectEntity)
					} else {
						connectingObjectsEntity = GetServer().GetEntityManager().NewBPMSConnectingObjectEntity(uuids[i], nil)
						connectingObjectsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(connectingObjectsEntity)
					}
					connectingObjectsEntity.AppendReferenced("connectingObjects", this)
					this.AppendChild("connectingObjects", connectingObjectsEntity)
				}
			}
		}

		/** ressources **/
		if results[0][15] != nil {
			uuidsStr := results[0][15].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var ressourcesEntity *BPMS_RessourceInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						ressourcesEntity = instance.(*BPMS_RessourceInstanceEntity)
					} else {
						ressourcesEntity = GetServer().GetEntityManager().NewBPMSRessourceInstanceEntity(uuids[i], nil)
						ressourcesEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(ressourcesEntity)
					}
					ressourcesEntity.AppendReferenced("ressources", this)
					this.AppendChild("ressources", ressourcesEntity)
				}
			}
		}

		/** associations of SubprocessInstance **/

		/** SubprocessInstancePtr **/
		if results[0][16] != nil {
			id := results[0][16].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.SubprocessInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_SubprocessInstancePtr = id
				GetServer().GetEntityManager().appendReference("SubprocessInstancePtr", this.object.UUID, id_)
			}
		}

		/** processInstancePtr **/
		if results[0][17] != nil {
			id := results[0][17].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.ProcessInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_processInstancePtr = id
				GetServer().GetEntityManager().appendReference("processInstancePtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][18].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][19].(string)
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
func (this *EntityManager) NewBPMSSubprocessInstanceEntityFromObject(object *BPMS.SubprocessInstance) *BPMS_SubprocessInstanceEntity {
	return this.NewBPMSSubprocessInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_SubprocessInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSSubprocessInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.SubprocessInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_SubprocessInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_SubprocessInstanceEntity) AppendReference(reference Entity) {

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
//              			GatewayInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_GatewayInstanceEntity struct {
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
	object         *BPMS.GatewayInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMSGatewayInstanceEntity(objectId string, object interface{}) *BPMS_GatewayInstanceEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSGatewayInstanceExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.GatewayInstance).TYPENAME = "BPMS.GatewayInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.GatewayInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_GatewayInstanceEntity)
		}
	} else {
		uuidStr = "BPMS.GatewayInstance%" + Utility.RandomUUID()
	}
	entity := new(BPMS_GatewayInstanceEntity)
	if object == nil {
		entity.object = new(BPMS.GatewayInstance)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.GatewayInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.GatewayInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.GatewayInstance", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_GatewayInstanceEntity) GetTypeName() string {
	return "BPMS.GatewayInstance"
}
func (this *BPMS_GatewayInstanceEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_GatewayInstanceEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_GatewayInstanceEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_GatewayInstanceEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_GatewayInstanceEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_GatewayInstanceEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_GatewayInstanceEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_GatewayInstanceEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_GatewayInstanceEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_GatewayInstanceEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_GatewayInstanceEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_GatewayInstanceEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_GatewayInstanceEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_GatewayInstanceEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_GatewayInstanceEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_GatewayInstanceEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_GatewayInstanceEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_GatewayInstanceEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_GatewayInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_GatewayInstanceEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_GatewayInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_GatewayInstanceEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_GatewayInstanceEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.GatewayInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_GatewayInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_GatewayInstanceEntityPrototype() {

	var gatewayInstanceEntityProto EntityPrototype
	gatewayInstanceEntityProto.TypeName = "BPMS.GatewayInstance"
	gatewayInstanceEntityProto.SuperTypeNames = append(gatewayInstanceEntityProto.SuperTypeNames, "BPMS.Instance")
	gatewayInstanceEntityProto.SuperTypeNames = append(gatewayInstanceEntityProto.SuperTypeNames, "BPMS.FlowNodeInstance")
	gatewayInstanceEntityProto.Ids = append(gatewayInstanceEntityProto.Ids, "uuid")
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "uuid")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "xs.string")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 0)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, false)
	gatewayInstanceEntityProto.Indexs = append(gatewayInstanceEntityProto.Indexs, "parentUuid")
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "parentUuid")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "xs.string")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 1)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, false)

	/** members of Instance **/
	gatewayInstanceEntityProto.Ids = append(gatewayInstanceEntityProto.Ids, "M_id")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 2)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_id")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "xs.ID")
	gatewayInstanceEntityProto.Indexs = append(gatewayInstanceEntityProto.Indexs, "M_bpmnElementId")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 3)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_bpmnElementId")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "xs.string")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 4)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_participants")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "[]xs.string")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 5)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_dataRef")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 6)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_data")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 7)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_logInfoRef")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "[]BPMS.LogInfo:Ref")

	/** members of FlowNodeInstance **/
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 8)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_flowNodeType")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "enum:FlowNodeType_AbstractTask:FlowNodeType_ServiceTask:FlowNodeType_UserTask:FlowNodeType_ManualTask:FlowNodeType_BusinessRuleTask:FlowNodeType_ScriptTask:FlowNodeType_EmbeddedSubprocess:FlowNodeType_EventSubprocess:FlowNodeType_AdHocSubprocess:FlowNodeType_Transaction:FlowNodeType_CallActivity:FlowNodeType_ParallelGateway:FlowNodeType_ExclusiveGateway:FlowNodeType_InclusiveGateway:FlowNodeType_EventBasedGateway:FlowNodeType_ComplexGateway:FlowNodeType_StartEvent:FlowNodeType_IntermediateCatchEvent:FlowNodeType_BoundaryEvent:FlowNodeType_EndEvent:FlowNodeType_IntermediateThrowEvent")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 9)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_lifecycleState")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "enum:LifecycleState_Completed:LifecycleState_Compensated:LifecycleState_Failed:LifecycleState_Terminated:LifecycleState_Completing:LifecycleState_Compensating:LifecycleState_Failing:LifecycleState_Terminating:LifecycleState_Ready:LifecycleState_Active")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 10)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_inputRef")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject:Ref")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 11)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_outputRef")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject:Ref")

	/** members of GatewayInstance **/
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 12)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_gatewayType")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "enum:GatewayType_ParallelGateway:GatewayType_ExclusiveGateway:GatewayType_InclusiveGateway:GatewayType_EventBasedGateway:GatewayType_ComplexGateway")

	/** associations of GatewayInstance **/
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 13)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, false)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_SubprocessInstancePtr")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "BPMS.SubprocessInstance:Ref")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 14)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, false)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "M_processInstancePtr")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "BPMS.ProcessInstance:Ref")
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "childsUuid")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "[]xs.string")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 15)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, false)

	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields, "referenced")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType, "[]EntityRef")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder, 16)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&gatewayInstanceEntityProto)

}

/** Create **/
func (this *BPMS_GatewayInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.GatewayInstance"

	var query EntityQuery
	query.TypeName = "BPMS.GatewayInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of FlowNodeInstance **/
	query.Fields = append(query.Fields, "M_flowNodeType")
	query.Fields = append(query.Fields, "M_lifecycleState")
	query.Fields = append(query.Fields, "M_inputRef")
	query.Fields = append(query.Fields, "M_outputRef")

	/** members of GatewayInstance **/
	query.Fields = append(query.Fields, "M_gatewayType")

	/** associations of GatewayInstance **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var GatewayInstanceInfo []interface{}

	GatewayInstanceInfo = append(GatewayInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		GatewayInstanceInfo = append(GatewayInstanceInfo, this.parentPtr.GetUuid())
	} else {
		GatewayInstanceInfo = append(GatewayInstanceInfo, "")
	}

	/** members of Instance **/
	GatewayInstanceInfo = append(GatewayInstanceInfo, this.object.M_id)
	GatewayInstanceInfo = append(GatewayInstanceInfo, this.object.M_bpmnElementId)
	GatewayInstanceInfo = append(GatewayInstanceInfo, this.object.M_participants)

	/** Save dataRef type ItemAwareElementInstance **/
	dataRefStr, _ := json.Marshal(this.object.M_dataRef)
	GatewayInstanceInfo = append(GatewayInstanceInfo, string(dataRefStr))

	/** Save data type ItemAwareElementInstance **/
	dataIds := make([]string, 0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity := GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(this.object.M_data[i].UUID, this.object.M_data[i])
		dataIds = append(dataIds, dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data", dataEntity)
		if dataEntity.NeedSave() {
			dataEntity.SaveEntity()
		}
	}
	dataStr, _ := json.Marshal(dataIds)
	GatewayInstanceInfo = append(GatewayInstanceInfo, string(dataStr))

	/** Save logInfoRef type LogInfo **/
	logInfoRefStr, _ := json.Marshal(this.object.M_logInfoRef)
	GatewayInstanceInfo = append(GatewayInstanceInfo, string(logInfoRefStr))

	/** members of FlowNodeInstance **/

	/** Save flowNodeType type FlowNodeType **/
	if this.object.M_flowNodeType == BPMS.FlowNodeType_AbstractTask {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ServiceTask {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 1)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_UserTask {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 2)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ManualTask {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 3)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_BusinessRuleTask {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 4)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ScriptTask {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 5)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EmbeddedSubprocess {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 6)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EventSubprocess {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 7)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_AdHocSubprocess {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 8)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_Transaction {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 9)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_CallActivity {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 10)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ParallelGateway {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 11)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ExclusiveGateway {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 12)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_InclusiveGateway {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 13)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EventBasedGateway {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 14)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ComplexGateway {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 15)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_StartEvent {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 16)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_IntermediateCatchEvent {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 17)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_BoundaryEvent {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 18)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EndEvent {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 19)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_IntermediateThrowEvent {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 20)
	} else {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	}

	/** Save lifecycleState type LifecycleState **/
	if this.object.M_lifecycleState == BPMS.LifecycleState_Completed {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Compensated {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 1)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Failed {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 2)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Terminated {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 3)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Completing {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 4)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Compensating {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 5)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Failing {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 6)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Terminating {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 7)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Ready {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 8)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Active {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 9)
	} else {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	}

	/** Save inputRef type ConnectingObject **/
	inputRefStr, _ := json.Marshal(this.object.M_inputRef)
	GatewayInstanceInfo = append(GatewayInstanceInfo, string(inputRefStr))

	/** Save outputRef type ConnectingObject **/
	outputRefStr, _ := json.Marshal(this.object.M_outputRef)
	GatewayInstanceInfo = append(GatewayInstanceInfo, string(outputRefStr))

	/** members of GatewayInstance **/

	/** Save gatewayType type GatewayType **/
	if this.object.M_gatewayType == BPMS.GatewayType_ParallelGateway {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	} else if this.object.M_gatewayType == BPMS.GatewayType_ExclusiveGateway {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 1)
	} else if this.object.M_gatewayType == BPMS.GatewayType_InclusiveGateway {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 2)
	} else if this.object.M_gatewayType == BPMS.GatewayType_EventBasedGateway {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 3)
	} else if this.object.M_gatewayType == BPMS.GatewayType_ComplexGateway {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 4)
	} else {
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	}

	/** associations of GatewayInstance **/

	/** Save SubprocessInstance type SubprocessInstance **/
	GatewayInstanceInfo = append(GatewayInstanceInfo, this.object.M_SubprocessInstancePtr)

	/** Save processInstance type ProcessInstance **/
	GatewayInstanceInfo = append(GatewayInstanceInfo, this.object.M_processInstancePtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	GatewayInstanceInfo = append(GatewayInstanceInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	GatewayInstanceInfo = append(GatewayInstanceInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), GatewayInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), GatewayInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_GatewayInstanceEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_GatewayInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.GatewayInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of FlowNodeInstance **/
	query.Fields = append(query.Fields, "M_flowNodeType")
	query.Fields = append(query.Fields, "M_lifecycleState")
	query.Fields = append(query.Fields, "M_inputRef")
	query.Fields = append(query.Fields, "M_outputRef")

	/** members of GatewayInstance **/
	query.Fields = append(query.Fields, "M_gatewayType")

	/** associations of GatewayInstance **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of GatewayInstance...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.GatewayInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.GatewayInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** bpmnElementId **/
		if results[0][3] != nil {
			this.object.M_bpmnElementId = results[0][3].(string)
		}

		/** participants **/
		if results[0][4] != nil {
			this.object.M_participants = append(this.object.M_participants, results[0][4].([]string)...)
		}

		/** dataRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ItemAwareElementInstance"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef, ids[i])
					GetServer().GetEntityManager().appendReference("dataRef", this.object.UUID, id_)
				}
			}
		}

		/** data **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_ItemAwareElementInstanceEntity)
					} else {
						dataEntity = GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data", dataEntity)
				}
			}
		}

		/** logInfoRef **/
		if results[0][7] != nil {
			idsStr := results[0][7].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.LogInfo"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef, ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef", this.object.UUID, id_)
				}
			}
		}

		/** members of FlowNodeInstance **/

		/** flowNodeType **/
		if results[0][8] != nil {
			enumIndex := results[0][8].(int)
			if enumIndex == 0 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_AbstractTask
			} else if enumIndex == 1 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ServiceTask
			} else if enumIndex == 2 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_UserTask
			} else if enumIndex == 3 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ManualTask
			} else if enumIndex == 4 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_BusinessRuleTask
			} else if enumIndex == 5 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ScriptTask
			} else if enumIndex == 6 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EmbeddedSubprocess
			} else if enumIndex == 7 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EventSubprocess
			} else if enumIndex == 8 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_AdHocSubprocess
			} else if enumIndex == 9 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_Transaction
			} else if enumIndex == 10 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_CallActivity
			} else if enumIndex == 11 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ParallelGateway
			} else if enumIndex == 12 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ExclusiveGateway
			} else if enumIndex == 13 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_InclusiveGateway
			} else if enumIndex == 14 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EventBasedGateway
			} else if enumIndex == 15 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ComplexGateway
			} else if enumIndex == 16 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_StartEvent
			} else if enumIndex == 17 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_IntermediateCatchEvent
			} else if enumIndex == 18 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_BoundaryEvent
			} else if enumIndex == 19 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EndEvent
			} else if enumIndex == 20 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_IntermediateThrowEvent
			}
		}

		/** lifecycleState **/
		if results[0][9] != nil {
			enumIndex := results[0][9].(int)
			if enumIndex == 0 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Completed
			} else if enumIndex == 1 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Compensated
			} else if enumIndex == 2 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Failed
			} else if enumIndex == 3 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Terminated
			} else if enumIndex == 4 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Completing
			} else if enumIndex == 5 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Compensating
			} else if enumIndex == 6 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Failing
			} else if enumIndex == 7 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Terminating
			} else if enumIndex == 8 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Ready
			} else if enumIndex == 9 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Active
			}
		}

		/** inputRef **/
		if results[0][10] != nil {
			idsStr := results[0][10].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ConnectingObject"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_inputRef = append(this.object.M_inputRef, ids[i])
					GetServer().GetEntityManager().appendReference("inputRef", this.object.UUID, id_)
				}
			}
		}

		/** outputRef **/
		if results[0][11] != nil {
			idsStr := results[0][11].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ConnectingObject"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_outputRef = append(this.object.M_outputRef, ids[i])
					GetServer().GetEntityManager().appendReference("outputRef", this.object.UUID, id_)
				}
			}
		}

		/** members of GatewayInstance **/

		/** gatewayType **/
		if results[0][12] != nil {
			enumIndex := results[0][12].(int)
			if enumIndex == 0 {
				this.object.M_gatewayType = BPMS.GatewayType_ParallelGateway
			} else if enumIndex == 1 {
				this.object.M_gatewayType = BPMS.GatewayType_ExclusiveGateway
			} else if enumIndex == 2 {
				this.object.M_gatewayType = BPMS.GatewayType_InclusiveGateway
			} else if enumIndex == 3 {
				this.object.M_gatewayType = BPMS.GatewayType_EventBasedGateway
			} else if enumIndex == 4 {
				this.object.M_gatewayType = BPMS.GatewayType_ComplexGateway
			}
		}

		/** associations of GatewayInstance **/

		/** SubprocessInstancePtr **/
		if results[0][13] != nil {
			id := results[0][13].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.SubprocessInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_SubprocessInstancePtr = id
				GetServer().GetEntityManager().appendReference("SubprocessInstancePtr", this.object.UUID, id_)
			}
		}

		/** processInstancePtr **/
		if results[0][14] != nil {
			id := results[0][14].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.ProcessInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_processInstancePtr = id
				GetServer().GetEntityManager().appendReference("processInstancePtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][15].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][16].(string)
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
func (this *EntityManager) NewBPMSGatewayInstanceEntityFromObject(object *BPMS.GatewayInstance) *BPMS_GatewayInstanceEntity {
	return this.NewBPMSGatewayInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_GatewayInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSGatewayInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.GatewayInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_GatewayInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_GatewayInstanceEntity) AppendReference(reference Entity) {

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
//              			EventInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_EventInstanceEntity struct {
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
	object         *BPMS.EventInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMSEventInstanceEntity(objectId string, object interface{}) *BPMS_EventInstanceEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSEventInstanceExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.EventInstance).TYPENAME = "BPMS.EventInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.EventInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_EventInstanceEntity)
		}
	} else {
		uuidStr = "BPMS.EventInstance%" + Utility.RandomUUID()
	}
	entity := new(BPMS_EventInstanceEntity)
	if object == nil {
		entity.object = new(BPMS.EventInstance)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.EventInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.EventInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.EventInstance", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_EventInstanceEntity) GetTypeName() string {
	return "BPMS.EventInstance"
}
func (this *BPMS_EventInstanceEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_EventInstanceEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_EventInstanceEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_EventInstanceEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_EventInstanceEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_EventInstanceEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_EventInstanceEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_EventInstanceEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_EventInstanceEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_EventInstanceEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_EventInstanceEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_EventInstanceEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_EventInstanceEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_EventInstanceEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_EventInstanceEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_EventInstanceEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_EventInstanceEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_EventInstanceEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_EventInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_EventInstanceEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_EventInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_EventInstanceEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_EventInstanceEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.EventInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_EventInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_EventInstanceEntityPrototype() {

	var eventInstanceEntityProto EntityPrototype
	eventInstanceEntityProto.TypeName = "BPMS.EventInstance"
	eventInstanceEntityProto.SuperTypeNames = append(eventInstanceEntityProto.SuperTypeNames, "BPMS.Instance")
	eventInstanceEntityProto.SuperTypeNames = append(eventInstanceEntityProto.SuperTypeNames, "BPMS.FlowNodeInstance")
	eventInstanceEntityProto.Ids = append(eventInstanceEntityProto.Ids, "uuid")
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "uuid")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "xs.string")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 0)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, false)
	eventInstanceEntityProto.Indexs = append(eventInstanceEntityProto.Indexs, "parentUuid")
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "parentUuid")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "xs.string")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 1)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, false)

	/** members of Instance **/
	eventInstanceEntityProto.Ids = append(eventInstanceEntityProto.Ids, "M_id")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 2)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_id")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "xs.ID")
	eventInstanceEntityProto.Indexs = append(eventInstanceEntityProto.Indexs, "M_bpmnElementId")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 3)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_bpmnElementId")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "xs.string")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 4)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_participants")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "[]xs.string")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 5)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_dataRef")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 6)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_data")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 7)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_logInfoRef")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "[]BPMS.LogInfo:Ref")

	/** members of FlowNodeInstance **/
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 8)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_flowNodeType")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "enum:FlowNodeType_AbstractTask:FlowNodeType_ServiceTask:FlowNodeType_UserTask:FlowNodeType_ManualTask:FlowNodeType_BusinessRuleTask:FlowNodeType_ScriptTask:FlowNodeType_EmbeddedSubprocess:FlowNodeType_EventSubprocess:FlowNodeType_AdHocSubprocess:FlowNodeType_Transaction:FlowNodeType_CallActivity:FlowNodeType_ParallelGateway:FlowNodeType_ExclusiveGateway:FlowNodeType_InclusiveGateway:FlowNodeType_EventBasedGateway:FlowNodeType_ComplexGateway:FlowNodeType_StartEvent:FlowNodeType_IntermediateCatchEvent:FlowNodeType_BoundaryEvent:FlowNodeType_EndEvent:FlowNodeType_IntermediateThrowEvent")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 9)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_lifecycleState")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "enum:LifecycleState_Completed:LifecycleState_Compensated:LifecycleState_Failed:LifecycleState_Terminated:LifecycleState_Completing:LifecycleState_Compensating:LifecycleState_Failing:LifecycleState_Terminating:LifecycleState_Ready:LifecycleState_Active")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 10)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_inputRef")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject:Ref")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 11)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_outputRef")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject:Ref")

	/** members of EventInstance **/
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 12)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_eventType")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "enum:EventType_StartEvent:EventType_IntermediateCatchEvent:EventType_BoundaryEvent:EventType_EndEvent:EventType_IntermediateThrowEvent")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 13)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_eventDefintionInstances")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "[]BPMS.EventDefinitionInstance")

	/** associations of EventInstance **/
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 14)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, false)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_SubprocessInstancePtr")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "BPMS.SubprocessInstance:Ref")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 15)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, false)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "M_processInstancePtr")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "BPMS.ProcessInstance:Ref")
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "childsUuid")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "[]xs.string")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 16)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, false)

	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields, "referenced")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType, "[]EntityRef")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder, 17)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&eventInstanceEntityProto)

}

/** Create **/
func (this *BPMS_EventInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.EventInstance"

	var query EntityQuery
	query.TypeName = "BPMS.EventInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of FlowNodeInstance **/
	query.Fields = append(query.Fields, "M_flowNodeType")
	query.Fields = append(query.Fields, "M_lifecycleState")
	query.Fields = append(query.Fields, "M_inputRef")
	query.Fields = append(query.Fields, "M_outputRef")

	/** members of EventInstance **/
	query.Fields = append(query.Fields, "M_eventType")
	query.Fields = append(query.Fields, "M_eventDefintionInstances")

	/** associations of EventInstance **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var EventInstanceInfo []interface{}

	EventInstanceInfo = append(EventInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		EventInstanceInfo = append(EventInstanceInfo, this.parentPtr.GetUuid())
	} else {
		EventInstanceInfo = append(EventInstanceInfo, "")
	}

	/** members of Instance **/
	EventInstanceInfo = append(EventInstanceInfo, this.object.M_id)
	EventInstanceInfo = append(EventInstanceInfo, this.object.M_bpmnElementId)
	EventInstanceInfo = append(EventInstanceInfo, this.object.M_participants)

	/** Save dataRef type ItemAwareElementInstance **/
	dataRefStr, _ := json.Marshal(this.object.M_dataRef)
	EventInstanceInfo = append(EventInstanceInfo, string(dataRefStr))

	/** Save data type ItemAwareElementInstance **/
	dataIds := make([]string, 0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity := GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(this.object.M_data[i].UUID, this.object.M_data[i])
		dataIds = append(dataIds, dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data", dataEntity)
		if dataEntity.NeedSave() {
			dataEntity.SaveEntity()
		}
	}
	dataStr, _ := json.Marshal(dataIds)
	EventInstanceInfo = append(EventInstanceInfo, string(dataStr))

	/** Save logInfoRef type LogInfo **/
	logInfoRefStr, _ := json.Marshal(this.object.M_logInfoRef)
	EventInstanceInfo = append(EventInstanceInfo, string(logInfoRefStr))

	/** members of FlowNodeInstance **/

	/** Save flowNodeType type FlowNodeType **/
	if this.object.M_flowNodeType == BPMS.FlowNodeType_AbstractTask {
		EventInstanceInfo = append(EventInstanceInfo, 0)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ServiceTask {
		EventInstanceInfo = append(EventInstanceInfo, 1)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_UserTask {
		EventInstanceInfo = append(EventInstanceInfo, 2)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ManualTask {
		EventInstanceInfo = append(EventInstanceInfo, 3)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_BusinessRuleTask {
		EventInstanceInfo = append(EventInstanceInfo, 4)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ScriptTask {
		EventInstanceInfo = append(EventInstanceInfo, 5)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EmbeddedSubprocess {
		EventInstanceInfo = append(EventInstanceInfo, 6)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EventSubprocess {
		EventInstanceInfo = append(EventInstanceInfo, 7)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_AdHocSubprocess {
		EventInstanceInfo = append(EventInstanceInfo, 8)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_Transaction {
		EventInstanceInfo = append(EventInstanceInfo, 9)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_CallActivity {
		EventInstanceInfo = append(EventInstanceInfo, 10)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ParallelGateway {
		EventInstanceInfo = append(EventInstanceInfo, 11)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ExclusiveGateway {
		EventInstanceInfo = append(EventInstanceInfo, 12)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_InclusiveGateway {
		EventInstanceInfo = append(EventInstanceInfo, 13)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EventBasedGateway {
		EventInstanceInfo = append(EventInstanceInfo, 14)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_ComplexGateway {
		EventInstanceInfo = append(EventInstanceInfo, 15)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_StartEvent {
		EventInstanceInfo = append(EventInstanceInfo, 16)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_IntermediateCatchEvent {
		EventInstanceInfo = append(EventInstanceInfo, 17)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_BoundaryEvent {
		EventInstanceInfo = append(EventInstanceInfo, 18)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_EndEvent {
		EventInstanceInfo = append(EventInstanceInfo, 19)
	} else if this.object.M_flowNodeType == BPMS.FlowNodeType_IntermediateThrowEvent {
		EventInstanceInfo = append(EventInstanceInfo, 20)
	} else {
		EventInstanceInfo = append(EventInstanceInfo, 0)
	}

	/** Save lifecycleState type LifecycleState **/
	if this.object.M_lifecycleState == BPMS.LifecycleState_Completed {
		EventInstanceInfo = append(EventInstanceInfo, 0)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Compensated {
		EventInstanceInfo = append(EventInstanceInfo, 1)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Failed {
		EventInstanceInfo = append(EventInstanceInfo, 2)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Terminated {
		EventInstanceInfo = append(EventInstanceInfo, 3)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Completing {
		EventInstanceInfo = append(EventInstanceInfo, 4)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Compensating {
		EventInstanceInfo = append(EventInstanceInfo, 5)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Failing {
		EventInstanceInfo = append(EventInstanceInfo, 6)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Terminating {
		EventInstanceInfo = append(EventInstanceInfo, 7)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Ready {
		EventInstanceInfo = append(EventInstanceInfo, 8)
	} else if this.object.M_lifecycleState == BPMS.LifecycleState_Active {
		EventInstanceInfo = append(EventInstanceInfo, 9)
	} else {
		EventInstanceInfo = append(EventInstanceInfo, 0)
	}

	/** Save inputRef type ConnectingObject **/
	inputRefStr, _ := json.Marshal(this.object.M_inputRef)
	EventInstanceInfo = append(EventInstanceInfo, string(inputRefStr))

	/** Save outputRef type ConnectingObject **/
	outputRefStr, _ := json.Marshal(this.object.M_outputRef)
	EventInstanceInfo = append(EventInstanceInfo, string(outputRefStr))

	/** members of EventInstance **/

	/** Save eventType type EventType **/
	if this.object.M_eventType == BPMS.EventType_StartEvent {
		EventInstanceInfo = append(EventInstanceInfo, 0)
	} else if this.object.M_eventType == BPMS.EventType_IntermediateCatchEvent {
		EventInstanceInfo = append(EventInstanceInfo, 1)
	} else if this.object.M_eventType == BPMS.EventType_BoundaryEvent {
		EventInstanceInfo = append(EventInstanceInfo, 2)
	} else if this.object.M_eventType == BPMS.EventType_EndEvent {
		EventInstanceInfo = append(EventInstanceInfo, 3)
	} else if this.object.M_eventType == BPMS.EventType_IntermediateThrowEvent {
		EventInstanceInfo = append(EventInstanceInfo, 4)
	} else {
		EventInstanceInfo = append(EventInstanceInfo, 0)
	}

	/** Save eventDefintionInstances type EventDefinitionInstance **/
	eventDefintionInstancesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_eventDefintionInstances); i++ {
		eventDefintionInstancesEntity := GetServer().GetEntityManager().NewBPMSEventDefinitionInstanceEntity(this.object.M_eventDefintionInstances[i].UUID, this.object.M_eventDefintionInstances[i])
		eventDefintionInstancesIds = append(eventDefintionInstancesIds, eventDefintionInstancesEntity.uuid)
		eventDefintionInstancesEntity.AppendReferenced("eventDefintionInstances", this)
		this.AppendChild("eventDefintionInstances", eventDefintionInstancesEntity)
		if eventDefintionInstancesEntity.NeedSave() {
			eventDefintionInstancesEntity.SaveEntity()
		}
	}
	eventDefintionInstancesStr, _ := json.Marshal(eventDefintionInstancesIds)
	EventInstanceInfo = append(EventInstanceInfo, string(eventDefintionInstancesStr))

	/** associations of EventInstance **/

	/** Save SubprocessInstance type SubprocessInstance **/
	EventInstanceInfo = append(EventInstanceInfo, this.object.M_SubprocessInstancePtr)

	/** Save processInstance type ProcessInstance **/
	EventInstanceInfo = append(EventInstanceInfo, this.object.M_processInstancePtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	EventInstanceInfo = append(EventInstanceInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	EventInstanceInfo = append(EventInstanceInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), EventInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), EventInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_EventInstanceEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_EventInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.EventInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of FlowNodeInstance **/
	query.Fields = append(query.Fields, "M_flowNodeType")
	query.Fields = append(query.Fields, "M_lifecycleState")
	query.Fields = append(query.Fields, "M_inputRef")
	query.Fields = append(query.Fields, "M_outputRef")

	/** members of EventInstance **/
	query.Fields = append(query.Fields, "M_eventType")
	query.Fields = append(query.Fields, "M_eventDefintionInstances")

	/** associations of EventInstance **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of EventInstance...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.EventInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.EventInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** bpmnElementId **/
		if results[0][3] != nil {
			this.object.M_bpmnElementId = results[0][3].(string)
		}

		/** participants **/
		if results[0][4] != nil {
			this.object.M_participants = append(this.object.M_participants, results[0][4].([]string)...)
		}

		/** dataRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ItemAwareElementInstance"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef, ids[i])
					GetServer().GetEntityManager().appendReference("dataRef", this.object.UUID, id_)
				}
			}
		}

		/** data **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_ItemAwareElementInstanceEntity)
					} else {
						dataEntity = GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data", dataEntity)
				}
			}
		}

		/** logInfoRef **/
		if results[0][7] != nil {
			idsStr := results[0][7].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.LogInfo"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef, ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef", this.object.UUID, id_)
				}
			}
		}

		/** members of FlowNodeInstance **/

		/** flowNodeType **/
		if results[0][8] != nil {
			enumIndex := results[0][8].(int)
			if enumIndex == 0 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_AbstractTask
			} else if enumIndex == 1 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ServiceTask
			} else if enumIndex == 2 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_UserTask
			} else if enumIndex == 3 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ManualTask
			} else if enumIndex == 4 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_BusinessRuleTask
			} else if enumIndex == 5 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ScriptTask
			} else if enumIndex == 6 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EmbeddedSubprocess
			} else if enumIndex == 7 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EventSubprocess
			} else if enumIndex == 8 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_AdHocSubprocess
			} else if enumIndex == 9 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_Transaction
			} else if enumIndex == 10 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_CallActivity
			} else if enumIndex == 11 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ParallelGateway
			} else if enumIndex == 12 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ExclusiveGateway
			} else if enumIndex == 13 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_InclusiveGateway
			} else if enumIndex == 14 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EventBasedGateway
			} else if enumIndex == 15 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_ComplexGateway
			} else if enumIndex == 16 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_StartEvent
			} else if enumIndex == 17 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_IntermediateCatchEvent
			} else if enumIndex == 18 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_BoundaryEvent
			} else if enumIndex == 19 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_EndEvent
			} else if enumIndex == 20 {
				this.object.M_flowNodeType = BPMS.FlowNodeType_IntermediateThrowEvent
			}
		}

		/** lifecycleState **/
		if results[0][9] != nil {
			enumIndex := results[0][9].(int)
			if enumIndex == 0 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Completed
			} else if enumIndex == 1 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Compensated
			} else if enumIndex == 2 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Failed
			} else if enumIndex == 3 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Terminated
			} else if enumIndex == 4 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Completing
			} else if enumIndex == 5 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Compensating
			} else if enumIndex == 6 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Failing
			} else if enumIndex == 7 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Terminating
			} else if enumIndex == 8 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Ready
			} else if enumIndex == 9 {
				this.object.M_lifecycleState = BPMS.LifecycleState_Active
			}
		}

		/** inputRef **/
		if results[0][10] != nil {
			idsStr := results[0][10].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ConnectingObject"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_inputRef = append(this.object.M_inputRef, ids[i])
					GetServer().GetEntityManager().appendReference("inputRef", this.object.UUID, id_)
				}
			}
		}

		/** outputRef **/
		if results[0][11] != nil {
			idsStr := results[0][11].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ConnectingObject"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_outputRef = append(this.object.M_outputRef, ids[i])
					GetServer().GetEntityManager().appendReference("outputRef", this.object.UUID, id_)
				}
			}
		}

		/** members of EventInstance **/

		/** eventType **/
		if results[0][12] != nil {
			enumIndex := results[0][12].(int)
			if enumIndex == 0 {
				this.object.M_eventType = BPMS.EventType_StartEvent
			} else if enumIndex == 1 {
				this.object.M_eventType = BPMS.EventType_IntermediateCatchEvent
			} else if enumIndex == 2 {
				this.object.M_eventType = BPMS.EventType_BoundaryEvent
			} else if enumIndex == 3 {
				this.object.M_eventType = BPMS.EventType_EndEvent
			} else if enumIndex == 4 {
				this.object.M_eventType = BPMS.EventType_IntermediateThrowEvent
			}
		}

		/** eventDefintionInstances **/
		if results[0][13] != nil {
			uuidsStr := results[0][13].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var eventDefintionInstancesEntity *BPMS_EventDefinitionInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						eventDefintionInstancesEntity = instance.(*BPMS_EventDefinitionInstanceEntity)
					} else {
						eventDefintionInstancesEntity = GetServer().GetEntityManager().NewBPMSEventDefinitionInstanceEntity(uuids[i], nil)
						eventDefintionInstancesEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(eventDefintionInstancesEntity)
					}
					eventDefintionInstancesEntity.AppendReferenced("eventDefintionInstances", this)
					this.AppendChild("eventDefintionInstances", eventDefintionInstancesEntity)
				}
			}
		}

		/** associations of EventInstance **/

		/** SubprocessInstancePtr **/
		if results[0][14] != nil {
			id := results[0][14].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.SubprocessInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_SubprocessInstancePtr = id
				GetServer().GetEntityManager().appendReference("SubprocessInstancePtr", this.object.UUID, id_)
			}
		}

		/** processInstancePtr **/
		if results[0][15] != nil {
			id := results[0][15].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.ProcessInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_processInstancePtr = id
				GetServer().GetEntityManager().appendReference("processInstancePtr", this.object.UUID, id_)
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
func (this *EntityManager) NewBPMSEventInstanceEntityFromObject(object *BPMS.EventInstance) *BPMS_EventInstanceEntity {
	return this.NewBPMSEventInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_EventInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSEventInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.EventInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_EventInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_EventInstanceEntity) AppendReference(reference Entity) {

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
//              			DefinitionsInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_DefinitionsInstanceEntity struct {
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
	object         *BPMS.DefinitionsInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMSDefinitionsInstanceEntity(objectId string, object interface{}) *BPMS_DefinitionsInstanceEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSDefinitionsInstanceExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.DefinitionsInstance).TYPENAME = "BPMS.DefinitionsInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.DefinitionsInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_DefinitionsInstanceEntity)
		}
	} else {
		uuidStr = "BPMS.DefinitionsInstance%" + Utility.RandomUUID()
	}
	entity := new(BPMS_DefinitionsInstanceEntity)
	if object == nil {
		entity.object = new(BPMS.DefinitionsInstance)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.DefinitionsInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.DefinitionsInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.DefinitionsInstance", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_DefinitionsInstanceEntity) GetTypeName() string {
	return "BPMS.DefinitionsInstance"
}
func (this *BPMS_DefinitionsInstanceEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_DefinitionsInstanceEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_DefinitionsInstanceEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_DefinitionsInstanceEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_DefinitionsInstanceEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_DefinitionsInstanceEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_DefinitionsInstanceEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_DefinitionsInstanceEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_DefinitionsInstanceEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_DefinitionsInstanceEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_DefinitionsInstanceEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_DefinitionsInstanceEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_DefinitionsInstanceEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_DefinitionsInstanceEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_DefinitionsInstanceEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_DefinitionsInstanceEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_DefinitionsInstanceEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_DefinitionsInstanceEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_DefinitionsInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_DefinitionsInstanceEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_DefinitionsInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_DefinitionsInstanceEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_DefinitionsInstanceEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.DefinitionsInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_DefinitionsInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_DefinitionsInstanceEntityPrototype() {

	var definitionsInstanceEntityProto EntityPrototype
	definitionsInstanceEntityProto.TypeName = "BPMS.DefinitionsInstance"
	definitionsInstanceEntityProto.SuperTypeNames = append(definitionsInstanceEntityProto.SuperTypeNames, "BPMS.Instance")
	definitionsInstanceEntityProto.Ids = append(definitionsInstanceEntityProto.Ids, "uuid")
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "uuid")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "xs.string")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 0)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, false)
	definitionsInstanceEntityProto.Indexs = append(definitionsInstanceEntityProto.Indexs, "parentUuid")
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "parentUuid")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "xs.string")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 1)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, false)

	/** members of Instance **/
	definitionsInstanceEntityProto.Ids = append(definitionsInstanceEntityProto.Ids, "M_id")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 2)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "M_id")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "xs.ID")
	definitionsInstanceEntityProto.Indexs = append(definitionsInstanceEntityProto.Indexs, "M_bpmnElementId")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 3)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "M_bpmnElementId")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "xs.string")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 4)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "M_participants")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "[]xs.string")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 5)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "M_dataRef")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 6)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "M_data")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 7)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "M_logInfoRef")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "[]BPMS.LogInfo:Ref")

	/** members of DefinitionsInstance **/
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 8)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "M_processInstances")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "[]BPMS.ProcessInstance")

	/** associations of DefinitionsInstance **/
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 9)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, false)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "M_parentPtr")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "BPMS.Runtimes:Ref")
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "childsUuid")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "[]xs.string")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 10)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, false)

	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields, "referenced")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType, "[]EntityRef")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder, 11)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&definitionsInstanceEntityProto)

}

/** Create **/
func (this *BPMS_DefinitionsInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.DefinitionsInstance"

	var query EntityQuery
	query.TypeName = "BPMS.DefinitionsInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of DefinitionsInstance **/
	query.Fields = append(query.Fields, "M_processInstances")

	/** associations of DefinitionsInstance **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var DefinitionsInstanceInfo []interface{}

	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, this.parentPtr.GetUuid())
	} else {
		DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, "")
	}

	/** members of Instance **/
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, this.object.M_id)
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, this.object.M_bpmnElementId)
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, this.object.M_participants)

	/** Save dataRef type ItemAwareElementInstance **/
	dataRefStr, _ := json.Marshal(this.object.M_dataRef)
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, string(dataRefStr))

	/** Save data type ItemAwareElementInstance **/
	dataIds := make([]string, 0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity := GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(this.object.M_data[i].UUID, this.object.M_data[i])
		dataIds = append(dataIds, dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data", dataEntity)
		if dataEntity.NeedSave() {
			dataEntity.SaveEntity()
		}
	}
	dataStr, _ := json.Marshal(dataIds)
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, string(dataStr))

	/** Save logInfoRef type LogInfo **/
	logInfoRefStr, _ := json.Marshal(this.object.M_logInfoRef)
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, string(logInfoRefStr))

	/** members of DefinitionsInstance **/

	/** Save processInstances type ProcessInstance **/
	processInstancesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_processInstances); i++ {
		processInstancesEntity := GetServer().GetEntityManager().NewBPMSProcessInstanceEntity(this.object.M_processInstances[i].UUID, this.object.M_processInstances[i])
		processInstancesIds = append(processInstancesIds, processInstancesEntity.uuid)
		processInstancesEntity.AppendReferenced("processInstances", this)
		this.AppendChild("processInstances", processInstancesEntity)
		if processInstancesEntity.NeedSave() {
			processInstancesEntity.SaveEntity()
		}
	}
	processInstancesStr, _ := json.Marshal(processInstancesIds)
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, string(processInstancesStr))

	/** associations of DefinitionsInstance **/

	/** Save parent type Runtimes **/
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, this.object.M_parentPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), DefinitionsInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), DefinitionsInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_DefinitionsInstanceEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_DefinitionsInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.DefinitionsInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of DefinitionsInstance **/
	query.Fields = append(query.Fields, "M_processInstances")

	/** associations of DefinitionsInstance **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of DefinitionsInstance...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.DefinitionsInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.DefinitionsInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** bpmnElementId **/
		if results[0][3] != nil {
			this.object.M_bpmnElementId = results[0][3].(string)
		}

		/** participants **/
		if results[0][4] != nil {
			this.object.M_participants = append(this.object.M_participants, results[0][4].([]string)...)
		}

		/** dataRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ItemAwareElementInstance"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef, ids[i])
					GetServer().GetEntityManager().appendReference("dataRef", this.object.UUID, id_)
				}
			}
		}

		/** data **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_ItemAwareElementInstanceEntity)
					} else {
						dataEntity = GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data", dataEntity)
				}
			}
		}

		/** logInfoRef **/
		if results[0][7] != nil {
			idsStr := results[0][7].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.LogInfo"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef, ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef", this.object.UUID, id_)
				}
			}
		}

		/** members of DefinitionsInstance **/

		/** processInstances **/
		if results[0][8] != nil {
			uuidsStr := results[0][8].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var processInstancesEntity *BPMS_ProcessInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						processInstancesEntity = instance.(*BPMS_ProcessInstanceEntity)
					} else {
						processInstancesEntity = GetServer().GetEntityManager().NewBPMSProcessInstanceEntity(uuids[i], nil)
						processInstancesEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(processInstancesEntity)
					}
					processInstancesEntity.AppendReferenced("processInstances", this)
					this.AppendChild("processInstances", processInstancesEntity)
				}
			}
		}

		/** associations of DefinitionsInstance **/

		/** parentPtr **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.Runtimes"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][10].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][11].(string)
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
func (this *EntityManager) NewBPMSDefinitionsInstanceEntityFromObject(object *BPMS.DefinitionsInstance) *BPMS_DefinitionsInstanceEntity {
	return this.NewBPMSDefinitionsInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_DefinitionsInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSDefinitionsInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.DefinitionsInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_DefinitionsInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_DefinitionsInstanceEntity) AppendReference(reference Entity) {

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
//              			ProcessInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_ProcessInstanceEntity struct {
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
	object         *BPMS.ProcessInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMSProcessInstanceEntity(objectId string, object interface{}) *BPMS_ProcessInstanceEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSProcessInstanceExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.ProcessInstance).TYPENAME = "BPMS.ProcessInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.ProcessInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_ProcessInstanceEntity)
		}
	} else {
		uuidStr = "BPMS.ProcessInstance%" + Utility.RandomUUID()
	}
	entity := new(BPMS_ProcessInstanceEntity)
	if object == nil {
		entity.object = new(BPMS.ProcessInstance)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.ProcessInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.ProcessInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.ProcessInstance", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_ProcessInstanceEntity) GetTypeName() string {
	return "BPMS.ProcessInstance"
}
func (this *BPMS_ProcessInstanceEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_ProcessInstanceEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_ProcessInstanceEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_ProcessInstanceEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_ProcessInstanceEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_ProcessInstanceEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_ProcessInstanceEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_ProcessInstanceEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_ProcessInstanceEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_ProcessInstanceEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_ProcessInstanceEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_ProcessInstanceEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_ProcessInstanceEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_ProcessInstanceEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_ProcessInstanceEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_ProcessInstanceEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_ProcessInstanceEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_ProcessInstanceEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_ProcessInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_ProcessInstanceEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_ProcessInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_ProcessInstanceEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_ProcessInstanceEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.ProcessInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_ProcessInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_ProcessInstanceEntityPrototype() {

	var processInstanceEntityProto EntityPrototype
	processInstanceEntityProto.TypeName = "BPMS.ProcessInstance"
	processInstanceEntityProto.SuperTypeNames = append(processInstanceEntityProto.SuperTypeNames, "BPMS.Instance")
	processInstanceEntityProto.Ids = append(processInstanceEntityProto.Ids, "uuid")
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "uuid")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 0)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, false)
	processInstanceEntityProto.Indexs = append(processInstanceEntityProto.Indexs, "parentUuid")
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "parentUuid")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 1)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, false)

	/** members of Instance **/
	processInstanceEntityProto.Ids = append(processInstanceEntityProto.Ids, "M_id")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 2)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_id")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "xs.ID")
	processInstanceEntityProto.Indexs = append(processInstanceEntityProto.Indexs, "M_bpmnElementId")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 3)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_bpmnElementId")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 4)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_participants")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "[]xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 5)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_dataRef")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 6)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_data")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 7)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_logInfoRef")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "[]BPMS.LogInfo:Ref")

	/** members of ProcessInstance **/
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 8)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_number")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "xs.int")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 9)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_colorName")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 10)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_colorNumber")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 11)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_flowNodeInstances")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "[]BPMS.FlowNodeInstance")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 12)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_connectingObjects")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "[]BPMS.ConnectingObject")

	/** associations of ProcessInstance **/
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 13)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, false)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "M_parentPtr")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "BPMS.DefinitionsInstance:Ref")
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "childsUuid")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "[]xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 14)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, false)

	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields, "referenced")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType, "[]EntityRef")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder, 15)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&processInstanceEntityProto)

}

/** Create **/
func (this *BPMS_ProcessInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.ProcessInstance"

	var query EntityQuery
	query.TypeName = "BPMS.ProcessInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of ProcessInstance **/
	query.Fields = append(query.Fields, "M_number")
	query.Fields = append(query.Fields, "M_colorName")
	query.Fields = append(query.Fields, "M_colorNumber")
	query.Fields = append(query.Fields, "M_flowNodeInstances")
	query.Fields = append(query.Fields, "M_connectingObjects")

	/** associations of ProcessInstance **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ProcessInstanceInfo []interface{}

	ProcessInstanceInfo = append(ProcessInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		ProcessInstanceInfo = append(ProcessInstanceInfo, this.parentPtr.GetUuid())
	} else {
		ProcessInstanceInfo = append(ProcessInstanceInfo, "")
	}

	/** members of Instance **/
	ProcessInstanceInfo = append(ProcessInstanceInfo, this.object.M_id)
	ProcessInstanceInfo = append(ProcessInstanceInfo, this.object.M_bpmnElementId)
	ProcessInstanceInfo = append(ProcessInstanceInfo, this.object.M_participants)

	/** Save dataRef type ItemAwareElementInstance **/
	dataRefStr, _ := json.Marshal(this.object.M_dataRef)
	ProcessInstanceInfo = append(ProcessInstanceInfo, string(dataRefStr))

	/** Save data type ItemAwareElementInstance **/
	dataIds := make([]string, 0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity := GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(this.object.M_data[i].UUID, this.object.M_data[i])
		dataIds = append(dataIds, dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data", dataEntity)
		if dataEntity.NeedSave() {
			dataEntity.SaveEntity()
		}
	}
	dataStr, _ := json.Marshal(dataIds)
	ProcessInstanceInfo = append(ProcessInstanceInfo, string(dataStr))

	/** Save logInfoRef type LogInfo **/
	logInfoRefStr, _ := json.Marshal(this.object.M_logInfoRef)
	ProcessInstanceInfo = append(ProcessInstanceInfo, string(logInfoRefStr))

	/** members of ProcessInstance **/
	ProcessInstanceInfo = append(ProcessInstanceInfo, this.object.M_number)
	ProcessInstanceInfo = append(ProcessInstanceInfo, this.object.M_colorName)
	ProcessInstanceInfo = append(ProcessInstanceInfo, this.object.M_colorNumber)

	/** Save flowNodeInstances type FlowNodeInstance **/
	flowNodeInstancesIds := make([]string, 0)
	for i := 0; i < len(this.object.M_flowNodeInstances); i++ {
		switch v := this.object.M_flowNodeInstances[i].(type) {
		case *BPMS.ActivityInstance:
			flowNodeInstancesEntity := GetServer().GetEntityManager().NewBPMSActivityInstanceEntity(v.UUID, v)
			flowNodeInstancesIds = append(flowNodeInstancesIds, flowNodeInstancesEntity.uuid)
			flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
			this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
			if flowNodeInstancesEntity.NeedSave() {
				flowNodeInstancesEntity.SaveEntity()
			}
		case *BPMS.SubprocessInstance:
			flowNodeInstancesEntity := GetServer().GetEntityManager().NewBPMSSubprocessInstanceEntity(v.UUID, v)
			flowNodeInstancesIds = append(flowNodeInstancesIds, flowNodeInstancesEntity.uuid)
			flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
			this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
			if flowNodeInstancesEntity.NeedSave() {
				flowNodeInstancesEntity.SaveEntity()
			}
		case *BPMS.GatewayInstance:
			flowNodeInstancesEntity := GetServer().GetEntityManager().NewBPMSGatewayInstanceEntity(v.UUID, v)
			flowNodeInstancesIds = append(flowNodeInstancesIds, flowNodeInstancesEntity.uuid)
			flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
			this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
			if flowNodeInstancesEntity.NeedSave() {
				flowNodeInstancesEntity.SaveEntity()
			}
		case *BPMS.EventInstance:
			flowNodeInstancesEntity := GetServer().GetEntityManager().NewBPMSEventInstanceEntity(v.UUID, v)
			flowNodeInstancesIds = append(flowNodeInstancesIds, flowNodeInstancesEntity.uuid)
			flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
			this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
			if flowNodeInstancesEntity.NeedSave() {
				flowNodeInstancesEntity.SaveEntity()
			}
		}
	}
	flowNodeInstancesStr, _ := json.Marshal(flowNodeInstancesIds)
	ProcessInstanceInfo = append(ProcessInstanceInfo, string(flowNodeInstancesStr))

	/** Save connectingObjects type ConnectingObject **/
	connectingObjectsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_connectingObjects); i++ {
		connectingObjectsEntity := GetServer().GetEntityManager().NewBPMSConnectingObjectEntity(this.object.M_connectingObjects[i].UUID, this.object.M_connectingObjects[i])
		connectingObjectsIds = append(connectingObjectsIds, connectingObjectsEntity.uuid)
		connectingObjectsEntity.AppendReferenced("connectingObjects", this)
		this.AppendChild("connectingObjects", connectingObjectsEntity)
		if connectingObjectsEntity.NeedSave() {
			connectingObjectsEntity.SaveEntity()
		}
	}
	connectingObjectsStr, _ := json.Marshal(connectingObjectsIds)
	ProcessInstanceInfo = append(ProcessInstanceInfo, string(connectingObjectsStr))

	/** associations of ProcessInstance **/

	/** Save parent type DefinitionsInstance **/
	ProcessInstanceInfo = append(ProcessInstanceInfo, this.object.M_parentPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ProcessInstanceInfo = append(ProcessInstanceInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ProcessInstanceInfo = append(ProcessInstanceInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), ProcessInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), ProcessInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_ProcessInstanceEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_ProcessInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.ProcessInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of ProcessInstance **/
	query.Fields = append(query.Fields, "M_number")
	query.Fields = append(query.Fields, "M_colorName")
	query.Fields = append(query.Fields, "M_colorNumber")
	query.Fields = append(query.Fields, "M_flowNodeInstances")
	query.Fields = append(query.Fields, "M_connectingObjects")

	/** associations of ProcessInstance **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ProcessInstance...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.ProcessInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.ProcessInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** bpmnElementId **/
		if results[0][3] != nil {
			this.object.M_bpmnElementId = results[0][3].(string)
		}

		/** participants **/
		if results[0][4] != nil {
			this.object.M_participants = append(this.object.M_participants, results[0][4].([]string)...)
		}

		/** dataRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ItemAwareElementInstance"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef, ids[i])
					GetServer().GetEntityManager().appendReference("dataRef", this.object.UUID, id_)
				}
			}
		}

		/** data **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_ItemAwareElementInstanceEntity)
					} else {
						dataEntity = GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data", dataEntity)
				}
			}
		}

		/** logInfoRef **/
		if results[0][7] != nil {
			idsStr := results[0][7].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.LogInfo"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef, ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef", this.object.UUID, id_)
				}
			}
		}

		/** members of ProcessInstance **/

		/** number **/
		if results[0][8] != nil {
			this.object.M_number = results[0][8].(int)
		}

		/** colorName **/
		if results[0][9] != nil {
			this.object.M_colorName = results[0][9].(string)
		}

		/** colorNumber **/
		if results[0][10] != nil {
			this.object.M_colorNumber = results[0][10].(string)
		}

		/** flowNodeInstances **/
		if results[0][11] != nil {
			uuidsStr := results[0][11].(string)
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
				if typeName == "BPMS.ActivityInstance" {
					if len(uuids[i]) > 0 {
						var flowNodeInstancesEntity *BPMS_ActivityInstanceEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							flowNodeInstancesEntity = instance.(*BPMS_ActivityInstanceEntity)
						} else {
							flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMSActivityInstanceEntity(uuids[i], nil)
							flowNodeInstancesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
						}
						flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
						this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
					}
				} else if typeName == "BPMS.SubprocessInstance" {
					if len(uuids[i]) > 0 {
						var flowNodeInstancesEntity *BPMS_SubprocessInstanceEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							flowNodeInstancesEntity = instance.(*BPMS_SubprocessInstanceEntity)
						} else {
							flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMSSubprocessInstanceEntity(uuids[i], nil)
							flowNodeInstancesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
						}
						flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
						this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
					}
				} else if typeName == "BPMS.GatewayInstance" {
					if len(uuids[i]) > 0 {
						var flowNodeInstancesEntity *BPMS_GatewayInstanceEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							flowNodeInstancesEntity = instance.(*BPMS_GatewayInstanceEntity)
						} else {
							flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMSGatewayInstanceEntity(uuids[i], nil)
							flowNodeInstancesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
						}
						flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
						this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
					}
				} else if typeName == "BPMS.EventInstance" {
					if len(uuids[i]) > 0 {
						var flowNodeInstancesEntity *BPMS_EventInstanceEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							flowNodeInstancesEntity = instance.(*BPMS_EventInstanceEntity)
						} else {
							flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMSEventInstanceEntity(uuids[i], nil)
							flowNodeInstancesEntity.InitEntity(uuids[i])
							GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
						}
						flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
						this.AppendChild("flowNodeInstances", flowNodeInstancesEntity)
					}
				}
			}
		}

		/** connectingObjects **/
		if results[0][12] != nil {
			uuidsStr := results[0][12].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var connectingObjectsEntity *BPMS_ConnectingObjectEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						connectingObjectsEntity = instance.(*BPMS_ConnectingObjectEntity)
					} else {
						connectingObjectsEntity = GetServer().GetEntityManager().NewBPMSConnectingObjectEntity(uuids[i], nil)
						connectingObjectsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(connectingObjectsEntity)
					}
					connectingObjectsEntity.AppendReferenced("connectingObjects", this)
					this.AppendChild("connectingObjects", connectingObjectsEntity)
				}
			}
		}

		/** associations of ProcessInstance **/

		/** parentPtr **/
		if results[0][13] != nil {
			id := results[0][13].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.DefinitionsInstance"
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
func (this *EntityManager) NewBPMSProcessInstanceEntityFromObject(object *BPMS.ProcessInstance) *BPMS_ProcessInstanceEntity {
	return this.NewBPMSProcessInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_ProcessInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSProcessInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.ProcessInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_ProcessInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_ProcessInstanceEntity) AppendReference(reference Entity) {

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
//              			EventDefinitionInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_EventDefinitionInstanceEntity struct {
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
	object         *BPMS.EventDefinitionInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMSEventDefinitionInstanceEntity(objectId string, object interface{}) *BPMS_EventDefinitionInstanceEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSEventDefinitionInstanceExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.EventDefinitionInstance).TYPENAME = "BPMS.EventDefinitionInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.EventDefinitionInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_EventDefinitionInstanceEntity)
		}
	} else {
		uuidStr = "BPMS.EventDefinitionInstance%" + Utility.RandomUUID()
	}
	entity := new(BPMS_EventDefinitionInstanceEntity)
	if object == nil {
		entity.object = new(BPMS.EventDefinitionInstance)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.EventDefinitionInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.EventDefinitionInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.EventDefinitionInstance", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_EventDefinitionInstanceEntity) GetTypeName() string {
	return "BPMS.EventDefinitionInstance"
}
func (this *BPMS_EventDefinitionInstanceEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_EventDefinitionInstanceEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_EventDefinitionInstanceEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_EventDefinitionInstanceEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_EventDefinitionInstanceEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_EventDefinitionInstanceEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_EventDefinitionInstanceEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_EventDefinitionInstanceEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_EventDefinitionInstanceEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_EventDefinitionInstanceEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_EventDefinitionInstanceEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_EventDefinitionInstanceEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_EventDefinitionInstanceEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_EventDefinitionInstanceEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_EventDefinitionInstanceEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_EventDefinitionInstanceEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_EventDefinitionInstanceEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_EventDefinitionInstanceEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_EventDefinitionInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_EventDefinitionInstanceEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_EventDefinitionInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_EventDefinitionInstanceEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_EventDefinitionInstanceEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.EventDefinitionInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_EventDefinitionInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_EventDefinitionInstanceEntityPrototype() {

	var eventDefinitionInstanceEntityProto EntityPrototype
	eventDefinitionInstanceEntityProto.TypeName = "BPMS.EventDefinitionInstance"
	eventDefinitionInstanceEntityProto.SuperTypeNames = append(eventDefinitionInstanceEntityProto.SuperTypeNames, "BPMS.Instance")
	eventDefinitionInstanceEntityProto.Ids = append(eventDefinitionInstanceEntityProto.Ids, "uuid")
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "uuid")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "xs.string")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 0)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, false)
	eventDefinitionInstanceEntityProto.Indexs = append(eventDefinitionInstanceEntityProto.Indexs, "parentUuid")
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "parentUuid")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "xs.string")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 1)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, false)

	/** members of Instance **/
	eventDefinitionInstanceEntityProto.Ids = append(eventDefinitionInstanceEntityProto.Ids, "M_id")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 2)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "M_id")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "xs.ID")
	eventDefinitionInstanceEntityProto.Indexs = append(eventDefinitionInstanceEntityProto.Indexs, "M_bpmnElementId")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 3)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "M_bpmnElementId")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "xs.string")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 4)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "M_participants")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "[]xs.string")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 5)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "M_dataRef")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 6)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "M_data")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 7)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "M_logInfoRef")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "[]BPMS.LogInfo:Ref")

	/** members of EventDefinitionInstance **/
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 8)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "M_eventDefinitionType")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "enum:EventDefinitionType_MessageEventDefinition:EventDefinitionType_LinkEventDefinition:EventDefinitionType_ErrorEventDefinition:EventDefinitionType_TerminateEventDefinition:EventDefinitionType_CompensationEventDefinition:EventDefinitionType_ConditionalEventDefinition:EventDefinitionType_TimerEventDefinition:EventDefinitionType_CancelEventDefinition:EventDefinitionType_EscalationEventDefinition")

	/** associations of EventDefinitionInstance **/
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 9)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, false)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "M_eventInstancePtr")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "BPMS.EventInstance:Ref")
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "childsUuid")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "[]xs.string")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 10)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, false)

	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields, "referenced")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType, "[]EntityRef")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder, 11)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&eventDefinitionInstanceEntityProto)

}

/** Create **/
func (this *BPMS_EventDefinitionInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.EventDefinitionInstance"

	var query EntityQuery
	query.TypeName = "BPMS.EventDefinitionInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of EventDefinitionInstance **/
	query.Fields = append(query.Fields, "M_eventDefinitionType")

	/** associations of EventDefinitionInstance **/
	query.Fields = append(query.Fields, "M_eventInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var EventDefinitionInstanceInfo []interface{}

	EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, this.parentPtr.GetUuid())
	} else {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, "")
	}

	/** members of Instance **/
	EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, this.object.M_id)
	EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, this.object.M_bpmnElementId)
	EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, this.object.M_participants)

	/** Save dataRef type ItemAwareElementInstance **/
	dataRefStr, _ := json.Marshal(this.object.M_dataRef)
	EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, string(dataRefStr))

	/** Save data type ItemAwareElementInstance **/
	dataIds := make([]string, 0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity := GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(this.object.M_data[i].UUID, this.object.M_data[i])
		dataIds = append(dataIds, dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data", dataEntity)
		if dataEntity.NeedSave() {
			dataEntity.SaveEntity()
		}
	}
	dataStr, _ := json.Marshal(dataIds)
	EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, string(dataStr))

	/** Save logInfoRef type LogInfo **/
	logInfoRefStr, _ := json.Marshal(this.object.M_logInfoRef)
	EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, string(logInfoRefStr))

	/** members of EventDefinitionInstance **/

	/** Save eventDefinitionType type EventDefinitionType **/
	if this.object.M_eventDefinitionType == BPMS.EventDefinitionType_MessageEventDefinition {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 0)
	} else if this.object.M_eventDefinitionType == BPMS.EventDefinitionType_LinkEventDefinition {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 1)
	} else if this.object.M_eventDefinitionType == BPMS.EventDefinitionType_ErrorEventDefinition {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 2)
	} else if this.object.M_eventDefinitionType == BPMS.EventDefinitionType_TerminateEventDefinition {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 3)
	} else if this.object.M_eventDefinitionType == BPMS.EventDefinitionType_CompensationEventDefinition {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 4)
	} else if this.object.M_eventDefinitionType == BPMS.EventDefinitionType_ConditionalEventDefinition {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 5)
	} else if this.object.M_eventDefinitionType == BPMS.EventDefinitionType_TimerEventDefinition {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 6)
	} else if this.object.M_eventDefinitionType == BPMS.EventDefinitionType_CancelEventDefinition {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 7)
	} else if this.object.M_eventDefinitionType == BPMS.EventDefinitionType_EscalationEventDefinition {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 8)
	} else {
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 0)
	}

	/** associations of EventDefinitionInstance **/

	/** Save eventInstance type EventInstance **/
	EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, this.object.M_eventInstancePtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), EventDefinitionInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), EventDefinitionInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_EventDefinitionInstanceEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_EventDefinitionInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.EventDefinitionInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Instance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_participants")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_logInfoRef")

	/** members of EventDefinitionInstance **/
	query.Fields = append(query.Fields, "M_eventDefinitionType")

	/** associations of EventDefinitionInstance **/
	query.Fields = append(query.Fields, "M_eventInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of EventDefinitionInstance...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.EventDefinitionInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.EventDefinitionInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** bpmnElementId **/
		if results[0][3] != nil {
			this.object.M_bpmnElementId = results[0][3].(string)
		}

		/** participants **/
		if results[0][4] != nil {
			this.object.M_participants = append(this.object.M_participants, results[0][4].([]string)...)
		}

		/** dataRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ItemAwareElementInstance"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef, ids[i])
					GetServer().GetEntityManager().appendReference("dataRef", this.object.UUID, id_)
				}
			}
		}

		/** data **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_ItemAwareElementInstanceEntity)
					} else {
						dataEntity = GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data", dataEntity)
				}
			}
		}

		/** logInfoRef **/
		if results[0][7] != nil {
			idsStr := results[0][7].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.LogInfo"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef, ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef", this.object.UUID, id_)
				}
			}
		}

		/** members of EventDefinitionInstance **/

		/** eventDefinitionType **/
		if results[0][8] != nil {
			enumIndex := results[0][8].(int)
			if enumIndex == 0 {
				this.object.M_eventDefinitionType = BPMS.EventDefinitionType_MessageEventDefinition
			} else if enumIndex == 1 {
				this.object.M_eventDefinitionType = BPMS.EventDefinitionType_LinkEventDefinition
			} else if enumIndex == 2 {
				this.object.M_eventDefinitionType = BPMS.EventDefinitionType_ErrorEventDefinition
			} else if enumIndex == 3 {
				this.object.M_eventDefinitionType = BPMS.EventDefinitionType_TerminateEventDefinition
			} else if enumIndex == 4 {
				this.object.M_eventDefinitionType = BPMS.EventDefinitionType_CompensationEventDefinition
			} else if enumIndex == 5 {
				this.object.M_eventDefinitionType = BPMS.EventDefinitionType_ConditionalEventDefinition
			} else if enumIndex == 6 {
				this.object.M_eventDefinitionType = BPMS.EventDefinitionType_TimerEventDefinition
			} else if enumIndex == 7 {
				this.object.M_eventDefinitionType = BPMS.EventDefinitionType_CancelEventDefinition
			} else if enumIndex == 8 {
				this.object.M_eventDefinitionType = BPMS.EventDefinitionType_EscalationEventDefinition
			}
		}

		/** associations of EventDefinitionInstance **/

		/** eventInstancePtr **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.EventInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_eventInstancePtr = id
				GetServer().GetEntityManager().appendReference("eventInstancePtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][10].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][11].(string)
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
func (this *EntityManager) NewBPMSEventDefinitionInstanceEntityFromObject(object *BPMS.EventDefinitionInstance) *BPMS_EventDefinitionInstanceEntity {
	return this.NewBPMSEventDefinitionInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_EventDefinitionInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSEventDefinitionInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.EventDefinitionInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_EventDefinitionInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_EventDefinitionInstanceEntity) AppendReference(reference Entity) {

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
//              			CorrelationInfo
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_CorrelationInfoEntity struct {
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
	object         *BPMS.CorrelationInfo
}

/** Constructor function **/
func (this *EntityManager) NewBPMSCorrelationInfoEntity(objectId string, object interface{}) *BPMS_CorrelationInfoEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSCorrelationInfoExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.CorrelationInfo).TYPENAME = "BPMS.CorrelationInfo"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.CorrelationInfo).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_CorrelationInfoEntity)
		}
	} else {
		uuidStr = "BPMS.CorrelationInfo%" + Utility.RandomUUID()
	}
	entity := new(BPMS_CorrelationInfoEntity)
	if object == nil {
		entity.object = new(BPMS.CorrelationInfo)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.CorrelationInfo)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.CorrelationInfo"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.CorrelationInfo", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_CorrelationInfoEntity) GetTypeName() string {
	return "BPMS.CorrelationInfo"
}
func (this *BPMS_CorrelationInfoEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_CorrelationInfoEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_CorrelationInfoEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_CorrelationInfoEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_CorrelationInfoEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_CorrelationInfoEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_CorrelationInfoEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_CorrelationInfoEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_CorrelationInfoEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_CorrelationInfoEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_CorrelationInfoEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_CorrelationInfoEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_CorrelationInfoEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_CorrelationInfoEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_CorrelationInfoEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_CorrelationInfoEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_CorrelationInfoEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_CorrelationInfoEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_CorrelationInfoEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_CorrelationInfoEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_CorrelationInfoEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_CorrelationInfoEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_CorrelationInfoEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.CorrelationInfo"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_CorrelationInfoEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_CorrelationInfoEntityPrototype() {

	var correlationInfoEntityProto EntityPrototype
	correlationInfoEntityProto.TypeName = "BPMS.CorrelationInfo"
	correlationInfoEntityProto.Ids = append(correlationInfoEntityProto.Ids, "uuid")
	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields, "uuid")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType, "xs.string")
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder, 0)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility, false)
	correlationInfoEntityProto.Indexs = append(correlationInfoEntityProto.Indexs, "parentUuid")
	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields, "parentUuid")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType, "xs.string")
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder, 1)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility, false)

	/** members of CorrelationInfo **/
	correlationInfoEntityProto.Ids = append(correlationInfoEntityProto.Ids, "M_id")
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder, 2)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility, true)
	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields, "M_id")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType, "xs.ID")

	/** associations of CorrelationInfo **/
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder, 3)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility, false)
	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields, "M_runtimesPtr")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType, "BPMS.Runtimes:Ref")
	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields, "childsUuid")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType, "[]xs.string")
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder, 4)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility, false)

	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields, "referenced")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType, "[]EntityRef")
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder, 5)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&correlationInfoEntityProto)

}

/** Create **/
func (this *BPMS_CorrelationInfoEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.CorrelationInfo"

	var query EntityQuery
	query.TypeName = "BPMS.CorrelationInfo"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of CorrelationInfo **/
	query.Fields = append(query.Fields, "M_id")

	/** associations of CorrelationInfo **/
	query.Fields = append(query.Fields, "M_runtimesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var CorrelationInfoInfo []interface{}

	CorrelationInfoInfo = append(CorrelationInfoInfo, this.GetUuid())
	if this.parentPtr != nil {
		CorrelationInfoInfo = append(CorrelationInfoInfo, this.parentPtr.GetUuid())
	} else {
		CorrelationInfoInfo = append(CorrelationInfoInfo, "")
	}

	/** members of CorrelationInfo **/
	CorrelationInfoInfo = append(CorrelationInfoInfo, this.object.M_id)

	/** associations of CorrelationInfo **/

	/** Save runtimes type Runtimes **/
	CorrelationInfoInfo = append(CorrelationInfoInfo, this.object.M_runtimesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	CorrelationInfoInfo = append(CorrelationInfoInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	CorrelationInfoInfo = append(CorrelationInfoInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), CorrelationInfoInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), CorrelationInfoInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_CorrelationInfoEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_CorrelationInfoEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.CorrelationInfo"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of CorrelationInfo **/
	query.Fields = append(query.Fields, "M_id")

	/** associations of CorrelationInfo **/
	query.Fields = append(query.Fields, "M_runtimesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of CorrelationInfo...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.CorrelationInfo)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.CorrelationInfo"

		this.parentUuid = results[0][1].(string)

		/** members of CorrelationInfo **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** associations of CorrelationInfo **/

		/** runtimesPtr **/
		if results[0][3] != nil {
			id := results[0][3].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.Runtimes"
				id_ := refTypeName + "$$" + id
				this.object.M_runtimesPtr = id
				GetServer().GetEntityManager().appendReference("runtimesPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][4].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][5].(string)
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
func (this *EntityManager) NewBPMSCorrelationInfoEntityFromObject(object *BPMS.CorrelationInfo) *BPMS_CorrelationInfoEntity {
	return this.NewBPMSCorrelationInfoEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_CorrelationInfoEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSCorrelationInfoExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.CorrelationInfo"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_CorrelationInfoEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_CorrelationInfoEntity) AppendReference(reference Entity) {

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
//              			ItemAwareElementInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_ItemAwareElementInstanceEntity struct {
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
	object         *BPMS.ItemAwareElementInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMSItemAwareElementInstanceEntity(objectId string, object interface{}) *BPMS_ItemAwareElementInstanceEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSItemAwareElementInstanceExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.ItemAwareElementInstance).TYPENAME = "BPMS.ItemAwareElementInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.ItemAwareElementInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_ItemAwareElementInstanceEntity)
		}
	} else {
		uuidStr = "BPMS.ItemAwareElementInstance%" + Utility.RandomUUID()
	}
	entity := new(BPMS_ItemAwareElementInstanceEntity)
	if object == nil {
		entity.object = new(BPMS.ItemAwareElementInstance)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.ItemAwareElementInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.ItemAwareElementInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.ItemAwareElementInstance", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_ItemAwareElementInstanceEntity) GetTypeName() string {
	return "BPMS.ItemAwareElementInstance"
}
func (this *BPMS_ItemAwareElementInstanceEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_ItemAwareElementInstanceEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_ItemAwareElementInstanceEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_ItemAwareElementInstanceEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_ItemAwareElementInstanceEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_ItemAwareElementInstanceEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_ItemAwareElementInstanceEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_ItemAwareElementInstanceEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_ItemAwareElementInstanceEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_ItemAwareElementInstanceEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_ItemAwareElementInstanceEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_ItemAwareElementInstanceEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_ItemAwareElementInstanceEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_ItemAwareElementInstanceEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_ItemAwareElementInstanceEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_ItemAwareElementInstanceEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_ItemAwareElementInstanceEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_ItemAwareElementInstanceEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_ItemAwareElementInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_ItemAwareElementInstanceEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_ItemAwareElementInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_ItemAwareElementInstanceEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_ItemAwareElementInstanceEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.ItemAwareElementInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_ItemAwareElementInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_ItemAwareElementInstanceEntityPrototype() {

	var itemAwareElementInstanceEntityProto EntityPrototype
	itemAwareElementInstanceEntityProto.TypeName = "BPMS.ItemAwareElementInstance"
	itemAwareElementInstanceEntityProto.Ids = append(itemAwareElementInstanceEntityProto.Ids, "uuid")
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields, "uuid")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType, "xs.string")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder, 0)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility, false)
	itemAwareElementInstanceEntityProto.Indexs = append(itemAwareElementInstanceEntityProto.Indexs, "parentUuid")
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields, "parentUuid")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType, "xs.string")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder, 1)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility, false)

	/** members of ItemAwareElementInstance **/
	itemAwareElementInstanceEntityProto.Ids = append(itemAwareElementInstanceEntityProto.Ids, "M_id")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder, 2)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility, true)
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields, "M_id")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType, "xs.ID")
	itemAwareElementInstanceEntityProto.Ids = append(itemAwareElementInstanceEntityProto.Ids, "M_bpmnElementId")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder, 3)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility, true)
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields, "M_bpmnElementId")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType, "xs.ID")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder, 4)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility, true)
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields, "M_data")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType, "xs.[]uint8")

	/** associations of ItemAwareElementInstance **/
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder, 5)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility, false)
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields, "M_parentPtr")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType, "BPMS.Instance:Ref")
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields, "childsUuid")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType, "[]xs.string")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder, 6)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility, false)

	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields, "referenced")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType, "[]EntityRef")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder, 7)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&itemAwareElementInstanceEntityProto)

}

/** Create **/
func (this *BPMS_ItemAwareElementInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.ItemAwareElementInstance"

	var query EntityQuery
	query.TypeName = "BPMS.ItemAwareElementInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of ItemAwareElementInstance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_data")

	/** associations of ItemAwareElementInstance **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ItemAwareElementInstanceInfo []interface{}

	ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, this.parentPtr.GetUuid())
	} else {
		ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, "")
	}

	/** members of ItemAwareElementInstance **/
	ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, this.object.M_id)
	ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, this.object.M_bpmnElementId)
	ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, this.object.M_data)

	/** associations of ItemAwareElementInstance **/

	/** Save parent type Instance **/
	ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, this.object.M_parentPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), ItemAwareElementInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), ItemAwareElementInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_ItemAwareElementInstanceEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_ItemAwareElementInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.ItemAwareElementInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of ItemAwareElementInstance **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_data")

	/** associations of ItemAwareElementInstance **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ItemAwareElementInstance...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.ItemAwareElementInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.ItemAwareElementInstance"

		this.parentUuid = results[0][1].(string)

		/** members of ItemAwareElementInstance **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** bpmnElementId **/
		if results[0][3] != nil {
			this.object.M_bpmnElementId = results[0][3].(string)
		}

		/** data **/
		if results[0][4] != nil {
			this.object.M_data = results[0][4].([]uint8)
		}

		/** associations of ItemAwareElementInstance **/

		/** parentPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.Instance"
				id_ := refTypeName + "$$" + id
				this.object.M_parentPtr = id
				GetServer().GetEntityManager().appendReference("parentPtr", this.object.UUID, id_)
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
func (this *EntityManager) NewBPMSItemAwareElementInstanceEntityFromObject(object *BPMS.ItemAwareElementInstance) *BPMS_ItemAwareElementInstanceEntity {
	return this.NewBPMSItemAwareElementInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_ItemAwareElementInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSItemAwareElementInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.ItemAwareElementInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_ItemAwareElementInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_ItemAwareElementInstanceEntity) AppendReference(reference Entity) {

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
//              			RessourceParameterInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_RessourceParameterInstanceEntity struct {
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
	object         *BPMS.RessourceParameterInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMSRessourceParameterInstanceEntity(objectId string, object interface{}) *BPMS_RessourceParameterInstanceEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSRessourceParameterInstanceExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.RessourceParameterInstance).TYPENAME = "BPMS.RessourceParameterInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.RessourceParameterInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_RessourceParameterInstanceEntity)
		}
	} else {
		uuidStr = "BPMS.RessourceParameterInstance%" + Utility.RandomUUID()
	}
	entity := new(BPMS_RessourceParameterInstanceEntity)
	if object == nil {
		entity.object = new(BPMS.RessourceParameterInstance)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.RessourceParameterInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.RessourceParameterInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.RessourceParameterInstance", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_RessourceParameterInstanceEntity) GetTypeName() string {
	return "BPMS.RessourceParameterInstance"
}
func (this *BPMS_RessourceParameterInstanceEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_RessourceParameterInstanceEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_RessourceParameterInstanceEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_RessourceParameterInstanceEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_RessourceParameterInstanceEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_RessourceParameterInstanceEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_RessourceParameterInstanceEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_RessourceParameterInstanceEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_RessourceParameterInstanceEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_RessourceParameterInstanceEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_RessourceParameterInstanceEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_RessourceParameterInstanceEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_RessourceParameterInstanceEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_RessourceParameterInstanceEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_RessourceParameterInstanceEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_RessourceParameterInstanceEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_RessourceParameterInstanceEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_RessourceParameterInstanceEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_RessourceParameterInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_RessourceParameterInstanceEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_RessourceParameterInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_RessourceParameterInstanceEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_RessourceParameterInstanceEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.RessourceParameterInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_RessourceParameterInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_RessourceParameterInstanceEntityPrototype() {

	var ressourceParameterInstanceEntityProto EntityPrototype
	ressourceParameterInstanceEntityProto.TypeName = "BPMS.RessourceParameterInstance"
	ressourceParameterInstanceEntityProto.Ids = append(ressourceParameterInstanceEntityProto.Ids, "uuid")
	ressourceParameterInstanceEntityProto.Fields = append(ressourceParameterInstanceEntityProto.Fields, "uuid")
	ressourceParameterInstanceEntityProto.FieldsType = append(ressourceParameterInstanceEntityProto.FieldsType, "xs.string")
	ressourceParameterInstanceEntityProto.FieldsOrder = append(ressourceParameterInstanceEntityProto.FieldsOrder, 0)
	ressourceParameterInstanceEntityProto.FieldsVisibility = append(ressourceParameterInstanceEntityProto.FieldsVisibility, false)
	ressourceParameterInstanceEntityProto.Indexs = append(ressourceParameterInstanceEntityProto.Indexs, "parentUuid")
	ressourceParameterInstanceEntityProto.Fields = append(ressourceParameterInstanceEntityProto.Fields, "parentUuid")
	ressourceParameterInstanceEntityProto.FieldsType = append(ressourceParameterInstanceEntityProto.FieldsType, "xs.string")
	ressourceParameterInstanceEntityProto.FieldsOrder = append(ressourceParameterInstanceEntityProto.FieldsOrder, 1)
	ressourceParameterInstanceEntityProto.FieldsVisibility = append(ressourceParameterInstanceEntityProto.FieldsVisibility, false)

	/** members of RessourceParameterInstance **/
	ressourceParameterInstanceEntityProto.Ids = append(ressourceParameterInstanceEntityProto.Ids, "M_bpmnElementId")
	ressourceParameterInstanceEntityProto.FieldsOrder = append(ressourceParameterInstanceEntityProto.FieldsOrder, 2)
	ressourceParameterInstanceEntityProto.FieldsVisibility = append(ressourceParameterInstanceEntityProto.FieldsVisibility, true)
	ressourceParameterInstanceEntityProto.Fields = append(ressourceParameterInstanceEntityProto.Fields, "M_bpmnElementId")
	ressourceParameterInstanceEntityProto.FieldsType = append(ressourceParameterInstanceEntityProto.FieldsType, "xs.ID")
	ressourceParameterInstanceEntityProto.FieldsOrder = append(ressourceParameterInstanceEntityProto.FieldsOrder, 3)
	ressourceParameterInstanceEntityProto.FieldsVisibility = append(ressourceParameterInstanceEntityProto.FieldsVisibility, true)
	ressourceParameterInstanceEntityProto.Fields = append(ressourceParameterInstanceEntityProto.Fields, "M_data")
	ressourceParameterInstanceEntityProto.FieldsType = append(ressourceParameterInstanceEntityProto.FieldsType, "BPMS.ItemAwareElementInstance")

	/** associations of RessourceParameterInstance **/
	ressourceParameterInstanceEntityProto.FieldsOrder = append(ressourceParameterInstanceEntityProto.FieldsOrder, 4)
	ressourceParameterInstanceEntityProto.FieldsVisibility = append(ressourceParameterInstanceEntityProto.FieldsVisibility, false)
	ressourceParameterInstanceEntityProto.Fields = append(ressourceParameterInstanceEntityProto.Fields, "M_parentPtr")
	ressourceParameterInstanceEntityProto.FieldsType = append(ressourceParameterInstanceEntityProto.FieldsType, "BPMS.RessourceInstance:Ref")
	ressourceParameterInstanceEntityProto.Fields = append(ressourceParameterInstanceEntityProto.Fields, "childsUuid")
	ressourceParameterInstanceEntityProto.FieldsType = append(ressourceParameterInstanceEntityProto.FieldsType, "[]xs.string")
	ressourceParameterInstanceEntityProto.FieldsOrder = append(ressourceParameterInstanceEntityProto.FieldsOrder, 5)
	ressourceParameterInstanceEntityProto.FieldsVisibility = append(ressourceParameterInstanceEntityProto.FieldsVisibility, false)

	ressourceParameterInstanceEntityProto.Fields = append(ressourceParameterInstanceEntityProto.Fields, "referenced")
	ressourceParameterInstanceEntityProto.FieldsType = append(ressourceParameterInstanceEntityProto.FieldsType, "[]EntityRef")
	ressourceParameterInstanceEntityProto.FieldsOrder = append(ressourceParameterInstanceEntityProto.FieldsOrder, 6)
	ressourceParameterInstanceEntityProto.FieldsVisibility = append(ressourceParameterInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&ressourceParameterInstanceEntityProto)

}

/** Create **/
func (this *BPMS_RessourceParameterInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.RessourceParameterInstance"

	var query EntityQuery
	query.TypeName = "BPMS.RessourceParameterInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of RessourceParameterInstance **/
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_data")

	/** associations of RessourceParameterInstance **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var RessourceParameterInstanceInfo []interface{}

	RessourceParameterInstanceInfo = append(RessourceParameterInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		RessourceParameterInstanceInfo = append(RessourceParameterInstanceInfo, this.parentPtr.GetUuid())
	} else {
		RessourceParameterInstanceInfo = append(RessourceParameterInstanceInfo, "")
	}

	/** members of RessourceParameterInstance **/
	RessourceParameterInstanceInfo = append(RessourceParameterInstanceInfo, this.object.M_bpmnElementId)

	/** Save data type ItemAwareElementInstance **/
	if this.object.M_data != nil {
		dataEntity := GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(this.object.M_data.UUID, this.object.M_data)
		RessourceParameterInstanceInfo = append(RessourceParameterInstanceInfo, dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data", dataEntity)
		if dataEntity.NeedSave() {
			dataEntity.SaveEntity()
		}
	} else {
		RessourceParameterInstanceInfo = append(RessourceParameterInstanceInfo, "")
	}

	/** associations of RessourceParameterInstance **/

	/** Save parent type RessourceInstance **/
	/** attribute RessourceInstance has no method GetId, must be an error here...*/
	RessourceParameterInstanceInfo = append(RessourceParameterInstanceInfo, "")
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	RessourceParameterInstanceInfo = append(RessourceParameterInstanceInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	RessourceParameterInstanceInfo = append(RessourceParameterInstanceInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), RessourceParameterInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), RessourceParameterInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_RessourceParameterInstanceEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_RessourceParameterInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.RessourceParameterInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of RessourceParameterInstance **/
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_data")

	/** associations of RessourceParameterInstance **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of RessourceParameterInstance...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.RessourceParameterInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.RessourceParameterInstance"

		this.parentUuid = results[0][1].(string)

		/** members of RessourceParameterInstance **/

		/** bpmnElementId **/
		if results[0][2] != nil {
			this.object.M_bpmnElementId = results[0][2].(string)
		}

		/** data **/
		if results[0][3] != nil {
			uuid := results[0][3].(string)
			if len(uuid) > 0 {
				var dataEntity *BPMS_ItemAwareElementInstanceEntity
				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
					dataEntity = instance.(*BPMS_ItemAwareElementInstanceEntity)
				} else {
					dataEntity = GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity(uuid, nil)
					dataEntity.InitEntity(uuid)
					GetServer().GetEntityManager().insert(dataEntity)
				}
				dataEntity.AppendReferenced("data", this)
				this.AppendChild("data", dataEntity)
			}
		}

		/** associations of RessourceParameterInstance **/

		/** parentPtr **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.RessourceInstance"
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
func (this *EntityManager) NewBPMSRessourceParameterInstanceEntityFromObject(object *BPMS.RessourceParameterInstance) *BPMS_RessourceParameterInstanceEntity {
	return this.NewBPMSRessourceParameterInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_RessourceParameterInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSRessourceParameterInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.RessourceParameterInstance"
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_RessourceParameterInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_RessourceParameterInstanceEntity) AppendReference(reference Entity) {

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
//              			RessourceInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_RessourceInstanceEntity struct {
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
	object         *BPMS.RessourceInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMSRessourceInstanceEntity(objectId string, object interface{}) *BPMS_RessourceInstanceEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSRessourceInstanceExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.RessourceInstance).TYPENAME = "BPMS.RessourceInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.RessourceInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_RessourceInstanceEntity)
		}
	} else {
		uuidStr = "BPMS.RessourceInstance%" + Utility.RandomUUID()
	}
	entity := new(BPMS_RessourceInstanceEntity)
	if object == nil {
		entity.object = new(BPMS.RessourceInstance)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.RessourceInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.RessourceInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.RessourceInstance", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_RessourceInstanceEntity) GetTypeName() string {
	return "BPMS.RessourceInstance"
}
func (this *BPMS_RessourceInstanceEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_RessourceInstanceEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_RessourceInstanceEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_RessourceInstanceEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_RessourceInstanceEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_RessourceInstanceEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_RessourceInstanceEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_RessourceInstanceEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_RessourceInstanceEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_RessourceInstanceEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_RessourceInstanceEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_RessourceInstanceEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_RessourceInstanceEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_RessourceInstanceEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_RessourceInstanceEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_RessourceInstanceEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_RessourceInstanceEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_RessourceInstanceEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_RessourceInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_RessourceInstanceEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_RessourceInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_RessourceInstanceEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_RessourceInstanceEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.RessourceInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_RessourceInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_RessourceInstanceEntityPrototype() {

	var ressourceInstanceEntityProto EntityPrototype
	ressourceInstanceEntityProto.TypeName = "BPMS.RessourceInstance"
	ressourceInstanceEntityProto.Ids = append(ressourceInstanceEntityProto.Ids, "uuid")
	ressourceInstanceEntityProto.Fields = append(ressourceInstanceEntityProto.Fields, "uuid")
	ressourceInstanceEntityProto.FieldsType = append(ressourceInstanceEntityProto.FieldsType, "xs.string")
	ressourceInstanceEntityProto.FieldsOrder = append(ressourceInstanceEntityProto.FieldsOrder, 0)
	ressourceInstanceEntityProto.FieldsVisibility = append(ressourceInstanceEntityProto.FieldsVisibility, false)
	ressourceInstanceEntityProto.Indexs = append(ressourceInstanceEntityProto.Indexs, "parentUuid")
	ressourceInstanceEntityProto.Fields = append(ressourceInstanceEntityProto.Fields, "parentUuid")
	ressourceInstanceEntityProto.FieldsType = append(ressourceInstanceEntityProto.FieldsType, "xs.string")
	ressourceInstanceEntityProto.FieldsOrder = append(ressourceInstanceEntityProto.FieldsOrder, 1)
	ressourceInstanceEntityProto.FieldsVisibility = append(ressourceInstanceEntityProto.FieldsVisibility, false)

	/** members of RessourceInstance **/
	ressourceInstanceEntityProto.Ids = append(ressourceInstanceEntityProto.Ids, "M_bpmnElementId")
	ressourceInstanceEntityProto.FieldsOrder = append(ressourceInstanceEntityProto.FieldsOrder, 2)
	ressourceInstanceEntityProto.FieldsVisibility = append(ressourceInstanceEntityProto.FieldsVisibility, true)
	ressourceInstanceEntityProto.Fields = append(ressourceInstanceEntityProto.Fields, "M_bpmnElementId")
	ressourceInstanceEntityProto.FieldsType = append(ressourceInstanceEntityProto.FieldsType, "xs.ID")
	ressourceInstanceEntityProto.FieldsOrder = append(ressourceInstanceEntityProto.FieldsOrder, 3)
	ressourceInstanceEntityProto.FieldsVisibility = append(ressourceInstanceEntityProto.FieldsVisibility, true)
	ressourceInstanceEntityProto.Fields = append(ressourceInstanceEntityProto.Fields, "M_parameters")
	ressourceInstanceEntityProto.FieldsType = append(ressourceInstanceEntityProto.FieldsType, "[]BPMS.RessourceParameterInstance")

	/** associations of RessourceInstance **/
	ressourceInstanceEntityProto.FieldsOrder = append(ressourceInstanceEntityProto.FieldsOrder, 4)
	ressourceInstanceEntityProto.FieldsVisibility = append(ressourceInstanceEntityProto.FieldsVisibility, false)
	ressourceInstanceEntityProto.Fields = append(ressourceInstanceEntityProto.Fields, "M_parentPtr")
	ressourceInstanceEntityProto.FieldsType = append(ressourceInstanceEntityProto.FieldsType, "BPMS.Instance:Ref")
	ressourceInstanceEntityProto.Fields = append(ressourceInstanceEntityProto.Fields, "childsUuid")
	ressourceInstanceEntityProto.FieldsType = append(ressourceInstanceEntityProto.FieldsType, "[]xs.string")
	ressourceInstanceEntityProto.FieldsOrder = append(ressourceInstanceEntityProto.FieldsOrder, 5)
	ressourceInstanceEntityProto.FieldsVisibility = append(ressourceInstanceEntityProto.FieldsVisibility, false)

	ressourceInstanceEntityProto.Fields = append(ressourceInstanceEntityProto.Fields, "referenced")
	ressourceInstanceEntityProto.FieldsType = append(ressourceInstanceEntityProto.FieldsType, "[]EntityRef")
	ressourceInstanceEntityProto.FieldsOrder = append(ressourceInstanceEntityProto.FieldsOrder, 6)
	ressourceInstanceEntityProto.FieldsVisibility = append(ressourceInstanceEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&ressourceInstanceEntityProto)

}

/** Create **/
func (this *BPMS_RessourceInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.RessourceInstance"

	var query EntityQuery
	query.TypeName = "BPMS.RessourceInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of RessourceInstance **/
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_parameters")

	/** associations of RessourceInstance **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var RessourceInstanceInfo []interface{}

	RessourceInstanceInfo = append(RessourceInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		RessourceInstanceInfo = append(RessourceInstanceInfo, this.parentPtr.GetUuid())
	} else {
		RessourceInstanceInfo = append(RessourceInstanceInfo, "")
	}

	/** members of RessourceInstance **/
	RessourceInstanceInfo = append(RessourceInstanceInfo, this.object.M_bpmnElementId)

	/** Save parameters type RessourceParameterInstance **/
	parametersIds := make([]string, 0)
	for i := 0; i < len(this.object.M_parameters); i++ {
		parametersEntity := GetServer().GetEntityManager().NewBPMSRessourceParameterInstanceEntity(this.object.M_parameters[i].UUID, this.object.M_parameters[i])
		parametersIds = append(parametersIds, parametersEntity.uuid)
		parametersEntity.AppendReferenced("parameters", this)
		this.AppendChild("parameters", parametersEntity)
		if parametersEntity.NeedSave() {
			parametersEntity.SaveEntity()
		}
	}
	parametersStr, _ := json.Marshal(parametersIds)
	RessourceInstanceInfo = append(RessourceInstanceInfo, string(parametersStr))

	/** associations of RessourceInstance **/

	/** Save parent type Instance **/
	RessourceInstanceInfo = append(RessourceInstanceInfo, this.object.M_parentPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	RessourceInstanceInfo = append(RessourceInstanceInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	RessourceInstanceInfo = append(RessourceInstanceInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), RessourceInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), RessourceInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_RessourceInstanceEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_RessourceInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.RessourceInstance"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of RessourceInstance **/
	query.Fields = append(query.Fields, "M_bpmnElementId")
	query.Fields = append(query.Fields, "M_parameters")

	/** associations of RessourceInstance **/
	query.Fields = append(query.Fields, "M_parentPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of RessourceInstance...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.RessourceInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.RessourceInstance"

		this.parentUuid = results[0][1].(string)

		/** members of RessourceInstance **/

		/** bpmnElementId **/
		if results[0][2] != nil {
			this.object.M_bpmnElementId = results[0][2].(string)
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
					var parametersEntity *BPMS_RessourceParameterInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						parametersEntity = instance.(*BPMS_RessourceParameterInstanceEntity)
					} else {
						parametersEntity = GetServer().GetEntityManager().NewBPMSRessourceParameterInstanceEntity(uuids[i], nil)
						parametersEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(parametersEntity)
					}
					parametersEntity.AppendReferenced("parameters", this)
					this.AppendChild("parameters", parametersEntity)
				}
			}
		}

		/** associations of RessourceInstance **/

		/** parentPtr **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.Instance"
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
func (this *EntityManager) NewBPMSRessourceInstanceEntityFromObject(object *BPMS.RessourceInstance) *BPMS_RessourceInstanceEntity {
	return this.NewBPMSRessourceInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_RessourceInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSRessourceInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.RessourceInstance"
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_RessourceInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_RessourceInstanceEntity) AppendReference(reference Entity) {

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
//              			EventData
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_EventDataEntity struct {
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
	object         *BPMS.EventData
}

/** Constructor function **/
func (this *EntityManager) NewBPMSEventDataEntity(objectId string, object interface{}) *BPMS_EventDataEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSEventDataExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.EventData).TYPENAME = "BPMS.EventData"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.EventData).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_EventDataEntity)
		}
	} else {
		uuidStr = "BPMS.EventData%" + Utility.RandomUUID()
	}
	entity := new(BPMS_EventDataEntity)
	if object == nil {
		entity.object = new(BPMS.EventData)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.EventData)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.EventData"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.EventData", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_EventDataEntity) GetTypeName() string {
	return "BPMS.EventData"
}
func (this *BPMS_EventDataEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_EventDataEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_EventDataEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_EventDataEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_EventDataEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_EventDataEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_EventDataEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_EventDataEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_EventDataEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_EventDataEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_EventDataEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_EventDataEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_EventDataEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_EventDataEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_EventDataEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_EventDataEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_EventDataEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_EventDataEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_EventDataEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_EventDataEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_EventDataEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_EventDataEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_EventDataEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.EventData"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_EventDataEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_EventDataEntityPrototype() {

	var eventDataEntityProto EntityPrototype
	eventDataEntityProto.TypeName = "BPMS.EventData"
	eventDataEntityProto.Ids = append(eventDataEntityProto.Ids, "uuid")
	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields, "uuid")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType, "xs.string")
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder, 0)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility, false)
	eventDataEntityProto.Indexs = append(eventDataEntityProto.Indexs, "parentUuid")
	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields, "parentUuid")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType, "xs.string")
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder, 1)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility, false)

	/** members of EventData **/
	eventDataEntityProto.Ids = append(eventDataEntityProto.Ids, "M_id")
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder, 2)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility, true)
	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields, "M_id")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType, "xs.ID")

	/** associations of EventData **/
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder, 3)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility, false)
	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields, "M_triggerPtr")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType, "BPMS.Trigger:Ref")
	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields, "childsUuid")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType, "[]xs.string")
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder, 4)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility, false)

	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields, "referenced")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType, "[]EntityRef")
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder, 5)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&eventDataEntityProto)

}

/** Create **/
func (this *BPMS_EventDataEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.EventData"

	var query EntityQuery
	query.TypeName = "BPMS.EventData"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of EventData **/
	query.Fields = append(query.Fields, "M_id")

	/** associations of EventData **/
	query.Fields = append(query.Fields, "M_triggerPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var EventDataInfo []interface{}

	EventDataInfo = append(EventDataInfo, this.GetUuid())
	if this.parentPtr != nil {
		EventDataInfo = append(EventDataInfo, this.parentPtr.GetUuid())
	} else {
		EventDataInfo = append(EventDataInfo, "")
	}

	/** members of EventData **/
	EventDataInfo = append(EventDataInfo, this.object.M_id)

	/** associations of EventData **/

	/** Save trigger type Trigger **/
	EventDataInfo = append(EventDataInfo, this.object.M_triggerPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	EventDataInfo = append(EventDataInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	EventDataInfo = append(EventDataInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), EventDataInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), EventDataInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_EventDataEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_EventDataEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.EventData"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of EventData **/
	query.Fields = append(query.Fields, "M_id")

	/** associations of EventData **/
	query.Fields = append(query.Fields, "M_triggerPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of EventData...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.EventData)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.EventData"

		this.parentUuid = results[0][1].(string)

		/** members of EventData **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** associations of EventData **/

		/** triggerPtr **/
		if results[0][3] != nil {
			id := results[0][3].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.Trigger"
				id_ := refTypeName + "$$" + id
				this.object.M_triggerPtr = id
				GetServer().GetEntityManager().appendReference("triggerPtr", this.object.UUID, id_)
			}
		}
		childsUuidStr := results[0][4].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][5].(string)
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
func (this *EntityManager) NewBPMSEventDataEntityFromObject(object *BPMS.EventData) *BPMS_EventDataEntity {
	return this.NewBPMSEventDataEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_EventDataEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSEventDataExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.EventData"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_EventDataEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_EventDataEntity) AppendReference(reference Entity) {

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
//              			Trigger
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_TriggerEntity struct {
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
	object         *BPMS.Trigger
}

/** Constructor function **/
func (this *EntityManager) NewBPMSTriggerEntity(objectId string, object interface{}) *BPMS_TriggerEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSTriggerExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.Trigger).TYPENAME = "BPMS.Trigger"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.Trigger).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_TriggerEntity)
		}
	} else {
		uuidStr = "BPMS.Trigger%" + Utility.RandomUUID()
	}
	entity := new(BPMS_TriggerEntity)
	if object == nil {
		entity.object = new(BPMS.Trigger)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.Trigger)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.Trigger"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.Trigger", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_TriggerEntity) GetTypeName() string {
	return "BPMS.Trigger"
}
func (this *BPMS_TriggerEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_TriggerEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_TriggerEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_TriggerEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_TriggerEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_TriggerEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_TriggerEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_TriggerEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_TriggerEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_TriggerEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_TriggerEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_TriggerEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_TriggerEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_TriggerEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_TriggerEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_TriggerEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_TriggerEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_TriggerEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_TriggerEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_TriggerEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_TriggerEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_TriggerEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_TriggerEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.Trigger"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_TriggerEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_TriggerEntityPrototype() {

	var triggerEntityProto EntityPrototype
	triggerEntityProto.TypeName = "BPMS.Trigger"
	triggerEntityProto.Ids = append(triggerEntityProto.Ids, "uuid")
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "uuid")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 0)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, false)
	triggerEntityProto.Indexs = append(triggerEntityProto.Indexs, "parentUuid")
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "parentUuid")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 1)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, false)

	/** members of Trigger **/
	triggerEntityProto.Ids = append(triggerEntityProto.Ids, "M_id")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 2)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "M_id")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "xs.ID")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 3)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "M_processUUID")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 4)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "M_sessionId")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 5)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "M_eventTriggerType")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "enum:EventTriggerType_None:EventTriggerType_Timer:EventTriggerType_Conditional:EventTriggerType_Message:EventTriggerType_Signal:EventTriggerType_Multiple:EventTriggerType_ParallelMultiple:EventTriggerType_Escalation:EventTriggerType_Error:EventTriggerType_Compensation:EventTriggerType_Terminate:EventTriggerType_Cancel:EventTriggerType_Link:EventTriggerType_Start")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 6)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "M_eventDatas")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "[]BPMS.EventData")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 7)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "M_dataRef")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "[]BPMS.ItemAwareElementInstance:Ref")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 8)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "M_sourceRef")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "BPMS.EventDefinitionInstance:Ref")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 9)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "M_targetRef")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "BPMS.FlowNodeInstance:Ref")

	/** associations of Trigger **/
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 10)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, false)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "M_runtimesPtr")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "BPMS.Runtimes:Ref")
	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "childsUuid")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "[]xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 11)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, false)

	triggerEntityProto.Fields = append(triggerEntityProto.Fields, "referenced")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType, "[]EntityRef")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder, 12)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&triggerEntityProto)

}

/** Create **/
func (this *BPMS_TriggerEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.Trigger"

	var query EntityQuery
	query.TypeName = "BPMS.Trigger"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Trigger **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_processUUID")
	query.Fields = append(query.Fields, "M_sessionId")
	query.Fields = append(query.Fields, "M_eventTriggerType")
	query.Fields = append(query.Fields, "M_eventDatas")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_sourceRef")
	query.Fields = append(query.Fields, "M_targetRef")

	/** associations of Trigger **/
	query.Fields = append(query.Fields, "M_runtimesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var TriggerInfo []interface{}

	TriggerInfo = append(TriggerInfo, this.GetUuid())
	if this.parentPtr != nil {
		TriggerInfo = append(TriggerInfo, this.parentPtr.GetUuid())
	} else {
		TriggerInfo = append(TriggerInfo, "")
	}

	/** members of Trigger **/
	TriggerInfo = append(TriggerInfo, this.object.M_id)
	TriggerInfo = append(TriggerInfo, this.object.M_processUUID)
	TriggerInfo = append(TriggerInfo, this.object.M_sessionId)

	/** Save eventTriggerType type EventTriggerType **/
	if this.object.M_eventTriggerType == BPMS.EventTriggerType_None {
		TriggerInfo = append(TriggerInfo, 0)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Timer {
		TriggerInfo = append(TriggerInfo, 1)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Conditional {
		TriggerInfo = append(TriggerInfo, 2)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Message {
		TriggerInfo = append(TriggerInfo, 3)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Signal {
		TriggerInfo = append(TriggerInfo, 4)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Multiple {
		TriggerInfo = append(TriggerInfo, 5)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_ParallelMultiple {
		TriggerInfo = append(TriggerInfo, 6)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Escalation {
		TriggerInfo = append(TriggerInfo, 7)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Error {
		TriggerInfo = append(TriggerInfo, 8)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Compensation {
		TriggerInfo = append(TriggerInfo, 9)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Terminate {
		TriggerInfo = append(TriggerInfo, 10)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Cancel {
		TriggerInfo = append(TriggerInfo, 11)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Link {
		TriggerInfo = append(TriggerInfo, 12)
	} else if this.object.M_eventTriggerType == BPMS.EventTriggerType_Start {
		TriggerInfo = append(TriggerInfo, 13)
	} else {
		TriggerInfo = append(TriggerInfo, 0)
	}

	/** Save eventDatas type EventData **/
	eventDatasIds := make([]string, 0)
	for i := 0; i < len(this.object.M_eventDatas); i++ {
		eventDatasEntity := GetServer().GetEntityManager().NewBPMSEventDataEntity(this.object.M_eventDatas[i].UUID, this.object.M_eventDatas[i])
		eventDatasIds = append(eventDatasIds, eventDatasEntity.uuid)
		eventDatasEntity.AppendReferenced("eventDatas", this)
		this.AppendChild("eventDatas", eventDatasEntity)
		if eventDatasEntity.NeedSave() {
			eventDatasEntity.SaveEntity()
		}
	}
	eventDatasStr, _ := json.Marshal(eventDatasIds)
	TriggerInfo = append(TriggerInfo, string(eventDatasStr))

	/** Save dataRef type ItemAwareElementInstance **/
	dataRefStr, _ := json.Marshal(this.object.M_dataRef)
	TriggerInfo = append(TriggerInfo, string(dataRefStr))

	/** Save sourceRef type EventDefinitionInstance **/
	TriggerInfo = append(TriggerInfo, this.object.M_sourceRef)

	/** Save targetRef type FlowNodeInstance **/
	TriggerInfo = append(TriggerInfo, this.object.M_targetRef)

	/** associations of Trigger **/

	/** Save runtimes type Runtimes **/
	TriggerInfo = append(TriggerInfo, this.object.M_runtimesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	TriggerInfo = append(TriggerInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	TriggerInfo = append(TriggerInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), TriggerInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), TriggerInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_TriggerEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_TriggerEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.Trigger"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Trigger **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_processUUID")
	query.Fields = append(query.Fields, "M_sessionId")
	query.Fields = append(query.Fields, "M_eventTriggerType")
	query.Fields = append(query.Fields, "M_eventDatas")
	query.Fields = append(query.Fields, "M_dataRef")
	query.Fields = append(query.Fields, "M_sourceRef")
	query.Fields = append(query.Fields, "M_targetRef")

	/** associations of Trigger **/
	query.Fields = append(query.Fields, "M_runtimesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Trigger...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.Trigger)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.Trigger"

		this.parentUuid = results[0][1].(string)

		/** members of Trigger **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** processUUID **/
		if results[0][3] != nil {
			this.object.M_processUUID = results[0][3].(string)
		}

		/** sessionId **/
		if results[0][4] != nil {
			this.object.M_sessionId = results[0][4].(string)
		}

		/** eventTriggerType **/
		if results[0][5] != nil {
			enumIndex := results[0][5].(int)
			if enumIndex == 0 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_None
			} else if enumIndex == 1 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Timer
			} else if enumIndex == 2 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Conditional
			} else if enumIndex == 3 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Message
			} else if enumIndex == 4 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Signal
			} else if enumIndex == 5 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Multiple
			} else if enumIndex == 6 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_ParallelMultiple
			} else if enumIndex == 7 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Escalation
			} else if enumIndex == 8 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Error
			} else if enumIndex == 9 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Compensation
			} else if enumIndex == 10 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Terminate
			} else if enumIndex == 11 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Cancel
			} else if enumIndex == 12 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Link
			} else if enumIndex == 13 {
				this.object.M_eventTriggerType = BPMS.EventTriggerType_Start
			}
		}

		/** eventDatas **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var eventDatasEntity *BPMS_EventDataEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						eventDatasEntity = instance.(*BPMS_EventDataEntity)
					} else {
						eventDatasEntity = GetServer().GetEntityManager().NewBPMSEventDataEntity(uuids[i], nil)
						eventDatasEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(eventDatasEntity)
					}
					eventDatasEntity.AppendReferenced("eventDatas", this)
					this.AppendChild("eventDatas", eventDatasEntity)
				}
			}
		}

		/** dataRef **/
		if results[0][7] != nil {
			idsStr := results[0][7].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "BPMS.ItemAwareElementInstance"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef, ids[i])
					GetServer().GetEntityManager().appendReference("dataRef", this.object.UUID, id_)
				}
			}
		}

		/** sourceRef **/
		if results[0][8] != nil {
			id := results[0][8].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.EventDefinitionInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_sourceRef = id
				GetServer().GetEntityManager().appendReference("sourceRef", this.object.UUID, id_)
			}
		}

		/** targetRef **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.FlowNodeInstance"
				id_ := refTypeName + "$$" + id
				this.object.M_targetRef = id
				GetServer().GetEntityManager().appendReference("targetRef", this.object.UUID, id_)
			}
		}

		/** associations of Trigger **/

		/** runtimesPtr **/
		if results[0][10] != nil {
			id := results[0][10].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.Runtimes"
				id_ := refTypeName + "$$" + id
				this.object.M_runtimesPtr = id
				GetServer().GetEntityManager().appendReference("runtimesPtr", this.object.UUID, id_)
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
func (this *EntityManager) NewBPMSTriggerEntityFromObject(object *BPMS.Trigger) *BPMS_TriggerEntity {
	return this.NewBPMSTriggerEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_TriggerEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSTriggerExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.Trigger"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_TriggerEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_TriggerEntity) AppendReference(reference Entity) {

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
//              			Exception
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_ExceptionEntity struct {
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
	object         *BPMS.Exception
}

/** Constructor function **/
func (this *EntityManager) NewBPMSExceptionEntity(objectId string, object interface{}) *BPMS_ExceptionEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSExceptionExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.Exception).TYPENAME = "BPMS.Exception"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.Exception).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_ExceptionEntity)
		}
	} else {
		uuidStr = "BPMS.Exception%" + Utility.RandomUUID()
	}
	entity := new(BPMS_ExceptionEntity)
	if object == nil {
		entity.object = new(BPMS.Exception)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.Exception)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.Exception"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.Exception", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_ExceptionEntity) GetTypeName() string {
	return "BPMS.Exception"
}
func (this *BPMS_ExceptionEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_ExceptionEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_ExceptionEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_ExceptionEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_ExceptionEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_ExceptionEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_ExceptionEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_ExceptionEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_ExceptionEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_ExceptionEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_ExceptionEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_ExceptionEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_ExceptionEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_ExceptionEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_ExceptionEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_ExceptionEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_ExceptionEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_ExceptionEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_ExceptionEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_ExceptionEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_ExceptionEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_ExceptionEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_ExceptionEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.Exception"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_ExceptionEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_ExceptionEntityPrototype() {

	var exceptionEntityProto EntityPrototype
	exceptionEntityProto.TypeName = "BPMS.Exception"
	exceptionEntityProto.Ids = append(exceptionEntityProto.Ids, "uuid")
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields, "uuid")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType, "xs.string")
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder, 0)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility, false)
	exceptionEntityProto.Indexs = append(exceptionEntityProto.Indexs, "parentUuid")
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields, "parentUuid")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType, "xs.string")
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder, 1)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility, false)

	/** members of Exception **/
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder, 2)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility, true)
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields, "M_id")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType, "xs.string")
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder, 3)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility, true)
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields, "M_exceptionType")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType, "enum:ExceptionType_NoIORuleException:ExceptionType_GatewayException:ExceptionType_NoAvailableOutputSetException:ExceptionType_NotMatchingIOSpecification:ExceptionType_IllegalStartEventException")

	/** associations of Exception **/
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder, 4)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility, false)
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields, "M_runtimesPtr")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType, "BPMS.Runtimes:Ref")
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields, "childsUuid")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType, "[]xs.string")
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder, 5)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility, false)

	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields, "referenced")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType, "[]EntityRef")
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder, 6)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&exceptionEntityProto)

}

/** Create **/
func (this *BPMS_ExceptionEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.Exception"

	var query EntityQuery
	query.TypeName = "BPMS.Exception"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Exception **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_exceptionType")

	/** associations of Exception **/
	query.Fields = append(query.Fields, "M_runtimesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ExceptionInfo []interface{}

	ExceptionInfo = append(ExceptionInfo, this.GetUuid())
	if this.parentPtr != nil {
		ExceptionInfo = append(ExceptionInfo, this.parentPtr.GetUuid())
	} else {
		ExceptionInfo = append(ExceptionInfo, "")
	}

	/** members of Exception **/
	ExceptionInfo = append(ExceptionInfo, this.object.M_id)

	/** Save exceptionType type ExceptionType **/
	if this.object.M_exceptionType == BPMS.ExceptionType_NoIORuleException {
		ExceptionInfo = append(ExceptionInfo, 0)
	} else if this.object.M_exceptionType == BPMS.ExceptionType_GatewayException {
		ExceptionInfo = append(ExceptionInfo, 1)
	} else if this.object.M_exceptionType == BPMS.ExceptionType_NoAvailableOutputSetException {
		ExceptionInfo = append(ExceptionInfo, 2)
	} else if this.object.M_exceptionType == BPMS.ExceptionType_NotMatchingIOSpecification {
		ExceptionInfo = append(ExceptionInfo, 3)
	} else if this.object.M_exceptionType == BPMS.ExceptionType_IllegalStartEventException {
		ExceptionInfo = append(ExceptionInfo, 4)
	} else {
		ExceptionInfo = append(ExceptionInfo, 0)
	}

	/** associations of Exception **/

	/** Save runtimes type Runtimes **/
	ExceptionInfo = append(ExceptionInfo, this.object.M_runtimesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	ExceptionInfo = append(ExceptionInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	ExceptionInfo = append(ExceptionInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), ExceptionInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), ExceptionInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_ExceptionEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_ExceptionEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.Exception"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Exception **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_exceptionType")

	/** associations of Exception **/
	query.Fields = append(query.Fields, "M_runtimesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Exception...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.Exception)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.Exception"

		this.parentUuid = results[0][1].(string)

		/** members of Exception **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** exceptionType **/
		if results[0][3] != nil {
			enumIndex := results[0][3].(int)
			if enumIndex == 0 {
				this.object.M_exceptionType = BPMS.ExceptionType_NoIORuleException
			} else if enumIndex == 1 {
				this.object.M_exceptionType = BPMS.ExceptionType_GatewayException
			} else if enumIndex == 2 {
				this.object.M_exceptionType = BPMS.ExceptionType_NoAvailableOutputSetException
			} else if enumIndex == 3 {
				this.object.M_exceptionType = BPMS.ExceptionType_NotMatchingIOSpecification
			} else if enumIndex == 4 {
				this.object.M_exceptionType = BPMS.ExceptionType_IllegalStartEventException
			}
		}

		/** associations of Exception **/

		/** runtimesPtr **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.Runtimes"
				id_ := refTypeName + "$$" + id
				this.object.M_runtimesPtr = id
				GetServer().GetEntityManager().appendReference("runtimesPtr", this.object.UUID, id_)
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
func (this *EntityManager) NewBPMSExceptionEntityFromObject(object *BPMS.Exception) *BPMS_ExceptionEntity {
	return this.NewBPMSExceptionEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_ExceptionEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSExceptionExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.Exception"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_ExceptionEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_ExceptionEntity) AppendReference(reference Entity) {

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
//              			LogInfo
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_LogInfoEntity struct {
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
	object         *BPMS.LogInfo
}

/** Constructor function **/
func (this *EntityManager) NewBPMSLogInfoEntity(objectId string, object interface{}) *BPMS_LogInfoEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSLogInfoExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.LogInfo).TYPENAME = "BPMS.LogInfo"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.LogInfo).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_LogInfoEntity)
		}
	} else {
		uuidStr = "BPMS.LogInfo%" + Utility.RandomUUID()
	}
	entity := new(BPMS_LogInfoEntity)
	if object == nil {
		entity.object = new(BPMS.LogInfo)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.LogInfo)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.LogInfo"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.LogInfo", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_LogInfoEntity) GetTypeName() string {
	return "BPMS.LogInfo"
}
func (this *BPMS_LogInfoEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_LogInfoEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_LogInfoEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_LogInfoEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_LogInfoEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_LogInfoEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_LogInfoEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_LogInfoEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_LogInfoEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_LogInfoEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_LogInfoEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_LogInfoEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_LogInfoEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_LogInfoEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_LogInfoEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_LogInfoEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_LogInfoEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_LogInfoEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_LogInfoEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_LogInfoEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_LogInfoEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_LogInfoEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_LogInfoEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.LogInfo"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_LogInfoEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_LogInfoEntityPrototype() {

	var logInfoEntityProto EntityPrototype
	logInfoEntityProto.TypeName = "BPMS.LogInfo"
	logInfoEntityProto.Ids = append(logInfoEntityProto.Ids, "uuid")
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "uuid")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "xs.string")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 0)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, false)
	logInfoEntityProto.Indexs = append(logInfoEntityProto.Indexs, "parentUuid")
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "parentUuid")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "xs.string")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 1)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, false)

	/** members of LogInfo **/
	logInfoEntityProto.Ids = append(logInfoEntityProto.Ids, "M_id")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 2)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "M_id")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "xs.ID")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 3)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "M_date")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "xs.int64")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 4)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "M_actor")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "xs.interface{}:Ref")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 5)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "M_action")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "xs.string")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 6)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "M_object")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "BPMS.Instance:Ref")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 7)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "M_description")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "xs.string")

	/** associations of LogInfo **/
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 8)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, false)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "M_runtimesPtr")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "BPMS.Runtimes:Ref")
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "childsUuid")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "[]xs.string")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 9)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, false)

	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields, "referenced")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType, "[]EntityRef")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder, 10)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&logInfoEntityProto)

}

/** Create **/
func (this *BPMS_LogInfoEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.LogInfo"

	var query EntityQuery
	query.TypeName = "BPMS.LogInfo"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of LogInfo **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_date")
	query.Fields = append(query.Fields, "M_actor")
	query.Fields = append(query.Fields, "M_action")
	query.Fields = append(query.Fields, "M_object")
	query.Fields = append(query.Fields, "M_description")

	/** associations of LogInfo **/
	query.Fields = append(query.Fields, "M_runtimesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var LogInfoInfo []interface{}

	LogInfoInfo = append(LogInfoInfo, this.GetUuid())
	if this.parentPtr != nil {
		LogInfoInfo = append(LogInfoInfo, this.parentPtr.GetUuid())
	} else {
		LogInfoInfo = append(LogInfoInfo, "")
	}

	/** members of LogInfo **/
	LogInfoInfo = append(LogInfoInfo, this.object.M_id)
	LogInfoInfo = append(LogInfoInfo, this.object.M_date)
	LogInfoInfo = append(LogInfoInfo, this.object.M_actor)
	LogInfoInfo = append(LogInfoInfo, this.object.M_action)

	/** Save object type Instance **/
	LogInfoInfo = append(LogInfoInfo, this.object.M_object)
	LogInfoInfo = append(LogInfoInfo, this.object.M_description)

	/** associations of LogInfo **/

	/** Save runtimes type Runtimes **/
	LogInfoInfo = append(LogInfoInfo, this.object.M_runtimesPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	LogInfoInfo = append(LogInfoInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	LogInfoInfo = append(LogInfoInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), LogInfoInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), LogInfoInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_LogInfoEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_LogInfoEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.LogInfo"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of LogInfo **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_date")
	query.Fields = append(query.Fields, "M_actor")
	query.Fields = append(query.Fields, "M_action")
	query.Fields = append(query.Fields, "M_object")
	query.Fields = append(query.Fields, "M_description")

	/** associations of LogInfo **/
	query.Fields = append(query.Fields, "M_runtimesPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of LogInfo...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.LogInfo)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.LogInfo"

		this.parentUuid = results[0][1].(string)

		/** members of LogInfo **/

		/** id **/
		if results[0][2] != nil {
			this.object.M_id = results[0][2].(string)
		}

		/** date **/
		if results[0][3] != nil {
			this.object.M_date = results[0][3].(int64)
		}

		/** actor **/
		if results[0][4] != nil {
			id := results[0][4].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.interface{}"
				id_ := refTypeName + "$$" + id
				this.object.M_actor = id
				GetServer().GetEntityManager().appendReference("actor", this.object.UUID, id_)
			}
		}

		/** action **/
		if results[0][5] != nil {
			this.object.M_action = results[0][5].(string)
		}

		/** object **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.Instance"
				id_ := refTypeName + "$$" + id
				this.object.M_object = id
				GetServer().GetEntityManager().appendReference("object", this.object.UUID, id_)
			}
		}

		/** description **/
		if results[0][7] != nil {
			this.object.M_description = results[0][7].(string)
		}

		/** associations of LogInfo **/

		/** runtimesPtr **/
		if results[0][8] != nil {
			id := results[0][8].(string)
			if len(id) > 0 {
				refTypeName := "BPMS.Runtimes"
				id_ := refTypeName + "$$" + id
				this.object.M_runtimesPtr = id
				GetServer().GetEntityManager().appendReference("runtimesPtr", this.object.UUID, id_)
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
func (this *EntityManager) NewBPMSLogInfoEntityFromObject(object *BPMS.LogInfo) *BPMS_LogInfoEntity {
	return this.NewBPMSLogInfoEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_LogInfoEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSLogInfoExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.LogInfo"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_LogInfoEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_LogInfoEntity) AppendReference(reference Entity) {

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
//              			Runtimes
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_RuntimesEntity struct {
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
	object         *BPMS.Runtimes
}

/** Constructor function **/
func (this *EntityManager) NewBPMSRuntimesEntity(objectId string, object interface{}) *BPMS_RuntimesEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = BPMSRuntimesExists(objectId)
		}
	}
	if object != nil {
		object.(*BPMS.Runtimes).TYPENAME = "BPMS.Runtimes"
	}
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*BPMS.Runtimes).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_RuntimesEntity)
		}
	} else {
		uuidStr = "BPMS.Runtimes%" + Utility.RandomUUID()
	}
	entity := new(BPMS_RuntimesEntity)
	if object == nil {
		entity.object = new(BPMS.Runtimes)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*BPMS.Runtimes)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS.Runtimes"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS.Runtimes", "BPMS")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func (this *BPMS_RuntimesEntity) GetTypeName() string {
	return "BPMS.Runtimes"
}
func (this *BPMS_RuntimesEntity) GetUuid() string {
	return this.uuid
}
func (this *BPMS_RuntimesEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *BPMS_RuntimesEntity) SetParentPtr(parentPtr Entity) {
	this.parentPtr = parentPtr
}

func (this *BPMS_RuntimesEntity) AppendReferenced(name string, owner Entity) {
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

func (this *BPMS_RuntimesEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *BPMS_RuntimesEntity) RemoveReferenced(name string, owner Entity) {
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

func (this *BPMS_RuntimesEntity) RemoveReference(name string, reference Entity) {
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

func (this *BPMS_RuntimesEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

func (this *BPMS_RuntimesEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

func (this *BPMS_RuntimesEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *BPMS_RuntimesEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *BPMS_RuntimesEntity) RemoveChild(name string, uuid string) {
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

func (this *BPMS_RuntimesEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *BPMS_RuntimesEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *BPMS_RuntimesEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

func (this *BPMS_RuntimesEntity) SetReferencesPtr(refsPtr []Entity) {
	this.referencesPtr = refsPtr
}

func (this *BPMS_RuntimesEntity) GetObject() interface{} {
	return this.object
}

func (this *BPMS_RuntimesEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *BPMS_RuntimesEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *BPMS_RuntimesEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *BPMS_RuntimesEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *BPMS_RuntimesEntity) GetChecksum() string {
	objectStr, _ := json.Marshal(this.object)
	return Utility.GetMD5Hash(string(objectStr))
}

func (this *BPMS_RuntimesEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "BPMS.Runtimes"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *BPMS_RuntimesEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_BPMS_RuntimesEntityPrototype() {

	var runtimesEntityProto EntityPrototype
	runtimesEntityProto.TypeName = "BPMS.Runtimes"
	runtimesEntityProto.Ids = append(runtimesEntityProto.Ids, "uuid")
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "uuid")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 0)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, false)
	runtimesEntityProto.Indexs = append(runtimesEntityProto.Indexs, "parentUuid")
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "parentUuid")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 1)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, false)

	/** members of Runtimes **/
	runtimesEntityProto.Ids = append(runtimesEntityProto.Ids, "M_id")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 2)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "M_id")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "xs.ID")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 3)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "M_name")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 4)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "M_version")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 5)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "M_definitions")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "[]BPMS.DefinitionsInstance")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 6)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "M_exceptions")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "[]BPMS.Exception")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 7)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "M_triggers")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "[]BPMS.Trigger")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 8)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "M_correlationInfos")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "[]BPMS.CorrelationInfo")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 9)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "M_logInfos")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "[]BPMS.LogInfo")
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "childsUuid")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "[]xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 10)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, false)

	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields, "referenced")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType, "[]EntityRef")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder, 11)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(BPMSDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&runtimesEntityProto)

}

/** Create **/
func (this *BPMS_RuntimesEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS.Runtimes"

	var query EntityQuery
	query.TypeName = "BPMS.Runtimes"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Runtimes **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_version")
	query.Fields = append(query.Fields, "M_definitions")
	query.Fields = append(query.Fields, "M_exceptions")
	query.Fields = append(query.Fields, "M_triggers")
	query.Fields = append(query.Fields, "M_correlationInfos")
	query.Fields = append(query.Fields, "M_logInfos")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var RuntimesInfo []interface{}

	RuntimesInfo = append(RuntimesInfo, this.GetUuid())
	if this.parentPtr != nil {
		RuntimesInfo = append(RuntimesInfo, this.parentPtr.GetUuid())
	} else {
		RuntimesInfo = append(RuntimesInfo, "")
	}

	/** members of Runtimes **/
	RuntimesInfo = append(RuntimesInfo, this.object.M_id)
	RuntimesInfo = append(RuntimesInfo, this.object.M_name)
	RuntimesInfo = append(RuntimesInfo, this.object.M_version)

	/** Save definitions type DefinitionsInstance **/
	definitionsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_definitions); i++ {
		definitionsEntity := GetServer().GetEntityManager().NewBPMSDefinitionsInstanceEntity(this.object.M_definitions[i].UUID, this.object.M_definitions[i])
		definitionsIds = append(definitionsIds, definitionsEntity.uuid)
		definitionsEntity.AppendReferenced("definitions", this)
		this.AppendChild("definitions", definitionsEntity)
		if definitionsEntity.NeedSave() {
			definitionsEntity.SaveEntity()
		}
	}
	definitionsStr, _ := json.Marshal(definitionsIds)
	RuntimesInfo = append(RuntimesInfo, string(definitionsStr))

	/** Save exceptions type Exception **/
	exceptionsIds := make([]string, 0)
	for i := 0; i < len(this.object.M_exceptions); i++ {
		exceptionsEntity := GetServer().GetEntityManager().NewBPMSExceptionEntity(this.object.M_exceptions[i].UUID, this.object.M_exceptions[i])
		exceptionsIds = append(exceptionsIds, exceptionsEntity.uuid)
		exceptionsEntity.AppendReferenced("exceptions", this)
		this.AppendChild("exceptions", exceptionsEntity)
		if exceptionsEntity.NeedSave() {
			exceptionsEntity.SaveEntity()
		}
	}
	exceptionsStr, _ := json.Marshal(exceptionsIds)
	RuntimesInfo = append(RuntimesInfo, string(exceptionsStr))

	/** Save triggers type Trigger **/
	triggersIds := make([]string, 0)
	for i := 0; i < len(this.object.M_triggers); i++ {
		triggersEntity := GetServer().GetEntityManager().NewBPMSTriggerEntity(this.object.M_triggers[i].UUID, this.object.M_triggers[i])
		triggersIds = append(triggersIds, triggersEntity.uuid)
		triggersEntity.AppendReferenced("triggers", this)
		this.AppendChild("triggers", triggersEntity)
		if triggersEntity.NeedSave() {
			triggersEntity.SaveEntity()
		}
	}
	triggersStr, _ := json.Marshal(triggersIds)
	RuntimesInfo = append(RuntimesInfo, string(triggersStr))

	/** Save correlationInfos type CorrelationInfo **/
	correlationInfosIds := make([]string, 0)
	for i := 0; i < len(this.object.M_correlationInfos); i++ {
		correlationInfosEntity := GetServer().GetEntityManager().NewBPMSCorrelationInfoEntity(this.object.M_correlationInfos[i].UUID, this.object.M_correlationInfos[i])
		correlationInfosIds = append(correlationInfosIds, correlationInfosEntity.uuid)
		correlationInfosEntity.AppendReferenced("correlationInfos", this)
		this.AppendChild("correlationInfos", correlationInfosEntity)
		if correlationInfosEntity.NeedSave() {
			correlationInfosEntity.SaveEntity()
		}
	}
	correlationInfosStr, _ := json.Marshal(correlationInfosIds)
	RuntimesInfo = append(RuntimesInfo, string(correlationInfosStr))

	/** Save logInfos type LogInfo **/
	logInfosIds := make([]string, 0)
	for i := 0; i < len(this.object.M_logInfos); i++ {
		logInfosEntity := GetServer().GetEntityManager().NewBPMSLogInfoEntity(this.object.M_logInfos[i].UUID, this.object.M_logInfos[i])
		logInfosIds = append(logInfosIds, logInfosEntity.uuid)
		logInfosEntity.AppendReferenced("logInfos", this)
		this.AppendChild("logInfos", logInfosEntity)
		if logInfosEntity.NeedSave() {
			logInfosEntity.SaveEntity()
		}
	}
	logInfosStr, _ := json.Marshal(logInfosIds)
	RuntimesInfo = append(RuntimesInfo, string(logInfosStr))
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	RuntimesInfo = append(RuntimesInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	RuntimesInfo = append(RuntimesInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMSDB, string(queryStr), RuntimesInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(BPMSDB, string(queryStr), RuntimesInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_RuntimesEntity) InitEntity(id string) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_RuntimesEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS.Runtimes"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Runtimes **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_version")
	query.Fields = append(query.Fields, "M_definitions")
	query.Fields = append(query.Fields, "M_exceptions")
	query.Fields = append(query.Fields, "M_triggers")
	query.Fields = append(query.Fields, "M_correlationInfos")
	query.Fields = append(query.Fields, "M_logInfos")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Runtimes...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(BPMS.Runtimes)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS.Runtimes"

		this.parentUuid = results[0][1].(string)

		/** members of Runtimes **/

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

		/** definitions **/
		if results[0][5] != nil {
			uuidsStr := results[0][5].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var definitionsEntity *BPMS_DefinitionsInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						definitionsEntity = instance.(*BPMS_DefinitionsInstanceEntity)
					} else {
						definitionsEntity = GetServer().GetEntityManager().NewBPMSDefinitionsInstanceEntity(uuids[i], nil)
						definitionsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(definitionsEntity)
					}
					definitionsEntity.AppendReferenced("definitions", this)
					this.AppendChild("definitions", definitionsEntity)
				}
			}
		}

		/** exceptions **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var exceptionsEntity *BPMS_ExceptionEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						exceptionsEntity = instance.(*BPMS_ExceptionEntity)
					} else {
						exceptionsEntity = GetServer().GetEntityManager().NewBPMSExceptionEntity(uuids[i], nil)
						exceptionsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(exceptionsEntity)
					}
					exceptionsEntity.AppendReferenced("exceptions", this)
					this.AppendChild("exceptions", exceptionsEntity)
				}
			}
		}

		/** triggers **/
		if results[0][7] != nil {
			uuidsStr := results[0][7].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var triggersEntity *BPMS_TriggerEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						triggersEntity = instance.(*BPMS_TriggerEntity)
					} else {
						triggersEntity = GetServer().GetEntityManager().NewBPMSTriggerEntity(uuids[i], nil)
						triggersEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(triggersEntity)
					}
					triggersEntity.AppendReferenced("triggers", this)
					this.AppendChild("triggers", triggersEntity)
				}
			}
		}

		/** correlationInfos **/
		if results[0][8] != nil {
			uuidsStr := results[0][8].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var correlationInfosEntity *BPMS_CorrelationInfoEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						correlationInfosEntity = instance.(*BPMS_CorrelationInfoEntity)
					} else {
						correlationInfosEntity = GetServer().GetEntityManager().NewBPMSCorrelationInfoEntity(uuids[i], nil)
						correlationInfosEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(correlationInfosEntity)
					}
					correlationInfosEntity.AppendReferenced("correlationInfos", this)
					this.AppendChild("correlationInfos", correlationInfosEntity)
				}
			}
		}

		/** logInfos **/
		if results[0][9] != nil {
			uuidsStr := results[0][9].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if len(uuids[i]) > 0 {
					var logInfosEntity *BPMS_LogInfoEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						logInfosEntity = instance.(*BPMS_LogInfoEntity)
					} else {
						logInfosEntity = GetServer().GetEntityManager().NewBPMSLogInfoEntity(uuids[i], nil)
						logInfosEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(logInfosEntity)
					}
					logInfosEntity.AppendReferenced("logInfos", this)
					this.AppendChild("logInfos", logInfosEntity)
				}
			}
		}
		childsUuidStr := results[0][10].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][11].(string)
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
func (this *EntityManager) NewBPMSRuntimesEntityFromObject(object *BPMS.Runtimes) *BPMS_RuntimesEntity {
	return this.NewBPMSRuntimesEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_RuntimesEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMSRuntimesExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS.Runtimes"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMSDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_RuntimesEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMS_RuntimesEntity) AppendReference(reference Entity) {

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
func (this *EntityManager) registerBPMSObjects() {
	Utility.RegisterType((*BPMS.ConnectingObject)(nil))
	Utility.RegisterType((*BPMS.ActivityInstance)(nil))
	Utility.RegisterType((*BPMS.SubprocessInstance)(nil))
	Utility.RegisterType((*BPMS.GatewayInstance)(nil))
	Utility.RegisterType((*BPMS.EventInstance)(nil))
	Utility.RegisterType((*BPMS.DefinitionsInstance)(nil))
	Utility.RegisterType((*BPMS.ProcessInstance)(nil))
	Utility.RegisterType((*BPMS.EventDefinitionInstance)(nil))
	Utility.RegisterType((*BPMS.CorrelationInfo)(nil))
	Utility.RegisterType((*BPMS.ItemAwareElementInstance)(nil))
	Utility.RegisterType((*BPMS.RessourceParameterInstance)(nil))
	Utility.RegisterType((*BPMS.RessourceInstance)(nil))
	Utility.RegisterType((*BPMS.EventData)(nil))
	Utility.RegisterType((*BPMS.Trigger)(nil))
	Utility.RegisterType((*BPMS.Exception)(nil))
	Utility.RegisterType((*BPMS.LogInfo)(nil))
	Utility.RegisterType((*BPMS.Runtimes)(nil))
}

/** Create entity prototypes contain in a package **/
func (this *EntityManager) createBPMSPrototypes() {
	this.create_BPMS_InstanceEntityPrototype()
	this.create_BPMS_ConnectingObjectEntityPrototype()
	this.create_BPMS_FlowNodeInstanceEntityPrototype()
	this.create_BPMS_ActivityInstanceEntityPrototype()
	this.create_BPMS_SubprocessInstanceEntityPrototype()
	this.create_BPMS_GatewayInstanceEntityPrototype()
	this.create_BPMS_EventInstanceEntityPrototype()
	this.create_BPMS_DefinitionsInstanceEntityPrototype()
	this.create_BPMS_ProcessInstanceEntityPrototype()
	this.create_BPMS_EventDefinitionInstanceEntityPrototype()
	this.create_BPMS_CorrelationInfoEntityPrototype()
	this.create_BPMS_ItemAwareElementInstanceEntityPrototype()
	this.create_BPMS_RessourceParameterInstanceEntityPrototype()
	this.create_BPMS_RessourceInstanceEntityPrototype()
	this.create_BPMS_EventDataEntityPrototype()
	this.create_BPMS_TriggerEntityPrototype()
	this.create_BPMS_ExceptionEntityPrototype()
	this.create_BPMS_LogInfoEntityPrototype()
	this.create_BPMS_RuntimesEntityPrototype()
}
