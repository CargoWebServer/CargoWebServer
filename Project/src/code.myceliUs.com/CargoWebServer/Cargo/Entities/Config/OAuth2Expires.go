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
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of OAuth2Expires **/
	M_id string
	m_token *OAuth2Refresh
	/** If the ref is a string and not an object **/
	M_token string
	M_expiresAt int64


	/** Associations **/
	m_parentPtr *OAuth2Configuration
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2Expires **/
type XsdOAuth2Expires struct {
	XMLName xml.Name	`xml:"oauth2Expires"`

}
/** UUID **/
func (this *OAuth2Expires) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *OAuth2Expires) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Expires) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Token **/
func (this *OAuth2Expires) GetToken() *OAuth2Refresh{
	return this.m_token
}

/** Init reference Token **/
func (this *OAuth2Expires) SetToken(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_token = ref.(string)
	}else{
		this.m_token = ref.(*OAuth2Refresh)
		this.M_token = ref.(*OAuth2Refresh).GetUUID()
	}
}

/** Remove reference Token **/
func (this *OAuth2Expires) RemoveToken(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Refresh)
	if toDelete.GetUUID() == this.m_token.GetUUID() {
		this.m_token = nil
		this.M_token = ""
	}
}

/** ExpiresAt **/
func (this *OAuth2Expires) GetExpiresAt() int64{
	return this.M_expiresAt
}

/** Init reference ExpiresAt **/
func (this *OAuth2Expires) SetExpiresAt(ref interface{}){
	this.NeedSave = true
	this.M_expiresAt = ref.(int64)
}

/** Remove reference ExpiresAt **/

/** Parent **/
func (this *OAuth2Expires) GetParentPtr() *OAuth2Configuration{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Expires) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*OAuth2Configuration)
		this.M_parentPtr = ref.(Configuration).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *OAuth2Expires) RemoveParentPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Configuration)
	if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
		this.m_parentPtr = nil
		this.M_parentPtr = ""
	}
}
