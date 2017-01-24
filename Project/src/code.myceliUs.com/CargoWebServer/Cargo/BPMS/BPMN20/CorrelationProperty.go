// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type CorrelationProperty struct {

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

	/** members of RootElement **/
	/** No members **/

	/** members of CorrelationProperty **/
	M_correlationPropertyRetrievalExpression []*CorrelationPropertyRetrievalExpression
	M_name                                   string
	m_type                                   *ItemDefinition
	/** If the ref is a string and not an object **/
	M_type string

	/** Associations **/
	m_correlationKeyPtr []*CorrelationKey
	/** If the ref is a string and not an object **/
	M_correlationKeyPtr             []string
	m_correlationPropertyBindingPtr []*CorrelationPropertyBinding
	/** If the ref is a string and not an object **/
	M_correlationPropertyBindingPtr []string
	m_lanePtr                       []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr    []string
	m_definitionsPtr *Definitions
	/** If the ref is a string and not an object **/
	M_definitionsPtr string
}

/** Xml parser for CorrelationProperty **/
type XsdCorrelationProperty struct {
	XMLName xml.Name `xml:"correlationProperty"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	/** RootElement **/

	M_correlationPropertyRetrievalExpression []*XsdCorrelationPropertyRetrievalExpression `xml:"correlationPropertyRetrievalExpression,omitempty"`
	M_name                                   string                                       `xml:"name,attr"`
	M_type                                   string                                       `xml:"type,attr"`
}

/** UUID **/
func (this *CorrelationProperty) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *CorrelationProperty) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *CorrelationProperty) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *CorrelationProperty) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *CorrelationProperty) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *CorrelationProperty) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *CorrelationProperty) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *CorrelationProperty) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *CorrelationProperty) SetExtensionDefinitions(ref interface{}) {
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
func (this *CorrelationProperty) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *CorrelationProperty) SetExtensionValues(ref interface{}) {
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
func (this *CorrelationProperty) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *CorrelationProperty) SetDocumentation(ref interface{}) {
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
func (this *CorrelationProperty) RemoveDocumentation(ref interface{}) {
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

/** CorrelationPropertyRetrievalExpression **/
func (this *CorrelationProperty) GetCorrelationPropertyRetrievalExpression() []*CorrelationPropertyRetrievalExpression {
	return this.M_correlationPropertyRetrievalExpression
}

/** Init reference CorrelationPropertyRetrievalExpression **/
func (this *CorrelationProperty) SetCorrelationPropertyRetrievalExpression(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var correlationPropertyRetrievalExpressions []*CorrelationPropertyRetrievalExpression
	for i := 0; i < len(this.M_correlationPropertyRetrievalExpression); i++ {
		if this.M_correlationPropertyRetrievalExpression[i].GetUUID() != ref.(BaseElement).GetUUID() {
			correlationPropertyRetrievalExpressions = append(correlationPropertyRetrievalExpressions, this.M_correlationPropertyRetrievalExpression[i])
		} else {
			isExist = true
			correlationPropertyRetrievalExpressions = append(correlationPropertyRetrievalExpressions, ref.(*CorrelationPropertyRetrievalExpression))
		}
	}
	if !isExist {
		correlationPropertyRetrievalExpressions = append(correlationPropertyRetrievalExpressions, ref.(*CorrelationPropertyRetrievalExpression))
	}
	this.M_correlationPropertyRetrievalExpression = correlationPropertyRetrievalExpressions
}

/** Remove reference CorrelationPropertyRetrievalExpression **/
func (this *CorrelationProperty) RemoveCorrelationPropertyRetrievalExpression(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	correlationPropertyRetrievalExpression_ := make([]*CorrelationPropertyRetrievalExpression, 0)
	for i := 0; i < len(this.M_correlationPropertyRetrievalExpression); i++ {
		if toDelete.GetUUID() != this.M_correlationPropertyRetrievalExpression[i].GetUUID() {
			correlationPropertyRetrievalExpression_ = append(correlationPropertyRetrievalExpression_, this.M_correlationPropertyRetrievalExpression[i])
		}
	}
	this.M_correlationPropertyRetrievalExpression = correlationPropertyRetrievalExpression_
}

/** Name **/
func (this *CorrelationProperty) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *CorrelationProperty) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Type **/
func (this *CorrelationProperty) GetType() *ItemDefinition {
	return this.m_type
}

/** Init reference Type **/
func (this *CorrelationProperty) SetType(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_type = ref.(string)
	} else {
		this.m_type = ref.(*ItemDefinition)
		this.M_type = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Type **/
func (this *CorrelationProperty) RemoveType(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_type.GetUUID() {
		this.m_type = nil
		this.M_type = ""
	}
}

/** CorrelationKey **/
func (this *CorrelationProperty) GetCorrelationKeyPtr() []*CorrelationKey {
	return this.m_correlationKeyPtr
}

/** Init reference CorrelationKey **/
func (this *CorrelationProperty) SetCorrelationKeyPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_correlationKeyPtr); i++ {
			if this.M_correlationKeyPtr[i] == refStr {
				return
			}
		}
		this.M_correlationKeyPtr = append(this.M_correlationKeyPtr, ref.(string))
	} else {
		this.RemoveCorrelationKeyPtr(ref)
		this.m_correlationKeyPtr = append(this.m_correlationKeyPtr, ref.(*CorrelationKey))
		this.M_correlationKeyPtr = append(this.M_correlationKeyPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CorrelationKey **/
func (this *CorrelationProperty) RemoveCorrelationKeyPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	correlationKeyPtr_ := make([]*CorrelationKey, 0)
	correlationKeyPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_correlationKeyPtr); i++ {
		if toDelete.GetUUID() != this.m_correlationKeyPtr[i].GetUUID() {
			correlationKeyPtr_ = append(correlationKeyPtr_, this.m_correlationKeyPtr[i])
			correlationKeyPtrUuid = append(correlationKeyPtrUuid, this.M_correlationKeyPtr[i])
		}
	}
	this.m_correlationKeyPtr = correlationKeyPtr_
	this.M_correlationKeyPtr = correlationKeyPtrUuid
}

/** CorrelationPropertyBinding **/
func (this *CorrelationProperty) GetCorrelationPropertyBindingPtr() []*CorrelationPropertyBinding {
	return this.m_correlationPropertyBindingPtr
}

/** Init reference CorrelationPropertyBinding **/
func (this *CorrelationProperty) SetCorrelationPropertyBindingPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_correlationPropertyBindingPtr); i++ {
			if this.M_correlationPropertyBindingPtr[i] == refStr {
				return
			}
		}
		this.M_correlationPropertyBindingPtr = append(this.M_correlationPropertyBindingPtr, ref.(string))
	} else {
		this.RemoveCorrelationPropertyBindingPtr(ref)
		this.m_correlationPropertyBindingPtr = append(this.m_correlationPropertyBindingPtr, ref.(*CorrelationPropertyBinding))
		this.M_correlationPropertyBindingPtr = append(this.M_correlationPropertyBindingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CorrelationPropertyBinding **/
func (this *CorrelationProperty) RemoveCorrelationPropertyBindingPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	correlationPropertyBindingPtr_ := make([]*CorrelationPropertyBinding, 0)
	correlationPropertyBindingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_correlationPropertyBindingPtr); i++ {
		if toDelete.GetUUID() != this.m_correlationPropertyBindingPtr[i].GetUUID() {
			correlationPropertyBindingPtr_ = append(correlationPropertyBindingPtr_, this.m_correlationPropertyBindingPtr[i])
			correlationPropertyBindingPtrUuid = append(correlationPropertyBindingPtrUuid, this.M_correlationPropertyBindingPtr[i])
		}
	}
	this.m_correlationPropertyBindingPtr = correlationPropertyBindingPtr_
	this.M_correlationPropertyBindingPtr = correlationPropertyBindingPtrUuid
}

/** Lane **/
func (this *CorrelationProperty) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *CorrelationProperty) SetLanePtr(ref interface{}) {
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
func (this *CorrelationProperty) RemoveLanePtr(ref interface{}) {
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
func (this *CorrelationProperty) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *CorrelationProperty) SetOutgoingPtr(ref interface{}) {
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
func (this *CorrelationProperty) RemoveOutgoingPtr(ref interface{}) {
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
func (this *CorrelationProperty) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *CorrelationProperty) SetIncomingPtr(ref interface{}) {
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
func (this *CorrelationProperty) RemoveIncomingPtr(ref interface{}) {
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
func (this *CorrelationProperty) GetDefinitionsPtr() *Definitions {
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *CorrelationProperty) SetDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	} else {
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *CorrelationProperty) RemoveDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}
