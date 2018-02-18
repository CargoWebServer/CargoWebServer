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

	/** members of Entities **/
	M_id string
	M_name string
	M_version string
	M_entities []Entity
	M_roles []*Role
	M_permissions []*Permission
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

/** Evaluate if an entity needs to be saved. **/
func (this *Entities) IsNeedSave() bool{
	return this.NeedSave
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Entities) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *Entities) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Entities) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Name **/
func (this *Entities) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Entities) SetName(ref interface{}){
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Name **/

/** Version **/
func (this *Entities) GetVersion() string{
	return this.M_version
}

/** Init reference Version **/
func (this *Entities) SetVersion(ref interface{}){
	if this.M_version != ref.(string) {
		this.M_version = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Version **/

/** Entities **/
func (this *Entities) GetEntities() []Entity{
	return this.M_entities
}

/** Init reference Entities **/
func (this *Entities) SetEntities(ref interface{}){
	isExist := false
	var entitiess []Entity
	for i:=0; i<len(this.M_entities); i++ {
		if this.M_entities[i].(Entity).GetUuid() != ref.(Entity).GetUuid() {
			entitiess = append(entitiess, this.M_entities[i])
		} else {
			isExist = true
			entitiess = append(entitiess, ref.(Entity))
		}
	}
	if !isExist {
		entitiess = append(entitiess, ref.(Entity))
		this.NeedSave = true
		this.M_entities = entitiess
	}
}

/** Remove reference Entities **/
func (this *Entities) RemoveEntities(ref interface{}){
	toDelete := ref.(Entity)
	entities_ := make([]Entity, 0)
	for i := 0; i < len(this.M_entities); i++ {
		if toDelete.GetUuid() != this.M_entities[i].(Entity).GetUuid() {
			entities_ = append(entities_, this.M_entities[i])
		}else{
			this.NeedSave = true
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
	isExist := false
	var roless []*Role
	for i:=0; i<len(this.M_roles); i++ {
		if this.M_roles[i].GetUuid() != ref.(*Role).GetUuid() {
			roless = append(roless, this.M_roles[i])
		} else {
			isExist = true
			roless = append(roless, ref.(*Role))
		}
	}
	if !isExist {
		roless = append(roless, ref.(*Role))
		this.NeedSave = true
		this.M_roles = roless
	}
}

/** Remove reference Roles **/
func (this *Entities) RemoveRoles(ref interface{}){
	toDelete := ref.(*Role)
	roles_ := make([]*Role, 0)
	for i := 0; i < len(this.M_roles); i++ {
		if toDelete.GetUuid() != this.M_roles[i].GetUuid() {
			roles_ = append(roles_, this.M_roles[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_roles = roles_
}

/** Permissions **/
func (this *Entities) GetPermissions() []*Permission{
	return this.M_permissions
}

/** Init reference Permissions **/
func (this *Entities) SetPermissions(ref interface{}){
	isExist := false
	var permissionss []*Permission
	for i:=0; i<len(this.M_permissions); i++ {
		if this.M_permissions[i].GetUuid() != ref.(*Permission).GetUuid() {
			permissionss = append(permissionss, this.M_permissions[i])
		} else {
			isExist = true
			permissionss = append(permissionss, ref.(*Permission))
		}
	}
	if !isExist {
		permissionss = append(permissionss, ref.(*Permission))
		this.NeedSave = true
		this.M_permissions = permissionss
	}
}

/** Remove reference Permissions **/
func (this *Entities) RemovePermissions(ref interface{}){
	toDelete := ref.(*Permission)
	permissions_ := make([]*Permission, 0)
	for i := 0; i < len(this.M_permissions); i++ {
		if toDelete.GetUuid() != this.M_permissions[i].GetUuid() {
			permissions_ = append(permissions_, this.M_permissions[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_permissions = permissions_
}

/** Actions **/
func (this *Entities) GetActions() []*Action{
	return this.M_actions
}

/** Init reference Actions **/
func (this *Entities) SetActions(ref interface{}){
	isExist := false
	var actionss []*Action
	for i:=0; i<len(this.M_actions); i++ {
		if this.M_actions[i].GetUuid() != ref.(*Action).GetUuid() {
			actionss = append(actionss, this.M_actions[i])
		} else {
			isExist = true
			actionss = append(actionss, ref.(*Action))
		}
	}
	if !isExist {
		actionss = append(actionss, ref.(*Action))
		this.NeedSave = true
		this.M_actions = actionss
	}
}

/** Remove reference Actions **/
func (this *Entities) RemoveActions(ref interface{}){
	toDelete := ref.(*Action)
	actions_ := make([]*Action, 0)
	for i := 0; i < len(this.M_actions); i++ {
		if toDelete.GetUuid() != this.M_actions[i].GetUuid() {
			actions_ = append(actions_, this.M_actions[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_actions = actions_
}
