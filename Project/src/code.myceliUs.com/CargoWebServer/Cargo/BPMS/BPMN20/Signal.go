package BPMN20

import(
"encoding/xml"
)

type Signal struct{

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

	/** members of RootElement **/
	/** No members **/

	/** members of Signal **/
	m_structureRef *ItemDefinition
	/** If the ref is a string and not an object **/
	M_structureRef string
	M_name string


	/** Associations **/
	m_signalEventDefinitionPtr []*SignalEventDefinition
	/** If the ref is a string and not an object **/
	M_signalEventDefinitionPtr []string
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
	m_definitionsPtr *Definitions
	/** If the ref is a string and not an object **/
	M_definitionsPtr string
}

/** Xml parser for Signal **/
type XsdSignal struct {
	XMLName xml.Name	`xml:"signal"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** RootElement **/


	M_name	string	`xml:"name,attr"`
	M_structureRef	string	`xml:"structureRef,attr"`

}
/** UUID **/
func (this *Signal) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Signal) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Signal) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Signal) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *Signal) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Signal) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Signal) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Signal) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Signal) SetExtensionDefinitions(ref interface{}){
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
func (this *Signal) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Signal) SetExtensionValues(ref interface{}){
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
func (this *Signal) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Signal) SetDocumentation(ref interface{}){
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
func (this *Signal) RemoveDocumentation(ref interface{}){
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

/** StructureRef **/
func (this *Signal) GetStructureRef() *ItemDefinition{
	return this.m_structureRef
}

/** Init reference StructureRef **/
func (this *Signal) SetStructureRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_structureRef = ref.(string)
	}else{
		this.m_structureRef = ref.(*ItemDefinition)
		this.M_structureRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference StructureRef **/
func (this *Signal) RemoveStructureRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_structureRef.GetUUID() {
		this.m_structureRef = nil
		this.M_structureRef = ""
	}
}

/** Name **/
func (this *Signal) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Signal) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** SignalEventDefinition **/
func (this *Signal) GetSignalEventDefinitionPtr() []*SignalEventDefinition{
	return this.m_signalEventDefinitionPtr
}

/** Init reference SignalEventDefinition **/
func (this *Signal) SetSignalEventDefinitionPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_signalEventDefinitionPtr); i++ {
			if this.M_signalEventDefinitionPtr[i] == refStr {
				return
			}
		}
		this.M_signalEventDefinitionPtr = append(this.M_signalEventDefinitionPtr, ref.(string))
	}else{
		this.RemoveSignalEventDefinitionPtr(ref)
		this.m_signalEventDefinitionPtr = append(this.m_signalEventDefinitionPtr, ref.(*SignalEventDefinition))
		this.M_signalEventDefinitionPtr = append(this.M_signalEventDefinitionPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference SignalEventDefinition **/
func (this *Signal) RemoveSignalEventDefinitionPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	signalEventDefinitionPtr_ := make([]*SignalEventDefinition, 0)
	signalEventDefinitionPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_signalEventDefinitionPtr); i++ {
		if toDelete.GetUUID() != this.m_signalEventDefinitionPtr[i].GetUUID() {
			signalEventDefinitionPtr_ = append(signalEventDefinitionPtr_, this.m_signalEventDefinitionPtr[i])
			signalEventDefinitionPtrUuid = append(signalEventDefinitionPtrUuid, this.M_signalEventDefinitionPtr[i])
		}
	}
	this.m_signalEventDefinitionPtr = signalEventDefinitionPtr_
	this.M_signalEventDefinitionPtr = signalEventDefinitionPtrUuid
}

/** Lane **/
func (this *Signal) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Signal) SetLanePtr(ref interface{}){
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
func (this *Signal) RemoveLanePtr(ref interface{}){
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
func (this *Signal) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Signal) SetOutgoingPtr(ref interface{}){
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
func (this *Signal) RemoveOutgoingPtr(ref interface{}){
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
func (this *Signal) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Signal) SetIncomingPtr(ref interface{}){
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
func (this *Signal) RemoveIncomingPtr(ref interface{}){
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

/** Definitions **/
func (this *Signal) GetDefinitionsPtr() *Definitions{
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *Signal) SetDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	}else{
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *Signal) RemoveDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}
