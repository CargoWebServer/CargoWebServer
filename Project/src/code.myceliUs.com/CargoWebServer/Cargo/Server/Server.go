package Server

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
	//"github.com/skratchdot/open-golang/open"
)

var (
	server *Server
)

type Server struct {

	// The map of logger.
	loggers map[string]*Logger

	// The network related stuff...
	hub              *Hub
	messageProcessor *MessageProcessor

	// That map contain list of other server on the network.
	// Use by services manager as example.
	peers map[string]connection

	// The address information.
	addressInfo *Utility.IPInfo
}

/**
 * Create a new server...
 */
func newServer() *Server {
	// The server object itself...
	server = new(Server)

	// Network...
	server.hub = NewHub()
	server.messageProcessor = newMessageProcessor()

	// Initialyse with the default configuration.
	server.initialize()

	// Get the server address information.
	server.addressInfo, _ = Utility.MyIP()

	// if Cargoroot is not set...
	if len(os.Getenv("CARGOROOT")) == 0 {
		// In that case I will install the server...
	}

	// if the admin has password adminadmin I will display the setup wizard..
	adminAccountEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.Account", "admin")
	if err == nil {
		adminAccount := adminAccountEntity.GetObject().(*CargoEntities.Account)
		if adminAccount.GetPassword() == "adminadmin" {
			//open.Run("http://127.0.0.1:9393/Bridge")
		}
	}

	log.Println("adminAccount", adminAccountEntity)
	return server
}

/**
 * Do intialysation stuff here.
 */
func (this *Server) initialize() {

	// Contain the reference to other services on the network.
	this.peers = make(map[string]connection)

	// Must be call first.
	this.GetConfigurationManager()
	this.GetServiceManager()

	// Start the cache manager.
	this.GetCacheManager().initialize()
	this.GetCacheManager().start()

	// The basic services.
	this.GetServiceManager().registerService(this.GetEventManager())
	this.GetServiceManager().registerService(this.GetConfigurationManager())

	// Those service are not manage by the service manager.
	this.GetServiceManager().registerService(this.GetDataManager())
	this.GetServiceManager().registerService(this.GetEntityManager())
	this.GetServiceManager().registerService(this.GetSessionManager())
	this.GetServiceManager().registerService(this.GetAccountManager())
	this.GetServiceManager().registerService(this.GetSecurityManager())

	// call other service in order to register theire configuration.
	this.GetServiceManager().registerService(this.GetLdapManager())
	this.GetServiceManager().registerService(this.GetFileManager())
	this.GetServiceManager().registerService(this.GetEmailManager())
	this.GetServiceManager().registerService(this.GetProjectManager())
	this.GetServiceManager().registerService(this.GetSchemaManager())

	// BPMN stuff
	//this.GetServiceManager().registerService(this.GetWorkflowManager())
	//this.GetServiceManager().registerService(this.GetWorkflowProcessor())

	// The other services are initialyse by the service manager.
	this.GetServiceManager().initialize()

	// The map of loggers.
	this.loggers = make(map[string]*Logger)
	// The default error logger...
	logger := NewLogger("defaultErrorLogger")
	this.loggers["defaultErrorLogger"] = logger
}

/**
 * Return a connection with a given id.
 */
func (this *Server) getConnectionById(id string) connection {
	// Get the conncetion with a given id if it exist...
	if connection, ok := this.hub.connections[id]; ok {
		//Return the connection...
		return connection
	}
	// The connection dosen't exist anymore...
	return nil
}

/**
 * Start processing messages
 */
func (this *Server) startMessageProcessor() {
	go this.messageProcessor.run()
}

/**
 * Start routing message
 */
func (this *Server) startHub() {
	go this.hub.run()
}

/**
 * If something goes wrong that funtion report the error to the client side.
 */
func (this *Server) reportErrorMessage(messageId string, sessionId string, errorObject *CargoEntities.Error) {
	conn := this.getConnectionById(sessionId)
	if conn != nil {
		to := make([]connection, 1)
		to[0] = conn
		errorObjectStr, _ := json.Marshal(errorObject)
		errMsg := NewErrorMessage(messageId, int32(errorObject.GetCode()), errorObject.GetBody(), errorObjectStr, to)
		conn.Send(errMsg.GetBytes())
	}
}

//////////////////////////////////////////////////////////
// Api
//////////////////////////////////////////////////////////

/**
 * Get the server singleton.
 */
func GetServer() *Server {
	if server == nil {
		server = newServer()
	}
	return server
}

/**
 * Start the server.
 */
func (this *Server) Start() {

	log.Println("Start the server...")
	// Start the server...
	server.startMessageProcessor()
	server.startHub()

	/*this.GetEventManager().start()

	// Starting basic services...
	this.GetConfigurationManager().start()
	this.GetDataManager().start()
	this.GetEntityManager().start()
	this.GetSessionManager().start()
	this.GetAccountManager().start()
	this.GetSecurityManager().start()*/

	// the service manager will start previous service depending of there
	// configurations.
	this.GetServiceManager().start()
}

/**
 * Stop the server.
 */
func (this *Server) Stop() {

	// Stop processing...
	server.messageProcessor.abortedByEnvironment <- true
	server.hub.abortedByEnvironment <- true

	// must be call last
	this.GetServiceManager().stop()

	log.Println("Bye Bye :-)")

	// Now stop the process.
	os.Exit(0)
}

/**
 * Restart the server.
 */
func (this *Server) Restart() {
	// Stop processing...
	server.Stop()

	// I will reinit the server...
	server = newServer()

	// Start the server...
	server.Start()

}

/**
 * Set the applications root path.
 */
func (this *Server) SetRootPath(path string) error {
	err := os.Setenv("CARGOROOT", path)
	if err != nil {
		log.Println("Set Root path error:", err)
	}
	return err
}

//////////////////////////////////////////////////////////
// Interface to other servers...
//////////////////////////////////////////////////////////
/**
 * Open a new connection with server on the network...
 */
func (this *Server) Connect(host string, port int) error {

	// Create the new connection.
	conn := new(tcpSocketConnection)

	// Open the a new connection with the server.
	err := conn.Open(host, port)
	if err != nil {
		return err // The connection fail...
	}

	// Here I will create the new connection...
	this.hub.register <- conn

	// Start the writing loop...
	go conn.Writer()

	// Start the reading loop...
	go conn.Reader()

	// Keep the reference... host:port will be the id.
	this.peers[host+":"+strconv.Itoa(port)] = conn

	return nil
}

/**
 * Close the connection with the other server.
 */
func (this *Server) Disconnect(host string, port int) {
	conn := this.peers[host+":"+strconv.Itoa(port)]
	if conn != nil {
		conn.Close()
		delete(this.peers, host+":"+strconv.Itoa(port))
	}
}

//////////////////////////////////////////////////////////
// Getter
//////////////////////////////////////////////////////////

func (this *Server) GetProcessor() *MessageProcessor {
	return this.messageProcessor
}

func (this *Server) GetHub() *Hub {
	return this.hub
}

func (this *Server) AppendLogger(logger *Logger) {
	this.loggers[logger.id] = logger
}

func (this *Server) GetLoggerById(id string) *Logger {
	return this.loggers[id]
}

func (this *Server) GetDefaultErrorLogger() *Logger {
	return this.loggers["defaultErrorLogger"]
}
