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
	"code.myceliUs.com/Utility"

	"github.com/xrash/smetrics"

	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/ast"
	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/lexer"
	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/parser"

	// Xapian datastore.
	base64 "encoding/base64"

	"code.myceliUs.com/GoXapian"
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

	// This the channel to comminicate with the datastore.
	m_db_channel chan map[string]interface{}
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
	store.m_db_channel = make(chan map[string]interface{})

	// datastore operation presscessing function.
	go func(db_channel chan map[string]interface{}) {
		// Keep store in memory.
		for {
			select {
			case op := <-db_channel:
				if op["operation"].(string) == "openDataStore" {
					paths := op["paths"].([]string)
					var db xapian.Database
					var err error
					for i := 0; i < len(paths); i++ {
						path := paths[i]
						if !op["isWritable"].(bool) {
							if Utility.Exists(path) {
								db = xapian.NewDatabase()
								// log.Println(db)
								db.Add_database(xapian.NewDatabase(path))
							} else {
								err = errors.New("fail to open " + path)
							}

						} else {
							// create of open de db.
							db = xapian.NewWritableDatabase(path, xapian.DB_CREATE_OR_OPEN)
						}
					}
					// return the result to the channel.
					op["results"].(chan []interface{}) <- []interface{}{db, err}
				}
			}
		}

	}(store.m_db_channel)

	return
}

// Open database
func (this *GraphStore) OpenDataStoreWrite(paths []string) (xapian.WritableDatabase, error) {
	op := make(map[string]interface{})
	op["operation"] = "openDataStore"
	op["paths"] = paths
	op["isWritable"] = true
	op["results"] = make(chan []interface{})

	this.m_db_channel <- op
	results := <-op["results"].(chan []interface{})

	if results[1] != nil {
		return nil, results[1].(error)
	}

	return results[0].(xapian.WritableDatabase), nil
}

func (this *GraphStore) OpenDataStoreRead(paths []string) (xapian.Database, error) {
	op := make(map[string]interface{})
	op["operation"] = "openDataStore"
	op["paths"] = paths
	op["isWritable"] = false
	op["results"] = make(chan []interface{})

	this.m_db_channel <- op
	results := <-op["results"].(chan []interface{})

	if results[1] != nil {
		return nil, results[1].(error)
	}

	return results[0].(xapian.Database), nil
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

func generatePrefix(typeName string, field string) string {
	// remove the M_ part of the field name
	prefix := typeName + "." + field
	prefix = strings.Replace(prefix, ".", "_", -1) + "%"
	return "X" + strings.ToLower(prefix)
}

// Index entity string field.
func (this *GraphStore) indexStringField(data string, field string, typeName string, termGenerator xapian.TermGenerator) {

	if Utility.IsUriBase64(data) {
		data_, err := base64.StdEncoding.DecodeString(data)
		if err == nil {
			if strings.Index(data, ":text/") > -1 || strings.Index(data, ":application/") > -1 {
				log.Println("1092 ----> base64", data_)
			}
		}
	} else if Utility.IsStdBase64(data) {
		data_, err := base64.StdEncoding.DecodeString(data)
		if err == nil {
			log.Println("1087 ----> base64", data_)
			termGenerator.Index_text(strings.ToLower(string(data_)))
			termGenerator.Index_text(strings.ToLower(string(data_)), uint(1), generatePrefix(typeName, field))
		}
	} else {
		termGenerator.Index_text(strings.ToLower(data))
		termGenerator.Index_text(strings.ToLower(data), uint(1), generatePrefix(typeName, field))
	}
}

// Index entity field
func (this *GraphStore) indexField(data interface{}, field string, fieldType string, typeName string, termGenerator xapian.TermGenerator, doc xapian.Document, index uint) {
	if data != nil {
		if reflect.TypeOf(data).Kind() == reflect.Slice {
			s := reflect.ValueOf(data)
			for i := 0; i < s.Len(); i++ {
				this.indexField(s.Index(i).Interface(), field, fieldType, typeName, termGenerator, doc, index)
			}
		} else if reflect.TypeOf(data).Kind() == reflect.String {
			this.indexStringField(data.(string), field, typeName, termGenerator)
		} else {
			doc.Add_value(index, Utility.ToString(data))
		}
	}
}

// index entity information.
func (this *GraphStore) indexEntity(doc xapian.Document, entity Entity) {

	// The term generator
	termGenerator := xapian.NewTermGenerator()

	// set english by default.
	termGenerator.Set_stemmer(xapian.NewStem("en"))
	termGenerator.Set_document(doc)

	// Index the typename
	termGenerator.Index_text(entity.GetTypeName(), uint(1), "S")

	// Index each part it typename.
	values_ := strings.Split(entity.GetTypeName(), ".")
	for i := 0; i < len(values_); i++ {
		termGenerator.Index_text(values_[i])
		termGenerator.Increase_termpos()
	}

	prototype, _ := this.GetEntityPrototype(entity.GetTypeName())

	// also index value supertype...
	for i := 0; i < len(prototype.SuperTypeNames); i++ {
		termGenerator.Index_text(prototype.SuperTypeNames[i], uint(1), "S")
		values_ := strings.Split(prototype.SuperTypeNames[i], ".")
		for j := 0; j < len(values_); j++ {
			termGenerator.Index_text(values_[j])
			termGenerator.Increase_termpos()
		}
	}

	// Here I will append boolean term.
	for i := 0; i < len(prototype.Fields); i++ {
		value := Utility.GetProperty(entity, prototype.Fields[i])
		this.indexField(value, prototype.Fields[i], prototype.FieldsType[i], prototype.TypeName, termGenerator, doc, uint(i))
		if Utility.Contains(prototype.Ids, prototype.Fields[i]) || Utility.Contains(prototype.Indexs, prototype.Fields[i]) {
			// append a boolean term.
			term := generatePrefix(prototype.TypeName, prototype.Fields[i]) + strings.TrimSpace(strings.ToLower(Utility.ToString(value)))
			doc.Add_boolean_term(term)
		}
	}
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

	// Creation of triples...
	for i := 0; i < len(values); i++ {
		v := values[i]
		// If the value is an entity...
		if reflect.TypeOf(v).Kind() == reflect.Ptr || reflect.TypeOf(v).Kind() == reflect.Struct {
			typeName := Utility.GetProperty(v, "TYPENAME")
			uuid := Utility.GetProperty(v, "UUID")
			if typeName != nil {
				db, err := this.OpenDataStoreWrite([]string{this.m_path + "/" + typeName.(string) + ".glass"})
				if err == nil {
					defer db.Close()
					// So here I will index the property found in the entity.
					doc := xapian.NewDocument()

					// Keep json data...
					data, _ := Utility.ToJson(v)
					this.indexEntity(doc, v.(Entity))
					doc.Set_data(data)
					doc.Add_boolean_term("Q" + uuid.(string))
					db.Replace_document("Q"+uuid.(string), doc)
					log.Println("---> append entity ", v.(Entity))
				}
			}
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

	if len(query.Query) > 0 {
		var err error
		results, err = this.executeSearchQuery(query.Query, query.Fields)
		if err != nil {
			return nil, err
		}
	}

	return
}

/**
 * Update a entity value.
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
	log.Println("---> update entity ", values)
	// So in that case I will udate the values.

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
	log.Println("---> remove entity ", values)

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
				expected = expectedDateValue.Unix() // get the unix time for calcul
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

// Get potential document from the database.
func (this *GraphStore) getIndexation(typeName string, fieldName string, expected interface{}) ([]map[string]interface{}, error) {
	db, err := this.OpenDataStoreRead([]string{this.m_path + "/" + typeName + ".glass"})
	if err != nil {
		return nil, err
	}
	defer db.Close()

	term := generatePrefix(typeName, fieldName) + strings.TrimSpace(strings.ToLower(Utility.ToString(expected)))
	var results []map[string]interface{}

	// Here the term will be search.
	q := xapian.NewQuery(term)
	enquire := xapian.NewEnquire(db)
	enquire.Set_query(q)
	mset := enquire.Get_mset(uint(0), uint(100000))
	for i := 0; i < mset.Size(); i++ {
		it := mset.Get_hit(uint(i))
		var v map[string]interface{}
		err := json.Unmarshal([]byte(it.Get_document().Get_data()), &v)
		if err != nil {
			return nil, err
		}
		results = append(results, v)
	}
	return results, nil
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

		store := GetServer().GetDataManager().getDataStore(strings.Split(typeName, ".")[0])

		// Need the prototype here.
		prototype, err := store.GetEntityPrototype(typeName)
		if err != nil {
			return nil, err
		}

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
								values := indexations[i]
								if err != nil {
									return nil, err
								}
								var isMatch bool
								if isArray {
									// Here I have an array of values to test.
									for j := 0; j < len(values[fieldName].([]interface{})); j++ {
										isMatch, err = this.evaluate(typeName, fieldName, comparator, expected, values[fieldName].([]interface{})[j])
									}
								} else {
									isMatch, err = this.evaluate(typeName, fieldName, comparator, expected, values[fieldName])
								}

								if err != nil {
									return nil, err
								}
								if isMatch {
									// if the result match I put it inside the map result.
									results[indexations[i]["UUID"].(string)] = values
								}
							}
						}
					}
				} else if comparator == "~=" || comparator == "!=" || comparator == "^=" || comparator == "$=" || (isRegex && comparator == "==") {
					// Here I will use the typename as indexation key...
					indexations, err := this.getIndexation(typeName, "", "")
					if err == nil {
						for i := 0; i < len(indexations); i++ {
							values := indexations[i]
							if err != nil {
								return nil, err
							}

							isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[fieldName])
							if err != nil {
								return nil, err
							}
							if isMatch {
								// if the result match I put it inside the map result.
								results[indexations[i]["UUID"].(string)] = values
							}
						}
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
				// TODO get all value for that field and test if it match the constraint...
				// ex. (?, typeName:fieldName:fieldType, ?) if the object respect the contraint put the predicate in the
				// result map.
				if comparator == "==" || comparator == "!=" {

					/*q := "( ?, TYPENAME, " + typeName + " )"
					uuids, err := this.Read(q, []interface{}{}, []interface{}{})

					if err != nil {
						return nil, err
					}

					fieldType_ := strings.Replace(fieldType, ":Ref", "", -1)
					predicat := typeName + ":" + fieldType_ + ":" + fieldName

					if comparator == "==" {
						for i := 0; i < len(uuids); i++ {
							q := "( " + uuids[i][0].(string) + ", " + predicat + ", ? )"
							_, err := this.Read(q, []interface{}{}, []interface{}{})
							if err != nil { // No value found.
								values, err := this.getValues(uuids[i][0].(string), fields)
								if err == nil {
									results[uuids[i][0].(string)] = values
								}
							}
						}
					} else if comparator == "!=" {
						for i := 0; i < len(uuids); i++ {
							q := "( " + uuids[i][0].(string) + ", " + predicat + ", ? )"
							_, err := this.Read(q, []interface{}{}, []interface{}{})
							if err == nil { // if there is no error that mean a value is found.
								values, err := this.getValues(uuids[i][0].(string), fields)
								if err == nil {
									results[uuids[i][0].(string)] = values
								}
							}
						}
					}*/

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
					values := indexations[i]
					if err != nil {
						return nil, err
					}

					isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[fieldName])
					if err != nil {
						return nil, err
					}
					if isMatch {
						// if the result match I put it inside the map result.
						results[indexations[i]["UUID"].(string)] = values
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
						values := indexations[i]
						if err != nil {
							return nil, err
						}

						isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[fieldName])
						if err != nil {
							return nil, err
						}
						if isMatch {
							// if the result match I put it inside the map result.
							results[indexations[i]["UUID"].(string)] = values
						}
					}
				}
			} else {
				// for the other comparator I will get all the entities of the given type and test each of those.
				indexations, err := this.getIndexation(typeName, "", "")
				if err == nil {
					for i := 0; i < len(indexations); i++ {
						values := indexations[i]
						if err != nil {
							return nil, err
						}

						isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[fieldName])
						if err != nil {
							return nil, err
						}
						if isMatch {
							// if the result match I put it inside the map result.
							results[indexations[i]["UUID"].(string)] = values
						}
					}
				}
			}
		}
	}

	if len(results) == 0 {
		return nil, errors.New("no result found!")
	}

	log.Println("----> 2103", results)
	return results, nil
}

/**
 * Execute a search query.
 */
func (this *GraphStore) executeSearchQuery(query string, fields []string) ([][]interface{}, error) {

	query = `(CargoEntities.User.M_firstName ~= "Eric" || CargoEntities.User.M_firstName == "Louis") && CargoEntities.User.M_lastName != "Boucher"`
	query = `CargoEntities.Entities.M_id == "CARGO_ENTITIES"`

	log.Println("---> run query: ", query)
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
			results_ := make([]interface{}, 0)
			// if the fields are specefied i will retrun it.
			if fieldLength > 0 {
				for i := 0; i < fieldLength; i++ {
					// I will keep only required fields.
					results_ = append(results_, object[fields[i]])
				}
				results = append(results, results_)
			} else {
				// append the object in the array.
				results = append(results, []interface{}{object})
			}
		}

		return results, err
	} else {
		log.Panicln("--> search error ", err)
	}
	return nil, err
}
