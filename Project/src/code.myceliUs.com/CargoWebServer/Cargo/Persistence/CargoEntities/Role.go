package CargoEntities

import(
"encoding/xml"
)

type Role struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Role **/
	M_id string
	m_accountsRef []*Account
	/** If the ref is a string and not an object **/
	M_accountsRef []string


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
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** AccountsRef **/
func (this *Role) GetAccountsRef() []*Account{
	return this.m_accountsRef
}

/** Init reference AccountsRef **/
func (this *Role) SetAccountsRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_accountsRef); i++ {
			if this.M_accountsRef[i] == refStr {
				return
			}
		}
		this.M_accountsRef = append(this.M_accountsRef, ref.(string))
	}else{
		this.RemoveAccountsRef(ref)
		this.m_accountsRef = append(this.m_accountsRef, ref.(*Account))
		this.M_accountsRef = append(this.M_accountsRef, ref.(Entity).GetUUID())
	}
}

/** Remove reference AccountsRef **/
func (this *Role) RemoveAccountsRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Entity)
	accountsRef_ := make([]*Account, 0)
	accountsRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_accountsRef); i++ {
		if toDelete.GetUUID() != this.m_accountsRef[i].GetUUID() {
			accountsRef_ = append(accountsRef_, this.m_accountsRef[i])
			accountsRefUuid = append(accountsRefUuid, this.m_accountsRef[i].GetUUID())
		}
	}
	this.m_accountsRef = accountsRef_
	this.M_accountsRef = accountsRefUuid
}

/** Entities **/
func (this *Role) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Role) SetEntitiesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *Role) RemoveEntitiesPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Entities)
	if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
		this.m_entitiesPtr = nil
		this.M_entitiesPtr = ""
	}
}
