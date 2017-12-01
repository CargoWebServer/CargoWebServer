package Server

import (
	"encoding/json"
	"errors"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"reflect"
	"strconv"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
)

type ServiceManager struct {
	// info about connection on smtp server...
	m_services         map[string]Service
	m_servicesLst      []Service
	m_serviceClientSrc map[string]string
	m_serviceServerSrc map[string]string

	// Keep list of action by services.
	m_serviceAction map[string][]*CargoEntities.Action
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
	serviceManager.m_servicesLst = make([]Service, 0)           // Keep the order of intialisation.
	serviceManager.m_serviceClientSrc = make(map[string]string) // Keep the services sources.
	serviceManager.m_serviceServerSrc = make(map[string]string)
	serviceManager.m_serviceAction = make(map[string][]*CargoEntities.Action) // Keep the list of action here.

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

	for i := 0; i < len(this.m_servicesLst); i++ {
		// Initialyse the service.
		this.m_servicesLst[i].initialize()
	}

}

func (this *ServiceManager) getId() string {
	return "ServiceManager"
}

func (this *ServiceManager) start() {
	log.Println("--> Start ServiceManager")

	// I will create new action if there one's...
	for i := 0; i < len(this.m_servicesLst); i++ {
		// register the action inside the service.
		this.registerServiceActions(this.m_servicesLst[i])
		// Generate the service code.
		this.generateActionCode(this.m_servicesLst[i].getId())

	}
	// register itself as service.
	this.registerServiceActions(this)
	this.generateActionCode(this.getId())

	for i := 0; i < len(this.m_servicesLst); i++ {
		// Get the service configuration information.
		config := GetServer().GetConfigurationManager().getServiceConfigurationById(this.m_servicesLst[i].getId())
		if config != nil {
			if config.M_start == true {
				this.m_servicesLst[i].start()
			}
		}
	}
}

func (this *ServiceManager) stop() {
	log.Println("--> Stop ServiceManager")
	for _, service := range this.m_services {
		service.stop()
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
 * Here I will register external services. (WS or TCP)
 */
func (this *ServiceManager) registerServiceContainerActions(config *Config.ServiceConfiguration) {
	// If the action array does not exist.
	if this.m_serviceAction[config.GetId()] == nil {
		this.m_serviceAction[config.GetId()] = make([]*CargoEntities.Action, 0)
	}

	address := GetServer().GetConfigurationManager().GetIpv4() + ":" + strconv.Itoa(config.GetPort())
	if config.GetPort() == GetServer().GetConfigurationManager().GetWsConfigurationServicePort() {
		address = "ws://" + address
	}

	// I will open a connection with the service and get it list of actions.
	conn, err := GetServer().connect(address)

	if err != nil {
		return
	}

	// So here I will request the list of actions.
	id := Utility.RandomUUID()
	method := "GetActionInfos"
	params := make([]*MessageData, 0)

	to := make([]connection, 1)
	to[0] = conn

	// The success callback.
	successCallback_ := func(rspMsg *message, caller_ interface{}) {
		// So here i will get the message value...
		results := rspMsg.msg.Rsp.Results
		for i := 0; i < len(results); i++ {
			infos := make([]map[string]interface{}, 0)
			err := json.Unmarshal(results[i].DataBytes, &infos)
			if err == nil {
				// Here I got the informations..
				for j := 0; j < len(infos); j++ {
					// Represent the id of the interface that contain the
					// the method...
					interfaceId := infos[j]["IID"].(string)
					actionInfos := infos[j]["actions"].([]interface{})

					for k := 0; k < len(actionInfos); k++ {
						actionId := interfaceId + "." + actionInfos[k].(map[string]interface{})["name"].(string)
						actionUuid := CargoEntitiesActionExists(actionId)
						var action *CargoEntities.Action
						if len(actionUuid) == 0 {
							action = new(CargoEntities.Action)
							action.TYPENAME = "CargoEntities.Action"
							action.SetName(actionId)
							action.SetAccessType(CargoEntities.AccessType_Public)

							// Now the documentation.
							var doc string
							for l := 0; l < len(actionInfos[k].(map[string]interface{})["doc"].([]interface{})); l++ {
								doc += actionInfos[k].(map[string]interface{})["doc"].([]interface{})[l].(string) + "\n"
							}

							action.SetDoc(doc)
							actionEntity, _ := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntities().GetUuid(), "M_actions", "CargoEntities.Action", "", action)

							// Now the parameters.
							parameterInfos := actionInfos[k].(map[string]interface{})["parameters"].([]interface{})
							for l := 0; l < len(parameterInfos); l++ {
								parameterInfo := parameterInfos[l]
								parameter := new(CargoEntities.Parameter)
								parameter.SetName(parameterInfo.(map[string]interface{})["name"].(string))
								isArray, _ := strconv.ParseBool(parameterInfo.(map[string]interface{})["isArray"].(string))
								parameter.SetIsArray(isArray)
								parameter.SetType(parameterInfo.(map[string]interface{})["type"].(string))
								GetServer().GetEntityManager().createEntity(action.GetUUID(), "M_parameters", "CargoEntities.Parameter", "", parameter)
								action.SetParameters(parameter)
							}

							// Now the return values.
							resultInfos := actionInfos[k].(map[string]interface{})["results"].([]interface{})
							for l := 0; l < len(resultInfos); l++ {
								resultInfo := resultInfos[l]
								result := new(CargoEntities.Parameter)
								result.UUID = Utility.RandomUUID()
								result.SetName(resultInfo.(map[string]interface{})["name"].(string))
								isArray, _ := strconv.ParseBool(resultInfo.(map[string]interface{})["isArray"].(string))
								result.SetIsArray(isArray)
								result.SetType(resultInfo.(map[string]interface{})["type"].(string))
								GetServer().GetEntityManager().createEntity(action.GetUUID(), "M_results", "CargoEntities.Parameter", "", result)
								action.SetResults(result)
							}

							actionEntity.SaveEntity() // Save the entity
							log.Println("service container action ", action.GetName(), "was created susscessfully with uuid ", action.UUID)
						} else {
							entity, _ := GetServer().GetEntityManager().getEntityByUuid(actionUuid, false)
							action = entity.GetObject().(*CargoEntities.Action)
							log.Println("-->Load service container action ", action.GetName(), " informations.")
						}
						this.m_serviceAction[config.GetId()] = append(this.m_serviceAction[config.GetId()], action)
					}
				}
			}
		}
	}

	// The error callback.
	errorCallback_ := func(errMsg *message, caller interface{}) {
		errStr := errMsg.msg.Err.Message
		log.Println("--------> error: 207 ", errStr)
	}

	rqst, _ := NewRequestMessage(id, method, params, to, successCallback_, nil, errorCallback_, nil)

	GetServer().GetProcessor().m_sendRequest <- rqst

	// Close the connection.
	//conn.Close()
}

/**
 * That function use reflection to create the actions information contain in a
 * given service. The information will be use by role. It must be run at least once
 * the results will by action information in CargoEntities.
 */
func (this *ServiceManager) registerServiceActions(service Service) {

	// Now  I will register action for that service.
	// If the action array does not exist.
	if this.m_serviceAction[service.getId()] == nil {
		this.m_serviceAction[service.getId()] = make([]*CargoEntities.Action, 0)
	}

	// I will use the reflection to retreive method inside the service
	serviceType := reflect.TypeOf(service)

	// Here I will get the documentation information necessary to
	// create the action js code.
	methodsDoc := make(map[string]*doc.Func, 0)
	fset := token.NewFileSet() // positions are relative to fset
	d, err := parser.ParseDir(fset, "./Cargo/Server", nil, parser.ParseComments)
	if err != nil {
		// In that case i dont have access to the server code so i will get
		// action form information in the db.
		query := new(EntityQuery)
		query.TYPENAME = "Server.EntityQuery"
		query.TypeName = "CargoEntities.Action"
		query.Query = `CargoEntities.Action.M_name == /Server.` + service.getId() + `/`

		actions, _ := GetServer().GetEntityManager().getEntities("CargoEntities.Action", query, "CargoEntities", false)
		for i := 0; i < len(actions); i++ {
			action := actions[i].GetObject().(*CargoEntities.Action)
			this.m_serviceAction[service.getId()] = append(this.m_serviceAction[service.getId()], action)
			log.Println("-->Load service ", service.getId(), " action ", action.GetName(), " informations.")
		}
		return // Nothing todo in that case.
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
		methodUuid := CargoEntitiesActionExists(methodName)
		var action *CargoEntities.Action
		if len(methodUuid) == 0 {
			action = new(CargoEntities.Action)
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(methodUuid, false)
			action = entity.GetObject().(*CargoEntities.Action)
		}

		if len(methodUuid) == 0 && (len(methodName) > 0 && !(strings.HasPrefix(method.Name, "New") && (strings.HasSuffix(method.Name, "Entity") || strings.HasSuffix(method.Name, "EntityFromObject")))) {
			action.SetName(methodName)
			m := methodsDoc[methodName[strings.LastIndex(methodName, ".")+1:]]
			if m != nil {
				if len(action.UUID) == 0 {
					action.SetDoc(m.Doc)
					if strings.Index(action.M_doc, "@api ") != -1 { // Only api action are exported...
						// Set the uuid if is not set...
						actionEntity, _ := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntities().GetUuid(), "M_actions", "CargoEntities.Action", "", action)

						// The input
						for j := 0; j < method.Type.NumIn(); j++ {
							// The first paramters is the object itself.
							if j >= 1 {
								in := method.Type.In(j)
								parameter := new(CargoEntities.Parameter)
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

								parameterId := action.GetName() + ":" + parameter.GetName() + ":" + parameter.GetType()
								parameter.UUID = "CargoEntities.Parameter%" + Utility.GenerateUUID(parameterId) // Ok must be random
								GetServer().GetEntityManager().createEntity(action.GetUUID(), "M_parameters", "CargoEntities.Parameter", "", parameter)
								action.SetParameters(parameter)

							}
						}

						// The output
						for j := 0; j < method.Type.NumOut(); j++ {
							out := method.Type.Out(j)
							parameter := new(CargoEntities.Parameter)

							parameter.TYPENAME = "CargoEntities.Parameter"
							parameter.SetType(out.String())
							parameter.SetName("r" + strconv.Itoa(j))
							if strings.HasPrefix(out.String(), "[]") {
								parameter.SetIsArray(true)
							} else {
								parameter.SetIsArray(false)
							}

							parameterId := action.GetName() + ":" + parameter.GetName() + ":" + parameter.GetType()
							parameter.UUID = "CargoEntities.Parameter%" + Utility.GenerateUUID(parameterId) // Ok must be random
							GetServer().GetEntityManager().createEntity(action.GetUUID(), "M_results", "CargoEntities.Parameter", "", parameter)
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

						actionEntity.SaveEntity() // Save the entity
						log.Println("service action ", action.GetName(), "was created susscessfully with uuid ", action.UUID)
					}
				}
			}
		}
		// Export only action with api...
		if strings.Index(action.M_doc, "@api ") != -1 { // Only api action are exported...
			this.m_serviceAction[service.getId()] = append(this.m_serviceAction[service.getId()], action)
			log.Println("-->Load service action ", action.GetName(), " informations.")
		}
	}
}

func (this *ServiceManager) generateActionCode(serviceId string) {

	// Here I will generate the client service class.
	var clientSrc string
	// And the server side code to.
	var serverSrc string

	// Now the server side function...
	serverSrc += "require(\"Cargo/eventHub\")\n"
	serverSrc += "require(\"Cargo/utility\")\n\n"

	var eventTypename = strings.Replace(serviceId, "Manager", "Event", -1)

	// Here I will generate the javascript code use by client side.
	clientSrc = "// ============================= " + serviceId + " ========================================\n"
	clientSrc += "require(\"Cargo/eventHub\")\n"
	clientSrc += "require(\"Cargo/utility\")\n\n"

	serverSrc = clientSrc // same comment.

	clientSrc += "\nvar " + serviceId + " = function(){\n"
	clientSrc += "	if (server == undefined) {\n"
	clientSrc += "		return\n"
	clientSrc += "	}\n"
	clientSrc += "	EventHub.call(this, " + eventTypename + ")\n\n"
	clientSrc += "	return this\n"
	clientSrc += "}\n\n"

	clientSrc += serviceId + ".prototype = new EventHub(null);\n"
	clientSrc += serviceId + ".prototype.constructor = " + serviceId + ";\n\n"

	actions := this.GetServiceActions(serviceId, "", "")

	for i := 0; i < len(actions); i++ {
		action := actions[i]
		name := action.M_name[strings.LastIndex(action.M_name, ".")+1:]
		doc := action.GetDoc()
		if strings.Index(doc, "@api ") != -1 {

			if strings.Index(doc, "@src\n") != -1 {
				// Here the code of the method is defined in the documentation.
				clientSrc += doc[strings.Index(doc, "@src\n")+5:] + "\n"
			} else {
				clientSrc += serviceId + ".prototype." + strings.ToLower(name[0:1]) + name[1:] + " = function("

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
						} else if strings.HasPrefix(paramTypeName, "interface") {
							clientSrc += "	params.push(createRpcData(" + param.GetName() + ", \"JSON_STR\", \"" + param.GetName() + "\"))\n"
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
							clientSrc += "	params.push(createRpcData(" + param.GetName() + ", \"JSON_STR\", \"" + param.GetName() + "\", \"" + param.GetType() + "\"))\n"
						}
					}
				}

				// Now will generate the code for executeJsFunction.
				clientSrc += "\n	server.executeJsFunction(\n"
				clientSrc += "	\"" + serviceId + name + "\",\n"
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
								clientSrc += "			if(caller.results[0] == null){\n"
								clientSrc += "				if(caller.successCallback != undefined){\n"
								clientSrc += "					caller.successCallback(entities, caller.caller)\n"
								clientSrc += "					caller.successCallback = undefined\n"
								clientSrc += "				}\n"
								clientSrc += "				return\n"
								clientSrc += "			}\n"
								clientSrc += "			for (var i = 0; i < caller.results[0].length; i++) {\n"
								clientSrc += "				var entity = eval(\"new \" + prototype.TypeName + \"()\")\n"
								clientSrc += "				if (i == caller.results[0].length - 1) {\n"
								clientSrc += "					entity.initCallback = function (caller) {\n"
								clientSrc += "						return function (entity) {\n"
								clientSrc += "							server.entityManager.setEntity(entity)\n"
								clientSrc += "							if(caller.successCallback != undefined){\n"
								clientSrc += "								caller.successCallback(entities, caller.caller)\n"
								clientSrc += "								caller.successCallback = undefined\n"
								clientSrc += "							}\n"
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
								clientSrc += "				if(caller.successCallback != undefined){\n"
								clientSrc += "					caller.successCallback(entities[caller.results[0].UUID], caller.caller)\n"
								clientSrc += "					caller.successCallback=undefined\n"
								clientSrc += "				}\n"
								clientSrc += "				return // break it here.\n"
								clientSrc += "			}\n\n"

								clientSrc += "			var entity = eval(\"new \" + prototype.TypeName + \"()\")\n"
								clientSrc += "				entity.initCallback = function () {\n"
								clientSrc += "					return function (entity) {\n"
								clientSrc += "					if(caller.successCallback != undefined){\n"
								clientSrc += "						caller.successCallback(entity, caller.caller)\n"
								clientSrc += "						caller.successCallback=undefined\n"
								clientSrc += "					}\n"
								clientSrc += "				}\n"
								clientSrc += "			}(caller)\n"
								clientSrc += "			entity.init(caller.results[0])\n"
							}

							clientSrc += "			},\n"
							clientSrc += "			function (errMsg, caller) { // Error Callback\n"
							clientSrc += "				if(caller.errorCallback != undefined){\n"
							clientSrc += "					caller.errorCallback(errMsg, caller.caller)\n"
							clientSrc += "					caller.errorCallback = undefined\n"
							clientSrc += "				}\n"
							clientSrc += "				server.errorManager.onError(errMsg)\n"
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
							clientSrc += "		if(caller.successCallback != undefined){\n"
							clientSrc += "			caller.successCallback(results, caller.caller)\n"
							clientSrc += "			caller.successCallback=undefined\n"
							clientSrc += "		}\n"
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
					clientSrc += "		if(caller.errorCallback != undefined){\n"
					clientSrc += "			caller.errorCallback(errMsg, caller.caller)\n"
					clientSrc += "			caller.errorCallback = undefined\n"
					clientSrc += "		}\n"
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

		serverSrc += "function " + serviceId + name + "("
		params_ := ""
		if action.M_parameters != nil {
			// The last tow parameters are sessionId and message Id
			for j := 0; j < len(action.M_parameters)-2; j++ {
				params_ += action.M_parameters[j].GetName()
				if j < len(action.M_parameters)-3 {
					params_ += ", "
				}
			}
		}
		serverSrc += params_
		serverSrc += "){\n"

		if len(params_) > 0 {
			params_ += ", "
		}

		params_ += "messageId, sessionId"

		// The content of the action code will depend of the parameter output.
		if len(action.M_results) > 0 {
			// It can be an array or not...
			if action.M_results[0].IsArray() {
				serverSrc += "	var " + action.M_results[0].M_name + " = []\n"
			} else {
				serverSrc += "	var " + action.M_results[0].M_name + " = null\n"
			}
			serverSrc += "	" + action.M_results[0].M_name + " = GetServer().Get" + serviceId + "()." + name + "(" + params_ + ")\n"
			serverSrc += "	return " + action.M_results[0].M_name + "\n"
		} else {
			// Here I will simply call the method on the service object..
			serverSrc += "	GetServer().Get" + serviceId + "()." + name + "(" + params_ + ")\n"
		}

		serverSrc += "}\n\n"
	}

	// Keep the service javaScript code in the map.
	serviceManager.m_serviceClientSrc[serviceId] = clientSrc
	serviceManager.m_serviceServerSrc[serviceId] = serverSrc

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
	//for i := 0; i < len(this.m_serviceAction[serviceName]); i++ {
	//log.Println(this.m_serviceAction[serviceName][i])
	//}
	return this.m_serviceAction[serviceName]
}
