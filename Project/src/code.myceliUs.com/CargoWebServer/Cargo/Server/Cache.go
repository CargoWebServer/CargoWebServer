package Server

import (
	"log"
	"reflect"
	"strings"
	"time"

	"code.myceliUs.com/Utility"
	"github.com/allegro/bigcache"
)

var (
	cache *Cache
)

type Cache struct {
	// Cache the entitie in memory...
	m_cache *bigcache.BigCache

	// The operation channel.
	m_operations chan map[string]interface{}
}

/**
 *
 */
func newCache() *Cache {
	if cache == nil {
		cache = new(Cache)
		// TODO set those value in the server config...
		config := bigcache.Config{
			// number of shards (must be a power of 2)
			Shards: 1024,
			// time after which entry can be evicted
			LifeWindow: 10 * time.Minute,
			// rps * lifeWindow, used only in initial memory allocation
			MaxEntriesInWindow: 1000 * 10 * 60,
			// max entry size in bytes, used only in initial memory allocation
			MaxEntrySize: 500,
			// prints information about additional memory allocation
			Verbose: true,
			// cache will not allocate more memory than this limit, value in MB
			// if value is reached then the oldest entries can be overridden for the new ones
			// 0 value means no size limit
			HardMaxCacheSize: 8192,
			// callback fired when the oldest entry is removed because of its
			// expiration time or no space left for the new entry. Default value is nil which
			// means no callback and it prevents from unwrapping the oldest entry.
			OnRemove: nil,
		}

		// The Cache...
		cache.m_cache, _ = bigcache.NewBigCache(config)

		// The operation channel.
		cache.m_operations = make(chan map[string]interface{}, 0)
	}

	// Cache processing loop...
	go func(cache *Cache) {
		for {
			select {
			case operation := <-cache.m_operations:
				if operation["name"] == "getEntity" {
					uuid := operation["uuid"].(string)
					getEntity := operation["getEntity"].(chan Entity)
					typeName := strings.Split(uuid, "%")[0]
					var entity Entity
					if entry, err := cache.m_cache.Get(uuid); err == nil {
						val, err := Utility.FromBytes(entry, typeName)
						if err == nil {
							if reflect.TypeOf(val).String() != "map[string]interface {}" {
								entity = val.(Entity)
							} else {
								entity = NewDynamicEntity()
								entity.(*DynamicEntity).setObject(val.(map[string]interface{}))
							}
						} else {
							log.Println("--> go error ", uuid, err)
						}
					}

					// Return the found values.
					getEntity <- entity

				} else if operation["name"] == "setEntity" {
					entity := operation["entity"].(Entity)
					// Append in the map:
					if reflect.TypeOf(entity).String() != "*Server.DynamicEntity" {
						var bytes, err = Utility.ToBytes(entity)
						if err == nil {
							// By id
							if len(entity.Ids()) > 0 {
								id := generateEntityUuid(entity.GetTypeName(), "", entity.Ids())
								cache.m_cache.Set(id, []byte(entity.GetUuid()))
							}
							// By uuid
							cache.m_cache.Set(entity.GetUuid(), bytes)
						}
					} else {
						obj := entity.(*DynamicEntity).getValues()
						var bytes, err = Utility.ToBytes(obj)
						if err == nil {
							// By id
							if len(entity.Ids()) > 0 {
								id := generateEntityUuid(entity.GetTypeName(), "", entity.Ids())
								cache.m_cache.Set(id, []byte(entity.GetUuid()))
							}
							// By uuid
							cache.m_cache.Set(entity.GetUuid(), bytes)
						}
					}
				} else if operation["name"] == "removeEntity" {
					entity := operation["entity"].(Entity)
					// Remove it from the map.
					if len(entity.Ids()) > 0 {
						id := generateEntityUuid(entity.GetTypeName(), "", entity.Ids())
						cache.m_cache.Delete(id)
					}
					// Remove from the cache.
					cache.m_cache.Delete(entity.GetUuid())
					log.Println("Entity was remove successfully from cache ", entity.GetUuid())
				}

			}
		}
	}(cache)

	return cache
}
