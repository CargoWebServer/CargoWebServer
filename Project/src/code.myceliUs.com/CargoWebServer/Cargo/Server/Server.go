package Server

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
	//"github.com/skratchdot/open-golang/open"
	"code.myceliUs.com/CargoWebServer/Cargo/JS"
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

	// The address information.
	addressInfo *Utility.IPInfo

	// The list of open sub-connection by connection id.
	subConnectionIds map[string][]string
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
	ids := []interface{}{"admin"}
	adminAccountEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.Account", ids, false)
	if err == nil {
		adminAccount := adminAccountEntity.GetObject().(*CargoEntities.Account)
		if adminAccount.GetPassword() == "adminadmin" {
			//open.Run("http://127.0.0.1:9393")
		}
	}

	server.subConnectionIds = make(map[string][]string, 0)

	return server
}

/**
 * Do intialysation stuff here.
 */
func (this *Server) initialize() {

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
	this.GetServiceManager().registerService(this.GetOAuth2Manager())
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
 * Remove all session connections.
 */
func (this *Server) removeAllOpenSubConnections(connectionId string) {
	log.Println("---------> remove all open sub-connection...")
	// Now I will remove it connections.
	subConnectionIds := this.subConnectionIds[connectionId]
	if subConnectionIds != nil {
		for i := 0; i < len(subConnectionIds); i++ {
			connection := this.getConnectionById(subConnectionIds[i])
			connection.Close()
		}
	}
	// Remove it from memory...
	this.subConnectionIds[connectionId] = make([]string, 0)
}

func (this *Server) appendSubConnectionId(connectionId string, subConnectionId string) {
	if this.subConnectionIds[connectionId] == nil {
		this.subConnectionIds[connectionId] = make([]string, 0)
	}

	if !Utility.Contains(this.subConnectionIds[connectionId], subConnectionId) {
		this.subConnectionIds[connectionId] = append(this.subConnectionIds[connectionId], subConnectionId)
	}
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
	// Display the error on the server console.
	log.Println(errorObject.GetBody())
}

/**
 * Execute a vb script cmd.
 * * Windows only...
 */
func (this *Server) runVbs(scriptName string, args []string) ([]string, error) {
	path := this.GetConfigurationManager().GetScriptPath() + "/" + scriptName
	args_ := make([]string, 0)
	args_ = append(args_, "/Nologo") // Remove the trademark...
	args_ = append(args_, path)

	args_ = append(args_, args...)
	//log.Println("-----> ", args_)
	out, err := exec.Command("C:/WINDOWS/system32/cscript.exe", args_...).Output()
	results := strings.Split(string(out), "\n")
	results = results[0 : len(results)-1]
	return results, err
}

//////////////////////////////////////////////////////////
// Api
//////////////////////////////////////////////////////////

/**
 * Execute a visual basic script command.
 */
func (this *Server) ExecuteVbScript(scriptName string, args []string, messageId string, sessionId string) []string {

	//log.Println("----------> run ", scriptName, "args: ", args)
	// Run the given script on the server side.
	results, err := this.runVbs(scriptName, args)

	// Get the session id and the message id...
	if err != nil {
		// Create the error object.
		cargoError := NewError(Utility.FileLine(), ACTION_EXECUTE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	return results
}

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

	// the service manager will start previous service depending of there
	// configurations.
	this.GetServiceManager().start()

	// Here I will set the services code...
	apiSrc := ""
	for _, src := range this.GetServiceManager().m_serviceServerSrc {
		apiSrc += src + "\n"
	}

	// Server side binded functions.
	JS.GetJsRuntimeManager().AppendScript(apiSrc)

	// Now I will set JS function needed by the server side.

	// Init connection is call when a Server object need to be connect on the net work.
	JS.GetJsRuntimeManager().AppendFunction("initConnection",
		func(adress string, openCallback string, closeCallback string, messageCallback string, connectionId string, caller interface{}) {
			log.Println(" init connection with : ", adress)

			// Get the new connection id.
			subConnectionId, err := GetServer().Connect(adress)

			if err != nil {
				log.Println("-------> connection fail: ", err)
				return
			}

			// I will append the connection the session.
			GetServer().appendSubConnectionId(connectionId, subConnectionId)

			// I will execute the openCallback function...
			functionParams := make([]interface{}, 1)
			functionParams[0] = subConnectionId
			JS.GetJsRuntimeManager().ExecuteJsFunction("", connectionId, openCallback, functionParams)

		})

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
func (this *Server) Connect(address string) (string, error) {

	var conn connection
	values := strings.Split(address, ":")

	var host string
	var socket string
	var port int
	if len(values) == 3 {
		socket = values[0]
		host = strings.Replace(values[1], "//", "", -1)
		port, _ = strconv.Atoi(values[2])
	} else if len(values) == 2 {
		socket = "tcp"
		host = values[0]
		port, _ = strconv.Atoi(values[1])
	}

	// Create the new connection.
	if socket == "ws" {
		conn = NewWebSocketConnection()
	} else {
		conn = NewTcpSocketConnection()
	}

	// Open the a new connection with the server.
	log.Println("---> try to open ", host, port)
	err := conn.Open(host, port)
	if err != nil {
		log.Println("=----------> 329: ", err)
		return "", err // The connection fail...
	}

	// Test if the connection is open.
	if !conn.IsOpen() {
		return "", errors.New("Fail to open connection with socket " + host + " at port " + strconv.Itoa(port))
	}

	return conn.GetUuid(), nil
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
