// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type Notification struct{

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

	/** members of Message **/
	M_body string

	/** members of Notification **/
	M_fromRef string
	M_toRef string
	M_type string
	M_code int


	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for Notification **/
type XsdNotification struct {
	XMLName xml.Name	`xml:"notification"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	/** Message **/
	M_body	string	`xml:"body,attr"`


	M_fromRef	*string	`xml:"fromRef"`
	M_toRef	*string	`xml:"toRef"`
	M_type	string	`xml:"type,attr"`
	M_code	string	`xml:"code,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Notification) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Notification) SetUuid(uuid string){
	this.UUID = uuid
}

/** Need save **/
func (this *Notification) IsNeedSave() bool{
	return this.needSave
}
func (this *Notification) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *Notification) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *Notification) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *Notification) RemoveReferenced(uuid string, field string) {
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

func (this *Notification) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *Notification) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *Notification) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Notification) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Notification"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Notification) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Notification) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Notification) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Notification) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *Notification) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Notification) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *Notification) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Notification) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Notification) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Notification) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Notification) GetId()string{
	return this.M_id
}

func (this *Notification) SetId(val string){
	this.M_id= val
}




func (this *Notification) GetBody()string{
	return this.M_body
}

func (this *Notification) SetBody(val string){
	this.M_body= val
}




func (this *Notification) GetFromRef()*Account{
	entity, err := this.getEntityByUuid(this.M_fromRef)
	if err == nil {
		return entity.(*Account)
	}
	return nil
}

func (this *Notification) SetFromRef(val *Account){
	this.M_fromRef= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Notification) ResetFromRef(){
	this.M_fromRef= ""
	this.setEntity(this)
}


func (this *Notification) GetToRef()*Account{
	entity, err := this.getEntityByUuid(this.M_toRef)
	if err == nil {
		return entity.(*Account)
	}
	return nil
}

func (this *Notification) SetToRef(val *Account){
	this.M_toRef= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Notification) ResetToRef(){
	this.M_toRef= ""
	this.setEntity(this)
}


func (this *Notification) GetType()string{
	return this.M_type
}

func (this *Notification) SetType(val string){
	this.M_type= val
}




func (this *Notification) GetCode()int{
	return this.M_code
}

func (this *Notification) SetCode(val int){
	this.M_code= val
}




func (this *Notification) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Notification) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Notification) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
	this.setEntity(this)
}

