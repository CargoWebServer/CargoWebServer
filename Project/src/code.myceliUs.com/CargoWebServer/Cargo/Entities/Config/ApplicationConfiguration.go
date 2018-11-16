// +build Config

package Config

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type ApplicationConfiguration struct{

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

	/** members of Configuration **/
	M_id string

	/** members of ApplicationConfiguration **/
	M_indexPage string


	/** Associations **/
	M_parentPtr string
}

/** Xml parser for ApplicationConfiguration **/
type XsdApplicationConfiguration struct {
	XMLName xml.Name	`xml:"applicationConfiguration"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_indexPage	string	`xml:"indexPage,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *ApplicationConfiguration) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *ApplicationConfiguration) SetUuid(uuid string){
	this.UUID = uuid
}

/** Need save **/
func (this *ApplicationConfiguration) IsNeedSave() bool{
	return this.needSave
}
func (this *ApplicationConfiguration) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *ApplicationConfiguration) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *ApplicationConfiguration) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *ApplicationConfiguration) RemoveReferenced(uuid string, field string) {
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

func (this *ApplicationConfiguration) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *ApplicationConfiguration) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *ApplicationConfiguration) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *ApplicationConfiguration) GetTypeName() string{
	this.TYPENAME = "Config.ApplicationConfiguration"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *ApplicationConfiguration) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *ApplicationConfiguration) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ApplicationConfiguration) GetParentLnk() string{
	return this.ParentLnk
}
func (this *ApplicationConfiguration) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *ApplicationConfiguration) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ApplicationConfiguration) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *ApplicationConfiguration) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *ApplicationConfiguration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *ApplicationConfiguration) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *ApplicationConfiguration) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *ApplicationConfiguration) GetId()string{
	return this.M_id
}

func (this *ApplicationConfiguration) SetId(val string){
	this.M_id= val
}




func (this *ApplicationConfiguration) GetIndexPage()string{
	return this.M_indexPage
}

func (this *ApplicationConfiguration) SetIndexPage(val string){
	this.M_indexPage= val
}




func (this *ApplicationConfiguration) GetParentPtr()*Configurations{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*Configurations)
	}
	return nil
}

func (this *ApplicationConfiguration) SetParentPtr(val *Configurations){
	this.M_parentPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *ApplicationConfiguration) ResetParentPtr(){
	this.M_parentPtr= ""
	this.setEntity(this)
}

