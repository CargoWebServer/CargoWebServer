// GoJerryScriptClient project GoJerryScriptClient.go
package GoJerryScriptClient

import (
	"fmt"
	"log"

	"strconv"
	"time"

	"encoding/json"
	//	"errors"
	"os/exec"
	"reflect"

	"code.myceliUs.com/GoJerryScript"
	"code.myceliUs.com/Utility"
)

var (
	// Callback function used by dynamic type, it's call when an entity is set.
	// Can be use to store dynamic type in a cache.
	SetEntity func(interface{}) = func(val interface{}) {
		log.Println("---> set entity ", val)
	}
)

// That act as a remote connection with the engine.
type Client struct {
	isRunning        bool
	peer             *GoJerryScript.Peer
	exec_action_chan chan *GoJerryScript.Action
	srv              *exec.Cmd
}

// Create a new client session with jerry script server.
func NewClient(address string, port int) *Client {

	client := new(Client)
	client.isRunning = true

	// Open the action channel.
	client.exec_action_chan = make(chan *GoJerryScript.Action, 0)

	// Create the peer.
	client.peer = GoJerryScript.NewPeer(address, port, client.exec_action_chan)

	var err error
	// Here I will start the external server process.
	client.srv = exec.Command("/home/dave/Documents/CargoWebServer/Project/src/code.myceliUs.com/GoJerryScript/GoJerryScriptServer/GoJerryScriptServer", strconv.Itoa(port))
	err = client.srv.Start()

	if err != nil {
		log.Println("Fail to start GoJerryScriptServer", err)
		return nil
	}

	// Create the client connection, try 5 time and wait 200 millisecond each try.
	for i := 0; i < 5; i++ {
		time.Sleep(50 * time.Millisecond)
		err = client.peer.Connect(address, port)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Println("Fail to connect with server ", address, ":", port, err)
		return nil
	}

	// process actions.
	go client.processActions()

	return client
}

// Call a go function and return it result.
func (self *Client) callGoFunction(name string, params ...interface{}) (interface{}, error) {
	results, err := Utility.CallFunction(name, params...)
	if err != nil {
		return nil, err
	}

	if len(results) > 0 {
		if len(results) == 1 {
			// One result here...
			switch goError := results[0].Interface().(type) {
			case error:
				return nil, goError
			}
			return results[0].Interface(), nil
		} else {
			// Return the result and the error after.
			results_ := make([]interface{}, 0)
			for i := 0; i < len(results); i++ {
				switch goError := results[i].Interface().(type) {
				case error:
					return nil, goError
				}
				results_ = append(results_, results[i].Interface())
			}
			return results_, nil
		}
	}
	return nil, nil
}

// Run requested actions.
func (self *Client) processActions() {
	for self.isRunning {
		select {
		case action := <-self.exec_action_chan:
			go func(a *GoJerryScript.Action) {
				log.Println("---> client exec action: ", a.UUID, a.Name)

				var target interface{}
				isJsObject := true
				if len(a.Target) > 0 {
					if a.Target == "Client" {
						// Call a methode of client
						target = self
						isJsObject = false
					} else {
						target = GoJerryScript.GetCache().GetObject(a.Target)
						isJsObject = reflect.TypeOf(target).String() == "GoJerryScript.Object" ||
							reflect.TypeOf(target).String() == "*GoJerryScript.Object"
					}
				}

				params := make([]interface{}, 0)
				for i := 0; i < len(a.Params); i++ {
					// So here I will append the value to parameters.
					log.Println("---> params: ", a.Params[i].Name, a.Params[i].Type, a.Params[i].Value)
					// TODO test if the parameter is a ObjectRef.
					params = append(params, a.Params[i].Value)
				}

				// I will call the function.
				var results interface{}
				var err interface{}

				if isJsObject {
					results, err = self.callGoFunction(a.Name, params...)
				} else {
					// Now I will call the method on the object.
					results, err = Utility.CallMethod(target, a.Name, params)
				}

				// if results are go struct
				// I must transfer it on the server side here...
				if results != nil {
					if reflect.TypeOf(results).Kind() == reflect.Slice {
						// Here if the result is a slice I will test if it contains struct...
						slice := reflect.ValueOf(results)
						results_ := make([]interface{}, 0)

						for i := 0; i < slice.Len(); i++ {
							e := slice.Index(i)
							// I will derefence the pointer if it's a pointer.
							if reflect.TypeOf(e.Interface()).Kind() == reflect.Ptr {
								e = e.Elem()
							}

							// Test it
							if reflect.TypeOf(e.Interface()).Kind() == reflect.Struct {
								// results will be register.
								var uuid string
								if reflect.TypeOf(e.Interface()).String() != "GoJerryScript.Object" {
									uuid = self.RegisterGoObject(slice.Index(i).Interface(), "")
								} else {
									// No need to export the object function here
									// because it's already exist on the server side.
									uuid = results.(GoJerryScript.Object).UUID
								}
								// I will set the results a object reference.
								results_ = append(results_, GoJerryScript.ObjectRef{UUID: uuid})
							}
						}
						// Set the array of object references.
						results = results_
					} else {
						// I will test if the result is a structure or not...
						e := reflect.ValueOf(results)
						// I will derefence the pointer if it a pointer.
						if reflect.TypeOf(results).Kind() == reflect.Ptr {
							e = e.Elem()
						}
						// Test it
						if reflect.TypeOf(e.Interface()).Kind() == reflect.Struct {
							// results will be register.
							var uuid string
							if reflect.TypeOf(e.Interface()).String() != "GoJerryScript.Object" {
								uuid = self.RegisterGoObject(results, "")
							} else {
								// No need to export the object function here
								// because it's already exist on the server side.
								uuid = results.(GoJerryScript.Object).UUID
							}
							// I will set the results a object reference.
							results = GoJerryScript.ObjectRef{UUID: uuid}
						}
					}
				}

				// set the action result.
				a.AppendResults(results, err)

				// Here I will create the response and send it back to the client.
				a.Done <- a
			}(action)
		}
	}
}

////////////////////////////-- Api --////////////////////////////

// Create Go object
func (self *Client) CreateGoObject(jsonStr string) (interface{}, error) {
	log.Println("---> create object ", jsonStr)
	var value interface{}
	data := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err == nil {
		if data["TYPENAME"] != nil {
			relfectValue := Utility.MakeInstance(data["TYPENAME"].(string), data, SetEntity)
			value = relfectValue.Interface()
		} else {
			// Here map[string]interface{} will be use.
			return nil, nil //errors.New("No object are registered in go to back the js object!")

		}
	}

	return value, err
}

/**
 * Register a go type to be usable as JS type.
 */
func (self *Client) RegisterGoType(value interface{}) {
	// Register local object.
	Utility.RegisterType(value)
}

/**
 * That function is use to create Js object representation in the Js engine from
 * go object. The tow object (JS/Go) are releated by a uuid. The Go object live
 * in the client side and the Js in the server side where the Js engine is.
 * obj The go object to register
 * name If a name is given It will be a global variable.
 */
func (self *Client) RegisterGoObject(obj interface{}, name string) string {
	log.Println("---> register object: ", obj)
	// First of all I will create the uuid property.
	ptrString := fmt.Sprintf("%d", obj)
	uuid := Utility.GenerateUUID(ptrString)

	if GoJerryScript.GetCache().GetObject(uuid) != nil {
		// if the object is already register I will return
		// this avoid circular call
		return uuid
	}

	// Here I will keep the object in the client cache.
	GoJerryScript.GetCache().SetObject(uuid, obj)

	// Now I will create the object on the server.
	createObjectAction := new(GoJerryScript.Action)
	createObjectAction.UUID = Utility.RandomUUID()
	createObjectAction.Name = "CreateObject"
	createObjectAction.AppendParam("uuid", uuid)
	createObjectAction.AppendParam("name", name) // no name is nessarry here.

	// I will create the object.
	self.peer.CallRemoteAction(createObjectAction)

	// Now I will register it member attributes.
	element := reflect.ValueOf(obj)
	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		element = reflect.ValueOf(obj).Elem()
	}

	// Register object method.
	for i := 0; i < element.Addr().NumMethod(); i++ {
		typeMethod := element.Addr().Type().Method(i)
		methodName := typeMethod.Name
		action := new(GoJerryScript.Action)
		action.UUID = Utility.RandomUUID()
		action.Name = "SetGoObjectMethod"
		action.AppendParam("uuid", uuid)
		action.AppendParam("name", methodName)
		self.peer.CallRemoteAction(action)
	}

	for i := 0; i < element.NumField(); i++ {
		valueField := element.Field(i)
		typeField := element.Type().Field(i)

		// Here I will set the property
		// Bust the number of field handler here.
		if valueField.CanInterface() {
			// So here is the field.
			fieldValue := valueField.Interface()
			fieldName := typeField.Name

			// Now depending of the type of the field I will do a recursion.
			var typeOf = reflect.TypeOf(fieldValue)

			// Dereference the pointer as needed.
			if typeOf.Kind() == reflect.Ptr {
				// Here I will access the value and not the pointer...
				fieldValue = valueField.Elem().Interface()
			}

			if typeOf.Kind() == reflect.Slice {
				if valueField.Len() > 0 {
					isObject := false
					for i := 0; i < valueField.Len(); i++ {
						fieldValue := valueField.Index(i).Interface()
						// If the value is a struct I need to register it to JS
						if reflect.TypeOf(fieldValue).Kind() == reflect.Struct || reflect.TypeOf(fieldValue).Kind() == reflect.Ptr {
							if reflect.TypeOf(reflect.ValueOf(fieldValue).Elem()).Kind() == reflect.Struct {
								// Here the element is a structure so I need to create it representation.
								uuid_ := self.RegisterGoObject(fieldValue, "")
								if i == 0 {
									action := new(GoJerryScript.Action)
									action.UUID = Utility.RandomUUID()
									action.Name = "CreateObjectArray"
									action.AppendParam("uuid", uuid)
									action.AppendParam("name", fieldName)
									action.AppendParam("size", uint32(valueField.Len()))
									self.peer.CallRemoteAction(action)
								}

								// Set the object in the array the object property.
								action := new(GoJerryScript.Action)
								action.UUID = Utility.RandomUUID()
								action.Name = "SetObjectPropertyAtIndex"
								action.AppendParam("uuid", uuid)
								action.AppendParam("name", fieldName)
								action.AppendParam("index", uint32(i))
								action.AppendParam("value", GoJerryScript.ObjectRef{UUID: uuid_})
								self.peer.CallRemoteAction(action)
								isObject = true
							}
						}
					}

					// if the array dosent contain structure I will set it values...
					if !isObject {
						action := new(GoJerryScript.Action)
						action.UUID = Utility.RandomUUID()
						action.Name = "SetObjectProperty"
						action.AppendParam("uuid", uuid)
						action.AppendParam("name", fieldName)
						action.AppendParam("value", fieldValue)
						self.peer.CallRemoteAction(action)
					}
				}
			} else if typeOf.Kind() == reflect.Struct {
				// Set the struct...
				uuid_ := self.RegisterGoObject(fieldValue, "")

				// Set the struct as object property.
				action := new(GoJerryScript.Action)
				action.UUID = Utility.RandomUUID()
				action.Name = "SetObjectProperty"
				action.AppendParam("uuid", uuid)
				action.AppendParam("name", fieldName)
				action.AppendParam("value", GoJerryScript.ObjectRef{UUID: uuid_})
				self.peer.CallRemoteAction(action)

			} else {
				// basic type.
				action := new(GoJerryScript.Action)
				action.UUID = Utility.RandomUUID()
				action.Name = "SetObjectProperty"
				action.AppendParam("uuid", uuid)
				action.AppendParam("name", fieldName)
				action.AppendParam("value", fieldValue)
				self.peer.CallRemoteAction(action)
			}
		}
	}

	return uuid
}

/**
 * Register a go function to bu usable in JS (in the global object)
 */
func (self *Client) RegisterGoFunction(name string, fct interface{}) {
	log.Println("---> register go function: ", name)

	// Keep the function in the local client.
	Utility.RegisterFunction(name, fct)

	// Register the function in the server.

	// Create the action.
	action := new(GoJerryScript.Action)
	action.UUID = Utility.RandomUUID()

	// The name of the action to execute.
	action.Name = "RegisterGoFunction"

	// Append the name parameter.
	action.AppendParam("name", name)

	action = self.peer.CallRemoteAction(action)
}

/**
 * Register a JavaScript function in the interpreter.
 * name The name of the js function
 * args The list of arguments
 * src  The function js code.
 */
func (self *Client) RegisterJsFunction(name string, src string) error {
	// Create the action.
	action := new(GoJerryScript.Action)
	action.UUID = Utility.RandomUUID()

	// The name of the action to execute.
	action.Name = "RegisterJsFunction"

	action.AppendParam("name", name)
	action.AppendParam("src", src)

	// Call the action.
	action = self.peer.CallRemoteAction(action)

	if action.Results[0] != nil {
		return action.Results[0].(error)
	}

	return nil
}

/**
 * Create a new Js object.
 */
func (self *Client) CreateObject(name string) GoJerryScript.Object {
	// Create the object.
	obj := GoJerryScript.NewObject(name)

	// Give object the peer so it can register itself with the server.
	obj.SetPeer(self.peer)

	return *obj
}

/**
 * Evaluate sript.
 * The list of global variables to be set before executing the script.
 */
func (self *Client) EvalScript(script string, variables GoJerryScript.Variables) (GoJerryScript.Value, error) {
	// So here I will create the function parameters.
	action := new(GoJerryScript.Action)
	action.UUID = Utility.RandomUUID()

	// The name of the action to execute.
	action.Name = "EvalScript"

	action.AppendParam("script", script)
	action.AppendParam("variables", variables)

	// Call the remote action
	action = self.peer.CallRemoteAction(action)

	var err error

	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	return action.Results[0].(GoJerryScript.Value), err
}

/**
 * Call a JS/Go function with a given name and a list of parameter.
 */
func (self *Client) CallFunction(name string, params ...interface{}) (GoJerryScript.Value, error) {
	// So here I will create the function parameters.
	action := new(GoJerryScript.Action)
	action.UUID = Utility.RandomUUID()

	// The name of the action to execute.
	action.Name = "CallFunction"
	action.AppendParam("name", name)
	action.AppendParam("params", params)

	// Call the remote action
	action = self.peer.CallRemoteAction(action)
	var err error
	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	return action.Results[0].(GoJerryScript.Value), err
}

/**
 * Set the global variable name.
 */
func (self *Client) SetGlobalVariable(name string, value interface{}) {
	// So here I will create the function parameters.
	action := new(GoJerryScript.Action)
	action.UUID = Utility.RandomUUID()

	// The name of the action to execute.
	action.Name = "SetGlobalVariable"
	action.AppendParam("name", name)
	action.AppendParam("value", value)

	// Call the remote action
	action = self.peer.CallRemoteAction(action)
}

/**
 * Get the global variable name.
 */
func (self *Client) GetGlobalVariable(name string) (GoJerryScript.Value, error) {
	action := new(GoJerryScript.Action)
	action.UUID = Utility.RandomUUID()

	// The name of the action to execute.
	action.Name = "GetGlobalVariable"
	action.AppendParam("name", name)

	// Call the remote action
	action = self.peer.CallRemoteAction(action)

	var err error
	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	return action.Results[0].(GoJerryScript.Value), err
}

/**
 * Stop JerryScript.
 */
func (self *Client) Stop() bool {
	// Create the action.
	/*action := new(GoJerryScript.Action)
	action.UUID = Utility.RandomUUID()

	// The name of the action to execute.
	action.Name = "Stop"

	action = self.peer.CallRemoteAction(action)*/

	// Stop action proecessing
	self.isRunning = false

	// stop it peer.
	self.peer.Close()

	// stop it server.
	self.srv.Process.Kill()

	return true
}
