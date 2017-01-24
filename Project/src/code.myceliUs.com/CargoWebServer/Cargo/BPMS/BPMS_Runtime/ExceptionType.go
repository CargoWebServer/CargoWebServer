//+build BPMN
package BPMS_Runtime

type ExceptionType int

const (
	ExceptionType_NoIORuleException ExceptionType = 1 + iota
	ExceptionType_GatewayException
	ExceptionType_NoAvailableOutputSetException
	ExceptionType_NotMatchingIOSpecification
	ExceptionType_IllegalStartEventException
)
