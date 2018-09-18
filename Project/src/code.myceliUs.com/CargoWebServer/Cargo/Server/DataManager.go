package Server

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"

	"compress/gzip"
	"io/ioutil"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
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

	/** That channel is use to access the dataStores map **/
	m_dataStoreChan chan map[string]interface{}
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
	dataManager.m_dataStoreChan = make(chan map[string]interface{}, 0)
	// Concurrency...
	go func() {
		for {
			select {
			case op := <-dataManager.m_dataStoreChan:
				if op["op"] == "getDataStore" {
					storeId := op["storeId"].(string)
					op["store"].(chan DataStore) <- dataManager.m_dataStores[storeId]
				} else if op["op"] == "getDataStores" {
					stores := make([]DataStore, 0)
					for _, store := range dataManager.m_dataStores {
						stores = append(stores, store)
					}
					op["stores"].(chan []DataStore) <- stores
				} else if op["op"] == "setDataStore" {
					store := op["store"].(DataStore)
					dataManager.m_dataStores[store.GetId()] = store

				} else if op["op"] == "removeDataStore" {
					storeId := op["storeId"].(string)
					delete(dataManager.m_dataStores, storeId)
				}
			}
		}
	}()

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
	storeConfigurations := GetServer().GetConfigurationManager().getActiveConfigurations().GetDataStoreConfigs()

	log.Println("--> initialyze DataManager")
	for i := 0; i < len(storeConfigurations); i++ {
		if this.getDataStore(storeConfigurations[i].GetId()) == nil {
			store, err := NewDataStore(storeConfigurations[i])
			if err != nil {
			} else {
				this.setDataStore(store)
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
	for i := 0; i < len(this.getDataStores()); i++ {
		err := this.getDataStores()[i].Connect()
		if err == nil {
			log.Println("--> connection ", this.getDataStores()[i].GetId(), " is opened!")
		} else {
			log.Println("--> fail to open connection ", this.getDataStores()[i].GetId())
		}
	}
}

func (this *DataManager) appendDefaultDataStore(config *Config.DataStoreConfiguration) {
	store, err := NewDataStore(config)
	if err != nil {
		log.Println(err)
	}

	this.setDataStore(store)
	store.Connect()
}

/**
 * Access Map functions...
 */
func (this *DataManager) getDataStore(id string) DataStore {
	arguments := make(map[string]interface{})
	arguments["op"] = "getDataStore"
	arguments["storeId"] = id
	result := make(chan DataStore)
	arguments["store"] = result
	this.m_dataStoreChan <- arguments
	return <-result
}

func (this *DataManager) getDataStores() []DataStore {
	arguments := make(map[string]interface{})
	arguments["op"] = "getDataStores"
	result := make(chan []DataStore)
	arguments["stores"] = result
	this.m_dataStoreChan <- arguments
	return <-result
}

func (this *DataManager) setDataStore(store DataStore) {
	arguments := make(map[string]interface{})
	arguments["op"] = "setDataStore"
	arguments["store"] = store
	this.m_dataStoreChan <- arguments
}

func (this *DataManager) removeDataStore(id string) {
	arguments := make(map[string]interface{})
	arguments["storeId"] = id
	arguments["op"] = "removeDataStore"
	this.m_dataStoreChan <- arguments
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

	return data, err
}

/**
 * Execute a query that create a new data. The data contains the new
 * value to insert in the DB.
 */
func (this *DataManager) createData(storeName string, query string, d []interface{}) (lastId interface{}, err error) {
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

	return
}

func (this *DataManager) deleteData(storeName string, query string, params []interface{}) (err error) {
	store := this.getDataStore(storeName)

	if store == nil {
		return errors.New("Data store " + storeName + " does not exist.")
	}

	err = store.Delete(query, params)
	if err != nil {
		err = errors.New("Query '" + query + "' failed with error '" + err.Error() + "'.")
		return
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

	return
}

func (this *DataManager) createDataStore(storeId string, storeName string, hostName string, ipv4 string, port int, storeType Config.DataStoreType, storeVendor Config.DataStoreVendor) (DataStore, *CargoEntities.Error) {

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
	storeConfigEntity, err_ := GetServer().GetEntityManager().getEntityById("Config.DataStoreConfiguration", "Config", ids)

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
		configEntity := GetServer().GetConfigurationManager().getActiveConfigurations()
		storeConfigEntity, err_ = GetServer().GetEntityManager().createEntity(configEntity, "M_dataStoreConfigs", storeConfig)
		if err_ != nil {
			return nil, err_
		}

	} else {
		storeConfig = storeConfigEntity.(*Config.DataStoreConfiguration)
		storeConfig.M_id = storeId
		storeConfig.M_storeName = storeName
		storeConfig.M_dataStoreVendor = storeVendor
		storeConfig.M_dataStoreType = storeType
		storeConfig.M_ipv4 = ipv4
		storeConfig.M_hostName = hostName
		storeConfig.M_port = port
		err := GetServer().GetEntityManager().saveEntity(storeConfig)
		if err != nil {
			return nil, err
		}
	}

	// Create the store here.
	store, err := NewDataStore(storeConfig)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err)
		return nil, cargoError
	}

	// Set the newly created datastore on the map.
	this.setDataStore(store)

	return store, nil
}

func (this *DataManager) deleteDataStore(storeId string) *CargoEntities.Error {

	if !Utility.IsValidVariableName(storeId) {
		cargoError := NewError(Utility.FileLine(), INVALID_VARIABLE_NAME_ERROR, SERVER_ERROR_CODE, errors.New("The storeId '"+storeId+"' is not valid."))
		return cargoError
	}

	store := this.getDataStore(storeId)
	if this.getDataStore(storeId) == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The storeId '"+storeId+"' doesn't exist."))
		log.Println("---> Store with id", storeId, "dosen't exist!")
		return cargoError
	}

	// Here I will remove all prototype.
	store.DeleteEntityPrototypes()
	store.Close()

	// Remove the storeObject from the storeMap
	this.removeDataStore(storeId)

	// Delete the directory
	filePath := GetServer().GetConfigurationManager().GetDataPath() + "/" + storeId
	err := os.RemoveAll(filePath)

	// I will also remove it schema if there is one.
	schemaPath := GetServer().GetConfigurationManager().GetSchemasPath() + "/" + storeId + ".xsd"
	if Utility.Exists(schemaPath) {
		// remove it schemas
		os.Remove(schemaPath)
	}

	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Failed to delete directory '"+filePath+"' with error '"+err.Error()+"'."))
		log.Println("---> Fail to remove ", storeId, err)
		return cargoError
	}
	return nil

}

func (this *DataManager) close() {
	// Close the data manager.
	for _, v := range this.m_dataStores {
		v.Close()
	}

}

func toJsonStr(object interface{}) string {
	b, _ := json.Marshal(object)
	// Convert bytes to string.
	b, _ = Utility.PrettyPrint(b)
	s := string(b)
	return s
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
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The datastore '"+storeName+"' does not exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}
	err := store.Connect()

	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Fail to open the data store connection "+err.Error()+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}
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
	/*errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}*/

	err := this.deleteData(storeName, query, parameters)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err))
	}
}

// @api 1.0
// Determine if the datastore exist in the server.
// @param {string} storeId The data server connection (configuration) id
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {bool} Return true if the datastore exist.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) HasDataStore(storeId string, messageId string, sessionId string) bool {
	storeConfig, _ := GetServer().GetEntityManager().getEntityById("Config.DataStoreConfiguration", "Config", []interface{}{storeId})
	return storeConfig != nil
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
func (this *DataManager) CreateDataStore(storeId string, storeName string, hostName string, ipv4 string, port float64, storeType float64, storeVendor float64, messageId string, sessionId string) {
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

	// Here I will send the event.
	var eventDatas []*MessageData
	evtData := new(MessageData)
	evtData.TYPENAME = "Server.MessageData"
	evtData.Name = "storeConfig"

	storeConfig, _ := GetServer().GetEntityManager().getEntityById("Config.DataStoreConfiguration", "Config", []interface{}{storeId})
	evtData.Value = storeConfig
	eventDatas = append(eventDatas, evtData)
	evt, _ := NewEvent(NewDataStoreEvent, DataEvent, eventDatas)
	GetServer().GetEventManager().BroadcastEvent(evt)

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

	// Here I will send the event.
	var eventDatas []*MessageData
	evtData := new(MessageData)
	evtData.TYPENAME = "Server.MessageData"
	evtData.Name = "storeId"

	evtData.Value = storeId
	eventDatas = append(eventDatas, evtData)
	evt, _ := NewEvent(DeleteDataStoreEvent, DataEvent, eventDatas)
	GetServer().GetEventManager().BroadcastEvent(evt)
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

// @api 1.0
// DataStore infos and all it shcemas infos. It can be use to import datastore
// infos.
// @param {string} storeId The store id from where the informations will be exported.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {string} Return DataStore information as json format.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) ExportJsonSchemas(storeId string, messageId string, sessionId string) string {
	storeConfig, errObj := GetServer().GetEntityManager().getEntityById("Config.DataStoreConfiguration", "Config", []interface{}{storeId})
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return ""
	}

	store := this.getDataStore(storeId)
	prototypes, _ := store.GetEntityPrototypes()

	infos := new(struct {
		DataStoreConfig *Config.DataStoreConfiguration
		Schemas         []*EntityPrototype
	})

	infos.DataStoreConfig = storeConfig.(*Config.DataStoreConfiguration)
	infos.Schemas = prototypes

	return toJsonStr(infos)
}

// @api 1.0
// Create schemas from inforamtion in the infos(json) string. If the datastore
// dosent exist a new datastore in created and prototype are create in it after.
// @param {string} jsonStr A json string with all necessary information in it.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {protected}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) ImportJsonSchema(jsonStr string, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	infos := new(struct {
		DataStoreConfig *Config.DataStoreConfiguration
		Schemas         []*EntityPrototype
	})

	err := json.Unmarshal([]byte(jsonStr), &infos)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	// So first of all I will test if the data store exist...
	var store DataStore
	if infos.DataStoreConfig.M_ipv4 != "127.0.0.1" {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Import is permitted on local server only "+infos.DataStoreConfig.M_id+" is a remote server with addresse "+infos.DataStoreConfig.M_ipv4))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	hasDataStore := this.HasDataStore(infos.DataStoreConfig.M_id, messageId, sessionId)
	if !hasDataStore {
		// Here I will create the new data store.
		var cargoError *CargoEntities.Error
		store, cargoError = this.createDataStore(infos.DataStoreConfig.M_id, infos.DataStoreConfig.M_storeName, infos.DataStoreConfig.M_hostName, infos.DataStoreConfig.M_ipv4, infos.DataStoreConfig.M_port, infos.DataStoreConfig.M_dataStoreType, infos.DataStoreConfig.M_dataStoreVendor)
		if cargoError != nil {
			GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		}
	} else {
		store = this.getDataStore(infos.DataStoreConfig.M_id)
	}

	// Prototypes will be create only if they dosent exist.
	for i := 0; i < len(infos.Schemas); i++ {
		prototype := infos.Schemas[i]
		_, errObj := store.GetEntityPrototype(prototype.TypeName)
		if errObj != nil {
			// Here I will create the prototype.
			store.CreateEntityPrototype(prototype)
		} else {
			// Here I will save it.
			store.SaveEntityPrototype(prototype)
		}
	}

	// Open connection.
	store.Connect()
}

// @api 1.0
// Return Dump the content of a data store into json string and compress it.
// The archive contain the information of the schemas and the data it contain.
// That function can be use to create backup of the data store.
// @param {string} storeId The store id from where the informations will be exported.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {string} Return A link to the created archive, the caller must delete the when done.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *DataManager) ExportJsonData(storeId string, messageId string, sessionId string) string {

	storeConfig, errObj := GetServer().GetEntityManager().getEntityById("Config.DataStoreConfiguration", "Config", []interface{}{storeId})
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return ""
	}

	store := this.getDataStore(storeId)
	prototypes, _ := store.GetEntityPrototypes()

	infos := new(struct {
		DataStoreConfig *Config.DataStoreConfiguration
		Schemas         []*EntityPrototype
		Data            []interface{} // That will contain all data
	})

	infos.DataStoreConfig = storeConfig.(*Config.DataStoreConfiguration)
	infos.Schemas = prototypes
	infos.Data = make([]interface{}, 0)

	// Now the data...
	for i := 0; i < len(prototypes); i++ {
		entities, errObj := GetServer().GetEntityManager().getEntities(prototypes[i].TypeName, storeId, nil)
		if errObj == nil {
			for j := 0; j < len(entities); j++ {
				// append the object.
				if reflect.TypeOf(entities[i]).String() == "*Server.DynamicEntity" {
					infos.Data = append(infos.Data, entities[j].(*DynamicEntity).getValues())
				} else {
					infos.Data = append(infos.Data, entities[j])
				}
			}
		}
	}

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(toJsonStr(infos))); err != nil {
		panic(err)
	}

	if err := gz.Flush(); err != nil {
		panic(err)
	}

	if err := gz.Close(); err != nil {
		panic(err)
	}

	var path = GetServer().GetConfigurationManager().GetApplicationDirectoryPath() + "/" + storeId + ".gz"
	err := ioutil.WriteFile(path, b.Bytes(), 0666)

	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return ""
	}

	return storeId + ".gz"
}

// @api 1.0
// Import JSON archived data.
// @param {string} filename The name of the file to create.
// @param {[]byte} filedata The data to import (.gz file).
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} progressCallback The function is call when chunk of response is received.
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
//DataManager.prototype.importJsonData = function (filename, filedata, successCallback, progressCallback, errorCallback, caller) {
//    // server is the client side singleton.
//    var params = [];
//    // The file data (filedata) will be upload with the http protocol...
//    params.push(createRpcData(filename, "STRING", "filename"));
//    // Here I will create a new data form...
//    var formData = new FormData();
//    formData.append("multiplefiles", filedata, filename)
//    // Use the post function to upload the file to the server.
//    var xhr = new XMLHttpRequest();
//    xhr.open('POST', '/uploads', true);
//    // In case of error or success...
//    xhr.onload = function (params, xhr) {
//        return function (e) {
//            if (xhr.readyState === 4) {
//                if (xhr.status === 200) {
//                    // Here I will create the file...
//                    server.executeJsFunction(
//                        "DataManagerImportJsonData", // The function to execute remotely on server
//                        params, // The parameters to pass to that function
//                        function (index, total, caller) { // The progress callback
//                            // Keep track of the file transfert.
//                            caller.progressCallback(index, total, caller.caller);
//                        },
//                        function (result, caller) {
//                            caller.successCallback(result[0], caller.caller);
//                        },
//                        function (errMsg, caller) {
//                            // display the message in the console.
//                            // call the immediate error callback.
//                            caller.errorCallback(errMsg, caller.caller);
//                            // dispatch the message.
//                            server.errorManager.onError(errMsg);
//                        }, // Error callback
//                        { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback } // The caller
//                    )
//                } else {
//                    console.error(xhr.statusText);
//                }
//            }
//        }
//    } (params, xhr)
//    // now the progress event...
//    xhr.upload.onprogress = function (progressCallback, caller) {
//        return function (e) {
//            if (e.lengthComputable) {
//                progressCallback(e.loaded, e.total, caller);
//            }
//        }
//    } (progressCallback, caller)
//    xhr.send(formData);
//}
func (this *DataManager) ImportJsonData(filename string, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	tmpPath := GetServer().GetConfigurationManager().GetTmpPath() + "/" + filename

	// I will open the file form the tmp directory.
	f, err := os.Open(tmpPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	defer gr.Close()

	data, err := ioutil.ReadAll(gr)
	if err != nil {
		//log.Println(err)
		return
	}

	// here I will initialyse data.
	infos := new(struct {
		DataStoreConfig *Config.DataStoreConfiguration
		Schemas         []*EntityPrototype
		Data            []interface{} // That will contain all data
	})

	err = json.Unmarshal(data, &infos)
	if err != nil {
		//log.Println(err)
		return
	}

	// The first step will be to be sure that the store exist.
	// So first of all I will test if the data store exist...
	var store DataStore
	if infos.DataStoreConfig.M_ipv4 != "127.0.0.1" {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Import is permitted on local server only "+infos.DataStoreConfig.M_id+" is a remote server with addresse "+infos.DataStoreConfig.M_ipv4))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	store = this.getDataStore(infos.DataStoreConfig.M_id)
	if store == nil {
		var cargoError *CargoEntities.Error
		store, cargoError = this.createDataStore(infos.DataStoreConfig.M_id, infos.DataStoreConfig.M_storeName, infos.DataStoreConfig.M_hostName, infos.DataStoreConfig.M_ipv4, infos.DataStoreConfig.M_port, infos.DataStoreConfig.M_dataStoreType, infos.DataStoreConfig.M_dataStoreVendor)
		if cargoError != nil {
			GetServer().reportErrorMessage(messageId, sessionId, cargoError)
			return
		}
	}

	// Open connection.
	err = store.Connect()
	// Prototypes will be create only if they dosent exist.
	for i := 0; i < len(infos.Schemas); i++ {
		prototype := infos.Schemas[i]
		_, errObj := store.GetEntityPrototype(prototype.TypeName)
		if errObj != nil {
			// Here I will create the prototype.
			store.CreateEntityPrototype(prototype)
		} else {
			// Here I will save it.
			store.SaveEntityPrototype(prototype)
		}
	}
	if err != nil {
		//log.Println(err)
		return
	}
	// Now I will save the data...
	var entities []Entity
	for i := 0; i < len(infos.Data); i++ {
		values := infos.Data[i].(map[string]interface{})
		value, err := Utility.InitializeStructure(values, setEntityFct)
		obj := value.Interface()

		if err == nil {
			if reflect.TypeOf(obj).String() == "map[string]interface {}" {
				entity := NewDynamicEntity()
				entity.setObject(obj.(map[string]interface{}))
				entities = append(entities, entity)
			} else {
				// In that case I will create a static entity.
				entities = append(entities, obj.(Entity))
			}
		}
	}

	// Now I can save it.
	for i := 0; i < len(entities); i++ {
		//log.Println("---> create: ", entities[i])
		GetServer().GetEntityManager().saveEntity(entities[i])
	}
	// remove the tmp file if it file path is not empty... otherwise the
	// file will bee remove latter.
	defer os.Remove(tmpPath)
}
