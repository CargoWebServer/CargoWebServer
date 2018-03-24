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
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of Entity **/
	M_id string

	/** members of Message **/
	M_body string

	/** members of Error **/
	M_errorPath string
	M_code int
	M_accountRef string


	/** Associations **/
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
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
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

func (this *Error) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Error) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *Error) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Error) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *Error) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *Error) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *Error) GetId()string{
	return this.M_id
}

func (this *Error) SetId(val string){
	this.M_id= val
}




func (this *Error) GetBody()string{
	return this.M_body
}

func (this *Error) SetBody(val string){
	this.M_body= val
}




func (this *Error) GetErrorPath()string{
	return this.M_errorPath
}

func (this *Error) SetErrorPath(val string){
	this.M_errorPath= val
}




func (this *Error) GetCode()int{
	return this.M_code
}

func (this *Error) SetCode(val int){
	this.M_code= val
}




func (this *Error) GetAccountRef()*Account{
	entity, err := this.getEntityByUuid(this.M_accountRef)
	if err == nil {
		return entity.(*Account)
	}
	return nil
}

func (this *Error) SetAccountRef(val *Account){
	this.M_accountRef= val.GetUuid()
	this.setEntity(this)
}


func (this *Error) ResetAccountRef(){
	this.M_accountRef= ""
}


func (this *Error) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Error) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
	this.setEntity(this)
}


func (this *Error) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

