// +build Config

package Config

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
	"strings"
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
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

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
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
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

func (this *OAuth2Configuration) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Configuration) GetChilds() []interface{}{
	var childs []interface{}
	var child interface{}
	var err error
	for i:=0; i < len(this.M_clients); i++ {
		child, err = this.getEntityByUuid( this.M_clients[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	for i:=0; i < len(this.M_authorize); i++ {
		child, err = this.getEntityByUuid( this.M_authorize[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	for i:=0; i < len(this.M_access); i++ {
		child, err = this.getEntityByUuid( this.M_access[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	for i:=0; i < len(this.M_ids); i++ {
		child, err = this.getEntityByUuid( this.M_ids[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	for i:=0; i < len(this.M_refresh); i++ {
		child, err = this.getEntityByUuid( this.M_refresh[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	for i:=0; i < len(this.M_expire); i++ {
		child, err = this.getEntityByUuid( this.M_expire[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	return childs
}
/** Return the list of all childs uuid **/
func (this *OAuth2Configuration) GetChildsUuid() []string{
	var childs []string
	childs = append( childs, this.M_clients...)
	childs = append( childs, this.M_authorize...)
	childs = append( childs, this.M_access...)
	childs = append( childs, this.M_ids...)
	childs = append( childs, this.M_refresh...)
	childs = append( childs, this.M_expire...)
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Configuration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *OAuth2Configuration) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *OAuth2Configuration) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *OAuth2Configuration) GetId()string{
	return this.M_id
}

func (this *OAuth2Configuration) SetId(val string){
	this.M_id= val
}




func (this *OAuth2Configuration) GetAuthorizationExpiration()int{
	return this.M_authorizationExpiration
}

func (this *OAuth2Configuration) SetAuthorizationExpiration(val int){
	this.M_authorizationExpiration= val
}




func (this *OAuth2Configuration) GetAccessExpiration()int64{
	return this.M_accessExpiration
}

func (this *OAuth2Configuration) SetAccessExpiration(val int64){
	this.M_accessExpiration= val
}




func (this *OAuth2Configuration) GetTokenType()string{
	return this.M_tokenType
}

func (this *OAuth2Configuration) SetTokenType(val string){
	this.M_tokenType= val
}




func (this *OAuth2Configuration) GetErrorStatusCode()int{
	return this.M_errorStatusCode
}

func (this *OAuth2Configuration) SetErrorStatusCode(val int){
	this.M_errorStatusCode= val
}




func (this *OAuth2Configuration) IsAllowClientSecretInParams()bool{
	return this.M_allowClientSecretInParams
}

func (this *OAuth2Configuration) SetAllowClientSecretInParams(val bool){
	this.M_allowClientSecretInParams= val
}




func (this *OAuth2Configuration) IsAllowGetAccessRequest()bool{
	return this.M_allowGetAccessRequest
}

func (this *OAuth2Configuration) SetAllowGetAccessRequest(val bool){
	this.M_allowGetAccessRequest= val
}




func (this *OAuth2Configuration) GetRedirectUriSeparator()string{
	return this.M_redirectUriSeparator
}

func (this *OAuth2Configuration) SetRedirectUriSeparator(val string){
	this.M_redirectUriSeparator= val
}




func (this *OAuth2Configuration) GetPrivateKey()string{
	return this.M_privateKey
}

func (this *OAuth2Configuration) SetPrivateKey(val string){
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
}



func (this *OAuth2Configuration) GetAllowedAccessTypes()[]string{
	return this.M_allowedAccessTypes
}

func (this *OAuth2Configuration) SetAllowedAccessTypes(val []string){
	this.M_allowedAccessTypes= val
}


func (this *OAuth2Configuration) AppendAllowedAccessTypes(val string){
	this.M_allowedAccessTypes=append(this.M_allowedAccessTypes, val)
}



func (this *OAuth2Configuration) GetClients()[]*OAuth2Client{
	values := make([]*OAuth2Client, 0)
	for i := 0; i < len(this.M_clients); i++ {
		entity, err := this.getEntityByUuid(this.M_clients[i])
		if err == nil {
			values = append( values, entity.(*OAuth2Client))
		}
	}
	return values
}

func (this *OAuth2Configuration) SetClients(val []*OAuth2Client){
	this.M_clients= make([]string,0)
	for i:=0; i < len(val); i++{
		if len(val[i].GetParentUuid()) > 0  &&  len(val[i].GetParentLnk()) > 0 {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_clients")
		this.M_clients=append(this.M_clients, val[i].GetUuid())
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *OAuth2Configuration) AppendClients(val *OAuth2Client){
	for i:=0; i < len(this.M_clients); i++{
		if this.M_clients[i] == val.GetUuid() {
			return
		}
	}
	if len(val.GetParentUuid()) > 0 &&  len(val.GetParentLnk()) > 0 {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_clients")
  this.setEntity(val)
	this.M_clients = append(this.M_clients, val.GetUuid())
	this.setEntity(this)
}

func (this *OAuth2Configuration) RemoveClients(val *OAuth2Client){
	values := make([]string,0)
	for i:=0; i < len(this.M_clients); i++{
		if this.M_clients[i] != val.GetUuid() {
			values = append(values, this.M_clients[i])
		}
	}
	this.M_clients = values
	this.setEntity(this)
}


func (this *OAuth2Configuration) GetAuthorize()[]*OAuth2Authorize{
	values := make([]*OAuth2Authorize, 0)
	for i := 0; i < len(this.M_authorize); i++ {
		entity, err := this.getEntityByUuid(this.M_authorize[i])
		if err == nil {
			values = append( values, entity.(*OAuth2Authorize))
		}
	}
	return values
}

func (this *OAuth2Configuration) SetAuthorize(val []*OAuth2Authorize){
	this.M_authorize= make([]string,0)
	for i:=0; i < len(val); i++{
		if len(val[i].GetParentUuid()) > 0  &&  len(val[i].GetParentLnk()) > 0 {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_authorize")
		this.M_authorize=append(this.M_authorize, val[i].GetUuid())
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *OAuth2Configuration) AppendAuthorize(val *OAuth2Authorize){
	for i:=0; i < len(this.M_authorize); i++{
		if this.M_authorize[i] == val.GetUuid() {
			return
		}
	}
	if len(val.GetParentUuid()) > 0 &&  len(val.GetParentLnk()) > 0 {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_authorize")
  this.setEntity(val)
	this.M_authorize = append(this.M_authorize, val.GetUuid())
	this.setEntity(this)
}

func (this *OAuth2Configuration) RemoveAuthorize(val *OAuth2Authorize){
	values := make([]string,0)
	for i:=0; i < len(this.M_authorize); i++{
		if this.M_authorize[i] != val.GetUuid() {
			values = append(values, this.M_authorize[i])
		}
	}
	this.M_authorize = values
	this.setEntity(this)
}


func (this *OAuth2Configuration) GetAccess()[]*OAuth2Access{
	values := make([]*OAuth2Access, 0)
	for i := 0; i < len(this.M_access); i++ {
		entity, err := this.getEntityByUuid(this.M_access[i])
		if err == nil {
			values = append( values, entity.(*OAuth2Access))
		}
	}
	return values
}

func (this *OAuth2Configuration) SetAccess(val []*OAuth2Access){
	this.M_access= make([]string,0)
	for i:=0; i < len(val); i++{
		if len(val[i].GetParentUuid()) > 0  &&  len(val[i].GetParentLnk()) > 0 {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_access")
		this.M_access=append(this.M_access, val[i].GetUuid())
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *OAuth2Configuration) AppendAccess(val *OAuth2Access){
	for i:=0; i < len(this.M_access); i++{
		if this.M_access[i] == val.GetUuid() {
			return
		}
	}
	if len(val.GetParentUuid()) > 0 &&  len(val.GetParentLnk()) > 0 {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_access")
  this.setEntity(val)
	this.M_access = append(this.M_access, val.GetUuid())
	this.setEntity(this)
}

func (this *OAuth2Configuration) RemoveAccess(val *OAuth2Access){
	values := make([]string,0)
	for i:=0; i < len(this.M_access); i++{
		if this.M_access[i] != val.GetUuid() {
			values = append(values, this.M_access[i])
		}
	}
	this.M_access = values
	this.setEntity(this)
}


func (this *OAuth2Configuration) GetIds()[]*OAuth2IdToken{
	values := make([]*OAuth2IdToken, 0)
	for i := 0; i < len(this.M_ids); i++ {
		entity, err := this.getEntityByUuid(this.M_ids[i])
		if err == nil {
			values = append( values, entity.(*OAuth2IdToken))
		}
	}
	return values
}

func (this *OAuth2Configuration) SetIds(val []*OAuth2IdToken){
	this.M_ids= make([]string,0)
	for i:=0; i < len(val); i++{
		if len(val[i].GetParentUuid()) > 0  &&  len(val[i].GetParentLnk()) > 0 {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_ids")
		this.M_ids=append(this.M_ids, val[i].GetUuid())
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *OAuth2Configuration) AppendIds(val *OAuth2IdToken){
	for i:=0; i < len(this.M_ids); i++{
		if this.M_ids[i] == val.GetUuid() {
			return
		}
	}
	if len(val.GetParentUuid()) > 0 &&  len(val.GetParentLnk()) > 0 {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_ids")
  this.setEntity(val)
	this.M_ids = append(this.M_ids, val.GetUuid())
	this.setEntity(this)
}

func (this *OAuth2Configuration) RemoveIds(val *OAuth2IdToken){
	values := make([]string,0)
	for i:=0; i < len(this.M_ids); i++{
		if this.M_ids[i] != val.GetUuid() {
			values = append(values, this.M_ids[i])
		}
	}
	this.M_ids = values
	this.setEntity(this)
}


func (this *OAuth2Configuration) GetRefresh()[]*OAuth2Refresh{
	values := make([]*OAuth2Refresh, 0)
	for i := 0; i < len(this.M_refresh); i++ {
		entity, err := this.getEntityByUuid(this.M_refresh[i])
		if err == nil {
			values = append( values, entity.(*OAuth2Refresh))
		}
	}
	return values
}

func (this *OAuth2Configuration) SetRefresh(val []*OAuth2Refresh){
	this.M_refresh= make([]string,0)
	for i:=0; i < len(val); i++{
		if len(val[i].GetParentUuid()) > 0  &&  len(val[i].GetParentLnk()) > 0 {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_refresh")
		this.M_refresh=append(this.M_refresh, val[i].GetUuid())
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *OAuth2Configuration) AppendRefresh(val *OAuth2Refresh){
	for i:=0; i < len(this.M_refresh); i++{
		if this.M_refresh[i] == val.GetUuid() {
			return
		}
	}
	if len(val.GetParentUuid()) > 0 &&  len(val.GetParentLnk()) > 0 {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_refresh")
  this.setEntity(val)
	this.M_refresh = append(this.M_refresh, val.GetUuid())
	this.setEntity(this)
}

func (this *OAuth2Configuration) RemoveRefresh(val *OAuth2Refresh){
	values := make([]string,0)
	for i:=0; i < len(this.M_refresh); i++{
		if this.M_refresh[i] != val.GetUuid() {
			values = append(values, this.M_refresh[i])
		}
	}
	this.M_refresh = values
	this.setEntity(this)
}


func (this *OAuth2Configuration) GetExpire()[]*OAuth2Expires{
	values := make([]*OAuth2Expires, 0)
	for i := 0; i < len(this.M_expire); i++ {
		entity, err := this.getEntityByUuid(this.M_expire[i])
		if err == nil {
			values = append( values, entity.(*OAuth2Expires))
		}
	}
	return values
}

func (this *OAuth2Configuration) SetExpire(val []*OAuth2Expires){
	this.M_expire= make([]string,0)
	for i:=0; i < len(val); i++{
		if len(val[i].GetParentUuid()) > 0  &&  len(val[i].GetParentLnk()) > 0 {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_expire")
		this.M_expire=append(this.M_expire, val[i].GetUuid())
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *OAuth2Configuration) AppendExpire(val *OAuth2Expires){
	for i:=0; i < len(this.M_expire); i++{
		if this.M_expire[i] == val.GetUuid() {
			return
		}
	}
	if len(val.GetParentUuid()) > 0 &&  len(val.GetParentLnk()) > 0 {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_expire")
  this.setEntity(val)
	this.M_expire = append(this.M_expire, val.GetUuid())
	this.setEntity(this)
}

func (this *OAuth2Configuration) RemoveExpire(val *OAuth2Expires){
	values := make([]string,0)
	for i:=0; i < len(this.M_expire); i++{
		if this.M_expire[i] != val.GetUuid() {
			values = append(values, this.M_expire[i])
		}
	}
	this.M_expire = values
	this.setEntity(this)
}


func (this *OAuth2Configuration) GetParentPtr()*Configurations{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*Configurations)
	}
	return nil
}

func (this *OAuth2Configuration) SetParentPtr(val *Configurations){
	this.M_parentPtr= val.GetUuid()
	this.setEntity(this)
}


func (this *OAuth2Configuration) ResetParentPtr(){
	this.M_parentPtr= ""
}

