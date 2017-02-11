package CargoEntities

import(
"encoding/xml"
)

type Parameter struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Parameter **/
	M_name string
	M_type string
	M_isArray bool


	/** Associations **/
	m_parametersPtr *Parameter
	/** If the ref is a string and not an object **/
	M_parametersPtr string
}

/** Xml parser for Parameter **/
type XsdParameter struct {
	XMLName xml.Name	`xml:"parameter"`
	M_name	string	`xml:"name,attr"`
	M_type	string	`xml:"type,attr"`
	M_isArray	bool	`xml:"isArray,attr"`

}
/** UUID **/
func (this *Parameter) GetUUID() string{
	return this.UUID
}

/** Name **/
func (this *Parameter) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Parameter) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Type **/
func (this *Parameter) GetType() string{
	return this.M_type
}

/** Init reference Type **/
func (this *Parameter) SetType(ref interface{}){
	this.NeedSave = true
	this.M_type = ref.(string)
}

/** Remove reference Type **/

/** IsArray **/
func (this *Parameter) IsArray() bool{
	return this.M_isArray
}

/** Init reference IsArray **/
func (this *Parameter) SetIsArray(ref interface{}){
	this.NeedSave = true
	this.M_isArray = ref.(bool)
}

/** Remove reference IsArray **/

/** Parameters **/
func (this *Parameter) GetParametersPtr() *Parameter{
	return this.m_parametersPtr
}

/** Init reference Parameters **/
func (this *Parameter) SetParametersPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parametersPtr = ref.(string)
	}else{
		this.m_parametersPtr = ref.(*Parameter)
		this.M_parametersPtr = ref.(*Parameter).GetName()
	}
}

/** Remove reference Parameters **/
