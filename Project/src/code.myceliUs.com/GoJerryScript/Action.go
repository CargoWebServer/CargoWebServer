package GoJerryScript

import "reflect"

type Message struct {
	// Can be 0 request or 1 response
	Type int

	// Same as it action.
	UUID string

	// The remote action to execute.
	Remote Action
}

type Param struct {
	// The name of the parameter if any
	Name string

	// It data type
	Type string

	// It contain value.
	Value interface{}
}

// Message can contain action to be done by both client and server.
type Action struct {

	// The same uuid as the message.
	UUID string

	// If the action must be
	Target string

	// The name of the action.
	Name string

	// The list of action parameters.
	Params []Param

	// The action results.
	Results []interface{}

	// That channel is use to return the action
	// when execution is done, so it no
	Done chan *Action
}

// Append a new parameter to an action.
// TODO test if the parameter is an object reference in that case send
// object info and not the object itself. To the same for results.
func (self *Action) AppendParam(name string, value interface{}) {

	var param Param
	param.Name = name

	// Here I will use the go reflection to get the type of
	// the value...
	param.Type = reflect.TypeOf(value).String()

	// Set the parameter value.
	param.Value = value

	if self.Params == nil {
		self.Params = make([]Param, 0)
	}
	self.Params = append(self.Params, param)
}

// Append the action results.
// TODO test if the parameter is an object reference in that case send
// object info and not the object itself. To the same for results.
func (self *Action) AppendResults(results ...interface{}) {
	self.Results = make([]interface{}, 0)
	self.Results = append(self.Results, results...)
}
