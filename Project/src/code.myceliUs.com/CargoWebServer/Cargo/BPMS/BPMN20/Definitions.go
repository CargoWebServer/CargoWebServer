// +build BPMN
package BPMN20

import (
	"encoding/xml"

	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/BPMNDI"
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/DI"
)

type Definitions struct {

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

	/** members of Definitions **/
	M_name               string
	M_targetNamespace    string
	M_expressionLanguage string
	M_typeLanguage       string
	M_import             []*Import
	M_extension          []*Extension
	M_relationship       []*Relationship
	M_rootElement        []RootElement
	M_BPMNDiagram        []*BPMNDI.BPMNDiagram
	M_exporter           string
	M_exporterVersion    string

	/** Associations **/
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
}

/** Xml parser for Definitions **/
type XsdDefinitions struct {
	XMLName xml.Name `xml:"definitions"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_import         []*XsdImport                     `xml:"import,omitempty"`
	M_extension      []*XsdExtension                  `xml:"extension,omitempty"`
	M_rootElement_0  []*XsdCategory                   `xml:"category,omitempty"`
	M_rootElement_1  []*XsdGlobalChoreographyTask     `xml:"globalChoreographyTask,omitempty"`
	M_rootElement_2  []*XsdChoreography               `xml:"choreography,omitempty"`
	M_rootElement_3  []*XsdGlobalConversation         `xml:"globalConversation,omitempty"`
	M_rootElement_4  []*XsdCollaboration              `xml:"collaboration,omitempty"`
	M_rootElement_5  []*XsdCorrelationProperty        `xml:"correlationProperty,omitempty"`
	M_rootElement_6  []*XsdDataStore                  `xml:"dataStore,omitempty"`
	M_rootElement_7  []*XsdEndPoint                   `xml:"endPoint,omitempty"`
	M_rootElement_8  []*XsdError                      `xml:"error,omitempty"`
	M_rootElement_9  []*XsdEscalation                 `xml:"escalation,omitempty"`
	M_rootElement_10 []*XsdCancelEventDefinition      `xml:"cancelEventDefinition,omitempty"`
	M_rootElement_11 []*XsdCompensateEventDefinition  `xml:"compensateEventDefinition,omitempty"`
	M_rootElement_12 []*XsdConditionalEventDefinition `xml:"conditionalEventDefinition,omitempty"`
	M_rootElement_13 []*XsdErrorEventDefinition       `xml:"errorEventDefinition,omitempty"`
	M_rootElement_14 []*XsdEscalationEventDefinition  `xml:"escalationEventDefinition,omitempty"`
	M_rootElement_15 []*XsdLinkEventDefinition        `xml:"linkEventDefinition,omitempty"`
	M_rootElement_16 []*XsdMessageEventDefinition     `xml:"messageEventDefinition,omitempty"`
	M_rootElement_17 []*XsdSignalEventDefinition      `xml:"signalEventDefinition,omitempty"`
	M_rootElement_18 []*XsdTerminateEventDefinition   `xml:"terminateEventDefinition,omitempty"`
	M_rootElement_19 []*XsdTimerEventDefinition       `xml:"timerEventDefinition,omitempty"`
	M_rootElement_20 []*XsdGlobalBusinessRuleTask     `xml:"globalBusinessRuleTask,omitempty"`
	M_rootElement_21 []*XsdGlobalManualTask           `xml:"globalManualTask,omitempty"`
	M_rootElement_22 []*XsdGlobalScriptTask           `xml:"globalScriptTask,omitempty"`
	M_rootElement_23 []*XsdGlobalTask                 `xml:"globalTask,omitempty"`
	M_rootElement_24 []*XsdGlobalUserTask             `xml:"globalUserTask,omitempty"`
	M_rootElement_25 []*XsdInterface                  `xml:"interface,omitempty"`
	M_rootElement_26 []*XsdItemDefinition             `xml:"itemDefinition,omitempty"`
	M_rootElement_27 []*XsdMessage                    `xml:"message,omitempty"`
	M_rootElement_28 []*XsdPartnerEntity              `xml:"partnerEntity,omitempty"`
	M_rootElement_29 []*XsdPartnerRole                `xml:"partnerRole,omitempty"`
	M_rootElement_30 []*XsdProcess                    `xml:"process,omitempty"`
	M_rootElement_31 []*XsdResource                   `xml:"resource,omitempty"`
	M_rootElement_32 []*XsdSignal                     `xml:"signal,omitempty"`

	M_BPMNDiagram        []*BPMNDI.XsdBPMNDiagram `xml:"BPMNDiagram,omitempty"`
	M_relationship       []*XsdRelationship       `xml:"relationship,omitempty"`
	M_name               string                   `xml:"name,attr"`
	M_targetNamespace    string                   `xml:"targetNamespace,attr"`
	M_expressionLanguage string                   `xml:"expressionLanguage,attr"`
	M_typeLanguage       string                   `xml:"typeLanguage,attr"`
	M_exporter           string                   `xml:"exporter,attr"`
	M_exporterVersion    string                   `xml:"exporterVersion,attr"`
}

/** UUID **/
func (this *Definitions) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Definitions) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Definitions) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Definitions) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Definitions) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Definitions) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Definitions) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Definitions) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Definitions) SetExtensionDefinitions(ref interface{}) {
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
func (this *Definitions) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Definitions) SetExtensionValues(ref interface{}) {
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
func (this *Definitions) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Definitions) SetDocumentation(ref interface{}) {
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
func (this *Definitions) RemoveDocumentation(ref interface{}) {
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
func (this *Definitions) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Definitions) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** TargetNamespace **/
func (this *Definitions) GetTargetNamespace() string {
	return this.M_targetNamespace
}

/** Init reference TargetNamespace **/
func (this *Definitions) SetTargetNamespace(ref interface{}) {
	this.NeedSave = true
	this.M_targetNamespace = ref.(string)
}

/** Remove reference TargetNamespace **/

/** ExpressionLanguage **/
func (this *Definitions) GetExpressionLanguage() string {
	return this.M_expressionLanguage
}

/** Init reference ExpressionLanguage **/
func (this *Definitions) SetExpressionLanguage(ref interface{}) {
	this.NeedSave = true
	this.M_expressionLanguage = ref.(string)
}

/** Remove reference ExpressionLanguage **/

/** TypeLanguage **/
func (this *Definitions) GetTypeLanguage() string {
	return this.M_typeLanguage
}

/** Init reference TypeLanguage **/
func (this *Definitions) SetTypeLanguage(ref interface{}) {
	this.NeedSave = true
	this.M_typeLanguage = ref.(string)
}

/** Remove reference TypeLanguage **/

/** Import **/
func (this *Definitions) GetImport() []*Import {
	return this.M_import
}

/** Init reference Import **/
func (this *Definitions) SetImport(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var imports []*Import
	for i := 0; i < len(this.M_import); i++ {
		if this.M_import[i].GetUUID() != ref.(*Import).GetUUID() {
			imports = append(imports, this.M_import[i])
		} else {
			isExist = true
			imports = append(imports, ref.(*Import))
		}
	}
	if !isExist {
		imports = append(imports, ref.(*Import))
	}
	this.M_import = imports
}

/** Remove reference Import **/

/** Extension **/
func (this *Definitions) GetExtension() []*Extension {
	return this.M_extension
}

/** Init reference Extension **/
func (this *Definitions) SetExtension(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var extensions []*Extension
	for i := 0; i < len(this.M_extension); i++ {
		if this.M_extension[i].GetUUID() != ref.(*Extension).GetUUID() {
			extensions = append(extensions, this.M_extension[i])
		} else {
			isExist = true
			extensions = append(extensions, ref.(*Extension))
		}
	}
	if !isExist {
		extensions = append(extensions, ref.(*Extension))
	}
	this.M_extension = extensions
}

/** Remove reference Extension **/

/** Relationship **/
func (this *Definitions) GetRelationship() []*Relationship {
	return this.M_relationship
}

/** Init reference Relationship **/
func (this *Definitions) SetRelationship(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var relationships []*Relationship
	for i := 0; i < len(this.M_relationship); i++ {
		if this.M_relationship[i].GetUUID() != ref.(BaseElement).GetUUID() {
			relationships = append(relationships, this.M_relationship[i])
		} else {
			isExist = true
			relationships = append(relationships, ref.(*Relationship))
		}
	}
	if !isExist {
		relationships = append(relationships, ref.(*Relationship))
	}
	this.M_relationship = relationships
}

/** Remove reference Relationship **/
func (this *Definitions) RemoveRelationship(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	relationship_ := make([]*Relationship, 0)
	for i := 0; i < len(this.M_relationship); i++ {
		if toDelete.GetUUID() != this.M_relationship[i].GetUUID() {
			relationship_ = append(relationship_, this.M_relationship[i])
		}
	}
	this.M_relationship = relationship_
}

/** RootElement **/
func (this *Definitions) GetRootElement() []RootElement {
	return this.M_rootElement
}

/** Init reference RootElement **/
func (this *Definitions) SetRootElement(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var rootElements []RootElement
	for i := 0; i < len(this.M_rootElement); i++ {
		if this.M_rootElement[i].(BaseElement).GetUUID() != ref.(BaseElement).GetUUID() {
			rootElements = append(rootElements, this.M_rootElement[i])
		} else {
			isExist = true
			rootElements = append(rootElements, ref.(RootElement))
		}
	}
	if !isExist {
		rootElements = append(rootElements, ref.(RootElement))
	}
	this.M_rootElement = rootElements
}

/** Remove reference RootElement **/
func (this *Definitions) RemoveRootElement(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	rootElement_ := make([]RootElement, 0)
	for i := 0; i < len(this.M_rootElement); i++ {
		if toDelete.GetUUID() != this.M_rootElement[i].(BaseElement).GetUUID() {
			rootElement_ = append(rootElement_, this.M_rootElement[i])
		}
	}
	this.M_rootElement = rootElement_
}

/** BPMNDiagram **/
func (this *Definitions) GetBPMNDiagram() []*BPMNDI.BPMNDiagram {
	return this.M_BPMNDiagram
}

/** Init reference BPMNDiagram **/
func (this *Definitions) SetBPMNDiagram(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var BPMNDiagrams []*BPMNDI.BPMNDiagram
	for i := 0; i < len(this.M_BPMNDiagram); i++ {
		if this.M_BPMNDiagram[i].GetUUID() != ref.(DI.Diagram).GetUUID() {
			BPMNDiagrams = append(BPMNDiagrams, this.M_BPMNDiagram[i])
		} else {
			isExist = true
			BPMNDiagrams = append(BPMNDiagrams, ref.(*BPMNDI.BPMNDiagram))
		}
	}
	if !isExist {
		BPMNDiagrams = append(BPMNDiagrams, ref.(*BPMNDI.BPMNDiagram))
	}
	this.M_BPMNDiagram = BPMNDiagrams
}

/** Remove reference BPMNDiagram **/
func (this *Definitions) RemoveBPMNDiagram(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.Diagram)
	BPMNDiagram_ := make([]*BPMNDI.BPMNDiagram, 0)
	for i := 0; i < len(this.M_BPMNDiagram); i++ {
		if toDelete.GetUUID() != this.M_BPMNDiagram[i].GetUUID() {
			BPMNDiagram_ = append(BPMNDiagram_, this.M_BPMNDiagram[i])
		}
	}
	this.M_BPMNDiagram = BPMNDiagram_
}

/** Exporter **/
func (this *Definitions) GetExporter() string {
	return this.M_exporter
}

/** Init reference Exporter **/
func (this *Definitions) SetExporter(ref interface{}) {
	this.NeedSave = true
	this.M_exporter = ref.(string)
}

/** Remove reference Exporter **/

/** ExporterVersion **/
func (this *Definitions) GetExporterVersion() string {
	return this.M_exporterVersion
}

/** Init reference ExporterVersion **/
func (this *Definitions) SetExporterVersion(ref interface{}) {
	this.NeedSave = true
	this.M_exporterVersion = ref.(string)
}

/** Remove reference ExporterVersion **/

/** Lane **/
func (this *Definitions) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Definitions) SetLanePtr(ref interface{}) {
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
func (this *Definitions) RemoveLanePtr(ref interface{}) {
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
func (this *Definitions) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Definitions) SetOutgoingPtr(ref interface{}) {
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
func (this *Definitions) RemoveOutgoingPtr(ref interface{}) {
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
func (this *Definitions) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Definitions) SetIncomingPtr(ref interface{}) {
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
func (this *Definitions) RemoveIncomingPtr(ref interface{}) {
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
