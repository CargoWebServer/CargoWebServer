// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type DataInputAssociation struct {

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

	/** members of DataAssociation **/
	M_transformation *FormalExpression
	M_assignment     []*Assignment
	m_targetRef      ItemAwareElement
	/** If the ref is a string and not an object **/
	M_targetRef string
	m_sourceRef []ItemAwareElement
	/** If the ref is a string and not an object **/
	M_sourceRef []string

	/** members of DataInputAssociation **/
	/** No members **/

	/** Associations **/
	m_throwEventPtr ThrowEvent
	/** If the ref is a string and not an object **/
	M_throwEventPtr string
	m_activityPtr   Activity
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
	M_incomingPtr []string
}

/** Xml parser for DataInputAssociation **/
type XsdDataInputAssociation struct {
	XMLName xml.Name `xml:"dataInputAssociation"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	/** DataAssociation **/
	M_sourceRef      []string           `xml:"sourceRef"`
	M_targetRef      string             `xml:"targetRef"`
	M_transformation *XsdTransformation `xml:"transformation,omitempty"`
	M_assignment     []*XsdAssignment   `xml:"assignment,omitempty"`
}

/** UUID **/
func (this *DataInputAssociation) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *DataInputAssociation) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *DataInputAssociation) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *DataInputAssociation) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *DataInputAssociation) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *DataInputAssociation) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *DataInputAssociation) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *DataInputAssociation) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *DataInputAssociation) SetExtensionDefinitions(ref interface{}) {
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
func (this *DataInputAssociation) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *DataInputAssociation) SetExtensionValues(ref interface{}) {
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
func (this *DataInputAssociation) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *DataInputAssociation) SetDocumentation(ref interface{}) {
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
func (this *DataInputAssociation) RemoveDocumentation(ref interface{}) {
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

/** Transformation **/
func (this *DataInputAssociation) GetTransformation() *FormalExpression {
	return this.M_transformation
}

/** Init reference Transformation **/
func (this *DataInputAssociation) SetTransformation(ref interface{}) {
	this.NeedSave = true
	this.M_transformation = ref.(*FormalExpression)
}

/** Remove reference Transformation **/
func (this *DataInputAssociation) RemoveTransformation(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_transformation.GetUUID() {
		this.M_transformation = nil
	}
}

/** Assignment **/
func (this *DataInputAssociation) GetAssignment() []*Assignment {
	return this.M_assignment
}

/** Init reference Assignment **/
func (this *DataInputAssociation) SetAssignment(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var assignments []*Assignment
	for i := 0; i < len(this.M_assignment); i++ {
		if this.M_assignment[i].GetUUID() != ref.(BaseElement).GetUUID() {
			assignments = append(assignments, this.M_assignment[i])
		} else {
			isExist = true
			assignments = append(assignments, ref.(*Assignment))
		}
	}
	if !isExist {
		assignments = append(assignments, ref.(*Assignment))
	}
	this.M_assignment = assignments
}

/** Remove reference Assignment **/
func (this *DataInputAssociation) RemoveAssignment(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	assignment_ := make([]*Assignment, 0)
	for i := 0; i < len(this.M_assignment); i++ {
		if toDelete.GetUUID() != this.M_assignment[i].GetUUID() {
			assignment_ = append(assignment_, this.M_assignment[i])
		}
	}
	this.M_assignment = assignment_
}

/** TargetRef **/
func (this *DataInputAssociation) GetTargetRef() ItemAwareElement {
	return this.m_targetRef
}

/** Init reference TargetRef **/
func (this *DataInputAssociation) SetTargetRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_targetRef = ref.(string)
	} else {
		this.m_targetRef = ref.(ItemAwareElement)
		this.M_targetRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference TargetRef **/
func (this *DataInputAssociation) RemoveTargetRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_targetRef.(BaseElement).GetUUID() {
		this.m_targetRef = nil
		this.M_targetRef = ""
	}
}

/** SourceRef **/
func (this *DataInputAssociation) GetSourceRef() []ItemAwareElement {
	return this.m_sourceRef
}

/** Init reference SourceRef **/
func (this *DataInputAssociation) SetSourceRef(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_sourceRef); i++ {
			if this.M_sourceRef[i] == refStr {
				return
			}
		}
		this.M_sourceRef = append(this.M_sourceRef, ref.(string))
	} else {
		this.RemoveSourceRef(ref)
		this.m_sourceRef = append(this.m_sourceRef, ref.(ItemAwareElement))
		this.M_sourceRef = append(this.M_sourceRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference SourceRef **/
func (this *DataInputAssociation) RemoveSourceRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	sourceRef_ := make([]ItemAwareElement, 0)
	sourceRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_sourceRef); i++ {
		if toDelete.GetUUID() != this.m_sourceRef[i].(BaseElement).GetUUID() {
			sourceRef_ = append(sourceRef_, this.m_sourceRef[i])
			sourceRefUuid = append(sourceRefUuid, this.M_sourceRef[i])
		}
	}
	this.m_sourceRef = sourceRef_
	this.M_sourceRef = sourceRefUuid
}

/** ThrowEvent **/
func (this *DataInputAssociation) GetThrowEventPtr() ThrowEvent {
	return this.m_throwEventPtr
}

/** Init reference ThrowEvent **/
func (this *DataInputAssociation) SetThrowEventPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_throwEventPtr = ref.(string)
	} else {
		this.m_throwEventPtr = ref.(ThrowEvent)
		this.M_throwEventPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ThrowEvent **/
func (this *DataInputAssociation) RemoveThrowEventPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_throwEventPtr.(BaseElement).GetUUID() {
		this.m_throwEventPtr = nil
		this.M_throwEventPtr = ""
	}
}

/** Activity **/
func (this *DataInputAssociation) GetActivityPtr() Activity {
	return this.m_activityPtr
}

/** Init reference Activity **/
func (this *DataInputAssociation) SetActivityPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_activityPtr = ref.(string)
	} else {
		this.m_activityPtr = ref.(Activity)
		this.M_activityPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Activity **/
func (this *DataInputAssociation) RemoveActivityPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_activityPtr.(BaseElement).GetUUID() {
		this.m_activityPtr = nil
		this.M_activityPtr = ""
	}
}

/** Lane **/
func (this *DataInputAssociation) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *DataInputAssociation) SetLanePtr(ref interface{}) {
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
func (this *DataInputAssociation) RemoveLanePtr(ref interface{}) {
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
func (this *DataInputAssociation) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *DataInputAssociation) SetOutgoingPtr(ref interface{}) {
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
func (this *DataInputAssociation) RemoveOutgoingPtr(ref interface{}) {
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
func (this *DataInputAssociation) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *DataInputAssociation) SetIncomingPtr(ref interface{}) {
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
func (this *DataInputAssociation) RemoveIncomingPtr(ref interface{}) {
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
