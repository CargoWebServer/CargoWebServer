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
	m_accounts []*Account
	/** If the ref is a string and not an object **/
	M_accounts []string
	m_actions []*Action
	/** If the ref is a string and not an object **/
	M_actions []string


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
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

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Role) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *Role) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Role) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Accounts **/
func (this *Role) GetAccounts() []*Account{
	if this.m_accounts == nil {
		this.m_accounts = make([]*Account, 0)
		for i := 0; i < len(this.M_accounts); i++ {
			entity, err := this.getEntityByUuid(this.M_accounts[i])
			if err == nil {
				this.m_accounts = append(this.m_accounts, entity.(*Account))
			}
		}
	}
	return this.m_accounts
}
func (this *Role) GetAccountsStr() []string{
	return this.M_accounts
}

/** Init reference Accounts **/
func (this *Role) SetAccounts(ref interface{}){
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_accounts); i++ {
			if this.M_accounts[i] == refStr {
				return
			}
		}
		this.M_accounts = append(this.M_accounts, ref.(string))
		this.NeedSave = true
	}else{
		for i:=0; i < len(this.m_accounts); i++ {
			if this.m_accounts[i].GetUuid() == ref.(*Account).GetUuid() {
				return
			}
		}
		isExist := false
		for i:=0; i < len(this.M_accounts); i++ {
			if this.M_accounts[i] == ref.(*Account).GetUuid() {
				isExist = true
			}
		}
		this.m_accounts = append(this.m_accounts, ref.(*Account))
	if !isExist {
		this.M_accounts = append(this.M_accounts, ref.(Entity).GetUuid())
		this.NeedSave = true
	}
	}
}

/** Remove reference Accounts **/
func (this *Role) RemoveAccounts(ref interface{}){
	toDelete := ref.(Entity)
	accounts_ := make([]*Account, 0)
	accountsUuid := make([]string, 0)
	for i := 0; i < len(this.m_accounts); i++ {
		if toDelete.GetUuid() != this.m_accounts[i].GetUuid() {
			accounts_ = append(accounts_, this.m_accounts[i])
			accountsUuid = append(accountsUuid, this.M_accounts[i])
		}else{
			this.NeedSave = true
		}
	}
	this.m_accounts = accounts_
	this.M_accounts = accountsUuid
}

/** Actions **/
func (this *Role) GetActions() []*Action{
	if this.m_actions == nil {
		this.m_actions = make([]*Action, 0)
		for i := 0; i < len(this.M_actions); i++ {
			entity, err := this.getEntityByUuid(this.M_actions[i])
			if err == nil {
				this.m_actions = append(this.m_actions, entity.(*Action))
			}
		}
	}
	return this.m_actions
}
func (this *Role) GetActionsStr() []string{
	return this.M_actions
}

/** Init reference Actions **/
func (this *Role) SetActions(ref interface{}){
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_actions); i++ {
			if this.M_actions[i] == refStr {
				return
			}
		}
		this.M_actions = append(this.M_actions, ref.(string))
		this.NeedSave = true
	}else{
		for i:=0; i < len(this.m_actions); i++ {
			if this.m_actions[i].GetUuid() == ref.(*Action).GetUuid() {
				return
			}
		}
		isExist := false
		for i:=0; i < len(this.M_actions); i++ {
			if this.M_actions[i] == ref.(*Action).GetUuid() {
				isExist = true
			}
		}
		this.m_actions = append(this.m_actions, ref.(*Action))
	if !isExist {
		this.M_actions = append(this.M_actions, ref.(*Action).GetUuid())
		this.NeedSave = true
	}
	}
}

/** Remove reference Actions **/
func (this *Role) RemoveActions(ref interface{}){
	toDelete := ref.(*Action)
	actions_ := make([]*Action, 0)
	actionsUuid := make([]string, 0)
	for i := 0; i < len(this.m_actions); i++ {
		if toDelete.GetUuid() != this.m_actions[i].GetUuid() {
			actions_ = append(actions_, this.m_actions[i])
			actionsUuid = append(actionsUuid, this.M_actions[i])
		}else{
			this.NeedSave = true
		}
	}
	this.m_actions = actions_
	this.M_actions = actionsUuid
}

/** Entities **/
func (this *Role) GetEntitiesPtr() *Entities{
	if this.m_entitiesPtr == nil {
		entity, err := this.getEntityByUuid(this.M_entitiesPtr)
		if err == nil {
			this.m_entitiesPtr = entity.(*Entities)
		}
	}
	return this.m_entitiesPtr
}
func (this *Role) GetEntitiesPtrStr() string{
	return this.M_entitiesPtr
}

/** Init reference Entities **/
func (this *Role) SetEntitiesPtr(ref interface{}){
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
func (this *Role) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
