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
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

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
/** UUID **/
func (this *LogEntry) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *LogEntry) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *LogEntry) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** CreationTime **/
func (this *LogEntry) GetCreationTime() int64{
	return this.M_creationTime
}

/** Init reference CreationTime **/
func (this *LogEntry) SetCreationTime(ref interface{}){
	this.NeedSave = true
	this.M_creationTime = ref.(int64)
}

/** Remove reference CreationTime **/

/** EntityRef **/
func (this *LogEntry) GetEntityRef() Entity{
	return this.m_entityRef
}

/** Init reference EntityRef **/
func (this *LogEntry) SetEntityRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entityRef = ref.(string)
	}else{
		this.m_entityRef = ref.(Entity)
		this.M_entityRef = ref.(Entity).GetUUID()
	}
}

/** Remove reference EntityRef **/
func (this *LogEntry) RemoveEntityRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_entityRef!= nil {
		if toDelete.GetUUID() == this.m_entityRef.(Entity).GetUUID() {
			this.m_entityRef = nil
			this.M_entityRef = ""
		}else{
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
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_loggerPtr = ref.(string)
	}else{
		this.m_loggerPtr = ref.(*Log)
		this.M_loggerPtr = ref.(Entity).GetUUID()
	}
}

/** Remove reference Logger **/
func (this *LogEntry) RemoveLoggerPtr(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_loggerPtr!= nil {
		if toDelete.GetUUID() == this.m_loggerPtr.GetUUID() {
			this.m_loggerPtr = nil
			this.M_loggerPtr = ""
		}else{
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
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *LogEntry) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
		}else{
			this.NeedSave = true
		}
	}
}
