package BPMN20

import(
"encoding/xml"
)

type DataStoreReference struct{

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

	/** members of ItemAwareElement **/
	m_itemSubjectRef *ItemDefinition
	/** If the ref is a string and not an object **/
	M_itemSubjectRef string
	M_dataState *DataState

	/** members of FlowElement **/
	M_name string
	M_auditing *Auditing
	M_monitoring *Monitoring
	m_categoryValueRef []*CategoryValue
	/** If the ref is a string and not an object **/
	M_categoryValueRef []string

	/** members of DataStoreReference **/
	m_dataStoreRef *DataStore
	/** If the ref is a string and not an object **/
	M_dataStoreRef string


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
	m_dataAssociationPtr []DataAssociation
	/** If the ref is a string and not an object **/
	M_dataAssociationPtr []string
	m_multiInstanceLoopCharacteristicsPtr *MultiInstanceLoopCharacteristics
	/** If the ref is a string and not an object **/
	M_multiInstanceLoopCharacteristicsPtr string
	m_containerPtr FlowElementsContainer
	/** If the ref is a string and not an object **/
	M_containerPtr string
}

/** Xml parser for DataStoreReference **/
type XsdDataStoreReference struct {
	XMLName xml.Name	`xml:"dataStoreReference"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** FlowElement **/
	M_auditing	*XsdAuditing	`xml:"auditing,omitempty"`
	M_monitoring	*XsdMonitoring	`xml:"monitoring,omitempty"`
	M_categoryValueRef	[]string	`xml:"categoryValueRef"`
	M_name	string	`xml:"name,attr"`


	M_dataState	*XsdDataState	`xml:"dataState,omitempty"`
	M_itemSubjectRef	string	`xml:"itemSubjectRef,attr"`
	M_dataStoreRef	string	`xml:"dataStoreRef,attr"`

}
/** UUID **/
func (this *DataStoreReference) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *DataStoreReference) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *DataStoreReference) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *DataStoreReference) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *DataStoreReference) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *DataStoreReference) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *DataStoreReference) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *DataStoreReference) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *DataStoreReference) SetExtensionDefinitions(ref interface{}){
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
func (this *DataStoreReference) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *DataStoreReference) SetExtensionValues(ref interface{}){
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
func (this *DataStoreReference) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *DataStoreReference) SetDocumentation(ref interface{}){
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
func (this *DataStoreReference) RemoveDocumentation(ref interface{}){
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

/** ItemSubjectRef **/
func (this *DataStoreReference) GetItemSubjectRef() *ItemDefinition{
	return this.m_itemSubjectRef
}

/** Init reference ItemSubjectRef **/
func (this *DataStoreReference) SetItemSubjectRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_itemSubjectRef = ref.(string)
	}else{
		this.m_itemSubjectRef = ref.(*ItemDefinition)
		this.M_itemSubjectRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ItemSubjectRef **/
func (this *DataStoreReference) RemoveItemSubjectRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_itemSubjectRef.GetUUID() {
		this.m_itemSubjectRef = nil
		this.M_itemSubjectRef = ""
	}
}

/** DataState **/
func (this *DataStoreReference) GetDataState() *DataState{
	return this.M_dataState
}

/** Init reference DataState **/
func (this *DataStoreReference) SetDataState(ref interface{}){
	this.NeedSave = true
	this.M_dataState = ref.(*DataState)
}

/** Remove reference DataState **/
func (this *DataStoreReference) RemoveDataState(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_dataState.GetUUID() {
		this.M_dataState = nil
	}
}

/** Name **/
func (this *DataStoreReference) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *DataStoreReference) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Auditing **/
func (this *DataStoreReference) GetAuditing() *Auditing{
	return this.M_auditing
}

/** Init reference Auditing **/
func (this *DataStoreReference) SetAuditing(ref interface{}){
	this.NeedSave = true
	this.M_auditing = ref.(*Auditing)
}

/** Remove reference Auditing **/
func (this *DataStoreReference) RemoveAuditing(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_auditing.GetUUID() {
		this.M_auditing = nil
	}
}

/** Monitoring **/
func (this *DataStoreReference) GetMonitoring() *Monitoring{
	return this.M_monitoring
}

/** Init reference Monitoring **/
func (this *DataStoreReference) SetMonitoring(ref interface{}){
	this.NeedSave = true
	this.M_monitoring = ref.(*Monitoring)
}

/** Remove reference Monitoring **/
func (this *DataStoreReference) RemoveMonitoring(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_monitoring.GetUUID() {
		this.M_monitoring = nil
	}
}

/** CategoryValueRef **/
func (this *DataStoreReference) GetCategoryValueRef() []*CategoryValue{
	return this.m_categoryValueRef
}

/** Init reference CategoryValueRef **/
func (this *DataStoreReference) SetCategoryValueRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_categoryValueRef); i++ {
			if this.M_categoryValueRef[i] == refStr {
				return
			}
		}
		this.M_categoryValueRef = append(this.M_categoryValueRef, ref.(string))
	}else{
		this.RemoveCategoryValueRef(ref)
		this.m_categoryValueRef = append(this.m_categoryValueRef, ref.(*CategoryValue))
		this.M_categoryValueRef = append(this.M_categoryValueRef, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference CategoryValueRef **/
func (this *DataStoreReference) RemoveCategoryValueRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	categoryValueRef_ := make([]*CategoryValue, 0)
	categoryValueRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_categoryValueRef); i++ {
		if toDelete.GetUUID() != this.m_categoryValueRef[i].GetUUID() {
			categoryValueRef_ = append(categoryValueRef_, this.m_categoryValueRef[i])
			categoryValueRefUuid = append(categoryValueRefUuid, this.M_categoryValueRef[i])
		}
	}
	this.m_categoryValueRef = categoryValueRef_
	this.M_categoryValueRef = categoryValueRefUuid
}

/** DataStoreRef **/
func (this *DataStoreReference) GetDataStoreRef() *DataStore{
	return this.m_dataStoreRef
}

/** Init reference DataStoreRef **/
func (this *DataStoreReference) SetDataStoreRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_dataStoreRef = ref.(string)
	}else{
		this.m_dataStoreRef = ref.(*DataStore)
		this.M_dataStoreRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference DataStoreRef **/
func (this *DataStoreReference) RemoveDataStoreRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_dataStoreRef.GetUUID() {
		this.m_dataStoreRef = nil
		this.M_dataStoreRef = ""
	}
}

/** Lane **/
func (this *DataStoreReference) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *DataStoreReference) SetLanePtr(ref interface{}){
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
func (this *DataStoreReference) RemoveLanePtr(ref interface{}){
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
func (this *DataStoreReference) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *DataStoreReference) SetOutgoingPtr(ref interface{}){
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
func (this *DataStoreReference) RemoveOutgoingPtr(ref interface{}){
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
func (this *DataStoreReference) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *DataStoreReference) SetIncomingPtr(ref interface{}){
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
func (this *DataStoreReference) RemoveIncomingPtr(ref interface{}){
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

/** DataAssociation **/
func (this *DataStoreReference) GetDataAssociationPtr() []DataAssociation{
	return this.m_dataAssociationPtr
}

/** Init reference DataAssociation **/
func (this *DataStoreReference) SetDataAssociationPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_dataAssociationPtr); i++ {
			if this.M_dataAssociationPtr[i] == refStr {
				return
			}
		}
		this.M_dataAssociationPtr = append(this.M_dataAssociationPtr, ref.(string))
	}else{
		this.RemoveDataAssociationPtr(ref)
		this.m_dataAssociationPtr = append(this.m_dataAssociationPtr, ref.(DataAssociation))
		this.M_dataAssociationPtr = append(this.M_dataAssociationPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference DataAssociation **/
func (this *DataStoreReference) RemoveDataAssociationPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataAssociationPtr_ := make([]DataAssociation, 0)
	dataAssociationPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_dataAssociationPtr); i++ {
		if toDelete.GetUUID() != this.m_dataAssociationPtr[i].(BaseElement).GetUUID() {
			dataAssociationPtr_ = append(dataAssociationPtr_, this.m_dataAssociationPtr[i])
			dataAssociationPtrUuid = append(dataAssociationPtrUuid, this.M_dataAssociationPtr[i])
		}
	}
	this.m_dataAssociationPtr = dataAssociationPtr_
	this.M_dataAssociationPtr = dataAssociationPtrUuid
}

/** MultiInstanceLoopCharacteristics **/
func (this *DataStoreReference) GetMultiInstanceLoopCharacteristicsPtr() *MultiInstanceLoopCharacteristics{
	return this.m_multiInstanceLoopCharacteristicsPtr
}

/** Init reference MultiInstanceLoopCharacteristics **/
func (this *DataStoreReference) SetMultiInstanceLoopCharacteristicsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(string)
	}else{
		this.m_multiInstanceLoopCharacteristicsPtr = ref.(*MultiInstanceLoopCharacteristics)
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference MultiInstanceLoopCharacteristics **/
func (this *DataStoreReference) RemoveMultiInstanceLoopCharacteristicsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_multiInstanceLoopCharacteristicsPtr.GetUUID() {
		this.m_multiInstanceLoopCharacteristicsPtr = nil
		this.M_multiInstanceLoopCharacteristicsPtr = ""
	}
}

/** Container **/
func (this *DataStoreReference) GetContainerPtr() FlowElementsContainer{
	return this.m_containerPtr
}

/** Init reference Container **/
func (this *DataStoreReference) SetContainerPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_containerPtr = ref.(string)
	}else{
		this.m_containerPtr = ref.(FlowElementsContainer)
		this.M_containerPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Container **/
func (this *DataStoreReference) RemoveContainerPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_containerPtr.(BaseElement).GetUUID() {
		this.m_containerPtr = nil
		this.M_containerPtr = ""
	}
}
