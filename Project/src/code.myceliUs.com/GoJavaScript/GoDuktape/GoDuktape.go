package GoDuktape

//#cgo LDFLAGS:  -lm
/*
#include "duktape.h"
typedef duk_context* duk_context_ptr;
extern duk_idx_t push_c_function(duk_context_ptr ctx, const char* name);
extern duk_int_t eval_string(duk_context_ptr context, const char* src);
extern duk_int_t compile_function_string(duk_context_ptr ctx, const char* src);
extern const char* safe_to_string(duk_context_ptr ctx, duk_idx_t index);
extern duk_bool_t is_error(duk_context_ptr ctx, duk_idx_t index);
extern const char* safe_to_string(duk_context_ptr ctx, duk_idx_t idx);
extern duk_idx_t set_finalizer(duk_context_ptr ctx, const char* uuid);
*/
import "C"
import "reflect"
import "unsafe"
import "errors"
import "strconv"
import "code.myceliUs.com/GoJavaScript"
import "log"
import "strings"
import "encoding/json"

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

func isArray(ctx C.duk_context_ptr, index int) bool {
	return int(C.duk_is_array(ctx, C.int(index))) > 0
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

func isError(ctx C.duk_context_ptr, index int) bool {

	return int(C.is_error(ctx, C.int(index))) > 0
}

// Retreive an object by it uuid as a global object property.
func getJsObjectByUuid(uuid string, ctx C.duk_context_ptr) {

	// So here I will try to create a local Js representation of the object.
	objInfos, err := GoJavaScript.CallGoFunction("Client", "GetGoObjectInfos", uuid)
	if err == nil {
		// So here I got an object map info.
		// Create the object JS object.
		obj_idx := C.duk_push_object(ctx)

		// Now I will set the uuid property.
		uuid_value := C.CString(uuid)

		C.set_finalizer(ctx, uuid_value)
		C.duk_push_string(ctx, uuid_value)
		defer C.free(unsafe.Pointer(uuid_value))

		uuid_name := C.CString("uuid_")
		C.duk_put_prop_string(ctx, obj_idx, uuid_name)

		defer C.free(unsafe.Pointer(uuid_name))

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
			cname := C.CString(name)
			if reflect.TypeOf(value).Kind() == reflect.Slice {
				slice := reflect.ValueOf(value)
				values := C.duk_push_array(ctx)
				for i := 0; i < slice.Len(); i++ {
					e := slice.Index(i).Interface()
					if reflect.TypeOf(e).Kind() == reflect.Map {
						// Here The value contain a map... so I will append
						if e.(map[string]interface{})["TYPENAME"] != nil {
							if e.(map[string]interface{})["TYPENAME"].(string) == "GoJavaScript.ObjectRef" {
								// Here I will pup the object in the stack value.
								getJsObjectByUuid(e.(map[string]interface{})["UUID"].(string), ctx)
								C.duk_put_prop_index(ctx, values, C.uint(i))
							}
						} else {
							log.Println("---> unknow object propertie type 231")
						}
					} else {
						// set the value on the stack.
						setValue(ctx, e)
						// Release the result
						C.duk_put_prop_index(ctx, values, C.uint(i))
					}
				}
				C.duk_put_prop_string(ctx, obj_idx, cname)
			} else if reflect.TypeOf(value).Kind() == reflect.Map {
				if value.(map[string]interface{})["TYPENAME"] != nil {
					if value.(map[string]interface{})["TYPENAME"].(string) == "GoJavaScript.ObjectRef" {
						getJsObjectByUuid(value.(map[string]interface{})["UUID"].(string), ctx)
						C.duk_put_prop_string(ctx, obj_idx, cname)
					} else {
						log.Println("---> unknow object propertie type 245")
					}
				}
			} else {
				// Standard object property, int, string, float...
				setValue(ctx, value)
				C.duk_put_prop_string(ctx, obj_idx, cname)
			}

			C.free(unsafe.Pointer(cname))
		}
	}
}

/**
 * Set a go value in JavaScript context.
 */
func setValue(ctx C.duk_context_ptr, value interface{}) {
	// if the value is null I will push a null value
	// on the context.
	if value == nil {
		C.duk_push_null(ctx)
		return
	}

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
	} else if reflect.TypeOf(value).String() == "GoJavaScript.ObjectRef" {
		// I got a Js object reference.
		uuid := value.(GoJavaScript.ObjectRef).UUID
		getJsObjectByUuid(uuid, ctx)
	} else if reflect.TypeOf(value).String() == "map[string]interface {}" {
		jsonStr, err := json.Marshal(value)
		if err == nil {
			if value.(map[string]interface{})["TYPENAME"] != nil {
				ref, err := GoJavaScript.CallGoFunction("Client", "CreateGoObject", string(jsonStr))
				if err == nil {
					// In that case an object exist in the case...
					getJsObjectByUuid(ref.(*GoJavaScript.ObjectRef).UUID, ctx)
				} else {
					log.Println("--> fail to Create Go object ", string(jsonStr), err)
				}
			} else {
				// Not a registered type...
				cstr := C.CString(string(jsonStr))
				defer C.free(unsafe.Pointer(cstr))
				C.duk_push_string(ctx, cstr)
				C.duk_json_decode(ctx, C.int(-1))
			}
		}

	} else {
		log.Println("----> no type found for ", reflect.TypeOf(value).String(), value)
	}
}

/**
 * Return a value at a given index inside a given context.
 */
func getValue(ctx C.duk_context_ptr, index int) (interface{}, error) {
	if isString(ctx, index) {
		//log.Println("---> string found")
		cstr := C.duk_get_string(ctx, C.int(index))
		return C.GoString(cstr), nil
	} else if isBool(ctx, index) {
		//log.Println("---> boolean found")
		return uint(C.duk_get_boolean(ctx, C.int(index))) > 0, nil
	} else if isNumber(ctx, index) {
		n := float64(C.duk_get_number(ctx, C.int(index)))
		return n, nil
	} else if isError(ctx, index) {
		//log.Println("---> error found")
		stack := C.CString("stack")
		C.duk_get_prop_string(ctx, C.int(index), stack)
		err := errors.New(C.GoString(C.safe_to_string(ctx, C.int(-1))))
		C.free(unsafe.Pointer(stack))
		log.Println("323 ----> ", err)
		return nil, err
	} else if isArray(ctx, index) {
		// The object is an array
		size := int(C.duk_get_length(ctx, C.int(index)))
		array := make([]interface{}, size)
		for i := 0; i < size; i++ {
			C.duk_get_prop_index(ctx, C.int(index), C.uint(i))
			v, err := getValue(ctx, -1)
			if err == nil {
				// Set back the go value.
				array[i] = v
			}
			C.duk_pop(ctx) // back to the array and not the value.
		}
		return array, nil
	} else if isFunction(ctx, index) {
		log.Println("---> is a function: 309")
	} else if isObject(ctx, index) {
		// Here it can be an array or an object...
		// The go object will be a copy of the Js object.
		uuid_name := C.CString("uuid_")
		if int(C.duk_has_prop_string(ctx, C.int(index), uuid_name)) > 0 {
			// Get the uuid.
			C.duk_get_prop_string(ctx, C.int(index), uuid_name)
			uuid, _ := getValue(ctx, -1)
			C.duk_pop(ctx) // pop the string.
			// Return and object reference.
			return GoJavaScript.NewObjectRef(uuid.(string)), nil

		} else {
			// if there is no error
			jsonStr := C.GoString(C.duk_json_encode(ctx, C.int(index)))
			C.duk_pop(ctx) // remove the string from the top of the stack
			if strings.Index(jsonStr, "TYPENAME") != -1 {
				// So here I will create a remote action and tell the client to
				// create a Go object from jsonStr. The object will be set by
				// the client on the server.
				return GoJavaScript.CallGoFunction("Client", "CreateGoObject", jsonStr)
			}
			// In that case the object has no go representation...
			// and must be use only in JS.
			return nil, nil

		}
		C.free(unsafe.Pointer(uuid_name))

	} else if C.duk_get_type(ctx, C.int(index)) == C.DUK_TYPE_MASK_NONE ||
		C.duk_get_type(ctx, C.int(index)) == C.DUK_TYPE_MASK_NULL ||
		C.duk_get_type(ctx, C.int(index)) == C.DUK_TYPE_MASK_UNDEFINED {
		return nil, nil
	}

	return nil, errors.New("no value found at index " + strconv.Itoa(index))
}

//export c_finalizer
func c_finalizer(ctx C.duk_context_ptr) C.duk_ret_t {
	// Push the current function on the context
	C.duk_push_current_function(ctx)

	// Get it name..
	C.duk_get_prop_string(ctx, -1, C.CString("uuid"))

	uuid, err := getValue(ctx, -1)
	C.duk_pop(ctx) // back to the context calling context.

	if err == nil {
		// delete the object from the client cache.
		// Now I will ask the client side to remove it object reference to.
		GoJavaScript.CallGoFunction("Client", "DeleteGoObject", uuid)
		return C.duk_ret_t(1)
	}

	return C.duk_ret_t(0) // return undefined.
}

//export c_function_handler
func c_function_handler(ctx C.duk_context_ptr) C.duk_ret_t {
	// Push the current function on the context
	C.duk_push_current_function(ctx)

	// Get it name..
	C.duk_get_prop_string(ctx, -1, C.CString("name"))

	name, err := getValue(ctx, -1)
	C.duk_pop(ctx) // back to the context calling context.

	if err == nil {
		// Now I will get back the list of arguments.
		size := int(C.duk_get_top(ctx)) - 1
		params := make([]interface{}, size)
		for i := 0; i < size; i++ {
			value, err := getValue(ctx, i)
			if err == nil {
				params[i] = value
			} else {
				log.Panicln("---> fail 393")
			}
		}

		// Now I will get the this
		C.duk_push_this(ctx)

		if isObject(ctx, -1) {
			uuid_ := C.CString("uuid_")
			C.duk_push_string(ctx, uuid_)
			defer C.free(unsafe.Pointer(uuid_))
			if int(C.duk_has_prop_string(ctx, -2, uuid_)) > 0 {
				C.duk_get_prop_string(ctx, -2, uuid_)
				uuid, _ := getValue(ctx, -1)
				C.duk_pop(ctx)
				// I will now call the function.
				result, err := GoJavaScript.CallGoFunction(uuid.(string), name.(string), params...)
				if err == nil && result != nil {
					// So here I will set the value
					setValue(ctx, result)
					return C.duk_ret_t(1)
				} else if err != nil {
					// error occured here.
				}
			}
		} else {
			result, err := GoJavaScript.CallGoFunction("", name.(string), params...)
			if err == nil && result != nil {
				setValue(ctx, result)
				// The value must be pop by the caller.
				return C.duk_ret_t(1)
			} else if err != nil {
				// error occured here.
			}
		}
	}
	return C.duk_ret_t(0) // return undefined.

}
