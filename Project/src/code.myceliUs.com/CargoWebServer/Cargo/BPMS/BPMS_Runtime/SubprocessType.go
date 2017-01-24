//+build BPMN
package BPMS_Runtime

type SubprocessType int

const (
	SubprocessType_EmbeddedSubprocess SubprocessType = 1 + iota
	SubprocessType_EventSubprocess
	SubprocessType_AdHocSubprocess
	SubprocessType_Transaction
)
