package GoJavaScript

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"code.myceliUs.com/Utility"
)

func NewObjectRef(uuid string) *ObjectRef {
	ref := new(ObjectRef)
	ref.UUID = uuid
	ref.TYPENAME = "GoJavaScript.ObjectRef"
	return ref
}

var (
	// Callback function used by dynamic type, it's call when an entity is set.
	// Can be use to store dynamic type in a cache.
	SetEntity func(interface{}) = func(val interface{}) {
		/** nothing todo here... **/
	}
)

/**
 * Client are use to call remote function over javascript engine.
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
	pending_actions_chan *sync.Map // map[string]chan *Action

	// The channel where the action are executed.
	exec_action_chan chan *Action
}

var (
	logChannel = make(chan string)
)

func Log(infos ...interface{}) {

	// also display in the command prompt.
	logChannel <- fmt.Sprintln(infos)
}

func NewPeer(address string, port int, exec_action_chan chan *Action) *Peer {
	// Transferable objects types over the client/server.
	Utility.RegisterType((*Message)(nil))
	Utility.RegisterType((*Action)(nil))
	Utility.RegisterType((*Param)(nil))
	Utility.RegisterType((*Variable)(nil))
	Utility.RegisterType((*Object)(nil))
	Utility.RegisterType((*ObjectRef)(nil))

	p := new(Peer)
	p.port = port
	p.address = address
	p.isRunning = true

	// Open the channels.
	p.send_chan = make(chan *Message)
	p.receive_chan = make(chan *Message)

	// pending action.
	p.pending_actions_chan = new(sync.Map) // make(map[string]chan *Action, 0)

	// action exec channel, this is given by the channel owner.return
	p.exec_action_chan = exec_action_chan

	// Stop the process.
	p.stop_chan = make(chan bool)

	// The log processing information.
	go func(p *Peer) {
		for {
			select {
			case msg := <-logChannel:
				// Open the log file.
				f, err := os.OpenFile("./GoJavaScript_"+strconv.Itoa(p.port)+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
				if err == nil {
					logger := log.New(f, "", log.LstdFlags)
					logger.Println(msg)
					// set the message.
					f.Close()
				}
			}
		}
	}(p)

	return p
}

// The main processing loop.
func (self *Peer) run() {

	// Listen incomming information.
	go func() {
		for self.isRunning {
			data := make(map[string]interface{})
			decoder := json.NewDecoder(self.conn)
			if decoder != nil {
				err := decoder.Decode(&data)
				if err == nil {
					// Here I will create a message from the generic map of interface
					msg, err := Utility.InitializeStructure(data, SetEntity)
					if err != nil {
						return
					}
					self.receive_chan <- msg.Interface().(*Message)
				} else {
					if err.Error() == "EOF" {
						Log("136 Peer.go --> Connection close normaly for peer ", self.address, self.port)
					}
					return
				}
			} else {
				Log("Peer.go line 107 fail to create decoder tcp connection!")
				return
			}
		}
	}()

	// Process
	for self.isRunning {
		select {
		case msg := <-self.receive_chan:
			// So here the message can be a request or a response.
			if msg.Type == Request {
				// Process request in a separated go routine.
				go func(self *Peer, msg *Message) {
					// Get the action
					action := msg.Remote

					// Open answer channel
					action.SetDone()

					// In that case I will run the action.
					// Set the action on the channel to be execute by the peer owner.
					self.exec_action_chan <- action

					// Wait for the result.
					action = <-action.GetDone()

					// Here I will create the response and send it back to the client.
					rsp := NewMessage(Response, action)

					// Send the message back to the asking peer.
					self.SendMessage(rsp)

				}(self, msg)

			} else {
				go func(self *Peer, msg *Message) {
					// Here I receive a response.
					// I will get the action...
					channel, _ := self.pending_actions_chan.Load(msg.UUID)
					action := <-channel.(chan *Action)

					// The response result in the action.
					action.Results = msg.Remote.Results

					// unblock the function call.
					channel.(chan *Action) <- action

				}(self, msg)
			}
		case msg := <-self.send_chan:
			go func(self *Peer, msg *Message) {
				// Send the message over the network.
				encoder := json.NewEncoder(self.conn)
				err := encoder.Encode(msg)
				if err != nil {
					Log("---> marshaling error: ", err, os.Args)
					os.Exit(0)
				}
			}(self, msg)
		case <-self.stop_chan:
			Log("---> close peer ", self.address, self.port)
			self.conn.Close()
			self.isRunning = false

			//erase map
			self.pending_actions_chan.Range(func(key interface{}, value interface{}) bool {
				// clear the map.
				action := value.(*Action)
				action.AppendResults(errors.New("Action cancel by peer " + self.address))
				action.GetDone() <- action
				self.pending_actions_chan.Delete(key)
				return true
			})

			// Close the channels
			close(self.send_chan)
			close(self.receive_chan)
			close(self.exec_action_chan)

			break
		}
	}
}

func (self *Peer) SetPort(port int) {
	self.port = port
}

func (self *Peer) Ping() {
	// Create the action.
	ping := NewAction("Ping", "")
	ping.Target = "Client"

	// Append the name parameter.
	self.CallRemoteAction(ping)
}

// Send a message over the network
func (self *Peer) SendMessage(msg *Message) {
	// Here I will create the message from the data...
	self.send_chan <- msg
}

// If the peer act as server you must use that function
func (self *Peer) Listen() {
	ln, _ := net.Listen("tcp", ":"+strconv.Itoa(self.port))

	// Set the port from the actual connection.
	self.port = ln.Addr().(*net.TCPAddr).Port
	Log("248 connection open from server at port ---> ", self.port)

	// print the value to the console.
	fmt.Println(self.port)

	self.conn, _ = ln.Accept() // only one connection per server here...

	// Process message (incoming and outgoing)
	go self.run()

	// test if the connection is alive.
	go func(self *Peer) {
		// Call ping a interval of 5 second...
		for ; true; <-time.Tick(2 * time.Second) {
			self.Ping()
		}

	}(self)
}

func (self *Peer) Close() {
	// Send stop request.
	self.stop_chan <- true
}

// If the peer act as client you must use that one.
func (self *Peer) Connect(address string, port int) error {

	var err error
	ipv4 := address + ":" + strconv.Itoa(port)
	self.conn, err = net.Dial("tcp", ipv4)

	if err != nil {
		Log("--> connection error: ", err)
		return err
	}

	// Start runing.
	go self.run()

	return nil
}

// Run a remote action.
func (self *Peer) CallRemoteAction(action *Action) *Action {

	// So here I will create a pending action channel a wait for the completion
	// of the function before return it.
	channel := make(chan *Action)
	self.pending_actions_chan.Store(action.UUID, channel)

	// Create the action request.
	msg := NewMessage(Request, action)

	// Send the request and wait for it answer.
	//
	go func() {
		// ** The message is sent in it own goroutine so the action is
		// append faster in the pending action map and chance of response arrive
		// before the action is waiting are null.
		self.SendMessage(msg)
	}()

	// Put the action in the channel and wait for the response...
	channel <- action

	// Call remote action and wait for it result and set it back in the action.
	action = <-channel

	// Remove it from the map.
	self.pending_actions_chan.Delete(action.UUID)

	return action
}
