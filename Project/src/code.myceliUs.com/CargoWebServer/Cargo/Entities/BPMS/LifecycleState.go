package BPMS

type LifecycleState int
const(
	LifecycleState_Completed LifecycleState = 1+iota
	LifecycleState_Compensated
	LifecycleState_Failed
	LifecycleState_Terminated
	LifecycleState_Completing
	LifecycleState_Compensating
	LifecycleState_Failing
	LifecycleState_Terminating
	LifecycleState_Ready
	LifecycleState_Active
)
