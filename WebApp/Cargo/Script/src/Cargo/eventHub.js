AccountEvent = "AccountEvent"
AccountRegisterSuccessEvent = 0
AccountConfirmationSucessEvent = 1
SessionEvent = "SessionEvent"
LoginEvent = 4
LogoutEvent = 5
StateChangeEvent = 6
BpmnEvent = "BpmnEvent"
NewProcessInstanceEvent = 7
UpdateProcessInstanceEvent = 8
NewDefinitionsEvent = 9
DeleteDefinitionsEvent = 10
UpdateDefinitionsEvent = 11
EntityEvent = "EntityEvent"
NewEntityEvent = 12
UpdateEntityEvent = 13
DeleteEntityEvent = 14
OpenEntityEvent = 15
CloseEntityEvent = 16
FileEvent = "FileEvent"
NewFileEvent = 17
DeleteFileEvent = 18
UpdateFileEvent = 19
OpenFileEvent = 20
CloseFileEvent = 21
DataEvent = "DataEvent"
DeleteRowEvent = 22
NewRowEvent = 23
UpdateRowEvent = 24
NewDataStoreEvent = 25
DeleteDataStoreEvent = 26
SecurityEvent = "SecurityEvent"
NewRoleEvent = 27
DeleteRoleEvent = 28
UpdateRoleEvent = 29
PrototypeEvent = "PrototypeEvent"
NewPrototypeEvent = 30
UpdatePrototypeEvent = 31
DeletePrototypeEvent = 32
ProjectEvent = "ProjectEvent"
EmailEvent = "EmailEvent"
ServiceEvent = "ServiceEvent"
ConfigurationEvent = "ConfigurationEvent"
NewTaskEvent = 33
UpdateTaskEvent = 34
EventEvent = "EventEvent"
LdapEvent = "LdapEvent"
OAuth2Event = "OAuth2Event"
SchemaEvent = "SchemaEvent"
WorkflowEvent = "WorkflowEvent"

exports.AccountEvent = AccountEvent
exports.AccountRegisterSuccessEvent = AccountRegisterSuccessEvent
exports.AccountConfirmationSucessEvent = AccountConfirmationSucessEvent
exports.SessionEvent = SessionEvent
exports.LoginEvent = LoginEvent
exports.LogoutEvent = LogoutEvent
exports.StateChangeEvent = StateChangeEvent
exports.BpmnEvent = BpmnEvent
exports.NewProcessInstanceEvent = NewProcessInstanceEvent
exports.UpdateProcessInstanceEvent = UpdateProcessInstanceEvent
exports.NewDefinitionsEvent = NewDefinitionsEvent
exports.DeleteDefinitionsEvent = DeleteDefinitionsEvent
exports.UpdateDefinitionsEvent = UpdateDefinitionsEvent
exports.EntityEvent = EntityEvent
exports.NewEntityEvent = NewEntityEvent
exports.UpdateEntityEvent = UpdateEntityEvent
exports.DeleteEntityEvent = DeleteEntityEvent
exports.OpenEntityEvent = OpenEntityEvent
exports.CloseEntityEvent = CloseEntityEvent
exports.FileEvent = FileEvent
exports.NewFileEvent = NewFileEvent
exports.DeleteFileEvent = DeleteFileEvent
exports.UpdateFileEvent = UpdateFileEvent
exports.OpenFileEvent = OpenFileEvent
exports.CloseFileEvent = CloseFileEvent
exports.DataEvent = DataEvent
exports.DeleteRowEvent = DeleteRowEvent
exports.NewRowEvent = NewRowEvent
exports.UpdateRowEvent = UpdateRowEvent
exports.NewDataStoreEvent = NewDataStoreEvent
exports.DeleteDataStoreEvent = DeleteDataStoreEvent
exports.SecurityEvent = SecurityEvent
exports.NewRoleEvent = NewRoleEvent
exports.DeleteRoleEvent = DeleteRoleEvent
exports.UpdateRoleEvent = UpdateRoleEvent
exports.PrototypeEvent = PrototypeEvent
exports.NewPrototypeEvent = NewPrototypeEvent
exports.UpdatePrototypeEvent = UpdatePrototypeEvent
exports.DeletePrototypeEvent = DeletePrototypeEvent
exports.ProjectEvent = ProjectEvent
exports.EmailEvent = EmailEvent
exports.ServiceEvent = ServiceEvent
exports.ConfigurationEvent = ConfigurationEvent
exports.NewTaskEvent = NewTaskEvent
exports.UpdateTaskEvent = UpdateTaskEvent
exports.EventEvent = EventEvent
exports.LdapEvent = LdapEvent
exports.OAuth2Event = OAuth2Event
exports.SchemaEvent = SchemaEvent
exports.WorkflowEvent = WorkflowEvent

require("Cargo/utility")
 
/**
* EventHub contructor
* @constructor
* @param {string} channelId The id of the channel of events to manage
* @returns {EventHub}
* @stability 2
* @public true
*/
var EventHub = function (channelId) {

    this.id = randomUUID();

    this.channelId = channelId;

    this.observers = {}

    return this
}

/**
 * Register the hub as a listener to a given channel.
 */
EventHub.prototype.registerListener = function () {
    // Append to the event handler
    server.eventHandler.addEventListener(this,
        // callback
        function (listener) {
            console.log("Listener" + listener.id + "was registered to the channel ", listener.channelId)
        }
    )
}

/**
* Attach observer to a specific event.
* @param obeserver The observer to attach.
* @param eventId The event id.
* @param {function} updateFct The function to execute when the event is received.
* @stability 1
* @public true
*/
EventHub.prototype.attach = function (observer, eventId, updateFct) {
    observer.observable = this

    if (observer.id == undefined) {
        // observer needs a UUID
        observer.id = randomUUID()
    }

    if (this.observers[eventId] == undefined) {
        this.observers[eventId] = []
    }

    var observerExistsForEventId = false
    for (var i = 0; i < this.observers[eventId].length; i++) {
        if (this.observers[eventId][i].id == observer.id) {
            // only on obeserver with the same id are allowed.
            observerExistsForEventId = true
        }
    }

    if (!observerExistsForEventId) {
        this.observers[eventId].push(observer)
    }

    if (observer.updateFunctions == undefined) {
        observer.updateFunctions = {}
    }

    observer.updateFunctions[this.id + "_" + eventId] = updateFct
}

/**
* Detach observer from event.
* @param obeserver The to detach
* @param eventId The event id
* @stability 1
* @public true
*/
EventHub.prototype.detach = function (observer, eventId) {
    if (observer.observable != null) {
        observer.observable = null
    }
    if (observer.updateFunctions != undefined) {
        if (observer.updateFunctions[this.id + "_" + eventId] != null) {
            delete observer.updateFunctions[this.id + "_" + eventId]
            if (Object.keys(observer.updateFunctions).length == 0) {
                this.observers[eventId].pop(observer)
            }
        }
    }
}

/**
* When an event is received, the observer callback function is called.
* @param evt The event to dispatch.
* @stability 1
* @public false
*/
EventHub.prototype.onEvent = function (evt) {
    //console.log("Event received: ", evt)
    var observers = this.observers[evt.code]
    if (observers != undefined) {
        for (var i = 0; i < observers.length; i++) {
            if (observers[i].updateFunctions != undefined) {
                if (observers[i].updateFunctions[this.id + "_" + evt.code] != null) {
                    observers[i].updateFunctions[this.id + "_" + evt.code](evt, observers[i])
                } else {
                    if (Object.keys(observers[i].updateFunctions).length == 0) {
                        this.observers[eventId].pop(observers[i])
                    }
                }
            }
        }
    }
}

// exported class.
exports.EventHub = EventHub


