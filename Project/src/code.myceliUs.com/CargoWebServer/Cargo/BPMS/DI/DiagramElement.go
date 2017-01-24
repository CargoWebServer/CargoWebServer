//+build DI
package DI

type DiagramElement interface {
	/** Method of DiagramElement **/

	/** UUID **/
	GetUUID() string

	/** OwningDiagram **/
	GetOwningDiagram() Diagram

	/** OwningElement **/
	GetOwningElement() DiagramElement

	/** ModelElement **/
	GetModelElement() interface{}

	/** Style **/
	GetStyle() Style

	/** OwnedElement **/
	GetOwnedElement() []DiagramElement

	/** Id **/
	GetId() string
}
