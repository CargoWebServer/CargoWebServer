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
			// Call remote action
			action = self.peer.CallRemoteAction(action)

			// Set back the action on the channel.
			self.call_remote_actions_chan <- action
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
			// Run action here.
			if action.Name == "AppendJsFunction" {
				action.AppendResults(self.engine.AppendJsFunction(action.Params[0].Value.(string), action.Params[1].Value.([]string), action.Params[2].Value.(string)))
				self.exec_action_chan <- action

			} else if action.Name == "EvalScript" {
				script := action.Params[0].Value.(string)
				variables := action.Params[1].Value.(GoJerryScript.Variables)
				// So here I will call the function and return it value.
				action.AppendResults(self.engine.EvalScript(script, variables))
				self.exec_action_chan <- action
			} else if action.Name == "RegisterGoFunction" {
				self.engine.RegisterGoFunction(action.Params[0].Value.(string))
				self.exec_action_chan <- action
			} else if action.Name == "Stop" {
				log.Println("--> Stop JerryScript!")
				// TODO fix it.
				// Stop the server.
				/*self.isRunning = false
				self.engine.Clear()

				// Send back the result to client.
				self.exec_action_chan <- action

				// Close the connection.
				self.peer.Close()

				// return.
				break*/
			}
		}
	}
}
