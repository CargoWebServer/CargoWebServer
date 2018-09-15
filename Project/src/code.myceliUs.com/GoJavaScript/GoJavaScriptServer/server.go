package main

import (
	"log"

	"code.myceliUs.com/GoJavaScript"
	"code.myceliUs.com/GoJavaScript/GoChakra"
	"code.myceliUs.com/GoJavaScript/GoDuktape"
	"code.myceliUs.com/GoJavaScript/GoJerryScript"
)

// The server will redirect client command to JavaScript.
type Server struct {
	isRunning bool
	peer      *GoJavaScript.Peer

	// The underlying engine, can be:
	// - JerryScript
	engine GoJavaScript.Engine

	// execute remote action.
	exec_action_chan chan *GoJavaScript.Action

	// call remote actions from the engine.
	call_remote_actions_chan chan *GoJavaScript.Action
}

func NewServer(address string, port int, name string) *Server {

	// I will create a new server and start listen for incomming request.
	server := new(Server)
	server.isRunning = true
	// open the action channel.
	server.exec_action_chan = make(chan *GoJavaScript.Action)

	// open remote action call.
	server.call_remote_actions_chan = make(chan *GoJavaScript.Action, 0)

	// Global variable use in server side.
	GoJavaScript.Call_remote_actions_chan = server.call_remote_actions_chan

	// Create the peer.
	server.peer = GoJavaScript.NewPeer(address, port, server.exec_action_chan)

	// The underlying engine.

	// The engine.
	if name == "jerryscript" {
		server.engine = new(GoJerryScript.Engine)
	} else if name == "duktape" {
		server.engine = new(GoDuktape.Engine)
	} else if name == "chakracore" {
		server.engine = new(GoChakra.Engine)
	}

	// Start the engine
	server.engine.Start(port)

	// Start listen
	go server.peer.Listen()

	// Return the newly created server.
	return server
}

// In case the server made call to it client.
func (self *Server) processRemoteActions() {
	for self.isRunning {
		select {
		case action := <-self.call_remote_actions_chan:
			go func(a *GoJavaScript.Action, s *Server) {
				// Call remote action
				a = s.peer.CallRemoteAction(a)
				// Set back the action on the channel.
				action.GetDone() <- a
			}(action, self)
		}

	}
}

// Run requested actions.
func (self *Server) processActions() {

	// Process remote action in it own goroutine.
	go self.processRemoteActions()
	for self.isRunning {
		select {
		case action := <-self.exec_action_chan:
			// Here the action will be execute in a non-blocking way so
			// other exec action will be possible.
			go func() {
				log.Println("---> call action: ", action.Name)
				if action.Name == "RegisterJsFunction" {
					action.AppendResults(self.engine.RegisterJsFunction(action.Params[0].Value.(string), action.Params[1].Value.(string)))
				} else if action.Name == "EvalScript" {
					// So here I will call the function and return it value.
					if action.Params[1].Value != nil {
						action.AppendResults(self.engine.EvalScript(action.Params[0].Value.(string), action.Params[1].Value.([]interface{})))
					} else {
						action.AppendResults(self.engine.EvalScript(action.Params[0].Value.(string), []interface{}{}))
					}
				} else if action.Name == "CallFunction" {
					// So here I will call the function and return it value.
					if action.Params[1].Value != nil {
						action.AppendResults(self.engine.CallFunction(action.Params[0].Value.(string), action.Params[1].Value.([]interface{})))
					} else {
						action.AppendResults(self.engine.CallFunction(action.Params[0].Value.(string), []interface{}{}))
					}
				} else if action.Name == "RegisterGoFunction" {
					self.engine.RegisterGoFunction(action.Params[0].Value.(string))
				} else if action.Name == "CreateObject" {
					self.engine.CreateObject(action.Params[0].Value.(string), action.Params[1].Value.(string))
				} else if action.Name == "CallObjectMethod" {
					action.AppendResults(self.engine.CallObjectMethod(action.Params[0].Value.(string), action.Params[1].Value.(string), action.Params[2].Value.([]interface{})...))
				} else if action.Name == "SetGlobalVariable" {
					self.engine.SetGlobalVariable(action.Params[0].Value.(string), action.Params[1].Value)
				} else if action.Name == "GetGlobalVariable" {
					action.AppendResults(self.engine.GetGlobalVariable(action.Params[0].Value.(string)))
				} else if action.Name == "Stop" {
					log.Println("--> Stop JavaScript exec!")
					// Send back the result to client.
					self.isRunning = false
				}
				action.GetDone() <- action
			}()

		}
	}
	self.peer.Close()
	log.Println("---> server stop processing incomming messages!")
}
