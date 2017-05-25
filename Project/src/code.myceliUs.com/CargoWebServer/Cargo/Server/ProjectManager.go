/**
 *
 */
package Server

import (
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
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
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())

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

/**
 * Synchronize all the project from application root directory.
 */
func (this *ProjectManager) synchronize() {
	// Each directory contain an application...
	files, _ := ioutil.ReadDir(this.root)

	for _, f := range files {
		if f.IsDir() {
			log.Println("Synchronize project ", f.Name())
			if !strings.HasPrefix(".", f.Name()) {

				projectUUID := CargoEntitiesProjectExists(f.Name())
				var project *CargoEntities.Project
				if len(projectUUID) == 0 {
					project = new(CargoEntities.Project)
					project.SetName(f.Name())
					project.SetId(f.Name())
					cargoEntities := server.GetEntityManager().getCargoEntities()

					// first of all i will see if the project exist...
					GetServer().GetEntityManager().NewCargoEntitiesProjectEntity(cargoEntities.GetUuid(), "", project)

					// Here I will synchronyse the project...
					this.synchronizeProject(project, this.root+"/"+f.Name())

					cargoEntities.GetObject().(*CargoEntities.Entities).SetEntities(project)
					cargoEntities.SaveEntity()
				} else {
					projectEntity, _ := GetServer().GetEntityManager().getEntityByUuid(projectUUID, false)
					this.synchronizeProject(projectEntity.GetObject().(*CargoEntities.Project), this.root+"/"+f.Name())
					projectEntity.SaveEntity()
				}
			}
		}
	}
}

/**
 * Synchronize a project.
 */
func (this *ProjectManager) synchronizeProject(project *CargoEntities.Project, projectPath string) {

	// Set the file id...
	ids := []interface{}{Utility.CreateSha1Key([]byte("/" + project.GetId()))}

	// I will keep reference to the project directory only...
	file, err := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.File", ids, false) // get the first file level only...
	if err == nil {
		project.SetFilesRef(file.GetObject().(*CargoEntities.File))
	}
}
