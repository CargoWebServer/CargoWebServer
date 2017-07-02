var AccountManager = function () {
    if (server == undefined) {
        return
    }
    EventHub.call(this, AccountEvent)

    return this
}

AccountManager.prototype = new EventHub(null);
AccountManager.prototype.constructor = AccountManager;

AccountManager.prototype.onEvent = function (evt) {
    EventHub.prototype.onEvent.call(this, evt)
}
AccountManager.prototype.getAccountById = function (id, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    server.executeJsFunction(
        "AccountManagerGetAccountById",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("CargoEntities.Account", "CargoEntities",
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

AccountManager.prototype.getUserById = function (id, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(id, "STRING", "id"))

    server.executeJsFunction(
        "AccountManagerGetUserById",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("CargoEntities.User", "CargoEntities",
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

AccountManager.prototype.me = function (connectionId, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(connectionId, "STRING", "connectionId"))

    server.executeJsFunction(
        "AccountManagerMe",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("CargoEntities.Account", "CargoEntities",
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

AccountManager.prototype.register = function (name, password, email, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(name, "STRING", "name"))
    params.push(createRpcData(password, "STRING", "password"))
    params.push(createRpcData(email, "STRING", "email"))

    server.executeJsFunction(
        "AccountManagerRegister",
        params,
        undefined, //progress callback
        function (results, caller) { // Success callback
            server.entityManager.getEntityPrototype("CargoEntities.Account", "CargoEntities",
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

