package Server

import (
	"log"
	"time"
)

type CacheManager struct {

	/**
	 * This map will keep entities in memory.
	 */
	entitiesMap map[string]Entity

	/**
	 * Each entity has a lifetime of 30 second. If an entity is not use
	 * inside that interval it will be flush from the cache.
	 */
	timersMap map[string]*time.Timer

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
	return cacheManager
}

var (
	timeout = time.Duration(5 * time.Minute)
)

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Intilialization of the cacheManager
 */
func (this *CacheManager) initialize() {

	log.Println("--> Initialize CacheManager")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())

	this.entitiesMap = make(map[string]Entity, 0)
	this.timersMap = make(map[string]*time.Timer)

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
			this.entitiesMap[inputEntity.GetUuid()] = inputEntity
			if this.timersMap[inputEntity.GetUuid()] != nil {
				// In that case I will reset the timer time to 30 sec...
				this.timersMap[inputEntity.GetUuid()].Reset(time.Second * timeout)
			} else {
				this.timersMap[inputEntity.GetUuid()] = time.NewTimer(time.Second * timeout)
			}

			go func(timer *time.Timer, uuid string, removeChannel chan string) {

				// Remove the entity from the cache.
				<-timer.C
				//log.Println("Timer expired for entity ", uuid)
				removeChannel <- uuid

			}(this.timersMap[inputEntity.GetUuid()], inputEntity.GetUuid(), this.removeEntityChannel)

		case outputEntity := <-this.outputEntityChannel:
			outputEntity_ := this.entitiesMap[outputEntity.entityUuid]
			// Reset the timeout value here.
			if this.timersMap[outputEntity.entityUuid] != nil {
				this.timersMap[outputEntity.entityUuid].Reset(time.Second * timeout)
			}
			outputEntity.entityOutputChannel <- outputEntity_

		case entityUuidToRemove := <-this.removeEntityChannel:
			// The entity to remove.
			delete(this.entitiesMap, entityUuidToRemove)
			delete(this.timersMap, entityUuidToRemove)

		case done := <-this.abortedByEnvironment:
			if done {
				return
			}
		}

		//log.Println("Number of items in the cache ", len(this.entitiesMap))

	}
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
