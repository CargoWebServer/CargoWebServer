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
 * @fileOverview Network messages handler.
 * @author Dave Courtois
 * @version 1.0
 */

////////////////////////////////////////////////////////////////////////////////
//                      Network stuff..
////////////////////////////////////////////////////////////////////////////////
/**
 * @constant {int} This is the maximum message size.
 */
MAX_MSG_SIZE = 17740;

/**
 *  @constant {int}  Double data type identifier.
 */
Data_DOUBLE = 0
/**
 *  @constant {int}  Interger data type identifier.
 */
Data_INTEGER = 1
/**
 *  @constant {int}  String data type identifier.
 */
Data_STRING = 2
/**
 *  @constant {int}  Bytes data type identifier.
 */
Data_BYTES = 3
/**
 *  @constant {int}  Struct (JSON) data type identifier.
 */
Data_JSON_STR = 4
/**
 *  @constant {int}  Boolean data type identifier.
 */
Data_BOOLEAN = 5

/**
 *  @constant {int}  Request message type identifier.
 */
Msg_REQUEST = 0;

/**
 *  @constant {int}  Response message type identifier.
 */
Msg_RESPONSE = 1;

/**
 *  @constant {int}  Error message type identifier.
 */
Msg_ERROR = 2;

/**
 *  @constant {int}  Event message type identifier.
 */
Msg_EVENT = 3;

/**
 *  @constant {int}  Transfert message type identifier.
 */
Msg_TRANSFER = 4;

// Contain the pending request waiting to be ask back.
var pendingRequest = {};

// Map of array of message chunk.
var pendingMessage = {}

// OAuth2 dialogs.
// That dialog is use to give access to application to protected user
// ressources. With it You can access to cargo, github, google, facebook user ressources
var oauth2Dialog

////////////////////////////////////////////////////////////////////////////////
//                      The proto buffer objects.
////////////////////////////////////////////////////////////////////////////////

// This is the content of the file rpc.proto in form of string.
var protobufSrc = "package com.mycelius.message;";
protobufSrc += "message Data {required string name = 1;required bytes dataBytes = 2;"
protobufSrc += "enum DataType{ DOUBLE = 0;INTEGER = 1;STRING = 2;BYTES = 3;JSON_STR = 4;BOOLEAN = 5;}required DataType type = 3 [default = BYTES]; optional string typeName = 4;}"
protobufSrc += "message Request{required string method = 1;repeated Data params = 2;required string id = 3;}message Response{repeated Data results = 1;required string id = 2;}"
protobufSrc += "message Error{required int32 code = 1;required string message = 2;required string id = 3;optional bytes data = 4;}"
protobufSrc += "message Event{required int32 code = 1;required string name = 2;repeated Data evtData = 3;}"
protobufSrc += "message Message{enum MessageType {REQUEST = 0;RESPONSE = 1;ERROR = 2;EVENT=3;TRANSFER=4;} required MessageType type = 1 [default = ERROR];required sint32 index = 2;required int32 total = 3;optional Request rqst = 4;optional Response rsp = 5;optional Error err = 6;optional Event evt = 7; optional bytes data = 8;optional string id = 9;}"

var root = dcodeIO.ProtoBuf.protoFromString(protobufSrc).build(); //.protoFromFile("proto/rpc.proto").build();

// Use these object to create and serialyse message.
var RpcMessage = root.com.mycelius.message.Message;
var RpcRequest = root.com.mycelius.message.Request;
var RpcData = root.com.mycelius.message.Data;
var RpcResponse = root.com.mycelius.message.Response;
var RpcError = root.com.mycelius.message.Error;
var RpcEvent = root.com.mycelius.message.Event;


////////////////////////////////////////////////////////////////////////////////
//                      Web socket initialisation
////////////////////////////////////////////////////////////////////////////////

/*
 * Initialisation of the web socket handler.
 */
function initConnection(adress, onOpenCallback, onCloseCallback) {
    if ("WebSocket" in window) {
        // Let us open a web socket
        var connection = new WebSocket(adress);
        connection.onopen = function () {
            // Web Socket is connected, send data using send()
            //console.log("The web socket is open and ready to use.");
            if (onOpenCallback != undefined) {
                onOpenCallback();
            }
        };

        connection.onmessage = function (evt) {
            var arrayBuffer;
            var fileReader = new FileReader();
            var self = this;

            fileReader.onload = function () {
                arrayBuffer = this.result;
                server.handleMessage(self, arrayBuffer);
            };

            fileReader.readAsArrayBuffer(evt.data);
        };

        connection.onclose = function () {
            // websocket is closed.
            if (onCloseCallback != undefined) {
                onCloseCallback();
            }
        };


    }
    else {
        console.log("WebSocket NOT supported by your Browser!");
    }

    return connection;
}

/**
 * Simple data representation use to exange data between participants.
 */
var MessageData = function (name, value) {
    // Hint about type name.
    this.TYPENAME = "Server.MessageData"

    /**
     * @param {string} Name The message identification.
     */
    this.Name = name

    /**
     * @param {string} Value The message content.
     */
    this.Value = value

    return this
}

////////////////////////////////////////////////////////////////////////////////
//                     Message base class
////////////////////////////////////////////////////////////////////////////////

/**
 * Wrapping RPC message.
 * @param id The message id.
 * @param conn The web socket connection reference.
 * @param progressCallback The progress callback
 * @param successCallback The success callback
 * @param errorCallback The error callback
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
var Message = function (id, conn, progressCallback, successCallback, errorCallback, caller) {
    if (id == undefined) {
        return;
    }

    // The id is generate
    this.id = id;
    this.conn = conn; // Must have a send method. WebSocket and WebRTC has one.
    this.index = -1
    this.total = 1
    this.data = null

    // The link to the caller
    if (caller != null) {
        this.caller = caller
    }

    // Call back function.
    if (progressCallback != null) {
        this.progressCallback = progressCallback;
    }

    if (successCallback != null) {
        this.successCallback = successCallback;
    }

    if (errorCallback != null) {
        this.errorCallback = errorCallback;
    }
    return this;
};

// Must be implement by child's.
Message.prototype.getRpcMessageData = function () {
    var msg = new RpcMessage({ "id": this.id, "type": Msg_TRANSFER, "index": this.index, "total": this.total, "data": utf8_to_b64(this.data) });
    return msg.toArrayBuffer();
}

/**
 * Parse the data contain in the data/param/result map.
 * @param {bytes} The binary data sent by the server.
 */
Message.prototype.parseData = function (data) {
    if (data.type == Data_DOUBLE) {
        var begin = data.dataBytes.offset;
        var end = data.dataBytes.limit;
        var val = new Float64Array(data.dataBytes.buffer.slice(begin, end))[0];
        return val
    } else if (data.type == Data_INTEGER) {
        var begin = data.dataBytes.offset;
        var end = data.dataBytes.limit;
        var val = new Int32Array(data.dataBytes.buffer.slice(begin, end))[0];
        return val
    } else if (data.type == Data_STRING) {
        var begin = data.dataBytes.offset;
        var end = data.dataBytes.limit;
        var str = new StringView(data.dataBytes.view.subarray(begin, end)).toString();
        return str;
    } else if (data.type == Data_BYTES) {
        var begin = this.data.dataBytes.offset;
        var end = this.data.dataBytes.length;
        var bytes = this.data.dataBytes.view.subarray(begin, end)
        return bytes
    } else if (data.type == Data_JSON_STR) {
        var begin = data.dataBytes.offset;
        var end = data.dataBytes.limit;
        var str = new StringView(data.dataBytes.view.subarray(begin, end)).toString();
        return JSON.parse(str);
    } else if (data.type == Data_BOOLEAN) {
        var begin = data.dataBytes.offset;
        var end = data.dataBytes.limit;
        var str = new StringView(data[i].dataBytes.view.subarray(begin, end)).toString();
        var val = str == "true" ? true : false
        return val
    }

    return null
}

/**
 * Send message to the server.
 */
Message.prototype.send = function () {

    // First i will create the array for the array buffer.
    var bytes = new Uint8Array(this.getRpcMessageData())

    // Control the size of the message.
    if ((this.total == 1 && this.index == -1) || (this.total > 1 && this.index > -1)) {
        this.conn.send(bytes);
    } else if (this.total > 1) {
        // Bytes will contain the message to send to the server,
        // because it can't be send in one pass I will get it's bytes and
        // chunk it and put it inside a new message with the same id.
        var size = Math.round((bytes.byteLength / MAX_MSG_SIZE) + .5);
        for (var i = 0; i < size; i++) {
            var fragment;
            if (MAX_MSG_SIZE * (i + 1) < bytes.length) {
                fragment = bytes.subarray(MAX_MSG_SIZE * i, MAX_MSG_SIZE * (i + 1));
            } else {
                fragment = bytes.subarray(MAX_MSG_SIZE * i, bytes.length);
            }
            var self = this
            var msg = new Message(this.id, this.conn,
                // progress call back
                self.progressCallback,
                function (id) {
                    // Call the progression call back.
                    if (this.progressCallback != null) {
                        if (pendingMessage[id] != undefined) {
                            this.progressCallback(pendingMessage[id].length)
                        }
                    }
                    return false;
                },
                // function error call back
                self.errorCallback,
                self.caller
            );

            if (pendingMessage[this.id] == undefined) {
                // This is the message to be sent.
                pendingMessage[this.id] = new Array();
            }

            // Set the msg data.
            msg.data = Uint8ToBase64(fragment)
            msg.index = i
            msg.total = self.total
            pendingMessage[this.id][i] = msg;
        }

        // Process the next message for that id.
        processPendingMessage(this.id);
    }
};

/*
 * That function is call recursively until there no more pending message.
 */
function processPendingMessage(msgId) {
    if (pendingMessage[msgId].length > 0) {
        var msg = pendingMessage[msgId].shift();
        if (msg.progressCallback != null) {
            msg.progressCallback(msg.index, msg.total, msg.caller)
        }
        msg.send();
    } else {
        // Remove the entry
        delete pendingMessage[msgId]
    }
}

////////////////////////////////////////////////////////////////////////////////
//                     Request class
////////////////////////////////////////////////////////////////////////////////

/**
 * A class that contain value to be sent to the server.
 * @param id The message id.
 * @param conn The web socket connection reference.
 * @param method The method to be called.
 * @param params The list of parameters to pass to the methode.
 * @param progressCallback The progress callback
 * @param successCallback The success callback
 * @param errorCallback The error callback
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 * @returns {*}
 * @constructor
 * @extends Message
 */
var Request = function (id, conn, method, params, progressCallback, successCallback, errorCallback, caller) {
    Message.call(this, id, conn, progressCallback, successCallback, errorCallback, caller);

    this.method = method;
    this.params = params;
    this.paramsMap = {}

    for (var i = 0; i < params.length; ++i) {
        this.paramsMap[params[i].name] = params[i];
    }

    if (successCallback != null) {
        pendingRequest[this.id] = this; // Set the request on the map.
    } else if (id != "") {
        pendingRequest[id] = this; // Insert in the pending request map.
    }

    return this;
};

// inherit from request object.
Request.prototype = new Message(null);
Request.prototype.constructor = Request;

/**
 * Create the protobuffer request object. 
 */
Request.prototype.getRpcMessageData = function () {
    var rqst = new RpcRequest({ "id": this.id, "method": this.method, "params": this.params });

    // Calculate the total number of message here.
    var msg = new RpcMessage({ "id": this.id, "type": Msg_REQUEST, "rqst": rqst, "index": 1, "total": 1 });

    var bytes = msg.toArrayBuffer();
    msg.total = Math.ceil(bytes.byteLength / MAX_MSG_SIZE)
    this.total = msg.total

    return bytes
}

// global var use in process
var w

/**
 * Read and parse data from the proto format to js format.
 */
Request.prototype.process = function () {
    for (var i = 0; i < this.params.length; ++i) {
        this.paramsMap[this.params[i].name] = this.parseData(this.params[i])
    }

    // The ping request is a special case of request.
    if (this.method == "Ping") {
        // Send a ping response.
        var response = new Response(this.id, this.conn, [], null, null, null);
        response.send()
    }

    if (this.method == "OAuth2Authorization") {
        // Here I will create the  dialog.
        if (oauth2Dialog == null) {
            // Here We receice an OAuth request.
            var href = this.paramsMap["authorizationLnk"]
            oauth2Dialog = new Dialog("oauth2Dialog", document.getElementsByTagName("body")[0], true, "Authorization")
            oauth2Dialog.footer.element.style.display = "none"

            // The content will be the html receive from the request.
            if (href.startsWith("https://www.facebook.com/dialog/oauth")
                || href.startsWith("https://accounts.google.com")) {
                var lnk = oauth2Dialog.content.appendElement({ "tag": "a", "href": "#" }).down()
                // Set specif oauth provider information here...
                if (href.startsWith("https://accounts.google.com")) {
                    lnk.element.innerHTML = "Google authentication"
                } else if (href.startsWith("https://www.facebook.com/dialog/oauth")) {
                    lnk.element.innerHTML = "Facebook authentication"
                } else {
                    lnk.element.innerHTML = "OAuth2 authentication"
                }

                // TODO append other provider here.

                lnk.element.onclick = function (href) {
                    return function () {
                        // Set the window...
                        w = window.open(href, '_blank');
                    }
                } (href)
            } else {
                var content = new Element(oauth2Dialog.content, { "tag": "iFrame" })
                content.element.contentWindow.document.location.href = href;

                // Set the width and heigth of the dialog to fit the content.
                oauth2Dialog.div.element.style.width = content.element.offsetWidth + "px"
                oauth2Dialog.div.element.style.heigth = content.element.offsetHeigth + "px"
            }
        }
        oauth2Dialog.setCentered()
    } else if (this.method == "closeAuthorizeDialog") {
        // Here I will simply destroy the oauth2 Dialog
        oauth2Dialog.close()
        delete oauth2Dialog
        // Send an empty response, the information will be sent via a form inside the content.
        var response = new Response(this.id, this.conn, [], null, null, null);
        response.send()
    } else if (this.method == "finalyseAuthorize") {

        var idTokenUuid = this.paramsMap["idTokenUuid"]
        // So here I will store the idTokenUuid in the local storage and close the issuer window.
        localStorage.setItem("idTokenUuid", idTokenUuid)
        if (w != undefined) {
            w.close()
            // remove the w ref.
            w = null
        }
    }

    // Now I will create the function prototype and try to call it.
    var fn = window[this.method];
    var fnparams = []
    for (var i = 0; i < this.params.length; ++i) {
        // add the param.
        fnparams.push(this.paramsMap[this.params[i].name])
    }

    // is object a function?
    if (typeof fn === "function") {
        var result = fn.apply(null, fnparams);
        // Todo get the response from the funtion and create a response to send back to the server.
    }
}

////////////////////////////////////////////////////////////////////////////////
//                     Response class
////////////////////////////////////////////////////////////////////////////////
/**
 * This is the base class for a response.
 * @param {string} id The message id.
 * @param {bytes} results The result data 
 * @param {object} conn The web socket connection reference.
 * @param {function} progressCallback The progress callback
 * @param {function} successCallback The success callback
 * @param {function} errorCallback The error callback
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 * @returns {*}
 * @constructor
 * @extends Message
 */
var Response = function (id, conn, results, progressCallback, successCallback, errorCallback, caller) {
    Message.call(this, id, conn, progressCallback, successCallback, errorCallback, caller);

    // I will get the results.
    this.resultsMap = {}
    this.results = results
    for (var i = 0; i < results.length; ++i) {
        this.resultsMap[results[i].name] = this.parseData(results[i])
    }

    return this;
};

// inherit from request object.
Response.prototype = new Message(null);
Response.prototype.constructor = Response;

/*
 * Create the protobuffer RpcResponse from information in this class.
 */
Response.prototype.getRpcMessageData = function () {
    var resp = new RpcResponse({ "id": this.id, "results": this.results });
    var msg = new RpcMessage({ "id": this.id, "type": Msg_RESPONSE, "rsp": resp, "index": -1, "total": 1 });
    var bytes = msg.toArrayBuffer();
    msg.total = Math.ceil(bytes.byteLength / MAX_MSG_SIZE)
    this.total = msg.total
    return bytes;
}

/*
 * Call the success callback if the response is completed.
 */
Response.prototype.execute = function (rqst) {
    delete pendingRequest[rqst.id]; 
    if (rqst.successCallback != null) {
        rqst.successCallback(this.id, this.resultsMap, rqst.caller)
    }
}

////////////////////////////////////////////////////////////////////////////////
//                     Event class
////////////////////////////////////////////////////////////////////////////////

// Sample use of server.
// var evt = new EventMsg(randomUUID(), server.conn, SessionEvent, LoginEvent, [], undefined, undefined, undefined)
// evt.dataMap["sessionInfo"] = result
// throw the event.
// server.sessionManager.onEvent(evt)
var EventMsg = function (id, conn, name, code, data, progressCallback, successCallback, errorCallback) {
    Message.call(this, id, conn, progressCallback, successCallback, errorCallback);

    // int value that correspond to the code.
    this.code = code

    // the event name
    this.name = name

    // This contain event specific data.
    this.dataMap = {}

    // the data byte array ready to be sent on the network.
    this.data = data

    for (var i = 0; i < data.length; ++i) {
        this.dataMap[data[i].name] = this.parseData(data[i])
    }
    return this
}

// inherit from request object.
EventMsg.prototype = new Message(null);
EventMsg.prototype.constructor = EventMsg;

// Create the data to be sent over the network
EventMsg.prototype.getRpcMessageData = function () {
    var evt = new RpcEvent({ "code": this.code, "name": this.name, "evtData": this.data });
    var msg = new RpcMessage({ "type": Msg_EVENT, "evt": evt, "index": -1, "total": 1 });
    var bytes = msg.toArrayBuffer();
    msg.total = Math.ceil(bytes.byteLength / MAX_MSG_SIZE)
    this.total = msg.total

    return bytes;
}


////////////////////////////////////////////////////////////////////////////////
//                     Error class
////////////////////////////////////////////////////////////////////////////////
var ErrorMsg = function (id, conn, code, message, data, progressCallback, successCallback, errorCallback) {
    Message.call(this, id, conn, progressCallback, successCallback, errorCallback);

    // The error code
    this.code = code
    // The error message
    this.message = message
    // The error specific data.
    this.data = data

    // This contain error specific data.
    this.dataMap = {}
    if (data != null) {
        // The data is an
        var begin = data.offset;
        var end = data.limit;
        var str = new StringView(data.view.subarray(begin, end)).toString();
        if (isJsonString(str)) {
            var errorObj = JSON.parse(str)
            this.dataMap["errorObj"] = errorObj
        } else {
            this.dataMap["errorStr"] = str
        }
    }
    return this
}

// inherit from request object.
ErrorMsg.prototype = new Message(null);
ErrorMsg.prototype.constructor = ErrorMsg;

// Create the data to be sent over the network
ErrorMsg.prototype.getRpcMessageData = function () {
    var err = new RpcError({ "code": this.code, "message": this.message, "data": this.data });
    var msg = new RpcMessage({ "type": Msg_ERROR, "err": err, "index": 1, "total": 1 });
    var bytes = msg.toArrayBuffer();
    msg.total = Math.ceil(bytes.byteLength / MAX_MSG_SIZE)
    this.total = msg.total

    return bytes;
}

ErrorMsg.prototype.catch = function () {
    var rqst = pendingRequest[this.id]
    if (rqst !== undefined) {
        if (rqst.errorCallback != null) {
            // Call the error callback with this err information.
            rqst.errorCallback(this, rqst.caller)
        }
    }
    // Delete pending request in that case.
    delete pendingRequest[this.id]; // remove the request from the queue,,,
}