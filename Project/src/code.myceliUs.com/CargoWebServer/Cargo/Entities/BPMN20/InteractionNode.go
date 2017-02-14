// +build BPMN20

package BPMN20

type InteractionNode interface{
	/** Method of InteractionNode **/

	/** UUID **/
	GetUUID() string

	/** IncomingConversationLinks **/
	GetIncomingConversationLinks() []*ConversationLink
	SetIncomingConversationLinks(interface{}) 

	/** OutgoingConversationLinks **/
	GetOutgoingConversationLinks() []*ConversationLink
	SetOutgoingConversationLinks(interface{}) 

}