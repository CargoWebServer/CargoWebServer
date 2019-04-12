package Utility

import (
	"reflect"
)

type TypeManager struct {
	/**
	 * That map will contain the list of all type to be created dynamicaly
	 */
	typeRegistry map[string]reflect.Type

	/**
	 * That map will contain the list of function callable dynamicaly.
	 */
	functionRegistry map[string]interface{}

	/**
	 * The channel to be call to access internal map.
	 */
	operationChannel chan map[string]interface{}
}

func (this *TypeManager) getType(name string) (t reflect.Type, exist bool) {
	op := make(map[string]interface{})
	op["opName"] = "getType"
	op["name"] = name
	op["result"] = make(chan interface{}, 0)
	this.operationChannel <- op

	// return the result.
	result := <-op["result"].(chan interface{})
	if result == nil {
		return t, false
	}

	return result.(reflect.Type), true
}

func (this *TypeManager) setType(name string, val reflect.Type) {
	op := make(map[string]interface{})
	op["opName"] = "setType"
	op["name"] = name
	op["value"] = val
	op["result"] = make(chan bool, 0)

	this.operationChannel <- op

	// wait before return.
	<-op["result"].(chan bool)

	return
}

func (this *TypeManager) getFunction(name string) interface{} {

	op := make(map[string]interface{})
	op["opName"] = "getFunction"
	op["name"] = name
	op["result"] = make(chan interface{}, 0)

	this.operationChannel <- op

	// return the result.
	result := <-op["result"].(chan interface{})
	return result
}

func (this *TypeManager) setFunction(name string, val interface{}) {
	op := make(map[string]interface{})
	op["opName"] = "setFunction"
	op["name"] = name
	op["value"] = val
	op["result"] = make(chan bool, 0)

	this.operationChannel <- op

	// wait before return.
	<-op["result"].(chan bool)

	return
}

// the singleton.
var typeManager *TypeManager

/**
 * The singleton package.
 */
func getTypeManager() *TypeManager {

	if typeManager != nil {
		return typeManager
	}

	typeManager = new(TypeManager)

	typeManager.typeRegistry = make(map[string]reflect.Type, 0)
	typeManager.functionRegistry = make(map[string]interface{}, 0)
	typeManager.operationChannel = make(chan map[string]interface{}, 0)

	// process type manager call here.
	go func() {
		for {
			select {
			case op := <-typeManager.operationChannel:

				if op["opName"] == "getType" {
					name := op["name"].(string)
					if t, ok := typeManager.typeRegistry[name]; ok {
						op["result"].(chan interface{}) <- t
					} else {
						op["result"].(chan interface{}) <- nil
					}

				} else if op["opName"] == "setType" {
					name := op["name"].(string)
					val := op["value"].(reflect.Type)
					typeManager.typeRegistry[name] = val
					op["result"].(chan bool) <- true

				} else if op["opName"] == "getFunction" {
					name := op["name"].(string)
					if fct, ok := typeManager.functionRegistry[name]; ok {
						op["result"].(chan interface{}) <- fct
					} else {
						op["result"].(chan interface{}) <- nil
					}
				} else if op["opName"] == "setFunction" {
					name := op["name"].(string)
					typeManager.functionRegistry[name] = op["value"]
					op["result"].(chan bool) <- true
				}
			}
		}
	}()
	return typeManager
}
