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
	return jerry_call_function ( func_obj_val, this_val, args_p, args_count );
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
	return jerry_create_string (str_p);
}

jerry_size_t
get_string_size (const jerry_value_t value){
	return jerry_get_string_size(value);
}

jerry_size_t
string_to_char_buffer (const jerry_value_t value, char *buffer_p, size_t buffer_size){
	// Here I will set the string value inside the buffer and return
	// the size of the written data.
	return jerry_string_to_char_buffer (value, buffer_p, buffer_size);
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

void create_native_object(const char* uuid, jerry_value_t object){
	// Here I will create the native object reference.
	object_reference_t *native_obj = malloc(sizeof(*native_obj));

	// Set it uuid.
	native_obj->uuid = uuid;

	// Set the native pointer.
	jerry_set_object_native_pointer (object, native_obj, &native_obj_type_info);
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
import "log"
import "unsafe"

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
	// Set the propertie in the global context..
	Jerry_set_object_property(getGlobalObject(), name, value)
}

/**
 * Return a variable define in the global object.
 */
func (self *Engine) GetGlobalVariable(name string) (Value, error) {

	// first of all I will initialyse the arguments.
	property := Jerry_get_object_property(getGlobalObject(), name)

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

	// Create the object JS object.
	obj := Jerry_create_object()
	if !Jerry_value_is_object(obj) {
		log.Panicln("---> fail to create a new object! ", uuid)
	}

	// Set the uuid property.
	Jerry_set_object_property(obj, "uuid_", uuid)

	// keep the object in the global namespace.
	// set is uuid as global object property
	GetCache().setJsObject(uuid, obj)

	if len(name) > 0 {
		// set is name as global object property
		Jerry_set_object_property(getGlobalObject(), name, obj)
	}

	// Set native object to the object.
	C.create_native_object(C.CString(uuid), uint32_t_To_Jerry_value_t(obj))
}

/**
 * Set an object property.
 * uuid The object reference.
 * name The name of the property to set
 * value The value of the property
 */
func (self *Engine) SetObjectProperty(uuid string, name string, value interface{}) error {
	// Get the object from the cache
	obj := getJsObjectByUuid(uuid)
	if !Jerry_value_is_undefined(obj) {
		// Set the property value.
		Jerry_set_object_property(obj, name, value)
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
func (self *Engine) GetObjectProperty(uuid string, name string) (Value, error) {
	var value Value
	// I will get the object reference from the cache.
	obj := getJsObjectByUuid(uuid)
	if !Jerry_value_is_undefined(obj) {
		// Return the value from an object.
		property := Jerry_get_object_property(obj, name)
		defer Jerry_release_value(property)
		if Jerry_value_is_error(property) {
			return value, errors.New("no property found with name " + name)
		}
		return *NewValue(property), nil

	} else {
		err := errors.New("302 Object " + uuid + " dosent exist!")
		log.Println("---> ", err)
		return value, err
	}
}

/**
 * Create an empty array of a given size and set it as object property.
 */
func (self *Engine) CreateObjectArray(uuid string, name string, size uint32) error {
	obj := getJsObjectByUuid(uuid)
	if Jerry_value_is_undefined(obj) {
		err := errors.New("340 Object " + uuid + " dosent exist!")
		log.Println("---> ", err)
		return err
	}

	arr_ := jerry_value_t_To_uint32_t(C.create_array(C.uint32_t(size)))
	if !Jerry_value_is_array(arr_) {
		err := errors.New("321 fail to create array property for Object " + uuid)
		return err
	}

	// set it property.
	Jerry_set_object_property(obj, name, arr_)

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
	obj := getJsObjectByUuid(uuid)
	if !Jerry_value_is_undefined(obj) {
		// I will retreive the array...
		arr := Jerry_get_object_property(obj, name)
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
	var value Value
	obj := getJsObjectByUuid(uuid)
	if !Jerry_value_is_undefined(obj) {
		// Return the value from an object.
		arr := Jerry_get_object_property(obj, name)
		defer Jerry_release_value(arr)

		if Jerry_value_is_error(arr) {
			return value, errors.New("no property found with name " + name)
		}

		// Here I will get the value.
		e := jerry_value_t_To_uint32_t(C.get_property_by_index(uint32_t_To_Jerry_value_t(arr), C.uint32_t(index)))
		defer Jerry_release_value(e)

		return *NewValue(e), nil

	} else {
		return value, errors.New("395 Object " + uuid + " dosent exist!")
	}
}

/**
 * set object methode.
 */
func (self *Engine) SetGoObjectMethod(uuid, name string) error {
	obj := getJsObjectByUuid(uuid)
	var err error
	if Jerry_value_is_undefined(obj) {
		err = errors.New("389 Object " + uuid + " dosent exist!")
		log.Println(err)
		return err
	}

	if Jerry_value_is_object(obj) {

		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(cstr))

		C.setGoMethod(C.CString(name), uint32_t_To_Jerry_value_t(obj))
		return nil
	}

	err = errors.New("397 " + uuid + " is not an object!")
	log.Println(err)
	return err
}

func (self *Engine) SetJsObjectMethod(uuid, name string, src string) error {
	obj := getJsObjectByUuid(uuid)
	if Jerry_value_is_undefined(obj) {
		return errors.New("418 Object " + uuid + " dosent exist!")
	}

	// In that case I want to associate a js function to the object.
	err := appendJsFunction(obj, name, src)
	log.Println("---> 413", err)
	return err
}

/**
 * Call object methode.
 */
func (self *Engine) CallObjectMethod(uuid string, name string, params ...interface{}) (Value, error) {

	var value Value
	obj := getJsObjectByUuid(uuid)
	if Jerry_value_is_undefined(obj) {
		return value, errors.New("432 Object " + uuid + " dosent exist!")
	}
	return callJsFunction(obj, name, params)
}

/////////////////// Functions //////////////////////

/**
 * Register a go function in JS
 */
func (self *Engine) RegisterGoFunction(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.setGoMethod(cstr, uint32_t_To_Jerry_value_t(getGlobalObject()))
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

	// Call function on the global object here.
	return callJsFunction(getGlobalObject(), name, params)
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
