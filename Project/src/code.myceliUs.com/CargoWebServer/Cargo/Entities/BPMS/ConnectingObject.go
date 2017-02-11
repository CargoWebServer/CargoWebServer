package BPMS

import(
"encoding/xml"
)

type ConnectingObject struct{

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

	/** members of ConnectingObject **/
	M_connectingObjectType ConnectingObjectType
	m_sourceRef FlowNodeInstance
	/** If the ref is a string and not an object **/
	M_sourceRef string
	m_targetRef FlowNodeInstance
	/** If the ref is a string and not an object **/
	M_targetRef string


	/** Associations **/
	m_SubprocessInstancePtr *SubprocessInstance
	/** If the ref is a string and not an object **/
	M_SubprocessInstancePtr string
	m_processInstancePtr *ProcessInstance
	/** If the ref is a string and not an object **/
	M_processInstancePtr string
}

/** Xml parser for ConnectingObject **/
type XsdConnectingObject struct {
	XMLName xml.Name	`xml:"connectingObject"`
	/** Instance **/
	M_id	string	`xml:"id,attr"`
	M_bpmnElementId	string	`xml:"bpmnElementId,attr"`


	M_sourceRef	*string	`xml:"sourceRef"`
	M_targetRef	*string	`xml:"targetRef"`
	M_connectingObjectType	string	`xml:"connectingObjectType,attr"`

}
/** UUID **/
func (this *ConnectingObject) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *ConnectingObject) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ConnectingObject) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** BpmnElementId **/
func (this *ConnectingObject) GetBpmnElementId() string{
	return this.M_bpmnElementId
}

/** Init reference BpmnElementId **/
func (this *ConnectingObject) SetBpmnElementId(ref interface{}){
	this.NeedSave = true
	this.M_bpmnElementId = ref.(string)
}

/** Remove reference BpmnElementId **/

/** Participants **/
func (this *ConnectingObject) GetParticipants() []string{
	return this.M_participants
}

/** Init reference Participants **/
func (this *ConnectingObject) SetParticipants(ref interface{}){
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
func (this *ConnectingObject) GetDataRef() []*ItemAwareElementInstance{
	return this.m_dataRef
}

/** Init reference DataRef **/
func (this *ConnectingObject) SetDataRef(ref interface{}){
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
func (this *ConnectingObject) RemoveDataRef(ref interface{}){
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
func (this *ConnectingObject) GetData() []*ItemAwareElementInstance{
	return this.M_data
}

/** Init reference Data **/
func (this *ConnectingObject) SetData(ref interface{}){
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
func (this *ConnectingObject) RemoveData(ref interface{}){
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
func (this *ConnectingObject) GetLogInfoRef() []*LogInfo{
	return this.m_logInfoRef
}

/** Init reference LogInfoRef **/
func (this *ConnectingObject) SetLogInfoRef(ref interface{}){
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
func (this *ConnectingObject) RemoveLogInfoRef(ref interface{}){
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

/** ConnectingObjectType **/
func (this *ConnectingObject) GetConnectingObjectType() ConnectingObjectType{
	return this.M_connectingObjectType
}

/** Init reference ConnectingObjectType **/
func (this *ConnectingObject) SetConnectingObjectType(ref interface{}){
	this.NeedSave = true
	this.M_connectingObjectType = ref.(ConnectingObjectType)
}

/** Remove reference ConnectingObjectType **/

/** SourceRef **/
func (this *ConnectingObject) GetSourceRef() FlowNodeInstance{
	return this.m_sourceRef
}

/** Init reference SourceRef **/
func (this *ConnectingObject) SetSourceRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_sourceRef = ref.(string)
	}else{
		this.m_sourceRef = ref.(FlowNodeInstance)
		this.M_sourceRef = ref.(Instance).GetUUID()
	}
}

/** Remove reference SourceRef **/
func (this *ConnectingObject) RemoveSourceRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_sourceRef.(Instance).GetUUID() {
		this.m_sourceRef = nil
		this.M_sourceRef = ""
	}
}

/** TargetRef **/
func (this *ConnectingObject) GetTargetRef() FlowNodeInstance{
	return this.m_targetRef
}

/** Init reference TargetRef **/
func (this *ConnectingObject) SetTargetRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_targetRef = ref.(string)
	}else{
		this.m_targetRef = ref.(FlowNodeInstance)
		this.M_targetRef = ref.(Instance).GetUUID()
	}
}

/** Remove reference TargetRef **/
func (this *ConnectingObject) RemoveTargetRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_targetRef.(Instance).GetUUID() {
		this.m_targetRef = nil
		this.M_targetRef = ""
	}
}

/** SubprocessInstance **/
func (this *ConnectingObject) GetSubprocessInstancePtr() *SubprocessInstance{
	return this.m_SubprocessInstancePtr
}

/** Init reference SubprocessInstance **/
func (this *ConnectingObject) SetSubprocessInstancePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_SubprocessInstancePtr = ref.(string)
	}else{
		this.m_SubprocessInstancePtr = ref.(*SubprocessInstance)
		this.M_SubprocessInstancePtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference SubprocessInstance **/
func (this *ConnectingObject) RemoveSubprocessInstancePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_SubprocessInstancePtr.GetUUID() {
		this.m_SubprocessInstancePtr = nil
		this.M_SubprocessInstancePtr = ""
	}
}

/** ProcessInstance **/
func (this *ConnectingObject) GetProcessInstancePtr() *ProcessInstance{
	return this.m_processInstancePtr
}

/** Init reference ProcessInstance **/
func (this *ConnectingObject) SetProcessInstancePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_processInstancePtr = ref.(string)
	}else{
		this.m_processInstancePtr = ref.(*ProcessInstance)
		this.M_processInstancePtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference ProcessInstance **/
func (this *ConnectingObject) RemoveProcessInstancePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_processInstancePtr.GetUUID() {
		this.m_processInstancePtr = nil
		this.M_processInstancePtr = ""
	}
}
