package BPMS

import(
"encoding/xml"
)

type EventInstance struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Instance **/
	M_id string
	M_bpmnElementId string
	M_participants []string
	m_dataRef []*ItemAwareElementInstance
	/** If the ref is a string and not an object **/
	M_dataRef []string
	M_data []*ItemAwareElementInstance
	m_logInfoRef []*LogInfo
	/** If the ref is a string and not an object **/
	M_logInfoRef []string

	/** members of FlowNodeInstance **/
	M_flowNodeType FlowNodeType
	M_lifecycleState LifecycleState
	m_inputRef []*ConnectingObject
	/** If the ref is a string and not an object **/
	M_inputRef []string
	m_outputRef []*ConnectingObject
	/** If the ref is a string and not an object **/
	M_outputRef []string

	/** members of EventInstance **/
	M_eventType EventType
	M_eventDefintionInstances []*EventDefinitionInstance


	/** Associations **/
	m_SubprocessInstancePtr *SubprocessInstance
	/** If the ref is a string and not an object **/
	M_SubprocessInstancePtr string
	m_processInstancePtr *ProcessInstance
	/** If the ref is a string and not an object **/
	M_processInstancePtr string
}

/** Xml parser for EventInstance **/
type XsdEventInstance struct {
	XMLName xml.Name	`xml:"eventInstance"`
	/** Instance **/
	M_id	string	`xml:"id,attr"`
	M_bpmnElementId	string	`xml:"bpmnElementId,attr"`


	/** FlowNodeInstance **/
	M_inputRef	[]string	`xml:"inputRef"`
	M_outputRef	[]string	`xml:"outputRef"`
	M_flowNodeType	string	`xml:"flowNodeType,attr"`
	M_lifecycleState	string	`xml:"lifecycleState,attr"`


	M_eventDefintionInstances	[]*XsdEventDefinitionInstance	`xml:"eventDefintionInstances,omitempty"`
	M_eventType	string	`xml:"eventType,attr"`

}
/** UUID **/
func (this *EventInstance) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *EventInstance) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *EventInstance) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** BpmnElementId **/
func (this *EventInstance) GetBpmnElementId() string{
	return this.M_bpmnElementId
}

/** Init reference BpmnElementId **/
func (this *EventInstance) SetBpmnElementId(ref interface{}){
	this.NeedSave = true
	this.M_bpmnElementId = ref.(string)
}

/** Remove reference BpmnElementId **/

/** Participants **/
func (this *EventInstance) GetParticipants() []string{
	return this.M_participants
}

/** Init reference Participants **/
func (this *EventInstance) SetParticipants(ref interface{}){
	this.NeedSave = true
	isExist := false
	var participantss []string
	for i:=0; i<len(this.M_participants); i++ {
		if this.M_participants[i] != ref.(string) {
			participantss = append(participantss, this.M_participants[i])
		} else {
			isExist = true
			participantss = append(participantss, ref.(string))
		}
	}
	if !isExist {
		participantss = append(participantss, ref.(string))
	}
	this.M_participants = participantss
}

/** Remove reference Participants **/

/** DataRef **/
func (this *EventInstance) GetDataRef() []*ItemAwareElementInstance{
	return this.m_dataRef
}

/** Init reference DataRef **/
func (this *EventInstance) SetDataRef(ref interface{}){
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
func (this *EventInstance) RemoveDataRef(ref interface{}){
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

/** Data **/
func (this *EventInstance) GetData() []*ItemAwareElementInstance{
	return this.M_data
}

/** Init reference Data **/
func (this *EventInstance) SetData(ref interface{}){
	this.NeedSave = true
	isExist := false
	var datas []*ItemAwareElementInstance
	for i:=0; i<len(this.M_data); i++ {
		if this.M_data[i].GetUUID() != ref.(*ItemAwareElementInstance).GetUUID() {
			datas = append(datas, this.M_data[i])
		} else {
			isExist = true
			datas = append(datas, ref.(*ItemAwareElementInstance))
		}
	}
	if !isExist {
		datas = append(datas, ref.(*ItemAwareElementInstance))
	}
	this.M_data = datas
}

/** Remove reference Data **/
func (this *EventInstance) RemoveData(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ItemAwareElementInstance)
	data_ := make([]*ItemAwareElementInstance, 0)
	for i := 0; i < len(this.M_data); i++ {
		if toDelete.GetUUID() != this.M_data[i].GetUUID() {
			data_ = append(data_, this.M_data[i])
		}
	}
	this.M_data = data_
}

/** LogInfoRef **/
func (this *EventInstance) GetLogInfoRef() []*LogInfo{
	return this.m_logInfoRef
}

/** Init reference LogInfoRef **/
func (this *EventInstance) SetLogInfoRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_logInfoRef); i++ {
			if this.M_logInfoRef[i] == refStr {
				return
			}
		}
		this.M_logInfoRef = append(this.M_logInfoRef, ref.(string))
	}else{
		this.RemoveLogInfoRef(ref)
		this.m_logInfoRef = append(this.m_logInfoRef, ref.(*LogInfo))
		this.M_logInfoRef = append(this.M_logInfoRef, ref.(*LogInfo).GetUUID())
	}
}

/** Remove reference LogInfoRef **/
func (this *EventInstance) RemoveLogInfoRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*LogInfo)
	logInfoRef_ := make([]*LogInfo, 0)
	logInfoRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_logInfoRef); i++ {
		if toDelete.GetUUID() != this.m_logInfoRef[i].GetUUID() {
			logInfoRef_ = append(logInfoRef_, this.m_logInfoRef[i])
			logInfoRefUuid = append(logInfoRefUuid, this.M_logInfoRef[i])
		}
	}
	this.m_logInfoRef = logInfoRef_
	this.M_logInfoRef = logInfoRefUuid
}

/** FlowNodeType **/
func (this *EventInstance) GetFlowNodeType() FlowNodeType{
	return this.M_flowNodeType
}

/** Init reference FlowNodeType **/
func (this *EventInstance) SetFlowNodeType(ref interface{}){
	this.NeedSave = true
	this.M_flowNodeType = ref.(FlowNodeType)
}

/** Remove reference FlowNodeType **/

/** LifecycleState **/
func (this *EventInstance) GetLifecycleState() LifecycleState{
	return this.M_lifecycleState
}

/** Init reference LifecycleState **/
func (this *EventInstance) SetLifecycleState(ref interface{}){
	this.NeedSave = true
	this.M_lifecycleState = ref.(LifecycleState)
}

/** Remove reference LifecycleState **/

/** InputRef **/
func (this *EventInstance) GetInputRef() []*ConnectingObject{
	return this.m_inputRef
}

/** Init reference InputRef **/
func (this *EventInstance) SetInputRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_inputRef); i++ {
			if this.M_inputRef[i] == refStr {
				return
			}
		}
		this.M_inputRef = append(this.M_inputRef, ref.(string))
	}else{
		this.RemoveInputRef(ref)
		this.m_inputRef = append(this.m_inputRef, ref.(*ConnectingObject))
		this.M_inputRef = append(this.M_inputRef, ref.(Instance).GetUUID())
	}
}

/** Remove reference InputRef **/
func (this *EventInstance) RemoveInputRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	inputRef_ := make([]*ConnectingObject, 0)
	inputRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_inputRef); i++ {
		if toDelete.GetUUID() != this.m_inputRef[i].GetUUID() {
			inputRef_ = append(inputRef_, this.m_inputRef[i])
			inputRefUuid = append(inputRefUuid, this.M_inputRef[i])
		}
	}
	this.m_inputRef = inputRef_
	this.M_inputRef = inputRefUuid
}

/** OutputRef **/
func (this *EventInstance) GetOutputRef() []*ConnectingObject{
	return this.m_outputRef
}

/** Init reference OutputRef **/
func (this *EventInstance) SetOutputRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_outputRef); i++ {
			if this.M_outputRef[i] == refStr {
				return
			}
		}
		this.M_outputRef = append(this.M_outputRef, ref.(string))
	}else{
		this.RemoveOutputRef(ref)
		this.m_outputRef = append(this.m_outputRef, ref.(*ConnectingObject))
		this.M_outputRef = append(this.M_outputRef, ref.(Instance).GetUUID())
	}
}

/** Remove reference OutputRef **/
func (this *EventInstance) RemoveOutputRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	outputRef_ := make([]*ConnectingObject, 0)
	outputRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_outputRef); i++ {
		if toDelete.GetUUID() != this.m_outputRef[i].GetUUID() {
			outputRef_ = append(outputRef_, this.m_outputRef[i])
			outputRefUuid = append(outputRefUuid, this.M_outputRef[i])
		}
	}
	this.m_outputRef = outputRef_
	this.M_outputRef = outputRefUuid
}

/** EventType **/
func (this *EventInstance) GetEventType() EventType{
	return this.M_eventType
}

/** Init reference EventType **/
func (this *EventInstance) SetEventType(ref interface{}){
	this.NeedSave = true
	this.M_eventType = ref.(EventType)
}

/** Remove reference EventType **/

/** EventDefintionInstances **/
func (this *EventInstance) GetEventDefintionInstances() []*EventDefinitionInstance{
	return this.M_eventDefintionInstances
}

/** Init reference EventDefintionInstances **/
func (this *EventInstance) SetEventDefintionInstances(ref interface{}){
	this.NeedSave = true
	isExist := false
	var eventDefintionInstancess []*EventDefinitionInstance
	for i:=0; i<len(this.M_eventDefintionInstances); i++ {
		if this.M_eventDefintionInstances[i].GetUUID() != ref.(Instance).GetUUID() {
			eventDefintionInstancess = append(eventDefintionInstancess, this.M_eventDefintionInstances[i])
		} else {
			isExist = true
			eventDefintionInstancess = append(eventDefintionInstancess, ref.(*EventDefinitionInstance))
		}
	}
	if !isExist {
		eventDefintionInstancess = append(eventDefintionInstancess, ref.(*EventDefinitionInstance))
	}
	this.M_eventDefintionInstances = eventDefintionInstancess
}

/** Remove reference EventDefintionInstances **/
func (this *EventInstance) RemoveEventDefintionInstances(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	eventDefintionInstances_ := make([]*EventDefinitionInstance, 0)
	for i := 0; i < len(this.M_eventDefintionInstances); i++ {
		if toDelete.GetUUID() != this.M_eventDefintionInstances[i].GetUUID() {
			eventDefintionInstances_ = append(eventDefintionInstances_, this.M_eventDefintionInstances[i])
		}
	}
	this.M_eventDefintionInstances = eventDefintionInstances_
}

/** SubprocessInstance **/
func (this *EventInstance) GetSubprocessInstancePtr() *SubprocessInstance{
	return this.m_SubprocessInstancePtr
}

/** Init reference SubprocessInstance **/
func (this *EventInstance) SetSubprocessInstancePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_SubprocessInstancePtr = ref.(string)
	}else{
		this.m_SubprocessInstancePtr = ref.(*SubprocessInstance)
		this.M_SubprocessInstancePtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference SubprocessInstance **/
func (this *EventInstance) RemoveSubprocessInstancePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_SubprocessInstancePtr.GetUUID() {
		this.m_SubprocessInstancePtr = nil
		this.M_SubprocessInstancePtr = ""
	}
}

/** ProcessInstance **/
func (this *EventInstance) GetProcessInstancePtr() *ProcessInstance{
	return this.m_processInstancePtr
}

/** Init reference ProcessInstance **/
func (this *EventInstance) SetProcessInstancePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_processInstancePtr = ref.(string)
	}else{
		this.m_processInstancePtr = ref.(*ProcessInstance)
		this.M_processInstancePtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference ProcessInstance **/
func (this *EventInstance) RemoveProcessInstancePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_processInstancePtr.GetUUID() {
		this.m_processInstancePtr = nil
		this.M_processInstancePtr = ""
	}
}
