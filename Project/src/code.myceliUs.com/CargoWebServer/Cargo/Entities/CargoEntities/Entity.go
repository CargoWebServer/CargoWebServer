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
	 * Return link to entity childs.
	 */
	GetChilds() []interface{}

	/**
	 * Return list of childs uuid.
	 */
	GetChildsUuid() []string

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

	/**
	 * Return true if the entity need to be saved.
	 */
	IsNeedSave() bool

	/**
	 * Set the needsave value.
	 */
	SetNeedSave(needsave bool)


GetId()string
SetId(val string)



}