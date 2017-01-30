package Server

import (
	"log"
	"os/exec"
	"runtime"

	"code.myceliUs.com/CargoWebServer/Cargo/Config/CargoConfig"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
)

type ServiceManager struct {

	// info about connection on smtp server...
	m_services            map[string]Service
	m_servicesConfigs     map[string]CargoConfig.ServiceConfiguration
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
	serviceManager.m_servicesConfigs = make(map[string]CargoConfig.ServiceConfiguration)

	// The list of optional services.
	serviceConfigurations := GetServer().GetConfigurationManager().GetLocalServiceConfigurations()

	// Also ge the other services...
	serviceConfigurations = append(serviceConfigurations, GetServer().GetConfigurationManager().GetServiceConfigurations()...)

	// Now I will inialyse the services.
	for i := 0; i < len(serviceConfigurations); i++ {
		serviceConfiguration := serviceConfigurations[i]
		// First of all I will retreive the service object...
		params := make([]interface{}, 0)
		service, err := Utility.CallMethod(GetServer(), "Get"+serviceConfiguration.M_id, params)

		// In case of local service...
		if serviceConfiguration.M_ipv4 == "127.0.0.1" {
			if err != nil {
				log.Println("--> service whit name ", serviceConfiguration.M_id, " dosen't exist!")
			} else {
				serviceManager.m_servicesConfigs[serviceConfiguration.M_id] = serviceConfiguration
				serviceManager.m_services[serviceConfiguration.M_id] = service.(Service)
			}
		} else {
			// I case of distant services...
			log.Println("--> try to connect to service with name ", serviceConfiguration.M_id, " at adresse ", serviceConfiguration.GetIpv4())
			// So here I will create the connection...
		}
	}

	return serviceManager
}

/**
 * Do intialysation stuff here.
 */
func (this *ServiceManager) Initialize() {
	// Here I will start the c++ service container...
	log.Println("--> Initialize ServiceManager")
	for _, service := range this.m_services {
		service.Initialize()
	}
}

func (this *ServiceManager) GetId() string {
	return "ServiceManager"
}

func (this *ServiceManager) Start() {

	log.Println("--> Start ServiceManager")
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

	for _, service := range this.m_services {
		// Get the service configuration information.
		serviceConfiguration := this.m_servicesConfigs[service.GetId()]
		if serviceConfiguration.M_start == true {
			service.Start()
		}
	}
}

func (this *ServiceManager) Stop() {
	log.Println("--> Stop ServiceManager")
	for _, service := range this.m_services {
		service.Stop()
	}

	// Stop the process...
	serviceManager.m_serviceContainerCmd.Process.Kill()
}
