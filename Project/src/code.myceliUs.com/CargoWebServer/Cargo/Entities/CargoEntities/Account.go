// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
	"strings"
)

type Account struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** Keep reference to entity that made use of thit entity **/
	Referenced []string
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of Entity **/
	M_id string

	/** members of Account **/
	M_name string
	M_password string
	M_email string
	M_sessions []string
	M_messages []string
	M_userRef string
	M_rolesRef []string
	M_permissionsRef []string


	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for Account **/
type XsdAccount struct {
	XMLName xml.Name	`xml:"accountsRef"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_userRef	*string	`xml:"userRef"`
	M_rolesRef	[]string	`xml:"rolesRef"`
	M_permissionsRef	[]string	`xml:"permissionsRef"`
	M_sessions	[]*XsdSession	`xml:"sessions,omitempty"`

	M_name	string	`xml:"name,attr"`
	M_password	string	`xml:"password,attr"`
	M_email	string	`xml:"email,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Account) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Account) SetUuid(uuid string){
	this.UUID = uuid
}

func (this *Account) GetReferenced() []string {
 if this.Referenced == nil {
 	this.Referenced = make([]string, 0)
 }
	// return the list of references
	return this.Referenced
}

func (this *Account) SetReferenced(uuid string, field string){
 if this.Referenced == nil {
 	this.Referenced = make([]string, 0)
 }
 if !Utility.Contains(this.Referenced, uuid+":"+field) {
 	this.Referenced = append(this.Referenced, uuid+":"+field)
 }
}

func (this *Account) RemoveReferenced(uuid string, field string){
 if this.Referenced == nil {
 	return
 }
 referenced := make([]string,0)
 for i:=0; i < len(this.Referenced); i++ {
 	if this.Referenced[i] != uuid+":"+field {
 		referenced = append(referenced, uuid+":"+field)
 	}
 }
 	this.Referenced = referenced
}

func (this *Account) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *Account) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *Account) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Account) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Account"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Account) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Account) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Account) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Account) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *Account) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Account) GetChilds() []interface{}{
	var childs []interface{}
	var child interface{}
	var err error
	for i:=0; i < len(this.M_sessions); i++ {
		child, err = this.getEntityByUuid( this.M_sessions[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	for i:=0; i < len(this.M_messages); i++ {
		child, err = this.getEntityByUuid( this.M_messages[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	return childs
}
/** Return the list of all childs uuid **/
func (this *Account) GetChildsUuid() []string{
	var childs []string
	childs = append( childs, this.M_sessions...)
	childs = append( childs, this.M_messages...)
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Account) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Account) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Account) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Account) GetId()string{
	return this.M_id
}

func (this *Account) SetId(val string){
	this.M_id= val
}




func (this *Account) GetName()string{
	return this.M_name
}

func (this *Account) SetName(val string){
	this.M_name= val
}




func (this *Account) GetPassword()string{
	return this.M_password
}

func (this *Account) SetPassword(val string){
	this.M_password= val
}




func (this *Account) GetEmail()string{
	return this.M_email
}

func (this *Account) SetEmail(val string){
	this.M_email= val
}




func (this *Account) GetSessions()[]*Session{
	values := make([]*Session, 0)
	for i := 0; i < len(this.M_sessions); i++ {
		entity, err := this.getEntityByUuid(this.M_sessions[i])
		if err == nil {
			values = append( values, entity.(*Session))
		}
	}
	return values
}

func (this *Account) SetSessions(val []*Session){
	this.M_sessions= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_sessions=append(this.M_sessions, val[i].GetUuid())
		if len(val[i].GetParentUuid()) > 0  &&  len(val[i].GetParentLnk()) > 0 && this.GetUuid() != val[i].GetParentUuid(){
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_sessions")
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *Account) AppendSessions(val *Session){
	for i:=0; i < len(this.M_sessions); i++{
		if this.M_sessions[i] == val.GetUuid() {
			return
		}
	}
	this.M_sessions = append(this.M_sessions, val.GetUuid())
	if len(val.GetParentUuid()) > 0 &&  len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_sessions")
  this.setEntity(val)
	this.setEntity(this)
}

func (this *Account) RemoveSessions(val *Session){
	values := make([]string,0)
	for i:=0; i < len(this.M_sessions); i++{
		if this.M_sessions[i] != val.GetUuid() {
			values = append(values, this.M_sessions[i])
		}
	}
	this.M_sessions = values
	this.setEntity(this)
}


func (this *Account) GetMessages()[]Message{
	values := make([]Message, 0)
	for i := 0; i < len(this.M_messages); i++ {
		entity, err := this.getEntityByUuid(this.M_messages[i])
		if err == nil {
			values = append( values, entity.(Message))
		}
	}
	return values
}

func (this *Account) SetMessages(val []Message){
	this.M_messages= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_messages=append(this.M_messages, val[i].GetUuid())
		if len(val[i].GetParentUuid()) > 0  &&  len(val[i].GetParentLnk()) > 0 && this.GetUuid() != val[i].GetParentUuid(){
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_messages")
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *Account) AppendMessages(val Message){
	for i:=0; i < len(this.M_messages); i++{
		if this.M_messages[i] == val.GetUuid() {
			return
		}
	}
	this.M_messages = append(this.M_messages, val.GetUuid())
	if len(val.GetParentUuid()) > 0 &&  len(val.GetParentLnk()) > 0 && val.GetParentUuid() != this.GetUuid() {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_messages")
  this.setEntity(val)
	this.setEntity(this)
}

func (this *Account) RemoveMessages(val Message){
	values := make([]string,0)
	for i:=0; i < len(this.M_messages); i++{
		if this.M_messages[i] != val.GetUuid() {
			values = append(values, this.M_messages[i])
		}
	}
	this.M_messages = values
	this.setEntity(this)
}


func (this *Account) GetUserRef()*User{
	entity, err := this.getEntityByUuid(this.M_userRef)
	if err == nil {
		return entity.(*User)
	}
	return nil
}

func (this *Account) SetUserRef(val *User){
	this.M_userRef= val.GetUuid()
		val.SetReferenced(this.UUID,"M_userRef")
	this.setEntity(this)
}


func (this *Account) ResetUserRef(){
	this.M_userRef= ""
}


func (this *Account) GetRolesRef()[]*Role{
	values := make([]*Role, 0)
	for i := 0; i < len(this.M_rolesRef); i++ {
		entity, err := this.getEntityByUuid(this.M_rolesRef[i])
		if err == nil {
			values = append( values, entity.(*Role))
		}
	}
	return values
}

func (this *Account) SetRolesRef(val []*Role){
	this.M_rolesRef= make([]string,0)
	for i:=0; i < len(val); i++{
		val[i].SetReferenced(this.UUID,"M_rolesRef")
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *Account) AppendRolesRef(val *Role){
	for i:=0; i < len(this.M_rolesRef); i++{
		if this.M_rolesRef[i] == val.GetUuid() {
			return
		}
	}
	this.M_rolesRef = append(this.M_rolesRef, val.GetUuid())
	val.SetReferenced(this.UUID,"M_rolesRef")
	this.setEntity(this)
}

func (this *Account) RemoveRolesRef(val *Role){
	values := make([]string,0)
	for i:=0; i < len(this.M_rolesRef); i++{
		if this.M_rolesRef[i] != val.GetUuid() {
			values = append(values, this.M_rolesRef[i])
		}
	}
	this.M_rolesRef = values
	this.setEntity(this)
	val.RemoveReferenced(this.UUID,"M_rolesRef")
}


func (this *Account) GetPermissionsRef()[]*Permission{
	values := make([]*Permission, 0)
	for i := 0; i < len(this.M_permissionsRef); i++ {
		entity, err := this.getEntityByUuid(this.M_permissionsRef[i])
		if err == nil {
			values = append( values, entity.(*Permission))
		}
	}
	return values
}

func (this *Account) SetPermissionsRef(val []*Permission){
	this.M_permissionsRef= make([]string,0)
	for i:=0; i < len(val); i++{
		val[i].SetReferenced(this.UUID,"M_permissionsRef")
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *Account) AppendPermissionsRef(val *Permission){
	for i:=0; i < len(this.M_permissionsRef); i++{
		if this.M_permissionsRef[i] == val.GetUuid() {
			return
		}
	}
	this.M_permissionsRef = append(this.M_permissionsRef, val.GetUuid())
	val.SetReferenced(this.UUID,"M_permissionsRef")
	this.setEntity(this)
}

func (this *Account) RemovePermissionsRef(val *Permission){
	values := make([]string,0)
	for i:=0; i < len(this.M_permissionsRef); i++{
		if this.M_permissionsRef[i] != val.GetUuid() {
			values = append(values, this.M_permissionsRef[i])
		}
	}
	this.M_permissionsRef = values
	this.setEntity(this)
	val.RemoveReferenced(this.UUID,"M_permissionsRef")
}


func (this *Account) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Account) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
		val.SetReferenced(this.UUID,"M_entitiesPtr")
	this.setEntity(this)
}


func (this *Account) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

