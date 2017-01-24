// +build BPMN
package BPMN20

type ProcessType int

const (
	ProcessType_None ProcessType = 1 + iota
	ProcessType_Public
	ProcessType_Private
)
