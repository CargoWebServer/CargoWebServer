// +build Config

package Config

import(
	"encoding/xml"
)

type OAuth2Expires struct{

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

	/** members of OAuth2Expires **/
	M_id string
	M_expiresAt int64


	/** Associations **/
	m_parentPtr *OAuth2Configuration
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for OAuth2Expires **/
type XsdOAuth2Expires struct {
	XMLName xml.Name	`xml:"oauth2Expires"`
	M_id	string	`xml:"id,attr"`
	M_expiresAt	int64	`xml:"expiresAt,attr"`

}
/** UUID **/
func (this *OAuth2Expires) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *OAuth2Expires) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *OAuth2Expires) SetId(ref interface{}){
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** ExpiresAt **/
func (this *OAuth2Expires) GetExpiresAt() int64{
	return this.M_expiresAt
}

/** Init reference ExpiresAt **/
func (this *OAuth2Expires) SetExpiresAt(ref interface{}){
	this.M_expiresAt = ref.(int64)
}

/** Remove reference ExpiresAt **/

/** Parent **/
func (this *OAuth2Expires) GetParentPtr() *OAuth2Configuration{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *OAuth2Expires) SetParentPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.M_parentPtr = ref.(Configuration).GetUUID()
		this.m_parentPtr = ref.(*OAuth2Configuration)
	}
}

/** Remove reference Parent **/
func (this *OAuth2Expires) RemoveParentPtr(ref interface{}){
	toDelete := ref.(Configuration)
	if this.m_parentPtr!= nil {
		if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
		}
	}
}
