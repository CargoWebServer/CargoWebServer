var WorkflowManager = function (id) {

    if (server == undefined) {
        return
    }
    if (id == undefined) {
        id = randomUUID()
    }

    EventManager.call(this, id, BpmnEvent)

    /*
     * The list of all bpmn element.
     */
    this.bpmnElements = {}

    return this
}

WorkflowManager.prototype = new EventManager(null)
WorkflowManager.prototype.constructor = WorkflowManager

WorkflowManager.prototype.RegisterListener = function () {
    // Append to the event handler...
    server.eventHandler.AddEventManager(this,
        // callback
        function () {
            console.log("Listener registered!!!!")
        }
    )
}

/* The event handling **/
WorkflowManager.prototype.onEvent = function (evt) {

    EventManager.prototype.onEvent.call(this, evt)
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Definitions related functions
/////////////////////////////////////////////////////////////////////////////////////////////

/*
 * Import all definition from the definitions directory on the server.
 */
function ImportXmlBpmnDefinitions(content) {
    err = server.GetWorkflowManager().ImportXmlBpmnDefinitions(content, messageId, sessionId)
    return err
}

WorkflowManager.prototype.importXmlBpmnDefinitions = function (content, successCallback, errorCallback, caller) {
    // server is the client side singleton...
    var params = []
    params[0] = new RpcData({ "name": "content", "type": 2, "dataBytes": utf8_to_b64(content) })

    // Call it on the server.
    server.executeJsFunction(
        ImportXmlBpmnDefinitions.toString(), // The function to execute remotely on server
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
function GetDefinitionsIds() {
    var ids = null
    ids = server.GetWorkflowManager().GetDefinitionsIds(messageId, sessionId)
    return ids
}

WorkflowManager.prototype.getDefinitionsIds = function (successCallback, errorCallback, caller) {
    // server is the client side singleton...
    var params = []

    // Call it on the server.
    server.executeJsFunction(
        GetDefinitionsIds.toString(), // The function to execute remotely on server
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
function GetDefinitionsById(id) {
    var definition = null
    definition = server.GetWorkflowManager().GetDefinitionsById(id, messageId, sessionId)
    return definition
}

WorkflowManager.prototype.getDefinitionsById = function (id, successCallback, errorCallback, caller) {
    // server is the client side singleton...
    var params = []
    params[0] = new RpcData({ "name": "id", "type": 2, "dataBytes": utf8_to_b64(id) })

    // Call it on the server.
    server.executeJsFunction(
        GetDefinitionsById.toString(), // The function to execute remotely on server
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
function GetAllDefinitions() {
    var allDefintions = null
    allDefintions = server.GetWorkflowManager().GetAllDefinitions(messageId, sessionId)
    return allDefintions
}

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
                GetAllDefinitions.toString(), // The function to execute remotely on server
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
function GetDefinitionInstances(definitionsId) {
    var instances = null
    instances = server.GetWorkflowManager().GetDefinitionInstances(definitionsId, messageId, sessionId)
    return instances
}

WorkflowManager.prototype.getDefinitionInstances = function (definitions, successCallback, errorCallback, caller) {
    // server is the client side singleton...
    var params = []
    params[0] = new RpcData({ "name": "definitionsId", "type": 2, "dataBytes": utf8_to_b64(definitions.UUID) })

    // Call it on the server.
    server.executeJsFunction(
        GetDefinitionInstances.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {

            var definitions = caller.definitions
            var results = result[0]
            for (var i = 0; i < results.length; i++) {
                setReferences(results[i], definitions.getReferences())
                // Here I will link the definition itself...
                results[i].getDefinitions = function (definitions) {
                    return function () {
                        return definitions
                    }
                } (definitions)
            }

            caller.successCallback(results, caller.caller)
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
function NewItemAwareElementInstance(bpmnElementId, data) {
    // Now create the new process...
    var itemDefinitionEntity = null
    itemDefinitionEntity = server.GetWorkflowProcessor().NewItemAwareElementInstance(bpmnElementId, data, messageId, sessionId)
    return itemDefinitionEntity
}

WorkflowManager.prototype.newItemAwareElementInstance = function (bpmnElementId, data, successCallback, errorCallback, caller) {

    // server is the client side singleton...
    var params = []
    params[0] = new RpcData({ "name": "bpmnElementId", "type": 2, "dataBytes": utf8_to_b64(bpmnElementId) })
    params[1] = new RpcData({ "name": "data", "type": 2, "dataBytes": utf8_to_b64(data) })
    
    // Call it on the server.
    server.executeJsFunction(
        NewItemAwareElementInstance.toString(), // The function to execute remotely on server
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

/*
 * Start a new process instance.
 */
function StartProcess(processUUID, eventData, eventDefinitionData) {
    // Now create the new process...
    server.GetWorkflowManager().StartProcess(processUUID, eventData, eventDefinitionData, messageId, sessionId)

}

WorkflowManager.prototype.startProcess = function (processUUID, eventData, eventDefinitionData, successCallback, errorCallback, caller) {

    // server is the client side singleton...
    var params = []
    params[0] = new RpcData({ "name": "processId", "type": 2, "dataBytes": utf8_to_b64(processUUID) })
    params[1] = new RpcData({ "name": "data", "type": 4, "dataBytes": utf8_to_b64(JSON.stringify(eventData)) })
    params[2] = new RpcData({ "name": "data", "type": 4, "dataBytes": utf8_to_b64(JSON.stringify(eventDefinitionData)) })

    // Call it on the server.
    server.executeJsFunction(
        StartProcess.toString(), // The function to execute remotely on server
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


