// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Operation struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of BaseElement **/
	M_id    string
	m_other interface{}
	/** If the ref is a string and not an object **/
	M_other                string
	M_extensionElements    *ExtensionElements
	M_extensionDefinitions []*ExtensionDefinition
	M_extensionValues      []*ExtensionAttributeValue
	M_documentation        []*Documentation

	/** members of Operation **/
	M_name         string
	m_inMessageRef *Message
	/** If the ref is a string and not an object **/
	M_inMessageRef  string
	m_outMessageRef *Message
	/** If the ref is a string and not an object **/
	M_outMessageRef string
	m_errorRef      []*Error
	/** If the ref is a string and not an object **/
	M_errorRef          []string
	m_implementationRef interface{}
	/** If the ref is a string and not an object **/
	M_implementationRef string

	/** Associations **/
	m_interfacePtr *Interface
	/** If the ref is a string and not an object **/
	M_interfacePtr              string
	m_messageEventDefinitionPtr []*MessageEventDefinition
	/** If the ref is a string and not an object **/
	M_messageEventDefinitionPtr []string
	m_ioBindingPtr              []*InputOutputBinding
	/** If the ref is a string and not an object **/
	M_ioBindingPtr   []string
	m_serviceTaskPtr []*ServiceTask
	/** If the ref is a string and not an object **/
	M_serviceTaskPtr []string
	m_sendTaskPtr    []*SendTask
	/** If the ref is a string and not an object **/
	M_sendTaskPtr    []string
	m_receiveTaskPtr []*ReceiveTask
	/** If the ref is a string and not an object **/
	M_receiveTaskPtr []string
	m_lanePtr        []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
}

/** Xml parser for Operation **/
type XsdOperation struct {
	XMLName xml.Name `xml:"operation"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_inMessageRef      string   `xml:"inMessageRef"`
	M_outMessageRef     *string  `xml:"outMessageRef"`
	M_errorRef          []string `xml:"errorRef"`
	M_name              string   `xml:"name,attr"`
	M_implementationRef string   `xml:"implementationRef,attr"`
}

/** UUID **/
func (this *Operation) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Operation) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Operation) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Operation) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Operation) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Operation) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Operation) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Operation) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Operation) SetExtensionDefinitions(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var extensionDefinitionss []*ExtensionDefinition
	for i := 0; i < len(this.M_extensionDefinitions); i++ {
		if this.M_extensionDefinitions[i].GetName() != ref.(*ExtensionDefinition).GetName() {
			extensionDefinitionss = append(extensionDefinitionss, this.M_extensionDefinitions[i])
		} else {
			isExist = true
			extensionDefinitionss = append(extensionDefinitionss, ref.(*ExtensionDefinition))
		}
	}
	if !isExist {
		extensionDefinitionss = append(extensionDefinitionss, ref.(*ExtensionDefinition))
	}
	this.M_extensionDefinitions = extensionDefinitionss
}

/** Remove reference ExtensionDefinitions **/

/** ExtensionValues **/
func (this *Operation) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Operation) SetExtensionValues(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var extensionValuess []*ExtensionAttributeValue
	for i := 0; i < len(this.M_extensionValues); i++ {
		if this.M_extensionValues[i].GetUUID() != ref.(*ExtensionAttributeValue).GetUUID() {
			extensionValuess = append(extensionValuess, this.M_extensionValues[i])
		} else {
			isExist = true
			extensionValuess = append(extensionValuess, ref.(*ExtensionAttributeValue))
		}
	}
	if !isExist {
		extensionValuess = append(extensionValuess, ref.(*ExtensionAttributeValue))
	}
	this.M_extensionValues = extensionValuess
}

/** Remove reference ExtensionValues **/

/** Documentation **/
func (this *Operation) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Operation) SetDocumentation(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var documentations []*Documentation
	for i := 0; i < len(this.M_documentation); i++ {
		if this.M_documentation[i].GetUUID() != ref.(BaseElement).GetUUID() {
			documentations = append(documentations, this.M_documentation[i])
		} else {
			isExist = true
			documentations = append(documentations, ref.(*Documentation))
		}
	}
	if !isExist {
		documentations = append(documentations, ref.(*Documentation))
	}
	this.M_documentation = documentations
}

/** Remove reference Documentation **/
func (this *Operation) RemoveDocumentation(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	documentation_ := make([]*Documentation, 0)
	for i := 0; i < len(this.M_documentation); i++ {
		if toDelete.GetUUID() != this.M_documentation[i].GetUUID() {
			documentation_ = append(documentation_, this.M_documentation[i])
		}
	}
	this.M_documentation = documentation_
}

/** Name **/
func (this *Operation) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Operation) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** InMessageRef **/
func (this *Operation) GetInMessageRef() *Message {
	return this.m_inMessageRef
}

/** Init reference InMessageRef **/
func (this *Operation) SetInMessageRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_inMessageRef = ref.(string)
	} else {
		this.m_inMessageRef = ref.(*Message)
		this.M_inMessageRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference InMessageRef **/
func (this *Operation) RemoveInMessageRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_inMessageRef.GetUUID() {
		this.m_inMessageRef = nil
		this.M_inMessageRef = ""
	}
}

/** OutMessageRef **/
func (this *Operation) GetOutMessageRef() *Message {
	return this.m_outMessageRef
}

/** Init reference OutMessageRef **/
func (this *Operation) SetOutMessageRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_outMessageRef = ref.(string)
	} else {
		this.m_outMessageRef = ref.(*Message)
		this.M_outMessageRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference OutMessageRef **/
func (this *Operation) RemoveOutMessageRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_outMessageRef.GetUUID() {
		this.m_outMessageRef = nil
		this.M_outMessageRef = ""
	}
}

/** ErrorRef **/
func (this *Operation) GetErrorRef() []*Error {
	return this.m_errorRef
}

/** Init reference ErrorRef **/
func (this *Operation) SetErrorRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_errorRef); i++ {
			if this.M_errorRef[i] == refStr {
				return
			}
		}
		this.M_errorRef = append(this.M_errorRef, ref.(string))
	} else {
		this.RemoveErrorRef(ref)
		this.m_errorRef = append(this.m_errorRef, ref.(*Error))
		this.M_errorRef = append(this.M_errorRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ErrorRef **/
func (this *Operation) RemoveErrorRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	errorRef_ := make([]*Error, 0)
	errorRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_errorRef); i++ {
		if toDelete.GetUUID() != this.m_errorRef[i].GetUUID() {
			errorRef_ = append(errorRef_, this.m_errorRef[i])
			errorRefUuid = append(errorRefUuid, this.M_errorRef[i])
		}
	}
	this.m_errorRef = errorRef_
	this.M_errorRef = errorRefUuid
}

/** ImplementationRef **/
func (this *Operation) GetImplementationRef() interface{} {
	return this.m_implementationRef
}

/** Init reference ImplementationRef **/
func (this *Operation) SetImplementationRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_implementationRef = ref.(string)
	} else {
		this.m_implementationRef = ref.(interface{})
	}
}

/** Remove reference ImplementationRef **/

/** Interface **/
func (this *Operation) GetInterfacePtr() *Interface {
	return this.m_interfacePtr
}

/** Init reference Interface **/
func (this *Operation) SetInterfacePtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_interfacePtr = ref.(string)
	} else {
		this.m_interfacePtr = ref.(*Interface)
		this.M_interfacePtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Interface **/
func (this *Operation) RemoveInterfacePtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_interfacePtr.GetUUID() {
		this.m_interfacePtr = nil
		this.M_interfacePtr = ""
	}
}

/** MessageEventDefinition **/
func (this *Operation) GetMessageEventDefinitionPtr() []*MessageEventDefinition {
	return this.m_messageEventDefinitionPtr
}

/** Init reference MessageEventDefinition **/
func (this *Operation) SetMessageEventDefinitionPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_messageEventDefinitionPtr); i++ {
			if this.M_messageEventDefinitionPtr[i] == refStr {
				return
			}
		}
		this.M_messageEventDefinitionPtr = append(this.M_messageEventDefinitionPtr, ref.(string))
	} else {
		this.RemoveMessageEventDefinitionPtr(ref)
		this.m_messageEventDefinitionPtr = append(this.m_messageEventDefinitionPtr, ref.(*MessageEventDefinition))
		this.M_messageEventDefinitionPtr = append(this.M_messageEventDefinitionPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference MessageEventDefinition **/
func (this *Operation) RemoveMessageEventDefinitionPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	messageEventDefinitionPtr_ := make([]*MessageEventDefinition, 0)
	messageEventDefinitionPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_messageEventDefinitionPtr); i++ {
		if toDelete.GetUUID() != this.m_messageEventDefinitionPtr[i].GetUUID() {
			messageEventDefinitionPtr_ = append(messageEventDefinitionPtr_, this.m_messageEventDefinitionPtr[i])
			messageEventDefinitionPtrUuid = append(messageEventDefinitionPtrUuid, this.M_messageEventDefinitionPtr[i])
		}
	}
	this.m_messageEventDefinitionPtr = messageEventDefinitionPtr_
	this.M_messageEventDefinitionPtr = messageEventDefinitionPtrUuid
}

/** IoBinding **/
func (this *Operation) GetIoBindingPtr() []*InputOutputBinding {
	return this.m_ioBindingPtr
}

/** Init reference IoBinding **/
func (this *Operation) SetIoBindingPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_ioBindingPtr); i++ {
			if this.M_ioBindingPtr[i] == refStr {
				return
			}
		}
		this.M_ioBindingPtr = append(this.M_ioBindingPtr, ref.(string))
	} else {
		this.RemoveIoBindingPtr(ref)
		this.m_ioBindingPtr = append(this.m_ioBindingPtr, ref.(*InputOutputBinding))
		this.M_ioBindingPtr = append(this.M_ioBindingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference IoBinding **/
func (this *Operation) RemoveIoBindingPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	ioBindingPtr_ := make([]*InputOutputBinding, 0)
	ioBindingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_ioBindingPtr); i++ {
		if toDelete.GetUUID() != this.m_ioBindingPtr[i].GetUUID() {
			ioBindingPtr_ = append(ioBindingPtr_, this.m_ioBindingPtr[i])
			ioBindingPtrUuid = append(ioBindingPtrUuid, this.M_ioBindingPtr[i])
		}
	}
	this.m_ioBindingPtr = ioBindingPtr_
	this.M_ioBindingPtr = ioBindingPtrUuid
}

/** ServiceTask **/
func (this *Operation) GetServiceTaskPtr() []*ServiceTask {
	return this.m_serviceTaskPtr
}

/** Init reference ServiceTask **/
func (this *Operation) SetServiceTaskPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_serviceTaskPtr); i++ {
			if this.M_serviceTaskPtr[i] == refStr {
				return
			}
		}
		this.M_serviceTaskPtr = append(this.M_serviceTaskPtr, ref.(string))
	} else {
		this.RemoveServiceTaskPtr(ref)
		this.m_serviceTaskPtr = append(this.m_serviceTaskPtr, ref.(*ServiceTask))
		this.M_serviceTaskPtr = append(this.M_serviceTaskPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ServiceTask **/
func (this *Operation) RemoveServiceTaskPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	serviceTaskPtr_ := make([]*ServiceTask, 0)
	serviceTaskPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_serviceTaskPtr); i++ {
		if toDelete.GetUUID() != this.m_serviceTaskPtr[i].GetUUID() {
			serviceTaskPtr_ = append(serviceTaskPtr_, this.m_serviceTaskPtr[i])
			serviceTaskPtrUuid = append(serviceTaskPtrUuid, this.M_serviceTaskPtr[i])
		}
	}
	this.m_serviceTaskPtr = serviceTaskPtr_
	this.M_serviceTaskPtr = serviceTaskPtrUuid
}

/** SendTask **/
func (this *Operation) GetSendTaskPtr() []*SendTask {
	return this.m_sendTaskPtr
}

/** Init reference SendTask **/
func (this *Operation) SetSendTaskPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_sendTaskPtr); i++ {
			if this.M_sendTaskPtr[i] == refStr {
				return
			}
		}
		this.M_sendTaskPtr = append(this.M_sendTaskPtr, ref.(string))
	} else {
		this.RemoveSendTaskPtr(ref)
		this.m_sendTaskPtr = append(this.m_sendTaskPtr, ref.(*SendTask))
		this.M_sendTaskPtr = append(this.M_sendTaskPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference SendTask **/
func (this *Operation) RemoveSendTaskPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	sendTaskPtr_ := make([]*SendTask, 0)
	sendTaskPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_sendTaskPtr); i++ {
		if toDelete.GetUUID() != this.m_sendTaskPtr[i].GetUUID() {
			sendTaskPtr_ = append(sendTaskPtr_, this.m_sendTaskPtr[i])
			sendTaskPtrUuid = append(sendTaskPtrUuid, this.M_sendTaskPtr[i])
		}
	}
	this.m_sendTaskPtr = sendTaskPtr_
	this.M_sendTaskPtr = sendTaskPtrUuid
}

/** ReceiveTask **/
func (this *Operation) GetReceiveTaskPtr() []*ReceiveTask {
	return this.m_receiveTaskPtr
}

/** Init reference ReceiveTask **/
func (this *Operation) SetReceiveTaskPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_receiveTaskPtr); i++ {
			if this.M_receiveTaskPtr[i] == refStr {
				return
			}
		}
		this.M_receiveTaskPtr = append(this.M_receiveTaskPtr, ref.(string))
	} else {
		this.RemoveReceiveTaskPtr(ref)
		this.m_receiveTaskPtr = append(this.m_receiveTaskPtr, ref.(*ReceiveTask))
		this.M_receiveTaskPtr = append(this.M_receiveTaskPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ReceiveTask **/
func (this *Operation) RemoveReceiveTaskPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	receiveTaskPtr_ := make([]*ReceiveTask, 0)
	receiveTaskPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_receiveTaskPtr); i++ {
		if toDelete.GetUUID() != this.m_receiveTaskPtr[i].GetUUID() {
			receiveTaskPtr_ = append(receiveTaskPtr_, this.m_receiveTaskPtr[i])
			receiveTaskPtrUuid = append(receiveTaskPtrUuid, this.M_receiveTaskPtr[i])
		}
	}
	this.m_receiveTaskPtr = receiveTaskPtr_
	this.M_receiveTaskPtr = receiveTaskPtrUuid
}

/** Lane **/
func (this *Operation) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Operation) SetLanePtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_lanePtr); i++ {
			if this.M_lanePtr[i] == refStr {
				return
			}
		}
		this.M_lanePtr = append(this.M_lanePtr, ref.(string))
	} else {
		this.RemoveLanePtr(ref)
		this.m_lanePtr = append(this.m_lanePtr, ref.(*Lane))
		this.M_lanePtr = append(this.M_lanePtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Lane **/
func (this *Operation) RemoveLanePtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	lanePtr_ := make([]*Lane, 0)
	lanePtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_lanePtr); i++ {
		if toDelete.GetUUID() != this.m_lanePtr[i].GetUUID() {
			lanePtr_ = append(lanePtr_, this.m_lanePtr[i])
			lanePtrUuid = append(lanePtrUuid, this.M_lanePtr[i])
		}
	}
	this.m_lanePtr = lanePtr_
	this.M_lanePtr = lanePtrUuid
}

/** Outgoing **/
func (this *Operation) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Operation) SetOutgoingPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_outgoingPtr); i++ {
			if this.M_outgoingPtr[i] == refStr {
				return
			}
		}
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(string))
	} else {
		this.RemoveOutgoingPtr(ref)
		this.m_outgoingPtr = append(this.m_outgoingPtr, ref.(*Association))
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Outgoing **/
func (this *Operation) RemoveOutgoingPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	outgoingPtr_ := make([]*Association, 0)
	outgoingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_outgoingPtr); i++ {
		if toDelete.GetUUID() != this.m_outgoingPtr[i].GetUUID() {
			outgoingPtr_ = append(outgoingPtr_, this.m_outgoingPtr[i])
			outgoingPtrUuid = append(outgoingPtrUuid, this.M_outgoingPtr[i])
		}
	}
	this.m_outgoingPtr = outgoingPtr_
	this.M_outgoingPtr = outgoingPtrUuid
}

/** Incoming **/
func (this *Operation) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Operation) SetIncomingPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_incomingPtr); i++ {
			if this.M_incomingPtr[i] == refStr {
				return
			}
		}
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(string))
	} else {
		this.RemoveIncomingPtr(ref)
		this.m_incomingPtr = append(this.m_incomingPtr, ref.(*Association))
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Incoming **/
func (this *Operation) RemoveIncomingPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	incomingPtr_ := make([]*Association, 0)
	incomingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_incomingPtr); i++ {
		if toDelete.GetUUID() != this.m_incomingPtr[i].GetUUID() {
			incomingPtr_ = append(incomingPtr_, this.m_incomingPtr[i])
			incomingPtrUuid = append(incomingPtrUuid, this.M_incomingPtr[i])
		}
	}
	this.m_incomingPtr = incomingPtr_
	this.M_incomingPtr = incomingPtrUuid
}
