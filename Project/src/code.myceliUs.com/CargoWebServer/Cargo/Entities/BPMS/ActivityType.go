// +build BPMS

package BPMS

type ActivityType int
const(
	ActivityType_AbstractTask ActivityType = 1+iota
	ActivityType_ServiceTask
	ActivityType_UserTask
	ActivityType_ManualTask
	ActivityType_BusinessRuleTask
	ActivityType_ScriptTask
	ActivityType_EmbeddedSubprocess
	ActivityType_EventSubprocess
	ActivityType_AdHocSubprocess
	ActivityType_Transaction
	ActivityType_CallActivity
)
