// +build Config

package Config

import(
"encoding/xml"
)

type OAuth2Access struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of OAuth2Access **/
	m_client *OAuth2Client
	/** If the ref is a string and not an object **/
	M_client string
	m_authorize *OAuth2Authorize
	/** If the ref is a string and not an object **/
	M_authorize string
	M_previous string
	M_accessToken string
	m_refreshToken *OAuth2Refresh
	/** If the ref is a string and not an object **/
	M_refreshToken string
	M_expiresIn int64
	M_scope string
	M_redirectUri string
	M_extra []uint8
	M_createdAt int64


	/** Associations **/
	m_parentPtr *OAuth2Configuration
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2Access **/
type XsdOAuth2Access struct {
	XMLName xml.Name	`xml:"oauth2Access"`

}
/** UUID **/
func (this *OAuth2Access) GetUUID() string{
	return this.UUID
}

/** Client **/
func (this *OAuth2Access) GetClient() *OAuth2Client{
	return this.m_client
}

/** Init reference Client **/
func (this *OAuth2Access) SetClient(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_client = ref.(string)
	}else{
		this.m_client = ref.(*OAuth2Client)
		this.M_client = ref.(*OAuth2Client).GetUUID()
	}
}

/** Remove reference Client **/
func (this *OAuth2Access) RemoveClient(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Client)
	if toDelete.GetUUID() == this.m_client.GetUUID() {
		this.m_client = nil
		this.M_client = ""
	}
}

/** Authorize **/
func (this *OAuth2Access) GetAuthorize() *OAuth2Authorize{
	return this.m_authorize
}

/** Init reference Authorize **/
func (this *OAuth2Access) SetAuthorize(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_authorize = ref.(string)
	}else{
		this.m_authorize = ref.(*OAuth2Authorize)
		this.M_authorize = ref.(*OAuth2Authorize).GetUUID()
	}
}

/** Remove reference Authorize **/
func (this *OAuth2Access) RemoveAuthorize(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Authorize)
	if toDelete.GetUUID() == this.m_authorize.GetUUID() {
		this.m_authorize = nil
		this.M_authorize = ""
	}
}

/** Previous **/
func (this *OAuth2Access) GetPrevious() string{
	return this.M_previous
}

/** Init reference Previous **/
func (this *OAuth2Access) SetPrevious(ref interface{}){
	this.NeedSave = true
	this.M_previous = ref.(string)
}

/** Remove reference Previous **/

/** AccessToken **/
func (this *OAuth2Access) GetAccessToken() string{
	return this.M_accessToken
}

/** Init reference AccessToken **/
func (this *OAuth2Access) SetAccessToken(ref interface{}){
	this.NeedSave = true
	this.M_accessToken = ref.(string)
}

/** Remove reference AccessToken **/

/** RefreshToken **/
func (this *OAuth2Access) GetRefreshToken() *OAuth2Refresh{
	return this.m_refreshToken
}

/** Init reference RefreshToken **/
func (this *OAuth2Access) SetRefreshToken(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_refreshToken = ref.(string)
	}else{
		this.m_refreshToken = ref.(*OAuth2Refresh)
		this.M_refreshToken = ref.(*OAuth2Refresh).GetUUID()
	}
}

/** Remove reference RefreshToken **/
func (this *OAuth2Access) RemoveRefreshToken(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Refresh)
	if toDelete.GetUUID() == this.m_refreshToken.GetUUID() {
		this.m_refreshToken = nil
		this.M_refreshToken = ""
	}
}

/** ExpiresIn **/
func (this *OAuth2Access) GetExpiresIn() int64{
	return this.M_expiresIn
}

/** Init reference ExpiresIn **/
func (this *OAuth2Access) SetExpiresIn(ref interface{}){
	this.NeedSave = true
	this.M_expiresIn = ref.(int64)
}

/** Remove reference ExpiresIn **/

/** Scope **/
func (this *OAuth2Access) GetScope() string{
	return this.M_scope
}

/** Init reference Scope **/
func (this *OAuth2Access) SetScope(ref interface{}){
	this.NeedSave = true
	this.M_scope = ref.(string)
}

/** Remove reference Scope **/

/** RedirectUri **/
func (this *OAuth2Access) GetRedirectUri() string{
	return this.M_redirectUri
}

/** Init reference RedirectUri **/
func (this *OAuth2Access) SetRedirectUri(ref interface{}){
	this.NeedSave = true
	this.M_redirectUri = ref.(string)
}

/** Remove reference RedirectUri **/

/** Extra **/
func (this *OAuth2Access) GetExtra() []uint8{
	return this.M_extra
}

/** Init reference Extra **/
func (this *OAuth2Access) SetExtra(ref interface{}){
	this.NeedSave = true
	this.M_extra = ref.([]uint8)
}

/** Remove reference Extra **/

/** CreatedAt **/
func (this *OAuth2Access) GetCreatedAt() int64{
	return this.M_createdAt
}

/** Init reference CreatedAt **/
func (this *OAuth2Access) SetCreatedAt(ref interface{}){
	this.NeedSave = true
	this.M_createdAt = ref.(int64)
}

/** Remove reference CreatedAt **/

/** Parent **/
func (this *OAuth2Access) GetParentPtr() *OAuth2Configuration{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Access) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*OAuth2Configuration)
		this.M_parentPtr = ref.(Configuration).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *OAuth2Access) RemoveParentPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Configuration)
	if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
		this.m_parentPtr = nil
		this.M_parentPtr = ""
	}
}
