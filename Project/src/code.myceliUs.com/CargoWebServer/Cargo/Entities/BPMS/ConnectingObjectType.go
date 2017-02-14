// +build BPMS

package BPMS

type ConnectingObjectType int
const(
	ConnectingObjectType_SequenceFlow ConnectingObjectType = 1+iota
	ConnectingObjectType_MessageFlow
	ConnectingObjectType_Association
	ConnectingObjectType_DataAssociation
)
