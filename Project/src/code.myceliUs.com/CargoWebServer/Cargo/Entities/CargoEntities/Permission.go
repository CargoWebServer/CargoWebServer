// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type Permission struct{

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

	/** members of Permission **/
	M_id string
	M_types int
	M_accountsRef []string


	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for Permission **/
type XsdPermission struct {
	XMLName xml.Name	`xml:"permissionsRef"`
	M_accountsRef	[]string	`xml:"accountsRef"`
	M_id	string	`xml:"id,attr"`
	M_types	int	`xml:"types,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Permission) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Permission) SetUuid(uuid string){
	this.UUID = uuid
}

func (this *Permission) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *Permission) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *Permission) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Permission) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Permission"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Permission) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Permission) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Permission) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Permission) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *Permission) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Permission) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *Permission) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Permission) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Permission) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Permission) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Permission) GetId()string{
	return this.M_id
}

func (this *Permission) SetId(val string){
	this.M_id= val
}




func (this *Permission) GetTypes()int{
	return this.M_types
}

func (this *Permission) SetTypes(val int){
	this.M_types= val
}




func (this *Permission) GetAccountsRef()[]*Account{
	values := make([]*Account, 0)
	for i := 0; i < len(this.M_accountsRef); i++ {
		entity, err := this.getEntityByUuid(this.M_accountsRef[i])
		if err == nil {
			values = append( values, entity.(*Account))
		}
	}
	return values
}

func (this *Permission) SetAccountsRef(val []*Account){
	this.M_accountsRef= make([]string,0)
	for i:=0; i < len(val); i++{
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *Permission) AppendAccountsRef(val *Account){
	for i:=0; i < len(this.M_accountsRef); i++{
		if this.M_accountsRef[i] == val.GetUuid() {
			return
		}
	}
	this.M_accountsRef = append(this.M_accountsRef, val.GetUuid())
	this.setEntity(this)
}

func (this *Permission) RemoveAccountsRef(val *Account){
	values := make([]string,0)
	for i:=0; i < len(this.M_accountsRef); i++{
		if this.M_accountsRef[i] != val.GetUuid() {
			values = append(values, this.M_accountsRef[i])
		}
	}
	this.M_accountsRef = values
	this.setEntity(this)
}


func (this *Permission) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Permission) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
	this.setEntity(this)
}


func (this *Permission) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

