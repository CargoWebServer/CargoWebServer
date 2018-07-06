package Server

import (
	"bufio"
	"bytes"
	base64 "encoding/base64"
	"encoding/csv"
	"encoding/json"
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
	"golang.org/x/net/html"

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

}

func (this *FileManager) stop() {
	log.Println("--> Stop FileManager")
}

func (this *FileManager) synchronizeAll() {
	// Now I will synchronize files...
	rootDir := this.synchronize(this.root)
	GetServer().GetEntityManager().saveEntity(rootDir)
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

	if filePath_ == "/lib" {
		return nil
	}

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
						dirEntity.AppendFiles(file)
					} else {
						log.Panicln("--------------> fail to create file ", filePath__, err)
					}
				}
			} else {
				// make a recursion...
				if !strings.HasPrefix(f.Name(), ".") {
					subDir := this.synchronize(this.root + filePath__)
					if subDir != nil {
						dirEntity.AppendFiles(subDir)
					}
				}
			}
		} else {
			// I will test the checksum to see if the file has change...
			if !f.IsDir() {
				// Update the file checksum and save it if the file has change...
				if !strings.HasPrefix(f.Name(), ".") {
					filedata, _ := ioutil.ReadFile(this.root + filePath__)
					this.saveFile(fileEntity.GetUuid(), filedata, "", 128, 128, false)
					dirEntity.AppendFiles(fileEntity)
				}
			} else {
				// make a recursion...
				if !strings.HasPrefix(f.Name(), ".") {
					subDir := this.synchronize(this.root + filePath__)
					if subDir != nil {
						dirEntity.AppendFiles(subDir)
					}
				}
			}
		}
	}

	// I will now remove the remaining files in the map...
	for _, fileToDelete := range toDelete {
		// Delete the associated entity...
		err := this.deleteFile(fileToDelete.GetUuid())
		if err == nil {
			log.Println("--> Delete file: ", fileToDelete.GetPath()+"/"+fileToDelete.GetName())
		} else {
			log.Println("---> ", err)
		}
	}

	return dirEntity
}

/**
 * Create a new directory and it's associated file entity on the server
 */
// TODO add the logic for db files...
func (this *FileManager) createDir(dirName string, dirPath string, sessionId string) (*CargoEntities.File, *CargoEntities.Error) {
	// no file entities must be created for that dir.
	if dirPath+"/"+dirName == "/lib" {
		cargoError := NewError(Utility.FileLine(), INVALID_DIRECTORY_PATH_ERROR, SERVER_ERROR_CODE, errors.New("lib directory can not contain file entities."))
		return nil, cargoError
	}

	// Return the dir entity if it already exist.
	dirId := Utility.CreateSha1Key([]byte(dirPath + "/" + dirName))
	dirEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.File", "CargoEntities", []interface{}{dirId})
	if errObj == nil {
		return dirEntity.(*CargoEntities.File), nil
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
	parentDirEntityId := Utility.CreateSha1Key([]byte(parentDirPath))
	parentDir, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.File", "CargoEntities", []interface{}{parentDirEntityId})

	if parentDir == nil {
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
	// function pointer setting here.
	dir.SetEntityGetter(getEntityFct)
	dir.SetEntitySetter(setEntityFct)
	dir.SetUuidGenerator(generateUuidFct)

	// Set the cargo entities object.
	entities := GetServer().GetEntityManager().getCargoEntities()
	dir.SetEntitiesPtr(entities)

	if parentDir != nil {
		// That will set the uuid
		dirEntity, _ := GetServer().GetEntityManager().createEntity(parentDir.GetUuid(), "M_files", dir)
		dir = dirEntity.(*CargoEntities.File)
		dir.SetParentDirPtr(parentDir.(*CargoEntities.File))
		dir.SetEntitiesPtr(entities)
		parentDir.(*CargoEntities.File).AppendFiles(dir)
		GetServer().GetEntityManager().saveEntity(dir)
	} else {
		// Save the dir.
		GetServer().GetEntityManager().saveEntity(dir)
	}

	return dir, nil

}

/**
 * Create a file.
 */
func (this *FileManager) createFile(parentDir *CargoEntities.File, filename string, filepath string, filedata []byte, sessionId string, thumbnailMaxHeight int, thumbnailMaxWidth int, dbFile bool) (*CargoEntities.File, *CargoEntities.Error) {
	// Not create files from lib directory
	if strings.HasPrefix(filepath+"/"+filename, "/lib/") {
		cargoError := NewError(Utility.FileLine(), INVALID_DIRECTORY_PATH_ERROR, SERVER_ERROR_CODE, errors.New("lib directory can not contain file entities."))
		return nil, cargoError
	}

	var file *CargoEntities.File
	var f *os.File
	var err error
	var checksum string

	// Use to determine if the file already exist or not.
	isNew := true

	parentDirEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(parentDir.GetUuid())
	if errObj != nil {
		return nil, errObj
	}

	// I will retreive the file.
	fileId := Utility.CreateSha1Key([]byte(filepath + "/" + filename))
	fileEntity, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.File", "CargoEntities", []interface{}{fileId})

	if fileEntity != nil {
		isNew = false
		file = fileEntity.(*CargoEntities.File)
		log.Println("Get entity for file: ", filepath+"/"+filename)
	} else {
		log.Println("Create entity for file: ", filepath+"/"+filename)
		file = new(CargoEntities.File)
		file.SetId(fileId)
		file.SetParentUuid(parentDir.GetUuid())
		file.SetParentLnk("M_files")
		// function pointer setting here.
		file.SetEntityGetter(getEntityFct)
		file.SetEntitySetter(setEntityFct)
		file.SetUuidGenerator(generateUuidFct)
	}

	// Disk file.
	if !dbFile {
		// Here I will get the parent put the file in it and save it.
		parentDir := parentDirEntity.(*CargoEntities.File)
		file.SetParentDirPtr(parentDir)
		parentDir.AppendFiles(file)

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
		checksum_ := Utility.CreateFileChecksum(f_)

		f_.Close()
		os.Remove(f_.Name())

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
		f_, err := ioutil.TempFile("", "_cargo_tmp_file_")
		if err != nil {
			// Create the error message
			cargoError := NewError(Utility.FileLine(), FILE_OPEN_ERROR, SERVER_ERROR_CODE, errors.New("Failed to open _cargo_tmp_file_ for file '"+filename+"'. "))
			return nil, cargoError

		}
		f_.Write(filedata)

		// Create the checksum
		checksum = Utility.CreateFileChecksum(f_)

		// Keep the data into the data base as a base 64 string
		file.SetData(base64.StdEncoding.EncodeToString(filedata))

		// delete temporary file now.
		f_.Close()
		os.Remove(f_.Name())
	}

	// Set general information.
	file.SetPath(filepath)
	file.SetName(filename)
	file.SetChecksum(checksum)

	// Set the entity object.
	entities := GetServer().GetEntityManager().getCargoEntities()
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

	// Create the file if is new, genereate it uuid...
	if isNew {
		fileEntity, _ := GetServer().GetEntityManager().createEntity(parentDirEntity.GetUuid(), "M_files", file)
		file = fileEntity.(*CargoEntities.File) // cast...
	}

	if err == nil {
		eventData := make([]*MessageData, 2)
		fileInfo := new(MessageData)
		fileInfo.TYPENAME = "Server.MessageData"
		fileInfo.Name = "fileInfo"
		if strings.HasPrefix(file.M_mime, "text/") || strings.HasPrefix(file.M_mime, "application/") {
			file.SetData(base64.StdEncoding.EncodeToString(filedata))
		}

		fileInfo.Value = file
		eventData[0] = fileInfo

		prototypeInfo := new(MessageData)
		prototypeInfo.TYPENAME = "Server.MessageData"
		prototypeInfo.Name = "prototype"
		prototypeInfo.Value, _ = GetServer().GetEntityManager().getEntityPrototype("CargoEntities.File", "CargoEntities")
		eventData[1] = prototypeInfo

		var evt *Event
		if isNew {
			// New file create
			evt, _ = NewEvent(NewFileEvent, FileEvent, eventData)
		} else {
			// Existing file was update
			evt, _ = NewEvent(UpdateFileEvent, FileEvent, eventData)
		}

		GetServer().GetEventManager().BroadcastEvent(evt)

		// append the ?checksum in includes path.
		if file.GetMime() == "text/html" {
			filedata = this.setHtmlIncludes(file.GetPath(), filedata)
		}
	}

	return file, nil
}

/**
 * Save the content of the file...
 */
func (this *FileManager) saveFile(uuid string, filedata []byte, sessionId string, thumbnailMaxHeight int, thumbnailMaxWidth int, dbFile bool) error {

	// Create a temporary file object from the data...
	f_, err := ioutil.TempFile(os.TempDir(), uuid)
	if err != nil {
		log.Println("Fail to open _cargo_tmp_file_ for file ", uuid, " ", err)
		return err
	}

	f_.Write(filedata)

	checksum := Utility.CreateFileChecksum(f_)

	// close the file and remove it.
	f_.Close()
	os.Remove(f_.Name())

	// I will retreive the file, the uuid must exist in that case.
	fileEntity, _ := GetServer().GetEntityManager().getEntityByUuid(uuid)
	file := fileEntity.(*CargoEntities.File)

	if file.GetChecksum() != checksum {
		// append the ?checksum in includes path.
		var oldChecksum = file.GetChecksum()
		if file.GetMime() == "text/html" {
			filedata = this.setHtmlIncludes(file.GetPath(), filedata)
		}

		if !dbFile {
			// Save the data to the disck...
			filename := file.GetName()
			filepath := file.GetPath()
			err = ioutil.WriteFile(this.root+"/"+filepath+"/"+filename, filedata, 0644)
			if err != nil {
				log.Println("---> fail to write file: ", this.root+"/"+filepath+"/"+filename, err)
				return err
			}
		} else {
			// Set the new data...
			file.SetData(string(filedata))
		}

		file.SetChecksum(checksum)
		err := GetServer().GetEntityManager().saveEntity(file)
		if err != nil {
			log.Println("---> save file error ", err)
		} else if file.GetMime() == "application/javascript" || file.GetMime() == "text/css" || file.GetMime() == "text/json" {
			// Here I will remplace the old checksum with the new one.
			this.updateHtmlChecksum(oldChecksum, checksum, file.GetPath()+"/"+file.GetName())
		}

	} /*else {
		return nil
	}*/

	if err == nil {
		if strings.HasPrefix(file.M_mime, "text/") || strings.HasPrefix(file.M_mime, "application/") {
			file.SetData(base64.StdEncoding.EncodeToString(filedata))
		}

		eventData := make([]*MessageData, 2)
		fileInfo := new(MessageData)
		fileInfo.TYPENAME = "Server.MessageData"
		fileInfo.Name = "fileInfo"
		fileInfo.Value = file
		eventData[0] = fileInfo

		prototypeInfo := new(MessageData)
		prototypeInfo.TYPENAME = "Server.MessageData"
		prototypeInfo.Name = "prototype"
		prototypeInfo.Value, _ = GetServer().GetEntityManager().getEntityPrototype("CargoEntities.File", "CargoEntities")
		eventData[1] = prototypeInfo
		evt, _ := NewEvent(UpdateFileEvent, FileEvent, eventData)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}

	return nil
}

func (this *FileManager) updateHtmlChecksum(oldChecksum string, newChecksum string, filePath string) {
	// So first of all I need to get all html files...
	var query EntityQuery
	query.Fields = []string{"UUID", "M_path", "M_name", "M_mime"}
	query.TYPENAME = "Server.EntityQuery"
	query.TypeName = "CargoEntities.File"
	query.Query = `CargoEntities.File.M_mime == "text/html"`
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData("CargoEntities", string(queryStr), []interface{}{}, []interface{}{})
	if err == nil {
		for i := 0; i < len(results); i++ {
			fileInfo := results[i]
			// So here I will read the file content...
			txt, err := ioutil.ReadFile(this.root + fileInfo[1].(string) + "/" + fileInfo[2].(string))
			if err == nil {
				if strings.Index(string(txt), oldChecksum) != -1 {
					// The old reference was found!
					data := strings.Replace(string(txt), oldChecksum, newChecksum, -1)
					this.saveFile(fileInfo[0].(string), []byte(data), "", 256, 256, false)
				} else if strings.Index(string(txt), filePath) != -1 {
					// In case the file was not exist before.
					data := strings.Replace(string(txt), filePath, filePath+"?"+newChecksum, -1)
					this.saveFile(fileInfo[0].(string), []byte(data), "", 256, 256, false)
				} else if strings.HasPrefix(filePath, fileInfo[1].(string)) {
					// In case the file is in the same directory as the html file and
					// it path is express relatively.
					relativeFilePath := filePath[len(fileInfo[1].(string))+1:]
					data := strings.Replace(string(txt), relativeFilePath, relativeFilePath+"?"+newChecksum, -1)
					this.saveFile(fileInfo[0].(string), []byte(data), "", 256, 256, false)
				}
			}
		}
	}
}

/**
 * If the file is an html file and it contain link to file entities I will append
 * the file checksum in the import link so the cache will reload that file
 * correctly.
 */
func (this *FileManager) setHtmlIncludes(path string, filedata []byte) []byte {
	doc, _ := html.Parse(strings.NewReader(string(filedata)))
	var f func(*html.Node)
	var refs = make([]*html.Node, 0)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "link" || n.Data == "script") {
			refs = append(refs, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	f_ := func(basePath string, path string, a *html.Attribute) {
		// Retreive the existing file...
		fileId := Utility.CreateSha1Key([]byte(path))
		file, _ := this.getFileById(fileId)

		var checksum string
		if len(strings.Split(path, "?")) > 1 {
			checksum = strings.Split(path, "?")[1]
			path = strings.Split(path, "?")[0]
		}
		if file != nil {
			if checksum != file.GetChecksum() {
				log.Println("--> checksum: ", file.GetChecksum())
				// Change the actual value with the new checksum
				a.Val = path + "?" + file.GetChecksum()
			}
		} else {
			log.Println(basePath + "/" + path)
			fileId := Utility.CreateSha1Key([]byte(basePath + "/" + path))
			file, _ := this.getFileById(fileId)
			if file != nil {
				if checksum != file.GetChecksum() {
					log.Println("--> checksum: ", file.GetChecksum())
					// Change the actual value with the new checksum
					a.Val = path + "?" + file.GetChecksum()
				}
			}
		}
	}

	// iteration over found refs...
	for i := 0; i < len(refs); i++ {
		n := refs[i]
		if n.Data == "script" {
			for j := 0; j < len(n.Attr); j++ {
				if n.Attr[j].Key == "src" {
					f_(path, n.Attr[j].Val, &n.Attr[j])
				}
			}
		} else if n.Data == "link" {
			for j := 0; j < len(n.Attr); j++ {
				if n.Attr[j].Key == "href" {
					f_(path, n.Attr[j].Val, &n.Attr[j])
				}
			}
		}
	}

	// Render the modified document.
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, doc)
	// Return the modified html document.
	return buf.Bytes()
}

/**
 * Delete a file with a given uuid
 */
func (this *FileManager) deleteFile(uuid string) error {
	fileEntity, err := GetServer().GetEntityManager().getEntityByUuid(uuid)
	if err != nil {
		return errors.New(err.GetBody())
	}

	file := fileEntity.(*CargoEntities.File)

	if file.GetFileType() == CargoEntities.FileType_DiskFile {
		var filePath string
		if _, err := os.Stat(this.root + file.GetPath() + "/" + file.GetName()); err != nil {
			// TODO Throw an error if err != nil ?
			if os.IsNotExist(err) {
				filePath = file.GetPath() + "/" + file.GetName()
			}
		} else {
			filePath = this.root + file.GetPath() + "/" + file.GetName()
		}

		if !file.IsDir() {
			// Remove the file from the disck
			os.Remove(filePath) // The file can be already remove...
		} else {
			os.RemoveAll(filePath)
		}
	}

	// Here i will remove the entity...
	GetServer().GetEntityManager().deleteEntity(fileEntity)

	eventData := make([]*MessageData, 2)
	fileInfo := new(MessageData)
	fileInfo.TYPENAME = "Server.MessageData"
	fileInfo.Name = "fileInfo"
	fileInfo.Value = file
	eventData[0] = fileInfo

	prototypeInfo := new(MessageData)
	prototypeInfo.TYPENAME = "Server.MessageData"
	prototypeInfo.Name = "prototype"
	prototypeInfo.Value, _ = GetServer().GetEntityManager().getEntityPrototype("CargoEntities.File", "CargoEntities")
	eventData[1] = prototypeInfo

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
 * Return a file with a given uuid, or id.
 */
func (this *FileManager) getFileById(id string) (*CargoEntities.File, *CargoEntities.Error) {
	ids := []interface{}{id}
	fileEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.File", "CargoEntities", ids)
	if fileEntity != nil {
		return fileEntity.(*CargoEntities.File), errObj
	}
	return nil, errObj
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
	defer mimeTypeFile.Close()
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

		defer file.Close()

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
	dbFile.SetEntityGetter(getEntityFct)
	dbFile.SetEntitySetter(setEntityFct)
	dbFile.SetUuidGenerator(generateUuidFct)
	// Create the file.
	GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntitiesUuid(), "M_entities", dbFile)

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

	// relative to the sever root if it start with a /
	if strings.HasPrefix(filePath, "/") {
		filePath = this.root + "/" + filePath
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

	defer csvfile.Close()

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

	// if the file exist in the root...
	log.Println("--> try to remove file ", this.root+"/"+filePath)
	if _, err := os.Stat(this.root + "/" + filePath); err == nil {
		// TODO Throw an error if err != nil ?
		if os.IsNotExist(err) {
			filePath = this.root + "/" + filePath
		}
	}

	err := os.Remove(this.root + "/" + filePath)
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
// @param {string} filepath The path of the directory where the file is.
// @param {string} filename The name of the file to download
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
//                // return the success callback with the result.
//                successCallback(blob, caller)
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
// Save a file on the server.
// @param {*CargoEntities.File} file The file to save.
// @param {string} filedata The data of the file.
// @param {int} thumbnailMaxHeight The maximum height size of the thumbnail associated with the file (keep the ratio).
// @param {int} thumbnailMaxWidth The maximum width size of the thumbnail associated with the file (keep the ratio).
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.File} The created file entity.
// @scope {public}
// @param {callback} progressCallback The function is call when chunk of response is received.
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
//FileManager.prototype.saveFile = function (file, filedata, thumbnailMaxHeight, thumbnailMaxWidth, successCallback, progressCallback, errorCallback, caller) {
//    // server is the client side singleton.
//    var params = []
//    // The file data (filedata) will be upload with the http protocol...
//    params.push(createRpcData(file, "JSON_STR", "file"))
//    params.push(createRpcData(thumbnailMaxHeight, "INTEGER", "thumbnailMaxHeight"))
//    params.push(createRpcData(thumbnailMaxWidth, "INTEGER", "thumbnailMaxWidth"))
//    // Here I will create a new data form...
//    var formData = new FormData()
//    formData.append("multiplefiles", filedata, file.M_name)
//    // Use the post function to upload the file to the server.
//    var xhr = new XMLHttpRequest()
//    xhr.open('POST', '/uploads', true)
//    // In case of error or success...
//    xhr.onload = function (params, xhr) {
//        return function (e) {
//            if (xhr.readyState === 4) {
//                if (xhr.status === 200) {
//                    // Here I will create the file...
//                    server.executeJsFunction(
//                        "FileManagerSaveFile", // The function to execute remotely on server
//                        params, // The parameters to pass to that function
//                        function (index, total, caller) { // The progress callback
//                            // Keep track of the file transfert.
//                            caller.progressCallback(index, total, caller.caller)
//                        },
//                        function (result, caller) {
//                            caller.successCallback(result[0], caller.caller)
//                        },
//                        function (errMsg, caller) {
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
func (this *FileManager) SaveFile(file *CargoEntities.File, thumbnailMaxHeight int64, thumbnailMaxWidth int64, messageId string, sessionId string) {

	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	tmpPath := GetServer().GetConfigurationManager().GetTmpPath() + "/" + file.GetName()

	// I will open the file from the tmp directory.
	filedata, err := ioutil.ReadFile(tmpPath)

	// remove the tmp file if it file path is not empty... otherwise the
	// file will bee remove latter.
	defer os.Remove(tmpPath)

	if err != nil {
		errObj := NewError(Utility.FileLine(), FILE_NOT_FOUND_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	err = this.saveFile(file.UUID, filedata, sessionId, int(thumbnailMaxHeight), int(thumbnailMaxWidth), file.M_fileType == CargoEntities.FileType_DbFile)
	if err != nil {
		errObj := NewError(Utility.FileLine(), FILE_WRITE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	// Now I will change file checksum in the all index.html files...

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
// @src
//FileManager.prototype.isFileExist = function (filename, filepath, successCallback, errorCallback, caller) {
//    var params = []
//    params.push(createRpcData(filename, "STRING", "filename"))
//    params.push(createRpcData(filepath, "STRING", "filepath"))
//    // Call it on the server.
//    server.executeJsFunction(
//        "FileManagerIsFileExist", // The function to execute remotely on server
//        params, // The parameters to pass to that function
//        function (index, total, caller) { // The progress callback
//            // Nothing special to do here.
//        },
//        function (result, caller) {
//        	if (caller.successCallback != undefined) {
//        		caller.successCallback(result[0], caller.caller)
//          	caller.successCallback = undefined
//          }
//        },
//        function (errMsg, caller) {
//          server.errorManager.onError(errMsg)
//         	if( caller.errorCallback != undefined){
//          	caller.errorCallback(errMsg, caller.caller)
//				caller.errorCallback = undefined
//			}
//        }, // Error callback
//        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
//    )
//}
func (this *FileManager) IsFileExist(filename string, filepath string, messageId string, sessionId string) bool {

	fileId := Utility.CreateSha1Key([]byte(filepath + "/" + filename))
	uuid, err := GetServer().GetEntityManager().getEntityUuidById("CargoEntities.File", "CargoEntities", []interface{}{fileId})
	if err != nil {
		return false
	}
	return len(uuid) > 0
}

// @api 1.0
// Rename a file.
// @param {string} uuid The uuid of the file
// @param {string} filename The new file name.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *FileManager) RenameFile(uuid string, filename string, messageId string, sessionId string) {

	fileEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(uuid)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	// Change the file id and the fileName.
	file := fileEntity.(*CargoEntities.File)

	var path string
	if file.GetParentDirPtr() != nil {
		path = file.GetParentDirPtr().GetPath() + "/" + file.GetParentDirPtr().GetName()
	} else {
		path = file.GetPath()
	}

	if file.GetName() != filename {
		oldFilePath := fileManager.root + file.GetPath() + "/" + file.GetName()
		newFilePath := fileManager.root + path + "/" + filename
		err := os.Rename(oldFilePath, newFilePath)
		if err != nil {
			cargoError := NewError(Utility.FileLine(), FILE_MANAGER_ERROR, SERVER_ERROR_CODE, err)
			GetServer().reportErrorMessage(messageId, sessionId, cargoError)
			return
		}
	} else {
		// nothing to do here.
		return
	}

	fileId := Utility.CreateSha1Key([]byte(file.GetPath() + "/" + filename))
	file.SetPath(path)
	file.SetId(fileId)
	file.SetName(filename)
	GetServer().GetEntityManager().saveEntity(fileEntity)

	// Here If the file is a directory I must change all it child path to...
	if file.IsDir() {
		for i := 0; i < len(file.GetFiles()); i++ {
			this.RenameFile(file.GetFiles()[i].GetUuid(), file.GetFiles()[i].GetName(), messageId, sessionId)
		}
	}

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

	// Append an event listener for that file.
	//conn := GetServer().getConnectionById(sessionId)
	//listener := NewEventListener(file.UUID+"_editor", conn)
	//GetServer().GetEventManager().AddEventListener(listener)

	// Generate the openFileEvent.
	eventData := make([]*MessageData, 1)
	fileInfo := new(MessageData)
	fileInfo.TYPENAME = "Server.MessageData"
	fileInfo.Name = "FileInfo"

	session := GetServer().GetSessionManager().getActiveSessionById(sessionId)
	var account *CargoEntities.Account
	if session != nil {
		account = session.GetAccountPtr()
	}

	fileInfo.Value = map[string]interface{}{"accountId": account.GetUuid(), "fileId": file.GetUuid()}

	eventData[0] = fileInfo

	//evt, _ := NewEvent(OpenFileEvent, file.GetUuid()+"_editor", eventData)
	GetServer().GetEventManager().BroadcastEventData(OpenFileEvent, file.GetUuid()+"_editor", eventData, Utility.RandomUUID(), sessionId)

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

// @api 1.0
// Return the list of file/directory contain in a directory at given path.
// @param {string} path The file path of the file to read.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {[]string} The list of files.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this FileManager) ReadDir(path string, messageId string, sessionId string) []string {
	if strings.HasPrefix(path, "/") {
		path = this.root + path
	}
	var lst []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), FILE_READ_ERROR, SERVER_ERROR_CODE, err))
		return nil
	}

	for _, f := range files {
		lst = append(lst, f.Name())
	}

	return lst
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
	var path string                         // grab the filenames
	if len(r.FormValue("path")) == 0 {
		path = tmpPath
	} else {
		path = r.FormValue("path")
		if strings.HasPrefix(path, "/") {
			path = GetServer().GetFileManager().root + path
		}
	}

	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Println(w, err)
			return
		}
		log.Println("--> upload file to ", path+"/"+files[i].Filename)
		out, err := os.Create(path + "/" + files[i].Filename)
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
