package BPMN20

import(
"encoding/xml"
)

type DataStore struct{

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

	/** members of ItemAwareElement **/
	m_itemSubjectRef *ItemDefinition
	/** If the ref is a string and not an object **/
	M_itemSubjectRef string
	M_dataState *DataState

	/** members of DataStore **/
	M_name string
	M_capacity int
	M_isUnlimited bool


	/** Associations **/
	m_dataStoreReferencePtr []*DataStoreReference
	/** If the ref is a string and not an object **/
	M_dataStoreReferencePtr []string
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
	m_dataAssociationPtr []DataAssociation
	/** If the ref is a string and not an object **/
	M_dataAssociationPtr []string
	m_multiInstanceLoopCharacteristicsPtr *MultiInstanceLoopCharacteristics
	/** If the ref is a string and not an object **/
	M_multiInstanceLoopCharacteristicsPtr string
}

/** Xml parser for DataStore **/
type XsdDataStore struct {
	XMLName xml.Name	`xml:"dataStore"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** RootElement **/


	M_dataState	*XsdDataState	`xml:"dataState,omitempty"`
	M_name	string	`xml:"name,attr"`
	M_capacity	int	`xml:"capacity,attr"`
	M_isUnlimited	bool	`xml:"isUnlimited,attr"`
	M_itemSubjectRef	string	`xml:"itemSubjectRef,attr"`

}
/** UUID **/
func (this *DataStore) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *DataStore) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *DataStore) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *DataStore) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *DataStore) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *DataStore) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *DataStore) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *DataStore) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *DataStore) SetExtensionDefinitions(ref interface{}){
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
func (this *DataStore) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *DataStore) SetExtensionValues(ref interface{}){
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
func (this *DataStore) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *DataStore) SetDocumentation(ref interface{}){
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
func (this *DataStore) RemoveDocumentation(ref interface{}){
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
func (this *DataStore) GetItemSubjectRef() *ItemDefinition{
	return this.m_itemSubjectRef
}

/** Init reference ItemSubjectRef **/
func (this *DataStore) SetItemSubjectRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_itemSubjectRef = ref.(string)
	}else{
		this.m_itemSubjectRef = ref.(*ItemDefinition)
		this.M_itemSubjectRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ItemSubjectRef **/
func (this *DataStore) RemoveItemSubjectRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_itemSubjectRef.GetUUID() {
		this.m_itemSubjectRef = nil
		this.M_itemSubjectRef = ""
	}
}

/** DataState **/
func (this *DataStore) GetDataState() *DataState{
	return this.M_dataState
}

/** Init reference DataState **/
func (this *DataStore) SetDataState(ref interface{}){
	this.NeedSave = true
	this.M_dataState = ref.(*DataState)
}

/** Remove reference DataState **/
func (this *DataStore) RemoveDataState(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_dataState.GetUUID() {
		this.M_dataState = nil
	}
}

/** Name **/
func (this *DataStore) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *DataStore) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Capacity **/
func (this *DataStore) GetCapacity() int{
	return this.M_capacity
}

/** Init reference Capacity **/
func (this *DataStore) SetCapacity(ref interface{}){
	this.NeedSave = true
	this.M_capacity = ref.(int)
}

/** Remove reference Capacity **/

/** IsUnlimited **/
func (this *DataStore) IsUnlimited() bool{
	return this.M_isUnlimited
}

/** Init reference IsUnlimited **/
func (this *DataStore) SetIsUnlimited(ref interface{}){
	this.NeedSave = true
	this.M_isUnlimited = ref.(bool)
}

/** Remove reference IsUnlimited **/

/** DataStoreReference **/
func (this *DataStore) GetDataStoreReferencePtr() []*DataStoreReference{
	return this.m_dataStoreReferencePtr
}

/** Init reference DataStoreReference **/
func (this *DataStore) SetDataStoreReferencePtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_dataStoreReferencePtr); i++ {
			if this.M_dataStoreReferencePtr[i] == refStr {
				return
			}
		}
		this.M_dataStoreReferencePtr = append(this.M_dataStoreReferencePtr, ref.(string))
	}else{
		this.RemoveDataStoreReferencePtr(ref)
		this.m_dataStoreReferencePtr = append(this.m_dataStoreReferencePtr, ref.(*DataStoreReference))
		this.M_dataStoreReferencePtr = append(this.M_dataStoreReferencePtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference DataStoreReference **/
func (this *DataStore) RemoveDataStoreReferencePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataStoreReferencePtr_ := make([]*DataStoreReference, 0)
	dataStoreReferencePtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_dataStoreReferencePtr); i++ {
		if toDelete.GetUUID() != this.m_dataStoreReferencePtr[i].GetUUID() {
			dataStoreReferencePtr_ = append(dataStoreReferencePtr_, this.m_dataStoreReferencePtr[i])
			dataStoreReferencePtrUuid = append(dataStoreReferencePtrUuid, this.M_dataStoreReferencePtr[i])
		}
	}
	this.m_dataStoreReferencePtr = dataStoreReferencePtr_
	this.M_dataStoreReferencePtr = dataStoreReferencePtrUuid
}

/** Lane **/
func (this *DataStore) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *DataStore) SetLanePtr(ref interface{}){
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
func (this *DataStore) RemoveLanePtr(ref interface{}){
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
func (this *DataStore) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *DataStore) SetOutgoingPtr(ref interface{}){
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
func (this *DataStore) RemoveOutgoingPtr(ref interface{}){
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
func (this *DataStore) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *DataStore) SetIncomingPtr(ref interface{}){
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
func (this *DataStore) RemoveIncomingPtr(ref interface{}){
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
func (this *DataStore) GetDefinitionsPtr() *Definitions{
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *DataStore) SetDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	}else{
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *DataStore) RemoveDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}

/** DataAssociation **/
func (this *DataStore) GetDataAssociationPtr() []DataAssociation{
	return this.m_dataAssociationPtr
}

/** Init reference DataAssociation **/
func (this *DataStore) SetDataAssociationPtr(ref interface{}){
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
func (this *DataStore) RemoveDataAssociationPtr(ref interface{}){
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
func (this *DataStore) GetMultiInstanceLoopCharacteristicsPtr() *MultiInstanceLoopCharacteristics{
	return this.m_multiInstanceLoopCharacteristicsPtr
}

/** Init reference MultiInstanceLoopCharacteristics **/
func (this *DataStore) SetMultiInstanceLoopCharacteristicsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(string)
	}else{
		this.m_multiInstanceLoopCharacteristicsPtr = ref.(*MultiInstanceLoopCharacteristics)
		this.M_multiInstanceLoopCharacteristicsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference MultiInstanceLoopCharacteristics **/
func (this *DataStore) RemoveMultiInstanceLoopCharacteristicsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_multiInstanceLoopCharacteristicsPtr.GetUUID() {
		this.m_multiInstanceLoopCharacteristicsPtr = nil
		this.M_multiInstanceLoopCharacteristicsPtr = ""
	}
}
