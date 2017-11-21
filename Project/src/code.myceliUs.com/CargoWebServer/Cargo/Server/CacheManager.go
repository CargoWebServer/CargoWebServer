package Server

import (
	"log"
	"time"
	// Use to see if there is memory leak remove for production.
	"runtime/debug"
	"strings"
)

/**
 * The cache manager is simply a map of Entity accessible via channel and where
 * entity has a limited lifespan of 10 minutes. It main purpose is to back data
 * the time of initislisation, and also for multiple access. The real db is the
 * KV store or the SQL store.
 */
type CacheManager struct {

	/**
	 * Contain the item
	 */
	entities map[string]Entity

	/**
	 * Channel used to append entity inside the entity map
	 */
	inputEntityChannel chan struct {
		// The entity to input
		entity Entity
		// The channel to block the execution util the entity is in the map.
		wait chan bool
	}

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
	removeEntityChannel chan struct {
		// The entity to input
		uuid string
		// The channel to block the execution util the entity is in the map.
		wait chan bool
	}

	// stop processing when that variable are set to true...
	abortedByEnvironment chan bool

	// Use to free the os memory.
	ticker *time.Ticker
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
	this.inputEntityChannel = make(chan struct {
		entity Entity
		wait   chan bool
	})

	this.outputEntityChannel = make(chan struct {
		entityUuid          string
		entityOutputChannel chan Entity
	})

	this.removeEntityChannel = make(chan struct {
		uuid string
		wait chan bool
	})

	this.abortedByEnvironment = make(chan bool)
	this.ticker = time.NewTicker(1 * time.Hour)
}

func (this *CacheManager) start() {
	log.Println("--> Start CacheManager")
	go this.run()

	// Here I will compact the memory after 10 minutes...
	go func(ticker *time.Ticker) {
		for t := range ticker.C {
			log.Println("--> call debug.FreeOSMemory()", t)
			debug.FreeOSMemory()
		}
	}(this.ticker)
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
			this.set(inputEntity.entity)
			inputEntity.wait <- false // Unblock the channel.

		case outputEntity := <-this.outputEntityChannel:
			outputEntity_ := this.get(outputEntity.entityUuid)
			outputEntity.entityOutputChannel <- outputEntity_

		case entityUuidToRemove := <-this.removeEntityChannel:
			// The entity to remove.
			this.remove(entityUuidToRemove.uuid)
			entityUuidToRemove.wait <- false

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

	// Remove the entity from the cache after 10 minutes.
	// exception are Config.*, session, action and account those entities
	// are keep in the cache until they are explicitly remove.
	if !strings.HasPrefix(entity.GetTypeName(), "Config.") && entity.GetTypeName() != "CargoEntities.Action" && entity.GetTypeName() != "CargoEntities.Session" && entity.GetTypeName() != "CargoEntities.Account" {
		go func(uuid string, lifespan time.Duration) {
			timer := time.NewTimer(lifespan * time.Minute)
			<-timer.C
			GetServer().GetCacheManager().removeEntity(uuid)
		}(entity.GetUuid(), 10)
	}

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
	defer close(outputInfo.entityOutputChannel)
	outputInfo.entityUuid = uuid
	this.outputEntityChannel <- *outputInfo
	entity := <-outputInfo.entityOutputChannel

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
	toRemove := new(struct {
		uuid string
		wait chan bool
	})

	toRemove.uuid = uuid
	toRemove.wait = make(chan bool)
	defer close(toRemove.wait)
	this.removeEntityChannel <- *toRemove
	<-toRemove.wait
}

/**
 * Insert entity if it doesn't already exist. Otherwise replace current entity.
 */
func (this *CacheManager) setEntity(entity Entity) {
	input := new(struct {
		entity Entity
		wait   chan bool
	})

	input.entity = entity
	input.wait = make(chan bool)
	defer close(input.wait)

	this.inputEntityChannel <- *input
	// Wait before enter other entity.
	<-input.wait
}
