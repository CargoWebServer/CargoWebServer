package Utility

import (
	"errors"
	"reflect"
	"strconv"
)

/**
 * Register a Go function to be able to call it dynamically latter.
 */
func RegisterFunction(name string, fct interface{}) {
	// keep the function in the map.
	getTypeManager().setFunction(name, fct)
}

/**
 * Get a function from it name.
 */
func GetFunction(name string) interface{} {
	return getTypeManager().getFunction(name)
}

func CallFunction(name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(getTypeManager().getFunction(name))
	if len(params) != f.Type().NumIn() && !f.Type().IsVariadic() {
		err = errors.New("Wrong number of parameter for " + name + " got " + strconv.Itoa(len(params)) + " but espect " + strconv.Itoa(f.Type().NumIn()))
		return
	}
	in := make([]reflect.Value, len(params))

	for k, param := range params {
		if param != nil {
			in[k] = reflect.ValueOf(param)
		} else if !f.Type().IsVariadic() {
			in[k] = reflect.Zero(f.Type().In(k))
		} else {
			var nilVal interface{}
			in[k] = reflect.ValueOf(&nilVal).Elem()
		}
	}

	result = f.Call(in)

	return
}
