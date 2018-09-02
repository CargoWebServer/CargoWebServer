package GoChakra

import (
	"log"
	"runtime"

	"code.myceliUs.com/GoJavaScript"
)

/**
 * The Chacra JS engine.
 */
type Engine struct {
	port int

	// That channel will be use to process action requested on the runtime.
	actions chan map[string]interface{}
}

/**
 * Init and start the engine.
 * port The port to communicate with the engine, or the debbuger.
 */
func (self *Engine) Start(port int) {

	// The channel of actions to be process on the runtime.
	self.actions = make(chan map[string]interface{}, 0)

	// Start processing action here.
	go func() {
		// The rutime must run in one thread only or context and rutime variable
		// will be lost.
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		runtime := new(Runtime)
		runtime.Start()

		for {
			select {
			// Process action here.
			case action := <-self.actions:
				//log.Println("---> action receive ", action)
				if action["id"] == "SetGlobalVariable" {
					runtime.SetGlobalVariable(action["name"].(string), action["value"])
				} else if action["id"] == "GetGlobalVariable" {
					results := make([]interface{}, 2)
					results[0], results[1] = runtime.GetGlobalVariable(action["name"].(string))
					action["results"].(chan []interface{}) <- results
				} else if action["id"] == "CreateObject" {
					runtime.CreateObject(action["uuid"].(string), action["name"].(string))
				} else if action["id"] == "SetObjectProperty" {
					action["error"].(chan error) <- runtime.SetObjectProperty(action["uuid"].(string), action["name"].(string), action["value"])
				} else if action["id"] == "GetObjectProperty" {
					results := make([]interface{}, 2)
					results[0], results[1] = runtime.GetObjectProperty(action["uuid"].(string), action["name"].(string))
					action["results"].(chan []interface{}) <- results
				} else if action["id"] == "CreateObjectArray" {
					action["error"].(chan error) <- runtime.CreateObjectArray(action["uuid"].(string), action["name"].(string), action["size"].(uint32))
				} else if action["id"] == "SetObjectPropertyAtIndex" {
					runtime.SetObjectPropertyAtIndex(action["uuid"].(string), action["name"].(string), action["i"].(uint32), action["value"])
				} else if action["id"] == "GetObjectPropertyAtIndex" {
					results := make([]interface{}, 2)
					results[0], results[1] = runtime.GetObjectPropertyAtIndex(action["uuid"].(string), action["name"].(string), action["i"].(uint32))
					action["results"].(chan []interface{}) <- results
				} else if action["id"] == "SetGoObjectMethod" {
					action["error"].(chan error) <- runtime.SetGoObjectMethod(action["uuid"].(string), action["name"].(string))
				} else if action["id"] == "SetJsObjectMethod" {
					action["error"].(chan error) <- runtime.SetJsObjectMethod(action["uuid"].(string), action["name"].(string), action["src"].(string))
				} else if action["id"] == "CallObjectMethod" {
					results := make([]interface{}, 2)
					results[0], results[1] = runtime.CallObjectMethod(action["uuid"].(string), action["name"].(string), action["params"].([]interface{})...)
					action["results"].(chan []interface{}) <- results
				} else if action["id"] == "RegisterGoFunction" {
					runtime.RegisterGoFunction(action["name"].(string))
				} else if action["id"] == "RegisterJsFunction" {
					log.Println("71 --> RegisterJsFunction")
					action["error"].(chan error) <- runtime.RegisterJsFunction(action["name"].(string), action["src"].(string))
				} else if action["id"] == "CallFunction" {
					results := make([]interface{}, 2)
					results[0], results[1] = runtime.CallFunction(action["name"].(string), action["params"].([]interface{}))
					action["results"].(chan []interface{}) <- results
				} else if action["id"] == "EvalScript" {
					results := make([]interface{}, 2)
					results[0], results[1] = runtime.EvalScript(action["script"].(string), action["variables"].([]interface{}))
					action["results"].(chan []interface{}) <- results
				} else if action["id"] == "Clear" {
					runtime.Clear()
				}
			}
		}

	}()

}

/////////////////// Global variables //////////////////////

/**
 * Set a variable on the global context.
 * name The name of the variable in the context.
 * value The value of the variable, can be a string, a number,
 */
func (self *Engine) SetGlobalVariable(name string, value interface{}) {
	action := make(map[string]interface{})
	action["id"] = "SetGlobalVariable"
	action["name"] = name
	action["value"] = value
	self.actions <- action
}

/**
 * Return a variable define in the global object.
 */
func (self *Engine) GetGlobalVariable(name string) (GoJavaScript.Value, error) {
	action := make(map[string]interface{})
	action["id"] = "GetGlobalVariable"
	action["name"] = name

	// The resturn channel
	action["results"] = make(chan []interface{}, 0)
	self.actions <- action

	results := <-action["results"].(chan []interface{})

	var err error
	if results[1] != nil {
		err = results[1].(error)
	}
	return results[0].(GoJavaScript.Value), err
}

/////////////////// Objects //////////////////////
/**
 * Create JavaScript object with given uuid. If name is given the object will be
 * set a global object property.
 */
func (self *Engine) CreateObject(uuid string, name string) {
	action := make(map[string]interface{})
	action["id"] = "CreateObject"
	action["uuid"] = uuid
	action["name"] = name
	self.actions <- action
}

/**
 * Set an object property.
 * uuid The object reference.
 * name The name of the property to set
 * value The value of the property
 */
func (self *Engine) SetObjectProperty(uuid string, name string, value interface{}) error {
	action := make(map[string]interface{})
	action["id"] = "SetObjectProperty"
	action["uuid"] = uuid
	action["name"] = name
	action["value"] = value

	// if there is error
	action["error"] = make(chan error, 0)

	self.actions <- action
	return <-action["error"].(chan error)

}

/**
 * That function is use to get Js object property
 */
func (self *Engine) GetObjectProperty(uuid string, name string) (GoJavaScript.Value, error) {
	action := make(map[string]interface{})
	action["id"] = "GetObjectProperty"
	action["uuid"] = uuid
	action["name"] = name

	// The resturn channel
	action["results"] = make(chan []interface{}, 0)

	self.actions <- action

	results := <-action["results"].(chan []interface{})
	var err error
	if results[1] != nil {
		err = results[1].(error)
	}
	return results[0].(GoJavaScript.Value), err
}

/**
 * Create an empty array of a given size and set it as object property.
 */
func (self *Engine) CreateObjectArray(uuid string, name string, size uint32) error {
	action := make(map[string]interface{})
	action["id"] = "CreateObjectArray"
	action["uuid"] = uuid
	action["name"] = name
	action["size"] = size

	// if there is error
	action["error"] = make(chan error, 0)
	self.actions <- action

	return <-action["error"].(chan error)
}

/**
 * Set an object property.
 * uuid The object reference.
 * name The name of the property to set
 * index The index of the object in the array
 * value The value of the property
 */
func (self *Engine) SetObjectPropertyAtIndex(uuid string, name string, i uint32, value interface{}) {
	action := make(map[string]interface{})
	action["id"] = "SetObjectPropertyAtIndex"
	action["uuid"] = uuid
	action["name"] = name
	action["i"] = i
	action["value"] = value
	self.actions <- action
}

/**
 * That function is use to get Js obeject property
 */
func (self *Engine) GetObjectPropertyAtIndex(uuid string, name string, i uint32) (GoJavaScript.Value, error) {
	action := make(map[string]interface{})
	action["id"] = "GetObjectPropertyAtIndex"
	action["uuid"] = uuid
	action["name"] = name
	action["i"] = i
	// The resturn channel
	action["results"] = make(chan []interface{}, 0)

	self.actions <- action

	results := <-action["results"].(chan []interface{})
	var err error
	if results[1] != nil {
		err = results[1].(error)
	}
	return results[0].(GoJavaScript.Value), err
}

/**
 * set object methode.
 */
func (self *Engine) SetGoObjectMethod(uuid, name string) error {
	action := make(map[string]interface{})
	action["id"] = "SetGoObjectMethod"
	action["uuid"] = uuid
	action["name"] = name
	// if there is error
	action["error"] = make(chan error, 0)
	self.actions <- action
	return <-action["error"].(chan error)
}

func (self *Engine) SetJsObjectMethod(uuid, name string, src string) error {
	action := make(map[string]interface{})
	action["id"] = "SetJsObjectMethod"
	action["uuid"] = uuid
	action["name"] = name
	action["src"] = src
	// if there is error
	action["error"] = make(chan error, 0)
	self.actions <- action
	return <-action["error"].(chan error)
}

/**
 * Call object methode.
 */
func (self *Engine) CallObjectMethod(uuid string, name string, params ...interface{}) (GoJavaScript.Value, error) {
	action := make(map[string]interface{})
	action["id"] = "CallObjectMethod"
	action["uuid"] = uuid
	action["name"] = name
	action["params"] = params
	// The resturn channel
	action["results"] = make(chan []interface{}, 0)
	self.actions <- action

	results := <-action["results"].(chan []interface{})
	var err error
	if results[1] != nil {
		err = results[1].(error)
	}

	return results[0].(GoJavaScript.Value), err
}

/////////////////// Functions //////////////////////

/**
 * Register a go function in JS
 */
func (self *Engine) RegisterGoFunction(name string) {
	action := make(map[string]interface{})
	action["id"] = "RegisterGoFunction"
	action["name"] = name
	self.actions <- action
}

/**
 * Parse and set a function in the Javascript.
 * name The name of the function (the function will be keep in the engine for it
 *      lifetime.
 * args The argument name for that function.
 * src  The body of the function
 * options Can be JERRY_PARSE_NO_OPTS or JERRY_PARSE_STRICT_MODE
 */
func (self *Engine) RegisterJsFunction(name string, src string) error {
	action := make(map[string]interface{})
	action["id"] = "RegisterJsFunction"
	action["name"] = name
	action["src"] = src
	// if there is error
	action["error"] = make(chan error, 0)
	self.actions <- action
	return <-action["error"].(chan error)
}

/**
 * Call a Javascript function. The function must exist...
 */
func (self *Engine) CallFunction(name string, params []interface{}) (GoJavaScript.Value, error) {
	action := make(map[string]interface{})
	action["id"] = "CallFunction"
	action["name"] = name
	action["params"] = params

	// The resturn channel
	action["results"] = make(chan []interface{}, 0)
	self.actions <- action

	results := <-action["results"].(chan []interface{})

	var err error
	if results[1] != nil {
		err = results[1].(error)
	}

	return results[0].(GoJavaScript.Value), err
}

/**
 * Evaluate a script.
 * script Contain the code to run.
 * variables Contain the list of variable to set on the global context before
 * running the script.
 */
func (self *Engine) EvalScript(script string, variables []interface{}) (GoJavaScript.Value, error) {
	action := make(map[string]interface{})
	action["id"] = "EvalScript"
	action["script"] = script
	action["variables"] = variables

	// The resturn channel
	action["results"] = make(chan []interface{}, 0)

	self.actions <- action
	results := <-action["results"].(chan []interface{})

	var err error
	if results[1] != nil {
		err = results[1].(error)
	}

	return results[0].(GoJavaScript.Value), err

}

func (self *Engine) Clear() {
	/* Cleanup the script engine. */
	action := make(map[string]interface{})
	action["id"] = "Clear"
	self.actions <- action
}
