// +build BPMN
package BPMN20

type ItemAwareElement interface {
	/** Method of ItemAwareElement **/

	/** UUID **/
	GetUUID() string

	/** ItemSubjectRef **/
	GetItemSubjectRef() *ItemDefinition

	/** DataState **/
	GetDataState() *DataState
}
