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
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of OAuth2Access **/
	M_id string
	M_client string
	M_authorize string
	M_previous string
	M_refreshToken string
	M_expiresIn int64
	M_scope string
	M_redirectUri string
	M_userData string
	M_createdAt int64


	/** Associations **/
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
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2Access) GetUuid() string{
	return this.UUID
}
func (this *OAuth2Access) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2Access) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2Access) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2Access"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2Access) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2Access) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Access) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2Access) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Access) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Evaluate if an entity needs to be saved. **/
func (this *OAuth2Access) IsNeedSave() bool{
	return this.NeedSave
}
func (this *OAuth2Access) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Access) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *OAuth2Access) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *OAuth2Access) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *OAuth2Access) GetId()string{
	return this.M_id
}

func (this *OAuth2Access) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *OAuth2Access) GetClient()*OAuth2Client{
	entity, err := this.getEntityByUuid(this.M_client)
	if err == nil {
		return entity.(*OAuth2Client)
	}
	return nil
}

func (this *OAuth2Access) SetClient(val *OAuth2Client){
	this.NeedSave = this.M_client != val.GetUuid()
	this.M_client= val.GetUuid()
}

func (this *OAuth2Access) ResetClient(){
	this.M_client= ""
}


func (this *OAuth2Access) GetAuthorize()string{
	return this.M_authorize
}

func (this *OAuth2Access) SetAuthorize(val string){
	this.NeedSave = this.M_authorize== val
	this.M_authorize= val
}


func (this *OAuth2Access) GetPrevious()string{
	return this.M_previous
}

func (this *OAuth2Access) SetPrevious(val string){
	this.NeedSave = this.M_previous== val
	this.M_previous= val
}


func (this *OAuth2Access) GetRefreshToken()*OAuth2Refresh{
	entity, err := this.getEntityByUuid(this.M_refreshToken)
	if err == nil {
		return entity.(*OAuth2Refresh)
	}
	return nil
}

func (this *OAuth2Access) SetRefreshToken(val *OAuth2Refresh){
	this.NeedSave = this.M_refreshToken != val.GetUuid()
	this.M_refreshToken= val.GetUuid()
}

func (this *OAuth2Access) ResetRefreshToken(){
	this.M_refreshToken= ""
}


func (this *OAuth2Access) GetExpiresIn()int64{
	return this.M_expiresIn
}

func (this *OAuth2Access) SetExpiresIn(val int64){
	this.NeedSave = this.M_expiresIn== val
	this.M_expiresIn= val
}


func (this *OAuth2Access) GetScope()string{
	return this.M_scope
}

func (this *OAuth2Access) SetScope(val string){
	this.NeedSave = this.M_scope== val
	this.M_scope= val
}


func (this *OAuth2Access) GetRedirectUri()string{
	return this.M_redirectUri
}

func (this *OAuth2Access) SetRedirectUri(val string){
	this.NeedSave = this.M_redirectUri== val
	this.M_redirectUri= val
}


func (this *OAuth2Access) GetUserData()*OAuth2IdToken{
	entity, err := this.getEntityByUuid(this.M_userData)
	if err == nil {
		return entity.(*OAuth2IdToken)
	}
	return nil
}

func (this *OAuth2Access) SetUserData(val *OAuth2IdToken){
	this.NeedSave = this.M_userData != val.GetUuid()
	this.M_userData= val.GetUuid()
}

func (this *OAuth2Access) ResetUserData(){
	this.M_userData= ""
}


func (this *OAuth2Access) GetCreatedAt()int64{
	return this.M_createdAt
}

func (this *OAuth2Access) SetCreatedAt(val int64){
	this.NeedSave = this.M_createdAt== val
	this.M_createdAt= val
}


func (this *OAuth2Access) GetParentPtr()*OAuth2Configuration{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*OAuth2Configuration)
	}
	return nil
}

func (this *OAuth2Access) SetParentPtr(val *OAuth2Configuration){
	this.NeedSave = this.M_parentPtr != val.GetUuid()
	this.M_parentPtr= val.GetUuid()
}

func (this *OAuth2Access) ResetParentPtr(){
	this.M_parentPtr= ""
}

