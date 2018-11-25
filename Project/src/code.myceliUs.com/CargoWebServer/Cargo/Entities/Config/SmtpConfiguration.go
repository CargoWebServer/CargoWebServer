// +build Config

package Config

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type SmtpConfiguration struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** keep track if the entity has change over time. **/
	needSave bool
	/** Keep reference to entity that made use of thit entity **/
	Referenced []string
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of Configuration **/
	M_id string

	/** members of SmtpConfiguration **/
	M_textEncoding Encoding
	M_hostName string
	M_ipv4 string
	M_port int
	M_user string
	M_pwd string


	/** Associations **/
	M_parentPtr string
}

/** Xml parser for SmtpConfiguration **/
type XsdSmtpConfiguration struct {
	XMLName xml.Name	`xml:"smtpConfiguration"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_hostName	string	`xml:"hostName,attr"`
	M_ipv4	string	`xml:"ipv4,attr"`
	M_port	int	`xml:"port,attr"`
	M_user	string	`xml:"user,attr"`
	M_pwd	string	`xml:"pwd,attr"`
	M_textEncoding	string	`xml:"textEncoding,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *SmtpConfiguration) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *SmtpConfiguration) SetUuid(uuid string){
	this.UUID = uuid
}

/** Need save **/
func (this *SmtpConfiguration) IsNeedSave() bool{
	return this.needSave
}
func (this *SmtpConfiguration) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *SmtpConfiguration) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *SmtpConfiguration) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *SmtpConfiguration) RemoveReferenced(uuid string, field string) {
	if this.Referenced == nil {
		return
	}
	referenced := make([]string, 0)
	for i := 0; i < len(this.Referenced); i++ {
		if this.Referenced[i] != uuid+":"+field {
			referenced = append(referenced, uuid+":"+field)
		}
	}
	this.Referenced = referenced
}

func (this *SmtpConfiguration) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *SmtpConfiguration) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *SmtpConfiguration) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *SmtpConfiguration) GetTypeName() string{
	this.TYPENAME = "Config.SmtpConfiguration"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *SmtpConfiguration) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *SmtpConfiguration) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *SmtpConfiguration) GetParentLnk() string{
	return this.ParentLnk
}
func (this *SmtpConfiguration) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *SmtpConfiguration) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *SmtpConfiguration) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *SmtpConfiguration) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *SmtpConfiguration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *SmtpConfiguration) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *SmtpConfiguration) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *SmtpConfiguration) GetId()string{
	return this.M_id
}

func (this *SmtpConfiguration) SetId(val string){
	this.M_id= val
}




func (this *SmtpConfiguration) GetTextEncoding()Encoding{
	return this.M_textEncoding
}

func (this *SmtpConfiguration) SetTextEncoding(val Encoding){
	this.M_textEncoding= val
}


func (this *SmtpConfiguration) ResetTextEncoding(){
	this.M_textEncoding= 0
	this.setEntity(this)
}


func (this *SmtpConfiguration) GetHostName()string{
	return this.M_hostName
}

func (this *SmtpConfiguration) SetHostName(val string){
	this.M_hostName= val
}




func (this *SmtpConfiguration) GetIpv4()string{
	return this.M_ipv4
}

func (this *SmtpConfiguration) SetIpv4(val string){
	this.M_ipv4= val
}




func (this *SmtpConfiguration) GetPort()int{
	return this.M_port
}

func (this *SmtpConfiguration) SetPort(val int){
	this.M_port= val
}




func (this *SmtpConfiguration) GetUser()string{
	return this.M_user
}

func (this *SmtpConfiguration) SetUser(val string){
	this.M_user= val
}




func (this *SmtpConfiguration) GetPwd()string{
	return this.M_pwd
}

func (this *SmtpConfiguration) SetPwd(val string){
	this.M_pwd= val
}




func (this *SmtpConfiguration) GetParentPtr()*Configurations{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*Configurations)
	}
	return nil
}

func (this *SmtpConfiguration) SetParentPtr(val *Configurations){
	this.M_parentPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *SmtpConfiguration) ResetParentPtr(){
	this.M_parentPtr= ""
	this.setEntity(this)
}

