// +build CargoEntities

package CargoEntities

import(
"encoding/xml"
)

type Action struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Action **/
	M_name string
	M_parameters []*Parameter
	M_results []*Parameter
	M_accessType AccessType


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for Action **/
type XsdAction struct {
	XMLName xml.Name	`xml:"action"`
	M_name	string	`xml:"name,attr"`

}
/** UUID **/
func (this *Action) GetUUID() string{
	return this.UUID
}

/** Name **/
func (this *Action) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Action) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Parameters **/
func (this *Action) GetParameters() []*Parameter{
	return this.M_parameters
}

/** Init reference Parameters **/
func (this *Action) SetParameters(ref interface{}){
	this.NeedSave = true
	isExist := false
	var parameterss []*Parameter
	for i:=0; i<len(this.M_parameters); i++ {
		if this.M_parameters[i].GetUUID() != ref.(*Parameter).GetUUID() {
			parameterss = append(parameterss, this.M_parameters[i])
		} else {
			isExist = true
			parameterss = append(parameterss, ref.(*Parameter))
		}
	}
	if !isExist {
		parameterss = append(parameterss, ref.(*Parameter))
	}
	this.M_parameters = parameterss
}

/** Remove reference Parameters **/
func (this *Action) RemoveParameters(ref interface{}){
	toDelete := ref.(*Parameter)
	parameters_ := make([]*Parameter, 0)
	for i := 0; i < len(this.M_parameters); i++ {
		if toDelete.GetUUID() != this.M_parameters[i].GetUUID() {
			parameters_ = append(parameters_, this.M_parameters[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_parameters = parameters_
}

/** Results **/
func (this *Action) GetResults() []*Parameter{
	return this.M_results
}

/** Init reference Results **/
func (this *Action) SetResults(ref interface{}){
	this.NeedSave = true
	isExist := false
	var resultss []*Parameter
	for i:=0; i<len(this.M_results); i++ {
		if this.M_results[i].GetUUID() != ref.(*Parameter).GetUUID() {
			resultss = append(resultss, this.M_results[i])
		} else {
			isExist = true
			resultss = append(resultss, ref.(*Parameter))
		}
	}
	if !isExist {
		resultss = append(resultss, ref.(*Parameter))
	}
	this.M_results = resultss
}

/** Remove reference Results **/
func (this *Action) RemoveResults(ref interface{}){
	toDelete := ref.(*Parameter)
	results_ := make([]*Parameter, 0)
	for i := 0; i < len(this.M_results); i++ {
		if toDelete.GetUUID() != this.M_results[i].GetUUID() {
			results_ = append(results_, this.M_results[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_results = results_
}

/** AccessType **/
func (this *Action) GetAccessType() AccessType{
	return this.M_accessType
}

/** Init reference AccessType **/
func (this *Action) SetAccessType(ref interface{}){
	this.NeedSave = true
	this.M_accessType = ref.(AccessType)
}

/** Remove reference AccessType **/

/** Entities **/
func (this *Action) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Action) SetEntitiesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *Action) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
		}else{
			this.NeedSave = true
		}
	}
}
