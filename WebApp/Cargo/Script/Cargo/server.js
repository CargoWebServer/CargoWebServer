// The default session id.
var sessionId = ""

/**
 * Connection class
 */
var Connection = function(){
	this.id = "" // The connection on the server side.
	return this;
}

/**
 * Contain the informations use by executeJsFunction function parameters.
 */
var RpcData = function(values){
	this.name = values.name
	this.type = values.type
	this.dataBytes = values.dataBytes
	this.typeName = values.typeName
}

/**
 * The sever class is a representation of a distant server into a local object.
 */
 var Server = function (hostName, ipv4, port) {
    /**
     * @property {string} hostName The server side domain name.
     * @example www.cargoWebserver.com
     * @example localhost
     */
    this.hostName = hostName

    /**
     * @property {string} ipv4 the ip adress of the server.
     * @example 127.0.0.1
     */
    this.ipv4 = ipv4

    /**
     * @property {int} port the port number.
     * @example 8080
     */
    this.port = port

    /**
     * @property {string} sessionId Identify the client with the server.
     */
    this.sessionId = null


	/**
	 * @property {string} The socket connection.
	 */
	 this.conn = null
	 
	 
	return this
 }
 
 /*
  * Open the connection with the other peer.
  */
 Server.prototype.init = function (onOpenConnectionCallback, onCloseConnectionCallback, caller) {
	var address = this.ipv4 + ":" + this.port.toString()

	// The connection will be set on the sever side.
	initConnection(address, onOpenConnectionCallback, onCloseConnectionCallback, sessionId, this, caller)
 }
 
 /*
 * Ping a given sever.
 */
 Server.prototype.ping = function (successCallback, errorCallback, caller) {
	ping(successCallback, errorCallback, caller, this.conn.id)
 }
 
/*
 * Execute JavaScript function.
 */
 Server.prototype.executeJsFunction = function (functionSrc, functionParams, progressCallback, successCallback, errorCallback, caller) {
	executeJsFunction(functionSrc, functionParams, progressCallback, successCallback, errorCallback, caller, this.conn.id)
 }

/**
 * Run a visual basic script with a given name on the server side.
 */
Server.prototype.executeVbSrcript = function (scriptName, args, successCallback, errorCallback, caller) {
    var params = new Array();
    params.push(createRpcData(scriptName, "STRING", "scriptName"))
    params.push(createRpcData(args, "JSON_STR", "args"))

    // Register this listener to the server.
    var rqst = new Request(randomUUID(), this.conn, "ExecuteVbScript", params,
        // Progress callback
        function () { },
        // Success callback
        function (id, results, caller) {
            // Keep the session id...
            caller.successCallback(results, caller.caller)
        },
        // Error callback...
        function (errorMsg, caller) {
            caller.errorCallback(errorMsg, caller.caller)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller });
    rqst.send();
}

/**
 * Run and executable command on the server and get the results.
 */
Server.prototype.runCmd = function (name, args, successCallback, errorCallback, caller) {
    var params = new Array();
    params.push(createRpcData(name, "STRING", "name"))
    params.push(createRpcData(args, "JSON_STR", "args"))

    // Register this listener to the server.
    var rqst = new Request(randomUUID(), this.conn, "RunCmd", params,
        // Progress callback
        function () { },
        // Success callback
        function (id, results, caller) {
            // Keep the session id...
            caller.successCallback(results, caller.caller)
        },
        // Error callback...
        function (errorMsg, caller) {
            caller.errorCallback(errorMsg, caller.caller)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller });
    rqst.send();
}

/**
 * Get the list of services and their respective source code. The code
 * permit to get access to service remote actions.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
Server.prototype.getServicesClientCode = function (successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = new Array();

    // Register this listener to the server.
    var rqst = new Request(randomUUID(), this.conn, "GetServicesClientCode", params,
        // Progress callback
        function () { },
        // Success callback
        function (id, results, caller) {
            // Keep the session id...
            caller.successCallback(results["result"], caller.caller)
        },
        // Error callback...
        function (errorMsg, caller) {
            caller.errorCallback(errorMsg, caller.caller)
        }, { "successCallback": successCallback, "errorCallback": errorCallback, "caller": caller });
    rqst.send();
}

/**
 * Close the server.
 */
Server.prototype.stop = function (successCallback, errorCallback, caller) {
    // server is the client side singleton...
    var params = []
    // Call it on the server.
    server.executeJsFunction(
        "Stop", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            //console.log(result)
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            caller.server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback, "server": this } // The caller
    )
}
