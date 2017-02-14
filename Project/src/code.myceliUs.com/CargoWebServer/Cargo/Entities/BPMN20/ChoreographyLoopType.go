// +build BPMN20

package BPMN20

type ChoreographyLoopType int
const(
	ChoreographyLoopType_None ChoreographyLoopType = 1+iota
	ChoreographyLoopType_Standard
	ChoreographyLoopType_MultiInstanceSequential
	ChoreographyLoopType_MultiInstanceParallel
)
