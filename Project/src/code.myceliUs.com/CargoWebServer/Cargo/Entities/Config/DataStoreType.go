// +build Config

package Config

type DataStoreType int
const(
	DataStoreType_SQL_STORE DataStoreType = 1+iota
	DataStoreType_KEY_VALUE_STORE
)
