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
duk_int_t eval_string(duk_context_ptr context, const char* src){
	return duk_peval_string(context,src);
}

extern duk_int_t compile_function_string(duk_context_ptr ctx, const char* src){
	return duk_pcompile_string(ctx, DUK_COMPILE_FUNCTION, src);
}

// The function handler.
extern duk_ret_t c_function_handler(duk_context_ptr ctx);

// Set C function handler.
duk_idx_t push_c_function(duk_context_ptr ctx, const char* name){
	duk_idx_t fct_idx = duk_push_c_function(ctx, c_function_handler, DUK_VARARGS);

	// Set the function name as property...
	duk_push_string(ctx, "name");
	duk_push_string(ctx, name);
	duk_def_prop(ctx, -3, DUK_DEFPROP_HAVE_VALUE);

	return fct_idx;
}

duk_bool_t is_error(duk_context_ptr ctx, duk_idx_t index){
	return duk_is_error(ctx, index);
}

const char* safe_to_string(duk_context_ptr ctx, duk_idx_t idx){
	return duk_safe_to_string(ctx, idx);
}

*/
import "C"

import "unsafe"
import "code.myceliUs.com/GoJavaScript"

//import "reflect"
import "errors"
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
	log.Println("--> create object: ", uuid, name)
	obj_idx := C.duk_push_object(self.context)

	// Now I will set the uuid property.

	// uuid value
	uuid_ := C.CString(uuid)
	C.duk_push_string(self.context, uuid_)
	C.free(unsafe.Pointer(uuid_))

	// uuid name
	uuid_ = C.CString("uuid_")
	C.duk_put_prop_string(self.context, obj_idx, uuid_)
	C.free(unsafe.Pointer(uuid_))

	if len(name) > 0 {
		name_ := C.CString(name)
		C.duk_put_global_string(self.context, name_)
		C.free(unsafe.Pointer(name_))
	}
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
	value.TYPENAME = "GoJavaScript.Value"

	// I will get (create the object on the stack)
	getJsObjectByUuid(uuid, self.context)

	// go to the object position.
	uuid_ := C.CString(uuid)
	C.duk_get_global_string(self.context, uuid_)

	// Set the function name to be call
	cstr := C.CString(name)
	C.duk_push_string(self.context, cstr)
	defer C.free(unsafe.Pointer(cstr))

	// Put the method in the context.
	C.duk_get_prop_string(self.context, C.int(-2), cstr)

	// Put the this object pointer.
	C.duk_dup(self.context, -2)

	// Set the arguments...
	for i := 0; i < len(params); i++ {
		// So here I will set argument on the context.
		setValue(self.context, params[i])
	}

	// Call the method.
	if int(C.duk_pcall_method(self.context, C.int(len(params)))) == 0 {
		// Now the result is at -1
		v, err := getValue(self.context, -1)
		if err == nil {
			value.Val = v
		}
		// remove the result from the stack.
		C.duk_pop(self.context) // Pop the call result.
	} else {
		err := errors.New(C.GoString(C.safe_to_string(self.context, -1)))
		log.Println("268 ---> error found!", err)
		return value, err
	}

	C.duk_pop(self.context) // Pop the instance.

	return value, nil
}

/////////////////// Functions //////////////////////

/**
 * Register a go function in JS
 */
func (self *Engine) RegisterGoFunction(name string) {

	// Keep the go function in memory...
	cstr := C.CString(name)
	C.push_c_function(self.context, cstr)
	C.duk_put_global_string(self.context, cstr)

	// The propertie name value
	C.free(unsafe.Pointer(cstr))
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
	defer C.free(unsafe.Pointer(cstr))

	// eval the script.
	if int(C.eval_string(self.context, cstr)) != 0 {
		return errors.New(C.GoString(C.safe_to_string(self.context, -1)))
	}
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
