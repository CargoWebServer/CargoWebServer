package Server

/**
 * A service is a class that can be use to extend the server functionality.
 */
type Service interface {

	/**
	* Return the id of a given service.
	 */
	GetId() string

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
