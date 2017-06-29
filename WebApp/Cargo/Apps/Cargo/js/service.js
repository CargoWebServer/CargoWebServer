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
 * @fileOverview Service related functions
 * @author Dave Courtois
 * @version 1.0
 */

/**
 * The service manager's is to interact with service.
 * @constructor
 * @extends EventHub
 */
var ServiceManager = function () {

    if (server == undefined) {
        return
    }

    EventHub.call(this, ServiceEvent)

    return this
}

ServiceManager.prototype = new EventHub(null);
ServiceManager.prototype.constructor = ServiceManager;

/*
 * Dispatch event.
 */
ServiceManager.prototype.onEvent = function (evt) {
    EventHub.prototype.onEvent.call(this, evt)
}

/**
 * Get the list of actions for a given service.
 * @param {string} serviceName The name of the service
 * @param {function} successCallback The function to execute in case of success
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
ServiceManager.prototype.getServiceActions = function (serviceName, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(serviceName, "STRING", "serviceName"))

    // Call it on the server.
    server.executeJsFunction(
        "GetServiceActions", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            //console.log(result)
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message.
            server.errorManager.onError(errMsg)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}
