package Server

////////////////////////////////////////////////////////////////////////////////
//						Entity Reference
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
//						Entity Query
////////////////////////////////////////////////////////////////////////////////
/**
 * The query is use to specifying the basic information it's like
 * the select, insert or update of sql...
 */
type EntityQuery struct {
	// Must be Server.EntityQuery
	TYPENAME string

	// The name of the entity
	TypeName string
	// The list of field to retreive, delete or modify
	Fields []string
	// The base index, this must be of form indexFieldName=indexFieldValue
	Indexs []string
	// The query to execute by the search engine.
	Query string
}

/**
 * The entity interface regroups methods needed by a structure to be
 * saved and initialized from the key value data store.
 */
type Entity interface {

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
	 * Evaluate if an entity needs to be saved.
	 */
	IsNeedSave() bool
}
