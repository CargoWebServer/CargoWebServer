// +build CargoEntities

package CargoEntities

type AccessType int
const(
	AccessType_Hidden AccessType = 1+iota
	AccessType_Public
	AccessType_Restricted
)
