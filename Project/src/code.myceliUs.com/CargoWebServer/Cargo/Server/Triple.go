package Server

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"

	"code.myceliUs.com/Utility"
	"github.com/robertkrimen/otto"
)

type Triple struct {
	Subject   string
	Predicate string
	Object    interface{}

	// True if the object must be indexed...
	IsIndex bool
}

/**
 * Create a map[string] interface{} from array of triples.
 */
func FromTriples(values [][]interface{}) map[string]interface{} {

	obj := make(map[string]interface{}, 0)
	var typeName string
	if len(values) > 0 {
		typeName = strings.Split(values[0][0].(string), "%")[0]
	} else {
		return nil
	}

	storeId := strings.Split(typeName, ".")[0]
	prototype, _ := entityManager.getEntityPrototype(typeName, storeId)
	obj["TYPENAME"] = typeName

	for i := 0; i < len(values); i++ {
		if values[i][1] != "TYPENAME" {
			propertie := strings.Split(values[i][1].(string), ":")[2]
			value := values[i][2]
			filedIndex := prototype.getFieldIndex(propertie)
			if filedIndex != -1 {
				fieldType := prototype.FieldsType[filedIndex]
				if strings.HasPrefix(fieldType, "[]") {
					if strings.HasSuffix(fieldType, ":Ref") {
						// In case of a reference I will create an array.
						if obj[propertie] == nil {
							obj[propertie] = make([]string, 0)
						}
						if Utility.IsValidEntityReferenceName(value.(string)) {
							obj[propertie] = append(obj[propertie].([]string), value.(string))
						}
					} else {
						if Utility.IsValidEntityReferenceName(value.(string)) {
							// So here I will get the value from the store.
							if obj[propertie] == nil {
								obj[propertie] = make([]string, 0)
							}
							obj[propertie] = append(obj[propertie].([]string), value.(string))
						} else {
							values_ := make([]interface{}, 0)
							err := json.Unmarshal([]byte(value.(string)), &values_)
							if err == nil {
								if values_ != nil {
									obj[propertie] = values_
								}
							}
						}
					}
				} else {
					// Set the value...
					if strings.HasSuffix(fieldType, ":Ref") {
						obj[propertie] = value
					} else {
						if strings.HasPrefix(fieldType, "xs.") {
							obj[propertie] = value
						} else {
							if reflect.TypeOf(value).Kind() == reflect.String {
								if Utility.IsValidEntityReferenceName(value.(string)) {
									obj[propertie] = values[i][2].(string)
								} else {
									obj[propertie] = value
								}
							} else {
								obj[propertie] = value
							}
						}
					}
				}
			}
		}
	}

	return obj
}

/**
 * Recursively generate Triple from structure values.
 */
func ToTriples(values map[string]interface{}, triples *[]interface{}) error {
	var uuid string
	if values["UUID"] != nil {
		uuid = values["UUID"].(string)
		typeName := values["TYPENAME"].(string)

		// append the type name as a relation.
		*triples = append(*triples, Triple{uuid, "TYPENAME", typeName, true})
		prototype, _ := entityManager.getEntityPrototype(typeName, strings.Split(typeName, ".")[0])
		for k, v := range values {
			fieldIndex := prototype.getFieldIndex(k)
			if fieldIndex != -1 {
				fieldType := prototype.FieldsType[fieldIndex]
				fieldType_ := strings.Replace(fieldType, "[]", "", -1)
				fieldType_ = strings.Replace(fieldType_, ":Ref", "", -1)

				// Satic entity enum type.
				if strings.HasPrefix(fieldType_, "enum:") {
					//ex: enum:AccessType_Hidden:AccessType_Public:AccessType_Restricted Here the type will be AccessType
					fieldType_ = strings.Split(fieldType_, ":")[1]
					fieldType_ = strings.Split(fieldType_, "_")[0]
				}

				if v != nil {
					isIndex := Utility.Contains(prototype.Ids, k)
					if !isIndex {
						isIndex = Utility.Contains(prototype.Indexs, k)
					}

					// In case of otto value...
					if reflect.TypeOf(v).String() == "otto.Value" {
						v, _ = v.(otto.Value).Export()
					}

					if strings.HasSuffix(fieldType, ":Ref") {
						if strings.HasPrefix(fieldType, "[]") {

							if reflect.TypeOf(v).String() == "[]string" {
								for i := 0; i < len(v.([]string)); i++ {
									*triples = append(*triples, Triple{uuid, typeName + ":" + fieldType_ + ":" + k, v.([]string)[i], isIndex})
								}
							} else if reflect.TypeOf(v).String() == "[]interface {}" {
								for i := 0; i < len(v.([]interface{})); i++ {
									*triples = append(*triples, Triple{uuid, typeName + ":" + fieldType_ + ":" + k, v.([]interface{})[i].(string), isIndex})
								}
							} else if reflect.TypeOf(v).String() == "[]map[string]interface {}" {
								for i := 0; i < len(v.([]map[string]interface{})); i++ {
									if v.([]map[string]interface{})[i]["UUID"] != nil {
										*triples = append(*triples, Triple{uuid, typeName + ":" + fieldType_ + ":" + k, v.([]map[string]interface{})[i]["UUID"].(string), isIndex})
									}
								}
							} else {
								log.Panicln("------> reference type fail!", reflect.TypeOf(v).String())
							}
						} else {
							if reflect.TypeOf(v).String() == "map[string]interface {}" {
								if v.(map[string]interface{})["TYPENAME"] == nil {
									ToTriples(v.(map[string]interface{}), triples)
								}
							} else if reflect.TypeOf(v).Kind() == reflect.String {
								// Here I will append attribute...
								*triples = append(*triples, Triple{uuid, typeName + ":" + fieldType_ + ":" + k, v, isIndex})
							} else {
								log.Panicln("------> reference type fail!", reflect.TypeOf(v).String())
							}
						}

					} else {
						if strings.HasPrefix(fieldType, "[]") {
							if reflect.TypeOf(v).String() == "[]interface {}" {
								if len(v.([]interface{})) > 0 {
									if reflect.TypeOf(v.([]interface{})[0]).String() == "map[string]interface {}" {
										if v.([]interface{})[0].(map[string]interface{})["TYPENAME"] == nil {
											for i := 0; i < len(v.([]interface{})); i++ {
												ToTriples(v.([]interface{})[i].(map[string]interface{}), triples)
											}
										}
									} else {
										// a regular array here.
										if reflect.TypeOf(v.([]interface{})[0]).Kind() == reflect.String {
											if Utility.IsValidEntityReferenceName(v.([]interface{})[0].(string)) {
												for i := 0; i < len(v.([]interface{})); i++ {
													uuid_ := v.([]interface{})[i].(string)
													*triples = append(*triples, Triple{uuid, typeName + ":" + strings.Split(uuid_, "%")[0] + ":" + k, uuid_, isIndex})
												}
											} else {
												str, err := json.Marshal(v)
												if err == nil {
													*triples = append(*triples, Triple{uuid, typeName + ":" + fieldType_ + ":" + k, string(str), isIndex})
												}
											}
										} else {
											str, err := json.Marshal(v)
											if err == nil {
												*triples = append(*triples, Triple{uuid, typeName + ":" + fieldType_ + ":" + k, string(str), isIndex})
											}
										}
									}
								}
							} else if reflect.TypeOf(v).String() == "[]string" {
								if len(v.([]string)) > 0 {
									if Utility.IsValidEntityReferenceName(v.([]string)[0]) {
										for i := 0; i < len(v.([]string)); i++ {
											uuid_ := v.([]string)[i]
											*triples = append(*triples, Triple{uuid, typeName + ":" + strings.Split(uuid_, "%")[0] + ":" + k, uuid_, isIndex})
										}
									} else {
										str, err := json.Marshal(v)
										if err == nil {
											*triples = append(*triples, Triple{uuid, typeName + ":" + fieldType_ + ":" + k, string(str), isIndex})
										}
									}
								}
							} else {
								str, err := json.Marshal(v)
								if err == nil {
									*triples = append(*triples, Triple{uuid, typeName + ":" + fieldType_ + ":" + k, string(str), isIndex})
								}
							}
						} else {
							if reflect.TypeOf(v).String() == "map[string]interface {}" {
								if v.(map[string]interface{})["TYPENAME"] == nil {
									ToTriples(v.(map[string]interface{}), triples)
								}
							} else {
								// Dont save the file disk data into the entity...
								if typeName == "CargoEntities.File" {
									if values["M_fileType"].(float64) == 2 {
										values["M_data"] = ""
									}
								}
								// Here I will append attribute...
								if reflect.TypeOf(v).Kind() == reflect.String {
									if Utility.IsValidEntityReferenceName(v.(string)) && k != "ParentUuid" && k != "UUID" {
										uuid_ := v.(string)
										*triples = append(*triples, Triple{uuid, typeName + ":" + strings.Split(uuid_, "%")[0] + ":" + k, uuid_, isIndex})
									} else {
										*triples = append(*triples, Triple{uuid, typeName + ":" + fieldType_ + ":" + k, v, isIndex})
									}
								} else {
									*triples = append(*triples, Triple{uuid, typeName + ":" + fieldType_ + ":" + k, v, isIndex})
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}
