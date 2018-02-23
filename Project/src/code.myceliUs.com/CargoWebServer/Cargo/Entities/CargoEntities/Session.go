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

	/** members of Session **/
	M_id string
	M_startTime int64
	M_endTime int64
	M_statusTime int64
	M_sessionState SessionState
	m_computerRef *Computer
	/** If the ref is a string and not an object **/
	M_computerRef string


	/** Associations **/
	m_accountPtr *Account
	/** If the ref is a string and not an object **/
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
	return this.UUID
}
func (this *Session) SetUuid(uuid string){
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

/** Id **/
func (this *Session) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Session) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** StartTime **/
func (this *Session) GetStartTime() int64{
	return this.M_startTime
}

/** Init reference StartTime **/
func (this *Session) SetStartTime(ref interface{}){
	if this.M_startTime != ref.(int64) {
		this.M_startTime = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference StartTime **/

/** EndTime **/
func (this *Session) GetEndTime() int64{
	return this.M_endTime
}

/** Init reference EndTime **/
func (this *Session) SetEndTime(ref interface{}){
	if this.M_endTime != ref.(int64) {
		this.M_endTime = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference EndTime **/

/** StatusTime **/
func (this *Session) GetStatusTime() int64{
	return this.M_statusTime
}

/** Init reference StatusTime **/
func (this *Session) SetStatusTime(ref interface{}){
	if this.M_statusTime != ref.(int64) {
		this.M_statusTime = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference StatusTime **/

/** SessionState **/
func (this *Session) GetSessionState() SessionState{
	return this.M_sessionState
}

/** Init reference SessionState **/
func (this *Session) SetSessionState(ref interface{}){
	if this.M_sessionState != ref.(SessionState) {
		this.M_sessionState = ref.(SessionState)
		this.NeedSave = true
	}
}

/** Remove reference SessionState **/

/** ComputerRef **/
func (this *Session) GetComputerRef() *Computer{
	if this.m_computerRef == nil {
		entity, err := this.getEntityByUuid(this.M_computerRef)
		if err == nil {
			this.m_computerRef = entity.(*Computer)
		}
	}
	return this.m_computerRef
}
func (this *Session) GetComputerRefStr() string{
	return this.M_computerRef
}

/** Init reference ComputerRef **/
func (this *Session) SetComputerRef(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_computerRef != ref.(string) {
			this.M_computerRef = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_computerRef != ref.(Entity).GetUuid() {
			this.M_computerRef = ref.(Entity).GetUuid()
			this.NeedSave = true
		}
		this.m_computerRef = ref.(*Computer)
	}
}

/** Remove reference ComputerRef **/
func (this *Session) RemoveComputerRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_computerRef!= nil {
		if toDelete.GetUuid() == this.m_computerRef.GetUuid() {
			this.m_computerRef = nil
			this.M_computerRef = ""
			this.NeedSave = true
		}
	}
}

/** Account **/
func (this *Session) GetAccountPtr() *Account{
	if this.m_accountPtr == nil {
		entity, err := this.getEntityByUuid(this.M_accountPtr)
		if err == nil {
			this.m_accountPtr = entity.(*Account)
		}
	}
	return this.m_accountPtr
}
func (this *Session) GetAccountPtrStr() string{
	return this.M_accountPtr
}

/** Init reference Account **/
func (this *Session) SetAccountPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_accountPtr != ref.(string) {
			this.M_accountPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_accountPtr != ref.(Entity).GetUuid() {
			this.M_accountPtr = ref.(Entity).GetUuid()
			this.NeedSave = true
		}
		this.m_accountPtr = ref.(*Account)
	}
}

/** Remove reference Account **/
func (this *Session) RemoveAccountPtr(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_accountPtr!= nil {
		if toDelete.GetUuid() == this.m_accountPtr.GetUuid() {
			this.m_accountPtr = nil
			this.M_accountPtr = ""
			this.NeedSave = true
		}
	}
}
