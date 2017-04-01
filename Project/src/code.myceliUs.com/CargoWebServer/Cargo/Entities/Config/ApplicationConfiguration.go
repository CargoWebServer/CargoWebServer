// +build Config

package Config

import(
"encoding/xml"
)

type ApplicationConfiguration struct{

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

	/** members of Configuration **/
	M_id string

	/** members of ApplicationConfiguration **/
	M_indexPage string


	/** Associations **/
	m_parentPtr *Configurations
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for ApplicationConfiguration **/
type XsdApplicationConfiguration struct {
	XMLName xml.Name	`xml:"applicationConfiguration"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_indexPage	string	`xml:"indexPage,attr"`

}
/** UUID **/
func (this *ApplicationConfiguration) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *ApplicationConfiguration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ApplicationConfiguration) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** IndexPage **/
func (this *ApplicationConfiguration) GetIndexPage() string{
	return this.M_indexPage
}

/** Init reference IndexPage **/
func (this *ApplicationConfiguration) SetIndexPage(ref interface{}){
	this.NeedSave = true
	this.M_indexPage = ref.(string)
}

/** Remove reference IndexPage **/

/** Parent **/
func (this *ApplicationConfiguration) GetParentPtr() *Configurations{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *ApplicationConfiguration) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*Configurations)
		this.M_parentPtr = ref.(*Configurations).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *ApplicationConfiguration) RemoveParentPtr(ref interface{}){
	toDelete := ref.(*Configurations)
	if this.m_parentPtr!= nil {
		if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
		}else{
			this.NeedSave = true
		}
	}
}
