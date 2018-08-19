// GoJerryScriptClient project GoJerryScriptClient.go
package GoJerryScriptClient

import (
	"log"
	"strconv"

	"time"

	"os/exec"
	"reflect"

	"code.myceliUs.com/GoJerryScript"
	"code.myceliUs.com/Utility"
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

	// Here I will start the external server process.
	client.srv = exec.Command("/home/dave/Documents/CargoWebServer/Project/src/code.myceliUs.com/GoJerryScript/GoJerryScriptServer/GoJerryScriptServer", strconv.Itoa(port))
	err := client.srv.Start()

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
			log.Println("---> client action: ", action.Name)
			var target interface{}
			isJsObject := true
			if len(action.Target) > 0 {
				target = GoJerryScript.GetCache().GetObject(action.Target)
				isJsObject = reflect.TypeOf(target).String() == "*GoJerryScript.Object"
			}

			params := make([]interface{}, 0)
			for i := 0; i < len(action.Params); i++ {
				// So here I will append the value to parameters.
				params = append(params, action.Params[i].Value)
			}

			// I will call the function.
			if isJsObject {
				action.AppendResults(self.callGoFunction(action.Name, params...))
			} else {
				// Call method on the object.
				log.Println("---> call method on object, implement me!!!")
			}

			// send back the response.
			self.exec_action_chan <- action
		}
	}
}

////////////////////////////-- Api --////////////////////////////

/**
 * Register a go type to be usable as JS type.
 */
func (self *Client) RegisterGoType(value interface{}) {
	// Register local object.
	Utility.RegisterType(value)
}

/**
 * Register a go function to bu usable in JS (in the global object)
 */
func (self *Client) RegisterGoFunction(name string, fct interface{}) {

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
func (self *Client) CreatObject(name string) *GoJerryScript.Object {
	// Create the object.
	obj := GoJerryScript.NewObject(name)

	// Give object the peer so it can register itself with the server.
	obj.SetPeer(self.peer)

	return obj
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

func (self *Client) CallFunction(name string, params []interface{}) (GoJerryScript.Value, error) {
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
