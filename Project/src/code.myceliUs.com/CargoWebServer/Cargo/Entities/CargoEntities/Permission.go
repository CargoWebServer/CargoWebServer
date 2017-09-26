// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type Permission struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Permission **/
	M_id string
	M_types int
	m_accountsRef []*Account
	/** If the ref is a string and not an object **/
	M_accountsRef []string


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for Permission **/
type XsdPermission struct {
	XMLName xml.Name	`xml:"permissionsRef"`
	M_id	string	`xml:"id,attr"`
	M_types	int	`xml:"types,attr"`

}
/** UUID **/
func (this *Permission) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Permission) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Permission) SetId(ref interface{}){
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Types **/
func (this *Permission) GetTypes() int{
	return this.M_types
}

/** Init reference Types **/
func (this *Permission) SetTypes(ref interface{}){
	this.M_types = ref.(int)
}

/** Remove reference Types **/

/** AccountsRef **/
func (this *Permission) GetAccountsRef() []*Account{
	return this.m_accountsRef
}

/** Init reference AccountsRef **/
func (this *Permission) SetAccountsRef(ref interface{}){
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_accountsRef); i++ {
			if this.M_accountsRef[i] == refStr {
				return
			}
		}
		this.M_accountsRef = append(this.M_accountsRef, ref.(string))
	}else{
		for i:=0; i < len(this.m_accountsRef); i++ {
			if this.m_accountsRef[i].GetUUID() == ref.(*Account).GetUUID() {
				return
			}
		}
		isExist := false
		for i:=0; i < len(this.M_accountsRef); i++ {
			if this.M_accountsRef[i] == ref.(*Account).GetUUID() {
				isExist = true
			}
		}
		this.m_accountsRef = append(this.m_accountsRef, ref.(*Account))
	if !isExist {
		this.M_accountsRef = append(this.M_accountsRef, ref.(Entity).GetUUID())
	}
	}
}

/** Remove reference AccountsRef **/
func (this *Permission) RemoveAccountsRef(ref interface{}){
	toDelete := ref.(Entity)
	accountsRef_ := make([]*Account, 0)
	accountsRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_accountsRef); i++ {
		if toDelete.GetUUID() != this.m_accountsRef[i].GetUUID() {
			accountsRef_ = append(accountsRef_, this.m_accountsRef[i])
			accountsRefUuid = append(accountsRefUuid, this.M_accountsRef[i])
		}
	}
	this.m_accountsRef = accountsRef_
	this.M_accountsRef = accountsRefUuid
}

/** Entities **/
func (this *Permission) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Permission) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Permission) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
		}
	}
}
