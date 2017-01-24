package Server

import (
	"errors"

	//	"log"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Persistence/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
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

/**
 * Do initialization stuff here.
 */
func (this *AccountManager) Initialize() {
	entities := GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities)
	// Create the admin account if it doesn't exist
	adminUuid := CargoEntitiesAccountExists("admin")
	if len(adminUuid) == 0 {
		// Create the account in memory.
		account := new(CargoEntities.Account)
		account.M_id = "admin"
		account.M_password = "adminadmin"
		account.M_name = "admin"
		account.M_email = ""

		// Append the newly create account into the cargo entities
		entities.SetEntities(account)
		account.SetEntitiesPtr(entities)
		// Save the account
		GetServer().GetEntityManager().getCargoEntities().SaveEntity()
	}

	// Create the guest account if it doesn't exist
	guestUuid := CargoEntitiesAccountExists("guest")
	if len(guestUuid) == 0 {
		// Create the account in memory.
		account := new(CargoEntities.Account)
		account.M_id = "guest"
		account.M_password = ""
		account.M_name = "guest"
		account.M_email = ""
		account.UUID = "CargoEntities.Account%" + Utility.RandomUUID()

		// Append the newly create account into the cargo entities
		entities.SetEntities(account)
		account.SetEntitiesPtr(entities)
		// Save the account
		GetServer().GetEntityManager().getCargoEntities().SaveEntity()
	}
}

/**
 * Function used to register an new account on the server.
 * @name of the new accout
 * @password the password
 * @email the email address
 * @messageId the id of the message at the root of the action.
 * @sessionId the id of the session at the root of the action.
 */
func (this *AccountManager) Register(name string, password string, email string, messageId string, sessionId string) *CargoEntities.Account {

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

		// Append the newly create account into the cargo entities
		entities := GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities)
		entities.SetEntities(account)
		account.SetEntitiesPtr(entities)

		// Save the account
		GetServer().GetEntityManager().getCargoEntities().SaveEntity()

		return account
	} else {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ACCOUNT_ALREADY_EXISTS_ERROR, SERVER_ERROR_CODE, errors.New("The account '"+name+"' already exists"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)

		return nil
	}

}

/**
 * Retreive an account with a given id.
 */
func (this *AccountManager) GetAccountById(name string, messageId string, sessionId string) *CargoEntities.Account {

	accountUuid := CargoEntitiesAccountExists(name)
	if len(accountUuid) == 0 {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ACCOUNT_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The account '"+name+"' doesn't exist"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	accountEntity, errObj := server.GetEntityManager().getEntityByUuid(accountUuid)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	return accountEntity.GetObject().(*CargoEntities.Account)
}

/**
* Retreive a user with a given id
 */
func (this *AccountManager) GetUserById(id string, messageId string, sessionId string) *CargoEntities.User {

	userUuid := CargoEntitiesUserExists(id)
	if len(userUuid) == 0 {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), USER_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The user with the id '"+id+"' doesn't exist"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	userEntity, errObj := server.GetEntityManager().getEntityByUuid(userUuid)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	return userEntity.GetObject().(*CargoEntities.User)
}
