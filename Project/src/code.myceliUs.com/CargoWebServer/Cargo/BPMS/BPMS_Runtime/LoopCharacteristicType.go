//+build BPMN
package BPMS_Runtime

type LoopCharacteristicType int

const (
	LoopCharacteristicType_StandardLoopCharacteristics LoopCharacteristicType = 1 + iota
	LoopCharacteristicType_MultiInstanceLoopCharacteristics
)
