package Server

import (
	"errors"
	"log"

	//"strings"

	"regexp"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

/**
 * Entity security functionality.
 */
type SecurityManager struct {
}

var securityManager *SecurityManager

func (this *Server) GetSecurityManager() *SecurityManager {
	if securityManager == nil {
		securityManager = newSecurityManager()
	}
	return securityManager
}

func newSecurityManager() *SecurityManager {
	securityManager := new(SecurityManager)

	return securityManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

func (this *SecurityManager) createAdminRole() {
	//adminRoleUuid Create the admin role if it doesn't exist
	entity, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{"adminRole"})
	if entity == nil {
		ids := []interface{}{"admin"}
		adminAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", ids)
		if errObj != nil {
			return
		}
		adminAccount := adminAccountEntity.(*CargoEntities.Account)

		// Create adminRole
		adminRole, _ := this.createRole("adminRole")
		adminRole.AppendAccountsRef(adminAccount)
		adminAccount.AppendRolesRef(adminRole)

		GetServer().GetEntityManager().saveEntity(adminRole)
		GetServer().GetEntityManager().saveEntity(adminAccount)

	}
}

func (this *SecurityManager) createGuestRole() {
	// Create the guest role if it doesn't exist
	entity, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{"guestRole"})
	if entity == nil {
		ids := []interface{}{"guest"}
		guestAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", ids)
		if errObj != nil {
			return
		}
		guestAccount := guestAccountEntity.(*CargoEntities.Account)

		// Create guestRole
		guestRole, _ := this.createRole("guestRole")
		guestRole.AppendAccountsRef(guestAccount)

		// Setting guestRole to guest account
		guestAccount.AppendRolesRef(guestRole)

		GetServer().GetEntityManager().saveEntity(guestRole)
		GetServer().GetEntityManager().saveEntity(guestAccount)

	}
}

/**
 * Initialize security related information.
 */
func (this *SecurityManager) initialize() {

	log.Println("--> Initialize Security manager")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)

	this.createAdminRole()
	this.createGuestRole()

}

func (this *SecurityManager) getId() string {
	return "SecurityManager"
}

func (this *SecurityManager) start() {
	log.Println("--> Start SecurityManager")
}

func (this *SecurityManager) stop() {
	log.Println("--> Stop SecurityManager")
}

/////////////////////////////////////////////////////////////////////////////
// Role
/////////////////////////////////////////////////////////////////////////////
/**
 * Create a new Role with a given id.
 */
func (this *SecurityManager) createRole(id string) (*CargoEntities.Role, *CargoEntities.Error) {

	// Create the role with this id if it doesn't exist
	var role *CargoEntities.Role
	entity, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{id})

	if entity == nil {
		role = new(CargoEntities.Role)
		role.SetId(id)

		// Create the role.
		GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntities(), "M_roles", role)
	} else {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ROLE_ID_ALEADY_EXISTS_ERROR, SERVER_ERROR_CODE, errors.New("The role id '"+id+"' is already attibuted to an existing role entity."))
		// Return the uuid of the created error in the err return param.
		return nil, cargoError
	}
	return role, nil
}

/**
 * Retreive a role with a given id.
 */
func (this *SecurityManager) getRole(id string) (*CargoEntities.Role, *CargoEntities.Error) {
	var role *CargoEntities.Role
	roleEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{id})

	if err != nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ROLE_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The role id '"+id+"' does not correspond to an existing role entity."))
		return nil, cargoError
	}

	role = roleEntity.(*CargoEntities.Role)

	return role, nil
}

/**
 * Delete a role with a given id.
 */
func (this *SecurityManager) deleteRole(id string) *CargoEntities.Error {
	roleEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{id})
	if err != nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ROLE_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The role id '"+id+"' does not correspond to an existing role entity."))

		// Return the uuid of the created error in the err return param.
		return cargoError
	}

	GetServer().GetEntityManager().deleteEntity(roleEntity)

	return nil
}

/////////////////////////////////////////////////////////////////////////////
// Account
/////////////////////////////////////////////////////////////////////////////
/**
 * Append a new account to a given role. Do nothing if the account is already in the role
 */
func (this *SecurityManager) appendAccount(roleId string, accountId string) *CargoEntities.Error {
	accountEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", []interface{}{accountId})
	if err != nil {
		return err
	}

	roleEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{roleId})
	if err != nil {
		return err
	}

	// Verify if the role doesn't already have the account
	if !(this.hasAccount(roleId, accountId)) {
		// Get the role entity
		role := roleEntity.(*CargoEntities.Role)
		account := accountEntity.(*CargoEntities.Account)

		// Set the account to the role
		role.AppendAccountsRef(account)
		account.AppendRolesRef(role)

		GetServer().GetEntityManager().saveEntity(roleEntity)
		GetServer().GetEntityManager().saveEntity(accountEntity)
	}

	return nil
}

/**
 * Return true if a role has an account.
 */
func (this *SecurityManager) hasAccount(roleId string, accountId string) bool {

	accountEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", []interface{}{accountId})
	if err != nil {
		return false
	}

	roleEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{roleId})
	if err != nil {
		return false
	}

	role := roleEntity.(*CargoEntities.Role)
	for i := 0; i < len(role.M_accountsRef); i++ {
		if role.M_accountsRef[i] == accountEntity.GetUuid() {
			return true
		}
	}

	return false
}

/**
 * Remove an account from a given role.
 */
func (this *SecurityManager) removeAccount(roleId string, accountId string) *CargoEntities.Error {

	accountEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", []interface{}{accountId})
	if err != nil {
		return err
	}

	roleEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{roleId})
	if err != nil {
		return err
	}

	// Verify if the role has the account
	if this.hasAccount(roleId, accountId) {
		// Get the role entity

		role := roleEntity.(*CargoEntities.Role)
		// Get the account entity

		account := accountEntity.(*CargoEntities.Account)

		// Remove the account from the role
		role.RemoveAccountsRef(account)
		account.RemoveRolesRef(role)

		GetServer().GetEntityManager().saveEntity(roleEntity)
		GetServer().GetEntityManager().saveEntity(accountEntity)

	} else {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ROLE_DOESNT_HAVE_ACCOUNT_ERROR, SERVER_ERROR_CODE, errors.New("The account with the id '"+accountId+"' is not related to the role with the id '"+roleId+"'."))

		// Return the uuid of the created error in the err return param.
		return cargoError
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////
// Action
/////////////////////////////////////////////////////////////////////////////
/**
 * Append a new action to a given role. Do nothing if the action is already in the role
 */
func (this *SecurityManager) appendAction(roleId string, actionName string) *CargoEntities.Error {
	actionEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Action", "CargoEntities", []interface{}{actionName})
	if err != nil {
		return err
	}

	roleEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{roleId})
	if err != nil {
		return err
	}

	// Verify if the role doesn't already have the account
	if !(this.hasAction(roleId, actionName)) {
		role := roleEntity.(*CargoEntities.Role)

		// Set the account to the role
		role.AppendActionsRef(actionEntity.(*CargoEntities.Action))

		// Save the entities.
		GetServer().GetEntityManager().saveEntity(roleEntity)
		GetServer().GetEntityManager().saveEntity(actionEntity)
	}

	return nil
}

/**
 * Return true if a role has an action.
 */
func (this *SecurityManager) hasAction(roleId string, actionName string) bool {

	actionEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Action", "CargoEntities", []interface{}{actionName})
	if err != nil {
		return false
	}

	roleEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{roleId})
	if err != nil {
		return false
	}

	role := roleEntity.(*CargoEntities.Role)
	for i := 0; i < len(role.M_actionsRef); i++ {
		if role.M_actionsRef[i] == actionEntity.GetUuid() {
			return true
		}
	}

	return false
}

/**
 * Remove an action from a given role.
 */
func (this *SecurityManager) removeAction(roleId string, actionName string) *CargoEntities.Error {

	actionEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Action", "CargoEntities", []interface{}{actionName})
	if err != nil {
		return err
	}

	roleEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Role", "CargoEntities", []interface{}{roleId})
	if err != nil {
		return err
	}

	// Verify if the role has the action
	if this.hasAction(roleId, actionName) {
		// Get the role entity

		role := roleEntity.(*CargoEntities.Role)
		// Get the account entity
		action := actionEntity.(*CargoEntities.Action)

		// Remove the account from the role
		role.RemoveActionsRef(action)
		GetServer().GetEntityManager().saveEntity(roleEntity)

	} else {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ROLE_DOESNT_HAVE_ACTION_ERROR, SERVER_ERROR_CODE, errors.New("The action with the name '"+actionName+"' is not related to the role with the id '"+roleId+"'."))

		// Return the uuid of the created error in the err return param.
		return cargoError
	}

	return nil
}

/**
 * Return permission deny if the account releated with the session id can
 * no excute the action.
 */
func (this *SecurityManager) canExecuteAction(sessionId string, actionName string) *CargoEntities.Error {
	// Local action, when no sessionId was given.
	return nil
	/*
		if len(sessionId) == 0 {
			return nil
		}

		// format the action name...
		actionName = actionName[strings.LastIndex(actionName, "/")+1:]
		actionName = strings.Replace(actionName, "*", "", -1)
		actionName = strings.Replace(actionName, ")", "", -1)
		actionName = strings.Replace(actionName, "(", "", -1)

		// Here is the list of function where no permission apply...
		actionEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Action", "CargoEntities", []interface{}{actionName})
		if err != nil {
			return err
		}

		action := actionEntity.(*CargoEntities.Action)

		// if the action is public or hidden no validation is required.
		if action.GetAccessType() == CargoEntities.AccessType_Hidden || action.GetAccessType() == CargoEntities.AccessType_Public {
			return nil
		}

		// Format action name from code.myceliUs.com/CargoWebServer/Cargo/Server.(*WorkflowManager).StartProcess
		// to Server.WorkflowManager.StartProcess
		actionName = strings.Replace(actionName, "code.myceliUs.com/CargoWebServer/Cargo/", "", -1)
		actionName = strings.Replace(actionName, "(", "", -1)
		actionName = strings.Replace(actionName, "*", "", -1)
		actionName = strings.Replace(actionName, ")", "", -1)

		session := GetServer().GetSessionManager().getActiveSessionById(sessionId)
		var account *CargoEntities.Account
		if session != nil {
			account = session.GetAccountPtr()
			roles := account.GetRolesRef()
			for i := 0; i < len(roles); i++ {
				if this.hasAction(roles[i].GetId(), actionName) {
					//log.Println("account ", account.GetId(), " can execute action ", action.GetName())
					return nil
				}
			}
		} else {
			// Create the error message
			cargoError := NewError(Utility.FileLine(), SESSION_UUID_NOT_FOUND_ERROR, SERVER_ERROR_CODE, errors.New("The session with the uuid '"+sessionId+"' dosen't exist!"))
			// Return the uuid of the created error in the err return param.
			return cargoError
		}

		// Create the error message
		cargoError := NewError(Utility.FileLine(), ACTION_EXECUTE_ERROR, SERVER_ERROR_CODE, errors.New("The action with the name '"+actionName+"' cannot be executed by '"+account.GetId()+"'."))

		// Return the uuid of the created error in the err return param.
		return cargoError
	*/
}

/////////////////////////////////////////////////////////////////////////////
// Permission
/////////////////////////////////////////////////////////////////////////////

/**
 * Return the list of entity for a given permission.
 */
func (this *SecurityManager) getEntitiesByPermission(permission *CargoEntities.Permission) []Entity {
	var entities []Entity
	query := permission.GetId()

	// I will retreive the datastore from the query...
	regex := "p([a-z]+)ch"
	match, _ := regexp.MatchString(regex, query)
	log.Println(match)

	//GetServer().GetEntityManager().getEntitiesByType()

	return entities
}

/**
 * Determine if a user session has the permission to Create, Read,
 * Update or Delete a given entity.
 */
func (this *SecurityManager) hasPermission(sessionId string, permissionType int, entity Entity) *CargoEntities.Error {
	session := GetServer().GetSessionManager().getActiveSessionById(sessionId)

	var account *CargoEntities.Account
	if session != nil {
		account = session.GetAccountPtr()
		owner := GetServer().GetEntityManager().getEntityOwner(entity)

		if owner != nil {
			if owner.GetUuid() == account.GetUuid() {
				// Owner has all entity permission.
				return nil
			}
		}

		permissions := account.GetPermissionsRef()
		for i := 0; i < len(permissions); i++ {
			// Here I will evaluate the pattern that is a regular expression.
			log.Println("--------> evaluate ", permissions[i].GetId())
			// If the regex is a success then i will return nil here.

		}
	} else {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), SESSION_UUID_NOT_FOUND_ERROR, SERVER_ERROR_CODE, errors.New("The session with the id '"+sessionId+"' dosen't exist!"))
		// Return the uuid of the created error in the err return param.
		return cargoError
	}

	var permissionName string
	if permissionType == 0 {
		permissionName = "---"
	} else if permissionType == 1 {
		permissionName = "--x"
	} else if permissionType == 2 {
		permissionName = "-w-"
	} else if permissionType == 3 {
		permissionName = "-wx"
	} else if permissionType == 4 {
		permissionName = "r--"
	} else if permissionType == 5 {
		permissionName = "r-x"
	} else if permissionType == 6 {
		permissionName = "rw-"
	} else if permissionType == 7 {
		permissionName = "rwx"
	}

	// Create the error message
	cargoError := NewError(Utility.FileLine(), PERMISSION_DENIED_ERROR, SERVER_ERROR_CODE, errors.New("User with account '"+account.GetId()+"' dosent have the permission to "+permissionName+" entity with uuid '"+entity.GetUuid()+"'"))

	// Return the uuid of the created error in the err return param.
	return cargoError
}

/**
 * Append a new action to a given role. Do nothing if the action is already in the role
 */
func (this *SecurityManager) appendPermission(accountId string, permissionType int, pattern string) *CargoEntities.Error {

	accountEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", []interface{}{accountId})

	if err == nil {
		// Get or create the permission.
		permissionEntity, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Permission", "CargoEntities", []interface{}{pattern})
		var permission *CargoEntities.Permission

		if permissionEntity != nil {
			permission = permissionEntity.(*CargoEntities.Permission)
			permission.SetTypes(permissionType)
			permission.SetId(pattern)
		} else {
			// So here I will create the permission.
			permission = new(CargoEntities.Permission)
			permission.TYPENAME = "CargoEntities.Permission"
			permission.SetId(pattern)
			permission.SetTypes(permissionType)

			// Set the uuid
			entity, errObj := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntities(), "M_permissions", permission)
			if errObj != nil {
				return errObj
			}
			permission = entity.(*CargoEntities.Permission)
		}

		// Append the account ref to the
		permission.AppendAccountsRef(accountEntity.(*CargoEntities.Account))
		accountEntity.(*CargoEntities.Account).AppendPermissionsRef(permission)

		GetServer().GetEntityManager().saveEntity(permission)
		GetServer().GetEntityManager().saveEntity(accountEntity)

	} else {
		// Account error
		cargoError := NewError(Utility.FileLine(), ACCOUNT_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The account id '"+accountId+"' does not correspond to an existing account entity."))

		// Return the uuid of the created error in the err return param.
		return cargoError
	}

	return nil
}

/**
 * Remove a premission with a given pattern from an account with a given id.
 */
func (this *SecurityManager) removePermission(accountId string, pattern string) *CargoEntities.Error {

	accountEntity, err := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", []interface{}{accountId})

	if err == nil {
		account := accountEntity.(*CargoEntities.Account)

		// Get or create the permission.
		permissionEntity, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Permission", "CargoEntities", []interface{}{pattern})
		var permission *CargoEntities.Permission
		if permissionEntity != nil {
			permission = permissionEntity.(*CargoEntities.Permission)
		} else {
			cargoError := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The permission with pattern '"+pattern+"' does not correspond to an existing account entity."))
			// Return the uuid of the created error in the err return param.
			return cargoError
		}

		// Remove the permission
		account.RemovePermissionsRef(permission)
		permission.RemoveAccountsRef(account)

		// Save the account
		GetServer().GetEntityManager().saveEntity(accountEntity)

		// Save or delete the permission.
		if len(permission.M_accountsRef) == 0 {
			// if the permission is not used anymore I will delete it
			GetServer().GetEntityManager().deleteEntity(permission)
		} else {
			GetServer().GetEntityManager().saveEntity(permission)
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
//SecurityManager.prototype.onEvent = function (evt) {
//    EventHub.prototype.onEvent.call(this, evt)
//}
func (this *SecurityManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

////////////// Role //////////////

// @api 1.0
// Create a new role
// @param {string} id The id of the role to create
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Role} The created role
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) CreateRole(id string, messageId string, sessionId string) *CargoEntities.Role {
	role, errObj := this.createRole(id)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}
	return role
}

// @api 1.0
// Retreive a role with a given id.
// @param {string} id The id of the role to retreive
// @return {*CargoEntities.Role} The retrieved role
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) GetRole(id string, messageId string, sessionId string) *CargoEntities.Role {
	role, errObj := this.getRole(id)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}
	return role
}

// @api 1.0
// Delete a role with a given id.
// @param {string} id The id the of role to delete
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) DeleteRole(id string, messageId string, sessionId string) {
	errObj := this.deleteRole(id)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

// @api 1.0
// Determines if a role has a given account.
// @param {string} roleId The id of the role to verify
// @param {string} accountId The id of the account to verify
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) HasAccount(roleId string, accountId string, messageId string, sessionId string) bool {
	return this.hasAccount(roleId, accountId)
}

// @api 1.0
// Append a new account to a given role. Does nothing if the account is already in the role
// @param {string} roleId The id of the role to append the account to
// @param {string} accountId The id of the account to append
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) AppendAccount(roleId string, accountId string, messageId string, sessionId string) {
	errObj := this.appendAccount(roleId, accountId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

// @api 1.0
// Remove an account from a given role.
// @param {string} roleId The id of the role to remove the account from
// @param {string} accountId The id of the account to remove from the role
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) RemoveAccount(roleId string, accountId string, messageId string, sessionId string) {
	errObj := this.removeAccount(roleId, accountId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

/////////// Action /////////////

// @api 1.0
// Determines if a role has a given action.
// @param {string} roleId The id of the role to verify
// @param {string} actionName The name of the action to verify
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) HasAction(roleId string, actionName string, messageId string, sessionId string) bool {
	return this.hasAction(roleId, actionName)
}

// @api 1.0
// Append a new action to a given role. Does nothing if the action is already in the role
// @param {string} roleId The id of the role to append the action to
// @param {string} actionName The name of the action to append
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) AppendAction(roleId string, accountName string, messageId string, sessionId string) {
	errObj := this.appendAction(roleId, accountName)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

// @api 1.0
// Remove an account from a given role.
// @param {string} roleId The id of the role to remove the account from
// @param {string} actionName The name of the action to remove from the role
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) RemoveAction(roleId string, accountName string, messageId string, sessionId string) {
	errObj := this.removeAction(roleId, accountName)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

////////////// Permission //////////////

// @api 1.0
// Append a new permission to a given account. Does nothing if the permission is already in the account
// @param {string} accountId The id of the account to append the permission to
// @param {string} permissionType The type of the permission to append
// @param {string} pattern The pattern of the permission to eveluate.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) AppendPermission(accountId string, permissionType float64, pattern string, messageId string, sessionId string) {
	errObj := this.appendPermission(accountId, int(permissionType), pattern)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

// @api 1.0
// Remove a permission from an account
// @param {string} accountId The id of the account to remove the permission from
// @param {string} permissionPattern The pattern of the permission to remove from the account
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) RemovePermission(accountId string, pattern string, messageId string, sessionId string) {
	errObj := this.removePermission(accountId, pattern)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

// @api 1.0
// Determines if actual the caller of the method can execute a given action.
// @param {string} actionName The name of the action the user want to execute.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) CanExecuteAction(actionName string, sessionId string, messageId string) bool {
	errObj := this.canExecuteAction(sessionId, actionName)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return false
	}

	return true
}

///////////////////////////////// Other stuff //////////////////////////////////

// @api 1.0
// Change the current password for the admin account.
// @param {string} pwd The current password
// @param {string} newPwd The new password
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *SecurityManager) ChangeAdminPassword(pwd string, newPwd string, messageId string, sessionId string) {
	ids := []interface{}{"admin"}
	adminAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "CargoEntities", ids)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}

	// If the actual password dosent match the password.
	account := adminAccountEntity.(*CargoEntities.Account)
	if account.GetPassword() != pwd {
		errObj = NewError(Utility.FileLine(), PERMISSION_DENIED_ERROR, SERVER_ERROR_CODE, errors.New("Wrong password!"))
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}

	// Now I will set the new password.
	account.SetPassword(newPwd)

	// I will save the entity.
	GetServer().GetEntityManager().saveEntity(adminAccountEntity)
}
