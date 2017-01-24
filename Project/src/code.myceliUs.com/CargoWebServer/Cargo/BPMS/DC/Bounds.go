//+build BPMN
package DC

import (
	"encoding/xml"
)

type Bounds struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of Bounds **/
	M_id     string
	M_x      float64
	M_y      float64
	M_width  float64
	M_height float64
}

/** Xml parser for Bounds **/
type XsdBounds struct {
	XMLName  xml.Name `xml:"Bounds"`
	M_x      float64  `xml:"x,attr"`
	M_y      float64  `xml:"y,attr"`
	M_width  float64  `xml:"width,attr"`
	M_height float64  `xml:"height,attr"`
}

/** UUID **/
func (this *Bounds) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Bounds) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Bounds) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** X **/
func (this *Bounds) GetX() float64 {
	return this.M_x
}

/** Init reference X **/
func (this *Bounds) SetX(ref interface{}) {
	this.NeedSave = true
	this.M_x = ref.(float64)
}

/** Remove reference X **/

/** Y **/
func (this *Bounds) GetY() float64 {
	return this.M_y
}

/** Init reference Y **/
func (this *Bounds) SetY(ref interface{}) {
	this.NeedSave = true
	this.M_y = ref.(float64)
}

/** Remove reference Y **/

/** Width **/
func (this *Bounds) GetWidth() float64 {
	return this.M_width
}

/** Init reference Width **/
func (this *Bounds) SetWidth(ref interface{}) {
	this.NeedSave = true
	this.M_width = ref.(float64)
}

/** Remove reference Width **/

/** Height **/
func (this *Bounds) GetHeight() float64 {
	return this.M_height
}

/** Init reference Height **/
func (this *Bounds) SetHeight(ref interface{}) {
	this.NeedSave = true
	this.M_height = ref.(float64)
}

/** Remove reference Height **/
