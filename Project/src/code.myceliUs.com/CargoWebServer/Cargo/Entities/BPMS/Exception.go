package BPMS

import(
"encoding/xml"
)

type Exception struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Exception **/
	M_id string
	M_exceptionType ExceptionType


	/** Associations **/
	m_runtimesPtr *Runtimes
	/** If the ref is a string and not an object **/
	M_runtimesPtr string
}

/** Xml parser for Exception **/
type XsdException struct {
	XMLName xml.Name	`xml:"exception"`
	M_name	string	`xml:"name,attr"`
	M_exceptionType	string	`xml:"exceptionType,attr"`

}
/** UUID **/
func (this *Exception) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Exception) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Exception) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** ExceptionType **/
func (this *Exception) GetExceptionType() ExceptionType{
	return this.M_exceptionType
}

/** Init reference ExceptionType **/
func (this *Exception) SetExceptionType(ref interface{}){
	this.NeedSave = true
	this.M_exceptionType = ref.(ExceptionType)
}

/** Remove reference ExceptionType **/

/** Runtimes **/
func (this *Exception) GetRuntimesPtr() *Runtimes{
	return this.m_runtimesPtr
}

/** Init reference Runtimes **/
func (this *Exception) SetRuntimesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_runtimesPtr = ref.(string)
	}else{
		this.m_runtimesPtr = ref.(*Runtimes)
		this.M_runtimesPtr = ref.(*Runtimes).GetUUID()
	}
}

/** Remove reference Runtimes **/
func (this *Exception) RemoveRuntimesPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Runtimes)
	if toDelete.GetUUID() == this.m_runtimesPtr.GetUUID() {
		this.m_runtimesPtr = nil
		this.M_runtimesPtr = ""
	}
}
