package BPMNDI

import(
"code.myceliUs.com/CargoWebServer/Cargo/Entities/DI"
"encoding/xml"
)

type BPMNPlane struct{

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

	/** members of Plane **/
	M_DiagramElement []DI.DiagramElement

	/** members of BPMNPlane **/
	m_bpmnElement interface{}
	/** If the ref is a string and not an object **/
	M_bpmnElement string


	/** Associations **/
	m_diagramPtr *BPMNDiagram
	/** If the ref is a string and not an object **/
	M_diagramPtr string
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

/** Xml parser for BPMNPlane **/
type XsdBPMNPlane struct {
	XMLName xml.Name	`xml:"BPMNPlane"`
	/** DiagramElement **/
	M_id	string	`xml:"id,attr"`


	/** Node **/


	/** Plane **/
	M_DiagramElement_0	[]*XsdBPMNShape	`xml:"BPMNShape,omitempty"`
	M_DiagramElement_1	[]*XsdBPMNEdge	`xml:"BPMNEdge,omitempty"`



	M_bpmnElement	string	`xml:"bpmnElement,attr"`

}
/** UUID **/
func (this *BPMNPlane) GetUUID() string{
	return this.UUID
}

/** OwningDiagram **/
func (this *BPMNPlane) GetOwningDiagram() DI.Diagram{
	return this.m_owningDiagram
}

/** Init reference OwningDiagram **/
func (this *BPMNPlane) SetOwningDiagram(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningDiagram = ref.(string)
	}else{
		this.m_owningDiagram = ref.(DI.Diagram)
	}
}

/** Remove reference OwningDiagram **/
func (this *BPMNPlane) RemoveOwningDiagram(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.Diagram)
	if toDelete.GetUUID() == this.m_owningDiagram.(DI.Diagram).GetUUID() {
		this.m_owningDiagram = nil
		this.M_owningDiagram = ""
	}
}

/** OwningElement **/
func (this *BPMNPlane) GetOwningElement() DI.DiagramElement{
	return this.m_owningElement
}

/** Init reference OwningElement **/
func (this *BPMNPlane) SetOwningElement(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningElement = ref.(string)
	}else{
		this.m_owningElement = ref.(DI.DiagramElement)
	}
}

/** Remove reference OwningElement **/
func (this *BPMNPlane) RemoveOwningElement(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_owningElement.(DI.DiagramElement).GetUUID() {
		this.m_owningElement = nil
		this.M_owningElement = ""
	}
}

/** ModelElement **/
func (this *BPMNPlane) GetModelElement() interface{}{
	return this.m_modelElement
}

/** Init reference ModelElement **/
func (this *BPMNPlane) SetModelElement(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_modelElement = ref.(string)
	}else{
		this.m_modelElement = ref.(interface{})
	}
}

/** Remove reference ModelElement **/

/** Style **/
func (this *BPMNPlane) GetStyle() DI.Style{
	return this.m_style
}

/** Init reference Style **/
func (this *BPMNPlane) SetStyle(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_style = ref.(string)
	}else{
		this.m_style = ref.(DI.Style)
	}
}

/** Remove reference Style **/
func (this *BPMNPlane) RemoveStyle(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.Style)
	if toDelete.GetUUID() == this.m_style.(DI.Style).GetUUID() {
		this.m_style = nil
		this.M_style = ""
	}
}

/** OwnedElement **/
func (this *BPMNPlane) GetOwnedElement() []DI.DiagramElement{
	return this.M_ownedElement
}

/** Init reference OwnedElement **/
func (this *BPMNPlane) SetOwnedElement(ref interface{}){
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
func (this *BPMNPlane) RemoveOwnedElement(ref interface{}){
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
func (this *BPMNPlane) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *BPMNPlane) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** DiagramElement **/
func (this *BPMNPlane) GetDiagramElement() []DI.DiagramElement{
	return this.M_DiagramElement
}

/** Init reference DiagramElement **/
func (this *BPMNPlane) SetDiagramElement(ref interface{}){
	this.NeedSave = true
	isExist := false
	var DiagramElements []DI.DiagramElement
	for i:=0; i<len(this.M_DiagramElement); i++ {
		if this.M_DiagramElement[i].GetUUID() != ref.(DI.DiagramElement).GetUUID() {
			DiagramElements = append(DiagramElements, this.M_DiagramElement[i])
		} else {
			isExist = true
			DiagramElements = append(DiagramElements, ref.(DI.DiagramElement))
		}
	}
	if !isExist {
		DiagramElements = append(DiagramElements, ref.(DI.DiagramElement))
	}
	this.M_DiagramElement = DiagramElements
}

/** Remove reference DiagramElement **/
func (this *BPMNPlane) RemoveDiagramElement(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	DiagramElement_ := make([]DI.DiagramElement, 0)
	for i := 0; i < len(this.M_DiagramElement); i++ {
		if toDelete.GetUUID() != this.M_DiagramElement[i].(DI.DiagramElement).GetUUID() {
			DiagramElement_ = append(DiagramElement_, this.M_DiagramElement[i])
		}
	}
	this.M_DiagramElement = DiagramElement_
}

/** BpmnElement **/
func (this *BPMNPlane) GetBpmnElement() interface{}{
	return this.m_bpmnElement
}

/** Init reference BpmnElement **/
func (this *BPMNPlane) SetBpmnElement(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_bpmnElement = ref.(string)
	}else{
		this.m_bpmnElement = ref.(interface{})
	}
}

/** Remove reference BpmnElement **/

/** Diagram **/
func (this *BPMNPlane) GetDiagramPtr() *BPMNDiagram{
	return this.m_diagramPtr
}

/** Init reference Diagram **/
func (this *BPMNPlane) SetDiagramPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_diagramPtr = ref.(string)
	}else{
		this.m_diagramPtr = ref.(*BPMNDiagram)
	}
}

/** Remove reference Diagram **/
func (this *BPMNPlane) RemoveDiagramPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.Diagram)
	if toDelete.GetUUID() == this.m_diagramPtr.GetUUID() {
		this.m_diagramPtr = nil
		this.M_diagramPtr = ""
	}
}

/** SourceEdge **/
func (this *BPMNPlane) GetSourceEdgePtr() []DI.Edge{
	return this.m_sourceEdgePtr
}

/** Init reference SourceEdge **/
func (this *BPMNPlane) SetSourceEdgePtr(ref interface{}){
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
func (this *BPMNPlane) RemoveSourceEdgePtr(ref interface{}){
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
func (this *BPMNPlane) GetTargetEdgePtr() []DI.Edge{
	return this.m_targetEdgePtr
}

/** Init reference TargetEdge **/
func (this *BPMNPlane) SetTargetEdgePtr(ref interface{}){
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
func (this *BPMNPlane) RemoveTargetEdgePtr(ref interface{}){
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
func (this *BPMNPlane) GetPlanePtr() DI.Plane{
	return this.m_planePtr
}

/** Init reference Plane **/
func (this *BPMNPlane) SetPlanePtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_planePtr = ref.(string)
	}else{
		this.m_planePtr = ref.(DI.Plane)
	}
}

/** Remove reference Plane **/
func (this *BPMNPlane) RemovePlanePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.m_planePtr.(DI.DiagramElement).GetUUID() {
		this.m_planePtr = nil
		this.M_planePtr = ""
	}
}
