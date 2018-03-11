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
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of OAuth2Authorize **/
	M_id string
	M_client string
	M_expiresIn int64
	M_scope string
	M_redirectUri string
	M_state string
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
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2Authorize) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *OAuth2Authorize) SetUuid(uuid string){
	this.NeedSave = this.UUID == uuid
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2Authorize) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2Authorize) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2Authorize"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2Authorize) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2Authorize) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Authorize) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2Authorize) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *OAuth2Authorize) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Authorize) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Evaluate if an entity needs to be saved. **/
func (this *OAuth2Authorize) IsNeedSave() bool{
	return this.NeedSave
}
func (this *OAuth2Authorize) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Authorize) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *OAuth2Authorize) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *OAuth2Authorize) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *OAuth2Authorize) GetId()string{
	return this.M_id
}

func (this *OAuth2Authorize) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}




func (this *OAuth2Authorize) GetClient()*OAuth2Client{
	entity, err := this.getEntityByUuid(this.M_client)
	if err == nil {
		return entity.(*OAuth2Client)
	}
	return nil
}

func (this *OAuth2Authorize) SetClient(val *OAuth2Client){
	this.NeedSave = this.M_client != val.GetUuid()
	this.M_client= val.GetUuid()
}


func (this *OAuth2Authorize) ResetClient(){
	this.M_client= ""
}


func (this *OAuth2Authorize) GetExpiresIn()int64{
	return this.M_expiresIn
}

func (this *OAuth2Authorize) SetExpiresIn(val int64){
	this.NeedSave = this.M_expiresIn== val
	this.M_expiresIn= val
}




func (this *OAuth2Authorize) GetScope()string{
	return this.M_scope
}

func (this *OAuth2Authorize) SetScope(val string){
	this.NeedSave = this.M_scope== val
	this.M_scope= val
}




func (this *OAuth2Authorize) GetRedirectUri()string{
	return this.M_redirectUri
}

func (this *OAuth2Authorize) SetRedirectUri(val string){
	this.NeedSave = this.M_redirectUri== val
	this.M_redirectUri= val
}




func (this *OAuth2Authorize) GetState()string{
	return this.M_state
}

func (this *OAuth2Authorize) SetState(val string){
	this.NeedSave = this.M_state== val
	this.M_state= val
}




func (this *OAuth2Authorize) GetUserData()*OAuth2IdToken{
	entity, err := this.getEntityByUuid(this.M_userData)
	if err == nil {
		return entity.(*OAuth2IdToken)
	}
	return nil
}

func (this *OAuth2Authorize) SetUserData(val *OAuth2IdToken){
	this.NeedSave = this.M_userData != val.GetUuid()
	this.M_userData= val.GetUuid()
}


func (this *OAuth2Authorize) ResetUserData(){
	this.M_userData= ""
}


func (this *OAuth2Authorize) GetCreatedAt()int64{
	return this.M_createdAt
}

func (this *OAuth2Authorize) SetCreatedAt(val int64){
	this.NeedSave = this.M_createdAt== val
	this.M_createdAt= val
}



