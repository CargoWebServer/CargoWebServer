package BPMN20

import(
"encoding/xml"
)

type CategoryValue struct{

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

	/** members of CategoryValue **/
	m_categorizedFlowElements []FlowElement
	/** If the ref is a string and not an object **/
	M_categorizedFlowElements []string
	M_value string


	/** Associations **/
	m_categoryValueRefPtr []*Group
	/** If the ref is a string and not an object **/
	M_categoryValueRefPtr []string
	m_categoryPtr *Category
	/** If the ref is a string and not an object **/
	M_categoryPtr string
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

/** Xml parser for CategoryValue **/
type XsdCategoryValue struct {
	XMLName xml.Name	`xml:"categoryValue"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	M_value	string	`xml:"value,attr"`

}
/** UUID **/
func (this *CategoryValue) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *CategoryValue) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *CategoryValue) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *CategoryValue) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *CategoryValue) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *CategoryValue) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *CategoryValue) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *CategoryValue) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *CategoryValue) SetExtensionDefinitions(ref interface{}){
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
func (this *CategoryValue) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *CategoryValue) SetExtensionValues(ref interface{}){
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
func (this *CategoryValue) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *CategoryValue) SetDocumentation(ref interface{}){
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
func (this *CategoryValue) RemoveDocumentation(ref interface{}){
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

/** CategorizedFlowElements **/
func (this *CategoryValue) GetCategorizedFlowElements() []FlowElement{
	return this.m_categorizedFlowElements
}

/** Init reference CategorizedFlowElements **/
func (this *CategoryValue) SetCategorizedFlowElements(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_categorizedFlowElements); i++ {
			if this.M_categorizedFlowElements[i] == refStr {
				return
			}
		}
		this.M_categorizedFlowElements = append(this.M_categorizedFlowElements, ref.(string))
	}else{
		this.RemoveCategorizedFlowElements(ref)
		this.m_categorizedFlowElements = append(this.m_categorizedFlowElements, ref.(FlowElement))
		this.M_categorizedFlowElements = append(this.M_categorizedFlowElements, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CategorizedFlowElements **/
func (this *CategoryValue) RemoveCategorizedFlowElements(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	categorizedFlowElements_ := make([]FlowElement, 0)
	categorizedFlowElementsUuid := make([]string, 0)
	for i := 0; i < len(this.m_categorizedFlowElements); i++ {
		if toDelete.GetUUID() != this.m_categorizedFlowElements[i].(BaseElement).GetUUID() {
			categorizedFlowElements_ = append(categorizedFlowElements_, this.m_categorizedFlowElements[i])
			categorizedFlowElementsUuid = append(categorizedFlowElementsUuid, this.M_categorizedFlowElements[i])
		}
	}
	this.m_categorizedFlowElements = categorizedFlowElements_
	this.M_categorizedFlowElements = categorizedFlowElementsUuid
}

/** Value **/
func (this *CategoryValue) GetValue() string{
	return this.M_value
}

/** Init reference Value **/
func (this *CategoryValue) SetValue(ref interface{}){
	this.NeedSave = true
	this.M_value = ref.(string)
}

/** Remove reference Value **/

/** CategoryValueRef **/
func (this *CategoryValue) GetCategoryValueRefPtr() []*Group{
	return this.m_categoryValueRefPtr
}

/** Init reference CategoryValueRef **/
func (this *CategoryValue) SetCategoryValueRefPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_categoryValueRefPtr); i++ {
			if this.M_categoryValueRefPtr[i] == refStr {
				return
			}
		}
		this.M_categoryValueRefPtr = append(this.M_categoryValueRefPtr, ref.(string))
	}else{
		this.RemoveCategoryValueRefPtr(ref)
		this.m_categoryValueRefPtr = append(this.m_categoryValueRefPtr, ref.(*Group))
		this.M_categoryValueRefPtr = append(this.M_categoryValueRefPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CategoryValueRef **/
func (this *CategoryValue) RemoveCategoryValueRefPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	categoryValueRefPtr_ := make([]*Group, 0)
	categoryValueRefPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_categoryValueRefPtr); i++ {
		if toDelete.GetUUID() != this.m_categoryValueRefPtr[i].GetUUID() {
			categoryValueRefPtr_ = append(categoryValueRefPtr_, this.m_categoryValueRefPtr[i])
			categoryValueRefPtrUuid = append(categoryValueRefPtrUuid, this.M_categoryValueRefPtr[i])
		}
	}
	this.m_categoryValueRefPtr = categoryValueRefPtr_
	this.M_categoryValueRefPtr = categoryValueRefPtrUuid
}

/** Category **/
func (this *CategoryValue) GetCategoryPtr() *Category{
	return this.m_categoryPtr
}

/** Init reference Category **/
func (this *CategoryValue) SetCategoryPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_categoryPtr = ref.(string)
	}else{
		this.m_categoryPtr = ref.(*Category)
		this.M_categoryPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Category **/
func (this *CategoryValue) RemoveCategoryPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_categoryPtr.GetUUID() {
		this.m_categoryPtr = nil
		this.M_categoryPtr = ""
	}
}

/** Lane **/
func (this *CategoryValue) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *CategoryValue) SetLanePtr(ref interface{}){
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
func (this *CategoryValue) RemoveLanePtr(ref interface{}){
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
func (this *CategoryValue) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *CategoryValue) SetOutgoingPtr(ref interface{}){
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
func (this *CategoryValue) RemoveOutgoingPtr(ref interface{}){
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
func (this *CategoryValue) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *CategoryValue) SetIncomingPtr(ref interface{}){
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
func (this *CategoryValue) RemoveIncomingPtr(ref interface{}){
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
