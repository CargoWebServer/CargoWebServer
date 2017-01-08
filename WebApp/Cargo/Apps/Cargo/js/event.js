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
* @fileOverview Event management functinality.
* @author Dave Courtois
* @version 1.0
*/

AccountEvent = "AccountEvent"
AccountRegisterSuccessEvent = 0
AccountConfirmationSucessEvent = 1
SessionEvent = "SessionEvent"
LoginEvent = 4
LogoutEvent = 5
StateChangeEvent = 6
BpmnEvent = "BpmnEvent"
NewProcessInstanceEvent = 7
UpdateProcessInstanceEvent = 8
NewDefinitionsEvent = 9
DeleteDefinitionsEvent = 10
UpdateDefinitionsEvent = 11
EntityEvent = "EntityEvent"
NewEntityEvent = 12
UpdateEntityEvent = 13
DeleteEntityEvent = 14
OpenEntityEvent = 15
CloseEntityEvent = 16
FileEvent = "FileEvent"
NewFileEvent = 17
DeleteFileEvent = 18
UpdateFileEvent = 19
TableEvent = "TableEvent"
DeleteRowEvent = 22
NewRowEvent = 23
UpdateRowEvent = 24
SecurityEvent = "SecurityEvent"
NewRoleEvent = 25
DeleteRoleEvent = 26
UpdateRoleEvent = 27
ProjectEvent = "ProjectEvent"
EmailEvent = "EmailEvent"

/**
* This function is call
* @param {string} id The string id.
* @param {string} name The name of the event to mananage.
* @returns {EventManager}
* @constructor
*/
var EventManager = function (id, name) {
    if (id == null) {
        return
    }

    /**
     * @property The event id.
     */
    this.id = id

    /**
     * @property The name of the event
     */
    this.name = name
    this.observers = {}

    return this
}

/**
* Attach observer to event.
* @param obeserver The objserver to attach.
* @param eventNumber The event number.
* @param {function} updateFct The function to execute when the event is receive.
*/
EventManager.prototype.attach = function (observer, eventNumber, updateFct) {
    observer.observable = this

    if (observer.id == undefined) {
        // observer need a uuid or an id.
        observer.id = randomUUID()
    }

    if (this.observers[eventNumber] == undefined) {
        this.observers[eventNumber] = []
    }

    var observerExistsForEventNumber = false
    for (var i = 0; i < this.observers[eventNumber].length; i++) {
        if (this.observers[eventNumber][i].id == observer.id) {
            // only on obeserver with the same id are allow.
            observerExistsForEventNumber = true
        }
    }

    if (!observerExistsForEventNumber) {
        this.observers[eventNumber].push(observer)
    }

    if (observer.updateFunctions == undefined){
        observer.updateFunctions = {}
    }

    observer.updateFunctions[this.id + "_" + eventNumber] = updateFct
}

/**
* Detach an observer from event.
* @param obeserver The to detach
* @param eventNumber The event number
*/
EventManager.prototype.detach = function (observer, eventNumber) {
    if (observer.observable != null) {
        observer.observable = null
    }
    if (observer.updateFunctions != undefined) {
        if (observer.updateFunctions[this.id + "_" + eventNumber] != null) {
            delete observer.updateFunctions[this.id + "_" + eventNumber]
            if (Object.keys(observer.updateFunctions).length == 0) {
                this.observers[eventNumber].pop(observer)
            }
        }
    }
}

/**
* When an event is received, the observer callback function is call.
* @param evt The event to dispatch.
*/
EventManager.prototype.onEvent = function (evt) {
    console.log("Event received: ", evt)
    var observers = this.observers[evt.code]
    if (observers != undefined) {
        for (var i = 0; i < observers.length; i++) {
            if (observers[i].updateFunctions != undefined) {
                if (observers[i].updateFunctions[this.id + "_" + evt.code] != null) {
                    observers[i].updateFunctions[this.id + "_" + evt.code](evt, observers[i])
                } else {
                    if (Object.keys(observers[i].updateFunctions).length == 0) {
                        this.observers[eventNumber].pop(observers[i])
                    }
                }
            }
        }
    }
}

/**
* Each event type has its own channel.
* @param id The channel id, most of the time the event type name.
* @returns {EventChannel}
* @constructor
*/
var EventChannel = function (id) {
    // The event id
    this.id = id
    // This map contain
    this.listeners = {}

    return this
}

/**
* Send event to listener.
* @param evt
* @constructor
*/
EventChannel.prototype.BroadcastEvent = function (evt) {
    for (var l in this.listeners) {
        var listener = this.listeners[l]
        listener.onEvent(evt)
    }
}

/**
* A singleton used to receive and send event. Use it instead of
* directly use event listener and channel.
* @returns {EventHandler}
* @constructor
*/
var EventHandler = function () {
    /**
     * @property channels The channel 
     */
    this.channels = {}

    return this
}

// Open listening channel,

/**
* Append a new event listner.
* @param listener The listener to append.
* @param callback The function to call when an event happen.
*/
EventHandler.prototype.AddEventManager = function (listener, callback) {
    /* Add it to the local event listener **/
    if (this.channels[listener.name] == undefined) {
        this.channels[listener.name] = new EventChannel(listener.name)
    }
    // append the listener
    this.channels[listener.name].listeners[listener.id] = listener

    /* Append to the remote event listener **/
    // Create a request
    var p1 = new RpcData({ "name": "name", "type": 2, "dataBytes": utf8_to_b64(listener.name) });

    var params = new Array();
    params[0] = p1;
    // Register this listener to the server.
    var rqst = new Request(randomUUID(), server.conn, "RegisterListener", params,
        // Progress callback
        function () { },
        // Success callback
        function () {
            // I will call the success call back function
            callback();
        },
        // Error callback.
        function () {

        });
    rqst.send();
}

/**
* Remove the listener and close the channel if is empty
* @param listener The listener to remove.
* @param callback The function to call when an event happen.
*/
EventHandler.prototype.RemoveEventManager = function (listener, callback) {
    /* Delete the local listener **/
    if (this.channels[listener.name] != undefined) {
        if (this.channels[listener.name].listeners[listener.id] != undefined) {
            delete this.channels[listener.name].listeners[listener.id]
        }
        if (Object.keys(this.channels[listener.name]).length == 0) {
            delete this.channels[listener.name]
        }
    }

    /* Delete the remote listener **/
    // Create a request
    var p1 = new RpcData({ "name": "name", "type": 2, "dataBytes": utf8_to_b64(listener.name) });

    var params = new Array();
    params[0] = p1;

    // Register this listener to the server.
    var rqst = new Request(randomUUID(), server.conn, "UnregisterListener", params, function () {
        // I will call the success call back function
        callback();
    });
    rqst.send();
}

/*
* Server side script
*/
function AppendEventFilter(filter, eventType) {
    server.GetEventManager().AppendEventFilter(filter, eventType, messageId, sessionId)
}

/**
* Append a new filter to a listener
* @param {string} filter The filter to append
* @param {string} eventType The name of the channel. Usually the eventType
* @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
* @param {function} errorCallback In case of error.
* @param {object} caller A place to store object from the request context and get it back from the response context.
*/
EventHandler.prototype.appendEventFilter = function (filter, eventType, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(filter, "STRING", "filter"))
    params.push(createRpcData(eventType, "STRING", "eventType"))

    // Call it on the server.
    server.executeJsFunction(
        AppendEventFilter.toString(), // The function to execute remotely on server
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

/**
* Broadcast an event localy over a given channel, usualy the event type..
* @param evt The event to broadcast locally.
* var evt = {"code":OpenEntityEvent, "name":FileEvent, "dataMap":{"fileInfo": file}}
* server.eventHandler.BroadcastEvent(evt)
*/
EventHandler.prototype.BroadcastEvent = function (evt) {
    var channel = this.channels[evt.name]
    if (channel != undefined) {
        channel.BroadcastEvent(evt)
    }
}

/*
* Server side script
*/
function BroadcastEventData(evtNumber, eventType, eventDatas) {
    // Call the method.
    server.GetEventManager().BroadcastEventData(evtNumber, eventType, eventDatas, messageId, sessionId)
}

/**
* Broadcast event over the network.
* @param {int} evtNumber The event number.
* @param {string} evtType The event type.
* @param {MessageData} eventDatas An array of Message Data structures.
* Here is an example To send a file open event over the network.
* var entityInfo = {"TYPENAME":"Server.MessageData", "Name":"entityInfo", "Value":file.stringify()}
* server.eventHandler.broadcastEventData(OpenEntityEvent, EntityEvent, [entityInfo], function(){}, function(){}, undefined) 
 */
EventHandler.prototype.broadcastEventData = function (evtNumber, eventType, eventDatas, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(new RpcData({ "name": "evtNumber", "type": 1, "dataBytes": utf8_to_b64(evtNumber) }))
    params.push(new RpcData({ "name": "eventType", "type": 2, "dataBytes": utf8_to_b64(eventType) }))
    params.push(new RpcData({ "name": "eventDatas", "type": 4, "dataBytes": utf8_to_b64(JSON.stringify(eventDatas)) }))

    // Call it on the server.
    server.executeJsFunction(
        BroadcastEventData.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            //console.log(result)
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            console.log(errMsg)
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}
