// +build Config

package Config

import (
	"encoding/xml"
)

type OAuth2Access struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of OAuth2Access **/
	M_id     string
	m_client *OAuth2Client
	/** If the ref is a string and not an object **/
	M_client       string
	M_authorize    string
	M_previous     string
	m_refreshToken *OAuth2Refresh
	/** If the ref is a string and not an object **/
	M_refreshToken string
	M_expiresIn    int64
	M_scope        string
	M_redirectUri  string
	m_userData     *OAuth2IdToken
	/** If the ref is a string and not an object **/
	M_userData  string
	M_createdAt int64

	/** Associations **/
	m_parentPtr *OAuth2Configuration
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2Access **/
type XsdOAuth2Access struct {
	XMLName            xml.Name `xml:"oauth2Access"`
	M_id               string   `xml:"id,attr"`
	M_authorize        string   `xml:"authorize,attr"`
	M_previous         string   `xml:"previous,attr"`
	M_expiresIn        int64    `xml:"expiresIn ,attr"`
	M_scope            string   `xml:"scope,attr"`
	M_redirectUri      string   `xml:"redirectUri,attr"`
	M_tokenUri         string   `xml:"tokenUri,attr"`
	M_authorizationUri string   `xml:"authorizationUri,attr"`
	M_createdAt        int64    `xml:"createdAt ,attr"`
}

/** UUID **/
func (this *OAuth2Access) GetUUID() string {
	return this.UUID
}

/** Id **/
func (this *OAuth2Access) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Access) SetId(ref interface{}) {
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference Id **/

/** Client **/
func (this *OAuth2Access) GetClient() *OAuth2Client {
	return this.m_client
}

/** Init reference Client **/
func (this *OAuth2Access) SetClient(ref interface{}) {
	if _, ok := ref.(string); ok {
		if this.M_client != ref.(string) {
			this.M_client = ref.(string)
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
	} else {
		if this.M_client != ref.(*OAuth2Client).GetUUID() {
			this.M_client = ref.(*OAuth2Client).GetUUID()
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
		this.m_client = ref.(*OAuth2Client)
	}
}

/** Remove reference Client **/
func (this *OAuth2Access) RemoveClient(ref interface{}) {
	toDelete := ref.(*OAuth2Client)
	if this.m_client != nil {
		if toDelete.GetUUID() == this.m_client.GetUUID() {
			this.m_client = nil
			this.M_client = ""
			this.NeedSave = true
		}
	}
}

/** Authorize **/
func (this *OAuth2Access) GetAuthorize() string {
	return this.M_authorize
}

/** Init reference Authorize **/
func (this *OAuth2Access) SetAuthorize(ref interface{}) {
	if this.M_authorize != ref.(string) {
		this.M_authorize = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference Authorize **/

/** Previous **/
func (this *OAuth2Access) GetPrevious() string {
	return this.M_previous
}

/** Init reference Previous **/
func (this *OAuth2Access) SetPrevious(ref interface{}) {
	if this.M_previous != ref.(string) {
		this.M_previous = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference Previous **/

/** RefreshToken **/
func (this *OAuth2Access) GetRefreshToken() *OAuth2Refresh {
	return this.m_refreshToken
}

/** Init reference RefreshToken **/
func (this *OAuth2Access) SetRefreshToken(ref interface{}) {
	if _, ok := ref.(string); ok {
		if this.M_refreshToken != ref.(string) {
			this.M_refreshToken = ref.(string)
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
	} else {
		if this.M_refreshToken != ref.(*OAuth2Refresh).GetUUID() {
			this.M_refreshToken = ref.(*OAuth2Refresh).GetUUID()
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
		this.m_refreshToken = ref.(*OAuth2Refresh)
	}
}

/** Remove reference RefreshToken **/
func (this *OAuth2Access) RemoveRefreshToken(ref interface{}) {
	toDelete := ref.(*OAuth2Refresh)
	if this.m_refreshToken != nil {
		if toDelete.GetUUID() == this.m_refreshToken.GetUUID() {
			this.m_refreshToken = nil
			this.M_refreshToken = ""
			this.NeedSave = true
		}
	}
}

/** ExpiresIn **/
func (this *OAuth2Access) GetExpiresIn() int64 {
	return this.M_expiresIn
}

/** Init reference ExpiresIn **/
func (this *OAuth2Access) SetExpiresIn(ref interface{}) {
	if this.M_expiresIn != ref.(int64) {
		this.M_expiresIn = ref.(int64)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference ExpiresIn **/

/** Scope **/
func (this *OAuth2Access) GetScope() string {
	return this.M_scope
}

/** Init reference Scope **/
func (this *OAuth2Access) SetScope(ref interface{}) {
	if this.M_scope != ref.(string) {
		this.M_scope = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference Scope **/

/** RedirectUri **/
func (this *OAuth2Access) GetRedirectUri() string {
	return this.M_redirectUri
}

/** Init reference RedirectUri **/
func (this *OAuth2Access) SetRedirectUri(ref interface{}) {
	if this.M_redirectUri != ref.(string) {
		this.M_redirectUri = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference RedirectUri **/

/** UserData **/
func (this *OAuth2Access) GetUserData() *OAuth2IdToken {
	return this.m_userData
}

/** Init reference UserData **/
func (this *OAuth2Access) SetUserData(ref interface{}) {
	if _, ok := ref.(string); ok {
		if this.M_userData != ref.(string) {
			this.M_userData = ref.(string)
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
	} else {
		if this.M_userData != ref.(*OAuth2IdToken).GetUUID() {
			this.M_userData = ref.(*OAuth2IdToken).GetUUID()
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
		this.m_userData = ref.(*OAuth2IdToken)
	}
}

/** Remove reference UserData **/
func (this *OAuth2Access) RemoveUserData(ref interface{}) {
	toDelete := ref.(*OAuth2IdToken)
	if this.m_userData != nil {
		if toDelete.GetUUID() == this.m_userData.GetUUID() {
			this.m_userData = nil
			this.M_userData = ""
			this.NeedSave = true
		}
	}
}

/** CreatedAt **/
func (this *OAuth2Access) GetCreatedAt() int64 {
	return this.M_createdAt
}

/** Init reference CreatedAt **/
func (this *OAuth2Access) SetCreatedAt(ref interface{}) {
	if this.M_createdAt != ref.(int64) {
		this.M_createdAt = ref.(int64)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference CreatedAt **/

/** Parent **/
func (this *OAuth2Access) GetParentPtr() *OAuth2Configuration {
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Access) SetParentPtr(ref interface{}) {
	if _, ok := ref.(string); ok {
		if this.M_parentPtr != ref.(string) {
			this.M_parentPtr = ref.(string)
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
	} else {
		if this.M_parentPtr != ref.(Configuration).GetUUID() {
			this.M_parentPtr = ref.(Configuration).GetUUID()
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
		this.m_parentPtr = ref.(*OAuth2Configuration)
	}
}

/** Remove reference Parent **/
func (this *OAuth2Access) RemoveParentPtr(ref interface{}) {
	toDelete := ref.(Configuration)
	if this.m_parentPtr != nil {
		if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
