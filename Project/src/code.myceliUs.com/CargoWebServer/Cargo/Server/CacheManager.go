package Server

import (
	"log"
	"time"
)

/**
 * I will made use of BoltDB as cache backend. The cache will store information
 * of the engine on the disk.
 */
type CacheManager struct {

	/**
	 * Contain the item
	 */
	entities map[string]Entity

	/**
	 * Channel used to append entity inside the entity map
	 */
	inputEntityChannel chan Entity

	/**
	* Channel used to output entity
	 */
	outputEntityChannel chan struct {
		// The uuid of the entity
		entityUuid string
		// The output channel
		entityOutputChannel chan Entity
	}

	/**
	 * The channel used to remove entities
	 */
	removeEntityChannel chan string

	// stop processing when that variable are set to true...
	abortedByEnvironment chan bool
}

var cacheManager *CacheManager

func (this *Server) GetCacheManager() *CacheManager {
	if cacheManager == nil {
		cacheManager = newCacheManager()
	}
	return cacheManager
}

/**
 * The cacheManager manages the memory and the lifetime of entities.
 */
func newCacheManager() *CacheManager {
	cacheManager := new(CacheManager)

	// The maximum size in memory allowed to the server.
	cacheManager.entities = make(map[string]Entity, 0)

	return cacheManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Intilialization of the cacheManager
 */
func (this *CacheManager) initialize() {

	log.Println("--> Initialize CacheManager")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)
	this.inputEntityChannel = make(chan Entity)
	this.outputEntityChannel = make(chan struct {
		entityUuid          string
		entityOutputChannel chan Entity
	})

	this.removeEntityChannel = make(chan string)
	this.abortedByEnvironment = make(chan bool)
}

func (this *CacheManager) start() {
	log.Println("--> Start CacheManager")
	go this.run()
}

func (this *CacheManager) getId() string {
	return "CacheManager"
}

func (this *CacheManager) stop() {
	log.Println("--> Stop CacheManager")

	// Free the cache
	this.abortedByEnvironment <- true
}

/**
 * Processing message from outside threads
 */
func (this *CacheManager) run() {
	for {
		select {
		case inputEntity := <-this.inputEntityChannel:
			// Append entity to the database.
			this.set(inputEntity)

		case outputEntity := <-this.outputEntityChannel:
			outputEntity_ := this.get(outputEntity.entityUuid)
			outputEntity.entityOutputChannel <- outputEntity_

		case entityUuidToRemove := <-this.removeEntityChannel:
			// The entity to remove.
			this.remove(entityUuidToRemove)

		case done := <-this.abortedByEnvironment:
			if done {
				return
			}
		}
	}
}

/**
 * Gets an entity with a given uuid from the entitiesMap
 */
func (this *CacheManager) set(entity Entity) {

	this.entities[entity.GetUuid()] = entity

	// Lifespan of entity in the cache manager is one minute.
	go func(uuid string, removeEntityChannel chan string) {
		timer := time.NewTimer(1 * time.Minute)
		<-timer.C
		removeEntityChannel <- uuid
	}(entity.GetUuid(), this.removeEntityChannel)

}

/**
 * Gets an entity with a given uuid from the entitiesMap
 */
func (this *CacheManager) get(uuid string) Entity {
	return this.entities[uuid]
}

/**
 * Gets an entity with a given uuid from the entitiesMap
 */
func (this *CacheManager) remove(uuid string) {
	delete(this.entities, uuid)
}

/**
 * Gets an entity with a given uuid from the entitiesMap
 */
func (this *CacheManager) getEntity(uuid string) Entity {
	outputInfo := new(struct {
		entityUuid          string
		entityOutputChannel chan Entity
	})

	outputInfo.entityOutputChannel = make(chan Entity)
	outputInfo.entityUuid = uuid
	this.outputEntityChannel <- *outputInfo
	entity := <-outputInfo.entityOutputChannel

	defer close(outputInfo.entityOutputChannel)
	return entity
}

/**
 * Determine if the entity exists in the map.
 */
func (this *CacheManager) contains(uuid string) (Entity, bool) {

	entity := this.getEntity(uuid)
	if entity != nil {
		return entity, true
	}

	return nil, false
}

/**
 * Remove an existing entity with a given uuid.
 */
func (this *CacheManager) removeEntity(uuid string) {
	this.removeEntityChannel <- uuid
}

/**
 * Insert entity if it doesn't already exist. Otherwise replace current entity.
 */
func (this *CacheManager) setEntity(entity Entity) {
	this.inputEntityChannel <- entity
}
