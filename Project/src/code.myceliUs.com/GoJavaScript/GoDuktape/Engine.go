package GoDuktape

//#cgo CFLAGS: -std=c99 -O2
/*
#define nullptr 0
#include "duktape.h"

// List of typedef.
typedef duk_context* duk_context_ptr;

// Create a default duk context.
duk_context_ptr create_default_context(){
	return duk_create_heap_default();
}

// Clear the default context.
void eval_string(duk_context_ptr context, const char* src){
	duk_eval_string(context,src);
}

// Compile a script.
void compile(duk_context_ptr context, int flags){
	duk_compile(context, flags);
}

*/
import "C"

import "unsafe"
import "code.myceliUs.com/GoJavaScript"

//import "reflect"

import "log"

/**
 * The duktape JavaScript engine.
 */
type Engine struct {
	// The debugger port.
	port int

	// The main duktape context.
	context C.duk_context_ptr
}

/**
 * Initialyse and start a new duktape JavaScript engine.
 */
func (self *Engine) Start(port int) {
	/* The port to communicate with the instance */
	self.port = port

	/** I will initialyse the main context. **/
	self.context = C.create_default_context()

}

func (self *Engine) Clear() {
	/* Clear the heap memory */
	C.duk_destroy_heap(self.context)
}

/////////////////// Global variables //////////////////////

/**
 * Set a variable on the global context.
 * name The name of the variable in the context.
 * value The value of the variable, can be a string, a number, an object.
 */
func (self *Engine) SetGlobalVariable(name string, value interface{}) {
	// Set the propertie in the global context..
	setValue(self.context, value)

	// Set the property in the global context.
	cstr := C.CString(name)
	C.duk_put_global_string(self.context, cstr)
	C.free(unsafe.Pointer(cstr))
}

/**
 * Return a variable define in the global object.
 */
func (self *Engine) GetGlobalVariable(name string) (GoJavaScript.Value, error) {

	// first of all I will initialyse the arguments.
	var value GoJavaScript.Value
	value.TYPENAME = "GoJavaScript.Value"

	cstr := C.CString(name)
	// Set the property at top of the stack.
	ret := C.int(C.duk_get_global_string(self.context, cstr))
	if ret > 0 {
		// In case of string
		v, err := getValue(self.context, -1)
		if err == nil {
			value.Val = v
		} else {
			return value, err
		}
	}

	C.free(unsafe.Pointer(cstr))

	return value, nil
}

/////////////////// Objects //////////////////////
/**
 * Create JavaScript object with given uuid. If name is given the object will be
 * set a global object property.
 */
func (self *Engine) CreateObject(uuid string, name string) {

}

/**
 * Set an object property.
 * uuid The object reference.
 * name The name of the property to set
 * value The value of the property
 */
func (self *Engine) SetObjectProperty(uuid string, name string, value interface{}) error {
	return nil
}

/**
 * That function is use to get Js object property
 */
func (self *Engine) GetObjectProperty(uuid string, name string) (GoJavaScript.Value, error) {
	var value GoJavaScript.Value
	return value, nil
}

/**
 * Create an empty array of a given size and set it as object property.
 */
func (self *Engine) CreateObjectArray(uuid string, name string, size uint32) error {
	return nil
}

/**
 * Set an object property.
 * uuid The object reference.
 * name The name of the property to set
 * index The index of the object in the array
 * value The value of the property
 */
func (self *Engine) SetObjectPropertyAtIndex(uuid string, name string, index uint32, value interface{}) {

}

/**
 * That function is use to get Js obeject property
 */
func (self *Engine) GetObjectPropertyAtIndex(uuid string, name string, index uint32) (GoJavaScript.Value, error) {
	var value GoJavaScript.Value
	return value, nil
}

/**
 * set to JS object a Go function.
 */
func (self *Engine) SetGoObjectMethod(uuid, name string) error {
	return nil
}

/**
 * Set to JS object a Js function.
 */
func (self *Engine) SetJsObjectMethod(uuid, name string, src string) error {
	return nil
}

/**
 * Call object methode.
 */
func (self *Engine) CallObjectMethod(uuid string, name string, params ...interface{}) (GoJavaScript.Value, error) {
	var value GoJavaScript.Value
	return value, nil
}

/////////////////// Functions //////////////////////

/**
 * Register a go function in JS
 */
func (self *Engine) RegisterGoFunction(name string) {

}

/**
 * Parse and set a function in the Javascript.
 * name The name of the function (the function will be keep in the engine for it
 *      lifetime.
 * args The argument name for that function.
 * src  The body of the function
 * options Can be JERRY_PARSE_NO_OPTS or JERRY_PARSE_STRICT_MODE
 */
func (self *Engine) RegisterJsFunction(name string, src string) error {
	cstr := C.CString(src)

	// eval the script.
	C.eval_string(self.context, cstr)
	C.free(unsafe.Pointer(cstr))
	return nil
}

/**
 * Call a Javascript function. The function must exist...
 */
func (self *Engine) CallFunction(name string, params []interface{}) (GoJavaScript.Value, error) {
	var value GoJavaScript.Value
	value.TYPENAME = "GoJavaScript.Value"

	cstr := C.CString(name)
	// Set the property at top of the stack.
	ret := C.int(C.duk_get_global_string(self.context, cstr))
	C.free(unsafe.Pointer(cstr))
	if ret > 0 {
		// so here the context point to the function.
		// I will append the list of arguments.
		// Now I will set the arguments...
		for i := 0; i < len(params); i++ {
			// So here I will set argument on the context.
			setValue(self.context, params[i])
		}

		C.duk_call(self.context, C.int(len(params)))

		// Now the result is at -1
		v, err := getValue(self.context, -1)
		if err == nil {
			log.Println("---> v ", v)
			value.Val = v
		}

		// remove the result from the stack.
		C.duk_pop(self.context)
	}

	return value, nil
}

/**
 * Evaluate a script.
 * src Contain the code to run.
 * variables Contain the list of variable to set on the global context before
 * running the script.
 */
func (self *Engine) EvalScript(src string, variables []interface{}) (GoJavaScript.Value, error) {
	var value GoJavaScript.Value
	value.TYPENAME = "GoJavaScript.Value"

	for i := 0; i < len(variables); i++ {
		self.SetGlobalVariable(variables[i].(*GoJavaScript.Variable).Name, variables[i].(*GoJavaScript.Variable).Value)
	}

	cstr := C.CString(src)

	// eval the script.
	C.eval_string(self.context, cstr)

	// Now the result is at -1
	v, err := getValue(self.context, -1)
	if err == nil {
		value.Val = v
	}

	return value, nil
}
