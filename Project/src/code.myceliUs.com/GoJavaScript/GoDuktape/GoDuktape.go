package GoDuktape

//#cgo LDFLAGS:  -lm
/*
#include "duktape.h"
typedef duk_context* duk_context_ptr;


*/
import "C"
import "reflect"
import "unsafe"
import "errors"
import "strconv"

//import "log"

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
