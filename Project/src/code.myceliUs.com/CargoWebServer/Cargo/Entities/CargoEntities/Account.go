// +build CargoEntities

package CargoEntities

import (
	"encoding/xml"
)

type Account struct {

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
	getEntityByUuid func(string) (interface{}, error)

	/** members of Entity **/
	M_id string

	/** members of Account **/
	M_name           string
	M_password       string
	M_email          string
	M_sessions       []string
	M_messages       []string
	M_userRef        string
	M_rolesRef       []string
	M_permissionsRef []string

	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for Account **/
type XsdAccount struct {
	XMLName xml.Name `xml:"toRef"`
	/** Entity **/
	M_id string `xml:"id,attr"`

	M_userRef        *string       `xml:"userRef"`
	M_rolesRef       []string      `xml:"rolesRef"`
	M_permissionsRef []string      `xml:"permissionsRef"`
	M_sessions       []*XsdSession `xml:"sessions,omitempty"`

	M_name     string `xml:"name,attr"`
	M_password string `xml:"password,attr"`
	M_email    string `xml:"email,attr"`
}

/***************** Entity **************************/

/** UUID **/
func (this *Account) GetUuid() string {
	return this.UUID
}
func (this *Account) SetUuid(uuid string) {
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Account) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Account) GetTypeName() string {
	this.TYPENAME = "CargoEntities.Account"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Account) GetParentUuid() string {
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Account) SetParentUuid(parentUuid string) {
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Account) GetParentLnk() string {
	return this.ParentLnk
}
func (this *Account) SetParentLnk(parentLnk string) {
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *Account) IsNeedSave() bool {
	return this.NeedSave
}
func (this *Account) ResetNeedSave() {
	this.NeedSave = false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Account) SetEntityGetter(fct func(uuid string) (interface{}, error)) {
	this.getEntityByUuid = fct
}

func (this *Account) GetId() string {
	return this.M_id
}

func (this *Account) SetId(val string) {
	this.NeedSave = this.M_id == val
	this.M_id = val
}

func (this *Account) GetName() string {
	return this.M_name
}

func (this *Account) SetName(val string) {
	this.NeedSave = this.M_name == val
	this.M_name = val
}

func (this *Account) GetPassword() string {
	return this.M_password
}

func (this *Account) SetPassword(val string) {
	this.NeedSave = this.M_password == val
	this.M_password = val
}

func (this *Account) GetEmail() string {
	return this.M_email
}

func (this *Account) SetEmail(val string) {
	this.NeedSave = this.M_email == val
	this.M_email = val
}

func (this *Account) GetSessions() []*Session {
	sessions := make([]*Session, 0)
	for i := 0; i < len(this.M_sessions); i++ {
		entity, err := this.getEntityByUuid(this.M_sessions[i])
		if err == nil {
			sessions = append(sessions, entity.(*Session))
		}
	}
	return sessions
}

func (this *Account) SetSessions(val []*Session) {
	this.M_sessions = make([]string, 0)
	for i := 0; i < len(val); i++ {
		this.M_sessions = append(this.M_sessions, val[i].GetUuid())
	}
}

func (this *Account) AppendSessions(val *Session) {
	for i := 0; i < len(this.M_sessions); i++ {
		if this.M_sessions[i] == val.GetUuid() {
			return
		}
	}
	this.M_sessions = append(this.M_sessions, val.GetUuid())
}

func (this *Account) RemoveSessions(val *Session) {
	sessions := make([]string, 0)
	for i := 0; i < len(this.M_sessions); i++ {
		if this.M_sessions[i] != val.GetUuid() {
			sessions = append(sessions, val.GetUuid())
		} else {
			this.NeedSave = true
		}
	}
	this.M_sessions = sessions
}

func (this *Account) GetMessages() []Message {
	messages := make([]Message, 0)
	for i := 0; i < len(this.M_messages); i++ {
		entity, err := this.getEntityByUuid(this.M_messages[i])
		if err == nil {
			messages = append(messages, entity.(Message))
		}
	}
	return messages
}

func (this *Account) SetMessages(val []Message) {
	this.M_messages = make([]string, 0)
	for i := 0; i < len(val); i++ {
		this.M_messages = append(this.M_messages, val[i].GetUuid())
	}
}

func (this *Account) AppendMessages(val Message) {
	for i := 0; i < len(this.M_messages); i++ {
		if this.M_messages[i] == val.GetUuid() {
			return
		}
	}
	this.M_messages = append(this.M_messages, val.GetUuid())
}

func (this *Account) RemoveMessages(val Message) {
	messages := make([]string, 0)
	for i := 0; i < len(this.M_messages); i++ {
		if this.M_messages[i] != val.GetUuid() {
			messages = append(messages, val.GetUuid())
		} else {
			this.NeedSave = true
		}
	}
	this.M_messages = messages
}

func (this *Account) GetUserRef() *User {
	entity, err := this.getEntityByUuid(this.M_userRef)
	if err == nil {
		return entity.(*User)
	}
	return nil
}

func (this *Account) SetUserRef(val *User) {
	this.M_userRef = val.GetUuid()
}

func (this *Account) ResetUserRef() {
	this.M_userRef = ""
}

func (this *Account) GetRolesRef() []*Role {
	rolesRef := make([]*Role, 0)
	for i := 0; i < len(this.M_rolesRef); i++ {
		entity, err := this.getEntityByUuid(this.M_rolesRef[i])
		if err == nil {
			rolesRef = append(rolesRef, entity.(*Role))
		}
	}
	return rolesRef
}

func (this *Account) SetRolesRef(val []*Role) {
	this.M_rolesRef = make([]string, 0)
	for i := 0; i < len(val); i++ {
		this.M_rolesRef = append(this.M_rolesRef, val[i].GetUuid())
	}
}

func (this *Account) AppendRolesRef(val *Role) {
	for i := 0; i < len(this.M_rolesRef); i++ {
		if this.M_rolesRef[i] == val.GetUuid() {
			return
		}
	}
	this.M_rolesRef = append(this.M_rolesRef, val.GetUuid())
}

func (this *Account) RemoveRolesRef(val *Role) {
	rolesRef := make([]string, 0)
	for i := 0; i < len(this.M_rolesRef); i++ {
		if this.M_rolesRef[i] != val.GetUuid() {
			rolesRef = append(rolesRef, val.GetUuid())
		} else {
			this.NeedSave = true
		}
	}
	this.M_rolesRef = rolesRef
}

func (this *Account) GetPermissionsRef() []*Permission {
	permissionsRef := make([]*Permission, 0)
	for i := 0; i < len(this.M_permissionsRef); i++ {
		entity, err := this.getEntityByUuid(this.M_permissionsRef[i])
		if err == nil {
			permissionsRef = append(permissionsRef, entity.(*Permission))
		}
	}
	return permissionsRef
}

func (this *Account) SetPermissionsRef(val []*Permission) {
	this.M_permissionsRef = make([]string, 0)
	for i := 0; i < len(val); i++ {
		this.M_permissionsRef = append(this.M_permissionsRef, val[i].GetUuid())
	}
}

func (this *Account) AppendPermissionsRef(val *Permission) {
	for i := 0; i < len(this.M_permissionsRef); i++ {
		if this.M_permissionsRef[i] == val.GetUuid() {
			return
		}
	}
	this.M_permissionsRef = append(this.M_permissionsRef, val.GetUuid())
}

func (this *Account) RemovePermissionsRef(val *Permission) {
	permissionsRef := make([]string, 0)
	for i := 0; i < len(this.M_permissionsRef); i++ {
		if this.M_permissionsRef[i] != val.GetUuid() {
			permissionsRef = append(permissionsRef, val.GetUuid())
		} else {
			this.NeedSave = true
		}
	}
	this.M_permissionsRef = permissionsRef
}

func (this *Account) GetEntitiesPtr() *Entities {
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Account) SetEntitiesPtr(val *Entities) {
	this.M_entitiesPtr = val.GetUuid()
}

func (this *Account) ResetEntitiesPtr() {
	this.M_entitiesPtr = ""
}
