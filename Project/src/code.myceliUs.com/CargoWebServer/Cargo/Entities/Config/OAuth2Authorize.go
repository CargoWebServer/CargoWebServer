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
	/** The relation name with the parent. **/
	ParentLnk string
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
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference Id **/

/** Client **/
func (this *OAuth2Authorize) GetClient() *OAuth2Client{
	return this.m_client
}

/** Init reference Client **/
func (this *OAuth2Authorize) SetClient(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_client != ref.(string) {
			this.M_client = ref.(string)
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_client != ref.(*OAuth2Client).GetUUID() {
			this.M_client = ref.(*OAuth2Client).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
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
			this.NeedSave = true
		}
	}
}

/** ExpiresIn **/
func (this *OAuth2Authorize) GetExpiresIn() int64{
	return this.M_expiresIn
}

/** Init reference ExpiresIn **/
func (this *OAuth2Authorize) SetExpiresIn(ref interface{}){
	if this.M_expiresIn != ref.(int64) {
		this.M_expiresIn = ref.(int64)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference ExpiresIn **/

/** Scope **/
func (this *OAuth2Authorize) GetScope() string{
	return this.M_scope
}

/** Init reference Scope **/
func (this *OAuth2Authorize) SetScope(ref interface{}){
	if this.M_scope != ref.(string) {
		this.M_scope = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference Scope **/

/** RedirectUri **/
func (this *OAuth2Authorize) GetRedirectUri() string{
	return this.M_redirectUri
}

/** Init reference RedirectUri **/
func (this *OAuth2Authorize) SetRedirectUri(ref interface{}){
	if this.M_redirectUri != ref.(string) {
		this.M_redirectUri = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference RedirectUri **/

/** State **/
func (this *OAuth2Authorize) GetState() string{
	return this.M_state
}

/** Init reference State **/
func (this *OAuth2Authorize) SetState(ref interface{}){
	if this.M_state != ref.(string) {
		this.M_state = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference State **/

/** UserData **/
func (this *OAuth2Authorize) GetUserData() *OAuth2IdToken{
	return this.m_userData
}

/** Init reference UserData **/
func (this *OAuth2Authorize) SetUserData(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_userData != ref.(string) {
			this.M_userData = ref.(string)
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_userData != ref.(*OAuth2IdToken).GetUUID() {
			this.M_userData = ref.(*OAuth2IdToken).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
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
			this.NeedSave = true
		}
	}
}

/** CreatedAt **/
func (this *OAuth2Authorize) GetCreatedAt() int64{
	return this.M_createdAt
}

/** Init reference CreatedAt **/
func (this *OAuth2Authorize) SetCreatedAt(ref interface{}){
	if this.M_createdAt != ref.(int64) {
		this.M_createdAt = ref.(int64)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference CreatedAt **/
