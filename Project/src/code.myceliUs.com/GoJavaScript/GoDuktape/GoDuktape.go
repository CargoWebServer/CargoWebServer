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
extern void load_byte_code(duk_context_ptr ctx, char* bytecode, int length);

*/
import "C"
import "reflect"
import "unsafe"
import "errors"
import "strconv"
import "code.myceliUs.com/GoJavaScript"
import b64 "encoding/base64"

import "log"
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

func isJsFunction(ctx C.duk_context_ptr, index int) bool {
	// Here the function is compile...
	if int(C.duk_is_ecmascript_function(ctx, C.int(index))) > 0 {
		return true
	}

	if int(C.duk_is_function(ctx, C.int(index))) > 0 {
		return true
	}

	return false
}

func isGoFunction(ctx C.duk_context_ptr, index int) bool {
	return int(C.duk_is_c_function(ctx, C.int(index))) > 0
}

func isError(ctx C.duk_context_ptr, index int) bool {

	return int(C.is_error(ctx, C.int(index))) > 0
}

// Retreive an object by it uuid as a global object property.
func getJsObjectByUuid(uuid string, ctx C.duk_context_ptr) error {

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
		if objInfos.(map[string]interface{})["Methods"] != nil {
			methods := objInfos.(map[string]interface{})["Methods"].(map[string]interface{})
			for name, src := range methods {
				cstr := C.CString(name)
				if len(src.(string)) == 0 {
					// Set the go function here.
					C.push_c_function(ctx, cstr)
					C.duk_put_prop_string(ctx, obj_idx, cstr)
				} else {
					byteCode, err := b64.StdEncoding.DecodeString(src.(string))
					if err == nil {
						byteCode_ := (*_Ctype_char)(C.CBytes(byteCode))
						C.load_byte_code(ctx, byteCode_, C.int(len(byteCode)))
						C.free(unsafe.Pointer(byteCode_))

						// replaces the buffer on the stack top with the original function
						C.duk_load_function(ctx)

						// Now the ctx contain a function
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
				}
				C.free(unsafe.Pointer(cstr))
			}

			// I can remove the methods from the infos.
			delete(objInfos.(map[string]interface{}), "Methods")
		}

		// Now the object properties.
		for name, value := range objInfos.(map[string]interface{}) {
			cname := C.CString(name)
			if value != nil {
				if reflect.TypeOf(value).Kind() == reflect.Slice {
					slice := reflect.ValueOf(value)
					values := C.duk_push_array(ctx)
					for i := 0; i < slice.Len(); i++ {
						if slice.Index(i).IsValid() {
							if slice.Index(i).IsNil() {
								// In case of null values.
								C.duk_push_null(ctx)
								C.duk_put_prop_index(ctx, values, C.uint(i))
							} else {
								e := slice.Index(i).Interface()
								if reflect.TypeOf(e).Kind() == reflect.Map {
									// Here The value contain a map... so I will append
									if e.(map[string]interface{})["TYPENAME"] != nil {
										if e.(map[string]interface{})["TYPENAME"].(string) == "GoJavaScript.ObjectRef" {
											// Here I will pup the object in the stack value.
											err := getJsObjectByUuid(e.(map[string]interface{})["UUID"].(string), ctx)
											if err == nil {
												C.duk_put_prop_index(ctx, values, C.uint(i))
											}
										} else {
											//log.Panicln("214 GoDukTape.go ---> unhandle case!!!!")
											setValue(ctx, e)
											C.duk_put_prop_index(ctx, values, C.uint(i))
										}
									} else {
										setValue(ctx, e)
										C.duk_put_prop_index(ctx, values, C.uint(i))
									}
								} else {
									// set the value on the stack.
									setValue(ctx, e)
									// Release the result
									C.duk_put_prop_index(ctx, values, C.uint(i))
								}
							}
						}
					}
					C.duk_put_prop_string(ctx, obj_idx, cname)
				} else if reflect.TypeOf(value).Kind() == reflect.Map {
					if value.(map[string]interface{})["TYPENAME"] != nil {
						if value.(map[string]interface{})["TYPENAME"].(string) == "GoJavaScript.ObjectRef" {
							err := getJsObjectByUuid(value.(map[string]interface{})["UUID"].(string), ctx)
							if err == nil {
								C.duk_put_prop_string(ctx, obj_idx, cname)
							}
						} else {
							setValue(ctx, value)
							C.duk_put_prop_string(ctx, obj_idx, cname)
						}
					} else {
						setValue(ctx, value)
						C.duk_put_prop_string(ctx, obj_idx, cname)
					}
				} else {
					// Standard object property, int, string, float...
					setValue(ctx, value)
					C.duk_put_prop_string(ctx, obj_idx, cname)
				}

				C.free(unsafe.Pointer(cname))
			} else {
				// set null value
				C.duk_push_null(ctx)
				C.duk_put_prop_string(ctx, obj_idx, cname)
			}
		}
	}
	return err
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
					// Put the object in the map.
					getJsObjectByUuid(ref.(*GoJavaScript.ObjectRef).UUID, ctx)
				} else {
					log.Println("--> fail to Create Go object ", string(jsonStr), err)
				}
			} else if value.(map[string]interface{})["UUID"] != nil { // It's a Go Javasacript object.
				uuid := value.(map[string]interface{})["UUID"].(string)
				getJsObjectByUuid(uuid, ctx)
			} else {
				// Not a registered type...
				cstr := C.CString(string(jsonStr))
				defer C.free(unsafe.Pointer(cstr))
				C.duk_push_string(ctx, cstr)
				C.duk_json_decode(ctx, C.int(-1))
			}
			/**/
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
		//log.Println("---> array found")
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
	} else if isGoFunction(ctx, index) {
		// Here I will return the function name
		C.duk_dup(ctx, C.int(index))
		str := C.GoString(C.duk_to_string(ctx, -1))
		C.duk_pop(ctx)
		// here i will retrun the string value of the go function.
		return str, nil
	} else if isJsFunction(ctx, index) {
		// I will duplicate the function on the stack and create it buffer version.
		C.duk_dup(ctx, C.int(index))
		// put the byte code in the stack
		C.duk_dump_function(ctx)
		var size int
		ptr := C.duk_get_buffer(ctx, -1, (*C.duk_size_t)(unsafe.Pointer(&size)))
		bytcode := C.GoBytes(ptr, C.int(size))
		C.duk_pop(ctx)
		return bytcode, nil
	} else if isObject(ctx, index) {
		// Here it can be an array or an object...
		// The go object will be a copy of the Js object.
		uuid_name := C.CString("uuid_")
		if int(C.duk_has_prop_string(ctx, C.int(index), uuid_name)) > 0 {
			// Get the uuid.
			C.duk_get_prop_string(ctx, C.int(index), uuid_name)
			uuid, _ := getValue(ctx, -1)
			C.duk_pop(ctx) // pop the string.

			// Here I must set the client object value with the value in memory.
			C.duk_dup(ctx, C.int(index))
			jsonStr := C.GoString(C.duk_json_encode(ctx, -1))
			C.duk_pop(ctx) // pop the json string.

			// Set the object values with values retreive in the vm object
			GoJavaScript.CallGoFunction("Client", "SetObjectValues", uuid.(string), string(jsonStr))

			// Return and object reference.
			ref := GoJavaScript.NewObjectRef(uuid.(string))
			return ref, nil
		} else {
			if int(C.duk_has_prop_string(ctx, C.int(index), C.CString("TYPENAME"))) > 0 {
				// The js object is an entity.
				C.duk_dup(ctx, C.int(index))
				jsonStr := C.GoString(C.duk_json_encode(ctx, -1))
				C.duk_pop(ctx) // pop the json string.
				return GoJavaScript.CallGoFunction("Client", "CreateGoObject", string(jsonStr))
			} else {
				// Here i will create a Js go object that will serve as object handle from
				// go code.
				obj, err := GoJavaScript.CallGoFunction("Client", "CreateObject", "")
				uuid := obj.(*GoJavaScript.ObjectRef).UUID

				properties := make([]string, 0)
				C.duk_enum(ctx, C.int(index), 0)
				for int(C.duk_next(ctx, -1, 0)) > 0 {
					// [ ... enum key ]
					properties = append(properties, C.GoString(C.duk_get_string(ctx, -1)))
					C.duk_pop(ctx) // pop key
				}

				C.duk_pop(ctx) // pop enum object
				for i := 0; i < len(properties); i++ {
					name := C.CString(properties[i])
					C.duk_get_prop_string(ctx, C.int(index), name)
					if isJsFunction(ctx, -1) {
						value, err := getValue(ctx, -1)
						if err == nil {
							GoJavaScript.CallGoFunction("Client", "SetObjectJsMethod", uuid, properties[i], value)
						}
					} else if isGoFunction(ctx, -1) {
						_, err := getValue(ctx, -1)
						if err == nil {
							GoJavaScript.CallGoFunction("Client", "SetObjectGoMethod", uuid, properties[i])
						}
					} else {
						value, err := getValue(ctx, -1)
						if err == nil {
							GoJavaScript.CallGoFunction("Client", "SetObjectProperty", uuid, properties[i], value)
						}
					}
					C.free(unsafe.Pointer(name))
					C.duk_pop(ctx) // pop value object
				}
				return obj, err
			}

		}
		C.free(unsafe.Pointer(uuid_name))

	} else if C.duk_get_type(ctx, C.int(index)) == C.DUK_TYPE_MASK_NONE ||
		C.duk_get_type(ctx, C.int(index)) == C.DUK_TYPE_MASK_NULL ||
		C.duk_get_type(ctx, C.int(index)) == C.DUK_TYPE_MASK_UNDEFINED {
		return nil, nil
	}

	return nil, errors.New("no value found at index " + strconv.Itoa(index) + " with type index " + strconv.Itoa(int(C.duk_get_type(ctx, C.int(index)))))
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
func c_function_handler(ctx C.duk_context_ptr, name_cstr *C.char, uuid_cstr *C.char) C.duk_ret_t {
	// Push the current function on the context
	name := C.GoString(name_cstr)
	var uuid string
	if uuid_cstr != nil {
		uuid = C.GoString(uuid_cstr)
	}

	// Now I will get back the list of arguments.
	size := int(C.duk_get_top(ctx)) - 1
	params := make([]interface{}, size)
	for i := 0; i < size; i++ {
		value, err := getValue(ctx, i)
		if err == nil {
			params[i] = value
		} else {
			log.Panicln("512 GoDukTape.go")
		}
	}

	if len(uuid) > 0 {
		// I will now call the function.
		result, err := GoJavaScript.CallGoFunction(uuid, name, params...)
		if err == nil && result != nil {
			// So here I will set the value
			setValue(ctx, result)
			return C.duk_ret_t(1)
		} else if err != nil {
			// error occured here.
			log.Println("525 GoDukTage.go ---> ", err)
		}

	} else {
		result, err := GoJavaScript.CallGoFunction("", name, params...)
		if err == nil && result != nil {
			setValue(ctx, result)
			// The value must be pop by the caller.
			return C.duk_ret_t(1)
		} else if err != nil {
			// error occured here.
			log.Println("536 GoDukTape.go ---> ", err)
		}
	}

	return C.duk_ret_t(0) // return undefined.

}
