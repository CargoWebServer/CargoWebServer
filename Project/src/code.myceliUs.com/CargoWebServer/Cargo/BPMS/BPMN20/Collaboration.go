// +build BPMN
package BPMN20

type Collaboration interface {
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

	/** MessageFlow **/
	GetMessageFlow() []*MessageFlow

	/** CorrelationKey **/
	GetCorrelationKey() []*CorrelationKey

	/** ConversationNode **/
	GetConversationNode() []ConversationNode

	/** ConversationLink **/
	GetConversationLink() []*ConversationLink
}
