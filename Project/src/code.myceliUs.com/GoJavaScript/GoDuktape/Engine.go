package GoDuktape

//#cgo CFLAGS: -std=c99 -O2
/*
#define nullptr 0

// enable debuger features.
#define DUK_USE_DEBUGGER_SUPPORT
#define DUK_USE_INTERRUPT_COUNTER

#include "duktape.h"

// List of typedef.
typedef duk_context* duk_context_ptr;

// Create a default duk context.
duk_context_ptr create_default_context(){
	return duk_create_heap_default();
}

// Clear the default context.
duk_int_t eval_string(duk_context_ptr ctx, const char* src){
	return duk_peval_string(ctx,src);
}

extern duk_int_t compile_function_string(duk_context_ptr ctx, const char* src){
	return duk_pcompile_string(ctx, DUK_COMPILE_FUNCTION, src);
}

// The object finalyser function.
extern duk_ret_t c_finalizer(duk_context_ptr);

// Set the delete callback function
duk_idx_t set_finalizer(duk_context_ptr ctx, const char* uuid){

	duk_idx_t fct_idx = duk_push_c_function(ctx, c_finalizer, 1);

	// Set the function name as property...
	duk_push_string(ctx, "uuid");
	duk_push_string(ctx, uuid);
	duk_def_prop(ctx, -3, DUK_DEFPROP_HAVE_VALUE);

	// set the object finaliser.
	duk_set_finalizer(ctx, -2);

	return fct_idx;
}

// The function handler.
extern duk_ret_t c_function_handler(duk_context_ptr ctx, const char* name, const char* uuid);

duk_ret_t c_function_handler_(duk_context_ptr ctx){
	// Push the current function on the context
	duk_push_current_function(ctx);

	// Get it name property.
	duk_get_prop_string(ctx, -1, "name");

	const char* name = duk_to_string(ctx, -1);

	duk_pop(ctx); // pop name.

	// push this in the context.
	duk_push_this(ctx);

	const char* uuid = 0;
	if(duk_is_object(ctx, -1)){
		if(duk_has_prop_string(ctx, -1, "uuid_")) {
			duk_get_prop_string(ctx, -1, "uuid_");
			uuid = duk_to_string(ctx, -1);
			duk_pop(ctx); // pop prop value (uuid)
		}
	}

	duk_pop(ctx); // pop this

	// Call the go handler function with it parematers.
	return c_function_handler(ctx, name, uuid);
}

// Set C function handler.
duk_idx_t push_c_function(duk_context_ptr ctx, const char* name){
	duk_idx_t fct_idx = duk_push_c_function(ctx, c_function_handler_, DUK_VARARGS);

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

// Load a function from bytecode.
void load_byte_code(duk_context_ptr ctx, char* bytecode, int length){
	char* dukBuff = (char*)duk_push_fixed_buffer(ctx, length); // push a duk buffer to stack
	memcpy(dukBuff, bytecode, length); // copy the bytecode to the duk buffer
}
*/
import "C"

import "unsafe"
import "code.myceliUs.com/GoJavaScript"
import "fmt"
import "errors"
import "log"
import "os"
import b64 "encoding/base64"

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
func (self *Engine) GetGlobalVariable(name string) (interface{}, error) {

	// first of all I will initialyse the arguments.
	var value interface{}
	cstr := C.CString(name)

	// Set the property at top of the stack.
	ret := C.int(C.duk_get_global_string(self.context, cstr))
	if ret > 0 {
		// In case of string
		v, err := getValue(self.context, -1)
		if err == nil {
			value = v
		} else {
			return value, err
		}
		C.duk_pop(self.context)
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

	obj_idx := C.duk_push_object(self.context)

	// uuid value
	uuid_value := C.CString(uuid)
	C.set_finalizer(self.context, uuid_value)
	C.duk_push_string(self.context, uuid_value)
	uuid_name := C.CString("uuid_")
	C.duk_put_prop_string(self.context, obj_idx, uuid_name)

	C.free(unsafe.Pointer(uuid_value))
	C.free(unsafe.Pointer(uuid_name))

	if len(name) > 0 {
		name_ := C.CString(name)
		C.duk_put_global_string(self.context, name_)
		C.free(unsafe.Pointer(name_))
	}
}

/**
 * Call object methode.
 */
func (self *Engine) CallObjectMethod(uuid string, name string, params ...interface{}) (interface{}, error) {
	var value interface{}

	// I will get (create the object on the stack)
	getJsObjectByUuid(uuid, self.context)

	// Set the function name to be call
	cstr := C.CString(name)
	C.duk_push_string(self.context, cstr)
	defer C.free(unsafe.Pointer(cstr))

	// Put the method in the context.
	C.duk_get_prop_string(self.context, C.int(-2), cstr)

	// Put the this object pointer.
	C.duk_dup(self.context, -3)

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
			value = v
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
func (self *Engine) CallFunction(name string, params []interface{}) (interface{}, error) {
	// So here I will not crash the engine if there is error but
	// I will display the error in the console and recover.
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
		}
	}()

	// so here the context point to the function.
	// I will append the list of arguments.
	// Now I will set the arguments...
	var value interface{}
	cstr := C.CString(name)

	// Set the property at top of the stack.
	ret := C.int(C.duk_get_global_string(self.context, cstr))
	C.free(unsafe.Pointer(cstr))

	if ret == 0 {
		byteCode, err := b64.StdEncoding.DecodeString(name)
		if err == nil {
			byteCode_ := (*C.char)(C.CBytes(byteCode))
			C.load_byte_code(self.context, byteCode_, C.int(len(byteCode)))
			C.free(unsafe.Pointer(byteCode_))

			// replaces the buffer on the stack top with the original function
			C.duk_load_function(self.context)
		}
	}

	for i := 0; i < len(params); i++ {
		// Set value
		setValue(self.context, params[i])
	}

	C.duk_call(self.context, C.int(len(params)))

	// Now the result is at -1
	v, err := getValue(self.context, -1)
	if err == nil {
		value = v
	}

	// remove the result from the stack.
	C.duk_pop(self.context)

	return value, err
}

/**
 * Evaluate a script.
 * src Contain the code to run.
 * variables Contain the list of variable to set on the global context before
 * running the script.
 */
func (self *Engine) EvalScript(src string, variables []interface{}) (interface{}, error) {
	var value interface{}

	for i := 0; i < len(variables); i++ {
		self.SetGlobalVariable(variables[i].(*GoJavaScript.Variable).Name, variables[i].(*GoJavaScript.Variable).Value)
	}

	cstr := C.CString(src)
	defer C.free(unsafe.Pointer(cstr))

	// eval the script.
	if int(C.eval_string(self.context, cstr)) != 0 {
		return value, errors.New(C.GoString(C.safe_to_string(self.context, -1)))
	}

	// Now the result is at -1
	v, err := getValue(self.context, -1)

	if err == nil {
		value = v
	}

	return value, err
}
