package CargoConfig

import(
"encoding/xml"
)

type DataStoreConfiguration struct{

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

	/** members of DataStoreConfiguration **/
	M_dataStoreType DataStoreType
	M_dataStoreVendor DataStoreVendor
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


	M_hostName	string	`xml:"hostName,attr"`
	M_ipv4	string	`xml:"ipv4,attr"`
	M_port	int	`xml:"port,attr"`
	M_user	string	`xml:"user,attr"`
	M_pwd	string	`xml:"pwd,attr"`
	M_dataStoreType	string	`xml:"dataStoreType,attr"`
	M_dataStoreVendor	string	`xml:"dataStoreVendor,attr"`

}
/** UUID **/
func (this *DataStoreConfiguration) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *DataStoreConfiguration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *DataStoreConfiguration) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** DataStoreType **/
func (this *DataStoreConfiguration) GetDataStoreType() DataStoreType{
	return this.M_dataStoreType
}

/** Init reference DataStoreType **/
func (this *DataStoreConfiguration) SetDataStoreType(ref interface{}){
	this.NeedSave = true
	this.M_dataStoreType = ref.(DataStoreType)
}

/** Remove reference DataStoreType **/

/** DataStoreVendor **/
func (this *DataStoreConfiguration) GetDataStoreVendor() DataStoreVendor{
	return this.M_dataStoreVendor
}

/** Init reference DataStoreVendor **/
func (this *DataStoreConfiguration) SetDataStoreVendor(ref interface{}){
	this.NeedSave = true
	this.M_dataStoreVendor = ref.(DataStoreVendor)
}

/** Remove reference DataStoreVendor **/

/** HostName **/
func (this *DataStoreConfiguration) GetHostName() string{
	return this.M_hostName
}

/** Init reference HostName **/
func (this *DataStoreConfiguration) SetHostName(ref interface{}){
	this.NeedSave = true
	this.M_hostName = ref.(string)
}

/** Remove reference HostName **/

/** Ipv4 **/
func (this *DataStoreConfiguration) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *DataStoreConfiguration) SetIpv4(ref interface{}){
	this.NeedSave = true
	this.M_ipv4 = ref.(string)
}

/** Remove reference Ipv4 **/

/** User **/
func (this *DataStoreConfiguration) GetUser() string{
	return this.M_user
}

/** Init reference User **/
func (this *DataStoreConfiguration) SetUser(ref interface{}){
	this.NeedSave = true
	this.M_user = ref.(string)
}

/** Remove reference User **/

/** Pwd **/
func (this *DataStoreConfiguration) GetPwd() string{
	return this.M_pwd
}

/** Init reference Pwd **/
func (this *DataStoreConfiguration) SetPwd(ref interface{}){
	this.NeedSave = true
	this.M_pwd = ref.(string)
}

/** Remove reference Pwd **/

/** Port **/
func (this *DataStoreConfiguration) GetPort() int{
	return this.M_port
}

/** Init reference Port **/
func (this *DataStoreConfiguration) SetPort(ref interface{}){
	this.NeedSave = true
	this.M_port = ref.(int)
}

/** Remove reference Port **/

/** Parent **/
func (this *DataStoreConfiguration) GetParentPtr() *Configurations{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *DataStoreConfiguration) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*Configurations)
		this.M_parentPtr = ref.(*Configurations).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *DataStoreConfiguration) RemoveParentPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Configurations)
	if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
		this.m_parentPtr = nil
		this.M_parentPtr = ""
	}
}
