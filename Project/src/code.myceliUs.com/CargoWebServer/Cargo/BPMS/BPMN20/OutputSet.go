// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type OutputSet struct {

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

	/** members of OutputSet **/
	m_dataOutputRefs []*DataOutput
	/** If the ref is a string and not an object **/
	M_dataOutputRefs []string
	M_name           string
	m_inputSetRefs   []*InputSet
	/** If the ref is a string and not an object **/
	M_inputSetRefs       []string
	m_optionalOutputRefs []*DataOutput
	/** If the ref is a string and not an object **/
	M_optionalOutputRefs       []string
	m_whileExecutingOutputRefs []*DataOutput
	/** If the ref is a string and not an object **/
	M_whileExecutingOutputRefs []string

	/** Associations **/
	m_catchEventPtr CatchEvent
	/** If the ref is a string and not an object **/
	M_catchEventPtr               string
	m_inputOutputSpecificationPtr *InputOutputSpecification
	/** If the ref is a string and not an object **/
	M_inputOutputSpecificationPtr string
	m_inputOutputBindingPtr       []*InputOutputBinding
	/** If the ref is a string and not an object **/
	M_inputOutputBindingPtr []string
	m_lanePtr               []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
}

/** Xml parser for OutputSet **/
type XsdOutputSet struct {
	XMLName xml.Name `xml:"outputSet"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	M_dataOutputRefs           []string `xml:"dataOutputRefs"`
	M_optionalOutputRefs       []string `xml:"optionalOutputRefs"`
	M_whileExecutingOutputRefs []string `xml:"whileExecutingOutputRefs"`
	M_inputSetRefs             []string `xml:"inputSetRefs"`
	M_name                     string   `xml:"name,attr"`
}

/** UUID **/
func (this *OutputSet) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *OutputSet) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *OutputSet) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *OutputSet) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *OutputSet) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *OutputSet) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *OutputSet) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *OutputSet) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *OutputSet) SetExtensionDefinitions(ref interface{}) {
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
func (this *OutputSet) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *OutputSet) SetExtensionValues(ref interface{}) {
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
func (this *OutputSet) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *OutputSet) SetDocumentation(ref interface{}) {
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
func (this *OutputSet) RemoveDocumentation(ref interface{}) {
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

/** DataOutputRefs **/
func (this *OutputSet) GetDataOutputRefs() []*DataOutput {
	return this.m_dataOutputRefs
}

/** Init reference DataOutputRefs **/
func (this *OutputSet) SetDataOutputRefs(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_dataOutputRefs); i++ {
			if this.M_dataOutputRefs[i] == refStr {
				return
			}
		}
		this.M_dataOutputRefs = append(this.M_dataOutputRefs, ref.(string))
	} else {
		this.RemoveDataOutputRefs(ref)
		this.m_dataOutputRefs = append(this.m_dataOutputRefs, ref.(*DataOutput))
		this.M_dataOutputRefs = append(this.M_dataOutputRefs, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference DataOutputRefs **/
func (this *OutputSet) RemoveDataOutputRefs(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	dataOutputRefs_ := make([]*DataOutput, 0)
	dataOutputRefsUuid := make([]string, 0)
	for i := 0; i < len(this.m_dataOutputRefs); i++ {
		if toDelete.GetUUID() != this.m_dataOutputRefs[i].GetUUID() {
			dataOutputRefs_ = append(dataOutputRefs_, this.m_dataOutputRefs[i])
			dataOutputRefsUuid = append(dataOutputRefsUuid, this.M_dataOutputRefs[i])
		}
	}
	this.m_dataOutputRefs = dataOutputRefs_
	this.M_dataOutputRefs = dataOutputRefsUuid
}

/** Name **/
func (this *OutputSet) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *OutputSet) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** InputSetRefs **/
func (this *OutputSet) GetInputSetRefs() []*InputSet {
	return this.m_inputSetRefs
}

/** Init reference InputSetRefs **/
func (this *OutputSet) SetInputSetRefs(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_inputSetRefs); i++ {
			if this.M_inputSetRefs[i] == refStr {
				return
			}
		}
		this.M_inputSetRefs = append(this.M_inputSetRefs, ref.(string))
	} else {
		this.RemoveInputSetRefs(ref)
		this.m_inputSetRefs = append(this.m_inputSetRefs, ref.(*InputSet))
		this.M_inputSetRefs = append(this.M_inputSetRefs, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference InputSetRefs **/
func (this *OutputSet) RemoveInputSetRefs(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	inputSetRefs_ := make([]*InputSet, 0)
	inputSetRefsUuid := make([]string, 0)
	for i := 0; i < len(this.m_inputSetRefs); i++ {
		if toDelete.GetUUID() != this.m_inputSetRefs[i].GetUUID() {
			inputSetRefs_ = append(inputSetRefs_, this.m_inputSetRefs[i])
			inputSetRefsUuid = append(inputSetRefsUuid, this.M_inputSetRefs[i])
		}
	}
	this.m_inputSetRefs = inputSetRefs_
	this.M_inputSetRefs = inputSetRefsUuid
}

/** OptionalOutputRefs **/
func (this *OutputSet) GetOptionalOutputRefs() []*DataOutput {
	return this.m_optionalOutputRefs
}

/** Init reference OptionalOutputRefs **/
func (this *OutputSet) SetOptionalOutputRefs(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_optionalOutputRefs); i++ {
			if this.M_optionalOutputRefs[i] == refStr {
				return
			}
		}
		this.M_optionalOutputRefs = append(this.M_optionalOutputRefs, ref.(string))
	} else {
		this.RemoveOptionalOutputRefs(ref)
		this.m_optionalOutputRefs = append(this.m_optionalOutputRefs, ref.(*DataOutput))
		this.M_optionalOutputRefs = append(this.M_optionalOutputRefs, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference OptionalOutputRefs **/
func (this *OutputSet) RemoveOptionalOutputRefs(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	optionalOutputRefs_ := make([]*DataOutput, 0)
	optionalOutputRefsUuid := make([]string, 0)
	for i := 0; i < len(this.m_optionalOutputRefs); i++ {
		if toDelete.GetUUID() != this.m_optionalOutputRefs[i].GetUUID() {
			optionalOutputRefs_ = append(optionalOutputRefs_, this.m_optionalOutputRefs[i])
			optionalOutputRefsUuid = append(optionalOutputRefsUuid, this.M_optionalOutputRefs[i])
		}
	}
	this.m_optionalOutputRefs = optionalOutputRefs_
	this.M_optionalOutputRefs = optionalOutputRefsUuid
}

/** WhileExecutingOutputRefs **/
func (this *OutputSet) GetWhileExecutingOutputRefs() []*DataOutput {
	return this.m_whileExecutingOutputRefs
}

/** Init reference WhileExecutingOutputRefs **/
func (this *OutputSet) SetWhileExecutingOutputRefs(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_whileExecutingOutputRefs); i++ {
			if this.M_whileExecutingOutputRefs[i] == refStr {
				return
			}
		}
		this.M_whileExecutingOutputRefs = append(this.M_whileExecutingOutputRefs, ref.(string))
	} else {
		this.RemoveWhileExecutingOutputRefs(ref)
		this.m_whileExecutingOutputRefs = append(this.m_whileExecutingOutputRefs, ref.(*DataOutput))
		this.M_whileExecutingOutputRefs = append(this.M_whileExecutingOutputRefs, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference WhileExecutingOutputRefs **/
func (this *OutputSet) RemoveWhileExecutingOutputRefs(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	whileExecutingOutputRefs_ := make([]*DataOutput, 0)
	whileExecutingOutputRefsUuid := make([]string, 0)
	for i := 0; i < len(this.m_whileExecutingOutputRefs); i++ {
		if toDelete.GetUUID() != this.m_whileExecutingOutputRefs[i].GetUUID() {
			whileExecutingOutputRefs_ = append(whileExecutingOutputRefs_, this.m_whileExecutingOutputRefs[i])
			whileExecutingOutputRefsUuid = append(whileExecutingOutputRefsUuid, this.M_whileExecutingOutputRefs[i])
		}
	}
	this.m_whileExecutingOutputRefs = whileExecutingOutputRefs_
	this.M_whileExecutingOutputRefs = whileExecutingOutputRefsUuid
}

/** CatchEvent **/
func (this *OutputSet) GetCatchEventPtr() CatchEvent {
	return this.m_catchEventPtr
}

/** Init reference CatchEvent **/
func (this *OutputSet) SetCatchEventPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_catchEventPtr = ref.(string)
	} else {
		this.m_catchEventPtr = ref.(CatchEvent)
		this.M_catchEventPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference CatchEvent **/
func (this *OutputSet) RemoveCatchEventPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_catchEventPtr.(BaseElement).GetUUID() {
		this.m_catchEventPtr = nil
		this.M_catchEventPtr = ""
	}
}

/** InputOutputSpecification **/
func (this *OutputSet) GetInputOutputSpecificationPtr() *InputOutputSpecification {
	return this.m_inputOutputSpecificationPtr
}

/** Init reference InputOutputSpecification **/
func (this *OutputSet) SetInputOutputSpecificationPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_inputOutputSpecificationPtr = ref.(string)
	} else {
		this.m_inputOutputSpecificationPtr = ref.(*InputOutputSpecification)
		this.M_inputOutputSpecificationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference InputOutputSpecification **/
func (this *OutputSet) RemoveInputOutputSpecificationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_inputOutputSpecificationPtr.GetUUID() {
		this.m_inputOutputSpecificationPtr = nil
		this.M_inputOutputSpecificationPtr = ""
	}
}

/** InputOutputBinding **/
func (this *OutputSet) GetInputOutputBindingPtr() []*InputOutputBinding {
	return this.m_inputOutputBindingPtr
}

/** Init reference InputOutputBinding **/
func (this *OutputSet) SetInputOutputBindingPtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_inputOutputBindingPtr); i++ {
			if this.M_inputOutputBindingPtr[i] == refStr {
				return
			}
		}
		this.M_inputOutputBindingPtr = append(this.M_inputOutputBindingPtr, ref.(string))
	} else {
		this.RemoveInputOutputBindingPtr(ref)
		this.m_inputOutputBindingPtr = append(this.m_inputOutputBindingPtr, ref.(*InputOutputBinding))
		this.M_inputOutputBindingPtr = append(this.M_inputOutputBindingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference InputOutputBinding **/
func (this *OutputSet) RemoveInputOutputBindingPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	inputOutputBindingPtr_ := make([]*InputOutputBinding, 0)
	inputOutputBindingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_inputOutputBindingPtr); i++ {
		if toDelete.GetUUID() != this.m_inputOutputBindingPtr[i].GetUUID() {
			inputOutputBindingPtr_ = append(inputOutputBindingPtr_, this.m_inputOutputBindingPtr[i])
			inputOutputBindingPtrUuid = append(inputOutputBindingPtrUuid, this.M_inputOutputBindingPtr[i])
		}
	}
	this.m_inputOutputBindingPtr = inputOutputBindingPtr_
	this.M_inputOutputBindingPtr = inputOutputBindingPtrUuid
}

/** Lane **/
func (this *OutputSet) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *OutputSet) SetLanePtr(ref interface{}) {
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
func (this *OutputSet) RemoveLanePtr(ref interface{}) {
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
func (this *OutputSet) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *OutputSet) SetOutgoingPtr(ref interface{}) {
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
func (this *OutputSet) RemoveOutgoingPtr(ref interface{}) {
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
func (this *OutputSet) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *OutputSet) SetIncomingPtr(ref interface{}) {
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
func (this *OutputSet) RemoveIncomingPtr(ref interface{}) {
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
