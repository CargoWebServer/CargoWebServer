// +build Config

package Config

import(
	"encoding/xml"
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
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

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
	m_parentPtr *Configurations
	/** If the ref is a string and not an object **/
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
	return this.UUID
}
func (this *DataStoreConfiguration) SetUuid(uuid string){
	this.UUID = uuid
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

/** Evaluate if an entity needs to be saved. **/
func (this *DataStoreConfiguration) IsNeedSave() bool{
	return this.NeedSave
}
func (this *DataStoreConfiguration) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *DataStoreConfiguration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *DataStoreConfiguration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *DataStoreConfiguration) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** DataStoreType **/
func (this *DataStoreConfiguration) GetDataStoreType() DataStoreType{
	return this.M_dataStoreType
}

/** Init reference DataStoreType **/
func (this *DataStoreConfiguration) SetDataStoreType(ref interface{}){
	if this.M_dataStoreType != ref.(DataStoreType) {
		this.M_dataStoreType = ref.(DataStoreType)
		this.NeedSave = true
	}
}

/** Remove reference DataStoreType **/

/** DataStoreVendor **/
func (this *DataStoreConfiguration) GetDataStoreVendor() DataStoreVendor{
	return this.M_dataStoreVendor
}

/** Init reference DataStoreVendor **/
func (this *DataStoreConfiguration) SetDataStoreVendor(ref interface{}){
	if this.M_dataStoreVendor != ref.(DataStoreVendor) {
		this.M_dataStoreVendor = ref.(DataStoreVendor)
		this.NeedSave = true
	}
}

/** Remove reference DataStoreVendor **/

/** TextEncoding **/
func (this *DataStoreConfiguration) GetTextEncoding() Encoding{
	return this.M_textEncoding
}

/** Init reference TextEncoding **/
func (this *DataStoreConfiguration) SetTextEncoding(ref interface{}){
	if this.M_textEncoding != ref.(Encoding) {
		this.M_textEncoding = ref.(Encoding)
		this.NeedSave = true
	}
}

/** Remove reference TextEncoding **/

/** StoreName **/
func (this *DataStoreConfiguration) GetStoreName() string{
	return this.M_storeName
}

/** Init reference StoreName **/
func (this *DataStoreConfiguration) SetStoreName(ref interface{}){
	if this.M_storeName != ref.(string) {
		this.M_storeName = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference StoreName **/

/** HostName **/
func (this *DataStoreConfiguration) GetHostName() string{
	return this.M_hostName
}

/** Init reference HostName **/
func (this *DataStoreConfiguration) SetHostName(ref interface{}){
	if this.M_hostName != ref.(string) {
		this.M_hostName = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference HostName **/

/** Ipv4 **/
func (this *DataStoreConfiguration) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *DataStoreConfiguration) SetIpv4(ref interface{}){
	if this.M_ipv4 != ref.(string) {
		this.M_ipv4 = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Ipv4 **/

/** User **/
func (this *DataStoreConfiguration) GetUser() string{
	return this.M_user
}

/** Init reference User **/
func (this *DataStoreConfiguration) SetUser(ref interface{}){
	if this.M_user != ref.(string) {
		this.M_user = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference User **/

/** Pwd **/
func (this *DataStoreConfiguration) GetPwd() string{
	return this.M_pwd
}

/** Init reference Pwd **/
func (this *DataStoreConfiguration) SetPwd(ref interface{}){
	if this.M_pwd != ref.(string) {
		this.M_pwd = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Pwd **/

/** Port **/
func (this *DataStoreConfiguration) GetPort() int{
	return this.M_port
}

/** Init reference Port **/
func (this *DataStoreConfiguration) SetPort(ref interface{}){
	if this.M_port != ref.(int) {
		this.M_port = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference Port **/

/** Parent **/
func (this *DataStoreConfiguration) GetParentPtr() *Configurations{
	if this.m_parentPtr == nil {
		entity, err := this.getEntityByUuid(this.M_parentPtr)
		if err == nil {
			this.m_parentPtr = entity.(*Configurations)
		}
	}
	return this.m_parentPtr
}
func (this *DataStoreConfiguration) GetParentPtrStr() string{
	return this.M_parentPtr
}

/** Init reference Parent **/
func (this *DataStoreConfiguration) SetParentPtr(ref interface{}){
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
func (this *DataStoreConfiguration) RemoveParentPtr(ref interface{}){
	toDelete := ref.(*Configurations)
	if this.m_parentPtr!= nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
