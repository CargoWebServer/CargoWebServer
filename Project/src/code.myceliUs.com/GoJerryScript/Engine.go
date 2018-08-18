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
                jerry_size_t args_count){

	// evaluate the result
	jerry_value_t ret = jerry_call_function ( func_obj_val, this_val,args_p, args_count );
	return ret;
}

jerry_value_t
parse_function (const char *resource_name_p,
                      size_t resource_name_length,
                      const char *arg_list_p,
                      size_t arg_list_size,
                      const char *source_p,
                      size_t source_size,
                      uint32_t parse_opts){

	return jerry_parse_function (resource_name_p, resource_name_length, arg_list_p,
		arg_list_size, source_p, source_size, parse_opts);
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

/**
 * Parse and set a function in the Javascript.
 * name The name of the function (the function will be keep in the engine for it
 *      lifetime.
 * args The argument name for that function.
 * src  The body of the function
 * options Can be JERRY_PARSE_NO_OPTS or JERRY_PARSE_STRICT_MODE
 */
func (self *Engine) AppendJsFunction(name string, args []string, src string) error {
	err := appendJsFunction(name, args, src)
	return err
}

/**
 * Register a go function in JS
 */
func (self *Engine) RegisterGoFunction(name string) {
	cs := C.CString(name)
	C.setGoFct(cs)
	defer C.free(unsafe.Pointer(cs))
}

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
func (self *Engine) GetGlobalVariable(name string) (*Value, error) {
	// first of all I will initialyse the arguments.
	globalObject := Jerry_get_global_object()
	defer Jerry_release_value(globalObject)
	propertyName := goToJs(name)
	defer Jerry_release_value(propertyName)
	property := Jerry_get_property(globalObject, propertyName)

	if Jerry_value_is_error(property) {
		return nil, errors.New("No variable found with name " + name)
	}

	value := NewValue(property)

	return value, nil
}

/**
 * Create an object with a given name.
 */
func (self *Engine) CreateObject(name string) *Object {
	// Create an object and return it.
	result := NewObject(name)
	return result
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
