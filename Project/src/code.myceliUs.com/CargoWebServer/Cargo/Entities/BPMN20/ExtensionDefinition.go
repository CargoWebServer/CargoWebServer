// +build BPMN20

package BPMN20

type ExtensionDefinition struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of ExtensionDefinition **/
	M_name string
	M_extensionAttributeDefinition []*ExtensionAttributeDefinition


	/** Associations **/
	m_baseElementPtr []BaseElement
	/** If the ref is a string and not an object **/
	M_baseElementPtr []string
	m_extensionPtr *Extension
	/** If the ref is a string and not an object **/
	M_extensionPtr string
}/** UUID **/
func (this *ExtensionDefinition) GetUUID() string{
	return this.UUID
}

/** Name **/
func (this *ExtensionDefinition) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *ExtensionDefinition) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** ExtensionAttributeDefinition **/
func (this *ExtensionDefinition) GetExtensionAttributeDefinition() []*ExtensionAttributeDefinition{
	return this.M_extensionAttributeDefinition
}

/** Init reference ExtensionAttributeDefinition **/
func (this *ExtensionDefinition) SetExtensionAttributeDefinition(ref interface{}){
	this.NeedSave = true
	isExist := false
	var extensionAttributeDefinitions []*ExtensionAttributeDefinition
	for i:=0; i<len(this.M_extensionAttributeDefinition); i++ {
		if this.M_extensionAttributeDefinition[i].GetUUID() != ref.(*ExtensionAttributeDefinition).GetUUID() {
			extensionAttributeDefinitions = append(extensionAttributeDefinitions, this.M_extensionAttributeDefinition[i])
		} else {
			isExist = true
			extensionAttributeDefinitions = append(extensionAttributeDefinitions, ref.(*ExtensionAttributeDefinition))
		}
	}
	if !isExist {
		extensionAttributeDefinitions = append(extensionAttributeDefinitions, ref.(*ExtensionAttributeDefinition))
	}
	this.M_extensionAttributeDefinition = extensionAttributeDefinitions
}

/** Remove reference ExtensionAttributeDefinition **/
func (this *ExtensionDefinition) RemoveExtensionAttributeDefinition(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*ExtensionAttributeDefinition)
	extensionAttributeDefinition_ := make([]*ExtensionAttributeDefinition, 0)
	for i := 0; i < len(this.M_extensionAttributeDefinition); i++ {
		if toDelete.GetUUID() != this.M_extensionAttributeDefinition[i].GetUUID() {
			extensionAttributeDefinition_ = append(extensionAttributeDefinition_, this.M_extensionAttributeDefinition[i])
		}
	}
	this.M_extensionAttributeDefinition = extensionAttributeDefinition_
}

/** BaseElement **/
func (this *ExtensionDefinition) GetBaseElementPtr() []BaseElement{
	return this.m_baseElementPtr
}

/** Init reference BaseElement **/
func (this *ExtensionDefinition) SetBaseElementPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_baseElementPtr); i++ {
			if this.M_baseElementPtr[i] == refStr {
				return
			}
		}
		this.M_baseElementPtr = append(this.M_baseElementPtr, ref.(string))
	}else{
		this.RemoveBaseElementPtr(ref)
		this.m_baseElementPtr = append(this.m_baseElementPtr, ref.(BaseElement))
		this.M_baseElementPtr = append(this.M_baseElementPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference BaseElement **/
func (this *ExtensionDefinition) RemoveBaseElementPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	baseElementPtr_ := make([]BaseElement, 0)
	baseElementPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_baseElementPtr); i++ {
		if toDelete.GetUUID() != this.m_baseElementPtr[i].(BaseElement).GetUUID() {
			baseElementPtr_ = append(baseElementPtr_, this.m_baseElementPtr[i])
			baseElementPtrUuid = append(baseElementPtrUuid, this.M_baseElementPtr[i])
		}
	}
	this.m_baseElementPtr = baseElementPtr_
	this.M_baseElementPtr = baseElementPtrUuid
}

/** Extension **/
func (this *ExtensionDefinition) GetExtensionPtr() *Extension{
	return this.m_extensionPtr
}

/** Init reference Extension **/
func (this *ExtensionDefinition) SetExtensionPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_extensionPtr = ref.(string)
	}else{
		this.m_extensionPtr = ref.(*Extension)
		this.M_extensionPtr = ref.(*Extension).GetUUID()
	}
}

/** Remove reference Extension **/
func (this *ExtensionDefinition) RemoveExtensionPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Extension)
	if toDelete.GetUUID() == this.m_extensionPtr.GetUUID() {
		this.m_extensionPtr = nil
		this.M_extensionPtr = ""
	}
}
