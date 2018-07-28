package GoJerryScript

import (
	"encoding/json"
	"reflect"
	"unsafe"

	"C"

	"code.myceliUs.com/Utility"
)

// TODO do like otto Value...
type Value struct {
}

/**
 * Create a go string from a JS string pointer.
 */
func jsStrToGoStr(input Uint32_t) string {
	// Size info, ptr and it value
	sizePtr := Jerry_get_string_size(input)
	sizeValue := *(*uint64)(unsafe.Pointer(sizePtr.Swigcptr()))
	var buffer Uint8
	buffer.ptr = C.malloc(C.ulong(sizeValue)) // allocated the size.

	// Test if the string is a valid utf8 string...
	isUtf8 := Jerry_is_valid_utf8_string(input, sizePtr)
	if isUtf8 {
		Jerry_string_to_utf8_char_buffer(input, buffer, sizePtr)
	} else {
		Jerry_string_to_char_buffer(input, buffer, sizePtr)
	}
	// Set the slice information...
	h := &reflect.SliceHeader{buffer.Swigcptr(), int(sizeValue), int(sizeValue)}
	NewSlice := *(*[]byte)(unsafe.Pointer(h))

	// Set the Go string value from the slice.
	value := string(NewSlice)

	// release the memory used by the buffer.
	buffer.Free()

	return value
}

func goToJs(value interface{}) Uint32_t {
	var propValue Uint32_t
	if reflect.TypeOf(value).Kind() == reflect.String {
		// String value
		propValue = Jerry_create_string_from_utf8(NewUint8FromString(value.(string)))
	} else if reflect.TypeOf(value).Kind() == reflect.Bool {
		// Boolean value
		propValue = Jerry_create_boolean(value.(bool))
	} else if reflect.TypeOf(value).Kind() == reflect.Int {
		propValue = Jerry_create_number(float64(value.(int)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int8 {
		propValue = Jerry_create_number(float64(value.(int8)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int16 {
		propValue = Jerry_create_number(float64(value.(int16)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int32 {
		propValue = Jerry_create_number(float64(value.(int32)))
	} else if reflect.TypeOf(value).Kind() == reflect.Int64 {
		propValue = Jerry_create_number(float64(value.(int64)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint {
		propValue = Jerry_create_number(float64(value.(uint)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint8 {
		propValue = Jerry_create_number(float64(value.(uint8)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint16 {
		propValue = Jerry_create_number(float64(value.(uint16)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint32 {
		propValue = Jerry_create_number(float64(value.(uint32)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint64 {
		propValue = Jerry_create_number(float64(value.(uint64)))
	} else if reflect.TypeOf(value).Kind() == reflect.Float32 {
		propValue = Jerry_create_number(float64(value.(float32)))
	} else if reflect.TypeOf(value).Kind() == reflect.Float64 {
		propValue = Jerry_create_number(value.(float64))
	} else if reflect.TypeOf(value).Kind() == reflect.Slice {
		// So here I will create a array and put value in it.
		var size Uint32
		s := reflect.ValueOf(value)
		l := s.Len()
		size.ptr = unsafe.Pointer(&l)
		propValue = Jerry_create_array(size)
		for i := 0; i < l; i++ {
			v := goToJs(s.Index(i).Interface())
			var index Uint32
			index.ptr = unsafe.Pointer(&i)
			Jerry_set_property_by_index(propValue, index, v)

			// Release memory
			Jerry_release_value(v)
		}
	} else if reflect.TypeOf(value).Kind() == reflect.Struct {

	}

	return propValue
}

/**
 * Return equivalent value of a 32 bit c pointer.
 */
func jsToGo(input Uint32_t) (interface{}, error) {

	// the Go value...
	var value interface{}
	typeInfo := Jerry_value_get_type(input)
	// Now I will get the result if any...
	if typeInfo == Jerry_type_t(JERRY_TYPE_NUMBER) {
		value = Jerry_get_number_value(input)
	} else if typeInfo == Jerry_type_t(JERRY_TYPE_STRING) {
		value = jsStrToGoStr(input)
	} else if typeInfo == Jerry_type_t(JERRY_TYPE_BOOLEAN) {
		value = Jerry_get_boolean_value(input)
	} else if Jerry_value_is_typedarray(input) {
		/** Not made use of typed array **/
	} else if Jerry_value_is_array(input) {
		countPtr := Jerry_get_array_length(input)
		count := *(*uint64)(unsafe.Pointer(countPtr.Swigcptr()))
		// So here I got a array without type so I will get it property by index
		// and interpret each result.
		value = make([]interface{}, 0)
		var i uint64
		for i = 0; i < count; i++ {
			var index Uint32
			index.ptr = unsafe.Pointer(&i)
			e := Jerry_get_property_by_index(input, index)
			v, err := jsToGo(e)
			if err == nil {
				value = append(value.([]interface{}), v)
			}
			Jerry_release_value(e)
		}

	} else if Jerry_value_is_object(input) {
		// The go object will be a copy of the Js object.
		stringified := Jerry_json_stringfy(input)
		// if there is no error
		if !Jerry_value_is_error(stringified) {
			jsonStr := jsStrToGoStr(stringified)
			data := make(map[string]interface{}, 0)
			err := json.Unmarshal([]byte(jsonStr), &data)

			if err == nil {
				if data["TYPENAME"] != nil {
					relfectValue := Utility.MakeInstance(data["TYPENAME"].(string), data, SetEntity)
					value = relfectValue.Interface()
				} else {
					// Here map[string]interface{} will be use.
					value = data
				}

			} else {
				return nil, err
			}
		}
	}

	return value, nil
}
