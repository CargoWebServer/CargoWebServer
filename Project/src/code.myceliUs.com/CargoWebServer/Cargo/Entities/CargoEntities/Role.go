// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type Role struct{

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

	/** members of Role **/
	M_id string
	M_accountsRef []string
	M_actionsRef []string


	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for Role **/
type XsdRole struct {
	XMLName xml.Name	`xml:"role"`
	M_accountsRef	[]string	`xml:"accountsRef"`
	M_actionsRef	[]string	`xml:"actionsRef"`
	M_id	string	`xml:"id,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Role) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Role) SetUuid(uuid string){
	this.UUID = uuid
}

/** Need save **/
func (this *Role) IsNeedSave() bool{
	return this.needSave
}
func (this *Role) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *Role) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *Role) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *Role) RemoveReferenced(uuid string, field string) {
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

func (this *Role) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *Role) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *Role) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Role) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Role"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Role) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Role) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Role) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Role) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *Role) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Role) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *Role) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Role) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Role) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Role) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Role) GetId()string{
	return this.M_id
}

func (this *Role) SetId(val string){
	this.M_id= val
}




func (this *Role) GetAccountsRef()[]*Account{
	values := make([]*Account, 0)
	for i := 0; i < len(this.M_accountsRef); i++ {
		entity, err := this.getEntityByUuid(this.M_accountsRef[i])
		if err == nil {
			values = append( values, entity.(*Account))
		}
	}
	return values
}

func (this *Role) SetAccountsRef(val []*Account){
	this.M_accountsRef= make([]string,0)
	for i:=0; i < len(val); i++{
		this.setEntity(val[i])
	}
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Role) AppendAccountsRef(val *Account){
	for i:=0; i < len(this.M_accountsRef); i++{
		if this.M_accountsRef[i] == val.GetUuid() {
			return
		}
	}
	if this.M_accountsRef== nil {
		this.M_accountsRef = make([]string, 0)
	}

	this.M_accountsRef = append(this.M_accountsRef, val.GetUuid())
	this.setEntity(this)
	this.SetNeedSave(true)
}

func (this *Role) RemoveAccountsRef(val *Account){
	values := make([]string,0)
	for i:=0; i < len(this.M_accountsRef); i++{
		if this.M_accountsRef[i] != val.GetUuid() {
			values = append(values, this.M_accountsRef[i])
		}
	}
	this.M_accountsRef = values
	this.setEntity(this)
}


func (this *Role) GetActionsRef()[]*Action{
	values := make([]*Action, 0)
	for i := 0; i < len(this.M_actionsRef); i++ {
		entity, err := this.getEntityByUuid(this.M_actionsRef[i])
		if err == nil {
			values = append( values, entity.(*Action))
		}
	}
	return values
}

func (this *Role) SetActionsRef(val []*Action){
	this.M_actionsRef= make([]string,0)
	for i:=0; i < len(val); i++{
		this.setEntity(val[i])
	}
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Role) AppendActionsRef(val *Action){
	for i:=0; i < len(this.M_actionsRef); i++{
		if this.M_actionsRef[i] == val.GetUuid() {
			return
		}
	}
	if this.M_actionsRef== nil {
		this.M_actionsRef = make([]string, 0)
	}

	this.M_actionsRef = append(this.M_actionsRef, val.GetUuid())
	this.setEntity(this)
	this.SetNeedSave(true)
}

func (this *Role) RemoveActionsRef(val *Action){
	values := make([]string,0)
	for i:=0; i < len(this.M_actionsRef); i++{
		if this.M_actionsRef[i] != val.GetUuid() {
			values = append(values, this.M_actionsRef[i])
		}
	}
	this.M_actionsRef = values
	this.setEntity(this)
}


func (this *Role) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Role) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *Role) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
	this.setEntity(this)
}

