// +build Config

package Config

import (
	"encoding/xml"
	"strings"

	"code.myceliUs.com/Utility"
)

type Configurations struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** Keep reference to entity that made use of thit entity **/
	Referenced []string
	/** Get entity by uuid function **/
	getEntityByUuid func(string) (interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of Configurations **/
	M_id                  string
	M_name                string
	M_version             string
	M_serverConfig        string
	M_oauth2Configuration string
	M_serviceConfigs      []string
	M_dataStoreConfigs    []string
	M_smtpConfigs         []string
	M_ldapConfigs         []string
	M_applicationConfigs  []string
	M_scheduledTasks      []string
}

/** Xml parser for Configurations **/
type XsdConfigurations struct {
	XMLName               xml.Name                       `xml:"configurations"`
	M_serverConfig        XsdServerConfiguration         `xml:"serverConfig"`
	M_applicationConfigs  []*XsdApplicationConfiguration `xml:"applicationConfigs,omitempty"`
	M_smtpConfigs         []*XsdSmtpConfiguration        `xml:"smtpConfigs,omitempty"`
	M_ldapConfigs         []*XsdLdapConfiguration        `xml:"ldapConfigs,omitempty"`
	M_dataStoreConfigs    []*XsdDataStoreConfiguration   `xml:"dataStoreConfigs,omitempty"`
	M_serviceConfigs      []*XsdServiceConfiguration     `xml:"serviceConfigs,omitempty"`
	M_oauth2Configuration *XsdOAuth2Configuration        `xml:"oauth2Configuration,omitempty"`
	M_scheduledTasks      []*XsdScheduledTask            `xml:"scheduledTasks,omitempty"`
	M_id                  string                         `xml:"id,attr"`
	M_name                string                         `xml:"name,attr"`
	M_version             string                         `xml:"version,attr"`
}

/***************** Entity **************************/

/** UUID **/
func (this *Configurations) GetUuid() string {
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Configurations) SetUuid(uuid string) {
	this.UUID = uuid
}

func (this *Configurations) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *Configurations) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *Configurations) RemoveReferenced(uuid string, field string) {
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

func (this *Configurations) SetFieldValue(field string, value interface{}) error {
	return Utility.SetProperty(this, field, value)
}

func (this *Configurations) GetFieldValue(field string) interface{} {
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *Configurations) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Configurations) GetTypeName() string {
	this.TYPENAME = "Config.Configurations"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Configurations) GetParentUuid() string {
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Configurations) SetParentUuid(parentUuid string) {
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Configurations) GetParentLnk() string {
	return this.ParentLnk
}
func (this *Configurations) SetParentLnk(parentLnk string) {
	this.ParentLnk = parentLnk
}

func (this *Configurations) GetParent() interface{} {
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Configurations) GetChilds() []interface{} {
	var childs []interface{}
	var child interface{}
	var err error
	child, err = this.getEntityByUuid(this.M_serverConfig)
	if err == nil {
		childs = append(childs, child)
	}
	child, err = this.getEntityByUuid(this.M_oauth2Configuration)
	if err == nil {
		childs = append(childs, child)
	}
	for i := 0; i < len(this.M_serviceConfigs); i++ {
		child, err = this.getEntityByUuid(this.M_serviceConfigs[i])
		if err == nil {
			childs = append(childs, child)
		}
	}
	for i := 0; i < len(this.M_dataStoreConfigs); i++ {
		child, err = this.getEntityByUuid(this.M_dataStoreConfigs[i])
		if err == nil {
			childs = append(childs, child)
		}
	}
	for i := 0; i < len(this.M_smtpConfigs); i++ {
		child, err = this.getEntityByUuid(this.M_smtpConfigs[i])
		if err == nil {
			childs = append(childs, child)
		}
	}
	for i := 0; i < len(this.M_ldapConfigs); i++ {
		child, err = this.getEntityByUuid(this.M_ldapConfigs[i])
		if err == nil {
			childs = append(childs, child)
		}
	}
	for i := 0; i < len(this.M_applicationConfigs); i++ {
		child, err = this.getEntityByUuid(this.M_applicationConfigs[i])
		if err == nil {
			childs = append(childs, child)
		}
	}
	for i := 0; i < len(this.M_scheduledTasks); i++ {
		child, err = this.getEntityByUuid(this.M_scheduledTasks[i])
		if err == nil {
			childs = append(childs, child)
		}
	}
	return childs
}

/** Return the list of all childs uuid **/
func (this *Configurations) GetChildsUuid() []string {
	var childs []string
	childs = append(childs, this.M_serverConfig)
	childs = append(childs, this.M_oauth2Configuration)
	childs = append(childs, this.M_serviceConfigs...)
	childs = append(childs, this.M_dataStoreConfigs...)
	childs = append(childs, this.M_smtpConfigs...)
	childs = append(childs, this.M_ldapConfigs...)
	childs = append(childs, this.M_applicationConfigs...)
	childs = append(childs, this.M_scheduledTasks...)
	return childs
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Configurations) SetEntityGetter(fct func(uuid string) (interface{}, error)) {
	this.getEntityByUuid = fct
}

/** Use it the set the entity on the cache. **/
func (this *Configurations) SetEntitySetter(fct func(entity interface{})) {
	this.setEntity = fct
}

/** Set the uuid generator function **/
func (this *Configurations) SetUuidGenerator(fct func(entity interface{}) string) {
	this.generateUuid = fct
}

func (this *Configurations) GetId() string {
	return this.M_id
}

func (this *Configurations) SetId(val string) {
	this.M_id = val
}

func (this *Configurations) GetName() string {
	return this.M_name
}

func (this *Configurations) SetName(val string) {
	this.M_name = val
}

func (this *Configurations) GetVersion() string {
	return this.M_version
}

func (this *Configurations) SetVersion(val string) {
	this.M_version = val
}

func (this *Configurations) GetServerConfig() *ServerConfiguration {
	entity, err := this.getEntityByUuid(this.M_serverConfig)
	if err == nil {
		return entity.(*ServerConfiguration)
	}
	return nil
}

func (this *Configurations) SetServerConfig(val *ServerConfiguration) {
	this.M_serverConfig = val.GetUuid()
	if len(val.GetParentUuid()) > 0 && len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Reset" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_serverConfig")
	this.setEntity(val)
	this.setEntity(this)
}

func (this *Configurations) ResetServerConfig() {
	val := this.GetServerConfig()
	val.RemoveReferenced(this.UUID, "M_serverConfig")
	this.M_serverConfig = ""
}

func (this *Configurations) GetOauth2Configuration() *OAuth2Configuration {
	entity, err := this.getEntityByUuid(this.M_oauth2Configuration)
	if err == nil {
		return entity.(*OAuth2Configuration)
	}
	return nil
}

func (this *Configurations) SetOauth2Configuration(val *OAuth2Configuration) {
	this.M_oauth2Configuration = val.GetUuid()
	if len(val.GetParentUuid()) > 0 && len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Reset" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_oauth2Configuration")
	this.setEntity(val)
	this.setEntity(this)
}

func (this *Configurations) ResetOauth2Configuration() {
	val := this.GetOauth2Configuration()
	val.RemoveReferenced(this.UUID, "M_oauth2Configuration")
	this.M_oauth2Configuration = ""
}

func (this *Configurations) GetServiceConfigs() []*ServiceConfiguration {
	values := make([]*ServiceConfiguration, 0)
	for i := 0; i < len(this.M_serviceConfigs); i++ {
		entity, err := this.getEntityByUuid(this.M_serviceConfigs[i])
		if err == nil {
			values = append(values, entity.(*ServiceConfiguration))
		}
	}
	return values
}

func (this *Configurations) SetServiceConfigs(val []*ServiceConfiguration) {
	this.M_serviceConfigs = make([]string, 0)
	for i := 0; i < len(val); i++ {
		this.M_serviceConfigs = append(this.M_serviceConfigs, val[i].GetUuid())
		if len(val[i].GetParentUuid()) > 0 && len(val[i].GetParentLnk()) > 0 && this.GetUuid() != val[i].GetParentUuid() {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_serviceConfigs")
		this.setEntity(val[i])
	}
	this.setEntity(this)
}

func (this *Configurations) AppendServiceConfigs(val *ServiceConfiguration) {
	for i := 0; i < len(this.M_serviceConfigs); i++ {
		if this.M_serviceConfigs[i] == val.GetUuid() {
			return
		}
	}
	this.M_serviceConfigs = append(this.M_serviceConfigs, val.GetUuid())
	if len(val.GetParentUuid()) > 0 && len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_serviceConfigs")
	this.setEntity(val)
	this.setEntity(this)
}

func (this *Configurations) RemoveServiceConfigs(val *ServiceConfiguration) {
	values := make([]string, 0)
	for i := 0; i < len(this.M_serviceConfigs); i++ {
		if this.M_serviceConfigs[i] != val.GetUuid() {
			values = append(values, this.M_serviceConfigs[i])
		}
	}
	this.M_serviceConfigs = values
	this.setEntity(this)
}

func (this *Configurations) GetDataStoreConfigs() []*DataStoreConfiguration {
	values := make([]*DataStoreConfiguration, 0)
	for i := 0; i < len(this.M_dataStoreConfigs); i++ {
		entity, err := this.getEntityByUuid(this.M_dataStoreConfigs[i])
		if err == nil {
			values = append(values, entity.(*DataStoreConfiguration))
		}
	}
	return values
}

func (this *Configurations) SetDataStoreConfigs(val []*DataStoreConfiguration) {
	this.M_dataStoreConfigs = make([]string, 0)
	for i := 0; i < len(val); i++ {
		this.M_dataStoreConfigs = append(this.M_dataStoreConfigs, val[i].GetUuid())
		if len(val[i].GetParentUuid()) > 0 && len(val[i].GetParentLnk()) > 0 && this.GetUuid() != val[i].GetParentUuid() {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_dataStoreConfigs")
		this.setEntity(val[i])
	}
	this.setEntity(this)
}

func (this *Configurations) AppendDataStoreConfigs(val *DataStoreConfiguration) {
	for i := 0; i < len(this.M_dataStoreConfigs); i++ {
		if this.M_dataStoreConfigs[i] == val.GetUuid() {
			return
		}
	}
	this.M_dataStoreConfigs = append(this.M_dataStoreConfigs, val.GetUuid())
	if len(val.GetParentUuid()) > 0 && len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_dataStoreConfigs")
	this.setEntity(val)
	this.setEntity(this)
}

func (this *Configurations) RemoveDataStoreConfigs(val *DataStoreConfiguration) {
	values := make([]string, 0)
	for i := 0; i < len(this.M_dataStoreConfigs); i++ {
		if this.M_dataStoreConfigs[i] != val.GetUuid() {
			values = append(values, this.M_dataStoreConfigs[i])
		}
	}
	this.M_dataStoreConfigs = values
	this.setEntity(this)
}

func (this *Configurations) GetSmtpConfigs() []*SmtpConfiguration {
	values := make([]*SmtpConfiguration, 0)
	for i := 0; i < len(this.M_smtpConfigs); i++ {
		entity, err := this.getEntityByUuid(this.M_smtpConfigs[i])
		if err == nil {
			values = append(values, entity.(*SmtpConfiguration))
		}
	}
	return values
}

func (this *Configurations) SetSmtpConfigs(val []*SmtpConfiguration) {
	this.M_smtpConfigs = make([]string, 0)
	for i := 0; i < len(val); i++ {
		this.M_smtpConfigs = append(this.M_smtpConfigs, val[i].GetUuid())
		if len(val[i].GetParentUuid()) > 0 && len(val[i].GetParentLnk()) > 0 && this.GetUuid() != val[i].GetParentUuid() {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_smtpConfigs")
		this.setEntity(val[i])
	}
	this.setEntity(this)
}

func (this *Configurations) AppendSmtpConfigs(val *SmtpConfiguration) {
	for i := 0; i < len(this.M_smtpConfigs); i++ {
		if this.M_smtpConfigs[i] == val.GetUuid() {
			return
		}
	}
	this.M_smtpConfigs = append(this.M_smtpConfigs, val.GetUuid())
	if len(val.GetParentUuid()) > 0 && len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_smtpConfigs")
	this.setEntity(val)
	this.setEntity(this)
}

func (this *Configurations) RemoveSmtpConfigs(val *SmtpConfiguration) {
	values := make([]string, 0)
	for i := 0; i < len(this.M_smtpConfigs); i++ {
		if this.M_smtpConfigs[i] != val.GetUuid() {
			values = append(values, this.M_smtpConfigs[i])
		}
	}
	this.M_smtpConfigs = values
	this.setEntity(this)
}

func (this *Configurations) GetLdapConfigs() []*LdapConfiguration {
	values := make([]*LdapConfiguration, 0)
	for i := 0; i < len(this.M_ldapConfigs); i++ {
		entity, err := this.getEntityByUuid(this.M_ldapConfigs[i])
		if err == nil {
			values = append(values, entity.(*LdapConfiguration))
		}
	}
	return values
}

func (this *Configurations) SetLdapConfigs(val []*LdapConfiguration) {
	this.M_ldapConfigs = make([]string, 0)
	for i := 0; i < len(val); i++ {
		this.M_ldapConfigs = append(this.M_ldapConfigs, val[i].GetUuid())
		if len(val[i].GetParentUuid()) > 0 && len(val[i].GetParentLnk()) > 0 && this.GetUuid() != val[i].GetParentUuid() {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_ldapConfigs")
		this.setEntity(val[i])
	}
	this.setEntity(this)
}

func (this *Configurations) AppendLdapConfigs(val *LdapConfiguration) {
	for i := 0; i < len(this.M_ldapConfigs); i++ {
		if this.M_ldapConfigs[i] == val.GetUuid() {
			return
		}
	}
	this.M_ldapConfigs = append(this.M_ldapConfigs, val.GetUuid())
	if len(val.GetParentUuid()) > 0 && len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_ldapConfigs")
	this.setEntity(val)
	this.setEntity(this)
}

func (this *Configurations) RemoveLdapConfigs(val *LdapConfiguration) {
	values := make([]string, 0)
	for i := 0; i < len(this.M_ldapConfigs); i++ {
		if this.M_ldapConfigs[i] != val.GetUuid() {
			values = append(values, this.M_ldapConfigs[i])
		}
	}
	this.M_ldapConfigs = values
	this.setEntity(this)
}

func (this *Configurations) GetApplicationConfigs() []*ApplicationConfiguration {
	values := make([]*ApplicationConfiguration, 0)
	for i := 0; i < len(this.M_applicationConfigs); i++ {
		entity, err := this.getEntityByUuid(this.M_applicationConfigs[i])
		if err == nil {
			values = append(values, entity.(*ApplicationConfiguration))
		}
	}
	return values
}

func (this *Configurations) SetApplicationConfigs(val []*ApplicationConfiguration) {
	this.M_applicationConfigs = make([]string, 0)
	for i := 0; i < len(val); i++ {
		this.M_applicationConfigs = append(this.M_applicationConfigs, val[i].GetUuid())
		if len(val[i].GetParentUuid()) > 0 && len(val[i].GetParentLnk()) > 0 && this.GetUuid() != val[i].GetParentUuid() {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_applicationConfigs")
		this.setEntity(val[i])
	}
	this.setEntity(this)
}

func (this *Configurations) AppendApplicationConfigs(val *ApplicationConfiguration) {
	for i := 0; i < len(this.M_applicationConfigs); i++ {
		if this.M_applicationConfigs[i] == val.GetUuid() {
			return
		}
	}
	this.M_applicationConfigs = append(this.M_applicationConfigs, val.GetUuid())
	if len(val.GetParentUuid()) > 0 && len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_applicationConfigs")
	this.setEntity(val)
	this.setEntity(this)
}

func (this *Configurations) RemoveApplicationConfigs(val *ApplicationConfiguration) {
	values := make([]string, 0)
	for i := 0; i < len(this.M_applicationConfigs); i++ {
		if this.M_applicationConfigs[i] != val.GetUuid() {
			values = append(values, this.M_applicationConfigs[i])
		}
	}
	this.M_applicationConfigs = values
	this.setEntity(this)
}

func (this *Configurations) GetScheduledTasks() []*ScheduledTask {
	values := make([]*ScheduledTask, 0)
	for i := 0; i < len(this.M_scheduledTasks); i++ {
		entity, err := this.getEntityByUuid(this.M_scheduledTasks[i])
		if err == nil {
			values = append(values, entity.(*ScheduledTask))
		}
	}
	return values
}

func (this *Configurations) SetScheduledTasks(val []*ScheduledTask) {
	this.M_scheduledTasks = make([]string, 0)
	for i := 0; i < len(val); i++ {
		this.M_scheduledTasks = append(this.M_scheduledTasks, val[i].GetUuid())
		if len(val[i].GetParentUuid()) > 0 && len(val[i].GetParentLnk()) > 0 && this.GetUuid() != val[i].GetParentUuid() {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_scheduledTasks")
		this.setEntity(val[i])
	}
	this.setEntity(this)
}

func (this *Configurations) AppendScheduledTasks(val *ScheduledTask) {
	for i := 0; i < len(this.M_scheduledTasks); i++ {
		if this.M_scheduledTasks[i] == val.GetUuid() {
			return
		}
	}
	this.M_scheduledTasks = append(this.M_scheduledTasks, val.GetUuid())
	if len(val.GetParentUuid()) > 0 && len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_scheduledTasks")
	this.setEntity(val)
	this.setEntity(this)
}

func (this *Configurations) RemoveScheduledTasks(val *ScheduledTask) {
	values := make([]string, 0)
	for i := 0; i < len(this.M_scheduledTasks); i++ {
		if this.M_scheduledTasks[i] != val.GetUuid() {
			values = append(values, this.M_scheduledTasks[i])
		}
	}
	this.M_scheduledTasks = values
	this.setEntity(this)
}
