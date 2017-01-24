//+build BPMN
package BPMS_Runtime

import (
	"encoding/xml"
)

type ProcessInstance struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of Instance **/
	M_id            string
	M_bpmnElementId string
	M_participants  []string
	m_dataRef       []*ItemAwareElementInstance
	/** If the ref is a string and not an object **/
	M_dataRef    []string
	M_data       []*ItemAwareElementInstance
	m_logInfoRef []*LogInfo
	/** If the ref is a string and not an object **/
	M_logInfoRef []string

	/** members of ProcessInstance **/
	M_number            int
	M_colorName         string
	M_colorNumber       string
	M_flowNodeInstances []FlowNodeInstance
	M_connectingObjects []*ConnectingObject

	/** Associations **/
	m_parentPtr *DefinitionsInstance
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for ProcessInstance **/
type XsdProcessInstance struct {
	XMLName xml.Name `xml:"process"`
	/** Instance **/
	M_id            string `xml:"id,attr"`
	M_bpmnElementId string `xml:"bpmnElementId,attr"`

	M_flowNodeInstances_0 []*XsdActivityInstance   `xml:"activityInstance,omitempty"`
	M_flowNodeInstances_1 []*XsdSubprocessInstance `xml:"SubprocessInstance,omitempty"`
	M_flowNodeInstances_2 []*XsdGatewayInstance    `xml:"gatewayInstance,omitempty"`
	M_flowNodeInstances_3 []*XsdEventInstance      `xml:"eventInstance,omitempty"`

	M_number      int    `xml:"number,attr"`
	M_colorNumber string `xml:"colorNumber,attr"`
	M_colorName   string `xml:"colorName,attr"`
}

/** UUID **/
func (this *ProcessInstance) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *ProcessInstance) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *ProcessInstance) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** BpmnElementId **/
func (this *ProcessInstance) GetBpmnElementId() string {
	return this.M_bpmnElementId
}

/** Init reference BpmnElementId **/
func (this *ProcessInstance) SetBpmnElementId(ref interface{}) {
	this.NeedSave = true
	this.M_bpmnElementId = ref.(string)
}

/** Remove reference BpmnElementId **/

/** Participants **/
func (this *ProcessInstance) GetParticipants() []string {
	return this.M_participants
}

/** Init reference Participants **/
func (this *ProcessInstance) SetParticipants(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var participantss []string
	for i := 0; i < len(this.M_participants); i++ {
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
func (this *ProcessInstance) GetDataRef() []*ItemAwareElementInstance {
	return this.m_dataRef
}

/** Init reference DataRef **/
func (this *ProcessInstance) SetDataRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_dataRef); i++ {
			if this.M_dataRef[i] == refStr {
				return
			}
		}
		this.M_dataRef = append(this.M_dataRef, ref.(string))
	} else {
		this.RemoveDataRef(ref)
		this.m_dataRef = append(this.m_dataRef, ref.(*ItemAwareElementInstance))
		this.M_dataRef = append(this.M_dataRef, ref.(*ItemAwareElementInstance).GetUUID())
	}
}

/** Remove reference DataRef **/
func (this *ProcessInstance) RemoveDataRef(ref interface{}) {
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
func (this *ProcessInstance) GetData() []*ItemAwareElementInstance {
	return this.M_data
}

/** Init reference Data **/
func (this *ProcessInstance) SetData(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var datas []*ItemAwareElementInstance
	for i := 0; i < len(this.M_data); i++ {
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
func (this *ProcessInstance) RemoveData(ref interface{}) {
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
func (this *ProcessInstance) GetLogInfoRef() []*LogInfo {
	return this.m_logInfoRef
}

/** Init reference LogInfoRef **/
func (this *ProcessInstance) SetLogInfoRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_logInfoRef); i++ {
			if this.M_logInfoRef[i] == refStr {
				return
			}
		}
		this.M_logInfoRef = append(this.M_logInfoRef, ref.(string))
	} else {
		this.RemoveLogInfoRef(ref)
		this.m_logInfoRef = append(this.m_logInfoRef, ref.(*LogInfo))
		this.M_logInfoRef = append(this.M_logInfoRef, ref.(*LogInfo).GetUUID())
	}
}

/** Remove reference LogInfoRef **/
func (this *ProcessInstance) RemoveLogInfoRef(ref interface{}) {
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

/** Number **/
func (this *ProcessInstance) GetNumber() int {
	return this.M_number
}

/** Init reference Number **/
func (this *ProcessInstance) SetNumber(ref interface{}) {
	this.NeedSave = true
	this.M_number = ref.(int)
}

/** Remove reference Number **/

/** ColorName **/
func (this *ProcessInstance) GetColorName() string {
	return this.M_colorName
}

/** Init reference ColorName **/
func (this *ProcessInstance) SetColorName(ref interface{}) {
	this.NeedSave = true
	this.M_colorName = ref.(string)
}

/** Remove reference ColorName **/

/** ColorNumber **/
func (this *ProcessInstance) GetColorNumber() string {
	return this.M_colorNumber
}

/** Init reference ColorNumber **/
func (this *ProcessInstance) SetColorNumber(ref interface{}) {
	this.NeedSave = true
	this.M_colorNumber = ref.(string)
}

/** Remove reference ColorNumber **/

/** FlowNodeInstances **/
func (this *ProcessInstance) GetFlowNodeInstances() []FlowNodeInstance {
	return this.M_flowNodeInstances
}

/** Init reference FlowNodeInstances **/
func (this *ProcessInstance) SetFlowNodeInstances(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var flowNodeInstancess []FlowNodeInstance
	for i := 0; i < len(this.M_flowNodeInstances); i++ {
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
func (this *ProcessInstance) RemoveFlowNodeInstances(ref interface{}) {
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
func (this *ProcessInstance) GetConnectingObjects() []*ConnectingObject {
	return this.M_connectingObjects
}

/** Init reference ConnectingObjects **/
func (this *ProcessInstance) SetConnectingObjects(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var connectingObjectss []*ConnectingObject
	for i := 0; i < len(this.M_connectingObjects); i++ {
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
func (this *ProcessInstance) RemoveConnectingObjects(ref interface{}) {
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

/** Parent **/
func (this *ProcessInstance) GetParentPtr() *DefinitionsInstance {
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *ProcessInstance) SetParentPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	} else {
		this.m_parentPtr = ref.(*DefinitionsInstance)
		this.M_parentPtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *ProcessInstance) RemoveParentPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
		this.m_parentPtr = nil
		this.M_parentPtr = ""
	}
}
