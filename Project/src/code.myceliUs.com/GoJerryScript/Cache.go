package GoJerryScript

//import "log"

// The cache is simply a map of object that keep go object accessible for JerryScript.
type Cache struct {
	// The map where object are store.
	m_objects map[string]interface{}

	// The js object references
	m_jsObjects map[string]Uint32_t

	// The object count.
	m_objectsCount map[string]uint

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
		cache.m_jsObjects = make(map[string]Uint32_t)
		cache.m_objectsCount = make(map[string]uint)

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
				} else if operation["name"] == "getJsObject" {
					var object = cache.m_jsObjects[operation["id"].(string)]
					operation["result"].(chan Uint32_t) <- object
				} else if operation["name"] == "setObject" {
					// Increment the reference count.
					if count, ok := cache.m_objectsCount[operation["id"].(string)]; ok {
						cache.m_objectsCount[operation["id"].(string)] = count + 1
					} else {
						cache.m_objectsCount[operation["id"].(string)] = uint(1)
					}
					// increment the reference count.
					cache.m_objects[operation["id"].(string)] = operation["object"]

				} else if operation["name"] == "setJsObject" {
					if count, ok := cache.m_objectsCount[operation["id"].(string)]; ok {
						cache.m_objectsCount[operation["id"].(string)] = count + 1
					} else {
						cache.m_objectsCount[operation["id"].(string)] = uint(1)
					}

					// set the object in the map.
					cache.m_jsObjects[operation["id"].(string)] = operation["jsObject"].(Uint32_t)

				} else if operation["name"] == "removeObject" {
					delete(cache.m_objects, operation["id"].(string))
					delete(cache.m_jsObjects, operation["id"].(string))
				}
			}
		}
	}()
}

// Return an object from the cache.
func (cache *Cache) getObject(id string) interface{} {
	// Here I will get object from the cache.
	values := make(map[string]interface{})
	values["name"] = "getObject"
	values["id"] = id
	values["result"] = make(chan interface{})
	cache.m_operations <- values
	// wait to the result to be found.
	return <-values["result"].(chan interface{})
}

func (cache *Cache) getJsObject(id string) Uint32_t {
	// Here I will get object from the cache.
	values := make(map[string]interface{})
	values["name"] = "getJsObject"
	values["id"] = id
	values["result"] = make(chan Uint32_t)
	cache.m_operations <- values
	// wait to the result to be found.
	return <-values["result"].(chan Uint32_t)
}

func (cache *Cache) setObject(id string, object interface{}) {
	// Here I will set object in the cache
	values := make(map[string]interface{})
	values["name"] = "setObject"
	values["id"] = id
	values["object"] = object
	cache.m_operations <- values
}

func (cache *Cache) setJsObject(id string, jsObject Uint32_t) {
	// Here I will set object in the cache
	values := make(map[string]interface{})
	values["name"] = "setJsObject"
	values["id"] = id
	values["jsObject"] = jsObject
	cache.m_operations <- values
}

func (cache *Cache) removeObject(id string) {
	values := make(map[string]interface{})
	values["name"] = "removeObject"
	values["id"] = id
	cache.m_operations <- values
}
