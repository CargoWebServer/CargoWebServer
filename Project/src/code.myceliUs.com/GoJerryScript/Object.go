package GoJerryScript

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "jerryscript.h"
#include "jerryscript-ext/handler.h"
#include "jerryscript-debugger.h"

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
*/
import "C"
import "errors"
import "reflect"

//import "log"
import "code.myceliUs.com/Utility"
import "fmt"

type Object struct {
	// The pointer to the object inside the vm.
	ptr Uint32_t
}

func NewObject(name string) *Object {

	// The object itself.
	obj := new(Object)

	ptrString := fmt.Sprintf("%d", obj)
	uuid := Utility.GenerateUUID(ptrString)

	// Keep the pointer.
	obj.ptr = jerry_value_t_To_uint32_t(C.create_native_object(C.CString(uuid)))

	// Set the object in the cache.
	GetCache().setJsObject(uuid, obj.ptr)

	globalObj := Jerry_get_global_object()
	defer Jerry_release_value(globalObj)

	propName := goToJs(name)
	defer Jerry_release_value(propName)

	// Set the object as a propety in the global object.
	Jerry_release_value(Jerry_set_property(globalObj, propName, obj.ptr))

	return obj
}

/**
 * Set the property of an object.
 */
func (self *Object) Set(name string, value interface{}) {

	// If the value is a function
	if reflect.TypeOf(value).Kind() == reflect.Func {
		setGoMethod(self.ptr, name, value)
	} else {
		// Set the property value.
		Jerry_release_value(Jerry_set_property(self.ptr, goToJs(name), goToJs(value)))
	}
}

/**
 * Return property value.
 */
func (self *Object) Get(name string) (*Value, error) {
	// Return the value from an object.
	ptr := Jerry_get_property(self.ptr, goToJs(name))
	if Jerry_value_is_error(ptr) {
		return nil, errors.New("no property found with name " + name)
	}

	value := NewValue(ptr)

	return value, nil
}

/**
 * Call a function over an object.
 */
func (self *Object) Call(name string, params ...interface{}) (Value, error) {
	return callJsFunction(self.ptr, name, params)
}
