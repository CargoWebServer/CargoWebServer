/**
 * Created by Dave Courtois on 2/14/2015.
 */
var applicationName = document.getElementsByTagName("title")[0].text

// Local event...
var ChangeFileEvent = 100
var ChangeThemeEvent = 101

// Set the address here
var mainPage = null
var catalog = null
var homePage = null

function init() {

    // Set style informations
    cargoThemeInfos = JSON.parse(localStorage.getItem("bridge_theme_infos"))
    if (cargoThemeInfos != undefined) {
        for (var ruleName in cargoThemeInfos) {
            var rule = getCSSRule(ruleName)
            for (property in cargoThemeInfos[ruleName]) {
                rule.style[property] = cargoThemeInfos[ruleName][property]
            }
        }
    }

    var bodyElement = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "height: 100%; width: 100%;" });

    // The body element....
    var mainLayout = new Element(bodyElement, { "tag": "div", "style": "position: absolute; top:0px; left:0px; right:0px; bottom: 0px;" });

    // The page to display when the user is logged in
    homePage = new HomePage()

    // The login page...
    var loginPage = new LoginPage(function (mainLayout) {
        return function (sessionsInfo) {
            homePage.init(mainLayout, sessionsInfo)
        }
    }(mainLayout),
        "SafranLdap" // Put the ldap sever id here
    )
    // the main page...
    mainPage = new MainPage(mainLayout, loginPage)

    /*bodyElement.element.oncontextmenu = function(){
        return false;
    }*/
}

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
                                                    // now the prototypes...
                                                    server.entityManager.getEntityPrototypes("BPMS",
                                                        function () {
                                                            init()
                                                        },
                                                        // error callback
                                                        function () {
                                                            // without bpmn
                                                            init()
                                                        }, {})
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