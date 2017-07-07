package Server

import (
	"log"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/JS"
	"code.myceliUs.com/Utility"
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
			for id, conn := range h.connections {
				if conn != nil {
					if conn.IsOpen() {
						id := Utility.RandomUUID()
						method := "Ping"
						params := make([]*MessageData, 0)
						to := make([]connection, 1)
						to[0] = conn
						successCallback := func(rspMsg *message, caller interface{}) {
							//log.Println("success!!!")
						}

						errorCallback := func(rspMsg *message, caller interface{}) {
							//log.Println("error!!!")
						}

						ping, err := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback)

						if err != nil {
							log.Println(err, " at time ", t)
							conn.Close() // Here I will close the connection.
						} else {
							GetServer().GetProcessor().m_pendingRequestChannel <- ping
						}

					} else {
						GetServer().GetSessionManager().removeClosedSession()
					}
				} else {
					delete(h.connections, id)
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
			log.Println("------>Hub append connection ", c.GetUuid())
			this.connections[c.GetUuid()] = c
			// initialyse js interpreter for the new connection.
			if c.GetPort() == GetServer().GetConfigurationManager().GetServerPort() {

				vm := JS.GetJsRuntimeManager().CreateVm(c.GetUuid())

				// put server object in it context.
				vm.Set("server", GetServer())

				// Init scripts
				JS.GetJsRuntimeManager().InitScripts(c.GetUuid())
			}

		case c := <-this.unregister:
			log.Println("----> remove connection ", c.GetAddrStr(), c.GetPort())
			delete(this.connections, c.GetUuid())
			GetServer().GetEventManager().removeClosedListener()
			GetServer().GetSessionManager().removeClosedSession()
			GetServer().onClose(c.GetUuid())

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
