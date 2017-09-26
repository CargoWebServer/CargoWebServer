// +build Config

package Config

import(
	"encoding/xml"
)

type OAuth2IdToken struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of OAuth2IdToken **/
	M_issuer string
	M_id string
	m_client *OAuth2Client
	/** If the ref is a string and not an object **/
	M_client string
	M_expiration int64
	M_issuedAt int64
	M_nonce string
	M_email string
	M_emailVerified bool
	M_name string
	M_familyName string
	M_givenName string
	M_local string


	/** Associations **/
	m_parentPtr *OAuth2Configuration
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2IdToken **/
type XsdOAuth2IdToken struct {
	XMLName xml.Name	`xml:"oauth2IdToken"`
	M_id	string	`xml:"id,attr"`
	M_issuer	string	`xml:"issuer,attr"`
	M_expiration	int64	`xml:"expiration,attr"`
	M_issuedAt	int64	`xml:"issuedAt,attr"`
	M_nonce	string	`xml:"nonce,attr"`
	M_email	string	`xml:"email,attr"`
	M_emailVerified	bool	`xml:"emailVerified,attr"`
	M_name	string	`xml:"name,attr"`
	M_familyName	string	`xml:"familyName,attr"`
	M_givenName	string	`xml:"givenName,attr"`
	M_local	string	`xml:"local,attr"`

}
/** UUID **/
func (this *OAuth2IdToken) GetUUID() string{
	return this.UUID
}

/** Issuer **/
func (this *OAuth2IdToken) GetIssuer() string{
	return this.M_issuer
}

/** Init reference Issuer **/
func (this *OAuth2IdToken) SetIssuer(ref interface{}){
	this.M_issuer = ref.(string)
}

/** Remove reference Issuer **/

/** Id **/
func (this *OAuth2IdToken) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2IdToken) SetId(ref interface{}){
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Client **/
func (this *OAuth2IdToken) GetClient() *OAuth2Client{
	return this.m_client
}

/** Init reference Client **/
func (this *OAuth2IdToken) SetClient(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_client = ref.(string)
	}else{
		this.M_client = ref.(*OAuth2Client).GetUUID()
		this.m_client = ref.(*OAuth2Client)
	}
}

/** Remove reference Client **/
func (this *OAuth2IdToken) RemoveClient(ref interface{}){
	toDelete := ref.(*OAuth2Client)
	if this.m_client!= nil {
		if toDelete.GetUUID() == this.m_client.GetUUID() {
			this.m_client = nil
			this.M_client = ""
		}
	}
}

/** Expiration **/
func (this *OAuth2IdToken) GetExpiration() int64{
	return this.M_expiration
}

/** Init reference Expiration **/
func (this *OAuth2IdToken) SetExpiration(ref interface{}){
	this.M_expiration = ref.(int64)
}

/** Remove reference Expiration **/

/** IssuedAt **/
func (this *OAuth2IdToken) GetIssuedAt() int64{
	return this.M_issuedAt
}

/** Init reference IssuedAt **/
func (this *OAuth2IdToken) SetIssuedAt(ref interface{}){
	this.M_issuedAt = ref.(int64)
}

/** Remove reference IssuedAt **/

/** Nonce **/
func (this *OAuth2IdToken) GetNonce() string{
	return this.M_nonce
}

/** Init reference Nonce **/
func (this *OAuth2IdToken) SetNonce(ref interface{}){
	this.M_nonce = ref.(string)
}

/** Remove reference Nonce **/

/** Email **/
func (this *OAuth2IdToken) GetEmail() string{
	return this.M_email
}

/** Init reference Email **/
func (this *OAuth2IdToken) SetEmail(ref interface{}){
	this.M_email = ref.(string)
}

/** Remove reference Email **/

/** EmailVerified **/
func (this *OAuth2IdToken) GetEmailVerified() bool{
	return this.M_emailVerified
}

/** Init reference EmailVerified **/
func (this *OAuth2IdToken) SetEmailVerified(ref interface{}){
	this.M_emailVerified = ref.(bool)
}

/** Remove reference EmailVerified **/

/** Name **/
func (this *OAuth2IdToken) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *OAuth2IdToken) SetName(ref interface{}){
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** FamilyName **/
func (this *OAuth2IdToken) GetFamilyName() string{
	return this.M_familyName
}

/** Init reference FamilyName **/
func (this *OAuth2IdToken) SetFamilyName(ref interface{}){
	this.M_familyName = ref.(string)
}

/** Remove reference FamilyName **/

/** GivenName **/
func (this *OAuth2IdToken) GetGivenName() string{
	return this.M_givenName
}

/** Init reference GivenName **/
func (this *OAuth2IdToken) SetGivenName(ref interface{}){
	this.M_givenName = ref.(string)
}

/** Remove reference GivenName **/

/** Local **/
func (this *OAuth2IdToken) GetLocal() string{
	return this.M_local
}

/** Init reference Local **/
func (this *OAuth2IdToken) SetLocal(ref interface{}){
	this.M_local = ref.(string)
}

/** Remove reference Local **/

/** Parent **/
func (this *OAuth2IdToken) GetParentPtr() *OAuth2Configuration{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2IdToken) SetParentPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.M_parentPtr = ref.(Configuration).GetUUID()
		this.m_parentPtr = ref.(*OAuth2Configuration)
	}
}

/** Remove reference Parent **/
func (this *OAuth2IdToken) RemoveParentPtr(ref interface{}){
	toDelete := ref.(Configuration)
	if this.m_parentPtr!= nil {
		if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
		}
	}
}
