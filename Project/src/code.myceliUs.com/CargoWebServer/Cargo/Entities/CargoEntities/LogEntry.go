// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type LogEntry struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool

	/** members of Entity **/
	M_id string

	/** members of LogEntry **/
	M_creationTime int64
	m_entityRef Entity
	/** If the ref is a string and not an object **/
	M_entityRef string


	/** Associations **/
	m_loggerPtr *Log
	/** If the ref is a string and not an object **/
	M_loggerPtr string
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for LogEntry **/
type XsdLogEntry struct {
	XMLName xml.Name	`xml:"logEntry"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_entityRef	string	`xml:"entityRef"`
	M_creationTime	int64	`xml:"creationTime,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *LogEntry) GetUuid() string{
	return this.UUID
}
func (this *LogEntry) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *LogEntry) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *LogEntry) GetTypeName() string{
	this.TYPENAME = "CargoEntities.LogEntry"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *LogEntry) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *LogEntry) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *LogEntry) GetParentLnk() string{
	return this.ParentLnk
}
func (this *LogEntry) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *LogEntry) IsNeedSave() bool{
	return this.NeedSave
}


/** Id **/
func (this *LogEntry) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *LogEntry) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** CreationTime **/
func (this *LogEntry) GetCreationTime() int64{
	return this.M_creationTime
}

/** Init reference CreationTime **/
func (this *LogEntry) SetCreationTime(ref interface{}){
	if this.M_creationTime != ref.(int64) {
		this.M_creationTime = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference CreationTime **/

/** EntityRef **/
func (this *LogEntry) GetEntityRef() Entity{
	return this.m_entityRef
}

/** Init reference EntityRef **/
func (this *LogEntry) SetEntityRef(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_entityRef != ref.(string) {
			this.M_entityRef = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_entityRef != ref.(Entity).GetUuid() {
			this.M_entityRef = ref.(Entity).GetUuid()
			this.NeedSave = true
		}
		this.m_entityRef = ref.(Entity)
	}
}

/** Remove reference EntityRef **/
func (this *LogEntry) RemoveEntityRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_entityRef!= nil {
		if toDelete.GetUuid() == this.m_entityRef.(Entity).GetUuid() {
			this.m_entityRef = nil
			this.M_entityRef = ""
			this.NeedSave = true
		}
	}
}

/** Logger **/
func (this *LogEntry) GetLoggerPtr() *Log{
	return this.m_loggerPtr
}

/** Init reference Logger **/
func (this *LogEntry) SetLoggerPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_loggerPtr != ref.(string) {
			this.M_loggerPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_loggerPtr != ref.(Entity).GetUuid() {
			this.M_loggerPtr = ref.(Entity).GetUuid()
			this.NeedSave = true
		}
		this.m_loggerPtr = ref.(*Log)
	}
}

/** Remove reference Logger **/
func (this *LogEntry) RemoveLoggerPtr(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_loggerPtr!= nil {
		if toDelete.GetUuid() == this.m_loggerPtr.GetUuid() {
			this.m_loggerPtr = nil
			this.M_loggerPtr = ""
			this.NeedSave = true
		}
	}
}

/** Entities **/
func (this *LogEntry) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *LogEntry) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_entitiesPtr != ref.(string) {
			this.M_entitiesPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_entitiesPtr != ref.(*Entities).GetUuid() {
			this.M_entitiesPtr = ref.(*Entities).GetUuid()
			this.NeedSave = true
		}
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *LogEntry) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
