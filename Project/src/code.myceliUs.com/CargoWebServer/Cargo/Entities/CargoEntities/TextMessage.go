// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type TextMessage struct{

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

	/** members of Message **/
	M_body string

	/** members of TextMessage **/
	M_creationTime int64
	M_fromRef string
	M_toRef string
	M_title string


	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for TextMessage **/
type XsdTextMessage struct {
	XMLName xml.Name	`xml:"textMessage"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	/** Message **/
	M_body	string	`xml:"body,attr"`


	M_fromRef	*string	`xml:"fromRef"`
	M_toRef	*string	`xml:"toRef"`
	M_title	string	`xml:"title,attr"`
	M_creationTime	int64	`xml:"creationTime,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *TextMessage) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *TextMessage) SetUuid(uuid string){
	this.NeedSave = this.UUID == uuid
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *TextMessage) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *TextMessage) GetTypeName() string{
	this.TYPENAME = "CargoEntities.TextMessage"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *TextMessage) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *TextMessage) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *TextMessage) GetParentLnk() string{
	return this.ParentLnk
}
func (this *TextMessage) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *TextMessage) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *TextMessage) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Evaluate if an entity needs to be saved. **/
func (this *TextMessage) IsNeedSave() bool{
	return this.NeedSave
}
func (this *TextMessage) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *TextMessage) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *TextMessage) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *TextMessage) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *TextMessage) GetId()string{
	return this.M_id
}

func (this *TextMessage) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}




func (this *TextMessage) GetBody()string{
	return this.M_body
}

func (this *TextMessage) SetBody(val string){
	this.NeedSave = this.M_body== val
	this.M_body= val
}




func (this *TextMessage) GetCreationTime()int64{
	return this.M_creationTime
}

func (this *TextMessage) SetCreationTime(val int64){
	this.NeedSave = this.M_creationTime== val
	this.M_creationTime= val
}




func (this *TextMessage) GetFromRef()*Account{
	entity, err := this.getEntityByUuid(this.M_fromRef)
	if err == nil {
		return entity.(*Account)
	}
	return nil
}

func (this *TextMessage) SetFromRef(val *Account){
	this.NeedSave = this.M_fromRef != val.GetUuid()
	this.M_fromRef= val.GetUuid()
}


func (this *TextMessage) ResetFromRef(){
	this.M_fromRef= ""
}


func (this *TextMessage) GetToRef()*Account{
	entity, err := this.getEntityByUuid(this.M_toRef)
	if err == nil {
		return entity.(*Account)
	}
	return nil
}

func (this *TextMessage) SetToRef(val *Account){
	this.NeedSave = this.M_toRef != val.GetUuid()
	this.M_toRef= val.GetUuid()
}


func (this *TextMessage) ResetToRef(){
	this.M_toRef= ""
}


func (this *TextMessage) GetTitle()string{
	return this.M_title
}

func (this *TextMessage) SetTitle(val string){
	this.NeedSave = this.M_title== val
	this.M_title= val
}




func (this *TextMessage) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *TextMessage) SetEntitiesPtr(val *Entities){
	this.NeedSave = this.M_entitiesPtr != val.GetUuid()
	this.M_entitiesPtr= val.GetUuid()
}


func (this *TextMessage) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

