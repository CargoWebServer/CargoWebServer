/**
 * That file contain the code for OAuth2 service.
 */

package Server

import (
	"log"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
)

type OAuth2Manager struct {
	// Contain the list of avalable ldap servers...
	m_configsInfo map[string]Config.OAuth2Configuration
}

var oauth2Manager *OAuth2Manager

func (this *Server) GetOAuth2Manager() *OAuth2Manager {
	if oauth2Manager == nil {
		oauth2Manager = newOAuth2Manager()
	}
	return oauth2Manager
}

func newOAuth2Manager() *OAuth2Manager {
	oauth2Manager := new(OAuth2Manager)
	oauth2Manager.m_configsInfo = make(map[string]Config.OAuth2Configuration)
	return oauth2Manager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * That function is use to synchronize the information of a ldap server
 * with a given id.
 */
func (this *OAuth2Manager) initialize() {
	// register service avalaible action here.
	log.Println("--> initialyze OAuth2Manager")
	// Create the default configurations
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())
}

func (this *OAuth2Manager) getId() string {
	return "OAuth2Manager"
}

func (this *OAuth2Manager) start() {
	log.Println("--> Start OAuth2Manager")
	configurations := GetServer().GetConfigurationManager().GetOAuth2Configuration()

	for i := 0; i < len(configurations); i++ {
		this.m_configsInfo[configurations[i].M_id] = configurations[i]
	}
}

func (this *OAuth2Manager) stop() {
	log.Println("--> Stop OAuth2Manager")
}
