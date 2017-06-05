
var HelloTest = function () {
    return this
}

HelloTest.prototype.SayHello = function (to, messageId, sessionId) {
    // Test register action.
    server.GetServiceManager().RegisterAction("HelloTest.SayHello", [], [])
    console.log("hello " + to + "!")
}
