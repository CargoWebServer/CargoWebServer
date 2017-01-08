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
 * @fileOverview Entities related functionalities.
 * @author Dave Courtois
 * @version 1.0
 */

/**
 * The project manager is use to manage project.
 * @constructor
 * @extends EventManager
 */
var ProjectManager = function (id) {

    if (server == undefined) {
        return
    }
    if (id == undefined) {
        id = randomUUID()
    }

    EventManager.call(this, id, ProjectEvent)

    /**
     * @property {object} entityPrototypes Keep track of prototypes in use.
     */
    this.entityPrototypes = {}

    return this
}

ProjectManager.prototype = new EventManager(null);
ProjectManager.prototype.constructor = ProjectManager;

/*
 * Dispatch event.
 */
ProjectManager.prototype.onEvent = function (evt) {
    EventManager.prototype.onEvent.call(this, evt)
}

ProjectManager.prototype.RegisterListener = function () {
    // Append to the event handler.
    server.eventHandler.AddEventManager(this,
        // callback
        function () {
            console.log("Entity manager is registered!")
        }
    )
}

/**
 * Retreive all project on the server.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
ProjectManager.prototype.GetAllProjects = function (successCallback, errorCallback, caller) {
    server.entityManager.getObjectsByType("CargoEntities.Project", "CargoEntities", "", function () { },
        function (result, caller) {
            caller.successCallback(result, caller.caller)
        },errorCallback, {"caller":caller, "successCallback": successCallback})
}