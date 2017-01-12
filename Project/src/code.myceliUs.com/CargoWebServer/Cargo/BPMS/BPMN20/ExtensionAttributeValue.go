package BPMN20

type ExtensionAttributeValue struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of ExtensionAttributeValue **/
	m_valueRef interface{}
	/** If the ref is a string and not an object **/
	M_valueRef string
	m_value interface{}
	/** If the ref is a string and not an object **/
	M_value string
	m_extensionAttributeDefinition *ExtensionAttributeDefinition
	/** If the ref is a string and not an object **/
	M_extensionAttributeDefinition string


	/** Associations **/
	m_baseElementPtr BaseElement
	/** If the ref is a string and not an object **/
	M_baseElementPtr string
}/** UUID **/
func (this *ExtensionAttributeValue) GetUUID() string{
	return this.UUID
}

/** ValueRef **/
func (this *ExtensionAttributeValue) GetValueRef() interface{}{
	return this.m_valueRef
}

/** Init reference ValueRef **/
func (this *ExtensionAttributeValue) SetValueRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_valueRef = ref.(string)
	}else{
		this.m_valueRef = ref.(interface{})
	}
}

/** Remove reference ValueRef **/

/** Value **/
func (this *ExtensionAttributeValue) GetValue() interface{}{
	return this.m_value
}

/** Init reference Value **/
func (this *ExtensionAttributeValue) SetValue(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_value = ref.(string)
	}else{
		this.m_value = ref.(interface{})
	}
}

/** Remove reference Value **/

/** ExtensionAttributeDefinition **/
func (this *ExtensionAttributeValue) GetExtensionAttributeDefinition() *ExtensionAttributeDefinition{
	return this.m_extensionAttributeDefinition
}

/** Init reference ExtensionAttributeDefinition **/
func (this *ExtensionAttributeValue) SetExtensionAttributeDefinition(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_extensionAttributeDefinition = ref.(string)
	}else{
		this.m_extensionAttributeDefinition = ref.(*ExtensionAttributeDefinition)
		this.M_extensionAttributeDefinition = ref.(*ExtensionAttributeDefinition).GetName()
	}
}

/** Remove reference ExtensionAttributeDefinition **/

/** BaseElement **/
func (this *ExtensionAttributeValue) GetBaseElementPtr() BaseElement{
	return this.m_baseElementPtr
}

/** Init reference BaseElement **/
func (this *ExtensionAttributeValue) SetBaseElementPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_baseElementPtr = ref.(string)
	}else{
		this.m_baseElementPtr = ref.(BaseElement)
		this.M_baseElementPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference BaseElement **/
func (this *ExtensionAttributeValue) RemoveBaseElementPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_baseElementPtr.(BaseElement).GetUUID() {
		this.m_baseElementPtr = nil
		this.M_baseElementPtr = ""
	}
}
