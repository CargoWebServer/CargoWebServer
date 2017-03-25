/**
 * That file contain the code for OAuth2 service.
 */

package Server

import (
	//"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	//"net/url"
	//"strconv"
	"strings"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
	"github.com/RangelReale/osin"
)

/**
 * Clear expiring authorization and (access/refresh)
 */
func clearCodeExpired(code string) {
	GetServer().GetOAuth2Manager().m_store.RemoveAuthorize(code)
	GetServer().GetOAuth2Manager().m_store.RemoveAccess(code)
}

/**
 * Use a timer to execute clearExpiredCode when it can at end...
 */
func setCodeExpiration(code string, duration time.Duration) {
	// Create a closure and wrap the code.
	f := func(code string) func() {
		return func() {
			clearCodeExpired(code)
		}
	}(code)

	// The function will be call after the duration.
	time.AfterFunc(duration, f)
}

type OAuth2Manager struct {

	// the data stores.
	m_store *OAuth2Store

	// the oauth sever.
	m_server *osin.Server
}

var oauth2Manager *OAuth2Manager

func (this *Server) GetOAuth2Manager() *OAuth2Manager {
	if oauth2Manager == nil {
		oauth2Manager = newOAuth2Manager()
	}
	return oauth2Manager
}

func newOAuth2Manager() *OAuth2Manager {
	oauth2Manager := new(OAuth2Manager)
	return oauth2Manager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * That function is use to synchronize the information of a ldap server
 * with a given id.
 */
func (this *OAuth2Manager) initialize() {
	// register service avalaible action here.
	log.Println("--> initialyze OAuth2Manager")
	// Create the default configurations
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())
}

func (this *OAuth2Manager) getId() string {
	return "OAuth2Manager"
}

func (this *OAuth2Manager) start() {
	log.Println("--> Start OAuth2Manager")

	// Here I will intialyse configurations.
	cfg := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	var sconfig *osin.ServerConfig
	if cfg == nil {
		// Get the default configuration.
		sconfig = osin.NewServerConfig()

		// Set default parameters here.
		sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
		sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE,
			osin.REFRESH_TOKEN, osin.PASSWORD, osin.CLIENT_CREDENTIALS, osin.ASSERTION}
		sconfig.AllowGetAccessRequest = true
		sconfig.AllowClientSecretInParams = true

		// Save it into the entity.
		cfg := new(Config.OAuth2Configuration)
		cfg.SetId("OAuth2Config")

		cfg.SetAccessExpiration(int(sconfig.AccessExpiration))
		cfg.SetAllowClientSecretInParams(sconfig.AllowClientSecretInParams)
		cfg.SetAllowGetAccessRequest(sconfig.AllowGetAccessRequest)

		for i := 0; i < len(sconfig.AllowedAuthorizeTypes); i++ {
			if sconfig.AllowedAuthorizeTypes[i] == osin.CODE {
				cfg.SetAllowedAuthorizeTypes("code")
			} else if sconfig.AllowedAuthorizeTypes[i] == osin.TOKEN {
				cfg.SetAllowedAuthorizeTypes("token")
			}
		}

		for i := 0; i < len(sconfig.AllowedAccessTypes); i++ {
			if sconfig.AllowedAccessTypes[i] == osin.AUTHORIZATION_CODE {
				cfg.SetAllowedAccessTypes("authorization_code")
			} else if sconfig.AllowedAccessTypes[i] == osin.REFRESH_TOKEN {
				cfg.SetAllowedAccessTypes("refresh_token")
			} else if sconfig.AllowedAccessTypes[i] == osin.PASSWORD {
				cfg.SetAllowedAccessTypes("password")
			} else if sconfig.AllowedAccessTypes[i] == osin.CLIENT_CREDENTIALS {
				cfg.SetAllowedAccessTypes("client_credentials")
			} else if sconfig.AllowedAccessTypes[i] == osin.ASSERTION {
				cfg.SetAllowedAccessTypes("assertion")
			} else if sconfig.AllowedAccessTypes[i] == osin.IMPLICIT {
				cfg.SetAllowedAccessTypes("__implicit")
			}
		}

		cfg.SetAuthorizationExpiration(int(sconfig.AuthorizationExpiration))
		cfg.SetErrorStatusCode(sconfig.ErrorStatusCode)
		cfg.SetRedirectUriSeparator(sconfig.RedirectUriSeparator)
		cfg.SetTokenType(sconfig.TokenType)

		// Append to the active configuration and save the entity.
		configurations := GetServer().GetConfigurationManager().getActiveConfigurations()
		configurations.SetOauth2Configuration(cfg)

		// Save it
		GetServer().GetConfigurationManager().m_activeConfigurationsEntity.SaveEntity()

	} else {
		sconfig = osin.NewServerConfig()

		// Set the access expiration time.
		sconfig.AccessExpiration = int32(cfg.GetAccessExpiration())
		sconfig.AllowClientSecretInParams = cfg.GetAllowClientSecretInParams()
		sconfig.AllowGetAccessRequest = cfg.GetAllowGetAccessRequest()

		for i := 0; i < len(cfg.GetAllowedAuthorizeTypes()); i++ {
			if cfg.GetAllowedAuthorizeTypes()[i] == "code" {
				sconfig.AllowedAuthorizeTypes = append(sconfig.AllowedAuthorizeTypes, osin.CODE)
			} else if cfg.GetAllowedAuthorizeTypes()[i] == "token" {
				sconfig.AllowedAuthorizeTypes = append(sconfig.AllowedAuthorizeTypes, osin.TOKEN)
			}
		}

		for i := 0; i < len(cfg.GetAllowedAccessTypes()); i++ {
			if cfg.GetAllowedAccessTypes()[i] == "authorization_code" {
				sconfig.AllowedAccessTypes = append(sconfig.AllowedAccessTypes, osin.AUTHORIZATION_CODE)
			} else if cfg.GetAllowedAccessTypes()[i] == "refresh_token" {
				sconfig.AllowedAccessTypes = append(sconfig.AllowedAccessTypes, osin.REFRESH_TOKEN)
			} else if cfg.GetAllowedAccessTypes()[i] == "password" {
				sconfig.AllowedAccessTypes = append(sconfig.AllowedAccessTypes, osin.PASSWORD)
			} else if cfg.GetAllowedAccessTypes()[i] == "client_credentials" {
				sconfig.AllowedAccessTypes = append(sconfig.AllowedAccessTypes, osin.CLIENT_CREDENTIALS)
			} else if cfg.GetAllowedAccessTypes()[i] == "assertion" {
				sconfig.AllowedAccessTypes = append(sconfig.AllowedAccessTypes, osin.ASSERTION)
			} else if cfg.GetAllowedAccessTypes()[i] == "__implicit" {
				sconfig.AllowedAccessTypes = append(sconfig.AllowedAccessTypes, osin.IMPLICIT)
			}
		}

		sconfig.AuthorizationExpiration = int32(cfg.GetAuthorizationExpiration())
		sconfig.ErrorStatusCode = cfg.GetErrorStatusCode()
		sconfig.RedirectUriSeparator = cfg.GetRedirectUriSeparator()
		sconfig.TokenType = cfg.GetTokenType()
	}

	// Start the oauth service.
	this.m_store = newOauth2Store()
	this.m_server = osin.NewServer(sconfig, this.m_store)

	// Here I will remove all expired authorization.
	for i := 0; i < len(cfg.GetExpire()); i++ {
		expireTime := time.Unix(cfg.GetExpire()[i].GetExpiresAt(), 0)
		if expireTime.Before(time.Now()) {
			log.Println("------> remove authorize...")
			// I that case the value must be remove expired values...
			this.m_store.RemoveAccess(cfg.GetExpire()[i].GetToken())
			this.m_store.RemoveAuthorize(cfg.GetExpire()[i].GetToken())
		} else {
			setCodeExpiration(cfg.GetExpire()[i].GetToken(), expireTime.Sub(time.Now()))
		}
	}
}

func (this *OAuth2Manager) stop() {
	log.Println("--> Stop OAuth2Manager")
}

////////////////////////////////////////////////////////////////////////////////
// OAuth2 Store
////////////////////////////////////////////////////////////////////////////////
type OAuth2Store struct {
}

// Create and intialyse the OAuth2 Store.
func newOauth2Store() *OAuth2Store {
	store := new(OAuth2Store)
	return store
}

/**
 * Close the store.
 */
func (this *OAuth2Store) Close() {

}

/**
 * Return pointer...
 */
func (s *OAuth2Store) Clone() osin.Storage {
	return s
}

/**
 * Retrun a given client.
 */
func (this *OAuth2Store) GetClient(id string) (osin.Client, error) {
	log.Println("Get Client")
	// From the list of registred client I will retreive the client
	// with the given id.
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	clients := config.GetClients()

	for i := 0; i < len(clients); i++ {
		if clients[i].M_id == id {
			c := new(osin.DefaultClient)
			c.Id = clients[i].M_id
			c.Secret = clients[i].M_secret
			c.RedirectUri = clients[i].M_redirectUri
			c.UserData = clients[i].M_extra
			return c, nil
		}
	}

	// No client with the given id was found.
	return nil, errors.New("No client found with id " + id)
}

/**
 * Set the client value.
 */
func (this *OAuth2Store) SetClient(id string, client osin.Client) error {
	log.Println("Set Client")
	// The configuration.
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	// Create a client configuration from the osin.Client.
	c := new(Config.OAuth2Client)
	c.M_id = id
	c.M_extra = client.GetUserData().([]uint8)
	c.M_secret = client.GetSecret()
	c.M_redirectUri = client.GetRedirectUri()

	// Set the uuid.
	GetServer().GetEntityManager().NewConfigOAuth2ClientEntity(config.GetUUID(), "", client)

	// append a new client.
	config.SetClients(c)

	// Save the entity.
	configEntity.SaveEntity()

	log.Println("Client with id ", id, " was save!")
	return nil
}

/**
 * Save a given autorization.
 */
func (this *OAuth2Store) SaveAuthorize(data *osin.AuthorizeData) error {
	log.Println("Save Authorize")
	// Get the config entity
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	a := new(Config.OAuth2Authorize)

	// Set the client.
	for i := 0; i < len(config.GetClients()); i++ {
		if config.GetClients()[i].M_id == data.Client.GetId() {
			a.SetClient(config.GetClients()[i])
			break
		}
	}

	// Set the value from the data.
	a.SetId(data.Code)
	a.SetExpiresIn(int64(data.ExpiresIn))
	a.SetScope(data.Scope)
	a.SetRedirectUri(data.RedirectUri)
	a.SetState(data.State)
	a.SetCreatedAt(data.CreatedAt.Unix())

	if data.UserData != nil {
		a.SetExtra(data.UserData)
	}

	// Set the uuid
	GetServer().GetEntityManager().NewConfigOAuth2AuthorizeEntity(config.GetUUID(), "", a)

	// Put in the config.
	config.SetAuthorize(a)

	// Save the entity.
	configEntity.SaveEntity()

	log.Println("Save Authorize completed!", a)
	log.Println("save authorize created at ", data.CreatedAt.String())
	log.Println("save authorize expire at ", data.ExpireAt().String())

	// Add expire data.
	if err := this.AddExpireAtData(data.Code, data.ExpireAt()); err != nil {
		log.Println("Fail to create access data ", err)
		return err
	}

	return nil
}

// LoadAuthorize looks up AuthorizeData by a code.
// Client information MUST be loaded together.
// Optionally can return error if expired.
func (this *OAuth2Store) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	log.Println("Load Authorize", code)
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()

	for i := 0; i < len(config.GetAuthorize()); i++ {
		if config.GetAuthorize()[i].GetId() == code {
			var data *osin.AuthorizeData
			data = new(osin.AuthorizeData)
			data.Code = code
			data.ExpiresIn = int32(config.GetAuthorize()[i].GetExpiresIn())
			data.Scope = config.GetAuthorize()[i].GetScope()
			data.RedirectUri = config.GetAuthorize()[i].GetRedirectUri()
			data.State = config.GetAuthorize()[i].GetState()
			data.CreatedAt = time.Unix(config.GetAuthorize()[i].GetCreatedAt(), 0)
			data.UserData = config.GetAuthorize()[i].GetExtra()
			c, err := this.GetClient(config.GetAuthorize()[i].GetClient().GetId())
			if err != nil {
				return nil, err
			}
			data.Client = c

			// Now I will test the expiration time.
			if data.ExpireAt().Before(time.Now()) {
				return nil, errors.New("Token expired at " + data.ExpireAt().String())
			}

			log.Println("Load Authorize Data completed!")
			log.Println("load authorize created at ", data.CreatedAt.String())
			log.Println("load authorize expire at ", data.ExpireAt().String())

			return data, nil
		}
	}

	// No data was found.
	return nil, errors.New("No authorize data found with code " + code)
}

/**
 * Remove authorize from the db.
 */
func (this *OAuth2Store) RemoveAuthorize(code string) error {
	log.Println("Remove Authorize", code)
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()

	for i := 0; i < len(config.GetAuthorize()); i++ {
		if config.GetAuthorize()[i].GetId() == code {
			entity, err := GetServer().GetEntityManager().getEntityByUuid(config.GetAuthorize()[i].GetUUID())
			entity.DeleteEntity()
			if err == nil {
				if err := this.RemoveExpireAtData(code); err != nil {
					return err
				}
			}
			log.Println("Remove authorize complete!")
			return nil
		}
	}
	return errors.New("No authorization with code " + code + " was found!")
}

/**
 * Save the access Data.
 */
func (this *OAuth2Store) SaveAccess(data *osin.AccessData) error {
	log.Println("Save Access")
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	access := new(Config.OAuth2Access)
	for i := 0; i < len(config.GetAccess()); i++ {
		if config.GetAccess()[i].GetId() == data.AccessToken {
			access = config.GetAccess()[i]
			log.Println("access already exist for code ", data.AccessToken)
			break
		}
	}

	access.SetId(data.AccessToken)

	// Set the uuid.
	GetServer().GetEntityManager().NewConfigOAuth2AccessEntity(config.GetUUID(), "", access)

	prev := ""
	authorizeData := &osin.AuthorizeData{}

	if data.AccessData != nil {
		prev = data.AccessData.AccessToken
	}

	if data.AuthorizeData != nil {
		authorizeData = data.AuthorizeData
	}

	if data.Client == nil {
		return errors.New("data.Client must not be nil")
	}

	// Set the client.
	for i := 0; i < len(config.GetClients()); i++ {
		if config.GetClients()[i].GetId() == data.Client.GetId() {
			access.SetClient(config.GetClients()[i])
		}
	}

	// Set the authorization.
	if authorizeData == nil {
		return errors.New("authorize data must not be nil")
	}

	// Keep only the code here no the object because it will be deleted after
	// the access creation.
	access.SetAuthorize(authorizeData.Code)

	// Set other values.
	access.SetPrevious(prev)
	access.SetId(data.AccessToken)

	access.SetExpiresIn(int64(data.ExpiresIn))
	access.SetScope(data.Scope)
	access.SetRedirectUri(data.RedirectUri)

	// Set the unix time.
	access.SetCreatedAt(data.CreatedAt.Unix())

	if data.UserData != nil {
		access.SetExtra(data.UserData)
	}

	// Add expire data.
	if err := this.AddExpireAtData(data.AccessToken, data.ExpireAt()); err != nil {
		log.Println("Fail to create access data ", err)
		return err
	}

	// Now the refresh token.
	if len(data.RefreshToken) > 0 {
		// In that case I will save the refresh token.
		var r *Config.OAuth2Refresh
		for i := 0; i < len(config.GetRefresh()); i++ {
			if config.GetRefresh()[i].GetId() == data.RefreshToken && config.GetRefresh()[i].GetAccess().GetId() == data.AccessToken {
				r = config.GetRefresh()[i]
			}
		}

		// Here the refresh token dosent exist so i will create it.
		if r == nil {
			r = new(Config.OAuth2Refresh)
			r.SetId(data.RefreshToken)
			r.SetAccess(access)
			// Set the object uuid...
			GetServer().GetEntityManager().NewConfigOAuth2RefreshEntity(config.GetUUID(), "", r)

			// Set the access
			access.SetRefreshToken(r)

			// Set into it parent.
			config.SetRefresh(r)
		}
	}

	// save the access.
	config.SetAccess(access)
	configEntity.SetNeedSave(true)
	configEntity.SaveEntity()

	return nil
}

/**
 * Load the access for a given code.
 */
func (this *OAuth2Store) LoadAccess(code string) (*osin.AccessData, error) {
	log.Println("Load Access ", code)
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()

	var access *osin.AccessData
	for i := 0; i < len(config.GetAccess()); i++ {
		if config.GetAccess()[i].GetId() == code {
			access = new(osin.AccessData)
			access.AccessToken = code
			access.ExpiresIn = int32(config.GetAccess()[i].GetExpiresIn())
			access.Scope = config.GetAccess()[i].GetScope()
			access.RedirectUri = config.GetAccess()[i].GetRedirectUri()
			access.CreatedAt = time.Unix(int64(config.GetAccess()[i].GetCreatedAt()), 0)
			access.UserData = config.GetAccess()[i].GetExtra()

			// The refresh token
			access.RefreshToken = config.GetAccess()[i].GetId()

			// The access token
			access.AccessToken = config.GetAccess()[i].GetId()

			// Now the client
			c, err := this.GetClient(config.GetAccess()[i].GetClient().GetId())
			if err != nil {
				return nil, err
			}
			access.Client = c

			// The authorize
			auth, err := this.LoadAuthorize(config.GetAccess()[i].GetAuthorize())
			if err != nil {
				return nil, err
			}
			access.AuthorizeData = auth
		}
	}

	log.Println("Load: access create at ", access.CreatedAt.String())
	log.Println("load: access expire at ", access.ExpireAt().String())

	return access, nil
}

/**
 * Remove an access code.
 */
func (this *OAuth2Store) RemoveAccess(code string) error {
	log.Println("Remove Access ", code)
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	for i := 0; i < len(config.GetAccess()); i++ {
		if config.GetAccess()[i].GetId() == code {
			// Get the entity and delete it.
			accessEntity, err := GetServer().GetEntityManager().getEntityByUuid(config.GetAccess()[i].GetUUID())
			if err == nil {
				accessEntity.DeleteEntity()
				// Remove the expire data.
				if err := this.RemoveExpireAtData(code); err != nil {
					return err
				}

				// Now remove the refresh.
				this.RemoveRefresh(config.GetAccess()[i].GetRefreshToken().GetId())
			}

		}
	}
	return errors.New("No access with code " + code + " was found.")
}

/**
 * Load the access data from it refresh code.
 */
func (this *OAuth2Store) LoadRefresh(code string) (*osin.AccessData, error) {
	log.Println("Load Refresh ", code)
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()

	for i := 0; i < len(config.GetRefresh()); i++ {
		if config.GetRefresh()[i].GetId() == code {
			return this.LoadAccess(config.GetRefresh()[i].GetAccess().GetId())
		}
	}

	return nil, errors.New("No access data was found for " + code)
}

/**
 * Remove refresh.
 */
func (this *OAuth2Store) RemoveRefresh(code string) error {
	log.Println("Remove Refresh ", code)
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	for i := 0; i < len(config.GetRefresh()); i++ {
		if config.GetRefresh()[i].GetId() == code {
			entity, err := GetServer().GetEntityManager().getEntityByUuid(config.GetRefresh()[i].GetUUID())
			if err == nil {
				entity.DeleteEntity()
				return nil
			}
			return errors.New("Fail to delete refresh with code " + code)
		}
	}

	return errors.New("No refresh was found with code " + code)
}

// AddExpireAtData add info in expires table
func (s *OAuth2Store) AddExpireAtData(code string, expireAt time.Time) error {
	expire := new(Config.OAuth2Expires)
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	lastId := 0
	for i := 0; i < len(config.GetExpire()); i++ {
		if config.GetExpire()[i].GetId() > lastId {
			lastId = config.GetExpire()[i].GetId()
		}
	}

	lastId++
	expire.SetId(lastId)
	expire.SetToken(code)
	expire.SetExpiresAt(expireAt.Unix())
	config.SetExpire(expire)
	configEntity.SaveEntity()

	// Start the timer.
	duration := expireAt.Sub(time.Now())
	setCodeExpiration(code, duration)

	return nil
}

// RemoveExpireAtData remove info in expires table
func (s *OAuth2Store) RemoveExpireAtData(code string) error {
	log.Println("Remove Expire At Data", code)
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()

	for i := 0; i < len(config.GetExpire()); i++ {

		if config.GetExpire()[i] == nil {
			entity, err := GetServer().GetEntityManager().getEntityByUuid(config.GetExpire()[i].GetUUID())
			if err == nil {
				entity.DeleteEntity()
			}
		} else {
			if config.GetExpire()[i].GetToken() == code {
				entity, err := GetServer().GetEntityManager().getEntityByUuid(config.GetExpire()[i].GetUUID())
				if err == nil {
					entity.DeleteEntity()
				}
			}
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// The api
////////////////////////////////////////////////////////////////////////////////

/**
 * That function is use to get a given ressource for a given client.
 */
func (this *OAuth2Manager) GetRessource(clientId string, scope string, query string, messageId string, sessionId string) (interface{}, *CargoEntities.Error) {

	log.Println("=----------------> GetRessource")

	// So here the first thing to do is to get the client.
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()

	var client *Config.OAuth2Client
	for i := 0; i < len(config.GetClients()); i++ {
		if clientId == config.GetClients()[i].GetId() {
			client = config.GetClients()[i]
		}
	}

	if client == nil {
		// Repport the error
		errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("No client was found with id "+clientId))
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}

	log.Println("=----------------> Client found!")

	var access *Config.OAuth2Access
	// Now the client was found I will try to get an access code for the given
	// scope and client.
	for i := 0; i < len(config.GetAccess()) && access == nil; i++ {
		a := config.GetAccess()[i]
		if a.GetClient().GetId() == client.GetId() {
			values := strings.Split(a.GetScope(), " ")
			values_ := strings.Split(scope, " ")
			hasScope := true
			for j := 0; j < len(values_); j++ {
				if !Utility.Contains(values, values_[j]) {
					hasScope = false
					break
				}
			}
			// If the access has the correct scope.
			if hasScope {
				access = a
			}
		}
	}

	if access != nil {
		// Here I will made the API call.

	} else {
		log.Println("=----------------> authorization needed!")
		// No access was found so here I will initiated the OAuth process...
		// To do so I will create the href where the user will be ask to
		// authorize the client application to access ressources.
		var authorizationLnk = client.GetAuthorizationUri()
		authorizationLnk += "?response_type=code&client_id=" + client.GetId()
		authorizationLnk += "&state=" + messageId + ":" + sessionId + "&scope=" + scope
		authorizationLnk += "&redirect_uri=" + client.GetRedirectUri() + "/code"

		// I will create the request and send it to the client...
		id := Utility.RandomUUID()
		method := "OAuth2Authorize"
		params := make([]*MessageData, 1)
		data := new(MessageData)
		data.Name = "authorizationLnk"
		data.Value = authorizationLnk
		params[0] = data
		to := make([]connection, 1)
		to[0] = GetServer().getConnectionById(sessionId)
		oauth2Authorize, err := NewRequestMessage(id, method, params, to, nil, nil, nil)
		if err == nil {
			// Send the request.
			GetServer().GetProcessor().m_pendingRequestChannel <- oauth2Authorize
		}

	}

	var result interface{}

	return result, nil
}

////////////////////////////////////////////////////////////////////////////////
// Auth Http handler.
////////////////////////////////////////////////////////////////////////////////

/**
 * If the use is not logged...
 */
func HandleLoginPage(ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()

	user := r.Form.Get("login")
	pwd := r.Form.Get("password")
	state := r.URL.Query()["state"][0]

	var sessionId string
	var messageId string
	if len(strings.Split(state, ":")) == 2 {
		messageId = strings.Split(state, ":")[0]
		sessionId = strings.Split(state, ":")[1]
	}

	if r.Method == "POST" && len(sessionId) > 0 && len(messageId) > 0 {
		session := GetServer().GetSessionManager().Login(user, pwd, "", messageId, sessionId)
		if session != nil {
			log.Println("--------> user " + user + " is logged in!")
			return true
		}
	}

	w.Write([]byte("<html><body>"))

	// if the user is no logged...
	w.Write([]byte(fmt.Sprintf("<form class='authorizationForm' action=\"/authorize?%s\" method=\"POST\">", r.URL.RawQuery)))
	w.Write([]byte("Login: <input type=\"text\" name=\"login\" /><br/>"))
	w.Write([]byte("Password: <input type=\"password\" name=\"password\" /><br/>"))
	w.Write([]byte("<input type=\"submit\"/>"))
	w.Write([]byte("</form>"))
	w.Write([]byte("</body></html>"))

	return false
}

/**
 * Ask for Authorization.
 */
func HandleAuthorizationPage(ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) bool {
	return false
}

/**
 * OAuth Authorization handler.
 */
func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Authorize Handler")
	server := GetServer().GetOAuth2Manager().m_server
	if server == nil {
		fmt.Printf("ERROR: %s\n", errors.New("no OAuth2 service configure!"))
		fmt.Fprintf(w, "no OAuth2 service configure!")
		return
	}

	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {
		if !HandleLoginPage(ar, w, r) {
			return
		}
		ar.Authorized = true
		server.FinishAuthorizeRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	osin.OutputJSON(resp, w, r)
}

/**
 * Access token endpoint
 */
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Token Handler")
	server := GetServer().GetOAuth2Manager().m_server
	if server == nil {
		fmt.Printf("ERROR: %s\n", errors.New("no OAuth2 service configure!"))
		fmt.Fprintf(w, "no OAuth2 service configure!")
		return
	}

	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAccessRequest(resp, r); ar != nil {
		switch ar.Type {
		case osin.AUTHORIZATION_CODE:
			ar.Authorized = true
		case osin.REFRESH_TOKEN:
			ar.Authorized = true
		case osin.PASSWORD:
			if ar.Username == "test" && ar.Password == "test" {
				ar.Authorized = true
			}
		case osin.CLIENT_CREDENTIALS:
			ar.Authorized = true
		case osin.ASSERTION:
			if ar.AssertionType == "urn:osin.example.complete" && ar.Assertion == "osin.data" {
				ar.Authorized = true
			}
		}
		server.FinishAccessRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		resp.Output["custom_parameter"] = 19923
	}
	osin.OutputJSON(resp, w, r)
}

/**
 * Information endpoint
 */
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Info Handler")
	server := GetServer().GetOAuth2Manager().m_server
	if server == nil {
		fmt.Printf("ERROR: %s\n", errors.New("no OAuth2 service configure!"))
		fmt.Fprintf(w, "no OAuth2 service configure!")
		return
	}
	resp := server.NewResponse()
	defer resp.Close()

	if ir := server.HandleInfoRequest(resp, r); ir != nil {
		server.FinishInfoRequest(resp, r, ir)
	}
	osin.OutputJSON(resp, w, r)
}

/**
 * Authorization code.
 */
// Application destination - CODE
func AppAuthCodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("App Auth Code Handler")
	/*	r.ParseForm()

		code := r.Form.Get("code")

		w.Write([]byte("<html><body>"))
		w.Write([]byte("APP AUTH - CODE<br/>"))
		defer w.Write([]byte("</body></html>"))

		if code == "" {
			w.Write([]byte("Nothing to do"))
			return
		}

		jr := make(map[string]interface{})

		service := GetServer().GetConfigurationManager().getServiceConfigurationById("OAuth2Manager")
		port := service.GetPort()
		hostName := service.GetHostName()

		// build access code url
		aurl := fmt.Sprintf("/token?grant_type=authorization_code&client_id=1234&client_secret=aabbccdd&state=xyz&redirect_uri=%s&code=%s",
			url.QueryEscape("http://"+hostName+":"+strconv.Itoa(port)+"/appauth/code"), url.QueryEscape(code))

		// if parse, download and parse json
		if r.Form.Get("doparse") == "1" {
			port := GetServer().GetConfigurationManager().GetServerPort()
			hostName := GetServer().GetConfigurationManager().GetHostName()
			err := DownloadAccessToken(fmt.Sprintf("http://"+hostName+":"+strconv.Itoa(port)+"%s", aurl),
				&osin.BasicAuth{"1234", "aabbccdd"}, jr)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.Write([]byte("<br/>"))
			}
		}

		// show json error
		if erd, ok := jr["error"]; ok {
			w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
		}

		// show json access token
		if at, ok := jr["access_token"]; ok {
			w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
		}

		w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

		// output links
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Goto Token URL</a><br/>", aurl)))

		cururl := *r.URL
		curq := cururl.Query()
		curq.Add("doparse", "1")
		cururl.RawQuery = curq.Encode()
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Download Token</a><br/>", cururl.String())))
	*/
}
