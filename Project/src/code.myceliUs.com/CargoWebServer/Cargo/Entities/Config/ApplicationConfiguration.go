// +build Config

package Config

import(
	"encoding/xml"
)

type ApplicationConfiguration struct{

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

	/** members of ApplicationConfiguration **/
	M_indexPage string


	/** Associations **/
	m_parentPtr *Configurations
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for ApplicationConfiguration **/
type XsdApplicationConfiguration struct {
	XMLName xml.Name	`xml:"applicationConfiguration"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_indexPage	string	`xml:"indexPage,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *ApplicationConfiguration) GetUuid() string{
	return this.UUID
}
func (this *ApplicationConfiguration) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *ApplicationConfiguration) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *ApplicationConfiguration) GetTypeName() string{
	this.TYPENAME = "Config.ApplicationConfiguration"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *ApplicationConfiguration) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *ApplicationConfiguration) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ApplicationConfiguration) GetParentLnk() string{
	return this.ParentLnk
}
func (this *ApplicationConfiguration) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *ApplicationConfiguration) IsNeedSave() bool{
	return this.NeedSave
}
func (this *ApplicationConfiguration) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *ApplicationConfiguration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *ApplicationConfiguration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ApplicationConfiguration) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** IndexPage **/
func (this *ApplicationConfiguration) GetIndexPage() string{
	return this.M_indexPage
}

/** Init reference IndexPage **/
func (this *ApplicationConfiguration) SetIndexPage(ref interface{}){
	if this.M_indexPage != ref.(string) {
		this.M_indexPage = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference IndexPage **/

/** Parent **/
func (this *ApplicationConfiguration) GetParentPtr() *Configurations{
	if this.m_parentPtr == nil {
		entity, err := this.getEntityByUuid(this.M_parentPtr)
		if err == nil {
			this.m_parentPtr = entity.(*Configurations)
		}
	}
	return this.m_parentPtr
}
func (this *ApplicationConfiguration) GetParentPtrStr() string{
	return this.M_parentPtr
}

/** Init reference Parent **/
func (this *ApplicationConfiguration) SetParentPtr(ref interface{}){
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
func (this *ApplicationConfiguration) RemoveParentPtr(ref interface{}){
	toDelete := ref.(*Configurations)
	if this.m_parentPtr!= nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
