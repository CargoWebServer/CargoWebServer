// +build BPMN
package BPMN20

type ResourceRole interface {
	/** Method of ResourceRole **/

	/** UUID **/
	GetUUID() string

	/** ResourceRef **/
	GetResourceRef() *Resource

	/** ResourceParameterBinding **/
	GetResourceParameterBinding() []*ResourceParameterBinding

	/** ResourceAssignmentExpression **/
	GetResourceAssignmentExpression() *ResourceAssignmentExpression

	/** Name **/
	GetName() string
}
