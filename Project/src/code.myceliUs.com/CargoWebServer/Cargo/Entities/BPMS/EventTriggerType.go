// +build BPMS

package BPMS

type EventTriggerType int
const(
	EventTriggerType_None EventTriggerType = 1+iota
	EventTriggerType_Timer
	EventTriggerType_Conditional
	EventTriggerType_Message
	EventTriggerType_Signal
	EventTriggerType_Multiple
	EventTriggerType_ParallelMultiple
	EventTriggerType_Escalation
	EventTriggerType_Error
	EventTriggerType_Compensation
	EventTriggerType_Terminate
	EventTriggerType_Cancel
	EventTriggerType_Link
	EventTriggerType_Start
)
