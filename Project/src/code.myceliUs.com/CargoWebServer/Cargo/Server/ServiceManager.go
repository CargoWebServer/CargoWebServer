package Server

import (
	"errors"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

const (
	serviceContainer = "ServiceContainer"
)

type ServiceManager struct {
	// info about connection on smtp server...
	m_services             map[string]Service
	m_servicesLst          []Service
	m_serviceContainerCmds []*exec.Cmd
	m_serviceSrc           map[string]string
}

var serviceManager *ServiceManager

func (this *Server) GetServiceManager() *ServiceManager {

	if serviceManager == nil {
		serviceManager = newServiceManager()
	}

	return serviceManager
}

/**
 * Singleton that return reference to the Smtp service.
 */
func newServiceManager() *ServiceManager {

	serviceManager := new(ServiceManager)

	// Here I will initialyse the optional services...
	serviceManager.m_services = make(map[string]Service)
	serviceManager.m_servicesLst = make([]Service, 0)     // Keep the order of intialisation.
	serviceManager.m_serviceSrc = make(map[string]string) // Keep the services sources.

	return serviceManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Do intialysation stuff here.
 */
func (this *ServiceManager) initialize() {
	// Here I will start the c++ service container...
	log.Println("--> Initialize ServiceManager")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)

	// Set the service container port.
	GetServer().GetConfigurationManager().setServiceConfiguration(serviceContainer, GetServer().GetConfigurationManager().GetWsConfigurationServicePort())

	for i := 0; i < len(this.m_servicesLst); i++ {
		// Initialyse the service.
		this.m_servicesLst[i].initialize()
	}
}

func (this *ServiceManager) getId() string {
	return "ServiceManager"
}

func (this *ServiceManager) startServiceContainer(name string, port int) error {
	// The first step will be to start the service manager.
	tcp_serviceContainerPath := GetServer().GetConfigurationManager().GetBinPath() + "/" + name
	if runtime.GOOS == "windows" {
		tcp_serviceContainerPath += ".exe"
	}

	// Set the command
	cmd := exec.Command(tcp_serviceContainerPath)
	cmd.Args = append(cmd.Args, strconv.Itoa(port))

	// Call it...
	err := cmd.Start()
	if err != nil {
		log.Println("---> fail to start the service container!")
		return err
	}

	// the command succed here.
	serviceManager.m_serviceContainerCmds = append(serviceManager.m_serviceContainerCmds, cmd)

	return nil
}

func (this *ServiceManager) start() {

	log.Println("--> Start ServiceManager")

	// I will create new action if there one's...
	for i := 0; i < len(this.m_servicesLst); i++ {
		// register the action inside the service.
		this.registerActions(this.m_servicesLst[i])
	}

	// register itself as service.
	this.registerActions(this)

	for i := 0; i < len(this.m_servicesLst); i++ {
		// Get the service configuration information.
		config := GetServer().GetConfigurationManager().getServiceConfigurationById(this.m_servicesLst[i].getId())
		if config != nil {
			if config.M_start == true {
				this.m_servicesLst[i].start()
			}

		}
	}

	// TCP
	// this.startServiceContainer("CargoServiceContainer_TCP", GetServer().GetConfigurationManager().GetTcpConfigurationServicePort())

	// WS
	// this.startServiceContainer("CargoServiceContainer_WS", GetServer().GetConfigurationManager().GetWsConfigurationServicePort())

}

func (this *ServiceManager) stop() {
	log.Println("--> Stop ServiceManager")
	for _, service := range this.m_services {
		service.stop()
	}

	// Stop the services containers...
	for i := 0; i < len(serviceManager.m_serviceContainerCmds); i++ {
		if serviceManager.m_serviceContainerCmds[i].Process != nil {
			serviceManager.m_serviceContainerCmds[i].Process.Kill()
		}
	}
}

/**
 * Register a new service.
 */
func (this *ServiceManager) registerService(service Service) {
	this.m_services[service.getId()] = service
	this.m_servicesLst = append(this.m_servicesLst, service)
}

/**
 * That function use reflection to create the actions information contain in a
 * given service. The information will be use by role.
 */
func (this *ServiceManager) registerActions(service Service) {

	// I will use the reflection to reteive method inside the service
	serviceType := reflect.TypeOf(service)

	// Here I will get the documentation information necessary to
	// create the action js code.
	methodsDoc := make(map[string]*doc.Func, 0)
	fset := token.NewFileSet() // positions are relative to fset
	d, err := parser.ParseDir(fset, "./Cargo/Server", nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}

	for _, f := range d {
		p := doc.New(f, "./", 0)
		for _, t := range p.Types {
			if t.Name == service.getId() {
				for _, m := range t.Methods {
					methodsDoc[m.Name] = m
				}
				break
			}
		}
	}

	// Now I will print it list of function.
	for i := 0; i < serviceType.NumMethod(); i++ {
		// I will try to find if the action was register
		method := serviceType.Method(i)
		methodName := strings.Replace(serviceType.String(), "*", "", -1) + "." + method.Name
		metodUuid := CargoEntitiesActionExists(methodName)
		if len(metodUuid) == 0 && !(strings.HasPrefix(method.Name, "New") && (strings.HasSuffix(method.Name, "Entity") || strings.HasSuffix(method.Name, "EntityFromObject"))) {
			action := new(CargoEntities.Action)
			action.SetName(methodName)
			m := methodsDoc[methodName[strings.LastIndex(methodName, ".")+1:]]
			if m != nil {
				action.SetDoc(m.Doc)

				if strings.Index(action.M_doc, "@api ") != -1 { // Only api action are exported...
					// Set the uuid
					GetServer().GetEntityManager().NewCargoEntitiesActionEntity(GetServer().GetEntityManager().getCargoEntities().GetUuid(), "", action)

					// The input
					for j := 0; j < method.Type.NumIn(); j++ {
						// The first paramters is the object itself.
						if j >= 1 {
							in := method.Type.In(j)
							parameter := new(CargoEntities.Parameter)
							parameter.UUID = "CargoEntities.Parameter%" + Utility.RandomUUID() // Ok must be random
							parameter.TYPENAME = "CargoEntities.Parameter"
							parameter.SetType(in.String())

							if m != nil {
								field := m.Decl.Type.Params.List[j-1]
								parameter.SetName(field.Names[0].String())
							} else {
								parameter.SetName("p" + strconv.Itoa(len(action.M_parameters)))
							}

							if strings.HasPrefix(in.String(), "[]") {
								parameter.SetIsArray(true)
							} else {
								parameter.SetIsArray(false)
							}
							action.SetParameters(parameter)
						}
					}

					// The output
					for j := 0; j < method.Type.NumOut(); j++ {
						out := method.Type.Out(j)
						parameter := new(CargoEntities.Parameter)
						parameter.UUID = "CargoEntities.Parameter%" + Utility.RandomUUID() // Ok must be random
						parameter.TYPENAME = "CargoEntities.Parameter"
						parameter.SetType(out.String())
						parameter.SetName("r" + strconv.Itoa(j))
						if strings.HasPrefix(out.String(), "[]") {
							parameter.SetIsArray(true)
						} else {
							parameter.SetIsArray(false)
						}
						action.SetResults(parameter)
					}

					action.SetAccessType(CargoEntities.AccessType_Public)

					// Now I will set the access type of the action before save it.
					if strings.Index(action.M_doc, "@scope {hidden}") != -1 {
						action.SetAccessType(CargoEntities.AccessType_Hidden)
					}

					if strings.Index(action.M_doc, "@scope {public}") != -1 {
						action.SetAccessType(CargoEntities.AccessType_Public)
					}

					if strings.Index(action.M_doc, "@scope {restricted}") != -1 {
						action.SetAccessType(CargoEntities.AccessType_Restricted)
					}

					// apend it to the entities action.
					action.SetEntitiesPtr(GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities))
					GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities).SetActions(action)

					// I will append the action into the admin role that has all permission.
					adminRoleUuid := CargoEntitiesRoleExists("adminRole")
					if len(adminRoleUuid) > 0 {
						adminRoleEntity, _ := GetServer().GetEntityManager().getEntityByUuid(adminRoleUuid, false)
						if action.GetAccessType() != CargoEntities.AccessType_Hidden {
							adminRoleEntity.GetObject().(*CargoEntities.Role).SetActions(action)
							adminRoleEntity.SaveEntity()
						}
					}

					guestRoleUuid := CargoEntitiesRoleExists("guestRole")
					if len(guestRoleUuid) > 0 {
						guestRoleEntity, _ := GetServer().GetEntityManager().getEntityByUuid(guestRoleUuid, false)
						if action.GetAccessType() == CargoEntities.AccessType_Public {
							guestRoleEntity.GetObject().(*CargoEntities.Role).SetActions(action)
							guestRoleEntity.SaveEntity()
						}
					}
				}
			}
		}
	}

	GetServer().GetEntityManager().getCargoEntities().SaveEntity()

	// Here I will generate the client service class.
	var clientSrc string
	var eventTypename = strings.Replace(service.getId(), "Manager", "Event", -1)

	// Here I will generate the javascript code use by client side.
	clientSrc = "// ============================= " + service.getId() + " ========================================\n"
	clientSrc += "\nvar " + service.getId() + " = function(){\n"
	clientSrc += "	if (server == undefined) {\n"
	clientSrc += "		return\n"
	clientSrc += "	}\n"
	clientSrc += "	EventHub.call(this, " + eventTypename + ")\n\n"
	clientSrc += "	return this\n"
	clientSrc += "}\n\n"

	clientSrc += service.getId() + ".prototype = new EventHub(null);\n"
	clientSrc += service.getId() + ".prototype.constructor = " + service.getId() + ";\n\n"

	actions := this.GetServiceActions(service.getId(), "", "")

	for i := 0; i < len(actions); i++ {
		action := actions[i]
		name := action.M_name[strings.LastIndex(action.M_name, ".")+1:]
		doc := action.GetDoc()
		if strings.Index(doc, "@api ") != -1 {
			if strings.Index(doc, "@src\n") != -1 {
				// Here the code of the method is defined in the documentation.
				clientSrc += doc[strings.Index(doc, "@src\n")+5:] + "\n"
			} else {
				clientSrc += service.getId() + ".prototype." + strings.ToLower(name[0:1]) + name[1:] + " = function("

				// Now the parameters...
				if action.M_parameters != nil {
					// The last tow parameters are sessionId and message Id
					for j := 0; j < len(action.M_parameters)-2; j++ {
						clientSrc += action.M_parameters[j].GetName()
						if j < len(action.M_parameters)-2 {
							clientSrc += ", "
						}
					}
				}

				// I will look for callback function.
				callbacks := make([]string, 0)

				if strings.Index(doc, "@param {callback} successCallback") != -1 {
					callbacks = append(callbacks, "successCallback")
				}

				if strings.Index(doc, "@param {callback} progressCallback") != -1 {
					callbacks = append(callbacks, "progressCallback")
				}

				if strings.Index(doc, "@param {callback} errorCallback") != -1 {
					callbacks = append(callbacks, "errorCallback")
				}

				for j := 0; j < len(callbacks); j++ {
					clientSrc += callbacks[j]
					if j < len(callbacks)-1 {
						clientSrc += ", "
					}
				}
				if len(callbacks) > 0 {
					clientSrc += ", "
				}
				clientSrc += "caller){\n"

				// Here I will generate the content of the function.
				if action.M_parameters != nil {
					clientSrc += "	var params = []\n"
					for j := 0; j < len(action.M_parameters)-2; j++ {
						param := action.M_parameters[j]
						paramTypeName := param.GetType()
						if paramTypeName == "string" {
							clientSrc += "	params.push(createRpcData(" + param.GetName() + ", \"STRING\", \"" + param.GetName() + "\"))\n"
						} else if strings.HasPrefix(paramTypeName, "int") {
							clientSrc += "	params.push(createRpcData(" + param.GetName() + ", \"INTEGER\", \"" + param.GetName() + "\"))\n"
						} else if paramTypeName == "bool" {
							clientSrc += "	params.push(createRpcData(" + param.GetName() + ", \"BOOLEAN\", \"" + param.GetName() + "\"))\n"
						} else if paramTypeName == "double" || strings.HasPrefix(paramTypeName, "float") {
							clientSrc += "	params.push(createRpcData(" + param.GetName() + ", \"DOUBLE\", \"" + param.GetName() + "\"))\n"
						} else if paramTypeName == "[]unit8" || paramTypeName == "[]byte" {
							clientSrc += "	params.push(createRpcData(" + param.GetName() + ", \"BYTES\", \"" + param.GetName() + "\"))\n"
						} else {
							// Array or Object or array of object...
							clientSrc += "	params.push(createRpcData(" + param.GetName() + ", \"JSON_STR\", \"" + param.GetName() + "\"))\n"
						}
					}
				}

				// Now will generate the code for executeJsFunction.
				clientSrc += "\n	server.executeJsFunction(\n"
				clientSrc += "	\"" + service.getId() + name + "\",\n"
				clientSrc += "	params, \n"
				caller := "{"
				if Utility.Contains(callbacks, "progressCallback") {
					// Set the progress callback.
					clientSrc += "	function (index, total, caller) { // Progress callback\n"
					clientSrc += "		caller.progressCallback(index, total, caller.caller)\n"
					clientSrc += "	},\n"

					// Set the caller.
					caller += "\"progressCallback\":progressCallback, "
				} else {
					clientSrc += "	undefined, //progress callback\n"
				}

				if Utility.Contains(callbacks, "successCallback") {
					// Set the progress callback.
					clientSrc += "	function (results, caller) { // Success callback\n"
					if len(action.M_results) > 0 {
						typeName := action.M_results[0].M_type
						isArray := strings.HasPrefix(typeName, "[]")
						typeName = strings.Replace(typeName, "[]", "", -1)
						typeName = strings.Replace(typeName, "*", "", -1)
						// Now I will test if the type is an entity...
						if strings.Index(typeName, ".") > -1 {
							// Here I got an entity...

							clientSrc += "		server.entityManager.getEntityPrototype(\"" + typeName + "\", \"" + typeName[0:strings.Index(typeName, ".")] + "\",\n"
							clientSrc += "			function (prototype, caller) { // Success Callback\n"
							// in case of an array...
							if isArray {
								clientSrc += "			var entities = []\n"
								clientSrc += "			for (var i = 0; i < caller.results[0].length; i++) {\n"
								clientSrc += "				var entity = eval(\"new \" + prototype.TypeName + \"()\")\n"
								clientSrc += "				if (i == caller.results[0].length - 1) {\n"
								clientSrc += "					entity.initCallback = function (caller) {\n"
								clientSrc += "						return function (entity) {\n"
								clientSrc += "							server.entityManager.setEntity(entity)\n"
								clientSrc += "							caller.successCallback(entities, caller.caller)\n"
								clientSrc += "						}\n"
								clientSrc += "					} (caller)\n"
								clientSrc += "				}else{\n"
								clientSrc += "					entity.initCallback = function (entity) {\n"
								clientSrc += "						server.entityManager.setEntity(entity)\n"
								clientSrc += "					}\n"
								clientSrc += "				}\n"
								clientSrc += "				entities.push(entity)\n"
								clientSrc += "				entity.init(caller.results[0][i])\n"
								clientSrc += "			}\n"
							} else {
								// In case of a regular entity.
								clientSrc += "			if (caller.results[0] == null) {\n"
								clientSrc += "				return\n"
								clientSrc += "			}\n"

								// In case of existing entity.
								clientSrc += "			if (entities[caller.results[0].UUID] != undefined && caller.results[0].TYPENAME == caller.results[0].__class__) {\n"
								clientSrc += "				caller.successCallback(entities[caller.results[0].UUID], caller.caller)\n"
								clientSrc += "				return // break it here.\n"
								clientSrc += "			}\n\n"

								clientSrc += "			var entity = eval(\"new \" + prototype.TypeName + \"()\")\n"
								clientSrc += "				entity.initCallback = function () {\n"
								clientSrc += "					return function (entity) {\n"
								clientSrc += "						caller.successCallback(entity, caller.caller)\n"
								clientSrc += "				}\n"
								clientSrc += "			}(caller)\n"
								clientSrc += "			entity.init(caller.results[0])\n"
							}

							clientSrc += "			},\n"
							clientSrc += "			function (errMsg, caller) { // Error Callback\n"
							clientSrc += "				caller.errorCallback(errMsg, caller.caller)\n"
							clientSrc += "			},\n"
							caller := "{ \"caller\": caller.caller"

							if Utility.Contains(callbacks, "progressCallback") {
								caller += ", \"progressCallback\": caller.progressCallback"
							}
							caller += ", \"successCallback\": caller.successCallback, \"errorCallback\": caller.errorCallback, \"results\": results }\n"
							clientSrc += "			" + caller

							clientSrc += "		)\n"

						} else {
							// Here I got a regulat type.
							clientSrc += "		caller.successCallback(results, caller.caller)\n"
						}
					} else {
						clientSrc += "		caller.successCallback(results, caller.caller)\n"
					}

					clientSrc += "	},\n"
					// Set the caller.
					caller += "\"successCallback\":successCallback, "
				} else {
					clientSrc += "	undefined, //success callback\n"
				}

				if Utility.Contains(callbacks, "errorCallback") {
					clientSrc += "	function (errMsg, caller) { // Error callback\n"
					clientSrc += "		caller.errorCallback(errMsg, caller.caller)\n"
					clientSrc += "		server.errorManager.onError(errMsg)\n"
					clientSrc += "	},"
					// Set the caller.
					caller += "\"errorCallback\":errorCallback, "
				} else {
					clientSrc += "	undefined, //error callback\n"
				}

				caller += "\"caller\": caller}"
				clientSrc += caller + ")\n"
				clientSrc += "}\n\n"
			}
		}
	}

	// Keep the service javaScript code in the map.
	serviceManager.m_serviceSrc[service.getId()] = clientSrc
}

func (this *ServiceManager) registerAction(methodName string, parameters []interface{}, results []interface{}) (*CargoEntities.Action, error) {

	// I will try to find if the action was register
	methodUuid := CargoEntitiesActionExists(methodName)
	if len(methodUuid) > 0 {
		return nil, errors.New("The method whit name '" + methodName + "' already exist!")
	}

	// The converted type.
	var parameters_ []*CargoEntities.Parameter
	var results_ []*CargoEntities.Parameter

	// Initialyse the parameters object of not already intialyse.
	values, err := Utility.InitializeStructures(parameters, "CargoEntities.Parameter")

	if err == nil {
		parameters_ = values.Interface().([]*CargoEntities.Parameter)
	} else {
		for i := 0; i < len(parameters); i++ {
			parameters_ = append(parameters_, parameters[i].(*CargoEntities.Parameter))
		}
	}

	values, err = Utility.InitializeStructures(results, "CargoEntities.Parameter")
	if err == nil {
		results_ = values.Interface().([]*CargoEntities.Parameter)
	} else {
		for i := 0; i < len(results); i++ {
			results_ = append(results_, results[i].(*CargoEntities.Parameter))
		}
	}

	action := new(CargoEntities.Action)
	action.SetName(methodName)

	// Set the uuid
	GetServer().GetEntityManager().NewCargoEntitiesActionEntity(GetServer().GetEntityManager().getCargoEntities().GetUuid(), "", action)

	// The input
	for j := 0; j < len(parameters_); j++ {
		action.SetParameters(parameters_[j])
	}

	// The output
	for j := 0; j < len(results_); j++ {
		action.SetResults(results_[j])
	}

	// Restricted by default.
	action.SetAccessType(CargoEntities.AccessType_Restricted)

	// apend it to the entities action.
	action.SetEntitiesPtr(GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities))
	GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities).SetActions(action)

	GetServer().GetEntityManager().getCargoEntities().SaveEntity()

	return action, nil

}

////////////////////////////////////////////////////////////////////////////////
// Api
////////////////////////////////////////////////////////////////////////////////

// @api 1.0
// Event handler function.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
//ServiceManager.prototype.onEvent = function (evt) {
//    EventHub.prototype.onEvent.call(this, evt)
//}
func (this *ServiceManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

// @api 1.0
// That function is run on the server side to register action from javascript
// files in the script folder.
// @param {string} name The name of the action to register.
// @param {[]interface{}} parameters The list of action parameters (*CargoEntities.Parameter).
// @param {[]interface{}} results The list of action return value (*CargoEntities.Parameter).
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *ServiceManager) RegisterAction(name string, parameters []interface{}, results []interface{}, messageId string, sessionId string) *CargoEntities.Action {
	action, err := this.registerAction(name, parameters, results)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), TYPENAME_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	return action
}

// @api 1.0
// Get the list of actions for a given service.
// @param {string} serviceName The name of the service
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *ServiceManager) GetServiceActions(serviceName string, messageId string, sessionId string) []*CargoEntities.Action {
	actionsEntity, _ := GetServer().GetEntityManager().getEntitiesByType("CargoEntities.Action", "", "CargoEntities", false)
	actions := make([]*CargoEntities.Action, 0)
	for i := 0; i < len(actionsEntity); i++ {
		action := actionsEntity[i].GetObject().(*CargoEntities.Action)
		if strings.Contains(action.M_name, serviceName) {
			actions = append(actions, action)
		}
	}
	return actions
}
