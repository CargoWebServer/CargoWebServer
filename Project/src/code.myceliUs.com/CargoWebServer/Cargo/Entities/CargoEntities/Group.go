// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type Group struct{

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

	/** members of Group **/
	M_name string
	M_membersRef []string


	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for Group **/
type XsdGroup struct {
	XMLName xml.Name	`xml:"memberOfRef"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_membersRef	[]string	`xml:"membersRef"`
	M_name	string	`xml:"name,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Group) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Group) SetUuid(uuid string){
	this.UUID = uuid
}

/** Need save **/
func (this *Group) IsNeedSave() bool{
	return this.needSave
}
func (this *Group) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *Group) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *Group) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *Group) RemoveReferenced(uuid string, field string) {
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

func (this *Group) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *Group) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *Group) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Group) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Group"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Group) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Group) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Group) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Group) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *Group) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Group) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *Group) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Group) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Group) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Group) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Group) GetId()string{
	return this.M_id
}

func (this *Group) SetId(val string){
	this.M_id= val
}




func (this *Group) GetName()string{
	return this.M_name
}

func (this *Group) SetName(val string){
	this.M_name= val
}




func (this *Group) GetMembersRef()[]*User{
	values := make([]*User, 0)
	for i := 0; i < len(this.M_membersRef); i++ {
		entity, err := this.getEntityByUuid(this.M_membersRef[i])
		if err == nil {
			values = append( values, entity.(*User))
		}
	}
	return values
}

func (this *Group) SetMembersRef(val []*User){
	this.M_membersRef= make([]string,0)
	for i:=0; i < len(val); i++{
		this.setEntity(val[i])
	}
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Group) AppendMembersRef(val *User){
	for i:=0; i < len(this.M_membersRef); i++{
		if this.M_membersRef[i] == val.GetUuid() {
			return
		}
	}
	if this.M_membersRef== nil {
		this.M_membersRef = make([]string, 0)
	}

	this.M_membersRef = append(this.M_membersRef, val.GetUuid())
	this.setEntity(this)
	this.SetNeedSave(true)
}

func (this *Group) RemoveMembersRef(val *User){
	values := make([]string,0)
	for i:=0; i < len(this.M_membersRef); i++{
		if this.M_membersRef[i] != val.GetUuid() {
			values = append(values, this.M_membersRef[i])
		}
	}
	this.M_membersRef = values
	this.setEntity(this)
}


func (this *Group) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Group) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Group) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
	this.setEntity(this)
}

