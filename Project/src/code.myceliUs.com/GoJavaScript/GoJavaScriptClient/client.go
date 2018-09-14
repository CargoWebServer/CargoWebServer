// GoJavaScriptClient project GoJavaScriptClient.go
package GoJavaScriptClient

import (
	"log"

	"strconv"
	"time"

	"encoding/json"
	"errors"

	//"fmt"
	"os/exec"
	"reflect"

	"code.myceliUs.com/GoJavaScript"
	"code.myceliUs.com/Utility"
)

var (
	// Callback function used by dynamic type, it's call when an entity is set.
	// Can be use to store dynamic type in a cache.
	SetEntity func(interface{}) = func(val interface{}) {
		/** nothing todo here... **/
	}
)

// That act as a remote connection with the engine.
type Client struct {
	isRunning        bool
	peer             *GoJavaScript.Peer
	exec_action_chan chan *GoJavaScript.Action
	srv              *exec.Cmd
}

// Create a new client session with jerry script server.
func NewClient(address string, port int, name string) *Client {

	client := new(Client)
	client.isRunning = true

	// Open the action channel.
	client.exec_action_chan = make(chan *GoJavaScript.Action, 0)

	// Create the peer.
	client.peer = GoJavaScript.NewPeer(address, port, client.exec_action_chan)

	var err error
	// Here I will start the external server process.
	// Make Go intall for the GoJerryScriptServer to be in the /bin of go.
	client.srv = exec.Command("GoJavaScriptServer", strconv.Itoa(port), name)
	err = client.srv.Start()

	if err != nil {
		log.Println("Fail to start GoJavaScriptServer", err)
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

// Replace object reference by actual object.
func setObjectRefs(values []interface{}) []interface{} {

	for i := 0; i < len(values); i++ {
		if reflect.TypeOf(values[i]).Kind() == reflect.Slice {
			// Here I got a slice...
			slice := reflect.ValueOf(values[i])
			for j := 0; j < slice.Len(); j++ {
				e := slice.Index(j)
				if reflect.TypeOf(e.Interface()).String() == "GoJavaScript.ObjectRef" {
					// Replace the object reference with it actual object.
					values[i].([]interface{})[j] = GoJavaScript.GetCache().GetObject(e.Interface().(GoJavaScript.ObjectRef).UUID)
				} else if reflect.TypeOf(e.Interface()).String() == "*GoJavaScript.ObjectRef" {
					// Replace the object reference with it actual object.
					values[i].([]interface{})[j] = GoJavaScript.GetCache().GetObject(e.Interface().(*GoJavaScript.ObjectRef).UUID)
				}
			}
		} else {
			if reflect.TypeOf(values[i]).String() == "GoJavaScript.ObjectRef" {
				// Replace the object reference with it actual value.
				values[i] = GoJavaScript.GetCache().GetObject(values[i].(GoJavaScript.ObjectRef).UUID)
			} else if reflect.TypeOf(values[i]).String() == "*GoJavaScript.ObjectRef" {
				// Replace the object reference with it actual value.
				values[i] = GoJavaScript.GetCache().GetObject(values[i].(*GoJavaScript.ObjectRef).UUID)
			}
		}
	}

	return values
}

// Call a go function and return it result.
func (self *Client) callGoFunction(name string, params []interface{}) (interface{}, error) {
	results, err := Utility.CallFunction(name, setObjectRefs(params)...)
	if err != nil {
		log.Println("---> call go function ", name, " fail with error: ", err)
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

/**
 * Convert object, or objects to their uuid reference.
 */
func (self *Client) objectToRef(objects interface{}) interface{} {
	if objects != nil {
		if reflect.TypeOf(objects).Kind() == reflect.Slice {
			// Here if the result is a slice I will test if it contains struct...
			slice := reflect.ValueOf(objects)
			objects_ := make([]interface{}, 0)
			for i := 0; i < slice.Len(); i++ {
				e := slice.Index(i)
				// I will derefence the pointer if it's a pointer.
				for reflect.TypeOf(e.Interface()).Kind() == reflect.Ptr {
					e = e.Elem()
				}

				if reflect.TypeOf(e.Interface()).Kind() == reflect.Struct {
					// results will be register.
					uuid := self.RegisterGoObject(slice.Index(i).Interface(), "")
					// I will set the results a object reference.
					objects_ = append(objects_, GoJavaScript.NewObjectRef(uuid))
				}
			}
			// Set the array of object references.
			objects = objects_
		} else {
			// I will test if the result is a structure or not...
			e := reflect.ValueOf(objects)
			// I will derefence the pointer if it a pointer.
			for reflect.TypeOf(e.Interface()).Kind() == reflect.Ptr {
				e = e.Elem()
			}

			// if the object is a structure.
			if reflect.TypeOf(e.Interface()).Kind() == reflect.Struct {
				// results will be register.
				uuid := self.RegisterGoObject(objects, "")
				// I will set the results a object reference.
				objects = GoJavaScript.NewObjectRef(uuid)
			}
		}
	}

	return objects
}

// Action request by the server.
func (self *Client) processActions() {
	for self.isRunning {
		select {
		case action := <-self.exec_action_chan:
			go func(a *GoJavaScript.Action) {
				//log.Println("---> client action: ", a.Name)
				var target interface{}
				isJsObject := true
				if len(a.Target) > 0 {
					if a.Target == "Client" {
						// Call a methode of client
						target = self
						isJsObject = false
					} else {
						target = GoJavaScript.GetCache().GetObject(a.Target)
						isJsObject = reflect.TypeOf(target).String() == "GoJavaScript.Object" ||
							reflect.TypeOf(target).String() == "*GoJavaScript.Object"
					}
				}

				params := make([]interface{}, 0)
				for i := 0; i < len(a.Params); i++ {
					// So here I will append the value to parameters.
					params = append(params, GoJavaScript.GetObject(a.Params[i].Value))
				}

				// I will call the function.
				var results interface{}
				var err interface{}

				if isJsObject {
					results, err = self.callGoFunction(a.Name, params)
				} else {
					// Now I will call the method on the object.
					results, err = Utility.CallMethod(target, a.Name, params)
					if err != nil {
						log.Println("---> target ", target)
						log.Println("---> call Method: ", a.Name, params)
						log.Println("---> error: ", err)
					}
				}
				// Keep go object in client and transfert only reference.
				results = self.objectToRef(results)
				//log.Println("---> retrun objects ref results: ", results)
				// set the action result.
				a.AppendResults(results, err)
				// Here I will create the response and send it back to the client.
				a.GetDone() <- a
			}(action)
		}
	}
	// No more action to process.
	self.peer.Close()
}

////////////////////////////-- Api --////////////////////////////

// Create Go object
func (self *Client) CreateGoObject(jsonStr string) (interface{}, error) {
	log.Println("----> ask client to create object: ", jsonStr)
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

// create a string representation of the object, with is properties and also he's
// method names
func (self *Client) GetGoObjectInfos(uuid string) (map[string]interface{}, error) {
	obj := GoJavaScript.GetCache().GetObject(uuid)
	var err error
	var infos map[string]interface{}
	if obj != nil {

		infos = make(map[string]interface{}, 0)
		if err != nil {
			return infos, err
		}

		infos["Methods"] = make(map[string]string, 0)
		element := reflect.ValueOf(obj)

		for reflect.TypeOf(element.Interface()).Kind() == reflect.Ptr {
			element = reflect.ValueOf(obj).Elem()
		}

		// Method of JavaScript object are define in map.
		if element.Type().String() != "GoJavaScript.Object" {
			for i := 0; i < element.Addr().NumMethod(); i++ {
				typeMethod := element.Addr().Type().Method(i)
				infos["Methods"].(map[string]string)[typeMethod.Name] = ""
			}
		}

		// Now I will reflect other objects properties...
		for i := 0; i < element.NumField(); i++ {
			valueField := element.Field(i)
			typeField := element.Type().Field(i)

			// Here I will set the property
			// Bust the number of field handler here.
			if valueField.CanInterface() {
				if valueField.IsValid() {
					// So here is the field.
					fieldValue := valueField.Interface()
					fieldName := typeField.Name
					// Dereference the pointer as needed.
					for reflect.TypeOf(fieldValue).Kind() == reflect.Ptr {
						// Here I will access the value and not the pointer...
						fieldValue = valueField.Elem().Interface()
					}

					if reflect.TypeOf(fieldValue).Kind() == reflect.Slice {
						if valueField.Len() > 0 {
							values := make([]interface{}, 0)
							for i := 0; i < valueField.Len(); i++ {
								fieldValue := valueField.Index(i).Interface()
								// If the value is a struct I need to register it to JS
								if reflect.TypeOf(fieldValue).Kind() == reflect.Struct || reflect.TypeOf(fieldValue).Kind() == reflect.Ptr {
									if reflect.TypeOf(reflect.ValueOf(fieldValue).Elem()).Kind() == reflect.Struct {
										// Here the element is a structure so I need to create it representation.
										uuid_ := self.RegisterGoObject(fieldValue, "")
										values = append(values, GoJavaScript.NewObjectRef(uuid_))
									}
								} else {
									values = append(values, fieldValue)
								}
							}
							// if the array dosent contain structure I will set it values...
							infos[fieldName] = values

						}
					} else if reflect.TypeOf(fieldValue).Kind() == reflect.Struct {
						// Set the struct...
						uuid_ := self.RegisterGoObject(fieldValue, "")
						infos[fieldName] = GoJavaScript.NewObjectRef(uuid_)
					} else if reflect.TypeOf(fieldValue).Kind() == reflect.Map {
						if fieldName == "Methods" {
							methods := fieldValue.(map[string]string)
							for name, src := range methods {
								infos["Methods"].(map[string]string)[name] = src
							}
						} else if fieldName == "Properties" {
							// can be recursive?
							properties := fieldValue.(map[string]interface{})
							for name, value := range properties {
								infos[name] = value
							}
						}
					} else {
						// basic type.
						infos[fieldName] = fieldValue
					}
				}
			}
		}

		// Return the list of infos.
		// log.Println("343 --->  ", infos)

		return infos, nil
	}

	return infos, errors.New("object " + uuid + " dosen't exist in the client!")
}

func (self *Client) DeleteGoObject(uuid string) {
	GoJavaScript.GetCache().RemoveObject(uuid)
}

/**
 * Register a go type to be usable as JS type.
 */
func (self *Client) RegisterGoType(value interface{}) {
	// Register local object.
	if reflect.TypeOf(value).String() != "GoJavaScript.Object" {
		Utility.RegisterType(value)
	}
}

/**
 * Localy register a go object.
 */
func (self *Client) RegisterGoObject(obj interface{}, name string) string {
	// Here I will dynamicaly register objet type in the utility cache...
	empty := reflect.New(reflect.TypeOf(obj))
	self.RegisterGoType(empty.Elem().Interface())

	// Random uuid.
	var uuid string
	if len(name) > 0 {
		// In that case the object is in the global scope.
		uuid = Utility.GenerateUUID(name)
	} else {
		// Not a global object.
		uuid = Utility.RandomUUID()
	}

	// Do not recreate already existing object.
	if GoJavaScript.GetCache().GetObject(uuid) != nil {
		return uuid
	}

	// Here I will keep the object in the client cache.
	GoJavaScript.GetCache().SetObject(uuid, obj)

	return uuid
}

/**
 * Register a go function to bu usable in JS (in the global object)
 */
func (self *Client) RegisterGoFunction(name string, fct interface{}) {

	// Keep the function in the local client.
	Utility.RegisterFunction(name, fct)

	// Register the function in the server.

	// Create the action.
	action := GoJavaScript.NewAction("RegisterGoFunction", "")

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
	action := GoJavaScript.NewAction("RegisterJsFunction", "")

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
func (self *Client) CreateObject(name string) GoJavaScript.Object {
	// Create the object.
	obj := GoJavaScript.NewObject(name)

	// Give object the peer so it can register itself with the server.
	obj.SetPeer(self.peer)

	return *obj
}

/**
 * Evaluate sript.
 * The list of global variables to be set before executing the script.
 */
func (self *Client) EvalScript(script string, variables []interface{}) (GoJavaScript.Value, error) {
	// So here I will create the function parameters.
	action := GoJavaScript.NewAction("EvalScript", "")

	action.AppendParam("script", script)
	action.AppendParam("variables", variables)

	// Call the remote action
	action = self.peer.CallRemoteAction(action)

	var err error
	if action.Results[1] != nil {
		if reflect.TypeOf(action.Results[1]).String() == "Error" {
			err = action.Results[1].(error)
		} else if reflect.TypeOf(action.Results[1]).String() == "map[string]interface {}" {
			err = errors.New("Script error")
		}
	}

	// Transform object ref into object as needed.
	value := action.Results[0].(*GoJavaScript.Value)

	value.Export()

	return *value, err
}

/**
 * Call a JS/Go function with a given name and a list of parameter.
 */
func (self *Client) CallFunction(name string, params ...interface{}) (GoJavaScript.Value, error) {
	// So here I will create the function parameters.
	action := GoJavaScript.NewAction("CallFunction", "")
	action.AppendParam("name", name)
	action.AppendParam("params", params)

	// Call the remote action
	action = self.peer.CallRemoteAction(action)
	var err error
	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	value := action.Results[0].(*GoJavaScript.Value)

	// Transform object ref into object as needed.
	value.Export()

	return *value, err
}

/**
 * Set the global variable name.
 */
func (self *Client) SetGlobalVariable(name string, value interface{}) {

	// Replace objects by their reference as needed.
	value = self.objectToRef(value)

	// So here I will create the function parameters.
	action := GoJavaScript.NewAction("SetGlobalVariable", "")
	action.AppendParam("name", name)
	action.AppendParam("value", value)

	// Call the remote action
	action = self.peer.CallRemoteAction(action)
}

/**
 * Get the global variable name.
 */
func (self *Client) GetGlobalVariable(name string) (GoJavaScript.Value, error) {
	action := GoJavaScript.NewAction("GetGlobalVariable", "")

	action.AppendParam("name", name)

	// Call the remote action
	action = self.peer.CallRemoteAction(action)

	var err error
	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	// Transform object ref into object as needed.
	value := action.Results[0].(*GoJavaScript.Value)
	value.Export()

	return *value, err
}

/**
 * Stop JavaScript.
 */
func (self *Client) Stop() bool {
	// Create the action.
	action := GoJavaScript.NewAction("Stop", "")

	// Clear server ressource
	action = self.peer.CallRemoteAction(action)

	// Stop action proecessing
	self.isRunning = false

	// stop it server.
	self.srv.Process.Kill()

	// Close the peers.
	self.peer.Close()

	return true
}
