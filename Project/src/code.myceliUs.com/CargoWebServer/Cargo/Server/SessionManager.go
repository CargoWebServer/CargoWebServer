package Server

import (
	"errors"
	"log"
	"sort"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/JS"
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
	activeSessions map[string]*CargoEntities.Session

	sessionToCloseChannel chan struct {
		session *CargoEntities.Session
		err     chan *CargoEntities.Error
	}

	activeSessionsChannel chan struct {
		activeSessionsChan chan []*CargoEntities.Session
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
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())

	this.activeSessions = make(map[string]*CargoEntities.Session, 0)
	this.sessionToCloseChannel = make(chan struct {
		session *CargoEntities.Session
		err     chan *CargoEntities.Error
	})
	this.activeSessionsChannel = make(chan struct {
		activeSessionsChan chan []*CargoEntities.Session
	})

	go this.removeClosedSession()

	go this.run()

}

func (this *SessionManager) getId() string {
	return "SessionManager"
}

func (this *SessionManager) start() {
	log.Println("--> Start SessionManager")
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
		case activeSessionsChannel_ := <-this.activeSessionsChannel:
			activeSessionsChannel_.activeSessionsChan <- this.getActiveSessions()
		case sessionToClose := <-this.sessionToCloseChannel:
			sessionToClose.err <- this.closeSession_(sessionToClose.session)
		}
	}
}

/**
 * This function is used to clean all closed sessions
 */
func (this *SessionManager) removeClosedSession() {

	sessions, _ := GetServer().GetEntityManager().getEntitiesByType("CargoEntities.Session", "", CargoEntitiesDB)

	for i := 0; i < len(sessions); i++ {
		sessionId := sessions[i].GetObject().(*CargoEntities.Session).GetId()
		if GetServer().getConnectionById(sessionId) == nil {
			// The session is closed
			this.closeSession(sessions[i].GetObject().(*CargoEntities.Session))
		}
	}
}

func (this *SessionManager) closeSession_(session *CargoEntities.Session) *CargoEntities.Error {

	// Remove the JS session
	JS.GetJsRuntimeManager().CloseSession(session.GetId())

	// Delete the session entity
	sessionEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(session.UUID)

	if errObj != nil {
		return NewError(Utility.FileLine(), SESSION_UUID_NOT_FOUND_ERROR, SERVER_ERROR_CODE, errors.New("The session with uuid '"+session.UUID+"' was not found."))
	}

	// Send session close event
	eventData := make([]*MessageData, 2)

	// The closed session
	sessionInfo := new(MessageData)
	sessionInfo.Name = "sessionInfo"
	sessionInfo.Value = session
	eventData[0] = sessionInfo

	// The active session
	sessionsInfo := new(MessageData)
	sessionsInfo.Name = "sessionsInfo"
	sessionsInfo.Value = this.getActiveSessionByAccountId(session.M_accountPtr)
	eventData[1] = sessionsInfo

	evt, _ := NewEvent(LogoutEvent, SessionEvent, eventData)

	// Remove the session from active session
	delete(this.activeSessions, session.GetId())
	sessionEntity.DeleteEntity()

	// Send the event
	GetServer().GetEventManager().BroadcastEvent(evt)

	// Delete the session
	connection := GetServer().getConnectionById(session.GetId())
	if connection != nil {
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
	for _, session := range this.activeSessions {
		sessions = append(sessions, session)
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

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

/**
 * Authenticate a user with a given account name and password on a given ldap server.
 */
func (this *SessionManager) Login(accountName string, psswd string, serverId string, messageId string, sessionId string) *CargoEntities.Session {

	var session *CargoEntities.Session
	accountUuid := CargoEntitiesAccountExists(accountName)

	// Verify if the account exists
	if len(accountUuid) > 0 {
		// The accout exists. It will be initialized
		var accountEntity Entity
		accountEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(accountUuid)

		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
		account := accountEntity.GetObject().(*CargoEntities.Account)

		// Verify if the password it correct
		if _, ok := GetServer().GetLdapManager().m_configsInfo[serverId]; ok {
			if GetServer().GetLdapManager().Authenticate(serverId, account.M_id, psswd) == false {
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

			// Set the account holding this session
			session.SetAccountPtr(account)

			// Get the session by id.
			connection := GetServer().getConnectionById(sessionId)
			if connection != nil {
				addr := connection.GetAddrStr()
				computer, err := GetServer().GetLdapManager().GetComputerByIp(addr)
				if err == nil {
					session.SetComputerRef(computer)
				}
			}

			// Append the session in the account entity
			account.SetSessions(session)
			account.NeedSave = true

			// Save it
			accountEntity.SaveEntity()
			this.activeSessions[session.GetId()] = session

			// Send session close event
			eventData := make([]*MessageData, 2)

			sessionInfo := new(MessageData)
			sessionInfo.Name = "sessionInfo"
			sessionInfo.Value = session
			eventData[0] = sessionInfo

			// Append the user data
			sessionsInfo := new(MessageData)
			sessionsInfo.Name = "sessionsInfo"
			sessionsInfo.Value = this.GetActiveSessionByAccountId(account.M_id)
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

/**
 * Close a user session
 */
func (this *SessionManager) Logout(toCloseId string, messageId string, sessionId string) {
	// Simply close the session
	currentSession := this.GetActiveSessionById(toCloseId)

	if currentSession != nil {
		this.closeSession(currentSession)
	} else {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), SESSION_ID_NOT_ACTIVE, SERVER_ERROR_CODE, errors.New("The session with the id '"+sessionId+"' is not active"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
	}
}

/**
 * Return the list of all active sessions on the server.
 */
func (this *SessionManager) GetActiveSessions() []*CargoEntities.Session {
	activeSessions := new(struct {
		activeSessionsChan chan []*CargoEntities.Session
	})

	activeSessions.activeSessionsChan = make(chan []*CargoEntities.Session)

	this.activeSessionsChannel <- *activeSessions

	return <-activeSessions.activeSessionsChan
}

/**
 * This function returns the session with a given id, if the session is active.
 */
func (this *SessionManager) GetActiveSessionById(sessionId string) *CargoEntities.Session {

	activeSessions := this.GetActiveSessions()

	for i := 0; i < len(activeSessions); i++ {
		if activeSessions[i].GetId() == sessionId {
			return activeSessions[i]
		}
	}
	return nil
}

/**
 * Returns the list of sessions for a given accout.
 */
func (this *SessionManager) GetActiveSessionByAccountId(accountId string) Sessions {
	var sessions Sessions
	for _, session := range this.GetActiveSessions() {
		if session.M_accountPtr == accountId {
			sessions = append(sessions, session)
		}
	}
	sort.Sort(sessions)
	return sessions
}

/**
 * Update the state of a session.
 */
func (this *SessionManager) UpdateSessionState(state int64, messageId string, sessionId string) {

	// Get the session object...
	session := this.GetActiveSessionById(sessionId)

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
	sessionInfo.Name = "sessionInfo"
	sessionInfo.Value = session
	eventData[0] = sessionInfo

	// Append the user data
	sessionsInfo := new(MessageData)
	sessionsInfo.Name = "sessionsInfo"
	sessionsInfo.Value = this.GetActiveSessionByAccountId(session.M_accountPtr)
	eventData[1] = sessionsInfo

	evt, _ := NewEvent(StateChangeEvent, SessionEvent, eventData)
	GetServer().GetEventManager().BroadcastEvent(evt)
}
