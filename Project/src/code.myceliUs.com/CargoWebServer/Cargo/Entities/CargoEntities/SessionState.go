// +build CargoEntities

package CargoEntities

type SessionState int

const (
	SessionState_Online SessionState = 1 + iota
	SessionState_Away
	SessionState_Offline
)
