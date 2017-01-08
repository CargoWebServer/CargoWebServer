package BPMN20

import(
"encoding/xml"
)

type TimerEventDefinition struct{

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

	/** members of EventDefinition **/
	/** No members **/

	/** members of TimerEventDefinition **/
	M_timeDate Expression
	M_timeCycle Expression
	M_timeDuration Expression


	/** Associations **/
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
	m_throwEventPtr []ThrowEvent
	/** If the ref is a string and not an object **/
	M_throwEventPtr []string
	m_catchEventPtr []CatchEvent
	/** If the ref is a string and not an object **/
	M_catchEventPtr []string
	m_multiInstanceLoopCharacteristicsPtr []*MultiInstanceLoopCharacteristics
	/** If the ref is a string and not an object **/
	M_multiInstanceLoopCharacteristicsPtr []string
}

/** Xml parser for TimerEventDefinition **/
type XsdTimerEventDefinition struct {
	XMLName xml.Name	`xml:"timerEventDefinition"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** RootElement **/


	/** EventDefinition **/



}
/** UUID **/
func (this *TimerEventDefinition) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *TimerEventDefinition) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *TimerEventDefinition) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *TimerEventDefinition) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *TimerEventDefinition) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *TimerEventDefinition) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *TimerEventDefinition) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *TimerEventDefinition) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *TimerEventDefinition) SetExtensionDefinitions(ref interface{}){
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
func (this *TimerEventDefinition) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *TimerEventDefinition) SetExtensionValues(ref interface{}){
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
func (this *TimerEventDefinition) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *TimerEventDefinition) SetDocumentation(ref interface{}){
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
func (this *TimerEventDefinition) RemoveDocumentation(ref interface{}){
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

/** TimeDate **/
func (this *TimerEventDefinition) GetTimeDate() Expression{
	return this.M_timeDate
}

/** Init reference TimeDate **/
func (this *TimerEventDefinition) SetTimeDate(ref interface{}){
	this.NeedSave = true
	this.M_timeDate = ref.(Expression)
}

/** Remove reference TimeDate **/
func (this *TimerEventDefinition) RemoveTimeDate(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_timeDate.(BaseElement).GetUUID() {
		this.M_timeDate = nil
	}
}

/** TimeCycle **/
func (this *TimerEventDefinition) GetTimeCycle() Expression{
	return this.M_timeCycle
}

/** Init reference TimeCycle **/
func (this *TimerEventDefinition) SetTimeCycle(ref interface{}){
	this.NeedSave = true
	this.M_timeCycle = ref.(Expression)
}

/** Remove reference TimeCycle **/
func (this *TimerEventDefinition) RemoveTimeCycle(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_timeCycle.(BaseElement).GetUUID() {
		this.M_timeCycle = nil
	}
}

/** TimeDuration **/
func (this *TimerEventDefinition) GetTimeDuration() Expression{
	return this.M_timeDuration
}

/** Init reference TimeDuration **/
func (this *TimerEventDefinition) SetTimeDuration(ref interface{}){
	this.NeedSave = true
	this.M_timeDuration = ref.(Expression)
}

/** Remove reference TimeDuration **/
func (this *TimerEventDefinition) RemoveTimeDuration(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_timeDuration.(BaseElement).GetUUID() {
		this.M_timeDuration = nil
	}
}

/** Lane **/
func (this *TimerEventDefinition) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *TimerEventDefinition) SetLanePtr(ref interface{}){
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
func (this *TimerEventDefinition) RemoveLanePtr(ref interface{}){
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
func (this *TimerEventDefinition) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *TimerEventDefinition) SetOutgoingPtr(ref interface{}){
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
func (this *TimerEventDefinition) RemoveOutgoingPtr(ref interface{}){
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
func (this *TimerEventDefinition) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *TimerEventDefinition) SetIncomingPtr(ref interface{}){
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
func (this *TimerEventDefinition) RemoveIncomingPtr(ref interface{}){
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
func (this *TimerEventDefinition) GetDefinitionsPtr() *Definitions{
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *TimerEventDefinition) SetDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	}else{
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *TimerEventDefinition) RemoveDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}

/** ThrowEvent **/
func (this *TimerEventDefinition) GetThrowEventPtr() []ThrowEvent{
	return this.m_throwEventPtr
}

/** Init reference ThrowEvent **/
func (this *TimerEventDefinition) SetThrowEventPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_throwEventPtr); i++ {
			if this.M_throwEventPtr[i] == refStr {
				return
			}
		}
		this.M_throwEventPtr = append(this.M_throwEventPtr, ref.(string))
	}else{
		this.RemoveThrowEventPtr(ref)
		this.m_throwEventPtr = append(this.m_throwEventPtr, ref.(ThrowEvent))
		this.M_throwEventPtr = append(this.M_throwEventPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ThrowEvent **/
func (this *TimerEventDefinition) RemoveThrowEventPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	throwEventPtr_ := make([]ThrowEvent, 0)
	throwEventPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_throwEventPtr); i++ {
		if toDelete.GetUUID() != this.m_throwEventPtr[i].(BaseElement).GetUUID() {
			throwEventPtr_ = append(throwEventPtr_, this.m_throwEventPtr[i])
			throwEventPtrUuid = append(throwEventPtrUuid, this.M_throwEventPtr[i])
		}
	}
	this.m_throwEventPtr = throwEventPtr_
	this.M_throwEventPtr = throwEventPtrUuid
}

/** CatchEvent **/
func (this *TimerEventDefinition) GetCatchEventPtr() []CatchEvent{
	return this.m_catchEventPtr
}

/** Init reference CatchEvent **/
func (this *TimerEventDefinition) SetCatchEventPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_catchEventPtr); i++ {
			if this.M_catchEventPtr[i] == refStr {
				return
			}
		}
		this.M_catchEventPtr = append(this.M_catchEventPtr, ref.(string))
	}else{
		this.RemoveCatchEventPtr(ref)
		this.m_catchEventPtr = append(this.m_catchEventPtr, ref.(CatchEvent))
		this.M_catchEventPtr = append(this.M_catchEventPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CatchEvent **/
func (this *TimerEventDefinition) RemoveCatchEventPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	catchEventPtr_ := make([]CatchEvent, 0)
	catchEventPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_catchEventPtr); i++ {
		if toDelete.GetUUID() != this.m_catchEventPtr[i].(BaseElement).GetUUID() {
			catchEventPtr_ = append(catchEventPtr_, this.m_catchEventPtr[i])
			catchEventPtrUuid = append(catchEventPtrUuid, this.M_catchEventPtr[i])
		}
	}
	this.m_catchEventPtr = catchEventPtr_
	this.M_catchEventPtr = catchEventPtrUuid
}

/** MultiInstanceLoopCharacteristics **/
func (this *TimerEventDefinition) GetMultiInstanceLoopCharacteristicsPtr() []*MultiInstanceLoopCharacteristics{
	return this.m_multiInstanceLoopCharacteristicsPtr
}

/** Init reference MultiInstanceLoopCharacteristics **/
func (this *TimerEventDefinition) SetMultiInstanceLoopCharacteristicsPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_multiInstanceLoopCharacteristicsPtr); i++ {
			if this.M_multiInstanceLoopCharacteristicsPtr[i] == refStr {
				return
			}
		}
		this.M_multiInstanceLoopCharacteristicsPtr = append(this.M_multiInstanceLoopCharacteristicsPtr, ref.(string))
	}else{
		this.RemoveMultiInstanceLoopCharacteristicsPtr(ref)
		this.m_multiInstanceLoopCharacteristicsPtr = append(this.m_multiInstanceLoopCharacteristicsPtr, ref.(*MultiInstanceLoopCharacteristics))
		this.M_multiInstanceLoopCharacteristicsPtr = append(this.M_multiInstanceLoopCharacteristicsPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference MultiInstanceLoopCharacteristics **/
func (this *TimerEventDefinition) RemoveMultiInstanceLoopCharacteristicsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	multiInstanceLoopCharacteristicsPtr_ := make([]*MultiInstanceLoopCharacteristics, 0)
	multiInstanceLoopCharacteristicsPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_multiInstanceLoopCharacteristicsPtr); i++ {
		if toDelete.GetUUID() != this.m_multiInstanceLoopCharacteristicsPtr[i].GetUUID() {
			multiInstanceLoopCharacteristicsPtr_ = append(multiInstanceLoopCharacteristicsPtr_, this.m_multiInstanceLoopCharacteristicsPtr[i])
			multiInstanceLoopCharacteristicsPtrUuid = append(multiInstanceLoopCharacteristicsPtrUuid, this.M_multiInstanceLoopCharacteristicsPtr[i])
		}
	}
	this.m_multiInstanceLoopCharacteristicsPtr = multiInstanceLoopCharacteristicsPtr_
	this.M_multiInstanceLoopCharacteristicsPtr = multiInstanceLoopCharacteristicsPtrUuid
}
