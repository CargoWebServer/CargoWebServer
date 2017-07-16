package Server

import (
	b64 "encoding/base64"
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
	"github.com/robertkrimen/otto"
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

	// That map contain Javascript connection object.
	subConnections map[string]otto.Value
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
	server.subConnections = make(map[string]otto.Value, 0)

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
 * Remove a sub-connection.
 */
func (this *Server) removeSubConnection(connectionId string) {
	subConnectionIds := make([]string, 0)
	for i := 0; i < len(this.subConnectionIds[connectionId]); i++ {
		if connectionId == this.subConnectionIds[connectionId][i] {
			connection := this.getConnectionById(subConnectionIds[i])
			connection.Close()
		} else {
			subConnectionIds = append(subConnectionIds, connectionId)
		}
	}

	this.subConnectionIds[connectionId] = subConnectionIds
}

/**
 * Remove a given subconnection from the list.
 */
func (this *Server) removeSubConnections(connectionId string, subConnectionId string) {
	// Now I will remove it connections.
	subConnectionIds := make([]string, 0)

	if subConnectionIds != nil {
		for i := 0; i < len(this.subConnectionIds[connectionId]); i++ {
			if subConnectionId == this.subConnectionIds[connectionId][i] {
				connection := this.getConnectionById(this.subConnectionIds[connectionId][i])
				connection.Close()
			} else {
				subConnectionIds = append(subConnectionIds, subConnectionId)
			}
		}
	}

	this.subConnectionIds[connectionId] = subConnectionIds

	// Remove from the map of js object to.
	delete(this.subConnections, subConnectionId)
}

/**
 * Trigger the onclose function over the js object.
 */
func (this *Server) onClose(subConnectionId string) {
	subConnection := this.subConnections[subConnectionId]
	if subConnection.Object() != nil {
		subConnection.Object().Call("onclose")
	}
}

/**
 * Trigger the onMessage function over the js object.
 */
func (this *Server) onMessage(subConnectionId string) {
	subConnection := this.subConnections[subConnectionId]
	if subConnection.Object() != nil {
		subConnection.Object().Call("onmessage")
	}
}

/**
 * Remove all session connections.
 */
func (this *Server) removeAllOpenSubConnections(connectionId string) {
	// Now I will remove it connections.
	subConnectionIds := this.subConnectionIds[connectionId]
	if subConnectionIds != nil {
		for i := 0; i < len(subConnectionIds); i++ {
			connection := this.getConnectionById(subConnectionIds[i])
			connection.Close()

			// remove from js map to.
			delete(this.subConnections, subConnectionIds[i])
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

	// Convert a utf8 string to a base 64 string
	JS.GetJsRuntimeManager().AppendFunction("utf8_to_b64", func(data string) string {
		sEnc := b64.StdEncoding.EncodeToString([]byte(data))
		return sEnc
	})

	/**
	 * Made other connection side execute JS code.
	 */
	JS.GetJsRuntimeManager().AppendFunction("executeJsFunction", func(functionSrc string, functionParams []otto.Value, progressCallback string, successCallback string, errorCallback string, caller otto.Value, subConnectionId string) {
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]connection, 1)
		to[0] = GetServer().getConnectionById(subConnectionId)

		param := new(MessageData)
		param.Name = "functionSrc"
		param.Value = functionSrc

		// Append the params.
		params = append(params, param)

		// I will create the function parameters.
		for i := 0; i < len(functionParams); i++ {
			param := new(MessageData)
			paramName, _ := functionParams[i].Object().Get("name")
			param.Name = paramName.String()

			paramValue, _ := functionParams[i].Object().Get("dataBytes")
			if paramValue.IsString() {
				val, _ := paramValue.ToString()
				param.Value = val
			} else if paramValue.IsBoolean() {
				val, _ := paramValue.ToBoolean()
				param.Value = val
			} else if paramValue.IsNull() {
				param.Value = nil
			} else if paramValue.IsNumber() {
				val, _ := paramValue.ToFloat()
				param.Value = val
			} else if paramValue.IsUndefined() {
				param.Value = nil
			} else {
				val, _ := paramValue.Export()
				param.Value = val
			}

			params = append(params, param)
		}

		// The success callback.
		successCallback_ := func(successCallback string, subConnectionId string, caller otto.Value) func(*message, interface{}) {
			return func(rspMsg *message, caller_ interface{}) {
				results := make([]interface{}, 0)
				// So here i will get the message value...
				for i := 0; i < len(rspMsg.msg.Rsp.Results); i++ {

					param := rspMsg.msg.Rsp.Results[i]
					if param.GetType() == Data_DOUBLE {
						val, err := strconv.ParseFloat(string(param.GetDataBytes()), 64)
						if err != nil {
							panic(err)
						}
						results = append(results, val)
					} else if param.GetType() == Data_INTEGER {

						val, err := strconv.ParseInt(string(param.GetDataBytes()), 10, 64)
						if err != nil {
							panic(err)
						}
						results = append(results, val)
					} else if param.GetType() == Data_BOOLEAN {
						val, err := strconv.ParseBool(string(param.GetDataBytes()))
						if err != nil {
							panic(err)
						}
						results = append(results, val)
					} else if param.GetType() == Data_STRING {
						val := string(param.GetDataBytes())
						results = append(results, val)
					} else if param.GetType() == Data_BYTES {
						results = append(results, param.GetDataBytes())
					} else if param.GetType() == Data_JSON_STR {
						val := string(param.GetDataBytes())
						val_, err := b64.StdEncoding.DecodeString(val)
						if err == nil {
							val = string(val_)
						}

						// Only registered type will be process sucessfully here.
						// how the server will be able to know what to do otherwise.
						if strings.HasPrefix(val, "[") && (strings.HasSuffix(val, "]") || strings.HasSuffix(val, "]\n")) {
							// It contain an array of values to be init
							var values interface{}
							if param.GetTypeName() == "[]string" {
								values = make([]string, 0)
							} else {
								values = make([]interface{}, 0)
							}

							err = json.Unmarshal([]byte(val), &values)

							if err == nil {
								p, err := Utility.InitializeStructures(values.([]interface{}), param.GetTypeName())
								if err == nil {

									results = append(results, p.Interface())
								} else {
									//log.Println("Error:", err)
									// Here I will try to create a the array of object.
									if err.Error() == "NotDynamicObject" {
										p, err := Utility.InitializeArray(values.([]interface{}), param.GetTypeName())
										if err == nil {
											if p.IsValid() {
												results = append(results, p.Interface())
											} else {
												// here i will set an empty generic array.
												results = append(results, make([]interface{}, 0))
											}
										}
									}
								}
							}

						} else {
							// It contain an object.
							var valMap map[string]interface{}
							err = json.Unmarshal([]byte(val), &valMap)
							if err == nil {
								p, err := Utility.InitializeStructure(valMap)
								if err != nil {
									log.Println("Error:", err)
									results = append(results, valMap)
								} else {
									results = append(results, p.Interface())
								}
							} else {
								// I will set a nil value to the parameter in that case.
								results = append(results, nil)
							}
						}
					}
				}

				params := make([]interface{}, 2)
				params[0] = results
				params[1] = caller
				// run the success callback.
				for k, v := range GetServer().subConnectionIds {
					if Utility.Contains(v, subConnectionId) {
						if JS.GetJsRuntimeManager().GetVm(k) != nil {
							JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), k, successCallback, params)
						}
					}
				}

			}
		}(successCallback, subConnectionId, caller)

		// The error callback.
		errorCallback_ := func(errorCallback string, subConnectionId string, caller otto.Value) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				errStr := errMsg.msg.Err.Message
				params := make([]interface{}, 2)
				params[0] = *errStr
				params[1] = caller

				// run the error callback.
				for k, v := range GetServer().subConnectionIds {
					if Utility.Contains(v, subConnectionId) {
						if JS.GetJsRuntimeManager().GetVm(k) != nil {
							JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), k, errorCallback, params)
						}
					}
				}
			}
		}(errorCallback, subConnectionId, caller)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback_, nil, errorCallback_)

		GetServer().GetProcessor().m_pendingRequestChannel <- rqst
	})

	/**
	 * Set a ping message to the other end connection...
	 */
	JS.GetJsRuntimeManager().AppendFunction("ping", func(successCallback string, errorCallback string, caller otto.Value, subConnectionId string) {

		id := Utility.RandomUUID()
		method := "Ping"
		params := make([]*MessageData, 0)

		to := make([]connection, 1)
		to[0] = GetServer().getConnectionById(subConnectionId)

		// The success callback.
		successCallback_ := func(successCallback string, subConnectionId string, caller otto.Value) func(*message, interface{}) {
			return func(rspMsg *message, caller_ interface{}) {

				// So here i will get the message value...
				result := string(rspMsg.msg.Rsp.Results[0].DataBytes)
				params := make([]interface{}, 2)
				params[0] = result
				params[1] = caller
				// run the success callback.
				for k, v := range GetServer().subConnectionIds {
					if Utility.Contains(v, subConnectionId) {
						if JS.GetJsRuntimeManager().GetVm(k) != nil {
							JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), k, successCallback, params)
						}
					}
				}
			}
		}(successCallback, subConnectionId, caller)

		// The error callback.
		errorCallback_ := func(errorCallback string, subConnectionId string, caller otto.Value) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				errStr := errMsg.msg.Err.Message
				params := make([]interface{}, 2)
				params[0] = *errStr
				params[1] = caller

				// run the error callback.
				for k, v := range GetServer().subConnectionIds {
					if Utility.Contains(v, subConnectionId) {
						if JS.GetJsRuntimeManager().GetVm(k) != nil {
							JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), k, errorCallback, params)
						}
					}
				}
			}
		}(errorCallback, subConnectionId, caller)

		ping, _ := NewRequestMessage(id, method, params, to, successCallback_, nil, errorCallback_)

		GetServer().GetProcessor().m_pendingRequestChannel <- ping

	})

	/**
	 * Init connection is call when a Server object need to be connect on the net work.
	 */
	JS.GetJsRuntimeManager().AppendFunction("initConnection",
		func(adress string, openCallback string, closeCallback string, connectionId string, service otto.Value, caller otto.Value) otto.Value {
			log.Println(" init connection with : ", adress, " session id: ", connectionId)

			// Get the new connection id.
			subConnection, err := GetServer().connect(adress)

			// The new created connection Js object.
			var conn otto.Value

			if err != nil {
				log.Println("-------> connection fail: ", err)
				return conn
			}

			subConnectionId := subConnection.GetUuid()

			vm := JS.GetJsRuntimeManager().GetVm(connectionId)

			// I will append the connection the session.
			GetServer().appendSubConnectionId(connectionId, subConnectionId)

			// Here I will create the connection object...
			conn, err = vm.Run("new Connection()")

			if err == nil {
				// I will set the connection id.
				conn.Object().Set("id", subConnectionId)

				// Set the connection in the caller.
				service.Object().Set("conn", conn)

				// I will set the open callback.
				_, err := vm.Run("Connection.prototype.onopen = " + openCallback)
				if err != nil {
					log.Println("-----> error!", err)
				}

				// Now the close callback.
				_, err = vm.Run("Connection.prototype.onclose = " + closeCallback)
				if err != nil {
					log.Println("-----> error!", err)
				}

			} else {
				log.Println(err)
			}

			// Keep the connection link...
			GetServer().subConnections[subConnectionId] = conn

			// I will get the client code and inject it in the vm.
			id := Utility.RandomUUID()
			method := "GetServicesClientCode"
			params := make([]*MessageData, 0)
			to := make([]connection, 1)
			to[0] = subConnection
			successCallback := func(connectionId string, conn otto.Value, service otto.Value, caller_ otto.Value) func(rspMsg *message, caller interface{}) {
				return func(rspMsg *message, caller interface{}) {
					src := string(rspMsg.msg.Rsp.Results[0].DataBytes)
					JS.GetJsRuntimeManager().AppendScript(src)
					// Call on open...
					conn.Object().Call("onopen", service, caller_)
				}
			}(connectionId, conn, service, caller)

			errorCallback := func(rspMsg *message, caller interface{}) {
				log.Println("GetServicesClientCode error!!!")
			}

			rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback)
			GetServer().GetProcessor().m_pendingRequestChannel <- rqst

			return conn

		})

	// Now I will create the empty session...
	JS.GetJsRuntimeManager().CreateVm("")    // The anonymous session.
	JS.GetJsRuntimeManager().InitScripts("") // Run the script for the default session.

	// Test compile analyse...
	JS.GetJsRuntimeManager().GetVm("").Set("server", this)
	JS.GetJsRuntimeManager().GetVm("").Run("TestMessageContainer(30)")
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
func (this *Server) connect(address string) (connection, error) {

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
		log.Println("--------------> connection fail! ", host, port, err)
		return nil, err // The connection fail...
	}

	// Test if the connection is open.
	if !conn.IsOpen() {
		return nil, errors.New("Fail to open connection with socket " + host + " at port " + strconv.Itoa(port))
	}

	return conn, nil
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
