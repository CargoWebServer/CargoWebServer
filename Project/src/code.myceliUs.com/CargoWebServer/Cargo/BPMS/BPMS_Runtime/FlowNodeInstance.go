//+build BPMN
package BPMS_Runtime

type FlowNodeInstance interface {
	/** Method of FlowNodeInstance **/

	/** UUID **/
	GetUUID() string

	/** FlowNodeType **/
	GetFlowNodeType() FlowNodeType

	/** LifecycleState **/
	GetLifecycleState() LifecycleState
	SetLifecycleState(interface{})

	/** InputRef **/
	GetInputRef() []*ConnectingObject
	SetInputRef(interface{})

	/** OutputRef **/
	GetOutputRef() []*ConnectingObject
	SetOutputRef(interface{})

	/** Get the parent process instance pointer **/
	GetProcessInstancePtr() *ProcessInstance
}
