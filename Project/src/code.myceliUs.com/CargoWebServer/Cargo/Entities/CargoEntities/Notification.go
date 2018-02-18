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
	m_fromRef *Account
	/** If the ref is a string and not an object **/
	M_fromRef string
	m_toRef *Account
	/** If the ref is a string and not an object **/
	M_toRef string
	M_type string
	M_code int


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
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

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Notification) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *Notification) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Notification) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Body **/
func (this *Notification) GetBody() string{
	return this.M_body
}

/** Init reference Body **/
func (this *Notification) SetBody(ref interface{}){
	if this.M_body != ref.(string) {
		this.M_body = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Body **/

/** FromRef **/
func (this *Notification) GetFromRef() *Account{
	if this.m_fromRef == nil {
		entity, err := this.getEntityByUuid(this.M_fromRef)
		if err == nil {
			this.m_fromRef = entity.(*Account)
		}
	}
	return this.m_fromRef
}
func (this *Notification) GetFromRefStr() string{
	return this.M_fromRef
}

/** Init reference FromRef **/
func (this *Notification) SetFromRef(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_fromRef != ref.(string) {
			this.M_fromRef = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_fromRef != ref.(Entity).GetUuid() {
			this.M_fromRef = ref.(Entity).GetUuid()
			this.NeedSave = true
		}
		this.m_fromRef = ref.(*Account)
	}
}

/** Remove reference FromRef **/
func (this *Notification) RemoveFromRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_fromRef!= nil {
		if toDelete.GetUuid() == this.m_fromRef.GetUuid() {
			this.m_fromRef = nil
			this.M_fromRef = ""
			this.NeedSave = true
		}
	}
}

/** ToRef **/
func (this *Notification) GetToRef() *Account{
	if this.m_toRef == nil {
		entity, err := this.getEntityByUuid(this.M_toRef)
		if err == nil {
			this.m_toRef = entity.(*Account)
		}
	}
	return this.m_toRef
}
func (this *Notification) GetToRefStr() string{
	return this.M_toRef
}

/** Init reference ToRef **/
func (this *Notification) SetToRef(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_toRef != ref.(string) {
			this.M_toRef = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_toRef != ref.(Entity).GetUuid() {
			this.M_toRef = ref.(Entity).GetUuid()
			this.NeedSave = true
		}
		this.m_toRef = ref.(*Account)
	}
}

/** Remove reference ToRef **/
func (this *Notification) RemoveToRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_toRef!= nil {
		if toDelete.GetUuid() == this.m_toRef.GetUuid() {
			this.m_toRef = nil
			this.M_toRef = ""
			this.NeedSave = true
		}
	}
}

/** Type **/
func (this *Notification) GetType() string{
	return this.M_type
}

/** Init reference Type **/
func (this *Notification) SetType(ref interface{}){
	if this.M_type != ref.(string) {
		this.M_type = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Type **/

/** Code **/
func (this *Notification) GetCode() int{
	return this.M_code
}

/** Init reference Code **/
func (this *Notification) SetCode(ref interface{}){
	if this.M_code != ref.(int) {
		this.M_code = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference Code **/

/** Entities **/
func (this *Notification) GetEntitiesPtr() *Entities{
	if this.m_entitiesPtr == nil {
		entity, err := this.getEntityByUuid(this.M_entitiesPtr)
		if err == nil {
			this.m_entitiesPtr = entity.(*Entities)
		}
	}
	return this.m_entitiesPtr
}
func (this *Notification) GetEntitiesPtrStr() string{
	return this.M_entitiesPtr
}

/** Init reference Entities **/
func (this *Notification) SetEntitiesPtr(ref interface{}){
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
func (this *Notification) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
