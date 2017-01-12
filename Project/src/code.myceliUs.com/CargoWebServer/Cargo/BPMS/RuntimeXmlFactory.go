package BPMS
import (
	"golang.org/x/net/html/charset"
	"encoding/xml"
	"strings"
	"os"
	"log"
	"path/filepath"
	"code.google.com/p/go-uuid/uuid"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/BPMS_Runtime"
)
type RuntimeXmlFactory struct {
	m_references map[string] interface{}
	m_object map[string]map[string][]string
}

/** Initialization function from xml file **/
func (this *RuntimeXmlFactory)InitXml(inputPath string, object *BPMS_Runtime.Runtimes) error{
	this.m_references = make(map[string]interface{})
	this.m_object = make(map[string]map[string][]string)
	xmlFilePath, err := filepath.Abs(inputPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	reader, err := os.Open(xmlFilePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	var xmlElement *BPMS_Runtime.XsdRuntimes
	xmlElement = new(BPMS_Runtime.XsdRuntimes)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(xmlElement); err != nil {
		return err
	}
	this.InitRuntimes(xmlElement, object)
	for ref0, refMap := range this.m_object {
		refOwner := this.m_references[ref0]
		if refOwner != nil {
			for ref1, _ := range refMap {
				refs := refMap[ref1]
				for i:=0; i<len(refs); i++{
					ref:= this.m_references[refs[i]]
					if  ref != nil {
						params := make([]interface {},0)
						params = append(params,ref)
						methodName := "Set" + strings.ToUpper(ref1[0:1]) + ref1[1:]
						Utility.CallMethod(refOwner, methodName, params )
					}else{
						params := make([]interface {},0)
						params = append(params,refs[i])
						methodName := "Set" + strings.ToUpper(ref1[0:1]) + ref1[1:]
						Utility.CallMethod(refOwner, methodName, params)
					}
				}
			}
		}
	}
	return nil
}

/** Serialization to xml file **/
func (this *RuntimeXmlFactory)SerializeXml(outputPath string, toSerialize *BPMS_Runtime.Runtimes) error{
	xmlFilePath, err := filepath.Abs(outputPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fo, err := os.Create(xmlFilePath)
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	var xmlElement *BPMS_Runtime.XsdRuntimes
	xmlElement = new(BPMS_Runtime.XsdRuntimes)

	this.SerialyzeRuntimes(xmlElement, toSerialize)
	output, err := xml.MarshalIndent(xmlElement, "  ", "    ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fileContent := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n")
	fileContent = append(fileContent, output...)
	_, err = fo.Write(fileContent)
	return nil
}

/** inititialisation of GatewayInstance **/
func (this *RuntimeXmlFactory) InitGatewayInstance(xmlElement *BPMS_Runtime.XsdGatewayInstance,object *BPMS_Runtime.GatewayInstance){
	log.Println("Initialize GatewayInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.GatewayInstance%" + Utility.RandomUUID()	}

	/** GatewayInstance **/
	object.M_id= xmlElement.M_id

	/** GatewayInstance **/
	object.M_bpmnElementId= xmlElement.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref inputRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_inputRef); i++ {
		if _, ok:= this.m_object[object.M_id]["inputRef"]; !ok {
			this.m_object[object.M_id]["inputRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["inputRef"] = append(this.m_object[object.M_id]["inputRef"], xmlElement.M_inputRef[i])
	}


	/** Init ref outputRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_outputRef); i++ {
		if _, ok:= this.m_object[object.M_id]["outputRef"]; !ok {
			this.m_object[object.M_id]["outputRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["outputRef"] = append(this.m_object[object.M_id]["outputRef"], xmlElement.M_outputRef[i])
	}


	/** FlowNodeType **/
	if xmlElement.M_flowNodeType=="##AbstractTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AbstractTask
	} else if xmlElement.M_flowNodeType=="##ServiceTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ServiceTask
	} else if xmlElement.M_flowNodeType=="##UserTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_UserTask
	} else if xmlElement.M_flowNodeType=="##ManualTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ManualTask
	} else if xmlElement.M_flowNodeType=="##BusinessRuleTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BusinessRuleTask
	} else if xmlElement.M_flowNodeType=="##ScriptTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ScriptTask
	} else if xmlElement.M_flowNodeType=="##EmbeddedSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EmbeddedSubprocess
	} else if xmlElement.M_flowNodeType=="##EventSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventSubprocess
	} else if xmlElement.M_flowNodeType=="##AdHocSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AdHocSubprocess
	} else if xmlElement.M_flowNodeType=="##Transaction"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_Transaction
	} else if xmlElement.M_flowNodeType=="##CallActivity"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_CallActivity
	} else if xmlElement.M_flowNodeType=="##ParallelGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ParallelGateway
	} else if xmlElement.M_flowNodeType=="##ExclusiveGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ExclusiveGateway
	} else if xmlElement.M_flowNodeType=="##InclusiveGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_InclusiveGateway
	} else if xmlElement.M_flowNodeType=="##EventBasedGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventBasedGateway
	} else if xmlElement.M_flowNodeType=="##ComplexGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ComplexGateway
	} else if xmlElement.M_flowNodeType=="##StartEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_StartEvent
	} else if xmlElement.M_flowNodeType=="##IntermediateCatchEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateCatchEvent
	} else if xmlElement.M_flowNodeType=="##BoundaryEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BoundaryEvent
	} else if xmlElement.M_flowNodeType=="##EndEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EndEvent
	} else if xmlElement.M_flowNodeType=="##IntermediateThrowEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateThrowEvent
	}

	/** LifecycleState **/
	if xmlElement.M_lifecycleState=="##Completed"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completed
	} else if xmlElement.M_lifecycleState=="##Compensated"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensated
	} else if xmlElement.M_lifecycleState=="##Failed"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failed
	} else if xmlElement.M_lifecycleState=="##Terminated"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminated
	} else if xmlElement.M_lifecycleState=="##Ready"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Ready
	} else if xmlElement.M_lifecycleState=="##Active"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Active
	} else if xmlElement.M_lifecycleState=="##Completing"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completing
	} else if xmlElement.M_lifecycleState=="##Compensating"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensating
	} else if xmlElement.M_lifecycleState=="##Failing"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failing
	} else if xmlElement.M_lifecycleState=="##Terminating"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminating
	}

	/** GatewayType **/
	if xmlElement.M_gatewayType=="##ParallelGateway"{
		object.M_gatewayType=BPMS_Runtime.GatewayType_ParallelGateway
	} else if xmlElement.M_gatewayType=="##ExclusiveGateway"{
		object.M_gatewayType=BPMS_Runtime.GatewayType_ExclusiveGateway
	} else if xmlElement.M_gatewayType=="##InclusiveGateway"{
		object.M_gatewayType=BPMS_Runtime.GatewayType_InclusiveGateway
	} else if xmlElement.M_gatewayType=="##EventBasedGateway"{
		object.M_gatewayType=BPMS_Runtime.GatewayType_EventBasedGateway
	} else if xmlElement.M_gatewayType=="##ComplexGateway"{
		object.M_gatewayType=BPMS_Runtime.GatewayType_ComplexGateway
	}
}

/** inititialisation of EventDefinitionInstance **/
func (this *RuntimeXmlFactory) InitEventDefinitionInstance(xmlElement *BPMS_Runtime.XsdEventDefinitionInstance,object *BPMS_Runtime.EventDefinitionInstance){
	log.Println("Initialize EventDefinitionInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.EventDefinitionInstance%" + Utility.RandomUUID()	}

	/** EventDefinitionInstance **/
	object.M_id= xmlElement.M_id

	/** EventDefinitionInstance **/
	object.M_bpmnElementId= xmlElement.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** EventDefinitionType **/
	if xmlElement.M_eventDefinitionType=="##MessageEventDefinition"{
		object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_MessageEventDefinition
	} else if xmlElement.M_eventDefinitionType=="##LinkEventDefinition"{
		object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_LinkEventDefinition
	} else if xmlElement.M_eventDefinitionType=="##ErrorEventDefinition"{
		object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_ErrorEventDefinition
	} else if xmlElement.M_eventDefinitionType=="##TerminateEventDefinition"{
		object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_TerminateEventDefinition
	} else if xmlElement.M_eventDefinitionType=="##CompensationEventDefinition"{
		object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_CompensationEventDefinition
	} else if xmlElement.M_eventDefinitionType=="##ConditionalEventDefinition"{
		object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_ConditionalEventDefinition
	} else if xmlElement.M_eventDefinitionType=="##TimerEventDefinition"{
		object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_TimerEventDefinition
	} else if xmlElement.M_eventDefinitionType=="##CancelEventDefinition"{
		object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_CancelEventDefinition
	} else if xmlElement.M_eventDefinitionType=="##EscalationEventDefinition"{
		object.M_eventDefinitionType=BPMS_Runtime.EventDefinitionType_EscalationEventDefinition
	}
}

/** inititialisation of Trigger **/
func (this *RuntimeXmlFactory) InitTrigger(xmlElement *BPMS_Runtime.XsdTrigger,object *BPMS_Runtime.Trigger){
	log.Println("Initialize Trigger")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.Trigger%" + Utility.RandomUUID()	}

	/** Init eventData **/
	object.M_eventDatas= make([]*BPMS_Runtime.EventData,0)
	for i:=0;i<len(xmlElement.M_eventDatas); i++{
		val:=new(BPMS_Runtime.EventData)
		this.InitEventData(xmlElement.M_eventDatas[i],val)
		object.M_eventDatas= append(object.M_eventDatas, val)

		/** association initialisation **/
		val.SetTriggerPtr(object)
	}

	/** Init ref dataRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_dataRef); i++ {
		if _, ok:= this.m_object[object.M_id]["dataRef"]; !ok {
			this.m_object[object.M_id]["dataRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["dataRef"] = append(this.m_object[object.M_id]["dataRef"], xmlElement.M_dataRef[i])
	}


	/** Init ref sourceRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	if xmlElement.M_sourceRef !=nil {
		if _, ok:= this.m_object[object.M_id]["sourceRef"]; !ok {
			this.m_object[object.M_id]["sourceRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["sourceRef"] = append(this.m_object[object.M_id]["sourceRef"], *xmlElement.M_sourceRef)
	}


	/** Init ref targetRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	if xmlElement.M_targetRef !=nil {
		if _, ok:= this.m_object[object.M_id]["targetRef"]; !ok {
			this.m_object[object.M_id]["targetRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["targetRef"] = append(this.m_object[object.M_id]["targetRef"], *xmlElement.M_targetRef)
	}


	/** Trigger **/
	object.M_id= xmlElement.M_id

	/** Trigger **/
	object.M_processUUID= xmlElement.M_processUUID

	/** Trigger **/
	object.M_sessionId= xmlElement.M_sessionId

	/** EventTriggerType **/
	if xmlElement.M_eventTriggerType=="##None"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_None
	} else if xmlElement.M_eventTriggerType=="##Timer"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Timer
	} else if xmlElement.M_eventTriggerType=="##Conditional"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Conditional
	} else if xmlElement.M_eventTriggerType=="##Message"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Message
	} else if xmlElement.M_eventTriggerType=="##Signal"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Signal
	} else if xmlElement.M_eventTriggerType=="##Multiple"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Multiple
	} else if xmlElement.M_eventTriggerType=="##ParallelMultiple"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_ParallelMultiple
	} else if xmlElement.M_eventTriggerType=="##Escalation"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Escalation
	} else if xmlElement.M_eventTriggerType=="##Error"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Error
	} else if xmlElement.M_eventTriggerType=="##Compensation"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Compensation
	} else if xmlElement.M_eventTriggerType=="##Terminate"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Terminate
	} else if xmlElement.M_eventTriggerType=="##Cancel"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Cancel
	} else if xmlElement.M_eventTriggerType=="##Link"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Link
	} else if xmlElement.M_eventTriggerType=="##Start"{
		object.M_eventTriggerType=BPMS_Runtime.EventTriggerType_Start
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ItemAwareElementInstance **/
func (this *RuntimeXmlFactory) InitItemAwareElementInstance(xmlElement *BPMS_Runtime.XsdItemAwareElementInstance,object *BPMS_Runtime.ItemAwareElementInstance){
	log.Println("Initialize ItemAwareElementInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.ItemAwareElementInstance%" + Utility.RandomUUID()	}

	/** ItemAwareElementInstance **/
	object.M_id= xmlElement.M_id

	/** ItemAwareElementInstance **/
	object.M_bpmnElementId= xmlElement.M_bpmnElementId

	/** ItemAwareElementInstance **/
	object.M_data= xmlElement.M_data
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ConnectingObject **/
func (this *RuntimeXmlFactory) InitConnectingObject(xmlElement *BPMS_Runtime.XsdConnectingObject,object *BPMS_Runtime.ConnectingObject){
	log.Println("Initialize ConnectingObject")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.ConnectingObject%" + Utility.RandomUUID()	}

	/** ConnectingObject **/
	object.M_id= xmlElement.M_id

	/** ConnectingObject **/
	object.M_bpmnElementId= xmlElement.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref sourceRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	if xmlElement.M_sourceRef !=nil {
		if _, ok:= this.m_object[object.M_id]["sourceRef"]; !ok {
			this.m_object[object.M_id]["sourceRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["sourceRef"] = append(this.m_object[object.M_id]["sourceRef"], *xmlElement.M_sourceRef)
	}


	/** Init ref targetRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	if xmlElement.M_targetRef !=nil {
		if _, ok:= this.m_object[object.M_id]["targetRef"]; !ok {
			this.m_object[object.M_id]["targetRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["targetRef"] = append(this.m_object[object.M_id]["targetRef"], *xmlElement.M_targetRef)
	}


	/** ConnectingObjectType **/
	if xmlElement.M_connectingObjectType=="##SequenceFlow"{
		object.M_connectingObjectType=BPMS_Runtime.ConnectingObjectType_SequenceFlow
	} else if xmlElement.M_connectingObjectType=="##MessageFlow"{
		object.M_connectingObjectType=BPMS_Runtime.ConnectingObjectType_MessageFlow
	} else if xmlElement.M_connectingObjectType=="##Association"{
		object.M_connectingObjectType=BPMS_Runtime.ConnectingObjectType_Association
	} else if xmlElement.M_connectingObjectType=="##DataAssociation"{
		object.M_connectingObjectType=BPMS_Runtime.ConnectingObjectType_DataAssociation
	}
}

/** inititialisation of CorrelationInfo **/
func (this *RuntimeXmlFactory) InitCorrelationInfo(xmlElement *BPMS_Runtime.XsdCorrelationInfo,object *BPMS_Runtime.CorrelationInfo){
	log.Println("Initialize CorrelationInfo")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.CorrelationInfo%" + Utility.RandomUUID()	}

	/** CorrelationInfo **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of DefinitionsInstance **/
func (this *RuntimeXmlFactory) InitDefinitionsInstance(xmlElement *BPMS_Runtime.XsdDefinitionsInstance,object *BPMS_Runtime.DefinitionsInstance){
	log.Println("Initialize DefinitionsInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.DefinitionsInstance%" + Utility.RandomUUID()	}

	/** DefinitionsInstance **/
	object.M_id= xmlElement.M_id

	/** DefinitionsInstance **/
	object.M_bpmnElementId= xmlElement.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init process **/
	object.M_processInstances= make([]*BPMS_Runtime.ProcessInstance,0)
	for i:=0;i<len(xmlElement.M_processInstances); i++{
		val:=new(BPMS_Runtime.ProcessInstance)
		this.InitProcessInstance(xmlElement.M_processInstances[i],val)
		object.M_processInstances= append(object.M_processInstances, val)

		/** association initialisation **/
	}
}

/** inititialisation of ProcessInstance **/
func (this *RuntimeXmlFactory) InitProcessInstance(xmlElement *BPMS_Runtime.XsdProcessInstance,object *BPMS_Runtime.ProcessInstance){
	log.Println("Initialize ProcessInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.ProcessInstance%" + Utility.RandomUUID()	}

	/** ProcessInstance **/
	object.M_id= xmlElement.M_id

	/** ProcessInstance **/
	object.M_bpmnElementId= xmlElement.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init activityInstance **/
	object.M_flowNodeInstances= make([]BPMS_Runtime.FlowNodeInstance,0)
	for i:=0;i<len(xmlElement.M_flowNodeInstances_0); i++{
		val:=new(BPMS_Runtime.ActivityInstance)
		this.InitActivityInstance(xmlElement.M_flowNodeInstances_0[i],val)
		object.M_flowNodeInstances= append(object.M_flowNodeInstances, val)

		/** association initialisation **/
		val.SetProcessInstancePtr(object)
	}

	/** Init SubprocessInstance **/
	for i:=0;i<len(xmlElement.M_flowNodeInstances_1); i++{
		val:=new(BPMS_Runtime.SubprocessInstance)
		this.InitSubprocessInstance(xmlElement.M_flowNodeInstances_1[i],val)
		object.M_flowNodeInstances= append(object.M_flowNodeInstances, val)

		/** association initialisation **/
		val.SetProcessInstancePtr(object)
	}

	/** Init gatewayInstance **/
	for i:=0;i<len(xmlElement.M_flowNodeInstances_2); i++{
		val:=new(BPMS_Runtime.GatewayInstance)
		this.InitGatewayInstance(xmlElement.M_flowNodeInstances_2[i],val)
		object.M_flowNodeInstances= append(object.M_flowNodeInstances, val)

		/** association initialisation **/
		val.SetProcessInstancePtr(object)
	}

	/** Init eventInstance **/
	for i:=0;i<len(xmlElement.M_flowNodeInstances_3); i++{
		val:=new(BPMS_Runtime.EventInstance)
		this.InitEventInstance(xmlElement.M_flowNodeInstances_3[i],val)
		object.M_flowNodeInstances= append(object.M_flowNodeInstances, val)

		/** association initialisation **/
		val.SetProcessInstancePtr(object)
	}

	/** Instance **/
	object.M_number= xmlElement.M_number

	/** Instance **/
	object.M_colorNumber= xmlElement.M_colorNumber

	/** Instance **/
	object.M_colorName= xmlElement.M_colorName
}

/** inititialisation of ActivityInstance **/
func (this *RuntimeXmlFactory) InitActivityInstance(xmlElement *BPMS_Runtime.XsdActivityInstance,object *BPMS_Runtime.ActivityInstance){
	log.Println("Initialize ActivityInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.ActivityInstance%" + Utility.RandomUUID()	}

	/** ActivityInstance **/
	object.M_id= xmlElement.M_id

	/** ActivityInstance **/
	object.M_bpmnElementId= xmlElement.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref inputRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_inputRef); i++ {
		if _, ok:= this.m_object[object.M_id]["inputRef"]; !ok {
			this.m_object[object.M_id]["inputRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["inputRef"] = append(this.m_object[object.M_id]["inputRef"], xmlElement.M_inputRef[i])
	}


	/** Init ref outputRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_outputRef); i++ {
		if _, ok:= this.m_object[object.M_id]["outputRef"]; !ok {
			this.m_object[object.M_id]["outputRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["outputRef"] = append(this.m_object[object.M_id]["outputRef"], xmlElement.M_outputRef[i])
	}


	/** FlowNodeType **/
	if xmlElement.M_flowNodeType=="##AbstractTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AbstractTask
	} else if xmlElement.M_flowNodeType=="##ServiceTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ServiceTask
	} else if xmlElement.M_flowNodeType=="##UserTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_UserTask
	} else if xmlElement.M_flowNodeType=="##ManualTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ManualTask
	} else if xmlElement.M_flowNodeType=="##BusinessRuleTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BusinessRuleTask
	} else if xmlElement.M_flowNodeType=="##ScriptTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ScriptTask
	} else if xmlElement.M_flowNodeType=="##EmbeddedSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EmbeddedSubprocess
	} else if xmlElement.M_flowNodeType=="##EventSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventSubprocess
	} else if xmlElement.M_flowNodeType=="##AdHocSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AdHocSubprocess
	} else if xmlElement.M_flowNodeType=="##Transaction"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_Transaction
	} else if xmlElement.M_flowNodeType=="##CallActivity"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_CallActivity
	} else if xmlElement.M_flowNodeType=="##ParallelGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ParallelGateway
	} else if xmlElement.M_flowNodeType=="##ExclusiveGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ExclusiveGateway
	} else if xmlElement.M_flowNodeType=="##InclusiveGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_InclusiveGateway
	} else if xmlElement.M_flowNodeType=="##EventBasedGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventBasedGateway
	} else if xmlElement.M_flowNodeType=="##ComplexGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ComplexGateway
	} else if xmlElement.M_flowNodeType=="##StartEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_StartEvent
	} else if xmlElement.M_flowNodeType=="##IntermediateCatchEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateCatchEvent
	} else if xmlElement.M_flowNodeType=="##BoundaryEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BoundaryEvent
	} else if xmlElement.M_flowNodeType=="##EndEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EndEvent
	} else if xmlElement.M_flowNodeType=="##IntermediateThrowEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateThrowEvent
	}

	/** LifecycleState **/
	if xmlElement.M_lifecycleState=="##Completed"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completed
	} else if xmlElement.M_lifecycleState=="##Compensated"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensated
	} else if xmlElement.M_lifecycleState=="##Failed"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failed
	} else if xmlElement.M_lifecycleState=="##Terminated"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminated
	} else if xmlElement.M_lifecycleState=="##Ready"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Ready
	} else if xmlElement.M_lifecycleState=="##Active"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Active
	} else if xmlElement.M_lifecycleState=="##Completing"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completing
	} else if xmlElement.M_lifecycleState=="##Compensating"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensating
	} else if xmlElement.M_lifecycleState=="##Failing"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failing
	} else if xmlElement.M_lifecycleState=="##Terminating"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminating
	}

	/** ActivityType **/
	if xmlElement.M_activityType=="##AbstractTask"{
		object.M_activityType=BPMS_Runtime.ActivityType_AbstractTask
	} else if xmlElement.M_activityType=="##ServiceTask"{
		object.M_activityType=BPMS_Runtime.ActivityType_ServiceTask
	} else if xmlElement.M_activityType=="##UserTask"{
		object.M_activityType=BPMS_Runtime.ActivityType_UserTask
	} else if xmlElement.M_activityType=="##ManualTask"{
		object.M_activityType=BPMS_Runtime.ActivityType_ManualTask
	} else if xmlElement.M_activityType=="##BusinessRuleTask"{
		object.M_activityType=BPMS_Runtime.ActivityType_BusinessRuleTask
	} else if xmlElement.M_activityType=="##ScriptTask"{
		object.M_activityType=BPMS_Runtime.ActivityType_ScriptTask
	} else if xmlElement.M_activityType=="##EmbeddedSubprocess"{
		object.M_activityType=BPMS_Runtime.ActivityType_EmbeddedSubprocess
	} else if xmlElement.M_activityType=="##EventSubprocess"{
		object.M_activityType=BPMS_Runtime.ActivityType_EventSubprocess
	} else if xmlElement.M_activityType=="##AdHocSubprocess"{
		object.M_activityType=BPMS_Runtime.ActivityType_AdHocSubprocess
	} else if xmlElement.M_activityType=="##Transaction"{
		object.M_activityType=BPMS_Runtime.ActivityType_Transaction
	} else if xmlElement.M_activityType=="##CallActivity"{
		object.M_activityType=BPMS_Runtime.ActivityType_CallActivity
	}

	/** LoopCharacteristicType **/
	if xmlElement.M_loopCharacteristicType=="##StandardLoopCharacteristics"{
		object.M_loopCharacteristicType=BPMS_Runtime.LoopCharacteristicType_StandardLoopCharacteristics
	} else if xmlElement.M_loopCharacteristicType=="##MultiInstanceLoopCharacteristics"{
		object.M_loopCharacteristicType=BPMS_Runtime.LoopCharacteristicType_MultiInstanceLoopCharacteristics
	}

	/** MultiInstanceBehaviorType **/
	if xmlElement.M_multiInstanceBehaviorType=="##None"{
		object.M_multiInstanceBehaviorType=BPMS_Runtime.MultiInstanceBehaviorType_None
	} else if xmlElement.M_multiInstanceBehaviorType=="##One"{
		object.M_multiInstanceBehaviorType=BPMS_Runtime.MultiInstanceBehaviorType_One
	} else if xmlElement.M_multiInstanceBehaviorType=="##All"{
		object.M_multiInstanceBehaviorType=BPMS_Runtime.MultiInstanceBehaviorType_All
	} else if xmlElement.M_multiInstanceBehaviorType=="##Complex"{
		object.M_multiInstanceBehaviorType=BPMS_Runtime.MultiInstanceBehaviorType_Complex
	}
}

/** inititialisation of EventData **/
func (this *RuntimeXmlFactory) InitEventData(xmlElement *BPMS_Runtime.XsdEventData,object *BPMS_Runtime.EventData){
	log.Println("Initialize EventData")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.EventData%" + Utility.RandomUUID()	}

	/** EventData **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Runtimes **/
func (this *RuntimeXmlFactory) InitRuntimes(xmlElement *BPMS_Runtime.XsdRuntimes,object *BPMS_Runtime.Runtimes){
	log.Println("Initialize Runtimes")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.Runtimes%" + Utility.RandomUUID()	}

	/** Init definitions **/
	object.M_definitions= make([]*BPMS_Runtime.DefinitionsInstance,0)
	for i:=0;i<len(xmlElement.M_definitions); i++{
		val:=new(BPMS_Runtime.DefinitionsInstance)
		this.InitDefinitionsInstance(xmlElement.M_definitions[i],val)
		object.M_definitions= append(object.M_definitions, val)

		/** association initialisation **/
	}

	/** Init exception **/
	object.M_exceptions= make([]*BPMS_Runtime.Exception,0)
	for i:=0;i<len(xmlElement.M_exceptions); i++{
		val:=new(BPMS_Runtime.Exception)
		this.InitException(xmlElement.M_exceptions[i],val)
		object.M_exceptions= append(object.M_exceptions, val)

		/** association initialisation **/
		val.SetRuntimesPtr(object)
	}

	/** Init trigger **/
	object.M_triggers= make([]*BPMS_Runtime.Trigger,0)
	for i:=0;i<len(xmlElement.M_triggers); i++{
		val:=new(BPMS_Runtime.Trigger)
		this.InitTrigger(xmlElement.M_triggers[i],val)
		object.M_triggers= append(object.M_triggers, val)

		/** association initialisation **/
		val.SetRuntimesPtr(object)
	}

	/** Init correlationInfo **/
	object.M_correlationInfos= make([]*BPMS_Runtime.CorrelationInfo,0)
	for i:=0;i<len(xmlElement.M_correlationInfos); i++{
		val:=new(BPMS_Runtime.CorrelationInfo)
		this.InitCorrelationInfo(xmlElement.M_correlationInfos[i],val)
		object.M_correlationInfos= append(object.M_correlationInfos, val)

		/** association initialisation **/
		val.SetRuntimesPtr(object)
	}

	/** Runtimes **/
	object.M_id= xmlElement.M_id

	/** Runtimes **/
	object.M_name= xmlElement.M_name

	/** Runtimes **/
	object.M_version= xmlElement.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of SubprocessInstance **/
func (this *RuntimeXmlFactory) InitSubprocessInstance(xmlElement *BPMS_Runtime.XsdSubprocessInstance,object *BPMS_Runtime.SubprocessInstance){
	log.Println("Initialize SubprocessInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.SubprocessInstance%" + Utility.RandomUUID()	}

	/** SubprocessInstance **/
	object.M_id= xmlElement.M_id

	/** SubprocessInstance **/
	object.M_bpmnElementId= xmlElement.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref inputRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_inputRef); i++ {
		if _, ok:= this.m_object[object.M_id]["inputRef"]; !ok {
			this.m_object[object.M_id]["inputRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["inputRef"] = append(this.m_object[object.M_id]["inputRef"], xmlElement.M_inputRef[i])
	}


	/** Init ref outputRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_outputRef); i++ {
		if _, ok:= this.m_object[object.M_id]["outputRef"]; !ok {
			this.m_object[object.M_id]["outputRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["outputRef"] = append(this.m_object[object.M_id]["outputRef"], xmlElement.M_outputRef[i])
	}


	/** FlowNodeType **/
	if xmlElement.M_flowNodeType=="##AbstractTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AbstractTask
	} else if xmlElement.M_flowNodeType=="##ServiceTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ServiceTask
	} else if xmlElement.M_flowNodeType=="##UserTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_UserTask
	} else if xmlElement.M_flowNodeType=="##ManualTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ManualTask
	} else if xmlElement.M_flowNodeType=="##BusinessRuleTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BusinessRuleTask
	} else if xmlElement.M_flowNodeType=="##ScriptTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ScriptTask
	} else if xmlElement.M_flowNodeType=="##EmbeddedSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EmbeddedSubprocess
	} else if xmlElement.M_flowNodeType=="##EventSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventSubprocess
	} else if xmlElement.M_flowNodeType=="##AdHocSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AdHocSubprocess
	} else if xmlElement.M_flowNodeType=="##Transaction"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_Transaction
	} else if xmlElement.M_flowNodeType=="##CallActivity"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_CallActivity
	} else if xmlElement.M_flowNodeType=="##ParallelGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ParallelGateway
	} else if xmlElement.M_flowNodeType=="##ExclusiveGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ExclusiveGateway
	} else if xmlElement.M_flowNodeType=="##InclusiveGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_InclusiveGateway
	} else if xmlElement.M_flowNodeType=="##EventBasedGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventBasedGateway
	} else if xmlElement.M_flowNodeType=="##ComplexGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ComplexGateway
	} else if xmlElement.M_flowNodeType=="##StartEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_StartEvent
	} else if xmlElement.M_flowNodeType=="##IntermediateCatchEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateCatchEvent
	} else if xmlElement.M_flowNodeType=="##BoundaryEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BoundaryEvent
	} else if xmlElement.M_flowNodeType=="##EndEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EndEvent
	} else if xmlElement.M_flowNodeType=="##IntermediateThrowEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateThrowEvent
	}

	/** LifecycleState **/
	if xmlElement.M_lifecycleState=="##Completed"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completed
	} else if xmlElement.M_lifecycleState=="##Compensated"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensated
	} else if xmlElement.M_lifecycleState=="##Failed"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failed
	} else if xmlElement.M_lifecycleState=="##Terminated"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminated
	} else if xmlElement.M_lifecycleState=="##Ready"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Ready
	} else if xmlElement.M_lifecycleState=="##Active"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Active
	} else if xmlElement.M_lifecycleState=="##Completing"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completing
	} else if xmlElement.M_lifecycleState=="##Compensating"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensating
	} else if xmlElement.M_lifecycleState=="##Failing"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failing
	} else if xmlElement.M_lifecycleState=="##Terminating"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminating
	}

	/** Init activityInstance **/
	object.M_flowNodeInstances= make([]BPMS_Runtime.FlowNodeInstance,0)
	for i:=0;i<len(xmlElement.M_flowNodeInstances_0); i++{
		val:=new(BPMS_Runtime.ActivityInstance)
		this.InitActivityInstance(xmlElement.M_flowNodeInstances_0[i],val)
		object.M_flowNodeInstances= append(object.M_flowNodeInstances, val)

		/** association initialisation **/
	}

	/** Init SubprocessInstance **/
	for i:=0;i<len(xmlElement.M_flowNodeInstances_1); i++{
		val:=new(BPMS_Runtime.SubprocessInstance)
		this.InitSubprocessInstance(xmlElement.M_flowNodeInstances_1[i],val)
		object.M_flowNodeInstances= append(object.M_flowNodeInstances, val)

		/** association initialisation **/
	}

	/** Init gatewayInstance **/
	for i:=0;i<len(xmlElement.M_flowNodeInstances_2); i++{
		val:=new(BPMS_Runtime.GatewayInstance)
		this.InitGatewayInstance(xmlElement.M_flowNodeInstances_2[i],val)
		object.M_flowNodeInstances= append(object.M_flowNodeInstances, val)

		/** association initialisation **/
	}

	/** Init eventInstance **/
	for i:=0;i<len(xmlElement.M_flowNodeInstances_3); i++{
		val:=new(BPMS_Runtime.EventInstance)
		this.InitEventInstance(xmlElement.M_flowNodeInstances_3[i],val)
		object.M_flowNodeInstances= append(object.M_flowNodeInstances, val)

		/** association initialisation **/
	}

	/** Init outputRef **/
	object.M_connectingObjects= make([]*BPMS_Runtime.ConnectingObject,0)
	for i:=0;i<len(xmlElement.M_connectingObjects); i++{
		val:=new(BPMS_Runtime.ConnectingObject)
		this.InitConnectingObject(xmlElement.M_connectingObjects[i],val)
		object.M_connectingObjects= append(object.M_connectingObjects, val)

		/** association initialisation **/
	}

	/** SubprocessType **/
	if xmlElement.M_SubprocessType=="##EmbeddedSubprocess"{
		object.M_SubprocessType=BPMS_Runtime.SubprocessType_EmbeddedSubprocess
	} else if xmlElement.M_SubprocessType=="##EventSubprocess"{
		object.M_SubprocessType=BPMS_Runtime.SubprocessType_EventSubprocess
	} else if xmlElement.M_SubprocessType=="##AdHocSubprocess"{
		object.M_SubprocessType=BPMS_Runtime.SubprocessType_AdHocSubprocess
	} else if xmlElement.M_SubprocessType=="##Transaction"{
		object.M_SubprocessType=BPMS_Runtime.SubprocessType_Transaction
	}
}

/** inititialisation of EventInstance **/
func (this *RuntimeXmlFactory) InitEventInstance(xmlElement *BPMS_Runtime.XsdEventInstance,object *BPMS_Runtime.EventInstance){
	log.Println("Initialize EventInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.EventInstance%" + Utility.RandomUUID()	}

	/** EventInstance **/
	object.M_id= xmlElement.M_id

	/** EventInstance **/
	object.M_bpmnElementId= xmlElement.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref inputRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_inputRef); i++ {
		if _, ok:= this.m_object[object.M_id]["inputRef"]; !ok {
			this.m_object[object.M_id]["inputRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["inputRef"] = append(this.m_object[object.M_id]["inputRef"], xmlElement.M_inputRef[i])
	}


	/** Init ref outputRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_outputRef); i++ {
		if _, ok:= this.m_object[object.M_id]["outputRef"]; !ok {
			this.m_object[object.M_id]["outputRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["outputRef"] = append(this.m_object[object.M_id]["outputRef"], xmlElement.M_outputRef[i])
	}


	/** FlowNodeType **/
	if xmlElement.M_flowNodeType=="##AbstractTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AbstractTask
	} else if xmlElement.M_flowNodeType=="##ServiceTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ServiceTask
	} else if xmlElement.M_flowNodeType=="##UserTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_UserTask
	} else if xmlElement.M_flowNodeType=="##ManualTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ManualTask
	} else if xmlElement.M_flowNodeType=="##BusinessRuleTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BusinessRuleTask
	} else if xmlElement.M_flowNodeType=="##ScriptTask"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ScriptTask
	} else if xmlElement.M_flowNodeType=="##EmbeddedSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EmbeddedSubprocess
	} else if xmlElement.M_flowNodeType=="##EventSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventSubprocess
	} else if xmlElement.M_flowNodeType=="##AdHocSubprocess"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_AdHocSubprocess
	} else if xmlElement.M_flowNodeType=="##Transaction"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_Transaction
	} else if xmlElement.M_flowNodeType=="##CallActivity"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_CallActivity
	} else if xmlElement.M_flowNodeType=="##ParallelGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ParallelGateway
	} else if xmlElement.M_flowNodeType=="##ExclusiveGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ExclusiveGateway
	} else if xmlElement.M_flowNodeType=="##InclusiveGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_InclusiveGateway
	} else if xmlElement.M_flowNodeType=="##EventBasedGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EventBasedGateway
	} else if xmlElement.M_flowNodeType=="##ComplexGateway"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_ComplexGateway
	} else if xmlElement.M_flowNodeType=="##StartEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_StartEvent
	} else if xmlElement.M_flowNodeType=="##IntermediateCatchEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateCatchEvent
	} else if xmlElement.M_flowNodeType=="##BoundaryEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_BoundaryEvent
	} else if xmlElement.M_flowNodeType=="##EndEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_EndEvent
	} else if xmlElement.M_flowNodeType=="##IntermediateThrowEvent"{
		object.M_flowNodeType=BPMS_Runtime.FlowNodeType_IntermediateThrowEvent
	}

	/** LifecycleState **/
	if xmlElement.M_lifecycleState=="##Completed"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completed
	} else if xmlElement.M_lifecycleState=="##Compensated"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensated
	} else if xmlElement.M_lifecycleState=="##Failed"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failed
	} else if xmlElement.M_lifecycleState=="##Terminated"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminated
	} else if xmlElement.M_lifecycleState=="##Ready"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Ready
	} else if xmlElement.M_lifecycleState=="##Active"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Active
	} else if xmlElement.M_lifecycleState=="##Completing"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Completing
	} else if xmlElement.M_lifecycleState=="##Compensating"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Compensating
	} else if xmlElement.M_lifecycleState=="##Failing"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Failing
	} else if xmlElement.M_lifecycleState=="##Terminating"{
		object.M_lifecycleState=BPMS_Runtime.LifecycleState_Terminating
	}

	/** Init sourceRef **/
	object.M_eventDefintionInstances= make([]*BPMS_Runtime.EventDefinitionInstance,0)
	for i:=0;i<len(xmlElement.M_eventDefintionInstances); i++{
		val:=new(BPMS_Runtime.EventDefinitionInstance)
		this.InitEventDefinitionInstance(xmlElement.M_eventDefintionInstances[i],val)
		object.M_eventDefintionInstances= append(object.M_eventDefintionInstances, val)

		/** association initialisation **/
		val.SetEventInstancePtr(object)
	}

	/** EventType **/
	if xmlElement.M_eventType=="##StartEvent"{
		object.M_eventType=BPMS_Runtime.EventType_StartEvent
	} else if xmlElement.M_eventType=="##IntermediateCatchEvent"{
		object.M_eventType=BPMS_Runtime.EventType_IntermediateCatchEvent
	} else if xmlElement.M_eventType=="##BoundaryEvent"{
		object.M_eventType=BPMS_Runtime.EventType_BoundaryEvent
	} else if xmlElement.M_eventType=="##EndEvent"{
		object.M_eventType=BPMS_Runtime.EventType_EndEvent
	} else if xmlElement.M_eventType=="##IntermediateThrowEvent"{
		object.M_eventType=BPMS_Runtime.EventType_IntermediateThrowEvent
	}
}

/** inititialisation of Exception **/
func (this *RuntimeXmlFactory) InitException(xmlElement *BPMS_Runtime.XsdException,object *BPMS_Runtime.Exception){
	log.Println("Initialize Exception")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS_Runtime.Exception%" + Utility.RandomUUID()	}

	/** ExceptionType **/
	if xmlElement.M_exceptionType=="##GatewayException"{
		object.M_exceptionType=BPMS_Runtime.ExceptionType_GatewayException
	} else if xmlElement.M_exceptionType=="##NoIORuleException"{
		object.M_exceptionType=BPMS_Runtime.ExceptionType_NoIORuleException
	} else if xmlElement.M_exceptionType=="##NoAvailableOutputSetException"{
		object.M_exceptionType=BPMS_Runtime.ExceptionType_NoAvailableOutputSetException
	} else if xmlElement.M_exceptionType=="##NotMatchingIOSpecification"{
		object.M_exceptionType=BPMS_Runtime.ExceptionType_NotMatchingIOSpecification
	} else if xmlElement.M_exceptionType=="##IllegalStartEventException"{
		object.M_exceptionType=BPMS_Runtime.ExceptionType_IllegalStartEventException
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of DefinitionsInstance **/
func (this *RuntimeXmlFactory) SerialyzeDefinitionsInstance(xmlElement *BPMS_Runtime.XsdDefinitionsInstance,object *BPMS_Runtime.DefinitionsInstance){
	if xmlElement == nil{
		return
	}

	/** DefinitionsInstance **/
	xmlElement.M_id= object.M_id

	/** DefinitionsInstance **/
	xmlElement.M_bpmnElementId= object.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ProcessInstance **/
	if len(object.M_processInstances) > 0 {
		xmlElement.M_processInstances= make([]*BPMS_Runtime.XsdProcessInstance,0)
	}

	/** Now I will save the value of processInstances **/
	for i:=0; i<len(object.M_processInstances);i++{
		xmlElement.M_processInstances=append(xmlElement.M_processInstances,new(BPMS_Runtime.XsdProcessInstance))
		this.SerialyzeProcessInstance(xmlElement.M_processInstances[i],object.M_processInstances[i])
	}
}

/** serialysation of SubprocessInstance **/
func (this *RuntimeXmlFactory) SerialyzeSubprocessInstance(xmlElement *BPMS_Runtime.XsdSubprocessInstance,object *BPMS_Runtime.SubprocessInstance){
	if xmlElement == nil{
		return
	}

	/** SubprocessInstance **/
	xmlElement.M_id= object.M_id

	/** SubprocessInstance **/
	xmlElement.M_bpmnElementId= object.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref inputRef **/
		xmlElement.M_inputRef= object.M_inputRef


	/** Serialyze ref outputRef **/
		xmlElement.M_outputRef= object.M_outputRef


	/** FlowNodeType **/
	if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AbstractTask{
		xmlElement.M_flowNodeType="##AbstractTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ServiceTask{
		xmlElement.M_flowNodeType="##ServiceTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_UserTask{
		xmlElement.M_flowNodeType="##UserTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ManualTask{
		xmlElement.M_flowNodeType="##ManualTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BusinessRuleTask{
		xmlElement.M_flowNodeType="##BusinessRuleTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ScriptTask{
		xmlElement.M_flowNodeType="##ScriptTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EmbeddedSubprocess{
		xmlElement.M_flowNodeType="##EmbeddedSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventSubprocess{
		xmlElement.M_flowNodeType="##EventSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AdHocSubprocess{
		xmlElement.M_flowNodeType="##AdHocSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_Transaction{
		xmlElement.M_flowNodeType="##Transaction"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_CallActivity{
		xmlElement.M_flowNodeType="##CallActivity"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ParallelGateway{
		xmlElement.M_flowNodeType="##ParallelGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ExclusiveGateway{
		xmlElement.M_flowNodeType="##ExclusiveGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_InclusiveGateway{
		xmlElement.M_flowNodeType="##InclusiveGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventBasedGateway{
		xmlElement.M_flowNodeType="##EventBasedGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ComplexGateway{
		xmlElement.M_flowNodeType="##ComplexGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_StartEvent{
		xmlElement.M_flowNodeType="##StartEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateCatchEvent{
		xmlElement.M_flowNodeType="##IntermediateCatchEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BoundaryEvent{
		xmlElement.M_flowNodeType="##BoundaryEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EndEvent{
		xmlElement.M_flowNodeType="##EndEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateThrowEvent{
		xmlElement.M_flowNodeType="##IntermediateThrowEvent"
	}

	/** LifecycleState **/
	if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completed{
		xmlElement.M_lifecycleState="##Completed"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensated{
		xmlElement.M_lifecycleState="##Compensated"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failed{
		xmlElement.M_lifecycleState="##Failed"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminated{
		xmlElement.M_lifecycleState="##Terminated"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Ready{
		xmlElement.M_lifecycleState="##Ready"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Active{
		xmlElement.M_lifecycleState="##Active"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completing{
		xmlElement.M_lifecycleState="##Completing"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensating{
		xmlElement.M_lifecycleState="##Compensating"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failing{
		xmlElement.M_lifecycleState="##Failing"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminating{
		xmlElement.M_lifecycleState="##Terminating"
	}

	/** Serialyze FlowNodeInstance **/
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_0= make([]*BPMS_Runtime.XsdActivityInstance,0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_1= make([]*BPMS_Runtime.XsdSubprocessInstance,0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_2= make([]*BPMS_Runtime.XsdGatewayInstance,0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_3= make([]*BPMS_Runtime.XsdEventInstance,0)
	}

	/** Now I will save the value of flowNodeInstances **/
	for i:=0; i<len(object.M_flowNodeInstances);i++{
		switch v:= object.M_flowNodeInstances[i].(type){
			case *BPMS_Runtime.ActivityInstance:
				xmlElement.M_flowNodeInstances_0=append(xmlElement.M_flowNodeInstances_0,new(BPMS_Runtime.XsdActivityInstance))
				this.SerialyzeActivityInstance(xmlElement.M_flowNodeInstances_0[len(xmlElement.M_flowNodeInstances_0)-1],v)
				log.Println("Serialyze SubprocessInstance:flowNodeInstances:ActivityInstance")
			case *BPMS_Runtime.SubprocessInstance:
				xmlElement.M_flowNodeInstances_1=append(xmlElement.M_flowNodeInstances_1,new(BPMS_Runtime.XsdSubprocessInstance))
				this.SerialyzeSubprocessInstance(xmlElement.M_flowNodeInstances_1[len(xmlElement.M_flowNodeInstances_1)-1],v)
				log.Println("Serialyze SubprocessInstance:flowNodeInstances:SubprocessInstance")
			case *BPMS_Runtime.GatewayInstance:
				xmlElement.M_flowNodeInstances_2=append(xmlElement.M_flowNodeInstances_2,new(BPMS_Runtime.XsdGatewayInstance))
				this.SerialyzeGatewayInstance(xmlElement.M_flowNodeInstances_2[len(xmlElement.M_flowNodeInstances_2)-1],v)
				log.Println("Serialyze SubprocessInstance:flowNodeInstances:GatewayInstance")
			case *BPMS_Runtime.EventInstance:
				xmlElement.M_flowNodeInstances_3=append(xmlElement.M_flowNodeInstances_3,new(BPMS_Runtime.XsdEventInstance))
				this.SerialyzeEventInstance(xmlElement.M_flowNodeInstances_3[len(xmlElement.M_flowNodeInstances_3)-1],v)
				log.Println("Serialyze SubprocessInstance:flowNodeInstances:EventInstance")
		}
	}

	/** Serialyze ConnectingObject **/
	if len(object.M_connectingObjects) > 0 {
		xmlElement.M_connectingObjects= make([]*BPMS_Runtime.XsdConnectingObject,0)
	}

	/** Now I will save the value of connectingObjects **/
	for i:=0; i<len(object.M_connectingObjects);i++{
		xmlElement.M_connectingObjects=append(xmlElement.M_connectingObjects,new(BPMS_Runtime.XsdConnectingObject))
		this.SerialyzeConnectingObject(xmlElement.M_connectingObjects[i],object.M_connectingObjects[i])
	}

	/** SubprocessType **/
	if object.M_SubprocessType==BPMS_Runtime.SubprocessType_EmbeddedSubprocess{
		xmlElement.M_SubprocessType="##EmbeddedSubprocess"
	} else if object.M_SubprocessType==BPMS_Runtime.SubprocessType_EventSubprocess{
		xmlElement.M_SubprocessType="##EventSubprocess"
	} else if object.M_SubprocessType==BPMS_Runtime.SubprocessType_AdHocSubprocess{
		xmlElement.M_SubprocessType="##AdHocSubprocess"
	} else if object.M_SubprocessType==BPMS_Runtime.SubprocessType_Transaction{
		xmlElement.M_SubprocessType="##Transaction"
	}
}

/** serialysation of ConnectingObject **/
func (this *RuntimeXmlFactory) SerialyzeConnectingObject(xmlElement *BPMS_Runtime.XsdConnectingObject,object *BPMS_Runtime.ConnectingObject){
	if xmlElement == nil{
		return
	}

	/** ConnectingObject **/
	xmlElement.M_id= object.M_id

	/** ConnectingObject **/
	xmlElement.M_bpmnElementId= object.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref sourceRef **/
	xmlElement.M_sourceRef=&object.M_sourceRef


	/** Serialyze ref targetRef **/
	xmlElement.M_targetRef=&object.M_targetRef


	/** ConnectingObjectType **/
	if object.M_connectingObjectType==BPMS_Runtime.ConnectingObjectType_SequenceFlow{
		xmlElement.M_connectingObjectType="##SequenceFlow"
	} else if object.M_connectingObjectType==BPMS_Runtime.ConnectingObjectType_MessageFlow{
		xmlElement.M_connectingObjectType="##MessageFlow"
	} else if object.M_connectingObjectType==BPMS_Runtime.ConnectingObjectType_Association{
		xmlElement.M_connectingObjectType="##Association"
	} else if object.M_connectingObjectType==BPMS_Runtime.ConnectingObjectType_DataAssociation{
		xmlElement.M_connectingObjectType="##DataAssociation"
	}
}

/** serialysation of Runtimes **/
func (this *RuntimeXmlFactory) SerialyzeRuntimes(xmlElement *BPMS_Runtime.XsdRuntimes,object *BPMS_Runtime.Runtimes){
	if xmlElement == nil{
		return
	}

	/** Serialyze DefinitionsInstance **/
	if len(object.M_definitions) > 0 {
		xmlElement.M_definitions= make([]*BPMS_Runtime.XsdDefinitionsInstance,0)
	}

	/** Now I will save the value of definitions **/
	for i:=0; i<len(object.M_definitions);i++{
		xmlElement.M_definitions=append(xmlElement.M_definitions,new(BPMS_Runtime.XsdDefinitionsInstance))
		this.SerialyzeDefinitionsInstance(xmlElement.M_definitions[i],object.M_definitions[i])
	}

	/** Serialyze Exception **/
	if len(object.M_exceptions) > 0 {
		xmlElement.M_exceptions= make([]*BPMS_Runtime.XsdException,0)
	}

	/** Now I will save the value of exceptions **/
	for i:=0; i<len(object.M_exceptions);i++{
		xmlElement.M_exceptions=append(xmlElement.M_exceptions,new(BPMS_Runtime.XsdException))
		this.SerialyzeException(xmlElement.M_exceptions[i],object.M_exceptions[i])
	}

	/** Serialyze Trigger **/
	if len(object.M_triggers) > 0 {
		xmlElement.M_triggers= make([]*BPMS_Runtime.XsdTrigger,0)
	}

	/** Now I will save the value of triggers **/
	for i:=0; i<len(object.M_triggers);i++{
		xmlElement.M_triggers=append(xmlElement.M_triggers,new(BPMS_Runtime.XsdTrigger))
		this.SerialyzeTrigger(xmlElement.M_triggers[i],object.M_triggers[i])
	}

	/** Serialyze CorrelationInfo **/
	if len(object.M_correlationInfos) > 0 {
		xmlElement.M_correlationInfos= make([]*BPMS_Runtime.XsdCorrelationInfo,0)
	}

	/** Now I will save the value of correlationInfos **/
	for i:=0; i<len(object.M_correlationInfos);i++{
		xmlElement.M_correlationInfos=append(xmlElement.M_correlationInfos,new(BPMS_Runtime.XsdCorrelationInfo))
		this.SerialyzeCorrelationInfo(xmlElement.M_correlationInfos[i],object.M_correlationInfos[i])
	}

	/** Runtimes **/
	xmlElement.M_id= object.M_id

	/** Runtimes **/
	xmlElement.M_name= object.M_name

	/** Runtimes **/
	xmlElement.M_version= object.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of CorrelationInfo **/
func (this *RuntimeXmlFactory) SerialyzeCorrelationInfo(xmlElement *BPMS_Runtime.XsdCorrelationInfo,object *BPMS_Runtime.CorrelationInfo){
	if xmlElement == nil{
		return
	}

	/** CorrelationInfo **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of EventData **/
func (this *RuntimeXmlFactory) SerialyzeEventData(xmlElement *BPMS_Runtime.XsdEventData,object *BPMS_Runtime.EventData){
	if xmlElement == nil{
		return
	}

	/** EventData **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of GatewayInstance **/
func (this *RuntimeXmlFactory) SerialyzeGatewayInstance(xmlElement *BPMS_Runtime.XsdGatewayInstance,object *BPMS_Runtime.GatewayInstance){
	if xmlElement == nil{
		return
	}

	/** GatewayInstance **/
	xmlElement.M_id= object.M_id

	/** GatewayInstance **/
	xmlElement.M_bpmnElementId= object.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref inputRef **/
		xmlElement.M_inputRef= object.M_inputRef


	/** Serialyze ref outputRef **/
		xmlElement.M_outputRef= object.M_outputRef


	/** FlowNodeType **/
	if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AbstractTask{
		xmlElement.M_flowNodeType="##AbstractTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ServiceTask{
		xmlElement.M_flowNodeType="##ServiceTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_UserTask{
		xmlElement.M_flowNodeType="##UserTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ManualTask{
		xmlElement.M_flowNodeType="##ManualTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BusinessRuleTask{
		xmlElement.M_flowNodeType="##BusinessRuleTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ScriptTask{
		xmlElement.M_flowNodeType="##ScriptTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EmbeddedSubprocess{
		xmlElement.M_flowNodeType="##EmbeddedSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventSubprocess{
		xmlElement.M_flowNodeType="##EventSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AdHocSubprocess{
		xmlElement.M_flowNodeType="##AdHocSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_Transaction{
		xmlElement.M_flowNodeType="##Transaction"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_CallActivity{
		xmlElement.M_flowNodeType="##CallActivity"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ParallelGateway{
		xmlElement.M_flowNodeType="##ParallelGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ExclusiveGateway{
		xmlElement.M_flowNodeType="##ExclusiveGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_InclusiveGateway{
		xmlElement.M_flowNodeType="##InclusiveGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventBasedGateway{
		xmlElement.M_flowNodeType="##EventBasedGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ComplexGateway{
		xmlElement.M_flowNodeType="##ComplexGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_StartEvent{
		xmlElement.M_flowNodeType="##StartEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateCatchEvent{
		xmlElement.M_flowNodeType="##IntermediateCatchEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BoundaryEvent{
		xmlElement.M_flowNodeType="##BoundaryEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EndEvent{
		xmlElement.M_flowNodeType="##EndEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateThrowEvent{
		xmlElement.M_flowNodeType="##IntermediateThrowEvent"
	}

	/** LifecycleState **/
	if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completed{
		xmlElement.M_lifecycleState="##Completed"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensated{
		xmlElement.M_lifecycleState="##Compensated"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failed{
		xmlElement.M_lifecycleState="##Failed"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminated{
		xmlElement.M_lifecycleState="##Terminated"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Ready{
		xmlElement.M_lifecycleState="##Ready"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Active{
		xmlElement.M_lifecycleState="##Active"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completing{
		xmlElement.M_lifecycleState="##Completing"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensating{
		xmlElement.M_lifecycleState="##Compensating"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failing{
		xmlElement.M_lifecycleState="##Failing"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminating{
		xmlElement.M_lifecycleState="##Terminating"
	}

	/** GatewayType **/
	if object.M_gatewayType==BPMS_Runtime.GatewayType_ParallelGateway{
		xmlElement.M_gatewayType="##ParallelGateway"
	} else if object.M_gatewayType==BPMS_Runtime.GatewayType_ExclusiveGateway{
		xmlElement.M_gatewayType="##ExclusiveGateway"
	} else if object.M_gatewayType==BPMS_Runtime.GatewayType_InclusiveGateway{
		xmlElement.M_gatewayType="##InclusiveGateway"
	} else if object.M_gatewayType==BPMS_Runtime.GatewayType_EventBasedGateway{
		xmlElement.M_gatewayType="##EventBasedGateway"
	} else if object.M_gatewayType==BPMS_Runtime.GatewayType_ComplexGateway{
		xmlElement.M_gatewayType="##ComplexGateway"
	}
}

/** serialysation of EventDefinitionInstance **/
func (this *RuntimeXmlFactory) SerialyzeEventDefinitionInstance(xmlElement *BPMS_Runtime.XsdEventDefinitionInstance,object *BPMS_Runtime.EventDefinitionInstance){
	if xmlElement == nil{
		return
	}

	/** EventDefinitionInstance **/
	xmlElement.M_id= object.M_id

	/** EventDefinitionInstance **/
	xmlElement.M_bpmnElementId= object.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** EventDefinitionType **/
	if object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_MessageEventDefinition{
		xmlElement.M_eventDefinitionType="##MessageEventDefinition"
	} else if object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_LinkEventDefinition{
		xmlElement.M_eventDefinitionType="##LinkEventDefinition"
	} else if object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_ErrorEventDefinition{
		xmlElement.M_eventDefinitionType="##ErrorEventDefinition"
	} else if object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_TerminateEventDefinition{
		xmlElement.M_eventDefinitionType="##TerminateEventDefinition"
	} else if object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_CompensationEventDefinition{
		xmlElement.M_eventDefinitionType="##CompensationEventDefinition"
	} else if object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_ConditionalEventDefinition{
		xmlElement.M_eventDefinitionType="##ConditionalEventDefinition"
	} else if object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_TimerEventDefinition{
		xmlElement.M_eventDefinitionType="##TimerEventDefinition"
	} else if object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_CancelEventDefinition{
		xmlElement.M_eventDefinitionType="##CancelEventDefinition"
	} else if object.M_eventDefinitionType==BPMS_Runtime.EventDefinitionType_EscalationEventDefinition{
		xmlElement.M_eventDefinitionType="##EscalationEventDefinition"
	}
}

/** serialysation of Exception **/
func (this *RuntimeXmlFactory) SerialyzeException(xmlElement *BPMS_Runtime.XsdException,object *BPMS_Runtime.Exception){
	if xmlElement == nil{
		return
	}

	/** ExceptionType **/
	if object.M_exceptionType==BPMS_Runtime.ExceptionType_GatewayException{
		xmlElement.M_exceptionType="##GatewayException"
	} else if object.M_exceptionType==BPMS_Runtime.ExceptionType_NoIORuleException{
		xmlElement.M_exceptionType="##NoIORuleException"
	} else if object.M_exceptionType==BPMS_Runtime.ExceptionType_NoAvailableOutputSetException{
		xmlElement.M_exceptionType="##NoAvailableOutputSetException"
	} else if object.M_exceptionType==BPMS_Runtime.ExceptionType_NotMatchingIOSpecification{
		xmlElement.M_exceptionType="##NotMatchingIOSpecification"
	} else if object.M_exceptionType==BPMS_Runtime.ExceptionType_IllegalStartEventException{
		xmlElement.M_exceptionType="##IllegalStartEventException"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Trigger **/
func (this *RuntimeXmlFactory) SerialyzeTrigger(xmlElement *BPMS_Runtime.XsdTrigger,object *BPMS_Runtime.Trigger){
	if xmlElement == nil{
		return
	}

	/** Serialyze EventData **/
	if len(object.M_eventDatas) > 0 {
		xmlElement.M_eventDatas= make([]*BPMS_Runtime.XsdEventData,0)
	}

	/** Now I will save the value of eventDatas **/
	for i:=0; i<len(object.M_eventDatas);i++{
		xmlElement.M_eventDatas=append(xmlElement.M_eventDatas,new(BPMS_Runtime.XsdEventData))
		this.SerialyzeEventData(xmlElement.M_eventDatas[i],object.M_eventDatas[i])
	}

	/** Serialyze ref dataRef **/
		xmlElement.M_dataRef= object.M_dataRef


	/** Serialyze ref sourceRef **/
	xmlElement.M_sourceRef=&object.M_sourceRef


	/** Serialyze ref targetRef **/
	xmlElement.M_targetRef=&object.M_targetRef


	/** Trigger **/
	xmlElement.M_id= object.M_id

	/** Trigger **/
	xmlElement.M_processUUID= object.M_processUUID

	/** Trigger **/
	xmlElement.M_sessionId= object.M_sessionId

	/** EventTriggerType **/
	if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_None{
		xmlElement.M_eventTriggerType="##None"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Timer{
		xmlElement.M_eventTriggerType="##Timer"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Conditional{
		xmlElement.M_eventTriggerType="##Conditional"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Message{
		xmlElement.M_eventTriggerType="##Message"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Signal{
		xmlElement.M_eventTriggerType="##Signal"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Multiple{
		xmlElement.M_eventTriggerType="##Multiple"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_ParallelMultiple{
		xmlElement.M_eventTriggerType="##ParallelMultiple"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Escalation{
		xmlElement.M_eventTriggerType="##Escalation"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Error{
		xmlElement.M_eventTriggerType="##Error"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Compensation{
		xmlElement.M_eventTriggerType="##Compensation"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Terminate{
		xmlElement.M_eventTriggerType="##Terminate"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Cancel{
		xmlElement.M_eventTriggerType="##Cancel"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Link{
		xmlElement.M_eventTriggerType="##Link"
	} else if object.M_eventTriggerType==BPMS_Runtime.EventTriggerType_Start{
		xmlElement.M_eventTriggerType="##Start"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ActivityInstance **/
func (this *RuntimeXmlFactory) SerialyzeActivityInstance(xmlElement *BPMS_Runtime.XsdActivityInstance,object *BPMS_Runtime.ActivityInstance){
	if xmlElement == nil{
		return
	}

	/** ActivityInstance **/
	xmlElement.M_id= object.M_id

	/** ActivityInstance **/
	xmlElement.M_bpmnElementId= object.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref inputRef **/
		xmlElement.M_inputRef= object.M_inputRef


	/** Serialyze ref outputRef **/
		xmlElement.M_outputRef= object.M_outputRef


	/** FlowNodeType **/
	if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AbstractTask{
		xmlElement.M_flowNodeType="##AbstractTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ServiceTask{
		xmlElement.M_flowNodeType="##ServiceTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_UserTask{
		xmlElement.M_flowNodeType="##UserTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ManualTask{
		xmlElement.M_flowNodeType="##ManualTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BusinessRuleTask{
		xmlElement.M_flowNodeType="##BusinessRuleTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ScriptTask{
		xmlElement.M_flowNodeType="##ScriptTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EmbeddedSubprocess{
		xmlElement.M_flowNodeType="##EmbeddedSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventSubprocess{
		xmlElement.M_flowNodeType="##EventSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AdHocSubprocess{
		xmlElement.M_flowNodeType="##AdHocSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_Transaction{
		xmlElement.M_flowNodeType="##Transaction"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_CallActivity{
		xmlElement.M_flowNodeType="##CallActivity"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ParallelGateway{
		xmlElement.M_flowNodeType="##ParallelGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ExclusiveGateway{
		xmlElement.M_flowNodeType="##ExclusiveGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_InclusiveGateway{
		xmlElement.M_flowNodeType="##InclusiveGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventBasedGateway{
		xmlElement.M_flowNodeType="##EventBasedGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ComplexGateway{
		xmlElement.M_flowNodeType="##ComplexGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_StartEvent{
		xmlElement.M_flowNodeType="##StartEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateCatchEvent{
		xmlElement.M_flowNodeType="##IntermediateCatchEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BoundaryEvent{
		xmlElement.M_flowNodeType="##BoundaryEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EndEvent{
		xmlElement.M_flowNodeType="##EndEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateThrowEvent{
		xmlElement.M_flowNodeType="##IntermediateThrowEvent"
	}

	/** LifecycleState **/
	if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completed{
		xmlElement.M_lifecycleState="##Completed"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensated{
		xmlElement.M_lifecycleState="##Compensated"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failed{
		xmlElement.M_lifecycleState="##Failed"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminated{
		xmlElement.M_lifecycleState="##Terminated"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Ready{
		xmlElement.M_lifecycleState="##Ready"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Active{
		xmlElement.M_lifecycleState="##Active"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completing{
		xmlElement.M_lifecycleState="##Completing"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensating{
		xmlElement.M_lifecycleState="##Compensating"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failing{
		xmlElement.M_lifecycleState="##Failing"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminating{
		xmlElement.M_lifecycleState="##Terminating"
	}

	/** ActivityType **/
	if object.M_activityType==BPMS_Runtime.ActivityType_AbstractTask{
		xmlElement.M_activityType="##AbstractTask"
	} else if object.M_activityType==BPMS_Runtime.ActivityType_ServiceTask{
		xmlElement.M_activityType="##ServiceTask"
	} else if object.M_activityType==BPMS_Runtime.ActivityType_UserTask{
		xmlElement.M_activityType="##UserTask"
	} else if object.M_activityType==BPMS_Runtime.ActivityType_ManualTask{
		xmlElement.M_activityType="##ManualTask"
	} else if object.M_activityType==BPMS_Runtime.ActivityType_BusinessRuleTask{
		xmlElement.M_activityType="##BusinessRuleTask"
	} else if object.M_activityType==BPMS_Runtime.ActivityType_ScriptTask{
		xmlElement.M_activityType="##ScriptTask"
	} else if object.M_activityType==BPMS_Runtime.ActivityType_EmbeddedSubprocess{
		xmlElement.M_activityType="##EmbeddedSubprocess"
	} else if object.M_activityType==BPMS_Runtime.ActivityType_EventSubprocess{
		xmlElement.M_activityType="##EventSubprocess"
	} else if object.M_activityType==BPMS_Runtime.ActivityType_AdHocSubprocess{
		xmlElement.M_activityType="##AdHocSubprocess"
	} else if object.M_activityType==BPMS_Runtime.ActivityType_Transaction{
		xmlElement.M_activityType="##Transaction"
	} else if object.M_activityType==BPMS_Runtime.ActivityType_CallActivity{
		xmlElement.M_activityType="##CallActivity"
	}

	/** LoopCharacteristicType **/
	if object.M_loopCharacteristicType==BPMS_Runtime.LoopCharacteristicType_StandardLoopCharacteristics{
		xmlElement.M_loopCharacteristicType="##StandardLoopCharacteristics"
	} else if object.M_loopCharacteristicType==BPMS_Runtime.LoopCharacteristicType_MultiInstanceLoopCharacteristics{
		xmlElement.M_loopCharacteristicType="##MultiInstanceLoopCharacteristics"
	}

	/** MultiInstanceBehaviorType **/
	if object.M_multiInstanceBehaviorType==BPMS_Runtime.MultiInstanceBehaviorType_None{
		xmlElement.M_multiInstanceBehaviorType="##None"
	} else if object.M_multiInstanceBehaviorType==BPMS_Runtime.MultiInstanceBehaviorType_One{
		xmlElement.M_multiInstanceBehaviorType="##One"
	} else if object.M_multiInstanceBehaviorType==BPMS_Runtime.MultiInstanceBehaviorType_All{
		xmlElement.M_multiInstanceBehaviorType="##All"
	} else if object.M_multiInstanceBehaviorType==BPMS_Runtime.MultiInstanceBehaviorType_Complex{
		xmlElement.M_multiInstanceBehaviorType="##Complex"
	}
}

/** serialysation of EventInstance **/
func (this *RuntimeXmlFactory) SerialyzeEventInstance(xmlElement *BPMS_Runtime.XsdEventInstance,object *BPMS_Runtime.EventInstance){
	if xmlElement == nil{
		return
	}

	/** EventInstance **/
	xmlElement.M_id= object.M_id

	/** EventInstance **/
	xmlElement.M_bpmnElementId= object.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref inputRef **/
		xmlElement.M_inputRef= object.M_inputRef


	/** Serialyze ref outputRef **/
		xmlElement.M_outputRef= object.M_outputRef


	/** FlowNodeType **/
	if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AbstractTask{
		xmlElement.M_flowNodeType="##AbstractTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ServiceTask{
		xmlElement.M_flowNodeType="##ServiceTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_UserTask{
		xmlElement.M_flowNodeType="##UserTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ManualTask{
		xmlElement.M_flowNodeType="##ManualTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BusinessRuleTask{
		xmlElement.M_flowNodeType="##BusinessRuleTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ScriptTask{
		xmlElement.M_flowNodeType="##ScriptTask"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EmbeddedSubprocess{
		xmlElement.M_flowNodeType="##EmbeddedSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventSubprocess{
		xmlElement.M_flowNodeType="##EventSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_AdHocSubprocess{
		xmlElement.M_flowNodeType="##AdHocSubprocess"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_Transaction{
		xmlElement.M_flowNodeType="##Transaction"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_CallActivity{
		xmlElement.M_flowNodeType="##CallActivity"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ParallelGateway{
		xmlElement.M_flowNodeType="##ParallelGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ExclusiveGateway{
		xmlElement.M_flowNodeType="##ExclusiveGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_InclusiveGateway{
		xmlElement.M_flowNodeType="##InclusiveGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EventBasedGateway{
		xmlElement.M_flowNodeType="##EventBasedGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_ComplexGateway{
		xmlElement.M_flowNodeType="##ComplexGateway"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_StartEvent{
		xmlElement.M_flowNodeType="##StartEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateCatchEvent{
		xmlElement.M_flowNodeType="##IntermediateCatchEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_BoundaryEvent{
		xmlElement.M_flowNodeType="##BoundaryEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_EndEvent{
		xmlElement.M_flowNodeType="##EndEvent"
	} else if object.M_flowNodeType==BPMS_Runtime.FlowNodeType_IntermediateThrowEvent{
		xmlElement.M_flowNodeType="##IntermediateThrowEvent"
	}

	/** LifecycleState **/
	if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completed{
		xmlElement.M_lifecycleState="##Completed"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensated{
		xmlElement.M_lifecycleState="##Compensated"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failed{
		xmlElement.M_lifecycleState="##Failed"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminated{
		xmlElement.M_lifecycleState="##Terminated"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Ready{
		xmlElement.M_lifecycleState="##Ready"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Active{
		xmlElement.M_lifecycleState="##Active"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Completing{
		xmlElement.M_lifecycleState="##Completing"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Compensating{
		xmlElement.M_lifecycleState="##Compensating"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Failing{
		xmlElement.M_lifecycleState="##Failing"
	} else if object.M_lifecycleState==BPMS_Runtime.LifecycleState_Terminating{
		xmlElement.M_lifecycleState="##Terminating"
	}

	/** Serialyze EventDefinitionInstance **/
	if len(object.M_eventDefintionInstances) > 0 {
		xmlElement.M_eventDefintionInstances= make([]*BPMS_Runtime.XsdEventDefinitionInstance,0)
	}

	/** Now I will save the value of eventDefintionInstances **/
	for i:=0; i<len(object.M_eventDefintionInstances);i++{
		xmlElement.M_eventDefintionInstances=append(xmlElement.M_eventDefintionInstances,new(BPMS_Runtime.XsdEventDefinitionInstance))
		this.SerialyzeEventDefinitionInstance(xmlElement.M_eventDefintionInstances[i],object.M_eventDefintionInstances[i])
	}

	/** EventType **/
	if object.M_eventType==BPMS_Runtime.EventType_StartEvent{
		xmlElement.M_eventType="##StartEvent"
	} else if object.M_eventType==BPMS_Runtime.EventType_IntermediateCatchEvent{
		xmlElement.M_eventType="##IntermediateCatchEvent"
	} else if object.M_eventType==BPMS_Runtime.EventType_BoundaryEvent{
		xmlElement.M_eventType="##BoundaryEvent"
	} else if object.M_eventType==BPMS_Runtime.EventType_EndEvent{
		xmlElement.M_eventType="##EndEvent"
	} else if object.M_eventType==BPMS_Runtime.EventType_IntermediateThrowEvent{
		xmlElement.M_eventType="##IntermediateThrowEvent"
	}
}

/** serialysation of ProcessInstance **/
func (this *RuntimeXmlFactory) SerialyzeProcessInstance(xmlElement *BPMS_Runtime.XsdProcessInstance,object *BPMS_Runtime.ProcessInstance){
	if xmlElement == nil{
		return
	}

	/** ProcessInstance **/
	xmlElement.M_id= object.M_id

	/** ProcessInstance **/
	xmlElement.M_bpmnElementId= object.M_bpmnElementId
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze FlowNodeInstance **/
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_0= make([]*BPMS_Runtime.XsdActivityInstance,0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_1= make([]*BPMS_Runtime.XsdSubprocessInstance,0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_2= make([]*BPMS_Runtime.XsdGatewayInstance,0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_3= make([]*BPMS_Runtime.XsdEventInstance,0)
	}

	/** Now I will save the value of flowNodeInstances **/
	for i:=0; i<len(object.M_flowNodeInstances);i++{
		switch v:= object.M_flowNodeInstances[i].(type){
			case *BPMS_Runtime.ActivityInstance:
				xmlElement.M_flowNodeInstances_0=append(xmlElement.M_flowNodeInstances_0,new(BPMS_Runtime.XsdActivityInstance))
				this.SerialyzeActivityInstance(xmlElement.M_flowNodeInstances_0[len(xmlElement.M_flowNodeInstances_0)-1],v)
				log.Println("Serialyze ProcessInstance:flowNodeInstances:ActivityInstance")
			case *BPMS_Runtime.SubprocessInstance:
				xmlElement.M_flowNodeInstances_1=append(xmlElement.M_flowNodeInstances_1,new(BPMS_Runtime.XsdSubprocessInstance))
				this.SerialyzeSubprocessInstance(xmlElement.M_flowNodeInstances_1[len(xmlElement.M_flowNodeInstances_1)-1],v)
				log.Println("Serialyze ProcessInstance:flowNodeInstances:SubprocessInstance")
			case *BPMS_Runtime.GatewayInstance:
				xmlElement.M_flowNodeInstances_2=append(xmlElement.M_flowNodeInstances_2,new(BPMS_Runtime.XsdGatewayInstance))
				this.SerialyzeGatewayInstance(xmlElement.M_flowNodeInstances_2[len(xmlElement.M_flowNodeInstances_2)-1],v)
				log.Println("Serialyze ProcessInstance:flowNodeInstances:GatewayInstance")
			case *BPMS_Runtime.EventInstance:
				xmlElement.M_flowNodeInstances_3=append(xmlElement.M_flowNodeInstances_3,new(BPMS_Runtime.XsdEventInstance))
				this.SerialyzeEventInstance(xmlElement.M_flowNodeInstances_3[len(xmlElement.M_flowNodeInstances_3)-1],v)
				log.Println("Serialyze ProcessInstance:flowNodeInstances:EventInstance")
		}
	}

	/** Instance **/
	xmlElement.M_number= object.M_number

	/** Instance **/
	xmlElement.M_colorNumber= object.M_colorNumber

	/** Instance **/
	xmlElement.M_colorName= object.M_colorName
}
