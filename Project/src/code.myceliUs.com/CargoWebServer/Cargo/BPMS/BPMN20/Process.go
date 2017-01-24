// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Process struct {

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

	/** members of CallableElement **/
	M_name                  string
	M_ioSpecification       *InputOutputSpecification
	m_supportedInterfaceRef []*Interface
	/** If the ref is a string and not an object **/
	M_supportedInterfaceRef []string
	M_ioBinding             []*InputOutputBinding

	/** members of Process **/
	M_processType ProcessType
	M_isClosed    bool
	M_auditing    *Auditing
	M_monitoring  *Monitoring
	M_property    []*Property
	m_supports    []*Process
	/** If the ref is a string and not an object **/
	M_supports                     []string
	m_definitionalCollaborationRef Collaboration
	/** If the ref is a string and not an object **/
	M_definitionalCollaborationRef string
	M_isExecutable                 bool
	M_resourceRole                 []ResourceRole
	M_artifact                     []Artifact
	M_correlationSubscription      []*CorrelationSubscription

	/** Associations **/
	m_processPtr []*Process
	/** If the ref is a string and not an object **/
	M_processPtr     []string
	m_participantPtr []*Participant
	/** If the ref is a string and not an object **/
	M_participantPtr []string
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
	M_definitionsPtr  string
	m_callActivityPtr []*CallActivity
	/** If the ref is a string and not an object **/
	M_callActivityPtr []string
}

/** Xml parser for Process **/
type XsdProcess struct {
	XMLName xml.Name `xml:"process"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	/** RootElement **/

	/** CallableElement **/
	M_supportedInterfaceRef []string                     `xml:"supportedInterfaceRef"`
	M_ioSpecification       *XsdInputOutputSpecification `xml:"ioSpecification,omitempty"`
	M_ioBinding             []*XsdInputOutputBinding     `xml:"ioBinding,omitempty"`
	M_name                  string                       `xml:"name,attr"`

	M_auditing       *XsdAuditing                 `xml:"auditing,omitempty"`
	M_monitoring     *XsdMonitoring               `xml:"monitoring,omitempty"`
	M_property       []*XsdProperty               `xml:"property,omitempty"`
	M_laneSet        []*XsdLaneSet                `xml:"laneSet,omitempty"`
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

	M_artifact_0 []*XsdAssociation    `xml:"association,omitempty"`
	M_artifact_1 []*XsdGroup          `xml:"group,omitempty"`
	M_artifact_2 []*XsdTextAnnotation `xml:"textAnnotation,omitempty"`

	M_resourceRole_0               []*XsdHumanPerformer          `xml:"humanPerformer,omitempty"`
	M_resourceRole_1               []*XsdPotentialOwner          `xml:"potentialOwner,omitempty"`
	M_resourceRole_2               []*XsdPerformer               `xml:"performer,omitempty"`
	M_resourceRole_3               []*XsdResourceRole            `xml:"resourceRole,omitempty"`
	M_correlationSubscription      []*XsdCorrelationSubscription `xml:"correlationSubscription,omitempty"`
	M_supports                     []string                      `xml:"supports"`
	M_processType                  string                        `xml:"processType,attr"`
	M_isClosed                     bool                          `xml:"isClosed,attr"`
	M_isExecutable                 bool                          `xml:"isExecutable,attr"`
	M_definitionalCollaborationRef string                        `xml:"definitionalCollaborationRef,attr"`
}

/** UUID **/
func (this *Process) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Process) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Process) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Process) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Process) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Process) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Process) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Process) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Process) SetExtensionDefinitions(ref interface{}) {
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
func (this *Process) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Process) SetExtensionValues(ref interface{}) {
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
func (this *Process) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Process) SetDocumentation(ref interface{}) {
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
func (this *Process) RemoveDocumentation(ref interface{}) {
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
func (this *Process) GetFlowElement() []FlowElement {
	return this.M_flowElement
}

/** Init reference FlowElement **/
func (this *Process) SetFlowElement(ref interface{}) {
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
func (this *Process) RemoveFlowElement(ref interface{}) {
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
func (this *Process) GetLaneSet() []*LaneSet {
	return this.M_laneSet
}

/** Init reference LaneSet **/
func (this *Process) SetLaneSet(ref interface{}) {
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
func (this *Process) RemoveLaneSet(ref interface{}) {
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
func (this *Process) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Process) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** IoSpecification **/
func (this *Process) GetIoSpecification() *InputOutputSpecification {
	return this.M_ioSpecification
}

/** Init reference IoSpecification **/
func (this *Process) SetIoSpecification(ref interface{}) {
	this.NeedSave = true
	this.M_ioSpecification = ref.(*InputOutputSpecification)
}

/** Remove reference IoSpecification **/
func (this *Process) RemoveIoSpecification(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_ioSpecification.GetUUID() {
		this.M_ioSpecification = nil
	}
}

/** SupportedInterfaceRef **/
func (this *Process) GetSupportedInterfaceRef() []*Interface {
	return this.m_supportedInterfaceRef
}

/** Init reference SupportedInterfaceRef **/
func (this *Process) SetSupportedInterfaceRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_supportedInterfaceRef); i++ {
			if this.M_supportedInterfaceRef[i] == refStr {
				return
			}
		}
		this.M_supportedInterfaceRef = append(this.M_supportedInterfaceRef, ref.(string))
	} else {
		this.RemoveSupportedInterfaceRef(ref)
		this.m_supportedInterfaceRef = append(this.m_supportedInterfaceRef, ref.(*Interface))
		this.M_supportedInterfaceRef = append(this.M_supportedInterfaceRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference SupportedInterfaceRef **/
func (this *Process) RemoveSupportedInterfaceRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	supportedInterfaceRef_ := make([]*Interface, 0)
	supportedInterfaceRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_supportedInterfaceRef); i++ {
		if toDelete.GetUUID() != this.m_supportedInterfaceRef[i].GetUUID() {
			supportedInterfaceRef_ = append(supportedInterfaceRef_, this.m_supportedInterfaceRef[i])
			supportedInterfaceRefUuid = append(supportedInterfaceRefUuid, this.M_supportedInterfaceRef[i])
		}
	}
	this.m_supportedInterfaceRef = supportedInterfaceRef_
	this.M_supportedInterfaceRef = supportedInterfaceRefUuid
}

/** IoBinding **/
func (this *Process) GetIoBinding() []*InputOutputBinding {
	return this.M_ioBinding
}

/** Init reference IoBinding **/
func (this *Process) SetIoBinding(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var ioBindings []*InputOutputBinding
	for i := 0; i < len(this.M_ioBinding); i++ {
		if this.M_ioBinding[i].GetUUID() != ref.(BaseElement).GetUUID() {
			ioBindings = append(ioBindings, this.M_ioBinding[i])
		} else {
			isExist = true
			ioBindings = append(ioBindings, ref.(*InputOutputBinding))
		}
	}
	if !isExist {
		ioBindings = append(ioBindings, ref.(*InputOutputBinding))
	}
	this.M_ioBinding = ioBindings
}

/** Remove reference IoBinding **/
func (this *Process) RemoveIoBinding(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	ioBinding_ := make([]*InputOutputBinding, 0)
	for i := 0; i < len(this.M_ioBinding); i++ {
		if toDelete.GetUUID() != this.M_ioBinding[i].GetUUID() {
			ioBinding_ = append(ioBinding_, this.M_ioBinding[i])
		}
	}
	this.M_ioBinding = ioBinding_
}

/** ProcessType **/
func (this *Process) GetProcessType() ProcessType {
	return this.M_processType
}

/** Init reference ProcessType **/
func (this *Process) SetProcessType(ref interface{}) {
	this.NeedSave = true
	this.M_processType = ref.(ProcessType)
}

/** Remove reference ProcessType **/

/** IsClosed **/
func (this *Process) IsClosed() bool {
	return this.M_isClosed
}

/** Init reference IsClosed **/
func (this *Process) SetIsClosed(ref interface{}) {
	this.NeedSave = true
	this.M_isClosed = ref.(bool)
}

/** Remove reference IsClosed **/

/** Auditing **/
func (this *Process) GetAuditing() *Auditing {
	return this.M_auditing
}

/** Init reference Auditing **/
func (this *Process) SetAuditing(ref interface{}) {
	this.NeedSave = true
	this.M_auditing = ref.(*Auditing)
}

/** Remove reference Auditing **/
func (this *Process) RemoveAuditing(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_auditing.GetUUID() {
		this.M_auditing = nil
	}
}

/** Monitoring **/
func (this *Process) GetMonitoring() *Monitoring {
	return this.M_monitoring
}

/** Init reference Monitoring **/
func (this *Process) SetMonitoring(ref interface{}) {
	this.NeedSave = true
	this.M_monitoring = ref.(*Monitoring)
}

/** Remove reference Monitoring **/
func (this *Process) RemoveMonitoring(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_monitoring.GetUUID() {
		this.M_monitoring = nil
	}
}

/** Property **/
func (this *Process) GetProperty() []*Property {
	return this.M_property
}

/** Init reference Property **/
func (this *Process) SetProperty(ref interface{}) {
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
func (this *Process) RemoveProperty(ref interface{}) {
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

/** Supports **/
func (this *Process) GetSupports() []*Process {
	return this.m_supports
}

/** Init reference Supports **/
func (this *Process) SetSupports(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_supports); i++ {
			if this.M_supports[i] == refStr {
				return
			}
		}
		this.M_supports = append(this.M_supports, ref.(string))
	} else {
		this.RemoveSupports(ref)
		this.m_supports = append(this.m_supports, ref.(*Process))
		this.M_supports = append(this.M_supports, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Supports **/
func (this *Process) RemoveSupports(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	supports_ := make([]*Process, 0)
	supportsUuid := make([]string, 0)
	for i := 0; i < len(this.m_supports); i++ {
		if toDelete.GetUUID() != this.m_supports[i].GetUUID() {
			supports_ = append(supports_, this.m_supports[i])
			supportsUuid = append(supportsUuid, this.M_supports[i])
		}
	}
	this.m_supports = supports_
	this.M_supports = supportsUuid
}

/** DefinitionalCollaborationRef **/
func (this *Process) GetDefinitionalCollaborationRef() Collaboration {
	return this.m_definitionalCollaborationRef
}

/** Init reference DefinitionalCollaborationRef **/
func (this *Process) SetDefinitionalCollaborationRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionalCollaborationRef = ref.(string)
	} else {
		this.m_definitionalCollaborationRef = ref.(Collaboration)
		this.M_definitionalCollaborationRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference DefinitionalCollaborationRef **/
func (this *Process) RemoveDefinitionalCollaborationRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionalCollaborationRef.(BaseElement).GetUUID() {
		this.m_definitionalCollaborationRef = nil
		this.M_definitionalCollaborationRef = ""
	}
}

/** IsExecutable **/
func (this *Process) IsExecutable() bool {
	return this.M_isExecutable
}

/** Init reference IsExecutable **/
func (this *Process) SetIsExecutable(ref interface{}) {
	this.NeedSave = true
	this.M_isExecutable = ref.(bool)
}

/** Remove reference IsExecutable **/

/** ResourceRole **/
func (this *Process) GetResourceRole() []ResourceRole {
	return this.M_resourceRole
}

/** Init reference ResourceRole **/
func (this *Process) SetResourceRole(ref interface{}) {
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
func (this *Process) RemoveResourceRole(ref interface{}) {
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

/** Artifact **/
func (this *Process) GetArtifact() []Artifact {
	return this.M_artifact
}

/** Init reference Artifact **/
func (this *Process) SetArtifact(ref interface{}) {
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
func (this *Process) RemoveArtifact(ref interface{}) {
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

/** CorrelationSubscription **/
func (this *Process) GetCorrelationSubscription() []*CorrelationSubscription {
	return this.M_correlationSubscription
}

/** Init reference CorrelationSubscription **/
func (this *Process) SetCorrelationSubscription(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var correlationSubscriptions []*CorrelationSubscription
	for i := 0; i < len(this.M_correlationSubscription); i++ {
		if this.M_correlationSubscription[i].GetUUID() != ref.(BaseElement).GetUUID() {
			correlationSubscriptions = append(correlationSubscriptions, this.M_correlationSubscription[i])
		} else {
			isExist = true
			correlationSubscriptions = append(correlationSubscriptions, ref.(*CorrelationSubscription))
		}
	}
	if !isExist {
		correlationSubscriptions = append(correlationSubscriptions, ref.(*CorrelationSubscription))
	}
	this.M_correlationSubscription = correlationSubscriptions
}

/** Remove reference CorrelationSubscription **/
func (this *Process) RemoveCorrelationSubscription(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	correlationSubscription_ := make([]*CorrelationSubscription, 0)
	for i := 0; i < len(this.M_correlationSubscription); i++ {
		if toDelete.GetUUID() != this.M_correlationSubscription[i].GetUUID() {
			correlationSubscription_ = append(correlationSubscription_, this.M_correlationSubscription[i])
		}
	}
	this.M_correlationSubscription = correlationSubscription_
}

/** Process **/
func (this *Process) GetProcessPtr() []*Process {
	return this.m_processPtr
}

/** Init reference Process **/
func (this *Process) SetProcessPtr(ref interface{}) {
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
func (this *Process) RemoveProcessPtr(ref interface{}) {
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

/** Participant **/
func (this *Process) GetParticipantPtr() []*Participant {
	return this.m_participantPtr
}

/** Init reference Participant **/
func (this *Process) SetParticipantPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_participantPtr); i++ {
			if this.M_participantPtr[i] == refStr {
				return
			}
		}
		this.M_participantPtr = append(this.M_participantPtr, ref.(string))
	} else {
		this.RemoveParticipantPtr(ref)
		this.m_participantPtr = append(this.m_participantPtr, ref.(*Participant))
		this.M_participantPtr = append(this.M_participantPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Participant **/
func (this *Process) RemoveParticipantPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	participantPtr_ := make([]*Participant, 0)
	participantPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_participantPtr); i++ {
		if toDelete.GetUUID() != this.m_participantPtr[i].GetUUID() {
			participantPtr_ = append(participantPtr_, this.m_participantPtr[i])
			participantPtrUuid = append(participantPtrUuid, this.M_participantPtr[i])
		}
	}
	this.m_participantPtr = participantPtr_
	this.M_participantPtr = participantPtrUuid
}

/** Lane **/
func (this *Process) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Process) SetLanePtr(ref interface{}) {
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
func (this *Process) RemoveLanePtr(ref interface{}) {
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
func (this *Process) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Process) SetOutgoingPtr(ref interface{}) {
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
func (this *Process) RemoveOutgoingPtr(ref interface{}) {
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
func (this *Process) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Process) SetIncomingPtr(ref interface{}) {
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
func (this *Process) RemoveIncomingPtr(ref interface{}) {
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
func (this *Process) GetDefinitionsPtr() *Definitions {
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *Process) SetDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	} else {
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *Process) RemoveDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}

/** CallActivity **/
func (this *Process) GetCallActivityPtr() []*CallActivity {
	return this.m_callActivityPtr
}

/** Init reference CallActivity **/
func (this *Process) SetCallActivityPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_callActivityPtr); i++ {
			if this.M_callActivityPtr[i] == refStr {
				return
			}
		}
		this.M_callActivityPtr = append(this.M_callActivityPtr, ref.(string))
	} else {
		this.RemoveCallActivityPtr(ref)
		this.m_callActivityPtr = append(this.m_callActivityPtr, ref.(*CallActivity))
		this.M_callActivityPtr = append(this.M_callActivityPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CallActivity **/
func (this *Process) RemoveCallActivityPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	callActivityPtr_ := make([]*CallActivity, 0)
	callActivityPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_callActivityPtr); i++ {
		if toDelete.GetUUID() != this.m_callActivityPtr[i].GetUUID() {
			callActivityPtr_ = append(callActivityPtr_, this.m_callActivityPtr[i])
			callActivityPtrUuid = append(callActivityPtrUuid, this.M_callActivityPtr[i])
		}
	}
	this.m_callActivityPtr = callActivityPtr_
	this.M_callActivityPtr = callActivityPtrUuid
}
