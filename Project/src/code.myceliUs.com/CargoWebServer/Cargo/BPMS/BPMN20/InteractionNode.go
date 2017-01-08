package BPMN20

type InteractionNode interface{
	/** Method of InteractionNode **/

	/** UUID **/
	GetUUID() string

	/** IncomingConversationLinks **/
	GetIncomingConversationLinks() []*ConversationLink

	/** OutgoingConversationLinks **/
	GetOutgoingConversationLinks() []*ConversationLink

}