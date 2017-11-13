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
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
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
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference Name **/

/** MembersRef **/
func (this *Group) GetMembersRef() []*User{
	return this.m_membersRef
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
		if this.IsInit == true {			this.NeedSave = true
		}
	}else{
		for i:=0; i < len(this.m_membersRef); i++ {
			if this.m_membersRef[i].GetUUID() == ref.(*User).GetUUID() {
				return
			}
		}
		isExist := false
		for i:=0; i < len(this.M_membersRef); i++ {
			if this.M_membersRef[i] == ref.(*User).GetUUID() {
				isExist = true
			}
		}
		this.m_membersRef = append(this.m_membersRef, ref.(*User))
	if !isExist {
		this.M_membersRef = append(this.M_membersRef, ref.(Entity).GetUUID())
		if this.IsInit == true {			this.NeedSave = true
		}
	}
	}
}

/** Remove reference MembersRef **/
func (this *Group) RemoveMembersRef(ref interface{}){
	toDelete := ref.(Entity)
	membersRef_ := make([]*User, 0)
	membersRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_membersRef); i++ {
		if toDelete.GetUUID() != this.m_membersRef[i].GetUUID() {
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
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Group) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_entitiesPtr != ref.(string) {
			this.M_entitiesPtr = ref.(string)
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_entitiesPtr != ref.(*Entities).GetUUID() {
			this.M_entitiesPtr = ref.(*Entities).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Group) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
