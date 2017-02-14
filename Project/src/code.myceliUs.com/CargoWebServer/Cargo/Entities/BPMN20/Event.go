// +build BPMN20

package BPMN20

type Event interface{
	/** Method of Event **/

	/** UUID **/
	GetUUID() string

	/** Property **/
	GetProperty() []*Property
	SetProperty(interface{}) 

}