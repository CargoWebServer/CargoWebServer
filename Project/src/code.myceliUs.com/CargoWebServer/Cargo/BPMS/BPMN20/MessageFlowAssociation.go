package BPMN20

import(
"encoding/xml"
)

type MessageFlowAssociation struct{

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

	/** members of MessageFlowAssociation **/
	m_innerMessageFlowRef *MessageFlow
	/** If the ref is a string and not an object **/
	M_innerMessageFlowRef string
	m_outerMessageFlowRef *MessageFlow
	/** If the ref is a string and not an object **/
	M_outerMessageFlowRef string


	/** Associations **/
	m_collaborationPtr Collaboration
	/** If the ref is a string and not an object **/
	M_collaborationPtr string
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

/** Xml parser for MessageFlowAssociation **/
type XsdMessageFlowAssociation struct {
	XMLName xml.Name	`xml:"messageFlowAssociation"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	M_innerMessageFlowRef	string	`xml:"innerMessageFlowRef,attr"`
	M_outerMessageFlowRef	string	`xml:"outerMessageFlowRef,attr"`

}
/** UUID **/
func (this *MessageFlowAssociation) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *MessageFlowAssociation) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *MessageFlowAssociation) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *MessageFlowAssociation) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *MessageFlowAssociation) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *MessageFlowAssociation) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *MessageFlowAssociation) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *MessageFlowAssociation) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *MessageFlowAssociation) SetExtensionDefinitions(ref interface{}){
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
func (this *MessageFlowAssociation) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *MessageFlowAssociation) SetExtensionValues(ref interface{}){
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
func (this *MessageFlowAssociation) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *MessageFlowAssociation) SetDocumentation(ref interface{}){
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
func (this *MessageFlowAssociation) RemoveDocumentation(ref interface{}){
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

/** InnerMessageFlowRef **/
func (this *MessageFlowAssociation) GetInnerMessageFlowRef() *MessageFlow{
	return this.m_innerMessageFlowRef
}

/** Init reference InnerMessageFlowRef **/
func (this *MessageFlowAssociation) SetInnerMessageFlowRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_innerMessageFlowRef = ref.(string)
	}else{
		this.m_innerMessageFlowRef = ref.(*MessageFlow)
		this.M_innerMessageFlowRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference InnerMessageFlowRef **/
func (this *MessageFlowAssociation) RemoveInnerMessageFlowRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_innerMessageFlowRef.GetUUID() {
		this.m_innerMessageFlowRef = nil
		this.M_innerMessageFlowRef = ""
	}
}

/** OuterMessageFlowRef **/
func (this *MessageFlowAssociation) GetOuterMessageFlowRef() *MessageFlow{
	return this.m_outerMessageFlowRef
}

/** Init reference OuterMessageFlowRef **/
func (this *MessageFlowAssociation) SetOuterMessageFlowRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_outerMessageFlowRef = ref.(string)
	}else{
		this.m_outerMessageFlowRef = ref.(*MessageFlow)
		this.M_outerMessageFlowRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference OuterMessageFlowRef **/
func (this *MessageFlowAssociation) RemoveOuterMessageFlowRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_outerMessageFlowRef.GetUUID() {
		this.m_outerMessageFlowRef = nil
		this.M_outerMessageFlowRef = ""
	}
}

/** Collaboration **/
func (this *MessageFlowAssociation) GetCollaborationPtr() Collaboration{
	return this.m_collaborationPtr
}

/** Init reference Collaboration **/
func (this *MessageFlowAssociation) SetCollaborationPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_collaborationPtr = ref.(string)
	}else{
		this.m_collaborationPtr = ref.(Collaboration)
		this.M_collaborationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Collaboration **/
func (this *MessageFlowAssociation) RemoveCollaborationPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_collaborationPtr.(BaseElement).GetUUID() {
		this.m_collaborationPtr = nil
		this.M_collaborationPtr = ""
	}
}

/** Lane **/
func (this *MessageFlowAssociation) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *MessageFlowAssociation) SetLanePtr(ref interface{}){
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
func (this *MessageFlowAssociation) RemoveLanePtr(ref interface{}){
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
func (this *MessageFlowAssociation) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *MessageFlowAssociation) SetOutgoingPtr(ref interface{}){
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
func (this *MessageFlowAssociation) RemoveOutgoingPtr(ref interface{}){
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
func (this *MessageFlowAssociation) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *MessageFlowAssociation) SetIncomingPtr(ref interface{}){
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
func (this *MessageFlowAssociation) RemoveIncomingPtr(ref interface{}){
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
