// +build CargoEntities

package CargoEntities

type Message interface{
	/** Method of Message **/

	/** UUID **/
	GetUUID() string

	/** Body **/
	GetBody() string

}