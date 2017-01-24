// +build BPMN
package BPMN20

type ExtensionAttributeDefinition struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of ExtensionAttributeDefinition **/
	M_name                string
	M_type                string
	M_isReference         bool
	m_extensionDefinition *ExtensionDefinition
	/** If the ref is a string and not an object **/
	M_extensionDefinition string

	/** Associations **/
	m_extensionAttributeValuePtr []*ExtensionAttributeValue
	/** If the ref is a string and not an object **/
	M_extensionAttributeValuePtr []string
} /** UUID **/
func (this *ExtensionAttributeDefinition) GetUUID() string {
	return this.UUID
}

/** Name **/
func (this *ExtensionAttributeDefinition) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *ExtensionAttributeDefinition) SetName(ref interface{}) {
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Type **/
func (this *ExtensionAttributeDefinition) GetType() string {
	return this.M_type
}

/** Init reference Type **/
func (this *ExtensionAttributeDefinition) SetType(ref interface{}) {
	this.NeedSave = true
	this.M_type = ref.(string)
}

/** Remove reference Type **/

/** IsReference **/
func (this *ExtensionAttributeDefinition) IsReference() bool {
	return this.M_isReference
}

/** Init reference IsReference **/
func (this *ExtensionAttributeDefinition) SetIsReference(ref interface{}) {
	this.NeedSave = true
	this.M_isReference = ref.(bool)
}

/** Remove reference IsReference **/

/** ExtensionDefinition **/
func (this *ExtensionAttributeDefinition) GetExtensionDefinition() *ExtensionDefinition {
	return this.m_extensionDefinition
}

/** Init reference ExtensionDefinition **/
func (this *ExtensionAttributeDefinition) SetExtensionDefinition(ref interface{}) {
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_extensionDefinition = ref.(string)
	} else {
		this.m_extensionDefinition = ref.(*ExtensionDefinition)
		this.M_extensionDefinition = ref.(*ExtensionDefinition).GetName()
	}
}

/** Remove reference ExtensionDefinition **/

/** ExtensionAttributeValue **/
func (this *ExtensionAttributeDefinition) GetExtensionAttributeValuePtr() []*ExtensionAttributeValue {
	return this.m_extensionAttributeValuePtr
}

/** Init reference ExtensionAttributeValue **/
func (this *ExtensionAttributeDefinition) SetExtensionAttributeValuePtr(ref interface{}) {
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_extensionAttributeValuePtr); i++ {
			if this.M_extensionAttributeValuePtr[i] == refStr {
				return
			}
		}
		this.M_extensionAttributeValuePtr = append(this.M_extensionAttributeValuePtr, ref.(string))
	} else {
		this.m_extensionAttributeValuePtr = append(this.m_extensionAttributeValuePtr, ref.(*ExtensionAttributeValue))
	}
}

/** Remove reference ExtensionAttributeValue **/
