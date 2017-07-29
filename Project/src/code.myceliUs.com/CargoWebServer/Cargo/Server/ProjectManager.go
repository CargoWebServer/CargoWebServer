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
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)

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

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

// @api 1.0
// Event handler function.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
//ProjectManager.prototype.onEvent = function (evt) {
//    EventHub.prototype.onEvent.call(this, evt)
//}
func (this *ProjectManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

// @api 1.0
// Synchronize the project found in the server Apps directory.
// @param {string} id The LDAP server connection id.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Account} The new registered account.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *ProjectManager) Synchronize(sessionId string, messageId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}
	// Synchronize the content of project.
	this.synchronize()
}

// @api 1.0
// Return the list of all projects on the server.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Account} The new registered account.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *ProjectManager) GetAllProjects(sessionId string, messageId string) []*CargoEntities.Project {
	projects := make([]*CargoEntities.Project, 0)
	/*errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return projects
	}*/
	entities, errObj := GetServer().GetEntityManager().getEntitiesByType("CargoEntities.Project", "", "CargoEntities", false)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return projects
	}

	// Here I will parse the list of entities retreived...
	for i := 0; i < len(entities); i++ {
		projects = append(projects, entities[i].GetObject().(*CargoEntities.Project))
	}
	return projects
}
