// +build BPMN
package BPMN20

type MultiInstanceFlowCondition int

const (
	MultiInstanceFlowCondition_None MultiInstanceFlowCondition = 1 + iota
	MultiInstanceFlowCondition_One
	MultiInstanceFlowCondition_All
	MultiInstanceFlowCondition_Complex
)
