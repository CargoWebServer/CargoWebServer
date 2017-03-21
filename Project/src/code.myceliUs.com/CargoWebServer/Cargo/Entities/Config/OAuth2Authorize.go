// +build Config

package Config

import(
"encoding/xml"
)

type OAuth2Authorize struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of OAuth2Authorize **/
	m_client *OAuth2Client
	/** If the ref is a string and not an object **/
	M_client string
	M_code string
	M_expiresIn int64
	M_scope string
	M_redirectUri string
	M_state string
	M_extra []uint8
	M_createdAt int64

}

/** Xml parser for OAuth2Authorize **/
type XsdOAuth2Authorize struct {
	XMLName xml.Name	`xml:"oauth2Authorize"`

}
/** UUID **/
func (this *OAuth2Authorize) GetUUID() string{
	return this.UUID
}

/** Client **/
func (this *OAuth2Authorize) GetClient() *OAuth2Client{
	return this.m_client
}

/** Init reference Client **/
func (this *OAuth2Authorize) SetClient(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_client = ref.(string)
	}else{
		this.m_client = ref.(*OAuth2Client)
		this.M_client = ref.(*OAuth2Client).GetUUID()
	}
}

/** Remove reference Client **/
func (this *OAuth2Authorize) RemoveClient(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Client)
	if toDelete.GetUUID() == this.m_client.GetUUID() {
		this.m_client = nil
		this.M_client = ""
	}
}

/** Code **/
func (this *OAuth2Authorize) GetCode() string{
	return this.M_code
}

/** Init reference Code **/
func (this *OAuth2Authorize) SetCode(ref interface{}){
	this.NeedSave = true
	this.M_code = ref.(string)
}

/** Remove reference Code **/

/** ExpiresIn **/
func (this *OAuth2Authorize) GetExpiresIn() int64{
	return this.M_expiresIn
}

/** Init reference ExpiresIn **/
func (this *OAuth2Authorize) SetExpiresIn(ref interface{}){
	this.NeedSave = true
	this.M_expiresIn = ref.(int64)
}

/** Remove reference ExpiresIn **/

/** Scope **/
func (this *OAuth2Authorize) GetScope() string{
	return this.M_scope
}

/** Init reference Scope **/
func (this *OAuth2Authorize) SetScope(ref interface{}){
	this.NeedSave = true
	this.M_scope = ref.(string)
}

/** Remove reference Scope **/

/** RedirectUri **/
func (this *OAuth2Authorize) GetRedirectUri() string{
	return this.M_redirectUri
}

/** Init reference RedirectUri **/
func (this *OAuth2Authorize) SetRedirectUri(ref interface{}){
	this.NeedSave = true
	this.M_redirectUri = ref.(string)
}

/** Remove reference RedirectUri **/

/** State **/
func (this *OAuth2Authorize) GetState() string{
	return this.M_state
}

/** Init reference State **/
func (this *OAuth2Authorize) SetState(ref interface{}){
	this.NeedSave = true
	this.M_state = ref.(string)
}

/** Remove reference State **/

/** Extra **/
func (this *OAuth2Authorize) GetExtra() []uint8{
	return this.M_extra
}

/** Init reference Extra **/
func (this *OAuth2Authorize) SetExtra(ref interface{}){
	this.NeedSave = true
	this.M_extra = ref.([]uint8)
}

/** Remove reference Extra **/

/** CreatedAt **/
func (this *OAuth2Authorize) GetCreatedAt() int64{
	return this.M_createdAt
}

/** Init reference CreatedAt **/
func (this *OAuth2Authorize) SetCreatedAt(ref interface{}){
	this.NeedSave = true
	this.M_createdAt = ref.(int64)
}

/** Remove reference CreatedAt **/
