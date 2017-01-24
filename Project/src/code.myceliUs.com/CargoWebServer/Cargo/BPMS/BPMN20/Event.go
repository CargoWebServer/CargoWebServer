// +build BPMN
package BPMN20

type Event interface {
	/** Method of Event **/

	/** UUID **/
	GetUUID() string

	/** Property **/
	GetProperty() []*Property
}
