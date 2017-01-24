package Server

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"os"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Utility"

	"code.myceliUs.com/CargoWebServer/Cargo/Config/CargoConfig"
	"code.myceliUs.com/CargoWebServer/Cargo/Persistence/CargoEntities"
	LDAP "github.com/mavricknz/ldap"
)

type LdapManager struct {
	// Contain the list of avalable ldap servers...
	m_configsInfo map[string]CargoConfig.LdapConfiguration
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
	ldapManager.m_configsInfo = make(map[string]CargoConfig.LdapConfiguration)

	return ldapManager
}

/**
 * That function is use to synchronize the information of a ldap server
 * with a given id.
 */
func (this *LdapManager) Initialize() {

}

func (this *LdapManager) Start() {
	log.Println("--> Start LdapManager")
	ldapConfigurations := GetServer().GetConfigurationManager().GetLdapConfigurations()

	for i := 0; i < len(ldapConfigurations); i++ {
		this.m_configsInfo[ldapConfigurations[i].M_id] = ldapConfigurations[i]
	}

	// configure all information from the servers...
	for _, info := range this.m_configsInfo {

		// Synchronize the list of user...
		err := this.SynchronizeUsers(info.M_id)
		if err != nil {
			log.Println("Synchronize Users Error ", err)
		}

		// Synchronize the list of group...
		err = this.SynchronizeGroups(info.M_id)
		if err != nil {
			log.Println("Synchronize Groups Error ", err)
		}

		// Synchronize the list of computer...
		err = this.SynchronizeComputers(info.M_id)
		if err != nil {
			log.Println("Synchronize Computers Error ", err)
		}

	}
}

func (this *LdapManager) Stop() {
	log.Println("--> Stop LdapManager")
}

/**
 * Connect to a ldap server...
 */
func (this *LdapManager) Connect(id string, userId string, psswd string) (*LDAP.LDAPConnection, error) {
	ldapConfigInfo := this.m_configsInfo[id]

	conn := LDAP.NewLDAPConnection(ldapConfigInfo.M_ipv4, uint16(ldapConfigInfo.M_port))
	err := conn.Connect()
	if err != nil {
		// handle error
		log.Println("Cannot open the connection: ", ldapConfigInfo.M_hostName)
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
func (this *LdapManager) Search(id string, login string, psswd string, base_dn string, filter string, attributes []string) ([][]interface{}, error) {
	ldapConfigInfo := this.m_configsInfo[id]
	// Try to connect to the ldap server...
	var err error
	var conn *LDAP.LDAPConnection
	if len(login) == 0 {
		conn, err = this.Connect(ldapConfigInfo.M_id, ldapConfigInfo.M_user, ldapConfigInfo.M_pwd)
		if err != nil {
			log.Println("Fail to connect:", ldapConfigInfo.M_id)
			return nil, err
		}
	} else {
		conn, err = this.Connect(ldapConfigInfo.M_id, login, psswd)
		if err != nil {
			log.Println("Fail to connect:", ldapConfigInfo.M_id)
			return nil, err
		}
	}

	// The connection will be close latter...
	defer conn.Close()

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

func (this *LdapManager) Authenticate(id string, login string, psswd string) bool {

	ldapConfigInfo := this.m_configsInfo[id]
	// Now I will try to make a simple query if it fail that's mean the user
	// does have the permission...
	var filter string = "(objectClass=user)"

	// Test get some user...
	var attributes []string = []string{"sAMAccountName"}
	//log.Println("authenticate account ", name, " login ", login)
	_, err := this.Search(id, login, psswd, ldapConfigInfo.M_searchBase, filter, attributes)
	if err != nil {
		return false
	}

	return true
}

////////////////////////////////////////////////////////////////////////////////
//	User
////////////////////////////////////////////////////////////////////////////////
func (this *LdapManager) SynchronizeUsers(id string) error {

	// Now i will create the user entry found in the ldap server...
	var base_dn string = "OU=Users,OU=MON,OU=CA,DC=UD6,DC=UF6"
	var filter string = "(objectClass=user)"
	var accountId string

	// a configuration file...
	var attributes []string = []string{"sAMAccountName", "name", "mail", "telephoneNumber", "userPrincipalName", "distinguishedName"}

	results, err := this.Search(id, "", "", base_dn, filter, attributes)

	if err != nil {
		log.Println("error, fail to search the information on ldap...")
		return err
	}

	cargoEntities := GetServer().GetEntityManager().getCargoEntities()

	// Here I will print the information...
	for i := 0; i < len(results); i++ {
		row := results[i]
		userEntity := GetServer().GetEntityManager().NewCargoEntitiesUserEntity("", nil)
		user := userEntity.GetObject().(*CargoEntities.User)
		for j := 0; j < len(row); j++ {
			// Print the result...
			if attributes[j] == "name" {
				// Here I will split the name to get the first name, last name
				// and middle letter as needed...
				values := strings.Split(row[j].(string), ",")
				if len(values) > 0 {
					user.SetLastName(strings.TrimSpace(values[0]))
				}
				if len(values) > 1 {
					values_ := strings.Split(values[1], " ")
					if len(values_) > 1 {
						if len(values_[1]) > 2 {
							user.SetFirstName(strings.TrimSpace(values[1]))
						} else {
							user.SetFirstName(strings.TrimSpace(values_[0]))
							user.SetMiddle(strings.TrimSpace(values_[1]))
						}
					} else {
						user.SetFirstName(strings.TrimSpace(values[1]))
					}
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
		userUuid := CargoEntitiesUserExists(user.M_id)
		if len(userUuid) == 0 && len(user.GetEmail()) > 0 {
			// The user must be save...
			log.Println("Save user ", user.GetEmail(), "an id ", user.GetId())
			if len(user.GetEmail()) > 0 {
				accountEntity := GetServer().GetEntityManager().NewCargoEntitiesAccountEntity(accountId, nil)
				// Create the account in memory...
				account := accountEntity.GetObject().(*CargoEntities.Account)
				account.M_id = accountId
				account.M_password = "Dowty123"
				account.M_name = user.GetId()
				account.M_email = user.GetEmail()

				// Append the newly create account into the cargo entities
				entities := cargoEntities.GetObject().(*CargoEntities.Entities)
				entities.SetEntities(account)
				account.SetEntitiesPtr(entities)

				// Link the account and the user...
				account.SetUserRef(user)
				account.NeedSave = true
				user.SetAccounts(account)

				// Now append the user to the entities...
				entities.SetEntities(user)
				user.SetEntitiesPtr(entities)
			}
		}

	}
	// Call save on Entities...
	cargoEntities.SaveEntity()

	// No error...
	return nil
}

/**
 * Get the list of group for a user...
 */
func (this *LdapManager) getLdapUserMemberOf(id string, userId string) ([]string, error) {

	var base_dn string = "OU=MON,OU=CA,DC=UD6,DC=UF6"

	var filter string = "(&(objectClass=group)(objectcategory=Group)(member=" + userId + "))"
	var attributes []string = []string{"sAMAccountName"}
	results, err := this.Search(id, "", "", base_dn, filter, attributes)
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

/**
 * That function return the list of all user's register in the database.
 */
func (this *LdapManager) GetAllUsers() ([]*CargoEntities.User, *CargoEntities.Error) {
	var allUsers []*CargoEntities.User
	entities, errObj := GetServer().GetEntityManager().getEntitiesByType("CargoEntities.User", "", CargoEntitiesDB)
	if errObj != nil {
		return nil, errObj
	}
	for i := 0; i < len(entities); i++ {
		allUsers = append(allUsers, entities[i].GetObject().(*CargoEntities.User))
	}
	return allUsers, nil
}

/**
 * Return a user with a given id.
 */
func (this *LdapManager) GetUserById(id string) (*CargoEntities.User, *CargoEntities.Error) {
	userEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.User", id)
	if errObj == nil {
		user := userEntity.GetObject().(*CargoEntities.User)
		return user, nil
	}
	return nil, errObj
}

////////////////////////////////////////////////////////////////////////////////
//	Group
////////////////////////////////////////////////////////////////////////////////

/**
 * This Get the LDAP groups from the DB...
 */
func (this *LdapManager) SynchronizeGroups(id string) error {
	var base_dn string = "OU=Groups,OU=MON,OU=CA,DC=UD6,DC=UF6"
	var filter string = "(objectClass=group)"
	var attributes []string = []string{"name", "distinguishedName"}

	results, err := this.Search(id, "", "", base_dn, filter, attributes)

	if err != nil {
		log.Println("error, fail to search the information on ldap...")
		return err
	}

	// Here I will print the information...
	for i := 0; i < len(results); i++ {
		// Here I will get the user in the group...
		row := results[i]
		groupEntity := GetServer().GetEntityManager().NewCargoEntitiesGroupEntity("", nil)
		group := groupEntity.GetObject().(*CargoEntities.Group)
		for j := 0; j < len(row); j++ {
			// Print the result...
			if attributes[j] == "distinguishedName" {
				// First of all I will retreive the group itself...
				// Now I will retrive user inside this group...
				membersRef, err := this.getLdapGroupMembers(id, row[j].(string))
				if err == nil {
					for k := 0; k < len(membersRef); k++ {
						member, err := GetServer().GetEntityManager().getEntityById("CargoEntities.User", membersRef[k])
						if err == nil {
							group.SetMembersRef(member.GetObject().(*CargoEntities.User))
							member.GetObject().(*CargoEntities.User).SetMemberOfRef(group)
							member.SaveEntity() // save the user...
						}
					}
					// if the number of members is not null...
					if len(membersRef) > 0 {
						groupUuid := CargoEntitiesGroupExists(group.M_id)
						if len(groupUuid) == 0 {
							// Here i will save the group...
							entities := GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities)
							entities.SetEntities(group)
							group.SetEntitiesPtr(entities)
						}
					}
				}

			} else if attributes[j] == "name" {
				// Now I will retrive user inside this group...
				group.SetId(row[j].(string))
				group.SetName(row[j].(string))
			}
		}
	}

	// Call save on Entities...
	GetServer().GetEntityManager().getCargoEntities().SaveEntity()
	return nil
}

func (this *LdapManager) getLdapGroupMembers(id string, groupId string) ([]string, error) {
	var base_dn string = "OU=MON,OU=CA,DC=UD6,DC=UF6"

	var filter string = "(&(objectClass=user)(objectcategory=Person)(memberOf=" + groupId + "))"
	var attributes []string = []string{"sAMAccountName"}
	results, err := this.Search(id, "", "", base_dn, filter, attributes)
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
				//if strings.HasPrefix(userId, "mm") == true || strings.HasPrefix(userId, "mtmx") == true || strings.HasPrefix(userId, "mrmfct") == true {
				// Save user from ldap directory to the database...
				members = append(members, userId)
				//}
			}
		}
	}
	return members, nil
}

/**
 * That function return the list of all group's register in the database.
 */
func (this *LdapManager) GetAllGroups() ([]*CargoEntities.Group, *CargoEntities.Error) {
	var allGroups []*CargoEntities.Group
	entities, errObj := GetServer().GetEntityManager().getEntitiesByType("CargoEntities.Group", "", CargoEntitiesDB)
	if errObj != nil {
		return nil, errObj
	}
	for i := 0; i < len(entities); i++ {
		allGroups = append(allGroups, entities[i].GetObject().(*CargoEntities.Group))
	}
	return allGroups, nil
}

/**
 * Return a group with a given id.
 */
func (this *LdapManager) GetGroupById(id string) (*CargoEntities.Group, *CargoEntities.Error) {
	groupEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Group", id)
	if errObj == nil {
		group := groupEntity.GetObject().(*CargoEntities.Group)
		return group, nil
	}
	return nil, errObj
}

////////////////////////////////////////////////////////////////////////////////
//	Computer
////////////////////////////////////////////////////////////////////////////////

/**
 * Synchronize the computers from ldap information.
 */
func (this *LdapManager) SynchronizeComputers(id string) error {
	// This is the list of computer.
	var base_dn string = "OU=Computers,OU=MON,OU=CA,DC=UD6,DC=UF6"
	var filter string = "(objectClass=computer)"
	var attributes []string = []string{"dNSHostName", "distinguishedName"}

	results, err := this.Search(id, "", "", base_dn, filter, attributes)

	if err != nil {
		log.Println("error, fail to search the information on ldap...")
		return err
	}

	// Here I will print the information...
	for i := 0; i < len(results); i++ {
		// Here I will get the user in the group...
		row := results[i]
		computerEntity := GetServer().GetEntityManager().NewCargoEntitiesComputerEntity("", nil)
		computer := computerEntity.GetObject().(*CargoEntities.Computer)

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
		/**/
		computerUuid := CargoEntitiesComputerExists(computer.M_id)
		if len(computerUuid) == 0 {
			/*addrs, err := net.LookupIP(computer.GetName())
			for _, addr := range addrs {
				if ipv4 := addr.To4(); ipv4 != nil {
					computer.SetIpv4(ipv4.String())
				}
			}
			if err != nil {
				log.Println("Adress not found!", computer.GetName())
			} else {
				log.Println("Save computer", computer.GetName(), computer.GetIpv4())
			}*/
			entities := GetServer().GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities)
			entities.SetEntities(computer)
			computer.SetEntitiesPtr(entities)
		}
	}
	// Call save on Entities...
	GetServer().GetEntityManager().getCargoEntities().SaveEntity()
	return nil
}

/**
 * Return a computer with a given id.
 */
func (this *LdapManager) GetComputer(id string) (*CargoEntities.Computer, *CargoEntities.Error) {
	computerEntity, errObj := GetServer().GetEntityManager().getEntityById("CargoEntities.Computer", id)
	if errObj == nil {
		computer := computerEntity.GetObject().(*CargoEntities.Computer)
		return computer, errObj
	}
	return nil, errObj
}

/**
 * Return a computer with a given name
 */
func (this *LdapManager) GetComputerByName(name string) (*CargoEntities.Computer, *CargoEntities.Error) {

	var query EntityQuery
	query.TypeName = "CargoEntities.Computer"
	query.Indexs = append(query.Indexs, "M_name="+name)
	query.Fields = append(query.Fields, "uuid")
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
	computerEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(results[0][0].(string))
	if errObj == nil {
		computer := computerEntity.GetObject().(*CargoEntities.Computer)
		return computer, errObj
	}

	return nil, errObj
}

/**
 * Return a computer with a given Ip adress...
 */
func (this *LdapManager) GetComputerByIp(ip string) (*CargoEntities.Computer, *CargoEntities.Error) {

	// Il will make a lookup first and test if the computer contain the name
	ids, _ := net.LookupAddr(ip)
	if len(ids) > 0 {
		return this.GetComputer(strings.ToUpper(ids[0]))
	} else {
		hostname, _ := os.Hostname()
		computer, err := this.GetComputerByName(strings.ToUpper(hostname))
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
			return computer, nil
		}
	}

	return nil, NewError(Utility.FileLine(), COMPUTER_IP_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The computer with the ip '"+ip+"' was not found. "))
}
