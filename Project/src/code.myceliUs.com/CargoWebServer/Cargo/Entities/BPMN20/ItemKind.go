// +build BPMN20

package BPMN20

type ItemKind int
const(
	ItemKind_Physical ItemKind = 1+iota
	ItemKind_Information
)
