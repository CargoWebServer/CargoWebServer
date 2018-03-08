// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type Entities struct{

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
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of Entities **/
	M_id string
	M_name string
	M_version string
	M_entities []string
	M_roles []string
	M_permissions []string
	M_actions []string

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
	M_permissions	[]*XsdPermission	`xml:"permissions,omitempty"`
	M_id	string	`xml:"id,attr"`
	M_name	string	`xml:"name,attr"`
	M_version	string	`xml:"version,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Entities) GetUuid() string{
	return this.UUID
}
func (this *Entities) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Entities) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Entities) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Entities"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Entities) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Entities) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Entities) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Entities) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Entities) GetChilds() []interface{}{
	var childs []interface{}
	var child interface{}
	var err error
	for i:=0; i < len(this.M_entities); i++ {
		child, err = this.getEntityByUuid( this.M_entities[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	for i:=0; i < len(this.M_roles); i++ {
		child, err = this.getEntityByUuid( this.M_roles[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	for i:=0; i < len(this.M_permissions); i++ {
		child, err = this.getEntityByUuid( this.M_permissions[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	for i:=0; i < len(this.M_actions); i++ {
		child, err = this.getEntityByUuid( this.M_actions[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	return childs
}
/** Evaluate if an entity needs to be saved. **/
func (this *Entities) IsNeedSave() bool{
	return this.NeedSave
}
func (this *Entities) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Entities) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Entities) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Entities) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Entities) GetId()string{
	return this.M_id
}

func (this *Entities) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *Entities) GetName()string{
	return this.M_name
}

func (this *Entities) SetName(val string){
	this.NeedSave = this.M_name== val
	this.M_name= val
}


func (this *Entities) GetVersion()string{
	return this.M_version
}

func (this *Entities) SetVersion(val string){
	this.NeedSave = this.M_version== val
	this.M_version= val
}


func (this *Entities) GetEntities()[]Entity{
	entities := make([]Entity, 0)
	for i := 0; i < len(this.M_entities); i++ {
		entity, err := this.getEntityByUuid(this.M_entities[i])
		if err == nil {
			entities = append(entities, entity.(Entity))
		}
	}
	return entities
}

func (this *Entities) SetEntities(val []Entity){
	this.M_entities= make([]string,0)
	for i:=0; i < len(val); i++{
		val[i].SetParentUuid(this.UUID)
		val[i].SetParentLnk("M_entities")
		if len(val[i].GetUuid()) == 0 {
			val[i].SetUuid(this.generateUuid(val[i]))
		}
		this.setEntity(val[i])

		this.M_entities=append(this.M_entities, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Entities) AppendEntities(val Entity){
	for i:=0; i < len(this.M_entities); i++{
		if this.M_entities[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	val.SetParentUuid(this.UUID)
	val.SetParentLnk("M_entities")
	if len(val.GetUuid()) == 0 {
		val.SetUuid(this.generateUuid(val))
	}
	this.setEntity(val)

	this.M_entities = append(this.M_entities, val.GetUuid())
}

func (this *Entities) RemoveEntities(val Entity){
	entities := make([]string,0)
	for i:=0; i < len(this.M_entities); i++{
		if this.M_entities[i] != val.GetUuid() {
			entities = append(entities, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_entities = entities
}


func (this *Entities) GetRoles()[]*Role{
	roles := make([]*Role, 0)
	for i := 0; i < len(this.M_roles); i++ {
		entity, err := this.getEntityByUuid(this.M_roles[i])
		if err == nil {
			roles = append(roles, entity.(*Role))
		}
	}
	return roles
}

func (this *Entities) SetRoles(val []*Role){
	this.M_roles= make([]string,0)
	for i:=0; i < len(val); i++{
		val[i].SetParentUuid(this.UUID)
		val[i].SetParentLnk("M_roles")
		if len(val[i].GetUuid()) == 0 {
			val[i].SetUuid(this.generateUuid(val[i]))
		}
		this.setEntity(val[i])

		this.M_roles=append(this.M_roles, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Entities) AppendRoles(val *Role){
	for i:=0; i < len(this.M_roles); i++{
		if this.M_roles[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	val.SetParentUuid(this.UUID)
	val.SetParentLnk("M_roles")
	if len(val.GetUuid()) == 0 {
		val.SetUuid(this.generateUuid(val))
	}
	this.setEntity(val)

	this.M_roles = append(this.M_roles, val.GetUuid())
}

func (this *Entities) RemoveRoles(val *Role){
	roles := make([]string,0)
	for i:=0; i < len(this.M_roles); i++{
		if this.M_roles[i] != val.GetUuid() {
			roles = append(roles, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_roles = roles
}


func (this *Entities) GetPermissions()[]*Permission{
	permissions := make([]*Permission, 0)
	for i := 0; i < len(this.M_permissions); i++ {
		entity, err := this.getEntityByUuid(this.M_permissions[i])
		if err == nil {
			permissions = append(permissions, entity.(*Permission))
		}
	}
	return permissions
}

func (this *Entities) SetPermissions(val []*Permission){
	this.M_permissions= make([]string,0)
	for i:=0; i < len(val); i++{
		val[i].SetParentUuid(this.UUID)
		val[i].SetParentLnk("M_permissions")
		if len(val[i].GetUuid()) == 0 {
			val[i].SetUuid(this.generateUuid(val[i]))
		}
		this.setEntity(val[i])

		this.M_permissions=append(this.M_permissions, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Entities) AppendPermissions(val *Permission){
	for i:=0; i < len(this.M_permissions); i++{
		if this.M_permissions[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	val.SetParentUuid(this.UUID)
	val.SetParentLnk("M_permissions")
	if len(val.GetUuid()) == 0 {
		val.SetUuid(this.generateUuid(val))
	}
	this.setEntity(val)

	this.M_permissions = append(this.M_permissions, val.GetUuid())
}

func (this *Entities) RemovePermissions(val *Permission){
	permissions := make([]string,0)
	for i:=0; i < len(this.M_permissions); i++{
		if this.M_permissions[i] != val.GetUuid() {
			permissions = append(permissions, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_permissions = permissions
}


func (this *Entities) GetActions()[]*Action{
	actions := make([]*Action, 0)
	for i := 0; i < len(this.M_actions); i++ {
		entity, err := this.getEntityByUuid(this.M_actions[i])
		if err == nil {
			actions = append(actions, entity.(*Action))
		}
	}
	return actions
}

func (this *Entities) SetActions(val []*Action){
	this.M_actions= make([]string,0)
	for i:=0; i < len(val); i++{
		val[i].SetParentUuid(this.UUID)
		val[i].SetParentLnk("M_actions")
		if len(val[i].GetUuid()) == 0 {
			val[i].SetUuid(this.generateUuid(val[i]))
		}
		this.setEntity(val[i])

		this.M_actions=append(this.M_actions, val[i].GetUuid())
	}
	this.NeedSave= true
}

func (this *Entities) AppendActions(val *Action){
	for i:=0; i < len(this.M_actions); i++{
		if this.M_actions[i] == val.GetUuid() {
			return
		}
	}
	this.NeedSave= true
	val.SetParentUuid(this.UUID)
	val.SetParentLnk("M_actions")
	if len(val.GetUuid()) == 0 {
		val.SetUuid(this.generateUuid(val))
	}
	this.setEntity(val)

	this.M_actions = append(this.M_actions, val.GetUuid())
}

func (this *Entities) RemoveActions(val *Action){
	actions := make([]string,0)
	for i:=0; i < len(this.M_actions); i++{
		if this.M_actions[i] != val.GetUuid() {
			actions = append(actions, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_actions = actions
}

