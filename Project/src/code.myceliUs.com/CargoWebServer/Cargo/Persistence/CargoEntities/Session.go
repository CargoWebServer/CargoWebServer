package CargoEntities

import(
"encoding/xml"
)

type Session struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

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
/** UUID **/
func (this *Session) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Session) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Session) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** StartTime **/
func (this *Session) GetStartTime() int64{
	return this.M_startTime
}

/** Init reference StartTime **/
func (this *Session) SetStartTime(ref interface{}){
	this.NeedSave = true
	this.M_startTime = ref.(int64)
}

/** Remove reference StartTime **/

/** EndTime **/
func (this *Session) GetEndTime() int64{
	return this.M_endTime
}

/** Init reference EndTime **/
func (this *Session) SetEndTime(ref interface{}){
	this.NeedSave = true
	this.M_endTime = ref.(int64)
}

/** Remove reference EndTime **/

/** StatusTime **/
func (this *Session) GetStatusTime() int64{
	return this.M_statusTime
}

/** Init reference StatusTime **/
func (this *Session) SetStatusTime(ref interface{}){
	this.NeedSave = true
	this.M_statusTime = ref.(int64)
}

/** Remove reference StatusTime **/

/** SessionState **/
func (this *Session) GetSessionState() SessionState{
	return this.M_sessionState
}

/** Init reference SessionState **/
func (this *Session) SetSessionState(ref interface{}){
	this.NeedSave = true
	this.M_sessionState = ref.(SessionState)
}

/** Remove reference SessionState **/

/** ComputerRef **/
func (this *Session) GetComputerRef() *Computer{
	return this.m_computerRef
}

/** Init reference ComputerRef **/
func (this *Session) SetComputerRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_computerRef = ref.(string)
	}else{
		this.m_computerRef = ref.(*Computer)
		this.M_computerRef = ref.(Entity).GetUUID()
	}
}

/** Remove reference ComputerRef **/
func (this *Session) RemoveComputerRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Entity)
	if toDelete.GetUUID() == this.m_computerRef.GetUUID() {
		this.m_computerRef = nil
		this.M_computerRef = ""
	}
}

/** Account **/
func (this *Session) GetAccountPtr() *Account{
	return this.m_accountPtr
}

/** Init reference Account **/
func (this *Session) SetAccountPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_accountPtr = ref.(string)
	}else{
		this.m_accountPtr = ref.(*Account)
		this.M_accountPtr = ref.(Entity).GetUUID()
	}
}

/** Remove reference Account **/
func (this *Session) RemoveAccountPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Entity)
	if toDelete.GetUUID() == this.m_accountPtr.GetUUID() {
		this.m_accountPtr = nil
		this.M_accountPtr = ""
	}
}
