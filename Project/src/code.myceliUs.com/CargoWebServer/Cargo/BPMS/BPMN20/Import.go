package BPMN20

import(
"encoding/xml"
)

type Import struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Import **/
	M_importType string
	M_location string
	M_namespace string


	/** Associations **/
	m_itemDefinitionPtr []*ItemDefinition
	/** If the ref is a string and not an object **/
	M_itemDefinitionPtr []string
	m_definitionsPtr *Definitions
	/** If the ref is a string and not an object **/
	M_definitionsPtr string
}

/** Xml parser for Import **/
type XsdImport struct {
	XMLName xml.Name	`xml:"import"`
	M_namespace	string	`xml:"namespace,attr"`
	M_location	string	`xml:"location,attr"`
	M_importType	string	`xml:"importType,attr"`

}
/** UUID **/
func (this *Import) GetUUID() string{
	return this.UUID
}

/** ImportType **/
func (this *Import) GetImportType() string{
	return this.M_importType
}

/** Init reference ImportType **/
func (this *Import) SetImportType(ref interface{}){
	this.NeedSave = true
	this.M_importType = ref.(string)
}

/** Remove reference ImportType **/

/** Location **/
func (this *Import) GetLocation() string{
	return this.M_location
}

/** Init reference Location **/
func (this *Import) SetLocation(ref interface{}){
	this.NeedSave = true
	this.M_location = ref.(string)
}

/** Remove reference Location **/

/** Namespace **/
func (this *Import) GetNamespace() string{
	return this.M_namespace
}

/** Init reference Namespace **/
func (this *Import) SetNamespace(ref interface{}){
	this.NeedSave = true
	this.M_namespace = ref.(string)
}

/** Remove reference Namespace **/

/** ItemDefinition **/
func (this *Import) GetItemDefinitionPtr() []*ItemDefinition{
	return this.m_itemDefinitionPtr
}

/** Init reference ItemDefinition **/
func (this *Import) SetItemDefinitionPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_itemDefinitionPtr); i++ {
			if this.M_itemDefinitionPtr[i] == refStr {
				return
			}
		}
		this.M_itemDefinitionPtr = append(this.M_itemDefinitionPtr, ref.(string))
	}else{
		this.RemoveItemDefinitionPtr(ref)
		this.m_itemDefinitionPtr = append(this.m_itemDefinitionPtr, ref.(*ItemDefinition))
		this.M_itemDefinitionPtr = append(this.M_itemDefinitionPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ItemDefinition **/
func (this *Import) RemoveItemDefinitionPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	itemDefinitionPtr_ := make([]*ItemDefinition, 0)
	itemDefinitionPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_itemDefinitionPtr); i++ {
		if toDelete.GetUUID() != this.m_itemDefinitionPtr[i].GetUUID() {
			itemDefinitionPtr_ = append(itemDefinitionPtr_, this.m_itemDefinitionPtr[i])
			itemDefinitionPtrUuid = append(itemDefinitionPtrUuid, this.M_itemDefinitionPtr[i])
		}
	}
	this.m_itemDefinitionPtr = itemDefinitionPtr_
	this.M_itemDefinitionPtr = itemDefinitionPtrUuid
}

/** Definitions **/
func (this *Import) GetDefinitionsPtr() *Definitions{
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *Import) SetDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	}else{
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *Import) RemoveDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}
