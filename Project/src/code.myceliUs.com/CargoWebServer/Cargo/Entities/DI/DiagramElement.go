// +build DI

package DI

type DiagramElement interface{
	/** Method of DiagramElement **/

	/** UUID **/
	GetUUID() string

	/** OwningDiagram **/
	GetOwningDiagram() Diagram
	SetOwningDiagram(interface{}) 

	/** OwningElement **/
	GetOwningElement() DiagramElement
	SetOwningElement(interface{}) 

	/** ModelElement **/
	GetModelElement() interface{}
	SetModelElement(interface{}) 

	/** Style **/
	GetStyle() Style
	SetStyle(interface{}) 

	/** OwnedElement **/
	GetOwnedElement() []DiagramElement
	SetOwnedElement(interface{}) 

	/** Id **/
	GetId() string
	SetId(interface{}) 

}