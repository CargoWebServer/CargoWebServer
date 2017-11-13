// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type Parameter struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
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
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference Name **/

/** Type **/
func (this *Parameter) GetType() string{
	return this.M_type
}

/** Init reference Type **/
func (this *Parameter) SetType(ref interface{}){
	if this.M_type != ref.(string) {
		this.M_type = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference Type **/

/** IsArray **/
func (this *Parameter) IsArray() bool{
	return this.M_isArray
}

/** Init reference IsArray **/
func (this *Parameter) SetIsArray(ref interface{}){
	if this.M_isArray != ref.(bool) {
		this.M_isArray = ref.(bool)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference IsArray **/

/** Parameters **/
func (this *Parameter) GetParametersPtr() *Parameter{
	return this.m_parametersPtr
}

/** Init reference Parameters **/
func (this *Parameter) SetParametersPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_parametersPtr != ref.(string) {
			this.M_parametersPtr = ref.(string)
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_parametersPtr != ref.(*Parameter).GetUUID() {
			this.M_parametersPtr = ref.(*Parameter).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_parametersPtr = ref.(*Parameter)
	}
}

/** Remove reference Parameters **/
func (this *Parameter) RemoveParametersPtr(ref interface{}){
	toDelete := ref.(*Parameter)
	if this.m_parametersPtr!= nil {
		if toDelete.GetUUID() == this.m_parametersPtr.GetUUID() {
			this.m_parametersPtr = nil
			this.M_parametersPtr = ""
			this.NeedSave = true
		}
	}
}
