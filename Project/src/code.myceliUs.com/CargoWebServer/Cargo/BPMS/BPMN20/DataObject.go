// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type DataObject struct {

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

	/** members of FlowElement **/
	M_name             string
	M_auditing         *Auditing
	M_monitoring       *Monitoring
	m_categoryValueRef []*CategoryValue
	/** If the ref is a string and not an object **/
	M_categoryValueRef []string

	/** members of ItemAwareElement **/
	m_itemSubjectRef *ItemDefinition
	/** If the ref is a string and not an object **/
	M_itemSubjectRef string
	M_dataState      *DataState

	/** members of DataObject **/
	M_isCollection bool

	/** Associations **/
	m_dataObjectPtr []*DataObjectReference
	/** If the ref is a string and not an object **/
	M_dataObjectPtr []string
	m_lanePtr       []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr  []string
	m_containerPtr FlowElementsContainer
	/** If the ref is a string and not an object **/
	M_containerPtr       string
	m_dataAssociationPtr []DataAssociation
	/** If the ref is a string and not an object **/
	M_dataAssociationPtr                  []string
	m_multiInstanceLoopCharacteristicsPtr *MultiInstanceLoopCharacteristics
	/** If the ref is a string and not an object **/
	M_multiInstanceLoopCharacteristicsPtr string
}

/** Xml parser for DataObject **/
type XsdDataObject struct {
	XMLName xml.Name `xml:"dataObject"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	/** FlowElement **/
	M_auditing         *XsdAuditing   `xml:"auditing,omitempty"`
	M_monitoring       *XsdMonitoring `xml:"monitoring,omitempty"`
	M_categoryValueRef []string       `xml:"categoryValueRef"`
	M_name             string         `xml:"name,attr"`

	M_dataState      *XsdDataState `xml:"dataState,omitempty"`
	M_itemSubjectRef string        `xml:"itemSubjectRef,attr"`
	M_isCollection   bool          `xml:"isCollection,attr"`
}

/** UUID **/
func (this *DataObject) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *DataObject) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *DataObject) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *DataObject) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *DataObject) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *DataObject) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *DataObject) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *DataObject) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *DataObject) SetExtensionDefinitions(ref interface{}) {
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
func (this *DataObject) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *DataObject) SetExtensionValues(ref interface{}) {
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
func (this *DataObject) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *DataObject) SetDocumentation(ref interface{}) {
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
func (this *DataObject) RemoveDocumentation(ref interface{}) {
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
func (this *DataObject) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *DataObject) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Auditing **/
func (this *DataObject) GetAuditing() *Auditing {
	return this.M_auditing
}

/** Init reference Auditing **/
func (this *DataObject) SetAuditing(ref interface{}) {
	this.NeedSave = true
	this.M_auditing = ref.(*Auditing)
}

/** Remove reference Auditing **/
func (this *DataObject) RemoveAuditing(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_auditing.GetUUID() {
		this.M_auditing = nil
	}
}

/** Monitoring **/
func (this *DataObject) GetMonitoring() *Monitoring {
	return this.M_monitoring
}

/** Init reference Monitoring **/
func (this *DataObject) SetMonitoring(ref interface{}) {
	this.NeedSave = true
	this.M_monitoring = ref.(*Monitoring)
}

/** Remove reference Monitoring **/
func (this *DataObject) RemoveMonitoring(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_monitoring.GetUUID() {
		this.M_monitoring = nil
	}
}

/** CategoryValueRef **/
func (this *DataObject) GetCategoryValueRef() []*CategoryValue {
	return this.m_categoryValueRef
}

/** Init reference CategoryValueRef **/
func (this *DataObject) SetCategoryValueRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_categoryValueRef); i++ {
			if this.M_categoryValueRef[i] == refStr {
				return
			}
		}
		this.M_categoryValueRef = append(this.M_categoryValueRef, ref.(string))
	} else {
		this.RemoveCategoryValueRef(ref)
		this.m_categoryValueRef = append(this.m_categoryValueRef, ref.(*CategoryValue))
		this.M_categoryValueRef = append(this.M_categoryValueRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CategoryValueRef **/
func (this *DataObject) RemoveCategoryValueRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	categoryValueRef_ := make([]*CategoryValue, 0)
	categoryValueRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_categoryValueRef); i++ {
		if toDelete.GetUUID() != this.m_categoryValueRef[i].GetUUID() {
			categoryValueRef_ = append(categoryValueRef_, this.m_categoryValueRef[i])
			categoryValueRefUuid = append(categoryValueRefUuid, this.M_categoryValueRef[i])
		}
	}
	this.m_categoryValueRef = categoryValueRef_
	this.M_categoryValueRef = categoryValueRefUuid
}

/** ItemSubjectRef **/
func (this *DataObject) GetItemSubjectRef() *ItemDefinition {
	return this.m_itemSubjectRef
}

/** Init reference ItemSubjectRef **/
func (this *DataObject) SetItemSubjectRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_itemSubjectRef = ref.(string)
	} else {
		this.m_itemSubjectRef = ref.(*ItemDefinition)
		this.M_itemSubjectRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ItemSubjectRef **/
func (this *DataObject) RemoveItemSubjectRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_itemSubjectRef.GetUUID() {
		this.m_itemSubjectRef = nil
		this.M_itemSubjectRef = ""
	}
}

/** DataState **/
func (this *DataObject) GetDataState() *DataState {
	return this.M_dataState
}

/** Init reference DataState **/
func (this *DataObject) SetDataState(ref interface{}) {
	this.NeedSave = true
	this.M_dataState = ref.(*DataState)
}

/** Remove reference DataState **/
func (this *DataObject) RemoveDataState(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_dataState.GetUUID() {
		this.M_dataState = nil
	}
}

/** IsCollection **/
func (this *DataObject) IsCollection() bool {
	return this.M_isCollection
}

/** Init reference IsCollection **/
func (this *DataObject) SetIsCollection(ref interface{}) {
	this.NeedSave = true
	this.M_isCollection = ref.(bool)
}

/** Remove reference IsCollection **/

/** DataObject **/
func (this *DataObject) GetDataObjectPtr() []*DataObjectReference {
	return this.m_dataObjectPtr
}

/** Init reference DataObject **/
func (this *DataObject) SetDataObjectPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_dataObjectPtr); i++ {
			if this.M_dataObjectPtr[i] == refStr {
				return
			}
		}
		this.M_dataObjectPtr = append(this.M_dataObjectPtr, ref.(string))
	} else {
		this.RemoveDataObjectPtr(ref)
		this.m_dataObjectPtr = append(this.m_dataObjectPtr, ref.(*DataObjectReference))
		this.M_dataObjectPtr = append(this.M_dataObjectPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference DataObject **/
func (this *DataObject) RemoveDataObjectPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataObjectPtr_ := make([]*DataObjectReference, 0)
	dataObjectPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_dataObjectPtr); i++ {
		if toDelete.GetUUID() != this.m_dataObjectPtr[i].GetUUID() {
			dataObjectPtr_ = append(dataObjectPtr_, this.m_dataObjectPtr[i])
			dataObjectPtrUuid = append(dataObjectPtrUuid, this.M_dataObjectPtr[i])
		}
	}
	this.m_dataObjectPtr = dataObjectPtr_
	this.M_dataObjectPtr = dataObjectPtrUuid
}

/** Lane **/
func (this *DataObject) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *DataObject) SetLanePtr(ref interface{}) {
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
func (this *DataObject) RemoveLanePtr(ref interface{}) {
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
func (this *DataObject) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *DataObject) SetOutgoingPtr(ref interface{}) {
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
func (this *DataObject) RemoveOutgoingPtr(ref interface{}) {
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
func (this *DataObject) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *DataObject) SetIncomingPtr(ref interface{}) {
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
func (this *DataObject) RemoveIncomingPtr(ref interface{}) {
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

/** Container **/
func (this *DataObject) GetContainerPtr() FlowElementsContainer {
	return this.m_containerPtr
}

/** Init reference Container **/
func (this *DataObject) SetContainerPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_containerPtr = ref.(string)
	} else {
		this.m_containerPtr = ref.(FlowElementsContainer)
		this.M_containerPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Container **/
func (this *DataObject) RemoveContainerPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_containerPtr.(BaseElement).GetUUID() {
		this.m_containerPtr = nil
		this.M_containerPtr = ""
	}
}

/** DataAssociation **/
func (this *DataObject) GetDataAssociationPtr() []DataAssociation {
	return this.m_dataAssociationPtr
}

/** Init reference DataAssociation **/
func (this *DataObject) SetDataAssociationPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_dataAssociationPtr); i++ {
			if this.M_dataAssociationPtr[i] == refStr {
				return
			}
		}
		this.M_dataAssociationPtr = append(this.M_dataAssociationPtr, ref.(string))
	} else {
		this.RemoveDataAssociationPtr(ref)
		this.m_dataAssociationPtr = append(this.m_dataAssociationPtr, ref.(DataAssociation))
		this.M_dataAssociationPtr = append(this.M_dataAssociationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference DataAssociation **/
func (this *DataObject) RemoveDataAssociationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataAssociationPtr_ := make([]DataAssociation, 0)
	dataAssociationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_dataAssociationPtr); i++ {
		if toDelete.GetUUID() != this.m_dataAssociationPtr[i].(BaseElement).GetUUID() {
			dataAssociationPtr_ = append(dataAssociationPtr_, this.m_dataAssociationPtr[i])
			dataAssociationPtrUuid = append(dataAssociationPtrUuid, this.M_dataAssociationPtr[i])
		}
	}
	this.m_dataAssociationPtr = dataAssociationPtr_
	this.M_dataAssociationPtr = dataAssociationPtrUuid
}

/** MultiInstanceLoopCharacteristics **/
func (this *DataObject) GetMultiInstanceLoopCharacteristicsPtr() *MultiInstanceLoopCharacteristics {
	return this.m_multiInstanceLoopCharacteristicsPtr
}

/** Init reference MultiInstanceLoopCharacteristics **/
func (this *DataObject) SetMultiInstanceLoopCharacteristicsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(string)
	} else {
		this.m_multiInstanceLoopCharacteristicsPtr = ref.(*MultiInstanceLoopCharacteristics)
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference MultiInstanceLoopCharacteristics **/
func (this *DataObject) RemoveMultiInstanceLoopCharacteristicsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_multiInstanceLoopCharacteristicsPtr.GetUUID() {
		this.m_multiInstanceLoopCharacteristicsPtr = nil
		this.M_multiInstanceLoopCharacteristicsPtr = ""
	}
}
