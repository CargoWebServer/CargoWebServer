// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
	"strings"
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
	/** keep track if the entity has change over time. **/
	needSave bool
	/** Keep reference to entity that made use of thit entity **/
	Referenced []string
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
	this.UUID = uuid
}

/** Need save **/
func (this *Log) IsNeedSave() bool{
	return this.needSave
}
func (this *Log) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *Log) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *Log) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *Log) RemoveReferenced(uuid string, field string) {
	if this.Referenced == nil {
		return
	}
	referenced := make([]string, 0)
	for i := 0; i < len(this.Referenced); i++ {
		if this.Referenced[i] != uuid+":"+field {
			referenced = append(referenced, uuid+":"+field)
		}
	}
	this.Referenced = referenced
}

func (this *Log) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *Log) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
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
/** Return the list of all childs uuid **/
func (this *Log) GetChildsUuid() []string{
	var childs []string
	childs = append( childs, this.M_entries...)
	return childs
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
		this.M_entries=append(this.M_entries, val[i].GetUuid())
		if len(val[i].GetParentUuid()) > 0  &&  len(val[i].GetParentLnk()) > 0 && this.GetUuid() != val[i].GetParentUuid(){
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_entries")
		this.setEntity(val[i])
	}
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Log) AppendEntries(val *LogEntry){
	for i:=0; i < len(this.M_entries); i++{
		if this.M_entries[i] == val.GetUuid() {
			return
		}
	}
	if this.M_entries== nil {
		this.M_entries = make([]string, 0)
	}

	this.M_entries = append(this.M_entries, val.GetUuid())
	if len(val.GetParentUuid()) > 0 &&  len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_entries")
  this.setEntity(val)
	this.setEntity(this)
	this.SetNeedSave(true)
}

func (this *Log) RemoveEntries(val *LogEntry){
	values := make([]string,0)
	for i:=0; i < len(this.M_entries); i++{
		if this.M_entries[i] != val.GetUuid() {
			values = append(values, this.M_entries[i])
		}
	}
	this.M_entries = values
	this.setEntity(this)
}


func (this *Log) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Log) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Log) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
	this.setEntity(this)
}

