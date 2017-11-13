package Server

import (
	"errors"
	"log"
	"sort"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

/**
 * This interface is used to sort sessions according to the time
 */
type Sessions []*CargoEntities.Session

// Len is part of sort.Interface.
func (s Sessions) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s Sessions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (s Sessions) Less(i, j int) bool {
	return s[i].M_startTime < s[j].M_startTime
}

type SessionManager struct {
	sessionToCloseChannel chan struct {
		session *CargoEntities.Session
		err     chan *CargoEntities.Error
	}
}

var sessionManager *SessionManager

func (this *Server) GetSessionManager() *SessionManager {
	if sessionManager == nil {
		sessionManager = newSessionManager()
	}
	return sessionManager
}

/**
 * This function creates and return the session manager...
 */
func newSessionManager() *SessionManager {
	sessionManager := new(SessionManager)
	return sessionManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Do initialization stuff here.
 */
func (this *SessionManager) initialize() {

	log.Println("--> Initialize SessionManager")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)

	this.sessionToCloseChannel = make(chan struct {
		session *CargoEntities.Session
		err     chan *CargoEntities.Error
	})

	go this.removeClosedSession()

	go this.run()

}

func (this *SessionManager) getId() string {
	return "SessionManager"
}

func (this *SessionManager) start() {
	log.Println("--> Start SessionManager")
	this.removeClosedSession()
}

func (this *SessionManager) stop() {
	log.Println("--> Stop SessionManager")
}

/**
 * Processing message from outside threads
 */
func (this *SessionManager) run() {

	for {
		select {
		case sessionToClose := <-this.sessionToCloseChannel:
			sessionToClose.err <- this.closeSession_(sessionToClose.session)
		}
	}
}

/**
 * This function is used to clean all closed sessions
 */
func (this *SessionManager) removeClosedSession() {

	sessions := this.getActiveSessions()

	for i := 0; i < len(sessions); i++ {
		sessionId := sessions[i].GetId()
		if GetServer().getConnectionById(sessionId) == nil {
			// The session is closed
			this.closeSession(sessions[i])
		}
	}

}

func (this *SessionManager) closeSession_(session *CargoEntities.Session) *CargoEntities.Error {
	// Delete the session entity
	sessionEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(session.UUID, false)

	if errObj != nil {
		return NewError(Utility.FileLine(), SESSION_UUID_NOT_FOUND_ERROR, SERVER_ERROR_CODE, errors.New("The session with uuid '"+session.UUID+"' was not found."))
	}

	// Send session close event
	eventData := make([]*MessageData, 2)

	// The closed session
	sessionInfo := new(MessageData)
	sessionInfo.TYPENAME = "Server.MessageData"
	sessionInfo.Name = "sessionInfo"
	sessionInfo.Value = session
	eventData[0] = sessionInfo

	// The active session
	sessionsInfo := new(MessageData)
	sessionsInfo.TYPENAME = "Server.MessageData"
	sessionsInfo.Name = "sessionsInfo"
	sessionsInfo.Value = this.getActiveSessionByAccountId(session.M_accountPtr)
	eventData[1] = sessionsInfo

	evt, _ := NewEvent(LogoutEvent, SessionEvent, eventData)

	sessionEntity.DeleteEntity()

	// Send the event
	GetServer().GetEventManager().BroadcastEvent(evt)

	// Delete the session
	connection := GetServer().getConnectionById(session.GetId())
	if connection != nil {
		// Remove the vm for the session.
		connection.Close()
	}

	return nil
}

/**
 * Remove the connection for a given user
 */
func (this *SessionManager) closeSession(session *CargoEntities.Session) *CargoEntities.Error {

	sessionToClose := new(struct {
		session *CargoEntities.Session
		err     chan *CargoEntities.Error
	})

	sessionToClose.session = session

	sessionToClose.err = make(chan *CargoEntities.Error)

	this.sessionToCloseChannel <- *sessionToClose

	err := <-sessionToClose.err

	return err

}

func (this *SessionManager) getActiveSessions() []*CargoEntities.Session {

	var sessions []*CargoEntities.Session
	entities, _ := GetServer().GetEntityManager().getEntities("CargoEntities.Account", "", "CargoEntities", false)
	for i := 0; i < len(entities); i++ {
		account := entities[i].GetObject().(*CargoEntities.Account)
		sessions = append(sessions, account.GetSessions()...)
	}

	return sessions

}

/**
 * Returns the list of sessions for a given accout.
 * Use locally only
 */
func (this *SessionManager) getActiveSessionByAccountId(accountId string) Sessions {
	var sessions Sessions
	for _, session := range this.getActiveSessions() {
		if session.M_accountPtr == accountId {
			sessions = append(sessions, session)
		}
	}
	sort.Sort(sessions)
	return sessions
}

func (this *SessionManager) getActiveSessionById(id string) *CargoEntities.Session {

	activeSessions := this.getActiveSessions()

	for i := 0; i < len(activeSessions); i++ {
		if activeSessions[i].GetId() == id {
			return activeSessions[i]
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

// @api 1.0
// Event handler function.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
//SessionManager.prototype.onEvent = function (evt) {
//    EventHub.prototype.onEvent.call(this, evt)
//}
func (this *SessionManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

// @api 1.0
// Authenticate the user's account name and password on the server.
// @see LoginPage
// @param {string} name The account name
// @param {string} password The account password
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Session} The created session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SessionManager) Login(accountName string, psswd string, serverId string, messageId string, sessionId string) *CargoEntities.Session {

	var session *CargoEntities.Session
	accountUuid := CargoEntitiesAccountExists(accountName)

	// Verify if the account exists
	if len(accountUuid) > 0 {
		// The accout exists. It will be initialized
		var accountEntity Entity
		accountEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(accountUuid, false)

		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
		account := accountEntity.GetObject().(*CargoEntities.Account)

		// Verify if the password is correct
		if _, ok := GetServer().GetLdapManager().getConfigsInfo()[serverId]; ok {
			if GetServer().GetLdapManager().authenticate(serverId, account.M_id, psswd) == false {
				if account.M_password != psswd {

					// Create the error message
					cargoError := NewError(Utility.FileLine(), PASSWORD_MISMATCH_ERROR, SERVER_ERROR_CODE, errors.New("The password '"+psswd+"' does not match the account name '"+accountName+"'. "))
					GetServer().reportErrorMessage(messageId, sessionId, cargoError)

					return nil
				}
			}
		} else {
			if account.M_password != psswd {
				// Create the error message
				cargoError := NewError(Utility.FileLine(), PASSWORD_MISMATCH_ERROR, SERVER_ERROR_CODE, errors.New("The password '"+psswd+"' does not match the account name '"+accountName+"'. "))
				GetServer().reportErrorMessage(messageId, sessionId, cargoError)
				return nil
			}
		}

		sessionUuid := CargoEntitiesSessionExists(sessionId)
		if len(sessionUuid) == 0 {

			// If the session does not exist
			session = new(CargoEntities.Session)

			// The connection id is the same as the session id.
			session.UUID = "CargoEntities.Session%" + sessionId
			session.TYPENAME = "CargoEntities.Session"
			session.M_id = sessionId

			// Initialization of fields
			session.M_startTime = int64(time.Now().Unix())
			session.M_statusTime = int64(time.Now().Unix())
			session.M_sessionState = CargoEntities.SessionState_Online
			session.ParentLnk = "M_sessions"

			//Set the computer reference.
			connection := GetServer().getConnectionById(sessionId)
			if connection != nil {
				addr := connection.GetAddrStr()
				computer, err := GetServer().GetLdapManager().getComputerByIp(addr)
				if err == nil {
					session.SetComputerRef(computer)
				}
			}

			// Set the account ptr.
			session.SetAccountPtr(account)

			GetServer().GetEntityManager().createEntity(account.GetUUID(), "M_sessions", "CargoEntities.Session", session.GetId(), session)

			// Send session close event
			eventData := make([]*MessageData, 2)

			sessionInfo := new(MessageData)
			sessionInfo.TYPENAME = "Server.MessageData"
			sessionInfo.Name = "sessionInfo"
			sessionInfo.Value = session
			eventData[0] = sessionInfo

			// Append the user data
			sessionsInfo := new(MessageData)
			sessionsInfo.Name = "sessionsInfo"
			sessionsInfo.Value = this.getActiveSessionByAccountId(account.M_id)
			eventData[1] = sessionsInfo

			evt, _ := NewEvent(LoginEvent, SessionEvent, eventData)

			// Send the event
			GetServer().GetEventManager().BroadcastEvent(evt)

		} else {
			// Power clicker protection
			log.Println("session aleready exist...", sessionUuid)
		}

		// Return the active session for that account
		return session

	} else {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ACCOUNT_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The account '"+accountName+"' doesn't exist"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)

		return nil
	}
	return nil
}

// @api 1.0
// Close a user session on the server. A logout event is throw to inform other participant that the session
// is closed.
// @param {string} toCloseId The id of the session to close.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SessionManager) Logout(toCloseId string, messageId string, sessionId string) {
	// Simply close the session
	currentSession := this.getActiveSessionById(toCloseId)

	if currentSession != nil {
		this.closeSession(currentSession)
	} else {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), SESSION_ID_NOT_ACTIVE, SERVER_ERROR_CODE, errors.New("The session with the id '"+sessionId+"' is not active"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
	}
}

// @api 1.0
// Get the list of all active session on the server for a given account name
// @param {string} accountId The account name.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {[]*CargoEntities.Session} Return an array of session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SessionManager) GetActiveSessionByAccountId(accountId string, messageId string, sessionId string) Sessions {
	sessions := this.getActiveSessionByAccountId(accountId)
	return sessions
}

// @api 1.0
// Change the state of a given session.
// @param {int} state 1: Online, 2:Away, other: Offline.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SessionManager) UpdateSessionState(state int64, messageId string, sessionId string) {

	// Get the session object...
	session := this.getActiveSessionById(sessionId)

	if session == nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), SESSION_ID_NOT_ACTIVE, SERVER_ERROR_CODE, errors.New("The session with the id '"+sessionId+"' is not active"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	// The last session time
	if state == 1 {
		// Online
		session.M_statusTime = session.M_startTime // reset to start time
		session.M_sessionState = CargoEntities.SessionState_Online
	} else if state == 2 {
		// Away
		session.M_statusTime = int64(time.Now().Unix())
		session.M_sessionState = CargoEntities.SessionState_Away
	} else {
		// Offline
		session.M_statusTime = int64(time.Now().Unix())
		session.M_sessionState = CargoEntities.SessionState_Offline
	}

	// Send session close event
	eventData := make([]*MessageData, 2)

	// Append the user data
	sessionInfo := new(MessageData)
	sessionInfo.TYPENAME = "Server.MessageData"
	sessionInfo.Name = "sessionInfo"
	sessionInfo.Value = session
	eventData[0] = sessionInfo

	// Append the user data
	sessionsInfo := new(MessageData)
	sessionsInfo.Name = "sessionsInfo"
	sessionsInfo.Value = this.getActiveSessionByAccountId(session.M_accountPtr)
	eventData[1] = sessionsInfo

	evt, _ := NewEvent(StateChangeEvent, SessionEvent, eventData)
	GetServer().GetEventManager().BroadcastEvent(evt)
}
