// +build linux darwin
package GoChakra

import (
	"log"

	"runtime"

	"code.myceliUs.com/GoJavaScript"
	"code.myceliUs.com/Utility"
)

/**
 * The Chacra JS engine.
 */
type Engine struct {
	port int

	// That channel will be use to process action requested on the runtime.
	actions chan map[string]interface{}

	// Global value that must be set on each newly created runtime.

	// That map contain the js function source code.
	jsFunctions map[string]string

	// That map contain the list of go function callable on the client side form
	// the server.
	goFunctions []string

	// The list of global variables
	variables map[string]interface{}

	// The main runtime.
	runtime *Runtime

	// Stop main runtime channel.
	stop chan bool
}

/**
 * Init and start the engine.
 * port The port to communicate with the engine, or the debbuger.
 */
func (self *Engine) Start(port int) {

	// The channel of actions to be process on the runtime.
	self.actions = make(chan map[string]interface{}, 0)

	// Initialisation of memory...
	self.jsFunctions = make(map[string]string, 0)
	self.goFunctions = make([]string, 0)
	self.variables = make(map[string]interface{}, 0)
	self.stop = make(chan bool)
	self.runtime = self.newRuntime(self.stop)
}

/**
 * Each engine will run in it own thread so
 */
func (self *Engine) newRuntime(stop chan bool) *Runtime {
	// Start processing action here.

	// The rutime must run in one thread only or context and rutime variable
	// will be lost.

	// The default runtime.
	var jsRuntime *Runtime
	wait := make(chan bool)

	// Start the runtime processing loop.
	go func() {
		// Fix the goroutine in it own thread.
		runtime.LockOSThread()

		// Start the runtime in it thread
		jsRuntime = new(Runtime)
		jsRuntime.Start()

		// Here I will set the global thing.

		// The js functions
		for name, src := range self.jsFunctions {
			jsRuntime.RegisterJsFunction(name, src)
		}

		// The go functions
		for i := 0; i < len(self.goFunctions); i++ {
			jsRuntime.RegisterGoFunction(self.goFunctions[i])
		}

		// The variables.
		for name, variable := range self.variables {
			jsRuntime.SetGlobalVariable(name, variable)
		}

		log.Println("---> start process message. 91 ", getCurrentContext())
		wait <- true

		// Start the processing loop.
		for {
			select {
			case <-stop:
				// exist process loop.
				log.Println("---> stop runtime ", getCurrentContext())
				// give back the thread to the os
				runtime.UnlockOSThread()
				return // exit the process function.

			case action := <-self.actions:
				log.Println("---> 46 action receive ", action["id"].(string))
				if action["id"] == "SetGlobalVariable" {
					jsRuntime.SetGlobalVariable(action["name"].(string), action["value"])
					action["done"].(chan bool) <- true
				} else if action["id"] == "CreateObject" {
					jsRuntime.CreateObject(action["uuid"].(string), action["name"].(string))
					action["done"].(chan bool) <- true
				} else if action["id"] == "SetObjectPropertyAtIndex" {
					jsRuntime.SetObjectPropertyAtIndex(action["uuid"].(string), action["name"].(string), action["i"].(uint32), action["value"])
					action["done"].(chan bool) <- true
				} else if action["id"] == "Clear" {
					jsRuntime.Clear()
					action["done"].(chan bool) <- true
					break // exit the processing loop.
				} else if action["id"] == "RegisterJsFunction" {
					action["error"].(chan error) <- jsRuntime.RegisterJsFunction(action["name"].(string), action["src"].(string))
				} else if action["id"] == "RegisterGoFunction" {
					jsRuntime.RegisterGoFunction(action["name"].(string))
					action["done"].(chan bool) <- true
				} else if action["id"] == "SetGoObjectMethod" {
					action["error"].(chan error) <- jsRuntime.SetGoObjectMethod(action["uuid"].(string), action["name"].(string))
				} else if action["id"] == "SetJsObjectMethod" {
					action["error"].(chan error) <- jsRuntime.SetJsObjectMethod(action["uuid"].(string), action["name"].(string), action["src"].(string))
				} else if action["id"] == "SetObjectProperty" {
					action["error"].(chan error) <- jsRuntime.SetObjectProperty(action["uuid"].(string), action["name"].(string), action["value"])
				} else if action["id"] == "CreateObjectArray" {
					action["error"].(chan error) <- jsRuntime.CreateObjectArray(action["uuid"].(string), action["name"].(string), action["size"].(uint32))
				} else if action["id"] == "GetObjectProperty" {
					results := make([]interface{}, 2)
					results[0], results[1] = jsRuntime.GetObjectProperty(action["uuid"].(string), action["name"].(string))
					action["results"].(chan []interface{}) <- results
				} else if action["id"] == "GetObjectPropertyAtIndex" {
					results := make([]interface{}, 2)
					results[0], results[1] = jsRuntime.GetObjectPropertyAtIndex(action["uuid"].(string), action["name"].(string), action["i"].(uint32))
					action["results"].(chan []interface{}) <- results
				} else if action["id"] == "GetGlobalVariable" {
					results := make([]interface{}, 2)
					results[0], results[1] = jsRuntime.GetGlobalVariable(action["name"].(string))
					action["results"].(chan []interface{}) <- results
				} else if action["id"] == "CallFunction" {
					results := make([]interface{}, 2)
					results[0], results[1] = jsRuntime.CallFunction(action["name"].(string), action["params"].([]interface{}))
					log.Println("---> function results: ", results)
					action["results"].(chan []interface{}) <- results
				} else if action["id"] == "EvalScript" {
					// Each script will be call in it own js runtime,
					// this is the only way to not blocking the main processing
					// loop.
					stop_ := make(chan bool)
					jsRuntime_ := self.newRuntime(stop_)

					results := make([]interface{}, 2)
					results[0], results[1] = jsRuntime_.EvalScript(action["script"].(string), action["variables"].([]interface{}))

					// Set back global variable after the script was done.
					for name, src := range self.jsFunctions {
						jsRuntime.RegisterJsFunction(name, src)
					}

					for i := 0; i < len(self.goFunctions); i++ {
						jsRuntime.RegisterGoFunction(self.goFunctions[i])
					}

					for name, variable := range self.variables {
						jsRuntime.SetGlobalVariable(name, variable)
					}

					// stop the jsRuntime.
					jsRuntime_.Clear()

					// stop the processing loop.
					stop_ <- true
					action["results"].(chan []interface{}) <- results

				} else if action["id"] == "CallObjectMethod" {
					results := make([]interface{}, 2)
					results[0], results[1] = jsRuntime.CallObjectMethod(action["uuid"].(string), action["name"].(string), action["params"].([]interface{})...)
					action["results"].(chan []interface{}) <- results
				}
				log.Println("---> 98 action done ", action["id"].(string))
			}
		}
	}()

	// Wait for the runtime to start in it own thread...
	<-wait

	return jsRuntime
}

/////////////////// Global variables //////////////////////

/**
 * Set a variable on the global context.
 * name The name of the variable in the context.
 * value The value of the variable, can be a string, a number,
 */
func (self *Engine) SetGlobalVariable(name string, value interface{}) {
	// Keep the variable on the map.
	self.variables[name] = value

	action := make(map[string]interface{})
	action["id"] = "SetGlobalVariable"
	action["name"] = name
	action["value"] = value

	action["done"] = make(chan bool, 0)
	self.actions <- action

	<-action["done"].(chan bool)
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

	// Keep the object as global variable.
	if len(name) > 0 {
		// Keep an object reference in the global variable map in that
		// case.
		self.variables[name] = GoJavaScript.NewObjectRef(uuid)
	}

	action := make(map[string]interface{})
	action["id"] = "CreateObject"
	action["uuid"] = uuid
	action["name"] = name

	action["done"] = make(chan bool, 0)
	self.actions <- action

	<-action["done"].(chan bool)
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
	// Keep the go function in memory...
	if !Utility.Contains(self.goFunctions, name) {
		self.goFunctions = append(self.goFunctions, name)
	}

	action := make(map[string]interface{})
	action["id"] = "RegisterGoFunction"
	action["name"] = name

	action["done"] = make(chan bool, 0)
	self.actions <- action

	<-action["done"].(chan bool)
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
	// I will keep the JS function in the map of functions
	self.jsFunctions[name] = src

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
	//log.Println("---> 493 ", getCurrentContext(), script)
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

	action["done"] = make(chan bool, 0)
	self.actions <- action

	<-action["done"].(chan bool)
}
