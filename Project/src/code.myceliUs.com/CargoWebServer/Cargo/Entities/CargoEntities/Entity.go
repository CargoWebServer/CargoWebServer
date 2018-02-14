// +build CargoEntities

package CargoEntities

type Entity interface{
	/** Method of Entity **/

	/** UUID **/
	GetUuid() string

	/** Id **/
	GetId() string

}