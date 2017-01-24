package Server

import (
	//"log"

	"code.myceliUs.com/CargoWebServer/Cargo/Persistence/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
)

/**
 * The logger construnctor..
 */
func NewLogger(id string) *Logger {
	logger := new(Logger)
	logger.id = id

	// I will get it related log with the same id from the entity store.
	uuid := CargoEntitiesLogExists(id)

	if len(uuid) > 0 {
		entity, _ := GetServer().GetEntityManager().getEntityByUuid(uuid)
		logger.logEntity = entity.(*CargoEntities_LogEntity)
	} else {
		logger.logEntity = GetServer().GetEntityManager().NewCargoEntitiesLogEntity(id, nil)
		logger.logEntity.GetObject().(*CargoEntities.Log).SetId(id)

		server.GetEntityManager().getCargoEntities().GetObject().(*CargoEntities.Entities).SetEntities(logger.logEntity.GetObject())
		server.GetEntityManager().getCargoEntities().SaveEntity()
	}

	return logger
}

/**
 * A logger object.
 */
type Logger struct {
	// The id for the group of entries.
	id string
	// The log object....
	logEntity *CargoEntities_LogEntity
}

/**
 * Create the entity...
 */
func (this *Logger) AppendLogEntry(toLog CargoEntities.Entity) {
	// Set the log entry informations.
	logEntry := new(CargoEntities.LogEntry)
	logEntry.SetLoggerPtr(this.logEntity.GetObject())
	logEntry.SetEntityRef(toLog)
	logEntry.SetCreationTime(Utility.MakeTimestamp())

	// Append the entry in the logger.
	this.logEntity.GetObject().(*CargoEntities.Log).SetEntries(logEntry)

	// Save the logEntry.
	this.logEntity.SaveEntity()
}
