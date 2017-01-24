// +build BPMN
package BPMN20

import (
	"encoding/xml"
)

type Extension struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of Extension **/
	M_mustUnderstand      bool
	M_documentation       []*Documentation
	M_extensionDefinition *ExtensionDefinition

	/** Associations **/
	m_definitionsPtr *Definitions
	/** If the ref is a string and not an object **/
	M_definitionsPtr string
}

/** Xml parser for Extension **/
type XsdExtension struct {
	XMLName               xml.Name            `xml:"extension"`
	M_documentation       []*XsdDocumentation `xml:"documentation,omitempty"`
	M_extensionDefinition string              `xml:"extensionDefinition,attr"`
	M_mustUnderstand      bool                `xml:"mustUnderstand,attr"`
}

/** UUID **/
func (this *Extension) GetUUID() string {
	return this.UUID
}

/** MustUnderstand **/
func (this *Extension) GetMustUnderstand() bool {
	return this.M_mustUnderstand
}

/** Init reference MustUnderstand **/
func (this *Extension) SetMustUnderstand(ref interface{}) {
	this.NeedSave = true
	this.M_mustUnderstand = ref.(bool)
}

/** Remove reference MustUnderstand **/

/** Documentation **/
func (this *Extension) GetDocumentation() []*Documentation {
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Extension) SetDocumentation(ref interface{}) {
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
func (this *Extension) RemoveDocumentation(ref interface{}) {
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

/** ExtensionDefinition **/
func (this *Extension) GetExtensionDefinition() *ExtensionDefinition {
	return this.M_extensionDefinition
}

/** Init reference ExtensionDefinition **/
func (this *Extension) SetExtensionDefinition(ref interface{}) {
	this.NeedSave = true
	this.M_extensionDefinition = ref.(*ExtensionDefinition)
}

/** Remove reference ExtensionDefinition **/

/** Definitions **/
func (this *Extension) GetDefinitionsPtr() *Definitions {
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *Extension) SetDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	} else {
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *Extension) RemoveDefinitionsPtr(ref interface{}) {
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}
