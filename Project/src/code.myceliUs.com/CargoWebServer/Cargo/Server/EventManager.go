package Server

import (
	"encoding/json"

	"log"
	"regexp"
	"sync"

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
	TableEvent                     = "TableEvent"
	DeleteRowEvent                 = 22
	NewRowEvent                    = 23
	UpdateRowEvent                 = 24
	SecurityEvent                  = "SecurityEvent"
	NewRoleEvent                   = 25
	DeleteRoleEvent                = 26
	UpdateRoleEvent                = 27
	PrototypeEvent                 = "PrototypeEvent"
	NewPrototypeEvent              = 28
	UpdatePrototypeEvent           = 29
	DeletePrototypeEvent           = 30
	ProjectEvent                   = "ProjectEvent"
	EmailEvent                     = "EmailEvent"
	ServiceEvent                   = "ServiceEvent"
	DataEvent                      = "DataEvent"
	ConfigurationEvent             = "ConfigurationEvent"
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

	// Use it to synchronize ressources.
	sync.Mutex
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
	this.Lock()
	defer this.Unlock()

	// No event need to be sent if the map is not initialyse...
	if this.m_eventDataMap != nil {
		this.m_eventDataMap[evt] = dataStr
	}
}

/**
 * 	Get an event string
 */
func (this *EventManager) getEventData(evt *Event) string {
	this.Lock()
	defer this.Unlock()
	return this.m_eventDataMap[evt]
}

/**
 * Remove close listener when one connection is close...
 */
func (this *EventManager) removeClosedListener() {
	this.Lock()
	defer this.Unlock()

	for _, channel := range this.m_channels {
		for _, listener := range channel.m_listeners {
			if listener.m_addr.IsOpen() == false {
				channel.removeEventListener(listener)
			}
		}
		if len(channel.m_listeners) == 0 {
			delete(this.m_channels, channel.m_eventName)
		}
	}
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
	m_addr connection

	m_filters []string

	sync.Mutex
}

/**
 * Append a new filter
 */
func (this *EventListener) appendFilter(filter string) {
	this.Lock()
	defer this.Unlock()
	if !Utility.Contains(this.m_filters, filter) {
		this.m_filters = append(this.m_filters, filter)
	}
}

/**
 * Remove a filter
 */
func (this *EventListener) removeFilter(filter string) {
	this.Lock()
	defer this.Unlock()
	var filters []string
	for _, f := range this.m_filters {
		if f != filter {
			filters = append(filters, f)
		}
	}
	this.m_filters = filters
}

/**
 * Get a filter by index
 */
func (this *EventListener) GetFilter(index int) string {
	this.Lock()
	defer this.Unlock()
	return this.m_filters[index]
}

/**
 * Create a new listener with a given name...
 */
func NewEventListener(eventName string, conn connection) *EventListener {
	listner := new(EventListener)
	listner.m_addr = conn
	listner.m_eventName = eventName
	listner.m_id = conn.GetUuid()
	return listner
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
		m.from = this.m_addr
		m.to = append(m.to, this.m_addr)

		// Set the type to response
		m.msg = new(Message)
		index_ := int32(-1)
		total := int32(1)
		m.msg.Index = &index_
		m.msg.Total = &total
		m.msg.Type = new(Message_MessageType)
		*m.msg.Type = Message_EVENT
		m.msg.Evt = evt

		// I will sent the event message to the listener...
		this.m_addr.Send(m.GetBytes())
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
	//log.Println("-----------> 285 ", evt)
	for _, listener := range this.m_listeners {
		//log.Println("----------> evt broadcast: ", evt, listener.m_id)
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
	this.Lock()
	defer this.Unlock()

	channel := this.m_channels[evt.GetName()]
	if channel != nil {
		channel.broadcastEvent(evt)
	}
	delete(this.m_eventDataMap, evt)
}

// @api 1.0
// Send event to specific account.
// @param {*Server.Event} evt Event to broadcast over the network.
// @param {*CargoEntities.Account} to The target account
// @scope {hidden}
func (this *EventManager) BroadcastEventTo(evt *Event, to *CargoEntities.Account) {
	this.Lock()
	defer this.Unlock()
	// Broadcast event over listener over the channel.
	channel := this.m_channels[evt.GetName()]
	if channel != nil {
		channel.broadcastEventTo(evt, to)
	}
}

// @api 1.0
// Add a new event listener
// @param {*Server.EventListener} The event listener to append.
// @scope {public}
func (this *EventManager) AddEventListener(listener *EventListener) {
	this.Lock()
	defer this.Unlock()
	// Create the channel if is not exist
	channel := this.m_channels[listener.getEventName()]

	if this.m_channels[listener.getEventName()] == nil {
		channel = new(EventChannel)
		channel.m_eventName = listener.getEventName()
		channel.m_listeners = make(map[string]*EventListener, 0)
		this.m_channels[channel.m_eventName] = channel
	}

	// append the listener
	channel.m_listeners[listener.getId()] = listener
	log.Println("---------> 395 ", this.m_channels)
}

// @api 1.0
// Remove a event listener from a given channel
// @param {string} id The event listener id to remove.
// @param {string} name The name of the channel where the listner lisen to.
// @scope {public}
func (this *EventManager) RemoveEventListener(id string, name string) {
	this.Lock()
	defer this.Unlock()
	// Remove the listener
	listener := this.m_channels[name].m_listeners[id]
	this.m_channels[name].removeEventListener(listener)

	// if the channel is empty remove the channel...
	if len(this.m_channels) == 0 {
		delete(this.m_channels, name)
	}
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
	if this.m_channels[channelId] != nil {
		listener := this.m_channels[channelId].m_listeners[sessionId]
		if listener != nil {
			listener.appendFilter(filter)
		}
	}
}
