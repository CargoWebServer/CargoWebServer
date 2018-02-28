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

func (this *Computer) GetId()string{
	return this.M_id
}

func (this *Computer) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *Computer) GetName()string{
	return this.M_name
}

func (this *Computer) SetName(val string){
	this.NeedSave = this.M_name== val
	this.M_name= val
}


func (this *Computer) GetIpv4()string{
	return this.M_ipv4
}

func (this *Computer) SetIpv4(val string){
	this.NeedSave = this.M_ipv4== val
	this.M_ipv4= val
}


func (this *Computer) GetOsType()OsType{
	return this.M_osType
}

func (this *Computer) SetOsType(val OsType){
	this.NeedSave = this.M_osType== val
	this.M_osType= val
}

func (this *Computer) ResetOsType(){
	this.M_osType= 0
}


func (this *Computer) GetPlatformType()PlatformType{
	return this.M_platformType
}

func (this *Computer) SetPlatformType(val PlatformType){
	this.NeedSave = this.M_platformType== val
	this.M_platformType= val
}

func (this *Computer) ResetPlatformType(){
	this.M_platformType= 0
}


func (this *Computer) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Computer) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
}

func (this *Computer) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

