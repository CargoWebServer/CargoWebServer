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
// var server = new Server("localhost", "127.0.0.1", 9393)
var server = new Server("www.cargowebserver.com", "10.67.44.63", 9393)

// Amazon ec2 sever...
//var server = new Server("www.cargowebserver.com", "54.218.110.52", 9393)

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
            server.serviceManager = new ServiceManager()

            // Register the listener to the server.
            server.accountManager.registerListener()
            server.sessionManager.registerListener()
            server.fileManager.registerListener()
            server.entityManager.registerListener()
            server.dataManager.registerListener()
            server.emailManager.registerListener()
            server.projectManager.registerListener()
            server.securityManager.registerListener()
            server.serviceManager.registerListener()

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