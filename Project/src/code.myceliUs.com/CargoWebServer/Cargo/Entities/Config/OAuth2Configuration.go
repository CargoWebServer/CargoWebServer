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
	M_clients []string
	M_authorize []string
	M_access []string
	M_ids []string
	M_refresh []string
	M_expire []string


	/** Associations **/
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
func (this *OAuth2Configuration) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Configuration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

func (this *OAuth2Configuration) GetId()string{
	return this.M_id
}

func (this *OAuth2Configuration) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *OAuth2Configuration) GetAuthorizationExpiration()int{
	return this.M_authorizationExpiration
}

func (this *OAuth2Configuration) SetAuthorizationExpiration(val int){
	this.NeedSave = this.M_authorizationExpiration== val
	this.M_authorizationExpiration= val
}


func (this *OAuth2Configuration) GetAccessExpiration()int64{
	return this.M_accessExpiration
}

func (this *OAuth2Configuration) SetAccessExpiration(val int64){
	this.NeedSave = this.M_accessExpiration== val
	this.M_accessExpiration= val
}


func (this *OAuth2Configuration) GetTokenType()string{
	return this.M_tokenType
}

func (this *OAuth2Configuration) SetTokenType(val string){
	this.NeedSave = this.M_tokenType== val
	this.M_tokenType= val
}


func (this *OAuth2Configuration) GetErrorStatusCode()int{
	return this.M_errorStatusCode
}

func (this *OAuth2Configuration) SetErrorStatusCode(val int){
	this.NeedSave = this.M_errorStatusCode== val
	this.M_errorStatusCode= val
}


func (this *OAuth2Configuration) IsAllowClientSecretInParams()bool{
	return this.M_allowClientSecretInParams
}

func (this *OAuth2Configuration) SetAllowClientSecretInParams(val bool){
	this.NeedSave = this.M_allowClientSecretInParams== val
	this.M_allowClientSecretInParams= val
}


func (this *OAuth2Configuration) IsAllowGetAccessRequest()bool{
	return this.M_allowGetAccessRequest
}

func (this *OAuth2Configuration) SetAllowGetAccessRequest(val bool){
	this.NeedSave = this.M_allowGetAccessRequest== val
	this.M_allowGetAccessRequest= val
}


func (this *OAuth2Configuration) GetRedirectUriSeparator()string{
	return this.M_redirectUriSeparator
}

func (this *OAuth2Configuration) SetRedirectUriSeparator(val string){
	this.NeedSave = this.M_redirectUriSeparator== val
	this.M_redirectUriSeparator= val
}


func (this *OAuth2Configuration) GetPrivateKey()string{
	return this.M_privateKey
}

func (this *OAuth2Configuration) SetPrivateKey(val string){
	this.NeedSave = this.M_privateKey== val
	this.M_privateKey= val
}


func (this *OAuth2Configuration) GetAllowedAuthorizeTypes()[]string{
	return this.M_allowedAuthorizeTypes
}

func (this *OAuth2Configuration) SetAllowedAuthorizeTypes(val []string){
	this.M_allowedAuthorizeTypes= val
}

func (this *OAuth2Configuration) AppendAllowedAuthorizeTypes(val string){
	this.M_allowedAuthorizeTypes=append(this.M_allowedAuthorizeTypes, val)
	this.NeedSave= true
}


func (this *OAuth2Configuration) GetAllowedAccessTypes()[]string{
	return this.M_allowedAccessTypes
}

func (this *OAuth2Configuration) SetAllowedAccessTypes(val []string){
	this.M_allowedAccessTypes= val
}

func (this *OAuth2Configuration) AppendAllowedAccessTypes(val string){
	this.M_allowedAccessTypes=append(this.M_allowedAccessTypes, val)
	this.NeedSave= true
}


func (this *OAuth2Configuration) GetClients()[]*OAuth2Client{
	clients := make([]*OAuth2Client, 0)
	for i := 0; i < len(this.M_clients); i++ {
		entity, err := this.getEntityByUuid(this.M_clients[i])
		if err == nil {
			clients = append(clients, entity.(*OAuth2Client))
		}
	}
	return clients
}

func (this *OAuth2Configuration) SetClients(val []*OAuth2Client){
	this.M_clients= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_clients=append(this.M_clients, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *OAuth2Configuration) AppendClients(val *OAuth2Client){
	for i:=0; i < len(this.M_clients); i++{
		if this.M_clients[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_clients = append(this.M_clients, val.GetUuid())
}

func (this *OAuth2Configuration) RemoveClients(val *OAuth2Client){
	clients := make([]string,0)
	for i:=0; i < len(this.M_clients); i++{
		if this.M_clients[i] != val.GetUuid() {
			clients = append(clients, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_clients = clients
}


func (this *OAuth2Configuration) GetAuthorize()[]*OAuth2Authorize{
	authorize := make([]*OAuth2Authorize, 0)
	for i := 0; i < len(this.M_authorize); i++ {
		entity, err := this.getEntityByUuid(this.M_authorize[i])
		if err == nil {
			authorize = append(authorize, entity.(*OAuth2Authorize))
		}
	}
	return authorize
}

func (this *OAuth2Configuration) SetAuthorize(val []*OAuth2Authorize){
	this.M_authorize= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_authorize=append(this.M_authorize, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *OAuth2Configuration) AppendAuthorize(val *OAuth2Authorize){
	for i:=0; i < len(this.M_authorize); i++{
		if this.M_authorize[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_authorize = append(this.M_authorize, val.GetUuid())
}

func (this *OAuth2Configuration) RemoveAuthorize(val *OAuth2Authorize){
	authorize := make([]string,0)
	for i:=0; i < len(this.M_authorize); i++{
		if this.M_authorize[i] != val.GetUuid() {
			authorize = append(authorize, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_authorize = authorize
}


func (this *OAuth2Configuration) GetAccess()[]*OAuth2Access{
	access := make([]*OAuth2Access, 0)
	for i := 0; i < len(this.M_access); i++ {
		entity, err := this.getEntityByUuid(this.M_access[i])
		if err == nil {
			access = append(access, entity.(*OAuth2Access))
		}
	}
	return access
}

func (this *OAuth2Configuration) SetAccess(val []*OAuth2Access){
	this.M_access= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_access=append(this.M_access, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *OAuth2Configuration) AppendAccess(val *OAuth2Access){
	for i:=0; i < len(this.M_access); i++{
		if this.M_access[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_access = append(this.M_access, val.GetUuid())
}

func (this *OAuth2Configuration) RemoveAccess(val *OAuth2Access){
	access := make([]string,0)
	for i:=0; i < len(this.M_access); i++{
		if this.M_access[i] != val.GetUuid() {
			access = append(access, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_access = access
}


func (this *OAuth2Configuration) GetIds()[]*OAuth2IdToken{
	ids := make([]*OAuth2IdToken, 0)
	for i := 0; i < len(this.M_ids); i++ {
		entity, err := this.getEntityByUuid(this.M_ids[i])
		if err == nil {
			ids = append(ids, entity.(*OAuth2IdToken))
		}
	}
	return ids
}

func (this *OAuth2Configuration) SetIds(val []*OAuth2IdToken){
	this.M_ids= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_ids=append(this.M_ids, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *OAuth2Configuration) AppendIds(val *OAuth2IdToken){
	for i:=0; i < len(this.M_ids); i++{
		if this.M_ids[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_ids = append(this.M_ids, val.GetUuid())
}

func (this *OAuth2Configuration) RemoveIds(val *OAuth2IdToken){
	ids := make([]string,0)
	for i:=0; i < len(this.M_ids); i++{
		if this.M_ids[i] != val.GetUuid() {
			ids = append(ids, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_ids = ids
}


func (this *OAuth2Configuration) GetRefresh()[]*OAuth2Refresh{
	refresh := make([]*OAuth2Refresh, 0)
	for i := 0; i < len(this.M_refresh); i++ {
		entity, err := this.getEntityByUuid(this.M_refresh[i])
		if err == nil {
			refresh = append(refresh, entity.(*OAuth2Refresh))
		}
	}
	return refresh
}

func (this *OAuth2Configuration) SetRefresh(val []*OAuth2Refresh){
	this.M_refresh= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_refresh=append(this.M_refresh, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *OAuth2Configuration) AppendRefresh(val *OAuth2Refresh){
	for i:=0; i < len(this.M_refresh); i++{
		if this.M_refresh[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_refresh = append(this.M_refresh, val.GetUuid())
}

func (this *OAuth2Configuration) RemoveRefresh(val *OAuth2Refresh){
	refresh := make([]string,0)
	for i:=0; i < len(this.M_refresh); i++{
		if this.M_refresh[i] != val.GetUuid() {
			refresh = append(refresh, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_refresh = refresh
}


func (this *OAuth2Configuration) GetExpire()[]*OAuth2Expires{
	expire := make([]*OAuth2Expires, 0)
	for i := 0; i < len(this.M_expire); i++ {
		entity, err := this.getEntityByUuid(this.M_expire[i])
		if err == nil {
			expire = append(expire, entity.(*OAuth2Expires))
		}
	}
	return expire
}

func (this *OAuth2Configuration) SetExpire(val []*OAuth2Expires){
	this.M_expire= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_expire=append(this.M_expire, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *OAuth2Configuration) AppendExpire(val *OAuth2Expires){
	for i:=0; i < len(this.M_expire); i++{
		if this.M_expire[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_expire = append(this.M_expire, val.GetUuid())
}

func (this *OAuth2Configuration) RemoveExpire(val *OAuth2Expires){
	expire := make([]string,0)
	for i:=0; i < len(this.M_expire); i++{
		if this.M_expire[i] != val.GetUuid() {
			expire = append(expire, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_expire = expire
}


func (this *OAuth2Configuration) GetParentPtr()*Configurations{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*Configurations)
	}
	return nil
}

func (this *OAuth2Configuration) SetParentPtr(val *Configurations){
	this.NeedSave = this.M_parentPtr != val.GetUuid()
	this.M_parentPtr= val.GetUuid()
}

func (this *OAuth2Configuration) ResetParentPtr(){
	this.M_parentPtr= ""
}

