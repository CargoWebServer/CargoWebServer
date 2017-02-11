package BPMS

type GatewayType int
const(
	GatewayType_ParallelGateway GatewayType = 1+iota
	GatewayType_ExclusiveGateway
	GatewayType_InclusiveGateway
	GatewayType_EventBasedGateway
	GatewayType_ComplexGateway
)
