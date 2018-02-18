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
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

	/** members of Configuration **/
	M_id string

	/** members of OAuth2Configuration **/
	M_authorizationExpiration int
	M_accessExpiration int64
	M_tokenType string
	M_errorStatusCode int
	M_allowClientSecretInParams bool
	M_allowGetAccessRequest bool
	M_redirectUriSeparator string
	M_privateKey string
	M_allowedAuthorizeTypes []string
	M_allowedAccessTypes []string
	M_clients []*OAuth2Client
	M_authorize []*OAuth2Authorize
	M_access []*OAuth2Access
	M_ids []*OAuth2IdToken
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
	M_accessExpiration	int64	`xml:"accessExpiration,attr"`
	M_tokenType	string	`xml:"tokenType,attr"`
	M_errorStatusCode	int	`xml:"errorStatusCode,attr"`
	M_allowClientSecretInParams	bool	`xml:"allowClientSecretInParams,attr"`
	M_allowGetAccessRequest	bool	`xml:"allowGetAccessRequest,attr"`
	M_requirePKCEForPublicClients	bool	`xml:"requirePKCEForPublicClients,attr"`
	M_redirectUriSeparator	string	`xml:"redirectUriSeparator,attr"`
	M_privateKey	string	`xml:"privateKey,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2Configuration) GetUuid() string{
	return this.UUID
}
func (this *OAuth2Configuration) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2Configuration) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2Configuration) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2Configuration"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2Configuration) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2Configuration) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Configuration) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2Configuration) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *OAuth2Configuration) IsNeedSave() bool{
	return this.NeedSave
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Configuration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *OAuth2Configuration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Configuration) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** AuthorizationExpiration **/
func (this *OAuth2Configuration) GetAuthorizationExpiration() int{
	return this.M_authorizationExpiration
}

/** Init reference AuthorizationExpiration **/
func (this *OAuth2Configuration) SetAuthorizationExpiration(ref interface{}){
	if this.M_authorizationExpiration != ref.(int) {
		this.M_authorizationExpiration = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference AuthorizationExpiration **/

/** AccessExpiration **/
func (this *OAuth2Configuration) GetAccessExpiration() int64{
	return this.M_accessExpiration
}

/** Init reference AccessExpiration **/
func (this *OAuth2Configuration) SetAccessExpiration(ref interface{}){
	if this.M_accessExpiration != ref.(int64) {
		this.M_accessExpiration = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference AccessExpiration **/

/** TokenType **/
func (this *OAuth2Configuration) GetTokenType() string{
	return this.M_tokenType
}

/** Init reference TokenType **/
func (this *OAuth2Configuration) SetTokenType(ref interface{}){
	if this.M_tokenType != ref.(string) {
		this.M_tokenType = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference TokenType **/

/** ErrorStatusCode **/
func (this *OAuth2Configuration) GetErrorStatusCode() int{
	return this.M_errorStatusCode
}

/** Init reference ErrorStatusCode **/
func (this *OAuth2Configuration) SetErrorStatusCode(ref interface{}){
	if this.M_errorStatusCode != ref.(int) {
		this.M_errorStatusCode = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference ErrorStatusCode **/

/** AllowClientSecretInParams **/
func (this *OAuth2Configuration) GetAllowClientSecretInParams() bool{
	return this.M_allowClientSecretInParams
}

/** Init reference AllowClientSecretInParams **/
func (this *OAuth2Configuration) SetAllowClientSecretInParams(ref interface{}){
	if this.M_allowClientSecretInParams != ref.(bool) {
		this.M_allowClientSecretInParams = ref.(bool)
		this.NeedSave = true
	}
}

/** Remove reference AllowClientSecretInParams **/

/** AllowGetAccessRequest **/
func (this *OAuth2Configuration) GetAllowGetAccessRequest() bool{
	return this.M_allowGetAccessRequest
}

/** Init reference AllowGetAccessRequest **/
func (this *OAuth2Configuration) SetAllowGetAccessRequest(ref interface{}){
	if this.M_allowGetAccessRequest != ref.(bool) {
		this.M_allowGetAccessRequest = ref.(bool)
		this.NeedSave = true
	}
}

/** Remove reference AllowGetAccessRequest **/

/** RedirectUriSeparator **/
func (this *OAuth2Configuration) GetRedirectUriSeparator() string{
	return this.M_redirectUriSeparator
}

/** Init reference RedirectUriSeparator **/
func (this *OAuth2Configuration) SetRedirectUriSeparator(ref interface{}){
	if this.M_redirectUriSeparator != ref.(string) {
		this.M_redirectUriSeparator = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference RedirectUriSeparator **/

/** PrivateKey **/
func (this *OAuth2Configuration) GetPrivateKey() string{
	return this.M_privateKey
}

/** Init reference PrivateKey **/
func (this *OAuth2Configuration) SetPrivateKey(ref interface{}){
	if this.M_privateKey != ref.(string) {
		this.M_privateKey = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference PrivateKey **/

/** AllowedAuthorizeTypes **/
func (this *OAuth2Configuration) GetAllowedAuthorizeTypes() []string{
	return this.M_allowedAuthorizeTypes
}

/** Init reference AllowedAuthorizeTypes **/
func (this *OAuth2Configuration) SetAllowedAuthorizeTypes(ref interface{}){
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
		this.NeedSave = true
		this.M_allowedAuthorizeTypes = allowedAuthorizeTypess
	}
}

/** Remove reference AllowedAuthorizeTypes **/

/** AllowedAccessTypes **/
func (this *OAuth2Configuration) GetAllowedAccessTypes() []string{
	return this.M_allowedAccessTypes
}

/** Init reference AllowedAccessTypes **/
func (this *OAuth2Configuration) SetAllowedAccessTypes(ref interface{}){
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
		this.NeedSave = true
		this.M_allowedAccessTypes = allowedAccessTypess
	}
}

/** Remove reference AllowedAccessTypes **/

/** Clients **/
func (this *OAuth2Configuration) GetClients() []*OAuth2Client{
	return this.M_clients
}

/** Init reference Clients **/
func (this *OAuth2Configuration) SetClients(ref interface{}){
	isExist := false
	var clientss []*OAuth2Client
	for i:=0; i<len(this.M_clients); i++ {
		if this.M_clients[i].GetUuid() != ref.(*OAuth2Client).GetUuid() {
			clientss = append(clientss, this.M_clients[i])
		} else {
			isExist = true
			clientss = append(clientss, ref.(*OAuth2Client))
		}
	}
	if !isExist {
		clientss = append(clientss, ref.(*OAuth2Client))
		this.NeedSave = true
		this.M_clients = clientss
	}
}

/** Remove reference Clients **/
func (this *OAuth2Configuration) RemoveClients(ref interface{}){
	toDelete := ref.(*OAuth2Client)
	clients_ := make([]*OAuth2Client, 0)
	for i := 0; i < len(this.M_clients); i++ {
		if toDelete.GetUuid() != this.M_clients[i].GetUuid() {
			clients_ = append(clients_, this.M_clients[i])
		}else{
			this.NeedSave = true
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
	isExist := false
	var authorizes []*OAuth2Authorize
	for i:=0; i<len(this.M_authorize); i++ {
		if this.M_authorize[i].GetUuid() != ref.(*OAuth2Authorize).GetUuid() {
			authorizes = append(authorizes, this.M_authorize[i])
		} else {
			isExist = true
			authorizes = append(authorizes, ref.(*OAuth2Authorize))
		}
	}
	if !isExist {
		authorizes = append(authorizes, ref.(*OAuth2Authorize))
		this.NeedSave = true
		this.M_authorize = authorizes
	}
}

/** Remove reference Authorize **/
func (this *OAuth2Configuration) RemoveAuthorize(ref interface{}){
	toDelete := ref.(*OAuth2Authorize)
	authorize_ := make([]*OAuth2Authorize, 0)
	for i := 0; i < len(this.M_authorize); i++ {
		if toDelete.GetUuid() != this.M_authorize[i].GetUuid() {
			authorize_ = append(authorize_, this.M_authorize[i])
		}else{
			this.NeedSave = true
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
	isExist := false
	var accesss []*OAuth2Access
	for i:=0; i<len(this.M_access); i++ {
		if this.M_access[i].GetUuid() != ref.(*OAuth2Access).GetUuid() {
			accesss = append(accesss, this.M_access[i])
		} else {
			isExist = true
			accesss = append(accesss, ref.(*OAuth2Access))
		}
	}
	if !isExist {
		accesss = append(accesss, ref.(*OAuth2Access))
		this.NeedSave = true
		this.M_access = accesss
	}
}

/** Remove reference Access **/
func (this *OAuth2Configuration) RemoveAccess(ref interface{}){
	toDelete := ref.(*OAuth2Access)
	access_ := make([]*OAuth2Access, 0)
	for i := 0; i < len(this.M_access); i++ {
		if toDelete.GetUuid() != this.M_access[i].GetUuid() {
			access_ = append(access_, this.M_access[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_access = access_
}

/** Ids **/
func (this *OAuth2Configuration) GetIds() []*OAuth2IdToken{
	return this.M_ids
}

/** Init reference Ids **/
func (this *OAuth2Configuration) SetIds(ref interface{}){
	isExist := false
	var idss []*OAuth2IdToken
	for i:=0; i<len(this.M_ids); i++ {
		if this.M_ids[i].GetUuid() != ref.(*OAuth2IdToken).GetUuid() {
			idss = append(idss, this.M_ids[i])
		} else {
			isExist = true
			idss = append(idss, ref.(*OAuth2IdToken))
		}
	}
	if !isExist {
		idss = append(idss, ref.(*OAuth2IdToken))
		this.NeedSave = true
		this.M_ids = idss
	}
}

/** Remove reference Ids **/
func (this *OAuth2Configuration) RemoveIds(ref interface{}){
	toDelete := ref.(*OAuth2IdToken)
	ids_ := make([]*OAuth2IdToken, 0)
	for i := 0; i < len(this.M_ids); i++ {
		if toDelete.GetUuid() != this.M_ids[i].GetUuid() {
			ids_ = append(ids_, this.M_ids[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_ids = ids_
}

/** Refresh **/
func (this *OAuth2Configuration) GetRefresh() []*OAuth2Refresh{
	return this.M_refresh
}

/** Init reference Refresh **/
func (this *OAuth2Configuration) SetRefresh(ref interface{}){
	isExist := false
	var refreshs []*OAuth2Refresh
	for i:=0; i<len(this.M_refresh); i++ {
		if this.M_refresh[i].GetUuid() != ref.(*OAuth2Refresh).GetUuid() {
			refreshs = append(refreshs, this.M_refresh[i])
		} else {
			isExist = true
			refreshs = append(refreshs, ref.(*OAuth2Refresh))
		}
	}
	if !isExist {
		refreshs = append(refreshs, ref.(*OAuth2Refresh))
		this.NeedSave = true
		this.M_refresh = refreshs
	}
}

/** Remove reference Refresh **/
func (this *OAuth2Configuration) RemoveRefresh(ref interface{}){
	toDelete := ref.(*OAuth2Refresh)
	refresh_ := make([]*OAuth2Refresh, 0)
	for i := 0; i < len(this.M_refresh); i++ {
		if toDelete.GetUuid() != this.M_refresh[i].GetUuid() {
			refresh_ = append(refresh_, this.M_refresh[i])
		}else{
			this.NeedSave = true
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
	isExist := false
	var expires []*OAuth2Expires
	for i:=0; i<len(this.M_expire); i++ {
		if this.M_expire[i].GetUuid() != ref.(*OAuth2Expires).GetUuid() {
			expires = append(expires, this.M_expire[i])
		} else {
			isExist = true
			expires = append(expires, ref.(*OAuth2Expires))
		}
	}
	if !isExist {
		expires = append(expires, ref.(*OAuth2Expires))
		this.NeedSave = true
		this.M_expire = expires
	}
}

/** Remove reference Expire **/
func (this *OAuth2Configuration) RemoveExpire(ref interface{}){
	toDelete := ref.(*OAuth2Expires)
	expire_ := make([]*OAuth2Expires, 0)
	for i := 0; i < len(this.M_expire); i++ {
		if toDelete.GetUuid() != this.M_expire[i].GetUuid() {
			expire_ = append(expire_, this.M_expire[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_expire = expire_
}

/** Parent **/
func (this *OAuth2Configuration) GetParentPtr() *Configurations{
	if this.m_parentPtr == nil {
		entity, err := this.getEntityByUuid(this.M_parentPtr)
		if err == nil {
			this.m_parentPtr = entity.(*Configurations)
		}
	}
	return this.m_parentPtr
}
func (this *OAuth2Configuration) GetParentPtrStr() string{
	return this.M_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Configuration) SetParentPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_parentPtr != ref.(string) {
			this.M_parentPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_parentPtr != ref.(*Configurations).GetUuid() {
			this.M_parentPtr = ref.(*Configurations).GetUuid()
			this.NeedSave = true
		}
		this.m_parentPtr = ref.(*Configurations)
	}
}

/** Remove reference Parent **/
func (this *OAuth2Configuration) RemoveParentPtr(ref interface{}){
	toDelete := ref.(*Configurations)
	if this.m_parentPtr!= nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
