// +build BPMN20

package BPMN20

import(
"encoding/xml"
)

type CorrelationKey struct{

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

	/** members of CorrelationKey **/
	m_correlationPropertyRef []*CorrelationProperty
	/** If the ref is a string and not an object **/
	M_correlationPropertyRef []string
	M_name string


	/** Associations **/
	m_conversationNodePtr ConversationNode
	/** If the ref is a string and not an object **/
	M_conversationNodePtr string
	m_correlationSubscriptionPtr []*CorrelationSubscription
	/** If the ref is a string and not an object **/
	M_correlationSubscriptionPtr []string
	m_collaborationPtr Collaboration
	/** If the ref is a string and not an object **/
	M_collaborationPtr string
	m_choreographyActivityPtr ChoreographyActivity
	/** If the ref is a string and not an object **/
	M_choreographyActivityPtr string
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

/** Xml parser for CorrelationKey **/
type XsdCorrelationKey struct {
	XMLName xml.Name	`xml:"correlationKey"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	M_correlationPropertyRef	[]string	`xml:"correlationPropertyRef"`
	M_name	string	`xml:"name,attr"`

}
/** UUID **/
func (this *CorrelationKey) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *CorrelationKey) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *CorrelationKey) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *CorrelationKey) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *CorrelationKey) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *CorrelationKey) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *CorrelationKey) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/
func (this *CorrelationKey) RemoveExtensionElements(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ExtensionElements)
	if toDelete.GetUUID() == this.M_extensionElements.GetUUID() {
		this.M_extensionElements = nil
	}
}

/** ExtensionDefinitions **/
func (this *CorrelationKey) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *CorrelationKey) SetExtensionDefinitions(ref interface{}){
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
func (this *CorrelationKey) RemoveExtensionDefinitions(ref interface{}){
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
func (this *CorrelationKey) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *CorrelationKey) SetExtensionValues(ref interface{}){
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
func (this *CorrelationKey) RemoveExtensionValues(ref interface{}){
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
func (this *CorrelationKey) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *CorrelationKey) SetDocumentation(ref interface{}){
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
func (this *CorrelationKey) RemoveDocumentation(ref interface{}){
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

/** CorrelationPropertyRef **/
func (this *CorrelationKey) GetCorrelationPropertyRef() []*CorrelationProperty{
	return this.m_correlationPropertyRef
}

/** Init reference CorrelationPropertyRef **/
func (this *CorrelationKey) SetCorrelationPropertyRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_correlationPropertyRef); i++ {
			if this.M_correlationPropertyRef[i] == refStr {
				return
			}
		}
		this.M_correlationPropertyRef = append(this.M_correlationPropertyRef, ref.(string))
	}else{
		this.RemoveCorrelationPropertyRef(ref)
		this.m_correlationPropertyRef = append(this.m_correlationPropertyRef, ref.(*CorrelationProperty))
		this.M_correlationPropertyRef = append(this.M_correlationPropertyRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CorrelationPropertyRef **/
func (this *CorrelationKey) RemoveCorrelationPropertyRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	correlationPropertyRef_ := make([]*CorrelationProperty, 0)
	correlationPropertyRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_correlationPropertyRef); i++ {
		if toDelete.GetUUID() != this.m_correlationPropertyRef[i].GetUUID() {
			correlationPropertyRef_ = append(correlationPropertyRef_, this.m_correlationPropertyRef[i])
			correlationPropertyRefUuid = append(correlationPropertyRefUuid, this.M_correlationPropertyRef[i])
		}
	}
	this.m_correlationPropertyRef = correlationPropertyRef_
	this.M_correlationPropertyRef = correlationPropertyRefUuid
}

/** Name **/
func (this *CorrelationKey) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *CorrelationKey) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** ConversationNode **/
func (this *CorrelationKey) GetConversationNodePtr() ConversationNode{
	return this.m_conversationNodePtr
}

/** Init reference ConversationNode **/
func (this *CorrelationKey) SetConversationNodePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_conversationNodePtr = ref.(string)
	}else{
		this.m_conversationNodePtr = ref.(ConversationNode)
		this.M_conversationNodePtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ConversationNode **/
func (this *CorrelationKey) RemoveConversationNodePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_conversationNodePtr.(BaseElement).GetUUID() {
		this.m_conversationNodePtr = nil
		this.M_conversationNodePtr = ""
	}
}

/** CorrelationSubscription **/
func (this *CorrelationKey) GetCorrelationSubscriptionPtr() []*CorrelationSubscription{
	return this.m_correlationSubscriptionPtr
}

/** Init reference CorrelationSubscription **/
func (this *CorrelationKey) SetCorrelationSubscriptionPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_correlationSubscriptionPtr); i++ {
			if this.M_correlationSubscriptionPtr[i] == refStr {
				return
			}
		}
		this.M_correlationSubscriptionPtr = append(this.M_correlationSubscriptionPtr, ref.(string))
	}else{
		this.RemoveCorrelationSubscriptionPtr(ref)
		this.m_correlationSubscriptionPtr = append(this.m_correlationSubscriptionPtr, ref.(*CorrelationSubscription))
		this.M_correlationSubscriptionPtr = append(this.M_correlationSubscriptionPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CorrelationSubscription **/
func (this *CorrelationKey) RemoveCorrelationSubscriptionPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	correlationSubscriptionPtr_ := make([]*CorrelationSubscription, 0)
	correlationSubscriptionPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_correlationSubscriptionPtr); i++ {
		if toDelete.GetUUID() != this.m_correlationSubscriptionPtr[i].GetUUID() {
			correlationSubscriptionPtr_ = append(correlationSubscriptionPtr_, this.m_correlationSubscriptionPtr[i])
			correlationSubscriptionPtrUuid = append(correlationSubscriptionPtrUuid, this.M_correlationSubscriptionPtr[i])
		}
	}
	this.m_correlationSubscriptionPtr = correlationSubscriptionPtr_
	this.M_correlationSubscriptionPtr = correlationSubscriptionPtrUuid
}

/** Collaboration **/
func (this *CorrelationKey) GetCollaborationPtr() Collaboration{
	return this.m_collaborationPtr
}

/** Init reference Collaboration **/
func (this *CorrelationKey) SetCollaborationPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_collaborationPtr = ref.(string)
	}else{
		this.m_collaborationPtr = ref.(Collaboration)
		this.M_collaborationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Collaboration **/
func (this *CorrelationKey) RemoveCollaborationPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_collaborationPtr.(BaseElement).GetUUID() {
		this.m_collaborationPtr = nil
		this.M_collaborationPtr = ""
	}
}

/** ChoreographyActivity **/
func (this *CorrelationKey) GetChoreographyActivityPtr() ChoreographyActivity{
	return this.m_choreographyActivityPtr
}

/** Init reference ChoreographyActivity **/
func (this *CorrelationKey) SetChoreographyActivityPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_choreographyActivityPtr = ref.(string)
	}else{
		this.m_choreographyActivityPtr = ref.(ChoreographyActivity)
		this.M_choreographyActivityPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ChoreographyActivity **/
func (this *CorrelationKey) RemoveChoreographyActivityPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_choreographyActivityPtr.(BaseElement).GetUUID() {
		this.m_choreographyActivityPtr = nil
		this.M_choreographyActivityPtr = ""
	}
}

/** Lane **/
func (this *CorrelationKey) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *CorrelationKey) SetLanePtr(ref interface{}){
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
func (this *CorrelationKey) RemoveLanePtr(ref interface{}){
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
func (this *CorrelationKey) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *CorrelationKey) SetOutgoingPtr(ref interface{}){
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
func (this *CorrelationKey) RemoveOutgoingPtr(ref interface{}){
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
func (this *CorrelationKey) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *CorrelationKey) SetIncomingPtr(ref interface{}){
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
func (this *CorrelationKey) RemoveIncomingPtr(ref interface{}){
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
