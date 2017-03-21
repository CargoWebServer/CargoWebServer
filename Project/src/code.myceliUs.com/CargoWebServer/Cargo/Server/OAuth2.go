/**
 * That file contain the code for OAuth2 service.
 */

package Server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	//"code.myceliUs.com/Utility"
	"github.com/RangelReale/osin"
	//"github.com/RangelReale/osincli"
)

type OAuth2Manager struct {
	// Contain the list of avalable ldap servers...
	m_configsInfo *Config.OAuth2Configuration

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
	this.m_configsInfo = GetServer().GetConfigurationManager().GetOAuth2Configuration()
	if this.m_configsInfo == nil {
		// Get the default configuration.
		c := osin.NewServerConfig()
		this.m_server = osin.NewServer(c, this.m_store)

		// Save it into the entity.
		cfg := new(Config.OAuth2Configuration)
		cfg.SetId("OAuth2Config")

		cfg.SetAccessExpiration(int(c.AccessExpiration))

		cfg.SetAllowClientSecretInParams(c.AllowClientSecretInParams)
		//cfg.SetAllowedAccessTypes(c.AllowedAccessTypes)
		//cfg.SetAllowedAuthorizeTypes(c.AllowedAuthorizeTypes)
		cfg.SetAllowGetAccessRequest(c.AllowGetAccessRequest)
		for i := 0; i < len(c.AllowedAuthorizeTypes); i++ {
			if c.AllowedAuthorizeTypes[i] == osin.CODE {
				cfg.SetAllowedAuthorizeTypes("code")
			} else if c.AllowedAuthorizeTypes[i] == osin.TOKEN {
				cfg.SetAllowedAuthorizeTypes("token")
			}
		}

		for i := 0; i < len(c.AllowedAccessTypes); i++ {
			if c.AllowedAccessTypes[i] == osin.AUTHORIZATION_CODE {
				cfg.SetAllowedAccessTypes("authorization_code")
			} else if c.AllowedAccessTypes[i] == osin.REFRESH_TOKEN {
				cfg.SetAllowedAuthorizeTypes("refresh_token")
			} else if c.AllowedAccessTypes[i] == osin.PASSWORD {
				cfg.SetAllowedAuthorizeTypes("password")
			} else if c.AllowedAccessTypes[i] == osin.CLIENT_CREDENTIALS {
				cfg.SetAllowedAuthorizeTypes("client_credentials")
			} else if c.AllowedAccessTypes[i] == osin.ASSERTION {
				cfg.SetAllowedAuthorizeTypes("assertion")
			} else if c.AllowedAccessTypes[i] == osin.IMPLICIT {
				cfg.SetAllowedAuthorizeTypes("__implicit")
			}
		}

		cfg.SetAuthorizationExpiration(int(c.AuthorizationExpiration))
		cfg.SetErrorStatusCode(c.ErrorStatusCode)
		cfg.SetRedirectUriSeparator(c.RedirectUriSeparator)
		cfg.SetTokenType(c.TokenType)
		// Append to the active configuration and save the entity.
		GetServer().GetConfigurationManager().m_activeConfigurations.SetOauth2Configuration(cfg)
		GetServer().GetConfigurationManager().m_configurationEntity.SaveEntity()
	} else {
		c := osin.NewServerConfig()
		// TODO set the c value from this.m_configsInfo.
		this.m_server = osin.NewServer(c, this.m_store)
	}

	this.m_store = newOauth2Store()

}

func (this *OAuth2Manager) stop() {
	log.Println("--> Stop OAuth2Manager")
}

////////////////////////////////////////////////////////////////////////////////
// Api
////////////////////////////////////////////////////////////////////////////////

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

			// return the client.
			log.Println("Client with id ", id, " found!")
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
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	// Create a client configuration from the osin.Client.
	c := new(Config.OAuth2Client)
	c.M_id = id
	c.M_extra = client.GetUserData().([]uint8)
	c.M_secret = client.GetSecret()
	c.M_redirectUri = client.GetRedirectUri()

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
	a.SetCode(data.Code)
	a.SetExpiresIn(int64(data.ExpiresIn))
	a.SetScope(data.Scope)
	a.SetRedirectUri(data.RedirectUri)
	a.SetState(data.State)
	a.SetCreatedAt(data.CreatedAt.Unix())

	if data.UserData != nil {
		a.SetExtra(data.UserData)
	}

	// Put in the config.
	config.SetAuthorize(a)

	// Save the entity.
	configEntity.SaveEntity()

	log.Println("Save Authorize completed!")
	return nil
}

// LoadAuthorize looks up AuthorizeData by a code.
// Client information MUST be loaded together.
// Optionally can return error if expired.
func (this *OAuth2Store) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()

	var data *osin.AuthorizeData
	for i := 0; i < len(config.GetAuthorize()); i++ {
		if config.GetAuthorize()[i].GetCode() == code {
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
		}
	}

	// No data was found.
	if data == nil {
		return nil, errors.New("No authorize data found with code " + code)
	}

	// Now I will test the expiration time.
	if data.ExpireAt().Before(time.Now()) {
		return nil, errors.New("Token expired at " + data.ExpireAt().String())
	}

	log.Println("Load Authorize Data completed!")
	return data, nil
}

/**
 * Remove authorize from the db.
 */
func (this *OAuth2Store) RemoveAuthorize(code string) error {

	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	config.RemoveAuthorize(code)
	configEntity.SaveEntity()
	if err := this.RemoveExpireAtData(code); err != nil {
		return err
	}
	log.Println("Remove authorize complete!")
	return nil
}

/**
 * Save the access Data.
 */
func (this *OAuth2Store) SaveAccess(data *osin.AccessData) error {
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	access := new(Config.OAuth2Access)
	for i := 0; i < len(config.GetAccess()); i++ {
		if config.GetAccess()[i].M_accessToken == data.AccessToken {
			access = config.GetAccess()[i]
			break
		}
	}

	prev := ""
	authorizeData := &osin.AuthorizeData{}

	if data.AccessData != nil {
		prev = data.AccessData.AccessToken
	}

	if data.AuthorizeData != nil {
		authorizeData = data.AuthorizeData
	}

	var r *Config.OAuth2Refresh
	for i := 0; i < len(config.GetRefresh()); i++ {
		if config.GetRefresh()[i].GetToken() == data.RefreshToken && config.GetRefresh()[i].GetAccess().GetAccessToken() == data.AccessToken {
			r = config.GetRefresh()[i]
		}
	}

	// Here the refresh token dosent exist so i will create it.
	if r == nil {
		r = new(Config.OAuth2Refresh)
		r.SetToken(data.RefreshToken)
		r.SetAccess(access)
		config.SetRefresh(r)
	}

	if data.Client == nil {
		return errors.New("data.Client must not be nil")
	}

	// Set the client.
	c, err := this.GetClient(data.Client.GetId())
	if err != nil {
		return err
	}
	access.SetClient(c)

	// Set the authorization.
	a, err := this.LoadAuthorize(authorizeData.Code)
	if err != nil {
		return err
	}
	access.SetAuthorize(a)

	// Set other values.
	access.SetPrevious(prev)
	access.SetAccessToken(data.AccessToken)
	access.SetRefreshToken(data.RefreshToken)
	access.SetExpiresIn(data.ExpiresIn)
	access.SetScope(data.Scope)
	access.SetRedirectUri(data.RedirectUri)
	access.SetCreatedAt(data.CreatedAt)

	if data.UserData != nil {
		access.SetExtra(data.UserData)
	}

	// Add expire data.
	if err = this.AddExpireAtData(data.AccessToken, data.ExpireAt()); err != nil {
		return err
	}

	configEntity.SaveEntity()
	return nil
}

/**
 * Load the access for a given code.
 */
func (this *OAuth2Store) LoadAccess(code string) (*osin.AccessData, error) {

	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()

	var access *osin.AccessData
	for i := 0; i < len(config.GetAccess()); i++ {
		if config.GetAccess()[i].GetAccessToken() == code {
			access = new(osin.AccessData)
			access.AccessToken = code
			access.ExpiresIn = int32(config.GetAccess()[i].GetExpiresIn())
			access.Scope = config.GetAccess()[i].GetScope()
			access.RedirectUri = config.GetAccess()[i].GetRedirectUri()
			access.CreatedAt = time.Unix(int64(config.GetAccess()[i].GetCreatedAt()), 0)
			access.UserData = config.GetAccess()[i].GetExtra()

			// The refresh token
			access.RefreshToken = config.GetAccess()[i].GetAccessToken()

			// The access token
			access.AccessToken = config.GetAccess()[i].GetAccessToken()

			// Now the client
			c, err := this.GetClient(config.GetAccess()[i].GetClient().GetId())
			if err != nil {
				return nil, err
			}
			access.Client = c

			// The authorize
			auth, err := this.LoadAuthorize(config.GetAccess()[i].GetAuthorize().GetCode())
			if err != nil {
				return nil, err
			}
			access.AuthorizeData = auth
		}
	}

	return access, nil
}

/**
 * Remove an access code.
 */
func (this *OAuth2Store) RemoveAccess(code string) error {
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	config.RemoveAccess(code)
	configEntity.SaveEntity()

	// Remove the expire data.
	if err := this.RemoveExpireAtData(code); err != nil {
		return err
	}

	return nil
}

/**
 * Load the access data from it refresh code.
 */
func (this *OAuth2Store) LoadRefresh(code string) (*osin.AccessData, error) {
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()

	for i := 0; i < len(config.GetRefresh()); i++ {
		if config.GetRefresh()[i].GetToken() == code {
			return this.LoadAccess(config.GetRefresh()[i].GetAccess().GetAccessToken())
		}
	}

	return nil, errors.New("No access data was found for " + code)
}

/**
 * Remove refresh.
 */
func (this *OAuth2Store) RemoveRefresh(code string) error {
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	config.RemoveRefresh(code)
	configEntity.SaveEntity()
	return nil
}

// AddExpireAtData add info in expires table
func (s *OAuth2Store) AddExpireAtData(code string, expireAt time.Time) error {
	var expire *Config.OAuth2Expires
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	for i := 0; i < len(config.GetExpire()); i++ {
		if config.GetExpire()[i].GetToken().GetToken() == code {
			expire = config.GetExpire()[i]
		}
	}

	expire.SetExpiresAt(expireAt.Unix())
	config.SetExpire(expire)
	configEntity.SaveEntity()

	return nil
}

// RemoveExpireAtData remove info in expires table
func (s *OAuth2Store) RemoveExpireAtData(code string) error {
	config := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	configEntity, _ := GetServer().GetEntityManager().getEntityByUuid(config.GetUUID())

	for i := 0; i < len(config.GetExpire()); i++ {
		if config.GetExpire()[i].GetToken().GetToken() == code {
			config.RemoveExpire(config.GetExpire()[i])
		}
	}

	configEntity.SaveEntity()
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Auth Http handler.
////////////////////////////////////////////////////////////////////////////////
// TODO wrote the code to use the web socket here...

/**
 * That contain the authorization page. With the account manager test for
 * the password and user...
 */
func HandleLoginPage(ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) bool {

	log.Println("---------> Handle Login Page!")
	r.ParseForm()

	// Test the values...
	if r.Method == "POST" && r.Form.Get("login") == "test" && r.Form.Get("password") == "test" {
		return true
	}

	w.Write([]byte("<html><body>"))

	w.Write([]byte(fmt.Sprintf("LOGIN %s (use test/test)<br/>", ar.Client.GetId())))
	w.Write([]byte(fmt.Sprintf("<form action=\"/authorize?%s\" method=\"POST\">", r.URL.RawQuery)))

	w.Write([]byte("Login: <input type=\"text\" name=\"login\" /><br/>"))
	w.Write([]byte("Password: <input type=\"password\" name=\"password\" /><br/>"))
	w.Write([]byte("<input type=\"submit\"/>"))

	w.Write([]byte("</form>"))

	w.Write([]byte("</body></html>"))

	return false
}

/**
 * That contain the access token.
 */
func DownloadAccessToken(url string, auth *osin.BasicAuth, output map[string]interface{}) error {
	// download access token
	log.Println("-------> download access token! ", url)
	preq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	if auth != nil {
		preq.SetBasicAuth(auth.Username, auth.Password)
	}

	pclient := &http.Client{}
	presp, err := pclient.Do(preq)
	if err != nil {
		return err
	}

	if presp.StatusCode != 200 {
		return errors.New("Invalid status code")
	}

	jdec := json.NewDecoder(presp.Body)
	err = jdec.Decode(&output)
	return err
}

/**
 * OAuth Authorization handler.
 */
func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {

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
	server := GetServer().GetOAuth2Manager().m_server
	if server == nil {
		fmt.Printf("ERROR: %s\n", errors.New("no OAuth2 service configure!"))
		fmt.Fprintf(w, "no OAuth2 service configure!")
		return
	}
	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAccessRequest(resp, r); ar != nil {
		ar.Authorized = true
		server.FinishAccessRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", errors.New("no OAuth2 service configure!"))
	}
	osin.OutputJSON(resp, w, r)
}

/**
 * Information endpoint
 */
func InfoHandler(w http.ResponseWriter, r *http.Request) {
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
 * Application home endpoint
 */
func AppHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body>"))

	service := GetServer().GetConfigurationManager().getServiceConfigurationById("OAuth2Manager")
	port := service.GetPort()
	hostName := service.GetHostName()

	w.Write([]byte(fmt.Sprintf("<a href=\"/authorize?response_type=code&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Login</a><br/>", url.QueryEscape("http://"+hostName+":"+strconv.Itoa(port)+"/appauth/code"))))
	w.Write([]byte("</body></html>"))
}

/**
 * Authorization code.
 */
// Application destination - CODE
func AppAuthCodeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

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
}
