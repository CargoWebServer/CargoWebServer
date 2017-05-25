package Server

import (
	"log"
	"sort"
	"time"

	"code.myceliUs.com/Utility"
)

/**
 * Cache item contain complentary informations to manage item in the cache.
 */
type CacheItem struct {
	entity     Entity
	hit        int
	lastAccess int64
}

/**
 * Return a metric value that represent the weight of the item in the cache.
 */
func (this *CacheItem) Weight() int64 {
	now := Utility.MakeTimestamp()
	elapsed := now - this.lastAccess
	return int64(elapsed) * int64(this.hit)
}

// An array of Cache Item.
type CacheItems []*CacheItem

/**
 * Return the size of an array
 */
func (this CacheItems) Len() int {
	return len(this)
}

/**
 * Return if an element must be considère less than other item.
 */
func (this CacheItems) Less(i, j int) bool {
	return this[i].Weight() < this[j].Weight()
}

/**
 * Swap tow element of the array.
 */
func (this CacheItems) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

/**
 * Swap tow element of the array.
 */
func (this CacheItems) getItemIndex(uuid string) int {
	for i := 0; i < this.Len(); i++ {
		if this[i].entity.GetUuid() == uuid {
			return i
		}
	}

	return this.Len()
}

/**
 * Return the size in memory of items.
 */
func (this CacheItems) getSize() uint {
	var size uint
	for i := 0; i < this.Len(); i++ {
		size += this[i].entity.GetSize()
	}

	return size
}

/**
 * I will made use of BoltDB as cache backend. The cache will store information
 * of the engine on the disk.
 */
type CacheManager struct {
	/**
	 * The maximum cache size.
	 */
	max uint

	/**
	 * Contain the item
	 */
	orderedItems CacheItems

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

	// The maximum size in memory allowed to the server.
	cacheManager.max = 268435456 // 256 megabytes...
	cacheManager.orderedItems = make(CacheItems, 0)

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
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())
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
	// Set the entity.

	index := this.orderedItems.getItemIndex(entity.GetUuid())
	var item *CacheItem

	if this.orderedItems.getSize() >= this.max {
		item = new(CacheItem)
		item.entity = entity
		item.lastAccess = Utility.MakeTimestamp()

		// reorder the array
		sort.Sort(this.orderedItems)

		index := len(this.orderedItems) - 1
		this.orderedItems[index] = item

	} else if index != this.orderedItems.Len() {
		item = this.orderedItems[index]
		item.entity = entity
		item.lastAccess = Utility.MakeTimestamp()
		// Set the item at the end.
		this.orderedItems[index] = item

	} else {
		item = new(CacheItem)
		item.hit = 1
		item.entity = entity
		item.lastAccess = Utility.MakeTimestamp()
		this.orderedItems = append(this.orderedItems, item)
	}
}

/**
 * Gets an entity with a given uuid from the entitiesMap
 */
func (this *CacheManager) get(uuid string) Entity {

	index := this.orderedItems.getItemIndex(uuid)
	if index != this.orderedItems.Len() {
		this.orderedItems[index].hit += 1
		return this.orderedItems[index].entity
	}

	return nil
}

/**
 * Gets an entity with a given uuid from the entitiesMap
 */
func (this *CacheManager) remove(uuid string) {
	log.Println("------> remove item: ", uuid, " from cache")
	var orderedItems CacheItems
	for i := 0; i < len(this.orderedItems); i++ {
		if this.orderedItems[i].entity.GetUuid() != uuid {
			orderedItems = append(orderedItems, this.orderedItems[i])
		}
	}
	this.orderedItems = orderedItems
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
