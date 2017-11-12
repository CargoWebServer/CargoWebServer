package Server

import (
	//	"log"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	//	"code.myceliUs.com/Utility"
)

/**
 * The logger construnctor..
 */
func NewLogger(id string) *Logger {
	logger := new(Logger)
	logger.id = id

	return logger
}

/**
 * A logger object.
 */
type Logger struct {
	// The id for the group of entries.
	id string
}

/**
 * Create the entity...
 */
func (this *Logger) AppendLogEntry(toLog CargoEntities.Entity) {

}
