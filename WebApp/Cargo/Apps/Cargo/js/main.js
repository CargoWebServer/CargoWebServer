/*
 * (C) Copyright 2016 Mycelius SA (http://mycelius.com/).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * @fileOverview Initialization stuff.
 * @author Dave Courtois
 * @version 1.0
 */


/**
 * This global object represent a reference to the distant server.
 * @see Server
 */
//var server = new Server("localhost", "127.0.0.1", 9393)
//var server = new Server("mon-util-01", "10.2.128.70", 9393)
var server = new Server("mon176", "10.67.44.52", 9393)

// Amazon ec2 sever...
//var server = new Server("www.cargowebserver.com", "54.218.110.52", 9393)

// Empty language info, that object keep interface text that need to be translate.
var languageInfo = {}

// Keep the initialysed entitie in memory
var entities = {}

// Keep the list of entity prototype in memory
var entityPrototypes = {}

// Keep the active user session id here.
var activeSessionAccountId = null

/**
 * Each application must have a function called main. The main function will
 * be call rigth after the initialization is completed.
 */
var main = null

/*
 * Append load function to the windows load event listener.
 */
function load() {

    // Open a new connection whit the web server.
    server.init(function () {

        // Get the session id from the server.
        server.setSessionId(function () {

            server.errorManager = new ErrorManager()

            // Inject service code accessors.
            server.getServicesClientCode(
                function (results, caller) {
                    server.getRootPath(
                        function (path, caller) {
                            // set the path.
                            server.root = path
                            for (var key in results) {
                                // create the listener if is not already exist.
                                if (server[key] == undefined) {
                                    // inject the code in the client memory
                                    try{
                                        eval(results[key])
                                    }catch(err){
                                        console.log(err)
                                    }
                                    
                                    // Now I will create the listener
                                    var listenerName = key.charAt(0).toLowerCase() + key.slice(1);
                                    server[listenerName] = eval("new " + key + "()")
                                    server[listenerName].registerListener()
                                }
                            }

                            // Register prototype manager as listener.
                            server["prototypeManager"] = new EntityPrototypeManager(PrototypeEvent)
                            server["prototypeManager"].registerListener(PrototypeEvent)

                            // Go to the main entry point
                            // Append the listener for the entity.
                            // The session listener.
                            server.entityManager.getEntityPrototypes("Config", function (result, caller) {
                                server.entityManager.getEntityPrototypes("CargoEntities", function (result, initCallback) {
                                    /** Here I will set the list of available data source... */
                                    server.configurationManager.getActiveConfigurations(
                                        function (activeConfigurations) {
                                            // Get the active configuration.
                                            server.activeConfigurations = activeConfigurations
                                            if (main != null) {
                                                // So here I will connect the file listener...
                                                server.fileManager.attach(this, UpdateFileEvent, function (evt, codeEditor) {
                                                    if (evt.dataMap.fileInfo !== undefined) {
                                                        var fileInfo = evt.dataMap["fileInfo"]
                                                        var filePath = "http://" + server.hostName + ":" + server.port + fileInfo.M_path + "/" + fileInfo.M_name
                                                        /*if(fileInfo.M_mime == "application/javascript"){
                                                            // First I will unload the script file.
                                                            var old = null;
                                                            for(var i=0; i < document.scripts.length; i++){
                                                                if(document.scripts[i].src.startsWith(filePath)){
                                                                    old = document.scripts[i]
                                                                    break;
                                                                }
                                                            }
                                                            if(old != null){
                                                                old.parentElement.removeChild(old)
                                                            }
                                
                                                            // Now I will append the new script.
                                                            loadjscssfile(filePath + "?" + fileInfo.M_checksum, "application/javascript", function(){
                                                                console.log("---> the file is loaded!..")
                                                            })
                                                        }
                                                        
                                                        
                                                        // Read html text file.
                                                        if(fileInfo.M_mime == "text/html"){
                                                            server.fileManager.readTextFile(fileInfo.M_path + "/" + fileInfo.M_name, 
                                                            function(results, caller){
                                                                
                                                                if(caller.M_name == "index.html"){
                                                                    //document.body.parentElement.innerHTML = results[0]
                                                                    //document.open("text/html", "replace");
                                                                    //document.write(results[0]);  // htmlCode is the variable you called newDocument
                                                                }
                                                            },
                                                            function(){
                                                                
                                                            }, fileInfo)
                                                        }*/
                                                    }
                                                })
                                                main()
                                            }
                                        },
                                        function () {

                                        }, {})

                                }, function () {/* Error callback */ }, null)
                            }, function () {/* Error callback */ }, {})
                        },
                        function (erroObj, caller) {

                        },
                        {})
                },
                function (errObj, caller) {
                    // Here no client service code was found.
                }, {})
        })

    },  // onOpen callback
        function () { // onClose callback
            //alert('The connection was closed!!!')
            location.reload()
        });
}

if (window.addEventListener) {
    window.addEventListener('load', load, false);
}
else if (window.attachEvent) {
    window.attachEvent('onload', load);
}