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
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

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
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *ServiceConfiguration) SetUuid(uuid string){
	this.NeedSave = this.UUID == uuid
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

func (this *ServiceConfiguration) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ServiceConfiguration) GetChilds() []interface{}{
	var childs []interface{}
	return childs
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
/** Use it the set the entity on the cache. **/
func (this *ServiceConfiguration) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *ServiceConfiguration) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *ServiceConfiguration) GetId()string{
	return this.M_id
}

func (this *ServiceConfiguration) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}




func (this *ServiceConfiguration) GetHostName()string{
	return this.M_hostName
}

func (this *ServiceConfiguration) SetHostName(val string){
	this.NeedSave = this.M_hostName== val
	this.M_hostName= val
}




func (this *ServiceConfiguration) GetIpv4()string{
	return this.M_ipv4
}

func (this *ServiceConfiguration) SetIpv4(val string){
	this.NeedSave = this.M_ipv4== val
	this.M_ipv4= val
}




func (this *ServiceConfiguration) GetPort()int{
	return this.M_port
}

func (this *ServiceConfiguration) SetPort(val int){
	this.NeedSave = this.M_port== val
	this.M_port= val
}




func (this *ServiceConfiguration) GetUser()string{
	return this.M_user
}

func (this *ServiceConfiguration) SetUser(val string){
	this.NeedSave = this.M_user== val
	this.M_user= val
}




func (this *ServiceConfiguration) GetPwd()string{
	return this.M_pwd
}

func (this *ServiceConfiguration) SetPwd(val string){
	this.NeedSave = this.M_pwd== val
	this.M_pwd= val
}




func (this *ServiceConfiguration) IsStart()bool{
	return this.M_start
}

func (this *ServiceConfiguration) SetStart(val bool){
	this.NeedSave = this.M_start== val
	this.M_start= val
}




func (this *ServiceConfiguration) GetParentPtr()*Configurations{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*Configurations)
	}
	return nil
}

func (this *ServiceConfiguration) SetParentPtr(val *Configurations){
	this.NeedSave = this.M_parentPtr != val.GetUuid()
	this.M_parentPtr= val.GetUuid()
}


func (this *ServiceConfiguration) ResetParentPtr(){
	this.M_parentPtr= ""
}

