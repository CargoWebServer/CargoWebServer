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
var FileManager = function () {
    if (server == undefined) {
        return
    }
    EventHub.call(this, FileEvent)

    return this
}

FileManager.prototype = new EventHub(null);
FileManager.prototype.constructor = FileManager;

FileManager.prototype.onEvent = function (evt) {
    EventHub.prototype.onEvent.call(this, evt)
}
FileManager.prototype.createDir = function (dirName, dirPath, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(dirName, "STRING", "dirName"))
    params.push(createRpcData(dirPath, "STRING", "dirPath"))

    server.executeJsFunction(
        "FileManagerCreateDir",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("CargoEntities.File", "CargoEntities",
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
                        "FileManagerCreateFile", // The function to execute remotely on server
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
                } else {
                    console.error(xhr.statusText);
                }
            }
        }
    }(params, xhr)
    // now the progress event...
    xhr.upload.onprogress = function (progressCallback, caller) {
        return function (e) {
            if (e.lengthComputable) {
                progressCallback(e.loaded, e.total, caller)
            }
        }
    }(progressCallback, caller)
    xhr.send(formData);
}

FileManager.prototype.deleteFile = function (uuid, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(uuid, "STRING", "uuid"))

    server.executeJsFunction(
        "FileManagerDeleteFile",
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
                }(successCallback, caller)
                reader.readAsDataURL(blob);
            }
        }
    }(successCallback, caller)
    xhr.onprogress = function (progressCallback, caller) {
        return function (e) {
            progressCallback(e.loaded, e.total, caller)
        }
    }(progressCallback, caller)
    xhr.send();
}

FileManager.prototype.getFileByPath = function (path, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(path, "STRING", "path"))

    server.executeJsFunction(
        "FileManagerGetFileByPath",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("CargoEntities.File", "CargoEntities",
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

FileManager.prototype.getMimeTypeByExtension = function (fileExtension, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(fileExtension, "STRING", "fileExtension"))

    server.executeJsFunction(
        "FileManagerGetMimeTypeByExtension",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("Server.MimeType", "Server",
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

FileManager.prototype.isFileExist = function (successCallback, errorCallback, caller) {
    var params = []

    server.executeJsFunction(
        "FileManagerIsFileExist",
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

FileManager.prototype.openFile = function (fileId, progressCallback, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(createRpcData(fileId, "STRING", "fileId"))
    server.executeJsFunction(
        "FileManagerOpenFile", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            caller.progressCallback(index, total, caller.caller)
        },
        function (result, caller) {
            var file = new CargoEntities.File()
            file.init(result[0])
            server.entityManager.setEntity(file)
            caller.successCallback(file, caller.caller)
            var evt = { "code": OpenEntityEvent, "name": FileEvent, "dataMap": { "fileInfo": file } }
            server.eventHandler.broadcastLocalEvent(evt)
        },
        function (errMsg, caller) {
            caller.errorCallback(errMsg, caller.caller)
            server.errorHandler.onError(errMsg)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback } // The caller
    )
}

FileManager.prototype.readCsvFile = function (filePath, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(filePath, "STRING", "filePath"))

    server.executeJsFunction(
        "FileManagerReadCsvFile",
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

FileManager.prototype.readTextFile = function (filePath, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(filePath, "STRING", "filePath"))

    server.executeJsFunction(
        "FileManagerReadTextFile",
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

FileManager.prototype.removeFile = function (filePath, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(filePath, "STRING", "filePath"))

    server.executeJsFunction(
        "FileManagerRemoveFile",
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
