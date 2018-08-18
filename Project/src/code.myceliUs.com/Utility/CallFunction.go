package Utility

import (
	"errors"
	"reflect"
)

/**
 * That map will contain the list of function callable dynamicaly.
 */
var functionRegistry = make(map[string]interface{})

/**
 * Register a Go function to be able to call it dynamically latter.
 */
func RegisterFunction(name string, fct interface{}) {
	// keep the function in the map.
	functionRegistry[name] = fct
}

func CallFunction(name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(functionRegistry[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)

	return
}
