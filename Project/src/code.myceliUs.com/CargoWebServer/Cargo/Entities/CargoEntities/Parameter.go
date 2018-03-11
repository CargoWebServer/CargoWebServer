// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type Parameter struct{

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

	/** members of Parameter **/
	M_name string
	M_type string
	M_isArray bool


	/** Associations **/
	M_parametersPtr string
}

/** Xml parser for Parameter **/
type XsdParameter struct {
	XMLName xml.Name	`xml:"parameter"`
	M_name	string	`xml:"name,attr"`
	M_type	string	`xml:"type,attr"`
	M_isArray	bool	`xml:"isArray,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Parameter) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Parameter) SetUuid(uuid string){
	this.NeedSave = this.UUID == uuid
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Parameter) Ids() []interface{} {
	ids := make([]interface{}, 0)
	return ids
}

/** The type name **/
func (this *Parameter) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Parameter"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Parameter) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Parameter) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Parameter) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Parameter) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *Parameter) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Parameter) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Evaluate if an entity needs to be saved. **/
func (this *Parameter) IsNeedSave() bool{
	return this.NeedSave
}
func (this *Parameter) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Parameter) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Parameter) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Parameter) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Parameter) GetName()string{
	return this.M_name
}

func (this *Parameter) SetName(val string){
	this.NeedSave = this.M_name== val
	this.M_name= val
}




func (this *Parameter) GetType()string{
	return this.M_type
}

func (this *Parameter) SetType(val string){
	this.NeedSave = this.M_type== val
	this.M_type= val
}




func (this *Parameter) IsArray()bool{
	return this.M_isArray
}

func (this *Parameter) SetIsArray(val bool){
	this.NeedSave = this.M_isArray== val
	this.M_isArray= val
}




func (this *Parameter) GetParametersPtr()*Parameter{
	entity, err := this.getEntityByUuid(this.M_parametersPtr)
	if err == nil {
		return entity.(*Parameter)
	}
	return nil
}

func (this *Parameter) SetParametersPtr(val *Parameter){
	this.NeedSave = this.M_parametersPtr != val.GetUuid()
	this.M_parametersPtr= val.GetUuid()
}


func (this *Parameter) ResetParametersPtr(){
	this.M_parametersPtr= ""
}

