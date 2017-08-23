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
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of OAuth2Access **/
	M_id string
	m_client *OAuth2Client
	/** If the ref is a string and not an object **/
	M_client string
	M_authorize string
	M_previous string
	m_refreshToken *OAuth2Refresh
	/** If the ref is a string and not an object **/
	M_refreshToken string
	M_expiresIn int64
	M_scope string
	M_redirectUri string
	m_userData *OAuth2IdToken
	/** If the ref is a string and not an object **/
	M_userData string
	M_createdAt int64


	/** Associations **/
	m_parentPtr *OAuth2Configuration
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2Access **/
type XsdOAuth2Access struct {
	XMLName xml.Name	`xml:"oauth2Access"`
	M_id	string	`xml:"id,attr"`
	M_authorize	string	`xml:"authorize,attr"`
	M_previous	string	`xml:"previous,attr"`
	M_expiresIn 	int64	`xml:"expiresIn ,attr"`
	M_scope	string	`xml:"scope,attr"`
	M_redirectUri	string	`xml:"redirectUri,attr"`
	M_tokenUri	string	`xml:"tokenUri,attr"`
	M_authorizationUri	string	`xml:"authorizationUri,attr"`
	M_createdAt 	int64	`xml:"createdAt ,attr"`

}
/** UUID **/
func (this *OAuth2Access) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *OAuth2Access) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Access) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

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
	toDelete := ref.(*OAuth2Client)
	if this.m_client!= nil {
		if toDelete.GetUUID() == this.m_client.GetUUID() {
			this.m_client = nil
			this.M_client = ""
		}else{
			this.NeedSave = true
		}
	}
}

/** Authorize **/
func (this *OAuth2Access) GetAuthorize() string{
	return this.M_authorize
}

/** Init reference Authorize **/
func (this *OAuth2Access) SetAuthorize(ref interface{}){
	this.NeedSave = true
	this.M_authorize = ref.(string)
}

/** Remove reference Authorize **/

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
	toDelete := ref.(*OAuth2Refresh)
	if this.m_refreshToken!= nil {
		if toDelete.GetUUID() == this.m_refreshToken.GetUUID() {
			this.m_refreshToken = nil
			this.M_refreshToken = ""
		}else{
			this.NeedSave = true
		}
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

/** UserData **/
func (this *OAuth2Access) GetUserData() *OAuth2IdToken{
	return this.m_userData
}

/** Init reference UserData **/
func (this *OAuth2Access) SetUserData(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_userData = ref.(string)
	}else{
		this.m_userData = ref.(*OAuth2IdToken)
		this.M_userData = ref.(*OAuth2IdToken).GetUUID()
	}
}

/** Remove reference UserData **/
func (this *OAuth2Access) RemoveUserData(ref interface{}){
	toDelete := ref.(*OAuth2IdToken)
	if this.m_userData!= nil {
		if toDelete.GetUUID() == this.m_userData.GetUUID() {
			this.m_userData = nil
			this.M_userData = ""
		}else{
			this.NeedSave = true
		}
	}
}

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
	toDelete := ref.(Configuration)
	if this.m_parentPtr!= nil {
		if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
		}else{
			this.NeedSave = true
		}
	}
}
