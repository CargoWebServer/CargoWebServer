package DI

type LabeledEdge interface{
	/** Method of LabeledEdge **/

	/** UUID **/
	GetUUID() string

	/** OwnedLabel **/
	GetOwnedLabel() []Label

}