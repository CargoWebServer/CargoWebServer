package Utility

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/gob"
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"
)

/**
 * Use it to get UUID of referenceable object.
 */
type Referenceable interface {
	GetUUID() string
}

/**
 * That map will contain the list of all type to be created dynamicaly
 */
var typeRegistry = make(map[string]reflect.Type)

func GetTypeOf(typeName string) reflect.Type {
	if t, ok := typeRegistry[typeName]; ok {
		return reflect.New(t).Type()
	}
	return nil
}

/**
 * Register an instance of the type.
 */
func RegisterType(typedNil interface{}) {
	t := reflect.TypeOf(typedNil).Elem()
	index := strings.LastIndex(t.PkgPath(), "/")
	var typeName = t.Name()
	if _, ok := typeRegistry[t.PkgPath()[index+1:]+"."+typeName]; !ok {
		if index > 0 {
			typeRegistry[t.PkgPath()[index+1:]+"."+typeName] = t
			gob.RegisterName(t.PkgPath()[index+1:]+"."+typeName, typedNil)
			log.Println("------> type: ", t.PkgPath()[index+1:]+"."+typeName, " was register as dynamic type.")
		} else {
			typeRegistry[t.PkgPath()+"."+typeName] = t
			gob.RegisterName(t.PkgPath()+"."+typeName, typedNil)
			log.Println("------> type: ", t.PkgPath()+"."+typeName, " was register as dynamic type.")
		}
	}
}

func toInt(value interface{}) int {
	switch v := value.(type) {
	case string:
		result, err := strconv.Atoi(v)
		if err == nil {
			return int(result)
		}
		return 0
	case int8:
		return int(value.(int8))
	case int16:
		return int(value.(int16))
	case int32:
		return int(value.(int32))
	case int64:
		return int(value.(int64))
	case uint8:
		return int(value.(uint8))
	case uint16:
		return int(value.(uint16))
	case uint32:
		return int(value.(uint32))
	case uint64:
		return int(value.(uint64))
	case float64:
		return int(value.(float64))
	case float32:
		return int(value.(float32))
	default:
		return 0
	}
}

/**
 * Initialyse base type value.
 */
func InitializeBaseTypeValue(t reflect.Type, value interface{}) reflect.Value {
	if value == nil {
		return reflect.ValueOf(nil)
	}

	var v reflect.Value

	switch t.Kind() {
	case reflect.String:
		// Here it's possible that the value contain the map of values...
		// I that case I will
		v = reflect.ValueOf(value.(string))
	case reflect.Bool:
		if value != nil {
			v = reflect.ValueOf(value.(bool))
		} else {
			v = reflect.ValueOf(false)
		}
	case reflect.Int:
		v = reflect.ValueOf(toInt(value))
	case reflect.Int8:
		v = reflect.ValueOf(toInt(value))
	case reflect.Int32:
		v = reflect.ValueOf(toInt(value))
	case reflect.Int64:
		v = reflect.ValueOf(toInt(value))
	case reflect.Uint:
		v = reflect.ValueOf(value.(uint))
	case reflect.Uint8:
		v = reflect.ValueOf(value.(uint8))
	case reflect.Uint32:
		v = reflect.ValueOf(value.(uint32))
	case reflect.Uint64:
		v = reflect.ValueOf(value.(uint64))
	case reflect.Float32:
		v = reflect.ValueOf(value.(float32))
	case reflect.Float64:
		v = reflect.ValueOf(value.(float64))
	case reflect.Array:
		log.Panicln("--------> array found!")
	default:
		log.Println("unexpected type %T\n", t)
	}

	return v
}

/**
 * Create an instance of the type with it name.
 */
func MakeInstance(typeName string, data map[string]interface{}, setEntity func(interface{})) reflect.Value {
	value := initializeStructureValue(typeName, data, setEntity)
	if setEntity != nil {
		setEntity(value.Interface())
	}
	return value
}

/**
 * Intialyse the struct fields with the values contain in the map.
 */
func initializeStructureValue(typeName string, data map[string]interface{}, setEntity func(interface{})) reflect.Value {
	// Here I will create the value...
	t := typeRegistry[typeName]
	if t == nil {
		return reflect.ValueOf(data)
	}
	v := reflect.New(t)
	for name, value := range data {
		ft, exist := t.FieldByName(name)
		if exist && value != nil {
			initializeStructureFieldValue(v, name, ft.Type, value, setEntity)
		}
	}

	// Return the initialysed value...
	return v
}

// Return an initialyse field value.
func InitializeStructureFieldArrayValue(fieldName string, fieldType reflect.Type, values reflect.Value, setEntity func(interface{})) reflect.Value {
	// Create a slice here.
	slice := reflect.MakeSlice(fieldType, values.Len(), values.Len())

	// Here I will iterate over the slice
	for i := 0; i < values.Len(); i++ {
		// the slice value.
		v_ := values.Index(i).Interface()

		// If the type is register as dynamic type.
		if v_ != nil {
			if reflect.TypeOf(v_).String() == "map[string]interface {}" {
				// The value is a dynamic value.
				v_ := v_.(map[string]interface{})
				if v_["TYPENAME"] != nil {
					fv := initializeStructureValue(v_["TYPENAME"].(string), v_, setEntity)
					// I will set the reference in the parent object.
					setEntity(fv.Interface())

					// Special case for entity (Cargo)...
					if strings.HasPrefix(fieldName, "M_") {
						if v_["UUID"] != nil {
							index := slice.Index(i)
							index.Set(reflect.ValueOf(v_["UUID"].(string)))
						}
					} else {
						index := slice.Index(i)
						index.Set(fv)
					}

				} else {
					// A generic map not a dynamic type.
					index := slice.Index(i)
					index.Set(reflect.ValueOf(v_))
				}
			} else {
				// Not an array of map[string]interface {}
				index := slice.Index(i)
				if reflect.TypeOf(v_).Kind() == reflect.Slice {
					slice := InitializeStructureFieldArrayValue(fieldName, reflect.TypeOf(v_), reflect.ValueOf(v_), setEntity)
					// Set the slice...
					if slice.IsValid() {
						// A sub-array.
						index.Set(slice)
					}
				} else {
					// Not an array
					fv := InitializeBaseTypeValue(reflect.TypeOf(v_), v_).Convert(reflect.TypeOf(v_))
					if fv.IsValid() {
						index.Set(fv)
					}
				}
			}
		}
	}
	return slice
}

func initializeStructureFieldValue(v reflect.Value, fieldName string, fieldType reflect.Type, fieldValue interface{}, setEntity func(interface{})) {
	switch fieldType.Kind() {
	case reflect.Slice:
		// That's mean the value contain an array...
		if reflect.TypeOf(fieldValue).String() == "[]uint8" || reflect.TypeOf(fieldValue).String() == "[]byte" {
			fv := InitializeBaseTypeValue(reflect.TypeOf(fieldValue), fieldValue)
			val := fv.String()
			val_, err := b64.StdEncoding.DecodeString(val)
			if err == nil {
				val = string(val_)
			}
			// Set the value...
			v.Elem().FieldByName(fieldName).Set(reflect.ValueOf([]byte(val)))
		} else {
			slice := InitializeStructureFieldArrayValue(fieldName, fieldType, reflect.ValueOf(fieldValue), setEntity)
			// Set the slice...
			if slice.IsValid() {
				v.Elem().FieldByName(fieldName).Set(slice)
			}
		}

	case reflect.Struct:
		fv, _ := InitializeStructure(fieldValue.(map[string]interface{}), setEntity)
		if fv.IsValid() {
			v.Elem().FieldByName(fieldName).Set(fv.Elem())
		}
	case reflect.Ptr:
		fv, _ := InitializeStructure(fieldValue.(map[string]interface{}), setEntity)
		if fv.IsValid() {
			v.Elem().FieldByName(fieldName).Set(fv)
		}
	case reflect.Interface:
		// Here the type of the actual value will determine the value to initialyse...
		initializeStructureFieldValue(v, fieldName, reflect.TypeOf(fieldValue), fieldValue, setEntity)

	case reflect.Map:
		fv, err := InitializeStructure(fieldValue.(map[string]interface{}), setEntity)
		if err == nil {
			if fv.IsValid() {
				v.Elem().FieldByName(fieldName).Set(fv.Elem())
			}
		} else {
			// In that case I dont have a map with a define type so i will simply set
			// the actual value to the field.
			v.Elem().FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
		}
	default:

		// Convert is use to enumeration type who are int and must be convert to
		// it const type representation.
		fv := InitializeBaseTypeValue(fieldType, fieldValue).Convert(fieldType)
		if fv.IsValid() {
			v.Elem().FieldByName(fieldName).Set(fv)
		}
	}

}

/**
 * Initialyse an array of structures, return it as interface (array of the actual
 * objects)
 */
func InitializeStructures(data []interface{}, typeName string, setEntity func(interface{})) (reflect.Value, error) {
	// Here I will get the type name, only dynamic type can be use here...
	var values reflect.Value
	if len(data) > 0 {
		// Structure data must be a map[string]interface{}
		if _, ok := data[0].(map[string]interface{}); ok {
			if typeName_, ok := data[0].(map[string]interface{})["TYPENAME"]; ok {
				// Now I will create empty structure and initialyse it with the value found in the map values.
				for i := 0; i < len(data); i++ {
					obj := MakeInstance(typeName_.(string), data[i].(map[string]interface{}), setEntity)
					if i == 0 {
						if len(typeName) == 0 {
							values = reflect.MakeSlice(reflect.SliceOf(obj.Type()), 0, 0)
						} else if t, ok := typeRegistry[typeName]; ok {
							values = reflect.MakeSlice(reflect.SliceOf(reflect.New(t).Type()), 0, 0)
						} else {
							emptyInterfaceArray := make([]interface{}, 0, 0)
							values = reflect.ValueOf(emptyInterfaceArray)
						}
					}
					values = reflect.Append(values, obj)
				}
				return values, nil
			} else {
				return reflect.ValueOf(data), nil
			}
		} else {
			return values, errors.New("NotDynamicObject")
		}
	} else {
		// Here there is no value in the array.
		if t, ok := typeRegistry[typeName]; ok {
			values = reflect.MakeSlice(reflect.SliceOf(reflect.New(t).Type()), 0, 0)
		} else {
			emptyInterfaceArray := make([]interface{}, 0, 0)
			values = reflect.ValueOf(emptyInterfaceArray)
		}
	}
	return values, nil
}

/**
 * Serialyse the entity to a byte array.
 */
func ToBytes(val interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(val)
	return buf.Bytes(), err
}

/**
 * Read entity from byte array.
 */
func FromBytes(data []byte, typeName string) (interface{}, error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	if t, ok := typeRegistry[typeName]; ok {
		v := reflect.New(t).Interface()
		err := dec.Decode(v)
		return v, err
	} else {
		v := make(map[string]interface{})
		err := dec.Decode(&v)
		return v, err
	}
	return nil, errors.New("Fail to instantiate value!")
}

/**
 * Initialyse a single object from it value.
 */
func InitializeStructure(data map[string]interface{}, setEntity func(interface{})) (reflect.Value, error) {
	// Here I will get the type name, only dynamic type can be use here...
	var value reflect.Value
	if typeName, ok := data["TYPENAME"]; ok {
		if _, ok := typeRegistry[typeName.(string)]; ok {
			value = MakeInstance(typeName.(string), data, setEntity)
			setEntity(value.Interface())
			return value, nil
		} else {
			// Return the value itself...
			return reflect.ValueOf(data), nil
		}
	} else {
		return value, errors.New("NotDynamicObject")
	}
}

/**
 * Initialyse an array of values other than structure...
 */
func InitializeArray(data []interface{}) (reflect.Value, error) {
	var values reflect.Value

	sameType := true
	if len(data) > 1 {
		for i := 1; i < len(data) && sameType; i++ {
			if data[i] != nil {
				sameType = reflect.TypeOf(data[i]).String() == reflect.TypeOf(data[i-1]).String()
			}
		}
	}

	for i := 0; i < len(data); i++ {
		if data[i] != nil {
			if i == 0 {
				if sameType {
					values = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(data[i])), 0, 0)
				} else {
					emptyInterfaceArray := make([]interface{}, 0, 0)
					values = reflect.ValueOf(emptyInterfaceArray)
				}
			}
			values = reflect.Append(values, reflect.ValueOf(data[i]))
		}
	}
	return values, nil
}
