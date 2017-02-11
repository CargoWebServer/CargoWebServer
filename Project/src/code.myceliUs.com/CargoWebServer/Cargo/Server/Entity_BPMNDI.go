package Server
import(
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/BPMNDI"
	"encoding/json"
	"code.myceliUs.com/Utility"
	"log"
	"strings"
)


////////////////////////////////////////////////////////////////////////////////
//              			BPMNDiagram
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMNDI_BPMNDiagramEntity struct{
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
	object *BPMNDI.BPMNDiagram
}

/** Constructor function **/
func (this *EntityManager) NewBPMNDIBPMNDiagramEntity(objectId string, object interface{}) *BPMNDI_BPMNDiagramEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMNDIBPMNDiagramExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMNDI.BPMNDiagram).TYPENAME = "BPMNDI.BPMNDiagram"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMNDI.BPMNDiagram).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMNDI_BPMNDiagramEntity)
		}
	}else{
		uuidStr = "BPMNDI.BPMNDiagram%" + Utility.RandomUUID()
	}
	entity := new(BPMNDI_BPMNDiagramEntity)
	if object == nil{
		entity.object = new(BPMNDI.BPMNDiagram)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMNDI.BPMNDiagram)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMNDI.BPMNDiagram"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMNDI.BPMNDiagram","BPMN20")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMNDI_BPMNDiagramEntity) GetTypeName()string{
	return "BPMNDI.BPMNDiagram"
}
func(this *BPMNDI_BPMNDiagramEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMNDI_BPMNDiagramEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMNDI_BPMNDiagramEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMNDI_BPMNDiagramEntity) AppendReferenced(name string, owner Entity){
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

func(this *BPMNDI_BPMNDiagramEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMNDI_BPMNDiagramEntity) RemoveReferenced(name string, owner Entity) {
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

func(this *BPMNDI_BPMNDiagramEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMNDI_BPMNDiagramEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMNDI_BPMNDiagramEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMNDI_BPMNDiagramEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMNDI_BPMNDiagramEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMNDI_BPMNDiagramEntity) RemoveChild(name string, uuid string) {
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

func(this *BPMNDI_BPMNDiagramEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMNDI_BPMNDiagramEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMNDI_BPMNDiagramEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMNDI_BPMNDiagramEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMNDI_BPMNDiagramEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMNDI_BPMNDiagramEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMNDI_BPMNDiagramEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMNDI_BPMNDiagramEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMNDI_BPMNDiagramEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMNDI_BPMNDiagramEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMNDI_BPMNDiagramEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNDiagram"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMNDI_BPMNDiagramEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) create_BPMNDI_BPMNDiagramEntityPrototype() {

	var bPMNDiagramEntityProto EntityPrototype
	bPMNDiagramEntityProto.TypeName = "BPMNDI.BPMNDiagram"
	bPMNDiagramEntityProto.SuperTypeNames = append(bPMNDiagramEntityProto.SuperTypeNames, "DI.Diagram")
	bPMNDiagramEntityProto.Ids = append(bPMNDiagramEntityProto.Ids,"uuid")
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"uuid")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"xs.string")
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,0)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,false)
	bPMNDiagramEntityProto.Indexs = append(bPMNDiagramEntityProto.Indexs,"parentUuid")
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"parentUuid")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"xs.string")
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,1)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,false)

	/** members of Diagram **/
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,2)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,true)
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"M_rootElement")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"BPMNDI.DI.DiagramElement")
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,3)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,true)
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"M_name")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"xs.string")
	bPMNDiagramEntityProto.Ids = append(bPMNDiagramEntityProto.Ids,"M_id")
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,4)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,true)
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"M_id")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"xs.ID")
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,5)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,true)
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"M_documentation")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"xs.string")
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,6)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,true)
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"M_resolution")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"xs.float64")
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,7)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,true)
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"M_ownedStyle")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"[]BPMNDI.DI.Style")

	/** members of BPMNDiagram **/
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,8)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,true)
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"M_BPMNPlane")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"BPMNDI.BPMNPlane")
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,9)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,true)
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"M_BPMNLabelStyle")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"[]BPMNDI.BPMNLabelStyle")
	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"childsUuid")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"[]xs.string")
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,10)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,false)

	bPMNDiagramEntityProto.Fields = append(bPMNDiagramEntityProto.Fields,"referenced")
	bPMNDiagramEntityProto.FieldsType = append(bPMNDiagramEntityProto.FieldsType,"[]EntityRef")
	bPMNDiagramEntityProto.FieldsOrder = append(bPMNDiagramEntityProto.FieldsOrder,11)
	bPMNDiagramEntityProto.FieldsVisibility = append(bPMNDiagramEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMNDIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&bPMNDiagramEntityProto)

}

/** Create **/
func (this *BPMNDI_BPMNDiagramEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMNDI.BPMNDiagram"

	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNDiagram"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Diagram **/
	query.Fields = append(query.Fields, "M_rootElement")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_documentation")
	query.Fields = append(query.Fields, "M_resolution")
	query.Fields = append(query.Fields, "M_ownedStyle")

	/** members of BPMNDiagram **/
	query.Fields = append(query.Fields, "M_BPMNPlane")
	query.Fields = append(query.Fields, "M_BPMNLabelStyle")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var BPMNDiagramInfo []interface{}

	BPMNDiagramInfo = append(BPMNDiagramInfo, this.GetUuid())
	if this.parentPtr != nil {
		BPMNDiagramInfo = append(BPMNDiagramInfo, this.parentPtr.GetUuid())
	}else{
		BPMNDiagramInfo = append(BPMNDiagramInfo, "")
	}

	/** members of Diagram **/

	/** Save rootElement type DiagramElement **/
	if this.object.M_rootElement != nil {
		switch v := this.object.M_rootElement.(type) {
		case *BPMNDI.BPMNShape:
			rootElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(v.UUID, v)
			BPMNDiagramInfo = append(BPMNDiagramInfo, rootElementEntity.uuid)
		    rootElementEntity.AppendReferenced("rootElement", this)
			this.AppendChild("rootElement",rootElementEntity)
			if rootElementEntity.NeedSave() {
				rootElementEntity.SaveEntity()
			}
		case *BPMNDI.BPMNPlane:
			rootElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(v.UUID, v)
			BPMNDiagramInfo = append(BPMNDiagramInfo, rootElementEntity.uuid)
		    rootElementEntity.AppendReferenced("rootElement", this)
			this.AppendChild("rootElement",rootElementEntity)
			if rootElementEntity.NeedSave() {
				rootElementEntity.SaveEntity()
			}
		case *BPMNDI.BPMNLabel:
			rootElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
			BPMNDiagramInfo = append(BPMNDiagramInfo, rootElementEntity.uuid)
		    rootElementEntity.AppendReferenced("rootElement", this)
			this.AppendChild("rootElement",rootElementEntity)
			if rootElementEntity.NeedSave() {
				rootElementEntity.SaveEntity()
			}
		case *BPMNDI.BPMNEdge:
			rootElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(v.UUID, v)
			BPMNDiagramInfo = append(BPMNDiagramInfo, rootElementEntity.uuid)
		    rootElementEntity.AppendReferenced("rootElement", this)
			this.AppendChild("rootElement",rootElementEntity)
			if rootElementEntity.NeedSave() {
				rootElementEntity.SaveEntity()
			}
			}
	}else{
		BPMNDiagramInfo = append(BPMNDiagramInfo, "")
	}
	BPMNDiagramInfo = append(BPMNDiagramInfo, this.object.M_name)
	BPMNDiagramInfo = append(BPMNDiagramInfo, this.object.M_id)
	BPMNDiagramInfo = append(BPMNDiagramInfo, this.object.M_documentation)
	BPMNDiagramInfo = append(BPMNDiagramInfo, this.object.M_resolution)

	/** Save ownedStyle type Style **/
	ownedStyleIds := make([]string,0)
	for i := 0; i < len(this.object.M_ownedStyle); i++ {
		switch v := this.object.M_ownedStyle[i].(type) {
		case *BPMNDI.BPMNLabelStyle:
		ownedStyleEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelStyleEntity(v.UUID, v)
		ownedStyleIds=append(ownedStyleIds,ownedStyleEntity.uuid)
		ownedStyleEntity.AppendReferenced("ownedStyle", this)
		this.AppendChild("ownedStyle",ownedStyleEntity)
		if ownedStyleEntity.NeedSave() {
			ownedStyleEntity.SaveEntity()
		}
		}
	}
	ownedStyleStr, _ := json.Marshal(ownedStyleIds)
	BPMNDiagramInfo = append(BPMNDiagramInfo, string(ownedStyleStr))

	/** members of BPMNDiagram **/

	/** Save BPMNPlane type BPMNPlane **/
	if this.object.M_BPMNPlane != nil {
		BPMNPlaneEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(this.object.M_BPMNPlane.UUID, this.object.M_BPMNPlane)
		BPMNDiagramInfo = append(BPMNDiagramInfo, BPMNPlaneEntity.uuid)
		BPMNPlaneEntity.AppendReferenced("BPMNPlane", this)
		this.AppendChild("BPMNPlane",BPMNPlaneEntity)
		if BPMNPlaneEntity.NeedSave() {
			BPMNPlaneEntity.SaveEntity()
		}
	}else{
		BPMNDiagramInfo = append(BPMNDiagramInfo, "")
	}

	/** Save BPMNLabelStyle type BPMNLabelStyle **/
	BPMNLabelStyleIds := make([]string,0)
	for i := 0; i < len(this.object.M_BPMNLabelStyle); i++ {
		BPMNLabelStyleEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelStyleEntity(this.object.M_BPMNLabelStyle[i].UUID,this.object.M_BPMNLabelStyle[i])
		BPMNLabelStyleIds=append(BPMNLabelStyleIds,BPMNLabelStyleEntity.uuid)
		BPMNLabelStyleEntity.AppendReferenced("BPMNLabelStyle", this)
		this.AppendChild("BPMNLabelStyle",BPMNLabelStyleEntity)
		if BPMNLabelStyleEntity.NeedSave() {
			BPMNLabelStyleEntity.SaveEntity()
		}
	}
	BPMNLabelStyleStr, _ := json.Marshal(BPMNLabelStyleIds)
	BPMNDiagramInfo = append(BPMNDiagramInfo, string(BPMNLabelStyleStr))
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	BPMNDiagramInfo = append(BPMNDiagramInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	BPMNDiagramInfo = append(BPMNDiagramInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMNDIDB, string(queryStr), BPMNDiagramInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMNDIDB, string(queryStr), BPMNDiagramInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMNDI_BPMNDiagramEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMNDI_BPMNDiagramEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNDiagram"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Diagram **/
	query.Fields = append(query.Fields, "M_rootElement")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_documentation")
	query.Fields = append(query.Fields, "M_resolution")
	query.Fields = append(query.Fields, "M_ownedStyle")

	/** members of BPMNDiagram **/
	query.Fields = append(query.Fields, "M_BPMNPlane")
	query.Fields = append(query.Fields, "M_BPMNLabelStyle")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of BPMNDiagram...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMNDI.BPMNDiagram)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMNDI.BPMNDiagram"

		this.parentUuid = results[0][1].(string)

		/** members of Diagram **/

		/** rootElement **/
 		if results[0][2] != nil{
			uuid :=results[0][2].(string)
			if len(uuid) > 0 {
				typeName := uuid[0:strings.Index(uuid, "%")]
				if err!=nil{
					log.Println("type ", typeName, " not found!")
					return err
				}
			if typeName == "BPMNDI.BPMNPlane"{
				if len(uuid) > 0 {
					var rootElementEntity *BPMNDI_BPMNPlaneEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						rootElementEntity = instance.(*BPMNDI_BPMNPlaneEntity)
					}else{
						rootElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(uuid, nil)
						rootElementEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(rootElementEntity)
					}
					rootElementEntity.AppendReferenced("rootElement", this)
					this.AppendChild("rootElement",rootElementEntity)
				}
			} else if typeName == "BPMNDI.BPMNLabel"{
				if len(uuid) > 0 {
					var rootElementEntity *BPMNDI_BPMNLabelEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						rootElementEntity = instance.(*BPMNDI_BPMNLabelEntity)
					}else{
						rootElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuid, nil)
						rootElementEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(rootElementEntity)
					}
					rootElementEntity.AppendReferenced("rootElement", this)
					this.AppendChild("rootElement",rootElementEntity)
				}
			} else if typeName == "BPMNDI.BPMNShape"{
				if len(uuid) > 0 {
					var rootElementEntity *BPMNDI_BPMNShapeEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						rootElementEntity = instance.(*BPMNDI_BPMNShapeEntity)
					}else{
						rootElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(uuid, nil)
						rootElementEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(rootElementEntity)
					}
					rootElementEntity.AppendReferenced("rootElement", this)
					this.AppendChild("rootElement",rootElementEntity)
				}
			} else if typeName == "BPMNDI.BPMNEdge"{
				if len(uuid) > 0 {
					var rootElementEntity *BPMNDI_BPMNEdgeEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						rootElementEntity = instance.(*BPMNDI_BPMNEdgeEntity)
					}else{
						rootElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(uuid, nil)
						rootElementEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(rootElementEntity)
					}
					rootElementEntity.AppendReferenced("rootElement", this)
					this.AppendChild("rootElement",rootElementEntity)
				}
			}
			}
 		}

		/** name **/
 		if results[0][3] != nil{
 			this.object.M_name=results[0][3].(string)
 		}

		/** id **/
 		if results[0][4] != nil{
 			this.object.M_id=results[0][4].(string)
 		}

		/** documentation **/
 		if results[0][5] != nil{
 			this.object.M_documentation=results[0][5].(string)
 		}

		/** resolution **/
 		if results[0][6] != nil{
 			this.object.M_resolution=results[0][6].(float64)
 		}

		/** ownedStyle **/
 		if results[0][7] != nil{
			uuidsStr :=results[0][7].(string)
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
				if typeName == "BPMNDI.BPMNLabelStyle"{
						if len(uuids[i]) > 0 {
							var ownedStyleEntity *BPMNDI_BPMNLabelStyleEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedStyleEntity = instance.(*BPMNDI_BPMNLabelStyleEntity)
							}else{
								ownedStyleEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelStyleEntity(uuids[i], nil)
								ownedStyleEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedStyleEntity)
							}
							ownedStyleEntity.AppendReferenced("ownedStyle", this)
							this.AppendChild("ownedStyle",ownedStyleEntity)
						}
				}
 			}
 		}

		/** members of BPMNDiagram **/

		/** BPMNPlane **/
 		if results[0][8] != nil{
			uuid :=results[0][8].(string)
			if len(uuid) > 0 {
				var BPMNPlaneEntity *BPMNDI_BPMNPlaneEntity
				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
					BPMNPlaneEntity = instance.(*BPMNDI_BPMNPlaneEntity)
				}else{
					BPMNPlaneEntity = GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(uuid, nil)
					BPMNPlaneEntity.InitEntity(uuid)
					GetServer().GetEntityManager().insert( BPMNPlaneEntity)
				}
				BPMNPlaneEntity.AppendReferenced("BPMNPlane", this)
				this.AppendChild("BPMNPlane",BPMNPlaneEntity)
			}
 		}

		/** BPMNLabelStyle **/
 		if results[0][9] != nil{
			uuidsStr :=results[0][9].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var BPMNLabelStyleEntity *BPMNDI_BPMNLabelStyleEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						BPMNLabelStyleEntity = instance.(*BPMNDI_BPMNLabelStyleEntity)
					}else{
						BPMNLabelStyleEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelStyleEntity(uuids[i], nil)
						BPMNLabelStyleEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(BPMNLabelStyleEntity)
					}
					BPMNLabelStyleEntity.AppendReferenced("BPMNLabelStyle", this)
					this.AppendChild("BPMNLabelStyle",BPMNLabelStyleEntity)
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
func (this *EntityManager) NewBPMNDIBPMNDiagramEntityFromObject(object *BPMNDI.BPMNDiagram) *BPMNDI_BPMNDiagramEntity {
	 return this.NewBPMNDIBPMNDiagramEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMNDI_BPMNDiagramEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMNDIBPMNDiagramExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNDiagram"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMNDI_BPMNDiagramEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMNDI_BPMNDiagramEntity) AppendReference(reference Entity) {

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
//              			BPMNPlane
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMNDI_BPMNPlaneEntity struct{
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
	object *BPMNDI.BPMNPlane
}

/** Constructor function **/
func (this *EntityManager) NewBPMNDIBPMNPlaneEntity(objectId string, object interface{}) *BPMNDI_BPMNPlaneEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMNDIBPMNPlaneExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMNDI.BPMNPlane).TYPENAME = "BPMNDI.BPMNPlane"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMNDI.BPMNPlane).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMNDI_BPMNPlaneEntity)
		}
	}else{
		uuidStr = "BPMNDI.BPMNPlane%" + Utility.RandomUUID()
	}
	entity := new(BPMNDI_BPMNPlaneEntity)
	if object == nil{
		entity.object = new(BPMNDI.BPMNPlane)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMNDI.BPMNPlane)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMNDI.BPMNPlane"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMNDI.BPMNPlane","BPMN20")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMNDI_BPMNPlaneEntity) GetTypeName()string{
	return "BPMNDI.BPMNPlane"
}
func(this *BPMNDI_BPMNPlaneEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMNDI_BPMNPlaneEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMNDI_BPMNPlaneEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMNDI_BPMNPlaneEntity) AppendReferenced(name string, owner Entity){
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

func(this *BPMNDI_BPMNPlaneEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMNDI_BPMNPlaneEntity) RemoveReferenced(name string, owner Entity) {
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

func(this *BPMNDI_BPMNPlaneEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMNDI_BPMNPlaneEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMNDI_BPMNPlaneEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMNDI_BPMNPlaneEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMNDI_BPMNPlaneEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMNDI_BPMNPlaneEntity) RemoveChild(name string, uuid string) {
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

func(this *BPMNDI_BPMNPlaneEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMNDI_BPMNPlaneEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMNDI_BPMNPlaneEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMNDI_BPMNPlaneEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMNDI_BPMNPlaneEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMNDI_BPMNPlaneEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMNDI_BPMNPlaneEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMNDI_BPMNPlaneEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMNDI_BPMNPlaneEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMNDI_BPMNPlaneEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMNDI_BPMNPlaneEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNPlane"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMNDI_BPMNPlaneEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) create_BPMNDI_BPMNPlaneEntityPrototype() {

	var bPMNPlaneEntityProto EntityPrototype
	bPMNPlaneEntityProto.TypeName = "BPMNDI.BPMNPlane"
	bPMNPlaneEntityProto.SuperTypeNames = append(bPMNPlaneEntityProto.SuperTypeNames, "DI.DiagramElement")
	bPMNPlaneEntityProto.SuperTypeNames = append(bPMNPlaneEntityProto.SuperTypeNames, "DI.Node")
	bPMNPlaneEntityProto.SuperTypeNames = append(bPMNPlaneEntityProto.SuperTypeNames, "DI.Plane")
	bPMNPlaneEntityProto.Ids = append(bPMNPlaneEntityProto.Ids,"uuid")
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"uuid")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"xs.string")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,0)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,false)
	bPMNPlaneEntityProto.Indexs = append(bPMNPlaneEntityProto.Indexs,"parentUuid")
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"parentUuid")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"xs.string")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,1)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,false)

	/** members of DiagramElement **/
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,2)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,true)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_owningDiagram")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"BPMNDI.DI.Diagram:Ref")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,3)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,true)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_owningElement")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"BPMNDI.DI.DiagramElement:Ref")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,4)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,true)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_modelElement")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"xs.interface{}:Ref")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,5)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,true)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_style")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"BPMNDI.DI.Style:Ref")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,6)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,true)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_ownedElement")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"[]BPMNDI.DI.DiagramElement")
	bPMNPlaneEntityProto.Ids = append(bPMNPlaneEntityProto.Ids,"M_id")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,7)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,true)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_id")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"xs.ID")

	/** members of Node **/
	/** No members **/

	/** members of Plane **/
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,8)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,true)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_DiagramElement")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"[]BPMNDI.DI.DiagramElement")

	/** members of BPMNPlane **/
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,9)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,true)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_bpmnElement")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"xs.interface{}:Ref")

	/** associations of BPMNPlane **/
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,10)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,false)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_diagramPtr")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"BPMNDI.BPMNDiagram:Ref")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,11)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,false)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_sourceEdgePtr")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"[]BPMNDI.DI.Edge:Ref")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,12)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,false)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_targetEdgePtr")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"[]BPMNDI.DI.Edge:Ref")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,13)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,false)
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"M_planePtr")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"BPMNDI.DI.Plane:Ref")
	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"childsUuid")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"[]xs.string")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,14)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,false)

	bPMNPlaneEntityProto.Fields = append(bPMNPlaneEntityProto.Fields,"referenced")
	bPMNPlaneEntityProto.FieldsType = append(bPMNPlaneEntityProto.FieldsType,"[]EntityRef")
	bPMNPlaneEntityProto.FieldsOrder = append(bPMNPlaneEntityProto.FieldsOrder,15)
	bPMNPlaneEntityProto.FieldsVisibility = append(bPMNPlaneEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMNDIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&bPMNPlaneEntityProto)

}

/** Create **/
func (this *BPMNDI_BPMNPlaneEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMNDI.BPMNPlane"

	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNPlane"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of DiagramElement **/
	query.Fields = append(query.Fields, "M_owningDiagram")
	query.Fields = append(query.Fields, "M_owningElement")
	query.Fields = append(query.Fields, "M_modelElement")
	query.Fields = append(query.Fields, "M_style")
	query.Fields = append(query.Fields, "M_ownedElement")
	query.Fields = append(query.Fields, "M_id")

	/** members of Node **/
	/** No members **/

	/** members of Plane **/
	query.Fields = append(query.Fields, "M_DiagramElement")

	/** members of BPMNPlane **/
	query.Fields = append(query.Fields, "M_bpmnElement")

		/** associations of BPMNPlane **/
	query.Fields = append(query.Fields, "M_diagramPtr")
	query.Fields = append(query.Fields, "M_sourceEdgePtr")
	query.Fields = append(query.Fields, "M_targetEdgePtr")
	query.Fields = append(query.Fields, "M_planePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var BPMNPlaneInfo []interface{}

	BPMNPlaneInfo = append(BPMNPlaneInfo, this.GetUuid())
	if this.parentPtr != nil {
		BPMNPlaneInfo = append(BPMNPlaneInfo, this.parentPtr.GetUuid())
	}else{
		BPMNPlaneInfo = append(BPMNPlaneInfo, "")
	}

	/** members of DiagramElement **/

	/** Save owningDiagram type Diagram **/
		BPMNPlaneInfo = append(BPMNPlaneInfo,this.object.M_owningDiagram)

	/** Save owningElement type DiagramElement **/
		BPMNPlaneInfo = append(BPMNPlaneInfo,this.object.M_owningElement)
	BPMNPlaneInfo = append(BPMNPlaneInfo, this.object.M_modelElement)

	/** Save style type Style **/
		BPMNPlaneInfo = append(BPMNPlaneInfo,this.object.M_style)

	/** Save ownedElement type DiagramElement **/
	ownedElementIds := make([]string,0)
	for i := 0; i < len(this.object.M_ownedElement); i++ {
		switch v := this.object.M_ownedElement[i].(type) {
		case *BPMNDI.BPMNLabel:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNShape:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNPlane:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNEdge:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		}
	}
	ownedElementStr, _ := json.Marshal(ownedElementIds)
	BPMNPlaneInfo = append(BPMNPlaneInfo, string(ownedElementStr))
	BPMNPlaneInfo = append(BPMNPlaneInfo, this.object.M_id)

	/** members of Node **/
	/** No members **/

	/** members of Plane **/

	/** Save DiagramElement type DiagramElement **/
	DiagramElementIds := make([]string,0)
	for i := 0; i < len(this.object.M_DiagramElement); i++ {
		switch v := this.object.M_DiagramElement[i].(type) {
		case *BPMNDI.BPMNShape:
		DiagramElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(v.UUID, v)
		DiagramElementIds=append(DiagramElementIds,DiagramElementEntity.uuid)
		DiagramElementEntity.AppendReferenced("DiagramElement", this)
		this.AppendChild("DiagramElement",DiagramElementEntity)
		if DiagramElementEntity.NeedSave() {
			DiagramElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNPlane:
		DiagramElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(v.UUID, v)
		DiagramElementIds=append(DiagramElementIds,DiagramElementEntity.uuid)
		DiagramElementEntity.AppendReferenced("DiagramElement", this)
		this.AppendChild("DiagramElement",DiagramElementEntity)
		if DiagramElementEntity.NeedSave() {
			DiagramElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNLabel:
		DiagramElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
		DiagramElementIds=append(DiagramElementIds,DiagramElementEntity.uuid)
		DiagramElementEntity.AppendReferenced("DiagramElement", this)
		this.AppendChild("DiagramElement",DiagramElementEntity)
		if DiagramElementEntity.NeedSave() {
			DiagramElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNEdge:
		DiagramElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(v.UUID, v)
		DiagramElementIds=append(DiagramElementIds,DiagramElementEntity.uuid)
		DiagramElementEntity.AppendReferenced("DiagramElement", this)
		this.AppendChild("DiagramElement",DiagramElementEntity)
		if DiagramElementEntity.NeedSave() {
			DiagramElementEntity.SaveEntity()
		}
		}
	}
	DiagramElementStr, _ := json.Marshal(DiagramElementIds)
	BPMNPlaneInfo = append(BPMNPlaneInfo, string(DiagramElementStr))

	/** members of BPMNPlane **/
	BPMNPlaneInfo = append(BPMNPlaneInfo, this.object.M_bpmnElement)

	/** associations of BPMNPlane **/

	/** Save diagram type BPMNDiagram **/
		BPMNPlaneInfo = append(BPMNPlaneInfo,this.object.M_diagramPtr)

	/** Save sourceEdge type Edge **/
	sourceEdgePtrStr, _ := json.Marshal(this.object.M_sourceEdgePtr)
	BPMNPlaneInfo = append(BPMNPlaneInfo, string(sourceEdgePtrStr))

	/** Save targetEdge type Edge **/
	targetEdgePtrStr, _ := json.Marshal(this.object.M_targetEdgePtr)
	BPMNPlaneInfo = append(BPMNPlaneInfo, string(targetEdgePtrStr))

	/** Save plane type Plane **/
		BPMNPlaneInfo = append(BPMNPlaneInfo,this.object.M_planePtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	BPMNPlaneInfo = append(BPMNPlaneInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	BPMNPlaneInfo = append(BPMNPlaneInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMNDIDB, string(queryStr), BPMNPlaneInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMNDIDB, string(queryStr), BPMNPlaneInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMNDI_BPMNPlaneEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMNDI_BPMNPlaneEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNPlane"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of DiagramElement **/
	query.Fields = append(query.Fields, "M_owningDiagram")
	query.Fields = append(query.Fields, "M_owningElement")
	query.Fields = append(query.Fields, "M_modelElement")
	query.Fields = append(query.Fields, "M_style")
	query.Fields = append(query.Fields, "M_ownedElement")
	query.Fields = append(query.Fields, "M_id")

	/** members of Node **/
	/** No members **/

	/** members of Plane **/
	query.Fields = append(query.Fields, "M_DiagramElement")

	/** members of BPMNPlane **/
	query.Fields = append(query.Fields, "M_bpmnElement")

		/** associations of BPMNPlane **/
	query.Fields = append(query.Fields, "M_diagramPtr")
	query.Fields = append(query.Fields, "M_sourceEdgePtr")
	query.Fields = append(query.Fields, "M_targetEdgePtr")
	query.Fields = append(query.Fields, "M_planePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of BPMNPlane...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMNDI.BPMNPlane)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMNDI.BPMNPlane"

		this.parentUuid = results[0][1].(string)

		/** members of DiagramElement **/

		/** owningDiagram **/
 		if results[0][2] != nil{
			id :=results[0][2].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Diagram"
				id_:= refTypeName + "$$" + id
				this.object.M_owningDiagram= id
				GetServer().GetEntityManager().appendReference("owningDiagram",this.object.UUID, id_)
			}
 		}

		/** owningElement **/
 		if results[0][3] != nil{
			id :=results[0][3].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.DiagramElement"
				id_:= refTypeName + "$$" + id
				this.object.M_owningElement= id
				GetServer().GetEntityManager().appendReference("owningElement",this.object.UUID, id_)
			}
 		}

		/** modelElement **/
 		if results[0][4] != nil{
			id :=results[0][4].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.interface{}"
				id_:= refTypeName + "$$" + id
				this.object.M_modelElement= id
				GetServer().GetEntityManager().appendReference("modelElement",this.object.UUID, id_)
			}
 		}

		/** style **/
 		if results[0][5] != nil{
			id :=results[0][5].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Style"
				id_:= refTypeName + "$$" + id
				this.object.M_style= id
				GetServer().GetEntityManager().appendReference("style",this.object.UUID, id_)
			}
 		}

		/** ownedElement **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
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
				if typeName == "BPMNDI.BPMNEdge"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNEdgeEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNEdgeEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNShape"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNShapeEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNShapeEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNPlane"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNPlaneEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNPlaneEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNLabel"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNLabelEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNLabelEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				}
 			}
 		}

		/** id **/
 		if results[0][7] != nil{
 			this.object.M_id=results[0][7].(string)
 		}

		/** members of Node **/
		/** No members **/

		/** members of Plane **/

		/** DiagramElement **/
 		if results[0][8] != nil{
			uuidsStr :=results[0][8].(string)
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
				if typeName == "BPMNDI.BPMNEdge"{
						if len(uuids[i]) > 0 {
							var DiagramElementEntity *BPMNDI_BPMNEdgeEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								DiagramElementEntity = instance.(*BPMNDI_BPMNEdgeEntity)
							}else{
								DiagramElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(uuids[i], nil)
								DiagramElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(DiagramElementEntity)
							}
							DiagramElementEntity.AppendReferenced("DiagramElement", this)
							this.AppendChild("DiagramElement",DiagramElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNShape"{
						if len(uuids[i]) > 0 {
							var DiagramElementEntity *BPMNDI_BPMNShapeEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								DiagramElementEntity = instance.(*BPMNDI_BPMNShapeEntity)
							}else{
								DiagramElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(uuids[i], nil)
								DiagramElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(DiagramElementEntity)
							}
							DiagramElementEntity.AppendReferenced("DiagramElement", this)
							this.AppendChild("DiagramElement",DiagramElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNPlane"{
						if len(uuids[i]) > 0 {
							var DiagramElementEntity *BPMNDI_BPMNPlaneEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								DiagramElementEntity = instance.(*BPMNDI_BPMNPlaneEntity)
							}else{
								DiagramElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(uuids[i], nil)
								DiagramElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(DiagramElementEntity)
							}
							DiagramElementEntity.AppendReferenced("DiagramElement", this)
							this.AppendChild("DiagramElement",DiagramElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNLabel"{
						if len(uuids[i]) > 0 {
							var DiagramElementEntity *BPMNDI_BPMNLabelEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								DiagramElementEntity = instance.(*BPMNDI_BPMNLabelEntity)
							}else{
								DiagramElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuids[i], nil)
								DiagramElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(DiagramElementEntity)
							}
							DiagramElementEntity.AppendReferenced("DiagramElement", this)
							this.AppendChild("DiagramElement",DiagramElementEntity)
						}
				}
 			}
 		}

		/** members of BPMNPlane **/

		/** bpmnElement **/
 		if results[0][9] != nil{
			id :=results[0][9].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.interface{}"
				id_:= refTypeName + "$$" + id
				this.object.M_bpmnElement= id
				GetServer().GetEntityManager().appendReference("bpmnElement",this.object.UUID, id_)
				this.object.M_bpmnElement = id
			}
 		}

		/** associations of BPMNPlane **/

		/** diagramPtr **/
 		if results[0][10] != nil{
			id :=results[0][10].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.BPMNDiagram"
				id_:= refTypeName + "$$" + id
				this.object.M_diagramPtr= id
				GetServer().GetEntityManager().appendReference("diagramPtr",this.object.UUID, id_)
			}
 		}

		/** sourceEdgePtr **/
 		if results[0][11] != nil{
			idsStr :=results[0][11].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMNDI.Edge"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_sourceEdgePtr = append(this.object.M_sourceEdgePtr,ids[i])
					GetServer().GetEntityManager().appendReference("sourceEdgePtr",this.object.UUID, id_)
				}
			}
 		}

		/** targetEdgePtr **/
 		if results[0][12] != nil{
			idsStr :=results[0][12].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMNDI.Edge"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_targetEdgePtr = append(this.object.M_targetEdgePtr,ids[i])
					GetServer().GetEntityManager().appendReference("targetEdgePtr",this.object.UUID, id_)
				}
			}
 		}

		/** planePtr **/
 		if results[0][13] != nil{
			id :=results[0][13].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Plane"
				id_:= refTypeName + "$$" + id
				this.object.M_planePtr= id
				GetServer().GetEntityManager().appendReference("planePtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMNDIBPMNPlaneEntityFromObject(object *BPMNDI.BPMNPlane) *BPMNDI_BPMNPlaneEntity {
	 return this.NewBPMNDIBPMNPlaneEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMNDI_BPMNPlaneEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMNDIBPMNPlaneExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNPlane"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMNDI_BPMNPlaneEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMNDI_BPMNPlaneEntity) AppendReference(reference Entity) {

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
//              			BPMNShape
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMNDI_BPMNShapeEntity struct{
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
	object *BPMNDI.BPMNShape
}

/** Constructor function **/
func (this *EntityManager) NewBPMNDIBPMNShapeEntity(objectId string, object interface{}) *BPMNDI_BPMNShapeEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMNDIBPMNShapeExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMNDI.BPMNShape).TYPENAME = "BPMNDI.BPMNShape"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMNDI.BPMNShape).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMNDI_BPMNShapeEntity)
		}
	}else{
		uuidStr = "BPMNDI.BPMNShape%" + Utility.RandomUUID()
	}
	entity := new(BPMNDI_BPMNShapeEntity)
	if object == nil{
		entity.object = new(BPMNDI.BPMNShape)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMNDI.BPMNShape)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMNDI.BPMNShape"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMNDI.BPMNShape","BPMN20")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMNDI_BPMNShapeEntity) GetTypeName()string{
	return "BPMNDI.BPMNShape"
}
func(this *BPMNDI_BPMNShapeEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMNDI_BPMNShapeEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMNDI_BPMNShapeEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMNDI_BPMNShapeEntity) AppendReferenced(name string, owner Entity){
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

func(this *BPMNDI_BPMNShapeEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMNDI_BPMNShapeEntity) RemoveReferenced(name string, owner Entity) {
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

func(this *BPMNDI_BPMNShapeEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMNDI_BPMNShapeEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMNDI_BPMNShapeEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMNDI_BPMNShapeEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMNDI_BPMNShapeEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMNDI_BPMNShapeEntity) RemoveChild(name string, uuid string) {
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

func(this *BPMNDI_BPMNShapeEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMNDI_BPMNShapeEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMNDI_BPMNShapeEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMNDI_BPMNShapeEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMNDI_BPMNShapeEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMNDI_BPMNShapeEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMNDI_BPMNShapeEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMNDI_BPMNShapeEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMNDI_BPMNShapeEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMNDI_BPMNShapeEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMNDI_BPMNShapeEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNShape"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMNDI_BPMNShapeEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) create_BPMNDI_BPMNShapeEntityPrototype() {

	var bPMNShapeEntityProto EntityPrototype
	bPMNShapeEntityProto.TypeName = "BPMNDI.BPMNShape"
	bPMNShapeEntityProto.SuperTypeNames = append(bPMNShapeEntityProto.SuperTypeNames, "DI.DiagramElement")
	bPMNShapeEntityProto.SuperTypeNames = append(bPMNShapeEntityProto.SuperTypeNames, "DI.Node")
	bPMNShapeEntityProto.SuperTypeNames = append(bPMNShapeEntityProto.SuperTypeNames, "DI.Shape")
	bPMNShapeEntityProto.SuperTypeNames = append(bPMNShapeEntityProto.SuperTypeNames, "DI.LabeledShape")
	bPMNShapeEntityProto.Ids = append(bPMNShapeEntityProto.Ids,"uuid")
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"uuid")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"xs.string")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,0)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,false)
	bPMNShapeEntityProto.Indexs = append(bPMNShapeEntityProto.Indexs,"parentUuid")
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"parentUuid")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"xs.string")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,1)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,false)

	/** members of DiagramElement **/
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,2)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_owningDiagram")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"BPMNDI.DI.Diagram:Ref")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,3)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_owningElement")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"BPMNDI.DI.DiagramElement:Ref")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,4)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_modelElement")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"xs.interface{}:Ref")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,5)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_style")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"BPMNDI.DI.Style:Ref")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,6)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_ownedElement")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"[]BPMNDI.DI.DiagramElement")
	bPMNShapeEntityProto.Ids = append(bPMNShapeEntityProto.Ids,"M_id")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,7)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_id")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"xs.ID")

	/** members of Node **/
	/** No members **/

	/** members of Shape **/
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,8)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_Bounds")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"BPMNDI.DC.Bounds")

	/** members of LabeledShape **/
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,9)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_ownedLabel")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"[]BPMNDI.DI.Label")

	/** members of BPMNShape **/
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,10)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_bpmnElement")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"xs.interface{}:Ref")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,11)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_isHorizontal")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"xs.bool")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,12)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_isExpanded")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"xs.bool")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,13)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_isMarkerVisible")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"xs.bool")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,14)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_BPMNLabel")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"BPMNDI.BPMNLabel")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,15)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_isMessageVisible")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"xs.bool")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,16)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_participantBandKind")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"enum:ParticipantBandKind_Top_initiating:ParticipantBandKind_Middle_initiating:ParticipantBandKind_Bottom_initiating:ParticipantBandKind_Top_non_initiating:ParticipantBandKind_Middle_non_initiating:ParticipantBandKind_Bottom_non_initiating")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,17)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,true)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_choreographyActivityShape")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"BPMNDI.BPMNShape:Ref")

	/** associations of BPMNShape **/
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,18)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,false)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_participantBandShapePtr")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"BPMNDI.BPMNShape:Ref")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,19)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,false)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_sourceEdgePtr")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"[]BPMNDI.DI.Edge:Ref")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,20)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,false)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_targetEdgePtr")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"[]BPMNDI.DI.Edge:Ref")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,21)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,false)
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"M_planePtr")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"BPMNDI.DI.Plane:Ref")
	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"childsUuid")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"[]xs.string")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,22)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,false)

	bPMNShapeEntityProto.Fields = append(bPMNShapeEntityProto.Fields,"referenced")
	bPMNShapeEntityProto.FieldsType = append(bPMNShapeEntityProto.FieldsType,"[]EntityRef")
	bPMNShapeEntityProto.FieldsOrder = append(bPMNShapeEntityProto.FieldsOrder,23)
	bPMNShapeEntityProto.FieldsVisibility = append(bPMNShapeEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMNDIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&bPMNShapeEntityProto)

}

/** Create **/
func (this *BPMNDI_BPMNShapeEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMNDI.BPMNShape"

	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNShape"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of DiagramElement **/
	query.Fields = append(query.Fields, "M_owningDiagram")
	query.Fields = append(query.Fields, "M_owningElement")
	query.Fields = append(query.Fields, "M_modelElement")
	query.Fields = append(query.Fields, "M_style")
	query.Fields = append(query.Fields, "M_ownedElement")
	query.Fields = append(query.Fields, "M_id")

	/** members of Node **/
	/** No members **/

	/** members of Shape **/
	query.Fields = append(query.Fields, "M_Bounds")

	/** members of LabeledShape **/
	query.Fields = append(query.Fields, "M_ownedLabel")

	/** members of BPMNShape **/
	query.Fields = append(query.Fields, "M_bpmnElement")
	query.Fields = append(query.Fields, "M_isHorizontal")
	query.Fields = append(query.Fields, "M_isExpanded")
	query.Fields = append(query.Fields, "M_isMarkerVisible")
	query.Fields = append(query.Fields, "M_BPMNLabel")
	query.Fields = append(query.Fields, "M_isMessageVisible")
	query.Fields = append(query.Fields, "M_participantBandKind")
	query.Fields = append(query.Fields, "M_choreographyActivityShape")

		/** associations of BPMNShape **/
	query.Fields = append(query.Fields, "M_participantBandShapePtr")
	query.Fields = append(query.Fields, "M_sourceEdgePtr")
	query.Fields = append(query.Fields, "M_targetEdgePtr")
	query.Fields = append(query.Fields, "M_planePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var BPMNShapeInfo []interface{}

	BPMNShapeInfo = append(BPMNShapeInfo, this.GetUuid())
	if this.parentPtr != nil {
		BPMNShapeInfo = append(BPMNShapeInfo, this.parentPtr.GetUuid())
	}else{
		BPMNShapeInfo = append(BPMNShapeInfo, "")
	}

	/** members of DiagramElement **/

	/** Save owningDiagram type Diagram **/
		BPMNShapeInfo = append(BPMNShapeInfo,this.object.M_owningDiagram)

	/** Save owningElement type DiagramElement **/
		BPMNShapeInfo = append(BPMNShapeInfo,this.object.M_owningElement)
	BPMNShapeInfo = append(BPMNShapeInfo, this.object.M_modelElement)

	/** Save style type Style **/
		BPMNShapeInfo = append(BPMNShapeInfo,this.object.M_style)

	/** Save ownedElement type DiagramElement **/
	ownedElementIds := make([]string,0)
	for i := 0; i < len(this.object.M_ownedElement); i++ {
		switch v := this.object.M_ownedElement[i].(type) {
		case *BPMNDI.BPMNShape:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNPlane:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNLabel:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNEdge:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		}
	}
	ownedElementStr, _ := json.Marshal(ownedElementIds)
	BPMNShapeInfo = append(BPMNShapeInfo, string(ownedElementStr))
	BPMNShapeInfo = append(BPMNShapeInfo, this.object.M_id)

	/** members of Node **/
	/** No members **/

	/** members of Shape **/

	/** Save Bounds type Bounds **/
	if this.object.M_Bounds != nil {
		BoundsEntity:= GetServer().GetEntityManager().NewDCBoundsEntity(this.object.M_Bounds.UUID, this.object.M_Bounds)
		BPMNShapeInfo = append(BPMNShapeInfo, BoundsEntity.uuid)
		BoundsEntity.AppendReferenced("Bounds", this)
		this.AppendChild("Bounds",BoundsEntity)
		if BoundsEntity.NeedSave() {
			BoundsEntity.SaveEntity()
		}
	}else{
		BPMNShapeInfo = append(BPMNShapeInfo, "")
	}

	/** members of LabeledShape **/

	/** Save ownedLabel type Label **/
	ownedLabelIds := make([]string,0)
	for i := 0; i < len(this.object.M_ownedLabel); i++ {
		switch v := this.object.M_ownedLabel[i].(type) {
		case *BPMNDI.BPMNLabel:
		ownedLabelEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
		ownedLabelIds=append(ownedLabelIds,ownedLabelEntity.uuid)
		ownedLabelEntity.AppendReferenced("ownedLabel", this)
		this.AppendChild("ownedLabel",ownedLabelEntity)
		if ownedLabelEntity.NeedSave() {
			ownedLabelEntity.SaveEntity()
		}
		}
	}
	ownedLabelStr, _ := json.Marshal(ownedLabelIds)
	BPMNShapeInfo = append(BPMNShapeInfo, string(ownedLabelStr))

	/** members of BPMNShape **/
	BPMNShapeInfo = append(BPMNShapeInfo, this.object.M_bpmnElement)
	BPMNShapeInfo = append(BPMNShapeInfo, this.object.M_isHorizontal)
	BPMNShapeInfo = append(BPMNShapeInfo, this.object.M_isExpanded)
	BPMNShapeInfo = append(BPMNShapeInfo, this.object.M_isMarkerVisible)

	/** Save BPMNLabel type BPMNLabel **/
	if this.object.M_BPMNLabel != nil {
		BPMNLabelEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(this.object.M_BPMNLabel.UUID, this.object.M_BPMNLabel)
		BPMNShapeInfo = append(BPMNShapeInfo, BPMNLabelEntity.uuid)
		BPMNLabelEntity.AppendReferenced("BPMNLabel", this)
		this.AppendChild("BPMNLabel",BPMNLabelEntity)
		if BPMNLabelEntity.NeedSave() {
			BPMNLabelEntity.SaveEntity()
		}
	}else{
		BPMNShapeInfo = append(BPMNShapeInfo, "")
	}
	BPMNShapeInfo = append(BPMNShapeInfo, this.object.M_isMessageVisible)

	/** Save participantBandKind type ParticipantBandKind **/
	if this.object.M_participantBandKind==BPMNDI.ParticipantBandKind_Top_initiating{
		BPMNShapeInfo = append(BPMNShapeInfo, 0)
	} else if this.object.M_participantBandKind==BPMNDI.ParticipantBandKind_Middle_initiating{
		BPMNShapeInfo = append(BPMNShapeInfo, 1)
	} else if this.object.M_participantBandKind==BPMNDI.ParticipantBandKind_Bottom_initiating{
		BPMNShapeInfo = append(BPMNShapeInfo, 2)
	} else if this.object.M_participantBandKind==BPMNDI.ParticipantBandKind_Top_non_initiating{
		BPMNShapeInfo = append(BPMNShapeInfo, 3)
	} else if this.object.M_participantBandKind==BPMNDI.ParticipantBandKind_Middle_non_initiating{
		BPMNShapeInfo = append(BPMNShapeInfo, 4)
	} else if this.object.M_participantBandKind==BPMNDI.ParticipantBandKind_Bottom_non_initiating{
		BPMNShapeInfo = append(BPMNShapeInfo, 5)
	}else{
		BPMNShapeInfo = append(BPMNShapeInfo, 0)
	}

	/** Save choreographyActivityShape type BPMNShape **/
		BPMNShapeInfo = append(BPMNShapeInfo,this.object.M_choreographyActivityShape)

	/** associations of BPMNShape **/

	/** Save participantBandShape type BPMNShape **/
		BPMNShapeInfo = append(BPMNShapeInfo,this.object.M_participantBandShapePtr)

	/** Save sourceEdge type Edge **/
	sourceEdgePtrStr, _ := json.Marshal(this.object.M_sourceEdgePtr)
	BPMNShapeInfo = append(BPMNShapeInfo, string(sourceEdgePtrStr))

	/** Save targetEdge type Edge **/
	targetEdgePtrStr, _ := json.Marshal(this.object.M_targetEdgePtr)
	BPMNShapeInfo = append(BPMNShapeInfo, string(targetEdgePtrStr))

	/** Save plane type Plane **/
		BPMNShapeInfo = append(BPMNShapeInfo,this.object.M_planePtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	BPMNShapeInfo = append(BPMNShapeInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	BPMNShapeInfo = append(BPMNShapeInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMNDIDB, string(queryStr), BPMNShapeInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMNDIDB, string(queryStr), BPMNShapeInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMNDI_BPMNShapeEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMNDI_BPMNShapeEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNShape"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of DiagramElement **/
	query.Fields = append(query.Fields, "M_owningDiagram")
	query.Fields = append(query.Fields, "M_owningElement")
	query.Fields = append(query.Fields, "M_modelElement")
	query.Fields = append(query.Fields, "M_style")
	query.Fields = append(query.Fields, "M_ownedElement")
	query.Fields = append(query.Fields, "M_id")

	/** members of Node **/
	/** No members **/

	/** members of Shape **/
	query.Fields = append(query.Fields, "M_Bounds")

	/** members of LabeledShape **/
	query.Fields = append(query.Fields, "M_ownedLabel")

	/** members of BPMNShape **/
	query.Fields = append(query.Fields, "M_bpmnElement")
	query.Fields = append(query.Fields, "M_isHorizontal")
	query.Fields = append(query.Fields, "M_isExpanded")
	query.Fields = append(query.Fields, "M_isMarkerVisible")
	query.Fields = append(query.Fields, "M_BPMNLabel")
	query.Fields = append(query.Fields, "M_isMessageVisible")
	query.Fields = append(query.Fields, "M_participantBandKind")
	query.Fields = append(query.Fields, "M_choreographyActivityShape")

		/** associations of BPMNShape **/
	query.Fields = append(query.Fields, "M_participantBandShapePtr")
	query.Fields = append(query.Fields, "M_sourceEdgePtr")
	query.Fields = append(query.Fields, "M_targetEdgePtr")
	query.Fields = append(query.Fields, "M_planePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of BPMNShape...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMNDI.BPMNShape)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMNDI.BPMNShape"

		this.parentUuid = results[0][1].(string)

		/** members of DiagramElement **/

		/** owningDiagram **/
 		if results[0][2] != nil{
			id :=results[0][2].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Diagram"
				id_:= refTypeName + "$$" + id
				this.object.M_owningDiagram= id
				GetServer().GetEntityManager().appendReference("owningDiagram",this.object.UUID, id_)
			}
 		}

		/** owningElement **/
 		if results[0][3] != nil{
			id :=results[0][3].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.DiagramElement"
				id_:= refTypeName + "$$" + id
				this.object.M_owningElement= id
				GetServer().GetEntityManager().appendReference("owningElement",this.object.UUID, id_)
			}
 		}

		/** modelElement **/
 		if results[0][4] != nil{
			id :=results[0][4].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.interface{}"
				id_:= refTypeName + "$$" + id
				this.object.M_modelElement= id
				GetServer().GetEntityManager().appendReference("modelElement",this.object.UUID, id_)
			}
 		}

		/** style **/
 		if results[0][5] != nil{
			id :=results[0][5].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Style"
				id_:= refTypeName + "$$" + id
				this.object.M_style= id
				GetServer().GetEntityManager().appendReference("style",this.object.UUID, id_)
			}
 		}

		/** ownedElement **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
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
				if typeName == "BPMNDI.BPMNShape"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNShapeEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNShapeEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNPlane"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNPlaneEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNPlaneEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNLabel"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNLabelEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNLabelEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNEdge"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNEdgeEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNEdgeEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				}
 			}
 		}

		/** id **/
 		if results[0][7] != nil{
 			this.object.M_id=results[0][7].(string)
 		}

		/** members of Node **/
		/** No members **/

		/** members of Shape **/

		/** Bounds **/
 		if results[0][8] != nil{
			uuid :=results[0][8].(string)
			if len(uuid) > 0 {
				var BoundsEntity *DC_BoundsEntity
				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
					BoundsEntity = instance.(*DC_BoundsEntity)
				}else{
					BoundsEntity = GetServer().GetEntityManager().NewDCBoundsEntity(uuid, nil)
					BoundsEntity.InitEntity(uuid)
					GetServer().GetEntityManager().insert( BoundsEntity)
				}
				BoundsEntity.AppendReferenced("Bounds", this)
				this.AppendChild("Bounds",BoundsEntity)
			}
 		}

		/** members of LabeledShape **/

		/** ownedLabel **/
 		if results[0][9] != nil{
			uuidsStr :=results[0][9].(string)
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
				if typeName == "BPMNDI.BPMNLabel"{
						if len(uuids[i]) > 0 {
							var ownedLabelEntity *BPMNDI_BPMNLabelEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedLabelEntity = instance.(*BPMNDI_BPMNLabelEntity)
							}else{
								ownedLabelEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuids[i], nil)
								ownedLabelEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedLabelEntity)
							}
							ownedLabelEntity.AppendReferenced("ownedLabel", this)
							this.AppendChild("ownedLabel",ownedLabelEntity)
						}
				}
 			}
 		}

		/** members of BPMNShape **/

		/** bpmnElement **/
 		if results[0][10] != nil{
			id :=results[0][10].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.interface{}"
				id_:= refTypeName + "$$" + id
				this.object.M_bpmnElement= id
				GetServer().GetEntityManager().appendReference("bpmnElement",this.object.UUID, id_)
				this.object.M_bpmnElement = id
			}
 		}

		/** isHorizontal **/
 		if results[0][11] != nil{
 			this.object.M_isHorizontal=results[0][11].(bool)
 		}

		/** isExpanded **/
 		if results[0][12] != nil{
 			this.object.M_isExpanded=results[0][12].(bool)
 		}

		/** isMarkerVisible **/
 		if results[0][13] != nil{
 			this.object.M_isMarkerVisible=results[0][13].(bool)
 		}

		/** BPMNLabel **/
 		if results[0][14] != nil{
			uuid :=results[0][14].(string)
			if len(uuid) > 0 {
				var BPMNLabelEntity *BPMNDI_BPMNLabelEntity
				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
					BPMNLabelEntity = instance.(*BPMNDI_BPMNLabelEntity)
				}else{
					BPMNLabelEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuid, nil)
					BPMNLabelEntity.InitEntity(uuid)
					GetServer().GetEntityManager().insert( BPMNLabelEntity)
				}
				BPMNLabelEntity.AppendReferenced("BPMNLabel", this)
				this.AppendChild("BPMNLabel",BPMNLabelEntity)
			}
 		}

		/** isMessageVisible **/
 		if results[0][15] != nil{
 			this.object.M_isMessageVisible=results[0][15].(bool)
 		}

		/** participantBandKind **/
 		if results[0][16] != nil{
 			enumIndex := results[0][16].(int)
			if enumIndex == 0{
 				this.object.M_participantBandKind=BPMNDI.ParticipantBandKind_Top_initiating
			} else if enumIndex == 1{
 				this.object.M_participantBandKind=BPMNDI.ParticipantBandKind_Middle_initiating
			} else if enumIndex == 2{
 				this.object.M_participantBandKind=BPMNDI.ParticipantBandKind_Bottom_initiating
			} else if enumIndex == 3{
 				this.object.M_participantBandKind=BPMNDI.ParticipantBandKind_Top_non_initiating
			} else if enumIndex == 4{
 				this.object.M_participantBandKind=BPMNDI.ParticipantBandKind_Middle_non_initiating
			} else if enumIndex == 5{
 				this.object.M_participantBandKind=BPMNDI.ParticipantBandKind_Bottom_non_initiating
 			}
 		}

		/** choreographyActivityShape **/
 		if results[0][17] != nil{
			id :=results[0][17].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.BPMNShape"
				id_:= refTypeName + "$$" + id
				this.object.M_choreographyActivityShape= id
				GetServer().GetEntityManager().appendReference("choreographyActivityShape",this.object.UUID, id_)
			}
 		}

		/** associations of BPMNShape **/

		/** participantBandShapePtr **/
 		if results[0][18] != nil{
			id :=results[0][18].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.BPMNShape"
				id_:= refTypeName + "$$" + id
				this.object.M_participantBandShapePtr= id
				GetServer().GetEntityManager().appendReference("participantBandShapePtr",this.object.UUID, id_)
			}
 		}

		/** sourceEdgePtr **/
 		if results[0][19] != nil{
			idsStr :=results[0][19].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMNDI.Edge"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_sourceEdgePtr = append(this.object.M_sourceEdgePtr,ids[i])
					GetServer().GetEntityManager().appendReference("sourceEdgePtr",this.object.UUID, id_)
				}
			}
 		}

		/** targetEdgePtr **/
 		if results[0][20] != nil{
			idsStr :=results[0][20].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMNDI.Edge"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_targetEdgePtr = append(this.object.M_targetEdgePtr,ids[i])
					GetServer().GetEntityManager().appendReference("targetEdgePtr",this.object.UUID, id_)
				}
			}
 		}

		/** planePtr **/
 		if results[0][21] != nil{
			id :=results[0][21].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Plane"
				id_:= refTypeName + "$$" + id
				this.object.M_planePtr= id
				GetServer().GetEntityManager().appendReference("planePtr",this.object.UUID, id_)
			}
 		}
		childsUuidStr := results[0][22].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][23].(string)
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
func (this *EntityManager) NewBPMNDIBPMNShapeEntityFromObject(object *BPMNDI.BPMNShape) *BPMNDI_BPMNShapeEntity {
	 return this.NewBPMNDIBPMNShapeEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMNDI_BPMNShapeEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMNDIBPMNShapeExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNShape"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMNDI_BPMNShapeEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMNDI_BPMNShapeEntity) AppendReference(reference Entity) {

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
//              			BPMNEdge
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMNDI_BPMNEdgeEntity struct{
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
	object *BPMNDI.BPMNEdge
}

/** Constructor function **/
func (this *EntityManager) NewBPMNDIBPMNEdgeEntity(objectId string, object interface{}) *BPMNDI_BPMNEdgeEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMNDIBPMNEdgeExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMNDI.BPMNEdge).TYPENAME = "BPMNDI.BPMNEdge"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMNDI.BPMNEdge).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMNDI_BPMNEdgeEntity)
		}
	}else{
		uuidStr = "BPMNDI.BPMNEdge%" + Utility.RandomUUID()
	}
	entity := new(BPMNDI_BPMNEdgeEntity)
	if object == nil{
		entity.object = new(BPMNDI.BPMNEdge)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMNDI.BPMNEdge)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMNDI.BPMNEdge"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMNDI.BPMNEdge","BPMN20")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMNDI_BPMNEdgeEntity) GetTypeName()string{
	return "BPMNDI.BPMNEdge"
}
func(this *BPMNDI_BPMNEdgeEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMNDI_BPMNEdgeEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMNDI_BPMNEdgeEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMNDI_BPMNEdgeEntity) AppendReferenced(name string, owner Entity){
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

func(this *BPMNDI_BPMNEdgeEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMNDI_BPMNEdgeEntity) RemoveReferenced(name string, owner Entity) {
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

func(this *BPMNDI_BPMNEdgeEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMNDI_BPMNEdgeEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMNDI_BPMNEdgeEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMNDI_BPMNEdgeEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMNDI_BPMNEdgeEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMNDI_BPMNEdgeEntity) RemoveChild(name string, uuid string) {
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

func(this *BPMNDI_BPMNEdgeEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMNDI_BPMNEdgeEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMNDI_BPMNEdgeEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMNDI_BPMNEdgeEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMNDI_BPMNEdgeEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMNDI_BPMNEdgeEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMNDI_BPMNEdgeEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMNDI_BPMNEdgeEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMNDI_BPMNEdgeEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMNDI_BPMNEdgeEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMNDI_BPMNEdgeEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNEdge"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMNDI_BPMNEdgeEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) create_BPMNDI_BPMNEdgeEntityPrototype() {

	var bPMNEdgeEntityProto EntityPrototype
	bPMNEdgeEntityProto.TypeName = "BPMNDI.BPMNEdge"
	bPMNEdgeEntityProto.SuperTypeNames = append(bPMNEdgeEntityProto.SuperTypeNames, "DI.DiagramElement")
	bPMNEdgeEntityProto.SuperTypeNames = append(bPMNEdgeEntityProto.SuperTypeNames, "DI.Edge")
	bPMNEdgeEntityProto.SuperTypeNames = append(bPMNEdgeEntityProto.SuperTypeNames, "DI.LabeledEdge")
	bPMNEdgeEntityProto.Ids = append(bPMNEdgeEntityProto.Ids,"uuid")
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"uuid")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"xs.string")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,0)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,false)
	bPMNEdgeEntityProto.Indexs = append(bPMNEdgeEntityProto.Indexs,"parentUuid")
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"parentUuid")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"xs.string")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,1)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,false)

	/** members of DiagramElement **/
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,2)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_owningDiagram")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"BPMNDI.DI.Diagram:Ref")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,3)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_owningElement")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"BPMNDI.DI.DiagramElement:Ref")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,4)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_modelElement")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"xs.interface{}:Ref")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,5)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_style")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"BPMNDI.DI.Style:Ref")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,6)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_ownedElement")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"[]BPMNDI.DI.DiagramElement")
	bPMNEdgeEntityProto.Ids = append(bPMNEdgeEntityProto.Ids,"M_id")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,7)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_id")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"xs.ID")

	/** members of Edge **/
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,8)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_source")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"BPMNDI.DI.DiagramElement")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,9)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_target")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"BPMNDI.DI.DiagramElement")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,10)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_waypoint")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"[]BPMNDI.DC.Point")

	/** members of LabeledEdge **/
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,11)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_ownedLabel")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"[]BPMNDI.DI.Label")

	/** members of BPMNEdge **/
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,12)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_BPMNLabel")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"BPMNDI.BPMNLabel")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,13)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_bpmnElement")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"xs.interface{}:Ref")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,14)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_sourceElement")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"BPMNDI.DI.DiagramElement")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,15)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_targetElement")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"BPMNDI.DI.DiagramElement:Ref")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,16)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,true)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_messageVisibleKind")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"enum:MessageVisibleKind_Initiating:MessageVisibleKind_Non_initiating")

	/** associations of BPMNEdge **/
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,17)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,false)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_sourceEdgePtr")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"[]BPMNDI.DI.Edge:Ref")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,18)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,false)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_targetEdgePtr")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"[]BPMNDI.DI.Edge:Ref")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,19)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,false)
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"M_planePtr")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"BPMNDI.DI.Plane:Ref")
	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"childsUuid")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"[]xs.string")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,20)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,false)

	bPMNEdgeEntityProto.Fields = append(bPMNEdgeEntityProto.Fields,"referenced")
	bPMNEdgeEntityProto.FieldsType = append(bPMNEdgeEntityProto.FieldsType,"[]EntityRef")
	bPMNEdgeEntityProto.FieldsOrder = append(bPMNEdgeEntityProto.FieldsOrder,21)
	bPMNEdgeEntityProto.FieldsVisibility = append(bPMNEdgeEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMNDIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&bPMNEdgeEntityProto)

}

/** Create **/
func (this *BPMNDI_BPMNEdgeEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMNDI.BPMNEdge"

	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNEdge"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of DiagramElement **/
	query.Fields = append(query.Fields, "M_owningDiagram")
	query.Fields = append(query.Fields, "M_owningElement")
	query.Fields = append(query.Fields, "M_modelElement")
	query.Fields = append(query.Fields, "M_style")
	query.Fields = append(query.Fields, "M_ownedElement")
	query.Fields = append(query.Fields, "M_id")

	/** members of Edge **/
	query.Fields = append(query.Fields, "M_source")
	query.Fields = append(query.Fields, "M_target")
	query.Fields = append(query.Fields, "M_waypoint")

	/** members of LabeledEdge **/
	query.Fields = append(query.Fields, "M_ownedLabel")

	/** members of BPMNEdge **/
	query.Fields = append(query.Fields, "M_BPMNLabel")
	query.Fields = append(query.Fields, "M_bpmnElement")
	query.Fields = append(query.Fields, "M_sourceElement")
	query.Fields = append(query.Fields, "M_targetElement")
	query.Fields = append(query.Fields, "M_messageVisibleKind")

		/** associations of BPMNEdge **/
	query.Fields = append(query.Fields, "M_sourceEdgePtr")
	query.Fields = append(query.Fields, "M_targetEdgePtr")
	query.Fields = append(query.Fields, "M_planePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var BPMNEdgeInfo []interface{}

	BPMNEdgeInfo = append(BPMNEdgeInfo, this.GetUuid())
	if this.parentPtr != nil {
		BPMNEdgeInfo = append(BPMNEdgeInfo, this.parentPtr.GetUuid())
	}else{
		BPMNEdgeInfo = append(BPMNEdgeInfo, "")
	}

	/** members of DiagramElement **/

	/** Save owningDiagram type Diagram **/
		BPMNEdgeInfo = append(BPMNEdgeInfo,this.object.M_owningDiagram)

	/** Save owningElement type DiagramElement **/
		BPMNEdgeInfo = append(BPMNEdgeInfo,this.object.M_owningElement)
	BPMNEdgeInfo = append(BPMNEdgeInfo, this.object.M_modelElement)

	/** Save style type Style **/
		BPMNEdgeInfo = append(BPMNEdgeInfo,this.object.M_style)

	/** Save ownedElement type DiagramElement **/
	ownedElementIds := make([]string,0)
	for i := 0; i < len(this.object.M_ownedElement); i++ {
		switch v := this.object.M_ownedElement[i].(type) {
		case *BPMNDI.BPMNPlane:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNLabel:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNShape:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNEdge:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		}
	}
	ownedElementStr, _ := json.Marshal(ownedElementIds)
	BPMNEdgeInfo = append(BPMNEdgeInfo, string(ownedElementStr))
	BPMNEdgeInfo = append(BPMNEdgeInfo, this.object.M_id)

	/** members of Edge **/

	/** Save source type DiagramElement **/
	if this.object.M_source != nil {
		switch v := this.object.M_source.(type) {
		case *BPMNDI.BPMNShape:
			sourceEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, sourceEntity.uuid)
		    sourceEntity.AppendReferenced("source", this)
			this.AppendChild("source",sourceEntity)
			if sourceEntity.NeedSave() {
				sourceEntity.SaveEntity()
			}
		case *BPMNDI.BPMNPlane:
			sourceEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, sourceEntity.uuid)
		    sourceEntity.AppendReferenced("source", this)
			this.AppendChild("source",sourceEntity)
			if sourceEntity.NeedSave() {
				sourceEntity.SaveEntity()
			}
		case *BPMNDI.BPMNLabel:
			sourceEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, sourceEntity.uuid)
		    sourceEntity.AppendReferenced("source", this)
			this.AppendChild("source",sourceEntity)
			if sourceEntity.NeedSave() {
				sourceEntity.SaveEntity()
			}
		case *BPMNDI.BPMNEdge:
			sourceEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, sourceEntity.uuid)
		    sourceEntity.AppendReferenced("source", this)
			this.AppendChild("source",sourceEntity)
			if sourceEntity.NeedSave() {
				sourceEntity.SaveEntity()
			}
			}
	}else{
		BPMNEdgeInfo = append(BPMNEdgeInfo, "")
	}

	/** Save target type DiagramElement **/
	if this.object.M_target != nil {
		switch v := this.object.M_target.(type) {
		case *BPMNDI.BPMNLabel:
			targetEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, targetEntity.uuid)
		    targetEntity.AppendReferenced("target", this)
			this.AppendChild("target",targetEntity)
			if targetEntity.NeedSave() {
				targetEntity.SaveEntity()
			}
		case *BPMNDI.BPMNShape:
			targetEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, targetEntity.uuid)
		    targetEntity.AppendReferenced("target", this)
			this.AppendChild("target",targetEntity)
			if targetEntity.NeedSave() {
				targetEntity.SaveEntity()
			}
		case *BPMNDI.BPMNPlane:
			targetEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, targetEntity.uuid)
		    targetEntity.AppendReferenced("target", this)
			this.AppendChild("target",targetEntity)
			if targetEntity.NeedSave() {
				targetEntity.SaveEntity()
			}
		case *BPMNDI.BPMNEdge:
			targetEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, targetEntity.uuid)
		    targetEntity.AppendReferenced("target", this)
			this.AppendChild("target",targetEntity)
			if targetEntity.NeedSave() {
				targetEntity.SaveEntity()
			}
			}
	}else{
		BPMNEdgeInfo = append(BPMNEdgeInfo, "")
	}

	/** Save waypoint type Point **/
	waypointIds := make([]string,0)
	for i := 0; i < len(this.object.M_waypoint); i++ {
		waypointEntity:= GetServer().GetEntityManager().NewDCPointEntity(this.object.M_waypoint[i].UUID,this.object.M_waypoint[i])
		waypointIds=append(waypointIds,waypointEntity.uuid)
		waypointEntity.AppendReferenced("waypoint", this)
		this.AppendChild("waypoint",waypointEntity)
		if waypointEntity.NeedSave() {
			waypointEntity.SaveEntity()
		}
	}
	waypointStr, _ := json.Marshal(waypointIds)
	BPMNEdgeInfo = append(BPMNEdgeInfo, string(waypointStr))

	/** members of LabeledEdge **/

	/** Save ownedLabel type Label **/
	ownedLabelIds := make([]string,0)
	for i := 0; i < len(this.object.M_ownedLabel); i++ {
		switch v := this.object.M_ownedLabel[i].(type) {
		case *BPMNDI.BPMNLabel:
		ownedLabelEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
		ownedLabelIds=append(ownedLabelIds,ownedLabelEntity.uuid)
		ownedLabelEntity.AppendReferenced("ownedLabel", this)
		this.AppendChild("ownedLabel",ownedLabelEntity)
		if ownedLabelEntity.NeedSave() {
			ownedLabelEntity.SaveEntity()
		}
		}
	}
	ownedLabelStr, _ := json.Marshal(ownedLabelIds)
	BPMNEdgeInfo = append(BPMNEdgeInfo, string(ownedLabelStr))

	/** members of BPMNEdge **/

	/** Save BPMNLabel type BPMNLabel **/
	if this.object.M_BPMNLabel != nil {
		BPMNLabelEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(this.object.M_BPMNLabel.UUID, this.object.M_BPMNLabel)
		BPMNEdgeInfo = append(BPMNEdgeInfo, BPMNLabelEntity.uuid)
		BPMNLabelEntity.AppendReferenced("BPMNLabel", this)
		this.AppendChild("BPMNLabel",BPMNLabelEntity)
		if BPMNLabelEntity.NeedSave() {
			BPMNLabelEntity.SaveEntity()
		}
	}else{
		BPMNEdgeInfo = append(BPMNEdgeInfo, "")
	}
	BPMNEdgeInfo = append(BPMNEdgeInfo, this.object.M_bpmnElement)

	/** Save sourceElement type DiagramElement **/
	if this.object.M_sourceElement != nil {
		switch v := this.object.M_sourceElement.(type) {
		case *BPMNDI.BPMNShape:
			sourceElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, sourceElementEntity.uuid)
		    sourceElementEntity.AppendReferenced("sourceElement", this)
			this.AppendChild("sourceElement",sourceElementEntity)
			if sourceElementEntity.NeedSave() {
				sourceElementEntity.SaveEntity()
			}
		case *BPMNDI.BPMNPlane:
			sourceElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, sourceElementEntity.uuid)
		    sourceElementEntity.AppendReferenced("sourceElement", this)
			this.AppendChild("sourceElement",sourceElementEntity)
			if sourceElementEntity.NeedSave() {
				sourceElementEntity.SaveEntity()
			}
		case *BPMNDI.BPMNLabel:
			sourceElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, sourceElementEntity.uuid)
		    sourceElementEntity.AppendReferenced("sourceElement", this)
			this.AppendChild("sourceElement",sourceElementEntity)
			if sourceElementEntity.NeedSave() {
				sourceElementEntity.SaveEntity()
			}
		case *BPMNDI.BPMNEdge:
			sourceElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(v.UUID, v)
			BPMNEdgeInfo = append(BPMNEdgeInfo, sourceElementEntity.uuid)
		    sourceElementEntity.AppendReferenced("sourceElement", this)
			this.AppendChild("sourceElement",sourceElementEntity)
			if sourceElementEntity.NeedSave() {
				sourceElementEntity.SaveEntity()
			}
			}
	}else{
		BPMNEdgeInfo = append(BPMNEdgeInfo, "")
	}

	/** Save targetElement type DiagramElement **/
		BPMNEdgeInfo = append(BPMNEdgeInfo,this.object.M_targetElement)

	/** Save messageVisibleKind type MessageVisibleKind **/
	if this.object.M_messageVisibleKind==BPMNDI.MessageVisibleKind_Initiating{
		BPMNEdgeInfo = append(BPMNEdgeInfo, 0)
	} else if this.object.M_messageVisibleKind==BPMNDI.MessageVisibleKind_Non_initiating{
		BPMNEdgeInfo = append(BPMNEdgeInfo, 1)
	}else{
		BPMNEdgeInfo = append(BPMNEdgeInfo, 0)
	}

	/** associations of BPMNEdge **/

	/** Save sourceEdge type Edge **/
	sourceEdgePtrStr, _ := json.Marshal(this.object.M_sourceEdgePtr)
	BPMNEdgeInfo = append(BPMNEdgeInfo, string(sourceEdgePtrStr))

	/** Save targetEdge type Edge **/
	targetEdgePtrStr, _ := json.Marshal(this.object.M_targetEdgePtr)
	BPMNEdgeInfo = append(BPMNEdgeInfo, string(targetEdgePtrStr))

	/** Save plane type Plane **/
		BPMNEdgeInfo = append(BPMNEdgeInfo,this.object.M_planePtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	BPMNEdgeInfo = append(BPMNEdgeInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	BPMNEdgeInfo = append(BPMNEdgeInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMNDIDB, string(queryStr), BPMNEdgeInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMNDIDB, string(queryStr), BPMNEdgeInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMNDI_BPMNEdgeEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMNDI_BPMNEdgeEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNEdge"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of DiagramElement **/
	query.Fields = append(query.Fields, "M_owningDiagram")
	query.Fields = append(query.Fields, "M_owningElement")
	query.Fields = append(query.Fields, "M_modelElement")
	query.Fields = append(query.Fields, "M_style")
	query.Fields = append(query.Fields, "M_ownedElement")
	query.Fields = append(query.Fields, "M_id")

	/** members of Edge **/
	query.Fields = append(query.Fields, "M_source")
	query.Fields = append(query.Fields, "M_target")
	query.Fields = append(query.Fields, "M_waypoint")

	/** members of LabeledEdge **/
	query.Fields = append(query.Fields, "M_ownedLabel")

	/** members of BPMNEdge **/
	query.Fields = append(query.Fields, "M_BPMNLabel")
	query.Fields = append(query.Fields, "M_bpmnElement")
	query.Fields = append(query.Fields, "M_sourceElement")
	query.Fields = append(query.Fields, "M_targetElement")
	query.Fields = append(query.Fields, "M_messageVisibleKind")

		/** associations of BPMNEdge **/
	query.Fields = append(query.Fields, "M_sourceEdgePtr")
	query.Fields = append(query.Fields, "M_targetEdgePtr")
	query.Fields = append(query.Fields, "M_planePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of BPMNEdge...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMNDI.BPMNEdge)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMNDI.BPMNEdge"

		this.parentUuid = results[0][1].(string)

		/** members of DiagramElement **/

		/** owningDiagram **/
 		if results[0][2] != nil{
			id :=results[0][2].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Diagram"
				id_:= refTypeName + "$$" + id
				this.object.M_owningDiagram= id
				GetServer().GetEntityManager().appendReference("owningDiagram",this.object.UUID, id_)
			}
 		}

		/** owningElement **/
 		if results[0][3] != nil{
			id :=results[0][3].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.DiagramElement"
				id_:= refTypeName + "$$" + id
				this.object.M_owningElement= id
				GetServer().GetEntityManager().appendReference("owningElement",this.object.UUID, id_)
			}
 		}

		/** modelElement **/
 		if results[0][4] != nil{
			id :=results[0][4].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.interface{}"
				id_:= refTypeName + "$$" + id
				this.object.M_modelElement= id
				GetServer().GetEntityManager().appendReference("modelElement",this.object.UUID, id_)
			}
 		}

		/** style **/
 		if results[0][5] != nil{
			id :=results[0][5].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Style"
				id_:= refTypeName + "$$" + id
				this.object.M_style= id
				GetServer().GetEntityManager().appendReference("style",this.object.UUID, id_)
			}
 		}

		/** ownedElement **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
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
				if typeName == "BPMNDI.BPMNPlane"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNPlaneEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNPlaneEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNLabel"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNLabelEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNLabelEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNShape"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNShapeEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNShapeEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNEdge"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNEdgeEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNEdgeEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				}
 			}
 		}

		/** id **/
 		if results[0][7] != nil{
 			this.object.M_id=results[0][7].(string)
 		}

		/** members of Edge **/

		/** source **/
 		if results[0][8] != nil{
			uuid :=results[0][8].(string)
			if len(uuid) > 0 {
				typeName := uuid[0:strings.Index(uuid, "%")]
				if err!=nil{
					log.Println("type ", typeName, " not found!")
					return err
				}
			if typeName == "BPMNDI.BPMNShape"{
				if len(uuid) > 0 {
					var sourceEntity *BPMNDI_BPMNShapeEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						sourceEntity = instance.(*BPMNDI_BPMNShapeEntity)
					}else{
						sourceEntity = GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(uuid, nil)
						sourceEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(sourceEntity)
					}
					sourceEntity.AppendReferenced("source", this)
					this.AppendChild("source",sourceEntity)
				}
			} else if typeName == "BPMNDI.BPMNPlane"{
				if len(uuid) > 0 {
					var sourceEntity *BPMNDI_BPMNPlaneEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						sourceEntity = instance.(*BPMNDI_BPMNPlaneEntity)
					}else{
						sourceEntity = GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(uuid, nil)
						sourceEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(sourceEntity)
					}
					sourceEntity.AppendReferenced("source", this)
					this.AppendChild("source",sourceEntity)
				}
			} else if typeName == "BPMNDI.BPMNLabel"{
				if len(uuid) > 0 {
					var sourceEntity *BPMNDI_BPMNLabelEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						sourceEntity = instance.(*BPMNDI_BPMNLabelEntity)
					}else{
						sourceEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuid, nil)
						sourceEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(sourceEntity)
					}
					sourceEntity.AppendReferenced("source", this)
					this.AppendChild("source",sourceEntity)
				}
			} else if typeName == "BPMNDI.BPMNEdge"{
				if len(uuid) > 0 {
					var sourceEntity *BPMNDI_BPMNEdgeEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						sourceEntity = instance.(*BPMNDI_BPMNEdgeEntity)
					}else{
						sourceEntity = GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(uuid, nil)
						sourceEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(sourceEntity)
					}
					sourceEntity.AppendReferenced("source", this)
					this.AppendChild("source",sourceEntity)
				}
			}
			}
 		}

		/** target **/
 		if results[0][9] != nil{
			uuid :=results[0][9].(string)
			if len(uuid) > 0 {
				typeName := uuid[0:strings.Index(uuid, "%")]
				if err!=nil{
					log.Println("type ", typeName, " not found!")
					return err
				}
			if typeName == "BPMNDI.BPMNShape"{
				if len(uuid) > 0 {
					var targetEntity *BPMNDI_BPMNShapeEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						targetEntity = instance.(*BPMNDI_BPMNShapeEntity)
					}else{
						targetEntity = GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(uuid, nil)
						targetEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(targetEntity)
					}
					targetEntity.AppendReferenced("target", this)
					this.AppendChild("target",targetEntity)
				}
			} else if typeName == "BPMNDI.BPMNPlane"{
				if len(uuid) > 0 {
					var targetEntity *BPMNDI_BPMNPlaneEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						targetEntity = instance.(*BPMNDI_BPMNPlaneEntity)
					}else{
						targetEntity = GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(uuid, nil)
						targetEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(targetEntity)
					}
					targetEntity.AppendReferenced("target", this)
					this.AppendChild("target",targetEntity)
				}
			} else if typeName == "BPMNDI.BPMNLabel"{
				if len(uuid) > 0 {
					var targetEntity *BPMNDI_BPMNLabelEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						targetEntity = instance.(*BPMNDI_BPMNLabelEntity)
					}else{
						targetEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuid, nil)
						targetEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(targetEntity)
					}
					targetEntity.AppendReferenced("target", this)
					this.AppendChild("target",targetEntity)
				}
			} else if typeName == "BPMNDI.BPMNEdge"{
				if len(uuid) > 0 {
					var targetEntity *BPMNDI_BPMNEdgeEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						targetEntity = instance.(*BPMNDI_BPMNEdgeEntity)
					}else{
						targetEntity = GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(uuid, nil)
						targetEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(targetEntity)
					}
					targetEntity.AppendReferenced("target", this)
					this.AppendChild("target",targetEntity)
				}
			}
			}
 		}

		/** waypoint **/
 		if results[0][10] != nil{
			uuidsStr :=results[0][10].(string)
			uuids :=make([]string,0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i:=0; i<len(uuids); i++{
				if len(uuids[i]) > 0 {
					var waypointEntity *DC_PointEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
						waypointEntity = instance.(*DC_PointEntity)
					}else{
						waypointEntity = GetServer().GetEntityManager().NewDCPointEntity(uuids[i], nil)
						waypointEntity.InitEntity(uuids[i])
						GetServer().GetEntityManager().insert(waypointEntity)
					}
					waypointEntity.AppendReferenced("waypoint", this)
					this.AppendChild("waypoint",waypointEntity)
				}
 			}
 		}

		/** members of LabeledEdge **/

		/** ownedLabel **/
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
				if typeName == "BPMNDI.BPMNLabel"{
						if len(uuids[i]) > 0 {
							var ownedLabelEntity *BPMNDI_BPMNLabelEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedLabelEntity = instance.(*BPMNDI_BPMNLabelEntity)
							}else{
								ownedLabelEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuids[i], nil)
								ownedLabelEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedLabelEntity)
							}
							ownedLabelEntity.AppendReferenced("ownedLabel", this)
							this.AppendChild("ownedLabel",ownedLabelEntity)
						}
				}
 			}
 		}

		/** members of BPMNEdge **/

		/** BPMNLabel **/
 		if results[0][12] != nil{
			uuid :=results[0][12].(string)
			if len(uuid) > 0 {
				var BPMNLabelEntity *BPMNDI_BPMNLabelEntity
				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
					BPMNLabelEntity = instance.(*BPMNDI_BPMNLabelEntity)
				}else{
					BPMNLabelEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuid, nil)
					BPMNLabelEntity.InitEntity(uuid)
					GetServer().GetEntityManager().insert( BPMNLabelEntity)
				}
				BPMNLabelEntity.AppendReferenced("BPMNLabel", this)
				this.AppendChild("BPMNLabel",BPMNLabelEntity)
			}
 		}

		/** bpmnElement **/
 		if results[0][13] != nil{
			id :=results[0][13].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.interface{}"
				id_:= refTypeName + "$$" + id
				this.object.M_bpmnElement= id
				GetServer().GetEntityManager().appendReference("bpmnElement",this.object.UUID, id_)
				this.object.M_bpmnElement = id
			}
 		}

		/** sourceElement **/
 		if results[0][14] != nil{
			uuid :=results[0][14].(string)
			if len(uuid) > 0 {
				typeName := uuid[0:strings.Index(uuid, "%")]
				if err!=nil{
					log.Println("type ", typeName, " not found!")
					return err
				}
			if typeName == "BPMNDI.BPMNShape"{
				if len(uuid) > 0 {
					var sourceElementEntity *BPMNDI_BPMNShapeEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						sourceElementEntity = instance.(*BPMNDI_BPMNShapeEntity)
					}else{
						sourceElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(uuid, nil)
						sourceElementEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(sourceElementEntity)
					}
					sourceElementEntity.AppendReferenced("sourceElement", this)
					this.AppendChild("sourceElement",sourceElementEntity)
				}
			} else if typeName == "BPMNDI.BPMNPlane"{
				if len(uuid) > 0 {
					var sourceElementEntity *BPMNDI_BPMNPlaneEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						sourceElementEntity = instance.(*BPMNDI_BPMNPlaneEntity)
					}else{
						sourceElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(uuid, nil)
						sourceElementEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(sourceElementEntity)
					}
					sourceElementEntity.AppendReferenced("sourceElement", this)
					this.AppendChild("sourceElement",sourceElementEntity)
				}
			} else if typeName == "BPMNDI.BPMNLabel"{
				if len(uuid) > 0 {
					var sourceElementEntity *BPMNDI_BPMNLabelEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						sourceElementEntity = instance.(*BPMNDI_BPMNLabelEntity)
					}else{
						sourceElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuid, nil)
						sourceElementEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(sourceElementEntity)
					}
					sourceElementEntity.AppendReferenced("sourceElement", this)
					this.AppendChild("sourceElement",sourceElementEntity)
				}
			} else if typeName == "BPMNDI.BPMNEdge"{
				if len(uuid) > 0 {
					var sourceElementEntity *BPMNDI_BPMNEdgeEntity
					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
						sourceElementEntity = instance.(*BPMNDI_BPMNEdgeEntity)
					}else{
						sourceElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(uuid, nil)
						sourceElementEntity.InitEntity(uuid)
						GetServer().GetEntityManager().insert(sourceElementEntity)
					}
					sourceElementEntity.AppendReferenced("sourceElement", this)
					this.AppendChild("sourceElement",sourceElementEntity)
				}
			}
			}
 		}

		/** targetElement **/
 		if results[0][15] != nil{
			id :=results[0][15].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.DiagramElement"
				id_:= refTypeName + "$$" + id
				this.object.M_targetElement= id
				GetServer().GetEntityManager().appendReference("targetElement",this.object.UUID, id_)
			}
 		}

		/** messageVisibleKind **/
 		if results[0][16] != nil{
 			enumIndex := results[0][16].(int)
			if enumIndex == 0{
 				this.object.M_messageVisibleKind=BPMNDI.MessageVisibleKind_Initiating
			} else if enumIndex == 1{
 				this.object.M_messageVisibleKind=BPMNDI.MessageVisibleKind_Non_initiating
 			}
 		}

		/** associations of BPMNEdge **/

		/** sourceEdgePtr **/
 		if results[0][17] != nil{
			idsStr :=results[0][17].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMNDI.Edge"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_sourceEdgePtr = append(this.object.M_sourceEdgePtr,ids[i])
					GetServer().GetEntityManager().appendReference("sourceEdgePtr",this.object.UUID, id_)
				}
			}
 		}

		/** targetEdgePtr **/
 		if results[0][18] != nil{
			idsStr :=results[0][18].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMNDI.Edge"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_targetEdgePtr = append(this.object.M_targetEdgePtr,ids[i])
					GetServer().GetEntityManager().appendReference("targetEdgePtr",this.object.UUID, id_)
				}
			}
 		}

		/** planePtr **/
 		if results[0][19] != nil{
			id :=results[0][19].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Plane"
				id_:= refTypeName + "$$" + id
				this.object.M_planePtr= id
				GetServer().GetEntityManager().appendReference("planePtr",this.object.UUID, id_)
			}
 		}
		childsUuidStr := results[0][20].(string)
		this.childsUuid = make([]string, 0)
		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
		if err != nil {
			return err
		}

		referencedStr := results[0][21].(string)
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
func (this *EntityManager) NewBPMNDIBPMNEdgeEntityFromObject(object *BPMNDI.BPMNEdge) *BPMNDI_BPMNEdgeEntity {
	 return this.NewBPMNDIBPMNEdgeEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMNDI_BPMNEdgeEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMNDIBPMNEdgeExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNEdge"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMNDI_BPMNEdgeEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMNDI_BPMNEdgeEntity) AppendReference(reference Entity) {

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
//              			BPMNLabel
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMNDI_BPMNLabelEntity struct{
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
	object *BPMNDI.BPMNLabel
}

/** Constructor function **/
func (this *EntityManager) NewBPMNDIBPMNLabelEntity(objectId string, object interface{}) *BPMNDI_BPMNLabelEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMNDIBPMNLabelExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMNDI.BPMNLabel).TYPENAME = "BPMNDI.BPMNLabel"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMNDI.BPMNLabel).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMNDI_BPMNLabelEntity)
		}
	}else{
		uuidStr = "BPMNDI.BPMNLabel%" + Utility.RandomUUID()
	}
	entity := new(BPMNDI_BPMNLabelEntity)
	if object == nil{
		entity.object = new(BPMNDI.BPMNLabel)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMNDI.BPMNLabel)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMNDI.BPMNLabel"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMNDI.BPMNLabel","BPMN20")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMNDI_BPMNLabelEntity) GetTypeName()string{
	return "BPMNDI.BPMNLabel"
}
func(this *BPMNDI_BPMNLabelEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMNDI_BPMNLabelEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMNDI_BPMNLabelEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMNDI_BPMNLabelEntity) AppendReferenced(name string, owner Entity){
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

func(this *BPMNDI_BPMNLabelEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMNDI_BPMNLabelEntity) RemoveReferenced(name string, owner Entity) {
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

func(this *BPMNDI_BPMNLabelEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMNDI_BPMNLabelEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMNDI_BPMNLabelEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMNDI_BPMNLabelEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMNDI_BPMNLabelEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMNDI_BPMNLabelEntity) RemoveChild(name string, uuid string) {
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

func(this *BPMNDI_BPMNLabelEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMNDI_BPMNLabelEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMNDI_BPMNLabelEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMNDI_BPMNLabelEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMNDI_BPMNLabelEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMNDI_BPMNLabelEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMNDI_BPMNLabelEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMNDI_BPMNLabelEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMNDI_BPMNLabelEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMNDI_BPMNLabelEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMNDI_BPMNLabelEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNLabel"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMNDI_BPMNLabelEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) create_BPMNDI_BPMNLabelEntityPrototype() {

	var bPMNLabelEntityProto EntityPrototype
	bPMNLabelEntityProto.TypeName = "BPMNDI.BPMNLabel"
	bPMNLabelEntityProto.SuperTypeNames = append(bPMNLabelEntityProto.SuperTypeNames, "DI.DiagramElement")
	bPMNLabelEntityProto.SuperTypeNames = append(bPMNLabelEntityProto.SuperTypeNames, "DI.Node")
	bPMNLabelEntityProto.SuperTypeNames = append(bPMNLabelEntityProto.SuperTypeNames, "DI.Label")
	bPMNLabelEntityProto.Ids = append(bPMNLabelEntityProto.Ids,"uuid")
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"uuid")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"xs.string")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,0)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)
	bPMNLabelEntityProto.Indexs = append(bPMNLabelEntityProto.Indexs,"parentUuid")
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"parentUuid")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"xs.string")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,1)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)

	/** members of DiagramElement **/
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,2)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,true)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_owningDiagram")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"BPMNDI.DI.Diagram:Ref")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,3)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,true)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_owningElement")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"BPMNDI.DI.DiagramElement:Ref")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,4)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,true)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_modelElement")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"xs.interface{}:Ref")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,5)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,true)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_style")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"BPMNDI.DI.Style:Ref")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,6)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,true)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_ownedElement")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"[]BPMNDI.DI.DiagramElement")
	bPMNLabelEntityProto.Ids = append(bPMNLabelEntityProto.Ids,"M_id")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,7)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,true)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_id")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"xs.ID")

	/** members of Node **/
	/** No members **/

	/** members of Label **/
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,8)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,true)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_Bounds")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"BPMNDI.DC.Bounds")

	/** members of BPMNLabel **/
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,9)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,true)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_labelStyle")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"BPMNDI.BPMNLabelStyle:Ref")

	/** associations of BPMNLabel **/
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,10)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_shapePtr")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"BPMNDI.BPMNShape:Ref")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,11)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_edgePtr")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"BPMNDI.BPMNEdge:Ref")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,12)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_sourceEdgePtr")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"[]BPMNDI.DI.Edge:Ref")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,13)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_targetEdgePtr")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"[]BPMNDI.DI.Edge:Ref")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,14)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_planePtr")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"BPMNDI.DI.Plane:Ref")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,15)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_owningEdgePtr")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"BPMNDI.DI.LabeledEdge:Ref")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,16)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"M_owningShapePtr")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"BPMNDI.DI.LabeledShape:Ref")
	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"childsUuid")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"[]xs.string")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,17)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)

	bPMNLabelEntityProto.Fields = append(bPMNLabelEntityProto.Fields,"referenced")
	bPMNLabelEntityProto.FieldsType = append(bPMNLabelEntityProto.FieldsType,"[]EntityRef")
	bPMNLabelEntityProto.FieldsOrder = append(bPMNLabelEntityProto.FieldsOrder,18)
	bPMNLabelEntityProto.FieldsVisibility = append(bPMNLabelEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMNDIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&bPMNLabelEntityProto)

}

/** Create **/
func (this *BPMNDI_BPMNLabelEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMNDI.BPMNLabel"

	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNLabel"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of DiagramElement **/
	query.Fields = append(query.Fields, "M_owningDiagram")
	query.Fields = append(query.Fields, "M_owningElement")
	query.Fields = append(query.Fields, "M_modelElement")
	query.Fields = append(query.Fields, "M_style")
	query.Fields = append(query.Fields, "M_ownedElement")
	query.Fields = append(query.Fields, "M_id")

	/** members of Node **/
	/** No members **/

	/** members of Label **/
	query.Fields = append(query.Fields, "M_Bounds")

	/** members of BPMNLabel **/
	query.Fields = append(query.Fields, "M_labelStyle")

		/** associations of BPMNLabel **/
	query.Fields = append(query.Fields, "M_shapePtr")
	query.Fields = append(query.Fields, "M_edgePtr")
	query.Fields = append(query.Fields, "M_sourceEdgePtr")
	query.Fields = append(query.Fields, "M_targetEdgePtr")
	query.Fields = append(query.Fields, "M_planePtr")
	query.Fields = append(query.Fields, "M_owningEdgePtr")
	query.Fields = append(query.Fields, "M_owningShapePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var BPMNLabelInfo []interface{}

	BPMNLabelInfo = append(BPMNLabelInfo, this.GetUuid())
	if this.parentPtr != nil {
		BPMNLabelInfo = append(BPMNLabelInfo, this.parentPtr.GetUuid())
	}else{
		BPMNLabelInfo = append(BPMNLabelInfo, "")
	}

	/** members of DiagramElement **/

	/** Save owningDiagram type Diagram **/
		BPMNLabelInfo = append(BPMNLabelInfo,this.object.M_owningDiagram)

	/** Save owningElement type DiagramElement **/
		BPMNLabelInfo = append(BPMNLabelInfo,this.object.M_owningElement)
	BPMNLabelInfo = append(BPMNLabelInfo, this.object.M_modelElement)

	/** Save style type Style **/
		BPMNLabelInfo = append(BPMNLabelInfo,this.object.M_style)

	/** Save ownedElement type DiagramElement **/
	ownedElementIds := make([]string,0)
	for i := 0; i < len(this.object.M_ownedElement); i++ {
		switch v := this.object.M_ownedElement[i].(type) {
		case *BPMNDI.BPMNPlane:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNLabel:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNShape:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		case *BPMNDI.BPMNEdge:
		ownedElementEntity:= GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(v.UUID, v)
		ownedElementIds=append(ownedElementIds,ownedElementEntity.uuid)
		ownedElementEntity.AppendReferenced("ownedElement", this)
		this.AppendChild("ownedElement",ownedElementEntity)
		if ownedElementEntity.NeedSave() {
			ownedElementEntity.SaveEntity()
		}
		}
	}
	ownedElementStr, _ := json.Marshal(ownedElementIds)
	BPMNLabelInfo = append(BPMNLabelInfo, string(ownedElementStr))
	BPMNLabelInfo = append(BPMNLabelInfo, this.object.M_id)

	/** members of Node **/
	/** No members **/

	/** members of Label **/

	/** Save Bounds type Bounds **/
	if this.object.M_Bounds != nil {
		BoundsEntity:= GetServer().GetEntityManager().NewDCBoundsEntity(this.object.M_Bounds.UUID, this.object.M_Bounds)
		BPMNLabelInfo = append(BPMNLabelInfo, BoundsEntity.uuid)
		BoundsEntity.AppendReferenced("Bounds", this)
		this.AppendChild("Bounds",BoundsEntity)
		if BoundsEntity.NeedSave() {
			BoundsEntity.SaveEntity()
		}
	}else{
		BPMNLabelInfo = append(BPMNLabelInfo, "")
	}

	/** members of BPMNLabel **/

	/** Save labelStyle type BPMNLabelStyle **/
		BPMNLabelInfo = append(BPMNLabelInfo,this.object.M_labelStyle)

	/** associations of BPMNLabel **/

	/** Save shape type BPMNShape **/
		BPMNLabelInfo = append(BPMNLabelInfo,this.object.M_shapePtr)

	/** Save edge type BPMNEdge **/
		BPMNLabelInfo = append(BPMNLabelInfo,this.object.M_edgePtr)

	/** Save sourceEdge type Edge **/
	sourceEdgePtrStr, _ := json.Marshal(this.object.M_sourceEdgePtr)
	BPMNLabelInfo = append(BPMNLabelInfo, string(sourceEdgePtrStr))

	/** Save targetEdge type Edge **/
	targetEdgePtrStr, _ := json.Marshal(this.object.M_targetEdgePtr)
	BPMNLabelInfo = append(BPMNLabelInfo, string(targetEdgePtrStr))

	/** Save plane type Plane **/
		BPMNLabelInfo = append(BPMNLabelInfo,this.object.M_planePtr)

	/** Save owningEdge type LabeledEdge **/
		BPMNLabelInfo = append(BPMNLabelInfo,this.object.M_owningEdgePtr)

	/** Save owningShape type LabeledShape **/
		BPMNLabelInfo = append(BPMNLabelInfo,this.object.M_owningShapePtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	BPMNLabelInfo = append(BPMNLabelInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	BPMNLabelInfo = append(BPMNLabelInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMNDIDB, string(queryStr), BPMNLabelInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMNDIDB, string(queryStr), BPMNLabelInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMNDI_BPMNLabelEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMNDI_BPMNLabelEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNLabel"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of DiagramElement **/
	query.Fields = append(query.Fields, "M_owningDiagram")
	query.Fields = append(query.Fields, "M_owningElement")
	query.Fields = append(query.Fields, "M_modelElement")
	query.Fields = append(query.Fields, "M_style")
	query.Fields = append(query.Fields, "M_ownedElement")
	query.Fields = append(query.Fields, "M_id")

	/** members of Node **/
	/** No members **/

	/** members of Label **/
	query.Fields = append(query.Fields, "M_Bounds")

	/** members of BPMNLabel **/
	query.Fields = append(query.Fields, "M_labelStyle")

		/** associations of BPMNLabel **/
	query.Fields = append(query.Fields, "M_shapePtr")
	query.Fields = append(query.Fields, "M_edgePtr")
	query.Fields = append(query.Fields, "M_sourceEdgePtr")
	query.Fields = append(query.Fields, "M_targetEdgePtr")
	query.Fields = append(query.Fields, "M_planePtr")
	query.Fields = append(query.Fields, "M_owningEdgePtr")
	query.Fields = append(query.Fields, "M_owningShapePtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of BPMNLabel...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMNDI.BPMNLabel)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMNDI.BPMNLabel"

		this.parentUuid = results[0][1].(string)

		/** members of DiagramElement **/

		/** owningDiagram **/
 		if results[0][2] != nil{
			id :=results[0][2].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Diagram"
				id_:= refTypeName + "$$" + id
				this.object.M_owningDiagram= id
				GetServer().GetEntityManager().appendReference("owningDiagram",this.object.UUID, id_)
			}
 		}

		/** owningElement **/
 		if results[0][3] != nil{
			id :=results[0][3].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.DiagramElement"
				id_:= refTypeName + "$$" + id
				this.object.M_owningElement= id
				GetServer().GetEntityManager().appendReference("owningElement",this.object.UUID, id_)
			}
 		}

		/** modelElement **/
 		if results[0][4] != nil{
			id :=results[0][4].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.interface{}"
				id_:= refTypeName + "$$" + id
				this.object.M_modelElement= id
				GetServer().GetEntityManager().appendReference("modelElement",this.object.UUID, id_)
			}
 		}

		/** style **/
 		if results[0][5] != nil{
			id :=results[0][5].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Style"
				id_:= refTypeName + "$$" + id
				this.object.M_style= id
				GetServer().GetEntityManager().appendReference("style",this.object.UUID, id_)
			}
 		}

		/** ownedElement **/
 		if results[0][6] != nil{
			uuidsStr :=results[0][6].(string)
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
				if typeName == "BPMNDI.BPMNEdge"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNEdgeEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNEdgeEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNEdgeEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNPlane"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNPlaneEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNPlaneEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNPlaneEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNLabel"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNLabelEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNLabelEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNLabelEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				} else if typeName == "BPMNDI.BPMNShape"{
						if len(uuids[i]) > 0 {
							var ownedElementEntity *BPMNDI_BPMNShapeEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								ownedElementEntity = instance.(*BPMNDI_BPMNShapeEntity)
							}else{
								ownedElementEntity = GetServer().GetEntityManager().NewBPMNDIBPMNShapeEntity(uuids[i], nil)
								ownedElementEntity.InitEntity(uuids[i])
								GetServer().GetEntityManager().insert(ownedElementEntity)
							}
							ownedElementEntity.AppendReferenced("ownedElement", this)
							this.AppendChild("ownedElement",ownedElementEntity)
						}
				}
 			}
 		}

		/** id **/
 		if results[0][7] != nil{
 			this.object.M_id=results[0][7].(string)
 		}

		/** members of Node **/
		/** No members **/

		/** members of Label **/

		/** Bounds **/
 		if results[0][8] != nil{
			uuid :=results[0][8].(string)
			if len(uuid) > 0 {
				var BoundsEntity *DC_BoundsEntity
				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
					BoundsEntity = instance.(*DC_BoundsEntity)
				}else{
					BoundsEntity = GetServer().GetEntityManager().NewDCBoundsEntity(uuid, nil)
					BoundsEntity.InitEntity(uuid)
					GetServer().GetEntityManager().insert( BoundsEntity)
				}
				BoundsEntity.AppendReferenced("Bounds", this)
				this.AppendChild("Bounds",BoundsEntity)
			}
 		}

		/** members of BPMNLabel **/

		/** labelStyle **/
 		if results[0][9] != nil{
			id :=results[0][9].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.BPMNLabelStyle"
				id_:= refTypeName + "$$" + id
				this.object.M_labelStyle= id
				GetServer().GetEntityManager().appendReference("labelStyle",this.object.UUID, id_)
			}
 		}

		/** associations of BPMNLabel **/

		/** shapePtr **/
 		if results[0][10] != nil{
			id :=results[0][10].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.BPMNShape"
				id_:= refTypeName + "$$" + id
				this.object.M_shapePtr= id
				GetServer().GetEntityManager().appendReference("shapePtr",this.object.UUID, id_)
			}
 		}

		/** edgePtr **/
 		if results[0][11] != nil{
			id :=results[0][11].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.BPMNEdge"
				id_:= refTypeName + "$$" + id
				this.object.M_edgePtr= id
				GetServer().GetEntityManager().appendReference("edgePtr",this.object.UUID, id_)
			}
 		}

		/** sourceEdgePtr **/
 		if results[0][12] != nil{
			idsStr :=results[0][12].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMNDI.Edge"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_sourceEdgePtr = append(this.object.M_sourceEdgePtr,ids[i])
					GetServer().GetEntityManager().appendReference("sourceEdgePtr",this.object.UUID, id_)
				}
			}
 		}

		/** targetEdgePtr **/
 		if results[0][13] != nil{
			idsStr :=results[0][13].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMNDI.Edge"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_targetEdgePtr = append(this.object.M_targetEdgePtr,ids[i])
					GetServer().GetEntityManager().appendReference("targetEdgePtr",this.object.UUID, id_)
				}
			}
 		}

		/** planePtr **/
 		if results[0][14] != nil{
			id :=results[0][14].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Plane"
				id_:= refTypeName + "$$" + id
				this.object.M_planePtr= id
				GetServer().GetEntityManager().appendReference("planePtr",this.object.UUID, id_)
			}
 		}

		/** owningEdgePtr **/
 		if results[0][15] != nil{
			id :=results[0][15].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.LabeledEdge"
				id_:= refTypeName + "$$" + id
				this.object.M_owningEdgePtr= id
				GetServer().GetEntityManager().appendReference("owningEdgePtr",this.object.UUID, id_)
			}
 		}

		/** owningShapePtr **/
 		if results[0][16] != nil{
			id :=results[0][16].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.LabeledShape"
				id_:= refTypeName + "$$" + id
				this.object.M_owningShapePtr= id
				GetServer().GetEntityManager().appendReference("owningShapePtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMNDIBPMNLabelEntityFromObject(object *BPMNDI.BPMNLabel) *BPMNDI_BPMNLabelEntity {
	 return this.NewBPMNDIBPMNLabelEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMNDI_BPMNLabelEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMNDIBPMNLabelExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNLabel"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMNDI_BPMNLabelEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMNDI_BPMNLabelEntity) AppendReference(reference Entity) {

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
//              			BPMNLabelStyle
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type BPMNDI_BPMNLabelStyleEntity struct{
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
	object *BPMNDI.BPMNLabelStyle
}

/** Constructor function **/
func (this *EntityManager) NewBPMNDIBPMNLabelStyleEntity(objectId string, object interface{}) *BPMNDI_BPMNLabelStyleEntity{
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId){
			uuidStr = objectId
		}else{
			uuidStr  = BPMNDIBPMNLabelStyleExists(objectId)
		}
	}
	if object != nil{
		object.(*BPMNDI.BPMNLabelStyle).TYPENAME = "BPMNDI.BPMNLabelStyle"
	}
	if len(uuidStr) > 0 {
		if object != nil{
			object.(*BPMNDI.BPMNLabelStyle).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr);ok {
			if object != nil{
				this.setObjectValues(val, object)

			}
			return val.(*BPMNDI_BPMNLabelStyleEntity)
		}
	}else{
		uuidStr = "BPMNDI.BPMNLabelStyle%" + Utility.RandomUUID()
	}
	entity := new(BPMNDI_BPMNLabelStyleEntity)
	if object == nil{
		entity.object = new(BPMNDI.BPMNLabelStyle)
		entity.SetNeedSave(true)
	}else{
		entity.object = object.(*BPMNDI.BPMNLabelStyle)
		entity.SetNeedSave(true)
	}
	entity.object.TYPENAME = "BPMNDI.BPMNLabelStyle"

	entity.object.UUID = uuidStr
	entity.SetInit(false)
	entity.uuid = uuidStr
	this.insert(entity)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("BPMNDI.BPMNLabelStyle","BPMN20")
	entity.prototype = prototype
	return entity
}

/** Entity functions **/
func(this *BPMNDI_BPMNLabelStyleEntity) GetTypeName()string{
	return "BPMNDI.BPMNLabelStyle"
}
func(this *BPMNDI_BPMNLabelStyleEntity) GetUuid()string{
	return this.uuid
}
func(this *BPMNDI_BPMNLabelStyleEntity) GetParentPtr()Entity{
	return this.parentPtr
}

func(this *BPMNDI_BPMNLabelStyleEntity) SetParentPtr(parentPtr Entity){
	this.parentPtr=parentPtr
}

func(this *BPMNDI_BPMNLabelStyleEntity) AppendReferenced(name string, owner Entity){
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

func(this *BPMNDI_BPMNLabelStyleEntity) GetReferenced() []EntityRef{
	return this.referenced
}

func(this *BPMNDI_BPMNLabelStyleEntity) RemoveReferenced(name string, owner Entity) {
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

func(this *BPMNDI_BPMNLabelStyleEntity) RemoveReference(name string, reference Entity){
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

func(this *BPMNDI_BPMNLabelStyleEntity) GetChildsPtr() []Entity{
	return this.childsPtr
}

func(this *BPMNDI_BPMNLabelStyleEntity) SetChildsPtr(childsPtr[]Entity){
	this.childsPtr = childsPtr
}

func(this *BPMNDI_BPMNLabelStyleEntity) GetChildsUuid() []string{
	return this.childsUuid
}

func(this *BPMNDI_BPMNLabelStyleEntity) SetChildsUuid(childsUuid[]string){
	this.childsUuid = childsUuid
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func(this *BPMNDI_BPMNLabelStyleEntity) RemoveChild(name string, uuid string) {
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

func(this *BPMNDI_BPMNLabelStyleEntity) GetReferencesUuid() []string{
	return this.referencesUuid
}

func(this *BPMNDI_BPMNLabelStyleEntity) SetReferencesUuid(refsUuid[]string){
	this.referencesUuid = refsUuid
}

func(this *BPMNDI_BPMNLabelStyleEntity) GetReferencesPtr() []Entity{
	return this.referencesPtr
}

func(this *BPMNDI_BPMNLabelStyleEntity) SetReferencesPtr(refsPtr[]Entity){
	this.referencesPtr = refsPtr
}

func(this *BPMNDI_BPMNLabelStyleEntity) GetObject() interface{}{
	return this.object
}

func(this *BPMNDI_BPMNLabelStyleEntity) NeedSave() bool{
	return this.object.NeedSave
}

func(this *BPMNDI_BPMNLabelStyleEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func(this *BPMNDI_BPMNLabelStyleEntity) IsInit() bool{
	return this.object.IsInit
}

func(this *BPMNDI_BPMNLabelStyleEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func(this *BPMNDI_BPMNLabelStyleEntity) GetChecksum() string{
	objectStr, _ := json.Marshal(this.object)
	return  Utility.GetMD5Hash(string(objectStr))
}

func(this *BPMNDI_BPMNLabelStyleEntity) Exist() bool{
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNLabelStyle"
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
*/
func(this *BPMNDI_BPMNLabelStyleEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}
/** Entity Prototype creation **/
func (this *EntityManager) create_BPMNDI_BPMNLabelStyleEntityPrototype() {

	var bPMNLabelStyleEntityProto EntityPrototype
	bPMNLabelStyleEntityProto.TypeName = "BPMNDI.BPMNLabelStyle"
	bPMNLabelStyleEntityProto.SuperTypeNames = append(bPMNLabelStyleEntityProto.SuperTypeNames, "DI.Style")
	bPMNLabelStyleEntityProto.Ids = append(bPMNLabelStyleEntityProto.Ids,"uuid")
	bPMNLabelStyleEntityProto.Fields = append(bPMNLabelStyleEntityProto.Fields,"uuid")
	bPMNLabelStyleEntityProto.FieldsType = append(bPMNLabelStyleEntityProto.FieldsType,"xs.string")
	bPMNLabelStyleEntityProto.FieldsOrder = append(bPMNLabelStyleEntityProto.FieldsOrder,0)
	bPMNLabelStyleEntityProto.FieldsVisibility = append(bPMNLabelStyleEntityProto.FieldsVisibility,false)
	bPMNLabelStyleEntityProto.Indexs = append(bPMNLabelStyleEntityProto.Indexs,"parentUuid")
	bPMNLabelStyleEntityProto.Fields = append(bPMNLabelStyleEntityProto.Fields,"parentUuid")
	bPMNLabelStyleEntityProto.FieldsType = append(bPMNLabelStyleEntityProto.FieldsType,"xs.string")
	bPMNLabelStyleEntityProto.FieldsOrder = append(bPMNLabelStyleEntityProto.FieldsOrder,1)
	bPMNLabelStyleEntityProto.FieldsVisibility = append(bPMNLabelStyleEntityProto.FieldsVisibility,false)

	/** members of Style **/
	bPMNLabelStyleEntityProto.Ids = append(bPMNLabelStyleEntityProto.Ids,"M_id")
	bPMNLabelStyleEntityProto.FieldsOrder = append(bPMNLabelStyleEntityProto.FieldsOrder,2)
	bPMNLabelStyleEntityProto.FieldsVisibility = append(bPMNLabelStyleEntityProto.FieldsVisibility,true)
	bPMNLabelStyleEntityProto.Fields = append(bPMNLabelStyleEntityProto.Fields,"M_id")
	bPMNLabelStyleEntityProto.FieldsType = append(bPMNLabelStyleEntityProto.FieldsType,"xs.ID")

	/** members of BPMNLabelStyle **/
	bPMNLabelStyleEntityProto.FieldsOrder = append(bPMNLabelStyleEntityProto.FieldsOrder,3)
	bPMNLabelStyleEntityProto.FieldsVisibility = append(bPMNLabelStyleEntityProto.FieldsVisibility,true)
	bPMNLabelStyleEntityProto.Fields = append(bPMNLabelStyleEntityProto.Fields,"M_Font")
	bPMNLabelStyleEntityProto.FieldsType = append(bPMNLabelStyleEntityProto.FieldsType,"BPMNDI.DC.Font")

	/** associations of BPMNLabelStyle **/
	bPMNLabelStyleEntityProto.FieldsOrder = append(bPMNLabelStyleEntityProto.FieldsOrder,4)
	bPMNLabelStyleEntityProto.FieldsVisibility = append(bPMNLabelStyleEntityProto.FieldsVisibility,false)
	bPMNLabelStyleEntityProto.Fields = append(bPMNLabelStyleEntityProto.Fields,"M_diagramPtr")
	bPMNLabelStyleEntityProto.FieldsType = append(bPMNLabelStyleEntityProto.FieldsType,"BPMNDI.BPMNDiagram:Ref")
	bPMNLabelStyleEntityProto.FieldsOrder = append(bPMNLabelStyleEntityProto.FieldsOrder,5)
	bPMNLabelStyleEntityProto.FieldsVisibility = append(bPMNLabelStyleEntityProto.FieldsVisibility,false)
	bPMNLabelStyleEntityProto.Fields = append(bPMNLabelStyleEntityProto.Fields,"M_labelPtr")
	bPMNLabelStyleEntityProto.FieldsType = append(bPMNLabelStyleEntityProto.FieldsType,"[]BPMNDI.BPMNLabel:Ref")
	bPMNLabelStyleEntityProto.FieldsOrder = append(bPMNLabelStyleEntityProto.FieldsOrder,6)
	bPMNLabelStyleEntityProto.FieldsVisibility = append(bPMNLabelStyleEntityProto.FieldsVisibility,false)
	bPMNLabelStyleEntityProto.Fields = append(bPMNLabelStyleEntityProto.Fields,"M_diagramElementPtr")
	bPMNLabelStyleEntityProto.FieldsType = append(bPMNLabelStyleEntityProto.FieldsType,"[]BPMNDI.DI.DiagramElement:Ref")
	bPMNLabelStyleEntityProto.FieldsOrder = append(bPMNLabelStyleEntityProto.FieldsOrder,7)
	bPMNLabelStyleEntityProto.FieldsVisibility = append(bPMNLabelStyleEntityProto.FieldsVisibility,false)
	bPMNLabelStyleEntityProto.Fields = append(bPMNLabelStyleEntityProto.Fields,"M_owningDiagramPtr")
	bPMNLabelStyleEntityProto.FieldsType = append(bPMNLabelStyleEntityProto.FieldsType,"BPMNDI.DI.Diagram:Ref")
	bPMNLabelStyleEntityProto.Fields = append(bPMNLabelStyleEntityProto.Fields,"childsUuid")
	bPMNLabelStyleEntityProto.FieldsType = append(bPMNLabelStyleEntityProto.FieldsType,"[]xs.string")
	bPMNLabelStyleEntityProto.FieldsOrder = append(bPMNLabelStyleEntityProto.FieldsOrder,8)
	bPMNLabelStyleEntityProto.FieldsVisibility = append(bPMNLabelStyleEntityProto.FieldsVisibility,false)

	bPMNLabelStyleEntityProto.Fields = append(bPMNLabelStyleEntityProto.Fields,"referenced")
	bPMNLabelStyleEntityProto.FieldsType = append(bPMNLabelStyleEntityProto.FieldsType,"[]EntityRef")
	bPMNLabelStyleEntityProto.FieldsOrder = append(bPMNLabelStyleEntityProto.FieldsOrder,9)
	bPMNLabelStyleEntityProto.FieldsVisibility = append(bPMNLabelStyleEntityProto.FieldsVisibility,false)

	store := GetServer().GetDataManager().getDataStore(BPMNDIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&bPMNLabelStyleEntityProto)

}

/** Create **/
func (this *BPMNDI_BPMNLabelStyleEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	this.object.UUID = this.uuid
	this.object.TYPENAME = "BPMNDI.BPMNLabelStyle"

	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNLabelStyle"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Style **/
	query.Fields = append(query.Fields, "M_id")

	/** members of BPMNLabelStyle **/
	query.Fields = append(query.Fields, "M_Font")

		/** associations of BPMNLabelStyle **/
	query.Fields = append(query.Fields, "M_diagramPtr")
	query.Fields = append(query.Fields, "M_labelPtr")
	query.Fields = append(query.Fields, "M_diagramElementPtr")
	query.Fields = append(query.Fields, "M_owningDiagramPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	var BPMNLabelStyleInfo []interface{}

	BPMNLabelStyleInfo = append(BPMNLabelStyleInfo, this.GetUuid())
	if this.parentPtr != nil {
		BPMNLabelStyleInfo = append(BPMNLabelStyleInfo, this.parentPtr.GetUuid())
	}else{
		BPMNLabelStyleInfo = append(BPMNLabelStyleInfo, "")
	}

	/** members of Style **/
	BPMNLabelStyleInfo = append(BPMNLabelStyleInfo, this.object.M_id)

	/** members of BPMNLabelStyle **/

	/** Save Font type Font **/
	if this.object.M_Font != nil {
		FontEntity:= GetServer().GetEntityManager().NewDCFontEntity(this.object.M_Font.UUID, this.object.M_Font)
		BPMNLabelStyleInfo = append(BPMNLabelStyleInfo, FontEntity.uuid)
		FontEntity.AppendReferenced("Font", this)
		this.AppendChild("Font",FontEntity)
		if FontEntity.NeedSave() {
			FontEntity.SaveEntity()
		}
	}else{
		BPMNLabelStyleInfo = append(BPMNLabelStyleInfo, "")
	}

	/** associations of BPMNLabelStyle **/

	/** Save diagram type BPMNDiagram **/
		BPMNLabelStyleInfo = append(BPMNLabelStyleInfo,this.object.M_diagramPtr)

	/** Save label type BPMNLabel **/
	labelPtrStr, _ := json.Marshal(this.object.M_labelPtr)
	BPMNLabelStyleInfo = append(BPMNLabelStyleInfo, string(labelPtrStr))

	/** Save diagramElement type DiagramElement **/
	diagramElementPtrStr, _ := json.Marshal(this.object.M_diagramElementPtr)
	BPMNLabelStyleInfo = append(BPMNLabelStyleInfo, string(diagramElementPtrStr))

	/** Save owningDiagram type Diagram **/
		BPMNLabelStyleInfo = append(BPMNLabelStyleInfo,this.object.M_owningDiagramPtr)
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	BPMNLabelStyleInfo = append(BPMNLabelStyleInfo, string(childsUuidStr))
	referencedStr, _ := json.Marshal(this.referenced)
	BPMNLabelStyleInfo = append(BPMNLabelStyleInfo, string(referencedStr))
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
		err = GetServer().GetDataManager().updateData(BPMNDIDB, string(queryStr), BPMNLabelStyleInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err =  GetServer().GetDataManager().createData(BPMNDIDB, string(queryStr), BPMNLabelStyleInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *BPMNDI_BPMNLabelStyleEntity) InitEntity(id string) error{
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*BPMNDI_BPMNLabelStyleEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.uuid = id

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNLabelStyle"

	query.Fields = append(query.Fields, "uuid")
	query.Fields = append(query.Fields, "parentUuid")

	/** members of Style **/
	query.Fields = append(query.Fields, "M_id")

	/** members of BPMNLabelStyle **/
	query.Fields = append(query.Fields, "M_Font")

		/** associations of BPMNLabelStyle **/
	query.Fields = append(query.Fields, "M_diagramPtr")
	query.Fields = append(query.Fields, "M_labelPtr")
	query.Fields = append(query.Fields, "M_diagramElementPtr")
	query.Fields = append(query.Fields, "M_owningDiagramPtr")

	query.Fields = append(query.Fields, "childsUuid")
	query.Fields = append(query.Fields, "referenced")
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of BPMNLabelStyle...
	if len(results) > 0 {

	/** initialyzation of the entity object **/
		this.object = new(BPMNDI.BPMNLabelStyle)
		this.object.UUID = this.uuid
		this.object.TYPENAME = "BPMNDI.BPMNLabelStyle"

		this.parentUuid = results[0][1].(string)

		/** members of Style **/

		/** id **/
 		if results[0][2] != nil{
 			this.object.M_id=results[0][2].(string)
 		}

		/** members of BPMNLabelStyle **/

		/** Font **/
 		if results[0][3] != nil{
			uuid :=results[0][3].(string)
			if len(uuid) > 0 {
				var FontEntity *DC_FontEntity
				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
					FontEntity = instance.(*DC_FontEntity)
				}else{
					FontEntity = GetServer().GetEntityManager().NewDCFontEntity(uuid, nil)
					FontEntity.InitEntity(uuid)
					GetServer().GetEntityManager().insert( FontEntity)
				}
				FontEntity.AppendReferenced("Font", this)
				this.AppendChild("Font",FontEntity)
			}
 		}

		/** associations of BPMNLabelStyle **/

		/** diagramPtr **/
 		if results[0][4] != nil{
			id :=results[0][4].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.BPMNDiagram"
				id_:= refTypeName + "$$" + id
				this.object.M_diagramPtr= id
				GetServer().GetEntityManager().appendReference("diagramPtr",this.object.UUID, id_)
			}
 		}

		/** labelPtr **/
 		if results[0][5] != nil{
			idsStr :=results[0][5].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMNDI.BPMNLabel"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_labelPtr = append(this.object.M_labelPtr,ids[i])
					GetServer().GetEntityManager().appendReference("labelPtr",this.object.UUID, id_)
				}
			}
 		}

		/** diagramElementPtr **/
 		if results[0][6] != nil{
			idsStr :=results[0][6].(string)
			ids :=make([]string,0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i:=0; i<len(ids); i++{
				if len(ids[i]) > 0 {
					refTypeName:="BPMNDI.DiagramElement"
					id_:= refTypeName + "$$" + ids[i]
					this.object.M_diagramElementPtr = append(this.object.M_diagramElementPtr,ids[i])
					GetServer().GetEntityManager().appendReference("diagramElementPtr",this.object.UUID, id_)
				}
			}
 		}

		/** owningDiagramPtr **/
 		if results[0][7] != nil{
			id :=results[0][7].(string)
			if len(id) > 0 {
				refTypeName:="BPMNDI.Diagram"
				id_:= refTypeName + "$$" + id
				this.object.M_owningDiagramPtr= id
				GetServer().GetEntityManager().appendReference("owningDiagramPtr",this.object.UUID, id_)
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
func (this *EntityManager) NewBPMNDIBPMNLabelStyleEntityFromObject(object *BPMNDI.BPMNLabelStyle) *BPMNDI_BPMNLabelStyleEntity {
	 return this.NewBPMNDIBPMNLabelStyleEntity(object.UUID, object)
}

/** Delete **/
func (this *BPMNDI_BPMNLabelStyleEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func BPMNDIBPMNLabelStyleExists(val string) string {
	var query EntityQuery
	query.TypeName = "BPMNDI.BPMNLabelStyle"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface {} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(BPMNDIDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *BPMNDI_BPMNLabelStyleEntity) AppendChild(attributeName string, child Entity) error {

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
func (this *BPMNDI_BPMNLabelStyleEntity) AppendReference(reference Entity) {

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
func (this *EntityManager) registerBPMNDIObjects(){
	Utility.RegisterType((*BPMNDI.BPMNDiagram)(nil))
	Utility.RegisterType((*BPMNDI.BPMNPlane)(nil))
	Utility.RegisterType((*BPMNDI.BPMNShape)(nil))
	Utility.RegisterType((*BPMNDI.BPMNEdge)(nil))
	Utility.RegisterType((*BPMNDI.BPMNLabel)(nil))
	Utility.RegisterType((*BPMNDI.BPMNLabelStyle)(nil))
}

/** Create entity prototypes contain in a package **/
func (this *EntityManager) createBPMNDIPrototypes(){
	this.create_BPMNDI_BPMNDiagramEntityPrototype() 
	this.create_BPMNDI_BPMNPlaneEntityPrototype() 
	this.create_BPMNDI_BPMNShapeEntityPrototype() 
	this.create_BPMNDI_BPMNEdgeEntityPrototype() 
	this.create_BPMNDI_BPMNLabelEntityPrototype() 
	this.create_BPMNDI_BPMNLabelStyleEntityPrototype() 
}

