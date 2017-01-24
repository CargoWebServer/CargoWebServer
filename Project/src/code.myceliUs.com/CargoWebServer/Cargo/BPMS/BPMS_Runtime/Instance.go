//+build BPMN
package BPMS_Runtime

type Instance interface {
	/** Method of Instance **/

	/** UUID **/
	GetUUID() string

	/** Id **/
	GetId() string

	/** BpmnElementId **/
	GetBpmnElementId() string

	/** Participants **/
	GetParticipants() []string

	/** DataRef **/
	GetDataRef() []*ItemAwareElementInstance

	/** Data **/
	GetData() []*ItemAwareElementInstance
	SetData(interface{})

	/** LogInfoRef **/
	GetLogInfoRef() []*LogInfo
	SetLogInfoRef(interface{})
}
