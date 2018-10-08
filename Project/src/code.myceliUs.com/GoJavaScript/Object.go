package GoJavaScript

import "code.myceliUs.com/Utility"
import "reflect"

// The object must be contain entirely in the client a get by the server when he's
// ready to use it with getObjectByUuid... so no channel will dead lock.
/**
 * Go representation of a JS object.
 */
type Object struct {
	// The the type name.
	//TYPENAME string

	// The unique object identifier.
	UUID string

	// If a name is given the object will be set
	// as a property of the global JS object
	Name string

	// Contain method name for Go and Js method
	// and Js code for Js method.
	// the store value can be a string or []uint8 depending if the code

	Methods map[string]interface{}

	// Contain the properties of the object.
	Properties map[string]interface{}

	// The peer where the object live.
	peer *Peer
}

/**
 * Go handle of JS object.
 */
func NewObject(name string) *Object {

	// The object itself.
	obj := new(Object)

	//obj.TYPENAME = "GoJavaScript.Object"

	// If the name is given that's mean the object will be set as a global
	// object so it uuid will be generated from it name.
	if len(name) > 0 {
		// Thats mean the value is a global variable.
		obj.UUID = Utility.GenerateUUID(name)
	} else {
		obj.UUID = Utility.RandomUUID()
	}

	// The name of the object.
	obj.Name = name

	// Set the properties.

	// Contain method name for Go and Js method
	// and Js code for Js method.
	obj.Methods = make(map[string]interface{}, 0)

	// Contain the properties of the object.
	obj.Properties = make(map[string]interface{}, 0)

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

	// Here I have to register the object on the server if it name is define
	// that will be a global variable.
	if len(self.Name) > 0 {
		action := NewAction("CreateObject", "")
		action.AppendParam("uuid", self.UUID)
		action.AppendParam("name", self.Name)

		// Call the action here.
		action = self.peer.CallRemoteAction(action)
	}
}

/**
 * Return property value.
 */
func (self *Object) Get(name string) (interface{}, error) {
	return self.Properties[name], nil
}

/**
 * Set object property
 * name The name of the method in the object.
 * value The value that method return, it can be a Go function.
 */
func (self *Object) Set(name string, value interface{}) {
	// If the value is a function
	if reflect.TypeOf(value).Kind() == reflect.Func {
		// Set a go method.
		Utility.RegisterFunction(name, value)
		self.Methods[name] = ""
	} else {
		// Set object property
		self.Properties[name] = value
	}
}

/**
 * Set a JS function as Object methode.
 */
func (self *Object) SetJsMethode(name string, src string) {
	self.Methods[name] = src
}

/**
 * Call a function over an object.
 */
func (self *Object) Call(name string, params ...interface{}) (interface{}, error) {
	action := NewAction("CallObjectMethod", "")
	action.AppendParam("uuid", self.UUID)
	action.AppendParam("name", name)

	// I will set parameter object ref before call.
	action.AppendParam("params", ObjectToRef(params))

	// Call the action here.
	action = self.peer.CallRemoteAction(action)

	var err error
	var result interface{}

	if action.Results[0] != nil {
		result = RefToObject(action.Results[0])
	}

	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	return result, err
}
