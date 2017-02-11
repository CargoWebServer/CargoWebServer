package BPMS

import(
"encoding/xml"
)

type EventData struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of EventData **/
	M_id string


	/** Associations **/
	m_triggerPtr *Trigger
	/** If the ref is a string and not an object **/
	M_triggerPtr string
}

/** Xml parser for EventData **/
type XsdEventData struct {
	XMLName xml.Name	`xml:"eventData"`
	M_id	string	`xml:"id,attr"`

}
/** UUID **/
func (this *EventData) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *EventData) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *EventData) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Trigger **/
func (this *EventData) GetTriggerPtr() *Trigger{
	return this.m_triggerPtr
}

/** Init reference Trigger **/
func (this *EventData) SetTriggerPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_triggerPtr = ref.(string)
	}else{
		this.m_triggerPtr = ref.(*Trigger)
		this.M_triggerPtr = ref.(*Trigger).GetUUID()
	}
}

/** Remove reference Trigger **/
func (this *EventData) RemoveTriggerPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Trigger)
	if toDelete.GetUUID() == this.m_triggerPtr.GetUUID() {
		this.m_triggerPtr = nil
		this.M_triggerPtr = ""
	}
}
