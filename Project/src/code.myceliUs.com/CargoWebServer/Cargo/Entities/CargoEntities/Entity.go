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
	 * Get an entity's parent pointer.
	 */
	 GetParent() interface{}

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
	 * Return list of childs uuid.
	 */
	GetChildsUuid() []string

	/**
	 * Set the function GetEntityByUuid as a pointer. The entity manager can't
	 * be access by Entities package...
	 */
	 SetEntityGetter(func(uuid string) (interface{}, error))


GetId()string
SetId(val string)



}