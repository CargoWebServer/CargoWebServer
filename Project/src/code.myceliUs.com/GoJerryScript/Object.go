package GoJerryScript

//import "log"
import "code.myceliUs.com/Utility"
import "fmt"
import "reflect"

/**
 * Go representation of a JS object.
 */
type Object struct {

	// The peer where the object live.
	peer *Peer

	// If a name is given the object will be set
	// as a property of the global JS object
	name string

	// The unique object identifier.
	uuid string
}

/**
 * Go handle of JS object.
 */
func NewObject(name string) *Object {

	// The object itself.
	obj := new(Object)
	ptrString := fmt.Sprintf("%d", obj)
	uuid := Utility.GenerateUUID(ptrString)
	obj.uuid = uuid
	obj.name = name

	// Here I will keep the object in the client cache.
	GetCache().SetObject(obj.uuid, obj)

	return obj
}

/**
 * Peer are use to transfert the object informations between client and server side
 */
func (self *Object) SetPeer(peer *Peer) {
	// I will set the object peer.
	self.peer = peer

	// So here I will create the action.
	action := new(Action)
	action.UUID = Utility.RandomUUID()
	action.Name = "CreateObject"
	action.AppendParam("uuid", self.uuid)
	action.AppendParam("name", self.name)

	// Call the action here.
	peer.CallRemoteAction(action)
}

/**
 * Return property value.
 */
func (self *Object) Get(name string) (*Value, error) {
	action := new(Action)
	action.UUID = Utility.RandomUUID()
	action.Name = "GetObjectProperty"
	action.AppendParam("uuid", self.uuid)
	action.AppendParam("name", name)

	// Call the action here.
	action = self.peer.CallRemoteAction(action)
	var result Value
	var err error

	if action.Results[0] != nil {
		result = action.Results[0].(Value)
	}

	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	return &result, err
}

/**
 * Set object property
 * name The name of the method in the object.
 * value The value that method return, it can be a Go function.
 */
func (self *Object) Set(name string, value interface{}) {
	// If the value is a function
	if reflect.TypeOf(value).Kind() == reflect.Func {
		// In that case I will register the function.
		Utility.RegisterFunction(name, value)
		action := new(Action)
		action.UUID = Utility.RandomUUID()
		action.Name = "SetGoObjectMethod"
		action.AppendParam("uuid", self.uuid)
		action.AppendParam("name", name)

		// Call the action here.
		self.peer.CallRemoteAction(action)

	} else {

		action := new(Action)
		action.UUID = Utility.RandomUUID()
		action.Name = "SetObjectProperty"
		action.AppendParam("uuid", self.uuid)
		action.AppendParam("name", name)
		action.AppendParam("value", value)

		// Call the action here.
		self.peer.CallRemoteAction(action)
	}
}

/**
 * Set a JS function as Object methode.
 */
func (self *Object) SetJsMethode(name string, src string) {
	action := new(Action)
	action.UUID = Utility.RandomUUID()
	action.Name = "SetJsObjectMethod"
	action.AppendParam("uuid", self.uuid)
	action.AppendParam("name", name)
	action.AppendParam("src", src)

	// Call the action here.
	self.peer.CallRemoteAction(action)
}

/**
 * Call a function over an object.
 */
func (self *Object) Call(name string, params ...interface{}) (Value, error) {

	action := new(Action)
	action.UUID = Utility.RandomUUID()
	action.Name = "CallObjectMethod"
	action.AppendParam("uuid", self.uuid)
	action.AppendParam("name", name)

	action.AppendParam("params", params)

	// Call the action here.
	action = self.peer.CallRemoteAction(action)

	var result Value
	var err error

	if action.Results[0] != nil {
		result = action.Results[0].(Value)
	}

	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	return result, err
}
