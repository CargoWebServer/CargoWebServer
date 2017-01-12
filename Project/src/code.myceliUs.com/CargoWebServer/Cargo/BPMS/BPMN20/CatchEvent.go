package BPMN20

type CatchEvent interface{
	/** Method of CatchEvent **/

	/** UUID **/
	GetUUID() string

	/** ParallelMultiple **/
	GetParallelMultiple() bool

	/** OutputSet **/
	GetOutputSet() *OutputSet

	/** EventDefinitionRef **/
	GetEventDefinitionRef() []EventDefinition

	/** DataOutputAssociation **/
	GetDataOutputAssociation() []*DataOutputAssociation

	/** DataOutput **/
	GetDataOutput() []*DataOutput

	/** EventDefinition **/
	GetEventDefinition() []EventDefinition

}