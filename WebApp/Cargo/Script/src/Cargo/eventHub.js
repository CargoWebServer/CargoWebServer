AccountEvent = "AccountEvent";
AccountRegisterSuccessEvent = 0;
AccountConfirmationSucessEvent = 1;
SessionEvent = "SessionEvent";
LoginEvent = 4;
LogoutEvent = 5;
StateChangeEvent = 6;
EntityEvent = "EntityEvent";
NewEntityEvent = 7;
UpdateEntityEvent = 8;
DeleteEntityEvent = 9;
OpenEntityEvent = 10;
CloseEntityEvent = 11;
FileEvent = "FileEvent";
NewFileEvent = 12;
DeleteFileEvent = 13;
UpdateFileEvent = 14;
OpenFileEvent = 15;
CloseFileEvent = 16;
FileEditEvent = 17;
DataEvent = "DataEvent";
DeleteRowEvent = 18;
NewRowEvent = 19;
UpdateRowEvent = 20;
NewDataStoreEvent = 21;
DeleteDataStoreEvent = 22;
SecurityEvent = "SecurityEvent";
NewRoleEvent = 23;
DeleteRoleEvent = 24;
UpdateRoleEvent = 25;
PrototypeEvent = "PrototypeEvent";
NewPrototypeEvent = 26;
UpdatePrototypeEvent = 27;
DeletePrototypeEvent = 28;
ProjectEvent = "ProjectEvent";
EmailEvent = "EmailEvent";
ServiceEvent = "ServiceEvent";
ConfigurationEvent = "ConfigurationEvent";
NewTaskEvent = 29;
UpdateTaskEvent = 30;
EventEvent = "EventEvent";
LdapEvent = "LdapEvent";
OAuth2Event = "OAuth2Event";
SchemaEvent = "SchemaEvent";
WorkflowEvent = "WorkflowEvent";
NewBpmnDefinitionsEvent = 31;
DeleteBpmnDefinitionsEvent = 32;
UpdateBpmnDefinitionsEvent = 33;
StartProcessEvent = 34;

exports.AccountEvent = AccountEvent;
exports.AccountRegisterSuccessEvent = AccountRegisterSuccessEvent;
exports.AccountConfirmationSucessEvent = AccountConfirmationSucessEvent;
exports.SessionEvent = SessionEvent;
exports.LoginEvent = LoginEvent;
exports.LogoutEvent = LogoutEvent;
exports.StateChangeEvent = StateChangeEvent;
exports.EntityEvent = EntityEvent;
exports.NewEntityEvent = NewEntityEvent;
exports.UpdateEntityEvent = UpdateEntityEvent;
exports.DeleteEntityEvent = DeleteEntityEvent;
exports.OpenEntityEvent = OpenEntityEvent;
exports.CloseEntityEvent = CloseEntityEvent;
exports.FileEvent = FileEvent;
exports.NewFileEvent = NewFileEvent;
exports.DeleteFileEvent = DeleteFileEvent;
exports.UpdateFileEvent = UpdateFileEvent;
exports.OpenFileEvent = OpenFileEvent;
exports.CloseFileEvent = CloseFileEvent;
exports.DataEvent = DataEvent;
exports.DeleteRowEvent = DeleteRowEvent;
exports.NewRowEvent = NewRowEvent;
exports.UpdateRowEvent = UpdateRowEvent;
exports.NewDataStoreEvent = NewDataStoreEvent;
exports.DeleteDataStoreEvent = DeleteDataStoreEvent;
exports.SecurityEvent = SecurityEvent;
exports.NewRoleEvent = NewRoleEvent;
exports.DeleteRoleEvent = DeleteRoleEvent;
exports.UpdateRoleEvent = UpdateRoleEvent;
exports.PrototypeEvent = PrototypeEvent;
exports.NewPrototypeEvent = NewPrototypeEvent;
exports.UpdatePrototypeEvent = UpdatePrototypeEvent;
exports.DeletePrototypeEvent = DeletePrototypeEvent;
exports.ProjectEvent = ProjectEvent;
exports.EmailEvent = EmailEvent;
exports.ServiceEvent = ServiceEvent;
exports.ConfigurationEvent = ConfigurationEvent;
exports.NewTaskEvent = NewTaskEvent;
exports.UpdateTaskEvent = UpdateTaskEvent;
exports.EventEvent = EventEvent;
exports.LdapEvent = LdapEvent;
exports.OAuth2Event = OAuth2Event;
exports.SchemaEvent = SchemaEvent;
exports.WorkflowEvent = WorkflowEvent;
exports.NewBpmnDefinitionsEvent = NewBpmnDefinitionsEvent;
exports.DeleteBpmnDefinitionsEvent = DeleteBpmnDefinitionsEvent;
exports.UpdateBpmnDefinitionsEvent = UpdateBpmnDefinitionsEvent;
exports.StartProcessEvent = StartProcessEvent;

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

    this.observers = {};

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
    );
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
        observer.id = randomUUID();
    }

    if (this.observers[eventId] == undefined) {
        this.observers[eventId] = [];
    }

    var observerExistsForEventId = false
    for (var i = 0; i < this.observers[eventId].length; i++) {
        if (this.observers[eventId][i].id == observer.id) {
            // only on obeserver with the same id are allowed.
            observerExistsForEventId = true;
        }
    }

    if (!observerExistsForEventId) {
        this.observers[eventId].push(observer);
    }

    if (observer.updateFunctions == undefined) {
        observer.updateFunctions = {};
    }

    observer.updateFunctions[this.id + "_" + eventId] = updateFct;
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
        observer.observable = null;
    }
    if (observer.updateFunctions != undefined) {
        if (observer.updateFunctions[this.id + "_" + eventId] != null) {
            delete observer.updateFunctions[this.id + "_" + eventId]
            if (Object.keys(observer.updateFunctions).length == 0) {
                this.observers[eventId].pop(observer);
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
    var observers = this.observers[evt.code];
    if (observers != undefined) {
        for (var i = 0; i < observers.length; i++) {
            if (observers[i].updateFunctions != undefined) {
                if (observers[i].updateFunctions[this.id + "_" + evt.code] != null) {
                    observers[i].updateFunctions[this.id + "_" + evt.code](evt, observers[i]);
                } else {
                    if (Object.keys(observers[i].updateFunctions).length == 0) {
                        this.observers[eventId].pop(observers[i]);
                    }
                }
            }
        }
    }
}

// exported class.
exports.EventHub = EventHub;


