package CargoEntities

import(
"encoding/xml"
)

type Entities struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Entities **/
	M_id string
	M_name string
	M_version string
	M_entities []Entity
	M_roles []*Role
	M_actions []*Action

}

/** Xml parser for Entities **/
type XsdEntities struct {
	XMLName xml.Name	`xml:"entities"`
	M_entities_0	[]*XsdAccount	`xml:"toRef,omitempty"`
	M_entities_1	[]*XsdFile	`xml:"filesRef,omitempty"`
	M_entities_2	[]*XsdProject	`xml:"project,omitempty"`
	M_entities_3	[]*XsdComputer	`xml:"computerRef,omitempty"`
	M_entities_4	[]*XsdUser	`xml:"membersRef,omitempty"`
	M_entities_5	[]*XsdGroup	`xml:"memberOfRef,omitempty"`
	M_entities_6	[]*XsdLogEntry	`xml:"logEntry,omitempty"`
	M_entities_7	[]*XsdLog	`xml:"log,omitempty"`

	M_roles	[]*XsdRole	`xml:"roles,omitempty"`
	M_actions	[]*XsdAction	`xml:"actions,omitempty"`
	M_id	string	`xml:"id,attr"`
	M_name	string	`xml:"name,attr"`
	M_version	string	`xml:"version,attr"`

}
/** UUID **/
func (this *Entities) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Entities) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Entities) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Name **/
func (this *Entities) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Entities) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Version **/
func (this *Entities) GetVersion() string{
	return this.M_version
}

/** Init reference Version **/
func (this *Entities) SetVersion(ref interface{}){
	this.NeedSave = true
	this.M_version = ref.(string)
}

/** Remove reference Version **/

/** Entities **/
func (this *Entities) GetEntities() []Entity{
	return this.M_entities
}

/** Init reference Entities **/
func (this *Entities) SetEntities(ref interface{}){
	this.NeedSave = true
	isExist := false
	var entitiess []Entity
	for i:=0; i<len(this.M_entities); i++ {
		if this.M_entities[i].(Entity).GetUUID() != ref.(Entity).GetUUID() {
			entitiess = append(entitiess, this.M_entities[i])
		} else {
			isExist = true
			entitiess = append(entitiess, ref.(Entity))
		}
	}
	if !isExist {
		entitiess = append(entitiess, ref.(Entity))
	}
	this.M_entities = entitiess
}

/** Remove reference Entities **/
func (this *Entities) RemoveEntities(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Entity)
	entities_ := make([]Entity, 0)
	for i := 0; i < len(this.M_entities); i++ {
		if toDelete.GetUUID() != this.M_entities[i].(Entity).GetUUID() {
			entities_ = append(entities_, this.M_entities[i])
		}
	}
	this.M_entities = entities_
}

/** Roles **/
func (this *Entities) GetRoles() []*Role{
	return this.M_roles
}

/** Init reference Roles **/
func (this *Entities) SetRoles(ref interface{}){
	this.NeedSave = true
	isExist := false
	var roless []*Role
	for i:=0; i<len(this.M_roles); i++ {
		if this.M_roles[i].GetUUID() != ref.(*Role).GetUUID() {
			roless = append(roless, this.M_roles[i])
		} else {
			isExist = true
			roless = append(roless, ref.(*Role))
		}
	}
	if !isExist {
		roless = append(roless, ref.(*Role))
	}
	this.M_roles = roless
}

/** Remove reference Roles **/
func (this *Entities) RemoveRoles(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Role)
	roles_ := make([]*Role, 0)
	for i := 0; i < len(this.M_roles); i++ {
		if toDelete.GetUUID() != this.M_roles[i].GetUUID() {
			roles_ = append(roles_, this.M_roles[i])
		}
	}
	this.M_roles = roles_
}

/** Actions **/
func (this *Entities) GetActions() []*Action{
	return this.M_actions
}

/** Init reference Actions **/
func (this *Entities) SetActions(ref interface{}){
	this.NeedSave = true
	isExist := false
	var actionss []*Action
	for i:=0; i<len(this.M_actions); i++ {
		if this.M_actions[i].GetName() != ref.(*Action).GetName() {
			actionss = append(actionss, this.M_actions[i])
		} else {
			isExist = true
			actionss = append(actionss, ref.(*Action))
		}
	}
	if !isExist {
		actionss = append(actionss, ref.(*Action))
	}
	this.M_actions = actionss
}

/** Remove reference Actions **/
