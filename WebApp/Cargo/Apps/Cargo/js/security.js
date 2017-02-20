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

/**
 * The security manager's role is to manipulate 
 * roles and permissions on entities
 * @constructor
 * @extends EventManager
 */
var SecurityManager = function () {

    if (server == undefined) {
        return
    }

    EventManager.call(this, SecurityEvent)

    return this
}

SecurityManager.prototype = new EventManager(null);
SecurityManager.prototype.constructor = SecurityManager;

/*
 * Dispatch event.
 */
SecurityManager.prototype.onEvent = function (evt) {
    EventManager.prototype.onEvent.call(this, evt)
}

SecurityManager.prototype.RegisterListener = function () {
    // Append to the event handler.
    server.eventHandler.AddEventManager(this,
        // callback
        function () {
            console.log("Security manager is registered")
        }
    )
}

/*
 * Sever side code.
 */
function CreateRole(id) {
    var newRole = null
    newRole = server.GetSecurityManager().CreateRole(id, messageId, sessionId)
    return newRole
}

/**
 * Create a new role
 * @param {string} id The id of the role to create
 * @param {function} successCallback The function to execute in case of role creation success
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
SecurityManager.prototype.createRole = function (id, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    // Call it on the server.
    server.executeJsFunction(
        CreateRole.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            if (result[0] != null) {
                var role = eval("new " + result[0].TYPENAME + "()")
                role.init(result[0])
                caller.successCallback(role, caller.caller)
            }
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
 * Sever side code.
 */
function GetRole(id) {
    var role = null
    role = server.GetSecurityManager().GetRole(id, messageId, sessionId)
    return role
}

/**
 * Retreive a role with a given id.
 * @param {string} id The id of the role to retreive
 * @param {function} successCallback The function to execute in case of success
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
SecurityManager.prototype.getRole = function (id, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    // Call it on the server.
    server.executeJsFunction(
        GetRole.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            if (result[0] != null) {
                var role = eval("new " + result[0].TYPENAME + "()")
                role.init(result[0])
                caller.successCallback(role, caller.caller)
            }
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
 * Sever side code.
 */
function DeleteRole(id) {
    server.GetSecurityManager().DeleteRole(id, messageId, sessionId)
}

/**
 * Delete a role with a given id.
 * @param {string} id The id the of role to retreive
 * @param {function} successCallback The function to execute in case of success
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
SecurityManager.prototype.deleteRole = function (id, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    // Call it on the server.
    server.executeJsFunction(
        DeleteRole.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
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
 * Sever side code.
 */
function HasAccount(roleId, accountId) {
    var roleHasAccount = server.GetSecurityManager().HasAccount(roleId, accountId, messageId, sessionId)
    return roleHasAccount
}

/**
 * Determines if a role has a given account.
 * @param {string} roleId The id of the role to verify
 * @param {string} accountId The id of the account to verify
 * @param {function} successCallback The function to execute in case of success
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
SecurityManager.prototype.hasAccount = function (roleId, accountId, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(roleId, "STRING", "roleId"))
    params.push(createRpcData(accountId, "STRING", "accountId"))

    // Call it on the server.
    server.executeJsFunction(
        HasAccount.toString(), // The function to execute remotely on server
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
 * Sever side code.
 */
function AppendAccount(roleId, accountId) {
    server.GetSecurityManager().AppendAccount(roleId, accountId, messageId, sessionId)
}

/**
 * Append a new account to a given role. Does nothing if the account is already in the role
 * @param {string} roleId The id of the role to append the account to 
 * @param {string} accountId The id of the account to append
 * @param {function} successCallback The function to execute in case of success
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
SecurityManager.prototype.appendAccount = function (roleId, accountId, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(roleId, "STRING", "roleId"))
    params.push(createRpcData(accountId, "STRING", "accountId"))

    // Call it on the server.
    server.executeJsFunction(
        AppendAccount.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
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
 * Sever side code.
 */
function RemoveAccount(roleId, accountId) {
    server.GetSecurityManager().RemoveAccount(roleId, accountId, messageId, sessionId)
}

/**
 * Remove an account from a given role.
 * @param {string} roleId The id of the role to remove the account from 
 * @param {string} accountId The id of the account to remove from the role
 * @param {function} successCallback The function to execute in case of success
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
SecurityManager.prototype.removeAccount = function (roleId, accountId, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(roleId, "STRING", "roleId"))
    params.push(createRpcData(accountId, "STRING", "accountId"))

    // Call it on the server.
    server.executeJsFunction(
        RemoveAccount.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
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
 * Sever side code.
 */
function ChangeAdminPassword(pwd, newPwd) {
    server.GetSecurityManager().ChangeAdminPassword(pwd, newPwd, messageId, sessionId)
}

/**
 * Change the current password for the admin account.
 * @param {string} pwd The current password
 * @param {string} newPwd The new password
 */
SecurityManager.prototype.changeAdminPassword = function (pwd, newPwd, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(pwd, "STRING", "pwd"))
    params.push(createRpcData(newPwd, "STRING", "newPwd"))

    // Call it on the server.
    server.executeJsFunction(
        ChangeAdminPassword.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
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