// TestJerryScriot project main.go
package main

import (
	//"log"

	"code.myceliUs.com/GoJerryScript"
)

func main() {
	engine := GoJerryScript.NewEngine(9696, GoJerryScript.JERRY_INIT_EMPTY)
	engine.AppendFunction("SayHelloTo", []string{"to"}, "print ('Hello ' + to);", GoJerryScript.JERRY_PARSE_NO_OPTS)
	engine.EvalFunction("SayHelloTo", []string{"Jerry"})
	engine.Clear()
}
