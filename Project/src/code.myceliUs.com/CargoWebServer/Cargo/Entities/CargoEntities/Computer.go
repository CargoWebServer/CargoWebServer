// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type Computer struct{

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

	/** members of Entity **/
	M_id string

	/** members of Computer **/
	M_name string
	M_ipv4 string
	M_osType OsType
	M_platformType PlatformType


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for Computer **/
type XsdComputer struct {
	XMLName xml.Name	`xml:"computerRef"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_osType	string	`xml:"osType,attr"`
	M_platformType	string	`xml:"platformType,attr"`
	M_name	string	`xml:"name,attr"`
	M_ipv4	string	`xml:"ipv4,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Computer) GetUuid() string{
	return this.UUID
}
func (this *Computer) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Computer) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Computer) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Computer"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Computer) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Computer) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Computer) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Computer) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *Computer) IsNeedSave() bool{
	return this.NeedSave
}
func (this *Computer) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Computer) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *Computer) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Computer) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Name **/
func (this *Computer) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Computer) SetName(ref interface{}){
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Name **/

/** Ipv4 **/
func (this *Computer) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *Computer) SetIpv4(ref interface{}){
	if this.M_ipv4 != ref.(string) {
		this.M_ipv4 = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Ipv4 **/

/** OsType **/
func (this *Computer) GetOsType() OsType{
	return this.M_osType
}

/** Init reference OsType **/
func (this *Computer) SetOsType(ref interface{}){
	if this.M_osType != ref.(OsType) {
		this.M_osType = ref.(OsType)
		this.NeedSave = true
	}
}

/** Remove reference OsType **/

/** PlatformType **/
func (this *Computer) GetPlatformType() PlatformType{
	return this.M_platformType
}

/** Init reference PlatformType **/
func (this *Computer) SetPlatformType(ref interface{}){
	if this.M_platformType != ref.(PlatformType) {
		this.M_platformType = ref.(PlatformType)
		this.NeedSave = true
	}
}

/** Remove reference PlatformType **/

/** Entities **/
func (this *Computer) GetEntitiesPtr() *Entities{
	if this.m_entitiesPtr == nil {
		entity, err := this.getEntityByUuid(this.M_entitiesPtr)
		if err == nil {
			this.m_entitiesPtr = entity.(*Entities)
		}
	}
	return this.m_entitiesPtr
}
func (this *Computer) GetEntitiesPtrStr() string{
	return this.M_entitiesPtr
}

/** Init reference Entities **/
func (this *Computer) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_entitiesPtr != ref.(string) {
			this.M_entitiesPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_entitiesPtr != ref.(*Entities).GetUuid() {
			this.M_entitiesPtr = ref.(*Entities).GetUuid()
			this.NeedSave = true
		}
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Computer) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
