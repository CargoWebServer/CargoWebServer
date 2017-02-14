// +build Config

package Config

import(
"encoding/xml"
)

type ServiceConfiguration struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Configuration **/
	M_id string

	/** members of ServiceConfiguration **/
	M_hostName string
	M_ipv4 string
	M_port int
	M_user string
	M_pwd string
	M_start bool


	/** Associations **/
	m_parentPtr *Configurations
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for ServiceConfiguration **/
type XsdServiceConfiguration struct {
	XMLName xml.Name	`xml:"serviceConfiguration"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_hostName	string	`xml:"hostName,attr"`
	M_ipv4	string	`xml:"ipv4,attr"`
	M_port	int	`xml:"port,attr"`
	M_user	string	`xml:"user,attr"`
	M_pwd	string	`xml:"pwd,attr"`
	M_start	bool	`xml:"start,attr"`

}
/** UUID **/
func (this *ServiceConfiguration) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *ServiceConfiguration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ServiceConfiguration) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** HostName **/
func (this *ServiceConfiguration) GetHostName() string{
	return this.M_hostName
}

/** Init reference HostName **/
func (this *ServiceConfiguration) SetHostName(ref interface{}){
	this.NeedSave = true
	this.M_hostName = ref.(string)
}

/** Remove reference HostName **/

/** Ipv4 **/
func (this *ServiceConfiguration) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *ServiceConfiguration) SetIpv4(ref interface{}){
	this.NeedSave = true
	this.M_ipv4 = ref.(string)
}

/** Remove reference Ipv4 **/

/** Port **/
func (this *ServiceConfiguration) GetPort() int{
	return this.M_port
}

/** Init reference Port **/
func (this *ServiceConfiguration) SetPort(ref interface{}){
	this.NeedSave = true
	this.M_port = ref.(int)
}

/** Remove reference Port **/

/** User **/
func (this *ServiceConfiguration) GetUser() string{
	return this.M_user
}

/** Init reference User **/
func (this *ServiceConfiguration) SetUser(ref interface{}){
	this.NeedSave = true
	this.M_user = ref.(string)
}

/** Remove reference User **/

/** Pwd **/
func (this *ServiceConfiguration) GetPwd() string{
	return this.M_pwd
}

/** Init reference Pwd **/
func (this *ServiceConfiguration) SetPwd(ref interface{}){
	this.NeedSave = true
	this.M_pwd = ref.(string)
}

/** Remove reference Pwd **/

/** Start **/
func (this *ServiceConfiguration) GetStart() bool{
	return this.M_start
}

/** Init reference Start **/
func (this *ServiceConfiguration) SetStart(ref interface{}){
	this.NeedSave = true
	this.M_start = ref.(bool)
}

/** Remove reference Start **/

/** Parent **/
func (this *ServiceConfiguration) GetParentPtr() *Configurations{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *ServiceConfiguration) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*Configurations)
		this.M_parentPtr = ref.(*Configurations).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *ServiceConfiguration) RemoveParentPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Configurations)
	if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
		this.m_parentPtr = nil
		this.M_parentPtr = ""
	}
}
