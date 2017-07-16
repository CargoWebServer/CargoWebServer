/**
 * Test send message from the sever to the service container.
 */
function TestMessageContainer(count) {
	// New service object
	service = new Server("localhost", "127.0.0.1", 9595)

	// Initialyse the server connection.
	service.init(
		// onOpenConnectionCallback
		function (service, caller) {
			console.log("Try to Send " + caller + " message to the service container")
			var sayHello = new com.mycelius.SayHelloInterface(service)

			for (var i = 0; i < 10; i++) {
				// Call say hello. 
				sayHello.sayHelloTo("message " + i,
					// Success Callback
					function (result, caller) {
						console.log("------> success!", result)
					},
					// Error Callback
					function (errObj, caller) {

					}, i)
			}
		},
		// onCloseConnectionCallback
		function () {
		}, count)

}
