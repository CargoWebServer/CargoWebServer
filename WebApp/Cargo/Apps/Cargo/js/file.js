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
 * @fileOverview Contain the defintion of the FileManager class.
 * @author Dave Courtois
 * @version 1.0
 */

/**
 * The FileManager contain functionality to work with files.
  *@constructor
 * @extends EventManager
 */
var FileManager = function () {

    if (server == undefined) {
        return
    }

    EventManager.call(this, FileEvent)

    return this
}

FileManager.prototype = new EventManager(null);
FileManager.prototype.constructor = FileManager;


FileManager.prototype.RegisterListener = function () {
    // Append to the event handler.
    server.eventHandler.AddEventManager(this,
        // callback
        function () {
            console.log("Listener registered!!!!")
        }
    )
}

/*
 * Sever side code.
 */
function CreateDir(dirName, dirPath) {
    var dir = null
    dir = server.GetFileManager().CreateDir(dirName, dirPath, messageId, sessionId)
    return dir
}

/**
 * Create a new directory on the server.
 * @param {string} dirName The name of the new directory.
 * @param {string} dirPath The path of the parent of the new directory.
 * @param {function} successCallback The function to execute in case of success
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
FileManager.prototype.createDir = function (dirName, dirPath, successCallback, progressCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(dirName, "STRING", "dirName"))
    params.push(createRpcData(dirPath, "STRING", "dirPath"))

    server.executeJsFunction(
        CreateDir.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Keep track of the file transfert.
            caller.progressCallback(index, total, caller.caller)
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
        { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * server side code.
 */
function CreateFile(filename, filepath, thumbnailMaxHeight, thumbnailMaxWidth, dbfile) {
    var file = null
    file = server.GetFileManager().CreateFile(filename, filepath, thumbnailMaxHeight, thumbnailMaxWidth, dbfile, messageId, sessionId)
    return file
}

/**
 * Create a new file on the server.
 * @param {string} filename The name of the file to create.
 * @param {string} filepath The path of the directory where to create the file.
 * @param {string} filedata The data of the file.
 * @param {int} thumbnailMaxHeight The maximum height size of the thumbnail associated with the file (keep the ratio).
 * @param {int} thumbnailMaxWidth The maximum width size of the thumbnail associated with the file (keep the ratio).
 * @param {bool} dbFile If it set to true the file will be save on the server local object store, otherwize a file on disck will be created.
 * @param {function} progressCallback The function is call when chunk of response is received.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
FileManager.prototype.createFile = function (filename, filepath, filedata, thumbnailMaxHeight, thumbnailMaxWidth, dbFile, successCallback, progressCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    // The file data (filedata) will be upload with the http protocol...
    params.push(createRpcData(filename, "STRING", "filename"))
    params.push(createRpcData(filepath, "STRING", "filepath"))
    params.push(createRpcData(thumbnailMaxHeight, "INTEGER", "thumbnailMaxHeight"))
    params.push(createRpcData(thumbnailMaxWidth, "INTEGER", "thumbnailMaxWidth"))
    params.push(createRpcData(dbFile, "BOOLEAN", "dbFile"))

    // Here I will create a new data form...
    var formData = new FormData()
    formData.append("multiplefiles", filedata, filename)

    // Use the post function to upload the file to the server.
    var xhr = new XMLHttpRequest()
    xhr.open('POST', '/uploads', true)

    // In case of error or success...
    xhr.onload = function (params, xhr) {
        return function (e) {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    console.log(xhr.responseText);
                    // Here I will create the file...
                    server.executeJsFunction(
                        CreateFile.toString(), // The function to execute remotely on server
                        params, // The parameters to pass to that function
                        function (index, total, caller) { // The progress callback
                            // Keep track of the file transfert.
                            caller.progressCallback(index, total, caller.caller)
                        },
                        function (result, caller) {
                            console.log(result)
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
                        { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback } // The caller
                    )
                } else {
                    console.error(xhr.statusText);
                }
            }
        }
    } (params, xhr)

    // now the progress event...
    xhr.upload.onprogress = function (progressCallback, caller) {
        return function (e) {
            if (e.lengthComputable) {
                progressCallback(e.loaded, e.total, caller)
            }
        }
    } (progressCallback, caller)

    xhr.send(formData);
}

/*
 * server side code.
 */
function GetFileByPath(path) {
    var file = null
    file = server.GetFileManager().GetFileByPath(path, messageId, sessionId)
    return file
}

/**
 * Retreive a file with a given id
 * @param {string} id The file id.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
FileManager.prototype.downloadFile = function (path, fileName, mimeType, progressCallback, successCallback, errorCallback, caller) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', path + '/' + fileName, true);
    xhr.responseType = 'blob';

    xhr.onload = function (successCallback) {
        return function (e) {
            if (this.status == 200) {
                // Note: .response instead of .responseText
                var blob = new Blob([this.response], { type: mimeType });

                // I will read the file as data url...
                var reader = new FileReader();
                reader.onload = function (successCallback, caller) {
                    return function (e) {
                        var dataURL = e.target.result;
                        // return the success callback with the result.
                        successCallback(dataURL, caller)
                    }
                } (successCallback, caller)
                reader.readAsDataURL(blob);
            }
        }
    } (successCallback, caller)

    xhr.onprogress = function (progressCallback, caller) {
        return function (e) {
            progressCallback(e.loaded, e.total, caller)

        }
    } (progressCallback, caller)

    xhr.send();
}

/**
 * Retreive a file with a given id
 * @param {string} id The file id.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
FileManager.prototype.getFileByPath = function (path, progressCallback, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(path, "STRING", "path"))

    server.executeJsFunction(
        GetFileByPath.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Keep track of the file transfert.
            caller.progressCallback(index, total, caller.caller)
        },
        function (result, caller) {
            var file = new CargoEntities.File()
            file.init(result[0])
            server.entityManager.setEntity(file)
            caller.successCallback(file, caller.caller)
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
function OpenFile(fileId) {
    var file = null
    file = server.GetFileManager().OpenFile(fileId, messageId, sessionId)
    return file
}

/**
 * Open a file with a given id
 * @param {string} id The file id.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
FileManager.prototype.openFile = function (fileId, progressCallback, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(fileId, "STRING", "fileId"))

    server.executeJsFunction(
        OpenFile.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Keep track of the file transfert.
            caller.progressCallback(index, total, caller.caller)
        },
        function (result, caller) {
            var file = new CargoEntities.File()
            file.init(result[0])
            server.entityManager.setEntity(file)
            caller.successCallback(file, caller.caller)
            var evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": file } }
            server.eventHandler.BroadcastEvent(evt)
        },
        function (errMsg, caller) {
            // display the message in the console.
            console.log(errMsg)
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message.
            server.errorHandler.onError(errMsg)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * sever side code.
 */
function GetMimeTypeByExtension(fileExtension) {
    var file = null
    file = server.GetFileManager().GetMimeTypeByExtension(fileExtension, messageId, sessionId)
    return file
}

/**
 * Retreive the mime type information from a given extention.
 * @param {string} fileExtension The file extention ex. txt, xls, html, css
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
FileManager.prototype.getMimeTypeByExtension = function (fileExtension, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(fileExtension, "STRING", "fileExtension"))

    server.executeJsFunction(
        GetMimeTypeByExtension.toString(), // The function to execute remotely on server
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
 * sever side code.
 */
function IsFileExist(fileName, filePath) {
    var isFileExist
    isFileExist = server.GetFileManager().IsFileExist(fileName, filePath)
    return isFileExist
}

/**
 * Test if a given file exist.
 * @param {string} filename The name of the file to create.
 * @param {string} filepath The path of the directory where to create the file.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
FileManager.prototype.isFileExist = function (fileName, filePath, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(filename, "STRING", "filename"))
    params.push(createRpcData(filepath, "STRING", "filepath"))

    server.executeJsFunction(
        IsFileExist.toString(), // The function to execute remotely on server
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
 * sever side code.
 */
function DeleteFile(uuid) {
    var err
    err = server.GetFileManager().DeleteFile(uuid, messageId, sessionId)
    return err
}

/**
 * Remove a file with a given uuid.
 * @param {string} uuid The file uuid.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
FileManager.prototype.deleteFile = function (uuid, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(uuid, "STRING", "uuid"))

    server.executeJsFunction(
        DeleteFile.toString(), // The function to execute remotely on server
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