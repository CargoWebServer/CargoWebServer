package BPMS_Runtime

import(
"encoding/xml"
)

type LogInfo struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of LogInfo **/
	M_id string
	M_date int64
	m_actor interface{}
	/** If the ref is a string and not an object **/
	M_actor string
	M_action string
	m_object Instance
	/** If the ref is a string and not an object **/
	M_object string
	M_description string


	/** Associations **/
	m_runtimesPtr *Runtimes
	/** If the ref is a string and not an object **/
	M_runtimesPtr string
}

/** Xml parser for LogInfo **/
type XsdLogInfo struct {
	XMLName xml.Name	`xml:"logInfo"`
	M_id	string	`xml:"id,attr"`
	M_date	int64	`xml:"date,attr"`
	M_actor	string	`xml:"actor,attr"`
	M_action	string	`xml:"action,attr"`
	M_description	string	`xml:"description,attr"`

}
/** UUID **/
func (this *LogInfo) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *LogInfo) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *LogInfo) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Date **/
func (this *LogInfo) GetDate() int64{
	return this.M_date
}

/** Init reference Date **/
func (this *LogInfo) SetDate(ref interface{}){
	this.NeedSave = true
	this.M_date = ref.(int64)
}

/** Remove reference Date **/

/** Actor **/
func (this *LogInfo) GetActor() interface{}{
	return this.m_actor
}

/** Init reference Actor **/
func (this *LogInfo) SetActor(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_actor = ref.(string)
	}else{
		this.m_actor = ref.(interface{})
	}
}

/** Remove reference Actor **/

/** Action **/
func (this *LogInfo) GetAction() string{
	return this.M_action
}

/** Init reference Action **/
func (this *LogInfo) SetAction(ref interface{}){
	this.NeedSave = true
	this.M_action = ref.(string)
}

/** Remove reference Action **/

/** Object **/
func (this *LogInfo) GetObject() Instance{
	return this.m_object
}

/** Init reference Object **/
func (this *LogInfo) SetObject(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_object = ref.(string)
	}else{
		this.m_object = ref.(Instance)
		this.M_object = ref.(Instance).GetUUID()
	}
}

/** Remove reference Object **/
func (this *LogInfo) RemoveObject(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_object.(Instance).GetUUID() {
		this.m_object = nil
		this.M_object = ""
	}
}

/** Description **/
func (this *LogInfo) GetDescription() string{
	return this.M_description
}

/** Init reference Description **/
func (this *LogInfo) SetDescription(ref interface{}){
	this.NeedSave = true
	this.M_description = ref.(string)
}

/** Remove reference Description **/

/** Runtimes **/
func (this *LogInfo) GetRuntimesPtr() *Runtimes{
	return this.m_runtimesPtr
}

/** Init reference Runtimes **/
func (this *LogInfo) SetRuntimesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_runtimesPtr = ref.(string)
	}else{
		this.m_runtimesPtr = ref.(*Runtimes)
		this.M_runtimesPtr = ref.(*Runtimes).GetUUID()
	}
}

/** Remove reference Runtimes **/
func (this *LogInfo) RemoveRuntimesPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Runtimes)
	if toDelete.GetUUID() == this.m_runtimesPtr.GetUUID() {
		this.m_runtimesPtr = nil
		this.M_runtimesPtr = ""
	}
}
