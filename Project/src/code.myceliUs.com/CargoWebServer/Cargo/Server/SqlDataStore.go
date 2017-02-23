package Server

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	_ "github.com/alexbrainman/odbc"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

////////////////////////////////////////////////////////////////////////////////
//              			Sql Data Store
////////////////////////////////////////////////////////////////////////////////
type SqlDataStore struct {
	/** The name of the store... **/
	m_id string

	/** The name of the drivers use **/
	m_vendor Config.DataStoreVendor

	/** The store port **/
	m_port int

	/** The store host **/
	m_host string

	/** The password **/
	m_password string

	/** The username **/
	m_user string

	/** The driver **/
	m_driver string

	/** The underlying database **/
	m_db *sql.DB
}

func NewSqlDataStore(info *Config.DataStoreConfiguration) (*SqlDataStore, error) {

	// The store...
	store := new(SqlDataStore)
	store.m_id = info.M_id

	// Keep this info...
	store.m_vendor = info.M_dataStoreVendor
	store.m_host = info.M_hostName
	store.m_user = info.M_user
	store.m_password = info.M_pwd
	store.m_port = info.M_port

	// Connect the store...
	err := store.Connect()

	return store, err
}

func (this *SqlDataStore) Connect() error {

	// The error message.
	var err error
	var connectionString string
	var driver string
	if this.m_vendor == Config.DataStoreVendor_MSSQL {
		/** Connect to Microsoft Sql server here... **/
		// So I will create the connection string from info...
		connectionString += "server=" + this.m_host + ";"
		connectionString += "user=" + this.m_user + ";"
		connectionString += "password=" + this.m_password + ";"
		connectionString += "port=" + strconv.Itoa(this.m_port) + ";"
		connectionString += "database=" + this.m_id + ";"
		connectionString += "driver=mssql"
		//connectionString += "encrypt=false;"
		driver = "mssql"

	} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
		/** Connect to oracle MySql server here... **/
		connectionString += this.m_user + ":"
		connectionString += this.m_password + "@tcp("
		connectionString += this.m_host + ":" + strconv.Itoa(this.m_port) + ")"
		connectionString += "/" + this.m_id
		//connectionString += "encrypt=false;"
		driver = "mysql"

	} else if this.m_vendor == Config.DataStoreVendor_ODBC {
		/** Connect with ODBC here... **/
		if runtime.GOOS == "windows" {
			connectionString += "driver=sql server;"
		} else {
			connectionString += "driver=freetds;"
		}
		connectionString += "server=" + this.m_host + ";"
		connectionString += "database=" + this.m_id + ";"
		connectionString += "uid=" + this.m_user + ";"
		connectionString += "pwd=" + this.m_password + ";"
		connectionString += "port=" + strconv.Itoa(this.m_port) + ";"
		connectionString += "clientcharset=UTF-8;"
		driver = "odbc"
	}

	this.m_db, err = sql.Open(driver, connectionString)

	return err
}

func (this *SqlDataStore) GetId() string {
	return this.m_id
}

/** Open a new Connection with the data store **/
func (this *SqlDataStore) Ping() (err error) {
	err = this.m_db.Ping()
	return err
}

/** Crud interface **/
func (this *SqlDataStore) Create(query string, data_ []interface{}) (lastId interface{}, err error) {

	err = this.Ping()
	if err != nil {
		err = this.Connect()
		if err != nil {
			return nil, err
		}
	}

	lastId = int64(0)
	stmt, err := this.m_db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(data_...)
	if err != nil {
		log.Fatal(err)
	}

	// First of all I need to find the key column id...
	// From the query I will retreive the table name.
	// the query follow this pattern... INSERT INTO TableNameHere(...
	startIndex := strings.Index(strings.ToUpper(query), " INTO ") + 5 // 5 is the len of INTO and one space...
	endIndex := strings.Index(query, "(")
	tableName := strings.TrimSpace(query[startIndex:endIndex])

	var colKeyQuery string
	if this.m_vendor == Config.DataStoreVendor_MSSQL || this.m_vendor == Config.DataStoreVendor_MYSQL {
		colKeyQuery = "SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE (TABLE_SCHEMA = '" + this.GetId() + "') AND (TABLE_NAME = ?) AND (COLUMN_KEY = 'PRI')"
	} else if this.m_vendor == Config.DataStoreVendor_ODBC {
		colKeyQuery = "SELECT COLUMN_NAME FROM " + this.GetId() + ".INFORMATION_SCHEMA.KEY_COLUMN_USAGE WHERE TABLE_NAME LIKE ? AND CONSTRAINT_NAME LIKE 'PK%'"
	}

	fieldType := make([]interface{}, 1, 1)
	fieldType[0] = "string"

	params := make([]interface{}, 1, 1)
	params[0] = tableName
	results, _ := GetServer().GetDataManager().readData(this.GetId(), colKeyQuery, fieldType, params)
	// Work for one column...
	if len(results) != 0 {
		if colKey, ok := results[0][0].(string); ok {
			// Now the second query...
			var lastIndexQuery string
			if this.m_vendor == Config.DataStoreVendor_MYSQL {
				lastIndexQuery = "SELECT " + colKey + " FROM " + tableName + " ORDER BY " + colKey + " DESC LIMIT 1"
			} else {
				lastIndexQuery = "SELECT TOP 1 " + colKey + " FROM " + tableName + " ORDER BY " + colKey + " DESC"
			}

			params = make([]interface{}, 0, 0)
			fieldType[0] = "int"
			results, _ = GetServer().GetDataManager().readData(this.GetId(), lastIndexQuery, fieldType, params)
			if len(results) > 0 {
				lastId = int64(results[0][0].(int))

			}
		}
	}

	if err == nil {
		eventData := make([]*MessageData, 3)

		tableName_ := new(MessageData)
		tableName_.Name = "tableName"
		tableName_.Value = tableName
		eventData[0] = tableName_

		id := new(MessageData)
		id.Name = "id"
		id.Value = lastId
		eventData[1] = id

		row := new(MessageData)
		row.Name = "values"
		row.Value = data_
		eventData[2] = row

		evt, _ := NewEvent(NewRowEvent, TableEvent, eventData)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}

	return
}

/**
 * Read a query execute it and return the result as an array of interface...
 */
func (this *SqlDataStore) Read(query string, fieldsType []interface{}, params []interface{}) ([][]interface{}, error) {

	err := this.Ping()
	if err != nil {
		err = this.Connect()
		if err != nil {
			log.Println("---> sql connection error:", err)
			return nil, err
		}
	}

	// TODO Try to figure out why the connection is lost...
	// Lost of connection so I will reconnect anyway...
	//this.Connect()

	rows, err := this.m_db.Query(query, params...)
	if err != nil {
		log.Println("---> sql read query error:", err)
		log.Println("query: ", query)
		return nil, err
	}

	defer rows.Close()
	results := make([][]interface{}, 0, 0)

	fields := make([]interface{}, len(fieldsType))
	dest := make([]interface{}, len(fieldsType)) // A temporary interface{} slice

	for i, _ := range fieldsType {
		dest[i] = &fields[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err := rows.Scan(dest...)
		if err != nil {

			return nil, err
		}
		result := make([]interface{}, 0, 0)

		for i := 0; i < len(fields); i++ {
			if fields[i] != nil {

				if fieldsType[i] == "int" {
					var val int
					switch fields[i].(type) {
					case int:
						{
							val = fields[i].(int)
						}
					case int32:
						{
							val = int(fields[i].(int32))
						}
					case int64:
						{
							val = int(fields[i].(int64))
						}
					case []uint8:
						{
							val, _ = strconv.Atoi(string(fields[i].([]uint8)))
						}
					}

					result = append(result, val)

				} else if fieldsType[i] == "bit" {
					val := fields[i].(bool)
					result = append(result, val)

				} else if fieldsType[i] == "real" || fieldsType[i] == "float" {

					switch fields[i].(type) {
					case float32:
						{
							val := fields[i].(float32)
							result = append(result, val)

						}
					case float64:
						{
							val := fields[i].(float64)
							result = append(result, val)

						}
					case []uint8:
						{
							val, _ := strconv.ParseFloat(string(fields[i].([]uint8)), 64)
							result = append(result, val)
						}
					}

				} else if fieldsType[i] == "bytes" {
					var val []byte
					val = fields[i].([]byte)
					result = append(result, val)

				} else if fieldsType[i] == "date" {
					var val time.Time
					val = fields[i].(time.Time)
					result = append(result, val)

				} else if fieldsType[i] == "string" {
					var val string
					switch fields[i].(type) {
					case []uint8: // Bytes...
						{
							val = string(fields[i].([]byte))
						}
					case string:
						{
							val = fields[i].(string)
						}
					}

					result = append(result, val)

				} else {
					log.Println("Type not found!!!")
				}
			} else {
				result = append(result, nil)
			}
		}
		results = append(results, result)
	}
	err = rows.Err()
	if err != nil {
		log.Println("---> sql error:", err)
		log.Println("quey: ", query)
		return nil, err
	}
	return results, nil
}

func (this *SqlDataStore) Update(query string, fields []interface{}, params []interface{}) (err error) {
	err = this.Ping()
	if err != nil {
		err = this.Connect()
		if err != nil {
			return err
		}
	}

	stmt, err := this.m_db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	var values []interface{}
	values = append(values, fields...)
	values = append(values, params...)

	_, err = stmt.Exec(values...)

	if err != nil {
		log.Fatal(err)
	}

	startIndex := strings.Index(strings.ToUpper(query), "UPDATE") + 6 // 5 is the len of INTO and one space...
	endIndex := strings.Index(strings.ToUpper(query), "SET")
	tableName := strings.TrimSpace(query[startIndex:endIndex])
	tableName = strings.Replace(tableName, "[", "", -1)
	tableName = strings.Replace(tableName, "]", "", -1)

	if err == nil {
		eventData := make([]*MessageData, 2)

		tableName_ := new(MessageData)
		tableName_.Name = "tableName"
		tableName_.Value = tableName
		eventData[0] = tableName_

		id := new(MessageData)
		id.Name = "id"
		if reflect.TypeOf(params[0]).String() == "string" {
			id.Value = params[0].(string) // Be sure to convert to the good type here...
		} else if reflect.TypeOf(params[0]).String() == "int" {
			id.Value = params[0].(int)
		} else if reflect.TypeOf(params[0]).String() == "int32" {
			id.Value = params[0].(int32)
		} else if reflect.TypeOf(params[0]).String() == "int64" {
			id.Value = params[0].(int64)
		} else if reflect.TypeOf(params[0]).String() == "float32" {
			id.Value = params[0].(float32)
		} else if reflect.TypeOf(params[0]).String() == "float64" {
			id.Value = params[0].(float64)
		} else {
			log.Println("----> unknow type: ", reflect.TypeOf(params[0]).String())
		}

		eventData[1] = id
		evt, _ := NewEvent(UpdateRowEvent, TableEvent, eventData)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}

	return
}

/**
 * Delete data that match a given pattern.
 */
func (this *SqlDataStore) Delete(query string, params []interface{}) (err error) {
	err = this.Ping()
	if err != nil {
		err = this.Connect()
		if err != nil {
			return err
		}
	}

	stmt, err := this.m_db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(params...)
	if err != nil {
		log.Println(err)
	}

	startIndex := strings.Index(strings.ToUpper(query), "FROM") + 5 // 5 is the len of INTO and one space...
	endIndex := strings.Index(query, "WHERE")
	tableName := strings.TrimSpace(query[startIndex:endIndex])
	tableName = strings.Replace(tableName, "[", "", -1)
	tableName = strings.Replace(tableName, "]", "", -1)

	if err == nil {
		var eventData []*MessageData

		tableName_ := new(MessageData)
		tableName_.Name = "tableName"
		tableName_.Value = tableName
		eventData = append(eventData, tableName_)

		for i := 0; i < len(params); i++ {
			id := new(MessageData)
			id.Name = "id_" + strconv.Itoa(i)
			if reflect.TypeOf(params[i]).String() == "string" {
				id.Value = params[i].(string) // Be sure to convert to the good type here...
			} else if reflect.TypeOf(params[i]).String() == "int" {
				id.Value = params[i].(int)
			} else if reflect.TypeOf(params[i]).String() == "int32" {
				id.Value = params[i].(int32)
			} else if reflect.TypeOf(params[i]).String() == "int64" {
				id.Value = params[i].(int64)
			} else if reflect.TypeOf(params[i]).String() == "float32" {
				id.Value = params[i].(float32)
			} else if reflect.TypeOf(params[i]).String() == "float64" {
				id.Value = params[i].(float64)
			} else {
				log.Println("----> unknow type: ", reflect.TypeOf(params[i]).String())
			}
			eventData = append(eventData, id)
		}

		evt, _ := NewEvent(DeleteRowEvent, TableEvent, eventData)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}

	return err
}

/**
 * Close the backend store.
 */
func (this *SqlDataStore) Close() error {
	return this.m_db.Close()
}

/**
 * Return the prototypes for all accessible table of a database.
 */
func (this *SqlDataStore) GetEntityPrototypes() ([]*EntityPrototype, error) {
	var prototypes []*EntityPrototype
	// Try to ping the sql server.
	err := this.Ping()
	if err != nil {
		err = this.Connect()
		if err != nil {
			return prototypes, err
		}
	}

	// So here I will get the list of table name that will be prototypes...
	query := "SELECT sobjects.name "
	query += "FROM sysobjects sobjects "
	query += "WHERE sobjects.xtype = 'U'"

	fieldsType := make([]interface{}, 1)
	fieldsType[0] = "string"

	var params []interface{}

	// Read the
	values, err := this.Read(query, fieldsType, params)
	if err != nil {
		return prototypes, err
	}
	for i := 0; i < len(values); i++ {
		prototype, err := this.GetEntityPrototype(values[i][0].(string))
		if err != nil {
			return prototypes, err
		}
		prototypes = append(prototypes, prototype)
	}

	return prototypes, nil
}

/**
 * Return the prototype of a given table.
 */
func (this *SqlDataStore) GetEntityPrototype(id string) (*EntityPrototype, error) {

	var prototype *EntityPrototype
	err := this.Ping()
	if err != nil {
		err = this.Connect()
		if err != nil {
			return prototype, err
		}
	}

	// Retreive the schema id.
	var schemaId string
	schemaId, err = this.getSchemaId(id)

	if err != nil {
		return nil, err
	}

	// If the data store is not found.
	store := GetServer().GetDataManager().m_dataStores["sql_info"]
	if store == nil {
		store, _ = GetServer().GetDataManager().createDataStore("sql_info", Config.DataStoreType_KEY_VALUE_STORE, Config.DataStoreVendor_MYCELIUS)
	} else {
		prototype, err = store.GetEntityPrototype(schemaId + "." + id)
		if err == nil {
			return prototype, nil
		}
	}

	// Initialyse the prototype.
	prototype = NewEntityPrototype()
	prototype.TypeName = schemaId + "." + id

	// Now I will retreive the list of field of the table.
	query := "SELECT "
	query += "c.name 'Column Name',"
	query += "t.Name 'Data type',"
	query += "c.max_length 'Max Length',"
	query += "c.precision ,"
	query += "c.scale ,"
	query += "c.is_nullable,"
	query += "ISNULL(i.is_primary_key, 0) 'Primary Key' "
	query += "FROM "
	query += "sys.columns c "
	query += "INNER JOIN "
	query += "sys.types t ON c.user_type_id = t.user_type_id "
	query += "LEFT OUTER JOIN "
	query += "sys.index_columns ic ON ic.object_id = c.object_id AND ic.column_id = c.column_id "
	query += "LEFT OUTER JOIN  "
	query += "sys.indexes i ON ic.object_id = i.object_id AND ic.index_id = i.index_id "
	query += "WHERE "
	query += "c.object_id = OBJECT_ID('" + id + "')"

	fieldsType := make([]interface{}, 7)
	fieldsType[0] = "string"
	fieldsType[1] = "string"
	fieldsType[2] = "int"
	fieldsType[3] = "int"
	fieldsType[4] = "int"
	fieldsType[5] = "bit"
	fieldsType[6] = "bit"

	var params []interface{}

	// Read the
	values, err := this.Read(query, fieldsType, params)
	if err != nil {
		return prototype, err
	}

	for i := 0; i < len(values); i++ {
		fieldName := values[i][0].(string)
		fieldName = strings.Replace(fieldName, "'", "''", -1)
		fieldTypeName := values[i][1].(string)
		fieldType, _ := this.getSqlTypePrototype(fieldTypeName)
		// If the field can be null
		isNilAble := values[i][5].(bool)
		// if the field is a key
		isId := values[i][6].(bool)
		// So here I will create new field...
		if fieldType != nil {
			prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
			refName, refType, err := this.getTypeNameRef(id, fieldName)
			if err == nil {
				// So here the attribute is a ref...
				prototype.Fields = append(prototype.Fields, "M_"+refName)
				prototype.FieldsType = append(prototype.FieldsType, refType)
			} else {
				prototype.Fields = append(prototype.Fields, "M_"+fieldName)
				prototype.FieldsType = append(prototype.FieldsType, fieldType.TypeName)
				if isId {
					prototype.Ids = append(prototype.Ids, "M_"+fieldName)
				}
			}
			if isNilAble {
				prototype.FieldsNillable = append(prototype.FieldsNillable, true)
			} else {
				prototype.FieldsNillable = append(prototype.FieldsNillable, false)
			}
			prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)

		}
	}

	err = store.(*KeyValueDataStore).SetEntityPrototype(prototype)
	if err == nil {
		log.Println("---> create ", prototype.TypeName, " prototype.")
	}
	return prototype, err
}

func (this *SqlDataStore) getSchemaId(name string) (string, error) {
	query := "SELECT SCHEMA_NAME(schema_id)"
	query += "AS SchemaTable "
	query += "FROM sys.tables "
	query += "WHERE  sys.tables.name = '" + name + "'"

	fieldsType := make([]interface{}, 1)
	fieldsType[0] = "string"
	var params []interface{}

	// Read the
	values, err := this.Read(query, fieldsType, params)
	if err != nil {
		return "", err
	}
	if len(values) > 0 {
		return this.m_id + "." + values[0][0].(string), nil
	}

	return "", errors.New("No schema found for table " + name)

}

func (this *SqlDataStore) getTypeNameRef(tableName string, fieldName string) (string, string, error) {
	fieldName = strings.Replace(fieldName, "'", "''", -1)

	query := "SELECT "
	query += "  OBJECT_NAME (f.referenced_object_id) AS ReferenceTableName, "
	query += "  f.name AS ForeignKey "
	query += "FROM "
	query += "  sys.foreign_keys AS f "
	query += "  INNER JOIN sys.foreign_key_columns AS fc ON f.OBJECT_ID = fc.constraint_object_id "
	query += "  INNER JOIN sys.objects AS o ON o.OBJECT_ID = fc.referenced_object_id "
	query += "WHERE "
	query += "    OBJECT_NAME(f.parent_object_id) = '" + tableName + "' and COL_NAME(fc.parent_object_id,fc.parent_column_id) = '" + fieldName + "'"

	fieldsType := make([]interface{}, 2)
	fieldsType[0] = "string"
	fieldsType[1] = "string"
	var params []interface{}

	// Read the
	values, err := this.Read(query, fieldsType, params)
	if err != nil {
		return "", "", err
	}

	if len(values) > 0 {
		if len(values[0]) == 2 {
			refName := values[0][1].(string)
			refTypeName := values[0][0].(string)

			refSchema, err := this.getSchemaId(refTypeName)
			if err != nil {
				log.Println(query)
				return "", "", err
			}
			return refName, refSchema + "." + refTypeName + ":Ref", nil
		}
	}

	return "", "", errors.New("No reference found for attribute " + tableName + "." + fieldName)
}

/**
 * Get the mapping of a given sql type.
 */
func (this *SqlDataStore) getSqlTypePrototype(typeName string) (*EntityPrototype, error) {
	prototype, err := GetServer().GetEntityManager().getEntityPrototype("sqltypes."+typeName, "sqltypes")
	return prototype, err
}
