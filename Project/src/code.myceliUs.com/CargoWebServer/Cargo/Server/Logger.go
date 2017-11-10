package Server

import (
	"log"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
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
		entity, _ := GetServer().GetEntityManager().getEntityByUuid(uuid, false)
		logger.logEntity = entity.(*CargoEntities_LogEntity)
	} else {
		logObject := new(CargoEntities.Log)
		logObject.SetId(id)
		entities := server.GetEntityManager().getCargoEntities()

		// Set the log entity.
		logEntity, err := GetServer().GetEntityManager().createEntity(entities.GetUuid(), "M_entities", "CargoEntities.Log", id, logObject)
		if err == nil {
			logger.logEntity = logEntity.(*CargoEntities_LogEntity)
		} else {
			log.Panicln("-----> error ", err)
		}
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
