package GoJavaScript

/**
 * Variable with name and value.
 */
type Variable struct {
	// The typename.
	TYPENAME string
	// The variable name
	Name string
	// It value.
	Value interface{}
}

/**
 * Create a new variable type.
 */
func NewVariable(name string, value interface{}) *Variable {
	v := new(Variable)
	v.TYPENAME = "GoJavaScript.Variable"
	v.Name = name
	v.Value = value
	return v
}

/**
 * Generic 32bit reference pointer. (unit 32)
 */
type Uint32_t interface {
	Swigcptr() uintptr
}

// Go Object reference.
type ObjectRef struct {
	// The uuid of the referenced object.
	UUID     string
	TYPENAME string
}

// Contain byte code.
type ByteCode struct {
	TYPENAME string
	Data     []uint8
}
