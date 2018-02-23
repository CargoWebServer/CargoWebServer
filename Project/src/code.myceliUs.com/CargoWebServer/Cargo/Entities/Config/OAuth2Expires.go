// +build Config

package Config

import(
	"encoding/xml"
)

type OAuth2Expires struct{

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

	/** members of OAuth2Expires **/
	M_id string
	M_expiresAt int64


	/** Associations **/
	m_parentPtr *OAuth2Configuration
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2Expires **/
type XsdOAuth2Expires struct {
	XMLName xml.Name	`xml:"oauth2Expires"`
	M_id	string	`xml:"id,attr"`
	M_expiresAt	int64	`xml:"expiresAt,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2Expires) GetUuid() string{
	return this.UUID
}
func (this *OAuth2Expires) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2Expires) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2Expires) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2Expires"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2Expires) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2Expires) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Expires) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2Expires) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *OAuth2Expires) IsNeedSave() bool{
	return this.NeedSave
}
func (this *OAuth2Expires) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Expires) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *OAuth2Expires) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Expires) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** ExpiresAt **/
func (this *OAuth2Expires) GetExpiresAt() int64{
	return this.M_expiresAt
}

/** Init reference ExpiresAt **/
func (this *OAuth2Expires) SetExpiresAt(ref interface{}){
	if this.M_expiresAt != ref.(int64) {
		this.M_expiresAt = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference ExpiresAt **/

/** Parent **/
func (this *OAuth2Expires) GetParentPtr() *OAuth2Configuration{
	if this.m_parentPtr == nil {
		entity, err := this.getEntityByUuid(this.M_parentPtr)
		if err == nil {
			this.m_parentPtr = entity.(*OAuth2Configuration)
		}
	}
	return this.m_parentPtr
}
func (this *OAuth2Expires) GetParentPtrStr() string{
	return this.M_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Expires) SetParentPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_parentPtr != ref.(string) {
			this.M_parentPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_parentPtr != ref.(Configuration).GetUuid() {
			this.M_parentPtr = ref.(Configuration).GetUuid()
			this.NeedSave = true
		}
		this.m_parentPtr = ref.(*OAuth2Configuration)
	}
}

/** Remove reference Parent **/
func (this *OAuth2Expires) RemoveParentPtr(ref interface{}){
	toDelete := ref.(Configuration)
	if this.m_parentPtr!= nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
