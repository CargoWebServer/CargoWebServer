package Server

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"code.myceliUs.com/CargoWebServer/Cargo/Persistence/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"

	"github.com/pkg/profile"
	//"github.com/skratchdot/open-golang/open"
)

var (
	server *Server
)

type Server struct {
	// Singleton...
	// The configuration manager.
	configurationManager *ConfigurationManager

	// The security manager.
	securityManager *SecurityManager

	// The event manager
	eventManager *EventManager

	// Session manager
	sessionManager *SessionManager

	// The account manager
	accountManager *AccountManager

	// Database accessor...
	dataManager *DataManager

	// The ldap manager
	ldapManager *LdapManager

	// The mail
	emailManager *EmailManager

	// The map of logger.
	loggers map[string]*Logger

	// The workflow manager
	workflowManager *WorkflowManager

	// The workflow processor
	workflowProcessor *WorkflowProcessor

	// The entity manager
	entityManager *EntityManager

	// File manager
	fileManager *FileManager

	// Project manager
	projectManager *ProjectManager

	// XSD/XML schma...
	schemaManager *SchemaManager

	// The network related stuff...
	hub              *Hub
	messageProcessor *MessageProcessor

	// That map contain list of other server on the network.
	peers map[string]connection

	// Cache Manager
	cacheManager *CacheManager

	// The address information.
	addressInfo *Utility.IPInfo

	// The profiler...
	profiler *profile.Profile
}

/**
 * Create a new server...
 */
func newServer() *Server {
	// The server object itself...
	server = new(Server)

	// The configuration manager...
	server.configurationManager = newConfigurationManager()

	// Database accessor...
	server.dataManager = newDataManager()

	// The mail
	// output mail...
	server.emailManager = newEmailManager()

	// The event manager.
	server.eventManager = newEventManager()

	// The account manager.
	server.accountManager = newAccountManager()

	// Session
	server.sessionManager = newSessionManager()

	// The ldap manager
	server.ldapManager = newLdapManager()

	// The entity manager
	server.entityManager = newEntityManager()

	// File manager
	server.fileManager = newFileManager()

	// Schema manager
	server.schemaManager = newSchemaManager()

	// The project manager
	server.projectManager = newProjectManager()

	// The security manager
	server.securityManager = newSecurityManager()

	// Network...
	server.hub = NewHub()
	server.messageProcessor = newMessageProcessor()

	// The cache manager
	server.cacheManager = newCacheManager()

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
	this.cacheManager.Initialize()
	this.eventManager.Initialize()
	this.configurationManager.Initialyze()
	this.dataManager.Initialyze()
	this.entityManager.Initialize()
	this.sessionManager.Initialize()
	this.accountManager.Initialize()
	this.emailManager.Initialyze()
	this.schemaManager.Initialyze()

	// The BPMN functionality...
	// The workflow manager.
	server.workflowManager = newWorkflowManager()

	// The workflow processor.
	server.workflowProcessor = newWorkflowProcessor()

	this.workflowManager.Initialize()
	this.workflowProcessor.Initialize()

	this.securityManager.Initialize()

	// The map of loggers.
	this.loggers = make(map[string]*Logger)

	// The default error logger...
	logger := NewLogger("defaultErrorLogger")
	this.loggers["defaultErrorLogger"] = logger

	this.fileManager.Initialize()
	//this.ldapManager.Initialize()
	this.projectManager.Initialyze()
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

func (this *Server) startWorkflowProcessor() {
	go this.workflowProcessor.run()
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
	//this.profiler = profile.Start(profile.CPUProfile, profile.ProfilePath("."), profile.NoShutdownHook).(*profile.Profile)

	log.Println("Start the server..." /*, this.addressInfo.IP*/)
	// Start the server...
	server.startMessageProcessor()
	server.startHub()

	server.startWorkflowProcessor()

}

/**
 * Stop the server.
 */
func (this *Server) Stop() {

	// Stop processing...
	server.messageProcessor.abortedByEnvironment <- true
	server.hub.abortedByEnvironment <- true
	server.cacheManager.abortedByEnvironment <- true
	// Close all connection.
	server.dataManager.close()

	log.Println("Bye Bye :-)")

	//this.profiler.Stop()

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
func (this *Server) GetAccountManager() *AccountManager {
	return this.accountManager
}

func (this *Server) GetEntityManager() *EntityManager {
	return this.entityManager
}

func (this *Server) GetLdapManager() *LdapManager {
	return this.ldapManager
}

func (this *Server) GetDataManager() *DataManager {
	return this.dataManager
}

func (this *Server) GetSessionManager() *SessionManager {
	return this.sessionManager
}

func (this *Server) GetEventManager() *EventManager {
	return this.eventManager
}

func (this *Server) GetSmtpManager() *EmailManager {
	return this.emailManager
}

func (this *Server) GetSchemaManager() *SchemaManager {
	return this.schemaManager
}

func (this *Server) GetProcessor() *MessageProcessor {
	return this.messageProcessor
}

func (this *Server) GetWorkflowManager() *WorkflowManager {
	return this.workflowManager
}

func (this *Server) GetWorkflowProcessor() *WorkflowProcessor {
	return this.workflowProcessor
}

func (this *Server) GetHub() *Hub {
	return this.hub
}

func (this *Server) GetFileManager() *FileManager {
	return this.fileManager
}

func (this *Server) GetConfigurationManager() *ConfigurationManager {
	return this.configurationManager
}

func (this *Server) GetSecurityManager() *SecurityManager {
	return this.securityManager
}

func (this *Server) GetCacheManager() *CacheManager {
	return this.cacheManager
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
