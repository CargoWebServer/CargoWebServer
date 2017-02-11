package Config

type DataStoreVendor int
const(
	DataStoreVendor_MYCELIUS DataStoreVendor = 1+iota
	DataStoreVendor_MYSQL
	DataStoreVendor_MSSQL
	DataStoreVendor_ODBC
	DataStoreVendor_KNOWLEDGEBASE
)
