package Server

import (
	"errors"
	"log"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

/**
 * Entity security functionality.
 */
type SecurityManager struct {
	adminRole *CargoEntities.Role
	guestRole *CargoEntities.Role
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

/**
 * Initialize security related information.
 */
func (this *SecurityManager) initialize() {

	log.Println("--> Initialize SessionManager")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())

	// Create the admin role if it doesn't exist
	adminRoleUuid := CargoEntitiesRoleExists("adminRole")
	cargoEntities := server.GetEntityManager().getCargoEntities()

	if len(adminRoleUuid) == 0 {

		adminAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.Account", "admin")
		if errObj != nil {
			return
		}
		adminAccount := adminAccountEntity.GetObject().(*CargoEntities.Account)

		// Create adminRole
		this.adminRole, _ = this.createRole("adminRole")
		this.adminRole.SetAccounts(adminAccount)

		//adminAccountEntity.SaveEntity()
		adminAccount.SetRolesRef(this.adminRole)

		// Setting adminRole to admin account
		cargoEntities.GetObject().(*CargoEntities.Entities).SetRoles(this.adminRole)
		cargoEntities.SaveEntity()

	}

	// Create the guest role if it doesn't exist
	guestRoleUuid := CargoEntitiesRoleExists("guestRole")
	if len(guestRoleUuid) == 0 {

		guestAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.Account", "guest")
		if errObj != nil {
			return
		}
		guestAccount := guestAccountEntity.GetObject().(*CargoEntities.Account)

		// Create guestRole
		this.guestRole, _ = this.createRole("guestRole")
		this.guestRole.SetAccounts(guestAccount)

		// Setting guestRole to guest account
		cargoEntities.GetObject().(*CargoEntities.Entities).SetRoles(this.guestRole)
		guestAccount.SetRolesRef(this.guestRole)
		//guestAccountEntity.SaveEntity()

		cargoEntities.SaveEntity()
	}

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
		roleEntity := GetServer().GetEntityManager().NewCargoEntitiesRoleEntity(cargoEntities.GetUuid(), id, nil)
		role = roleEntity.GetObject().(*CargoEntities.Role)
		role.SetId(id)
		cargoEntities.GetObject().(*CargoEntities.Entities).SetRoles(role)
		cargoEntities.SaveEntity()
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
	roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid)

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

	roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(uuid)
	if err != nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ROLE_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The role id '"+id+"' does not correspond to an existing role entity."))

		// Return the uuid of the created error in the err return param.
		return cargoError
	}

	// Remove the role from all accounts that have this role
	accounts := roleEntity.GetObject().(*CargoEntities.Role).GetAccounts()
	for i := 0; i < len(accounts); i++ {
		accounts[i].RemoveRolesRef(roleEntity.GetObject())
		accountEntity := GetServer().GetEntityManager().NewCargoEntitiesAccountEntityFromObject(accounts[i])
		accountEntity.SaveEntity()
	}

	roleEntity.DeleteEntity()

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
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			// Get the account entity
			accountEntity, err := GetServer().GetEntityManager().getEntityByUuid(accountUuid)
			if err == nil {
				account := accountEntity.GetObject().(*CargoEntities.Account)
				// Set the account to the role
				role.SetAccounts(account)
				account.SetRolesRef(role)

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
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			for i := 0; i < len(role.M_accounts); i++ {
				if role.M_accounts[i] == accountId {
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
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			// Get the account entity
			accountEntity, err := GetServer().GetEntityManager().getEntityByUuid(accountUuid)
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
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			// Get the account entity
			actionEntity, err := GetServer().GetEntityManager().getEntityByUuid(actionUuid)
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
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid)
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
		roleEntity, err := GetServer().GetEntityManager().getEntityByUuid(roleUuid)
		if err == nil {
			role := roleEntity.GetObject().(*CargoEntities.Role)
			// Get the account entity
			actionEntity, err := GetServer().GetEntityManager().getEntityByUuid(actionName)
			if err == nil {
				action := actionEntity.GetObject().(*CargoEntities.Action)
				// Remove the account from the role
				role.RemoveActions(action)
				roleEntity.SaveEntity()
			}
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

	// Format action name from code.myceliUs.com/CargoWebServer/Cargo/Server.(*WorkflowManager).StartProcess
	// to Server.WorkflowManager.StartProcess
	actionName = strings.Replace(actionName, "code.myceliUs.com/CargoWebServer/Cargo/", "", -1)
	actionName = strings.Replace(actionName, "(", "", -1)
	actionName = strings.Replace(actionName, "*", "", -1)
	actionName = strings.Replace(actionName, ")", "", -1)

	session := GetServer().GetSessionManager().activeSessions[sessionId]

	var account *CargoEntities.Account
	if session != nil {
		account = session.GetAccountPtr()
		roles := account.GetRolesRef()
		for i := 0; i < len(roles); i++ {
			if this.hasAction(roles[i].GetId(), actionName) {
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
 * Determine if a user session has the permission to Create, Read,
 * Update or Delete a given entity.
 */
func (this *SecurityManager) hasPermission(sessionId string, permissionType CargoEntities.PermissionType, entity Entity) *CargoEntities.Error {
	session := GetServer().GetSessionManager().activeSessions[sessionId]

	var account *CargoEntities.Account
	if session != nil {
		account = session.GetAccountPtr()
		permissions := account.GetPermissions()
		for i := 0; i < len(permissions); i++ {
			// Here I will evaluate the pattern that is a regular expression.
			log.Println("--------> evaluate ", permissions[i].GetPattern())
			// If the regex is a success then i will return nil here.
		}
	} else {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), SESSION_UUID_NOT_FOUND_ERROR, SERVER_ERROR_CODE, errors.New("The session with the id '"+sessionId+"' dosen't exist!"))
		// Return the uuid of the created error in the err return param.
		return cargoError
	}

	var permissionName string
	if permissionType == CargoEntities.PermissionType_Create {
		permissionName = "create"
	} else if permissionType == CargoEntities.PermissionType_Read {
		permissionName = "read"
	} else if permissionType == CargoEntities.PermissionType_Update {
		permissionName = "update"
	} else if permissionType == CargoEntities.PermissionType_Delete {
		permissionName = "delete"
	}

	// Create the error message
	cargoError := NewError(Utility.FileLine(), PERMISSION_DENIED_ERROR, SERVER_ERROR_CODE, errors.New("User with account '"+account.GetId()+"' dosent have the permission to "+permissionName+" entity with uuid '"+entity.GetUuid()+"'"))

	// Return the uuid of the created error in the err return param.
	return cargoError
}

/**
 * Append a new action to a given role. Do nothing if the action is already in the role
 */
func (this *SecurityManager) appendPermission(accountId string, permissionType CargoEntities.PermissionType, pattern string) *CargoEntities.Error {
	accountUuid := CargoEntitiesAccountExists(accountId)
	accountEntity, err := GetServer().GetEntityManager().getEntityByUuid(accountUuid)
	if err == nil {
		// Here I will get the account and set the permission if is not already
		// exist.
		account := accountEntity.GetObject().(*CargoEntities.Account)
		for i := 0; i < len(account.GetPermissions()); i++ {
			permission := account.GetPermissions()[i]
			if permission.GetPattern() == pattern && permission.GetType() == permissionType {
				return nil // not a error but nothing to do here.
			}
		}
		// So here I will create the permission.
		permission := new(CargoEntities.Permission)
		permission.UUID = "CargoEntities.Permission%" + Utility.RandomUUID()
		permission.TYPENAME = "CargoEntities.Permission"
		permission.SetPattern(pattern)
		permission.SetType(permissionType)
		account.SetPermissions(permission)
		accountEntity.SaveEntity()

	} else {
		// Account error
		cargoError := NewError(Utility.FileLine(), ACCOUNT_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The account id '"+accountId+"' does not correspond to an existing account entity."))

		// Return the uuid of the created error in the err return param.
		return cargoError
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

/**
 * Create a new Role with a given id.
 */
func (this *SecurityManager) CreateRole(id string, messageId string, sessionId string) interface{} {
	role, errObj := this.createRole(id)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}
	return role
}

/**
 * Retreive a role with a given id.
 */
func (this *SecurityManager) GetRole(id string, messageId string, sessionId string) interface{} {
	role, errObj := this.getRole(id)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}
	return role
}

/**
 * Delete a role with a given id.
 */
func (this *SecurityManager) DeleteRole(id string, messageId string, sessionId string) {
	errObj := this.deleteRole(id)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

/**
 * Return true if a role has an account.
 */
func (this *SecurityManager) HasAccount(roleId string, accountId string, messageId string, sessionId string) bool {
	return this.hasAccount(roleId, accountId)
}

/**
 * Append a new account to a given role. Do nothing if the account is already in the role
 */
func (this *SecurityManager) AppendAccount(roleId string, accountId string, messageId string, sessionId string) {
	errObj := this.appendAccount(roleId, accountId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

/**
 * Remove an account from a given role.
 */
func (this *SecurityManager) RemoveAccount(roleId string, accountId string, messageId string, sessionId string) {
	errObj := this.removeAccount(roleId, accountId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

/**
 * Change the admin password.
 */
func (this *SecurityManager) ChangeAdminPassword(pwd string, newPwd string, messageId string, sessionId string) {

	adminAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities", "CargoEntities.Account", "admin")
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
