
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
 * @fileOverview Contains the defintion of the AccountManager class.
 * @author Dave Courtois, Eric Kavalec
 * @version 1.0
 */

/**
 * The AccountManager is used to get information about user accounts.
  *@constructor
 * @extends EventHub
 */
var AccountManager = function () {

    this.account = null

    if (server == undefined) {
        return
    }

    EventHub.call(this, AccountEvent)

    return this
}

AccountManager.prototype = new EventHub(null);
AccountManager.prototype.constructor = AccountManager;

/*
 * Dispatch event
 */
AccountManager.prototype.onEvent = function (evt) {
    if (evt.code == AccountRegisterSuccessEvent) {
        this.account = evt.dataMap["accountInfo"][0]
    } else if (evt.code == ContactInvitationReceivedsuccessEvent) {
        this.account = evt.dataMap["toInfo"][0]
    }
    EventHub.prototype.onEvent.call(this, evt)
}

/**
 * Returns the account type.
 * @returns {string} The account type.
 */
AccountManager.prototype.getAccountType = function () {
    return this.account.M_accountType
}

/*
 * Server side code. 
 */
function registerAccount(name, password, email) {
    var newAccount = null
    newAccount = server.GetAccountManager().Register(name, password, email, messageId, sessionId)
    return newAccount
}

/**
 * Register a new account.
 * @param {string} name The name of the new account.
 * @param {string} password The password associated with the new account.
 * @param {string} email The email of the new account.
 * @param {function} successCallback The function to execute in case of role creation success
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
AccountManager.prototype.register = function (name, password, email, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(name, "STRING", "name"))
    params.push(createRpcData(password, "STRING", "password"))
    params.push(createRpcData(email, "STRING", "email"))

    // Call it on the server.
    server.executeJsFunction(
        registerAccount.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // Progress callback
            // Nothing to do here.
        },
        function (result, caller) { // Success callback
            /*
            // Create a new event localy with the data receive from the server
            var evt = new EventMsg(randomUUID(), server.conn, AccountEvent, AccountRegisterSuccessEvent, [], undefined, undefined, undefined)
            evt.dataMap["accountInfo"] = result
            // Throw the event
            server.accountManager.onEvent(evt)
            // Return the account in the successcallback
            */
            caller.successCallback(result[0], caller.caller)

        },
        function (errMsg, caller) { // Error callback
            // display the message in the console.
            console.log(errMsg)
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message.
            server.errorManager.onError(errMsg)
        },
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/* 
 * Server side code.
 */
function GetAccountById(id) {
    var account
    account = server.GetAccountManager().GetAccountById(id, messageId, sessionId)
    return account
}

/**
 * Retreive an account with a given id.
 * @param {string} id The id of the account.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
AccountManager.prototype.getAccountById = function (id, successCallback, errorCallback, caller) {

    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    // Call it on the server.
    server.executeJsFunction(
        GetAccountById.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {

            var account = eval("new " + result[0].TYPENAME + "()")
            account.initCallback = function (caller) {
                return function (account) {
                    caller.successCallback(account, caller.caller)
                }
            } (caller)

            account.init(result[0])
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
function GetUserById(id) {
    var objects
    objects = server.GetAccountManager().GetUserById(id, messageId, sessionId)
    return objects
}

/**
 * Retreive a user with a given id.
 * @param {string} id The id of the user.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
AccountManager.prototype.getUserById = function (id, successCallback, errorCallback, caller) {

    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    // Call it on the server.
    server.executeJsFunction(
        GetUserById.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            var user = eval("new " + result[0].TYPENAME + "()")
            account.initCallback = function (caller) {
                return function (account) {
                    caller.successCallback(account, caller.caller)
                }
            } (caller)

            user.init(result[0])
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