// +build BPMN20

package Server

import (
	"log"

	"bufio"
	"bytes"
	b64 "encoding/base64"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/BPMN20"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/BPMS"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/CargoWebServer/Cargo/JS"
	"code.myceliUs.com/Utility"
	"github.com/robertkrimen/otto"
)

const (
	// make the processor more responsive...
	resolution = 5

	// Datastore names...
	BPMSDB = "BPMS"
)

/**
 * Tag used to identified a process instance during it lifetime.
 * The tag will be reatributed latter...
 */
type ColorTag struct {
	Name       string // Acid Green
	Number     string // Hexadecimal #B0BF1A
	InstanceId string // if is used by an instance.
}

type WorkflowProcessor struct {

	// The workflow manage stop when that variable
	// is set to true.
	abortedByEnvironment chan bool

	// The trigger channel.
	triggerChan chan *BPMS.Trigger

	// The runtime...
	runtime *BPMS_RuntimesEntity

	// Process instance have a color tag to be
	// easier to regonized by the end user. The color must be unique
	// for the lifetime of the instance, but the color can be reuse latter by
	// other instance.
	colorTagMap map[string]*ColorTag

	// Contain the list of available tag number...
	availableTagNumber []string
}

var workflowProcessor *WorkflowProcessor

func (this *Server) GetWorkflowProcessor() *WorkflowProcessor {
	if workflowProcessor == nil {
		workflowProcessor = newWorkflowProcessor()
	}
	return workflowProcessor
}

func newWorkflowProcessor() *WorkflowProcessor {
	// The workflow manager runtime entities
	bpmnRuntimeDB := new(Config.DataStoreConfiguration)
	bpmnRuntimeDB.M_id = BPMSDB
	bpmnRuntimeDB.M_dataStoreVendor = Config.DataStoreVendor_MYCELIUS
	bpmnRuntimeDB.M_dataStoreType = Config.DataStoreType_KEY_VALUE_STORE
	bpmnRuntimeDB.NeedSave = true

	GetServer().GetConfigurationManager().appendDefaultDataStoreConfiguration(bpmnRuntimeDB)
	GetServer().GetDataManager().appendDefaultDataStore(bpmnRuntimeDB)

	// Create and register entity prototypes for dynamic typing system.
	GetServer().GetEntityManager().createBPMSPrototypes()
	GetServer().GetEntityManager().registerBPMSObjects()

	// The workflow manger...
	workflowProcessor := new(WorkflowProcessor)
	workflowProcessor.abortedByEnvironment = make(chan bool)

	// The trigger chanel...
	workflowProcessor.triggerChan = make(chan *BPMS.Trigger)

	return workflowProcessor
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

func (this *WorkflowProcessor) initialize() {

	log.Println("--> Initialize WorkflowProcessor")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())

	// The runtime data...
	runtimeUUID := BPMSRuntimesExists("runtime")
	if len(runtimeUUID) == 0 {
		// Create a new runtime...
		this.runtime = GetServer().GetEntityManager().NewBPMSRuntimesEntity("runtime", nil)
		this.runtime.SaveEntity()

	} else {
		// Get the current runtime...
		this.runtime = GetServer().GetEntityManager().NewBPMSRuntimesEntity(runtimeUUID, nil)
		// Initialyse the runtime...
		this.runtime.InitEntity(runtimeUUID)
	}

	// Here I will read the color tag...
	this.colorTagMap = make(map[string]*ColorTag, 0)
	colorNamePath := GetServer().GetConfigurationManager().GetDataPath() + "/colorName.csv"
	colorNameFile, _ := os.Open(colorNamePath)
	csvReader := csv.NewReader(bufio.NewReader(colorNameFile))

	for {
		record, err := csvReader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		tag := new(ColorTag)
		for i := 0; i < len(record); i++ {
			if i == 0 {
				tag.Name = record[i]
			} else if i == 1 {
				tag.Number = record[i]
				this.availableTagNumber = append(this.availableTagNumber, tag.Number)
			}
		}
		this.colorTagMap[tag.Number] = tag
	}

	// That function will evaluate the processes and put it in the good
	// state to start...
	this.runTopLevelProcesses()
	this.getNextProcessInstanceNumber()
}

func (this *WorkflowProcessor) getId() string {
	return "WorkflowProcessor"
}

/**
 * Start the service.
 */
func (this *WorkflowProcessor) start() {
	log.Println("--> Start WorkflowProcessor")
	go this.run()
}

func (this *WorkflowProcessor) stop() {
	log.Println("--> Stop WorkflowProcessor")

}

// Start processing the workflow...
func (this *WorkflowProcessor) run() {

	for {
		select {
		case trigger := <-this.triggerChan:
			this.processTrigger(trigger) // Process the triggers...
		case done := <-this.abortedByEnvironment:
			if done {
				log.Println("Stop processing workflow")
				return
			}
		default:
			// Run the to level process each 300 millisecond.
			this.runTopLevelProcesses()
		}
		// Does nothing for the next 300 millisecond...
		time.Sleep(300 * time.Millisecond)

	}
}

/**
 * Process new trigger... The trigger contain information, data, that the
 * workflow processor use to evaluate the next action evalute...
 */
func (this *WorkflowProcessor) processTrigger(trigger *BPMS.Trigger) *CargoEntities.Error {
	// Now i will process the trigger...

	// New process trigger was created...
	if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Start {

		// Get the bpmn process...
		processEntity, err := GetServer().GetEntityManager().getEntityByUuid(trigger.M_processUUID)
		if err != nil {
			return err
		}

		// Get it definitions.
		definitionsEntity := processEntity.GetParentPtr()
		definitionsInstance := this.getDefinitionInstance(definitionsEntity.GetObject().(*BPMN20.Definitions).GetId(), definitionsEntity.GetObject().(*BPMN20.Definitions).GetUUID())

		/** So Here I will create a new process instance **/
		processInstance := new(BPMS.ProcessInstance)
		processInstance.UUID = "BPMS.ProcessInstance%" + Utility.RandomUUID()
		processInstance.M_number = this.getNextProcessInstanceNumber()

		index := rand.Intn(len(this.availableTagNumber))
		colorTag := this.colorTagMap[this.availableTagNumber[index]]
		processInstance.M_colorNumber = colorTag.Number
		processInstance.M_colorName = colorTag.Name

		// Here I will associate the bmpn process instance id.
		processInstance.M_bpmnElementId = trigger.M_processUUID

		// Now the start event...
		var startEvent *BPMN20.StartEvent
		for i := 0; i < len(processEntity.GetChildsPtr()); i++ {
			if processEntity.GetChildsPtr()[i].GetTypeName() == "BPMN20.StartEvent" {
				startEvent = processEntity.GetChildsPtr()[i].GetObject().(*BPMN20.StartEvent)
				break
			}
		}

		// Evaluation of the source expression if there is so...
		// Get a reference to the process.
		process := processEntity.GetObject().(*BPMN20.Process)
		processInstance.SetId(process.M_id)

		// Set the process instances...
		definitionsInstance.SetProcessInstances(processInstance)
		processInstance.SetParentPtr(definitionsInstance)

		// Create empty itemaware elements in the process instance.
		this.createItemAwareElementInstances(processInstance)

		// Now I will create the start event...
		startEventInstance := this.createInstance(startEvent, processInstance, nil, trigger.GetDataRef(), trigger.M_sessionId)

		// Set the new definition instance...
		this.runtime.object.SetDefinitions(definitionsInstance)

		// Save the new entity...
		this.runtime.SaveEntity()

		// we are done with the start event.
		startEventInstance.(BPMS.FlowNodeInstance).SetLifecycleState(BPMS.LifecycleState_Completing)

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_None {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Timer {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Conditional {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Message {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Signal {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Multiple {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_ParallelMultiple {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Escalation {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Error {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Compensation {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Terminate {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Cancel {

	} else if trigger.GetEventTriggerType() == BPMS.EventTriggerType_Link {

	}

	// No error here...
	return nil
}

/**
 * Evaluate processes instance for each bpmn processes.
 */
func (this *WorkflowProcessor) runTopLevelProcesses() {
	// Get all processes...
	processes := GetServer().GetWorkflowManager().getProcess()
	for i := 0; i < len(processes); i++ {
		this.workflowTransitionInterpreter(processes[i])
	}
}

/**
 * Evaluate the transition.
 */
func (this *WorkflowProcessor) workflowTransitionInterpreter(process *BPMN20.Process) {
	activeProcessInstances := this.getActiveProcessInstances(process)
	for i := 0; i < len(activeProcessInstances); i++ {
		for j := 0; j < len(activeProcessInstances[i].GetFlowNodeInstances()); j++ {
			this.workflowTransition(activeProcessInstances[i].GetFlowNodeInstances()[j], activeProcessInstances[i], "")
		}
	}
}

/**
 * Evaluate the transition...
 */
func (this *WorkflowProcessor) workflowTransition(flowNode BPMS.FlowNodeInstance, processInstance *BPMS.ProcessInstance, sessionId string) {

	// I will retreive the entity related to the instance...
	bpmnElementEntity, err := GetServer().GetEntityManager().getEntityByUuid(flowNode.(BPMS.Instance).GetBpmnElementId())
	if err == nil {
		if flowNode.GetLifecycleState() == BPMS.LifecycleState_Active {
			// Here I will process the active instances...
			if reflect.TypeOf(flowNode).String() == "*BPMS.ActivityInstance" {
				if strings.HasPrefix(flowNode.(*BPMS.ActivityInstance).M_bpmnElementId, "BPMN20.UserTask%") {
					err := this.executeActivityInstance(flowNode.(*BPMS.ActivityInstance), "")
					if err == nil {
						flowNode.SetLifecycleState(BPMS.LifecycleState_Completing)
					}
				}
			} else if reflect.TypeOf(flowNode).String() == "*BPMS.EventInstance" {
				if reflect.TypeOf(bpmnElementEntity.GetObject()).String() == "*BPMN20.EndEvent" {
					// here I just set the life cycle to complete...
					flowNode.SetLifecycleState(BPMS.LifecycleState_Completing)
				}

			} else if reflect.TypeOf(flowNode).String() == "*BPMS.GatewayInstance" {

			} else if reflect.TypeOf(flowNode).String() == "*BPMS.SubprocessInstance" {

			}
		} else if flowNode.GetLifecycleState() == BPMS.LifecycleState_Completing {
			bpmnElement := bpmnElementEntity.GetObject().(BPMN20.FlowNode)

			// Here i will process the completed instance...
			for i := 0; i < len(bpmnElement.GetOutgoing()); i++ {
				// TODO in case of a gateway evaluate the sequence flow condition...
				seqFlow := bpmnElement.GetOutgoing()[i]

				connectingObj := new(BPMS.ConnectingObject)
				connectingObj.UUID = "BPMS.ConnectingObject%" + Utility.RandomUUID()
				connectingObj.M_bpmnElementId = seqFlow.GetUUID()
				connectingObj.SetSourceRef(flowNode)
				connectingObj.SetId(bpmnElement.GetOutgoing()[i].GetId())

				// Now I will create the next node

				// The item reference must be set...
				items := make([]*BPMS.ItemAwareElementInstance, 0)

				// Create the new instance.
				nextStep := this.createInstance(seqFlow.GetTargetRef(), processInstance, connectingObj, items, "")

				connectingObj.SetTargetRef(nextStep)

				flowNode.SetOutputRef(connectingObj)

				// Set the connecting object and save it...
				processInstance.SetConnectingObjects(connectingObj)

				processInstanceEntity := GetServer().GetEntityManager().NewBPMSProcessInstanceEntityFromObject(processInstance)

				// Save the change...
				processInstanceEntity.SaveEntity()
			}

			// Set completed...
			this.completeInstance(flowNode, sessionId)
		}
	}
}

/**
 * Do the work inside the activity instance if there some work todo...
 */
func (this *WorkflowProcessor) executeActivityInstance(instance BPMS.Instance, sessionId string) *CargoEntities.Error {

	// I will retreive the activity entity related to the the instance.
	activityEntity, err := GetServer().GetEntityManager().getEntityByUuid(instance.GetBpmnElementId())
	if err != nil {
		return err
	}

	// Get the activity object.
	activity := activityEntity.GetObject()

	processEntity, _ := GetServer().GetEntityManager().getEntityByUuid(instance.(*BPMS.ActivityInstance).GetProcessInstancePtr().GetBpmnElementId())
	definitions := processEntity.GetObject().(*BPMN20.Process).GetDefinitionsPtr()

	switch v := activity.(type) {
	case *BPMN20.AdHocSubProcess:

	case *BPMN20.BusinessRuleTask:

	case *BPMN20.CallActivity:

	case *BPMN20.SubProcess_impl:

	case *BPMN20.ManualTask:

	case *BPMN20.ScriptTask:
		if activity.(*BPMN20.ScriptTask).GetScript() != nil {
			activity.(*BPMN20.ScriptTask).GetScript()
			// List the variable put in the vm...
			instances := make(map[string]*BPMS.ItemAwareElementInstance)
			this.appendItemAwareElementInstance(instance, instances)

			var definitionInstances []BPMS.Instance
			// append defintion instance item aware... (DataStore)
			definitionInstances, _ = GetServer().GetWorkflowManager().getInstances(definitions.UUID, "BPMS.DefinitionsInstance")
			for i := 0; i < len(definitionInstances); i++ {
				this.appendItemAwareElementInstance(definitionInstances[i], instances)
			}

			script := "function run(){\n" + activity.(*BPMN20.ScriptTask).GetScript().GetScript() + "\n}"
			instances, err := this.runScript(definitions, script, "run", instances, sessionId)

			if err != nil {
				return NewError(Utility.FileLine(), ACTION_EXECUTE_ERROR, SERVER_ERROR_CODE, err)
			}
		}
		return nil
	case *BPMN20.UserTask:
		log.Println("----------> user task ", v.GetUUID())

	case *BPMN20.Task_impl:

	default:
		log.Println("-------> Unknow activity type ", v)
	}

	return nil
}

/**
 * That function is use to evaluate a given script.
 */
func (this *WorkflowProcessor) runScript(definitions *BPMN20.Definitions, functionStr string, functionName string, instances map[string]*BPMS.ItemAwareElementInstance, sessionId string) (map[string]*BPMS.ItemAwareElementInstance, error) {
	var err error

	// Remove the xml tag here...
	functionStr = strings.Replace(functionStr, "<![CDATA[", "", -1)
	functionStr = strings.Replace(functionStr, "]]>", "", -1)

	// Set the instance on the context...
	vm := JS.GetJsRuntimeManager().GetVm(sessionId).Copy()
	vm.Set("sessionId", sessionId)

	// Set the server pointer and their services...
	vm.Set("server", GetServer())

	// I will get the list of stores...
	for i := 0; i < len(definitions.GetRootElement()); i++ {
		if reflect.TypeOf(definitions.GetRootElement()[i]).String() == "*BPMN20.DataStore" {
			dataStore := definitions.GetRootElement()[i].(*BPMN20.DataStore)
			this.initDatastore(dataStore.GetName(), vm)
		}
	}

	// Here I will set instance data on the vm context...
	for k, v := range instances {
		objects, isCollection, err := this.getItemawareElementData(v)
		if err == nil {
			if isCollection {
				// set an empty array
				object, _ := vm.Object(k + "=[]")

				// push the value in it.
				for i := 0; i < len(objects); i++ {
					object.Call("push", objects[i])
				}

			} else {
				if len(objects) > 0 {
					vm.Set(k, objects[0])
				} else {
					vm.Set(k, nil)
				}
			}
		} else {
			log.Println("----->462 data error: ", err.Error())
		}
	}

	// Set the function string
	_, err = vm.Run(functionStr)
	if err != nil {
		log.Println("Error in code of ", functionName)
		return instances, err
	}

	// Call the function.
	_, err = vm.Run(functionName + "();")
	if err != nil {
		log.Println("Error in code of ", functionName)
		return instances, err
	}

	// Here all was correctly runing, I will get back the items data and set
	// back to the item objects.
	for k, v := range instances {
		data, _ := vm.Get(k)
		if v != nil {
			if !data.IsUndefined() {
				this.setItemawareElementData(v, &data)
				// set back the instance.
				instances[v.M_id] = v
			}
		} else {
			log.Println("---> value assiciated with ", k, " is null!!!")
		}
	}

	return instances, err
}

////////////////////////////////////////////////////////////////////////////////
// Various constructor function.
////////////////////////////////////////////////////////////////////////////////
/**
 * Create the instance from
 */
func (this *WorkflowProcessor) createInstance(flowNode BPMN20.FlowNode, processInstance *BPMS.ProcessInstance, input *BPMS.ConnectingObject, items []*BPMS.ItemAwareElementInstance, sessionId string) BPMS.Instance {
	var instance BPMS.Instance

	switch v := flowNode.(type) {
	case BPMN20.Activity:
		//log.Println("-------> create activity ", v)
		instance = this.createActivityInstance(v, processInstance, items, sessionId)
		instance.(*BPMS.ActivityInstance).SetInputRef(input)
		//log.Println("-------> after create activity process data:", processInstance.GetData())
	case BPMN20.Gateway:
		//log.Println("-------> create gateway ", v)
		instance.(*BPMS.GatewayInstance).SetInputRef(input)

	case BPMN20.Event:
		//log.Println("-------> create event ", v)
		instance = this.createEventInstance(v, processInstance, items, sessionId)
		if input != nil {
			instance.(*BPMS.EventInstance).SetInputRef(input)
		}
		//log.Println("-------> after start event creation process data:", processInstance.GetData())
	case *BPMN20.Transaction:
		// Nothing todo here... for now...

	default:
		log.Println("-------> not define ", v)
	}

	// Set the log information here...
	if instance != nil {
		this.setLogInfo(instance.(BPMS.FlowNodeInstance), "New instance created", sessionId)
	}

	// Get the entity for the process instance.
	processInstance.NeedSave = true
	processEntity := GetServer().GetEntityManager().NewBPMSProcessInstanceEntityFromObject(processInstance)
	processEntity.SaveEntity()

	return instance
}

/**
 * Create Activity instance
 */
func (this *WorkflowProcessor) createActivityInstance(activity BPMN20.Activity, processInstance *BPMS.ProcessInstance, items []*BPMS.ItemAwareElementInstance, sessionId string) *BPMS.ActivityInstance {

	instanceEntity := GetServer().GetEntityManager().NewBPMSActivityInstanceEntity("BPMS.ActivityInstance%"+Utility.RandomUUID(), nil)
	instance := instanceEntity.GetObject().(*BPMS.ActivityInstance)
	instance.M_bpmnElementId = activity.GetUUID()
	instance.M_id = activity.(BPMN20.BaseElement).GetId()
	instance.M_bpmnElementId = activity.GetUUID()

	// Set the process instance here.
	instance.SetProcessInstancePtr(processInstance)

	// Set state to ready...
	instance.M_lifecycleState = BPMS.LifecycleState_Ready

	// I will copy the data reference into the start event...
	for i := 0; i < len(items); i++ {
		instance.SetDataRef(items[i])
		items[i].SetParentPtr(instance)
	}

	// Create missing itemware elements in the instance.
	this.createItemAwareElementInstances(instance)

	switch v := activity.(type) {
	case *BPMN20.AdHocSubProcess:
		instance.M_flowNodeType = BPMS.FlowNodeType_AdHocSubprocess
		instance.M_activityType = BPMS.ActivityType_AdHocSubprocess
	case *BPMN20.BusinessRuleTask:
		instance.M_flowNodeType = BPMS.FlowNodeType_BusinessRuleTask
		instance.M_activityType = BPMS.ActivityType_BusinessRuleTask
	case *BPMN20.CallActivity:
		instance.M_flowNodeType = BPMS.FlowNodeType_CallActivity
		instance.M_activityType = BPMS.ActivityType_CallActivity
	case *BPMN20.SubProcess_impl:
		instance.M_flowNodeType = BPMS.FlowNodeType_EmbeddedSubprocess
		instance.M_activityType = BPMS.ActivityType_EmbeddedSubprocess
	case *BPMN20.ManualTask:
		instance.M_flowNodeType = BPMS.FlowNodeType_ManualTask
		instance.M_activityType = BPMS.ActivityType_ManualTask
	case *BPMN20.ScriptTask:
		instance.M_flowNodeType = BPMS.FlowNodeType_ScriptTask
		instance.M_activityType = BPMS.ActivityType_ScriptTask
	case *BPMN20.UserTask:
		instance.M_flowNodeType = BPMS.FlowNodeType_UserTask
		instance.M_activityType = BPMS.ActivityType_UserTask
	case *BPMN20.Task_impl:
		instance.M_flowNodeType = BPMS.FlowNodeType_AbstractTask
		instance.M_activityType = BPMS.ActivityType_AbstractTask
	default:
		log.Println("-------> Unknow activity type ", v)
	}

	// Set the start event for the processInstance...
	processInstance.SetFlowNodeInstances(instance)

	// I can activate the instance...
	this.activateInstance(instance, sessionId)

	return instance
}

/**
 * Create Event Instance
 */
func (this *WorkflowProcessor) createEventInstance(event BPMN20.Event, processInstance *BPMS.ProcessInstance, items []*BPMS.ItemAwareElementInstance, sessionId string) *BPMS.EventInstance {
	// Now I will create the start event...
	instanceEntity := GetServer().GetEntityManager().NewBPMSEventInstanceEntity("BPMS.EventInstance%"+Utility.RandomUUID(), nil)
	instance := instanceEntity.GetObject().(*BPMS.EventInstance)
	instance.M_bpmnElementId = event.GetUUID()
	instance.M_id = event.(BPMN20.BaseElement).GetId()
	instance.SetProcessInstancePtr(processInstance)

	// I will copy the data reference into the start event...
	for i := 0; i < len(items); i++ {
		instance.SetDataRef(items[i])
		items[i].SetParentPtr(instance)
	}

	// Create missing itemware elements in the instance.
	this.createItemAwareElementInstances(instance)

	// Select the good event type...
	switch v := event.(type) {
	case *BPMN20.StartEvent:
		instance.M_flowNodeType = BPMS.FlowNodeType_StartEvent
		instance.M_eventType = BPMS.EventType_StartEvent
	case *BPMN20.EndEvent:
		instance.M_flowNodeType = BPMS.FlowNodeType_EndEvent
		instance.M_eventType = BPMS.EventType_EndEvent
	case *BPMN20.BoundaryEvent:
		instance.M_flowNodeType = BPMS.FlowNodeType_BoundaryEvent
		instance.M_eventType = BPMS.EventType_BoundaryEvent
	case *BPMN20.IntermediateCatchEvent:
		instance.M_flowNodeType = BPMS.FlowNodeType_IntermediateCatchEvent
		instance.M_eventType = BPMS.EventType_IntermediateCatchEvent
	case *BPMN20.IntermediateThrowEvent:
		instance.M_flowNodeType = BPMS.FlowNodeType_IntermediateThrowEvent
		instance.M_eventType = BPMS.EventType_IntermediateThrowEvent
	default:
		log.Println("------> event ", v, " has not instance type...")
	}

	// Set the instance to ready state...
	instance.M_lifecycleState = BPMS.LifecycleState_Ready

	// Set the start event for the processInstance...
	processInstance.SetFlowNodeInstances(instance)

	// I can activate the instance...
	this.activateInstance(instance, sessionId)

	return instance
}

/**
 * Create a new Item aware element instance for a given bpmnElementId...
 * The bpmn element must be a ItemAwareElement.
 */
func (this *WorkflowProcessor) createItemAwareElementInstance(bpmnElementId string, data interface{}) (*BPMS.ItemAwareElementInstance, *CargoEntities.Error) {

	bpmnElementEntity, err := GetServer().GetEntityManager().getEntityByUuid(bpmnElementId)
	if err != nil {
		return nil, err
	}

	// Here I will simply serialyse the data and save it...
	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	err_ := enc.Encode(data)
	if err != nil {
		return nil, NewError(Utility.FileLine(), JSON_MARSHALING_ERROR, SERVER_ERROR_CODE, err_)
	}

	// Here I will set the instance id...
	instanceEntity := GetServer().GetEntityManager().NewBPMSItemAwareElementInstanceEntity("", nil)

	instance := instanceEntity.GetObject().(*BPMS.ItemAwareElementInstance)
	instance.M_bpmnElementId = bpmnElementId
	instance.M_id = bpmnElementEntity.GetObject().(BPMN20.BaseElement).GetId()
	instance.M_data = buffer.Bytes()

	// Now I will save the entity
	instanceEntity.SaveEntity()

	return instance, nil
}

////////////////////////////////////////////////////////////////////////////////
// State transition function.
////////////////////////////////////////////////////////////////////////////////

// Utility function to save a given entity
func (this *WorkflowProcessor) saveInstance(instance BPMS.Instance) {
	var entity Entity

	switch v := instance.(type) {
	case *BPMS.ActivityInstance:
		v.GetProcessInstancePtr().SetFlowNodeInstances(instance)
		entity = GetServer().GetEntityManager().NewBPMSProcessInstanceEntityFromObject(v.GetProcessInstancePtr())
	case *BPMS.ProcessInstance:
		entity = GetServer().GetEntityManager().NewBPMSProcessInstanceEntityFromObject(v)
	case *BPMS.EventInstance:
		v.GetProcessInstancePtr().SetFlowNodeInstances(instance)
		entity = GetServer().GetEntityManager().NewBPMSProcessInstanceEntityFromObject(v.GetProcessInstancePtr())
	case *BPMS.GatewayInstance:
		v.GetProcessInstancePtr().SetFlowNodeInstances(instance)
		entity = GetServer().GetEntityManager().NewBPMSProcessInstanceEntityFromObject(v.GetProcessInstancePtr())
	}

	if entity != nil {
		entity.SaveEntity()
	}
}

/**
 * Activate a given instance.
 */
func (this *WorkflowProcessor) activateInstance(instance BPMS.FlowNodeInstance, sessionId string) error {
	if instance.GetLifecycleState() != BPMS.LifecycleState_Ready {
		return errors.New("The input flow node instance must have life cycle state at Ready")
	}

	isUserTask := false

	// Initialisation of the data input.
	switch v := instance.(type) {
	case *BPMS.ActivityInstance:
		isUserTask = strings.HasPrefix(v.M_bpmnElementId, "BPMN20.UserTask")
		bpmnElement, _ := GetServer().GetEntityManager().getEntityByUuid(v.M_bpmnElementId)
		activity := bpmnElement.GetObject().(BPMN20.Activity)

		var dataInputAssociations []BPMN20.DataAssociation
		for i := 0; i < len(activity.GetDataInputAssociation()); i++ {
			dataInputAssociations = append(dataInputAssociations, activity.GetDataInputAssociation()[i])
		}

		this.setDataAssocication(v.GetProcessInstancePtr(), v, dataInputAssociations, sessionId)

	case *BPMS.EventInstance:

		// Get the data objects from the flow elements
		var dataAssociations []BPMN20.DataAssociation
		properties := make([]BPMN20.BaseElement, 0)
		bpmnElement, _ := GetServer().GetEntityManager().getEntityByUuid(v.M_bpmnElementId)
		event := bpmnElement.GetObject().(BPMN20.Event)
		processInstance := v.GetProcessInstancePtr()

		for i := 0; i < len(event.GetProperty()); i++ {
			properties = append(properties, event.GetProperty()[i])
		}

		// Set output data association...
		switch v_ := event.(type) {
		case BPMN20.CatchEvent:
			for i := 0; i < len(v_.GetDataOutputAssociation()); i++ {
				dataAssociations = append(dataAssociations, v_.GetDataOutputAssociation()[i])
			}
			if len(v_.GetDataOutputAssociation()) > 0 {
				this.setDataAssocication(v, processInstance, dataAssociations, sessionId)
			}
		case BPMN20.ThrowEvent:
			for i := 0; i < len(v_.GetDataInputAssociation()); i++ {
				dataAssociations = append(dataAssociations, v_.GetDataInputAssociation()[i])
			}
			if len(v_.GetDataInputAssociation()) > 0 {
				this.setDataAssocication(processInstance, v, dataAssociations, sessionId)
			}
		}

	case *BPMS.GatewayInstance:

	}

	// User task stay at ready state.
	if !isUserTask {
		instance.SetLifecycleState(BPMS.LifecycleState_Active)
	}

	this.saveInstance(instance.(BPMS.Instance))

	return nil
}

/**
 * Terminate a given instance.
 */
func (this *WorkflowProcessor) completeInstance(instance BPMS.FlowNodeInstance, sessionId string) error {
	if instance.GetLifecycleState() != BPMS.LifecycleState_Completing {
		return errors.New("The input flow node instance must have life cycle state at Completing")
	}

	// Set the data associations.
	switch v := instance.(type) {
	case *BPMS.ActivityInstance:
		bpmnElement, _ := GetServer().GetEntityManager().getEntityByUuid(v.M_bpmnElementId)
		activity := bpmnElement.GetObject().(BPMN20.Activity)

		var dataOutputAssociations []BPMN20.DataAssociation
		for i := 0; i < len(activity.GetDataOutputAssociation()); i++ {
			dataOutputAssociations = append(dataOutputAssociations, activity.GetDataOutputAssociation()[i])
		}

		// Contain regular data associations.
		this.setDataAssocication(v, v.GetProcessInstancePtr(), dataOutputAssociations, sessionId)

	case *BPMS.EventInstance:

	case *BPMS.GatewayInstance:

	}

	instance.SetLifecycleState(BPMS.LifecycleState_Completed)
	this.saveInstance(instance.(BPMS.Instance))
	return nil
}

/**
 * Terminate a given instance.
 */
func (this *WorkflowProcessor) terminateInstance(instance BPMS.FlowNodeInstance) error {
	if instance.GetLifecycleState() != BPMS.LifecycleState_Terminating {
		return errors.New("The input flow node instance must have life cycle state at Terminating")
	}

	// TODO do terminating stuff here.

	instance.SetLifecycleState(BPMS.LifecycleState_Terminated)
	this.saveInstance(instance.(BPMS.Instance))

	return nil
}

/**
 * Compensate a given instance.
 */
func (this *WorkflowProcessor) compensateInstance(instance BPMS.FlowNodeInstance) error {
	if instance.GetLifecycleState() != BPMS.LifecycleState_Compensating {
		return errors.New("The input flow node instance must have life cycle state at Compensating")
	}

	// TODO do compensating stuff here.

	instance.SetLifecycleState(BPMS.LifecycleState_Compensated)
	this.saveInstance(instance.(BPMS.Instance))

	return nil
}

/**
 * Abord a given instance.
 */
func (this *WorkflowProcessor) abordInstance(instance BPMS.FlowNodeInstance) error {
	if instance.GetLifecycleState() != BPMS.LifecycleState_Failing {
		return errors.New("The input flow node instance must have life cycle state at Failing")
	}

	// TODO do failling stuff here.
	instance.SetLifecycleState(BPMS.LifecycleState_Failed)
	this.saveInstance(instance.(BPMS.Instance))

	return nil
}

/**
 * Delete a process instance from the runtime when is no more needed...
 */
func (this *WorkflowProcessor) deleteInstance(processInstance *BPMS.ProcessInstance) {
	processInstaceEntity := GetServer().GetEntityManager().NewBPMSProcessInstanceEntityFromObject(processInstance)
	if processInstaceEntity != nil {
		//log.Println("---------> remove process instance: ", processInstance.GetUUID())
		// Remove the entity...
		//processInstaceEntity.DeleteEntity()
	}
}

/**
 * Determine if a process is still active.
 */
func (this *WorkflowProcessor) stillActive(processInstance *BPMS.ProcessInstance) bool {
	// Here If the process instance contain a end event instance with state
	// completed then it must be considere completed
	for i := 0; i < len(processInstance.GetFlowNodeInstances()); i++ {
		instance := processInstance.GetFlowNodeInstances()[i]
		if reflect.TypeOf(instance).String() == "*BPMS.EventInstance" {
			evt := instance.(*BPMS.EventInstance)
			if strings.HasPrefix(evt.GetBpmnElementId(), "BPMN20.EndEvent%") {
				if evt.GetLifecycleState() == BPMS.LifecycleState_Completed {
					return false
				}
			}
		}
	}

	// TODO evalute here other case where the process instance must be considere
	// innactive.

	return true
}

//////////////////////////////////////////////////////////////////////////////
// Data evaluation function.
//////////////////////////////////////////////////////////////////////////////

/**
 * That function is use to initialyse the classes contained inside a datastore.
 */
func (this *WorkflowProcessor) initDatastore(storeId string, vm *otto.Otto) {

	// Create the data store if he not already exist...
	dataStore_ := GetServer().GetDataManager().getDataStore(storeId)
	if dataStore_ == nil {
		dataStore_, _ = GetServer().GetDataManager().createDataStore(storeId, Config.DataStoreType_KEY_VALUE_STORE, Config.DataStoreVendor_MYCELIUS)
	}

	// So here I will get the prototypes from the store...
	prototypes, err := GetServer().GetEntityManager().getEntityPrototypes(storeId)
	if err == nil {
		for i := 0; i < len(prototypes); i++ {
			// Generate the constructor code.
			constructorSrc := prototypes[i].generateConstructor()
			_, err = vm.Run(constructorSrc)
			if err != nil {
				log.Println("878 ----> ", err)
			}
		}
	}
}

func (this *WorkflowProcessor) appendItemAwareElementInstance(instance BPMS.Instance, instances map[string]*BPMS.ItemAwareElementInstance) {
	// Initialisation of the data instances.
	for i := 0; i < len(instance.GetDataRef()); i++ {
		instances[instance.GetDataRef()[i].M_id] = instance.GetDataRef()[i]
	}

	// Ref
	for i := 0; i < len(instance.GetData()); i++ {
		instances[instance.GetData()[i].M_id] = instance.GetData()[i]
	}
}

func (this *WorkflowProcessor) getItemAwareElementId(ref BPMN20.ItemAwareElement) string {

	var id string

	if reflect.TypeOf(ref).String() == "*BPMN20.DataObjectReference" {
		id = ref.(*BPMN20.DataObjectReference).GetDataObjectRef().GetId()
	} else if reflect.TypeOf(ref).String() == "*BPMN20.DataStoreReference" {
		id = ref.(*BPMN20.DataStoreReference).GetDataStoreRef().GetId()
	} else {
		id = ref.(BPMN20.BaseElement).GetId()
	}

	return id
}

func (this *WorkflowProcessor) containItemAwareElementInstance(id string, instance BPMS.Instance) bool {
	for i := 0; i < len(instance.GetData()); i++ {
		if instance.GetData()[i].M_id == id {
			return true
		}
	}
	for i := 0; i < len(instance.GetDataRef()); i++ {
		if instance.GetDataRef()[i].M_id == id {
			return true
		}
	}
	return false
}

/**
 * Create instance if there are not already existing.
 */
func (this *WorkflowProcessor) createItemAwareElementInstances(instance BPMS.Instance) {

	bpmnElementEntity, err := GetServer().GetEntityManager().getEntityByUuid(instance.GetBpmnElementId())
	if err == nil {
		var itemAwareElements []BPMN20.ItemAwareElement
		bpmnElement := bpmnElementEntity.GetObject()

		// So here I will get data element from different type.
		switch v := bpmnElement.(type) {
		case *BPMN20.Definitions:
			definitions := v
			for i := 0; i < len(definitions.GetRootElement()); i++ {
				switch v := definitions.GetRootElement()[i].(type) {
				case *BPMN20.DataStore:
					itemAwareElements = append(itemAwareElements, v)
				}
			}
		case *BPMN20.Process:
			process := v
			// The property...
			for i := 0; i < len(process.GetProperty()); i++ {
				itemAwareElements = append(itemAwareElements, process.GetProperty()[i])
			}

			// The io specification.
			if process.GetIoSpecification() != nil {
				for i := 0; i < len(process.GetIoSpecification().GetDataOutput()); i++ {
					itemAwareElements = append(itemAwareElements, process.GetIoSpecification().GetDataOutput()[i])
				}

				for i := 0; i < len(process.GetIoSpecification().GetDataInput()); i++ {
					itemAwareElements = append(itemAwareElements, process.GetIoSpecification().GetDataInput()[i])
				}
			}

			for i := 0; i < len(process.GetFlowElement()); i++ {
				switch v := process.GetFlowElement()[i].(type) {
				case *BPMN20.DataObject:
					itemAwareElements = append(itemAwareElements, v)
				}
			}

		case BPMN20.Activity:
			activity := v
			// The property...
			for i := 0; i < len(activity.GetProperty()); i++ {
				itemAwareElements = append(itemAwareElements, activity.GetProperty()[i])
			}

			// The io specification.
			if activity.GetIoSpecification() != nil {
				for i := 0; i < len(activity.GetIoSpecification().GetDataOutput()); i++ {
					itemAwareElements = append(itemAwareElements, activity.GetIoSpecification().GetDataOutput()[i])
				}

				for i := 0; i < len(activity.GetIoSpecification().GetDataInput()); i++ {
					itemAwareElements = append(itemAwareElements, activity.GetIoSpecification().GetDataInput()[i])
				}
			}
		case BPMN20.Event:
			event := v
			// The property...
			for i := 0; i < len(event.GetProperty()); i++ {
				itemAwareElements = append(itemAwareElements, event.GetProperty()[i])
			}

			switch v := bpmnElement.(type) {
			case BPMN20.CatchEvent:
				catchEvent := v
				for i := 0; i < len(catchEvent.GetDataOutput()); i++ {
					itemAwareElements = append(itemAwareElements, catchEvent.GetDataOutput()[i])
				}
			case BPMN20.ThrowEvent:
				throwEvent := v
				for i := 0; i < len(throwEvent.GetDataInput()); i++ {
					itemAwareElements = append(itemAwareElements, throwEvent.GetDataInput()[i])
				}
			}
		}

		for i := 0; i < len(itemAwareElements); i++ {
			ref := itemAwareElements[i]
			if !this.containItemAwareElementInstance(ref.(BPMN20.BaseElement).GetId(), instance) {
				// Create new source element in case it not exist...
				itemawareElement := new(BPMS.ItemAwareElementInstance)
				itemawareElement.UUID = "BPMS.ItemAwareElementInstance%" + Utility.RandomUUID()
				itemawareElement.M_bpmnElementId = ref.(BPMN20.BaseElement).GetUUID()
				itemawareElement.M_id = ref.(BPMN20.BaseElement).GetId()
				itemawareElement.NeedSave = true
				itemawareElement.SetParentPtr(instance)
				instance.SetData(itemawareElement)
				instance.SetData(itemawareElement)
				// save the instance.
				this.saveInstance(instance)
			}
		}
	}
}

/**
 * Set data associations
 * source represent the instance where the data to be are.
 * target represent the instance where the data must be created.
 * dataAssociations contain the list of dataAssociations to set.
 */
func (this *WorkflowProcessor) setDataAssocication(source BPMS.Instance, target BPMS.Instance, dataAssociations []BPMN20.DataAssociation, sessionId string) {
	// Index the item aware in a map to access it by theire id's
	instances := make(map[string]*BPMS.ItemAwareElementInstance, 0)

	// I will found the definition from the source.
	var definitions *BPMN20.Definitions
	switch v := source.(type) {
	case BPMS.FlowNodeInstance:
		processEntity, _ := GetServer().GetEntityManager().getEntityByUuid(v.GetProcessInstancePtr().GetBpmnElementId())
		definitions = processEntity.GetObject().(*BPMN20.Process).GetDefinitionsPtr()
	case *BPMS.ProcessInstance:
		processEntity, _ := GetServer().GetEntityManager().getEntityByUuid(v.GetBpmnElementId())
		definitions = processEntity.GetObject().(*BPMN20.Process).GetDefinitionsPtr()
	}
	var err error
	var definitionInstances []BPMS.Instance

	// append defintion instance item aware... (DataStore)
	definitionInstances, _ = GetServer().GetWorkflowManager().getInstances(definitions.UUID, "BPMS.DefinitionsInstance")
	for i := 0; i < len(definitionInstances); i++ {
		this.appendItemAwareElementInstance(definitionInstances[i], instances)
	}

	// Apppend the source and the target
	this.appendItemAwareElementInstance(source, instances)
	this.appendItemAwareElementInstance(target, instances)

	// Set the data from the source instance to the
	// target instance.
	for i := 0; i < len(dataAssociations); i++ {
		dataAssociation := dataAssociations[i]
		sourceRefs := dataAssociation.GetSourceRef()
		targetRef := dataAssociation.GetTargetRef()
		targetId := this.getItemAwareElementId(targetRef)

		// Get transtformamtion if there one.
		transformation := dataAssociation.GetTransformation()

		// Get assignement (to, from)
		assignments := dataAssociation.GetAssignment()
		trgItemawareElement := instances[targetId]

		// The 'from' assignements evaluations...
		for j := 0; j < len(assignments); j++ {
			assignment := assignments[j]
			if assignment.GetFrom() != nil {
				from := assignment.GetFrom()
				// In that case no value are contain in the item aware
				// element...
				if len(from.GetOther().(string)) > 0 {
					script := "function setFrom(){\n" + from.GetOther().(string) + "\n}"
					instances, err = this.runScript(definitions, script, "setFrom", instances, sessionId)
					if err != nil {
						// TODO log error here.
						log.Println("--> 992 script error ", err)
					}
				}
			}
		}

		// Here I will iterate over the list of sources, create it itemaware
		// elements if there's not exist and apply transformation on it.
		for j := 0; j < len(sourceRefs); j++ {
			sourceRef := sourceRefs[j]
			sourceId := this.getItemAwareElementId(sourceRef)
			srcItemawareElement := instances[sourceId]

			//  The 'transformation' evaluations...
			if transformation != nil {
				// The value display in the javasript side are the value of the M_data...
				if len(transformation.GetOther().(string)) > 0 {
					script := "function transform(){\n" + transformation.GetOther().(string) + "\n}"
					instances, err = this.runScript(definitions, script, "transform", instances, sessionId)
					if err != nil {
						// TODO log error here.
						log.Println("--> script error 1025 ", err)
					}
				}
			} else {
				// Set the data directly...
				trgItemawareElement.SetData(srcItemawareElement.M_data)
			}
		}

		// The 'to' assignements evaluations...
		for j := 0; j < len(assignments); j++ {
			assignment := assignments[j]
			if assignment.GetTo() != nil {
				to := assignment.GetTo()
				// In that case no value are contain in the item aware
				// element...
				if len(to.GetOther().(string)) > 0 {
					script := "function setTo(){\n" + to.GetOther().(string) + "\n}"
					instances, err = this.runScript(definitions, script, "setTo", instances, sessionId)
					if err != nil {
						log.Println("----->1047 data error: ", err.Error())
					}
				}
			}
		}
	}
}

/**
 * That function is use to set the
 */
func (this *WorkflowProcessor) setItemawareElementData(instance *BPMS.ItemAwareElementInstance, data *otto.Value) error {
	var strVal string

	bpmnElementEntity, _ := GetServer().GetEntityManager().getEntityByUuid(instance.M_bpmnElementId)
	itemAwareElement := bpmnElementEntity.GetObject().(BPMN20.ItemAwareElement)
	//itemawareElementData, _, _ := this.getItemawareElementData(instance)

	// get general information.
	typeName, isCollection := this.getItemawareElementItemDefinitionType(itemAwareElement)

	if strings.HasPrefix(typeName, "xsd:") == true {
		// Here I have a base type
		if isCollection == true {
			object, err := data.Export()
			if err == nil {
				objectStr, err := json.Marshal(object)
				if err == nil {
					strVal = string(objectStr)
				} else {
					return err
				}

			} else {
				return err
			}
		} else {
			// Here I got a simple value...
			if typeName == "xsd:boolean" {
				val, err := data.ToBoolean()
				if err == nil {
					strVal = strconv.FormatBool(val)
				}
			} else if typeName == "xsd:anyURI" || typeName == "xsd:base61Binary" || typeName == "xsd:string" || typeName == "xsd:byte" {
				val, err := data.ToString()
				if err == nil {
					strVal = val
				}
			} else if typeName == "xsd:int" || typeName == "xsd:integer" || typeName == "xsd:long" {
				val, err := data.ToInteger()
				if err == nil {
					strVal = strconv.FormatInt(val, 10)
				}
			} else if typeName == "xsd:date" {
				// TODO implement it...

			} else if typeName == "xsd:float" || typeName == "xsd:double" {
				val, err := data.ToFloat()
				if err == nil {
					strVal = strconv.FormatFloat(val, 'E', -1, 64)
				}
			}
		}
	} else {
		object, err := data.Export()
		if err == nil {
			objectStr, err := json.Marshal(object)
			if err == nil {
				strVal = string(objectStr)
			} else {
				return err
			}

		} else {
			return err
		}

	}

	// Set back the value here.
	instance.SetData([]byte(b64.StdEncoding.EncodeToString([]byte(strVal))))

	// Save it new value.
	instance.GetParentPtr().SetData(instance) // create or update the data.
	this.saveInstance(instance.GetParentPtr())

	return nil
}

/**
 * That function is use to recreate the information from itemawre element data.
 */
func (this *WorkflowProcessor) getItemawareElementData(itemawareElement *BPMS.ItemAwareElementInstance) ([]interface{}, bool, error) {

	objects := make([]interface{}, 0)

	// Double encoding here...
	input := Utility.BytesToString(itemawareElement.M_data)
	input = strings.Replace(input, `"`, "", -1)

	data, err := b64.StdEncoding.DecodeString(input)

	if err != nil {
		return nil, false, err
	}

	// So  now I will get the item definition...
	bpmnElementEntity, _ := GetServer().GetEntityManager().getEntityByUuid(itemawareElement.M_bpmnElementId)
	itemAwareElement := bpmnElementEntity.GetObject().(BPMN20.ItemAwareElement)

	// The item definition object can be either an itemDefinition objet or a string.
	typeName, isCollection := this.getItemawareElementItemDefinitionType(itemAwareElement)

	if strings.HasPrefix(typeName, "xsd:") {
		// Here I have a base type
		if isCollection {
			dataStr := strings.TrimSpace(string(data))
			// In case of a collection.
			if len(dataStr) == 0 {
				// empty array...
				return objects, true, nil
			}

			if !strings.HasPrefix(dataStr, "[") && !strings.HasSuffix(dataStr, "[") {
				dataStr = "[" + dataStr + "]"
			}

			// Unmarshal the data values.
			err := json.Unmarshal([]byte(dataStr), &objects)
			if err == nil {
				return objects, true, nil
			} else {
				return nil, true, err
			}

		} else {
			// Here I got a json structure...
			if typeName == "xsd:boolean" {
				val, err := strconv.ParseBool(string(data))
				if err != nil {
					return nil, false, err
				}
				objects = append(objects, val)
				return objects, false, nil
			} else if typeName == "xsd:anyURI" || typeName == "xsd:base61Binary" || typeName == "xsd:string" {
				objects = append(objects, string(data))
				return objects, false, nil
			} else if typeName == "xsd:int" || typeName == "xsd:integer" || typeName == "xsd:long" {
				val, err := strconv.ParseInt(string(data), 10, 64)
				if err != nil {
					return nil, false, err
				}
				objects = append(objects, val)
				return objects, false, nil
			} else if typeName == "xsd:byte" {
				objects = append(objects, data)
				return objects, false, nil

			} else if typeName == "xsd:date" {
				// TODO implement it...

			} else if typeName == "xsd:float" || typeName == "xsd:double" {
				val, err := strconv.ParseFloat(string(data), 64)
				if err != nil {
					return nil, false, err
				}
				objects = append(objects, val)
				return objects, false, nil
			}
		}
	} else {
		if isCollection {
			values := make([]string, 0)
			err := json.Unmarshal(data, &values)
			// Now each element of the values array contain a json string.
			if err == nil {
				for i := 0; i < len(values); i++ {
					var object interface{}
					err := json.Unmarshal([]byte(values[i]), &object)
					if err == nil {
						// Here I got a json object.
						objects = append(objects, object)
					} else {
						if Utility.IsValidEntityReferenceName(values[i]) {

							// Here I have a reference value.
							entity, errObj := GetServer().GetEntityManager().getEntityByUuid(values[i])
							if errObj == nil {
								// Append the object...
								objects = append(objects, entity.GetObject())
							} else {
								// TODO log error here.
							}

						} else {
							return nil, false, err
						}
					}
				}
			} else {
				return nil, false, err
			}
		} else {
			var object interface{}
			err := json.Unmarshal(data, &object)
			if err == nil {
				objects = append(objects, object)
			}
		}
	}
	return objects, isCollection, nil
}

//////////////////////////////////////////////////////////////////////////////
// Ressource evaluation function.
//////////////////////////////////////////////////////////////////////////////

/**
 * Evaluate ressource of an instance, instance must be subProcess or activity.
 */
func (this *WorkflowProcessor) evaluateRessources(instance BPMS.Instance) {

}

////////////////////////////////////////////////////////////////////////////////
// Getter/setter
////////////////////////////////////////////////////////////////////////////////

/**
 * Get the definition instance with a given id.
 */
func (this *WorkflowProcessor) getDefinitionInstance(definitionsId string, bpmnDefinitionsId string) *BPMS.DefinitionsInstance {
	var definitionsInstance *BPMS.DefinitionsInstance
	// Now I will try to find the definitions instance for that definitions.
	definitionsInstanceEntity, err := GetServer().GetEntityManager().getEntityById("BPMS.DefinitionsInstance", definitionsId)

	if err != nil {
		// Here i will create the definitions instance.
		definitionsInstanceEntity = GetServer().GetEntityManager().NewBPMSDefinitionsInstanceEntity(definitionsId, nil)
		definitionsInstance = definitionsInstanceEntity.GetObject().(*BPMS.DefinitionsInstance)
		definitionsInstance.M_id = definitionsId
		definitionsInstance.M_bpmnElementId = bpmnDefinitionsId
		// Create dataStore object.
		this.createItemAwareElementInstances(definitionsInstance)
	} else {
		definitionsInstance = definitionsInstanceEntity.GetObject().(*BPMS.DefinitionsInstance)
	}

	return definitionsInstance
}

/**
 * That function return the list of active process instances for a given process
 * It also delete innactive instances en passant.
 */
func (this *WorkflowProcessor) getActiveProcessInstances(process *BPMN20.Process) []*BPMS.ProcessInstance {
	instances, _ := GetServer().GetWorkflowManager().getInstances(process.GetUUID(), "BPMS.ProcessInstance")
	var activeProcessInstances []*BPMS.ProcessInstance

	// The list of processes instance...
	for i := 0; i < len(instances); i++ {
		processInstance := instances[i].(*BPMS.ProcessInstance)
		if this.stillActive(processInstance) == false {
			this.deleteInstance(processInstance)
			this.availableTagNumber = append(this.availableTagNumber, processInstance.M_colorNumber)
		} else {
			activeProcessInstances = append(activeProcessInstances, processInstance)

			if tag, ok := this.colorTagMap[processInstance.M_colorNumber]; ok {
				// Set the color...
				tag.InstanceId = processInstance.UUID
				var availableTagNumber_ []string
				for j := 0; j < len(this.availableTagNumber); j++ {
					if this.availableTagNumber[j] != processInstance.M_colorNumber {
						availableTagNumber_ = append(availableTagNumber_, this.availableTagNumber[j])
					}
				}
				this.availableTagNumber = availableTagNumber_
			}
		}
	}

	return activeProcessInstances
}

/**
 * That function return the next process number...
 */
func (this *WorkflowProcessor) getNextProcessInstanceNumber() int {
	var number int

	// Now I will get all defintions names...
	var intancesQuery EntityQuery
	intancesQuery.TypeName = "BPMS.ProcessInstance"

	intancesQuery.Fields = append(intancesQuery.Fields, "M_number")

	var filedsType []interface{} // not use...
	var params []interface{}
	query, _ := json.Marshal(intancesQuery)

	values, err := GetServer().GetDataManager().readData(BPMSDB, string(query), filedsType, params)

	if err == nil {
		for i := 0; i < len(values); i++ {
			n := values[i][0].(int)
			if n > number {
				number = n
			}
		}
	} else {
		log.Println("-------> process number error: ", err)
	}

	// The next number...
	number += 1
	return number
}

/**
 * Utility function to retreive the item definitions.
 */
func (this *WorkflowProcessor) getItemawareElementItemDefinitionType(itemawareElement BPMN20.ItemAwareElement) (string, bool) {
	// Can be one of those class...
	var typeName string
	var isCollection bool

	if itemawareElement.GetItemSubjectRef() != nil {
		typeName = itemawareElement.GetItemSubjectRef().M_structureRef
		isCollection = itemawareElement.GetItemSubjectRef().IsCollection()
	} else if reflect.TypeOf(itemawareElement).String() == "*BPMN20.ItemAwareElement_impl" {
		typeName = itemawareElement.(*BPMN20.ItemAwareElement_impl).M_itemSubjectRef
	} else if reflect.TypeOf(itemawareElement).String() == "*BPMN20.Property" {
		typeName = itemawareElement.(*BPMN20.Property).M_itemSubjectRef
	} else if reflect.TypeOf(itemawareElement).String() == "*BPMN20.DataOutput" {
		typeName = itemawareElement.(*BPMN20.DataOutput).M_itemSubjectRef
		isCollection = itemawareElement.(*BPMN20.DataOutput).IsCollection()
	} else if reflect.TypeOf(itemawareElement).String() == "*BPMN20.DataInput" {
		typeName = itemawareElement.(*BPMN20.DataInput).M_itemSubjectRef
		isCollection = itemawareElement.(*BPMN20.DataInput).M_isCollection
	} else if reflect.TypeOf(itemawareElement).String() == "*BPMN20.DataObject" {
		typeName = itemawareElement.(*BPMN20.DataObject).M_itemSubjectRef
		isCollection = itemawareElement.(*BPMN20.DataObject).M_isCollection
	} else if reflect.TypeOf(itemawareElement).String() == "*BPMN20.DataStore" {
		typeName = itemawareElement.(*BPMN20.DataStore).M_itemSubjectRef
		isCollection = true
	}

	return typeName, isCollection
}

/**
 * Set log information for a given evenement.
 */
func (this *WorkflowProcessor) setLogInfo(instance BPMS.FlowNodeInstance, descripion string, sessionId string) {

	// Now the log information...
	logInfo := new(BPMS.LogInfo)
	logInfo.UUID = "BPMS.LogInfo%" + Utility.RandomUUID()
	logInfo.SetRuntimesPtr(this.runtime.object)
	logInfo.M_date = Utility.MakeTimestamp()
	logInfo.M_id = Utility.RandomUUID()
	logInfo.SetObject(instance)

	if instance.GetFlowNodeType() == BPMS.FlowNodeType_AbstractTask {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_ServiceTask {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_ManualTask {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_BusinessRuleTask {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_ScriptTask {
		logInfo.M_action = "Run Script" // Start a new process...
	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_EmbeddedSubprocess {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_EventSubprocess {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_AdHocSubprocess {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_Transaction {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_CallActivity {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_ParallelGateway {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_ExclusiveGateway {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_EventBasedGateway {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_ComplexGateway {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_StartEvent {
		logInfo.M_action = "Start Process" // Start a new process...
	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_IntermediateCatchEvent {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_BoundaryEvent {

	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_EndEvent {
		logInfo.M_action = "End Process" // Start a new process...
	} else if instance.GetFlowNodeType() == BPMS.FlowNodeType_IntermediateThrowEvent {

	}

	if len(sessionId) > 0 {
		// Set the actor...
		session := GetServer().GetSessionManager().GetActiveSessionById(sessionId)
		// Set the account pointer to the session...
		logInfo.SetActor(session.GetAccountPtr().GetUUID())
	} else {
		// The runtime is the actor in that case.
		logInfo.SetActor(this.runtime.GetUuid())
	}

	// Set as a reference
	instance.(BPMS.Instance).SetLogInfoRef(logInfo)

	// The parent of the logInfo...
	this.runtime.object.SetLogInfos(logInfo)

}

////////////////////////////////////////////////////////////////////////////////
// Api.
////////////////////////////////////////////////////////////////////////////////

/**
 * That function create a new Item aware element instance for a given bpmn element
 * id, who is an instance of ItemAwareElement.
 */
func (this *WorkflowProcessor) NewItemAwareElementInstance(bpmnElementId string, data interface{}, messageId string, sessionId string) *BPMS.ItemAwareElementInstance {
	instance, err := this.createItemAwareElementInstance(bpmnElementId, data)

	// Retunr an error in that case...
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, err)
		return nil
	}

	return instance
}

/**
 * Return the list of active process instances for a given bpmn process.
 */
func (this *WorkflowProcessor) GetActiveProcessInstances(bpmnElementId string, messageId string, sessionId string) []*BPMS.ProcessInstance {

	processEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(bpmnElementId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return nil
	}

	instances := this.getActiveProcessInstances(processEntity.GetObject().(*BPMN20.Process))
	log.Println("--------> get active instances: ", instances)
	return instances
}

/**
 * Execute an activity instance with a given uuid.
 */
func (this *WorkflowProcessor) ActivateActivityInstance(uuid string, messageId string, sessionId string) {
	log.Println("----------> execute activity ", uuid)
	entity, errObj := GetServer().GetEntityManager().getEntityByUuid(uuid)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	// I will get the instance and activate it...
	instance := entity.GetObject().(*BPMS.ActivityInstance)

	// Activate the user task.
	instance.SetLifecycleState(BPMS.LifecycleState_Active)
	this.saveInstance(instance)
}
