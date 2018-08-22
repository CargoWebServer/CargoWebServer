package GoJerryScript

import (
	"encoding/gob"
	"log"
	"net"
	"strconv"
)

// Object reference type.
type ObjectRef struct {
	// The uuid of the referenced object.
	UUID string
}

/**
 * Client are use to call remote function over JerryScript js engine.
 */
type Peer struct {

	// The connection with the server side.
	port      int
	address   string
	conn      net.Conn
	isRunning bool

	// communication channel.
	send_chan    chan *Message
	receive_chan chan *Message

	// To stop the peer.
	stop_chan chan bool

	// Actions related channel.

	// Action asked by the peer.
	pending_actions_chan map[string]chan *Action

	// The channel where the action are executed.
	exec_action_chan chan *Action
}

func NewPeer(address string, port int, exec_action_chan chan *Action) *Peer {
	// Transferable objects types over the client/server.
	gob.Register(Message{})
	gob.Register(Action{})
	gob.Register(Param{})
	gob.Register(Variables{})
	gob.Register(Value{})
	gob.Register(Object{})
	gob.Register(ObjectRef{})
	gob.Register([]interface{}{})

	p := new(Peer)
	p.port = port
	p.address = address
	p.isRunning = true

	// Open the channels.
	p.send_chan = make(chan *Message)
	p.receive_chan = make(chan *Message)

	// pending action.
	p.pending_actions_chan = make(map[string]chan *Action, 0)

	// action exec channel, this is given by the channel owner.
	p.exec_action_chan = exec_action_chan

	// Stop the process.
	p.stop_chan = make(chan bool)

	return p
}

// The main processing loop.
func (self *Peer) run() {

	// Listen incomming information.
	go func(self *Peer) {
		for self.isRunning {
			msg := new(Message)
			log.Println(self.conn.LocalAddr().String())
			decoder := gob.NewDecoder(self.conn)
			err := decoder.Decode(msg)
			if err == nil {
				self.receive_chan <- msg
			} else {
				log.Panicln("---> unmarchaling error: ", err)
			}
		}
	}(self)

	// Process
	for {
		select {
		case msg := <-self.receive_chan:
			// So here the message can be a request or a response.
			if msg.Type == 0 {
				// Process request in a separated go routine.
				go func(self *Peer, msg *Message) {
					action := &msg.Remote

					// Open ansew channel
					action.Done = make(chan *Action)

					// In that case I will run the action.
					// Set the action on the channel to be execute by the peer owner.
					self.exec_action_chan <- action

					// Wait for the result.
					action = <-action.Done
					// Here I will create the response and send it back to the client.
					rsp := new(Message)
					rsp.UUID = msg.UUID
					rsp.Type = 1

					// Set back the action
					rsp.Remote = *action

					// Send the message back to the asking peer.
					self.SendMessage(rsp)

				}(self, msg)

			} else {
				go func(self *Peer, msg *Message) {
					// Here I receive a response.
					// I will get the action...
					//log.Println("---> 130 receive response action: ", msg.Remote.UUID, msg.Remote.Name)
					action := <-self.pending_actions_chan[msg.UUID]

					// The response result in the action.
					action.Results = msg.Remote.Results

					// unblock the function call.
					self.pending_actions_chan[msg.UUID] <- action

				}(self, msg)
			}
		case msg := <-self.send_chan:
			go func(self *Peer, msg *Message) {

				// Send the message over the network.
				encoder := gob.NewEncoder(self.conn)
				log.Println("----> send action ", msg.Remote)
				err := encoder.Encode(msg)
				if err != nil {
					log.Println("---> ", msg.Remote.Name)
					for i := 0; i < len(msg.Remote.Results); i++ {
						log.Println("---> ", msg.Remote.Results[i])
					}
					log.Panicln("---> marshaling error: ", err)
				}
			}(self, msg)
		case <-self.stop_chan:
			self.isRunning = false
			break // stop the processin loop.
		}
	}
}

// Send a message over the network
func (self *Peer) SendMessage(msg *Message) {
	// Here I will create the message from the data...
	self.send_chan <- msg
}

// If the peer act as server you must use that function
func (self *Peer) Listen() {
	ln, _ := net.Listen("tcp", ":"+strconv.Itoa(self.port))
	self.conn, _ = ln.Accept()

	// Process message (incoming and outgoing)
	go self.run()
}

func (self *Peer) Close() {
	// Close the connection.
	self.conn.Close()

	// Send stop request.
	self.stop_chan <- true
}

// If the peer act as client you must use that one.
func (self *Peer) Connect(address string, port int) error {

	var err error
	ipv4 := address + ":" + strconv.Itoa(port)
	self.conn, err = net.Dial("tcp", ipv4)

	// Start runing.
	go self.run()

	return err
}

// Run a remote action.
func (self *Peer) CallRemoteAction(action *Action) *Action {

	// So here I will create a pending action channel a wait for the completion
	// of the function before return it.
	self.pending_actions_chan[action.UUID] = make(chan *Action)

	// Create the action request.
	msg := new(Message)
	msg.Type = 0
	msg.UUID = action.UUID

	// Set the action as message data.
	msg.Remote = *action

	// Send the request and wait for it answer.
	self.SendMessage(msg)

	// Put the action in the channel and wait for the response...
	self.pending_actions_chan[action.UUID] <- action

	// Call remote action and wait for it result and set it back in the action.
	action = <-self.pending_actions_chan[action.UUID]
	// Remove it from the map.
	delete(self.pending_actions_chan, action.UUID)

	return action
}
