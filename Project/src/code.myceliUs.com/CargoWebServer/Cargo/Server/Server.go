package Server

// Minimal build options...
// -i -tags "Config CargoEntities"
// With the workflowmanager
// -i -tags "Config CargoEntities BPMN20 BPMNDI BPMS DI DC"

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"os"

	"strconv"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
	//"github.com/skratchdot/open-golang/open"
	"os/exec"
	"runtime"

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

	// Contain the list of active command.
	cmds []*exec.Cmd

	// Contain the list of active command and their calling session.
	sessionCmds map[string][]*exec.Cmd
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

	// Active commands by session.
	server.sessionCmds = make(map[string][]*exec.Cmd, 0)

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
 * Retunr a connection with a given addresse.
 */
func (this *Server) getConnectionByIp(ipv4 string, port int) connection {
	// Get the conncetion with a given id if it exist...
	for _, connection := range this.hub.connections {
		if connection.GetAddrStr() == ipv4 && connection.GetPort() == port {
			return connection
		}
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

	// Remove ressource use by scripts.
	callback := func(sessionId string) func() {
		return func() {
			log.Println("--> session " + sessionId + " is now close!")
			// Now I will kill it commands.
			cmds := GetServer().sessionCmds[sessionId]
			for i := 0; i < len(cmds); i++ {
				// Remove the command
				GetServer().removeCmd(cmds[i])
			}
		}
	}(subConnectionId)

	// Remove the JS session
	JS.GetJsRuntimeManager().CloseSession(subConnectionId, callback)

	// I will also kill running cmd started by this (connection/session).
	cmds := this.sessionCmds[subConnectionId]

	// Remove command.
	for i := 0; i < len(cmds); i++ {
		// Remove the command
		this.removeCmd(cmds[i])
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
 *
 */

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
	for id, src := range this.GetServiceManager().m_serviceServerSrc {
		// Server side binded functions.
		JS.GetJsRuntimeManager().AppendScript("CargoWebServer/"+id, src)
	}

	////////////////////////////////////////////////////////////////////////////
	// SOM function.
	////////////////////////////////////////////////////////////////////////////
	JS.GetJsRuntimeManager().AppendFunction("CargoWebServer.GetServer", func() *Server {
		return GetServer()
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
		param.TYPENAME = "Server.MessageData"
		// Append the params.
		params = append(params, param)

		// I will create the function parameters.
		for i := 0; i < len(functionParams); i++ {
			param := new(MessageData)
			paramName, _ := functionParams[i].Object().Get("name")
			param.Name = paramName.String()
			param.TYPENAME = "Server.MessageData"
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
		successCallback_ := func(successCallback string, subConnectionId string) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
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
									// I will set unmarshal values in that case.
									results = values.([]interface{})
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

				if rspMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), "", successCallback, params)
				} else {
					// run the success callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							// Todo test if the subconnection id is equal to (to)
							JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), k, successCallback, params)
						}
					}
				}

			}
		}(successCallback, subConnectionId)

		// The error callback.
		errorCallback_ := func(errorCallback string, subConnectionId string) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				errStr := errMsg.msg.Err.Message
				params := make([]interface{}, 2)
				params[0] = *errStr
				params[1] = caller
				// run the error callback.
				if errMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), "", errorCallback, params)
				} else {
					// run the error callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), k, errorCallback, params)
						}
					}
				}
			}
		}(errorCallback, subConnectionId)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback_, nil, errorCallback_, caller)

		go func(rqst *message) {
			GetServer().GetProcessor().m_sendRequest <- rqst
		}(rqst)

	})

	/**
	 * Set a ping message to the other end connection...
	 */
	JS.GetJsRuntimeManager().AppendFunction("CargoWebServer.ping", func(successCallback string, errorCallback string, caller otto.Value, subConnectionId string) {

		id := Utility.RandomUUID()
		method := "Ping"
		params := make([]*MessageData, 0)

		to := make([]connection, 1)
		to[0] = GetServer().getConnectionById(subConnectionId)

		// The success callback.
		successCallback_ := func(successCallback string, subConnectionId string) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {

				// So here i will get the message value...
				result := string(rspMsg.msg.Rsp.Results[0].DataBytes)
				params := make([]interface{}, 2)
				params[0] = result
				params[1] = caller

				// run the success callback.
				if rspMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), "", successCallback, params)
				} else {
					// run the success callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							// Todo test if the subconnection id is equal to (to)
							JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), k, successCallback, params)
						}
					}
				}
			}
		}(successCallback, subConnectionId)

		// The error callback.
		errorCallback_ := func(errorCallback string, subConnectionId string) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				errStr := errMsg.msg.Err.Message
				params := make([]interface{}, 2)
				params[0] = *errStr
				params[1] = caller

				// run the error callback.
				if errMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), "", errorCallback, params)
				} else {
					// run the error callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), k, errorCallback, params)
						}
					}
				}
			}
		}(errorCallback, subConnectionId)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback_, nil, errorCallback_, caller)

		go func(rqst *message) {
			GetServer().GetProcessor().m_sendRequest <- rqst
		}(rqst)

	})

	/**
	 * Set a executeVbSrcript message to the other end connection...
	 */
	JS.GetJsRuntimeManager().AppendFunction("CargoWebServer.executeVbSrcript", func(scriptName string, args []string, successCallback string, errorCallback string, caller otto.Value, subConnectionId string) {
		id := Utility.RandomUUID()
		method := "ExecuteVbScript"
		params := make([]*MessageData, 0)

		param0 := new(MessageData)
		param0.Name = "scriptName"
		param0.Value = scriptName

		param1 := new(MessageData)
		param1.Name = "args"
		param1.TYPENAME = "[]string"
		param1.Value = args

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)

		to := make([]connection, 1)
		to[0] = GetServer().getConnectionById(subConnectionId)

		// The success callback.
		successCallback_ := func(successCallback string, subConnectionId string) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {

				// So here i will get the message value...
				result := string(rspMsg.msg.Rsp.Results[0].DataBytes)
				params := make([]interface{}, 2)
				params[0] = result
				params[1] = caller
				// run the success callback.
				if rspMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), "", successCallback, params)
				} else {
					// run the success callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							// Todo test if the subconnection id is equal to (to)
							JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), k, successCallback, params)
						}
					}
				}
			}
		}(successCallback, subConnectionId)

		// The error callback.
		errorCallback_ := func(errorCallback string, subConnectionId string) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				errStr := errMsg.msg.Err.Message
				params := make([]interface{}, 2)
				params[0] = *errStr
				params[1] = caller

				// run the error callback.
				if errMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), "", errorCallback, params)
				} else {
					// run the error callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), k, errorCallback, params)
						}
					}
				}
			}
		}(errorCallback, subConnectionId)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback_, nil, errorCallback_, caller)

		go func(rqst *message) {
			GetServer().GetProcessor().m_sendRequest <- rqst
		}(rqst)

	})

	/**
	 * Execute external command on the server.
	 */
	JS.GetJsRuntimeManager().AppendFunction("CargoWebServer.runCmd", func(scriptName string, args []string, successCallback string, errorCallback string, caller otto.Value, subConnectionId string) {
		id := Utility.RandomUUID()
		method := "RunCmd"
		params := make([]*MessageData, 0)

		param0 := new(MessageData)
		param0.Name = "scriptName"
		param0.Value = scriptName

		param1 := new(MessageData)
		param1.Name = "args"
		param1.Value = args

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)

		to := make([]connection, 1)
		to[0] = GetServer().getConnectionById(subConnectionId)

		// The success callback.
		successCallback_ := func(successCallback string, subConnectionId string) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {

				// So here i will get the message value...
				result := string(rspMsg.msg.Rsp.Results[0].DataBytes)
				params := make([]interface{}, 2)
				params[0] = result
				params[1] = caller

				// run the success callback.
				if rspMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), "", successCallback, params)
				} else {
					// run the success callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							// Todo test if the subconnection id is equal to (to)
							JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), k, successCallback, params)
						}
					}
				}
			}
		}(successCallback, subConnectionId)

		// The error callback.
		errorCallback_ := func(errorCallback string, subConnectionId string) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				errStr := errMsg.msg.Err.Message
				params := make([]interface{}, 2)
				params[0] = *errStr
				params[1] = caller

				// run the error callback.
				if errMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), "", errorCallback, params)
				} else {
					// run the error callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), k, errorCallback, params)
						}
					}
				}
			}
		}(errorCallback, subConnectionId)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback_, nil, errorCallback_, caller)

		go func(rqst *message) {
			GetServer().GetProcessor().m_sendRequest <- rqst
		}(rqst)

	})

	/**
	 * Get the list of services and their respective source code. The code
	 * permit to get access to service remote actions.
	 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
	 * @param {function} errorCallback In case of error.
	 * @param {object} caller A place to store object from the request context and get it back from the response context.
	 */
	JS.GetJsRuntimeManager().AppendFunction("CargoWebServer.getServicesClientCode", func(successCallback string, errorCallback string, caller otto.Value, subConnectionId string) {
		id := Utility.RandomUUID()
		method := "GetServicesClientCode"
		params := make([]*MessageData, 0)

		to := make([]connection, 1)
		to[0] = GetServer().getConnectionById(subConnectionId)

		// The success callback.
		successCallback_ := func(successCallback string, subConnectionId string) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {

				// So here i will get the message value...
				result := string(rspMsg.msg.Rsp.Results[0].DataBytes)
				params := make([]interface{}, 2)
				params[0] = result
				params[1] = caller

				// run the success callback.
				if rspMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), "", successCallback, params)
				} else {
					// run the success callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							// Todo test if the subconnection id is equal to (to)
							JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), k, successCallback, params)
						}
					}
				}
			}
		}(successCallback, subConnectionId)

		// The error callback.
		errorCallback_ := func(errorCallback string, subConnectionId string) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				errStr := errMsg.msg.Err.Message
				params := make([]interface{}, 2)
				params[0] = *errStr
				params[1] = caller

				// run the error callback.
				if errMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), "", errorCallback, params)
				} else {
					// run the error callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), k, errorCallback, params)
						}
					}
				}
			}
		}(errorCallback, subConnectionId)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback_, nil, errorCallback_, caller)

		go func(rqst *message) {
			GetServer().GetProcessor().m_sendRequest <- rqst
		}(rqst)

	})

	JS.GetJsRuntimeManager().AppendFunction("CargoWebServer.stop", func(successCallback string, errorCallback string, caller otto.Value, subConnectionId string) {
		id := Utility.RandomUUID()
		method := "Stop"
		params := make([]*MessageData, 0)

		to := make([]connection, 1)
		to[0] = GetServer().getConnectionById(subConnectionId)

		// The success callback.
		successCallback_ := func(successCallback string, subConnectionId string) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {

				// So here i will get the message value...
				result := string(rspMsg.msg.Rsp.Results[0].DataBytes)
				params := make([]interface{}, 2)
				params[0] = result
				params[1] = caller

				// run the success callback.
				if rspMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), "", successCallback, params)
				} else {
					// run the success callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							// Todo test if the subconnection id is equal to (to)
							JS.GetJsRuntimeManager().ExecuteJsFunction(rspMsg.GetId(), k, successCallback, params)
						}
					}
				}
			}
		}(successCallback, subConnectionId)

		// The error callback.
		errorCallback_ := func(errorCallback string, subConnectionId string) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				errStr := errMsg.msg.Err.Message
				params := make([]interface{}, 2)
				params[0] = *errStr
				params[1] = caller

				// run the error callback.
				if errMsg.from == nil {
					// Here it's a request from a local JS script.
					JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), "", errorCallback, params)
				} else {
					// run the error callback.
					for k, v := range GetServer().subConnectionIds {
						if Utility.Contains(v, subConnectionId) {
							JS.GetJsRuntimeManager().ExecuteJsFunction(errMsg.GetId(), k, errorCallback, params)
						}
					}
				}
			}
		}(errorCallback, subConnectionId)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback_, nil, errorCallback_, caller)

		go func(rqst *message) {
			GetServer().GetProcessor().m_sendRequest <- rqst
		}(rqst)
	})

	/**
	 * Init connection is call when a Server object need to be connect on the net work.
	 */
	JS.GetJsRuntimeManager().AppendFunction("CargoWebServer.initConnection",
		func(adress string, openCallback string, closeCallback string, connectionId string, service otto.Value, caller otto.Value) otto.Value {
			log.Println("--> init connection with : ", adress, " session id: ", connectionId)

			// Get the new connection id.
			subConnection, err := GetServer().connect(adress)

			// The new created connection Js object.
			var conn otto.Value

			if err != nil {
				return conn
			}

			subConnectionId := subConnection.GetUuid()

			// I will append the connection the session.
			GetServer().appendSubConnectionId(connectionId, subConnectionId)

			// Here I will create the connection object...
			// .RunScript(connectionId, "new Connection()")
			conn, err = JS.GetJsRuntimeManager().GetSession(connectionId).Run("new Connection()")
			if err != nil {
				log.Println("---------> error found!", err)
			}

			// I will set the connection id.
			conn.Object().Set("id", subConnectionId)

			// Set the connection in the caller.
			service.Object().Set("conn", conn)

			// I will set the open callback.
			//.RunScript(connectionId, "Connection.prototype.onopen = "+openCallback)
			_, err = JS.GetJsRuntimeManager().GetSession(connectionId).Run("Connection.prototype.onopen = " + openCallback)
			if err != nil {
				log.Println("-----> error!", err)
			}

			// Now the close callback.
			//.RunScript(connectionId, "Connection.prototype.onclose = "+closeCallback)
			_, err = JS.GetJsRuntimeManager().GetSession(connectionId).Run("Connection.prototype.onclose = " + closeCallback)
			if err != nil {
				log.Println("-----> error!", err)
			}

			// Keep the connection link...
			GetServer().subConnections[subConnectionId] = conn

			// I will get the client code and inject it in the vm.
			id := Utility.RandomUUID()

			method := "GetServicesClientCode"
			params := make([]*MessageData, 0)
			to := make([]connection, 1)
			to[0] = subConnection

			successCallback := func(connectionId string, conn otto.Value, service otto.Value) func(rspMsg *message, caller interface{}) {
				return func(rspMsg *message, caller interface{}) {
					defer func() {
						// Stahp mean the VM was kill by the admin.
						if caught := recover(); caught != nil {
							if caught.(error).Error() == "Stahp" {
								// Here the task was cancel.
								return
							} else {
								panic(caught) // Something else happened, repanic!
							}
						}
					}()

					src := string(rspMsg.msg.Rsp.Results[0].DataBytes)
					JS.GetJsRuntimeManager().AppendScript("CargoWebServer", src)
					// Call on open...
					conn.Object().Call("onopen", service, caller)
				}
			}(connectionId, conn, service)

			errorCallback := func(rspMsg *message, caller interface{}) {
				log.Println("GetServicesClientCode error!!!")
			}

			rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, caller)

			go func(rqst *message) {
				GetServer().GetProcessor().m_sendRequest <- rqst
			}(rqst)

			return conn

		})

	////////////////////////////////////////////////////////////////////////////
	// Services intialisation.
	////////////////////////////////////////////////////////////////////////////
	defer func() {
		// Now after all initialisation are done I will open connection with
		// other servers.
		this.GetDataManager().openConnections() // That will also append entities scripts.

		// Javacript initialisation here.
		JS.GetJsRuntimeManager().OpenSession("") // Set the anonymous session.

		// Append services scripts.
		for id, src := range GetServer().GetServiceManager().m_serviceClientSrc {
			JS.GetJsRuntimeManager().AppendScript("CargoWebServer/"+id, src)
		}

		// Initialyse the script for the default session.
		JS.GetJsRuntimeManager().InitScripts() // Run the script for the default session.

		// Set service in the server object.
		for serviceName, _ := range GetServer().GetServiceManager().m_serviceClientSrc {
			log.Println("--> Load", serviceName, "service script")
			_, err := JS.GetJsRuntimeManager().RunScript("", "server."+strings.ToLower(serviceName[0:1])+serviceName[1:]+" = new "+serviceName+"();")
			if err != nil {
				log.Fatal(err)
			}
		}

		// Now I will register actions for services container.
		activeConfigurationsEntity, err := GetServer().GetConfigurationManager().getActiveConfigurationsEntity()
		if err != nil {
			log.Panicln(err)
		}

		activeConfigurations := activeConfigurationsEntity.GetObject().(*Config.Configurations)

		// Now I will set scheduled task.
		for i := 0; i < len(activeConfigurations.M_scheduledTasks); i++ {
			task := activeConfigurations.M_scheduledTasks[i]
			GetTaskManager().scheduleTask(task)
		}

		// Here I will initialise the search engine.
		// Sync files
		GetServer().GetFileManager().synchronizeAll()
		// Sync users, computers and groups.
		GetServer().GetLdapManager().synchronizeAll()
		// Sync projects.
		GetServer().GetProjectManager().synchronize()
	}()
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

	// Kill the running command if there so...
	for i := 0; i < len(this.cmds); i++ {
		if this.cmds[i].Process != nil {
			this.cmds[i].Process.Kill()
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
func (this *Server) connect(address string) (connection, error) {

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

	// Open the a new connection with the server.
	if host == this.GetConfigurationManager().GetHostName() && this.GetConfigurationManager().GetServerPort() == port {
		return nil, errors.New("Loopback connection!")
	}

	// If a connection already exist I will use it...
	conn := this.getConnectionByIp(address, port)
	if conn != nil {
		return conn, nil
	}

	// Create the new connection.
	if socket == "ws" {
		conn = NewWebSocketConnection()
	} else {
		conn = NewTcpSocketConnection()
	}

	err := conn.Open(host, port)
	if err != nil {
		return nil, err // The connection fail...
	}

	// Test if the connection is open.
	if !conn.IsOpen() {
		return nil, errors.New("Fail to open connection with socket " + host + " at port " + strconv.Itoa(port))
	}

	log.Println("--> connection whit ", host, " at port ", port, " is now open!")

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

/////////////////////////////////////////////////////////
// Call cmd from server.
/////////////////////////////////////////////////////////

/**
 * Remove a command from list of running command.
 */
func (server *Server) removeCmd(cmd *exec.Cmd) {
	var cmds []*exec.Cmd
	for i := 0; i < len(server.cmds); i++ {
		if server.cmds[i] != cmd {
			cmds = append(cmds, cmd)
		}
	}
	server.cmds = cmds

	for sessionId, cmds_ := range server.sessionCmds {
		var cmds []*exec.Cmd
		for i := 0; i < len(cmds_); i++ {
			if cmds_[i] != cmd {
				cmds = append(cmds, cmds_[i])
			}
		}
		server.sessionCmds[sessionId] = cmds
	}

	// Kill it process.
	if cmd != nil {
		if cmd.Process != nil {
			err := cmd.Process.Kill()
			if err != nil {
				log.Println("Fail to kill command ", err)
			}
		}
	}
}

// Run starts the specified command and waits for it to complete.
// @param {string} name The name of the command to run.
// @param {[]string} The list of command arguments.
// @param {string} sessionId The user session.
func (server *Server) RunCmd(name string, args []string, sessionId string) interface{} {
	if runtime.GOOS == "windows" && !strings.HasSuffix(name, ".exe") {
		name += ".exe"
	}

	// The first step will be to start the service manager.
	path := server.GetConfigurationManager().GetBinPath() + "/" + name

	// In the case that the command is not in the bin path I will
	// try to run it from the system path.
	if !Utility.Exists(path) {
		path = name
	}

	// Set the command
	cmd := exec.Command(path)
	cmd.Args = append(cmd.Args, args...)

	// the command succed here.
	server.cmds = append(server.cmds, cmd)

	// Now I will register the command with the session.
	server.sessionCmds[sessionId] = append(server.sessionCmds[sessionId], cmd)

	// Call it...
	output, err := cmd.Output()

	if err != nil {
		log.Println("Fail to run cmd: ", name)
		log.Println("error: ", err)
		return err
	}

	return string(output)
}
