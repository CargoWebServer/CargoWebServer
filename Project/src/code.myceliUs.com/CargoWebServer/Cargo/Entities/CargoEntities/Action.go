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
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool

	/** members of Action **/
	M_name string
	M_doc string
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
	M_doc	string	`xml:"doc,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Action) GetUuid() string{
	return this.UUID
}
func (this *Action) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Action) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_name)
	return ids
}

/** The type name **/
func (this *Action) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Action"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Action) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Action) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Action) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Action) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *Action) IsNeedSave() bool{
	return this.NeedSave
}


/** Name **/
func (this *Action) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Action) SetName(ref interface{}){
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Name **/

/** Doc **/
func (this *Action) GetDoc() string{
	return this.M_doc
}

/** Init reference Doc **/
func (this *Action) SetDoc(ref interface{}){
	if this.M_doc != ref.(string) {
		this.M_doc = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Doc **/

/** Parameters **/
func (this *Action) GetParameters() []*Parameter{
	return this.M_parameters
}

/** Init reference Parameters **/
func (this *Action) SetParameters(ref interface{}){
	isExist := false
	var parameterss []*Parameter
	for i:=0; i<len(this.M_parameters); i++ {
		if this.M_parameters[i].GetUuid() != ref.(*Parameter).GetUuid() {
			parameterss = append(parameterss, this.M_parameters[i])
		} else {
			isExist = true
			parameterss = append(parameterss, ref.(*Parameter))
		}
	}
	if !isExist {
		parameterss = append(parameterss, ref.(*Parameter))
		this.NeedSave = true
		this.M_parameters = parameterss
	}
}

/** Remove reference Parameters **/
func (this *Action) RemoveParameters(ref interface{}){
	toDelete := ref.(*Parameter)
	parameters_ := make([]*Parameter, 0)
	for i := 0; i < len(this.M_parameters); i++ {
		if toDelete.GetUuid() != this.M_parameters[i].GetUuid() {
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
	isExist := false
	var resultss []*Parameter
	for i:=0; i<len(this.M_results); i++ {
		if this.M_results[i].GetUuid() != ref.(*Parameter).GetUuid() {
			resultss = append(resultss, this.M_results[i])
		} else {
			isExist = true
			resultss = append(resultss, ref.(*Parameter))
		}
	}
	if !isExist {
		resultss = append(resultss, ref.(*Parameter))
		this.NeedSave = true
		this.M_results = resultss
	}
}

/** Remove reference Results **/
func (this *Action) RemoveResults(ref interface{}){
	toDelete := ref.(*Parameter)
	results_ := make([]*Parameter, 0)
	for i := 0; i < len(this.M_results); i++ {
		if toDelete.GetUuid() != this.M_results[i].GetUuid() {
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
	if this.M_accessType != ref.(AccessType) {
		this.M_accessType = ref.(AccessType)
		this.NeedSave = true
	}
}

/** Remove reference AccessType **/

/** Entities **/
func (this *Action) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Action) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_entitiesPtr != ref.(string) {
			this.M_entitiesPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_entitiesPtr != ref.(*Entities).GetUuid() {
			this.M_entitiesPtr = ref.(*Entities).GetUuid()
			this.NeedSave = true
		}
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Action) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
