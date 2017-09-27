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

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Configuration **/
	M_id string

	/** members of DataStoreConfiguration **/
	M_dataStoreType DataStoreType
	M_dataStoreVendor DataStoreVendor
	M_textEncoding Encoding
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
	M_textEncoding	string	`xml:"textEncoding,attr"`

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
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** DataStoreType **/
func (this *DataStoreConfiguration) GetDataStoreType() DataStoreType{
	return this.M_dataStoreType
}

/** Init reference DataStoreType **/
func (this *DataStoreConfiguration) SetDataStoreType(ref interface{}){
	this.M_dataStoreType = ref.(DataStoreType)
}

/** Remove reference DataStoreType **/

/** DataStoreVendor **/
func (this *DataStoreConfiguration) GetDataStoreVendor() DataStoreVendor{
	return this.M_dataStoreVendor
}

/** Init reference DataStoreVendor **/
func (this *DataStoreConfiguration) SetDataStoreVendor(ref interface{}){
	this.M_dataStoreVendor = ref.(DataStoreVendor)
}

/** Remove reference DataStoreVendor **/

/** TextEncoding **/
func (this *DataStoreConfiguration) GetTextEncoding() Encoding{
	return this.M_textEncoding
}

/** Init reference TextEncoding **/
func (this *DataStoreConfiguration) SetTextEncoding(ref interface{}){
	this.M_textEncoding = ref.(Encoding)
}

/** Remove reference TextEncoding **/

/** HostName **/
func (this *DataStoreConfiguration) GetHostName() string{
	return this.M_hostName
}

/** Init reference HostName **/
func (this *DataStoreConfiguration) SetHostName(ref interface{}){
	this.M_hostName = ref.(string)
}

/** Remove reference HostName **/

/** Ipv4 **/
func (this *DataStoreConfiguration) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *DataStoreConfiguration) SetIpv4(ref interface{}){
	this.M_ipv4 = ref.(string)
}

/** Remove reference Ipv4 **/

/** User **/
func (this *DataStoreConfiguration) GetUser() string{
	return this.M_user
}

/** Init reference User **/
func (this *DataStoreConfiguration) SetUser(ref interface{}){
	this.M_user = ref.(string)
}

/** Remove reference User **/

/** Pwd **/
func (this *DataStoreConfiguration) GetPwd() string{
	return this.M_pwd
}

/** Init reference Pwd **/
func (this *DataStoreConfiguration) SetPwd(ref interface{}){
	this.M_pwd = ref.(string)
}

/** Remove reference Pwd **/

/** Port **/
func (this *DataStoreConfiguration) GetPort() int{
	return this.M_port
}

/** Init reference Port **/
func (this *DataStoreConfiguration) SetPort(ref interface{}){
	this.M_port = ref.(int)
}

/** Remove reference Port **/

/** Parent **/
func (this *DataStoreConfiguration) GetParentPtr() *Configurations{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *DataStoreConfiguration) SetParentPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.M_parentPtr = ref.(*Configurations).GetUUID()
		this.m_parentPtr = ref.(*Configurations)
	}
}

/** Remove reference Parent **/
func (this *DataStoreConfiguration) RemoveParentPtr(ref interface{}){
	toDelete := ref.(*Configurations)
	if this.m_parentPtr!= nil {
		if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
		}
	}
}
