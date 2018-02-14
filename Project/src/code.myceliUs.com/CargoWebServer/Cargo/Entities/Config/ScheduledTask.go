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
	m_parentPtr *Configurations
	/** If the ref is a string and not an object **/
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


/** Id **/
func (this *ScheduledTask) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ScheduledTask) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** IsActive **/
func (this *ScheduledTask) IsActive() bool{
	return this.M_isActive
}

/** Init reference IsActive **/
func (this *ScheduledTask) SetIsActive(ref interface{}){
	if this.M_isActive != ref.(bool) {
		this.M_isActive = ref.(bool)
		this.NeedSave = true
	}
}

/** Remove reference IsActive **/

/** Script **/
func (this *ScheduledTask) GetScript() string{
	return this.M_script
}

/** Init reference Script **/
func (this *ScheduledTask) SetScript(ref interface{}){
	if this.M_script != ref.(string) {
		this.M_script = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Script **/

/** StartTime **/
func (this *ScheduledTask) GetStartTime() int64{
	return this.M_startTime
}

/** Init reference StartTime **/
func (this *ScheduledTask) SetStartTime(ref interface{}){
	if this.M_startTime != ref.(int64) {
		this.M_startTime = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference StartTime **/

/** ExpirationTime **/
func (this *ScheduledTask) GetExpirationTime() int64{
	return this.M_expirationTime
}

/** Init reference ExpirationTime **/
func (this *ScheduledTask) SetExpirationTime(ref interface{}){
	if this.M_expirationTime != ref.(int64) {
		this.M_expirationTime = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference ExpirationTime **/

/** Frequency **/
func (this *ScheduledTask) GetFrequency() int{
	return this.M_frequency
}

/** Init reference Frequency **/
func (this *ScheduledTask) SetFrequency(ref interface{}){
	if this.M_frequency != ref.(int) {
		this.M_frequency = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference Frequency **/

/** FrequencyType **/
func (this *ScheduledTask) GetFrequencyType() FrequencyType{
	return this.M_frequencyType
}

/** Init reference FrequencyType **/
func (this *ScheduledTask) SetFrequencyType(ref interface{}){
	if this.M_frequencyType != ref.(FrequencyType) {
		this.M_frequencyType = ref.(FrequencyType)
		this.NeedSave = true
	}
}

/** Remove reference FrequencyType **/

/** Offsets **/
func (this *ScheduledTask) GetOffsets() []int{
	return this.M_offsets
}

/** Init reference Offsets **/
func (this *ScheduledTask) SetOffsets(ref interface{}){
	isExist := false
	var offsetss []int
	for i:=0; i<len(this.M_offsets); i++ {
		if this.M_offsets[i] != ref.(int) {
			offsetss = append(offsetss, this.M_offsets[i])
		} else {
			isExist = true
			offsetss = append(offsetss, ref.(int))
		}
	}
	if !isExist {
		offsetss = append(offsetss, ref.(int))
		this.NeedSave = true
		this.M_offsets = offsetss
	}
}

/** Remove reference Offsets **/

/** Parent **/
func (this *ScheduledTask) GetParentPtr() *Configurations{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *ScheduledTask) SetParentPtr(ref interface{}){
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
func (this *ScheduledTask) RemoveParentPtr(ref interface{}){
	toDelete := ref.(*Configurations)
	if this.m_parentPtr!= nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
