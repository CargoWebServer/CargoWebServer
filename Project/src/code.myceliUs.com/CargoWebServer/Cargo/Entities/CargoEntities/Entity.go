// +build CargoEntities

package CargoEntities

type Entity interface {
	/** Method of Entity **/

	/** UUID **/
	GetUUID() string

	/** Id **/
	GetId() string
}
