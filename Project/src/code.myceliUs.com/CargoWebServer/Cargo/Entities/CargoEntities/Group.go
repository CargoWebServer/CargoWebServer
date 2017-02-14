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
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

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
/** UUID **/
func (this *Group) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Group) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Group) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Name **/
func (this *Group) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Group) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** MembersRef **/
func (this *Group) GetMembersRef() []*User{
	return this.m_membersRef
}

/** Init reference MembersRef **/
func (this *Group) SetMembersRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_membersRef); i++ {
			if this.M_membersRef[i] == refStr {
				return
			}
		}
		this.M_membersRef = append(this.M_membersRef, ref.(string))
	}else{
		this.RemoveMembersRef(ref)
		this.m_membersRef = append(this.m_membersRef, ref.(*User))
		this.M_membersRef = append(this.M_membersRef, ref.(Entity).GetUUID())
	}
}

/** Remove reference MembersRef **/
func (this *Group) RemoveMembersRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Entity)
	membersRef_ := make([]*User, 0)
	membersRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_membersRef); i++ {
		if toDelete.GetUUID() != this.m_membersRef[i].GetUUID() {
			membersRef_ = append(membersRef_, this.m_membersRef[i])
			membersRefUuid = append(membersRefUuid, this.M_membersRef[i])
		}
	}
	this.m_membersRef = membersRef_
	this.M_membersRef = membersRefUuid
}

/** Entities **/
func (this *Group) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Group) SetEntitiesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *Group) RemoveEntitiesPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Entities)
	if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
		this.m_entitiesPtr = nil
		this.M_entitiesPtr = ""
	}
}
