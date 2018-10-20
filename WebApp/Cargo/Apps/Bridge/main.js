/**
 * Created by Dave Courtois on 2/14/2015.
 */
 
var applicationName = document.getElementsByTagName("title")[0].text;

// Local event.
var ChangeFileEvent = 100;
var ChangeThemeEvent = 101;

// Set the address here.
var mainPage = null;
var catalog = null;
var homePage = null;
var bodyElement = null;

function init() {
    // console.log("welcome to bridge!")
    // Set style informations.
    cargoThemeInfos = JSON.parse(localStorage.getItem("bridge_theme_infos"));
    if (cargoThemeInfos !== undefined) {
        for (var ruleName in cargoThemeInfos) {
            var rule = getCSSRule(ruleName)
            if (rule !== undefined) {
                for (var property in cargoThemeInfos[ruleName]) {
                    rule.style[property] = cargoThemeInfos[ruleName][property]
                }
            } else {
                console.log("no css rule found with name ", ruleName, "!")
            }
        }
    }
    
    bodyElement = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "height: 100%; width: 100%;" });

    // The body element....
    var mainLayout = new Element(bodyElement, { "tag": "div", "style": "position: absolute; top:0px; left:0px; right:0px; bottom: 0px;" });

    // The page to display when the user is logged in.
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
}

/**
 * That is a connection with the service container.
 */
var service = new Server("localhost", "127.0.0.1", 9494)
//var service = new Server("mon-util-01", "10.2.128.70", 9494)
//var service = new Server("mon104", "10.67.44.73", 9494)
//var service = new Server("mon176", "10.67.46.210", 9494)

var xapian = null

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
                                                    server.entityManager.getEntityPrototypes("BPMN20",
                                                        function () {
                                                            service.conn = initConnection("ws://" + service.ipv4 + ":" + service.port.toString(),
                                                                function (service) {
                                                                    return function () {
                                                                        service.getServicesClientCode(
                                                                            // success callback
                                                                            function (results, caller) {
                                                                                // eval in that case contain the code to use the service.
                                                                                eval(results);
                                                                                // Xapian test...
                                                                                xapian = new com.mycelius.XapianInterface(caller.service);
                                                                                
                                                                                // Because require is already define with an empty function
                                                                                // i will get it from require.js and injected it here.
                                                                                require = undefined // reset the actual one (empty)
                                                                                server.fileManager.readTextFile("/lib/require.js", 
                                                                                    function(results, caller){
                                                                                        eval(results[0])
                                                                                        init();
                                                                                    },
                                                                                    function(err, caller){
                                                                                        
                                                                                    }, {})
                                                                               
                                                                            },
                                                                            // error callback.
                                                                            function () {
                                                                            }, { "service": service })
                                                                    }
                                                                }(service),
                                                                function () {
                                                                    console.log("Service is close!")
                                                                })

                                                        },
                                                        // error callback
                                                        function () {
                                                            // without bpmn
                                                            service.conn = initConnection("ws://" + service.ipv4 + ":" + service.port.toString(),
                                                                function (service) {
                                                                    return function () {
                                                                        console.log("Service is open!")
                                                                        service.getServicesClientCode(
                                                                            // success callback
                                                                            function (results, caller) {
                                                                                // eval in that case contain the code to use the service.
                                                                                eval(results)
                                                                                // Initialyse the search engine.
                                                                                xapian = new com.mycelius.XapianInterface(caller.service)
                                                                                init()
                                                                            },
                                                                            // error callback.
                                                                            function () {

                                                                            }, { "service": service })
                                                                    }
                                                                }(service),
                                                                function () {
                                                                    console.log("Service is close!")
                                                                })
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
