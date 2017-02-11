package BPMS

import(
"encoding/xml"
)

type CorrelationInfo struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of CorrelationInfo **/
	M_id string


	/** Associations **/
	m_runtimesPtr *Runtimes
	/** If the ref is a string and not an object **/
	M_runtimesPtr string
}

/** Xml parser for CorrelationInfo **/
type XsdCorrelationInfo struct {
	XMLName xml.Name	`xml:"correlationInfo"`
	M_id	string	`xml:"id,attr"`

}
/** UUID **/
func (this *CorrelationInfo) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *CorrelationInfo) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *CorrelationInfo) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Runtimes **/
func (this *CorrelationInfo) GetRuntimesPtr() *Runtimes{
	return this.m_runtimesPtr
}

/** Init reference Runtimes **/
func (this *CorrelationInfo) SetRuntimesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_runtimesPtr = ref.(string)
	}else{
		this.m_runtimesPtr = ref.(*Runtimes)
		this.M_runtimesPtr = ref.(*Runtimes).GetUUID()
	}
}

/** Remove reference Runtimes **/
func (this *CorrelationInfo) RemoveRuntimesPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Runtimes)
	if toDelete.GetUUID() == this.m_runtimesPtr.GetUUID() {
		this.m_runtimesPtr = nil
		this.M_runtimesPtr = ""
	}
}
