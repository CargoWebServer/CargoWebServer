package Server

import (
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"

	"log"

	"strings"
	"time"
)

type CacheManager struct {

	/**
	 * This map will keep entities in memory.
	 */
	entitiesMap map[string]Entity

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
	 * This map will contain the sessions using the entities
	 */
	sessionIdsMap map[string][]string

	/**
	 * The channel used to populate the sessionIdsMap
	 */
	sessionIdsChannel chan struct {
		// The uuid of the entity
		entityUuid string
		// The sessionId
		sessionId string
	}

	/**
	 * The channel used to remove entities
	 */
	removeEntityChannel chan string

	/**
	 * The channel used to remove sessionIds
	 */
	removeSessionIdChannel chan string

	// stop processing when that variable are set to true...
	abortedByEnvironment chan bool

	// Ticker to ping the connection...
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
	return cacheManager
}

/**
 * Intilialization of the cacheManager
 */
func (this *CacheManager) Initialize() {
	this.entitiesMap = make(map[string]Entity, 0)
	this.sessionIdsMap = make(map[string][]string, 0)

	this.inputEntityChannel = make(chan Entity)
	this.outputEntityChannel = make(chan struct {
		entityUuid          string
		entityOutputChannel chan Entity
	})

	this.sessionIdsChannel = make(chan struct {
		entityUuid string
		sessionId  string
	})
	this.removeEntityChannel = make(chan string)
	this.removeSessionIdChannel = make(chan string)
	this.abortedByEnvironment = make(chan bool)
}

func (this *CacheManager) Start() {
	log.Println("--> Start CacheManager")
	go this.run()

	this.ticker = time.NewTicker(time.Millisecond * 2000)

	go func(ticker *time.Ticker) {
		for _ = range ticker.C {
			GetServer().GetCacheManager().flushCache()
		}
	}(this.ticker)
}

func (this *CacheManager) Stop() {
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

		case sessionInfo := <-this.sessionIdsChannel:
			this.appendEntitySessionId(this.entitiesMap[sessionInfo.entityUuid], sessionInfo.sessionId, "")

		case outputEntity := <-this.outputEntityChannel:
			outputEntity_ := this.entitiesMap[outputEntity.entityUuid]
			outputEntity.entityOutputChannel <- outputEntity_

		case sessionId := <-this.removeSessionIdChannel:
			for k, v := range this.sessionIdsMap {
				sessionIds_ := make([]string, 0)
				for i := 0; i < len(v); i++ {
					if sessionId != v[i] {
						sessionIds_ = append(sessionIds_, v[i])
					}
				}
				this.sessionIdsMap[k] = sessionIds_
			}

		case entityUuidToRemove := <-this.removeEntityChannel:
			// Delete the corresponding sessionId in the sessionIdsMap
			delete(this.sessionIdsMap, entityUuidToRemove)
			// The entity to remove.
			delete(this.entitiesMap, entityUuidToRemove)

		case done := <-this.abortedByEnvironment:
			if done {
				return
			}
		}

		//log.Println("Number of items in the cache ", len(this.entitiesMap))

	}

	// Stop the ticker.
	this.ticker.Stop()
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
 * Remove session from the cache
 */
func (this *CacheManager) removeSession(sessionId string) {
	this.removeSessionIdChannel <- sessionId
}

/**
 * Append a cacheEntityInfo to the map of cacheEntityInfoMap if it doesn't already exist, otherwise update it.
 * Must only be used locally
 */
func (this *CacheManager) appendEntitySessionId(entity Entity, sessionId string, recursivePath string) {
	if entity == nil {
		return
	}

	if strings.Contains(recursivePath, entity.GetUuid()) {
		return
	}

	recursivePath += (entity.GetUuid() + " ")

	if _, ok := this.sessionIdsMap[entity.GetUuid()]; !ok {
		// The entity is not already in the cache
		this.sessionIdsMap[entity.GetUuid()] = make([]string, 0)
	}

	if Utility.Contains(this.sessionIdsMap[entity.GetUuid()], sessionId) == false && len(sessionId) > 0 {
		this.sessionIdsMap[entity.GetUuid()] = append(this.sessionIdsMap[entity.GetUuid()], sessionId)
	}

	for i := 0; i < len(entity.GetChildsPtr()); i++ {
		this.appendEntitySessionId(entity.GetChildsPtr()[i], sessionId, recursivePath)
	}
}

/**
 * Remove unused entities from the cache
 * Must only be used locally
 */
func (this *CacheManager) flushCache() {
	for k, v := range this.sessionIdsMap {
		if len(v) == 0 {
			delete(this.entitiesMap, k)
			delete(this.sessionIdsMap, k)
		}
	}
}

/**
 * Associate the sessionId with the entityUuid in the cache
 * Can be used from the outside of this class
 */
func (this *CacheManager) register(entityUuid string, sessionId string) {
	this.sessionIdsChannel <- struct {
		entityUuid string
		sessionId  string
	}{entityUuid, sessionId}
}

/**
 * Insert entity if it doesn't already exist. Otherwise replace current entity.
 */
func (this *CacheManager) setEntity(entity Entity) {
	this.inputEntityChannel <- entity
}
