// +build CargoEntities

package CargoEntities

import(
"encoding/xml"
)

type Account struct{

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

	/** members of Entity **/
	M_id string

	/** members of Account **/
	M_name string
	M_password string
	M_email string
	M_sessions []*Session
	M_permissions []*Permission
	M_messages []Message
	m_userRef *User
	/** If the ref is a string and not an object **/
	M_userRef string
	m_rolesRef []*Role
	/** If the ref is a string and not an object **/
	M_rolesRef []string


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for Account **/
type XsdAccount struct {
	XMLName xml.Name	`xml:"toRef"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_userRef	*string	`xml:"userRef"`
	M_rolesRef	[]string	`xml:"rolesRef"`
	M_permissions	[]*XsdPermission	`xml:"permissions,omitempty"`
	M_sessions	[]*XsdSession	`xml:"sessions,omitempty"`

	M_name	string	`xml:"name,attr"`
	M_password	string	`xml:"password,attr"`
	M_email	string	`xml:"email,attr"`

}
/** UUID **/
func (this *Account) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Account) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Account) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Name **/
func (this *Account) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Account) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Password **/
func (this *Account) GetPassword() string{
	return this.M_password
}

/** Init reference Password **/
func (this *Account) SetPassword(ref interface{}){
	this.NeedSave = true
	this.M_password = ref.(string)
}

/** Remove reference Password **/

/** Email **/
func (this *Account) GetEmail() string{
	return this.M_email
}

/** Init reference Email **/
func (this *Account) SetEmail(ref interface{}){
	this.NeedSave = true
	this.M_email = ref.(string)
}

/** Remove reference Email **/

/** Sessions **/
func (this *Account) GetSessions() []*Session{
	return this.M_sessions
}

/** Init reference Sessions **/
func (this *Account) SetSessions(ref interface{}){
	this.NeedSave = true
	isExist := false
	var sessionss []*Session
	for i:=0; i<len(this.M_sessions); i++ {
		if this.M_sessions[i].GetUUID() != ref.(*Session).GetUUID() {
			sessionss = append(sessionss, this.M_sessions[i])
		} else {
			isExist = true
			sessionss = append(sessionss, ref.(*Session))
		}
	}
	if !isExist {
		sessionss = append(sessionss, ref.(*Session))
	}
	this.M_sessions = sessionss
}

/** Remove reference Sessions **/
func (this *Account) RemoveSessions(ref interface{}){
	toDelete := ref.(*Session)
	sessions_ := make([]*Session, 0)
	for i := 0; i < len(this.M_sessions); i++ {
		if toDelete.GetUUID() != this.M_sessions[i].GetUUID() {
			sessions_ = append(sessions_, this.M_sessions[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_sessions = sessions_
}

/** Permissions **/
func (this *Account) GetPermissions() []*Permission{
	return this.M_permissions
}

/** Init reference Permissions **/
func (this *Account) SetPermissions(ref interface{}){
	this.NeedSave = true
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
func (this *Account) RemovePermissions(ref interface{}){
	toDelete := ref.(*Permission)
	permissions_ := make([]*Permission, 0)
	for i := 0; i < len(this.M_permissions); i++ {
		if toDelete.GetUUID() != this.M_permissions[i].GetUUID() {
			permissions_ = append(permissions_, this.M_permissions[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_permissions = permissions_
}

/** Messages **/
func (this *Account) GetMessages() []Message{
	return this.M_messages
}

/** Init reference Messages **/
func (this *Account) SetMessages(ref interface{}){
	this.NeedSave = true
	isExist := false
	var messagess []Message
	for i:=0; i<len(this.M_messages); i++ {
		if this.M_messages[i].(Entity).GetUUID() != ref.(Entity).GetUUID() {
			messagess = append(messagess, this.M_messages[i])
		} else {
			isExist = true
			messagess = append(messagess, ref.(Message))
		}
	}
	if !isExist {
		messagess = append(messagess, ref.(Message))
	}
	this.M_messages = messagess
}

/** Remove reference Messages **/
func (this *Account) RemoveMessages(ref interface{}){
	toDelete := ref.(Entity)
	messages_ := make([]Message, 0)
	for i := 0; i < len(this.M_messages); i++ {
		if toDelete.GetUUID() != this.M_messages[i].(Entity).GetUUID() {
			messages_ = append(messages_, this.M_messages[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_messages = messages_
}

/** UserRef **/
func (this *Account) GetUserRef() *User{
	return this.m_userRef
}

/** Init reference UserRef **/
func (this *Account) SetUserRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_userRef = ref.(string)
	}else{
		this.m_userRef = ref.(*User)
		this.M_userRef = ref.(Entity).GetUUID()
	}
}

/** Remove reference UserRef **/
func (this *Account) RemoveUserRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_userRef!= nil {
		if toDelete.GetUUID() == this.m_userRef.GetUUID() {
			this.m_userRef = nil
			this.M_userRef = ""
		}else{
			this.NeedSave = true
		}
	}
}

/** RolesRef **/
func (this *Account) GetRolesRef() []*Role{
	return this.m_rolesRef
}

/** Init reference RolesRef **/
func (this *Account) SetRolesRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_rolesRef); i++ {
			if this.M_rolesRef[i] == refStr {
				return
			}
		}
		this.M_rolesRef = append(this.M_rolesRef, ref.(string))
	}else{
		this.RemoveRolesRef(ref)
		this.m_rolesRef = append(this.m_rolesRef, ref.(*Role))
		this.M_rolesRef = append(this.M_rolesRef, ref.(*Role).GetUUID())
	}
}

/** Remove reference RolesRef **/
func (this *Account) RemoveRolesRef(ref interface{}){
	toDelete := ref.(*Role)
	rolesRef_ := make([]*Role, 0)
	rolesRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_rolesRef); i++ {
		if toDelete.GetUUID() != this.m_rolesRef[i].GetUUID() {
			rolesRef_ = append(rolesRef_, this.m_rolesRef[i])
			rolesRefUuid = append(rolesRefUuid, this.M_rolesRef[i])
		}else{
			this.NeedSave = true
		}
	}
	this.m_rolesRef = rolesRef_
	this.M_rolesRef = rolesRefUuid
}

/** Entities **/
func (this *Account) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Account) SetEntitiesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *Account) RemoveEntitiesPtr(ref interface{}){
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
