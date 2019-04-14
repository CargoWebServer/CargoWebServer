// GoJavaScriptClient project GoJavaScriptClient.go
package GoJavaScriptClient

import (
	"log"

	"strconv"
	"time"

	"encoding/json"
	"errors"

	//"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

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
	// Make sure the command exist in the os path...
	log.Println("--> start new GoJavaScriptServer", port, name)
	client.srv = exec.Command("GoJavaScriptServer", strconv.Itoa(port), name)

	err = client.srv.Start()

	if err != nil {
		// In that case I will try to start the GoJavaScriptServer at the save level of the client application.
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err == nil {
			if runtime.GOOS == "windows" {
				client.srv = exec.Command(dir+"\\GoJavaScriptServer.exe", strconv.Itoa(port), name)
			} else {
				client.srv = exec.Command(dir+"/GoJavaScriptServer", strconv.Itoa(port), name)
			}

			err = client.srv.Start()
			if err != nil {
				log.Println("Fail to start GoJavaScriptServer", err)
				return nil
			}
		}

	}

	// Create the client connection, try 5 time and wait 200 millisecond each try.
	for i := 0; i < 100; i++ {
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

	// Set the console.
	client.SetGlobalVariable("console", getConsole())

	return client
}

// Call a go function and return it result.
func (self *Client) callGoFunction(name string, params []interface{}) (interface{}, error) {
	// Replace function string by their pointer.
	for i := 0; i < len(params); i++ {
		if params[i] != nil {
			if reflect.TypeOf(params[i]).Kind() == reflect.String {
				if strings.HasPrefix(params[i].(string), "function ") && strings.HasSuffix(params[i].(string), "() { [native code] }") {
					startIndex := strings.Index(params[i].(string), " ") + 1
					endIndex := strings.Index(params[i].(string), "(")
					name := params[i].(string)[startIndex:endIndex]
					params[i] = Utility.GetFunction(name)
				}
			}
		}
	}

	//log.Println("----> call go function ", name, params)
	results, err := Utility.CallFunction(name, GoJavaScript.RefToObject(params).([]interface{})...)

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

// Action request by the server.
func (self *Client) processActions() {
	for self.isRunning {
		select {
		case action := <-self.exec_action_chan:
			//log.Println("---> action receive: ", action.Name, action.Target)
			go func(a *GoJavaScript.Action) {
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
						if target == nil {
							log.Println("---> fail to retreive target ", a.Target)
						}
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
						log.Println("---> call Method: ", a.Name)
						for i := 0; i < len(params); i++ {
							log.Println("param ", i, params[i], reflect.TypeOf(params[i]).String())
						}
						str, _ := Utility.ToJson(target)
						log.Println("---> target: ", str)
						log.Println("---> error: ", err)
					}
				}
				// Keep go object in client and transfert only reference.
				results = GoJavaScript.ObjectToRef(results)

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
	var value interface{}
	data := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(jsonStr), &data)

	if err == nil {
		// Here Only entity will be update see if other type must be updated...
		if data["TYPENAME"] != nil {
			relfectValue := Utility.MakeInstance(data["TYPENAME"].(string), data, SetEntity)
			value = relfectValue.Interface()
		} else {
			// Here map[string]interface{} will be use.
			err = errors.New("No object are registered in go to back the js object!")
		}
	}

	return value, err
}

// Set object values.
func (self *Client) SetObjectValues(uuid string, jsonStr string) {
	var obj interface{}
	var err error
	obj, err = self.CreateGoObject(jsonStr)
	if err == nil {
		// Set back the object with it new values.
		GoJavaScript.GetCache().SetObject(uuid, obj)
	}
}

// create a string representation of the object, with is properties and also he's
// method names
func (self *Client) GetGoObjectInfos(uuid string) (map[string]interface{}, error) {
	obj := GoJavaScript.GetCache().GetObject(uuid)
	var err error
	var infos map[string]interface{}

	if obj != nil {
		infos = make(map[string]interface{}, 0)

		// define object infos...
		infos["__object_infos__"] = ""
		infos["Methods"] = make(map[string]interface{}, 0)
		infos["uuid_"] = uuid

		// in case the object is a generic map with no method...
		if reflect.TypeOf(obj).String() == "map[string]interface {}" {
			for name, value := range obj.(map[string]interface{}) {
				infos[name] = value
			}
			return infos, nil
		}

		if err != nil {
			return infos, err
		}

		element := reflect.ValueOf(obj)

		for reflect.TypeOf(element.Interface()).Kind() == reflect.Ptr {
			element = reflect.ValueOf(obj).Elem()
		}

		// Method of JavaScript object are define in map.
		if element.IsValid() {
			if element.Type().String() != "GoJavaScript.Object" {
				for i := 0; i < element.Addr().NumMethod(); i++ {
					typeMethod := element.Addr().Type().Method(i)
					infos["Methods"].(map[string]interface{})[typeMethod.Name] = ""
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
										// Here the element is a structure so I need to create it representation.
										uuid_ := GoJavaScript.RegisterGoObject(fieldValue, "")
										values = append(values, GoJavaScript.NewObjectRef(uuid_))
									} else {
										values = append(values, fieldValue)
									}
								}

								// if the array dosent contain structure I will set it values...
								infos[fieldName] = values

							}
						} else if reflect.TypeOf(fieldValue).Kind() == reflect.Struct || reflect.TypeOf(fieldValue).Kind() == reflect.Ptr {
							// Set the struct...
							uuid_ := GoJavaScript.RegisterGoObject(fieldValue, "")
							infos[fieldName] = GoJavaScript.NewObjectRef(uuid_)
						} else if reflect.TypeOf(fieldValue).Kind() == reflect.Map {
							if fieldName == "Methods" {
								methods := fieldValue.(map[string]interface{})
								for name, src := range methods {
									infos["Methods"].(map[string]interface{})[name] = src // src can by byte code.
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
				} else {
					//log.Println("---> field cant interface ", typeField.Name)
				}
			}
		}

		// Return the list of infos.
		return infos, nil
	}

	return infos, errors.New("object " + uuid + " dosen't exist in the client!")
}

func (self *Client) DeleteGoObject(uuid string) {
	GoJavaScript.GetCache().RemoveObject(uuid)
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

// Set Js object property.
func (self *Client) SetObjectProperty(uuid string, name string, value interface{}) {
	obj := GoJavaScript.GetCache().GetObject(uuid).(GoJavaScript.Object)

	// set the object property.
	obj.Set(name, value)
}

func (self *Client) SetObjectJsMethod(uuid string, name string, value interface{}) {
	obj := GoJavaScript.GetCache().GetObject(uuid).(GoJavaScript.Object)
	// set the object property.
	obj.Methods[name] = value
}

/**
 * Evaluate sript.
 * The list of global variables to be set before executing the script.
 */
func (self *Client) EvalScript(script string, variables []interface{}) (interface{}, error) {

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
	value := GoJavaScript.RefToObject(action.Results[0])

	return value, err
}

/**
 * Call a JS/Go function with a given name and a list of parameter.
 */
func (self *Client) CallFunction(name string, params ...interface{}) (interface{}, error) {

	// So here I will create the function parameters.
	action := GoJavaScript.NewAction("CallFunction", "")
	action.AppendParam("name", name)
	action.AppendParam("params", GoJavaScript.ObjectToRef(params))

	// Call the remote action
	action = self.peer.CallRemoteAction(action)
	var err error
	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	value := GoJavaScript.RefToObject(action.Results[0])

	return value, err
}

/**
 * Set the global variable name.
 */
func (self *Client) SetGlobalVariable(name string, value interface{}) {

	// Replace objects by their reference as needed.
	value = GoJavaScript.ObjectToRef(value)

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
func (self *Client) GetGlobalVariable(name string) (interface{}, error) {
	action := GoJavaScript.NewAction("GetGlobalVariable", "")

	action.AppendParam("name", name)

	// Call the remote action
	action = self.peer.CallRemoteAction(action)

	var err error
	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	// Transform object ref into object as needed.
	value := GoJavaScript.RefToObject(action.Results[0])

	return value, err
}

/**
 * Use by the server to test if the client is alive.
 */
func (self *Client) Ping() string {
	return "Pong"
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
