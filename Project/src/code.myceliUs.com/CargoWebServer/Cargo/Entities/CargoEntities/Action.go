// +build CargoEntities

package CargoEntities

import (
	"encoding/xml"
)

type Action struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** Get entity by uuid function **/
	getEntityByUuid func(string) (interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of Action **/
	M_name       string
	M_doc        string
	M_parameters []string
	M_results    []string
	M_accessType AccessType

	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for Action **/
type XsdAction struct {
	XMLName      xml.Name        `xml:"action"`
	M_parameters []*XsdParameter `xml:"parameters,omitempty"`
	M_results    []*XsdParameter `xml:"results,omitempty"`
	M_name       string          `xml:"name,attr"`
	M_doc        string          `xml:"doc,attr"`
	M_accessType string          `xml:"accessType,attr"`
}

/***************** Entity **************************/

/** UUID **/
func (this *Action) GetUuid() string {
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Action) SetUuid(uuid string) {
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Action) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_name)
	return ids
}

/** The type name **/
func (this *Action) GetTypeName() string {
	this.TYPENAME = "CargoEntities.Action"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Action) GetParentUuid() string {
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Action) SetParentUuid(parentUuid string) {
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Action) GetParentLnk() string {
	return this.ParentLnk
}
func (this *Action) SetParentLnk(parentLnk string) {
	this.ParentLnk = parentLnk
}

func (this *Action) GetParent() interface{} {
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Action) GetChilds() []interface{} {
	var childs []interface{}
	var child interface{}
	var err error
	for i := 0; i < len(this.M_parameters); i++ {
		child, err = this.getEntityByUuid(this.M_parameters[i])
		if err == nil {
			childs = append(childs, child)
		}
	}
	for i := 0; i < len(this.M_results); i++ {
		child, err = this.getEntityByUuid(this.M_results[i])
		if err == nil {
			childs = append(childs, child)
		}
	}
	return childs
}

/** Return the list of all childs uuid **/
func (this *Action) GetChildsUuid() []string {
	var childs []string
	childs = append(childs, this.M_parameters...)
	childs = append(childs, this.M_results...)
	return childs
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Action) SetEntityGetter(fct func(uuid string) (interface{}, error)) {
	this.getEntityByUuid = fct
}

/** Use it the set the entity on the cache. **/
func (this *Action) SetEntitySetter(fct func(entity interface{})) {
	this.setEntity = fct
}

/** Set the uuid generator function **/
func (this *Action) SetUuidGenerator(fct func(entity interface{}) string) {
	this.generateUuid = fct
}

func (this *Action) GetName() string {
	return this.M_name
}

func (this *Action) SetName(val string) {
	this.M_name = val
}

func (this *Action) GetDoc() string {
	return this.M_doc
}

func (this *Action) SetDoc(val string) {
	this.M_doc = val
}

func (this *Action) GetParameters() []*Parameter {
	values := make([]*Parameter, 0)
	for i := 0; i < len(this.M_parameters); i++ {
		entity, err := this.getEntityByUuid(this.M_parameters[i])
		if err == nil {
			values = append(values, entity.(*Parameter))
		}
	}
	return values
}

func (this *Action) SetParameters(val []*Parameter) {
	this.M_parameters = make([]string, 0)
	for i := 0; i < len(val); i++ {
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_parameters")
		this.M_parameters = append(this.M_parameters, val[i].GetUuid())
		this.setEntity(val[i])
	}
	this.setEntity(this)
}

func (this *Action) AppendParameters(val *Parameter) {
	for i := 0; i < len(this.M_parameters); i++ {
		if this.M_parameters[i] == val.GetUuid() {
			return
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_parameters")
	this.setEntity(val)
	this.M_parameters = append(this.M_parameters, val.GetUuid())
	this.setEntity(this)
}

func (this *Action) RemoveParameters(val *Parameter) {
	values := make([]string, 0)
	for i := 0; i < len(this.M_parameters); i++ {
		if this.M_parameters[i] != val.GetUuid() {
			values = append(values, val.GetUuid())
		}
	}
	this.M_parameters = values
	this.setEntity(this)
}

func (this *Action) GetResults() []*Parameter {
	values := make([]*Parameter, 0)
	for i := 0; i < len(this.M_results); i++ {
		entity, err := this.getEntityByUuid(this.M_results[i])
		if err == nil {
			values = append(values, entity.(*Parameter))
		}
	}
	return values
}

func (this *Action) SetResults(val []*Parameter) {
	this.M_results = make([]string, 0)
	for i := 0; i < len(val); i++ {
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_results")
		this.M_results = append(this.M_results, val[i].GetUuid())
		this.setEntity(val[i])
	}
	this.setEntity(this)
}

func (this *Action) AppendResults(val *Parameter) {
	for i := 0; i < len(this.M_results); i++ {
		if this.M_results[i] == val.GetUuid() {
			return
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_results")
	this.setEntity(val)
	this.M_results = append(this.M_results, val.GetUuid())
	this.setEntity(this)
}

func (this *Action) RemoveResults(val *Parameter) {
	values := make([]string, 0)
	for i := 0; i < len(this.M_results); i++ {
		if this.M_results[i] != val.GetUuid() {
			values = append(values, val.GetUuid())
		}
	}
	this.M_results = values
	this.setEntity(this)
}

func (this *Action) GetAccessType() AccessType {
	return this.M_accessType
}

func (this *Action) SetAccessType(val AccessType) {
	this.M_accessType = val
}

func (this *Action) ResetAccessType() {
	this.M_accessType = 0
}

func (this *Action) GetEntitiesPtr() *Entities {
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Action) SetEntitiesPtr(val *Entities) {
	this.M_entitiesPtr = val.GetUuid()
	this.setEntity(this)
}

func (this *Action) ResetEntitiesPtr() {
	this.M_entitiesPtr = ""
}
