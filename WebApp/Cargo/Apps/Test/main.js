var applicationName = document.getElementsByTagName("title")[0].text


var languageInfo = {
    "en": {
    }
}

/**
 * This function is the entry point of the application...
 */
 function main() {
    // Append filter to receive all session event message
    // on the sessionEvent channel.
    //securityTests()
    /*
    server.eventHandler.appendEventFilter(
         "CargoEntities.",
         "EntityEvent",
         function () {
             entitiesDump("CargoEntities", "CargoEntities.User")
         },
         function () { },
         undefined
     )
 

    server.eventHandler.appendEventFilter(
         "COLLADASchema.",
         "EntityEvent",
         function () {
             entitiesDump("COLLADASchema", "COLLADASchema.COLLADA")
         },
         function () { },
         undefined
     )
    */
     

    // utilityTests()
    //serverTests()
    //sessionTests()
    //languageManagerTests()
    elementTests()

    //accountTests()
    //fileTests()

    //dataTests()

    //entityTests()

    //testDynamicEntity()

    // entityDump("item_1", "Test.Item")

    //entitiesDump("COLLADASchema.COLLADA")

    //entitiesDump("XPDMXML.ProcessStructureType")
    //entitiesDump("DT3_informations.Workpoint")
    
    //testEntityQuery()

    //TestWebRtc2()

    // Test get media source...
    // TestUploadFile()

    // Test get bmpn defintion instance...
    /*
    server.entityManager.getEntityPrototypes("Test",
        // Success callback.
        function (result, caller) {
            server.entityManager.getEntityPrototypes("BPMN20",
                // Success callback.
                function (result, caller) {
                    server.entityManager.getEntityPrototypes("BPMS_Runtime",
                        // Success callback.
                        function (result, caller) {
                            server.entityManager.getEntityById("BPMN20.Definitions", "_1484846640138",
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
                                            } (result), undefined, false, result, "")
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
}


function testEntityQuery() {
    //{"TypeName":"CargoEntities.Log","Fields":["uuid"],"Indexs":["M_id=defaultErrorLogger"],"Query":""}
    var query = {}
    query.TypeName = "Test.Item"
    query.Fields = ["M_name", "M_description"]
    // Regex
    //query.Query = 'Test.Item.M_description == /Ceci est [a-z|\s|0-9]+/ && Test.Item.M_id != "item_5"'
    //query.Query = 'Test.Item.M_stringLst == /t[a-z]t[a-z](\\.)?/'
    query.Query = 'Test.Item.M_description ^= "Ceci"'
    // bool value
    // query.Query = 'Test.Item.M_inStock == true'
    // int value
    //query.Query = 'Test.Item.M_qte <= 10'
    // float value
    //query.Query = 'Test.Item.M_price <= 3.0'
    // Date... using the 8601 string format.
    //query.Query = 'Test.Item.M_date >= "2016-07-12T15:42:22.720Z" && Test.Item.M_date <= "2016-09-12T15:42:22.720Z"'

    server.dataManager.read("Test", JSON.stringify(query), [], [],
        function (results, caller) {
            console.log("-------> results: ", results)
        },
        function (index, total, caller) {

        }, function (errMsg, caller) {

        }, undefined)

    query.TypeName = "CargoEntities.User"
    query.Fields = ["M_id", "M_firstName", "M_lastName", "M_email"]
    query.Query = '(CargoEntities.User.M_firstName ~= "Eric" || CargoEntities.User.M_firstName == "Louis") && CargoEntities.User.M_lastName != "Boucher"'

    server.dataManager.read("CargoEntities", JSON.stringify(query), [], [],
        function (results, caller) {
            console.log("-------> results: ", results)
        }, function (errMsg, caller) {

        }, undefined)

    server.entityManager.getObjectsByType("CargoEntities.User", "CargoEntities", '(CargoEntities.User.M_firstName ~= "Eric" || CargoEntities.User.M_firstName == "Louis") && CargoEntities.User.M_lastName != "Boucher"',
        // Progress...
        function () {

        },
        // Sucess...
        function (results, caller) {
            console.log(results)
        },
        function () {

        }, undefined)
}

function entityDump(id, typeName) {
    server.entityManager.getEntityPrototypes(typeName.split(".")[0],
        function (result, caller) {
            // Here I will initialyse the catalog...
            server.entityManager.getEntityById(typeName, id,
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
                        }
                    } (result), undefined, false, result, "")
                },
                function () {
                })
        })
}

function entitiesDump(typeName) {
    server.entityManager.getEntityPrototypes(typeName.split(".")[0],
        function (result) {
            // Here I will initialyse the catalog...
            server.entityManager.getObjectsByType(typeName, typeName.split(".")[0], "",
                // Progress callback...
                function () {

                },
                // Success callback.
                function (results, caller) {
                    var parent = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "position: absolute; width: auto; height: auto;" })
                    for (var i = 0; i < results.length; i++) {
                        // Initialyse entities references..
                        new EntityPanel(parent, typeName, function (entity) {
                            return function (panel) {
                                panel.setEntity(entity)
                            }
                        } (results[i]), undefined, false, results[i], "")
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
            } (videoPanel),
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
        } (video, canvas),
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
            } (canvas.element, video.element));
    } else {
        alert("Sorry, your browser does not support getUserMedia.");
    }
}




