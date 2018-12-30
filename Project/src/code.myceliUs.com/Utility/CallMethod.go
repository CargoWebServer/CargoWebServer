package Utility

import (
	"errors"
	"log"

	//"os"
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
		} else if !method.Type().IsVariadic() {
			in[k] = reflect.Zero(method.Type().In(k))
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
					*results = []interface{}{nil, err}
				}
				wait <- *results
			}(wait, &results_)

			results := finalMethod.Call(in)

			if len(results) > 0 {
				if len(results) == 1 {
					// One result here...
					result := results[0]
					zeroValue := reflect.Zero(result.Type())
					if result.IsValid() {
						if result != zeroValue {
							switch goError := result.Interface().(type) {
							case error:
								results_ = []interface{}{nil, goError}
								return
							}
							results_ = []interface{}{result.Interface(), nil}
						}
					}
					return
				} else if len(results) == 2 {
					// Return the result and the error after.
					result0 := results[0]
					zeroValue0 := reflect.Zero(result0.Type())
					result1 := results[0]
					zeroValue1 := reflect.Zero(result1.Type())
					if result0 != zeroValue0 && result1 != zeroValue1 {
						results_ = []interface{}{results[0].Interface(), results[1].Interface()}
					} else if result0 == zeroValue0 && result1 != zeroValue1 {
						results_ = []interface{}{nil, results[1].Interface()}
					} else if result0 != zeroValue0 && result1 == zeroValue1 {
						results_ = []interface{}{results[0].Interface(), nil}
					}
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

func SetProperty(i interface{}, field string, value interface{}) error {
	// pointer to struct - addressable
	ps := reflect.ValueOf(i)
	// struct
	s := ps.Elem()

	if s.Kind() == reflect.Struct {
		// exported field
		f := s.FieldByName(field)
		if f.IsValid() {
			// A Value can be changed only if it is
			// addressable and was not obtained by
			// the use of unexported struct fields.
			if f.CanSet() {
				// change value of N
				if f.Kind() == reflect.TypeOf(value).Kind() {
					f.Set(reflect.ValueOf(value))
				} else {
					log.Println("---> field ", field, " has type ", f.Kind(), " not ", reflect.TypeOf(value).Kind())
				}
			} else {
				log.Println("---> field ", field, " can not be set!")
			}
		} else {
			log.Println("---> field ", field, " is no valid!")
		}
	}
	return nil
}

func GetProperty(i interface{}, field string) interface{} {

	ps := reflect.ValueOf(i)
	// struct
	s := ps.Elem()

	// exported field
	f := s.FieldByName(field)
	if f.IsValid() {
		// Return value as interface.
		return f.Interface()
	}

	return nil
}
