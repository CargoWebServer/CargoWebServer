package BPMN20

import(
"encoding/xml"
)

type ParticipantMultiplicity struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of BaseElement **/
	M_id string
	m_other interface{}
	/** If the ref is a string and not an object **/
	M_other string
	M_extensionElements *ExtensionElements
	M_extensionDefinitions []*ExtensionDefinition
	M_extensionValues []*ExtensionAttributeValue
	M_documentation []*Documentation

	/** members of ParticipantMultiplicity **/
	M_minimum int
	M_maximum int


	/** Associations **/
	m_participantPtr *Participant
	/** If the ref is a string and not an object **/
	M_participantPtr string
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
}

/** Xml parser for ParticipantMultiplicity **/
type XsdParticipantMultiplicity struct {
	XMLName xml.Name	`xml:"participantMultiplicity"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	M_minimum	int	`xml:"minimum,attr"`
	M_maximum	int	`xml:"maximum,attr"`

}
/** UUID **/
func (this *ParticipantMultiplicity) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *ParticipantMultiplicity) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ParticipantMultiplicity) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *ParticipantMultiplicity) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *ParticipantMultiplicity) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *ParticipantMultiplicity) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *ParticipantMultiplicity) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *ParticipantMultiplicity) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *ParticipantMultiplicity) SetExtensionDefinitions(ref interface{}){
	this.NeedSave = true
	isExist := false
	var extensionDefinitionss []*ExtensionDefinition
	for i:=0; i<len(this.M_extensionDefinitions); i++ {
		if this.M_extensionDefinitions[i].GetName() != ref.(*ExtensionDefinition).GetName() {
			extensionDefinitionss = append(extensionDefinitionss, this.M_extensionDefinitions[i])
		} else {
			isExist = true
			extensionDefinitionss = append(extensionDefinitionss, ref.(*ExtensionDefinition))
		}
	}
	if !isExist {
		extensionDefinitionss = append(extensionDefinitionss, ref.(*ExtensionDefinition))
	}
	this.M_extensionDefinitions = extensionDefinitionss
}

/** Remove reference ExtensionDefinitions **/

/** ExtensionValues **/
func (this *ParticipantMultiplicity) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *ParticipantMultiplicity) SetExtensionValues(ref interface{}){
	this.NeedSave = true
	isExist := false
	var extensionValuess []*ExtensionAttributeValue
	for i:=0; i<len(this.M_extensionValues); i++ {
		if this.M_extensionValues[i].GetUUID() != ref.(*ExtensionAttributeValue).GetUUID() {
			extensionValuess = append(extensionValuess, this.M_extensionValues[i])
		} else {
			isExist = true
			extensionValuess = append(extensionValuess, ref.(*ExtensionAttributeValue))
		}
	}
	if !isExist {
		extensionValuess = append(extensionValuess, ref.(*ExtensionAttributeValue))
	}
	this.M_extensionValues = extensionValuess
}

/** Remove reference ExtensionValues **/

/** Documentation **/
func (this *ParticipantMultiplicity) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *ParticipantMultiplicity) SetDocumentation(ref interface{}){
	this.NeedSave = true
	isExist := false
	var documentations []*Documentation
	for i:=0; i<len(this.M_documentation); i++ {
		if this.M_documentation[i].GetUUID() != ref.(BaseElement).GetUUID() {
			documentations = append(documentations, this.M_documentation[i])
		} else {
			isExist = true
			documentations = append(documentations, ref.(*Documentation))
		}
	}
	if !isExist {
		documentations = append(documentations, ref.(*Documentation))
	}
	this.M_documentation = documentations
}

/** Remove reference Documentation **/
func (this *ParticipantMultiplicity) RemoveDocumentation(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	documentation_ := make([]*Documentation, 0)
	for i := 0; i < len(this.M_documentation); i++ {
		if toDelete.GetUUID() != this.M_documentation[i].GetUUID() {
			documentation_ = append(documentation_, this.M_documentation[i])
		}
	}
	this.M_documentation = documentation_
}

/** Minimum **/
func (this *ParticipantMultiplicity) GetMinimum() int{
	return this.M_minimum
}

/** Init reference Minimum **/
func (this *ParticipantMultiplicity) SetMinimum(ref interface{}){
	this.NeedSave = true
	this.M_minimum = ref.(int)
}

/** Remove reference Minimum **/

/** Maximum **/
func (this *ParticipantMultiplicity) GetMaximum() int{
	return this.M_maximum
}

/** Init reference Maximum **/
func (this *ParticipantMultiplicity) SetMaximum(ref interface{}){
	this.NeedSave = true
	this.M_maximum = ref.(int)
}

/** Remove reference Maximum **/

/** Participant **/
func (this *ParticipantMultiplicity) GetParticipantPtr() *Participant{
	return this.m_participantPtr
}

/** Init reference Participant **/
func (this *ParticipantMultiplicity) SetParticipantPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_participantPtr = ref.(string)
	}else{
		this.m_participantPtr = ref.(*Participant)
		this.M_participantPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Participant **/
func (this *ParticipantMultiplicity) RemoveParticipantPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_participantPtr.GetUUID() {
		this.m_participantPtr = nil
		this.M_participantPtr = ""
	}
}

/** Lane **/
func (this *ParticipantMultiplicity) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *ParticipantMultiplicity) SetLanePtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_lanePtr); i++ {
			if this.M_lanePtr[i] == refStr {
				return
			}
		}
		this.M_lanePtr = append(this.M_lanePtr, ref.(string))
	}else{
		this.RemoveLanePtr(ref)
		this.m_lanePtr = append(this.m_lanePtr, ref.(*Lane))
		this.M_lanePtr = append(this.M_lanePtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Lane **/
func (this *ParticipantMultiplicity) RemoveLanePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	lanePtr_ := make([]*Lane, 0)
	lanePtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_lanePtr); i++ {
		if toDelete.GetUUID() != this.m_lanePtr[i].GetUUID() {
			lanePtr_ = append(lanePtr_, this.m_lanePtr[i])
			lanePtrUuid = append(lanePtrUuid, this.M_lanePtr[i])
		}
	}
	this.m_lanePtr = lanePtr_
	this.M_lanePtr = lanePtrUuid
}

/** Outgoing **/
func (this *ParticipantMultiplicity) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *ParticipantMultiplicity) SetOutgoingPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_outgoingPtr); i++ {
			if this.M_outgoingPtr[i] == refStr {
				return
			}
		}
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(string))
	}else{
		this.RemoveOutgoingPtr(ref)
		this.m_outgoingPtr = append(this.m_outgoingPtr, ref.(*Association))
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Outgoing **/
func (this *ParticipantMultiplicity) RemoveOutgoingPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	outgoingPtr_ := make([]*Association, 0)
	outgoingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_outgoingPtr); i++ {
		if toDelete.GetUUID() != this.m_outgoingPtr[i].GetUUID() {
			outgoingPtr_ = append(outgoingPtr_, this.m_outgoingPtr[i])
			outgoingPtrUuid = append(outgoingPtrUuid, this.M_outgoingPtr[i])
		}
	}
	this.m_outgoingPtr = outgoingPtr_
	this.M_outgoingPtr = outgoingPtrUuid
}

/** Incoming **/
func (this *ParticipantMultiplicity) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *ParticipantMultiplicity) SetIncomingPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_incomingPtr); i++ {
			if this.M_incomingPtr[i] == refStr {
				return
			}
		}
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(string))
	}else{
		this.RemoveIncomingPtr(ref)
		this.m_incomingPtr = append(this.m_incomingPtr, ref.(*Association))
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Incoming **/
func (this *ParticipantMultiplicity) RemoveIncomingPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	incomingPtr_ := make([]*Association, 0)
	incomingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_incomingPtr); i++ {
		if toDelete.GetUUID() != this.m_incomingPtr[i].GetUUID() {
			incomingPtr_ = append(incomingPtr_, this.m_incomingPtr[i])
			incomingPtrUuid = append(incomingPtrUuid, this.M_incomingPtr[i])
		}
	}
	this.m_incomingPtr = incomingPtr_
	this.M_incomingPtr = incomingPtrUuid
}
