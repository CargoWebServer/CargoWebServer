// +build BPMN20

package BPMN20

import(
"encoding/xml"
)

type Interface struct{

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

	/** members of RootElement **/
	/** No members **/

	/** members of Interface **/
	M_name string
	M_operation []*Operation
	m_implementationRef interface{}
	/** If the ref is a string and not an object **/
	M_implementationRef string
	M_implementation Implementation
	M_implementationStr string


	/** Associations **/
	m_callableElementsPtr []CallableElement
	/** If the ref is a string and not an object **/
	M_callableElementsPtr []string
	m_participantPtr []*Participant
	/** If the ref is a string and not an object **/
	M_participantPtr []string
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
	m_definitionsPtr *Definitions
	/** If the ref is a string and not an object **/
	M_definitionsPtr string
}

/** Xml parser for Interface **/
type XsdInterface struct {
	XMLName xml.Name	`xml:"interface"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** RootElement **/


	M_operation	[]*XsdOperation	`xml:"operation,omitempty"`
	M_name	string	`xml:"name,attr"`
	M_implementationRef	string	`xml:"implementationRef,attr"`

}
/** UUID **/
func (this *Interface) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Interface) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Interface) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Interface) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *Interface) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Interface) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Interface) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/
func (this *Interface) RemoveExtensionElements(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ExtensionElements)
	if toDelete.GetUUID() == this.M_extensionElements.GetUUID() {
		this.M_extensionElements = nil
	}
}

/** ExtensionDefinitions **/
func (this *Interface) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Interface) SetExtensionDefinitions(ref interface{}){
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
func (this *Interface) RemoveExtensionDefinitions(ref interface{}){
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
func (this *Interface) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Interface) SetExtensionValues(ref interface{}){
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
func (this *Interface) RemoveExtensionValues(ref interface{}){
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
func (this *Interface) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Interface) SetDocumentation(ref interface{}){
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
func (this *Interface) RemoveDocumentation(ref interface{}){
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
func (this *Interface) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Interface) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Operation **/
func (this *Interface) GetOperation() []*Operation{
	return this.M_operation
}

/** Init reference Operation **/
func (this *Interface) SetOperation(ref interface{}){
	this.NeedSave = true
	isExist := false
	var operations []*Operation
	for i:=0; i<len(this.M_operation); i++ {
		if this.M_operation[i].GetUUID() != ref.(BaseElement).GetUUID() {
			operations = append(operations, this.M_operation[i])
		} else {
			isExist = true
			operations = append(operations, ref.(*Operation))
		}
	}
	if !isExist {
		operations = append(operations, ref.(*Operation))
	}
	this.M_operation = operations
}

/** Remove reference Operation **/
func (this *Interface) RemoveOperation(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	operation_ := make([]*Operation, 0)
	for i := 0; i < len(this.M_operation); i++ {
		if toDelete.GetUUID() != this.M_operation[i].GetUUID() {
			operation_ = append(operation_, this.M_operation[i])
		}
	}
	this.M_operation = operation_
}

/** ImplementationRef **/
func (this *Interface) GetImplementationRef() interface{}{
	return this.m_implementationRef
}

/** Init reference ImplementationRef **/
func (this *Interface) SetImplementationRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_implementationRef = ref.(string)
	}else{
		this.m_implementationRef = ref.(interface{})
	}
}

/** Remove reference ImplementationRef **/

/** Implementation **/
func (this *Interface) GetImplementation() Implementation{
	return this.M_implementation
}

/** Init reference Implementation **/
func (this *Interface) SetImplementation(ref interface{}){
	this.NeedSave = true
	this.M_implementation = ref.(Implementation)
}

/** Remove reference Implementation **/

/** ImplementationStr **/
func (this *Interface) GetImplementationStr() string{
	return this.M_implementationStr
}

/** Init reference ImplementationStr **/
func (this *Interface) SetImplementationStr(ref interface{}){
	this.NeedSave = true
	this.M_implementationStr = ref.(string)
}

/** Remove reference ImplementationStr **/

/** CallableElements **/
func (this *Interface) GetCallableElementsPtr() []CallableElement{
	return this.m_callableElementsPtr
}

/** Init reference CallableElements **/
func (this *Interface) SetCallableElementsPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_callableElementsPtr); i++ {
			if this.M_callableElementsPtr[i] == refStr {
				return
			}
		}
		this.M_callableElementsPtr = append(this.M_callableElementsPtr, ref.(string))
	}else{
		this.RemoveCallableElementsPtr(ref)
		this.m_callableElementsPtr = append(this.m_callableElementsPtr, ref.(CallableElement))
		this.M_callableElementsPtr = append(this.M_callableElementsPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CallableElements **/
func (this *Interface) RemoveCallableElementsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	callableElementsPtr_ := make([]CallableElement, 0)
	callableElementsPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_callableElementsPtr); i++ {
		if toDelete.GetUUID() != this.m_callableElementsPtr[i].(BaseElement).GetUUID() {
			callableElementsPtr_ = append(callableElementsPtr_, this.m_callableElementsPtr[i])
			callableElementsPtrUuid = append(callableElementsPtrUuid, this.M_callableElementsPtr[i])
		}
	}
	this.m_callableElementsPtr = callableElementsPtr_
	this.M_callableElementsPtr = callableElementsPtrUuid
}

/** Participant **/
func (this *Interface) GetParticipantPtr() []*Participant{
	return this.m_participantPtr
}

/** Init reference Participant **/
func (this *Interface) SetParticipantPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_participantPtr); i++ {
			if this.M_participantPtr[i] == refStr {
				return
			}
		}
		this.M_participantPtr = append(this.M_participantPtr, ref.(string))
	}else{
		this.RemoveParticipantPtr(ref)
		this.m_participantPtr = append(this.m_participantPtr, ref.(*Participant))
		this.M_participantPtr = append(this.M_participantPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Participant **/
func (this *Interface) RemoveParticipantPtr(ref interface{}){
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
func (this *Interface) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Interface) SetLanePtr(ref interface{}){
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
func (this *Interface) RemoveLanePtr(ref interface{}){
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
func (this *Interface) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Interface) SetOutgoingPtr(ref interface{}){
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
func (this *Interface) RemoveOutgoingPtr(ref interface{}){
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
func (this *Interface) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Interface) SetIncomingPtr(ref interface{}){
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
func (this *Interface) RemoveIncomingPtr(ref interface{}){
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
func (this *Interface) GetDefinitionsPtr() *Definitions{
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *Interface) SetDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	}else{
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *Interface) RemoveDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}
