//+build BPMN
package BPMS_Runtime

import (
	"encoding/xml"
)

type ItemAwareElementInstance struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of ItemAwareElementInstance **/
	M_id            string
	M_bpmnElementId string
	M_data          []uint8

	/** Associations **/
	m_parentPtr Instance
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for ItemAwareElementInstance **/
type XsdItemAwareElementInstance struct {
	XMLName         xml.Name `xml:"itemAwareElementInstance"`
	M_id            string   `xml:"id,attr"`
	M_bpmnElementId string   `xml:"bpmnElementId,attr"`
	M_data          []uint8  `xml:"data,attr"`
}

/** UUID **/
func (this *ItemAwareElementInstance) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *ItemAwareElementInstance) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *ItemAwareElementInstance) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** BpmnElementId **/
func (this *ItemAwareElementInstance) GetBpmnElementId() string {
	return this.M_bpmnElementId
}

/** Init reference BpmnElementId **/
func (this *ItemAwareElementInstance) SetBpmnElementId(ref interface{}) {
	this.NeedSave = true
	this.M_bpmnElementId = ref.(string)
}

/** Remove reference BpmnElementId **/

/** Data **/
func (this *ItemAwareElementInstance) GetData() []uint8 {
	return this.M_data
}

/** Init reference Data **/
func (this *ItemAwareElementInstance) SetData(ref interface{}) {
	this.NeedSave = true
	this.M_data = ref.([]uint8)
}

/** Remove reference Data **/

/** Parent **/
func (this *ItemAwareElementInstance) GetParentPtr() Instance {
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *ItemAwareElementInstance) SetParentPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	} else {
		this.m_parentPtr = ref.(Instance)
		this.M_parentPtr = ref.(Instance).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *ItemAwareElementInstance) RemoveParentPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(Instance)
	if toDelete.GetUUID() == this.m_parentPtr.(Instance).GetUUID() {
		this.m_parentPtr = nil
		this.M_parentPtr = ""
	}
}
