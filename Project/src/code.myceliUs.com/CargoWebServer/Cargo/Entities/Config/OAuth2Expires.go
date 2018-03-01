// +build Config

package Config

import(
	"encoding/xml"
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
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

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
	return this.UUID
}
func (this *OAuth2Expires) SetUuid(uuid string){
	this.UUID = uuid
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

/** Evaluate if an entity needs to be saved. **/
func (this *OAuth2Expires) IsNeedSave() bool{
	return this.NeedSave
}
func (this *OAuth2Expires) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Expires) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

func (this *OAuth2Expires) GetId()string{
	return this.M_id
}

func (this *OAuth2Expires) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *OAuth2Expires) GetExpiresAt()int64{
	return this.M_expiresAt
}

func (this *OAuth2Expires) SetExpiresAt(val int64){
	this.NeedSave = this.M_expiresAt== val
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
	this.NeedSave = this.M_parentPtr != val.GetUuid()
	this.M_parentPtr= val.GetUuid()
}

func (this *OAuth2Expires) ResetParentPtr(){
	this.M_parentPtr= ""
}

