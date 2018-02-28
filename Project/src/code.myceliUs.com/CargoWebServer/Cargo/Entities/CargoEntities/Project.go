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
	M_filesRef []string


	/** Associations **/
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
func (this *Project) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *Project) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

func (this *Project) GetId()string{
	return this.M_id
}

func (this *Project) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *Project) GetName()string{
	return this.M_name
}

func (this *Project) SetName(val string){
	this.NeedSave = this.M_name== val
	this.M_name= val
}


func (this *Project) GetFilesRef()[]*File{
	filesRef := make([]*File, 0)
	for i := 0; i < len(this.M_filesRef); i++ {
		entity, err := this.getEntityByUuid(this.M_filesRef[i])
		if err == nil {
			filesRef = append(filesRef, entity.(*File))
		}
	}
	return filesRef
}

func (this *Project) SetFilesRef(val []*File){
	this.M_filesRef= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_filesRef=append(this.M_filesRef, val[i].GetUuid())
	}
}

func (this *Project) AppendFilesRef(val *File){
	for i:=0; i < len(this.M_filesRef); i++{
		if this.M_filesRef[i] == val.GetUuid() {
			return
		}
	}
	this.M_filesRef = append(this.M_filesRef, val.GetUuid())
}

func (this *Project) RemoveFilesRef(val *File){
	filesRef := make([]string,0)
	for i:=0; i < len(this.M_filesRef); i++{
		if this.M_filesRef[i] != val.GetUuid() {
			filesRef = append(filesRef, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_filesRef = filesRef
}


func (this *Project) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *Project) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
}

func (this *Project) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

