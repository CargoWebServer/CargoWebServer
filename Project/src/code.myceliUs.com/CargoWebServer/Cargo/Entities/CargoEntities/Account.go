// +build CargoEntities

package CargoEntities

import (
	"encoding/xml"
	//	"log"
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

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of Entity **/
	M_id string

	/** members of Account **/
	M_name     string
	M_password string
	M_email    string
	M_sessions []*Session
	M_messages []Message
	m_userRef  *User
	/** If the ref is a string and not an object **/
	M_userRef  string
	m_rolesRef []*Role
	/** If the ref is a string and not an object **/
	M_rolesRef       []string
	m_permissionsRef []*Permission
	/** If the ref is a string and not an object **/
	M_permissionsRef []string

	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
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

/** UUID **/
func (this *Account) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *Account) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *Account) SetId(ref interface{}) {
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference Id **/

/** Name **/
func (this *Account) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Account) SetName(ref interface{}) {
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference Name **/

/** Password **/
func (this *Account) GetPassword() string {
	return this.M_password
}

/** Init reference Password **/
func (this *Account) SetPassword(ref interface{}) {
	if this.M_password != ref.(string) {
		this.M_password = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference Password **/

/** Email **/
func (this *Account) GetEmail() string {
	return this.M_email
}

/** Init reference Email **/
func (this *Account) SetEmail(ref interface{}) {
	if this.M_email != ref.(string) {
		this.M_email = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference Email **/

/** Sessions **/
func (this *Account) GetSessions() []*Session {
	return this.M_sessions
}

/** Init reference Sessions **/
func (this *Account) SetSessions(ref interface{}) {
	isExist := false
	var sessionss []*Session
	for i := 0; i < len(this.M_sessions); i++ {
		if this.M_sessions[i].GetUUID() != ref.(*Session).GetUUID() {
			sessionss = append(sessionss, this.M_sessions[i])
		} else {
			isExist = true
			sessionss = append(sessionss, ref.(*Session))
		}
	}
	if !isExist {
		sessionss = append(sessionss, ref.(*Session))
		if this.IsInit == true {
			this.NeedSave = true
		}
		this.M_sessions = sessionss
	}
}

/** Remove reference Sessions **/
func (this *Account) RemoveSessions(ref interface{}) {
	toDelete := ref.(*Session)
	sessions_ := make([]*Session, 0)
	for i := 0; i < len(this.M_sessions); i++ {
		if toDelete.GetUUID() != this.M_sessions[i].GetUUID() {
			sessions_ = append(sessions_, this.M_sessions[i])
		} else {
			this.NeedSave = true
		}
	}
	this.M_sessions = sessions_
}

/** Messages **/
func (this *Account) GetMessages() []Message {
	return this.M_messages
}

/** Init reference Messages **/
func (this *Account) SetMessages(ref interface{}) {
	isExist := false
	var messagess []Message
	for i := 0; i < len(this.M_messages); i++ {
		if this.M_messages[i].(Entity).GetUUID() != ref.(Entity).GetUUID() {
			messagess = append(messagess, this.M_messages[i])
		} else {
			isExist = true
			messagess = append(messagess, ref.(Message))
		}
	}
	if !isExist {
		messagess = append(messagess, ref.(Message))
		if this.IsInit == true {
			this.NeedSave = true
		}
		this.M_messages = messagess
	}
}

/** Remove reference Messages **/
func (this *Account) RemoveMessages(ref interface{}) {
	toDelete := ref.(Entity)
	messages_ := make([]Message, 0)
	for i := 0; i < len(this.M_messages); i++ {
		if toDelete.GetUUID() != this.M_messages[i].(Entity).GetUUID() {
			messages_ = append(messages_, this.M_messages[i])
		} else {
			this.NeedSave = true
		}
	}
	this.M_messages = messages_
}

/** UserRef **/
func (this *Account) GetUserRef() *User {
	return this.m_userRef
}

/** Init reference UserRef **/
func (this *Account) SetUserRef(ref interface{}) {
	if _, ok := ref.(string); ok {
		if this.M_userRef != ref.(string) {
			this.M_userRef = ref.(string)
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
	} else {
		if this.M_userRef != ref.(Entity).GetUUID() {
			this.M_userRef = ref.(Entity).GetUUID()
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
		this.m_userRef = ref.(*User)
	}
}

/** Remove reference UserRef **/
func (this *Account) RemoveUserRef(ref interface{}) {
	toDelete := ref.(Entity)
	if this.m_userRef != nil {
		if toDelete.GetUUID() == this.m_userRef.GetUUID() {
			this.m_userRef = nil
			this.M_userRef = ""
			this.NeedSave = true
		}
	}
}

/** RolesRef **/
func (this *Account) GetRolesRef() []*Role {
	return this.m_rolesRef
}

/** Init reference RolesRef **/
func (this *Account) SetRolesRef(ref interface{}) {
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_rolesRef); i++ {
			if this.M_rolesRef[i] == refStr {
				return
			}
		}
		this.M_rolesRef = append(this.M_rolesRef, ref.(string))
		if this.IsInit == true {
			this.NeedSave = true
		}
	} else {
		for i := 0; i < len(this.m_rolesRef); i++ {
			if this.m_rolesRef[i].GetUUID() == ref.(*Role).GetUUID() {
				return
			}
		}
		isExist := false
		for i := 0; i < len(this.M_rolesRef); i++ {
			if this.M_rolesRef[i] == ref.(*Role).GetUUID() {
				isExist = true
			}
		}
		this.m_rolesRef = append(this.m_rolesRef, ref.(*Role))
		if !isExist {
			this.M_rolesRef = append(this.M_rolesRef, ref.(*Role).GetUUID())
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
	}
}

/** Remove reference RolesRef **/
func (this *Account) RemoveRolesRef(ref interface{}) {
	toDelete := ref.(*Role)
	rolesRef_ := make([]*Role, 0)
	rolesRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_rolesRef); i++ {
		if toDelete.GetUUID() != this.m_rolesRef[i].GetUUID() {
			rolesRef_ = append(rolesRef_, this.m_rolesRef[i])
			rolesRefUuid = append(rolesRefUuid, this.M_rolesRef[i])
		} else {
			this.NeedSave = true
		}
	}
	this.m_rolesRef = rolesRef_
	this.M_rolesRef = rolesRefUuid
}

/** PermissionsRef **/
func (this *Account) GetPermissionsRef() []*Permission {
	return this.m_permissionsRef
}

/** Init reference PermissionsRef **/
func (this *Account) SetPermissionsRef(ref interface{}) {
	if refStr, ok := ref.(string); ok {
		for i := 0; i < len(this.M_permissionsRef); i++ {
			if this.M_permissionsRef[i] == refStr {
				return
			}
		}
		this.M_permissionsRef = append(this.M_permissionsRef, ref.(string))
		if this.IsInit == true {
			this.NeedSave = true
		}
	} else {
		for i := 0; i < len(this.m_permissionsRef); i++ {
			if this.m_permissionsRef[i].GetUUID() == ref.(*Permission).GetUUID() {
				return
			}
		}
		isExist := false
		for i := 0; i < len(this.M_permissionsRef); i++ {
			if this.M_permissionsRef[i] == ref.(*Permission).GetUUID() {
				isExist = true
			}
		}
		this.m_permissionsRef = append(this.m_permissionsRef, ref.(*Permission))
		if !isExist {
			this.M_permissionsRef = append(this.M_permissionsRef, ref.(*Permission).GetUUID())
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
	}
}

/** Remove reference PermissionsRef **/
func (this *Account) RemovePermissionsRef(ref interface{}) {
	toDelete := ref.(*Permission)
	permissionsRef_ := make([]*Permission, 0)
	permissionsRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_permissionsRef); i++ {
		if toDelete.GetUUID() != this.m_permissionsRef[i].GetUUID() {
			permissionsRef_ = append(permissionsRef_, this.m_permissionsRef[i])
			permissionsRefUuid = append(permissionsRefUuid, this.M_permissionsRef[i])
		} else {
			this.NeedSave = true
		}
	}
	this.m_permissionsRef = permissionsRef_
	this.M_permissionsRef = permissionsRefUuid
}

/** Entities **/
func (this *Account) GetEntitiesPtr() *Entities {
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Account) SetEntitiesPtr(ref interface{}) {
	if _, ok := ref.(string); ok {
		if this.M_entitiesPtr != ref.(string) {
			this.M_entitiesPtr = ref.(string)
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
	} else {
		if this.M_entitiesPtr != ref.(*Entities).GetUUID() {
			this.M_entitiesPtr = ref.(*Entities).GetUUID()
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Account) RemoveEntitiesPtr(ref interface{}) {
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr != nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
