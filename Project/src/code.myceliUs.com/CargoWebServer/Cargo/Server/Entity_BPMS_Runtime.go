package Server
import(
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/BPMS_Runtime"
	"encoding/json"
	"code.google.com/p/go-uuid/uuid"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
	"log"
	"strings"
)

/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_InstanceEntityPrototype() {

	var instanceEntityProto EntityPrototype
	instanceEntityProto.TypeName = "BPMS_Runtime.Instance"
	instanceEntityProto.IsAbstract=true
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS_Runtime.ProcessInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS_Runtime.EventDefinitionInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS_Runtime.DefinitionsInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS_Runtime.ActivityInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS_Runtime.SubprocessInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS_Runtime.GatewayInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS_Runtime.EventInstance")
	instanceEntityProto.SubstitutionGroup = append(instanceEntityProto.SubstitutionGroup, "BPMS_Runtime.ConnectingObject")
	instanceEntityProto.Ids = append(instanceEntityProto.Ids,"uuid")
	instanceEntityProto.Fields = append(instanceEntityProto.Fields,"uuid")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType,"xs.string")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder,0)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility,false)
	instanceEntityProto.Indexs = append(instanceEntityProto.Indexs,"parentUuid")
	instanceEntityProto.Fields = append(instanceEntityProto.Fields,"parentUuid")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType,"xs.string")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder,1)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility,false)

	/** members of Instance **/
	instanceEntityProto.Ids = append(instanceEntityProto.Ids,"M_id")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder,2)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility,true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields,"M_id")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType,"xs.ID")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder,3)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility,true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields,"M_bpmnElementId")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType,"xs.string")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder,4)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility,true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields,"M_participants")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType,"[]xs.string")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder,5)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility,true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields,"M_dataRef")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder,6)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility,true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields,"M_data")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder,7)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility,true)
	instanceEntityProto.Fields = append(instanceEntityProto.Fields,"M_logInfoRef")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo:Ref")
	instanceEntityProto.Fields = append(instanceEntityProto.Fields,"childsUuid")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType,"[]xs.string")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder,8)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility,false)

	instanceEntityProto.Fields = append(instanceEntityProto.Fields,"referenced")
	instanceEntityProto.FieldsType = append(instanceEntityProto.FieldsType,"[]EntityRef")
	instanceEntityProto.FieldsOrder = append(instanceEntityProto.FieldsOrder,9)
	instanceEntityProto.FieldsVisibility = append(instanceEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&instanceEntityProto)

}


////////////////////////////////////////////////////////////////////////////////
//              			ConnectingObject
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_ConnectingObjectEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.ConnectingObject
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeConnectingObjectEntity(objectId string, object interface{}) *BPMS_Runtime_ConnectingObjectEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeConnectingObjectExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.ConnectingObject).TYPENAME = "BPMS_Runtime.ConnectingObject"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.ConnectingObject).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_ConnectingObjectEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.ConnectingObject%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_ConnectingObjectEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.ConnectingObject)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.ConnectingObject)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.ConnectingObject"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.ConnectingObject","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_ConnectingObjectEntity) GetTypeName()string{
	return "BPMS_Runtime.ConnectingObject"
}
func(this *BPMS_Runtime_ConnectingObjectEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_ConnectingObjectEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_ConnectingObjectEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_ConnectingObjectEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_ConnectingObjectEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_ConnectingObjectEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_ConnectingObjectEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_ConnectingObjectEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_ConnectingObjectEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_ConnectingObjectEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_ConnectingObjectEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_ConnectingObjectEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_ConnectingObjectEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_ConnectingObjectEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_ConnectingObjectEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_ConnectingObjectEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_ConnectingObjectEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_ConnectingObjectEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_ConnectingObjectEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_ConnectingObjectEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_ConnectingObjectEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_ConnectingObjectEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_ConnectingObjectEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ConnectingObject"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_ConnectingObjectEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_ConnectingObjectEntityPrototype() {

	var connectingObjectEntityProto EntityPrototype
	connectingObjectEntityProto.TypeName = "BPMS_Runtime.ConnectingObject"
	connectingObjectEntityProto.SuperTypeNames = append(connectingObjectEntityProto.SuperTypeNames, "BPMS_Runtime.Instance")
	connectingObjectEntityProto.Ids = append(connectingObjectEntityProto.Ids,"uuid")
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"uuid")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"xs.string")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,0)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,false)
	connectingObjectEntityProto.Indexs = append(connectingObjectEntityProto.Indexs,"parentUuid")
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"parentUuid")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"xs.string")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,1)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,false)

	/** members of Instance **/
	connectingObjectEntityProto.Ids = append(connectingObjectEntityProto.Ids,"M_id")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,2)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_id")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"xs.ID")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,3)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_bpmnElementId")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"xs.string")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,4)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_participants")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"[]xs.string")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,5)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_dataRef")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,6)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_data")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,7)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_logInfoRef")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo:Ref")

	/** members of ConnectingObject **/
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,8)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_connectingObjectType")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"enum:ConnectingObjectType_SequenceFlow:ConnectingObjectType_MessageFlow:ConnectingObjectType_Association:ConnectingObjectType_DataAssociation")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,9)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_sourceRef")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"BPMS_Runtime.FlowNodeInstance:Ref")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,10)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,true)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_targetRef")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"BPMS_Runtime.FlowNodeInstance:Ref")

	/** associations of ConnectingObject **/
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,11)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,false)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_SubprocessInstancePtr")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"BPMS_Runtime.SubprocessInstance:Ref")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,12)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,false)
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"M_processInstancePtr")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"BPMS_Runtime.ProcessInstance:Ref")
	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"childsUuid")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"[]xs.string")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,13)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,false)

	connectingObjectEntityProto.Fields = append(connectingObjectEntityProto.Fields,"referenced")
	connectingObjectEntityProto.FieldsType = append(connectingObjectEntityProto.FieldsType,"[]EntityRef")
	connectingObjectEntityProto.FieldsOrder = append(connectingObjectEntityProto.FieldsOrder,14)
	connectingObjectEntityProto.FieldsVisibility = append(connectingObjectEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&connectingObjectEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_ConnectingObjectEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.ConnectingObject"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ConnectingObject"

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
	}else{
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
	dataIds := make([]string,0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(this.object.M_data[i].UUID,this.object.M_data[i])
		dataIds=append(dataIds,dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data",dataEntity)
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
	if this.object.M_connectingObjectType==BPMS_Runtime.ConnectingObjectType_SequenceFlow{
		ConnectingObjectInfo = append(ConnectingObjectInfo, 0)
	} else if this.object.M_connectingObjectType==BPMS_Runtime.ConnectingObjectType_MessageFlow{
		ConnectingObjectInfo = append(ConnectingObjectInfo, 1)
	} else if this.object.M_connectingObjectType==BPMS_Runtime.ConnectingObjectType_Association{
		ConnectingObjectInfo = append(ConnectingObjectInfo, 2)
	} else if this.object.M_connectingObjectType==BPMS_Runtime.ConnectingObjectType_DataAssociation{
		ConnectingObjectInfo = append(ConnectingObjectInfo, 3)
	}else{
		ConnectingObjectInfo = append(ConnectingObjectInfo, 0)
	}

	/** Save sourceRef type FlowNodeInstance **/
		ConnectingObjectInfo = append(ConnectingObjectInfo,this.object.M_sourceRef)

	/** Save targetRef type FlowNodeInstance **/
		ConnectingObjectInfo = append(ConnectingObjectInfo,this.object.M_targetRef)

	/** associations of ConnectingObject **/

	/** Save SubprocessInstance type SubprocessInstance **/
		ConnectingObjectInfo = append(ConnectingObjectInfo,this.object.M_SubprocessInstancePtr)

	/** Save processInstance type ProcessInstance **/
		ConnectingObjectInfo = append(ConnectingObjectInfo,this.object.M_processInstancePtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), ConnectingObjectInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), ConnectingObjectInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_ConnectingObjectEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_ConnectingObjectEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ConnectingObject"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ConnectingObject...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.ConnectingObject)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.ConnectingObject"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** bpmnElementId **/
 		if results[0][3] != nil{
 			this.object.M_bpmnElementId=results[0][3].(string)
 		}

		/** participants **/
 		if results[0][4] != nil{
 			this.object.M_participants= append(this.object.M_participants, results[0][4].([]string) ...)
 		}

		/** dataRef **/
 		if results[0][5] != nil{
			idsStr :=results[0][5].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ItemAwareElementInstance"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef,ids[i])
					GetServer().GetEntityManager().appendReference("dataRef",this.object.UUID, id_)
				}
			}
 		}

		/** data **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_Runtime_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_Runtime_ItemAwareElementInstanceEntity)
					}else{
						dataEntity = GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data",dataEntity)
				}
 			}
 		}

		/** logInfoRef **/
 		if results[0][7] != nil{
			idsStr :=results[0][7].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.LogInfo"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef,ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of ConnectingObject **/

		/** connectingObjectType **/
 		if results[0][8] != nil{
 			enumIndex := results[0][8].(int)
			if enumIndex == 0{
 				this.object.M_connectingObjectType=BPMS_Runtime.ConnectingObjectType_SequenceFlow
			} else if enumIndex == 1{
 				this.object.M_connectingObjectType=BPMS_Runtime.ConnectingObjectType_MessageFlow
			} else if enumIndex == 2{
 				this.object.M_connectingObjectType=BPMS_Runtime.ConnectingObjectType_Association
			} else if enumIndex == 3{
 				this.object.M_connectingObjectType=BPMS_Runtime.ConnectingObjectType_DataAssociation
 			}
 		}

		/** sourceRef **/
 		if results[0][9] != nil{
			id :=results[0][9].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.FlowNodeInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_sourceRef= id
				GetServer().GetEntityManager().appendReference("sourceRef",this.object.UUID, id_)
			}
 		}

		/** targetRef **/
 		if results[0][10] != nil{
			id :=results[0][10].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.FlowNodeInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_targetRef= id
				GetServer().GetEntityManager().appendReference("targetRef",this.object.UUID, id_)
			}
 		}

		/** associations of ConnectingObject **/

		/** SubprocessInstancePtr **/
 		if results[0][11] != nil{
			id :=results[0][11].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.SubprocessInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_SubprocessInstancePtr= id
				GetServer().GetEntityManager().appendReference("SubprocessInstancePtr",this.object.UUID, id_)
			}
 		}

		/** processInstancePtr **/
 		if results[0][12] != nil{
			id :=results[0][12].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.ProcessInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_processInstancePtr= id
				GetServer().GetEntityManager().appendReference("processInstancePtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeConnectingObjectEntityFromObject(object *BPMS_Runtime.ConnectingObject) *BPMS_Runtime_ConnectingObjectEntity {
	 return this.NewBPMS_RuntimeConnectingObjectEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_ConnectingObjectEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeConnectingObjectExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ConnectingObject"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_ConnectingObjectEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_ConnectingObjectEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_FlowNodeInstanceEntityPrototype() {

	var flowNodeInstanceEntityProto EntityPrototype
	flowNodeInstanceEntityProto.TypeName = "BPMS_Runtime.FlowNodeInstance"
	flowNodeInstanceEntityProto.IsAbstract=true
	flowNodeInstanceEntityProto.SuperTypeNames = append(flowNodeInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.Instance")
	flowNodeInstanceEntityProto.SubstitutionGroup = append(flowNodeInstanceEntityProto.SubstitutionGroup, "BPMS_Runtime.SubprocessInstance")
	flowNodeInstanceEntityProto.SubstitutionGroup = append(flowNodeInstanceEntityProto.SubstitutionGroup, "BPMS_Runtime.GatewayInstance")
	flowNodeInstanceEntityProto.SubstitutionGroup = append(flowNodeInstanceEntityProto.SubstitutionGroup, "BPMS_Runtime.EventInstance")
	flowNodeInstanceEntityProto.SubstitutionGroup = append(flowNodeInstanceEntityProto.SubstitutionGroup, "BPMS_Runtime.ActivityInstance")
	flowNodeInstanceEntityProto.Ids = append(flowNodeInstanceEntityProto.Ids,"uuid")
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"uuid")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"xs.string")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,0)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,false)
	flowNodeInstanceEntityProto.Indexs = append(flowNodeInstanceEntityProto.Indexs,"parentUuid")
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"parentUuid")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"xs.string")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,1)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,false)

	/** members of Instance **/
	flowNodeInstanceEntityProto.Ids = append(flowNodeInstanceEntityProto.Ids,"M_id")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,2)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_id")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"xs.ID")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,3)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_bpmnElementId")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"xs.string")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,4)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_participants")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"[]xs.string")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,5)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_dataRef")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,6)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_data")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,7)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_logInfoRef")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo:Ref")

	/** members of FlowNodeInstance **/
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,8)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_flowNodeType")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"enum:FlowNodeType_AbstractTask:FlowNodeType_ServiceTask:FlowNodeType_UserTask:FlowNodeType_ManualTask:FlowNodeType_BusinessRuleTask:FlowNodeType_ScriptTask:FlowNodeType_EmbeddedSubprocess:FlowNodeType_EventSubprocess:FlowNodeType_AdHocSubprocess:FlowNodeType_Transaction:FlowNodeType_CallActivity:FlowNodeType_ParallelGateway:FlowNodeType_ExclusiveGateway:FlowNodeType_InclusiveGateway:FlowNodeType_EventBasedGateway:FlowNodeType_ComplexGateway:FlowNodeType_StartEvent:FlowNodeType_IntermediateCatchEvent:FlowNodeType_BoundaryEvent:FlowNodeType_EndEvent:FlowNodeType_IntermediateThrowEvent")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,9)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_lifecycleState")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"enum:LifecycleState_Completed:LifecycleState_Compensated:LifecycleState_Failed:LifecycleState_Terminated:LifecycleState_Completing:LifecycleState_Compensating:LifecycleState_Failing:LifecycleState_Terminating:LifecycleState_Ready:LifecycleState_Active")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,10)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_inputRef")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject:Ref")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,11)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,true)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_outputRef")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject:Ref")

	/** associations of FlowNodeInstance **/
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,12)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,false)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_SubprocessInstancePtr")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"BPMS_Runtime.SubprocessInstance:Ref")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,13)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,false)
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"M_processInstancePtr")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"BPMS_Runtime.ProcessInstance:Ref")
	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"childsUuid")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"[]xs.string")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,14)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,false)

	flowNodeInstanceEntityProto.Fields = append(flowNodeInstanceEntityProto.Fields,"referenced")
	flowNodeInstanceEntityProto.FieldsType = append(flowNodeInstanceEntityProto.FieldsType,"[]EntityRef")
	flowNodeInstanceEntityProto.FieldsOrder = append(flowNodeInstanceEntityProto.FieldsOrder,15)
	flowNodeInstanceEntityProto.FieldsVisibility = append(flowNodeInstanceEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&flowNodeInstanceEntityProto)

}


////////////////////////////////////////////////////////////////////////////////
//              			ActivityInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_ActivityInstanceEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.ActivityInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeActivityInstanceEntity(objectId string, object interface{}) *BPMS_Runtime_ActivityInstanceEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeActivityInstanceExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.ActivityInstance).TYPENAME = "BPMS_Runtime.ActivityInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.ActivityInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_ActivityInstanceEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.ActivityInstance%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_ActivityInstanceEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.ActivityInstance)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.ActivityInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.ActivityInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.ActivityInstance","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_ActivityInstanceEntity) GetTypeName()string{
	return "BPMS_Runtime.ActivityInstance"
}
func(this *BPMS_Runtime_ActivityInstanceEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_ActivityInstanceEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_ActivityInstanceEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_ActivityInstanceEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_ActivityInstanceEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_ActivityInstanceEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_ActivityInstanceEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_ActivityInstanceEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_ActivityInstanceEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_ActivityInstanceEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_ActivityInstanceEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_ActivityInstanceEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_ActivityInstanceEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_ActivityInstanceEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_ActivityInstanceEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_ActivityInstanceEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_ActivityInstanceEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_ActivityInstanceEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_ActivityInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_ActivityInstanceEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_ActivityInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_ActivityInstanceEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_ActivityInstanceEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ActivityInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_ActivityInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_ActivityInstanceEntityPrototype() {

	var activityInstanceEntityProto EntityPrototype
	activityInstanceEntityProto.TypeName = "BPMS_Runtime.ActivityInstance"
	activityInstanceEntityProto.SuperTypeNames = append(activityInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.Instance")
	activityInstanceEntityProto.SuperTypeNames = append(activityInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.FlowNodeInstance")
	activityInstanceEntityProto.Ids = append(activityInstanceEntityProto.Ids,"uuid")
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"uuid")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"xs.string")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,0)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,false)
	activityInstanceEntityProto.Indexs = append(activityInstanceEntityProto.Indexs,"parentUuid")
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"parentUuid")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"xs.string")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,1)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,false)

	/** members of Instance **/
	activityInstanceEntityProto.Ids = append(activityInstanceEntityProto.Ids,"M_id")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,2)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_id")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"xs.ID")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,3)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_bpmnElementId")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"xs.string")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,4)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_participants")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"[]xs.string")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,5)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_dataRef")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,6)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_data")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,7)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_logInfoRef")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo:Ref")

	/** members of FlowNodeInstance **/
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,8)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_flowNodeType")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"enum:FlowNodeType_AbstractTask:FlowNodeType_ServiceTask:FlowNodeType_UserTask:FlowNodeType_ManualTask:FlowNodeType_BusinessRuleTask:FlowNodeType_ScriptTask:FlowNodeType_EmbeddedSubprocess:FlowNodeType_EventSubprocess:FlowNodeType_AdHocSubprocess:FlowNodeType_Transaction:FlowNodeType_CallActivity:FlowNodeType_ParallelGateway:FlowNodeType_ExclusiveGateway:FlowNodeType_InclusiveGateway:FlowNodeType_EventBasedGateway:FlowNodeType_ComplexGateway:FlowNodeType_StartEvent:FlowNodeType_IntermediateCatchEvent:FlowNodeType_BoundaryEvent:FlowNodeType_EndEvent:FlowNodeType_IntermediateThrowEvent")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,9)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_lifecycleState")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"enum:LifecycleState_Completed:LifecycleState_Compensated:LifecycleState_Failed:LifecycleState_Terminated:LifecycleState_Completing:LifecycleState_Compensating:LifecycleState_Failing:LifecycleState_Terminating:LifecycleState_Ready:LifecycleState_Active")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,10)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_inputRef")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject:Ref")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,11)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_outputRef")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject:Ref")

	/** members of ActivityInstance **/
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,12)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_activityType")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"enum:ActivityType_AbstractTask:ActivityType_ServiceTask:ActivityType_UserTask:ActivityType_ManualTask:ActivityType_BusinessRuleTask:ActivityType_ScriptTask:ActivityType_EmbeddedSubprocess:ActivityType_EventSubprocess:ActivityType_AdHocSubprocess:ActivityType_Transaction:ActivityType_CallActivity")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,13)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_multiInstanceBehaviorType")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"enum:MultiInstanceBehaviorType_None:MultiInstanceBehaviorType_One:MultiInstanceBehaviorType_All:MultiInstanceBehaviorType_Complex")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,14)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_loopCharacteristicType")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"enum:LoopCharacteristicType_StandardLoopCharacteristics:LoopCharacteristicType_MultiInstanceLoopCharacteristics")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,15)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,true)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_tokenCount")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"xs.int")

	/** associations of ActivityInstance **/
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,16)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,false)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_SubprocessInstancePtr")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"BPMS_Runtime.SubprocessInstance:Ref")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,17)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,false)
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"M_processInstancePtr")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"BPMS_Runtime.ProcessInstance:Ref")
	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"childsUuid")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"[]xs.string")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,18)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,false)

	activityInstanceEntityProto.Fields = append(activityInstanceEntityProto.Fields,"referenced")
	activityInstanceEntityProto.FieldsType = append(activityInstanceEntityProto.FieldsType,"[]EntityRef")
	activityInstanceEntityProto.FieldsOrder = append(activityInstanceEntityProto.FieldsOrder,19)
	activityInstanceEntityProto.FieldsVisibility = append(activityInstanceEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&activityInstanceEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_ActivityInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.ActivityInstance"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ActivityInstance"

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

		/** associations of ActivityInstance **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var ActivityInstanceInfo []interface{}

	ActivityInstanceInfo = append(ActivityInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		ActivityInstanceInfo = append(ActivityInstanceInfo, this.parentPtr.GetUuid())
	}else{
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
	dataIds := make([]string,0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(this.object.M_data[i].UUID,this.object.M_data[i])
		dataIds=append(dataIds,dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data",dataEntity)
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
	if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AbstractTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ServiceTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 1)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_UserTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 2)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ManualTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 3)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BusinessRuleTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 4)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ScriptTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 5)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EmbeddedSubprocess{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 6)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventSubprocess{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 7)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AdHocSubprocess{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 8)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_Transaction{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 9)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_CallActivity{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 10)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ParallelGateway{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 11)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ExclusiveGateway{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 12)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_InclusiveGateway{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 13)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventBasedGateway{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 14)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ComplexGateway{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 15)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_StartEvent{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 16)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateCatchEvent{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 17)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BoundaryEvent{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 18)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EndEvent{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 19)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateThrowEvent{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 20)
	}else{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	}

	/** Save lifecycleState type LifecycleState **/
	if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completed{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensated{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 1)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failed{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 2)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminated{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 3)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completing{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 4)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensating{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 5)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failing{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 6)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminating{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 7)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Ready{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 8)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Active{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 9)
	}else{
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
	if this.object.M_activityType==BPMS_Runtime.ActivityType_AbstractTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	} else if this.object.M_activityType==BPMS_Runtime.ActivityType_ServiceTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 1)
	} else if this.object.M_activityType==BPMS_Runtime.ActivityType_UserTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 2)
	} else if this.object.M_activityType==BPMS_Runtime.ActivityType_ManualTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 3)
	} else if this.object.M_activityType==BPMS_Runtime.ActivityType_BusinessRuleTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 4)
	} else if this.object.M_activityType==BPMS_Runtime.ActivityType_ScriptTask{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 5)
	} else if this.object.M_activityType==BPMS_Runtime.ActivityType_EmbeddedSubprocess{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 6)
	} else if this.object.M_activityType==BPMS_Runtime.ActivityType_EventSubprocess{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 7)
	} else if this.object.M_activityType==BPMS_Runtime.ActivityType_AdHocSubprocess{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 8)
	} else if this.object.M_activityType==BPMS_Runtime.ActivityType_Transaction{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 9)
	} else if this.object.M_activityType==BPMS_Runtime.ActivityType_CallActivity{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 10)
	}else{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	}

	/** Save multiInstanceBehaviorType type MultiInstanceBehaviorType **/
	if this.object.M_multiInstanceBehaviorType==BPMS_Runtime.MultiInstanceBehaviorType_None{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	} else if this.object.M_multiInstanceBehaviorType==BPMS_Runtime.MultiInstanceBehaviorType_One{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 1)
	} else if this.object.M_multiInstanceBehaviorType==BPMS_Runtime.MultiInstanceBehaviorType_All{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 2)
	} else if this.object.M_multiInstanceBehaviorType==BPMS_Runtime.MultiInstanceBehaviorType_Complex{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 3)
	}else{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	}

	/** Save loopCharacteristicType type LoopCharacteristicType **/
	if this.object.M_loopCharacteristicType==BPMS_Runtime.LoopCharacteristicType_StandardLoopCharacteristics{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	} else if this.object.M_loopCharacteristicType==BPMS_Runtime.LoopCharacteristicType_MultiInstanceLoopCharacteristics{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 1)
	}else{
		ActivityInstanceInfo = append(ActivityInstanceInfo, 0)
	}
	ActivityInstanceInfo = append(ActivityInstanceInfo, this.object.M_tokenCount)

	/** associations of ActivityInstance **/

	/** Save SubprocessInstance type SubprocessInstance **/
		ActivityInstanceInfo = append(ActivityInstanceInfo,this.object.M_SubprocessInstancePtr)

	/** Save processInstance type ProcessInstance **/
		ActivityInstanceInfo = append(ActivityInstanceInfo,this.object.M_processInstancePtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), ActivityInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), ActivityInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_ActivityInstanceEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_ActivityInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ActivityInstance"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ActivityInstance...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.ActivityInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.ActivityInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** bpmnElementId **/
 		if results[0][3] != nil{
 			this.object.M_bpmnElementId=results[0][3].(string)
 		}

		/** participants **/
 		if results[0][4] != nil{
 			this.object.M_participants= append(this.object.M_participants, results[0][4].([]string) ...)
 		}

		/** dataRef **/
 		if results[0][5] != nil{
			idsStr :=results[0][5].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ItemAwareElementInstance"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef,ids[i])
					GetServer().GetEntityManager().appendReference("dataRef",this.object.UUID, id_)
				}
			}
 		}

		/** data **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_Runtime_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_Runtime_ItemAwareElementInstanceEntity)
					}else{
						dataEntity = GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data",dataEntity)
				}
 			}
 		}

		/** logInfoRef **/
 		if results[0][7] != nil{
			idsStr :=results[0][7].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.LogInfo"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef,ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of FlowNodeInstance **/

		/** flowNodeType **/
 		if results[0][8] != nil{
 			enumIndex := results[0][8].(int)
			if enumIndex == 0{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AbstractTask
			} else if enumIndex == 1{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ServiceTask
			} else if enumIndex == 2{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_UserTask
			} else if enumIndex == 3{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ManualTask
			} else if enumIndex == 4{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BusinessRuleTask
			} else if enumIndex == 5{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ScriptTask
			} else if enumIndex == 6{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EmbeddedSubprocess
			} else if enumIndex == 7{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventSubprocess
			} else if enumIndex == 8{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AdHocSubprocess
			} else if enumIndex == 9{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_Transaction
			} else if enumIndex == 10{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_CallActivity
			} else if enumIndex == 11{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ParallelGateway
			} else if enumIndex == 12{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ExclusiveGateway
			} else if enumIndex == 13{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_InclusiveGateway
			} else if enumIndex == 14{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventBasedGateway
			} else if enumIndex == 15{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ComplexGateway
			} else if enumIndex == 16{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_StartEvent
			} else if enumIndex == 17{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateCatchEvent
			} else if enumIndex == 18{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BoundaryEvent
			} else if enumIndex == 19{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EndEvent
			} else if enumIndex == 20{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateThrowEvent
 			}
 		}

		/** lifecycleState **/
 		if results[0][9] != nil{
 			enumIndex := results[0][9].(int)
			if enumIndex == 0{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completed
			} else if enumIndex == 1{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensated
			} else if enumIndex == 2{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failed
			} else if enumIndex == 3{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminated
			} else if enumIndex == 4{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completing
			} else if enumIndex == 5{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensating
			} else if enumIndex == 6{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failing
			} else if enumIndex == 7{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminating
			} else if enumIndex == 8{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Ready
			} else if enumIndex == 9{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Active
 			}
 		}

		/** inputRef **/
 		if results[0][10] != nil{
			idsStr :=results[0][10].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ConnectingObject"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_inputRef = append(this.object.M_inputRef,ids[i])
					GetServer().GetEntityManager().appendReference("inputRef",this.object.UUID, id_)
				}
			}
 		}

		/** outputRef **/
 		if results[0][11] != nil{
			idsStr :=results[0][11].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ConnectingObject"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_outputRef = append(this.object.M_outputRef,ids[i])
					GetServer().GetEntityManager().appendReference("outputRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of ActivityInstance **/

		/** activityType **/
 		if results[0][12] != nil{
 			enumIndex := results[0][12].(int)
			if enumIndex == 0{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_AbstractTask
			} else if enumIndex == 1{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_ServiceTask
			} else if enumIndex == 2{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_UserTask
			} else if enumIndex == 3{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_ManualTask
			} else if enumIndex == 4{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_BusinessRuleTask
			} else if enumIndex == 5{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_ScriptTask
			} else if enumIndex == 6{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_EmbeddedSubprocess
			} else if enumIndex == 7{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_EventSubprocess
			} else if enumIndex == 8{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_AdHocSubprocess
			} else if enumIndex == 9{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_Transaction
			} else if enumIndex == 10{
 				this.object.M_activityType=BPMS_Runtime.ActivityType_CallActivity
 			}
 		}

		/** multiInstanceBehaviorType **/
 		if results[0][13] != nil{
 			enumIndex := results[0][13].(int)
			if enumIndex == 0{
 				this.object.M_multiInstanceBehaviorType=BPMS_Runtime.MultiInstanceBehaviorType_None
			} else if enumIndex == 1{
 				this.object.M_multiInstanceBehaviorType=BPMS_Runtime.MultiInstanceBehaviorType_One
			} else if enumIndex == 2{
 				this.object.M_multiInstanceBehaviorType=BPMS_Runtime.MultiInstanceBehaviorType_All
			} else if enumIndex == 3{
 				this.object.M_multiInstanceBehaviorType=BPMS_Runtime.MultiInstanceBehaviorType_Complex
 			}
 		}

		/** loopCharacteristicType **/
 		if results[0][14] != nil{
 			enumIndex := results[0][14].(int)
			if enumIndex == 0{
 				this.object.M_loopCharacteristicType=BPMS_Runtime.LoopCharacteristicType_StandardLoopCharacteristics
			} else if enumIndex == 1{
 				this.object.M_loopCharacteristicType=BPMS_Runtime.LoopCharacteristicType_MultiInstanceLoopCharacteristics
 			}
 		}

		/** tokenCount **/
 		if results[0][15] != nil{
 			this.object.M_tokenCount=results[0][15].(int)
 		}

		/** associations of ActivityInstance **/

		/** SubprocessInstancePtr **/
 		if results[0][16] != nil{
			id :=results[0][16].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.SubprocessInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_SubprocessInstancePtr= id
				GetServer().GetEntityManager().appendReference("SubprocessInstancePtr",this.object.UUID, id_)
			}
 		}

		/** processInstancePtr **/
 		if results[0][17] != nil{
			id :=results[0][17].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.ProcessInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_processInstancePtr= id
				GetServer().GetEntityManager().appendReference("processInstancePtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeActivityInstanceEntityFromObject(object *BPMS_Runtime.ActivityInstance) *BPMS_Runtime_ActivityInstanceEntity {
	 return this.NewBPMS_RuntimeActivityInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_ActivityInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeActivityInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ActivityInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_ActivityInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_ActivityInstanceEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			SubprocessInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_SubprocessInstanceEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.SubprocessInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeSubprocessInstanceEntity(objectId string, object interface{}) *BPMS_Runtime_SubprocessInstanceEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeSubprocessInstanceExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.SubprocessInstance).TYPENAME = "BPMS_Runtime.SubprocessInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.SubprocessInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_SubprocessInstanceEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.SubprocessInstance%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_SubprocessInstanceEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.SubprocessInstance)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.SubprocessInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.SubprocessInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.SubprocessInstance","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_SubprocessInstanceEntity) GetTypeName()string{
	return "BPMS_Runtime.SubprocessInstance"
}
func(this *BPMS_Runtime_SubprocessInstanceEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_SubprocessInstanceEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_SubprocessInstanceEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_SubprocessInstanceEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_SubprocessInstanceEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_SubprocessInstanceEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.SubprocessInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_SubprocessInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_SubprocessInstanceEntityPrototype() {

	var subprocessInstanceEntityProto EntityPrototype
	subprocessInstanceEntityProto.TypeName = "BPMS_Runtime.SubprocessInstance"
	subprocessInstanceEntityProto.SuperTypeNames = append(subprocessInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.Instance")
	subprocessInstanceEntityProto.SuperTypeNames = append(subprocessInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.FlowNodeInstance")
	subprocessInstanceEntityProto.Ids = append(subprocessInstanceEntityProto.Ids,"uuid")
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"uuid")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"xs.string")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,0)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,false)
	subprocessInstanceEntityProto.Indexs = append(subprocessInstanceEntityProto.Indexs,"parentUuid")
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"parentUuid")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"xs.string")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,1)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,false)

	/** members of Instance **/
	subprocessInstanceEntityProto.Ids = append(subprocessInstanceEntityProto.Ids,"M_id")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,2)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_id")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"xs.ID")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,3)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_bpmnElementId")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"xs.string")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,4)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_participants")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"[]xs.string")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,5)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_dataRef")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,6)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_data")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,7)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_logInfoRef")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo:Ref")

	/** members of FlowNodeInstance **/
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,8)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_flowNodeType")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"enum:FlowNodeType_AbstractTask:FlowNodeType_ServiceTask:FlowNodeType_UserTask:FlowNodeType_ManualTask:FlowNodeType_BusinessRuleTask:FlowNodeType_ScriptTask:FlowNodeType_EmbeddedSubprocess:FlowNodeType_EventSubprocess:FlowNodeType_AdHocSubprocess:FlowNodeType_Transaction:FlowNodeType_CallActivity:FlowNodeType_ParallelGateway:FlowNodeType_ExclusiveGateway:FlowNodeType_InclusiveGateway:FlowNodeType_EventBasedGateway:FlowNodeType_ComplexGateway:FlowNodeType_StartEvent:FlowNodeType_IntermediateCatchEvent:FlowNodeType_BoundaryEvent:FlowNodeType_EndEvent:FlowNodeType_IntermediateThrowEvent")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,9)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_lifecycleState")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"enum:LifecycleState_Completed:LifecycleState_Compensated:LifecycleState_Failed:LifecycleState_Terminated:LifecycleState_Completing:LifecycleState_Compensating:LifecycleState_Failing:LifecycleState_Terminating:LifecycleState_Ready:LifecycleState_Active")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,10)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_inputRef")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject:Ref")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,11)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_outputRef")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject:Ref")

	/** members of SubprocessInstance **/
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,12)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_SubprocessType")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"enum:SubprocessType_EmbeddedSubprocess:SubprocessType_EventSubprocess:SubprocessType_AdHocSubprocess:SubprocessType_Transaction")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,13)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_flowNodeInstances")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"[]BPMS_Runtime.FlowNodeInstance")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,14)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,true)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_connectingObjects")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject")

	/** associations of SubprocessInstance **/
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,15)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,false)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_SubprocessInstancePtr")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"BPMS_Runtime.SubprocessInstance:Ref")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,16)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,false)
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"M_processInstancePtr")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"BPMS_Runtime.ProcessInstance:Ref")
	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"childsUuid")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"[]xs.string")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,17)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,false)

	subprocessInstanceEntityProto.Fields = append(subprocessInstanceEntityProto.Fields,"referenced")
	subprocessInstanceEntityProto.FieldsType = append(subprocessInstanceEntityProto.FieldsType,"[]EntityRef")
	subprocessInstanceEntityProto.FieldsOrder = append(subprocessInstanceEntityProto.FieldsOrder,18)
	subprocessInstanceEntityProto.FieldsVisibility = append(subprocessInstanceEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&subprocessInstanceEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_SubprocessInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.SubprocessInstance"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.SubprocessInstance"

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

		/** associations of SubprocessInstance **/
	query.Fields = append(query.Fields, "M_SubprocessInstancePtr")
	query.Fields = append(query.Fields, "M_processInstancePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var SubprocessInstanceInfo []interface{}

	SubprocessInstanceInfo = append(SubprocessInstanceInfo, this.GetUuid())
	if this.parentPtr != nil {
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, this.parentPtr.GetUuid())
	}else{
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
	dataIds := make([]string,0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(this.object.M_data[i].UUID,this.object.M_data[i])
		dataIds=append(dataIds,dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data",dataEntity)
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
	if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AbstractTask{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ServiceTask{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 1)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_UserTask{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 2)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ManualTask{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 3)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BusinessRuleTask{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 4)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ScriptTask{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 5)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EmbeddedSubprocess{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 6)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventSubprocess{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 7)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AdHocSubprocess{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 8)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_Transaction{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 9)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_CallActivity{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 10)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ParallelGateway{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 11)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ExclusiveGateway{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 12)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_InclusiveGateway{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 13)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventBasedGateway{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 14)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ComplexGateway{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 15)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_StartEvent{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 16)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateCatchEvent{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 17)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BoundaryEvent{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 18)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EndEvent{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 19)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateThrowEvent{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 20)
	}else{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	}

	/** Save lifecycleState type LifecycleState **/
	if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completed{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensated{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 1)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failed{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 2)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminated{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 3)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completing{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 4)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensating{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 5)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failing{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 6)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminating{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 7)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Ready{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 8)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Active{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 9)
	}else{
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
	if this.object.M_SubprocessType==BPMS_Runtime.SubprocessType_EmbeddedSubprocess{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	} else if this.object.M_SubprocessType==BPMS_Runtime.SubprocessType_EventSubprocess{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 1)
	} else if this.object.M_SubprocessType==BPMS_Runtime.SubprocessType_AdHocSubprocess{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 2)
	} else if this.object.M_SubprocessType==BPMS_Runtime.SubprocessType_Transaction{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 3)
	}else{
		SubprocessInstanceInfo = append(SubprocessInstanceInfo, 0)
	}

	/** Save flowNodeInstances type FlowNodeInstance **/
	flowNodeInstancesIds := make([]string,0)
	for i := 0; i < len(this.object.M_flowNodeInstances); i++ {
		switch v := this.object.M_flowNodeInstances[i].(type) {
		case *BPMS_Runtime.ActivityInstance:
		flowNodeInstancesEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeActivityInstanceEntity(v.UUID, v)
		flowNodeInstancesIds=append(flowNodeInstancesIds,flowNodeInstancesEntity.uuid)
		flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
		this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
		if flowNodeInstancesEntity.NeedSave() {
			flowNodeInstancesEntity.SaveEntity()
		}
		case *BPMS_Runtime.SubprocessInstance:
		flowNodeInstancesEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeSubprocessInstanceEntity(v.UUID, v)
		flowNodeInstancesIds=append(flowNodeInstancesIds,flowNodeInstancesEntity.uuid)
		flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
		this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
		if flowNodeInstancesEntity.NeedSave() {
			flowNodeInstancesEntity.SaveEntity()
		}
		case *BPMS_Runtime.GatewayInstance:
		flowNodeInstancesEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeGatewayInstanceEntity(v.UUID, v)
		flowNodeInstancesIds=append(flowNodeInstancesIds,flowNodeInstancesEntity.uuid)
		flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
		this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
		if flowNodeInstancesEntity.NeedSave() {
			flowNodeInstancesEntity.SaveEntity()
		}
		case *BPMS_Runtime.EventInstance:
		flowNodeInstancesEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeEventInstanceEntity(v.UUID, v)
		flowNodeInstancesIds=append(flowNodeInstancesIds,flowNodeInstancesEntity.uuid)
		flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
		this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
		if flowNodeInstancesEntity.NeedSave() {
			flowNodeInstancesEntity.SaveEntity()
		}
		}
	}
	flowNodeInstancesStr, _ := json.Marshal(flowNodeInstancesIds)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(flowNodeInstancesStr))

	/** Save connectingObjects type ConnectingObject **/
	connectingObjectsIds := make([]string,0)
	for i := 0; i < len(this.object.M_connectingObjects); i++ {
		connectingObjectsEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeConnectingObjectEntity(this.object.M_connectingObjects[i].UUID,this.object.M_connectingObjects[i])
		connectingObjectsIds=append(connectingObjectsIds,connectingObjectsEntity.uuid)
		connectingObjectsEntity.AppendReferenced("connectingObjects", this)
		this.AppendChild("connectingObjects",connectingObjectsEntity)
		if connectingObjectsEntity.NeedSave() {
			connectingObjectsEntity.SaveEntity()
		}
	}
	connectingObjectsStr, _ := json.Marshal(connectingObjectsIds)
	SubprocessInstanceInfo = append(SubprocessInstanceInfo, string(connectingObjectsStr))

	/** associations of SubprocessInstance **/

	/** Save SubprocessInstance type SubprocessInstance **/
		SubprocessInstanceInfo = append(SubprocessInstanceInfo,this.object.M_SubprocessInstancePtr)

	/** Save processInstance type ProcessInstance **/
		SubprocessInstanceInfo = append(SubprocessInstanceInfo,this.object.M_processInstancePtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), SubprocessInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), SubprocessInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_SubprocessInstanceEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_SubprocessInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.SubprocessInstance"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of SubprocessInstance...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.SubprocessInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.SubprocessInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** bpmnElementId **/
 		if results[0][3] != nil{
 			this.object.M_bpmnElementId=results[0][3].(string)
 		}

		/** participants **/
 		if results[0][4] != nil{
 			this.object.M_participants= append(this.object.M_participants, results[0][4].([]string) ...)
 		}

		/** dataRef **/
 		if results[0][5] != nil{
			idsStr :=results[0][5].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ItemAwareElementInstance"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef,ids[i])
					GetServer().GetEntityManager().appendReference("dataRef",this.object.UUID, id_)
				}
			}
 		}

		/** data **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_Runtime_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_Runtime_ItemAwareElementInstanceEntity)
					}else{
						dataEntity = GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data",dataEntity)
				}
 			}
 		}

		/** logInfoRef **/
 		if results[0][7] != nil{
			idsStr :=results[0][7].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.LogInfo"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef,ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of FlowNodeInstance **/

		/** flowNodeType **/
 		if results[0][8] != nil{
 			enumIndex := results[0][8].(int)
			if enumIndex == 0{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AbstractTask
			} else if enumIndex == 1{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ServiceTask
			} else if enumIndex == 2{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_UserTask
			} else if enumIndex == 3{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ManualTask
			} else if enumIndex == 4{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BusinessRuleTask
			} else if enumIndex == 5{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ScriptTask
			} else if enumIndex == 6{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EmbeddedSubprocess
			} else if enumIndex == 7{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventSubprocess
			} else if enumIndex == 8{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AdHocSubprocess
			} else if enumIndex == 9{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_Transaction
			} else if enumIndex == 10{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_CallActivity
			} else if enumIndex == 11{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ParallelGateway
			} else if enumIndex == 12{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ExclusiveGateway
			} else if enumIndex == 13{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_InclusiveGateway
			} else if enumIndex == 14{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventBasedGateway
			} else if enumIndex == 15{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ComplexGateway
			} else if enumIndex == 16{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_StartEvent
			} else if enumIndex == 17{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateCatchEvent
			} else if enumIndex == 18{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BoundaryEvent
			} else if enumIndex == 19{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EndEvent
			} else if enumIndex == 20{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateThrowEvent
 			}
 		}

		/** lifecycleState **/
 		if results[0][9] != nil{
 			enumIndex := results[0][9].(int)
			if enumIndex == 0{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completed
			} else if enumIndex == 1{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensated
			} else if enumIndex == 2{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failed
			} else if enumIndex == 3{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminated
			} else if enumIndex == 4{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completing
			} else if enumIndex == 5{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensating
			} else if enumIndex == 6{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failing
			} else if enumIndex == 7{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminating
			} else if enumIndex == 8{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Ready
			} else if enumIndex == 9{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Active
 			}
 		}

		/** inputRef **/
 		if results[0][10] != nil{
			idsStr :=results[0][10].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ConnectingObject"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_inputRef = append(this.object.M_inputRef,ids[i])
					GetServer().GetEntityManager().appendReference("inputRef",this.object.UUID, id_)
				}
			}
 		}

		/** outputRef **/
 		if results[0][11] != nil{
			idsStr :=results[0][11].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ConnectingObject"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_outputRef = append(this.object.M_outputRef,ids[i])
					GetServer().GetEntityManager().appendReference("outputRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of SubprocessInstance **/

		/** SubprocessType **/
 		if results[0][12] != nil{
 			enumIndex := results[0][12].(int)
			if enumIndex == 0{
 				this.object.M_SubprocessType=BPMS_Runtime.SubprocessType_EmbeddedSubprocess
			} else if enumIndex == 1{
 				this.object.M_SubprocessType=BPMS_Runtime.SubprocessType_EventSubprocess
			} else if enumIndex == 2{
 				this.object.M_SubprocessType=BPMS_Runtime.SubprocessType_AdHocSubprocess
			} else if enumIndex == 3{
 				this.object.M_SubprocessType=BPMS_Runtime.SubprocessType_Transaction
 			}
 		}

		/** flowNodeInstances **/
 		if results[0][13] != nil{
			uuidsStr :=results[0][13].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
			typeName := uuids[i][0:strings.Index(uuids[i], "%")]
			if err!=nil{
				log.Println("type ", typeName, " not found!")
				return err
			}
				if typeName == "BPMS_Runtime.ActivityInstance"{
						if len(uuids[i]) > 0 {
							var flowNodeInstancesEntity *BPMS_Runtime_ActivityInstanceEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								flowNodeInstancesEntity = instance.(*BPMS_Runtime_ActivityInstanceEntity)
							}else{
								flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMS_RuntimeActivityInstanceEntity(uuids[i], nil)
								flowNodeInstancesEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
							}
							flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
							this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
						}
				} else if typeName == "BPMS_Runtime.SubprocessInstance"{
						if len(uuids[i]) > 0 {
							var flowNodeInstancesEntity *BPMS_Runtime_SubprocessInstanceEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								flowNodeInstancesEntity = instance.(*BPMS_Runtime_SubprocessInstanceEntity)
							}else{
								flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMS_RuntimeSubprocessInstanceEntity(uuids[i], nil)
								flowNodeInstancesEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
							}
							flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
							this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
						}
				} else if typeName == "BPMS_Runtime.GatewayInstance"{
						if len(uuids[i]) > 0 {
							var flowNodeInstancesEntity *BPMS_Runtime_GatewayInstanceEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								flowNodeInstancesEntity = instance.(*BPMS_Runtime_GatewayInstanceEntity)
							}else{
								flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMS_RuntimeGatewayInstanceEntity(uuids[i], nil)
								flowNodeInstancesEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
							}
							flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
							this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
						}
				} else if typeName == "BPMS_Runtime.EventInstance"{
						if len(uuids[i]) > 0 {
							var flowNodeInstancesEntity *BPMS_Runtime_EventInstanceEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								flowNodeInstancesEntity = instance.(*BPMS_Runtime_EventInstanceEntity)
							}else{
								flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMS_RuntimeEventInstanceEntity(uuids[i], nil)
								flowNodeInstancesEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
							}
							flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
							this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
						}
				}
 			}
 		}

		/** connectingObjects **/
 		if results[0][14] != nil{
			uuidsStr :=results[0][14].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var connectingObjectsEntity *BPMS_Runtime_ConnectingObjectEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						connectingObjectsEntity = instance.(*BPMS_Runtime_ConnectingObjectEntity)
					}else{
						connectingObjectsEntity = GetServer().GetEntityManager().NewBPMS_RuntimeConnectingObjectEntity(uuids[i], nil)
						connectingObjectsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(connectingObjectsEntity)
					}
					connectingObjectsEntity.AppendReferenced("connectingObjects", this)
					this.AppendChild("connectingObjects",connectingObjectsEntity)
				}
 			}
 		}

		/** associations of SubprocessInstance **/

		/** SubprocessInstancePtr **/
 		if results[0][15] != nil{
			id :=results[0][15].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.SubprocessInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_SubprocessInstancePtr= id
				GetServer().GetEntityManager().appendReference("SubprocessInstancePtr",this.object.UUID, id_)
			}
 		}

		/** processInstancePtr **/
 		if results[0][16] != nil{
			id :=results[0][16].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.ProcessInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_processInstancePtr= id
				GetServer().GetEntityManager().appendReference("processInstancePtr",this.object.UUID, id_)
			}
 		}
		childsUuidStr := results[0][17].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][18].(string)
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
func (this *EntityManager) NewBPMS_RuntimeSubprocessInstanceEntityFromObject(object *BPMS_Runtime.SubprocessInstance) *BPMS_Runtime_SubprocessInstanceEntity {
	 return this.NewBPMS_RuntimeSubprocessInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_SubprocessInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeSubprocessInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.SubprocessInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_SubprocessInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_SubprocessInstanceEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			GatewayInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_GatewayInstanceEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.GatewayInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeGatewayInstanceEntity(objectId string, object interface{}) *BPMS_Runtime_GatewayInstanceEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeGatewayInstanceExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.GatewayInstance).TYPENAME = "BPMS_Runtime.GatewayInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.GatewayInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_GatewayInstanceEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.GatewayInstance%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_GatewayInstanceEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.GatewayInstance)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.GatewayInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.GatewayInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.GatewayInstance","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_GatewayInstanceEntity) GetTypeName()string{
	return "BPMS_Runtime.GatewayInstance"
}
func(this *BPMS_Runtime_GatewayInstanceEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_GatewayInstanceEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_GatewayInstanceEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_GatewayInstanceEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_GatewayInstanceEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_GatewayInstanceEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_GatewayInstanceEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_GatewayInstanceEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_GatewayInstanceEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_GatewayInstanceEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_GatewayInstanceEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_GatewayInstanceEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_GatewayInstanceEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_GatewayInstanceEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_GatewayInstanceEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_GatewayInstanceEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_GatewayInstanceEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_GatewayInstanceEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_GatewayInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_GatewayInstanceEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_GatewayInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_GatewayInstanceEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_GatewayInstanceEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.GatewayInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_GatewayInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_GatewayInstanceEntityPrototype() {

	var gatewayInstanceEntityProto EntityPrototype
	gatewayInstanceEntityProto.TypeName = "BPMS_Runtime.GatewayInstance"
	gatewayInstanceEntityProto.SuperTypeNames = append(gatewayInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.Instance")
	gatewayInstanceEntityProto.SuperTypeNames = append(gatewayInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.FlowNodeInstance")
	gatewayInstanceEntityProto.Ids = append(gatewayInstanceEntityProto.Ids,"uuid")
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"uuid")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"xs.string")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,0)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,false)
	gatewayInstanceEntityProto.Indexs = append(gatewayInstanceEntityProto.Indexs,"parentUuid")
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"parentUuid")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"xs.string")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,1)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,false)

	/** members of Instance **/
	gatewayInstanceEntityProto.Ids = append(gatewayInstanceEntityProto.Ids,"M_id")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,2)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_id")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"xs.ID")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,3)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_bpmnElementId")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"xs.string")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,4)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_participants")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"[]xs.string")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,5)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_dataRef")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,6)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_data")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,7)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_logInfoRef")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo:Ref")

	/** members of FlowNodeInstance **/
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,8)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_flowNodeType")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"enum:FlowNodeType_AbstractTask:FlowNodeType_ServiceTask:FlowNodeType_UserTask:FlowNodeType_ManualTask:FlowNodeType_BusinessRuleTask:FlowNodeType_ScriptTask:FlowNodeType_EmbeddedSubprocess:FlowNodeType_EventSubprocess:FlowNodeType_AdHocSubprocess:FlowNodeType_Transaction:FlowNodeType_CallActivity:FlowNodeType_ParallelGateway:FlowNodeType_ExclusiveGateway:FlowNodeType_InclusiveGateway:FlowNodeType_EventBasedGateway:FlowNodeType_ComplexGateway:FlowNodeType_StartEvent:FlowNodeType_IntermediateCatchEvent:FlowNodeType_BoundaryEvent:FlowNodeType_EndEvent:FlowNodeType_IntermediateThrowEvent")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,9)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_lifecycleState")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"enum:LifecycleState_Completed:LifecycleState_Compensated:LifecycleState_Failed:LifecycleState_Terminated:LifecycleState_Completing:LifecycleState_Compensating:LifecycleState_Failing:LifecycleState_Terminating:LifecycleState_Ready:LifecycleState_Active")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,10)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_inputRef")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject:Ref")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,11)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_outputRef")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject:Ref")

	/** members of GatewayInstance **/
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,12)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,true)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_gatewayType")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"enum:GatewayType_ParallelGateway:GatewayType_ExclusiveGateway:GatewayType_InclusiveGateway:GatewayType_EventBasedGateway:GatewayType_ComplexGateway")

	/** associations of GatewayInstance **/
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,13)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,false)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_SubprocessInstancePtr")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"BPMS_Runtime.SubprocessInstance:Ref")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,14)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,false)
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"M_processInstancePtr")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"BPMS_Runtime.ProcessInstance:Ref")
	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"childsUuid")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"[]xs.string")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,15)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,false)

	gatewayInstanceEntityProto.Fields = append(gatewayInstanceEntityProto.Fields,"referenced")
	gatewayInstanceEntityProto.FieldsType = append(gatewayInstanceEntityProto.FieldsType,"[]EntityRef")
	gatewayInstanceEntityProto.FieldsOrder = append(gatewayInstanceEntityProto.FieldsOrder,16)
	gatewayInstanceEntityProto.FieldsVisibility = append(gatewayInstanceEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&gatewayInstanceEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_GatewayInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.GatewayInstance"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.GatewayInstance"

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
	}else{
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
	dataIds := make([]string,0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(this.object.M_data[i].UUID,this.object.M_data[i])
		dataIds=append(dataIds,dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data",dataEntity)
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
	if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AbstractTask{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ServiceTask{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 1)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_UserTask{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 2)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ManualTask{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 3)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BusinessRuleTask{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 4)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ScriptTask{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 5)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EmbeddedSubprocess{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 6)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventSubprocess{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 7)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AdHocSubprocess{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 8)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_Transaction{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 9)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_CallActivity{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 10)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ParallelGateway{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 11)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ExclusiveGateway{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 12)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_InclusiveGateway{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 13)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventBasedGateway{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 14)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ComplexGateway{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 15)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_StartEvent{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 16)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateCatchEvent{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 17)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BoundaryEvent{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 18)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EndEvent{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 19)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateThrowEvent{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 20)
	}else{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	}

	/** Save lifecycleState type LifecycleState **/
	if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completed{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensated{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 1)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failed{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 2)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminated{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 3)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completing{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 4)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensating{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 5)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failing{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 6)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminating{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 7)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Ready{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 8)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Active{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 9)
	}else{
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
	if this.object.M_gatewayType==BPMS_Runtime.GatewayType_ParallelGateway{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	} else if this.object.M_gatewayType==BPMS_Runtime.GatewayType_ExclusiveGateway{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 1)
	} else if this.object.M_gatewayType==BPMS_Runtime.GatewayType_InclusiveGateway{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 2)
	} else if this.object.M_gatewayType==BPMS_Runtime.GatewayType_EventBasedGateway{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 3)
	} else if this.object.M_gatewayType==BPMS_Runtime.GatewayType_ComplexGateway{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 4)
	}else{
		GatewayInstanceInfo = append(GatewayInstanceInfo, 0)
	}

	/** associations of GatewayInstance **/

	/** Save SubprocessInstance type SubprocessInstance **/
		GatewayInstanceInfo = append(GatewayInstanceInfo,this.object.M_SubprocessInstancePtr)

	/** Save processInstance type ProcessInstance **/
		GatewayInstanceInfo = append(GatewayInstanceInfo,this.object.M_processInstancePtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), GatewayInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), GatewayInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_GatewayInstanceEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_GatewayInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.GatewayInstance"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of GatewayInstance...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.GatewayInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.GatewayInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** bpmnElementId **/
 		if results[0][3] != nil{
 			this.object.M_bpmnElementId=results[0][3].(string)
 		}

		/** participants **/
 		if results[0][4] != nil{
 			this.object.M_participants= append(this.object.M_participants, results[0][4].([]string) ...)
 		}

		/** dataRef **/
 		if results[0][5] != nil{
			idsStr :=results[0][5].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ItemAwareElementInstance"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef,ids[i])
					GetServer().GetEntityManager().appendReference("dataRef",this.object.UUID, id_)
				}
			}
 		}

		/** data **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_Runtime_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_Runtime_ItemAwareElementInstanceEntity)
					}else{
						dataEntity = GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data",dataEntity)
				}
 			}
 		}

		/** logInfoRef **/
 		if results[0][7] != nil{
			idsStr :=results[0][7].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.LogInfo"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef,ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of FlowNodeInstance **/

		/** flowNodeType **/
 		if results[0][8] != nil{
 			enumIndex := results[0][8].(int)
			if enumIndex == 0{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AbstractTask
			} else if enumIndex == 1{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ServiceTask
			} else if enumIndex == 2{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_UserTask
			} else if enumIndex == 3{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ManualTask
			} else if enumIndex == 4{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BusinessRuleTask
			} else if enumIndex == 5{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ScriptTask
			} else if enumIndex == 6{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EmbeddedSubprocess
			} else if enumIndex == 7{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventSubprocess
			} else if enumIndex == 8{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AdHocSubprocess
			} else if enumIndex == 9{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_Transaction
			} else if enumIndex == 10{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_CallActivity
			} else if enumIndex == 11{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ParallelGateway
			} else if enumIndex == 12{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ExclusiveGateway
			} else if enumIndex == 13{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_InclusiveGateway
			} else if enumIndex == 14{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventBasedGateway
			} else if enumIndex == 15{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ComplexGateway
			} else if enumIndex == 16{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_StartEvent
			} else if enumIndex == 17{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateCatchEvent
			} else if enumIndex == 18{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BoundaryEvent
			} else if enumIndex == 19{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EndEvent
			} else if enumIndex == 20{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateThrowEvent
 			}
 		}

		/** lifecycleState **/
 		if results[0][9] != nil{
 			enumIndex := results[0][9].(int)
			if enumIndex == 0{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completed
			} else if enumIndex == 1{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensated
			} else if enumIndex == 2{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failed
			} else if enumIndex == 3{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminated
			} else if enumIndex == 4{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completing
			} else if enumIndex == 5{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensating
			} else if enumIndex == 6{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failing
			} else if enumIndex == 7{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminating
			} else if enumIndex == 8{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Ready
			} else if enumIndex == 9{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Active
 			}
 		}

		/** inputRef **/
 		if results[0][10] != nil{
			idsStr :=results[0][10].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ConnectingObject"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_inputRef = append(this.object.M_inputRef,ids[i])
					GetServer().GetEntityManager().appendReference("inputRef",this.object.UUID, id_)
				}
			}
 		}

		/** outputRef **/
 		if results[0][11] != nil{
			idsStr :=results[0][11].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ConnectingObject"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_outputRef = append(this.object.M_outputRef,ids[i])
					GetServer().GetEntityManager().appendReference("outputRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of GatewayInstance **/

		/** gatewayType **/
 		if results[0][12] != nil{
 			enumIndex := results[0][12].(int)
			if enumIndex == 0{
 				this.object.M_gatewayType=BPMS_Runtime.GatewayType_ParallelGateway
			} else if enumIndex == 1{
 				this.object.M_gatewayType=BPMS_Runtime.GatewayType_ExclusiveGateway
			} else if enumIndex == 2{
 				this.object.M_gatewayType=BPMS_Runtime.GatewayType_InclusiveGateway
			} else if enumIndex == 3{
 				this.object.M_gatewayType=BPMS_Runtime.GatewayType_EventBasedGateway
			} else if enumIndex == 4{
 				this.object.M_gatewayType=BPMS_Runtime.GatewayType_ComplexGateway
 			}
 		}

		/** associations of GatewayInstance **/

		/** SubprocessInstancePtr **/
 		if results[0][13] != nil{
			id :=results[0][13].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.SubprocessInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_SubprocessInstancePtr= id
				GetServer().GetEntityManager().appendReference("SubprocessInstancePtr",this.object.UUID, id_)
			}
 		}

		/** processInstancePtr **/
 		if results[0][14] != nil{
			id :=results[0][14].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.ProcessInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_processInstancePtr= id
				GetServer().GetEntityManager().appendReference("processInstancePtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeGatewayInstanceEntityFromObject(object *BPMS_Runtime.GatewayInstance) *BPMS_Runtime_GatewayInstanceEntity {
	 return this.NewBPMS_RuntimeGatewayInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_GatewayInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeGatewayInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.GatewayInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_GatewayInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_GatewayInstanceEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			EventInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_EventInstanceEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.EventInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeEventInstanceEntity(objectId string, object interface{}) *BPMS_Runtime_EventInstanceEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeEventInstanceExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.EventInstance).TYPENAME = "BPMS_Runtime.EventInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.EventInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_EventInstanceEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.EventInstance%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_EventInstanceEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.EventInstance)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.EventInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.EventInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.EventInstance","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_EventInstanceEntity) GetTypeName()string{
	return "BPMS_Runtime.EventInstance"
}
func(this *BPMS_Runtime_EventInstanceEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_EventInstanceEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_EventInstanceEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_EventInstanceEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_EventInstanceEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_EventInstanceEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_EventInstanceEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_EventInstanceEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_EventInstanceEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_EventInstanceEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_EventInstanceEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_EventInstanceEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_EventInstanceEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_EventInstanceEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_EventInstanceEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_EventInstanceEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_EventInstanceEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_EventInstanceEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_EventInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_EventInstanceEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_EventInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_EventInstanceEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_EventInstanceEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_EventInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_EventInstanceEntityPrototype() {

	var eventInstanceEntityProto EntityPrototype
	eventInstanceEntityProto.TypeName = "BPMS_Runtime.EventInstance"
	eventInstanceEntityProto.SuperTypeNames = append(eventInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.Instance")
	eventInstanceEntityProto.SuperTypeNames = append(eventInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.FlowNodeInstance")
	eventInstanceEntityProto.Ids = append(eventInstanceEntityProto.Ids,"uuid")
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"uuid")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"xs.string")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,0)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,false)
	eventInstanceEntityProto.Indexs = append(eventInstanceEntityProto.Indexs,"parentUuid")
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"parentUuid")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"xs.string")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,1)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,false)

	/** members of Instance **/
	eventInstanceEntityProto.Ids = append(eventInstanceEntityProto.Ids,"M_id")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,2)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_id")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"xs.ID")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,3)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_bpmnElementId")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"xs.string")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,4)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_participants")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"[]xs.string")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,5)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_dataRef")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,6)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_data")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,7)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_logInfoRef")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo:Ref")

	/** members of FlowNodeInstance **/
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,8)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_flowNodeType")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"enum:FlowNodeType_AbstractTask:FlowNodeType_ServiceTask:FlowNodeType_UserTask:FlowNodeType_ManualTask:FlowNodeType_BusinessRuleTask:FlowNodeType_ScriptTask:FlowNodeType_EmbeddedSubprocess:FlowNodeType_EventSubprocess:FlowNodeType_AdHocSubprocess:FlowNodeType_Transaction:FlowNodeType_CallActivity:FlowNodeType_ParallelGateway:FlowNodeType_ExclusiveGateway:FlowNodeType_InclusiveGateway:FlowNodeType_EventBasedGateway:FlowNodeType_ComplexGateway:FlowNodeType_StartEvent:FlowNodeType_IntermediateCatchEvent:FlowNodeType_BoundaryEvent:FlowNodeType_EndEvent:FlowNodeType_IntermediateThrowEvent")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,9)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_lifecycleState")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"enum:LifecycleState_Completed:LifecycleState_Compensated:LifecycleState_Failed:LifecycleState_Terminated:LifecycleState_Completing:LifecycleState_Compensating:LifecycleState_Failing:LifecycleState_Terminating:LifecycleState_Ready:LifecycleState_Active")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,10)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_inputRef")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject:Ref")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,11)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_outputRef")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject:Ref")

	/** members of EventInstance **/
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,12)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_eventType")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"enum:EventType_StartEvent:EventType_IntermediateCatchEvent:EventType_BoundaryEvent:EventType_EndEvent:EventType_IntermediateThrowEvent")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,13)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,true)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_eventDefintionInstances")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"[]BPMS_Runtime.EventDefinitionInstance")

	/** associations of EventInstance **/
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,14)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,false)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_SubprocessInstancePtr")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"BPMS_Runtime.SubprocessInstance:Ref")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,15)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,false)
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"M_processInstancePtr")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"BPMS_Runtime.ProcessInstance:Ref")
	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"childsUuid")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"[]xs.string")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,16)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,false)

	eventInstanceEntityProto.Fields = append(eventInstanceEntityProto.Fields,"referenced")
	eventInstanceEntityProto.FieldsType = append(eventInstanceEntityProto.FieldsType,"[]EntityRef")
	eventInstanceEntityProto.FieldsOrder = append(eventInstanceEntityProto.FieldsOrder,17)
	eventInstanceEntityProto.FieldsVisibility = append(eventInstanceEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&eventInstanceEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_EventInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.EventInstance"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventInstance"

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
	}else{
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
	dataIds := make([]string,0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(this.object.M_data[i].UUID,this.object.M_data[i])
		dataIds=append(dataIds,dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data",dataEntity)
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
	if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AbstractTask{
		EventInstanceInfo = append(EventInstanceInfo, 0)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ServiceTask{
		EventInstanceInfo = append(EventInstanceInfo, 1)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_UserTask{
		EventInstanceInfo = append(EventInstanceInfo, 2)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ManualTask{
		EventInstanceInfo = append(EventInstanceInfo, 3)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BusinessRuleTask{
		EventInstanceInfo = append(EventInstanceInfo, 4)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ScriptTask{
		EventInstanceInfo = append(EventInstanceInfo, 5)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EmbeddedSubprocess{
		EventInstanceInfo = append(EventInstanceInfo, 6)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventSubprocess{
		EventInstanceInfo = append(EventInstanceInfo, 7)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AdHocSubprocess{
		EventInstanceInfo = append(EventInstanceInfo, 8)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_Transaction{
		EventInstanceInfo = append(EventInstanceInfo, 9)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_CallActivity{
		EventInstanceInfo = append(EventInstanceInfo, 10)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ParallelGateway{
		EventInstanceInfo = append(EventInstanceInfo, 11)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ExclusiveGateway{
		EventInstanceInfo = append(EventInstanceInfo, 12)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_InclusiveGateway{
		EventInstanceInfo = append(EventInstanceInfo, 13)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventBasedGateway{
		EventInstanceInfo = append(EventInstanceInfo, 14)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ComplexGateway{
		EventInstanceInfo = append(EventInstanceInfo, 15)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_StartEvent{
		EventInstanceInfo = append(EventInstanceInfo, 16)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateCatchEvent{
		EventInstanceInfo = append(EventInstanceInfo, 17)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BoundaryEvent{
		EventInstanceInfo = append(EventInstanceInfo, 18)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EndEvent{
		EventInstanceInfo = append(EventInstanceInfo, 19)
	} else if this.object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateThrowEvent{
		EventInstanceInfo = append(EventInstanceInfo, 20)
	}else{
		EventInstanceInfo = append(EventInstanceInfo, 0)
	}

	/** Save lifecycleState type LifecycleState **/
	if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completed{
		EventInstanceInfo = append(EventInstanceInfo, 0)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensated{
		EventInstanceInfo = append(EventInstanceInfo, 1)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failed{
		EventInstanceInfo = append(EventInstanceInfo, 2)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminated{
		EventInstanceInfo = append(EventInstanceInfo, 3)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completing{
		EventInstanceInfo = append(EventInstanceInfo, 4)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensating{
		EventInstanceInfo = append(EventInstanceInfo, 5)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failing{
		EventInstanceInfo = append(EventInstanceInfo, 6)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminating{
		EventInstanceInfo = append(EventInstanceInfo, 7)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Ready{
		EventInstanceInfo = append(EventInstanceInfo, 8)
	} else if this.object.M_lifecycleState==BPMS_Runtime.LifecycleState_Active{
		EventInstanceInfo = append(EventInstanceInfo, 9)
	}else{
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
	if this.object.M_eventType==BPMS_Runtime.EventType_StartEvent{
		EventInstanceInfo = append(EventInstanceInfo, 0)
	} else if this.object.M_eventType==BPMS_Runtime.EventType_IntermediateCatchEvent{
		EventInstanceInfo = append(EventInstanceInfo, 1)
	} else if this.object.M_eventType==BPMS_Runtime.EventType_BoundaryEvent{
		EventInstanceInfo = append(EventInstanceInfo, 2)
	} else if this.object.M_eventType==BPMS_Runtime.EventType_EndEvent{
		EventInstanceInfo = append(EventInstanceInfo, 3)
	} else if this.object.M_eventType==BPMS_Runtime.EventType_IntermediateThrowEvent{
		EventInstanceInfo = append(EventInstanceInfo, 4)
	}else{
		EventInstanceInfo = append(EventInstanceInfo, 0)
	}

	/** Save eventDefintionInstances type EventDefinitionInstance **/
	eventDefintionInstancesIds := make([]string,0)
	for i := 0; i < len(this.object.M_eventDefintionInstances); i++ {
		eventDefintionInstancesEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeEventDefinitionInstanceEntity(this.object.M_eventDefintionInstances[i].UUID,this.object.M_eventDefintionInstances[i])
		eventDefintionInstancesIds=append(eventDefintionInstancesIds,eventDefintionInstancesEntity.uuid)
		eventDefintionInstancesEntity.AppendReferenced("eventDefintionInstances", this)
		this.AppendChild("eventDefintionInstances",eventDefintionInstancesEntity)
		if eventDefintionInstancesEntity.NeedSave() {
			eventDefintionInstancesEntity.SaveEntity()
		}
	}
	eventDefintionInstancesStr, _ := json.Marshal(eventDefintionInstancesIds)
	EventInstanceInfo = append(EventInstanceInfo, string(eventDefintionInstancesStr))

	/** associations of EventInstance **/

	/** Save SubprocessInstance type SubprocessInstance **/
		EventInstanceInfo = append(EventInstanceInfo,this.object.M_SubprocessInstancePtr)

	/** Save processInstance type ProcessInstance **/
		EventInstanceInfo = append(EventInstanceInfo,this.object.M_processInstancePtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), EventInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), EventInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_EventInstanceEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_EventInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventInstance"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of EventInstance...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.EventInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.EventInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** bpmnElementId **/
 		if results[0][3] != nil{
 			this.object.M_bpmnElementId=results[0][3].(string)
 		}

		/** participants **/
 		if results[0][4] != nil{
 			this.object.M_participants= append(this.object.M_participants, results[0][4].([]string) ...)
 		}

		/** dataRef **/
 		if results[0][5] != nil{
			idsStr :=results[0][5].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ItemAwareElementInstance"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef,ids[i])
					GetServer().GetEntityManager().appendReference("dataRef",this.object.UUID, id_)
				}
			}
 		}

		/** data **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_Runtime_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_Runtime_ItemAwareElementInstanceEntity)
					}else{
						dataEntity = GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data",dataEntity)
				}
 			}
 		}

		/** logInfoRef **/
 		if results[0][7] != nil{
			idsStr :=results[0][7].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.LogInfo"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef,ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of FlowNodeInstance **/

		/** flowNodeType **/
 		if results[0][8] != nil{
 			enumIndex := results[0][8].(int)
			if enumIndex == 0{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AbstractTask
			} else if enumIndex == 1{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ServiceTask
			} else if enumIndex == 2{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_UserTask
			} else if enumIndex == 3{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ManualTask
			} else if enumIndex == 4{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BusinessRuleTask
			} else if enumIndex == 5{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ScriptTask
			} else if enumIndex == 6{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EmbeddedSubprocess
			} else if enumIndex == 7{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventSubprocess
			} else if enumIndex == 8{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AdHocSubprocess
			} else if enumIndex == 9{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_Transaction
			} else if enumIndex == 10{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_CallActivity
			} else if enumIndex == 11{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ParallelGateway
			} else if enumIndex == 12{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ExclusiveGateway
			} else if enumIndex == 13{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_InclusiveGateway
			} else if enumIndex == 14{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventBasedGateway
			} else if enumIndex == 15{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ComplexGateway
			} else if enumIndex == 16{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_StartEvent
			} else if enumIndex == 17{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateCatchEvent
			} else if enumIndex == 18{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BoundaryEvent
			} else if enumIndex == 19{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EndEvent
			} else if enumIndex == 20{
 				this.object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateThrowEvent
 			}
 		}

		/** lifecycleState **/
 		if results[0][9] != nil{
 			enumIndex := results[0][9].(int)
			if enumIndex == 0{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completed
			} else if enumIndex == 1{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensated
			} else if enumIndex == 2{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failed
			} else if enumIndex == 3{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminated
			} else if enumIndex == 4{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completing
			} else if enumIndex == 5{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensating
			} else if enumIndex == 6{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failing
			} else if enumIndex == 7{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminating
			} else if enumIndex == 8{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Ready
			} else if enumIndex == 9{
 				this.object.M_lifecycleState=BPMS_Runtime.LifecycleState_Active
 			}
 		}

		/** inputRef **/
 		if results[0][10] != nil{
			idsStr :=results[0][10].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ConnectingObject"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_inputRef = append(this.object.M_inputRef,ids[i])
					GetServer().GetEntityManager().appendReference("inputRef",this.object.UUID, id_)
				}
			}
 		}

		/** outputRef **/
 		if results[0][11] != nil{
			idsStr :=results[0][11].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ConnectingObject"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_outputRef = append(this.object.M_outputRef,ids[i])
					GetServer().GetEntityManager().appendReference("outputRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of EventInstance **/

		/** eventType **/
 		if results[0][12] != nil{
 			enumIndex := results[0][12].(int)
			if enumIndex == 0{
 				this.object.M_eventType=BPMS_Runtime.EventType_StartEvent
			} else if enumIndex == 1{
 				this.object.M_eventType=BPMS_Runtime.EventType_IntermediateCatchEvent
			} else if enumIndex == 2{
 				this.object.M_eventType=BPMS_Runtime.EventType_BoundaryEvent
			} else if enumIndex == 3{
 				this.object.M_eventType=BPMS_Runtime.EventType_EndEvent
			} else if enumIndex == 4{
 				this.object.M_eventType=BPMS_Runtime.EventType_IntermediateThrowEvent
 			}
 		}

		/** eventDefintionInstances **/
 		if results[0][13] != nil{
			uuidsStr :=results[0][13].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var eventDefintionInstancesEntity *BPMS_Runtime_EventDefinitionInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						eventDefintionInstancesEntity = instance.(*BPMS_Runtime_EventDefinitionInstanceEntity)
					}else{
						eventDefintionInstancesEntity = GetServer().GetEntityManager().NewBPMS_RuntimeEventDefinitionInstanceEntity(uuids[i], nil)
						eventDefintionInstancesEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(eventDefintionInstancesEntity)
					}
					eventDefintionInstancesEntity.AppendReferenced("eventDefintionInstances", this)
					this.AppendChild("eventDefintionInstances",eventDefintionInstancesEntity)
				}
 			}
 		}

		/** associations of EventInstance **/

		/** SubprocessInstancePtr **/
 		if results[0][14] != nil{
			id :=results[0][14].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.SubprocessInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_SubprocessInstancePtr= id
				GetServer().GetEntityManager().appendReference("SubprocessInstancePtr",this.object.UUID, id_)
			}
 		}

		/** processInstancePtr **/
 		if results[0][15] != nil{
			id :=results[0][15].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.ProcessInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_processInstancePtr= id
				GetServer().GetEntityManager().appendReference("processInstancePtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeEventInstanceEntityFromObject(object *BPMS_Runtime.EventInstance) *BPMS_Runtime_EventInstanceEntity {
	 return this.NewBPMS_RuntimeEventInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_EventInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeEventInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_EventInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_EventInstanceEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			DefinitionsInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_DefinitionsInstanceEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.DefinitionsInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeDefinitionsInstanceEntity(objectId string, object interface{}) *BPMS_Runtime_DefinitionsInstanceEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeDefinitionsInstanceExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.DefinitionsInstance).TYPENAME = "BPMS_Runtime.DefinitionsInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.DefinitionsInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_DefinitionsInstanceEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.DefinitionsInstance%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_DefinitionsInstanceEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.DefinitionsInstance)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.DefinitionsInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.DefinitionsInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.DefinitionsInstance","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetTypeName()string{
	return "BPMS_Runtime.DefinitionsInstance"
}
func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_DefinitionsInstanceEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_DefinitionsInstanceEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.DefinitionsInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_DefinitionsInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_DefinitionsInstanceEntityPrototype() {

	var definitionsInstanceEntityProto EntityPrototype
	definitionsInstanceEntityProto.TypeName = "BPMS_Runtime.DefinitionsInstance"
	definitionsInstanceEntityProto.SuperTypeNames = append(definitionsInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.Instance")
	definitionsInstanceEntityProto.Ids = append(definitionsInstanceEntityProto.Ids,"uuid")
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"uuid")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"xs.string")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,0)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,false)
	definitionsInstanceEntityProto.Indexs = append(definitionsInstanceEntityProto.Indexs,"parentUuid")
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"parentUuid")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"xs.string")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,1)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,false)

	/** members of Instance **/
	definitionsInstanceEntityProto.Ids = append(definitionsInstanceEntityProto.Ids,"M_id")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,2)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"M_id")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"xs.ID")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,3)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"M_bpmnElementId")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"xs.string")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,4)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"M_participants")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"[]xs.string")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,5)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"M_dataRef")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,6)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"M_data")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,7)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"M_logInfoRef")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo:Ref")

	/** members of DefinitionsInstance **/
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,8)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,true)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"M_processInstances")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ProcessInstance")

	/** associations of DefinitionsInstance **/
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,9)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,false)
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"M_parentPtr")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"BPMS_Runtime.Runtimes:Ref")
	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"childsUuid")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"[]xs.string")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,10)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,false)

	definitionsInstanceEntityProto.Fields = append(definitionsInstanceEntityProto.Fields,"referenced")
	definitionsInstanceEntityProto.FieldsType = append(definitionsInstanceEntityProto.FieldsType,"[]EntityRef")
	definitionsInstanceEntityProto.FieldsOrder = append(definitionsInstanceEntityProto.FieldsOrder,11)
	definitionsInstanceEntityProto.FieldsVisibility = append(definitionsInstanceEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&definitionsInstanceEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_DefinitionsInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.DefinitionsInstance"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.DefinitionsInstance"

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
	}else{
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
	dataIds := make([]string,0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(this.object.M_data[i].UUID,this.object.M_data[i])
		dataIds=append(dataIds,dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data",dataEntity)
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
	processInstancesIds := make([]string,0)
	for i := 0; i < len(this.object.M_processInstances); i++ {
		processInstancesEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeProcessInstanceEntity(this.object.M_processInstances[i].UUID,this.object.M_processInstances[i])
		processInstancesIds=append(processInstancesIds,processInstancesEntity.uuid)
		processInstancesEntity.AppendReferenced("processInstances", this)
		this.AppendChild("processInstances",processInstancesEntity)
		if processInstancesEntity.NeedSave() {
			processInstancesEntity.SaveEntity()
		}
	}
	processInstancesStr, _ := json.Marshal(processInstancesIds)
	DefinitionsInstanceInfo = append(DefinitionsInstanceInfo, string(processInstancesStr))

	/** associations of DefinitionsInstance **/

	/** Save parent type Runtimes **/
		DefinitionsInstanceInfo = append(DefinitionsInstanceInfo,this.object.M_parentPtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), DefinitionsInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), DefinitionsInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_DefinitionsInstanceEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_DefinitionsInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.DefinitionsInstance"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of DefinitionsInstance...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.DefinitionsInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.DefinitionsInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** bpmnElementId **/
 		if results[0][3] != nil{
 			this.object.M_bpmnElementId=results[0][3].(string)
 		}

		/** participants **/
 		if results[0][4] != nil{
 			this.object.M_participants= append(this.object.M_participants, results[0][4].([]string) ...)
 		}

		/** dataRef **/
 		if results[0][5] != nil{
			idsStr :=results[0][5].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ItemAwareElementInstance"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef,ids[i])
					GetServer().GetEntityManager().appendReference("dataRef",this.object.UUID, id_)
				}
			}
 		}

		/** data **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_Runtime_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_Runtime_ItemAwareElementInstanceEntity)
					}else{
						dataEntity = GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data",dataEntity)
				}
 			}
 		}

		/** logInfoRef **/
 		if results[0][7] != nil{
			idsStr :=results[0][7].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.LogInfo"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef,ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of DefinitionsInstance **/

		/** processInstances **/
 		if results[0][8] != nil{
			uuidsStr :=results[0][8].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var processInstancesEntity *BPMS_Runtime_ProcessInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						processInstancesEntity = instance.(*BPMS_Runtime_ProcessInstanceEntity)
					}else{
						processInstancesEntity = GetServer().GetEntityManager().NewBPMS_RuntimeProcessInstanceEntity(uuids[i], nil)
						processInstancesEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(processInstancesEntity)
					}
					processInstancesEntity.AppendReferenced("processInstances", this)
					this.AppendChild("processInstances",processInstancesEntity)
				}
 			}
 		}

		/** associations of DefinitionsInstance **/

		/** parentPtr **/
 		if results[0][9] != nil{
			id :=results[0][9].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.Runtimes"
				id_:= refTypeName + "$$" + id
				this.object.M_parentPtr= id
				GetServer().GetEntityManager().appendReference("parentPtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeDefinitionsInstanceEntityFromObject(object *BPMS_Runtime.DefinitionsInstance) *BPMS_Runtime_DefinitionsInstanceEntity {
	 return this.NewBPMS_RuntimeDefinitionsInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_DefinitionsInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeDefinitionsInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.DefinitionsInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_DefinitionsInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_DefinitionsInstanceEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			ProcessInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_ProcessInstanceEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.ProcessInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeProcessInstanceEntity(objectId string, object interface{}) *BPMS_Runtime_ProcessInstanceEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeProcessInstanceExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.ProcessInstance).TYPENAME = "BPMS_Runtime.ProcessInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.ProcessInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_ProcessInstanceEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.ProcessInstance%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_ProcessInstanceEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.ProcessInstance)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.ProcessInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.ProcessInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.ProcessInstance","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_ProcessInstanceEntity) GetTypeName()string{
	return "BPMS_Runtime.ProcessInstance"
}
func(this *BPMS_Runtime_ProcessInstanceEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_ProcessInstanceEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_ProcessInstanceEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_ProcessInstanceEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_ProcessInstanceEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_ProcessInstanceEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_ProcessInstanceEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_ProcessInstanceEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_ProcessInstanceEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_ProcessInstanceEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_ProcessInstanceEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_ProcessInstanceEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_ProcessInstanceEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_ProcessInstanceEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_ProcessInstanceEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_ProcessInstanceEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_ProcessInstanceEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_ProcessInstanceEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_ProcessInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_ProcessInstanceEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_ProcessInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_ProcessInstanceEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_ProcessInstanceEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ProcessInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_ProcessInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_ProcessInstanceEntityPrototype() {

	var processInstanceEntityProto EntityPrototype
	processInstanceEntityProto.TypeName = "BPMS_Runtime.ProcessInstance"
	processInstanceEntityProto.SuperTypeNames = append(processInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.Instance")
	processInstanceEntityProto.Ids = append(processInstanceEntityProto.Ids,"uuid")
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"uuid")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,0)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,false)
	processInstanceEntityProto.Indexs = append(processInstanceEntityProto.Indexs,"parentUuid")
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"parentUuid")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,1)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,false)

	/** members of Instance **/
	processInstanceEntityProto.Ids = append(processInstanceEntityProto.Ids,"M_id")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,2)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_id")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"xs.ID")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,3)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_bpmnElementId")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,4)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_participants")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"[]xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,5)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_dataRef")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,6)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_data")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,7)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_logInfoRef")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo:Ref")

	/** members of ProcessInstance **/
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,8)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_number")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"xs.int")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,9)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_colorName")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,10)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_colorNumber")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,11)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_flowNodeInstances")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"[]BPMS_Runtime.FlowNodeInstance")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,12)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,true)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_connectingObjects")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ConnectingObject")

	/** associations of ProcessInstance **/
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,13)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,false)
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"M_parentPtr")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"BPMS_Runtime.DefinitionsInstance:Ref")
	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"childsUuid")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"[]xs.string")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,14)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,false)

	processInstanceEntityProto.Fields = append(processInstanceEntityProto.Fields,"referenced")
	processInstanceEntityProto.FieldsType = append(processInstanceEntityProto.FieldsType,"[]EntityRef")
	processInstanceEntityProto.FieldsOrder = append(processInstanceEntityProto.FieldsOrder,15)
	processInstanceEntityProto.FieldsVisibility = append(processInstanceEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&processInstanceEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_ProcessInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.ProcessInstance"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ProcessInstance"

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
	}else{
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
	dataIds := make([]string,0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(this.object.M_data[i].UUID,this.object.M_data[i])
		dataIds=append(dataIds,dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data",dataEntity)
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
	flowNodeInstancesIds := make([]string,0)
	for i := 0; i < len(this.object.M_flowNodeInstances); i++ {
		switch v := this.object.M_flowNodeInstances[i].(type) {
		case *BPMS_Runtime.ActivityInstance:
		flowNodeInstancesEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeActivityInstanceEntity(v.UUID, v)
		flowNodeInstancesIds=append(flowNodeInstancesIds,flowNodeInstancesEntity.uuid)
		flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
		this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
		if flowNodeInstancesEntity.NeedSave() {
			flowNodeInstancesEntity.SaveEntity()
		}
		case *BPMS_Runtime.SubprocessInstance:
		flowNodeInstancesEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeSubprocessInstanceEntity(v.UUID, v)
		flowNodeInstancesIds=append(flowNodeInstancesIds,flowNodeInstancesEntity.uuid)
		flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
		this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
		if flowNodeInstancesEntity.NeedSave() {
			flowNodeInstancesEntity.SaveEntity()
		}
		case *BPMS_Runtime.GatewayInstance:
		flowNodeInstancesEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeGatewayInstanceEntity(v.UUID, v)
		flowNodeInstancesIds=append(flowNodeInstancesIds,flowNodeInstancesEntity.uuid)
		flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
		this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
		if flowNodeInstancesEntity.NeedSave() {
			flowNodeInstancesEntity.SaveEntity()
		}
		case *BPMS_Runtime.EventInstance:
		flowNodeInstancesEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeEventInstanceEntity(v.UUID, v)
		flowNodeInstancesIds=append(flowNodeInstancesIds,flowNodeInstancesEntity.uuid)
		flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
		this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
		if flowNodeInstancesEntity.NeedSave() {
			flowNodeInstancesEntity.SaveEntity()
		}
		}
	}
	flowNodeInstancesStr, _ := json.Marshal(flowNodeInstancesIds)
	ProcessInstanceInfo = append(ProcessInstanceInfo, string(flowNodeInstancesStr))

	/** Save connectingObjects type ConnectingObject **/
	connectingObjectsIds := make([]string,0)
	for i := 0; i < len(this.object.M_connectingObjects); i++ {
		connectingObjectsEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeConnectingObjectEntity(this.object.M_connectingObjects[i].UUID,this.object.M_connectingObjects[i])
		connectingObjectsIds=append(connectingObjectsIds,connectingObjectsEntity.uuid)
		connectingObjectsEntity.AppendReferenced("connectingObjects", this)
		this.AppendChild("connectingObjects",connectingObjectsEntity)
		if connectingObjectsEntity.NeedSave() {
			connectingObjectsEntity.SaveEntity()
		}
	}
	connectingObjectsStr, _ := json.Marshal(connectingObjectsIds)
	ProcessInstanceInfo = append(ProcessInstanceInfo, string(connectingObjectsStr))

	/** associations of ProcessInstance **/

	/** Save parent type DefinitionsInstance **/
		ProcessInstanceInfo = append(ProcessInstanceInfo,this.object.M_parentPtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), ProcessInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), ProcessInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_ProcessInstanceEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_ProcessInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ProcessInstance"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ProcessInstance...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.ProcessInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.ProcessInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** bpmnElementId **/
 		if results[0][3] != nil{
 			this.object.M_bpmnElementId=results[0][3].(string)
 		}

		/** participants **/
 		if results[0][4] != nil{
 			this.object.M_participants= append(this.object.M_participants, results[0][4].([]string) ...)
 		}

		/** dataRef **/
 		if results[0][5] != nil{
			idsStr :=results[0][5].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ItemAwareElementInstance"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef,ids[i])
					GetServer().GetEntityManager().appendReference("dataRef",this.object.UUID, id_)
				}
			}
 		}

		/** data **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_Runtime_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_Runtime_ItemAwareElementInstanceEntity)
					}else{
						dataEntity = GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data",dataEntity)
				}
 			}
 		}

		/** logInfoRef **/
 		if results[0][7] != nil{
			idsStr :=results[0][7].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.LogInfo"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef,ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of ProcessInstance **/

		/** number **/
 		if results[0][8] != nil{
 			this.object.M_number=results[0][8].(int)
 		}

		/** colorName **/
 		if results[0][9] != nil{
 			this.object.M_colorName=results[0][9].(string)
 		}

		/** colorNumber **/
 		if results[0][10] != nil{
 			this.object.M_colorNumber=results[0][10].(string)
 		}

		/** flowNodeInstances **/
 		if results[0][11] != nil{
			uuidsStr :=results[0][11].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
			typeName := uuids[i][0:strings.Index(uuids[i], "%")]
			if err!=nil{
				log.Println("type ", typeName, " not found!")
				return err
			}
				if typeName == "BPMS_Runtime.SubprocessInstance"{
						if len(uuids[i]) > 0 {
							var flowNodeInstancesEntity *BPMS_Runtime_SubprocessInstanceEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								flowNodeInstancesEntity = instance.(*BPMS_Runtime_SubprocessInstanceEntity)
							}else{
								flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMS_RuntimeSubprocessInstanceEntity(uuids[i], nil)
								flowNodeInstancesEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
							}
							flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
							this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
						}
				} else if typeName == "BPMS_Runtime.GatewayInstance"{
						if len(uuids[i]) > 0 {
							var flowNodeInstancesEntity *BPMS_Runtime_GatewayInstanceEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								flowNodeInstancesEntity = instance.(*BPMS_Runtime_GatewayInstanceEntity)
							}else{
								flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMS_RuntimeGatewayInstanceEntity(uuids[i], nil)
								flowNodeInstancesEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
							}
							flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
							this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
						}
				} else if typeName == "BPMS_Runtime.EventInstance"{
						if len(uuids[i]) > 0 {
							var flowNodeInstancesEntity *BPMS_Runtime_EventInstanceEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								flowNodeInstancesEntity = instance.(*BPMS_Runtime_EventInstanceEntity)
							}else{
								flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMS_RuntimeEventInstanceEntity(uuids[i], nil)
								flowNodeInstancesEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
							}
							flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
							this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
						}
				} else if typeName == "BPMS_Runtime.ActivityInstance"{
						if len(uuids[i]) > 0 {
							var flowNodeInstancesEntity *BPMS_Runtime_ActivityInstanceEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								flowNodeInstancesEntity = instance.(*BPMS_Runtime_ActivityInstanceEntity)
							}else{
								flowNodeInstancesEntity = GetServer().GetEntityManager().NewBPMS_RuntimeActivityInstanceEntity(uuids[i], nil)
								flowNodeInstancesEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(flowNodeInstancesEntity)
							}
							flowNodeInstancesEntity.AppendReferenced("flowNodeInstances", this)
							this.AppendChild("flowNodeInstances",flowNodeInstancesEntity)
						}
				}
 			}
 		}

		/** connectingObjects **/
 		if results[0][12] != nil{
			uuidsStr :=results[0][12].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var connectingObjectsEntity *BPMS_Runtime_ConnectingObjectEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						connectingObjectsEntity = instance.(*BPMS_Runtime_ConnectingObjectEntity)
					}else{
						connectingObjectsEntity = GetServer().GetEntityManager().NewBPMS_RuntimeConnectingObjectEntity(uuids[i], nil)
						connectingObjectsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(connectingObjectsEntity)
					}
					connectingObjectsEntity.AppendReferenced("connectingObjects", this)
					this.AppendChild("connectingObjects",connectingObjectsEntity)
				}
 			}
 		}

		/** associations of ProcessInstance **/

		/** parentPtr **/
 		if results[0][13] != nil{
			id :=results[0][13].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.DefinitionsInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_parentPtr= id
				GetServer().GetEntityManager().appendReference("parentPtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeProcessInstanceEntityFromObject(object *BPMS_Runtime.ProcessInstance) *BPMS_Runtime_ProcessInstanceEntity {
	 return this.NewBPMS_RuntimeProcessInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_ProcessInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeProcessInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ProcessInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_ProcessInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_ProcessInstanceEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			EventDefinitionInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_EventDefinitionInstanceEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.EventDefinitionInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeEventDefinitionInstanceEntity(objectId string, object interface{}) *BPMS_Runtime_EventDefinitionInstanceEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeEventDefinitionInstanceExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.EventDefinitionInstance).TYPENAME = "BPMS_Runtime.EventDefinitionInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.EventDefinitionInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_EventDefinitionInstanceEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.EventDefinitionInstance%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_EventDefinitionInstanceEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.EventDefinitionInstance)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.EventDefinitionInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.EventDefinitionInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.EventDefinitionInstance","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetTypeName()string{
	return "BPMS_Runtime.EventDefinitionInstance"
}
func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_EventDefinitionInstanceEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_EventDefinitionInstanceEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventDefinitionInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_EventDefinitionInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_EventDefinitionInstanceEntityPrototype() {

	var eventDefinitionInstanceEntityProto EntityPrototype
	eventDefinitionInstanceEntityProto.TypeName = "BPMS_Runtime.EventDefinitionInstance"
	eventDefinitionInstanceEntityProto.SuperTypeNames = append(eventDefinitionInstanceEntityProto.SuperTypeNames, "BPMS_Runtime.Instance")
	eventDefinitionInstanceEntityProto.Ids = append(eventDefinitionInstanceEntityProto.Ids,"uuid")
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"uuid")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"xs.string")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,0)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,false)
	eventDefinitionInstanceEntityProto.Indexs = append(eventDefinitionInstanceEntityProto.Indexs,"parentUuid")
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"parentUuid")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"xs.string")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,1)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,false)

	/** members of Instance **/
	eventDefinitionInstanceEntityProto.Ids = append(eventDefinitionInstanceEntityProto.Ids,"M_id")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,2)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"M_id")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"xs.ID")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,3)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"M_bpmnElementId")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"xs.string")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,4)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"M_participants")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"[]xs.string")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,5)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"M_dataRef")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,6)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"M_data")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,7)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"M_logInfoRef")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo:Ref")

	/** members of EventDefinitionInstance **/
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,8)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,true)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"M_eventDefinitionType")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"enum:EventDefinitionType_MessageEventDefinition:EventDefinitionType_LinkEventDefinition:EventDefinitionType_ErrorEventDefinition:EventDefinitionType_TerminateEventDefinition:EventDefinitionType_CompensationEventDefinition:EventDefinitionType_ConditionalEventDefinition:EventDefinitionType_TimerEventDefinition:EventDefinitionType_CancelEventDefinition:EventDefinitionType_EscalationEventDefinition")

	/** associations of EventDefinitionInstance **/
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,9)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,false)
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"M_eventInstancePtr")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"BPMS_Runtime.EventInstance:Ref")
	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"childsUuid")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"[]xs.string")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,10)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,false)

	eventDefinitionInstanceEntityProto.Fields = append(eventDefinitionInstanceEntityProto.Fields,"referenced")
	eventDefinitionInstanceEntityProto.FieldsType = append(eventDefinitionInstanceEntityProto.FieldsType,"[]EntityRef")
	eventDefinitionInstanceEntityProto.FieldsOrder = append(eventDefinitionInstanceEntityProto.FieldsOrder,11)
	eventDefinitionInstanceEntityProto.FieldsVisibility = append(eventDefinitionInstanceEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&eventDefinitionInstanceEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_EventDefinitionInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.EventDefinitionInstance"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventDefinitionInstance"

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
	}else{
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
	dataIds := make([]string,0)
	for i := 0; i < len(this.object.M_data); i++ {
		dataEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(this.object.M_data[i].UUID,this.object.M_data[i])
		dataIds=append(dataIds,dataEntity.uuid)
		dataEntity.AppendReferenced("data", this)
		this.AppendChild("data",dataEntity)
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
	if this.object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_MessageEventDefinition{
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 0)
	} else if this.object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_LinkEventDefinition{
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 1)
	} else if this.object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_ErrorEventDefinition{
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 2)
	} else if this.object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_TerminateEventDefinition{
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 3)
	} else if this.object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_CompensationEventDefinition{
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 4)
	} else if this.object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_ConditionalEventDefinition{
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 5)
	} else if this.object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_TimerEventDefinition{
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 6)
	} else if this.object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_CancelEventDefinition{
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 7)
	} else if this.object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_EscalationEventDefinition{
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 8)
	}else{
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo, 0)
	}

	/** associations of EventDefinitionInstance **/

	/** Save eventInstance type EventInstance **/
		EventDefinitionInstanceInfo = append(EventDefinitionInstanceInfo,this.object.M_eventInstancePtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), EventDefinitionInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), EventDefinitionInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_EventDefinitionInstanceEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_EventDefinitionInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventDefinitionInstance"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of EventDefinitionInstance...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.EventDefinitionInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.EventDefinitionInstance"

		this.parentUuid = results[0][1].(string)

		/** members of Instance **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** bpmnElementId **/
 		if results[0][3] != nil{
 			this.object.M_bpmnElementId=results[0][3].(string)
 		}

		/** participants **/
 		if results[0][4] != nil{
 			this.object.M_participants= append(this.object.M_participants, results[0][4].([]string) ...)
 		}

		/** dataRef **/
 		if results[0][5] != nil{
			idsStr :=results[0][5].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ItemAwareElementInstance"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef,ids[i])
					GetServer().GetEntityManager().appendReference("dataRef",this.object.UUID, id_)
				}
			}
 		}

		/** data **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var dataEntity *BPMS_Runtime_ItemAwareElementInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						dataEntity = instance.(*BPMS_Runtime_ItemAwareElementInstanceEntity)
					}else{
						dataEntity = GetServer().GetEntityManager().NewBPMS_RuntimeItemAwareElementInstanceEntity(uuids[i], nil)
						dataEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(dataEntity)
					}
					dataEntity.AppendReferenced("data", this)
					this.AppendChild("data",dataEntity)
				}
 			}
 		}

		/** logInfoRef **/
 		if results[0][7] != nil{
			idsStr :=results[0][7].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.LogInfo"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_logInfoRef = append(this.object.M_logInfoRef,ids[i])
					GetServer().GetEntityManager().appendReference("logInfoRef",this.object.UUID, id_)
				}
			}
 		}

		/** members of EventDefinitionInstance **/

		/** eventDefinitionType **/
 		if results[0][8] != nil{
 			enumIndex := results[0][8].(int)
			if enumIndex == 0{
 				this.object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_MessageEventDefinition
			} else if enumIndex == 1{
 				this.object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_LinkEventDefinition
			} else if enumIndex == 2{
 				this.object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_ErrorEventDefinition
			} else if enumIndex == 3{
 				this.object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_TerminateEventDefinition
			} else if enumIndex == 4{
 				this.object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_CompensationEventDefinition
			} else if enumIndex == 5{
 				this.object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_ConditionalEventDefinition
			} else if enumIndex == 6{
 				this.object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_TimerEventDefinition
			} else if enumIndex == 7{
 				this.object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_CancelEventDefinition
			} else if enumIndex == 8{
 				this.object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_EscalationEventDefinition
 			}
 		}

		/** associations of EventDefinitionInstance **/

		/** eventInstancePtr **/
 		if results[0][9] != nil{
			id :=results[0][9].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.EventInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_eventInstancePtr= id
				GetServer().GetEntityManager().appendReference("eventInstancePtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeEventDefinitionInstanceEntityFromObject(object *BPMS_Runtime.EventDefinitionInstance) *BPMS_Runtime_EventDefinitionInstanceEntity {
	 return this.NewBPMS_RuntimeEventDefinitionInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_EventDefinitionInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeEventDefinitionInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventDefinitionInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_EventDefinitionInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_EventDefinitionInstanceEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			CorrelationInfo
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_CorrelationInfoEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.CorrelationInfo
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeCorrelationInfoEntity(objectId string, object interface{}) *BPMS_Runtime_CorrelationInfoEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeCorrelationInfoExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.CorrelationInfo).TYPENAME = "BPMS_Runtime.CorrelationInfo"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.CorrelationInfo).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_CorrelationInfoEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.CorrelationInfo%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_CorrelationInfoEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.CorrelationInfo)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.CorrelationInfo)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.CorrelationInfo"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.CorrelationInfo","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_CorrelationInfoEntity) GetTypeName()string{
	return "BPMS_Runtime.CorrelationInfo"
}
func(this *BPMS_Runtime_CorrelationInfoEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_CorrelationInfoEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_CorrelationInfoEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_CorrelationInfoEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_CorrelationInfoEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_CorrelationInfoEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_CorrelationInfoEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_CorrelationInfoEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_CorrelationInfoEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_CorrelationInfoEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_CorrelationInfoEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_CorrelationInfoEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_CorrelationInfoEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_CorrelationInfoEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_CorrelationInfoEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_CorrelationInfoEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_CorrelationInfoEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_CorrelationInfoEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_CorrelationInfoEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_CorrelationInfoEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_CorrelationInfoEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_CorrelationInfoEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_CorrelationInfoEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.CorrelationInfo"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_CorrelationInfoEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_CorrelationInfoEntityPrototype() {

	var correlationInfoEntityProto EntityPrototype
	correlationInfoEntityProto.TypeName = "BPMS_Runtime.CorrelationInfo"
	correlationInfoEntityProto.Ids = append(correlationInfoEntityProto.Ids,"uuid")
	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields,"uuid")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType,"xs.string")
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder,0)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility,false)
	correlationInfoEntityProto.Indexs = append(correlationInfoEntityProto.Indexs,"parentUuid")
	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields,"parentUuid")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType,"xs.string")
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder,1)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility,false)

	/** members of CorrelationInfo **/
	correlationInfoEntityProto.Ids = append(correlationInfoEntityProto.Ids,"M_id")
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder,2)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility,true)
	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields,"M_id")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType,"xs.ID")

	/** associations of CorrelationInfo **/
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder,3)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility,false)
	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields,"M_runtimesPtr")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType,"BPMS_Runtime.Runtimes:Ref")
	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields,"childsUuid")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType,"[]xs.string")
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder,4)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility,false)

	correlationInfoEntityProto.Fields = append(correlationInfoEntityProto.Fields,"referenced")
	correlationInfoEntityProto.FieldsType = append(correlationInfoEntityProto.FieldsType,"[]EntityRef")
	correlationInfoEntityProto.FieldsOrder = append(correlationInfoEntityProto.FieldsOrder,5)
	correlationInfoEntityProto.FieldsVisibility = append(correlationInfoEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&correlationInfoEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_CorrelationInfoEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.CorrelationInfo"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.CorrelationInfo"

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
	}else{
		CorrelationInfoInfo = append(CorrelationInfoInfo, "")
	}

	/** members of CorrelationInfo **/
	CorrelationInfoInfo = append(CorrelationInfoInfo, this.object.M_id)

	/** associations of CorrelationInfo **/

	/** Save runtimes type Runtimes **/
		CorrelationInfoInfo = append(CorrelationInfoInfo,this.object.M_runtimesPtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), CorrelationInfoInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), CorrelationInfoInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_CorrelationInfoEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_CorrelationInfoEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.CorrelationInfo"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of CorrelationInfo...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.CorrelationInfo)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.CorrelationInfo"

		this.parentUuid = results[0][1].(string)

		/** members of CorrelationInfo **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** associations of CorrelationInfo **/

		/** runtimesPtr **/
 		if results[0][3] != nil{
			id :=results[0][3].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.Runtimes"
				id_:= refTypeName + "$$" + id
				this.object.M_runtimesPtr= id
				GetServer().GetEntityManager().appendReference("runtimesPtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeCorrelationInfoEntityFromObject(object *BPMS_Runtime.CorrelationInfo) *BPMS_Runtime_CorrelationInfoEntity {
	 return this.NewBPMS_RuntimeCorrelationInfoEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_CorrelationInfoEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeCorrelationInfoExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.CorrelationInfo"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_CorrelationInfoEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_CorrelationInfoEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			ItemAwareElementInstance
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_ItemAwareElementInstanceEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.ItemAwareElementInstance
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeItemAwareElementInstanceEntity(objectId string, object interface{}) *BPMS_Runtime_ItemAwareElementInstanceEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeItemAwareElementInstanceExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.ItemAwareElementInstance).TYPENAME = "BPMS_Runtime.ItemAwareElementInstance"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.ItemAwareElementInstance).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_ItemAwareElementInstanceEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.ItemAwareElementInstance%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_ItemAwareElementInstanceEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.ItemAwareElementInstance)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.ItemAwareElementInstance)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.ItemAwareElementInstance"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.ItemAwareElementInstance","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetTypeName()string{
	return "BPMS_Runtime.ItemAwareElementInstance"
}
func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ItemAwareElementInstance"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_ItemAwareElementInstanceEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_ItemAwareElementInstanceEntityPrototype() {

	var itemAwareElementInstanceEntityProto EntityPrototype
	itemAwareElementInstanceEntityProto.TypeName = "BPMS_Runtime.ItemAwareElementInstance"
	itemAwareElementInstanceEntityProto.Ids = append(itemAwareElementInstanceEntityProto.Ids,"uuid")
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields,"uuid")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType,"xs.string")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder,0)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility,false)
	itemAwareElementInstanceEntityProto.Indexs = append(itemAwareElementInstanceEntityProto.Indexs,"parentUuid")
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields,"parentUuid")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType,"xs.string")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder,1)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility,false)

	/** members of ItemAwareElementInstance **/
	itemAwareElementInstanceEntityProto.Ids = append(itemAwareElementInstanceEntityProto.Ids,"M_id")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder,2)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility,true)
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields,"M_id")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType,"xs.ID")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder,3)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility,true)
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields,"M_bpmnElementId")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType,"xs.string")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder,4)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility,true)
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields,"M_data")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType,"xs.[]uint8")

	/** associations of ItemAwareElementInstance **/
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder,5)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility,false)
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields,"M_parentPtr")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType,"BPMS_Runtime.Instance:Ref")
	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields,"childsUuid")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType,"[]xs.string")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder,6)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility,false)

	itemAwareElementInstanceEntityProto.Fields = append(itemAwareElementInstanceEntityProto.Fields,"referenced")
	itemAwareElementInstanceEntityProto.FieldsType = append(itemAwareElementInstanceEntityProto.FieldsType,"[]EntityRef")
	itemAwareElementInstanceEntityProto.FieldsOrder = append(itemAwareElementInstanceEntityProto.FieldsOrder,7)
	itemAwareElementInstanceEntityProto.FieldsVisibility = append(itemAwareElementInstanceEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&itemAwareElementInstanceEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_ItemAwareElementInstanceEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.ItemAwareElementInstance"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ItemAwareElementInstance"

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
	}else{
		ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, "")
	}

	/** members of ItemAwareElementInstance **/
	ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, this.object.M_id)
	ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, this.object.M_bpmnElementId)
	ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo, this.object.M_data)

	/** associations of ItemAwareElementInstance **/

	/** Save parent type Instance **/
		ItemAwareElementInstanceInfo = append(ItemAwareElementInstanceInfo,this.object.M_parentPtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), ItemAwareElementInstanceInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), ItemAwareElementInstanceInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_ItemAwareElementInstanceEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_ItemAwareElementInstanceEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ItemAwareElementInstance"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of ItemAwareElementInstance...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.ItemAwareElementInstance)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.ItemAwareElementInstance"

		this.parentUuid = results[0][1].(string)

		/** members of ItemAwareElementInstance **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** bpmnElementId **/
 		if results[0][3] != nil{
 			this.object.M_bpmnElementId=results[0][3].(string)
 		}

		/** data **/
 		if results[0][4] != nil{
 			this.object.M_data=results[0][4].([]uint8)
 		}

		/** associations of ItemAwareElementInstance **/

		/** parentPtr **/
 		if results[0][5] != nil{
			id :=results[0][5].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.Instance"
				id_:= refTypeName + "$$" + id
				this.object.M_parentPtr= id
				GetServer().GetEntityManager().appendReference("parentPtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeItemAwareElementInstanceEntityFromObject(object *BPMS_Runtime.ItemAwareElementInstance) *BPMS_Runtime_ItemAwareElementInstanceEntity {
	 return this.NewBPMS_RuntimeItemAwareElementInstanceEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_ItemAwareElementInstanceEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeItemAwareElementInstanceExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.ItemAwareElementInstance"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_ItemAwareElementInstanceEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_ItemAwareElementInstanceEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			EventData
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_EventDataEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.EventData
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeEventDataEntity(objectId string, object interface{}) *BPMS_Runtime_EventDataEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeEventDataExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.EventData).TYPENAME = "BPMS_Runtime.EventData"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.EventData).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_EventDataEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.EventData%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_EventDataEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.EventData)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.EventData)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.EventData"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.EventData","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_EventDataEntity) GetTypeName()string{
	return "BPMS_Runtime.EventData"
}
func(this *BPMS_Runtime_EventDataEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_EventDataEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_EventDataEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_EventDataEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_EventDataEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_EventDataEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_EventDataEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_EventDataEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_EventDataEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_EventDataEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_EventDataEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_EventDataEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_EventDataEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_EventDataEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_EventDataEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_EventDataEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_EventDataEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_EventDataEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_EventDataEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_EventDataEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_EventDataEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_EventDataEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_EventDataEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventData"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_EventDataEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_EventDataEntityPrototype() {

	var eventDataEntityProto EntityPrototype
	eventDataEntityProto.TypeName = "BPMS_Runtime.EventData"
	eventDataEntityProto.Ids = append(eventDataEntityProto.Ids,"uuid")
	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields,"uuid")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType,"xs.string")
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder,0)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility,false)
	eventDataEntityProto.Indexs = append(eventDataEntityProto.Indexs,"parentUuid")
	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields,"parentUuid")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType,"xs.string")
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder,1)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility,false)

	/** members of EventData **/
	eventDataEntityProto.Ids = append(eventDataEntityProto.Ids,"M_id")
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder,2)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility,true)
	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields,"M_id")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType,"xs.ID")

	/** associations of EventData **/
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder,3)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility,false)
	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields,"M_triggerPtr")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType,"BPMS_Runtime.Trigger:Ref")
	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields,"childsUuid")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType,"[]xs.string")
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder,4)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility,false)

	eventDataEntityProto.Fields = append(eventDataEntityProto.Fields,"referenced")
	eventDataEntityProto.FieldsType = append(eventDataEntityProto.FieldsType,"[]EntityRef")
	eventDataEntityProto.FieldsOrder = append(eventDataEntityProto.FieldsOrder,5)
	eventDataEntityProto.FieldsVisibility = append(eventDataEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&eventDataEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_EventDataEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.EventData"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventData"

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
	}else{
		EventDataInfo = append(EventDataInfo, "")
	}

	/** members of EventData **/
	EventDataInfo = append(EventDataInfo, this.object.M_id)

	/** associations of EventData **/

	/** Save trigger type Trigger **/
		EventDataInfo = append(EventDataInfo,this.object.M_triggerPtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), EventDataInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), EventDataInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_EventDataEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_EventDataEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventData"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of EventData...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.EventData)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.EventData"

		this.parentUuid = results[0][1].(string)

		/** members of EventData **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** associations of EventData **/

		/** triggerPtr **/
 		if results[0][3] != nil{
			id :=results[0][3].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.Trigger"
				id_:= refTypeName + "$$" + id
				this.object.M_triggerPtr= id
				GetServer().GetEntityManager().appendReference("triggerPtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeEventDataEntityFromObject(object *BPMS_Runtime.EventData) *BPMS_Runtime_EventDataEntity {
	 return this.NewBPMS_RuntimeEventDataEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_EventDataEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeEventDataExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.EventData"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_EventDataEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_EventDataEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			Trigger
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_TriggerEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.Trigger
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeTriggerEntity(objectId string, object interface{}) *BPMS_Runtime_TriggerEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeTriggerExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.Trigger).TYPENAME = "BPMS_Runtime.Trigger"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.Trigger).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_TriggerEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.Trigger%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_TriggerEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.Trigger)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.Trigger)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.Trigger"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.Trigger","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_TriggerEntity) GetTypeName()string{
	return "BPMS_Runtime.Trigger"
}
func(this *BPMS_Runtime_TriggerEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_TriggerEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_TriggerEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_TriggerEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_TriggerEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_TriggerEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_TriggerEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_TriggerEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_TriggerEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_TriggerEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_TriggerEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_TriggerEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_TriggerEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_TriggerEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_TriggerEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_TriggerEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_TriggerEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_TriggerEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_TriggerEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_TriggerEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_TriggerEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_TriggerEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_TriggerEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Trigger"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_TriggerEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_TriggerEntityPrototype() {

	var triggerEntityProto EntityPrototype
	triggerEntityProto.TypeName = "BPMS_Runtime.Trigger"
	triggerEntityProto.Ids = append(triggerEntityProto.Ids,"uuid")
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"uuid")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,0)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,false)
	triggerEntityProto.Indexs = append(triggerEntityProto.Indexs,"parentUuid")
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"parentUuid")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,1)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,false)

	/** members of Trigger **/
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,2)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"M_id")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,3)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"M_processUUID")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,4)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"M_sessionId")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,5)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"M_eventTriggerType")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"enum:EventTriggerType_None:EventTriggerType_Timer:EventTriggerType_Conditional:EventTriggerType_Message:EventTriggerType_Signal:EventTriggerType_Multiple:EventTriggerType_ParallelMultiple:EventTriggerType_Escalation:EventTriggerType_Error:EventTriggerType_Compensation:EventTriggerType_Terminate:EventTriggerType_Cancel:EventTriggerType_Link:EventTriggerType_Start")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,6)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"M_eventDatas")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"[]BPMS_Runtime.EventData")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,7)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"M_dataRef")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"[]BPMS_Runtime.ItemAwareElementInstance:Ref")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,8)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"M_sourceRef")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"BPMS_Runtime.EventDefinitionInstance:Ref")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,9)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,true)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"M_targetRef")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"BPMS_Runtime.FlowNodeInstance:Ref")

	/** associations of Trigger **/
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,10)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,false)
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"M_runtimesPtr")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"BPMS_Runtime.Runtimes:Ref")
	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"childsUuid")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"[]xs.string")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,11)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,false)

	triggerEntityProto.Fields = append(triggerEntityProto.Fields,"referenced")
	triggerEntityProto.FieldsType = append(triggerEntityProto.FieldsType,"[]EntityRef")
	triggerEntityProto.FieldsOrder = append(triggerEntityProto.FieldsOrder,12)
	triggerEntityProto.FieldsVisibility = append(triggerEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&triggerEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_TriggerEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.Trigger"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Trigger"

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
	}else{
		TriggerInfo = append(TriggerInfo, "")
	}

	/** members of Trigger **/
	TriggerInfo = append(TriggerInfo, this.object.M_id)
	TriggerInfo = append(TriggerInfo, this.object.M_processUUID)
	TriggerInfo = append(TriggerInfo, this.object.M_sessionId)

	/** Save eventTriggerType type EventTriggerType **/
	if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_None{
		TriggerInfo = append(TriggerInfo, 0)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Timer{
		TriggerInfo = append(TriggerInfo, 1)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Conditional{
		TriggerInfo = append(TriggerInfo, 2)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Message{
		TriggerInfo = append(TriggerInfo, 3)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Signal{
		TriggerInfo = append(TriggerInfo, 4)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Multiple{
		TriggerInfo = append(TriggerInfo, 5)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_ParallelMultiple{
		TriggerInfo = append(TriggerInfo, 6)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Escalation{
		TriggerInfo = append(TriggerInfo, 7)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Error{
		TriggerInfo = append(TriggerInfo, 8)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Compensation{
		TriggerInfo = append(TriggerInfo, 9)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Terminate{
		TriggerInfo = append(TriggerInfo, 10)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Cancel{
		TriggerInfo = append(TriggerInfo, 11)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Link{
		TriggerInfo = append(TriggerInfo, 12)
	} else if this.object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Start{
		TriggerInfo = append(TriggerInfo, 13)
	}else{
		TriggerInfo = append(TriggerInfo, 0)
	}

	/** Save eventDatas type EventData **/
	eventDatasIds := make([]string,0)
	for i := 0; i < len(this.object.M_eventDatas); i++ {
		eventDatasEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeEventDataEntity(this.object.M_eventDatas[i].UUID,this.object.M_eventDatas[i])
		eventDatasIds=append(eventDatasIds,eventDatasEntity.uuid)
		eventDatasEntity.AppendReferenced("eventDatas", this)
		this.AppendChild("eventDatas",eventDatasEntity)
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
		TriggerInfo = append(TriggerInfo,this.object.M_sourceRef)

	/** Save targetRef type FlowNodeInstance **/
		TriggerInfo = append(TriggerInfo,this.object.M_targetRef)

	/** associations of Trigger **/

	/** Save runtimes type Runtimes **/
		TriggerInfo = append(TriggerInfo,this.object.M_runtimesPtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), TriggerInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), TriggerInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_TriggerEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_TriggerEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Trigger"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Trigger...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.Trigger)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.Trigger"

		this.parentUuid = results[0][1].(string)

		/** members of Trigger **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** processUUID **/
 		if results[0][3] != nil{
 			this.object.M_processUUID=results[0][3].(string)
 		}

		/** sessionId **/
 		if results[0][4] != nil{
 			this.object.M_sessionId=results[0][4].(string)
 		}

		/** eventTriggerType **/
 		if results[0][5] != nil{
 			enumIndex := results[0][5].(int)
			if enumIndex == 0{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_None
			} else if enumIndex == 1{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Timer
			} else if enumIndex == 2{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Conditional
			} else if enumIndex == 3{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Message
			} else if enumIndex == 4{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Signal
			} else if enumIndex == 5{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Multiple
			} else if enumIndex == 6{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_ParallelMultiple
			} else if enumIndex == 7{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Escalation
			} else if enumIndex == 8{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Error
			} else if enumIndex == 9{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Compensation
			} else if enumIndex == 10{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Terminate
			} else if enumIndex == 11{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Cancel
			} else if enumIndex == 12{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Link
			} else if enumIndex == 13{
 				this.object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Start
 			}
 		}

		/** eventDatas **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var eventDatasEntity *BPMS_Runtime_EventDataEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						eventDatasEntity = instance.(*BPMS_Runtime_EventDataEntity)
					}else{
						eventDatasEntity = GetServer().GetEntityManager().NewBPMS_RuntimeEventDataEntity(uuids[i], nil)
						eventDatasEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(eventDatasEntity)
					}
					eventDatasEntity.AppendReferenced("eventDatas", this)
					this.AppendChild("eventDatas",eventDatasEntity)
				}
 			}
 		}

		/** dataRef **/
 		if results[0][7] != nil{
			idsStr :=results[0][7].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMS_Runtime.ItemAwareElementInstance"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_dataRef = append(this.object.M_dataRef,ids[i])
					GetServer().GetEntityManager().appendReference("dataRef",this.object.UUID, id_)
				}
			}
 		}

		/** sourceRef **/
 		if results[0][8] != nil{
			id :=results[0][8].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.EventDefinitionInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_sourceRef= id
				GetServer().GetEntityManager().appendReference("sourceRef",this.object.UUID, id_)
			}
 		}

		/** targetRef **/
 		if results[0][9] != nil{
			id :=results[0][9].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.FlowNodeInstance"
				id_:= refTypeName + "$$" + id
				this.object.M_targetRef= id
				GetServer().GetEntityManager().appendReference("targetRef",this.object.UUID, id_)
			}
 		}

		/** associations of Trigger **/

		/** runtimesPtr **/
 		if results[0][10] != nil{
			id :=results[0][10].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.Runtimes"
				id_:= refTypeName + "$$" + id
				this.object.M_runtimesPtr= id
				GetServer().GetEntityManager().appendReference("runtimesPtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeTriggerEntityFromObject(object *BPMS_Runtime.Trigger) *BPMS_Runtime_TriggerEntity {
	 return this.NewBPMS_RuntimeTriggerEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_TriggerEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeTriggerExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Trigger"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_TriggerEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_TriggerEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			Exception
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_ExceptionEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.Exception
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeExceptionEntity(objectId string, object interface{}) *BPMS_Runtime_ExceptionEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeExceptionExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.Exception).TYPENAME = "BPMS_Runtime.Exception"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.Exception).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_ExceptionEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.Exception%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_ExceptionEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.Exception)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.Exception)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.Exception"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.Exception","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_ExceptionEntity) GetTypeName()string{
	return "BPMS_Runtime.Exception"
}
func(this *BPMS_Runtime_ExceptionEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_ExceptionEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_ExceptionEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_ExceptionEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_ExceptionEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_ExceptionEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_ExceptionEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_ExceptionEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_ExceptionEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_ExceptionEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_ExceptionEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_ExceptionEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_ExceptionEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_ExceptionEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_ExceptionEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_ExceptionEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_ExceptionEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_ExceptionEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_ExceptionEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_ExceptionEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_ExceptionEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_ExceptionEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_ExceptionEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Exception"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_ExceptionEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_ExceptionEntityPrototype() {

	var exceptionEntityProto EntityPrototype
	exceptionEntityProto.TypeName = "BPMS_Runtime.Exception"
	exceptionEntityProto.Ids = append(exceptionEntityProto.Ids,"uuid")
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields,"uuid")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType,"xs.string")
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder,0)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility,false)
	exceptionEntityProto.Indexs = append(exceptionEntityProto.Indexs,"parentUuid")
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields,"parentUuid")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType,"xs.string")
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder,1)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility,false)

	/** members of Exception **/
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder,2)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility,true)
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields,"M_id")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType,"xs.string")
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder,3)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility,true)
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields,"M_exceptionType")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType,"enum:ExceptionType_NoIORuleException:ExceptionType_GatewayException:ExceptionType_NoAvailableOutputSetException:ExceptionType_NotMatchingIOSpecification:ExceptionType_IllegalStartEventException")

	/** associations of Exception **/
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder,4)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility,false)
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields,"M_runtimesPtr")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType,"BPMS_Runtime.Runtimes:Ref")
	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields,"childsUuid")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType,"[]xs.string")
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder,5)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility,false)

	exceptionEntityProto.Fields = append(exceptionEntityProto.Fields,"referenced")
	exceptionEntityProto.FieldsType = append(exceptionEntityProto.FieldsType,"[]EntityRef")
	exceptionEntityProto.FieldsOrder = append(exceptionEntityProto.FieldsOrder,6)
	exceptionEntityProto.FieldsVisibility = append(exceptionEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&exceptionEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_ExceptionEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.Exception"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Exception"

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
	}else{
		ExceptionInfo = append(ExceptionInfo, "")
	}

	/** members of Exception **/
	ExceptionInfo = append(ExceptionInfo, this.object.M_id)

	/** Save exceptionType type ExceptionType **/
	if this.object.M_exceptionType==BPMS_Runtime.ExceptionType_NoIORuleException{
		ExceptionInfo = append(ExceptionInfo, 0)
	} else if this.object.M_exceptionType==BPMS_Runtime.ExceptionType_GatewayException{
		ExceptionInfo = append(ExceptionInfo, 1)
	} else if this.object.M_exceptionType==BPMS_Runtime.ExceptionType_NoAvailableOutputSetException{
		ExceptionInfo = append(ExceptionInfo, 2)
	} else if this.object.M_exceptionType==BPMS_Runtime.ExceptionType_NotMatchingIOSpecification{
		ExceptionInfo = append(ExceptionInfo, 3)
	} else if this.object.M_exceptionType==BPMS_Runtime.ExceptionType_IllegalStartEventException{
		ExceptionInfo = append(ExceptionInfo, 4)
	}else{
		ExceptionInfo = append(ExceptionInfo, 0)
	}

	/** associations of Exception **/

	/** Save runtimes type Runtimes **/
		ExceptionInfo = append(ExceptionInfo,this.object.M_runtimesPtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), ExceptionInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), ExceptionInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_ExceptionEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_ExceptionEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Exception"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Exception...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.Exception)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.Exception"

		this.parentUuid = results[0][1].(string)

		/** members of Exception **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** exceptionType **/
 		if results[0][3] != nil{
 			enumIndex := results[0][3].(int)
			if enumIndex == 0{
 				this.object.M_exceptionType=BPMS_Runtime.ExceptionType_NoIORuleException
			} else if enumIndex == 1{
 				this.object.M_exceptionType=BPMS_Runtime.ExceptionType_GatewayException
			} else if enumIndex == 2{
 				this.object.M_exceptionType=BPMS_Runtime.ExceptionType_NoAvailableOutputSetException
			} else if enumIndex == 3{
 				this.object.M_exceptionType=BPMS_Runtime.ExceptionType_NotMatchingIOSpecification
			} else if enumIndex == 4{
 				this.object.M_exceptionType=BPMS_Runtime.ExceptionType_IllegalStartEventException
 			}
 		}

		/** associations of Exception **/

		/** runtimesPtr **/
 		if results[0][4] != nil{
			id :=results[0][4].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.Runtimes"
				id_:= refTypeName + "$$" + id
				this.object.M_runtimesPtr= id
				GetServer().GetEntityManager().appendReference("runtimesPtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeExceptionEntityFromObject(object *BPMS_Runtime.Exception) *BPMS_Runtime_ExceptionEntity {
	 return this.NewBPMS_RuntimeExceptionEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_ExceptionEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeExceptionExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Exception"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_ExceptionEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_ExceptionEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			LogInfo
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_LogInfoEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.LogInfo
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeLogInfoEntity(objectId string, object interface{}) *BPMS_Runtime_LogInfoEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeLogInfoExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.LogInfo).TYPENAME = "BPMS_Runtime.LogInfo"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.LogInfo).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_LogInfoEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.LogInfo%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_LogInfoEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.LogInfo)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.LogInfo)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.LogInfo"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.LogInfo","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_LogInfoEntity) GetTypeName()string{
	return "BPMS_Runtime.LogInfo"
}
func(this *BPMS_Runtime_LogInfoEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_LogInfoEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_LogInfoEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_LogInfoEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_LogInfoEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_LogInfoEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_LogInfoEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_LogInfoEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_LogInfoEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_LogInfoEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_LogInfoEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_LogInfoEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_LogInfoEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_LogInfoEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_LogInfoEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_LogInfoEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_LogInfoEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_LogInfoEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_LogInfoEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_LogInfoEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_LogInfoEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_LogInfoEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_LogInfoEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.LogInfo"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_LogInfoEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_LogInfoEntityPrototype() {

	var logInfoEntityProto EntityPrototype
	logInfoEntityProto.TypeName = "BPMS_Runtime.LogInfo"
	logInfoEntityProto.Ids = append(logInfoEntityProto.Ids,"uuid")
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"uuid")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"xs.string")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,0)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,false)
	logInfoEntityProto.Indexs = append(logInfoEntityProto.Indexs,"parentUuid")
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"parentUuid")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"xs.string")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,1)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,false)

	/** members of LogInfo **/
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,2)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"M_id")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"xs.string")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,3)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"M_date")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"xs.int64")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,4)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"M_actor")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"xs.interface{}:Ref")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,5)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"M_action")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"xs.string")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,6)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"M_object")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"BPMS_Runtime.Instance:Ref")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,7)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,true)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"M_description")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"xs.string")

	/** associations of LogInfo **/
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,8)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,false)
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"M_runtimesPtr")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"BPMS_Runtime.Runtimes:Ref")
	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"childsUuid")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"[]xs.string")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,9)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,false)

	logInfoEntityProto.Fields = append(logInfoEntityProto.Fields,"referenced")
	logInfoEntityProto.FieldsType = append(logInfoEntityProto.FieldsType,"[]EntityRef")
	logInfoEntityProto.FieldsOrder = append(logInfoEntityProto.FieldsOrder,10)
	logInfoEntityProto.FieldsVisibility = append(logInfoEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&logInfoEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_LogInfoEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.LogInfo"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.LogInfo"

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
	}else{
		LogInfoInfo = append(LogInfoInfo, "")
	}

	/** members of LogInfo **/
	LogInfoInfo = append(LogInfoInfo, this.object.M_id)
	LogInfoInfo = append(LogInfoInfo, this.object.M_date)
	LogInfoInfo = append(LogInfoInfo, this.object.M_actor)
	LogInfoInfo = append(LogInfoInfo, this.object.M_action)

	/** Save object type Instance **/
		LogInfoInfo = append(LogInfoInfo,this.object.M_object)
	LogInfoInfo = append(LogInfoInfo, this.object.M_description)

	/** associations of LogInfo **/

	/** Save runtimes type Runtimes **/
		LogInfoInfo = append(LogInfoInfo,this.object.M_runtimesPtr)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), LogInfoInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), LogInfoInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_LogInfoEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_LogInfoEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.LogInfo"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of LogInfo...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.LogInfo)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.LogInfo"

		this.parentUuid = results[0][1].(string)

		/** members of LogInfo **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** date **/
 		if results[0][3] != nil{
 			this.object.M_date=results[0][3].(int64)
 		}

		/** actor **/
 		if results[0][4] != nil{
			id :=results[0][4].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.interface{}"
				id_:= refTypeName + "$$" + id
				this.object.M_actor= id
				GetServer().GetEntityManager().appendReference("actor",this.object.UUID, id_)
			}
 		}

		/** action **/
 		if results[0][5] != nil{
 			this.object.M_action=results[0][5].(string)
 		}

		/** object **/
 		if results[0][6] != nil{
			id :=results[0][6].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.Instance"
				id_:= refTypeName + "$$" + id
				this.object.M_object= id
				GetServer().GetEntityManager().appendReference("object",this.object.UUID, id_)
			}
 		}

		/** description **/
 		if results[0][7] != nil{
 			this.object.M_description=results[0][7].(string)
 		}

		/** associations of LogInfo **/

		/** runtimesPtr **/
 		if results[0][8] != nil{
			id :=results[0][8].(string)
			if len(id) > 0 {
				refTypeName:="BPMS_Runtime.Runtimes"
				id_:= refTypeName + "$$" + id
				this.object.M_runtimesPtr= id
				GetServer().GetEntityManager().appendReference("runtimesPtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMS_RuntimeLogInfoEntityFromObject(object *BPMS_Runtime.LogInfo) *BPMS_Runtime_LogInfoEntity {
	 return this.NewBPMS_RuntimeLogInfoEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_LogInfoEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeLogInfoExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.LogInfo"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_LogInfoEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_LogInfoEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}

////////////////////////////////////////////////////////////////////////////////
//              			Runtimes
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMS_Runtime_RuntimesEntity struct{
	/** not the object id, except for the definition **/
	uuid string
	parentPtr 			Entity
	parentUuid 			string
	childsPtr  			[]Entity
	childsUuid  		[]string
	referencesUuid  	[]string
	referencesPtr  	    []Entity
	prototype      		*EntityPrototype
	referenced  		[]EntityRef
	object *BPMS_Runtime.Runtimes
}

/** Constructor function **/
func (this *EntityManager) NewBPMS_RuntimeRuntimesEntity(objectId string, object interface{}) *BPMS_Runtime_RuntimesEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMS_RuntimeRuntimesExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMS_Runtime.Runtimes).TYPENAME = "BPMS_Runtime.Runtimes"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMS_Runtime.Runtimes).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMS_Runtime_RuntimesEntity)
		}
	}else{
		uuidStr = "BPMS_Runtime.Runtimes%" + uuid.NewRandom().String()
	}
	entity := new(BPMS_Runtime_RuntimesEntity)
	if object == nil{
		entity.object = new(BPMS_Runtime.Runtimes)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMS_Runtime.Runtimes)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMS_Runtime.Runtimes"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMS_Runtime.Runtimes","BPMS_Runtime")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMS_Runtime_RuntimesEntity) GetTypeName()string{
	return "BPMS_Runtime.Runtimes"
}
func(this *BPMS_Runtime_RuntimesEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMS_Runtime_RuntimesEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMS_Runtime_RuntimesEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMS_Runtime_RuntimesEntity) AppendReferenced(name string, owner Entity){
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i:=0; i<len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { 
			return;
		}
	}
	this.referenced = append(this.referenced, ref)
}

func(this *BPMS_Runtime_RuntimesEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMS_Runtime_RuntimesEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef,0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func(this *BPMS_Runtime_RuntimesEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMS_Runtime_RuntimesEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMS_Runtime_RuntimesEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMS_Runtime_RuntimesEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMS_Runtime_RuntimesEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMS_Runtime_RuntimesEntity) RemoveChild(name string, uuid string) {
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
 		}else{
			params[0] = this.GetChildsPtr()[i].GetObject()
 		}
 	}
 	this.childsPtr = childsPtr

	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	Utility.CallMethod(this.GetObject(), removeMethode, params)
 }

func(this *BPMS_Runtime_RuntimesEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMS_Runtime_RuntimesEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMS_Runtime_RuntimesEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMS_Runtime_RuntimesEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMS_Runtime_RuntimesEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMS_Runtime_RuntimesEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMS_Runtime_RuntimesEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMS_Runtime_RuntimesEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMS_Runtime_RuntimesEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMS_Runtime_RuntimesEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMS_Runtime_RuntimesEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Runtimes"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMS_Runtime_RuntimesEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) Create_BPMS_Runtime_RuntimesEntityPrototype() {

	var runtimesEntityProto EntityPrototype
	runtimesEntityProto.TypeName = "BPMS_Runtime.Runtimes"
	runtimesEntityProto.Ids = append(runtimesEntityProto.Ids,"uuid")
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"uuid")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,0)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,false)
	runtimesEntityProto.Indexs = append(runtimesEntityProto.Indexs,"parentUuid")
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"parentUuid")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,1)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,false)

	/** members of Runtimes **/
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,2)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"M_id")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,3)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"M_name")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,4)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"M_version")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,5)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"M_definitions")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"[]BPMS_Runtime.DefinitionsInstance")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,6)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"M_exceptions")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"[]BPMS_Runtime.Exception")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,7)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"M_triggers")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"[]BPMS_Runtime.Trigger")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,8)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"M_correlationInfos")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"[]BPMS_Runtime.CorrelationInfo")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,9)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,true)
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"M_logInfos")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"[]BPMS_Runtime.LogInfo")
	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"childsUuid")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"[]xs.string")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,10)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,false)

	runtimesEntityProto.Fields = append(runtimesEntityProto.Fields,"referenced")
	runtimesEntityProto.FieldsType = append(runtimesEntityProto.FieldsType,"[]EntityRef")
	runtimesEntityProto.FieldsOrder = append(runtimesEntityProto.FieldsOrder,11)
	runtimesEntityProto.FieldsVisibility = append(runtimesEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMS_RuntimeDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&runtimesEntityProto)

}

/** Create **/
func (this *BPMS_Runtime_RuntimesEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMS_Runtime.Runtimes"

	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Runtimes"

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
	}else{
		RuntimesInfo = append(RuntimesInfo, "")
	}

	/** members of Runtimes **/
	RuntimesInfo = append(RuntimesInfo, this.object.M_id)
	RuntimesInfo = append(RuntimesInfo, this.object.M_name)
	RuntimesInfo = append(RuntimesInfo, this.object.M_version)

	/** Save definitions type DefinitionsInstance **/
	definitionsIds := make([]string,0)
	for i := 0; i < len(this.object.M_definitions); i++ {
		definitionsEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeDefinitionsInstanceEntity(this.object.M_definitions[i].UUID,this.object.M_definitions[i])
		definitionsIds=append(definitionsIds,definitionsEntity.uuid)
		definitionsEntity.AppendReferenced("definitions", this)
		this.AppendChild("definitions",definitionsEntity)
		if definitionsEntity.NeedSave() {
			definitionsEntity.SaveEntity()
		}
	}
	definitionsStr, _ := json.Marshal(definitionsIds)
	RuntimesInfo = append(RuntimesInfo, string(definitionsStr))

	/** Save exceptions type Exception **/
	exceptionsIds := make([]string,0)
	for i := 0; i < len(this.object.M_exceptions); i++ {
		exceptionsEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeExceptionEntity(this.object.M_exceptions[i].UUID,this.object.M_exceptions[i])
		exceptionsIds=append(exceptionsIds,exceptionsEntity.uuid)
		exceptionsEntity.AppendReferenced("exceptions", this)
		this.AppendChild("exceptions",exceptionsEntity)
		if exceptionsEntity.NeedSave() {
			exceptionsEntity.SaveEntity()
		}
	}
	exceptionsStr, _ := json.Marshal(exceptionsIds)
	RuntimesInfo = append(RuntimesInfo, string(exceptionsStr))

	/** Save triggers type Trigger **/
	triggersIds := make([]string,0)
	for i := 0; i < len(this.object.M_triggers); i++ {
		triggersEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeTriggerEntity(this.object.M_triggers[i].UUID,this.object.M_triggers[i])
		triggersIds=append(triggersIds,triggersEntity.uuid)
		triggersEntity.AppendReferenced("triggers", this)
		this.AppendChild("triggers",triggersEntity)
		if triggersEntity.NeedSave() {
			triggersEntity.SaveEntity()
		}
	}
	triggersStr, _ := json.Marshal(triggersIds)
	RuntimesInfo = append(RuntimesInfo, string(triggersStr))

	/** Save correlationInfos type CorrelationInfo **/
	correlationInfosIds := make([]string,0)
	for i := 0; i < len(this.object.M_correlationInfos); i++ {
		correlationInfosEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeCorrelationInfoEntity(this.object.M_correlationInfos[i].UUID,this.object.M_correlationInfos[i])
		correlationInfosIds=append(correlationInfosIds,correlationInfosEntity.uuid)
		correlationInfosEntity.AppendReferenced("correlationInfos", this)
		this.AppendChild("correlationInfos",correlationInfosEntity)
		if correlationInfosEntity.NeedSave() {
			correlationInfosEntity.SaveEntity()
		}
	}
	correlationInfosStr, _ := json.Marshal(correlationInfosIds)
	RuntimesInfo = append(RuntimesInfo, string(correlationInfosStr))

	/** Save logInfos type LogInfo **/
	logInfosIds := make([]string,0)
	for i := 0; i < len(this.object.M_logInfos); i++ {
		logInfosEntity:= GetServer().GetEntityManager().NewBPMS_RuntimeLogInfoEntity(this.object.M_logInfos[i].UUID,this.object.M_logInfos[i])
		logInfosIds=append(logInfosIds,logInfosEntity.uuid)
		logInfosEntity.AppendReferenced("logInfos", this)
		this.AppendChild("logInfos",logInfosEntity)
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
		err = GetServer().GetDataManager().updateData(BPMS_RuntimeDB, string(queryStr), RuntimesInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMS_RuntimeDB, string(queryStr), RuntimesInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMS_Runtime_RuntimesEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMS_Runtime_RuntimesEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Runtimes"

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

	results, err = GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Runtimes...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMS_Runtime.Runtimes)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMS_Runtime.Runtimes"

		this.parentUuid = results[0][1].(string)

		/** members of Runtimes **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** name **/
 		if results[0][3] != nil{
 			this.object.M_name=results[0][3].(string)
 		}

		/** version **/
 		if results[0][4] != nil{
 			this.object.M_version=results[0][4].(string)
 		}

		/** definitions **/
 		if results[0][5] != nil{
			uuidsStr :=results[0][5].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var definitionsEntity *BPMS_Runtime_DefinitionsInstanceEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						definitionsEntity = instance.(*BPMS_Runtime_DefinitionsInstanceEntity)
					}else{
						definitionsEntity = GetServer().GetEntityManager().NewBPMS_RuntimeDefinitionsInstanceEntity(uuids[i], nil)
						definitionsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(definitionsEntity)
					}
					definitionsEntity.AppendReferenced("definitions", this)
					this.AppendChild("definitions",definitionsEntity)
				}
 			}
 		}

		/** exceptions **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var exceptionsEntity *BPMS_Runtime_ExceptionEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						exceptionsEntity = instance.(*BPMS_Runtime_ExceptionEntity)
					}else{
						exceptionsEntity = GetServer().GetEntityManager().NewBPMS_RuntimeExceptionEntity(uuids[i], nil)
						exceptionsEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(exceptionsEntity)
					}
					exceptionsEntity.AppendReferenced("exceptions", this)
					this.AppendChild("exceptions",exceptionsEntity)
				}
 			}
 		}

		/** triggers **/
 		if results[0][7] != nil{
			uuidsStr :=results[0][7].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var triggersEntity *BPMS_Runtime_TriggerEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						triggersEntity = instance.(*BPMS_Runtime_TriggerEntity)
					}else{
						triggersEntity = GetServer().GetEntityManager().NewBPMS_RuntimeTriggerEntity(uuids[i], nil)
						triggersEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(triggersEntity)
					}
					triggersEntity.AppendReferenced("triggers", this)
					this.AppendChild("triggers",triggersEntity)
				}
 			}
 		}

		/** correlationInfos **/
 		if results[0][8] != nil{
			uuidsStr :=results[0][8].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var correlationInfosEntity *BPMS_Runtime_CorrelationInfoEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						correlationInfosEntity = instance.(*BPMS_Runtime_CorrelationInfoEntity)
					}else{
						correlationInfosEntity = GetServer().GetEntityManager().NewBPMS_RuntimeCorrelationInfoEntity(uuids[i], nil)
						correlationInfosEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(correlationInfosEntity)
					}
					correlationInfosEntity.AppendReferenced("correlationInfos", this)
					this.AppendChild("correlationInfos",correlationInfosEntity)
				}
 			}
 		}

		/** logInfos **/
 		if results[0][9] != nil{
			uuidsStr :=results[0][9].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var logInfosEntity *BPMS_Runtime_LogInfoEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						logInfosEntity = instance.(*BPMS_Runtime_LogInfoEntity)
					}else{
						logInfosEntity = GetServer().GetEntityManager().NewBPMS_RuntimeLogInfoEntity(uuids[i], nil)
						logInfosEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(logInfosEntity)
					}
					logInfosEntity.AppendReferenced("logInfos", this)
					this.AppendChild("logInfos",logInfosEntity)
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
func (this *EntityManager) NewBPMS_RuntimeRuntimesEntityFromObject(object *BPMS_Runtime.Runtimes) *BPMS_Runtime_RuntimesEntity {
	 return this.NewBPMS_RuntimeRuntimesEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMS_Runtime_RuntimesEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMS_RuntimeRuntimesExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMS_Runtime.Runtimes"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMS_RuntimeDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMS_Runtime_RuntimesEntity) AppendChild(attributeName string, child Entity) error {

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
	attributeName = strings.Replace(attributeName,"M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}
/** Append reference entity into parent entity. **/
func (this *BPMS_Runtime_RuntimesEntity) AppendReference(reference Entity) {

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
	 }else{
	 	// The reference must be update in that case.
	 	this.referencesPtr[index]  = reference
	 }
}
/** Register the entity to the dynamic typing system. **/
func (this *EntityManager) RegisterBPMS_RuntimeObjects(){
	Utility.RegisterType((*BPMS_Runtime.ConnectingObject)(nil))
	Utility.RegisterType((*BPMS_Runtime.ActivityInstance)(nil))
	Utility.RegisterType((*BPMS_Runtime.SubprocessInstance)(nil))
	Utility.RegisterType((*BPMS_Runtime.GatewayInstance)(nil))
	Utility.RegisterType((*BPMS_Runtime.EventInstance)(nil))
	Utility.RegisterType((*BPMS_Runtime.DefinitionsInstance)(nil))
	Utility.RegisterType((*BPMS_Runtime.ProcessInstance)(nil))
	Utility.RegisterType((*BPMS_Runtime.EventDefinitionInstance)(nil))
	Utility.RegisterType((*BPMS_Runtime.CorrelationInfo)(nil))
	Utility.RegisterType((*BPMS_Runtime.ItemAwareElementInstance)(nil))
	Utility.RegisterType((*BPMS_Runtime.EventData)(nil))
	Utility.RegisterType((*BPMS_Runtime.Trigger)(nil))
	Utility.RegisterType((*BPMS_Runtime.Exception)(nil))
	Utility.RegisterType((*BPMS_Runtime.LogInfo)(nil))
	Utility.RegisterType((*BPMS_Runtime.Runtimes)(nil))
}

/** Create entity prototypes contain in a package **/
func (this *EntityManager) CreateBPMS_RuntimePrototypes(){
	this.Create_BPMS_Runtime_InstanceEntityPrototype() 
	this.Create_BPMS_Runtime_ConnectingObjectEntityPrototype() 
	this.Create_BPMS_Runtime_FlowNodeInstanceEntityPrototype() 
	this.Create_BPMS_Runtime_ActivityInstanceEntityPrototype() 
	this.Create_BPMS_Runtime_SubprocessInstanceEntityPrototype() 
	this.Create_BPMS_Runtime_GatewayInstanceEntityPrototype() 
	this.Create_BPMS_Runtime_EventInstanceEntityPrototype() 
	this.Create_BPMS_Runtime_DefinitionsInstanceEntityPrototype() 
	this.Create_BPMS_Runtime_ProcessInstanceEntityPrototype() 
	this.Create_BPMS_Runtime_EventDefinitionInstanceEntityPrototype() 
	this.Create_BPMS_Runtime_CorrelationInfoEntityPrototype() 
	this.Create_BPMS_Runtime_ItemAwareElementInstanceEntityPrototype() 
	this.Create_BPMS_Runtime_EventDataEntityPrototype() 
	this.Create_BPMS_Runtime_TriggerEntityPrototype() 
	this.Create_BPMS_Runtime_ExceptionEntityPrototype() 
	this.Create_BPMS_Runtime_LogInfoEntityPrototype() 
	this.Create_BPMS_Runtime_RuntimesEntityPrototype() 
}

