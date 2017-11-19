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
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
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
		if this.IsInit == true {			this.NeedSave = true
		}
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
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_entityRef != ref.(Entity).GetUUID() {
			this.M_entityRef = ref.(Entity).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_entityRef = ref.(Entity)
	}
}

/** Remove reference EntityRef **/
func (this *LogEntry) RemoveEntityRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_entityRef!= nil {
		if toDelete.GetUUID() == this.m_entityRef.(Entity).GetUUID() {
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
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_loggerPtr != ref.(Entity).GetUUID() {
			this.M_loggerPtr = ref.(Entity).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_loggerPtr = ref.(*Log)
	}
}

/** Remove reference Logger **/
func (this *LogEntry) RemoveLoggerPtr(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_loggerPtr!= nil {
		if toDelete.GetUUID() == this.m_loggerPtr.GetUUID() {
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
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_entitiesPtr != ref.(*Entities).GetUUID() {
			this.M_entitiesPtr = ref.(*Entities).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *LogEntry) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
