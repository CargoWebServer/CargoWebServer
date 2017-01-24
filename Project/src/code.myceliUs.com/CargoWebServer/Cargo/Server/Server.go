package Server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"code.myceliUs.com/CargoWebServer/Cargo/Persistence/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
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
	peers map[string]connection

	// The address information.
	addressInfo *Utility.IPInfo

	// The list of services...
	services []ServiceInfo
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

	server.services = make([]ServiceInfo, 0)

	// Initialyse with the default configuration.
	server.initialize()

	// Get the server address information.
	server.addressInfo, _ = Utility.MyIP()

	// if Cargoroot is not set...
	if len(os.Getenv("CARGOROOT")) == 0 {
		// In that case I will install the server...
	}

	// if the admin has password adminadmin I will display the setup wizard..
	adminAccountEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "admin")
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

	// Contain the reference to other server on the network.
	this.peers = make(map[string]connection)

	// Initialyse managers...
	this.GetCacheManager().Initialize()
	// Must be start before other service initialization.
	this.GetCacheManager().Start()

	this.GetEventManager().Initialize()

	// Basic services...
	this.GetConfigurationManager().Initialize()
	this.GetDataManager().Initialize()

	this.GetEntityManager().Initialize()
	this.GetSessionManager().Initialize()
	this.GetAccountManager().Initialize()
	this.GetSecurityManager().Initialize()

	// The map of loggers.
	this.loggers = make(map[string]*Logger)
	// The default error logger...
	logger := NewLogger("defaultErrorLogger")
	this.loggers["defaultErrorLogger"] = logger

	// TODO load other module here...
	root := this.GetConfigurationManager().GetApplicationDirectoryPath()
	servicesInfo_, err := ioutil.ReadFile(root + "/services.json")
	if err != nil {
		log.Println("--> no services configuration was found!")
	} else {
		json.Unmarshal(servicesInfo_, &this.services)
	}

	// Optional services Initialysation...
	for i := 0; i < len(this.services); i++ {
		// First of all I will retreive the service object...
		params := make([]interface{}, 0)
		service, err := Utility.CallMethod(this, "Get"+this.services[i].Name, params)
		if err != nil {
			log.Println("--> service whit name ", this.services[i].Name, " dosen't exist!")
		} else {
			service.(Service).Initialize()
		}
	}

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

	this.GetEventManager().Start()

	// Starting basic services...
	this.GetConfigurationManager().Start()
	this.GetDataManager().Start()
	this.GetEntityManager().Start()
	this.GetSessionManager().Start()
	this.GetAccountManager().Start()
	this.GetSecurityManager().Start()

	// Optional services Initialysation...
	for i := 0; i < len(this.services); i++ {
		// First of all I will retreive the service object...
		params := make([]interface{}, 0)
		service, err := Utility.CallMethod(this, "Get"+this.services[i].Name, params)
		if err != nil {
			log.Println("--> service whit name ", this.services[i].Name, " dosen't exist!")
		} else {
			if this.services[i].Start {
				service.(Service).Start()
			}
		}
	}

}

/**
 * Stop the server.
 */
func (this *Server) Stop() {

	// Stop processing...
	server.messageProcessor.abortedByEnvironment <- true
	server.hub.abortedByEnvironment <- true

	// Stoping service...
	this.GetConfigurationManager().Stop()
	this.GetDataManager().Stop()
	this.GetEntityManager().Stop()
	this.GetSessionManager().Stop()
	this.GetAccountManager().Stop()
	this.GetSecurityManager().Stop()
	this.GetCacheManager().Stop()

	// Optional services Initialysation...
	for i := 0; i < len(this.services); i++ {
		// First of all I will retreive the service object...
		params := make([]interface{}, 0)
		service, err := Utility.CallMethod(this, "Get"+this.services[i].Name, params)
		if err != nil {
			log.Println("--> service whit name ", this.services[i].Name, " dosen't exist!")
		} else {
			service.(Service).Stop()
		}
	}

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
