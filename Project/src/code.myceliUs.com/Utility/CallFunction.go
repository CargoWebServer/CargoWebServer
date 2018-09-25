package Utility

import (
	"errors"
	"reflect"
	"strconv"
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

/**
 * Get a function from it name.
 */
func GetFunction(name string) interface{} {
	return functionRegistry[name]
}

func CallFunction(name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(functionRegistry[name])
	if len(params) != f.Type().NumIn() && !f.Type().IsVariadic() {
		err = errors.New("Wrong number of parameter for " + name + " got " + strconv.Itoa(len(params)) + " but espect " + strconv.Itoa(f.Type().NumIn()))
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)

	return
}
