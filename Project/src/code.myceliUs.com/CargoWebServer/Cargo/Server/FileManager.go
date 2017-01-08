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
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Persistence/CargoEntities"
	"github.com/nfnt/resize"

	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
	"github.com/polds/imgbase64"
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

func (this *FileManager) Initialize() {
	// Now I will synchronize files...
	rootDir := this.synchronize(this.root)
	rootEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntityFromObject(rootDir)
	rootEntity.SaveEntity()

	//server.entityManager.cargoEntities.GetObject().(*CargoEntities.Entities).SetEntities(rootEntity.GetObject())
	//server.entityManager.cargoEntities.SaveEntity()
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
			log.Println("Root Dir entity ", dirEntity.UUID)
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
		log.Println("Get file ", filePath__)
		fileEntity, err := this.getFileById(fileId)

		// Remove from file to delete...
		delete(toDelete, fileId)

		if err != nil {
			// here I will create the new entity...
			if !f.IsDir() {
				// Now I will open the file and create the entry in the DB.
				filedata, _ := ioutil.ReadFile(this.root + filePath__)
				if !strings.HasPrefix(f.Name(), ".") {
					file, err := this.createFile(f.Name(), filePath_, filedata, "", 128, 128, false)
					if err == nil {
						dirEntity.SetFiles(file)
					} else {
						log.Println("--------------> fail to create file ", filePath__)
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

	// The directory entity.
	var dirEntity *CargoEntities_FileEntity
	var dir *CargoEntities.File

	dirId := Utility.CreateSha1Key([]byte(dirPath_))

	dirEntity = GetServer().GetEntityManager().NewCargoEntitiesFileEntity(dirId, nil)
	dir = dirEntity.GetObject().(*CargoEntities.File)
	dir.SetId(dirId)
	dir.SetIsDir(true)
	dir.SetName(dirName)
	dir.SetPath(dirPath)
	dir.NeedSave = true
	dir.SetFileType(CargoEntities.FileType_DiskFile)

	// Set the cargo entities object.
	entities := GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities)
	dir.SetEntitiesPtr(entities)

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
	if len(parentDirPath) > 0 {
		// Retreive the perentDirEntity, put the file in it and save the parent
		var parentDir *CargoEntities.File
		var parentDirEntity *CargoEntities_FileEntity
		parentDirEntityId := Utility.CreateSha1Key([]byte(parentDirPath))
		parentDirUuid := CargoEntitiesFileExists(parentDirEntityId)

		if len(parentDirUuid) == 0 {
			// Here I will create the parent directory...
			index := strings.LastIndex(parentDirPath, "/")
			var cargoError *CargoEntities.Error
			parentDir, cargoError = this.createDir(parentDirPath[index+1:], parentDirPath[0:index], sessionId)
			// Create the error message
			if cargoError != nil {
				return nil, cargoError
			}
			parentDirUuid = parentDir.GetUUID()
		}

		parentDirEntity = GetServer().GetEntityManager().NewCargoEntitiesFileEntity(parentDirUuid, nil)
		parentDir = parentDirEntity.GetObject().(*CargoEntities.File)
		dir.SetParentDirPtr(parentDir)
		parentDir.SetFiles(dir)

		// Save the parent directory. This will also save the directory.
		parentDirEntity.SaveEntity()

	}

	dirEntity.SaveEntity()

	return dir, nil
}

/**
 * Create a file.
 */
func (this *FileManager) createFile(filename string, filepath string, filedata []byte, sessionId string, thumbnailMaxHeight int, thumbnailMaxWidth int, dbFile bool) (*CargoEntities.File, *CargoEntities.Error) {
	var fileEntity *CargoEntities_FileEntity
	var file *CargoEntities.File
	var fileId string
	var f *os.File
	var err error
	var checksum string
	var parentDirPath string

	// The parent is the root.
	if len(filepath) == 0 {
		parentDirPath = "CARGOROOT"
	} else {
		parentDirPath = filepath
	}

	if !dbFile {
		// Write the file data. Try to decode it if it is encoded, and decode it if it is not.
		if _, err := os.Stat(this.root + filepath + "/" + filename); err != nil {
			// TODO Throw an error if err != nil ?
			if os.IsNotExist(err) {
				ioutil.WriteFile(this.root+filepath+"/"+filename, filedata, 0644)
			}
		}

		// The user wants to save the data into a physical file.
		fileId = Utility.CreateSha1Key([]byte(filepath + "/" + filename))

		log.Println("-----------> ", filepath+"/"+filename)

		fileUuid := CargoEntitiesFileExists(fileId)
		if len(fileUuid) > 0 {
			// Create the error message
			cargoError := NewError(Utility.FileLine(), FILE_ALREADY_EXISTS_ERROR, SERVER_ERROR_CODE, errors.New("The file '"+fileId+"' already exists. "))
			return nil, cargoError
		}

		fileEntity = GetServer().GetEntityManager().NewCargoEntitiesFileEntity(fileId, nil)
		file = fileEntity.GetObject().(*CargoEntities.File)
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
		fileId = Utility.RandomUUID()
		file.SetFileType(CargoEntities.FileType_DbFile)
		fileEntity = GetServer().GetEntityManager().NewCargoEntitiesFileEntity(fileId, nil)
		file = fileEntity.GetObject().(*CargoEntities.File)

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

	// Set the basic information.
	file.SetId(fileId)
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

	// Save the entity.
	if dbFile {
		fileEntity.SaveEntity()
	} else {
		parentDirEntityId := Utility.CreateSha1Key([]byte(parentDirPath))
		parentDirUuid := CargoEntitiesFileExists(parentDirEntityId)
		if len(parentDirUuid) == 0 {
			// Create the error message
			cargoError := NewError(Utility.FileLine(), INVALID_DIRECTORY_PATH_ERROR, SERVER_ERROR_CODE, errors.New("The path '"+parentDirPath+"' was not found. "))
			return nil, cargoError
		}
		parentDirEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity(parentDirUuid, nil)
		parentDir := parentDirEntity.GetObject().(*CargoEntities.File)
		file.SetParentDirPtr(parentDir)
		parentDir.SetFiles(file)
		// Save the parent and the file.
		parentDirEntity.SaveEntity()
	}

	defer f.Close()

	if err == nil {
		eventData := make([]*MessageData, 1)

		fileInfo := new(MessageData)
		fileInfo.Name = "fileInfo"
		fileInfo.Value = file
		eventData[0] = fileInfo

		evt, _ := NewEvent(NewFileEvent, FileEvent, eventData)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}

	return file, nil
}

/**
 * Return a file with a given uuid, or id.
 */
func (this *FileManager) getFileById(id string) (*CargoEntities.File, *CargoEntities.Error) {
	fileEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.File", id)
	if errObj == nil {
		// Initialize it's content.
		fileEntity.InitEntity(fileEntity.GetUuid())
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

	// I will retreive the file...
	fileEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity(uuid, nil)
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
	fileEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity(uuid, nil)
	fileEntity.InitEntity(uuid)
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
		parentDirEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity(file.GetParentDirPtr().UUID, nil)
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

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

/**
 * Create a new directory and it's associated file entity on the server
 */
func (this *FileManager) CreateDir(dirName string, dirPath string, messageId string, sessionId string) *CargoEntities.File {
	dir, errObj := this.createDir(dirName, dirPath, sessionId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	return dir
}

/**
 * Create a file
 */
func (this *FileManager) CreateFile(filename string, filepath string, thumbnailMaxHeight int64, thumbnailMaxWidth int64, dbFile bool, messageId string, sessionId string) *CargoEntities.File {

	tmpPath := GetServer().GetConfigurationManager().GetTmpPath() + "/" + filename

	// I will open the file form the tmp directory.
	filedata, err := ioutil.ReadFile(tmpPath)

	if err != nil {
		errObj := NewError(Utility.FileLine(), FILE_NOT_FOUND_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	file, errObj := this.createFile(filename, filepath, filedata, sessionId, int(thumbnailMaxHeight), int(thumbnailMaxWidth), dbFile)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// remove the tmp file.
	os.Remove(tmpPath)

	return file
}

/**
 * Delete a file with a given uuid
 */
func (this *FileManager) DeleteFile(uuid string, messageId string, sessionId string) {
	err := this.deleteFile(uuid)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), FILE_DELETE_ERROR, SERVER_ERROR_CODE, errors.New("Failed to delete file with uuid '"+uuid+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
	}
}

/**
 * Test if a given file exist.
 */
func (this *FileManager) IsFileExist(filename string, filepath string) bool {
	fileId := Utility.CreateSha1Key([]byte(filepath + "/" + filename))
	fileUuid := CargoEntitiesFileExists(fileId)
	if len(fileUuid) > 0 {
		return true
	}
	return false
}

/**
 * Get the mime type information...
 */
func (this *FileManager) GetMimeTypeByExtension(fileExtension string, messageId string, sessionId string) *MimeType {
	mimeType := this.mimeTypeMap[strings.ToLower(fileExtension)]
	if mimeType == nil {
		cargoError := NewError(Utility.FileLine(), MIMETYPE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("MimeType for file extension '"+fileExtension+"' doesn't exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}
	return mimeType
}

/**
 * Return a file with a given path.
 */
func (this *FileManager) GetFileByPath(path string, messageId string, sessionId string) *CargoEntities.File {
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

/**
 * Function call when the user open a file.
 */
func (this *FileManager) OpenFile(fileId string, messageId string, sessionId string) *CargoEntities.File {

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
