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
	admin, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", []interface{}{"admin"})
	if admin == nil {

		// Create the account in memory.
		account := new(CargoEntities.Account)
		account.M_id = "admin"
		account.M_password = "adminadmin"
		account.M_name = "admin"
		account.M_email = ""
		_, err := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntitiesUuid(), "M_entities", "CargoEntities.Account", "admin", account)
		if err != nil {
			log.Panicln("--> create account admin: ", err)
		}
	}

	// Create the guest account if it doesn't exist
	guest, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", []interface{}{"guest"})
	if guest == nil {
		// Create the account in memory.
		account := new(CargoEntities.Account)
		account.M_id = "guest"
		account.M_password = "guest"
		account.M_name = "guest"
		account.M_email = ""
		_, err := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntitiesUuid(), "M_entities", "CargoEntities.Account", "guest", account)
		if err != nil {
			log.Panicln("--> create account admin: ", err)
		}
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
	account, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", []interface{}{name})
	if account == nil {
		// Create the account in memory.
		account := new(CargoEntities.Account)
		account.M_id = name
		account.M_password = password
		account.M_name = name
		account.M_email = email

		// Append the newly create account into the cargo entities
		GetServer().GetEntityManager().CreateEntity(GetServer().GetEntityManager().getCargoEntitiesUuid(), "M_entities", "CargoEntities.Account", name, account, "", "")
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

	account, cargoError := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", []interface{}{id})
	if cargoError != nil {
		// Create the error message
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	return account.(*CargoEntities.Account)
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

	user, cargoError := GetServer().GetEntityManager().getEntityById("CargoEntities.User", "CargoEntities", []interface{}{id})
	if cargoError != nil {
		// Create the error message
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	return user.(*CargoEntities.User)
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

	session := GetServer().GetSessionManager().getActiveSessionById(connectionId)

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

	return session.GetAccountPtr()
}
