package Server

import (
	"errors"

	"log"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

/**
 * This class implements functionalities to manage accounts
 */
type AccountManager struct {
}

var accountManager *AccountManager

/**
 * Create a new account manager
 */
func newAccountManager() *AccountManager {
	accountManager := new(AccountManager)
	return accountManager
}

func (this *Server) GetAccountManager() *AccountManager {
	if accountManager == nil {
		accountManager = newAccountManager()
	}
	return accountManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Do initialization stuff here.
 */
func (this *AccountManager) initialize() {

	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)

	// Create the admin account if it doesn't exist
	adminUuid := CargoEntitiesAccountExists("admin")
	if len(adminUuid) == 0 {
		// Create the account in memory.
		account := new(CargoEntities.Account)
		account.M_id = "admin"
		account.M_password = "adminadmin"
		account.M_name = "admin"
		account.NeedSave = true
		account.M_email = ""
		GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntities().GetUuid(), "M_entities", "CargoEntities.Account", "admin", account)
	}

	// Create the guest account if it doesn't exist
	guestUuid := CargoEntitiesAccountExists("guest")
	if len(guestUuid) == 0 {
		// Create the account in memory.
		account := new(CargoEntities.Account)
		account.M_id = "guest"
		account.M_password = "guest"
		account.M_name = "guest"
		account.M_email = ""
		account.NeedSave = true
		GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntities().GetUuid(), "M_entities", "CargoEntities.Account", "guest", account)
	}

}

func (this *AccountManager) getId() string {
	return "AccountManager"
}

func (this *AccountManager) start() {
	log.Println("--> Start AccountManager")
}

func (this *AccountManager) stop() {
	log.Println("--> Stop AccountManager")
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

// @api 1.0
// Event handler function.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
//AccountManager.prototype.onEvent = function (evt) {
//    EventHub.prototype.onEvent.call(this, evt)
//}
func (this *AccountManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

// @api 1.0
// Register a new account.
// @param {string} name The name of the new account.
// @param {string} password The password associated with the new account.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Account} The new registered account.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *AccountManager) Register(name string, password string, email string, messageId string, sessionId string) *CargoEntities.Account {

	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// The email address must be written in lower case.
	email = strings.ToLower(email)

	accountUuid := CargoEntitiesAccountExists(name)
	if len(accountUuid) == 0 {
		// Create the account in memory.
		account := new(CargoEntities.Account)
		account.M_id = name
		account.M_password = password
		account.M_name = name
		account.M_email = email
		account.NeedSave = true

		// Append the newly create account into the cargo entities
		GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntities().GetUuid(), "M_entities", "CargoEntities.Account", name, account)
		return account
	} else {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ACCOUNT_ALREADY_EXISTS_ERROR, SERVER_ERROR_CODE, errors.New("The account '"+name+"' already exists"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)

		return nil
	}

}

// @api 1.0
// Retreive an account with a given id.
// @param {string} id The id of the account.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Account} The retreived account.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *AccountManager) GetAccountById(id string, messageId string, sessionId string) *CargoEntities.Account {

	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	accountUuid := CargoEntitiesAccountExists(id)
	if len(accountUuid) == 0 {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ACCOUNT_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The account '"+id+"' doesn't exist"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	accountEntity, errObj := server.GetEntityManager().getEntityByUuid(accountUuid, false)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	return accountEntity.GetObject().(*CargoEntities.Account)
}

// @api 1.0
// Retreive a user with a given id.
// @param {string} id The id of the user.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *AccountManager) GetUserById(id string, messageId string, sessionId string) *CargoEntities.User {

	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	userUuid := CargoEntitiesUserExists(id)
	if len(userUuid) == 0 {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), USER_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The user with the id '"+id+"' doesn't exist"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	userEntity, errObj := server.GetEntityManager().getEntityByUuid(userUuid, false)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	return userEntity.GetObject().(*CargoEntities.User)
}

// @api 1.0
// Retreive a account with a given session id, sessionId is in fact the entity
// uuid and not the connection id.
// @param {string} connectionId The current socket connection id.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *AccountManager) Me(connectionId string, messageId string, sessionId string) *CargoEntities.Account {

	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	session := GetServer().GetSessionManager().activeSessions[connectionId]

	if session == nil {
		cargoError := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The session with the id '"+connectionId+"' doesn't exist"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	if session.GetAccountPtr() == nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The session with the id '"+connectionId+"' has not associated account."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	/*errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		log.Println("------> action: ", Utility.FunctionName())
		log.Println("------> error: ", errObj)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}*/

	return session.GetAccountPtr()
}
