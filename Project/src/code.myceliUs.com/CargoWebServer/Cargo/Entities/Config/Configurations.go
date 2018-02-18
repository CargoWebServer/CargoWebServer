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
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

	/** members of Configurations **/
	M_id string
	M_name string
	M_version string
	M_serverConfig *ServerConfiguration
	M_oauth2Configuration *OAuth2Configuration
	M_serviceConfigs []*ServiceConfiguration
	M_dataStoreConfigs []*DataStoreConfiguration
	M_smtpConfigs []*SmtpConfiguration
	M_ldapConfigs []*LdapConfiguration
	M_applicationConfigs []*ApplicationConfiguration
	M_scheduledTasks []*ScheduledTask

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
	M_scheduledTasks	[]*XsdScheduledTask	`xml:"scheduledTasks,omitempty"`
	M_id	string	`xml:"id,attr"`
	M_name	string	`xml:"name,attr"`
	M_version	string	`xml:"version,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Configurations) GetUuid() string{
	return this.UUID
}
func (this *Configurations) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Configurations) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Configurations) GetTypeName() string{
	this.TYPENAME = "Config.Configurations"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Configurations) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Configurations) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Configurations) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Configurations) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *Configurations) IsNeedSave() bool{
	return this.NeedSave
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Configurations) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *Configurations) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Configurations) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Name **/
func (this *Configurations) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Configurations) SetName(ref interface{}){
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Name **/

/** Version **/
func (this *Configurations) GetVersion() string{
	return this.M_version
}

/** Init reference Version **/
func (this *Configurations) SetVersion(ref interface{}){
	if this.M_version != ref.(string) {
		this.M_version = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Version **/

/** ServerConfig **/
func (this *Configurations) GetServerConfig() *ServerConfiguration{
	return this.M_serverConfig
}

/** Init reference ServerConfig **/
func (this *Configurations) SetServerConfig(ref interface{}){
	if this.M_serverConfig != ref.(*ServerConfiguration) {
		this.M_serverConfig = ref.(*ServerConfiguration)
		this.NeedSave = true
	}
}

/** Remove reference ServerConfig **/
func (this *Configurations) RemoveServerConfig(ref interface{}){
	toDelete := ref.(Configuration)
	if toDelete.GetUuid() == this.M_serverConfig.GetUuid() {
		this.M_serverConfig = nil
		this.NeedSave = true
	}
}

/** Oauth2Configuration **/
func (this *Configurations) GetOauth2Configuration() *OAuth2Configuration{
	return this.M_oauth2Configuration
}

/** Init reference Oauth2Configuration **/
func (this *Configurations) SetOauth2Configuration(ref interface{}){
	if this.M_oauth2Configuration != ref.(*OAuth2Configuration) {
		this.M_oauth2Configuration = ref.(*OAuth2Configuration)
		this.NeedSave = true
	}
}

/** Remove reference Oauth2Configuration **/
func (this *Configurations) RemoveOauth2Configuration(ref interface{}){
	toDelete := ref.(Configuration)
	if toDelete.GetUuid() == this.M_oauth2Configuration.GetUuid() {
		this.M_oauth2Configuration = nil
		this.NeedSave = true
	}
}

/** ServiceConfigs **/
func (this *Configurations) GetServiceConfigs() []*ServiceConfiguration{
	return this.M_serviceConfigs
}

/** Init reference ServiceConfigs **/
func (this *Configurations) SetServiceConfigs(ref interface{}){
	isExist := false
	var serviceConfigss []*ServiceConfiguration
	for i:=0; i<len(this.M_serviceConfigs); i++ {
		if this.M_serviceConfigs[i].GetUuid() != ref.(Configuration).GetUuid() {
			serviceConfigss = append(serviceConfigss, this.M_serviceConfigs[i])
		} else {
			isExist = true
			serviceConfigss = append(serviceConfigss, ref.(*ServiceConfiguration))
		}
	}
	if !isExist {
		serviceConfigss = append(serviceConfigss, ref.(*ServiceConfiguration))
		this.NeedSave = true
		this.M_serviceConfigs = serviceConfigss
	}
}

/** Remove reference ServiceConfigs **/
func (this *Configurations) RemoveServiceConfigs(ref interface{}){
	toDelete := ref.(Configuration)
	serviceConfigs_ := make([]*ServiceConfiguration, 0)
	for i := 0; i < len(this.M_serviceConfigs); i++ {
		if toDelete.GetUuid() != this.M_serviceConfigs[i].GetUuid() {
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
	isExist := false
	var dataStoreConfigss []*DataStoreConfiguration
	for i:=0; i<len(this.M_dataStoreConfigs); i++ {
		if this.M_dataStoreConfigs[i].GetUuid() != ref.(Configuration).GetUuid() {
			dataStoreConfigss = append(dataStoreConfigss, this.M_dataStoreConfigs[i])
		} else {
			isExist = true
			dataStoreConfigss = append(dataStoreConfigss, ref.(*DataStoreConfiguration))
		}
	}
	if !isExist {
		dataStoreConfigss = append(dataStoreConfigss, ref.(*DataStoreConfiguration))
		this.NeedSave = true
		this.M_dataStoreConfigs = dataStoreConfigss
	}
}

/** Remove reference DataStoreConfigs **/
func (this *Configurations) RemoveDataStoreConfigs(ref interface{}){
	toDelete := ref.(Configuration)
	dataStoreConfigs_ := make([]*DataStoreConfiguration, 0)
	for i := 0; i < len(this.M_dataStoreConfigs); i++ {
		if toDelete.GetUuid() != this.M_dataStoreConfigs[i].GetUuid() {
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
	isExist := false
	var smtpConfigss []*SmtpConfiguration
	for i:=0; i<len(this.M_smtpConfigs); i++ {
		if this.M_smtpConfigs[i].GetUuid() != ref.(Configuration).GetUuid() {
			smtpConfigss = append(smtpConfigss, this.M_smtpConfigs[i])
		} else {
			isExist = true
			smtpConfigss = append(smtpConfigss, ref.(*SmtpConfiguration))
		}
	}
	if !isExist {
		smtpConfigss = append(smtpConfigss, ref.(*SmtpConfiguration))
		this.NeedSave = true
		this.M_smtpConfigs = smtpConfigss
	}
}

/** Remove reference SmtpConfigs **/
func (this *Configurations) RemoveSmtpConfigs(ref interface{}){
	toDelete := ref.(Configuration)
	smtpConfigs_ := make([]*SmtpConfiguration, 0)
	for i := 0; i < len(this.M_smtpConfigs); i++ {
		if toDelete.GetUuid() != this.M_smtpConfigs[i].GetUuid() {
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
	isExist := false
	var ldapConfigss []*LdapConfiguration
	for i:=0; i<len(this.M_ldapConfigs); i++ {
		if this.M_ldapConfigs[i].GetUuid() != ref.(Configuration).GetUuid() {
			ldapConfigss = append(ldapConfigss, this.M_ldapConfigs[i])
		} else {
			isExist = true
			ldapConfigss = append(ldapConfigss, ref.(*LdapConfiguration))
		}
	}
	if !isExist {
		ldapConfigss = append(ldapConfigss, ref.(*LdapConfiguration))
		this.NeedSave = true
		this.M_ldapConfigs = ldapConfigss
	}
}

/** Remove reference LdapConfigs **/
func (this *Configurations) RemoveLdapConfigs(ref interface{}){
	toDelete := ref.(Configuration)
	ldapConfigs_ := make([]*LdapConfiguration, 0)
	for i := 0; i < len(this.M_ldapConfigs); i++ {
		if toDelete.GetUuid() != this.M_ldapConfigs[i].GetUuid() {
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
	isExist := false
	var applicationConfigss []*ApplicationConfiguration
	for i:=0; i<len(this.M_applicationConfigs); i++ {
		if this.M_applicationConfigs[i].GetUuid() != ref.(Configuration).GetUuid() {
			applicationConfigss = append(applicationConfigss, this.M_applicationConfigs[i])
		} else {
			isExist = true
			applicationConfigss = append(applicationConfigss, ref.(*ApplicationConfiguration))
		}
	}
	if !isExist {
		applicationConfigss = append(applicationConfigss, ref.(*ApplicationConfiguration))
		this.NeedSave = true
		this.M_applicationConfigs = applicationConfigss
	}
}

/** Remove reference ApplicationConfigs **/
func (this *Configurations) RemoveApplicationConfigs(ref interface{}){
	toDelete := ref.(Configuration)
	applicationConfigs_ := make([]*ApplicationConfiguration, 0)
	for i := 0; i < len(this.M_applicationConfigs); i++ {
		if toDelete.GetUuid() != this.M_applicationConfigs[i].GetUuid() {
			applicationConfigs_ = append(applicationConfigs_, this.M_applicationConfigs[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_applicationConfigs = applicationConfigs_
}

/** ScheduledTasks **/
func (this *Configurations) GetScheduledTasks() []*ScheduledTask{
	return this.M_scheduledTasks
}

/** Init reference ScheduledTasks **/
func (this *Configurations) SetScheduledTasks(ref interface{}){
	isExist := false
	var scheduledTaskss []*ScheduledTask
	for i:=0; i<len(this.M_scheduledTasks); i++ {
		if this.M_scheduledTasks[i].GetUuid() != ref.(Configuration).GetUuid() {
			scheduledTaskss = append(scheduledTaskss, this.M_scheduledTasks[i])
		} else {
			isExist = true
			scheduledTaskss = append(scheduledTaskss, ref.(*ScheduledTask))
		}
	}
	if !isExist {
		scheduledTaskss = append(scheduledTaskss, ref.(*ScheduledTask))
		this.NeedSave = true
		this.M_scheduledTasks = scheduledTaskss
	}
}

/** Remove reference ScheduledTasks **/
func (this *Configurations) RemoveScheduledTasks(ref interface{}){
	toDelete := ref.(Configuration)
	scheduledTasks_ := make([]*ScheduledTask, 0)
	for i := 0; i < len(this.M_scheduledTasks); i++ {
		if toDelete.GetUuid() != this.M_scheduledTasks[i].GetUuid() {
			scheduledTasks_ = append(scheduledTasks_, this.M_scheduledTasks[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_scheduledTasks = scheduledTasks_
}
