package BPMNDI

type MessageVisibleKind int
const(
	MessageVisibleKind_Initiating MessageVisibleKind = 1+iota
	MessageVisibleKind_Non_initiating
)
