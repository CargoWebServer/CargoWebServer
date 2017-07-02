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
 * @fileOverview Security related functions
 * @author Dave Courtois, Eric Kavalec
 * @version 1.0
 */
var SecurityManager = function () {
    if (server == undefined) {
        return
    }
    EventHub.call(this, SecurityEvent)

    return this
}

SecurityManager.prototype = new EventHub(null);
SecurityManager.prototype.constructor = SecurityManager;

SecurityManager.prototype.onEvent = function (evt) {
    EventHub.prototype.onEvent.call(this, evt)
}
SecurityManager.prototype.appendAccount = function (roleId, accountId, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(roleId, "STRING", "roleId"))
    params.push(createRpcData(accountId, "STRING", "accountId"))

    server.executeJsFunction(
        "SecurityManagerAppendAccount",
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

SecurityManager.prototype.appendAction = function (roleId, accountName, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(roleId, "STRING", "roleId"))
    params.push(createRpcData(accountName, "STRING", "accountName"))

    server.executeJsFunction(
        "SecurityManagerAppendAction",
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

SecurityManager.prototype.appendPermission = function (accountId, permissionType, pattern, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(accountId, "STRING", "accountId"))
    params.push(createRpcData(permissionType, "INTEGER", "permissionType"))
    params.push(createRpcData(pattern, "STRING", "pattern"))

    server.executeJsFunction(
        "SecurityManagerAppendPermission",
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

SecurityManager.prototype.canExecuteAction = function (actionName, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(actionName, "STRING", "actionName"))

    server.executeJsFunction(
        "SecurityManagerCanExecuteAction",
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

SecurityManager.prototype.changeAdminPassword = function (pwd, newPwd, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(pwd, "STRING", "pwd"))
    params.push(createRpcData(newPwd, "STRING", "newPwd"))

    server.executeJsFunction(
        "SecurityManagerChangeAdminPassword",
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

SecurityManager.prototype.createRole = function (id, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    server.executeJsFunction(
        "SecurityManagerCreateRole",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("CargoEntities.Role", "CargoEntities",
                function (prototype, caller) { // Success Callback
                    if (caller.results[0] == null) {
                        return
                    }
                    if (entities[caller.results[0].UUID] != undefined && caller.results[0].TYPENAME == caller.results[0].__class__) {
                        caller.successCallback(entities[caller.results[0].UUID], caller.caller)
                        return // break it here.
                    }

                    var entity = eval("new " + prototype.TypeName + "()")
                    entity.initCallback = function () {
                        return function (entity) {
                            caller.successCallback(entity, caller.caller)
                        }
                    }(caller)
                    entity.init(caller.results[0])
                },
                function (errMsg, caller) { // Error Callback
                    caller.errorCallback(errMsg, caller.caller)
                },
                { "caller": caller.caller, "successCallback": caller.successCallback, "errorCallback": caller.errorCallback, "results": results }
            )
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

SecurityManager.prototype.deleteRole = function (id, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    server.executeJsFunction(
        "SecurityManagerDeleteRole",
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

SecurityManager.prototype.getRole = function (id, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    server.executeJsFunction(
        "SecurityManagerGetRole",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("CargoEntities.Role", "CargoEntities",
                function (prototype, caller) { // Success Callback
                    if (caller.results[0] == null) {
                        return
                    }
                    if (entities[caller.results[0].UUID] != undefined && caller.results[0].TYPENAME == caller.results[0].__class__) {
                        caller.successCallback(entities[caller.results[0].UUID], caller.caller)
                        return // break it here.
                    }

                    var entity = eval("new " + prototype.TypeName + "()")
                    entity.initCallback = function () {
                        return function (entity) {
                            caller.successCallback(entity, caller.caller)
                        }
                    }(caller)
                    entity.init(caller.results[0])
                },
                function (errMsg, caller) { // Error Callback
                    caller.errorCallback(errMsg, caller.caller)
                },
                { "caller": caller.caller, "successCallback": caller.successCallback, "errorCallback": caller.errorCallback, "results": results }
            )
        },
        function (errMsg, caller) { // Error callback
            caller.errorCallback(errMsg, caller.caller)
            server.errorManager.onError(errMsg)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })
}

SecurityManager.prototype.hasAccount = function (roleId, accountId, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(roleId, "STRING", "roleId"))
    params.push(createRpcData(accountId, "STRING", "accountId"))

    server.executeJsFunction(
        "SecurityManagerHasAccount",
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

SecurityManager.prototype.hasAction = function (roleId, actionName, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(roleId, "STRING", "roleId"))
    params.push(createRpcData(actionName, "STRING", "actionName"))

    server.executeJsFunction(
        "SecurityManagerHasAction",
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

SecurityManager.prototype.removeAccount = function (roleId, accountId, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(roleId, "STRING", "roleId"))
    params.push(createRpcData(accountId, "STRING", "accountId"))

    server.executeJsFunction(
        "SecurityManagerRemoveAccount",
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

SecurityManager.prototype.removeAction = function (roleId, accountName, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(roleId, "STRING", "roleId"))
    params.push(createRpcData(accountName, "STRING", "accountName"))

    server.executeJsFunction(
        "SecurityManagerRemoveAction",
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

SecurityManager.prototype.removePermission = function (accountId, pattern, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(accountId, "STRING", "accountId"))
    params.push(createRpcData(pattern, "STRING", "pattern"))

    server.executeJsFunction(
        "SecurityManagerRemovePermission",
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

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// OAuth2 Ressource access... The client must be configure first to be able to get access to ressources.
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Execute a query to retreive given ressource. 
 * @param {string} clientId The clien id, defined in configuration.
 * @param {string} scope The access scope need for the query.
 * @param {string} query The query that the OAuth2 provider will execute to retreive the information.
 */
SecurityManager.prototype.getResource = function (clientId, scope, query, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    // Account uuid are set and reset at the time of login and logout respectively.
    var idTokenUuid = ""
    if (localStorage.getItem("idTokenUuid") != undefined) {
        idTokenUuid = localStorage.getItem("idTokenUuid")
    }

    var params = []
    params.push(createRpcData(clientId, "STRING", "clientId"))
    params.push(createRpcData(scope, "STRING", "scope"))
    params.push(createRpcData(query, "STRING", "query"))
    params.push(createRpcData(idTokenUuid, "STRING", "idTokenUuid"))

    // Call it on the server.
    server.executeJsFunction(
        "SecurityManagerGetResource", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (results, caller) {
            caller.successCallback(results[0], caller.caller)
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

