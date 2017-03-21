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
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of OAuth2Refresh **/
	M_token string
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

}
/** UUID **/
func (this *OAuth2Refresh) GetUUID() string{
	return this.UUID
}

/** Token **/
func (this *OAuth2Refresh) GetToken() string{
	return this.M_token
}

/** Init reference Token **/
func (this *OAuth2Refresh) SetToken(ref interface{}){
	this.NeedSave = true
	this.M_token = ref.(string)
}

/** Remove reference Token **/

/** Access **/
func (this *OAuth2Refresh) GetAccess() *OAuth2Access{
	return this.m_access
}

/** Init reference Access **/
func (this *OAuth2Refresh) SetAccess(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_access = ref.(string)
	}else{
		this.m_access = ref.(*OAuth2Access)
		this.M_access = ref.(*OAuth2Access).GetUUID()
	}
}

/** Remove reference Access **/
func (this *OAuth2Refresh) RemoveAccess(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Access)
	if toDelete.GetUUID() == this.m_access.GetUUID() {
		this.m_access = nil
		this.M_access = ""
	}
}

/** Parent **/
func (this *OAuth2Refresh) GetParentPtr() *OAuth2Configuration{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Refresh) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*OAuth2Configuration)
		this.M_parentPtr = ref.(Configuration).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *OAuth2Refresh) RemoveParentPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Configuration)
	if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
		this.m_parentPtr = nil
		this.M_parentPtr = ""
	}
}
