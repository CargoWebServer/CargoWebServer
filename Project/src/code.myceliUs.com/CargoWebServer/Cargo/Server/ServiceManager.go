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
					parameter.SetName("p" + strconv.Itoa(j-1))
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

			// apend it to the entities action.
			action.SetEntitiesPtr(GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities))
			GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities).SetActions(action)

			// I will append the action into the admin role that has all permission.
			adminRoleUuid := CargoEntitiesRoleExists("adminRole")
			if len(adminRoleUuid) > 0 {
				adminRoleEntity, _ := GetServer().GetEntityManager().getEntityByUuid(adminRoleUuid, false)
				adminRoleEntity.GetObject().(*CargoEntities.Role).SetActions(action)
			}
		}
	}
	GetServer().GetEntityManager().getCargoEntities().SaveEntity()
}
