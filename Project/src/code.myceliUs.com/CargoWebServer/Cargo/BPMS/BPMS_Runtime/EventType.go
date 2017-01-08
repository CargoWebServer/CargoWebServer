package BPMS_Runtime

type EventType int
const(
	EventType_StartEvent EventType = 1+iota
	EventType_IntermediateCatchEvent
	EventType_BoundaryEvent
	EventType_EndEvent
	EventType_IntermediateThrowEvent
)
