// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Participant struct {

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

	/** members of Participant **/
	M_name         string
	m_interfaceRef []*Interface
	/** If the ref is a string and not an object **/
	M_interfaceRef            []string
	M_participantMultiplicity *ParticipantMultiplicity
	m_endPointRef             []*EndPoint
	/** If the ref is a string and not an object **/
	M_endPointRef []string
	m_processRef  *Process
	/** If the ref is a string and not an object **/
	M_processRef string

	/** Associations **/
	m_conversationNodePtr []ConversationNode
	/** If the ref is a string and not an object **/
	M_conversationNodePtr []string
	m_partnerEntityRefPtr []*PartnerEntity
	/** If the ref is a string and not an object **/
	M_partnerEntityRefPtr []string
	m_partnerRoleRefPtr   []*PartnerRole
	/** If the ref is a string and not an object **/
	M_partnerRoleRefPtr         []string
	m_participantAssociationPtr []*ParticipantAssociation
	/** If the ref is a string and not an object **/
	M_participantAssociationPtr []string
	m_collaborationPtr          Collaboration
	/** If the ref is a string and not an object **/
	M_collaborationPtr        string
	m_choreographyActivityPtr []ChoreographyActivity
	/** If the ref is a string and not an object **/
	M_choreographyActivityPtr   []string
	m_globalChoreographyTaskPtr []*GlobalChoreographyTask
	/** If the ref is a string and not an object **/
	M_globalChoreographyTaskPtr []string
	m_lanePtr                   []*Lane
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
	M_messageFlowPtr []string
}

/** Xml parser for Participant **/
type XsdParticipant struct {
	XMLName xml.Name `xml:"participant"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_interfaceRef            []string                    `xml:"interfaceRef"`
	M_endPointRef             []string                    `xml:"endPointRef"`
	M_participantMultiplicity *XsdParticipantMultiplicity `xml:"participantMultiplicity,omitempty"`
	M_name                    string                      `xml:"name,attr"`
	M_processRef              string                      `xml:"processRef,attr"`
}

/** UUID **/
func (this *Participant) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Participant) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Participant) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Participant) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Participant) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Participant) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Participant) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Participant) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Participant) SetExtensionDefinitions(ref interface{}) {
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
func (this *Participant) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Participant) SetExtensionValues(ref interface{}) {
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
func (this *Participant) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Participant) SetDocumentation(ref interface{}) {
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
func (this *Participant) RemoveDocumentation(ref interface{}) {
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
func (this *Participant) GetIncomingConversationLinks() []*ConversationLink {
	return this.m_incomingConversationLinks
}

/** Init reference IncomingConversationLinks **/
func (this *Participant) SetIncomingConversationLinks(ref interface{}) {
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
func (this *Participant) RemoveIncomingConversationLinks(ref interface{}) {
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
func (this *Participant) GetOutgoingConversationLinks() []*ConversationLink {
	return this.m_outgoingConversationLinks
}

/** Init reference OutgoingConversationLinks **/
func (this *Participant) SetOutgoingConversationLinks(ref interface{}) {
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
func (this *Participant) RemoveOutgoingConversationLinks(ref interface{}) {
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
func (this *Participant) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Participant) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** InterfaceRef **/
func (this *Participant) GetInterfaceRef() []*Interface {
	return this.m_interfaceRef
}

/** Init reference InterfaceRef **/
func (this *Participant) SetInterfaceRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_interfaceRef); i++ {
			if this.M_interfaceRef[i] == refStr {
				return
			}
		}
		this.M_interfaceRef = append(this.M_interfaceRef, ref.(string))
	} else {
		this.RemoveInterfaceRef(ref)
		this.m_interfaceRef = append(this.m_interfaceRef, ref.(*Interface))
		this.M_interfaceRef = append(this.M_interfaceRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference InterfaceRef **/
func (this *Participant) RemoveInterfaceRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	interfaceRef_ := make([]*Interface, 0)
	interfaceRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_interfaceRef); i++ {
		if toDelete.GetUUID() != this.m_interfaceRef[i].GetUUID() {
			interfaceRef_ = append(interfaceRef_, this.m_interfaceRef[i])
			interfaceRefUuid = append(interfaceRefUuid, this.M_interfaceRef[i])
		}
	}
	this.m_interfaceRef = interfaceRef_
	this.M_interfaceRef = interfaceRefUuid
}

/** ParticipantMultiplicity **/
func (this *Participant) GetParticipantMultiplicity() *ParticipantMultiplicity {
	return this.M_participantMultiplicity
}

/** Init reference ParticipantMultiplicity **/
func (this *Participant) SetParticipantMultiplicity(ref interface{}) {
	this.NeedSave = true
	this.M_participantMultiplicity = ref.(*ParticipantMultiplicity)
}

/** Remove reference ParticipantMultiplicity **/
func (this *Participant) RemoveParticipantMultiplicity(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_participantMultiplicity.GetUUID() {
		this.M_participantMultiplicity = nil
	}
}

/** EndPointRef **/
func (this *Participant) GetEndPointRef() []*EndPoint {
	return this.m_endPointRef
}

/** Init reference EndPointRef **/
func (this *Participant) SetEndPointRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_endPointRef); i++ {
			if this.M_endPointRef[i] == refStr {
				return
			}
		}
		this.M_endPointRef = append(this.M_endPointRef, ref.(string))
	} else {
		this.RemoveEndPointRef(ref)
		this.m_endPointRef = append(this.m_endPointRef, ref.(*EndPoint))
		this.M_endPointRef = append(this.M_endPointRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference EndPointRef **/
func (this *Participant) RemoveEndPointRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	endPointRef_ := make([]*EndPoint, 0)
	endPointRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_endPointRef); i++ {
		if toDelete.GetUUID() != this.m_endPointRef[i].GetUUID() {
			endPointRef_ = append(endPointRef_, this.m_endPointRef[i])
			endPointRefUuid = append(endPointRefUuid, this.M_endPointRef[i])
		}
	}
	this.m_endPointRef = endPointRef_
	this.M_endPointRef = endPointRefUuid
}

/** ProcessRef **/
func (this *Participant) GetProcessRef() *Process {
	return this.m_processRef
}

/** Init reference ProcessRef **/
func (this *Participant) SetProcessRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_processRef = ref.(string)
	} else {
		this.m_processRef = ref.(*Process)
		this.M_processRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ProcessRef **/
func (this *Participant) RemoveProcessRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_processRef.GetUUID() {
		this.m_processRef = nil
		this.M_processRef = ""
	}
}

/** ConversationNode **/
func (this *Participant) GetConversationNodePtr() []ConversationNode {
	return this.m_conversationNodePtr
}

/** Init reference ConversationNode **/
func (this *Participant) SetConversationNodePtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_conversationNodePtr); i++ {
			if this.M_conversationNodePtr[i] == refStr {
				return
			}
		}
		this.M_conversationNodePtr = append(this.M_conversationNodePtr, ref.(string))
	} else {
		this.RemoveConversationNodePtr(ref)
		this.m_conversationNodePtr = append(this.m_conversationNodePtr, ref.(ConversationNode))
		this.M_conversationNodePtr = append(this.M_conversationNodePtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ConversationNode **/
func (this *Participant) RemoveConversationNodePtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	conversationNodePtr_ := make([]ConversationNode, 0)
	conversationNodePtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_conversationNodePtr); i++ {
		if toDelete.GetUUID() != this.m_conversationNodePtr[i].(BaseElement).GetUUID() {
			conversationNodePtr_ = append(conversationNodePtr_, this.m_conversationNodePtr[i])
			conversationNodePtrUuid = append(conversationNodePtrUuid, this.M_conversationNodePtr[i])
		}
	}
	this.m_conversationNodePtr = conversationNodePtr_
	this.M_conversationNodePtr = conversationNodePtrUuid
}

/** PartnerEntityRef **/
func (this *Participant) GetPartnerEntityRefPtr() []*PartnerEntity {
	return this.m_partnerEntityRefPtr
}

/** Init reference PartnerEntityRef **/
func (this *Participant) SetPartnerEntityRefPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_partnerEntityRefPtr); i++ {
			if this.M_partnerEntityRefPtr[i] == refStr {
				return
			}
		}
		this.M_partnerEntityRefPtr = append(this.M_partnerEntityRefPtr, ref.(string))
	} else {
		this.RemovePartnerEntityRefPtr(ref)
		this.m_partnerEntityRefPtr = append(this.m_partnerEntityRefPtr, ref.(*PartnerEntity))
		this.M_partnerEntityRefPtr = append(this.M_partnerEntityRefPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference PartnerEntityRef **/
func (this *Participant) RemovePartnerEntityRefPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	partnerEntityRefPtr_ := make([]*PartnerEntity, 0)
	partnerEntityRefPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_partnerEntityRefPtr); i++ {
		if toDelete.GetUUID() != this.m_partnerEntityRefPtr[i].GetUUID() {
			partnerEntityRefPtr_ = append(partnerEntityRefPtr_, this.m_partnerEntityRefPtr[i])
			partnerEntityRefPtrUuid = append(partnerEntityRefPtrUuid, this.M_partnerEntityRefPtr[i])
		}
	}
	this.m_partnerEntityRefPtr = partnerEntityRefPtr_
	this.M_partnerEntityRefPtr = partnerEntityRefPtrUuid
}

/** PartnerRoleRef **/
func (this *Participant) GetPartnerRoleRefPtr() []*PartnerRole {
	return this.m_partnerRoleRefPtr
}

/** Init reference PartnerRoleRef **/
func (this *Participant) SetPartnerRoleRefPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_partnerRoleRefPtr); i++ {
			if this.M_partnerRoleRefPtr[i] == refStr {
				return
			}
		}
		this.M_partnerRoleRefPtr = append(this.M_partnerRoleRefPtr, ref.(string))
	} else {
		this.RemovePartnerRoleRefPtr(ref)
		this.m_partnerRoleRefPtr = append(this.m_partnerRoleRefPtr, ref.(*PartnerRole))
		this.M_partnerRoleRefPtr = append(this.M_partnerRoleRefPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference PartnerRoleRef **/
func (this *Participant) RemovePartnerRoleRefPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	partnerRoleRefPtr_ := make([]*PartnerRole, 0)
	partnerRoleRefPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_partnerRoleRefPtr); i++ {
		if toDelete.GetUUID() != this.m_partnerRoleRefPtr[i].GetUUID() {
			partnerRoleRefPtr_ = append(partnerRoleRefPtr_, this.m_partnerRoleRefPtr[i])
			partnerRoleRefPtrUuid = append(partnerRoleRefPtrUuid, this.M_partnerRoleRefPtr[i])
		}
	}
	this.m_partnerRoleRefPtr = partnerRoleRefPtr_
	this.M_partnerRoleRefPtr = partnerRoleRefPtrUuid
}

/** ParticipantAssociation **/
func (this *Participant) GetParticipantAssociationPtr() []*ParticipantAssociation {
	return this.m_participantAssociationPtr
}

/** Init reference ParticipantAssociation **/
func (this *Participant) SetParticipantAssociationPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_participantAssociationPtr); i++ {
			if this.M_participantAssociationPtr[i] == refStr {
				return
			}
		}
		this.M_participantAssociationPtr = append(this.M_participantAssociationPtr, ref.(string))
	} else {
		this.RemoveParticipantAssociationPtr(ref)
		this.m_participantAssociationPtr = append(this.m_participantAssociationPtr, ref.(*ParticipantAssociation))
		this.M_participantAssociationPtr = append(this.M_participantAssociationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ParticipantAssociation **/
func (this *Participant) RemoveParticipantAssociationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	participantAssociationPtr_ := make([]*ParticipantAssociation, 0)
	participantAssociationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_participantAssociationPtr); i++ {
		if toDelete.GetUUID() != this.m_participantAssociationPtr[i].GetUUID() {
			participantAssociationPtr_ = append(participantAssociationPtr_, this.m_participantAssociationPtr[i])
			participantAssociationPtrUuid = append(participantAssociationPtrUuid, this.M_participantAssociationPtr[i])
		}
	}
	this.m_participantAssociationPtr = participantAssociationPtr_
	this.M_participantAssociationPtr = participantAssociationPtrUuid
}

/** Collaboration **/
func (this *Participant) GetCollaborationPtr() Collaboration {
	return this.m_collaborationPtr
}

/** Init reference Collaboration **/
func (this *Participant) SetCollaborationPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_collaborationPtr = ref.(string)
	} else {
		this.m_collaborationPtr = ref.(Collaboration)
		this.M_collaborationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Collaboration **/
func (this *Participant) RemoveCollaborationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_collaborationPtr.(BaseElement).GetUUID() {
		this.m_collaborationPtr = nil
		this.M_collaborationPtr = ""
	}
}

/** ChoreographyActivity **/
func (this *Participant) GetChoreographyActivityPtr() []ChoreographyActivity {
	return this.m_choreographyActivityPtr
}

/** Init reference ChoreographyActivity **/
func (this *Participant) SetChoreographyActivityPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_choreographyActivityPtr); i++ {
			if this.M_choreographyActivityPtr[i] == refStr {
				return
			}
		}
		this.M_choreographyActivityPtr = append(this.M_choreographyActivityPtr, ref.(string))
	} else {
		this.RemoveChoreographyActivityPtr(ref)
		this.m_choreographyActivityPtr = append(this.m_choreographyActivityPtr, ref.(ChoreographyActivity))
		this.M_choreographyActivityPtr = append(this.M_choreographyActivityPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ChoreographyActivity **/
func (this *Participant) RemoveChoreographyActivityPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	choreographyActivityPtr_ := make([]ChoreographyActivity, 0)
	choreographyActivityPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_choreographyActivityPtr); i++ {
		if toDelete.GetUUID() != this.m_choreographyActivityPtr[i].(BaseElement).GetUUID() {
			choreographyActivityPtr_ = append(choreographyActivityPtr_, this.m_choreographyActivityPtr[i])
			choreographyActivityPtrUuid = append(choreographyActivityPtrUuid, this.M_choreographyActivityPtr[i])
		}
	}
	this.m_choreographyActivityPtr = choreographyActivityPtr_
	this.M_choreographyActivityPtr = choreographyActivityPtrUuid
}

/** GlobalChoreographyTask **/
func (this *Participant) GetGlobalChoreographyTaskPtr() []*GlobalChoreographyTask {
	return this.m_globalChoreographyTaskPtr
}

/** Init reference GlobalChoreographyTask **/
func (this *Participant) SetGlobalChoreographyTaskPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_globalChoreographyTaskPtr); i++ {
			if this.M_globalChoreographyTaskPtr[i] == refStr {
				return
			}
		}
		this.M_globalChoreographyTaskPtr = append(this.M_globalChoreographyTaskPtr, ref.(string))
	} else {
		this.RemoveGlobalChoreographyTaskPtr(ref)
		this.m_globalChoreographyTaskPtr = append(this.m_globalChoreographyTaskPtr, ref.(*GlobalChoreographyTask))
		this.M_globalChoreographyTaskPtr = append(this.M_globalChoreographyTaskPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference GlobalChoreographyTask **/
func (this *Participant) RemoveGlobalChoreographyTaskPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	globalChoreographyTaskPtr_ := make([]*GlobalChoreographyTask, 0)
	globalChoreographyTaskPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_globalChoreographyTaskPtr); i++ {
		if toDelete.GetUUID() != this.m_globalChoreographyTaskPtr[i].GetUUID() {
			globalChoreographyTaskPtr_ = append(globalChoreographyTaskPtr_, this.m_globalChoreographyTaskPtr[i])
			globalChoreographyTaskPtrUuid = append(globalChoreographyTaskPtrUuid, this.M_globalChoreographyTaskPtr[i])
		}
	}
	this.m_globalChoreographyTaskPtr = globalChoreographyTaskPtr_
	this.M_globalChoreographyTaskPtr = globalChoreographyTaskPtrUuid
}

/** Lane **/
func (this *Participant) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Participant) SetLanePtr(ref interface{}) {
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
func (this *Participant) RemoveLanePtr(ref interface{}) {
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
func (this *Participant) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Participant) SetOutgoingPtr(ref interface{}) {
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
func (this *Participant) RemoveOutgoingPtr(ref interface{}) {
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
func (this *Participant) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Participant) SetIncomingPtr(ref interface{}) {
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
func (this *Participant) RemoveIncomingPtr(ref interface{}) {
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
func (this *Participant) GetMessageFlowPtr() []*MessageFlow {
	return this.m_messageFlowPtr
}

/** Init reference MessageFlow **/
func (this *Participant) SetMessageFlowPtr(ref interface{}) {
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
func (this *Participant) RemoveMessageFlowPtr(ref interface{}) {
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
