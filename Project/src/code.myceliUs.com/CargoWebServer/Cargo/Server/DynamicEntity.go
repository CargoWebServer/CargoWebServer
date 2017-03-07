package Server

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"code.myceliUs.com/Utility"
	"github.com/pborman/uuid"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
)

var (
	// The stirng represent the id of the reference.
	entityById = make(map[string]*DynamicEntity)
)

/**
 * Implementation of the Entity. Dynamic entity use a map[string]interface {}
 * to store object informations. All information in the Entity itself other
 * than object are store.
 */
type DynamicEntity struct {

	// The entity uuid.
	uuid string

	// Hard link parent->[]childs
	parentPtr  Entity
	parentUuid string
	childsPtr  []Entity
	childsUuid []string

	// Soft link []referenced->[]references
	referencesPtr  []Entity
	referencesUuid []string
	referenced     []EntityRef
	prototype      *EntityPrototype

	/** object will be a map... **/
	object map[string]interface{}

	/**
	 * Use to protected the entitiesMap access...
	 */
	sync.RWMutex

	/**
	 * Mutex to use with recursion.
	 */
	lock Utility.Rmutex
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
func (this *EntityManager) newDynamicEntity(values map[string]interface{}) (*DynamicEntity, *CargoEntities.Error) {
	this.Lock()
	defer this.Unlock()

	var entity *DynamicEntity

	if len(values["UUID"].(string)) > 0 {
		if val, ok := this.contain(values["UUID"].(string)); ok {
			if val != nil {
				entity = val.(*DynamicEntity)

				if Utility.GetChecksum(values) != entity.GetChecksum() {
					entity.SetObjectValues(values)
				} else {
					entity.SetNeedSave(false)
				}
				return entity, nil
			}
		}

		// Here if the enity is nil I will create a new instance.
		if entity == nil {
			entity = new(DynamicEntity)

			// If the object contain an id...
			entity.uuid = values["UUID"].(string)
			prototype, err := getEntityPrototype(values)

			if err != nil {
				// Create the error message
				cargoError := NewError(Utility.FileLine(), PROTOTYPE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Prototype not found for type '"+values["TYPENAME"].(string)+"'."))
				return nil, cargoError
			}

			// keep the reference to the prototype.
			entity.prototype = prototype

			// If the entity does not exist I will generate it new uuid...
			if len(entity.uuid) == 0 {
				// Here there is a new entity...
				entity.uuid = values["TYPENAME"].(string) + "%" + uuid.NewRandom().String()
				values["UUID"] = entity.uuid
			}

			entity.childsUuid = make([]string, 0)
			entity.referencesUuid = make([]string, 0)
			entity.referenced = make([]EntityRef, 0)

			// I will set it parent ptr...
			if values["parentUuid"] != nil {
				if len(values["parentUuid"].(string)) > 0 {
					parentPtr, _ := GetServer().GetEntityManager().getDynamicEntityByUuid(values["parentUuid"].(string))
					entity.SetParentPtr(parentPtr)
				}
			}
		}
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
 * Determine is an entity is initialyse or not.
 */
func (this *DynamicEntity) IsInit() bool {
	return this.GetObject().(map[string]interface{})["IsInit"].(bool)
}

/**
 * Set if an entity must be inityalyse.
 */
func (this *DynamicEntity) SetInit(isInit bool) {
	this.GetObject().(map[string]interface{})["IsInit"] = isInit
}

/**
 * Test if an entity need to be save.
 */
func (this *DynamicEntity) NeedSave() bool {
	return this.GetObject().(map[string]interface{})["NeedSave"].(bool)
}

/**
 * Set if an entity need to be save.
 */
func (this *DynamicEntity) SetNeedSave(needSave bool) {
	this.GetObject().(map[string]interface{})["NeedSave"] = needSave
}

/**
 * Initialyse a entity with a given id.
 */
func (this *DynamicEntity) InitEntity(id string) error {

	token := Utility.NewToken()
	this.lock.Lock(token)

	err := this.initEntity(id, "")
	this.lock.Unlock(token)
	//log.Println("After init:", toJsonStr(this.object))
	return err
}

func (this *DynamicEntity) initEntity(id string, path string) error {

	// cut infinite recursion here.
	if strings.Index(path, id) != -1 {
		return nil
	}

	// If the value is already in the cache I have nothing todo...
	if this.object["IsInit"].(bool) == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*DynamicEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.SetInit(false)
	}

	// I will set the id, (must be a uuid...)
	this.uuid = id

	typeName := id[0:strings.Index(id, "%")]
	packageName := typeName[0:strings.Index(typeName, ".")]

	var query EntityQuery
	query.TypeName = typeName

	// Here I will append the rest of the fields...
	// append the list of fields...
	query.Fields = append(query.Fields, this.prototype.Fields...)

	// The index of search...
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	var fieldsType []interface{} // not use...
	var params []interface{}

	storeId := packageName
	if reflect.TypeOf(dataManager.getDataStore(storeId)).String() == "*Server.SqlDataStore" {
		storeId = "sql_info" // Must save or update value from sql info instead.
	}

	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(storeId, string(queryStr), fieldsType, params)
	if err != nil {
		log.Fatalln(Utility.FileLine(), "--> No object found for entity ", id, " ", err)
		return err
	}

	// Initialisation of information of Interface...
	if len(results) > 0 {
		// Set the common values...
		this.object["UUID"] = results[0][0].(string)
		this.object["TYPENAME"] = typeName // Set the typeName

		// Set the parent uuid...
		this.parentUuid = results[0][1].(string)
		this.object["parentUuid"] = this.parentUuid // Set the parent uuid.

		//init the child...
		childsUuidStr := results[0][this.prototype.getFieldIndex("childsUuid")].(string)

		if len(childsUuidStr) > 0 {
			err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)
			if err != nil {
				log.Println(Utility.FileLine(), " unable to read child uuid! for entity ", this.uuid, " ", childsUuidStr)
				return err
			}
		}

		referencedStr := results[0][this.prototype.getFieldIndex("referenced")].(string)
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
		for i := 2; i < len(results[0])-2; i++ {
			fieldName := this.prototype.Fields[i]
			fieldType := this.prototype.FieldsType[i]
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
							this.object[fieldName] = values
						}
					} else {
						// If the field is a reference.
						if this.isRef(fieldType, fieldName) {
							// Here I got an array for reference string.
							var referencesId []string

							json.Unmarshal([]byte(results[0][i].(string)), &referencesId)

							refTypeName := fieldType

							this.object[fieldName] = referencesId
							for i := 0; i < len(referencesId); i++ {
								id_ := refTypeName + "$$" + referencesId[i]
								GetServer().GetEntityManager().appendReference(fieldName, this.object["UUID"].(string), id_)
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
								// Set an empty array here...
								if Utility.IsValidEntityReferenceName(uuids[0]) {
									this.object[fieldName] = make([]map[string]interface{}, 0)
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
												values["TYPENAME"] = typeName

												// I will try to create a static entity...
												newEntityMethod := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
												params := make([]interface{}, 2)
												params[0] = uuids[i]
												params[1] = values
												staticEntity, err := Utility.CallMethod(GetServer().GetEntityManager(), newEntityMethod, params)

												if err != nil {
													// In that case I will try with dynamic entity.
													dynamicEntity, errObj := GetServer().GetEntityManager().newDynamicEntity(values)
													if errObj == nil {
														// initialise the sub entity.
														err := dynamicEntity.initEntity(uuids[i], path+"|"+this.GetUuid())
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
													staticEntity.(Entity).InitEntity(uuids[i])
													staticEntity.(Entity).AppendReferenced(fieldName, this)
													this.AppendChild(fieldName, staticEntity.(Entity))
												}

											}
										}
									}
								}
							}
						}
					}
				} else {
					if strings.HasPrefix(fieldType, "xs.") || strings.HasPrefix(fieldType, "sqltypes.") || fieldName == "M_valueOf" {
						this.object[fieldName] = results[0][i]
					} else {
						// Not an array here.
						if this.isRef(fieldType, fieldName) {
							// In that case the reference must be a simple string.
							refTypeName := fieldType
							id_ := refTypeName + "$$" + results[0][i].(string)
							this.object[fieldName] = results[0][i].(string)
							GetServer().GetEntityManager().appendReference(fieldName, this.object["UUID"].(string), id_)
						} else {
							// Here I have an object...
							if reflect.TypeOf(results[0][i]).String() == "string" {
								uuid := results[0][i].(string)
								if Utility.IsValidEntityReferenceName(uuid) {
									if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {
										dynamicEntity := instance.(*DynamicEntity)
										dynamicEntity.AppendReferenced(fieldName, this)
										this.AppendChild(fieldName, dynamicEntity)
									} else {
										typeName := strings.Replace(fieldType, ":Ref", "", -1)
										values := make(map[string]interface{}, 0)
										values["UUID"] = uuid
										values["TYPENAME"] = typeName

										// I will try to create a static entity...
										newEntityMethod := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
										params := make([]interface{}, 2)
										params[0] = uuid
										params[1] = values
										staticEntity, err := Utility.CallMethod(GetServer().GetEntityManager(), newEntityMethod, params)

										if err != nil {
											// In that case I will try with dynamic entity.
											dynamicEntity, errObj := GetServer().GetEntityManager().newDynamicEntity(values)
											if errObj == nil {
												// initialise the sub entity.
												err := dynamicEntity.initEntity(uuid, path+"|"+this.GetUuid())
												if err != nil {
													// I will try to remove it from it child if is there...
													this.RemoveChild(fieldName, uuid)
													this.SetNeedSave(true)
												} else {
													dynamicEntity.AppendReferenced(fieldName, this)
													this.AppendChild(fieldName, dynamicEntity)
												}
											}
										} else {
											staticEntity.(Entity).InitEntity(uuid)
											staticEntity.(Entity).AppendReferenced(fieldName, this)
											this.AppendChild(fieldName, staticEntity.(Entity))
										}
									}
								} else {
									// A plain string...
									this.object[fieldName] = results[0][i]
								}
							} else {
								// A base type other than strings... bool, float64, int64 etc...
								this.object[fieldName] = results[0][i]
							}
						}
					}
				}
			} else {
				this.object[fieldName] = nil
			}
		}
	}

	// set init done.
	this.SetInit(true)
	this.SetNeedSave(false)

	GetServer().GetEntityManager().InitEntity(this)

	// if some change are found at initialysation I will update the values.
	if this.NeedSave() == true {
		// The entity need to be save...
		this.saveEntity("")
	}

	return nil
}

/**
 * Save entity
 */
func (this *DynamicEntity) SaveEntity() {
	token := Utility.NewToken()

	this.lock.Lock(token)
	this.saveEntity("")
	this.lock.Unlock(token)

	//log.Println("After save:", toJsonStr(this.object))

}

func (this *DynamicEntity) saveEntity(path string) {

	// Do not save if is nt necessary...
	if !this.NeedSave() || strings.Index(path, this.GetUuid()) != -1 {
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
	query.Fields = append(query.Fields, "uuid")
	DynamicEntityInfo = append(DynamicEntityInfo, this.GetUuid())

	query.Fields = append(query.Fields, "parentUuid")
	if len(this.parentUuid) > 0 {
		DynamicEntityInfo = append(DynamicEntityInfo, this.parentUuid)
		this.object["parentUuid"] = this.parentUuid
	} else {
		DynamicEntityInfo = append(DynamicEntityInfo, "")
	}

	// Now the fields of the object, from the prototype...
	// The first tow field are uuid and parentUuid
	// and the last tow fields are childsuuid and referenced...
	for i := 2; i < len(this.prototype.Fields)-2; i++ {
		fieldName := this.prototype.Fields[i]
		fieldType := this.prototype.FieldsType[i]

		// append the field in the query fields list.
		query.Fields = append(query.Fields, fieldName)

		// Print the field name.
		if this.object[fieldName] != nil {
			log.Println("--------> save ", fieldName, " whit value ", this.object[fieldName], " field type ", fieldType)
			// Array's
			if strings.HasPrefix(fieldType, "[]") {
				if strings.HasPrefix(fieldType, "[]xs.") || strings.HasPrefix(fieldType, "[]sqltypes.") || fieldName == "M_listOf" {

					valuesStr, err := json.Marshal(this.object[fieldName])
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
						if reflect.TypeOf(this.object[fieldName]).String() == "[]interface {}" {
							// An array of interface type.
							for j := 0; j < len(this.object[fieldName].([]interface{})); j++ {
								// Append reference.
								obj := this.object[fieldName].([]interface{})[j]
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

						} else if reflect.TypeOf(this.object[fieldName]).String() == "[]string" {
							for j := 0; j < len(this.object[fieldName].([]string)); j++ {
								refUuid := this.object[fieldName].([]string)[j]
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
							log.Println(Utility.FileLine(), "---> field no found!", fieldName, ":", fieldType, ":", reflect.TypeOf(this.object[fieldName]).String())
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
						if reflect.TypeOf(this.object[fieldName]).String() == "[]map[string]interface {}" {
							subEntityIds := make([]string, 0)
							for j := 0; j < len(this.object[fieldName].([]map[string]interface{})); j++ {
								// I will get the values of the sub item...
								subValues := this.object[fieldName].([]map[string]interface{})[j]

								typeName := strings.Replace(strings.Replace(fieldType, ":Ref", "", -1), "[]", "", -1)
								uuid := subValues["UUID"].(string)

								// I will try to create a static entity...
								newEntityMethod := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
								params := make([]interface{}, 2)
								params[0] = uuid
								params[1] = subValues
								staticEntity, err := Utility.CallMethod(GetServer().GetEntityManager(), newEntityMethod, params)

								if err != nil {
									// In that case I will try with dynamic entity.
									// I will create the sub value...
									subEntity, errObj := GetServer().GetEntityManager().newDynamicEntity(subValues)
									if errObj == nil {
										subEntity.AppendReferenced(fieldName, this)
										this.AppendChild(fieldName, subEntity)
										subEntityIds = append(subEntityIds, subEntity.uuid)
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

						} else if reflect.TypeOf(this.object[fieldName]).String() == "[]interface {}" {
							subEntityIds := make([]string, 0)
							for j := 0; j < len(this.object[fieldName].([]interface{})); j++ {
								if reflect.TypeOf(this.object[fieldName].([]interface{})[j]).String() == "map[string]interface {}" {
									// I will get the values of the sub item...
									subValues := this.object[fieldName].([]interface{})[j].(map[string]interface{})
									// I will create the sub value...
									typeName := strings.Replace(strings.Replace(fieldType, ":Ref", "", -1), "[]", "", -1)
									uuid := subValues["UUID"].(string)

									// I will try to create a static entity...
									newEntityMethod := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
									params := make([]interface{}, 2)
									params[0] = uuid
									params[1] = subValues
									staticEntity, err := Utility.CallMethod(GetServer().GetEntityManager(), newEntityMethod, params)

									if err != nil {
										// In that case I will try with dynamic entity.
										// I will create the sub value...
										subEntity, errObj := GetServer().GetEntityManager().newDynamicEntity(subValues)
										if errObj == nil {
											subEntity.AppendReferenced(fieldName, this)
											this.AppendChild(fieldName, subEntity)
											subEntityIds = append(subEntityIds, subEntity.uuid)
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
									log.Println(Utility.FileLine(), "---> field no found!", fieldName, ":", fieldType, ":", reflect.TypeOf(this.object[fieldName].([]interface{})[j]).String())
								}
							}
							subEntityIdsStr, _ := json.Marshal(subEntityIds)
							DynamicEntityInfo = append(DynamicEntityInfo, string(subEntityIdsStr))
						}
					}
				}
			} else {
				// Not an array...
				if strings.HasPrefix(fieldType, "xs.") || strings.HasPrefix(fieldType, "sqltypes.") || fieldName == "M_valueOf" {
					// Little fix to convert float into int type as needed.
					if fieldType == "xs.byte" || fieldType == "xs.short" || fieldType == "xs.int" || fieldType == "xs.integer" || fieldType == "xs.long" || fieldType == "xs.unsignedInt" || fieldType == "xs.unsignedByte" || fieldType == "xs.unsignedShort" || fieldType == "xs.unsignedLong" {
						if reflect.TypeOf(this.object[fieldName]).Kind() == reflect.Float32 {
							if fieldType == "xs.int" || fieldType == "xs.integer" {
								this.object[fieldName] = int32(this.object[fieldName].(float32))
							} else if fieldType == "xs.byte" {
								this.object[fieldName] = int8(this.object[fieldName].(float32))
							} else if fieldType == "xs.short" {
								this.object[fieldName] = int16(this.object[fieldName].(float32))
							} else if fieldType == "xs.unsignedByte" {
								this.object[fieldName] = uint8(this.object[fieldName].(float32))
							} else if fieldType == "xs.unsignedShort" {
								this.object[fieldName] = uint16(this.object[fieldName].(float32))
							} else if fieldType == "xs.unsignedInt" {
								this.object[fieldName] = uint32(this.object[fieldName].(float32))
							} else if fieldType == "xs.unsignedLong" {
								this.object[fieldName] = uint64(this.object[fieldName].(float32))
							} else {
								this.object[fieldName] = int64(this.object[fieldName].(float32))
							}
						} else if reflect.TypeOf(this.object[fieldName]).Kind() == reflect.Float64 {
							if fieldType == "xs.int" || fieldType == "xs.integer" {
								this.object[fieldName] = int32(this.object[fieldName].(float64))
							} else if fieldType == "xs.byte" {
								this.object[fieldName] = int8(this.object[fieldName].(float64))
							} else if fieldType == "xs.short" {
								this.object[fieldName] = int16(this.object[fieldName].(float64))
							} else if fieldType == "xs.unsignedByte" {
								this.object[fieldName] = uint8(this.object[fieldName].(float64))
							} else if fieldType == "xs.unsignedShort" {
								this.object[fieldName] = uint16(this.object[fieldName].(float64))
							} else if fieldType == "xs.unsignedInt" {
								this.object[fieldName] = uint32(this.object[fieldName].(float64))
							} else if fieldType == "xs.unsignedLong" {
								this.object[fieldName] = uint64(this.object[fieldName].(float64))
							} else {
								this.object[fieldName] = int64(this.object[fieldName].(float64))
							}
						}
					}

					//log.Println(fieldType, this.object[fieldName], reflect.TypeOf(this.object[fieldName]))
					DynamicEntityInfo = append(DynamicEntityInfo, this.object[fieldName])
				} else {
					if this.isRef(fieldType, fieldName) {
						// Must be a string or an object...
						var refId string
						if reflect.TypeOf(this.object[fieldName]).String() == "string" {
							refId = this.object[fieldName].(string)
						} else if reflect.TypeOf(this.object[fieldName]).String() == "map[string]interface {}" {
							refId = this.object[fieldName].(map[string]interface{})["UUID"].(string)
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
						if this.object[fieldName] != nil {
							// Only value with object map will be convert into ther given type...
							if reflect.TypeOf(this.object[fieldName]).String() == "map[string]interface {}" {
								subValues := this.object[fieldName].(map[string]interface{})

								// I will create the sub value...
								typeName := strings.Replace(strings.Replace(fieldType, ":Ref", "", -1), "[]", "", -1)
								uuid := subValues["UUID"].(string)

								// I will try to create a static entity...
								newEntityMethod := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
								params := make([]interface{}, 2)
								params[0] = uuid
								params[1] = subValues
								staticEntity, err := Utility.CallMethod(GetServer().GetEntityManager(), newEntityMethod, params)

								if err != nil {
									// In that case I will try with dynamic entity.
									// I will create the sub value...
									subEntity, errObj := GetServer().GetEntityManager().newDynamicEntity(subValues)
									if errObj == nil {
										subEntity.AppendReferenced(fieldName, this)
										this.AppendChild(fieldName, subEntity)
										DynamicEntityInfo = append(DynamicEntityInfo, subEntity.uuid)
										if subEntity.NeedSave() {
											subEntity.saveEntity(path + "|" + this.GetUuid())
										}
									}
								} else {
									staticEntity.(Entity).AppendReferenced(fieldName, this)
									this.AppendChild(fieldName, staticEntity.(Entity))
									DynamicEntityInfo = append(DynamicEntityInfo, uuid)
									if staticEntity.(Entity).NeedSave() {
										staticEntity.(Entity).SaveEntity()
									}
								}
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
	msgData.Name = "entity"
	msgData.Value = this.object
	eventData[0] = msgData

	var evt *Event
	var err error

	// Now I will register the entity in the entityByIdMap if is not nil
	if len(this.prototype.Ids) > 1 {
		id := this.prototype.Ids[1]
		if this.object[id] != nil {
			var objectId string
			if reflect.TypeOf(this.object[id]).Kind() == reflect.String {
				objectId = this.object[id].(string)
			} else if reflect.TypeOf(this.object[id]).Kind() == reflect.Int {
				objectId = strconv.Itoa(this.object[id].(int))
			} else if reflect.TypeOf(this.object[id]).Kind() == reflect.Int8 {
				objectId = strconv.Itoa(int(this.object[id].(int8)))
			} else if reflect.TypeOf(this.object[id]).Kind() == reflect.Int16 {
				objectId = strconv.Itoa(int(this.object[id].(int16)))
			} else if reflect.TypeOf(this.object[id]).Kind() == reflect.Int32 {
				objectId = strconv.Itoa(int(this.object[id].(int32)))
			} else if reflect.TypeOf(this.object[id]).Kind() == reflect.Int64 {
				objectId = strconv.Itoa(int(this.object[id].(int64)))
			}
			entityById[objectId] = this
		}
	}

	storeId := this.GetPackageName()
	if reflect.TypeOf(dataManager.getDataStore(storeId)).String() == "*Server.SqlDataStore" {
		storeId = "sql_info" // Must save or update value from sql info instead.
	}

	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "uuid="+this.uuid)
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(storeId, string(queryStr), DynamicEntityInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		// Save the values for that entity.
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(storeId, string(queryStr), DynamicEntityInfo)
	}

	if err == nil {
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
	token := Utility.NewToken()
	this.lock.Lock(token)
	GetServer().GetEntityManager().deleteEntity(this)
	this.lock.Unlock(token)
}

/**
 * Remove a chidl uuid form the list of child in an entity.
 */
func (this *DynamicEntity) RemoveChild(name string, uuid string) {

	childsUuid := make([]string, 0)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		}
	}

	this.childsUuid = childsUuid
	childsPtr := make([]Entity, 0)
	for i := 0; i < len(this.GetChildsPtr()); i++ {
		if this.GetChildsPtr()[i].GetUuid() != uuid {
			childsPtr = append(childsPtr, this.GetChildsPtr()[i])
		}
	}
	this.childsPtr = childsPtr

	// Now I will remove it's value in the object...
	filedType := this.prototype.FieldsType[this.prototype.getFieldIndex(name)]
	isArray := strings.HasPrefix(filedType, "[]")
	if isArray {
		childs := make([]map[string]interface{}, 0)
		if reflect.TypeOf(this.object[name]).String() == "[]map[string]interface {}" {
			for i := 0; i < len(this.object[name].([]map[string]interface{})); i++ {
				if uuid != this.object[name].([]map[string]interface{})[i]["UUID"] {
					childs = append(childs, this.object[name].([]map[string]interface{})[i])
				}
			}
		} else if reflect.TypeOf(this.object[name]).String() == "[]interface {}" {
			for i := 0; i < len(this.object[name].([]interface{})); i++ {
				if uuid != this.object[name].([]interface{})[i].(map[string]interface{})["UUID"] {
					childs = append(childs, this.object[name].([]interface{})[i].(map[string]interface{}))
				}
			}
		}
		this.object[name] = childs
	} else {
		delete(this.object, name)
	}

}

/**
 * Return the type name of an entity
 */
func (this *DynamicEntity) GetTypeName() string {
	return this.object["TYPENAME"].(string)
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
	return this.uuid
}

/**
 * If the entity is created by a parent entity.
 */
func (this *DynamicEntity) GetParentPtr() Entity {
	return this.parentPtr
}

func (this *DynamicEntity) SetParentPtr(parent Entity) {
	this.parentUuid = parent.GetUuid()
	this.parentPtr = parent
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
 * Return the list of reference of an entity
 */
func (this *DynamicEntity) GetReferencesPtr() []Entity {
	return this.referencesPtr
}

/**
 * Set reference uuid
 */
func (this *DynamicEntity) SetReferencesPtr(ref []Entity) {
	this.referencesPtr = ref
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
		this.referencesPtr = append(this.referencesPtr, reference)
	} else {
		// The reference must be update in that case.
		this.referencesPtr[index] = reference
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
	this.SetReferencesPtr(refsPtr)

	// Now I will remove it from it it internal object to.
	fieldType := this.prototype.FieldsType[this.prototype.getFieldIndex(name)]
	isArray := strings.HasPrefix(fieldType, "[]")
	if isArray {
		refs := make([]string, 0)
		for i := 0; i < len(this.object[name].([]string)); i++ {
			if reference.GetUuid() != this.object[name].([]string)[i] {
				refs = append(refs, this.object[name].([]string)[i])
			}
		}
		this.object[name] = refs
	} else {
		delete(this.object, name)
	}

}

/**
 * Return the list of child entity
 */
func (this *DynamicEntity) GetChildsPtr() []Entity {
	return this.childsPtr
}

/**
 * Set the array of childs ptr...
 */
func (this *DynamicEntity) SetChildsPtr(childsPtr []Entity) {
	this.childsPtr = childsPtr
}

/**
 * Append a child...
 */
func (this *DynamicEntity) AppendChild(attributeName string, child Entity) error {
	//log.Println("--------> append child: ", this.uuid, ":", attributeName, child.GetUuid())
	// Set or reset the child ptr.
	child.SetParentPtr(this)

	// I will retreive the field type.
	fieldTypeIndex := this.prototype.getFieldIndex(attributeName)
	if fieldTypeIndex > 0 {
		fieldTypeName := this.prototype.FieldsType[fieldTypeIndex]
		if strings.HasPrefix(fieldTypeName, "[]") {
			// if the array is nil...
			if this.object[attributeName] == nil {
				if strings.HasSuffix(fieldTypeName, ":Ref") {
					this.object[attributeName] = make([]string, 0)
				} else {
					this.object[attributeName] = make([]map[string]interface{}, 0)
				}
			}
			// In case of a reference.
			if reflect.TypeOf(this.object[attributeName]).String() == "[]string" {
				if !Utility.Contains(this.object[attributeName].([]string), child.GetUuid()) {
					this.object[attributeName] = append(this.object[attributeName].([]string), child.GetUuid())
				}

			} else if reflect.TypeOf(this.object[attributeName]).String() == "[]map[string]interface {}" {
				// In case of an object...
				if this.object[attributeName] != nil {
					objects := this.object[attributeName].([]map[string]interface{})
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
					this.object[attributeName] = objects_
				} else {
					this.object[attributeName] = make([]map[string]interface{}, 0)
					if reflect.TypeOf(child.GetObject()).String() == "map[string]interface {}" {
						this.object[attributeName] = append(this.object[attributeName].([]map[string]interface{}), child.GetObject().(map[string]interface{}))
					} else {
						object, _ := Utility.ToMap(child.GetObject())
						this.object[attributeName] = append(this.object[attributeName].([]map[string]interface{}), object)
					}
				}
			} else if reflect.TypeOf(this.object[attributeName]).String() == "[]interface {}" {
				if this.object[attributeName] != nil {
					objects := this.object[attributeName].([]interface{})
					objects_ := make([]interface{}, 0)
					isExist := false
					for i := 0; i < len(objects); i++ {
						if objects[i].(map[string]interface{})["UUID"].(string) != child.GetUuid() {
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
					this.object[attributeName] = objects_
				} else {
					this.object[attributeName] = make([]map[string]interface{}, 0)
					if reflect.TypeOf(child.GetObject()).String() == "map[string]interface {}" {
						this.object[attributeName] = append(this.object[attributeName].([]map[string]interface{}), child.GetObject().(map[string]interface{}))
					} else {
						object, _ := Utility.ToMap(child.GetObject())
						this.object[attributeName] = append(this.object[attributeName].([]map[string]interface{}), object)
					}
				}
			}

		} else {
			// Set or replace the value...
			if reflect.TypeOf(child.GetObject()).String() == "map[string]interface {}" {
				this.object[attributeName] = child.GetObject()
			} else {
				object, _ := Utility.ToMap(child.GetObject())
				this.object[attributeName] = object
			}
		}
	}

	// Set the child list...
	if Utility.Contains(this.childsUuid, child.GetUuid()) == false {
		// Append it to the list of UUID
		this.childsUuid = append(this.childsUuid, child.GetUuid())
		// Append it to the child
		this.childsPtr = append(this.childsPtr, child)
	} else {

		// In that case I will update the value inside the childsPtr...
		isExist := false
		childsPtr := make([]Entity, 0)
		for i := 0; i < len(this.childsPtr); i++ {
			if this.childsPtr[i].GetUuid() != child.GetUuid() {
				childsPtr = append(childsPtr, this.childsPtr[i])
			} else {
				// Replace the child ptr.
				childsPtr = append(childsPtr, child)
				isExist = true
			}
		}
		// Append the child ptr in that case...
		if !isExist {
			childsPtr = append(childsPtr, child)
			this.SetChildsPtr(childsPtr)
		}
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
	this.Lock()
	defer this.Unlock()

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
	log.Println("---------> ", this.object["UUID"])
	token := Utility.NewToken()
	this.lock.Lock(token)
	defer this.lock.Unlock(token)
	return this.object
}

/**
 * Return the object wrapped by this entity...
 */
func (this *DynamicEntity) SetObjectValues(values map[string]interface{}) {
	// Here I will clean up field inside the object that are not inside the new
	// value.
	token := Utility.NewToken()
	this.lock.Lock(token)
	defer this.Unlock()

	// nothing to do if the pointer is nil...
	if this == nil {
		return
	}

	entity, exist := GetServer().GetEntityManager().contain(this.GetUuid())
	if !exist {
		// If the map is not in the cache I can set the values directly.
		this.object = values
		this.SetNeedSave(true)
		return
	} else {
		// here I will set the need save attribute.
		entity.SetNeedSave(entity.GetChecksum() != Utility.GetChecksum(values))
		this.object = entity.GetObject().(map[string]interface{})
	}

	for i := 0; i < len(this.prototype.Fields); i++ {
		field := this.prototype.Fields[i]
		fieldType := this.prototype.FieldsType[i]
		isArray := strings.HasPrefix(fieldType, "[]")
		isRef := strings.HasSuffix(fieldType, ":Ref")
		isBaseType := strings.HasPrefix(fieldType, "[]xs.") || strings.HasPrefix(fieldType, "xs.") || strings.HasPrefix(fieldType, "sqltypes.") || strings.HasPrefix(fieldType, "[]sqltypes.")

		// The goal here is to remove base value or reference if there is no
		// more in the values.
		if this.object[field] != nil {
			if strings.HasPrefix(field, "M_") {
				// remove nil value
				if isBaseType {
					if this.object[field] != nil && values[field] == nil {
						this.object = nil
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
						if reflect.TypeOf(this.object[field]).String() == "[]string" {
							for i := 0; i < len(this.object[field].([]string)); i++ {
								if !Utility.Contains(v_, this.object[field].([]string)[i]) {
									toRemove := this.object[field].([]string)[i]
									ref, err := GetServer().GetEntityManager().getEntityByUuid(toRemove)
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
							ref, err := GetServer().GetEntityManager().getEntityByUuid(v_[i])
							if err == nil {
								ref.AppendReferenced(field, this)
								this.AppendReference(ref)
							}
						}

						// Set the ref values...
						this.object[field] = v_

					} else {
						var toRemove string

						if v == nil && this.object[field] != nil {
							toRemove = this.object[field].(string)
						} else if v != nil && this.object[field] != nil {
							if reflect.TypeOf(this.object[field]).String() == "map[string]interface {}" {
								if v.(map[string]interface{})["UUID"] != this.object[field] {
									toRemove = this.object[field].(string)
								}
							} else if reflect.TypeOf(this.object[field]).String() == "string" {
								if v != this.object[field] {
									toRemove = this.object[field].(string)
								}
							}
						}
						if len(toRemove) > 0 {
							ref, err := GetServer().GetEntityManager().getEntityByUuid(toRemove)
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
								ref, err := GetServer().GetEntityManager().getEntityByUuid(refUuid)
								if err == nil {
									ref.AppendReferenced(field, this)
									this.AppendReference(ref)
								}
								this.object[field] = refUuid
							}
						} else {
							this.object[field] = nil
						}
					}
				}
			}
		}
	}

	// Here the object exist so I need to copy values inside the existing map.
	for k, v := range values {

		if v == nil {
			delete(this.object, k)
		} else {
			// Only propertie with M_ will be set here.
			fieldIndex := this.prototype.getFieldIndex(k)
			if strings.HasPrefix(k, "M_") && fieldIndex != -1 {
				fieldType := this.prototype.FieldsType[fieldIndex]
				isRef := strings.HasSuffix(fieldType, ":Ref")
				isArray := strings.HasPrefix(fieldType, "[]")

				if !isRef {
					if this.object[k] == nil {
						// simply set the value...
						if v != nil {
							this.object[k] = v
						}
					} else {
						if reflect.TypeOf(this.object[k]).String() == "map[string]interface {}" {
							subValues := v.(map[string]interface{})
							if subValues["UUID"] != nil {
								// Here I have another dynamic object...
								subEntity, exist := GetServer().GetEntityManager().contain(subValues["UUID"].(string))
								if exist {
									// Set the sub entity
									subEntity.(*DynamicEntity).SetObjectValues(subValues)
								} else {
									var errObj *CargoEntities.Error
									subEntity, errObj = GetServer().GetEntityManager().newDynamicEntity(subValues)
									if errObj == nil {
										this.object[k] = subEntity.GetObject()
									}
								}
							} else {
								log.Println(Utility.FileLine(), " Map of interface without UUID found...!")
								this.object[k] = v
							}
						} else {
							if this.object[k] != nil {
								if reflect.TypeOf(this.object[k]).String() == "[]map[string]interface {}" {
									for i := 0; i < len(this.object[k].([]map[string]interface{})); i++ {
										subValues := this.object[k].([]map[string]interface{})[i]
										if subValues["UUID"] != nil {
											// Here I have another dynamic object...
											subEntity, exist := GetServer().GetEntityManager().contain(subValues["UUID"].(string))
											if exist {
												// Set the sub entity
												subEntity.(*DynamicEntity).SetObjectValues(subValues)
											} else {
												var errObj *CargoEntities.Error
												subEntity, errObj = GetServer().GetEntityManager().newDynamicEntity(subValues)
												if errObj == nil {
													this.object[k].([]map[string]interface{})[i] = subEntity.GetObject().(map[string]interface{})
												}
											}
										} else {
											log.Println(Utility.FileLine(), " Map of interface without UUID found...!")
											this.object[k].([]map[string]interface{})[i] = v.(map[string]interface{})
										}
									}
								} else if reflect.TypeOf(this.object[k]).String() == "[]interface {}" {
									if len(this.object[k].([]interface{})) == 0 {
										this.object[k] = v
									} else {
										if reflect.TypeOf(this.object[k].([]interface{})[0]).String() == "map[string]interface {}" {
											for i := 0; i < len(this.object[k].([]interface{})); i++ {
												subValues := this.object[k].([]interface{})[i].(map[string]interface{})
												if subValues["UUID"] != nil {
													// Here I have another dynamic object...
													subEntity, exist := GetServer().GetEntityManager().contain(subValues["UUID"].(string))
													if exist {
														// Set the sub entity
														subEntity.(*DynamicEntity).SetObjectValues(subValues)
													} else {
														var errObj *CargoEntities.Error
														subEntity, errObj = GetServer().GetEntityManager().newDynamicEntity(subValues)
														if errObj == nil {
															this.object[k].([]interface{})[i] = subEntity.GetObject()
														}
													}
												} else {
													log.Println(Utility.FileLine(), " Map of interface without UUID found...!")
													this.object[k].([]interface{})[i] = v
												}
											}
										} else {
											// Replace the array with the new value.
											this.object[k] = v
										}
									}
								} else if reflect.TypeOf(this.object[k]).String() == "[]string" {
									this.object[k] = v // Replace the array of string with the new value.
								} else if this.object[k] != v {
									if reflect.TypeOf(this.object[k]).Kind() == reflect.TypeOf(v).Kind() {
										this.object[k] = v
									} else {
										// Here I need to convert the v type...
										if reflect.TypeOf(this.object[k]).Kind() == reflect.Float64 && reflect.TypeOf(v).Kind() == reflect.String {
											this.object[k], _ = strconv.ParseFloat(v.(string), 64)
										} else if reflect.TypeOf(this.object[k]).Kind() == reflect.Float32 && reflect.TypeOf(v).Kind() == reflect.String {
											this.object[k], _ = strconv.ParseFloat(v.(string), 32)
										} else if reflect.TypeOf(this.object[k]).Kind() == reflect.Int64 && reflect.TypeOf(v).Kind() == reflect.String {
											this.object[k], _ = strconv.ParseInt(v.(string), 10, 64)
										} else if reflect.TypeOf(this.object[k]).Kind() == reflect.Int8 && reflect.TypeOf(v).Kind() == reflect.String {
											this.object[k], _ = strconv.ParseInt(v.(string), 10, 8)
										} else if reflect.TypeOf(this.object[k]).Kind() == reflect.Int && reflect.TypeOf(v).Kind() == reflect.String {
											this.object[k], _ = strconv.ParseInt(v.(string), 10, 64)
										} else if reflect.TypeOf(this.object[k]).Kind() == reflect.Int32 && reflect.TypeOf(v).Kind() == reflect.String {
											this.object[k], _ = strconv.ParseInt(v.(string), 10, 32)
										} else if reflect.TypeOf(this.object[k]).Kind() == reflect.Int16 && reflect.TypeOf(v).Kind() == reflect.String {
											this.object[k], _ = strconv.ParseInt(v.(string), 10, 16)
										} else if reflect.TypeOf(this.object[k]).Kind() == reflect.Bool && reflect.TypeOf(v).Kind() == reflect.String {
											this.object[k], _ = strconv.ParseBool(v.(string))
										} else if reflect.TypeOf(this.object[k]).Kind() == reflect.Int64 && reflect.TypeOf(v).Kind() == reflect.Float64 {
											// the json parser transform all numerical value to float... that not what we want here...
											this.object[k] = int64(v.(float64))
										}
									}

								}
							} else {
								// Append the new value here...
								this.object[k] = v // Replace the array of string with the new value.
							}
						}
					}
				} else {
					// Here I have a reference. I need to keep only the uuid
					// in case of I receive the complete objet value.
					if isArray {
						if this.object[k] == nil {
							this.object[k] = make([]string, 0)
						}
						if reflect.TypeOf(v).String() == "[]map[string] interface{}" {
							for i := 0; i < len(v.([]map[string]interface{})); i++ {
								this.object[k] = append(this.object[k].([]string), v.(map[string]interface{})["UUID"].(string))
							}
						} else if reflect.TypeOf(v).String() == "[]interface {}" {
							for i := 0; i < len(v.([]interface{})); i++ {
								if reflect.TypeOf(v.([]interface{})[i]).String() == "map[string] interface{}" {
									this.object[k] = append(this.object[k].([]string), v.([]interface{})[i].(map[string]interface{})["UUID"].(string))
								} else if reflect.TypeOf(v.([]interface{})[i]).String() == "string" {
									this.object[k] = append(this.object[k].([]string), v.([]interface{})[i].(string))
								}
							}
						} else if reflect.TypeOf(v).String() == "[]string" {
							// and array of string.
							this.object[k] = v
						}
					} else {
						if reflect.TypeOf(v).String() == "map[string]interface {}" {
							this.object[k] = v.(map[string]interface{})["UUID"]
						} else if reflect.TypeOf(v).String() == "string" {
							this.object[k] = v
						}
					}
				}
			} else {
				if strings.HasPrefix(k, "M_") {
					log.Println(this.prototype.TypeName, " has no field ", k)
				}
			}
		}
	}
}

/**
 * Return the entity prototype.
 */
func (this *DynamicEntity) GetPrototype() *EntityPrototype {
	return this.prototype
}

/**
 * Calculate a unique value for a given entity object value...
 */
func (this *DynamicEntity) GetChecksum() string {
	return Utility.GetChecksum(this.object)
}

/**
 * Check if the an entity exist...
 */
func (this *DynamicEntity) Exist() bool {
	// if the 'this' pointer is not initialyse it means the entity was deleted.
	if this == nil || this.object["TYPENAME"] == nil {
		return false
	}

	var query EntityQuery
	query.TypeName = this.GetTypeName()
	query.Indexs = append(query.Indexs, "uuid="+this.uuid)
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	storeId := this.GetPackageName()
	if reflect.TypeOf(dataManager.getDataStore(storeId)).String() == "*Server.SqlDataStore" {
		storeId = "sql_info" // Must save or update value from sql info instead.
	}
	results, err := GetServer().GetDataManager().readData(storeId, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}

	return len(results[0][0].(string)) > 0
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
	s := string(b)
	return s
}
