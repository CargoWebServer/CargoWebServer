package GoJerryScript

import "testing"
import "log"

func TestHelloJerry(t *testing.T) {
	// Init the script engine.
	Jerry_init(Jerry_init_flag_t(JERRY_INIT_EMPTY))

	/* Register 'print' function  */
	RegisterPrintHandler()

	//str := "function add(a, b){return a+b;}; add(1, 2);"
	str := "print ('Hello, World!');"
	var arg0 Uint8                     // nil pointer
	var arg1 = int64(0)                // 0 length
	var arg2 = NewUint8FromString(str) // The script.
	var arg3 = int64(len(str) + 1)
	var arg4 = NewUint32FromInt(int32(JERRY_PARSE_NO_OPTS))

	parsed_code := Jerry_parse(arg0, arg1, arg2, arg3, arg4)
	if !Jerry_value_is_error(parsed_code) {

		/* Execute the parsed source code in the Global scope */
		jerry_value_t := Jerry_run(parsed_code)

		/* Returned value must be freed */
		Jerry_release_value(jerry_value_t)
	}

	/* Parsed source code must be freed */
	Jerry_release_value(parsed_code)

	// Cleanup the script engine.
	Jerry_cleanup()
}

func TestEvalFunction(t *testing.T) {
	engine := NewEngine(9696, JERRY_INIT_EMPTY)

	// Test eval string function.
	engine.AppendFunction("SayHelloTo", []string{"to"}, "'Hello ' + to + '!'", JERRY_PARSE_NO_OPTS)
	str, _ := engine.EvalFunction("SayHelloTo", []interface{}{"Jerry Script"})
	log.Println("---> ", str)

	// Test numeric function
	engine.AppendFunction("Add", []string{"a", "b"}, "a + b;", JERRY_PARSE_NO_OPTS)
	number, _ := engine.EvalFunction("Add", []interface{}{1, 2.25})
	log.Println("---> ", number)

	engine.Clear()
}
