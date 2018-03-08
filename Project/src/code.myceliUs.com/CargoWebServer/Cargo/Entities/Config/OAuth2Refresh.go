// +build Config

package Config

import(
	"encoding/xml"
)

type OAuth2Refresh struct{

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

	/** members of OAuth2Refresh **/
	M_id string
	M_access string


	/** Associations **/
	M_parentPtr string
}

/** Xml parser for OAuth2Refresh **/
type XsdOAuth2Refresh struct {
	XMLName xml.Name	`xml:"oauth2Refresh"`
	M_id	string	`xml:"id,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2Refresh) GetUuid() string{
	return this.UUID
}
func (this *OAuth2Refresh) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2Refresh) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2Refresh) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2Refresh"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2Refresh) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2Refresh) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Refresh) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2Refresh) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Refresh) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Evaluate if an entity needs to be saved. **/
func (this *OAuth2Refresh) IsNeedSave() bool{
	return this.NeedSave
}
func (this *OAuth2Refresh) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Refresh) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *OAuth2Refresh) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *OAuth2Refresh) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *OAuth2Refresh) GetId()string{
	return this.M_id
}

func (this *OAuth2Refresh) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *OAuth2Refresh) GetAccess()*OAuth2Access{
	entity, err := this.getEntityByUuid(this.M_access)
	if err == nil {
		return entity.(*OAuth2Access)
	}
	return nil
}

func (this *OAuth2Refresh) SetAccess(val *OAuth2Access){
	this.NeedSave = this.M_access != val.GetUuid()
	this.M_access= val.GetUuid()
}

func (this *OAuth2Refresh) ResetAccess(){
	this.M_access= ""
}


func (this *OAuth2Refresh) GetParentPtr()*OAuth2Configuration{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*OAuth2Configuration)
	}
	return nil
}

func (this *OAuth2Refresh) SetParentPtr(val *OAuth2Configuration){
	this.NeedSave = this.M_parentPtr != val.GetUuid()
	this.M_parentPtr= val.GetUuid()
}

func (this *OAuth2Refresh) ResetParentPtr(){
	this.M_parentPtr= ""
}

