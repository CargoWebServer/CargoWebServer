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
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

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
/** UUID **/
func (this *LdapConfiguration) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *LdapConfiguration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *LdapConfiguration) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** HostName **/
func (this *LdapConfiguration) GetHostName() string{
	return this.M_hostName
}

/** Init reference HostName **/
func (this *LdapConfiguration) SetHostName(ref interface{}){
	this.NeedSave = true
	this.M_hostName = ref.(string)
}

/** Remove reference HostName **/

/** Ipv4 **/
func (this *LdapConfiguration) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *LdapConfiguration) SetIpv4(ref interface{}){
	this.NeedSave = true
	this.M_ipv4 = ref.(string)
}

/** Remove reference Ipv4 **/

/** Port **/
func (this *LdapConfiguration) GetPort() int{
	return this.M_port
}

/** Init reference Port **/
func (this *LdapConfiguration) SetPort(ref interface{}){
	this.NeedSave = true
	this.M_port = ref.(int)
}

/** Remove reference Port **/

/** User **/
func (this *LdapConfiguration) GetUser() string{
	return this.M_user
}

/** Init reference User **/
func (this *LdapConfiguration) SetUser(ref interface{}){
	this.NeedSave = true
	this.M_user = ref.(string)
}

/** Remove reference User **/

/** Pwd **/
func (this *LdapConfiguration) GetPwd() string{
	return this.M_pwd
}

/** Init reference Pwd **/
func (this *LdapConfiguration) SetPwd(ref interface{}){
	this.NeedSave = true
	this.M_pwd = ref.(string)
}

/** Remove reference Pwd **/

/** Domain **/
func (this *LdapConfiguration) GetDomain() string{
	return this.M_domain
}

/** Init reference Domain **/
func (this *LdapConfiguration) SetDomain(ref interface{}){
	this.NeedSave = true
	this.M_domain = ref.(string)
}

/** Remove reference Domain **/

/** SearchBase **/
func (this *LdapConfiguration) GetSearchBase() string{
	return this.M_searchBase
}

/** Init reference SearchBase **/
func (this *LdapConfiguration) SetSearchBase(ref interface{}){
	this.NeedSave = true
	this.M_searchBase = ref.(string)
}

/** Remove reference SearchBase **/

/** Parent **/
func (this *LdapConfiguration) GetParentPtr() *Configurations{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *LdapConfiguration) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*Configurations)
		this.M_parentPtr = ref.(*Configurations).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *LdapConfiguration) RemoveParentPtr(ref interface{}){
	toDelete := ref.(*Configurations)
	if this.m_parentPtr!= nil {
		if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
		}else{
			this.NeedSave = true
		}
	}
}
