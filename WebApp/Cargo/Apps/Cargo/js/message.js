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
 * @fileOverview The definiton of Server class.
 * @author Dave Courtois
 * @version 1.0
 */

/**
 * The email manager can be use to send email or get acces to mailbox.
 * @constructor
 * @extends EventHub
 */
var EmailManager = function (id) {

    if (server == undefined) {
        return
    }
    if (id == undefined) {
        id = randomUUID()
    }

    EventHub.call(this, id, EmailEvent)

    return this
}

EmailManager.prototype = new EventHub(null);
EmailManager.prototype.constructor = EntityManager;

/*
 * Dispatch event.
 */
EmailManager.prototype.onEvent = function (evt) {

    EventHub.prototype.onEvent.call(this, evt)
}

/**
* This little structure is use to keep the cc information.
* @param name  Just a string can be John Doe, mr. Doe etc.
* @param email  The email of the cc.
* @constructor
*/
var CarbonCopy = function (name, email) {
    // Hint about type name.
    this.TYPENAME = "Server.CarbonCopy"

    /**
     * @property {string} Name The name of the carbon copy.
     */
    this.Name = name

    /**
     * @property {string} Mail The email addresse of the carbon copy.
     */
    this.Mail = email

    return this
}

/**
 * This structure is use in the transfer of files use in attach.
 * @param {string} fileName  The name of the file
 * @param {string} fileData  The file data
 * @constructor
 */
var Attachment = function (fileName, fileData) {
    this.TYPENAME = "Server.Attachment"
    /**
     * @property {string} FileName the name of the file attachement. It can also contain the file path on the server side,
     * if is the case the file data will be take from this file.
     */
    this.FileName = fileName
    /**
     * @property {string} FileData The data of the file. If it's null the file data will be taked from the server side.
     */
    this.FileData = fileData
    return this
}

/*
 * Server side function.
 */
function SendEmail(serverId, from_, to, cc, title, msg, attachs, bodyType) {
    var err = null
    err = server.GetSmtpManager().SendEmail(serverId, from_, to, cc, title, msg, attachs, bodyType, messageId, sessionId)
    return err
}

/**
 * This function is use to send a email message to a given addresse.
 * @param serverId The server configuration id
 * @param from  The email of the sender
 * @param to The destination email's list
 * @param cc The carbon copy's.
 * @param title The email title.
 * @param msg The message to send must be html format.
 * @param attachs The list of local file to upload to the server and attach to the message.
 * @param bodyType text/html or text/plain
 * @param {callback} progressCallback The function called when information is transfer from the server to client.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 * @constructor
 */
EmailManager.prototype.sendEmail = function (serverId, from_, to, cc, title, msg, attachs, bodyType, progressCallback, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(serverId, "STRING", "serverId"))
    params.push(createRpcData(from_, "STRING", "from"))
    params.push(createRpcData(to, "STRING", "to"))
    params.push(createRpcData(cc, "JSON_STR", "cc", "Server.CarbonCopy"))
    params.push(createRpcData(title, "STRING", "title"))
    params.push(createRpcData(msg, "STRING", "msg"))
    params.push(createRpcData(attachs, "JSON_STR", "attachs", "Server.Attachment"))
    params.push(createRpcData(bodyType, "STRING", "bodyType"))

    // Call it on the server.
    server.executeJsFunction(
        // The session id on the server.
        SendEmail.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            caller.progressCallback(index, total, caller.caller)
        },
        function (result, caller) {
            // Nothing special to do here.
            caller.successCallback(result)
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