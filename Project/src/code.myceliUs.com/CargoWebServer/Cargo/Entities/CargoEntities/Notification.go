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
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
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
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference Body **/

/** FromRef **/
func (this *Notification) GetFromRef() *Account{
	return this.m_fromRef
}

/** Init reference FromRef **/
func (this *Notification) SetFromRef(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_fromRef != ref.(string) {
			this.M_fromRef = ref.(string)
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_fromRef != ref.(Entity).GetUUID() {
			this.M_fromRef = ref.(Entity).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_fromRef = ref.(*Account)
	}
}

/** Remove reference FromRef **/
func (this *Notification) RemoveFromRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_fromRef!= nil {
		if toDelete.GetUUID() == this.m_fromRef.GetUUID() {
			this.m_fromRef = nil
			this.M_fromRef = ""
			this.NeedSave = true
		}
	}
}

/** ToRef **/
func (this *Notification) GetToRef() *Account{
	return this.m_toRef
}

/** Init reference ToRef **/
func (this *Notification) SetToRef(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_toRef != ref.(string) {
			this.M_toRef = ref.(string)
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_toRef != ref.(Entity).GetUUID() {
			this.M_toRef = ref.(Entity).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_toRef = ref.(*Account)
	}
}

/** Remove reference ToRef **/
func (this *Notification) RemoveToRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_toRef!= nil {
		if toDelete.GetUUID() == this.m_toRef.GetUUID() {
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
		if this.IsInit == true {			this.NeedSave = true
		}
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
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference Code **/

/** Entities **/
func (this *Notification) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Notification) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_entitiesPtr != ref.(string) {
			this.M_entitiesPtr = ref.(string)
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_entitiesPtr != ref.(*Entities).GetUUID() {
			this.M_entitiesPtr = ref.(*Entities).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Notification) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
