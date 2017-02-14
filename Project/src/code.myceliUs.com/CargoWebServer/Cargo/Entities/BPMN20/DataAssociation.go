// +build BPMN20

package BPMN20

type DataAssociation interface{
	/** Method of DataAssociation **/

	/** UUID **/
	GetUUID() string

	/** Transformation **/
	GetTransformation() *FormalExpression

	/** Assignment **/
	GetAssignment() []*Assignment

	/** TargetRef **/
	GetTargetRef() ItemAwareElement

	/** SourceRef **/
	GetSourceRef() []ItemAwareElement

}