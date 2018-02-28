// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
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
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

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

/** Evaluate if an entity needs to be saved. **/
func (this *File) IsNeedSave() bool{
	return this.NeedSave
}
func (this *File) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *File) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

func (this *File) GetId()string{
	return this.M_id
}

func (this *File) SetId(val string){
	this.NeedSave = this.M_id== val
	this.M_id= val
}


func (this *File) GetName()string{
	return this.M_name
}

func (this *File) SetName(val string){
	this.NeedSave = this.M_name== val
	this.M_name= val
}


func (this *File) GetPath()string{
	return this.M_path
}

func (this *File) SetPath(val string){
	this.NeedSave = this.M_path== val
	this.M_path= val
}


func (this *File) GetSize()int{
	return this.M_size
}

func (this *File) SetSize(val int){
	this.NeedSave = this.M_size== val
	this.M_size= val
}


func (this *File) GetModeTime()int64{
	return this.M_modeTime
}

func (this *File) SetModeTime(val int64){
	this.NeedSave = this.M_modeTime== val
	this.M_modeTime= val
}


func (this *File) IsDir()bool{
	return this.M_isDir
}

func (this *File) SetIsDir(val bool){
	this.NeedSave = this.M_isDir== val
	this.M_isDir= val
}


func (this *File) GetChecksum()string{
	return this.M_checksum
}

func (this *File) SetChecksum(val string){
	this.NeedSave = this.M_checksum== val
	this.M_checksum= val
}


func (this *File) GetData()string{
	return this.M_data
}

func (this *File) SetData(val string){
	this.NeedSave = this.M_data== val
	this.M_data= val
}


func (this *File) GetThumbnail()string{
	return this.M_thumbnail
}

func (this *File) SetThumbnail(val string){
	this.NeedSave = this.M_thumbnail== val
	this.M_thumbnail= val
}


func (this *File) GetMime()string{
	return this.M_mime
}

func (this *File) SetMime(val string){
	this.NeedSave = this.M_mime== val
	this.M_mime= val
}


func (this *File) GetFiles()[]*File{
	files := make([]*File, 0)
	for i := 0; i < len(this.M_files); i++ {
		entity, err := this.getEntityByUuid(this.M_files[i])
		if err == nil {
			files = append(files, entity.(*File))
		}
	}
	return files
}

func (this *File) SetFiles(val []*File){
	this.M_files= make([]string,0)
	for i:=0; i < len(val); i++{
		this.M_files=append(this.M_files, val[i].GetUuid())
	}
}

func (this *File) AppendFiles(val *File){
	for i:=0; i < len(this.M_files); i++{
		if this.M_files[i] == val.GetUuid() {
			return
		}
	}
	this.M_files = append(this.M_files, val.GetUuid())
}

func (this *File) RemoveFiles(val *File){
	files := make([]string,0)
	for i:=0; i < len(this.M_files); i++{
		if this.M_files[i] != val.GetUuid() {
			files = append(files, val.GetUuid())
		}else{
			this.NeedSave = true
		}
	}
	this.M_files = files
}


func (this *File) GetFileType()FileType{
	return this.M_fileType
}

func (this *File) SetFileType(val FileType){
	this.NeedSave = this.M_fileType== val
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
}

func (this *File) ResetEntitiesPtr(){
	this.M_entitiesPtr= ""
}

