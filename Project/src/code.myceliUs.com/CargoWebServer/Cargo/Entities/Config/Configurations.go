// +build Config

package Config

import(
"encoding/xml"
)

type Configurations struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Configurations **/
	M_id string
	M_name string
	M_version string
	M_filePath string
	M_serverConfig *ServerConfiguration
	M_oauth2Configuration *OAuth2Configuration
	M_serviceConfigs []*ServiceConfiguration
	M_dataStoreConfigs []*DataStoreConfiguration
	M_smtpConfigs []*SmtpConfiguration
	M_ldapConfigs []*LdapConfiguration
	M_applicationConfigs []*ApplicationConfiguration

}

/** Xml parser for Configurations **/
type XsdConfigurations struct {
	XMLName xml.Name	`xml:"configurations"`
	M_serverConfig	XsdServerConfiguration	`xml:"serverConfig"`
	M_applicationConfigs	[]*XsdApplicationConfiguration	`xml:"applicationConfigs,omitempty"`
	M_smtpConfigs	[]*XsdSmtpConfiguration	`xml:"smtpConfigs,omitempty"`
	M_ldapConfigs	[]*XsdLdapConfiguration	`xml:"ldapConfigs,omitempty"`
	M_dataStoreConfigs	[]*XsdDataStoreConfiguration	`xml:"dataStoreConfigs,omitempty"`
	M_serviceConfigs	[]*XsdServiceConfiguration	`xml:"serviceConfigs,omitempty"`
	M_oauth2Configuration	*XsdOAuth2Configuration	`xml:"oauth2Configuration,omitempty"`
	M_id	string	`xml:"id,attr"`
	M_name	string	`xml:"name,attr"`
	M_version	string	`xml:"version,attr"`

}
/** UUID **/
func (this *Configurations) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Configurations) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Configurations) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Name **/
func (this *Configurations) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Configurations) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Version **/
func (this *Configurations) GetVersion() string{
	return this.M_version
}

/** Init reference Version **/
func (this *Configurations) SetVersion(ref interface{}){
	this.NeedSave = true
	this.M_version = ref.(string)
}

/** Remove reference Version **/

/** FilePath **/
func (this *Configurations) GetFilePath() string{
	return this.M_filePath
}

/** Init reference FilePath **/
func (this *Configurations) SetFilePath(ref interface{}){
	this.NeedSave = true
	this.M_filePath = ref.(string)
}

/** Remove reference FilePath **/

/** ServerConfig **/
func (this *Configurations) GetServerConfig() *ServerConfiguration{
	return this.M_serverConfig
}

/** Init reference ServerConfig **/
func (this *Configurations) SetServerConfig(ref interface{}){
	this.NeedSave = true
	this.M_serverConfig = ref.(*ServerConfiguration)
}

/** Remove reference ServerConfig **/
func (this *Configurations) RemoveServerConfig(ref interface{}){
	toDelete := ref.(Configuration)
	if toDelete.GetUUID() == this.M_serverConfig.GetUUID() {
		this.M_serverConfig = nil
	}
}

/** Oauth2Configuration **/
func (this *Configurations) GetOauth2Configuration() *OAuth2Configuration{
	return this.M_oauth2Configuration
}

/** Init reference Oauth2Configuration **/
func (this *Configurations) SetOauth2Configuration(ref interface{}){
	this.NeedSave = true
	this.M_oauth2Configuration = ref.(*OAuth2Configuration)
}

/** Remove reference Oauth2Configuration **/
func (this *Configurations) RemoveOauth2Configuration(ref interface{}){
	toDelete := ref.(Configuration)
	if toDelete.GetUUID() == this.M_oauth2Configuration.GetUUID() {
		this.M_oauth2Configuration = nil
	}
}

/** ServiceConfigs **/
func (this *Configurations) GetServiceConfigs() []*ServiceConfiguration{
	return this.M_serviceConfigs
}

/** Init reference ServiceConfigs **/
func (this *Configurations) SetServiceConfigs(ref interface{}){
	this.NeedSave = true
	isExist := false
	var serviceConfigss []*ServiceConfiguration
	for i:=0; i<len(this.M_serviceConfigs); i++ {
		if this.M_serviceConfigs[i].GetUUID() != ref.(Configuration).GetUUID() {
			serviceConfigss = append(serviceConfigss, this.M_serviceConfigs[i])
		} else {
			isExist = true
			serviceConfigss = append(serviceConfigss, ref.(*ServiceConfiguration))
		}
	}
	if !isExist {
		serviceConfigss = append(serviceConfigss, ref.(*ServiceConfiguration))
	}
	this.M_serviceConfigs = serviceConfigss
}

/** Remove reference ServiceConfigs **/
func (this *Configurations) RemoveServiceConfigs(ref interface{}){
	toDelete := ref.(Configuration)
	serviceConfigs_ := make([]*ServiceConfiguration, 0)
	for i := 0; i < len(this.M_serviceConfigs); i++ {
		if toDelete.GetUUID() != this.M_serviceConfigs[i].GetUUID() {
			serviceConfigs_ = append(serviceConfigs_, this.M_serviceConfigs[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_serviceConfigs = serviceConfigs_
}

/** DataStoreConfigs **/
func (this *Configurations) GetDataStoreConfigs() []*DataStoreConfiguration{
	return this.M_dataStoreConfigs
}

/** Init reference DataStoreConfigs **/
func (this *Configurations) SetDataStoreConfigs(ref interface{}){
	this.NeedSave = true
	isExist := false
	var dataStoreConfigss []*DataStoreConfiguration
	for i:=0; i<len(this.M_dataStoreConfigs); i++ {
		if this.M_dataStoreConfigs[i].GetUUID() != ref.(Configuration).GetUUID() {
			dataStoreConfigss = append(dataStoreConfigss, this.M_dataStoreConfigs[i])
		} else {
			isExist = true
			dataStoreConfigss = append(dataStoreConfigss, ref.(*DataStoreConfiguration))
		}
	}
	if !isExist {
		dataStoreConfigss = append(dataStoreConfigss, ref.(*DataStoreConfiguration))
	}
	this.M_dataStoreConfigs = dataStoreConfigss
}

/** Remove reference DataStoreConfigs **/
func (this *Configurations) RemoveDataStoreConfigs(ref interface{}){
	toDelete := ref.(Configuration)
	dataStoreConfigs_ := make([]*DataStoreConfiguration, 0)
	for i := 0; i < len(this.M_dataStoreConfigs); i++ {
		if toDelete.GetUUID() != this.M_dataStoreConfigs[i].GetUUID() {
			dataStoreConfigs_ = append(dataStoreConfigs_, this.M_dataStoreConfigs[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_dataStoreConfigs = dataStoreConfigs_
}

/** SmtpConfigs **/
func (this *Configurations) GetSmtpConfigs() []*SmtpConfiguration{
	return this.M_smtpConfigs
}

/** Init reference SmtpConfigs **/
func (this *Configurations) SetSmtpConfigs(ref interface{}){
	this.NeedSave = true
	isExist := false
	var smtpConfigss []*SmtpConfiguration
	for i:=0; i<len(this.M_smtpConfigs); i++ {
		if this.M_smtpConfigs[i].GetUUID() != ref.(Configuration).GetUUID() {
			smtpConfigss = append(smtpConfigss, this.M_smtpConfigs[i])
		} else {
			isExist = true
			smtpConfigss = append(smtpConfigss, ref.(*SmtpConfiguration))
		}
	}
	if !isExist {
		smtpConfigss = append(smtpConfigss, ref.(*SmtpConfiguration))
	}
	this.M_smtpConfigs = smtpConfigss
}

/** Remove reference SmtpConfigs **/
func (this *Configurations) RemoveSmtpConfigs(ref interface{}){
	toDelete := ref.(Configuration)
	smtpConfigs_ := make([]*SmtpConfiguration, 0)
	for i := 0; i < len(this.M_smtpConfigs); i++ {
		if toDelete.GetUUID() != this.M_smtpConfigs[i].GetUUID() {
			smtpConfigs_ = append(smtpConfigs_, this.M_smtpConfigs[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_smtpConfigs = smtpConfigs_
}

/** LdapConfigs **/
func (this *Configurations) GetLdapConfigs() []*LdapConfiguration{
	return this.M_ldapConfigs
}

/** Init reference LdapConfigs **/
func (this *Configurations) SetLdapConfigs(ref interface{}){
	this.NeedSave = true
	isExist := false
	var ldapConfigss []*LdapConfiguration
	for i:=0; i<len(this.M_ldapConfigs); i++ {
		if this.M_ldapConfigs[i].GetUUID() != ref.(Configuration).GetUUID() {
			ldapConfigss = append(ldapConfigss, this.M_ldapConfigs[i])
		} else {
			isExist = true
			ldapConfigss = append(ldapConfigss, ref.(*LdapConfiguration))
		}
	}
	if !isExist {
		ldapConfigss = append(ldapConfigss, ref.(*LdapConfiguration))
	}
	this.M_ldapConfigs = ldapConfigss
}

/** Remove reference LdapConfigs **/
func (this *Configurations) RemoveLdapConfigs(ref interface{}){
	toDelete := ref.(Configuration)
	ldapConfigs_ := make([]*LdapConfiguration, 0)
	for i := 0; i < len(this.M_ldapConfigs); i++ {
		if toDelete.GetUUID() != this.M_ldapConfigs[i].GetUUID() {
			ldapConfigs_ = append(ldapConfigs_, this.M_ldapConfigs[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_ldapConfigs = ldapConfigs_
}

/** ApplicationConfigs **/
func (this *Configurations) GetApplicationConfigs() []*ApplicationConfiguration{
	return this.M_applicationConfigs
}

/** Init reference ApplicationConfigs **/
func (this *Configurations) SetApplicationConfigs(ref interface{}){
	this.NeedSave = true
	isExist := false
	var applicationConfigss []*ApplicationConfiguration
	for i:=0; i<len(this.M_applicationConfigs); i++ {
		if this.M_applicationConfigs[i].GetUUID() != ref.(Configuration).GetUUID() {
			applicationConfigss = append(applicationConfigss, this.M_applicationConfigs[i])
		} else {
			isExist = true
			applicationConfigss = append(applicationConfigss, ref.(*ApplicationConfiguration))
		}
	}
	if !isExist {
		applicationConfigss = append(applicationConfigss, ref.(*ApplicationConfiguration))
	}
	this.M_applicationConfigs = applicationConfigss
}

/** Remove reference ApplicationConfigs **/
func (this *Configurations) RemoveApplicationConfigs(ref interface{}){
	toDelete := ref.(Configuration)
	applicationConfigs_ := make([]*ApplicationConfiguration, 0)
	for i := 0; i < len(this.M_applicationConfigs); i++ {
		if toDelete.GetUUID() != this.M_applicationConfigs[i].GetUUID() {
			applicationConfigs_ = append(applicationConfigs_, this.M_applicationConfigs[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_applicationConfigs = applicationConfigs_
}
