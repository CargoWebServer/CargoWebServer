// +build BPMN20

package Server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"code.google.com/p/go-uuid/uuid"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/BPMN20"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/BPMS"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
)

const (
	// Datastore names...
	BPMN20DB = "BPMN20"
	BPMNDIDB = "BPMN20"
	DCDB     = "BPMN20"
	DIDB     = "BPMN20"
)

type WorkflowManager struct {
}

func newWorkflowManager() *WorkflowManager {
	// The workflow manger...
	workflowManager := new(WorkflowManager)

	// The BPMN 2.0 entities
	bpmn20DB := new(Config.DataStoreConfiguration)
	bpmn20DB.M_id = BPMN20DB
	bpmn20DB.M_dataStoreVendor = Config.DataStoreVendor_MYCELIUS
	bpmn20DB.M_dataStoreType = Config.DataStoreType_KEY_VALUE_STORE
	bpmn20DB.NeedSave = true
	GetServer().GetConfigurationManager().appendDefaultDataStoreConfiguration(bpmn20DB)
	GetServer().GetDataManager().appendDefaultDataStore(bpmn20DB)

	// Create prototypes...
	GetServer().GetEntityManager().createBPMN20Prototypes()
	GetServer().GetEntityManager().createBPMNDIPrototypes()
	GetServer().GetEntityManager().createDCPrototypes()
	GetServer().GetEntityManager().createDIPrototypes()

	// Register objects...
	GetServer().GetEntityManager().registerBPMNDIObjects()
	GetServer().GetEntityManager().registerDCObjects()
	GetServer().GetEntityManager().registerDIObjects()
	GetServer().GetEntityManager().registerBPMN20Objects()

	return workflowManager
}

var workflowManager *WorkflowManager

func (this *Server) GetWorkflowManager() *WorkflowManager {
	if workflowManager == nil {
		workflowManager = newWorkflowManager()
	}
	return workflowManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Do intialysation stuff here.
 */
func (this *WorkflowManager) initialize() {
	log.Println("--> Initialize WorkflowManager")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())

	// Here i will load the list of all definition into the
	// entity mananager entity map.
	GetServer().GetEntityManager().getEntitiesByType("BPMN20.Definitions", "", "BPMN20")

}

func (this *WorkflowManager) getId() string {
	return "WorkflowManager"
}

/**
 * Start the service.
 */
func (this *WorkflowManager) start() {
	log.Println("--> Start WorkflowManager")
}

func (this *WorkflowManager) stop() {
	log.Println("--> Stop WorkflowManager")

}

////////////////////////////////////////////////////////////////////////////////
// BPMN private functions...
////////////////////////////////////////////////////////////////////////////////
/**
 * That function return the list of definitions id on the server.
 */
func (this *WorkflowManager) getDefinitionsIds() ([]string, *CargoEntities.Error) {
	names := make([]string, 0)

	// Now I will get all defintions names...
	var definitionsQuery EntityQuery
	definitionsQuery.TypeName = "BPMN20.Definitions"
	definitionsQuery.Fields = append(definitionsQuery.Fields, "M_id")

	var filedsType []interface{} // not use...
	var params []interface{}
	query, _ := json.Marshal(definitionsQuery)

	values, err := GetServer().GetDataManager().readData(BPMN20DB, string(query), filedsType, params)

	if err == nil {
		for i := 0; i < len(values); i++ {
			names = append(names, values[i][0].(string))
		}
	} else {
		cargoError := NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("Definitions Ids cant be found!"))
		return names, cargoError
	}

	return names, nil
}

/**
 * That function return the defintion with a given id.
 */
func (this *WorkflowManager) getDefinitionsById(id string) (*BPMN20.Definitions, *CargoEntities.Error) {

	/** Create a new entity with a given name **/
	definitionEntity, err := GetServer().GetEntityManager().getEntityById("BPMN20.Definitions", id)

	// Test only...
	if err != nil {
		return nil, err
	}

	// Initialyse the entity...

	// Here the defnition need to be initialyse...
	GetServer().GetEntityManager().InitEntity(definitionEntity)

	// Get the object...
	definition := definitionEntity.GetObject().(*BPMN20.Definitions)

	// Return the initialyse entity...
	return definition, nil
}

////////////////////////////////////////////////////////////////////////////////
// utility function section
////////////////////////////////////////////////////////////////////////////////

/**
 * That function return the process definition entity from a given definition.
 */
func (this *WorkflowManager) getProcessEntityById(processId string, defintionId string) (*BPMN20_ProcessEntity, *CargoEntities.Error) {

	defintions, err := this.getDefinitionsById(defintionId)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(defintions.GetRootElement()); i++ {
		element := defintions.GetRootElement()[i].(BPMN20.BaseElement)
		//log.Println("Element id ", element.GetId())
		if element.GetId() == processId {
			processEntity := GetServer().GetEntityManager().NewBPMN20ProcessEntity(element.GetId(), nil)
			// Initialyse the reference...
			GetServer().GetEntityManager().InitEntity(processEntity)
			return processEntity, nil
		}
	}
	return nil, NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("No procees with id"+processId+"was found in definition "+defintionId))
}

////////////////////////////////////////////////////////////////////////////////
// universe...
////////////////////////////////////////////////////////////////////////////////
func (this *WorkflowManager) getProcess() []*BPMN20.Process {
	entities, _ := GetServer().GetEntityManager().getEntitiesByType("BPMN20.Process", "", "BPMN20")
	var processes []*BPMN20.Process
	for i := 0; i < len(entities); i++ {
		processes = append(processes, entities[i].GetObject().(*BPMN20.Process))
	}
	return processes
}

/**
 * That function return the list of instance for a given bpmn element id.
 */
func (this *WorkflowManager) getInstances(bpmnElementId string, typeName string) ([]BPMS.Instance, *CargoEntities.Error) {

	// Now I will get all defintions names...
	var intancesQuery EntityQuery
	intancesQuery.TypeName = typeName
	intancesQuery.Fields = append(intancesQuery.Fields, "uuid")
	intancesQuery.Indexs = append(intancesQuery.Indexs, "M_bpmnElementId="+bpmnElementId)

	var filedsType []interface{} // not use...
	var params []interface{}
	query, _ := json.Marshal(intancesQuery)

	// The array of instance...
	var instances []BPMS.Instance

	values, err := GetServer().GetDataManager().readData(BPMSDB, string(query), filedsType, params)
	if err == nil {
		for i := 0; i < len(values); i++ {
			uuid := values[i][0].(string)
			instance, err := GetServer().GetEntityManager().getEntityByUuid(uuid)
			// Initialyse the reference...
			GetServer().GetEntityManager().InitEntity(instance)
			if err != nil {
				return instances, err
			}
			// Append to the list...
			instances = append(instances, instance.GetObject().(BPMS.Instance))
		}
	} else {
		return instances, NewError(Utility.FileLine(), ENTITY_ID_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("No instance was found with id "+bpmnElementId+" and type name"+typeName))
	}

	return instances, nil

}

/**
 * Load the content of the bpmn definition...
 */
func (this *WorkflowManager) loadXmlBpmnDefinitions(messageId string, sessionId string) {

	dump := GetServer().GetConfigurationManager().GetDefinitionsPath()
	files, _ := ioutil.ReadDir(dump)
	if len(files) > 0 {
		log.Println("Import file in dir", dump)
	}

	// The files
	for _, f := range files {
		log.Println("Load file ", dump+"/"+f.Name())
		definition := new(BPMN20.Definitions)
		xmlFactory := new(Entities.BPMN_XmlFactory)
		err := xmlFactory.InitXml(dump+"/"+f.Name(), definition)
		if err != nil {

			GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), XML_READ_ERROR, SERVER_ERROR_CODE, err))
			os.Remove(f.Name())
			return
		} else {
			definitionEntity := new(BPMN20_DefinitionsEntity)
			definitionEntity.object = definition
			if len(definitionEntity.uuid) == 0 {
				definitionEntity.uuid = definitionEntity.GetTypeName() + "%" + uuid.NewRandom().String()
			}

			definitionEntity.SaveEntity()
			os.Remove(f.Name())
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////

/**
 * Load the bpmn defintions from the Definitions directory...
 */
func (this *WorkflowManager) ImportXmlBpmnDefinitions(content string, messageId string, sessionId string) {
	definitions := new(BPMN20.Definitions)
	xmlFactory := new(Entities.BPMN_XmlFactory)

	// Here I will create a temporary file
	tmp := GetServer().GetConfigurationManager().GetTmpPath()
	f, err := os.Create(tmp + "/" + Utility.RandomUUID())

	f.WriteString(content)
	f.Close()

	// Remove the file when done.
	defer os.Remove(f.Name())

	err = xmlFactory.InitXml(f.Name(), definitions)

	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), XML_READ_ERROR, SERVER_ERROR_CODE, err))
		f.Close()
		return
	} else {
		definitionEntity := GetServer().GetEntityManager().NewBPMN20DefinitionsEntityFromObject(definitions)
		log.Println(definitionEntity)
		definitionEntity.SaveEntity()
		f.Close()

		// Here a will create the event...
		eventData := make([]*MessageData, 1)

		// The file itself
		definitionsInfo := new(MessageData)
		definitionsInfo.Name = "definitionsInfo"
		definitionsInfo.Value = definitionEntity.object
		eventData[0] = definitionsInfo

		evt, _ := NewEvent(NewDefinitionsEvent, BpmnEvent, eventData)
		// Send the message to the remaining users...
		GetServer().GetEventManager().BroadcastEvent(evt)
	}

}

/**
 * Return the list of all definitions ids...
 */
func (this *WorkflowManager) GetDefinitionsIds(sessionId string, messageId string) []string {
	ids, err := this.getDefinitionsIds()
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, err)
	}

	return ids
}

/**
 * Return the list of all definitions.
 */
func (this *WorkflowManager) GetAllDefinitions(messageId string, sessionId string) []*BPMN20.Definitions {
	definitions, err := GetServer().GetEntityManager().getEntitiesByType("BPMN20.Definitions", "", "BPMN20")
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, err)
	}

	allDefinitions := make([]*BPMN20.Definitions, 0)
	for i := 0; i < len(definitions); i++ {
		if definitions[i].GetObject() != nil {
			allDefinitions = append(allDefinitions, definitions[i].GetObject().(*BPMN20.Definitions))
		} else {
			definitions[i].DeleteEntity()
		}
	}

	return allDefinitions
}

/**
 * Return a definition with a given id.
 */
func (this *WorkflowManager) GetDefinitionsById(id string, messageId string, sessionId string) *BPMN20.Definitions {
	definition, err := this.getDefinitionsById(id)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, err)
	}
	return definition
}

////////////////////////////////////////////////////////////////////////////////
// Runtime section
////////////////////////////////////////////////////////////////////////////////

/**
 * That function is use to create a new process. The definitions id and processId
 * must be know...
 * The start event data is the data ouput, those data will be attach to an
 * item aware element instance as a []uint8.
 * The start event definition data will be attach to the event definition,
 * as message data or signal data...
 */
func (this *WorkflowManager) StartProcess(processUUID string, startEventData interface{}, startEventDefinitionData interface{}, messageId string, sessionId string) {
	log.Println("-------> start process: ", processUUID)

	trigger := new(BPMS.Trigger)
	trigger.UUID = "BPMS.Trigger%" + Utility.RandomUUID()
	trigger.M_processUUID = processUUID

	// The event data...
	if reflect.TypeOf(startEventData).String() == "[]*BPMS.ItemAwareElementInstance" {
		for i := 0; i < len(startEventData.([]*BPMS.ItemAwareElementInstance)); i++ {
			trigger.SetDataRef(startEventData.([]*BPMS.ItemAwareElementInstance)[i])
		}
	}

	trigger.M_eventTriggerType = BPMS.EventTriggerType_Start

	trigger.M_sessionId = sessionId

	// So from the start event data i will create the item aware element
	// and set is data...
	GetServer().GetWorkflowProcessor().triggerChan <- trigger
}

/**
 * That function return the list of instance for a given bpmn element id.
 */
func (this *WorkflowManager) GetDefinitionInstances(id string, messageId string, sessionId string) []*BPMS.DefinitionsInstance {
	instances, err := this.getInstances(id, "BPMS.DefinitionsInstance")
	if err != nil {
		log.Println("--------> Definitions", id, "not found!!!")
		GetServer().reportErrorMessage(messageId, sessionId, err)
	}

	definitionInstances := make([]*BPMS.DefinitionsInstance, len(instances))
	for i := 0; i < len(instances); i++ {
		definitionInstances[i] = instances[i].(*BPMS.DefinitionsInstance)
	}

	return definitionInstances
}
