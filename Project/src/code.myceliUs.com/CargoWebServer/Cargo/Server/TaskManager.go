/**
 * Internal task manager.
 */
package Server

import (
	b64 "encoding/base64"
	"time"

	"errors"
	"log"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/CargoWebServer/Cargo/JS"
	"code.myceliUs.com/Utility"
)

type TaskInstanceInfo struct {
	TYPENAME     string
	UUID         string
	TaskId       string
	CreationTime int64
	StartTime    int64
	EndTime      int64
	CancelTime   int64
	Error        string
	Status       string
}

/**
 * That class contain the logic to process JS task. The interface to external
 * user are in the class configuration manager.
 */
type TaskManager struct {
	// synchronize task...
	m_setScheduledTasksChan    chan *Config.ScheduledTask
	m_cancelScheduledTasksChan chan string
	m_getTaskInstancesInfo     chan chan []*TaskInstanceInfo
}

var taskManager *TaskManager

func newTaskManager() *TaskManager {

	// Open the channel.
	taskManager = new(TaskManager)
	taskManager.m_setScheduledTasksChan = make(chan *Config.ScheduledTask)
	taskManager.m_cancelScheduledTasksChan = make(chan string)
	taskManager.m_getTaskInstancesInfo = make(chan chan []*TaskInstanceInfo)

	// start the task manager.
	go taskManager.start()

	return taskManager
}

/**
 * Task processing function.
 */
func (this *TaskManager) start() {

	// Keep timers by task id to be able to cancel it latter.
	tasks := make(map[string]*time.Timer, 0)
	// Keep info of task instances since the server started.
	instancesInfos := make([]*TaskInstanceInfo, 0)

	for {
		select {
		// set the task.
		case task := <-this.m_setScheduledTasksChan:

			// Only one instance at time can running at any time...
			for i := 0; i < len(instancesInfos); i++ {
				if task.GetId() == instancesInfos[i].TaskId {
					if tasks[instancesInfos[i].TaskId] != nil {
						if instancesInfos[i].Status != "Completed" && instancesInfos[i].Status != "Failed" {
							// Stop and remove the timer from the map.
							instancesInfos[i].CancelTime = time.Now().Unix()
							instancesInfos[i].Status = "Canceled"
							tasks[instancesInfos[i].TaskId].Stop()
							delete(tasks, instancesInfos[i].TaskId)

							// Send event message...
							var eventDatas []*MessageData
							evtData := new(MessageData)
							evtData.TYPENAME = "Server.MessageData"
							evtData.Name = "taskInfos"
							evtData.Value = instancesInfos[i]

							eventDatas = append(eventDatas, evtData)
							evt, _ := NewEvent(UpdateTaskEvent, ConfigurationEvent, eventDatas)
							GetServer().GetEventManager().BroadcastEvent(evt)
						}
					}
				}
			}

			// Set the timer.
			var instanceInfos *TaskInstanceInfo

			// Plan the next instance.
			startTime := time.Unix(task.GetStartTime(), 0)
			delay := startTime.Sub(time.Now())
			timer := time.NewTimer(delay)
			tasks[task.GetId()] = timer

			// Keep task info here.
			instanceInfos = new(TaskInstanceInfo)
			instanceInfos.UUID = Utility.RandomUUID()
			instanceInfos.CreationTime = time.Now().Unix() // Time of task creation.
			// append to the list of task.
			instancesInfos = append(instancesInfos, instanceInfos)

			instanceInfos.TYPENAME = "Server.TaskInstanceInfo"
			instanceInfos.TaskId = task.GetId()
			instanceInfos.StartTime = task.GetStartTime() // When the task is planed to start

			// Run the task.
			runTask := func() {
				err := GetTaskManager().runTask(task) // run the task.
				instanceInfos.EndTime = time.Now().Unix()
				if err != nil {
					instanceInfos.Status = "Failed"
					instanceInfos.Error = err.Error()
				}
				if instanceInfos.Status != "Failed" && instanceInfos.Status != "Canceled" {
					instanceInfos.Status = "Completed"
				}
				// Send event message...
				var eventDatas []*MessageData
				evtData := new(MessageData)
				evtData.TYPENAME = "Server.MessageData"
				evtData.Name = "taskInfos"
				evtData.Value = instanceInfos
				eventDatas = append(eventDatas, evtData)
				evt, _ := NewEvent(UpdateTaskEvent, ConfigurationEvent, eventDatas)
				GetServer().GetEventManager().BroadcastEvent(evt)
			}

			if delay > 0 {
				instanceInfos.Status = "Scheduled"
				go func(task *Config.ScheduledTask, timer *time.Timer, instanceInfos *TaskInstanceInfo) {
					<-timer.C // wait util the delay expire...
					runTask()
				}(task, timer, instanceInfos)
			} else {
				instanceInfos.Status = "Running"
				go runTask()
			}

			// Send event message...
			var eventDatas []*MessageData
			evtData := new(MessageData)
			evtData.TYPENAME = "Server.MessageData"
			evtData.Name = "taskInfos"
			evtData.Value = instanceInfos

			eventDatas = append(eventDatas, evtData)
			evt, _ := NewEvent(NewTaskEvent, ConfigurationEvent, eventDatas)
			GetServer().GetEventManager().BroadcastEvent(evt)

		// reset the task.
		case uuid := <-this.m_cancelScheduledTasksChan:
			for i := 0; i < len(instancesInfos); i++ {
				if instancesInfos[i].UUID == uuid {
					instancesInfos[i].CancelTime = time.Now().Unix()
					instancesInfos[i].Status = "Canceled"
					if tasks[instancesInfos[i].TaskId] != nil {
						// Stop and remove the timer from the map.
						tasks[instancesInfos[i].TaskId].Stop()
						delete(tasks, instancesInfos[i].TaskId)
						// I will stop the vm that run that task.
						log.Println("--> Supend the task: ", instancesInfos[i].TaskId)
						JS.GetJsRuntimeManager().CloseSession(instancesInfos[i].TaskId,
							func(taskId string) func() {
								return func() {
									log.Println("--> task with id " + taskId + " is now stop!")
									// Now I will kill it commands.
									cmds := GetServer().sessionCmds[taskId]
									for i := 0; i < len(cmds); i++ {
										// Remove the command
										GetServer().removeCmd(cmds[i])
									}
								}
							}(instancesInfos[i].TaskId))

						// Send event message...
						var eventDatas []*MessageData
						evtData := new(MessageData)
						evtData.TYPENAME = "Server.MessageData"
						evtData.Name = "taskInfos"
						evtData.Value = instancesInfos[i]

						eventDatas = append(eventDatas, evtData)
						evt, _ := NewEvent(UpdateTaskEvent, ConfigurationEvent, eventDatas)
						GetServer().GetEventManager().BroadcastEvent(evt)

						break
					}
				}
			}

		// Return the list of scheduled task.
		case getTasksChan := <-this.m_getTaskInstancesInfo:
			getTasksChan <- instancesInfos
		}
	}

}

/**
 * Return the singleton.
 */
func GetTaskManager() *TaskManager {
	if taskManager == nil {
		return newTaskManager()
	}
	return taskManager
}

/**
 * Run a task.
 */
func (this *TaskManager) runTask(task *Config.ScheduledTask) error {

	// Error handler.
	defer func() {
		// Stahp mean the VM was kill by the admin.
		if caught := recover(); caught != nil {
			if caught.(error).Error() == "Stahp" {
				// Here the task was cancel.
				return
			} else {
				panic(caught) // Something else happened, repanic!
			}
		}
	}()

	// first of all I will test if the task is active.
	if task.IsActive() == false {
		return nil // Nothing to do here.
	}

	dbFile, err := GetServer().GetEntityManager().getEntityById("CargoEntities.File", "CargoEntities", []interface{}{task.M_script})
	if err == nil {
		script, err := b64.StdEncoding.DecodeString(dbFile.(*CargoEntities.File).GetData())
		// Now I will run the script...
		if err == nil {
			// Open a new session if none is already open.
			JS.GetJsRuntimeManager().OpenSession(task.GetId())
			_, err := JS.GetJsRuntimeManager().RunScript(task.GetId(), string(script))
			if err != nil {
				log.Println("--> script error: ", err)
				return err
			}

		} else {
			log.Println("--> script error: ", err)
			return err
		}
		if task.GetFrequencyType() != Config.FrequencyType_ONCE {
			// So here I will re-schedule the task, it will not be schedule if is it
			// expired or it must run once.
			GetTaskManager().scheduleTask(task)
		} else {
			if task.GetStartTime() > 0 {
				GetTaskManager().scheduleTask(task)
			}
		}
	} else {
		log.Println("--> script error: ", err.GetBody())
		return errors.New(err.GetBody())
	}

	log.Println("--> task ", task.GetId(), "run successfully!")
	return nil
}

// daysIn returns the number of days in a month for a given year.
func daysIn(m time.Month, year int) int {
	// This is equivalent to time.daysIn(m, year).
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

/**
 * That function retun the list of task instances.
 */
func (this *TaskManager) getTaskInstancesInfos() []*TaskInstanceInfo {
	taskInfosChan := make(chan []*TaskInstanceInfo)
	this.m_getTaskInstancesInfo <- taskInfosChan
	taskInfos := <-taskInfosChan
	return taskInfos
}

/**
 * That function is use to schedule a task.
 */
func (this *TaskManager) scheduleTask(task *Config.ScheduledTask) {

	// first of all I will test if the task is active.
	if task.IsActive() == false {
		return // Nothing to do here.
	}

	// If the task is expired
	if task.M_expirationTime > 0 {
		if task.M_expirationTime < time.Now().Unix() {
			// The task has expire!
			task.SetIsActive(false)
			return
		}
	}

	var nextTime time.Time
	if task.GetFrequencyType() != Config.FrequencyType_ONCE {
		startTime := time.Unix(task.M_startTime, 0)

		// Now I will get the next time when the task must be executed.
		nextTime = startTime
		var previous time.Time
		frequency := task.GetFrequency()
		if frequency == 0 {
			frequency = 1 // Must be one by default if not specify.
		}

		for nextTime.Sub(time.Now()) < 0 {
			previous = nextTime
			// I will append
			if task.GetFrequencyType() == Config.FrequencyType_DAILY {
				f := time.Duration((24 * 60 * 60 * 1000) / frequency)
				nextTime = nextTime.Add(f * time.Millisecond)
			} else if task.GetFrequencyType() == Config.FrequencyType_WEEKELY {
				f := time.Duration((7 * 24 * 60 * 60 * 1000) / frequency)
				nextTime = nextTime.Add(f * time.Millisecond)
			} else if task.GetFrequencyType() == Config.FrequencyType_MONTHLY {
				numberOfDay := daysIn(nextTime.Month(), nextTime.Year())
				f := time.Duration((numberOfDay * 24 * 60 * 60 * 1000) / frequency)
				nextTime = nextTime.Add(f * time.Millisecond)
			}
		}

		// Here I will test if the previous time combine with offset value can
		// be use as nextTime.
		for i := 0; i < len(task.M_offsets); i++ {
			if task.GetFrequencyType() == Config.FrequencyType_WEEKELY || task.GetFrequencyType() == Config.FrequencyType_MONTHLY {
				offset := time.Duration(task.M_offsets[i] * 24 * 60 * 60 * 1000) // in hours
				nextTime_ := previous.Add(offset * time.Millisecond)
				if nextTime_.Sub(time.Now()) > 0 {
					nextTime = nextTime_
					break
				}
			} else if task.GetFrequencyType() == Config.FrequencyType_DAILY {
				// Here the offset represent hours and not days.
				offset := time.Duration(task.M_offsets[i] * 60 * 60 * 1000) // in hours
				nextTime_ := previous.Add(offset * time.Millisecond)
				if nextTime_.Sub(time.Now()) > 0 {
					nextTime = nextTime_
					break
				}
			}
		}

		// Set it start time...
		task.SetStartTime(nextTime.Unix())

	} else {
		if task.GetStartTime() == 0 {
			// Run the task directly in that case.
			this.m_setScheduledTasksChan <- task
			return // nothing more to do...
		} else {
			startTime := time.Unix(task.GetStartTime(), 0)
			if startTime.Sub(time.Now()) < 0 {
				// innactivate the task in that case.
				task.SetIsActive(false)
			}
		}
	}

	// Save modification.
	GetServer().GetEntityManager().saveEntity(task)

	// Process the task.
	startTime := time.Unix(task.GetStartTime(), 0)
	if startTime.Sub(time.Now()) >= 0 {
		this.m_setScheduledTasksChan <- task
	}
}
