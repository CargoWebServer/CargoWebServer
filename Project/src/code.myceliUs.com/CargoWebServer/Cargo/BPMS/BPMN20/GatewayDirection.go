package BPMN20

type GatewayDirection int
const(
	GatewayDirection_Unspecified GatewayDirection = 1+iota
	GatewayDirection_Converging
	GatewayDirection_Diverging
	GatewayDirection_Mixed
)
