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

	/** If the entity is fully initialyse **/
	IsInit   bool

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
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Name **/
func (this *Entities) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Entities) SetName(ref interface{}){
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Version **/
func (this *Entities) GetVersion() string{
	return this.M_version
}

/** Init reference Version **/
func (this *Entities) SetVersion(ref interface{}){
	this.M_version = ref.(string)
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
	toDelete := ref.(*Role)
	roles_ := make([]*Role, 0)
	for i := 0; i < len(this.M_roles); i++ {
		if toDelete.GetUUID() != this.M_roles[i].GetUUID() {
			roles_ = append(roles_, this.M_roles[i])
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
		if this.M_permissions[i].GetUUID() != ref.(*Permission).GetUUID() {
			permissionss = append(permissionss, this.M_permissions[i])
		} else {
			isExist = true
			permissionss = append(permissionss, ref.(*Permission))
		}
	}
	if !isExist {
		permissionss = append(permissionss, ref.(*Permission))
	}
	this.M_permissions = permissionss
}

/** Remove reference Permissions **/
func (this *Entities) RemovePermissions(ref interface{}){
	toDelete := ref.(*Permission)
	permissions_ := make([]*Permission, 0)
	for i := 0; i < len(this.M_permissions); i++ {
		if toDelete.GetUUID() != this.M_permissions[i].GetUUID() {
			permissions_ = append(permissions_, this.M_permissions[i])
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
		if this.M_actions[i].GetUUID() != ref.(*Action).GetUUID() {
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
func (this *Entities) RemoveActions(ref interface{}){
	toDelete := ref.(*Action)
	actions_ := make([]*Action, 0)
	for i := 0; i < len(this.M_actions); i++ {
		if toDelete.GetUUID() != this.M_actions[i].GetUUID() {
			actions_ = append(actions_, this.M_actions[i])
		}
	}
	this.M_actions = actions_
}
