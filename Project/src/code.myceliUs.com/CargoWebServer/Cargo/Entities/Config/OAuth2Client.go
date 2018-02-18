// +build Config

package Config

import (
	"encoding/xml"
)

type OAuth2Client struct {

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
	/** Get entity by uuid function **/
	getEntityByUuid func(string) (interface{}, error)

	/** members of OAuth2Client **/
	M_id               string
	M_secret           string
	M_redirectUri      string
	M_tokenUri         string
	M_authorizationUri string
	M_extra            []uint8

	/** Associations **/
	m_parentPtr *OAuth2Configuration
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2Client **/
type XsdOAuth2Client struct {
	XMLName       xml.Name `xml:"oauth2Client"`
	M_id          string   `xml:"id,attr"`
	M_secret      string   `xml:"secret,attr"`
	M_redirectUri string   `xml:"redirectUri,attr"`
	M_extra       []uint8  `xml:"extra,attr"`
}

/***************** Entity **************************/

/** UUID **/
func (this *OAuth2Client) GetUuid() string {
	return this.UUID
}
func (this *OAuth2Client) SetUuid(uuid string) {
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2Client) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2Client) GetTypeName() string {
	this.TYPENAME = "Config.OAuth2Client"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2Client) GetParentUuid() string {
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2Client) SetParentUuid(parentUuid string) {
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Client) GetParentLnk() string {
	return this.ParentLnk
}
func (this *OAuth2Client) SetParentLnk(parentLnk string) {
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *OAuth2Client) IsNeedSave() bool {
	return this.NeedSave
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Client) SetEntityGetter(fct func(uuid string) (interface{}, error)) {
	this.getEntityByUuid = fct
}

/** Id **/
func (this *OAuth2Client) GetId() string {
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Client) SetId(ref interface{}) {
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Secret **/
func (this *OAuth2Client) GetSecret() string {
	return this.M_secret
}

/** Init reference Secret **/
func (this *OAuth2Client) SetSecret(ref interface{}) {
	if this.M_secret != ref.(string) {
		this.M_secret = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Secret **/

/** RedirectUri **/
func (this *OAuth2Client) GetRedirectUri() string {
	return this.M_redirectUri
}

/** Init reference RedirectUri **/
func (this *OAuth2Client) SetRedirectUri(ref interface{}) {
	if this.M_redirectUri != ref.(string) {
		this.M_redirectUri = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference RedirectUri **/

/** TokenUri **/
func (this *OAuth2Client) GetTokenUri() string {
	return this.M_tokenUri
}

/** Init reference TokenUri **/
func (this *OAuth2Client) SetTokenUri(ref interface{}) {
	if this.M_tokenUri != ref.(string) {
		this.M_tokenUri = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference TokenUri **/

/** AuthorizationUri **/
func (this *OAuth2Client) GetAuthorizationUri() string {
	return this.M_authorizationUri
}

/** Init reference AuthorizationUri **/
func (this *OAuth2Client) SetAuthorizationUri(ref interface{}) {
	if this.M_authorizationUri != ref.(string) {
		this.M_authorizationUri = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference AuthorizationUri **/

/** Extra **/
func (this *OAuth2Client) GetExtra() []uint8 {
	return this.M_extra
}

/** Init reference Extra **/
func (this *OAuth2Client) SetExtra(ref interface{}) {
	/*if this.M_extra != ref.([]uint8) {
		this.M_extra = ref.([]uint8)
		this.NeedSave = true
	}*/
}

/** Remove reference Extra **/

/** Parent **/
func (this *OAuth2Client) GetParentPtr() *OAuth2Configuration {
	if this.m_parentPtr == nil {
		entity, err := this.getEntityByUuid(this.M_parentPtr)
		if err == nil {
			this.m_parentPtr = entity.(*OAuth2Configuration)
		}
	}
	return this.m_parentPtr
}
func (this *OAuth2Client) GetParentPtrStr() string {
	return this.M_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Client) SetParentPtr(ref interface{}) {
	if _, ok := ref.(string); ok {
		if this.M_parentPtr != ref.(string) {
			this.M_parentPtr = ref.(string)
			this.NeedSave = true
		}
	} else {
		if this.M_parentPtr != ref.(Configuration).GetUuid() {
			this.M_parentPtr = ref.(Configuration).GetUuid()
			this.NeedSave = true
		}
		this.m_parentPtr = ref.(*OAuth2Configuration)
	}
}

/** Remove reference Parent **/
func (this *OAuth2Client) RemoveParentPtr(ref interface{}) {
	toDelete := ref.(Configuration)
	if this.m_parentPtr != nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
