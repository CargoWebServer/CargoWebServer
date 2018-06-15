package GoJerryScript

import "testing"

func TestHelloJerry(t *testing.T) {
	// Init the script engine.
	Jerry_init(Jerry_init_flag_t(JERRY_INIT_EMPTY))

	/* Register 'print' function from the extensions */

	var arg0 Uint8
	var arg1 *_swig_fnptr

	Jerryx_handler_register_global(arg0, arg1)

	// Cleanup the script engine.
	Jerry_cleanup()
}
