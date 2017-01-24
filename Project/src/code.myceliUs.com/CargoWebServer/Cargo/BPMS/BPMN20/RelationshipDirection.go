// +build BPMN
package BPMN20

type RelationshipDirection int

const (
	RelationshipDirection_None RelationshipDirection = 1 + iota
	RelationshipDirection_Forward
	RelationshipDirection_Backward
	RelationshipDirection_Both
)
