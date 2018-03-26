// +build Config

package Config

import(
	"encoding/xml"
)

type OAuth2Client struct{

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

	/** members of OAuth2Client **/
	M_id string
	M_secret string
	M_redirectUri string
	M_tokenUri string
	M_authorizationUri string
	M_extra []uint8


	/** Associations **/
	M_parentPtr string
}

/** Xml parser for OAuth2Client **/
type XsdOAuth2Client struct {
	XMLName xml.Name	`xml:"oauth2Client"`
	M_id	string	`xml:"id,attr"`
	M_secret	string	`xml:"secret,attr"`
	M_redirectUri	string	`xml:"redirectUri,attr"`
	M_extra	[]uint8	`xml:"extra,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2Client) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *OAuth2Client) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2Client) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2Client) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2Client"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2Client) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2Client) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Client) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2Client) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *OAuth2Client) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Client) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *OAuth2Client) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Client) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *OAuth2Client) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *OAuth2Client) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *OAuth2Client) GetId()string{
	return this.M_id
}

func (this *OAuth2Client) SetId(val string){
	this.M_id= val
}




func (this *OAuth2Client) GetSecret()string{
	return this.M_secret
}

func (this *OAuth2Client) SetSecret(val string){
	this.M_secret= val
}




func (this *OAuth2Client) GetRedirectUri()string{
	return this.M_redirectUri
}

func (this *OAuth2Client) SetRedirectUri(val string){
	this.M_redirectUri= val
}




func (this *OAuth2Client) GetTokenUri()string{
	return this.M_tokenUri
}

func (this *OAuth2Client) SetTokenUri(val string){
	this.M_tokenUri= val
}




func (this *OAuth2Client) GetAuthorizationUri()string{
	return this.M_authorizationUri
}

func (this *OAuth2Client) SetAuthorizationUri(val string){
	this.M_authorizationUri= val
}




func (this *OAuth2Client) GetExtra()[]uint8{
	return this.M_extra
}

func (this *OAuth2Client) SetExtra(val []uint8){
	this.M_extra= val
}




func (this *OAuth2Client) GetParentPtr()*OAuth2Configuration{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*OAuth2Configuration)
	}
	return nil
}

func (this *OAuth2Client) SetParentPtr(val *OAuth2Configuration){
	this.M_parentPtr= val.GetUuid()
	this.setEntity(this)
}


func (this *OAuth2Client) ResetParentPtr(){
	this.M_parentPtr= ""
}

