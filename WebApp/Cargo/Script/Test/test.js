
var HelloTest = function () {
	// Input parameter's
	var to = new CargoEntities.Parameter()
	to.M_isArray = false
	to.M_name = "to"
	to.M_type = "string"
	
	// Output values
	var result = new CargoEntities.Parameter()
	result.M_isArray = false
	result.M_name = "result"
	result.M_type = "string"
	
	// Register action to be able to apply role and permission.
	server.GetServiceManager().RegisterAction("HelloTest.SayHello", [to], [result])

    return this
}

HelloTest.prototype.SayHello = function (to, messageId, sessionId) {
	// validate the message.
	var canExecute = server.GetSecurityManager().CanExecuteAction("HelloTest.SayHello", sessionId,  messageId)
	if(canExecute == false){
		return "Permission denied!"
	}
    // Test register action.
    return "hello " + to + "!";
}

function SayHello(to){
	var hello = new HelloTest()
	
	return hello.SayHello(to)
}
