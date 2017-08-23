// +build Config

package Config

import(
	"encoding/xml"
)

type ScheduledTask struct{

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

	/** members of ScheduledTask **/
	M_script string
	M_interval int
	M_keepAlive bool


	/** Associations **/
	m_parentPtr *Configurations
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for ScheduledTask **/
type XsdScheduledTask struct {
	XMLName xml.Name	`xml:"scheduledTask"`
	M_script	string	`xml:"script,attr"`
	M_interval	int	`xml:"interval,attr"`
	M_keepAlive	bool	`xml:"keepAlive,attr"`

}
/** UUID **/
func (this *ScheduledTask) GetUUID() string{
	return this.UUID
}

/** Script **/
func (this *ScheduledTask) GetScript() string{
	return this.M_script
}

/** Init reference Script **/
func (this *ScheduledTask) SetScript(ref interface{}){
	this.NeedSave = true
	this.M_script = ref.(string)
}

/** Remove reference Script **/

/** Interval **/
func (this *ScheduledTask) GetInterval() int{
	return this.M_interval
}

/** Init reference Interval **/
func (this *ScheduledTask) SetInterval(ref interface{}){
	this.NeedSave = true
	this.M_interval = ref.(int)
}

/** Remove reference Interval **/

/** KeepAlive **/
func (this *ScheduledTask) GetKeepAlive() bool{
	return this.M_keepAlive
}

/** Init reference KeepAlive **/
func (this *ScheduledTask) SetKeepAlive(ref interface{}){
	this.NeedSave = true
	this.M_keepAlive = ref.(bool)
}

/** Remove reference KeepAlive **/

/** Parent **/
func (this *ScheduledTask) GetParentPtr() *Configurations{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *ScheduledTask) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*Configurations)
		this.M_parentPtr = ref.(*Configurations).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *ScheduledTask) RemoveParentPtr(ref interface{}){
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
