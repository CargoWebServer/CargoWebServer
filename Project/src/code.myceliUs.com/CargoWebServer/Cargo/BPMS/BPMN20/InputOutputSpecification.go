// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type InputOutputSpecification struct {

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

	/** members of InputOutputSpecification **/
	M_inputSet   []*InputSet
	M_outputSet  []*OutputSet
	M_dataInput  []*DataInput
	M_dataOutput []*DataOutput

	/** Associations **/
	m_callableElementPtr CallableElement
	/** If the ref is a string and not an object **/
	M_callableElementPtr string
	m_activityPtr        Activity
	/** If the ref is a string and not an object **/
	M_activityPtr string
	m_lanePtr     []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
}

/** Xml parser for InputOutputSpecification **/
type XsdInputOutputSpecification struct {
	XMLName xml.Name `xml:"ioSpecification"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_dataInput  []*XsdDataInput  `xml:"dataInput,omitempty"`
	M_dataOutput []*XsdDataOutput `xml:"dataOutput,omitempty"`
	M_inputSet   []*XsdInputSet   `xml:"inputSet,omitempty"`
	M_outputSet  []*XsdOutputSet  `xml:"outputSet,omitempty"`
}

/** UUID **/
func (this *InputOutputSpecification) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *InputOutputSpecification) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *InputOutputSpecification) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *InputOutputSpecification) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *InputOutputSpecification) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *InputOutputSpecification) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *InputOutputSpecification) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *InputOutputSpecification) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *InputOutputSpecification) SetExtensionDefinitions(ref interface{}) {
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
func (this *InputOutputSpecification) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *InputOutputSpecification) SetExtensionValues(ref interface{}) {
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
func (this *InputOutputSpecification) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *InputOutputSpecification) SetDocumentation(ref interface{}) {
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
func (this *InputOutputSpecification) RemoveDocumentation(ref interface{}) {
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

/** InputSet **/
func (this *InputOutputSpecification) GetInputSet() []*InputSet {
	return this.M_inputSet
}

/** Init reference InputSet **/
func (this *InputOutputSpecification) SetInputSet(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var inputSets []*InputSet
	for i := 0; i < len(this.M_inputSet); i++ {
		if this.M_inputSet[i].GetUUID() != ref.(BaseElement).GetUUID() {
			inputSets = append(inputSets, this.M_inputSet[i])
		} else {
			isExist = true
			inputSets = append(inputSets, ref.(*InputSet))
		}
	}
	if !isExist {
		inputSets = append(inputSets, ref.(*InputSet))
	}
	this.M_inputSet = inputSets
}

/** Remove reference InputSet **/
func (this *InputOutputSpecification) RemoveInputSet(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	inputSet_ := make([]*InputSet, 0)
	for i := 0; i < len(this.M_inputSet); i++ {
		if toDelete.GetUUID() != this.M_inputSet[i].GetUUID() {
			inputSet_ = append(inputSet_, this.M_inputSet[i])
		}
	}
	this.M_inputSet = inputSet_
}

/** OutputSet **/
func (this *InputOutputSpecification) GetOutputSet() []*OutputSet {
	return this.M_outputSet
}

/** Init reference OutputSet **/
func (this *InputOutputSpecification) SetOutputSet(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var outputSets []*OutputSet
	for i := 0; i < len(this.M_outputSet); i++ {
		if this.M_outputSet[i].GetUUID() != ref.(BaseElement).GetUUID() {
			outputSets = append(outputSets, this.M_outputSet[i])
		} else {
			isExist = true
			outputSets = append(outputSets, ref.(*OutputSet))
		}
	}
	if !isExist {
		outputSets = append(outputSets, ref.(*OutputSet))
	}
	this.M_outputSet = outputSets
}

/** Remove reference OutputSet **/
func (this *InputOutputSpecification) RemoveOutputSet(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	outputSet_ := make([]*OutputSet, 0)
	for i := 0; i < len(this.M_outputSet); i++ {
		if toDelete.GetUUID() != this.M_outputSet[i].GetUUID() {
			outputSet_ = append(outputSet_, this.M_outputSet[i])
		}
	}
	this.M_outputSet = outputSet_
}

/** DataInput **/
func (this *InputOutputSpecification) GetDataInput() []*DataInput {
	return this.M_dataInput
}

/** Init reference DataInput **/
func (this *InputOutputSpecification) SetDataInput(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var dataInputs []*DataInput
	for i := 0; i < len(this.M_dataInput); i++ {
		if this.M_dataInput[i].GetUUID() != ref.(BaseElement).GetUUID() {
			dataInputs = append(dataInputs, this.M_dataInput[i])
		} else {
			isExist = true
			dataInputs = append(dataInputs, ref.(*DataInput))
		}
	}
	if !isExist {
		dataInputs = append(dataInputs, ref.(*DataInput))
	}
	this.M_dataInput = dataInputs
}

/** Remove reference DataInput **/
func (this *InputOutputSpecification) RemoveDataInput(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataInput_ := make([]*DataInput, 0)
	for i := 0; i < len(this.M_dataInput); i++ {
		if toDelete.GetUUID() != this.M_dataInput[i].GetUUID() {
			dataInput_ = append(dataInput_, this.M_dataInput[i])
		}
	}
	this.M_dataInput = dataInput_
}

/** DataOutput **/
func (this *InputOutputSpecification) GetDataOutput() []*DataOutput {
	return this.M_dataOutput
}

/** Init reference DataOutput **/
func (this *InputOutputSpecification) SetDataOutput(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var dataOutputs []*DataOutput
	for i := 0; i < len(this.M_dataOutput); i++ {
		if this.M_dataOutput[i].GetUUID() != ref.(BaseElement).GetUUID() {
			dataOutputs = append(dataOutputs, this.M_dataOutput[i])
		} else {
			isExist = true
			dataOutputs = append(dataOutputs, ref.(*DataOutput))
		}
	}
	if !isExist {
		dataOutputs = append(dataOutputs, ref.(*DataOutput))
	}
	this.M_dataOutput = dataOutputs
}

/** Remove reference DataOutput **/
func (this *InputOutputSpecification) RemoveDataOutput(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataOutput_ := make([]*DataOutput, 0)
	for i := 0; i < len(this.M_dataOutput); i++ {
		if toDelete.GetUUID() != this.M_dataOutput[i].GetUUID() {
			dataOutput_ = append(dataOutput_, this.M_dataOutput[i])
		}
	}
	this.M_dataOutput = dataOutput_
}

/** CallableElement **/
func (this *InputOutputSpecification) GetCallableElementPtr() CallableElement {
	return this.m_callableElementPtr
}

/** Init reference CallableElement **/
func (this *InputOutputSpecification) SetCallableElementPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_callableElementPtr = ref.(string)
	} else {
		this.m_callableElementPtr = ref.(CallableElement)
		this.M_callableElementPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference CallableElement **/
func (this *InputOutputSpecification) RemoveCallableElementPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_callableElementPtr.(BaseElement).GetUUID() {
		this.m_callableElementPtr = nil
		this.M_callableElementPtr = ""
	}
}

/** Activity **/
func (this *InputOutputSpecification) GetActivityPtr() Activity {
	return this.m_activityPtr
}

/** Init reference Activity **/
func (this *InputOutputSpecification) SetActivityPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_activityPtr = ref.(string)
	} else {
		this.m_activityPtr = ref.(Activity)
		this.M_activityPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Activity **/
func (this *InputOutputSpecification) RemoveActivityPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_activityPtr.(BaseElement).GetUUID() {
		this.m_activityPtr = nil
		this.M_activityPtr = ""
	}
}

/** Lane **/
func (this *InputOutputSpecification) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *InputOutputSpecification) SetLanePtr(ref interface{}) {
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
func (this *InputOutputSpecification) RemoveLanePtr(ref interface{}) {
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
func (this *InputOutputSpecification) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *InputOutputSpecification) SetOutgoingPtr(ref interface{}) {
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
func (this *InputOutputSpecification) RemoveOutgoingPtr(ref interface{}) {
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
func (this *InputOutputSpecification) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *InputOutputSpecification) SetIncomingPtr(ref interface{}) {
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
func (this *InputOutputSpecification) RemoveIncomingPtr(ref interface{}) {
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
