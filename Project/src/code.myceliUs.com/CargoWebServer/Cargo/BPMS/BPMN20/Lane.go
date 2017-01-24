// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Lane struct {

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

	/** members of Lane **/
	M_name                string
	M_childLaneSet        *LaneSet
	m_partitionElementRef BaseElement
	/** If the ref is a string and not an object **/
	M_partitionElementRef string
	m_flowNodeRef         []FlowNode
	/** If the ref is a string and not an object **/
	M_flowNodeRef      []string
	M_partitionElement BaseElement

	/** Associations **/
	m_laneSetPtr *LaneSet
	/** If the ref is a string and not an object **/
	M_laneSetPtr string
	m_lanePtr    []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
}

/** Xml parser for Lane **/
type XsdLane struct {
	XMLName xml.Name `xml:"lane"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_flowNodeRef         []string         `xml:"flowNodeRef"`
	M_childLaneSet        *XsdChildLaneSet `xml:"childLaneSet,omitempty"`
	M_name                string           `xml:"name,attr"`
	M_partitionElementRef string           `xml:"partitionElementRef,attr"`
}

/** UUID **/
func (this *Lane) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Lane) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Lane) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Lane) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Lane) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Lane) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Lane) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Lane) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Lane) SetExtensionDefinitions(ref interface{}) {
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
func (this *Lane) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Lane) SetExtensionValues(ref interface{}) {
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
func (this *Lane) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Lane) SetDocumentation(ref interface{}) {
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
func (this *Lane) RemoveDocumentation(ref interface{}) {
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
func (this *Lane) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Lane) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** ChildLaneSet **/
func (this *Lane) GetChildLaneSet() *LaneSet {
	return this.M_childLaneSet
}

/** Init reference ChildLaneSet **/
func (this *Lane) SetChildLaneSet(ref interface{}) {
	this.NeedSave = true
	this.M_childLaneSet = ref.(*LaneSet)
}

/** Remove reference ChildLaneSet **/
func (this *Lane) RemoveChildLaneSet(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_childLaneSet.GetUUID() {
		this.M_childLaneSet = nil
	}
}

/** PartitionElementRef **/
func (this *Lane) GetPartitionElementRef() BaseElement {
	return this.m_partitionElementRef
}

/** Init reference PartitionElementRef **/
func (this *Lane) SetPartitionElementRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_partitionElementRef = ref.(string)
	} else {
		this.m_partitionElementRef = ref.(BaseElement)
		this.M_partitionElementRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference PartitionElementRef **/
func (this *Lane) RemovePartitionElementRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_partitionElementRef.(BaseElement).GetUUID() {
		this.m_partitionElementRef = nil
		this.M_partitionElementRef = ""
	}
}

/** FlowNodeRef **/
func (this *Lane) GetFlowNodeRef() []FlowNode {
	return this.m_flowNodeRef
}

/** Init reference FlowNodeRef **/
func (this *Lane) SetFlowNodeRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_flowNodeRef); i++ {
			if this.M_flowNodeRef[i] == refStr {
				return
			}
		}
		this.M_flowNodeRef = append(this.M_flowNodeRef, ref.(string))
	} else {
		this.RemoveFlowNodeRef(ref)
		this.m_flowNodeRef = append(this.m_flowNodeRef, ref.(FlowNode))
		this.M_flowNodeRef = append(this.M_flowNodeRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference FlowNodeRef **/
func (this *Lane) RemoveFlowNodeRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	flowNodeRef_ := make([]FlowNode, 0)
	flowNodeRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_flowNodeRef); i++ {
		if toDelete.GetUUID() != this.m_flowNodeRef[i].(BaseElement).GetUUID() {
			flowNodeRef_ = append(flowNodeRef_, this.m_flowNodeRef[i])
			flowNodeRefUuid = append(flowNodeRefUuid, this.M_flowNodeRef[i])
		}
	}
	this.m_flowNodeRef = flowNodeRef_
	this.M_flowNodeRef = flowNodeRefUuid
}

/** PartitionElement **/
func (this *Lane) GetPartitionElement() BaseElement {
	return this.M_partitionElement
}

/** Init reference PartitionElement **/
func (this *Lane) SetPartitionElement(ref interface{}) {
	this.NeedSave = true
	this.M_partitionElement = ref.(BaseElement)
}

/** Remove reference PartitionElement **/
func (this *Lane) RemovePartitionElement(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_partitionElement.(BaseElement).GetUUID() {
		this.M_partitionElement = nil
	}
}

/** LaneSet **/
func (this *Lane) GetLaneSetPtr() *LaneSet {
	return this.m_laneSetPtr
}

/** Init reference LaneSet **/
func (this *Lane) SetLaneSetPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_laneSetPtr = ref.(string)
	} else {
		this.m_laneSetPtr = ref.(*LaneSet)
		this.M_laneSetPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference LaneSet **/
func (this *Lane) RemoveLaneSetPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_laneSetPtr.GetUUID() {
		this.m_laneSetPtr = nil
		this.M_laneSetPtr = ""
	}
}

/** Lane **/
func (this *Lane) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Lane) SetLanePtr(ref interface{}) {
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
func (this *Lane) RemoveLanePtr(ref interface{}) {
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
func (this *Lane) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Lane) SetOutgoingPtr(ref interface{}) {
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
func (this *Lane) RemoveOutgoingPtr(ref interface{}) {
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
func (this *Lane) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Lane) SetIncomingPtr(ref interface{}) {
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
func (this *Lane) RemoveIncomingPtr(ref interface{}) {
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
