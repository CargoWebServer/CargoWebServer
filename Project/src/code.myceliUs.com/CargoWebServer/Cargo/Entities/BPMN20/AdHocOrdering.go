// +build BPMN20

package BPMN20

type AdHocOrdering int
const(
	AdHocOrdering_Parallel AdHocOrdering = 1+iota
	AdHocOrdering_Sequential
)
