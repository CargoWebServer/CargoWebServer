package Server

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"regexp"
	"strconv"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/CargoWebServer/Cargo/JS"
	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/ast"
	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/lexer"
	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/parser"
	"code.myceliUs.com/Utility"
	"github.com/xrash/smetrics"

	// Xapian datastore.
	base64 "encoding/base64"

	"sync"

	"code.myceliUs.com/GoXapian"
	"code.myceliUs.com/XML_Schemas"
)

////////////////////////////////////////////////////////////////////////////////
//                              DataStore function
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
//			Key value Data Store
////////////////////////////////////////////////////////////////////////////////
type GraphStore struct {
	/** The store name **/
	m_id string

	// In case of remote store
	m_conn *WebSocketConnection

	m_port int

	m_ipv4 string

	m_hostName string

	m_storeName string

	m_pwd string

	m_user string

	m_prototypes map[string]*EntityPrototype

	// In case of local sotre
	/** The store path **/
	m_path string

	// Operation channels
	m_set_entity_channel    chan interface{}
	m_get_entity_channel    chan interface{}
	m_delete_entity_channel chan interface{}
}

func getServiceContainerConnection() *WebSocketConnection {
	var conn *WebSocketConnection
	var port int
	port = 9494 // Try to get it from the db...
	conn = GetServer().getConnectionByIp("127.0.0.1", port)
	return conn
}

func NewGraphStore(info *Config.DataStoreConfiguration) (store *GraphStore, err error) {
	store = new(GraphStore)
	store.m_id = info.M_id

	// Connection information.
	store.m_ipv4 = info.M_ipv4
	store.m_port = info.M_port
	store.m_user = info.M_user
	store.m_pwd = info.M_pwd
	store.m_hostName = info.M_hostName
	store.m_storeName = info.M_storeName
	store.m_prototypes = make(map[string]*EntityPrototype, 0)

	// if the store is a local store.
	if store.m_ipv4 == "127.0.0.1" {
		store.m_path = GetServer().GetConfigurationManager().GetDataPath() + "/" + store.m_id
		if _, err := os.Stat(store.m_path); os.IsNotExist(err) {
			os.Mkdir(store.m_path, 0777)
		}
	}

	if err != nil {
		log.Println("open:", err)
	}

	// Here I will register all class in the vm.
	prototypes, err := store.GetEntityPrototypes()
	if err == nil {
		for i := 0; i < len(prototypes); i++ {
			// The script will be put in global context (CargoWebServer)
			JS.GetJsRuntimeManager().AppendScript("CargoWebServer/"+prototypes[i].TypeName, prototypes[i].generateConstructor(), false)
		}
	}

	// Open a db channel.
	store.m_set_entity_channel = make(chan interface{})
	store.m_get_entity_channel = make(chan interface{})
	store.m_delete_entity_channel = make(chan interface{})

	// datastore operation presscessing function.
	go func(store *GraphStore) {

		// The readable datastore.
		wstores := make(map[string]xapian.WritableDatabase)

		// Keep store in memory.
		for {
			select {
			case op := <-store.m_set_entity_channel:
				values, _ := op.(*sync.Map).Load("values")
				if values.(map[string]interface{})["TYPENAME"] != nil {
					//log.Println("---> save values ", toJsonStr(values))
					typeName := values.(map[string]interface{})["TYPENAME"].(string)
					uuid := values.(map[string]interface{})["UUID"].(string)
					path := store.m_path + "/" + typeName + ".glass"

					if wstores[path] == nil {
						wstores[path] = xapian.NewWritableDatabase(path, xapian.DB_CREATE_OR_OPEN)
					}

					wstores[path].Begin_transaction()

					// So here I will index the property found in the entity.
					doc := xapian.NewDocument()

					// Keep json data...
					var data string
					data, err = Utility.ToJson(values)

					if err == nil && len(data) > 0 {
						store.indexEntity(doc, values.(map[string]interface{}))
						doc.Set_data(data)
						doc.Add_boolean_term("Q" + formalize(uuid))
						wstores[path].Replace_document("Q"+formalize(uuid), doc)
					}

					wstores[path].Commit_transaction()

					// Release the document memory.
					xapian.DeleteDocument(doc)
				}
				done, _ := op.(*sync.Map).Load("done")
				done.(chan bool) <- true

			case op := <-store.m_delete_entity_channel:
				uuid := op.(map[string]interface{})["uuid"].(string)
				query := xapian.NewQuery("Q" + formalize(uuid))
				path := store.m_path + "/" + uuid[0:strings.Index(uuid, "%")] + ".glass"

				if wstores[path] == nil {
					wstores[path] = xapian.NewWritableDatabase(path, xapian.DB_CREATE_OR_OPEN)
				}

				enquire := xapian.NewEnquire(wstores[path])
				enquire.Set_query(query)

				mset := enquire.Get_mset(uint(0), uint(10000))

				// Now I will process the results.
				for i := 0; i < mset.Size(); i++ {
					it := mset.Get_hit(uint(i))
					doc := it.Get_document()
					// Remove the document
					wstores[path].Delete_document(doc.Get_docid())
					xapian.DeleteDocument(doc)
					xapian.DeleteMSetIterator(it)
				}

				xapian.DeleteQuery(query)
				xapian.DeleteMSet(mset)
				xapian.DeleteEnquire(enquire)

				wstores[path].Flush()
				op.(map[string]interface{})["done"].(chan bool) <- true

			case op := <-store.m_get_entity_channel:

				queryString := op.(map[string]interface{})["queryString"].(string)
				typeName := op.(map[string]interface{})["typeName"].(string)
				fields := op.(map[string]interface{})["fields"].([]string)
				path := store.m_path + "/" + typeName + ".glass"

				// The results.
				results := make([][]interface{}, 0)
				if !Utility.Exists(path) {
					// Here no database was found.
					err = errors.New("Datastore " + path + " dosent exit!")
				} else {
					if len(queryString) > 0 {
						results, err = store.executeSearchQuery(queryString, fields)
					} else {
						typeNameIndex := generatePrefix(typeName, "TYPENAME") + formalize(typeName)
						query := xapian.NewQuery(typeNameIndex)
						results, err = store.runXapianQuery(typeName, query, fields)
					}
					if len(results) == 0 {
						err = errors.New("No results found!")
					}
				}
				op.(map[string]interface{})["results"].(chan []interface{}) <- []interface{}{results, err}
			}
		}

	}(store)

	return
}

//////////////////////////////////////////////////////////////////////////////////
// Synchronized operations.
//////////////////////////////////////////////////////////////////////////////////

/**
 * Create or Save entity in it store.
 */
func (this *GraphStore) setEntity(entity Entity) {
	var values map[string]interface{}

	if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
		values = entity.(*DynamicEntity).getValues()
	} else {
		values, _ = Utility.ToMap(entity)
	}

	op := new(sync.Map)
	op.Store("values", values)
	done := make(chan bool)
	op.Store("done", done)

	this.m_set_entity_channel <- op

	<-done
}

/**
 * Remove an entity from the datastore.
 */
func (this *GraphStore) deleteEntity(uuid string) {
	op := make(map[string]interface{})
	op["uuid"] = uuid
	op["done"] = make(chan bool)

	this.m_delete_entity_channel <- op

	<-op["done"].(chan bool)
}

/**
 * Get entity or values from a datastore.
 */
func (this *GraphStore) getValues(queryString string, typeName string, fields []string) ([][]interface{}, error) {
	op := make(map[string]interface{})
	op["queryString"] = queryString
	op["typeName"] = typeName
	op["fields"] = fields
	op["results"] = make(chan []interface{})

	this.m_get_entity_channel <- op

	results := <-op["results"].(chan []interface{})

	if results[1] != nil {
		return nil, results[1].(error)
	}

	return results[0].([][]interface{}), nil
}

/**
 * This function is use to create a new entity prototype and save it value.
 * in db.
 * It must be create once per type
 */
func (this *GraphStore) CreateEntityPrototype(prototype *EntityPrototype) error {

	if len(prototype.TypeName) == 0 {
		return errors.New("Entity prototype type name must contain a value!")
	}

	if this.m_ipv4 != "127.0.0.1" {
		// I will use execute JS function to get the list of entity prototypes.
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		param0 := new(MessageData)
		param0.TYPENAME = "Server.MessageData"
		param0.Name = "functionSrc"
		param0.Value = `function GetEntityPrototype(storeId, prototype){ return GetServer().GetEntityManager().CreateEntityPrototype(storeId, prototype, sessionId, messageId) }`

		param1 := new(MessageData)
		param1.TYPENAME = "Server.MessageData"
		param1.Name = "storeId"
		param1.Value = this.m_id

		param2 := new(MessageData)
		param2.TYPENAME = "Server.MessageData"
		param2.Name = "prototype"
		param2.Value = prototype

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)
		params = append(params, param2)

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				// So here I will marchal the values from a json string and
				// initialyse the entity values from the values the contain.
				var results []map[string]interface{}
				json.Unmarshal(rspMsg.msg.Rsp.Results[0].DataBytes, &results)

				// Set the TYPENAME property here.
				results[0]["TYPENAME"] = "Server.EntityPrototype"
				value, err := Utility.InitializeStructure(results[0], setEntityFct)
				if err != nil {
					resultsChan <- err
				} else {
					resultsChan <- value.Interface().(*EntityPrototype)
				}
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan
		if reflect.TypeOf(results).String() == "*Server.EntityPrototype" {
			return nil
		}

		return results.(error) // return an error message instead.
	}

	// Here i will append super type fields...
	prototype.setSuperTypeFields()

	// Register it to the vm...
	JS.GetJsRuntimeManager().AppendScript("CargoWebServer", prototype.generateConstructor(), true)

	// Send event message...
	var eventDatas []*MessageData
	evtData := new(MessageData)
	evtData.TYPENAME = "Server.MessageData"
	evtData.Name = "prototype"

	evtData.Value = prototype
	eventDatas = append(eventDatas, evtData)
	evt, _ := NewEvent(NewPrototypeEvent, PrototypeEvent, eventDatas)
	GetServer().GetEventManager().BroadcastEvent(evt)
	if len(prototype.TypeName) == 0 {
		return errors.New("Entity prototype type name must contain a value!")
	}
	// I will serialyse the prototype.
	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	err := enc.Encode(prototype)

	if err != nil {
		log.Println("Prototype encode:", err)
		return err
	}

	if len(prototype.TypeName) == 0 {
		// The typeName cant be nil!
		panic(prototype)
	}

	// I will save the entity prototype in a file.
	if strings.HasPrefix(prototype.TypeName, this.GetId()) {
		file, err := os.Create(this.m_path + "/" + prototype.TypeName + ".gob")
		defer file.Close()

		if err == nil {
			encoder := gob.NewEncoder(file)
			encoder.Encode(prototype)
		} else {
			return err
		}
		this.m_prototypes[prototype.TypeName] = prototype
	}

	return nil
}

/**
 * Save an entity prototype.
 */
func (this *GraphStore) SaveEntityPrototype(prototype *EntityPrototype) error {
	if len(prototype.TypeName) == 0 {
		return errors.New("Entity prototype type name must contain a value!")
	}
	if this.m_ipv4 != "127.0.0.1" {
		// I will use execute JS function to get the list of entity prototypes.
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		param0 := new(MessageData)
		param0.TYPENAME = "Server.MessageData"
		param0.Name = "functionSrc"
		param0.Value = `function SaveEntityPrototype(storeId, prototype){ return GetServer().GetEntityManager().SaveEntityPrototype(storeId, prototype, sessionId, messageId) }`

		param1 := new(MessageData)
		param1.TYPENAME = "Server.MessageData"
		param1.Name = "storeId"
		param1.Value = this.m_id

		param2 := new(MessageData)
		param2.TYPENAME = "Server.MessageData"
		param2.Name = "prototype"
		param2.Value = prototype

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)
		params = append(params, param2)

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				// So here I will marchal the values from a json string and
				// initialyse the entity values from the values the contain.
				var results []map[string]interface{}
				json.Unmarshal(rspMsg.msg.Rsp.Results[0].DataBytes, &results)

				// Set the TYPENAME property here.
				results[0]["TYPENAME"] = "Server.EntityPrototype"
				value, err := Utility.InitializeStructure(results[0], setEntityFct)
				if err != nil {
					resultsChan <- err
				} else {
					resultsChan <- value.Interface().(*EntityPrototype)
				}
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan
		if reflect.TypeOf(results).String() == "*Server.EntityPrototype" {
			return nil
		}

		return results.(error) // return an error message instead.
	}

	// Get the current entity prototype.
	prototype_, err := GetServer().GetEntityManager().getEntityPrototype(prototype.TypeName, this.m_id)
	if err != nil {
		return err
	}

	// I will serialyse the prototype.
	prototype.setSuperTypeFields()

	// I will remove it from substitution group as neeeded...
	for i := 0; i < len(prototype_.SuperTypeNames); i++ {
		if !Utility.Contains(prototype.SuperTypeNames, prototype_.SuperTypeNames[i]) {
			// Here I will remove the prototype from superType substitution group.
			superTypeName := prototype_.SuperTypeNames[i]
			superType, err := GetServer().GetEntityManager().getEntityPrototype(superTypeName, superTypeName[0:strings.Index(superTypeName, ".")])
			if err != nil {
				return err
			}

			substitutionGroup := make([]string, 0)
			for j := 0; j < len(superType.SubstitutionGroup); j++ {
				if superType.SubstitutionGroup[j] != prototype_.TypeName {
					substitutionGroup = append(substitutionGroup, superType.SubstitutionGroup[j])
				}
			}
			superType.SubstitutionGroup = substitutionGroup
			store := GetServer().GetDataManager().getDataStore(superTypeName[0:strings.Index(superTypeName, ".")])
			err = store.SaveEntityPrototype(superType)

			if err != nil {
				return err
			}
		}
	}

	// Register it to the vm...
	JS.GetJsRuntimeManager().AppendScript("CargoWebServer/"+prototype.TypeName, prototype.generateConstructor(), true)

	file, err := os.Create(this.m_path + "/" + prototype.TypeName + ".gob")
	defer file.Close()

	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(prototype)
	} else {
		return err
	}

	this.m_prototypes[prototype.TypeName] = prototype

	var eventDatas []*MessageData
	evtData := new(MessageData)
	evtData.TYPENAME = "Server.MessageData"
	evtData.Name = "prototype"

	evtData.Value = prototype
	eventDatas = append(eventDatas, evtData)
	evt, _ := NewEvent(UpdatePrototypeEvent, PrototypeEvent, eventDatas)
	GetServer().GetEventManager().BroadcastEvent(evt)

	return nil
}

/**
 * Remove an entity prototype and all it releated values.
 */
func (this *GraphStore) DeleteEntityPrototype(typeName string) error {
	// In case of remote data store.
	if this.m_ipv4 != "127.0.0.1" {
		// I will use execute JS function to get the list of entity prototypes.
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		param0 := new(MessageData)
		param0.TYPENAME = "Server.MessageData"
		param0.Name = "functionSrc"
		param0.Value = `function DeleteEntityPrototype(typeName, storeId){ GetServer().GetEntityManager().DeleteEntityPrototype(typeName, storeId, sessionId, messageId) }`

		param1 := new(MessageData)
		param1.TYPENAME = "Server.MessageData"
		param1.Name = "typeName"
		param1.Value = typeName

		param2 := new(MessageData)
		param2.TYPENAME = "Server.MessageData"
		param2.Name = "storeId"
		param2.Value = this.m_id

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)
		params = append(params, param2)

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				log.Println("---> entity protoype deleted!")
				// update success
				resultsChan <- nil
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan

		// in case of error
		if results != nil {
			if reflect.TypeOf(results).String() == "*string" {
				return errors.New(*results.(*string))
			}
		}

		return nil
	}

	prototype := this.m_prototypes[typeName]
	// The prototype does not exist.
	if prototype == nil {
		// not exist so no need to be removed...
		return nil
	}

	// Remove substitution group from it parent.
	for i := 0; i < len(prototype.SuperTypeNames); i++ {
		storeId := prototype.SuperTypeNames[i][0:strings.Index(prototype.SuperTypeNames[i], ".")]
		superPrototype, err := GetServer().GetEntityManager().getEntityPrototype(prototype.SuperTypeNames[i], storeId)
		if err == nil {
			substitutionGroup := make([]string, 0)
			for j := 0; j < len(superPrototype.SubstitutionGroup); j++ {
				if superPrototype.SubstitutionGroup[j] != typeName {
					substitutionGroup = append(substitutionGroup, superPrototype.SubstitutionGroup[j])
				}
			}
			// Save the prototype.
			superPrototype.SubstitutionGroup = substitutionGroup
			store := GetServer().GetDataManager().getDataStore(storeId)
			store.SaveEntityPrototype(superPrototype)
		}
	}

	// I will delete all entity...
	entities, _ := GetServer().GetEntityManager().getEntities(prototype.TypeName, this.m_id, nil)
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		// remove it...
		GetServer().GetEntityManager().deleteEntity(entity)
	}

	delete(this.m_prototypes, typeName)

	err := os.Remove(this.m_path + "/" + prototype.TypeName + ".gob")

	return err
}

/**
 * Remove all prototypes.
 */
func (this *GraphStore) DeleteEntityPrototypes() error {
	if this.m_ipv4 == "127.0.0.1" {
		for typeName, prototype := range this.m_prototypes {
			// Remove substitution group from it parent.
			for i := 0; i < len(prototype.SuperTypeNames); i++ {
				storeId := prototype.SuperTypeNames[i][0:strings.Index(prototype.SuperTypeNames[i], ".")]
				if storeId != this.m_id {
					superPrototype, err := GetServer().GetEntityManager().getEntityPrototype(prototype.SuperTypeNames[i], storeId)
					if err == nil {
						substitutionGroup := make([]string, 0)
						for j := 0; j < len(superPrototype.SubstitutionGroup); j++ {
							if superPrototype.SubstitutionGroup[j] != typeName {
								substitutionGroup = append(substitutionGroup, superPrototype.SubstitutionGroup[j])
							}
						}
						// Save the prototype.
						superPrototype.SubstitutionGroup = substitutionGroup
						store := GetServer().GetDataManager().getDataStore(storeId)
						store.SaveEntityPrototype(superPrototype)
					}
				}
			}

			// Remove the entity from the cache and send delete event.
			entities, _ := GetServer().GetEntityManager().getEntities(typeName, this.m_id, nil)
			for i := 0; i < len(entities); i++ {
				entity := entities[i]
				// remove it from the cache...
				if len(entity.GetParentUuid()) > 0 {
					if !strings.HasPrefix(entity.GetParentUuid(), this.m_id) {
						// I will get the parent uuid link.
						parent, err := GetServer().GetEntityManager().getEntityByUuid(entity.GetParentUuid())
						if err != nil {
							return errors.New(err.GetBody())
						}

						// Here I will remove it from it parent...
						// Get values as map[string]interface{} and also set the entity in it parent.
						if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
							parent.(*DynamicEntity).removeValue(entity.GetParentLnk(), entity.GetUuid())
						} else {
							removeMethode := strings.Replace(entity.GetParentLnk(), "M_", "", -1)
							removeMethode = "Remove" + strings.ToUpper(removeMethode[0:1]) + removeMethode[1:]
							params := make([]interface{}, 1)
							params[0] = entity
							_, err_ := Utility.CallMethod(parent, removeMethode, params)
							if err_ != nil {
								cargoError := NewError(Utility.FileLine(), ATTRIBUTE_NAME_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, err_.(error))
								return errors.New(cargoError.GetBody())
							}
						}

						// Update the parent here.
						var eventDatas []*MessageData
						evtData := new(MessageData)
						evtData.TYPENAME = "Server.MessageData"
						evtData.Name = "entity"
						if reflect.TypeOf(parent).String() == "*Server.DynamicEntity" {
							evtData.Value = parent.(*DynamicEntity).getValues()
						} else {
							evtData.Value = parent
						}
						eventDatas = append(eventDatas, evtData)
						evt, _ := NewEvent(UpdateEntityEvent, EntityEvent, eventDatas)
						GetServer().GetEventManager().BroadcastEvent(evt)
					}
				}

				GetServer().GetEntityManager().removeEntity(entity)

				// Send event message...
				var eventDatas []*MessageData
				evtData := new(MessageData)
				evtData.TYPENAME = "Server.MessageData"
				evtData.Name = "entity"
				if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
					evtData.Value = entity.(*DynamicEntity).getValues()
				} else {
					evtData.Value = entity
				}

				eventDatas = append(eventDatas, evtData)
				evt, _ := NewEvent(DeleteEntityEvent, EntityEvent, eventDatas)
				GetServer().GetEventManager().BroadcastEvent(evt)

			}
		}

		// Remove all prototypes from the map.
		for typeName, _ := range this.m_prototypes {
			delete(this.m_prototypes, typeName)
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
//                              DataStore function
////////////////////////////////////////////////////////////////////////////////

/**
 * This function is use to retreive an existing entity prototype...
 */
func (this *GraphStore) GetEntityPrototype(typeName string) (*EntityPrototype, error) {
	if len(typeName) == 0 {
		return nil, errors.New("Entity prototype type name must contain a value!")
	}

	// Here the store is not a local, so I will use a remote call to get the
	// list of it entity prototypes.
	if this.m_ipv4 != "127.0.0.1" {

		// I will use execute JS function to get the list of entity prototypes.
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		param0 := new(MessageData)
		param0.TYPENAME = "Server.MessageData"
		param0.Name = "functionSrc"
		param0.Value = `function GetEntityPrototype(typeName, storeId){ return GetServer().GetEntityManager().GetEntityPrototype(typeName, storeId, sessionId, messageId) }`

		param1 := new(MessageData)
		param1.TYPENAME = "Server.MessageData"
		param1.Name = "typeName"
		param1.Value = typeName

		param2 := new(MessageData)
		param2.TYPENAME = "Server.MessageData"
		param2.Name = "storeId"
		param2.Value = this.m_id

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)
		params = append(params, param2)

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				// So here I will marchal the values from a json string and
				// initialyse the entity values from the values the contain.
				var results []map[string]interface{}
				json.Unmarshal(rspMsg.msg.Rsp.Results[0].DataBytes, &results)

				// Set the TYPENAME property here.
				results[0]["TYPENAME"] = "Server.EntityPrototype"
				value, err := Utility.InitializeStructure(results[0], setEntityFct)
				if err != nil {
					resultsChan <- err
				} else {
					resultsChan <- value.Interface().(*EntityPrototype)
				}
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan
		if reflect.TypeOf(results).String() == "*Server.EntityPrototype" {
			return results.(*EntityPrototype), nil
		}

		return nil, results.(error) // return an error message instead.
	}

	if this.m_prototypes[typeName] != nil {
		return this.m_prototypes[typeName], nil
	} else {
		// Local store stuff...
		var prototype *EntityPrototype
		prototype = new(EntityPrototype)
		file, err := os.Open(this.m_path + "/" + typeName + ".gob")
		defer file.Close()
		if err == nil {
			decoder := gob.NewDecoder(file)
			err = decoder.Decode(prototype)
		} else {
			file, err = os.Open(this.m_path + "/" + typeName + "_impl.gob")
			if err == nil {
				decoder := gob.NewDecoder(file)
				err = decoder.Decode(prototype)
			}
		}
		if err != nil {
			//log.Panicln("---> ", typeName, err)
			return nil, err
		}

		this.m_prototypes[typeName] = prototype

		return prototype, err
	}

}

/**
 * Retreive the list of all entity prototype in a given store.
 */
func (this *GraphStore) GetEntityPrototypes() ([]*EntityPrototype, error) {

	var prototypes []*EntityPrototype
	// Here the store is not a local, so I will use a remote call to get the
	// list of it entity prototypes.
	if this.m_ipv4 == "" {
		this.m_ipv4 = "127.0.0.1"
	}

	if this.m_ipv4 != "127.0.0.1" {
		if !this.m_conn.IsOpen() {
			err := this.Connect()
			if err != nil {
				return nil, err
			}
		}
		// I will use execute JS function to get the list of entity prototypes.
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		param0 := new(MessageData)
		param0.TYPENAME = "Server.MessageData"
		param0.Name = "functionSrc"
		param0.Value = `function GetEntityPrototypes(storeId){ return GetServer().GetEntityManager().GetEntityPrototypes(storeId, sessionId, messageId) }`

		param1 := new(MessageData)
		param1.TYPENAME = "Server.MessageData"
		param1.Name = "storeId"
		param1.Value = this.m_id

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				// So here I will marchal the values from a json string and
				// initialyse the entity values from the values the contain.
				var results [][]map[string]interface{}
				var prototypes []*EntityPrototype
				json.Unmarshal(rspMsg.msg.Rsp.Results[0].DataBytes, &results)
				for i := 0; i < len(results[0]); i++ {
					// Set the TYPENAME property here.
					results[0][i]["TYPENAME"] = "Server.EntityPrototype"
					values, err := Utility.InitializeStructure(results[0][i], setEntityFct)
					if err == nil {
						prototypes = append(prototypes, values.Interface().(*EntityPrototype))
					}
				}
				resultsChan <- prototypes
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan
		if reflect.TypeOf(results).String() == "[]*Server.EntityPrototype" {
			return results.([]*EntityPrototype), nil
		}

		return prototypes, errors.New(*results.(*string)) // return an error message instead.
	}

	// Get prototypes from files.
	files, err := ioutil.ReadDir(this.m_path)
	if err != nil {
		return nil, err
	}

	for _, info := range files {
		if strings.HasSuffix(info.Name(), ".gob") {
			if err == nil {
				prototype, err := this.GetEntityPrototype(strings.Split(info.Name(), ".gob")[0])
				if err == nil {
					prototypes = append(prototypes, prototype)
				}
			}
		}
	}

	return prototypes, nil
}

/**
 * Return the name of a store.
 */
func (this *GraphStore) GetId() string {
	return this.m_id
}

// TODO validate the user and password here...
func (this *GraphStore) Connect() error {

	if this.m_ipv4 != "127.0.0.1" {
		// I will not try to connect if a connection already exist.
		if this.m_conn != nil {
			if this.m_conn.IsOpen() {
				return nil
			}
		}

		// Here I will connect to a remote server.
		var err error
		this.m_conn, err = GetServer().connect(this.m_ipv4, this.m_port)

		if err != nil {
			return err
		}

		// Here I will use the user and password in the connection to validate
		// that the user can get data from the store.

		// I will use execute JS function to get the list of entity prototypes.
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		param0 := new(MessageData)
		param0.TYPENAME = "Server.MessageData"
		param0.Name = "functionSrc"
		param0.Value = `function Login(accountName, psswd, serverId){ return GetServer().GetSessionManager().Login(accountName, psswd, serverId, sessionId, messageId) }`

		param1 := new(MessageData)
		param1.TYPENAME = "Server.MessageData"
		param1.Name = "accountName"
		param1.Value = this.m_user

		param2 := new(MessageData)
		param2.TYPENAME = "Server.MessageData"
		param2.Name = "psswd"
		param2.Value = this.m_pwd

		param3 := new(MessageData)
		param3.TYPENAME = "Server.MessageData"
		param3.Name = "serverId"
		param3.Value = this.m_hostName

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)
		params = append(params, param2)
		params = append(params, param3)

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				// So here I will marchal the values from a json string and
				// initialyse the entity values from the values the contain.

				var results []map[string]interface{}
				json.Unmarshal(rspMsg.msg.Rsp.Results[0].DataBytes, &results)
				if results[0] == nil {
					resultsChan <- "Fail to open session!"
					return
				}
				results[0]["TYPENAME"] = "CargoEntities.Session"
				values, err := Utility.InitializeStructure(results[0], setEntityFct)

				if err == nil {
					resultsChan <- values.Interface().(*CargoEntities.Session)
				} else {
					resultsChan <- err.Error() // send the error instead...
				}
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan
		if reflect.TypeOf(results).String() != "*CargoEntities.Session" {
			return errors.New(*results.(*string)) // return an error message instead.
		}
	}

	return nil
}

/**
 * Help to know if a store is connect or existing...
 */
func (this *GraphStore) Ping() error {
	if this.m_ipv4 != "127.0.0.1" {
		if this.m_conn != nil {
			if !this.m_conn.IsOpen() {
				err := this.Connect()
				if err != nil {
					return err
				}
			}
		}

		// Call ping on the distant server.
		id := Utility.RandomUUID()
		method := "Ping"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				// So here I will marchal the values from a json string and
				// initialyse the entity values from the values the contain.
				resultsChan <- string(rspMsg.msg.Rsp.Results[0].DataBytes)
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan
		if reflect.TypeOf(results).String() != "string" {
			return errors.New(*results.(*string)) // return an error message instead.
		}

		return nil
	}

	// Local store ping...
	path := GetServer().GetConfigurationManager().GetDataPath() + "/" + this.GetId()
	_, err := os.Stat(path)
	return err
}

/**
 * Create a new entry in the database.
 */
func (this *GraphStore) Create(queryStr string, values []interface{}) (lastId interface{}, err error) {
	if this.m_ipv4 != "127.0.0.1" {
		if this.m_conn != nil {
			if !this.m_conn.IsOpen() {
				err := this.Connect()
				if err != nil {
					return nil, err
				}
			}
		}

		// I will use execute JS function to get the list of entity prototypes.
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		param0 := new(MessageData)
		param0.TYPENAME = "Server.MessageData"
		param0.Name = "functionSrc"
		param0.Value = `function CreateData(storeId, query, data){ return GetServer().GetDataManager().Create(storeId, query, data, sessionId, messageId) }`

		param1 := new(MessageData)
		param1.TYPENAME = "Server.MessageData"
		param1.Name = "storeId"
		param1.Value = this.m_id

		param2 := new(MessageData)
		param2.TYPENAME = "Server.MessageData"
		param2.Name = "query"
		param2.Value = queryStr

		param3 := new(MessageData)
		param3.TYPENAME = "Server.MessageData"
		param3.Name = "data"
		param3.Value = values

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)
		params = append(params, param2)
		params = append(params, param3)

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				// So here I will marchal the values from a json string and
				resultsChan <- string(rspMsg.msg.Rsp.Results[0].DataBytes) // Return the last created id if there is some.
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan

		// in case of error
		if reflect.TypeOf(results).String() == "*string" {
			return -1, errors.New(*results.(*string))
		}

		return results, nil
	}

	// Creation of entity
	for i := 0; i < len(values); i++ {
		v := values[i]
		// If the value is a dynamic entity...
		if reflect.TypeOf(v).Kind() == reflect.Ptr || reflect.TypeOf(v).Kind() == reflect.Struct || reflect.TypeOf(v).Kind() == reflect.Map {
			this.setEntity(v.(Entity))
		}
	}

	return
}

/**
 * Get the value list...
 */
func (this *GraphStore) Read(queryStr string, fieldsType []interface{}, params []interface{}) (results [][]interface{}, err error) {
	if this.m_ipv4 != "127.0.0.1" {
		if this.m_conn != nil {
			if !this.m_conn.IsOpen() {
				err := this.Connect()
				if err != nil {
					return nil, err
				}
			}
		}
		// I will use execute JS function to get the list of entity prototypes.
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		param0 := new(MessageData)
		param0.TYPENAME = "Server.MessageData"
		param0.Name = "functionSrc"
		param0.Value = `function ReadData(storeId, query, fieldsType, parameters){ return GetServer().GetDataManager().Read(storeId, query, fieldsType, parameters, sessionId, messageId) }`

		param1 := new(MessageData)
		param1.TYPENAME = "Server.MessageData"
		param1.Name = "storeId"
		param1.Value = this.m_id

		param2 := new(MessageData)
		param2.TYPENAME = "Server.MessageData"
		param2.Name = "query"
		param2.Value = queryStr

		param3 := new(MessageData)
		param3.TYPENAME = "Server.MessageData"
		param3.Name = "fieldsType"
		param3.Value = fieldsType

		param4 := new(MessageData)
		param4.TYPENAME = "Server.MessageData"
		param4.Name = "parameters"
		param4.Value = params

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)
		params = append(params, param2)
		params = append(params, param3)
		params = append(params, param4)

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				// So here I will marchal the values from a json string and
				// initialyse the entity values from the values the contain.
				var results [][][]interface{} // Tree dimension array of values
				err := json.Unmarshal(rspMsg.msg.Rsp.Results[0].DataBytes, &results)
				if err != nil {
					resultsChan <- err
					return
				}
				resultsChan <- results[0] // the first element contain the results.
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan

		// in case of error
		if reflect.TypeOf(results).String() == "error" {
			return nil, results.(error) // return an error message instead.
		} else if reflect.TypeOf(results).String() == "*string" {
			return nil, errors.New(*results.(*string))
		}

		return results.([][]interface{}), nil
	}

	// First of all i will init the query...
	var query EntityQuery
	err = json.Unmarshal([]byte(queryStr), &query)
	if err != nil {
		return nil, err
	}

	results, err = this.getValues(query.Query, query.TypeName, query.Fields)

	return
}

/**
 * Update a entity value.
 * TODO think about a cute way to modify part of the entity and not the whole thing...
 */
func (this *GraphStore) Update(queryStr string, values []interface{}, params []interface{}) (err error) {
	// Remote server.
	if this.m_ipv4 != "127.0.0.1" {
		if this.m_conn != nil {
			if !this.m_conn.IsOpen() {
				err := this.Connect()
				if err != nil {
					return err
				}
			}
		}
		// I will use execute JS function to get the list of entity prototypes.
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		param0 := new(MessageData)
		param0.TYPENAME = "Server.MessageData"
		param0.Name = "functionSrc"
		param0.Value = `function UpdateData(storeId, query, fields, parameters){ return GetServer().GetDataManager().Update(storeId, query, fields, parameters, sessionId, messageId) }`

		param1 := new(MessageData)
		param1.TYPENAME = "Server.MessageData"
		param1.Name = "storeId"
		param1.Value = this.m_id

		param2 := new(MessageData)
		param2.TYPENAME = "Server.MessageData"
		param2.Name = "query"
		param2.Value = queryStr

		param3 := new(MessageData)
		param3.TYPENAME = "Server.MessageData"
		param3.Name = "fields"
		param3.Value = values

		param4 := new(MessageData)
		param4.TYPENAME = "Server.MessageData"
		param4.Name = "parameters"
		param4.Value = params

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)
		params = append(params, param2)
		params = append(params, param3)
		params = append(params, param4)

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				// update success
				resultsChan <- nil
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan

		// in case of error
		if results != nil {
			if reflect.TypeOf(results).String() == "*string" {
				return errors.New(*results.(*string))
			}
		}

		return nil
	}

	// The value to be save.
	for i := 0; i < len(values); i++ {
		v := values[i]
		// If the value is an entity...
		if reflect.TypeOf(v).Kind() == reflect.Ptr || reflect.TypeOf(v).Kind() == reflect.Struct || reflect.TypeOf(v).Kind() == reflect.Map {
			this.setEntity(v.(Entity))
		}
	}

	return
}

/**
 * Delete entity from the store...
 */
func (this *GraphStore) Delete(queryStr string, values []interface{}) (err error) {
	// Remote server.
	if this.m_ipv4 != "127.0.0.1" {
		if this.m_conn != nil {
			if !this.m_conn.IsOpen() {
				err := this.Connect()
				if err != nil {
					return err
				}
			}
		}
		// I will use execute JS function to get the list of entity prototypes.
		id := Utility.RandomUUID()
		method := "ExecuteJsFunction"
		params := make([]*MessageData, 0)

		to := make([]*WebSocketConnection, 1)
		to[0] = this.m_conn

		param0 := new(MessageData)
		param0.TYPENAME = "Server.MessageData"
		param0.Name = "functionSrc"
		param0.Value = `function UpdateData(storeId, query, parameters){ return GetServer().GetDataManager().Delete(storeId, query, parameters, sessionId, messageId) }`

		param1 := new(MessageData)
		param1.TYPENAME = "Server.MessageData"
		param1.Name = "storeId"
		param1.Value = this.m_id

		param2 := new(MessageData)
		param2.TYPENAME = "Server.MessageData"
		param2.Name = "query"
		param2.Value = queryStr

		param3 := new(MessageData)
		param3.TYPENAME = "Server.MessageData"
		param3.Name = "parameters"
		param3.Value = params

		// Append the params.
		params = append(params, param0)
		params = append(params, param1)
		params = append(params, param2)
		params = append(params, param3)

		// The channel will be use to wait for results.
		resultsChan := make(chan interface{})

		// The success callback.
		successCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(rspMsg *message, caller interface{}) {
				// update success
				resultsChan <- nil
			}
		}(resultsChan)

		// The error callback.
		errorCallback := func(resultsChan chan interface{}) func(*message, interface{}) {
			return func(errMsg *message, caller interface{}) {
				resultsChan <- errMsg.msg.Err.Message
			}
		}(resultsChan)

		rqst, _ := NewRequestMessage(id, method, params, to, successCallback, nil, errorCallback, nil)

		go func(rqst *message) {
			GetServer().getProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan

		// in case of error
		if results != nil {
			if reflect.TypeOf(results).String() == "*string" {
				return errors.New(*results.(*string))
			}
		}

		return nil
	}

	// Remove the list of obsolete triples from the datastore.
	for i := 0; i < len(values); i++ {
		toDelete := values[i]
		if reflect.TypeOf(toDelete).Kind() == reflect.String {
			if Utility.IsValidEntityReferenceName(toDelete.(string)) {
				this.deleteEntity(toDelete.(string))
			}
		}
	}

	return
}

/**
 * Close the backend store.
 */
func (this *GraphStore) Close() error {
	// Remote server.
	if this.m_ipv4 != "127.0.0.1" {
		// Close the connection.
		if this.m_conn != nil {
			this.m_conn.Close()
		}
		return nil
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Search functionality.
////////////////////////////////////////////////////////////////////////////////

// Generate prefix is use to create a indexation key for a given document.
// the field must be in the index or id's.
func generatePrefix(typeName string, field string) string {
	// remove the M_ part of the field name
	prefix := typeName

	if len(field) > 0 {
		prefix += "." + field
	}

	// replace unwanted character's
	prefix = strings.Replace(prefix, ".", "_", -1) + "%"
	prefix = "X" + strings.ToLower(prefix)

	return prefix
}

// Remove ambiquous query symbols % - . and replace it with _
func formalize(uuid string) string {
	return strings.TrimSpace(strings.ToLower(Utility.ToString(strings.Replace(strings.Replace(strings.Replace(uuid, "-", "_", -1), ".", "_", -1), "%", "_", -1))))
}

// Index entity string field.
func (this *GraphStore) indexStringField(data string, field string, typeName string, termGenerator xapian.TermGenerator) {
	// I will index all string field to be able to found it back latter.
	termGenerator.Index_text(strings.ToLower(data), uint(1), strings.ToUpper(field))
	if Utility.IsUriBase64(data) {
		data_, err := base64.StdEncoding.DecodeString(data)
		if err == nil {
			if strings.Index(data, ":text/") > -1 || strings.Index(data, ":application/") > -1 {
				termGenerator.Index_text(strings.ToLower(string(data_)))
			}
		}
	} else if Utility.IsStdBase64(data) {
		data_, err := base64.StdEncoding.DecodeString(data)
		if err == nil {
			termGenerator.Index_text(strings.ToLower(string(data_)))
			if field != "M_data" {
				termGenerator.Index_text(strings.ToLower(string(data)))
			}
		}
	} else {
		termGenerator.Index_text(strings.ToLower(data))
	}
}

// Index entity field
func (this *GraphStore) indexField(data interface{}, field string, fieldType string, typeName string, termGenerator xapian.TermGenerator, doc xapian.Document, index int) {
	// This will give possibility to search for given fields.
	if data != nil {
		if reflect.TypeOf(data).Kind() == reflect.Slice {
			str_, err := Utility.ToJson(data)
			if err == nil {
				s := reflect.ValueOf(data)
				for i := 0; i < s.Len(); i++ {
					// I will remove nil values.
					item := s.Index(i)
					if item.IsValid() {
						zeroValue := reflect.Zero(item.Type())
						if zeroValue != item {
							this.indexField(s.Index(i).Interface(), field, fieldType, typeName, termGenerator, doc, -1)
						} else {
							this.indexField(nil, field, fieldType, typeName, termGenerator, doc, -1)
						}

					}
				}
				doc.Add_value(uint(index), Utility.ToString(str_))
			} else {
				doc.Add_value(uint(index), "null")
			}
		} else {
			if index != -1 {
				doc.Add_value(uint(index), Utility.ToString(data))
			}
			if (XML_Schemas.IsXsNumeric(fieldType) || XML_Schemas.IsXsInt(fieldType)) && index != -1 {
				value := Utility.ToNumeric(data)
				doc.Add_value(uint(index), xapian.Sortable_serialise(value))
			} else if reflect.TypeOf(data).Kind() == reflect.String {
				str := Utility.ToString(data)
				// If the the value is a valid entity reference i I will use boolean term.
				if Utility.IsValidEntityReferenceName(str) {
					term := generatePrefix(typeName, field) + formalize(str)
					doc.Add_boolean_term(term)
				} else {
					this.indexStringField(str, field, typeName, termGenerator)
				}
			}
		}
	} else {
		doc.Add_value(uint(index), "null")
	}
}

// index entity information.
func (this *GraphStore) indexEntity(doc xapian.Document, values map[string]interface{}) {

	// The term generator
	termGenerator := xapian.NewTermGenerator()

	// set english by default.
	stemmer := xapian.NewStem("en")

	termGenerator.Set_stemmer(stemmer)
	termGenerator.Set_document(doc)

	// Regular text indexation...
	termGenerator.Index_text(values["TYPENAME"].(string), uint(1), "TYPENAME")

	// Boolean term indexation exact match.
	typeNameIndex := generatePrefix(values["TYPENAME"].(string), "TYPENAME") + formalize(values["TYPENAME"].(string))
	doc.Add_boolean_term(typeNameIndex)

	prototype, _ := this.GetEntityPrototype(values["TYPENAME"].(string))

	// also index value supertype...
	for i := 0; i < len(prototype.SuperTypeNames); i++ {
		termGenerator.Index_text(prototype.SuperTypeNames[i], uint(1), "TYPENAME")
		typeNameIndex := generatePrefix(prototype.SuperTypeNames[i], "TYPENAME") + formalize(prototype.SuperTypeNames[i])
		doc.Add_boolean_term(typeNameIndex)
	}

	// Here I will append boolean term.
	for i := 0; i < len(prototype.Fields); i++ {
		// Index the value.
		var value interface{}
		value = values[prototype.Fields[i]]
		if Utility.Contains(prototype.Ids, prototype.Fields[i]) || Utility.Contains(prototype.Indexs, prototype.Fields[i]) {
			// Index the unique value index for the typeName and this field.
			if value != nil {
				term := generatePrefix(prototype.TypeName, prototype.Fields[i]) + formalize(Utility.ToString(value))
				doc.Add_boolean_term(term)
			}
		}
		index := prototype.getFieldIndex(prototype.Fields[i])
		this.indexField(value, prototype.Fields[i], prototype.FieldsType[i], prototype.TypeName, termGenerator, doc, index)
	}

	xapian.DeleteStem(stemmer)
	xapian.DeleteTermGenerator(termGenerator)
}

////////////////////////////////////////////////////////////////////////////////
// Search functionality.
////////////////////////////////////////////////////////////////////////////////

/**
 * Merge tow results in one...
 */
func (this *GraphStore) merge(r1 map[string]map[string]interface{}, r2 map[string]map[string]interface{}) map[string]map[string]interface{} {

	for k, v := range r1 {
		r2[k] = v
	}
	return r2
}

/**
 * Evaluate an expression.
 */
func (this *GraphStore) evaluate(typeName string, fieldName string, comparator string, expected interface{}, value interface{}) (bool, error) {
	isMatch := false
	// if the value is nil i will automatically return
	if value == nil {
		return isMatch, nil
	}

	prototype, err := this.GetEntityPrototype(typeName)
	if err != nil {
		return false, err
	}

	// The type name.
	fieldType := prototype.FieldsType[prototype.getFieldIndex(fieldName)]
	fieldType = strings.Replace(fieldType, "[]", "", -1)

	// here for the date I will get it unix time value...
	if fieldType == "xs.date" || fieldType == "xs.dateTime" {
		expectedDateValue, err := Utility.MatchISO8601_Date(expected.(string))
		if err == nil {
			dateValue, _ := Utility.MatchISO8601_Date(value.(string))
			if fieldType == "xs.dateTime" {
				expected = expectedDateValue.Unix() // get the unix resultstime for calcul
				value = dateValue.Unix()            // get the unix time for calcul
			} else {
				expected = expectedDateValue.Truncate(24 * time.Hour).Unix() // get the unix time for calcul
				value = dateValue.Truncate(24 * time.Hour).Unix()            // get the unix time for calcul
			}
		} else {
			// I will try with data time instead.
			expectedDateValue, err := Utility.MatchISO8601_DateTime(expected.(string))
			if err == nil {
				dateValue, _ := Utility.MatchISO8601_DateTime(value.(string))
				if fieldType == "xs.dateTime" {
					expected = expectedDateValue.Unix() // get the unix time for calcul
					value = dateValue.Unix()            // get the unix time for calcul
				} else {
					expected = expectedDateValue.Truncate(24 * time.Hour).Unix() // get the unix time for calcul
					value = dateValue.Truncate(24 * time.Hour).Unix()            // get the unix time for calcul
				}
			} else {
				return false, err
			}
		}
	}

	if comparator == "==" {
		// Equality comparator.
		// Case of string type.
		if reflect.TypeOf(expected).Kind() == reflect.String && reflect.TypeOf(value).Kind() == reflect.String {
			isRegex := strings.HasPrefix(expected.(string), "/") && strings.HasSuffix(expected.(string), "/")
			if isRegex {
				// here I will try to match the regular expression.
				var err error
				isMatch, err = regexp.MatchString(expected.(string)[1:len(expected.(string))-1], value.(string))
				if err != nil {
					return false, err
				}
			} else {
				isMatch = Utility.RemoveAccent(expected.(string)) == Utility.RemoveAccent(value.(string))
			}
		} else if reflect.TypeOf(expected).Kind() == reflect.Bool && reflect.TypeOf(value).Kind() == reflect.Bool {
			return expected.(bool) == value.(bool), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return expected.(int64) == value.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return expected.(float64) == value.(float64), nil
		}
	} else if comparator == "~=" {
		// Approximation comparator, string only...
		// Case of string types.
		if reflect.TypeOf(expected).Kind() == reflect.String && reflect.TypeOf(value).Kind() == reflect.String {
			distance := smetrics.JaroWinkler(Utility.RemoveAccent(expected.(string)), Utility.RemoveAccent(value.(string)), 0.7, 4)
			isMatch = distance >= .85
		} else {
			return false, errors.New("Operator ~= can be only used with strings.")
		}
	} else if comparator == "!=" {
		// Equality comparator.
		// Case of string type.
		if reflect.TypeOf(expected).Kind() == reflect.String && reflect.TypeOf(value).Kind() == reflect.String {
			isMatch = Utility.RemoveAccent(expected.(string)) != Utility.RemoveAccent(value.(string))
		} else if reflect.TypeOf(expected).Kind() == reflect.Bool && reflect.TypeOf(value).Kind() == reflect.Bool {
			return expected.(bool) != value.(bool), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return expected.(int64) != value.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return expected.(float64) != value.(float64), nil
		}
	} else if comparator == "^=" {
		if reflect.TypeOf(expected).Kind() == reflect.String && reflect.TypeOf(value).Kind() == reflect.String {
			return strings.HasPrefix(value.(string), expected.(string)), nil
		} else {
			return false, nil
		}
	} else if comparator == "$=" {
		if reflect.TypeOf(expected).Kind() == reflect.String && reflect.TypeOf(value).Kind() == reflect.String {
			return strings.HasSuffix(value.(string), expected.(string)), nil
		} else {
			return false, nil
		}
	} else if comparator == "<" {
		// Number operator only...
		if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return value.(int64) < expected.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return value.(float64) < expected.(float64), nil
		}
	} else if comparator == "<=" {
		if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return value.(int64) <= expected.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return value.(float64) <= expected.(float64), nil
		}
	} else if comparator == ">" {
		if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return value.(int64) > expected.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return value.(float64) > expected.(float64), nil
		}
	} else if comparator == ">=" {
		if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return value.(int64) >= expected.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return value.(float64) >= expected.(float64), nil
		}
	}
	
	return isMatch, nil
}

/**
 * That function test if a given value match all expressions of a given ast...
 */
func (this *GraphStore) match(ast *ast.QueryAst, values map[string]interface{}) (bool, error) {

	// test if the value is composite.
	if ast.IsComposite() {
		ast1, _, ast2 := ast.GetSubQueries()
		// both side of the tree must match.
		isMatch, err := this.match(ast1, values)
		if err != nil {
			return false, err
		}
		if isMatch == false {
			return false, nil
		}

		isMatch, err = this.match(ast2, values)
		if err != nil {
			return false, err
		}

		if isMatch == false {
			return false, nil
		}
	} else {
		// I will evaluate the expression...
		typeName, fieldName, comparator, expected := ast.GetExpression()
		return this.evaluate(typeName, fieldName, comparator, expected, values[fieldName])
	}

	return true, nil
}

func (this *GraphStore) getIndexationValues(uuid string, fields []string) (map[string]interface{}, error) {
	object := make(map[string]interface{}, 0)

	// So here I will retreive values for given object fields...
	typeName := strings.Split(uuid, "%")[0]

	prefix := generatePrefix(typeName, "UUID") + formalize(uuid)
	query := xapian.NewQuery(prefix)

	results, err := this.runXapianQuery(typeName, query, fields)
	object["TYPENAME"] = typeName

	if len(fields) > 0 {
		for i := 0; i < len(fields); i++ {
			object[fields[i]] = results[0][i]
		}
	} else {
		object = results[0][0].(map[string]interface{})
	}
	return object, err
}

/**
 * Run a query and return it result.
 */
func (this *GraphStore) runXapianQuery(typeName string, query xapian.Query, fields []string) ([][]interface{}, error) {
	var err error
	var results [][]interface{}
	var prototype *EntityPrototype
	prototype, err = this.GetEntityPrototype(typeName)
	if err != nil {
		return results, err
	}

	path := this.m_path + "/" + typeName + ".glass"
	db := xapian.NewDatabase(path)

	enquire := xapian.NewEnquire(db)
	enquire.Set_query(query)
	mset := enquire.Get_mset(uint(0), uint(10000))

	// Now I will process the results.
	for i := 0; i < mset.Size(); i++ {
		it := mset.Get_hit(uint(i))
		doc := it.Get_document()
		if len(fields) > 0 {
			values := make([]interface{}, 0)
			for j := 0; j < len(fields); j++ {
				// Get the field index.
				fieldIndex := prototype.getFieldIndex(fields[j])
				fieldType := prototype.FieldsType[fieldIndex]
				value := doc.Get_value(uint(fieldIndex))
				if XML_Schemas.IsXsNumeric(fieldType) {
					values = append(values, xapian.Sortable_unserialise(value))
				} else if XML_Schemas.IsXsString(fieldType) {
					values = append(values, value)
				} else if XML_Schemas.IsXsId(fieldType) {
					values = append(values, value)
				} else if XML_Schemas.IsXsDate(fieldType) {
					values = append(values, xapian.Sortable_unserialise(value))
				} else {
					var v map[string]interface{}
					err := json.Unmarshal([]byte(doc.Get_data()), &v)
					if err == nil {
						values = append(values, v[fields[j]])
					} else {
						values = append(values, nil)
					}
				}
			}
			results = append(results, values)

		} else {
			// In that case the data contain in the document are return.
			var v map[string]interface{}
			err := json.Unmarshal([]byte(doc.Get_data()), &v)
			if err == nil {
				results = append(results, []interface{}{v})
			}
		}
		xapian.DeleteDocument(doc)
		xapian.DeleteMSetIterator(it)
	}

	if len(results) == 0 {
		err = errors.New("No results found!")
	}

	db.Close()
	xapian.DeleteMSet(mset)
	xapian.DeleteEnquire(enquire)
	xapian.DeleteQuery(query)
	xapian.DeleteDatabase(db)

	return results, err
}

// Return the UUID for a given
func (this *GraphStore) getIndexation(typeName string, fieldName string, expected interface{}) ([]interface{}, error) {

	// Indexations contain array of string
	var ids []interface{}
	var prototype, err = this.GetEntityPrototype(typeName)
	if err != nil {
		return ids, err
	}
	// I will retreive the value...
	if len(fieldName) == 0 {
		// Indexation by typeName...
		typeNameIndex := generatePrefix(typeName, "TYPENAME") + formalize(typeName)
		query := xapian.NewQuery(typeNameIndex)
		results, err := this.runXapianQuery(typeName, query, []string{"UUID"})
		if err != nil {
			return ids, err
		}
		for i := 0; i < len(results); i++ {
			ids = append(ids, results[i][0].(string))
		}
	} else {
		fieldIndex := prototype.getFieldIndex(fieldName)
		if fieldIndex != -1 {
			fieldType := prototype.FieldsType[fieldIndex]
			fieldType_ := strings.Replace(fieldType, "[]", "", -1)
			fieldType_ = strings.Replace(fieldType_, ":Ref", "", -1)
			var query xapian.Query
			if fieldName == "UUID" || fieldName == "TYPENAME" || fieldName == "ParentUuid" || fieldName == "ParentLnk" {
				prefix := generatePrefix(typeName, fieldName) + formalize(Utility.ToString(expected))
				query = xapian.NewQuery(prefix)
			} else if strings.HasPrefix(fieldName, "M_") {
				if XML_Schemas.IsXsString(fieldType) || XML_Schemas.IsXsId(fieldType) {
					p := xapian.NewQueryParser()
					// I ill add the prefix to the query parser.
					p.Add_prefix(fieldName[2:], strings.ToUpper(fieldName))
					query = p.Parse_query(fieldName[2:] + ":" + expected.(string))
					xapian.DeleteQueryParser(p)
				}
			}

			results, err := this.runXapianQuery(typeName, query, []string{"UUID"})
			if err != nil {
				return ids, err
			}
			// Now I will get the results...
			for i := 0; i < len(results); i++ {
				ids = append(ids, results[i][0].(string))
			}

		}
	}
	return ids, nil
}

/**
 * Here i will walk the tree and generate the query.
 */
func (this *GraphStore) runQuery(ast *ast.QueryAst, fields []string) (map[string]map[string]interface{}, error) {
	// I will create the array if it dosent exist.
	results := make(map[string]map[string]interface{}, 0)

	if ast.IsComposite() {
		// Get the sub-queries
		ast1, operator, ast2 := ast.GetSubQueries()

		r1, err := this.runQuery(ast1, fields)
		if err != nil {
			return nil, err
		}

		r2, err := this.runQuery(ast2, fields)
		if err != nil {
			return nil, err
		}

		if operator == "&&" { // conjonction
			for k, v := range r2 {
				isMatch, err := this.match(ast1, v)
				if err != nil {
					return nil, err
				}
				if isMatch {
					results[k] = v
				}
			}
			for k, v := range r1 {
				isMatch, err := this.match(ast2, v)
				if err != nil {
					return nil, err
				}
				if isMatch {
					results[k] = v
				}
			}
		} else if operator == "||" { // disjonction
			results = this.merge(r1, r2)
		}

	} else {

		typeName, fieldName, comparator, expected := ast.GetExpression()
		values := make(map[string]map[string]interface{}, 0)
		// Need the prototype here.
		prototype, err := this.GetEntityPrototype(typeName)
		if err != nil {
			return nil, err
		}

		if fieldName == "TYPENAME" {
			indexations, err := this.getIndexation(typeName, "", "")
			if err != nil {
				return nil, err
			}
			for i := 0; i < len(indexations); i++ {
				values_, err := this.getIndexationValues(indexations[i].(string), fields)
				if err == nil {
					results[indexations[i].(string)] = values_
				}
			}

		} else {
			fieldType := prototype.FieldsType[prototype.getFieldIndex(fieldName)]
			isArray := strings.HasPrefix(fieldType, "[]")
			isRef := strings.HasSuffix(fieldType, ":Ref")
			fieldType = strings.Replace(fieldType, "[]", "", -1)
			isString := fieldType == "xs.string" || fieldType == "xs.token" || fieldType == "xs.anyURI" || fieldType == "xs.anyURI" || fieldType == "xs.IDREF" || fieldType == "xs.QName" || fieldType == "xs.NOTATION" || fieldType == "xs.normalizedString" || fieldType == "xs.Name" || fieldType == "xs.NCName" || fieldType == "xs.ID" || fieldType == "xs.language"

			// Integers types.
			isInt := fieldType == "xs.int" || fieldType == "xs.integer" || fieldType == "xs.long" || fieldType == "xs.unsignedInt" || fieldType == "xs.short" || fieldType == "xs.unsignedLong"

			// decimal value
			isDecimal := fieldType == "xs.float" || fieldType == "xs.decimal" || fieldType == "xs.double"

			// Date time
			isDate := fieldType == "xs.date" || fieldType == "xs.dateTime"

			fields = prototype.Fields // all field must be search...

			// Strings or references...
			if isString || isRef {
				// The string expected value...
				if expected != nil {
					expectedStr := expected.(string)
					isRegex := strings.HasPrefix(expectedStr, "/") && strings.HasSuffix(expectedStr, "/")
					if comparator == "==" && !isRegex {
						// Now i will get the value from the indexation.
						if len(expectedStr) > 0 {
							indexations, err := this.getIndexation(typeName, fieldName, expectedStr)
							if err == nil {
								for i := 0; i < len(indexations); i++ {
									values[indexations[i].(string)], err = this.getIndexationValues(indexations[i].(string), []string{})
									if err != nil {
										return nil, err
									}
									var isMatch bool
									if isArray {
										// Here I have an array of values to test.
										var strValues []string
										err = json.Unmarshal([]byte(values[indexations[i].(string)][fieldName].(string)), &strValues)
										if err != nil {
											return nil, err
										}
										for j := 0; j < len(strValues); j++ {
											isMatch, err = this.evaluate(typeName, fieldName, comparator, expected, strValues[j])
										}
									} else {
										isMatch, err = this.evaluate(typeName, fieldName, comparator, expected, values[indexations[i].(string)][fieldName])
									}

									if err != nil {
										return nil, err
									}
									if isMatch {
										// if the result match I put it inside the map result.
										results[indexations[i].(string)] = values[indexations[i].(string)]
									}
								}
							}
						}
					} else if comparator == "~=" || comparator == "!=" || comparator == "^=" || comparator == "$=" || (isRegex && comparator == "==") {
						// Here I will use the typename as indexation key...
						indexations, err := this.getIndexation(typeName, "", "")
						if err == nil {
							for i := 0; i < len(indexations); i++ {
								values[indexations[i].(string)], err = this.getIndexationValues(indexations[i].(string), fields)
								if err != nil {
									return nil, err
								}
								isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[indexations[i].(string)][fieldName])
								if err != nil {
									return nil, err
								}
								if isMatch {
									// if the result match I put it inside the map result.
									results[indexations[i].(string)] = values[indexations[i].(string)]
								}
							}
						} else {
							log.Println("---> 2230 ", err)
						}
					} else {
						if !isRegex {
							return nil, errors.New("Unexpexted comparator " + comparator + " for type \"string\".")
						} else {
							return nil, errors.New("Unexpexted comparator " + comparator + " for regex, use \"==\" insted")
						}
					}
				} else if isRef {
					// In that case the only the operato == and != are define.
					if comparator == "==" || comparator == "!=" {

						typeNameIndex := generatePrefix(typeName, "TYPENAME") + formalize(typeName)
						query := xapian.NewQuery(typeNameIndex)
						uuids, err := this.runXapianQuery(typeName, query, []string{"UUID"})

						if err != nil {
							return nil, err
						}

						if comparator == "==" {
							for i := 0; i < len(uuids); i++ {
								values, err := this.getIndexationValues(uuids[i][0].(string), fields)
								if err == nil {
									results[uuids[i][0].(string)] = values
								}
							}
						} else if comparator == "!=" {
							for i := 0; i < len(uuids); i++ {
								values, err := this.getIndexationValues(uuids[i][0].(string), fields)
								if err == nil {
									results[uuids[i][0].(string)] = values
								}
							}
						}

					} else {
						return nil, errors.New("Unexpexted comparator " + comparator + " for regex, use \"==\" insted")
					}
				}

			} else if fieldType == "xs.boolean" {
				if !(comparator == "==" || comparator == "!=") {
					return nil, errors.New("Unexpexted comparator " + comparator + " for bool values, use \"==\" or  \"!=\"")
				}

				// Get the boolean value.
				indexations, err := this.getIndexation(typeName, fieldName, strconv.FormatBool(expected.(bool)))
				if err == nil {
					for i := 0; i < len(indexations); i++ {
						values[indexations[i].(string)], err = this.getIndexationValues(indexations[i].(string), fields)
						if err != nil {
							return nil, err
						}

						isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[indexations[i].(string)][fieldName])
						if err != nil {
							return nil, err
						}
						if isMatch {
							// if the result match I put it inside the map result.
							results[indexations[i].(string)] = values[indexations[i].(string)]
						}
					}
				}
			} else if isInt || isDecimal || isDate { // Numeric values or date that are covert at evaluation time as integer.
				if comparator == "~=" {
					return nil, errors.New("Unexpexted comparator " + comparator + " for type numeric value.")
				}
				// Get the boolean value.
				if comparator == "==" {
					indexations, err := this.getIndexation(typeName, fieldName, expected)
					if err == nil {
						for i := 0; i < len(indexations); i++ {
							values[indexations[i].(string)], err = this.getIndexationValues(indexations[i].(string), fields)
							if err != nil {
								return nil, err
							}

							isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[indexations[i].(string)][fieldName])
							if err != nil {
								return nil, err
							}
							if isMatch {
								// if the result match I put it inside the map result.
								results[indexations[i].(string)] = values[indexations[i].(string)]
							}
						}
					}
				} else {
					// for the other comparator I will get all the entities of the given type and test each of those.
					indexations, err := this.getIndexation(typeName, "", "")
					if err == nil {
						for i := 0; i < len(indexations); i++ {
							values[indexations[i].(string)], err = this.getIndexationValues(indexations[i].(string), fields)
							if err != nil {
								return nil, err
							}

							isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[indexations[i].(string)][fieldName])
							if err != nil {
								return nil, err
							}
							if isMatch {
								// if the result match I put it inside the map result.
								results[indexations[i].(string)] = values[indexations[i].(string)]
							}
						}
					}
				}
			}
		}
	}
	//log.Println("----> 2341 ", results)
	return results, nil
}

/**
 * Execute a search query.
 */
func (this *GraphStore) executeSearchQuery(query string, fields []string) ([][]interface{}, error) {
	s := lexer.NewLexer([]byte(query))
	p := parser.NewParser()
	a, err := p.Parse(s)

	if err == nil {
		astree := a.(*ast.QueryAst)
		fieldLength := len(fields)
		r, err := this.runQuery(astree, fields)
		if err != nil {
			return nil, err
		}

		// Here I will keep the result part...
		results := make([][]interface{}, 0)
		for _, object := range r {
			if fieldLength == 0 {
				// In that case the whole object will be set in the result.
				results = append(results, []interface{}{object})
			} else {
				results_ := make([]interface{}, 0)
				for i := 0; i < fieldLength; i++ {
					results_ = append(results_, object[fields[i]])
				}
				results = append(results, results_)
			}
		}
		return results, err
	} else {
		log.Println("-------> query: ", query, fields)
		log.Println("--> search error ", err)
	}
	return nil, err
}
