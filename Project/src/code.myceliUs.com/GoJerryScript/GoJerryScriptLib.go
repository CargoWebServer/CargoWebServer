package GoJerryScript

//#cgo LDFLAGS: -L/usr/local/lib -ljerry-core -ljerry-ext -ljerry-libm -ljerry-port-default
/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "jerryscript.h"
#include "jerryscript-ext/handler.h"
#include "jerryscript-debugger.h"

extern uintptr_t _wrap_jerry_create_number_GoJerryScript_8ba631a7412d8d81(double arg1);

static jerry_value_t create_number(float number){
	return jerry_create_number(number);
}
*/
import "C"
import "reflect"
import "unsafe"
import "encoding/binary"
import "math"
import "sync"
import "code.myceliUs.com/Utility"

//import "strconv"
import "log"

// Global variable.
var (
	// Callback function used by dynamic type, it's call when an entity is set.
	// Can be use to store dynamic type in a cache.
	SetEntity func(interface{}) = func(val interface{}) {
		log.Println("---> set entity ", val)
	}

	fctNameMap = make(map[uint32]string, 0)
	mu         sync.Mutex
)

func setFctPointerName(fct uint32, name string) {
	mu.Lock()
	defer mu.Unlock()
	fctNameMap[fct] = name
}

func getFctPointerName(fct uint32) string {
	return fctNameMap[fct]
}

func CallGoFunction(name string, params ...interface{}) (interface{}, error) {
	results, err := Utility.CallFunction(name, params...)
	if err != nil {
		return nil, err
	}
	if len(results) > 0 {
		if len(results) == 1 {
			// One result here...
			switch goError := results[0].Interface().(type) {
			case error:
				return nil, goError
			}
			return results[0].Interface(), nil
		} else {
			// Return the result and the error after.
			results_ := make([]interface{}, 0)
			for i := 0; i < len(results); i++ {
				switch goError := results[i].Interface().(type) {
				case error:
					return nil, goError
				}
				results_ = append(results_, results[i].Interface())
			}
			return results_, nil
		}
	}
	return nil, nil
}

//export handler
func handler(fct C.jerry_value_t, this C.jerry_value_t, args C.uintptr_t, length int) C.jerry_value_t {

	name := getFctPointerName(uint32(fct))
	params := make([]interface{}, 0)

	for i := 0; i < length; i++ {

		val, err := jsToGo((Uint32_t)(SwigcptrUint32_t(C.uintptr_t(args))))
		if err == nil {
			params = append(params, val)

		}
		args += 4
	}

	result, err := CallGoFunction(name, params...)

	if err == nil && result != nil {
		jsVal := goToJs(result)
		val := (*uintptr)(unsafe.Pointer(jsVal.Swigcptr()))
		return C.jerry_value_t(*val)
	}

	// here i will retrun a null value
	null := Jerry_create_null()

	// So here the result value must be dereference before return
	val := (*uintptr)(unsafe.Pointer(null.Swigcptr()))
	return C.jerry_value_t(*val)
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
 * Variable with name and value.
 */
type Variable struct {
	Name  string
	Value interface{}
}

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

func uintptrToBytes(u *uintptr) []byte {
	return (*[sizeOfUintPtr]byte)(unsafe.Pointer(u))[:]
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
