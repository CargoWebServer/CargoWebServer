package Utility

import (
	"errors"
	//"log"
	"reflect"
)

/**
 * Make use of reflexion to call the method specified in the message.
 */
func CallMethod(i interface{}, methodName string, params []interface{}) (interface{}, interface{}) {

	//log.Println("Call method ", methodName, " with params ", params)
	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value
	value = reflect.ValueOf(i)

	// In case of a nil pointer...
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return "", errors.New("Nil pointer!")
		}
	}

	// if we start with a pointer, we need to get value pointed to
	// if we start with a value, we need to get a pointer to that value
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(i))
		temp := ptr.Elem()
		temp.Set(value)
	}

	// check for method on value
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	// check for method on pointer
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		if param != nil {
			in[k] = reflect.ValueOf(param)
		} else {
			var nilVal interface{}
			in[k] = reflect.ValueOf(&nilVal).Elem()
		}
	}

	if finalMethod.IsValid() {
		results := finalMethod.Call(in)
		if len(results) > 0 {
			if len(results) == 1 {
				// One result here...
				switch goError := results[0].Interface().(type) {
				case error:
					return nil, goError
				}

				return results[0].Interface(), nil
			} else if len(results) == 2 {
				// Return the result and the error after.
				return results[0].Interface(), results[1].Interface()
			}
		}
	} else {
		return nil, errors.New("Method dosen't exist!")
	}

	// return or panic, method not found of either type
	return "", nil
}
