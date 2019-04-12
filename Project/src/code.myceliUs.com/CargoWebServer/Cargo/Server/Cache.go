package Server

import (
	//	"errors"
	"log"
	"reflect"
	"runtime/debug"
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
	m_getEntry     chan map[string]interface{}
	m_setEntity    chan map[string]interface{}
	m_getEntity    chan map[string]interface{}
	m_removeEntity chan map[string]interface{}
	m_remove       chan map[string]interface{}
	m_getValue     chan map[string]interface{}
	m_getValues    chan map[string]interface{}
	m_setValues    chan map[string]interface{}
	m_setValue     chan map[string]interface{}
	m_deleteValue  chan map[string]interface{}
}

func (cache *Cache) getEntity(uuid string) (Entity, error) {
	typeName := strings.Split(uuid, "%")[0]
	var entity Entity
	if entry, err := cache.m_cache.Get(uuid); err == nil {
		val, err := Utility.FromBytes(entry, typeName)
		if err == nil {
			if reflect.TypeOf(val).String() != "map[string]interface {}" {
				entity = val.(Entity)
			} else {
				entity = NewDynamicEntity()
				// Set the basic entity properties only.
				entity.(*DynamicEntity).typeName = val.(map[string]interface{})["TYPENAME"].(string)
				entity.(*DynamicEntity).uuid = val.(map[string]interface{})["UUID"].(string)
				if val.(map[string]interface{})["ParentUuid"] != nil {
					entity.(*DynamicEntity).parentUuid = val.(map[string]interface{})["ParentUuid"].(string)
				}
				if val.(map[string]interface{})["ParentLnk"] != nil {
					entity.(*DynamicEntity).parentLnk = val.(map[string]interface{})["ParentLnk"].(string)
				}
			}
		} else if Utility.IsValidEntityReferenceName(string(entry)) {
			return cache.getEntity(string(entry))
		} else {
			log.Panicln("--> go error ", uuid, err)
			return nil, err
		}
	}

	return entity, nil
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
			LifeWindow: 2 * time.Minute,
			// rps * lifeWindow, used only in initial memory allocation
			MaxEntriesInWindow: 1000 * 10 * 60,
			// max entry size in bytes, used only in initial memory allocation
			MaxEntrySize: 500,
			// prints information about additional memory allocation
			Verbose: true,
			// cache will not allocate more memory than this limit, value in MB
			// if value is reached then the oldest entries can be overridden for the new ones
			// 0 value means no size limit
			HardMaxCacheSize: 4000,
			// callback fired when the oldest entry is removed because of its
			// expiration time or no space left for the new entry. Default value is nil which
			// means no callback and it prevents from unwrapping the oldest entry.
			OnRemove: func(key string, data []byte) {
				debug.FreeOSMemory()
			},
		}

		// The Cache...
		cache.m_cache, _ = bigcache.NewBigCache(config)

		// The operations channel.
		cache.m_getEntry = make(chan map[string]interface{}, 0)
		cache.m_setEntity = make(chan map[string]interface{}, 0)
		cache.m_getEntity = make(chan map[string]interface{}, 0)
		cache.m_removeEntity = make(chan map[string]interface{}, 0)
		cache.m_remove = make(chan map[string]interface{}, 0)
		cache.m_getValue = make(chan map[string]interface{}, 0)
		cache.m_getValues = make(chan map[string]interface{}, 0)
		cache.m_setValues = make(chan map[string]interface{}, 0)
		cache.m_setValue = make(chan map[string]interface{}, 0)
		cache.m_deleteValue = make(chan map[string]interface{}, 0)

	}

	// Cache processing loop...
	go func(cache *Cache) {
		// Keep information if an entity need to be saved.
		for {
			select {
			case operation := <-cache.m_getEntry:

				getEntry := operation["getEntry"].(chan []interface{})
				uuid := operation["uuid"].(string)
				entry, err := cache.getEntity(uuid)
				getEntry <- []interface{}{entry, err}

			case operation := <-cache.m_getEntity:

				uuid := operation["uuid"].(string)
				getEntity := operation["getEntity"].(chan Entity)
				entity, _ := cache.getEntity(uuid)

				// Return the found values.
				getEntity <- entity

			case operation := <-cache.m_setEntity:

				entity := operation["entity"].(Entity)
				// Append in the map: setObject set the value for DynamicEntity...
				if reflect.TypeOf(entity).String() != "*Server.DynamicEntity" {
					var bytes, err = Utility.ToBytes(entity)
					if err == nil {
						bytes_, err := cache.m_cache.Get(entity.GetUuid())
						if err == nil {
							if string(bytes_) != string(bytes) {
								// By id
								if len(entity.Ids()) > 0 {
									id := generateEntityUuid(entity.GetTypeName(), "", entity.Ids())
									cache.m_cache.Set(id, []byte(entity.GetUuid()))
								}
								// By uuid
								cache.m_cache.Set(entity.GetUuid(), bytes)
							}
						} else {
							// By id
							if len(entity.Ids()) > 0 {
								id := generateEntityUuid(entity.GetTypeName(), "", entity.Ids())
								cache.m_cache.Set(id, []byte(entity.GetUuid()))
							}
							// By uuid
							cache.m_cache.Set(entity.GetUuid(), bytes)
						}

					} else {
						LogInfo("---> fail to serialyse entity!")
					}
				}
			case operation := <-cache.m_removeEntity:

				entity := operation["entity"].(Entity)
				// Remove it from the map.
				if len(entity.Ids()) > 0 {
					id := generateEntityUuid(entity.GetTypeName(), "", entity.Ids())
					cache.m_cache.Delete(id)
				}
				LogInfo("--->cache delete entity: ", entity.GetUuid())
				// Remove from the cache.
				cache.m_cache.Delete(entity.GetUuid())

			case operation := <-cache.m_remove:

				uuid := operation["uuid"].(string)
				// Remove from the cache.
				cache.m_cache.Delete(uuid)

			case operation := <-cache.m_getValue:

				uuid := operation["uuid"].(string)
				field := operation["field"].(string)
				//LogInfo("--->cache getValue ", uuid)
				getValue := operation["getValue"].(chan interface{})
				var value interface{}
				typeName := strings.Split(uuid, "%")[0]
				if entry, err := cache.m_cache.Get(uuid); err == nil {
					val, err := Utility.FromBytes(entry, typeName)
					if err == nil {
						value = val.(map[string]interface{})[field]
					}
				}
				// Return the found values.
				getValue <- value

			case operation := <-cache.m_getValues:
				uuid := operation["uuid"].(string)
				getValues := operation["getValues"].(chan map[string]interface{})
				var values map[string]interface{}
				typeName := strings.Split(uuid, "%")[0]
				if entry, err := cache.m_cache.Get(uuid); err == nil {
					val, err := Utility.FromBytes(entry, typeName)
					if err == nil {
						values = val.(map[string]interface{})
					}
				} else {
					// get values from the db instead.
					values, _ = getEntityByUuid(uuid)
					LogInfo("---> no values found in the cache for uuid ", uuid, err.Error())
				}
				// Return the found values.
				getValues <- values

			case operation := <-cache.m_setValues:
				values := operation["values"].(map[string]interface{})

				var bytes, err = Utility.ToBytes(values)
				if err == nil {
					bytes_, err := cache.m_cache.Get(values["UUID"].(string))
					if err == nil {
						needSave := string(bytes) != string(bytes_)
						if needSave {
							// By id
							if values["Ids"] != nil {
								id := generateEntityUuid(values["TYPENAME"].(string), "", values["Ids"].([]interface{}))
								cache.m_cache.Set(id, []byte(values["UUID"].(string)))
							}
							// By uuid
							cache.m_cache.Set(values["UUID"].(string), bytes)
						}
						// set if the entity need to be save.
						operation["needSave"].(chan bool) <- needSave
					} else {
						if values["Ids"] != nil {
							id := generateEntityUuid(values["TYPENAME"].(string), "", values["Ids"].([]interface{}))
							cache.m_cache.Set(id, []byte(values["UUID"].(string)))
						}
						// By uuid
						cache.m_cache.Set(values["UUID"].(string), bytes)
						// true by default.
						operation["needSave"].(chan bool) <- true
					}
				} else {
					operation["needSave"].(chan bool) <- true
				}

			case operation := <-cache.m_setValue:

				uuid := operation["uuid"].(string)
				field := operation["field"].(string)
				value := operation["value"].(interface{})

				typeName := strings.Split(uuid, "%")[0]
				if entry, err := cache.m_cache.Get(uuid); err == nil {
					val, err := Utility.FromBytes(entry, typeName)
					if err == nil {
						values := val.(map[string]interface{})

						// Test if the value has change.
						needSave := false
						if values[field] != nil {
							needSave = Utility.GetChecksum(value) != Utility.GetChecksum(values[field])
						} else {
							needSave = true
						}

						if needSave {
							//LogInfo("---> setValue ", value)
							// Set the field value...
							values[field] = value
							var bytes, err = Utility.ToBytes(values)
							if err == nil {
								// By id
								if values["ids"] != nil {
									id := generateEntityUuid(values["TYPENAME"].(string), "", values["Ids"].([]interface{}))
									cache.m_cache.Set(id, []byte(values["UUID"].(string)))
								}
								// By uuid
								cache.m_cache.Set(values["UUID"].(string), bytes)
							}
						}

						operation["needSave"].(chan bool) <- needSave

					} else {
						// new value here...
						operation["needSave"].(chan bool) <- true
					}

				}
			case operation := <-cache.m_deleteValue:
				uuid := operation["uuid"].(string)
				field := operation["field"].(string)
				// LogInfo("--->cache deleteValue: ", uuid)
				typeName := strings.Split(uuid, "%")[0]
				if entry, err := cache.m_cache.Get(uuid); err == nil {
					val, err := Utility.FromBytes(entry, typeName)
					if err == nil {
						values := val.(map[string]interface{})
						delete(values, field)
						var bytes, err = Utility.ToBytes(values)
						if err == nil {
							// By id
							if values["ids"] != nil {
								id := generateEntityUuid(values["TYPENAME"].(string), "", values["Ids"].([]interface{}))
								cache.m_cache.Set(id, []byte(values["UUID"].(string)))
							}
							// By uuid
							cache.m_cache.Set(values["UUID"].(string), bytes)
						}
					}
				}
			}
		}
	}(cache)

	return cache
}
