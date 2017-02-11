package DI

type LabeledShape interface{
	/** Method of LabeledShape **/

	/** UUID **/
	GetUUID() string

	/** OwnedLabel **/
	GetOwnedLabel() []Label

}