// Action
package Server

import (
	//"log"
	"os/exec"
	"reflect"
	"strconv"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/JS"
	"code.myceliUs.com/Utility"
)

/**
 * This data structure keep information about Action to be done by the
 */
type Action struct {
	// The message who generate this Action...
	msg *message

	// The name of the Action, can contain a dot if is a method of an object,
	// ex like the data store manager, dataStore.get
	Name string

	/** The parameters **/
	Params []interface{}

	/** The parameters type name **/
	ParamTypeNames []string

	/** The parameters name **/
	ParamNames []string
}

/**
 * Create and intialyse basic Action info.
 */
func newAction(name string, msg *message) *Action {
	a := new(Action)
	a.Name = name
	a.msg = msg
	a.Params = make([]interface{}, 0, 0)
	a.ParamTypeNames = make([]string, 0, 0)
	a.ParamNames = make([]string, 0, 0)
	return a
}

/**
 * Send a response to the caller of Action.
 */
func (self *Action) sendResponse(result []*MessageData) {
	// Respond back to the source...
	to := make([]connection, 1)
	to[0] = self.msg.from
	resultMsg, _ := NewResponseMessage(self.msg.GetId(), result, to)
	GetServer().GetProcessor().appendResponse(resultMsg)
}

/**
 * Execute the Action iteself...
 */
func (self *Action) execute() {

	// That function use reflection to retreive the
	// method to call on a given object.
	x, errMsg := Utility.CallMethod(*self, self.Name, self.Params)

	// Get the session id and the message id...
	var sessionId string
	if self.msg.from != nil {
		sessionId = self.msg.from.GetUuid()
	}

	messageId := self.msg.GetId()

	if errMsg != nil {
		err := errMsg.(error)

		// Create the error object.
		cargoError := NewError(Utility.FileLine(), ACTION_EXECUTE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	xt := reflect.TypeOf(x).Kind()
	// Create the data element...
	r := new(MessageData)
	r.TYPENAME = "Server.MessageData"
	r.Name = "result"

	if xt == reflect.Int || xt == reflect.Int64 {
		r.Value = strconv.FormatInt(x.(int64), 10)
	} else if xt == reflect.Bool {
		r.Value = strconv.FormatBool(x.(bool))
	} else if xt == reflect.String {
		r.Value = x
	} else if xt == reflect.Struct {
		r.Value = x
	} else { // Must be array of bytes...
		r.Value = x
	}

	// Set the result and put in the
	result := make([]*MessageData, 1)
	result[0] = r
	self.sendResponse(result)
}

////////////////////////////////////////////////////////////////////////////////
//              			Listeners Actions...
////////////////////////////////////////////////////////////////////////////////

/**
 * Register a new listener with a given name.
 */
func (self *Action) RegisterListener(name string) {
	if self.msg.from.IsOpen() {
		listener := NewEventListener(name, self.msg.from)
		GetServer().GetEventManager().AddEventListener(listener)
	}
}

/**
 * Remove a listener with a given name.
 */
func (self *Action) UnregisterListener(name string) {
	GetServer().GetEventManager().RemoveEventListener(self.msg.from.GetUuid(), name)
}

/**
 * The client must know it session id, it simply return it...
 */
func (self *Action) GetSessionId() string {
	if self.msg.from.IsOpen() {
		return self.msg.from.GetUuid()
	}
	return ""
}

/**
 * That function return the client services code.
 * The code must be inject in the client JS interpreter in order
 * to access server side service.
 * The map contain the service id as key and the service source code as value.
 */
func (self *Action) GetServicesClientCode() map[string]string {
	// Simply return the internal map.
	return GetServer().GetServiceManager().m_serviceClientSrc
}

/**
 * Execute a vb script cmd.
 * * Windows only...
 */
func runVbs(scriptName string, args []string) ([]string, error) {
	path := GetServer().GetConfigurationManager().GetScriptPath() + "/" + scriptName
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

/**
 * Execute a vb script.
 */
func (self *Action) ExecuteVbScript(scriptName string, args []string) []string {

	// Run the given script on the server side.
	results, err := runVbs(scriptName, args)

	// Get the session id and the message id...
	if err != nil {
		sessionId := self.msg.from.GetUuid()
		messageId := self.msg.GetId()

		// Create the error object.
		cargoError := NewError(Utility.FileLine(), ACTION_EXECUTE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	return results
}

////////////////////////////////////////////////////////////////////////////////
//					Script execution releated Actions...
////////////////////////////////////////////////////////////////////////////////
/**
 * That function is the most important function of the framework. It use
 */
func (self *Action) ExecuteJsFunction(funtionStr string, funtionParams ...interface{}) (results []interface{}, jsError error) {
	var sessionId string
	if self.msg.from != nil {
		sessionId = self.msg.from.GetUuid()
	}

	// Call the function on the Js runtime.
	results, jsError = JS.GetJsRuntimeManager().ExecuteJsFunction(self.msg.GetId(), sessionId, funtionStr, funtionParams)

	if jsError != nil {
		// Here the user made an error inside is js code, i will simply report
		// he's error...
		return nil, jsError
	}

	// If the results[0] is an error, I will return an error...
	// The error is throw by the golang functor
	if len(results) > 0 {
		switch goError := results[0].(type) {
		case error:
			return nil, goError
		}
	}

	// Here there is no error, the functor do it's job...
	return
}

/**
 * Return the pong message to keep connection alive.
 */
func (self *Action) Ping() (string, error) {
	return "pong", nil
}

/**
 * Simply call RunCmd on the server.
 */
func (self *Action) RunCmd(name string, args []string) (result string, err error) {
	return GetServer().RunCmd(name, args)
}
