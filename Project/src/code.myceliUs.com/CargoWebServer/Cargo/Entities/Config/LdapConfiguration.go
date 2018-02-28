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
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

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
func (this *LdapConfiguration) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *LdapConfiguration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

func (this *LdapConfiguration) GetId()string{
	return this.M_id
}

func (this *LdapConfiguration) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *LdapConfiguration) GetHostName()string{
	return this.M_hostName
}

func (this *LdapConfiguration) SetHostName(val string){
	this.NeedSave = this.M_hostName== val
	this.M_hostName= val
}


func (this *LdapConfiguration) GetIpv4()string{
	return this.M_ipv4
}

func (this *LdapConfiguration) SetIpv4(val string){
	this.NeedSave = this.M_ipv4== val
	this.M_ipv4= val
}


func (this *LdapConfiguration) GetPort()int{
	return this.M_port
}

func (this *LdapConfiguration) SetPort(val int){
	this.NeedSave = this.M_port== val
	this.M_port= val
}


func (this *LdapConfiguration) GetUser()string{
	return this.M_user
}

func (this *LdapConfiguration) SetUser(val string){
	this.NeedSave = this.M_user== val
	this.M_user= val
}


func (this *LdapConfiguration) GetPwd()string{
	return this.M_pwd
}

func (this *LdapConfiguration) SetPwd(val string){
	this.NeedSave = this.M_pwd== val
	this.M_pwd= val
}


func (this *LdapConfiguration) GetDomain()string{
	return this.M_domain
}

func (this *LdapConfiguration) SetDomain(val string){
	this.NeedSave = this.M_domain== val
	this.M_domain= val
}


func (this *LdapConfiguration) GetSearchBase()string{
	return this.M_searchBase
}

func (this *LdapConfiguration) SetSearchBase(val string){
	this.NeedSave = this.M_searchBase== val
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
}

func (this *LdapConfiguration) ResetParentPtr(){
	this.M_parentPtr= ""
}

