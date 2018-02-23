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

/** Id **/
func (this *File) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *File) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Name **/
func (this *File) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *File) SetName(ref interface{}){
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Name **/

/** Path **/
func (this *File) GetPath() string{
	return this.M_path
}

/** Init reference Path **/
func (this *File) SetPath(ref interface{}){
	if this.M_path != ref.(string) {
		this.M_path = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Path **/

/** Size **/
func (this *File) GetSize() int{
	return this.M_size
}

/** Init reference Size **/
func (this *File) SetSize(ref interface{}){
	if this.M_size != ref.(int) {
		this.M_size = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference Size **/

/** ModeTime **/
func (this *File) GetModeTime() int64{
	return this.M_modeTime
}

/** Init reference ModeTime **/
func (this *File) SetModeTime(ref interface{}){
	if this.M_modeTime != ref.(int64) {
		this.M_modeTime = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference ModeTime **/

/** IsDir **/
func (this *File) IsDir() bool{
	return this.M_isDir
}

/** Init reference IsDir **/
func (this *File) SetIsDir(ref interface{}){
	if this.M_isDir != ref.(bool) {
		this.M_isDir = ref.(bool)
		this.NeedSave = true
	}
}

/** Remove reference IsDir **/

/** Checksum **/
func (this *File) GetChecksum() string{
	return this.M_checksum
}

/** Init reference Checksum **/
func (this *File) SetChecksum(ref interface{}){
	if this.M_checksum != ref.(string) {
		this.M_checksum = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Checksum **/

/** Data **/
func (this *File) GetData() string{
	return this.M_data
}

/** Init reference Data **/
func (this *File) SetData(ref interface{}){
	if this.M_data != ref.(string) {
		this.M_data = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Data **/

/** Thumbnail **/
func (this *File) GetThumbnail() string{
	return this.M_thumbnail
}

/** Init reference Thumbnail **/
func (this *File) SetThumbnail(ref interface{}){
	if this.M_thumbnail != ref.(string) {
		this.M_thumbnail = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Thumbnail **/

/** Mime **/
func (this *File) GetMime() string{
	return this.M_mime
}

/** Init reference Mime **/
func (this *File) SetMime(ref interface{}){
	if this.M_mime != ref.(string) {
		this.M_mime = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Mime **/

/** Files **/
func (this *File) GetFiles() []*File{
	return this.M_files
}

/** Init reference Files **/
func (this *File) SetFiles(ref interface{}){
	isExist := false
	var filess []*File
	for i:=0; i<len(this.M_files); i++ {
		if this.M_files[i].GetUuid() != ref.(Entity).GetUuid() {
			filess = append(filess, this.M_files[i])
		} else {
			isExist = true
			filess = append(filess, ref.(*File))
		}
	}
	if !isExist {
		filess = append(filess, ref.(*File))
		this.NeedSave = true
		this.M_files = filess
	}
}

/** Remove reference Files **/
func (this *File) RemoveFiles(ref interface{}){
	toDelete := ref.(Entity)
	files_ := make([]*File, 0)
	for i := 0; i < len(this.M_files); i++ {
		if toDelete.GetUuid() != this.M_files[i].GetUuid() {
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
	if this.M_fileType != ref.(FileType) {
		this.M_fileType = ref.(FileType)
		this.NeedSave = true
	}
}

/** Remove reference FileType **/

/** ParentDir **/
func (this *File) GetParentDirPtr() *File{
	if this.m_parentDirPtr == nil {
		entity, err := this.getEntityByUuid(this.M_parentDirPtr)
		if err == nil {
			this.m_parentDirPtr = entity.(*File)
		}
	}
	return this.m_parentDirPtr
}
func (this *File) GetParentDirPtrStr() string{
	return this.M_parentDirPtr
}

/** Init reference ParentDir **/
func (this *File) SetParentDirPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_parentDirPtr != ref.(string) {
			this.M_parentDirPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_parentDirPtr != ref.(Entity).GetUuid() {
			this.M_parentDirPtr = ref.(Entity).GetUuid()
			this.NeedSave = true
		}
		this.m_parentDirPtr = ref.(*File)
	}
}

/** Remove reference ParentDir **/
func (this *File) RemoveParentDirPtr(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_parentDirPtr!= nil {
		if toDelete.GetUuid() == this.m_parentDirPtr.GetUuid() {
			this.m_parentDirPtr = nil
			this.M_parentDirPtr = ""
			this.NeedSave = true
		}
	}
}

/** Entities **/
func (this *File) GetEntitiesPtr() *Entities{
	if this.m_entitiesPtr == nil {
		entity, err := this.getEntityByUuid(this.M_entitiesPtr)
		if err == nil {
			this.m_entitiesPtr = entity.(*Entities)
		}
	}
	return this.m_entitiesPtr
}
func (this *File) GetEntitiesPtrStr() string{
	return this.M_entitiesPtr
}

/** Init reference Entities **/
func (this *File) SetEntitiesPtr(ref interface{}){
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
func (this *File) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUuid() == this.m_entitiesPtr.GetUuid() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
