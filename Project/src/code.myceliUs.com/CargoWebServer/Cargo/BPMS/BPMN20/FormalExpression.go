package BPMN20

import (
	"encoding/xml"
)

type FormalExpression struct {

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

	/** members of Expression **/
	/** No members **/

	/** members of FormalExpression **/
	M_language           string
	m_evaluatesToTypeRef *ItemDefinition
	/** If the ref is a string and not an object **/
	M_evaluatesToTypeRef string

	/** Associations **/
	m_complexGatewayPtr *ComplexGateway
	/** If the ref is a string and not an object **/
	M_complexGatewayPtr             string
	m_conditionalEventDefinitionPtr *ConditionalEventDefinition
	/** If the ref is a string and not an object **/
	M_conditionalEventDefinitionPtr string
	m_dataAssociationPtr            DataAssociation
	/** If the ref is a string and not an object **/
	M_dataAssociationPtr string
	m_assignmentPtr      *Assignment
	/** If the ref is a string and not an object **/
	M_assignmentPtr   string
	m_sequenceFlowPtr *SequenceFlow
	/** If the ref is a string and not an object **/
	M_sequenceFlowPtr   string
	m_correlationsetPtr *CorrelationPropertyRetrievalExpression
	/** If the ref is a string and not an object **/
	M_correlationsetPtr             string
	m_correlationPropertyBindingPtr *CorrelationPropertyBinding
	/** If the ref is a string and not an object **/
	M_correlationPropertyBindingPtr       string
	m_multiInstanceLoopCharacteristicsPtr *MultiInstanceLoopCharacteristics
	/** If the ref is a string and not an object **/
	M_multiInstanceLoopCharacteristicsPtr string
	m_standardLoopCharacteristicsPtr      *StandardLoopCharacteristics
	/** If the ref is a string and not an object **/
	M_standardLoopCharacteristicsPtr string
	m_adHocSubProcessPtr             *AdHocSubProcess
	/** If the ref is a string and not an object **/
	M_adHocSubProcessPtr           string
	m_complexBehaviorDefinitionPtr *ComplexBehaviorDefinition
	/** If the ref is a string and not an object **/
	M_complexBehaviorDefinitionPtr string
	m_lanePtr                      []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr             []string
	m_timerEventDefinitionPtr *TimerEventDefinition
	/** If the ref is a string and not an object **/
	M_timerEventDefinitionPtr     string
	m_resourceParameterBindingPtr *ResourceParameterBinding
	/** If the ref is a string and not an object **/
	M_resourceParameterBindingPtr     string
	m_resourceAssignmentExpressionPtr *ResourceAssignmentExpression
	/** If the ref is a string and not an object **/
	M_resourceAssignmentExpressionPtr string
}

/** Xml parser for FormalExpression **/
type XsdFormalExpression struct {
	XMLName xml.Name `xml:"formalExpression"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

/** Alias Xsd parser **/

type XsdCompletionCondition struct {
	XMLName xml.Name `xml:"completionCondition"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

type XsdLoopCondition struct {
	XMLName xml.Name `xml:"loopCondition"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

type XsdConditionExpression struct {
	XMLName xml.Name `xml:"conditionExpression"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

type XsdTransformation struct {
	XMLName xml.Name `xml:"transformation"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

type XsdMessagePath struct {
	XMLName xml.Name `xml:"messagePath"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

type XsdFrom struct {
	XMLName xml.Name `xml:"from"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

type XsdActivationCondition struct {
	XMLName xml.Name `xml:"activationCondition"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

type XsdTo struct {
	XMLName xml.Name `xml:"to"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

type XsdDataPath struct {
	XMLName xml.Name `xml:"dataPath"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

type XsdCondition struct {
	XMLName xml.Name `xml:"condition"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

type XsdLoopCardinality struct {
	XMLName xml.Name `xml:"loopCardinality"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	M_other             string                `xml:",innerxml"`

	/** Expression **/

	M_language           string `xml:"language,attr"`
	M_evaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

/** UUID **/
func (this *FormalExpression) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *FormalExpression) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *FormalExpression) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *FormalExpression) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *FormalExpression) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *FormalExpression) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *FormalExpression) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *FormalExpression) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *FormalExpression) SetExtensionDefinitions(ref interface{}) {
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
func (this *FormalExpression) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *FormalExpression) SetExtensionValues(ref interface{}) {
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
func (this *FormalExpression) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *FormalExpression) SetDocumentation(ref interface{}) {
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
func (this *FormalExpression) RemoveDocumentation(ref interface{}) {
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

/** Language **/
func (this *FormalExpression) GetLanguage() string {
	return this.M_language
}

/** Init reference Language **/
func (this *FormalExpression) SetLanguage(ref interface{}) {
	this.NeedSave = true
	this.M_language = ref.(string)
}

/** Remove reference Language **/

/** EvaluatesToTypeRef **/
func (this *FormalExpression) GetEvaluatesToTypeRef() *ItemDefinition {
	return this.m_evaluatesToTypeRef
}

/** Init reference EvaluatesToTypeRef **/
func (this *FormalExpression) SetEvaluatesToTypeRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_evaluatesToTypeRef = ref.(string)
	} else {
		this.m_evaluatesToTypeRef = ref.(*ItemDefinition)
		this.M_evaluatesToTypeRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference EvaluatesToTypeRef **/
func (this *FormalExpression) RemoveEvaluatesToTypeRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_evaluatesToTypeRef.GetUUID() {
		this.m_evaluatesToTypeRef = nil
		this.M_evaluatesToTypeRef = ""
	}
}

/** ComplexGateway **/
func (this *FormalExpression) GetComplexGatewayPtr() *ComplexGateway {
	return this.m_complexGatewayPtr
}

/** Init reference ComplexGateway **/
func (this *FormalExpression) SetComplexGatewayPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_complexGatewayPtr = ref.(string)
	} else {
		this.m_complexGatewayPtr = ref.(*ComplexGateway)
		this.M_complexGatewayPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ComplexGateway **/
func (this *FormalExpression) RemoveComplexGatewayPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_complexGatewayPtr.GetUUID() {
		this.m_complexGatewayPtr = nil
		this.M_complexGatewayPtr = ""
	}
}

/** ConditionalEventDefinition **/
func (this *FormalExpression) GetConditionalEventDefinitionPtr() *ConditionalEventDefinition {
	return this.m_conditionalEventDefinitionPtr
}

/** Init reference ConditionalEventDefinition **/
func (this *FormalExpression) SetConditionalEventDefinitionPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_conditionalEventDefinitionPtr = ref.(string)
	} else {
		this.m_conditionalEventDefinitionPtr = ref.(*ConditionalEventDefinition)
		this.M_conditionalEventDefinitionPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ConditionalEventDefinition **/
func (this *FormalExpression) RemoveConditionalEventDefinitionPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_conditionalEventDefinitionPtr.GetUUID() {
		this.m_conditionalEventDefinitionPtr = nil
		this.M_conditionalEventDefinitionPtr = ""
	}
}

/** DataAssociation **/
func (this *FormalExpression) GetDataAssociationPtr() DataAssociation {
	return this.m_dataAssociationPtr
}

/** Init reference DataAssociation **/
func (this *FormalExpression) SetDataAssociationPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_dataAssociationPtr = ref.(string)
	} else {
		this.m_dataAssociationPtr = ref.(DataAssociation)
		this.M_dataAssociationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference DataAssociation **/
func (this *FormalExpression) RemoveDataAssociationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_dataAssociationPtr.(BaseElement).GetUUID() {
		this.m_dataAssociationPtr = nil
		this.M_dataAssociationPtr = ""
	}
}

/** Assignment **/
func (this *FormalExpression) GetAssignmentPtr() *Assignment {
	return this.m_assignmentPtr
}

/** Init reference Assignment **/
func (this *FormalExpression) SetAssignmentPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_assignmentPtr = ref.(string)
	} else {
		this.m_assignmentPtr = ref.(*Assignment)
		this.M_assignmentPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Assignment **/
func (this *FormalExpression) RemoveAssignmentPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_assignmentPtr.GetUUID() {
		this.m_assignmentPtr = nil
		this.M_assignmentPtr = ""
	}
}

/** SequenceFlow **/
func (this *FormalExpression) GetSequenceFlowPtr() *SequenceFlow {
	return this.m_sequenceFlowPtr
}

/** Init reference SequenceFlow **/
func (this *FormalExpression) SetSequenceFlowPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_sequenceFlowPtr = ref.(string)
	} else {
		this.m_sequenceFlowPtr = ref.(*SequenceFlow)
		this.M_sequenceFlowPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference SequenceFlow **/
func (this *FormalExpression) RemoveSequenceFlowPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_sequenceFlowPtr.GetUUID() {
		this.m_sequenceFlowPtr = nil
		this.M_sequenceFlowPtr = ""
	}
}

/** Correlationset **/
func (this *FormalExpression) GetCorrelationsetPtr() *CorrelationPropertyRetrievalExpression {
	return this.m_correlationsetPtr
}

/** Init reference Correlationset **/
func (this *FormalExpression) SetCorrelationsetPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_correlationsetPtr = ref.(string)
	} else {
		this.m_correlationsetPtr = ref.(*CorrelationPropertyRetrievalExpression)
		this.M_correlationsetPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Correlationset **/
func (this *FormalExpression) RemoveCorrelationsetPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_correlationsetPtr.GetUUID() {
		this.m_correlationsetPtr = nil
		this.M_correlationsetPtr = ""
	}
}

/** CorrelationPropertyBinding **/
func (this *FormalExpression) GetCorrelationPropertyBindingPtr() *CorrelationPropertyBinding {
	return this.m_correlationPropertyBindingPtr
}

/** Init reference CorrelationPropertyBinding **/
func (this *FormalExpression) SetCorrelationPropertyBindingPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_correlationPropertyBindingPtr = ref.(string)
	} else {
		this.m_correlationPropertyBindingPtr = ref.(*CorrelationPropertyBinding)
		this.M_correlationPropertyBindingPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference CorrelationPropertyBinding **/
func (this *FormalExpression) RemoveCorrelationPropertyBindingPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_correlationPropertyBindingPtr.GetUUID() {
		this.m_correlationPropertyBindingPtr = nil
		this.M_correlationPropertyBindingPtr = ""
	}
}

/** MultiInstanceLoopCharacteristics **/
func (this *FormalExpression) GetMultiInstanceLoopCharacteristicsPtr() *MultiInstanceLoopCharacteristics {
	return this.m_multiInstanceLoopCharacteristicsPtr
}

/** Init reference MultiInstanceLoopCharacteristics **/
func (this *FormalExpression) SetMultiInstanceLoopCharacteristicsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(string)
	} else {
		this.m_multiInstanceLoopCharacteristicsPtr = ref.(*MultiInstanceLoopCharacteristics)
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference MultiInstanceLoopCharacteristics **/
func (this *FormalExpression) RemoveMultiInstanceLoopCharacteristicsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_multiInstanceLoopCharacteristicsPtr.GetUUID() {
		this.m_multiInstanceLoopCharacteristicsPtr = nil
		this.M_multiInstanceLoopCharacteristicsPtr = ""
	}
}

/** StandardLoopCharacteristics **/
func (this *FormalExpression) GetStandardLoopCharacteristicsPtr() *StandardLoopCharacteristics {
	return this.m_standardLoopCharacteristicsPtr
}

/** Init reference StandardLoopCharacteristics **/
func (this *FormalExpression) SetStandardLoopCharacteristicsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_standardLoopCharacteristicsPtr = ref.(string)
	} else {
		this.m_standardLoopCharacteristicsPtr = ref.(*StandardLoopCharacteristics)
		this.M_standardLoopCharacteristicsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference StandardLoopCharacteristics **/
func (this *FormalExpression) RemoveStandardLoopCharacteristicsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_standardLoopCharacteristicsPtr.GetUUID() {
		this.m_standardLoopCharacteristicsPtr = nil
		this.M_standardLoopCharacteristicsPtr = ""
	}
}

/** AdHocSubProcess **/
func (this *FormalExpression) GetAdHocSubProcessPtr() *AdHocSubProcess {
	return this.m_adHocSubProcessPtr
}

/** Init reference AdHocSubProcess **/
func (this *FormalExpression) SetAdHocSubProcessPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_adHocSubProcessPtr = ref.(string)
	} else {
		this.m_adHocSubProcessPtr = ref.(*AdHocSubProcess)
		this.M_adHocSubProcessPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference AdHocSubProcess **/
func (this *FormalExpression) RemoveAdHocSubProcessPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_adHocSubProcessPtr.GetUUID() {
		this.m_adHocSubProcessPtr = nil
		this.M_adHocSubProcessPtr = ""
	}
}

/** ComplexBehaviorDefinition **/
func (this *FormalExpression) GetComplexBehaviorDefinitionPtr() *ComplexBehaviorDefinition {
	return this.m_complexBehaviorDefinitionPtr
}

/** Init reference ComplexBehaviorDefinition **/
func (this *FormalExpression) SetComplexBehaviorDefinitionPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_complexBehaviorDefinitionPtr = ref.(string)
	} else {
		this.m_complexBehaviorDefinitionPtr = ref.(*ComplexBehaviorDefinition)
		this.M_complexBehaviorDefinitionPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ComplexBehaviorDefinition **/
func (this *FormalExpression) RemoveComplexBehaviorDefinitionPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_complexBehaviorDefinitionPtr.GetUUID() {
		this.m_complexBehaviorDefinitionPtr = nil
		this.M_complexBehaviorDefinitionPtr = ""
	}
}

/** Lane **/
func (this *FormalExpression) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *FormalExpression) SetLanePtr(ref interface{}) {
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
func (this *FormalExpression) RemoveLanePtr(ref interface{}) {
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
func (this *FormalExpression) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *FormalExpression) SetOutgoingPtr(ref interface{}) {
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
func (this *FormalExpression) RemoveOutgoingPtr(ref interface{}) {
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
func (this *FormalExpression) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *FormalExpression) SetIncomingPtr(ref interface{}) {
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
func (this *FormalExpression) RemoveIncomingPtr(ref interface{}) {
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

/** TimerEventDefinition **/
func (this *FormalExpression) GetTimerEventDefinitionPtr() *TimerEventDefinition {
	return this.m_timerEventDefinitionPtr
}

/** Init reference TimerEventDefinition **/
func (this *FormalExpression) SetTimerEventDefinitionPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_timerEventDefinitionPtr = ref.(string)
	} else {
		this.m_timerEventDefinitionPtr = ref.(*TimerEventDefinition)
		this.M_timerEventDefinitionPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference TimerEventDefinition **/
func (this *FormalExpression) RemoveTimerEventDefinitionPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_timerEventDefinitionPtr.GetUUID() {
		this.m_timerEventDefinitionPtr = nil
		this.M_timerEventDefinitionPtr = ""
	}
}

/** ResourceParameterBinding **/
func (this *FormalExpression) GetResourceParameterBindingPtr() *ResourceParameterBinding {
	return this.m_resourceParameterBindingPtr
}

/** Init reference ResourceParameterBinding **/
func (this *FormalExpression) SetResourceParameterBindingPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_resourceParameterBindingPtr = ref.(string)
	} else {
		this.m_resourceParameterBindingPtr = ref.(*ResourceParameterBinding)
		this.M_resourceParameterBindingPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ResourceParameterBinding **/
func (this *FormalExpression) RemoveResourceParameterBindingPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_resourceParameterBindingPtr.GetUUID() {
		this.m_resourceParameterBindingPtr = nil
		this.M_resourceParameterBindingPtr = ""
	}
}

/** ResourceAssignmentExpression **/
func (this *FormalExpression) GetResourceAssignmentExpressionPtr() *ResourceAssignmentExpression {
	return this.m_resourceAssignmentExpressionPtr
}

/** Init reference ResourceAssignmentExpression **/
func (this *FormalExpression) SetResourceAssignmentExpressionPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_resourceAssignmentExpressionPtr = ref.(string)
	} else {
		this.m_resourceAssignmentExpressionPtr = ref.(*ResourceAssignmentExpression)
		this.M_resourceAssignmentExpressionPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ResourceAssignmentExpression **/
func (this *FormalExpression) RemoveResourceAssignmentExpressionPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_resourceAssignmentExpressionPtr.GetUUID() {
		this.m_resourceAssignmentExpressionPtr = nil
		this.M_resourceAssignmentExpressionPtr = ""
	}
}
