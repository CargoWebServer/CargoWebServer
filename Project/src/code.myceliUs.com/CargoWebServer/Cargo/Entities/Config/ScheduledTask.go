// +build Config

package Config

import (
	"encoding/xml"

	"code.myceliUs.com/Utility"
)

type ScheduledTask struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** keep track if the entity has change over time. **/
	needSave bool
	/** Keep reference to entity that made use of thit entity **/
	Referenced []string
	/** Get entity by uuid function **/
	getEntityByUuid func(string) (interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of Configuration **/
	M_id string

	/** members of ScheduledTask **/
	M_isActive       bool
	M_script         string
	M_startTime      int64
	M_expirationTime int64
	M_frequency      int
	M_frequencyType  FrequencyType
	M_offsets        []int

	/** Associations **/
	M_parentPtr string
}

/** Xml parser for ScheduledTask **/
type XsdScheduledTask struct {
	XMLName xml.Name `xml:"scheduledTask"`
	/** Configuration **/
	M_id string `xml:"id,attr"`

	M_isActive       bool   `xml:"isActive,attr"`
	M_script         string `xml:"script,attr"`
	M_startTime      int64  `xml:"startTime,attr"`
	M_expirationTime int64  `xml:"expirationTime,attr"`
	M_frequency      int    `xml:"frequency,attr"`
	M_frequencyType  string `xml:"frequencyType,attr"`
	M_offsets        []int  `xml:"offsets,attr"`
}

/***************** Entity **************************/

/** UUID **/
func (this *ScheduledTask) GetUuid() string {
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *ScheduledTask) SetUuid(uuid string) {
	this.UUID = uuid
}

/** Need save **/
func (this *ScheduledTask) IsNeedSave() bool {
	return this.needSave
}
func (this *ScheduledTask) SetNeedSave(needSave bool) {
	this.needSave = needSave
}

func (this *ScheduledTask) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *ScheduledTask) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *ScheduledTask) RemoveReferenced(uuid string, field string) {
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

func (this *ScheduledTask) SetFieldValue(field string, value interface{}) error {
	return Utility.SetProperty(this, field, value)
}

func (this *ScheduledTask) GetFieldValue(field string) interface{} {
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *ScheduledTask) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *ScheduledTask) GetTypeName() string {
	this.TYPENAME = "Config.ScheduledTask"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *ScheduledTask) GetParentUuid() string {
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *ScheduledTask) SetParentUuid(parentUuid string) {
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ScheduledTask) GetParentLnk() string {
	return this.ParentLnk
}
func (this *ScheduledTask) SetParentLnk(parentLnk string) {
	this.ParentLnk = parentLnk
}

func (this *ScheduledTask) GetParent() interface{} {
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ScheduledTask) GetChilds() []interface{} {
	var childs []interface{}
	return childs
}

/** Return the list of all childs uuid **/
func (this *ScheduledTask) GetChildsUuid() []string {
	var childs []string
	return childs
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *ScheduledTask) SetEntityGetter(fct func(uuid string) (interface{}, error)) {
	this.getEntityByUuid = fct
}

/** Use it the set the entity on the cache. **/
func (this *ScheduledTask) SetEntitySetter(fct func(entity interface{})) {
	this.setEntity = fct
}

/** Set the uuid generator function **/
func (this *ScheduledTask) SetUuidGenerator(fct func(entity interface{}) string) {
	this.generateUuid = fct
}

func (this *ScheduledTask) GetId() string {
	return this.M_id
}

func (this *ScheduledTask) SetId(val string) {
	this.M_id = val
}

func (this *ScheduledTask) IsActive() bool {
	return this.M_isActive
}

func (this *ScheduledTask) SetIsActive(val bool) {
	this.M_isActive = val
}

func (this *ScheduledTask) GetScript() string {
	return this.M_script
}

func (this *ScheduledTask) SetScript(val string) {
	this.M_script = val
}

func (this *ScheduledTask) GetStartTime() int64 {
	return this.M_startTime
}

func (this *ScheduledTask) SetStartTime(val int64) {
	this.M_startTime = val
}

func (this *ScheduledTask) GetExpirationTime() int64 {
	return this.M_expirationTime
}

func (this *ScheduledTask) SetExpirationTime(val int64) {
	this.M_expirationTime = val
}

func (this *ScheduledTask) GetFrequency() int {
	return this.M_frequency
}

func (this *ScheduledTask) SetFrequency(val int) {
	this.M_frequency = val
}

func (this *ScheduledTask) GetFrequencyType() FrequencyType {
	return this.M_frequencyType
}

func (this *ScheduledTask) SetFrequencyType(val FrequencyType) {
	this.M_frequencyType = val
}

func (this *ScheduledTask) ResetFrequencyType() {
	this.M_frequencyType = 0
	this.setEntity(this)
}

func (this *ScheduledTask) GetOffsets() []int {
	return this.M_offsets
}

func (this *ScheduledTask) SetOffsets(val []int) {
	this.M_offsets = val
}

func (this *ScheduledTask) AppendOffsets(val int) {
	if this.M_offsets == nil {
		this.M_offsets = make([]int, 0)
	}

	this.M_offsets = append(this.M_offsets, val)
	this.setEntity(this)
	this.SetNeedSave(true)
}

func (this *ScheduledTask) GetParentPtr() *Configurations {
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*Configurations)
	}
	return nil
}

func (this *ScheduledTask) SetParentPtr(val *Configurations) {
	this.M_parentPtr = val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}

func (this *ScheduledTask) ResetParentPtr() {
	this.M_parentPtr = ""
	this.setEntity(this)
}
