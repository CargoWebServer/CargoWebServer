// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Performer_impl struct {

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

	/** members of ResourceRole **/
	m_resourceRef *Resource
	/** If the ref is a string and not an object **/
	M_resourceRef                  string
	M_resourceParameterBinding     []*ResourceParameterBinding
	M_resourceAssignmentExpression *ResourceAssignmentExpression
	M_name                         string

	/** members of Performer **/
	/** No members **/

	/** Associations **/
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr   []string
	m_globalTaskPtr GlobalTask
	/** If the ref is a string and not an object **/
	M_globalTaskPtr string
	m_processPtr    *Process
	/** If the ref is a string and not an object **/
	M_processPtr  string
	m_activityPtr Activity
	/** If the ref is a string and not an object **/
	M_activityPtr string
}

/** Xml parser for Performer **/
type XsdPerformer struct {
	XMLName xml.Name `xml:"performer"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	/** ResourceRole **/
	M_name                         string                           `xml:"name,attr"`
	M_resourceRef                  *string                          `xml:"resourceRef"`
	M_resourceParameterBinding     []*XsdResourceParameterBinding   `xml:"resourceParameterBinding,omitempty"`
	M_resourceAssignmentExpression *XsdResourceAssignmentExpression `xml:"resourceAssignmentExpression,omitempty"`
}

/** UUID **/
func (this *Performer_impl) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Performer_impl) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Performer_impl) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Performer_impl) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Performer_impl) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Performer_impl) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Performer_impl) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Performer_impl) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Performer_impl) SetExtensionDefinitions(ref interface{}) {
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
func (this *Performer_impl) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Performer_impl) SetExtensionValues(ref interface{}) {
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
func (this *Performer_impl) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Performer_impl) SetDocumentation(ref interface{}) {
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
func (this *Performer_impl) RemoveDocumentation(ref interface{}) {
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

/** ResourceRef **/
func (this *Performer_impl) GetResourceRef() *Resource {
	return this.m_resourceRef
}

/** Init reference ResourceRef **/
func (this *Performer_impl) SetResourceRef(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_resourceRef = ref.(string)
	} else {
		this.m_resourceRef = ref.(*Resource)
		this.M_resourceRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ResourceRef **/
func (this *Performer_impl) RemoveResourceRef(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_resourceRef.GetUUID() {
		this.m_resourceRef = nil
		this.M_resourceRef = ""
	}
}

/** ResourceParameterBinding **/
func (this *Performer_impl) GetResourceParameterBinding() []*ResourceParameterBinding {
	return this.M_resourceParameterBinding
}

/** Init reference ResourceParameterBinding **/
func (this *Performer_impl) SetResourceParameterBinding(ref interface{}) {
	this.NeedSave = true
	isExist := false
	var resourceParameterBindings []*ResourceParameterBinding
	for i := 0; i < len(this.M_resourceParameterBinding); i++ {
		if this.M_resourceParameterBinding[i].GetUUID() != ref.(BaseElement).GetUUID() {
			resourceParameterBindings = append(resourceParameterBindings, this.M_resourceParameterBinding[i])
		} else {
			isExist = true
			resourceParameterBindings = append(resourceParameterBindings, ref.(*ResourceParameterBinding))
		}
	}
	if !isExist {
		resourceParameterBindings = append(resourceParameterBindings, ref.(*ResourceParameterBinding))
	}
	this.M_resourceParameterBinding = resourceParameterBindings
}

/** Remove reference ResourceParameterBinding **/
func (this *Performer_impl) RemoveResourceParameterBinding(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	resourceParameterBinding_ := make([]*ResourceParameterBinding, 0)
	for i := 0; i < len(this.M_resourceParameterBinding); i++ {
		if toDelete.GetUUID() != this.M_resourceParameterBinding[i].GetUUID() {
			resourceParameterBinding_ = append(resourceParameterBinding_, this.M_resourceParameterBinding[i])
		}
	}
	this.M_resourceParameterBinding = resourceParameterBinding_
}

/** ResourceAssignmentExpression **/
func (this *Performer_impl) GetResourceAssignmentExpression() *ResourceAssignmentExpression {
	return this.M_resourceAssignmentExpression
}

/** Init reference ResourceAssignmentExpression **/
func (this *Performer_impl) SetResourceAssignmentExpression(ref interface{}) {
	this.NeedSave = true
	this.M_resourceAssignmentExpression = ref.(*ResourceAssignmentExpression)
}

/** Remove reference ResourceAssignmentExpression **/
func (this *Performer_impl) RemoveResourceAssignmentExpression(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.M_resourceAssignmentExpression.GetUUID() {
		this.M_resourceAssignmentExpression = nil
	}
}

/** Name **/
func (this *Performer_impl) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Performer_impl) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Lane **/
func (this *Performer_impl) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Performer_impl) SetLanePtr(ref interface{}) {
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
func (this *Performer_impl) RemoveLanePtr(ref interface{}) {
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
func (this *Performer_impl) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Performer_impl) SetOutgoingPtr(ref interface{}) {
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
func (this *Performer_impl) RemoveOutgoingPtr(ref interface{}) {
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
func (this *Performer_impl) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Performer_impl) SetIncomingPtr(ref interface{}) {
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
func (this *Performer_impl) RemoveIncomingPtr(ref interface{}) {
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

/** GlobalTask **/
func (this *Performer_impl) GetGlobalTaskPtr() GlobalTask {
	return this.m_globalTaskPtr
}

/** Init reference GlobalTask **/
func (this *Performer_impl) SetGlobalTaskPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_globalTaskPtr = ref.(string)
	} else {
		this.m_globalTaskPtr = ref.(GlobalTask)
		this.M_globalTaskPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference GlobalTask **/
func (this *Performer_impl) RemoveGlobalTaskPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_globalTaskPtr.(BaseElement).GetUUID() {
		this.m_globalTaskPtr = nil
		this.M_globalTaskPtr = ""
	}
}

/** Process **/
func (this *Performer_impl) GetProcessPtr() *Process {
	return this.m_processPtr
}

/** Init reference Process **/
func (this *Performer_impl) SetProcessPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_processPtr = ref.(string)
	} else {
		this.m_processPtr = ref.(*Process)
		this.M_processPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Process **/
func (this *Performer_impl) RemoveProcessPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_processPtr.GetUUID() {
		this.m_processPtr = nil
		this.M_processPtr = ""
	}
}

/** Activity **/
func (this *Performer_impl) GetActivityPtr() Activity {
	return this.m_activityPtr
}

/** Init reference Activity **/
func (this *Performer_impl) SetActivityPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_activityPtr = ref.(string)
	} else {
		this.m_activityPtr = ref.(Activity)
		this.M_activityPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Activity **/
func (this *Performer_impl) RemoveActivityPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_activityPtr.(BaseElement).GetUUID() {
		this.m_activityPtr = nil
		this.M_activityPtr = ""
	}
}
