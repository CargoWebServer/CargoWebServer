// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type Error struct{

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

	/** members of Entity **/
	M_id string

	/** members of Message **/
	M_body string

	/** members of Error **/
	M_errorPath string
	M_code int
	m_accountRef *Account
	/** If the ref is a string and not an object **/
	M_accountRef string


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for Error **/
type XsdError struct {
	XMLName xml.Name	`xml:"error"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	/** Message **/
	M_body	string	`xml:"body,attr"`


	M_code	string	`xml:"code,attr"`
	M_errorPath	string	`xml:"errorPath,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Error) GetUuid() string{
	return this.UUID
}
func (this *Error) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Error) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Error) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Error"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Error) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Error) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Error) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Error) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *Error) IsNeedSave() bool{
	return this.NeedSave
}


/** Id **/
func (this *Error) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Error) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Body **/
func (this *Error) GetBody() string{
	return this.M_body
}

/** Init reference Body **/
func (this *Error) SetBody(ref interface{}){
	if this.M_body != ref.(string) {
		this.M_body = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Body **/

/** ErrorPath **/
func (this *Error) GetErrorPath() string{
	return this.M_errorPath
}

/** Init reference ErrorPath **/
func (this *Error) SetErrorPath(ref interface{}){
	if this.M_errorPath != ref.(string) {
		this.M_errorPath = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference ErrorPath **/

/** Code **/
func (this *Error) GetCode() int{
	return this.M_code
}

/** Init reference Code **/
func (this *Error) SetCode(ref interface{}){
	if this.M_code != ref.(int) {
		this.M_code = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference Code **/

/** AccountRef **/
func (this *Error) GetAccountRef() *Account{
	return this.m_accountRef
}

/** Init reference AccountRef **/
func (this *Error) SetAccountRef(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_accountRef != ref.(string) {
			this.M_accountRef = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_accountRef != ref.(Entity).GetUuid() {
			this.M_accountRef = ref.(Entity).GetUuid()
			this.NeedSave = true
		}
		this.m_accountRef = ref.(*Account)
	}
}

/** Remove reference AccountRef **/
func (this *Error) RemoveAccountRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_accountRef!= nil {
		if toDelete.GetUuid() == this.m_accountRef.GetUuid() {
			this.m_accountRef = nil
			this.M_accountRef = ""
			this.NeedSave = true
		}
	}
}

/** Entities **/
func (this *Error) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Error) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_entitiesPtr != ref.(string) {
			this.M_entitiesPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_entitiesPtr != ref.(*Entities).GetUuid() {
			this.M_entitiesPtr = ref.(*Entities).GetUuid()
			this.NeedSave = true
		}
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Error) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
