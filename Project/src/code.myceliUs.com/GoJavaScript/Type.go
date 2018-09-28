package GoJavaScript

import "reflect"
import "code.myceliUs.com/Utility"

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

// Replace object reference by actual object.
func RefToObject(values []interface{}) []interface{} {
	for i := 0; i < len(values); i++ {
		if values[i] != nil {
			if reflect.TypeOf(values[i]).Kind() == reflect.Slice {
				// Here I got a slice...
				slice := reflect.ValueOf(values[i])
				for j := 0; j < slice.Len(); j++ {
					e := slice.Index(j)
					if reflect.TypeOf(e.Interface()).String() == "GoJavaScript.ObjectRef" {
						// Replace the object reference with it actual object.
						values[i].([]interface{})[j] = GetCache().GetObject(e.Interface().(ObjectRef).UUID)
					} else if reflect.TypeOf(e.Interface()).String() == "*GoJavaScript.ObjectRef" {
						// Replace the object reference with it actual object.
						values[i].([]interface{})[j] = GetCache().GetObject(e.Interface().(*ObjectRef).UUID)
					}
				}
			} else {
				if reflect.TypeOf(values[i]).String() == "GoJavaScript.ObjectRef" {
					// Replace the object reference with it actual value.
					values[i] = GetCache().GetObject(values[i].(ObjectRef).UUID)
				} else if reflect.TypeOf(values[i]).String() == "*GoJavaScript.ObjectRef" {
					// Replace the object reference with it actual value.
					values[i] = GetCache().GetObject(values[i].(*ObjectRef).UUID)
				}
			}
		}
	}

	return values
}

/**
 * Register a go type to be usable as JS type.
 */
func RegisterGoType(value interface{}) {
	// Register local object.
	if reflect.TypeOf(value).String() != "GoJavaScript.Object" {
		Utility.RegisterType(value)
	}
}

/**
 * Localy register a go object.
 */
func RegisterGoObject(obj interface{}, name string) string {
	// Here I will dynamicaly register objet type in the utility cache...
	empty := reflect.New(reflect.TypeOf(obj))
	RegisterGoType(empty.Elem().Interface())

	// Random uuid.
	var uuid string
	if len(name) > 0 {
		// In that case the object is in the global scope.
		uuid = Utility.GenerateUUID(name)
	} else {
		// Not a global object.
		uuid = Utility.RandomUUID()
	}

	// Do not recreate already existing object.
	if GetCache().GetObject(uuid) != nil {
		return uuid
	}

	// Here I will keep the object in the client cache.
	GetCache().SetObject(uuid, obj)

	return uuid
}

/**
 * Convert object, or objects to their uuid reference.
 */
func ObjectToRef(objects interface{}) interface{} {
	if objects != nil {
		if reflect.TypeOf(objects).Kind() == reflect.Slice {
			// Here if the result is a slice I will test if it contains struct...
			slice := reflect.ValueOf(objects)
			objects_ := make([]interface{}, 0)
			for i := 0; i < slice.Len(); i++ {
				e := slice.Index(i)
				// I will derefence the pointer if it's a pointer.
				for reflect.TypeOf(e.Interface()).Kind() == reflect.Ptr {
					e = e.Elem()
				}
				if reflect.TypeOf(e.Interface()).Kind() == reflect.Struct {
					// results will be register.
					uuid := RegisterGoObject(slice.Index(i).Interface(), "")
					// I will set the results a object reference.
					objects_ = append(objects_, NewObjectRef(uuid))
				} else {
					objects_ = append(objects_, slice.Index(i).Interface())
				}
			}
			// Set the array of object references.
			objects = objects_
		} else {
			// I will test if the result is a structure or not...
			if objects != nil {
				if reflect.TypeOf(objects).Kind() == reflect.Ptr || reflect.TypeOf(objects).Kind() == reflect.Struct {
					e := reflect.ValueOf(objects)
					if e.IsValid() {
						if reflect.TypeOf(objects).Kind() == reflect.Ptr {
							if e.IsNil() {
								return objects
							}
						}
						// I will derefence the pointer if it a pointer.
						for reflect.TypeOf(e.Interface()).Kind() == reflect.Ptr {
							e = e.Elem()
						}
						// if the object is a structure.
						if reflect.TypeOf(e.Interface()).Kind() == reflect.Struct {
							// results will be register.
							uuid := RegisterGoObject(objects, "")
							// I will set the results a object reference.
							objects = NewObjectRef(uuid)
						}
					}
				}
			}
		}
	}

	return objects
}
