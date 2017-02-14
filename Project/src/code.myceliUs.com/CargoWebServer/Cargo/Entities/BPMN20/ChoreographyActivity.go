// +build BPMN20

package BPMN20

type ChoreographyActivity interface{
	/** Method of ChoreographyActivity **/

	/** UUID **/
	GetUUID() string

	/** ParticipantRef **/
	GetParticipantRef() []*Participant

	/** InitiatingParticipantRef **/
	GetInitiatingParticipantRef() *Participant

	/** CorrelationKey **/
	GetCorrelationKey() []*CorrelationKey
	SetCorrelationKey(interface{}) 

	/** LoopType **/
	GetLoopType() ChoreographyLoopType
	SetLoopType(interface{}) 

}