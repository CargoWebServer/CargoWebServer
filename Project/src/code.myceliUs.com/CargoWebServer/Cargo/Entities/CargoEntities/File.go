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
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

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
	M_files []*File
	M_fileType FileType


	/** Associations **/
	m_parentDirPtr *File
	/** If the ref is a string and not an object **/
	M_parentDirPtr string
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
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
/** UUID **/
func (this *File) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *File) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *File) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Name **/
func (this *File) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *File) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Path **/
func (this *File) GetPath() string{
	return this.M_path
}

/** Init reference Path **/
func (this *File) SetPath(ref interface{}){
	this.NeedSave = true
	this.M_path = ref.(string)
}

/** Remove reference Path **/

/** Size **/
func (this *File) GetSize() int{
	return this.M_size
}

/** Init reference Size **/
func (this *File) SetSize(ref interface{}){
	this.NeedSave = true
	this.M_size = ref.(int)
}

/** Remove reference Size **/

/** ModeTime **/
func (this *File) GetModeTime() int64{
	return this.M_modeTime
}

/** Init reference ModeTime **/
func (this *File) SetModeTime(ref interface{}){
	this.NeedSave = true
	this.M_modeTime = ref.(int64)
}

/** Remove reference ModeTime **/

/** IsDir **/
func (this *File) IsDir() bool{
	return this.M_isDir
}

/** Init reference IsDir **/
func (this *File) SetIsDir(ref interface{}){
	this.NeedSave = true
	this.M_isDir = ref.(bool)
}

/** Remove reference IsDir **/

/** Checksum **/
func (this *File) GetChecksum() string{
	return this.M_checksum
}

/** Init reference Checksum **/
func (this *File) SetChecksum(ref interface{}){
	this.NeedSave = true
	this.M_checksum = ref.(string)
}

/** Remove reference Checksum **/

/** Data **/
func (this *File) GetData() string{
	return this.M_data
}

/** Init reference Data **/
func (this *File) SetData(ref interface{}){
	this.NeedSave = true
	this.M_data = ref.(string)
}

/** Remove reference Data **/

/** Thumbnail **/
func (this *File) GetThumbnail() string{
	return this.M_thumbnail
}

/** Init reference Thumbnail **/
func (this *File) SetThumbnail(ref interface{}){
	this.NeedSave = true
	this.M_thumbnail = ref.(string)
}

/** Remove reference Thumbnail **/

/** Mime **/
func (this *File) GetMime() string{
	return this.M_mime
}

/** Init reference Mime **/
func (this *File) SetMime(ref interface{}){
	this.NeedSave = true
	this.M_mime = ref.(string)
}

/** Remove reference Mime **/

/** Files **/
func (this *File) GetFiles() []*File{
	return this.M_files
}

/** Init reference Files **/
func (this *File) SetFiles(ref interface{}){
	this.NeedSave = true
	isExist := false
	var filess []*File
	for i:=0; i<len(this.M_files); i++ {
		if this.M_files[i].GetUUID() != ref.(Entity).GetUUID() {
			filess = append(filess, this.M_files[i])
		} else {
			isExist = true
			filess = append(filess, ref.(*File))
		}
	}
	if !isExist {
		filess = append(filess, ref.(*File))
	}
	this.M_files = filess
}

/** Remove reference Files **/
func (this *File) RemoveFiles(ref interface{}){
	toDelete := ref.(Entity)
	files_ := make([]*File, 0)
	for i := 0; i < len(this.M_files); i++ {
		if toDelete.GetUUID() != this.M_files[i].GetUUID() {
			files_ = append(files_, this.M_files[i])
		}else{
			this.NeedSave = true
		}
	}
	this.M_files = files_
}

/** FileType **/
func (this *File) GetFileType() FileType{
	return this.M_fileType
}

/** Init reference FileType **/
func (this *File) SetFileType(ref interface{}){
	this.NeedSave = true
	this.M_fileType = ref.(FileType)
}

/** Remove reference FileType **/

/** ParentDir **/
func (this *File) GetParentDirPtr() *File{
	return this.m_parentDirPtr
}

/** Init reference ParentDir **/
func (this *File) SetParentDirPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentDirPtr = ref.(string)
	}else{
		this.m_parentDirPtr = ref.(*File)
		this.M_parentDirPtr = ref.(Entity).GetUUID()
	}
}

/** Remove reference ParentDir **/
func (this *File) RemoveParentDirPtr(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_parentDirPtr!= nil {
		if toDelete.GetUUID() == this.m_parentDirPtr.GetUUID() {
			this.m_parentDirPtr = nil
			this.M_parentDirPtr = ""
		}else{
			this.NeedSave = true
		}
	}
}

/** Entities **/
func (this *File) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *File) SetEntitiesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *File) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
		}else{
			this.NeedSave = true
		}
	}
}
