/**
 *
 */
package Server

import (
	"io/ioutil"
	"log"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

type ProjectManager struct {
	root string
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
}

func (this *ProjectManager) stop() {
	log.Println("--> Stop ProjectManager")
}

/**
 * Synchronize all the project from application root directory.
 */
func (this *ProjectManager) synchronize() {
	log.Println("---> synchronize projects")
	// Each directory contain an application...
	files, _ := ioutil.ReadDir(this.root)
	for _, f := range files {
		if f.IsDir() {
			log.Println("--> Synchronize project ", f.Name())
			if !strings.HasPrefix(".", f.Name()) && f.Name() != "lib" && f.Name() != "queries" {
				projectEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Project", "CargoEntities", []interface{}{f.Name()})
				if err != nil {
					project := new(CargoEntities.Project)
					project.SetName(f.Name())
					project.SetId(f.Name())

					// first of all i will see if the project exist...
					projectEntity, err := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntitiesUuid(), "M_entities", "CargoEntities.Project", f.Name(), project)

					// Here I will synchronyse the project...
					if err == nil {
						this.synchronizeProject(projectEntity.(*CargoEntities.Project), this.root+"/"+f.Name())
					}
				} else {
					this.synchronizeProject(projectEntity.(*CargoEntities.Project), this.root+"/"+f.Name())
					// Save as needed.
					GetServer().GetEntityManager().saveEntity(projectEntity)
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
	file, err := GetServer().GetEntityManager().getEntityById("CargoEntities.File", "CargoEntities", ids) // get the first file level only...
	if err == nil {
		project.AppendFilesRef(file.(*CargoEntities.File))
	}
	GetServer().GetEntityManager().saveEntity(project)
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
	entities, errObj := GetServer().GetEntityManager().getEntities("CargoEntities.Project", "CargoEntities", nil)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return projects
	}

	// Here I will parse the list of entities retreived...
	for i := 0; i < len(entities); i++ {
		projects = append(projects, entities[i].(*CargoEntities.Project))
	}
	return projects
}
