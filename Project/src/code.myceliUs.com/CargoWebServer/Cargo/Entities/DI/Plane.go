// +build DI

package DI

type Plane interface{
	/** Method of Plane **/

	/** UUID **/
	GetUUID() string

	/** DiagramElement **/
	GetDiagramElement() []DiagramElement
	SetDiagramElement(interface{}) 

}