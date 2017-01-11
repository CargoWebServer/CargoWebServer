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

	/** LoopType **/
	GetLoopType() ChoreographyLoopType

}