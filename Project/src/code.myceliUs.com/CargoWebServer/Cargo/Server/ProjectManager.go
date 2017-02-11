/**
 *
 */
package Server

import (
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
	//"bytes"
	//"encoding/xml"
	"io/ioutil"
	"log"
	"strings"
)

type ProjectManager struct {
	activeProjects map[string]*CargoEntities.Project
	root           string
	m_config       *Config.ServiceConfiguration
}

var projectManager *ProjectManager

func (this *Server) GetProjectManager() *ProjectManager {
	if projectManager == nil {
		projectManager = newProjectManager()
	}
	return projectManager
}

/**
 * This function create and return the session manager...
 */
func newProjectManager() *ProjectManager {
	projectManager := new(ProjectManager)
	projectManager.activeProjects = make(map[string]*CargoEntities.Project, 0)
	projectManager.root = GetServer().GetConfigurationManager().GetApplicationDirectoryPath()
	return projectManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Do intialysation stuff here.
 */
func (this *ProjectManager) initialize() {
	// register service avalaible action here.
	log.Println("--> Initialize ProjectManager")
	this.m_config = GetServer().GetConfigurationManager().getServiceConfiguration(this.getId())
	this.m_config.M_start = false
}

func (this *ProjectManager) getId() string {
	return "ProjectManager"
}

func (this *ProjectManager) start() {
	log.Println("--> Start ProjectManager")
	// First of all I will scan the application directory to see if project
	// project exist...
	this.synchronize()
}

func (this *ProjectManager) stop() {
	log.Println("--> Stop ProjectManager")
}

func (this *ProjectManager) getConfig() *Config.ServiceConfiguration {
	return this.m_config
}

/**
 * Synchronize all the project from application root directory.
 */
func (this *ProjectManager) synchronize() {
	// Each directory contain an application...
	files, _ := ioutil.ReadDir(this.root)
	cargoEntities := server.GetEntityManager().getCargoEntities()
	for _, f := range files {
		if f.IsDir() {
			log.Println("Synchronize project ", f.Name())
			if !strings.HasPrefix(".", f.Name()) {
				// first of all i will see if the project exist...
				projectEntity := GetServer().GetEntityManager().NewCargoEntitiesProjectEntity(f.Name(), nil)

				// Get the reference to the object...
				project := projectEntity.GetObject().(*CargoEntities.Project)

				// Set the project id...
				project.SetId(f.Name())

				// Set the projet name if is not already set.
				if len(project.GetName()) == 0 {
					project.SetName(f.Name())
				}

				// Here I will synchronyse the project...
				this.synchronizeProject(project, this.root+"/"+f.Name())

				// Here i will save the entity...
				projectEntity.SetNeedSave(true)
				cargoEntities.GetObject().(*CargoEntities.Entities).SetEntities(project)
			}
		}
	}
	cargoEntities.SaveEntity()
}

/**
 * Synchronize a project.
 */
func (this *ProjectManager) synchronizeProject(project *CargoEntities.Project, projectPath string) {

	// Set the file id...
	fileId := Utility.CreateSha1Key([]byte("/" + project.GetId()))

	// I will keep reference to the project directory only...
	file, err := GetServer().GetEntityManager().getEntityById("CargoEntities.File", fileId) // get the first file level only...
	if err == nil {
		project.SetFilesRef(file.GetObject().(*CargoEntities.File))
	} else {
		log.Println("--------> file ", fileId, " was not found!")
	}
}
