package Utility

import (
	"errors"
	"log"
	"os"
	"reflect"
	"strconv"
)

/**
 * Make use of reflexion to call the method specified in the message.
 */
func CallMethod(i interface{}, methodName string, params []interface{}) (interface{}, interface{}) {

	if i == nil {
		return "", errors.New("Nil pointer!")
	}

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
	if method.IsValid() { // Variadic can't be restricted here...
		finalMethod = method
		// Here I will validate the parameters.
		if !method.Type().IsVariadic() {
			if method.Type().NumIn() != len(params) {
				errMsg := "Wrong number of parameter for method " + methodName + " expected " + strconv.Itoa(method.Type().NumIn()-1) + " got " + strconv.Itoa(len(params))
				return nil, errors.New(errMsg)
			}
		}
	}

	in := make([]reflect.Value, len(params))
	index := 0
	for k, param := range params {
		if param != nil {
			in[k] = reflect.ValueOf(param)
		} else {
			var nilVal interface{}
			in[k] = reflect.ValueOf(&nilVal).Elem()
		}
		index++
	}

	if finalMethod.IsValid() {
		// here because the method call can cause panic I will not call
		// it directly...
		wait := make(chan []interface{})
		go func(wait chan []interface{}) {
			var results_ []interface{}
			// In case of error.
			defer func(wait chan []interface{}, results *[]interface{}) { //catch or finally
				if err := recover(); err != nil { //catch
					log.Println(os.Stderr, "Exception: %v\n", err)
					*results = []interface{}{nil, err}
				}
				wait <- *results
			}(wait, &results_)

			results := finalMethod.Call(in)
			if len(results) > 0 {
				if len(results) == 1 {
					// One result here...
					switch goError := results[0].Interface().(type) {
					case error:
						results_ = []interface{}{nil, goError}
						return
					}
					results_ = []interface{}{results[0].Interface(), nil}
					return
				} else if len(results) == 2 {
					// Return the result and the error after.
					results_ = []interface{}{results[0].Interface(), results[1].Interface()}
					return
				}
			}
			results_ = []interface{}{"", nil}
		}(wait)

		// wait for the results...
		results := <-wait
		return results[0], results[1]

	} else {
		return nil, errors.New("Method dosen't exist!")
	}

	// return or panic, method not found of either type
	return "", nil
}
