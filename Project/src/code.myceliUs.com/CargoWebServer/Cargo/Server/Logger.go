package Server

import (
	"fmt"
	"log"
	"os"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

// Shortcut function.
func LogInfo(infos ...interface{}) {

	// also display in the command prompt.
	log.Println(infos...)
	GetServer().getDefaultInfoLogger().LogInfo(fmt.Sprintln(infos))
}

func LogError(infos ...interface{}) {
	var info string
	for i := 0; i < len(infos); i++ {
		info += Utility.ToString(infos[i])
		if i < len(infos)-1 {
			info += "Error"
		}
	}

	// also display in the command prompt.
	log.Println(info)
	GetServer().getDefaultErrorLogger().LogInfo(info)
}

/**
 * The logger construnctor..
 */
func NewLogger(id string) *Logger {
	logger := new(Logger)
	logger.id = id
	logger.logChannel = make(chan string)

	// The log processing information.
	go func() {
		for {
			select {
			case msg := <-logger.logChannel:
				// Open the log file.
				f, err := os.OpenFile("text.log",
					os.O_APPEND|os.O_CREATE|os.O_WRONLY, 777)
				if err != nil {
					log.Println(err)
				}

				// set the message.
				logger := log.New(f, "", log.LstdFlags)
				logger.Println(msg)

				f.Close()
			}
		}
	}()

	return logger
}

/**
 * A logger object.
 */
type Logger struct {
	// The id for the group of entries.
	id string

	// The channel that receive message.
	logChannel chan string
}

/**
 * Create the entity...
 */
func (this *Logger) AppendLogEntry(toLog CargoEntities.Entity) {

}

/**
 * Log information to file.
 */
func (self *Logger) LogInfo(info string) {
	// set the message on the channel.
	self.logChannel <- info
}
