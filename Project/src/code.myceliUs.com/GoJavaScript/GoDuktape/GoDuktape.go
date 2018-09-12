package GoDuktape

//#cgo LDFLAGS:  -lm
/*
#include "duktape.h"
typedef duk_context* duk_context_ptr;
extern duk_idx_t push_c_function(duk_context_ptr ctx, const char* name);
extern duk_int_t eval_string(duk_context_ptr context, const char* src);
extern duk_int_t compile_function_string(duk_context_ptr ctx, const char* src);
extern const char* safe_to_string(duk_context_ptr ctx, duk_idx_t index);

*/
import "C"
import "reflect"
import "unsafe"
import "errors"
import "strconv"
import "code.myceliUs.com/GoJavaScript"
import "log"

/**
 * Go and Javascript functions bindings.
 */

// Type checking functions.
// The first argument is the context pointer.
// The second is the index of the value in the context to test.
func isString(ctx C.duk_context_ptr, index int) bool {
	type_ := C.duk_get_type_mask(ctx, C.int(index))
	if type_ == C.DUK_TYPE_MASK_STRING {
		return true
	}
	return false
}

func isNumber(ctx C.duk_context_ptr, index int) bool {
	type_ := C.duk_get_type_mask(ctx, C.int(index))
	if type_ == C.DUK_TYPE_MASK_NUMBER {
		return true
	}
	return false
}

func isNone(ctx C.duk_context_ptr, index int) bool {
	type_ := C.duk_get_type_mask(ctx, C.int(index))
	if type_ == C.DUK_TYPE_MASK_NONE {
		return true
	}
	return false
}

func isNull(ctx C.duk_context_ptr, index int) bool {
	type_ := C.duk_get_type_mask(ctx, C.int(index))
	if type_ == C.DUK_TYPE_MASK_NULL {
		return true
	}
	return false
}

func isUndefined(ctx C.duk_context_ptr, index int) bool {
	type_ := C.duk_get_type_mask(ctx, C.int(index))
	if type_ == C.DUK_TYPE_MASK_UNDEFINED {
		return true
	}
	return false
}

func isBool(ctx C.duk_context_ptr, index int) bool {
	type_ := C.duk_get_type_mask(ctx, C.int(index))
	if type_ == C.DUK_TYPE_MASK_BOOLEAN {
		return true
	}
	return false
}

func isPointer(ctx C.duk_context_ptr, index int) bool {
	type_ := C.duk_get_type_mask(ctx, C.int(index))
	if type_ == C.DUK_TYPE_MASK_POINTER {
		return true
	}
	return false
}

func isObject(ctx C.duk_context_ptr, index int) bool {
	type_ := C.duk_get_type_mask(ctx, C.int(index))
	if type_ == C.DUK_TYPE_MASK_OBJECT {
		return true
	}
	return false
}

func isBuffer(ctx C.duk_context_ptr, index int) bool {
	type_ := C.duk_get_type_mask(ctx, C.int(index))
	if type_ == C.DUK_TYPE_MASK_BUFFER {
		return true
	}
	return false
}

func isFunction(ctx C.duk_context_ptr, index int) bool {
	type_ := C.duk_get_type_mask(ctx, C.int(index))
	if type_ == C.DUK_TYPE_MASK_LIGHTFUNC {
		return true
	}
	return false
}

// Retreive an object by it uuid as a global object property.
func getJsObjectByUuid(uuid string, ctx C.duk_context_ptr) {

	// So here I will try to create a local Js representation of the object.
	objInfos, err := GoJavaScript.CallGoFunction("Client", "GetGoObjectInfos", uuid)
	log.Println("----> object infos: ", objInfos)

	if err == nil {
		// So here I got an object map info.
		// Create the object JS object.
		obj_idx := C.duk_push_object(ctx)

		// Now I will set the uuid property.
		uuid_value := C.CString(uuid)
		C.duk_push_string(ctx, uuid_value)
		defer C.free(unsafe.Pointer(uuid_value))

		uuid_name := C.CString("uuid_")
		C.duk_put_prop_string(ctx, obj_idx, uuid_name)
		C.free(unsafe.Pointer(uuid_name))

		// Now I will set the object method.
		methods := objInfos.(map[string]interface{})["Methods"].(map[string]interface{})
		for name, src := range methods {
			cstr := C.CString(name)
			if len(src.(string)) == 0 {
				// Set the go function here.
				C.push_c_function(ctx, cstr)
				C.duk_put_prop_string(ctx, obj_idx, cstr)
			} else {
				src_ := C.CString(src.(string))
				if int(C.compile_function_string(ctx, src_)) != 0 {
					log.Println("---> compilation fail!", src)
					log.Println("---> error:", C.GoString(C.safe_to_string(ctx, -1)))
				} else {
					// Keep the function as object method.
					C.duk_put_prop_string(ctx, obj_idx, cstr)
				}
			}
			C.free(unsafe.Pointer(cstr))
		}

		// I can remove the methods from the infos.
		delete(objInfos.(map[string]interface{}), "Methods")

		// Now the object properties.
		for name, value := range objInfos.(map[string]interface{}) {
			if reflect.TypeOf(value).Kind() == reflect.Slice {
				slice := reflect.ValueOf(value)
				values := jerry_value_t_To_uint32_t(C.create_array(C.uint32_t(slice.Len())))
				for i := 0; i < slice.Len(); i++ {
					e := slice.Index(i).Interface()
					if reflect.TypeOf(e).Kind() == reflect.Map {
						// Here The value contain a map... so I will append
						if e.(map[string]interface{})["TYPENAME"] != nil {
							if e.(map[string]interface{})["TYPENAME"].(string) == "GoJavaScript.ObjectRef" {
								value_ := getJsObjectByUuid(e.(map[string]interface{})["UUID"].(string))
								r := C.set_property_by_index(uint32_t_To_Jerry_value_t(values), C.uint32_t(uint32(i)), uint32_t_To_Jerry_value_t(goToJs(value_)))
								// Release the result
								Jerry_release_value(jerry_value_t_To_uint32_t(r))
							}
						} else {
							log.Println("---> unknow object propertie type 231")
						}
					} else {
						r := C.set_property_by_index(uint32_t_To_Jerry_value_t(values), C.uint32_t(uint32(i)), uint32_t_To_Jerry_value_t(goToJs(e)))
						// Release the result
						Jerry_release_value(jerry_value_t_To_uint32_t(r))
					}
				}
				Jerry_set_object_property(obj, name, values)

			} else if reflect.TypeOf(value).Kind() == reflect.Map {
				if value.(map[string]interface{})["TYPENAME"] != nil {
					if value.(map[string]interface{})["TYPENAME"].(string) == "GoJavaScript.ObjectRef" {
						value_ := getJsObjectByUuid(value.(map[string]interface{})["UUID"].(string))
						Jerry_set_object_property(obj, name, value_)
					} else {
						log.Println("---> unknow object propertie type 245")
					}
				}
			} else {
				// Standard object property, int, string, float...
				Jerry_set_object_property(obj, name, value)
			}
		}

		// set on the global object.
		C.duk_put_global_string(ctx, uuid_value)

		// keep the object on the global object if name is define.
		if objInfos.(map[string]interface{})["Name"] != nil {
			if len(objInfos.(map[string]interface{})["Name"].(string)) > 0 {
				// set is name as global object property
				name_ := C.CString(objInfos.(map[string]interface{})["Name"].(string))
				C.duk_put_global_string(ctx, name_)
				C.free(unsafe.Pointer(name_))
			}
		}

	}

	log.Println("---> object ", uuid, "is undefined!")
}

/**
 * Set a go value in JavaScript context.
 */
func setValue(ctx C.duk_context_ptr, value interface{}) {
	if reflect.TypeOf(value).Kind() == reflect.String {
		cstr := C.CString(value.(string))
		C.duk_push_string(ctx, cstr)
		C.free(unsafe.Pointer(cstr))
	} else if reflect.TypeOf(value).Kind() == reflect.Bool {
		if value.(bool) {
			C.duk_push_boolean(ctx, C.duk_bool_t(uint(1)))
		} else {
			C.duk_push_boolean(ctx, C.duk_bool_t(uint(0)))
		}
	} else if reflect.TypeOf(value).Kind() == reflect.Int {
		C.duk_push_int(ctx, C.duk_int_t(value.(int)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int8 {
		C.duk_push_int(ctx, C.duk_int_t(value.(int8)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int16 {
		C.duk_push_int(ctx, C.duk_int_t(value.(int16)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int32 {
		C.duk_push_int(ctx, C.duk_int_t(value.(int32)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int64 {
		C.duk_push_int(ctx, C.duk_int_t(value.(int64)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint {
		C.duk_push_int(ctx, C.duk_int_t(value.(uint)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint8 {
		C.duk_push_int(ctx, C.duk_int_t(value.(uint8)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint16 {
		C.duk_push_int(ctx, C.duk_int_t(value.(uint16)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint32 {
		C.duk_push_int(ctx, C.duk_int_t(value.(uint32)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint64 {
		C.duk_push_int(ctx, C.duk_int_t(value.(uint64)))
	} else if reflect.TypeOf(value).Kind() == reflect.Float32 {
		C.duk_push_number(ctx, C.duk_double_t(value.(float32)))
	} else if reflect.TypeOf(value).Kind() == reflect.Float64 {
		C.duk_push_number(ctx, C.duk_double_t(value.(float64)))
	} else if reflect.TypeOf(value).Kind() == reflect.Slice {
		// So here I will create a array and put value in it.
		s := reflect.ValueOf(value)
		array := C.duk_push_array(ctx)
		for i := 0; i < s.Len(); i++ {
			// here I will set value...
			setValue(ctx, s.Index(i).Interface())
			// And I will push it index property
			C.duk_put_prop_index(ctx, array, C.uint(i))
		}
	} else if reflect.TypeOf(value).String() == "*GoJavaScript.ObjectRef" {
		// I got a Js object reference.
		uuid := value.(*GoJavaScript.ObjectRef).UUID
		getJsObjectByUuid(uuid, ctx)
	}

}

/**
 * Return a value at a given index inside a given context.
 */
func getValue(ctx C.duk_context_ptr, index int) (interface{}, error) {
	if isString(ctx, index) {
		cstr := C.duk_get_string(ctx, C.int(index))
		return C.GoString(cstr), nil
	} else if isBool(ctx, index) {
		return uint(C.duk_get_boolean(ctx, C.int(index))) > 0, nil
	} else if isNumber(ctx, index) {
		return float64(C.duk_get_number(ctx, C.int(index))), nil
	} else if isObject(ctx, index) {
		// Here it can be an array or an object...
		if int(C.duk_is_array(ctx, C.int(index))) > 0 {
			// The object is an array
			size := int(C.duk_get_length(ctx, C.int(index)))
			array := make([]interface{}, size)
			for i := 0; i < size; i++ {
				C.duk_get_prop_index(ctx, C.int(-1), C.uint(i))
				v, err := getValue(ctx, -1)
				if err == nil {
					// Set back the go value.
					array[i] = v
				}
				C.duk_pop(ctx)
			}
			return array, nil
		} else {
			// The object is an object

		}
	}
	return nil, errors.New("no value found at index " + strconv.Itoa(index))
}

//export c_function_handler
func c_function_handler(ctx C.duk_context_ptr) C.duk_ret_t {
	log.Println("---> c_function_handler call!")

	// Push the current function on the context
	C.duk_push_current_function(ctx)

	// Get it name..
	C.duk_get_prop_string(ctx, -1, C.CString("name"))

	name, err := getValue(ctx, -1)
	if err == nil {
		log.Println("call function ", name)
	}

	C.duk_pop(ctx) // back to the context calling context.

	// Now I will get back the list of arguments.
	size := int(C.duk_get_top(ctx)) - 1
	params := make([]interface{}, size)
	for i := 0; i < size; i++ {
		value, err := getValue(ctx, i)
		if err == nil {
			params[i] = value
		}
	}

	// Now I will get the this
	C.duk_push_this(ctx)
	uuid_ := C.CString("uuid_")
	C.duk_push_string(ctx, uuid_)
	C.free(unsafe.Pointer(uuid_))

	if int(C.duk_has_prop(ctx, -3)) > 0 {
		C.duk_pop(ctx)
	} else {
		C.duk_pop(ctx)
		result, err := GoJavaScript.CallGoFunction("", name.(string), params...)
		if err == nil && result != nil {
			// So here I will set the value
			setValue(ctx, result)
			// return value at top.
			return C.duk_ret_t(1)
		} else if err != nil {
			// error occured here.
		}
	}

	return C.duk_ret_t(0) // return undefined.
}
