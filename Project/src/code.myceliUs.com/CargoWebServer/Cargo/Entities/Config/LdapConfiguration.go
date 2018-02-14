// +build Config

package Config

import(
	"encoding/xml"
)

type LdapConfiguration struct{

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

	/** members of Configuration **/
	M_id string

	/** members of LdapConfiguration **/
	M_hostName string
	M_ipv4 string
	M_port int
	M_user string
	M_pwd string
	M_domain string
	M_searchBase string


	/** Associations **/
	m_parentPtr *Configurations
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for LdapConfiguration **/
type XsdLdapConfiguration struct {
	XMLName xml.Name	`xml:"ldapConfiguration"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_hostName	string	`xml:"hostName,attr"`
	M_ipv4	string	`xml:"ipv4,attr"`
	M_port	int	`xml:"port,attr"`
	M_user	string	`xml:"user,attr"`
	M_pwd	string	`xml:"pwd,attr"`
	M_domain	string	`xml:"domain,attr"`
	M_searchBase	string	`xml:"searchBase,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *LdapConfiguration) GetUuid() string{
	return this.UUID
}
func (this *LdapConfiguration) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *LdapConfiguration) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *LdapConfiguration) GetTypeName() string{
	this.TYPENAME = "Config.LdapConfiguration"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *LdapConfiguration) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *LdapConfiguration) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *LdapConfiguration) GetParentLnk() string{
	return this.ParentLnk
}
func (this *LdapConfiguration) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *LdapConfiguration) IsNeedSave() bool{
	return this.NeedSave
}


/** Id **/
func (this *LdapConfiguration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *LdapConfiguration) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** HostName **/
func (this *LdapConfiguration) GetHostName() string{
	return this.M_hostName
}

/** Init reference HostName **/
func (this *LdapConfiguration) SetHostName(ref interface{}){
	if this.M_hostName != ref.(string) {
		this.M_hostName = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference HostName **/

/** Ipv4 **/
func (this *LdapConfiguration) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *LdapConfiguration) SetIpv4(ref interface{}){
	if this.M_ipv4 != ref.(string) {
		this.M_ipv4 = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Ipv4 **/

/** Port **/
func (this *LdapConfiguration) GetPort() int{
	return this.M_port
}

/** Init reference Port **/
func (this *LdapConfiguration) SetPort(ref interface{}){
	if this.M_port != ref.(int) {
		this.M_port = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference Port **/

/** User **/
func (this *LdapConfiguration) GetUser() string{
	return this.M_user
}

/** Init reference User **/
func (this *LdapConfiguration) SetUser(ref interface{}){
	if this.M_user != ref.(string) {
		this.M_user = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference User **/

/** Pwd **/
func (this *LdapConfiguration) GetPwd() string{
	return this.M_pwd
}

/** Init reference Pwd **/
func (this *LdapConfiguration) SetPwd(ref interface{}){
	if this.M_pwd != ref.(string) {
		this.M_pwd = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Pwd **/

/** Domain **/
func (this *LdapConfiguration) GetDomain() string{
	return this.M_domain
}

/** Init reference Domain **/
func (this *LdapConfiguration) SetDomain(ref interface{}){
	if this.M_domain != ref.(string) {
		this.M_domain = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Domain **/

/** SearchBase **/
func (this *LdapConfiguration) GetSearchBase() string{
	return this.M_searchBase
}

/** Init reference SearchBase **/
func (this *LdapConfiguration) SetSearchBase(ref interface{}){
	if this.M_searchBase != ref.(string) {
		this.M_searchBase = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference SearchBase **/

/** Parent **/
func (this *LdapConfiguration) GetParentPtr() *Configurations{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *LdapConfiguration) SetParentPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_parentPtr != ref.(string) {
			this.M_parentPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_parentPtr != ref.(*Configurations).GetUuid() {
			this.M_parentPtr = ref.(*Configurations).GetUuid()
			this.NeedSave = true
		}
		this.m_parentPtr = ref.(*Configurations)
	}
}

/** Remove reference Parent **/
func (this *LdapConfiguration) RemoveParentPtr(ref interface{}){
	toDelete := ref.(*Configurations)
	if this.m_parentPtr!= nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
