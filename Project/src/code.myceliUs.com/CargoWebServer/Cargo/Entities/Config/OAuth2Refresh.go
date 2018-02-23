// +build Config

package Config

import(
	"encoding/xml"
)

type OAuth2Refresh struct{

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

	/** members of OAuth2Refresh **/
	M_id string
	m_access *OAuth2Access
	/** If the ref is a string and not an object **/
	M_access string


	/** Associations **/
	m_parentPtr *OAuth2Configuration
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2Refresh **/
type XsdOAuth2Refresh struct {
	XMLName xml.Name	`xml:"oauth2Refresh"`
	M_id	string	`xml:"id,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2Refresh) GetUuid() string{
	return this.UUID
}
func (this *OAuth2Refresh) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2Refresh) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2Refresh) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2Refresh"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2Refresh) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2Refresh) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Refresh) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2Refresh) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *OAuth2Refresh) IsNeedSave() bool{
	return this.NeedSave
}
func (this *OAuth2Refresh) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Refresh) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *OAuth2Refresh) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Refresh) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Access **/
func (this *OAuth2Refresh) GetAccess() *OAuth2Access{
	if this.m_access == nil {
		entity, err := this.getEntityByUuid(this.M_access)
		if err == nil {
			this.m_access = entity.(*OAuth2Access)
		}
	}
	return this.m_access
}
func (this *OAuth2Refresh) GetAccessStr() string{
	return this.M_access
}

/** Init reference Access **/
func (this *OAuth2Refresh) SetAccess(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_access != ref.(string) {
			this.M_access = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_access != ref.(*OAuth2Access).GetUuid() {
			this.M_access = ref.(*OAuth2Access).GetUuid()
			this.NeedSave = true
		}
		this.m_access = ref.(*OAuth2Access)
	}
}

/** Remove reference Access **/
func (this *OAuth2Refresh) RemoveAccess(ref interface{}){
	toDelete := ref.(*OAuth2Access)
	if this.m_access!= nil {
		if toDelete.GetUuid() == this.m_access.GetUuid() {
			this.m_access = nil
			this.M_access = ""
			this.NeedSave = true
		}
	}
}

/** Parent **/
func (this *OAuth2Refresh) GetParentPtr() *OAuth2Configuration{
	if this.m_parentPtr == nil {
		entity, err := this.getEntityByUuid(this.M_parentPtr)
		if err == nil {
			this.m_parentPtr = entity.(*OAuth2Configuration)
		}
	}
	return this.m_parentPtr
}
func (this *OAuth2Refresh) GetParentPtrStr() string{
	return this.M_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Refresh) SetParentPtr(ref interface{}){
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
func (this *OAuth2Refresh) RemoveParentPtr(ref interface{}){
	toDelete := ref.(Configuration)
	if this.m_parentPtr!= nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
