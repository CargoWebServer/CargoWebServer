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
 Server.prototype.init = function (onOpenConnectionCallback, onCloseConnectionCallback) {
	var address = this.ipv4 + ":" + this.port.toString()
	// The connection will be set on the sever side.
	initConnection(address, onOpenConnectionCallback, onCloseConnectionCallback, sessionId, this)
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
 