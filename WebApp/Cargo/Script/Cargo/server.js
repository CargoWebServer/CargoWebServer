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


	return this
 }
 
 /*
 * Initialisation of the socket handler.
 */
 function initConnection(adress, onOpenCallback, onCloseCallback, onMessageCallback, sessionId) {
	console.log("try to open connection with service: " + adress)
 }