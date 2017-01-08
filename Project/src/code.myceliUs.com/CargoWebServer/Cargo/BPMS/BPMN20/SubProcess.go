package BPMN20

type SubProcess interface{
	/** Method of SubProcess **/

	/** UUID **/
	GetUUID() string

	/** TriggeredByEvent **/
	GetTriggeredByEvent() bool

	/** Artifact **/
	GetArtifact() []Artifact

}