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
 * Set a variable on the global context.
 * name The name of the variable in the context.
 * value The value of the variable, can be a string, a number,
 */
func (self *Engine) SetVariable(name string, value interface{}) {
	// first of all I will initialyse the arguments.
	globalObject := Jerry_get_global_object()

	propName := Jerry_create_string(NewUint8FromString(name))

	var propValue Uint32_t
	if reflect.TypeOf(value).Kind() == reflect.String {
		// String value
		propValue = Jerry_create_string(NewUint8FromString(value.(string)))
	} else if reflect.TypeOf(value).Kind() == reflect.Bool {
		// Boolean value
		propValue = Jerry_create_boolean(value.(bool))
	} else if reflect.TypeOf(value).Kind() == reflect.Int {
		propValue = Jerry_create_number(float64(value.(int)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int8 {
		propValue = Jerry_create_number(float64(value.(int8)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int16 {
		propValue = Jerry_create_number(float64(value.(int16)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int32 {
		propValue = Jerry_create_number(float64(value.(int32)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int64 {
		propValue = Jerry_create_number(float64(value.(int64)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint {
		propValue = Jerry_create_number(float64(value.(uint)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint8 {
		propValue = Jerry_create_number(float64(value.(uint8)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint16 {
		propValue = Jerry_create_number(float64(value.(uint16)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint32 {
		propValue = Jerry_create_number(float64(value.(uint32)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint64 {
		propValue = Jerry_create_number(float64(value.(uint64)))
	} else if reflect.TypeOf(value).Kind() == reflect.Float32 {
		propValue = Jerry_create_number(float64(value.(float32)))
	} else if reflect.TypeOf(value).Kind() == reflect.Float64 {
		propValue = Jerry_create_number(value.(float64))
	}

	// Set the propertie in the global context..
	Jerry_set_property(globalObject, propName, propValue)

	// Release the resource as no more needed here.
	Jerry_release_value(propName)
	Jerry_release_value(propValue)

	Jerry_release_value(globalObject)

}

/**
 * Evaluate a javascript function with a given name and arguments list.
 * 			*The fuction must be set by AppendFunction before being use here.
 */
func (self *Engine) EvalFunction(name string, args []interface{}) (interface{}, error) {
	if fct, ok := self.functions[name]; ok {
		// Now I will test if argument are given.
		if len(fct.Args) != len(args) {
			return nil, errors.New("Function named " + name + " expect " + strconv.Itoa(len(fct.Args)) + " argument(s) not " + strconv.Itoa(len(args)))
		}

		for i := 0; i < len(args); i++ {
			self.SetVariable(fct.Args[i], args[i])
		}

		// Now I will evaluate the function...
		ret := Jerry_eval(NewUint8FromString(fct.Body), int64(len(fct.Body)), false)

		// Get the Go value...
		var value interface{}

		typeInfo := Jerry_value_get_type(ret)

		// Now I will get the result if any...
		if typeInfo == Jerry_type_t(JERRY_TYPE_NUMBER) {
			value = Jerry_get_number_value(ret)
		} else if typeInfo == Jerry_type_t(JERRY_TYPE_STRING) {
			// Size info, ptr and it value
			sizePtr := Jerry_get_string_size(ret)
			sizeValue := *(*uint64)(unsafe.Pointer(sizePtr.Swigcptr()))
			var buffer Uint8
			buffer.ptr = C.malloc(C.ulong(sizeValue)) // allocated the size.

			// Test if the string is a valid utf8 string...

			Jerry_string_to_char_buffer(ret, buffer, sizePtr)

			// Set the slice information...
			h := &reflect.SliceHeader{buffer.Swigcptr(), int(sizeValue), int(sizeValue)}
			NewSlice := *(*[]byte)(unsafe.Pointer(h))

			// Set the Go string value from the slice.
			value = string(NewSlice)

			// release the memory used by the buffer.
			buffer.Free()
		}

		// Free JavaScript value, returned by eval
		Jerry_release_value(ret)
		return value, nil
	} else {
		return nil, errors.New("No function found with name " + name)
	}
}

func (self *Engine) Clear() {

	/* Cleanup the script engine. */
	Jerry_cleanup()

}
