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

	/** If the entity is fully initialyse **/
	IsInit   bool

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
/** UUID **/
func (this *Error) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Error) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Error) SetId(ref interface{}){
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Body **/
func (this *Error) GetBody() string{
	return this.M_body
}

/** Init reference Body **/
func (this *Error) SetBody(ref interface{}){
	this.M_body = ref.(string)
}

/** Remove reference Body **/

/** ErrorPath **/
func (this *Error) GetErrorPath() string{
	return this.M_errorPath
}

/** Init reference ErrorPath **/
func (this *Error) SetErrorPath(ref interface{}){
	this.M_errorPath = ref.(string)
}

/** Remove reference ErrorPath **/

/** Code **/
func (this *Error) GetCode() int{
	return this.M_code
}

/** Init reference Code **/
func (this *Error) SetCode(ref interface{}){
	this.M_code = ref.(int)
}

/** Remove reference Code **/

/** AccountRef **/
func (this *Error) GetAccountRef() *Account{
	return this.m_accountRef
}

/** Init reference AccountRef **/
func (this *Error) SetAccountRef(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_accountRef = ref.(string)
	}else{
		this.M_accountRef = ref.(Entity).GetUUID()
		this.m_accountRef = ref.(*Account)
	}
}

/** Remove reference AccountRef **/
func (this *Error) RemoveAccountRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_accountRef!= nil {
		if toDelete.GetUUID() == this.m_accountRef.GetUUID() {
			this.m_accountRef = nil
			this.M_accountRef = ""
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
		this.M_entitiesPtr = ref.(string)
	}else{
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Error) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
		}
	}
}
