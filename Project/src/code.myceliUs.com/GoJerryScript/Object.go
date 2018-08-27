package GoJerryScript

//import "log"
import "code.myceliUs.com/Utility"
import "reflect"

/**
 * Go representation of a JS object.
 */
type Object struct {
	// The typename.
	TYPENAME string

	// The peer where the object live.
	peer *Peer

	// If a name is given the object will be set
	// as a property of the global JS object
	Name string

	// The unique object identifier.
	UUID string
}

/**
 * Go handle of JS object.
 */
func NewObject(name string) *Object {

	// The object itself.
	obj := new(Object)
	obj.TYPENAME = "GoJerryScript.Object"

	// If the name is given that's mean the object will be set as a global
	// object so it uuid will be generated from it name.
	if len(name) > 0 {
		// Thats mean the value is a global variable.
		obj.UUID = Utility.GenerateUUID(name)
	} else {
		obj.UUID = Utility.RandomUUID()
	}

	obj.Name = name

	// Here I will keep the object in the client cache.
	GetCache().SetObject(obj.UUID, obj)

	return obj
}

/**
 * Peer are use to transfert the object informations between client and server side
 */
func (self *Object) SetPeer(peer *Peer) {
	// I will set the object peer.
	self.peer = peer

	// So here I will create the action.
	action := NewAction("CreateObject", "")
	action.AppendParam("uuid", self.UUID)
	action.AppendParam("name", self.Name)

	// Call the action here.
	peer.CallRemoteAction(action)
}

/**
 * Return property value.
 */
func (self *Object) Get(name string) (Value, error) {

	action := NewAction("GetObjectProperty", "")
	action.AppendParam("uuid", self.UUID)
	action.AppendParam("name", name)

	// Call the action here.
	action = self.peer.CallRemoteAction(action)
	var result *Value
	var err error

	if action.Results[0] != nil {
		result = action.Results[0].(*Value)
	}

	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	return *result, err
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
		action := NewAction("SetGoObjectMethod", "")

		action.AppendParam("uuid", self.UUID)
		action.AppendParam("name", name)

		// Call the action here.
		self.peer.CallRemoteAction(action)

	} else {

		action := NewAction("SetObjectProperty", "")

		action.AppendParam("uuid", self.UUID)
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
	action := NewAction("SetJsObjectMethod", "")
	action.AppendParam("uuid", self.UUID)
	action.AppendParam("name", name)
	action.AppendParam("src", src)

	// Call the action here.
	self.peer.CallRemoteAction(action)
}

/**
 * Call a function over an object.
 */
func (self *Object) Call(name string, params ...interface{}) (Value, error) {

	action := NewAction("CallObjectMethod", "")
	action.AppendParam("uuid", self.UUID)
	action.AppendParam("name", name)

	action.AppendParam("params", params)

	// Call the action here.
	action = self.peer.CallRemoteAction(action)

	var result *Value
	var err error

	if action.Results[0] != nil {
		result = action.Results[0].(*Value)
	}

	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	return *result, err
}
