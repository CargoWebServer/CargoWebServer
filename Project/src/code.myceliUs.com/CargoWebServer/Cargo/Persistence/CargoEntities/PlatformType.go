package CargoEntities

type PlatformType int
const(
	PlatformType_Unknown PlatformType = 1+iota
	PlatformType_Tablet
	PlatformType_Phone
	PlatformType_Desktop
	PlatformType_Laptop
)
