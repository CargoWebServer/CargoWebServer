package BPMN20

type EventBasedGatewayType int
const(
	EventBasedGatewayType_Parallel EventBasedGatewayType = 1+iota
	EventBasedGatewayType_Exclusive
)
