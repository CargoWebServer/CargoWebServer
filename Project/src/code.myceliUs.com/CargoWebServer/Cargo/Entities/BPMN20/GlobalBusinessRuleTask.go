// +build BPMN20

package BPMN20

import(
"encoding/xml"
)

type GlobalBusinessRuleTask struct{

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

	/** members of GlobalTask **/
	M_resourceRole []ResourceRole

	/** members of GlobalBusinessRuleTask **/
	M_implementation Implementation
	M_implementationStr string


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
	m_definitionsPtr *Definitions
	/** If the ref is a string and not an object **/
	M_definitionsPtr string
	m_callActivityPtr []*CallActivity
	/** If the ref is a string and not an object **/
	M_callActivityPtr []string
}

/** Xml parser for GlobalBusinessRuleTask **/
type XsdGlobalBusinessRuleTask struct {
	XMLName xml.Name	`xml:"globalBusinessRuleTask"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** RootElement **/


	/** CallableElement **/
	M_supportedInterfaceRef	[]string	`xml:"supportedInterfaceRef"`
	M_ioSpecification	*XsdInputOutputSpecification	`xml:"ioSpecification,omitempty"`
	M_ioBinding	[]*XsdInputOutputBinding	`xml:"ioBinding,omitempty"`
	M_name	string	`xml:"name,attr"`


	/** GlobalTask **/
	M_resourceRole_0	[]*XsdHumanPerformer	`xml:"humanPerformer,omitempty"`
	M_resourceRole_1	[]*XsdPotentialOwner	`xml:"potentialOwner,omitempty"`
	M_resourceRole_2	[]*XsdPerformer	`xml:"performer,omitempty"`
	M_resourceRole_3	[]*XsdResourceRole	`xml:"resourceRole,omitempty"`


	M_implementation	string	`xml:"implementation,attr"`

}
/** UUID **/
func (this *GlobalBusinessRuleTask) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *GlobalBusinessRuleTask) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *GlobalBusinessRuleTask) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *GlobalBusinessRuleTask) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *GlobalBusinessRuleTask) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *GlobalBusinessRuleTask) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *GlobalBusinessRuleTask) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/
func (this *GlobalBusinessRuleTask) RemoveExtensionElements(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ExtensionElements)
	if toDelete.GetUUID() == this.M_extensionElements.GetUUID() {
		this.M_extensionElements = nil
	}
}

/** ExtensionDefinitions **/
func (this *GlobalBusinessRuleTask) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *GlobalBusinessRuleTask) SetExtensionDefinitions(ref interface{}){
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
func (this *GlobalBusinessRuleTask) RemoveExtensionDefinitions(ref interface{}){
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
func (this *GlobalBusinessRuleTask) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *GlobalBusinessRuleTask) SetExtensionValues(ref interface{}){
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
func (this *GlobalBusinessRuleTask) RemoveExtensionValues(ref interface{}){
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
func (this *GlobalBusinessRuleTask) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *GlobalBusinessRuleTask) SetDocumentation(ref interface{}){
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
func (this *GlobalBusinessRuleTask) RemoveDocumentation(ref interface{}){
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
func (this *GlobalBusinessRuleTask) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *GlobalBusinessRuleTask) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** IoSpecification **/
func (this *GlobalBusinessRuleTask) GetIoSpecification() *InputOutputSpecification{
	return this.M_ioSpecification
}

/** Init reference IoSpecification **/
func (this *GlobalBusinessRuleTask) SetIoSpecification(ref interface{}){
	this.NeedSave = true
	this.M_ioSpecification = ref.(*InputOutputSpecification)
}

/** Remove reference IoSpecification **/
func (this *GlobalBusinessRuleTask) RemoveIoSpecification(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_ioSpecification.GetUUID() {
		this.M_ioSpecification = nil
	}
}

/** SupportedInterfaceRef **/
func (this *GlobalBusinessRuleTask) GetSupportedInterfaceRef() []*Interface{
	return this.m_supportedInterfaceRef
}

/** Init reference SupportedInterfaceRef **/
func (this *GlobalBusinessRuleTask) SetSupportedInterfaceRef(ref interface{}){
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
func (this *GlobalBusinessRuleTask) RemoveSupportedInterfaceRef(ref interface{}){
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
func (this *GlobalBusinessRuleTask) GetIoBinding() []*InputOutputBinding{
	return this.M_ioBinding
}

/** Init reference IoBinding **/
func (this *GlobalBusinessRuleTask) SetIoBinding(ref interface{}){
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
func (this *GlobalBusinessRuleTask) RemoveIoBinding(ref interface{}){
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

/** ResourceRole **/
func (this *GlobalBusinessRuleTask) GetResourceRole() []ResourceRole{
	return this.M_resourceRole
}

/** Init reference ResourceRole **/
func (this *GlobalBusinessRuleTask) SetResourceRole(ref interface{}){
	this.NeedSave = true
	isExist := false
	var resourceRoles []ResourceRole
	for i:=0; i<len(this.M_resourceRole); i++ {
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
func (this *GlobalBusinessRuleTask) RemoveResourceRole(ref interface{}){
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

/** Implementation **/
func (this *GlobalBusinessRuleTask) GetImplementation() Implementation{
	return this.M_implementation
}

/** Init reference Implementation **/
func (this *GlobalBusinessRuleTask) SetImplementation(ref interface{}){
	this.NeedSave = true
	this.M_implementation = ref.(Implementation)
}

/** Remove reference Implementation **/

/** ImplementationStr **/
func (this *GlobalBusinessRuleTask) GetImplementationStr() string{
	return this.M_implementationStr
}

/** Init reference ImplementationStr **/
func (this *GlobalBusinessRuleTask) SetImplementationStr(ref interface{}){
	this.NeedSave = true
	this.M_implementationStr = ref.(string)
}

/** Remove reference ImplementationStr **/

/** Lane **/
func (this *GlobalBusinessRuleTask) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *GlobalBusinessRuleTask) SetLanePtr(ref interface{}){
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
func (this *GlobalBusinessRuleTask) RemoveLanePtr(ref interface{}){
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
func (this *GlobalBusinessRuleTask) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *GlobalBusinessRuleTask) SetOutgoingPtr(ref interface{}){
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
func (this *GlobalBusinessRuleTask) RemoveOutgoingPtr(ref interface{}){
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
func (this *GlobalBusinessRuleTask) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *GlobalBusinessRuleTask) SetIncomingPtr(ref interface{}){
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
func (this *GlobalBusinessRuleTask) RemoveIncomingPtr(ref interface{}){
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
func (this *GlobalBusinessRuleTask) GetDefinitionsPtr() *Definitions{
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *GlobalBusinessRuleTask) SetDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	}else{
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *GlobalBusinessRuleTask) RemoveDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}

/** CallActivity **/
func (this *GlobalBusinessRuleTask) GetCallActivityPtr() []*CallActivity{
	return this.m_callActivityPtr
}

/** Init reference CallActivity **/
func (this *GlobalBusinessRuleTask) SetCallActivityPtr(ref interface{}){
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
func (this *GlobalBusinessRuleTask) RemoveCallActivityPtr(ref interface{}){
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
