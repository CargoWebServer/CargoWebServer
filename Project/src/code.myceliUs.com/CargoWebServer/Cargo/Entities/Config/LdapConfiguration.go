// +build Config

package Config

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
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
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

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
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *LdapConfiguration) SetUuid(uuid string){
	this.UUID = uuid
}

func (this *LdapConfiguration) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *LdapConfiguration) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
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

func (this *LdapConfiguration) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *LdapConfiguration) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *LdapConfiguration) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *LdapConfiguration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *LdapConfiguration) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *LdapConfiguration) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *LdapConfiguration) GetId()string{
	return this.M_id
}

func (this *LdapConfiguration) SetId(val string){
	this.M_id= val
}




func (this *LdapConfiguration) GetHostName()string{
	return this.M_hostName
}

func (this *LdapConfiguration) SetHostName(val string){
	this.M_hostName= val
}




func (this *LdapConfiguration) GetIpv4()string{
	return this.M_ipv4
}

func (this *LdapConfiguration) SetIpv4(val string){
	this.M_ipv4= val
}




func (this *LdapConfiguration) GetPort()int{
	return this.M_port
}

func (this *LdapConfiguration) SetPort(val int){
	this.M_port= val
}




func (this *LdapConfiguration) GetUser()string{
	return this.M_user
}

func (this *LdapConfiguration) SetUser(val string){
	this.M_user= val
}




func (this *LdapConfiguration) GetPwd()string{
	return this.M_pwd
}

func (this *LdapConfiguration) SetPwd(val string){
	this.M_pwd= val
}




func (this *LdapConfiguration) GetDomain()string{
	return this.M_domain
}

func (this *LdapConfiguration) SetDomain(val string){
	this.M_domain= val
}




func (this *LdapConfiguration) GetSearchBase()string{
	return this.M_searchBase
}

func (this *LdapConfiguration) SetSearchBase(val string){
	this.M_searchBase= val
}




func (this *LdapConfiguration) GetParentPtr()*Configurations{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*Configurations)
	}
	return nil
}

func (this *LdapConfiguration) SetParentPtr(val *Configurations){
	this.M_parentPtr= val.GetUuid()
	this.setEntity(this)
}


func (this *LdapConfiguration) ResetParentPtr(){
	this.M_parentPtr= ""
}

