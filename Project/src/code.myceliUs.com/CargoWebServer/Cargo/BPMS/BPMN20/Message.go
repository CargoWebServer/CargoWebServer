// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Message struct {

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

	/** members of RootElement **/
	/** No members **/

	/** members of Message **/
	M_name    string
	m_itemRef *ItemDefinition
	/** If the ref is a string and not an object **/
	M_itemRef string

	/** Associations **/
	m_operationPtr []*Operation
	/** If the ref is a string and not an object **/
	M_operationPtr              []string
	m_messageEventDefinitionPtr []*MessageEventDefinition
	/** If the ref is a string and not an object **/
	M_messageEventDefinitionPtr                 []string
	m_correlationPropertyRetrievalExpressionPtr []*CorrelationPropertyRetrievalExpression
	/** If the ref is a string and not an object **/
	M_correlationPropertyRetrievalExpressionPtr []string
	m_messageFlowPtr                            []*MessageFlow
	/** If the ref is a string and not an object **/
	M_messageFlowPtr []string
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
	M_incomingPtr    []string
	m_definitionsPtr *Definitions
	/** If the ref is a string and not an object **/
	M_definitionsPtr string
}

/** Xml parser for Message **/
type XsdMessage struct {
	XMLName xml.Name `xml:"message"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	/** RootElement **/

	M_name    string `xml:"name,attr"`
	M_itemRef string `xml:"itemRef,attr"`
}

/** UUID **/
func (this *Message) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Message) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Message) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Message) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Message) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Message) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Message) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Message) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Message) SetExtensionDefinitions(ref interface{}) {
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
func (this *Message) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Message) SetExtensionValues(ref interface{}) {
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
func (this *Message) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Message) SetDocumentation(ref interface{}) {
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
func (this *Message) RemoveDocumentation(ref interface{}) {
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
func (this *Message) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Message) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** ItemRef **/
func (this *Message) GetItemRef() *ItemDefinition {
	return this.m_itemRef
}

/** Init reference ItemRef **/
func (this *Message) SetItemRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_itemRef = ref.(string)
	} else {
		this.m_itemRef = ref.(*ItemDefinition)
		this.M_itemRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ItemRef **/
func (this *Message) RemoveItemRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_itemRef.GetUUID() {
		this.m_itemRef = nil
		this.M_itemRef = ""
	}
}

/** Operation **/
func (this *Message) GetOperationPtr() []*Operation {
	return this.m_operationPtr
}

/** Init reference Operation **/
func (this *Message) SetOperationPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_operationPtr); i++ {
			if this.M_operationPtr[i] == refStr {
				return
			}
		}
		this.M_operationPtr = append(this.M_operationPtr, ref.(string))
	} else {
		this.RemoveOperationPtr(ref)
		this.m_operationPtr = append(this.m_operationPtr, ref.(*Operation))
		this.M_operationPtr = append(this.M_operationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Operation **/
func (this *Message) RemoveOperationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	operationPtr_ := make([]*Operation, 0)
	operationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_operationPtr); i++ {
		if toDelete.GetUUID() != this.m_operationPtr[i].GetUUID() {
			operationPtr_ = append(operationPtr_, this.m_operationPtr[i])
			operationPtrUuid = append(operationPtrUuid, this.M_operationPtr[i])
		}
	}
	this.m_operationPtr = operationPtr_
	this.M_operationPtr = operationPtrUuid
}

/** MessageEventDefinition **/
func (this *Message) GetMessageEventDefinitionPtr() []*MessageEventDefinition {
	return this.m_messageEventDefinitionPtr
}

/** Init reference MessageEventDefinition **/
func (this *Message) SetMessageEventDefinitionPtr(ref interface{}) {
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
func (this *Message) RemoveMessageEventDefinitionPtr(ref interface{}) {
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

/** CorrelationPropertyRetrievalExpression **/
func (this *Message) GetCorrelationPropertyRetrievalExpressionPtr() []*CorrelationPropertyRetrievalExpression {
	return this.m_correlationPropertyRetrievalExpressionPtr
}

/** Init reference CorrelationPropertyRetrievalExpression **/
func (this *Message) SetCorrelationPropertyRetrievalExpressionPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_correlationPropertyRetrievalExpressionPtr); i++ {
			if this.M_correlationPropertyRetrievalExpressionPtr[i] == refStr {
				return
			}
		}
		this.M_correlationPropertyRetrievalExpressionPtr = append(this.M_correlationPropertyRetrievalExpressionPtr, ref.(string))
	} else {
		this.RemoveCorrelationPropertyRetrievalExpressionPtr(ref)
		this.m_correlationPropertyRetrievalExpressionPtr = append(this.m_correlationPropertyRetrievalExpressionPtr, ref.(*CorrelationPropertyRetrievalExpression))
		this.M_correlationPropertyRetrievalExpressionPtr = append(this.M_correlationPropertyRetrievalExpressionPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CorrelationPropertyRetrievalExpression **/
func (this *Message) RemoveCorrelationPropertyRetrievalExpressionPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	correlationPropertyRetrievalExpressionPtr_ := make([]*CorrelationPropertyRetrievalExpression, 0)
	correlationPropertyRetrievalExpressionPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_correlationPropertyRetrievalExpressionPtr); i++ {
		if toDelete.GetUUID() != this.m_correlationPropertyRetrievalExpressionPtr[i].GetUUID() {
			correlationPropertyRetrievalExpressionPtr_ = append(correlationPropertyRetrievalExpressionPtr_, this.m_correlationPropertyRetrievalExpressionPtr[i])
			correlationPropertyRetrievalExpressionPtrUuid = append(correlationPropertyRetrievalExpressionPtrUuid, this.M_correlationPropertyRetrievalExpressionPtr[i])
		}
	}
	this.m_correlationPropertyRetrievalExpressionPtr = correlationPropertyRetrievalExpressionPtr_
	this.M_correlationPropertyRetrievalExpressionPtr = correlationPropertyRetrievalExpressionPtrUuid
}

/** MessageFlow **/
func (this *Message) GetMessageFlowPtr() []*MessageFlow {
	return this.m_messageFlowPtr
}

/** Init reference MessageFlow **/
func (this *Message) SetMessageFlowPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_messageFlowPtr); i++ {
			if this.M_messageFlowPtr[i] == refStr {
				return
			}
		}
		this.M_messageFlowPtr = append(this.M_messageFlowPtr, ref.(string))
	} else {
		this.RemoveMessageFlowPtr(ref)
		this.m_messageFlowPtr = append(this.m_messageFlowPtr, ref.(*MessageFlow))
		this.M_messageFlowPtr = append(this.M_messageFlowPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference MessageFlow **/
func (this *Message) RemoveMessageFlowPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	messageFlowPtr_ := make([]*MessageFlow, 0)
	messageFlowPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_messageFlowPtr); i++ {
		if toDelete.GetUUID() != this.m_messageFlowPtr[i].GetUUID() {
			messageFlowPtr_ = append(messageFlowPtr_, this.m_messageFlowPtr[i])
			messageFlowPtrUuid = append(messageFlowPtrUuid, this.M_messageFlowPtr[i])
		}
	}
	this.m_messageFlowPtr = messageFlowPtr_
	this.M_messageFlowPtr = messageFlowPtrUuid
}

/** SendTask **/
func (this *Message) GetSendTaskPtr() []*SendTask {
	return this.m_sendTaskPtr
}

/** Init reference SendTask **/
func (this *Message) SetSendTaskPtr(ref interface{}) {
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
func (this *Message) RemoveSendTaskPtr(ref interface{}) {
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
func (this *Message) GetReceiveTaskPtr() []*ReceiveTask {
	return this.m_receiveTaskPtr
}

/** Init reference ReceiveTask **/
func (this *Message) SetReceiveTaskPtr(ref interface{}) {
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
func (this *Message) RemoveReceiveTaskPtr(ref interface{}) {
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
func (this *Message) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Message) SetLanePtr(ref interface{}) {
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
func (this *Message) RemoveLanePtr(ref interface{}) {
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
func (this *Message) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Message) SetOutgoingPtr(ref interface{}) {
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
func (this *Message) RemoveOutgoingPtr(ref interface{}) {
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
func (this *Message) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Message) SetIncomingPtr(ref interface{}) {
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
func (this *Message) RemoveIncomingPtr(ref interface{}) {
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

/** Definitions **/
func (this *Message) GetDefinitionsPtr() *Definitions {
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *Message) SetDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	} else {
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *Message) RemoveDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}
