// +build BPMN20

package BPMN20

import(
"encoding/xml"
)

type EndEvent struct{

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

	/** members of InteractionNode **/
	m_incomingConversationLinks []*ConversationLink
	/** If the ref is a string and not an object **/
	M_incomingConversationLinks []string
	m_outgoingConversationLinks []*ConversationLink
	/** If the ref is a string and not an object **/
	M_outgoingConversationLinks []string

	/** members of Event **/
	M_property []*Property

	/** members of ThrowEvent **/
	M_inputSet *InputSet
	m_eventDefinitionRef []EventDefinition
	/** If the ref is a string and not an object **/
	M_eventDefinitionRef []string
	M_dataInputAssociation []*DataInputAssociation
	M_dataInput []*DataInput
	M_eventDefinition []EventDefinition

	/** members of EndEvent **/
	/** No members **/


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
	m_messageFlowPtr []*MessageFlow
	/** If the ref is a string and not an object **/
	M_messageFlowPtr []string
}

/** Xml parser for EndEvent **/
type XsdEndEvent struct {
	XMLName xml.Name	`xml:"endEvent"`
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


	/** Event **/
	M_property	[]*XsdProperty	`xml:"property,omitempty"`


	/** ThrowEvent **/
	M_dataInput	[]*XsdDataInput	`xml:"dataInput,omitempty"`
	M_dataInputAssociation	[]*XsdDataInputAssociation	`xml:"dataInputAssociation,omitempty"`
	M_inputSet	*XsdInputSet	`xml:"inputSet,omitempty"`
	M_eventDefinition_0	[]*XsdCancelEventDefinition	`xml:"cancelEventDefinition,omitempty"`
	M_eventDefinition_1	[]*XsdCompensateEventDefinition	`xml:"compensateEventDefinition,omitempty"`
	M_eventDefinition_2	[]*XsdConditionalEventDefinition	`xml:"conditionalEventDefinition,omitempty"`
	M_eventDefinition_3	[]*XsdErrorEventDefinition	`xml:"errorEventDefinition,omitempty"`
	M_eventDefinition_4	[]*XsdEscalationEventDefinition	`xml:"escalationEventDefinition,omitempty"`
	M_eventDefinition_5	[]*XsdLinkEventDefinition	`xml:"linkEventDefinition,omitempty"`
	M_eventDefinition_6	[]*XsdMessageEventDefinition	`xml:"messageEventDefinition,omitempty"`
	M_eventDefinition_7	[]*XsdSignalEventDefinition	`xml:"signalEventDefinition,omitempty"`
	M_eventDefinition_8	[]*XsdTerminateEventDefinition	`xml:"terminateEventDefinition,omitempty"`
	M_eventDefinition_9	[]*XsdTimerEventDefinition	`xml:"timerEventDefinition,omitempty"`

	M_eventDefinitionRef	[]string	`xml:"eventDefinitionRef"`



}
/** UUID **/
func (this *EndEvent) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *EndEvent) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *EndEvent) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *EndEvent) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *EndEvent) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *EndEvent) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *EndEvent) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/
func (this *EndEvent) RemoveExtensionElements(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ExtensionElements)
	if toDelete.GetUUID() == this.M_extensionElements.GetUUID() {
		this.M_extensionElements = nil
	}
}

/** ExtensionDefinitions **/
func (this *EndEvent) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *EndEvent) SetExtensionDefinitions(ref interface{}){
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
func (this *EndEvent) RemoveExtensionDefinitions(ref interface{}){
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
func (this *EndEvent) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *EndEvent) SetExtensionValues(ref interface{}){
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
func (this *EndEvent) RemoveExtensionValues(ref interface{}){
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
func (this *EndEvent) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *EndEvent) SetDocumentation(ref interface{}){
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
func (this *EndEvent) RemoveDocumentation(ref interface{}){
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
func (this *EndEvent) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *EndEvent) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Auditing **/
func (this *EndEvent) GetAuditing() *Auditing{
	return this.M_auditing
}

/** Init reference Auditing **/
func (this *EndEvent) SetAuditing(ref interface{}){
	this.NeedSave = true
	this.M_auditing = ref.(*Auditing)
}

/** Remove reference Auditing **/
func (this *EndEvent) RemoveAuditing(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_auditing.GetUUID() {
		this.M_auditing = nil
	}
}

/** Monitoring **/
func (this *EndEvent) GetMonitoring() *Monitoring{
	return this.M_monitoring
}

/** Init reference Monitoring **/
func (this *EndEvent) SetMonitoring(ref interface{}){
	this.NeedSave = true
	this.M_monitoring = ref.(*Monitoring)
}

/** Remove reference Monitoring **/
func (this *EndEvent) RemoveMonitoring(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_monitoring.GetUUID() {
		this.M_monitoring = nil
	}
}

/** CategoryValueRef **/
func (this *EndEvent) GetCategoryValueRef() []*CategoryValue{
	return this.m_categoryValueRef
}

/** Init reference CategoryValueRef **/
func (this *EndEvent) SetCategoryValueRef(ref interface{}){
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
func (this *EndEvent) RemoveCategoryValueRef(ref interface{}){
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
func (this *EndEvent) GetOutgoing() []*SequenceFlow{
	return this.m_outgoing
}

/** Init reference Outgoing **/
func (this *EndEvent) SetOutgoing(ref interface{}){
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
func (this *EndEvent) RemoveOutgoing(ref interface{}){
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
func (this *EndEvent) GetIncoming() []*SequenceFlow{
	return this.m_incoming
}

/** Init reference Incoming **/
func (this *EndEvent) SetIncoming(ref interface{}){
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
func (this *EndEvent) RemoveIncoming(ref interface{}){
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
func (this *EndEvent) GetLanes() []*Lane{
	return this.m_lanes
}

/** Init reference Lanes **/
func (this *EndEvent) SetLanes(ref interface{}){
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
func (this *EndEvent) RemoveLanes(ref interface{}){
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

/** IncomingConversationLinks **/
func (this *EndEvent) GetIncomingConversationLinks() []*ConversationLink{
	return this.m_incomingConversationLinks
}

/** Init reference IncomingConversationLinks **/
func (this *EndEvent) SetIncomingConversationLinks(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_incomingConversationLinks); i++ {
			if this.M_incomingConversationLinks[i] == refStr {
				return
			}
		}
		this.M_incomingConversationLinks = append(this.M_incomingConversationLinks, ref.(string))
	}else{
		this.RemoveIncomingConversationLinks(ref)
		this.m_incomingConversationLinks = append(this.m_incomingConversationLinks, ref.(*ConversationLink))
		this.M_incomingConversationLinks = append(this.M_incomingConversationLinks, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference IncomingConversationLinks **/
func (this *EndEvent) RemoveIncomingConversationLinks(ref interface{}){
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
func (this *EndEvent) GetOutgoingConversationLinks() []*ConversationLink{
	return this.m_outgoingConversationLinks
}

/** Init reference OutgoingConversationLinks **/
func (this *EndEvent) SetOutgoingConversationLinks(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_outgoingConversationLinks); i++ {
			if this.M_outgoingConversationLinks[i] == refStr {
				return
			}
		}
		this.M_outgoingConversationLinks = append(this.M_outgoingConversationLinks, ref.(string))
	}else{
		this.RemoveOutgoingConversationLinks(ref)
		this.m_outgoingConversationLinks = append(this.m_outgoingConversationLinks, ref.(*ConversationLink))
		this.M_outgoingConversationLinks = append(this.M_outgoingConversationLinks, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference OutgoingConversationLinks **/
func (this *EndEvent) RemoveOutgoingConversationLinks(ref interface{}){
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

/** Property **/
func (this *EndEvent) GetProperty() []*Property{
	return this.M_property
}

/** Init reference Property **/
func (this *EndEvent) SetProperty(ref interface{}){
	this.NeedSave = true
	isExist := false
	var propertys []*Property
	for i:=0; i<len(this.M_property); i++ {
		if this.M_property[i].GetUUID() != ref.(BaseElement).GetUUID() {
			propertys = append(propertys, this.M_property[i])
		} else {
			isExist = true
			propertys = append(propertys, ref.(*Property))
		}
	}
	if !isExist {
		propertys = append(propertys, ref.(*Property))
	}
	this.M_property = propertys
}

/** Remove reference Property **/
func (this *EndEvent) RemoveProperty(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	property_ := make([]*Property, 0)
	for i := 0; i < len(this.M_property); i++ {
		if toDelete.GetUUID() != this.M_property[i].GetUUID() {
			property_ = append(property_, this.M_property[i])
		}
	}
	this.M_property = property_
}

/** InputSet **/
func (this *EndEvent) GetInputSet() *InputSet{
	return this.M_inputSet
}

/** Init reference InputSet **/
func (this *EndEvent) SetInputSet(ref interface{}){
	this.NeedSave = true
	this.M_inputSet = ref.(*InputSet)
}

/** Remove reference InputSet **/
func (this *EndEvent) RemoveInputSet(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_inputSet.GetUUID() {
		this.M_inputSet = nil
	}
}

/** EventDefinitionRef **/
func (this *EndEvent) GetEventDefinitionRef() []EventDefinition{
	return this.m_eventDefinitionRef
}

/** Init reference EventDefinitionRef **/
func (this *EndEvent) SetEventDefinitionRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_eventDefinitionRef); i++ {
			if this.M_eventDefinitionRef[i] == refStr {
				return
			}
		}
		this.M_eventDefinitionRef = append(this.M_eventDefinitionRef, ref.(string))
	}else{
		this.RemoveEventDefinitionRef(ref)
		this.m_eventDefinitionRef = append(this.m_eventDefinitionRef, ref.(EventDefinition))
		this.M_eventDefinitionRef = append(this.M_eventDefinitionRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference EventDefinitionRef **/
func (this *EndEvent) RemoveEventDefinitionRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	eventDefinitionRef_ := make([]EventDefinition, 0)
	eventDefinitionRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_eventDefinitionRef); i++ {
		if toDelete.GetUUID() != this.m_eventDefinitionRef[i].(BaseElement).GetUUID() {
			eventDefinitionRef_ = append(eventDefinitionRef_, this.m_eventDefinitionRef[i])
			eventDefinitionRefUuid = append(eventDefinitionRefUuid, this.M_eventDefinitionRef[i])
		}
	}
	this.m_eventDefinitionRef = eventDefinitionRef_
	this.M_eventDefinitionRef = eventDefinitionRefUuid
}

/** DataInputAssociation **/
func (this *EndEvent) GetDataInputAssociation() []*DataInputAssociation{
	return this.M_dataInputAssociation
}

/** Init reference DataInputAssociation **/
func (this *EndEvent) SetDataInputAssociation(ref interface{}){
	this.NeedSave = true
	isExist := false
	var dataInputAssociations []*DataInputAssociation
	for i:=0; i<len(this.M_dataInputAssociation); i++ {
		if this.M_dataInputAssociation[i].GetUUID() != ref.(BaseElement).GetUUID() {
			dataInputAssociations = append(dataInputAssociations, this.M_dataInputAssociation[i])
		} else {
			isExist = true
			dataInputAssociations = append(dataInputAssociations, ref.(*DataInputAssociation))
		}
	}
	if !isExist {
		dataInputAssociations = append(dataInputAssociations, ref.(*DataInputAssociation))
	}
	this.M_dataInputAssociation = dataInputAssociations
}

/** Remove reference DataInputAssociation **/
func (this *EndEvent) RemoveDataInputAssociation(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataInputAssociation_ := make([]*DataInputAssociation, 0)
	for i := 0; i < len(this.M_dataInputAssociation); i++ {
		if toDelete.GetUUID() != this.M_dataInputAssociation[i].GetUUID() {
			dataInputAssociation_ = append(dataInputAssociation_, this.M_dataInputAssociation[i])
		}
	}
	this.M_dataInputAssociation = dataInputAssociation_
}

/** DataInput **/
func (this *EndEvent) GetDataInput() []*DataInput{
	return this.M_dataInput
}

/** Init reference DataInput **/
func (this *EndEvent) SetDataInput(ref interface{}){
	this.NeedSave = true
	isExist := false
	var dataInputs []*DataInput
	for i:=0; i<len(this.M_dataInput); i++ {
		if this.M_dataInput[i].GetUUID() != ref.(BaseElement).GetUUID() {
			dataInputs = append(dataInputs, this.M_dataInput[i])
		} else {
			isExist = true
			dataInputs = append(dataInputs, ref.(*DataInput))
		}
	}
	if !isExist {
		dataInputs = append(dataInputs, ref.(*DataInput))
	}
	this.M_dataInput = dataInputs
}

/** Remove reference DataInput **/
func (this *EndEvent) RemoveDataInput(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataInput_ := make([]*DataInput, 0)
	for i := 0; i < len(this.M_dataInput); i++ {
		if toDelete.GetUUID() != this.M_dataInput[i].GetUUID() {
			dataInput_ = append(dataInput_, this.M_dataInput[i])
		}
	}
	this.M_dataInput = dataInput_
}

/** EventDefinition **/
func (this *EndEvent) GetEventDefinition() []EventDefinition{
	return this.M_eventDefinition
}

/** Init reference EventDefinition **/
func (this *EndEvent) SetEventDefinition(ref interface{}){
	this.NeedSave = true
	isExist := false
	var eventDefinitions []EventDefinition
	for i:=0; i<len(this.M_eventDefinition); i++ {
		if this.M_eventDefinition[i].(BaseElement).GetUUID() != ref.(BaseElement).GetUUID() {
			eventDefinitions = append(eventDefinitions, this.M_eventDefinition[i])
		} else {
			isExist = true
			eventDefinitions = append(eventDefinitions, ref.(EventDefinition))
		}
	}
	if !isExist {
		eventDefinitions = append(eventDefinitions, ref.(EventDefinition))
	}
	this.M_eventDefinition = eventDefinitions
}

/** Remove reference EventDefinition **/
func (this *EndEvent) RemoveEventDefinition(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	eventDefinition_ := make([]EventDefinition, 0)
	for i := 0; i < len(this.M_eventDefinition); i++ {
		if toDelete.GetUUID() != this.M_eventDefinition[i].(BaseElement).GetUUID() {
			eventDefinition_ = append(eventDefinition_, this.M_eventDefinition[i])
		}
	}
	this.M_eventDefinition = eventDefinition_
}

/** Lane **/
func (this *EndEvent) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *EndEvent) SetLanePtr(ref interface{}){
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
func (this *EndEvent) RemoveLanePtr(ref interface{}){
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
func (this *EndEvent) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *EndEvent) SetOutgoingPtr(ref interface{}){
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
func (this *EndEvent) RemoveOutgoingPtr(ref interface{}){
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
func (this *EndEvent) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *EndEvent) SetIncomingPtr(ref interface{}){
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
func (this *EndEvent) RemoveIncomingPtr(ref interface{}){
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
func (this *EndEvent) GetContainerPtr() FlowElementsContainer{
	return this.m_containerPtr
}

/** Init reference Container **/
func (this *EndEvent) SetContainerPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_containerPtr = ref.(string)
	}else{
		this.m_containerPtr = ref.(FlowElementsContainer)
		this.M_containerPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Container **/
func (this *EndEvent) RemoveContainerPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_containerPtr.(BaseElement).GetUUID() {
		this.m_containerPtr = nil
		this.M_containerPtr = ""
	}
}

/** MessageFlow **/
func (this *EndEvent) GetMessageFlowPtr() []*MessageFlow{
	return this.m_messageFlowPtr
}

/** Init reference MessageFlow **/
func (this *EndEvent) SetMessageFlowPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_messageFlowPtr); i++ {
			if this.M_messageFlowPtr[i] == refStr {
				return
			}
		}
		this.M_messageFlowPtr = append(this.M_messageFlowPtr, ref.(string))
	}else{
		this.RemoveMessageFlowPtr(ref)
		this.m_messageFlowPtr = append(this.m_messageFlowPtr, ref.(*MessageFlow))
		this.M_messageFlowPtr = append(this.M_messageFlowPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference MessageFlow **/
func (this *EndEvent) RemoveMessageFlowPtr(ref interface{}){
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
