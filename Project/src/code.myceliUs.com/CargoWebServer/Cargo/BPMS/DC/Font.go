package DC

import(
"encoding/xml"
)

type Font struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Font **/
	M_name string
	M_size float64
	M_isBold bool
	M_isItalic bool
	M_isUnderline bool
	M_isStrikeThrough bool

}

/** Xml parser for Font **/
type XsdFont struct {
	XMLName xml.Name	`xml:"Font"`
	M_name	string	`xml:"name,attr"`
	M_size	float64	`xml:"size,attr"`
	M_isBold	bool	`xml:"isBold,attr"`
	M_isItalic	bool	`xml:"isItalic,attr"`
	M_isUnderline	bool	`xml:"isUnderline,attr"`
	M_isStrikeThrough	bool	`xml:"isStrikeThrough,attr"`

}
/** UUID **/
func (this *Font) GetUUID() string{
	return this.UUID
}

/** Name **/
func (this *Font) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Font) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Size **/
func (this *Font) GetSize() float64{
	return this.M_size
}

/** Init reference Size **/
func (this *Font) SetSize(ref interface{}){
	this.NeedSave = true
	this.M_size = ref.(float64)
}

/** Remove reference Size **/

/** IsBold **/
func (this *Font) IsBold() bool{
	return this.M_isBold
}

/** Init reference IsBold **/
func (this *Font) SetIsBold(ref interface{}){
	this.NeedSave = true
	this.M_isBold = ref.(bool)
}

/** Remove reference IsBold **/

/** IsItalic **/
func (this *Font) IsItalic() bool{
	return this.M_isItalic
}

/** Init reference IsItalic **/
func (this *Font) SetIsItalic(ref interface{}){
	this.NeedSave = true
	this.M_isItalic = ref.(bool)
}

/** Remove reference IsItalic **/

/** IsUnderline **/
func (this *Font) IsUnderline() bool{
	return this.M_isUnderline
}

/** Init reference IsUnderline **/
func (this *Font) SetIsUnderline(ref interface{}){
	this.NeedSave = true
	this.M_isUnderline = ref.(bool)
}

/** Remove reference IsUnderline **/

/** IsStrikeThrough **/
func (this *Font) IsStrikeThrough() bool{
	return this.M_isStrikeThrough
}

/** Init reference IsStrikeThrough **/
func (this *Font) SetIsStrikeThrough(ref interface{}){
	this.NeedSave = true
	this.M_isStrikeThrough = ref.(bool)
}

/** Remove reference IsStrikeThrough **/
