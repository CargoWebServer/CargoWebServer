package BPMS_Runtime

import(
"encoding/xml"
)

type GatewayInstance struct{

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
	m_inputRef *ConnectingObject
	/** If the ref is a string and not an object **/
	M_inputRef string
	m_outputRef *ConnectingObject
	/** If the ref is a string and not an object **/
	M_outputRef string

	/** members of GatewayInstance **/
	M_gatewayType GatewayType


	/** Associations **/
	m_SubprocessInstancePtr *SubprocessInstance
	/** If the ref is a string and not an object **/
	M_SubprocessInstancePtr string
	m_processInstancePtr *ProcessInstance
	/** If the ref is a string and not an object **/
	M_processInstancePtr string
}

/** Xml parser for GatewayInstance **/
type XsdGatewayInstance struct {
	XMLName xml.Name	`xml:"gatewayInstance"`
	/** Instance **/
	M_id	string	`xml:"id,attr"`
	M_bpmnElementId	string	`xml:"bpmnElementId,attr"`


	/** FlowNodeInstance **/
	M_inputRef	*string	`xml:"inputRef"`
	M_outputRef	*string	`xml:"outputRef"`
	M_flowNodeType	string	`xml:"flowNodeType,attr"`
	M_lifecycleState	string	`xml:"lifecycleState,attr"`


	M_gatewayType	string	`xml:"gatewayType,attr"`

}
/** UUID **/
func (this *GatewayInstance) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *GatewayInstance) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *GatewayInstance) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** BpmnElementId **/
func (this *GatewayInstance) GetBpmnElementId() string{
	return this.M_bpmnElementId
}

/** Init reference BpmnElementId **/
func (this *GatewayInstance) SetBpmnElementId(ref interface{}){
	this.NeedSave = true
	this.M_bpmnElementId = ref.(string)
}

/** Remove reference BpmnElementId **/

/** Participants **/
func (this *GatewayInstance) GetParticipants() []string{
	return this.M_participants
}

/** Init reference Participants **/
func (this *GatewayInstance) SetParticipants(ref interface{}){
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
func (this *GatewayInstance) GetDataRef() []*ItemAwareElementInstance{
	return this.m_dataRef
}

/** Init reference DataRef **/
func (this *GatewayInstance) SetDataRef(ref interface{}){
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
func (this *GatewayInstance) RemoveDataRef(ref interface{}){
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
func (this *GatewayInstance) GetData() []*ItemAwareElementInstance{
	return this.M_data
}

/** Init reference Data **/
func (this *GatewayInstance) SetData(ref interface{}){
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
func (this *GatewayInstance) RemoveData(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(ItemAwareElementInstance)
	data_ := make([]*ItemAwareElementInstance, 0)
	for i := 0; i < len(this.M_data); i++ {
		if toDelete.GetUUID() != this.M_data[i].GetUUID() {
			data_ = append(data_, this.M_data[i])
		}
	}
	this.M_data = data_
}

/** LogInfoRef **/
func (this *GatewayInstance) GetLogInfoRef() []*LogInfo{
	return this.m_logInfoRef
}

/** Init reference LogInfoRef **/
func (this *GatewayInstance) SetLogInfoRef(ref interface{}){
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
func (this *GatewayInstance) RemoveLogInfoRef(ref interface{}){
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
func (this *GatewayInstance) GetFlowNodeType() FlowNodeType{
	return this.M_flowNodeType
}

/** Init reference FlowNodeType **/
func (this *GatewayInstance) SetFlowNodeType(ref interface{}){
	this.NeedSave = true
	this.M_flowNodeType = ref.(FlowNodeType)
}

/** Remove reference FlowNodeType **/

/** LifecycleState **/
func (this *GatewayInstance) GetLifecycleState() LifecycleState{
	return this.M_lifecycleState
}

/** Init reference LifecycleState **/
func (this *GatewayInstance) SetLifecycleState(ref interface{}){
	this.NeedSave = true
	this.M_lifecycleState = ref.(LifecycleState)
}

/** Remove reference LifecycleState **/

/** InputRef **/
func (this *GatewayInstance) GetInputRef() *ConnectingObject{
	return this.m_inputRef
}

/** Init reference InputRef **/
func (this *GatewayInstance) SetInputRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_inputRef = ref.(string)
	}else{
		this.m_inputRef = ref.(*ConnectingObject)
		this.M_inputRef = ref.(Instance).GetUUID()
	}
}

/** Remove reference InputRef **/
func (this *GatewayInstance) RemoveInputRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_inputRef.GetUUID() {
		this.m_inputRef = nil
		this.M_inputRef = ""
	}
}

/** OutputRef **/
func (this *GatewayInstance) GetOutputRef() *ConnectingObject{
	return this.m_outputRef
}

/** Init reference OutputRef **/
func (this *GatewayInstance) SetOutputRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_outputRef = ref.(string)
	}else{
		this.m_outputRef = ref.(*ConnectingObject)
		this.M_outputRef = ref.(Instance).GetUUID()
	}
}

/** Remove reference OutputRef **/
func (this *GatewayInstance) RemoveOutputRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_outputRef.GetUUID() {
		this.m_outputRef = nil
		this.M_outputRef = ""
	}
}

/** GatewayType **/
func (this *GatewayInstance) GetGatewayType() GatewayType{
	return this.M_gatewayType
}

/** Init reference GatewayType **/
func (this *GatewayInstance) SetGatewayType(ref interface{}){
	this.NeedSave = true
	this.M_gatewayType = ref.(GatewayType)
}

/** Remove reference GatewayType **/

/** SubprocessInstance **/
func (this *GatewayInstance) GetSubprocessInstancePtr() *SubprocessInstance{
	return this.m_SubprocessInstancePtr
}

/** Init reference SubprocessInstance **/
func (this *GatewayInstance) SetSubprocessInstancePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_SubprocessInstancePtr = ref.(string)
	}else{
		this.m_SubprocessInstancePtr = ref.(*SubprocessInstance)
		this.M_SubprocessInstancePtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference SubprocessInstance **/
func (this *GatewayInstance) RemoveSubprocessInstancePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_SubprocessInstancePtr.GetUUID() {
		this.m_SubprocessInstancePtr = nil
		this.M_SubprocessInstancePtr = ""
	}
}

/** ProcessInstance **/
func (this *GatewayInstance) GetProcessInstancePtr() *ProcessInstance{
	return this.m_processInstancePtr
}

/** Init reference ProcessInstance **/
func (this *GatewayInstance) SetProcessInstancePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_processInstancePtr = ref.(string)
	}else{
		this.m_processInstancePtr = ref.(*ProcessInstance)
		this.M_processInstancePtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference ProcessInstance **/
func (this *GatewayInstance) RemoveProcessInstancePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_processInstancePtr.GetUUID() {
		this.m_processInstancePtr = nil
		this.M_processInstancePtr = ""
	}
}
