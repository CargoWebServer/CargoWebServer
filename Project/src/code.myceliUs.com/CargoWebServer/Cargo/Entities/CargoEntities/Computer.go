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
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

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
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
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

func (this *Computer) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Computer) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *Computer) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Computer) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Computer) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Computer) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Computer) GetId()string{
	return this.M_id
}

func (this *Computer) SetId(val string){
	this.M_id= val
}




func (this *Computer) GetName()string{
	return this.M_name
}

func (this *Computer) SetName(val string){
	this.M_name= val
}




func (this *Computer) GetIpv4()string{
	return this.M_ipv4
}

func (this *Computer) SetIpv4(val string){
	this.M_ipv4= val
}




func (this *Computer) GetOsType()OsType{
	return this.M_osType
}

func (this *Computer) SetOsType(val OsType){
	this.M_osType= val
}


func (this *Computer) ResetOsType(){
	this.M_osType= 0
}


func (this *Computer) GetPlatformType()PlatformType{
	return this.M_platformType
}

func (this *Computer) SetPlatformType(val PlatformType){
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
	this.setEntity(this)
}


func (this *Computer) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

