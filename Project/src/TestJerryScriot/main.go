// TestJerryScriot project main.go
package main

import (
	//"log"

	"code.myceliUs.com/GoJerryScript"
)

func main() {
	engine := GoJerryScript.NewEngine(9696, GoJerryScript.JERRY_INIT_EMPTY)

	// eval string value
	engine.AppendFunction("SayHelloTo", []string{"to"}, "print ('Hello ' + to);", GoJerryScript.JERRY_PARSE_NO_OPTS)
	engine.EvalFunction("SayHelloTo", []interface{}{"Jerry"})

	// eval numeric value
	engine.AppendFunction("Add", []string{"a", "b"}, "a+b;", GoJerryScript.JERRY_PARSE_NO_OPTS)
	engine.EvalFunction("Add", []interface{}{1, 2.25})

	engine.Clear()
}
