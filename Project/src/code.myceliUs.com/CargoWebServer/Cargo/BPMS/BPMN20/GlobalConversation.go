// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type GlobalConversation struct {

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

	/** members of Collaboration **/
	M_name            string
	M_isClosed        bool
	m_choreographyRef []Choreography
	/** If the ref is a string and not an object **/
	M_choreographyRef         []string
	M_artifact                []Artifact
	M_participantAssociation  []*ParticipantAssociation
	M_messageFlowAssociation  []*MessageFlowAssociation
	M_conversationAssociation []*ConversationAssociation
	M_participant             []*Participant
	M_messageFlow             []*MessageFlow
	M_correlationKey          []*CorrelationKey
	M_conversationNode        []ConversationNode
	M_conversationLink        []*ConversationLink

	/** members of GlobalConversation **/
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
	m_definitionsPtr *Definitions
	/** If the ref is a string and not an object **/
	M_definitionsPtr string
	m_processPtr     []*Process
	/** If the ref is a string and not an object **/
	M_processPtr          []string
	m_callConversationPtr []*CallConversation
	/** If the ref is a string and not an object **/
	M_callConversationPtr []string
}

/** Xml parser for GlobalConversation **/
type XsdGlobalConversation struct {
	XMLName xml.Name `xml:"globalConversation"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	/** RootElement **/

	/** Collaboration **/
	M_participant []*XsdParticipant    `xml:"participant,omitempty"`
	M_messageFlow []*XsdMessageFlow    `xml:"messageFlow,omitempty"`
	M_artifact_0  []*XsdAssociation    `xml:"association,omitempty"`
	M_artifact_1  []*XsdGroup          `xml:"group,omitempty"`
	M_artifact_2  []*XsdTextAnnotation `xml:"textAnnotation,omitempty"`

	M_conversationNode_0 []*XsdCallConversation `xml:"callConversation,omitempty"`
	M_conversationNode_1 []*XsdConversation     `xml:"conversation,omitempty"`
	M_conversationNode_2 []*XsdSubConversation  `xml:"subConversation,omitempty"`

	M_conversationAssociation []*XsdConversationAssociation `xml:"conversationAssociation,omitempty"`
	M_participantAssociation  []*XsdParticipantAssociation  `xml:"participantAssociation,omitempty"`
	M_messageFlowAssociation  []*XsdMessageFlowAssociation  `xml:"messageFlowAssociation,omitempty"`
	M_correlationKey          []*XsdCorrelationKey          `xml:"correlationKey,omitempty"`
	M_choreographyRef         []string                      `xml:"choreographyRef"`
	M_conversationLink        []*XsdConversationLink        `xml:"conversationLink,omitempty"`
	M_name                    string                        `xml:"name,attr"`
	M_isClosed                bool                          `xml:"isClosed,attr"`
}

/** UUID **/
func (this *GlobalConversation) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *GlobalConversation) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *GlobalConversation) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *GlobalConversation) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *GlobalConversation) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *GlobalConversation) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *GlobalConversation) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *GlobalConversation) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *GlobalConversation) SetExtensionDefinitions(ref interface{}) {
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
func (this *GlobalConversation) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *GlobalConversation) SetExtensionValues(ref interface{}) {
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
func (this *GlobalConversation) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *GlobalConversation) SetDocumentation(ref interface{}) {
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
func (this *GlobalConversation) RemoveDocumentation(ref interface{}) {
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
func (this *GlobalConversation) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *GlobalConversation) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** IsClosed **/
func (this *GlobalConversation) IsClosed() bool {
	return this.M_isClosed
}

/** Init reference IsClosed **/
func (this *GlobalConversation) SetIsClosed(ref interface{}) {
	this.NeedSave = true
	this.M_isClosed = ref.(bool)
}

/** Remove reference IsClosed **/

/** ChoreographyRef **/
func (this *GlobalConversation) GetChoreographyRef() []Choreography {
	return this.m_choreographyRef
}

/** Init reference ChoreographyRef **/
func (this *GlobalConversation) SetChoreographyRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_choreographyRef); i++ {
			if this.M_choreographyRef[i] == refStr {
				return
			}
		}
		this.M_choreographyRef = append(this.M_choreographyRef, ref.(string))
	} else {
		this.RemoveChoreographyRef(ref)
		this.m_choreographyRef = append(this.m_choreographyRef, ref.(Choreography))
		this.M_choreographyRef = append(this.M_choreographyRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ChoreographyRef **/
func (this *GlobalConversation) RemoveChoreographyRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	choreographyRef_ := make([]Choreography, 0)
	choreographyRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_choreographyRef); i++ {
		if toDelete.GetUUID() != this.m_choreographyRef[i].(BaseElement).GetUUID() {
			choreographyRef_ = append(choreographyRef_, this.m_choreographyRef[i])
			choreographyRefUuid = append(choreographyRefUuid, this.M_choreographyRef[i])
		}
	}
	this.m_choreographyRef = choreographyRef_
	this.M_choreographyRef = choreographyRefUuid
}

/** Artifact **/
func (this *GlobalConversation) GetArtifact() []Artifact {
	return this.M_artifact
}

/** Init reference Artifact **/
func (this *GlobalConversation) SetArtifact(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var artifacts []Artifact
	for i := 0; i < len(this.M_artifact); i++ {
		if this.M_artifact[i].(BaseElement).GetUUID() != ref.(BaseElement).GetUUID() {
			artifacts = append(artifacts, this.M_artifact[i])
		} else {
			isExist = true
			artifacts = append(artifacts, ref.(Artifact))
		}
	}
	if !isExist {
		artifacts = append(artifacts, ref.(Artifact))
	}
	this.M_artifact = artifacts
}

/** Remove reference Artifact **/
func (this *GlobalConversation) RemoveArtifact(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	artifact_ := make([]Artifact, 0)
	for i := 0; i < len(this.M_artifact); i++ {
		if toDelete.GetUUID() != this.M_artifact[i].(BaseElement).GetUUID() {
			artifact_ = append(artifact_, this.M_artifact[i])
		}
	}
	this.M_artifact = artifact_
}

/** ParticipantAssociation **/
func (this *GlobalConversation) GetParticipantAssociation() []*ParticipantAssociation {
	return this.M_participantAssociation
}

/** Init reference ParticipantAssociation **/
func (this *GlobalConversation) SetParticipantAssociation(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var participantAssociations []*ParticipantAssociation
	for i := 0; i < len(this.M_participantAssociation); i++ {
		if this.M_participantAssociation[i].GetUUID() != ref.(BaseElement).GetUUID() {
			participantAssociations = append(participantAssociations, this.M_participantAssociation[i])
		} else {
			isExist = true
			participantAssociations = append(participantAssociations, ref.(*ParticipantAssociation))
		}
	}
	if !isExist {
		participantAssociations = append(participantAssociations, ref.(*ParticipantAssociation))
	}
	this.M_participantAssociation = participantAssociations
}

/** Remove reference ParticipantAssociation **/
func (this *GlobalConversation) RemoveParticipantAssociation(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	participantAssociation_ := make([]*ParticipantAssociation, 0)
	for i := 0; i < len(this.M_participantAssociation); i++ {
		if toDelete.GetUUID() != this.M_participantAssociation[i].GetUUID() {
			participantAssociation_ = append(participantAssociation_, this.M_participantAssociation[i])
		}
	}
	this.M_participantAssociation = participantAssociation_
}

/** MessageFlowAssociation **/
func (this *GlobalConversation) GetMessageFlowAssociation() []*MessageFlowAssociation {
	return this.M_messageFlowAssociation
}

/** Init reference MessageFlowAssociation **/
func (this *GlobalConversation) SetMessageFlowAssociation(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var messageFlowAssociations []*MessageFlowAssociation
	for i := 0; i < len(this.M_messageFlowAssociation); i++ {
		if this.M_messageFlowAssociation[i].GetUUID() != ref.(BaseElement).GetUUID() {
			messageFlowAssociations = append(messageFlowAssociations, this.M_messageFlowAssociation[i])
		} else {
			isExist = true
			messageFlowAssociations = append(messageFlowAssociations, ref.(*MessageFlowAssociation))
		}
	}
	if !isExist {
		messageFlowAssociations = append(messageFlowAssociations, ref.(*MessageFlowAssociation))
	}
	this.M_messageFlowAssociation = messageFlowAssociations
}

/** Remove reference MessageFlowAssociation **/
func (this *GlobalConversation) RemoveMessageFlowAssociation(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	messageFlowAssociation_ := make([]*MessageFlowAssociation, 0)
	for i := 0; i < len(this.M_messageFlowAssociation); i++ {
		if toDelete.GetUUID() != this.M_messageFlowAssociation[i].GetUUID() {
			messageFlowAssociation_ = append(messageFlowAssociation_, this.M_messageFlowAssociation[i])
		}
	}
	this.M_messageFlowAssociation = messageFlowAssociation_
}

/** ConversationAssociation **/
func (this *GlobalConversation) GetConversationAssociation() []*ConversationAssociation {
	return this.M_conversationAssociation
}

/** Init reference ConversationAssociation **/
func (this *GlobalConversation) SetConversationAssociation(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var conversationAssociations []*ConversationAssociation
	for i := 0; i < len(this.M_conversationAssociation); i++ {
		if this.M_conversationAssociation[i].GetUUID() != ref.(BaseElement).GetUUID() {
			conversationAssociations = append(conversationAssociations, this.M_conversationAssociation[i])
		} else {
			isExist = true
			conversationAssociations = append(conversationAssociations, ref.(*ConversationAssociation))
		}
	}
	if !isExist {
		conversationAssociations = append(conversationAssociations, ref.(*ConversationAssociation))
	}
	this.M_conversationAssociation = conversationAssociations
}

/** Remove reference ConversationAssociation **/
func (this *GlobalConversation) RemoveConversationAssociation(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	conversationAssociation_ := make([]*ConversationAssociation, 0)
	for i := 0; i < len(this.M_conversationAssociation); i++ {
		if toDelete.GetUUID() != this.M_conversationAssociation[i].GetUUID() {
			conversationAssociation_ = append(conversationAssociation_, this.M_conversationAssociation[i])
		}
	}
	this.M_conversationAssociation = conversationAssociation_
}

/** Participant **/
func (this *GlobalConversation) GetParticipant() []*Participant {
	return this.M_participant
}

/** Init reference Participant **/
func (this *GlobalConversation) SetParticipant(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var participants []*Participant
	for i := 0; i < len(this.M_participant); i++ {
		if this.M_participant[i].GetUUID() != ref.(BaseElement).GetUUID() {
			participants = append(participants, this.M_participant[i])
		} else {
			isExist = true
			participants = append(participants, ref.(*Participant))
		}
	}
	if !isExist {
		participants = append(participants, ref.(*Participant))
	}
	this.M_participant = participants
}

/** Remove reference Participant **/
func (this *GlobalConversation) RemoveParticipant(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	participant_ := make([]*Participant, 0)
	for i := 0; i < len(this.M_participant); i++ {
		if toDelete.GetUUID() != this.M_participant[i].GetUUID() {
			participant_ = append(participant_, this.M_participant[i])
		}
	}
	this.M_participant = participant_
}

/** MessageFlow **/
func (this *GlobalConversation) GetMessageFlow() []*MessageFlow {
	return this.M_messageFlow
}

/** Init reference MessageFlow **/
func (this *GlobalConversation) SetMessageFlow(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var messageFlows []*MessageFlow
	for i := 0; i < len(this.M_messageFlow); i++ {
		if this.M_messageFlow[i].GetUUID() != ref.(BaseElement).GetUUID() {
			messageFlows = append(messageFlows, this.M_messageFlow[i])
		} else {
			isExist = true
			messageFlows = append(messageFlows, ref.(*MessageFlow))
		}
	}
	if !isExist {
		messageFlows = append(messageFlows, ref.(*MessageFlow))
	}
	this.M_messageFlow = messageFlows
}

/** Remove reference MessageFlow **/
func (this *GlobalConversation) RemoveMessageFlow(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	messageFlow_ := make([]*MessageFlow, 0)
	for i := 0; i < len(this.M_messageFlow); i++ {
		if toDelete.GetUUID() != this.M_messageFlow[i].GetUUID() {
			messageFlow_ = append(messageFlow_, this.M_messageFlow[i])
		}
	}
	this.M_messageFlow = messageFlow_
}

/** CorrelationKey **/
func (this *GlobalConversation) GetCorrelationKey() []*CorrelationKey {
	return this.M_correlationKey
}

/** Init reference CorrelationKey **/
func (this *GlobalConversation) SetCorrelationKey(ref interface{}) {
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
func (this *GlobalConversation) RemoveCorrelationKey(ref interface{}) {
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

/** ConversationNode **/
func (this *GlobalConversation) GetConversationNode() []ConversationNode {
	return this.M_conversationNode
}

/** Init reference ConversationNode **/
func (this *GlobalConversation) SetConversationNode(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var conversationNodes []ConversationNode
	for i := 0; i < len(this.M_conversationNode); i++ {
		if this.M_conversationNode[i].(BaseElement).GetUUID() != ref.(BaseElement).GetUUID() {
			conversationNodes = append(conversationNodes, this.M_conversationNode[i])
		} else {
			isExist = true
			conversationNodes = append(conversationNodes, ref.(ConversationNode))
		}
	}
	if !isExist {
		conversationNodes = append(conversationNodes, ref.(ConversationNode))
	}
	this.M_conversationNode = conversationNodes
}

/** Remove reference ConversationNode **/
func (this *GlobalConversation) RemoveConversationNode(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	conversationNode_ := make([]ConversationNode, 0)
	for i := 0; i < len(this.M_conversationNode); i++ {
		if toDelete.GetUUID() != this.M_conversationNode[i].(BaseElement).GetUUID() {
			conversationNode_ = append(conversationNode_, this.M_conversationNode[i])
		}
	}
	this.M_conversationNode = conversationNode_
}

/** ConversationLink **/
func (this *GlobalConversation) GetConversationLink() []*ConversationLink {
	return this.M_conversationLink
}

/** Init reference ConversationLink **/
func (this *GlobalConversation) SetConversationLink(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var conversationLinks []*ConversationLink
	for i := 0; i < len(this.M_conversationLink); i++ {
		if this.M_conversationLink[i].GetUUID() != ref.(BaseElement).GetUUID() {
			conversationLinks = append(conversationLinks, this.M_conversationLink[i])
		} else {
			isExist = true
			conversationLinks = append(conversationLinks, ref.(*ConversationLink))
		}
	}
	if !isExist {
		conversationLinks = append(conversationLinks, ref.(*ConversationLink))
	}
	this.M_conversationLink = conversationLinks
}

/** Remove reference ConversationLink **/
func (this *GlobalConversation) RemoveConversationLink(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	conversationLink_ := make([]*ConversationLink, 0)
	for i := 0; i < len(this.M_conversationLink); i++ {
		if toDelete.GetUUID() != this.M_conversationLink[i].GetUUID() {
			conversationLink_ = append(conversationLink_, this.M_conversationLink[i])
		}
	}
	this.M_conversationLink = conversationLink_
}

/** Lane **/
func (this *GlobalConversation) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *GlobalConversation) SetLanePtr(ref interface{}) {
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
func (this *GlobalConversation) RemoveLanePtr(ref interface{}) {
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
func (this *GlobalConversation) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *GlobalConversation) SetOutgoingPtr(ref interface{}) {
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
func (this *GlobalConversation) RemoveOutgoingPtr(ref interface{}) {
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
func (this *GlobalConversation) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *GlobalConversation) SetIncomingPtr(ref interface{}) {
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
func (this *GlobalConversation) RemoveIncomingPtr(ref interface{}) {
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
func (this *GlobalConversation) GetDefinitionsPtr() *Definitions {
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *GlobalConversation) SetDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	} else {
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *GlobalConversation) RemoveDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}

/** Process **/
func (this *GlobalConversation) GetProcessPtr() []*Process {
	return this.m_processPtr
}

/** Init reference Process **/
func (this *GlobalConversation) SetProcessPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_processPtr); i++ {
			if this.M_processPtr[i] == refStr {
				return
			}
		}
		this.M_processPtr = append(this.M_processPtr, ref.(string))
	} else {
		this.RemoveProcessPtr(ref)
		this.m_processPtr = append(this.m_processPtr, ref.(*Process))
		this.M_processPtr = append(this.M_processPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Process **/
func (this *GlobalConversation) RemoveProcessPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	processPtr_ := make([]*Process, 0)
	processPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_processPtr); i++ {
		if toDelete.GetUUID() != this.m_processPtr[i].GetUUID() {
			processPtr_ = append(processPtr_, this.m_processPtr[i])
			processPtrUuid = append(processPtrUuid, this.M_processPtr[i])
		}
	}
	this.m_processPtr = processPtr_
	this.M_processPtr = processPtrUuid
}

/** CallConversation **/
func (this *GlobalConversation) GetCallConversationPtr() []*CallConversation {
	return this.m_callConversationPtr
}

/** Init reference CallConversation **/
func (this *GlobalConversation) SetCallConversationPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_callConversationPtr); i++ {
			if this.M_callConversationPtr[i] == refStr {
				return
			}
		}
		this.M_callConversationPtr = append(this.M_callConversationPtr, ref.(string))
	} else {
		this.RemoveCallConversationPtr(ref)
		this.m_callConversationPtr = append(this.m_callConversationPtr, ref.(*CallConversation))
		this.M_callConversationPtr = append(this.M_callConversationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CallConversation **/
func (this *GlobalConversation) RemoveCallConversationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	callConversationPtr_ := make([]*CallConversation, 0)
	callConversationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_callConversationPtr); i++ {
		if toDelete.GetUUID() != this.m_callConversationPtr[i].GetUUID() {
			callConversationPtr_ = append(callConversationPtr_, this.m_callConversationPtr[i])
			callConversationPtrUuid = append(callConversationPtrUuid, this.M_callConversationPtr[i])
		}
	}
	this.m_callConversationPtr = callConversationPtr_
	this.M_callConversationPtr = callConversationPtrUuid
}
