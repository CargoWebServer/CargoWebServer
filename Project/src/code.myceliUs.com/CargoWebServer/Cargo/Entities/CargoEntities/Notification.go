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
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

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
/** UUID **/
func (this *Notification) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Notification) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Notification) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Body **/
func (this *Notification) GetBody() string{
	return this.M_body
}

/** Init reference Body **/
func (this *Notification) SetBody(ref interface{}){
	this.NeedSave = true
	this.M_body = ref.(string)
}

/** Remove reference Body **/

/** FromRef **/
func (this *Notification) GetFromRef() *Account{
	return this.m_fromRef
}

/** Init reference FromRef **/
func (this *Notification) SetFromRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_fromRef = ref.(string)
	}else{
		this.m_fromRef = ref.(*Account)
		this.M_fromRef = ref.(Entity).GetUUID()
	}
}

/** Remove reference FromRef **/
func (this *Notification) RemoveFromRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Entity)
	if toDelete.GetUUID() == this.m_fromRef.GetUUID() {
		this.m_fromRef = nil
		this.M_fromRef = ""
	}
}

/** ToRef **/
func (this *Notification) GetToRef() *Account{
	return this.m_toRef
}

/** Init reference ToRef **/
func (this *Notification) SetToRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_toRef = ref.(string)
	}else{
		this.m_toRef = ref.(*Account)
		this.M_toRef = ref.(Entity).GetUUID()
	}
}

/** Remove reference ToRef **/
func (this *Notification) RemoveToRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Entity)
	if toDelete.GetUUID() == this.m_toRef.GetUUID() {
		this.m_toRef = nil
		this.M_toRef = ""
	}
}

/** Type **/
func (this *Notification) GetType() string{
	return this.M_type
}

/** Init reference Type **/
func (this *Notification) SetType(ref interface{}){
	this.NeedSave = true
	this.M_type = ref.(string)
}

/** Remove reference Type **/

/** Code **/
func (this *Notification) GetCode() int{
	return this.M_code
}

/** Init reference Code **/
func (this *Notification) SetCode(ref interface{}){
	this.NeedSave = true
	this.M_code = ref.(int)
}

/** Remove reference Code **/

/** Entities **/
func (this *Notification) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Notification) SetEntitiesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *Notification) RemoveEntitiesPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Entities)
	if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
		this.m_entitiesPtr = nil
		this.M_entitiesPtr = ""
	}
}
