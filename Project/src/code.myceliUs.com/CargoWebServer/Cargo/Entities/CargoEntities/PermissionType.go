// +build CargoEntities

package CargoEntities

type PermissionType int
const(
	PermissionType_Create PermissionType = 1+iota
	PermissionType_Read
	PermissionType_Update
	PermissionType_Delete
)
