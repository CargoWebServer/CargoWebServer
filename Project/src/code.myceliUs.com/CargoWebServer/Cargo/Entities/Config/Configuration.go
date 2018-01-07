// +build Config

package Config

type Configuration interface {
	/** Method of Configuration **/

	/** UUID **/
	GetUUID() string

	/** Id **/
	GetId() string
}
