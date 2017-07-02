var SessionManager = function () {
    if (server == undefined) {
        return
    }
    EventHub.call(this, SessionEvent)

    return this
}

SessionManager.prototype = new EventHub(null);
SessionManager.prototype.constructor = SessionManager;

SessionManager.prototype.onEvent = function (evt) {
    EventHub.prototype.onEvent.call(this, evt)
}
SessionManager.prototype.getActiveSessionByAccountId = function (accountId, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(accountId, "STRING", "accountId"))

    server.executeJsFunction(
        "SessionManagerGetActiveSessionByAccountId",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("Server.Sessions", "Server",
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

SessionManager.prototype.getActiveSessions = function (successCallback, errorCallback, caller) {
    var params = []

    server.executeJsFunction(
        "SessionManagerGetActiveSessions",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("CargoEntities.Session", "CargoEntities",
                function (prototype, caller) { // Success Callback
                    var entities = []
                    for (var i = 0; i < caller.results[0].length; i++) {
                        var entity = eval("new " + prototype.TypeName + "()")
                        if (i == caller.results[0].length - 1) {
                            entity.initCallback = function (caller) {
                                return function (entity) {
                                    server.entityManager.setEntity(entity)
                                    caller.successCallback(entities, caller.caller)
                                }
                            }(caller)
                        } else {
                            entity.initCallback = function (entity) {
                                server.entityManager.setEntity(entity)
                            }
                        }
                        entities.push(entity)
                        entity.init(caller.results[0][i])
                    }
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

SessionManager.prototype.login = function (accountName, psswd, serverId, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(accountName, "STRING", "accountName"))
    params.push(createRpcData(psswd, "STRING", "psswd"))
    params.push(createRpcData(serverId, "STRING", "serverId"))

    server.executeJsFunction(
        "SessionManagerLogin",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("CargoEntities.Session", "CargoEntities",
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

SessionManager.prototype.logout = function (toCloseId, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(toCloseId, "STRING", "toCloseId"))

    server.executeJsFunction(
        "SessionManagerLogout",
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

SessionManager.prototype.updateSessionState = function (state, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(state, "INTEGER", "state"))

    server.executeJsFunction(
        "SessionManagerUpdateSessionState",
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