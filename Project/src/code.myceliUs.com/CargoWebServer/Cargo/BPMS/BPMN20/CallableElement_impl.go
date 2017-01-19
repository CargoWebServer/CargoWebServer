package BPMN20

import(
"encoding/xml"
)

type CallableElement_impl struct{

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

	/** members of CallableElement **/
	M_name string
	M_ioSpecification *InputOutputSpecification
	m_supportedInterfaceRef []*Interface
	/** If the ref is a string and not an object **/
	M_supportedInterfaceRef []string
	M_ioBinding []*InputOutputBinding


	/** Associations **/
	m_callActivityPtr []*CallActivity
	/** If the ref is a string and not an object **/
	M_callActivityPtr []string
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

/** Xml parser for CallableElement **/
type XsdCallableElement struct {
	XMLName xml.Name	`xml:"callableElement"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** RootElement **/


	M_supportedInterfaceRef	[]string	`xml:"supportedInterfaceRef"`
	M_ioSpecification	*XsdInputOutputSpecification	`xml:"ioSpecification,omitempty"`
	M_ioBinding	[]*XsdInputOutputBinding	`xml:"ioBinding,omitempty"`
	M_name	string	`xml:"name,attr"`

}
/** UUID **/
func (this *CallableElement_impl) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *CallableElement_impl) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *CallableElement_impl) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *CallableElement_impl) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *CallableElement_impl) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *CallableElement_impl) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *CallableElement_impl) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *CallableElement_impl) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *CallableElement_impl) SetExtensionDefinitions(ref interface{}){
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
func (this *CallableElement_impl) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *CallableElement_impl) SetExtensionValues(ref interface{}){
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
func (this *CallableElement_impl) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *CallableElement_impl) SetDocumentation(ref interface{}){
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
func (this *CallableElement_impl) RemoveDocumentation(ref interface{}){
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
func (this *CallableElement_impl) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *CallableElement_impl) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** IoSpecification **/
func (this *CallableElement_impl) GetIoSpecification() *InputOutputSpecification{
	return this.M_ioSpecification
}

/** Init reference IoSpecification **/
func (this *CallableElement_impl) SetIoSpecification(ref interface{}){
	this.NeedSave = true
	this.M_ioSpecification = ref.(*InputOutputSpecification)
}

/** Remove reference IoSpecification **/
func (this *CallableElement_impl) RemoveIoSpecification(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_ioSpecification.GetUUID() {
		this.M_ioSpecification = nil
	}
}

/** SupportedInterfaceRef **/
func (this *CallableElement_impl) GetSupportedInterfaceRef() []*Interface{
	return this.m_supportedInterfaceRef
}

/** Init reference SupportedInterfaceRef **/
func (this *CallableElement_impl) SetSupportedInterfaceRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_supportedInterfaceRef); i++ {
			if this.M_supportedInterfaceRef[i] == refStr {
				return
			}
		}
		this.M_supportedInterfaceRef = append(this.M_supportedInterfaceRef, ref.(string))
	}else{
		this.RemoveSupportedInterfaceRef(ref)
		this.m_supportedInterfaceRef = append(this.m_supportedInterfaceRef, ref.(*Interface))
		this.M_supportedInterfaceRef = append(this.M_supportedInterfaceRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference SupportedInterfaceRef **/
func (this *CallableElement_impl) RemoveSupportedInterfaceRef(ref interface{}){
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
func (this *CallableElement_impl) GetIoBinding() []*InputOutputBinding{
	return this.M_ioBinding
}

/** Init reference IoBinding **/
func (this *CallableElement_impl) SetIoBinding(ref interface{}){
	this.NeedSave = true
	isExist := false
	var ioBindings []*InputOutputBinding
	for i:=0; i<len(this.M_ioBinding); i++ {
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
func (this *CallableElement_impl) RemoveIoBinding(ref interface{}){
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

/** CallActivity **/
func (this *CallableElement_impl) GetCallActivityPtr() []*CallActivity{
	return this.m_callActivityPtr
}

/** Init reference CallActivity **/
func (this *CallableElement_impl) SetCallActivityPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_callActivityPtr); i++ {
			if this.M_callActivityPtr[i] == refStr {
				return
			}
		}
		this.M_callActivityPtr = append(this.M_callActivityPtr, ref.(string))
	}else{
		this.RemoveCallActivityPtr(ref)
		this.m_callActivityPtr = append(this.m_callActivityPtr, ref.(*CallActivity))
		this.M_callActivityPtr = append(this.M_callActivityPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CallActivity **/
func (this *CallableElement_impl) RemoveCallActivityPtr(ref interface{}){
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

/** Lane **/
func (this *CallableElement_impl) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *CallableElement_impl) SetLanePtr(ref interface{}){
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
func (this *CallableElement_impl) RemoveLanePtr(ref interface{}){
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
func (this *CallableElement_impl) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *CallableElement_impl) SetOutgoingPtr(ref interface{}){
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
func (this *CallableElement_impl) RemoveOutgoingPtr(ref interface{}){
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
func (this *CallableElement_impl) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *CallableElement_impl) SetIncomingPtr(ref interface{}){
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
func (this *CallableElement_impl) RemoveIncomingPtr(ref interface{}){
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
func (this *CallableElement_impl) GetDefinitionsPtr() *Definitions{
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *CallableElement_impl) SetDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	}else{
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *CallableElement_impl) RemoveDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}
