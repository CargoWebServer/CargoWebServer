package BPMN20

import(
"encoding/xml"
)

type ResourceParameter struct{

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

	/** members of ResourceParameter **/
	M_name string
	M_isRequired bool
	m_type *ItemDefinition
	/** If the ref is a string and not an object **/
	M_type string


	/** Associations **/
	m_resourcePtr *Resource
	/** If the ref is a string and not an object **/
	M_resourcePtr string
	m_resourceParameterBindingPtr []*ResourceParameterBinding
	/** If the ref is a string and not an object **/
	M_resourceParameterBindingPtr []string
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

/** Xml parser for ResourceParameter **/
type XsdResourceParameter struct {
	XMLName xml.Name	`xml:"resourceParameter"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	M_name	string	`xml:"name,attr"`
	M_type	string	`xml:"type,attr"`
	M_isRequired	bool	`xml:"isRequired,attr"`

}
/** UUID **/
func (this *ResourceParameter) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *ResourceParameter) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ResourceParameter) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *ResourceParameter) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *ResourceParameter) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *ResourceParameter) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *ResourceParameter) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *ResourceParameter) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *ResourceParameter) SetExtensionDefinitions(ref interface{}){
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
func (this *ResourceParameter) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *ResourceParameter) SetExtensionValues(ref interface{}){
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
func (this *ResourceParameter) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *ResourceParameter) SetDocumentation(ref interface{}){
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
func (this *ResourceParameter) RemoveDocumentation(ref interface{}){
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
func (this *ResourceParameter) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *ResourceParameter) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** IsRequired **/
func (this *ResourceParameter) IsRequired() bool{
	return this.M_isRequired
}

/** Init reference IsRequired **/
func (this *ResourceParameter) SetIsRequired(ref interface{}){
	this.NeedSave = true
	this.M_isRequired = ref.(bool)
}

/** Remove reference IsRequired **/

/** Type **/
func (this *ResourceParameter) GetType() *ItemDefinition{
	return this.m_type
}

/** Init reference Type **/
func (this *ResourceParameter) SetType(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_type = ref.(string)
	}else{
		this.m_type = ref.(*ItemDefinition)
		this.M_type = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Type **/
func (this *ResourceParameter) RemoveType(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_type.GetUUID() {
		this.m_type = nil
		this.M_type = ""
	}
}

/** Resource **/
func (this *ResourceParameter) GetResourcePtr() *Resource{
	return this.m_resourcePtr
}

/** Init reference Resource **/
func (this *ResourceParameter) SetResourcePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_resourcePtr = ref.(string)
	}else{
		this.m_resourcePtr = ref.(*Resource)
		this.M_resourcePtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Resource **/
func (this *ResourceParameter) RemoveResourcePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_resourcePtr.GetUUID() {
		this.m_resourcePtr = nil
		this.M_resourcePtr = ""
	}
}

/** ResourceParameterBinding **/
func (this *ResourceParameter) GetResourceParameterBindingPtr() []*ResourceParameterBinding{
	return this.m_resourceParameterBindingPtr
}

/** Init reference ResourceParameterBinding **/
func (this *ResourceParameter) SetResourceParameterBindingPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_resourceParameterBindingPtr); i++ {
			if this.M_resourceParameterBindingPtr[i] == refStr {
				return
			}
		}
		this.M_resourceParameterBindingPtr = append(this.M_resourceParameterBindingPtr, ref.(string))
	}else{
		this.RemoveResourceParameterBindingPtr(ref)
		this.m_resourceParameterBindingPtr = append(this.m_resourceParameterBindingPtr, ref.(*ResourceParameterBinding))
		this.M_resourceParameterBindingPtr = append(this.M_resourceParameterBindingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ResourceParameterBinding **/
func (this *ResourceParameter) RemoveResourceParameterBindingPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	resourceParameterBindingPtr_ := make([]*ResourceParameterBinding, 0)
	resourceParameterBindingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_resourceParameterBindingPtr); i++ {
		if toDelete.GetUUID() != this.m_resourceParameterBindingPtr[i].GetUUID() {
			resourceParameterBindingPtr_ = append(resourceParameterBindingPtr_, this.m_resourceParameterBindingPtr[i])
			resourceParameterBindingPtrUuid = append(resourceParameterBindingPtrUuid, this.M_resourceParameterBindingPtr[i])
		}
	}
	this.m_resourceParameterBindingPtr = resourceParameterBindingPtr_
	this.M_resourceParameterBindingPtr = resourceParameterBindingPtrUuid
}

/** Lane **/
func (this *ResourceParameter) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *ResourceParameter) SetLanePtr(ref interface{}){
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
func (this *ResourceParameter) RemoveLanePtr(ref interface{}){
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
func (this *ResourceParameter) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *ResourceParameter) SetOutgoingPtr(ref interface{}){
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
func (this *ResourceParameter) RemoveOutgoingPtr(ref interface{}){
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
func (this *ResourceParameter) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *ResourceParameter) SetIncomingPtr(ref interface{}){
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
func (this *ResourceParameter) RemoveIncomingPtr(ref interface{}){
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
