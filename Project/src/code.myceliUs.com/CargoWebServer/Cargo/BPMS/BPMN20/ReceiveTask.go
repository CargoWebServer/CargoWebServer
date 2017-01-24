// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type ReceiveTask struct {

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

	/** members of FlowElement **/
	M_name             string
	M_auditing         *Auditing
	M_monitoring       *Monitoring
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
	m_lanes    []*Lane
	/** If the ref is a string and not an object **/
	M_lanes []string

	/** members of Activity **/
	M_isForCompensation   bool
	M_loopCharacteristics LoopCharacteristics
	M_resourceRole        []ResourceRole
	m_default             *SequenceFlow
	/** If the ref is a string and not an object **/
	M_default           string
	M_property          []*Property
	M_ioSpecification   *InputOutputSpecification
	m_boundaryEventRefs []*BoundaryEvent
	/** If the ref is a string and not an object **/
	M_boundaryEventRefs     []string
	M_dataInputAssociation  []*DataInputAssociation
	M_dataOutputAssociation []*DataOutputAssociation
	M_startQuantity         int
	M_completionQuantity    int

	/** members of InteractionNode **/
	m_incomingConversationLinks []*ConversationLink
	/** If the ref is a string and not an object **/
	M_incomingConversationLinks []string
	m_outgoingConversationLinks []*ConversationLink
	/** If the ref is a string and not an object **/
	M_outgoingConversationLinks []string

	/** members of Task **/
	/** No members **/

	/** members of ReceiveTask **/
	M_implementation    Implementation
	M_implementationStr string
	M_instantiate       bool
	m_operationRef      *Operation
	/** If the ref is a string and not an object **/
	M_operationRef string
	m_messageRef   *Message
	/** If the ref is a string and not an object **/
	M_messageRef string

	/** Associations **/
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr  []string
	m_containerPtr FlowElementsContainer
	/** If the ref is a string and not an object **/
	M_containerPtr                 string
	m_compensateEventDefinitionPtr []*CompensateEventDefinition
	/** If the ref is a string and not an object **/
	M_compensateEventDefinitionPtr []string
	m_messageFlowPtr               []*MessageFlow
	/** If the ref is a string and not an object **/
	M_messageFlowPtr []string
}

/** Xml parser for ReceiveTask **/
type XsdReceiveTask struct {
	XMLName xml.Name `xml:"receiveTask"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	/** FlowElement **/
	M_auditing         *XsdAuditing   `xml:"auditing,omitempty"`
	M_monitoring       *XsdMonitoring `xml:"monitoring,omitempty"`
	M_categoryValueRef []string       `xml:"categoryValueRef"`
	M_name             string         `xml:"name,attr"`

	/** FlowNode **/
	M_incoming []string `xml:"incoming"`
	M_outgoing []string `xml:"outgoing"`

	/** Activity **/
	M_ioSpecification       *XsdInputOutputSpecification         `xml:"ioSpecification,omitempty"`
	M_property              []*XsdProperty                       `xml:"property,omitempty"`
	M_dataInputAssociation  []*XsdDataInputAssociation           `xml:"dataInputAssociation,omitempty"`
	M_dataOutputAssociation []*XsdDataOutputAssociation          `xml:"dataOutputAssociation,omitempty"`
	M_resourceRole_0        []*XsdHumanPerformer                 `xml:"humanPerformer,omitempty"`
	M_resourceRole_1        []*XsdPotentialOwner                 `xml:"potentialOwner,omitempty"`
	M_resourceRole_2        []*XsdPerformer                      `xml:"performer,omitempty"`
	M_resourceRole_3        []*XsdResourceRole                   `xml:"resourceRole,omitempty"`
	M_loopCharacteristics_0 *XsdMultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics,omitempty"`
	M_loopCharacteristics_1 *XsdStandardLoopCharacteristics      `xml:"standardLoopCharacteristics,omitempty"`

	M_isForCompensation  bool   `xml:"isForCompensation,attr"`
	M_startQuantity      int    `xml:"startQuantity,attr"`
	M_completionQuantity int    `xml:"completionQuantity,attr"`
	M_default            string `xml:"default,attr"`

	/** Task **/

	M_implementation string `xml:"implementation,attr"`
	M_instantiate    bool   `xml:"instantiate,attr"`
	M_messageRef     string `xml:"messageRef,attr"`
	M_operationRef   string `xml:"operationRef,attr"`
}

/** UUID **/
func (this *ReceiveTask) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *ReceiveTask) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *ReceiveTask) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *ReceiveTask) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *ReceiveTask) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *ReceiveTask) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *ReceiveTask) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *ReceiveTask) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *ReceiveTask) SetExtensionDefinitions(ref interface{}) {
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
func (this *ReceiveTask) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *ReceiveTask) SetExtensionValues(ref interface{}) {
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
func (this *ReceiveTask) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *ReceiveTask) SetDocumentation(ref interface{}) {
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
func (this *ReceiveTask) RemoveDocumentation(ref interface{}) {
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
func (this *ReceiveTask) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *ReceiveTask) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Auditing **/
func (this *ReceiveTask) GetAuditing() *Auditing {
	return this.M_auditing
}

/** Init reference Auditing **/
func (this *ReceiveTask) SetAuditing(ref interface{}) {
	this.NeedSave = true
	this.M_auditing = ref.(*Auditing)
}

/** Remove reference Auditing **/
func (this *ReceiveTask) RemoveAuditing(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_auditing.GetUUID() {
		this.M_auditing = nil
	}
}

/** Monitoring **/
func (this *ReceiveTask) GetMonitoring() *Monitoring {
	return this.M_monitoring
}

/** Init reference Monitoring **/
func (this *ReceiveTask) SetMonitoring(ref interface{}) {
	this.NeedSave = true
	this.M_monitoring = ref.(*Monitoring)
}

/** Remove reference Monitoring **/
func (this *ReceiveTask) RemoveMonitoring(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_monitoring.GetUUID() {
		this.M_monitoring = nil
	}
}

/** CategoryValueRef **/
func (this *ReceiveTask) GetCategoryValueRef() []*CategoryValue {
	return this.m_categoryValueRef
}

/** Init reference CategoryValueRef **/
func (this *ReceiveTask) SetCategoryValueRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_categoryValueRef); i++ {
			if this.M_categoryValueRef[i] == refStr {
				return
			}
		}
		this.M_categoryValueRef = append(this.M_categoryValueRef, ref.(string))
	} else {
		this.RemoveCategoryValueRef(ref)
		this.m_categoryValueRef = append(this.m_categoryValueRef, ref.(*CategoryValue))
		this.M_categoryValueRef = append(this.M_categoryValueRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CategoryValueRef **/
func (this *ReceiveTask) RemoveCategoryValueRef(ref interface{}) {
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
func (this *ReceiveTask) GetOutgoing() []*SequenceFlow {
	return this.m_outgoing
}

/** Init reference Outgoing **/
func (this *ReceiveTask) SetOutgoing(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_outgoing); i++ {
			if this.M_outgoing[i] == refStr {
				return
			}
		}
		this.M_outgoing = append(this.M_outgoing, ref.(string))
	} else {
		this.RemoveOutgoing(ref)
		this.m_outgoing = append(this.m_outgoing, ref.(*SequenceFlow))
		this.M_outgoing = append(this.M_outgoing, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Outgoing **/
func (this *ReceiveTask) RemoveOutgoing(ref interface{}) {
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
func (this *ReceiveTask) GetIncoming() []*SequenceFlow {
	return this.m_incoming
}

/** Init reference Incoming **/
func (this *ReceiveTask) SetIncoming(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_incoming); i++ {
			if this.M_incoming[i] == refStr {
				return
			}
		}
		this.M_incoming = append(this.M_incoming, ref.(string))
	} else {
		this.RemoveIncoming(ref)
		this.m_incoming = append(this.m_incoming, ref.(*SequenceFlow))
		this.M_incoming = append(this.M_incoming, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Incoming **/
func (this *ReceiveTask) RemoveIncoming(ref interface{}) {
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
func (this *ReceiveTask) GetLanes() []*Lane {
	return this.m_lanes
}

/** Init reference Lanes **/
func (this *ReceiveTask) SetLanes(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_lanes); i++ {
			if this.M_lanes[i] == refStr {
				return
			}
		}
		this.M_lanes = append(this.M_lanes, ref.(string))
	} else {
		this.RemoveLanes(ref)
		this.m_lanes = append(this.m_lanes, ref.(*Lane))
		this.M_lanes = append(this.M_lanes, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Lanes **/
func (this *ReceiveTask) RemoveLanes(ref interface{}) {
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

/** IsForCompensation **/
func (this *ReceiveTask) IsForCompensation() bool {
	return this.M_isForCompensation
}

/** Init reference IsForCompensation **/
func (this *ReceiveTask) SetIsForCompensation(ref interface{}) {
	this.NeedSave = true
	this.M_isForCompensation = ref.(bool)
}

/** Remove reference IsForCompensation **/

/** LoopCharacteristics **/
func (this *ReceiveTask) GetLoopCharacteristics() LoopCharacteristics {
	return this.M_loopCharacteristics
}

/** Init reference LoopCharacteristics **/
func (this *ReceiveTask) SetLoopCharacteristics(ref interface{}) {
	this.NeedSave = true
	this.M_loopCharacteristics = ref.(LoopCharacteristics)
}

/** Remove reference LoopCharacteristics **/
func (this *ReceiveTask) RemoveLoopCharacteristics(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_loopCharacteristics.(BaseElement).GetUUID() {
		this.M_loopCharacteristics = nil
	}
}

/** ResourceRole **/
func (this *ReceiveTask) GetResourceRole() []ResourceRole {
	return this.M_resourceRole
}

/** Init reference ResourceRole **/
func (this *ReceiveTask) SetResourceRole(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var resourceRoles []ResourceRole
	for i := 0; i < len(this.M_resourceRole); i++ {
		if this.M_resourceRole[i].(BaseElement).GetUUID() != ref.(BaseElement).GetUUID() {
			resourceRoles = append(resourceRoles, this.M_resourceRole[i])
		} else {
			isExist = true
			resourceRoles = append(resourceRoles, ref.(ResourceRole))
		}
	}
	if !isExist {
		resourceRoles = append(resourceRoles, ref.(ResourceRole))
	}
	this.M_resourceRole = resourceRoles
}

/** Remove reference ResourceRole **/
func (this *ReceiveTask) RemoveResourceRole(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	resourceRole_ := make([]ResourceRole, 0)
	for i := 0; i < len(this.M_resourceRole); i++ {
		if toDelete.GetUUID() != this.M_resourceRole[i].(BaseElement).GetUUID() {
			resourceRole_ = append(resourceRole_, this.M_resourceRole[i])
		}
	}
	this.M_resourceRole = resourceRole_
}

/** Default **/
func (this *ReceiveTask) GetDefault() *SequenceFlow {
	return this.m_default
}

/** Init reference Default **/
func (this *ReceiveTask) SetDefault(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_default = ref.(string)
	} else {
		this.m_default = ref.(*SequenceFlow)
		this.M_default = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Default **/
func (this *ReceiveTask) RemoveDefault(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_default.GetUUID() {
		this.m_default = nil
		this.M_default = ""
	}
}

/** Property **/
func (this *ReceiveTask) GetProperty() []*Property {
	return this.M_property
}

/** Init reference Property **/
func (this *ReceiveTask) SetProperty(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var propertys []*Property
	for i := 0; i < len(this.M_property); i++ {
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
func (this *ReceiveTask) RemoveProperty(ref interface{}) {
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

/** IoSpecification **/
func (this *ReceiveTask) GetIoSpecification() *InputOutputSpecification {
	return this.M_ioSpecification
}

/** Init reference IoSpecification **/
func (this *ReceiveTask) SetIoSpecification(ref interface{}) {
	this.NeedSave = true
	this.M_ioSpecification = ref.(*InputOutputSpecification)
}

/** Remove reference IoSpecification **/
func (this *ReceiveTask) RemoveIoSpecification(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_ioSpecification.GetUUID() {
		this.M_ioSpecification = nil
	}
}

/** BoundaryEventRefs **/
func (this *ReceiveTask) GetBoundaryEventRefs() []*BoundaryEvent {
	return this.m_boundaryEventRefs
}

/** Init reference BoundaryEventRefs **/
func (this *ReceiveTask) SetBoundaryEventRefs(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_boundaryEventRefs); i++ {
			if this.M_boundaryEventRefs[i] == refStr {
				return
			}
		}
		this.M_boundaryEventRefs = append(this.M_boundaryEventRefs, ref.(string))
	} else {
		this.RemoveBoundaryEventRefs(ref)
		this.m_boundaryEventRefs = append(this.m_boundaryEventRefs, ref.(*BoundaryEvent))
		this.M_boundaryEventRefs = append(this.M_boundaryEventRefs, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference BoundaryEventRefs **/
func (this *ReceiveTask) RemoveBoundaryEventRefs(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	boundaryEventRefs_ := make([]*BoundaryEvent, 0)
	boundaryEventRefsUuid := make([]string, 0)
	for i := 0; i < len(this.m_boundaryEventRefs); i++ {
		if toDelete.GetUUID() != this.m_boundaryEventRefs[i].GetUUID() {
			boundaryEventRefs_ = append(boundaryEventRefs_, this.m_boundaryEventRefs[i])
			boundaryEventRefsUuid = append(boundaryEventRefsUuid, this.M_boundaryEventRefs[i])
		}
	}
	this.m_boundaryEventRefs = boundaryEventRefs_
	this.M_boundaryEventRefs = boundaryEventRefsUuid
}

/** DataInputAssociation **/
func (this *ReceiveTask) GetDataInputAssociation() []*DataInputAssociation {
	return this.M_dataInputAssociation
}

/** Init reference DataInputAssociation **/
func (this *ReceiveTask) SetDataInputAssociation(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var dataInputAssociations []*DataInputAssociation
	for i := 0; i < len(this.M_dataInputAssociation); i++ {
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
func (this *ReceiveTask) RemoveDataInputAssociation(ref interface{}) {
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

/** DataOutputAssociation **/
func (this *ReceiveTask) GetDataOutputAssociation() []*DataOutputAssociation {
	return this.M_dataOutputAssociation
}

/** Init reference DataOutputAssociation **/
func (this *ReceiveTask) SetDataOutputAssociation(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var dataOutputAssociations []*DataOutputAssociation
	for i := 0; i < len(this.M_dataOutputAssociation); i++ {
		if this.M_dataOutputAssociation[i].GetUUID() != ref.(BaseElement).GetUUID() {
			dataOutputAssociations = append(dataOutputAssociations, this.M_dataOutputAssociation[i])
		} else {
			isExist = true
			dataOutputAssociations = append(dataOutputAssociations, ref.(*DataOutputAssociation))
		}
	}
	if !isExist {
		dataOutputAssociations = append(dataOutputAssociations, ref.(*DataOutputAssociation))
	}
	this.M_dataOutputAssociation = dataOutputAssociations
}

/** Remove reference DataOutputAssociation **/
func (this *ReceiveTask) RemoveDataOutputAssociation(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataOutputAssociation_ := make([]*DataOutputAssociation, 0)
	for i := 0; i < len(this.M_dataOutputAssociation); i++ {
		if toDelete.GetUUID() != this.M_dataOutputAssociation[i].GetUUID() {
			dataOutputAssociation_ = append(dataOutputAssociation_, this.M_dataOutputAssociation[i])
		}
	}
	this.M_dataOutputAssociation = dataOutputAssociation_
}

/** StartQuantity **/
func (this *ReceiveTask) GetStartQuantity() int {
	return this.M_startQuantity
}

/** Init reference StartQuantity **/
func (this *ReceiveTask) SetStartQuantity(ref interface{}) {
	this.NeedSave = true
	this.M_startQuantity = ref.(int)
}

/** Remove reference StartQuantity **/

/** CompletionQuantity **/
func (this *ReceiveTask) GetCompletionQuantity() int {
	return this.M_completionQuantity
}

/** Init reference CompletionQuantity **/
func (this *ReceiveTask) SetCompletionQuantity(ref interface{}) {
	this.NeedSave = true
	this.M_completionQuantity = ref.(int)
}

/** Remove reference CompletionQuantity **/

/** IncomingConversationLinks **/
func (this *ReceiveTask) GetIncomingConversationLinks() []*ConversationLink {
	return this.m_incomingConversationLinks
}

/** Init reference IncomingConversationLinks **/
func (this *ReceiveTask) SetIncomingConversationLinks(ref interface{}) {
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
func (this *ReceiveTask) RemoveIncomingConversationLinks(ref interface{}) {
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
func (this *ReceiveTask) GetOutgoingConversationLinks() []*ConversationLink {
	return this.m_outgoingConversationLinks
}

/** Init reference OutgoingConversationLinks **/
func (this *ReceiveTask) SetOutgoingConversationLinks(ref interface{}) {
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
func (this *ReceiveTask) RemoveOutgoingConversationLinks(ref interface{}) {
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

/** Implementation **/
func (this *ReceiveTask) GetImplementation() Implementation {
	return this.M_implementation
}

/** Init reference Implementation **/
func (this *ReceiveTask) SetImplementation(ref interface{}) {
	this.NeedSave = true
	this.M_implementation = ref.(Implementation)
}

/** Remove reference Implementation **/

/** ImplementationStr **/
func (this *ReceiveTask) GetImplementationStr() string {
	return this.M_implementationStr
}

/** Init reference ImplementationStr **/
func (this *ReceiveTask) SetImplementationStr(ref interface{}) {
	this.NeedSave = true
	this.M_implementationStr = ref.(string)
}

/** Remove reference ImplementationStr **/

/** Instantiate **/
func (this *ReceiveTask) GetInstantiate() bool {
	return this.M_instantiate
}

/** Init reference Instantiate **/
func (this *ReceiveTask) SetInstantiate(ref interface{}) {
	this.NeedSave = true
	this.M_instantiate = ref.(bool)
}

/** Remove reference Instantiate **/

/** OperationRef **/
func (this *ReceiveTask) GetOperationRef() *Operation {
	return this.m_operationRef
}

/** Init reference OperationRef **/
func (this *ReceiveTask) SetOperationRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_operationRef = ref.(string)
	} else {
		this.m_operationRef = ref.(*Operation)
		this.M_operationRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference OperationRef **/
func (this *ReceiveTask) RemoveOperationRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_operationRef.GetUUID() {
		this.m_operationRef = nil
		this.M_operationRef = ""
	}
}

/** MessageRef **/
func (this *ReceiveTask) GetMessageRef() *Message {
	return this.m_messageRef
}

/** Init reference MessageRef **/
func (this *ReceiveTask) SetMessageRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_messageRef = ref.(string)
	} else {
		this.m_messageRef = ref.(*Message)
		this.M_messageRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference MessageRef **/
func (this *ReceiveTask) RemoveMessageRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_messageRef.GetUUID() {
		this.m_messageRef = nil
		this.M_messageRef = ""
	}
}

/** Lane **/
func (this *ReceiveTask) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *ReceiveTask) SetLanePtr(ref interface{}) {
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
func (this *ReceiveTask) RemoveLanePtr(ref interface{}) {
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
func (this *ReceiveTask) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *ReceiveTask) SetOutgoingPtr(ref interface{}) {
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
func (this *ReceiveTask) RemoveOutgoingPtr(ref interface{}) {
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
func (this *ReceiveTask) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *ReceiveTask) SetIncomingPtr(ref interface{}) {
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
func (this *ReceiveTask) RemoveIncomingPtr(ref interface{}) {
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
func (this *ReceiveTask) GetContainerPtr() FlowElementsContainer {
	return this.m_containerPtr
}

/** Init reference Container **/
func (this *ReceiveTask) SetContainerPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_containerPtr = ref.(string)
	} else {
		this.m_containerPtr = ref.(FlowElementsContainer)
		this.M_containerPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Container **/
func (this *ReceiveTask) RemoveContainerPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_containerPtr.(BaseElement).GetUUID() {
		this.m_containerPtr = nil
		this.M_containerPtr = ""
	}
}

/** CompensateEventDefinition **/
func (this *ReceiveTask) GetCompensateEventDefinitionPtr() []*CompensateEventDefinition {
	return this.m_compensateEventDefinitionPtr
}

/** Init reference CompensateEventDefinition **/
func (this *ReceiveTask) SetCompensateEventDefinitionPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_compensateEventDefinitionPtr); i++ {
			if this.M_compensateEventDefinitionPtr[i] == refStr {
				return
			}
		}
		this.M_compensateEventDefinitionPtr = append(this.M_compensateEventDefinitionPtr, ref.(string))
	} else {
		this.RemoveCompensateEventDefinitionPtr(ref)
		this.m_compensateEventDefinitionPtr = append(this.m_compensateEventDefinitionPtr, ref.(*CompensateEventDefinition))
		this.M_compensateEventDefinitionPtr = append(this.M_compensateEventDefinitionPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CompensateEventDefinition **/
func (this *ReceiveTask) RemoveCompensateEventDefinitionPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	compensateEventDefinitionPtr_ := make([]*CompensateEventDefinition, 0)
	compensateEventDefinitionPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_compensateEventDefinitionPtr); i++ {
		if toDelete.GetUUID() != this.m_compensateEventDefinitionPtr[i].GetUUID() {
			compensateEventDefinitionPtr_ = append(compensateEventDefinitionPtr_, this.m_compensateEventDefinitionPtr[i])
			compensateEventDefinitionPtrUuid = append(compensateEventDefinitionPtrUuid, this.M_compensateEventDefinitionPtr[i])
		}
	}
	this.m_compensateEventDefinitionPtr = compensateEventDefinitionPtr_
	this.M_compensateEventDefinitionPtr = compensateEventDefinitionPtrUuid
}

/** MessageFlow **/
func (this *ReceiveTask) GetMessageFlowPtr() []*MessageFlow {
	return this.m_messageFlowPtr
}

/** Init reference MessageFlow **/
func (this *ReceiveTask) SetMessageFlowPtr(ref interface{}) {
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
func (this *ReceiveTask) RemoveMessageFlowPtr(ref interface{}) {
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
