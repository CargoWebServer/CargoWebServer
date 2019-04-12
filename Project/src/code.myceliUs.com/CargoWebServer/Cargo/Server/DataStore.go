package Server

import (
	"errors"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
)

/**
 * This is the factory function that create the correct store depending
 * of he's information.
 */
func NewDataStore(info *Config.DataStoreConfiguration) (DataStore, error) {
	if info.M_dataStoreType == Config.DataStoreType_SQL_STORE {
		dataStore, err := NewSqlDataStore(info)
		return dataStore, err
	} else if info.M_dataStoreType == Config.DataStoreType_GRAPH_STORE {
		dataStore, err := NewGraphStore(info)
		return dataStore, err
	} else {
		return nil, errors.New("No data store type was given")
	}
}

/**
 * DataStore is use to store data and do CRUD operation on it...
 */
type DataStore interface {
	/**
	 * Connection related stuff
	 */
	GetId() string

	/**
	 * Test if there's a connection with the server...
	 */
	Ping() error

	/** Crud interface **/
	Create(query string, data []interface{}) (lastId interface{}, err error)

	/**
	 * Param are filter to discard some element...
	 */
	Read(query string, fieldsType []interface{}, params []interface{}) ([][]interface{}, error)

	/**
	 * Update
	 */
	Update(query string, fields []interface{}, params []interface{}) error

	/**
	 * Delete values that match given parameter...
	 */
	Delete(query string, params []interface{}) error

	/**
	 * Close the data store, remove all connections or lnk to the data store.
	 */
	Close() error

	/**
	 * Open the data store connection.
	 */
	Connect() error

	/**
	 * Create a new entity prototype.
	 */
	CreateEntityPrototype(prototype *EntityPrototype) error

	/**
	 * Save entity prototype.
	 */
	SaveEntityPrototype(prototype *EntityPrototype) error

	/**
	 * Return the list of all entity prototypes from a dataStore
	 */
	GetEntityPrototypes() ([]*EntityPrototype, error)

	/**
	 * Return the prototype of a given type.
	 */
	GetEntityPrototype(id string) (*EntityPrototype, error)

	/**
	 * Remove a given entity prototype.
	 */
	DeleteEntityPrototype(id string) error

	/**
	 * Remove all prototypes.
	 */
	DeleteEntityPrototypes() error
}
