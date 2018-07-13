// TestJerryScriot project main.go
package main

import (
	//"log"

	"code.myceliUs.com/GoJerryScript"
)

func main() {
	engine := GoJerryScript.NewEngine(9696, GoJerryScript.JERRY_INIT_EMPTY)

	engine.AppendFunction("TestArray", []string{"arr, val"}, "function TestArray(arr, val){arr.push(val);return arr;}", GoJerryScript.JERRY_PARSE_NO_OPTS)
	engine.EvalScript("TestArray(arr, val);", []GoJerryScript.Variable{{Name: "arr", Value: []float64{1.0, 3.0, 4.0}}, {Name: "val", Value: 2.25}})

	engine.Clear()
}
