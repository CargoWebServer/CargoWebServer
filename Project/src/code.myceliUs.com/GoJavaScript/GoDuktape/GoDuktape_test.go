package GoDuktape

import "testing"
import "log"

var (
	engine = new(Engine)
)

/**
 * Test chakra function alone. Use GoJavascriptTest to test it in the context of an
 * application.
 */
func TestStart(t *testing.T) {
	engine.Start(8081)
}

func TestGlobalVariable(t *testing.T) {

	// First of all I will register the tow go function in the Engine.
	var toto = "This is JavaScript"
	engine.SetGlobalVariable("toto", toto)

	toto_, _ := engine.GetGlobalVariable("toto")
	toto__, _ := toto_.Export()

	if toto__ != toto {
		t.Error("Set/Get global variables fail! ")
	}

	// Here I will test for boolean value...
	isTrue := true
	engine.SetGlobalVariable("isTrue", isTrue)
	isTrue_, _ := engine.GetGlobalVariable("isTrue")
	if !isTrue_.Val.(bool) {
		t.Error("is not true")
	}

	isNumber := 0.125
	engine.SetGlobalVariable("isNumber", isNumber)

	isNumber_, _ := engine.GetGlobalVariable("isNumber")
	if isNumber_.Val.(float64) != 0.125 {
		t.Error("is not a number!")
	}

	isArray := []interface{}{toto, isTrue, isNumber, 10}
	engine.SetGlobalVariable("isArray", isArray)

	isArray_, _ := engine.GetGlobalVariable("isArray")
	log.Println("--> is array: ", isArray_)

}

/**
 * Stop the engine.
 */
func TestClear(t *testing.T) {
	engine.Clear()
}
