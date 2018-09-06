package GoJavaScript

import "strconv"
import "log"

// Global variable.
var (
	// Channel use by library code to call Go function on the client side of
	// the lib.
	Call_remote_actions_chan chan *Action
)

// Go function reside in the client, a remote call will be made here.
func CallGoFunction(target string, name string, params ...interface{}) (interface{}, error) {

	action := NewAction(name, target)

	// Set the list of parameters.
	for i := 0; i < len(params); i++ {
		action.AppendParam("arg"+strconv.Itoa(i), params[i])
	}

	// Create the channel to give back the action
	// when it's done.
	action.SetDone()

	// Send the action to the client side.
	Call_remote_actions_chan <- action

	// Set back the action with it results in it.
	action = <-action.GetDone()

	var err error
	if action.Results[1] != nil {
		log.Println("action  ", name, action.Results[0])
		log.Println("action error ", name, action.Results[1])
		err = action.Results[1].(error)
	}

	return action.Results[0], err
}

/**
 * The JavaScript JS engine.
 */
type Engine interface {

	/**
	 * Init and start the engine.
	 * port The port to communicate with the engine, or the debbuger.
	 */
	Start(port int)

	/////////////////// Global variables //////////////////////

	/**
	 * Set a variable on the global context.
	 * name The name of the variable in the context.
	 * value The value of the variable, can be a string, a number,
	 */
	SetGlobalVariable(name string, value interface{})

	/**
	 * Return a variable define in the global object.
	 */
	GetGlobalVariable(name string) (Value, error)

	/////////////////// Objects //////////////////////
	/**
	 * Create JavaScript object with given uuid. If name is given the object will be
	 * set a global object property.
	 */
	CreateObject(uuid string, name string)

	/**
	 * Set an object property.
	 * uuid The object reference.
	 * name The name of the property to set
	 * value The value of the property
	 */
	SetObjectProperty(uuid string, name string, value interface{}) error

	/**
	 * That function is use to get Js object property
	 */
	GetObjectProperty(uuid string, name string) (Value, error)

	/**
	 * Create an empty array of a given size and set it as object property.
	 */
	CreateObjectArray(uuid string, name string, size uint32) error

	/**
	 * Set an object property.
	 * uuid The object reference.
	 * name The name of the property to set
	 * index The index of the object in the array
	 * value The value of the property
	 */
	SetObjectPropertyAtIndex(uuid string, name string, index uint32, value interface{})

	/**
	 * That function is use to get Js obeject property
	 */
	GetObjectPropertyAtIndex(uuid string, name string, index uint32) (Value, error)

	/**
	 * set to JS object a Go function.
	 */
	SetGoObjectMethod(uuid, name string) error

	/**
	 * Set to JS object a Js function.
	 */
	SetJsObjectMethod(uuid, name string, src string) error

	/**
	 * Call object methode.
	 */
	CallObjectMethod(uuid string, name string, params ...interface{}) (Value, error)

	/////////////////// Functions //////////////////////

	/**
	 * Register a go function in JS
	 */
	RegisterGoFunction(name string)

	/**
	 * Parse and set a function in the Javascript.
	 * name The name of the function (the function will be keep in the engine for it
	 *      lifetime.
	 * args The argument name for that function.
	 * src  The body of the function
	 * options Can be JERRY_PARSE_NO_OPTS or JERRY_PARSE_STRICT_MODE
	 */
	RegisterJsFunction(name string, src string) error

	/**
	 * Call a Javascript function. The function must exist...
	 */
	CallFunction(name string, params []interface{}) (Value, error)

	/**
	 * Evaluate a script.
	 * script Contain the code to run.
	 * variables Contain the list of variable to set on the global context before
	 * running the script.
	 */
	EvalScript(script string, variables []interface{}) (Value, error)

	/**
	 * Clear the VM
	 */
	Clear()
}
