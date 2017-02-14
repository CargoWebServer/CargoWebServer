// +build BPMN20

package BPMN20

type Collaboration interface{
	/** Method of Collaboration **/

	/** UUID **/
	GetUUID() string

	/** Name **/
	GetName() string

	/** IsClosed **/
	IsClosed() bool

	/** ChoreographyRef **/
	GetChoreographyRef() []Choreography

	/** Artifact **/
	GetArtifact() []Artifact

	/** ParticipantAssociation **/
	GetParticipantAssociation() []*ParticipantAssociation

	/** MessageFlowAssociation **/
	GetMessageFlowAssociation() []*MessageFlowAssociation

	/** ConversationAssociation **/
	GetConversationAssociation() []*ConversationAssociation

	/** Participant **/
	GetParticipant() []*Participant
	SetParticipant(interface{}) 

	/** MessageFlow **/
	GetMessageFlow() []*MessageFlow
	SetMessageFlow(interface{}) 

	/** CorrelationKey **/
	GetCorrelationKey() []*CorrelationKey

	/** ConversationNode **/
	GetConversationNode() []ConversationNode
	SetConversationNode(interface{}) 

	/** ConversationLink **/
	GetConversationLink() []*ConversationLink
	SetConversationLink(interface{}) 

}