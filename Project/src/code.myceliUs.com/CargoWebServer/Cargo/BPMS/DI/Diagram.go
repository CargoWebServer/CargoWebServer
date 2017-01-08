package DI

type Diagram interface{
	/** Method of Diagram **/

	/** UUID **/
	GetUUID() string

	/** RootElement **/
	GetRootElement() DiagramElement

	/** Name **/
	GetName() string

	/** Id **/
	GetId() string

	/** Documentation **/
	GetDocumentation() string

	/** Resolution **/
	GetResolution() float64

	/** OwnedStyle **/
	GetOwnedStyle() []Style

}