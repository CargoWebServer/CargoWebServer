/**
 * That file contain the code for OAuth2 service.
 */

package Server

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	//"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
	"github.com/RangelReale/osin"
	"gopkg.in/square/go-jose.v1"
)

// variable use betheewen the http handlers and the OAuth service.
var (
	// The manager.
	oauth2Manager *OAuth2Manager

	// The key
	// TODO set in config???
	privateKeyBytes = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA4f5wg5l2hKsTeNem/V41fGnJm6gOdrj8ym3rFkEU/wT8RDtn
SgFEZOQpHEgQ7JL38xUfU0Y3g6aYw9QT0hJ7mCpz9Er5qLaMXJwZxzHzAahlfA0i
cqabvJOMvQtzD6uQv6wPEyZtDTWiQi9AXwBpHssPnpYGIn20ZZuNlX2BrClciHhC
PUIIZOQn/MmqTD31jSyjoQoV7MhhMTATKJx2XrHhR+1DcKJzQBSTAGnpYVaqpsAR
ap+nwRipr3nUTuxyGohBTSmjJ2usSeQXHI3bODIRe1AuTyHceAbewn8b462yEWKA
Rdpd9AjQW5SIVPfdsz5B6GlYQ5LdYKtznTuy7wIDAQABAoIBAQCwia1k7+2oZ2d3
n6agCAbqIE1QXfCmh41ZqJHbOY3oRQG3X1wpcGH4Gk+O+zDVTV2JszdcOt7E5dAy
MaomETAhRxB7hlIOnEN7WKm+dGNrKRvV0wDU5ReFMRHg31/Lnu8c+5BvGjZX+ky9
POIhFFYJqwCRlopGSUIxmVj5rSgtzk3iWOQXr+ah1bjEXvlxDOWkHN6YfpV5ThdE
KdBIPGEVqa63r9n2h+qazKrtiRqJqGnOrHzOECYbRFYhexsNFz7YT02xdfSHn7gM
IvabDDP/Qp0PjE1jdouiMaFHYnLBbgvlnZW9yuVf/rpXTUq/njxIXMmvmEyyvSDn
FcFikB8pAoGBAPF77hK4m3/rdGT7X8a/gwvZ2R121aBcdPwEaUhvj/36dx596zvY
mEOjrWfZhF083/nYWE2kVquj2wjs+otCLfifEEgXcVPTnEOPO9Zg3uNSL0nNQghj
FuD3iGLTUBCtM66oTe0jLSslHe8gLGEQqyMzHOzYxNqibxcOZIe8Qt0NAoGBAO+U
I5+XWjWEgDmvyC3TrOSf/KCGjtu0TSv30ipv27bDLMrpvPmD/5lpptTFwcxvVhCs
2b+chCjlghFSWFbBULBrfci2FtliClOVMYrlNBdUSJhf3aYSG2Doe6Bgt1n2CpNn
/iu37Y3NfemZBJA7hNl4dYe+f+uzM87cdQ214+jrAoGAXA0XxX8ll2+ToOLJsaNT
OvNB9h9Uc5qK5X5w+7G7O998BN2PC/MWp8H+2fVqpXgNENpNXttkRm1hk1dych86
EunfdPuqsX+as44oCyJGFHVBnWpm33eWQw9YqANRI+pCJzP08I5WK3osnPiwshd+
hR54yjgfYhBFNI7B95PmEQkCgYBzFSz7h1+s34Ycr8SvxsOBWxymG5zaCsUbPsL0
4aCgLScCHb9J+E86aVbbVFdglYa5Id7DPTL61ixhl7WZjujspeXZGSbmq0Kcnckb
mDgqkLECiOJW2NHP/j0McAkDLL4tysF8TLDO8gvuvzNC+WQ6drO2ThrypLVZQ+ry
eBIPmwKBgEZxhqa0gVvHQG/7Od69KWj4eJP28kq13RhKay8JOoN0vPmspXJo1HY3
CKuHRG+AP579dncdUnOMvfXOtkdM4vk0+hWASBQzM9xzVcztCa+koAugjVaLS9A+
9uQoqEeVNTckxx0S2bYevRy7hGQmUJTyQm3j1zEUR5jpdbL83Fbq
-----END RSA PRIVATE KEY-----`)
)

// The ID Token represents a JWT passed to the client as part of the token response.
//
// https://openid.net/specs/openid-connect-core-1_0.html#IDToken
type IDToken struct {
	Issuer     string `json:"iss"`
	UserID     string `json:"sub"`
	ClientID   string `json:"aud"`
	Expiration int64  `json:"exp"`
	IssuedAt   int64  `json:"iat"`

	Nonce string `json:"nonce,omitempty"` // Non-manditory fields MUST be "omitempty"

	// Custom claims supported by this server.
	//
	// See: https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims

	Email         string `json:"email,omitempty"`
	EmailVerified *bool  `json:"email_verified,omitempty"`

	Name       string `json:"name,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	GivenName  string `json:"given_name,omitempty"`
	Locale     string `json:"locale,omitempty"`
}

/**
 * The OAuth2 Server.
 */
type OAuth2Manager struct {

	// the data stores.
	m_store *OAuth2Store

	// the oauth sever.
	m_server *osin.Server

	// openId authentication.
	m_jwtSigner  jose.Signer
	m_publicKeys *jose.JsonWebKeySet
}

func (this *Server) GetOAuth2Manager() *OAuth2Manager {
	if oauth2Manager == nil {
		oauth2Manager = newOAuth2Manager()
	}
	return oauth2Manager
}

func newOAuth2Manager() *OAuth2Manager {

	// The oauth manager.
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

	// Load signing key.
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		log.Fatalf("no private key found")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("failed to parse key: %v", err)
	}

	// Configure jwtSigner and public keys.
	privateKey := &jose.JsonWebKey{
		Key:       key,
		Algorithm: "RS256",
		Use:       "sig",
		KeyID:     "1", // KeyID should use the key thumbprint.
	}

	this.m_jwtSigner, err = jose.NewSigner(jose.RS256, privateKey)
	if err != nil {
		log.Fatalf("failed to create jwtSigner: %v", err)
	}

	this.m_publicKeys = &jose.JsonWebKeySet{
		Keys: []jose.JsonWebKey{
			jose.JsonWebKey{Key: &key.PublicKey,
				Algorithm: "RS256",
				Use:       "sig",
				KeyID:     "1",
			},
		},
	}
}

func (this *OAuth2Manager) getId() string {
	return "OAuth2Manager"
}

func (this *OAuth2Manager) start() {
	log.Println("--> Start OAuth2Manager")

	// Here I will intialyse configurations.
	configEntity := GetServer().GetConfigurationManager().getActiveConfigurationsEntity()
	cfg := configEntity.GetObject().(*Config.Configurations).GetOauth2Configuration()
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
		sconfig.AccessExpiration = 60 // Test only...
		cfg.SetAccessExpiration(int64(sconfig.AccessExpiration))
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

		// Create the new configuration entity.
		GetServer().GetEntityManager().createEntity(configEntity.GetUuid(), "M_oauth2Configuration", "Config.OAuth2Configuration", cfg.GetId(), cfg)

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

		// Cleanup
		this.cleanup()
	}

	// Start the oauth service.
	this.m_store = newOauth2Store()
	this.m_server = osin.NewServer(sconfig, this.m_store)
}

func (this *OAuth2Manager) stop() {
	log.Println("--> Stop OAuth2Manager")
}

/**
 * That function remove expire access and authorization and renew access
 * if refresh exist.
 */
func (this *OAuth2Manager) cleanup() {
	config := GetServer().GetConfigurationManager().getActiveConfigurationsEntity().GetObject().(*Config.Configurations).GetOauth2Configuration()
	// First of all I will renew the access...
	for i := 0; i < len(config.GetAccess()); i++ {
		access := config.GetAccess()[i]
		expirationTime := time.Unix(access.GetCreatedAt(), 0).Add(time.Duration(access.GetExpiresIn()) * time.Second)
		if expirationTime.Before(time.Now()) {
			if access.GetRefreshToken() != nil {
				accessEntity, _ := GetServer().GetEntityManager().getEntityByUuid(access.UUID)
				// Reset the creation time instead of delete it and recreated it...
				access.SetCreatedAt(time.Now().Unix())
				accessEntity.SaveEntity() // update it

				// Now it expire time.
				expireEntity, _ := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Expires", access.GetId())
				expireEntity.GetObject().(*Config.OAuth2Expires).SetExpiresAt(time.Unix(access.GetCreatedAt(), 0).Add(time.Duration(access.GetExpiresIn()) * time.Second).Unix())
				expireEntity.SaveEntity()
			}
		}
	}

	// Here I will remove all expired authorization and access.
	for i := 0; i < len(config.GetExpire()); i++ {
		expireTime := time.Unix(config.GetExpire()[i].GetExpiresAt(), 0)
		if expireTime.Before(time.Now()) {
			// I that case the value must be remove expired values...
			this.m_store.RemoveAccess(config.GetExpire()[i].GetId())
			this.m_store.RemoveAuthorize(config.GetExpire()[i].GetId())
		} else {
			setCodeExpiration(config.GetExpire()[i].GetId(), expireTime.Sub(time.Now()))
		}
	}

	// Now I will remove all refresh without access...
	for i := 0; i < len(config.GetRefresh()); i++ {
		refresh := config.GetRefresh()[i]
		if refresh.GetAccess() == nil {
			refreshEntity, _ := GetServer().GetEntityManager().getEntityByUuid(refresh.GetUUID())
			GetServer().GetEntityManager().deleteEntity(refreshEntity)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// The api
////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////// OpenID ///////////////////////////////////

/**
 * handleDiscovery returns the OpenID Connect discovery object, allowing clients
 * to discover OAuth2 resources.
 */
func DiscoveryHandler(w http.ResponseWriter, r *http.Request) {
	// For other example see: https://accounts.google.com/.well-known/openid-configuration
	config := GetServer().GetConfigurationManager().getServiceConfigurationById(GetServer().GetOAuth2Manager().getId())
	issuer := "https://" + config.GetHostName() + ":" + strconv.Itoa(config.GetPort())
	data := map[string]interface{}{
		"issuer":                                issuer,
		"authorization_endpoint":                issuer + "/authorize",
		"token_endpoint":                        issuer + "/token",
		"jwks_uri":                              issuer + "/publickeys",
		"response_types_supported":              []string{"code"},
		"subject_types_supported":               []string{"public"},
		"id_token_signing_alg_values_supported": []string{"RS256"},
		"scopes_supported":                      []string{"openid", "email", "profile"},
		"token_endpoint_auth_methods_supported": []string{"client_secret_basic"},
		"claims_supported": []string{
			"aud", "email", "email_verified", "exp",
			"family_name", "given_name", "iat", "iss",
			"locale", "name", "sub",
		},
	}

	raw, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("failed to marshal data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(raw)))
	w.Write(raw)
}

/**
 * handlePublicKeys publishes the public part of this server's signing keys.
 * This allows clients to verify the signature of ID Tokens.
 */
func PublicKeysHandler(w http.ResponseWriter, r *http.Request) {
	raw, err := json.MarshalIndent(GetServer().GetOAuth2Manager().m_publicKeys, "", "  ")
	if err != nil {
		log.Printf("failed to marshal data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(raw)))
	w.Write(raw)
}

///////////////////////////////////// OAuth2 ///////////////////////////////////
/**
 * That function is use to get a given ressource for a given client.
 * The client id is the id define in the configuration
 * The scope are the scope of the ressources, ex. public_profile, email...
 * The query is an http query from various api like facebook graph api.
 * will start an authorization process if nothing is found.
 */
func (this *OAuth2Manager) GetResource(clientId string, scope string, query string, messageId string, sessionId string) interface{} {

	// So here the first thing to do is to get the client.
	config := GetServer().GetConfigurationManager().getActiveConfigurationsEntity().GetObject().(*Config.Configurations).GetOauth2Configuration()

	var client *Config.OAuth2Client
	for i := 0; i < len(config.GetClients()); i++ {
		if clientId == config.GetClients()[i].GetId() {
			client = config.GetClients()[i]
		}
	}

	if client == nil {
		// Repport the error
		errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Fail to execute query: "+query))
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

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
		// TODO give support to other token type.
		result, err := DownloadRessource(query, access.GetId(), "Bearer")
		if err == nil {
			return result
		} else {
			errObj := NewError(Utility.FileLine(), RESSOURCE_NOT_FOUND_ERROR, SERVER_ERROR_CODE, errors.New("No client was found with id "+clientId))
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
	} else {

		// No access was found so here I will initiated the OAuth process...
		// To do so I will create the href where the user will be ask to
		// authorize the client application to access ressources.
		var authorizationLnk = client.GetAuthorizationUri()
		authorizationLnk += "?response_type=code&client_id=" + client.GetId()
		authorizationLnk += "&state=" + messageId + ":" + sessionId + ":" + clientId + "&scope=" + scope
		authorizationLnk += "&redirect_uri=" + client.GetRedirectUri()

		// I will create the request and send it to the client...
		id := Utility.RandomUUID()

		// Here if there is no user logged for the given session I will send an authentication request.
		var method string
		method = "OAuth2Authorization"

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
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Helper function
////////////////////////////////////////////////////////////////////////////////

/**
 * Clear expiring authorization and (access/refresh)
 */
func clearCodeExpired(code string) {

	// Remove the expire
	expireUuid := ConfigOAuth2ExpiresExists(code)
	if len(expireUuid) > 0 {
		expireEntity, _ := GetServer().GetEntityManager().getEntityByUuid(expireUuid)
		GetServer().GetEntityManager().deleteEntity(expireEntity)
	}

	// Remove the authorization
	authorizationUuid := ConfigOAuth2AuthorizeExists(code)
	if len(authorizationUuid) > 0 {
		authorizationEntity, _ := GetServer().GetEntityManager().getEntityByUuid(authorizationUuid)
		GetServer().GetEntityManager().deleteEntity(authorizationEntity)
	}

	// Remove the access
	accessUuid := ConfigOAuth2AccessExists(code)
	if len(accessUuid) > 0 {
		accessEntity, _ := GetServer().GetEntityManager().getEntityByUuid(accessUuid)
		GetServer().GetEntityManager().deleteEntity(accessEntity)
	}

}

/**
 * Use a timer to execute clearExpiredCode when it can at end...
 */
func setCodeExpiration(code string, duration time.Duration) {
	// Create a closure and wrap the code.
	f := func(code string) func() {
		return func() {

			// In case of access token... Refresh it if it can.
			entity, errObj := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Access", code)
			if errObj == nil {
				access := entity.GetObject().(*Config.OAuth2Access)
				if access.GetRefreshToken() != nil {
					var state string
					createAccessToken("refresh_token", access.GetClient(), state, access.GetAuthorize(), access.GetRefreshToken().GetId())
				}
			}

			// Remove the old access...
			clearCodeExpired(code)
		}
	}(code)

	// The function will be call after the duration.
	time.AfterFunc(duration, f)
}

/**
 * AddExpireAtData add info in expires table
 */
func addExpireAtData(code string, expireAt time.Time) error {

	configEntity := GetServer().GetConfigurationManager().getOAuthConfigurationEntity()
	expireEntity, _ := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Expires", code)

	var expire *Config.OAuth2Expires
	if expireEntity == nil {
		expire = new(Config.OAuth2Expires)
		expire.SetId(code)

		// append to config.
		expireEntity, _ = GetServer().GetEntityManager().createEntity(configEntity.GetUuid(), "M_expire", "Config.OAuth2Expires", expire.GetId(), expire)

	} else {
		expire = expireEntity.GetObject().(*Config.OAuth2Expires)
	}

	// Set the date.
	expire.SetExpiresAt(expireAt.Unix())
	expireEntity.SaveEntity()

	// Start the timer.
	duration := expireAt.Sub(time.Now())

	// Set it expiration function.
	setCodeExpiration(code, duration)

	return nil
}

/**
 * Create Access token from refresh_token or authorization_code.
 */
func createAccessToken(grantType string, client *Config.OAuth2Client, state string, authorizationCode string, refreshToken string) {
	// The map that will contain the results
	jr := make(map[string]interface{})

	// build access code url
	var accessTokenQuery *url.URL
	accessTokenQuery, _ = url.Parse(client.GetTokenUri())
	parameters := url.Values{}
	parameters.Add("grant_type", grantType)
	parameters.Add("client_id", client.GetId())
	parameters.Add("client_secret", client.GetSecret())
	parameters.Add("state", state)
	parameters.Add("redirect_uri", client.GetRedirectUri())

	if grantType == "refresh_token" {
		if len(refreshToken) > 0 {
			parameters.Add("refresh_token", refreshToken)
		}
	} else if grantType == "authorization_code" {
		parameters.Add("code", authorizationCode)
	}

	accessTokenQuery.RawQuery = parameters.Encode()

	// if parse, download and parse json
	err := DownloadAccessToken(accessTokenQuery.String(),
		&osin.BasicAuth{client.GetId(), client.GetSecret()}, jr)

	if err != nil {
		fmt.Printf("refresh token failed with '%s'\n", err)
	} else {
		if jr["access_token"] != nil {
			// Here I will save the new access token.
			configEntity := GetServer().GetConfigurationManager().getOAuthConfigurationEntity()

			// Here I will create a new access token.
			var access *Config.OAuth2Access
			var accessEntity *Config_OAuth2AccessEntity
			entity, errObj := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Access", jr["access_token"].(string))
			if errObj != nil {
				// Here I will create a new access from json data.
				access = new(Config.OAuth2Access)
				// Set the id
				access.SetId(jr["access_token"].(string))

				// set the access uuid
				entity, _ = GetServer().GetEntityManager().createEntity(configEntity.GetUuid(), "M_access", "Config.OAuth2Access", access.GetId(), access)
				accessEntity = entity.(*Config_OAuth2AccessEntity)

				// Set the creation time.
				access.SetCreatedAt(time.Now().Unix())
				// Set the expiration delay.
				access.SetExpiresIn(int64(jr["expires_in"].(float64)))
				// Set it scope.
				if jr["scope"] != nil {
					access.SetScope(jr["scope"].(string))
				}

				/**
				// Set the custom parameters in the extra field.
				extra, err := json.Marshal(jr["custom_parameter"])
				if err == nil {
					accessToken.SetExtra(extra) // Set as json struct...
				}
				*/

				// Set the expiration...
				expirationTime := time.Unix(access.GetCreatedAt(), 0).Add(time.Duration(access.GetExpiresIn()) * time.Second)

				// Add the expire time.
				addExpireAtData(access.GetId(), expirationTime)

				// Set the client.
				access.SetClient(client)

				// Set the authorization code.
				access.SetAuthorize(authorizationCode)

				// If authorization object are found locally...
				authorizationEntity, errObj := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Authorize", authorizationCode)

				if errObj == nil {
					authorization := authorizationEntity.GetObject().(*Config.OAuth2Authorize)
					if len(authorization.GetRedirectUri()) > 0 {
						access.SetRedirectUri(authorization.GetRedirectUri())
					}
				}

				// Now the refresh token if there some.
				if jr["refresh_token"] != nil {
					// Here I will create the refresh token.
					refresh := new(Config.OAuth2Refresh)
					refresh.SetId(jr["refresh_token"].(string))
					refresh.SetAccess(access)

					// Set into it parent.
					GetServer().GetEntityManager().createEntity(configEntity.GetUuid(), "M_refresh", "Config.OAuth2Refresh", refresh.GetId(), refresh)

					// Set the access
					access.SetRefreshToken(refresh)

					// Save the entity with it refresh token object.
					accessEntity.SaveEntity()
				}

				// Save the new access token.
				configEntity.SaveEntity()
			}

			/** Nothing to do if the access is existing. **/

		} else if jr["error"] != nil {
			log.Println("Error found: ", jr["error"].(string))
		}
	}
}

/**
 * Download an access token.
 */
func DownloadAccessToken(url string, auth *osin.BasicAuth, output map[string]interface{}) error {
	// download access token
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
 * Download ressource specify with a given query.
 */
func DownloadRessource(query string, accessToken string, tokenType string) (map[string]interface{}, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", query, nil)
	req.Header.Add("Authorization", tokenType+" "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		output := make(map[string]interface{}, 0)
		json.Unmarshal(bodyBytes, &output)
		return output, nil
	}

	return nil, err
}

////////////////////////////////////////////////////////////////////////////////
// Auth Http handler.
////////////////////////////////////////////////////////////////////////////////

/**
 * If the use is not logged...
 */
func HandleAuthenticationPage(ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) bool {

	r.ParseForm()
	user := r.Form.Get("login")
	pwd := r.Form.Get("password")
	state := r.URL.Query()["state"][0]

	var sessionId string
	var messageId string
	if len(strings.Split(state, ":")) == 3 {
		messageId = strings.Split(state, ":")[0]
		sessionId = strings.Split(state, ":")[1]
	}

	// If the user is already logged in i will return true.
	if len(sessionId) > 0 {
		if GetServer().GetSessionManager().GetActiveSessionById(sessionId) != nil {
			if GetServer().GetSessionManager().GetActiveSessionById(sessionId).GetAccountPtr() != nil {
				return true
			}
		}
	}

	// Here I will authenticate the user...
	if r.Method == "POST" && len(sessionId) > 0 && len(messageId) > 0 {
		session := GetServer().GetSessionManager().Login(user, pwd, "", messageId, sessionId)
		if session != nil {
			return true
		} else {
			return false
		}
	}

	w.Write([]byte("<html><body>"))

	// if the user is no logged...
	w.Write([]byte(fmt.Sprintf("<form class='oauth2-form' action=\"/authorize?%s\" method=\"POST\">", r.URL.RawQuery)))
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

	// Here I will
	r.ParseForm()
	answer := r.FormValue("submitbutton")

	if answer == "Yes" {
		// Here the user accept the access to the ressource
		return true
	} else if answer == "No" {
		// Here The user refuse the acess to the ressource.
		// Here I will write the response that specifie that the user refuse
		// the authorization to the ressource.
		return false
	}

	// Here the user is logged and he need to ask if he give permission to
	// the request.
	w.Write([]byte("<html><body>"))
	w.Write([]byte(fmt.Sprintf("<form class='oauth2-form' action=\"/authorize?%s\" method=\"POST\">", r.URL.RawQuery)))
	w.Write([]byte("<span>Did you accept?</span></br>"))
	w.Write([]byte("<input type=\"submit\" name=\"submitbutton\" value=\"Yes\"/>"))
	w.Write([]byte("<input type=\"submit\" name=\"submitbutton\" value=\"No\"/>"))
	w.Write([]byte("</form>"))
	w.Write([]byte("</body></html>"))

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
		if !HandleAuthenticationPage(ar, w, r) {
			return
		}
		if !HandleAuthorizationPage(ar, w, r) {
			if len(r.FormValue("submitbutton")) == 0 {
				// No answer was given yet...
				return
			} else {
				// I that case the user refuse the authorization.
				ar.Authorized = false
				server.FinishAuthorizeRequest(resp, r, ar)
			}
		}
		// The user give the authorization.
		ar.Authorized = true

		// OpenId part.
		scopes := make(map[string]bool)
		for _, s := range strings.Fields(ar.Scope) {
			scopes[s] = true
		}

		// If the "openid" connect scope is specified, attach an ID Token to the
		// authorization response.
		//
		// The ID Token will be serialized and signed during the code for token exchange.
		if scopes["openid"] {
			config := GetServer().GetConfigurationManager().getServiceConfigurationById(GetServer().GetOAuth2Manager().getId())
			issuer := "https://" + config.GetHostName() + ":" + strconv.Itoa(config.GetPort())

			// These values would be tied to the end user authorizing the client.
			now := time.Now()
			idToken := IDToken{
				Issuer:     issuer,
				UserID:     "",
				ClientID:   ar.Client.GetId(),
				Expiration: now.Add(time.Hour).Unix(),
				IssuedAt:   now.Unix(),
				Nonce:      r.URL.Query().Get("nonce"),
			}

			// The state contain the messageId:sessionId:clientId
			state := r.URL.Query()["state"][0]
			var sessionId string
			if len(strings.Split(state, ":")) == 3 {
				sessionId = strings.Split(state, ":")[1]
			}

			// From the session I will retreive the user session->account->user
			session := GetServer().GetSessionManager().GetActiveSessionById(sessionId)

			// If the scope contain a profile
			if scopes["profile"] {
				idToken.UserID = session.GetAccountPtr().GetId()
				idToken.GivenName = session.GetAccountPtr().GetName()
				if session.GetAccountPtr().GetUserRef() != nil {
					idToken.Name = session.GetAccountPtr().GetUserRef().GetFirstName()
					idToken.FamilyName = session.GetAccountPtr().GetUserRef().GetLastName()
				}
				idToken.Locale = "us"
			}

			// Now if the scope contain a email...
			if scopes["email"] {
				t := true
				idToken.Email = session.GetAccountPtr().GetEmail()
				idToken.EmailVerified = &t
			}

			// NOTE: The storage must be able to encode and decode this object.
			ar.UserData = &idToken
		}
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
 * This is the client redirect handler.
 */
func AppAuthCodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("App Auth Code Handler")
	r.ParseForm()

	errorCode := r.Form.Get("error")
	state := r.Form.Get("state")
	// Send authentication end message...
	var sessionId string
	var clientId string
	if len(strings.Split(state, ":")) == 3 {
		sessionId = strings.Split(state, ":")[1]
		clientId = strings.Split(state, ":")[2]
	}

	// I will get a reference to the client who generate the request.
	clientEntity, _ := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Client", clientId)
	client := clientEntity.GetObject().(*Config.OAuth2Client)

	if len(errorCode) != 0 {
		// The authorization fail!
		errorDescription := r.Form.Get("error_description")
		log.Println("error description: ", errorDescription)
	} else {
		code := r.Form.Get("code")
		createAccessToken("authorization_code", client, state, code, "")
	}

	// The user take it decision so I will send the end authentication request.
	// I will create the request and send it to the client...
	id := Utility.RandomUUID()

	// Here if there is no user logged for the given session I will send an authentication request.
	var method string
	method = "OAuth2AuthorizationEnd"
	params := make([]*MessageData, 0)
	to := make([]connection, 1)
	to[0] = GetServer().getConnectionById(sessionId)
	oauth2AuthorizeEnd, err := NewRequestMessage(id, method, params, to, nil, nil, nil)
	if err == nil {
		// Send the request.
		GetServer().GetProcessor().m_pendingRequestChannel <- oauth2AuthorizeEnd
	}
}

////////////////////////////////////////////////////////////////////////////////
// OAuth2 Store Implementation.
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
	clientEntity, errObj := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Client", id)

	if errObj != nil {
		return nil, errors.New("No client found with id " + id)
	}

	// Set the client.
	client := clientEntity.GetObject().(*Config.OAuth2Client)

	// Create the corresponding client.
	c := new(osin.DefaultClient)
	c.Id = client.M_id
	c.Secret = client.M_secret
	c.RedirectUri = client.M_redirectUri
	c.UserData = client.M_extra
	return c, nil
}

/**
 * Set the client value.
 */
func (this *OAuth2Store) SetClient(id string, client osin.Client) error {
	log.Println("Set Client")
	// The configuration.
	configEntity := GetServer().GetConfigurationManager().getOAuthConfigurationEntity()

	// Create a client configuration from the osin.Client.
	c := new(Config.OAuth2Client)
	c.M_id = id
	c.M_extra = client.GetUserData().([]uint8)
	c.M_secret = client.GetSecret()
	c.M_redirectUri = client.GetRedirectUri()

	// append a new client.
	GetServer().GetEntityManager().createEntity(configEntity.GetUuid(), "M_client", "Config.OAuth2Client", c.GetId(), c)

	log.Println("Client with id ", id, " was save!")
	return nil
}

/**
 * Save a given autorization.
 */
func (this *OAuth2Store) SaveAuthorize(data *osin.AuthorizeData) error {
	log.Println("Save Authorize")

	// Get the config entity
	configEntity := GetServer().GetConfigurationManager().getOAuthConfigurationEntity()

	a := new(Config.OAuth2Authorize)

	// Set the client.
	c, errObj := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Client", data.Client.GetId())

	if errObj != nil {
		return errors.New("No client found with id " + data.Client.GetId())
	}

	// Set the client.
	a.SetClient(c.GetObject())

	// Set the value from the data.
	a.SetId(data.Code)
	a.SetExpiresIn(int64(data.ExpiresIn))
	a.SetScope(data.Scope)
	a.SetRedirectUri(data.RedirectUri)
	a.SetState(data.State)
	a.SetCreatedAt(data.CreatedAt.Unix())

	// Save the id token if found.
	if data.UserData != nil {
		idToken := this.saveIdToken(data.UserData.(*IDToken))
		// Set into the user user
		a.SetUserData(idToken)
	}

	// append a new Authorize.
	GetServer().GetEntityManager().createEntity(configEntity.GetUuid(), "M_authorize", "Config.OAuth2Authorize", a.GetId(), a)

	// Add expire data.
	if err := addExpireAtData(data.Code, data.ExpireAt()); err != nil {
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
	authorizeEntity, errObj := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Authorize", code)
	if errObj != nil {
		// No data was found.
		return nil, errors.New("No authorize data found with code " + code)
	}

	// Get the object.
	authorize := authorizeEntity.GetObject().(*Config.OAuth2Authorize)

	var data *osin.AuthorizeData
	data = new(osin.AuthorizeData)
	data.Code = code
	data.ExpiresIn = int32(authorize.GetExpiresIn())
	data.Scope = authorize.GetScope()
	data.RedirectUri = authorize.GetRedirectUri()
	data.State = authorize.GetState()
	data.CreatedAt = time.Unix(authorize.GetCreatedAt(), 0)

	// set the user data here.
	if authorize.GetUserData() != nil {
		data.UserData = this.loadIdToken(authorize.GetUserData())
	}

	c, err := this.GetClient(authorize.GetClient().GetId())
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

/**
 * Remove authorize from the db.
 */
func (this *OAuth2Store) RemoveAuthorize(code string) error {
	log.Println("Remove Authorize", code)
	uuid := ConfigOAuth2AuthorizeExists(code)
	if len(uuid) > 0 {
		entity, _ := GetServer().GetEntityManager().getEntityByUuid(uuid)
		GetServer().GetEntityManager().deleteEntity(entity)

		// Remove the related expire code if there one.
		uuid = ConfigOAuth2ExpiresExists(code)
		if len(uuid) > 0 {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(uuid)
			GetServer().GetEntityManager().deleteEntity(entity)
		}
		return nil
	}

	return errors.New("No authorization with code " + code + " was found!")
}

/**
 * Load a given id token.
 */
func (this *OAuth2Store) loadIdToken(idToken *Config.OAuth2IdToken) *IDToken {
	it := new(IDToken)

	it.ClientID = idToken.GetClient().GetId()
	it.Email = idToken.GetEmail()
	emailVerified := idToken.GetEmailVerified()
	it.EmailVerified = &emailVerified
	it.Expiration = idToken.GetExpiration()
	it.FamilyName = idToken.GetFamilyName()
	it.GivenName = idToken.GetGivenName()
	it.IssuedAt = idToken.GetIssuedAt()
	it.Issuer = idToken.GetIssuer()
	it.Locale = idToken.GetLocal()
	it.Name = idToken.GetName()
	it.Nonce = idToken.GetNonce()
	it.UserID = idToken.GetId()

	return it
}

/**
 * Save Token id.
 */
func (this *OAuth2Store) saveIdToken(data *IDToken) *Config.OAuth2IdToken {
	// Get needed entities.
	configEntity := GetServer().GetConfigurationManager().getOAuthConfigurationEntity()
	clientEntity, _ := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Client", data.ClientID)
	client := clientEntity.GetObject().(*Config.OAuth2Client)

	// Create id token (OpenId)
	idToken := new(Config.OAuth2IdToken)
	idToken.SetClient(client)
	idToken.SetId(data.UserID)
	idToken.SetEmail(data.Email)
	idToken.SetEmailVerified(*data.EmailVerified)
	idToken.SetExpiration(data.Expiration)
	idToken.SetFamilyName(data.FamilyName)
	idToken.SetGivenName(data.GivenName)
	idToken.SetIssuedAt(data.IssuedAt)
	idToken.SetIssuer(data.Issuer)
	idToken.SetLocal(data.Locale)
	idToken.SetName(data.Name)
	idToken.SetNonce(data.Nonce)

	// Save into the config.
	GetServer().GetEntityManager().createEntity(configEntity.GetUuid(), "M_ids", "Config.OAuth2IdToken", idToken.GetId(), idToken)

	return idToken
}

/**
 * Save the access Data.
 */
func (this *OAuth2Store) SaveAccess(data *osin.AccessData) error {
	log.Println("Save Access")

	configEntity := GetServer().GetConfigurationManager().getOAuthConfigurationEntity()
	accessEntity, errObj := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Access", data.AccessToken)
	var access *Config.OAuth2Access
	if errObj != nil {
		access = new(Config.OAuth2Access)
		access.SetId(data.AccessToken)
		// Set the uuid.
		entity, _ := GetServer().GetEntityManager().createEntity(configEntity.GetUuid(), "M_access", "Config.OAuth2Access", access.GetId(), access)
		accessEntity = entity.(*Config_OAuth2AccessEntity)
	} else {
		access = accessEntity.GetObject().(*Config.OAuth2Access)
	}

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
	clientEntity, _ := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Client", data.Client.GetId())
	client := clientEntity.GetObject().(*Config.OAuth2Client)
	access.SetClient(client)

	// Set the authorization.
	if authorizeData == nil {
		return errors.New("authorize data must not be nil")
	}

	// Keep only the code here no the object because it will be deleted after
	// the access creation.
	access.SetAuthorize(authorizeData.Code)

	// Set other values.
	access.SetPrevious(prev)

	access.SetExpiresIn(int64(data.ExpiresIn))
	access.SetScope(data.Scope)
	access.SetRedirectUri(data.RedirectUri)

	// Set the unix time.
	access.SetCreatedAt(data.CreatedAt.Unix())

	if data.UserData != nil {
		idToken := this.saveIdToken(data.UserData.(*IDToken))
		// Set into the user user
		access.SetUserData(idToken)
	}

	// Add expire data.
	if err := addExpireAtData(data.AccessToken, data.ExpireAt()); err != nil {
		log.Println("Fail to create access data ", err)
		return err
	}

	// Now the refresh token.
	if len(data.RefreshToken) > 0 {
		// In that case I will save the refresh token.
		refreshEntity, errObj := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Refresh", data.RefreshToken)
		if errObj != nil {
			refresh := new(Config.OAuth2Refresh)
			refresh.SetId(data.RefreshToken)
			// save the access.
			refreshEntity, _ = GetServer().GetEntityManager().createEntity(configEntity.GetUuid(), "M_refresh", "Config.OAuth2Refresh", refresh.GetId(), refresh)
		}

		// Here the refresh token dosent exist so i will create it.
		refreshEntity.GetObject().(*Config.OAuth2Refresh).SetAccess(access) // Ref.

		// Set the access
		access.SetRefreshToken(refreshEntity.GetObject().(*Config.OAuth2Refresh)) // Ref

		accessEntity.SaveEntity()
	}

	return nil
}

/**
 * Load the access for a given code.
 */
func (this *OAuth2Store) LoadAccess(code string) (*osin.AccessData, error) {
	log.Println("Load Access ", code)
	accessEntity, errObj := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Access", code)
	if errObj != nil {
		return nil, errors.New("No access found with code " + code)
	}

	var access *osin.AccessData
	a := accessEntity.GetObject().(*Config.OAuth2Access)

	access = new(osin.AccessData)
	access.AccessToken = code
	access.ExpiresIn = int32(a.GetExpiresIn())
	access.Scope = a.GetScope()
	access.RedirectUri = a.GetRedirectUri()
	access.CreatedAt = time.Unix(int64(a.GetCreatedAt()), 0)

	if a.GetUserData() != nil {
		access.UserData = this.loadIdToken(a.GetUserData())
	}

	// The refresh token
	if a.GetRefreshToken() != nil {
		access.RefreshToken = a.GetRefreshToken().GetId()
	}

	// The access token
	access.AccessToken = a.GetId()

	// Now the client
	c, err := this.GetClient(a.GetClient().GetId())
	if err != nil {
		return nil, err
	}
	access.Client = c

	// The authorize
	auth, err := this.LoadAuthorize(a.GetAuthorize())
	if err != nil {
		// Try to get authorize...
		refreshToken := a.GetRefreshToken()
		if refreshToken == nil {
			return nil, err
		}

		// Get the configuration object.
		config := GetServer().GetConfigurationManager().getActiveConfigurationsEntity().GetObject().(*Config.Configurations)

		// So here The refresh token is valid i will create a new authorization
		authorizeData := new(osin.AuthorizeData)
		authorizeData.Client = c

		// I will reuse the say authorization code...
		authorizeData.Code = a.GetAuthorize()
		authorizeData.CreatedAt = time.Now()
		authorizeData.ExpiresIn = int32(config.GetOauth2Configuration().GetAuthorizationExpiration())
		authorizeData.RedirectUri = c.GetRedirectUri()
		authorizeData.Scope = a.GetScope()

		// TODO does is needed?
		authorizeData.State = ""

		// Set the new authorization here...
		access.AuthorizeData = authorizeData
	}
	access.AuthorizeData = auth

	return access, nil
}

/**
 * Remove an access code.
 */
func (this *OAuth2Store) RemoveAccess(code string) error {
	log.Println("Remove Access ", code)
	uuid := ConfigOAuth2AccessExists(code)
	if len(uuid) > 0 {
		entity, _ := GetServer().GetEntityManager().getEntityByUuid(uuid)
		GetServer().GetEntityManager().deleteEntity(entity)
		// Remove the related expire code if there one.
		uuid = ConfigOAuth2ExpiresExists(code)
		if len(uuid) > 0 {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(uuid)
			GetServer().GetEntityManager().deleteEntity(entity)
		}
		return nil
	}
	return errors.New("No access with code " + code + " was found.")
}

/**
 * Load the access data from it refresh code.
 */
func (this *OAuth2Store) LoadRefresh(code string) (*osin.AccessData, error) {
	log.Println("Load Refresh ", code)

	refreshEntity, errObj := GetServer().GetEntityManager().getEntityById("Config", "Config.OAuth2Refresh", code)
	if errObj != nil {
		return nil, errors.New("Now refresh token found with code " + code)
	}

	refresh := refreshEntity.GetObject().(*Config.OAuth2Refresh)

	// Get the access...
	access, err := this.LoadAccess(refresh.GetAccess().GetId())
	if err != nil {
		// In that case I will create a new access and associated it
		// with the refresh...
		return nil, err
	}

	// Here the access token can be
	return access, err
}

/**
 * Remove refresh.
 */
func (this *OAuth2Store) RemoveRefresh(code string) error {
	log.Println("Remove Refresh ", code)
	uuid := ConfigOAuth2RefreshExists(code)
	if len(uuid) > 0 {
		entity, _ := GetServer().GetEntityManager().getEntityByUuid(uuid)
		GetServer().GetEntityManager().deleteEntity(entity)
		return nil
	}
	return errors.New("No refresh was found with code " + code)
}
