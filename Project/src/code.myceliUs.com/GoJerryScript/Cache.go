package GoJerryScript

import (
	"log"
	"reflect"
)

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
					// increment the reference count.
					cache.m_objects[operation["id"].(string)] = operation["object"]
				} else if operation["name"] == "setJsObject" {
					// set the object in the map.
					cache.m_jsObjects[operation["id"].(string)] = operation["jsObject"].(Uint32_t)
				} else if operation["name"] == "removeObject" {
					delete(cache.m_objects, operation["id"].(string))
					obj := cache.m_jsObjects[operation["id"].(string)]
					// remove it pointer from the interpreter.
					if obj != nil {
						//Jerry_release_value(obj)
						delete(cache.m_jsObjects, operation["id"].(string))
						Jerry_release_value(obj)
					}
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

func (cache *Cache) SetObject(id string, object interface{}) {
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

func (cache *Cache) RemoveObject(id string) {
	values := make(map[string]interface{})
	values["name"] = "removeObject"
	values["id"] = id
	cache.m_operations <- values
}

// Convert objectRef to object as needed.
func GetObject(val interface{}) interface{} {

	if reflect.TypeOf(val).String() == "GoJerryScript.ObjectRef" {
		ref := val.(ObjectRef)
		if GetCache().GetObject(ref.UUID) != nil {
			return GetCache().GetObject(ref.UUID)
		}
		return nil

	} else if reflect.TypeOf(val).Kind() == reflect.Slice {
		// In case of a slice I will transform the object ref with it actual values.
		slice := reflect.ValueOf(val)
		var values reflect.Value
		for i := 0; i < slice.Len(); i++ {
			e := slice.Index(i)
			if e.IsValid() {
				if !e.IsNil() {
					if reflect.TypeOf(e.Interface()).String() == "GoJerryScript.ObjectRef" {
						ref := e.Interface().(ObjectRef)
						if GetCache().GetObject(ref.UUID) != nil {
							obj := GetCache().GetObject(ref.UUID)
							if obj != nil {
								if i == 0 {
									values = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(obj)), 0, slice.Len())
								}
								values = reflect.Append(values, reflect.ValueOf(obj))
							} else {
								log.Println("---> fail to retreive object ", ref.UUID)
							}
						}
					}
				}
			}
		}
		// return values with object instead of object ref.
		if values.IsValid() {
			return values.Interface()
		}
	}

	// No conversion was necessary.
	return val
}
