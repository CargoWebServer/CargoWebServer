package DC

import(
"encoding/xml"
)

type Point struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Point **/
	M_id string
	M_x float64
	M_y float64

}

/** Xml parser for Point **/
type XsdPoint struct {
	XMLName xml.Name	`xml:"waypoint"`
	M_x	float64	`xml:"x,attr"`
	M_y	float64	`xml:"y,attr"`

}
/** UUID **/
func (this *Point) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Point) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Point) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** X **/
func (this *Point) GetX() float64{
	return this.M_x
}

/** Init reference X **/
func (this *Point) SetX(ref interface{}){
	this.NeedSave = true
	this.M_x = ref.(float64)
}

/** Remove reference X **/

/** Y **/
func (this *Point) GetY() float64{
	return this.M_y
}

/** Init reference Y **/
func (this *Point) SetY(ref interface{}){
	this.NeedSave = true
	this.M_y = ref.(float64)
}

/** Remove reference Y **/
