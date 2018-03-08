// +build Config

package Config

import(
	"encoding/xml"
)

type ServerConfiguration struct{

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

	/** members of ServerConfiguration **/
	M_hostName string
	M_ipv4 string
	M_serverPort int
	M_serviceContainerPort int
	M_applicationsPath string
	M_dataPath string
	M_scriptsPath string
	M_definitionsPath string
	M_schemasPath string
	M_tmpPath string
	M_binPath string
	M_shards int
	M_lifeWindow int
	M_maxEntriesInWindow int
	M_maxEntrySize int
	M_hardMaxCacheSize int
	M_verbose bool


	/** Associations **/
	M_parentPtr string
}

/** Xml parser for ServerConfiguration **/
type XsdServerConfiguration struct {
	XMLName xml.Name	`xml:"serverConfiguration"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_ipv4	string	`xml:"ipv4,attr"`
	M_hostName	string	`xml:"hostName,attr"`
	M_serverPort	int	`xml:"serverPort,attr"`
	M_serviceContainerPort	int	`xml:"serviceContainerPort,attr"`
	M_applicationsPath	string	`xml:"applicationsPath,attr"`
	M_dataPath	string	`xml:"dataPath,attr"`
	M_scriptsPath	string	`xml:"scriptsPath,attr"`
	M_definitionsPath	string	`xml:"definitionsPath,attr"`
	M_schemasPath	string	`xml:"schemasPath,attr"`
	M_tmpPath	string	`xml:"tmpPath,attr"`
	M_binPath	string	`xml:"binPath,attr"`
	M_shards	int	`xml:"shards,attr"`
	M_lifeWindow	int	`xml:"lifeWindow,attr"`
	M_maxEntriesInWindow	int	`xml:"maxEntriesInWindow,attr"`
	M_maxEntrySize	int	`xml:"maxEntrySize,attr"`
	M_verbose	bool	`xml:"verbose,attr"`
	M_hardMaxCacheSize	int	`xml:"hardMaxCacheSize,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *ServerConfiguration) GetUuid() string{
	return this.UUID
}
func (this *ServerConfiguration) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *ServerConfiguration) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *ServerConfiguration) GetTypeName() string{
	this.TYPENAME = "Config.ServerConfiguration"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *ServerConfiguration) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *ServerConfiguration) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ServerConfiguration) GetParentLnk() string{
	return this.ParentLnk
}
func (this *ServerConfiguration) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ServerConfiguration) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Evaluate if an entity needs to be saved. **/
func (this *ServerConfiguration) IsNeedSave() bool{
	return this.NeedSave
}
func (this *ServerConfiguration) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *ServerConfiguration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *ServerConfiguration) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *ServerConfiguration) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *ServerConfiguration) GetId()string{
	return this.M_id
}

func (this *ServerConfiguration) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *ServerConfiguration) GetHostName()string{
	return this.M_hostName
}

func (this *ServerConfiguration) SetHostName(val string){
	this.NeedSave = this.M_hostName== val
	this.M_hostName= val
}


func (this *ServerConfiguration) GetIpv4()string{
	return this.M_ipv4
}

func (this *ServerConfiguration) SetIpv4(val string){
	this.NeedSave = this.M_ipv4== val
	this.M_ipv4= val
}


func (this *ServerConfiguration) GetServerPort()int{
	return this.M_serverPort
}

func (this *ServerConfiguration) SetServerPort(val int){
	this.NeedSave = this.M_serverPort== val
	this.M_serverPort= val
}


func (this *ServerConfiguration) GetServiceContainerPort()int{
	return this.M_serviceContainerPort
}

func (this *ServerConfiguration) SetServiceContainerPort(val int){
	this.NeedSave = this.M_serviceContainerPort== val
	this.M_serviceContainerPort= val
}


func (this *ServerConfiguration) GetApplicationsPath()string{
	return this.M_applicationsPath
}

func (this *ServerConfiguration) SetApplicationsPath(val string){
	this.NeedSave = this.M_applicationsPath== val
	this.M_applicationsPath= val
}


func (this *ServerConfiguration) GetDataPath()string{
	return this.M_dataPath
}

func (this *ServerConfiguration) SetDataPath(val string){
	this.NeedSave = this.M_dataPath== val
	this.M_dataPath= val
}


func (this *ServerConfiguration) GetScriptsPath()string{
	return this.M_scriptsPath
}

func (this *ServerConfiguration) SetScriptsPath(val string){
	this.NeedSave = this.M_scriptsPath== val
	this.M_scriptsPath= val
}


func (this *ServerConfiguration) GetDefinitionsPath()string{
	return this.M_definitionsPath
}

func (this *ServerConfiguration) SetDefinitionsPath(val string){
	this.NeedSave = this.M_definitionsPath== val
	this.M_definitionsPath= val
}


func (this *ServerConfiguration) GetSchemasPath()string{
	return this.M_schemasPath
}

func (this *ServerConfiguration) SetSchemasPath(val string){
	this.NeedSave = this.M_schemasPath== val
	this.M_schemasPath= val
}


func (this *ServerConfiguration) GetTmpPath()string{
	return this.M_tmpPath
}

func (this *ServerConfiguration) SetTmpPath(val string){
	this.NeedSave = this.M_tmpPath== val
	this.M_tmpPath= val
}


func (this *ServerConfiguration) GetBinPath()string{
	return this.M_binPath
}

func (this *ServerConfiguration) SetBinPath(val string){
	this.NeedSave = this.M_binPath== val
	this.M_binPath= val
}


func (this *ServerConfiguration) GetShards()int{
	return this.M_shards
}

func (this *ServerConfiguration) SetShards(val int){
	this.NeedSave = this.M_shards== val
	this.M_shards= val
}


func (this *ServerConfiguration) GetLifeWindow()int{
	return this.M_lifeWindow
}

func (this *ServerConfiguration) SetLifeWindow(val int){
	this.NeedSave = this.M_lifeWindow== val
	this.M_lifeWindow= val
}


func (this *ServerConfiguration) GetMaxEntriesInWindow()int{
	return this.M_maxEntriesInWindow
}

func (this *ServerConfiguration) SetMaxEntriesInWindow(val int){
	this.NeedSave = this.M_maxEntriesInWindow== val
	this.M_maxEntriesInWindow= val
}


func (this *ServerConfiguration) GetMaxEntrySize()int{
	return this.M_maxEntrySize
}

func (this *ServerConfiguration) SetMaxEntrySize(val int){
	this.NeedSave = this.M_maxEntrySize== val
	this.M_maxEntrySize= val
}


func (this *ServerConfiguration) GetHardMaxCacheSize()int{
	return this.M_hardMaxCacheSize
}

func (this *ServerConfiguration) SetHardMaxCacheSize(val int){
	this.NeedSave = this.M_hardMaxCacheSize== val
	this.M_hardMaxCacheSize= val
}


func (this *ServerConfiguration) IsVerbose()bool{
	return this.M_verbose
}

func (this *ServerConfiguration) SetVerbose(val bool){
	this.NeedSave = this.M_verbose== val
	this.M_verbose= val
}


func (this *ServerConfiguration) GetParentPtr()*Configurations{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*Configurations)
	}
	return nil
}

func (this *ServerConfiguration) SetParentPtr(val *Configurations){
	this.NeedSave = this.M_parentPtr != val.GetUuid()
	this.M_parentPtr= val.GetUuid()
}

func (this *ServerConfiguration) ResetParentPtr(){
	this.M_parentPtr= ""
}

