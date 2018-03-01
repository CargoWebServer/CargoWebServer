// +build Config

package Config

import(
	"encoding/xml"
)

type ScheduledTask struct{

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

	/** members of ScheduledTask **/
	M_isActive bool
	M_script string
	M_startTime int64
	M_expirationTime int64
	M_frequency int
	M_frequencyType FrequencyType
	M_offsets []int


	/** Associations **/
	M_parentPtr string
}

/** Xml parser for ScheduledTask **/
type XsdScheduledTask struct {
	XMLName xml.Name	`xml:"scheduledTask"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_isActive	bool	`xml:"isActive,attr"`
	M_script	string	`xml:"script,attr"`
	M_startTime	int64	`xml:"startTime,attr"`
	M_expirationTime	int64	`xml:"expirationTime,attr"`
	M_frequency	int	`xml:"frequency,attr"`
	M_frequencyType	string	`xml:"frequencyType,attr"`
	M_offsets	int	`xml:"offsets,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *ScheduledTask) GetUuid() string{
	return this.UUID
}
func (this *ScheduledTask) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *ScheduledTask) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *ScheduledTask) GetTypeName() string{
	this.TYPENAME = "Config.ScheduledTask"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *ScheduledTask) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *ScheduledTask) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ScheduledTask) GetParentLnk() string{
	return this.ParentLnk
}
func (this *ScheduledTask) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *ScheduledTask) IsNeedSave() bool{
	return this.NeedSave
}
func (this *ScheduledTask) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *ScheduledTask) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

func (this *ScheduledTask) GetId()string{
	return this.M_id
}

func (this *ScheduledTask) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *ScheduledTask) IsActive()bool{
	return this.M_isActive
}

func (this *ScheduledTask) SetIsActive(val bool){
	this.NeedSave = this.M_isActive== val
	this.M_isActive= val
}


func (this *ScheduledTask) GetScript()string{
	return this.M_script
}

func (this *ScheduledTask) SetScript(val string){
	this.NeedSave = this.M_script== val
	this.M_script= val
}


func (this *ScheduledTask) GetStartTime()int64{
	return this.M_startTime
}

func (this *ScheduledTask) SetStartTime(val int64){
	this.NeedSave = this.M_startTime== val
	this.M_startTime= val
}


func (this *ScheduledTask) GetExpirationTime()int64{
	return this.M_expirationTime
}

func (this *ScheduledTask) SetExpirationTime(val int64){
	this.NeedSave = this.M_expirationTime== val
	this.M_expirationTime= val
}


func (this *ScheduledTask) GetFrequency()int{
	return this.M_frequency
}

func (this *ScheduledTask) SetFrequency(val int){
	this.NeedSave = this.M_frequency== val
	this.M_frequency= val
}


func (this *ScheduledTask) GetFrequencyType()FrequencyType{
	return this.M_frequencyType
}

func (this *ScheduledTask) SetFrequencyType(val FrequencyType){
	this.NeedSave = this.M_frequencyType== val
	this.M_frequencyType= val
}

func (this *ScheduledTask) ResetFrequencyType(){
	this.M_frequencyType= 0
}


func (this *ScheduledTask) GetOffsets()[]int{
	return this.M_offsets
}

func (this *ScheduledTask) SetOffsets(val []int){
	this.M_offsets= val
}

func (this *ScheduledTask) AppendOffsets(val int){
	this.M_offsets=append(this.M_offsets, val)
	this.NeedSave= true
}


func (this *ScheduledTask) GetParentPtr()*Configurations{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*Configurations)
	}
	return nil
}

func (this *ScheduledTask) SetParentPtr(val *Configurations){
	this.NeedSave = this.M_parentPtr != val.GetUuid()
	this.M_parentPtr= val.GetUuid()
}

func (this *ScheduledTask) ResetParentPtr(){
	this.M_parentPtr= ""
}

