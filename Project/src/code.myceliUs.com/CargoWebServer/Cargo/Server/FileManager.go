package Server

import (
	"bufio"
	"bytes"
	base64 "encoding/base64"
	"encoding/csv"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"github.com/nfnt/resize"

	"code.myceliUs.com/Utility"
	"github.com/polds/imgbase64"
	"github.com/tealeg/xlsx"
)

/**
 * That class will contain the list of mime type.
 */
type MimeType struct {
	Id             string
	Name           string
	FileExtensions []string
	Info           string
	Thumbnail      string
}

/**
 * This class implement functionality to manage account
 */
type FileManager struct {
	// This represent the root path of the server...
	root        string
	mimeTypeMap map[string]*MimeType
}

var fileManager *FileManager

func (this *Server) GetFileManager() *FileManager {
	if fileManager == nil {
		fileManager = newFileManager()
	}
	return fileManager
}

/**
 * Create a new file manager...
 */
func newFileManager() *FileManager {

	fileManager := new(FileManager)

	// Set the root to the application root path...
	fileManager.root = GetServer().GetConfigurationManager().GetApplicationDirectoryPath()
	fileManager.loadMimeType()

	return fileManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

func (this *FileManager) initialize() {
	// register service avalaible action here.
	log.Println("--> initialyze ConfigurationManager")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)

}

func (this *FileManager) getId() string {
	return "FileManager"
}

func (this *FileManager) start() {
	log.Println("--> Start FileManager")
	// Now I will synchronize files...
	rootDir := this.synchronize(this.root)
	rootEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntityFromObject(rootDir)
	rootEntity.SaveEntity()
}

func (this *FileManager) stop() {
	log.Println("--> Stop FileManager")
}

/**
 * That function will be call when the server start to synchronize file on the
 * disck with their related information in the db.
 */
func (this *FileManager) synchronize(filePath string) *CargoEntities.File {

	// That will contain the list of file that no more exist in a dir...
	var toDelete map[string]*CargoEntities.File

	// Keep only the part past the root...
	filePath_ := strings.Replace(filePath, this.root, "", -1)

	isRoot := len(filePath_) == 0
	if isRoot {
		filePath_ = "CARGOROOT"
	}

	// Here I will create the directory if it not exist...
	dirId := Utility.CreateSha1Key([]byte(filePath_))
	dirEntity, err := this.getFileById(dirId)

	if err != nil {
		log.Println("Dir ", filePath_, "not found!")
		if isRoot {
			dirEntity, _ = this.createDir("", "", "")
		} else {
			lastIndex := strings.LastIndex(filePath_, "/")
			dirName := ""
			dirPath := ""
			if lastIndex > 0 {
				dirName = filePath_[lastIndex+1:]
				dirPath = filePath_[:lastIndex]
			} else if lastIndex == 0 {
				dirName = filePath_[lastIndex+1:]
			}
			if !strings.HasPrefix(dirName, ".") {
				dirEntity, _ = this.createDir(dirName, dirPath, "")
			}
		}
	} else {
		// The directory already exist...
		toDelete = make(map[string]*CargoEntities.File, 0)

		// I will put all the file in the delete map and remove it
		// one by one latter.
		for i := 0; i < len(dirEntity.GetFiles()); i++ {
			// I will put all file here by default...
			_file := dirEntity.GetFiles()[i]
			toDelete[_file.GetId()] = _file
		}
	}

	// So I will get the the file from the root...
	files, _ := ioutil.ReadDir(filePath)
	if isRoot {
		// Reset the value of filePath_ to nothing
		filePath_ = ""
	}

	for _, f := range files {

		// Create the file information here...
		filePath__ := filePath_ + "/" + f.Name()
		fileId := Utility.CreateSha1Key([]byte(filePath__))
		//log.Println("Get file ", filePath__)
		fileEntity, err := this.getFileById(fileId)

		// Remove from file to delete...
		delete(toDelete, fileId)

		if err != nil {
			// here I will create the new entity...
			if !f.IsDir() {
				// Now I will open the file and create the entry in the DB.
				filedata, _ := ioutil.ReadFile(this.root + filePath__)
				if !strings.HasPrefix(f.Name(), ".") {
					file, err := this.createFile(dirEntity, f.Name(), filePath_, filedata, "", 128, 128, false)
					if err == nil {
						dirEntity.SetFiles(file)
					} else {
						log.Panicln("--------------> fail to create file ", filePath__)
					}
				}
			} else {
				// make a recursion...
				if !strings.HasPrefix(f.Name(), ".") {
					subDir := this.synchronize(this.root + filePath__)
					if subDir != nil {
						dirEntity.SetFiles(subDir)
					}
				}
			}
		} else {
			// I will test the checksum to see if the file has change...
			if !f.IsDir() {
				// Update the file checksum and save it if the file has change...
				if !strings.HasPrefix(f.Name(), ".") {
					filedata, _ := ioutil.ReadFile(this.root + filePath__)
					this.saveFile(fileEntity.UUID, filedata, "", 128, 128, false)
					dirEntity.SetFiles(fileEntity)
				}
			} else {
				// make a recursion...
				if !strings.HasPrefix(f.Name(), ".") {
					subDir := this.synchronize(this.root + filePath__)
					if subDir != nil {
						dirEntity.SetFiles(subDir)
					}
				}
			}
		}
	}

	// I will now remove the remaining files in the map...
	for _, fileToDelete := range toDelete {
		// Delete the associated entity...
		this.deleteFile(fileToDelete.UUID)
		log.Println("Delete file: ", fileToDelete.GetPath()+"/"+fileToDelete.GetName())
	}

	return dirEntity
}

/**
 * Create a new directory and it's associated file entity on the server
 */
// TODO add the logic for db files...
func (this *FileManager) createDir(dirName string, dirPath string, sessionId string) (*CargoEntities.File, *CargoEntities.Error) {

	// Return the dir entity if it already exist.
	dirId := Utility.CreateSha1Key([]byte(dirPath + "/" + dirName))
	dirUuid := CargoEntitiesFileExists(dirId)
	if len(dirUuid) > 0 {
		dirEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(dirUuid, false)
		if errObj == nil {
			return dirEntity.GetObject().(*CargoEntities.File), nil
		}
		return nil, errObj
	}

	// If the dir is the root dir
	isRoot := len(dirName) == 0 && len(dirPath) == 0

	var dirPath_ string
	var parentDirPath string

	if isRoot {
		dirPath_ = "CARGOROOT"
	} else if len(dirPath) == 0 {
		// The new directory is at the root.
		dirPath_ = "/" + dirName
		parentDirPath = "CARGOROOT"
	} else {
		// Not the root and not at root.
		dirPath_ = dirPath + "/" + dirName
		parentDirPath = dirPath
	}

	// Create the directory if it doesn't already exist.
	if !isRoot {
		// TODO Throw an error if err != nil ?
		if _, err := os.Stat(this.root + dirPath_); os.IsNotExist(err) {
			// path/to/whatever does not exist
			log.Println("Creating directory ", this.root+dirPath_)
			os.MkdirAll(this.root+dirPath_, 0777)
		}
	}

	// The parent directory

	// Retreive the perentDirEntity, put the file in it and save the parent
	var parentDir *CargoEntities.File
	var parentDirEntity *CargoEntities_FileEntity
	parentDirEntityId := Utility.CreateSha1Key([]byte(parentDirPath))
	parentDirUuid := CargoEntitiesFileExists(parentDirEntityId)

	if len(parentDirUuid) == 0 {
		// Here I will create the parent directory...
		index := strings.LastIndex(parentDirPath, "/")
		var cargoError *CargoEntities.Error

		if index > 0 {
			parentDir, cargoError = this.createDir(parentDirPath[index+1:], parentDirPath[0:index], sessionId)
		}

		// Create the error message
		if cargoError != nil {
			return nil, cargoError
		}
	}

	// The directory entity.
	dir := new(CargoEntities.File)
	dir.SetId(Utility.CreateSha1Key([]byte(dirPath_)))
	dir.SetIsDir(true)
	dir.SetName(dirName)
	dir.SetPath(dirPath)
	dir.SetFileType(CargoEntities.FileType_DiskFile)

	// Set the cargo entities object.
	entities := GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities)
	dir.SetEntitiesPtr(entities)

	if parentDir != nil {
		dir.SetParentDirPtr(parentDir)
		parentDir.SetFiles(dir)

		// Get the parent dir information.
		parentDirUuid = parentDir.GetUUID()
		parentDirEntity = GetServer().GetEntityManager().NewCargoEntitiesFileEntity("", parentDirUuid, nil)
		parentDir = parentDirEntity.GetObject().(*CargoEntities.File)

		// That will set the uuid
		GetServer().GetEntityManager().NewCargoEntitiesFileEntity(parentDirUuid, "", dir)

		// Save the parent directory. This will also save the directory.
		parentDirEntity.SaveEntity()
	} else {
		dirEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity("", "", dir)
		dirEntity.SaveEntity()
	}
	return dir, nil

}

/**
 * Create a file.
 */
func (this *FileManager) createFile(parentDir *CargoEntities.File, filename string, filepath string, filedata []byte, sessionId string, thumbnailMaxHeight int, thumbnailMaxWidth int, dbFile bool) (*CargoEntities.File, *CargoEntities.Error) {

	var file *CargoEntities.File
	var f *os.File
	var err error
	var checksum string

	// Use to determine if the file already exist or not.
	isNew := true

	// The user wants to save the data into a physical file.
	// File id's
	fileId := Utility.CreateSha1Key([]byte(filepath + "/" + filename))
	fileUuid := CargoEntitiesFileExists(fileId)

	parentDirEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(parentDir.GetUUID(), false)
	if errObj != nil {
		return nil, errObj
	}

	// I will retreive the file.
	if len(fileUuid) > 0 {
		isNew = false
		entity, cargoError := GetServer().GetEntityManager().getEntityByUuid(fileUuid, false)
		if cargoError != nil {
			return nil, cargoError
		}
		file = entity.GetObject().(*CargoEntities.File)
		log.Println("Get entity for file: ", filepath+"/"+filename)
	} else {
		file = new(CargoEntities.File)
		file.SetId(fileId)
		// Set the basic information.
		// Set the file uuid.
		GetServer().GetEntityManager().NewCargoEntitiesFileEntity(parentDirEntity.GetUuid(), "", file)
		log.Println("Create entity for file: ", filepath+"/"+filename)
	}

	// Disk file.
	if !dbFile {
		// Here I will get the parent put the file in it and save it.
		parentDir := parentDirEntity.GetObject().(*CargoEntities.File)
		file.SetParentDirPtr(parentDir)
		parentDir.SetFiles(file)

		// Write the file data. Try to decode it if it is encoded, and decode it if it is not.
		if _, err := os.Stat(this.root + filepath + "/" + filename); err != nil {
			// TODO Throw an error if err != nil ?
			if os.IsNotExist(err) {
				ioutil.WriteFile(this.root+filepath+"/"+filename, filedata, 0644)
			}
		}

		file.SetFileType(CargoEntities.FileType_DiskFile)
		f_, err := ioutil.TempFile("", "_create_file_")
		if err != nil {
			// Create the error message
			cargoError := NewError(Utility.FileLine(), FILE_OPEN_ERROR, SERVER_ERROR_CODE, errors.New("Failed to open tmp file for '"+filepath+"/"+filename+"'. "))
			return nil, cargoError
		}

		f_.Write(filedata)
		defer f_.Close()
		checksum_ := Utility.CreateFileChecksum(f_)

		// Open it
		f, err = os.Open(this.root + filepath + "/" + filename)
		if err != nil {
			// Create the error message
			cargoError := NewError(Utility.FileLine(), FILE_OPEN_ERROR, SERVER_ERROR_CODE, errors.New("Failed to open '"+this.root+filepath+"/"+filename+"'. "))
			return nil, cargoError
		}
		checksum = Utility.CreateFileChecksum(f)

		// Update the file content.
		if checksum != checksum_ {
			err = ioutil.WriteFile(this.root+filepath+"/"+filename, filedata, 0644)
			if err != nil {
				// Create the error message
				cargoError := NewError(Utility.FileLine(), FILE_MANAGER_ERROR, SERVER_ERROR_CODE, err)
				return nil, cargoError
			}
			// Reload the file
			f.Close()
			f, _ = os.Open(this.root + filepath + "/" + filename)
		}
	} else {

		// The file data will be saved in the database without physical file.
		// The id will be a uuid.
		file.SetFileType(CargoEntities.FileType_DbFile)
		f, err = ioutil.TempFile("", "_cargo_tmp_file_")
		if err != nil {
			// Create the error message
			cargoError := NewError(Utility.FileLine(), FILE_OPEN_ERROR, SERVER_ERROR_CODE, errors.New("Failed to open _cargo_tmp_file_ for file '"+filename+"'. "))
			return nil, cargoError

		}
		f.Write(filedata)

		// Create the checksum
		checksum = Utility.CreateFileChecksum(f)

		// Keep the data into the data base as a base 64 string
		file.SetData(base64.StdEncoding.EncodeToString(filedata))
	}

	// Set general information.
	file.SetPath(filepath)
	file.SetName(filename)
	file.SetChecksum(checksum)

	// Set the entity object.
	entities := GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities)
	file.SetEntitiesPtr(entities)

	var thumbnail string
	// Set the Thumnail
	fileName_ := strings.ToLower(filename)
	if strings.HasSuffix(fileName_, ".png") || strings.HasSuffix(fileName_, ".gif") || strings.HasSuffix(fileName_, ".bmp") || strings.HasSuffix(fileName_, ".jpeg") || strings.HasSuffix(fileName_, ".jpg") || strings.HasSuffix(fileName_, ".tiff") {
		thumbnail = this.createThumbnail(f, thumbnailMaxHeight, thumbnailMaxWidth)
	}

	// Mime type information.
	if strings.Index(fileName_, ".") > 0 {
		fileExtension := fileName_[strings.Index(fileName_, "."):]
		if len(fileExtension) > 0 {
			// Get the mimeType.
			mimeType := this.mimeTypeMap[strings.ToLower(fileExtension)]
			if mimeType != nil {
				if len(thumbnail) == 0 {
					thumbnail = mimeType.Thumbnail
				}
				file.SetMime(mimeType.Id)
			}
		}
	}

	file.SetThumbnail(thumbnail)

	// Other info.
	info, err := f.Stat()
	if err != nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), GET_FILE_STAT_ERROR, SERVER_ERROR_CODE, errors.New("Failed to get file '"+filename+"' stats. "))
		return nil, cargoError
	}

	// Set other information.
	file.SetSize(int(info.Size()))
	file.SetModeTime(info.ModTime().Unix())
	file.SetIsDir(info.IsDir())

	defer f.Close()

	if err == nil {
		eventData := make([]*MessageData, 1)
		fileInfo := new(MessageData)
		fileInfo.Name = "fileInfo"
		if !isNew {
			file.SetData(base64.StdEncoding.EncodeToString(filedata))
		}
		fileInfo.Value = file
		eventData[0] = fileInfo
		var evt *Event
		if isNew {
			// New file create
			evt, _ = NewEvent(NewFileEvent, FileEvent, eventData)
		} else {
			// Existing file was update
			evt, _ = NewEvent(UpdateFileEvent, FileEvent, eventData)
		}
		GetServer().GetEventManager().BroadcastEvent(evt)
	}

	// Save the file.
	parentDirEntity.SaveEntity()

	return file, nil
}

/**
 * Return a file with a given uuid, or id.
 */
func (this *FileManager) getFileById(id string) (*CargoEntities.File, *CargoEntities.Error) {
	ids := []interface{}{id}
	fileEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.File", ids, false)

	if errObj == nil {
		// Initialize it's content.
		fileEntity.InitEntity(fileEntity.GetUuid(), false)
		file := fileEntity.GetObject().(*CargoEntities.File)
		return file, nil
	}

	return nil, errObj
}

/**
 * Save the content of the file...
 */
func (this *FileManager) saveFile(uuid string, filedata []byte, sessionId string, thumbnailMaxHeight int, thumbnailMaxWidth int, dbFile bool) error {

	// Create a temporary file object from the data...
	f, err := ioutil.TempFile(os.TempDir(), uuid)
	if err != nil {
		log.Println("Fail to open _cargo_tmp_file_ for file ", uuid, " ", err)
		return err
	}

	f.Write(filedata)

	checksum := Utility.CreateFileChecksum(f)

	defer os.Remove(f.Name())

	// I will retreive the file, the uuid must exist in that case.
	fileEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity("", uuid, nil)
	file := fileEntity.GetObject().(*CargoEntities.File)

	if file.GetChecksum() != checksum {
		if !dbFile {
			// Save the data to the disck...
			filename := file.GetName()
			filepath := file.GetPath()
			log.Println("--> write file ln 617", this.root+filepath+"/"+filename)
			err = ioutil.WriteFile(this.root+"/"+filepath+"/"+filename, filedata, 0644)
			if err != nil {
				return err
			}
		} else {
			// Set the new data...
			file.SetData(filedata)
		}
		file.SetChecksum(checksum)
		fileEntity.SaveEntity()
	} else {
		file.NeedSave = false // Not need to be save...
	}

	if err == nil {
		eventData := make([]*MessageData, 1)

		fileInfo := new(MessageData)
		fileInfo.Name = "fileInfo"
		fileInfo.Value = file
		eventData[0] = fileInfo

		evt, _ := NewEvent(UpdateFileEvent, FileEvent, eventData)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}

	return nil
}

/**
 * Delete a file with a given uuid
 */
func (this *FileManager) deleteFile(uuid string) error {
	fileEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity("", uuid, nil)
	fileEntity.InitEntity(uuid, false)
	file := fileEntity.GetObject().(*CargoEntities.File)

	if file.GetFileType() == CargoEntities.FileType_DiskFile {
		filePath := this.root + file.GetPath() + "/" + file.GetName()

		if !file.IsDir() {
			// Remove the file from the disck
			os.Remove(filePath) // The file can be already remove...
		} else {
			Utility.RemoveContents(filePath)
		}
	}

	if file.GetParentDirPtr() != nil {
		file.GetParentDirPtr().RemoveFiles(file)
		parentDirEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity("", file.GetParentDirPtr().UUID, nil)
		parentDirEntity.SaveEntity()
	}

	// Here i will remove the entity...
	fileEntity.DeleteEntity()

	eventData := make([]*MessageData, 1)
	fileInfo := new(MessageData)
	fileInfo.Name = "fileInfo"
	fileInfo.Value = file
	eventData[0] = fileInfo

	evt, _ := NewEvent(DeleteFileEvent, FileEvent, eventData)
	GetServer().GetEventManager().BroadcastEvent(evt)

	return nil
}

/**
 * That function open a file and return it to the user who
 * asked for.
 * TODO add security on the file...
 */
func (this *FileManager) openFile(fileId string, sessionId string) (*CargoEntities.File, *CargoEntities.Error) {
	file, errObj := this.getFileById(fileId)

	// Here the file was not found.
	if errObj != nil {
		return nil, NewError(Utility.FileLine(), FILE_NOT_FOUND_ERROR, SERVER_ERROR_CODE, errors.New("The file '"+fileId+"' was not found."))
	}

	// I will return the file
	return file, nil
}

/**
 * That function load information about file type.
 * The informations is contain in the file mimeType.csv of the Data
 * directory. The associated icon are store in the Data/MimeTypeIcon
 */
func (this *FileManager) loadMimeType() {
	this.mimeTypeMap = make(map[string]*MimeType, 0)
	mimeTypeFilePath := GetServer().GetConfigurationManager().GetDataPath() + "/mimeType.csv"
	mimeTypeFile, _ := os.Open(mimeTypeFilePath)
	csvReader := csv.NewReader(bufio.NewReader(mimeTypeFile))
	for {
		record, err := csvReader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		mime := new(MimeType)
		for i := 0; i < len(record); i++ {
			if i == 0 {
				mime.Name = record[i]
			} else if i == 1 {
				mime.Id = record[i]
			} else if i == 2 {
				filesExtensions := record[i]
				if strings.HasPrefix(filesExtensions, "\"") && strings.HasSuffix(filesExtensions, "\"") {
					filesExtensions = filesExtensions[1 : len(filesExtensions)-1]
				}
				mime.FileExtensions = strings.Split(filesExtensions, ",")
				for j := 0; j < len(mime.FileExtensions); j++ {
					mime.FileExtensions[j] = strings.TrimSpace(mime.FileExtensions[j])
				}
			} else if i == 3 {
				mime.Info = record[i]
			}
		}

		// Now the icon...
		// The icon has the name of the file extension and is a png...
		ext := mime.FileExtensions[0]
		if strings.HasPrefix(ext, ".") {
			// remove the leading .
			ext = ext[1:]
		}

		file, err := os.Open(GetServer().GetConfigurationManager().GetDataPath() + "/MimeTypeIcon/" + ext + ".png")
		if err != nil {
			// The file extension has no image so i will get the unknow image...
			file, _ = os.Open(GetServer().GetConfigurationManager().GetDataPath() + "/MimeTypeIcon/unknown.png")
		}

		// Set the image as thumbnail...
		if file != nil {
			m, _ := png.Decode(file)
			var buf bytes.Buffer
			png.Encode(&buf, m)
			mime.Thumbnail = imgbase64.FromBuffer(buf)
		}

		// Keep in map by file extension...
		for i := 0; i < len(mime.FileExtensions); i++ {
			this.mimeTypeMap[mime.FileExtensions[i]] = mime
		}
	}
}

/**
 * Create a thumbnail...
 */
func (this *FileManager) createThumbnail(file *os.File, thumbnailMaxHeight int, thumbnailMaxWidth int) string {
	// Set the buffer pointer back to the begening of the file...
	file.Seek(0, 0)
	var originalImg image.Image
	var format string
	var err error

	if strings.HasSuffix(file.Name(), ".png") || strings.HasSuffix(file.Name(), ".PNG") {
		originalImg, err = png.Decode(file)
	} else if strings.HasSuffix(file.Name(), ".jpeg") || strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".JPEG") || strings.HasSuffix(file.Name(), ".JPG") {
		originalImg, err = jpeg.Decode(file)
	} else if strings.HasSuffix(file.Name(), ".gif") || strings.HasSuffix(file.Name(), ".GIF") {
		originalImg, err = gif.Decode(file)
	} else {
		return ""
	}

	if err != nil {
		log.Println("File ", file.Name(), " Format is ", format, " error: ", err)
		return ""
	}

	// I will get the ratio for the new image size to respect the scale.
	hRatio := thumbnailMaxHeight / originalImg.Bounds().Size().Y
	wRatio := thumbnailMaxWidth / originalImg.Bounds().Size().X

	var h int
	var w int

	// First I will try with the height
	if hRatio*originalImg.Bounds().Size().Y < thumbnailMaxWidth {
		h = thumbnailMaxHeight
		w = hRatio * originalImg.Bounds().Size().Y
	} else {
		// So here i will use it width
		h = wRatio * thumbnailMaxHeight
		w = thumbnailMaxWidth
	}

	// do not zoom...
	if hRatio > 1 {
		h = originalImg.Bounds().Size().Y
	}

	if wRatio > 1 {
		w = originalImg.Bounds().Size().X
	}

	// Now I will calculate the image size...
	img := resize.Resize(uint(h), uint(w), originalImg, resize.Lanczos3)
	var buf bytes.Buffer
	jpeg.Encode(&buf, img, &jpeg.Options{jpeg.DefaultQuality})

	// Now I will save the buffer containt to the thumbnail...
	thumbnail := imgbase64.FromBuffer(buf)
	file.Seek(0, 0) // Set the reader back to the begenin of the file...
	return thumbnail
}

func (this *FileManager) createDbFile(id string, name string, mimeType string, data string) *CargoEntities.File {
	dbFile := new(CargoEntities.File)
	dbFile.M_fileType = CargoEntities.FileType_DbFile
	dbFile.M_data = base64.StdEncoding.EncodeToString([]byte(data))
	dbFile.M_isDir = false
	dbFile.TYPENAME = "CargoEntities.File"
	dbFile.M_id = id
	dbFile.M_name = name
	dbFile.M_mime = mimeType
	entities := GetServer().GetEntityManager().getCargoEntities()

	// Create the file.
	GetServer().GetEntityManager().createEntity(entities.GetUuid(), "M_entities", "CargoEntities.File", id, dbFile)

	return dbFile
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

// @api 1.0
// Event handler function.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
//FileManager.prototype.onEvent = function (evt) {
//    EventHub.prototype.onEvent.call(this, evt)
//}
func (this *FileManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

// @api 1.0
// Read the content of a text file and return it.
// @param {string} path The path on the sever relative to the sever root.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {string} The text content of the file.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) ReadTextFile(filePath string, messageId string, sessionId string) string {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return ""
	}

	b, err := ioutil.ReadFile(filePath) // just pass the file name
	if err != nil {
		errObj := NewError(Utility.FileLine(), FILE_READ_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return ""
	}
	return string(b)
}

// @api 1.0
// Read the content of a comma separated values file (CSV)
// @param {string} path The path on the sever relative to the sever root.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {[][]} A tow dimensionnal array with values string
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) ReadCsvFile(filePath string, messageId string, sessionId string) [][]string {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	csvfile, err := os.Open(filePath)
	if err != nil {
		errObj := NewError(Utility.FileLine(), FILE_READ_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		errObj := NewError(Utility.FileLine(), FILE_READ_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	return rawCSVdata
}

// @api 1.0
// Remove a file (not a file entity) from the server at a given path.
// @param {string} path The path on the sever relative to the sever root.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) RemoveFile(filePath string, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	if !strings.HasPrefix(filePath, "/") {
		filePath = "/" + filePath
	}
	filePath = this.root + filePath

	err := os.Remove(filePath)
	if err != nil {
		errObj := NewError(Utility.FileLine(), FILE_DELETE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

// @api 1.0
// Create a new directory on the server.
// @param {string} dirName The name of the new directory.
// @param {string} dirPath The path of the parent of the new directory.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.File} The created directory entity.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) CreateDir(dirName string, dirPath string, messageId string, sessionId string) *CargoEntities.File {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	dir, errObj := this.createDir(dirName, dirPath, sessionId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	return dir
}

// @api 1.0
// Dowload a file (not entity) from the sever.
// @param {string} filepath The path of the directory where to create the file.
// @param {string} filename The name of the file to create.
// @param {*Server.MimeType} mimeType The file mime type.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.File} The created file entity.
// @scope {public}
// @param {callback} progressCallback The function is call when chunk of response is received.
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
//FileManager.prototype.downloadFile = function (path, fileName, mimeType, progressCallback, successCallback, errorCallback, caller) {
//    var xhr = new XMLHttpRequest();
//    xhr.open('GET', path + '/' + fileName, true);
//    xhr.responseType = 'blob';
//    xhr.onload = function (successCallback) {
//        return function (e) {
//            if (this.status == 200) {
//                // Note: .response instead of .responseText
//                var blob = new Blob([this.response], { type: mimeType });
//                // I will read the file as data url...
//                var reader = new FileReader();
//                reader.onload = function (successCallback, caller) {
//                    return function (e) {
//                        var dataURL = e.target.result;
//                        // return the success callback with the result.
//                        successCallback(dataURL, caller)
//                    }
//                } (successCallback, caller)
//                reader.readAsDataURL(blob);
//            }
//        }
//    } (successCallback, caller)
//    xhr.onprogress = function (progressCallback, caller) {
//        return function (e) {
//            progressCallback(e.loaded, e.total, caller)
//        }
//    } (progressCallback, caller)
//    xhr.send();
//}
func (this *FileManager) DownloadFile(path string, fileName string, mimeType *MimeType, messageId string, sessionId string) {
	/** empty funtion **/
}

// @api 1.0
// Create a new file on the server.
// @param {string} filename The name of the file to create.
// @param {string} filepath The path of the directory where to create the file.
// @param {string} filedata The data of the file.
// @param {int} thumbnailMaxHeight The maximum height size of the thumbnail associated with the file (keep the ratio).
// @param {int} thumbnailMaxWidth The maximum width size of the thumbnail associated with the file (keep the ratio).
// @param {bool} dbFile If it set to true the file will be save on the server local object store, otherwize a file on disck will be created.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.File} The created file entity.
// @scope {public}
// @param {callback} progressCallback The function is call when chunk of response is received.
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
//FileManager.prototype.createFile = function (filename, filepath, filedata, thumbnailMaxHeight, thumbnailMaxWidth, dbFile, successCallback, progressCallback, errorCallback, caller) {
//    // server is the client side singleton.
//    var params = []
//    // The file data (filedata) will be upload with the http protocol...
//    params.push(createRpcData(filename, "STRING", "filename"))
//    params.push(createRpcData(filepath, "STRING", "filepath"))
//    params.push(createRpcData(thumbnailMaxHeight, "INTEGER", "thumbnailMaxHeight"))
//    params.push(createRpcData(thumbnailMaxWidth, "INTEGER", "thumbnailMaxWidth"))
//    params.push(createRpcData(dbFile, "BOOLEAN", "dbFile"))
//    // Here I will create a new data form...
//    var formData = new FormData()
//    formData.append("multiplefiles", filedata, filename)
//    // Use the post function to upload the file to the server.
//    var xhr = new XMLHttpRequest()
//    xhr.open('POST', '/uploads', true)
//    // In case of error or success...
//    xhr.onload = function (params, xhr) {
//        return function (e) {
//            if (xhr.readyState === 4) {
//                if (xhr.status === 200) {
//                    console.log(xhr.responseText);
//                    // Here I will create the file...
//                    server.executeJsFunction(
//                        "FileManagerCreateFile", // The function to execute remotely on server
//                        params, // The parameters to pass to that function
//                        function (index, total, caller) { // The progress callback
//                            // Keep track of the file transfert.
//                            caller.progressCallback(index, total, caller.caller)
//                        },
//                        function (result, caller) {
//                            caller.successCallback(result[0], caller.caller)
//                        },
//                        function (errMsg, caller) {
//                            // display the message in the console.
//                            console.log(errMsg)
//                            // call the immediate error callback.
//                            caller.errorCallback(errMsg, caller.caller)
//                            // dispatch the message.
//                            server.errorManager.onError(errMsg)
//                        }, // Error callback
//                        { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback } // The caller
//                    )
//                } else {
//                    console.error(xhr.statusText);
//                }
//            }
//        }
//    } (params, xhr)
//    // now the progress event...
//    xhr.upload.onprogress = function (progressCallback, caller) {
//        return function (e) {
//            if (e.lengthComputable) {
//                progressCallback(e.loaded, e.total, caller)
//            }
//        }
//    } (progressCallback, caller)
//    xhr.send(formData);
//}
func (this *FileManager) CreateFile(filename string, filepath string, thumbnailMaxHeight int64, thumbnailMaxWidth int64, dbFile bool, messageId string, sessionId string) *CargoEntities.File {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	tmpPath := GetServer().GetConfigurationManager().GetTmpPath() + "/" + filename

	// I will open the file form the tmp directory.
	filedata, err := ioutil.ReadFile(tmpPath)

	// remove the tmp file if it file path is not empty... otherwise the
	// file will bee remove latter.
	defer os.Remove(tmpPath)

	if err != nil {
		errObj := NewError(Utility.FileLine(), FILE_NOT_FOUND_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	dirEntity, dirErrObj := this.getFileById(Utility.CreateSha1Key([]byte(filepath)))

	if dirErrObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, dirErrObj)
		return nil
	}

	file, errObj := this.createFile(dirEntity, filename, filepath, filedata, sessionId, int(thumbnailMaxHeight), int(thumbnailMaxWidth), dbFile)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	return file
}

// @api 1.0
// Remove a file entity with a given uuid.
// @param {string} uuid The file uuid.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) DeleteFile(uuid string, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	err := this.deleteFile(uuid)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), FILE_DELETE_ERROR, SERVER_ERROR_CODE, errors.New("Failed to delete file with uuid '"+uuid+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
	}
}

// @api 1.0
// Test if a given file exist.
// @param {string} filename The name of the file to create.
// @param {string} filepath The path of the directory where to create the file.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {bool} Return true if the file exist.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) IsFileExist(filename string, filepath string) bool {

	fileId := Utility.CreateSha1Key([]byte(filepath + "/" + filename))
	fileUuid := CargoEntitiesFileExists(fileId)
	if len(fileUuid) > 0 {
		return true
	}
	return false
}

// @api 1.0
// Retreive the mime type information from a given extention.
// @param {string} fileExtension The file extention ex. txt, xls, html, css
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*Server.MimeType} The mime type information.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) GetMimeTypeByExtension(fileExtension string, messageId string, sessionId string) *MimeType {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	mimeType := this.mimeTypeMap[strings.ToLower(fileExtension)]
	if mimeType == nil {
		cargoError := NewError(Utility.FileLine(), MIMETYPE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("MimeType for file extension '"+fileExtension+"' doesn't exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}
	return mimeType
}

// @api 1.0
// Retreive a file with a given id
// @param {string} path The file path.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) GetFileByPath(path string, messageId string, sessionId string) *CargoEntities.File {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	fileId := Utility.CreateSha1Key([]byte(path))
	file, errObj := this.getFileById(fileId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// I will also get the data from the disk...
	if !file.IsDir() && file.GetFileType() == CargoEntities.FileType_DiskFile {
		filedata, err := ioutil.ReadFile(this.root + path)
		if err != nil {
			GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), FILE_READ_ERROR, SERVER_ERROR_CODE, errors.New("File '"+this.root+path+"' could not be read.")))
			return nil
		}
		// encode to a string 64 oject...
		file.SetData(base64.StdEncoding.EncodeToString(filedata))
	}

	// Return the file object...
	return file
}

// @api 1.0
// Open a file with a given id
// @param {string} id The file id.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.File} Return the file with it content *In case of large file use downloadFile instead.
// @scope {public}
// @param {callback} progressCallback The function is call when chunk of response is received.
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
//FileManager.prototype.openFile = function (fileId, progressCallback, successCallback, errorCallback, caller) {
//    // server is the client side singleton.
//    var params = []
//    params.push(createRpcData(fileId, "STRING", "fileId"))
//    server.executeJsFunction(
//        "FileManagerOpenFile", // The function to execute remotely on server
//        params, // The parameters to pass to that function
//        function (index, total, caller) { // The progress callback
//            caller.progressCallback(index, total, caller.caller)
//        },
//        function (result, caller) {
//            var file = new CargoEntities.File()
//            file.init(result[0])
//            server.entityManager.setEntity(file)
//            caller.successCallback(file, caller.caller)
//            var evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": file } }
//            server.eventHandler.broadcastLocalEvent(evt)
//        },
//        function (errMsg, caller) {
//            caller.errorCallback(errMsg, caller.caller)
//            server.errorHandler.onError(errMsg)
//        }, // Error callback
//        { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback } // The caller
//    )
//}
func (this *FileManager) OpenFile(fileId string, messageId string, sessionId string) *CargoEntities.File {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	file, errObj := this.openFile(fileId, sessionId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// I will also get the data from the disk...
	if !file.IsDir() && file.GetFileType() == CargoEntities.FileType_DiskFile {
		filedata, err := ioutil.ReadFile(this.root + file.GetPath() + "/" + file.GetName())
		if err != nil {
			GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), FILE_READ_ERROR, SERVER_ERROR_CODE, errors.New("File '"+this.root+file.GetPath()+"/"+file.GetName()+"' could not be read.")))
			return nil
		}
		// encode to a string 64 oject...
		file.SetData(base64.StdEncoding.EncodeToString(filedata))
	}

	// Return the file object...
	return file
}

//////////////////// Excel Write/Read function /////////////////////////////////

// @api 1.0
// Write an array of value to an excel file.
// @param {string} path The xlsx file path.
// @param {string} sheetName The name of the sheet where to write the values.
// @param {values} the double dimension values of array.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) WriteExcelFile(filePath string, sheetName string, values [][]interface{}, messageId string, sessionId string) {
	// In case of temp dir...
	if strings.HasPrefix(filePath, "/tmp/") {
		// Here the file must be create in the temp path.
		filePath = GetServer().GetConfigurationManager().GetTmpPath() + "/" + filePath[len("/tmp/"):]
	} else {
		if !strings.HasPrefix(filePath, "/") {
			filePath = "/" + filePath
		}
		filePath = this.root + filePath
	}

	xlFile, err := xlsx.OpenFile(filePath)
	var xlSheet *xlsx.Sheet
	if err != nil {
		xlFile = xlsx.NewFile()
		xlSheet, _ = xlFile.AddSheet(sheetName)
	} else {
		xlSheet = xlFile.Sheet[sheetName]
		if xlSheet == nil {
			xlSheet, _ = xlFile.AddSheet(sheetName)
		}
	}

	// So here I got the xl file open and sheet ready to write into.
	for i := 0; i < len(values); i++ {
		row := xlSheet.AddRow()
		for j := 0; j < len(values[i]); j++ {
			if values[i][j] != nil {
				cell := row.AddCell()
				if reflect.TypeOf(values[i][j]).String() == "string" {
					str := values[i][j].(string)
					// here I will try to format the date time if it can be...
					dateTime, err := Utility.DateTimeFromString(str, "2006-01-02 15:04:05")
					if err != nil {
						cell.SetString(str)
					} else {
						cell.SetDateTime(dateTime)
					}
				} else {
					if values[i][j] != nil {
						cell.SetValue(values[i][j])
					}
				}
			}
		}
	}

	// Here I will save the file at the given path...
	err = xlFile.Save(filePath)

	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), FILE_WRITE_ERROR, SERVER_ERROR_CODE, err))
	}

}

// @api 1.0
// Read the content of an excel file and return it values as form of a map.
// @param {string} filePath The file path of the file to read.
// @param {string} sheetName The name of the page to read, if there no value the whole file is return.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {map[string][][]interface{}} The return map item represent page by name.
// @scope {public}
// @param {callback} progressCallback The function is call when chunk of response is received.
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) ReadExcelFile(filePath string, sheetName string, messageId string, sessionId string) map[string][][]string {
	values := make(map[string][][]string, 0)

	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), FILE_READ_ERROR, SERVER_ERROR_CODE, err))
		return values
	}

	if len(sheetName) > 0 { // Only the sheet whit that name will be read.
		xlSheet := xlFile.Sheet[sheetName]
		if xlSheet == nil {
			GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), FILE_READ_ERROR, SERVER_ERROR_CODE, errors.New("The file "+filePath+" has no sheet name "+sheetName)))
			return values
		}
		for _, row := range xlSheet.Rows {
			row_ := make([]string, 0)
			for _, cell := range row.Cells {
				row_ = append(row_, cell.Value)
			}
			values[xlSheet.Name] = append(values[xlSheet.Name], row_)
		}
	} else { // The whole sheet must be read in that case.
		for _, sheet := range xlFile.Sheets {
			values[sheet.Name] = make([][]string, 0)
			for _, row := range sheet.Rows {
				row_ := make([]string, 0)
				for _, cell := range row.Cells {
					row_ = append(row_, cell.Value)
				}
				values[sheet.Name] = append(values[sheet.Name], row_)
			}
		}
	}

	return values
}

/////////////////////////////////////////////////////////////////////////////
// Http handler
/////////////////////////////////////////////////////////////////////////////
/**
 * This code is use to upload a file into the tmp directory of the server
 * via http request.
 */
func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	// I will
	tmpPath := GetServer().GetConfigurationManager().GetTmpPath()

	err := r.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		log.Println(w, err)
		return
	}

	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	files := formdata.File["multiplefiles"] // grab the filenames
	for i, _ := range files {               // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Println(w, err)
			return
		}

		out, err := os.Create(tmpPath + "/" + files[i].Filename)
		defer out.Close()
		if err != nil {
			log.Println(w, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !
		if err != nil {
			log.Println(w, err)
			return
		}

	}

}
