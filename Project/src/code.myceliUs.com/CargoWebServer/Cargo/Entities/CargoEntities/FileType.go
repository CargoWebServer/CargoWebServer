// +build CargoEntities

package CargoEntities

type FileType int
const(
	FileType_DbFile FileType = 1+iota
	FileType_DiskFile
)
