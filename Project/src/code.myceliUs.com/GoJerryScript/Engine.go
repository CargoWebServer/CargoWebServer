package GoJerryScript

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "jerryscript.h"
#include "jerryscript-ext/handler.h"
#include "jerryscript-debugger.h"

extern  jerry_value_t
handler (const jerry_value_t,
         const jerry_value_t,
         const jerry_value_t[],
         const jerry_length_t);

// Define a new function handler.
void setGoFct(const char* name){
  // Function will have there own handler.
  jerry_value_t fct_handler = jerry_create_external_function (handler);
  jerry_value_t glob_obj = jerry_get_global_object ();
  jerry_value_t prop_name = jerry_create_string ((const jerry_char_t *) name);

  // keep the name in the function object itself, so it will be use at runtime
  // by the handler to know witch function to dynamicaly call.
  jerry_value_t prop_name_name = jerry_create_string ((const jerry_char_t *) "name");
  jerry_release_value (jerry_set_property (fct_handler, prop_name_name, prop_name));

  // set property and release the return value without any check
  jerry_release_value (jerry_set_property (glob_obj, prop_name, fct_handler));
  jerry_release_value (prop_name);
  jerry_release_value (prop_name_name);
  jerry_release_value (glob_obj);
  jerry_release_value (fct_handler);
}

typedef jerry_value_t* jerry_value_p;

// Define a new function handler.
void setGoMethod(const char* name, jerry_value_t obj){

  jerry_value_t fct_handler = jerry_create_external_function (handler);
  jerry_value_t prop_name = jerry_create_string ((const jerry_char_t *) name);

  // keep the name in the function object itself, so it will be use at runtime
  // by the handler to know witch function to dynamicaly call.
  jerry_value_t prop_name_name = jerry_create_string ((const jerry_char_t *) "name");
  jerry_release_value (jerry_set_property (fct_handler, prop_name_name, prop_name));

  // set property and release the return value without any check
  jerry_release_value (jerry_set_property (obj, prop_name, fct_handler));
  jerry_release_value (prop_name);
  jerry_release_value (prop_name_name);
  jerry_release_value (fct_handler);
}

// Call a JS function.
jerry_value_t
call_function ( const jerry_value_t func_obj_val,
                const jerry_value_t this_val,
                const jerry_value_p args_p,
                int32_t args_count){

	// evaluate the result
	jerry_value_t ret = jerry_call_function ( func_obj_val, this_val, args_p, args_count );
	return ret;
}

// Eval Script.
jerry_value_t
eval (const char *source_p,
            size_t source_size,
            bool is_strict){
	return jerry_eval (source_p, source_size, is_strict);
}

// Create the error message.
jerry_value_t
create_error (jerry_error_t error_type,
                    const char *message_p){
	// Create and return the error message.
	return jerry_create_error (error_type, message_p);

}

// The string helper function.

// Create a string and return it value.
jerry_value_t create_string (const char *str_p){
	return jerry_create_string_from_utf8 (str_p);
}

jerry_size_t
get_string_size (const jerry_value_t value){
	return jerry_get_utf8_string_size(value);
}

jerry_size_t
string_to_utf8_char_buffer (const jerry_value_t value, char *buffer_p, size_t buffer_size){

	// Here I will set the string value inside the buffer and return
	// the size of the written data.
	return jerry_string_to_utf8_char_buffer (value, buffer_p, buffer_size);
}

//extern jerry_value_t create_native_object(const char* uuid);
// Use to bind Js object and Go object lifecycle
typedef  struct {
	const char* uuid;
} object_reference_t;

// Return the object reference.
const char* get_object_reference_uuid(uintptr_t ref){
	return ((object_reference_t*)ref)->uuid;
}

void delete_object_reference(uintptr_t ref){
	free((char*)((object_reference_t*)ref)->uuid);
	free((object_reference_t*)ref);
}

// This is the destructor.
extern void object_native_free_callback(void* native_p);


// NOTE: The address (!) of type_info acts as a way to uniquely "identify" the
// C type `native_obj_t *`.
static const jerry_object_native_info_t native_obj_type_info =
{
  .free_cb = object_native_free_callback
};

jerry_value_t create_native_object(const char* uuid){
	// Create the js object that will be back by the
	// native object.
	jerry_value_t object = jerry_create_object();

	// Here I will create the native object reference.
	object_reference_t *native_obj = malloc(sizeof(*native_obj));
	native_obj->uuid = uuid;

	// Set the native pointer.
	jerry_set_object_native_pointer (object, native_obj, &native_obj_type_info);

	// Set the uuid property of the newly created object.
	jerry_value_t prop_uuid = jerry_create_string ((const jerry_char_t *) uuid);
 	jerry_value_t prop_name = jerry_create_string ((const jerry_char_t *) "uuid_");

	// Set the uuid property.
  	jerry_release_value (jerry_set_property (object, prop_name, prop_uuid));
	jerry_release_value (prop_uuid);
	jerry_release_value (prop_name);

	return object;
}

// Simplifyed array function.
jerry_value_t
create_array (uint32_t size){
	return jerry_create_array (size);
}

jerry_value_t
set_property_by_index (const jerry_value_t obj_val,
                             uint32_t index,
                             const jerry_value_t value_to_set){
	return jerry_set_property_by_index (obj_val, index, value_to_set);
}


jerry_value_t
get_property_by_index (const jerry_value_t obj_val,
                             uint32_t index){
	return jerry_get_property_by_index (obj_val, index);
}

uint32_t
get_array_length (const jerry_value_t value){
	return jerry_get_array_length (value);
}
*/
import "C"
import "errors"
import "unsafe"

//import "log"

/**
 * The JerryScript JS engine.
 */
type Engine struct {
	// The debugger port.
	port int
}

// Create new Javascript
func NewEngine(port int, options int) *Engine {
	// The engine.
	engine := new(Engine)
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

/////////////////// Global variables //////////////////////
/**
 * Set a variable on the global context.
 * name The name of the variable in the context.
 * value The value of the variable, can be a string, a number,
 */
func (self *Engine) SetGlobalVariable(name string, value interface{}) {

	// first of all I will initialyse the arguments.
	globalObject := Jerry_get_global_object()
	defer Jerry_release_value(globalObject)

	propName := goToJs(name)
	defer Jerry_release_value(propName)

	// Take a go value and initialyse a Uint32 representation.
	propValue := goToJs(value)
	defer Jerry_release_value(propValue)

	// Set the propertie in the global context..
	Jerry_release_value(Jerry_set_property(globalObject, propName, propValue))
}

/**
 * Return a variable define in the global object.
 */
func (self *Engine) GetGlobalVariable(name string) (Value, error) {
	// first of all I will initialyse the arguments.
	globalObject := Jerry_get_global_object()
	defer Jerry_release_value(globalObject)
	propertyName := goToJs(name)
	defer Jerry_release_value(propertyName)
	property := Jerry_get_property(globalObject, propertyName)

	var value Value

	if Jerry_value_is_error(property) {
		return value, errors.New("No variable found with name " + name)
	}

	value = *NewValue(property)

	return value, nil
}

/////////////////// Objects //////////////////////
/**
 * Create JavaScript object with given uuid. If name is given the object will be
 * set a global object property.
 */
func (self *Engine) CreateObject(uuid string, name string) {

	// Keep the pointer.
	obj := jerry_value_t_To_uint32_t(C.create_native_object(C.CString(uuid)))

	// Set the object in the cache.
	GetCache().setJsObject(uuid, obj)

	if len(name) > 0 {
		globalObj := Jerry_get_global_object()
		defer Jerry_release_value(globalObj)

		propName := goToJs(name)
		defer Jerry_release_value(propName)

		// Set the object as a propety in the global object.
		Jerry_release_value(Jerry_set_property(globalObj, propName, obj))
	}
}

/**
 * Set an object property.
 * uuid The object reference.
 * name The name of the property to set
 * value The value of the property
 */
func (self *Engine) SetObjectProperty(uuid string, name string, value interface{}) {
	// Get the object from the cache
	obj := GetCache().getJsObject(uuid)
	if obj != nil {
		// Set the property value.
		Jerry_release_value(Jerry_set_property(obj, goToJs(name), goToJs(value)))
	}
}

/**
 * That function is use to get Js obeject property
 */
func (self *Engine) GetObjectProperty(uuid string, name string) (Value, error) {
	// I will get the object reference from the cache.
	obj := GetCache().getJsObject(uuid)
	var value Value
	if obj != nil {
		// Return the value from an object.
		ptr := Jerry_get_property(obj, goToJs(name))
		defer Jerry_release_value(ptr)

		if Jerry_value_is_error(ptr) {
			return value, errors.New("no property found with name " + name)
		}

		return *NewValue(ptr), nil

	} else {
		return value, errors.New("Object " + uuid + " dosent exist!")
	}
}

/**
 * Create an empty array of a given size and set it as object property.
 */
func (self *Engine) CreateObjectArray(uuid string, name string, size uint32) error {
	obj := GetCache().getJsObject(uuid)
	if obj == nil {
		return errors.New("Object " + uuid + " dosent exist!")
	}

	arr_ := C.create_array(C.uint32_t(size))
	Jerry_release_value(Jerry_set_property(obj, goToJs(name), jerry_value_t_To_uint32_t(arr_)))
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
	// Get the object from the cache
	obj := GetCache().getJsObject(uuid)

	if obj != nil {
		// I will retreive the array...
		arr := Jerry_get_property(obj, goToJs(name))
		defer Jerry_release_value(arr)
		if Jerry_value_is_array(arr) {
			// So here I will set it property.
			v := goToJs(value)
			r := C.set_property_by_index(uint32_t_To_Jerry_value_t(arr), C.uint32_t(index), uint32_t_To_Jerry_value_t(v))
			// Release the result
			Jerry_release_value(jerry_value_t_To_uint32_t(r))
		}
	}
}

/**
 * That function is use to get Js obeject property
 */
func (self *Engine) GetObjectPropertyAtIndex(uuid string, name string, index uint32) (Value, error) {
	// I will get the object reference from the cache.
	obj := GetCache().getJsObject(uuid)
	var value Value
	if obj != nil {
		// Return the value from an object.
		arr := Jerry_get_property(obj, goToJs(name))
		defer Jerry_release_value(arr)

		if Jerry_value_is_error(arr) {
			return value, errors.New("no property found with name " + name)
		}

		// Here I will get the value.
		e := jerry_value_t_To_uint32_t(C.get_property_by_index(uint32_t_To_Jerry_value_t(arr), C.uint32_t(index)))
		defer Jerry_release_value(e)

		return *NewValue(e), nil

	} else {
		return value, errors.New("Object " + uuid + " dosent exist!")
	}
}

/**
 * set object methode.
 */
func (self *Engine) SetGoObjectMethod(uuid, name string) error {
	obj := GetCache().getJsObject(uuid)
	if obj == nil {
		return errors.New("Object " + uuid + " dosent exist!")
	}
	cs := C.CString(name)
	if Jerry_value_is_object(obj) {
		C.setGoMethod(cs, uint32_t_To_Jerry_value_t(obj))
	}
	defer C.free(unsafe.Pointer(cs))

	return nil
}

func (self *Engine) SetJsObjectMethod(uuid, name string, src string) error {
	obj := GetCache().getJsObject(uuid)
	if obj == nil {
		return errors.New("Object " + uuid + " dosent exist!")
	}
	// In that case I want to associate a js function to the object.
	err := appendJsFunction(obj, name, src)
	return err
}

/**
 * Call object methode.
 */
func (self *Engine) CallObjectMethod(uuid string, name string, params ...interface{}) (Value, error) {
	obj := GetCache().getJsObject(uuid)
	var result Value
	if obj == nil {
		return result, errors.New("Object " + uuid + " dosent exist!")
	}
	return callJsFunction(obj, name, params)
}

/////////////////// Functions //////////////////////

/**
 * Register a go function in JS
 */
func (self *Engine) RegisterGoFunction(name string) {
	cs := C.CString(name)
	C.setGoFct(cs)
	defer C.free(unsafe.Pointer(cs))
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
	err := appendJsFunction(nil, name, src)
	return err
}

/**
 * Call a Javascript function. The function must exist...
 */
func (self *Engine) CallFunction(name string, params []interface{}) (Value, error) {

	globalObject := Jerry_get_global_object()
	defer Jerry_release_value(globalObject)

	// Call function on the global object here.
	return callJsFunction(globalObject, name, params)
}

func (self *Engine) Clear() {
	/* Cleanup the script engine. */
	Jerry_cleanup()
}

/**
 * Evaluate a script.
 * script Contain the code to run.
 * variables Contain the list of variable to set on the global context before
 * running the script.
 */
func (self *Engine) EvalScript(script string, variables Variables) (Value, error) {

	// Here the values are put on the global contex before use in the function.
	for i := 0; i < len(variables); i++ {
		self.SetGlobalVariable(variables[i].Name, variables[i].Value)
	}

	return evalScript(script)
}
