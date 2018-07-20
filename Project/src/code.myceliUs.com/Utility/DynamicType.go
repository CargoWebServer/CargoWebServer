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
func initializeBaseTypeValue(t reflect.Type, value interface{}) reflect.Value {
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
		v = reflect.ValueOf(value.(uint64))
	case reflect.Uint8:
		v = reflect.ValueOf(value.(uint64))
	case reflect.Uint32:
		v = reflect.ValueOf(value.(uint64))
	case reflect.Uint64:
		v = reflect.ValueOf(value.(uint64))
	case reflect.Float32:
		v = reflect.ValueOf(value.(float64))
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
			switch ft.Type.Kind() {
			case reflect.Slice:
				// That's mean the value contain an array...
				if strings.HasPrefix(reflect.TypeOf(value).String(), "[]") {
					if reflect.TypeOf(value).String() == "[]interface {}" {
						values := value.([]interface{})
						for i := 0; i < len(values); i++ {
							var fv reflect.Value
							switch v_ := values[i].(type) {
							// Here i have a sub-value.
							case map[string]interface{}:
								if v_["TYPENAME"] != nil {
									fv = initializeStructureValue(v_["TYPENAME"].(string), v_, setEntity)
									if fv.IsValid() {
										// Here I got an dynamic data type.
										setEntity(fv.Interface())
										// I will set the reference in the parent object.
										v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), reflect.ValueOf(v_["UUID"].(string))))
									}
								}
							default:
								// A base type...
								// Here I will try to convert the base type to the one I have in
								// the structure.
								if v.Elem().FieldByName(name).Type().String() == "[]string" {
									var string_ string
									fv = initializeBaseTypeValue(reflect.TypeOf(string_), v_)
								} else if v.Elem().FieldByName(name).Type().String() == "[]int" {
									var int_ int
									fv = initializeBaseTypeValue(reflect.TypeOf(int_), v_)
								} else if v.Elem().FieldByName(name).Type().String() == "[]float" {
									var float_ float64
									fv = initializeBaseTypeValue(reflect.TypeOf(float_), v_)
								} else if v.Elem().FieldByName(name).Type().String() == "[]bool" {
									var bool_ bool
									fv = initializeBaseTypeValue(reflect.TypeOf(bool_), v_)
								} else {
									fv = initializeBaseTypeValue(reflect.TypeOf(v_), v_)
								}
								if fv.IsValid() {
									v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
								}
							}

						}
					} else if reflect.TypeOf(value).String() == "[]string" {
						values := value.([]string)
						for i := 0; i < len(values); i++ {
							fv := initializeBaseTypeValue(reflect.TypeOf(values[i]), values[i])
							if fv.IsValid() {
								v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
							}
						}
					} else if reflect.TypeOf(value).String() == "[]int" {
						values := value.([]int)
						for i := 0; i < len(values); i++ {
							fv := initializeBaseTypeValue(reflect.TypeOf(values[i]), values[i])
							if fv.IsValid() {
								v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
							}
						}
					} else if reflect.TypeOf(value).String() == "[]int64" {
						values := value.([]int64)
						for i := 0; i < len(values); i++ {
							fv := initializeBaseTypeValue(reflect.TypeOf(values[i]), values[i])
							if fv.IsValid() {
								v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
							}
						}
					} else if reflect.TypeOf(value).String() == "[]int32" {
						values := value.([]int32)
						for i := 0; i < len(values); i++ {
							fv := initializeBaseTypeValue(reflect.TypeOf(values[i]), values[i])
							if fv.IsValid() {
								v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
							}
						}
					} else if reflect.TypeOf(value).String() == "[]int16" {
						values := value.([]int16)
						for i := 0; i < len(values); i++ {
							fv := initializeBaseTypeValue(reflect.TypeOf(values[i]), values[i])
							if fv.IsValid() {
								v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
							}
						}
					} else if reflect.TypeOf(value).String() == "[]int8" {
						values := value.([]int8)
						for i := 0; i < len(values); i++ {
							fv := initializeBaseTypeValue(reflect.TypeOf(values[i]), values[i])
							if fv.IsValid() {
								v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
							}
						}
					} else if reflect.TypeOf(value).String() == "[]bool" {
						values := value.([]bool)
						for i := 0; i < len(values); i++ {
							fv := initializeBaseTypeValue(reflect.TypeOf(values[i]), values[i])
							if fv.IsValid() {
								v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
							}
						}
					} else if reflect.TypeOf(value).String() == "[]float64" {
						values := value.([]bool)
						for i := 0; i < len(values); i++ {
							fv := initializeBaseTypeValue(reflect.TypeOf(values[i]), values[i])
							if fv.IsValid() {
								v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
							}
						}
					} else if reflect.TypeOf(value).String() == "[]float32" {
						values := value.([]float32)
						for i := 0; i < len(values); i++ {
							fv := initializeBaseTypeValue(reflect.TypeOf(values[i]), values[i])
							if fv.IsValid() {
								v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
							}
						}
					} else if reflect.TypeOf(value).String() == "[]uint8" || reflect.TypeOf(value).String() == "[]byte" {
						fv := initializeBaseTypeValue(reflect.TypeOf(value), value)
						val := fv.String()
						val_, err := b64.StdEncoding.DecodeString(val)
						if err == nil {
							val = string(val_)
						}
						// Set the value...
						v.Elem().FieldByName(name).Set(reflect.ValueOf([]byte(val)))
					}
				} else {
					// Here the value is a base type...
					fv := initializeBaseTypeValue(reflect.TypeOf(value), value)
					if fv.IsValid() {
						if ft.Type.String() != fv.Type().String() {
							// So here a conversion is necessary...
							if ft.Type.String() == "[]uint8" || ft.Type.String() == "[]byte" || fv.Type().String() == "string" {
								val := fv.String()
								log.Println("----> encode value found: ", name, val)
								val_, err := b64.StdEncoding.DecodeString(val)
								if err == nil {
									log.Println("----> decoded value: ", name, val)
									val = string(val_)
								}
								// Set the value...
								v.Elem().FieldByName(name).Set(reflect.ValueOf([]byte(val)))
							}
						} else {
							v.Elem().FieldByName(name).Set(fv)
						}
					}
				}
			case reflect.Struct:
				fv, _ := InitializeStructure(value.(map[string]interface{}), setEntity)
				if fv.IsValid() {
					v.Elem().FieldByName(name).Set(fv.Elem())
				}
			case reflect.Ptr:
				fv, _ := InitializeStructure(value.(map[string]interface{}), setEntity)
				if fv.IsValid() {
					v.Elem().FieldByName(name).Set(fv)
				}
			case reflect.Interface:
				// To recurse is divine!-)
				if reflect.TypeOf(value).String() == "map[string]interface {}" {
					if typeName_, ok := value.(map[string]interface{})["TYPENAME"]; ok {
						if _, ok := typeRegistry[typeName_.(string)]; ok {
							fv, _ := InitializeStructure(value.(map[string]interface{}), setEntity)
							if fv.IsValid() {
								// I will set the reference in the parent object.
								v.Elem().FieldByName(name).Set(reflect.ValueOf(value.(map[string]interface{})["UUID"]))
							}
						} else {
							// Here it's a dynamic entity...
							v.Elem().FieldByName(name).Set(reflect.ValueOf(value))
						}
					} else {
						v.Elem().FieldByName(name).Set(reflect.ValueOf(value))
					}
				}
			case reflect.Map:
				fv, _ := InitializeStructure(value.(map[string]interface{}), setEntity)
				if fv.IsValid() {
					v.Elem().FieldByName(name).Set(fv.Elem())
				}

			default:
				if reflect.TypeOf(value).String() == "map[string]interface {}" {
					if typeName_, ok := value.(map[string]interface{})["TYPENAME"]; ok {
						if _, ok := typeRegistry[typeName_.(string)]; ok {
							fv, _ := InitializeStructure(value.(map[string]interface{}), setEntity)
							if fv.IsValid() {
								// I will set the reference in the parent object.
								v.Elem().FieldByName(name).Set(reflect.ValueOf(value.(map[string]interface{})["UUID"]))
							}
						} else {
							// Here it's a dynamic entity...
							v.Elem().FieldByName(name).Set(reflect.ValueOf(value))
						}
					} else {
						fv, _ := InitializeStructure(value.(map[string]interface{}), setEntity)

						if fv.IsValid() {
							v.Elem().FieldByName(name).Set(fv.Elem())
						}
					}

				} else {
					// Convert is use to enumeration type who are int and must be convert to
					// it const type representation.
					fv := initializeBaseTypeValue(ft.Type, value).Convert(ft.Type)
					if fv.IsValid() {
						v.Elem().FieldByName(name).Set(fv)
					}
				}
			}
		}
	}

	// Return the initialysed value...
	return v
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
