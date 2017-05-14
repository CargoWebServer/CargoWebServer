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
 * @fileOverview Entities related functionalities.
 * @author Dave Courtois
 * @version 1.0
 */


/**
 * The entity manager gives acces to ojects stored on the server.
 * @constructor
 * @extends EventHub
 */
var EntityManager = function () {

    if (server == undefined) {
        return
    }

    EventHub.call(this, EntityEvent)

    /**
     * @property {object} entityPrototypes Keeps track of prototypes in use.
     */
    this.entityPrototypes = {}

    /**
     * Keep the entity localy to reduce the network traffic and
     * prevent infinite recurion.
     */
    this.entities = {}

    // The list of item to init...
    this.toInit = {}

    return this
}

EntityManager.prototype = new EventHub(null);
EntityManager.prototype.constructor = EntityManager;

/*
 * Dispatch event.
 */
EntityManager.prototype.onEvent = function (evt) {
    // Set the internal object.
    if (evt.code == UpdateEntityEvent || evt.code == NewEntityEvent) {
        if (this.entityPrototypes[evt.dataMap["entity"].TYPENAME] == undefined) {
            console.log("Type " + evt.dataMap["entity"].TYPENAME + " not define!")
            return
        }
        if (this.entities[evt.dataMap["entity"].UUID] == undefined) {
            var entity = eval("new " + evt.dataMap["entity"].TYPENAME + "()")
            entity.initCallback = function (self, evt, entity) {
                return function (entity) {
                    server.entityManager.setEntity(entity)
                    EventHub.prototype.onEvent.call(self, evt)
                }
            } (this, evt, entity)
            entity.init(evt.dataMap["entity"])

        } else {
            // update the object values.
            // but before I call the event I will be sure the entity have 
            var entity = this.entities[evt.dataMap["entity"].UUID]
            entity.initCallback = function (self, evt, entity) {
                return function (entity) {
                    // Test if the object has change here befor calling it.
                    server.entityManager.setEntity(entity)
                    if (evt.done == undefined) {
                        EventHub.prototype.onEvent.call(self, evt)
                    }
                    evt.done = true // Cut the cyclic recursion.
                }
            } (this, evt, entity)
            // if (hasChange(entity, evt.dataMap["entity"])) {
            setObjectValues(entity, evt.dataMap["entity"])
            // }
        }
    } else if (evt.code == DeleteEntityEvent) {
        var entity = this.entities[evt.dataMap["entity"].UUID]
        if (entity != undefined) {
            this.resetEntity(entity)
            EventHub.prototype.onEvent.call(this, evt)
        }
    }
}

/*
 * Set an entity.
 */
EntityManager.prototype.setEntity = function (entity) {

    this.getEntityPrototype(entity.TYPENAME, entity.TYPENAME.split(".")[0],
        function (prototype, caller) {
            var id_ = entity.TYPENAME + ":"
            for (var i = 0; i < prototype.Ids.length; i++) {
                var id = prototype.Ids[i]
                if (id == "UUID") {
                    server.entityManager.entities[entity.UUID] = entity
                } else {
                    if (entity[id].length > 0) {
                        id_ += entity[id]
                        if (i < prototype.Ids.length - 1) {
                            id_ += "_"
                        }
                    }
                }
            }

            // Set the entity with it id.
            if (entity.IsInit) {
                server.entityManager.entities[id_] = entity
                if (entity.TYPENAME.startsWith("BPMN20")) {
                    server.workflowManager.bpmnElements[id_] = entity
                }
            }
        },
        function (errMsg, caller) {
            /** Nothing to do here. */
        },
        {})

}

/*
 * Remove an entity.
 */
EntityManager.prototype.resetEntity = function (entity) {
    var prototype = this.entityPrototypes[entity.TYPENAME]
    delete server.entityManager.entities[entity.UUID]

    var id = entity.TYPENAME + ":"
    for (var i = 0; i < prototype.Ids.length; i++) {
        id += entity[prototype.Ids[i]]
        if (i < prototype.Ids.length - 1) {
            id += "_"
        }
    }
    if (server.entityManager.entities[id] != undefined) {
        delete server.entityManager.entities[id]
    }
}

/*
 * Server side script
 */
function GetObjectsByType(typeName, queryStr, storeId) {
    var objects
    objects = server.GetEntityManager().GetObjectsByType(typeName, queryStr, storeId, messageId, sessionId)
    return objects
}

/**
 * That function is use to retreive objects with a given type.
 * @param {string} typeName The name of the type we looking for in the form packageName.typeName
 * @param {string} storeId The name of the store where the information is saved.
 * @param {string} queryStr It contain the code of a function to be executed by the server to filter specific values.
 * @param {function} progressCallback The function is call when chunk of response is received.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
EntityManager.prototype.getObjectsByType = function (typeName, storeId, queryStr, progressCallback, successCallback, errorCallback, caller) {

    // First of all i will get the entity prototype.
    server.entityManager.getEntityPrototype(typeName, storeId,
        // The success callback.
        function (result, caller) {
            // Set the parameters.
            var typeName = caller.typeName
            var storeId = caller.storeId
            var queryStr = caller.queryStr
            var successCallback = caller.successCallback
            var progressCallback = caller.progressCallback
            var errorCallback = caller.errorCallback
            var caller = caller.caller

            // Create the list of parameters.
            var params = []
            params.push(createRpcData(typeName, "STRING", "typeName"))
            params.push(createRpcData(queryStr, "STRING", "queryStr"))
            params.push(createRpcData(storeId, "STRING", "storeId"))

            // Call it on the server.
            server.executeJsFunction(
                GetObjectsByType, // The function to execute remotely on server
                params, // The parameters to pass to that function
                function (index, total, caller) { // The progress callback
                    // Keep track of the file transfert.
                    caller.progressCallback(index, total, caller.caller)
                },
                function (result, caller) {
                    var entities = []
                    if (result[0] != undefined) {
                        for (var i = 0; i < result[0].length; i++) {
                            var entity = eval("new " + caller.prototype.TypeName + "(caller.prototype)")
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
                function (errMsg, caller) {
                    // call the immediate error callback.
                    caller.errorCallback(errMsg, caller.caller)
                    // dispatch the message.
                    server.errorManager.onError(errMsg)
                }, // Error callback
                { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback, "prototype": result } // The caller
            )
        },
        // The error callback.
        function (errMsg, caller) {
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller)
            // dispatch the message.
            server.errorManager.onError(errMsg)
        }, { "typeName": typeName, "storeId": storeId, "queryStr": queryStr, "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback })
}

/*
 * Sever side script.
 */
function GetEntityLnks(uuid) {
    var lnkLst = null
    lnkLst = server.GetEntityManager().GetEntityLnks(uuid, messageId, sessionId)
    return lnkLst
}

/**
 * Retreive the list of all entity link's (dependencie) at once...
 * @param {string} uuid The of the entity that we want to retreive link's
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
EntityManager.prototype.getEntityLnks = function (uuid, progressCallback, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(createRpcData(uuid, "STRING", "uuid"))

    // Call it on the server.
    server.executeJsFunction(
        GetEntityLnks.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            caller.progressCallback(index, total, caller)
        },
        function (results, caller) {
            caller.successCallback(results[0], caller.caller)
        },
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback } // The caller
    )
}


/*
 * Server side script
 */
function GetEntityByUuid(uuid) {
    var entity = null
    entity = server.GetEntityManager().GetObjectByUuid(uuid, messageId, sessionId)
    return entity
}

/**
 * That function is use to retreive objects with a given type.
 * @param {string} uuid The uuid of the entity we looking for. The uuid must has form typeName%UUID.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
EntityManager.prototype.getEntityByUuid = function (uuid, successCallback, errorCallback, caller) {

    var entity = server.entityManager.entities[uuid]
    if (entity != undefined) {
        if (entity.TYPENAME == entity.__class__ && entity.IsInit == true) {
            successCallback(entity, caller)
            return // break it here.
        }
    }

    var typeName = uuid.substring(0, uuid.indexOf("%"))
    var storeId = typeName.substring(0, typeName.indexOf("."))

    // Create the entity prototype here.
    var entity = eval("new " + typeName + "(caller.prototype)")
    entity.UUID = uuid
    entity.TYPENAME = typeName
    server.entityManager.setEntity(entity)

    // First of all i will get the entity prototype.
    server.entityManager.getEntityPrototype(typeName, storeId,
        // The success callback.
        function (result, caller) {
            // Set the parameters.
            var uuid = caller.uuid
            var successCallback = caller.successCallback
            var progressCallback = caller.progressCallback
            var errorCallback = caller.errorCallback
            var caller = caller.caller

            var params = []
            params.push(createRpcData(uuid, "STRING", "uuid"))

            // Call it on the server.
            server.executeJsFunction(
                GetEntityByUuid.toString(), // The function to execute remotely on server
                params, // The parameters to pass to that function
                function (index, total, caller) { // The progress callback
                    // Nothing special to do here.
                },
                function (result, caller) {
                    var entity = server.entityManager.entities[result[0].UUID]
                    entity.initCallback = function (caller) {
                        return function (entity) {
                            server.entityManager.setEntity(entity)
                            caller.successCallback(entity, caller.caller)
                        }
                    } (caller)
                    if (entity.IsInit == false) {
                        entity.init(result[0])
                    } else {
                        caller.successCallback(entity, caller.caller)
                    }

                },
                function (errMsg, caller) {
                    server.errorManager.onError(errMsg)
                    caller.errorCallback(errMsg, caller.caller)
                }, // Error callback
                { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback, "prototype": result } // The caller
            )
        },
        // The error callback.
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller)
        }, { "uuid": uuid, "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback })
}


/*
 * Server side script
 */
function GetEntityById(storeId, typeName, ids) {
    var entity = null
    entity = server.GetEntityManager().GetObjectById(storeId, typeName, ids, messageId, sessionId)
    return entity
}

/**
 * Retrieve an entity with a given typename and id.
 * @param {string} typeName The object type name.
 * @param {string} ids The id's (not uuid) of the object to look for.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 * @param {parent} the parent object reference.
 */
EntityManager.prototype.getEntityById = function (storeId, typeName, ids, successCallback, errorCallback, caller, parent) {

    if (!isArray(ids)) {
        console.log("ids must be an array! ", ids)
    }

    // key in the server.
    var id = typeName + ":"
    for (var i = 0; i < ids.length; i++) {
        id += ids[i]
        if (i < ids.length - 1) {
            id += "_"
        }
    }

    if (server.entityManager.entities[id] != undefined) {
        successCallback(server.entityManager.entities[id], caller)
        return // break it here.
    }

    // First of all i will get the entity prototype.
    server.entityManager.getEntityPrototype(typeName, storeId,
        // The success callback.
        function (result, caller) {

            // Set the parameters.
            var storeId = caller.storeId
            var typeName = caller.typeName
            var ids = caller.ids
            var successCallback = caller.successCallback
            var progressCallback = caller.progressCallback
            var errorCallback = caller.errorCallback
            var caller = caller.caller

            var params = []
            params.push(createRpcData(storeId, "STRING", "storeId"))
            params.push(createRpcData(typeName, "STRING", "typeName"))
            params.push(createRpcData(ids, "JSON_STR", "ids")) // serialyse as an JSON object array...

            // Call it on the server.
            server.executeJsFunction(
                GetEntityById.toString(), // The function to execute remotely on server
                params, // The parameters to pass to that function
                function (index, total, caller) { // The progress callback
                    // Nothing special to do here.
                },
                function (result, caller) {
                    if (result[0] == null) {
                        return
                    }

                    // In case of existing entity.
                    if (server.entityManager.entities[result[0].UUID] != undefined && result[0].TYPENAME == result[0].__class__) {
                        caller.successCallback(server.entityManager.entities[result[0].UUID], caller.caller)
                        return // break it here.
                    }

                    var entity = eval("new " + caller.prototype.TypeName + "(caller.prototype)")
                    entity.initCallback = function () {
                        return function (entity) {
                            caller.successCallback(entity, caller.caller)
                        }
                    } (caller)

                    entity.init(result[0])
                },
                function (errMsg, caller) {
                    server.errorManager.onError(errMsg)
                    caller.errorCallback(errMsg, caller.caller)
                }, // Error callback
                { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback, "prototype": result, "parent": parent, "ids": ids } // The caller
            )
        },
        // The error callback.
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller)
        }, { "storeId": storeId, "typeName": typeName, "ids": ids, "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback })
}

/*
 * Server side code.
 */
function CreateEntity(parentUuid, attributeName, typeName, id, values) {
    var entity = null
    entity = server.GetEntityManager().CreateEntity(parentUuid, attributeName, typeName, id, values, messageId, sessionId)
    return entity
}

/**
 * That function is use to create a new entity of a given type..
 * @param {string} parentUuid The uuid of the parent entity if there is one, null otherwise.
 * @param {string} attributeName The attribute name is the name of the new entity in his parent. (parent.attributeName = this)
 * @param {string} typeName The type name of the new entity.
 * @param {string} id The id of the new entity. There is no restriction on the value entered.
 * @param {object} entity the entity to be save, it can be nil.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
EntityManager.prototype.createEntity = function (parentUuid, attributeName, typeName, id, entity, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(createRpcData(parentUuid, "STRING", "parentUuid"))
    params.push(createRpcData(attributeName, "STRING", "attributeName"))
    params.push(createRpcData(typeName, "STRING", "typeName"))
    params.push(createRpcData(id, "STRING", "id"))
    params.push(createRpcData(entity, "JSON_STR", "entity"))


    // Call it on the server.
    server.executeJsFunction(
        CreateEntity.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            var entity = eval("new " + result[0].TYPENAME + "()")
            entity.initCallback = function () {
                return function (entity) {
                    if (caller.successCallback != undefined) {
                        caller.successCallback(entity, caller.caller)
                        caller.successCallback = undefined
                    }
                }
            } (caller)
            entity.init(result[0])
        },
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * Sever side code.
 */
function RemoveEntity(uuid) {
    entity = server.GetEntityManager().RemoveEntity(uuid, messageId, sessionId)
}

/*
 * That function is use to remove an entity with a given uuid.
 * @param {string} uuid The uuid of entity to delete. Must have the form TypeName%UUID
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
EntityManager.prototype.removeEntity = function (uuid, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(createRpcData(uuid, "STRING", "uuid"))

    // Call it on the server.
    server.executeJsFunction(
        RemoveEntity.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            caller.successCallback(true, caller.caller)
        },
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * Server side script.
 */
function SaveEntity(entity, typeName) {
    // save the entity value.
    entity = server.GetEntityManager().SaveEntity(entity, typeName, messageId, sessionId)
    return entity
}

/**
 * Save The entity. If the entity does not exist it creates it.
 * @param {Entity} entity The entity to save.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
EntityManager.prototype.saveEntity = function (entity, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    entity.NeedSave = true
    var params = []
    params.push(createRpcData(entity, "JSON_STR", "entity"))
    params.push(createRpcData(entity.TYPENAME, "STRING", "typeName"))

    var functionStr = SaveEntity.toString()

    // Call it on the server.
    server.executeJsFunction(
        functionStr, // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            var entity = eval("new " + result[0].TYPENAME + "()")
            entity.initCallback = function () {
                return function (entity) {
                    // Set the new entity values...
                    server.entityManager.setEntity(entity)
                    if (caller.successCallback != undefined) {
                        caller.successCallback(entity, caller.caller)
                    }
                }
            } (caller)
            entity.init(result[0])
        },
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            if (caller.errorCallback != undefined) {
                caller.errorCallback(errMsg, caller.caller)
            }
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * Sever side script.
 */
function CreateEntityPrototype(storeId, prototype) {
    var proto = null
    proto = server.GetEntityManager().CreateEntityPrototype(storeId, prototype, messageId, sessionId)
    return proto
}

/**
 * Create a new entity prototype.
 * @param {string} storeId The store id, where to create the new prototype.
 * @param {EntityPrototype} prototype The prototype object to create.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
EntityManager.prototype.createEntityPrototype = function (storeId, prototype, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(createRpcData(storeId, "STRING", "storeId"))
    params.push(createRpcData(prototype, "JSON_STR", "prototype"))

    // Call it on the server.
    server.executeJsFunction(
        CreateEntityPrototype.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (results, caller) {
            var proto = new EntityPrototype()
            proto.init(results[0])
            server.entityManager.entityPrototypes[results[0].TypeName] = proto
            caller.successCallback(proto, caller.caller)
        },
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * Server side code.
 */
function GetEntityPrototype(typeName, storeId) {
    var proto = null
    proto = server.GetEntityManager().GetEntityPrototype(typeName, storeId, messageId, sessionId)
    return proto
}

/**
 * That function will retreive the entity prototype with a given type name.
 * @param {string} typeName The type name of the prototype to retreive.
 * @param {string} storeId The store id, where to create the new prototype.
 * @param {EntityPrototype} The prototype object to create.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
EntityManager.prototype.getEntityPrototype = function (typeName, storeId, successCallback, errorCallback, caller) {

    // Retrun entity prototype that aleady exist.
    if (server.entityManager.entityPrototypes[typeName] != undefined) {
        successCallback(server.entityManager.entityPrototypes[typeName], caller)
        return
    }

    // server is the client side singleton.
    var params = []
    params.push(createRpcData(typeName, "STRING", "typeName"))
    params.push(createRpcData(storeId, "STRING", "storeId"))

    // Call it on the server.
    server.executeJsFunction(
        GetEntityPrototype.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (results, caller) {
            var proto = new EntityPrototype()
            server.entityManager.entityPrototypes[results[0].TypeName] = proto
            proto.init(results[0])
            caller.successCallback(proto, caller.caller)
        },
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * Server side code.
 */
function GetEntityPrototypes(storeId) {
    var protos = null
    protos = server.GetEntityManager().GetEntityPrototypes(storeId, messageId, sessionId)
    return protos
}

/**
 * That function will retreive all prototypes of a store.
 * @param {string} storeId The store id, where to create the new prototype.
 * @param {EntityPrototype} The prototype object to create.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
EntityManager.prototype.getEntityPrototypes = function (storeId, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(createRpcData(storeId, "STRING", "storeId"))

    // Call it on the server.
    server.executeJsFunction(
        GetEntityPrototypes.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (results, caller) {
            var results = results[0]
            var protoypes = []
            if (results != null) {
                for (var i = 0; i < results.length; i++) {
                    var proto = new EntityPrototype()
                    server.entityManager.entityPrototypes[results[i].TypeName] = proto
                    proto.init(results[i])
                    protoypes.push(proto)
                }
            }

            caller.successCallback(protoypes, caller.caller)
        },
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/*
 * Server side code.
 */
function GetDerivedEntityPrototypes(typeName, storeId) {
    var proto = null
    proto = server.GetEntityManager().GetDerivedEntityPrototypes(typeName, messageId, sessionId)
    return proto
}

/**
 * That function will retreive the derived entity prototype from a given type.
 * @param {string} typeName The type name of the parent entity.
 * @param {string} storeId The store id, where to create the new prototype.
 * @param {EntityPrototype} The prototype object to create.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
EntityManager.prototype.getDerivedEntityPrototypes = function (typeName, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(createRpcData(typeName, "STRING", "typeName"))

    // Call it on the server.
    server.executeJsFunction(
        GetDerivedEntityPrototypes.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (results, caller) {
            var prototypes = []
            if (results[0] != null) {
                for (var i = 0; i < results[0].length; i++) {
                    var result = results[0][i]
                    if (server.entityManager.entityPrototypes[results[0][i].TypeName] != undefined) {
                        prototypes.push(server.entityManager.entityPrototypes[results[0][i].TypeName])
                    } else {
                        var proto = new EntityPrototype()
                        proto.init(results[0][i])
                        server.entityManager.entityPrototypes[results[0][i].TypeName] = proto
                    }
                }
            }
            // return the list of prototype object.
            caller.successCallback(prototypes, caller.caller)
        },
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

/**
  Restriction are expression defining limitation on the
  range of value that a variable can take. Type of restriction
  are :
    <ul>
        <li>Enumeration</li>
        <li>FractionDigits</li>
        <li>Length</li>
        <li>MaxExclusive</li>
        <li> MaxInclusive</li>
        <li>MaxLength</li>
        <li>MinExclusive</li>
        <li> MinInclusive</li>
        <li>MinLength</li>
        <li>Pattern</li>
        <li>TotalDigits</li>
        <li>WhiteSpace</li>
    </ul>
    */
var Restriction = function () {
    // The the of the restriction (Facet)
    this.Type

    // The value.
    this.Value
}

/**
 * The entity prototype define the schema of an entity. Entity prototype
 * are usefull to make entity persistent, to export and import schema from
 * other data source like xml schema. It's also possible to create at runtime
 * new kind of entity and make query over it.
 * </br>note: Fields, FieldsType, FieldsDocumentation, FieldsNillable and FieldsOrder 
 * must have the same number of elements.
 * @constructor 
 */
var EntityPrototype = function () {
    this.TYPENAME = "Server.EntityPrototype"

    /**
     * @property {string} TypeName The type name of the entity must be write 'packageName.className'
     */
    this.TypeName = "" // string

    /** 
     * @property {string} Documentation The entity type documentation.
     */
    this.Documentation

    /** 
     * @property {boolean} IsAbstract True if the entity prototype is an abstrac class.
     */
    this.IsAbstract = false

    /**
     * @property ListOf The prototype consiste of a list of other type
     */
    this.ListOf = ""

    /**
     * @property SubstitutionGroup The list of substitution group, must be use in abstract class.
     */
    this.SubstitutionGroup = []

    /**
     * @property SuperTypeNames The list of super class name.
     */
    this.SuperTypeNames = [] // string

    /**
     *  @property Restrictions The field restrictions.
     */
    this.Restrictions = [] // Restriction

    /**
     * @property Ids The list of fields used as ids.
     */
    this.Ids = [] // string

    /**
     * @property Indexs The list of fields used as index.
     */
    this.Indexs = [] // string

    ////////////////////////////////////////////////////////////////////////////
    //  Field's properties.
    ////////////////////////////////////////////////////////////////////////////

    /** 
     * @property Fields The fields name 
     */
    this.Fields = []

    /**
     * @property FieldsDocumentation The field documentation
     */
    this.FieldsDocumentation = []

    /**
     * @property FieldsType The fields type
     */
    this.FieldsType = [] // string

    /**
     * @property FieldsVisibility The fields visibility
     */
    this.FieldsVisibility = [] // bool

    /**
     * @property FieldsNillable If the field can be nil
     */
    this.FieldsNillable = [] // bool

    /**
     * @property FieldsOrder The field order
     */
    this.FieldsOrder = [] // int

    return this
}

/**
 * Return the list of field to be used as title.
 */
EntityPrototype.prototype.getTitles = function () {
    // The get title default function... can be overload.
    return [this.Ids[1]] // The first index only...
}

/**
 * Create a new class form json object.
 * @param {object} object The object that regroup the prototype properties.
 */
EntityPrototype.prototype.init = function (object) {
    if (object == null || object.TypeName == undefined) {
        return
    }

    // The type name.
    this.TypeName = object.TypeName

    // The package will be an object on the global scope.
    this.PackageName = object.TypeName.split(".")[0]
    if (window[this.PackageName] == undefined) {
        window[this.PackageName] = eval(this.PackageName + " = {}")
    }

    // The type name.
    this.ClassName = object.TypeName.substring(object.TypeName.indexOf(".") + 1)

    // The object ids
    this.appendIds(object.Ids)

    // The object indexs
    this.appendIndexs(object.Indexs)

    // Now the fields.
    if (object.Fields != null && object.FieldsType != null && object.FieldsVisibility != null && object.FieldsOrder != null) {

        // Append parent uuid if none is define.
        if (!contains(object.Fields, "ParentUuid")) {
            object.Fields.unshift("ParentUuid")
            object.FieldsType.unshift("xs.string")
            object.FieldsVisibility.unshift(false)
            object.FieldsOrder.push(object.FieldsOrder.length)
        }

        // Append the uuid if none is define.
        if (!contains(object.Fields, "UUID")) {
            // append the uuid...
            object.Fields.unshift("UUID")
            object.FieldsType.unshift("xs.string")
            object.FieldsVisibility.unshift(false)
            object.FieldsOrder.push(object.FieldsOrder.length)
            object.Ids.unshift("UUID")
        }

        for (var i = 0; i < object.Fields.length; i++) {
            this.appendField(object.Fields[i], object.FieldsType[i], object.FieldsVisibility[i], object.FieldsOrder[i])
            if (object.Fields[i] == "UUID") {
                if (!contains(this.Ids, "UUID")) {
                    this.Ids.unshift("UUID")
                }
            } else if (object.Fields[i] == "ParentUuid") {
                if (!contains(this.Indexs, "ParentUuid")) {
                    this.Indexs.unshift("ParentUuid")
                }
            }
        }

    } else {
        console.log(object.TypeName + " has no fields!!!")
    }

    // Now the restriction
    this.Restrictions = object.Restrictions

    // If the object is abstract
    this.IsAbstract = object.IsAbstract

    // The list of substitution type.
    this.SubstitutionGroup = object.SubstitutionGroup

    // The list of Supertype 
    this.SuperTypeNames = object.SuperTypeNames

    // if the type is a collection.
    this.ListOf = object.ListOf

    // other standard fields.
    this.appendField("childsUuid", "[]xs.string", false, this.Fields.length)
    this.appendField("referenced", "[]Server.EntityRef", false, this.Fields.length)

    // Generate the class code.
    this.generateConstructor()
}

/**
 * Append a new object value into an entity.
 */
function appendObjectValue(object, field, value) {
    var prototype = server.entityManager.entityPrototypes[object.TYPENAME]
    var fieldIndex = prototype.getFieldIndex(field)
    var prototype = server.entityManager.entityPrototypes[object.TYPENAME]
    var fieldType = prototype.FieldsType[fieldIndex]
    var isArray = fieldType.startsWith("[]")
    var isRef = fieldType.endsWith(":Ref")
    if (fieldIndex > -1) {
        if (isArray) {
            var index = 0
            var isExist = false
            // Create an array if is not exist.
            if (object[field] == undefined) {
                object[field] = []
            } else {
                for (var i = 0; i < object[field].length; i++) {
                    index = i
                    if (object[field][i].UUID == value.UUID) {
                        isExist = true
                        break
                    }
                }
            }

            // Set the reference in case of reference
            if (isRef) {
                setRef(object, field, value.UUID, true)
            } else {
                // Append or replace the value in case of an array.
                if (!isExist) {
                    index = object[field].length
                    object[field].push(value)
                } else {
                    object[field][index] = value
                }
            }

        } else {
            object[field] = value
            if (isRef) {
                setRef(object, field, value.UUID, false)
            }
        }
    }

    // Set need save to the object.
    object["NeedSave"] = true

    // Can be usefull to intercept change event...
    if (object.onChange != undefined) {
        object.onChange(object)
    }
}

/**
 * Remove an object from a given object.
 */
function removeObjectValue(object, field, value) {
    var prototype = server.entityManager.entityPrototypes[object.TYPENAME]
    var index = prototype.getFieldIndex(field)
    if (index > -1) {
        var fieldType = prototype.FieldsType[index]
        var isArray = fieldType.startsWith("[]")
        var isRef = fieldType.endsWith(":Ref")

        if (isArray) {
            // Here the entity is an array.
            for (var i = 0; i < object[field].length; i++) {
                var uuid = ""
                if (isObject(object[field][i])) {
                    uuid = object[field][i].UUID
                } else {
                    uuid = object[field][i]
                }

                if (uuid == value.UUID) {
                    object[field].splice(i, 1)
                    // remove the set and reset function.
                    if (isRef) {
                        delete object["reset_" + field + "_" + uuid + "_ref"]
                        delete object["set_" + field + "_" + uuid + "_ref"]
                    }
                }
            }
        } else {
            if (isRef) {
                var uuid = ""
                if (isObject(object[field])) {
                    uuid = object[field].UUID
                } else {
                    uuid = object[field]
                }
                delete object["reset_" + field + "_" + uuid + "_ref"]
                delete object["set_" + field + "_" + uuid + "_ref"]
            }
            delete object[field]
        }
    }
    object["NeedSave"] = true
}

/**
 * That function is use to reset the entity of it's original value.
 */
function resetObjectValues(object) {
    // Remove the object panel...
    delete object["panel"]
    var prototype = server.entityManager.entityPrototypes[object.TYPENAME]

    for (var propertyId in object) {
        var propretyType = prototype.FieldsType[prototype.getFieldIndex(propertyId)]
        if (propretyType != undefined && object[propertyId] != null) {
            var isRef = propretyType.endsWith(":Ref")
            var isArray = propretyType.startsWith("[]")
            var isBaseType = propretyType.startsWith("[]xs.") || propretyType.startsWith("xs.")
            if (isArray) {
                for (var i = 0; i < object[propertyId].length; i++) {
                    if (isObject(object[propertyId][i])) {
                        if (object[propertyId][i]["UUID"] != undefined) {
                            // Reset it's sub-objects.
                            if (!isRef && !isBaseType) {
                                resetObjectValues(object[propertyId][i])
                            }
                        }
                    }
                }
            } else if (isObject(object[propertyId])) {
                if (object[propertyId]["UUID"] != undefined) {
                    if (!isRef && !isBaseType) {
                        resetObjectValues(object[propertyId])
                    }
                }
            }
        }

        /** Only reference must be reset here. */
        if (propertyId.startsWith("reset_") && propertyId.endsWith("_ref")) {
            // Call the reset function.
            object[propertyId]()
        }
    }

}

/**
 * Return a property field type for a given field for a given type name.
 */
function getPropertyType(typeName, property) {
    var prototype = server.entityManager.entityPrototypes[typeName]
    var propertyType = null
    for (var i = 0; i < prototype.Fields.length; i++) {
        if (prototype.Fields[i] == property) {
            propertyType = prototype.FieldsType[i]
            break
        }
    }
    return propertyType
}

/////////////////////////////////////////////////////////////////////////
// Entity referenced information.

/**
 * That stucture is use to keep track of object 
 * referenced by another object.
 */
var EntityRef = function (name, owner, value) {
    this.TypeName = "Server.EntityRef"
    this.Name = name
    this.OwnerUuid = owner.UUID
    this.Value = value

    return this
}

/**
 * Append a reference to a target if it dosen't already exist.
 */
function appendReferenced(refName, target, owner) {

    if (target.referenced == undefined) {
        target.referenced = []
    }

    // Look if the value is not already there...
    for (var i = 0; i < target.referenced.length; i++) {
        var ref = target.referenced[i]
        if (ref.Name == refName && ref.OwnerUuid == owner.UUID) {
            // Nothing to do here.
            return
        }
    }

    // Append the referenced.
    target.referenced.push(new EntityRef(refName, owner, ""))
}

/////////////////////////////////////////////////////////////////////////
// Initialisation code here.

/**
 * Set object reference.
 */
function setRef(owner, property, refValue, isArray) {
    if (refValue.length == 0) {
        return owner
    }

    // Keep track of the references targets.
    if (owner.references.indexOf(refValue) == -1) {
        owner.references.push(refValue)
    }

    if (!isObjectReference(refValue)) {
        owner[property] = refValue
        return owner
    }


    if (isArray) {
        var index = owner[property].length
        if (owner[property].indexOf(refValue) == -1) {
            owner[property].push(refValue)
            /* The reset reference fucntion **/
            owner["reset_" + property + "_" + refValue + "_ref"] = function (entity, propertyName) {
                return function () {
                    for (var i = 0; i < entity[propertyName].length; i++) {
                        if (isObject(entity[propertyName][i])) {
                            entity[propertyName][i] = entity[propertyName][i].UUID
                        }
                    }
                }
            } (owner, property)

            /* The set reference fucntion **/
            owner["set_" + property + "_" + refValue + "_ref"] = function (entityUuid, propertyName, index, refValue) {
                return function (initCallback) {
                    var isExist = server.entityManager.entities[refValue] != undefined
                    var isInit = false

                    if (isExist) {
                        isInit = server.entityManager.entities[refValue].IsInit
                    }

                    if (isExist && isInit) {
                        // Here the reference exist on the server.
                        var entity = server.entityManager.entities[entityUuid]
                        var ref = server.entityManager.entities[refValue]
                        entity[propertyName][index] = ref
                        appendReferenced(propertyName, ref, entity)
                        if (initCallback != undefined) {
                            initCallback(ref)
                            initCallback = undefined
                        }

                    } else {
                        server.entityManager.getEntityByUuid(refValue,
                            function (result, caller) {
                                var propertyName = caller.propertyName
                                var index = caller.index
                                var entity = server.entityManager.entities[caller.entityUuid]
                                var ref = server.entityManager.entities[caller.refValue]
                                entity[propertyName][index] = ref
                                appendReferenced(propertyName, ref, entity)
                                if (caller.initCallback != undefined) {
                                    caller.initCallback(ref)
                                    caller.initCallback = undefined
                                }
                            },
                            function (errorMsg, caller) {
                            },
                            { "entityUuid": entityUuid, "propertyName": propertyName, "index": index, "refValue": refValue, "initCallback": initCallback }
                        )
                    }
                }
            } (owner.UUID, property, index, refValue)
        }
    } else {
        owner[property] = refValue

        /* The reset fucntion **/
        owner["reset_" + property + "_" + refValue + "_ref"] = function (entityUuid, propertyName) {
            return function () {
                var entity = server.entityManager.entities[entityUuid]
                // Set back the id of the reference
                if (isObject(entity[propertyName])) {
                    entity[propertyName] = entity[propertyName].UUID
                }
            }
        } (owner.UUID, property)

        /* The set fucntion **/
        owner["set_" + property + "_" + refValue + "_ref"] = function (entityUuid, propertyName, refValue) {
            return function (initCallback) {

                var isExist = server.entityManager.entities[refValue] != undefined
                var isInit = false

                if (isExist) {
                    isInit = server.entityManager.entities[refValue].IsInit
                }

                // If the entity is already on the client side...
                if (isExist && isInit) {
                    // Here the reference exist on the server.
                    var entity = server.entityManager.entities[entityUuid]
                    var ref = server.entityManager.entities[refValue]
                    entity[propertyName] = ref
                    appendReferenced(propertyName, ref, entity)
                    if (initCallback != undefined) {
                        initCallback(ref)
                        initCallback = undefined
                    }
                } else {
                    server.entityManager.getEntityByUuid(refValue,
                        function (result, caller) {
                            var propertyName = caller.propertyName
                            var entity = server.entityManager.entities[caller.entityUuid]
                            var ref = server.entityManager.entities[caller.refValue]
                            entity[propertyName] = ref
                            appendReferenced(propertyName, ref, entity)
                            if (caller.initCallback != undefined) {
                                caller.initCallback(ref)
                                caller.initCallback = undefined
                            }
                        },
                        function () { },
                        { "entityUuid": entityUuid, "propertyName": propertyName, "refValue": refValue, "initCallback": initCallback }
                    )
                }
            }
        } (owner.UUID, property, refValue)
    }
    return owner
}

function hasChange(entity, object) {
    return true
    /*
    // cut reference here.
    resetObjectValues(entity)

    // Now I will look if the object value has change.
    var prototype = server.entityManager.entityPrototypes[object["TYPENAME"]]

    for (var property in object) {
        var propertyType = getPropertyType(object["TYPENAME"], property)
        if (propertyType != null) {
            var isRef = propertyType.endsWith(":Ref")
            if (propertyType.startsWith("[]")) {
                // The property is an array.
                if (object[property] != null) {
                    if (object[property].length > 0) {
                        for (var i = 0; i < object[property].length; i++) {
                            if (isString(object[property][i])) {
                                if (isRef) {

                                }
                            }
                        }
                    }
                }
            } else {
                if (object[property] != undefined) {
                    if (isString(object[property])) {
                        if (isRef) {

                        }
                    }
                }
            }
        }
    }

    // set reference back.
    setObjectValues(entity)*/
}

/**
 * Set object, that function call setObjectValues in this path so it's recursive.
 */
function setSubObject(parent, property, values, isArray) {

    if (values.TYPENAME == undefined || values.UUID.length == 0) {
        return parent
    }

    server.entityManager.getEntityPrototype(values.TYPENAME, values.TYPENAME.split(".")[0],
        function (result, caller) {
            var parent = caller.parent
            var property = caller.property
            var values = caller.values
            var isArray = caller.isArray
            if (values.TYPENAME == "BPMN20.StartEvent") {
                i = 0;
            }

            var object = server.entityManager.entities[values.UUID]
            if (object == undefined) {
                object = eval("new " + values.TYPENAME + "()")
                // Keep track of the parent uuid in the child.
                object.UUID = values.UUID
                server.entityManager.setEntity(object)
            }

            // Keep track of the child uuid inside the parent.
            if (parent.childsUuid == undefined) {
                parent.childsUuid = []
            }

            if (parent.childsUuid.indexOf(object.UUID)) {
                parent.childsUuid.push(object.UUID)
            }

            if (isArray) {
                if (parent[property] == undefined) {
                    parent[property] = []
                }
                object.init(values)
                parent[property].push(object)

            } else {
                object.init(values)
                parent[property] = object
            }

            object.ParentUuid = parent.UUID

            server.entityManager.setEntity(object)

        },
        function () {

        }, { "parent": parent, "property": property, "values": values, "isArray": isArray })


    return parent
}

/**
 * That function initialyse an object created from a given prototype constructor with the values from a plain JSON object.
 * @param {object} object The object to initialyse.
 * @param {object} values The plain JSON object that contain values.
 */
function setObjectValues(object, values) {

    // Get the entity prototype.
    var prototype = server.entityManager.entityPrototypes[object["TYPENAME"]]
    if (prototype == undefined) {
        return
    }

    server.entityManager.setEntity(object)

    ////////////////////////////////////////////////////////////////////////
    // Set back the reference...
    if (values == undefined) {
        // call set_property on each object if there is a function defined.
        for (var property in object) {
            var propertyType = getPropertyType(object["TYPENAME"], property)
            if (propertyType != null) {
                var isRef = propertyType.endsWith(":Ref")
                if (propertyType.startsWith("[]")) {
                    // The property is an array.
                    if (object[property] != null) {
                        if (object[property].length > 0) {
                            for (var i = 0; i < object[property].length; i++) {
                                if (isString(object[property][i])) {
                                    if (isRef) {
                                        if (object["set_" + property + "_" + object[property] + "_ref"] != undefined) {
                                            // Call it.
                                            object["set_" + property + "_" + object[property] + "_ref"]()
                                        }
                                    }
                                }
                            }
                        }
                    }
                } else {
                    if (object[property] != undefined) {
                        if (isString(object[property])) {
                            if (isRef) {
                                if (object["set_" + property + "_" + object[property] + "_ref"] != undefined) {
                                    // Call it.
                                    object["set_" + property + "_" + object[property] + "_ref"]()
                                }
                            }
                        }
                    }
                }
            }
        }
        return
    }

    ////////////////////////////////////////////////////////////////////
    // Reset actual object fields and cound number of sub-objects...

    // The list of properties to set.
    // reference values must be put at end of the list.
    var properties = []
    for (var i = 0; i < prototype.FieldsType.length; i++) {
        var fieldType = prototype.FieldsType[i]
        var field = prototype.Fields[i]
        if (field.startsWith("M_") || field.startsWith("[]M_")) {
            // Reset the objet fields.
            if (!fieldType.startsWith("sqltypes.") && !fieldType.startsWith("[]sqltypes.") && !fieldType.startsWith("xs.") && !fieldType.startsWith("[]xs.")) {
                if (fieldType.startsWith("[]")) {
                    object[field] = []
                } else {
                    object[field] = ""
                }
            } else {
                // Reset the base type fields.
                if (fieldType.startsWith("[]xs.") || fieldType.startsWith("[]sqltypes.")) {
                    object[field] = []
                } else if (fieldType.startsWith("xs.") || fieldType.startsWith("sqltypes.")) {
                    object[field] = ""
                }
            }
        }
    }

    ////////////////////////////////////////////////////////////////////////////////
    // Generate sub-object and reference set and reset function
    for (var property in values) {

        var propertyType = prototype.FieldsType[prototype.getFieldIndex(property)]

        if (propertyType != null) {
            // Condition...
            var isRef = propertyType.endsWith(":Ref")

            // M_listOf, M_valueOf field or enumeration type contain plain value.
            var isBaseType = propertyType.startsWith("sqltypes.") && !propertyType.startsWith("[]sqltypes.") || propertyType.startsWith("[]xs.") || propertyType.startsWith("xs.") || property == "M_listOf" || property == "M_valueOf" || propertyType.startsWith("enum:")

            if (values[property] != null) {
                if (isBaseType) {
                    // String, int, double...
                    if (propertyType == "xs.base64Binary") {
                        // In case of binairy object I will try to create object if information is given for it.
                        // TODO see why to decode is necessary here...
                        var strVal = decode64(decode64(values[property]))
                        if (strVal.indexOf("TYPENAME") != -1 && strVal.indexOf("__class__") != -1) {
                            var jsonObj = JSON.parse(strVal)
                            if (!isArray(jsonObj)) {
                                // In case of the object is not an array...
                                var obj = eval("new " + jsonObj.TYPENAME + "()")
                                obj.initCallback = function (uuid, property) {
                                    return function (val) {
                                        server.entityManager.entities[uuid][property] = val
                                    }
                                } (object.UUID, property)

                                obj.init(jsonObj)
                            } else {
                                for (var i = 0; i < jsonObj.length; i++) {
                                    var jsonObj_ = JSON.parse(jsonObj[i])
                                    var obj = eval("new " + jsonObj_.TYPENAME + "()")
                                    obj.initCallback = function (uuid, property) {
                                        return function (val) {
                                            if (server.entityManager.entities[uuid][property] == "") {
                                                server.entityManager.entities[uuid][property] = []
                                            }
                                            server.entityManager.entities[uuid][property].push(val)
                                        }
                                    } (object.UUID, property)
                                    obj.init(jsonObj_)
                                }
                            }
                        } else {
                            // No information available to create an object, so the string will be use....
                            object[property] = strVal
                        }
                    } else {
                        object[property] = values[property]
                    }
                } else {
                    // Set object ref or values... only property begenin with M_ will be set here...
                    var isArray_ = propertyType.startsWith("[]")
                    if (property.startsWith("[]M_") || property.startsWith("M_")) {
                        if (isArray_) {
                            object[property] = []
                            for (var i = 0; i < values[property].length; i++) {
                                if (isRef) {
                                    object = setRef(object, property, values[property][i], isArray_)
                                } else {
                                    object = setSubObject(object, property, values[property][i], isArray_)
                                }
                            }
                        } else {
                            if (isRef) {
                                object = setRef(object, property, values[property], isArray_)
                            } else {
                                object = setSubObject(object, property, values[property], isArray_)
                            }
                        }
                    }
                }
            }
        }
    }

    //////////////////////////////////////////////////////
    // Set common values...
    object.UUID = values.UUID
    object.NeedSave = false
    object.exist = true
    object.IsInit = true // The object part only and not the refs...
    object.ParentUuid = values.ParentUuid // set the parent uuid.

    // Set the initialyse object.
    server.entityManager.setEntity(object)
    
    // Call the init callback.
    if (object.initCallback != undefined) {
        object.initCallback(object)
        object.initCallback == undefined
    }
}

/**
 * This function generate the js class base on the entity prototype.
 */
EntityPrototype.prototype.generateConstructor = function () {
    var constructorSrc = this.PackageName + " = function(){\n"
    var constructorSrc = this.PackageName + " || {};\n"

    var packageName = this.PackageName
    var classNames = this.ClassName.split(".")

    for (var i = 0; i < classNames.length - 1; i++) {
        packageName += "." + classNames[i]
        constructorSrc += packageName + " = " + packageName + " || {};\n"
    }

    // I will create the object constructor from the information
    // of the fields.
    constructorSrc += this.PackageName + "." + this.ClassName + " = function(){\n"

    // Common properties share by all entity.
    constructorSrc += " this.__class__ = \"" + this.PackageName + "." + this.ClassName + "\"\n"
    constructorSrc += " this.UUID = this.UUID\n"
    constructorSrc += " this.TYPENAME = \"" + this.TypeName + "\"\n"
    constructorSrc += " this.ParentUuid = \"\"\n"
    constructorSrc += " this.childsUuid = []\n"
    constructorSrc += " this.references = []\n"
    constructorSrc += " this.NeedSave = true\n"
    constructorSrc += " this.IsInit = false\n"
    constructorSrc += " this.exist = false\n"
    constructorSrc += " this.initCallback = undefined\n"
    constructorSrc += " this.panel = null\n"

    // Remove space accent '' from the field name
    function normalizeFieldName(fieldName) {
        // TODO make distinctive..
        fieldName = fieldName.replaceAll(" ", "_")
        fieldName = fieldName.replaceAll("'", "")
        return fieldName
    }

    // Fields.
    for (var i = 0; i < this.Fields.length; i++) {

        constructorSrc += " this." + normalizeFieldName(this.Fields[i])

        if (this.FieldsType[i].startsWith("[]")) {
            constructorSrc += " = undefined\n"
        } else {
            if (isXsString(this.FieldsType[i]) || isXsRef(this.FieldsType[i]) || isXsId(this.FieldsType[i])) {
                constructorSrc += " = \"\"\n"
            } else if (isXsInt(this.FieldsType[i])) {
                constructorSrc += " = 0\n"
            } else if (isXsNumeric(this.FieldsType[i])) {
                constructorSrc += " = 0.0\n"
            } else if (isXsDate(this.FieldsType[i])) {
                constructorSrc += " = new Date()\n"
            } else if (isXsBoolean(this.FieldsType[i])) {
                constructorSrc += " = false\n"
            } else if (this.FieldsType[i].startsWith("enum:")) {
                constructorSrc += " = 1\n"
            } else {
                // Object here.
                constructorSrc += " = undefined\n"
            }
        }
    }

    // Now the stringify function.
    constructorSrc += " this.stringify = function(){\n"
    constructorSrc += "       resetObjectValues(this)\n"
    constructorSrc += "       var cache = [];\n"
    constructorSrc += "       var entityStr = JSON.stringify(this, function(key, value) {\n"
    constructorSrc += "           if (typeof value === 'object' && value !== null) {\n"
    constructorSrc += "               if (cache.indexOf(value) !== -1) {\n"
    constructorSrc += "                   // Circular reference found, discard key\n"
    constructorSrc += "                   return;\n"
    constructorSrc += "               }\n"
    constructorSrc += "               // Store value in our collection\n"
    constructorSrc += "               cache.push(value);\n"
    constructorSrc += "           }\n"
    constructorSrc += "           return value;\n"
    constructorSrc += "       });\n"
    constructorSrc += "       cache = null; // Enable garbage collection\n"
    constructorSrc += "       setObjectValues(this)\n"
    constructorSrc += "       return entityStr\n"
    constructorSrc += "   }\n"

    // The get parent function
    constructorSrc += " this.getParent = function(){\n"
    constructorSrc += "       return server.entityManager.entities[this.ParentUuid]\n"
    constructorSrc += "  }\n"

    // The setter function.
    for (var i = 0; i < this.Fields.length; i++) {
        if (!this.FieldsType[i].startsWith("xs.") && !this.FieldsType[i].startsWith("[]xs.")) {
            // So its not a basic type.
            constructorSrc += " this.set" + normalizeFieldName(this.Fields[i]).replace("M_", "").capitalizeFirstLetter() + " = function(value){\n"
            constructorSrc += "     appendObjectValue(this,\"" + normalizeFieldName(this.Fields[i]) + "\", value)\n"
            constructorSrc += " }\n"
        }
    }

    // The remove function.
    for (var i = 0; i < this.Fields.length; i++) {
        if (!this.FieldsType[i].startsWith("xs.") && !this.FieldsType[i].startsWith("[]xs.")) {
            // So its not a basic type.
            constructorSrc += " this.remove" + normalizeFieldName(this.Fields[i]).replace("M_", "").capitalizeFirstLetter() + " = function(value){\n"
            constructorSrc += "     removeObjectValue(this,\"" + normalizeFieldName(this.Fields[i]) + "\", value)\n"
            constructorSrc += " }\n"
        }
    }

    // The get title default function... can be overload.
    for (var i = 1; i < this.Ids.length; i++) {
        var fieldIndex = this.getFieldIndex(this.Ids[i])
        var field = this.Fields[fieldIndex]
        if (this.FieldsVisibility[fieldIndex] == true) {
            constructorSrc += " this.getTitles = function(){\n"
            constructorSrc += "     return [this." + field + "]\n"
            constructorSrc += " }\n"
            break
        }
    }

    // Keep the reference on the entity prototype.
    // The class level.
    constructorSrc += " return this\n"
    constructorSrc += "}\n\n"

    constructorSrc += this.PackageName + "." + this.ClassName + ".prototype.init = function(object){\n"
    // First of all i will set reference in the result.
    constructorSrc += "   this.TYPENAME = object.TYPENAME\n"
    constructorSrc += "   this.UUID = object.UUID\n"
    constructorSrc += "   this.IsInit = false\n"
    constructorSrc += "   setObjectValues(this, object)\n"
    constructorSrc += "}\n\n"

    // Set the function.
    //console.log(constructorSrc)
    eval(constructorSrc)

}

/**
 * Create a new prototypeField.
 * @param {string} name The field name.
 * @param {string} name The field type name.
 * @param {boolean} isVisible True, if the field is visible.
 * @param {int} order The order the field will be return, usefull to display.
 */
EntityPrototype.prototype.appendField = function (name, typeName, isVisible, order) {
    // Set the field name.
    if (!contains(this.Fields, name)) {
        this.Fields.push(name)

        // Set the field type name
        this.FieldsType.push(typeName)

        // Set the field visibility
        this.FieldsVisibility.push(isVisible)
        // And the order (index in the list of fields.)
        this.FieldsOrder.push(parseInt(order))
    }
}

/**
 * Append the list of indexs
 * @param indexs The list of indexs to append.
 */
EntityPrototype.prototype.appendIndexs = function (indexs) {
    if (indexs != null) {
        for (var i = 0; i < indexs.length; i++) {
            this.Indexs.push(indexs[i])
        }
    }
}

/**
 * Append the list of id's
 * @param ids The list of ids to append.
 */
EntityPrototype.prototype.appendIds = function (ids) {
    if (ids != null) {
        for (var i = 0; i < ids.length; i++) {
            this.Ids.push(ids[i])
        }
    }
}

/**
 * Retreive the index of a given field in the prototype.
 * @param field The field we looking for.
 * @returns The index of the field or -1 if the field is not there.
 */
EntityPrototype.prototype.getFieldIndex = function (field) {
    for (var i = 0; i < this.Fields.length; i++) {
        if (this.Fields[i] == field) {
            return i
        }
    }
    return -1
}

/**
 * Return the name of the base type if the type is an extension of such a type.
 * @param {string} typeName The extension type name, ex. xs.string, xs.int, xs.date etc.
 */
EntityManager.prototype.getBaseTypeExtension = function (typeName, isArray) {
    if (!isArray) {
        isArray = typeName.startsWith("[]")
    }
    typeName = typeName.replace("[]", "").replace(":Ref", "")
    var prototype = this.entityPrototypes[typeName]

    if (prototype != null) {
        if (prototype.SuperTypeNames != null) {
            for (var i = 0; i < prototype.SuperTypeNames.length; i++) {
                if (prototype.SuperTypeNames[i].startsWith("xs.")) {
                    if (isArray) {
                        return "[]" + prototype.SuperTypeNames[i]
                    }
                    return prototype.SuperTypeNames[i]
                } else if (this.getBaseTypeExtension(prototype.SuperTypeNames[i]).length > 0) {
                    return this.getBaseTypeExtension(prototype.SuperTypeNames[i])
                }
            }
        } else if (prototype.ListOf != null) {
            if (prototype.ListOf.length > 0) {
                if (prototype.ListOf.startsWith("xs.")) {
                    return "[]" + prototype.ListOf
                }
                return this.getBaseTypeExtension(prototype.ListOf, true)
            }
        }
    }
    return ""
}

/**
 * Look if the given type is a list of other type.
 * @param {string} typeName The extension type name, ex. xs.string, xs.int, xs.date etc.
 */
EntityManager.prototype.isListOf = function (typeName) {
    typeName = typeName.replace("[]", "").replace(":Ref", "")
    var prototype = this.entityPrototypes[typeName]

    if (prototype != null) {
        if (prototype.ListOf != null) {
            if (prototype.ListOf.length > 0) {
                return true
            }
        }
        if (prototype.SuperTypeNames != null) {
            for (var i = 0; i < prototype.SuperTypeNames.length; i++) {
                if (prototype.SuperTypeNames[i] != prototype.TypeName) {
                    if (this.isListOf(prototype.SuperTypeNames[i])) {
                        return true
                    }
                }
            }
        }
    }
    return false
}