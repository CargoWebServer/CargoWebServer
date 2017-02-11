package Server

import (
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
)

/**
 * A service is a class that can be use to extend the server functionality.
 */
type Service interface {

	/**
	* Return the id of a given service.
	 */
	getId() string

	/**
	 * Initialisation of the service.
	 */
	initialize()

	/**
	 * Starting the service
	 */
	start()

	/**
	 * Stoping the service.
	 */
	stop()

	/**
	 * Get the service configurations
	 */
	getConfig() *Config.ServiceConfiguration
}
