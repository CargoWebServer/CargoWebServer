package BPMN20

type BaseElement interface{
	/** Method of BaseElement **/

	/** UUID **/
	GetUUID() string

	/** Id **/
	GetId() string

	/** Other **/
	GetOther() interface{}

	/** ExtensionElements **/
	GetExtensionElements() *ExtensionElements

	/** ExtensionDefinitions **/
	GetExtensionDefinitions() []*ExtensionDefinition

	/** ExtensionValues **/
	GetExtensionValues() []*ExtensionAttributeValue

	/** Documentation **/
	GetDocumentation() []*Documentation

}