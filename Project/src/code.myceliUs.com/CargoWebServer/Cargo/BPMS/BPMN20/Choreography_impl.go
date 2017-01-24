// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Choreography_impl struct {

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

	/** members of FlowElementsContainer **/
	M_flowElement []FlowElement
	M_laneSet     []*LaneSet

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

	/** members of Choreography **/
	/** No members **/

	/** Associations **/
	m_collaborationPtr []Collaboration
	/** If the ref is a string and not an object **/
	M_collaborationPtr            []string
	m_callChoreographyActivityPtr []*CallChoreography
	/** If the ref is a string and not an object **/
	M_callChoreographyActivityPtr []string
	m_lanePtr                     []*Lane
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

/** Xml parser for Choreography **/
type XsdChoreography struct {
	XMLName xml.Name `xml:"choreography"`
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

	M_flowElement_0  []*XsdAdHocSubProcess        `xml:"adHocSubProcess,omitempty"`
	M_flowElement_1  []*XsdBoundaryEvent          `xml:"boundaryEvent,omitempty"`
	M_flowElement_2  []*XsdBusinessRuleTask       `xml:"businessRuleTask,omitempty"`
	M_flowElement_3  []*XsdCallActivity           `xml:"callActivity,omitempty"`
	M_flowElement_4  []*XsdCallChoreography       `xml:"callChoreography,omitempty"`
	M_flowElement_5  []*XsdChoreographyTask       `xml:"choreographyTask,omitempty"`
	M_flowElement_6  []*XsdComplexGateway         `xml:"complexGateway,omitempty"`
	M_flowElement_7  []*XsdDataObject             `xml:"dataObject,omitempty"`
	M_flowElement_8  []*XsdDataObjectReference    `xml:"dataObjectReference,omitempty"`
	M_flowElement_9  []*XsdDataStoreReference     `xml:"dataStoreReference,omitempty"`
	M_flowElement_10 []*XsdEndEvent               `xml:"endEvent,omitempty"`
	M_flowElement_11 []*XsdEventBasedGateway      `xml:"eventBasedGateway,omitempty"`
	M_flowElement_12 []*XsdExclusiveGateway       `xml:"exclusiveGateway,omitempty"`
	M_flowElement_13 []*XsdImplicitThrowEvent     `xml:"implicitThrowEvent,omitempty"`
	M_flowElement_14 []*XsdInclusiveGateway       `xml:"inclusiveGateway,omitempty"`
	M_flowElement_15 []*XsdIntermediateCatchEvent `xml:"intermediateCatchEvent,omitempty"`
	M_flowElement_16 []*XsdIntermediateThrowEvent `xml:"intermediateThrowEvent,omitempty"`
	M_flowElement_17 []*XsdManualTask             `xml:"manualTask,omitempty"`
	M_flowElement_18 []*XsdParallelGateway        `xml:"parallelGateway,omitempty"`
	M_flowElement_19 []*XsdReceiveTask            `xml:"receiveTask,omitempty"`
	M_flowElement_20 []*XsdScriptTask             `xml:"scriptTask,omitempty"`
	M_flowElement_21 []*XsdSendTask               `xml:"sendTask,omitempty"`
	M_flowElement_22 []*XsdSequenceFlow           `xml:"sequenceFlow,omitempty"`
	M_flowElement_23 []*XsdServiceTask            `xml:"serviceTask,omitempty"`
	M_flowElement_24 []*XsdStartEvent             `xml:"startEvent,omitempty"`
	M_flowElement_25 []*XsdSubChoreography        `xml:"subChoreography,omitempty"`
	M_flowElement_26 []*XsdSubProcess             `xml:"subProcess,omitempty"`
	M_flowElement_27 []*XsdTask                   `xml:"task,omitempty"`
	M_flowElement_28 []*XsdTransaction            `xml:"transaction,omitempty"`
	M_flowElement_29 []*XsdUserTask               `xml:"userTask,omitempty"`
}

/** UUID **/
func (this *Choreography_impl) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Choreography_impl) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Choreography_impl) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Choreography_impl) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Choreography_impl) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Choreography_impl) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Choreography_impl) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Choreography_impl) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Choreography_impl) SetExtensionDefinitions(ref interface{}) {
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
func (this *Choreography_impl) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Choreography_impl) SetExtensionValues(ref interface{}) {
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
func (this *Choreography_impl) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Choreography_impl) SetDocumentation(ref interface{}) {
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
func (this *Choreography_impl) RemoveDocumentation(ref interface{}) {
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

/** FlowElement **/
func (this *Choreography_impl) GetFlowElement() []FlowElement {
	return this.M_flowElement
}

/** Init reference FlowElement **/
func (this *Choreography_impl) SetFlowElement(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var flowElements []FlowElement
	for i := 0; i < len(this.M_flowElement); i++ {
		if this.M_flowElement[i].(BaseElement).GetUUID() != ref.(BaseElement).GetUUID() {
			flowElements = append(flowElements, this.M_flowElement[i])
		} else {
			isExist = true
			flowElements = append(flowElements, ref.(FlowElement))
		}
	}
	if !isExist {
		flowElements = append(flowElements, ref.(FlowElement))
	}
	this.M_flowElement = flowElements
}

/** Remove reference FlowElement **/
func (this *Choreography_impl) RemoveFlowElement(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	flowElement_ := make([]FlowElement, 0)
	for i := 0; i < len(this.M_flowElement); i++ {
		if toDelete.GetUUID() != this.M_flowElement[i].(BaseElement).GetUUID() {
			flowElement_ = append(flowElement_, this.M_flowElement[i])
		}
	}
	this.M_flowElement = flowElement_
}

/** LaneSet **/
func (this *Choreography_impl) GetLaneSet() []*LaneSet {
	return this.M_laneSet
}

/** Init reference LaneSet **/
func (this *Choreography_impl) SetLaneSet(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var laneSets []*LaneSet
	for i := 0; i < len(this.M_laneSet); i++ {
		if this.M_laneSet[i].GetUUID() != ref.(BaseElement).GetUUID() {
			laneSets = append(laneSets, this.M_laneSet[i])
		} else {
			isExist = true
			laneSets = append(laneSets, ref.(*LaneSet))
		}
	}
	if !isExist {
		laneSets = append(laneSets, ref.(*LaneSet))
	}
	this.M_laneSet = laneSets
}

/** Remove reference LaneSet **/
func (this *Choreography_impl) RemoveLaneSet(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	laneSet_ := make([]*LaneSet, 0)
	for i := 0; i < len(this.M_laneSet); i++ {
		if toDelete.GetUUID() != this.M_laneSet[i].GetUUID() {
			laneSet_ = append(laneSet_, this.M_laneSet[i])
		}
	}
	this.M_laneSet = laneSet_
}

/** Name **/
func (this *Choreography_impl) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Choreography_impl) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** IsClosed **/
func (this *Choreography_impl) IsClosed() bool {
	return this.M_isClosed
}

/** Init reference IsClosed **/
func (this *Choreography_impl) SetIsClosed(ref interface{}) {
	this.NeedSave = true
	this.M_isClosed = ref.(bool)
}

/** Remove reference IsClosed **/

/** ChoreographyRef **/
func (this *Choreography_impl) GetChoreographyRef() []Choreography {
	return this.m_choreographyRef
}

/** Init reference ChoreographyRef **/
func (this *Choreography_impl) SetChoreographyRef(ref interface{}) {
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
func (this *Choreography_impl) RemoveChoreographyRef(ref interface{}) {
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
func (this *Choreography_impl) GetArtifact() []Artifact {
	return this.M_artifact
}

/** Init reference Artifact **/
func (this *Choreography_impl) SetArtifact(ref interface{}) {
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
func (this *Choreography_impl) RemoveArtifact(ref interface{}) {
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
func (this *Choreography_impl) GetParticipantAssociation() []*ParticipantAssociation {
	return this.M_participantAssociation
}

/** Init reference ParticipantAssociation **/
func (this *Choreography_impl) SetParticipantAssociation(ref interface{}) {
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
func (this *Choreography_impl) RemoveParticipantAssociation(ref interface{}) {
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
func (this *Choreography_impl) GetMessageFlowAssociation() []*MessageFlowAssociation {
	return this.M_messageFlowAssociation
}

/** Init reference MessageFlowAssociation **/
func (this *Choreography_impl) SetMessageFlowAssociation(ref interface{}) {
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
func (this *Choreography_impl) RemoveMessageFlowAssociation(ref interface{}) {
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
func (this *Choreography_impl) GetConversationAssociation() []*ConversationAssociation {
	return this.M_conversationAssociation
}

/** Init reference ConversationAssociation **/
func (this *Choreography_impl) SetConversationAssociation(ref interface{}) {
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
func (this *Choreography_impl) RemoveConversationAssociation(ref interface{}) {
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
func (this *Choreography_impl) GetParticipant() []*Participant {
	return this.M_participant
}

/** Init reference Participant **/
func (this *Choreography_impl) SetParticipant(ref interface{}) {
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
func (this *Choreography_impl) RemoveParticipant(ref interface{}) {
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
func (this *Choreography_impl) GetMessageFlow() []*MessageFlow {
	return this.M_messageFlow
}

/** Init reference MessageFlow **/
func (this *Choreography_impl) SetMessageFlow(ref interface{}) {
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
func (this *Choreography_impl) RemoveMessageFlow(ref interface{}) {
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
func (this *Choreography_impl) GetCorrelationKey() []*CorrelationKey {
	return this.M_correlationKey
}

/** Init reference CorrelationKey **/
func (this *Choreography_impl) SetCorrelationKey(ref interface{}) {
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
func (this *Choreography_impl) RemoveCorrelationKey(ref interface{}) {
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
func (this *Choreography_impl) GetConversationNode() []ConversationNode {
	return this.M_conversationNode
}

/** Init reference ConversationNode **/
func (this *Choreography_impl) SetConversationNode(ref interface{}) {
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
func (this *Choreography_impl) RemoveConversationNode(ref interface{}) {
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
func (this *Choreography_impl) GetConversationLink() []*ConversationLink {
	return this.M_conversationLink
}

/** Init reference ConversationLink **/
func (this *Choreography_impl) SetConversationLink(ref interface{}) {
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
func (this *Choreography_impl) RemoveConversationLink(ref interface{}) {
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

/** Collaboration **/
func (this *Choreography_impl) GetCollaborationPtr() []Collaboration {
	return this.m_collaborationPtr
}

/** Init reference Collaboration **/
func (this *Choreography_impl) SetCollaborationPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_collaborationPtr); i++ {
			if this.M_collaborationPtr[i] == refStr {
				return
			}
		}
		this.M_collaborationPtr = append(this.M_collaborationPtr, ref.(string))
	} else {
		this.RemoveCollaborationPtr(ref)
		this.m_collaborationPtr = append(this.m_collaborationPtr, ref.(Collaboration))
		this.M_collaborationPtr = append(this.M_collaborationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Collaboration **/
func (this *Choreography_impl) RemoveCollaborationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	collaborationPtr_ := make([]Collaboration, 0)
	collaborationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_collaborationPtr); i++ {
		if toDelete.GetUUID() != this.m_collaborationPtr[i].(BaseElement).GetUUID() {
			collaborationPtr_ = append(collaborationPtr_, this.m_collaborationPtr[i])
			collaborationPtrUuid = append(collaborationPtrUuid, this.M_collaborationPtr[i])
		}
	}
	this.m_collaborationPtr = collaborationPtr_
	this.M_collaborationPtr = collaborationPtrUuid
}

/** CallChoreographyActivity **/
func (this *Choreography_impl) GetCallChoreographyActivityPtr() []*CallChoreography {
	return this.m_callChoreographyActivityPtr
}

/** Init reference CallChoreographyActivity **/
func (this *Choreography_impl) SetCallChoreographyActivityPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_callChoreographyActivityPtr); i++ {
			if this.M_callChoreographyActivityPtr[i] == refStr {
				return
			}
		}
		this.M_callChoreographyActivityPtr = append(this.M_callChoreographyActivityPtr, ref.(string))
	} else {
		this.RemoveCallChoreographyActivityPtr(ref)
		this.m_callChoreographyActivityPtr = append(this.m_callChoreographyActivityPtr, ref.(*CallChoreography))
		this.M_callChoreographyActivityPtr = append(this.M_callChoreographyActivityPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CallChoreographyActivity **/
func (this *Choreography_impl) RemoveCallChoreographyActivityPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	callChoreographyActivityPtr_ := make([]*CallChoreography, 0)
	callChoreographyActivityPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_callChoreographyActivityPtr); i++ {
		if toDelete.GetUUID() != this.m_callChoreographyActivityPtr[i].GetUUID() {
			callChoreographyActivityPtr_ = append(callChoreographyActivityPtr_, this.m_callChoreographyActivityPtr[i])
			callChoreographyActivityPtrUuid = append(callChoreographyActivityPtrUuid, this.M_callChoreographyActivityPtr[i])
		}
	}
	this.m_callChoreographyActivityPtr = callChoreographyActivityPtr_
	this.M_callChoreographyActivityPtr = callChoreographyActivityPtrUuid
}

/** Lane **/
func (this *Choreography_impl) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Choreography_impl) SetLanePtr(ref interface{}) {
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
func (this *Choreography_impl) RemoveLanePtr(ref interface{}) {
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
func (this *Choreography_impl) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Choreography_impl) SetOutgoingPtr(ref interface{}) {
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
func (this *Choreography_impl) RemoveOutgoingPtr(ref interface{}) {
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
func (this *Choreography_impl) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Choreography_impl) SetIncomingPtr(ref interface{}) {
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
func (this *Choreography_impl) RemoveIncomingPtr(ref interface{}) {
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
func (this *Choreography_impl) GetDefinitionsPtr() *Definitions {
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *Choreography_impl) SetDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	} else {
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *Choreography_impl) RemoveDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}

/** Process **/
func (this *Choreography_impl) GetProcessPtr() []*Process {
	return this.m_processPtr
}

/** Init reference Process **/
func (this *Choreography_impl) SetProcessPtr(ref interface{}) {
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
func (this *Choreography_impl) RemoveProcessPtr(ref interface{}) {
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
func (this *Choreography_impl) GetCallConversationPtr() []*CallConversation {
	return this.m_callConversationPtr
}

/** Init reference CallConversation **/
func (this *Choreography_impl) SetCallConversationPtr(ref interface{}) {
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
func (this *Choreography_impl) RemoveCallConversationPtr(ref interface{}) {
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
