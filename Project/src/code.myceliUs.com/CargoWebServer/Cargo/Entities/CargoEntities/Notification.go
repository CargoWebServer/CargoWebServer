// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
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
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

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
	return this.UUID
}
func (this *Notification) SetUuid(uuid string){
	this.UUID = uuid
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

/** Evaluate if an entity needs to be saved. **/
func (this *Notification) IsNeedSave() bool{
	return this.NeedSave
}
func (this *Notification) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Notification) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

func (this *Notification) GetId()string{
	return this.M_id
}

func (this *Notification) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *Notification) GetBody()string{
	return this.M_body
}

func (this *Notification) SetBody(val string){
	this.NeedSave = this.M_body== val
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
	this.NeedSave = this.M_fromRef != val.GetUuid()
	this.M_fromRef= val.GetUuid()
}

func (this *Notification) ResetFromRef(){
	this.M_fromRef= ""
}


func (this *Notification) GetToRef()*Account{
	entity, err := this.getEntityByUuid(this.M_toRef)
	if err == nil {
		return entity.(*Account)
	}
	return nil
}

func (this *Notification) SetToRef(val *Account){
	this.NeedSave = this.M_toRef != val.GetUuid()
	this.M_toRef= val.GetUuid()
}

func (this *Notification) ResetToRef(){
	this.M_toRef= ""
}


func (this *Notification) GetType()string{
	return this.M_type
}

func (this *Notification) SetType(val string){
	this.NeedSave = this.M_type== val
	this.M_type= val
}


func (this *Notification) GetCode()int{
	return this.M_code
}

func (this *Notification) SetCode(val int){
	this.NeedSave = this.M_code== val
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
	this.NeedSave = this.M_entitiesPtr != val.GetUuid()
	this.M_entitiesPtr= val.GetUuid()
}

func (this *Notification) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

