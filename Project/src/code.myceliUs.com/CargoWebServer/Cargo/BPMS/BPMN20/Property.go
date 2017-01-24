// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Property struct {

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

	/** members of Property **/
	M_name string

	/** Associations **/
	m_processPtr *Process
	/** If the ref is a string and not an object **/
	M_processPtr string
	m_eventPtr   Event
	/** If the ref is a string and not an object **/
	M_eventPtr    string
	m_activityPtr Activity
	/** If the ref is a string and not an object **/
	M_activityPtr string
	m_lanePtr     []*Lane
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
	M_dataAssociationPtr                  []string
	m_multiInstanceLoopCharacteristicsPtr *MultiInstanceLoopCharacteristics
	/** If the ref is a string and not an object **/
	M_multiInstanceLoopCharacteristicsPtr string
}

/** Xml parser for Property **/
type XsdProperty struct {
	XMLName xml.Name `xml:"property"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_dataState      *XsdDataState `xml:"dataState,omitempty"`
	M_name           string        `xml:"name,attr"`
	M_itemSubjectRef string        `xml:"itemSubjectRef,attr"`
}

/** UUID **/
func (this *Property) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Property) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Property) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Property) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Property) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Property) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Property) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Property) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Property) SetExtensionDefinitions(ref interface{}) {
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
func (this *Property) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Property) SetExtensionValues(ref interface{}) {
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
func (this *Property) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Property) SetDocumentation(ref interface{}) {
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
func (this *Property) RemoveDocumentation(ref interface{}) {
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
func (this *Property) GetItemSubjectRef() *ItemDefinition {
	return this.m_itemSubjectRef
}

/** Init reference ItemSubjectRef **/
func (this *Property) SetItemSubjectRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_itemSubjectRef = ref.(string)
	} else {
		this.m_itemSubjectRef = ref.(*ItemDefinition)
		this.M_itemSubjectRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ItemSubjectRef **/
func (this *Property) RemoveItemSubjectRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_itemSubjectRef.GetUUID() {
		this.m_itemSubjectRef = nil
		this.M_itemSubjectRef = ""
	}
}

/** DataState **/
func (this *Property) GetDataState() *DataState {
	return this.M_dataState
}

/** Init reference DataState **/
func (this *Property) SetDataState(ref interface{}) {
	this.NeedSave = true
	this.M_dataState = ref.(*DataState)
}

/** Remove reference DataState **/
func (this *Property) RemoveDataState(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_dataState.GetUUID() {
		this.M_dataState = nil
	}
}

/** Name **/
func (this *Property) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Property) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Process **/
func (this *Property) GetProcessPtr() *Process {
	return this.m_processPtr
}

/** Init reference Process **/
func (this *Property) SetProcessPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_processPtr = ref.(string)
	} else {
		this.m_processPtr = ref.(*Process)
		this.M_processPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Process **/
func (this *Property) RemoveProcessPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_processPtr.GetUUID() {
		this.m_processPtr = nil
		this.M_processPtr = ""
	}
}

/** Event **/
func (this *Property) GetEventPtr() Event {
	return this.m_eventPtr
}

/** Init reference Event **/
func (this *Property) SetEventPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_eventPtr = ref.(string)
	} else {
		this.m_eventPtr = ref.(Event)
		this.M_eventPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Event **/
func (this *Property) RemoveEventPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_eventPtr.(BaseElement).GetUUID() {
		this.m_eventPtr = nil
		this.M_eventPtr = ""
	}
}

/** Activity **/
func (this *Property) GetActivityPtr() Activity {
	return this.m_activityPtr
}

/** Init reference Activity **/
func (this *Property) SetActivityPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_activityPtr = ref.(string)
	} else {
		this.m_activityPtr = ref.(Activity)
		this.M_activityPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Activity **/
func (this *Property) RemoveActivityPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_activityPtr.(BaseElement).GetUUID() {
		this.m_activityPtr = nil
		this.M_activityPtr = ""
	}
}

/** Lane **/
func (this *Property) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Property) SetLanePtr(ref interface{}) {
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
func (this *Property) RemoveLanePtr(ref interface{}) {
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
func (this *Property) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Property) SetOutgoingPtr(ref interface{}) {
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
func (this *Property) RemoveOutgoingPtr(ref interface{}) {
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
func (this *Property) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Property) SetIncomingPtr(ref interface{}) {
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
func (this *Property) RemoveIncomingPtr(ref interface{}) {
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
func (this *Property) GetDataAssociationPtr() []DataAssociation {
	return this.m_dataAssociationPtr
}

/** Init reference DataAssociation **/
func (this *Property) SetDataAssociationPtr(ref interface{}) {
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
func (this *Property) RemoveDataAssociationPtr(ref interface{}) {
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
func (this *Property) GetMultiInstanceLoopCharacteristicsPtr() *MultiInstanceLoopCharacteristics {
	return this.m_multiInstanceLoopCharacteristicsPtr
}

/** Init reference MultiInstanceLoopCharacteristics **/
func (this *Property) SetMultiInstanceLoopCharacteristicsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(string)
	} else {
		this.m_multiInstanceLoopCharacteristicsPtr = ref.(*MultiInstanceLoopCharacteristics)
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference MultiInstanceLoopCharacteristics **/
func (this *Property) RemoveMultiInstanceLoopCharacteristicsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_multiInstanceLoopCharacteristicsPtr.GetUUID() {
		this.m_multiInstanceLoopCharacteristicsPtr = nil
		this.M_multiInstanceLoopCharacteristicsPtr = ""
	}
}
