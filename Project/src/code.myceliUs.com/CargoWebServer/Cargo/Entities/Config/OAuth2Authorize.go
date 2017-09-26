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
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of OAuth2Authorize **/
	M_id string
	m_client *OAuth2Client
	/** If the ref is a string and not an object **/
	M_client string
	M_expiresIn int64
	M_scope string
	M_redirectUri string
	M_state string
	m_userData *OAuth2IdToken
	/** If the ref is a string and not an object **/
	M_userData string
	M_createdAt int64

}

/** Xml parser for OAuth2Authorize **/
type XsdOAuth2Authorize struct {
	XMLName xml.Name	`xml:"oauth2Authorize"`
	M_id	string	`xml:"id,attr"`
	M_expiresIn 	int64	`xml:"expiresIn ,attr"`
	M_scope	string	`xml:"scope,attr"`
	M_redirectUri	string	`xml:"redirectUri,attr"`
	M_state	string	`xml:"state,attr"`
	M_createdAt 	int64	`xml:"createdAt ,attr"`

}
/** UUID **/
func (this *OAuth2Authorize) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *OAuth2Authorize) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Authorize) SetId(ref interface{}){
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Client **/
func (this *OAuth2Authorize) GetClient() *OAuth2Client{
	return this.m_client
}

/** Init reference Client **/
func (this *OAuth2Authorize) SetClient(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_client = ref.(string)
	}else{
		this.M_client = ref.(*OAuth2Client).GetUUID()
		this.m_client = ref.(*OAuth2Client)
	}
}

/** Remove reference Client **/
func (this *OAuth2Authorize) RemoveClient(ref interface{}){
	toDelete := ref.(*OAuth2Client)
	if this.m_client!= nil {
		if toDelete.GetUUID() == this.m_client.GetUUID() {
			this.m_client = nil
			this.M_client = ""
		}
	}
}

/** ExpiresIn **/
func (this *OAuth2Authorize) GetExpiresIn() int64{
	return this.M_expiresIn
}

/** Init reference ExpiresIn **/
func (this *OAuth2Authorize) SetExpiresIn(ref interface{}){
	this.M_expiresIn = ref.(int64)
}

/** Remove reference ExpiresIn **/

/** Scope **/
func (this *OAuth2Authorize) GetScope() string{
	return this.M_scope
}

/** Init reference Scope **/
func (this *OAuth2Authorize) SetScope(ref interface{}){
	this.M_scope = ref.(string)
}

/** Remove reference Scope **/

/** RedirectUri **/
func (this *OAuth2Authorize) GetRedirectUri() string{
	return this.M_redirectUri
}

/** Init reference RedirectUri **/
func (this *OAuth2Authorize) SetRedirectUri(ref interface{}){
	this.M_redirectUri = ref.(string)
}

/** Remove reference RedirectUri **/

/** State **/
func (this *OAuth2Authorize) GetState() string{
	return this.M_state
}

/** Init reference State **/
func (this *OAuth2Authorize) SetState(ref interface{}){
	this.M_state = ref.(string)
}

/** Remove reference State **/

/** UserData **/
func (this *OAuth2Authorize) GetUserData() *OAuth2IdToken{
	return this.m_userData
}

/** Init reference UserData **/
func (this *OAuth2Authorize) SetUserData(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_userData = ref.(string)
	}else{
		this.M_userData = ref.(*OAuth2IdToken).GetUUID()
		this.m_userData = ref.(*OAuth2IdToken)
	}
}

/** Remove reference UserData **/
func (this *OAuth2Authorize) RemoveUserData(ref interface{}){
	toDelete := ref.(*OAuth2IdToken)
	if this.m_userData!= nil {
		if toDelete.GetUUID() == this.m_userData.GetUUID() {
			this.m_userData = nil
			this.M_userData = ""
		}
	}
}

/** CreatedAt **/
func (this *OAuth2Authorize) GetCreatedAt() int64{
	return this.M_createdAt
}

/** Init reference CreatedAt **/
func (this *OAuth2Authorize) SetCreatedAt(ref interface{}){
	this.M_createdAt = ref.(int64)
}

/** Remove reference CreatedAt **/
