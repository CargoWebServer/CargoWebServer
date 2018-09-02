package GoChakra

//#cgo CXXFLAGS: -std=c++11  -Wno-c++11-compat-deprecated-writable-strings -Wno-deprecated-declarations -Wno-unknown-warning-option
//#cgo LDFLAGS: -lChakraCore -lpthread  -lm -lrt -ldl -lstdc++
/*
#include "ChakraCore.h"
#include <stdlib.h>
#include <stdio.h>

// The null pointer value.
#define nullptr 0

extern const JsRuntimeHandle getRuntime();
extern const JsContextRef getActiveContext();
extern unsigned getCurrentSourceContext();

// String function...
JsErrorCode jsCopyString(
        _In_ JsValueRef value,
        _Out_opt_ char* buffer,
        _In_ size_t bufferSize){

	return JsCopyString( value, buffer, bufferSize, nullptr);
}

*/
import "C"
import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strings"
	"unsafe"

	"code.myceliUs.com/GoJavaScript"
	"code.myceliUs.com/Utility"
)

func getGlobalObject() uintptr {
	var globalObj uintptr
	JsGetGlobalObject(&globalObj)
	return globalObj
}

// Return the go error from the js error code.
func getError(code Enum_SS__JsErrorCode) error {
	var err error
	if code == JsNoError {
		return nil
	} else if code == JsErrorCategoryUsage {
		return errors.New("Category of errors that relates to incorrect usage of the API itself.")
	} else if code == JsErrorInvalidArgument {
		errors.New("An argument to a hosting API was invalid.")
	} else if code == JsErrorNullArgument {
		errors.New("An argument to a hosting API was null in a context where null is not allowed.")
	} else if code == JsErrorNoCurrentContext {
		errors.New("The hosting API requires that a context be current, but there is no current context.")
	} else if code == JsErrorInExceptionState {
		errors.New("The engine is in an exception state and no APIs can be called until the exception is cleared.")
	} else if code == JsErrorNotImplemented {
		errors.New("A hosting API is not yet implemented.")
	} else if code == JsErrorWrongThread {
		errors.New("A hosting API was called on the wrong thread.")
	} else if code == JsErrorRuntimeInUse {
		errors.New("A runtime that is still in use cannot be disposed.")
	} else if code == JsErrorBadSerializedScript {
		errors.New("A bad serialized script was used, or the serialized script was serialized by a different version of the Chakra engine.")
	} else if code == JsErrorInDisabledState {
		errors.New("The runtime is in a disabled state.")
	} else if code == JsErrorCannotDisableExecution {
		errors.New("Runtime does not support reliable script interruption.")
	} else if code == JsErrorHeapEnumInProgress {
		errors.New("A heap enumeration is currently underway in the script context.")
	} else if code == JsErrorArgumentNotObject {
		errors.New("A hosting API that operates on object values was called with a non-object value.")
	} else if code == JsErrorInProfileCallback {
		errors.New("A script context is in the middle of a profile callback.")
	} else if code == JsErrorInThreadServiceCallback {
		errors.New("A thread service callback is currently underway.")
	} else if code == JsErrorCannotSerializeDebugScript {
		errors.New("Scripts cannot be serialized in debug contexts.")
	} else if code == JsErrorAlreadyDebuggingContext {
		errors.New("The context cannot be put into a debug state because it is already in a debug state.")
	} else if code == JsErrorAlreadyProfilingContext {
		errors.New("The context cannot start profiling because it is already profiling.")
	} else if code == JsErrorIdleNotEnabled {
		errors.New("Idle notification given when the host did not enable idle processing.")
	} else if code == JsCannotSetProjectionEnqueueCallback {
		errors.New("The context did not accept the enqueue callback.")
	} else if code == JsErrorCannotStartProjection {
		errors.New("Failed to start projection.")
	} else if code == JsErrorInObjectBeforeCollectCallback {
		errors.New("The operation is not supported in an object before collect callback.")
	} else if code == JsErrorObjectNotInspectable {
		errors.New("Object cannot be unwrapped to IInspectable pointer.")
	} else if code == JsErrorPropertyNotSymbol {
		errors.New("A hosting API that operates on symbol property ids but was called with a non-symbol property id. The error code is returned by JsGetSymbolFromPropertyId if the function is called with non-symbol property id.")
	} else if code == JsErrorPropertyNotString {
		errors.New("A hosting API that operates on string property ids but was called with a non-string property id. The error code is returned by existing JsGetPropertyNamefromId if the function is called with non-string property id.")
	} else if code == JsErrorInvalidContext {
		errors.New("Module evaluation is called in wrong context.")
	} else if code == JsInvalidModuleHostInfoKind {
		errors.New("Module evaluation is called in wrong context.")
	} else if code == JsErrorModuleParsed {
		errors.New("Module was parsed already when JsParseModuleSource is called.")
	} else if code == JsNoWeakRefRequired {
		errors.New("Argument passed to JsCreateWeakReference is a primitive that is not managed by the GC. No weak reference is required, the value will never be collected.")
	} else if code == JsErrorPromisePending {
		errors.New("The Promise object is still in the pending state.")
	} else if code == JsErrorCategoryEngine {
		errors.New("Category of errors that relates to errors occurring within the engine itself.")
	} else if code == JsErrorOutOfMemory {
		errors.New("The Chakra engine has run out of memory.")
	} else if code == JsErrorBadFPUState {
		errors.New("The Chakra engine failed to set the Floating Point Unit state.")
	} else if code == JsErrorCategoryScript {
		errors.New("Category of errors that relates to errors in a script.")
	} else if code == JsErrorScriptException {
		errors.New("A JavaScript exception occurred while running a script.")
	} else if code == JsErrorScriptCompile {
		errors.New("JavaScript failed to compile.")
	} else if code == JsErrorScriptTerminated {
		errors.New("A script was terminated due to a request to 	jsonStr = `{}`suspend a runtime.")
	} else if code == JsErrorScriptEvalDisabled {
		errors.New("A script was terminated because it tried to use eval or function and eval was disabled.")
	} else if code == JsErrorCategoryFatal {
		errors.New("Category of errors that are fatal and signify failure of the engine.")
	} else if code == JsErrorFatal {
		errors.New("A fatal error in the engine has occurred.")
	} else if code == JsErrorWrongRuntime {
		errors.New("A hosting API was called with object created on different javascript runtime.")
	} else if code == JsErrorCategoryDiagError {
		errors.New("Category of errors that are related to failures during diagnostic operations.")
	} else if code == JsErrorDiagAlreadyInDebugMode {
		errors.New("The object for which the debugging API was called was not found.")
	} else if code == JsErrorDiagNotInDebugMode {
		errors.New("The debugging API can only be called when VM is in debug mode.")
	} else if code == JsErrorDiagNotAtBreak {
		errors.New("The debugging API can only be called when VM is at a break.")
	} else if code == JsErrorDiagInvalidHandle {
		errors.New("Debugging API was called with an invalid handle.")
	} else if code == JsErrorDiagObjectNotFound {
		errors.New("The object for which the debugging API was called was not found")
	} else if code == JsErrorDiagUnableToPerformAction {
		errors.New("VM was unable to perform the request action.")
	}
	return err
}

/**
 * Create a new value and set it finalyse methode.
 */
func NewValue(ptr uintptr) *GoJavaScript.Value {
	// Here I will create a new GoJavaScript value.
	v := new(GoJavaScript.Value)
	v.TYPENAME = "GoJavaScript.Value"
	var err error

	// Export the value.
	v.Val, err = jsToGo(ptr)
	if err != nil {
		log.Println("---> error: ", err)
		return nil
	}
	return v
}

// Retreive an object by it uuid as a global object property.
func getJsObjectByUuid(uuid string) uintptr {
	obj := GoJavaScript.GetCache().GetJsObject(uuid)
	if obj != nil {
		return obj.(uintptr)
	}
	log.Println("---> object ", uuid, "is undefined!")
	// The property is undefined.
	return uintptr(0)
}

// Convert a Go value to a JS value.
func goToJs(value interface{}) uintptr {
	var val uintptr
	var typeOf = reflect.TypeOf(value)
	//log.Println("---> go to js ", value, typeOf.String(), typeOf.Kind())
	if typeOf.Kind() == reflect.String {
		// Create a new str variable here.
		var str = value.(string)
		err := getError(JsCreateString(str, int64(len(str)), &val))
		if err != nil {
			log.Println("--> fail to create string ", value)
		}
	} else if typeOf.Kind() == reflect.Bool {
		// Boolean value
		err := getError(JsBoolToBoolean(value.(bool), &val))
		if err != nil {
			log.Println("--> fail to create bool value from ", value)
		}
	} else if typeOf.Kind() == reflect.Int {
		err := getError(JsIntToNumber(value.(int), &val))
		if err != nil {
			log.Println("--> fail to create number int value from ", value)
		}
	} else if typeOf.Kind() == reflect.Int8 {
		err := getError(JsIntToNumber(int(value.(int8)), &val))
		if err != nil {
			log.Println("--> fail to create number int8 value from ", value)
		}
	} else if typeOf.Kind() == reflect.Int16 {
		err := getError(JsIntToNumber(int(value.(int16)), &val))
		if err != nil {
			log.Println("--> fail to create create number int16 value from ", value)
		}
	} else if typeOf.Kind() == reflect.Int32 {
		err := getError(JsIntToNumber(int(value.(int32)), &val))
		if err != nil {
			log.Println("--> fail to create number int32 value from ", value)
		}
	} else if typeOf.Kind() == reflect.Int64 {
		err := getError(JsIntToNumber(int(value.(int64)), &val))
		if err != nil {
			log.Println("--> fail to create number int64 value from ", value)
		}
	} else if typeOf.Kind() == reflect.Uint {
		err := getError(JsIntToNumber(value.(int), &val))
		if err != nil {
			log.Println("--> fail to create number uint value from ", value)
		}
	} else if typeOf.Kind() == reflect.Uint8 {
		err := getError(JsIntToNumber(int(value.(uint8)), &val))
		if err != nil {
			log.Println("--> fail to create number uint8 value from ", value)
		}
	} else if typeOf.Kind() == reflect.Uint16 {
		err := getError(JsIntToNumber(int(value.(uint16)), &val))
		if err != nil {
			log.Println("--> fail to create create number uint16 value from ", value)
		}
	} else if typeOf.Kind() == reflect.Uint32 {
		err := getError(JsIntToNumber(int(value.(uint32)), &val))
		if err != nil {
			log.Println("--> fail to create number uint32 value from ", value)
		}
	} else if typeOf.Kind() == reflect.Uint64 {
		err := getError(JsIntToNumber(int(value.(uint64)), &val))
		if err != nil {
			log.Println("--> fail to create number uint64 value from ", value)
		}
	} else if typeOf.Kind() == reflect.Float32 {
		err := getError(JsDoubleToNumber(float64(value.(float32)), &val))
		if err != nil {
			log.Println("--> fail to create number float32 value from ", value)
		}
	} else if typeOf.Kind() == reflect.Float64 {
		err := getError(JsDoubleToNumber(value.(float64), &val))
		if err != nil {
			log.Println("--> fail to create number float64 value from ", value)
		}
	} else if typeOf.Kind() == reflect.Slice {
		// So here I will create a array and put value in it.
		s := reflect.ValueOf(value)

		err := getError(JsCreateArray(uint(s.Len()), &val))
		if err != nil {
			log.Println("--> fail to create number array value", value)
		}

		for i := 0; i < s.Len(); i++ {
			// I will get element by index.
			e := goToJs(s.Index(i).Interface())
			// Set the element at given index.
			var index uintptr
			JsIntToNumber(i, &index)
			JsSetIndexedProperty(val, index, e)
		}
	} else if typeOf.String() == "uintptr" {
		// already a js value
		return value.(uintptr)
	} else if typeOf.String() == "GoJavaScript.ObjectRef" {
		// I got a Js object reference.
		uuid := value.(GoJavaScript.ObjectRef).UUID
		val = getJsObjectByUuid(uuid)
		if val == uintptr(0) {
			// If the object is not in the cache...
			log.Panicln("----> object ", uuid, " dosent exist anymore!")
		}
	} else if typeOf.String() == "*GoJavaScript.ObjectRef" {
		// I got a Js object reference.
		uuid := value.(*GoJavaScript.ObjectRef).UUID
		val = getJsObjectByUuid(uuid)
		if val == uintptr(0) {
			// If the object is not in the cache...
			log.Panicln("----> object ", uuid, " dosent exist anymore!")
		}
	} else if typeOf.String() == "map[string]interface {}" {
		// In that case I will create a object from the value found in the map
		// and return it as prop value.
		data, err := json.Marshal(value)
		if err == nil {
			if value.(map[string]interface{})["TYPENAME"] != nil {
				ref, err := GoJavaScript.CallGoFunction("Client", "CreateGoObject", string(data))
				if err == nil {
					// In that case an object exist in the case...
					val = GoJavaScript.GetCache().GetJsObject(ref.(*GoJavaScript.ObjectRef).UUID).(uintptr)
				} else {
					log.Println("--> fail to Create Go object ", string(data), err)
				}
			} else {
				// Not a registered type...
				//propValue = jerry_value_t_To_uint32_t(C.json_parse(cstr, C.size_t(len(string(data)))))
				val, err = JsJsonParse(string(data))
				if err != nil {
					log.Println("--> ", err)
				}
			}
		}
	} else if typeOf.Kind() == reflect.Struct {
		val, err := Utility.ToMap(value)
		if err == nil {
			return goToJs(val)
		}
	} else {
		log.Panicln("---> type not found ", value, typeOf.String())
	}

	return val
}

// Shortcut function to get Js object property by it given name.
func JsGetObjectPropertyByName(obj uintptr, name string) (uintptr, error) {
	if !JsObjectHasPropertyByName(obj, name) {
		return uintptr(0), errors.New("no property found with name " + name)
	}
	var id uintptr
	JsCreatePropertyId(name, int64(len(name)), &id)

	var p uintptr
	err := getError(JsGetProperty(obj, id, &p))

	return p, err
}

func JsSetObjectPropertyByName(obj uintptr, name string, value interface{}) error {
	var id uintptr
	JsCreatePropertyId(name, int64(len(name)), &id)
	val := goToJs(value)
	return getError(JsSetProperty(obj, id, val, true))
}

func JsObjectHasPropertyByName(obj uintptr, name string) bool {
	var id uintptr
	JsCreatePropertyId(name, int64(len(name)), &id)

	hasProperty := false
	JsHasProperty(obj, id, &hasProperty)

	return hasProperty
}

func JsDeleteObjectPropertyByName(obj uintptr, name string) error {
	var id uintptr
	JsCreatePropertyId(name, int64(len(name)), &id)
	var ret uintptr

	err := getError(JsDeleteProperty(obj, id, true, &ret))

	return err
}

// Test if a value is null
func JsIsNull(val uintptr) bool {
	var type_ Enum_SS__JsValueType
	JsGetValueType(val, &type_)
	if type_ == JsNull {
		return true
	}
	return false
}

// Test if a value is undefined
func JsIsUndefined(val uintptr) bool {
	var type_ Enum_SS__JsValueType
	JsGetValueType(val, &type_)
	if type_ == JsUndefined {
		return true
	}
	return false
}

func JsIsFunction(val uintptr) bool {
	var type_ Enum_SS__JsValueType
	JsGetValueType(val, &type_)
	if type_ == JsFunction {
		return true
	}
	return false
}

func JsIsObject(val uintptr) bool {
	var type_ Enum_SS__JsValueType
	JsGetValueType(val, &type_)
	if type_ == JsObject {
		return true
	}
	return false
}

// Call an js function over an object, or the global object if obj is null
func callJsFunction(obj uintptr, name string, params []interface{}) (uintptr, error) {
	if obj == 0 {
		obj = getGlobalObject()
	}

	// get the function to be call.
	fct, err := JsGetObjectPropertyByName(obj, name)

	if !JsIsFunction(fct) {
		return uintptr(0), errors.New(name + " is not a function!")
	}

	// if the function dosent exist...
	if err != nil {
		return uintptr(0), err
	}

	var null uintptr
	JsGetNullValue(&null)

	// Now I will set the arguments...
	args := make([]uintptr, len(params)+1)

	// Set the this pointer.
	args[0] = obj

	for i := 0; i < len(params); i++ {
		if params[i] == nil {
			args[i+1] = null
		} else {
			p := goToJs(params[i])
			args[i+1] = p
		}
	}

	var ret uintptr
	err = getError(JsCallFunction(fct, (*uintptr)(unsafe.Pointer(&args[0])), uint16(len(args)), &ret))
	return ret, err
}

/**
 * Append a Js function to a given object.
 */
func appendJsFunction(object uintptr, name string, src string) error {

	// eval the script.
	_, err := evalScript(src, name)

	// in that case the function must be set as object function.
	if object != 0 && err == nil {
		fct, err := JsGetObjectPropertyByName(getGlobalObject(), name)
		if err == nil {
			// Set the function on the object.
			JsSetObjectPropertyByName(object, name, fct)

			// remove it from the global object.
			if object != getGlobalObject() {
				JsDeleteObjectPropertyByName(getGlobalObject(), name)
			}
		} else {
			return errors.New("no function found with name " + name)
		}
	}
	return err
}

// Eval a given script string.
func evalScript(script string, name string) (uintptr, error) {
	var ret uintptr
	var currentSourceContext = C.getCurrentSourceContext()
	err := getError(JsRun(goToJs(script), NewUintptr((uintptr)(unsafe.Pointer(&currentSourceContext))), goToJs(name), JsParseScriptAttributeNone, &ret))
	return ret, err
}

/**
 * Serialyse to json
 */
func JsJsonStringify(obj uintptr) (string, error) {
	// Get the JSON object from the global context.
	JSON, _ := JsGetObjectPropertyByName(getGlobalObject(), "JSON")
	val, err := callJsFunction(JSON, "stringify", []interface{}{obj})
	if err != nil {
		return "", err
	}

	val_, err := jsToGo(val)
	if err != nil {
		return "", err
	}

	return val_.(string), nil
}

/**
 * Create an object from json string.
 */
func JsJsonParse(jsonStr string) (uintptr, error) {

	// Get the JSON object from the global context.
	JSON, _ := JsGetObjectPropertyByName(getGlobalObject(), "JSON")

	return callJsFunction(JSON, "parse", []interface{}{jsonStr})
}

// Convert a Js value to Go value
func jsToGo(input uintptr) (interface{}, error) {

	// The js value type.
	var inputType Enum_SS__JsValueType

	err := getError(JsGetValueType(input, &inputType))
	if err != nil {
		return nil, err
	}

	// The input is a string.
	if inputType == JsNull {
		log.Println("---> the input value is nil!")
	} else if inputType == JsUndefined {
		log.Println("---> the input value is undefined!")
	} else if inputType == JsString {
		var size int
		JsGetStringLength(input, &size)

		// Create the memory that will old the value.
		buffer := (*C.char)(unsafe.Pointer(C.malloc(C.ulong(size))))

		// Put the value in the buffer.
		C.jsCopyString(C.JsRef(input), buffer, C.size_t(size))

		// Copy the value to a string.
		str := C.GoStringN(buffer, C.int(size))

		// free the buffer.
		C.free(unsafe.Pointer(buffer))

		return str, nil
	} else if inputType == JsBoolean {
		var val bool
		err = getError(JsBooleanToBool(input, &val))
		if err == nil {
			return val, nil
		}
	} else if inputType == JsNumber {
		// all number will be converted to float64
		var val float64
		err = getError(JsNumberToDouble(input, &val))
		if err == nil {
			return val, nil
		}
	} else if inputType == JsArray {
		// retreive the object property.
		l, _ := JsGetObjectPropertyByName(input, "length")
		var length int
		JsNumberToInt(l, &length)

		// So here I got a array without type so I will get it property by index
		// and interpret each result.
		val := make([]interface{}, 0)
		for i := 0; i < length; i++ {
			var index uintptr
			JsIntToNumber(i, &index)
			var e uintptr
			err := getError(JsGetIndexedProperty(input, index, &e))
			if err == nil {
				v, err := jsToGo(e)
				if err == nil {
					val = append(val, v)
				}
			}
		}
		return val, nil
	} else if inputType == JsArrayBuffer {
		/** not implemented **/
	} else if inputType == JsDataView {
		/** not implemented **/
	} else if inputType == JsTypedArray {
		/** not implemented **/
	} else if inputType == JsObject {
		// The go object will be a copy of the Js object.
		if JsObjectHasPropertyByName(input, "uuid_") {
			uuid_, _ := JsGetObjectPropertyByName(input, "uuid_")
			// Get the uuid string.
			uuid, _ := jsToGo(uuid_)
			// Return and object reference.
			val := GoJavaScript.NewObjectRef(uuid.(string))
			return val, nil
		} else {
			jsonStr, err := JsJsonStringify(input)
			// if there is no error
			if err == nil {
				if strings.Index(jsonStr, "TYPENAME") != -1 {
					// So here I will create a remote action and tell the client to
					// create a Go object from jsonStr. The object will be set by
					// the client on the server.
					return GoJavaScript.CallGoFunction("Client", "CreateGoObject", jsonStr)
				} else {
					// In that case I will use a map to instantited the object.
					val := make(map[string]interface{})
					err := json.Unmarshal([]byte(jsonStr), &val)
					return val, err
				}
			} else {
				// Continue any way with nil object instead of an error...
				return nil, err
			}
		}
	}

	// error found in convertion.
	return nil, err
}

// The Uint32 Type represent a 32 bit char.
type Uintptr struct {
	// The pointer that old the data.
	ptr unsafe.Pointer
}

func NewUintptr(ptr uintptr) Uintptr {
	var val Uintptr
	val.ptr = unsafe.Pointer(&ptr)
	return val
}

/**
 * Free the values.
 */
func (self Uintptr) Free() {
	C.free(unsafe.Pointer(self.ptr))
}

/**
 * Access the undelying memeory values pointer.
 */
func (self Uintptr) Swigcptr() uintptr {
	return uintptr(self.ptr)
}
