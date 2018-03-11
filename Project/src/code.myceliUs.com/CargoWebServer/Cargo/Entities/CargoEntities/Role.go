// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
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
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of Role **/
	M_id string
	M_accounts []string
	M_actions []string


	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for Role **/
type XsdRole struct {
	XMLName xml.Name	`xml:"role"`
	M_accounts	[]*XsdAccount	`xml:"accounts,omitempty"`
	M_actions	[]*XsdAction	`xml:"actions,omitempty"`
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
	this.NeedSave = this.UUID == uuid
	this.UUID = uuid
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
/** Evaluate if an entity needs to be saved. **/
func (this *Role) IsNeedSave() bool{
	return this.NeedSave
}
func (this *Role) ResetNeedSave(){
	this.NeedSave=false
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
	this.NeedSave = this.M_id== val
	this.M_id= val
}




func (this *Role) GetAccounts()[]*Account{
	values := make([]*Account, 0)
	for i := 0; i < len(this.M_accounts); i++ {
		entity, err := this.getEntityByUuid(this.M_accounts[i])
		if err == nil {
			values = append( values, entity.(*Account))
		}
	}
	return values
}

func (this *Role) SetAccounts(val []*Account){
	this.M_accounts= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_accounts=append(this.M_accounts, val[i].GetUuid())
	}
	this.NeedSave= true
}


func (this *Role) AppendAccounts(val *Account){
	for i:=0; i < len(this.M_accounts); i++{
		if this.M_accounts[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_accounts = append(this.M_accounts, val.GetUuid())
}

func (this *Role) RemoveAccounts(val *Account){
	values := make([]string,0)
	for i:=0; i < len(this.M_accounts); i++{
		if this.M_accounts[i] != val.GetUuid() {
			values = append(values, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_accounts = values
}


func (this *Role) GetActions()[]*Action{
	values := make([]*Action, 0)
	for i := 0; i < len(this.M_actions); i++ {
		entity, err := this.getEntityByUuid(this.M_actions[i])
		if err == nil {
			values = append( values, entity.(*Action))
		}
	}
	return values
}

func (this *Role) SetActions(val []*Action){
	this.M_actions= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_actions=append(this.M_actions, val[i].GetUuid())
	}
	this.NeedSave= true
}


func (this *Role) AppendActions(val *Action){
	for i:=0; i < len(this.M_actions); i++{
		if this.M_actions[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_actions = append(this.M_actions, val.GetUuid())
}

func (this *Role) RemoveActions(val *Action){
	values := make([]string,0)
	for i:=0; i < len(this.M_actions); i++{
		if this.M_actions[i] != val.GetUuid() {
			values = append(values, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_actions = values
}


func (this *Role) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Role) SetEntitiesPtr(val *Entities){
	this.NeedSave = this.M_entitiesPtr != val.GetUuid()
	this.M_entitiesPtr= val.GetUuid()
}


func (this *Role) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

