package Server

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"reflect"
	"sort"
	"strings"

	"compress/gzip"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

////////////////////////////////////////////////////////////////////////////////
//						Entity Manager
////////////////////////////////////////////////////////////////////////////////

type EntityManager struct {
	// Cache the entitie in memory...
	m_cache *Cache
}

var entityManager *EntityManager

// Function to be use a function pointer.
var getEntityFct func(uuid string) (interface{}, error)
var setEntityFct func(interface{})

// Function to set the entity uuid.
var (
	generateUuidFct = func(entity interface{}) string {
		// Here If the entity id's is not set I will set it...
		if entity.(Entity).GetFieldValue("UUID") != nil {
			if len(entity.(Entity).GetFieldValue("UUID").(string)) > 0 {
				return entity.(Entity).GetFieldValue("UUID").(string)
			}
		}

		uuid := generateEntityUuid(entity.(Entity).GetTypeName(), entity.(Entity).GetParentUuid(), entity.(Entity).Ids())
		return uuid
	}
)

func generateEntityUuid(typeName string, parentUuid string, ids []interface{}) string {

	if len(ids) == 0 {
		// if there is no ids in the entity I will generate a random uuid.
		return typeName + "%" + Utility.RandomUUID()
	}

	var keyInfo string
	keyInfo = typeName + ":"

	if len(parentUuid) > 0 {
		keyInfo += parentUuid + ":"
	}

	for i := 0; i < len(ids); i++ {
		if ids[i] != nil {
			keyInfo += Utility.ToString(ids[i])
			// Append underscore for readability in case of problem...
			if i < len(ids)-1 {
				keyInfo += "_"
			}
		}
	}

	uuid := typeName + "%" + Utility.GenerateUUID(keyInfo)

	// Return the uuid from the input information.
	return uuid
}

func (this *Server) GetEntityManager() *EntityManager {
	if entityManager == nil {
		entityManager = newEntityManager()
	}

	return entityManager
}

func newEntityManager() *EntityManager {

	entityManager = new(EntityManager)

	// Create prototypes for config objects and entities objects...
	entityManager.registerConfigObjects()
	entityManager.createConfigPrototypes()

	entityManager.registerCargoEntitiesObjects()
	entityManager.createCargoEntitiesPrototypes()

	// Entity prototype is a dynamic type.
	// ** Dynamic type must have the TYPENAME property!
	Utility.RegisterType((*EntityPrototype)(nil))
	Utility.RegisterType((*Restriction)(nil))
	Utility.RegisterType((*DynamicEntity)(nil))
	Utility.RegisterType((*MessageData)(nil))
	Utility.RegisterType((*TaskInstanceInfo)(nil))
	Utility.RegisterType((*EntityQuery)(nil))

	// Set the get Entity function
	getEntityFct = func(uuid string) (interface{}, error) {
		entity, err := entityManager.getEntityByUuid(uuid)
		if err != nil {
			return nil, errors.New(err.GetBody())
		}
		return entity, nil
	}

	setEntityFct = func(entity interface{}) {
		if v, ok := entity.(Entity); ok {
			GetServer().GetEntityManager().setEntity(v)
		}
	}

	// The Cache...
	entityManager.m_cache = newCache()

	return entityManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Initialization.
 */
func (this *EntityManager) initialize() {
	log.Println("--> Initialize EntityManager")

	// Create the default configurations
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)

	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})

	// must be call at least once at start time.
	this.getCargoEntities()
}

func (this *EntityManager) getId() string {
	return "EntityManager"
}

func (this *EntityManager) start() {
	log.Println("--> Start EntityManager")
}

func (this *EntityManager) stop() {
	log.Println("--> Stop EntityManager")
}

// Return the default Cargo Entities, create it if is not already exist.
func (this *EntityManager) getCargoEntities() *CargoEntities.Entities {
	var cargoEntities *CargoEntities.Entities
	entities, _ := GetServer().GetEntityManager().getEntityById("CargoEntities.Entities", "CargoEntities", []interface{}{"CARGO_ENTITIES"})
	if entities != nil {
		cargoEntities = entities.(*CargoEntities.Entities)
	} else {
		// I will create the cargo entities if it dosent already exist.
		cargoEntities = new(CargoEntities.Entities)
		cargoEntities.SetId("CARGO_ENTITIES")
		cargoEntities.SetName("Cargo entities")
		cargoEntities.SetVersion("1.0")
		cargoEntities.SetEntityGetter(getEntityFct)
		cargoEntities.SetEntitySetter(setEntityFct)
		cargoEntities.SetUuidGenerator(generateUuidFct)
		this.saveEntity(cargoEntities)
	}
	return cargoEntities
}

// Return the uuid for the CargoEntities
func (this *EntityManager) getCargoEntitiesUuid() string {
	return generateEntityUuid("CargoEntities.Entities", "", []interface{}{"CARGO_ENTITIES"})
}

func (this *EntityManager) getEntityPrototype(typeName string, storeId string) (*EntityPrototype, error) {
	store := GetServer().GetDataManager().getDataStore(storeId)
	if store != nil {
		prototype, err := store.GetEntityPrototype(typeName)
		return prototype, err
	}
	return nil, errors.New("No Data store found with id: " + storeId)
}

func (this *EntityManager) getEntityOwner(entity Entity) Entity {
	log.Println("getEntityOwner")
	return nil
}

func (this *EntityManager) isEntityExist(uuid string) bool {
	log.Println("isEntityExist")
	return false
}

func (this *EntityManager) getEntities(typeName string, storeId string, query *EntityQuery) ([]Entity, *CargoEntities.Error) {

	var q string
	if query == nil {
		query = new(EntityQuery)
		query.TYPENAME = "Server.EntityQuery"
		query.TypeName = typeName
		query.Query = query.TypeName + `.TYPENAME=="` + query.TypeName + `"`
	}

	// need all fields...
	query.Fields = make([]string, 0)

	val, err := json.Marshal(query)
	if err == nil {
		q = string(val)
	}

	store := GetServer().GetDataManager().getDataStore(storeId)
	results, err := store.Read(q, []interface{}{}, []interface{}{})

	if err != nil {
		errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Fail to retreive entity by id "))
		return nil, errObj
	}

	entities := make([]Entity, 0)

	for i := 0; i < len(results); i++ {
		if results[i][0] != nil {
			if reflect.TypeOf(results[i][0]).Kind() == reflect.Ptr || reflect.TypeOf(results[i][0]).Kind() == reflect.Struct || reflect.TypeOf(results[i][0]).Kind() == reflect.Map {
				obj, err := Utility.InitializeStructure(results[i][0].(map[string]interface{}), setEntityFct)
				if err == nil {
					var entity Entity
					// So here I will retreive the entity uuid from the entity id.
					if reflect.TypeOf(obj.Interface()).String() != "map[string]interface {}" {
						entity = obj.Interface().(Entity)
						// Set the entity in the cache
						this.setEntity(entity)
					} else {
						// Dynamic entity here.
						entity = NewDynamicEntity()
						entity.(*DynamicEntity).setObject(results[i][0].(map[string]interface{}))
						// set the entity in the cache.
						this.setEntity(entity)
					}
					entities = append(entities, entity)
				}
			} else {
				log.Println("----> results not an entity: ", results[i])
			}
		}
	}

	return entities, nil
}

func (this *EntityManager) removeEntity(entity Entity) {
	infos := make(map[string]interface{})
	infos["name"] = "removeEntity"
	infos["entity"] = entity

	// set the entity
	this.m_cache.m_removeEntity <- infos
}

func (this *EntityManager) setEntity(entity Entity) {
	// Set the entity function...
	entity.SetEntityGetter(getEntityFct)
	entity.SetEntitySetter(setEntityFct)
	entity.SetUuidGenerator(generateUuidFct)

	// When call those function also set the value for the field.
	entity.GetTypeName()
	entity.GetUuid()

	// Set the uuid if not already exist.
	// Now I will try to retreive the entity from the cache.
	infos := make(map[string]interface{})
	infos["name"] = "setEntity"
	infos["entity"] = entity

	// set the entity
	this.m_cache.m_setEntity <- infos
}

func (this *EntityManager) getEntity(uuid string) Entity {
	infos := make(map[string]interface{})
	infos["name"] = "getEntity"
	infos["uuid"] = uuid
	infos["getEntity"] = make(chan Entity)
	this.m_cache.m_getEntity <- infos
	entity := <-infos["getEntity"].(chan Entity)

	// Set the entity function pointer...
	if entity != nil {
		entity.SetEntityGetter(getEntityFct)
		entity.SetEntitySetter(setEntityFct)
		entity.SetUuidGenerator(generateUuidFct)
	}

	return entity
}

func (this *EntityManager) getEntityByUuid(uuid string) (Entity, *CargoEntities.Error) {

	if len(uuid) == 0 {
		errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("No uuid given!"))
		return nil, errObj
	}
	// Get it from cache if is there.
	entity := this.getEntity(uuid)
	if entity != nil {
		// that function must be set back.
		return entity, nil
	}

	if !Utility.IsValidEntityReferenceName(uuid) {
		return nil, NewError(Utility.FileLine(), INVALID_REFERENCE_NAME_ERROR, SERVER_ERROR_CODE, errors.New(uuid+" is not a valid reference!"))
	}

	// Todo get the entity from the datastore.
	typeName := strings.Split(uuid, "%")[0]
	storeId := typeName[0:strings.Index(typeName, ".")]

	// So now I will retreive the list of values associated with that uuid.
	var q EntityQuery
	q.TYPENAME = "Server.EntityQuery"
	q.TypeName = typeName
	q.Query = typeName + `.UUID=="` + uuid + `"`
	query, _ := json.Marshal(q)

	store := GetServer().GetDataManager().getDataStore(storeId)
	results, err := store.Read(string(query), []interface{}{}, []interface{}{})

	if err != nil {
		errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Fail to retreive entity with uuid "+uuid))
		return nil, errObj
	}

	obj, _ := Utility.InitializeStructure(results[0][0].(map[string]interface{}), setEntityFct)

	// So here I will retreive the entity uuid from the entity id.
	if obj.IsValid() {
		if obj.IsNil() {
			errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Fail to retreive entity with uuid "+uuid))
			return nil, errObj
		} else {

			if obj.Type().String() != "map[string]interface {}" {
				entity = obj.Interface().(Entity)
				entity.SetNeedSave(false)
				// Set the entity in the cache
				this.setEntity(entity)
			} else if obj.Type().String() == "map[string]interface {}" {
				// Dynamic entity here.
				entity = NewDynamicEntity()
				entity.(*DynamicEntity).setObject(results[0][0].(map[string]interface{}))
				// set the entity in the cache.
				entity.SetNeedSave(false)
				this.setEntity(entity)
			}
		}
	} else {
		errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Fail to retreive entity with uuid "+uuid))
		return nil, errObj
	}

	return entity, nil
}

func (this *EntityManager) getEntityById(typeName string, storeId string, ids []interface{}) (Entity, *CargoEntities.Error) {
	if len(ids) == 0 {
		errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("No ids given!"))
		return nil, errObj
	}

	// here the id is generated with parent uuid.
	uuid, errObj := this.getEntityUuidById(typeName, storeId, ids)
	if errObj == nil {
		return this.getEntityByUuid(uuid)
	}

	return nil, errObj
}

func (this *EntityManager) getEntityUuidById(typeName string, storeId string, ids []interface{}) (string, *CargoEntities.Error) {

	// So here I will retreive the entity uuid from the entity id.
	prototype, _ := this.getEntityPrototype(typeName, storeId)
	if prototype == nil {
		errObj := NewError(Utility.FileLine(), PROTOTYPE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Prototype with name "+typeName+" not found!"))
		return "", errObj
	}

	if len(prototype.Ids) <= 1 {
		errObj := NewError(Utility.FileLine(), PROTOTYPE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("No id define for the type "+typeName))
		return "", errObj

	}

	// First of all I will get the uuid...
	fieldType := prototype.FieldsType[prototype.getFieldIndex(prototype.Ids[1])]
	fieldType = strings.Replace(fieldType, "[]", "", -1)
	fieldType = strings.Replace(fieldType, ":Ref", "", -1)

	var q EntityQuery
	q.TYPENAME = "Server.EntityQuery"
	q.Fields = append(q.Fields, "UUID")
	q.TypeName = prototype.TypeName

	// The EQL query here.
	for i := 0; i < len(ids); i++ {
		q.Query += prototype.TypeName + `.` + prototype.Ids[i+1] + `==`
		if reflect.TypeOf(ids[i]).Kind() == reflect.String {
			q.Query += `"` + ids[i].(string) + `"`
		} else {
			q.Query += Utility.ToString(ids[i])
		}

		if i < len(ids)-1 && len(ids) > 1 {
			q.Query += " && "
		}
	}

	// Make the query over the store...
	store := GetServer().GetDataManager().getDataStore(storeId)
	query, _ := json.Marshal(q)
	results, err := store.Read(string(query), []interface{}{}, []interface{}{})

	if err != nil || len(results) == 0 {
		var idsStr string
		for i := 0; i < len(ids); i++ {
			idsStr += Utility.ToString(ids[i])
			if i < len(ids)-1 && i > 0 {
				idsStr += ", "
			}
		}

		errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Fail to retreive entity with id(s) "+idsStr))
		return "", errObj
	}

	return results[0][0].(string), nil
}

func (this *EntityManager) setParent(parent Entity, entity Entity) *CargoEntities.Error {

	var cargoError *CargoEntities.Error
	var parentPrototype *EntityPrototype

	parentPrototype, _ = GetServer().GetEntityManager().getEntityPrototype(parent.GetTypeName(), parent.GetTypeName()[0:strings.Index(parent.GetTypeName(), ".")])
	fieldIndex := parentPrototype.getFieldIndex(entity.GetParentLnk())
	if fieldIndex == -1 {
		// It happen when the parent is reattributed...
		return nil // Nothing todo here
	}

	fieldType := parentPrototype.FieldsType[fieldIndex]
	if cargoError != nil {
		return cargoError
	}

	// Get values as map[string]interface{} and also set the entity in it parent.
	if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
		if strings.HasPrefix(fieldType, "[]") {
			existingValues := make([]interface{}, 0)
			if parent.(*DynamicEntity).getValue(entity.GetParentLnk()) != nil {
				existingValues = parent.(*DynamicEntity).getValue(entity.GetParentLnk()).([]interface{})
			}
			exist := false
			for i := 0; i < len(existingValues); i++ {
				if existingValues[i].(string) == entity.(*DynamicEntity).GetUuid() {
					exist = true
					break
				}
			}
			if !exist {
				existingValues = append(existingValues, entity.(*DynamicEntity).GetUuid())
				parent.(*DynamicEntity).setValue(entity.GetParentLnk(), existingValues)
			}

		} else {
			parent.(*DynamicEntity).setValue(entity.GetParentLnk(), entity.(*DynamicEntity).GetUuid())
		}
	} else {
		setMethodName := strings.Replace(entity.GetParentLnk(), "M_", "", -1)
		if strings.HasPrefix(fieldType, "[]") {
			setMethodName = "Append" + strings.ToUpper(setMethodName[0:1]) + setMethodName[1:]
		} else {
			setMethodName = "Set" + strings.ToUpper(setMethodName[0:1]) + setMethodName[1:]
		}
		params := make([]interface{}, 1)
		params[0] = entity

		_, err_ := Utility.CallMethod(parent, setMethodName, params)
		//log.Println("---> call method ", parent.GetUuid(), setMethodName, parent.IsNeedSave())
		if err_ != nil {
			log.Println("fail to call method ", setMethodName, " on ", parent)
			cargoError := NewError(Utility.FileLine(), ATTRIBUTE_NAME_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, err_.(error))
			return cargoError
		}
	}

	//this.saveEntity(parent)
	if parent.IsNeedSave() {
		this.setEntity(parent)
		typeName := parent.GetTypeName()
		storeId := typeName[0:strings.Index(typeName, ".")]
		store := GetServer().GetDataManager().getDataStore(storeId)
		err := store.Update("", []interface{}{parent}, []interface{}{})
		if err != nil {
			cargoError := NewError(Utility.FileLine(), ENTITY_CREATION_ERROR, SERVER_ERROR_CODE, err)
			return cargoError
		}
	}
	return nil
}

func (this *EntityManager) saveChilds(entity Entity, prototype *EntityPrototype) {
	// The list of childs.
	childs := entity.GetChilds()
	// Save the childs...
	for i := 0; i < len(childs); i++ {
		childs[i].(Entity).SetParentUuid(entity.GetUuid())
		this.saveEntity(childs[i].(Entity))
	}
}

/**
 * Create an new entity.
 */
func (this *EntityManager) createEntity(parent Entity, attributeName string, entity Entity) (Entity, *CargoEntities.Error) {

	// Set the entity values here.
	typeName := entity.GetTypeName() // Set the type name if not already set...
	entity.SetParentLnk(attributeName)
	entity.SetParentUuid(parent.GetUuid())

	// Here I will set the entity on the cache...
	this.setEntity(entity)

	storeId := typeName[0:strings.Index(typeName, ".")]
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, storeId)

	// Now entity are quadify I will save it in the graph store.
	store := GetServer().GetDataManager().getDataStore(storeId)

	// I will create the entity.
	_, err := store.Create("", []interface{}{entity})
	if err != nil {
		cargoError := NewError(Utility.FileLine(), ENTITY_CREATION_ERROR, SERVER_ERROR_CODE, err)
		return nil, cargoError
	}

	//log.Println("---> create entity ", entity.GetUuid())
	entity.SetNeedSave(false)

	// also save it parent.
	if parent != nil {
		// Now I will save it in the datastore.
		// I will set the entity parent.
		cargoError := this.setParent(parent, entity)
		if cargoError != nil {
			return nil, cargoError
		}
	} else {
		cargoError := NewError(Utility.FileLine(), ENTITY_CREATION_ERROR, SERVER_ERROR_CODE, errors.New("parent must not be nil when createEntity is call."))
		return nil, cargoError
	}

	// The event data...
	eventData := make([]*MessageData, 2)
	msgData0 := new(MessageData)
	msgData0.Name = "entity"
	if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
		msgData0.Value = entity.(*DynamicEntity).getValues()
	} else {
		msgData0.Value = entity
	}
	eventData[0] = msgData0

	msgData1 := new(MessageData)
	msgData1.Name = "prototype"
	msgData1.Value = prototype
	eventData[1] = msgData1

	evt, _ := NewEvent(NewEntityEvent, EntityEvent, eventData)
	GetServer().GetEventManager().BroadcastEvent(evt)

	// Set it childs...
	this.saveChilds(entity, prototype)

	return entity, nil
}

func (this *EntityManager) entityExist(uuid string) bool {
	typeName := strings.Split(uuid, "%")[0]
	storeId := typeName[0:strings.Index(typeName, ".")]

	var query EntityQuery
	query.TYPENAME = typeName
	query.TypeName = typeName
	query.Fields = []string{"UUID"}
	query.Query = typeName + `.UUID=="` + uuid + `"`

	store := GetServer().GetDataManager().getDataStore(storeId)
	queryStr, _ := json.Marshal(query)
	_, err := store.Read(string(queryStr), []interface{}{}, []interface{}{})
	if err != nil {
		return false
	}
	return true
}

func (this *EntityManager) saveEntity(entity Entity) *CargoEntities.Error {
	// So now I will retreive the list of values associated with that uuid.
	// If the to string are the same no save is needed.
	// log.Println("---> save entity: ", entity.GetUuid())
	if this.entityExist(entity.GetUuid()) {
		if !entity.IsNeedSave() {
			// log.Println("--->not need save entity ", entity.GetUuid())
			return nil
		}
	}

	// Here I will set the entity on the cache...
	this.setEntity(entity)

	typeName := entity.GetTypeName() // Set the type name if not already set...
	storeId := typeName[0:strings.Index(typeName, ".")]
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, storeId)
	eventData := make([]*MessageData, 2)
	msgData0 := new(MessageData)
	msgData0.Name = "entity"
	if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
		msgData0.Value = entity.(*DynamicEntity).getValues()
		//log.Println("---> save entity ", toJsonStr(entity.(*DynamicEntity).getValues()))
	} else {
		msgData0.Value = entity
	}
	eventData[0] = msgData0

	msgData1 := new(MessageData)
	msgData1.Name = "prototype"
	msgData1.Value = prototype
	eventData[1] = msgData1

	// Send save event if something has change.
	// I will get the existing value from the datastore is it exist.
	var evt *Event
	var err error
	store := GetServer().GetDataManager().getDataStore(storeId)
	if this.entityExist(entity.GetUuid()) {
		// I will create the entity.
		err = store.Update("", []interface{}{entity}, []interface{}{})
		if err != nil {
			cargoError := NewError(Utility.FileLine(), ENTITY_CREATION_ERROR, SERVER_ERROR_CODE, err)
			return cargoError
		}
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
	} else {
		_, err = store.Create("", []interface{}{entity})
		if err != nil {
			cargoError := NewError(Utility.FileLine(), ENTITY_CREATION_ERROR, SERVER_ERROR_CODE, err)
			return cargoError
		}
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
	}

	// Send update entity event here.
	GetServer().GetEventManager().BroadcastEvent(evt)

	// Set need save to false.
	entity.SetNeedSave(false)

	// also save it parent.
	if len(entity.GetParentUuid()) > 0 {
		parent, cargoError := this.getEntityByUuid(entity.GetParentUuid())
		if cargoError == nil {
			// Now I will save it in the datastore.
			// I will set the entity parent.
			cargoError := this.setParent(parent, entity)
			if cargoError != nil {
				return cargoError
			}
		} else {
			return cargoError
		}
	}

	// Set it childs...
	this.saveChilds(entity, prototype)

	return nil
}

func (this *EntityManager) deleteEntity(entity Entity) *CargoEntities.Error {
	storeId := entity.GetTypeName()[0:strings.Index(entity.GetTypeName(), ".")]
	store := GetServer().GetDataManager().getDataStore(storeId)
	uuids := make([]string, 0)
	uuids = append(uuids, entity.GetUuid())

	// Now I will clear other references...
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(entity.GetTypeName(), storeId)

	// I will remove it references...
	toSave := make([]Entity, 0)

	referenced := entity.GetReferenced()
	for i := 0; i < len(referenced); i++ {
		values := strings.Split(referenced[i], ":")
		refOwer, err := this.getEntityByUuid(values[0])
		if err == nil {
			if reflect.TypeOf(refOwer).String() == "*Server.DynamicEntity" {
				refOwer.(*DynamicEntity).removeValue(values[1], entity.GetUuid())
			} else {
				removeMethodName := strings.Replace(values[1], "M_", "", -1)
				removeMethodName = "Remove" + strings.ToUpper(removeMethodName[0:1]) + removeMethodName[1:]
				params := make([]interface{}, 1)
				params[0] = entity
				_, err_ := Utility.CallMethod(refOwer, removeMethodName, params)
				if err_ != nil {
					cargoError := NewError(Utility.FileLine(), ATTRIBUTE_NAME_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, err_.(error))
					return cargoError
				}
			}
		}
		toSave = append(toSave, refOwer)
	}

	// Remove it child.
	for i := 0; i < len(entity.GetChildsUuid()); i++ {
		child, err := this.getEntityByUuid(entity.GetChildsUuid()[i])
		if err == nil {
			this.deleteEntity(child)
		}
	}

	if len(entity.GetParentUuid()) > 0 {
		parent, err := this.getEntityByUuid(entity.GetParentUuid())
		if err == nil {
			if reflect.TypeOf(parent).String() == "*Server.DynamicEntity" {
				parent.(*DynamicEntity).removeValue(entity.GetParentLnk(), entity.GetUuid())
			} else {
				removeMethodName := strings.Replace(entity.GetParentLnk(), "M_", "", -1)
				removeMethodName = "Remove" + strings.ToUpper(removeMethodName[0:1]) + removeMethodName[1:]
				params := make([]interface{}, 1)
				params[0] = entity
				_, err_ := Utility.CallMethod(parent, removeMethodName, params)
				if err_ != nil {
					log.Println("fail to call method ", removeMethodName, " on ", parent)
					cargoError := NewError(Utility.FileLine(), ATTRIBUTE_NAME_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, err_.(error))
					return cargoError
				}
			}
			toSave = append(toSave, parent)
		}
	}

	// I will remove it data.
	store.Delete("", []interface{}{entity.GetUuid()})

	// Now I will remove the values from the cache.
	for i := 0; i < len(uuids); i++ {
		infos := make(map[string]interface{})
		infos["name"] = "remove"
		infos["uuid"] = uuids[i]
		// set the entity
		this.m_cache.m_remove <- infos

		// also remove it by it ids indexation.
		if len(entity.Ids()) > 0 {
			id := generateEntityUuid(entity.GetTypeName(), "", entity.Ids())
			infos := make(map[string]interface{})
			infos["name"] = "remove"
			infos["uuid"] = id
		}

		this.m_cache.m_remove <- infos
	}

	// save releated entities.
	for i := 0; i < len(toSave); i++ {
		this.saveEntity(toSave[i])
	}

	// Send event message...
	var eventDatas []*MessageData
	evtData_0 := new(MessageData)
	evtData_0.TYPENAME = "Server.MessageData"
	evtData_0.Name = "entity"
	// I will send only necessary entity properties.
	evtData_0.Value = map[string]interface{}{"UUID": entity.GetUuid(), "TYPENAME": entity.GetTypeName()}

	eventDatas = append(eventDatas, evtData_0)

	evtData_1 := new(MessageData)
	evtData_1.TYPENAME = "Server.MessageData"
	evtData_1.Name = "prototype"
	// I will send only necessary entity properties.
	evtData_1.Value = prototype
	eventDatas = append(eventDatas, evtData_1)

	evt, _ := NewEvent(DeleteEntityEvent, EntityEvent, eventDatas)
	GetServer().GetEventManager().BroadcastEvent(evt)

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

////////////////////////////// Prototypes //////////////////////////////////////

// @api 1.0
// Create a new entity prototype.
// @param {string} storeId The store id, where to create the new prototype.
// @param {interface{}} prototype The prototype object to create.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{*EntityPrototype} Return the created entity prototype
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
// EntityManager.prototype.createEntityPrototype = function(storeId, prototype, successCallback, errorCallback, caller){
//	var params = []
//	params.push(createRpcData(storeId, "STRING", "storeId"))
//	params.push(createRpcData(prototype, "JSON_STR", "prototype"))
//	server.executeJsFunction(
//	"EntityManagerCreateEntityPrototype",
//	params,
//	undefined, //progress callback
//	function (results, caller) { // Success callback
// 	   if(caller.successCallback!=undefined){
// 			var prototype = new EntityPrototype();
//			prototype.init(results[0]);
//      		caller.successCallback(prototype, caller.caller);
//          caller.successCallback = undefined;
//		}
//	},
//	function (errMsg, caller) { // Error callback
//          server.errorManager.onError(errMsg);
//         	if( caller.errorCallback != undefined){
//          		caller.errorCallback(errMsg, caller.caller);
//				caller.errorCallback = undefined;
//			}
//	},{"successCallback":successCallback, "errorCallback":errorCallback, "caller": caller})
//}
func (this *EntityManager) CreateEntityPrototype(storeId string, prototype interface{}, messageId string, sessionId string) *EntityPrototype {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// Cast it as needed...
	if reflect.TypeOf(prototype).String() == "map[string]interface {}" {
		prototype.(map[string]interface{})["TYPENAME"] = "Server.EntityPrototype"
		values, err := Utility.InitializeStructure(prototype.(map[string]interface{}), setEntityFct)
		if err == nil {
			prototype = values.Interface()
		}
	}

	if reflect.TypeOf(prototype).String() != "*Server.EntityPrototype" {
		cargoError := NewError(Utility.FileLine(), PARAMETER_TYPE_ERROR, SERVER_ERROR_CODE, errors.New("Expected '*Server.EntityPrototype' but got '"+reflect.TypeOf(prototype).String()+"' instead."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	// Get the store...
	store := GetServer().GetDataManager().getDataStore(storeId)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Datastore '"+storeId+"' dosen't exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil

	}

	// Save the prototype...
	err := store.CreateEntityPrototype(prototype.(*EntityPrototype))

	if err != nil {
		cargoError := NewError(Utility.FileLine(), PROTOTYPE_CREATION_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	return prototype.(*EntityPrototype)
}

// @api 1.0
// Save existing entity prototype.
// @param {string} storeId The store id, where to create the new prototype.
// @param {interface{}} prototype The prototype object to create.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{*EntityPrototype} Return the saved entity prototype
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
// EntityManager.prototype.saveEntityPrototype = function(storeId, prototype, successCallback, errorCallback, caller){
//	var params = []
//	params.push(createRpcData(storeId, "STRING", "storeId"))
//	params.push(createRpcData(prototype, "JSON_STR", "prototype"))
//	server.executeJsFunction(
//	"EntityManagerSaveEntityPrototype",
//	params,
//	undefined, //progress callback
//	function (results, caller) { // Success callback
// 	   if(caller.successCallback!=undefined){
// 			 var prototype = new EntityPrototype();
//			 prototype.init(results[0]);
//      	 	 caller.successCallback(prototype, caller.caller);
//           caller.successCallback = undefined;
//		}
//	},
//	function (errMsg, caller) { // Error callback
//          server.errorManager.onError(errMsg);
//         	if( caller.errorCallback != undefined){
//          		caller.errorCallback(errMsg, caller.caller);
//				caller.errorCallback = undefined;
//			}
//	},{"successCallback":successCallback, "errorCallback":errorCallback, "caller": caller})
//}
func (this *EntityManager) SaveEntityPrototype(storeId string, prototype interface{}, messageId string, sessionId string) *EntityPrototype {

	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// Cast it as needed...
	if reflect.TypeOf(prototype).String() == "map[string]interface {}" {
		prototype.(map[string]interface{})["TYPENAME"] = "Server.EntityPrototype"
		values, err := Utility.InitializeStructure(prototype.(map[string]interface{}), setEntityFct)
		if err == nil {
			prototype = values.Interface()
		} else {
			log.Println("fail to initialyse EntityPrototype from map[string]interface {} ", err)
			cargoError := NewError(Utility.FileLine(), PARAMETER_TYPE_ERROR, SERVER_ERROR_CODE, err)
			GetServer().reportErrorMessage(messageId, sessionId, cargoError)
			return nil
		}
	}

	if reflect.TypeOf(prototype).String() != "*Server.EntityPrototype" {
		cargoError := NewError(Utility.FileLine(), PARAMETER_TYPE_ERROR, SERVER_ERROR_CODE, errors.New("Expected '*Server.EntityPrototype' but got '"+reflect.TypeOf(prototype).String()+"' instead."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	// Get the store...
	store := GetServer().GetDataManager().getDataStore(storeId)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Datastore '"+storeId+"' dosen't exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil

	}

	// Save the prototype...
	err := store.SaveEntityPrototype(prototype.(*EntityPrototype))
	if err != nil {
		cargoError := NewError(Utility.FileLine(), PROTOTYPE_UPDATE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	return prototype.(*EntityPrototype)
}

// @api 1.0
// Delete existing entity prototype.
// @param {string} typeName The prototype id.
// @param {string} storeId The store id, where to create the new prototype.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *EntityManager) DeleteEntityPrototype(typeName string, storeId string, messageId string, sessionId string) {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	prototype, err := this.getEntityPrototype(typeName, storeId)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	// Get the store...
	store := GetServer().GetDataManager().getDataStore(storeId)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Datastore '"+storeId+"' dosen't exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}

	err = store.DeleteEntityPrototype(prototype.TypeName)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), PROTOTYPE_DELETE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
	}

}

// @api 1.0
// Rename existing entity prototype.
// @param {string} typeName The new prototype name.
// @param {string} prototype The prototype to rename.
// @param {string} storeId The store id, where to create the new prototype.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{*EntityPrototype} Return the renamed entity prototype
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
// EntityManager.prototype.renameEntityPrototype = function(typeName, prototype, storeId, successCallback, errorCallback, caller){
//	var params = []
//	params.push(createRpcData(typeName, "STRING", "typeName"))
//	params.push(createRpcData(prototype, "JSON_STR", "prototype"))
//	params.push(createRpcData(storeId, "STRING", "storeId"))
//	server.executeJsFunction(
//	"EntityManagerRenameEntityPrototype",
//	params,
//	undefined, //progress callback
//	function (results, caller) { // Success callback
// 	   if(caller.successCallback!=undefined){
// 			 var prototype = new EntityPrototype();
//			 prototype.init(results[0]);
//      	 	 caller.successCallback(prototype, caller.caller);
//           caller.successCallback = undefined;
//		}
//	},
//	function (errMsg, caller) { // Error callback
//          server.errorManager.onError(errMsg)
//         	if( caller.errorCallback != undefined){
//          		caller.errorCallback(errMsg, caller.caller);
//				caller.errorCallback = undefined;
//			}
//	},{"successCallback":successCallback, "errorCallback":errorCallback, "caller": caller})
//}
func (this *EntityManager) RenameEntityPrototype(typeName string, prototype interface{}, storeId string, messageId string, sessionId string) *EntityPrototype {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// Cast it as needed...
	/*if reflect.TypeOf(prototype).String() == "map[string]interface {}" {
		prototype.(map[string]interface{})["TYPENAME"] = "Server.EntityPrototype"
		values, err := Utility.InitializeStructure(prototype.(map[string]interface{}))
		if err == nil {
			prototype = values.Interface()
		} else {
			log.Println("fail to initialyse EntityPrototype from map[string]interface {} ", err)
			cargoError := NewError(Utility.FileLine(), PARAMETER_TYPE_ERROR, SERVER_ERROR_CODE, err)
			GetServer().reportErrorMessage(messageId, sessionId, cargoError)
			return nil
		}
	}

	if reflect.TypeOf(prototype).String() != "*Server.EntityPrototype" {
		cargoError := NewError(Utility.FileLine(), PARAMETER_TYPE_ERROR, SERVER_ERROR_CODE, errors.New("Expected '*Server.EntityPrototype' but got '"+reflect.TypeOf(prototype).String()+"' instead."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	// Get the store...
	store := GetServer().GetDataManager().getDataStore(storeId)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Datastore '"+storeId+"' dosen't exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	oldName := prototype.(*EntityPrototype).TypeName
	// Those types can not be rename.
	if strings.HasPrefix(oldName, "xs.") || strings.HasPrefix(oldName, "sqltypes.") || strings.HasPrefix(oldName, "XMI_types.") || strings.HasPrefix(oldName, "Config.") || strings.HasPrefix(oldName, "CargoEntities.") || strings.HasPrefix(oldName, this.m_id) {
		cargoError := NewError(Utility.FileLine(), PROTOTYPE_UPDATE_ERROR, SERVER_ERROR_CODE, errors.New("Prototype "+oldName+" cannot be rename!"))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	// So here I will get the list of all entities for that type.
	entities, _ := this.getEntities(oldName, nil, storeId, false)

	// Now I will change the prototype name
	prototype.(*EntityPrototype).TypeName = typeName

	// Save info in substitution groups...
	for i := 0; i < len(prototype.(*EntityPrototype).SubstitutionGroup); i++ {
		subTypeName := prototype.(*EntityPrototype).SubstitutionGroup[i]
		subType, err := this.getEntityPrototype(subTypeName, subTypeName[0:strings.Index(subTypeName, ".")])
		if err == nil {
			for j := 0; j < len(subType.SuperTypeNames); j++ {
				if subType.SuperTypeNames[j] == oldName {
					subType.SuperTypeNames[j] = typeName
				}
			}
			// Save it...
			subType.Save(subTypeName[0:strings.Index(subTypeName, ".")])
		}

	}

	// Save info in supertypes
	for i := 0; i < len(prototype.(*EntityPrototype).SuperTypeNames); i++ {
		superTypeName := prototype.(*EntityPrototype).SuperTypeNames[i]
		superType, err := this.getEntityPrototype(superTypeName, superTypeName[0:strings.Index(superTypeName, ".")])
		if err == nil {
			for j := 0; j < len(superType.SubstitutionGroup); j++ {
				if superType.SubstitutionGroup[j] == oldName {
					superType.SubstitutionGroup[j] = typeName
				}
			}
			superType.Save(superTypeName[0:strings.Index(superTypeName, ".")])
		}
	}

	// Now I must make tour of all prototypes in the data store and replace
	// field that made use of that prototype with it new typename.
	prototypes, err := this.getEntityPrototypes(storeId, typeName[0:strings.Index(typeName, ".")])
	if err == nil {
		for i := 0; i < len(prototypes); i++ {
			p := prototypes[i]
			for j := 0; j < len(p.FieldsType); j++ {
				if strings.Index(p.FieldsType[j], oldName) > 0 {
					strings.Replace(p.FieldsType[j], oldName, typeName, -1)
				}
			}

			// save the prototype.
			p.Save(storeId)

		}
	} else {
		cargoError := NewError(Utility.FileLine(), PROTOTYPE_UPDATE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	prototype.(*EntityPrototype).TypeName = typeName
	store.(*GraphStore).DeleteEntityPrototype(oldName)
	store.(*GraphStore).saveEntityPrototype(prototype.(*EntityPrototype))

	// Now I must update entities...
	for i := 0; i < len(entities); i++ {
		if reflect.TypeOf(entities[i]).String() == "*Server.DynamicEntity" {
			//
			entity := entities[i].(*DynamicEntity)
			ids := make([]interface{}, 0)
			p, _ := this.getEntityPrototype(entity.GetTypeName(), entity.GetTypeName()[0:strings.Index(entity.GetTypeName(), ".")])
			for j := 0; j < len(p.Ids); j++ {
				ids = append(ids, entity.getValue(p.Ids[j]))
			}

			// Here I will delete the existing entity from the db...
			entity.setValue("UUID", nil)          // Set it uuid to nil
			entity.setValue("TYPENAME", typeName) // Set it new typeName
			// Recreate it with it new type
			newEntity, errObj := this.newDynamicEntity(entity.GetParentUuid(), entity.getValues().(map[string]interface{}))
			if errObj != nil {
				newEntity.SaveEntity() // Save the new entity
			}
		}

	}

	return prototype.(*EntityPrototype)*/
	return nil

}

// @api 1.0
// That function will retreive all prototypes of a store.
// @param {string} storeId The store id, where to create the new prototype.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{[]*EntityPrototype} Return the retreived list of entity prototype
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
// EntityManager.prototype.getEntityPrototypes = function (storeId, successCallback, errorCallback, caller) {
//    // server is the client side singleton.
//    var params = [];
//    params.push(createRpcData(storeId, "STRING", "storeId"));
//    // Call it on the server.
//    server.executeJsFunction(
//        "EntityManagerGetEntityPrototypes", // The function to execute remotely on server
//        params, // The parameters to pass to that function
//        function (index, total, caller) { // The progress callback
//            // Nothing special to do here.
//        },
//        function (results, caller) {
//            var results = results[0]
//            var prototypes = []
//            if (results != null) {
//                for (var i = 0; i < results.length; i++) {
//                    var proto = new EntityPrototype();
//                    entityPrototypes[results[i].TypeName] = proto;
//                    proto.init(results[i]);
//                    prototypes.push(proto);
//                }
//            }
// 			 if(caller.successCallback!=undefined){
//            	caller.successCallback(prototypes, caller.caller);
//            	caller.successCallback = undefined;
//			 }
//        },
//        function (errMsg, caller) {
//          server.errorManager.onError(errMsg)
//         	if( caller.errorCallback != undefined){
//          		caller.errorCallback(errMsg, caller.caller);
//				caller.errorCallback = undefined;
//			}
//        }, // Error callback
//        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
//    )
//}
func (this *EntityManager) GetEntityPrototypes(storeId string, messageId string, sessionId string) []*EntityPrototype {
	store := GetServer().GetDataManager().getDataStore(storeId)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("There is no store with id '"+storeId+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	protos, err := store.GetEntityPrototypes()
	if err != nil {
		cargoError := NewError(Utility.FileLine(), PROTOTYPE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("There is no prototypes in store '"+storeId+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	// Need to sort the prototype because of namespace evaluation.
	sort.Slice(protos[:], func(i, j int) bool {
		return protos[i].TypeName < protos[j].TypeName
	})

	return protos
}

// @api 1.0
// That function will retreive the entity prototype with a given type name.
// @param {string} typeName The type name of the prototype to retreive.
// @param {string} storeId The store id, where to create the new prototype.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{*EntityPrototype} Return the retreived entity prototype
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
// EntityManager.prototype.getEntityPrototype = function (typeName, storeId, successCallback, errorCallback, caller) {
//    // Retrun entity prototype that aleady exist.
//    if (entityPrototypes[typeName] != undefined) {
//        successCallback(entityPrototypes[typeName], caller);
//        successCallback = undefined;
//        return;
//    }
//    // server is the client side singleton.
//    var params = [];
//    params.push(createRpcData(typeName, "STRING", "typeName"));
//    params.push(createRpcData(storeId, "STRING", "storeId"));
//    // Call it on the server.
//    server.executeJsFunction(
//        "EntityManagerGetEntityPrototype", // The function to execute remotely on server
//        params, // The parameters to pass to that function
//        function (index, total, caller) { // The progress callback
//            // Nothing special to do here.
//        },
//        function (results, caller) {
//            var proto = new EntityPrototype();
//            entityPrototypes[results[0].TypeName] = proto;
//            proto.init(results[0]);
//			 if(caller.successCallback!=undefined){
//            	caller.successCallback(proto, caller.caller);
//            	caller.successCallback = undefined;
//        	}
//        },
//        function (errMsg, caller) {
//          server.errorManager.onError(errMsg);
//         	if( caller.errorCallback != undefined){
//          		caller.errorCallback(errMsg, caller.caller);
//				caller.errorCallback = undefined;
//			}
//        }, // Error callback
//        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
//    )
//}
func (this *EntityManager) GetEntityPrototype(typeName string, storeId string, messageId string, sessionId string) *EntityPrototype {

	proto, err := this.getEntityPrototype(typeName, storeId)
	if err != nil {
		cargoError := NewError(Utility.FileLine(), ENTITY_PROTOTYPE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return nil
	}

	return proto
}

//////////////////////////////// Entities //////////////////////////////////////

// @api 1.0
// Event handler function.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
// EntityManager.prototype.onEvent = function (evt) {
//     // Set the internal object.
//     if (evt.code == UpdateEntityEvent || evt.code == NewEntityEvent) {
//         if (entityPrototypes[evt.dataMap["entity"].TYPENAME] == undefined) {
//             return;
//         }
//         if (entities[evt.dataMap["entity"].UUID] == undefined) {
//             var entity = eval("new " + evt.dataMap["entity"].TYPENAME + "()");
//             var initCallback = function (self, evt, entity) {
//                 return function (entity) {
//                     server.entityManager.setEntity(entity);
//                     EventHub.prototype.onEvent.call(self, evt);
//                 }
//             }(this, evt, entity)
//             if (entity.initCallbacks == undefined) {
//                 entity.initCallbacks = [];
//             }
//             entity.initCallbacks.push(initCallback);
//             entity.init(evt.dataMap["entity"], false);
//         } else {
//             // update the object values.
//             // but before I call the event I will be sure the entity have
//             var entity = entities[evt.dataMap["entity"].UUID];
//             var initCallback = function (self, evt, entity) {
//                 return function (entity) {
//                     // Test if the object has change here befor calling it.
//                     server.entityManager.setEntity(entity);
//                     if (evt.done == undefined) {
//                         EventHub.prototype.onEvent.call(self, evt);
//                     }
//                     evt.done = true; // Cut the cyclic recursion.
//                 }
//             }(this, evt, entity)
//             if (entity.initCallbacks == undefined) {
//                 entity.initCallbacks = [];
//             }
//             entity.initCallbacks.push(initCallback);
//             setObjectValues(entity, evt.dataMap["entity"]);
//         }
//     } else if (evt.code == DeleteEntityEvent) {
//         var entity = entities[evt.dataMap["entity"].UUID];
//         if (entity != undefined) {
//             this.resetEntity(entity);
//             EventHub.prototype.onEvent.call(this, evt);
//         }
//     }
// }
func (this *EntityManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

// @api 1.0
// Set the value of an entity on the entityManager.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
// EntityManager.prototype.setEntity = function (entity) {
//    if(entity.UUID.length == 0){
//		 console.log("entity has no UUID!", entity);
//		 return;
//	  }
//	  entities[entity.UUID] = entity
//    this.getEntityPrototype(entity.TYPENAME, entity.TYPENAME.split(".")[0],
//        function (prototype, caller) {
//            var id_ = entity.TYPENAME + ":";
//            for (var i = 0; i < prototype.Ids.length; i++) {
//                var id = prototype.Ids[i];
//                if (id == "UUID" || id == "uuid") {
//					  if(entity.UUID != undefined){
//					  	if(entities[entity.UUID] != undefined && entity.UUID.length > 0){
//							entity.ParentLnk = entities[entity.UUID].ParentLnk;
//					  	}
//					  	entities[entity.UUID] = entity;
//					  }
//                } else if(entity[id] != undefined) {
//                    if (entity[id].length > 0) {
//                        id_ += entity[id];
//                        if (i < prototype.Ids.length - 1) {
//                            id_ += "_";
//                        }
//						  if(i == prototype.Ids.length - 1){
//					  		if(entities[id_] != undefined){
//								entity.ParentLnk = entities[id_].ParentLnk;
//					  		}
//							entities[id_] = entity;
//						  }
//                    }
//                }
//            }
//        },
//        function (errMsg, caller) {
//            /** Nothing to do here. */
//        },
//        {})
//}
func (this *EntityManager) SetEntity(values interface{}) {
	/** empty function here... **/
}

// @api 1.0
// Reset the value of an entity on the entityManager.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
// EntityManager.prototype.resetEntity = function (entity) {
//    var prototype = entityPrototypes[entity.TYPENAME]
//    delete entities[entity.UUID]
//    var id = entity.TYPENAME + ":";
//    for (var i = 0; i < prototype.Ids.length; i++) {
//        id += entity[prototype.Ids[i]];
//        if (i < prototype.Ids.length - 1) {
//            id += "_";
//        }
//    }
//    if (entities[id] != undefined) {
//        delete entities[id];
//    }
//}
func (this *EntityManager) ResetEntity(values interface{}) {
	/** empty function here... **/
}

// @api 1.0
// That function is use to create a new entity of a given type..
// @param {string} parentUuid The uuid of the parent entity if there is one, null otherwise.
// @param {string} attributeName The attribute name is the name of the new entity in his parent. (parent.attributeName = this)
// @param {interface{}} values the entity to be save, it can be nil.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{interface{}} Return the created entity
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
// EntityManager.prototype.createEntity = function (parentUuid, attributeName, entity, successCallback, errorCallback, caller) {
//     // server is the client side singleton.
//     var params = [];
//     params.push(createRpcData(parentUuid, "STRING", "parentUuid"));
//     params.push(createRpcData(attributeName, "STRING", "attributeName"));
//     params.push(createRpcData(entity, "JSON_STR", "entity"));
//     // Call it on the server.
//     server.executeJsFunction(
//         "EntityManagerCreateEntity", // The function to execute remotely on server
//         params, // The parameters to pass to that function
//         function (index, total, caller) { // The progress callback
//             // Nothing special to do here.
//         },
//         function (result, caller) {
//             var entity = eval("new " + result[0].TYPENAME + "()");
//             var initCallback = function () {
//                 return function (entity) {
//                     if (caller.successCallback != undefined) {
//                         caller.successCallback(entity, caller.caller);
//                         caller.successCallback = undefined;
//                     }
//                 }
//             }(caller)
//             if (entity.initCallbacks == undefined) {
//                 entity.initCallbacks = [];
//             }
//             entity.initCallbacks.push(initCallback);
//             entity.init(result[0], false);
//         },
//         function (errMsg, caller) {
//             server.errorManager.onError(errMsg);
//             if (caller.errorCallback != undefined) {
//                 caller.errorCallback(errMsg, caller.caller);
//                 caller.errorCallback = undefined;
//             }
//         }, // Error callback
//         { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
//     )
// }
func (this *EntityManager) CreateEntity(parentUuid string, attributeName string, values interface{}, messageId string, sessionId string) interface{} {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	var parent Entity
	parent, errObj = this.getEntityByUuid(parentUuid)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	if reflect.TypeOf(values).String() == "map[string]interface {}" {
		obj, err := Utility.InitializeStructure(values.(map[string]interface{}), setEntityFct)
		if err == nil {
			// Here I will take assumption I got an entity...
			// Now I will save the entity.
			if reflect.TypeOf(obj.Interface()).String() == "Server.Entity" {
				_, errObj = this.createEntity(parent, attributeName, obj.Interface().(Entity))
				if errObj != nil {
					GetServer().reportErrorMessage(messageId, sessionId, errObj)
					return nil
				}
				return obj.Interface().(Entity)
			} else {
				entity := NewDynamicEntity()
				entity.setObject(values.(map[string]interface{}))
				_, errObj = this.createEntity(parent, attributeName, entity)
				if errObj != nil {
					GetServer().reportErrorMessage(messageId, sessionId, errObj)
					return nil
				}
				return entity.getValues()
			}
		} else {
			errObj = NewError(Utility.FileLine(), ENTITY_CREATION_ERROR, SERVER_ERROR_CODE, err)
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
	} else {
		result, errObj := this.createEntity(parent, attributeName, values.(Entity))
		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
		return result
	}
}

// @api 1.0
// Save The entity. If the entity does not exist it creates it.
// @param {interface{}} values The entity to save.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{interface{}} Return an object (Entity)
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
// EntityManager.prototype.saveEntity = function (entity, successCallback, errorCallback, caller) {
//     // server is the client side singleton.
//     var params = [];
//     params.push(createRpcData(entity, "JSON_STR", "entity"));
//     // Call it on the server.
//     server.executeJsFunction(
//         "EntityManagerSaveEntity", // The function to execute remotely on server
//         params, // The parameters to pass to that function
//         function (index, total, caller) { // The progress callback
//             // Nothing special to do here.
//         },
//         function (result, caller) {
//             var entity = eval("new " + result[0].TYPENAME + "()");
//             var initCallback = function () {
//                 return function (entity) {
//                     // Set the new entity values...
//                     server.entityManager.setEntity(entity);
//                     if (caller.successCallback != undefined) {
//                         caller.successCallback(entity, caller.caller);
//                         caller.successCallback = undefined;
//                     }
//                 }
//             }(caller)
//             if (entity.initCallbacks == undefined) {
//                 entity.initCallbacks = [];
//             }
//             entity.initCallbacks.push(initCallback);
//             entity.init(result[0], false);
//         },
//         function (errMsg, caller) {
//             server.errorManager.onError(errMsg);
//             if (caller.errorCallback != undefined) {
//                 caller.errorCallback(errMsg, caller.caller);
//                 caller.errorCallback = undefined;
//             }
//         }, // Error callback
//         { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
//     )
// }
func (this *EntityManager) SaveEntity(values interface{}, messageId string, sessionId string) interface{} {
	var errObj *CargoEntities.Error
	errObj = GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	if reflect.TypeOf(values).String() == "map[string]interface {}" {
		obj, err := Utility.InitializeStructure(values.(map[string]interface{}), setEntityFct)
		if err == nil {
			// Here I will take assumption I got an entity...
			// Now I will save the entity.
			if reflect.TypeOf(obj.Interface()).String() == "Server.Entity" {
				errObj = this.saveEntity(obj.Interface().(Entity))
				if errObj != nil {
					GetServer().reportErrorMessage(messageId, sessionId, errObj)
					return nil
				}
				return obj.Interface().(Entity)
			} else {
				entity := NewDynamicEntity()
				entity.setObject(values.(map[string]interface{}))
				errObj = this.saveEntity(entity)
				if errObj != nil {
					GetServer().reportErrorMessage(messageId, sessionId, errObj)
					return nil
				}
				return entity.getValues()
			}
		} else {
			errObj = NewError(Utility.FileLine(), ENTITY_CREATION_ERROR, SERVER_ERROR_CODE, err)
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
	} else {
		// Set it propertie to need save.
		values.(Entity).SetNeedSave(true)
		errObj = this.saveEntity(values.(Entity))
		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
		values.(Entity).SetNeedSave(false)

		return values.(Entity)
	}

}

// @api 1.0
// That function is use to remove an entity with a given uuid.
// @param {string} uuid The uuid of entity to delete. Must have the form TypeName%UUID
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *EntityManager) RemoveEntity(uuid string, messageId string, sessionId string) {
	var errObj *CargoEntities.Error
	errObj = GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	// The entity to remove.
	var entity Entity
	entity, errObj = this.getEntityByUuid(uuid)

	if entity != nil {
		// validate over the entity
		/*errObj = GetServer().GetSecurityManager().hasPermission(sessionId, CargoEntities., entity)
		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return // exit here.
		}*/

		// Suppress the enitity...
		errObj := this.deleteEntity(entity)
		if errObj == nil {
			return
		}

	}

	// Repport the error
	GetServer().reportErrorMessage(messageId, sessionId, errObj)

}

// @api 1.0
// That function is use to retreive objects with a given type.
// @param {string} typeName The name of the type we looking for in the form packageName.typeName
// @param {string} storeId The name of the store where the information is saved.
// @param {EntityQuery} query It contain the code of a function to be executed by the server to filter specific values.
// @param {int} offset	Results offset
// @param {int} limit	The number of results to return. Can be use to create page of results.
// @param {[]string} orderBy the list of field that specifie the result order.
// @param {bool} asc the list of field that specifie the result order.
// @param {bool} lazy Load all child's value if false or just child's uuid if true.
// @result{[]interface{}} Return an array of object's (Entities)
// @scope {public}
// @param {callback} progressCallback The function is call when chunk of response is received.
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
//EntityManager.prototype.getEntities = function (typeName, storeId, query, offset, limit, orderBy, asc, lazy, progressCallback, successCallback, errorCallback, caller) {
//    // First of all i will get the entity prototype.
//    server.entityManager.getEntityPrototype(typeName, storeId,
//        // The success callback.
//        function (result, caller) {
//            // Set the parameters.
//            var typeName = caller.typeName;
//            var storeId = caller.storeId;
//            var query = caller.query;
//            var successCallback = caller.successCallback;
//            var progressCallback = caller.progressCallback;
//            var errorCallback = caller.errorCallback;
//            var lazy = caller.lazy;
//            var caller = caller.caller;
//            // Create the list of parameters.
//            var params = []
//            params.push(createRpcData(typeName, "STRING", "typeName"));
//            params.push(createRpcData(storeId, "STRING", "storeId"));
//            params.push(createRpcData(query, "JSON_STR", "query"));
//            params.push(createRpcData(offset, "INTEGER", "offset"));
//            params.push(createRpcData(limit, "INTEGER", "limit"));
//            params.push(createRpcData(orderBy, "JSON_STR", "orderBy", "[]string"));
//            params.push(createRpcData(asc, "BOOLEAN", "asc"));
//            params.push(createRpcData(lazy, "BOOLEAN", "lazy"));
//            // Call it on the server.
//            server.executeJsFunction(
//                "EntityManagerGetEntities", // The function to execute remotely on server
//                params, // The parameters to pass to that function
//                function (index, total, caller) { // The progress callback
//                    // Keep track of the file transfert.
//                    caller.progressCallback(index, total, caller.caller);
//                },
//                function (filePath, caller) { // Success callback
//                    var xhr = new XMLHttpRequest();
//                    xhr.open('GET', filePath, true);
//                    xhr.responseType = 'text';
//                    xhr.onload = function (caller) {
//                        return function (e) {
//                            if (this.status == 200) {
//                                // Note: .response instead of .responseText
//                                var jsonStr = this.responseText
//                                var entities = []
//                                if (jsonStr == "null") {
//                                    if (caller.successCallback != undefined) {
//                                        caller.successCallback(entities, caller.caller);
//                                        caller.successCallback = undefined;
//                                    }
//                                } else {
//                                    // Now I will initialyse entities from json objects.
//                                    var values = JSON.parse(jsonStr)
//                                    if (values.length > 0) {
//                                        var initEntitiesFct = function (values, caller, entities) {
//                                            var value = values.pop();
//                                            var entity = eval("new " + caller.prototype.TypeName + "()")
//                                            entities.push(entity);
//                                            if (values.length == 0) {
//                                                var initCallback = function (caller, entities) {
//                                                    return function (entity) {
//                                                        server.entityManager.setEntity(entity);
//                                                        if (caller.successCallback != undefined) {
//                                                            caller.successCallback(entities, caller.caller);
//                                                            caller.successCallback = undefined;
//                                                        }
//                                                    }
//                                                }(caller, entities)
//                                                if (entity.initCallbacks == undefined) {
//                                                    entity.initCallbacks = [];
//                                                }
//                                                entity.initCallbacks.push(initCallback);
//                                                entity.init(value, lazy);
//                                            } else {
//                                                var initCallback = function (values, caller, entities, initEntitiesFct) {
//                                                    return function (entity) {
//                                                        server.entityManager.setEntity(entity);
//                                                        initEntitiesFct(values, caller, entities);
//                                                    }
//                                                }(values, caller, entities, initEntitiesFct)
//                                                if (entity.initCallbacks == undefined) {
//                                                    entity.initCallbacks = [];
//                                                }
//                                                entity.initCallbacks.push(initCallback);
//                                                entity.init(value, lazy);
//                                            }
//                                        }
//                                        initEntitiesFct(values, caller, entities);
//                                    } else {
//                                        if (caller.successCallback != undefined) {
//                                            caller.successCallback(entities, caller.caller);
//                                            caller.successCallback = undefined;
//                                        }
//                                    }
//                                }
//                            }
//                        }
//                    }(caller)
//                    xhr.onprogress = function (progressCallback, caller) {
//                        return function (e) {
//							if (!e.lengthComputable) {
//					  			e.total = parseInt(e.target.getResponseHeader('x-decompressed-content-length'), 10);
//							}
//                          progressCallback(e.loaded, e.total, caller)
//                        }
//                    }(caller.progressCallback, caller.caller)
//                    xhr.send();
//                },
//                function (errMsg, caller) {
//                    // call the immediate error callback.
//                    if (caller.errorCallback != undefined) {
//                        caller.errorCallback(errMsg, caller.caller);
//                        caller.errorCallback = undefined;
//                    }
//                    // dispatch the message.
//                    server.errorManager.onError(errMsg);
//                }, // Error callback
//                { "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback, "prototype": result, "lazy": lazy } // The caller
//            )
//        },
//        // The error callback.
//        function (errMsg, caller) {
//            // call the immediate error callback.
//            if (caller.errorCallback != undefined) {
//                caller.errorCallback(errMsg, caller.caller);
//                caller.errorCallback = undefined;
//            }
//            // dispatch the message.
//            server.errorManager.onError(errMsg);
//        }, { "typeName": typeName, "storeId": storeId, "query": query, "caller": caller, "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback, "lazy": lazy })
//}
func (this *EntityManager) GetEntities(typeName string, storeId string, q interface{}, offset float64, limit float64, orderBy []interface{}, asc bool, lazy bool, messageId string, sessionId string) interface{} {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// The entity query.
	var query *EntityQuery
	if q != nil {
		query = q.(*EntityQuery)
	} else {
		query = new(EntityQuery)
		query.TYPENAME = "Server.EntityQuery"
		query.TypeName = typeName
	}

	query.Fields = make([]string, 0)

	// I will execute the query and get it results.
	store := GetServer().GetDataManager().getDataStore(storeId)
	queryStr, _ := json.Marshal(query)
	values, err := store.Read(string(queryStr), []interface{}{}, []interface{}{})
	results := make([]interface{}, 0)
	if err == nil {
		// Sort the entities
		/*this.sortEntities(entities, orderBy, 0, len(entities), asc)
		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return objects
		}*/
		if limit <= 0 {
			// all results are require.
			limit = float64(len(values))
		}

		for i := int(offset); i < int(limit); i++ {
			if !lazy {
				results = append(results, initChilds(values[i][0].(map[string]interface{})))
			} else {
				results = append(results, values[i][0].(map[string]interface{}))
			}
		}
	} else {
		errObj = NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// Here I will create a file in the tmp directory and send back it path to the
	// client.
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(toJsonStr(results))); err != nil {
		panic(err)
	}

	if err := gz.Flush(); err != nil {
		panic(err)
	}

	if err := gz.Close(); err != nil {
		panic(err)
	}

	// So here I will create the response as a compress file.
	var path = GetServer().GetConfigurationManager().GetTmpPath() + "/" + messageId + ".gz"
	err = ioutil.WriteFile(path, b.Bytes(), 0666)

	if err != nil {
		errObj := NewError(Utility.FileLine(), FILE_WRITE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	return "/tmp/" + messageId + ".gz"
}

// @api 1.0
// That function is use to retreive objects with a given type.
// @param {string} uuid The uuid of the entity we looking for. The uuid must has form typeName%UUID.
// @param {bool} lazy Load all child's value if false or just child's uuid if true.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{interface{}} Return an object (Entity)
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
//EntityManager.prototype.getEntityByUuid = function (uuid, lazy, successCallback, errorCallback, caller) {
//    if (uuid.length == 0) {
//        return;
//    }
//    var entity = entities[uuid]
//    if (entity != undefined) {
//        if (entity.TYPENAME == entity.class_name_ && entity.IsInit == true) {
//            successCallback(entity, caller);
//            return // break it here.
//        }
//    }
//    var typeName = uuid.substring(0, uuid.indexOf("%"));
//    var storeId = typeName.substring(0, typeName.indexOf("."));
//    // Create the entity prototype here.
//    var entity = eval("new " + typeName + "()");
//    entity.UUID = uuid;
//    entity.TYPENAME = typeName;
//    server.entityManager.setEntity(entity);
//    // First of all i will get the entity prototype.
//    server.entityManager.getEntityPrototype(typeName, storeId,
//        // The success callback.
//        function (result, caller) {
//            // Set the parameters.
//            var uuid = caller.uuid;
//            var successCallback = caller.successCallback;
//            var progressCallback = caller.progressCallback;
//            var errorCallback = caller.errorCallback;
//            var lazy = caller.lazy;
//            var caller = caller.caller;
//            var params = [];
//            params.push(createRpcData(uuid, "STRING", "uuid"));
//            params.push(createRpcData(lazy, "BOOLEAN", "lazy"));
//            // Call it on the server.
//            server.executeJsFunction(
//                "EntityManagerGetEntityByUuid", // The function to execute remotely on server
//                params, // The parameters to pass to that function
//                function (index, total, caller) { // The progress callback
//                    // Nothing special to do here.
//                },
//                function (results, caller) {
//                    // Initialyse the entity from the value.
//                    function initEntitiy(values, caller) {
//                        var entity = entities[values.UUID];
//                        if (entity == null) {
//                            entity = eval("new " + caller.prototype.TypeName + "()");
//                            entity.UUID = values.UUID;
//                            entity.TYPENAME = caller.prototype.TypeName;
//                            server.entityManager.setEntity(entity);
//                        }
//                        var initCallback = function (caller) {
//                            return function (entity) {
//                                server.entityManager.setEntity(entity);
//                                if (caller.successCallback != undefined) {
//                                    caller.successCallback(entity, caller.caller);
//                                    caller.successCallback = undefined;
//                                }
//                            }
//                        }(caller)
//                        if (entity.initCallbacks == undefined) {
//                            entity.initCallbacks = [];
//                        }
//                        entity.initCallbacks.push(initCallback);
//                        if (entity.IsInit == false) {
//                            entity.init(values, lazy);
//                        } else {
//                            if (caller.successCallback != undefined) {
//                                caller.successCallback(entity, caller.caller);
//                                caller.successCallback = undefined;
//                            }
//                        }
//                    }
//                    if (caller.lazy) {
//                        // In that case the result contain the entity.
//                        initEntitiy(results[0], caller)
//                    } else {
//                        // In that case the result contain the path to .gz file that contain the result as
//                        // json string.
//                        var xhr = new XMLHttpRequest();
//                        xhr.open('GET', results, true);
//                        xhr.responseType = 'text';
//                        xhr.onload = function (caller) {
//                            return function (e) {
//                                if (this.status == 200) {
//                                    var jsonStr = this.responseText;
//                                    var values = JSON.parse(jsonStr)
//									  try{
//										initEntitiy(values, caller)
//									  }catch(err){
//			               				caller.errorCallback(err, caller.caller);
//                        				caller.errorCallback = undefined;
//									  }
//                                }
//                            }
//                        }(caller)
//                        xhr.send();
//                    }
//                },
//                function (errMsg, caller) {
//                    server.errorManager.onError(errMsg);
//                    if (caller.errorCallback != undefined) {
//                        caller.errorCallback(errMsg, caller.caller);
//                        caller.errorCallback = undefined;
//                    }
//                }, // Error callback
//                { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback, "prototype": result, "lazy": lazy } // The caller
//            )
//        },
//        // The error callback.
//        function (errMsg, caller) {
//            server.errorManager.onError(errMsg);
//            if (caller.errorCallback != undefined) {
//                caller.errorCallback(errMsg, caller.caller);
//                caller.errorCallback = undefined;
//            }
//        }, { "uuid": uuid, "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback, "lazy": lazy })
//}
func (this *EntityManager) GetEntityByUuid(uuid string, lazy bool, messageId string, sessionId string) interface{} {

	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// In that case the code is server side synchronize code.
	if len(messageId) == 0 {
		entity, errObj := this.getEntityByUuid(uuid)
		if errObj != nil {
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
		// Here the entity with only it childs uuid's are return.
		if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
			return entity.(*DynamicEntity).getValues()
		}

		return entity
	} else {
		entity, err := getEntityByUuid(uuid)
		if err != nil {
			errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, err)
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
		if !lazy {
			entity = initChilds(entity)
		}

		// Here I will create a file in the tmp directory and send back it path to the
		// client.
		var values []byte
		values, err = json.Marshal(entity)
		if err != nil {
			errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, err)
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		if _, err := gz.Write(values); err != nil {
			panic(err)
		}

		if err := gz.Flush(); err != nil {
			panic(err)
		}

		if err := gz.Close(); err != nil {
			panic(err)
		}
		// So here I will create the response as a compress file.
		var path = GetServer().GetConfigurationManager().GetTmpPath() + "/" + messageId + ".gz"
		err = ioutil.WriteFile(path, b.Bytes(), 0666)
		if err != nil {
			errObj := NewError(Utility.FileLine(), FILE_WRITE_ERROR, SERVER_ERROR_CODE, err)
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}

		return "/tmp/" + messageId + ".gz"
	}

}

// Recursively replace child reference by theire actual values.
func initChilds(values map[string]interface{}) map[string]interface{} {
	typeName := values["TYPENAME"].(string)
	storeId := typeName[0:strings.Index(typeName, ".")]
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, storeId)
	for i := 0; i < len(prototype.Fields); i++ {
		if strings.HasPrefix(prototype.Fields[i], "M_") {
			isBaseType := strings.HasPrefix(prototype.FieldsType[i], "xs.") || strings.HasPrefix(prototype.FieldsType[i], "[]xs.") || strings.HasPrefix(prototype.FieldsType[i], "enum:")
			if !isBaseType {
				isRef := strings.HasSuffix(prototype.FieldsType[i], ":Ref") || strings.HasSuffix(prototype.Fields[i], "Ptr")
				if !isRef {
					v := reflect.ValueOf(values[prototype.Fields[i]])
					// In case of an array of references.
					if v.IsValid() {
						if v.Type().Kind() == reflect.Slice {
							for j := 0; j < v.Len(); j++ {
								if v.Index(j).Type().Kind() == reflect.Interface {
									if reflect.TypeOf(v.Index(j).Interface()).Kind() == reflect.String {
										if Utility.IsValidEntityReferenceName(v.Index(j).Interface().(string)) {
											v_, err := getEntityByUuid(v.Index(j).Interface().(string))
											if err == nil {
												values[prototype.Fields[i]].([]interface{})[j] = v_

											}
										}
									}
								}
							}
						} else {
							// In case of a reference.
							if v.Type().Kind() == reflect.String {
								// Here I will test if the value is a valid reference.
								if Utility.IsValidEntityReferenceName(v.String()) {
									v_, err := getEntityByUuid(v.String())
									if err == nil {
										values[prototype.Fields[i]] = v_
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return values
}

// Simple function that return an entity from it uuid if the entity exist.
func getEntityByUuid(uuid string) (map[string]interface{}, error) {

	typeName := strings.Split(uuid, "%")[0]
	storeId := typeName[0:strings.Index(typeName, ".")]

	var query EntityQuery
	query.TYPENAME = typeName
	query.TypeName = typeName
	query.Fields = []string{}
	query.Query = typeName + `.UUID=="` + uuid + `"`

	store := GetServer().GetDataManager().getDataStore(storeId)
	queryStr, _ := json.Marshal(query)
	values, err := store.Read(string(queryStr), []interface{}{}, []interface{}{})

	if err == nil {
		// init it child values.
		initChilds(values[0][0].(map[string]interface{}))
		// return the resulting map.
		return values[0][0].(map[string]interface{}), nil
	}

	return nil, err
}

// @api 1.0
// Return Retreive JSON object form the list of uuid's. Use this function if the number of result is high.
// @param {[]string} uuids The list of uuid's of entity to get.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} progressCallback display the progression of the data transfert.
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
//EntityManager.prototype.getEntitiesByUuid = function (uuids, progressCallback, successCallback, errorCallback, caller) {
//   var params = [];
//   params.push(createRpcData(uuids, "JSON_STR", "uuids", "[]string"));
//   server.executeJsFunction(
//       "EntityManagerGetEntitiesByUuid",
//       params,
//       function (index, total, caller) { // The progress callback
//           // Keep track of the file transfert.
//           caller.progressCallback(index, total, caller.caller);
//       }, //progress callback
//       function (filePath, caller) { // Success callback
//           var xhr = new XMLHttpRequest();
//           xhr.open('GET', filePath, true);
//           xhr.responseType = 'text';
//           xhr.onload = function (caller) {
//               return function (e) {
//                   if (this.status == 200) {
//                       var jsonStr = this.responseText;
//						 try{
//                       	caller.successCallback(JSON.parse(jsonStr), caller.caller)
//						 }catch(err){
//               			caller.errorCallback(err, caller.caller);
//						 }
//                   }
//               }
//           }(caller)
//           xhr.onprogress = function (progressCallback, caller) {
//               return function (e) {
//					if (!e.lengthComputable) {
//					  e.total = parseInt(e.target.getResponseHeader('x-decompressed-content-length'), 10);
//					}
//                  progressCallback(e.loaded, e.total, caller)
//               }
//           }(caller.progressCallback, caller.caller)
//           xhr.send();
//       },
//       function (errMsg, caller) { // Error callback
//           if (caller.errorCallback != undefined) {
//               caller.errorCallback(errMsg, caller.caller);
//               caller.errorCallback = undefined;
//           }
//           server.errorManager.onError(errMsg);
//       }, { "successCallback": successCallback, "progressCallback": progressCallback, "errorCallback": errorCallback, "caller": caller })
//}
func (this *EntityManager) GetEntitiesByUuid(uuids []interface{}, messageId string, sessionId string) interface{} {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// The results.
	results := make([]interface{}, 0)

	// I will use the data manager directly for performance purpose.
	for i := 0; i < len(uuids); i++ {
		uuid := uuids[i].(string)
		result, err := getEntityByUuid(uuid)
		if err == nil {
			results = append(results, result)
		} else {
			errObj := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, err)
			GetServer().reportErrorMessage(messageId, sessionId, errObj)
			return nil
		}
	}

	// Here I will create a file in the tmp directory and send back it path to the
	// client.
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(toJsonStr(results))); err != nil {
		panic(err)
	}

	if err := gz.Flush(); err != nil {
		panic(err)
	}

	if err := gz.Close(); err != nil {
		panic(err)
	}

	// So here I will create the response as a compress file.
	var path = GetServer().GetConfigurationManager().GetTmpPath() + "/" + messageId + ".gz"
	err := ioutil.WriteFile(path, b.Bytes(), 0666)
	if err != nil {
		errObj := NewError(Utility.FileLine(), FILE_WRITE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	return "/tmp/" + messageId + ".gz"
}

// @api 1.0
// Retrieve an entity with a given typename and id.
// @param {string} typeName The object type name.
// @param {string} storeId The object type name.
// @param {string} ids The id's (not uuid) of the object to look for.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{interface{}} Return an object (Entity)
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
// @src
// EntityManager.prototype.getEntityById = function (typeName, storeId, ids, lazy, successCallback, errorCallback, caller, parent) {
//     // key in the server.
//     var id = typeName + ":";
//     for (var i = 0; i < ids.length; i++) {
//         id += ids[i];
//         if (i < ids.length - 1) {
//             id += "_";
//         }
//     }
//     if (entities[id] != undefined) {
//         successCallback(entities[id], caller);
//         return;// break it here.
//     }
//     // First of all i will get the entity prototype.
//     server.entityManager.getEntityPrototype(typeName, storeId,
//         // The success callback.
//         function (result, caller) {
//             // Set the parameters.
//             var storeId = caller.storeId;
//             var typeName = caller.typeName;
//             var ids = caller.ids;
//             var successCallback = caller.successCallback;
//             var progressCallback = caller.progressCallback;
//             var errorCallback = caller.errorCallback;
//             var lazy = caller.lazy;
//             var caller = caller.caller;
//             var params = [];
//             params.push(createRpcData(typeName, "STRING", "typeName"));
//             params.push(createRpcData(storeId, "STRING", "storeId"));
//             params.push(createRpcData(ids, "JSON_STR", "ids")); // serialyse as an JSON object array...
// 			   params.push(createRpcData(lazy, "BOOLEAN", "lazy"));
//             // Call it on the server.
//             server.executeJsFunction(
//                 "EntityManagerGetEntityById", // The function to execute remotely on server
//                 params, // The parameters to pass to that function
//                 function (index, total, caller) { // The progress callback
//                     // Nothing special to do here.
//                 },
//                 function (result, caller) {
//                     if (result[0] == null) {
//                         return;
//                     }
//                     // In case of existing entity.
//                     if (entities[result[0].UUID] != undefined && result[0].TYPENAME == result[0].class_name_) {
//                         if (caller.successCallback != undefined) {
//                             caller.successCallback(entities[result[0].UUID], caller.caller);
//                             caller.successCallback = undefined;
//                         }
//                         return; // break it here.
//                     }
//                     var entity = eval("new " + caller.prototype.TypeName + "(caller.prototype)")
//                     var initCallback = function () {
//                         return function (entity) {
//                             if (caller.successCallback != undefined) {
//                                 caller.successCallback(entity, caller.caller);
//                                 caller.successCallback = undefined;
//                             }
//                         }
//                     }(caller)
//                     if (entity.initCallbacks == undefined) {
//                         entity.initCallbacks = [];
//                     }
//                     entity.initCallbacks.push(initCallback)
//                     entity.init(result[0], lazy);
//                 },
//                 function (errMsg, caller) {
//                     server.errorManager.onError(errMsg);
//                     if (caller.errorCallback != undefined) {
//                         caller.errorCallback(errMsg, caller.caller);
//                         caller.errorCallback = undefined;
//                     }
//                 }, // Error callback
//                 { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback, "prototype": result, "parent": parent, "ids": ids, "lazy": lazy } // The caller
//             )
//         },
//         // The error callback.
//         function (errMsg, caller) {
//             server.errorManager.onError(errMsg);
//             caller.errorCallback(errMsg, caller);
//         }, { "storeId": storeId, "typeName": typeName, "ids": ids, "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback, "lazy": lazy })
// }
func (this *EntityManager) GetEntityById(typeName string, storeId string, ids []interface{}, messageId string, sessionId string) interface{} {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}
	entity, errObj := this.getEntityById(typeName, storeId, ids)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	// Here the entity with only it childs uuid's are return.
	if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
		return entity.(*DynamicEntity).getValues()
	}
	return entity

}

// @api 1.0
// Test if an entity with a given id(s) exist in the db.
// @param {string} typeName The object type name.
// @param {string} storeId The object type name.
// @param {string} ids The id's (not uuid) of the object to look for.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{bool} Return true if the entity exist.
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *EntityManager) IsEntityExist(typeName string, storeId string, ids []interface{}, messageId string, sessionId string) bool {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return false
	}

	var uuid string
	uuid, errObj = this.getEntityUuidById(typeName, storeId, ids)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return false
	}

	return len(uuid) > 0
}

// @api 1.0
// Take an array of id's in the same order as the entity prototype Id's and
// generate a dertermistic UUID from it.
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @result{string} Return Derteministic Universal Unique Identifier string
// @scope {public}
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *EntityManager) GenerateEntityUUID(typeName string, parentUuid string, ids []interface{}, messageId string, sessionId string) string {
	errObj := GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return ""
	}

	return generateEntityUuid(typeName, parentUuid, ids)
}
