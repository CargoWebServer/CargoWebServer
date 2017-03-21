// +build Config

package Config

import(
"encoding/xml"
)

type OAuth2Configuration struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Configuration **/
	M_id string

	/** members of OAuth2Configuration **/
	M_authorizationExpiration int
	M_accessExpiration int
	M_tokenType string
	M_errorStatusCode int
	M_allowClientSecretInParams bool
	M_allowGetAccessRequest bool
	M_redirectUriSeparator string
	M_allowedAuthorizeTypes []string
	M_allowedAccessTypes []string
	M_clients []*OAuth2Client
	M_authorize []*OAuth2Authorize
	M_access []*OAuth2Access
	M_refresh []*OAuth2Refresh
	M_expire []*OAuth2Expires


	/** Associations **/
	m_parentPtr *Configurations
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2Configuration **/
type XsdOAuth2Configuration struct {
	XMLName xml.Name	`xml:"oauth2Configuration"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_authorizationExpiration	int	`xml:"authorizationExpiration,attr"`
	M_accessExpiration	int	`xml:"accessExpiration,attr"`
	M_tokenType	string	`xml:"tokenType,attr"`
	M_errorStatusCode	int	`xml:"errorStatusCode,attr"`
	M_allowClientSecretInParams	bool	`xml:"allowClientSecretInParams,attr"`
	M_allowGetAccessRequest	bool	`xml:"allowGetAccessRequest,attr"`
	M_requirePKCEForPublicClients	bool	`xml:"requirePKCEForPublicClients,attr"`
	M_redirectUriSeparator	string	`xml:"redirectUriSeparator,attr"`

}
/** UUID **/
func (this *OAuth2Configuration) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *OAuth2Configuration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Configuration) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** AuthorizationExpiration **/
func (this *OAuth2Configuration) GetAuthorizationExpiration() int{
	return this.M_authorizationExpiration
}

/** Init reference AuthorizationExpiration **/
func (this *OAuth2Configuration) SetAuthorizationExpiration(ref interface{}){
	this.NeedSave = true
	this.M_authorizationExpiration = ref.(int)
}

/** Remove reference AuthorizationExpiration **/

/** AccessExpiration **/
func (this *OAuth2Configuration) GetAccessExpiration() int{
	return this.M_accessExpiration
}

/** Init reference AccessExpiration **/
func (this *OAuth2Configuration) SetAccessExpiration(ref interface{}){
	this.NeedSave = true
	this.M_accessExpiration = ref.(int)
}

/** Remove reference AccessExpiration **/

/** TokenType **/
func (this *OAuth2Configuration) GetTokenType() string{
	return this.M_tokenType
}

/** Init reference TokenType **/
func (this *OAuth2Configuration) SetTokenType(ref interface{}){
	this.NeedSave = true
	this.M_tokenType = ref.(string)
}

/** Remove reference TokenType **/

/** ErrorStatusCode **/
func (this *OAuth2Configuration) GetErrorStatusCode() int{
	return this.M_errorStatusCode
}

/** Init reference ErrorStatusCode **/
func (this *OAuth2Configuration) SetErrorStatusCode(ref interface{}){
	this.NeedSave = true
	this.M_errorStatusCode = ref.(int)
}

/** Remove reference ErrorStatusCode **/

/** AllowClientSecretInParams **/
func (this *OAuth2Configuration) GetAllowClientSecretInParams() bool{
	return this.M_allowClientSecretInParams
}

/** Init reference AllowClientSecretInParams **/
func (this *OAuth2Configuration) SetAllowClientSecretInParams(ref interface{}){
	this.NeedSave = true
	this.M_allowClientSecretInParams = ref.(bool)
}

/** Remove reference AllowClientSecretInParams **/

/** AllowGetAccessRequest **/
func (this *OAuth2Configuration) GetAllowGetAccessRequest() bool{
	return this.M_allowGetAccessRequest
}

/** Init reference AllowGetAccessRequest **/
func (this *OAuth2Configuration) SetAllowGetAccessRequest(ref interface{}){
	this.NeedSave = true
	this.M_allowGetAccessRequest = ref.(bool)
}

/** Remove reference AllowGetAccessRequest **/

/** RedirectUriSeparator **/
func (this *OAuth2Configuration) GetRedirectUriSeparator() string{
	return this.M_redirectUriSeparator
}

/** Init reference RedirectUriSeparator **/
func (this *OAuth2Configuration) SetRedirectUriSeparator(ref interface{}){
	this.NeedSave = true
	this.M_redirectUriSeparator = ref.(string)
}

/** Remove reference RedirectUriSeparator **/

/** AllowedAuthorizeTypes **/
func (this *OAuth2Configuration) GetAllowedAuthorizeTypes() []string{
	return this.M_allowedAuthorizeTypes
}

/** Init reference AllowedAuthorizeTypes **/
func (this *OAuth2Configuration) SetAllowedAuthorizeTypes(ref interface{}){
	this.NeedSave = true
	isExist := false
	var allowedAuthorizeTypess []string
	for i:=0; i<len(this.M_allowedAuthorizeTypes); i++ {
		if this.M_allowedAuthorizeTypes[i] != ref.(string) {
			allowedAuthorizeTypess = append(allowedAuthorizeTypess, this.M_allowedAuthorizeTypes[i])
		} else {
			isExist = true
			allowedAuthorizeTypess = append(allowedAuthorizeTypess, ref.(string))
		}
	}
	if !isExist {
		allowedAuthorizeTypess = append(allowedAuthorizeTypess, ref.(string))
	}
	this.M_allowedAuthorizeTypes = allowedAuthorizeTypess
}

/** Remove reference AllowedAuthorizeTypes **/

/** AllowedAccessTypes **/
func (this *OAuth2Configuration) GetAllowedAccessTypes() []string{
	return this.M_allowedAccessTypes
}

/** Init reference AllowedAccessTypes **/
func (this *OAuth2Configuration) SetAllowedAccessTypes(ref interface{}){
	this.NeedSave = true
	isExist := false
	var allowedAccessTypess []string
	for i:=0; i<len(this.M_allowedAccessTypes); i++ {
		if this.M_allowedAccessTypes[i] != ref.(string) {
			allowedAccessTypess = append(allowedAccessTypess, this.M_allowedAccessTypes[i])
		} else {
			isExist = true
			allowedAccessTypess = append(allowedAccessTypess, ref.(string))
		}
	}
	if !isExist {
		allowedAccessTypess = append(allowedAccessTypess, ref.(string))
	}
	this.M_allowedAccessTypes = allowedAccessTypess
}

/** Remove reference AllowedAccessTypes **/

/** Clients **/
func (this *OAuth2Configuration) GetClients() []*OAuth2Client{
	return this.M_clients
}

/** Init reference Clients **/
func (this *OAuth2Configuration) SetClients(ref interface{}){
	this.NeedSave = true
	isExist := false
	var clientss []*OAuth2Client
	for i:=0; i<len(this.M_clients); i++ {
		if this.M_clients[i].GetUUID() != ref.(*OAuth2Client).GetUUID() {
			clientss = append(clientss, this.M_clients[i])
		} else {
			isExist = true
			clientss = append(clientss, ref.(*OAuth2Client))
		}
	}
	if !isExist {
		clientss = append(clientss, ref.(*OAuth2Client))
	}
	this.M_clients = clientss
}

/** Remove reference Clients **/
func (this *OAuth2Configuration) RemoveClients(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Client)
	clients_ := make([]*OAuth2Client, 0)
	for i := 0; i < len(this.M_clients); i++ {
		if toDelete.GetUUID() != this.M_clients[i].GetUUID() {
			clients_ = append(clients_, this.M_clients[i])
		}
	}
	this.M_clients = clients_
}

/** Authorize **/
func (this *OAuth2Configuration) GetAuthorize() []*OAuth2Authorize{
	return this.M_authorize
}

/** Init reference Authorize **/
func (this *OAuth2Configuration) SetAuthorize(ref interface{}){
	this.NeedSave = true
	isExist := false
	var authorizes []*OAuth2Authorize
	for i:=0; i<len(this.M_authorize); i++ {
		if this.M_authorize[i].GetUUID() != ref.(*OAuth2Authorize).GetUUID() {
			authorizes = append(authorizes, this.M_authorize[i])
		} else {
			isExist = true
			authorizes = append(authorizes, ref.(*OAuth2Authorize))
		}
	}
	if !isExist {
		authorizes = append(authorizes, ref.(*OAuth2Authorize))
	}
	this.M_authorize = authorizes
}

/** Remove reference Authorize **/
func (this *OAuth2Configuration) RemoveAuthorize(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Authorize)
	authorize_ := make([]*OAuth2Authorize, 0)
	for i := 0; i < len(this.M_authorize); i++ {
		if toDelete.GetUUID() != this.M_authorize[i].GetUUID() {
			authorize_ = append(authorize_, this.M_authorize[i])
		}
	}
	this.M_authorize = authorize_
}

/** Access **/
func (this *OAuth2Configuration) GetAccess() []*OAuth2Access{
	return this.M_access
}

/** Init reference Access **/
func (this *OAuth2Configuration) SetAccess(ref interface{}){
	this.NeedSave = true
	isExist := false
	var accesss []*OAuth2Access
	for i:=0; i<len(this.M_access); i++ {
		if this.M_access[i].GetUUID() != ref.(*OAuth2Access).GetUUID() {
			accesss = append(accesss, this.M_access[i])
		} else {
			isExist = true
			accesss = append(accesss, ref.(*OAuth2Access))
		}
	}
	if !isExist {
		accesss = append(accesss, ref.(*OAuth2Access))
	}
	this.M_access = accesss
}

/** Remove reference Access **/
func (this *OAuth2Configuration) RemoveAccess(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Access)
	access_ := make([]*OAuth2Access, 0)
	for i := 0; i < len(this.M_access); i++ {
		if toDelete.GetUUID() != this.M_access[i].GetUUID() {
			access_ = append(access_, this.M_access[i])
		}
	}
	this.M_access = access_
}

/** Refresh **/
func (this *OAuth2Configuration) GetRefresh() []*OAuth2Refresh{
	return this.M_refresh
}

/** Init reference Refresh **/
func (this *OAuth2Configuration) SetRefresh(ref interface{}){
	this.NeedSave = true
	isExist := false
	var refreshs []*OAuth2Refresh
	for i:=0; i<len(this.M_refresh); i++ {
		if this.M_refresh[i].GetUUID() != ref.(*OAuth2Refresh).GetUUID() {
			refreshs = append(refreshs, this.M_refresh[i])
		} else {
			isExist = true
			refreshs = append(refreshs, ref.(*OAuth2Refresh))
		}
	}
	if !isExist {
		refreshs = append(refreshs, ref.(*OAuth2Refresh))
	}
	this.M_refresh = refreshs
}

/** Remove reference Refresh **/
func (this *OAuth2Configuration) RemoveRefresh(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Refresh)
	refresh_ := make([]*OAuth2Refresh, 0)
	for i := 0; i < len(this.M_refresh); i++ {
		if toDelete.GetUUID() != this.M_refresh[i].GetUUID() {
			refresh_ = append(refresh_, this.M_refresh[i])
		}
	}
	this.M_refresh = refresh_
}

/** Expire **/
func (this *OAuth2Configuration) GetExpire() []*OAuth2Expires{
	return this.M_expire
}

/** Init reference Expire **/
func (this *OAuth2Configuration) SetExpire(ref interface{}){
	this.NeedSave = true
	isExist := false
	var expires []*OAuth2Expires
	for i:=0; i<len(this.M_expire); i++ {
		if this.M_expire[i].GetUUID() != ref.(*OAuth2Expires).GetUUID() {
			expires = append(expires, this.M_expire[i])
		} else {
			isExist = true
			expires = append(expires, ref.(*OAuth2Expires))
		}
	}
	if !isExist {
		expires = append(expires, ref.(*OAuth2Expires))
	}
	this.M_expire = expires
}

/** Remove reference Expire **/
func (this *OAuth2Configuration) RemoveExpire(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*OAuth2Expires)
	expire_ := make([]*OAuth2Expires, 0)
	for i := 0; i < len(this.M_expire); i++ {
		if toDelete.GetUUID() != this.M_expire[i].GetUUID() {
			expire_ = append(expire_, this.M_expire[i])
		}
	}
	this.M_expire = expire_
}

/** Parent **/
func (this *OAuth2Configuration) GetParentPtr() *Configurations{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Configuration) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*Configurations)
		this.M_parentPtr = ref.(*Configurations).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *OAuth2Configuration) RemoveParentPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Configurations)
	if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
		this.m_parentPtr = nil
		this.M_parentPtr = ""
	}
}
