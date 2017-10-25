package Server

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
	"code.myceliUs.com/XML_Schemas"
)

const (
	// The persistence db
	CargoEntitiesDB = "CargoEntities"
	createdFormat   = "2006-01-02 15:04:05"
)

/**
 * Data manager can be use to retreive data inside a data store, like Sql server
 * from file like xml file or any other source...
 */
type DataManager struct {
	/** This contain connection to know dataStore **/
	m_dataStores map[string]DataStore

	/**
	 * Use to protected the entitiesMap access...
	 */
	sync.RWMutex
}

var dataManager *DataManager

func (this *Server) GetDataManager() *DataManager {
	if dataManager == nil {
		dataManager = newDataManager()
	}
	return dataManager
}

/**
 * This is the accessing function to dataStore...
 */
func newDataManager() *DataManager {

	// Register dynamic type here...
	dataManager := new(DataManager)
	dataManager.m_dataStores = make(map[string]DataStore)

	/** Now I will initialyse data store one by one... **/
	defaultStoreConfigurations := GetServer().GetConfigurationManager().getDefaultDataStoreConfigurations()

	for i := 0; i < len(defaultStoreConfigurations); i++ {
		dataManager.appendDefaultDataStore(defaultStoreConfigurations[i])
	}

	/** Return the data manager pointer... **/
	return dataManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

func (this *DataManager) initialize() {

	// Create the default configurations
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)

	// Here I will get the datastore configuration...
	activeConfigurationsEntity, err := GetServer().GetConfigurationManager().getActiveConfigurationsEntity()
	if err != nil {
		log.Panicln(err)
	}

	storeConfigurations := activeConfigurationsEntity.GetObject().(*Config.Configurations).GetDataStoreConfigs()

	log.Println("--> initialyze DataManager")
	for i := 0; i < len(storeConfigurations); i++ {
		if this.m_dataStores[storeConfigurations[i].GetId()] == nil {
			store, err := NewDataStore(storeConfigurations[i])
			if err != nil {
			} else {
				this.m_dataStores[store.GetId()] = store
			}
		}
	}
}

func (this *DataManager) getId() string {
	return "DataManager"
}

func (this *DataManager) start() {
	log.Println("--> Start DataManager")
}

func (this *DataManager) stop() {
	log.Println("--> Stop DataManager")
	this.close()
}

////////////////////////////////////////////////////////////////////////////////
// private function
////////////////////////////////////////////////////////////////////////////////

func (this *DataManager) openConnections() {
	for _, store := range this.m_dataStores {
		err := store.Connect()
		// Call get entity prototype once to initialyse entity prototypes.
		if err == nil {
			store.GetEntityPrototypes()
		}
	}
}

func (this *DataManager) appendDefaultDataStore(config *Config.DataStoreConfiguration) {
	store, err := NewDataStore(config)
	if err != nil {
		log.Println(err)
	}
	this.m_dataStores[store.GetId()] = store
	store.Connect()
}

/**
 * Access a store with here given name...
 */
func (this *DataManager) getDataStore(id string) DataStore {
	this.Lock()
	defer this.Unlock()
	store := this.m_dataStores[id]
	return store
}

/**
 * Remove a dataStore from the map
 */
func (this *DataManager) removeDataStore(id string) {
	this.Lock()
	defer this.Unlock()

	// Close the connection.
	this.m_dataStores[id].Close()

	// Remove data. Deadlock error.
	// this.m_dataStores[name].DeleteEntityPrototypes()

	// Delete the reference from the map.
	delete(this.m_dataStores, id)
}

/**
 * Execute a query that read information from the store and
 * return the result and an array of interface...
 */
func (this *DataManager) readData(storeId string, query string, fieldsType []interface{}, params []interface{}) ([][]interface{}, error) {

	store := this.getDataStore(storeId)
	if store == nil {
		return nil, errors.New("The datastore '" + storeId + "' does not exist.")
	}

	data, err := store.Read(query, fieldsType, params)

	if err != nil {
		err = errors.New("Query '" + query + "' failed with error '" + err.Error() + "'.")
		return data, err
	}

	// In case of SQL data, the data we found will be use to get
	// sql data in a second pass.
	if storeId == "sql_info" && err == nil {
		if len(data) > 0 {
			if len(data[0]) > 0 {
				dataIndex := make([]int, 0)
				q := new(EntityQuery)
				json.Unmarshal([]byte(query), q)

				// In that case the value will be read from sql.
				for i := 0; i < len(data); i++ {
					if len(data[i]) > 1 {
						if reflect.TypeOf(data[i][0]).Kind() == reflect.String {
							if Utility.IsValidEntityReferenceName(data[i][0].(string)) {
								// So I will create the sql query to get it data back.
								uuid := data[i][0].(string)
								values := strings.Split(uuid[0:strings.Index(uuid, "%")], ".")
								dataBaseName := values[0]
								tableName := values[len(values)-1]
								schemaId := ""
								if len(values) == 3 {
									schemaId = values[1]
								}
								prototype, err := GetServer().GetEntityManager().getEntityPrototype(uuid[0:strings.Index(uuid, "%")], "sql_info")
								if err == nil {
									// Now I will create the query.
									var ids []interface{}
									var fields []string
									var fieldsType []interface{}
									for j := 0; j < len(q.Fields); j++ {
										fieldName := q.Fields[j]
										fieldType := prototype.FieldsType[prototype.getFieldIndex(fieldName)]
										// If the field is an id
										if !strings.HasSuffix(fieldType, ":Ref") && strings.HasPrefix(fieldName, "M_") && !isForeignKey(fieldName) {
											if j < len(data[i]) {
												if Utility.Contains(prototype.Ids, fieldName) {
													ids = append(ids, data[i][j]) // append the id
												}
												fields = append(fields, fieldName[2:])
												fieldsType = append(fieldsType, fieldType)
												dataIndex = append(dataIndex, j)
											}
										}
									}

									if len(fields) > 0 {
										// Now I will recreate the sql query string.
										query := "SELECT "
										for j := 0; j < len(fields); j++ {
											query += fields[j]
											if j < len(fields)-1 {
												query += ", "
											}
										}

										// The from close.
										query += " FROM " + dataBaseName
										if len(schemaId) > 0 {
											query += "." + schemaId
										}
										query += "." + tableName

										if len(ids) > 0 {
											query += " WHERE "
											for j := 0; j < len(ids); j++ {
												// The first ids in the list of ids are always the uuid so
												// the index is j+1
												id := Utility.ToString(ids[j])
												if id != "null" {
													query += strings.Replace(prototype.Ids[j+1], "M_", "", -1) + "=?"
													if j < len(ids)-1 {
														query += " AND "
													}
													params = append(params, id)
												}
											}
										}

										// Now I will get data from sql...
										sqlData, err := this.readData(dataBaseName, query, fieldsType, params)
										if err == nil {
											if len(sqlData) > 0 {
												// Now I will replace the data with the retreive values.
												for j := 0; j < len(sqlData[0]); j++ {
													// Set the value.
													data[i][dataIndex[j]] = sqlData[0][j]
												}

											} else {
												return data, errors.New("No sql data was found for entity " + data[i][0].(string))
											}

										} else {
											return data, err
										}
									} else {
										return data, errors.New("No sql data was found for entity " + data[i][0].(string))
									}
								}
							}
						}
					} else if len(data[i]) == 1 {
						return data, nil
					}
				}
			} else {
				err = errors.New("No data found!")
				data = nil
			}
		} else {
			err = errors.New("No data found!")
			data = nil
		}
	}

	return data, err
}

// One to many relationship
func (this *DataManager) setOneToManyEntityRelationship(name string, src *DynamicEntity, dest *DynamicEntity, isRef bool) {
	if isRef {
		// here the relation is an array...
		refsUuid := src.getValue(name)
		if refsUuid == nil {
			refsUuid = make([]string, 0)
		}

		if !Utility.Contains(refsUuid.([]string), dest.GetUuid()) {
			refsUuid = append(refsUuid.([]string), dest.GetUuid())
			src.SetNeedSave(true)
		}

		// Set the source uuid.
		src.setValue(name, refsUuid)

		// Set reference.
		src.AppendReference(dest)

		// Set referenced.
		dest.AppendReferenced(name, src)

	} else {
		src.AppendChild(name, dest)
	}

	// Now if will set the ref in dest...
	refUuid := dest.getValue(name)
	if src.GetUuid() != refUuid {
		dest.setValue(name, src.GetUuid())
		dest.SetNeedSave(true)
		src.AppendReferenced(name, dest)
		dest.AppendReference(src)
	}

	// Save entities
	dest.saveEntity(dest.GetUuid())
	src.saveEntity(src.GetUuid())
}

// One to one relaltionship.
func (this *DataManager) setOneToOneEntityRelationship(name string, src *DynamicEntity, dest *DynamicEntity, isRef bool) {
	if isRef {
		uuid := src.getValue(name).(string)
		// Set the other side of relation ship.
		if uuid != dest.GetUuid() {
			src.setValue(name, dest.GetUuid())
			src.SetNeedSave(true)
		}

		// Set reference.
		src.AppendReference(dest)
		// Set referenced.
		dest.AppendReferenced(name, src)

	} else {
		src.AppendChild(name, dest)
	}

	// Now if will set the ref in dest...
	refUuid := dest.getValue(name)
	if src.GetUuid() != refUuid {
		dest.setValue(name, src.GetUuid())
		dest.SetNeedSave(true)
		src.AppendReferenced(name, dest)
		dest.AppendReference(src)
	}

	// Save entities
	dest.saveEntity(dest.GetUuid())
	src.saveEntity(src.GetUuid())
}

/**
 * Return the list of entities for a given relationship from db...
 */
func (this *DataManager) getRelationshipEntities(prototype *EntityPrototype, fields []string, ids []string) ([]*DynamicEntity, error) {

	// The entities
	entities := make([]*DynamicEntity, 0)

	table := prototype.TypeName

	query := "SELECT "
	fieldsType := make([]interface{}, 0)

	// The first element is the uuid and are not store in sql.
	for i := 1; i < len(prototype.Ids); i++ {
		query += strings.Replace(prototype.Ids[i], "M_", "", -1)
		fieldsType = append(fieldsType, prototype.FieldsType[prototype.getFieldIndex(prototype.Ids[i])])
		if i < len(prototype.Ids)-1 {
			query += ","
		}
	}

	query += " FROM " + table + " WHERE "

	// append the fields
	for i := 0; i < len(fields); i++ {
		query += strings.Replace(fields[i], "M_", "", -1) + "=?"
		if i < len(fields)-1 {
			query += " AND "
		}
	}

	// append the ids
	params := make([]interface{}, 0)
	for i := 0; i < len(ids); i++ {
		params = append(params, ids[i])
	}

	storeName := table[0:strings.Index(table, ".")]
	store := this.getDataStore(storeName)
	if store == nil {
		return entities, errors.New("The datastore '" + storeName + "' does not exist.")
	}

	// Retreive id's
	data, err := store.Read(query, fieldsType, params)
	if err != nil {
		return entities, err
	}

	// No I will get entities from their id's
	for i := 0; i < len(data); i++ {
		entity, err := GetServer().GetEntityManager().getEntityById("sql_info", table, data[i], false)
		if err == nil {
			entities = append(entities, entity.(*DynamicEntity))
		}
	}

	return entities, nil
}

/**
 * From entity I will get references and set it.
 */
func (this *DataManager) setEntityRelationship(storeId string, name string, ref_0 *DynamicEntity) error {

	// First of all I will get the data store
	store := this.m_dataStores[storeId]
	if store == nil {
		return errors.New("No data store was found with id " + storeId)
	}

	// Now I will retreive the relationship information from sql
	// * remove the M_ prefix.
	refInfos, err := store.(*SqlDataStore).getRefInfos(name[2:])

	if err == nil && len(refInfos) > 0 {

		// In that case the ref_0 is the source entity.
		// Now I will determine the kind of relationship.
		typeName := storeId
		if len(refInfos[0][4]) > 0 && store.(*SqlDataStore).m_vendor != Config.DataStoreVendor_MYSQL {
			typeName += "." + refInfos[0][4]
		}

		typeName += "." + refInfos[0][0]
		prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, "sql_info")

		if store.(*SqlDataStore).isAssociative(prototype) && !store.(*SqlDataStore).isAssociative(ref_0.GetPrototype()) {
			// So here I want to get the list of associative value from the
			// sql db and create entities for it if it dosent exist.
			// I also want to set references in ref_0 with the associative value.

			// I will get all values from the sql associative table.
			fields := make([]string, 0)
			fieldsType := make([]interface{}, 0)
			for i := 1; i < len(prototype.Ids); i++ {
				fields = append(fields, prototype.Ids[i][2:])
				fieldsType = append(fieldsType, prototype.FieldsType[prototype.getFieldIndex(prototype.Ids[i])])
			}

			// In that case I want to initialyse the data from sql.
			query := "SELECT "
			for i := 0; i < len(fields); i++ {
				query += fields[i]

				if i < len(fields)-1 {
					query += ", "
				}
			}

			query += " FROM " + typeName + " WHERE "
			params := make([]interface{}, 0)
			for i := 0; i < len(refInfos); i++ {
				query += refInfos[i][1] + "=?"
				if strings.HasSuffix(ref_0.GetTypeName(), refInfos[i][0]) {
					params = append(params, ref_0.getValue("M_"+refInfos[i][1]))
				} else if strings.HasSuffix(ref_0.GetTypeName(), refInfos[i][2]) {
					params = append(params, ref_0.getValue("M_"+refInfos[i][3]))
				}
				if i < len(refInfos)-1 {
					query += " AND "
				}
			}

			storeName := typeName[0:strings.Index(typeName, ".")]
			store := this.getDataStore(storeName)
			if store == nil {
				return errors.New("The datastore '" + storeName + "' does not exist.")
			}

			data, err := store.Read(query, fieldsType, params)

			if err != nil {
				log.Println(err)
				return err
			}

			// Now from the data I will retreive the associative entity.
			for i := 0; i < len(data); i++ {
				keyInfo := typeName + ":"
				for j := 0; j < len(data[i]); j++ {
					keyInfo += Utility.ToString(data[i][j])
					// Append underscore for readability in case of problem...
					if j < len(data[i])-1 {
						keyInfo += "_"
					}
				}
				uuid := typeName + "%" + Utility.GenerateUUID(keyInfo)
				associativeEntitiy, _ := GetServer().GetEntityManager().getEntityByUuid(uuid, false)
				if associativeEntitiy != nil {
					// one side of relationship
					associativeEntitiy.(*DynamicEntity).AppendReference(ref_0)
					associativeEntitiy.(*DynamicEntity).setValue(name, ref_0.GetUuid())
					ref_0.AppendReferenced(name, associativeEntitiy)

					// other side of the relationship
					ref_0.AppendReference(associativeEntitiy)
					ref_0.appendValue(name, associativeEntitiy.GetUuid())
					associativeEntitiy.AppendReferenced(name, ref_0)
				}
			}
		} else if !store.(*SqlDataStore).isAssociative(ref_0.GetPrototype()) {
			// The ref_0 must be the source of relationship.
			if strings.HasSuffix(ref_0.GetTypeName(), refInfos[0][2]) {

				fieldType := ref_0.GetPrototype().FieldsType[ref_0.GetPrototype().getFieldIndex(name)]
				isArray := strings.HasPrefix(fieldType, "[]")
				isRef := strings.HasSuffix(fieldType, ":Ref")

				typeName := strings.Replace(fieldType, "[]", "", -1)
				typeName = strings.Replace(typeName, ":Ref", "", -1)

				// From the table i will retreive the entity prototype.
				prototype, err := GetServer().GetEntityManager().getEntityPrototype(typeName, "sql_info")
				if err != nil {
					return err
				}

				ids := make([]string, 0)
				fields := make([]string, 0)
				for i := 0; i < len(refInfos); i++ {
					id := Utility.ToString(ref_0.getValue("M_" + refInfos[0][3]))
					ids = append(ids, id)
					fields = append(fields, "M_"+refInfos[0][1])
				}

				// ref entities contain the list of existing ref in the database...
				ref_entities, err := this.getRelationshipEntities(prototype, fields, ids)

				if err != nil {
					log.Println("-------> error ", err)
					return err
				}

				if isArray {
					for i := 0; i < len(ref_entities); i++ {
						// The one to many relationship.
						this.setOneToManyEntityRelationship(name, ref_0, ref_entities[i], isRef)
					}
				} else {
					if len(ref_entities) == 1 {
						this.setOneToOneEntityRelationship(name, ref_0, ref_entities[0], isRef)
					}
				}

			} else {
				// The field type
				fieldType := ref_0.GetPrototype().FieldsType[ref_0.GetPrototype().getFieldIndex(name)]
				typeName := strings.Replace(fieldType, "[]", "", -1)
				typeName = strings.Replace(typeName, ":Ref", "", -1)

				prototype, err := GetServer().GetEntityManager().getEntityPrototype(typeName, "sql_info")
				if err != nil {
					return err
				}

				ids := make([]string, 0)
				fields := make([]string, 0)
				for i := 0; i < len(refInfos); i++ {

					// All id's must be set to create the lnk.
					if ref_0.getValue("M_"+refInfos[0][1]) == nil {
						return errors.New("No id field " + "M_" + refInfos[0][1] + " was set for entity " + ref_0.GetUuid())
					}

					id := Utility.ToString(ref_0.getValue("M_" + refInfos[0][1]))
					ids = append(ids, id)
					fields = append(fields, "M_"+refInfos[0][3])
				}

				ref_entities, err := this.getRelationshipEntities(prototype, fields, ids)

				// So here I will retreive the source value.
				if err == nil {
					if len(ref_entities) == 1 {
						ref_1 := ref_entities[0]
						index := ref_1.GetPrototype().getFieldIndex(name)
						if index > -1 {
							fieldType := ref_1.GetPrototype().FieldsType[index]
							isArray := strings.HasPrefix(fieldType, "[]")
							isRef := strings.HasSuffix(fieldType, ":Ref")
							if isArray {
								// The one to many relationship.
								this.setOneToManyEntityRelationship(name, ref_1, ref_0, isRef)
							} else {
								this.setOneToOneEntityRelationship(name, ref_1, ref_0, isRef)
							}
						}
					}
				}
			}
		}
	}

	return nil
}

/**
 * If the entity is save in sql database reference are not automaticaly set
 * instead a field that contain the reference id are save in the db. Here I
 * will retreive the associated entity and set it inside the M_FK_field_name.
 */
func (this *DataManager) setEntityReferences(uuid string, lazy bool) error {
	entity, err := GetServer().GetEntityManager().getEntityByUuid(uuid, lazy)

	if err != nil {
		return errors.New(err.GetBody())
	}
	prototype := entity.GetPrototype()

	// I will retreive reference fields.
	for i := 0; i < len(prototype.FieldsType); i++ {
		// I need to retreive the link between for example M_post_id and M_FK_blog_comment_blog_post.
		if isForeignKey(prototype.Fields[i]) {
			storeId := prototype.TypeName[0:strings.Index(prototype.TypeName, ".")]
			this.setEntityRelationship(storeId, prototype.Fields[i], entity.(*DynamicEntity))
		}
	}
	return nil
}

/**
 * Execute a query that create a new data. The data contains the new
 * value to insert in the DB.
 */
func (this *DataManager) createData(storeName string, query string, d []interface{}) (lastId interface{}, err error) {
	// If the store is sql_info in that case I will need to create the information
	// in the sql data store.
	store := this.getDataStore(storeName)
	if store == nil {
		return nil, errors.New("Data store '" + storeName + " does not exist.")
	}

	// Create the entity...
	lastId, err = store.Create(query, d)
	if err != nil {
		err = errors.New("Query '" + query + "' failed with error '" + err.Error() + "'.")
		return
	}

	// In the case of sql data I also need to save the information in the database.
	if storeName == "sql_info" && len(d) > 0 {
		if reflect.TypeOf(d[0]).Kind() == reflect.String {
			uuid := d[0].(string)
			values := strings.Split(uuid[0:strings.Index(uuid, "%")], ".")
			dataBaseName := values[0]
			tableName := values[len(values)-1]

			// we are not interested in system tables.
			if tableName == "sysdiagrams" {
				return nil, errors.New("system table " + tableName + " is not a valid table.")
			}
			schemaId := ""
			if len(values) == 3 {
				schemaId = values[1]
			}
			prototype, err := GetServer().GetEntityManager().getEntityPrototype(uuid[0:strings.Index(uuid, "%")], "sql_info")

			q := new(EntityQuery)
			json.Unmarshal([]byte(query), q)

			if err == nil {
				query := "INSERT INTO " + dataBaseName
				if len(schemaId) > 0 {
					query += "." + schemaId
				}
				data := make([]interface{}, 0)
				fields := make([]string, 0)
				fieldsType := make([]interface{}, 0)

				query += "." + tableName + "("
				values := "VALUES("
				for i := 0; i < len(q.Fields); i++ {
					fieldName := q.Fields[i]
					index := prototype.getFieldIndex(fieldName)
					if index > 0 {
						fieldType := prototype.FieldsType[index]
						if strings.HasPrefix(fieldName, "M_") && !strings.HasSuffix(fieldType, ":Ref") && !isForeignKey(fieldName) {
							fields = append(fields, fieldName)
							fieldsType = append(fieldsType, fieldType)
							// In case of null value...
							if reflect.TypeOf(d[i]).Kind() == reflect.String {
								if d[i] == "null" {
									// if the field is an id it must not be null
									if Utility.Contains(prototype.Ids, fieldName) {
										return -1, errors.New(prototype.TypeName + "." + fieldName + " is null.")
									}
									d[i] = "NULL"
								}
							}

							// Here I will convert the data
							if XML_Schemas.IsXsBoolean(fieldType) && reflect.TypeOf(d[i]).String() == "float64" {
								data = append(data, int8(d[i].(float64)))
							} else if XML_Schemas.IsXsInt(fieldType) && reflect.TypeOf(d[i]).String() == "float64" {
								data = append(data, int32(d[i].(float64)))
							} else if XML_Schemas.IsXsDate(fieldType) {
								dateTime, err := Utility.MatchISO8601_DateTime(d[i].(string))
								if err == nil {
									data = append(data, dateTime.Format(createdFormat))
								} else {
									data = append(data, d[i])
								}
							} else {
								data = append(data, d[i])
							}
						}
					}
				}

				for i := 0; i < len(fields); i++ {
					values += "?"
					query += fields[i][2:]
					if i < len(fields)-1 {
						values += ","
						query += ","
					}
				}

				// Set the values...
				query += ")" + values + ")"
				lastId, err = this.createData(dataBaseName, query, data)
				if err != nil {
					log.Println("---> data insert fail with err: ", err)
				}
			}
		}
	}

	return
}

func (this *DataManager) deleteData(storeName string, query string, params []interface{}) (err error) {
	store := this.getDataStore(storeName)

	if store == nil {
		return errors.New("Data store " + storeName + " does not exist.")
	}

	// Now if the entity has sql backend.
	if storeName == "sql_info" {

		var entityQuery *EntityQuery
		err = json.Unmarshal([]byte(query), &entityQuery)
		if err != nil {
			return err
		}

		// Get the entity uuid to delete.
		uuid := strings.Split(entityQuery.Indexs[0], "=")[1]

		if Utility.IsValidEntityReferenceName(uuid) {

			values := strings.Split(uuid[0:strings.Index(uuid, "%")], ".")
			dataBaseName := values[0]
			tableName := values[len(values)-1]
			schemaId := ""
			if len(values) == 3 {
				schemaId = values[1]
			}

			prototype, err := GetServer().GetEntityManager().getEntityPrototype(entityQuery.TypeName, "sql_info")

			ids := make([]interface{}, 0)
			query := "DELETE FROM " + dataBaseName
			if len(schemaId) > 0 {
				query += "." + schemaId
			}
			query += "." + tableName + " WHERE "

			if err == nil {
				entity, err := GetServer().GetEntityManager().getEntityByUuid(uuid, false)
				if err == nil {
					for i := 0; i < len(prototype.Ids); i++ {
						if strings.HasPrefix(prototype.Ids[i], "M_") {
							ids = append(ids, entity.(*DynamicEntity).getValue(prototype.Ids[i]))
							query += prototype.Ids[i][2:] + "=?"
							if i < len(prototype.Ids)-1 {
								query += " AND "
							}
						}
					}
					this.deleteData(dataBaseName, query, ids)
				}
			}
		}
	}

	err = store.Delete(query, params)
	if err != nil {
		err = errors.New("Query '" + query + "' failed with error '" + err.Error() + "'.")
	}

	return
}

func (this *DataManager) updateData(storeName string, query string, fields []interface{}, params []interface{}) (err error) {
	store := this.getDataStore(storeName)
	if store == nil {
		return errors.New("Data store " + storeName + " does not exist.")
	}

	err = store.Update(query, fields, params)
	if err != nil {
		err = errors.New("Query '" + query + "' failed with error '" + err.Error() + "'.")
		return
	}

	// In case of entity with sql database backend.
	if storeName == "sql_info" {
		var entityQuery EntityQuery
		err = json.Unmarshal([]byte(query), &entityQuery)
		if err != nil {
			return err
		}

		// Here I will use the entity query instead of the prototype.
		if reflect.TypeOf(fields[0]).Kind() == reflect.String {
			if Utility.IsValidEntityReferenceName(fields[0].(string)) {
				// I will get information from it uuid.
				uuid := fields[0].(string)
				values := strings.Split(uuid[0:strings.Index(uuid, "%")], ".")
				dataBaseName := values[0]
				tableName := values[len(values)-1]
				schemaId := ""
				if len(values) == 3 {
					schemaId = values[1]
				}

				fieldsName := make([]string, 0)
				data := make([]interface{}, 0)
				ids := make([]interface{}, 0)
				idsFieldsName := make([]string, 0)
				prototype, _ := GetServer().GetEntityManager().getEntityPrototype(entityQuery.TypeName, "sql_info")
				for i := 0; i < len(entityQuery.Fields); i++ {
					if strings.HasPrefix(entityQuery.Fields[i], "M_") && !isForeignKey(entityQuery.Fields[i]) {
						fieldType := prototype.FieldsType[prototype.getFieldIndex(entityQuery.Fields[i])]
						if !strings.HasSuffix(fieldType, ":Ref") {
							if Utility.Contains(prototype.Ids, entityQuery.Fields[i]) {
								ids = append(ids, fields[i])
								idsFieldsName = append(idsFieldsName, entityQuery.Fields[i][2:])
							} else {
								fieldsName = append(fieldsName, entityQuery.Fields[i][2:])
								if XML_Schemas.IsXsBoolean(fieldType) && reflect.TypeOf(fields[i]).String() == "float64" {
									data = append(data, int8(fields[i].(float64)))
								} else if XML_Schemas.IsXsInt(fieldType) && reflect.TypeOf(fields[i]).String() == "float64" {
									data = append(data, int32(fields[i].(float64)))
								} else if XML_Schemas.IsXsDate(fieldType) {
									if reflect.TypeOf(fields[i]).String() == "time.Time" {
										data = append(data, fields[i].(time.Time).Format(createdFormat))
									} else {
										dateTime, err := Utility.MatchISO8601_DateTime(fields[i].(string))
										if err == nil {
											data = append(data, dateTime.Format(createdFormat))
										} else {
											data = append(data, fields[i])
										}
									}
								} else {
									data = append(data, fields[i])
								}
							}
						}
					}
				}
				if len(fieldsName) > 0 {
					// The sql query.
					query := "UPDATE " + dataBaseName
					if len(schemaId) > 0 {
						query += "." + schemaId
					}
					query += "." + tableName + " SET "

					for i := 0; i < len(fieldsName); i++ {
						query += fieldsName[i] + "=?"
						if i < len(fieldsName)-1 {
							query += ", "
						}
					}

					if len(ids) > 0 {
						query += " WHERE "
						for i := 0; i < len(ids); i++ {
							query += idsFieldsName[i] + "=?"
							if i < len(ids)-1 {
								query += " AND "
							}
						}
					}

					// Update the entity.
					err = this.updateData(dataBaseName, query, data, ids)

				}
			}
		}
	}

	return
}

func (this *DataManager) createDataStore(storeId string, storeName string, hostName string, ipv4 string, port int, storeType Config.DataStoreType, storeVendor Config.DataStoreVendor) (DataStore, *CargoEntities.Error) {

	log.Println("------> create store id:", storeId, "name:", storeName, "host:", hostName, "ipv4:", ipv4, "port:", port, "type:", storeType, "vendor:", storeVendor)
	if !Utility.IsValidVariableName(storeId) {
		cargoError := NewError(Utility.FileLine(), INVALID_VARIABLE_NAME_ERROR, SERVER_ERROR_CODE, errors.New("The storeId '"+storeId+"' is not valid."))
		return nil, cargoError
	}

	if this.getDataStore(storeId) != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ALREADY_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The storeId '"+storeId+"' already exists."))
		return nil, cargoError
	}

	var storeConfig *Config.DataStoreConfiguration
	ids := []interface{}{storeId}
	storeConfigEntity, err_ := GetServer().GetEntityManager().getEntityById("Config", "Config.DataStoreConfiguration", ids, false)

	// Create the new store here.
	if err_ != nil {
		storeConfig = new(Config.DataStoreConfiguration)
		storeConfig.M_id = storeId
		storeConfig.M_storeName = storeName
		storeConfig.M_dataStoreVendor = storeVendor
		storeConfig.M_dataStoreType = storeType
		storeConfig.M_ipv4 = ipv4
		storeConfig.M_hostName = hostName
		storeConfig.M_port = port

		configEntity := GetServer().GetConfigurationManager().m_activeConfigurationsEntity
		storeConfigEntity, err_ = GetServer().GetEntityManager().createEntity(configEntity.GetUuid(), "M_dataStoreConfigs", "Config.DataStoreConfiguration", storeId, storeConfig)

		if err_ != nil {
			return nil, err_
		}

	} else {
		storeConfig = storeConfigEntity.GetObject().(*Config.DataStoreConfiguration)
		storeConfig.M_id = storeId
		storeConfig.M_storeName = storeName
		storeConfig.M_dataStoreVendor = storeVendor
		storeConfig.M_dataStoreType = storeType
		storeConfig.M_ipv4 = ipv4
		storeConfig.M_hostName = hostName
		storeConfig.M_port = port
		storeConfigEntity.SaveEntity()
	}

	// Create the store here.
	store, err := NewDataStore(storeConfig)
	if err == nil {
		// Append the new dataStore configuration.
		this.Lock()
		this.m_dataStores[storeId] = store
		this.Unlock()
		err := store.Connect()

		// Create entity prototypes.
		if err == nil {
			store.GetEntityPrototypes()
		} else {
			log.Println("---> fail to connect to ", hostName, err)
		}
	} else {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Failed to create dataStore with id '"+storeId+"' and with error '"+err.Error()+"'."))
		return nil, cargoError
	}

	return store, nil
}

func (this *DataManager) deleteDataStore(storeId string) *CargoEntities.Error {

	if !Utility.IsValidVariableName(storeId) {
		cargoError := NewError(Utility.FileLine(), INVALID_VARIABLE_NAME_ERROR, SERVER_ERROR_CODE, errors.New("The storeId '"+storeId+"' is not valid."))
		return cargoError
	}

	if this.getDataStore(storeId) == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The storeId '"+storeId+"' doesn't exist."))
		log.Println("------> Store with id", storeId, "dosen't exist!")
		return cargoError
	}

	// Delete the dataStore configuration
	dataStoreConfigurationUuid := ConfigDataStoreConfigurationExists(storeId)
	dataStoreConfigurationEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(dataStoreConfigurationUuid, false)

	// In case of the configuration is not already deleted...
	if errObj == nil {
		dataStoreConfigurationEntity.DeleteEntity()
	}

	// Remove the storeObject from the storeMap
	this.removeDataStore(storeId)

	// Delete the directory
	filePath := GetServer().GetConfigurationManager().GetDataPath() + "/" + storeId
	err := os.RemoveAll(filePath)

	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Failed to delete directory '"+filePath+"' with error '"+err.Error()+"'."))
		log.Println("------> Fail to remove ", storeId, err)
		return cargoError
	}

	return nil

}

func (this *DataManager) close() {
	this.Lock()
	defer this.Unlock()

	// Close the data manager.
	for _, v := range this.m_dataStores {
		v.Close()
	}

}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

// @api 1.0
// Event handler function.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
//DataManager.prototype.onEvent = function (evt) {
//    EventHub.prototype.onEvent.call(this, evt)
//}
func (this *DataManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

// @api 1.0
// Test if a datastore is reachable.
// @param {string} storeName The data server connection (configuration) id
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) Ping(storeName string, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	store := this.getDataStore(storeName)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("The datastore '"+storeName+"' does not exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}
	err := store.Ping()

	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Fail to ping the data store "+err.Error()+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}
}

// @api 1.0
// Open a new connection with the datastore.
// @param {string} storeName The data server connection (configuration) id
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) Connect(storeName string, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	store := this.getDataStore(storeName)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("The datastore '"+storeName+"' does not exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}
	err := store.Connect()

	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Fail to open the data store connection "+err.Error()+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	// I will get it entity prototypes.
	store.GetEntityPrototypes()

}

// @api 1.0
// Close the connection to the datastore with a given id.
// @param {string} storeName The data server connection (configuration) id
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) Close(storeName string, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	store := this.getDataStore(storeName)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("The datastore '"+storeName+"' does not exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}
	err := store.Close()

	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Fail to close the data store connection "+err.Error()+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}
}

// @api 1.0
// Execute a read query on the data sever.
// @param {string} storeName The data server connection (configuration) id
// @param {string} query The query string to execute.
// @param {[]string} fieldsType Contain the list of type of queryied data. ex. string, date, int, float.
// @param {[]interface{}} parameters Contain filter expression, ex. id=0, id != 3.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{[][]interface} Return a tow dimentionnal array of values.
// @scope {public}
// @param {callback} progressCallback The function is call when chunk of response is received.
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) Read(storeName string, query string, fieldsType []interface{}, parameters []interface{}, messageId string, sessionId string) [][]interface{} {

	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	data, err := this.readData(storeName, query, fieldsType, parameters)

	if err != nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	return data
}

// @api 1.0
// Create an new entry in the DB and return it's id(s)
// @param {string} storeName The data server connection (configuration) id
// @param {string} query The query string to execute.
// @param {[]interface{}} d values Contain the list of values associated with the fields in the query.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {interface{}} Return the last id if there is primary key set on the datastore table.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) Create(storeName string, query string, d []interface{}, messageId string, sessionId string) interface{} {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	lastId, err := this.createData(storeName, query, d)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err))
		return -1
	}
	return lastId
}

// @api 1.0
// Update existing database values.
// @param {string} connectionId The data server connection (configuration) id
// @param {string} query The query string to execute.
// @param {[]interfaces{}} fields Contain the list of type of queryied data. ex. string, date, int, float.
// @param {[]interfaces{}} parameters Contain filter expression, ex. id=0, id != 3.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) Update(storeName string, query string, fields []interface{}, parameters []interface{}, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	err := this.updateData(storeName, query, fields, parameters)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err))
	}
}

// @api 1.0
// Delete db value.
// @param {string} storeName The data server connection (configuration) id
// @param {string} query The query string to execute.
// @param {[]interface{}} parameters Contain filter expression, ex. id=0, id != 3.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) Delete(storeName string, query string, parameters []interface{}, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	err := this.deleteData(storeName, query, parameters)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err))
	}
}

// @api 1.0
// Determine if the datastore exist in the server.
// @param {string} storeName The data server connection (configuration) id
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) HasDataStore(storeName string, messageId string, sessionId string) bool {
	storeUuid := ConfigDataStoreConfigurationExists(storeName)
	return len(storeUuid) > 0
}

// @api 1.0
// Create a dataStore
// @param {string} storeId The id of the datastore to create (can also be see as connection id)
// @param {string} storeName The name of the dataStore to create
// @param {string} hostName The name host where the store is
// @param {string} ipv4 The ip address of the host where the store is.
// @param {int} port The port where the server listen at.
// @param {int} storeType The type of the store to create. SQL: 1; KEY_VALUE: 3
// @param {int} storeVendor The store vendor. DataStoreVendor_MYCELIUS: 1; 	DataStoreVendor_MYSQL: 2; DataStoreVendor_MSSQL:3; DataStoreVendor_ODBC
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) CreateDataStore(storeId string, storeName string, hostName string, ipv4 string, port int64, storeType int64, storeVendor int64, messageId string, sessionId string) {
	if storeType == 0 {
		storeType = 1
	}

	if storeVendor == 0 {
		storeVendor = 1
	}

	var errObj *CargoEntities.Error
	errObj = GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}
	var store DataStore
	store, errObj = this.createDataStore(storeId, storeName, hostName, ipv4, int(port), Config.DataStoreType(storeType), Config.DataStoreVendor(storeVendor))

	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	store.Connect()
}

// @api 1.0
// Delete a dataStore
// @param {string} storeId The id of the dataStore to delete
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) DeleteDataStore(storeId string, messageId string, sessionId string) {
	var errObj *CargoEntities.Error
	errObj = GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	errObj = this.deleteDataStore(storeId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

// @api 1.0
// Synchronize sql datastore with it sql data source.
// @param {string} storeId The id of the dataStore to synchronize
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) Synchronize(storeId string, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	store := this.getDataStore(storeId)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("The datastore '"+storeId+"' does not exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	err := store.Connect()

	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Fail to open the data store connection "+err.Error()+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	// I will get it entity prototypes.
	var prototypes []*EntityPrototype
	prototypes, err = store.GetEntityPrototypes()
	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Fail to get prototypes for store "+storeId+" error: "+err.Error()+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	if reflect.TypeOf(store).String() == "*Server.SqlDataStore" {
		store.(*SqlDataStore).synchronize(prototypes)
	}
}
