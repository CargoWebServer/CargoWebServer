var applicationName = document.getElementsByTagName("title")[0].text

var server = new Server("localhost", "127.0.0.1", 9393)

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

/**
 * Server side function
 */
function SayHello(to) {
    return "Hello " + to + "!"
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
        .appendElement({ "tag": "button", "id": "btn" })
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