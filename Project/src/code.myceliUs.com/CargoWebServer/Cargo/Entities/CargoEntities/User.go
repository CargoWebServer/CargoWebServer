// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type User struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** keep track if the entity has change over time. **/
	needSave bool
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

	/** members of User **/
	M_firstName string
	M_lastName string
	M_middle string
	M_phone string
	M_email string
	M_memberOfRef []string
	M_accounts []string


	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for User **/
type XsdUser struct {
	XMLName xml.Name	`xml:"userRef"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_memberOfRef	[]string	`xml:"memberOfRef"`
	M_firstName	string	`xml:"firstName,attr"`
	M_lastName	string	`xml:"lastName,attr"`
	M_middle	string	`xml:"middle,attr"`
	M_email	string	`xml:"email,attr"`
	M_phone	string	`xml:"phone,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *User) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *User) SetUuid(uuid string){
	this.UUID = uuid
}

/** Need save **/
func (this *User) IsNeedSave() bool{
	return this.needSave
}
func (this *User) SetNeedSave(needSave bool){
	this.needSave=needSave
}

func (this *User) GetReferenced() []string {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	// return the list of references
	return this.Referenced
}

func (this *User) SetReferenced(uuid string, field string) {
	if this.Referenced == nil {
		this.Referenced = make([]string, 0)
	}
	if !Utility.Contains(this.Referenced, uuid+":"+field) {
		this.Referenced = append(this.Referenced, uuid+":"+field)
	}
}

func (this *User) RemoveReferenced(uuid string, field string) {
	if this.Referenced == nil {
		return
	}
	referenced := make([]string, 0)
	for i := 0; i < len(this.Referenced); i++ {
		if this.Referenced[i] != uuid+":"+field {
			referenced = append(referenced, uuid+":"+field)
		}
	}
	this.Referenced = referenced
}

func (this *User) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *User) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *User) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *User) GetTypeName() string{
	this.TYPENAME = "CargoEntities.User"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *User) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *User) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *User) GetParentLnk() string{
	return this.ParentLnk
}
func (this *User) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *User) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *User) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *User) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *User) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *User) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *User) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *User) GetId()string{
	return this.M_id
}

func (this *User) SetId(val string){
	this.M_id= val
}




func (this *User) GetFirstName()string{
	return this.M_firstName
}

func (this *User) SetFirstName(val string){
	this.M_firstName= val
}




func (this *User) GetLastName()string{
	return this.M_lastName
}

func (this *User) SetLastName(val string){
	this.M_lastName= val
}




func (this *User) GetMiddle()string{
	return this.M_middle
}

func (this *User) SetMiddle(val string){
	this.M_middle= val
}




func (this *User) GetPhone()string{
	return this.M_phone
}

func (this *User) SetPhone(val string){
	this.M_phone= val
}




func (this *User) GetEmail()string{
	return this.M_email
}

func (this *User) SetEmail(val string){
	this.M_email= val
}




func (this *User) GetMemberOfRef()[]*Group{
	values := make([]*Group, 0)
	for i := 0; i < len(this.M_memberOfRef); i++ {
		entity, err := this.getEntityByUuid(this.M_memberOfRef[i])
		if err == nil {
			values = append( values, entity.(*Group))
		}
	}
	return values
}

func (this *User) SetMemberOfRef(val []*Group){
	this.M_memberOfRef= make([]string,0)
	for i:=0; i < len(val); i++{
		this.setEntity(val[i])
	}
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *User) AppendMemberOfRef(val *Group){
	for i:=0; i < len(this.M_memberOfRef); i++{
		if this.M_memberOfRef[i] == val.GetUuid() {
			return
		}
	}
	if this.M_memberOfRef== nil {
		this.M_memberOfRef = make([]string, 0)
	}

	this.M_memberOfRef = append(this.M_memberOfRef, val.GetUuid())
	this.setEntity(this)
	this.SetNeedSave(true)
}

func (this *User) RemoveMemberOfRef(val *Group){
	values := make([]string,0)
	for i:=0; i < len(this.M_memberOfRef); i++{
		if this.M_memberOfRef[i] != val.GetUuid() {
			values = append(values, this.M_memberOfRef[i])
		}
	}
	this.M_memberOfRef = values
	this.setEntity(this)
}


func (this *User) GetAccounts()[]*Account{
	values := make([]*Account, 0)
	for i := 0; i < len(this.M_accounts); i++ {
		entity, err := this.getEntityByUuid(this.M_accounts[i])
		if err == nil {
			values = append( values, entity.(*Account))
		}
	}
	return values
}

func (this *User) SetAccounts(val []*Account){
	this.M_accounts= make([]string,0)
	for i:=0; i < len(val); i++{
		this.setEntity(val[i])
	}
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *User) AppendAccounts(val *Account){
	for i:=0; i < len(this.M_accounts); i++{
		if this.M_accounts[i] == val.GetUuid() {
			return
		}
	}
	if this.M_accounts== nil {
		this.M_accounts = make([]string, 0)
	}

	this.M_accounts = append(this.M_accounts, val.GetUuid())
	this.setEntity(this)
	this.SetNeedSave(true)
}

func (this *User) RemoveAccounts(val *Account){
	values := make([]string,0)
	for i:=0; i < len(this.M_accounts); i++{
		if this.M_accounts[i] != val.GetUuid() {
			values = append(values, this.M_accounts[i])
		}
	}
	this.M_accounts = values
	this.setEntity(this)
}


func (this *User) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *User) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
	this.setEntity(this)
	this.SetNeedSave(true)
}


func (this *User) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
	this.setEntity(this)
}

