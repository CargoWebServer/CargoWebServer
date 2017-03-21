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
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Permission **/
	M_pattern string
	M_type PermissionType


	/** Associations **/
	m_accountPtr *Account
	/** If the ref is a string and not an object **/
	M_accountPtr string
}

/** Xml parser for Permission **/
type XsdPermission struct {
	XMLName xml.Name	`xml:"permission"`
	M_pattern	string	`xml:"pattern,attr"`
	M_type	string	`xml:"type,attr"`

}
/** UUID **/
func (this *Permission) GetUUID() string{
	return this.UUID
}

/** Pattern **/
func (this *Permission) GetPattern() string{
	return this.M_pattern
}

/** Init reference Pattern **/
func (this *Permission) SetPattern(ref interface{}){
	this.NeedSave = true
	this.M_pattern = ref.(string)
}

/** Remove reference Pattern **/

/** Type **/
func (this *Permission) GetType() PermissionType{
	return this.M_type
}

/** Init reference Type **/
func (this *Permission) SetType(ref interface{}){
	this.NeedSave = true
	this.M_type = ref.(PermissionType)
}

/** Remove reference Type **/

/** Account **/
func (this *Permission) GetAccountPtr() *Account{
	return this.m_accountPtr
}

/** Init reference Account **/
func (this *Permission) SetAccountPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_accountPtr = ref.(string)
	}else{
		this.m_accountPtr = ref.(*Account)
		this.M_accountPtr = ref.(Entity).GetUUID()
	}
}

/** Remove reference Account **/
func (this *Permission) RemoveAccountPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Entity)
	if toDelete.GetUUID() == this.m_accountPtr.GetUUID() {
		this.m_accountPtr = nil
		this.M_accountPtr = ""
	}
}
