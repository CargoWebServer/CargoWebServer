// +build BPMN20

package BPMN20

import(
"encoding/xml"
)

type InputSet struct{

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

	/** members of InputSet **/
	M_name string
	m_dataInputRefs []*DataInput
	/** If the ref is a string and not an object **/
	M_dataInputRefs []string
	m_optionalInputRefs []*DataInput
	/** If the ref is a string and not an object **/
	M_optionalInputRefs []string
	m_whileExecutingInputRefs []*DataInput
	/** If the ref is a string and not an object **/
	M_whileExecutingInputRefs []string
	m_outputSetRefs []*OutputSet
	/** If the ref is a string and not an object **/
	M_outputSetRefs []string


	/** Associations **/
	m_throwEventPtr ThrowEvent
	/** If the ref is a string and not an object **/
	M_throwEventPtr string
	m_inputOutputSpecificationPtr *InputOutputSpecification
	/** If the ref is a string and not an object **/
	M_inputOutputSpecificationPtr string
	m_inputOutputBindingPtr []*InputOutputBinding
	/** If the ref is a string and not an object **/
	M_inputOutputBindingPtr []string
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
}

/** Xml parser for InputSet **/
type XsdInputSet struct {
	XMLName xml.Name	`xml:"inputSet"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	M_dataInputRefs	[]string	`xml:"dataInputRefs"`
	M_optionalInputRefs	[]string	`xml:"optionalInputRefs"`
	M_whileExecutingInputRefs	[]string	`xml:"whileExecutingInputRefs"`
	M_outputSetRefs	[]string	`xml:"outputSetRefs"`
	M_name	string	`xml:"name,attr"`

}
/** UUID **/
func (this *InputSet) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *InputSet) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *InputSet) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *InputSet) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *InputSet) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *InputSet) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *InputSet) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/
func (this *InputSet) RemoveExtensionElements(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ExtensionElements)
	if toDelete.GetUUID() == this.M_extensionElements.GetUUID() {
		this.M_extensionElements = nil
	}
}

/** ExtensionDefinitions **/
func (this *InputSet) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *InputSet) SetExtensionDefinitions(ref interface{}){
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
func (this *InputSet) RemoveExtensionDefinitions(ref interface{}){
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
func (this *InputSet) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *InputSet) SetExtensionValues(ref interface{}){
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
func (this *InputSet) RemoveExtensionValues(ref interface{}){
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
func (this *InputSet) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *InputSet) SetDocumentation(ref interface{}){
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
func (this *InputSet) RemoveDocumentation(ref interface{}){
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
func (this *InputSet) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *InputSet) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** DataInputRefs **/
func (this *InputSet) GetDataInputRefs() []*DataInput{
	return this.m_dataInputRefs
}

/** Init reference DataInputRefs **/
func (this *InputSet) SetDataInputRefs(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_dataInputRefs); i++ {
			if this.M_dataInputRefs[i] == refStr {
				return
			}
		}
		this.M_dataInputRefs = append(this.M_dataInputRefs, ref.(string))
	}else{
		this.RemoveDataInputRefs(ref)
		this.m_dataInputRefs = append(this.m_dataInputRefs, ref.(*DataInput))
		this.M_dataInputRefs = append(this.M_dataInputRefs, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference DataInputRefs **/
func (this *InputSet) RemoveDataInputRefs(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataInputRefs_ := make([]*DataInput, 0)
	dataInputRefsUuid := make([]string, 0)
	for i := 0; i < len(this.m_dataInputRefs); i++ {
		if toDelete.GetUUID() != this.m_dataInputRefs[i].GetUUID() {
			dataInputRefs_ = append(dataInputRefs_, this.m_dataInputRefs[i])
			dataInputRefsUuid = append(dataInputRefsUuid, this.M_dataInputRefs[i])
		}
	}
	this.m_dataInputRefs = dataInputRefs_
	this.M_dataInputRefs = dataInputRefsUuid
}

/** OptionalInputRefs **/
func (this *InputSet) GetOptionalInputRefs() []*DataInput{
	return this.m_optionalInputRefs
}

/** Init reference OptionalInputRefs **/
func (this *InputSet) SetOptionalInputRefs(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_optionalInputRefs); i++ {
			if this.M_optionalInputRefs[i] == refStr {
				return
			}
		}
		this.M_optionalInputRefs = append(this.M_optionalInputRefs, ref.(string))
	}else{
		this.RemoveOptionalInputRefs(ref)
		this.m_optionalInputRefs = append(this.m_optionalInputRefs, ref.(*DataInput))
		this.M_optionalInputRefs = append(this.M_optionalInputRefs, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference OptionalInputRefs **/
func (this *InputSet) RemoveOptionalInputRefs(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	optionalInputRefs_ := make([]*DataInput, 0)
	optionalInputRefsUuid := make([]string, 0)
	for i := 0; i < len(this.m_optionalInputRefs); i++ {
		if toDelete.GetUUID() != this.m_optionalInputRefs[i].GetUUID() {
			optionalInputRefs_ = append(optionalInputRefs_, this.m_optionalInputRefs[i])
			optionalInputRefsUuid = append(optionalInputRefsUuid, this.M_optionalInputRefs[i])
		}
	}
	this.m_optionalInputRefs = optionalInputRefs_
	this.M_optionalInputRefs = optionalInputRefsUuid
}

/** WhileExecutingInputRefs **/
func (this *InputSet) GetWhileExecutingInputRefs() []*DataInput{
	return this.m_whileExecutingInputRefs
}

/** Init reference WhileExecutingInputRefs **/
func (this *InputSet) SetWhileExecutingInputRefs(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_whileExecutingInputRefs); i++ {
			if this.M_whileExecutingInputRefs[i] == refStr {
				return
			}
		}
		this.M_whileExecutingInputRefs = append(this.M_whileExecutingInputRefs, ref.(string))
	}else{
		this.RemoveWhileExecutingInputRefs(ref)
		this.m_whileExecutingInputRefs = append(this.m_whileExecutingInputRefs, ref.(*DataInput))
		this.M_whileExecutingInputRefs = append(this.M_whileExecutingInputRefs, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference WhileExecutingInputRefs **/
func (this *InputSet) RemoveWhileExecutingInputRefs(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	whileExecutingInputRefs_ := make([]*DataInput, 0)
	whileExecutingInputRefsUuid := make([]string, 0)
	for i := 0; i < len(this.m_whileExecutingInputRefs); i++ {
		if toDelete.GetUUID() != this.m_whileExecutingInputRefs[i].GetUUID() {
			whileExecutingInputRefs_ = append(whileExecutingInputRefs_, this.m_whileExecutingInputRefs[i])
			whileExecutingInputRefsUuid = append(whileExecutingInputRefsUuid, this.M_whileExecutingInputRefs[i])
		}
	}
	this.m_whileExecutingInputRefs = whileExecutingInputRefs_
	this.M_whileExecutingInputRefs = whileExecutingInputRefsUuid
}

/** OutputSetRefs **/
func (this *InputSet) GetOutputSetRefs() []*OutputSet{
	return this.m_outputSetRefs
}

/** Init reference OutputSetRefs **/
func (this *InputSet) SetOutputSetRefs(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_outputSetRefs); i++ {
			if this.M_outputSetRefs[i] == refStr {
				return
			}
		}
		this.M_outputSetRefs = append(this.M_outputSetRefs, ref.(string))
	}else{
		this.RemoveOutputSetRefs(ref)
		this.m_outputSetRefs = append(this.m_outputSetRefs, ref.(*OutputSet))
		this.M_outputSetRefs = append(this.M_outputSetRefs, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference OutputSetRefs **/
func (this *InputSet) RemoveOutputSetRefs(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	outputSetRefs_ := make([]*OutputSet, 0)
	outputSetRefsUuid := make([]string, 0)
	for i := 0; i < len(this.m_outputSetRefs); i++ {
		if toDelete.GetUUID() != this.m_outputSetRefs[i].GetUUID() {
			outputSetRefs_ = append(outputSetRefs_, this.m_outputSetRefs[i])
			outputSetRefsUuid = append(outputSetRefsUuid, this.M_outputSetRefs[i])
		}
	}
	this.m_outputSetRefs = outputSetRefs_
	this.M_outputSetRefs = outputSetRefsUuid
}

/** ThrowEvent **/
func (this *InputSet) GetThrowEventPtr() ThrowEvent{
	return this.m_throwEventPtr
}

/** Init reference ThrowEvent **/
func (this *InputSet) SetThrowEventPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_throwEventPtr = ref.(string)
	}else{
		this.m_throwEventPtr = ref.(ThrowEvent)
		this.M_throwEventPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ThrowEvent **/
func (this *InputSet) RemoveThrowEventPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_throwEventPtr.(BaseElement).GetUUID() {
		this.m_throwEventPtr = nil
		this.M_throwEventPtr = ""
	}
}

/** InputOutputSpecification **/
func (this *InputSet) GetInputOutputSpecificationPtr() *InputOutputSpecification{
	return this.m_inputOutputSpecificationPtr
}

/** Init reference InputOutputSpecification **/
func (this *InputSet) SetInputOutputSpecificationPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_inputOutputSpecificationPtr = ref.(string)
	}else{
		this.m_inputOutputSpecificationPtr = ref.(*InputOutputSpecification)
		this.M_inputOutputSpecificationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference InputOutputSpecification **/
func (this *InputSet) RemoveInputOutputSpecificationPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_inputOutputSpecificationPtr.GetUUID() {
		this.m_inputOutputSpecificationPtr = nil
		this.M_inputOutputSpecificationPtr = ""
	}
}

/** InputOutputBinding **/
func (this *InputSet) GetInputOutputBindingPtr() []*InputOutputBinding{
	return this.m_inputOutputBindingPtr
}

/** Init reference InputOutputBinding **/
func (this *InputSet) SetInputOutputBindingPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_inputOutputBindingPtr); i++ {
			if this.M_inputOutputBindingPtr[i] == refStr {
				return
			}
		}
		this.M_inputOutputBindingPtr = append(this.M_inputOutputBindingPtr, ref.(string))
	}else{
		this.RemoveInputOutputBindingPtr(ref)
		this.m_inputOutputBindingPtr = append(this.m_inputOutputBindingPtr, ref.(*InputOutputBinding))
		this.M_inputOutputBindingPtr = append(this.M_inputOutputBindingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference InputOutputBinding **/
func (this *InputSet) RemoveInputOutputBindingPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	inputOutputBindingPtr_ := make([]*InputOutputBinding, 0)
	inputOutputBindingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_inputOutputBindingPtr); i++ {
		if toDelete.GetUUID() != this.m_inputOutputBindingPtr[i].GetUUID() {
			inputOutputBindingPtr_ = append(inputOutputBindingPtr_, this.m_inputOutputBindingPtr[i])
			inputOutputBindingPtrUuid = append(inputOutputBindingPtrUuid, this.M_inputOutputBindingPtr[i])
		}
	}
	this.m_inputOutputBindingPtr = inputOutputBindingPtr_
	this.M_inputOutputBindingPtr = inputOutputBindingPtrUuid
}

/** Lane **/
func (this *InputSet) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *InputSet) SetLanePtr(ref interface{}){
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
func (this *InputSet) RemoveLanePtr(ref interface{}){
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
func (this *InputSet) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *InputSet) SetOutgoingPtr(ref interface{}){
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
func (this *InputSet) RemoveOutgoingPtr(ref interface{}){
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
func (this *InputSet) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *InputSet) SetIncomingPtr(ref interface{}){
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
func (this *InputSet) RemoveIncomingPtr(ref interface{}){
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
