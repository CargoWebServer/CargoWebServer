// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type Log struct{

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
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

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


	M_entries	[]*XsdLogEntry	`xml:"entries,omitempty"`
	M_creationTime	int64	`xml:"creationTime,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Log) GetUuid() string{
	return this.UUID
}
func (this *Log) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Log) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Log) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Log"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Log) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Log) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Log) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Log) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *Log) IsNeedSave() bool{
	return this.NeedSave
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Log) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *Log) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Log) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Entries **/
func (this *Log) GetEntries() []*LogEntry{
	return this.M_entries
}

/** Init reference Entries **/
func (this *Log) SetEntries(ref interface{}){
	isExist := false
	var entriess []*LogEntry
	for i:=0; i<len(this.M_entries); i++ {
		if this.M_entries[i].GetUuid() != ref.(Entity).GetUuid() {
			entriess = append(entriess, this.M_entries[i])
		} else {
			isExist = true
			entriess = append(entriess, ref.(*LogEntry))
		}
	}
	if !isExist {
		entriess = append(entriess, ref.(*LogEntry))
		this.NeedSave = true
		this.M_entries = entriess
	}
}

/** Remove reference Entries **/
func (this *Log) RemoveEntries(ref interface{}){
	toDelete := ref.(Entity)
	entries_ := make([]*LogEntry, 0)
	for i := 0; i < len(this.M_entries); i++ {
		if toDelete.GetUuid() != this.M_entries[i].GetUuid() {
			entries_ = append(entries_, this.M_entries[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_entries = entries_
}

/** Entities **/
func (this *Log) GetEntitiesPtr() *Entities{
	if this.m_entitiesPtr == nil {
		entity, err := this.getEntityByUuid(this.M_entitiesPtr)
		if err == nil {
			this.m_entitiesPtr = entity.(*Entities)
		}
	}
	return this.m_entitiesPtr
}
func (this *Log) GetEntitiesPtrStr() string{
	return this.M_entitiesPtr
}

/** Init reference Entities **/
func (this *Log) SetEntitiesPtr(ref interface{}){
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
func (this *Log) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
