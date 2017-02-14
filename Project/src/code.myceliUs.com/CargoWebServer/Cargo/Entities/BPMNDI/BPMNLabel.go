// +build BPMNDI

package BPMNDI

import(
"code.myceliUs.com/CargoWebServer/Cargo/Entities/DC"
"code.myceliUs.com/CargoWebServer/Cargo/Entities/DI"
"encoding/xml"
)

type BPMNLabel struct{

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

	/** members of Node **/
	/** No members **/

	/** members of Label **/
	M_Bounds *DC.Bounds

	/** members of BPMNLabel **/
	m_labelStyle *BPMNLabelStyle
	/** If the ref is a string and not an object **/
	M_labelStyle string


	/** Associations **/
	m_shapePtr *BPMNShape
	/** If the ref is a string and not an object **/
	M_shapePtr string
	m_edgePtr *BPMNEdge
	/** If the ref is a string and not an object **/
	M_edgePtr string
	m_sourceEdgePtr []DI.Edge
	/** If the ref is a string and not an object **/
	M_sourceEdgePtr []string
	m_targetEdgePtr []DI.Edge
	/** If the ref is a string and not an object **/
	M_targetEdgePtr []string
	m_planePtr DI.Plane
	/** If the ref is a string and not an object **/
	M_planePtr string
	m_owningEdgePtr DI.LabeledEdge
	/** If the ref is a string and not an object **/
	M_owningEdgePtr string
	m_owningShapePtr DI.LabeledShape
	/** If the ref is a string and not an object **/
	M_owningShapePtr string
}

/** Xml parser for BPMNLabel **/
type XsdBPMNLabel struct {
	XMLName xml.Name	`xml:"BPMNLabel"`
	/** DiagramElement **/
	M_id	string	`xml:"id,attr"`


	/** Node **/


	/** Label **/
	M_Bounds	*DC.XsdBounds	`xml:"Bounds,omitempty"`


	M_labelStyle	string	`xml:"labelStyle,attr"`

}
/** UUID **/
func (this *BPMNLabel) GetUUID() string{
	return this.UUID
}

/** OwningDiagram **/
func (this *BPMNLabel) GetOwningDiagram() DI.Diagram{
	return this.m_owningDiagram
}

/** Init reference OwningDiagram **/
func (this *BPMNLabel) SetOwningDiagram(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningDiagram = ref.(string)
	}else{
		this.m_owningDiagram = ref.(DI.Diagram)
	}
}

/** Remove reference OwningDiagram **/
func (this *BPMNLabel) RemoveOwningDiagram(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.Diagram)
	if toDelete.GetUUID() == this.m_owningDiagram.(DI.Diagram).GetUUID() {
		this.m_owningDiagram = nil
		this.M_owningDiagram = ""
	}
}

/** OwningElement **/
func (this *BPMNLabel) GetOwningElement() DI.DiagramElement{
	return this.m_owningElement
}

/** Init reference OwningElement **/
func (this *BPMNLabel) SetOwningElement(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningElement = ref.(string)
	}else{
		this.m_owningElement = ref.(DI.DiagramElement)
	}
}

/** Remove reference OwningElement **/
func (this *BPMNLabel) RemoveOwningElement(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_owningElement.(DI.DiagramElement).GetUUID() {
		this.m_owningElement = nil
		this.M_owningElement = ""
	}
}

/** ModelElement **/
func (this *BPMNLabel) GetModelElement() interface{}{
	return this.m_modelElement
}

/** Init reference ModelElement **/
func (this *BPMNLabel) SetModelElement(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_modelElement = ref.(string)
	}else{
		this.m_modelElement = ref.(interface{})
	}
}

/** Remove reference ModelElement **/

/** Style **/
func (this *BPMNLabel) GetStyle() DI.Style{
	return this.m_style
}

/** Init reference Style **/
func (this *BPMNLabel) SetStyle(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_style = ref.(string)
	}else{
		this.m_style = ref.(DI.Style)
	}
}

/** Remove reference Style **/
func (this *BPMNLabel) RemoveStyle(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.Style)
	if toDelete.GetUUID() == this.m_style.(DI.Style).GetUUID() {
		this.m_style = nil
		this.M_style = ""
	}
}

/** OwnedElement **/
func (this *BPMNLabel) GetOwnedElement() []DI.DiagramElement{
	return this.M_ownedElement
}

/** Init reference OwnedElement **/
func (this *BPMNLabel) SetOwnedElement(ref interface{}){
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
func (this *BPMNLabel) RemoveOwnedElement(ref interface{}){
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
func (this *BPMNLabel) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *BPMNLabel) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Bounds **/
func (this *BPMNLabel) GetBounds() *DC.Bounds{
	return this.M_Bounds
}

/** Init reference Bounds **/
func (this *BPMNLabel) SetBounds(ref interface{}){
	this.NeedSave = true
	this.M_Bounds = ref.(*DC.Bounds)
}

/** Remove reference Bounds **/
func (this *BPMNLabel) RemoveBounds(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*DC.Bounds)
	if toDelete.GetUUID() == this.M_Bounds.GetUUID() {
		this.M_Bounds = nil
	}
}

/** LabelStyle **/
func (this *BPMNLabel) GetLabelStyle() *BPMNLabelStyle{
	return this.m_labelStyle
}

/** Init reference LabelStyle **/
func (this *BPMNLabel) SetLabelStyle(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_labelStyle = ref.(string)
	}else{
		this.m_labelStyle = ref.(*BPMNLabelStyle)
	}
}

/** Remove reference LabelStyle **/
func (this *BPMNLabel) RemoveLabelStyle(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.Style)
	if toDelete.GetUUID() == this.m_labelStyle.GetUUID() {
		this.m_labelStyle = nil
		this.M_labelStyle = ""
	}
}

/** Shape **/
func (this *BPMNLabel) GetShapePtr() *BPMNShape{
	return this.m_shapePtr
}

/** Init reference Shape **/
func (this *BPMNLabel) SetShapePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_shapePtr = ref.(string)
	}else{
		this.m_shapePtr = ref.(*BPMNShape)
	}
}

/** Remove reference Shape **/
func (this *BPMNLabel) RemoveShapePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_shapePtr.GetUUID() {
		this.m_shapePtr = nil
		this.M_shapePtr = ""
	}
}

/** Edge **/
func (this *BPMNLabel) GetEdgePtr() *BPMNEdge{
	return this.m_edgePtr
}

/** Init reference Edge **/
func (this *BPMNLabel) SetEdgePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_edgePtr = ref.(string)
	}else{
		this.m_edgePtr = ref.(*BPMNEdge)
	}
}

/** Remove reference Edge **/
func (this *BPMNLabel) RemoveEdgePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_edgePtr.GetUUID() {
		this.m_edgePtr = nil
		this.M_edgePtr = ""
	}
}

/** SourceEdge **/
func (this *BPMNLabel) GetSourceEdgePtr() []DI.Edge{
	return this.m_sourceEdgePtr
}

/** Init reference SourceEdge **/
func (this *BPMNLabel) SetSourceEdgePtr(ref interface{}){
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
func (this *BPMNLabel) RemoveSourceEdgePtr(ref interface{}){
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
func (this *BPMNLabel) GetTargetEdgePtr() []DI.Edge{
	return this.m_targetEdgePtr
}

/** Init reference TargetEdge **/
func (this *BPMNLabel) SetTargetEdgePtr(ref interface{}){
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
func (this *BPMNLabel) RemoveTargetEdgePtr(ref interface{}){
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
func (this *BPMNLabel) GetPlanePtr() DI.Plane{
	return this.m_planePtr
}

/** Init reference Plane **/
func (this *BPMNLabel) SetPlanePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_planePtr = ref.(string)
	}else{
		this.m_planePtr = ref.(DI.Plane)
	}
}

/** Remove reference Plane **/
func (this *BPMNLabel) RemovePlanePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_planePtr.(DI.DiagramElement).GetUUID() {
		this.m_planePtr = nil
		this.M_planePtr = ""
	}
}

/** OwningEdge **/
func (this *BPMNLabel) GetOwningEdgePtr() DI.LabeledEdge{
	return this.m_owningEdgePtr
}

/** Init reference OwningEdge **/
func (this *BPMNLabel) SetOwningEdgePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningEdgePtr = ref.(string)
	}else{
		this.m_owningEdgePtr = ref.(DI.LabeledEdge)
	}
}

/** Remove reference OwningEdge **/
func (this *BPMNLabel) RemoveOwningEdgePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_owningEdgePtr.(DI.DiagramElement).GetUUID() {
		this.m_owningEdgePtr = nil
		this.M_owningEdgePtr = ""
	}
}

/** OwningShape **/
func (this *BPMNLabel) GetOwningShapePtr() DI.LabeledShape{
	return this.m_owningShapePtr
}

/** Init reference OwningShape **/
func (this *BPMNLabel) SetOwningShapePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningShapePtr = ref.(string)
	}else{
		this.m_owningShapePtr = ref.(DI.LabeledShape)
	}
}

/** Remove reference OwningShape **/
func (this *BPMNLabel) RemoveOwningShapePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_owningShapePtr.(DI.DiagramElement).GetUUID() {
		this.m_owningShapePtr = nil
		this.M_owningShapePtr = ""
	}
}
