// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type Project struct{

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

	/** members of Project **/
	M_name string
	m_filesRef []*File
	/** If the ref is a string and not an object **/
	M_filesRef []string


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for Project **/
type XsdProject struct {
	XMLName xml.Name	`xml:"project"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_filesRef	[]string	`xml:"filesRef"`
	M_name	string	`xml:"name,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *Project) GetUuid() string{
	return this.UUID
}
func (this *Project) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *Project) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *Project) GetTypeName() string{
	this.TYPENAME = "CargoEntities.Project"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *Project) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *Project) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *Project) GetParentLnk() string{
	return this.ParentLnk
}
func (this *Project) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *Project) IsNeedSave() bool{
	return this.NeedSave
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Project) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *Project) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Project) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Name **/
func (this *Project) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Project) SetName(ref interface{}){
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Name **/

/** FilesRef **/
func (this *Project) GetFilesRef() []*File{
	if this.m_filesRef == nil {
		this.m_filesRef = make([]*File, 0)
		for i := 0; i < len(this.M_filesRef); i++ {
			entity, err := this.getEntityByUuid(this.M_filesRef[i])
			if err == nil {
				this.m_filesRef = append(this.m_filesRef, entity.(*File))
			}
		}
	}
	return this.m_filesRef
}
func (this *Project) GetFilesRefStr() []string{
	return this.M_filesRef
}

/** Init reference FilesRef **/
func (this *Project) SetFilesRef(ref interface{}){
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_filesRef); i++ {
			if this.M_filesRef[i] == refStr {
				return
			}
		}
		this.M_filesRef = append(this.M_filesRef, ref.(string))
		this.NeedSave = true
	}else{
		for i:=0; i < len(this.m_filesRef); i++ {
			if this.m_filesRef[i].GetUuid() == ref.(*File).GetUuid() {
				return
			}
		}
		isExist := false
		for i:=0; i < len(this.M_filesRef); i++ {
			if this.M_filesRef[i] == ref.(*File).GetUuid() {
				isExist = true
			}
		}
		this.m_filesRef = append(this.m_filesRef, ref.(*File))
	if !isExist {
		this.M_filesRef = append(this.M_filesRef, ref.(Entity).GetUuid())
		this.NeedSave = true
	}
	}
}

/** Remove reference FilesRef **/
func (this *Project) RemoveFilesRef(ref interface{}){
	toDelete := ref.(Entity)
	filesRef_ := make([]*File, 0)
	filesRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_filesRef); i++ {
		if toDelete.GetUuid() != this.m_filesRef[i].GetUuid() {
			filesRef_ = append(filesRef_, this.m_filesRef[i])
			filesRefUuid = append(filesRefUuid, this.M_filesRef[i])
		}else{
			this.NeedSave = true
		}
	}
	this.m_filesRef = filesRef_
	this.M_filesRef = filesRefUuid
}

/** Entities **/
func (this *Project) GetEntitiesPtr() *Entities{
	if this.m_entitiesPtr == nil {
		entity, err := this.getEntityByUuid(this.M_entitiesPtr)
		if err == nil {
			this.m_entitiesPtr = entity.(*Entities)
		}
	}
	return this.m_entitiesPtr
}
func (this *Project) GetEntitiesPtrStr() string{
	return this.M_entitiesPtr
}

/** Init reference Entities **/
func (this *Project) SetEntitiesPtr(ref interface{}){
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
func (this *Project) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
