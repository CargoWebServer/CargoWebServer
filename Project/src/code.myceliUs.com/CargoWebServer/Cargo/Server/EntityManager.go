package Server

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

////////////////////////////////////////////////////////////////////////////////
//						Entity
////////////////////////////////////////////////////////////////////////////////

/**
 * Struct used to keep information about references.
 */
type EntityRef struct {
	Name      string
	OwnerUuid string
	Value     string
}

////////////////////////////////////////////////////////////////////////////////
//						Entity Manager
////////////////////////////////////////////////////////////////////////////////
type EntityManager struct {

	/**
	 * ref -> entity
	 */
	referenced map[string][]EntityRef

	/**
	 * entity -> ref
	 */
	reference map[string][]EntityRef

	/**
	 * The root of all object...
	 */
	cargoEntities *CargoEntities_EntitiesEntity

	/**
	* Lock referenced and references map.
	 */
	sync.Mutex
}

var entityManager *EntityManager

func (this *Server) GetEntityManager() *EntityManager {
	if entityManager == nil {
		entityManager = newEntityManager()
	}
	return entityManager
}

func newEntityManager() *EntityManager {

	entityManager = new(EntityManager)

	// Create prototypes for config objects and entities objects...
	entityManager.createConfigPrototypes()
	entityManager.createCargoEntitiesPrototypes()
	entityManager.registerConfigObjects()
	entityManager.registerCargoEntitiesObjects()

	// Entity prototype is a dynamic type.
	Utility.RegisterType((*EntityPrototype)(nil))
	Utility.RegisterType((*Restriction)(nil))
	Utility.RegisterType((*DynamicEntity)(nil))
	Utility.RegisterType((*MessageData)(nil))

	// References
	entityManager.referenced = make(map[string][]EntityRef, 0)
	entityManager.reference = make(map[string][]EntityRef, 0)

	return entityManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Initialization.
 */
func (this *EntityManager) initialize() {
	log.Println("--> Initialize EntityManager")

	// Create the default configurations
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())

	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})

	cargoEntitiesUuid := CargoEntitiesEntitiesExists("CARGO_ENTITIES")
	if len(cargoEntitiesUuid) > 0 {
		this.cargoEntities = this.NewCargoEntitiesEntitiesEntity(cargoEntitiesUuid, nil)
		this.cargoEntities.InitEntity(cargoEntitiesUuid)
	} else {
		this.cargoEntities = this.NewCargoEntitiesEntitiesEntity("CARGO_ENTITIES", nil)
		this.cargoEntities.object.M_id = "CARGO_ENTITIES"
		this.cargoEntities.object.M_name = "Cargo entities"
		this.cargoEntities.object.M_version = "1.0"
		this.cargoEntities.object.NeedSave = true
		this.cargoEntities.SaveEntity()
	}

}

func (this *EntityManager) getId() string {
	return "EntityManager"
}

func (this *EntityManager) start() {
	log.Println("--> Start EntityManager")
}

func (this *EntityManager) stop() {
	log.Println("--> Stop EntityManager")
}

/**
 * Cargo entities contains files, accounts, etc...
 */
func (this *EntityManager) getCargoEntities() *CargoEntities_EntitiesEntity {

	return this.cargoEntities

}

/**
 * Wrapper to the cacheManager's 'contains()' function.
 * Determines if the entity exists in the cacheManager map.
 */
func (this *EntityManager) contain(uuid string) (Entity, bool) {
	return server.GetCacheManager().contains(uuid)
}

/**
 * Wrapper to the cacheManager's 'setEntity()' function.
 * Inserts entity in the cacheManager's map if it doesn't already exist.
 * Otherwise replaces the entity in the map with this entity.
 */
func (this *EntityManager) insert(entity Entity) {
	// Set the cache...
	server.GetCacheManager().setEntity(entity)
}

/**
 * Wrapper to the cacheManager's 'removeEntity()' function.
 * Removes an existing entity with a given uuid from the cacheManager's map.
 */
func (this *EntityManager) removeEntity(uuid string) {
	server.GetCacheManager().removeEntity(uuid)
}

/**
 * That function is use to delete an entity from the store.
 */
func (this *EntityManager) deleteEntity(toDelete Entity) {
	// first of all i will remove it from the cache.
	this.removeEntity(toDelete.GetUuid())

	// remove it's data from the database.
	var deleteEntityQuery EntityQuery
	deleteEntityQuery.TypeName = toDelete.GetTypeName()
	deleteEntityQuery.Indexs = append(deleteEntityQuery.Indexs, "uuid="+toDelete.GetUuid())
	var params []interface{}
	query, _ := json.Marshal(deleteEntityQuery)

	GetServer().GetDataManager().deleteData(toDelete.GetTypeName()[0:strings.Index(toDelete.GetTypeName(), ".")], string(query), params)

	// delete it's childs.
	for i := 0; i < len(toDelete.GetChildsPtr()); i++ {
		toDelete.GetChildsPtr()[i].DeleteEntity()
	}

	// This variable will keep track of other entity to save after that entity will be
	// deleted.
	toSaves := make([]Entity, 0)

	// Now it's references/childs in it's owner/parent...
	for i := 0; i < len(toDelete.GetReferenced()); i++ {
		ref := toDelete.GetReferenced()[i]
		if this.isExist(ref.OwnerUuid) {
			refOwner, err := this.getEntityByUuid(ref.OwnerUuid)
			if err == nil {
				prototype := refOwner.GetPrototype()

				fieldIndex := prototype.getFieldIndex(ref.Name)
				if !strings.HasPrefix(ref.Name, "M_") && fieldIndex == -1 {
					fieldIndex = prototype.getFieldIndex("M_" + ref.Name)
				}

				fieldType := prototype.FieldsType[fieldIndex]

				isRef := strings.HasSuffix(fieldType, ":Ref")
				if isRef {
					refOwner.RemoveReference(ref.Name, toDelete)
				} else {
					refOwner.RemoveChild(ref.Name, toDelete.GetUuid())
				}
			}

			// append save only once.
			isExist := false
			for i := 0; i < len(toSaves); i++ {
				if toSaves[i].GetUuid() == refOwner.GetUuid() {
					isExist = true
				}
			}
			if !isExist {
				toSaves = append(toSaves, refOwner)
			}
		}
	}

	// Save refeferenced entity...
	for i := 0; i < len(toSaves); i++ {
		// Save it only if it dosen't already deleted.
		toSaves[i].SetNeedSave(true)
		toSaves[i].SaveEntity()
	}

	// Send event message...
	var eventDatas []*MessageData
	evtData := new(MessageData)
	evtData.Name = "entity"

	evtData.Value = toDelete.GetObject()
	eventDatas = append(eventDatas, evtData)
	evt, _ := NewEvent(DeleteEntityEvent, EntityEvent, eventDatas)
	GetServer().GetEventManager().BroadcastEvent(evt)

	log.Println("----------> entity ", toDelete.GetUuid(), " is remove ", !this.isExist(toDelete.GetUuid()))
}

/**
 * Set the list of reference of a given owner entity.
 */
func (this *EntityManager) setReferences(owner Entity) {
	prototype := owner.GetPrototype()
	for i := 0; i < len(prototype.FieldsType); i++ {
		fieldType := prototype.FieldsType[i]
		if strings.HasSuffix(fieldType, ":Ref") {
			fieldName := prototype.Fields[i]
			uuids := make([]string, 0)
			if reflect.TypeOf(owner.GetObject()).String() == "map[string]interface {}" {
				// Dynamic entity here.
				objectMap := owner.GetObject().(map[string]interface{})
				if objectMap[fieldName] != nil {
					if reflect.TypeOf(objectMap[fieldName]).String() == "string" {
						uuids = append(uuids, objectMap[fieldName].(string))
					} else if reflect.TypeOf(objectMap[fieldName]).String() == "[]interface {}" {
						for j := 0; j < len(objectMap[fieldName].([]interface{})); j++ {
							fieldValue := objectMap[fieldName].([]interface{})[j]
							if reflect.TypeOf(objectMap[fieldName]).String() == "string" {
								uuids = append(uuids, fieldValue.(string))
							}
						}
					} else if reflect.TypeOf(objectMap[fieldName]).String() == "[]string" {
						for j := 0; j < len(objectMap[fieldName].([]string)); j++ {
							uuids = append(uuids, objectMap[fieldName].([]string)[j])
						}
					}
				}
			} else {
				// In case of static entity...
				ps := reflect.ValueOf(owner.GetObject())
				s := ps.Elem()

				if s.Kind() == reflect.Struct {

					if strings.HasPrefix(fieldName, "M_") {
						fieldName = "m_" + fieldName[2:]
					}

					f := s.FieldByName(fieldName)
					if f.IsValid() {
						if f.Kind() == reflect.String {
							uuid := f.String()
							if len(uuid) > 0 {
								uuids = append(uuids, uuid)
							}
						} else if f.Kind() == reflect.Slice {
							for j := 0; j < f.Len(); j++ {
								if f.Index(j).Kind() == reflect.String {
									uuid := f.Index(j).String()
									if len(uuid) > 0 {
										uuids = append(uuids, uuid)
									}
								} else {
									if f.Index(j).Kind() == reflect.Ptr {
										// Here I have a structure.
										uuid := f.Index(j).Elem().FieldByName("UUID").String()
										if len(uuid) > 0 {
											uuids = append(uuids, uuid)
										}

									}
								}
							}
						} else if f.Kind() == reflect.Ptr {
							// Here I have a structure.
							if !f.IsNil() {
								uuid := f.Elem().FieldByName("UUID").String()
								if len(uuid) > 0 {
									uuids = append(uuids, uuid)
								}
							}
						}
					}
				}
			}

			// Now I will try to append the reference inside the
			// entity.
			for j := 0; j < len(uuids); j++ {
				reference, err := this.getEntityByUuid(uuids[j])
				if err == nil {
					owner.AppendReference(reference)
				}
			}
		}
	}

}

/**
 * Set the content of a target object whit the source object. use by static entity
 * only...
 */
func (this *EntityManager) setObjectValues(target Entity, source interface{}) {

	// in case of dynamic object...
	if reflect.TypeOf(source).String() == "map[string]interface {}" {
		target.(*DynamicEntity).SetObjectValues(source.(map[string]interface{}))
	}

	// here we have a static object...
	prototype := target.GetPrototype()

	// The need save evaluation...
	mapValues, _ := Utility.ToMap(source)
	needSave := target.GetChecksum() != Utility.GetChecksum(mapValues)
	target.SetNeedSave(needSave)

	// I will get the target object.
	targetReflexObject := reflect.ValueOf(target.GetObject())
	sourceFelfexObject := reflect.ValueOf(source)

	// First of all I will reset the values that are not in the target but are
	// in the source.
	for i := 0; i < len(prototype.FieldsType); i++ {
		fieldName := prototype.Fields[i]
		fieldType := prototype.FieldsType[i]
		isArray := strings.HasPrefix(fieldType, "[]")
		isRef := strings.HasSuffix(fieldType, ":Ref")
		isBaseType := strings.HasPrefix(fieldType, "[]xs.") || strings.HasPrefix(fieldType, "xs.")
		isEnum := strings.HasPrefix(fieldType, "enum")

		var targetField reflect.Value
		if targetReflexObject.Kind() == reflect.Ptr {
			targetField = targetReflexObject.Elem().FieldByName(fieldName)
		} else {
			targetField = targetReflexObject.FieldByName(fieldName)
		}

		var sourceField reflect.Value
		if targetReflexObject.Kind() == reflect.Ptr {
			sourceField = sourceFelfexObject.Elem().FieldByName(fieldName)
		} else {
			sourceField = sourceFelfexObject.FieldByName(fieldName)
		}

		if targetField.IsValid() {
			if isBaseType || isEnum {
				if !sourceField.IsValid() {
					if isArray {
						targetField.Set(reflect.ValueOf(make([]interface{}, 0)))
					} else {
						var val interface{}
						targetField.Set(reflect.ValueOf(val))
					}
				}
			} else if isRef {
				var removeMethode = "Remove" + strings.ToUpper(fieldName[2:3]) + fieldName[3:]
				var setMethode = "Set" + strings.ToUpper(fieldName[2:3]) + fieldName[3:]
				if isArray {
					for i := 0; i < targetField.Len(); i++ {
						ref := targetField.Index(i)
						needToBeRemove := true
						for j := 0; j < sourceField.Len(); j++ {
							if sourceField.Index(j) == ref {
								needToBeRemove = false
								break
							}
						}
						if needToBeRemove && len(ref.String()) > 0 {
							toRemove, _ := this.getEntityByUuid(ref.String())
							target.RemoveReference(fieldName, toRemove)
							toRemove.RemoveReferenced(fieldName, target)
							// I will call remove function...
							params := make([]interface{}, 1)
							params[0] = toRemove.GetObject()
							Utility.CallMethod(target.GetObject(), removeMethode, params)
						}
					}
					// Append the references...
					for i := 0; i < sourceField.Len(); i++ {
						if len(sourceField.Index(i).String()) > 0 {
							ref, err := this.getEntityByUuid(sourceField.Index(i).String())
							if err == nil {
								target.AppendReference(ref)
								ref.AppendReferenced(fieldName, target)
								// Now I will call the append method...
								params := make([]interface{}, 1)
								params[0] = ref.GetObject()
								Utility.CallMethod(target.GetObject(), setMethode, params)
							} else {
								params := make([]interface{}, 1)
								params[0] = sourceField.Index(i).String()
								Utility.CallMethod(target.GetObject(), setMethode, params)
							}
						}
					}
				} else {
					// Remove the reference...
					if sourceField.String() != targetField.String() {
						if len(targetField.String()) > 0 {
							toRemove, err := this.getEntityByUuid(targetField.String())
							if err == nil {
								target.RemoveReference(fieldName, toRemove)
								toRemove.RemoveReferenced(fieldName, target)
								// I will call remove function...
								params := make([]interface{}, 1)
								params[0] = toRemove.GetObject()
								Utility.CallMethod(target.GetObject(), removeMethode, params)
							}
						}
					}
					// Append the reference...
					if len(sourceField.String()) > 0 {
						ref, err := this.getEntityByUuid(sourceField.String())
						if err == nil {
							target.AppendReference(ref)
							ref.AppendReferenced(fieldName, target)
							// Now I will call the append method...
							params := make([]interface{}, 1)
							params[0] = ref.GetObject()
							Utility.CallMethod(target.GetObject(), setMethode, params)
						} else {
							params := make([]interface{}, 1)
							params[0] = sourceField.String()
							Utility.CallMethod(target.GetObject(), setMethode, params)
						}
					}
				}

			}
		}
	}

	// set the new field values.
	for i := 0; i < len(prototype.FieldsType); i++ {
		fieldName := prototype.Fields[i]
		fieldType := prototype.FieldsType[i]
		isArray := strings.HasPrefix(fieldType, "[]")
		isRef := strings.HasSuffix(fieldType, ":Ref")
		isBaseType := strings.HasPrefix(fieldType, "[]xs.") || strings.HasPrefix(fieldType, "xs.")
		isEnum := strings.HasPrefix(fieldType, "enum")

		var targetField reflect.Value
		if targetReflexObject.Kind() == reflect.Ptr {
			targetField = targetReflexObject.Elem().FieldByName(fieldName)
		} else {
			targetField = targetReflexObject.FieldByName(fieldName)
		}

		var sourceField reflect.Value
		if targetReflexObject.Kind() == reflect.Ptr {
			sourceField = sourceFelfexObject.Elem().FieldByName(fieldName)
		} else {
			sourceField = sourceFelfexObject.FieldByName(fieldName)
		}

		if sourceField.IsValid() {
			if targetField.IsValid() {

				if isBaseType || isEnum {
					// set the value(s)...
					targetField.Set(sourceField)
				} else if !isRef {
					var removeMethode = "Remove" + strings.ToUpper(fieldName[2:3]) + fieldName[3:]
					if isArray {
						// First I will remove the object in the target that are no more in the
						// source.
						for i := 0; i < targetField.Len(); i++ {
							toRemove := targetField.Index(i)
							needToBeRemove := true
							for j := 0; j < sourceField.Len(); j++ {
								subObject := sourceField.Index(j)
								var params []interface{}
								var subObjectUuid string
								val0, err0 := Utility.CallMethod(subObject.Interface(), "GetUUID", params)
								if err0 == nil {
									subObjectUuid = val0.(string)
								}

								var toRemoveUuid string
								val1, err1 := Utility.CallMethod(toRemove.Interface(), "GetUUID", params)
								if err1 == nil {
									toRemoveUuid = val1.(string)
								}

								if toRemoveUuid == subObjectUuid {
									needToBeRemove = false
									break
								}
							}
							if needToBeRemove {
								params := make([]interface{}, 1)
								params[0] = toRemove.Interface()
								Utility.CallMethod(target.GetObject(), removeMethode, params)
							}
						}

						// Now I will set the field...
						for j := 0; j < sourceField.Len(); j++ {
							subObject := sourceField.Index(j)
							if subObject.Interface() != nil {
								var params []interface{}
								val, err := Utility.CallMethod(subObject.Interface(), "GetUUID", params)
								if err == nil {
									subObjectUuid := val.(string)
									subEntity, err := this.getEntityByUuid(subObjectUuid)
									if err == nil {
										this.setObjectValues(subEntity, subObject.Interface())
									}
									setMethodName := strings.Replace(fieldName, "M_", "", -1)
									setMethodName = "Set" + strings.ToUpper(setMethodName[0:1]) + setMethodName[1:]
									params := make([]interface{}, 1)
									params[0] = subObject.Interface()
									_, err_ := Utility.CallMethod(target.GetObject(), setMethodName, params)
									if err_ != nil {
										log.Println("fail to call method ", setMethodName, " on ", target.GetObject())
									}
								} else {
									log.Println("----------> fail to call method GetUUID on ", subObject)
								}
							} else {
								// TODO remove the value here.
								// Remove the value here...
							}
						}
					} else {
						// Clear the actual value...
						if !sourceField.IsNil() {
							// remove the existing value

							// Set the new value.
							var subObjectUuid string
							var params []interface{}
							if sourceField.Interface() != nil {
								val, err := Utility.CallMethod(sourceField.Interface(), "GetUUID", params)
								if err == nil {
									subObjectUuid = val.(string)
								} else {
									log.Println("----------> fail to call method GetUUID on ", sourceField.Interface())
								}

								subEntity, err := this.getEntityByUuid(subObjectUuid)
								if err == nil {
									this.setObjectValues(subEntity, sourceField.Interface())
									setMethodName := strings.Replace(fieldName, "M_", "", -1)
									setMethodName = "Set" + strings.ToUpper(setMethodName[0:1]) + setMethodName[1:]
									params := make([]interface{}, 1)
									params[0] = sourceField.Interface()
									_, err := Utility.CallMethod(target.GetObject(), setMethodName, params)
									if err != nil {
										log.Println("fail to call method ", setMethodName, " on ", target.GetObject())
									}
								}
							}
						}
					}
				}
			}
		}
	}

}

/**
 * Return an entity with for a given type and id
 */
func (this *EntityManager) getEntityById(storeId string, typeName string, id string) (Entity, *CargoEntities.Error) {
	// Verify that typeName is valid
	// interface{} is an exception...
	if !Utility.IsValidPackageName(typeName) && !strings.HasSuffix(typeName, "interface{}") {
		cargoError := NewError(Utility.FileLine(), INVALID_PACKAGE_NAME_ERROR, SERVER_ERROR_CODE, errors.New("Type name '"+typeName+"' is not valid."))
		return nil, cargoError
	}

	// If the store is not found I will return an error.
	if GetServer().GetDataManager().getDataStore(storeId) == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Datastore '"+storeId+"' dosen't exist."))
		return nil, cargoError
	}

	// The entity information are in sql_info and not the given store id...
	if reflect.TypeOf(GetServer().GetDataManager().getDataStore(storeId)).String() == "*Server.SqlDataStore" {
		// I will try to found the prototype inside sql_info instead.
		storeId = "sql_info"
	}

	prototype, err := this.getEntityPrototype(typeName, storeId)

	var results [][]interface{}
	if err == nil {
		var query EntityQuery
		query.TypeName = typeName
		query.Fields = append(query.Fields, "uuid")
		var fieldsType []interface{} // not used
		var params []interface{}
		var ids []string
		ids = append(ids, prototype.Ids...)
		for i := 1; i < len(ids) && len(results) == 0; i++ {
			idField := ids[i]
			query.Indexs = make([]string, 0)
			query.Indexs = append(query.Indexs, idField+"="+id)
			queryStr, _ := json.Marshal(query)

			results, err = GetServer().GetDataManager().readData(storeId, string(queryStr), fieldsType, params)
			if err != nil {
				// Create the error message
				cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err)
				return nil, cargoError
			}
		}
	}

	// In that case not information are found.
	if len(results) == 0 {
		// Here I will send a an error...
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("No values found for type '"+typeName+"' and id '"+id+"'"))
		return nil, cargoError
	}

	entity, errObj := this.getEntityByUuid(results[0][0].(string))

	if errObj != nil || entity == nil {
		entity, errObj = this.getDynamicEntityByUuid(results[0][0].(string))
		if errObj != nil {
			return nil, errObj
		}
	}

	return entity, nil
}

/**
 * Return the list of entity type derived from a given type.
 */
func (this *EntityManager) getDerivedEntityType(typeName string) ([]*EntityPrototype, *CargoEntities.Error) {
	var derived []*EntityPrototype

	if !Utility.IsValidPackageName(typeName) {
		cargoError := NewError(Utility.FileLine(), INVALID_PACKAGE_NAME_ERROR, SERVER_ERROR_CODE, errors.New("Type name '"+typeName+"' is not valid."))
		return derived, cargoError
	}

	packageName := typeName[0:strings.Index(typeName, ".")]
	// Here I will retreive the supertype
	superTypePrototype, err := this.getEntityPrototype(typeName, packageName)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), PROTOTYPE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Type name '"+typeName+"' dosen't exist."))
		return derived, cargoError
	}

	for i := 0; i < len(superTypePrototype.SubstitutionGroup); i++ {
		packageName := superTypePrototype.SubstitutionGroup[i][0:strings.Index(superTypePrototype.SubstitutionGroup[i], ".")]
		substitutionGroup, err := this.getEntityPrototype(superTypePrototype.SubstitutionGroup[i], packageName)
		if err == nil {
			derived = append(derived, substitutionGroup)
		}
	}

	return derived, nil
}

/**
 * Return the list of entities for a given type name.
 */
func (this *EntityManager) getEntitiesByType(typeName string, queryStr string, storeId string) ([]Entity, *CargoEntities.Error) {

	var entities []Entity

	if GetServer().GetDataManager().getDataStore(storeId) == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Datastore '"+storeId+"' dosen't exist."))
		return nil, cargoError
	}
	var dataStore DataStore
	dataStore = GetServer().GetDataManager().getDataStore(storeId)
	if reflect.TypeOf(dataStore).String() == "*Server.SqlDataStore" {
		dataStore = GetServer().GetDataManager().getDataStore("sql_info")
	}

	if len(queryStr) == 0 {
		values, err := dataStore.(*KeyValueDataStore).getIndexation(typeName)
		if err != nil {
			return entities, NewError(Utility.FileLine(), DATASTORE_INDEXATION_ERROR, SERVER_ERROR_CODE, errors.New("No indexation for type '"+typeName+"'."))
		}

		for i := 0; i < len(values); i++ {
			key := values[i].(string)
			values_, err := dataStore.(*KeyValueDataStore).getValues(key)
			if err != nil {
				return entities, NewError(Utility.FileLine(), DATASTORE_KEY_NOT_FOUND_ERROR, SERVER_ERROR_CODE, errors.New("No value found for key '"+key+"'."))
			}
			if len(values_) > 0 {
				uuid := values_[0].(string)
				entity, errObj := this.getEntityByUuid(uuid)
				if errObj != nil {
					entity, errObj = this.getDynamicEntityByUuid(uuid)
					if errObj != nil {
						return entities, errObj
					}
				}
				if entity != nil {
					entities = append(entities, entity)
				}
			}
		}
	} else {
		// Here I will create a new query and execute it...
		var query EntityQuery
		query.TypeName = typeName
		query.Query = queryStr

		// I will retreive the uuid...
		query.TypeName = typeName
		query.Fields = append(query.Fields, "uuid")
		var fieldsType []interface{} // not used
		var params []interface{}

		// Now I will execute the query...
		queryStr_, _ := json.Marshal(query)

		results, err := GetServer().GetDataManager().readData(storeId, string(queryStr_), fieldsType, params)
		if err != nil {
			// Create the error message
			cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err)
			return nil, cargoError
		}

		if len(results) > 0 {
			for i := 0; i < len(results); i++ {
				uuid := results[i][0].(string)
				entity, _ := this.getEntityByUuid(uuid)
				entities = append(entities, entity)
			}
		}
	}

	return entities, nil
}

/**
 * Return the list of all entities for different types.
 */
func (this *EntityManager) getEntitiesByTypes(typeNames []string, storeId string) ([]Entity, *CargoEntities.Error) {
	entitiesMap := make(map[string]Entity)
	var entities []Entity

	for i := 0; i < len(typeNames); i++ {
		entities_, errObj := this.getEntitiesByType(typeNames[i], "", storeId)
		if errObj != nil {
			return entities, errObj
		}
		for j := 0; j < len(entities_); j++ {
			entitiesMap[entities_[j].GetUuid()] = entities_[j]
		}
	}

	for _, entity := range entitiesMap {
		entities = append(entities, entity)
	}

	return entities, nil
}

/**
 * Return the list of all entity links to one entity. If we see an entity as graph it return
 * all nodes links to a given node.
 */
func (this *EntityManager) getEntityLnkLst(entity Entity, visited *[]string, lnkLst *[]Entity) {

	if len(*visited) > 0 {
		if Utility.Contains(*visited, entity.GetUuid()) {
			return // nothing to do here...
		}
	}

	// append the entity...
	if entity.Exist() {
		*visited = append(*visited, entity.GetUuid())
		*lnkLst = append(*lnkLst, entity)

		// Append childs...
		for _, child := range entity.GetChildsPtr() {
			this.getEntityLnkLst(child, visited, lnkLst)
		}

		for _, ref := range entity.GetReferencesPtr() {
			this.getEntityLnkLst(ref, visited, lnkLst)
		}
	} else {
		// delete the entity...
		//entity.DeleteEntity()
	}
}

func (this *EntityManager) getEntityByUuid(uuid string) (Entity, *CargoEntities.Error) {
	if !Utility.IsValidEntityReferenceName(uuid) {
		return nil, NewError(Utility.FileLine(), INVALID_REFERENCE_NAME_ERROR, SERVER_ERROR_CODE, errors.New("The uuid '"+uuid+"' is not valid."))
	}

	if val, ok := this.contain(uuid); ok {
		return val, nil
	}

	typeName := strings.Split(uuid, "%")[0]

	// Remove the suffix in that particular case.
	if strings.HasSuffix(typeName, "_impl") {
		typeName = strings.Replace(typeName, "_impl", "", -1)
	}

	funcName := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"

	params := make([]interface{}, 2)
	params[0] = uuid
	params[1] = nil
	result, err := Utility.CallMethod(this, funcName, params)

	if err != nil {
		// Try with dynamic entity instead.
		entity, errObj := this.getDynamicEntityByUuid(uuid)
		if errObj != nil {
			return nil, errObj
		}
		return entity, nil
	}

	entity := result.(Entity)
	entity.InitEntity(uuid)

	// Here I will also set the reference for the entity...
	this.setReferences(entity)

	return entity, nil

}

func (this *EntityManager) getDynamicEntityByUuid(uuid string) (Entity, *CargoEntities.Error) {

	if val, ok := this.contain(uuid); ok {
		return val, nil
	}

	values := make(map[string]interface{}, 0)
	values["TYPENAME"] = strings.Split(uuid, "%")[0]
	values["UUID"] = uuid

	entity, errObj := this.newDynamicEntity(values)

	if errObj != nil {
		return nil, errObj
	}

	if !entity.Exist() {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ENTITY_UUID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The entity uuid '"+uuid+"' does not correspond to an existing entity."))
		// Return the error entity in the err return param.
		return nil, cargoError
	}

	// initialyse the entity.
	entity.InitEntity(uuid)

	return entity, nil

}

// Use at initialysation time (init)
func (this *EntityManager) appendReference(name string, ownerId string, value string) {

	if value != "null" && !strings.HasSuffix(value, "$$null") {
		var ref EntityRef
		ref.Name = name
		ref.OwnerUuid = ownerId
		ref.Value = value

		this.Lock()

		if this.reference[ownerId] == nil {
			this.reference[ownerId] = make([]EntityRef, 0)
		}
		this.reference[ownerId] = append(this.reference[ownerId], ref)

		this.Unlock()

		owner, err := this.getEntityByUuid(ownerId)
		if err == nil {
			targetId := strings.Split(value, "$$")[1]
			if Utility.IsValidEntityReferenceName(targetId) {
				target, err := this.getEntityByUuid(targetId)
				if err == nil {
					// Set the reference...
					owner.AppendReference(target)
				}
			}
		}

	}
}

// Use at creation time by dynamic entity... (Save)
func (this *EntityManager) appendReferenced(name string, ownerId string, value string) {
	//this.Lock()
	//defer this.Unlock()

	if value != "null" && !strings.HasSuffix(value, "$$null") {
		var ref EntityRef
		ref.Name = name
		ref.OwnerUuid = ownerId
		ref.Value = value

		// here the target id must be use to resolve references...
		targetId := value[strings.Index(value, "$$")+2:]

		// Here I will try to find the targeted object...
		if Utility.IsValidEntityReferenceName(targetId) {
			// The key is an uuid
			targetRef, err := GetServer().GetEntityManager().getEntityByUuid(targetId)
			if err == nil {
				ownerRef, _ := GetServer().GetEntityManager().getEntityByUuid(ownerId)
				ownerRef.AppendReference(targetRef)
				// Append the reference here...
				targetRef.AppendReferenced(name, ownerRef)
				//log.Println("-------> append referenced: ", targetId, ":", ownerId)
				// Save the target entity...
				targetRef.SaveEntity()
				return // nothing to do...
			}
		}

		// Here the target object does not exist so when it will be created and
		// save the reference to this owner object will be added to the target.
		if this.referenced[targetId] == nil {
			this.referenced[targetId] = make([]EntityRef, 0)
		}

		// Append the referenced here...
		this.referenced[targetId] = append(this.referenced[targetId], ref)
	}
}

// used by dynamic entity only...
func (this *EntityManager) saveReferenced(entity Entity) {
	// Here I will try to find if there is pending reference for the object...
	typeName := entity.GetTypeName()
	packageName := typeName[0:strings.Index(typeName, ".")]
	prototype, err := this.getEntityPrototype(typeName, packageName)
	if err == nil {
		for i := 0; i < len(prototype.Ids); i++ {
			var referenced []EntityRef
			if i == 0 {
				// Get reference by uuid...
				referenced = this.referenced[entity.GetUuid()]
			} else {
				// Reference registered by other id...
				id := prototype.Ids[i]

				if reflect.TypeOf(entity.GetObject()).String() == "map[string]interface {}" {
					// A dynamic entity here...
					if entity.GetObject().(map[string]interface{})[id] != nil {
						var refId string
						if reflect.TypeOf(entity.GetObject().(map[string]interface{})[id]).Kind() == reflect.String {
							refId += entity.GetObject().(map[string]interface{})[id].(string)
						} else if reflect.TypeOf(entity.GetObject().(map[string]interface{})[id]).Kind() == reflect.Int {
							refId += strconv.Itoa(entity.GetObject().(map[string]interface{})[id].(int))
						} else if reflect.TypeOf(entity.GetObject().(map[string]interface{})[id]).Kind() == reflect.Int8 {
							refId += strconv.Itoa(int(entity.GetObject().(map[string]interface{})[id].(int8)))
						} else if reflect.TypeOf(entity.GetObject().(map[string]interface{})[id]).Kind() == reflect.Int16 {
							refId += strconv.Itoa(int(entity.GetObject().(map[string]interface{})[id].(int16)))
						} else if reflect.TypeOf(entity.GetObject().(map[string]interface{})[id]).Kind() == reflect.Int32 {
							refId += strconv.Itoa(int(entity.GetObject().(map[string]interface{})[id].(int32)))
						} else if reflect.TypeOf(entity.GetObject().(map[string]interface{})[id]).Kind() == reflect.Int64 {
							refId += strconv.Itoa(int(entity.GetObject().(map[string]interface{})[id].(int64)))
						}
						referenced = this.referenced[refId]
					}
				} else {
					// A static entity here...
					id = strings.Replace(id, "M_", "", -1)
					methodName := "Get" + strings.ToUpper(id[0:1]) + id[1:]
					params := make([]interface{}, 0)
					result, err := Utility.CallMethod(entity.GetObject(), methodName, params)
					if err == nil {
						referenced = this.referenced[result.(string)]
					}
				}
			}

			for j := 0; j < len(referenced); j++ {
				// So here I will get the owner entity
				owner, _ := this.getEntityByUuid(referenced[j].OwnerUuid)
				if reflect.TypeOf(owner.GetObject()).String() == "map[string]interface {}" {
					// Set the uuid as value
					if owner.GetObject().(map[string]interface{})[referenced[j].Name] != entity.GetUuid() {
						owner.GetObject().(map[string]interface{})[referenced[j].Name] = entity.GetUuid()
						owner.SetNeedSave(true)
						owner.SaveEntity()
					}
				}

			}

		}
	}

}

/**
 * A function to initialyse an entity with a given id.
 */
func (this *EntityManager) InitEntity(entity Entity) {
	this.Lock()
	defer this.Unlock()

	// Get the list of references.
	references := this.reference[entity.GetUuid()]

	// Now I will call methode to initialyse the reference...
	for i := 0; i < len(references); i++ {
		ref := references[i]
		// Retreive the reference owner...
		refOwner, errObj := this.getEntityByUuid(ref.OwnerUuid)
		if errObj == nil {
			values := strings.Split(ref.Value, "$$")
			if len(values) == 2 && refOwner != nil {
				refUUID := values[1]
				typeName := values[0]
				var refTarget Entity
				if Utility.IsValidEntityReferenceName(refUUID) {
					// We have a uuid here
					refTarget, errObj = this.getEntityByUuid(refUUID)
				} else if len(refUUID) > 0 {
					// Here we have an id not a uuid...
					storeId := typeName[:strings.Index(typeName, ".")]
					refTarget, errObj = this.getEntityById(storeId, typeName, refUUID)
				}
				// The set methode name...
				if errObj == nil && refTarget != nil {
					methodName := "Set" + strings.ToUpper(ref.Name[0:1]) + ref.Name[1:]
					params := make([]interface{}, 1)
					params[0] = refTarget.GetObject()
					_, invalidMethod := Utility.CallMethod(refOwner.(Entity).GetObject(), methodName, params)
					if invalidMethod != nil {
						// Also append referenced into the ref owner.
						fieldIndex := refOwner.GetPrototype().getFieldIndex(ref.Name)
						fieldType := refOwner.GetPrototype().FieldsType[fieldIndex]
						if strings.HasSuffix(fieldType, ":Ref") {
							// TODO verify if the reference need to be set here.
							/*if strings.HasPrefix(fieldType, "[]") {
								if reflect.TypeOf(refOwner.(*DynamicEntity).object[ref.Name]).String() == "[]string" {
									if refOwner.(*DynamicEntity).object[ref.Name] == nil {
										refOwner.(*DynamicEntity).object[ref.Name] = make([]string, 0)
									}
									if !Utility.Contains(refOwner.(*DynamicEntity).object[ref.Name].([]string), refTarget.GetUuid()) {
										refOwner.(*DynamicEntity).object[ref.Name] = append(refOwner.(*DynamicEntity).object[ref.Name].([]string), refTarget.GetUuid())
									}

								} else if reflect.TypeOf(refOwner.(*DynamicEntity).object[ref.Name]).String() == "[]interface {}" {
									log.Println("--> wrong reference type")
								}
							} else {
								refOwner.(*DynamicEntity).object[ref.Name] = refTarget.GetUuid()
							}*/
							// Append the referenced
							refTarget.AppendReferenced(ref.Name, refOwner)
							refOwner.AppendReference(refTarget)
						}
					}

				} else if len(refUUID) > 0 {
					log.Println("--------> reference target not found:", refUUID)
				}
			}

		} else {
			log.Println("--------> reference owner not found:", ref.OwnerUuid)
			return
		}
	}

	// Now the entity is in the cache...
	this.insert(entity)

	// remove it from the list...
	delete(this.reference, entity.GetUuid())
}

/**
 * Return the list of entity prototype for a given package...
 */
func (this *EntityManager) getEntityPrototypes(storeId string, schemaId string) ([]*EntityPrototype, error) {

	if GetServer().GetDataManager().getDataStore(storeId) == nil {
		return nil, errors.New("The dataStore with id '" + storeId + "' doesn't exist.")
	}

	dataStore := GetServer().GetDataManager().getDataStore(storeId)
	if reflect.TypeOf(dataStore).String() == "*Server.SqlDataStore" {
		// I will try to found the prototype inside sql_info instead.
		dataStore = GetServer().GetDataManager().getDataStore("sql_info")
	}

	protos, err := dataStore.GetEntityPrototypes()
	prototypes := make([]*EntityPrototype, 0)

	if err == nil {
		for i := 0; i < len(protos); i++ {
			if strings.HasPrefix(protos[i].TypeName, schemaId) {
				prototypes = append(prototypes, protos[i])
			}
		}
	}

	return prototypes, err
}

/**
 * Return the entity prototype for an object of a given name.
 */
func (this *EntityManager) getEntityPrototype(typeName string, storeId string) (*EntityPrototype, error) {

	if GetServer().GetDataManager().getDataStore(storeId) == nil {
		return nil, errors.New("The dataStore with id '" + storeId + "' doesn't exist.")
	}

	dataStore := GetServer().GetDataManager().getDataStore(storeId)
	if reflect.TypeOf(dataStore).String() == "*Server.SqlDataStore" {
		// I will try to found the prototype inside sql_info instead.
		dataStore = GetServer().GetDataManager().getDataStore("sql_info")
	}

	proto, err := dataStore.GetEntityPrototype(typeName)

	if err != nil {
		err = errors.New("Prototype for entity '" + typeName + "' was not found.")
	}
	return proto, err
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

/**
 * Create a new entity with default value and append it inside it parent...
 *
 * TODO Est que "The attributeName is the name of the entity in it's parent whitout the M_" est vrai ou on doit lui donner avec le M_?
 */
func (this *EntityManager) CreateEntity(parentUuid string, attributeName string, typeName string, objectId string, values interface{}, messageId string, sessionId string) interface{} {

	if !Utility.IsValidPackageName(typeName) {
		cargoError := NewError(Utility.FileLine(), INVALID_PACKAGE_NAME_ERROR, SERVER_ERROR_CODE, errors.New("The typeName '"+typeName+"' is not valid."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	var parentPtr Entity

	if len(parentUuid) > 0 {

		if !Utility.IsValidEntityReferenceName(parentUuid) {
			cargoError := NewError(Utility.FileLine(), INVALID_REFERENCE_NAME_ERROR, SERVER_ERROR_CODE, errors.New("The parentUuid '"+parentUuid+"' is not valid."))
			GetServer().reportErrorMessage(messageId, sessionId, cargoError)
			return nil
		}

		var errObj *CargoEntities.Error
		parentPtr, errObj = this.getEntityByUuid(parentUuid)
		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
	}

	params := make([]interface{}, 2)
	params[0] = objectId
	params[1] = values

	methodName := "New" + strings.Replace(typeName, ".", "", -1) + "Entity"
	entity, invalidMethod := Utility.CallMethod(this, methodName, params)

	if invalidMethod != nil {
		log.Println("--------> invalid method name:", methodName)
		// Try to create a dynamic entity...
		if reflect.TypeOf(values).String() == "map[string]interface {}" {
			values.(map[string]interface{})["parentUuid"] = parentUuid
			var errObj *CargoEntities.Error
			entity, errObj = this.newDynamicEntity(values.(map[string]interface{}))
			if errObj != nil {
				GetServer().reportErrorMessage(messageId, sessionId, errObj)
				return nil
			}
		} else {
			cargoError := NewError(Utility.FileLine(), TYPENAME_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The typeName '"+typeName+"' is does not exist."))
			GetServer().reportErrorMessage(messageId, sessionId, cargoError)
			return nil
		}
	}

	// Save the entity...
	entity.(Entity).SetNeedSave(true)
	entity.(Entity).SaveEntity()

	// Now I will save it parent if there one.
	if parentPtr != nil {
		parentPtrTypeName := parentPtr.GetTypeName()
		parentPtrStoreId := parentPtrTypeName[:strings.Index(parentPtrTypeName, ".")]
		parentPrtPrototype, _ := this.getEntityPrototype(parentPtrTypeName, parentPtrStoreId)
		if parentPrtPrototype.getFieldIndex(attributeName) < 0 {
			cargoError := NewError(Utility.FileLine(), ATTRIBUTE_NAME_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The attribute name '"+attributeName+"' does not exist in the parent entity with uuid '"+parentUuid+"'."))
			GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		} else {
			// Append the child into it parent and save it.
			parentPtr.AppendChild(attributeName, entity.(Entity))
			// Set need save at true.
			parentPtr.SetNeedSave(true)
			parentPtr.SaveEntity()
		}
	}

	// Associate the sessionId with the entityUuid in the cache
	GetServer().GetCacheManager().register(entity.(Entity).GetUuid(), sessionId)

	// Return the object.
	return entity.(Entity).GetObject()
}

/**
 * Create a new entity prototype.
 */
func (this *EntityManager) CreateEntityPrototype(storeId string, prototype interface{}, messageId string, sessionId string) *EntityPrototype {

	if reflect.TypeOf(prototype).String() != "*Server.EntityPrototype" {
		cargoError := NewError(Utility.FileLine(), PARAMETER_TYPE_ERROR, SERVER_ERROR_CODE, errors.New("Expected '*Server.EntityPrototype' but got '"+reflect.TypeOf(prototype).String()+"' instead."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	// Get the store...
	store := GetServer().GetDataManager().getDataStore(storeId)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Datastore '"+storeId+"' dosen't exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil

	}

	// Save the prototype...
	err := store.(*KeyValueDataStore).SetEntityPrototype(prototype.(*EntityPrototype))

	// TODO append super type attribute into the prototype if there is not already there...
	if err != nil {
		cargoError := NewError(Utility.FileLine(), PROTOTYPE_CREATION_ERROR, SERVER_ERROR_CODE, errors.New("Failed to create the prototype '"+prototype.(*EntityPrototype).TypeName+"' in store '"+storeId+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	return prototype.(*EntityPrototype)
}

/**
 * Return the list of entity prototype for a given package...
 */
func (this *EntityManager) GetEntityPrototypes(storeId string, messageId string, sessionId string) []*EntityPrototype {
	var schemaId string
	if strings.Index(storeId, ".") > 0 {
		schemaId = storeId[strings.Index(storeId, ".")+1:]
		storeId = storeId[0:strings.Index(storeId, ".")]
	}
	protos, err := this.getEntityPrototypes(storeId, schemaId)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), PROTOTYPE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("There is no prototypes in store '"+storeId+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}
	return protos
}

/**
 * Return the list of derived type for a given type.
 */
func (this *EntityManager) GetDerivedEntityPrototypes(typeName string, messageId string, sessionId string) []*EntityPrototype {
	protos, errObj := this.getDerivedEntityType(typeName)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}
	return protos
}

/**
 * Return the entity prototype for an object of a given name.
 */
func (this *EntityManager) GetEntityPrototype(typeName string, storeId string, messageId string, sessionId string) *EntityPrototype {
	proto, err := this.getEntityPrototype(typeName, storeId)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), ENTITY_PROTOTYPE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}
	return proto
}

/**
 * Return the object contain in entity for a given type...
 */
func (this *EntityManager) GetObjectsByType(typeName string, queryStr string, storeId string, messageId string, sessionId string) []interface{} {

	entities, errObj := this.getEntitiesByType(typeName, queryStr, storeId)

	var objects []interface{}

	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return objects
	}

	for i := 0; i < len(entities); i++ {
		// If the entity was deleted before the time i sent it back to the
		//
		if entities[i] != nil {
			// Init the entity
			entities[i].InitEntity(entities[i].GetUuid())

			// Associate the sessionId with the entityUuid in the cache
			GetServer().GetCacheManager().register(entities[i].GetUuid(), sessionId)

			objects = append(objects, entities[i].GetObject())
		}
	}

	return objects
}

/**
 * Return true if an entity with a given uuid exist.
 */
func (this *EntityManager) isExist(uuid string) bool {
	_, contained := this.contain(uuid)
	if contained {
		return true
	}

	var query EntityQuery
	query.TypeName = uuid[0:strings.Index(uuid, ".")]
	query.Fields = append(query.Fields, "uuid")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	_, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return false
	}
	return true
}

/**
 * Return the underlying object, mostly use by the client side to get object..
 */
func (this *EntityManager) GetObjectByUuid(uuid string, messageId string, sessionId string) interface{} {
	entity, errObj := this.getEntityByUuid(uuid)
	if errObj != nil {
		entity, errObj = this.getDynamicEntityByUuid(uuid)
		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
	}

	// Associate the sessionId with the entityUuid in the cache
	GetServer().GetCacheManager().register(entity.GetUuid(), sessionId)

	return entity.GetObject()
}

/**
 * Return the underlying object, mostly use by the client side to get object..
 */
func (this *EntityManager) GetObjectById(storeId string, typeName string, id string, messageId string, sessionId string) interface{} {

	entity, errObj := this.getEntityById(storeId, typeName, id)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}
	// Associate the sessionId with the entityUuid in the cache
	GetServer().GetCacheManager().register(entity.GetUuid(), sessionId)

	return entity.GetObject()
}

/**
 * Save the vlaues of an entity.
 */
func (this *EntityManager) SaveEntity(values interface{}, typeName string, messageId string, sessionId string) interface{} {

	// Need to be call before any new entity function to test the new value
	// with the actual one.
	funcName := "New" + strings.Replace(typeName, ".", "", -1) + "EntityFromObject"

	params := make([]interface{}, 1)
	params[0] = values

	entity, err := Utility.CallMethod(this, funcName, params)

	if err != nil {
		// I will try with dynamic entity insted...
		var errObj *CargoEntities.Error
		entity, errObj = this.newDynamicEntity(values.(map[string]interface{}))
		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
	}

	// Now I will save the entity.
	entity.(Entity).SaveEntity()

	// Associate the sessionId with the entityUuid in the cache
	GetServer().GetCacheManager().register(entity.(Entity).GetUuid(), sessionId)

	return entity.(Entity).GetObject()
}

/**
 * Remove an existing entity with a given uuid.
 */
func (this *EntityManager) RemoveEntity(uuid string, messageId string, sessionId string) {

	// The entity to remove.
	var entity Entity
	var errObj *CargoEntities.Error

	// validate the action. TODO active it latter...
	/*errObj = GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return // exit here.
	}*/

	entity, errObj = this.getEntityByUuid(uuid)
	if errObj != nil {
		entity, errObj = this.getDynamicEntityByUuid(uuid)
	}

	if entity != nil {
		// validate over the entity TODO active it latter...
		/*errObj = GetServer().GetSecurityManager().hasPermission(sessionId, CargoEntities.PermissionType_Delete, entity)
		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return // exit here.
		}*/

		// Suppress the enitity...
		entity.DeleteEntity()

		// exit here.
		return
	}

	// Repport the error
	GetServer().reportErrorMessage(messageId, sessionId, errObj)
}

/**
 * Return the list of all link's for a given entity.
 */
func (this *EntityManager) GetEntityLnks(uuid string, messageId string, sessionId string) []Entity {

	// The entity to remove.
	var entity Entity
	var errObj *CargoEntities.Error
	entity, errObj = this.getEntityByUuid(uuid)
	visited := make([]string, 0)
	lnkLst := make([]Entity, 0)

	if entity != nil {
		this.getEntityLnkLst(entity, &visited, &lnkLst)
	}

	// Repport the error
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}

	return lnkLst
}

////////////////////////////////////////////////////////////////////////////////
// XSD/XML Schemas...
////////////////////////////////////////////////////////////////////////////////

func (this *EntityManager) generateXsdSchema(schemaId string, filePath string) (*EntityPrototype, error) {

	/** First of all I will generate the javascript file **/
	// Execute the command...

	// First of all I will

	return nil, nil
}
