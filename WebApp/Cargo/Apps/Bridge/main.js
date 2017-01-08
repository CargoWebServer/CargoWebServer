/**
 * Created by Dave on 2/14/2015 .
 */

var applicationName = document.getElementsByTagName("title")[0].text

// Set the address here
var mainPage = null
var catalog = null

/**
 * This function is the entry point of the application...
 */
function main() {
    server.eventHandler.appendEventFilter(
        "CargoEntities.",
        "EntityEvent",
        function () {
            server.eventHandler.appendEventFilter(
                "CargoEntities.",
                "FileEvent",
                function () {
                    server.eventHandler.appendEventFilter(
                        "CargoEntities.",
                        "SessionEvent",
                        function () {
                            server.eventHandler.appendEventFilter(
                                "CargoEntities.",
                                "TableEvent",
                                function () {
                                    server.eventHandler.appendEventFilter(
                                        "CargoEntities.",
                                        "AccountEvent",
                                        function () {
                                            server.eventHandler.appendEventFilter(
                                                "CargoEntities.",
                                                "SecurityEvent",
                                                function () {
                                                    var bodyElement = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "height: 100%; width: 100%;" });

                                                    // The body element....
                                                    var mainLayout = new Element(bodyElement, { "tag": "div", "style": "position: absolute; top:0px; left:0px; right:0px; bottom: 0px;" });

                                                    // The page to display when the user is logged in
                                                    homePage = new HomePage(mainLayout)

                                                    // The login page...
                                                    var loginPage = new LoginPage(function (mainLayout, homePage) {
                                                        return function (sessionsInfo) {
                                                            homePage.init(mainLayout, sessionsInfo)
                                                        }
                                                    } (mainLayout, homePage))

                                                    // the main page...
                                                    mainPage = new MainPage(mainLayout, loginPage)
                                                },
                                                function () { },
                                                undefined
                                            )
                                        },
                                        function () { },
                                        undefined
                                    )
                                },
                                function () { },
                                undefined
                            )
                        },
                        function () { },
                        undefined
                    )
                },
                function () { },
                undefined
            )
        },
        function () { },
        undefined
    )

}
