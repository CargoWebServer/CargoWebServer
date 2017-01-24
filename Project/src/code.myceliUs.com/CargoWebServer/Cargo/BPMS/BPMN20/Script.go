// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Script struct {

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

	/** members of Artifact **/
	/** No members **/

	/** members of Script **/
	M_script string

	/** Associations **/
	m_scriptTaskPtr *ScriptTask
	/** If the ref is a string and not an object **/
	M_scriptTaskPtr       string
	m_globalScriptTaskPtr *GlobalScriptTask
	/** If the ref is a string and not an object **/
	M_globalScriptTaskPtr string
	m_lanePtr             []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr     []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
	m_processPtr  *Process
	/** If the ref is a string and not an object **/
	M_processPtr       string
	m_collaborationPtr Collaboration
	/** If the ref is a string and not an object **/
	M_collaborationPtr   string
	m_subChoreographyPtr *SubChoreography
	/** If the ref is a string and not an object **/
	M_subChoreographyPtr string
	m_subProcessPtr      SubProcess
	/** If the ref is a string and not an object **/
	M_subProcessPtr string
}

/** Xml parser for Script **/
type XsdScript struct {
	XMLName xml.Name `xml:"script"`
	/** BaseElement **/
	M_documentation     []*XsdDocumentation   `xml:"documentation,omitempty"`
	M_extensionElements *XsdExtensionElements `xml:"extensionElements,omitempty"`
	M_id                string                `xml:"id,attr"`
	//	M_other	string	`xml:",innerxml"`

	/** Artifact **/

	M_script string `xml:",innerxml"`
}

/** UUID **/
func (this *Script) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Script) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Script) SetId(ref interface{}) {
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Script) GetOther() interface{} {
	return this.M_other
}

/** Init reference Other **/
func (this *Script) SetOther(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	} else {
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Script) GetExtensionElements() *ExtensionElements {
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Script) SetExtensionElements(ref interface{}) {
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Script) GetExtensionDefinitions() []*ExtensionDefinition {
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Script) SetExtensionDefinitions(ref interface{}) {
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
func (this *Script) GetExtensionValues() []*ExtensionAttributeValue {
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Script) SetExtensionValues(ref interface{}) {
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
func (this *Script) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Script) SetDocumentation(ref interface{}) {
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
func (this *Script) RemoveDocumentation(ref interface{}) {
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

/** Script **/
func (this *Script) GetScript() string {
	return this.M_script
}

/** Init reference Script **/
func (this *Script) SetScript(ref interface{}) {
	this.NeedSave = true
	this.M_script = ref.(string)
}

/** Remove reference Script **/

/** ScriptTask **/
func (this *Script) GetScriptTaskPtr() *ScriptTask {
	return this.m_scriptTaskPtr
}

/** Init reference ScriptTask **/
func (this *Script) SetScriptTaskPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_scriptTaskPtr = ref.(string)
	} else {
		this.m_scriptTaskPtr = ref.(*ScriptTask)
		this.M_scriptTaskPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference ScriptTask **/
func (this *Script) RemoveScriptTaskPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_scriptTaskPtr.GetUUID() {
		this.m_scriptTaskPtr = nil
		this.M_scriptTaskPtr = ""
	}
}

/** GlobalScriptTask **/
func (this *Script) GetGlobalScriptTaskPtr() *GlobalScriptTask {
	return this.m_globalScriptTaskPtr
}

/** Init reference GlobalScriptTask **/
func (this *Script) SetGlobalScriptTaskPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_globalScriptTaskPtr = ref.(string)
	} else {
		this.m_globalScriptTaskPtr = ref.(*GlobalScriptTask)
		this.M_globalScriptTaskPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference GlobalScriptTask **/
func (this *Script) RemoveGlobalScriptTaskPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_globalScriptTaskPtr.GetUUID() {
		this.m_globalScriptTaskPtr = nil
		this.M_globalScriptTaskPtr = ""
	}
}

/** Lane **/
func (this *Script) GetLanePtr() []*Lane {
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Script) SetLanePtr(ref interface{}) {
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
func (this *Script) RemoveLanePtr(ref interface{}) {
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
func (this *Script) GetOutgoingPtr() []*Association {
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Script) SetOutgoingPtr(ref interface{}) {
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
func (this *Script) RemoveOutgoingPtr(ref interface{}) {
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
func (this *Script) GetIncomingPtr() []*Association {
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Script) SetIncomingPtr(ref interface{}) {
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
func (this *Script) RemoveIncomingPtr(ref interface{}) {
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

/** Process **/
func (this *Script) GetProcessPtr() *Process {
	return this.m_processPtr
}

/** Init reference Process **/
func (this *Script) SetProcessPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_processPtr = ref.(string)
	} else {
		this.m_processPtr = ref.(*Process)
		this.M_processPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Process **/
func (this *Script) RemoveProcessPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_processPtr.GetUUID() {
		this.m_processPtr = nil
		this.M_processPtr = ""
	}
}

/** Collaboration **/
func (this *Script) GetCollaborationPtr() Collaboration {
	return this.m_collaborationPtr
}

/** Init reference Collaboration **/
func (this *Script) SetCollaborationPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_collaborationPtr = ref.(string)
	} else {
		this.m_collaborationPtr = ref.(Collaboration)
		this.M_collaborationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Collaboration **/
func (this *Script) RemoveCollaborationPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_collaborationPtr.(BaseElement).GetUUID() {
		this.m_collaborationPtr = nil
		this.M_collaborationPtr = ""
	}
}

/** SubChoreography **/
func (this *Script) GetSubChoreographyPtr() *SubChoreography {
	return this.m_subChoreographyPtr
}

/** Init reference SubChoreography **/
func (this *Script) SetSubChoreographyPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_subChoreographyPtr = ref.(string)
	} else {
		this.m_subChoreographyPtr = ref.(*SubChoreography)
		this.M_subChoreographyPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference SubChoreography **/
func (this *Script) RemoveSubChoreographyPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_subChoreographyPtr.GetUUID() {
		this.m_subChoreographyPtr = nil
		this.M_subChoreographyPtr = ""
	}
}

/** SubProcess **/
func (this *Script) GetSubProcessPtr() SubProcess {
	return this.m_subProcessPtr
}

/** Init reference SubProcess **/
func (this *Script) SetSubProcessPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_subProcessPtr = ref.(string)
	} else {
		this.m_subProcessPtr = ref.(SubProcess)
		this.M_subProcessPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference SubProcess **/
func (this *Script) RemoveSubProcessPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_subProcessPtr.(BaseElement).GetUUID() {
		this.m_subProcessPtr = nil
		this.M_subProcessPtr = ""
	}
}
