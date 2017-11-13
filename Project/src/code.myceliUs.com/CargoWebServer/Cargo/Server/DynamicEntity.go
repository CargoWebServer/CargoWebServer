package Server

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	"code.myceliUs.com/Utility"
	"github.com/pborman/uuid"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
)

/**
 * Implementation of the Entity. Dynamic entity use a map[string]interface {}
 * to store object informations. All information in the Entity itself other
 * than object are store.
 */
type DynamicEntity struct {

	// Hard link parent->[]childs
	childsUuid []string

	// Soft link []referenced->[]references
	referencesUuid []string
	referenced     []EntityRef

	/** object will be a map... **/
	object map[string]interface{}

	/**
	 * Use to protected the ressource access...
	 */
	sync.RWMutex
}

func getEntityPrototype(values map[string]interface{}) (*EntityPrototype, error) {
	if !Utility.IsValidPackageName(values["TYPENAME"].(string)) {
		return nil, errors.New("Package name '" + values["TYPENAME"].(string) + "' is not valid.")
	}

	typeName := values["TYPENAME"].(string)

	packageName := typeName[0:strings.Index(typeName, ".")]

	prototype, err := GetServer().GetEntityManager().getEntityPrototype(typeName, packageName)

	if err != nil {
		// Try to find the type in sql_info...
		prototype, err = GetServer().GetEntityManager().getEntityPrototype(typeName, "sql_info") //sql_info contain information from the sql schema
		if err != nil {
			return nil, err
		}
	}
	return prototype, nil
}

/**
 * Create a new dynamic entity...
 */
func (this *EntityManager) newDynamicEntity(parentUuid string, values map[string]interface{}) (*DynamicEntity, *CargoEntities.Error) {

	// I will test restriction before I save the entity.
	restrictionError := applyEntityRestrictions(values)
	if restrictionError != nil {
		errObj := NewError(Utility.FileLine(), PROTOTYPE_RESTRICTIONS_ERROR, SERVER_ERROR_CODE, restrictionError)
		return nil, errObj
	}

	// Set the parent uuid in that case.
	if len(parentUuid) == 0 && values["ParentUuid"] != nil {
		parentUuid = values["ParentUuid"].(string)
	} else {
		values["ParentUuid"] = parentUuid
	}

	var entity *DynamicEntity
	prototype, err := getEntityPrototype(values)
	if err != nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), PROTOTYPE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Prototype not found for type '"+values["TYPENAME"].(string)+"'."))
		return nil, cargoError
	}

	if values["UUID"] == nil {
		values["UUID"] = ""
	}

	// Keep track of the perent uuid inside the object.
	if len(values["UUID"].(string)) == 0 {
		// I will alway prefer a determistic key over a ramdom value...
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			values["UUID"] = values["TYPENAME"].(string) + "%" + uuid.NewRandom().String()
		} else {
			// The key will be compose by the parent uuid.
			var keyInfo string
			if len(parentUuid) > 0 {
				values["ParentUuid"] = parentUuid
				keyInfo += parentUuid + ":"
			}

			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				if values[prototype.Ids[i]] != nil {
					keyInfo += Utility.ToString(values[prototype.Ids[i]])
				} else {
					keyInfo += "undefined" // undefined is the default value here.
				}
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}
			// The uuid is in that case a MD5 value.
			values["UUID"] = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}

	// Try to retreive the entity.
	if val, ok := this.contain(values["UUID"].(string)); ok {
		if val != nil {
			entity = val.(*DynamicEntity)

			// Calculate the checksum.
			sum0 := Utility.GetChecksum(values)
			sum1 := entity.GetChecksum()
			if sum0 != sum1 {
				entity.SetObjectValues(values)
				entity.SetNeedSave(true)
			} else {
				// If the value dosent exist It need to be save
				// and is not already need save...
				if !entity.NeedSave() {
					if entity.Exist() == false {
						entity.SetNeedSave(true)
					} else {
						entity.SetNeedSave(false)
					}
				}
			}

			return entity, nil
		}
	}

	// Here if the enity is nil I will create a new instance.
	if entity == nil {
		// Create a new dynamic entity it that case.
		entity = new(DynamicEntity)

		// If the object contain an id...
		entity.setValue("UUID", values["UUID"].(string))

		entity.childsUuid = make([]string, 0)
		entity.referencesUuid = make([]string, 0)
		entity.referenced = make([]EntityRef, 0)

	}

	// Dynamic entity are not lazy by default.
	if values["lazy"] == nil {
		values["lazy"] = false
	}

	// Set the object value with the values, need save will be set
	// inside se object value.
	entity.SetObjectValues(values)

	// Set the init value at false.
	entity.SetInit(false)

	// insert it into the cache.
	this.insert(entity)

	return entity, nil
}

/**
 * Thread safe function
 */
func (this *DynamicEntity) setValue(field string, value interface{}) error {
	this.Lock()
	defer this.Unlock()
	if this.object == nil {
		this.object = make(map[string]interface{}, 0)
	}

	// Here the value is in the map.
	this.object[field] = value
	return nil
}

/**
 * Thread safe function
 */
func (this *DynamicEntity) getValue(field string) interface{} {
	this.Lock()
	defer this.Unlock()
	return this.object[field]
}

/**
 * Thread safe function, apply only on none array field...
 */
func (this *DynamicEntity) deleteValue(field string) {
	this.Lock()
	defer this.Unlock()
	prototype := this.GetPrototype()
	fieldType := prototype.FieldsType[prototype.getFieldIndex(field)]
	if !strings.HasPrefix(fieldType, "[]") {
		if strings.HasSuffix(fieldType, ":Ref") {
			refUuid := this.object[field].(string)
			ref, err := GetServer().GetEntityManager().getEntityByUuid(refUuid, false)
			if err == nil {
				this.RemoveReference(field, ref)
			}
		}
		// Remove the field itself.
		delete(this.object, field)
	}
}

/**
 * Set object.
 */
func (this *DynamicEntity) setObject(obj map[string]interface{}) {
	this.Lock()
	defer this.Unlock()
	this.object = obj
}

/**
 * Append a new value.
 */
func (this *DynamicEntity) appendValue(field string, value interface{}) {
	values := this.getValue(field)
	prototype := this.GetPrototype()
	if values == nil {
		// Here no value aready exist.
		if reflect.TypeOf(value).Kind() == reflect.String {
			values = make([]string, 1)
			values.([]string)[0] = value.(string)
			this.setValue(field, values)
		} else if reflect.TypeOf(value).String() == "*Server.DynamicEntity" {
			if strings.HasSuffix(prototype.FieldsType[prototype.getFieldIndex(field)], ":Ref") {
				values = make([]string, 1)
				values.([]string)[0] = value.(Entity).GetUuid()
				this.setValue(field, values)
			} else {
				this.AppendChild(field, value.(Entity))
			}
		} else {
			// Other types.
			values = make([]interface{}, 1)
			values.([]interface{})[0] = value
			this.setValue(field, values)
		}

	} else {
		// An array already exist in that case.
		if reflect.TypeOf(value).Kind() == reflect.String {
			values = append(values.([]string), value.(string))
			this.setValue(field, values)
		} else if reflect.TypeOf(value).String() == "*Server.DynamicEntity" {
			if strings.HasSuffix(prototype.FieldsType[prototype.getFieldIndex(field)], ":Ref") {
				var isExist = false
				for i := 0; i < len(values.([]*DynamicEntity)); i++ {
					if value.(Entity).GetUuid() == values.(*DynamicEntity).GetUuid() {
						isExist = true
					}
				}
				if !isExist {
					values = append(values.([]*DynamicEntity), value.(*DynamicEntity))
				}
			} else {
				this.AppendChild(field, value.(Entity))
			}
		} else {
			// Other types.
			values = append(values.([]interface{}), value)
			this.setValue(field, values)
		}
	}

}

/**
 * Determine is an entity is initialyse or not.
 */
func (this *DynamicEntity) IsInit() bool {
	return this.getValue("IsInit").(bool)
}

/**
 * Set if an entity must be inityalyse.
 */
func (this *DynamicEntity) SetInit(isInit bool) {
	this.setValue("IsInit", isInit)
}

/**
 * Test if an entity need to be save.
 */
func (this *DynamicEntity) NeedSave() bool {
	return this.getValue("NeedSave").(bool)
}

/**
 * Set if an entity need to be save.
 */
func (this *DynamicEntity) SetNeedSave(needSave bool) {
	this.setValue("NeedSave", needSave)
}

/**
 * Set if an entity need to be save.
 */
func (this *DynamicEntity) IsLazy() bool {
	return this.getValue("lazy").(bool)
}

/**
 * Initialyse a entity with a given id.
 */
func (this *DynamicEntity) InitEntity(id string, lazy bool) error {
	this.setValue("lazy", lazy)
	err := this.initEntity(id, "", lazy)

	//log.Println("After init:", toJsonStr(this.object))
	return err
}

func (this *DynamicEntity) initEntity(id string, path string, lazy bool) error {
	// cut infinite recursion here.
	if strings.Index(path, id) != -1 {
		return nil
	}

	// If the value is already in the cache I have nothing todo...
	if this.IsInit() == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*DynamicEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.SetInit(false)
	}

	// I will set the id, (must be a uuid...)
	this.setValue("UUID", id)

	typeName := id[0:strings.Index(id, "%")]
	packageName := typeName[0:strings.Index(typeName, ".")]

	var query EntityQuery
	query.TypeName = typeName
	prototype := this.GetPrototype()
	// Here I will append the rest of the fields...
	// append the list of fields...
	query.Fields = append(query.Fields, prototype.Fields...)

	// The index of search...
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	var fieldsType []interface{} // not use...
	var params []interface{}

	storeId := packageName
	if reflect.TypeOf(dataManager.getDataStore(storeId)).String() == "*Server.SqlDataStore" {
		storeId = "sql_info" // Must save or update value from sql info instead.
	}

	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(storeId, string(queryStr), fieldsType, params)
	if err != nil {
		log.Println(Utility.FileLine(), "--> No object found for entity ", id, " ", err)
		return err
	}

	// Initialisation of information of Interface...
	if len(results) > 0 {

		if len(results[0]) == 0 {
			return errors.New("No value found for entity " + this.GetUuid())
		}
		// Set the common values...
		this.setValue("TYPENAME", typeName) // Set the typeName

		this.setValue("UUID", results[0][0].(string))

		// Set the parent uuid...
		this.setValue("ParentUuid", results[0][1].(string)) // Set the parent uuid.

		this.setValue("ParentLnk", results[0][2].(string)) // Set the parent relation.

		//init the child...
		childsUuidStr := results[0][prototype.getFieldIndex("childsUuid")].(string)

		if len(childsUuidStr) > 0 {
			err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
			if err != nil {
				log.Println(Utility.FileLine(), " unable to read child uuid! for entity ", this.GetUuid(), " ", childsUuidStr)
				return err
			}
		}

		referencedStr := results[0][prototype.getFieldIndex("referenced")].(string)
		this.referenced = make([]EntityRef, 0)
		if len(referencedStr) > 0 {
			err = json.Unmarshal([]byte(referencedStr), &this.referenced)
			if err != nil {
				log.Println("---> fail to get ref ")
				return err
			}
		}

		// Now The value from the prototype...
		// The first tow field are uuid and parentUuid
		// and the last tow fields are childsuuid and referenced...
		for i := 3; i < len(results[0])-2; i++ {
			fieldName := prototype.Fields[i]
			fieldType := prototype.FieldsType[i]
			isNull := false

			if reflect.TypeOf(results[0][i]).String() == "string" {
				isNull = results[0][i].(string) == "null"
			}

			if !isNull {
				// Determine if the object is a reference.
				// Array's...
				if strings.HasPrefix(fieldType, "[]") {
					if strings.HasPrefix(fieldType, "[]xs.") || strings.HasPrefix(fieldType, "[]sqltypes.") || fieldName == "M_listOf" {
						values := make([]interface{}, 0)
						err := json.Unmarshal([]byte(results[0][i].(string)), &values)
						if err != nil {
							log.Println("fail to get value --------->", results[0][i].(string))
						} else {
							this.setValue(fieldName, values)
						}
					} else {
						// If the field is a reference.
						if this.isRef(fieldType, fieldName) {
							// Here I got an array for reference string.
							var referencesId []string

							json.Unmarshal([]byte(results[0][i].(string)), &referencesId)

							refTypeName := fieldType

							this.setValue(fieldName, referencesId)
							for i := 0; i < len(referencesId); i++ {
								id_ := refTypeName + "$$" + referencesId[i]
								GetServer().GetEntityManager().appendReference(fieldName, this.getValue("UUID").(string), id_)
							}

						} else {
							// Here I have objects.
							// A list of uuid's
							uuids := make([]string, 0)
							var err error

							if reflect.TypeOf(results[0][i]).String() == "string" {
								// Uuids
								err = json.Unmarshal([]byte(results[0][i].(string)), &uuids)
							}

							if err != nil {
								log.Println(Utility.FileLine(), "unable to read array of values at line 250!")
								return err
							}

							// Only uuid's are accepted value's here.
							if len(uuids) > 0 {
								log.Println("---------> uuids found! ", uuids)
								// Set an empty array here...
								if Utility.IsValidEntityReferenceName(uuids[0]) && !lazy {
									this.setValue(fieldName, make([]map[string]interface{}, 0))
									for i := 0; i < len(uuids); i++ {
										if len(uuids[i]) > 0 {
											if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
												// The entity already exist.
												dynamicEntity := instance.(*DynamicEntity)
												this.AppendChild(fieldName, dynamicEntity)
											} else {
												// Here I need to initialise it!
												typeName := strings.Replace(strings.Replace(fieldType, ":Ref", "", -1), "[]", "", -1)
												values := make(map[string]interface{}, 0)
												values["UUID"] = uuids[i]
												values["TYPENAME"] = uuids[i][0:strings.Index(uuids[i], "%")]

												// I will try to create a static entity...
												newEntityMethod := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
												params := make([]interface{}, 3)
												params[0] = ""
												params[1] = uuids[i]
												params[2] = values
												staticEntity, err := Utility.CallMethod(GetServer().GetEntityManager(), newEntityMethod, params)

												if err != nil {
													// In that case I will try with dynamic entity.
													dynamicEntity, errObj := GetServer().GetEntityManager().newDynamicEntity(this.GetUuid(), values)
													if errObj == nil {
														// initialise the sub entity.
														err := dynamicEntity.initEntity(uuids[i], path+"|"+this.GetUuid(), lazy)
														if err != nil {
															// I will try to remove it from it child if is there...
															this.RemoveChild(fieldName, uuids[i])
															this.SetNeedSave(true)
														} else {
															dynamicEntity.AppendReferenced(fieldName, this)
															this.AppendChild(fieldName, dynamicEntity)
														}
													}
												} else {
													staticEntity.(Entity).InitEntity(uuids[i], lazy)
													staticEntity.(Entity).AppendReferenced(fieldName, this)
													this.AppendChild(fieldName, staticEntity.(Entity))
												}

											}
										}
									}
								} else if lazy {
									// Lazy loaded object contain only the uuid of the object...
									var uuids []string
									json.Unmarshal([]byte(results[0][i].(string)), &uuids)
									this.setValue(fieldName, uuids)
								}
							}
						}
					}
				} else {
					if strings.HasPrefix(fieldType, "xs.") || strings.HasPrefix(fieldType, "sqltypes.") || fieldName == "M_valueOf" {
						this.setValue(fieldName, results[0][i])
					} else {
						// Not an array here.
						if this.isRef(fieldType, fieldName) {
							// In that case the reference must be a simple string.
							refTypeName := fieldType
							id_ := refTypeName + "$$" + results[0][i].(string)
							this.setValue(fieldName, results[0][i].(string))
							GetServer().GetEntityManager().appendReference(fieldName, this.getValue("UUID").(string), id_)
						} else {
							// Here I have an object...
							if reflect.TypeOf(results[0][i]).String() == "string" {
								uuid := results[0][i].(string)
								if Utility.IsValidEntityReferenceName(uuid) && !lazy {
									if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
										dynamicEntity := instance.(*DynamicEntity)
										dynamicEntity.AppendReferenced(fieldName, this)
										this.AppendChild(fieldName, dynamicEntity)
									} else {
										typeName := strings.Replace(fieldType, ":Ref", "", -1)
										values := make(map[string]interface{}, 0)
										values["UUID"] = uuid
										values["TYPENAME"] = uuid[0:strings.Index(uuid, "%")]

										// I will try to create a static entity...
										newEntityMethod := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
										params := make([]interface{}, 3)
										params[0] = ""
										params[1] = uuid
										params[2] = values
										staticEntity, err := Utility.CallMethod(GetServer().GetEntityManager(), newEntityMethod, params)

										if err != nil {
											// In that case I will try with dynamic entity.
											dynamicEntity, errObj := GetServer().GetEntityManager().newDynamicEntity(this.GetUuid(), values)
											if errObj == nil {
												// initialise the sub entity.
												err := dynamicEntity.initEntity(uuid, path+"|"+this.GetUuid(), lazy)
												if err != nil {
													// I will try to remove it from it child if is there...
													this.RemoveChild(fieldName, uuid)
													this.SetNeedSave(true)
												} else {
													dynamicEntity.AppendReferenced(fieldName, this)
													this.AppendChild(fieldName, dynamicEntity)
												}
											} else {
												log.Println("-------> 570 error ", errObj.GetBody())
											}
										} else {
											staticEntity.(Entity).InitEntity(uuid, lazy)
											staticEntity.(Entity).AppendReferenced(fieldName, this)
											this.AppendChild(fieldName, staticEntity.(Entity))
										}
									}
								} else {
									//log.Println("-------> plain string: ", fieldName, results[0][i])
									// A plain string...
									this.setValue(fieldName, results[0][i])
								}
							} else {
								// A base type other than strings... bool, float64, int64 etc...
								this.setValue(fieldName, results[0][i])
							}
						}
					}
				}
			} else {
				this.setValue(fieldName, nil)
			}
		}
	}

	// set init done.
	this.SetInit(true)
	this.SetNeedSave(false)
	GetServer().GetEntityManager().InitEntity(this, lazy)

	// if some change are found at initialysation I will update the values.
	if this.NeedSave() == true {
		// The entity need to be save...
		this.saveEntity("")
	}

	if storeId == "sql_info" {
		// Now I will initialyse references
		dataManager.setEntityReferences(this.GetUuid(), false)
	}
	return nil
}

/**
 * Save entity
 */
func (this *DynamicEntity) SaveEntity() {
	this.saveEntity("")
	//log.Println("After save:", toJsonStr(this.object))
}

func (this *DynamicEntity) saveEntity(path string) {
	// Do not save if is nt necessary...
	if !this.NeedSave() || strings.Index(path, this.GetUuid()) != -1 {
		log.Println("Entity ", this.GetUuid(), " Not need to be save.", this.NeedSave())
		return
	}

	// cut the recusion here.
	this.SetNeedSave(false)
	this.SetInit(true)

	// I will save the information into the database...
	var DynamicEntityInfo []interface{}
	var query EntityQuery
	query.TypeName = this.GetTypeName()

	// General information.
	query.Fields = append(query.Fields, "UUID")
	DynamicEntityInfo = append(DynamicEntityInfo, this.GetUuid())

	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	if len(this.GetParentUuid()) > 0 {
		DynamicEntityInfo = append(DynamicEntityInfo, this.GetParentUuid())
		DynamicEntityInfo = append(DynamicEntityInfo, this.GetParentLnk())
	} else {
		DynamicEntityInfo = append(DynamicEntityInfo, "")
		DynamicEntityInfo = append(DynamicEntityInfo, "")
	}

	prototype := this.GetPrototype()
	// Now the fields of the object, from the prototype...
	// The first tow field are uuid and parentUuid
	// and the last tow fields are childsuuid and referenced...
	for i := 3; i < len(prototype.Fields)-2; i++ {
		fieldName := prototype.Fields[i]
		fieldType := prototype.FieldsType[i]

		// append the field in the query fields list.
		query.Fields = append(query.Fields, fieldName)

		// Print the field name.
		if this.getValue(fieldName) != nil {
			// Array's
			if strings.HasPrefix(fieldType, "[]") {
				if strings.HasPrefix(fieldType, "[]xs.") || strings.HasPrefix(fieldType, "[]sqltypes.") || fieldName == "M_listOf" {
					valuesStr, err := json.Marshal(this.getValue(fieldName))
					if err != nil {
						DynamicEntityInfo = append(DynamicEntityInfo, "null")
					} else {
						DynamicEntityInfo = append(DynamicEntityInfo, string(valuesStr))
					}
				} else {
					// The reference is an array...
					// in the case of reference.
					if this.isRef(fieldType, fieldName) {
						// Keep the reference uuid here.
						refUuids := make([]string, 0)
						refIds := make([]string, 0)
						// Value can be of type []interface {} or [] of string.
						if reflect.TypeOf(this.getValue(fieldName)).String() == "[]interface {}" {
							// An array of interface type.
							for j := 0; j < len(this.getValue(fieldName).([]interface{})); j++ {
								// Append reference.
								obj := this.getValue(fieldName).([]interface{})[j]
								if obj != nil {
									var refUuid string
									// It can be an object or a string...
									if reflect.TypeOf(obj).String() == "map[string]interface {}" {
										refUuid = obj.(map[string]interface{})["UUID"].(string)
									} else if reflect.TypeOf(obj).String() == "string" {
										refUuid = obj.(string)
										if strings.HasPrefix(refUuid, "#") {
											refUuid = refUuid[1:]
										}
									}
									if len(refUuid) > 0 {
										if Utility.IsValidEntityReferenceName(refUuid) {
											if !Utility.Contains(refUuids, refUuid) {
												refUuids = append(refUuids, refUuid)
											}
										} else {
											// Append to the ids...
											if !Utility.Contains(refIds, refUuid) {
												refIds = append(refIds, refUuid)
											}
										}
									}
								}
							}

						} else if reflect.TypeOf(this.getValue(fieldName)).String() == "[]string" {
							for j := 0; j < len(this.getValue(fieldName).([]string)); j++ {
								refUuid := this.getValue(fieldName).([]string)[j]
								if len(refUuid) > 0 {
									if strings.HasPrefix(refUuid, "#") {
										refUuid = refUuid[1:]
									}

									if !Utility.Contains(refUuids, refUuid) {
										refUuids = append(refUuids, refUuid)
									}
								}
							}
						} else {
							log.Println(Utility.FileLine(), "---> field no found!", fieldName, ":", fieldType, ":", reflect.TypeOf(this.getValue(fieldName)).String())
						}

						// Set the pending uuid references.
						for j := 0; j < len(refUuids); j++ {
							refUuid := refUuids[j]
							refTypeName := strings.Replace(fieldType, "[]", "", -1)
							refTypeName = strings.Replace(refTypeName, ":Ref", "", -1)
							id := refTypeName + "$$" + refUuid

							// Append the referenced...
							GetServer().GetEntityManager().appendReferenced(fieldName, this.GetUuid(), id)
						}

						for j := 0; j < len(refIds); j++ {
							refId := refIds[j]
							refTypeName := strings.Replace(fieldType, "[]", "", -1)
							refTypeName = strings.Replace(refTypeName, ":Ref", "", -1)
							id := refTypeName + "$$" + refId
							// Append the referenced...
							GetServer().GetEntityManager().appendReferenced(fieldName, this.GetUuid(), id)
						}

						refUuidsStr, _ := json.Marshal(refUuids)
						DynamicEntityInfo = append(DynamicEntityInfo, string(refUuidsStr))

					} else {
						// In the case of object.
						if reflect.TypeOf(this.getValue(fieldName)).String() == "[]map[string]interface {}" {
							subEntityIds := make([]string, 0)
							for j := 0; j < len(this.getValue(fieldName).([]map[string]interface{})); j++ {
								// I will get the values of the sub item...
								uuid := this.getValue(fieldName).([]map[string]interface{})[j]["UUID"].(string)
								subValues := this.getValue(fieldName).([]map[string]interface{})[j]
								typeName := strings.Replace(fieldType, "[]", "", -1)

								// I will try to create a static entity...
								newEntityMethod := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
								params := make([]interface{}, 3)
								params[0] = ""
								params[1] = uuid
								params[2] = subValues
								staticEntity, err := Utility.CallMethod(GetServer().GetEntityManager(), newEntityMethod, params)

								if err != nil {
									// In that case I will try with dynamic entity.
									// I will create the sub value...
									subEntity, errObj := GetServer().GetEntityManager().newDynamicEntity(this.GetUuid(), subValues)
									if errObj == nil {
										subEntity.AppendReferenced(fieldName, this)
										this.AppendChild(fieldName, subEntity)
										subEntityIds = append(subEntityIds, subEntity.GetUuid())
										if subEntity.NeedSave() {
											subEntity.saveEntity(path + "|" + this.GetUuid())
										}
									}
								} else {
									staticEntity.(Entity).AppendReferenced(fieldName, this)
									this.AppendChild(fieldName, staticEntity.(Entity))
									subEntityIds = append(subEntityIds, uuid)
									if staticEntity.(Entity).NeedSave() {
										staticEntity.(Entity).SaveEntity()
									}
								}

							}
							subEntityIdsStr, _ := json.Marshal(subEntityIds)
							DynamicEntityInfo = append(DynamicEntityInfo, string(subEntityIdsStr))
						} else if reflect.TypeOf(this.getValue(fieldName)).String() == "[]interface {}" {
							subEntityIds := make([]string, 0)
							for j := 0; j < len(this.getValue(fieldName).([]interface{})); j++ {
								if reflect.TypeOf(this.getValue(fieldName).([]interface{})[j]).String() == "map[string]interface {}" {
									// I will get the values of the sub item...
									subValues := this.getValue(fieldName).([]interface{})[j].(map[string]interface{})
									// I will create the sub value...
									typeName := strings.Replace(strings.Replace(fieldType, ":Ref", "", -1), "[]", "", -1)

									var uuid string
									if subValues["UUID"] != nil {
										uuid = subValues["UUID"].(string)
									}

									// I will try to create a static entity...
									newEntityMethod := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
									params := make([]interface{}, 3)
									params[0] = ""
									params[1] = uuid
									params[2] = subValues
									staticEntity, err := Utility.CallMethod(GetServer().GetEntityManager(), newEntityMethod, params)

									if err != nil {
										// In that case I will try with dynamic entity.
										// I will create the sub value...
										subEntity, errObj := GetServer().GetEntityManager().newDynamicEntity(this.GetUuid(), subValues)
										if errObj == nil {
											subEntity.AppendReferenced(fieldName, this)
											this.AppendChild(fieldName, subEntity)
											subEntityIds = append(subEntityIds, subEntity.GetUuid())
											if subEntity.NeedSave() {
												subEntity.saveEntity(path + "|" + this.GetUuid())
											}
										}
									} else {
										staticEntity.(Entity).AppendReferenced(fieldName, this)
										this.AppendChild(fieldName, staticEntity.(Entity))
										subEntityIds = append(subEntityIds, uuid)
										if staticEntity.(Entity).NeedSave() {
											staticEntity.(Entity).SaveEntity()
										}
									}
								} else {
									log.Println(Utility.FileLine(), "---> field no found!", fieldName, ":", fieldType, ":", reflect.TypeOf(this.getValue(fieldName).([]interface{})[j]).String())
								}
							}
							subEntityIdsStr, _ := json.Marshal(subEntityIds)
							DynamicEntityInfo = append(DynamicEntityInfo, string(subEntityIdsStr))
						} else if reflect.TypeOf(this.getValue(fieldName)).String() == "[]string" {
							// Here it's a lazy intialysed object.
							subEntityIdsStr, _ := json.Marshal(this.getValue(fieldName).([]string))
							DynamicEntityInfo = append(DynamicEntityInfo, string(subEntityIdsStr))
						}
					}
				}
			} else {
				// Not an array...
				if strings.HasPrefix(fieldType, "xs.") || strings.HasPrefix(fieldType, "sqltypes.") || fieldName == "M_valueOf" {
					// Little fix to convert float into int type as needed.
					if fieldType == "xs.byte" || fieldType == "xs.short" || fieldType == "xs.int" || fieldType == "xs.integer" || fieldType == "xs.long" || fieldType == "xs.unsignedInt" || fieldType == "xs.unsignedByte" || fieldType == "xs.unsignedShort" || fieldType == "xs.unsignedLong" {
						if reflect.TypeOf(this.getValue(fieldName)).Kind() == reflect.Float32 {
							if fieldType == "xs.int" || fieldType == "xs.integer" {
								val := int32(this.getValue(fieldName).(float32))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.byte" {
								val := int8(this.getValue(fieldName).(float32))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.short" {
								val := int16(this.getValue(fieldName).(float32))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.unsignedByte" {
								val := uint8(this.getValue(fieldName).(float32))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.unsignedShort" {
								val := uint16(this.getValue(fieldName).(float32))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.unsignedInt" {
								val := uint32(this.getValue(fieldName).(float32))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.unsignedLong" {
								val := uint64(this.getValue(fieldName).(float32))
								this.setValue(fieldName, val)
							} else {
								val := int64(this.getValue(fieldName).(float32))
								this.setValue(fieldName, val)
							}
						} else if reflect.TypeOf(this.getValue(fieldName)).Kind() == reflect.Float64 {
							if fieldType == "xs.int" || fieldType == "xs.integer" {
								val := int32(this.getValue(fieldName).(float64))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.byte" {
								val := int8(this.getValue(fieldName).(float64))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.short" {
								val := int16(this.getValue(fieldName).(float64))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.unsignedByte" {
								val := uint8(this.getValue(fieldName).(float64))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.unsignedShort" {
								val := uint16(this.getValue(fieldName).(float64))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.unsignedInt" {
								val := uint32(this.getValue(fieldName).(float64))
								this.setValue(fieldName, val)
							} else if fieldType == "xs.unsignedLong" {
								val := uint64(this.getValue(fieldName).(float64))
								this.setValue(fieldName, val)
							} else {
								val := int64(this.getValue(fieldName).(float64))
								this.setValue(fieldName, val)
							}
						}
					}

					//log.Println(fieldType, this.object[fieldName], reflect.TypeOf(this.object[fieldName]))
					DynamicEntityInfo = append(DynamicEntityInfo, this.getValue(fieldName))
				} else {
					if this.isRef(fieldType, fieldName) {
						// Must be a string or an object...
						var refId string
						if reflect.TypeOf(this.getValue(fieldName)).String() == "string" {
							refId = this.getValue(fieldName).(string)
						} else if reflect.TypeOf(this.getValue(fieldName)).String() == "map[string]interface {}" {
							refId = this.getValue(fieldName).(map[string]interface{})["UUID"].(string)
						}

						if len(refId) > 0 {
							// Reference will be treated latter...
							var ref EntityRef
							ref.Name = fieldName
							ref.OwnerUuid = this.GetUuid()

							if strings.HasPrefix(refId, "#") {
								refId = refId[1:]
							}

							refUuid := refId
							refTypeName := strings.Replace(fieldType, ":Ref", "", -1)
							id := refTypeName + "$$" + refUuid
							GetServer().GetEntityManager().appendReferenced(fieldName, this.GetUuid(), id)

							// Append the reference here whit it base name...
							if Utility.IsValidEntityReferenceName(refId) {
								DynamicEntityInfo = append(DynamicEntityInfo, refId)
							} else {
								// need to be set latter...
								DynamicEntityInfo = append(DynamicEntityInfo, "null")
							}

						} else {
							DynamicEntityInfo = append(DynamicEntityInfo, "null")
						}

					} else {
						if this.getValue(fieldName) != nil {
							// Only value with object map will be convert into ther given type...
							if reflect.TypeOf(this.getValue(fieldName)).String() == "map[string]interface {}" {
								subValues := this.getValue(fieldName).(map[string]interface{})

								// I will create the sub value...
								typeName := strings.Replace(strings.Replace(fieldType, ":Ref", "", -1), "[]", "", -1)
								var uuid string
								if subValues["UUID"] != nil {
									uuid = subValues["UUID"].(string)
								}

								// I will try to create a static entity...
								newEntityMethod := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
								params := make([]interface{}, 3)
								params[0] = ""
								params[1] = uuid
								params[2] = subValues
								staticEntity, err := Utility.CallMethod(GetServer().GetEntityManager(), newEntityMethod, params)

								if err != nil {
									// In that case I will try with dynamic entity.
									// I will create the sub value...
									subEntity, errObj := GetServer().GetEntityManager().newDynamicEntity(this.GetUuid(), subValues)
									if errObj == nil {
										subEntity.AppendReferenced(fieldName, this)
										this.AppendChild(fieldName, subEntity)
										DynamicEntityInfo = append(DynamicEntityInfo, subEntity.GetUuid())
										if subEntity.NeedSave() {
											subEntity.saveEntity(path + "|" + this.GetUuid())
										}
									} else {
										log.Println("-----> 979 error  ", errObj.GetBody())
									}
								} else {
									staticEntity.(Entity).AppendReferenced(fieldName, this)
									this.AppendChild(fieldName, staticEntity.(Entity))
									DynamicEntityInfo = append(DynamicEntityInfo, uuid)
									if staticEntity.(Entity).NeedSave() {
										staticEntity.(Entity).SaveEntity()
									}
								}
							} else if reflect.TypeOf(this.getValue(fieldName)).String() == "string" {
								DynamicEntityInfo = append(DynamicEntityInfo, this.getValue(fieldName).(string))
							}
						}
					}
				}
			}
		} else {
			DynamicEntityInfo = append(DynamicEntityInfo, "null")
		}
	}

	// The childs uuid
	query.Fields = append(query.Fields, "childsUuid")

	// Finalyse save here...
	childsUuidStr, _ := json.Marshal(this.childsUuid)
	if len(childsUuidStr) > 0 {
		DynamicEntityInfo = append(DynamicEntityInfo, string(childsUuidStr))
	} else {
		DynamicEntityInfo = append(DynamicEntityInfo, "")
	}

	// The referenced
	query.Fields = append(query.Fields, "referenced")
	referencedStr, _ := json.Marshal(this.referenced)
	if len(childsUuidStr) > 0 {
		DynamicEntityInfo = append(DynamicEntityInfo, string(referencedStr))
	} else {
		DynamicEntityInfo = append(DynamicEntityInfo, "")
	}

	// The event data...
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.TYPENAME = "Server.MessageData"
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData

	var evt *Event
	var err error

	storeId := this.GetPackageName()
	if reflect.TypeOf(dataManager.getDataStore(storeId)).String() == "*Server.SqlDataStore" {
		storeId = "sql_info" // Must save or update value from sql info instead.
	}

	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(storeId, string(queryStr), DynamicEntityInfo, params)
	} else {
		log.Println("-----> ", this.GetUuid(), " not exist!")
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		// Save the values for that entity.
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(storeId, string(queryStr), DynamicEntityInfo)
	}

	if err == nil {
		if storeId == "sql_info" {
			// Now I will initialyse references
			dataManager.setEntityReferences(this.GetUuid(), false)
		}
		// Send the event.
		GetServer().GetEventManager().BroadcastEvent(evt)
		// resolved reference pointing to this entity and not already append...
		GetServer().GetEntityManager().saveReferenced(this)
		GetServer().GetEntityManager().setReferences(this)

	} else {
		log.Println(Utility.FileLine(), "Fail to save entity ", err)
	}

}

/**
 * Delete the entity
 */
func (this *DynamicEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/**
 * Remove a child uuid form the list of child's in an entity. The name is the
 * propertie name in the parent.
 */
func (this *DynamicEntity) RemoveChild(name string, uuid string) {

	childsUuid := make([]string, 0)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		}
	}

	this.childsUuid = childsUuid
	prototype := this.GetPrototype()

	// Now I will remove it's value in the object...
	fieldIndex := prototype.getFieldIndex(name)
	if fieldIndex > -1 {
		filedType := prototype.FieldsType[prototype.getFieldIndex(name)]
		isArray := strings.HasPrefix(filedType, "[]")
		if isArray {
			childs := make([]map[string]interface{}, 0)
			if this.getValue(name) != nil {
				if reflect.TypeOf(this.getValue(name)).String() == "[]map[string]interface {}" {
					for i := 0; i < len(this.getValue(name).([]map[string]interface{})); i++ {
						if uuid != this.getValue(name).([]map[string]interface{})[i]["UUID"] {
							childs = append(childs, this.getValue(name).([]map[string]interface{})[i])
						}
					}
				} else if reflect.TypeOf(this.getValue(name)).String() == "[]interface {}" {
					for i := 0; i < len(this.getValue(name).([]interface{})); i++ {
						if uuid != this.getValue(name).([]interface{})[i].(map[string]interface{})["UUID"] {
							childs = append(childs, this.getValue(name).([]interface{})[i].(map[string]interface{}))
						}
					}
				}
			}
			this.setValue(name, childs)
		} else {
			this.deleteValue(name)
		}
	} else {
		log.Println("--> "+prototype.TypeName+" dosent have field name ", name)
	}

}

/**
 * Return the type name of an entity
 */
func (this *DynamicEntity) GetTypeName() string {
	return this.getValue("TYPENAME").(string)
}

/**
 * The package name
 */
func (this *DynamicEntity) GetPackageName() string {
	typeName := this.GetTypeName()
	packageName := typeName[0:strings.Index(typeName, ".")]
	return packageName
}

/**
 * Each entity must have one uuid.
 */
func (this *DynamicEntity) GetUuid() string {
	return this.getValue("UUID").(string)
}

func (this *DynamicEntity) GetParentUuid() string {
	if this.getValue("ParentUuid") != nil {
		return this.getValue("ParentUuid").(string)
	}
	return ""
}

/**
 * The name of the relation with it parent.
 */
func (this *DynamicEntity) GetParentLnk() string {
	if this.getValue("ParentLnk") != nil {
		return this.getValue("ParentLnk").(string)
	}
	return ""
}

func (this *DynamicEntity) SetParentLnk(lnk string) {
	this.setValue("ParentLnk", lnk)
}

/**
 * If the entity is created by a parent entity.
 */
func (this *DynamicEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

/**
 * Return the list of references uuid of an entity
 */
func (this *DynamicEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

/**
 * Set reference uuid
 */
func (this *DynamicEntity) SetReferencesUuid(uuid []string) {
	this.referencesUuid = uuid
}

/**
 * Append a reference.
 */
func (this *DynamicEntity) AppendReference(reference Entity) {
	// Here i will append the reference uuid
	index := -1

	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}

	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

/**
 * Remove a reference from an entity.
 */
func (this *DynamicEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	refsPtr := make([]Entity, 0)

	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsPtr = append(refsPtr, reference)
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)

	// Now I will remove it from it it internal object to.
	prototype := this.GetPrototype()
	fieldType := prototype.FieldsType[prototype.getFieldIndex(name)]
	isArray := strings.HasPrefix(fieldType, "[]")
	if isArray {
		refs := make([]string, 0)
		for i := 0; i < len(this.getValue(name).([]string)); i++ {
			if reference.GetUuid() != this.getValue(name).([]string)[i] {
				refs = append(refs, this.getValue(name).([]string)[i])
			}
		}
		this.setValue(name, refs)
	} else {
		this.deleteValue(name)
	}

}

/**
 * Append a child...
 */
func (this *DynamicEntity) AppendChild(attributeName string, child Entity) error {

	// Append referenced.
	child.AppendReferenced(attributeName, this)

	// Set the parent Lnk.
	child.(*DynamicEntity).setValue("ParentLnk", attributeName)

	// I will retreive the field type.
	prototype := this.GetPrototype()
	fieldTypeIndex := prototype.getFieldIndex(attributeName)
	if fieldTypeIndex > 0 {
		fieldTypeName := prototype.FieldsType[fieldTypeIndex]
		if strings.HasPrefix(fieldTypeName, "[]") {
			// if the array is nil...
			if this.getValue(attributeName) == nil {
				if strings.HasSuffix(fieldTypeName, ":Ref") {
					this.setValue(attributeName, make([]string, 0))
				} else {
					this.setValue(attributeName, make([]map[string]interface{}, 0))
				}
			}
			// In case of a reference.
			if reflect.TypeOf(this.getValue(attributeName)).String() == "[]string" {
				if !Utility.Contains(this.getValue(attributeName).([]string), child.GetUuid()) {
					values := this.getValue(attributeName).([]string)
					this.setValue(attributeName, append(values, child.GetUuid()))
				}

			} else if reflect.TypeOf(this.getValue(attributeName)).String() == "[]map[string]interface {}" {
				// In case of an object...
				if this.getValue(attributeName) != nil {
					objects := this.getValue(attributeName).([]map[string]interface{})
					objects_ := make([]map[string]interface{}, 0)
					isExist := false
					for i := 0; i < len(objects); i++ {
						if objects[i]["UUID"].(string) != child.GetUuid() {
							objects_ = append(objects_, objects[i])
						} else {
							if reflect.TypeOf(child.GetObject()).String() == "map[string]interface {}" {
								objects_ = append(objects_, child.GetObject().(map[string]interface{}))
							} else {
								object, _ := Utility.ToMap(child.GetObject())
								objects_ = append(objects_, object)
							}
							isExist = true
						}
					}

					// append at the end of the list.
					if !isExist {
						if reflect.TypeOf(child.GetObject()).String() == "map[string]interface {}" {
							objects_ = append(objects_, child.GetObject().(map[string]interface{}))
						} else {
							object, _ := Utility.ToMap(child.GetObject())
							objects_ = append(objects_, object)
						}
					}

					// Set the array containing the new value.
					this.setValue(attributeName, objects_)
				} else {
					this.setValue(attributeName, make([]map[string]interface{}, 0))
					if reflect.TypeOf(child.GetObject()).String() == "map[string]interface {}" {
						values := this.getValue(attributeName).([]map[string]interface{})
						this.setValue(attributeName, append(values, child.GetObject().(map[string]interface{})))
					} else {
						object, _ := Utility.ToMap(child.GetObject())
						values := this.getValue(attributeName).([]map[string]interface{})
						this.setValue(attributeName, append(values, object))
					}
				}
			} else if reflect.TypeOf(this.getValue(attributeName)).String() == "[]interface {}" {
				if this.getValue(attributeName) != nil {
					objects := this.getValue(attributeName).([]interface{})
					objects_ := make([]interface{}, 0)
					isExist := false
					for i := 0; i < len(objects); i++ {
						var uuid string
						if objects[i].(map[string]interface{})["UUID"] != nil {
							uuid = objects[i].(map[string]interface{})["UUID"].(string)
						}
						if uuid != child.GetUuid() {
							objects_ = append(objects_, objects[i].(map[string]interface{}))
						} else {
							if reflect.TypeOf(child.GetObject()).String() == "map[string]interface {}" {
								objects_ = append(objects_, child.GetObject().(map[string]interface{}))
							} else {
								object, _ := Utility.ToMap(child.GetObject())
								objects_ = append(objects_, object)
							}
							isExist = true
						}
					}

					// append at the end of the list.
					if !isExist {
						if reflect.TypeOf(child.GetObject()).String() == "map[string]interface {}" {
							objects_ = append(objects_, child.GetObject().(map[string]interface{}))
						} else {
							object, _ := Utility.ToMap(child.GetObject())
							objects_ = append(objects_, object)
						}
					}

					// Set the array containing the new value.
					this.setValue(attributeName, objects_)
				} else {
					this.setValue(attributeName, make([]map[string]interface{}, 0))
					if reflect.TypeOf(child.GetObject()).String() == "map[string]interface {}" {
						values := this.getValue(attributeName).([]map[string]interface{})
						this.setValue(attributeName, append(values, child.GetObject().(map[string]interface{})))
					} else {
						object, _ := Utility.ToMap(child.GetObject())
						values := this.getValue(attributeName).([]map[string]interface{})
						this.setValue(attributeName, append(values, object))
					}
				}
			}

		} else {
			// Set or replace the value...
			if reflect.TypeOf(child.GetObject()).String() == "map[string]interface {}" {
				this.setValue(attributeName, child.GetObject())
			} else {
				object, _ := Utility.ToMap(child.GetObject())
				this.setValue(attributeName, object)
			}
		}
	}

	// Set the child list...
	if Utility.Contains(this.childsUuid, child.GetUuid()) == false {
		// Append it to the list of UUID
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}

	return nil
}

/**
 * Get the childs uuid...
 */
func (this *DynamicEntity) GetChildsUuid() []string {
	return this.childsUuid
}

/**
 * Set the array of childs uuid...
 */
func (this *DynamicEntity) SetChildsUuid(uuids []string) {
	this.childsUuid = uuids
}

/**
 * Append a entity that reference this entity.
 */
func (this *DynamicEntity) AppendReferenced(name string, owner Entity) {
	// Here I will try to find if the reference exist...
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == name && this.referenced[i].OwnerUuid == owner.GetUuid() {
			return
		}
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	this.referenced = append(this.referenced, ref)

	// Set need save.
	this.SetNeedSave(true)
}

/**
 * Return the list of reference.
 */
func (this *DynamicEntity) GetReferenced() []EntityRef {
	return this.referenced
}

/**
 * Remove the referenced.
 */
func (this *DynamicEntity) RemoveReferenced(name string, owner Entity) {
	this.Lock()
	defer this.Unlock()
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
}

/**
 * Return the object wrapped by this entity...
 */
func (this *DynamicEntity) GetObject() interface{} {
	this.Lock()
	defer this.Unlock()
	//return this.object.Items()
	return this.object
}

/**
 * Return the object wrapped by this entity...
 */
func (this *DynamicEntity) SetObjectValues(values map[string]interface{}) {

	// nothing to do if the pointer is nil...
	if this == nil {
		return
	}

	entity, exist := GetServer().GetEntityManager().contain(this.GetUuid())

	if !exist {
		// If the map is not in the cache I can set the values directly.
		this.setObject(values)
		this.SetNeedSave(true)
		return
	} else {
		// here I will set the need save attribute.
		sum0 := Utility.GetChecksum(values)
		sum1 := entity.GetChecksum()
		this.setObject(entity.GetObject().(map[string]interface{}))
		this.SetNeedSave(sum0 != sum1)
	}
	prototype := this.GetPrototype()
	for i := 0; i < len(prototype.Fields); i++ {
		field := prototype.Fields[i]
		fieldType := prototype.FieldsType[i]
		isArray := strings.HasPrefix(fieldType, "[]")
		isRef := strings.HasSuffix(fieldType, ":Ref")
		isBaseType := strings.HasPrefix(fieldType, "[]xs.") || strings.HasPrefix(fieldType, "xs.") || strings.HasPrefix(fieldType, "sqltypes.") || strings.HasPrefix(fieldType, "[]sqltypes.")

		// The goal here is to remove base value or reference if there is no
		// more in the values.
		if this.getValue(field) != nil {
			if strings.HasPrefix(field, "M_") {
				// remove nil value
				if isBaseType {
					if this.getValue(field) != nil && values[field] == nil {
						this.setValue(field, nil)
					}
				} else if isRef {
					v := values[field]
					if isArray {
						// Convert as needed...
						v_ := make([]string, 0)
						if v == nil {
							v = make([]string, 0)
						}
						if reflect.TypeOf(v).String() == "[]interface {}" {
							for i := 0; i < len(v.([]interface{})); i++ {
								if reflect.TypeOf(v.([]interface{})[i]).String() == "string" {
									v_ = append(v_, v.([]interface{})[i].(string))
								} else if reflect.TypeOf(v.([]interface{})[i]).String() == "map[string]interface {}" {
									v_ = append(v_, v.([]interface{})[i].(map[string]interface{})["UUID"].(string))
								}
							}
						} else if reflect.TypeOf(v).String() == "[]map[string]interface {}" {
							for i := 0; i < len(v.([]map[string]interface{})); i++ {
								v_ = append(v_, v.([]map[string]interface{})[i]["UUID"].(string))
							}
						} else if reflect.TypeOf(v).String() == "[]string" {
							v_ = v.([]string)
						}

						// Here I will remove deleted references.
						if reflect.TypeOf(this.getValue(field)).String() == "[]string" {
							for i := 0; i < len(this.getValue(field).([]string)); i++ {
								if !Utility.Contains(v_, this.getValue(field).([]string)[i]) {
									toRemove := this.getValue(field).([]string)[i]
									ref, err := GetServer().GetEntityManager().getEntityByUuid(toRemove, false)
									if err == nil {
										this.RemoveReference(field, ref)
										ref.RemoveReferenced(field, this)
									}
								}
							}
						} else {
							log.Println("---------> wrong references types object.")
						}

						// Set the references links...
						for i := 0; i < len(v_); i++ {
							ref, err := GetServer().GetEntityManager().getEntityByUuid(v_[i], true)
							if err == nil {
								ref.AppendReferenced(field, this)
								this.AppendReference(ref)
							}
						}

						// Set the ref values...
						this.setValue(field, v_)

					} else {
						var toRemove string

						if v == nil && this.getValue(field) != nil {
							toRemove = this.getValue(field).(string)
						} else if v != nil && this.getValue(field) != nil {
							if reflect.TypeOf(this.getValue(field)).String() == "map[string]interface {}" {
								if v.(map[string]interface{})["UUID"] != this.getValue(field) {
									toRemove = this.getValue(field).(string)
								}
							} else if reflect.TypeOf(this.getValue(field)).String() == "string" {
								if v != this.getValue(field) {
									toRemove = this.getValue(field).(string)
								}
							}
						}
						if len(toRemove) > 0 {
							ref, err := GetServer().GetEntityManager().getEntityByUuid(toRemove, true)
							if err == nil {
								this.RemoveReference(field, ref)
								ref.RemoveReferenced(field, this)
							}
						}

						// Set the reference.
						if v != nil {
							var refUuid string
							if reflect.TypeOf(v).String() == "map[string]interface {}" {
								refUuid = v.(map[string]interface{})["UUID"].(string)
							} else if reflect.TypeOf(v).String() == "string" {
								refUuid = v.(string)
							}

							if len(refUuid) > 0 {
								ref, err := GetServer().GetEntityManager().getEntityByUuid(refUuid, true)
								if err == nil {
									ref.AppendReferenced(field, this)
									this.AppendReference(ref)
								}
								this.setValue(field, refUuid)
							}
						} else {
							this.setValue(field, nil)
						}
					}
				}
			}
		}
	}

	// Here the object exist so I need to copy values inside the existing map.
	for k, v := range values {

		if v == nil {
			this.deleteValue(k)
		} else {
			// Only propertie with M_ will be set here.
			fieldIndex := prototype.getFieldIndex(k)
			if strings.HasPrefix(k, "M_") && fieldIndex != -1 {
				fieldType := prototype.FieldsType[fieldIndex]
				isRef := strings.HasSuffix(fieldType, ":Ref")
				isArray := strings.HasPrefix(fieldType, "[]")

				if !isRef {
					if this.getValue(k) == nil {
						// simply set the value...
						if v != nil {
							this.setValue(k, v)
						}
					} else {
						if reflect.TypeOf(this.getValue(k)).String() == "map[string]interface {}" {
							subValues := v.(map[string]interface{})
							if subValues["UUID"] != nil {
								// Here I have another dynamic object...
								subEntity, exist := GetServer().GetEntityManager().contain(subValues["UUID"].(string))
								if exist {
									// Set the sub entity
									subEntity.(*DynamicEntity).SetObjectValues(subValues)
								} else {
									var errObj *CargoEntities.Error
									subEntity, errObj = GetServer().GetEntityManager().newDynamicEntity(this.GetUuid(), subValues)
									if errObj == nil {
										this.setValue(k, subEntity.GetObject())
									}
								}
							} else {
								log.Println(Utility.FileLine(), " Map of interface without UUID found...!")
								this.setValue(k, v)
							}
						} else {
							if this.getValue(k) != nil {
								if reflect.TypeOf(this.getValue(k)).String() == "[]map[string]interface {}" {
									if reflect.TypeOf(v).String() == "[]map[string]interface {}" {
										for i := 0; i < len(v.([]map[string]interface{})); i++ {
											if v.([]map[string]interface{})[i]["UUID"] != nil {
												exist := false
												for j := 0; j < len(this.getValue(k).([]map[string]interface{})); j++ {
													if this.getValue(k).([]map[string]interface{})[j]["UUID"] == v.([]map[string]interface{})[i]["UUID"] {
														subEntity, _ := GetServer().GetEntityManager().getEntityByUuid(this.getValue(k).([]map[string]interface{})[j]["UUID"].(string), true)
														subEntity.(*DynamicEntity).SetObjectValues(v.([]map[string]interface{})[i])
														exist = true
													}
												}
												if !exist {
													// Here I need to append the new object.
													this.appendValue(k, v.([]map[string]interface{})[i])
												}
											} else {
												// Here the value is an object without uuid.
											}
										}
									} else if reflect.TypeOf(v).String() == "[]interface {}" {
										for i := 0; i < len(v.([]interface{})); i++ {
											if v.([]interface{})[i].(map[string]interface{})["UUID"] != nil {
												exist := false
												for j := 0; j < len(this.getValue(k).([]map[string]interface{})); j++ {
													if this.getValue(k).([]map[string]interface{})[j]["UUID"] == v.([]interface{})[i].(map[string]interface{})["UUID"] {
														subEntity, _ := GetServer().GetEntityManager().getEntityByUuid(this.getValue(k).([]map[string]interface{})[j]["UUID"].(string), true)
														subEntity.(*DynamicEntity).SetObjectValues(v.([]interface{})[i].(map[string]interface{}))
														exist = true
													}
												}
												if !exist {
													// Here I need to append the new object.
													this.appendValue(k, v.([]interface{})[i].(map[string]interface{}))
												}
											} else {
												// Here the value is an object without uuid.
											}
										}
									}
								} else if reflect.TypeOf(this.getValue(k)).String() == "[]interface {}" {
									if len(this.getValue(k).([]interface{})) == 0 {
										this.setValue(k, v)
									} else if reflect.TypeOf(this.getValue(k).([]interface{})[0]) != nil {
										if reflect.TypeOf(this.getValue(k).([]interface{})[0]).String() == "map[string]interface {}" {
											if reflect.TypeOf(v).String() == "[]map[string]interface {}" {
												for i := 0; i < len(v.([]map[string]interface{})); i++ {
													if v.([]map[string]interface{})[i]["UUID"] != nil {
														exist := false
														for j := 0; j < len(this.getValue(k).([]interface{})); j++ {
															if this.getValue(k).([]interface{})[j].(map[string]interface{})["UUID"] == v.([]map[string]interface{})[i]["UUID"] {
																subEntity, _ := GetServer().GetEntityManager().getEntityByUuid(this.getValue(k).([]interface{})[j].(map[string]interface{})["UUID"].(string), true)
																subEntity.(*DynamicEntity).SetObjectValues(v.([]map[string]interface{})[i])
																exist = true
															}
														}
														if !exist {
															// Here I need to append the new object.
															this.appendValue(k, v.([]map[string]interface{})[i])
														}
													} else {
														// Here the value is an object without uuid.
													}
												}
											} else if reflect.TypeOf(this.getValue(k).([]interface{})[0]).String() == "[]interface {}" {
												for i := 0; i < len(v.([]interface{})); i++ {
													if v.([]interface{})[i].(map[string]interface{})["UUID"] != nil {
														exist := false
														for j := 0; j < len(this.getValue(k).([]interface{})); j++ {
															if this.getValue(k).([]interface{})[j].(map[string]interface{})["UUID"] == v.([]interface{})[i].(map[string]interface{})["UUID"] {
																subEntity, _ := GetServer().GetEntityManager().getEntityByUuid(this.getValue(k).([]interface{})[j].(map[string]interface{})["UUID"].(string), true)
																subEntity.(*DynamicEntity).SetObjectValues(v.([]interface{})[i].(map[string]interface{}))
																exist = true
															}
														}
														if !exist {
															// Here I need to append the new object.
															this.appendValue(k, v.([]interface{})[i].(map[string]interface{}))
														}
													} else {
														// Here the value is an object without uuid.
													}
												}
											}
										} else {
											// Replace the array with the new value.
											this.setValue(k, v)
										}
									} else {
										// Replace the array with the new value.
										this.setValue(k, v)
									}
								} else if reflect.TypeOf(this.getValue(k)).String() == "[]string" {
									this.setValue(k, v) // Replace the array of string with the new value.
								} else if this.getValue(k) != v {
									if reflect.TypeOf(this.getValue(k)).Kind() == reflect.TypeOf(v).Kind() {
										this.setValue(k, v)
									} else {
										//log.Println("------> k : ", k, " v :", v, " kind ", reflect.TypeOf(this.getValue(k)).Kind(), " ", reflect.TypeOf(v).Kind())
										// Here I need to convert the v type...
										if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Float64 && reflect.TypeOf(v).Kind() == reflect.String {
											val, _ := strconv.ParseFloat(v.(string), 64)
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Float32 && reflect.TypeOf(v).Kind() == reflect.String {
											val, _ := strconv.ParseFloat(v.(string), 32)
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Int64 && reflect.TypeOf(v).Kind() == reflect.String {
											val, _ := strconv.ParseInt(v.(string), 10, 64)
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Int8 && reflect.TypeOf(v).Kind() == reflect.String {
											val, _ := strconv.ParseInt(v.(string), 10, 8)
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Int && reflect.TypeOf(v).Kind() == reflect.String {
											val, _ := strconv.ParseInt(v.(string), 10, 64)
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Int32 && reflect.TypeOf(v).Kind() == reflect.String {
											val, _ := strconv.ParseInt(v.(string), 10, 32)
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Int16 && reflect.TypeOf(v).Kind() == reflect.String {
											val, _ := strconv.ParseInt(v.(string), 10, 16)
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Bool && reflect.TypeOf(v).Kind() == reflect.String {
											val, _ := strconv.ParseBool(v.(string))
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Int64 && reflect.TypeOf(v).Kind() == reflect.Float64 {
											// the json parser transform all numerical value to float... that not what we want here...
											val := int64(v.(float64))
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Int32 && reflect.TypeOf(v).Kind() == reflect.Float64 {
											// the json parser transform all numerical value to float... that not what we want here...
											val := int32(v.(float64))
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Int16 && reflect.TypeOf(v).Kind() == reflect.Float64 {
											// the json parser transform all numerical value to float... that not what we want here...
											val := int16(v.(float64))
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.Int8 && reflect.TypeOf(v).Kind() == reflect.Float64 {
											// the json parser transform all numerical value to float... that not what we want here...
											val := int8(v.(float64))
											this.setValue(k, val)
										} else if reflect.TypeOf(this.getValue(k)).Kind() == reflect.String && reflect.TypeOf(v).Kind() == reflect.Float64 {
											val := Utility.ToString(v)
											this.setValue(k, val)
										} else {
											log.Println("-------------> k value", this.getValue(k))
											log.Println("-------------> v", v)
											log.Panicln("---------> Convertion error: from ", reflect.TypeOf(this.getValue(k)).Kind(), " to ", reflect.TypeOf(v).Kind())
										}
									}

								}
							} else {
								// Append the new value here...
								this.setValue(k, v) // Replace the array of string with the new value.
							}
						}
					}
				} else {
					// Here I have a reference. I need to keep only the uuid
					// in case of I receive the complete objet value.
					if isArray {
						if this.getValue(k) == nil {
							this.setValue(k, make([]string, 0))
						}
						if reflect.TypeOf(v).String() == "[]map[string] interface{}" {
							for i := 0; i < len(v.([]map[string]interface{})); i++ {
								val := this.getValue(k).([]string)
								this.setValue(k, append(val, v.(map[string]interface{})["UUID"].(string)))
							}
						} else if reflect.TypeOf(v).String() == "[]interface {}" {
							for i := 0; i < len(v.([]interface{})); i++ {
								if reflect.TypeOf(v.([]interface{})[i]).String() == "map[string] interface{}" {
									val := this.getValue(k).([]string)
									this.setValue(k, append(val, v.([]interface{})[i].(map[string]interface{})["UUID"].(string)))
								} else if reflect.TypeOf(v.([]interface{})[i]).String() == "string" {
									val := this.getValue(k).([]string)
									this.setValue(k, append(val, v.([]interface{})[i].(string)))
								}
							}
						} else if reflect.TypeOf(v).String() == "[]string" {
							// and array of string.
							this.setValue(k, v)
						}
					} else {
						if reflect.TypeOf(v).String() == "map[string]interface {}" {
							this.setValue(k, v.(map[string]interface{})["UUID"])
						} else if reflect.TypeOf(v).String() == "string" {
							this.setValue(k, v)
						}
					}
				}
			} else {
				if strings.HasPrefix(k, "M_") {
					log.Println(prototype.TypeName, " has no field ", k)
				}
			}
		}
	}
}

/**
 * Return the entity prototype.
 */
func (this *DynamicEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/**
 * Calculate a unique value for a given entity object value...
 */
func (this *DynamicEntity) GetChecksum() string {
	return Utility.GetChecksum(this.GetObject())
}

/**
 * Calculate a unique value for a given entity object value...
 */
func (this *DynamicEntity) GetSize() uint {
	return uint(unsafe.Sizeof(this.object))
}

/**
 * Check if the an entity exist...
 */
func (this *DynamicEntity) Exist() bool {
	// if the 'this' pointer is not initialyse it means the entity was deleted.
	if this == nil {
		return false
	}
	prototype := this.GetPrototype()
	var query EntityQuery
	query.TypeName = this.GetTypeName()
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, prototype.Ids...) // Get all it ids...
	var fieldsType []interface{}                          // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	storeId := this.GetPackageName()

	// In case of search of sql data.
	if reflect.TypeOf(dataManager.getDataStore(storeId)).String() == "*Server.SqlDataStore" {
		storeId = "sql_info" // Must save or update value from sql info instead.
	}

	results, err := GetServer().GetDataManager().readData(storeId, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		log.Println("--------------------> Not exist: ", this.GetUuid())
		return false
	}

	if len(results) == 0 {
		return false
	}

	if len(results[0]) == 0 {
		return false
	}

	return true
}

/**
 * Determine if the field is a reference.
 */
func (this *DynamicEntity) isRef(fieldType string, field string) bool {

	isRef_ := strings.HasSuffix(fieldType, ":Ref") || strings.HasSuffix(fieldType, ".anyURI") || strings.HasSuffix(fieldType, ".IDREF") || strings.HasSuffix(fieldType, ".IDREFS")
	return isRef_
}

func toJsonStr(object interface{}) string {
	b, _ := json.Marshal(object)
	// Convert bytes to string.
	b, _ = Utility.PrettyPrint(b)
	s := string(b)
	return s
}
