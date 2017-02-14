// +build BPMN20

package BPMN20

import(
"encoding/xml"
)

type SequenceFlow struct{

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

	/** members of SequenceFlow **/
	M_isImmediate bool
	M_conditionExpression *FormalExpression
	m_sourceRef FlowNode
	/** If the ref is a string and not an object **/
	M_sourceRef string
	m_targetRef FlowNode
	/** If the ref is a string and not an object **/
	M_targetRef string


	/** Associations **/
	m_complexGatewayPtr []*ComplexGateway
	/** If the ref is a string and not an object **/
	M_complexGatewayPtr []string
	m_exclusiveGatewayPtr []*ExclusiveGateway
	/** If the ref is a string and not an object **/
	M_exclusiveGatewayPtr []string
	m_inclusiveGatewayPtr []*InclusiveGateway
	/** If the ref is a string and not an object **/
	M_inclusiveGatewayPtr []string
	m_activityPtr Activity
	/** If the ref is a string and not an object **/
	M_activityPtr string
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

/** Xml parser for SequenceFlow **/
type XsdSequenceFlow struct {
	XMLName xml.Name	`xml:"sequenceFlow"`
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


	M_conditionExpression	*XsdConditionExpression	`xml:"conditionExpression,omitempty"`
	M_sourceRef	string	`xml:"sourceRef,attr"`
	M_targetRef	string	`xml:"targetRef,attr"`
	M_isImmediate	bool	`xml:"isImmediate,attr"`

}
/** UUID **/
func (this *SequenceFlow) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *SequenceFlow) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *SequenceFlow) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *SequenceFlow) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *SequenceFlow) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *SequenceFlow) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *SequenceFlow) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/
func (this *SequenceFlow) RemoveExtensionElements(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ExtensionElements)
	if toDelete.GetUUID() == this.M_extensionElements.GetUUID() {
		this.M_extensionElements = nil
	}
}

/** ExtensionDefinitions **/
func (this *SequenceFlow) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *SequenceFlow) SetExtensionDefinitions(ref interface{}){
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
func (this *SequenceFlow) RemoveExtensionDefinitions(ref interface{}){
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
func (this *SequenceFlow) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *SequenceFlow) SetExtensionValues(ref interface{}){
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
func (this *SequenceFlow) RemoveExtensionValues(ref interface{}){
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
func (this *SequenceFlow) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *SequenceFlow) SetDocumentation(ref interface{}){
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
func (this *SequenceFlow) RemoveDocumentation(ref interface{}){
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
func (this *SequenceFlow) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *SequenceFlow) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Auditing **/
func (this *SequenceFlow) GetAuditing() *Auditing{
	return this.M_auditing
}

/** Init reference Auditing **/
func (this *SequenceFlow) SetAuditing(ref interface{}){
	this.NeedSave = true
	this.M_auditing = ref.(*Auditing)
}

/** Remove reference Auditing **/
func (this *SequenceFlow) RemoveAuditing(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_auditing.GetUUID() {
		this.M_auditing = nil
	}
}

/** Monitoring **/
func (this *SequenceFlow) GetMonitoring() *Monitoring{
	return this.M_monitoring
}

/** Init reference Monitoring **/
func (this *SequenceFlow) SetMonitoring(ref interface{}){
	this.NeedSave = true
	this.M_monitoring = ref.(*Monitoring)
}

/** Remove reference Monitoring **/
func (this *SequenceFlow) RemoveMonitoring(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_monitoring.GetUUID() {
		this.M_monitoring = nil
	}
}

/** CategoryValueRef **/
func (this *SequenceFlow) GetCategoryValueRef() []*CategoryValue{
	return this.m_categoryValueRef
}

/** Init reference CategoryValueRef **/
func (this *SequenceFlow) SetCategoryValueRef(ref interface{}){
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
func (this *SequenceFlow) RemoveCategoryValueRef(ref interface{}){
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

/** IsImmediate **/
func (this *SequenceFlow) IsImmediate() bool{
	return this.M_isImmediate
}

/** Init reference IsImmediate **/
func (this *SequenceFlow) SetIsImmediate(ref interface{}){
	this.NeedSave = true
	this.M_isImmediate = ref.(bool)
}

/** Remove reference IsImmediate **/

/** ConditionExpression **/
func (this *SequenceFlow) GetConditionExpression() *FormalExpression{
	return this.M_conditionExpression
}

/** Init reference ConditionExpression **/
func (this *SequenceFlow) SetConditionExpression(ref interface{}){
	this.NeedSave = true
	this.M_conditionExpression = ref.(*FormalExpression)
}

/** Remove reference ConditionExpression **/
func (this *SequenceFlow) RemoveConditionExpression(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_conditionExpression.GetUUID() {
		this.M_conditionExpression = nil
	}
}

/** SourceRef **/
func (this *SequenceFlow) GetSourceRef() FlowNode{
	return this.m_sourceRef
}

/** Init reference SourceRef **/
func (this *SequenceFlow) SetSourceRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_sourceRef = ref.(string)
	}else{
		this.m_sourceRef = ref.(FlowNode)
		this.M_sourceRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference SourceRef **/
func (this *SequenceFlow) RemoveSourceRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_sourceRef.(BaseElement).GetUUID() {
		this.m_sourceRef = nil
		this.M_sourceRef = ""
	}
}

/** TargetRef **/
func (this *SequenceFlow) GetTargetRef() FlowNode{
	return this.m_targetRef
}

/** Init reference TargetRef **/
func (this *SequenceFlow) SetTargetRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_targetRef = ref.(string)
	}else{
		this.m_targetRef = ref.(FlowNode)
		this.M_targetRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference TargetRef **/
func (this *SequenceFlow) RemoveTargetRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_targetRef.(BaseElement).GetUUID() {
		this.m_targetRef = nil
		this.M_targetRef = ""
	}
}

/** ComplexGateway **/
func (this *SequenceFlow) GetComplexGatewayPtr() []*ComplexGateway{
	return this.m_complexGatewayPtr
}

/** Init reference ComplexGateway **/
func (this *SequenceFlow) SetComplexGatewayPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_complexGatewayPtr); i++ {
			if this.M_complexGatewayPtr[i] == refStr {
				return
			}
		}
		this.M_complexGatewayPtr = append(this.M_complexGatewayPtr, ref.(string))
	}else{
		this.RemoveComplexGatewayPtr(ref)
		this.m_complexGatewayPtr = append(this.m_complexGatewayPtr, ref.(*ComplexGateway))
		this.M_complexGatewayPtr = append(this.M_complexGatewayPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ComplexGateway **/
func (this *SequenceFlow) RemoveComplexGatewayPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	complexGatewayPtr_ := make([]*ComplexGateway, 0)
	complexGatewayPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_complexGatewayPtr); i++ {
		if toDelete.GetUUID() != this.m_complexGatewayPtr[i].GetUUID() {
			complexGatewayPtr_ = append(complexGatewayPtr_, this.m_complexGatewayPtr[i])
			complexGatewayPtrUuid = append(complexGatewayPtrUuid, this.M_complexGatewayPtr[i])
		}
	}
	this.m_complexGatewayPtr = complexGatewayPtr_
	this.M_complexGatewayPtr = complexGatewayPtrUuid
}

/** ExclusiveGateway **/
func (this *SequenceFlow) GetExclusiveGatewayPtr() []*ExclusiveGateway{
	return this.m_exclusiveGatewayPtr
}

/** Init reference ExclusiveGateway **/
func (this *SequenceFlow) SetExclusiveGatewayPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_exclusiveGatewayPtr); i++ {
			if this.M_exclusiveGatewayPtr[i] == refStr {
				return
			}
		}
		this.M_exclusiveGatewayPtr = append(this.M_exclusiveGatewayPtr, ref.(string))
	}else{
		this.RemoveExclusiveGatewayPtr(ref)
		this.m_exclusiveGatewayPtr = append(this.m_exclusiveGatewayPtr, ref.(*ExclusiveGateway))
		this.M_exclusiveGatewayPtr = append(this.M_exclusiveGatewayPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ExclusiveGateway **/
func (this *SequenceFlow) RemoveExclusiveGatewayPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	exclusiveGatewayPtr_ := make([]*ExclusiveGateway, 0)
	exclusiveGatewayPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_exclusiveGatewayPtr); i++ {
		if toDelete.GetUUID() != this.m_exclusiveGatewayPtr[i].GetUUID() {
			exclusiveGatewayPtr_ = append(exclusiveGatewayPtr_, this.m_exclusiveGatewayPtr[i])
			exclusiveGatewayPtrUuid = append(exclusiveGatewayPtrUuid, this.M_exclusiveGatewayPtr[i])
		}
	}
	this.m_exclusiveGatewayPtr = exclusiveGatewayPtr_
	this.M_exclusiveGatewayPtr = exclusiveGatewayPtrUuid
}

/** InclusiveGateway **/
func (this *SequenceFlow) GetInclusiveGatewayPtr() []*InclusiveGateway{
	return this.m_inclusiveGatewayPtr
}

/** Init reference InclusiveGateway **/
func (this *SequenceFlow) SetInclusiveGatewayPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_inclusiveGatewayPtr); i++ {
			if this.M_inclusiveGatewayPtr[i] == refStr {
				return
			}
		}
		this.M_inclusiveGatewayPtr = append(this.M_inclusiveGatewayPtr, ref.(string))
	}else{
		this.RemoveInclusiveGatewayPtr(ref)
		this.m_inclusiveGatewayPtr = append(this.m_inclusiveGatewayPtr, ref.(*InclusiveGateway))
		this.M_inclusiveGatewayPtr = append(this.M_inclusiveGatewayPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference InclusiveGateway **/
func (this *SequenceFlow) RemoveInclusiveGatewayPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	inclusiveGatewayPtr_ := make([]*InclusiveGateway, 0)
	inclusiveGatewayPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_inclusiveGatewayPtr); i++ {
		if toDelete.GetUUID() != this.m_inclusiveGatewayPtr[i].GetUUID() {
			inclusiveGatewayPtr_ = append(inclusiveGatewayPtr_, this.m_inclusiveGatewayPtr[i])
			inclusiveGatewayPtrUuid = append(inclusiveGatewayPtrUuid, this.M_inclusiveGatewayPtr[i])
		}
	}
	this.m_inclusiveGatewayPtr = inclusiveGatewayPtr_
	this.M_inclusiveGatewayPtr = inclusiveGatewayPtrUuid
}

/** Activity **/
func (this *SequenceFlow) GetActivityPtr() Activity{
	return this.m_activityPtr
}

/** Init reference Activity **/
func (this *SequenceFlow) SetActivityPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_activityPtr = ref.(string)
	}else{
		this.m_activityPtr = ref.(Activity)
		this.M_activityPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Activity **/
func (this *SequenceFlow) RemoveActivityPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_activityPtr.(BaseElement).GetUUID() {
		this.m_activityPtr = nil
		this.M_activityPtr = ""
	}
}

/** Lane **/
func (this *SequenceFlow) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *SequenceFlow) SetLanePtr(ref interface{}){
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
func (this *SequenceFlow) RemoveLanePtr(ref interface{}){
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
func (this *SequenceFlow) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *SequenceFlow) SetOutgoingPtr(ref interface{}){
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
func (this *SequenceFlow) RemoveOutgoingPtr(ref interface{}){
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
func (this *SequenceFlow) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *SequenceFlow) SetIncomingPtr(ref interface{}){
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
func (this *SequenceFlow) RemoveIncomingPtr(ref interface{}){
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
func (this *SequenceFlow) GetContainerPtr() FlowElementsContainer{
	return this.m_containerPtr
}

/** Init reference Container **/
func (this *SequenceFlow) SetContainerPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_containerPtr = ref.(string)
	}else{
		this.m_containerPtr = ref.(FlowElementsContainer)
		this.M_containerPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Container **/
func (this *SequenceFlow) RemoveContainerPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_containerPtr.(BaseElement).GetUUID() {
		this.m_containerPtr = nil
		this.M_containerPtr = ""
	}
}
