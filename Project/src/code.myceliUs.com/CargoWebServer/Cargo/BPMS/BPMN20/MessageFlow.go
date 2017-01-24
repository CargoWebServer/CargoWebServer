// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type MessageFlow struct {

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

	/** members of MessageFlow **/
	M_name      string
	m_sourceRef InteractionNode
	/** If the ref is a string and not an object **/
	M_sourceRef string
	m_targetRef InteractionNode
	/** If the ref is a string and not an object **/
	M_targetRef  string
	m_messageRef *Message
	/** If the ref is a string and not an object **/
	M_messageRef string

	/** Associations **/
	m_communicationPtr []ConversationNode
	/** If the ref is a string and not an object **/
	M_communicationPtr          []string
	m_messageFlowAssociationPtr []*MessageFlowAssociation
	/** If the ref is a string and not an object **/
	M_messageFlowAssociationPtr []string
	m_collaborationPtr          Collaboration
	/** If the ref is a string and not an object **/
	M_collaborationPtr    string
	m_choreographyTaskPtr *ChoreographyTask
	/** If the ref is a string and not an object **/
	M_choreographyTaskPtr string
	m_lanePtr             []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
}

/** Xml parser for MessageFlow **/
type XsdMessageFlow struct {
	XMLName xml.Name `xml:"messageFlow"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_name       string `xml:"name,attr"`
	M_sourceRef  string `xml:"sourceRef,attr"`
	M_targetRef  string `xml:"targetRef,attr"`
	M_messageRef string `xml:"messageRef,attr"`
}

/** UUID **/
func (this *MessageFlow) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *MessageFlow) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *MessageFlow) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *MessageFlow) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *MessageFlow) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *MessageFlow) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *MessageFlow) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *MessageFlow) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *MessageFlow) SetExtensionDefinitions(ref interface{}) {
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
func (this *MessageFlow) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *MessageFlow) SetExtensionValues(ref interface{}) {
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
func (this *MessageFlow) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *MessageFlow) SetDocumentation(ref interface{}) {
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
func (this *MessageFlow) RemoveDocumentation(ref interface{}) {
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
func (this *MessageFlow) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *MessageFlow) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** SourceRef **/
func (this *MessageFlow) GetSourceRef() InteractionNode {
	return this.m_sourceRef
}

/** Init reference SourceRef **/
func (this *MessageFlow) SetSourceRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_sourceRef = ref.(string)
	} else {
		this.m_sourceRef = ref.(InteractionNode)
		this.M_sourceRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference SourceRef **/
func (this *MessageFlow) RemoveSourceRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_sourceRef.(BaseElement).GetUUID() {
		this.m_sourceRef = nil
		this.M_sourceRef = ""
	}
}

/** TargetRef **/
func (this *MessageFlow) GetTargetRef() InteractionNode {
	return this.m_targetRef
}

/** Init reference TargetRef **/
func (this *MessageFlow) SetTargetRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_targetRef = ref.(string)
	} else {
		this.m_targetRef = ref.(InteractionNode)
		this.M_targetRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference TargetRef **/
func (this *MessageFlow) RemoveTargetRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_targetRef.(BaseElement).GetUUID() {
		this.m_targetRef = nil
		this.M_targetRef = ""
	}
}

/** MessageRef **/
func (this *MessageFlow) GetMessageRef() *Message {
	return this.m_messageRef
}

/** Init reference MessageRef **/
func (this *MessageFlow) SetMessageRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_messageRef = ref.(string)
	} else {
		this.m_messageRef = ref.(*Message)
		this.M_messageRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference MessageRef **/
func (this *MessageFlow) RemoveMessageRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_messageRef.GetUUID() {
		this.m_messageRef = nil
		this.M_messageRef = ""
	}
}

/** Communication **/
func (this *MessageFlow) GetCommunicationPtr() []ConversationNode {
	return this.m_communicationPtr
}

/** Init reference Communication **/
func (this *MessageFlow) SetCommunicationPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_communicationPtr); i++ {
			if this.M_communicationPtr[i] == refStr {
				return
			}
		}
		this.M_communicationPtr = append(this.M_communicationPtr, ref.(string))
	} else {
		this.RemoveCommunicationPtr(ref)
		this.m_communicationPtr = append(this.m_communicationPtr, ref.(ConversationNode))
		this.M_communicationPtr = append(this.M_communicationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Communication **/
func (this *MessageFlow) RemoveCommunicationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	communicationPtr_ := make([]ConversationNode, 0)
	communicationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_communicationPtr); i++ {
		if toDelete.GetUUID() != this.m_communicationPtr[i].(BaseElement).GetUUID() {
			communicationPtr_ = append(communicationPtr_, this.m_communicationPtr[i])
			communicationPtrUuid = append(communicationPtrUuid, this.M_communicationPtr[i])
		}
	}
	this.m_communicationPtr = communicationPtr_
	this.M_communicationPtr = communicationPtrUuid
}

/** MessageFlowAssociation **/
func (this *MessageFlow) GetMessageFlowAssociationPtr() []*MessageFlowAssociation {
	return this.m_messageFlowAssociationPtr
}

/** Init reference MessageFlowAssociation **/
func (this *MessageFlow) SetMessageFlowAssociationPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_messageFlowAssociationPtr); i++ {
			if this.M_messageFlowAssociationPtr[i] == refStr {
				return
			}
		}
		this.M_messageFlowAssociationPtr = append(this.M_messageFlowAssociationPtr, ref.(string))
	} else {
		this.RemoveMessageFlowAssociationPtr(ref)
		this.m_messageFlowAssociationPtr = append(this.m_messageFlowAssociationPtr, ref.(*MessageFlowAssociation))
		this.M_messageFlowAssociationPtr = append(this.M_messageFlowAssociationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference MessageFlowAssociation **/
func (this *MessageFlow) RemoveMessageFlowAssociationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	messageFlowAssociationPtr_ := make([]*MessageFlowAssociation, 0)
	messageFlowAssociationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_messageFlowAssociationPtr); i++ {
		if toDelete.GetUUID() != this.m_messageFlowAssociationPtr[i].GetUUID() {
			messageFlowAssociationPtr_ = append(messageFlowAssociationPtr_, this.m_messageFlowAssociationPtr[i])
			messageFlowAssociationPtrUuid = append(messageFlowAssociationPtrUuid, this.M_messageFlowAssociationPtr[i])
		}
	}
	this.m_messageFlowAssociationPtr = messageFlowAssociationPtr_
	this.M_messageFlowAssociationPtr = messageFlowAssociationPtrUuid
}

/** Collaboration **/
func (this *MessageFlow) GetCollaborationPtr() Collaboration {
	return this.m_collaborationPtr
}

/** Init reference Collaboration **/
func (this *MessageFlow) SetCollaborationPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_collaborationPtr = ref.(string)
	} else {
		this.m_collaborationPtr = ref.(Collaboration)
		this.M_collaborationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Collaboration **/
func (this *MessageFlow) RemoveCollaborationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_collaborationPtr.(BaseElement).GetUUID() {
		this.m_collaborationPtr = nil
		this.M_collaborationPtr = ""
	}
}

/** ChoreographyTask **/
func (this *MessageFlow) GetChoreographyTaskPtr() *ChoreographyTask {
	return this.m_choreographyTaskPtr
}

/** Init reference ChoreographyTask **/
func (this *MessageFlow) SetChoreographyTaskPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_choreographyTaskPtr = ref.(string)
	} else {
		this.m_choreographyTaskPtr = ref.(*ChoreographyTask)
		this.M_choreographyTaskPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ChoreographyTask **/
func (this *MessageFlow) RemoveChoreographyTaskPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_choreographyTaskPtr.GetUUID() {
		this.m_choreographyTaskPtr = nil
		this.M_choreographyTaskPtr = ""
	}
}

/** Lane **/
func (this *MessageFlow) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *MessageFlow) SetLanePtr(ref interface{}) {
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
func (this *MessageFlow) RemoveLanePtr(ref interface{}) {
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
func (this *MessageFlow) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *MessageFlow) SetOutgoingPtr(ref interface{}) {
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
func (this *MessageFlow) RemoveOutgoingPtr(ref interface{}) {
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
func (this *MessageFlow) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *MessageFlow) SetIncomingPtr(ref interface{}) {
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
func (this *MessageFlow) RemoveIncomingPtr(ref interface{}) {
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
