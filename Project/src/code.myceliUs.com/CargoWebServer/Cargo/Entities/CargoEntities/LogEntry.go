// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
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

	/** members of LogEntry **/
	M_creationTime int64
	M_entityRef string


	/** Associations **/
	M_loggerPtr string
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
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *LogEntry) SetUuid(uuid string){
	this.UUID = uuid
}

/** Need save **/
func (this *LogEntry) IsNeedSave() bool{
	return this.needSave
}
func (this *LogEntry) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *LogEntry) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *LogEntry) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *LogEntry) RemoveReferenced(uuid string, field string) {
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

func (this *LogEntry) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *LogEntry) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
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

func (this *LogEntry) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *LogEntry) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *LogEntry) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *LogEntry) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *LogEntry) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *LogEntry) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *LogEntry) GetId()string{
	return this.M_id
}

func (this *LogEntry) SetId(val string){
	this.M_id= val
}




func (this *LogEntry) GetCreationTime()int64{
	return this.M_creationTime
}

func (this *LogEntry) SetCreationTime(val int64){
	this.M_creationTime= val
}




func (this *LogEntry) GetEntityRef()Entity{
	entity, err := this.getEntityByUuid(this.M_entityRef)
	if err == nil {
		return entity.(Entity)
	}
	return nil
}

func (this *LogEntry) SetEntityRef(val Entity){
	this.M_entityRef= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}




func (this *LogEntry) GetLoggerPtr()*Log{
	entity, err := this.getEntityByUuid(this.M_loggerPtr)
	if err == nil {
		return entity.(*Log)
	}
	return nil
}

func (this *LogEntry) SetLoggerPtr(val *Log){
	this.M_loggerPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *LogEntry) ResetLoggerPtr(){
	this.M_loggerPtr= ""
	this.setEntity(this)
}


func (this *LogEntry) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *LogEntry) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *LogEntry) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
	this.setEntity(this)
}

