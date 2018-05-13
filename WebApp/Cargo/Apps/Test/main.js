var applicationName = document.getElementsByTagName("title")[0].text
// Fichier contenant les différents test.
var languageInfo = {
    "en": {
        "test":"test",
        "titi":"titi",
        "tutu":"tutu"
    }
}

/**
 * This function is the entry point of the application
 */
function main() {
    // eventTests()
    // Append filter to receive all session event message
    // on the sessionEvent channel.
    //securityTests()
    /*
    server.eventHandler.appendEventFilter(
         "CargoEntities.",
         "EntityEvent",
         function () {
             entitiesDump("CargoEntities.User")
         },
         function () { },
         undefined
     )
 
*/
    /* server.eventHandler.appendEventFilter(
          "CatalogSchema.",
          "EntityEvent",
          function () {
              entitiesDump("CatalogSchema.FiltreRectangulaire")
          },
          function () { },
          undefined
     )*/

    /*
    
        // utilityTests()
        //serverTests()
        //sessionTests()
        //languageManagerTests()
        //elementTests()
    
        //accountTests()
        //fileTests()
    
        //dataTests()
    
        //entityTests()
    
        /*server.sessionManager.login("admin", "adminadmin", "localhost",
            function () {
                // Create the dynamic entity here.
                testDynamicEntity()
            },
            function () {
                // Nothing to do here.
            }, {})*/


    // entityDump("item_1", "Test.Item")
    // entityDump("1", "dt_methode_development.dbo.DT_STATE")
    //entityDump("1", "dt_methode_development.dbo.DT_DTS")
    //entityDump("11362", "dt_methode_development.dbo.DT_DTS")
    // entityDump("1", "employees.employees")

    entitiesDump("COLLADASchema.COLLADA")

    // entitiesDump("DT3_informations.Department")
    // testServiceContainer()
    //entitiesDump("DT3_informations.Workpoint")
    //entitiesDump("CargoEntities.Action")

    //entitiesDump("XPDMXML.Product3DType")

    //testSayHello("Dave")
    // entityDump( "BPMN20", "BPMN20.Definitions")
    /*server.entityManager.getEntityPrototypes("COLLADASchema", 
    // success callback
    function(prototypes, caller){
        server.entityManager.getEntityByUuid("COLLADASchema.COLLADA%61933735-a65e-4f35-b8d6-920da06dee7e", false,
        // The success callback
        function(definitions, caller){
            console.log(definitions)
        },
        // The error callback
        function(error, caller){
            console.log(error)
        }, {})
    }, 
    // error callback.
    function(){

    })*/


    // entitiesDump("XPDMXML.ProductionGeneralSystem")
    // entitiesDump("CatalogSchema.CatalogType")

    //testEntityQuery()

    //TestWebRtc2()

    // Test get media source...
    // TestUploadFile()

    // Test get bmpn defintion instance...
    /*server.runCmd("cmd", ["/K","dir /Q C:\\Temp\\Erreur.txt"],
    // Success callback
    function(results, caller){
        console.log(results)
    },
    // Error callback.
    function(){

    });


    server.runCmd("cmd", ["/K","wmic datafile where name='c:\\\\windows\\\\system32\\\\notepad.exe' list full"],
    // Success callback
    function(results, caller){
        console.log(results)
    },
    // Error callback.
    function(){

    });*/

    /*
        function getFileInfos(path, callback) {
            server.runCmd("cmd", ["/K", "dir /Q " + path],
                // Success callback
                function (results, caller) {
                    var values = results["result"].split(/\s+/);

                    var author = ""
                    for(var i=0; i < values.length; i++){
                        if(values[i].startsWith("UD6")){
                            author = values[i]
                            break
                        }
                    }

                    var path = caller.path.replaceAll("\\", "\\\\")
                    server.runCmd("cmd", ["/K", "wmic datafile where name='" + path + "' list full"],
                        // Success callback
                        function (results, caller) {
                            var values = results["result"].split(/\s+/);
                            var fileInfos = {}
                            for (var i = 0; i < values.length; i++) {
                                if (values[i].indexOf("=") != -1) {
                                    var infos = values[i].split("=")
                                    var propertie = infos[0]
                                    var value = infos[1]
                                    if (propertie ==  "LastAccessed" || propertie == "LastModified" || propertie == "CreationDate" || propertie == "InstallDate") {
                                        
                                        value = moment(value, "YYYYMMDDHHmmSSSS").toDate();
                                    }
                                    fileInfos[propertie] = value
                                }
                            }
                            fileInfos["Author"] = author
                            caller.callback(fileInfos)
                        },
                        // Error callback.
                        function () {
    
                        }, { "callback": caller.callback, "author": author });
                },
                // Error callback.
                function () {
    
                }, { "path": path, "callback": callback });
        }
    
        getFileInfos("\\\\mon-util-01\\Demande_Travail_2\\1\\toto.txt", function (fileInfos) {
            console.log(fileInfos)
    
        })*/

    //LaunchImportNewSQLData_Process()

    /*
         server.entityManager.getEntityPrototypes("Test",
             // Success callback.
             function (result, caller) {
                 server.entityManager.getEntityPrototypes("BPMN20",
                     // Success callback.
                     function (result, caller) {
                         server.entityManager.getEntityPrototypes("BPMS",
                             // Success callback.
                             function (result, caller) {
                                 server.entityManager.getEntityById("BPMN20.Definitions", ["_1484846640138"],
                                     // success callbacak
                                     function (result, caller) {
                                         server.workflowManager.getDefinitionInstances(result,
                                             // success callback
                                             function (results, caller) {
                                                 var result = results[0]
                                                 var parent = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "position: absolute; width: auto; height: auto;" })
                                                 new EntityPanel(parent, result.TYPENAME, function (entity) {
                                                     return function (panel) {
                                                         panel.setEntity(entity)
                                                     }
                                                 }(result), undefined, false, result, "")
                                             },
                                             // error callback 
                                             function (errMsg, caller) {
                                             },
                                             {})
                                     },
                                     // Error callback 
                                     function (errMsg, caller) {
                                     }, {})
                             },
                             // Error callback.
                             function () {
                             }, {})
                     },
                     // Error callback.
                     function () {
     
                     }, {})
             },
             // Error callback.
             function () {
     
             }, {})
     */

    // Google OAuth
    /*server.oAuth2Manager.getResource("1010681964660.apps.googleusercontent.com", "profile", "", 
    function(result, caller){
    }, 
    function(errMsg, caller){
    }, {})*/

    // Facebook
    /*server.oAuth2Manager.getResource("821916804492503", "public_profile user_posts", "https://graph.facebook.com/v2.5/me/feed?limit=25", 
    function(results, caller){
        console.log("found results: ", results)
    },
    function(errMsg, caller){
    }, {})*/
    /*
         server.oAuth2Manager.getResource("1234", "openid profile email", "", 
         function(results, caller){
             console.log("found results: ", results)
         },
         function(errMsg, caller){
         }, {})
         
        
             var query = {}
             query.TypeName = "Proactive.AnalyseResult"
             query.Fields = ["M_NoTol", "M_NoFeat", "M_NoModele"]
             query.Query = ''
         
             server.dataManager.read("Proactive", JSON.stringify(query), [], [],
             function(){},
             function (results, caller) {
                 console.log("-------> results: ", results)
             }, function (errMsg, caller) {
        
             }, undefined)
             */
}

function testServiceContainer() {
    // Let us open a connection to a server... the service container.
    var service = new Server("localhost", "127.0.0.1", 9494)
    service.conn = initConnection("ws://" + service.ipv4 + ":" + service.port.toString(),
        function (service) {
            return function () {
                console.log("Service is open!")
                service.getServicesClientCode(
                    // success callback
                    function (results, caller) {
                        // eval in that case contain the code to use the service.
                        eval(results)
                        // Xapian test...
                        var xapian = new com.mycelius.XapianInterface(caller.service)

                        // Index csv file the file must exist on the server before that method is call.
                        /* Linux path */
                        //var datapath = "/home/dave/Documents/xapian/xapian-docsprint-master/data/100-objects-v1.csv"
                        //var dbpath = "/tmp/toto.glass";
                        var dbpath = "/home/dave/Documents/CargoWebServer/WebApp/Cargo/Data/CargoEntities/CargoEntities.glass"
                        /* Windows path */
                        // var datapath = "C:\\Users\\mm006819\\Documents\\xapian\\xapian-docsprint-master\\data\\100-objects-v1.csv"
                        //var dbpath = "C:\\Temp\\toto.glass";

                        /*xapian.indexCsv(
                            datapath,
                            dbpath,
                            ["Q:id_NUMBER","XD:DESCRIPTION","S:TITLE"],
                            "en",
                            // success callback
                            function (result, caller) {
                                console.log(result)
                            },
                            // error callback
                            function () {

                            }, {})*/

                        // Search for results...
                        xapian.search(
                            dbpath,
                            "Test",
                            ["XD:data"],
                            "en",
                            0,
                            10,
                            // success callback
                            function (result, caller) {
                                console.log(result)
                            },
                            // error callback
                            function () {

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
}

function testEntityQuery() {
    //{"TypeName":"CargoEntities.Log","Fields":["uuid"],"Indexs":["M_id=defaultErrorLogger"],"Query":""}
    var query = {}
    //query.TypeName = "Test.Item"
    //query.Fields = ["M_name", "M_description"]
    // Regex
    //query.Query = 'Test.Item.M_description == /Ceci est [a-z|\s|0-9]+/ && Test.Item.M_id != "item_5"'
    //query.Query = 'Test.Item.M_stringLst == /t[a-z]t[a-z](\\.)?/'
    //query.Query = 'Test.Item.M_description ^= "Ceci"'
    // bool value
    // query.Query = 'Test.Item.M_inStock == true'
    // int value
    //query.Query = 'Test.Item.M_qte <= 10'
    // float value
    //query.Query = 'Test.Item.M_price <= 3.0'
    // Date... using the 8601 string format.
    //query.Query = 'Test.Item.M_date >= "2016-07-12T15:42:22.720Z" && Test.Item.M_date <= "2016-09-12T15:42:22.720Z"'
    /*server.dataManager.read("Test", JSON.stringify(query), [], [],
        function (results, caller) {
            console.log("-------> results: ", results)
        },
        function (index, total, caller) {

        }, function (errMsg, caller) {

        }, undefined)*/

    query.TypeName = "CargoEntities.User"
    query.Fields = ["M_id", "M_firstName", "M_lastName", "M_email"]
    query.Query = '(CargoEntities.User.M_firstName ~= "Eric" || CargoEntities.User.M_firstName == "Louis") && CargoEntities.User.M_lastName != "Boucher"'

    server.dataManager.read("CargoEntities", JSON.stringify(query), [], [],
        function (results, caller) {
            console.log("-------> results: ", results)
        },
        function (index, total, caller) {

        },
        function (errMsg, caller) {

        }, undefined)

    server.entityManager.getEntities("CargoEntities.User", "CargoEntities", '(CargoEntities.User.M_firstName ~= "Eric" || CargoEntities.User.M_firstName == "Louis") && CargoEntities.User.M_lastName != "Boucher"', 0, -1, [], true, false,
        // Sucess...
        function (results, caller) {
            console.log(results)
        },
        // Progress...
        function (index, total) {

        },
        function () {

        }, undefined)
}

function entityDump(id, typeName) {
    server.entityManager.getEntityPrototypes(typeName.split(".")[0],
        function (result, caller) {
            // Here I will initialyse the catalog...
            server.entityManager.getEntityById(typeName, typeName.split(".")[0], [id], false,
                function (result) {

                    // Here I will overload the way to display the name in the interface.
                    CargoEntities.User.prototype.getTitles = function () {
                        this.displayName = this.M_firstName + " " + this.M_lastName
                        return [this.M_id, this.displayName]
                    }

                    // Initialyse entities references..
                    var parent = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "position: absolute; width: auto; height: auto;" })
                    new EntityPanel(parent, typeName, function (entity) {
                        return function (panel) {
                            panel.setEntity(entity)
                            panel.header.display()
                        }
                    }(result), undefined, false, result, "")
                },
                function (errObj, caller) {
                    console.log(errObj)
                })
        })
}

function entitiesDump(typeName) {
    server.entityManager.getEntityPrototypes(typeName.split(".")[0],
        function (result) {
            // Here I will initialyse the catalog...
            server.entityManager.getEntities(typeName, typeName.split(".")[0], "", 0, -1, [], true, false,
                // Progress callback...
                function () {

                },
                // Success callback.
                function (results, caller) {
                    console.log("entity: ", results)
                    var parent = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "position: absolute; width: auto; height: auto;" })
                    for (var i = 0; i < results.length; i++) {
                        // Initialyse entities references..
                        new EntityPanel(parent, typeName, function (entity) {
                            return function (panel) {
                                //panel.header.display()
                                panel.setEntity(entity)
                            }
                        }(results[i]), undefined, false, results[i], "")
                    }
                },
                // Error callback.
                function (errMsg, caller) {

                })
        }, typeName)
}

// The an uplad file panel.
function TestUploadFile() {
    var parent = new Element(document.getElementsByTagName("body")[0], { "tag": "div" })
    var path = "/Test/Upload"
    var fileUploadPanel = new FilesPanel(parent, path,
        // filesLoadCallback
        function (filePanel) {
            filePanel.uploadFile(path, function () {

            })
        },
        // filesReadCallback
        function () {

        })

}

//////////////////////////////////////////////////////////////////////
// Test JS extension and permission.
//////////////////////////////////////////////////////////////////////
function testSayHello(to) {
    // Try 
    var params = []
    params.push(createRpcData(to, "STRING", "to"))

    server.executeJsFunction(
        "SayHello", // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            console.log(result)
        },
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        {} // The caller
    )
}

//////////////////////////////////////////////////////////////////////
// WebRtc test.
//////////////////////////////////////////////////////////////////////
function TestWebRtc1() {

    // First of all I will append a video element inside the page.
    var videoPanel = new Element(document.getElementsByTagName("body")[0], { "tag": "video", autoplay: "" })
    var constraints = {
        video: {
            mandatory: {
                minWidth: 640,
                minHeight: 480
            }
        },
        audio: true
    };
    if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|OperaMini/i.test(navigator.userAgent)) {
        // The user is using a mobile device, lower our minimum resolution
        constraints = {
            video: {
                mandatory: {
                    minWidth: 480,
                    minHeight: 320,
                    maxWidth: 1024,
                    maxHeight: 768
                }
            },
            audio: true
        };
    }
    if (hasUserMedia()) {
        navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia;
        navigator.getUserMedia(constraints,
            function (videoPanel) {
                return function (stream) {
                    videoPanel.element.src = window.URL.createObjectURL(stream);
                }
            }(videoPanel),
            function (err) { }
        );
    } else {
        alert("Sorry, your browser does not support getUserMedia.");
    }
}

// Take a selfy...
function TestWebRtc2() {
    var panel = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "display: table" })
    panel.appendElement({ "tag": "div", "style": "display:table-row" }).down().appendElement({ "tag": "video", "id": "video", autoplay: "", "style": "diplay: table-cell" })
        .appendElement({ "tag": "canvas", "id": "canvas", "style": "diplay: table-cell, min-width: 640px;" }).up()
        .appendElement({ "tag": "div", "style": "display:table-row; text-align: center;" }).down().appendElement({ "tag": "button", "id": "button", "style": "display: table-cell;", "innerHtml": "Selfy!" })

    var video = panel.getChildById("video")
    var canvas = panel.getChildById("canvas")
    var button = panel.getChildById("button")


    if (hasUserMedia()) {
        navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia;
        var streaming = false;
        navigator.getUserMedia({
            video: true,
            audio: false
        }, function (video) {
            return function (stream) {
                video.element.src = window.URL.createObjectURL(stream);

                streaming = true
            }
        }(video, canvas),
            function (error) {
                console.log("Raised an error when capturing:", error);
            });

        button.element.addEventListener('click',
            function (canvas, video) {
                return function (event) {
                    if (streaming) {
                        canvas.width = video.clientWidth;
                        canvas.height = video.clientHeight;
                        var context = canvas.getContext('2d');
                        context.drawImage(video, 0, 0);
                    }
                }
            }(canvas.element, video.element));
    } else {
        alert("Sorry, your browser does not support getUserMedia.");
    }
}