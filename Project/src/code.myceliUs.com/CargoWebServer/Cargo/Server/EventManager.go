package Server

import (
	"encoding/json"

	"log"
	"regexp"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

const (
	AccountEvent                   = "AccountEvent"
	AccountRegisterSuccessEvent    = 0
	AccountConfirmationSucessEvent = 1
	SessionEvent                   = "SessionEvent"
	LoginEvent                     = 4
	LogoutEvent                    = 5
	StateChangeEvent               = 6
	BpmnEvent                      = "BpmnEvent"
	NewProcessInstanceEvent        = 7
	UpdateProcessInstanceEvent     = 8
	NewDefinitionsEvent            = 9
	DeleteDefinitionsEvent         = 10
	UpdateDefinitionsEvent         = 11
	EntityEvent                    = "EntityEvent"
	NewEntityEvent                 = 12
	UpdateEntityEvent              = 13
	DeleteEntityEvent              = 14
	OpenEntityEvent                = 15
	CloseEntityEvent               = 16
	FileEvent                      = "FileEvent"
	NewFileEvent                   = 17
	DeleteFileEvent                = 18
	UpdateFileEvent                = 19
	OpenFileEvent                  = 20
	CloseFileEvent                 = 21
	DataEvent                      = "DataEvent"
	DeleteRowEvent                 = 22
	NewRowEvent                    = 23
	UpdateRowEvent                 = 24
	NewDataStoreEvent              = 25
	DeleteDataStoreEvent           = 26
	SecurityEvent                  = "SecurityEvent"
	NewRoleEvent                   = 27
	DeleteRoleEvent                = 28
	UpdateRoleEvent                = 29
	PrototypeEvent                 = "PrototypeEvent"
	NewPrototypeEvent              = 30
	UpdatePrototypeEvent           = 31
	DeletePrototypeEvent           = 32
	ProjectEvent                   = "ProjectEvent"
	EmailEvent                     = "EmailEvent"
	ServiceEvent                   = "ServiceEvent"
	ConfigurationEvent             = "ConfigurationEvent"
	NewTaskEvent                   = 33
	UpdateTaskEvent                = 34
	EventEvent                     = "EventEvent"
	LdapEvent                      = "LdapEvent"
	OAuth2Event                    = "OAuth2Event"
	SchemaEvent                    = "SchemaEvent"
	WorkflowEvent                  = "WorkflowEvent"
)

////////////////////////////////////////////////////////////////////////////////
// The event manager
////////////////////////////////////////////////////////////////////////////////
type EventManager struct {

	// The map of avalaible event channels...
	m_channels     map[string]*EventChannel
	m_eventDataMap map[*Event]string

	// Concurent map access.
	m_opChannel chan map[string]interface{}
}

var eventManager *EventManager

func (this *Server) GetEventManager() *EventManager {
	if eventManager == nil {
		eventManager = newEventManager()
	}
	return eventManager
}

/**
 * A singleton that manage the event channels...
 */
func newEventManager() *EventManager {
	eventManager := new(EventManager)
	eventManager.m_channels = make(map[string]*EventChannel, 0)
	eventManager.m_opChannel = make(chan map[string]interface{}, 0)

	// Concurrency...
	go func() {
		for {
			select {
			case op := <-eventManager.m_opChannel:
				if op["op"] == "appendEventData" {
					dataStr := op["dataStr"].(string)
					evt := op["evt"].(*Event)
					if eventManager.m_eventDataMap != nil {
						eventManager.m_eventDataMap[evt] = dataStr
					}
				} else if op["op"] == "getEventData" {
					op["result"].(chan string) <- eventManager.m_eventDataMap[op["evt"].(*Event)]
				} else if op["op"] == "removeClosedListener" {
					for _, channel := range eventManager.m_channels {
						for _, listener := range channel.m_listeners {
							if listener.m_addr.IsOpen() == false {
								channel.removeEventListener(listener)
							}
						}
						if len(channel.m_listeners) == 0 {
							delete(eventManager.m_channels, channel.m_eventName)
						}
					}
				} else if op["op"] == "BroadcastEvent" {
					evt := op["evt"].(*Event)
					channel := eventManager.m_channels[evt.GetName()]
					if channel != nil {
						channel.broadcastEvent(evt)
					}
					delete(eventManager.m_eventDataMap, evt)
				} else if op["op"] == "BroadcastEventTo" {
					evt := op["evt"].(*Event)
					to := op["to"].(*CargoEntities.Account)
					channel := eventManager.m_channels[evt.GetName()]
					if channel != nil {
						channel.broadcastEventTo(evt, to)
					}
				} else if op["op"] == "AddEventListener" {
					listener := op["listener"].(*EventListener)
					channel := eventManager.m_channels[listener.getEventName()]

					if eventManager.m_channels[listener.getEventName()] == nil {
						channel = new(EventChannel)
						channel.m_eventName = listener.getEventName()
						channel.m_listeners = make(map[string]*EventListener, 0)
						eventManager.m_channels[channel.m_eventName] = channel
					}

					// append the listener
					channel.m_listeners[listener.getId()] = listener
				} else if op["op"] == "RemoveEventListener" {
					id := op["id"].(string)
					name := op["name"].(string)

					// Remove the listener
					listener := eventManager.m_channels[name].m_listeners[id]
					eventManager.m_channels[name].removeEventListener(listener)

					// if the channel is empty remove the channel...
					if len(eventManager.m_channels) == 0 {
						delete(eventManager.m_channels, name)
					}
				} else if op["op"] == "AppendEventFilter" {
					filter := op["filter"].(string)
					channelId := op["channelId"].(string)
					sessionId := op["sessionId"].(string)
					if eventManager.m_channels[channelId] != nil {
						listener := eventManager.m_channels[channelId].m_listeners[sessionId]
						if listener != nil {
							listener.appendFilter(filter)
						}
					}
				}
			}
		}
	}()
	return eventManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Do intialysation stuff here.
 */
func (this *EventManager) initialize() {
	log.Println("--> Initialize EventManager")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)

	this.m_eventDataMap = make(map[*Event]string, 0)
}

func (this *EventManager) getId() string {
	return "EventManager"
}

func (this *EventManager) start() {
	log.Println("--> Start EventManager")
}

func (this *EventManager) stop() {
	log.Println("--> Stop EventManager")
}

/**
 * 	Append event data to the m_eventDataMap
 */
func (this *EventManager) appendEventData(evt *Event, dataStr string) {
	arguments := make(map[string]interface{})
	arguments["op"] = "appendEventData"
	arguments["dataStr"] = dataStr
	arguments["evt"] = evt
	this.m_opChannel <- arguments
}

/**
 * 	Get an event string
 */
func (this *EventManager) getEventData(evt *Event) string {
	arguments := make(map[string]interface{})
	arguments["op"] = "getEventData"
	arguments["evt"] = evt
	result := make(chan string)
	arguments["result"] = result
	this.m_opChannel <- arguments
	return <-result
}

/**
 * Remove close listener when one connection is close...
 */
func (this *EventManager) removeClosedListener() {
	arguments := make(map[string]interface{})
	arguments["op"] = "removeClosedListener"
	this.m_opChannel <- arguments
}

////////////////////////////////////////////////////////////////////////////////
// The event listener
////////////////////////////////////////////////////////////////////////////////

/**
 * When event need to be handle by the server...
 */
type EventListener struct {
	// uuid
	m_id string
	// the type of event, use by channel
	m_eventName string
	// the listener addresse...
	m_addr *WebSocketConnection

	m_filters []string

	m_opChannel chan map[string]interface{}
}

/**
 * Create a new listener with a given name...
 */
func NewEventListener(eventName string, conn *WebSocketConnection) *EventListener {
	listner := new(EventListener)
	listner.m_addr = conn
	listner.m_eventName = eventName
	listner.m_id = conn.GetUuid()

	listner.m_opChannel = make(chan map[string]interface{}, 0)

	// Concurrency...
	go func() {
		for {
			select {
			case op := <-listner.m_opChannel:
				if op["op"] == "appendFilter" {
					filter := op["filter"].(string)
					if !Utility.Contains(listner.m_filters, filter) {
						listner.m_filters = append(listner.m_filters, filter)
					}
				} else if op["op"] == "removeFilter" {
					filter := op["filter"].(string)
					var filters []string
					for _, f := range listner.m_filters {
						if f != filter {
							filters = append(filters, f)
						}
					}
					listner.m_filters = filters
				} else if op["op"] == "GetFilter" {
					op["result"].(chan string) <- listner.m_filters[op["index"].(int)]
				}
			}
		}
	}()
	return listner
}

/**
 * Append a new filter
 */
func (this *EventListener) appendFilter(filter string) {
	arguments := make(map[string]interface{})
	arguments["op"] = "appendFilter"
	arguments["filter"] = filter
	this.m_opChannel <- arguments
}

/**
 * Remove a filter
 */
func (this *EventListener) removeFilter(filter string) {
	arguments := make(map[string]interface{})
	arguments["op"] = "removeFilter"
	arguments["filter"] = filter
	this.m_opChannel <- arguments
}

/**
 * Get a filter by index
 */
func (this *EventListener) GetFilter(index int) string {
	arguments := make(map[string]interface{})
	arguments["op"] = "GetFilter"
	arguments["index"] = index
	result := make(chan string)
	arguments["result"] = result
	this.m_opChannel <- arguments
	return <-result
}

// The uuid
func (this *EventListener) getId() string {
	return this.m_id
}

// Return the name of the listener, the same name as event...
func (this *EventListener) getEventName() string {
	return this.m_eventName
}

// Evaluates if an event needs to be sent by evaluating the filters
func (this *EventListener) evaluateFilter(evt *Event) bool {
	evtStr := GetServer().GetEventManager().getEventData(evt)
	//log.Println("284 ----------> ", this.m_filters)
	for _, filter := range this.m_filters {
		//log.Println("286 ----------> ", filter, evtStr)
		match, _ := regexp.MatchString(filter, evtStr)
		if match {
			//log.Println("289 ----------> filter match: ", evtStr)
			return true
		}
	}
	return false
}

func (this *EventListener) onEvent(evt *Event) {

	// Apply the filter
	// if the filter matches the event will be sent on the network
	needSend := true //this.evaluateFilter(evt)

	if needSend {
		// Do stuff here...
		// Create the protobuffer message...
		m := new(message)
		m.tryNb = 5
		m.from = this.m_addr
		m.to = append(m.to, this.m_addr)

		// Set the type to response
		m.msg = new(Message)
		uuid := Utility.RandomUUID()
		m.msg.Id = &uuid
		index_ := int32(-1)
		total := int32(1)
		m.msg.Index = &index_
		m.msg.Total = &total
		m.msg.Type = new(Message_MessageType)
		*m.msg.Type = Message_EVENT
		m.msg.Evt = evt

		// I will sent the event message to the listener...
		// Never send the message directly use the message processor for it.
		//log.Println("--> broadcast evt ", *m.msg.Id, " to ", this.m_addr.GetPort())
		GetServer().messageProcessor.m_outgoingChannel <- m
	}
}

////////////////////////////////////////////////////////////////////////////////
// The event channel
////////////////////////////////////////////////////////////////////////////////

/**
 * Event channel, each event type has a channel and listener subscribe to
 * to it.
 */
type EventChannel struct {
	// The name of the event manage by this channel...
	m_eventName string

	// The map of event listener...
	m_listeners map[string]*EventListener
}

/**
 * This funtion is use to broadcast the event over listener...
 */
func (this *EventChannel) broadcastEvent(evt *Event) {
	for _, listener := range this.m_listeners {
		listener.onEvent(evt)
	}
}

/**
 * Broadcast event to a specific account.
 */
func (this *EventChannel) broadcastEventTo(evt *Event, account *CargoEntities.Account) {
	for _, listener := range this.m_listeners {
		for i := 0; i < len(account.M_sessions); i++ {
			sessionId := account.M_sessions[i].M_id
			if sessionId == listener.m_id {
				listener.onEvent(evt)
			}
		}
	}
}

/**
 * Remove a listener from the channel
 */
func (this *EventChannel) removeEventListener(listener *EventListener) {
	delete(this.m_listeners, listener.m_id)
}

//////////////////////////////////////////////////////////////////////////////////
// API Event manager
//////////////////////////////////////////////////////////////////////////////////

// @api 1.0
// Event handler function.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
//EventManager.prototype.onEvent = function (evt) {
//    EventHub.prototype.onEvent.call(this, evt)
//}
func (this *EventManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

// @api 1.0
// Event to broadcast on the channel...
// @param {*Server.Event} evt Event to broadcast over the network.
// @scope {hidden}
func (this *EventManager) BroadcastEvent(evt *Event) {
	// Broadcast event over listener over the channel.
	arguments := make(map[string]interface{})
	arguments["op"] = "BroadcastEvent"
	arguments["evt"] = evt
	this.m_opChannel <- arguments
}

// @api 1.0
// Send event to specific account.
// @param {*Server.Event} evt Event to broadcast over the network.
// @param {*CargoEntities.Account} to The target account
// @scope {hidden}
func (this *EventManager) BroadcastEventTo(evt *Event, to *CargoEntities.Account) {
	// Broadcast event over listener over the channel.
	arguments := make(map[string]interface{})
	arguments["op"] = "BroadcastEventTo"
	arguments["evt"] = evt
	arguments["to"] = to
	this.m_opChannel <- arguments
}

// @api 1.0
// Add a new event listener
// @param {*Server.EventListener} The event listener to append.
// @scope {public}
func (this *EventManager) AddEventListener(listener *EventListener) {
	// Create the channel if is not exist
	arguments := make(map[string]interface{})
	arguments["op"] = "AddEventListener"
	arguments["listener"] = listener
	this.m_opChannel <- arguments
}

// @api 1.0
// Remove a event listener from a given channel
// @param {string} id The event listener id to remove.
// @param {string} name The name of the channel where the listner lisen to.
// @scope {public}
func (this *EventManager) RemoveEventListener(id string, name string) {
	arguments := make(map[string]interface{})
	arguments["op"] = "RemoveEventListener"
	arguments["id"] = id
	arguments["name"] = name
	this.m_opChannel <- arguments
}

// @api 1.0
// Broadcast event over the network.
// @param {int} eventNumber The event number.
// @param {string} channelId The event type.
// @param {int} eventDatas An array of Message Data structures.
// @param {interface{}} eventDatas An array of Message Data structures.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
func (this *EventManager) BroadcastEventData(eventNumber int64, channelId string, eventDatas interface{}, messageId string, sessionId string) {
	// Create the new event objet...
	evt, _ := NewEvent(int32(eventNumber), channelId, eventDatas.([]*MessageData))

	b, err := json.Marshal(eventDatas.([]*MessageData))
	this.appendEventData(evt, string(b))

	if err != nil {
		cargoError := NewError(Utility.FileLine(), EVENT_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
	}

	// Broadcast it...
	this.BroadcastEvent(evt)
}

// @api 1.0
// Append a new filter to a listener
// @param {string} filter The filter to append
// @param {string} channelId The id of the channel.
// @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
func (this *EventManager) AppendEventFilter(filter string, channelId string, messageId string, sessionId string) {
	//log.Println("append event filter ", filter, " for type ", eventType, " to session ", sessionId)
	arguments := make(map[string]interface{})
	arguments["op"] = "AppendEventFilter"
	arguments["filter"] = filter
	arguments["channelId"] = channelId
	arguments["sessionId"] = sessionId
	this.m_opChannel <- arguments
}
