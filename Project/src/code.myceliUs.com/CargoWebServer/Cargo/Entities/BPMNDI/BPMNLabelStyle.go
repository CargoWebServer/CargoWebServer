package BPMNDI

import(
"code.myceliUs.com/CargoWebServer/Cargo/Entities/DI"
"code.myceliUs.com/CargoWebServer/Cargo/Entities/DC"
"encoding/xml"
)

type BPMNLabelStyle struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Style **/
	M_id string

	/** members of BPMNLabelStyle **/
	M_Font *DC.Font


	/** Associations **/
	m_diagramPtr *BPMNDiagram
	/** If the ref is a string and not an object **/
	M_diagramPtr string
	m_labelPtr []*BPMNLabel
	/** If the ref is a string and not an object **/
	M_labelPtr []string
	m_diagramElementPtr []DI.DiagramElement
	/** If the ref is a string and not an object **/
	M_diagramElementPtr []string
	m_owningDiagramPtr DI.Diagram
	/** If the ref is a string and not an object **/
	M_owningDiagramPtr string
}

/** Xml parser for BPMNLabelStyle **/
type XsdBPMNLabelStyle struct {
	XMLName xml.Name	`xml:"BPMNLabelStyle"`
	/** Style **/
	M_id	string	`xml:"id,attr"`


	M_Font	DC.XsdFont	`xml:"Font"`

}
/** UUID **/
func (this *BPMNLabelStyle) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *BPMNLabelStyle) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *BPMNLabelStyle) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Font **/
func (this *BPMNLabelStyle) GetFont() *DC.Font{
	return this.M_Font
}

/** Init reference Font **/
func (this *BPMNLabelStyle) SetFont(ref interface{}){
	this.NeedSave = true
	this.M_Font = ref.(*DC.Font)
}

/** Remove reference Font **/

/** Diagram **/
func (this *BPMNLabelStyle) GetDiagramPtr() *BPMNDiagram{
	return this.m_diagramPtr
}

/** Init reference Diagram **/
func (this *BPMNLabelStyle) SetDiagramPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_diagramPtr = ref.(string)
	}else{
		this.m_diagramPtr = ref.(*BPMNDiagram)
	}
}

/** Remove reference Diagram **/
func (this *BPMNLabelStyle) RemoveDiagramPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.Diagram)
	if toDelete.GetUUID() == this.m_diagramPtr.GetUUID() {
		this.m_diagramPtr = nil
		this.M_diagramPtr = ""
	}
}

/** Label **/
func (this *BPMNLabelStyle) GetLabelPtr() []*BPMNLabel{
	return this.m_labelPtr
}

/** Init reference Label **/
func (this *BPMNLabelStyle) SetLabelPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_labelPtr); i++ {
			if this.M_labelPtr[i] == refStr {
				return
			}
		}
		this.M_labelPtr = append(this.M_labelPtr, ref.(string))
	}else{
		this.RemoveLabelPtr(ref)
		this.m_labelPtr = append(this.m_labelPtr, ref.(*BPMNLabel))
	}
}

/** Remove reference Label **/
func (this *BPMNLabelStyle) RemoveLabelPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	labelPtr_ := make([]*BPMNLabel, 0)
	labelPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_labelPtr); i++ {
		if toDelete.GetUUID() != this.m_labelPtr[i].GetUUID() {
			labelPtr_ = append(labelPtr_, this.m_labelPtr[i])
			labelPtrUuid = append(labelPtrUuid, this.M_labelPtr[i])
		}
	}
	this.m_labelPtr = labelPtr_
	this.M_labelPtr = labelPtrUuid
}

/** DiagramElement **/
func (this *BPMNLabelStyle) GetDiagramElementPtr() []DI.DiagramElement{
	return this.m_diagramElementPtr
}

/** Init reference DiagramElement **/
func (this *BPMNLabelStyle) SetDiagramElementPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_diagramElementPtr); i++ {
			if this.M_diagramElementPtr[i] == refStr {
				return
			}
		}
		this.M_diagramElementPtr = append(this.M_diagramElementPtr, ref.(string))
	}else{
		this.RemoveDiagramElementPtr(ref)
		this.m_diagramElementPtr = append(this.m_diagramElementPtr, ref.(DI.DiagramElement))
	}
}

/** Remove reference DiagramElement **/
func (this *BPMNLabelStyle) RemoveDiagramElementPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	diagramElementPtr_ := make([]DI.DiagramElement, 0)
	diagramElementPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_diagramElementPtr); i++ {
		if toDelete.GetUUID() != this.m_diagramElementPtr[i].(DI.DiagramElement).GetUUID() {
			diagramElementPtr_ = append(diagramElementPtr_, this.m_diagramElementPtr[i])
			diagramElementPtrUuid = append(diagramElementPtrUuid, this.M_diagramElementPtr[i])
		}
	}
	this.m_diagramElementPtr = diagramElementPtr_
	this.M_diagramElementPtr = diagramElementPtrUuid
}

/** OwningDiagram **/
func (this *BPMNLabelStyle) GetOwningDiagramPtr() DI.Diagram{
	return this.m_owningDiagramPtr
}

/** Init reference OwningDiagram **/
func (this *BPMNLabelStyle) SetOwningDiagramPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_owningDiagramPtr = ref.(string)
	}else{
		this.m_owningDiagramPtr = ref.(DI.Diagram)
	}
}

/** Remove reference OwningDiagram **/
func (this *BPMNLabelStyle) RemoveOwningDiagramPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(DI.Diagram)
	if toDelete.GetUUID() == this.m_owningDiagramPtr.(DI.Diagram).GetUUID() {
		this.m_owningDiagramPtr = nil
		this.M_owningDiagramPtr = ""
	}
}
