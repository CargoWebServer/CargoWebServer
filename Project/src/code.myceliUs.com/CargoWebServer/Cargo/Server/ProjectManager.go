/**
 *
 */
package Server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
	"github.com/yosssi/gohtml"
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
					projectEntity, err := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntities(), "M_entities", project)

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
// Compile TypeScrypt files
// @param {string} uuid The project uuid.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Account} The new registered account.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *ProjectManager) Compile(uuid string, sessionId string, messageId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	// so here I will cast the entity.
	entity, _ := GetServer().GetEntityManager().getEntityByUuid(uuid)
	dir := entity.(*CargoEntities.File)

	// Synchronize the content of project.
	path := GetServer().GetConfigurationManager().GetApplicationDirectoryPath() + "/" + dir.GetPath() + dir.GetName()

	for i := 0; i < len(dir.GetFiles()); i++ {
		if dir.GetFiles()[i].GetName() == "tsconfig.json" {
			// In that case I will compile the typeScript.
			cmd := exec.Command("tsc")
			cmd.Dir = path

			// Call it...
			_, err := cmd.Output()
			if err != nil {
				errObj := NewError(Utility.FileLine(), "RUN_CMD_ERROR", SERVER_ERROR_CODE, err)
				GetServer().reportErrorMessage(messageId, sessionId, errObj)
				return
			}

		} else if dir.GetFiles()[i].GetName() == "package.json" {
			cmd := exec.Command("npm")
			// install package as needed.
			cmd.Args = append(cmd.Args, "install")
			cmd.Dir = path

			// Call it...
			_, err := cmd.Output()
			if err != nil {
				errObj := NewError(Utility.FileLine(), "RUN_CMD_ERROR", SERVER_ERROR_CODE, err)
				GetServer().reportErrorMessage(messageId, sessionId, errObj)
				return
			}
		} else if dir.GetFiles()[i].GetName() == "pbiviz.json" {

			// Here I will generate the package.
			cmd := exec.Command("pbiviz")
			cmd.Dir = path
			cmd.Args = append(cmd.Args, "package")
			// Call it...
			_, err := cmd.Output()
			if err != nil {
				errObj := NewError(Utility.FileLine(), "RUN_CMD_ERROR", SERVER_ERROR_CODE, err)
				GetServer().reportErrorMessage(messageId, sessionId, errObj)
				return
			}

			// Here I will read the pbiviz.json file.
			b, _ := ioutil.ReadFile(path + "/pbiviz.json")
			pbiviz := make(map[string]interface{}, 0)
			json.Unmarshal(b, &pbiviz)

			// Now I will create the index.html file and I will include dependencie.
			index := `<html lang="en">`
			index += `<head>`
			index += `	<meta charset="utf-8">`
			index += `	<title>` + pbiviz["visual"].(map[string]interface{})["displayName"].(string) + `</title>`
			index += `	<meta charset="utf-8">`

			index += `	<meta name="author" content="` + pbiviz["author"].(map[string]interface{})["name"].(string) + `">`
			index += `	<meta name="version" content="` + pbiviz["visual"].(map[string]interface{})["version"].(string) + `">`

			// Now I will read the tsconfig.json to get the out file.
			b, _ = ioutil.ReadFile(path + "/tsconfig.json")
			tsconfig := make(map[string]interface{}, 0)
			json.Unmarshal(b, &tsconfig)

			out := tsconfig["compilerOptions"].(map[string]interface{})["out"].(string)
			out = strings.Replace(out, "./", "", -1)
			out = out[0:strings.Index(out, "/")]
			out = path + "/" + out

			// The sytle sheets.
			styleSheets := Utility.GetFilePathsByExtension(out, "css")

			for j := 0; j < len(styleSheets); j++ {
				index += `	<link rel="stylesheet" href="` + strings.Replace(styleSheets[j], GetServer().GetConfigurationManager().GetApplicationDirectoryPath(), "", -1) + `">`
			}

			// The container files...
			index += `	<link rel="stylesheet" href="/pbi/css/container.css">`

			index += `</head>`

			// Now the souce files.
			index += `<body onload="display();">`

			for j := 0; j < len(pbiviz["externalJS"].([]interface{})); j++ {
				index += `	<script src="` + pbiviz["externalJS"].([]interface{})[j].(string) + `"></script>`
			}

			jsFiles := Utility.GetFilePathsByExtension(out, "js")
			for j := 0; j < len(jsFiles); j++ {
				index += `	<script src="` + strings.Replace(jsFiles[j], GetServer().GetConfigurationManager().GetApplicationDirectoryPath(), "", -1) + `"></script>`
			}

			// I will now create the file that will display the component in the html file.
			jsFile := "/** Display the visual **/\n"
			jsFile += "function display(){\n"

			jsFile += "	var div = new Element(document.body, {\"tag\":\"div\", \"class\":\"pbi-container visual-" + pbiviz["visual"].(map[string]interface{})["guid"].(string) + "\"}); \n"
			jsFile += "	var host = new VisualHost();\n"
			jsFile += "	var options = new VisualConstructorOptions(div.element, host);\n"
			jsFile += "	var visual = new powerbi.extensibility.visual." + pbiviz["visual"].(map[string]interface{})["visualClassName"].(string) + "(options);\n\n"
			jsFile += "	// Init the container.\n"
			jsFile += "	initContainer(div);\n\n"

			// The update event.
			jsFile += "	window.onresize = function(visual, div){\n"
			jsFile += "		return function(){\n"
			jsFile += "			visual.update({\"viewport\":{\"width\":div.offsetWidth, \"height\":div.offsetHeight}})\n"
			jsFile += "		}\n"
			jsFile += "	}(visual, div.element)\n"

			jsFile += "	// fire resize event at start...\n"
			jsFile += "	var evt = document.createEvent('HTMLEvents');\n"
			jsFile += "	evt.initEvent('resize', true, false);\n"
			jsFile += "	div.element.dispatchEvent(evt);\n\n"
			jsFile += "}\n"

			// Save the js file
			Utility.WriteStringToFile(path+"/"+pbiviz["visual"].(map[string]interface{})["name"].(string)+".js", jsFile)

			index += `	<script src="/Cargo/js/utility.js"></script>`
			index += `	<script src="/Cargo/js/element.js"></script>`
			index += `	<script src="/pbi/js/visual.js"></script>`
			index += `	<script src="/pbi/js/container.js"></script>`

			index += `	<script src="` + pbiviz["visual"].(map[string]interface{})["name"].(string) + ".js" + `"></script>`
			index += `</body>`
			index += `</html>`

			Utility.WriteStringToFile(path+"/index.html", gohtml.Format(index))

		}
	}
	// Create the newly file if there is one...
	GetServer().GetFileManager().synchronize(path)
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
