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

	/** If the entity is fully initialyse **/
	IsInit   bool

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
/** UUID **/
func (this *Role) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Role) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Role) SetId(ref interface{}){
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Accounts **/
func (this *Role) GetAccounts() []*Account{
	return this.m_accounts
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
	}else{
		for i:=0; i < len(this.m_accounts); i++ {
			if this.m_accounts[i].GetUUID() == ref.(*Account).GetUUID() {
				return
			}
		}
		isExist := false
		for i:=0; i < len(this.M_accounts); i++ {
			if this.M_accounts[i] == ref.(*Account).GetUUID() {
				isExist = true
			}
		}
		this.m_accounts = append(this.m_accounts, ref.(*Account))
	if !isExist {
		this.M_accounts = append(this.M_accounts, ref.(Entity).GetUUID())
	}
	}
}

/** Remove reference Accounts **/
func (this *Role) RemoveAccounts(ref interface{}){
	toDelete := ref.(Entity)
	accounts_ := make([]*Account, 0)
	accountsUuid := make([]string, 0)
	for i := 0; i < len(this.m_accounts); i++ {
		if toDelete.GetUUID() != this.m_accounts[i].GetUUID() {
			accounts_ = append(accounts_, this.m_accounts[i])
			accountsUuid = append(accountsUuid, this.M_accounts[i])
		}
	}
	this.m_accounts = accounts_
	this.M_accounts = accountsUuid
}

/** Actions **/
func (this *Role) GetActions() []*Action{
	return this.m_actions
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
	}else{
		for i:=0; i < len(this.m_actions); i++ {
			if this.m_actions[i].GetUUID() == ref.(*Action).GetUUID() {
				return
			}
		}
		isExist := false
		for i:=0; i < len(this.M_actions); i++ {
			if this.M_actions[i] == ref.(*Action).GetUUID() {
				isExist = true
			}
		}
		this.m_actions = append(this.m_actions, ref.(*Action))
	if !isExist {
		this.M_actions = append(this.M_actions, ref.(*Action).GetUUID())
	}
	}
}

/** Remove reference Actions **/
func (this *Role) RemoveActions(ref interface{}){
	toDelete := ref.(*Action)
	actions_ := make([]*Action, 0)
	actionsUuid := make([]string, 0)
	for i := 0; i < len(this.m_actions); i++ {
		if toDelete.GetUUID() != this.m_actions[i].GetUUID() {
			actions_ = append(actions_, this.m_actions[i])
			actionsUuid = append(actionsUuid, this.M_actions[i])
		}
	}
	this.m_actions = actions_
	this.M_actions = actionsUuid
}

/** Entities **/
func (this *Role) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Role) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Role) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
		}
	}
}
