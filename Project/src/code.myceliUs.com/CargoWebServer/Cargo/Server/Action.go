// Action
package Server

import (
	//	"log"
	"reflect"
	"strconv"

	"code.myceliUs.com/CargoWebServer/Cargo/JS"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
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
}

/**
 * Create and intialyse basic Action info.
 */
func newAction(name string, msg *message) *Action {
	a := new(Action)
	a.Name = name
	a.msg = msg
	a.Params = make([]interface{}, 0, 0)
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

	// Remove the message from the pending message.
	GetServer().GetProcessor().removePendingRequest(self.msg)

	// That function use reflection to retreive the
	// method to call on a given object.
	x, errMsg := Utility.CallMethod(*self, self.Name, self.Params)

	if errMsg != nil {
		err := errMsg.(error)
		// Get the session id and the message id...
		sessionId := self.msg.from.GetUuid()
		messageId := self.msg.GetId()

		// Create the error object.
		cargoError := NewError(Utility.FileLine(), ACTION_EXECUTE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)

		return
	}

	xt := reflect.TypeOf(x).Kind()
	// Create the data element...
	r := new(MessageData)
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
		GetServer().GetProcessor().removePendingRequest(self.msg)
	}
}

/**
 * Remove a listener with a given name.
 */
func (self *Action) UnregisterListener(name string) {
	GetServer().GetEventManager().RemoveEventListener(self.msg.from.GetUuid(), name)
	GetServer().GetProcessor().removePendingRequest(self.msg)
}

/**
 * The client must know it session id, it simply return it...
 */
func (self *Action) GetSessionId() string {
	GetServer().GetProcessor().removePendingRequest(self.msg)
	if self.msg.from.IsOpen() {
		return self.msg.from.GetUuid()
	}
	return ""
}

////////////////////////////////////////////////////////////////////////////////
//					Script execution releated Actions...
////////////////////////////////////////////////////////////////////////////////
func (self *Action) ExecuteJsFunction(funtionStr string, funtionParams ...interface{}) (results []interface{}, jsError error) {

	// Call the function on the Js runtime.
	results, jsError = JS.GetJsRuntimeManager().ExecuteJsFunction(self.msg.GetId(), self.msg.from.GetUuid(), funtionStr, funtionParams)

	if jsError != nil {
		// Here the user made an error inside is js code, i will simply report
		// he's error...
		return nil, jsError
	}

	// If the results[0] is an error, I will return an error...
	// The error is throw by the golang functor
	switch goError := results[0].(type) {
	case error:
		return nil, goError
	}

	// Here there is no error, the functor do it's job...
	return results, nil
}
