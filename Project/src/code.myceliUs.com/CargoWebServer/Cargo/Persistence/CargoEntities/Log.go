package CargoEntities

import(
"encoding/xml"
)

type Log struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Entity **/
	M_id string

	/** members of Log **/
	M_entries []*LogEntry


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for Log **/
type XsdLog struct {
	XMLName xml.Name	`xml:"log"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_creationTime	int64	`xml:"creationTime,attr"`

}
/** UUID **/
func (this *Log) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Log) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Log) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Entries **/
func (this *Log) GetEntries() []*LogEntry{
	return this.M_entries
}

/** Init reference Entries **/
func (this *Log) SetEntries(ref interface{}){
	this.NeedSave = true
	isExist := false
	var entriess []*LogEntry
	for i:=0; i<len(this.M_entries); i++ {
		if this.M_entries[i].GetUUID() != ref.(Entity).GetUUID() {
			entriess = append(entriess, this.M_entries[i])
		} else {
			isExist = true
			entriess = append(entriess, ref.(*LogEntry))
		}
	}
	if !isExist {
		entriess = append(entriess, ref.(*LogEntry))
	}
	this.M_entries = entriess
}

/** Remove reference Entries **/
func (this *Log) RemoveEntries(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Entity)
	entries_ := make([]*LogEntry, 0)
	for i := 0; i < len(this.M_entries); i++ {
		if toDelete.GetUUID() != this.M_entries[i].GetUUID() {
			entries_ = append(entries_, this.M_entries[i])
		}
	}
	this.M_entries = entries_
}

/** Entities **/
func (this *Log) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Log) SetEntitiesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *Log) RemoveEntitiesPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Entities)
	if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
		this.m_entitiesPtr = nil
		this.M_entitiesPtr = ""
	}
}
