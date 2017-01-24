// +build BPMN
package BPMN20

type FlowNode interface {
	/** Method of FlowNode **/

	/** UUID **/
	GetUUID() string

	/** Outgoing **/
	GetOutgoing() []*SequenceFlow

	/** Incoming **/
	GetIncoming() []*SequenceFlow

	/** Lanes **/
	GetLanes() []*Lane
}
