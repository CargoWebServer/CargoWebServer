package main

import (
	"log"

	"code.myceliUs.com/GoJerryScript"
)

// The server will redirect client command to JerryScript.
type Server struct {
	isRunning bool
	peer      *GoJerryScript.Peer
	engine    *GoJerryScript.Engine

	// execute remote action.
	exec_action_chan chan *GoJerryScript.Action

	// call remote actions from the engine.
	call_remote_actions_chan chan *GoJerryScript.Action
}

func NewServer(address string, port int) *Server {

	// I will create a new server and start listen for incomming request.
	server := new(Server)
	server.isRunning = true
	// open the action channel.
	server.exec_action_chan = make(chan *GoJerryScript.Action)

	// open remote action call.
	server.call_remote_actions_chan = make(chan *GoJerryScript.Action, 0)

	// Global variable use in server side.
	GoJerryScript.Call_remote_actions_chan = server.call_remote_actions_chan

	// Create the peer.
	server.peer = GoJerryScript.NewPeer(address, port, server.exec_action_chan)

	// The underlying engine.
	server.engine = GoJerryScript.NewEngine(port, GoJerryScript.JERRY_INIT_EMPTY)

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
			go func(a *GoJerryScript.Action, s *Server) {
				// Call remote action
				a = s.peer.CallRemoteAction(a)
				// Set back the action on the channel.
				a.GetDone() <- a
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
			go func(a *GoJerryScript.Action, s *Server) {
				if a.Name == "RegisterJsFunction" {
					a.AppendResults(s.engine.RegisterJsFunction(a.Params[0].Value.(string), a.Params[1].Value.(string)))
				} else if a.Name == "EvalScript" {
					// So here I will call the function and return it value.
					if a.Params[1].Value != nil {
						a.AppendResults(s.engine.EvalScript(a.Params[0].Value.(string), a.Params[1].Value.([]interface{})))
					} else {
						a.AppendResults(s.engine.EvalScript(a.Params[0].Value.(string), []interface{}{}))
					}
				} else if a.Name == "CallFunction" {
					// So here I will call the function and return it value.
					if a.Params[1].Value != nil {
						a.AppendResults(s.engine.CallFunction(a.Params[0].Value.(string), a.Params[1].Value.([]interface{})))
					} else {
						a.AppendResults(s.engine.CallFunction(a.Params[0].Value.(string), []interface{}{}))
					}
				} else if a.Name == "RegisterGoFunction" {
					s.engine.RegisterGoFunction(a.Params[0].Value.(string))
				} else if a.Name == "CreateObject" {
					s.engine.CreateObject(a.Params[0].Value.(string), a.Params[1].Value.(string))
				} else if a.Name == "SetObjectProperty" {
					s.engine.SetObjectProperty(a.Params[0].Value.(string), a.Params[1].Value.(string), a.Params[2].Value)
				} else if a.Name == "SetGoObjectMethod" {
					s.engine.SetGoObjectMethod(a.Params[0].Value.(string), a.Params[1].Value.(string))
				} else if a.Name == "SetJsObjectMethod" {
					s.engine.SetJsObjectMethod(a.Params[0].Value.(string), a.Params[1].Value.(string), a.Params[2].Value.(string))
				} else if a.Name == "GetObjectProperty" {
					a.AppendResults(s.engine.GetObjectProperty(a.Params[0].Value.(string), a.Params[1].Value.(string)))
				} else if a.Name == "CallObjectMethod" {
					a.AppendResults(s.engine.CallObjectMethod(a.Params[0].Value.(string), a.Params[1].Value.(string), a.Params[2].Value.([]interface{})...))
				} else if a.Name == "CreateObjectArray" {
					a.AppendResults(s.engine.CreateObjectArray(a.Params[0].Value.(string), a.Params[1].Value.(string), uint32(a.Params[2].Value.(float64))))
				} else if a.Name == "SetObjectPropertyAtIndex" {
					s.engine.SetObjectPropertyAtIndex(a.Params[0].Value.(string), a.Params[1].Value.(string), uint32(a.Params[2].Value.(float64)), a.Params[3].Value)
				} else if a.Name == "GetObjectPropertyAtIndex" {
					a.AppendResults(s.engine.GetObjectPropertyAtIndex(a.Params[0].Value.(string), a.Params[1].Value.(string), uint32(a.Params[2].Value.(float64))))
				} else if a.Name == "SetGlobalVariable" {
					s.engine.SetGlobalVariable(a.Params[0].Value.(string), a.Params[1].Value)
				} else if a.Name == "GetGlobalVariable" {
					a.AppendResults(s.engine.GetGlobalVariable(a.Params[0].Value.(string)))
				} else if a.Name == "Stop" {
					log.Println("--> Stop JerryScript!")
					// TODO fix it.
					// Stop the server.
					/*self.isRunning = false
					self.engine.Clear()

					// Send back the result to client.
					self.exec_action_chan <- a

					// Close the connection.
					self.peer.Close()

					// return.
					break*/
				}

				a.GetDone() <- a
			}(action, self)
		}
	}
}
