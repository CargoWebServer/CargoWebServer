
var nbTests = 0;
var nbTestsAsserted = 0;

function AllFileTestsAsserted_Test() {
    QUnit.test("AllFileTestsAsserted_Test",
        function (assert) {
            assert.ok(nbTests == nbTestsAsserted, "nbTests: " + nbTests + ", nbTestsAsserted: " + nbTestsAsserted);
        })
}

/* FAIL: Pourtant le directory est bien crée
* Une fois que ca marche: essayer multiples endroits
*/
function createDir_Test() {
    nbTests++
    server.fileManager.createDir("createDir_Test", "",
        function (result) {
            QUnit.test("createDir_Test",
                function (result) {
                    return function (assert) {
                        console.log(result)
                        assert.ok(false);
                        nbTestsAsserted++
                    }
                }(result))
        }, function (result) {
        },
        function (result) {
        },
        undefined)
}

/**
 * PAS FINI DE TESTER TOUTES LES FONCTIONS
 */
function fileTests() {
    createDir_Test()

    setTimeout(function () {
        return function () {
            AllFileTestsAsserted_Test()
        }
    }(), 3000);
}


function testDownloadFile() {
    // Test file download...
    server.fileManager.downloadFile("/BrisOutil/Upload/700", "Hydrangeas.jpg", "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
        // Progress
        function (index, total, caller) {

        },
        // success
        function (result, caller) {
            // Create a new Blob object using the
            //response data of the onload object

            //Create a link element, hide it, direct
            //it towards the blob, and then 'click' it programatically
            let a = document.createElement("a");
            a.style = "display: none";
            document.body.appendChild(a);
            //Create a DOMString representing the blob
            //and point the link element towards it

            a.href = result;
            //programatically click the link to trigger the download
            a.click();

            //release the reference to the file by revoking the Object URL
            window.URL.revokeObjectURL(result);
        },
        // error
        function (errMsg, caller) {

        }, undefined)
}





//  EntityManager.prototype.getEntities = function (typeName, storeId, query, offset, limit, orderBy, asc, lazy, progressCallback, successCallback, errorCallback, caller) {
//     // First of all i will get the entity prototype.
//     server.entityManager.getEntityPrototype(typeName, storeId,
//         // The success callback.
//         function (result, caller) {
//             // Set the parameters.
//             var typeName = caller.typeName
//             var storeId = caller.storeId
//             var query = caller.query
//             var successCallback = caller.successCallback
//             var progressCallback = caller.progressCallback
//             var errorCallback = caller.errorCallback
//             var lazy = caller.lazy
//             var caller = caller.caller
//             // Create the list of parameters.
//             var params = []
//             params.push(createRpcData(typeName, "STRING", "typeName"))
//             params.push(createRpcData(storeId, "STRING", "storeId"))
//             params.push(createRpcData(query, "JSON_STR", "query"))
//             params.push(createRpcData(offset, "INTEGER", "offset"))
//             params.push(createRpcData(limit, "INTEGER", "limit"))
//             params.push(createRpcData(orderBy, "JSON_STR", "orderBy", "[]string"))
//             params.push(createRpcData(asc, "BOOLEAN", "asc"))
//             // Call it on the server.
//             server.executeJsFunction(
//                 "EntityManagerGetEntities", // The function to execute remotely on server
//                 params, // The parameters to pass to that function
//                 function (index, total, caller) { // The progress callback
//                     // Keep track of the file transfert.
//                     caller.progressCallback(index, total, caller.caller)
//                 },
//                 function (result, caller) {
//                     var entities = []
//                     if (result[0] == null) {
//                         if (caller.successCallback != undefined) {
//                             caller.successCallback(entities, caller.caller)
//                             caller.successCallback = undefined
//                         }
//                     } else {
//                         var values = result[0];
//                         if (values.length > 0) {
//                             var initEntitiesFct = function (values, caller, entities) {
//                                 var value = values.pop()
//                                 var entity = eval("new " + caller.prototype.TypeName + "()")
//                                 entities.push(entity)
//                                 if (values.length == 0) {
//                                     entity.initCallback = function (caller, entities) {
//                                         return function (entity) {
//                                             server.entityManager.setEntity(entity)
//                                             if (caller.successCallback != undefined) {
//                                                 caller.successCallback(entities, caller.caller)
//                                                 caller.successCallback = undefined
//                                             }
//                                         }
//                                     }(caller, entities)
//                                     entity.init(value, lazy)
//                                 } else {
//                                     entity.initCallback = function (entity) {
//                                         server.entityManager.setEntity(entity)
//                                     }
//                                     entity.init(value, lazy, function (values, caller, entities) {
//                                         return function () {
//                                             initEntitiesFct(values, caller, entities)
//                                         }
//                                     }(values, caller, entities))
//                                 }
//                             }
//                             initEntitiesFct(values, caller, entities)
//                         } else {
//                             if (caller.successCallback != undefined) {
//                                 caller.successCallback(entities, caller.caller)
//                                 caller.successCallback = undefined
//                             }
//                         }
//                     }
//                 },
//                 function (errMsg, caller) {
//                     // call the immediate error callback.
//                     if (caller.errorCallback != undefined) {
//                         caller.errorCallback(errMsg, caller.caller)
//                         caller.errorCallback = undefined
//                     }
//                     // dispatch the message.
//                     server.errorManager.onError(errMsg)
//                 }, // Error callback
//                 { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback, "prototype": result, "lazy": lazy } // The caller
//             )
//         },
//         // The error callback.
//         function (errMsg, caller) {
//             // call the immediate error callback.
//             if (caller.errorCallback != undefined) {
//                 caller.errorCallback(errMsg, caller.caller)
//                 caller.errorCallback = undefined
//             }
//             // dispatch the message.
//             server.errorManager.onError(errMsg)
//         }, { "typeName": typeName, "storeId": storeId, "query": query, "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback, "lazy": lazy })
// }


