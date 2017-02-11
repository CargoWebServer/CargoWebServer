package Server

import (
	"errors"

	"log"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
)

/**
 * Entity security functionality.
 */
type SecurityManager struct {
	adminRole *CargoEntities.Role
	guestRole *CargoEntities.Role
	m_config  *Config.ServiceConfiguration
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
	this.m_config = GetServer().GetConfigurationManager().getServiceConfiguration(this.getId())

	// Create the admin role if it doesn't exist
	adminRoleUuid := CargoEntitiesRoleExists("adminRole")
	cargoEntities := server.GetEntityManager().getCargoEntities()

	if len(adminRoleUuid) == 0 {

		adminAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "admin")
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

		guestAccountEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Account", "guest")
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

func (this *SecurityManager) getConfig() *Config.ServiceConfiguration {
	return this.m_config
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
