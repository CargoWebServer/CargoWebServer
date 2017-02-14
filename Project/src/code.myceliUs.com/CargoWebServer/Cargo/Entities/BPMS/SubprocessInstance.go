// +build BPMS

package BPMS

import(
"encoding/xml"
)

type SubprocessInstance struct{

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

	/** members of SubprocessInstance **/
	M_SubprocessType SubprocessType
	M_flowNodeInstances []FlowNodeInstance
	M_connectingObjects []*ConnectingObject
	M_ressources []*RessourceInstance


	/** Associations **/
	m_SubprocessInstancePtr *SubprocessInstance
	/** If the ref is a string and not an object **/
	M_SubprocessInstancePtr string
	m_processInstancePtr *ProcessInstance
	/** If the ref is a string and not an object **/
	M_processInstancePtr string
}

/** Xml parser for SubprocessInstance **/
type XsdSubprocessInstance struct {
	XMLName xml.Name	`xml:"SubprocessInstance"`
	/** Instance **/
	M_id	string	`xml:"id,attr"`
	M_bpmnElementId	string	`xml:"bpmnElementId,attr"`


	/** FlowNodeInstance **/
	M_inputRef	[]string	`xml:"inputRef"`
	M_outputRef	[]string	`xml:"outputRef"`
	M_flowNodeType	string	`xml:"flowNodeType,attr"`
	M_lifecycleState	string	`xml:"lifecycleState,attr"`


	M_flowNodeInstances_0	[]*XsdActivityInstance	`xml:"activityInstance,omitempty"`
	M_flowNodeInstances_1	[]*XsdSubprocessInstance	`xml:"SubprocessInstance,omitempty"`
	M_flowNodeInstances_2	[]*XsdGatewayInstance	`xml:"gatewayInstance,omitempty"`
	M_flowNodeInstances_3	[]*XsdEventInstance	`xml:"eventInstance,omitempty"`

	M_connectingObjects	[]*XsdConnectingObject	`xml:"connectingObjects,omitempty"`
	M_ressources	[]*XsdRessourceInstance	`xml:"ressources,omitempty"`
	M_SubprocessType	string	`xml:"SubprocessType,attr"`

}
/** UUID **/
func (this *SubprocessInstance) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *SubprocessInstance) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *SubprocessInstance) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** BpmnElementId **/
func (this *SubprocessInstance) GetBpmnElementId() string{
	return this.M_bpmnElementId
}

/** Init reference BpmnElementId **/
func (this *SubprocessInstance) SetBpmnElementId(ref interface{}){
	this.NeedSave = true
	this.M_bpmnElementId = ref.(string)
}

/** Remove reference BpmnElementId **/

/** Participants **/
func (this *SubprocessInstance) GetParticipants() []string{
	return this.M_participants
}

/** Init reference Participants **/
func (this *SubprocessInstance) SetParticipants(ref interface{}){
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
func (this *SubprocessInstance) GetDataRef() []*ItemAwareElementInstance{
	return this.m_dataRef
}

/** Init reference DataRef **/
func (this *SubprocessInstance) SetDataRef(ref interface{}){
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
func (this *SubprocessInstance) RemoveDataRef(ref interface{}){
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
func (this *SubprocessInstance) GetData() []*ItemAwareElementInstance{
	return this.M_data
}

/** Init reference Data **/
func (this *SubprocessInstance) SetData(ref interface{}){
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
func (this *SubprocessInstance) RemoveData(ref interface{}){
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
func (this *SubprocessInstance) GetLogInfoRef() []*LogInfo{
	return this.m_logInfoRef
}

/** Init reference LogInfoRef **/
func (this *SubprocessInstance) SetLogInfoRef(ref interface{}){
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
func (this *SubprocessInstance) RemoveLogInfoRef(ref interface{}){
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
func (this *SubprocessInstance) GetFlowNodeType() FlowNodeType{
	return this.M_flowNodeType
}

/** Init reference FlowNodeType **/
func (this *SubprocessInstance) SetFlowNodeType(ref interface{}){
	this.NeedSave = true
	this.M_flowNodeType = ref.(FlowNodeType)
}

/** Remove reference FlowNodeType **/

/** LifecycleState **/
func (this *SubprocessInstance) GetLifecycleState() LifecycleState{
	return this.M_lifecycleState
}

/** Init reference LifecycleState **/
func (this *SubprocessInstance) SetLifecycleState(ref interface{}){
	this.NeedSave = true
	this.M_lifecycleState = ref.(LifecycleState)
}

/** Remove reference LifecycleState **/

/** InputRef **/
func (this *SubprocessInstance) GetInputRef() []*ConnectingObject{
	return this.m_inputRef
}

/** Init reference InputRef **/
func (this *SubprocessInstance) SetInputRef(ref interface{}){
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
func (this *SubprocessInstance) RemoveInputRef(ref interface{}){
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
func (this *SubprocessInstance) GetOutputRef() []*ConnectingObject{
	return this.m_outputRef
}

/** Init reference OutputRef **/
func (this *SubprocessInstance) SetOutputRef(ref interface{}){
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
func (this *SubprocessInstance) RemoveOutputRef(ref interface{}){
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

/** SubprocessType **/
func (this *SubprocessInstance) GetSubprocessType() SubprocessType{
	return this.M_SubprocessType
}

/** Init reference SubprocessType **/
func (this *SubprocessInstance) SetSubprocessType(ref interface{}){
	this.NeedSave = true
	this.M_SubprocessType = ref.(SubprocessType)
}

/** Remove reference SubprocessType **/

/** FlowNodeInstances **/
func (this *SubprocessInstance) GetFlowNodeInstances() []FlowNodeInstance{
	return this.M_flowNodeInstances
}

/** Init reference FlowNodeInstances **/
func (this *SubprocessInstance) SetFlowNodeInstances(ref interface{}){
	this.NeedSave = true
	isExist := false
	var flowNodeInstancess []FlowNodeInstance
	for i:=0; i<len(this.M_flowNodeInstances); i++ {
		if this.M_flowNodeInstances[i].(Instance).GetUUID() != ref.(Instance).GetUUID() {
			flowNodeInstancess = append(flowNodeInstancess, this.M_flowNodeInstances[i])
		} else {
			isExist = true
			flowNodeInstancess = append(flowNodeInstancess, ref.(FlowNodeInstance))
		}
	}
	if !isExist {
		flowNodeInstancess = append(flowNodeInstancess, ref.(FlowNodeInstance))
	}
	this.M_flowNodeInstances = flowNodeInstancess
}

/** Remove reference FlowNodeInstances **/
func (this *SubprocessInstance) RemoveFlowNodeInstances(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	flowNodeInstances_ := make([]FlowNodeInstance, 0)
	for i := 0; i < len(this.M_flowNodeInstances); i++ {
		if toDelete.GetUUID() != this.M_flowNodeInstances[i].(Instance).GetUUID() {
			flowNodeInstances_ = append(flowNodeInstances_, this.M_flowNodeInstances[i])
		}
	}
	this.M_flowNodeInstances = flowNodeInstances_
}

/** ConnectingObjects **/
func (this *SubprocessInstance) GetConnectingObjects() []*ConnectingObject{
	return this.M_connectingObjects
}

/** Init reference ConnectingObjects **/
func (this *SubprocessInstance) SetConnectingObjects(ref interface{}){
	this.NeedSave = true
	isExist := false
	var connectingObjectss []*ConnectingObject
	for i:=0; i<len(this.M_connectingObjects); i++ {
		if this.M_connectingObjects[i].GetUUID() != ref.(Instance).GetUUID() {
			connectingObjectss = append(connectingObjectss, this.M_connectingObjects[i])
		} else {
			isExist = true
			connectingObjectss = append(connectingObjectss, ref.(*ConnectingObject))
		}
	}
	if !isExist {
		connectingObjectss = append(connectingObjectss, ref.(*ConnectingObject))
	}
	this.M_connectingObjects = connectingObjectss
}

/** Remove reference ConnectingObjects **/
func (this *SubprocessInstance) RemoveConnectingObjects(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	connectingObjects_ := make([]*ConnectingObject, 0)
	for i := 0; i < len(this.M_connectingObjects); i++ {
		if toDelete.GetUUID() != this.M_connectingObjects[i].GetUUID() {
			connectingObjects_ = append(connectingObjects_, this.M_connectingObjects[i])
		}
	}
	this.M_connectingObjects = connectingObjects_
}

/** Ressources **/
func (this *SubprocessInstance) GetRessources() []*RessourceInstance{
	return this.M_ressources
}

/** Init reference Ressources **/
func (this *SubprocessInstance) SetRessources(ref interface{}){
	this.NeedSave = true
	isExist := false
	var ressourcess []*RessourceInstance
	for i:=0; i<len(this.M_ressources); i++ {
		if this.M_ressources[i].GetUUID() != ref.(*RessourceInstance).GetUUID() {
			ressourcess = append(ressourcess, this.M_ressources[i])
		} else {
			isExist = true
			ressourcess = append(ressourcess, ref.(*RessourceInstance))
		}
	}
	if !isExist {
		ressourcess = append(ressourcess, ref.(*RessourceInstance))
	}
	this.M_ressources = ressourcess
}

/** Remove reference Ressources **/
func (this *SubprocessInstance) RemoveRessources(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*RessourceInstance)
	ressources_ := make([]*RessourceInstance, 0)
	for i := 0; i < len(this.M_ressources); i++ {
		if toDelete.GetUUID() != this.M_ressources[i].GetUUID() {
			ressources_ = append(ressources_, this.M_ressources[i])
		}
	}
	this.M_ressources = ressources_
}

/** SubprocessInstance **/
func (this *SubprocessInstance) GetSubprocessInstancePtr() *SubprocessInstance{
	return this.m_SubprocessInstancePtr
}

/** Init reference SubprocessInstance **/
func (this *SubprocessInstance) SetSubprocessInstancePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_SubprocessInstancePtr = ref.(string)
	}else{
		this.m_SubprocessInstancePtr = ref.(*SubprocessInstance)
		this.M_SubprocessInstancePtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference SubprocessInstance **/
func (this *SubprocessInstance) RemoveSubprocessInstancePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_SubprocessInstancePtr.GetUUID() {
		this.m_SubprocessInstancePtr = nil
		this.M_SubprocessInstancePtr = ""
	}
}

/** ProcessInstance **/
func (this *SubprocessInstance) GetProcessInstancePtr() *ProcessInstance{
	return this.m_processInstancePtr
}

/** Init reference ProcessInstance **/
func (this *SubprocessInstance) SetProcessInstancePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_processInstancePtr = ref.(string)
	}else{
		this.m_processInstancePtr = ref.(*ProcessInstance)
		this.M_processInstancePtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference ProcessInstance **/
func (this *SubprocessInstance) RemoveProcessInstancePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_processInstancePtr.GetUUID() {
		this.m_processInstancePtr = nil
		this.M_processInstancePtr = ""
	}
}
