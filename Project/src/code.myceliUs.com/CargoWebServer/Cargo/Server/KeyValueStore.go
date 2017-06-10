package Server

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/CargoWebServer/Cargo/JS"
	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/ast"
	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/lexer"
	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/parser"
	"code.myceliUs.com/Utility"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/xrash/smetrics"
)

////////////////////////////////////////////////////////////////////////////////
//                              DataStore function
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
//						Entity Query
////////////////////////////////////////////////////////////////////////////////
/**
 * The query is use to specifying the basic information it's like
 * the select, insert or update of sql...
 */
type EntityQuery struct {
	// The name of the entity
	TypeName string
	// The list of field to retreive, delete or modify
	Fields []string
	// The base index, this must be of form indexFieldName=indexFieldValue
	Indexs []string
	// The query to execute by the search engine.
	Query string
}

/**
 * This function is use to retreive the position in the array of a given field.
 */
func (this *EntityQuery) getFieldIndex(fieldName string) int {
	for i := 0; i < len(this.Fields); i++ {
		if this.Fields[i] == fieldName {
			return i
		}
	}
	return -1
}

/**
 * The search key is us to link the key value store with the seach engine.
 */
type SearchKey struct {
	TypeName string
	UUID     string
	Indexs   map[string]interface{}
}

////////////////////////////////////////////////////////////////////////////////
//			Key value Data Store
////////////////////////////////////////////////////////////////////////////////
type KeyValueDataStore struct {
	/** The store name **/
	m_id string

	/** The store path **/
	m_path string

	/** The runtime data store. **/
	m_db *leveldb.DB

	/**
	 * Use to protected the entitiesMap access...
	 */
	sync.RWMutex
}

func NewKeyValueDataStore(info *Config.DataStoreConfiguration) (store *KeyValueDataStore, err error) {
	store = new(KeyValueDataStore)
	store.m_id = info.M_id

	// Create the datastore if is not already exist and open it...
	store.m_path = GetServer().GetConfigurationManager().GetDataPath() + "/" + store.m_id
	store.m_db, err = leveldb.OpenFile(store.m_path, nil)

	if err != nil {
		log.Println("open:", err)
	}

	return
}

func (this *KeyValueDataStore) getValue(key string) ([]byte, error) {
	this.Lock()
	defer this.Unlock()
	values, err := this.m_db.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (this *KeyValueDataStore) getValues(uuid string) ([]interface{}, error) {
	this.Lock()
	defer this.Unlock()
	var results []interface{}
	values, err := this.m_db.Get([]byte(uuid), nil)
	if err != nil {
		return nil, err
	} else {
		dec := gob.NewDecoder(bytes.NewReader(values))
		dec.Decode(&results)
	}
	return results, nil
}

/**
 * Append value inside the key value store.
 */
func (this *KeyValueDataStore) setValue(key []byte, value []byte) error {
	this.Lock()
	defer this.Unlock()
	err := this.m_db.Put(key, value, nil)
	if err != nil {
		log.Println("Prototype encode:", err)
		return err
	}
	return nil
}

func (this *KeyValueDataStore) deleteValue(key string) error {
	this.Lock()
	defer this.Unlock()
	return this.m_db.Delete([]byte(key), nil)
}

/**
 * That function is use to create an extension of a given prototype.
 */
func (this *KeyValueDataStore) setSuperTypeFields(prototype *EntityPrototype) {

	for i := 0; i < len(prototype.SuperTypeNames); i++ {
		superTypeName := prototype.SuperTypeNames[0]
		superPrototype, err := GetServer().GetEntityManager().getEntityPrototype(superTypeName, superTypeName[0:strings.Index(superTypeName, ".")])
		if err == nil {
			// I will merge the fields
			// The first to fields are always the uuid and parentUuid and the last is the childUuids and referenced
			for j := 2; j < len(superPrototype.Fields)-2; j++ {
				if !Utility.Contains(prototype.Fields, superPrototype.Fields[j]) {
					Utility.InsertStringAt(2, superPrototype.Fields[j], &prototype.Fields)
					Utility.InsertStringAt(2, superPrototype.FieldsType[j], &prototype.FieldsType)
					Utility.InsertBoolAt(2, superPrototype.FieldsVisibility[j], &prototype.FieldsVisibility)

					// create a new index at the end...
					if superPrototype.FieldsNillable != nil {
						Utility.InsertBoolAt(2, superPrototype.FieldsNillable[j], &prototype.FieldsNillable)
					} else {
						prototype.FieldsNillable = append(prototype.FieldsNillable, true)
					}

					if superPrototype.FieldsDocumentation != nil {
						Utility.InsertStringAt(2, superPrototype.FieldsDocumentation[j], &prototype.FieldsDocumentation)
					} else {
						prototype.FieldsDocumentation = append(prototype.FieldsDocumentation, "")
					}
				}
			}
			// Now the index...
			for j := 0; j < len(superPrototype.Indexs); j++ {
				if !Utility.Contains(prototype.Indexs, superPrototype.Indexs[j]) {
					prototype.Indexs = append(prototype.Indexs, superPrototype.Indexs[j])
				}
			}
			// Now the ids
			for j := 0; j < len(superPrototype.Ids); j++ {
				if !Utility.Contains(prototype.Ids, superPrototype.Ids[j]) {
					prototype.Ids = append(prototype.Ids, superPrototype.Ids[j])
				}
			}
		}
	}

	// reset the field orders.
	prototype.FieldsOrder = make([]int, len(prototype.Fields))
	for i := 0; i < len(prototype.Fields); i++ {
		prototype.FieldsOrder[i] = i
	}

}

/**
 * This function is use to create a new entity prototype and save it value.
 * in db.
 * It must be create once per type
 */
func (this *KeyValueDataStore) SetEntityPrototype(prototype *EntityPrototype) error {

	// Save it only once...
	_, err := this.GetEntityPrototype(prototype.TypeName)

	if err != nil {
		// Here i will append super type fields...
		this.setSuperTypeFields(prototype)

		// I will serialyse the prototype.
		m := new(bytes.Buffer)
		enc := gob.NewEncoder(m)
		err := enc.Encode(prototype)
		err = this.setValue([]byte("prototype:"+prototype.TypeName), m.Bytes())
		if err != nil {
			log.Println("Prototype encode:", err)
			return err
		}
	}

	return nil
}

/**
 * Save an entity prototype.
 */
func (this *KeyValueDataStore) saveEntityPrototype(prototype *EntityPrototype) error {

	// Save it only once...
	_, err := this.GetEntityPrototype(prototype.TypeName)
	if err == nil {
		// I will serialyse the prototype.
		m := new(bytes.Buffer)
		enc := gob.NewEncoder(m)
		err := enc.Encode(prototype)
		err = this.setValue([]byte("prototype:"+prototype.TypeName), m.Bytes())
		if err != nil {
			log.Println("Prototype encode:", err)
			return err
		}
	}
	return nil
}

/**
 * Remove an entity prototype and all it releated values.
 */
func (this *KeyValueDataStore) DeleteEntityPrototype(typeName string) error {

	prototype, err := this.GetEntityPrototype(typeName)
	// The prototype does not exist.
	if err != nil {
		return err
	}

	// I will delete all entity...
	entities, _ := GetServer().GetEntityManager().getEntitiesByType(prototype.TypeName, "", this.m_id, false)
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		// remove it...
		entity.DeleteEntity()
	}

	// remove the prototype itself...
	this.deleteValue("prototype:" + prototype.TypeName)

	return nil
}

/**
 * Remove all prototypes.
 */
func (this *KeyValueDataStore) DeleteEntityPrototypes() error {
	prototypes, err := this.GetEntityPrototypes()
	if err != nil {
		return err
	}

	for i := 0; i < len(prototypes); i++ {
		err := this.DeleteEntityPrototype(prototypes[i].TypeName)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
 * This function is use to retreive an existing entity prototype...
 */
func (this *KeyValueDataStore) GetEntityPrototype(id string) (*EntityPrototype, error) {
	prototype := new(EntityPrototype)
	id = "prototype:" + id

	// Retreive the data from level db...
	data, err := this.getValue(id)
	if err != nil {
		return nil, err
	} else {
		dec := gob.NewDecoder(bytes.NewReader(data))
		dec.Decode(prototype)
	}

	return prototype, nil
}

/**
 * In case of multiple id value i will generate a unique key to be
 * use in the key value store. So The key can be regenerated as needed
 */
func (this *KeyValueDataStore) getKey(prototype *EntityPrototype, entity []interface{}) string {
	// The first value of the entity must be the key...
	return entity[prototype.getFieldIndex("UUID")].(string)
}

/**
 * Generate the indexation keys...
 */
func (this *KeyValueDataStore) getIndexationKeys(prototype *EntityPrototype, entity []interface{}) []string {
	var indexationKeys []string

	// I will also index supertype keys...
	prototypes := make([]*EntityPrototype, 1)
	prototypes[0] = prototype

	for i := 0; i < len(prototype.SuperTypeNames); i++ {
		superPrototype, _ := this.GetEntityPrototype(prototype.SuperTypeNames[i])
		if superPrototype != nil {
			prototypes = append(prototypes, superPrototype)
		}
	}

	// Now I will create the indexation...
	// first of all i will retreive existing indexation...
	for j := 0; j < len(prototypes); j++ {
		prototype := prototypes[j]

		// Set the type indexation
		indexationKeys = append(indexationKeys, prototype.TypeName)

		for i := 0; i < len(prototype.Indexs); i++ {
			var indexationKey string
			index := prototype.getFieldIndex(prototype.Indexs[i])
			// The indexation key is compose of the; type_name : filed_name : value_to_index
			switch v := entity[index].(type) {
			case string:
				if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
					values := make([]string, 0)
					err := json.Unmarshal([]byte(v), &values)
					if err == nil {
						for k := 0; k < len(values); k++ {
							indexationKey = prototype.TypeName + ":" + prototype.Indexs[i] + ":" + values[k]
							indexationKeys = append(indexationKeys, indexationKey)
						}
					}
				} else {
					if v != "null" {
						if len(strings.TrimSpace(v)) > 0 {
							indexationKey = prototype.TypeName + ":" + prototype.Indexs[i] + ":" + v
							indexationKeys = append(indexationKeys, indexationKey)
						}
					}
				}
			case int:
				indexationKey = prototype.TypeName + ":" + prototype.Indexs[i] + ":" + strconv.FormatInt(int64(v), 10)
				indexationKeys = append(indexationKeys, indexationKey)

			case int32:
				indexationKey = prototype.TypeName + ":" + prototype.Indexs[i] + ":" + strconv.FormatInt(int64(v), 10)
				indexationKeys = append(indexationKeys, indexationKey)

			case int64:
				indexationKey = prototype.TypeName + ":" + prototype.Indexs[i] + ":" + strconv.FormatInt(v, 10)
				indexationKeys = append(indexationKeys, indexationKey)

			case float64:
				indexationKey = prototype.TypeName + ":" + prototype.Indexs[i] + ":" + strconv.FormatFloat(v, 'f', 6, 64)
				indexationKeys = append(indexationKeys, indexationKey)

			case bool:
				indexationKey = prototype.TypeName + ":" + prototype.Indexs[i] + ":" + strconv.FormatBool(v)
				indexationKeys = append(indexationKeys, indexationKey)

			default:
				log.Println("--------> value can not be use as indexation key ", v, reflect.TypeOf(v).String())
			}
		}

		// Do the same for Id's
		for i := 0; i < len(prototype.Ids); i++ {
			var indexationKey string
			index := prototype.getFieldIndex(prototype.Ids[i])
			switch v := entity[index].(type) {
			case string:
				if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
					values := make([]string, 0)
					err := json.Unmarshal([]byte(v), &values)
					if err == nil {
						for k := 0; i < len(values); k++ {
							indexationKey = prototype.TypeName + ":" + prototype.Ids[i] + ":" + values[k]
							indexationKeys = append(indexationKeys, indexationKey)
						}
					}
				} else {
					if v != "null" {
						if len(strings.TrimSpace(v)) > 0 {
							indexationKey = prototype.TypeName + ":" + prototype.Ids[i] + ":" + v
							indexationKeys = append(indexationKeys, indexationKey)
						}
					}
				}

			case int8:
				indexationKey = prototype.TypeName + ":" + prototype.Ids[i] + ":" + strconv.Itoa(int(v))
				indexationKeys = append(indexationKeys, indexationKey)

			case int16:
				indexationKey = prototype.TypeName + ":" + prototype.Ids[i] + ":" + strconv.Itoa(int(v))
				indexationKeys = append(indexationKeys, indexationKey)

			case int:
				indexationKey = prototype.TypeName + ":" + prototype.Ids[i] + ":" + strconv.Itoa(int(v))
				indexationKeys = append(indexationKeys, indexationKey)

			case int32:
				indexationKey = prototype.TypeName + ":" + prototype.Ids[i] + ":" + strconv.Itoa(int(v))
				indexationKeys = append(indexationKeys, indexationKey)

			case int64:
				indexationKey = prototype.TypeName + ":" + prototype.Ids[i] + ":" + strconv.FormatInt(v, 10)
				indexationKeys = append(indexationKeys, indexationKey)

			case float64:
				indexationKey = prototype.TypeName + ":" + prototype.Ids[i] + ":" + strconv.FormatFloat(v, 'f', 6, 64)
				indexationKeys = append(indexationKeys, indexationKey)

			case bool:
				indexationKey = prototype.TypeName + ":" + prototype.Ids[i] + ":" + strconv.FormatBool(v)
				indexationKeys = append(indexationKeys, indexationKey)

			default:
				log.Println("--------> value can not be use as indexation key ", v, reflect.TypeOf(v).String())
			}
		}
	}
	return indexationKeys
}

/**
 * Retreive the list of all entity prototype in a given store.
 */
func (this *KeyValueDataStore) GetEntityPrototypes() ([]*EntityPrototype, error) {

	var prototypes []*EntityPrototype

	// Retreive values...
	this.Lock()
	iter := this.m_db.NewIterator(util.BytesPrefix([]byte("prototype:")), nil)
	this.Unlock()

	for iter.Next() {
		// Use key/value.
		value := iter.Value()

		// I will decode the prototype.
		prototype := new(EntityPrototype)
		dec := gob.NewDecoder(bytes.NewReader(value))
		dec.Decode(prototype)

		p, err := this.GetEntityPrototype(prototype.TypeName)
		if err != nil {
			return nil, err
		}

		// Append to the list of prototype
		prototypes = append(prototypes, p)
	}
	iter.Release()
	err := iter.Error()
	return prototypes, err
}

func (this *KeyValueDataStore) GetEntityByType(typeName string, storeId string) ([][]interface{}, error) {

	var entities [][]interface{}

	// Use key/value.
	// I will decode the prototype.
	ids, err := this.getIndexation(typeName)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(ids); i++ {
		id := ids[i].(string)
		var data []interface{}
		values, _ := this.getValue(id)
		dec := gob.NewDecoder(bytes.NewReader(values))
		dec.Decode(&data)

		// Append to the list of prototype
		entities = append(entities, data)
	}

	return entities, err
}

////////////////////////////////////////////////////////////////////////////////
//                              Indexation
////////////////////////////////////////////////////////////////////////////////

func (this *KeyValueDataStore) getIndexation(indexation string) ([]interface{}, error) {
	// Indexations contain array of string
	var ids []interface{}

	// I will retreive the value...
	values, err := this.getValue(indexation)

	if err != nil {
		return nil, err
	} else {
		dec := gob.NewDecoder(bytes.NewReader(values))
		dec.Decode(&ids)
	}

	return ids, nil
}

func (this *KeyValueDataStore) appendIndexation(indexation string, id string) error {

	// Indexations contain array of string
	ids, err := this.getIndexation(indexation)

	// I will retreive the value...
	if err != nil {
		// Here the indexation does not exist I will create it...
		ids = append(ids, id)
	} else {
		// I will look if the id is already there...
		exist := false
		for i := 0; i < len(ids) && exist == false; i++ {
			if ids[i] == id {
				exist = true
			}
		}
		if !exist {
			ids = append(ids, id)
		} else {
			// if the indexation already exit I do nothing...
			return nil
		}
	}

	// Encode the data and save it into the db.
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	enc.Encode(ids)

	err = this.setValue([]byte(indexation), data.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (this *KeyValueDataStore) deleteIndexation(indexationKey string, id string) error {

	// I will remove this id from the indexation.
	ids, err := this.getIndexation(indexationKey)
	if err != nil {
		// return...
		return err
	}

	newIds := make([]interface{}, 0)
	for i := 0; i < len(ids); i++ {
		if ids[i] != id {
			newIds = append(newIds, ids[i])
		}
	}

	if len(newIds) > 0 {
		var data bytes.Buffer
		enc := gob.NewEncoder(&data)
		enc.Encode(newIds)
		// Save the indexation whit he new value...
		err = this.setValue([]byte(indexationKey), data.Bytes())
	} else {
		// I will delete the indexation itself in that case...
		this.deleteValue(indexationKey)
	}

	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Search functionality.
////////////////////////////////////////////////////////////////////////////////

/**
 * Merge tow results in one...
 */
func (this *KeyValueDataStore) merge(r1 map[string]map[string]interface{}, r2 map[string]map[string]interface{}) map[string]map[string]interface{} {

	for k, v := range r1 {
		r2[k] = v
	}
	return r2
}

/**
 * Evaluate an expression.
 */
func (this *KeyValueDataStore) evaluate(typeName string, fieldName string, comparator string, expected interface{}, value interface{}) (bool, error) {
	isMatch := false

	// if the value is nil i will automatically return
	if value == nil {
		return isMatch, nil
	}

	prototype, err := this.GetEntityPrototype(typeName)
	if err != nil {
		return false, err
	}

	// The type name.
	fieldType := prototype.FieldsType[prototype.getFieldIndex(fieldName)]
	fieldType = strings.Replace(fieldType, "[]", "", -1)

	// here for the date I will get it unix time value...
	if fieldType == "xs.date" || fieldType == "xs.dateTime" {
		expectedDateValue, _ := Utility.MatchISO8601_Date(expected.(string))
		dateValue, _ := Utility.MatchISO8601_Date(value.(string))
		if fieldType == "xs.dateTime" {
			expected = expectedDateValue.Unix() // get the unix time for calcul
			value = dateValue.Unix()            // get the unix time for calcul
		} else {
			expected = expectedDateValue.Truncate(24 * time.Hour).Unix() // get the unix time for calcul
			value = dateValue.Truncate(24 * time.Hour).Unix()            // get the unix time for calcul
		}
	}

	if comparator == "==" {
		// Equality comparator.
		// Case of string type.
		if reflect.TypeOf(expected).Kind() == reflect.String && reflect.TypeOf(value).Kind() == reflect.String {
			isRegex := strings.HasPrefix(expected.(string), "/") && strings.HasSuffix(expected.(string), "/")
			if isRegex {
				// here I will try to match the regular expression.
				var err error
				isMatch, err = regexp.MatchString(expected.(string)[1:len(expected.(string))-1], value.(string))
				if err != nil {
					return false, err
				}
			} else {
				isMatch = Utility.RemoveAccent(expected.(string)) == Utility.RemoveAccent(value.(string))
			}
		} else if reflect.TypeOf(expected).Kind() == reflect.Bool && reflect.TypeOf(value).Kind() == reflect.Bool {
			return expected.(bool) == value.(bool), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return expected.(int64) == value.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return expected.(float64) == value.(float64), nil
		}
	} else if comparator == "~=" {
		// Approximation comparator, string only...
		// Case of string types.
		if reflect.TypeOf(expected).Kind() == reflect.String && reflect.TypeOf(value).Kind() == reflect.String {
			distance := smetrics.JaroWinkler(Utility.RemoveAccent(expected.(string)), Utility.RemoveAccent(value.(string)), 0.7, 4)
			isMatch = distance >= .85
		} else {
			return false, errors.New("Operator ~= can be only used with strings.")
		}
	} else if comparator == "!=" {
		// Equality comparator.
		// Case of string type.
		if reflect.TypeOf(expected).Kind() == reflect.String && reflect.TypeOf(value).Kind() == reflect.String {
			isMatch = Utility.RemoveAccent(expected.(string)) != Utility.RemoveAccent(value.(string))
		} else if reflect.TypeOf(expected).Kind() == reflect.Bool && reflect.TypeOf(value).Kind() == reflect.Bool {
			return expected.(bool) != value.(bool), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return expected.(int64) != value.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return expected.(float64) != value.(float64), nil
		}
	} else if comparator == "^=" {
		if reflect.TypeOf(expected).Kind() == reflect.String && reflect.TypeOf(value).Kind() == reflect.String {
			return strings.HasPrefix(value.(string), expected.(string)), nil
		} else {
			return false, nil
		}
	} else if comparator == "$=" {
		if reflect.TypeOf(expected).Kind() == reflect.String && reflect.TypeOf(value).Kind() == reflect.String {
			return strings.HasSuffix(value.(string), expected.(string)), nil
		} else {
			return false, nil
		}
	} else if comparator == "<" {
		// Number operator only...
		if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return value.(int64) < expected.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return value.(float64) < expected.(float64), nil
		}
	} else if comparator == "<=" {
		if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return value.(int64) <= expected.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return value.(float64) <= expected.(float64), nil
		}
	} else if comparator == ">" {
		if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return value.(int64) > expected.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return value.(float64) > expected.(float64), nil
		}
	} else if comparator == ">=" {
		if reflect.TypeOf(expected).Kind() == reflect.Int64 && reflect.TypeOf(value).Kind() == reflect.Int64 {
			return value.(int64) >= expected.(int64), nil
		} else if reflect.TypeOf(expected).Kind() == reflect.Float64 && reflect.TypeOf(value).Kind() == reflect.Float64 {
			return value.(float64) >= expected.(float64), nil
		}
	}

	return isMatch, nil
}

/**
 * That function test if a given value match all expressions of a given ast...
 */
func (this *KeyValueDataStore) match(ast *ast.QueryAst, values map[string]interface{}) (bool, error) {
	// test if the value is composite.
	if ast.IsComposite() {
		ast1, _, ast2 := ast.GetSubQueries()
		// both side of the tree must match.
		isMatch, err := this.match(ast1, values)
		if err != nil {
			return false, err
		}
		if isMatch == false {
			return false, nil
		}

		isMatch, err = this.match(ast2, values)
		if err != nil {
			return false, err
		}

		if isMatch == false {
			return false, nil
		}
	} else {
		// I will evaluate the expression...
		typeName, fieldName, comparator, expected := ast.GetExpression()
		return this.evaluate(typeName, fieldName, comparator, expected, values[fieldName])
	}

	return true, nil
}

/**
 * Here i will walk the tree and generate the query.
 */
func (this *KeyValueDataStore) runQuery(ast *ast.QueryAst, fields []string) (map[string]map[string]interface{}, error) {

	// I will create the array if it dosent exist.
	results := make(map[string]map[string]interface{}, 0)

	if ast.IsComposite() {
		// Get the sub-queries
		ast1, operator, ast2 := ast.GetSubQueries()

		r1, err := this.runQuery(ast1, fields)
		if err != nil {
			return nil, err
		}

		r2, err := this.runQuery(ast2, fields)
		if err != nil {
			return nil, err
		}

		if operator == "&&" { // conjonction
			for k, v := range r2 {
				isMatch, err := this.match(ast1, v)
				if err != nil {
					return nil, err
				}
				if isMatch {
					results[k] = v
				}
			}
			for k, v := range r1 {
				isMatch, err := this.match(ast2, v)
				if err != nil {
					return nil, err
				}
				if isMatch {
					results[k] = v
				}
			}
		} else if operator == "||" { // disjonction
			results = this.merge(r1, r2)
		}

	} else {

		typeName, fieldName, comparator, expected := ast.GetExpression()
		values := make(map[string]map[string]interface{}, 0)
		// Need the prototype here.
		prototype, err := this.GetEntityPrototype(typeName)
		if err != nil {
			return nil, err
		}
		fieldType := prototype.FieldsType[prototype.getFieldIndex(fieldName)]
		isArray := strings.HasPrefix(fieldType, "[]")
		isRef := strings.HasSuffix(fieldType, ":Ref")
		fieldType = strings.Replace(fieldType, "[]", "", -1)
		isString := fieldType == "xs.string" || fieldType == "xs.token" || fieldType == "xs.anyURI" || fieldType == "xs.anyURI" || fieldType == "xs.IDREF" || fieldType == "xs.QName" || fieldType == "xs.NOTATION" || fieldType == "xs.normalizedString" || fieldType == "xs.Name" || fieldType == "xs.NCName" || fieldType == "xs.ID" || fieldType == "xs.language"

		// Integers types.
		isInt := fieldType == "xs.int" || fieldType == "xs.integer" || fieldType == "xs.long" || fieldType == "xs.unsignedInt" || fieldType == "xs.short" || fieldType == "xs.unsignedLong"

		// decimal value
		isDecimal := fieldType == "xs.float" || fieldType == "xs.decimal" || fieldType == "xs.double"

		// Date time
		isDate := fieldType == "xs.date" || fieldType == "xs.dateTime"

		// I will append the indexs and the ids to the list of field if there's
		// not already in.
		for i := 0; i < len(prototype.Ids); i++ {
			if !Utility.Contains(fields, prototype.Ids[i]) {
				fields = append(fields, prototype.Ids[i])
			}
		}
		// The indexs
		for i := 0; i < len(prototype.Indexs); i++ {
			if !Utility.Contains(fields, prototype.Indexs[i]) {
				fields = append(fields, prototype.Indexs[i])
			}
		}

		// Strings or references...
		if isString || isRef {
			// The string expected value...
			expectedStr := expected.(string)
			isRegex := strings.HasPrefix(expectedStr, "/") && strings.HasSuffix(expectedStr, "/")
			if comparator == "==" && !isRegex {
				// Now i will get the value from the indexation.
				indexKey := typeName + ":" + fieldName + ":" + expectedStr
				indexations, err := this.getIndexation(indexKey)
				if err == nil {
					for i := 0; i < len(indexations); i++ {
						objects, err := this.getValues(indexations[i].(string))
						if err != nil {
							return nil, err
						}
						values[indexations[i].(string)] = make(map[string]interface{}, 0)
						for j := 0; j < len(fields); j++ {
							index := prototype.getFieldIndex(fields[j])
							values[indexations[i].(string)][fields[j]] = objects[index]
						}
						var isMatch bool
						if isArray {
							// Here I have an array of values to test.
							var strValues []string
							err = json.Unmarshal([]byte(values[indexations[i].(string)][fieldName].(string)), &strValues)
							if err != nil {
								return nil, err
							}
							for j := 0; j < len(strValues); j++ {
								isMatch, err = this.evaluate(typeName, fieldName, comparator, expected, strValues[j])
							}
						} else {
							isMatch, err = this.evaluate(typeName, fieldName, comparator, expected, values[indexations[i].(string)][fieldName])
						}

						if err != nil {
							return nil, err
						}
						if isMatch {
							// if the result match I put it inside the map result.
							results[indexations[i].(string)] = values[indexations[i].(string)]
						}
					}
				}
			} else if comparator == "~=" || comparator == "!=" || comparator == "^=" || comparator == "$=" || (isRegex && comparator == "==") {
				// Here I will use the typename as indexation key...
				indexations, err := this.getIndexation(typeName)
				if err == nil {
					for i := 0; i < len(indexations); i++ {
						objects, err := this.getValues(indexations[i].(string))
						if err != nil {
							return nil, err
						}
						values[indexations[i].(string)] = make(map[string]interface{}, 0)
						for j := 0; j < len(fields); j++ {
							index := prototype.getFieldIndex(fields[j])
							values[indexations[i].(string)][fields[j]] = objects[index]
						}
						isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[indexations[i].(string)][fieldName])
						if err != nil {
							return nil, err
						}
						if isMatch {
							// if the result match I put it inside the map result.
							results[indexations[i].(string)] = values[indexations[i].(string)]
						}
					}
				}
			} else {
				if !isRegex {
					return nil, errors.New("Unexpexted comparator " + comparator + " for type \"string\".")
				} else {
					return nil, errors.New("Unexpexted comparator " + comparator + " for regex, use \"==\" insted")
				}
			}
		} else if fieldType == "xs.boolean" {
			if !(comparator == "==" || comparator == "!=") {
				return nil, errors.New("Unexpexted comparator " + comparator + " for bool values, use \"==\" or  \"!=\"")
			}

			// Get the boolean value.
			indexKey := typeName + ":" + fieldName + ":" + strconv.FormatBool(expected.(bool))
			indexations, err := this.getIndexation(indexKey)
			if err == nil {
				for i := 0; i < len(indexations); i++ {
					objects, err := this.getValues(indexations[i].(string))
					if err != nil {
						return nil, err
					}
					values[indexations[i].(string)] = make(map[string]interface{}, 0)
					for j := 0; j < len(fields); j++ {
						index := prototype.getFieldIndex(fields[j])
						values[indexations[i].(string)][fields[j]] = objects[index]
					}
					isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[indexations[i].(string)][fieldName])
					if err != nil {
						return nil, err
					}
					if isMatch {
						// if the result match I put it inside the map result.
						results[indexations[i].(string)] = values[indexations[i].(string)]
					}
				}
			}
		} else if isInt || isDecimal || isDate { // Numeric values or date that are covert at evaluation time as integer.
			if comparator == "~=" {
				return nil, errors.New("Unexpexted comparator " + comparator + " for type numeric value.")
			}
			// Get the boolean value.
			if comparator == "==" {
				indexKey := typeName + ":" + fieldName + ":"
				if isDecimal {
					indexKey += strconv.FormatFloat(expected.(float64), 'f', 6, 64)
				} else if isInt {
					indexKey += strconv.FormatInt(expected.(int64), 10)
				} else if isDate {
					indexKey += expected.(string)
				}

				indexations, err := this.getIndexation(indexKey)
				if err == nil {
					for i := 0; i < len(indexations); i++ {
						objects, err := this.getValues(indexations[i].(string))
						if err != nil {
							return nil, err
						}

						values[indexations[i].(string)] = make(map[string]interface{}, 0)
						for j := 0; j < len(fields); j++ {
							index := prototype.getFieldIndex(fields[j])
							values[indexations[i].(string)][fields[j]] = objects[index]
						}
						isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[indexations[i].(string)][fieldName])
						if err != nil {
							return nil, err
						}
						if isMatch {
							// if the result match I put it inside the map result.
							results[indexations[i].(string)] = values[indexations[i].(string)]
						}
					}
				}
			} else {
				// for the other comparator I will get all the entities of the given type and test each of those.
				indexations, err := this.getIndexation(typeName)
				if err == nil {
					for i := 0; i < len(indexations); i++ {
						objects, err := this.getValues(indexations[i].(string))
						if err != nil {
							return nil, err
						}
						values[indexations[i].(string)] = make(map[string]interface{}, 0)
						for j := 0; j < len(fields); j++ {
							index := prototype.getFieldIndex(fields[j])
							values[indexations[i].(string)][fields[j]] = objects[index]
						}
						isMatch, err := this.evaluate(typeName, fieldName, comparator, expected, values[indexations[i].(string)][fieldName])
						if err != nil {
							return nil, err
						}
						if isMatch {
							// if the result match I put it inside the map result.
							results[indexations[i].(string)] = values[indexations[i].(string)]
						}
					}
				}
			}
		}
	}

	return results, nil
}

/**
 * Execute a search query.
 */
func (this *KeyValueDataStore) executeSearchQuery(query string, fields []string) ([][]interface{}, error) {
	s := lexer.NewLexer([]byte(query))
	p := parser.NewParser()
	a, err := p.Parse(s)
	if err == nil {
		astree := a.(*ast.QueryAst)
		fieldLength := len(fields)

		r, err := this.runQuery(astree, fields)
		if err != nil {
			return nil, err
		}

		// Here I will keep the result part...
		results := make([][]interface{}, 0)
		for _, object := range r {
			results_ := make([]interface{}, 0)
			for i := 0; i < fieldLength; i++ {
				results_ = append(results_, object[fields[i]])
			}
			results = append(results, results_)
		}

		return results, err
	} else {
		log.Println("--> search error ", err)
	}
	return nil, err
}

////////////////////////////////////////////////////////////////////////////////
//                              DataStore function
////////////////////////////////////////////////////////////////////////////////

/**
 * Return the name of a store.
 */
func (this *KeyValueDataStore) GetId() string {
	return this.m_id
}

// TODO validate the user and password here...
func (this *KeyValueDataStore) Connect() error {
	// Here I will register all class in the vm.
	prototypes, err := this.GetEntityPrototypes()
	if err == nil {
		for i := 0; i < len(prototypes); i++ {
			JS.GetJsRuntimeManager().AppendScript(prototypes[i].generateConstructor())
		}
	}

	return nil
}

/**
 * Help to know if a store is connect or existing...
 */
func (this *KeyValueDataStore) Ping() error {
	path := GetServer().GetConfigurationManager().GetDataPath() + "/" + this.GetId()
	_, err := os.Stat(path)
	return err
}

/**
 * Create a new entry in the database.
 */
func (this *KeyValueDataStore) Create(queryStr string, entity []interface{}) (lastId interface{}, err error) {
	// First of all i will init the query...
	var query EntityQuery
	json.Unmarshal([]byte(queryStr), &query)

	// now i will retreive the protoype for this entity...
	prototype, err := this.GetEntityPrototype(query.TypeName)

	// Key the key for that entity
	uuid := this.getKey(prototype, entity)

	// Each type of data will be insert
	// I will save the info for futur use...
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err = enc.Encode(entity)
	if err != nil {
		log.Panicln("encode:", err)
		return uuid, err
	}

	// Here I will save the data into the database...
	err = this.setValue([]byte(uuid), data.Bytes())
	if err != nil {
		log.Panicln("put value db:", err)
		return uuid, err
	}

	// Append the indexation for it type name...
	indexationKeys := this.getIndexationKeys(prototype, entity)
	for i := 0; i < len(indexationKeys); i++ {
		this.appendIndexation(indexationKeys[i], uuid)
	}

	// I will return the last uuid...
	lastId = uuid

	return uuid, err
}

/**
 * Get the value list...
 */
func (this *KeyValueDataStore) Read(queryStr string, fieldsType []interface{}, params []interface{}) (results [][]interface{}, err error) {

	// First of all i will init the query...
	var query EntityQuery
	json.Unmarshal([]byte(queryStr), &query)

	// now i will retreive the protoype for this entity...
	prototype, err := this.GetEntityPrototype(query.TypeName)
	if err != nil {
		return nil, err
	}

	if len(query.Query) > 0 {
		searchResult, err := this.executeSearchQuery(query.Query, query.Fields)
		if err != nil {
			return nil, err
		}
		return searchResult, nil
	} else {
		// Index are express as key values paire...
		// ex: "LogTime="+ strconv.FormatInt(this.LogTime, 10))
		indexs := make(map[string]string)
		for i := 0; i < len(query.Indexs); i++ {
			// create the function...
			index := query.Indexs[i]
			id := strings.TrimSpace(strings.Split(index, "=")[0])
			value := strings.TrimSpace(strings.Split(index, "=")[1])
			indexs[id] = value
		}

		// I will retreive ids here...
		var ids []string

		if len(indexs) > 0 {
			// The user give value of fields that match a given value.
			for key, value := range indexs {
				var indexationKey = query.TypeName + ":" + key + ":" + value
				keys, _ := this.getIndexation(indexationKey)
				for i := 0; i < len(keys); i++ {
					if !Utility.Contains(ids, keys[i].(string)) {
						ids = append(ids, keys[i].(string))
					}
				}
			}
		} else {
			// Use the type name as indexation...
			ids_, err := this.getIndexation(query.TypeName)
			if err == nil {
				for i := 0; i < len(ids_); i++ {
					ids = append(ids, ids_[i].(string))
				}
			}
		}

		// I will retreive the value...
		if err == nil {
			// Now I will get each entity by their id...
			for i := 0; i < len(ids); i++ {
				id := ids[i]
				var result []interface{}
				result, err = this.getValues(id)
				if err != nil {
					return
				} else {
					// Here I will insert only the value found in the
					// fields...
					var result_ []interface{}
					if len(query.Fields) > 0 && len(result) > 0 {
						// Here I will insert the fields asked by
						// the user...
						//log.Println(query.Fields, len(query.Fields))
						//log.Println(result, len(result))
						for k := 0; k < len(query.Fields); k++ {
							field := query.Fields[k]
							index := prototype.getFieldIndex(field)
							if index < len(result) && index != -1 {
								result_ = append(result_, result[index])
							} else {
								log.Println("------> Field ", field, " was not found for in result ", result, " at index ", index)
							}
						}
					}
					results = append(results, result_)
				}
			}
		} else {
			return
		}
	}
	return
}

/**
 * Update a entity value.
 */
func (this *KeyValueDataStore) Update(queryStr string, fields []interface{}, params []interface{}) (err error) {

	var query EntityQuery
	json.Unmarshal([]byte(queryStr), &query)

	// now i will retreive the protoype for this entity...
	prototype, err := this.GetEntityPrototype(query.TypeName)
	if err != nil {
		return err
	}

	fieldsType := make([]interface{}, 0)
	results, err := this.Read(string(queryStr), fieldsType, params)

	if err != nil {
		return err
	}

	// Here I will create a map of field and index...
	fieldIds := make(map[string]int, 0)
	for i := 0; i < len(prototype.Ids); i++ {
		fieldIds[prototype.Ids[i]] = i
	}

	fieldIndexs := make(map[string]int, 0)
	for i := 0; i < len(prototype.Indexs); i++ {
		fieldIndexs[prototype.Indexs[i]] = i
	}

	for i := 0; i < len(results); i++ {
		// The actual uuid...
		uuid := this.getKey(prototype, results[i])

		// get the actual entity
		var entity []interface{}
		entity, err = this.getValues(uuid)

		if err != nil {
			return
		}

		// Remove existing indexation
		old_indexationKeys := this.getIndexationKeys(prototype, entity)
		for j := 0; j < len(old_indexationKeys); j++ {
			this.deleteIndexation(old_indexationKeys[j], uuid)
		}

		// Update the entity values...
		for j := 0; j < len(fields); j++ {
			field := query.Fields[j]
			value := fields[j] // The new value to insert...

			// Now I will replace the value of the field...
			index := prototype.getFieldIndex(field)
			if index != -1 {
				entity[index] = value
			}
		}

		// Create the new indexations.
		new_indexationKeys := this.getIndexationKeys(prototype, entity)
		for j := 0; j < len(new_indexationKeys); j++ {
			this.appendIndexation(new_indexationKeys[j], uuid)
		}

		// And save the entity data..
		var data bytes.Buffer
		enc := gob.NewEncoder(&data)
		enc.Encode(entity)

		// Here I will save the data into the database...
		err = this.setValue([]byte(uuid), data.Bytes())
		if err != nil {
			log.Println("encode:", err)
			return
		}
	}
	return
}

/**
 * Delete entity from the store...
 */
func (this *KeyValueDataStore) Delete(queryStr string, params []interface{}) (err error) {
	// First of all i will init the query...
	var query EntityQuery
	json.Unmarshal([]byte(queryStr), &query)

	// now i will retreive the protoype for this entity...
	prototype, err := this.GetEntityPrototype(query.TypeName)
	if err != nil {
		return err
	}

	// Here I will create a map of field and index...
	fieldsIndex := make(map[string]int, 0)
	for i := 0; i < len(prototype.Ids); i++ {
		query.Fields = append(query.Fields, prototype.Ids[i])
		index := len(fieldsIndex)
		fieldsIndex[prototype.Ids[i]] = index
	}

	for i := 0; i < len(prototype.Indexs); i++ {
		query.Fields = append(query.Fields, prototype.Indexs[i])
		index := len(fieldsIndex)
		fieldsIndex[prototype.Indexs[i]] = index
	}

	// Here I will retreive the list of element to delete...
	fieldsType := make([]interface{}, 0)
	queryStr_, _ := json.Marshal(query)
	results, err := this.Read(string(queryStr_), fieldsType, params)

	for i := 0; i < len(results); i++ {

		// I will get the entity and remove it...
		uuid := this.getKey(prototype, results[i])

		// get the actual entity
		var entity []interface{}
		entity, err = this.getValues(uuid)

		// I will retreive it list of indexation key.
		indexationKeys := this.getIndexationKeys(prototype, entity)
		for j := 0; j < len(indexationKeys); j++ {
			this.deleteIndexation(indexationKeys[j], uuid)
		}

		err = this.deleteValue(uuid)
		if err != nil {
			//val, _ := this.getValue(uuid)
			log.Println("------------> fail to delete " + uuid)
		}

		val, err := this.getValue(uuid)
		if err == nil {
			log.Println("------------> fail to delete ", uuid, ":", string(val))
		}
	}

	return err
}

/**
 * Close the backend store.
 */
func (this *KeyValueDataStore) Close() error {
	this.Lock()
	defer this.Unlock()
	// Close the datastore.
	return this.m_db.Close()
}

////////////////////////////////////////////////////////////////////////////////
// Test for sql/xsd types.
////////////////////////////////////////////////////////////////////////////////
/**
 * Dertermine if the value is a base type.
 */
func isXsBaseType(fieldType string) bool {
	return isXsId(fieldType) || isXsRef(fieldType) || isXsInt(fieldType) || isXsString(fieldType) || isXsBinary(fieldType) || isXsNumeric(fieldType) || isXsBoolean(fieldType) || isXsDate(fieldType) || isXsTime(fieldType) || isXsMoney(fieldType)
}

/**
 * Helper function use to dertermine if a XS type must be considere integer.
 */
func isXsInt(fieldType string) bool {
	if strings.HasSuffix(fieldType, "byte") || strings.HasSuffix(fieldType, "int") || strings.HasSuffix(fieldType, "integer") || strings.HasSuffix(fieldType, "short") || strings.HasSuffix(fieldType, "unsignedInt") || strings.HasSuffix(fieldType, "unsignedBtype") || strings.HasSuffix(fieldType, "unsignedShort") || strings.HasSuffix(fieldType, "unsignedLong") || strings.HasSuffix(fieldType, "negativeInteger") || strings.HasSuffix(fieldType, "nonNegativeInteger") || strings.HasSuffix(fieldType, "nonPositiveInteger") || strings.HasSuffix(fieldType, "positiveInteger") || strings.HasSuffix(fieldType, "tinyint") || strings.HasSuffix(fieldType, "smallint") || strings.HasSuffix(fieldType, "bigint") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere String.
 */
func isXsString(fieldType string) bool {
	if strings.HasSuffix(fieldType, "string") || strings.HasSuffix(fieldType, "Name") || strings.HasSuffix(fieldType, "QName") || strings.HasSuffix(fieldType, "NMTOKEN") || strings.HasSuffix(fieldType, "gDay") || strings.HasSuffix(fieldType, "gMonth") || strings.HasSuffix(fieldType, "gMonthDay") || strings.HasSuffix(fieldType, "gYear") || strings.HasSuffix(fieldType, "gYearMonth") || strings.HasSuffix(fieldType, "token") || strings.HasSuffix(fieldType, "normalizedString") || strings.HasSuffix(fieldType, "hexBinary") || strings.HasSuffix(fieldType, "language") || strings.HasSuffix(fieldType, "NMTOKENS") || strings.HasSuffix(fieldType, "NOTATION") || strings.HasSuffix(fieldType, "char") || strings.HasSuffix(fieldType, "nchar") || strings.HasSuffix(fieldType, "varchar") || strings.HasSuffix(fieldType, "nvarchar") || strings.HasSuffix(fieldType, "text") || strings.HasSuffix(fieldType, "ntext") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere binary value.
 */
func isXsBinary(fieldType string) bool {
	if strings.HasSuffix(fieldType, "base64Binary") || strings.HasSuffix(fieldType, "varbinary") || strings.HasSuffix(fieldType, "binary") || strings.HasSuffix(fieldType, "image") || strings.HasSuffix(fieldType, "timestamp") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere numeric value.
 */
func isXsNumeric(fieldType string) bool {
	if strings.HasSuffix(fieldType, "double") || strings.HasSuffix(fieldType, "decimal") || strings.HasSuffix(fieldType, "float") || strings.HasSuffix(fieldType, "numeric") || strings.HasSuffix(fieldType, "real") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere boolean value.
 */
func isXsBoolean(fieldType string) bool {
	if strings.HasSuffix(fieldType, "boolean") || strings.HasSuffix(fieldType, "bit") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere date value.
 */
func isXsDate(fieldType string) bool {
	if strings.HasSuffix(fieldType, "date") || strings.HasSuffix(fieldType, "datetime") || strings.HasSuffix(fieldType, "datetime2") || strings.HasSuffix(fieldType, "smalldatetime") || strings.HasSuffix(fieldType, "datetimeoffset") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere time value.
 */
func isXsTime(fieldType string) bool {
	if strings.HasSuffix(fieldType, "time") || strings.HasSuffix(fieldType, "timestampNumeric") || strings.HasSuffix(fieldType, "timestamp") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere money value.
 */
func isXsMoney(fieldType string) bool {
	if strings.HasSuffix(fieldType, "money") || strings.HasSuffix(fieldType, "smallmoney") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere id value.
 */
func isXsId(fieldType string) bool {
	if strings.HasSuffix(fieldType, "ID") || strings.HasSuffix(fieldType, "NCName") || strings.HasSuffix(fieldType, "uniqueidentifier") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere id value.
 */
func isXsRef(fieldType string) bool {
	if strings.HasSuffix(fieldType, "anyURI") || strings.HasSuffix(fieldType, "IDREF") {
		return true
	}
	return false
}
