// +build BPMN
package BPMN20

type FlowElement interface {
	/** Method of FlowElement **/

	/** UUID **/
	GetUUID() string

	/** Name **/
	GetName() string

	/** Auditing **/
	GetAuditing() *Auditing

	/** Monitoring **/
	GetMonitoring() *Monitoring

	/** CategoryValueRef **/
	GetCategoryValueRef() []*CategoryValue
}
