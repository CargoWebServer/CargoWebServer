// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
	"strings"
)

type File struct{

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

	/** members of File **/
	M_name string
	M_path string
	M_size int
	M_modeTime int64
	M_isDir bool
	M_checksum string
	M_data string
	M_thumbnail string
	M_mime string
	M_files []string
	M_fileType FileType


	/** Associations **/
	M_parentDirPtr string
	M_entitiesPtr string
}

/** Xml parser for File **/
type XsdFile struct {
	XMLName xml.Name	`xml:"filesRef"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	M_files	[]*XsdFile	`xml:"files,omitempty"`
	M_name	string	`xml:"name,attr"`
	M_path	string	`xml:"path,attr"`
	M_size	int	`xml:"size,attr"`
	M_modeTime	int64	`xml:"modeTime,attr"`
	M_isDir	bool	`xml:"isDir,attr"`
	M_checksum	string	`xml:"checksum,attr"`
	M_data	string	`xml:"data,attr"`
	M_thumbnail	string	`xml:"thumbnail,attr"`
	M_mime	string	`xml:"mime,attr"`
	M_fileType	string	`xml:"fileType,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *File) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *File) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *File) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *File) GetTypeName() string{
	this.TYPENAME = "CargoEntities.File"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *File) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *File) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *File) GetParentLnk() string{
	return this.ParentLnk
}
func (this *File) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *File) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *File) GetChilds() []interface{}{
	var childs []interface{}
	var child interface{}
	var err error
	for i:=0; i < len(this.M_files); i++ {
		child, err = this.getEntityByUuid( this.M_files[i])
		if err == nil {
			childs = append( childs, child)
		}
	}
	return childs
}
/** Return the list of all childs uuid **/
func (this *File) GetChildsUuid() []string{
	var childs []string
	childs = append( childs, this.M_files...)
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *File) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *File) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *File) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *File) GetId()string{
	return this.M_id
}

func (this *File) SetId(val string){
	this.M_id= val
}




func (this *File) GetName()string{
	return this.M_name
}

func (this *File) SetName(val string){
	this.M_name= val
}




func (this *File) GetPath()string{
	return this.M_path
}

func (this *File) SetPath(val string){
	this.M_path= val
}




func (this *File) GetSize()int{
	return this.M_size
}

func (this *File) SetSize(val int){
	this.M_size= val
}




func (this *File) GetModeTime()int64{
	return this.M_modeTime
}

func (this *File) SetModeTime(val int64){
	this.M_modeTime= val
}




func (this *File) IsDir()bool{
	return this.M_isDir
}

func (this *File) SetIsDir(val bool){
	this.M_isDir= val
}




func (this *File) GetChecksum()string{
	return this.M_checksum
}

func (this *File) SetChecksum(val string){
	this.M_checksum= val
}




func (this *File) GetData()string{
	return this.M_data
}

func (this *File) SetData(val string){
	this.M_data= val
}




func (this *File) GetThumbnail()string{
	return this.M_thumbnail
}

func (this *File) SetThumbnail(val string){
	this.M_thumbnail= val
}




func (this *File) GetMime()string{
	return this.M_mime
}

func (this *File) SetMime(val string){
	this.M_mime= val
}




func (this *File) GetFiles()[]*File{
	values := make([]*File, 0)
	for i := 0; i < len(this.M_files); i++ {
		entity, err := this.getEntityByUuid(this.M_files[i])
		if err == nil {
			values = append( values, entity.(*File))
		}
	}
	return values
}

func (this *File) SetFiles(val []*File){
	this.M_files= make([]string,0)
	for i:=0; i < len(val); i++{
		if len(val[i].GetParentUuid()) > 0  &&  len(val[i].GetParentLnk()) > 0 {
			parent, _ := this.getEntityByUuid(val[i].GetParentUuid())
			if parent != nil {
				removeMethode := strings.Replace(val[i].GetParentLnk(), "M_", "", -1)
				removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
				params := make([]interface{}, 1)
				params[0] = val
				Utility.CallMethod(parent, removeMethode, params)
				this.setEntity(parent)
			}
		}
		val[i].SetParentUuid(this.GetUuid())
		val[i].SetParentLnk("M_files")
		this.M_files=append(this.M_files, val[i].GetUuid())
		this.setEntity(val[i])
	}
	this.setEntity(this)
}


func (this *File) AppendFiles(val *File){
	for i:=0; i < len(this.M_files); i++{
		if this.M_files[i] == val.GetUuid() {
			return
		}
	}
	if len(val.GetParentUuid()) > 0 &&  len(val.GetParentLnk()) > 0 {
		parent, _ := this.getEntityByUuid(val.GetParentUuid())
		if parent != nil {
			removeMethode := strings.Replace(val.GetParentLnk(), "M_", "", -1)
			removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
			params := make([]interface{}, 1)
			params[0] = val
			Utility.CallMethod(parent, removeMethode, params)
			this.setEntity(parent)
		}
	}
	val.SetParentUuid(this.GetUuid())
	val.SetParentLnk("M_files")
  this.setEntity(val)
	this.M_files = append(this.M_files, val.GetUuid())
	this.setEntity(this)
}

func (this *File) RemoveFiles(val *File){
	values := make([]string,0)
	for i:=0; i < len(this.M_files); i++{
		if this.M_files[i] != val.GetUuid() {
			values = append(values, this.M_files[i])
		}
	}
	this.M_files = values
	this.setEntity(this)
}


func (this *File) GetFileType()FileType{
	return this.M_fileType
}

func (this *File) SetFileType(val FileType){
	this.M_fileType= val
}


func (this *File) ResetFileType(){
	this.M_fileType= 0
}


func (this *File) GetParentDirPtr()*File{
	entity, err := this.getEntityByUuid(this.M_parentDirPtr)
	if err == nil {
		return entity.(*File)
	}
	return nil
}

func (this *File) SetParentDirPtr(val *File){
	this.M_parentDirPtr= val.GetUuid()
	this.setEntity(this)
}


func (this *File) ResetParentDirPtr(){
	this.M_parentDirPtr= ""
}


func (this *File) GetEntitiesPtr()*Entities{
	entity, err := this.getEntityByUuid(this.M_entitiesPtr)
	if err == nil {
		return entity.(*Entities)
	}
	return nil
}

func (this *File) SetEntitiesPtr(val *Entities){
	this.M_entitiesPtr= val.GetUuid()
	this.setEntity(this)
}


func (this *File) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

