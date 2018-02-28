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
	M_accountsRef []string


	/** Associations **/
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

func (this *Permission) GetId()string{
	return this.M_id
}

func (this *Permission) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *Permission) GetTypes()int{
	return this.M_types
}

func (this *Permission) SetTypes(val int){
	this.NeedSave = this.M_types== val
	this.M_types= val
}


func (this *Permission) GetAccountsRef()[]*Account{
	accountsRef := make([]*Account, 0)
	for i := 0; i < len(this.M_accountsRef); i++ {
		entity, err := this.getEntityByUuid(this.M_accountsRef[i])
		if err == nil {
			accountsRef = append(accountsRef, entity.(*Account))
		}
	}
	return accountsRef
}

func (this *Permission) SetAccountsRef(val []*Account){
	this.M_accountsRef= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_accountsRef=append(this.M_accountsRef, val[i].GetUuid())
	}
}

func (this *Permission) AppendAccountsRef(val *Account){
	for i:=0; i < len(this.M_accountsRef); i++{
		if this.M_accountsRef[i] == val.GetUuid() {
			return
		}
	}
	this.M_accountsRef = append(this.M_accountsRef, val.GetUuid())
}

func (this *Permission) RemoveAccountsRef(val *Account){
	accountsRef := make([]string,0)
	for i:=0; i < len(this.M_accountsRef); i++{
		if this.M_accountsRef[i] != val.GetUuid() {
			accountsRef = append(accountsRef, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_accountsRef = accountsRef
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
}

func (this *Permission) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

