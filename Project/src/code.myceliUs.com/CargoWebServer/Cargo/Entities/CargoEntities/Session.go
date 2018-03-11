// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type Session struct{

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

	/** members of Session **/
	M_id string
	M_startTime int64
	M_endTime int64
	M_statusTime int64
	M_sessionState SessionState
	M_computerRef string


	/** Associations **/
	M_accountPtr string
}

/** Xml parser for Session **/
type XsdSession struct {
	XMLName xml.Name	`xml:"session"`
	M_computerRef	*string	`xml:"computerRef"`
	M_id	string	`xml:"id,attr"`
	M_sessionState	string	`xml:"sessionState,attr"`
	M_startTime	int64	`xml:"startTime,attr"`
	M_endTime	int64	`xml:"endTime,attr"`
	M_statutTime	int64	`xml:"statutTime,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Session) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *Session) SetUuid(uuid string){
	this.NeedSave = this.UUID == uuid
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Session) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Session) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Session"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Session) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Session) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Session) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Session) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *Session) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Session) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Evaluate if an entity needs to be saved. **/
func (this *Session) IsNeedSave() bool{
	return this.NeedSave
}
func (this *Session) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Session) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Session) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Session) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Session) GetId()string{
	return this.M_id
}

func (this *Session) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}




func (this *Session) GetStartTime()int64{
	return this.M_startTime
}

func (this *Session) SetStartTime(val int64){
	this.NeedSave = this.M_startTime== val
	this.M_startTime= val
}




func (this *Session) GetEndTime()int64{
	return this.M_endTime
}

func (this *Session) SetEndTime(val int64){
	this.NeedSave = this.M_endTime== val
	this.M_endTime= val
}




func (this *Session) GetStatusTime()int64{
	return this.M_statusTime
}

func (this *Session) SetStatusTime(val int64){
	this.NeedSave = this.M_statusTime== val
	this.M_statusTime= val
}




func (this *Session) GetSessionState()SessionState{
	return this.M_sessionState
}

func (this *Session) SetSessionState(val SessionState){
	this.NeedSave = this.M_sessionState== val
	this.M_sessionState= val
}


func (this *Session) ResetSessionState(){
	this.M_sessionState= 0
}


func (this *Session) GetComputerRef()*Computer{
	entity, err := this.getEntityByUuid(this.M_computerRef)
	if err == nil {
		return entity.(*Computer)
	}
	return nil
}

func (this *Session) SetComputerRef(val *Computer){
	this.NeedSave = this.M_computerRef != val.GetUuid()
	this.M_computerRef= val.GetUuid()
}


func (this *Session) ResetComputerRef(){
	this.M_computerRef= ""
}


func (this *Session) GetAccountPtr()*Account{
	entity, err := this.getEntityByUuid(this.M_accountPtr)
	if err == nil {
		return entity.(*Account)
	}
	return nil
}

func (this *Session) SetAccountPtr(val *Account){
	this.NeedSave = this.M_accountPtr != val.GetUuid()
	this.M_accountPtr= val.GetUuid()
}


func (this *Session) ResetAccountPtr(){
	this.M_accountPtr= ""
}

