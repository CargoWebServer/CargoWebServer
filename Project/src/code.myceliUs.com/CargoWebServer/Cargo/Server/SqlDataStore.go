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
	/** The id of the store... **/
	m_id string

	/** The name of the store **/
	m_name string

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

	/** The list of reference informations **/
	m_refInfos map[string][][]string

	/** Keep prototypes here **/
	m_prototypes map[string]*EntityPrototype
}

func NewSqlDataStore(info *Config.DataStoreConfiguration) (*SqlDataStore, error) {

	// The store...
	store := new(SqlDataStore)
	store.m_id = info.M_id

	// Keep this info...
	store.m_name = info.M_storeName
	store.m_vendor = info.M_dataStoreVendor
	store.m_host = info.M_hostName
	store.m_user = info.M_user
	store.m_password = info.M_pwd
	store.m_port = info.M_port
	store.m_textEncoding = info.M_textEncoding

	// Keep the list of associations...
	store.m_associations = make(map[string]bool, 0)
	store.m_refInfos = make(map[string][][]string, 0)

	if store.m_vendor == Config.DataStoreVendor_ARANGODB {
		return nil, errors.New("ArgangoDB is store not a sql store.")
	}

	store.m_prototypes = make(map[string]*EntityPrototype, 0)

	return store, nil
}

func (this *SqlDataStore) Connect() error {
	//log.Println("-----> connect: ", this.GetId())
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
		connectionString += "database=" + this.m_name + ";"
		connectionString += "driver=mssql"
		//connectionString += "encrypt=false;"
		driver = "mssql"
	} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
		/** Connect to oracle MySql server here... **/
		connectionString += this.m_user + ":"
		connectionString += this.m_password + "@tcp("
		connectionString += this.m_host + ":" + strconv.Itoa(this.m_port) + ")"
		connectionString += "/" + this.m_name
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
		connectionString += "database=" + this.m_name + ";"
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
	if err != nil {
		return err
	}

	// Try to ping the sql server to be sure the connection is open.
	err = this.Ping()
	if err != nil {
		return err
	}

	// set the prototype in the map...
	this.GetEntityPrototypes()

	return err
}

func (this *SqlDataStore) GetId() string {
	return this.m_id
}

/** Open a new Connection with the data store **/
func (this *SqlDataStore) Ping() (err error) {
	if this.m_db == nil {
		err = errors.New("No connection was found for datastore " + this.m_name)
		//log.Println(err)
		return err
	}

	errChannel := make(chan error, 0)
	timer := time.NewTimer(5 * time.Second)

	// Set timer to 5 second before give up on connection.
	go func(errChannel chan error, dataStore DataStore, timer *time.Timer) {
		<-timer.C
		dataStore.Close() // Close the connection.
		err := errors.New("Ping " + dataStore.GetId() + " failed!")
		errChannel <- err
	}(errChannel, this, timer)

	// Try to ping the connection.
	go func(errChannel chan error, dataStore DataStore, timer *time.Timer) {
		errChannel <- dataStore.(*SqlDataStore).m_db.Ping()
		timer.Stop()
	}(errChannel, this, timer)

	err = <-errChannel

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
		return -1, err
	}
	_, err = stmt.Exec(data_...)
	if err != nil {
		return -1, err
	}

	// First of all I need to find the key column id...
	// From the query I will retreive the table name.
	// the query follow this pattern... INSERT INTO TableNameHere(...
	startIndex := strings.Index(strings.ToUpper(query), " INTO ") + 5 // 5 is the len of INTO and one space...
	endIndex := strings.Index(query, "(")
	tableName := strings.TrimSpace(query[startIndex:endIndex])
	if strings.LastIndex(tableName, ".") != -1 {
		tableName = tableName[strings.LastIndex(tableName, ".")+1:]
	}

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
	results, err := GetServer().GetDataManager().readData(this.GetId(), colKeyQuery, fieldType, params)
	// Work for one column...
	if err == nil {
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
				results, err = GetServer().GetDataManager().readData(this.GetId(), lastIndexQuery, fieldType, params)
				if err == nil {
					if len(results) > 0 {
						lastId = int64(results[0][0].(int))
					}
				} else {
					log.Println("-------> fail to execute query ", lastIndexQuery, err)
				}
			}
		}
	} else {
		log.Println("-------> fail to execute query ", colKeyQuery, " table Name ", tableName, err)
	}

	if err == nil {
		eventData := make([]*MessageData, 3)

		tableName_ := new(MessageData)
		tableName_.TYPENAME = "Server.MessageData"
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
		evt, _ := NewEvent(NewRowEvent, DataEvent, eventData)
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
		} else if reflect.TypeOf(value).String() == "int64" {
			if value.(int64) == 1 {
				return true
			} else {
				return false
			}
		} else {
			log.Println("-----------> bit cast fail! ", reflect.TypeOf(value).String())
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
			var val string
			val = value.(time.Time).Format(createdFormat)
			return val
		} else if reflect.TypeOf(value).String() == "[]uint8" {
			t, err := time.Parse(time.RFC3339Nano, string(value.([]uint8)))
			if err == nil {
				return t
			}
		}
	}

	/////////////////////////// string ////////////////////////////////
	if strings.HasSuffix(sqlTypeName, "string") || strings.HasSuffix(sqlTypeName, "nchar") || strings.HasSuffix(sqlTypeName, "varchar") || strings.HasSuffix(sqlTypeName, "nvarchar") || strings.HasSuffix(sqlTypeName, "text") || strings.HasSuffix(sqlTypeName, "ntext") {
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
 * Return true if a value is a foreign key false instead.
 */
func isForeignKey(val string) bool {
	if strings.Index(strings.ToUpper(val), "_FK_") != -1 {
		return true
	}

	if strings.HasPrefix(strings.ToUpper(val), "FK_") {
		return true
	}

	return false
}

/**
 * Read a query execute it and return the result as an array of interface...
 */
func (this *SqlDataStore) Read(query string, fieldsType []interface{}, params []interface{}) ([][]interface{}, error) {
	err := this.Ping()
	if err != nil {
		err = this.Connect()
		if err != nil {
			return nil, err
		}
	}
	rows, err := this.m_db.Query(query, params...)
	if err != nil {
		log.Println("---> sql read query error:", err)
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
		return err
	}

	// Close after use
	defer stmt.Close()

	var values []interface{}
	values = append(values, fields...)
	values = append(values, params...)

	_, err = stmt.Exec(values...)

	if err != nil {
		return err
	}

	startIndex := strings.Index(strings.ToUpper(query), "UPDATE") + 6 // 5 is the len of INTO and one space...
	endIndex := strings.Index(strings.ToUpper(query), "SET")
	tableName := strings.TrimSpace(query[startIndex:endIndex])
	tableName = strings.Replace(tableName, "[", "", -1)
	tableName = strings.Replace(tableName, "]", "", -1)

	if err == nil {
		eventData := make([]*MessageData, 2)

		tableName_ := new(MessageData)
		tableName_.TYPENAME = "Server.MessageData"
		tableName_.Name = "tableName"
		tableName_.Value = tableName
		eventData[0] = tableName_

		id := new(MessageData)
		id.TYPENAME = "Server.MessageData"
		id.Name = "id"
		id.Value = Utility.ToString(params[0])
		eventData[1] = id

		evt, _ := NewEvent(UpdateRowEvent, DataEvent, eventData)
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
		return err
	}

	// Close after use
	defer stmt.Close()

	_, err = stmt.Exec(params...)
	if err != nil {
		//log.Println(err)
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
		evt, _ := NewEvent(DeleteRowEvent, DataEvent, eventData)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}

	return err
}

/**
 * Close the backend store.
 */
func (this *SqlDataStore) Close() error {
	if this.m_db != nil {
		return this.m_db.Close()
	}

	// Clear the map.
	for typeName, _ := range this.m_prototypes {
		delete(this.m_prototypes, typeName)
	}

	return nil
}

/**
 * Create a new entity prototype.
 */
func (this *SqlDataStore) CreateEntityPrototype(prototype *EntityPrototype) error {
	/** Nothing to do here prototype are generated each time **/
	return nil
}

/**
 * Save entity prototype.
 */
func (this *SqlDataStore) SaveEntityPrototype(prototype *EntityPrototype) error {
	/** Nothing to do here, prototype are generated each time **/
	return nil
}

/**
 * Return the prototypes for all accessible table of a database.
 */
func (this *SqlDataStore) GetEntityPrototypes() ([]*EntityPrototype, error) {
	var prototypes []*EntityPrototype
	var query string

	fieldsType := make([]interface{}, 0)
	// So here I will get the list of table name that will be prototypes...
	if this.m_vendor == Config.DataStoreVendor_ODBC {
		query = "SELECT sobjects.name "
		query += "FROM sysobjects sobjects "
		query += "WHERE sobjects.xtype = 'U'"

	} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
		query = "SELECT table_name FROM information_schema.tables where table_schema='" + this.m_name + "'"
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
		if err == nil {
			prototypes = append(prototypes, prototype)
		}
	}

	// Complete the reference information.
	this.setRefs()

	return prototypes, nil
}

/**
 * Return the prototype of a given table.
 */
func (this *SqlDataStore) GetEntityPrototype(id string) (*EntityPrototype, error) {

	if !Utility.IsValidVariableName(id) {
		return nil, errors.New("Wrong type name " + id + "!")
	}

	if this.m_prototypes[id] != nil {
		return this.m_prototypes[id], nil
	}

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

	// Here remove space and dash symbol...
	if strings.Index(id, " ") > -1 || strings.Index(id, " ") > -1 {
		return nil, errors.New("Id \"" + id + "\" must not contain space or dash")
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
			if !Utility.Contains(prototype.Fields, "M_"+fieldName) && Utility.IsValidVariableName(fieldName) {
				prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
				prototype.Fields = append(prototype.Fields, "M_"+fieldName)
				prototype.FieldsType = append(prototype.FieldsType, fieldType.TypeName)
				prototype.FieldsDocumentation = append(prototype.FieldsDocumentation, "")
				setDefaultFieldValue(prototype, fieldTypeName)

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

	// keep in the map.
	this.m_prototypes[id] = prototype

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
				return this.m_name + "." + values[0][0].(string), nil
			}
		}
	} else if this.m_vendor == Config.DataStoreVendor_MYSQL {
		// with mysql the schema id is the id of the database.
		return this.m_name, nil
	}

	return "", errors.New("No schema found for table " + name)

}

func appendField(prototype *EntityPrototype, fieldName string, fieldType string) {

	if !Utility.Contains(prototype.Fields, fieldName) && Utility.IsValidVariableName(fieldName) {
		prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
		prototype.FieldsNillable = append(prototype.FieldsVisibility, true)
		prototype.Fields = append(prototype.Fields, fieldName)
		prototype.FieldsType = append(prototype.FieldsType, fieldType)
		prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
		prototype.FieldsDocumentation = append(prototype.FieldsDocumentation, "")

		// Now default value.
		setDefaultFieldValue(prototype, fieldName)
	}
}

// Return true if the prototype is an associative one...
func (this *SqlDataStore) isAssociative(prototype *EntityPrototype) bool {
	// if the asscociation already exist.
	if val, ok := this.m_associations[prototype.TypeName]; ok {
		return val
	}

	if len(prototype.Ids) == 2 {
		// Here the only value in table is a key...
		return false
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
func (this *SqlDataStore) getRefInfos(refId string) ([][]string, error) {

	// Return the previous found result here.
	if this.m_refInfos[refId] != nil {
		return this.m_refInfos[refId], nil
	}

	var results [][]string

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
		return results, err
	} else if len(values) > 0 {
		results = make([][]string, 0)
		for i := 0; i < len(values); i++ {
			if len(values[i]) == 5 {
				result := make([]string, 5)
				// Set the value inside the results.
				result[0] = values[i][0].(string)
				result[1] = values[i][1].(string)
				result[2] = values[i][2].(string)
				result[3] = values[i][3].(string)
				result[4] = values[i][4].(string)
				results = append(results, result)
			}
		}
		return results, nil
	}

	// Keep the results in memory for future use.
	this.m_refInfos[refId] = results

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
		query += "	and table_schema = '" + this.m_name + "'"
	}

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

			src, err := GetServer().GetEntityManager().getEntityPrototype(sourceTableName, this.m_name)
			if err != nil {
				log.Println("----------------> src not found: ", err)
				return err
			}

			trg, err := GetServer().GetEntityManager().getEntityPrototype(targetTableName, this.m_name)
			if err != nil {
				log.Println("----------------> trg not found: ", err)
				return err
			}

			// Test if source field and target field are id's
			sourceField_isId := Utility.Contains(src.Ids, "M_"+values[i][3].(string))
			targetField_isId := Utility.Contains(trg.Ids, "M_"+values[i][5].(string))

			// I will append the field if is not already there.
			if !this.isAssociative(src) {
				// Now the sources.
				if !Utility.Contains(src.Fields, "M_"+refName) {

					// Now the rule to determine the cardinality.
					var fieldType string
					fieldType = this.m_name + "." + schemasName + targetTableName + ":Ref"

					if !targetField_isId {
						// one to one relationship in that case.
						fieldType = "[]" + fieldType
					}

					appendField(src, "M_"+refName, fieldType)
				}

				if !Utility.Contains(trg.Fields, "M_"+refName) {
					// Now the rule to determine the cardinality.
					var fieldType string
					fieldType = this.m_name + "." + schemasName + sourceTableName

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
			trg, _ := GetServer().GetEntityManager().getEntityPrototype(associations_[i][1].(*EntityPrototype).TypeName, this.m_id)
			if !Utility.Contains(associativeTable.Fields, "M_"+refName) {
				fieldType := trg.TypeName
				fieldType = fieldType + ":Ref" // Associative table contain only references.
				if !this.isAssociative(trg) {
					appendField(associativeTable, "M_"+refName, fieldType)
				}
			}

			// Associative table must contain tow value.
			for j := 0; j < len(associations_); j++ {
				if j != i {
					trg1, err := GetServer().GetEntityManager().getEntityPrototype(associations_[j][1].(*EntityPrototype).TypeName, this.m_id)
					if err != nil {
						return err
					}
					if !this.isAssociative(trg) {
						appendField(trg, "M_"+refName, "[]"+associativeTable.TypeName+":Ref")
					}
					if !this.isAssociative(trg1) {
						appendField(trg1, "M_"+associations_[j][0].(string), "[]"+associativeTable.TypeName+":Ref")
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

	return nil /** prototype are generated object not saved **/
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

// local function to genereate the entity uuid...
// TODO handle multiple id key...
func generateUuid(key string, info map[string]interface{}, infos map[string]map[string]map[string]interface{}) string {
	keyInfo := key
	if info["parentId"] != nil {
		// here a parent exist for the entity so I will get it.
		parentId := info["parentId"].(string)
		parentType := parentId[0:strings.Index(parentId, ":")]
		parentInfo := infos[parentType][parentId]
		info["ParentUuid"] = generateUuid(parentId, parentInfo, infos)
		keyInfo = info["ParentUuid"].(string) + ":" + keyInfo
	}
	uuid := info["TYPENAME"].(string) + "%" + Utility.GenerateUUID(keyInfo)
	info["UUID"] = uuid // Keep the uuid.
	return uuid
}

// Generate the entity if is not already exist.
func createEntityFromInfo(key string, info map[string]interface{}, infos map[string]map[string]map[string]interface{}) Entity {
	var parentUuid string
	if info["ParentUuid"] != nil {
		parentUuid = info["ParentUuid"].(string)
		if strings.Index(parentUuid, "%") != -1 {
			// Here I will generate the parent entity...
			parentTypeName := parentUuid[0:strings.Index(parentUuid, "%")]
			// Here I will retreive the parent information.
			parentInfo := infos[parentTypeName][info["parentId"].(string)]

			/* Create the parent entity first... */
			createEntityFromInfo(info["parentId"].(string), parentInfo, infos)
		}
	}

	entity := NewDynamicEntity()
	entity.SetParentUuid(parentUuid)
	entity.setObject(info)

	return entity
}

/**
 * synchronize the content of database content. Only key's will be
 * save, the other field will be retreive as needed via sql querie's.
 */
func (this *SqlDataStore) synchronize(prototypes []*EntityPrototype) error {

	// Contain entity info without parent uuid.
	entityInfos := make(map[string]map[string]map[string]interface{}, 0)
	// First of all I will sychronize create the entities information if it dosen't exist.
	for i := 0; i < len(prototypes); i++ {
		prototype, _ := GetServer().GetEntityManager().getEntityPrototype(prototypes[i].TypeName, this.m_id)
		// Associative table object are not needed...
		if len(prototype.Ids) > 1 {
			query := "SELECT "
			fieldsType := make([]interface{}, 0)
			fieldsName := make([]string, 0)
			for j := 0; j < len(prototype.Ids); j++ {
				if strings.HasPrefix(prototype.Ids[j], "M_") {
					query += strings.Replace(prototype.Ids[j], "M_", "", -1)
					fieldsType = append(fieldsType, prototype.FieldsType[prototype.getFieldIndex(prototype.Ids[j])])
					fieldsName = append(fieldsName, prototype.Ids[j])
					if j < len(prototype.Ids)-1 {
						query += " ,"
					}
				}
			}

			// I will also append field's that are parts of relationship
			for j := 0; j < len(prototype.Fields); j++ {
				field := prototype.Fields[j]
				// Not an id...
				refInfos, err := this.getRefInfos(field[2:])
				if err == nil {
					// here a reference exist for that field so I will
					// get it from sql.
					for k := 0; k < len(refInfos); k++ {
						if strings.HasSuffix(prototype.TypeName, refInfos[k][0]) {
							if !Utility.Contains(fieldsName, "M_"+refInfos[k][1]) {
								query += " ,"
								query += refInfos[k][1]
								fieldType := prototype.FieldsType[prototype.getFieldIndex("M_"+refInfos[k][1])]
								fieldsType = append(fieldsType, fieldType)
								fieldsName = append(fieldsName, "M_"+refInfos[k][1])
							}
						}
					}
				}
			}

			// append index fields
			for j := 0; j < len(prototype.Fields); j++ {
				field := prototype.Fields[j]
				// Not an id...
				if !Utility.Contains(fieldsName, field) && strings.HasPrefix(field, "M_") && !isForeignKey(field) {
					query += " ,"
					query += field[2:]
					fieldType := prototype.FieldsType[prototype.getFieldIndex(field)]
					fieldsType = append(fieldsType, fieldType)
					fieldsName = append(fieldsName, field)
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
				// Here I will keep entity id's information.
				// UUID need ParentUuid before it can be generate so a second
				// pass will be necessary.
				entityInfo := make(map[string]interface{}, 0)
				entityInfo["TYPENAME"] = prototype.TypeName

				keyInfo := prototype.TypeName + ":"
				// -1 the first ids is the uuid and we dont have it now.
				for k := 0; k < len(prototype.Ids)-1; k++ {
					keyInfo += Utility.ToString(values[j][k])
					// Append underscore for readability in case of problem...
					if k < len(prototype.Ids)-2 {
						keyInfo += "_"
					}
				}

				for k := 0; k < len(values[j]); k++ {
					// Keep the id value in the entity field.
					if values[j][k] != nil {
						entityInfo[fieldsName[k]] = values[j][k] // also save entity values in the entity.
					}
				}

				// Keep the info.
				if entityInfos[prototype.TypeName] == nil {
					entityInfos[prototype.TypeName] = make(map[string]map[string]interface{}, 0)
				}

				// append the entity
				entityInfos[prototype.TypeName][keyInfo] = entityInfo
			}
		}
	}

	// Set the parent realationship here.
	for i := 0; i < len(prototypes); i++ {
		if !this.isAssociative(prototypes[i]) { // Associative table have no parent...
			prototype, _ := GetServer().GetEntityManager().getEntityPrototype(prototypes[i].TypeName, this.m_id)
			for _, info := range entityInfos[prototypes[i].TypeName] {
				for j := 0; j < len(prototype.FieldsType); j++ {
					if !strings.HasPrefix(prototype.FieldsType[j], "sqltypes") && strings.HasSuffix(prototype.FieldsType[j], ":Ref") && !strings.HasPrefix(prototype.FieldsType[j], "[]sqltypes") && strings.HasPrefix(prototype.Fields[j], "M_") {
						refInfos, err := this.getRefInfos(prototype.Fields[j][2:])
						if err == nil {
							ids := make([]string, 0)
							// Here If the other side or the relation is not a reference that's mean
							// it is he's parent.
							prototype_, _ := GetServer().GetEntityManager().getEntityPrototype(strings.Replace(strings.Replace(prototype.FieldsType[j], "[]", "", -1), ":Ref", "", -1), this.m_id)
							index := prototype_.getFieldIndex(prototype.Fields[j])
							for k := 0; k < len(refInfos); k++ {
								if info["M_"+refInfos[k][1]] != nil {
									if index > -1 {
										if !strings.HasSuffix(prototype_.FieldsType[index], ":Ref") {
											// So here I will create the parent id string to retreive it later.
											ids = append(ids, Utility.ToString(info["M_"+refInfos[k][1]]))
										}
									}
								}
							}
							// Now I will create the parent id from found ids.
							if len(ids) > 0 {
								parentId := prototype_.TypeName + ":"
								for k := 0; k < len(ids); k++ {
									parentId += ids[k]
									if k < len(ids)-1 {
										parentId += "_"
									}
								}
								info["parentId"] = parentId
							}
						}
					}
				}
			}
		}
	}

	// Release memory.
	entityInfos = make(map[string]map[string]map[string]interface{}, 0)

	// Note that relationship are not part of the object data itself but
	// are initialysed at runtime by the enity manager each time an entity
	// backed by sql is use.

	return nil
}
