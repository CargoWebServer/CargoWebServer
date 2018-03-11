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
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of Entity **/
	M_id string

	/** members of Log **/
	M_entries []string


	/** Associations **/
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
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Log) SetUuid(uuid string){
	this.NeedSave = this.UUID == uuid
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

func (this *Log) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Log) GetChilds() []interface{}{
	var childs []interface{}
	var child interface{}
	var err error
	for i:=0; i < len(this.M_entries); i++ {
		child, err = this.getEntityByUuid( this.M_entries[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	return childs
}
/** Evaluate if an entity needs to be saved. **/
func (this *Log) IsNeedSave() bool{
	return this.NeedSave
}
func (this *Log) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Log) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Log) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Log) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Log) GetId()string{
	return this.M_id
}

func (this *Log) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}




func (this *Log) GetEntries()[]*LogEntry{
	values := make([]*LogEntry, 0)
	for i := 0; i < len(this.M_entries); i++ {
		entity, err := this.getEntityByUuid(this.M_entries[i])
		if err == nil {
			values = append( values, entity.(*LogEntry))
		}
	}
	return values
}

func (this *Log) SetEntries(val []*LogEntry){
	this.M_entries= make([]string,0)
	for i:=0; i < len(val); i++{
		val[i].SetParentUuid(this.UUID)
		val[i].SetParentLnk("M_entries")
		this.setEntity(val[i])
		this.M_entries=append(this.M_entries, val[i].GetUuid())
	}
	this.NeedSave= true
}


func (this *Log) AppendEntries(val *LogEntry){
	for i:=0; i < len(this.M_entries); i++{
		if this.M_entries[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	val.SetParentUuid(this.UUID)
	val.SetParentLnk("M_entries")
	this.setEntity(val)
	this.M_entries = append(this.M_entries, val.GetUuid())
}

func (this *Log) RemoveEntries(val *LogEntry){
	values := make([]string,0)
	for i:=0; i < len(this.M_entries); i++{
		if this.M_entries[i] != val.GetUuid() {
			values = append(values, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_entries = values
}


func (this *Log) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Log) SetEntitiesPtr(val *Entities){
	this.NeedSave = this.M_entitiesPtr != val.GetUuid()
	this.M_entitiesPtr= val.GetUuid()
}


func (this *Log) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

