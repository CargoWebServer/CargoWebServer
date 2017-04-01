// +build CargoEntities

package CargoEntities

import(
"encoding/xml"
)

type User struct{

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

	/** members of Entity **/
	M_id string

	/** members of User **/
	M_firstName string
	M_lastName string
	M_middle string
	M_phone string
	M_email string
	m_memberOfRef []*Group
	/** If the ref is a string and not an object **/
	M_memberOfRef []string
	m_accounts []*Account
	/** If the ref is a string and not an object **/
	M_accounts []string


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for User **/
type XsdUser struct {
	XMLName xml.Name	`xml:"userRef"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_memberOfRef	[]string	`xml:"memberOfRef"`
	M_firstName	string	`xml:"firstName,attr"`
	M_lastName	string	`xml:"lastName,attr"`
	M_middle	string	`xml:"middle,attr"`
	M_email	string	`xml:"email,attr"`
	M_phone	string	`xml:"phone,attr"`

}
/** UUID **/
func (this *User) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *User) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *User) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** FirstName **/
func (this *User) GetFirstName() string{
	return this.M_firstName
}

/** Init reference FirstName **/
func (this *User) SetFirstName(ref interface{}){
	this.NeedSave = true
	this.M_firstName = ref.(string)
}

/** Remove reference FirstName **/

/** LastName **/
func (this *User) GetLastName() string{
	return this.M_lastName
}

/** Init reference LastName **/
func (this *User) SetLastName(ref interface{}){
	this.NeedSave = true
	this.M_lastName = ref.(string)
}

/** Remove reference LastName **/

/** Middle **/
func (this *User) GetMiddle() string{
	return this.M_middle
}

/** Init reference Middle **/
func (this *User) SetMiddle(ref interface{}){
	this.NeedSave = true
	this.M_middle = ref.(string)
}

/** Remove reference Middle **/

/** Phone **/
func (this *User) GetPhone() string{
	return this.M_phone
}

/** Init reference Phone **/
func (this *User) SetPhone(ref interface{}){
	this.NeedSave = true
	this.M_phone = ref.(string)
}

/** Remove reference Phone **/

/** Email **/
func (this *User) GetEmail() string{
	return this.M_email
}

/** Init reference Email **/
func (this *User) SetEmail(ref interface{}){
	this.NeedSave = true
	this.M_email = ref.(string)
}

/** Remove reference Email **/

/** MemberOfRef **/
func (this *User) GetMemberOfRef() []*Group{
	return this.m_memberOfRef
}

/** Init reference MemberOfRef **/
func (this *User) SetMemberOfRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_memberOfRef); i++ {
			if this.M_memberOfRef[i] == refStr {
				return
			}
		}
		this.M_memberOfRef = append(this.M_memberOfRef, ref.(string))
	}else{
		this.RemoveMemberOfRef(ref)
		this.m_memberOfRef = append(this.m_memberOfRef, ref.(*Group))
		this.M_memberOfRef = append(this.M_memberOfRef, ref.(Entity).GetUUID())
	}
}

/** Remove reference MemberOfRef **/
func (this *User) RemoveMemberOfRef(ref interface{}){
	toDelete := ref.(Entity)
	memberOfRef_ := make([]*Group, 0)
	memberOfRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_memberOfRef); i++ {
		if toDelete.GetUUID() != this.m_memberOfRef[i].GetUUID() {
			memberOfRef_ = append(memberOfRef_, this.m_memberOfRef[i])
			memberOfRefUuid = append(memberOfRefUuid, this.M_memberOfRef[i])
		}else{
			this.NeedSave = true
		}
	}
	this.m_memberOfRef = memberOfRef_
	this.M_memberOfRef = memberOfRefUuid
}

/** Accounts **/
func (this *User) GetAccounts() []*Account{
	return this.m_accounts
}

/** Init reference Accounts **/
func (this *User) SetAccounts(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_accounts); i++ {
			if this.M_accounts[i] == refStr {
				return
			}
		}
		this.M_accounts = append(this.M_accounts, ref.(string))
	}else{
		this.RemoveAccounts(ref)
		this.m_accounts = append(this.m_accounts, ref.(*Account))
		this.M_accounts = append(this.M_accounts, ref.(Entity).GetUUID())
	}
}

/** Remove reference Accounts **/
func (this *User) RemoveAccounts(ref interface{}){
	toDelete := ref.(Entity)
	accounts_ := make([]*Account, 0)
	accountsUuid := make([]string, 0)
	for i := 0; i < len(this.m_accounts); i++ {
		if toDelete.GetUUID() != this.m_accounts[i].GetUUID() {
			accounts_ = append(accounts_, this.m_accounts[i])
			accountsUuid = append(accountsUuid, this.M_accounts[i])
		}else{
			this.NeedSave = true
		}
	}
	this.m_accounts = accounts_
	this.M_accounts = accountsUuid
}

/** Entities **/
func (this *User) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *User) SetEntitiesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *User) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
		}else{
			this.NeedSave = true
		}
	}
}
