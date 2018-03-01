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
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

	/** members of Action **/
	M_name string
	M_doc string
	M_parameters []string
	M_results []string
	M_accessType AccessType


	/** Associations **/
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
func (this *Action) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Action) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

func (this *Action) GetName()string{
	return this.M_name
}

func (this *Action) SetName(val string){
	this.NeedSave = this.M_name== val
	this.M_name= val
}


func (this *Action) GetDoc()string{
	return this.M_doc
}

func (this *Action) SetDoc(val string){
	this.NeedSave = this.M_doc== val
	this.M_doc= val
}


func (this *Action) GetParameters()[]*Parameter{
	parameters := make([]*Parameter, 0)
	for i := 0; i < len(this.M_parameters); i++ {
		entity, err := this.getEntityByUuid(this.M_parameters[i])
		if err == nil {
			parameters = append(parameters, entity.(*Parameter))
		}
	}
	return parameters
}

func (this *Action) SetParameters(val []*Parameter){
	this.M_parameters= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_parameters=append(this.M_parameters, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Action) AppendParameters(val *Parameter){
	for i:=0; i < len(this.M_parameters); i++{
		if this.M_parameters[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_parameters = append(this.M_parameters, val.GetUuid())
}

func (this *Action) RemoveParameters(val *Parameter){
	parameters := make([]string,0)
	for i:=0; i < len(this.M_parameters); i++{
		if this.M_parameters[i] != val.GetUuid() {
			parameters = append(parameters, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_parameters = parameters
}


func (this *Action) GetResults()[]*Parameter{
	results := make([]*Parameter, 0)
	for i := 0; i < len(this.M_results); i++ {
		entity, err := this.getEntityByUuid(this.M_results[i])
		if err == nil {
			results = append(results, entity.(*Parameter))
		}
	}
	return results
}

func (this *Action) SetResults(val []*Parameter){
	this.M_results= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_results=append(this.M_results, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Action) AppendResults(val *Parameter){
	for i:=0; i < len(this.M_results); i++{
		if this.M_results[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	this.M_results = append(this.M_results, val.GetUuid())
}

func (this *Action) RemoveResults(val *Parameter){
	results := make([]string,0)
	for i:=0; i < len(this.M_results); i++{
		if this.M_results[i] != val.GetUuid() {
			results = append(results, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_results = results
}


func (this *Action) GetAccessType()AccessType{
	return this.M_accessType
}

func (this *Action) SetAccessType(val AccessType){
	this.NeedSave = this.M_accessType== val
	this.M_accessType= val
}

func (this *Action) ResetAccessType(){
	this.M_accessType= 0
}


func (this *Action) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Action) SetEntitiesPtr(val *Entities){
	this.NeedSave = this.M_entitiesPtr != val.GetUuid()
	this.M_entitiesPtr= val.GetUuid()
}

func (this *Action) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

