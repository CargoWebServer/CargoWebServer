// +build CargoEntities

package CargoEntities

type Entity interface{
	/** Method of Entity **/

	/**
	 * Return the type name of an entity
	 */
	GetTypeName() string

	/**
	 * Get an entity's uuid
	 * Each entity must have one uuid.
	 */
	GetUuid() string
	SetUuid(uuid string)

	/**
	 * Return the array of entity id's without it uuid.
	 */
	Ids() []interface{}

	/**
	 * Get an entity's parent UUID if it have a parent.
	 */
	GetParentUuid() string
	SetParentUuid(uuid string)

	/**
	 * The name of the relation with it parent.
	 */
	GetParentLnk() string
	SetParentLnk(string)

	/**
	 * Return link to entity childs.
	 */
	GetChilds() []interface{}

	/**
	 * Evaluate if an entity needs to be saved.
	 */
	IsNeedSave() bool

	/**
	 * Set the need save state to false.
	 */

	ResetNeedSave()

	/**
	 * Set the function GetEntityByUuid as a pointer. The entity manager can't
	 * be access by Entities package...
	 */
	 SetEntityGetter(func(uuid string) (interface{}, error))



}