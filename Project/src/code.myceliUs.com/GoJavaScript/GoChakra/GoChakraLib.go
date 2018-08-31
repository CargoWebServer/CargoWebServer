package GoChakra

//#cgo CXXFLAGS: -std=c++11  -Wno-c++11-compat-deprecated-writable-strings -Wno-deprecated-declarations -Wno-unknown-warning-option
//#cgo LDFLAGS: -lChakraCore -lpthread  -lm -lrt -ldl -lstdc++
/*
#include "ChakraCore.h"
#include <stdlib.h>
#include <stdio.h>

// The null pointer value.
#define nullptr 0

*/
import "C"
import (
	"log"
	//	"unsafe"
	"errors"
)

// Global variable.
var (
	// The global object.
	globalObj = getGlobalObject()
)

func getGlobalObject() uintptr {
	var globalObj uintptr

	err := getError(JsGetGlobalObject(&globalObj))
	if err != nil {
		log.Println("Print the error --> ", err)
		// Set the global object to the null value.
		JsGetNullValue(&globalObj)
	}

	return globalObj
}

// Return the go error from the js error code.
func getError(code Enum_SS__JsErrorCode) error {
	var err error
	if code == JsNoError {
		log.Println("---> no error found!")
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
		errors.New("A script was terminated due to a request to suspend a runtime.")
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
