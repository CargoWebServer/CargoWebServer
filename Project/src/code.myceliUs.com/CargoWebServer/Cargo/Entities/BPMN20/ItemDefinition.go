package BPMN20

import(
"encoding/xml"
)

type ItemDefinition struct{

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

	/** members of ItemDefinition **/
	M_itemKind ItemKind
	m_structureRef interface{}
	/** If the ref is a string and not an object **/
	M_structureRef string
	M_isCollection bool
	m_import *Import
	/** If the ref is a string and not an object **/
	M_import string


	/** Associations **/
	m_escalationPtr []*Escalation
	/** If the ref is a string and not an object **/
	M_escalationPtr []string
	m_signalPtr []*Signal
	/** If the ref is a string and not an object **/
	M_signalPtr []string
	m_itemAwareElementPtr []ItemAwareElement
	/** If the ref is a string and not an object **/
	M_itemAwareElementPtr []string
	m_correlationPropertyPtr []*CorrelationProperty
	/** If the ref is a string and not an object **/
	M_correlationPropertyPtr []string
	m_errorPtr []*Error
	/** If the ref is a string and not an object **/
	M_errorPtr []string
	m_formalExpressionPtr []*FormalExpression
	/** If the ref is a string and not an object **/
	M_formalExpressionPtr []string
	m_messagePtr []*Message
	/** If the ref is a string and not an object **/
	M_messagePtr []string
	m_resourceParameterPtr []*ResourceParameter
	/** If the ref is a string and not an object **/
	M_resourceParameterPtr []string
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

/** Xml parser for ItemDefinition **/
type XsdItemDefinition struct {
	XMLName xml.Name	`xml:"itemDefinition"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** RootElement **/


	M_structureRef	string	`xml:"structureRef,attr"`
	M_isCollection	bool	`xml:"isCollection,attr"`
	M_itemKind	string	`xml:"itemKind,attr"`

}
/** UUID **/
func (this *ItemDefinition) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *ItemDefinition) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ItemDefinition) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *ItemDefinition) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *ItemDefinition) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *ItemDefinition) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *ItemDefinition) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *ItemDefinition) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *ItemDefinition) SetExtensionDefinitions(ref interface{}){
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
func (this *ItemDefinition) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *ItemDefinition) SetExtensionValues(ref interface{}){
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
func (this *ItemDefinition) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *ItemDefinition) SetDocumentation(ref interface{}){
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
func (this *ItemDefinition) RemoveDocumentation(ref interface{}){
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

/** ItemKind **/
func (this *ItemDefinition) GetItemKind() ItemKind{
	return this.M_itemKind
}

/** Init reference ItemKind **/
func (this *ItemDefinition) SetItemKind(ref interface{}){
	this.NeedSave = true
	this.M_itemKind = ref.(ItemKind)
}

/** Remove reference ItemKind **/

/** StructureRef **/
func (this *ItemDefinition) GetStructureRef() interface{}{
	return this.m_structureRef
}

/** Init reference StructureRef **/
func (this *ItemDefinition) SetStructureRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_structureRef = ref.(string)
	}else{
		this.m_structureRef = ref.(interface{})
	}
}

/** Remove reference StructureRef **/

/** IsCollection **/
func (this *ItemDefinition) IsCollection() bool{
	return this.M_isCollection
}

/** Init reference IsCollection **/
func (this *ItemDefinition) SetIsCollection(ref interface{}){
	this.NeedSave = true
	this.M_isCollection = ref.(bool)
}

/** Remove reference IsCollection **/

/** Import **/
func (this *ItemDefinition) GetImport() *Import{
	return this.m_import
}

/** Init reference Import **/
func (this *ItemDefinition) SetImport(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_import = ref.(string)
	}else{
		this.m_import = ref.(*Import)
	}
}

/** Remove reference Import **/

/** Escalation **/
func (this *ItemDefinition) GetEscalationPtr() []*Escalation{
	return this.m_escalationPtr
}

/** Init reference Escalation **/
func (this *ItemDefinition) SetEscalationPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_escalationPtr); i++ {
			if this.M_escalationPtr[i] == refStr {
				return
			}
		}
		this.M_escalationPtr = append(this.M_escalationPtr, ref.(string))
	}else{
		this.RemoveEscalationPtr(ref)
		this.m_escalationPtr = append(this.m_escalationPtr, ref.(*Escalation))
		this.M_escalationPtr = append(this.M_escalationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Escalation **/
func (this *ItemDefinition) RemoveEscalationPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	escalationPtr_ := make([]*Escalation, 0)
	escalationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_escalationPtr); i++ {
		if toDelete.GetUUID() != this.m_escalationPtr[i].GetUUID() {
			escalationPtr_ = append(escalationPtr_, this.m_escalationPtr[i])
			escalationPtrUuid = append(escalationPtrUuid, this.M_escalationPtr[i])
		}
	}
	this.m_escalationPtr = escalationPtr_
	this.M_escalationPtr = escalationPtrUuid
}

/** Signal **/
func (this *ItemDefinition) GetSignalPtr() []*Signal{
	return this.m_signalPtr
}

/** Init reference Signal **/
func (this *ItemDefinition) SetSignalPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_signalPtr); i++ {
			if this.M_signalPtr[i] == refStr {
				return
			}
		}
		this.M_signalPtr = append(this.M_signalPtr, ref.(string))
	}else{
		this.RemoveSignalPtr(ref)
		this.m_signalPtr = append(this.m_signalPtr, ref.(*Signal))
		this.M_signalPtr = append(this.M_signalPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Signal **/
func (this *ItemDefinition) RemoveSignalPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	signalPtr_ := make([]*Signal, 0)
	signalPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_signalPtr); i++ {
		if toDelete.GetUUID() != this.m_signalPtr[i].GetUUID() {
			signalPtr_ = append(signalPtr_, this.m_signalPtr[i])
			signalPtrUuid = append(signalPtrUuid, this.M_signalPtr[i])
		}
	}
	this.m_signalPtr = signalPtr_
	this.M_signalPtr = signalPtrUuid
}

/** ItemAwareElement **/
func (this *ItemDefinition) GetItemAwareElementPtr() []ItemAwareElement{
	return this.m_itemAwareElementPtr
}

/** Init reference ItemAwareElement **/
func (this *ItemDefinition) SetItemAwareElementPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_itemAwareElementPtr); i++ {
			if this.M_itemAwareElementPtr[i] == refStr {
				return
			}
		}
		this.M_itemAwareElementPtr = append(this.M_itemAwareElementPtr, ref.(string))
	}else{
		this.RemoveItemAwareElementPtr(ref)
		this.m_itemAwareElementPtr = append(this.m_itemAwareElementPtr, ref.(ItemAwareElement))
		this.M_itemAwareElementPtr = append(this.M_itemAwareElementPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ItemAwareElement **/
func (this *ItemDefinition) RemoveItemAwareElementPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	itemAwareElementPtr_ := make([]ItemAwareElement, 0)
	itemAwareElementPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_itemAwareElementPtr); i++ {
		if toDelete.GetUUID() != this.m_itemAwareElementPtr[i].(BaseElement).GetUUID() {
			itemAwareElementPtr_ = append(itemAwareElementPtr_, this.m_itemAwareElementPtr[i])
			itemAwareElementPtrUuid = append(itemAwareElementPtrUuid, this.M_itemAwareElementPtr[i])
		}
	}
	this.m_itemAwareElementPtr = itemAwareElementPtr_
	this.M_itemAwareElementPtr = itemAwareElementPtrUuid
}

/** CorrelationProperty **/
func (this *ItemDefinition) GetCorrelationPropertyPtr() []*CorrelationProperty{
	return this.m_correlationPropertyPtr
}

/** Init reference CorrelationProperty **/
func (this *ItemDefinition) SetCorrelationPropertyPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_correlationPropertyPtr); i++ {
			if this.M_correlationPropertyPtr[i] == refStr {
				return
			}
		}
		this.M_correlationPropertyPtr = append(this.M_correlationPropertyPtr, ref.(string))
	}else{
		this.RemoveCorrelationPropertyPtr(ref)
		this.m_correlationPropertyPtr = append(this.m_correlationPropertyPtr, ref.(*CorrelationProperty))
		this.M_correlationPropertyPtr = append(this.M_correlationPropertyPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CorrelationProperty **/
func (this *ItemDefinition) RemoveCorrelationPropertyPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	correlationPropertyPtr_ := make([]*CorrelationProperty, 0)
	correlationPropertyPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_correlationPropertyPtr); i++ {
		if toDelete.GetUUID() != this.m_correlationPropertyPtr[i].GetUUID() {
			correlationPropertyPtr_ = append(correlationPropertyPtr_, this.m_correlationPropertyPtr[i])
			correlationPropertyPtrUuid = append(correlationPropertyPtrUuid, this.M_correlationPropertyPtr[i])
		}
	}
	this.m_correlationPropertyPtr = correlationPropertyPtr_
	this.M_correlationPropertyPtr = correlationPropertyPtrUuid
}

/** Error **/
func (this *ItemDefinition) GetErrorPtr() []*Error{
	return this.m_errorPtr
}

/** Init reference Error **/
func (this *ItemDefinition) SetErrorPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_errorPtr); i++ {
			if this.M_errorPtr[i] == refStr {
				return
			}
		}
		this.M_errorPtr = append(this.M_errorPtr, ref.(string))
	}else{
		this.RemoveErrorPtr(ref)
		this.m_errorPtr = append(this.m_errorPtr, ref.(*Error))
		this.M_errorPtr = append(this.M_errorPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Error **/
func (this *ItemDefinition) RemoveErrorPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	errorPtr_ := make([]*Error, 0)
	errorPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_errorPtr); i++ {
		if toDelete.GetUUID() != this.m_errorPtr[i].GetUUID() {
			errorPtr_ = append(errorPtr_, this.m_errorPtr[i])
			errorPtrUuid = append(errorPtrUuid, this.M_errorPtr[i])
		}
	}
	this.m_errorPtr = errorPtr_
	this.M_errorPtr = errorPtrUuid
}

/** FormalExpression **/
func (this *ItemDefinition) GetFormalExpressionPtr() []*FormalExpression{
	return this.m_formalExpressionPtr
}

/** Init reference FormalExpression **/
func (this *ItemDefinition) SetFormalExpressionPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_formalExpressionPtr); i++ {
			if this.M_formalExpressionPtr[i] == refStr {
				return
			}
		}
		this.M_formalExpressionPtr = append(this.M_formalExpressionPtr, ref.(string))
	}else{
		this.RemoveFormalExpressionPtr(ref)
		this.m_formalExpressionPtr = append(this.m_formalExpressionPtr, ref.(*FormalExpression))
		this.M_formalExpressionPtr = append(this.M_formalExpressionPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference FormalExpression **/
func (this *ItemDefinition) RemoveFormalExpressionPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	formalExpressionPtr_ := make([]*FormalExpression, 0)
	formalExpressionPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_formalExpressionPtr); i++ {
		if toDelete.GetUUID() != this.m_formalExpressionPtr[i].GetUUID() {
			formalExpressionPtr_ = append(formalExpressionPtr_, this.m_formalExpressionPtr[i])
			formalExpressionPtrUuid = append(formalExpressionPtrUuid, this.M_formalExpressionPtr[i])
		}
	}
	this.m_formalExpressionPtr = formalExpressionPtr_
	this.M_formalExpressionPtr = formalExpressionPtrUuid
}

/** Message **/
func (this *ItemDefinition) GetMessagePtr() []*Message{
	return this.m_messagePtr
}

/** Init reference Message **/
func (this *ItemDefinition) SetMessagePtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_messagePtr); i++ {
			if this.M_messagePtr[i] == refStr {
				return
			}
		}
		this.M_messagePtr = append(this.M_messagePtr, ref.(string))
	}else{
		this.RemoveMessagePtr(ref)
		this.m_messagePtr = append(this.m_messagePtr, ref.(*Message))
		this.M_messagePtr = append(this.M_messagePtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Message **/
func (this *ItemDefinition) RemoveMessagePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	messagePtr_ := make([]*Message, 0)
	messagePtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_messagePtr); i++ {
		if toDelete.GetUUID() != this.m_messagePtr[i].GetUUID() {
			messagePtr_ = append(messagePtr_, this.m_messagePtr[i])
			messagePtrUuid = append(messagePtrUuid, this.M_messagePtr[i])
		}
	}
	this.m_messagePtr = messagePtr_
	this.M_messagePtr = messagePtrUuid
}

/** ResourceParameter **/
func (this *ItemDefinition) GetResourceParameterPtr() []*ResourceParameter{
	return this.m_resourceParameterPtr
}

/** Init reference ResourceParameter **/
func (this *ItemDefinition) SetResourceParameterPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_resourceParameterPtr); i++ {
			if this.M_resourceParameterPtr[i] == refStr {
				return
			}
		}
		this.M_resourceParameterPtr = append(this.M_resourceParameterPtr, ref.(string))
	}else{
		this.RemoveResourceParameterPtr(ref)
		this.m_resourceParameterPtr = append(this.m_resourceParameterPtr, ref.(*ResourceParameter))
		this.M_resourceParameterPtr = append(this.M_resourceParameterPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ResourceParameter **/
func (this *ItemDefinition) RemoveResourceParameterPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	resourceParameterPtr_ := make([]*ResourceParameter, 0)
	resourceParameterPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_resourceParameterPtr); i++ {
		if toDelete.GetUUID() != this.m_resourceParameterPtr[i].GetUUID() {
			resourceParameterPtr_ = append(resourceParameterPtr_, this.m_resourceParameterPtr[i])
			resourceParameterPtrUuid = append(resourceParameterPtrUuid, this.M_resourceParameterPtr[i])
		}
	}
	this.m_resourceParameterPtr = resourceParameterPtr_
	this.M_resourceParameterPtr = resourceParameterPtrUuid
}

/** Lane **/
func (this *ItemDefinition) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *ItemDefinition) SetLanePtr(ref interface{}){
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
func (this *ItemDefinition) RemoveLanePtr(ref interface{}){
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
func (this *ItemDefinition) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *ItemDefinition) SetOutgoingPtr(ref interface{}){
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
func (this *ItemDefinition) RemoveOutgoingPtr(ref interface{}){
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
func (this *ItemDefinition) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *ItemDefinition) SetIncomingPtr(ref interface{}){
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
func (this *ItemDefinition) RemoveIncomingPtr(ref interface{}){
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
func (this *ItemDefinition) GetDefinitionsPtr() *Definitions{
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *ItemDefinition) SetDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	}else{
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *ItemDefinition) RemoveDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}
