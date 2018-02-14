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
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool

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
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2IdToken) GetUuid() string{
	return this.UUID
}
func (this *OAuth2IdToken) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2IdToken) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2IdToken) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2IdToken"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2IdToken) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2IdToken) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2IdToken) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2IdToken) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *OAuth2IdToken) IsNeedSave() bool{
	return this.NeedSave
}


/** Issuer **/
func (this *OAuth2IdToken) GetIssuer() string{
	return this.M_issuer
}

/** Init reference Issuer **/
func (this *OAuth2IdToken) SetIssuer(ref interface{}){
	if this.M_issuer != ref.(string) {
		this.M_issuer = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Issuer **/

/** Id **/
func (this *OAuth2IdToken) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2IdToken) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** Client **/
func (this *OAuth2IdToken) GetClient() *OAuth2Client{
	return this.m_client
}

/** Init reference Client **/
func (this *OAuth2IdToken) SetClient(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_client != ref.(string) {
			this.M_client = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_client != ref.(*OAuth2Client).GetUuid() {
			this.M_client = ref.(*OAuth2Client).GetUuid()
			this.NeedSave = true
		}
		this.m_client = ref.(*OAuth2Client)
	}
}

/** Remove reference Client **/
func (this *OAuth2IdToken) RemoveClient(ref interface{}){
	toDelete := ref.(*OAuth2Client)
	if this.m_client!= nil {
		if toDelete.GetUuid() == this.m_client.GetUuid() {
			this.m_client = nil
			this.M_client = ""
			this.NeedSave = true
		}
	}
}

/** Expiration **/
func (this *OAuth2IdToken) GetExpiration() int64{
	return this.M_expiration
}

/** Init reference Expiration **/
func (this *OAuth2IdToken) SetExpiration(ref interface{}){
	if this.M_expiration != ref.(int64) {
		this.M_expiration = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference Expiration **/

/** IssuedAt **/
func (this *OAuth2IdToken) GetIssuedAt() int64{
	return this.M_issuedAt
}

/** Init reference IssuedAt **/
func (this *OAuth2IdToken) SetIssuedAt(ref interface{}){
	if this.M_issuedAt != ref.(int64) {
		this.M_issuedAt = ref.(int64)
		this.NeedSave = true
	}
}

/** Remove reference IssuedAt **/

/** Nonce **/
func (this *OAuth2IdToken) GetNonce() string{
	return this.M_nonce
}

/** Init reference Nonce **/
func (this *OAuth2IdToken) SetNonce(ref interface{}){
	if this.M_nonce != ref.(string) {
		this.M_nonce = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Nonce **/

/** Email **/
func (this *OAuth2IdToken) GetEmail() string{
	return this.M_email
}

/** Init reference Email **/
func (this *OAuth2IdToken) SetEmail(ref interface{}){
	if this.M_email != ref.(string) {
		this.M_email = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Email **/

/** EmailVerified **/
func (this *OAuth2IdToken) GetEmailVerified() bool{
	return this.M_emailVerified
}

/** Init reference EmailVerified **/
func (this *OAuth2IdToken) SetEmailVerified(ref interface{}){
	if this.M_emailVerified != ref.(bool) {
		this.M_emailVerified = ref.(bool)
		this.NeedSave = true
	}
}

/** Remove reference EmailVerified **/

/** Name **/
func (this *OAuth2IdToken) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *OAuth2IdToken) SetName(ref interface{}){
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Name **/

/** FamilyName **/
func (this *OAuth2IdToken) GetFamilyName() string{
	return this.M_familyName
}

/** Init reference FamilyName **/
func (this *OAuth2IdToken) SetFamilyName(ref interface{}){
	if this.M_familyName != ref.(string) {
		this.M_familyName = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference FamilyName **/

/** GivenName **/
func (this *OAuth2IdToken) GetGivenName() string{
	return this.M_givenName
}

/** Init reference GivenName **/
func (this *OAuth2IdToken) SetGivenName(ref interface{}){
	if this.M_givenName != ref.(string) {
		this.M_givenName = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference GivenName **/

/** Local **/
func (this *OAuth2IdToken) GetLocal() string{
	return this.M_local
}

/** Init reference Local **/
func (this *OAuth2IdToken) SetLocal(ref interface{}){
	if this.M_local != ref.(string) {
		this.M_local = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Local **/

/** Parent **/
func (this *OAuth2IdToken) GetParentPtr() *OAuth2Configuration{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2IdToken) SetParentPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_parentPtr != ref.(string) {
			this.M_parentPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_parentPtr != ref.(Configuration).GetUuid() {
			this.M_parentPtr = ref.(Configuration).GetUuid()
			this.NeedSave = true
		}
		this.m_parentPtr = ref.(*OAuth2Configuration)
	}
}

/** Remove reference Parent **/
func (this *OAuth2IdToken) RemoveParentPtr(ref interface{}){
	toDelete := ref.(Configuration)
	if this.m_parentPtr!= nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
