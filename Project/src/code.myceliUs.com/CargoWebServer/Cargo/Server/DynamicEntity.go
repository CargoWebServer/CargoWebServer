package Server

import (
	"reflect"
	"strings"

	"code.myceliUs.com/Utility"
)

/**
 * Implementation of the Entity. Dynamic entity use a map[string]interface {}
 * to store object informations. All information in the Entity itself other
 * than object are store.
 */
type DynamicEntity struct {
	/** The entity uuid **/
	uuid string

	/** The entity typeName **/
	typeName string

	/** The parent uuid **/
	parentUuid string

	/** The relation with it parent **/
	parentLnk string

	/** Get entity by uuid function **/
	getEntityByUuid func(string) (interface{}, error)

	/** Put the entity on the cache **/
	setEntity func(interface{})

	/** Set the uuid function **/
	generateUuid func(interface{}) string
}

// Contain entity that need to be save.
var (
	needSave_ = make(map[string]bool)
)

func NewDynamicEntity() *DynamicEntity {
	entity := new(DynamicEntity)
	return entity
}

/**
 * Thread safe function
 */
func (this *DynamicEntity) setValue(field string, value interface{}) error {

	// Set the uuid if not already exist.
	// Now I will try to retreive the entity from the cache.
	infos := make(map[string]interface{})
	infos["name"] = "setValue"
	infos["uuid"] = this.uuid
	infos["field"] = field
	infos["value"] = value
	infos["needSave"] = make(chan bool)

	// set the values in the cache.
	cache.m_setValue <- infos

	needSave := <-infos["needSave"].(chan bool)
	if needSave {
		needSave_[this.uuid] = true
	}

	return nil
}

func (this *DynamicEntity) SetFieldValue(field string, value interface{}) error {
	needSave_[this.uuid] = true
	return this.setValue(field, value)
}

func (this *DynamicEntity) SetNeedSave(val bool) {
	if val {
		needSave_[this.uuid] = true
	} else {
		delete(needSave_, this.uuid)
	}
}

func (this *DynamicEntity) IsNeedSave() bool {
	if _, ok := needSave_[this.uuid]; ok {
		return true
	}
	return false
}

/**
 * Thread safe function
 */
func (this *DynamicEntity) getValue(field string) interface{} {
	// Set child uuid's here if there is not already sets...
	infos := make(map[string]interface{})
	infos["name"] = "getValue"
	infos["uuid"] = this.uuid
	infos["field"] = field
	infos["getValue"] = make(chan interface{})
	cache.m_getValue <- infos
	value := <-infos["getValue"].(chan interface{})
	return value
}

func (this *DynamicEntity) GetFieldValue(field string) interface{} {
	return this.getValue(field)
}

/**
 * Thread safe function, apply only on none array field...
 */
func (this *DynamicEntity) deleteValue(field string) {
	// Remove the field itself.
	infos := make(map[string]interface{})
	infos["name"] = "deleteValue"
	infos["uuid"] = this.uuid
	infos["field"] = field
	cache.m_deleteValue <- infos
}

/**
 * That function return a detach copy of the values contain in the object map.
 * Only the uuid of all child values are keep. That map can be use to save the content
 * of an entity in the cache.
 */
func (this *DynamicEntity) getValues() map[string]interface{} {
	// Set child uuid's here if there is not already sets...
	infos := make(map[string]interface{})
	infos["name"] = "getValues"
	infos["uuid"] = this.uuid
	infos["getValues"] = make(chan map[string]interface{})
	cache.m_getValues <- infos
	values := <-infos["getValues"].(chan map[string]interface{})
	return values
}

/**
 * Remove reference and create sub entity entity.
 */
func (this *DynamicEntity) setObject(obj map[string]interface{}) {

	// Here I will create the entity object.
	object := make(map[string]interface{})

	if obj["TYPENAME"] != nil {
		this.typeName = obj["TYPENAME"].(string)
	} else {
		return
	}

	if obj["UUID"] != nil {
		this.uuid = obj["UUID"].(string)
	}

	if obj["ParentUuid"] != nil {
		this.parentUuid = obj["ParentUuid"].(string)
	}

	if obj["ParentLnk"] != nil {
		this.parentLnk = obj["ParentLnk"].(string)
	}

	// Set uuid
	ids := make([]interface{}, 0)
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(this.typeName, this.typeName[0:strings.Index(this.typeName, ".")])
	// skip The first element, the uuid
	for i := 1; i < len(prototype.Ids); i++ {
		ids = append(ids, obj[prototype.Ids[i]])
	}
	obj["Ids"] = ids

	if len(this.uuid) == 0 {
		// In that case I will set it uuid.
		obj["UUID"] = generateEntityUuid(this.typeName, this.parentUuid, ids)
		this.uuid = obj["UUID"].(string)
	}

	// Here I will initilalyse sub-entity if there one, so theire map will not be part of this entity.
	for i := 0; i < len(prototype.Fields); i++ {
		field := prototype.Fields[i]
		val := obj[field]
		fieldType := prototype.FieldsType[i]
		if val != nil {
			if strings.HasPrefix(field, "M_") {
				if !strings.HasPrefix(fieldType, "[]xs.") && !strings.HasPrefix(fieldType, "xs.") {
					if strings.HasPrefix(fieldType, "[]") {
						if reflect.TypeOf(val).String() == "[]interface {}" {
							uuids := make([]interface{}, 0)
							for j := 0; j < len(val.([]interface{})); j++ {
								if val.([]interface{})[j] != nil {
									if reflect.TypeOf(val.([]interface{})[j]).String() == "map[string]interface {}" {
										if val.([]interface{})[j].(map[string]interface{})["TYPENAME"] != nil {
											// Set parent information if is not already set.
											if !strings.HasSuffix(fieldType, ":Ref") {
												val.([]interface{})[j].(map[string]interface{})["ParentUuid"] = obj["UUID"]
												val.([]interface{})[j].(map[string]interface{})["ParentLnk"] = field
												child := NewDynamicEntity()
												child.setObject(val.([]interface{})[j].(map[string]interface{}))
												GetServer().GetEntityManager().setEntity(child)
												val.([]interface{})[j].(map[string]interface{})["UUID"] = child.GetUuid()
											}
											// Keep it on the cache
											uuids = append(uuids, val.([]interface{})[j].(map[string]interface{})["UUID"])
										}
									} else {
										object[field] = val
									}
								}

							}
							if len(uuids) > 0 {
								object[field] = uuids
							}
						} else if reflect.TypeOf(val).String() == "[]map[string]interface {}" {
							uuids := make([]interface{}, 0)
							for j := 0; j < len(val.([]map[string]interface{})); j++ {
								if val.([]map[string]interface{})[j]["TYPENAME"] != nil {
									if !strings.HasSuffix(fieldType, ":Ref") {
										// Set parent information if is not already set.
										val.([]map[string]interface{})[j]["ParentUuid"] = obj["UUID"]
										val.([]map[string]interface{})[j]["ParentLnk"] = field
										child := NewDynamicEntity()
										child.setObject(val.([]map[string]interface{})[j])
										// keep the uuid in the map.
										val.([]map[string]interface{})[j]["UUID"] = child.GetUuid()
									}
									// Keep it on the cache
									uuids = append(uuids, val.([]map[string]interface{})[j]["UUID"])
								} else {
									object[field] = val
								}
							}
							if len(uuids) > 0 {
								object[field] = uuids
							}
						} else {
							object[field] = val
						}
					} else {
						if reflect.TypeOf(val).String() == "map[string]interface {}" {
							if val.(map[string]interface{})["TYPENAME"] != nil {
								if !strings.HasSuffix(fieldType, ":Ref") {
									// Set parent information if is not already set.
									val.(map[string]interface{})["ParentUuid"] = obj["UUID"]
									val.(map[string]interface{})["ParentLnk"] = field
									child := NewDynamicEntity()
									child.setObject(val.(map[string]interface{}))
									val.(map[string]interface{})["UUID"] = child.GetUuid()
								}
								// Keep it on the cache
								object[field] = val.(map[string]interface{})["UUID"]
							} else {
								object[field] = val
							}
						} else {
							object[field] = val
						}
					}
				} else {
					object[field] = val
				}
			} else {
				object[field] = val
			}
		} else {
			// Nil value...
			if strings.HasPrefix(fieldType, "[]") {
				// empty array...
				if strings.HasSuffix(fieldType, ":Ref") {
					object[field] = make([]string, 0)
				} else {
					object[field] = make([]interface{}, 0)
				}
			} else {
				object[field] = nil
			}
		}
	}

	object["UUID"] = this.uuid
	object["TYPENAME"] = this.typeName
	object["ParentUuid"] = this.parentUuid
	object["ParentLnk"] = this.parentLnk
	object["Ids"] = ids

	this.setValues(object)
}

/**
 * Save the map[string]interface {} for that entity.
 */
func (this *DynamicEntity) setValues(values map[string]interface{}) {
	// Set the uuid if not already exist.
	// Now I will try to retreive the entity from the cache.
	infos := make(map[string]interface{})
	infos["name"] = "setValues"
	infos["values"] = values
	infos["needSave"] = make(chan bool, 0)

	// set the values in the cache.
	cache.m_setValues <- infos

	// wait to see if the entity has change...
	this.SetNeedSave(<-infos["needSave"].(chan bool))
	LogInfo("---> entity ", this.uuid, "needSave", this.IsNeedSave())
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
 * Return the type name of an entity
 */
func (this *DynamicEntity) GetTypeName() string {
	return this.typeName
}

/**
 * Each entity must have one uuid.
 */
func (this *DynamicEntity) GetUuid() string {
	if len(this.uuid) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.uuid // Can be an error here.
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *DynamicEntity) SetEntityGetter(fct func(uuid string) (interface{}, error)) {
	this.getEntityByUuid = fct
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *DynamicEntity) SetEntitySetter(fct func(entity interface{})) {
	this.setEntity = fct
}

/** Set the uuid generator function **/
func (this *DynamicEntity) SetUuidGenerator(fct func(entity interface{}) string) {
	this.generateUuid = fct
}

/**
 * Return the array of id's for a given entity, it not contain it UUID.
 */
func (this *DynamicEntity) Ids() []interface{} {
	ids := make([]interface{}, 0)
	typeName := this.GetTypeName()
	prototype, err := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	if err != nil {
		return ids
	}
	// skip The first element, the uuid
	for i := 1; i < len(prototype.Ids); i++ {
		ids = append(ids, this.getValue(prototype.Ids[i]))
	}
	return ids
}

func (this *DynamicEntity) SetUuid(uuid string) {
	if this.uuid != uuid {
		this.setValue("UUID", uuid)
		this.uuid = uuid // Keep local...
	}
}

func (this *DynamicEntity) GetParentUuid() string {
	return this.parentUuid
}

func (this *DynamicEntity) SetParentUuid(parentUuid string) {
	if this.parentUuid != parentUuid {
		this.setValue("ParentUuid", parentUuid)
		this.parentUuid = parentUuid
	}
}

/**
 * The name of the relation with it parent.
 */
func (this *DynamicEntity) GetParentLnk() string {
	return this.parentLnk
}

func (this *DynamicEntity) SetParentLnk(lnk string) {
	if this.parentLnk != lnk {
		this.setValue("ParentLnk", lnk)
		this.parentLnk = lnk
	}
}

func (this *DynamicEntity) GetParent() interface{} {
	parent, err := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid())
	if err != nil {
		return err
	}
	return parent
}

/**
 * Return the list of all it childs.
 */
func (this *DynamicEntity) GetChilds() []interface{} {
	var childs []interface{}
	// I will get the childs by their uuid's if there already exist.
	uuids := this.GetChildsUuid()
	for i := 0; i < len(uuids); i++ {
		child, _ := GetServer().GetEntityManager().getEntityByUuid(uuids[i])
		if child != nil {
			childs = append(childs, child)
		}
	}
	return childs
}

/**
 * Return the list of all it childs.
 */
func (this *DynamicEntity) GetChildsUuid() []string {
	var childs []string
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(this.GetTypeName(), strings.Split(this.GetTypeName(), ".")[0])
	for i := 0; i < len(prototype.Fields); i++ {
		field := prototype.Fields[i]
		if strings.HasPrefix(field, "M_") {
			fieldType := prototype.FieldsType[i]
			if !strings.HasPrefix(fieldType, "[]xs.") && !strings.HasPrefix(fieldType, "xs.") && !strings.HasSuffix(fieldType, ":Ref") {
				val := this.getValue(field)
				if val != nil {
					if strings.HasPrefix(fieldType, "[]") {
						// The value is an array...
						if reflect.TypeOf(val).String() == "[]interface {}" {
							for j := 0; j < len(val.([]interface{})); j++ {
								if Utility.IsValidEntityReferenceName(val.([]interface{})[j].(string)) {
									childs = append(childs, val.([]interface{})[j].(string))
								}
							}
						} else if reflect.TypeOf(val).String() == "[]string" {
							for j := 0; j < len(val.([]string)); j++ {
								if Utility.IsValidEntityReferenceName(val.([]string)[j]) {
									childs = append(childs, val.([]string)[j])
								}
							}
						}
					} else {
						// The value is not an array.
						if Utility.IsValidEntityReferenceName(val.(string)) {
							childs = append(childs, val.(string))
						}
					}
				}
			}
		}
	}

	return childs
}

/**
 * Return the list of reference UUID:FieldName where that entity is referenced.
 */
func (this *DynamicEntity) GetReferenced() []string {
	referenced := this.getValue("Referenced").([]string)
	if referenced == nil {
		referenced = make([]string, 0)
		this.setValue("Referenced", referenced)
	}

	// return the list of references
	return referenced
}

/**
 * Set a reference
 */
func (this *DynamicEntity) SetReferenced(uuid string, field string) {
	referenced := this.GetReferenced()
	if !Utility.Contains(referenced, uuid+":"+field) {
		referenced = append(referenced, uuid+":"+field)
		this.setValue("Referenced", referenced)
	}
}

/**
 * Remove a reference.
 */
func (this *DynamicEntity) RemoveReferenced(uuid string, field string) {
	referenced := this.GetReferenced()
	referenced_ := make([]string, 0)
	for i := 0; i < len(referenced); i++ {
		if referenced[i] != uuid+":"+field {
			referenced_ = append(referenced_, referenced[i])
		}
	}

	// Set back the value without the uuid:field.
	this.setValue("Referenced", referenced_)
}
