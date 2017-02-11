package CargoEntities

type OsType int
const(
	OsType_Unknown OsType = 1+iota
	OsType_Linux
	OsType_Windows7
	OsType_Windows8
	OsType_Windows10
	OsType_OSX
	OsType_IOS
)
