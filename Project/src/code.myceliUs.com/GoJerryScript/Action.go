package GoJerryScript

import "reflect"
import "log"

import "encoding/gob"
import "bytes"

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
	Params []*Param

	// The action results.
	Results []interface{}
}

// Append a new parameter to an action.
// TODO test if the parameter is an object reference in that case send
// object info and not the object itself. To the same for results.
func (self *Action) AppendParam(name string, value interface{}) {
	param := new(Param)
	param.Name = name

	// Here I will use the go reflection to get the type of
	// the value...
	param.Type = reflect.TypeOf(value).String()

	// Set the parameter value.
	param.Value = value

	if self.Params == nil {
		self.Params = make([]*Param, 0)
	}

	self.Params = append(self.Params, param)
}

// Append the action results.
func (self *Action) AppendResults(results ...interface{}) {
	self.Results = make([]interface{}, 0)
	log.Println("---> append results: ", results)
	self.Results = append(self.Results, results...)
}

// Serialyse the content of the action
func MarshalAction(action *Action) []byte {
	// Be sure those struct are registered.
	gob.Register(Value{})
	gob.Register([]Variable{})

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(action)
	if err != nil {
		log.Println("---> encoding error: ", err)
		return nil
	}
	return buffer.Bytes()
}

func UnmarshalAction(data []byte) *Action {

	gob.Register(Value{})
	gob.Register([]Variable{})

	action := new(Action)
	buffer := bytes.NewReader(data)

	// Decode the action.
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(action)
	if err != nil {
		log.Println("---> decoding error: ", err)
		return nil
	}
	return action
}
