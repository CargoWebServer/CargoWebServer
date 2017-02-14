// +build BPMN20

package BPMN20

type CallableElement interface{
	/** Method of CallableElement **/

	/** UUID **/
	GetUUID() string

	/** Name **/
	GetName() string

	/** IoSpecification **/
	GetIoSpecification() *InputOutputSpecification

	/** SupportedInterfaceRef **/
	GetSupportedInterfaceRef() []*Interface

	/** IoBinding **/
	GetIoBinding() []*InputOutputBinding

}