/**
 * Test send message from the sever to the service container.
 */
function TestMessageContainer(count) {
	// New service object
	//service = new Server("localhost", "127.0.0.1", 9595)
	service = new Server("localhost", "ws:127.0.0.1", 9494)

	// Initialyse the server connection.
	service.init(
		// onOpenConnectionCallback
		function (service, count) {
			console.log("Try to Send " + count + " message to the service container")
			var sayHello = new com.mycelius.SayHelloInterface(service)
			var callback = function (sayHello) {
				return function (index, count) {
					console.log("---> send message "+ index)
					sayHello.sayHelloTo("message " + index,
						// Success Callback
						function (result, caller) {
							console.log("------> Receive message ", result)
							var index = caller.index + 1
							if (caller.index  < caller.count) {
								caller.callback(index, count)
							}
						},
						// Error Callback
						function (errObj, caller) {
						}, { "index": index, "count": count, "callback": callback })
				}
			} (sayHello)

			/* Start the recursive loop... */
			callback(0, count)
		},
		// onCloseConnectionCallback
		function () {
		}, count)

}
