package Server

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"code.myceliUs.com/Utility"
	"golang.org/x/net/websocket"
)

////////////////////////////////////////////////////////////////////////////////
//									WebSocket
////////////////////////////////////////////////////////////////////////////////
/**
 * The web socket connection...
 */
type WebSocketConnection struct {
	// The websocket connection.
	m_socket *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// The socket state.
	m_isOpen bool

	// number of try before giving up
	m_try int

	// The uuid of the connection
	m_uuid string
}

func NewWebSocketConnection() *WebSocketConnection {
	var conn = new(WebSocketConnection)
	conn.send = make(chan []byte /*, connection_channel_size*/)
	conn.m_uuid = Utility.RandomUUID()
	conn.m_isOpen = false
	return conn
}

func (c *WebSocketConnection) GetHostname() string {
	var address string
	if c.m_socket.IsServerConn() {
		address = c.m_socket.Request().RemoteAddr
	} else {
		address = c.m_socket.RemoteAddr().String()[5:] // ws:// (5 char to remove)
	}

	address = address[:strings.Index(address, ":")] // Remove the port...
	return address
}

func (c *WebSocketConnection) GetPort() int {
	address := c.m_socket.RemoteAddr().String()
	startIndex := strings.LastIndex(address, ":") + 1
	endIndex := strings.LastIndex(address, "/ws")
	var port int
	if endIndex == -1 {
		port, _ = strconv.Atoi(address[startIndex:])
	} else {
		port, _ = strconv.Atoi(address[startIndex:endIndex])
	}
	return port
}

func (c *WebSocketConnection) GetUuid() string {
	return c.m_uuid
}

func (c *WebSocketConnection) Open(host string, port int) (err error) {

	// Open the socket...
	url := "http://" + host + ":" + strconv.Itoa(port) + "/ws"
	origin := "ws://" + host + ":" + strconv.Itoa(port) + "/ws"
	c.m_socket, err = websocket.Dial(origin, "", url)

	if err != nil && c.m_try < 30 {
		time.Sleep(500 * time.Millisecond)
		c.m_try += 1
		return c.Open(host, port)
	} else if c.m_try == 3 {
		return errors.New("fail to connect with " + origin)
	} else if c.m_socket != nil {
		c.m_isOpen = true
		GetServer().getHub().register <- c
		// Start reading and writing loop's
		go c.Writer()
		go c.Reader()
	}

	return err
}

func (c *WebSocketConnection) Close() {
	if c.m_isOpen {
		c.m_isOpen = false
		c.m_socket.Close() // Close the socket..
		GetServer().getHub().unregister <- c
		GetServer().removeAllOpenSubConnections(c.GetUuid())
	}
}

/**
 * The connection state...
 */
func (c *WebSocketConnection) IsOpen() bool {
	return c.m_isOpen
}

func (c *WebSocketConnection) Send(data []byte) {
	c.send <- data
}

func (c *WebSocketConnection) Reader() {
	for c.m_isOpen == true {
		var in []byte
		if err := websocket.Message.Receive(c.m_socket, &in); err != nil {
			c.Close()
			break // Exit the reading loop.
		}
		msg, err := NewMessageFromData(in, c)
		if err == nil {
			GetServer().getHub().receivedMsg <- msg
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func (c *WebSocketConnection) Writer() {
	for c.m_isOpen == true {
		for message := range c.send {
			// I will get the message here...
			websocket.Message.Send(c.m_socket, message)
		}
		time.Sleep(1 * time.Millisecond)
	}
}

// The web socket handler function...
func HttpHandler(ws *websocket.Conn) {

	// Here I will create the new connection...
	c := NewWebSocketConnection()
	c.m_socket = ws
	c.m_isOpen = true
	c.send = make(chan []byte)

	// Register the connection with the hub.
	GetServer().getHub().register <- c

	defer func(c *WebSocketConnection) {
		//  I will remove all sub-connection associated with the connection
		GetServer().removeAllOpenSubConnections(c.GetUuid())
		c.Close()
		LogInfo("--> connection ", c.GetUuid(), "is close!")
	}(c)

	// Start the writing loop...
	go c.Writer()

	// we stay it stay in reader loop until the connection is close.
	c.Reader()
}
