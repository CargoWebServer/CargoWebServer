// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type ExtensionElements struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of ExtensionElements **/
	M_value string

	/** Associations **/
	m_baseElementPtr []BaseElement
	/** If the ref is a string and not an object **/
	M_baseElementPtr []string
}

/** Xml parser for ExtensionElements **/
type XsdExtensionElements struct {
	XMLName xml.Name `xml:"extensionElements"`
	M_value string   `xml:",innerxml"`
}

/** UUID **/
func (this *ExtensionElements) GetUUID() string {
	return this.UUID
}

/** Value **/
func (this *ExtensionElements) GetValue() string {
	return this.M_value
}

/** Init reference Value **/
func (this *ExtensionElements) SetValue(ref interface{}) {
	this.NeedSave = true
	this.M_value = ref.(string)
}

/** Remove reference Value **/

/** BaseElement **/
func (this *ExtensionElements) GetBaseElementPtr() []BaseElement {
	return this.m_baseElementPtr
}

/** Init reference BaseElement **/
func (this *ExtensionElements) SetBaseElementPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_baseElementPtr); i++ {
			if this.M_baseElementPtr[i] == refStr {
				return
			}
		}
		this.M_baseElementPtr = append(this.M_baseElementPtr, ref.(string))
	} else {
		this.RemoveBaseElementPtr(ref)
		this.m_baseElementPtr = append(this.m_baseElementPtr, ref.(BaseElement))
		this.M_baseElementPtr = append(this.M_baseElementPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference BaseElement **/
func (this *ExtensionElements) RemoveBaseElementPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	baseElementPtr_ := make([]BaseElement, 0)
	baseElementPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_baseElementPtr); i++ {
		if toDelete.GetUUID() != this.m_baseElementPtr[i].(BaseElement).GetUUID() {
			baseElementPtr_ = append(baseElementPtr_, this.m_baseElementPtr[i])
			baseElementPtrUuid = append(baseElementPtrUuid, this.M_baseElementPtr[i])
		}
	}
	this.m_baseElementPtr = baseElementPtr_
	this.M_baseElementPtr = baseElementPtrUuid
}
