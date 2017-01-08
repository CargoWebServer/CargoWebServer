package Server

import (
	"log"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/JS"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
)

type Hub struct {
	// Registered connections.
	connections map[string]connection

	// Inbound messages from the connections.
	// At this stage the message are just binairy data.
	receivedMsg chan *message

	// Register requests from the connections.
	register chan connection

	// Unregister requests from connections.
	unregister chan connection

	// Ticker to ping the connection...
	ticker *time.Ticker

	// stop processing whent that variable are set to true...
	abortedByEnvironment chan bool
}

func NewHub() *Hub {

	h := new(Hub)
	h.abortedByEnvironment = make(chan bool)
	h.connections = make(map[string]connection)
	h.receivedMsg = make(chan *message)
	h.register = make(chan connection)
	h.unregister = make(chan connection)

	// So here I will send empty message to keep socket alive
	// and clear the session if the connection is close...
	h.ticker = time.NewTicker(time.Millisecond * 2000)

	go func() {
		for t := range h.ticker.C {
			for _, conn := range h.connections {
				if conn.IsOpen() {
					id := Utility.RandomUUID()
					method := "Ping"
					params := make([]*MessageData, 0)

					to := make([]connection, 1)
					to[0] = conn

					// The success callback.
					successCallback := func(msg *message) {
						//log.Println("This is the end!!!")
						/* nothing todo here */
					}

					ping, err := NewRequestMessage(id, method, params, to, successCallback)
					if err != nil {
						log.Println(err, " at time ", t)
					}

					GetServer().GetProcessor().m_pendingRequestChannel <- ping

				} else {
					GetServer().GetSessionManager().removeClosedSession()
				}
			}
		}
	}()

	return h
}

func (this *Hub) run() {
	for {
		select {
		case c := <-this.register:
			this.connections[c.GetUuid()] = c
			// Set a new js server object...
			vm := JS.GetJsRuntimeManager().GetVm(c.GetUuid())
			vm.Set("server", GetServer())
		case c := <-this.unregister:
			delete(this.connections, c.GetUuid())
			c.Close()
			GetServer().GetEventManager().removeClosedListener()
			GetServer().GetSessionManager().removeClosedSession()
		case msg := <-this.receivedMsg:
			GetServer().GetProcessor().m_incomingChannel <- msg
		case done := <-this.abortedByEnvironment:
			if done {
				return
			}
		}
	}

	this.ticker.Stop()
}
