// +build BPMN
package BPMNDI

type ParticipantBandKind int

const (
	ParticipantBandKind_Top_initiating ParticipantBandKind = 1 + iota
	ParticipantBandKind_Middle_initiating
	ParticipantBandKind_Bottom_initiating
	ParticipantBandKind_Top_non_initiating
	ParticipantBandKind_Middle_non_initiating
	ParticipantBandKind_Bottom_non_initiating
)
