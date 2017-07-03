package Server

import (
	//	"errors"
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
}

var serviceManager *ServiceManager

func (this *Server) GetServiceManager() *ServiceManager {

	if serviceManager == nil {
		serviceManager = newServiceManager()
		/*
			// Here I will append action id's in there respective map.
			/////////////////////////  hidden /////////////////////////
			serviceManager.hidden = append(serviceManager.hidden, "Server.EventManager.Lock")
			serviceManager.hidden = append(serviceManager.hidden, "Server.EventManager.Unlock")
			serviceManager.hidden = append(serviceManager.hidden, "Server.EventManager.AddEventListener")
			serviceManager.hidden = append(serviceManager.hidden, "Server.EventManager.RemoveEventListener")
			serviceManager.restricted = append(serviceManager.restricted, "Server.EventManager.BroadcastEvent")
			serviceManager.restricted = append(serviceManager.restricted, "Server.EventManager.BroadcastEventTo")

			serviceManager.hidden = append(serviceManager.hidden, "Server.DataManager.Lock")
			serviceManager.hidden = append(serviceManager.hidden, "Server.DataManager.RLock")
			serviceManager.hidden = append(serviceManager.hidden, "Server.DataManager.RUnlock")
			serviceManager.hidden = append(serviceManager.hidden, "Server.DataManager.RLocker")
			serviceManager.hidden = append(serviceManager.hidden, "Server.DataManager.Unlock")

			serviceManager.hidden = append(serviceManager.hidden, "Server.EntityManager.RLock")
			serviceManager.hidden = append(serviceManager.hidden, "Server.EntityManager.RUnlock")
			serviceManager.hidden = append(serviceManager.hidden, "Server.EntityManager.RLocker")
			serviceManager.hidden = append(serviceManager.hidden, "Server.EntityManager.Lock")
			serviceManager.hidden = append(serviceManager.hidden, "Server.EntityManager.Unlock")
			serviceManager.hidden = append(serviceManager.hidden, "Server.EntityManager.InitEntity")

			serviceManager.hidden = append(serviceManager.hidden, "Server.SchemaManager.GetFieldsFieldsType")

			/////////////////////////  public /////////////////////////

			// Configuration manager
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetApplicationDirectoryPath")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetBinPath")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetDataPath")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetDefinitionsPath")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetHostName")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetIpv4")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetQueriesPath")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetSchemasPath")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetServerPort")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetScriptPath")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetServicePort")
			serviceManager.public = append(serviceManager.public, "Server.ConfigurationManager.GetTmpPath")

			// Session manager
			serviceManager.public = append(serviceManager.public, "Server.SessionManager.Logout")
			serviceManager.public = append(serviceManager.public, "Server.SessionManager.Login")
			serviceManager.public = append(serviceManager.public, "Server.SessionManager.GetActiveSessions")
			serviceManager.public = append(serviceManager.public, "Server.SessionManager.GetActiveSessionByAccountId")
			serviceManager.public = append(serviceManager.public, "Server.SessionManager.GetActiveSessionById")
			serviceManager.public = append(serviceManager.public, "Server.SessionManager.UpdateSessionState")

			// Data manager
			serviceManager.public = append(serviceManager.public, "Server.DataManager.Connect")
			serviceManager.public = append(serviceManager.public, "Server.DataManager.Ping")
			serviceManager.public = append(serviceManager.public, "Server.DataManager.Create")
			serviceManager.public = append(serviceManager.public, "Server.DataManager.Read")
			serviceManager.public = append(serviceManager.public, "Server.DataManager.Update")
			serviceManager.public = append(serviceManager.public, "Server.DataManager.Delete")

			// Entity manager
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.GenerateEntityUUID")
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.GetEntityPrototype")
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.GetDerivedEntityPrototypes")
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.GetEntityPrototypes")
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.GetEntityLnks")
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.GetObjectById")
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.GetObjectByUuid")
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.GetObjectsByType")
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.CreateEntity")
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.SaveEntity")
			serviceManager.public = append(serviceManager.public, "Server.EntityManager.RemoveEntity")

			// Account manager
			serviceManager.public = append(serviceManager.public, "Server.AccountManager.Me")
			serviceManager.public = append(serviceManager.public, "Server.AccountManager.GetUserById")
			serviceManager.public = append(serviceManager.public, "Server.AccountManager.GetAccountById")
			serviceManager.public = append(serviceManager.public, "Server.AccountManager.Register")

			// Security manager
			// Must be accessible to non admin role...
			serviceManager.public = append(serviceManager.public, "Server.SecurityManager.ChangeAdminPassword")
			serviceManager.public = append(serviceManager.public, "Server.SecurityManager.GetRole")
			serviceManager.public = append(serviceManager.public, "Server.SecurityManager.HasAccount")
			serviceManager.public = append(serviceManager.public, "Server.SecurityManager.HasAction")

			// LDAP manager
			serviceManager.public = append(serviceManager.public, "Server.LdapManager.Search")
			serviceManager.public = append(serviceManager.public, "Server.LdapManager.Authenticate")
			serviceManager.public = append(serviceManager.public, "Server.LdapManager.GetComputer")
			serviceManager.public = append(serviceManager.public, "Server.LdapManager.GetAllGroups")
			serviceManager.public = append(serviceManager.public, "Server.LdapManager.GetComputerByIp")
			serviceManager.public = append(serviceManager.public, "Server.LdapManager.GetAllUsers")
			serviceManager.public = append(serviceManager.public, "Server.LdapManager.GetGroupById")
			serviceManager.public = append(serviceManager.public, "Server.LdapManager.GetComputerByName")
			serviceManager.public = append(serviceManager.public, "Server.LdapManager.GetUserById")

			// OAuth2
			serviceManager.public = append(serviceManager.public, "Server.OAuth2Manager.GetResource")

			// File manager
			serviceManager.public = append(serviceManager.public, "Server.FileManager.GetFileByPath")
			serviceManager.public = append(serviceManager.public, "Server.FileManager.GetMimeTypeByExtension")
			serviceManager.public = append(serviceManager.public, "Server.FileManager.IsFileExist")
			serviceManager.public = append(serviceManager.public, "Server.FileManager.OpenFile")
			serviceManager.public = append(serviceManager.public, "Server.FileManager.ReadCsvFile")
			serviceManager.public = append(serviceManager.public, "Server.FileManager.ReadTextFile")
			serviceManager.public = append(serviceManager.public, "Server.FileManager.RemoveFile")
			serviceManager.public = append(serviceManager.public, "Server.FileManager.CreateDir")
			serviceManager.public = append(serviceManager.public, "Server.FileManager.CreateFile")
			serviceManager.public = append(serviceManager.public, "Server.FileManager.DeleteFile")

			// Email Manager
			serviceManager.public = append(serviceManager.public, "Server.EmailManager.ReceiveMailFunc")
			serviceManager.public = append(serviceManager.public, "Server.EmailManager.SendEmail")
			serviceManager.public = append(serviceManager.public, "Server.EmailManager.ValidateEmail")

			serviceManager.public = append(serviceManager.public, "Server.ServiceManager.RegisterAction")
			serviceManager.public = append(serviceManager.public, "Server.ServiceManager.GetServiceActions")

			///////////////////////// restricted /////////////////////////

			// Event manager
			serviceManager.restricted = append(serviceManager.restricted, "Server.EventManager.BroadcastEventData")
			serviceManager.restricted = append(serviceManager.restricted, "Server.EventManager.AppendEventFilter")

			// Data manager
			serviceManager.restricted = append(serviceManager.restricted, "Server.DataManager.Close")
			serviceManager.restricted = append(serviceManager.restricted, "Server.DataManager.CreateDataStore")
			serviceManager.restricted = append(serviceManager.restricted, "Server.DataManager.ImportXmlData")
			serviceManager.restricted = append(serviceManager.restricted, "Server.DataManager.DeleteDataStore")
			serviceManager.restricted = append(serviceManager.restricted, "Server.DataManager.ImportXsdSchema")
			serviceManager.restricted = append(serviceManager.restricted, "Server.DataManager.Synchronize")

			// Entity manager
			serviceManager.restricted = append(serviceManager.restricted, "Server.EntityManager.CreateEntityPrototype")

			// Security manager
			serviceManager.restricted = append(serviceManager.restricted, "Server.SecurityManager.AppendAccount")
			serviceManager.restricted = append(serviceManager.restricted, "Server.SecurityManager.AppendPermission")
			serviceManager.restricted = append(serviceManager.restricted, "Server.SecurityManager.AppendAction")
			serviceManager.restricted = append(serviceManager.restricted, "Server.SecurityManager.CreateRole")
			serviceManager.restricted = append(serviceManager.restricted, "Server.SecurityManager.DeleteRole")
			serviceManager.restricted = append(serviceManager.restricted, "Server.SecurityManager.RemoveAccount")
			serviceManager.restricted = append(serviceManager.restricted, "Server.SecurityManager.RemovePermission")
			serviceManager.restricted = append(serviceManager.restricted, "Server.SecurityManager.RemoveAction")

			// Ldap manager
			serviceManager.restricted = append(serviceManager.restricted, "Server.LdapManager.Connect")
			serviceManager.restricted = append(serviceManager.restricted, "Server.LdapManager.SynchronizeComputers")
			serviceManager.restricted = append(serviceManager.restricted, "Server.LdapManager.SynchronizeGroups")
			serviceManager.restricted = append(serviceManager.restricted, "Server.LdapManager.SynchronizeUsers")
		*/
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
	serviceManager.m_servicesLst = make([]Service, 0) // Keep the order of intialisation.

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
				// TODO uncomment when all service will be corrected with this code.
				//if strings.Index(action.M_doc, "@api ") != -1 { // Only api action are exported...
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
			//}
		}
	}

	GetServer().GetEntityManager().getCargoEntities().SaveEntity()

	// Here I will generate the client service class.
	var src string
	var eventTypename = strings.Replace(service.getId(), "Manager", "Event", -1)

	// Here I will generate the javascript code use by client side.
	src = "\nvar " + service.getId() + " = function(){\n"
	src += "	if (server == undefined) {\n"
	src += "		return\n"
	src += "	}\n"
	src += "	EventHub.call(this, " + eventTypename + ")\n\n"
	src += "	return this\n"
	src += "}\n\n"

	src += service.getId() + ".prototype = new EventHub(null);\n"
	src += service.getId() + ".prototype.constructor = " + service.getId() + ";\n\n"

	src += service.getId() + ".prototype.onEvent = function (evt) {\n"
	src += "	EventHub.prototype.onEvent.call(this, evt)\n"
	src += "}\n"

	actions := this.GetServiceActions(service.getId(), "", "")

	for i := 0; i < len(actions); i++ {
		action := actions[i]
		name := action.M_name[strings.LastIndex(action.M_name, ".")+1:]
		doc := action.GetDoc()
		if strings.Index(doc, "@api ") != -1 {
			if strings.Index(doc, "@src\n") != -1 {
				// Here the code of the method is defined in the documentation.
				src += doc[strings.Index(doc, "@src\n")+5:] + "\n"
			} else {
				src += service.getId() + ".prototype." + strings.ToLower(name[0:1]) + name[1:] + " = function("

				// Now the parameters...
				if action.M_parameters != nil {
					// The last tow parameters are sessionId and message Id
					for j := 0; j < len(action.M_parameters)-2; j++ {
						src += action.M_parameters[j].GetName()
						if j < len(action.M_parameters)-2 {
							src += ", "
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
					src += callbacks[j]
					if j < len(callbacks)-1 {
						src += ", "
					}
				}
				if len(callbacks) > 0 {
					src += ", "
				}
				src += "caller){\n"

				// Here I will generate the content of the function.
				if action.M_parameters != nil {
					src += "	var params = []\n"
					for j := 0; j < len(action.M_parameters)-2; j++ {
						param := action.M_parameters[j]
						paramTypeName := param.GetType()
						if paramTypeName == "string" {
							src += "	params.push(createRpcData(" + param.GetName() + ", \"STRING\", \"" + param.GetName() + "\"))\n"
						} else if strings.HasPrefix(paramTypeName, "int") {
							src += "	params.push(createRpcData(" + param.GetName() + ", \"INTEGER\", \"" + param.GetName() + "\"))\n"
						} else if paramTypeName == "bool" {
							src += "	params.push(createRpcData(" + param.GetName() + ", \"BOOLEAN\", \"" + param.GetName() + "\"))\n"
						} else if paramTypeName == "double" || strings.HasPrefix(paramTypeName, "float") {
							src += "	params.push(createRpcData(" + param.GetName() + ", \"DOUBLE\", \"" + param.GetName() + "\"))\n"
						} else if paramTypeName == "[]unit8" || paramTypeName == "[]byte" {
							src += "	params.push(createRpcData(" + param.GetName() + ", \"BYTES\", \"" + param.GetName() + "\"))\n"
						} else {
							// Array or Object or array of object...
							src += "	params.push(createRpcData(" + param.GetName() + ", \"JSON_STR\", \"" + param.GetName() + "\"))\n"
						}
					}
				}

				// Now will generate the code for executeJsFunction.
				src += "\n	server.executeJsFunction(\n"
				src += "	\"" + service.getId() + name + "\",\n"
				src += "	params, \n"
				caller := "{"
				if Utility.Contains(callbacks, "progressCallback") {
					// Set the progress callback.
					src += "	function (index, total, caller) { // Progress callback\n"
					src += "		caller.progressCallback(index, total, caller.caller)\n"
					src += "	},\n"

					// Set the caller.
					caller += "\"progressCallback\":progressCallback, "
				} else {
					src += "	undefined, //progress callback\n"
				}

				if Utility.Contains(callbacks, "successCallback") {
					// Set the progress callback.
					src += "	function (results, caller) { // Success callback\n"
					if len(action.M_results) > 0 {
						typeName := action.M_results[0].M_type
						isArray := strings.HasPrefix(typeName, "[]")
						typeName = strings.Replace(typeName, "[]", "", -1)
						typeName = strings.Replace(typeName, "*", "", -1)
						// Now I will test if the type is an entity...
						if strings.Index(typeName, ".") > -1 {
							// Here I got an entity...

							src += "		server.entityManager.getEntityPrototype(\"" + typeName + "\", \"" + typeName[0:strings.Index(typeName, ".")] + "\",\n"
							src += "			function (prototype, caller) { // Success Callback\n"
							// in case of an array...
							if isArray {
								src += "			var entities = []\n"
								src += "			for (var i = 0; i < caller.results[0].length; i++) {\n"
								src += "				var entity = eval(\"new \" + prototype.TypeName + \"()\")\n"
								src += "				if (i == caller.results[0].length - 1) {\n"
								src += "					entity.initCallback = function (caller) {\n"
								src += "						return function (entity) {\n"
								src += "							server.entityManager.setEntity(entity)\n"
								src += "							caller.successCallback(entities, caller.caller)\n"
								src += "						}\n"
								src += "					} (caller)\n"
								src += "				}else{\n"
								src += "					entity.initCallback = function (entity) {\n"
								src += "						server.entityManager.setEntity(entity)\n"
								src += "					}\n"
								src += "				}\n"
								src += "				entities.push(entity)\n"
								src += "				entity.init(caller.results[0][i])\n"
								src += "			}\n"
							} else {
								// In case of a regular entity.
								src += "			if (caller.results[0] == null) {\n"
								src += "				return\n"
								src += "			}\n"

								// In case of existing entity.
								src += "			if (entities[caller.results[0].UUID] != undefined && caller.results[0].TYPENAME == caller.results[0].__class__) {\n"
								src += "				caller.successCallback(entities[caller.results[0].UUID], caller.caller)\n"
								src += "				return // break it here.\n"
								src += "			}\n\n"

								src += "			var entity = eval(\"new \" + prototype.TypeName + \"()\")\n"
								src += "				entity.initCallback = function () {\n"
								src += "					return function (entity) {\n"
								src += "						caller.successCallback(entity, caller.caller)\n"
								src += "				}\n"
								src += "			}(caller)\n"
								src += "			entity.init(caller.results[0])\n"
							}

							src += "			},\n"
							src += "			function (errMsg, caller) { // Error Callback\n"
							src += "				caller.errorCallback(errMsg, caller.caller)\n"
							src += "			},\n"
							caller := "{ \"caller\": caller.caller"

							if Utility.Contains(callbacks, "progressCallback") {
								caller += ", \"progressCallback\": caller.progressCallback"
							}
							caller += ", \"successCallback\": caller.successCallback, \"errorCallback\": caller.errorCallback, \"results\": results }\n"
							src += "			" + caller

							src += "		)\n"

						} else {
							// Here I got a regulat type.
							src += "		caller.successCallback(results, caller.caller)\n"
						}
					} else {
						src += "		caller.successCallback(results, caller.caller)\n"
					}

					src += "	},\n"
					// Set the caller.
					caller += "\"successCallback\":successCallback, "
				} else {
					src += "	undefined, //success callback\n"
				}

				if Utility.Contains(callbacks, "errorCallback") {
					src += "	function (errMsg, caller) { // Error callback\n"
					src += "		caller.errorCallback(errMsg, caller.caller)\n"
					src += "		server.errorManager.onError(errMsg)\n"
					src += "	},"
					// Set the caller.
					caller += "\"errorCallback\":errorCallback, "
				} else {
					src += "	undefined, //error callback\n"
				}

				caller += "\"caller\": caller}"
				src += caller + ")\n"
				src += "}\n\n"
			}
		}
	}

	if service.getId() == "FileManager" {
		log.Println(src)
	}

}

/////////////////////// API ///////////////////////////

/**
 * That function is run on the server side to register action from javascript
 * files in the script folder.
 */
func (this *ServiceManager) RegisterAction(methodName string, parameters []interface{}, results []interface{}) error {

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

	// I will try to find if the action was register
	//	metodUuid := CargoEntitiesActionExists(methodName)
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

	return nil

}

/**
 * Return the list of action for a given service
 */
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
