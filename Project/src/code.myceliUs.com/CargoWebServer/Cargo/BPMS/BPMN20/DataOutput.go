package BPMN20

import(
"encoding/xml"
)

type DataOutput struct{

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

	/** members of ItemAwareElement **/
	m_itemSubjectRef *ItemDefinition
	/** If the ref is a string and not an object **/
	M_itemSubjectRef string
	M_dataState *DataState

	/** members of DataOutput **/
	M_name string
	M_isCollection bool
	m_outputSetRefs []*OutputSet
	/** If the ref is a string and not an object **/
	M_outputSetRefs []string
	m_outputSetWithOptional []*OutputSet
	/** If the ref is a string and not an object **/
	M_outputSetWithOptional []string
	m_outputSetWithWhileExecuting []*OutputSet
	/** If the ref is a string and not an object **/
	M_outputSetWithWhileExecuting []string


	/** Associations **/
	m_catchEventPtr CatchEvent
	/** If the ref is a string and not an object **/
	M_catchEventPtr string
	m_inputOutputSpecificationPtr *InputOutputSpecification
	/** If the ref is a string and not an object **/
	M_inputOutputSpecificationPtr string
	m_multiInstanceLoopCharacteristicsPtr *MultiInstanceLoopCharacteristics
	/** If the ref is a string and not an object **/
	M_multiInstanceLoopCharacteristicsPtr string
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
	m_dataAssociationPtr []DataAssociation
	/** If the ref is a string and not an object **/
	M_dataAssociationPtr []string
}

/** Xml parser for DataOutput **/
type XsdDataOutput struct {
	XMLName xml.Name	`xml:"dataOutput"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	M_dataState	*XsdDataState	`xml:"dataState,omitempty"`
	M_name	string	`xml:"name,attr"`
	M_itemSubjectRef	string	`xml:"itemSubjectRef,attr"`
	M_isCollection	bool	`xml:"isCollection,attr"`

}
/** Alias Xsd parser **/

 
type XsdOutputDataItem struct {
	XMLName xml.Name	`xml:"outputDataItem"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	M_dataState	*XsdDataState	`xml:"dataState,omitempty"`
	M_name	string	`xml:"name,attr"`
	M_itemSubjectRef	string	`xml:"itemSubjectRef,attr"`
	M_isCollection	bool	`xml:"isCollection,attr"`

}
/** UUID **/
func (this *DataOutput) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *DataOutput) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *DataOutput) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *DataOutput) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *DataOutput) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *DataOutput) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *DataOutput) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *DataOutput) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *DataOutput) SetExtensionDefinitions(ref interface{}){
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
func (this *DataOutput) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *DataOutput) SetExtensionValues(ref interface{}){
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
func (this *DataOutput) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *DataOutput) SetDocumentation(ref interface{}){
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
func (this *DataOutput) RemoveDocumentation(ref interface{}){
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

/** ItemSubjectRef **/
func (this *DataOutput) GetItemSubjectRef() *ItemDefinition{
	return this.m_itemSubjectRef
}

/** Init reference ItemSubjectRef **/
func (this *DataOutput) SetItemSubjectRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_itemSubjectRef = ref.(string)
	}else{
		this.m_itemSubjectRef = ref.(*ItemDefinition)
		this.M_itemSubjectRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ItemSubjectRef **/
func (this *DataOutput) RemoveItemSubjectRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_itemSubjectRef.GetUUID() {
		this.m_itemSubjectRef = nil
		this.M_itemSubjectRef = ""
	}
}

/** DataState **/
func (this *DataOutput) GetDataState() *DataState{
	return this.M_dataState
}

/** Init reference DataState **/
func (this *DataOutput) SetDataState(ref interface{}){
	this.NeedSave = true
	this.M_dataState = ref.(*DataState)
}

/** Remove reference DataState **/
func (this *DataOutput) RemoveDataState(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_dataState.GetUUID() {
		this.M_dataState = nil
	}
}

/** Name **/
func (this *DataOutput) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *DataOutput) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** IsCollection **/
func (this *DataOutput) IsCollection() bool{
	return this.M_isCollection
}

/** Init reference IsCollection **/
func (this *DataOutput) SetIsCollection(ref interface{}){
	this.NeedSave = true
	this.M_isCollection = ref.(bool)
}

/** Remove reference IsCollection **/

/** OutputSetRefs **/
func (this *DataOutput) GetOutputSetRefs() []*OutputSet{
	return this.m_outputSetRefs
}

/** Init reference OutputSetRefs **/
func (this *DataOutput) SetOutputSetRefs(ref interface{}){
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
func (this *DataOutput) RemoveOutputSetRefs(ref interface{}){
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

/** OutputSetWithOptional **/
func (this *DataOutput) GetOutputSetWithOptional() []*OutputSet{
	return this.m_outputSetWithOptional
}

/** Init reference OutputSetWithOptional **/
func (this *DataOutput) SetOutputSetWithOptional(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_outputSetWithOptional); i++ {
			if this.M_outputSetWithOptional[i] == refStr {
				return
			}
		}
		this.M_outputSetWithOptional = append(this.M_outputSetWithOptional, ref.(string))
	}else{
		this.RemoveOutputSetWithOptional(ref)
		this.m_outputSetWithOptional = append(this.m_outputSetWithOptional, ref.(*OutputSet))
		this.M_outputSetWithOptional = append(this.M_outputSetWithOptional, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference OutputSetWithOptional **/
func (this *DataOutput) RemoveOutputSetWithOptional(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	outputSetWithOptional_ := make([]*OutputSet, 0)
	outputSetWithOptionalUuid := make([]string, 0)
	for i := 0; i < len(this.m_outputSetWithOptional); i++ {
		if toDelete.GetUUID() != this.m_outputSetWithOptional[i].GetUUID() {
			outputSetWithOptional_ = append(outputSetWithOptional_, this.m_outputSetWithOptional[i])
			outputSetWithOptionalUuid = append(outputSetWithOptionalUuid, this.M_outputSetWithOptional[i])
		}
	}
	this.m_outputSetWithOptional = outputSetWithOptional_
	this.M_outputSetWithOptional = outputSetWithOptionalUuid
}

/** OutputSetWithWhileExecuting **/
func (this *DataOutput) GetOutputSetWithWhileExecuting() []*OutputSet{
	return this.m_outputSetWithWhileExecuting
}

/** Init reference OutputSetWithWhileExecuting **/
func (this *DataOutput) SetOutputSetWithWhileExecuting(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_outputSetWithWhileExecuting); i++ {
			if this.M_outputSetWithWhileExecuting[i] == refStr {
				return
			}
		}
		this.M_outputSetWithWhileExecuting = append(this.M_outputSetWithWhileExecuting, ref.(string))
	}else{
		this.RemoveOutputSetWithWhileExecuting(ref)
		this.m_outputSetWithWhileExecuting = append(this.m_outputSetWithWhileExecuting, ref.(*OutputSet))
		this.M_outputSetWithWhileExecuting = append(this.M_outputSetWithWhileExecuting, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference OutputSetWithWhileExecuting **/
func (this *DataOutput) RemoveOutputSetWithWhileExecuting(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	outputSetWithWhileExecuting_ := make([]*OutputSet, 0)
	outputSetWithWhileExecutingUuid := make([]string, 0)
	for i := 0; i < len(this.m_outputSetWithWhileExecuting); i++ {
		if toDelete.GetUUID() != this.m_outputSetWithWhileExecuting[i].GetUUID() {
			outputSetWithWhileExecuting_ = append(outputSetWithWhileExecuting_, this.m_outputSetWithWhileExecuting[i])
			outputSetWithWhileExecutingUuid = append(outputSetWithWhileExecutingUuid, this.M_outputSetWithWhileExecuting[i])
		}
	}
	this.m_outputSetWithWhileExecuting = outputSetWithWhileExecuting_
	this.M_outputSetWithWhileExecuting = outputSetWithWhileExecutingUuid
}

/** CatchEvent **/
func (this *DataOutput) GetCatchEventPtr() CatchEvent{
	return this.m_catchEventPtr
}

/** Init reference CatchEvent **/
func (this *DataOutput) SetCatchEventPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_catchEventPtr = ref.(string)
	}else{
		this.m_catchEventPtr = ref.(CatchEvent)
		this.M_catchEventPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference CatchEvent **/
func (this *DataOutput) RemoveCatchEventPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_catchEventPtr.(BaseElement).GetUUID() {
		this.m_catchEventPtr = nil
		this.M_catchEventPtr = ""
	}
}

/** InputOutputSpecification **/
func (this *DataOutput) GetInputOutputSpecificationPtr() *InputOutputSpecification{
	return this.m_inputOutputSpecificationPtr
}

/** Init reference InputOutputSpecification **/
func (this *DataOutput) SetInputOutputSpecificationPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_inputOutputSpecificationPtr = ref.(string)
	}else{
		this.m_inputOutputSpecificationPtr = ref.(*InputOutputSpecification)
		this.M_inputOutputSpecificationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference InputOutputSpecification **/
func (this *DataOutput) RemoveInputOutputSpecificationPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_inputOutputSpecificationPtr.GetUUID() {
		this.m_inputOutputSpecificationPtr = nil
		this.M_inputOutputSpecificationPtr = ""
	}
}

/** MultiInstanceLoopCharacteristics **/
func (this *DataOutput) GetMultiInstanceLoopCharacteristicsPtr() *MultiInstanceLoopCharacteristics{
	return this.m_multiInstanceLoopCharacteristicsPtr
}

/** Init reference MultiInstanceLoopCharacteristics **/
func (this *DataOutput) SetMultiInstanceLoopCharacteristicsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(string)
	}else{
		this.m_multiInstanceLoopCharacteristicsPtr = ref.(*MultiInstanceLoopCharacteristics)
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference MultiInstanceLoopCharacteristics **/
func (this *DataOutput) RemoveMultiInstanceLoopCharacteristicsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_multiInstanceLoopCharacteristicsPtr.GetUUID() {
		this.m_multiInstanceLoopCharacteristicsPtr = nil
		this.M_multiInstanceLoopCharacteristicsPtr = ""
	}
}

/** Lane **/
func (this *DataOutput) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *DataOutput) SetLanePtr(ref interface{}){
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
func (this *DataOutput) RemoveLanePtr(ref interface{}){
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
func (this *DataOutput) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *DataOutput) SetOutgoingPtr(ref interface{}){
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
func (this *DataOutput) RemoveOutgoingPtr(ref interface{}){
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
func (this *DataOutput) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *DataOutput) SetIncomingPtr(ref interface{}){
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
func (this *DataOutput) RemoveIncomingPtr(ref interface{}){
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

/** DataAssociation **/
func (this *DataOutput) GetDataAssociationPtr() []DataAssociation{
	return this.m_dataAssociationPtr
}

/** Init reference DataAssociation **/
func (this *DataOutput) SetDataAssociationPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_dataAssociationPtr); i++ {
			if this.M_dataAssociationPtr[i] == refStr {
				return
			}
		}
		this.M_dataAssociationPtr = append(this.M_dataAssociationPtr, ref.(string))
	}else{
		this.RemoveDataAssociationPtr(ref)
		this.m_dataAssociationPtr = append(this.m_dataAssociationPtr, ref.(DataAssociation))
		this.M_dataAssociationPtr = append(this.M_dataAssociationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference DataAssociation **/
func (this *DataOutput) RemoveDataAssociationPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataAssociationPtr_ := make([]DataAssociation, 0)
	dataAssociationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_dataAssociationPtr); i++ {
		if toDelete.GetUUID() != this.m_dataAssociationPtr[i].(BaseElement).GetUUID() {
			dataAssociationPtr_ = append(dataAssociationPtr_, this.m_dataAssociationPtr[i])
			dataAssociationPtrUuid = append(dataAssociationPtrUuid, this.M_dataAssociationPtr[i])
		}
	}
	this.m_dataAssociationPtr = dataAssociationPtr_
	this.M_dataAssociationPtr = dataAssociationPtrUuid
}
