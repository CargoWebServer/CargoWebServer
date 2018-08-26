package GoJerryScript

import "log"

// The cache is simply a map of object that keep go object accessible for JerryScript.
type Cache struct {
	// The map where object are store.
	m_objects map[string]interface{}

	// The channel to send cache operations.
	m_operations chan map[string]interface{}
}

// The singleton object.
var cache *Cache

func GetCache() *Cache {
	if cache == nil {
		newCache()
	}
	return cache
}

// Create a new object cache.
func newCache() {
	if cache == nil {
		cache = new(Cache)

		// The map of objects.
		cache.m_objects = make(map[string]interface{})

		// The operation channel.
		cache.m_operations = make(chan map[string]interface{}, 0)
	}

	// The cache
	go func() {
		for {
			select {
			case operation := <-cache.m_operations:
				if operation["name"] == "getObject" {
					var object = cache.m_objects[operation["id"].(string)]
					operation["result"].(chan interface{}) <- object
				} else if operation["name"] == "setObject" {
					cache.m_objects[operation["id"].(string)] = operation["object"]
					log.Println("---> cache now contain: ", len(cache.m_objects))
				} else if operation["name"] == "removeObject" {
					delete(cache.m_objects, operation["id"].(string))
				}
			}
		}
	}()
}

// Return an object from the cache.
func (cache *Cache) GetObject(id string) interface{} {
	// Here I will get object from the cache.
	values := make(map[string]interface{})
	values["name"] = "getObject"
	values["id"] = id
	values["result"] = make(chan interface{})
	cache.m_operations <- values
	// wait to the result to be found.
	return <-values["result"].(chan interface{})
}

func (cache *Cache) SetObject(id string, object interface{}) {
	// Here I will set object in the cache
	values := make(map[string]interface{})
	values["name"] = "setObject"
	values["id"] = id
	values["object"] = object
	cache.m_operations <- values
}

func (cache *Cache) RemoveObject(id string) {
	values := make(map[string]interface{})
	values["name"] = "removeObject"
	values["id"] = id
	cache.m_operations <- values
}
