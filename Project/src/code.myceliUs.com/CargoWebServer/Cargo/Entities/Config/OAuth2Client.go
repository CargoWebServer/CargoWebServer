// +build Config

package Config

import(
"encoding/xml"
)

type OAuth2Client struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of OAuth2Client **/
	M_id string
	M_secret string
	M_redirectUri string
	M_extra []uint8


	/** Associations **/
	m_parentPtr *OAuth2Configuration
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2Client **/
type XsdOAuth2Client struct {
	XMLName xml.Name	`xml:"oauth2Client"`

}
/** UUID **/
func (this *OAuth2Client) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *OAuth2Client) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Client) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Secret **/
func (this *OAuth2Client) GetSecret() string{
	return this.M_secret
}

/** Init reference Secret **/
func (this *OAuth2Client) SetSecret(ref interface{}){
	this.NeedSave = true
	this.M_secret = ref.(string)
}

/** Remove reference Secret **/

/** RedirectUri **/
func (this *OAuth2Client) GetRedirectUri() string{
	return this.M_redirectUri
}

/** Init reference RedirectUri **/
func (this *OAuth2Client) SetRedirectUri(ref interface{}){
	this.NeedSave = true
	this.M_redirectUri = ref.(string)
}

/** Remove reference RedirectUri **/

/** Extra **/
func (this *OAuth2Client) GetExtra() []uint8{
	return this.M_extra
}

/** Init reference Extra **/
func (this *OAuth2Client) SetExtra(ref interface{}){
	this.NeedSave = true
	this.M_extra = ref.([]uint8)
}

/** Remove reference Extra **/

/** Parent **/
func (this *OAuth2Client) GetParentPtr() *OAuth2Configuration{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Client) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*OAuth2Configuration)
		this.M_parentPtr = ref.(Configuration).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *OAuth2Client) RemoveParentPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Configuration)
	if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
		this.m_parentPtr = nil
		this.M_parentPtr = ""
	}
}
