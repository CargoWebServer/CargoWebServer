package GoChakra

import (
	"log"

	"code.myceliUs.com/GoJavaScript"
)

/*
#include "ChakraCore.h"
#include <stdlib.h>
#include <stdio.h>

// The null pointer value.
#define nullptr 0

// Only one runtime and context per process.
JsRuntimeHandle runtime;
JsContextRef context;
unsigned currentSourceContext = 0;

void initRuntime(){

	// Create the runtime.
	JsCreateRuntime(JsRuntimeAttributeNone, nullptr, &runtime);

	// create the context
	JsCreateContext(runtime, &context);

	// set the context.
	JsSetCurrentContext(context);
}

// Return the runtime
const JsRuntimeHandle getRuntime(){
	return runtime;
}

// Return the active context.
const JsContextRef getActiveContext(){
	return context;
}

unsigned getCurrentSourceContext(){
	return currentSourceContext++;
}

*/
import "C"

import "errors"

/**
 * The Chacra JS Runtime.
 */
type Runtime struct {
}

/**
 * Init and start the Runtime.
 * port The port to communicate with the Runtime, or the debbuger.
 */
func (self *Runtime) Start() {
	// Create the runtime and the context.
	C.initRuntime()
}

func (self *Runtime) DisplayContext(line int) {
	var context uintptr
	JsGetCurrentContext(&context)

	log.Println("--> context ", line, " ", context, getGlobalObject())

}

/////////////////// Global variables //////////////////////

/**
 * Set a variable on the global context.
 * name The name of the variable in the context.
 * value The value of the variable, can be a string, a number,
 */
func (self *Runtime) SetGlobalVariable(name string, value interface{}) {

	// Now I will set the value as a global object property.
	var propId uintptr

	// Create the property id with a given name.
	JsCreatePropertyId(name, int64(len(name)), &propId)

	// Set the property of the global object.
	JsSetProperty(getGlobalObject(), propId, goToJs(value), true)

	hasProperty := false
	JsHasProperty(getGlobalObject(), propId, &hasProperty)
	if !hasProperty {
		log.Println("---> fail to create global variable: ", name)
	}
}

/**
 * Return a variable define in the global object.
 */
func (self *Runtime) GetGlobalVariable(name string) (GoJavaScript.Value, error) {
	var val GoJavaScript.Value
	var jsVal uintptr
	var err error

	var propId uintptr
	JsCreatePropertyId(name, int64(len(name)), &propId)

	// I will retreive the property...
	err = getError(JsGetProperty(getGlobalObject(), propId, &jsVal))
	if err != nil {
		log.Println("---> err ", err)
	}

	val = *NewValue(jsVal)

	return val, err
}

/////////////////// Objects //////////////////////
/**
 * Create JavaScript object with given uuid. If name is given the object will be
 * set a global object property.
 */
func (self *Runtime) CreateObject(uuid string, name string) {

	// Create the object JS object.
	var obj uintptr

	// Set the uuid property.
	JsSetObjectPropertyByName(obj, "uuid_", uuid)

	// keep the object in the global namespace.
	// set is uuid as global object property
	GoJavaScript.GetCache().SetJsObject(uuid, obj)

	if len(name) > 0 {
		// set is name as global object property
		JsSetObjectPropertyByName(getGlobalObject(), name, obj)
	}

	// Set native object to the object.
	// C.create_native_object(C.CString(uuid), uint32_t_To_Jerry_value_t(obj))
}

/**
 * Set an object property.
 * uuid The object reference.
 * name The name of the property to set
 * value The value of the property
 */
func (self *Runtime) SetObjectProperty(uuid string, name string, value interface{}) error {
	// Get the object from the cache
	obj := getJsObjectByUuid(uuid)
	if !JsIsUndefined(obj) {
		// Set the property value.
		err := JsSetObjectPropertyByName(obj, name, value)
		if err != nil {
			return err
		}
	} else {
		err := errors.New("311 Object " + uuid + " dosent exist!")
		log.Println("---> ", err)
		return err
	}

	return nil
}

/**
 * That function is use to get Js object property
 */
func (self *Runtime) GetObjectProperty(uuid string, name string) (GoJavaScript.Value, error) {
	var value GoJavaScript.Value
	// I will get the object reference from the cache.
	obj := getJsObjectByUuid(uuid)
	if !JsIsUndefined(obj) {
		// Return the value from an object.
		property, err := JsGetObjectPropertyByName(obj, name)
		return *NewValue(property), err

	} else {
		err := errors.New("302 Object " + uuid + " dosent exist!")
		log.Println("---> ", err)
		return value, err
	}
}

/**
 * Create an empty array of a given size and set it as object property.
 */
func (self *Runtime) CreateObjectArray(uuid string, name string, size uint32) error {
	obj := getJsObjectByUuid(uuid)
	if JsIsUndefined(obj) {
		err := errors.New("340 Object " + uuid + " dosent exist!")
		log.Println("---> ", err)
		return err
	}

	var arr_ uintptr
	err := getError(JsCreateArray(uint(size), &arr_))
	if err != nil {
		return err
	}

	// set it property.
	JsSetObjectPropertyByName(obj, name, arr_)

	return nil
}

/**
 * Set an object property.
 * uuid The object reference.
 * name The name of the property to set
 * index The index of the object in the array
 * value The value of the property
 */
func (self *Runtime) SetObjectPropertyAtIndex(uuid string, name string, i uint32, value interface{}) {
	// Get the object from the cache
	obj := getJsObjectByUuid(uuid)
	if !JsIsUndefined(obj) {
		// I will retreive the array...
		arr, err := JsGetObjectPropertyByName(obj, name)
		if err == nil {
			// So here I will set it property.
			var index uintptr
			JsIntToNumber(int(i), &index)
			JsSetIndexedProperty(arr, index, goToJs(value))
		}
	}
}

/**
 * That function is use to get Js obeject property
 */
func (self *Runtime) GetObjectPropertyAtIndex(uuid string, name string, i uint32) (GoJavaScript.Value, error) {
	// I will get the object reference from the cache.
	var value GoJavaScript.Value
	obj := getJsObjectByUuid(uuid)
	if !JsIsUndefined(obj) {
		// Return the value from an object.
		arr, err := JsGetObjectPropertyByName(obj, name)

		if err != nil {
			return value, err
		}

		// Here I will get the value.
		var e uintptr
		var index uintptr
		JsIntToNumber(int(i), &index)
		JsGetIndexedProperty(arr, index, &e)

		return *NewValue(e), nil

	} else {
		return value, errors.New("395 Object " + uuid + " dosent exist!")
	}
}

/**
 * set object methode.
 */
func (self *Runtime) SetGoObjectMethod(uuid, name string) error {
	obj := getJsObjectByUuid(uuid)
	var err error
	if !JsIsUndefined(obj) {
		err = errors.New("389 Object " + uuid + " dosent exist!")
		log.Println(err)
		return err
	}

	if JsIsObject(obj) {

		// TODO set native method.
		// C.setGoMethod(C.CString(name), uint32_t_To_Jerry_value_t(obj))
		return nil
	}

	err = errors.New("397 " + uuid + " is not an object!")
	log.Println(err)
	return err
}

func (self *Runtime) SetJsObjectMethod(uuid, name string, src string) error {
	obj := getJsObjectByUuid(uuid)
	if JsIsUndefined(obj) {
		return errors.New("418 Object " + uuid + " dosent exist!")
	}
	// In that case I want to associate a js function to the object.
	err := appendJsFunction(obj, name, src)
	return err
}

/**
 * Call object methode.
 */
func (self *Runtime) CallObjectMethod(uuid string, name string, params ...interface{}) (GoJavaScript.Value, error) {
	var value GoJavaScript.Value
	obj := getJsObjectByUuid(uuid)
	if JsIsUndefined(obj) {
		return value, errors.New("432 Object " + uuid + " dosent exist!")
	}
	val, err := callJsFunction(obj, name, params)
	return *NewValue(val), err
}

/////////////////// Functions //////////////////////

/**
 * Register a go function in JS
 */
func (self *Runtime) RegisterGoFunction(name string) {
	// TODO implement it
	// C.setGoMethod(cstr, uint32_t_To_Jerry_value_t(getGlobalObject()))
}

/**
 * Parse and set a function in the Javascript.
 * name The name of the function (the function will be keep in the Runtime for it
 *      lifetime.
 * args The argument name for that function.
 * src  The body of the function
 * options Can be JERRY_PARSE_NO_OPTS or JERRY_PARSE_STRICT_MODE
 */
func (self *Runtime) RegisterJsFunction(name string, src string) error {
	err := appendJsFunction(uintptr(0), name, src)
	return err
}

/**
 * Call a Javascript function. The function must exist...
 */
func (self *Runtime) CallFunction(name string, params []interface{}) (GoJavaScript.Value, error) {
	val, err := callJsFunction(getGlobalObject(), name, params)
	// Call function on the global object here.
	return *NewValue(val), err
}

/**
 * Evaluate a script.
 * script Contain the code to run.
 * variables Contain the list of variable to set on the global context before
 * running the script.
 */
func (self *Runtime) EvalScript(script string, variables []interface{}) (GoJavaScript.Value, error) {

	// Here the values are put on the global contex before use in the function.
	for i := 0; i < len(variables); i++ {
		self.SetGlobalVariable(variables[i].(*GoJavaScript.Variable).Name, variables[i].(*GoJavaScript.Variable).Value)
	}

	val, err := evalScript(script, "unknow")

	return *NewValue(val), err
}

func (self *Runtime) Clear() {
	/* Cleanup the script Runtime. */
}
