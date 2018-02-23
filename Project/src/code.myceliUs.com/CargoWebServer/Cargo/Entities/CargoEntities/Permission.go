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
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

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
/***************** Entity **************************/

/** UUID **/
func (this *Permission) GetUuid() string{
	return this.UUID
}
func (this *Permission) SetUuid(uuid string){
	this.UUID = uuid
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

/** Evaluate if an entity needs to be saved. **/
func (this *Permission) IsNeedSave() bool{
	return this.NeedSave
}
func (this *Permission) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Permission) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *Permission) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Permission) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Types **/
func (this *Permission) GetTypes() int{
	return this.M_types
}

/** Init reference Types **/
func (this *Permission) SetTypes(ref interface{}){
	if this.M_types != ref.(int) {
		this.M_types = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference Types **/

/** AccountsRef **/
func (this *Permission) GetAccountsRef() []*Account{
	if this.m_accountsRef == nil {
		this.m_accountsRef = make([]*Account, 0)
		for i := 0; i < len(this.M_accountsRef); i++ {
			entity, err := this.getEntityByUuid(this.M_accountsRef[i])
			if err == nil {
				this.m_accountsRef = append(this.m_accountsRef, entity.(*Account))
			}
		}
	}
	return this.m_accountsRef
}
func (this *Permission) GetAccountsRefStr() []string{
	return this.M_accountsRef
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
		this.NeedSave = true
	}else{
		for i:=0; i < len(this.m_accountsRef); i++ {
			if this.m_accountsRef[i].GetUuid() == ref.(*Account).GetUuid() {
				return
			}
		}
		isExist := false
		for i:=0; i < len(this.M_accountsRef); i++ {
			if this.M_accountsRef[i] == ref.(*Account).GetUuid() {
				isExist = true
			}
		}
		this.m_accountsRef = append(this.m_accountsRef, ref.(*Account))
	if !isExist {
		this.M_accountsRef = append(this.M_accountsRef, ref.(Entity).GetUuid())
		this.NeedSave = true
	}
	}
}

/** Remove reference AccountsRef **/
func (this *Permission) RemoveAccountsRef(ref interface{}){
	toDelete := ref.(Entity)
	accountsRef_ := make([]*Account, 0)
	accountsRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_accountsRef); i++ {
		if toDelete.GetUuid() != this.m_accountsRef[i].GetUuid() {
			accountsRef_ = append(accountsRef_, this.m_accountsRef[i])
			accountsRefUuid = append(accountsRefUuid, this.M_accountsRef[i])
		}else{
			this.NeedSave = true
		}
	}
	this.m_accountsRef = accountsRef_
	this.M_accountsRef = accountsRefUuid
}

/** Entities **/
func (this *Permission) GetEntitiesPtr() *Entities{
	if this.m_entitiesPtr == nil {
		entity, err := this.getEntityByUuid(this.M_entitiesPtr)
		if err == nil {
			this.m_entitiesPtr = entity.(*Entities)
		}
	}
	return this.m_entitiesPtr
}
func (this *Permission) GetEntitiesPtrStr() string{
	return this.M_entitiesPtr
}

/** Init reference Entities **/
func (this *Permission) SetEntitiesPtr(ref interface{}){
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
func (this *Permission) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
