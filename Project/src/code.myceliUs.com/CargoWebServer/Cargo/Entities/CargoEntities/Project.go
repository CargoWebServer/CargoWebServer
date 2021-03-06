// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type Project struct{

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

	/** members of Project **/
	M_name string
	M_filesRef []string


	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for Project **/
type XsdProject struct {
	XMLName xml.Name	`xml:"project"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_filesRef	[]string	`xml:"filesRef"`
	M_name	string	`xml:"name,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Project) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Project) SetUuid(uuid string){
	this.UUID = uuid
}

/** Need save **/
func (this *Project) IsNeedSave() bool{
	return this.needSave
}
func (this *Project) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *Project) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *Project) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *Project) RemoveReferenced(uuid string, field string) {
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

func (this *Project) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *Project) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *Project) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Project) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Project"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Project) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Project) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Project) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Project) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *Project) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Project) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *Project) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Project) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Project) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Project) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Project) GetId()string{
	return this.M_id
}

func (this *Project) SetId(val string){
	this.M_id= val
}




func (this *Project) GetName()string{
	return this.M_name
}

func (this *Project) SetName(val string){
	this.M_name= val
}




func (this *Project) GetFilesRef()[]*File{
	values := make([]*File, 0)
	for i := 0; i < len(this.M_filesRef); i++ {
		entity, err := this.getEntityByUuid(this.M_filesRef[i])
		if err == nil {
			values = append( values, entity.(*File))
		}
	}
	return values
}

func (this *Project) SetFilesRef(val []*File){
	this.M_filesRef= make([]string,0)
	for i:=0; i < len(val); i++{
		this.setEntity(val[i])
	}
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Project) AppendFilesRef(val *File){
	for i:=0; i < len(this.M_filesRef); i++{
		if this.M_filesRef[i] == val.GetUuid() {
			return
		}
	}
	if this.M_filesRef== nil {
		this.M_filesRef = make([]string, 0)
	}

	this.M_filesRef = append(this.M_filesRef, val.GetUuid())
	this.setEntity(this)
	this.SetNeedSave(true)
}

func (this *Project) RemoveFilesRef(val *File){
	values := make([]string,0)
	for i:=0; i < len(this.M_filesRef); i++{
		if this.M_filesRef[i] != val.GetUuid() {
			values = append(values, this.M_filesRef[i])
		}
	}
	this.M_filesRef = values
	this.setEntity(this)
}


func (this *Project) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Project) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Project) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
	this.setEntity(this)
}

