package Server

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"code.myceliUs.com/Utility"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	LDAP "github.com/mavricknz/ldap"
)

type LdapManager struct {
	// Nothing here.
}

var ldapManager *LdapManager

func (this *Server) GetLdapManager() *LdapManager {
	if ldapManager == nil {
		ldapManager = newLdapManager()
	}
	return ldapManager
}

func newLdapManager() *LdapManager {
	ldapManager := new(LdapManager)
	return ldapManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * That function is use to synchronize the information of a ldap server
 * with a given id.
 */
func (this *LdapManager) initialize() {
	// register service avalaible action here.
	log.Println("--> initialyze LdapManager")
	// Create the default configurations
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)
}

func (this *LdapManager) getId() string {
	return "LdapManager"
}

func (this *LdapManager) getConfigsInfo() map[string]*Config.LdapConfiguration {
	configsInfo := make(map[string]*Config.LdapConfiguration, 0)

	activeConfigurations := GetServer().GetConfigurationManager().getActiveConfigurations()
	ldapConfigurations := activeConfigurations.GetLdapConfigs()

	for i := 0; i < len(ldapConfigurations); i++ {
		configsInfo[ldapConfigurations[i].M_id] = ldapConfigurations[i]
	}
	return configsInfo
}

func (this *LdapManager) start() {
	log.Println("--> Start LdapManager")
}

func (this *LdapManager) synchronizeAll() {

	// configure all information from the servers...
	for _, info := range this.getConfigsInfo() {
		// Synchronize the list of user...
		err := this.synchronizeUsers(info.M_id)
		if err != nil {
			log.Println("Synchronize Users Error ", err)
		}

		// Synchronize the list of group...
		err = this.synchronizeGroups(info.M_id)
		if err != nil {
			log.Println("Synchronize Groups Error ", err)
		}

		// Synchronize the list of computer...
		err = this.synchronizeComputers(info.M_id)
		if err != nil {
			log.Println("Synchronize Computers Error ", err)
		}
	}
}

func (this *LdapManager) stop() {
	log.Println("--> Stop LdapManager")
}

/**
 * Get the list of group for a user...
 */
func (this *LdapManager) getLdapUserMemberOf(id string, userId string, base_dn string) ([]string, error) {

	var filter string = "(&(objectClass=group)(objectcategory=Group)(member=" + userId + "))"
	var attributes []string = []string{"sAMAccountName"}
	results, err := this.search(id, "", "", base_dn, filter, attributes)
	var memberOf []string
	if err != nil {
		log.Println("error, fail to search the groups information for user on ldap...")
		return memberOf, err
	}

	// Here I will print the information...
	for i := 0; i < len(results); i++ {
		// Here I will get the user in the group...
		row := results[i]
		for j := 0; j < len(row); j++ {
			// Print the result...
			groupName := row[j].(string)
			// Save user from ldap directory to the database...
			if len(groupName) > 0 {
				memberOf = append(memberOf, groupName)
			}
		}
	}
	return memberOf, nil
}

func (this *LdapManager) getLdapGroupMembers(id string, groupId string) ([]string, error) {
	var base_dn = this.getConfigsInfo()[id].M_searchBase
	var filter string = "(&(objectClass=user)(objectcategory=Person)(memberOf=" + groupId + "))"
	var attributes []string = []string{"sAMAccountName"}
	results, err := this.search(id, "", "", base_dn, filter, attributes)
	var members []string

	if err != nil {
		log.Println("error, fail to search the members information for group on ldap...")
		return members, err
	}

	// Here I will print the information...
	for i := 0; i < len(results); i++ {
		// Here I will get the user in the group...
		row := results[i]
		for j := 0; j < len(row); j++ {
			// Print the result...

			if attributes[j] == "sAMAccountName" {
				userId := row[j].(string)
				// Save user from ldap directory to the database...
				members = append(members, userId)
			}
		}
	}
	return members, nil
}

/**
 * Connect to a ldap server...
 */
func (this *LdapManager) connect(id string, userId string, psswd string) (*LDAP.LDAPConnection, error) {

	ldapConfigInfo := this.getConfigsInfo()[id]

	conn := LDAP.NewLDAPConnection(ldapConfigInfo.M_ipv4, uint16(ldapConfigInfo.M_port))
	conn.NetworkConnectTimeout = time.Duration(3 * time.Second)
	conn.AbandonMessageOnReadTimeout = true
	err := conn.Connect()

	if err != nil {
		// handle error
		log.Println("---> Cannot open the connection: ", ldapConfigInfo.M_hostName, err)
		return nil, err
	}

	// Connect with the default user...
	if len(userId) > 0 {
		conn.Bind(userId, psswd)
	} else {
		conn.Bind(ldapConfigInfo.M_user, ldapConfigInfo.M_pwd)
	}
	return conn, nil
}

/**
 * Search for a list of value over the ldap server. if the base_dn is
 * not specify the default base is use. It return a list of values. This can
 * be interpret as a tow dimensional array.
 */
func (this *LdapManager) search(id string, login string, psswd string, base_dn string, filter string, attributes []string) ([][]interface{}, error) {
	ldapConfigInfo := this.getConfigsInfo()[id]
	// Try to connect to the ldap server...
	var err error
	var conn *LDAP.LDAPConnection
	if len(login) == 0 {
		conn, err = this.connect(ldapConfigInfo.M_id, ldapConfigInfo.M_user, ldapConfigInfo.M_pwd)
		if err != nil {
			log.Println("Fail to connect:", ldapConfigInfo.M_id)
			return nil, err
		}
	} else {
		conn, err = this.connect(ldapConfigInfo.M_id, login, psswd)
		if err != nil {
			log.Println("Fail to connect:", ldapConfigInfo.M_id)
			return nil, err
		}
	}

	// The connection will be close latter...
	defer conn.Close() // Close the connection after the request.

	if len(base_dn) == 0 {
		// Set to default base here...
		base_dn = ldapConfigInfo.M_searchBase
	}

	//Now I will execute the query...
	search_request := LDAP.NewSearchRequest(
		base_dn,
		LDAP.ScopeWholeSubtree, LDAP.NeverDerefAliases, 0, 0, false,
		filter,
		attributes,
		nil)

	sr, err := conn.Search(search_request)
	if err != nil {
		log.Println("--> search error", err)
		return nil, err
	}

	// Store the founded values in results...
	var results [][]interface{}
	for i := 0; i < len(sr.Entries); i++ {
		entry := sr.Entries[i]
		var row []interface{}
		for j := 0; j < len(attributes); j++ {
			attributeName := attributes[j]
			attributeValue := entry.GetAttributeValue(attributeName)
			row = append(row, attributeValue)
		}
		results = append(results, row)
	}

	return results, nil
}

func (this *LdapManager) authenticate(id string, login string, psswd string) bool {

	// Now I will try to make a simple query if it fail that's mean the user
	// does have the permission...
	var base_dn string = "OU=Users," + this.getConfigsInfo()[id].M_searchBase
	var filter string = "(objectClass=user)"

	// Test get some user...
	var attributes []string = []string{"sAMAccountName"}
	_, err := this.search(id, login, psswd, base_dn, filter, attributes)
	if err != nil {
		log.Println("---> ldap authenticate fail: ", login, err)
		return false
	}

	return true
}

////////////////////////////////////////////////////////////////////////////////
//	User
////////////////////////////////////////////////////////////////////////////////
func (this *LdapManager) synchronizeUsers(id string) error {

	// Now i will create the user entry found in the ldap server...
	var base_dn string = "OU=Users," + this.getConfigsInfo()[id].M_searchBase
	var filter string = "(objectClass=user)"
	var accountId string

	// a configuration file...
	var attributes []string = []string{"sAMAccountName", "givenName", "mail", "telephoneNumber", "userPrincipalName", "distinguishedName"}

	results, err := this.search(id, "", "", base_dn, filter, attributes)

	if err != nil {
		log.Println("error, fail to search the information on ldap...")
		return err
	}

	// Here I will print the information...
	for i := 0; i < len(results); i++ {
		row := results[i]
		user := new(CargoEntities.User)

		for j := 0; j < len(row); j++ {
			// Print the result...
			if attributes[j] == "givenName" {
				// Here I will split the name to get the first name, last name
				// and middle letter as needed...
				values := strings.Split(row[j].(string), " ")
				if len(values) > 0 {
					if len(values) == 1 {
						user.SetFirstName(strings.TrimSpace(values[0]))
					} else if len(values) == 2 {
						user.SetFirstName(strings.TrimSpace(values[0]))
						user.SetMiddle(strings.TrimSpace(values[1]))
					}
				}
			} else if attributes[j] == "distinguishedName" {
				index := strings.Index(row[j].(string), "\\")
				if index == -1 {
					index = strings.Index(row[j].(string), ",")
				}
				if index > 3 {
					lastName := strings.TrimSpace(strings.TrimSpace(row[j].(string)[3:index]))
					lastName = strings.ToUpper(lastName[0:1]) + strings.ToLower(lastName[1:])
					user.SetLastName(lastName)
				}
			} else if attributes[j] == "mail" {
				user.SetEmail(row[j].(string))
			} else if attributes[j] == "telephoneNumber" {
				user.SetPhone(row[j].(string))
			} else if attributes[j] == "sAMAccountName" {
				user.SetId(strings.ToLower(row[j].(string)))
			} else if attributes[j] == "userPrincipalName" {
				accountId = strings.ToLower(row[j].(string))
			}
		}

		// Specific ...
		// here i will test if the user exist...
		userEntity, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.User", "CargoEntities", []interface{}{user.GetId()})
		if userEntity == nil {
			if len(user.GetEmail()) > 0 {

				// The user must be save...
				if len(user.GetEmail()) > 0 {

					// Create the account in memory...
					if len(accountId) > 0 {
						account := new(CargoEntities.Account)
						account.M_id = accountId
						account.M_password = "Dowty123"
						account.M_name = user.GetId()
						account.M_email = user.GetEmail()

						// Set the account uuid.
						accountEntity, err := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntitiesUuid(), "M_entities", account)
						account = accountEntity.(*CargoEntities.Account)

						// Link the account and the user...
						if err == nil {
							userEntity, err := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntitiesUuid(), "M_entities", user)
							user = userEntity.(*CargoEntities.User)
							if err == nil {
								log.Println("--> Create user: ", user.GetId())
								account.SetUserRef(user)
								user.AppendAccounts(account)
								// Save both entity...
								GetServer().GetEntityManager().saveEntity(account)
								GetServer().GetEntityManager().saveEntity(user)
							} else {
								log.Fatal("------> fail to create user!")
							}
						}
					} else {
						// save only the user here.
						GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntitiesUuid(), "M_entities", user)
						log.Println("--> User ", user.GetFirstName()+" "+user.GetLastName(), " is not active")
					}
				}
			}
		}

	}

	// No error...
	return nil
}

////////////////////////////////////////////////////////////////////////////////
//	Group
////////////////////////////////////////////////////////////////////////////////

/**
 * This Get the LDAP groups from the DB...
 */
func (this *LdapManager) synchronizeGroups(id string) error {
	// log.Println("--> Synchronize Groups")
	var base_dn string = "OU=Groups," + this.getConfigsInfo()[id].M_searchBase
	var filter string = "(objectClass=group)"
	var attributes []string = []string{"name", "distinguishedName"}

	results, err := this.search(id, "", "", base_dn, filter, attributes)

	if err != nil {
		log.Println("error, fail to search the information on ldap...")
		return err
	}

	// Here I will print the information...
	for i := 0; i < len(results); i++ {
		// Here I will get the user in the group...
		row := results[i]
		group := new(CargoEntities.Group)

		for j := 0; j < len(row); j++ {
			// Print the result...
			if attributes[j] == "distinguishedName" {
				// First of all I will retreive the group itself...
				// Now I will retrive user inside this group...
				membersRef, err := this.getLdapGroupMembers(id, row[j].(string))
				if err == nil {
					// if the number of members is not null...
					if len(membersRef) > 0 {
						groupEntity, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Group", "CargoEntities", []interface{}{group.GetId()})
						if groupEntity == nil {
							// Here i will save the group...
							groupEntity, err := GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntitiesUuid(), "M_entities", group)
							if err == nil {
								group = groupEntity.(*CargoEntities.Group)
								log.Println("--> create group ", group.GetId())
							}
						} else {
							group = groupEntity.(*CargoEntities.Group)
						}

						for k := 0; k < len(membersRef); k++ {
							ids := []interface{}{strings.TrimSpace(strings.ToLower(membersRef[k]))}
							member, err := GetServer().GetEntityManager().getEntityById("CargoEntities.User", "CargoEntities", ids)
							if err == nil {
								group.AppendMembersRef(member.(*CargoEntities.User))
								member.(*CargoEntities.User).AppendMemberOfRef(group)
								GetServer().GetEntityManager().saveEntity(member)
							}
						}
					} else {
						group = nil
					}
				}

			} else if attributes[j] == "name" {
				// Now I will retrive user inside this group...
				group.SetId(row[j].(string))
				group.SetName(row[j].(string))
			}
		}
		if group != nil {
			GetServer().GetEntityManager().saveEntity(group)
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
//	Computer
////////////////////////////////////////////////////////////////////////////////
/**
 * Return a computer with a given id.
 */
func (this *LdapManager) getComputer(id string) (*CargoEntities.Computer, *CargoEntities.Error) {
	ids := []interface{}{id}
	computerEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Computer", "CargoEntities", ids)
	if errObj == nil {
		computer := computerEntity.(*CargoEntities.Computer)
		return computer, errObj
	}
	return nil, errObj
}

/**
 * Return a computer with a given name
 */
func (this *LdapManager) getComputerByName(name string) (*CargoEntities.Computer, *CargoEntities.Error) {

	var query EntityQuery
	query.TypeName = "CargoEntities.Computer"
	query.Query = "CargoEntities.Computer.M_name==\"" + name + "\""
	query.Fields = append(query.Fields, "M_id")

	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err)
		return nil, cargoError
	}

	// Here nothing was found...
	if len(results) == 0 {
		cargoError := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("No computer was found with name "+name))
		return nil, cargoError
	}

	// Get the computer with it name...
	computerEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Computer", "CargoEntities", results[0])
	if errObj == nil {
		computer := computerEntity.(*CargoEntities.Computer)
		addrs, err := net.LookupIP(computer.GetName())
		if err == nil {
			for _, addr := range addrs {
				if ipv4 := addr.To4(); ipv4 != nil {
					computer.SetIpv4(ipv4.String())
				}
			}
		}
		return computer, errObj
	}

	return nil, errObj
}

/**
 * Return a computer with a given Ip adress...
 */
func (this *LdapManager) getComputerByIp(ip string) (*CargoEntities.Computer, *CargoEntities.Error) {

	// Il will make a lookup first and test if the computer contain the name
	ids, _ := net.LookupAddr(ip)
	if len(ids) > 0 {
		return this.getComputer(strings.ToUpper(ids[0]))
	} else {
		hostname, _ := os.Hostname()
		computer, err := this.getComputerByName(strings.ToUpper(hostname))
		if err == nil {
			return computer, nil
		}

		if len(hostname) > 0 {
			computer := new(CargoEntities.Computer)
			computer.M_name = strings.ToUpper(hostname)
			computer.M_id = strings.ToUpper(hostname)
			computer.M_osType = CargoEntities.OsType_Unknown
			computer.M_platformType = CargoEntities.PlatformType_Unknown
			computer.M_ipv4 = "127.0.0.1"
			computer.SetEntityGetter(getEntityFct)
			computer.SetEntitySetter(setEntityFct)
			computer.SetUuidGenerator(generateUuidFct)
			return computer, nil
		}
	}

	return nil, NewError(Utility.FileLine(), COMPUTER_IP_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The computer with the ip '"+ip+"' was not found. "))
}

/**
 * Synchronize the computers from ldap information.
 */
func (this *LdapManager) synchronizeComputers(id string) error {

	// This is the list of computer.
	var base_dn string = "OU=Computers," + this.getConfigsInfo()[id].M_searchBase
	var filter string = "(objectClass=computer)"
	var attributes []string = []string{"dNSHostName", "distinguishedName"}

	results, err := this.search(id, "", "", base_dn, filter, attributes)

	if err != nil {
		log.Println("error, fail to search the information on ldap...")
		return err
	}

	// Here I will print the information...
	for i := 0; i < len(results); i++ {
		// Here I will get the user in the group...
		row := results[i]
		computer := new(CargoEntities.Computer)
		for j := 0; j < len(row); j++ {
			// Print the result...
			if attributes[j] == "distinguishedName" {
				// First of all I will retreive the group itself...
				values := strings.Split(row[j].(string), ",")
				for k := 0; k < len(values); k++ {
					value := values[k][strings.Index(values[k], "=")+1:]
					if k == 0 {
						computer.SetName(strings.ToUpper(value))
					} else if k == 1 {
						if value == "Windows7" {
							computer.SetOsType(CargoEntities.OsType_Windows7)
						} else if value == "Windows8" {
							computer.SetOsType(CargoEntities.OsType_Windows8)
						} else if value == "OSX" {
							computer.SetOsType(CargoEntities.OsType_OSX)
						} else if value == "Linux" {
							computer.SetOsType(CargoEntities.OsType_Linux)
						} else if value == "IOS" {
							computer.SetOsType(CargoEntities.OsType_IOS)
						} else {
							computer.SetOsType(CargoEntities.OsType_Unknown)
						}
					} else if k == 2 {
						if value == "Desktops" {
							computer.SetPlatformType(CargoEntities.PlatformType_Desktop)
						} else if value == "Laptops" {
							computer.SetPlatformType(CargoEntities.PlatformType_Laptop)
						} else if value == "Phone" {
							computer.SetPlatformType(CargoEntities.PlatformType_Phone)
						} else if value == "Tablet" {
							computer.SetPlatformType(CargoEntities.PlatformType_Tablet)
						} else {
							computer.SetPlatformType(CargoEntities.PlatformType_Unknown)
						}
					}
				}
			} else if attributes[j] == "dNSHostName" {
				computer.SetId(strings.ToUpper(row[j].(string)))
			}
		}

		computerEntity, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Computer", "CargoEntities", []interface{}{computer.GetId()})
		if computerEntity == nil {
			addrs, err := net.LookupIP(computer.GetName())
			for _, addr := range addrs {
				if ipv4 := addr.To4(); ipv4 != nil {
					computer.SetIpv4(ipv4.String())
				}
			}
			if err != nil {
				log.Println("Adress not found!", computer.GetName())
			} else {
				log.Println("Save computer", computer.GetName(), computer.GetIpv4())
			}
			computerEntity, _ = GetServer().GetEntityManager().createEntity(GetServer().GetEntityManager().getCargoEntitiesUuid(), "M_entities", computer)
			if computerEntity != nil {
				computer = computerEntity.(*CargoEntities.Computer)
			}
			log.Println("----> create computer ", computer.GetId())
		} /*else {
			addrs, err := net.LookupIP(computer.GetName())
			if err == nil {
				for _, addr := range addrs {
					if ipv4 := addr.To4(); ipv4 != nil {
						computer.SetIpv4(ipv4.String())
					}
				}
			}
			// Call save on Entities...
			computer.UUID = computerEntity.GetUuid()
			computer.ParentUuid = computerEntity.GetParentUuid()
			computer.ParentLnk = computerEntity.GetParentLnk()
			GetServer().GetEntityManager().saveEntity(computer)
			log.Println("----> save computer ", computer.GetName(), computer.GetIpv4())
		}*/
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
//LdapManager.prototype.onEvent = function (evt) {
//    EventHub.prototype.onEvent.call(this, evt)
//}
func (this *LdapManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

// @api 1.0
// Synchronize the computers, users and group of an LDAP server.
// @param {string} id The LDAP server connection id.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Account} The new registered account.
// @scope {restricted}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *LdapManager) Synchronize(id string, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	// Synchronize the list of user...
	err := this.synchronizeUsers(id)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), LDAP_ERROR, SERVER_ERROR_CODE, err))
	}

	// Synchronize the list of group...
	err = this.synchronizeGroups(id)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), LDAP_ERROR, SERVER_ERROR_CODE, err))
	}

	// Synchronize the list of computer...
	err = this.synchronizeComputers(id)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), LDAP_ERROR, SERVER_ERROR_CODE, err))
	}
}

// @api 1.0
// Authenticate a user with a given account id and psswd.
// @param {string} id The account id
// @param {string} password The password associated with the new account.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Account} The new registered account.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *LdapManager) Authenticate(id string, login string, psswd string) bool {
	ok := this.authenticate(id, login, psswd)

	return ok
}

// @api 1.0
// Execute a search over a LDAP server with a given connection id
// @param {string} id The LDAP connection id
// @param {string} login Account name to login over the LDAP server.
// @param {string} password The password associated with the login.
// @param {string} base_dn The base dns query
// @param {string} filter The query filter
// @param {[]string} attributes The query attributes
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {[][]interface{}} A tow dimensional array of values.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *LdapManager) Search(id string, login string, psswd string, base_dn string, filter string, attributes []string, messageId string, sessionId string) [][]interface{} {
	values, err := this.search(id, login, psswd, base_dn, filter, attributes)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), LDAP_ERROR, SERVER_ERROR_CODE, err))
	}

	return values
}

// @api 1.0
// Return computer object from a given IPV4 address
// @param {string} ip The IPV4 computer address
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Computer} The computer associated with the address
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *LdapManager) GetComputerByIp(ip string, messageId string, sessionId string) *CargoEntities.Computer {
	computer, err := this.getComputerByIp(ip)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, err)
		return nil
	}

	return computer
}

// @api 1.0
// Return computer object with a given name
// @param {string} name The computer domain name on the network
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @return {*CargoEntities.Computer} The computer associated with the address
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *LdapManager) GetComputerByName(name string, messageId string, sessionId string) *CargoEntities.Computer {
	computer, err := this.getComputerByName(name)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, err)
		return nil
	}
	return computer
}
