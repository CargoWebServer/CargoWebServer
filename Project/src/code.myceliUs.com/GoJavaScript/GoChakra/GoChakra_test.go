package GoChakra

import "testing"

import "log"

type Dog struct {
	Name string
	Race string
	Age  int
}

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

// Test setting various data type.
func TestSetVariables(t *testing.T) {

	// Test use to create type go/js type convesion...
	// Set a global variable.
	engine.SetGlobalVariable("toto", "this is a string value")
	// get it back...
	val, err := engine.GetGlobalVariable("toto")

	if err == nil {
		val_, _ := val.Export()
		if val_.(string) != "this is a string value" {
			t.Error("---> fail to get string value!")
		}
	}

	// Test with an array
	engine.SetGlobalVariable("arr", []interface{}{"val 1", 1.0, 1, true})
	arr, err := engine.GetGlobalVariable("arr")

	/*if err == nil {
		arr_, _ := arr.Export()
		if arr_.([]interface{})[0] != "val 1" && arr_.([]interface{})[1] != 1.0 &&
			arr_.([]interface{})[2] != 1 && arr_.([]interface{})[3] != true {
			t.Error("---> fail to set and get array")
		}
	}*/
	log.Println(arr)

	// Test exporting an object.
	engine.SetGlobalVariable("ada", Dog{Name: "ada", Race: "Border Collie", Age: 4})

	// Return a map in case of unregistered go type...
	ada, _ := engine.GetGlobalVariable("ada")
	log.Println(ada)

}

// Test runing simple script...
func TestRunSimpleJsFunction(t *testing.T) {

	engine.RegisterJsFunction("SayHelloTo", "function SayHelloTo(greething, to){return greething + ' ' + to + '!';}")
	str, _ := engine.CallFunction("SayHelloTo", []interface{}{"Hello", "Java Script"})
	str_, _ := str.ToString()

	if str_ != "Hello Java Script!" {
		t.Error("Expected 'Hello Java Script!', got ", str)
	} else {
		// display hello jerry!
		log.Println(str_)
	}

}
