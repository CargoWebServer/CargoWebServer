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
	M_id	string	`xml:"id,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Role) GetUuid() string{
	return this.UUID
}
func (this *Role) SetUuid(uuid string){
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

func (this *Role) GetId()string{
	return this.M_id
}

func (this *Role) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *Role) GetAccounts()[]*Account{
	accounts := make([]*Account, 0)
	for i := 0; i < len(this.M_accounts); i++ {
		entity, err := this.getEntityByUuid(this.M_accounts[i])
		if err == nil {
			accounts = append(accounts, entity.(*Account))
		}
	}
	return accounts
}

func (this *Role) SetAccounts(val []*Account){
	this.M_accounts= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_accounts=append(this.M_accounts, val[i].GetUuid())
	}
}

func (this *Role) AppendAccounts(val *Account){
	for i:=0; i < len(this.M_accounts); i++{
		if this.M_accounts[i] == val.GetUuid() {
			return
		}
	}
	this.M_accounts = append(this.M_accounts, val.GetUuid())
}

func (this *Role) RemoveAccounts(val *Account){
	accounts := make([]string,0)
	for i:=0; i < len(this.M_accounts); i++{
		if this.M_accounts[i] != val.GetUuid() {
			accounts = append(accounts, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_accounts = accounts
}


func (this *Role) GetActions()[]*Action{
	actions := make([]*Action, 0)
	for i := 0; i < len(this.M_actions); i++ {
		entity, err := this.getEntityByUuid(this.M_actions[i])
		if err == nil {
			actions = append(actions, entity.(*Action))
		}
	}
	return actions
}

func (this *Role) SetActions(val []*Action){
	this.M_actions= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_actions=append(this.M_actions, val[i].GetUuid())
	}
}

func (this *Role) AppendActions(val *Action){
	for i:=0; i < len(this.M_actions); i++{
		if this.M_actions[i] == val.GetUuid() {
			return
		}
	}
	this.M_actions = append(this.M_actions, val.GetUuid())
}

func (this *Role) RemoveActions(val *Action){
	actions := make([]string,0)
	for i:=0; i < len(this.M_actions); i++{
		if this.M_actions[i] != val.GetUuid() {
			actions = append(actions, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_actions = actions
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
}

func (this *Role) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

