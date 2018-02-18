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
	m_membersRef []*User
	/** If the ref is a string and not an object **/
	M_membersRef []string


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
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

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Group) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *Group) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Group) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Name **/
func (this *Group) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Group) SetName(ref interface{}){
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Name **/

/** MembersRef **/
func (this *Group) GetMembersRef() []*User{
	if this.m_membersRef == nil {
		this.m_membersRef = make([]*User, 0)
		for i := 0; i < len(this.M_membersRef); i++ {
			entity, err := this.getEntityByUuid(this.M_membersRef[i])
			if err == nil {
				this.m_membersRef = append(this.m_membersRef, entity.(*User))
			}
		}
	}
	return this.m_membersRef
}
func (this *Group) GetMembersRefStr() []string{
	return this.M_membersRef
}

/** Init reference MembersRef **/
func (this *Group) SetMembersRef(ref interface{}){
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_membersRef); i++ {
			if this.M_membersRef[i] == refStr {
				return
			}
		}
		this.M_membersRef = append(this.M_membersRef, ref.(string))
		this.NeedSave = true
	}else{
		for i:=0; i < len(this.m_membersRef); i++ {
			if this.m_membersRef[i].GetUuid() == ref.(*User).GetUuid() {
				return
			}
		}
		isExist := false
		for i:=0; i < len(this.M_membersRef); i++ {
			if this.M_membersRef[i] == ref.(*User).GetUuid() {
				isExist = true
			}
		}
		this.m_membersRef = append(this.m_membersRef, ref.(*User))
	if !isExist {
		this.M_membersRef = append(this.M_membersRef, ref.(Entity).GetUuid())
		this.NeedSave = true
	}
	}
}

/** Remove reference MembersRef **/
func (this *Group) RemoveMembersRef(ref interface{}){
	toDelete := ref.(Entity)
	membersRef_ := make([]*User, 0)
	membersRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_membersRef); i++ {
		if toDelete.GetUuid() != this.m_membersRef[i].GetUuid() {
			membersRef_ = append(membersRef_, this.m_membersRef[i])
			membersRefUuid = append(membersRefUuid, this.M_membersRef[i])
		}else{
			this.NeedSave = true
		}
	}
	this.m_membersRef = membersRef_
	this.M_membersRef = membersRefUuid
}

/** Entities **/
func (this *Group) GetEntitiesPtr() *Entities{
	if this.m_entitiesPtr == nil {
		entity, err := this.getEntityByUuid(this.M_entitiesPtr)
		if err == nil {
			this.m_entitiesPtr = entity.(*Entities)
		}
	}
	return this.m_entitiesPtr
}
func (this *Group) GetEntitiesPtrStr() string{
	return this.M_entitiesPtr
}

/** Init reference Entities **/
func (this *Group) SetEntitiesPtr(ref interface{}){
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
func (this *Group) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
