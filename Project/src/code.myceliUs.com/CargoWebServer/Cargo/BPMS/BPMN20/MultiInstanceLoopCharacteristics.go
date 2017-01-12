package BPMN20

import(
"encoding/xml"
)

type MultiInstanceLoopCharacteristics struct{

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

	/** members of LoopCharacteristics **/
	/** No members **/

	/** members of MultiInstanceLoopCharacteristics **/
	M_isSequential bool
	M_behavior MultiInstanceFlowCondition
	M_loopCardinality *FormalExpression
	m_loopDataInputRef ItemAwareElement
	/** If the ref is a string and not an object **/
	M_loopDataInputRef string
	m_loopDataOutputRef ItemAwareElement
	/** If the ref is a string and not an object **/
	M_loopDataOutputRef string
	M_inputDataItem *DataInput
	M_outputDataItem *DataOutput
	M_completionCondition *FormalExpression
	M_complexBehaviorDefinition []*ComplexBehaviorDefinition
	m_oneBehaviorEventRef EventDefinition
	/** If the ref is a string and not an object **/
	M_oneBehaviorEventRef string
	m_noneBehaviorEventRef EventDefinition
	/** If the ref is a string and not an object **/
	M_noneBehaviorEventRef string


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
	m_activityPtr Activity
	/** If the ref is a string and not an object **/
	M_activityPtr string
}

/** Xml parser for MultiInstanceLoopCharacteristics **/
type XsdMultiInstanceLoopCharacteristics struct {
	XMLName xml.Name	`xml:"multiInstanceLoopCharacteristics"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** LoopCharacteristics **/


	M_loopCardinality	*XsdLoopCardinality	`xml:"loopCardinality,omitempty"`
	M_loopDataInputRef	*string	`xml:"loopDataInputRef"`
	M_loopDataOutputRef	*string	`xml:"loopDataOutputRef"`
	M_inputDataItem	*XsdInputDataItem	`xml:"inputDataItem,omitempty"`
	M_outputDataItem	*XsdOutputDataItem	`xml:"outputDataItem,omitempty"`
	M_complexBehaviorDefinition	[]*XsdComplexBehaviorDefinition	`xml:"complexBehaviorDefinition,omitempty"`
	M_completionCondition	*XsdCompletionCondition	`xml:"completionCondition,omitempty"`
	M_isSequential	bool	`xml:"isSequential,attr"`
	M_behavior	string	`xml:"behavior,attr"`
	M_oneBehaviorEventRef	string	`xml:"oneBehaviorEventRef,attr"`
	M_noneBehaviorEventRef	string	`xml:"noneBehaviorEventRef,attr"`

}
/** UUID **/
func (this *MultiInstanceLoopCharacteristics) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *MultiInstanceLoopCharacteristics) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *MultiInstanceLoopCharacteristics) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *MultiInstanceLoopCharacteristics) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *MultiInstanceLoopCharacteristics) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *MultiInstanceLoopCharacteristics) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *MultiInstanceLoopCharacteristics) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *MultiInstanceLoopCharacteristics) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *MultiInstanceLoopCharacteristics) SetExtensionDefinitions(ref interface{}){
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
func (this *MultiInstanceLoopCharacteristics) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *MultiInstanceLoopCharacteristics) SetExtensionValues(ref interface{}){
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
func (this *MultiInstanceLoopCharacteristics) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *MultiInstanceLoopCharacteristics) SetDocumentation(ref interface{}){
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
func (this *MultiInstanceLoopCharacteristics) RemoveDocumentation(ref interface{}){
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

/** IsSequential **/
func (this *MultiInstanceLoopCharacteristics) IsSequential() bool{
	return this.M_isSequential
}

/** Init reference IsSequential **/
func (this *MultiInstanceLoopCharacteristics) SetIsSequential(ref interface{}){
	this.NeedSave = true
	this.M_isSequential = ref.(bool)
}

/** Remove reference IsSequential **/

/** Behavior **/
func (this *MultiInstanceLoopCharacteristics) GetBehavior() MultiInstanceFlowCondition{
	return this.M_behavior
}

/** Init reference Behavior **/
func (this *MultiInstanceLoopCharacteristics) SetBehavior(ref interface{}){
	this.NeedSave = true
	this.M_behavior = ref.(MultiInstanceFlowCondition)
}

/** Remove reference Behavior **/

/** LoopCardinality **/
func (this *MultiInstanceLoopCharacteristics) GetLoopCardinality() *FormalExpression{
	return this.M_loopCardinality
}

/** Init reference LoopCardinality **/
func (this *MultiInstanceLoopCharacteristics) SetLoopCardinality(ref interface{}){
	this.NeedSave = true
	this.M_loopCardinality = ref.(*FormalExpression)
}

/** Remove reference LoopCardinality **/
func (this *MultiInstanceLoopCharacteristics) RemoveLoopCardinality(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_loopCardinality.GetUUID() {
		this.M_loopCardinality = nil
	}
}

/** LoopDataInputRef **/
func (this *MultiInstanceLoopCharacteristics) GetLoopDataInputRef() ItemAwareElement{
	return this.m_loopDataInputRef
}

/** Init reference LoopDataInputRef **/
func (this *MultiInstanceLoopCharacteristics) SetLoopDataInputRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_loopDataInputRef = ref.(string)
	}else{
		this.m_loopDataInputRef = ref.(ItemAwareElement)
		this.M_loopDataInputRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference LoopDataInputRef **/
func (this *MultiInstanceLoopCharacteristics) RemoveLoopDataInputRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_loopDataInputRef.(BaseElement).GetUUID() {
		this.m_loopDataInputRef = nil
		this.M_loopDataInputRef = ""
	}
}

/** LoopDataOutputRef **/
func (this *MultiInstanceLoopCharacteristics) GetLoopDataOutputRef() ItemAwareElement{
	return this.m_loopDataOutputRef
}

/** Init reference LoopDataOutputRef **/
func (this *MultiInstanceLoopCharacteristics) SetLoopDataOutputRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_loopDataOutputRef = ref.(string)
	}else{
		this.m_loopDataOutputRef = ref.(ItemAwareElement)
		this.M_loopDataOutputRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference LoopDataOutputRef **/
func (this *MultiInstanceLoopCharacteristics) RemoveLoopDataOutputRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_loopDataOutputRef.(BaseElement).GetUUID() {
		this.m_loopDataOutputRef = nil
		this.M_loopDataOutputRef = ""
	}
}

/** InputDataItem **/
func (this *MultiInstanceLoopCharacteristics) GetInputDataItem() *DataInput{
	return this.M_inputDataItem
}

/** Init reference InputDataItem **/
func (this *MultiInstanceLoopCharacteristics) SetInputDataItem(ref interface{}){
	this.NeedSave = true
	this.M_inputDataItem = ref.(*DataInput)
}

/** Remove reference InputDataItem **/
func (this *MultiInstanceLoopCharacteristics) RemoveInputDataItem(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_inputDataItem.GetUUID() {
		this.M_inputDataItem = nil
	}
}

/** OutputDataItem **/
func (this *MultiInstanceLoopCharacteristics) GetOutputDataItem() *DataOutput{
	return this.M_outputDataItem
}

/** Init reference OutputDataItem **/
func (this *MultiInstanceLoopCharacteristics) SetOutputDataItem(ref interface{}){
	this.NeedSave = true
	this.M_outputDataItem = ref.(*DataOutput)
}

/** Remove reference OutputDataItem **/
func (this *MultiInstanceLoopCharacteristics) RemoveOutputDataItem(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_outputDataItem.GetUUID() {
		this.M_outputDataItem = nil
	}
}

/** CompletionCondition **/
func (this *MultiInstanceLoopCharacteristics) GetCompletionCondition() *FormalExpression{
	return this.M_completionCondition
}

/** Init reference CompletionCondition **/
func (this *MultiInstanceLoopCharacteristics) SetCompletionCondition(ref interface{}){
	this.NeedSave = true
	this.M_completionCondition = ref.(*FormalExpression)
}

/** Remove reference CompletionCondition **/
func (this *MultiInstanceLoopCharacteristics) RemoveCompletionCondition(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_completionCondition.GetUUID() {
		this.M_completionCondition = nil
	}
}

/** ComplexBehaviorDefinition **/
func (this *MultiInstanceLoopCharacteristics) GetComplexBehaviorDefinition() []*ComplexBehaviorDefinition{
	return this.M_complexBehaviorDefinition
}

/** Init reference ComplexBehaviorDefinition **/
func (this *MultiInstanceLoopCharacteristics) SetComplexBehaviorDefinition(ref interface{}){
	this.NeedSave = true
	isExist := false
	var complexBehaviorDefinitions []*ComplexBehaviorDefinition
	for i:=0; i<len(this.M_complexBehaviorDefinition); i++ {
		if this.M_complexBehaviorDefinition[i].GetUUID() != ref.(BaseElement).GetUUID() {
			complexBehaviorDefinitions = append(complexBehaviorDefinitions, this.M_complexBehaviorDefinition[i])
		} else {
			isExist = true
			complexBehaviorDefinitions = append(complexBehaviorDefinitions, ref.(*ComplexBehaviorDefinition))
		}
	}
	if !isExist {
		complexBehaviorDefinitions = append(complexBehaviorDefinitions, ref.(*ComplexBehaviorDefinition))
	}
	this.M_complexBehaviorDefinition = complexBehaviorDefinitions
}

/** Remove reference ComplexBehaviorDefinition **/
func (this *MultiInstanceLoopCharacteristics) RemoveComplexBehaviorDefinition(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	complexBehaviorDefinition_ := make([]*ComplexBehaviorDefinition, 0)
	for i := 0; i < len(this.M_complexBehaviorDefinition); i++ {
		if toDelete.GetUUID() != this.M_complexBehaviorDefinition[i].GetUUID() {
			complexBehaviorDefinition_ = append(complexBehaviorDefinition_, this.M_complexBehaviorDefinition[i])
		}
	}
	this.M_complexBehaviorDefinition = complexBehaviorDefinition_
}

/** OneBehaviorEventRef **/
func (this *MultiInstanceLoopCharacteristics) GetOneBehaviorEventRef() EventDefinition{
	return this.m_oneBehaviorEventRef
}

/** Init reference OneBehaviorEventRef **/
func (this *MultiInstanceLoopCharacteristics) SetOneBehaviorEventRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_oneBehaviorEventRef = ref.(string)
	}else{
		this.m_oneBehaviorEventRef = ref.(EventDefinition)
		this.M_oneBehaviorEventRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference OneBehaviorEventRef **/
func (this *MultiInstanceLoopCharacteristics) RemoveOneBehaviorEventRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_oneBehaviorEventRef.(BaseElement).GetUUID() {
		this.m_oneBehaviorEventRef = nil
		this.M_oneBehaviorEventRef = ""
	}
}

/** NoneBehaviorEventRef **/
func (this *MultiInstanceLoopCharacteristics) GetNoneBehaviorEventRef() EventDefinition{
	return this.m_noneBehaviorEventRef
}

/** Init reference NoneBehaviorEventRef **/
func (this *MultiInstanceLoopCharacteristics) SetNoneBehaviorEventRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_noneBehaviorEventRef = ref.(string)
	}else{
		this.m_noneBehaviorEventRef = ref.(EventDefinition)
		this.M_noneBehaviorEventRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference NoneBehaviorEventRef **/
func (this *MultiInstanceLoopCharacteristics) RemoveNoneBehaviorEventRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_noneBehaviorEventRef.(BaseElement).GetUUID() {
		this.m_noneBehaviorEventRef = nil
		this.M_noneBehaviorEventRef = ""
	}
}

/** Lane **/
func (this *MultiInstanceLoopCharacteristics) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *MultiInstanceLoopCharacteristics) SetLanePtr(ref interface{}){
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
func (this *MultiInstanceLoopCharacteristics) RemoveLanePtr(ref interface{}){
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
func (this *MultiInstanceLoopCharacteristics) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *MultiInstanceLoopCharacteristics) SetOutgoingPtr(ref interface{}){
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
func (this *MultiInstanceLoopCharacteristics) RemoveOutgoingPtr(ref interface{}){
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
func (this *MultiInstanceLoopCharacteristics) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *MultiInstanceLoopCharacteristics) SetIncomingPtr(ref interface{}){
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
func (this *MultiInstanceLoopCharacteristics) RemoveIncomingPtr(ref interface{}){
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

/** Activity **/
func (this *MultiInstanceLoopCharacteristics) GetActivityPtr() Activity{
	return this.m_activityPtr
}

/** Init reference Activity **/
func (this *MultiInstanceLoopCharacteristics) SetActivityPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_activityPtr = ref.(string)
	}else{
		this.m_activityPtr = ref.(Activity)
		this.M_activityPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Activity **/
func (this *MultiInstanceLoopCharacteristics) RemoveActivityPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_activityPtr.(BaseElement).GetUUID() {
		this.m_activityPtr = nil
		this.M_activityPtr = ""
	}
}
