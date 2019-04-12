// +build Config

package Config

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type OAuth2Expires struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** keep track if the entity has change over time. **/
	needSave bool
	/** Keep reference to entity that made use of thit entity **/
	Referenced []string
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of OAuth2Expires **/
	M_id string
	M_expiresAt int64


	/** Associations **/
	M_parentPtr string
}

/** Xml parser for OAuth2Expires **/
type XsdOAuth2Expires struct {
	XMLName xml.Name	`xml:"oauth2Expires"`
	M_id	string	`xml:"id,attr"`
	M_expiresAt	int64	`xml:"expiresAt,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2Expires) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *OAuth2Expires) SetUuid(uuid string){
	this.UUID = uuid
}

/** Need save **/
func (this *OAuth2Expires) IsNeedSave() bool{
	return this.needSave
}
func (this *OAuth2Expires) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *OAuth2Expires) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *OAuth2Expires) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *OAuth2Expires) RemoveReferenced(uuid string, field string) {
	if this.Referenced == nil {
		return
	}
	referenced := make([]string, 0)
	for i := 0; i < len(this.Referenced); i++ {
		if this.Referenced[i] != uuid+":"+field {
			referenced = append(referenced, uuid+":"+field)
		}
	}
	this.Referenced = referenced
}

func (this *OAuth2Expires) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *OAuth2Expires) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2Expires) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2Expires) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2Expires"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2Expires) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2Expires) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Expires) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2Expires) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *OAuth2Expires) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Expires) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *OAuth2Expires) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Expires) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *OAuth2Expires) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *OAuth2Expires) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *OAuth2Expires) GetId()string{
	return this.M_id
}

func (this *OAuth2Expires) SetId(val string){
	this.M_id= val
}




func (this *OAuth2Expires) GetExpiresAt()int64{
	return this.M_expiresAt
}

func (this *OAuth2Expires) SetExpiresAt(val int64){
	this.M_expiresAt= val
}




func (this *OAuth2Expires) GetParentPtr()*OAuth2Configuration{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*OAuth2Configuration)
	}
	return nil
}

func (this *OAuth2Expires) SetParentPtr(val *OAuth2Configuration){
	this.M_parentPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *OAuth2Expires) ResetParentPtr(){
	this.M_parentPtr= ""
	this.setEntity(this)
}

