// +build BPMN20

package BPMN20

type ThrowEvent interface{
	/** Method of ThrowEvent **/

	/** UUID **/
	GetUUID() string

	/** InputSet **/
	GetInputSet() *InputSet

	/** EventDefinitionRef **/
	GetEventDefinitionRef() []EventDefinition

	/** DataInputAssociation **/
	GetDataInputAssociation() []*DataInputAssociation

	/** DataInput **/
	GetDataInput() []*DataInput

	/** EventDefinition **/
	GetEventDefinition() []EventDefinition

}