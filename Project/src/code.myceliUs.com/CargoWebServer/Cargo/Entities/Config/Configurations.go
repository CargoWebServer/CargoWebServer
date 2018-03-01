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
	M_serverConfig string
	M_oauth2Configuration string
	M_serviceConfigs []string
	M_dataStoreConfigs []string
	M_smtpConfigs []string
	M_ldapConfigs []string
	M_applicationConfigs []string
	M_scheduledTasks []string

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
func (this *Configurations) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Configurations) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

func (this *Configurations) GetId()string{
	return this.M_id
}

func (this *Configurations) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *Configurations) GetName()string{
	return this.M_name
}

func (this *Configurations) SetName(val string){
	this.NeedSave = this.M_name== val
	this.M_name= val
}


func (this *Configurations) GetVersion()string{
	return this.M_version
}

func (this *Configurations) SetVersion(val string){
	this.NeedSave = this.M_version== val
	this.M_version= val
}


func (this *Configurations) GetServerConfig()*ServerConfiguration{
	entity, err := this.getEntityByUuid(this.M_serverConfig)
	if err == nil {
		return entity.(*ServerConfiguration)
	}
	return nil
}

func (this *Configurations) SetServerConfig(val *ServerConfiguration){
	this.NeedSave = this.M_serverConfig != val.GetUuid()
	this.M_serverConfig= val.GetUuid()
}

func (this *Configurations) ResetServerConfig(){
	this.M_serverConfig= ""
}


func (this *Configurations) GetOauth2Configuration()*OAuth2Configuration{
	entity, err := this.getEntityByUuid(this.M_oauth2Configuration)
	if err == nil {
		return entity.(*OAuth2Configuration)
	}
	return nil
}

func (this *Configurations) SetOauth2Configuration(val *OAuth2Configuration){
	this.NeedSave = this.M_oauth2Configuration != val.GetUuid()
	this.M_oauth2Configuration= val.GetUuid()
}

func (this *Configurations) ResetOauth2Configuration(){
	this.M_oauth2Configuration= ""
}


func (this *Configurations) GetServiceConfigs()[]*ServiceConfiguration{
	serviceConfigs := make([]*ServiceConfiguration, 0)
	for i := 0; i < len(this.M_serviceConfigs); i++ {
		entity, err := this.getEntityByUuid(this.M_serviceConfigs[i])
		if err == nil {
			serviceConfigs = append(serviceConfigs, entity.(*ServiceConfiguration))
		}
	}
	return serviceConfigs
}

func (this *Configurations) SetServiceConfigs(val []*ServiceConfiguration){
	this.M_serviceConfigs= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_serviceConfigs=append(this.M_serviceConfigs, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Configurations) AppendServiceConfigs(val *ServiceConfiguration){
	for i:=0; i < len(this.M_serviceConfigs); i++{
		if this.M_serviceConfigs[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_serviceConfigs = append(this.M_serviceConfigs, val.GetUuid())
}

func (this *Configurations) RemoveServiceConfigs(val *ServiceConfiguration){
	serviceConfigs := make([]string,0)
	for i:=0; i < len(this.M_serviceConfigs); i++{
		if this.M_serviceConfigs[i] != val.GetUuid() {
			serviceConfigs = append(serviceConfigs, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_serviceConfigs = serviceConfigs
}


func (this *Configurations) GetDataStoreConfigs()[]*DataStoreConfiguration{
	dataStoreConfigs := make([]*DataStoreConfiguration, 0)
	for i := 0; i < len(this.M_dataStoreConfigs); i++ {
		entity, err := this.getEntityByUuid(this.M_dataStoreConfigs[i])
		if err == nil {
			dataStoreConfigs = append(dataStoreConfigs, entity.(*DataStoreConfiguration))
		}
	}
	return dataStoreConfigs
}

func (this *Configurations) SetDataStoreConfigs(val []*DataStoreConfiguration){
	this.M_dataStoreConfigs= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_dataStoreConfigs=append(this.M_dataStoreConfigs, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Configurations) AppendDataStoreConfigs(val *DataStoreConfiguration){
	for i:=0; i < len(this.M_dataStoreConfigs); i++{
		if this.M_dataStoreConfigs[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_dataStoreConfigs = append(this.M_dataStoreConfigs, val.GetUuid())
}

func (this *Configurations) RemoveDataStoreConfigs(val *DataStoreConfiguration){
	dataStoreConfigs := make([]string,0)
	for i:=0; i < len(this.M_dataStoreConfigs); i++{
		if this.M_dataStoreConfigs[i] != val.GetUuid() {
			dataStoreConfigs = append(dataStoreConfigs, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_dataStoreConfigs = dataStoreConfigs
}


func (this *Configurations) GetSmtpConfigs()[]*SmtpConfiguration{
	smtpConfigs := make([]*SmtpConfiguration, 0)
	for i := 0; i < len(this.M_smtpConfigs); i++ {
		entity, err := this.getEntityByUuid(this.M_smtpConfigs[i])
		if err == nil {
			smtpConfigs = append(smtpConfigs, entity.(*SmtpConfiguration))
		}
	}
	return smtpConfigs
}

func (this *Configurations) SetSmtpConfigs(val []*SmtpConfiguration){
	this.M_smtpConfigs= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_smtpConfigs=append(this.M_smtpConfigs, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Configurations) AppendSmtpConfigs(val *SmtpConfiguration){
	for i:=0; i < len(this.M_smtpConfigs); i++{
		if this.M_smtpConfigs[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_smtpConfigs = append(this.M_smtpConfigs, val.GetUuid())
}

func (this *Configurations) RemoveSmtpConfigs(val *SmtpConfiguration){
	smtpConfigs := make([]string,0)
	for i:=0; i < len(this.M_smtpConfigs); i++{
		if this.M_smtpConfigs[i] != val.GetUuid() {
			smtpConfigs = append(smtpConfigs, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_smtpConfigs = smtpConfigs
}


func (this *Configurations) GetLdapConfigs()[]*LdapConfiguration{
	ldapConfigs := make([]*LdapConfiguration, 0)
	for i := 0; i < len(this.M_ldapConfigs); i++ {
		entity, err := this.getEntityByUuid(this.M_ldapConfigs[i])
		if err == nil {
			ldapConfigs = append(ldapConfigs, entity.(*LdapConfiguration))
		}
	}
	return ldapConfigs
}

func (this *Configurations) SetLdapConfigs(val []*LdapConfiguration){
	this.M_ldapConfigs= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_ldapConfigs=append(this.M_ldapConfigs, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Configurations) AppendLdapConfigs(val *LdapConfiguration){
	for i:=0; i < len(this.M_ldapConfigs); i++{
		if this.M_ldapConfigs[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_ldapConfigs = append(this.M_ldapConfigs, val.GetUuid())
}

func (this *Configurations) RemoveLdapConfigs(val *LdapConfiguration){
	ldapConfigs := make([]string,0)
	for i:=0; i < len(this.M_ldapConfigs); i++{
		if this.M_ldapConfigs[i] != val.GetUuid() {
			ldapConfigs = append(ldapConfigs, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_ldapConfigs = ldapConfigs
}


func (this *Configurations) GetApplicationConfigs()[]*ApplicationConfiguration{
	applicationConfigs := make([]*ApplicationConfiguration, 0)
	for i := 0; i < len(this.M_applicationConfigs); i++ {
		entity, err := this.getEntityByUuid(this.M_applicationConfigs[i])
		if err == nil {
			applicationConfigs = append(applicationConfigs, entity.(*ApplicationConfiguration))
		}
	}
	return applicationConfigs
}

func (this *Configurations) SetApplicationConfigs(val []*ApplicationConfiguration){
	this.M_applicationConfigs= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_applicationConfigs=append(this.M_applicationConfigs, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Configurations) AppendApplicationConfigs(val *ApplicationConfiguration){
	for i:=0; i < len(this.M_applicationConfigs); i++{
		if this.M_applicationConfigs[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_applicationConfigs = append(this.M_applicationConfigs, val.GetUuid())
}

func (this *Configurations) RemoveApplicationConfigs(val *ApplicationConfiguration){
	applicationConfigs := make([]string,0)
	for i:=0; i < len(this.M_applicationConfigs); i++{
		if this.M_applicationConfigs[i] != val.GetUuid() {
			applicationConfigs = append(applicationConfigs, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_applicationConfigs = applicationConfigs
}


func (this *Configurations) GetScheduledTasks()[]*ScheduledTask{
	scheduledTasks := make([]*ScheduledTask, 0)
	for i := 0; i < len(this.M_scheduledTasks); i++ {
		entity, err := this.getEntityByUuid(this.M_scheduledTasks[i])
		if err == nil {
			scheduledTasks = append(scheduledTasks, entity.(*ScheduledTask))
		}
	}
	return scheduledTasks
}

func (this *Configurations) SetScheduledTasks(val []*ScheduledTask){
	this.M_scheduledTasks= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_scheduledTasks=append(this.M_scheduledTasks, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Configurations) AppendScheduledTasks(val *ScheduledTask){
	for i:=0; i < len(this.M_scheduledTasks); i++{
		if this.M_scheduledTasks[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_scheduledTasks = append(this.M_scheduledTasks, val.GetUuid())
}

func (this *Configurations) RemoveScheduledTasks(val *ScheduledTask){
	scheduledTasks := make([]string,0)
	for i:=0; i < len(this.M_scheduledTasks); i++{
		if this.M_scheduledTasks[i] != val.GetUuid() {
			scheduledTasks = append(scheduledTasks, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_scheduledTasks = scheduledTasks
}

