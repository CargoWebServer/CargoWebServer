package GoJerryScript

import (
	"code.myceliUs.com/Utility"
)

// Message are use to send information between peers.
type Message struct {
	// The unique identifier of the message.
	UUID string

	// Contain the message data (Json string)
	Data []byte

	Type int // Can be 0: Request, 1:Response
}

// Create a message...
func NewMessage(Type int, Data []byte) *Message {
	msg := new(Message)
	msg.UUID = Utility.RandomUUID()
	msg.Data = Data
	return msg
}
