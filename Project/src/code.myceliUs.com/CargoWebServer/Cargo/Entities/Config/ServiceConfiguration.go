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
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

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
/***************** Entity **************************/

/** UUID **/
func (this *ServiceConfiguration) GetUuid() string{
	return this.UUID
}
func (this *ServiceConfiguration) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *ServiceConfiguration) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *ServiceConfiguration) GetTypeName() string{
	this.TYPENAME = "Config.ServiceConfiguration"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *ServiceConfiguration) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *ServiceConfiguration) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ServiceConfiguration) GetParentLnk() string{
	return this.ParentLnk
}
func (this *ServiceConfiguration) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *ServiceConfiguration) IsNeedSave() bool{
	return this.NeedSave
}
func (this *ServiceConfiguration) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *ServiceConfiguration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *ServiceConfiguration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ServiceConfiguration) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** HostName **/
func (this *ServiceConfiguration) GetHostName() string{
	return this.M_hostName
}

/** Init reference HostName **/
func (this *ServiceConfiguration) SetHostName(ref interface{}){
	if this.M_hostName != ref.(string) {
		this.M_hostName = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference HostName **/

/** Ipv4 **/
func (this *ServiceConfiguration) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *ServiceConfiguration) SetIpv4(ref interface{}){
	if this.M_ipv4 != ref.(string) {
		this.M_ipv4 = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Ipv4 **/

/** Port **/
func (this *ServiceConfiguration) GetPort() int{
	return this.M_port
}

/** Init reference Port **/
func (this *ServiceConfiguration) SetPort(ref interface{}){
	if this.M_port != ref.(int) {
		this.M_port = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference Port **/

/** User **/
func (this *ServiceConfiguration) GetUser() string{
	return this.M_user
}

/** Init reference User **/
func (this *ServiceConfiguration) SetUser(ref interface{}){
	if this.M_user != ref.(string) {
		this.M_user = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference User **/

/** Pwd **/
func (this *ServiceConfiguration) GetPwd() string{
	return this.M_pwd
}

/** Init reference Pwd **/
func (this *ServiceConfiguration) SetPwd(ref interface{}){
	if this.M_pwd != ref.(string) {
		this.M_pwd = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Pwd **/

/** Start **/
func (this *ServiceConfiguration) GetStart() bool{
	return this.M_start
}

/** Init reference Start **/
func (this *ServiceConfiguration) SetStart(ref interface{}){
	if this.M_start != ref.(bool) {
		this.M_start = ref.(bool)
		this.NeedSave = true
	}
}

/** Remove reference Start **/

/** Parent **/
func (this *ServiceConfiguration) GetParentPtr() *Configurations{
	if this.m_parentPtr == nil {
		entity, err := this.getEntityByUuid(this.M_parentPtr)
		if err == nil {
			this.m_parentPtr = entity.(*Configurations)
		}
	}
	return this.m_parentPtr
}
func (this *ServiceConfiguration) GetParentPtrStr() string{
	return this.M_parentPtr
}

/** Init reference Parent **/
func (this *ServiceConfiguration) SetParentPtr(ref interface{}){
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
func (this *ServiceConfiguration) RemoveParentPtr(ref interface{}){
	toDelete := ref.(*Configurations)
	if this.m_parentPtr!= nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
