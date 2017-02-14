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
var server = new Server("localhost", "127.0.0.1", 9393)
//var server = new Server("MON104", "10.67.44.73", 9393)

// Amazon ec2 sever...
//var server = new Server("www.cargowebserver.com", "54.214.130.226", 8080)

//var server = new Server("MON176", "10.67.44.61", 9393)
//var server = new Server("MON-UTIL-01", "10.2.128.70", 9393)
//var server = new Server("localhost", "192.168.2.10", 9393)
// The main entry point

/**
 * Each application must have a function called main. The main function will
 * be call rigth after the initialization is completed.
 */
var main = null

// If the bpmn service is use..
var BPMS = false

/*
 * Append load function to the windows load event listener.
 */
function load() {

    // Open a new connection whit the web server.
    server.conn = initConnection("ws://" + server.ipv4 + ":" + server.port.toString() + "/ws",
        function () {
            // Get the session id from the server.
            server.setSessionId()

            // Create the listener
            server.accountManager = new AccountManager()
            server.sessionManager = new SessionManager()
            server.errorManager = new ErrorManager()
            server.fileManager = new FileManager()
            server.entityManager = new EntityManager()
            server.dataManager = new DataManager()
            server.emailManager = new EmailManager()
            server.projectManager = new ProjectManager()
            server.securityManager = new SecurityManager()

            // Register the listener to the server.
            server.accountManager.RegisterListener()
            server.sessionManager.RegisterListener()
            server.fileManager.RegisterListener()
            server.entityManager.RegisterListener()
            server.dataManager.RegisterListener()
            server.emailManager.RegisterListener()
            server.projectManager.RegisterListener()
            //server.securityManager.RegisterListener()
            if (BPMS) {
                server.workflowManager = new WorkflowManager()
                server.workflowManager.RegisterListener()
            }

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