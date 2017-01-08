package BPMNDI

import (
	"encoding/xml"

	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/DC"
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/DI"
)

type BPMNShape struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of DiagramElement **/
	m_owningDiagram DI.Diagram
	/** If the ref is a string and not an object **/
	M_owningDiagram string
	m_owningElement DI.DiagramElement
	/** If the ref is a string and not an object **/
	M_owningElement string
	m_modelElement  interface{}
	/** If the ref is a string and not an object **/
	M_modelElement string
	m_style        DI.Style
	/** If the ref is a string and not an object **/
	M_style        string
	M_ownedElement []DI.DiagramElement
	M_id           string

	/** members of Node **/
	/** No members **/

	/** members of Shape **/
	M_Bounds *DC.Bounds

	/** members of LabeledShape **/
	M_ownedLabel []DI.Label

	/** members of BPMNShape **/
	m_bpmnElement interface{}
	/** If the ref is a string and not an object **/
	M_bpmnElement               string
	M_isHorizontal              bool
	M_isExpanded                bool
	M_isMarkerVisible           bool
	M_BPMNLabel                 *BPMNLabel
	M_isMessageVisible          bool
	M_participantBandKind       ParticipantBandKind
	m_choreographyActivityShape *BPMNShape
	/** If the ref is a string and not an object **/
	M_choreographyActivityShape string

	/** Associations **/
	m_participantBandShapePtr *BPMNShape
	/** If the ref is a string and not an object **/
	M_participantBandShapePtr string
	m_sourceEdgePtr           []DI.Edge
	/** If the ref is a string and not an object **/
	M_sourceEdgePtr []string
	m_targetEdgePtr []DI.Edge
	/** If the ref is a string and not an object **/
	M_targetEdgePtr []string
	m_planePtr      DI.Plane
	/** If the ref is a string and not an object **/
	M_planePtr string
}

/** Xml parser for BPMNShape **/
type XsdBPMNShape struct {
	XMLName xml.Name `xml:"BPMNShape"`
	/** DiagramElement **/
	M_id string `xml:"id,attr"`

	/** Node **/

	/** Shape **/
	M_Bounds DC.XsdBounds `xml:"Bounds"`

	/** LabeledShape **/

	M_BPMNLabel                 *XsdBPMNLabel `xml:"BPMNLabel,omitempty"`
	M_bpmnElement               string        `xml:"bpmnElement,attr"`
	M_isHorizontal              bool          `xml:"isHorizontal,attr"`
	M_isExpanded                bool          `xml:"isExpanded,attr"`
	M_isMarkerVisible           bool          `xml:"isMarkerVisible,attr"`
	M_isMessageVisible          bool          `xml:"isMessageVisible,attr"`
	M_participantBandKind       string        `xml:"participantBandKind,attr"`
	M_choreographyActivityShape string        `xml:"choreographyActivityShape,attr"`
}

/** UUID **/
func (this *BPMNShape) GetUUID() string {
	return this.UUID
}

/** OwningDiagram **/
func (this *BPMNShape) GetOwningDiagram() DI.Diagram {
	return this.m_owningDiagram
}

/** Init reference OwningDiagram **/
func (this *BPMNShape) SetOwningDiagram(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningDiagram = ref.(string)
	} else {
		this.m_owningDiagram = ref.(DI.Diagram)
	}
}

/** Remove reference OwningDiagram **/
func (this *BPMNShape) RemoveOwningDiagram(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.Diagram)
	if toDelete.GetUUID() == this.m_owningDiagram.(DI.Diagram).GetUUID() {
		this.m_owningDiagram = nil
		this.M_owningDiagram = ""
	}
}

/** OwningElement **/
func (this *BPMNShape) GetOwningElement() DI.DiagramElement {
	return this.m_owningElement
}

/** Init reference OwningElement **/
func (this *BPMNShape) SetOwningElement(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningElement = ref.(string)
	} else {
		this.m_owningElement = ref.(DI.DiagramElement)
	}
}

/** Remove reference OwningElement **/
func (this *BPMNShape) RemoveOwningElement(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_owningElement.(DI.DiagramElement).GetUUID() {
		this.m_owningElement = nil
		this.M_owningElement = ""
	}
}

/** ModelElement **/
func (this *BPMNShape) GetModelElement() interface{} {
	return this.m_modelElement
}

/** Init reference ModelElement **/
func (this *BPMNShape) SetModelElement(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_modelElement = ref.(string)
	} else {
		this.m_modelElement = ref.(interface{})
	}
}

/** Remove reference ModelElement **/

/** Style **/
func (this *BPMNShape) GetStyle() DI.Style {
	return this.m_style
}

/** Init reference Style **/
func (this *BPMNShape) SetStyle(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_style = ref.(string)
	} else {
		this.m_style = ref.(DI.Style)
	}
}

/** Remove reference Style **/
func (this *BPMNShape) RemoveStyle(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.Style)
	if toDelete.GetUUID() == this.m_style.(DI.Style).GetUUID() {
		this.m_style = nil
		this.M_style = ""
	}
}

/** OwnedElement **/
func (this *BPMNShape) GetOwnedElement() []DI.DiagramElement {
	return this.M_ownedElement
}

/** Init reference OwnedElement **/
func (this *BPMNShape) SetOwnedElement(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var ownedElements []DI.DiagramElement
	for i := 0; i < len(this.M_ownedElement); i++ {
		if this.M_ownedElement[i].GetUUID() != ref.(DI.DiagramElement).GetUUID() {
			ownedElements = append(ownedElements, this.M_ownedElement[i])
		} else {
			isExist = true
			ownedElements = append(ownedElements, ref.(DI.DiagramElement))
		}
	}
	if !isExist {
		ownedElements = append(ownedElements, ref.(DI.DiagramElement))
	}
	this.M_ownedElement = ownedElements
}

/** Remove reference OwnedElement **/
func (this *BPMNShape) RemoveOwnedElement(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	ownedElement_ := make([]DI.DiagramElement, 0)
	for i := 0; i < len(this.M_ownedElement); i++ {
		if toDelete.GetUUID() != this.M_ownedElement[i].(DI.DiagramElement).GetUUID() {
			ownedElement_ = append(ownedElement_, this.M_ownedElement[i])
		}
	}
	this.M_ownedElement = ownedElement_
}

/** Id **/
func (this *BPMNShape) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *BPMNShape) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Bounds **/
func (this *BPMNShape) GetBounds() *DC.Bounds {
	return this.M_Bounds
}

/** Init reference Bounds **/
func (this *BPMNShape) SetBounds(ref interface{}) {
	this.NeedSave = true
	this.M_Bounds = ref.(*DC.Bounds)
}

/** Remove reference Bounds **/
func (this *BPMNShape) RemoveBounds(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(*DC.Bounds)
	if toDelete.GetUUID() == this.M_Bounds.GetUUID() {
		this.M_Bounds = nil
	}
}

/** OwnedLabel **/
func (this *BPMNShape) GetOwnedLabel() []DI.Label {
	return this.M_ownedLabel
}

/** Init reference OwnedLabel **/
func (this *BPMNShape) SetOwnedLabel(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var ownedLabels []DI.Label
	for i := 0; i < len(this.M_ownedLabel); i++ {
		if this.M_ownedLabel[i].GetUUID() != ref.(DI.DiagramElement).GetUUID() {
			ownedLabels = append(ownedLabels, this.M_ownedLabel[i])
		} else {
			isExist = true
			ownedLabels = append(ownedLabels, ref.(DI.Label))
		}
	}
	if !isExist {
		ownedLabels = append(ownedLabels, ref.(DI.Label))
	}
	this.M_ownedLabel = ownedLabels
}

/** Remove reference OwnedLabel **/
func (this *BPMNShape) RemoveOwnedLabel(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	ownedLabel_ := make([]DI.Label, 0)
	for i := 0; i < len(this.M_ownedLabel); i++ {
		if toDelete.GetUUID() != this.M_ownedLabel[i].(DI.DiagramElement).GetUUID() {
			ownedLabel_ = append(ownedLabel_, this.M_ownedLabel[i])
		}
	}
	this.M_ownedLabel = ownedLabel_
}

/** BpmnElement **/
func (this *BPMNShape) GetBpmnElement() interface{} {
	return this.m_bpmnElement
}

/** Init reference BpmnElement **/
func (this *BPMNShape) SetBpmnElement(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_bpmnElement = ref.(string)
	} else {
		this.m_bpmnElement = ref.(interface{})
	}
}

/** Remove reference BpmnElement **/

/** IsHorizontal **/
func (this *BPMNShape) IsHorizontal() bool {
	return this.M_isHorizontal
}

/** Init reference IsHorizontal **/
func (this *BPMNShape) SetIsHorizontal(ref interface{}) {
	this.NeedSave = true
	this.M_isHorizontal = ref.(bool)
}

/** Remove reference IsHorizontal **/

/** IsExpanded **/
func (this *BPMNShape) IsExpanded() bool {
	return this.M_isExpanded
}

/** Init reference IsExpanded **/
func (this *BPMNShape) SetIsExpanded(ref interface{}) {
	this.NeedSave = true
	this.M_isExpanded = ref.(bool)
}

/** Remove reference IsExpanded **/

/** IsMarkerVisible **/
func (this *BPMNShape) IsMarkerVisible() bool {
	return this.M_isMarkerVisible
}

/** Init reference IsMarkerVisible **/
func (this *BPMNShape) SetIsMarkerVisible(ref interface{}) {
	this.NeedSave = true
	this.M_isMarkerVisible = ref.(bool)
}

/** Remove reference IsMarkerVisible **/

/** BPMNLabel **/
func (this *BPMNShape) GetBPMNLabel() *BPMNLabel {
	return this.M_BPMNLabel
}

/** Init reference BPMNLabel **/
func (this *BPMNShape) SetBPMNLabel(ref interface{}) {
	this.NeedSave = true
	this.M_BPMNLabel = ref.(*BPMNLabel)
}

/** Remove reference BPMNLabel **/
func (this *BPMNShape) RemoveBPMNLabel(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.M_BPMNLabel.GetUUID() {
		this.M_BPMNLabel = nil
	}
}

/** IsMessageVisible **/
func (this *BPMNShape) IsMessageVisible() bool {
	return this.M_isMessageVisible
}

/** Init reference IsMessageVisible **/
func (this *BPMNShape) SetIsMessageVisible(ref interface{}) {
	this.NeedSave = true
	this.M_isMessageVisible = ref.(bool)
}

/** Remove reference IsMessageVisible **/

/** ParticipantBandKind **/
func (this *BPMNShape) GetParticipantBandKind() ParticipantBandKind {
	return this.M_participantBandKind
}

/** Init reference ParticipantBandKind **/
func (this *BPMNShape) SetParticipantBandKind(ref interface{}) {
	this.NeedSave = true
	this.M_participantBandKind = ref.(ParticipantBandKind)
}

/** Remove reference ParticipantBandKind **/

/** ChoreographyActivityShape **/
func (this *BPMNShape) GetChoreographyActivityShape() *BPMNShape {
	return this.m_choreographyActivityShape
}

/** Init reference ChoreographyActivityShape **/
func (this *BPMNShape) SetChoreographyActivityShape(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_choreographyActivityShape = ref.(string)
	} else {
		this.m_choreographyActivityShape = ref.(*BPMNShape)
	}
}

/** Remove reference ChoreographyActivityShape **/
func (this *BPMNShape) RemoveChoreographyActivityShape(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_choreographyActivityShape.GetUUID() {
		this.m_choreographyActivityShape = nil
		this.M_choreographyActivityShape = ""
	}
}

/** ParticipantBandShape **/
func (this *BPMNShape) GetParticipantBandShapePtr() *BPMNShape {
	return this.m_participantBandShapePtr
}

/** Init reference ParticipantBandShape **/
func (this *BPMNShape) SetParticipantBandShapePtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_participantBandShapePtr = ref.(string)
	} else {
		this.m_participantBandShapePtr = ref.(*BPMNShape)
	}
}

/** Remove reference ParticipantBandShape **/
func (this *BPMNShape) RemoveParticipantBandShapePtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_participantBandShapePtr.GetUUID() {
		this.m_participantBandShapePtr = nil
		this.M_participantBandShapePtr = ""
	}
}

/** SourceEdge **/
func (this *BPMNShape) GetSourceEdgePtr() []DI.Edge {
	return this.m_sourceEdgePtr
}

/** Init reference SourceEdge **/
func (this *BPMNShape) SetSourceEdgePtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_sourceEdgePtr); i++ {
			if this.M_sourceEdgePtr[i] == refStr {
				return
			}
		}
		this.M_sourceEdgePtr = append(this.M_sourceEdgePtr, ref.(string))
	} else {
		this.RemoveSourceEdgePtr(ref)
		this.m_sourceEdgePtr = append(this.m_sourceEdgePtr, ref.(DI.Edge))
	}
}

/** Remove reference SourceEdge **/
func (this *BPMNShape) RemoveSourceEdgePtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	sourceEdgePtr_ := make([]DI.Edge, 0)
	sourceEdgePtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_sourceEdgePtr); i++ {
		if toDelete.GetUUID() != this.m_sourceEdgePtr[i].(DI.DiagramElement).GetUUID() {
			sourceEdgePtr_ = append(sourceEdgePtr_, this.m_sourceEdgePtr[i])
			sourceEdgePtrUuid = append(sourceEdgePtrUuid, this.M_sourceEdgePtr[i])
		}
	}
	this.m_sourceEdgePtr = sourceEdgePtr_
	this.M_sourceEdgePtr = sourceEdgePtrUuid
}

/** TargetEdge **/
func (this *BPMNShape) GetTargetEdgePtr() []DI.Edge {
	return this.m_targetEdgePtr
}

/** Init reference TargetEdge **/
func (this *BPMNShape) SetTargetEdgePtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_targetEdgePtr); i++ {
			if this.M_targetEdgePtr[i] == refStr {
				return
			}
		}
		this.M_targetEdgePtr = append(this.M_targetEdgePtr, ref.(string))
	} else {
		this.RemoveTargetEdgePtr(ref)
		this.m_targetEdgePtr = append(this.m_targetEdgePtr, ref.(DI.Edge))
	}
}

/** Remove reference TargetEdge **/
func (this *BPMNShape) RemoveTargetEdgePtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	targetEdgePtr_ := make([]DI.Edge, 0)
	targetEdgePtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_targetEdgePtr); i++ {
		if toDelete.GetUUID() != this.m_targetEdgePtr[i].(DI.DiagramElement).GetUUID() {
			targetEdgePtr_ = append(targetEdgePtr_, this.m_targetEdgePtr[i])
			targetEdgePtrUuid = append(targetEdgePtrUuid, this.M_targetEdgePtr[i])
		}
	}
	this.m_targetEdgePtr = targetEdgePtr_
	this.M_targetEdgePtr = targetEdgePtrUuid
}

/** Plane **/
func (this *BPMNShape) GetPlanePtr() DI.Plane {
	return this.m_planePtr
}

/** Init reference Plane **/
func (this *BPMNShape) SetPlanePtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_planePtr = ref.(string)
	} else {
		this.m_planePtr = ref.(DI.Plane)
	}
}

/** Remove reference Plane **/
func (this *BPMNShape) RemovePlanePtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_planePtr.(DI.DiagramElement).GetUUID() {
		this.m_planePtr = nil
		this.M_planePtr = ""
	}
}
