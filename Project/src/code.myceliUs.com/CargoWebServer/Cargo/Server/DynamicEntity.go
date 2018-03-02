package Server

import (
	"reflect"
	"strings"
	"sync"
)

/**
 * Implementation of the Entity. Dynamic entity use a map[string]interface {}
 * to store object informations. All information in the Entity itself other
 * than object are store.
 */
type DynamicEntity struct {

	/** object will be a map... **/
	object map[string]interface{}

	/** Get entity by uuid function **/
	getEntityByUuid func(string) (interface{}, error)

	/**
	 * Use to protected the ressource access...
	 */
	sync.RWMutex
}

func NewDynamicEntity() *DynamicEntity {
	entity := new(DynamicEntity)
	entity.object = make(map[string]interface{}, 0)
	return entity
}

/**
 * Thread safe function
 */
func (this *DynamicEntity) setValue(field string, value interface{}) error {
	this.Lock()
	defer this.Unlock()

	// Here the value is in the map.
	this.object[field] = value
	this.object["NeedSave"] = true
	return nil
}

/**
 * Thread safe function
 */
func (this *DynamicEntity) getValue(field string) interface{} {
	this.Lock()
	defer this.Unlock()
	return this.object[field]
}

/**
 * That function return a detach copy of the values contain in the object map.
 * Only the uuid of all child values are keep. That map can be use to save the content
 * of an entity in the cache.
 */
func (this *DynamicEntity) getValues() map[string]interface{} {
	this.Lock()
	defer this.Unlock()
	values := make(map[string]interface{})
	// return the values without all sub-entity values
	for k, v := range this.object {
		if v != nil {
			if reflect.TypeOf(v).String() == "[]interface {}" {
				if len(v.([]interface{})) > 0 {
					if reflect.TypeOf(v.([]interface{})[0]).String() == "map[string]interface {}" {
						if v.([]interface{})[0].(map[string]interface{})["UUID"] != nil {
							childs := make([]string, 0)
							for i := 0; i < len(v.([]interface{})); i++ {
								childs = append(childs, v.([]interface{})[0].(map[string]interface{})["UUID"].(string))
							}
							values[k] = childs
						} else {
							values[k] = v
						}
					} else {
						values[k] = v
					}
				}
			} else if reflect.TypeOf(v).String() == "map[string]interface {}" {
				if v.(map[string]interface{})["UUID"] != nil {
					values[k] = v.(map[string]interface{})["UUID"]
				} else {
					values[k] = v
				}
			} else {
				values[k] = v
			}
		} /*else {
			values[k] = v
		}*/
	}
	return values
}

/**
 * Thread safe function, apply only on none array field...
 */
func (this *DynamicEntity) deleteValue(field string) {
	this.Lock()
	defer this.Unlock()

	// Remove the field itself.
	delete(this.object, field)

}

/**
 * Set object.
 */
func (this *DynamicEntity) setObject(obj map[string]interface{}) {
	this.Lock()
	defer this.Unlock()
	this.object = obj
}

/**
 * Append a new value to array field
 */
func (this *DynamicEntity) appendValue(field string, value interface{}) {
	values := this.getValue(field)
	if values == nil {
		// Here no value aready exist.
		if reflect.TypeOf(value).Kind() == reflect.String {
			values = make([]string, 1)
			values.([]string)[0] = value.(string)
			this.setValue(field, values)
		} else {
			// Other types.
			values = make([]interface{}, 1)
			values.([]interface{})[0] = value
			this.setValue(field, values)
		}
	} else {
		// An array already exist in that case.
		if reflect.TypeOf(value).Kind() == reflect.String {
			if reflect.TypeOf(values).Kind().String() == "[]interface {}" {
				values = append(values.([]interface{}), value.(string))
			} else if reflect.TypeOf(values).Kind().String() == " []string" {
				values = append(values.([]string), value.(string))
			}
			this.setValue(field, values)
		} else {
			// Other types.
			values = append(values.([]interface{}), value)
			this.setValue(field, values)
		}
	}

}

/**
 * This is function is use to remove a child entity from it parent.
 * To remove other field type simply call 'setValue' with the new array values.
 */
func (this *DynamicEntity) removeValue(field string, uuid interface{}) {

	values_ := this.getValue(field)

	// Here no value aready exist.
	if reflect.TypeOf(values_).String() == "[]string" {
		values := make([]string, 0)
		for i := 0; i < len(values_.([]string)); i++ {
			if values_.([]string)[i] != uuid {
				values = append(values, values_.([]string)[i])
			}
		}
		this.setValue(field, values)
	} else if reflect.TypeOf(values_).String() == "[]interface {}" {
		values := make([]interface{}, 0)
		for i := 0; i < len(values_.([]interface{})); i++ {
			if values_.([]interface{})[i].(string) != uuid {
				values = append(values, values_.([]interface{})[i])
			}
		}
		this.setValue(field, values)
	} else if reflect.TypeOf(values_).String() == "[]map[string]interface {}" {
		values := make([]map[string]interface{}, 0)
		for i := 0; i < len(values_.([]map[string]interface{})); i++ {
			if values_.([]map[string]interface{})[i]["UUID"].(string) != uuid {
				values = append(values, values_.([]map[string]interface{})[i])
			}
		}
		this.setValue(field, values)
	}
}

/**
 * Return the internal object.
 */
func (this *DynamicEntity) getObject() map[string]interface{} {
	this.Lock()
	defer this.Unlock()
	return this.object
}

/**
 * Return the type name of an entity
 */
func (this *DynamicEntity) GetTypeName() string {
	return this.getValue("TYPENAME").(string)
}

/**
 * Each entity must have one uuid.
 */
func (this *DynamicEntity) GetUuid() string {
	if this.getValue("UUID") != nil {
		return this.getValue("UUID").(string)
	}
	return "" // Can be an error here.
}

/**
 * Each entity must have one uuid.
 */
func (this *DynamicEntity) ResetNeedSave() {
	this.setValue("NeedSave", false)
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *DynamicEntity) SetEntityGetter(fct func(uuid string) (interface{}, error)) {
	this.getEntityByUuid = fct
}

/**
 * Return the array of id's for a given entity, it not contain it UUID.
 */
func (this *DynamicEntity) Ids() []interface{} {
	typeName := this.GetTypeName()
	prototype, err := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	if err != nil {
		return nil
	}

	ids := make([]interface{}, 0)
	for i := 1; i < len(prototype.Ids); i++ {
		ids = append(ids, this.getValue(prototype.Ids[i]))
	}
	return ids
}

func (this *DynamicEntity) SetUuid(uuid string) {
	this.setValue("UUID", uuid)
}

func (this *DynamicEntity) GetParentUuid() string {
	if this.getValue("ParentUuid") != nil {
		return this.getValue("ParentUuid").(string)
	}
	return ""
}

func (this *DynamicEntity) SetParentUuid(parentUuid string) {
	this.setValue("ParentUuid", parentUuid)
}

/**
 * The name of the relation with it parent.
 */
func (this *DynamicEntity) GetParentLnk() string {
	parentLnk := this.getValue("ParentLnk")
	if parentLnk != nil {
		return parentLnk.(string)
	}
	return ""
}

func (this *DynamicEntity) SetParentLnk(lnk string) {
	this.setValue("ParentLnk", lnk)
}

/**
 * Test if an entity need to be save.
 */
func (this *DynamicEntity) IsNeedSave() bool {
	return this.getValue("NeedSave").(bool)
}
