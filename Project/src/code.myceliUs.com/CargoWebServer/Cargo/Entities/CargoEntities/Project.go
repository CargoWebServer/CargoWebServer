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
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

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
/** UUID **/
func (this *Project) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Project) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Project) SetId(ref interface{}){
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Name **/
func (this *Project) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Project) SetName(ref interface{}){
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** FilesRef **/
func (this *Project) GetFilesRef() []*File{
	return this.m_filesRef
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
	}else{
		for i:=0; i < len(this.m_filesRef); i++ {
			if this.m_filesRef[i].GetUUID() == ref.(*File).GetUUID() {
				return
			}
		}
		isExist := false
		for i:=0; i < len(this.M_filesRef); i++ {
			if this.M_filesRef[i] == ref.(*File).GetUUID() {
				isExist = true
			}
		}
		this.m_filesRef = append(this.m_filesRef, ref.(*File))
	if !isExist {
		this.M_filesRef = append(this.M_filesRef, ref.(Entity).GetUUID())
	}
	}
}

/** Remove reference FilesRef **/
func (this *Project) RemoveFilesRef(ref interface{}){
	toDelete := ref.(Entity)
	filesRef_ := make([]*File, 0)
	filesRefUuid := make([]string, 0)
	for i := 0; i < len(this.m_filesRef); i++ {
		if toDelete.GetUUID() != this.m_filesRef[i].GetUUID() {
			filesRef_ = append(filesRef_, this.m_filesRef[i])
			filesRefUuid = append(filesRefUuid, this.M_filesRef[i])
		}
	}
	this.m_filesRef = filesRef_
	this.M_filesRef = filesRefUuid
}

/** Entities **/
func (this *Project) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Project) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Project) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
		}
	}
}
