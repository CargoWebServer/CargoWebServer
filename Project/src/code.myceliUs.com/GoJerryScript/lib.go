package GoJerryScript

//#cgo LDFLAGS: -L/usr/local/lib -ljerry-core -ljerry-ext -ljerry-libm -ljerry-port-default
/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "jerryscript.h"
#include "jerryscript-ext/handler.h"
#include "jerryscript-debugger.h"

void register_print_handler() {
     jerryx_handler_register_global ((const jerry_char_t *) "print", jerryx_handler_print);
}

*/
import "C"
import "reflect"
import "unsafe"
import "errors"
import "strconv"

//import "log"

/**
 * Register the print handler.
 */
func RegisterPrintHandler() {
	C.register_print_handler()
}

// Various type conversion functions.
func unsafeStrToByte(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	var b []byte
	byteHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	byteHeader.Data = strHeader.Data

	// need to take the length of s here to ensure s is live until after we update b's Data
	// field since the garbage collector can collect a variable once it is no longer used
	// not when it goes out of scope, for more details see https://github.com/golang/go/issues/9046
	l := len(s)
	byteHeader.Len = l
	byteHeader.Cap = l
	return b
}

////////////// Uint 8 //////////////
// The Uint8 Type represent a 8 bit char.
type Uint8 struct {
	// The pointer that old the data.
	ptr unsafe.Pointer
}

/**
 * Free the values.
 */
func (self Uint8) Free() {
	C.free(unsafe.Pointer(self.ptr))
}

/**
 * Access the undelying memeory values pointer.
 */
func (self Uint8) Swigcptr() uintptr {
	return uintptr(self.ptr)
}

func NewUint8FromString(str string) Uint8 {
	// convert string to []bytes
	data := unsafeStrToByte(str)
	var val Uint8

	// keep the pointer value inside the
	val.ptr = C.CBytes(data)
	return val
}

func NewUint8FromBytes(data []uint8) Uint8 {
	var val Uint8
	val.ptr = C.CBytes(data)
	return val
}

////////////// Uint 16 //////////////
// The Uint16 Type represent a 16 bit char.
type Uint16 struct {
	// The pointer that old the data.
	ptr unsafe.Pointer
}

/**
 * Free the values.
 */
func (self Uint16) Free() {
	C.free(unsafe.Pointer(self.ptr))
}

/**
 * Access the undelying memeory values pointer.
 */
func (self Uint16) Swigcptr() uintptr {
	return uintptr(self.ptr)
}

////////////// Uint 32 //////////////
// The Uint32 Type represent a 32 bit char.
type Uint32 struct {
	// The pointer that old the data.
	ptr unsafe.Pointer
}

func NewUint32FromInt(i int32) Uint32 {
	var val Uint32
	val.ptr = unsafe.Pointer(&i)
	return val
}

/**
 * Free the values.
 */
func (self Uint32) Free() {
	C.free(unsafe.Pointer(self.ptr))
}

/**
 * Access the undelying memeory values pointer.
 */
func (self Uint32) Swigcptr() uintptr {
	return uintptr(self.ptr)
}

////////////// Instance //////////////

// Reference to an object.
type Instance struct {
	// The pointer that old the data.
	ptr unsafe.Pointer
}

func NewInstance(obj interface{}) Instance {
	var instance Instance
	return instance
}

/**
 * Free the values.
 */
func (self Instance) Free() {
	C.free(unsafe.Pointer(self.ptr))
}

/**
 * Access the undelying memeory values pointer.
 */
func (self Instance) Swigcptr() uintptr {
	return uintptr(self.ptr)
}

/**
 * Keep function informations.
 */
type Function struct {
	Name string
	Args []string
	Body string
}

/**
 * The JerryScript JS engine.
 */
type Engine struct {
	port      int
	functions map[string]Function
}

func NewEngine(port int, options int) *Engine {
	engine := new(Engine)
	// keep function pointer here.
	engine.functions = make(map[string]Function, 0)
	engine.start(port, options)
	return engine
}

/**
 * Init and start the engine.
 * port The port to communicate with the engine, or the debbuger.
 * option Can be JERRY_INIT_EMPTY, JERRY_INIT_SHOW_OPCODES, JERRY_INIT_SHOW_REGEXP_OPCODES
 *		  JERRY_INIT_MEM_STATS, 	JERRY_INIT_MEM_STATS_SEPARATE (or a combination option)
 */
func (self *Engine) start(port int, options int) {
	/* The port to communicate with the instance */
	self.port = port

	/* Init the script engine. */
	Jerry_init(Jerry_init_flag_t(options))

	/* Register 'print' function  */
	RegisterPrintHandler()
}

/**
 * Parse and set a function in the Javascript.
 * name The name of the function (the function will be keep in the engine for it
 *      lifetime.
 * args The argument name for that function.
 * src  The body of the function
 * options Can be JERRY_PARSE_NO_OPTS or JERRY_PARSE_STRICT_MODE
 */
func (self *Engine) AppendFunction(name string, args []string, src string, options int) error {
	/* The name of the function */
	arg0 := NewUint8FromString(name)
	args_ := ""
	for i := 0; i < len(args); i++ {
		args_ += args[i]
		if i < len(args)-1 {
			args_ += ", "
		}
	}
	arg1 := NewUint8FromString(args_)
	arg2 := NewUint8FromString(src)
	arg3 := NewUint32FromInt(int32(options))

	parsed_code := Jerry_parse_function(arg0, int64(len(name)), arg1, int64(len(args_)), arg2, int64(len(src)), arg3)
	if !Jerry_value_is_error(parsed_code) {
		Jerry_release_value(parsed_code)
		self.functions[name] = Function{Name: name, Args: args, Body: src}
		return nil
	} else {
		Jerry_release_value(parsed_code)
		return errors.New("Fail to parse function " + name)
	}
}

/**
 * Evaluate a function with a given name and arguments.
 */
func (self *Engine) EvalFunction(name string, args []string) error {
	if fct, ok := self.functions[name]; ok {
		// Now I will test if argument are given.
		if len(fct.Args) != len(args) {
			return errors.New("Function named " + name + " expect " + strconv.Itoa(len(fct.Args)) + " argument(s) not " + strconv.Itoa(len(args)))
		}

		// first of all I will initialyse the arguments.
		globalObject := Jerry_get_global_object()

		for i := 0; i < len(args); i++ {
			propName := Jerry_create_string(NewUint8FromString(fct.Args[i]))
			arg := args[i]
			propValue := Jerry_create_string(NewUint8FromString(arg))

			// Set the propertie in the global context..
			Jerry_set_property(globalObject, propName, propValue)

			// Release the resource as no more needed here.
			Jerry_release_value(propName)
			Jerry_release_value(propValue)
		}

		Jerry_release_value(globalObject)

		// Now I will evaluate the function...
		ret := Jerry_eval(NewUint8FromString(fct.Body), int64(len(fct.Body)), false)

		// Free JavaScript value, returned by eval
		Jerry_release_value(ret)

		return nil
	} else {
		return errors.New("No function found with name " + name)
	}
}

func (self *Engine) Clear() {

	/* Cleanup the script engine. */
	Jerry_cleanup()

}
