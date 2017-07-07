var applicationName = document.getElementsByTagName("title")[0].text

var languageInfo = {
    "en": {
        "label": "What is your name?",
        "btn": "greeting"
    },
    "fr": {
        "label": "Quel est votre nom?",
        "btn": "salutation"
    }
}

// Depending of the language the correct text will be set.
server.languageManager.appendLanguageInfo(languageInfo)
var service = null

/**
 * Server side function
 */
function SayHello(to) {

    // Test service usage here...
    var hostName = "localhost"
    var ipv4 = "127.0.0.1"
    var port = 9494

    service = new Server(hostName, ipv4, port)

    var address = "ws://" + service.ipv4 + ":" + service.port.toString()
    //var address =  service.ipv4 + ":" + service.port.toString()
    service.conn = initConnection(address,
        // on open connection callback
        function (caller) {
            // display the connection id...
            caller.conn = this
            caller.ping(
                // success callback
                function (result, caller) {
                    // Here I received the pong message.
                    console.log("----------> ln 41 Ping response: ", result)
  
                    // Now I will call a JS function...
                    function TestSayHelloPlugin(to) {
                        var msg = SayHello.sayHelloTo(to)
                        return msg
                    }

                    var params = []
                    params.push(createRpcData("Cargo !!!!", "STRING", "str"))

                    // Call it on the server.
                    caller.executeJsFunction(
                        TestSayHelloPlugin.toString(), // The function to execute remotely on server
                        params, // The parameters to pass to that function
                        function (index, total, caller) { // The progress callback
                            // Nothing special to do here.
                        },
                        function (results, caller) {
                            // Here I received the result.
                            console.log("----------> ln 61 result received: ", results[0])
                        },
                        function (errMsg, caller) {

                        }, // Error callback
                        {} // The caller
                    )

                },
                // error callback
                function (errObj, caller) {
                }, caller)
        },
        // on close callback
        function () {

        },
        // on message callback
        function () {

        },
        // the local session id.
        sessionId, service)

    return "Hello " + to + "!!!"
}

/**
 * This function is the entry point of the application...
 */
function main() {
    // Creation of the element.
    var panel = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "id": "panel" })

    // Now I will append the element label, input and button.
    panel.appendElement({ "tag": "div", "style": "display: table-row" }).down()
        .appendElement({ "tag": "span", "id": "label" })
        .appendElement({ "tag": "input", "id": "input" }).up()
        .appendElement({ "tag": "div", "style": "display: table-row" }).down()
        .appendElement({ "tag": "button", "id": "btn" }).up()
        .appendElement({ "tag": "div", "style": "display: table-row" }).down()
        .appendElement({ "tag": "span", "id": "greeting" })

    // Now the action.
    var btn = panel.getChildById("btn")
    var greeting = panel.getChildById("greeting")
    var input = panel.getChildById("input")

    btn.element.onclick = function (greeting, input) {
        return function () {
            var params = []
            params.push(createRpcData(input.element.value, "STRING", "id"))
            // Call it on the server.
            server.executeJsFunction(
                SayHello.toString(), // The function to execute remotely on server
                params, // The parameters to pass to that function
                // The progress callback
                function (index, total, caller) {
                    // Nothing to do here.
                },
                // The success callback
                function (results, caller) {
                    // Caller is the greething reference.
                    caller.element.innerHTML = results[0]
                },
                // The error callback
                function (errMsg, caller) {
                    // display the message in the console.
                    console.log(errMsg)
                }, // Error callback
                greeting // The caller
            )
        }
    } (greeting, input)
}
