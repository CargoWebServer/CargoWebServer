package Server

import (
	"errors"
	"log"
	"strings"

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
	adminRoleUuid := CargoEntitiesRoleExists("adminRole")

	if len(adminRoleUuid) == 0 {
		ids := []interface{}{"admin"}
		adminAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.Account", ids, false)
		if errObj != nil {
			return
		}

		adminAccount := adminAccountEntity.GetObject().(*CargoEntities.Account)

		// Create adminRole
		adminRole, _ := this.createRole("adminRole")
		adminRole.SetAccounts(adminAccount)
		adminRoleEntity, _ := GetServer().GetEntityManager().getEntityByUuid(adminRole.GetUUID(), false)
		adminRoleEntity.SetNeedSave(true)
		adminRoleEntity.SaveEntity()
		adminAccount.SetRolesRef(adminRole)
		adminAccountEntity.SetNeedSave(true)
		adminAccountEntity.SaveEntity()
	}
}

func (this *SecurityManager) createGuestRole() {
	// Create the guest role if it doesn't exist
	guestRoleUuid := CargoEntitiesRoleExists("guestRole")
	if len(guestRoleUuid) == 0 {
		ids := []interface{}{"guest"}
		guestAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.Account", ids, false)
		if errObj != nil {
			return
		}
		guestAccount := guestAccountEntity.GetObject().(*CargoEntities.Account)

		// Create guestRole
		guestRole, _ := this.createRole("guestRole")
		guestRole.SetAccounts(guestAccount)
		guestRoleEntity, _ := GetServer().GetEntityManager().getEntityByUuid(guestRole.GetUUID(), false)
		guestRoleEntity.SetNeedSave(true)
		guestRoleEntity.SaveEntity()

		// Setting guestRole to guest account
		guestAccount.SetRolesRef(guestRole)
		guestAccountEntity.SetNeedSave(true)
		guestAccountEntity.SaveEntity()
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
	uuid := CargoEntitiesRoleExists(id)
	var role *CargoEntities.Role
	cargoEntities := server.GetEntityManager().getCargoEntities()

	if len(uuid) == 0 {
		role = new(CargoEntities.Role)
		role.SetId(id)
		GetServer().GetEntityManager().NewCargoEntitiesRoleEntity(cargoEntities.GetUuid(), "", role)
		// Create the role.
		GetServer().GetEntityManager().createEntity(cargoEntities.GetUuid(), "M_roles", "CargoEntities.Role", id, role)
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
	roleUuid := CargoEntitiesRoleExists(id)
	roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid, false)

	if err != nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ROLE_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The role id '"+id+"' does not correspond to an existing role entity."))
		return nil, cargoError
	}

	role = roleEntity.GetObject().(*CargoEntities.Role)

	return role, nil
}

/**
 * Delete a role with a given id.
 */
func (this *SecurityManager) deleteRole(id string) *CargoEntities.Error {
	uuid := CargoEntitiesRoleExists(id)

	roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(uuid, false)
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

	roleUuid := CargoEntitiesRoleExists(roleId)
	accountUuid := CargoEntitiesAccountExists(accountId)

	// Verify if the role doesn't already have the account
	if !(this.hasAccount(roleId, accountId)) {
		// Get the role entity
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid, false)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			// Get the account entity
			accountEntity, err := GetServer().GetEntityManager().getEntityByUuid(accountUuid, false)
			if err == nil {
				account := accountEntity.GetObject().(*CargoEntities.Account)
				// Set the account to the role
				role.SetAccounts(account)
				account.SetRolesRef(role)
				roleEntity.SetNeedSave(true)
				roleEntity.SaveEntity()
				accountEntity.SaveEntity()
			} else {
				// Create the error message
				cargoError := NewError(Utility.FileLine(), ACCOUNT_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The account id '"+accountId+"' does not correspond to an existing account entity."))

				// Return the uuid of the created error in the err return param.
				return cargoError
			}
		} else {
			// Create the error message
			cargoError := NewError(Utility.FileLine(), ROLE_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The role id '"+roleId+"' does not correspond to an existing role entity."))

			// Return the uuid of the created error in the err return param.
			return cargoError
		}
	}

	return nil
}

/**
 * Return true if a role has an account.
 */
func (this *SecurityManager) hasAccount(roleId string, accountId string) bool {

	roleUuid := CargoEntitiesRoleExists(roleId)
	accountUuid := CargoEntitiesAccountExists(accountId)

	if (len(accountUuid) != 0) && (len(roleUuid) != 0) {
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid, false)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			for i := 0; i < len(role.M_accounts); i++ {
				if role.M_accounts[i] == accountUuid {
					return true
				}
			}
		}
	}

	return false
}

/**
 * Remove an account from a given role.
 */
func (this *SecurityManager) removeAccount(roleId string, accountId string) *CargoEntities.Error {

	roleUuid := CargoEntitiesRoleExists(roleId)
	accountUuid := CargoEntitiesAccountExists(accountId)

	// Verify if the role has the account
	if this.hasAccount(roleId, accountId) {
		// Get the role entity
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid, false)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			// Get the account entity
			accountEntity, err := GetServer().GetEntityManager().getEntityByUuid(accountUuid, false)
			if err == nil {
				account := accountEntity.GetObject().(*CargoEntities.Account)
				// Remove the account from the role
				role.RemoveAccounts(account)
				roleEntity.SaveEntity()
				account.RemoveRolesRef(role)
				accountEntity.SaveEntity()
			}
		}
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
	roleUuid := CargoEntitiesRoleExists(roleId)
	actionUuid := CargoEntitiesActionExists(actionName)

	// Verify if the role doesn't already have the account
	if !(this.hasAction(roleId, actionName)) {
		// Get the role entity
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid, false)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			// Get the account entity
			actionEntity, err := GetServer().GetEntityManager().getEntityByUuid(actionUuid, false)
			if err == nil {
				action := actionEntity.GetObject().(*CargoEntities.Action)
				// Set the account to the role
				role.SetActions(action)

				// Save the entities.
				roleEntity.SaveEntity()
				actionEntity.SaveEntity()

			} else {
				// Create the error message
				cargoError := NewError(Utility.FileLine(), ACTION_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The action name '"+actionName+"' does not correspond to an existing action entity."))

				// Return the uuid of the created error in the err return param.
				return cargoError
			}
		} else {
			// Create the error message
			cargoError := NewError(Utility.FileLine(), ROLE_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The role id '"+roleId+"' does not correspond to an existing role entity."))

			// Return the uuid of the created error in the err return param.
			return cargoError
		}
	}

	return nil
}

/**
 * Return true if a role has an action.
 */
func (this *SecurityManager) hasAction(roleId string, actionName string) bool {

	roleUuid := CargoEntitiesRoleExists(roleId)
	actionUuid := CargoEntitiesActionExists(actionName)

	if (len(actionUuid) != 0) && (len(roleUuid) != 0) {
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid, false)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			for i := 0; i < len(role.M_actions); i++ {
				if role.M_actions[i] == actionUuid {
					return true
				}
			}
		}
	}

	return false
}

/**
 * Remove an action from a given role.
 */
func (this *SecurityManager) removeAction(roleId string, actionName string) *CargoEntities.Error {

	roleUuid := CargoEntitiesRoleExists(roleId)

	// Verify if the role has the action
	if this.hasAction(roleId, actionName) {
		// Get the role entity
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid, false)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			// Get the account entity
			actionUuid := CargoEntitiesActionExists(actionName)
			actionEntity, err := GetServer().GetEntityManager().getEntityByUuid(actionUuid, false)
			if err == nil {
				action := actionEntity.GetObject().(*CargoEntities.Action)
				// Remove the account from the role
				role.RemoveActions(action)
				roleEntity.SaveEntity()
			} else {
				cargoError := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The action with name '"+actionName+"' dosent exist."))
				return cargoError
			}
		} else {
			// If the role does not exist.
			cargoError := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The role with the id '"+roleId+"' dosent exist."))
			return cargoError
		}
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
	if len(sessionId) == 0 {
		return nil
	}

	// format the action name...
	actionName = actionName[strings.LastIndex(actionName, "/")+1:]
	actionName = strings.Replace(actionName, "*", "", -1)
	actionName = strings.Replace(actionName, ")", "", -1)
	actionName = strings.Replace(actionName, "(", "", -1)

	// Here is the list of function where no permission apply...
	actionUuid := CargoEntitiesActionExists(actionName)
	if len(actionUuid) == 0 {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The action with the name '"+actionName+"' dosen't exist!"))
		// Return the uuid of the created error in the err return param.
		return cargoError
	}

	actionEntity, _ := GetServer().GetEntityManager().getEntityByUuid(actionUuid, true)
	action := actionEntity.GetObject().(*CargoEntities.Action)

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
				log.Println("account ", account.GetId(), " can execute action ", action.GetName())
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
			if owner.GetUUID() == account.GetUUID() {
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
	accountUuid := CargoEntitiesAccountExists(accountId)
	accountEntity, err := GetServer().GetEntityManager().getEntityByUuid(accountUuid, false)

	if err == nil {
		// Get or create the permission.
		permissionUuid := CargoEntitiesPermissionExists(pattern)
		var permissionEntity *CargoEntities_PermissionEntity
		var permission *CargoEntities.Permission

		if len(permissionUuid) > 0 {
			entity, errObj := GetServer().GetEntityManager().getEntityByUuid(permissionUuid, false)
			if errObj != nil {
				return errObj
			}
			permissionEntity = entity.(*CargoEntities_PermissionEntity)
			permission = permissionEntity.GetObject().(*CargoEntities.Permission)
			permission.SetTypes(permissionType)
			permission.SetId(pattern)
		} else {
			// So here I will create the permission.
			permission = new(CargoEntities.Permission)
			permission.TYPENAME = "CargoEntities.Permission"
			permission.SetId(pattern)
			permission.SetTypes(permissionType)
			// Set the uuid
			GetServer().GetEntityManager().NewCargoEntitiesPermissionEntity(GetServer().GetEntityManager().getCargoEntities().GetUuid(), "", permission)
			var errObj *CargoEntities.Error
			entity, errObj := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntities().GetUuid(), "M_permissions", "CargoEntities.Permission", "", permission)
			if errObj != nil {
				return errObj
			}
			permissionEntity = entity.(*CargoEntities_PermissionEntity)
			permission = permissionEntity.GetObject().(*CargoEntities.Permission)
		}

		// Append the account ref to the
		permission.SetAccountsRef(accountEntity.GetObject().(*CargoEntities.Account))
		accountEntity.GetObject().(*CargoEntities.Account).SetPermissionsRef(permission)

		//accountEntity.SaveEntity()
		//permissionEntity.SaveEntity()

		// Save the entity...
		GetServer().GetEntityManager().getCargoEntities().SaveEntity()

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

	accountUuid := CargoEntitiesAccountExists(accountId)
	accountEntity, err := GetServer().GetEntityManager().getEntityByUuid(accountUuid, false)

	if err == nil {
		account := accountEntity.GetObject().(*CargoEntities.Account)

		// Get or create the permission.
		permissionUuid := CargoEntitiesPermissionExists(pattern)
		var permissionEntity *CargoEntities_PermissionEntity
		var permission *CargoEntities.Permission
		if len(permissionUuid) > 0 {
			entity, errObj := GetServer().GetEntityManager().getEntityByUuid(permissionUuid, false)
			if errObj != nil {
				return errObj
			}
			permissionEntity = entity.(*CargoEntities_PermissionEntity)
		} else {
			cargoError := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The permission with pattern '"+pattern+"' does not correspond to an existing account entity."))
			// Return the uuid of the created error in the err return param.
			return cargoError
		}

		// Remove the permission
		permission = permissionEntity.GetObject().(*CargoEntities.Permission)
		account.RemovePermissionsRef(permission)
		permission.RemoveAccountsRef(account)

		// Save the account
		accountEntity.SaveEntity()

		// Save or delete the permission.
		if len(permission.M_accountsRef) == 0 {
			// if the permission is not used anymore I will delete it
			permissionEntity.DeleteEntity()
		} else {
			permissionEntity.SaveEntity()
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
func (this *SecurityManager) AppendPermission(accountId string, permissionType int, pattern string, messageId string, sessionId string) {
	errObj := this.appendPermission(accountId, permissionType, pattern)
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
	adminAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.Account", ids, false)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}

	// If the actual password dosent match the password.
	account := adminAccountEntity.GetObject().(*CargoEntities.Account)
	if account.GetPassword() != pwd {
		errObj = NewError(Utility.FileLine(), PERMISSION_DENIED_ERROR, SERVER_ERROR_CODE, errors.New("Wrong password!"))
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}

	// Now I will set the new password.
	account.SetPassword(newPwd)

	// I will save the entity.
	adminAccountEntity.SaveEntity()
}
