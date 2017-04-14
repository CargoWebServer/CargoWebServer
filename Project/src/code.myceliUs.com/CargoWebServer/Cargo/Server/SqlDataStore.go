package Server

/**
* TODO append the other encoding charset...
 */
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
	"code.myceliUs.com/Utility"

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

	/** The charset **/
	m_textEncoding Config.Encoding

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

	/** The list of associations **/
	m_associations map[string]bool
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
	store.m_textEncoding = info.M_textEncoding

	// Keep the list of associations...
	store.m_associations = make(map[string]bool, 0)

	return store, nil
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
		connectionString += "?"
		// The encoding
		if this.m_textEncoding == Config.Encoding_UTF8 {
			connectionString += "charset=utf8;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_1 {
			connectionString += "charset=ISO8859_1;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_2 {
			connectionString += "charset=ISO8859_2;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_3 {
			connectionString += "charset=ISO8859_3;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_4 {
			connectionString += "charset=ISO8859_4;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_5 {
			connectionString += "charset=ISO8859_5;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_6 {
			connectionString += "charset=ISO8859_6;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_7 {
			connectionString += "charset=ISO8859_7;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_8 {
			connectionString += "charset=ISO8859_8;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_9 {
			connectionString += "charset=ISO8859_9;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_10 {
			connectionString += "charset=ISO8859_10;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_13 {
			connectionString += "charset=ISO8859_13;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_14 {
			connectionString += "charset=ISO8859_14;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_15 {
			connectionString += "charset=ISO8859_15;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_16 {
			connectionString += "charset=ISO8859_16;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1250 {
			connectionString += "charset=Windows1250;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1251 {
			connectionString += "charset=Windows1251;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1252 {
			connectionString += "charset=Windows1252;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1253 {
			connectionString += "charset=Windows1253;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1254 {
			connectionString += "charset=Windows1254;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1255 {
			connectionString += "charset=Windows1255;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1256 {
			connectionString += "charset=Windows1256;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1257 {
			connectionString += "charset=Windows1257;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1258 {
			connectionString += "charset=Windows1258;"
		} else if this.m_textEncoding == Config.Encoding_KOI8R {
			connectionString += "charset=KOI8R;"
		} else if this.m_textEncoding == Config.Encoding_KOI8U {
			connectionString += "charset=KOI8U;"
		}
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

		// The encoding
		if this.m_textEncoding == Config.Encoding_UTF8 {
			connectionString += "charset=UTF8;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_1 {
			connectionString += "charset=ISO8859_1;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_2 {
			connectionString += "charset=ISO8859_2;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_3 {
			connectionString += "charset=ISO8859_3;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_4 {
			connectionString += "charset=ISO8859_4;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_5 {
			connectionString += "charset=ISO8859_5;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_6 {
			connectionString += "charset=ISO8859_6;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_7 {
			connectionString += "charset=ISO8859_7;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_8 {
			connectionString += "charset=ISO8859_8;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_9 {
			connectionString += "charset=ISO8859_9;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_10 {
			connectionString += "charset=ISO8859_10;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_13 {
			connectionString += "charset=ISO8859_13;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_14 {
			connectionString += "charset=ISO8859_14;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_15 {
			connectionString += "charset=ISO8859_15;"
		} else if this.m_textEncoding == Config.Encoding_ISO8859_16 {
			connectionString += "charset=ISO8859_16;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1250 {
			connectionString += "charset=Windows1250;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1251 {
			connectionString += "charset=Windows1251;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1252 {
			connectionString += "charset=Windows1252;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1253 {
			connectionString += "charset=Windows1253;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1254 {
			connectionString += "charset=Windows1254;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1255 {
			connectionString += "charset=Windows1255;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1256 {
			connectionString += "charset=Windows1256;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1257 {
			connectionString += "charset=Windows1257;"
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1258 {
			connectionString += "charset=Windows1258;"
		} else if this.m_textEncoding == Config.Encoding_KOI8R {
			connectionString += "charset=KOI8R;"
		} else if this.m_textEncoding == Config.Encoding_KOI8U {
			connectionString += "charset=KOI8U;"
		}
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
	if this.m_db == nil {
		err = errors.New("No connection was found for datastore " + this.m_id)
		return err
	}

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
		//log.Fatal(err)
		return -1, err
	}
	_, err = stmt.Exec(data_...)
	if err != nil {
		//log.Fatal(err)
		return -1, err
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
 * Cast sql type into it go equivalent type.
 */
func (this *SqlDataStore) castSqlType(sqlTypeName string, value interface{}) interface{} {

	/////////////////////////// Integer ////////////////////////////////
	if strings.HasSuffix(sqlTypeName, "tinyint") || strings.HasSuffix(sqlTypeName, "smallint") || strings.HasSuffix(sqlTypeName, "int") || strings.HasSuffix(sqlTypeName, "bigint") || strings.HasSuffix(sqlTypeName, "timestampNumeric") {
		var val int
		switch value.(type) {
		case int:
			{
				val = value.(int)
			}
		case int32:
			{
				val = int(value.(int32))
			}
		case int64:
			{
				val = int(value.(int64))
			}
		case []uint8:
			{
				val, _ = strconv.Atoi(string(value.([]uint8)))
			}
		}
		return val
	}

	/////////////////////////// Boolean ////////////////////////////////
	if strings.HasSuffix(sqlTypeName, "bit") {
		if reflect.TypeOf(value).Kind() == reflect.Bool {
			return value.(bool)
		} else if reflect.TypeOf(value).String() == "[]uint8" {
			if string(value.([]uint8)) == "1" {
				return true
			} else {
				return false
			}
		}
	}

	/////////////////////////// Numeric ////////////////////////////////
	if strings.HasSuffix(sqlTypeName, "real") || strings.HasSuffix(sqlTypeName, "float") || strings.HasSuffix(sqlTypeName, "decimal") || strings.HasSuffix(sqlTypeName, "numeric") || strings.HasSuffix(sqlTypeName, "money") || strings.HasSuffix(sqlTypeName, "smallmoney") {
		switch value.(type) {
		case float32:
			{
				val := value.(float32)
				return val
			}
		case float64:
			{
				val := value.(float64)
				return val
			}
		case []uint8:
			{
				val, _ := strconv.ParseFloat(string(value.([]uint8)), 64)
				return val
			}
		}
	}

	//////////////////////////// Date ////////////////////////////////
	if strings.HasSuffix(sqlTypeName, "date") || strings.HasSuffix(sqlTypeName, "datetime") {
		if reflect.TypeOf(value).String() == "time.Time" {
			var val time.Time
			val = value.(time.Time)
			return val
		} else if reflect.TypeOf(value).String() == "[]uint8" {
			t, err := time.Parse(time.RFC3339Nano, string(value.([]uint8)))
			if err == nil {
				return t
			}
		}
	}

	/////////////////////////// string ////////////////////////////////
	if strings.HasSuffix(sqlTypeName, "nchar") || strings.HasSuffix(sqlTypeName, "varchar") || strings.HasSuffix(sqlTypeName, "nvarchar") || strings.HasSuffix(sqlTypeName, "text") || strings.HasSuffix(sqlTypeName, "ntext") {
		var val string

		switch value.(type) {
		case []uint8: // Bytes...
			{
				val = string(value.([]byte))
			}
		case string:
			{
				val = value.(string)
			}
		}

		// If values need decoding.
		if this.m_textEncoding == Config.Encoding_ISO8859_1 {
			val, _ = Utility.DecodeISO8859_1(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_2 {
			val, _ = Utility.DecodeISO8859_2(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_3 {
			val, _ = Utility.DecodeISO8859_3(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_4 {
			val, _ = Utility.DecodeISO8859_4(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_5 {
			val, _ = Utility.DecodeISO8859_5(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_6 {
			val, _ = Utility.DecodeISO8859_6(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_7 {
			val, _ = Utility.DecodeISO8859_7(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_8 {
			val, _ = Utility.DecodeISO8859_8(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_9 {
			val, _ = Utility.DecodeISO8859_9(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_10 {
			val, _ = Utility.DecodeISO8859_10(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_13 {
			val, _ = Utility.DecodeISO8859_13(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_14 {
			val, _ = Utility.DecodeISO8859_14(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_15 {
			val, _ = Utility.DecodeISO8859_15(val)
		} else if this.m_textEncoding == Config.Encoding_ISO8859_16 {
			val, _ = Utility.DecodeISO8859_16(val)
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1250 {
			val, _ = Utility.DecodeWindows1250(val)
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1251 {
			val, _ = Utility.DecodeWindows1251(val)
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1252 {
			val, _ = Utility.DecodeWindows1252(val)
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1253 {
			val, _ = Utility.DecodeWindows1253(val)
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1254 {
			val, _ = Utility.DecodeWindows1254(val)
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1255 {
			val, _ = Utility.DecodeWindows1255(val)
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1256 {
			val, _ = Utility.DecodeWindows1256(val)
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1257 {
			val, _ = Utility.DecodeWindows1257(val)
		} else if this.m_textEncoding == Config.Encoding_WINDOWS_1258 {
			val, _ = Utility.DecodeWindows1258(val)
		} else if this.m_textEncoding == Config.Encoding_KOI8R {
			val, _ = Utility.DecodeKOI8R(val)
		} else if this.m_textEncoding == Config.Encoding_KOI8U {
			val, _ = Utility.DecodeKOI8U(val)
		}
		return val
	}

	/////////////////////////// binary ////////////////////////////////
	if strings.HasSuffix(sqlTypeName, "varbinary") || strings.HasSuffix(sqlTypeName, "binary") || strings.HasSuffix(sqlTypeName, "image") || strings.HasSuffix(sqlTypeName, "timestamp") {
		var val string

		switch value.(type) {
		case []uint8: // Bytes...
			{
				val = string(value.([]byte))
			}
		case string:
			{
				val = value.(string)
			}
		}

		// Must be a base 64 string.
		return val
	}

	/////////////////////////// other values ////////////////////////////////
	// uniqueidentifier, time
	var val string
	switch value.(type) {
	case []uint8: // Bytes...
		{
			val = string(value.([]byte))
		}
	case string:
		{
			val = value.(string)
		}
	}

	// Must be a base 64 string.
	return val

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
				val := this.castSqlType(fieldsType[i].(string), fields[i])
				result = append(result, val)
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

	// Close after use
	defer stmt.Close()

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
		id.Value = Utility.ToString(params[0])
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

	// Close after use
	defer stmt.Close()

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
			id.Value = Utility.ToString(params[i])
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
	var query string

	fieldsType := make([]interface{}, 0)
	// So here I will get the list of table name that will be prototypes...
	if this.m_vendor == Config.DataStoreVendor_ODBC {
		query = "SELECT sobjects.name "
		query += "FROM sysobjects sobjects "
		query += "WHERE sobjects.xtype = 'U'"

	} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
		query = "SELECT table_name FROM information_schema.tables where table_schema='" + this.m_id + "'"
	}

	fieldsType = append(fieldsType, "nvarchar")
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

	// Complete the reference information.
	this.setRefs()

	// Synchronize actual data
	this.synchronize(prototypes)

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
	var query string
	if this.m_vendor == Config.DataStoreVendor_ODBC {
		query = "SELECT "
		query += "c.name 'Column Name',"
		query += "t.Name 'Data type',"
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
	} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
		query = "SELECT "
		query += "column_name,"
		query += "data_type,"
		query += "IS_NULLABLE = 'YES',"
		query += "COLUMN_KEY = 'PRI' "
		query += "FROM information_schema.columns "
		query += "WHERE table_name='" + id + "';"
	}

	fieldsType := make([]interface{}, 4)
	fieldsType[0] = "nvarchar"
	fieldsType[1] = "nvarchar"
	fieldsType[2] = "bit"
	fieldsType[3] = "bit"

	var params []interface{}

	// Read the
	values, err := this.Read(query, fieldsType, params)
	if err != nil {
		log.Println("-------------->820 prototype: ", prototype)
		return prototype, err
	}

	for i := 0; i < len(values); i++ {
		fieldName := values[i][0].(string)
		fieldName = strings.Replace(fieldName, "'", "''", -1)
		fieldTypeName := values[i][1].(string)
		fieldType, _ := this.getSqlTypePrototype(fieldTypeName)
		// If the field can be null
		isNilAble := values[i][2].(bool)
		// if the field is a key
		isId := values[i][3].(bool)
		// So here I will create new field...
		if fieldType != nil {
			if !Utility.Contains(prototype.Fields, "M_"+fieldName) {
				prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
				prototype.Fields = append(prototype.Fields, "M_"+fieldName)
				prototype.FieldsType = append(prototype.FieldsType, fieldType.TypeName)
				prototype.FieldsDocumentation = append(prototype.FieldsDocumentation, "")

				// Set as id.
				if isId {
					prototype.Ids = append(prototype.Ids, "M_"+fieldName)
				}

				if isNilAble {
					prototype.FieldsNillable = append(prototype.FieldsNillable, true)
				} else {
					prototype.FieldsNillable = append(prototype.FieldsNillable, false)
				}

				prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
			}
		}
	}

	// Create the new prototype in sql_info store.
	prototype.Create("sql_info")

	log.Println("-------------->860 prototype: ", prototype)

	return prototype, err
}

/**
 * Determine if a relation is a reference or not.
 */
func (this *SqlDataStore) isRef(relationName string) bool {
	var query string
	if this.m_vendor == Config.DataStoreVendor_ODBC {
		query = "SELECT delete_referential_action_desc "
		query += "FROM sys.foreign_keys "
		query += "WHERE  name = '" + relationName + "'"
	} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
		query = "SELECT DELETE_RULE FROM "
		query += "information_schema.REFERENTIAL_CONSTRAINTS "
		query += "WHERE CONSTRAINT_NAME = '" + relationName + "'"
	}

	fieldsType := make([]interface{}, 1)
	fieldsType[0] = "nvarchar"
	var params []interface{}

	// Read the
	values, err := this.Read(query, fieldsType, params)
	if err != nil {
		return true
	}

	if len(values) > 0 {
		if values[0][0] != nil {
			return values[0][0].(string) != "CASCADE"
		}
	}
	return true

}

func (this *SqlDataStore) getSchemaId(name string) (string, error) {

	if this.m_vendor == Config.DataStoreVendor_ODBC {
		var query string

		query = "SELECT SCHEMA_NAME(schema_id)"
		query += "AS SchemaTable "
		query += "FROM sys.tables "
		query += "WHERE  sys.tables.name = '" + name + "'"
		fieldsType := make([]interface{}, 1)
		fieldsType[0] = "nvarchar"
		var params []interface{}

		// Read the
		values, err := this.Read(query, fieldsType, params)
		if err != nil {
			return "", err
		}
		if len(values) > 0 {
			if values[0][0] != nil {
				return this.m_id + "." + values[0][0].(string), nil
			}
		}
	} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
		// with mysql the schema id is the id of the database.
		return this.m_id, nil
	}

	return "", errors.New("No schema found for table " + name)

}

func appendField(prototype *EntityPrototype, fieldName string, fieldType string) {

	if !Utility.Contains(prototype.Fields, fieldName) {

		prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
		prototype.FieldsNillable = append(prototype.FieldsVisibility, true)

		fields := prototype.Fields
		fieldsType := prototype.FieldsType
		fieldsVisibility := prototype.FieldsVisibility
		//fieldsDocumentation := prototype.FieldsDocumentation

		// Set empty array...
		prototype.Fields = make([]string, len(prototype.FieldsOrder))
		prototype.FieldsType = make([]string, len(prototype.FieldsOrder))
		prototype.FieldsVisibility = make([]bool, len(prototype.FieldsOrder))
		prototype.FieldsDocumentation = make([]string, len(prototype.FieldsOrder))

		for i := 0; i < len(fields)-2; i++ {
			prototype.Fields[i] = fields[i]
			prototype.FieldsType[i] = fieldsType[i]
			prototype.FieldsVisibility[i] = fieldsVisibility[i]
			//prototype.FieldsDocumentation[i] = fieldsDocumentation[i]
		}

		lastIndex := len(prototype.FieldsOrder) - 1
		prototype.Fields[lastIndex-2] = fieldName
		prototype.FieldsType[lastIndex-2] = fieldType
		prototype.FieldsVisibility[lastIndex-2] = true
		//prototype.FieldsDocumentation[lastIndex-2] = ""

		// childsUuid
		prototype.Fields[lastIndex-1] = fields[len(fields)-2]
		prototype.FieldsType[lastIndex-1] = fieldsType[len(fields)-2]
		prototype.FieldsVisibility[lastIndex-1] = fieldsVisibility[len(fields)-2]
		//prototype.FieldsDocumentation[lastIndex-1] = fieldsDocumentation[len(fields)-2]

		// referenced
		prototype.Fields[lastIndex] = fields[len(fields)-1]
		prototype.FieldsType[lastIndex] = fieldsType[len(fields)-1]
		prototype.FieldsVisibility[lastIndex] = fieldsVisibility[len(fields)-1]
		//prototype.FieldsDocumentation[lastIndex] = fieldsDocumentation[len(fields)-1]

		// Save it back.
		GetServer().GetDataManager().getDataStore("sql_info").(*KeyValueDataStore).saveEntityPrototype(prototype)
	}
}

// Return true if the prototype is an associative one...
func (this *SqlDataStore) isAssociative(prototype *EntityPrototype) bool {
	// if the asscociation already exist.
	if val, ok := this.m_associations[prototype.TypeName]; ok {
		return val
	}

	for j := 0; j < len(prototype.Fields); j++ {
		if strings.HasPrefix(prototype.FieldsType[j], "sqltypes.") {
			if strings.HasPrefix(prototype.Fields[j], "M_") {
				refInfos, _ := this.getRefInfos(prototype.Fields[j][2:])
				if len(refInfos) == 0 {
					if !Utility.Contains(prototype.Ids, prototype.Fields[j]) {
						//log.Println("-----> ", prototype.TypeName, " is not associative")
						this.m_associations[prototype.TypeName] = false
						return false
					}
				}
			}
		}
	}

	this.m_associations[prototype.TypeName] = true
	//log.Println("-----> ", prototype.TypeName, " is associative")
	return true
}

// For a given ref id it return, the schma name, the table name,
// the table column name, the referenced_table name and the referenced column name.
// TODO handle multiple refs...
func (this *SqlDataStore) getRefInfos(refId string) ([]string, error) {

	var results []string
	var query string
	if this.m_vendor == Config.DataStoreVendor_ODBC {
		query = "SELECT "
		query += "	tab1.name     AS [table],"
		query += "	col1.name     AS [column],"
		query += "	tab2.name     AS [referenced_table],"
		query += "	col2.name     AS [referenced_column], "
		query += "	sch.name      AS [schema_name] "
		query += "FROM "
		query += "	sys.foreign_key_columns fkc "
		query += "INNER JOIN sys.objects obj "
		query += "	ON obj.object_id = fkc.constraint_object_id "
		query += "INNER JOIN sys.tables tab1 "
		query += "	ON tab1.object_id = fkc.parent_object_id "
		query += "INNER JOIN sys.schemas sch "
		query += "	ON tab1.schema_id = sch.schema_id "
		query += "INNER JOIN sys.columns col1 "
		query += "	ON col1.column_id = parent_column_id AND col1.object_id = tab1.object_id "
		query += "INNER JOIN sys.tables tab2 "
		query += "	ON tab2.object_id = fkc.referenced_object_id "
		query += "INNER JOIN sys.columns col2 "
		query += "	ON col2.column_id = referenced_column_id "
		query += "	AND col2.object_id =  tab2.object_id "
		query += "WHERE obj.name ='" + refId + "'"
	} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
		query = "select "
		query += "	table_name, "
		query += "	column_name, "
		query += "	referenced_table_name, "
		query += "	referenced_column_name, "
		query += "	table_schema "
		query += "from "
		query += "	information_schema.key_column_usage "
		query += "where "
		query += "	referenced_table_name is not null "
		query += "	and constraint_name = '" + refId + "'"
	}

	fieldsType := make([]interface{}, 5)
	fieldsType[0] = "nvarchar"
	fieldsType[1] = "nvarchar"
	fieldsType[2] = "nvarchar"
	fieldsType[3] = "nvarchar"
	fieldsType[4] = "nvarchar"

	var params []interface{}

	// Read the
	values, err := this.Read(query, fieldsType, params)
	if err != nil {
		log.Println("=-------|> ", err)
		log.Println(query)
		return results, err
	} else if len(values) > 0 {
		if len(values[0]) == 5 {
			results = make([]string, 5)
			// Set the value inside the results.
			results[0] = values[0][0].(string)
			results[1] = values[0][1].(string)
			results[2] = values[0][2].(string)
			results[3] = values[0][3].(string)
			results[4] = values[0][4].(string)
		}
		return results, nil
	}
	return results, nil
}

// Must be call after all protoypes are created.
// That function will complete references information.
func (this *SqlDataStore) setRefs() error {

	var query string
	if this.m_vendor == Config.DataStoreVendor_ODBC {
		query = "SELECT "
		query += "	obj.name      AS FK_NAME,"
		query += "	sch.name      AS [schema_name],"
		query += "	tab1.name     AS [table],"
		query += "	col1.name     AS [column],"
		query += "	tab2.name     AS [referenced_table],"
		query += "	col2.name     AS [referenced_column] "
		query += "FROM "
		query += "	sys.foreign_key_columns fkc "
		query += "INNER JOIN sys.objects obj "
		query += "	ON obj.object_id = fkc.constraint_object_id "
		query += "INNER JOIN sys.tables tab1 "
		query += "	ON tab1.object_id = fkc.parent_object_id "
		query += "INNER JOIN sys.schemas sch "
		query += "	ON tab1.schema_id = sch.schema_id "
		query += "INNER JOIN sys.columns col1 "
		query += "	ON col1.column_id = parent_column_id AND col1.object_id = tab1.object_id "
		query += "INNER JOIN sys.tables tab2 "
		query += "	ON tab2.object_id = fkc.referenced_object_id "
		query += "INNER JOIN sys.columns col2 "
		query += "	ON col2.column_id = referenced_column_id "
		query += "	AND col2.object_id =  tab2.object_id; "
	} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
		query = "select "
		query += "	constraint_name,"
		query += "	table_schema,"
		query += "	table_name, "
		query += "	column_name, "
		query += "	referenced_table_name, "
		query += "	referenced_column_name "
		query += "from "
		query += "	information_schema.key_column_usage "
		query += "where "
		query += "	referenced_table_name is not null "
		query += "	and table_schema = '" + this.m_id + "'"
	}

	// TODO create the query for mySql here.
	fieldsType := make([]interface{}, 6)
	fieldsType[0] = "nvarchar"
	fieldsType[1] = "nvarchar"
	fieldsType[2] = "nvarchar"
	fieldsType[3] = "nvarchar"
	fieldsType[4] = "nvarchar"
	fieldsType[5] = "nvarchar"

	var params []interface{}

	// Read the
	values, err := this.Read(query, fieldsType, params)
	if err != nil {
		log.Println("------> error ", err)
		return err
	}

	// Keep assiative table in a map.
	associativeTables := make(map[string]*EntityPrototype, 0)

	// Keep association values temporary...
	associations := make(map[string][][]interface{}, 0)

	if len(values) > 0 {
		for i := 0; i < len(values); i++ {
			refName := values[i][0].(string)
			schemasName := values[i][1].(string) + "."
			if this.m_vendor == Config.DataStoreVendor_MYSQL {
				schemasName = ""
			}

			// Source
			sourceTableName := values[i][2].(string)

			// Target.
			targetTableName := values[i][4].(string)

			src, err := GetServer().GetEntityManager().getEntityPrototype(this.m_id+"."+schemasName+sourceTableName, "sql_info")
			if err != nil {
				log.Println("----------------> src not found: ", err)
				return err
			}

			trg, err := GetServer().GetEntityManager().getEntityPrototype(this.m_id+"."+schemasName+targetTableName, "sql_info")
			if err != nil {
				log.Println("----------------> trg not found: ", err)
				return err
			}

			// Test if source field and target field are id's
			sourceField_isId := Utility.Contains(src.Ids, "M_"+values[i][3].(string))

			// I will set the field as non visible, the ref will be display
			//src.FieldsVisibility[src.getFieldIndex("M_"+values[i][3].(string))] = false

			targetField_isId := Utility.Contains(trg.Ids, "M_"+values[i][5].(string))

			// I will append the field if is not already there.
			if !this.isAssociative(src) {
				// Now the sources.
				if !Utility.Contains(src.Fields, "M_"+refName) {

					// Now the rule to determine the cardinality.
					var fieldType string
					fieldType = this.m_id + "." + schemasName + targetTableName + ":Ref"

					if !targetField_isId {
						// one to one relationship in that case.
						fieldType = "[]" + fieldType
					}

					appendField(src, "M_"+refName, fieldType)
				}

				if !Utility.Contains(trg.Fields, "M_"+refName) {
					// Now the rule to determine the cardinality.
					var fieldType string
					fieldType = this.m_id + "." + schemasName + sourceTableName

					if this.isRef(refName) {
						fieldType = fieldType + ":Ref"
					}

					if !sourceField_isId {
						// one to one relationship in that case.
						fieldType = "[]" + fieldType
					} else if len(src.Ids) > len(trg.Ids) {
						fieldType = "[]" + fieldType
					}

					appendField(trg, "M_"+refName, fieldType)
				}

			} else {
				associativeTables[src.TypeName] = src
				associations[src.TypeName] = append(associations[src.TypeName], make([]interface{}, 2))
				associations[src.TypeName][len(associations[src.TypeName])-1][0] = refName
				associations[src.TypeName][len(associations[src.TypeName])-1][1] = trg
			}
		}
	}

	// Now I will set the associations *Cross references
	for _, associativeTable := range associativeTables {
		associations_ := associations[associativeTable.TypeName]
		for i := 0; i < len(associations_); i++ {
			refName := associations_[i][0].(string)
			trg, _ := GetServer().GetEntityManager().getEntityPrototype(associations_[i][1].(*EntityPrototype).TypeName, "sql_info")
			if !Utility.Contains(associativeTable.Fields, "M_"+refName) {
				fieldType := trg.TypeName
				if this.isRef(refName) {
					fieldType = fieldType + ":Ref"
				}
				if !this.isAssociative(trg) {
					appendField(associativeTable, "M_"+refName, fieldType)
				}
			}

			// Associative table must contain tow value.
			for j := 0; j < len(associations_); j++ {
				if j != i {
					trg1, err := GetServer().GetEntityManager().getEntityPrototype(associations_[j][1].(*EntityPrototype).TypeName, "sql_info")
					if err != nil {
						return err
					}
					if !this.isAssociative(trg) {
						appendField(trg1, "M_"+refName, "[]"+trg.TypeName+":Ref")
					}
					if !this.isAssociative(trg1) {
						appendField(trg, "M_"+associations_[j][0].(string), "[]"+trg1.TypeName+":Ref")
					}
				}
			}
		}
	}

	return nil
}

/**
 * Get the mapping of a given sql type.
 */
func (this *SqlDataStore) getSqlTypePrototype(typeName string) (*EntityPrototype, error) {
	if typeName == "enum" {
		// enum is not a basic sql type...
		typeName = "varchar"
	}
	prototype, err := GetServer().GetEntityManager().getEntityPrototype("sqltypes."+typeName, "sqltypes")
	return prototype, err
}

/**
 * Remove a given entity prototype.
 */
func (this *SqlDataStore) DeleteEntityPrototype(id string) error {

	store := GetServer().GetDataManager().getDataStore("sql_info")
	if store != nil {
		return store.DeleteEntityPrototype(id)
	}
	return errors.New("Store with id sql_info dosen't exist!")
}

/**
 * Remove all prototypes.
 */
func (this *SqlDataStore) DeleteEntityPrototypes() error {
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
 * synchronize the content of database with sql_info content. Only key's will be
 * save, the other field will be retreive as needed via sql querie's.
 */
func (this *SqlDataStore) synchronize(prototypes []*EntityPrototype) error {

	// Keep entities before save...
	toSave := make(map[string]Entity, 0)
	entitiesByType := make(map[string][]Entity, 0)

	// First of all I will sychronize create the enities information if it dosen't exist.
	for i := 0; i < len(prototypes); i++ {
		prototype, _ := GetServer().GetEntityManager().getEntityPrototype(prototypes[i].TypeName, "sql_info")
		if !this.isAssociative(prototype) { // Associative table object are not needed...
			if len(prototype.Ids) > 1 {
				query := "SELECT "
				fieldsType := make([]interface{}, 0)
				for j := 0; j < len(prototype.Ids); j++ {
					if strings.HasPrefix(prototype.Ids[j], "M_") {
						query += strings.Replace(prototype.Ids[j], "M_", "", -1)
						fieldsType = append(fieldsType, prototype.FieldsType[prototype.getFieldIndex(prototype.Ids[j])])
						if j < len(prototype.Ids)-1 {
							query += " ,"
						}
					}
				}
				query += " FROM " + prototype.TypeName
				var params []interface{}

				// Execute the query...
				values, err := this.Read(query, fieldsType, params)
				if err != nil {
					return err
				}

				// Now I will generate a unique key for the retreive information.
				for j := 0; j < len(values); j++ {
					//log.Println(prototype.TypeName, j, ":", len(values))
					keyInfo := prototype.TypeName + ":"
					for k := 0; k < len(values[j]); k++ {
						keyInfo += Utility.ToString(values[j][k])
						// Append underscore for readability in case of problem...
						if k < len(values[j])-1 {
							keyInfo += "_"
						}
					}

					// The uuid is in that case a MD5 value.
					uuid := prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)

					// Now I will create the entity if it dosen't exist.
					_, errObj := GetServer().GetEntityManager().getDynamicEntityByUuid(uuid)
					if errObj != nil {
						// Here I will create the Dynamic entity.
						infos := make(map[string]interface{}, 0)
						infos["TYPENAME"] = prototype.TypeName
						infos["UUID"] = uuid

						// The 0 value is the uuid
						for k := 1; k < len(prototype.Ids); k++ {
							id := prototype.Ids[k]
							infos[id] = values[j][k-1]
						}

						entity, errObj := GetServer().GetEntityManager().newDynamicEntity("", infos)
						if errObj == nil {
							// Save the entity.
							toSave[uuid] = entity
							if entitiesByType[entity.GetTypeName()] == nil {
								entitiesByType[entity.GetTypeName()] = make([]Entity, 0)
							}
							entitiesByType[entity.GetTypeName()] = append(entitiesByType[entity.GetTypeName()], entity)
						} else {
							log.Println("------> error", errObj)
						}
					}
				}
			}
		}
	}

	// Now I will set the entity reliationships
	for i := 0; i < len(prototypes); i++ {
		if !this.isAssociative(prototypes[i]) { // Associative table object are not needed...
			prototype, _ := GetServer().GetEntityManager().getEntityPrototype(prototypes[i].TypeName, "sql_info")
			for j := 0; j < len(prototype.FieldsType); j++ {
				entities := entitiesByType[prototype.TypeName]
				for k := 0; k < len(entities); k++ {
					entity := entities[k]
					log.Println("1428 ----> val ", k, "/", len(entities))
					if !strings.HasPrefix(prototype.FieldsType[j], "sqltypes") && !strings.HasPrefix(prototype.FieldsType[j], "[]sqltypes") && strings.HasPrefix(prototype.Fields[j], "M_") {

						// If the relation is aggregation or association.
						isRef := strings.HasSuffix(prototype.FieldsType[j], ":Ref")
						// This is an aggregation releationship.
						refInfos, err := this.getRefInfos(prototype.Fields[j][2:])
						var childTypeName string
						if this.m_vendor == Config.DataStoreVendor_ODBC {
							childTypeName = this.m_id + "." + refInfos[4] + "." + refInfos[0]
						} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
							childTypeName = this.m_id + "." + refInfos[0]
						}

						if err == nil && !this.m_associations[childTypeName] {
							id := entity.GetObject().(map[string]interface{})["M_"+refInfos[3]]
							query := "SELECT " + refInfos[3] + " FROM " + refInfos[0] + " WHERE " + refInfos[1] + "=" + Utility.ToString(id)
							var params []interface{}
							var fieldsType []interface{}
							fieldsType = append(fieldsType, prototype.FieldsType[prototype.getFieldIndex("M_"+refInfos[3])])

							// Execute the query...
							values, err := this.Read(query, fieldsType, params)
							if err == nil {
								for n := 0; n < len(values); n++ {
									log.Println("1468 val n ", n, "/", len(values))
									// Set the child uuid.
									childUuid := childTypeName + "%" + Utility.GenerateUUID(entity.GetUuid()+":"+childTypeName+":"+Utility.ToString(values[n][0]))
									var childEntity Entity
									if toSave[childUuid] != nil {
										childEntity = toSave[childUuid]
									} else {
										childEntity, _ = GetServer().GetEntityManager().getEntityByUuid(childUuid)
									}

									if childEntity != nil {
										if !isRef {
											entity.AppendChild(prototype.Fields[j], childEntity)
											// The parent will save the child...
											delete(toSave, childUuid)
											log.Println("-------> child uuid ", childEntity.GetUuid(), " found!!!!")
										} else {
											if !strings.HasPrefix(prototype.FieldsType[j], "[]") {
												entity.(*DynamicEntity).setValue(prototype.Fields[j], childEntity)
											} else {
												entity.(*DynamicEntity).appendValue(prototype.Fields[j], childEntity)
											}
											entity.AppendReference(childEntity)
											childEntity.AppendReferenced(prototype.Fields[j], entity)
											log.Println("-------> ref uuid ", childEntity.GetUuid(), " found!!!!")
										}
									} else {
										log.Println("-------> child uuid ", childUuid, " not found!!!!")
									}
								}
							}

						}
					}
				}
			}
		}
	}

	// Save changed entity.
	for uuid, entity := range toSave {
		log.Println("----> entity ", uuid, " was save successfully!")
		entity.SaveEntity()
	}

	log.Println("------------> end of sync!")
	return nil
}
