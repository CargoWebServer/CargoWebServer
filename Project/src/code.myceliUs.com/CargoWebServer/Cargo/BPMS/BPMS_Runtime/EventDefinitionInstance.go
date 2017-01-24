//+build BPMN
package BPMS_Runtime

import (
	"encoding/xml"
)

type EventDefinitionInstance struct {

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

	/** members of EventDefinitionInstance **/
	M_eventDefinitionType EventDefinitionType

	/** Associations **/
	m_eventInstancePtr *EventInstance
	/** If the ref is a string and not an object **/
	M_eventInstancePtr string
}

/** Xml parser for EventDefinitionInstance **/
type XsdEventDefinitionInstance struct {
	XMLName xml.Name `xml:"eventDefinitionInstance"`
	/** Instance **/
	M_id            string `xml:"id,attr"`
	M_bpmnElementId string `xml:"bpmnElementId,attr"`

	M_eventDefinitionType string `xml:"eventDefinitionType,attr"`
}

/** UUID **/
func (this *EventDefinitionInstance) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *EventDefinitionInstance) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *EventDefinitionInstance) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** BpmnElementId **/
func (this *EventDefinitionInstance) GetBpmnElementId() string {
	return this.M_bpmnElementId
}

/** Init reference BpmnElementId **/
func (this *EventDefinitionInstance) SetBpmnElementId(ref interface{}) {
	this.NeedSave = true
	this.M_bpmnElementId = ref.(string)
}

/** Remove reference BpmnElementId **/

/** Participants **/
func (this *EventDefinitionInstance) GetParticipants() []string {
	return this.M_participants
}

/** Init reference Participants **/
func (this *EventDefinitionInstance) SetParticipants(ref interface{}) {
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
func (this *EventDefinitionInstance) GetDataRef() []*ItemAwareElementInstance {
	return this.m_dataRef
}

/** Init reference DataRef **/
func (this *EventDefinitionInstance) SetDataRef(ref interface{}) {
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
func (this *EventDefinitionInstance) RemoveDataRef(ref interface{}) {
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
func (this *EventDefinitionInstance) GetData() []*ItemAwareElementInstance {
	return this.M_data
}

/** Init reference Data **/
func (this *EventDefinitionInstance) SetData(ref interface{}) {
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
func (this *EventDefinitionInstance) RemoveData(ref interface{}) {
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
func (this *EventDefinitionInstance) GetLogInfoRef() []*LogInfo {
	return this.m_logInfoRef
}

/** Init reference LogInfoRef **/
func (this *EventDefinitionInstance) SetLogInfoRef(ref interface{}) {
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
func (this *EventDefinitionInstance) RemoveLogInfoRef(ref interface{}) {
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

/** EventDefinitionType **/
func (this *EventDefinitionInstance) GetEventDefinitionType() EventDefinitionType {
	return this.M_eventDefinitionType
}

/** Init reference EventDefinitionType **/
func (this *EventDefinitionInstance) SetEventDefinitionType(ref interface{}) {
	this.NeedSave = true
	this.M_eventDefinitionType = ref.(EventDefinitionType)
}

/** Remove reference EventDefinitionType **/

/** EventInstance **/
func (this *EventDefinitionInstance) GetEventInstancePtr() *EventInstance {
	return this.m_eventInstancePtr
}

/** Init reference EventInstance **/
func (this *EventDefinitionInstance) SetEventInstancePtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_eventInstancePtr = ref.(string)
	} else {
		this.m_eventInstancePtr = ref.(*EventInstance)
		this.M_eventInstancePtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference EventInstance **/
func (this *EventDefinitionInstance) RemoveEventInstancePtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_eventInstancePtr.GetUUID() {
		this.m_eventInstancePtr = nil
		this.M_eventInstancePtr = ""
	}
}
