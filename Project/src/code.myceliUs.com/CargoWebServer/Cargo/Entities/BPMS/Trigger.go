// +build BPMS

package BPMS

import(
"encoding/xml"
)

type Trigger struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Trigger **/
	M_id string
	M_processUUID string
	M_sessionId string
	M_eventTriggerType EventTriggerType
	M_eventDatas []*EventData
	m_dataRef []*ItemAwareElementInstance
	/** If the ref is a string and not an object **/
	M_dataRef []string
	m_sourceRef *EventDefinitionInstance
	/** If the ref is a string and not an object **/
	M_sourceRef string
	m_targetRef FlowNodeInstance
	/** If the ref is a string and not an object **/
	M_targetRef string


	/** Associations **/
	m_runtimesPtr *Runtimes
	/** If the ref is a string and not an object **/
	M_runtimesPtr string
}

/** Xml parser for Trigger **/
type XsdTrigger struct {
	XMLName xml.Name	`xml:"trigger"`
	M_eventDatas	[]*XsdEventData	`xml:"eventDatas,omitempty"`
	M_dataRef	[]string	`xml:"dataRef"`
	M_sourceRef	*string	`xml:"sourceRef"`
	M_targetRef	*string	`xml:"targetRef"`
	M_id	string	`xml:"id,attr"`
	M_processUUID	string	`xml:"processUUID,attr"`
	M_sessionId	string	`xml:"sessionId,attr"`
	M_eventTriggerType	string	`xml:"eventTriggerType,attr"`

}
/** UUID **/
func (this *Trigger) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Trigger) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Trigger) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** ProcessUUID **/
func (this *Trigger) GetProcessUUID() string{
	return this.M_processUUID
}

/** Init reference ProcessUUID **/
func (this *Trigger) SetProcessUUID(ref interface{}){
	this.NeedSave = true
	this.M_processUUID = ref.(string)
}

/** Remove reference ProcessUUID **/

/** SessionId **/
func (this *Trigger) GetSessionId() string{
	return this.M_sessionId
}

/** Init reference SessionId **/
func (this *Trigger) SetSessionId(ref interface{}){
	this.NeedSave = true
	this.M_sessionId = ref.(string)
}

/** Remove reference SessionId **/

/** EventTriggerType **/
func (this *Trigger) GetEventTriggerType() EventTriggerType{
	return this.M_eventTriggerType
}

/** Init reference EventTriggerType **/
func (this *Trigger) SetEventTriggerType(ref interface{}){
	this.NeedSave = true
	this.M_eventTriggerType = ref.(EventTriggerType)
}

/** Remove reference EventTriggerType **/

/** EventDatas **/
func (this *Trigger) GetEventDatas() []*EventData{
	return this.M_eventDatas
}

/** Init reference EventDatas **/
func (this *Trigger) SetEventDatas(ref interface{}){
	this.NeedSave = true
	isExist := false
	var eventDatass []*EventData
	for i:=0; i<len(this.M_eventDatas); i++ {
		if this.M_eventDatas[i].GetUUID() != ref.(*EventData).GetUUID() {
			eventDatass = append(eventDatass, this.M_eventDatas[i])
		} else {
			isExist = true
			eventDatass = append(eventDatass, ref.(*EventData))
		}
	}
	if !isExist {
		eventDatass = append(eventDatass, ref.(*EventData))
	}
	this.M_eventDatas = eventDatass
}

/** Remove reference EventDatas **/
func (this *Trigger) RemoveEventDatas(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*EventData)
	eventDatas_ := make([]*EventData, 0)
	for i := 0; i < len(this.M_eventDatas); i++ {
		if toDelete.GetUUID() != this.M_eventDatas[i].GetUUID() {
			eventDatas_ = append(eventDatas_, this.M_eventDatas[i])
		}
	}
	this.M_eventDatas = eventDatas_
}

/** DataRef **/
func (this *Trigger) GetDataRef() []*ItemAwareElementInstance{
	return this.m_dataRef
}

/** Init reference DataRef **/
func (this *Trigger) SetDataRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_dataRef); i++ {
			if this.M_dataRef[i] == refStr {
				return
			}
		}
		this.M_dataRef = append(this.M_dataRef, ref.(string))
	}else{
		this.RemoveDataRef(ref)
		this.m_dataRef = append(this.m_dataRef, ref.(*ItemAwareElementInstance))
		this.M_dataRef = append(this.M_dataRef, ref.(*ItemAwareElementInstance).GetUUID())
	}
}

/** Remove reference DataRef **/
func (this *Trigger) RemoveDataRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ItemAwareElementInstance)
	dataRef_ := make([]*ItemAwareElementInstance, 0)
	dataRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_dataRef); i++ {
		if toDelete.GetUUID() != this.m_dataRef[i].GetUUID() {
			dataRef_ = append(dataRef_, this.m_dataRef[i])
			dataRefUuid = append(dataRefUuid, this.M_dataRef[i])
		}
	}
	this.m_dataRef = dataRef_
	this.M_dataRef = dataRefUuid
}

/** SourceRef **/
func (this *Trigger) GetSourceRef() *EventDefinitionInstance{
	return this.m_sourceRef
}

/** Init reference SourceRef **/
func (this *Trigger) SetSourceRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_sourceRef = ref.(string)
	}else{
		this.m_sourceRef = ref.(*EventDefinitionInstance)
		this.M_sourceRef = ref.(Instance).GetUUID()
	}
}

/** Remove reference SourceRef **/
func (this *Trigger) RemoveSourceRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_sourceRef.GetUUID() {
		this.m_sourceRef = nil
		this.M_sourceRef = ""
	}
}

/** TargetRef **/
func (this *Trigger) GetTargetRef() FlowNodeInstance{
	return this.m_targetRef
}

/** Init reference TargetRef **/
func (this *Trigger) SetTargetRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_targetRef = ref.(string)
	}else{
		this.m_targetRef = ref.(FlowNodeInstance)
		this.M_targetRef = ref.(Instance).GetUUID()
	}
}

/** Remove reference TargetRef **/
func (this *Trigger) RemoveTargetRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_targetRef.(Instance).GetUUID() {
		this.m_targetRef = nil
		this.M_targetRef = ""
	}
}

/** Runtimes **/
func (this *Trigger) GetRuntimesPtr() *Runtimes{
	return this.m_runtimesPtr
}

/** Init reference Runtimes **/
func (this *Trigger) SetRuntimesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_runtimesPtr = ref.(string)
	}else{
		this.m_runtimesPtr = ref.(*Runtimes)
		this.M_runtimesPtr = ref.(*Runtimes).GetUUID()
	}
}

/** Remove reference Runtimes **/
func (this *Trigger) RemoveRuntimesPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Runtimes)
	if toDelete.GetUUID() == this.m_runtimesPtr.GetUUID() {
		this.m_runtimesPtr = nil
		this.M_runtimesPtr = ""
	}
}
