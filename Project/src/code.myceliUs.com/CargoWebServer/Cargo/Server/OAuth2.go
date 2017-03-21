/**
 * That file contain the code for OAuth2 service.
 */

package Server

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
	"github.com/RangelReale/osin"
	"github.com/RangelReale/osincli"
	"github.com/syndtr/goleveldb/leveldb"
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
	configuration := GetServer().GetConfigurationManager().GetOAuth2Configuration()
	this.m_configsInfo = configuration
	if configuration != nil {
		this.m_store = newOauth2Store(configuration)

		// Set the OAuth server configuration.
		cfg := osin.NewServerConfig()

		// Boolean values.
		//cfg.AllowGetAccessRequest = configuration.GetAllowGetAccessRequest()
		//cfg.AllowClientSecretInParams = configuration.GetAllowClientSecretInParams()

		// Expiration times.
		/*cfg.AccessExpiration = configuration.GetAccessExpiration()
		cfg.AuthorizationExpiration = configuration.GetAuthorizationExpiration()

		// Init properly...
		// cfg.AllowedAccessTypes = configuration.GetAllowedAuthorizeTypes()
		cfg.ErrorStatusCode = configuration.GetErrorStatusCode()
		cfg.RedirectUriSeparator = configuration.GetRedirectUriSeparator()
		cfg.TokenType = configuration.GetTokenType()*/

		// Test...
		cfg.AllowGetAccessRequest = true
		cfg.AllowClientSecretInParams = true

		this.m_server = osin.NewServer(cfg, this.m_store)
		log.Println("------> ", this.m_configsInfo.GetClients())
		// Here I will synchronise the clients from the configuration
		// and the datastore.
		for i := 0; i < len(this.m_configsInfo.GetClients()); i++ {
			clientInfo := this.m_configsInfo.GetClients()[i]
			client, err := this.m_store.GetClient(clientInfo.M_id)
			if err != nil {
				// Here I will append the client...
				client = new(osin.DefaultClient)
				client.(*osin.DefaultClient).Id = clientInfo.M_id
				client.(*osin.DefaultClient).Secret = clientInfo.M_secret
				client.(*osin.DefaultClient).RedirectUri = clientInfo.M_redirectUri
				client.(*osin.DefaultClient).UserData = clientInfo.M_userData
				this.m_store.SetClient(clientInfo.M_id, client)
			}
		}
	}

}

func (this *OAuth2Manager) stop() {
	log.Println("--> Stop OAuth2Manager")
	// Close level db.
	this.m_store.m_db.Close()
}

////////////////////////////////////////////////////////////////////////////////
// Api
////////////////////////////////////////////////////////////////////////////////

/**
 * That function is use to register an OAuth2 client.
 */
func (this *OAuth2Manager) RegisterDefaultClient(messageId string, sessionId string, id string, secret string, redirectUri string, data interface{}) {

	client := new(osin.DefaultClient)
	client.Id = id
	client.Secret = secret
	client.RedirectUri = redirectUri
	client.UserData = data

	// Here I will try to store the clients.
	err := this.m_store.SetClient(id, client)

	if err != nil {
		// Send the error message in that case...
		cargoError := NewError(Utility.FileLine(), FILE_DELETE_ERROR, SERVER_ERROR_CODE, errors.New("Failed to register client with id "+id))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
	}
}

////////////////////////////////////////////////////////////////////////////////
// OAuth2 Store
////////////////////////////////////////////////////////////////////////////////
type OAuth2Store struct {
	/** The runtime data store. **/
	m_db *leveldb.DB

	/**
	 * Use to protected the entitiesMap access...
	 */
	sync.RWMutex
}

// Create and intialyse the OAuth2 Store.
func newOauth2Store(config *Config.OAuth2Configuration) *OAuth2Store {

	// Register type to be saved in level db.
	Utility.RegisterType((*osin.Client)(nil))
	Utility.RegisterType((*osin.AuthorizeData)(nil))
	Utility.RegisterType((*osin.AccessData)(nil))

	var err error
	store := new(OAuth2Store)
	storePath := GetServer().GetConfigurationManager().GetDataPath() + "/" + config.GetId()
	// The path will be taken from the store configuration.
	store.m_db, err = leveldb.OpenFile(storePath, nil)

	if err != nil {
		log.Fatal("open:", err)
	}

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
	this.Lock()
	defer this.Unlock()
	/*log.Printf("GetClient: %s\n", id)

	values, err := this.m_db.Get([]byte(id), nil)
	if err != nil {
		return nil, err
	}

	// So here I will try to instansiate the client.
	var client osin.Client
	dec := gob.NewDecoder(bytes.NewReader(values))
	dec.Decode(&client)*/
	// create client
	cliconfig := &osincli.ClientConfig{
		ClientId:     "1234",
		ClientSecret: "aabbccdd",
		AuthorizeUrl: "http://localhost:9393/authorize",
		TokenUrl:     "http://localhost:9393/token",
		RedirectUrl:  "http://localhost:9393/appauth",
	}
	client, err := osincli.NewClient(cliconfig)
	if err != nil {
		panic(err)
	}
	return client, nil
}

/**
 * Set the client value.
 */
func (this *OAuth2Store) SetClient(id string, client osin.Client) error {
	this.Lock()
	defer this.Unlock()

	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	err := enc.Encode(client)

	err = this.m_db.Put([]byte(id), m.Bytes(), nil)
	if err != nil {
		log.Println("Client encode:", err)
		return err
	}

	return nil
}

/**
 * Save a given autorization.
 */
func (this *OAuth2Store) SaveAuthorize(data *osin.AuthorizeData) error {
	this.Lock()
	defer this.Unlock()
	log.Printf("SaveAuthorize: %s\n", data.Code)
	log.Println("==----------> 260: ", data)
	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	err := enc.Encode(data)

	err = this.m_db.Put([]byte(data.Code), m.Bytes(), nil)
	if err != nil {
		log.Println("Client encode:", err)
		return err
	}

	return nil
}

/**
 * Load an authorization by it code.
 */
func (this *OAuth2Store) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	this.Lock()
	defer this.Unlock()
	log.Printf("LoadAuthorize: %s\n", code)

	values, err := this.m_db.Get([]byte(code), nil)
	if err != nil {
		return nil, err
	}

	// So here I will try to instansiate the client.
	authorizeData := new(osin.AuthorizeData)
	dec := gob.NewDecoder(bytes.NewReader(values))
	dec.Decode(authorizeData)

	log.Println("==----------> 292: ", authorizeData)
	return authorizeData, nil
}

/**
 * Remove authorize from the db.
 */
func (this *OAuth2Store) RemoveAuthorize(code string) error {
	this.Lock()
	defer this.Unlock()
	log.Printf("RemoveAuthorize: %s\n", code)
	return this.m_db.Delete([]byte(code), nil)
}

/**
 * Save the access Data.
 */
func (this *OAuth2Store) SaveAccess(data *osin.AccessData) error {
	this.Lock()
	defer this.Unlock()
	log.Printf("SaveAccess: %s\n", data.AccessToken)

	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	err := enc.Encode(data)

	err = this.m_db.Put([]byte(data.AccessToken), m.Bytes(), nil)
	if err != nil {
		return err
	}

	// keep the refresh token.
	err = this.m_db.Put([]byte(data.RefreshToken), []byte(data.AccessToken), nil)
	if err != nil {
		return err
	}

	return nil
}

/**
 * Load the access for a given code.
 */
func (this *OAuth2Store) LoadAccess(code string) (*osin.AccessData, error) {
	this.Lock()
	defer this.Unlock()
	log.Printf("LoadAccess: %s\n", code)

	values, err := this.m_db.Get([]byte(code), nil)
	if err != nil {
		return nil, err
	}

	// So here I will try to instansiate the client.
	accessData := new(osin.AccessData)
	dec := gob.NewDecoder(bytes.NewReader(values))
	dec.Decode(accessData)

	return accessData, nil
}

/**
 * Remove an access code.
 */
func (this *OAuth2Store) RemoveAccess(code string) error {
	this.Lock()
	defer this.Unlock()
	log.Printf("RemoveAccess: %s\n", code)
	return this.m_db.Delete([]byte(code), nil)
}

/**
 * Load the access data from it refresh code.
 */
func (this *OAuth2Store) LoadRefresh(code string) (*osin.AccessData, error) {

	this.Lock()
	defer this.Unlock()
	log.Printf("LoadRefresh: %s\n", code)

	values, err := this.m_db.Get([]byte(code), nil)
	if err != nil {
		return nil, err
	}

	// So here I will try to instansiate the client.
	accessData := new(osin.AccessData)
	dec := gob.NewDecoder(bytes.NewReader(values))
	dec.Decode(accessData)

	return accessData, nil
}

func (this *OAuth2Store) RemoveRefresh(code string) error {
	this.Lock()
	defer this.Unlock()
	log.Printf("RemoveRefresh: %s\n", code)
	return this.m_db.Delete([]byte(code), nil)
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
