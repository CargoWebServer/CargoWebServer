// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type DataInput struct {

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

	/** members of ItemAwareElement **/
	m_itemSubjectRef *ItemDefinition
	/** If the ref is a string and not an object **/
	M_itemSubjectRef string
	M_dataState      *DataState

	/** members of DataInput **/
	M_name         string
	M_isCollection bool
	m_inputSetRefs []*InputSet
	/** If the ref is a string and not an object **/
	M_inputSetRefs         []string
	m_inputSetWithOptional []*InputSet
	/** If the ref is a string and not an object **/
	M_inputSetWithOptional       []string
	m_inputSetWithWhileExecuting []*InputSet
	/** If the ref is a string and not an object **/
	M_inputSetWithWhileExecuting []string

	/** Associations **/
	m_throwEventPtr ThrowEvent
	/** If the ref is a string and not an object **/
	M_throwEventPtr               string
	m_inputOutputSpecificationPtr *InputOutputSpecification
	/** If the ref is a string and not an object **/
	M_inputOutputSpecificationPtr         string
	m_multiInstanceLoopCharacteristicsPtr *MultiInstanceLoopCharacteristics
	/** If the ref is a string and not an object **/
	M_multiInstanceLoopCharacteristicsPtr string
	m_lanePtr                             []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr        []string
	m_dataAssociationPtr []DataAssociation
	/** If the ref is a string and not an object **/
	M_dataAssociationPtr []string
}

/** Xml parser for DataInput **/
type XsdDataInput struct {
	XMLName xml.Name `xml:"dataInput"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_dataState      *XsdDataState `xml:"dataState,omitempty"`
	M_name           string        `xml:"name,attr"`
	M_itemSubjectRef string        `xml:"itemSubjectRef,attr"`
	M_isCollection   bool          `xml:"isCollection,attr"`
}

/** Alias Xsd parser **/

type XsdInputDataItem struct {
	XMLName xml.Name `xml:"inputDataItem"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_dataState      *XsdDataState `xml:"dataState,omitempty"`
	M_name           string        `xml:"name,attr"`
	M_itemSubjectRef string        `xml:"itemSubjectRef,attr"`
	M_isCollection   bool          `xml:"isCollection,attr"`
}

/** UUID **/
func (this *DataInput) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *DataInput) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *DataInput) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *DataInput) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *DataInput) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *DataInput) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *DataInput) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *DataInput) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *DataInput) SetExtensionDefinitions(ref interface{}) {
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
func (this *DataInput) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *DataInput) SetExtensionValues(ref interface{}) {
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
func (this *DataInput) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *DataInput) SetDocumentation(ref interface{}) {
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
func (this *DataInput) RemoveDocumentation(ref interface{}) {
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

/** ItemSubjectRef **/
func (this *DataInput) GetItemSubjectRef() *ItemDefinition {
	return this.m_itemSubjectRef
}

/** Init reference ItemSubjectRef **/
func (this *DataInput) SetItemSubjectRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_itemSubjectRef = ref.(string)
	} else {
		this.m_itemSubjectRef = ref.(*ItemDefinition)
		this.M_itemSubjectRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ItemSubjectRef **/
func (this *DataInput) RemoveItemSubjectRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_itemSubjectRef.GetUUID() {
		this.m_itemSubjectRef = nil
		this.M_itemSubjectRef = ""
	}
}

/** DataState **/
func (this *DataInput) GetDataState() *DataState {
	return this.M_dataState
}

/** Init reference DataState **/
func (this *DataInput) SetDataState(ref interface{}) {
	this.NeedSave = true
	this.M_dataState = ref.(*DataState)
}

/** Remove reference DataState **/
func (this *DataInput) RemoveDataState(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_dataState.GetUUID() {
		this.M_dataState = nil
	}
}

/** Name **/
func (this *DataInput) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *DataInput) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** IsCollection **/
func (this *DataInput) IsCollection() bool {
	return this.M_isCollection
}

/** Init reference IsCollection **/
func (this *DataInput) SetIsCollection(ref interface{}) {
	this.NeedSave = true
	this.M_isCollection = ref.(bool)
}

/** Remove reference IsCollection **/

/** InputSetRefs **/
func (this *DataInput) GetInputSetRefs() []*InputSet {
	return this.m_inputSetRefs
}

/** Init reference InputSetRefs **/
func (this *DataInput) SetInputSetRefs(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_inputSetRefs); i++ {
			if this.M_inputSetRefs[i] == refStr {
				return
			}
		}
		this.M_inputSetRefs = append(this.M_inputSetRefs, ref.(string))
	} else {
		this.RemoveInputSetRefs(ref)
		this.m_inputSetRefs = append(this.m_inputSetRefs, ref.(*InputSet))
		this.M_inputSetRefs = append(this.M_inputSetRefs, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference InputSetRefs **/
func (this *DataInput) RemoveInputSetRefs(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	inputSetRefs_ := make([]*InputSet, 0)
	inputSetRefsUuid := make([]string, 0)
	for i := 0; i < len(this.m_inputSetRefs); i++ {
		if toDelete.GetUUID() != this.m_inputSetRefs[i].GetUUID() {
			inputSetRefs_ = append(inputSetRefs_, this.m_inputSetRefs[i])
			inputSetRefsUuid = append(inputSetRefsUuid, this.M_inputSetRefs[i])
		}
	}
	this.m_inputSetRefs = inputSetRefs_
	this.M_inputSetRefs = inputSetRefsUuid
}

/** InputSetWithOptional **/
func (this *DataInput) GetInputSetWithOptional() []*InputSet {
	return this.m_inputSetWithOptional
}

/** Init reference InputSetWithOptional **/
func (this *DataInput) SetInputSetWithOptional(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_inputSetWithOptional); i++ {
			if this.M_inputSetWithOptional[i] == refStr {
				return
			}
		}
		this.M_inputSetWithOptional = append(this.M_inputSetWithOptional, ref.(string))
	} else {
		this.RemoveInputSetWithOptional(ref)
		this.m_inputSetWithOptional = append(this.m_inputSetWithOptional, ref.(*InputSet))
		this.M_inputSetWithOptional = append(this.M_inputSetWithOptional, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference InputSetWithOptional **/
func (this *DataInput) RemoveInputSetWithOptional(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	inputSetWithOptional_ := make([]*InputSet, 0)
	inputSetWithOptionalUuid := make([]string, 0)
	for i := 0; i < len(this.m_inputSetWithOptional); i++ {
		if toDelete.GetUUID() != this.m_inputSetWithOptional[i].GetUUID() {
			inputSetWithOptional_ = append(inputSetWithOptional_, this.m_inputSetWithOptional[i])
			inputSetWithOptionalUuid = append(inputSetWithOptionalUuid, this.M_inputSetWithOptional[i])
		}
	}
	this.m_inputSetWithOptional = inputSetWithOptional_
	this.M_inputSetWithOptional = inputSetWithOptionalUuid
}

/** InputSetWithWhileExecuting **/
func (this *DataInput) GetInputSetWithWhileExecuting() []*InputSet {
	return this.m_inputSetWithWhileExecuting
}

/** Init reference InputSetWithWhileExecuting **/
func (this *DataInput) SetInputSetWithWhileExecuting(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_inputSetWithWhileExecuting); i++ {
			if this.M_inputSetWithWhileExecuting[i] == refStr {
				return
			}
		}
		this.M_inputSetWithWhileExecuting = append(this.M_inputSetWithWhileExecuting, ref.(string))
	} else {
		this.RemoveInputSetWithWhileExecuting(ref)
		this.m_inputSetWithWhileExecuting = append(this.m_inputSetWithWhileExecuting, ref.(*InputSet))
		this.M_inputSetWithWhileExecuting = append(this.M_inputSetWithWhileExecuting, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference InputSetWithWhileExecuting **/
func (this *DataInput) RemoveInputSetWithWhileExecuting(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	inputSetWithWhileExecuting_ := make([]*InputSet, 0)
	inputSetWithWhileExecutingUuid := make([]string, 0)
	for i := 0; i < len(this.m_inputSetWithWhileExecuting); i++ {
		if toDelete.GetUUID() != this.m_inputSetWithWhileExecuting[i].GetUUID() {
			inputSetWithWhileExecuting_ = append(inputSetWithWhileExecuting_, this.m_inputSetWithWhileExecuting[i])
			inputSetWithWhileExecutingUuid = append(inputSetWithWhileExecutingUuid, this.M_inputSetWithWhileExecuting[i])
		}
	}
	this.m_inputSetWithWhileExecuting = inputSetWithWhileExecuting_
	this.M_inputSetWithWhileExecuting = inputSetWithWhileExecutingUuid
}

/** ThrowEvent **/
func (this *DataInput) GetThrowEventPtr() ThrowEvent {
	return this.m_throwEventPtr
}

/** Init reference ThrowEvent **/
func (this *DataInput) SetThrowEventPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_throwEventPtr = ref.(string)
	} else {
		this.m_throwEventPtr = ref.(ThrowEvent)
		this.M_throwEventPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ThrowEvent **/
func (this *DataInput) RemoveThrowEventPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_throwEventPtr.(BaseElement).GetUUID() {
		this.m_throwEventPtr = nil
		this.M_throwEventPtr = ""
	}
}

/** InputOutputSpecification **/
func (this *DataInput) GetInputOutputSpecificationPtr() *InputOutputSpecification {
	return this.m_inputOutputSpecificationPtr
}

/** Init reference InputOutputSpecification **/
func (this *DataInput) SetInputOutputSpecificationPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_inputOutputSpecificationPtr = ref.(string)
	} else {
		this.m_inputOutputSpecificationPtr = ref.(*InputOutputSpecification)
		this.M_inputOutputSpecificationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference InputOutputSpecification **/
func (this *DataInput) RemoveInputOutputSpecificationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_inputOutputSpecificationPtr.GetUUID() {
		this.m_inputOutputSpecificationPtr = nil
		this.M_inputOutputSpecificationPtr = ""
	}
}

/** MultiInstanceLoopCharacteristics **/
func (this *DataInput) GetMultiInstanceLoopCharacteristicsPtr() *MultiInstanceLoopCharacteristics {
	return this.m_multiInstanceLoopCharacteristicsPtr
}

/** Init reference MultiInstanceLoopCharacteristics **/
func (this *DataInput) SetMultiInstanceLoopCharacteristicsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(string)
	} else {
		this.m_multiInstanceLoopCharacteristicsPtr = ref.(*MultiInstanceLoopCharacteristics)
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference MultiInstanceLoopCharacteristics **/
func (this *DataInput) RemoveMultiInstanceLoopCharacteristicsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_multiInstanceLoopCharacteristicsPtr.GetUUID() {
		this.m_multiInstanceLoopCharacteristicsPtr = nil
		this.M_multiInstanceLoopCharacteristicsPtr = ""
	}
}

/** Lane **/
func (this *DataInput) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *DataInput) SetLanePtr(ref interface{}) {
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
func (this *DataInput) RemoveLanePtr(ref interface{}) {
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
func (this *DataInput) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *DataInput) SetOutgoingPtr(ref interface{}) {
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
func (this *DataInput) RemoveOutgoingPtr(ref interface{}) {
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
func (this *DataInput) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *DataInput) SetIncomingPtr(ref interface{}) {
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
func (this *DataInput) RemoveIncomingPtr(ref interface{}) {
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

/** DataAssociation **/
func (this *DataInput) GetDataAssociationPtr() []DataAssociation {
	return this.m_dataAssociationPtr
}

/** Init reference DataAssociation **/
func (this *DataInput) SetDataAssociationPtr(ref interface{}) {
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
func (this *DataInput) RemoveDataAssociationPtr(ref interface{}) {
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
