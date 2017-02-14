// +build BPMN20

package BPMN20

import(
"encoding/xml"
)

type Text struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Text **/
	M_text string


	/** Associations **/
	m_textAnnotationPtr *TextAnnotation
	/** If the ref is a string and not an object **/
	M_textAnnotationPtr string
}

/** Xml parser for Text **/
type XsdText struct {
	XMLName xml.Name	`xml:"text"`
	M_text	string	`xml:",innerxml"`

}
/** UUID **/
func (this *Text) GetUUID() string{
	return this.UUID
}

/** Text **/
func (this *Text) GetText() string{
	return this.M_text
}

/** Init reference Text **/
func (this *Text) SetText(ref interface{}){
	this.NeedSave = true
	this.M_text = ref.(string)
}

/** Remove reference Text **/

/** TextAnnotation **/
func (this *Text) GetTextAnnotationPtr() *TextAnnotation{
	return this.m_textAnnotationPtr
}

/** Init reference TextAnnotation **/
func (this *Text) SetTextAnnotationPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_textAnnotationPtr = ref.(string)
	}else{
		this.m_textAnnotationPtr = ref.(*TextAnnotation)
		this.M_textAnnotationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference TextAnnotation **/
func (this *Text) RemoveTextAnnotationPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_textAnnotationPtr.GetUUID() {
		this.m_textAnnotationPtr = nil
		this.M_textAnnotationPtr = ""
	}
}
