// Global variable.
var loginPage = null;
var mainPage = null;
var bodyElement = new Element(document.body, { "tag": "div", "id": "body-element" });
// Here I will display the spinner.
var spinner = new Spinner(bodyElement, 30);
    
//var service = new Server("localhost", "127.0.0.1", 9494)
var service = new Server("mon176", "10.67.44.52", 9494)
var xapian = null

// The list of datastore to looking in.
var dbpaths = []
// The list of item types.
var itemTypes = []

var welcomeEvent = 2000
var createOrderEvent = 2001
var modifiedOrderEvent = 2002
var cancelOrderEvent = 2004
var catalogMessage = "catalogMessage"
var catalog = null
// Create the event listener for the current editor.
var catalogMessageHub = new EventHub(catalogMessage)


function main() {
    
    /*var url= "http://free.currencyconverterapi.com/api/v5/convert?q=EUR_USD&compact=y"
     $.getJSON(url,function(data){
        console.log(data)
      });*/

    // Initialisation of the interface.
    function main_() {
        // set the login page in the body to display it.
        service.conn = initConnection("ws://" + service.ipv4 + ":" + service.port.toString(),
            function (service) {
                return function () {
                    service.getServicesClientCode(
                        // success callback
                        function (results, caller) {
                            // eval in that case contain the code to use the service.
                            eval(results)
                            // Xapian test...
                            xapian = new com.mycelius.XapianInterface(caller.service)
                            server.eventHandler.addEventListener(
                            catalogMessageHub,
                                function () {
                                    
                                }
                            )
                            mainPage = new MainPage();
                            bodyElement.appendElement(mainPage.panel);

                            loginPage = new LoginPage(function (mainPage) {
                                return function (session) {
                                    mainPage.setActiveSession(session)
                                }
                            }(mainPage), "SafranLdap") // The name of the ldap server to authenticate with.

                        },
                        // error callback.
                        function () {
                        }, { "service": service })
                }
            }(service),
            function () {
                console.log("Service is close!")
            })
    }
    // the main page.
    main_()

    // Initialisation of the catalogue data here.
    server.entityManager.getEntityPrototypes("xs",
        // success callback
        function (result, caller) {
            server.entityManager.getEntityPrototypes("CatalogSchema", function (result, caller) {
                for (var i = 0; i < result.length; i++) {
                    var dbpath = server.root + "/Data/" + result[i].TypeName.split(".")[0] + "/" + result[i].TypeName + ".glass"
                    itemTypes.push(result[i].TypeName.split(".")[1])
                    dbpaths.push(dbpath)
                }
                // Connect all entity event related to the catalog schema.
                server.eventHandler.appendEventFilter(
                    "CatalogSchema.",
                    "EntityEvent",
                    function () {
                        // Here I will initialyse the catalog...
                        server.entityManager.getEntityById("CatalogSchema.CatalogType", "CatalogSchema", ["CatalogueMD"], true,
                            function (results, caller) {
                                // Comment it if the data is already imported.
                                // Import the filter informations here.
                                //var filePath = "C:\\Users\\mm006819\\Documents\\CargoWebServer\\WebApp\\Cargo\\Apps\\Catalogue\\csv\\filtres.csv"
                                //var filePath = "/home/dave/Documents/CargoWebServer/WebApp/Cargo/Apps/Catalog/csv/filtres.csv"
                                //if (results) {
                                    //importEntityFromCSV(filePath, getFiltersMapping(filePath), catalog, "M_items")
                                //}
                                // display the interface.
                                spinner.panel.element.style.display = "none";
                                catalog = results
                            },
                            function (err) {
                                var errObj = err.dataMap["errorObj"]
                                if (errObj.M_id == "ENTITY_ID_DOESNT_EXIST_ERROR") {
                                    // Here I will create the new entity...
                                    catalog = new CatalogSchema.CatalogType()
                                    catalog.M_schemaVersion = 1
                                    catalog.M_id = "CatalogueMD"
                                    // Save the empty caltalog...
                                    server.entityManager.saveEntity(catalog, function (result, caller) {
                                        catalog.init(result)
                                        spinner.panel.element.style.display = "none";
                                    })
                                }
                            },
                            null)
                    },
                    function () {},
                    undefined
                )

            }, null)
        },
        // error callback
        function () {

        },
        null)
}