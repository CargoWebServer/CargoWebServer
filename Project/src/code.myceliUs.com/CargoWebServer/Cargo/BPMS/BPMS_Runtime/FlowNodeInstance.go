package BPMS_Runtime

type FlowNodeInstance interface {
	/** Method of FlowNodeInstance **/

	/** UUID **/
	GetUUID() string

	/** FlowNodeType **/
	GetFlowNodeType() FlowNodeType

	/** LifecycleState **/
	GetLifecycleState() LifecycleState

	/** InputRef **/
	GetInputRef() *ConnectingObject
	SetInputRef(interface{})

	/** OutputRef **/
	GetOutputRef() *ConnectingObject
	SetOutputRef(interface{})
}
