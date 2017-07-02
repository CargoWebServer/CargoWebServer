var WorkflowManager = function () {

    if (server == undefined) {
        return
    }

    EventHub.call(this, BpmnEvent)

    /*
     * The list of all bpmn element.
     */
    this.bpmnElements = {}

    return this
}

WorkflowManager.prototype = new EventHub(null)
WorkflowManager.prototype.constructor = WorkflowManager

/* The event handling **/
WorkflowManager.prototype.onEvent = function (evt) {

    EventHub.prototype.onEvent.call(this, evt)
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Definitions related functions
/////////////////////////////////////////////////////////////////////////////////////////////

/*
 * Import all definition from the definitions directory on the server.
 */
WorkflowManager.prototype.importXmlBpmnDefinitions = function (content, successCallback, errorCallback, caller) {
    // server is the client side singleton...
    var params = []
    params.push(createRpcData(content, "STRING", "content"))

    // Call it on the server.
    server.executeJsFunction(
        "WorkflowManagerImportXmlBpmnDefinitions", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
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

/*
 * Get the list of all definition ids...
 */
WorkflowManager.prototype.getDefinitionsIds = function (successCallback, errorCallback, caller) {
    // server is the client side singleton...
    var params = []

    // Call it on the server.
    server.executeJsFunction(
        "WorkflowManagerGetDefinitionsIds", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            initDefinitions(result[0])
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


/*
 * Get the defintions with a given id.
 */
WorkflowManager.prototype.getDefinitionsById = function (id, successCallback, errorCallback, caller) {
    // server is the client side singleton...
    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    // Call it on the server.
    server.executeJsFunction(
        "WorkflowManagerGetDefinitionsById", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            console.log(result)
            initDefinitions(result[0])
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

function initDefinitions(definitions) {

    for (var i = 0; i < definitions.M_BPMNDiagram.length; i++) {
        // Here I will use function insted of reference, so there will no circulare ref at serialysation time...
        definitions.M_BPMNDiagram[i].getParentDefinitions = function (definitions) {
            return function () {
                return definitions
            }
        } (definitions)

        for (var j = 0; j < definitions.M_BPMNDiagram[i].M_BPMNPlane.M_DiagramElement.length; j++) {
            definitions.M_BPMNDiagram[i].M_BPMNPlane.M_DiagramElement[j].getParentPlane = function (parentPlane) {
                return function () {
                    return parentPlane
                }
            } (definitions.M_BPMNDiagram[i].M_BPMNPlane)

            definitions.M_BPMNDiagram[i].M_BPMNPlane.M_DiagramElement[j].getParentDiagram = function (parentDiagram) {
                return function () {
                    return parentDiagram
                }
            } (definitions.M_BPMNDiagram[i])
        }
    }
}

/*
 * Return the list of all definitions on the server.
 */
WorkflowManager.prototype.getAllDefinitions = function (successCallback, errorCallback, caller) {
    // server is the client side singleton...
    server.entityManager.getEntityPrototypes("BPMN20",
        function (result, caller) {
            // Set the variables.
            var successCallback = caller.successCallback
            var errorCallback = caller.errorCallback
            var caller = caller.caller
            var params = []
            // Call it on the server.
            server.executeJsFunction(
                "WorkflowManagerGetAllDefinitions", // The function to execute remotely on server
                params, // The parameters to pass to that function
                function (index, total, caller) { // The progress callback
                    // Nothing special to do here.
                },
                function (result, caller) {
                    var entities = []
                    if (result[0] != undefined) {
                        for (var i = 0; i < result[0].length; i++) {
                            var entity = eval("new " + result[0][i].TYPENAME + "(caller.prototype)")
                            entity.initCallback = function (entities, count, caller) {
                                return function (entity) {
                                    entities.push(entity)
                                    initDefinitions(entity)
                                    server.entityManager.setEntity(entity)
                                    if (count == entities.length) {
                                        caller.successCallback(entities, caller.caller)
                                    }
                                }
                            } (entities, result[0].length, caller)

                            entity.init(result[0][i])
                        }
                    }
                    if (result[0] == null) {
                        caller.successCallback(entities, caller.caller)
                    }
                },
                function (errMsg, caller) {
                    console.log(errMsg)
                    server.errorManager.onError(errMsg)
                    caller.errorCallback(errMsg, caller.caller)
                }, // Error callback
                { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
            )
        },
        function () { },
        { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller })


}


////////////////////////////////////////////////////////////////////////////////
// Runtime section
////////////////////////////////////////////////////////////////////////////////

/*
 * Get the process entity.
 */
WorkflowManager.prototype.getDefinitionInstances = function (definitions, successCallback, errorCallback, caller) {
    // server is the client side singleton...
    var params = []
    params.push(createRpcData(definitions.UUID, "STRING", "definitionsUUID"))

    // Call it on the server.
    server.executeJsFunction(
        "WorkflowManagerGetDefinitionInstances", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {

            var entities = []
            if (result[0] != undefined) {
                for (var i = 0; i < result[0].length; i++) {
                    var entity = eval("new " + result[0][i].TYPENAME + "(caller.prototype)")
                    if (i == result[0].length - 1) {
                        entity.initCallback = function (caller) {
                            return function (entity) {
                                server.entityManager.setEntity(entity)
                                caller.successCallback(entities, caller.caller)
                            }
                        } (caller)
                    } else {
                        entity.initCallback = function (entity) {
                            server.entityManager.setEntity(entity)
                        }
                    }

                    var definitions = caller.definitions

                    // Here I will link the definition itself...
                    entity.getDefinitions = function (definitions) {
                        return function () {
                            return definitions
                        }
                    } (definitions)

                    // push the entitie before init it...
                    entities.push(entity)

                    // call init...
                    entity.init(result[0][i])

                }
            }
            if (result[0] == null) {
                caller.successCallback(entities, caller.caller)
            }
        },
        function (errMsg) {
            console.log(errMsg)
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "definitions": definitions, "errorCallback": errorCallback } // The caller
    )
}


/*
 * Create a new item definition instance.
 */
WorkflowManager.prototype.newItemAwareElementInstance = function (bpmnElementId, data, successCallback, errorCallback, caller) {

    // server is the client side singleton...
    var params = []
    params.push(createRpcData(bpmnElementId, "STRING", "bpmnElementId"))
    params.push(createRpcData(data, "BYTES", "data"))

    // Call it on the server.
    server.executeJsFunction(
        "WorkflowManagerNewItemAwareElementInstance", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg) {
            console.log(errMsg)
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/**
 * Start a new process instance.
 */
WorkflowManager.prototype.startProcess = function (processUUID, eventData, eventDefinitionData, successCallback, errorCallback, caller) {

    // server is the client side singleton...
    var params = []
    params.push(createRpcData(processUUID, "STRING", "processUUID"))
    params.push(createRpcData(eventData, "JSON_STR", "params"))
    params.push(createRpcData(eventDefinitionData, "JSON_STR", "params"))

    // Call it on the server.
    server.executeJsFunction(
        "WorkflowManagerStartProcess", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg) {
            console.log(errMsg)
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/**
 * Get the list of active process instances for a bpmn process with given uuid.
 */
WorkflowManager.prototype.getActiveProcessInstances = function (uuid, successCallback, errorCallback, caller) {

    // server is the client side singleton...
    var params = []
    params.push(createRpcData(uuid, "STRING", "uuid"))

    // Call it on the server.
    server.executeJsFunction(
        "WorkflowManagerGetActiveProcessInstances", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            var entities = []
            if (result[0] != undefined) {
                for (var i = 0; i < result[0].length; i++) {
                    var entity = eval("new " + result[0][i].TYPENAME + "(caller.prototype)")
                    entity.initCallback = function (entities, count, caller) {
                        return function (entity) {
                            entities.push(entity)
                            server.entityManager.setEntity(entity)
                            if (count == entities.length) {
                                caller.successCallback(entities, caller.caller)
                            }
                        }
                    } (entities, result[0].length, caller)
                    entity.init(result[0][i])
                }
            }
            if (result[0] == null) {
                caller.successCallback(entities, caller.caller)
            }
        },
        function (errMsg) {
            console.log(errMsg)
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/**
 * Run an activity with a given uuid.
 */
WorkflowManager.prototype.activateActivityInstance = function (uuid, successCallback, errorCallback, caller) {

    // server is the client side singleton...
    var params = []
    params.push(createRpcData(uuid, "STRING", "uuid"))

    // Call it on the server.
    server.executeJsFunction(
        "WorkflowManagerActivateActivityInstance", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            // Nothing to do here.
        },
        function (errMsg) {
            console.log(errMsg)
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}