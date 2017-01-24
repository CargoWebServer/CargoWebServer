// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type InputOutputBinding struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of BaseElement **/
	M_id    string
	m_other interface{}
	/** If the ref is a string and not an object **/
	M_other                string
	M_extensionElements    *ExtensionElements
	M_extensionDefinitions []*ExtensionDefinition
	M_extensionValues      []*ExtensionAttributeValue
	M_documentation        []*Documentation

	/** members of InputOutputBinding **/
	m_inputDataRef *InputSet
	/** If the ref is a string and not an object **/
	M_inputDataRef  string
	m_outputDataRef *OutputSet
	/** If the ref is a string and not an object **/
	M_outputDataRef string
	m_operationRef  *Operation
	/** If the ref is a string and not an object **/
	M_operationRef string

	/** Associations **/
	m_callableElementPtr CallableElement
	/** If the ref is a string and not an object **/
	M_callableElementPtr string
	m_lanePtr            []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
}

/** Xml parser for InputOutputBinding **/
type XsdInputOutputBinding struct {
	XMLName xml.Name `xml:"ioBinding"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_operationRef  string `xml:"operationRef,attr"`
	M_inputDataRef  string `xml:"inputDataRef,attr"`
	M_outputDataRef string `xml:"outputDataRef,attr"`
}

/** UUID **/
func (this *InputOutputBinding) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *InputOutputBinding) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *InputOutputBinding) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *InputOutputBinding) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *InputOutputBinding) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *InputOutputBinding) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *InputOutputBinding) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *InputOutputBinding) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *InputOutputBinding) SetExtensionDefinitions(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var extensionDefinitionss []*ExtensionDefinition
	for i := 0; i < len(this.M_extensionDefinitions); i++ {
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
func (this *InputOutputBinding) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *InputOutputBinding) SetExtensionValues(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var extensionValuess []*ExtensionAttributeValue
	for i := 0; i < len(this.M_extensionValues); i++ {
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
func (this *InputOutputBinding) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *InputOutputBinding) SetDocumentation(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var documentations []*Documentation
	for i := 0; i < len(this.M_documentation); i++ {
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
func (this *InputOutputBinding) RemoveDocumentation(ref interface{}) {
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

/** InputDataRef **/
func (this *InputOutputBinding) GetInputDataRef() *InputSet {
	return this.m_inputDataRef
}

/** Init reference InputDataRef **/
func (this *InputOutputBinding) SetInputDataRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_inputDataRef = ref.(string)
	} else {
		this.m_inputDataRef = ref.(*InputSet)
		this.M_inputDataRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference InputDataRef **/
func (this *InputOutputBinding) RemoveInputDataRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_inputDataRef.GetUUID() {
		this.m_inputDataRef = nil
		this.M_inputDataRef = ""
	}
}

/** OutputDataRef **/
func (this *InputOutputBinding) GetOutputDataRef() *OutputSet {
	return this.m_outputDataRef
}

/** Init reference OutputDataRef **/
func (this *InputOutputBinding) SetOutputDataRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_outputDataRef = ref.(string)
	} else {
		this.m_outputDataRef = ref.(*OutputSet)
		this.M_outputDataRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference OutputDataRef **/
func (this *InputOutputBinding) RemoveOutputDataRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_outputDataRef.GetUUID() {
		this.m_outputDataRef = nil
		this.M_outputDataRef = ""
	}
}

/** OperationRef **/
func (this *InputOutputBinding) GetOperationRef() *Operation {
	return this.m_operationRef
}

/** Init reference OperationRef **/
func (this *InputOutputBinding) SetOperationRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_operationRef = ref.(string)
	} else {
		this.m_operationRef = ref.(*Operation)
		this.M_operationRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference OperationRef **/
func (this *InputOutputBinding) RemoveOperationRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_operationRef.GetUUID() {
		this.m_operationRef = nil
		this.M_operationRef = ""
	}
}

/** CallableElement **/
func (this *InputOutputBinding) GetCallableElementPtr() CallableElement {
	return this.m_callableElementPtr
}

/** Init reference CallableElement **/
func (this *InputOutputBinding) SetCallableElementPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_callableElementPtr = ref.(string)
	} else {
		this.m_callableElementPtr = ref.(CallableElement)
		this.M_callableElementPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference CallableElement **/
func (this *InputOutputBinding) RemoveCallableElementPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_callableElementPtr.(BaseElement).GetUUID() {
		this.m_callableElementPtr = nil
		this.M_callableElementPtr = ""
	}
}

/** Lane **/
func (this *InputOutputBinding) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *InputOutputBinding) SetLanePtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_lanePtr); i++ {
			if this.M_lanePtr[i] == refStr {
				return
			}
		}
		this.M_lanePtr = append(this.M_lanePtr, ref.(string))
	} else {
		this.RemoveLanePtr(ref)
		this.m_lanePtr = append(this.m_lanePtr, ref.(*Lane))
		this.M_lanePtr = append(this.M_lanePtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Lane **/
func (this *InputOutputBinding) RemoveLanePtr(ref interface{}) {
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
func (this *InputOutputBinding) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *InputOutputBinding) SetOutgoingPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_outgoingPtr); i++ {
			if this.M_outgoingPtr[i] == refStr {
				return
			}
		}
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(string))
	} else {
		this.RemoveOutgoingPtr(ref)
		this.m_outgoingPtr = append(this.m_outgoingPtr, ref.(*Association))
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Outgoing **/
func (this *InputOutputBinding) RemoveOutgoingPtr(ref interface{}) {
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
func (this *InputOutputBinding) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *InputOutputBinding) SetIncomingPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_incomingPtr); i++ {
			if this.M_incomingPtr[i] == refStr {
				return
			}
		}
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(string))
	} else {
		this.RemoveIncomingPtr(ref)
		this.m_incomingPtr = append(this.m_incomingPtr, ref.(*Association))
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Incoming **/
func (this *InputOutputBinding) RemoveIncomingPtr(ref interface{}) {
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
