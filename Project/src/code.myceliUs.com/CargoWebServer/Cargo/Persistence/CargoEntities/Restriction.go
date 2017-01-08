package CargoEntities

import(
"encoding/xml"
)

type Restriction struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Restriction **/
	m_actionRef *Action
	/** If the ref is a string and not an object **/
	M_actionRef string
	m_rolesRef []*Role
	/** If the ref is a string and not an object **/
	M_rolesRef []string


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for Restriction **/
type XsdRestriction struct {
	XMLName xml.Name	`xml:"restriction"`

}
/** UUID **/
func (this *Restriction) GetUUID() string{
	return this.UUID
}

/** ActionRef **/
func (this *Restriction) GetActionRef() *Action{
	return this.m_actionRef
}

/** Init reference ActionRef **/
func (this *Restriction) SetActionRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_actionRef = ref.(string)
	}else{
		this.m_actionRef = ref.(*Action)
		this.M_actionRef = ref.(Entity).GetUUID()
	}
}

/** Remove reference ActionRef **/
func (this *Restriction) RemoveActionRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(Entity)
	if toDelete.GetUUID() == this.m_actionRef.GetUUID() {
		this.m_actionRef = nil
		this.M_actionRef = ""
	}
}

/** RolesRef **/
func (this *Restriction) GetRolesRef() []*Role{
	return this.m_rolesRef
}

/** Init reference RolesRef **/
func (this *Restriction) SetRolesRef(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_rolesRef); i++ {
			if this.M_rolesRef[i] == refStr {
				return
			}
		}
		this.M_rolesRef = append(this.M_rolesRef, ref.(string))
	}else{
		this.RemoveRolesRef(ref)
		this.m_rolesRef = append(this.m_rolesRef, ref.(*Role))
		this.M_rolesRef = append(this.M_rolesRef, ref.(*Role).GetUUID())
	}
}

/** Remove reference RolesRef **/
func (this *Restriction) RemoveRolesRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Role)
	rolesRef_ := make([]*Role, 0)
	rolesRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_rolesRef); i++ {
		if toDelete.GetUUID() != this.m_rolesRef[i].GetUUID() {
			rolesRef_ = append(rolesRef_, this.m_rolesRef[i])
			rolesRefUuid = append(rolesRefUuid, this.m_rolesRef[i].GetUUID())
		}
	}
	this.m_rolesRef = rolesRef_
	this.M_rolesRef = rolesRefUuid
}

/** Entities **/
func (this *Restriction) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Restriction) SetEntitiesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *Restriction) RemoveEntitiesPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Entities)
	if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
		this.m_entitiesPtr = nil
		this.M_entitiesPtr = ""
	}
}
