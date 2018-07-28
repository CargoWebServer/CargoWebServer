package GoJerryScript

import (
	"C"
	"errors"

	"code.myceliUs.com/Utility"
)

/**
 * Init and start the engine.
 * port The port to communicate with the engine, or the debbuger.
 * option Can be JERRY_INIT_EMPTY, JERRY_INIT_SHOW_OPCODES, JERRY_INIT_SHOW_REGEXP_OPCODES
 *		  JERRY_INIT_MEM_STATS, 	JERRY_INIT_MEM_STATS_SEPARATE (or a combination option)
 */
func (self *Engine) start(port int, options int) {
	/* The port to communicate with the instance */
	self.port = port

	/* Init the script engine. */
	Jerry_init(Jerry_init_flag_t(options))
}

/**
 * Parse and set a function in the Javascript.
 * name The name of the function (the function will be keep in the engine for it
 *      lifetime.
 * args The argument name for that function.
 * src  The body of the function
 * options Can be JERRY_PARSE_NO_OPTS or JERRY_PARSE_STRICT_MODE
 */
func (self *Engine) AppendFunction(name string, args []string, src string, options int) error {
	/* The name of the function */
	arg0 := NewUint8FromString(name)
	args_ := ""
	for i := 0; i < len(args); i++ {
		args_ += args[i]
		if i < len(args)-1 {
			args_ += ", "
		}
	}
	arg1 := NewUint8FromString(args_)
	arg2 := NewUint8FromString(src)
	arg3 := NewUint32FromInt(int32(options))

	parsed_code := Jerry_parse_function(arg0, int64(len(name)), arg1, int64(len(args_)), arg2, int64(len(src)), arg3)

	if !Jerry_value_is_error(parsed_code) {
		self.functions[name] = Function{Name: name, Args: args, Body: src, Obj: parsed_code}
		_, err := self.EvalScript(src, []Variable{})
		return err
	} else {
		Jerry_release_value(parsed_code)
		return errors.New("Fail to parse function " + name)
	}
}

/**
 * Register a go type to be usable as JS type.
 */
func (self *Engine) RegisterGoType(value interface{}) {
	Utility.RegisterType(value)
}

/**
 * Set a variable on the global context.
 * name The name of the variable in the context.
 * value The value of the variable, can be a string, a number,
 */
func (self *Engine) SetGlobalVariable(name string, value interface{}) {
	// first of all I will initialyse the arguments.
	globalObject := Jerry_get_global_object()

	propName := Jerry_create_string(NewUint8FromString(name))

	// Take a go value and initialyse a Uint32 representation.
	propValue := goToJs(value)

	// Set the propertie in the global context..
	Jerry_set_property(globalObject, propName, propValue)

	// Release the resource as no more needed here.
	Jerry_release_value(propName)
	Jerry_release_value(propValue)

	Jerry_release_value(globalObject)
}

/**
 * Evaluate a script.
 * script Contain the code to run.
 * variables Contain the list of variable to set on the global context before
 * running the script.
 */
func (self *Engine) EvalScript(script string, variables []Variable) (interface{}, error) {
	// Here the values are put on the global contex before use in the function.
	for i := 0; i < len(variables); i++ {
		self.SetGlobalVariable(variables[i].Name, variables[i].Value)
	}

	// Now I will evaluate the function...
	ret := Jerry_eval(NewUint8FromString(script), int64(len(script)), false)

	// Convert Js value to Go value.
	value, err := jsToGo(ret)

	// Free JavaScript value, returned by eval
	Jerry_release_value(ret)
	return value, err
}

func (self *Engine) Clear() {
	for _, fct := range self.functions {
		Jerry_release_value(fct.Obj)
	}
	/* Cleanup the script engine. */
	Jerry_cleanup()

}
