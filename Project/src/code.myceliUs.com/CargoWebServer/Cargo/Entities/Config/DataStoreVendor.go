// +build Config

package Config

type DataStoreVendor int
const(
	DataStoreVendor_CAYLEY DataStoreVendor = 1+iota
	DataStoreVendor_MYSQL
	DataStoreVendor_MSSQL
	DataStoreVendor_ODBC
)
