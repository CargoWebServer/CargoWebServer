package BPMN20

type ConversationNode interface{
	/** Method of ConversationNode **/

	/** UUID **/
	GetUUID() string

	/** Name **/
	GetName() string

	/** ParticipantRef **/
	GetParticipantRef() []*Participant

	/** MessageFlowRef **/
	GetMessageFlowRef() []*MessageFlow

	/** CorrelationKey **/
	GetCorrelationKey() []*CorrelationKey

}