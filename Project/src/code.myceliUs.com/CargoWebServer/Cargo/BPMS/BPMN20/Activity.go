package BPMN20

type Activity interface {
	/** Method of Activity **/

	/** UUID **/
	GetUUID() string

	/** IsForCompensation **/
	IsForCompensation() bool

	/** LoopCharacteristics **/
	GetLoopCharacteristics() LoopCharacteristics

	/** ResourceRole **/
	GetResourceRole() []ResourceRole

	/** Default **/
	GetDefault() *SequenceFlow

	/** Property **/
	GetProperty() []*Property

	/** IoSpecification **/
	GetIoSpecification() *InputOutputSpecification

	/** BoundaryEventRefs **/
	GetBoundaryEventRefs() []*BoundaryEvent

	/** DataInputAssociation **/
	GetDataInputAssociation() []*DataInputAssociation

	/** DataOutputAssociation **/
	GetDataOutputAssociation() []*DataOutputAssociation

	/** StartQuantity **/
	GetStartQuantity() int

	/** CompletionQuantity **/
	GetCompletionQuantity() int
}
