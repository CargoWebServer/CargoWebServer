package Server

// Information about the service.
type ServiceInfo struct {
	// The name must be the name of the struct that implement
	// the service interface.
	Name string

	// If is set as true the service is start automaticaly.
	Start bool
}

/**
 * A service is a class that can be use to extend the server functionality.
 */
type Service interface {
	/**
	 * Initialisation of the service.
	 */
	Initialize()

	/**
	 * Starting the service
	 */
	Start()

	/**
	 * Stoping the service.
	 */
	Stop()
}
