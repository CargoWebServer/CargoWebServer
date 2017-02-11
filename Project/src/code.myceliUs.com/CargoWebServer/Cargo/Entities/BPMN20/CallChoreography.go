package BPMN20

import(
"encoding/xml"
)

type CallChoreography struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of BaseElement **/
	M_id string
	m_other interface{}
	/** If the ref is a string and not an object **/
	M_other string
	M_extensionElements *ExtensionElements
	M_extensionDefinitions []*ExtensionDefinition
	M_extensionValues []*ExtensionAttributeValue
	M_documentation []*Documentation

	/** members of FlowElement **/
	M_name string
	M_auditing *Auditing
	M_monitoring *Monitoring
	m_categoryValueRef []*CategoryValue
	/** If the ref is a string and not an object **/
	M_categoryValueRef []string

	/** members of FlowNode **/
	m_outgoing []*SequenceFlow
	/** If the ref is a string and not an object **/
	M_outgoing []string
	m_incoming []*SequenceFlow
	/** If the ref is a string and not an object **/
	M_incoming []string
	m_lanes []*Lane
	/** If the ref is a string and not an object **/
	M_lanes []string

	/** members of ChoreographyActivity **/
	m_participantRef []*Participant
	/** If the ref is a string and not an object **/
	M_participantRef []string
	m_initiatingParticipantRef *Participant
	/** If the ref is a string and not an object **/
	M_initiatingParticipantRef string
	M_correlationKey []*CorrelationKey
	M_loopType ChoreographyLoopType

	/** members of CallChoreography **/
	m_calledChoreographyRef Choreography
	/** If the ref is a string and not an object **/
	M_calledChoreographyRef string
	M_participantAssociation []*ParticipantAssociation


	/** Associations **/
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
	m_containerPtr FlowElementsContainer
	/** If the ref is a string and not an object **/
	M_containerPtr string
}

/** Xml parser for CallChoreography **/
type XsdCallChoreography struct {
	XMLName xml.Name	`xml:"callChoreography"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** FlowElement **/
	M_auditing	*XsdAuditing	`xml:"auditing,omitempty"`
	M_monitoring	*XsdMonitoring	`xml:"monitoring,omitempty"`
	M_categoryValueRef	[]string	`xml:"categoryValueRef"`
	M_name	string	`xml:"name,attr"`


	/** FlowNode **/
	M_incoming	[]string	`xml:"incoming"`
	M_outgoing	[]string	`xml:"outgoing"`


	/** ChoreographyActivity **/
	M_participantRef	[]string	`xml:"participantRef"`
	M_correlationKey	[]*XsdCorrelationKey	`xml:"correlationKey,omitempty"`
	M_initiatingParticipantRef	string	`xml:"initiatingParticipantRef,attr"`
	M_loopType	string	`xml:"loopType,attr"`


	M_participantAssociation	[]*XsdParticipantAssociation	`xml:"participantAssociation,omitempty"`
	M_calledChoreographyRef	string	`xml:"calledChoreographyRef,attr"`

}
/** UUID **/
func (this *CallChoreography) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *CallChoreography) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *CallChoreography) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *CallChoreography) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *CallChoreography) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *CallChoreography) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *CallChoreography) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *CallChoreography) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *CallChoreography) SetExtensionDefinitions(ref interface{}){
	this.NeedSave = true
	isExist := false
	var extensionDefinitionss []*ExtensionDefinition
	for i:=0; i<len(this.M_extensionDefinitions); i++ {
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
func (this *CallChoreography) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *CallChoreography) SetExtensionValues(ref interface{}){
	this.NeedSave = true
	isExist := false
	var extensionValuess []*ExtensionAttributeValue
	for i:=0; i<len(this.M_extensionValues); i++ {
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
func (this *CallChoreography) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *CallChoreography) SetDocumentation(ref interface{}){
	this.NeedSave = true
	isExist := false
	var documentations []*Documentation
	for i:=0; i<len(this.M_documentation); i++ {
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
func (this *CallChoreography) RemoveDocumentation(ref interface{}){
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
func (this *CallChoreography) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *CallChoreography) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Auditing **/
func (this *CallChoreography) GetAuditing() *Auditing{
	return this.M_auditing
}

/** Init reference Auditing **/
func (this *CallChoreography) SetAuditing(ref interface{}){
	this.NeedSave = true
	this.M_auditing = ref.(*Auditing)
}

/** Remove reference Auditing **/
func (this *CallChoreography) RemoveAuditing(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_auditing.GetUUID() {
		this.M_auditing = nil
	}
}

/** Monitoring **/
func (this *CallChoreography) GetMonitoring() *Monitoring{
	return this.M_monitoring
}

/** Init reference Monitoring **/
func (this *CallChoreography) SetMonitoring(ref interface{}){
	this.NeedSave = true
	this.M_monitoring = ref.(*Monitoring)
}

/** Remove reference Monitoring **/
func (this *CallChoreography) RemoveMonitoring(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_monitoring.GetUUID() {
		this.M_monitoring = nil
	}
}

/** CategoryValueRef **/
func (this *CallChoreography) GetCategoryValueRef() []*CategoryValue{
	return this.m_categoryValueRef
}

/** Init reference CategoryValueRef **/
func (this *CallChoreography) SetCategoryValueRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_categoryValueRef); i++ {
			if this.M_categoryValueRef[i] == refStr {
				return
			}
		}
		this.M_categoryValueRef = append(this.M_categoryValueRef, ref.(string))
	}else{
		this.RemoveCategoryValueRef(ref)
		this.m_categoryValueRef = append(this.m_categoryValueRef, ref.(*CategoryValue))
		this.M_categoryValueRef = append(this.M_categoryValueRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CategoryValueRef **/
func (this *CallChoreography) RemoveCategoryValueRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	categoryValueRef_ := make([]*CategoryValue, 0)
	categoryValueRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_categoryValueRef); i++ {
		if toDelete.GetUUID() != this.m_categoryValueRef[i].GetUUID() {
			categoryValueRef_ = append(categoryValueRef_, this.m_categoryValueRef[i])
			categoryValueRefUuid = append(categoryValueRefUuid, this.M_categoryValueRef[i])
		}
	}
	this.m_categoryValueRef = categoryValueRef_
	this.M_categoryValueRef = categoryValueRefUuid
}

/** Outgoing **/
func (this *CallChoreography) GetOutgoing() []*SequenceFlow{
	return this.m_outgoing
}

/** Init reference Outgoing **/
func (this *CallChoreography) SetOutgoing(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_outgoing); i++ {
			if this.M_outgoing[i] == refStr {
				return
			}
		}
		this.M_outgoing = append(this.M_outgoing, ref.(string))
	}else{
		this.RemoveOutgoing(ref)
		this.m_outgoing = append(this.m_outgoing, ref.(*SequenceFlow))
		this.M_outgoing = append(this.M_outgoing, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Outgoing **/
func (this *CallChoreography) RemoveOutgoing(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	outgoing_ := make([]*SequenceFlow, 0)
	outgoingUuid := make([]string, 0)
	for i := 0; i < len(this.m_outgoing); i++ {
		if toDelete.GetUUID() != this.m_outgoing[i].GetUUID() {
			outgoing_ = append(outgoing_, this.m_outgoing[i])
			outgoingUuid = append(outgoingUuid, this.M_outgoing[i])
		}
	}
	this.m_outgoing = outgoing_
	this.M_outgoing = outgoingUuid
}

/** Incoming **/
func (this *CallChoreography) GetIncoming() []*SequenceFlow{
	return this.m_incoming
}

/** Init reference Incoming **/
func (this *CallChoreography) SetIncoming(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_incoming); i++ {
			if this.M_incoming[i] == refStr {
				return
			}
		}
		this.M_incoming = append(this.M_incoming, ref.(string))
	}else{
		this.RemoveIncoming(ref)
		this.m_incoming = append(this.m_incoming, ref.(*SequenceFlow))
		this.M_incoming = append(this.M_incoming, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Incoming **/
func (this *CallChoreography) RemoveIncoming(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	incoming_ := make([]*SequenceFlow, 0)
	incomingUuid := make([]string, 0)
	for i := 0; i < len(this.m_incoming); i++ {
		if toDelete.GetUUID() != this.m_incoming[i].GetUUID() {
			incoming_ = append(incoming_, this.m_incoming[i])
			incomingUuid = append(incomingUuid, this.M_incoming[i])
		}
	}
	this.m_incoming = incoming_
	this.M_incoming = incomingUuid
}

/** Lanes **/
func (this *CallChoreography) GetLanes() []*Lane{
	return this.m_lanes
}

/** Init reference Lanes **/
func (this *CallChoreography) SetLanes(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_lanes); i++ {
			if this.M_lanes[i] == refStr {
				return
			}
		}
		this.M_lanes = append(this.M_lanes, ref.(string))
	}else{
		this.RemoveLanes(ref)
		this.m_lanes = append(this.m_lanes, ref.(*Lane))
		this.M_lanes = append(this.M_lanes, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Lanes **/
func (this *CallChoreography) RemoveLanes(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	lanes_ := make([]*Lane, 0)
	lanesUuid := make([]string, 0)
	for i := 0; i < len(this.m_lanes); i++ {
		if toDelete.GetUUID() != this.m_lanes[i].GetUUID() {
			lanes_ = append(lanes_, this.m_lanes[i])
			lanesUuid = append(lanesUuid, this.M_lanes[i])
		}
	}
	this.m_lanes = lanes_
	this.M_lanes = lanesUuid
}

/** ParticipantRef **/
func (this *CallChoreography) GetParticipantRef() []*Participant{
	return this.m_participantRef
}

/** Init reference ParticipantRef **/
func (this *CallChoreography) SetParticipantRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_participantRef); i++ {
			if this.M_participantRef[i] == refStr {
				return
			}
		}
		this.M_participantRef = append(this.M_participantRef, ref.(string))
	}else{
		this.RemoveParticipantRef(ref)
		this.m_participantRef = append(this.m_participantRef, ref.(*Participant))
		this.M_participantRef = append(this.M_participantRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ParticipantRef **/
func (this *CallChoreography) RemoveParticipantRef(ref interface{}){
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

/** InitiatingParticipantRef **/
func (this *CallChoreography) GetInitiatingParticipantRef() *Participant{
	return this.m_initiatingParticipantRef
}

/** Init reference InitiatingParticipantRef **/
func (this *CallChoreography) SetInitiatingParticipantRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_initiatingParticipantRef = ref.(string)
	}else{
		this.m_initiatingParticipantRef = ref.(*Participant)
		this.M_initiatingParticipantRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference InitiatingParticipantRef **/
func (this *CallChoreography) RemoveInitiatingParticipantRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_initiatingParticipantRef.GetUUID() {
		this.m_initiatingParticipantRef = nil
		this.M_initiatingParticipantRef = ""
	}
}

/** CorrelationKey **/
func (this *CallChoreography) GetCorrelationKey() []*CorrelationKey{
	return this.M_correlationKey
}

/** Init reference CorrelationKey **/
func (this *CallChoreography) SetCorrelationKey(ref interface{}){
	this.NeedSave = true
	isExist := false
	var correlationKeys []*CorrelationKey
	for i:=0; i<len(this.M_correlationKey); i++ {
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
func (this *CallChoreography) RemoveCorrelationKey(ref interface{}){
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

/** LoopType **/
func (this *CallChoreography) GetLoopType() ChoreographyLoopType{
	return this.M_loopType
}

/** Init reference LoopType **/
func (this *CallChoreography) SetLoopType(ref interface{}){
	this.NeedSave = true
	this.M_loopType = ref.(ChoreographyLoopType)
}

/** Remove reference LoopType **/

/** CalledChoreographyRef **/
func (this *CallChoreography) GetCalledChoreographyRef() Choreography{
	return this.m_calledChoreographyRef
}

/** Init reference CalledChoreographyRef **/
func (this *CallChoreography) SetCalledChoreographyRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_calledChoreographyRef = ref.(string)
	}else{
		this.m_calledChoreographyRef = ref.(Choreography)
		this.M_calledChoreographyRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference CalledChoreographyRef **/
func (this *CallChoreography) RemoveCalledChoreographyRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_calledChoreographyRef.(BaseElement).GetUUID() {
		this.m_calledChoreographyRef = nil
		this.M_calledChoreographyRef = ""
	}
}

/** ParticipantAssociation **/
func (this *CallChoreography) GetParticipantAssociation() []*ParticipantAssociation{
	return this.M_participantAssociation
}

/** Init reference ParticipantAssociation **/
func (this *CallChoreography) SetParticipantAssociation(ref interface{}){
	this.NeedSave = true
	isExist := false
	var participantAssociations []*ParticipantAssociation
	for i:=0; i<len(this.M_participantAssociation); i++ {
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
func (this *CallChoreography) RemoveParticipantAssociation(ref interface{}){
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

/** Lane **/
func (this *CallChoreography) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *CallChoreography) SetLanePtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_lanePtr); i++ {
			if this.M_lanePtr[i] == refStr {
				return
			}
		}
		this.M_lanePtr = append(this.M_lanePtr, ref.(string))
	}else{
		this.RemoveLanePtr(ref)
		this.m_lanePtr = append(this.m_lanePtr, ref.(*Lane))
		this.M_lanePtr = append(this.M_lanePtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Lane **/
func (this *CallChoreography) RemoveLanePtr(ref interface{}){
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
func (this *CallChoreography) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *CallChoreography) SetOutgoingPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_outgoingPtr); i++ {
			if this.M_outgoingPtr[i] == refStr {
				return
			}
		}
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(string))
	}else{
		this.RemoveOutgoingPtr(ref)
		this.m_outgoingPtr = append(this.m_outgoingPtr, ref.(*Association))
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Outgoing **/
func (this *CallChoreography) RemoveOutgoingPtr(ref interface{}){
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
func (this *CallChoreography) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *CallChoreography) SetIncomingPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_incomingPtr); i++ {
			if this.M_incomingPtr[i] == refStr {
				return
			}
		}
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(string))
	}else{
		this.RemoveIncomingPtr(ref)
		this.m_incomingPtr = append(this.m_incomingPtr, ref.(*Association))
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Incoming **/
func (this *CallChoreography) RemoveIncomingPtr(ref interface{}){
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

/** Container **/
func (this *CallChoreography) GetContainerPtr() FlowElementsContainer{
	return this.m_containerPtr
}

/** Init reference Container **/
func (this *CallChoreography) SetContainerPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_containerPtr = ref.(string)
	}else{
		this.m_containerPtr = ref.(FlowElementsContainer)
		this.M_containerPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Container **/
func (this *CallChoreography) RemoveContainerPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_containerPtr.(BaseElement).GetUUID() {
		this.m_containerPtr = nil
		this.M_containerPtr = ""
	}
}
