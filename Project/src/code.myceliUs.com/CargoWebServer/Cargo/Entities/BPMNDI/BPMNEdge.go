// +build BPMNDI

package BPMNDI

import(
"code.myceliUs.com/CargoWebServer/Cargo/Entities/DC"
"code.myceliUs.com/CargoWebServer/Cargo/Entities/DI"
"encoding/xml"
)

type BPMNEdge struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of DiagramElement **/
	m_owningDiagram DI.Diagram
	/** If the ref is a string and not an object **/
	M_owningDiagram string
	m_owningElement DI.DiagramElement
	/** If the ref is a string and not an object **/
	M_owningElement string
	m_modelElement interface{}
	/** If the ref is a string and not an object **/
	M_modelElement string
	m_style DI.Style
	/** If the ref is a string and not an object **/
	M_style string
	M_ownedElement []DI.DiagramElement
	M_id string

	/** members of Edge **/
	M_source DI.DiagramElement
	M_target DI.DiagramElement
	M_waypoint []*DC.Point

	/** members of LabeledEdge **/
	M_ownedLabel []DI.Label

	/** members of BPMNEdge **/
	M_BPMNLabel *BPMNLabel
	m_bpmnElement interface{}
	/** If the ref is a string and not an object **/
	M_bpmnElement string
	M_sourceElement DI.DiagramElement
	m_targetElement DI.DiagramElement
	/** If the ref is a string and not an object **/
	M_targetElement string
	M_messageVisibleKind MessageVisibleKind


	/** Associations **/
	m_sourceEdgePtr []DI.Edge
	/** If the ref is a string and not an object **/
	M_sourceEdgePtr []string
	m_targetEdgePtr []DI.Edge
	/** If the ref is a string and not an object **/
	M_targetEdgePtr []string
	m_planePtr DI.Plane
	/** If the ref is a string and not an object **/
	M_planePtr string
}

/** Xml parser for BPMNEdge **/
type XsdBPMNEdge struct {
	XMLName xml.Name	`xml:"BPMNEdge"`
	/** DiagramElement **/
	M_id	string	`xml:"id,attr"`


	/** Edge **/
	M_waypoint	[]*DC.XsdPoint	`xml:"waypoint,omitempty"`


	/** LabeledEdge **/


	M_BPMNLabel	*XsdBPMNLabel	`xml:"BPMNLabel,omitempty"`
	M_bpmnElement	string	`xml:"bpmnElement,attr"`
	M_sourceElement	string	`xml:"sourceElement,attr"`
	M_targetElement	string	`xml:"targetElement,attr"`
	M_messageVisibleKind	string	`xml:"messageVisibleKind,attr"`

}
/** UUID **/
func (this *BPMNEdge) GetUUID() string{
	return this.UUID
}

/** OwningDiagram **/
func (this *BPMNEdge) GetOwningDiagram() DI.Diagram{
	return this.m_owningDiagram
}

/** Init reference OwningDiagram **/
func (this *BPMNEdge) SetOwningDiagram(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningDiagram = ref.(string)
	}else{
		this.m_owningDiagram = ref.(DI.Diagram)
	}
}

/** Remove reference OwningDiagram **/
func (this *BPMNEdge) RemoveOwningDiagram(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.Diagram)
	if toDelete.GetUUID() == this.m_owningDiagram.(DI.Diagram).GetUUID() {
		this.m_owningDiagram = nil
		this.M_owningDiagram = ""
	}
}

/** OwningElement **/
func (this *BPMNEdge) GetOwningElement() DI.DiagramElement{
	return this.m_owningElement
}

/** Init reference OwningElement **/
func (this *BPMNEdge) SetOwningElement(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningElement = ref.(string)
	}else{
		this.m_owningElement = ref.(DI.DiagramElement)
	}
}

/** Remove reference OwningElement **/
func (this *BPMNEdge) RemoveOwningElement(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_owningElement.(DI.DiagramElement).GetUUID() {
		this.m_owningElement = nil
		this.M_owningElement = ""
	}
}

/** ModelElement **/
func (this *BPMNEdge) GetModelElement() interface{}{
	return this.m_modelElement
}

/** Init reference ModelElement **/
func (this *BPMNEdge) SetModelElement(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_modelElement = ref.(string)
	}else{
		this.m_modelElement = ref.(interface{})
	}
}

/** Remove reference ModelElement **/

/** Style **/
func (this *BPMNEdge) GetStyle() DI.Style{
	return this.m_style
}

/** Init reference Style **/
func (this *BPMNEdge) SetStyle(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_style = ref.(string)
	}else{
		this.m_style = ref.(DI.Style)
	}
}

/** Remove reference Style **/
func (this *BPMNEdge) RemoveStyle(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.Style)
	if toDelete.GetUUID() == this.m_style.(DI.Style).GetUUID() {
		this.m_style = nil
		this.M_style = ""
	}
}

/** OwnedElement **/
func (this *BPMNEdge) GetOwnedElement() []DI.DiagramElement{
	return this.M_ownedElement
}

/** Init reference OwnedElement **/
func (this *BPMNEdge) SetOwnedElement(ref interface{}){
	this.NeedSave = true
	isExist := false
	var ownedElements []DI.DiagramElement
	for i:=0; i<len(this.M_ownedElement); i++ {
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
func (this *BPMNEdge) RemoveOwnedElement(ref interface{}){
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
func (this *BPMNEdge) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *BPMNEdge) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Source **/
func (this *BPMNEdge) GetSource() DI.DiagramElement{
	return this.M_source
}

/** Init reference Source **/
func (this *BPMNEdge) SetSource(ref interface{}){
	this.NeedSave = true
	this.M_source = ref.(DI.DiagramElement)
}

/** Remove reference Source **/
func (this *BPMNEdge) RemoveSource(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.M_source.(DI.DiagramElement).GetUUID() {
		this.M_source = nil
	}
}

/** Target **/
func (this *BPMNEdge) GetTarget() DI.DiagramElement{
	return this.M_target
}

/** Init reference Target **/
func (this *BPMNEdge) SetTarget(ref interface{}){
	this.NeedSave = true
	this.M_target = ref.(DI.DiagramElement)
}

/** Remove reference Target **/
func (this *BPMNEdge) RemoveTarget(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.M_target.(DI.DiagramElement).GetUUID() {
		this.M_target = nil
	}
}

/** Waypoint **/
func (this *BPMNEdge) GetWaypoint() []*DC.Point{
	return this.M_waypoint
}

/** Init reference Waypoint **/
func (this *BPMNEdge) SetWaypoint(ref interface{}){
	this.NeedSave = true
	isExist := false
	var waypoints []*DC.Point
	for i:=0; i<len(this.M_waypoint); i++ {
		if this.M_waypoint[i].GetUUID() != ref.(*DC.Point).GetUUID() {
			waypoints = append(waypoints, this.M_waypoint[i])
		} else {
			isExist = true
			waypoints = append(waypoints, ref.(*DC.Point))
		}
	}
	if !isExist {
		waypoints = append(waypoints, ref.(*DC.Point))
	}
	this.M_waypoint = waypoints
}

/** Remove reference Waypoint **/
func (this *BPMNEdge) RemoveWaypoint(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*DC.Point)
	waypoint_ := make([]*DC.Point, 0)
	for i := 0; i < len(this.M_waypoint); i++ {
		if toDelete.GetUUID() != this.M_waypoint[i].GetUUID() {
			waypoint_ = append(waypoint_, this.M_waypoint[i])
		}
	}
	this.M_waypoint = waypoint_
}

/** OwnedLabel **/
func (this *BPMNEdge) GetOwnedLabel() []DI.Label{
	return this.M_ownedLabel
}

/** Init reference OwnedLabel **/
func (this *BPMNEdge) SetOwnedLabel(ref interface{}){
	this.NeedSave = true
	isExist := false
	var ownedLabels []DI.Label
	for i:=0; i<len(this.M_ownedLabel); i++ {
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
func (this *BPMNEdge) RemoveOwnedLabel(ref interface{}){
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

/** BPMNLabel **/
func (this *BPMNEdge) GetBPMNLabel() *BPMNLabel{
	return this.M_BPMNLabel
}

/** Init reference BPMNLabel **/
func (this *BPMNEdge) SetBPMNLabel(ref interface{}){
	this.NeedSave = true
	this.M_BPMNLabel = ref.(*BPMNLabel)
}

/** Remove reference BPMNLabel **/
func (this *BPMNEdge) RemoveBPMNLabel(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.M_BPMNLabel.GetUUID() {
		this.M_BPMNLabel = nil
	}
}

/** BpmnElement **/
func (this *BPMNEdge) GetBpmnElement() interface{}{
	return this.m_bpmnElement
}

/** Init reference BpmnElement **/
func (this *BPMNEdge) SetBpmnElement(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_bpmnElement = ref.(string)
	}else{
		this.m_bpmnElement = ref.(interface{})
	}
}

/** Remove reference BpmnElement **/

/** SourceElement **/
func (this *BPMNEdge) GetSourceElement() DI.DiagramElement{
	return this.M_sourceElement
}

/** Init reference SourceElement **/
func (this *BPMNEdge) SetSourceElement(ref interface{}){
	this.NeedSave = true
	this.M_sourceElement = ref.(DI.DiagramElement)
}

/** Remove reference SourceElement **/
func (this *BPMNEdge) RemoveSourceElement(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.M_sourceElement.(DI.DiagramElement).GetUUID() {
		this.M_sourceElement = nil
	}
}

/** TargetElement **/
func (this *BPMNEdge) GetTargetElement() DI.DiagramElement{
	return this.m_targetElement
}

/** Init reference TargetElement **/
func (this *BPMNEdge) SetTargetElement(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_targetElement = ref.(string)
	}else{
		this.m_targetElement = ref.(DI.DiagramElement)
	}
}

/** Remove reference TargetElement **/
func (this *BPMNEdge) RemoveTargetElement(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_targetElement.(DI.DiagramElement).GetUUID() {
		this.m_targetElement = nil
		this.M_targetElement = ""
	}
}

/** MessageVisibleKind **/
func (this *BPMNEdge) GetMessageVisibleKind() MessageVisibleKind{
	return this.M_messageVisibleKind
}

/** Init reference MessageVisibleKind **/
func (this *BPMNEdge) SetMessageVisibleKind(ref interface{}){
	this.NeedSave = true
	this.M_messageVisibleKind = ref.(MessageVisibleKind)
}

/** Remove reference MessageVisibleKind **/

/** SourceEdge **/
func (this *BPMNEdge) GetSourceEdgePtr() []DI.Edge{
	return this.m_sourceEdgePtr
}

/** Init reference SourceEdge **/
func (this *BPMNEdge) SetSourceEdgePtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_sourceEdgePtr); i++ {
			if this.M_sourceEdgePtr[i] == refStr {
				return
			}
		}
		this.M_sourceEdgePtr = append(this.M_sourceEdgePtr, ref.(string))
	}else{
		this.RemoveSourceEdgePtr(ref)
		this.m_sourceEdgePtr = append(this.m_sourceEdgePtr, ref.(DI.Edge))
	}
}

/** Remove reference SourceEdge **/
func (this *BPMNEdge) RemoveSourceEdgePtr(ref interface{}){
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
func (this *BPMNEdge) GetTargetEdgePtr() []DI.Edge{
	return this.m_targetEdgePtr
}

/** Init reference TargetEdge **/
func (this *BPMNEdge) SetTargetEdgePtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_targetEdgePtr); i++ {
			if this.M_targetEdgePtr[i] == refStr {
				return
			}
		}
		this.M_targetEdgePtr = append(this.M_targetEdgePtr, ref.(string))
	}else{
		this.RemoveTargetEdgePtr(ref)
		this.m_targetEdgePtr = append(this.m_targetEdgePtr, ref.(DI.Edge))
	}
}

/** Remove reference TargetEdge **/
func (this *BPMNEdge) RemoveTargetEdgePtr(ref interface{}){
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
func (this *BPMNEdge) GetPlanePtr() DI.Plane{
	return this.m_planePtr
}

/** Init reference Plane **/
func (this *BPMNEdge) SetPlanePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_planePtr = ref.(string)
	}else{
		this.m_planePtr = ref.(DI.Plane)
	}
}

/** Remove reference Plane **/
func (this *BPMNEdge) RemovePlanePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_planePtr.(DI.DiagramElement).GetUUID() {
		this.m_planePtr = nil
		this.M_planePtr = ""
	}
}
