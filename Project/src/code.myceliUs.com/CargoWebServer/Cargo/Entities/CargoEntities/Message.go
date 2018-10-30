// +build CargoEntities

package CargoEntities

type Message interface{
	/** Method of Message **/

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
	 * Return the list of reference UUID:FieldName where that entity is referenced.
	 */
	GetReferenced() []string

	/**
	 * Set a reference
	 */
	SetReferenced(uuid string, field string)

	/**
	 * Remove a reference.
	 */
	RemoveReferenced(uuid string, field string)

	/**
	 * Set a given field
	 */
	SetFieldValue(field string, value interface{}) error
	/**
	 * Return the value of a given field
	 */
	GetFieldValue(field string) interface{}

	/**
	 * Set the function GetEntityByUuid as a pointer. The entity manager can't
	 * be access by Entities package...
	 */
	SetEntityGetter(func(uuid string) (interface{}, error))

	/** Use it the set the entity on the cache. **/
	SetEntitySetter(fct func(entity interface{}))

	/** Use that function to generate the entity UUID**/
	SetUuidGenerator(fct func(entity interface{}) string)


GetBody()string
SetBody(val string)



}