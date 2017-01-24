// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Conversation struct {

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

	/** members of InteractionNode **/
	m_incomingConversationLinks []*ConversationLink
	/** If the ref is a string and not an object **/
	M_incomingConversationLinks []string
	m_outgoingConversationLinks []*ConversationLink
	/** If the ref is a string and not an object **/
	M_outgoingConversationLinks []string

	/** members of ConversationNode **/
	M_name           string
	m_participantRef []*Participant
	/** If the ref is a string and not an object **/
	M_participantRef []string
	m_messageFlowRef []*MessageFlow
	/** If the ref is a string and not an object **/
	M_messageFlowRef []string
	M_correlationKey []*CorrelationKey

	/** members of Conversation **/
	/** No members **/

	/** Associations **/
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr    []string
	m_messageFlowPtr []*MessageFlow
	/** If the ref is a string and not an object **/
	M_messageFlowPtr             []string
	m_conversationAssociationPtr []*ConversationAssociation
	/** If the ref is a string and not an object **/
	M_conversationAssociationPtr []string
	m_subConversationPtr         *SubConversation
	/** If the ref is a string and not an object **/
	M_subConversationPtr string
	m_collaborationPtr   Collaboration
	/** If the ref is a string and not an object **/
	M_collaborationPtr string
}

/** Xml parser for Conversation **/
type XsdConversation struct {
	XMLName xml.Name `xml:"conversation"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	/** ConversationNode **/
	M_participantRef []string             `xml:"participantRef"`
	M_messageFlowRef []string             `xml:"messageFlowRef"`
	M_correlationKey []*XsdCorrelationKey `xml:"correlationKey,omitempty"`
	M_name           string               `xml:"name,attr"`
}

/** UUID **/
func (this *Conversation) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Conversation) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Conversation) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Conversation) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Conversation) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Conversation) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Conversation) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Conversation) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Conversation) SetExtensionDefinitions(ref interface{}) {
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
func (this *Conversation) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Conversation) SetExtensionValues(ref interface{}) {
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
func (this *Conversation) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Conversation) SetDocumentation(ref interface{}) {
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
func (this *Conversation) RemoveDocumentation(ref interface{}) {
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

/** IncomingConversationLinks **/
func (this *Conversation) GetIncomingConversationLinks() []*ConversationLink {
	return this.m_incomingConversationLinks
}

/** Init reference IncomingConversationLinks **/
func (this *Conversation) SetIncomingConversationLinks(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_incomingConversationLinks); i++ {
			if this.M_incomingConversationLinks[i] == refStr {
				return
			}
		}
		this.M_incomingConversationLinks = append(this.M_incomingConversationLinks, ref.(string))
	} else {
		this.RemoveIncomingConversationLinks(ref)
		this.m_incomingConversationLinks = append(this.m_incomingConversationLinks, ref.(*ConversationLink))
		this.M_incomingConversationLinks = append(this.M_incomingConversationLinks, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference IncomingConversationLinks **/
func (this *Conversation) RemoveIncomingConversationLinks(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	incomingConversationLinks_ := make([]*ConversationLink, 0)
	incomingConversationLinksUuid := make([]string, 0)
	for i := 0; i < len(this.m_incomingConversationLinks); i++ {
		if toDelete.GetUUID() != this.m_incomingConversationLinks[i].GetUUID() {
			incomingConversationLinks_ = append(incomingConversationLinks_, this.m_incomingConversationLinks[i])
			incomingConversationLinksUuid = append(incomingConversationLinksUuid, this.M_incomingConversationLinks[i])
		}
	}
	this.m_incomingConversationLinks = incomingConversationLinks_
	this.M_incomingConversationLinks = incomingConversationLinksUuid
}

/** OutgoingConversationLinks **/
func (this *Conversation) GetOutgoingConversationLinks() []*ConversationLink {
	return this.m_outgoingConversationLinks
}

/** Init reference OutgoingConversationLinks **/
func (this *Conversation) SetOutgoingConversationLinks(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_outgoingConversationLinks); i++ {
			if this.M_outgoingConversationLinks[i] == refStr {
				return
			}
		}
		this.M_outgoingConversationLinks = append(this.M_outgoingConversationLinks, ref.(string))
	} else {
		this.RemoveOutgoingConversationLinks(ref)
		this.m_outgoingConversationLinks = append(this.m_outgoingConversationLinks, ref.(*ConversationLink))
		this.M_outgoingConversationLinks = append(this.M_outgoingConversationLinks, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference OutgoingConversationLinks **/
func (this *Conversation) RemoveOutgoingConversationLinks(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	outgoingConversationLinks_ := make([]*ConversationLink, 0)
	outgoingConversationLinksUuid := make([]string, 0)
	for i := 0; i < len(this.m_outgoingConversationLinks); i++ {
		if toDelete.GetUUID() != this.m_outgoingConversationLinks[i].GetUUID() {
			outgoingConversationLinks_ = append(outgoingConversationLinks_, this.m_outgoingConversationLinks[i])
			outgoingConversationLinksUuid = append(outgoingConversationLinksUuid, this.M_outgoingConversationLinks[i])
		}
	}
	this.m_outgoingConversationLinks = outgoingConversationLinks_
	this.M_outgoingConversationLinks = outgoingConversationLinksUuid
}

/** Name **/
func (this *Conversation) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Conversation) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** ParticipantRef **/
func (this *Conversation) GetParticipantRef() []*Participant {
	return this.m_participantRef
}

/** Init reference ParticipantRef **/
func (this *Conversation) SetParticipantRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_participantRef); i++ {
			if this.M_participantRef[i] == refStr {
				return
			}
		}
		this.M_participantRef = append(this.M_participantRef, ref.(string))
	} else {
		this.RemoveParticipantRef(ref)
		this.m_participantRef = append(this.m_participantRef, ref.(*Participant))
		this.M_participantRef = append(this.M_participantRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ParticipantRef **/
func (this *Conversation) RemoveParticipantRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	participantRef_ := make([]*Participant, 0)
	participantRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_participantRef); i++ {
		if toDelete.GetUUID() != this.m_participantRef[i].GetUUID() {
			participantRef_ = append(participantRef_, this.m_participantRef[i])
			participantRefUuid = append(participantRefUuid, this.M_participantRef[i])
		}
	}
	this.m_participantRef = participantRef_
	this.M_participantRef = participantRefUuid
}

/** MessageFlowRef **/
func (this *Conversation) GetMessageFlowRef() []*MessageFlow {
	return this.m_messageFlowRef
}

/** Init reference MessageFlowRef **/
func (this *Conversation) SetMessageFlowRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_messageFlowRef); i++ {
			if this.M_messageFlowRef[i] == refStr {
				return
			}
		}
		this.M_messageFlowRef = append(this.M_messageFlowRef, ref.(string))
	} else {
		this.RemoveMessageFlowRef(ref)
		this.m_messageFlowRef = append(this.m_messageFlowRef, ref.(*MessageFlow))
		this.M_messageFlowRef = append(this.M_messageFlowRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference MessageFlowRef **/
func (this *Conversation) RemoveMessageFlowRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	messageFlowRef_ := make([]*MessageFlow, 0)
	messageFlowRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_messageFlowRef); i++ {
		if toDelete.GetUUID() != this.m_messageFlowRef[i].GetUUID() {
			messageFlowRef_ = append(messageFlowRef_, this.m_messageFlowRef[i])
			messageFlowRefUuid = append(messageFlowRefUuid, this.M_messageFlowRef[i])
		}
	}
	this.m_messageFlowRef = messageFlowRef_
	this.M_messageFlowRef = messageFlowRefUuid
}

/** CorrelationKey **/
func (this *Conversation) GetCorrelationKey() []*CorrelationKey {
	return this.M_correlationKey
}

/** Init reference CorrelationKey **/
func (this *Conversation) SetCorrelationKey(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var correlationKeys []*CorrelationKey
	for i := 0; i < len(this.M_correlationKey); i++ {
		if this.M_correlationKey[i].GetUUID() != ref.(BaseElement).GetUUID() {
			correlationKeys = append(correlationKeys, this.M_correlationKey[i])
		} else {
			isExist = true
			correlationKeys = append(correlationKeys, ref.(*CorrelationKey))
		}
	}
	if !isExist {
		correlationKeys = append(correlationKeys, ref.(*CorrelationKey))
	}
	this.M_correlationKey = correlationKeys
}

/** Remove reference CorrelationKey **/
func (this *Conversation) RemoveCorrelationKey(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	correlationKey_ := make([]*CorrelationKey, 0)
	for i := 0; i < len(this.M_correlationKey); i++ {
		if toDelete.GetUUID() != this.M_correlationKey[i].GetUUID() {
			correlationKey_ = append(correlationKey_, this.M_correlationKey[i])
		}
	}
	this.M_correlationKey = correlationKey_
}

/** Lane **/
func (this *Conversation) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Conversation) SetLanePtr(ref interface{}) {
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
func (this *Conversation) RemoveLanePtr(ref interface{}) {
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
func (this *Conversation) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Conversation) SetOutgoingPtr(ref interface{}) {
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
func (this *Conversation) RemoveOutgoingPtr(ref interface{}) {
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
func (this *Conversation) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Conversation) SetIncomingPtr(ref interface{}) {
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
func (this *Conversation) RemoveIncomingPtr(ref interface{}) {
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

/** MessageFlow **/
func (this *Conversation) GetMessageFlowPtr() []*MessageFlow {
	return this.m_messageFlowPtr
}

/** Init reference MessageFlow **/
func (this *Conversation) SetMessageFlowPtr(ref interface{}) {
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
func (this *Conversation) RemoveMessageFlowPtr(ref interface{}) {
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

/** ConversationAssociation **/
func (this *Conversation) GetConversationAssociationPtr() []*ConversationAssociation {
	return this.m_conversationAssociationPtr
}

/** Init reference ConversationAssociation **/
func (this *Conversation) SetConversationAssociationPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_conversationAssociationPtr); i++ {
			if this.M_conversationAssociationPtr[i] == refStr {
				return
			}
		}
		this.M_conversationAssociationPtr = append(this.M_conversationAssociationPtr, ref.(string))
	} else {
		this.RemoveConversationAssociationPtr(ref)
		this.m_conversationAssociationPtr = append(this.m_conversationAssociationPtr, ref.(*ConversationAssociation))
		this.M_conversationAssociationPtr = append(this.M_conversationAssociationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ConversationAssociation **/
func (this *Conversation) RemoveConversationAssociationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	conversationAssociationPtr_ := make([]*ConversationAssociation, 0)
	conversationAssociationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_conversationAssociationPtr); i++ {
		if toDelete.GetUUID() != this.m_conversationAssociationPtr[i].GetUUID() {
			conversationAssociationPtr_ = append(conversationAssociationPtr_, this.m_conversationAssociationPtr[i])
			conversationAssociationPtrUuid = append(conversationAssociationPtrUuid, this.M_conversationAssociationPtr[i])
		}
	}
	this.m_conversationAssociationPtr = conversationAssociationPtr_
	this.M_conversationAssociationPtr = conversationAssociationPtrUuid
}

/** SubConversation **/
func (this *Conversation) GetSubConversationPtr() *SubConversation {
	return this.m_subConversationPtr
}

/** Init reference SubConversation **/
func (this *Conversation) SetSubConversationPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_subConversationPtr = ref.(string)
	} else {
		this.m_subConversationPtr = ref.(*SubConversation)
		this.M_subConversationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference SubConversation **/
func (this *Conversation) RemoveSubConversationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_subConversationPtr.GetUUID() {
		this.m_subConversationPtr = nil
		this.M_subConversationPtr = ""
	}
}

/** Collaboration **/
func (this *Conversation) GetCollaborationPtr() Collaboration {
	return this.m_collaborationPtr
}

/** Init reference Collaboration **/
func (this *Conversation) SetCollaborationPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_collaborationPtr = ref.(string)
	} else {
		this.m_collaborationPtr = ref.(Collaboration)
		this.M_collaborationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Collaboration **/
func (this *Conversation) RemoveCollaborationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_collaborationPtr.(BaseElement).GetUUID() {
		this.m_collaborationPtr = nil
		this.M_collaborationPtr = ""
	}
}
