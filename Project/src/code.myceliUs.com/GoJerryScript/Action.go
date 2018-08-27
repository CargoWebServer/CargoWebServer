package GoJerryScript

import "reflect"
import "code.myceliUs.com/Utility"
import "log"

type MessageType int

const (
	Request MessageType = 1 + iota
	Response
)

type Message struct {

	// The typename.
	TYPENAME string

	// Can be 0 request or 1 response
	Type MessageType

	// Same as it action.
	UUID string

	// The remote action to execute.
	Remote *Action
}

func NewMessage(messageType MessageType, action *Action) *Message {
	msg := new(Message)
	msg.TYPENAME = "GoJerryScript.Message"
	msg.Remote = action
	msg.Type = messageType
	msg.UUID = action.UUID
	return msg
}

type Param struct {
	// The typename.
	TYPENAME string

	// The name of the parameter if any
	Name string

	// It data type
	Type string

	// It contain value.
	Value interface{}
}

// Message can contain action to be done by both client and server.
type Action struct {
	// The typename.
	TYPENAME string

	// The same uuid as the message.
	UUID string

	// If the action must be
	Target string

	// The name of the action.
	Name string

	// The list of action parameters.
	Params []*Param

	// The action results.
	Results []interface{}

	// That channel is use to return the action
	// when execution is done, so it no
	done chan *Action
}

// Create a new action.
func NewAction(name string, target string) *Action {
	a := new(Action)
	a.TYPENAME = "GoJerryScript.Action"
	a.Name = name
	a.Target = target
	a.UUID = Utility.RandomUUID()
	return a
}

func (self *Action) SetDone() {
	self.done = make(chan *Action)
}

// Use to get the action result.
func (self *Action) GetDone() chan *Action {
	return self.done
}

// Append a new parameter to an action.
func (self *Action) AppendParam(name string, value interface{}) {
	// Create a new paremeter.
	param := new(Param)
	param.TYPENAME = "GoJerryScript.Param"
	param.Name = name

	// Here I will use the go reflection to get the type of
	// the value...
	param.Type = reflect.TypeOf(value).String()

	// Set the parameter value.
	param.Value = value

	log.Println("-----> param: ", name, " value ", value)

	if self.Params == nil {
		self.Params = make([]*Param, 0)
	}

	self.Params = append(self.Params, param)
}

// Append the action results.
func (self *Action) AppendResults(results ...interface{}) {
	for i := 0; i < len(results); i++ {
		log.Println("---> ", results[i])
	}
	self.Results = make([]interface{}, 0)
	self.Results = append(self.Results, results...)
}
