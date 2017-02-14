// +build BPMN20

package BPMN20

type FlowElementsContainer interface{
	/** Method of FlowElementsContainer **/

	/** UUID **/
	GetUUID() string

	/** FlowElement **/
	GetFlowElement() []FlowElement

	/** LaneSet **/
	GetLaneSet() []*LaneSet
	SetLaneSet(interface{}) 

}