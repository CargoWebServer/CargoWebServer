package Server

import (
	"log"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

type ServiceManager struct {

	// info about connection on smtp server...
	m_services            map[string]Service
	m_servicesLst         []Service
	m_serviceContainerCmd *exec.Cmd

	// Here is the list of action with theire initial
	// access type.
	hidden     []string
	public     []string
	restricted []string
}

var serviceManager *ServiceManager

func (this *Server) GetServiceManager() *ServiceManager {
	if serviceManager == nil {
		serviceManager = newServiceManager()
		serviceManager.hidden = make([]string, 0)
		serviceManager.public = make([]string, 0)
		serviceManager.restricted = make([]string, 0)

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
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())

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
		this.registerActions(this.m_servicesLst[i])
	}

	// The first step will be to start the service manager.
	serviceContainerPath := GetServer().GetConfigurationManager().GetBinPath() + "/CargoServiceContainer"
	if runtime.GOOS == "windows" {
		serviceContainerPath += ".exe"
	}

	serviceManager.m_serviceContainerCmd = exec.Command(serviceContainerPath)

	// Set the port number.
	// Uncomment to start the service.
	/*serviceManager.m_serviceContainerCmd.Args = append(serviceManager.m_serviceContainerCmd.Args, strconv.Itoa(GetServer().GetConfigurationManager().GetServicePort()))
	//var err error
	err := serviceManager.m_serviceContainerCmd.Start()
	if err != nil {
		log.Println("---> fail to start the service container!")
	}*/

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

	// Stop the process...
	if serviceManager.m_serviceContainerCmd != nil {
		if serviceManager.m_serviceContainerCmd.Process != nil {
			serviceManager.m_serviceContainerCmd.Process.Kill()
		}
	}
}

/**
 * Register a new service.
 */
func (this *ServiceManager) registerService(service Service) {
	log.Println("--------> register service ", service.getId())
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

	// Now I will print it list of function.
	for i := 0; i < serviceType.NumMethod(); i++ {
		// I will try to find if the action was register
		method := serviceType.Method(i)
		methodName := strings.Replace(serviceType.String(), "*", "", -1) + "." + method.Name
		metodUuid := CargoEntitiesActionExists(methodName)

		if len(metodUuid) == 0 && !(strings.HasPrefix(method.Name, "New") && (strings.HasSuffix(method.Name, "Entity") || strings.HasSuffix(method.Name, "EntityFromObject"))) {

			action := new(CargoEntities.Action)
			action.SetName(methodName)

			// Set the uuid
			GetServer().GetEntityManager().NewCargoEntitiesActionEntity(GetServer().GetEntityManager().getCargoEntities().GetUuid(), "", action)

			// The input
			for j := 0; j < method.Type.NumIn(); j++ {
				in := method.Type.In(j)
				// The first paramters is the object itself.
				if j > 2 {
					parameter := new(CargoEntities.Parameter)
					parameter.UUID = "CargoEntities.Parameter%" + Utility.RandomUUID() // Ok must be random
					parameter.TYPENAME = "CargoEntities.Parameter"
					parameter.SetType(in.String())
					parameter.SetName("p" + strconv.Itoa(len(action.M_parameters)))
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

			// Now I will set the access type of the action before save it.
			if Utility.Contains(this.hidden, action.M_name) {
				action.SetAccessType(CargoEntities.AccessType_Hidden)
			}

			if Utility.Contains(this.public, action.M_name) {
				action.SetAccessType(CargoEntities.AccessType_Public)
			}

			if Utility.Contains(this.restricted, action.M_name) {
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
				}
			}

			guestRoleUuid := CargoEntitiesRoleExists("guestRole")
			if len(guestRoleUuid) > 0 {
				guestRoleEntity, _ := GetServer().GetEntityManager().getEntityByUuid(guestRoleUuid, false)
				if action.GetAccessType() == CargoEntities.AccessType_Public {
					guestRoleEntity.GetObject().(*CargoEntities.Role).SetActions(action)
				}
			}
		}
	}
	GetServer().GetEntityManager().getCargoEntities().SaveEntity()
}
