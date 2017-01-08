/*
 * (C) Copyright 2016 Mycelius SA (http://mycelius.com/).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * @fileOverview Contain the defintion of the DataManager class.
 * @author Dave Courtois
 * @version 1.0
 */

/**
 * The datamanager is use to get information from data source like SQL or Cargo object store.
  *@constructor
 * @extends EventManager
 */
var DataManager = function (id) {

    if (server == undefined) {
        return
    }

    if (id == undefined) {
        id = randomUUID()
    }

    EventManager.call(this, id, TableEvent)

    return this
}

DataManager.prototype = new EventManager(null);
DataManager.prototype.constructor = DataManager;

/*
 * Dispatch event.
 */
DataManager.prototype.onEvent = function (evt) {
    EventManager.prototype.onEvent.call(this, evt)
}

DataManager.prototype.RegisterListener = function () {
    // Append to the event handler.
    server.eventHandler.AddEventManager(this,
        // callback
        function () {
            console.log("Listener registered!!!!")
        }
    )
}

/*
 * Server side code.
 */
function Read(connectionId, query, fields, params) {
    var values = null
    values = server.GetDataManager().Read(connectionId, query, fields, params, messageId, sessionId)
    return values
}

/**
 * Execute a read query on the data sever.
 * @param {string} connectionId The data server connection (configuration) id
 * @param {string} query The query string to execute.
 * @param {} fields Contain the list of type of queryied data. ex. string, date, int, float.
 * @param {} params Contain filter expression, ex. id=0, id != 3.
 * @param {function} progressCallback The function is call when chunk of response is received.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
DataManager.prototype.read = function (connectionId, query, fields, params, successCallback, progressCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params_ = []
    params_.push(new RpcData({ "name": "connectionId", "type": 2, "dataBytes": utf8_to_b64(connectionId) }))
    params_.push(new RpcData({ "name": "query", "type": 2, "dataBytes": utf8_to_b64(query) }))
    params_.push(new RpcData({ "name": "fields", "type": 4, "dataBytes": utf8_to_b64(JSON.stringify(fields)), "typeName": "[]string"}))
    params_.push(new RpcData({ "name": "params", "type": 4, "dataBytes": utf8_to_b64(JSON.stringify(params)), "typeName": "[]interface{}" }))

    // Call it on the server.
    server.executeJsFunction(
        Read.toString(), // The function to execute remotely on server
        params_, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
            caller.progressCallback(index, total, caller.caller)
        },
        function (result, caller) {
            //console.log(result[0])
            caller.successCallback(result, caller.caller)
        },
        function (errMsg, caller) {
            // display the message in the console.
            console.log(errMsg)
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message.
            server.errorManager.onError(errMsg)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * server side code.
 */
function Create(connectionId, query, values) {
    var id = null
    id = server.GetDataManager().Create(connectionId, query, values, messageId, sessionId)
    return id
}

/**
 * Create an new entry in the DB and return it's id(s)
 * @param {string} connectionId The data server connection (configuration) id
 * @param {string} query The query string to execute.
 * @param {} values Contain the list of values associated with the fields in the query.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
DataManager.prototype.create = function (connectionId, query, values, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params_ = []
    params_.push(new RpcData({ "name": "connectionId", "type": 2, "dataBytes": utf8_to_b64(connectionId) }))
    params_.push(new RpcData({ "name": "query", "type": 2, "dataBytes": utf8_to_b64(query) }))
    params_.push(new RpcData({ "name": "values", "type": 4, "dataBytes": utf8_to_b64(JSON.stringify(values)), "typeName": "[]interface{}" }))

    // Call it on the server.
    server.executeJsFunction(
        Create.toString(), // The function to execute remotely on server
        params_, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            //console.log(result[0])
            caller.successCallback(result, caller.caller)
        },
        function (errMsg, caller) {
            // display the message in the console.
            console.log(errMsg)
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message.
            server.errorManager.onError(errMsg)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * server side code
 */
function Update(connectionId, query, fields, params) {
    // No value are return.
    server.GetDataManager().Update(connectionId, query, fields, params, messageId, sessionId)
}

/**
 * Update existing database values.
 * @param {string} connectionId The data server connection (configuration) id
 * @param {string} query The query string to execute.
 * @param {} fields Contain the list of type of queryied data. ex. string, date, int, float.
 * @param {} params Contain filter expression, ex. id=0, id != 3.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
DataManager.prototype.update = function (connectionId, query, fields, params, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params_ = []
    params_.push(new RpcData({ "name": "connectionId", "type": 2, "dataBytes": utf8_to_b64(connectionId) }))
    params_.push(new RpcData({ "name": "query", "type": 2, "dataBytes": utf8_to_b64(query) }))
    params_.push(new RpcData({ "name": "fields", "type": 4, "dataBytes": utf8_to_b64(JSON.stringify(fields)), "typeName": "[]interface{}" }))
    params_.push(new RpcData({ "name": "params", "type": 4, "dataBytes": utf8_to_b64(JSON.stringify(params)), "typeName": "[]interface{}" }))

    // Call it on the server.
    server.executeJsFunction(
        Update.toString(), // The function to execute remotely on server
        params_, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            console.log(result[0])
            caller.successCallback(result, caller.caller)
        },
        function (errMsg, caller) {
            // display the message in the console.
            console.log(errMsg)
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message.
            server.errorManager.onError(errMsg)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * Server side code.
 */
function Delete(connectionId, query, params) {
    // No value are return.
    server.GetDataManager().Delete(connectionId, query, params, messageId, sessionId)
}

/**
 * Delete db value.
 * @param {string} connectionId The data server connection (configuration) id
 * @param {string} query The query string to execute.
 * @param {} params Contain filter expression, ex. id=0, id != 3.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
DataManager.prototype.delete = function (connectionId, query, params, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params_ = []
    params_.push(new RpcData({ "name": "connectionId", "type": 2, "dataBytes": utf8_to_b64(connectionId) }))
    params_.push(new RpcData({ "name": "query", "type": 2, "dataBytes": utf8_to_b64(query) }))
    params_.push(new RpcData({ "name": "params", "type": 4, "dataBytes": utf8_to_b64(JSON.stringify(params)), "typeName": "[]interface{}" }))

    // Call it on the server.
    server.executeJsFunction(
        Delete.toString(), // The function to execute remotely on server
        params_, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            console.log(result[0])
            caller.successCallback(result, caller.caller)
        },
        function (errMsg, caller) {
            // display the message in the console.
            console.log(errMsg)
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message.
            server.errorManager.onError(errMsg)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}


/*
 * server side code
 */
function CreateDataStore(storeId, storeType, storeVendor) {
    server.GetDataManager().CreateDataStore(storeId, storeType, storeVendor, messageId, sessionId)
}

/**
 * Create a dataStore
 * @param {string} storeId The id of the dataStore to create
 * @param {int} storeType The type of the store to create. SQL: 1; KEY_VALUE: 3
 * @param {int} storeVendor The store vendor. DataStoreVendor_MYCELIUS: 1; 	DataStoreVendor_MYSQL: 2; DataStoreVendor_MSSQL:3; DataStoreVendor_ODBC
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
DataManager.prototype.createDataStore = function (storeId, storeType, storeVendor, successCallback, errorCallback, caller) {
    var params = []

    params.push(createRpcData(storeId, "STRING", "storeId"))
    params.push(createRpcData(storeType, "INTEGER", "storeType"))
    params.push(createRpcData(storeVendor, "INTEGER", "storeVendor"))

    // Call it on the server.
    server.executeJsFunction(
        CreateDataStore.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            // display the message in the console.
            console.log(errMsg)
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message.
            server.errorManager.onError(errMsg)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * server side code
 */
function DeleteDataStore(storeId) {
    server.GetDataManager().DeleteDataStore(storeId, messageId, sessionId)
}

/**
 * Delete a dataStore
 * @param {string} storeId The id of the dataStore to delete
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
DataManager.prototype.deleteDataStore = function (storeId, successCallback, errorCallback, caller) {
    var params = []

    params.push(createRpcData(storeId, "STRING", "storeId"))

    // Call it on the server.
    server.executeJsFunction(
        DeleteDataStore.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            // display the message in the console.
            console.log(errMsg)
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message.
            server.errorManager.onError(errMsg)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}