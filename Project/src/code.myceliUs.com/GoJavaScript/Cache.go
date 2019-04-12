package GoJavaScript

import (
	"reflect"
	"time"
)

// The cache is simply a map of object that keep go object accessible for JavaScript.
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
					// increment the reference count.
					cache.m_objects[operation["id"].(string)] = operation["object"]

					//log.Println("set object ----> ", operation["id"].(string), len(cache.m_objects))
				} else if operation["name"] == "removeObject" {
					// The object will leave for more 30 second the time the
					// script get object from it reference.
					delete(cache.m_objects, operation["id"].(string))
					//log.Println("delete object ----> ", operation["id"].(string), len(cache.m_objects))
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
	result := <-values["result"].(chan interface{})

	return result
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
	// The server delete the reference before the cache
	// return it result so I will give a 30 second more lifespan to
	// the object before remove it form the client cache, other wise null
	// values will be return.
	delay := time.NewTimer(30 * time.Second)
	go func() {
		<-delay.C
		values := make(map[string]interface{})
		values["name"] = "removeObject"
		values["id"] = id
		cache.m_operations <- values
	}()
}

// Convert objectRef to object as needed.
func GetObject(val interface{}) interface{} {
	if val == nil {
		return nil
	}

	if reflect.TypeOf(val).String() == "GoJavaScript.ObjectRef" {
		ref := val.(ObjectRef)
		if GetCache().GetObject(ref.UUID) != nil {
			return GetCache().GetObject(ref.UUID)
		}
		return nil

	} else if reflect.TypeOf(val).String() == "*GoJavaScript.ObjectRef" {
		ref := val.(*ObjectRef)
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
				if reflect.TypeOf(e.Interface()).String() == "GoJavaScript.ObjectRef" {
					ref := e.Interface().(ObjectRef)
					if GetCache().GetObject(ref.UUID) != nil {
						obj := GetCache().GetObject(ref.UUID)
						if obj != nil {
							if i == 0 {
								values = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(obj)), 0, slice.Len())
							}
							values = reflect.Append(values, reflect.ValueOf(obj))
						}
					}
				} else if reflect.TypeOf(e.Interface()).String() == "*GoJavaScript.ObjectRef" {
					if !e.IsNil() {
						ref := e.Interface().(*ObjectRef)
						if GetCache().GetObject(ref.UUID) != nil {
							obj := GetCache().GetObject(ref.UUID)
							if obj != nil {
								if i == 0 {
									values = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(obj)), 0, slice.Len())
								}
								values = reflect.Append(values, reflect.ValueOf(obj))
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
