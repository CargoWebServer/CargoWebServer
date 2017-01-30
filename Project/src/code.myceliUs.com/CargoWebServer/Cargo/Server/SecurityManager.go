package Server

import (
	"errors"

	"log"
	"reflect"

	"code.myceliUs.com/CargoWebServer/Cargo/Persistence/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
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

/**
 * Initialize security related information.
 */
func (this *SecurityManager) Initialize() *CargoEntities.Error {

	// Create the admin role if it doesn't exist
	adminRoleUuid := CargoEntitiesRoleExists("adminRole")
	cargoEntities := server.GetEntityManager().getCargoEntities()

	if len(adminRoleUuid) == 0 {

		adminAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "admin")
		if errObj != nil {
			return errObj
		}
		adminAccount := adminAccountEntity.GetObject().(*CargoEntities.Account)

		// Create adminRole
		this.adminRole, _ = this.createRole("adminRole")
		this.adminRole.SetAccountsRef(adminAccount)

		//adminAccountEntity.SaveEntity()
		adminAccount.SetRolesRef(this.adminRole)

		// Setting adminRole to admin account
		cargoEntities.GetObject().(*CargoEntities.Entities).SetRoles(this.adminRole)
		cargoEntities.SaveEntity()

	}

	// Create the guest role if it doesn't exist
	guestRoleUuid := CargoEntitiesRoleExists("guestRole")
	if len(guestRoleUuid) == 0 {

		guestAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "guest")
		if errObj != nil {
			return errObj
		}
		guestAccount := guestAccountEntity.GetObject().(*CargoEntities.Account)

		// Create guestRole
		this.guestRole, _ = this.createRole("guestRole")
		this.guestRole.SetAccountsRef(guestAccount)

		// Setting guestRole to guest account
		cargoEntities.GetObject().(*CargoEntities.Entities).SetRoles(this.guestRole)
		guestAccount.SetRolesRef(this.guestRole)
		//guestAccountEntity.SaveEntity()

		cargoEntities.SaveEntity()
	}

	return nil
}

func (this *SecurityManager) GetId() string {
	return "SecurityManager"
}

func (this *SecurityManager) Start() {
	log.Println("--> Start SecurityManager")
}

func (this *SecurityManager) Stop() {
	log.Println("--> Stop SecurityManager")
}

/**
 * Create a new Role with a given id.
 */
func (this *SecurityManager) createRole(id string) (*CargoEntities.Role, *CargoEntities.Error) {

	// Create the role with this id if it doesn't exist
	uuid := CargoEntitiesRoleExists(id)
	var role *CargoEntities.Role
	cargoEntities := server.GetEntityManager().getCargoEntities()

	if len(uuid) == 0 {
		roleEntity := GetServer().GetEntityManager().NewCargoEntitiesRoleEntity(id, nil)
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
				role.SetAccountsRef(account)
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
			for i := 0; i < len(role.M_accountsRef); i++ {
				if role.M_accountsRef[i] == accountId {
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
				role.RemoveAccountsRef(account)
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
	accountsRef := roleEntity.GetObject().(*CargoEntities.Role).GetAccountsRef()
	for i := 0; i < len(accountsRef); i++ {
		accountsRef[i].RemoveRolesRef(roleEntity.GetObject())
		accountEntityRef := GetServer().GetEntityManager().NewCargoEntitiesAccountEntityFromObject(accountsRef[i])
		accountEntityRef.SaveEntity()
	}

	roleEntity.DeleteEntity()

	return nil
}

/**
 * Set a restriction with a given role (by id) to an action.
 */
func (this *SecurityManager) setRestrictionRole(roleId string, action *CargoEntities.Action) *CargoEntities.Error {

	roleUuid := CargoEntitiesRoleExists(roleId)
	// Get the role uuid, return error if it doesn't exist
	roleEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(roleUuid)
	if errObj != nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ROLE_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The role id '"+roleId+"' does not correspond to an existing role entity."))
		return cargoError
	}

	// Find all the restrictions and verify if one exists for the action
	restrictions, _ := GetServer().GetEntityManager().getEntitiesByType("CargoEntities.Restriction", "", "CargoEntities")

	for i := 0; i < len(restrictions); i++ {
		if restrictions[i].GetObject().(*CargoEntities.Restriction).GetActionRef().GetUUID() == action.GetUUID() {
			restrictions[i].GetObject().(*CargoEntities.Restriction).SetRolesRef(roleEntity.GetObject())
			return nil
		}
	}

	// The restriction on the role for the action doesn't already exist

	var restrictionEntity *CargoEntities_RestrictionEntity
	restrictionEntity = GetServer().GetEntityManager().NewCargoEntitiesRestrictionEntity("", nil)
	restriction := restrictionEntity.GetObject().(*CargoEntities.Restriction)
	restriction.SetActionRef(action)
	restriction.SetRolesRef(roleEntity.GetObject())

	// Save the information.
	cargoEntities := server.GetEntityManager().getCargoEntities()
	cargoEntities.GetObject().(*CargoEntities.Entities).SetRestrictions(restriction)
	cargoEntities.SaveEntity()

	return nil
}

/**
 * Remove a role (by Id) from the restriction associated to an action
 */
func (this *SecurityManager) removeRestrictionRole(roleId string, action *CargoEntities.Action) *CargoEntities.Error {

	roleUuid := CargoEntitiesRoleExists(roleId)
	// Get the role entity, return error if it doesn't exist
	roleEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(roleUuid)
	if errObj != nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ROLE_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The role id '"+roleId+"' does not correspond to an existing role entity."))
		return cargoError
	}

	// Find all the restrictions and verify if one exists for the action
	restrictions, _ := GetServer().GetEntityManager().getEntitiesByType("CargoEntities.Restriction", "", "CargoEntities")

	for i := 0; i < len(restrictions); i++ {
		if restrictions[i].GetObject().(*CargoEntities.Restriction).GetActionRef().GetUUID() == action.GetUUID() {
			restrictions[i].GetObject().(*CargoEntities.Restriction).RemoveRolesRef(roleEntity.GetObject())
			return nil
		}
	}

	cargoError := NewError(Utility.FileLine(), RESTRICTION_ACTION_ROLE_ERROR, SERVER_ERROR_CODE, errors.New("There is no restriction for the role '"+roleId+"' on the action '"+action.GetName()+"'."))
	return cargoError

}

/**
 * Verify if a role has a restriction for a given action
 */
func (this *SecurityManager) roleHasRestrictionForAction(roleId string, action *CargoEntities.Action) (bool, *CargoEntities.Error) {

	// Get the role uuid, return error if it doesn't exist
	roleUuid := CargoEntitiesRoleExists(roleId)
	if len(roleUuid) == 0 {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), ROLE_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The role id '"+roleId+"' does not correspond to an existing role entity."))
		return false, cargoError
	}

	// Find all the restrictions and verify if one exists for the action
	restrictions, _ := GetServer().GetEntityManager().getEntitiesByType("CargoEntities.Restriction", "", "CargoEntities")

	for i := 0; i < len(restrictions); i++ {
		if restrictions[i].GetObject().(*CargoEntities.Restriction).GetActionRef().GetUUID() == action.GetUUID() {

			// Verify if the resriction on the action for the role exists

			for j := 0; j < len(restrictions[i].GetObject().(*CargoEntities.Restriction).GetRolesRef()); j++ {
				if restrictions[i].GetObject().(*CargoEntities.Restriction).GetRolesRef()[j].GetId() == roleId {
					return true, nil

				}
			}
		}
	}
	return false, nil

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
 * Set a restriction with a given role (by id) to an action.
 */
func (this *SecurityManager) SetRestrictionRole(roleId string, action interface{}, messageId string, sessionId string) {

	if reflect.TypeOf(action).String() != "*CargoEntities.Action" {
		cargoError := NewError(Utility.FileLine(), PARAMETER_TYPE_ERROR, SERVER_ERROR_CODE, errors.New("Expected '*CargoEntities.Action' but got '"+reflect.TypeOf(action).String()+"' instead."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	errObj := this.setRestrictionRole(roleId, action.(*CargoEntities.Action))
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

/**
 * Remove a role (by Id) from the restriction associated to an action
 */
func (this *SecurityManager) RemoveRestrictionRole(roleId string, action interface{}, messageId string, sessionId string) {

	if reflect.TypeOf(action).String() != "*CargoEntities.Action" {
		cargoError := NewError(Utility.FileLine(), PARAMETER_TYPE_ERROR, SERVER_ERROR_CODE, errors.New("Expected '*CargoEntities.Action' but got '"+reflect.TypeOf(action).String()+"' instead."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	errObj := this.removeRestrictionRole(roleId, action.(*CargoEntities.Action))
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

/**
 * Verify if a role has a restriction for a given action
 */
func (this *SecurityManager) RoleHasRestrictionForAction(roleId string, action interface{}, messageId string, sessionId string) bool {

	if reflect.TypeOf(action).String() != "*CargoEntities.Action" {
		cargoError := NewError(Utility.FileLine(), PARAMETER_TYPE_ERROR, SERVER_ERROR_CODE, errors.New("Expected '*CargoEntities.Action' but got '"+reflect.TypeOf(action).String()+"' instead."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return false
	}

	result, errObj := this.roleHasRestrictionForAction(roleId, action.(*CargoEntities.Action))
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
	return result
}

/*
* Validate the restriction for an action
 */
func validateRestriction(actionId string, sessionId string) *CargoEntities.Error {

	actionUuid := CargoEntitiesActionExists(actionId)
	var action *CargoEntities.Action
	if len(actionUuid) == 0 {
		// create the action
		actionEntity := GetServer().GetEntityManager().NewCargoEntitiesActionEntity(actionId, nil)
		action = actionEntity.GetObject().(*CargoEntities.Action)
		action.SetId(actionId)
		cargoEntities := server.GetEntityManager().getCargoEntities()
		cargoEntities.GetObject().(*CargoEntities.Entities).SetEntities(action)
		cargoEntities.SaveEntity()

		return nil
	}

	actionEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(actionUuid)
	action = actionEntity.GetObject().(*CargoEntities.Action)
	if errObj != nil {
		return errObj
	}

	// get the account for the given session
	activeSessionEntity := GetServer().GetSessionManager().GetActiveSessionById(sessionId)

	if activeSessionEntity == nil {
		return nil
	}

	accountPtr := activeSessionEntity.GetAccountPtr()
	if accountPtr == nil {
		return nil
	}

	// get the roles for the account
	roles := accountPtr.GetRolesRef()

	// verify if role has restriction for action
	for i := 0; i < len(roles); i++ {
		roleHasRestrictionForAction, _ := GetServer().GetSecurityManager().roleHasRestrictionForAction(roles[i].GetId(), action)
		if roleHasRestrictionForAction {
			return NewError(Utility.FileLine(), PERMISSION_DENIED_ERROR, SERVER_ERROR_CODE, errors.New("There is a restriction on the action '"+actionId+"' with the role '"+roles[i].GetId()+"' for the account '"+accountPtr.GetId()+"'."))
		}
	}

	return nil
}
