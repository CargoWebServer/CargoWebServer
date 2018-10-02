package GoJavaScript

import "reflect"
import "code.myceliUs.com/Utility"
import "log"

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
func RefToObject(ref interface{}) interface{} {
	if ref == nil {
		return nil
	}

	if reflect.TypeOf(ref).Kind() == reflect.Slice {
		// Here I got a slice...
		slice := reflect.ValueOf(ref)
		for i := 0; i < slice.Len(); i++ {
			e := slice.Index(i)
			if reflect.TypeOf(e.Interface()).String() == "GoJavaScript.ObjectRef" {
				// Replace the object reference with it actual object.
				ref.([]interface{})[i] = GetCache().GetObject(e.Interface().(ObjectRef).UUID)
			} else if reflect.TypeOf(e.Interface()).String() == "*GoJavaScript.ObjectRef" {
				// Replace the object reference with it actual object.
				ref.([]interface{})[i] = GetCache().GetObject(e.Interface().(*ObjectRef).UUID)
			}
		}
	} else {
		if reflect.TypeOf(ref).String() == "GoJavaScript.ObjectRef" {
			// Replace the object reference with it actual value.
			ref = GetCache().GetObject(ref.(ObjectRef).UUID)
		} else if reflect.TypeOf(ref).String() == "*GoJavaScript.ObjectRef" {
			// Replace the object reference with it actual value.
			ref = GetCache().GetObject(ref.(*ObjectRef).UUID)
		}
	}
	return ref
}

/**
 * Register a go type to be usable as JS type.
 */
func RegisterGoType(value interface{}) {
	// Register local object.
	if reflect.TypeOf(value).String() != "GoJavaScript.Object" && reflect.TypeOf(value).String() != "map[string]interface {}" {
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
				if e.IsNil() {
					objects_ = append(objects_, nil)
				} else {
					// I will derefence the pointer if it's a pointer.
					for reflect.TypeOf(e.Interface()).Kind() == reflect.Ptr {
						e = e.Elem()
					}
					if reflect.TypeOf(e.Interface()).Kind() == reflect.Struct {
						// results will be register.
						uuid := RegisterGoObject(slice.Index(i).Interface(), "")
						// I will set the results a object reference.
						objects_ = append(objects_, NewObjectRef(uuid))

					} else if reflect.TypeOf(e.Interface()).String() == "map[string]interface {}" {
						// Here if the object is Entity I will create it object and
						// return a reference to it.
						if e.Interface().(map[string]interface{})["TYPENAME"] != nil && e.Interface().(map[string]interface{})["__object_infos__"] == nil {
							// In case of object...
							typeName := e.Interface().(map[string]interface{})["TYPENAME"].(string)
							var uuid string
							if e.Interface().(map[string]interface{})["UUID"] != nil {
								uuid = e.Interface().(map[string]interface{})["UUID"].(string)
							} else {
								uuid = Utility.RandomUUID()
							}

							// In that case I will initialyse the object.
							obj := Utility.MakeInstance(typeName, e.Interface().(map[string]interface{}), func(interface{}) {})
							GetCache().SetObject(uuid, obj)

							// I will set the results a object reference.
							objects_ = append(objects_, NewObjectRef(uuid))
						} else {
							objects_ = append(objects_, slice.Index(i).Interface())
						}
					} else {
						objects_ = append(objects_, slice.Index(i).Interface())
					}
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
				} else if reflect.TypeOf(objects).String() == "map[string]interface {}" {
					// Here if the object is Entity I will create it object and
					// return a reference to it.
					if objects.(map[string]interface{})["TYPENAME"] != nil && objects.(map[string]interface{})["__object_infos__"] == nil {
						// In case of object...
						typeName := objects.(map[string]interface{})["TYPENAME"].(string)

						var uuid string
						if objects.(map[string]interface{})["UUID"] != nil {
							uuid = objects.(map[string]interface{})["UUID"].(string)
						} else {
							uuid = Utility.RandomUUID()
						}

						// In that case I will initialyse the object.
						obj := Utility.MakeInstance(typeName, objects.(map[string]interface{}), func(interface{}) {})
						GetCache().SetObject(uuid, obj)

						// I will set the results a object reference.
						objects = NewObjectRef(uuid)
					}

				}
			}
		}
	}

	log.Println("225 ---> objects: ", objects)

	return objects
}
