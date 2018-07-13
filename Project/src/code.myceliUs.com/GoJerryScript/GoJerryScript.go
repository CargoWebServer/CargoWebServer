package GoJerryScript

//#cgo LDFLAGS: -L/usr/local/lib -ljerry-core -ljerry-ext -ljerry-libm -ljerry-port-default
/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "jerryscript.h"
#include "jerryscript-ext/handler.h"
#include "jerryscript-debugger.h"

// External function call.
static jerry_value_t
handler (const jerry_value_t function_obj,
         const jerry_value_t this_val,
         const jerry_value_t args_p[],
         const jerry_length_t args_cnt)
{
	// So here I will call the go function.

}

*/
import "C"
import "reflect"
import "unsafe"
import "errors"
import "encoding/binary"
import "math"
import "encoding/json"
import "code.myceliUs.com/Utility"

//import "strconv"
import "log"

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

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
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
	Obj  Uint32_t
}

/**
 * Property are variable with name and value.
 */
type Variable struct {
	Name  string
	Value interface{}
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
		self.functions[name] = Function{Name: name, Args: args, Body: src, Obj: parsed_code}
		_, err := self.EvalScript(src, []Variable{})
		return err
	} else {
		Jerry_release_value(parsed_code)
		return errors.New("Fail to parse function " + name)
	}
}

/**
 * Register a go type to be usable as JS type.
 */
func (self *Engine) RegisterGoType(value interface{}) {
	Utility.RegisterType(value)
}

/**
 * Set a variable on the global context.
 * name The name of the variable in the context.
 * value The value of the variable, can be a string, a number,
 */
func (self *Engine) SetGlobalVariable(name string, value interface{}) {
	// first of all I will initialyse the arguments.
	globalObject := Jerry_get_global_object()

	propName := Jerry_create_string(NewUint8FromString(name))

	propValue := goToJs(value)

	// Set the propertie in the global context..
	Jerry_set_property(globalObject, propName, propValue)

	// Release the resource as no more needed here.
	Jerry_release_value(propName)
	Jerry_release_value(propValue)

	Jerry_release_value(globalObject)
}

func goToJs(value interface{}) Uint32_t {

	var propValue Uint32_t
	if reflect.TypeOf(value).Kind() == reflect.String {
		// String value
		propValue = Jerry_create_string_from_utf8(NewUint8FromString(value.(string)))
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
	} else if reflect.TypeOf(value).Kind() == reflect.Slice {
		// So here I will create a array and put value in it.
		var size Uint32
		s := reflect.ValueOf(value)
		l := s.Len()
		size.ptr = unsafe.Pointer(&l)
		propValue = Jerry_create_array(size)
		for i := 0; i < l; i++ {
			v := goToJs(s.Index(i).Interface())
			var index Uint32
			index.ptr = unsafe.Pointer(&i)
			Jerry_set_property_by_index(propValue, index, v)

			// Release memory
			Jerry_release_value(v)
		}
	} else if reflect.TypeOf(value).Kind() == reflect.Struct {

	}

	return propValue
}

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

func uintptrToBytes(u *uintptr) []byte {
	return (*[sizeOfUintPtr]byte)(unsafe.Pointer(u))[:]
}

/**
 * Create a go string from a JS string pointer.
 */
func jsStrToGoStr(input Uint32_t) string {
	// Size info, ptr and it value
	sizePtr := Jerry_get_string_size(input)
	sizeValue := *(*uint64)(unsafe.Pointer(sizePtr.Swigcptr()))
	var buffer Uint8
	buffer.ptr = C.malloc(C.ulong(sizeValue)) // allocated the size.

	// Test if the string is a valid utf8 string...
	isUtf8 := Jerry_is_valid_utf8_string(input, sizePtr)
	if isUtf8 {
		Jerry_string_to_utf8_char_buffer(input, buffer, sizePtr)
	} else {
		Jerry_string_to_char_buffer(input, buffer, sizePtr)
	}
	// Set the slice information...
	h := &reflect.SliceHeader{buffer.Swigcptr(), int(sizeValue), int(sizeValue)}
	NewSlice := *(*[]byte)(unsafe.Pointer(h))

	// Set the Go string value from the slice.
	value := string(NewSlice)

	// release the memory used by the buffer.
	buffer.Free()

	return value
}

/**
 * Return equivalent value of a 32 bit c pointer.
 */
func jsToGo(input Uint32_t) (interface{}, error) {

	// the Go value...
	var value interface{}
	typeInfo := Jerry_value_get_type(input)
	// Now I will get the result if any...
	if typeInfo == Jerry_type_t(JERRY_TYPE_NUMBER) {
		value = Jerry_get_number_value(input)
	} else if typeInfo == Jerry_type_t(JERRY_TYPE_STRING) {
		value = jsStrToGoStr(input)
	} else if typeInfo == Jerry_type_t(JERRY_TYPE_BOOLEAN) {
		value = Jerry_get_boolean_value(input)
	} else if Jerry_value_is_typedarray(input) {
		/** Not made use of typed array **/
	} else if Jerry_value_is_array(input) {
		countPtr := Jerry_get_array_length(input)
		count := *(*uint64)(unsafe.Pointer(countPtr.Swigcptr()))
		// So here I got a array without type so I will get it property by index
		// and interpret each result.
		value = make([]interface{}, 0)
		var i uint64
		for i = 0; i < count; i++ {
			var index Uint32
			index.ptr = unsafe.Pointer(&i)
			e := Jerry_get_property_by_index(input, index)
			v, err := jsToGo(e)
			if err == nil {
				value = append(value.([]interface{}), v)
			}
			Jerry_release_value(e)
		}

	} else if Jerry_value_is_object(input) {
		// The go object will be a copy of the Js object.
		stringified := Jerry_json_stringfy(input)
		// if there is no error
		if !Jerry_value_is_error(stringified) {
			jsonStr := jsStrToGoStr(stringified)
			data := make(map[string]interface{}, 0)
			err := json.Unmarshal([]byte(jsonStr), &data)

			if err == nil {
				if data["TYPENAME"] != nil {
					relfectValue := Utility.MakeInstance(data["TYPENAME"].(string), data)
					value = relfectValue.Interface()
				} else {
					// Here map[string]interface{} will be use.
					value = data
				}

			} else {
				return nil, err
			}
		}
	}

	return value, nil
}

/**
 * That function is use to give access to a native golang function
 * from inside JerryScript.
 * name The name of the go function to be call.
 *
 */
//export CallFunction
func CallFunction(name string) {
	//
	log.Println("---> call funt", name)
}

/**
 * Evaluate a script.
 * script Contain the code to run.
 * variables Contain the list of variable to set on the global context before
 * running the script.
 */
func (self *Engine) EvalScript(script string, variables []Variable) (interface{}, error) {
	// Here the values are put on the global contex before use in the function.
	for i := 0; i < len(variables); i++ {
		self.SetGlobalVariable(variables[i].Name, variables[i].Value)
	}

	// Now I will evaluate the function...
	ret := Jerry_eval(NewUint8FromString(script), int64(len(script)), false)

	// Convert Js value to Go value.
	value, err := jsToGo(ret)

	// Free JavaScript value, returned by eval
	Jerry_release_value(ret)
	return value, err
}

func (self *Engine) Clear() {
	for _, fct := range self.functions {
		Jerry_release_value(fct.Obj)
	}
	/* Cleanup the script engine. */
	Jerry_cleanup()

}
