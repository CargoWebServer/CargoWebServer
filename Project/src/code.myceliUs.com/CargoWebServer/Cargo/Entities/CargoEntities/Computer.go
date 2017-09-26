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
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

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
/** UUID **/
func (this *Computer) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Computer) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Computer) SetId(ref interface{}){
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Name **/
func (this *Computer) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Computer) SetName(ref interface{}){
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Ipv4 **/
func (this *Computer) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *Computer) SetIpv4(ref interface{}){
	this.M_ipv4 = ref.(string)
}

/** Remove reference Ipv4 **/

/** OsType **/
func (this *Computer) GetOsType() OsType{
	return this.M_osType
}

/** Init reference OsType **/
func (this *Computer) SetOsType(ref interface{}){
	this.M_osType = ref.(OsType)
}

/** Remove reference OsType **/

/** PlatformType **/
func (this *Computer) GetPlatformType() PlatformType{
	return this.M_platformType
}

/** Init reference PlatformType **/
func (this *Computer) SetPlatformType(ref interface{}){
	this.M_platformType = ref.(PlatformType)
}

/** Remove reference PlatformType **/

/** Entities **/
func (this *Computer) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Computer) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Computer) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
		}
	}
}
