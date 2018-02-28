// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type Group struct{

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

	/** members of Entity **/
	M_id string

	/** members of Group **/
	M_name string
	M_membersRef []string


	/** Associations **/
	M_entitiesPtr string
}

/** Xml parser for Group **/
type XsdGroup struct {
	XMLName xml.Name	`xml:"memberOfRef"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_membersRef	[]string	`xml:"membersRef"`
	M_name	string	`xml:"name,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Group) GetUuid() string{
	return this.UUID
}
func (this *Group) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Group) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Group) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Group"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Group) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Group) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Group) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Group) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *Group) IsNeedSave() bool{
	return this.NeedSave
}
func (this *Group) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Group) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

func (this *Group) GetId()string{
	return this.M_id
}

func (this *Group) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *Group) GetName()string{
	return this.M_name
}

func (this *Group) SetName(val string){
	this.NeedSave = this.M_name== val
	this.M_name= val
}


func (this *Group) GetMembersRef()[]*User{
	membersRef := make([]*User, 0)
	for i := 0; i < len(this.M_membersRef); i++ {
		entity, err := this.getEntityByUuid(this.M_membersRef[i])
		if err == nil {
			membersRef = append(membersRef, entity.(*User))
		}
	}
	return membersRef
}

func (this *Group) SetMembersRef(val []*User){
	this.M_membersRef= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_membersRef=append(this.M_membersRef, val[i].GetUuid())
	}
}

func (this *Group) AppendMembersRef(val *User){
	for i:=0; i < len(this.M_membersRef); i++{
		if this.M_membersRef[i] == val.GetUuid() {
			return
		}
	}
	this.M_membersRef = append(this.M_membersRef, val.GetUuid())
}

func (this *Group) RemoveMembersRef(val *User){
	membersRef := make([]string,0)
	for i:=0; i < len(this.M_membersRef); i++{
		if this.M_membersRef[i] != val.GetUuid() {
			membersRef = append(membersRef, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_membersRef = membersRef
}


func (this *Group) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Group) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
}

func (this *Group) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

