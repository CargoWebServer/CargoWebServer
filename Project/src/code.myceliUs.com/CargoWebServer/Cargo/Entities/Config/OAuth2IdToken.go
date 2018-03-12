// +build Config

package Config

import(
	"encoding/xml"
)

type OAuth2IdToken struct{

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

	/** members of OAuth2IdToken **/
	M_issuer string
	M_id string
	M_client string
	M_expiration int64
	M_issuedAt int64
	M_nonce string
	M_email string
	M_emailVerified bool
	M_name string
	M_familyName string
	M_givenName string
	M_local string


	/** Associations **/
	M_parentPtr string
}

/** Xml parser for OAuth2IdToken **/
type XsdOAuth2IdToken struct {
	XMLName xml.Name	`xml:"oauth2IdToken"`
	M_id	string	`xml:"id,attr"`
	M_issuer	string	`xml:"issuer,attr"`
	M_expiration	int64	`xml:"expiration,attr"`
	M_issuedAt	int64	`xml:"issuedAt,attr"`
	M_nonce	string	`xml:"nonce,attr"`
	M_email	string	`xml:"email,attr"`
	M_emailVerified	bool	`xml:"emailVerified,attr"`
	M_name	string	`xml:"name,attr"`
	M_familyName	string	`xml:"familyName,attr"`
	M_givenName	string	`xml:"givenName,attr"`
	M_local	string	`xml:"local,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2IdToken) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *OAuth2IdToken) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2IdToken) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2IdToken) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2IdToken"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2IdToken) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2IdToken) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2IdToken) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2IdToken) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *OAuth2IdToken) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2IdToken) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2IdToken) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *OAuth2IdToken) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *OAuth2IdToken) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *OAuth2IdToken) GetIssuer()string{
	return this.M_issuer
}

func (this *OAuth2IdToken) SetIssuer(val string){
	this.M_issuer= val
}




func (this *OAuth2IdToken) GetId()string{
	return this.M_id
}

func (this *OAuth2IdToken) SetId(val string){
	this.M_id= val
}




func (this *OAuth2IdToken) GetClient()*OAuth2Client{
	entity, err := this.getEntityByUuid(this.M_client)
	if err == nil {
		return entity.(*OAuth2Client)
	}
	return nil
}

func (this *OAuth2IdToken) SetClient(val *OAuth2Client){
	this.M_client= val.GetUuid()
	this.setEntity(this)
}


func (this *OAuth2IdToken) ResetClient(){
	this.M_client= ""
}


func (this *OAuth2IdToken) GetExpiration()int64{
	return this.M_expiration
}

func (this *OAuth2IdToken) SetExpiration(val int64){
	this.M_expiration= val
}




func (this *OAuth2IdToken) GetIssuedAt()int64{
	return this.M_issuedAt
}

func (this *OAuth2IdToken) SetIssuedAt(val int64){
	this.M_issuedAt= val
}




func (this *OAuth2IdToken) GetNonce()string{
	return this.M_nonce
}

func (this *OAuth2IdToken) SetNonce(val string){
	this.M_nonce= val
}




func (this *OAuth2IdToken) GetEmail()string{
	return this.M_email
}

func (this *OAuth2IdToken) SetEmail(val string){
	this.M_email= val
}




func (this *OAuth2IdToken) IsEmailVerified()bool{
	return this.M_emailVerified
}

func (this *OAuth2IdToken) SetEmailVerified(val bool){
	this.M_emailVerified= val
}




func (this *OAuth2IdToken) GetName()string{
	return this.M_name
}

func (this *OAuth2IdToken) SetName(val string){
	this.M_name= val
}




func (this *OAuth2IdToken) GetFamilyName()string{
	return this.M_familyName
}

func (this *OAuth2IdToken) SetFamilyName(val string){
	this.M_familyName= val
}




func (this *OAuth2IdToken) GetGivenName()string{
	return this.M_givenName
}

func (this *OAuth2IdToken) SetGivenName(val string){
	this.M_givenName= val
}




func (this *OAuth2IdToken) GetLocal()string{
	return this.M_local
}

func (this *OAuth2IdToken) SetLocal(val string){
	this.M_local= val
}




func (this *OAuth2IdToken) GetParentPtr()*OAuth2Configuration{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*OAuth2Configuration)
	}
	return nil
}

func (this *OAuth2IdToken) SetParentPtr(val *OAuth2Configuration){
	this.M_parentPtr= val.GetUuid()
	this.setEntity(this)
}


func (this *OAuth2IdToken) ResetParentPtr(){
	this.M_parentPtr= ""
}

