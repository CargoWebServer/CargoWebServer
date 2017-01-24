// +build BPMN
package BPMN20

type InteractionNode_impl struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of BaseElement **/
	M_id    string
	m_other interface{}
	/** If the ref is a string and not an object **/
	M_other                string
	M_extensionElements    *ExtensionElements
	M_extensionDefinitions []*ExtensionDefinition
	M_extensionValues      []*ExtensionAttributeValue
	M_documentation        []*Documentation

	/** members of InteractionNode **/
	m_incomingConversationLinks []*ConversationLink
	/** If the ref is a string and not an object **/
	M_incomingConversationLinks []string
	m_outgoingConversationLinks []*ConversationLink
	/** If the ref is a string and not an object **/
	M_outgoingConversationLinks []string

	/** Associations **/
	m_messageFlowPtr []*MessageFlow
	/** If the ref is a string and not an object **/
	M_messageFlowPtr []string
	m_lanePtr        []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
} /** UUID **/
func (this *InteractionNode_impl) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *InteractionNode_impl) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *InteractionNode_impl) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *InteractionNode_impl) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *InteractionNode_impl) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *InteractionNode_impl) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *InteractionNode_impl) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *InteractionNode_impl) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *InteractionNode_impl) SetExtensionDefinitions(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var extensionDefinitionss []*ExtensionDefinition
	for i := 0; i < len(this.M_extensionDefinitions); i++ {
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
func (this *InteractionNode_impl) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *InteractionNode_impl) SetExtensionValues(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var extensionValuess []*ExtensionAttributeValue
	for i := 0; i < len(this.M_extensionValues); i++ {
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
func (this *InteractionNode_impl) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *InteractionNode_impl) SetDocumentation(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var documentations []*Documentation
	for i := 0; i < len(this.M_documentation); i++ {
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
func (this *InteractionNode_impl) RemoveDocumentation(ref interface{}) {
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

/** IncomingConversationLinks **/
func (this *InteractionNode_impl) GetIncomingConversationLinks() []*ConversationLink {
	return this.m_incomingConversationLinks
}

/** Init reference IncomingConversationLinks **/
func (this *InteractionNode_impl) SetIncomingConversationLinks(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_incomingConversationLinks); i++ {
			if this.M_incomingConversationLinks[i] == refStr {
				return
			}
		}
		this.M_incomingConversationLinks = append(this.M_incomingConversationLinks, ref.(string))
	} else {
		this.RemoveIncomingConversationLinks(ref)
		this.m_incomingConversationLinks = append(this.m_incomingConversationLinks, ref.(*ConversationLink))
		this.M_incomingConversationLinks = append(this.M_incomingConversationLinks, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference IncomingConversationLinks **/
func (this *InteractionNode_impl) RemoveIncomingConversationLinks(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	incomingConversationLinks_ := make([]*ConversationLink, 0)
	incomingConversationLinksUuid := make([]string, 0)
	for i := 0; i < len(this.m_incomingConversationLinks); i++ {
		if toDelete.GetUUID() != this.m_incomingConversationLinks[i].GetUUID() {
			incomingConversationLinks_ = append(incomingConversationLinks_, this.m_incomingConversationLinks[i])
			incomingConversationLinksUuid = append(incomingConversationLinksUuid, this.M_incomingConversationLinks[i])
		}
	}
	this.m_incomingConversationLinks = incomingConversationLinks_
	this.M_incomingConversationLinks = incomingConversationLinksUuid
}

/** OutgoingConversationLinks **/
func (this *InteractionNode_impl) GetOutgoingConversationLinks() []*ConversationLink {
	return this.m_outgoingConversationLinks
}

/** Init reference OutgoingConversationLinks **/
func (this *InteractionNode_impl) SetOutgoingConversationLinks(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_outgoingConversationLinks); i++ {
			if this.M_outgoingConversationLinks[i] == refStr {
				return
			}
		}
		this.M_outgoingConversationLinks = append(this.M_outgoingConversationLinks, ref.(string))
	} else {
		this.RemoveOutgoingConversationLinks(ref)
		this.m_outgoingConversationLinks = append(this.m_outgoingConversationLinks, ref.(*ConversationLink))
		this.M_outgoingConversationLinks = append(this.M_outgoingConversationLinks, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference OutgoingConversationLinks **/
func (this *InteractionNode_impl) RemoveOutgoingConversationLinks(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	outgoingConversationLinks_ := make([]*ConversationLink, 0)
	outgoingConversationLinksUuid := make([]string, 0)
	for i := 0; i < len(this.m_outgoingConversationLinks); i++ {
		if toDelete.GetUUID() != this.m_outgoingConversationLinks[i].GetUUID() {
			outgoingConversationLinks_ = append(outgoingConversationLinks_, this.m_outgoingConversationLinks[i])
			outgoingConversationLinksUuid = append(outgoingConversationLinksUuid, this.M_outgoingConversationLinks[i])
		}
	}
	this.m_outgoingConversationLinks = outgoingConversationLinks_
	this.M_outgoingConversationLinks = outgoingConversationLinksUuid
}

/** MessageFlow **/
func (this *InteractionNode_impl) GetMessageFlowPtr() []*MessageFlow {
	return this.m_messageFlowPtr
}

/** Init reference MessageFlow **/
func (this *InteractionNode_impl) SetMessageFlowPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_messageFlowPtr); i++ {
			if this.M_messageFlowPtr[i] == refStr {
				return
			}
		}
		this.M_messageFlowPtr = append(this.M_messageFlowPtr, ref.(string))
	} else {
		this.RemoveMessageFlowPtr(ref)
		this.m_messageFlowPtr = append(this.m_messageFlowPtr, ref.(*MessageFlow))
		this.M_messageFlowPtr = append(this.M_messageFlowPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference MessageFlow **/
func (this *InteractionNode_impl) RemoveMessageFlowPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	messageFlowPtr_ := make([]*MessageFlow, 0)
	messageFlowPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_messageFlowPtr); i++ {
		if toDelete.GetUUID() != this.m_messageFlowPtr[i].GetUUID() {
			messageFlowPtr_ = append(messageFlowPtr_, this.m_messageFlowPtr[i])
			messageFlowPtrUuid = append(messageFlowPtrUuid, this.M_messageFlowPtr[i])
		}
	}
	this.m_messageFlowPtr = messageFlowPtr_
	this.M_messageFlowPtr = messageFlowPtrUuid
}

/** Lane **/
func (this *InteractionNode_impl) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *InteractionNode_impl) SetLanePtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_lanePtr); i++ {
			if this.M_lanePtr[i] == refStr {
				return
			}
		}
		this.M_lanePtr = append(this.M_lanePtr, ref.(string))
	} else {
		this.RemoveLanePtr(ref)
		this.m_lanePtr = append(this.m_lanePtr, ref.(*Lane))
		this.M_lanePtr = append(this.M_lanePtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Lane **/
func (this *InteractionNode_impl) RemoveLanePtr(ref interface{}) {
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
func (this *InteractionNode_impl) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *InteractionNode_impl) SetOutgoingPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_outgoingPtr); i++ {
			if this.M_outgoingPtr[i] == refStr {
				return
			}
		}
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(string))
	} else {
		this.RemoveOutgoingPtr(ref)
		this.m_outgoingPtr = append(this.m_outgoingPtr, ref.(*Association))
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Outgoing **/
func (this *InteractionNode_impl) RemoveOutgoingPtr(ref interface{}) {
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
func (this *InteractionNode_impl) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *InteractionNode_impl) SetIncomingPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_incomingPtr); i++ {
			if this.M_incomingPtr[i] == refStr {
				return
			}
		}
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(string))
	} else {
		this.RemoveIncomingPtr(ref)
		this.m_incomingPtr = append(this.m_incomingPtr, ref.(*Association))
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Incoming **/
func (this *InteractionNode_impl) RemoveIncomingPtr(ref interface{}) {
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
