// +build DI

package DI

type Diagram interface{
	/** Method of Diagram **/

	/** UUID **/
	GetUUID() string

	/** RootElement **/
	GetRootElement() DiagramElement
	SetRootElement(interface{}) 

	/** Name **/
	GetName() string
	SetName(interface{}) 

	/** Id **/
	GetId() string
	SetId(interface{}) 

	/** Documentation **/
	GetDocumentation() string
	SetDocumentation(interface{}) 

	/** Resolution **/
	GetResolution() float64
	SetResolution(interface{}) 

	/** OwnedStyle **/
	GetOwnedStyle() []Style
	SetOwnedStyle(interface{}) 

}