// +build Config

package Config

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type DataStoreConfiguration struct{

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

	/** members of DataStoreConfiguration **/
	M_dataStoreType DataStoreType
	M_dataStoreVendor DataStoreVendor
	M_textEncoding Encoding
	M_storeName string
	M_hostName string
	M_ipv4 string
	M_user string
	M_pwd string
	M_port int


	/** Associations **/
	M_parentPtr string
}

/** Xml parser for DataStoreConfiguration **/
type XsdDataStoreConfiguration struct {
	XMLName xml.Name	`xml:"dataStoreConfiguration"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_storeName	string	`xml:"storeName,attr"`
	M_hostName	string	`xml:"hostName,attr"`
	M_ipv4	string	`xml:"ipv4,attr"`
	M_port	int	`xml:"port,attr"`
	M_user	string	`xml:"user,attr"`
	M_pwd	string	`xml:"pwd,attr"`
	M_dataStoreType	string	`xml:"dataStoreType,attr"`
	M_dataStoreVendor	string	`xml:"dataStoreVendor,attr"`
	M_textEncoding	string	`xml:"textEncoding,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *DataStoreConfiguration) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *DataStoreConfiguration) SetUuid(uuid string){
	this.UUID = uuid
}

/** Need save **/
func (this *DataStoreConfiguration) IsNeedSave() bool{
	return this.needSave
}
func (this *DataStoreConfiguration) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *DataStoreConfiguration) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *DataStoreConfiguration) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *DataStoreConfiguration) RemoveReferenced(uuid string, field string) {
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

func (this *DataStoreConfiguration) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *DataStoreConfiguration) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *DataStoreConfiguration) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *DataStoreConfiguration) GetTypeName() string{
	this.TYPENAME = "Config.DataStoreConfiguration"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *DataStoreConfiguration) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *DataStoreConfiguration) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *DataStoreConfiguration) GetParentLnk() string{
	return this.ParentLnk
}
func (this *DataStoreConfiguration) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *DataStoreConfiguration) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *DataStoreConfiguration) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *DataStoreConfiguration) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *DataStoreConfiguration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *DataStoreConfiguration) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *DataStoreConfiguration) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *DataStoreConfiguration) GetId()string{
	return this.M_id
}

func (this *DataStoreConfiguration) SetId(val string){
	this.M_id= val
}




func (this *DataStoreConfiguration) GetDataStoreType()DataStoreType{
	return this.M_dataStoreType
}

func (this *DataStoreConfiguration) SetDataStoreType(val DataStoreType){
	this.M_dataStoreType= val
}


func (this *DataStoreConfiguration) ResetDataStoreType(){
	this.M_dataStoreType= 0
	this.setEntity(this)
}


func (this *DataStoreConfiguration) GetDataStoreVendor()DataStoreVendor{
	return this.M_dataStoreVendor
}

func (this *DataStoreConfiguration) SetDataStoreVendor(val DataStoreVendor){
	this.M_dataStoreVendor= val
}


func (this *DataStoreConfiguration) ResetDataStoreVendor(){
	this.M_dataStoreVendor= 0
	this.setEntity(this)
}


func (this *DataStoreConfiguration) GetTextEncoding()Encoding{
	return this.M_textEncoding
}

func (this *DataStoreConfiguration) SetTextEncoding(val Encoding){
	this.M_textEncoding= val
}


func (this *DataStoreConfiguration) ResetTextEncoding(){
	this.M_textEncoding= 0
	this.setEntity(this)
}


func (this *DataStoreConfiguration) GetStoreName()string{
	return this.M_storeName
}

func (this *DataStoreConfiguration) SetStoreName(val string){
	this.M_storeName= val
}




func (this *DataStoreConfiguration) GetHostName()string{
	return this.M_hostName
}

func (this *DataStoreConfiguration) SetHostName(val string){
	this.M_hostName= val
}




func (this *DataStoreConfiguration) GetIpv4()string{
	return this.M_ipv4
}

func (this *DataStoreConfiguration) SetIpv4(val string){
	this.M_ipv4= val
}




func (this *DataStoreConfiguration) GetUser()string{
	return this.M_user
}

func (this *DataStoreConfiguration) SetUser(val string){
	this.M_user= val
}




func (this *DataStoreConfiguration) GetPwd()string{
	return this.M_pwd
}

func (this *DataStoreConfiguration) SetPwd(val string){
	this.M_pwd= val
}




func (this *DataStoreConfiguration) GetPort()int{
	return this.M_port
}

func (this *DataStoreConfiguration) SetPort(val int){
	this.M_port= val
}




func (this *DataStoreConfiguration) GetParentPtr()*Configurations{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*Configurations)
	}
	return nil
}

func (this *DataStoreConfiguration) SetParentPtr(val *Configurations){
	this.M_parentPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *DataStoreConfiguration) ResetParentPtr(){
	this.M_parentPtr= ""
	this.setEntity(this)
}

