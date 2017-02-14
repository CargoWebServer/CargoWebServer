// +build BPMN20

package BPMN20

import(
"encoding/xml"
)

type SubChoreography struct{

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

	/** members of FlowElementsContainer **/
	M_flowElement []FlowElement
	M_laneSet []*LaneSet

	/** members of SubChoreography **/
	M_artifact []Artifact


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

/** Xml parser for SubChoreography **/
type XsdSubChoreography struct {
	XMLName xml.Name	`xml:"subChoreography"`
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


	M_flowElement_0	[]*XsdAdHocSubProcess	`xml:"adHocSubProcess,omitempty"`
	M_flowElement_1	[]*XsdBoundaryEvent	`xml:"boundaryEvent,omitempty"`
	M_flowElement_2	[]*XsdBusinessRuleTask	`xml:"businessRuleTask,omitempty"`
	M_flowElement_3	[]*XsdCallActivity	`xml:"callActivity,omitempty"`
	M_flowElement_4	[]*XsdCallChoreography	`xml:"callChoreography,omitempty"`
	M_flowElement_5	[]*XsdChoreographyTask	`xml:"choreographyTask,omitempty"`
	M_flowElement_6	[]*XsdComplexGateway	`xml:"complexGateway,omitempty"`
	M_flowElement_7	[]*XsdDataObject	`xml:"dataObject,omitempty"`
	M_flowElement_8	[]*XsdDataObjectReference	`xml:"dataObjectReference,omitempty"`
	M_flowElement_9	[]*XsdDataStoreReference	`xml:"dataStoreReference,omitempty"`
	M_flowElement_10	[]*XsdEndEvent	`xml:"endEvent,omitempty"`
	M_flowElement_11	[]*XsdEventBasedGateway	`xml:"eventBasedGateway,omitempty"`
	M_flowElement_12	[]*XsdExclusiveGateway	`xml:"exclusiveGateway,omitempty"`
	M_flowElement_13	[]*XsdImplicitThrowEvent	`xml:"implicitThrowEvent,omitempty"`
	M_flowElement_14	[]*XsdInclusiveGateway	`xml:"inclusiveGateway,omitempty"`
	M_flowElement_15	[]*XsdIntermediateCatchEvent	`xml:"intermediateCatchEvent,omitempty"`
	M_flowElement_16	[]*XsdIntermediateThrowEvent	`xml:"intermediateThrowEvent,omitempty"`
	M_flowElement_17	[]*XsdManualTask	`xml:"manualTask,omitempty"`
	M_flowElement_18	[]*XsdParallelGateway	`xml:"parallelGateway,omitempty"`
	M_flowElement_19	[]*XsdReceiveTask	`xml:"receiveTask,omitempty"`
	M_flowElement_20	[]*XsdScriptTask	`xml:"scriptTask,omitempty"`
	M_flowElement_21	[]*XsdSendTask	`xml:"sendTask,omitempty"`
	M_flowElement_22	[]*XsdSequenceFlow	`xml:"sequenceFlow,omitempty"`
	M_flowElement_23	[]*XsdServiceTask	`xml:"serviceTask,omitempty"`
	M_flowElement_24	[]*XsdStartEvent	`xml:"startEvent,omitempty"`
	M_flowElement_25	[]*XsdSubChoreography	`xml:"subChoreography,omitempty"`
	M_flowElement_26	[]*XsdSubProcess	`xml:"subProcess,omitempty"`
	M_flowElement_27	[]*XsdTask	`xml:"task,omitempty"`
	M_flowElement_28	[]*XsdTransaction	`xml:"transaction,omitempty"`
	M_flowElement_29	[]*XsdUserTask	`xml:"userTask,omitempty"`

	M_artifact_0	[]*XsdAssociation	`xml:"association,omitempty"`
	M_artifact_1	[]*XsdGroup	`xml:"group,omitempty"`
	M_artifact_2	[]*XsdTextAnnotation	`xml:"textAnnotation,omitempty"`


}
/** UUID **/
func (this *SubChoreography) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *SubChoreography) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *SubChoreography) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *SubChoreography) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *SubChoreography) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *SubChoreography) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *SubChoreography) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/
func (this *SubChoreography) RemoveExtensionElements(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ExtensionElements)
	if toDelete.GetUUID() == this.M_extensionElements.GetUUID() {
		this.M_extensionElements = nil
	}
}

/** ExtensionDefinitions **/
func (this *SubChoreography) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *SubChoreography) SetExtensionDefinitions(ref interface{}){
	this.NeedSave = true
	isExist := false
	var extensionDefinitionss []*ExtensionDefinition
	for i:=0; i<len(this.M_extensionDefinitions); i++ {
		if this.M_extensionDefinitions[i].GetUUID() != ref.(*ExtensionDefinition).GetUUID() {
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
func (this *SubChoreography) RemoveExtensionDefinitions(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ExtensionDefinition)
	extensionDefinitions_ := make([]*ExtensionDefinition, 0)
	for i := 0; i < len(this.M_extensionDefinitions); i++ {
		if toDelete.GetUUID() != this.M_extensionDefinitions[i].GetUUID() {
			extensionDefinitions_ = append(extensionDefinitions_, this.M_extensionDefinitions[i])
		}
	}
	this.M_extensionDefinitions = extensionDefinitions_
}

/** ExtensionValues **/
func (this *SubChoreography) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *SubChoreography) SetExtensionValues(ref interface{}){
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
func (this *SubChoreography) RemoveExtensionValues(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ExtensionAttributeValue)
	extensionValues_ := make([]*ExtensionAttributeValue, 0)
	for i := 0; i < len(this.M_extensionValues); i++ {
		if toDelete.GetUUID() != this.M_extensionValues[i].GetUUID() {
			extensionValues_ = append(extensionValues_, this.M_extensionValues[i])
		}
	}
	this.M_extensionValues = extensionValues_
}

/** Documentation **/
func (this *SubChoreography) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *SubChoreography) SetDocumentation(ref interface{}){
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
func (this *SubChoreography) RemoveDocumentation(ref interface{}){
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
func (this *SubChoreography) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *SubChoreography) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Auditing **/
func (this *SubChoreography) GetAuditing() *Auditing{
	return this.M_auditing
}

/** Init reference Auditing **/
func (this *SubChoreography) SetAuditing(ref interface{}){
	this.NeedSave = true
	this.M_auditing = ref.(*Auditing)
}

/** Remove reference Auditing **/
func (this *SubChoreography) RemoveAuditing(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_auditing.GetUUID() {
		this.M_auditing = nil
	}
}

/** Monitoring **/
func (this *SubChoreography) GetMonitoring() *Monitoring{
	return this.M_monitoring
}

/** Init reference Monitoring **/
func (this *SubChoreography) SetMonitoring(ref interface{}){
	this.NeedSave = true
	this.M_monitoring = ref.(*Monitoring)
}

/** Remove reference Monitoring **/
func (this *SubChoreography) RemoveMonitoring(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_monitoring.GetUUID() {
		this.M_monitoring = nil
	}
}

/** CategoryValueRef **/
func (this *SubChoreography) GetCategoryValueRef() []*CategoryValue{
	return this.m_categoryValueRef
}

/** Init reference CategoryValueRef **/
func (this *SubChoreography) SetCategoryValueRef(ref interface{}){
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
func (this *SubChoreography) RemoveCategoryValueRef(ref interface{}){
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
func (this *SubChoreography) GetOutgoing() []*SequenceFlow{
	return this.m_outgoing
}

/** Init reference Outgoing **/
func (this *SubChoreography) SetOutgoing(ref interface{}){
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
func (this *SubChoreography) RemoveOutgoing(ref interface{}){
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
func (this *SubChoreography) GetIncoming() []*SequenceFlow{
	return this.m_incoming
}

/** Init reference Incoming **/
func (this *SubChoreography) SetIncoming(ref interface{}){
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
func (this *SubChoreography) RemoveIncoming(ref interface{}){
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
func (this *SubChoreography) GetLanes() []*Lane{
	return this.m_lanes
}

/** Init reference Lanes **/
func (this *SubChoreography) SetLanes(ref interface{}){
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
func (this *SubChoreography) RemoveLanes(ref interface{}){
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
func (this *SubChoreography) GetParticipantRef() []*Participant{
	return this.m_participantRef
}

/** Init reference ParticipantRef **/
func (this *SubChoreography) SetParticipantRef(ref interface{}){
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
func (this *SubChoreography) RemoveParticipantRef(ref interface{}){
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
func (this *SubChoreography) GetInitiatingParticipantRef() *Participant{
	return this.m_initiatingParticipantRef
}

/** Init reference InitiatingParticipantRef **/
func (this *SubChoreography) SetInitiatingParticipantRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_initiatingParticipantRef = ref.(string)
	}else{
		this.m_initiatingParticipantRef = ref.(*Participant)
		this.M_initiatingParticipantRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference InitiatingParticipantRef **/
func (this *SubChoreography) RemoveInitiatingParticipantRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_initiatingParticipantRef.GetUUID() {
		this.m_initiatingParticipantRef = nil
		this.M_initiatingParticipantRef = ""
	}
}

/** CorrelationKey **/
func (this *SubChoreography) GetCorrelationKey() []*CorrelationKey{
	return this.M_correlationKey
}

/** Init reference CorrelationKey **/
func (this *SubChoreography) SetCorrelationKey(ref interface{}){
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
func (this *SubChoreography) RemoveCorrelationKey(ref interface{}){
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
func (this *SubChoreography) GetLoopType() ChoreographyLoopType{
	return this.M_loopType
}

/** Init reference LoopType **/
func (this *SubChoreography) SetLoopType(ref interface{}){
	this.NeedSave = true
	this.M_loopType = ref.(ChoreographyLoopType)
}

/** Remove reference LoopType **/

/** FlowElement **/
func (this *SubChoreography) GetFlowElement() []FlowElement{
	return this.M_flowElement
}

/** Init reference FlowElement **/
func (this *SubChoreography) SetFlowElement(ref interface{}){
	this.NeedSave = true
	isExist := false
	var flowElements []FlowElement
	for i:=0; i<len(this.M_flowElement); i++ {
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
func (this *SubChoreography) RemoveFlowElement(ref interface{}){
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
func (this *SubChoreography) GetLaneSet() []*LaneSet{
	return this.M_laneSet
}

/** Init reference LaneSet **/
func (this *SubChoreography) SetLaneSet(ref interface{}){
	this.NeedSave = true
	isExist := false
	var laneSets []*LaneSet
	for i:=0; i<len(this.M_laneSet); i++ {
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
func (this *SubChoreography) RemoveLaneSet(ref interface{}){
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

/** Artifact **/
func (this *SubChoreography) GetArtifact() []Artifact{
	return this.M_artifact
}

/** Init reference Artifact **/
func (this *SubChoreography) SetArtifact(ref interface{}){
	this.NeedSave = true
	isExist := false
	var artifacts []Artifact
	for i:=0; i<len(this.M_artifact); i++ {
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
func (this *SubChoreography) RemoveArtifact(ref interface{}){
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

/** Lane **/
func (this *SubChoreography) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *SubChoreography) SetLanePtr(ref interface{}){
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
func (this *SubChoreography) RemoveLanePtr(ref interface{}){
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
func (this *SubChoreography) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *SubChoreography) SetOutgoingPtr(ref interface{}){
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
func (this *SubChoreography) RemoveOutgoingPtr(ref interface{}){
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
func (this *SubChoreography) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *SubChoreography) SetIncomingPtr(ref interface{}){
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
func (this *SubChoreography) RemoveIncomingPtr(ref interface{}){
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
func (this *SubChoreography) GetContainerPtr() FlowElementsContainer{
	return this.m_containerPtr
}

/** Init reference Container **/
func (this *SubChoreography) SetContainerPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_containerPtr = ref.(string)
	}else{
		this.m_containerPtr = ref.(FlowElementsContainer)
		this.M_containerPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Container **/
func (this *SubChoreography) RemoveContainerPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_containerPtr.(BaseElement).GetUUID() {
		this.m_containerPtr = nil
		this.M_containerPtr = ""
	}
}
