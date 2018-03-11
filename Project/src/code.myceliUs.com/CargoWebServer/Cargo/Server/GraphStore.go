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

	// The backend...
	"github.com/boltdb/bolt"
	"github.com/syndtr/goleveldb/leveldb"
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

	// The B+Tree That database is use to indexation of triple.
	m_index *bolt.DB

	// That level Db is fater in write operation so triple values will
	// be store in it.
	m_db *leveldb.DB
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

	// Now I will open the db.
	store.m_index, err = bolt.Open(store.m_path+"/"+store.m_id+".db", 0777, nil)
	if err != nil {
		log.Fatal(err)
	}

	// The triple store, speed write is better with levelDb than BoltDb
	store.m_db, err = leveldb.OpenFile(store.m_path+"/"+store.m_id+".lvl", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Here I will register all class in the vm.
	prototypes, err := store.GetEntityPrototypes()
	if err == nil {
		for i := 0; i < len(prototypes); i++ {
			// The script will be put in global context (CargoWebServer)
			JS.GetJsRuntimeManager().AppendScript("CargoWebServer/"+prototypes[i].TypeName, prototypes[i].generateConstructor(), false)
		}
	}

	return
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
				value, err := Utility.InitializeStructure(results[0])
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
			GetServer().GetProcessor().m_sendRequest <- rqst
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

	// I will save the entity prototype in a file.
	file, err := os.Create(this.m_path + "/" + prototype.TypeName + ".gob")
	defer file.Close()

	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(prototype)
	} else {
		return err
	}
	this.m_prototypes[prototype.TypeName] = prototype

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
				value, err := Utility.InitializeStructure(results[0])
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
			GetServer().GetProcessor().m_sendRequest <- rqst
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

	// Update local entities if the store is local.
	/*entities, _ := GetServer().GetEntityManager().getEntities(prototype_.TypeName, this.m_id, nil)

	// Remove the fields
	for i := 0; i < len(entities); i++ {
		entity := entities[i] // Must be a dynamic entity.

		// remove it...
		for j := 0; j < len(prototype.FieldsToDelete); j++ {
			field := prototype_.Fields[prototype.FieldsToDelete[j]]
			if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
				// Dynamic entity.
				entity.(*DynamicEntity).deleteValue(field)
			}

			GetServer().GetEntityManager().saveEntity(entity)
		}

		// update it...
		for j := 0; j < len(prototype.FieldsToUpdate); j++ {
			values := strings.Split(prototype.FieldsToUpdate[j], ":")
			if len(values) == 2 {
				indexFrom := prototype_.getFieldIndex(values[0])
				indexTo := prototype.getFieldIndex(values[1])
				if indexFrom > -1 && indexTo > -1 {
					if values[0] != values[1] {
						if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
							// Set the new value with the old one
							entity.(*DynamicEntity).setValue(values[1], entity.(*DynamicEntity).getValue(values[0]))
							// Delete the old one.
							entity.(*DynamicEntity).deleteValue(values[0])
						}

						prototype_.Fields[indexFrom] = values[1]
					}
					var fieldTypeTo = prototype_.FieldsType[indexTo]
					var fieldTypeFrom = prototype.FieldsType[indexFrom]
					if fieldTypeFrom != fieldTypeTo {
						log.Println("------> change field type from ", fieldTypeFrom, "with", fieldTypeTo)
						// TODO set conversion rules here for each possible types.
					}
				}
			}
		}

		// Now set new fields value inside existing entities with their default
		// value.
		for j := 0; j < len(prototype.Fields); j++ {
			if !Utility.Contains(prototype_.Fields, prototype.Fields[j]) {
				// I that case I will set the new field value inside the prototype.
				var value interface{}
				if strings.HasPrefix(prototype.FieldsType[j], "[]") {
					value = "undefined"
				} else {
					if XML_Schemas.IsXsString(prototype.FieldsType[j]) {
						value = prototype.FieldsDefaultValue[j]
					} else if XML_Schemas.IsXsInt(prototype.FieldsType[j]) || XML_Schemas.IsXsTime(prototype.FieldsType[j]) {
						value, _ = strconv.ParseInt(prototype.FieldsDefaultValue[j], 10, 64)
					} else if XML_Schemas.IsXsNumeric(prototype.FieldsType[j]) {
						value, _ = strconv.ParseFloat(prototype.FieldsDefaultValue[j], 64)
					} else if XML_Schemas.IsXsDate(prototype.FieldsType[j]) {
						value = Utility.MakeTimestamp()
					} else if XML_Schemas.IsXsBoolean(prototype.FieldsType[j]) {
						if prototype.FieldsDefaultValue[j] == "false" {
							value = false
						} else {
							value = true
						}
					} else {
						// Object here.
						value = "undefined"
					}
				}
				if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
					entity.(*DynamicEntity).setValue(prototype.Fields[j], value)
				}
			}
		}

		// Save the entity.
		GetServer().GetEntityManager().saveEntity(entity)
	}*/

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
			GetServer().GetProcessor().m_sendRequest <- rqst
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
							evtData.Value = parent.(*DynamicEntity).getObject()
						} else {
							evtData.Value = parent
						}
						eventDatas = append(eventDatas, evtData)
						evt, _ := NewEvent(UpdateEntityEvent, EntityEvent, eventDatas)
						GetServer().GetEventManager().BroadcastEvent(evt)
					}
				}
				GetServer().GetEntityManager().m_removeEntityChan <- entity

				// Send event message...
				var eventDatas []*MessageData
				evtData := new(MessageData)
				evtData.TYPENAME = "Server.MessageData"
				evtData.Name = "entity"
				if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
					evtData.Value = entity.(*DynamicEntity).getObject()
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

func (this *GraphStore) getValues(uuid string, fields []string) (map[string]interface{}, error) {
	object := make(map[string]interface{}, 0)

	// So here I will retreive values for given object fields...
	typeName := strings.Split(uuid, "%")[0]
	var prototype, err = this.GetEntityPrototype(typeName)
	if err != nil {
		return object, err
	}

	for i := 0; i < len(fields); i++ {
		fieldName := fields[i]
		fieldIndex := prototype.getFieldIndex(fieldName)
		if fieldIndex != -1 {
			fieldType := prototype.FieldsType[fieldIndex]
			fieldType_ := strings.Replace(fieldType, "[]", "", -1)
			fieldType_ = strings.Replace(fieldType_, ":Ref", "", -1)
			predicat := typeName + ":" + fieldType_ + ":" + fieldName

			q := "( " + uuid + ", " + predicat + ", ? )"
			results, err := this.Read(q, []interface{}{}, []interface{}{})
			if err == nil {
				if len(results) > 0 {
					// Only the first value will be keep here.
					// In case of array values are store as json string so no more than
					// one element must be here.
					object[fieldName] = results[0][2]
				}
			}
		}
	}
	return object, nil
}

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
		q := "( ?, TYPENAME, " + typeName + " )"
		results, err := this.Read(q, []interface{}{}, []interface{}{})
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
			predicat := typeName + ":" + fieldType_ + ":" + fieldName

			q := "( ?, " + predicat + ", " + Utility.ToString(expected) + " )"
			results, err := this.Read(q, []interface{}{}, []interface{}{})
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

		// I will append the indexs and the ids to the list of field if there's
		// not already in.
		for i := 0; i < len(prototype.Ids); i++ {
			if !Utility.Contains(fields, prototype.Ids[i]) {
				fields = append(fields, prototype.Ids[i])
			}
		}
		// The indexs
		for i := 0; i < len(prototype.Indexs); i++ {
			if !Utility.Contains(fields, prototype.Indexs[i]) {
				fields = append(fields, prototype.Indexs[i])
			}
		}

		// Strings or references...
		if isString || isRef {
			// The string expected value...
			expectedStr := expected.(string)
			isRegex := strings.HasPrefix(expectedStr, "/") && strings.HasSuffix(expectedStr, "/")
			if comparator == "==" && !isRegex {
				// Now i will get the value from the indexation.
				indexations, err := this.getIndexation(typeName, fieldName, expectedStr)
				if err == nil {
					for i := 0; i < len(indexations); i++ {

						values[indexations[i].(string)], err = this.getValues(indexations[i].(string), fields)
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
			} else if comparator == "~=" || comparator == "!=" || comparator == "^=" || comparator == "$=" || (isRegex && comparator == "==") {
				// Here I will use the typename as indexation key...
				indexations, err := this.getIndexation(typeName, "", "")
				if err == nil {
					for i := 0; i < len(indexations); i++ {
						values[indexations[i].(string)], err = this.getValues(indexations[i].(string), fields)
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
				if !isRegex {
					return nil, errors.New("Unexpexted comparator " + comparator + " for type \"string\".")
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
					values[indexations[i].(string)], err = this.getValues(indexations[i].(string), fields)
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
						values[indexations[i].(string)], err = this.getValues(indexations[i].(string), fields)
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
						values[indexations[i].(string)], err = this.getValues(indexations[i].(string), fields)
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
			results_ := make([]interface{}, 0)
			for i := 0; i < fieldLength; i++ {
				results_ = append(results_, object[fields[i]])
			}
			results = append(results, results_)
		}

		return results, err
	} else {
		log.Println("--> search error ", err)
	}
	return nil, err
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
				value, err := Utility.InitializeStructure(results[0])
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
			GetServer().GetProcessor().m_sendRequest <- rqst
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
					values, err := Utility.InitializeStructure(results[0][i])
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
			GetServer().GetProcessor().m_sendRequest <- rqst
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
				log.Println(this.m_conn)
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
				values, err := Utility.InitializeStructure(results[0])

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
			GetServer().GetProcessor().m_sendRequest <- rqst
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
			GetServer().GetProcessor().m_sendRequest <- rqst
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
func (this *GraphStore) Create(queryStr string, triples []interface{}) (lastId interface{}, err error) {
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
		param3.Value = triples

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
			GetServer().GetProcessor().m_sendRequest <- rqst
		}(rqst)

		// wait for result here.
		results := <-resultsChan

		// in case of error
		if reflect.TypeOf(results).String() == "*string" {
			return -1, errors.New(*results.(*string))
		}

		return results, nil
	}

	// So here I will index the triple...
	for i := 0; i < len(triples); i++ {
		// This will contain the value of the triple.
		triple := triples[i].(Triple)
		data, err := json.Marshal(&triple)

		if err == nil {
			// The triple uuid
			uuid := Utility.GenerateUUID(string(data))
			err = this.indexTriple(uuid, data)
			if err == nil {
				// S, ?, ?
				err = this.index(triple.Subject, uuid, "S")
				if err == nil {
					// ?, P, ?
					err = this.index(triple.Predicate, uuid, "P")
					if err == nil {
						if triple.IsIndex || Utility.IsValidEntityReferenceName(Utility.ToString(triple.Object)) {
							// ?, ?, O
							this.index(Utility.ToString(triple.Object), uuid, "O")
						}
						// Combine value.
						//  S, P, ?
						sp := Utility.GenerateUUID(triple.Subject + ":" + triple.Predicate)
						err = this.index(sp, uuid, "SP")
						if err == nil {
							so := Utility.GenerateUUID(triple.Subject + ":" + Utility.ToString(triple.Object))
							err = this.index(so, uuid, "SO")
							if err == nil {
								po := Utility.GenerateUUID(triple.Predicate + ":" + Utility.ToString(triple.Object))
								err = this.index(po, uuid, "PO")
							}
						}
					}
				}
			}
		}
	}

	// The triples to save...
	for i := 0; i < len(triples); i++ {
		log.Println("------> save triple ", triples[i])
	}

	return
}

// Append a triple into the database.
func (this *GraphStore) indexTriple(uuid string, data []byte) error {
	err := this.m_db.Put([]byte(uuid), data, nil)
	return err
}

func (this *GraphStore) removeTriple(uuid string) error {
	err := this.m_db.Delete([]byte(uuid), nil)
	return err
}

func (this *GraphStore) index(key string, uuid string, bucketId string) error {
	return this.m_index.Update(func(key string, uuid string, bucketId string) func(tx *bolt.Tx) error {
		return func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte(bucketId))
			if err == nil {
				v := b.Get([]byte(key))
				if v != nil {
					if strings.Index(string(v), uuid) == -1 {
						v = []byte(string(v) + ":" + uuid)
					}
				} else {
					v = []byte(uuid)
				}
				err = b.Put([]byte(key), v)
			}
			return err
		}
	}(key, uuid, bucketId))
}

// Remove the triple indexation from the DB.
func (this *GraphStore) removeIndex(key string, uuid string, bucketId string) error {
	return this.m_index.Update(func(key string, uuid string, bucketId string) func(tx *bolt.Tx) error {
		return func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucketId))
			var err error
			if b != nil {
				v := b.Get([]byte(key))
				v = []byte(strings.Replace(string(v), ":"+uuid, "", -1))
				v = []byte(strings.Replace(string(v), uuid, "", -1))
				if len(v) == 0 {
					err = b.Delete([]byte(key))
				} else {
					err = b.Put([]byte(key), v)
				}
			}
			return err
		}
	}(key, uuid, bucketId))
}

// Here I will retreive triple values.
func (this *GraphStore) getTriple(key string) (*Triple, error) {
	data, err := this.m_db.Get([]byte(key), nil)
	triple := new(Triple)
	err = json.Unmarshal(data, triple)
	if err == nil {
		return triple, nil
	}
	return nil, err
}

// Here I will retreive triple values.
func (this *GraphStore) getTriples(key string, bucketId string) []Triple {
	results := make(chan []Triple, 0)
	go this.m_index.View(func(key string, bucketId string, results chan []Triple, store *GraphStore) func(*bolt.Tx) error {
		return func(tx *bolt.Tx) error {
			triples := make([]Triple, 0)
			b := tx.Bucket([]byte(bucketId))
			if b == nil {
				results <- triples
				return errors.New("Bucket " + bucketId + " not found!")
			}
			v := b.Get([]byte(key))

			// Now I will return the found tripple.
			keys := strings.Split(string(v), ":")
			for i := 0; i < len(keys); i++ {
				if len(keys[i]) > 0 {
					triple, err := store.getTriple(keys[i])
					if err == nil {
						triples = append(triples, *triple)
					}
				}
			}

			results <- triples
			return nil
		}
	}(key, bucketId, results, this))
	return <-results
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
			GetServer().GetProcessor().m_sendRequest <- rqst
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

	var triples []Triple
	// So query will be of the form...
	// In case of simple query of the form (?, ?, ?)
	if strings.HasPrefix(queryStr, "(") && strings.HasSuffix(queryStr, ")") {
		values := strings.Split(queryStr[1:len(queryStr)-1], ",")
		s := strings.TrimSpace(values[0])
		p := strings.TrimSpace(values[1])
		o := strings.TrimSpace(values[2])
		if s == "?" && p != "?" && o != "?" { // ?, P, O
			bucketId := "PO"
			key := Utility.GenerateUUID(p + ":" + o)
			triples = this.getTriples(key, bucketId)
		} else if s != "?" && p == "?" && o != "?" { // S, ?, O
			bucketId := "SO"
			key := Utility.GenerateUUID(s + ":" + o)
			triples = this.getTriples(key, bucketId)
		} else if s != "?" && p != "?" && o == "?" { // S, P, ?
			bucketId := "SP"
			key := Utility.GenerateUUID(s + ":" + p)
			triples = this.getTriples(key, bucketId)
		} else if s != "?" && p == "?" && o == "?" { // ?, P, O
			bucketId := "S"
			triples = this.getTriples(s, bucketId)
		} else if s != "?" && p != "?" && o == "?" { // S, ?, O
			bucketId := "P"
			triples = this.getTriples(p, bucketId)
		} else if s == "?" && p == "?" && o != "?" { // S, P, ?
			bucketId := "O"
			triples = this.getTriples(o, bucketId)
		}

		if len(triples) == 0 {
			return nil, errors.New("No values found!")
		}

		results = make([][]interface{}, 0)
		for i := 0; i < len(triples); i++ {
			results = append(results, []interface{}{triples[i].Subject, triples[i].Predicate, triples[i].Object})
		}

		return results, nil

	} else {
		// First of all i will init the query...
		var query EntityQuery
		err := json.Unmarshal([]byte(queryStr), &query)

		if err != nil {
			return nil, err
		}

		if len(query.Query) > 0 {
			searchResult, err := this.executeSearchQuery(query.Query, query.Fields)
			if err != nil {
				return nil, err
			}
			return searchResult, nil
		}
	}
	return
}

/**
 * Update a entity value.
 */
func (this *GraphStore) Update(queryStr string, triples []interface{}, params []interface{}) (err error) {
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
		param3.Value = triples

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
			GetServer().GetProcessor().m_sendRequest <- rqst
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

	// The triples to save...
	for i := 0; i < len(triples); i++ {
		log.Println("------> save triple ", triples[i])
	}

	return
}

/**
 * Delete entity from the store...
 */
func (this *GraphStore) Delete(queryStr string, triples []interface{}) (err error) {
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
			GetServer().GetProcessor().m_sendRequest <- rqst
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
	for i := 0; i < len(triples); i++ {
		log.Println("remove triple: ", triples[i])
		data, err := json.Marshal(&triples[i])
		uuid := Utility.GenerateUUID(string(data))
		if err == nil {
			this.removeTriple(uuid)
			triple := triples[i].(Triple)

			this.removeIndex(triple.Subject, uuid, "S")

			this.removeIndex(triple.Predicate, uuid, "P")

			if triple.IsIndex {
				this.removeIndex(Utility.ToString(triple.Object), uuid, "O")
			}

			sp := Utility.GenerateUUID(triple.Subject + ":" + triple.Predicate)
			this.removeIndex(sp, uuid, "SP")

			so := Utility.GenerateUUID(triple.Predicate + ":" + Utility.ToString(triple.Object))
			this.removeIndex(so, uuid, "SO")

			po := Utility.GenerateUUID(triple.Predicate + ":" + Utility.ToString(triple.Object))
			this.removeIndex(po, uuid, "PO")
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

	// Close the datastore.
	this.m_index.Close()

	// Close level db.
	this.m_db.Close()

	return nil
}
