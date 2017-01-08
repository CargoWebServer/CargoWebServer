package BPMNDI

import (
	"encoding/xml"

	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/DI"
)

type BPMNDiagram struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of Diagram **/
	M_rootElement   DI.DiagramElement
	M_name          string
	M_id            string
	M_documentation string
	M_resolution    float64
	M_ownedStyle    []DI.Style

	/** members of BPMNDiagram **/
	M_BPMNPlane      *BPMNPlane
	M_BPMNLabelStyle []*BPMNLabelStyle
}

/** Xml parser for BPMNDiagram **/
type XsdBPMNDiagram struct {
	XMLName xml.Name `xml:"BPMNDiagram"`
	/** Diagram **/
	M_name          string  `xml:"name,attr"`
	M_documentation string  `xml:"documentation,attr"`
	M_resolution    float64 `xml:"resolution,attr"`
	M_id            string  `xml:"id,attr"`

	M_BPMNPlane      *XsdBPMNPlane        `xml:"BPMNPlane,omitempty"`
	M_BPMNLabelStyle []*XsdBPMNLabelStyle `xml:"BPMNLabelStyle,omitempty"`
}

/** UUID **/
func (this *BPMNDiagram) GetUUID() string {
	return this.UUID
}

/** RootElement **/
func (this *BPMNDiagram) GetRootElement() DI.DiagramElement {
	return this.M_rootElement
}

/** Init reference RootElement **/
func (this *BPMNDiagram) SetRootElement(ref interface{}) {
	this.NeedSave = true
	this.M_rootElement = ref.(DI.DiagramElement)
}

/** Remove reference RootElement **/
func (this *BPMNDiagram) RemoveRootElement(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.M_rootElement.(DI.DiagramElement).GetUUID() {
		this.M_rootElement = nil
	}
}

/** Name **/
func (this *BPMNDiagram) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *BPMNDiagram) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Id **/
func (this *BPMNDiagram) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *BPMNDiagram) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Documentation **/
func (this *BPMNDiagram) GetDocumentation() string {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *BPMNDiagram) SetDocumentation(ref interface{}) {
	this.NeedSave = true
	this.M_documentation = ref.(string)
}

/** Remove reference Documentation **/

/** Resolution **/
func (this *BPMNDiagram) GetResolution() float64 {
	return this.M_resolution
}

/** Init reference Resolution **/
func (this *BPMNDiagram) SetResolution(ref interface{}) {
	this.NeedSave = true
	this.M_resolution = ref.(float64)
}

/** Remove reference Resolution **/

/** OwnedStyle **/
func (this *BPMNDiagram) GetOwnedStyle() []DI.Style {
	return this.M_ownedStyle
}

/** Init reference OwnedStyle **/
func (this *BPMNDiagram) SetOwnedStyle(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var ownedStyles []DI.Style
	for i := 0; i < len(this.M_ownedStyle); i++ {
		if this.M_ownedStyle[i].GetUUID() != ref.(DI.Style).GetUUID() {
			ownedStyles = append(ownedStyles, this.M_ownedStyle[i])
		} else {
			isExist = true
			ownedStyles = append(ownedStyles, ref.(DI.Style))
		}
	}
	if !isExist {
		ownedStyles = append(ownedStyles, ref.(DI.Style))
	}
	this.M_ownedStyle = ownedStyles
}

/** Remove reference OwnedStyle **/
func (this *BPMNDiagram) RemoveOwnedStyle(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.Style)
	ownedStyle_ := make([]DI.Style, 0)
	for i := 0; i < len(this.M_ownedStyle); i++ {
		if toDelete.GetUUID() != this.M_ownedStyle[i].(DI.Style).GetUUID() {
			ownedStyle_ = append(ownedStyle_, this.M_ownedStyle[i])
		}
	}
	this.M_ownedStyle = ownedStyle_
}

/** BPMNPlane **/
func (this *BPMNDiagram) GetBPMNPlane() *BPMNPlane {
	return this.M_BPMNPlane
}

/** Init reference BPMNPlane **/
func (this *BPMNDiagram) SetBPMNPlane(ref interface{}) {
	this.NeedSave = true
	this.M_BPMNPlane = ref.(*BPMNPlane)
}

/** Remove reference BPMNPlane **/
func (this *BPMNDiagram) RemoveBPMNPlane(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.DiagramElement)
	if toDelete.GetUUID() == this.M_BPMNPlane.GetUUID() {
		this.M_BPMNPlane = nil
	}
}

/** BPMNLabelStyle **/
func (this *BPMNDiagram) GetBPMNLabelStyle() []*BPMNLabelStyle {
	return this.M_BPMNLabelStyle
}

/** Init reference BPMNLabelStyle **/
func (this *BPMNDiagram) SetBPMNLabelStyle(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var BPMNLabelStyles []*BPMNLabelStyle
	for i := 0; i < len(this.M_BPMNLabelStyle); i++ {
		if this.M_BPMNLabelStyle[i].GetUUID() != ref.(DI.Style).GetUUID() {
			BPMNLabelStyles = append(BPMNLabelStyles, this.M_BPMNLabelStyle[i])
		} else {
			isExist = true
			BPMNLabelStyles = append(BPMNLabelStyles, ref.(*BPMNLabelStyle))
		}
	}
	if !isExist {
		BPMNLabelStyles = append(BPMNLabelStyles, ref.(*BPMNLabelStyle))
	}
	this.M_BPMNLabelStyle = BPMNLabelStyles
}

/** Remove reference BPMNLabelStyle **/
func (this *BPMNDiagram) RemoveBPMNLabelStyle(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(DI.Style)
	BPMNLabelStyle_ := make([]*BPMNLabelStyle, 0)
	for i := 0; i < len(this.M_BPMNLabelStyle); i++ {
		if toDelete.GetUUID() != this.M_BPMNLabelStyle[i].GetUUID() {
			BPMNLabelStyle_ = append(BPMNLabelStyle_, this.M_BPMNLabelStyle[i])
		}
	}
	this.M_BPMNLabelStyle = BPMNLabelStyle_
}
