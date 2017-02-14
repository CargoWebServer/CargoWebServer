// +build Config

package Config

type DataStoreType int
const(
	DataStoreType_SQL_STORE DataStoreType = 1+iota
	DataStoreType_MEMORY_STORE
	DataStoreType_KEY_VALUE_STORE
	DataStoreType_GRAPH_STORE
)
