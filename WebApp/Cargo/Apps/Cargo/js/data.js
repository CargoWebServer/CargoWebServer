var DataManager = function () {
    if (server == undefined) {
        return
    }
    EventHub.call(this, DataEvent)

    return this
}

DataManager.prototype = new EventHub(null);
DataManager.prototype.constructor = DataManager;

DataManager.prototype.onEvent = function (evt) {
    EventHub.prototype.onEvent.call(this, evt)
}
DataManager.prototype.close = function (storeName, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeName, "STRING", "storeName"))

    server.executeJsFunction(
        "DataManagerClose",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.connect = function (storeName, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeName, "STRING", "storeName"))

    server.executeJsFunction(
        "DataManagerConnect",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.create = function (storeName, query, d, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeName, "STRING", "storeName"))
    params.push(createRpcData(query, "STRING", "query"))
    params.push(createRpcData(d, "JSON_STR", "d"))

    server.executeJsFunction(
        "DataManagerCreate",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.createDataStore = function (storeId, storeType, storeVendor, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeId, "STRING", "storeId"))
    params.push(createRpcData(storeType, "INTEGER", "storeType"))
    params.push(createRpcData(storeVendor, "INTEGER", "storeVendor"))

    server.executeJsFunction(
        "DataManagerCreateDataStore",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.delete = function (storeName, query, parameters, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeName, "STRING", "storeName"))
    params.push(createRpcData(query, "STRING", "query"))
    params.push(createRpcData(parameters, "JSON_STR", "parameters"))

    server.executeJsFunction(
        "DataManagerDelete",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.deleteDataStore = function (storeId, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeId, "STRING", "storeId"))

    server.executeJsFunction(
        "DataManagerDeleteDataStore",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.hasDataStore = function (storeName, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeName, "STRING", "storeName"))

    server.executeJsFunction(
        "HasDataStore",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.importXmlData = function (content, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(content, "STRING", "content"))

    server.executeJsFunction(
        "DataManagerImportXmlData",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.importXsdSchema = function (storeId, content, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeId, "STRING", "storeId"))
    params.push(createRpcData(content, "STRING", "content"))

    server.executeJsFunction(
        "DataManagerImportXsdSchema",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.ping = function (storeName, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeName, "STRING", "storeName"))

    server.executeJsFunction(
        "DataManagerPing",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.read = function (storeName, query, fieldsType, parameters, successCallback, progressCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeName, "STRING", "storeName"))
    params.push(createRpcData(query, "STRING", "query"))
    params.push(createRpcData(fieldsType, "JSON_STR", "fieldsType"))
    params.push(createRpcData(parameters, "JSON_STR", "parameters"))

    server.executeJsFunction(
        "DataManagerRead",
        params,
        function (index, total, caller) { // Progress callback
            caller.progressCallback(index, total, caller.caller)
        },
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "progressCallback": progressCallback, "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

DataManager.prototype.synchronize = function (storeId, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(storeId, "STRING", "storeId"))

    server.executeJsFunction(
        "DataManagerSynchronize",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            caller.successCallback(results, caller.caller)
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}