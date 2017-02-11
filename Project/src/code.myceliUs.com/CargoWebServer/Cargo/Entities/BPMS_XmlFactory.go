package Entities

import (
	"encoding/xml"
	"log"
	"os"
	"path/filepath"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/BPMS"
	"code.myceliUs.com/Utility"
	"golang.org/x/net/html/charset"
)

type BPMS_XmlFactory struct {
	m_references map[string]interface{}
	m_object     map[string]map[string][]string
}

/** Initialization function from xml file **/
func (this *BPMS_XmlFactory) InitXml(inputPath string, object *BPMS.Runtimes) error {
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
	var xmlElement *BPMS.XsdRuntimes
	xmlElement = new(BPMS.XsdRuntimes)
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
				for i := 0; i < len(refs); i++ {
					ref := this.m_references[refs[i]]
					if ref != nil {
						params := make([]interface{}, 0)
						params = append(params, ref)
						methodName := "Set" + strings.ToUpper(ref1[0:1]) + ref1[1:]
						Utility.CallMethod(refOwner, methodName, params)
					} else {
						params := make([]interface{}, 0)
						params = append(params, refs[i])
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
func (this *BPMS_XmlFactory) SerializeXml(outputPath string, toSerialize *BPMS.Runtimes) error {
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
	var xmlElement *BPMS.XsdRuntimes
	xmlElement = new(BPMS.XsdRuntimes)

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

/** inititialisation of EventData **/
func (this *BPMS_XmlFactory) InitEventData(xmlElement *BPMS.XsdEventData, object *BPMS.EventData) {
	log.Println("Initialize EventData")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.EventData%" + Utility.RandomUUID()
	}

	/** EventData **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ActivityInstance **/
func (this *BPMS_XmlFactory) InitActivityInstance(xmlElement *BPMS.XsdActivityInstance, object *BPMS.ActivityInstance) {
	log.Println("Initialize ActivityInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.ActivityInstance%" + Utility.RandomUUID()
	}

	/** ActivityInstance **/
	object.M_id = xmlElement.M_id

	/** Init ref bpmnElementId **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_bpmnElementId) > 0 {
		if _, ok := this.m_object[object.M_id]["bpmnElementId"]; !ok {
			this.m_object[object.M_id]["bpmnElementId"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElementId"] = append(this.m_object[object.M_id]["bpmnElementId"], xmlElement.M_bpmnElementId)
	}

	/** Init ref inputRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_inputRef); i++ {
		if _, ok := this.m_object[object.M_id]["inputRef"]; !ok {
			this.m_object[object.M_id]["inputRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["inputRef"] = append(this.m_object[object.M_id]["inputRef"], xmlElement.M_inputRef[i])
	}

	/** Init ref outputRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outputRef); i++ {
		if _, ok := this.m_object[object.M_id]["outputRef"]; !ok {
			this.m_object[object.M_id]["outputRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outputRef"] = append(this.m_object[object.M_id]["outputRef"], xmlElement.M_outputRef[i])
	}

	/** FlowNodeType **/
	if xmlElement.M_flowNodeType == "##AbstractTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_AbstractTask
	} else if xmlElement.M_flowNodeType == "##ServiceTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ServiceTask
	} else if xmlElement.M_flowNodeType == "##UserTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_UserTask
	} else if xmlElement.M_flowNodeType == "##ManualTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ManualTask
	} else if xmlElement.M_flowNodeType == "##BusinessRuleTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_BusinessRuleTask
	} else if xmlElement.M_flowNodeType == "##ScriptTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ScriptTask
	} else if xmlElement.M_flowNodeType == "##EmbeddedSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_EmbeddedSubprocess
	} else if xmlElement.M_flowNodeType == "##EventSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_EventSubprocess
	} else if xmlElement.M_flowNodeType == "##AdHocSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_AdHocSubprocess
	} else if xmlElement.M_flowNodeType == "##Transaction" {
		object.M_flowNodeType = BPMS.FlowNodeType_Transaction
	} else if xmlElement.M_flowNodeType == "##CallActivity" {
		object.M_flowNodeType = BPMS.FlowNodeType_CallActivity
	} else if xmlElement.M_flowNodeType == "##ParallelGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ParallelGateway
	} else if xmlElement.M_flowNodeType == "##ExclusiveGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ExclusiveGateway
	} else if xmlElement.M_flowNodeType == "##InclusiveGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_InclusiveGateway
	} else if xmlElement.M_flowNodeType == "##EventBasedGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_EventBasedGateway
	} else if xmlElement.M_flowNodeType == "##ComplexGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ComplexGateway
	} else if xmlElement.M_flowNodeType == "##StartEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_StartEvent
	} else if xmlElement.M_flowNodeType == "##IntermediateCatchEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_IntermediateCatchEvent
	} else if xmlElement.M_flowNodeType == "##BoundaryEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_BoundaryEvent
	} else if xmlElement.M_flowNodeType == "##EndEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_EndEvent
	} else if xmlElement.M_flowNodeType == "##IntermediateThrowEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_IntermediateThrowEvent
	}

	/** LifecycleState **/
	if xmlElement.M_lifecycleState == "##Completed" {
		object.M_lifecycleState = BPMS.LifecycleState_Completed
	} else if xmlElement.M_lifecycleState == "##Compensated" {
		object.M_lifecycleState = BPMS.LifecycleState_Compensated
	} else if xmlElement.M_lifecycleState == "##Failed" {
		object.M_lifecycleState = BPMS.LifecycleState_Failed
	} else if xmlElement.M_lifecycleState == "##Terminated" {
		object.M_lifecycleState = BPMS.LifecycleState_Terminated
	} else if xmlElement.M_lifecycleState == "##Ready" {
		object.M_lifecycleState = BPMS.LifecycleState_Ready
	} else if xmlElement.M_lifecycleState == "##Active" {
		object.M_lifecycleState = BPMS.LifecycleState_Active
	} else if xmlElement.M_lifecycleState == "##Completing" {
		object.M_lifecycleState = BPMS.LifecycleState_Completing
	} else if xmlElement.M_lifecycleState == "##Compensating" {
		object.M_lifecycleState = BPMS.LifecycleState_Compensating
	} else if xmlElement.M_lifecycleState == "##Failing" {
		object.M_lifecycleState = BPMS.LifecycleState_Failing
	} else if xmlElement.M_lifecycleState == "##Terminating" {
		object.M_lifecycleState = BPMS.LifecycleState_Terminating
	}

	/** ActivityType **/
	if xmlElement.M_activityType == "##AbstractTask" {
		object.M_activityType = BPMS.ActivityType_AbstractTask
	} else if xmlElement.M_activityType == "##ServiceTask" {
		object.M_activityType = BPMS.ActivityType_ServiceTask
	} else if xmlElement.M_activityType == "##UserTask" {
		object.M_activityType = BPMS.ActivityType_UserTask
	} else if xmlElement.M_activityType == "##ManualTask" {
		object.M_activityType = BPMS.ActivityType_ManualTask
	} else if xmlElement.M_activityType == "##BusinessRuleTask" {
		object.M_activityType = BPMS.ActivityType_BusinessRuleTask
	} else if xmlElement.M_activityType == "##ScriptTask" {
		object.M_activityType = BPMS.ActivityType_ScriptTask
	} else if xmlElement.M_activityType == "##EmbeddedSubprocess" {
		object.M_activityType = BPMS.ActivityType_EmbeddedSubprocess
	} else if xmlElement.M_activityType == "##EventSubprocess" {
		object.M_activityType = BPMS.ActivityType_EventSubprocess
	} else if xmlElement.M_activityType == "##AdHocSubprocess" {
		object.M_activityType = BPMS.ActivityType_AdHocSubprocess
	} else if xmlElement.M_activityType == "##Transaction" {
		object.M_activityType = BPMS.ActivityType_Transaction
	} else if xmlElement.M_activityType == "##CallActivity" {
		object.M_activityType = BPMS.ActivityType_CallActivity
	}

	/** LoopCharacteristicType **/
	if xmlElement.M_loopCharacteristicType == "##StandardLoopCharacteristics" {
		object.M_loopCharacteristicType = BPMS.LoopCharacteristicType_StandardLoopCharacteristics
	} else if xmlElement.M_loopCharacteristicType == "##MultiInstanceLoopCharacteristics" {
		object.M_loopCharacteristicType = BPMS.LoopCharacteristicType_MultiInstanceLoopCharacteristics
	}

	/** MultiInstanceBehaviorType **/
	if xmlElement.M_multiInstanceBehaviorType == "##None" {
		object.M_multiInstanceBehaviorType = BPMS.MultiInstanceBehaviorType_None
	} else if xmlElement.M_multiInstanceBehaviorType == "##One" {
		object.M_multiInstanceBehaviorType = BPMS.MultiInstanceBehaviorType_One
	} else if xmlElement.M_multiInstanceBehaviorType == "##All" {
		object.M_multiInstanceBehaviorType = BPMS.MultiInstanceBehaviorType_All
	} else if xmlElement.M_multiInstanceBehaviorType == "##Complex" {
		object.M_multiInstanceBehaviorType = BPMS.MultiInstanceBehaviorType_Complex
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of GatewayInstance **/
func (this *BPMS_XmlFactory) InitGatewayInstance(xmlElement *BPMS.XsdGatewayInstance, object *BPMS.GatewayInstance) {
	log.Println("Initialize GatewayInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.GatewayInstance%" + Utility.RandomUUID()
	}

	/** GatewayInstance **/
	object.M_id = xmlElement.M_id

	/** Init ref bpmnElementId **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_bpmnElementId) > 0 {
		if _, ok := this.m_object[object.M_id]["bpmnElementId"]; !ok {
			this.m_object[object.M_id]["bpmnElementId"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElementId"] = append(this.m_object[object.M_id]["bpmnElementId"], xmlElement.M_bpmnElementId)
	}

	/** Init ref inputRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_inputRef); i++ {
		if _, ok := this.m_object[object.M_id]["inputRef"]; !ok {
			this.m_object[object.M_id]["inputRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["inputRef"] = append(this.m_object[object.M_id]["inputRef"], xmlElement.M_inputRef[i])
	}

	/** Init ref outputRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outputRef); i++ {
		if _, ok := this.m_object[object.M_id]["outputRef"]; !ok {
			this.m_object[object.M_id]["outputRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outputRef"] = append(this.m_object[object.M_id]["outputRef"], xmlElement.M_outputRef[i])
	}

	/** FlowNodeType **/
	if xmlElement.M_flowNodeType == "##AbstractTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_AbstractTask
	} else if xmlElement.M_flowNodeType == "##ServiceTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ServiceTask
	} else if xmlElement.M_flowNodeType == "##UserTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_UserTask
	} else if xmlElement.M_flowNodeType == "##ManualTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ManualTask
	} else if xmlElement.M_flowNodeType == "##BusinessRuleTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_BusinessRuleTask
	} else if xmlElement.M_flowNodeType == "##ScriptTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ScriptTask
	} else if xmlElement.M_flowNodeType == "##EmbeddedSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_EmbeddedSubprocess
	} else if xmlElement.M_flowNodeType == "##EventSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_EventSubprocess
	} else if xmlElement.M_flowNodeType == "##AdHocSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_AdHocSubprocess
	} else if xmlElement.M_flowNodeType == "##Transaction" {
		object.M_flowNodeType = BPMS.FlowNodeType_Transaction
	} else if xmlElement.M_flowNodeType == "##CallActivity" {
		object.M_flowNodeType = BPMS.FlowNodeType_CallActivity
	} else if xmlElement.M_flowNodeType == "##ParallelGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ParallelGateway
	} else if xmlElement.M_flowNodeType == "##ExclusiveGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ExclusiveGateway
	} else if xmlElement.M_flowNodeType == "##InclusiveGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_InclusiveGateway
	} else if xmlElement.M_flowNodeType == "##EventBasedGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_EventBasedGateway
	} else if xmlElement.M_flowNodeType == "##ComplexGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ComplexGateway
	} else if xmlElement.M_flowNodeType == "##StartEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_StartEvent
	} else if xmlElement.M_flowNodeType == "##IntermediateCatchEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_IntermediateCatchEvent
	} else if xmlElement.M_flowNodeType == "##BoundaryEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_BoundaryEvent
	} else if xmlElement.M_flowNodeType == "##EndEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_EndEvent
	} else if xmlElement.M_flowNodeType == "##IntermediateThrowEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_IntermediateThrowEvent
	}

	/** LifecycleState **/
	if xmlElement.M_lifecycleState == "##Completed" {
		object.M_lifecycleState = BPMS.LifecycleState_Completed
	} else if xmlElement.M_lifecycleState == "##Compensated" {
		object.M_lifecycleState = BPMS.LifecycleState_Compensated
	} else if xmlElement.M_lifecycleState == "##Failed" {
		object.M_lifecycleState = BPMS.LifecycleState_Failed
	} else if xmlElement.M_lifecycleState == "##Terminated" {
		object.M_lifecycleState = BPMS.LifecycleState_Terminated
	} else if xmlElement.M_lifecycleState == "##Ready" {
		object.M_lifecycleState = BPMS.LifecycleState_Ready
	} else if xmlElement.M_lifecycleState == "##Active" {
		object.M_lifecycleState = BPMS.LifecycleState_Active
	} else if xmlElement.M_lifecycleState == "##Completing" {
		object.M_lifecycleState = BPMS.LifecycleState_Completing
	} else if xmlElement.M_lifecycleState == "##Compensating" {
		object.M_lifecycleState = BPMS.LifecycleState_Compensating
	} else if xmlElement.M_lifecycleState == "##Failing" {
		object.M_lifecycleState = BPMS.LifecycleState_Failing
	} else if xmlElement.M_lifecycleState == "##Terminating" {
		object.M_lifecycleState = BPMS.LifecycleState_Terminating
	}

	/** GatewayType **/
	if xmlElement.M_gatewayType == "##ParallelGateway" {
		object.M_gatewayType = BPMS.GatewayType_ParallelGateway
	} else if xmlElement.M_gatewayType == "##ExclusiveGateway" {
		object.M_gatewayType = BPMS.GatewayType_ExclusiveGateway
	} else if xmlElement.M_gatewayType == "##InclusiveGateway" {
		object.M_gatewayType = BPMS.GatewayType_InclusiveGateway
	} else if xmlElement.M_gatewayType == "##EventBasedGateway" {
		object.M_gatewayType = BPMS.GatewayType_EventBasedGateway
	} else if xmlElement.M_gatewayType == "##ComplexGateway" {
		object.M_gatewayType = BPMS.GatewayType_ComplexGateway
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of EventInstance **/
func (this *BPMS_XmlFactory) InitEventInstance(xmlElement *BPMS.XsdEventInstance, object *BPMS.EventInstance) {
	log.Println("Initialize EventInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.EventInstance%" + Utility.RandomUUID()
	}

	/** EventInstance **/
	object.M_id = xmlElement.M_id

	/** Init ref bpmnElementId **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_bpmnElementId) > 0 {
		if _, ok := this.m_object[object.M_id]["bpmnElementId"]; !ok {
			this.m_object[object.M_id]["bpmnElementId"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElementId"] = append(this.m_object[object.M_id]["bpmnElementId"], xmlElement.M_bpmnElementId)
	}

	/** Init ref inputRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_inputRef); i++ {
		if _, ok := this.m_object[object.M_id]["inputRef"]; !ok {
			this.m_object[object.M_id]["inputRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["inputRef"] = append(this.m_object[object.M_id]["inputRef"], xmlElement.M_inputRef[i])
	}

	/** Init ref outputRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outputRef); i++ {
		if _, ok := this.m_object[object.M_id]["outputRef"]; !ok {
			this.m_object[object.M_id]["outputRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outputRef"] = append(this.m_object[object.M_id]["outputRef"], xmlElement.M_outputRef[i])
	}

	/** FlowNodeType **/
	if xmlElement.M_flowNodeType == "##AbstractTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_AbstractTask
	} else if xmlElement.M_flowNodeType == "##ServiceTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ServiceTask
	} else if xmlElement.M_flowNodeType == "##UserTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_UserTask
	} else if xmlElement.M_flowNodeType == "##ManualTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ManualTask
	} else if xmlElement.M_flowNodeType == "##BusinessRuleTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_BusinessRuleTask
	} else if xmlElement.M_flowNodeType == "##ScriptTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ScriptTask
	} else if xmlElement.M_flowNodeType == "##EmbeddedSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_EmbeddedSubprocess
	} else if xmlElement.M_flowNodeType == "##EventSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_EventSubprocess
	} else if xmlElement.M_flowNodeType == "##AdHocSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_AdHocSubprocess
	} else if xmlElement.M_flowNodeType == "##Transaction" {
		object.M_flowNodeType = BPMS.FlowNodeType_Transaction
	} else if xmlElement.M_flowNodeType == "##CallActivity" {
		object.M_flowNodeType = BPMS.FlowNodeType_CallActivity
	} else if xmlElement.M_flowNodeType == "##ParallelGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ParallelGateway
	} else if xmlElement.M_flowNodeType == "##ExclusiveGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ExclusiveGateway
	} else if xmlElement.M_flowNodeType == "##InclusiveGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_InclusiveGateway
	} else if xmlElement.M_flowNodeType == "##EventBasedGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_EventBasedGateway
	} else if xmlElement.M_flowNodeType == "##ComplexGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ComplexGateway
	} else if xmlElement.M_flowNodeType == "##StartEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_StartEvent
	} else if xmlElement.M_flowNodeType == "##IntermediateCatchEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_IntermediateCatchEvent
	} else if xmlElement.M_flowNodeType == "##BoundaryEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_BoundaryEvent
	} else if xmlElement.M_flowNodeType == "##EndEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_EndEvent
	} else if xmlElement.M_flowNodeType == "##IntermediateThrowEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_IntermediateThrowEvent
	}

	/** LifecycleState **/
	if xmlElement.M_lifecycleState == "##Completed" {
		object.M_lifecycleState = BPMS.LifecycleState_Completed
	} else if xmlElement.M_lifecycleState == "##Compensated" {
		object.M_lifecycleState = BPMS.LifecycleState_Compensated
	} else if xmlElement.M_lifecycleState == "##Failed" {
		object.M_lifecycleState = BPMS.LifecycleState_Failed
	} else if xmlElement.M_lifecycleState == "##Terminated" {
		object.M_lifecycleState = BPMS.LifecycleState_Terminated
	} else if xmlElement.M_lifecycleState == "##Ready" {
		object.M_lifecycleState = BPMS.LifecycleState_Ready
	} else if xmlElement.M_lifecycleState == "##Active" {
		object.M_lifecycleState = BPMS.LifecycleState_Active
	} else if xmlElement.M_lifecycleState == "##Completing" {
		object.M_lifecycleState = BPMS.LifecycleState_Completing
	} else if xmlElement.M_lifecycleState == "##Compensating" {
		object.M_lifecycleState = BPMS.LifecycleState_Compensating
	} else if xmlElement.M_lifecycleState == "##Failing" {
		object.M_lifecycleState = BPMS.LifecycleState_Failing
	} else if xmlElement.M_lifecycleState == "##Terminating" {
		object.M_lifecycleState = BPMS.LifecycleState_Terminating
	}

	/** Init sourceRef **/
	object.M_eventDefintionInstances = make([]*BPMS.EventDefinitionInstance, 0)
	for i := 0; i < len(xmlElement.M_eventDefintionInstances); i++ {
		val := new(BPMS.EventDefinitionInstance)
		this.InitEventDefinitionInstance(xmlElement.M_eventDefintionInstances[i], val)
		object.M_eventDefintionInstances = append(object.M_eventDefintionInstances, val)

		/** association initialisation **/
		val.SetEventInstancePtr(object)
	}

	/** EventType **/
	if xmlElement.M_eventType == "##StartEvent" {
		object.M_eventType = BPMS.EventType_StartEvent
	} else if xmlElement.M_eventType == "##IntermediateCatchEvent" {
		object.M_eventType = BPMS.EventType_IntermediateCatchEvent
	} else if xmlElement.M_eventType == "##BoundaryEvent" {
		object.M_eventType = BPMS.EventType_BoundaryEvent
	} else if xmlElement.M_eventType == "##EndEvent" {
		object.M_eventType = BPMS.EventType_EndEvent
	} else if xmlElement.M_eventType == "##IntermediateThrowEvent" {
		object.M_eventType = BPMS.EventType_IntermediateThrowEvent
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Exception **/
func (this *BPMS_XmlFactory) InitException(xmlElement *BPMS.XsdException, object *BPMS.Exception) {
	log.Println("Initialize Exception")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.Exception%" + Utility.RandomUUID()
	}

	/** ExceptionType **/
	if xmlElement.M_exceptionType == "##GatewayException" {
		object.M_exceptionType = BPMS.ExceptionType_GatewayException
	} else if xmlElement.M_exceptionType == "##NoIORuleException" {
		object.M_exceptionType = BPMS.ExceptionType_NoIORuleException
	} else if xmlElement.M_exceptionType == "##NoAvailableOutputSetException" {
		object.M_exceptionType = BPMS.ExceptionType_NoAvailableOutputSetException
	} else if xmlElement.M_exceptionType == "##NotMatchingIOSpecification" {
		object.M_exceptionType = BPMS.ExceptionType_NotMatchingIOSpecification
	} else if xmlElement.M_exceptionType == "##IllegalStartEventException" {
		object.M_exceptionType = BPMS.ExceptionType_IllegalStartEventException
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Trigger **/
func (this *BPMS_XmlFactory) InitTrigger(xmlElement *BPMS.XsdTrigger, object *BPMS.Trigger) {
	log.Println("Initialize Trigger")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.Trigger%" + Utility.RandomUUID()
	}

	/** Init eventData **/
	object.M_eventDatas = make([]*BPMS.EventData, 0)
	for i := 0; i < len(xmlElement.M_eventDatas); i++ {
		val := new(BPMS.EventData)
		this.InitEventData(xmlElement.M_eventDatas[i], val)
		object.M_eventDatas = append(object.M_eventDatas, val)

		/** association initialisation **/
		val.SetTriggerPtr(object)
	}

	/** Init ref dataRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_dataRef); i++ {
		if _, ok := this.m_object[object.M_id]["dataRef"]; !ok {
			this.m_object[object.M_id]["dataRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["dataRef"] = append(this.m_object[object.M_id]["dataRef"], xmlElement.M_dataRef[i])
	}

	/** Init ref sourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_sourceRef != nil {
		if _, ok := this.m_object[object.M_id]["sourceRef"]; !ok {
			this.m_object[object.M_id]["sourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["sourceRef"] = append(this.m_object[object.M_id]["sourceRef"], *xmlElement.M_sourceRef)
	}

	/** Init ref targetRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_targetRef != nil {
		if _, ok := this.m_object[object.M_id]["targetRef"]; !ok {
			this.m_object[object.M_id]["targetRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["targetRef"] = append(this.m_object[object.M_id]["targetRef"], *xmlElement.M_targetRef)
	}

	/** Trigger **/
	object.M_id = xmlElement.M_id

	/** Trigger **/
	object.M_processUUID = xmlElement.M_processUUID

	/** Trigger **/
	object.M_sessionId = xmlElement.M_sessionId

	/** EventTriggerType **/
	if xmlElement.M_eventTriggerType == "##None" {
		object.M_eventTriggerType = BPMS.EventTriggerType_None
	} else if xmlElement.M_eventTriggerType == "##Timer" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Timer
	} else if xmlElement.M_eventTriggerType == "##Conditional" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Conditional
	} else if xmlElement.M_eventTriggerType == "##Message" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Message
	} else if xmlElement.M_eventTriggerType == "##Signal" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Signal
	} else if xmlElement.M_eventTriggerType == "##Multiple" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Multiple
	} else if xmlElement.M_eventTriggerType == "##ParallelMultiple" {
		object.M_eventTriggerType = BPMS.EventTriggerType_ParallelMultiple
	} else if xmlElement.M_eventTriggerType == "##Escalation" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Escalation
	} else if xmlElement.M_eventTriggerType == "##Error" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Error
	} else if xmlElement.M_eventTriggerType == "##Compensation" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Compensation
	} else if xmlElement.M_eventTriggerType == "##Terminate" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Terminate
	} else if xmlElement.M_eventTriggerType == "##Cancel" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Cancel
	} else if xmlElement.M_eventTriggerType == "##Link" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Link
	} else if xmlElement.M_eventTriggerType == "##Start" {
		object.M_eventTriggerType = BPMS.EventTriggerType_Start
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ItemAwareElementInstance **/
func (this *BPMS_XmlFactory) InitItemAwareElementInstance(xmlElement *BPMS.XsdItemAwareElementInstance, object *BPMS.ItemAwareElementInstance) {
	log.Println("Initialize ItemAwareElementInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.ItemAwareElementInstance%" + Utility.RandomUUID()
	}

	/** ItemAwareElementInstance **/
	object.M_id = xmlElement.M_id

	/** ItemAwareElementInstance **/
	object.M_bpmnElementId = xmlElement.M_bpmnElementId

	/** ItemAwareElementInstance **/
	object.M_data = xmlElement.M_data
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of CorrelationInfo **/
func (this *BPMS_XmlFactory) InitCorrelationInfo(xmlElement *BPMS.XsdCorrelationInfo, object *BPMS.CorrelationInfo) {
	log.Println("Initialize CorrelationInfo")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.CorrelationInfo%" + Utility.RandomUUID()
	}

	/** CorrelationInfo **/
	object.M_id = xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ConnectingObject **/
func (this *BPMS_XmlFactory) InitConnectingObject(xmlElement *BPMS.XsdConnectingObject, object *BPMS.ConnectingObject) {
	log.Println("Initialize ConnectingObject")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.ConnectingObject%" + Utility.RandomUUID()
	}

	/** ConnectingObject **/
	object.M_id = xmlElement.M_id

	/** Init ref bpmnElementId **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_bpmnElementId) > 0 {
		if _, ok := this.m_object[object.M_id]["bpmnElementId"]; !ok {
			this.m_object[object.M_id]["bpmnElementId"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElementId"] = append(this.m_object[object.M_id]["bpmnElementId"], xmlElement.M_bpmnElementId)
	}

	/** Init ref sourceRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_sourceRef != nil {
		if _, ok := this.m_object[object.M_id]["sourceRef"]; !ok {
			this.m_object[object.M_id]["sourceRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["sourceRef"] = append(this.m_object[object.M_id]["sourceRef"], *xmlElement.M_sourceRef)
	}

	/** Init ref targetRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if xmlElement.M_targetRef != nil {
		if _, ok := this.m_object[object.M_id]["targetRef"]; !ok {
			this.m_object[object.M_id]["targetRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["targetRef"] = append(this.m_object[object.M_id]["targetRef"], *xmlElement.M_targetRef)
	}

	/** ConnectingObjectType **/
	if xmlElement.M_connectingObjectType == "##SequenceFlow" {
		object.M_connectingObjectType = BPMS.ConnectingObjectType_SequenceFlow
	} else if xmlElement.M_connectingObjectType == "##MessageFlow" {
		object.M_connectingObjectType = BPMS.ConnectingObjectType_MessageFlow
	} else if xmlElement.M_connectingObjectType == "##Association" {
		object.M_connectingObjectType = BPMS.ConnectingObjectType_Association
	} else if xmlElement.M_connectingObjectType == "##DataAssociation" {
		object.M_connectingObjectType = BPMS.ConnectingObjectType_DataAssociation
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of DefinitionsInstance **/
func (this *BPMS_XmlFactory) InitDefinitionsInstance(xmlElement *BPMS.XsdDefinitionsInstance, object *BPMS.DefinitionsInstance) {
	log.Println("Initialize DefinitionsInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.DefinitionsInstance%" + Utility.RandomUUID()
	}

	/** DefinitionsInstance **/
	object.M_id = xmlElement.M_id

	/** Init ref bpmnElementId **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_bpmnElementId) > 0 {
		if _, ok := this.m_object[object.M_id]["bpmnElementId"]; !ok {
			this.m_object[object.M_id]["bpmnElementId"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElementId"] = append(this.m_object[object.M_id]["bpmnElementId"], xmlElement.M_bpmnElementId)
	}

	/** Init process **/
	object.M_processInstances = make([]*BPMS.ProcessInstance, 0)
	for i := 0; i < len(xmlElement.M_processInstances); i++ {
		val := new(BPMS.ProcessInstance)
		this.InitProcessInstance(xmlElement.M_processInstances[i], val)
		object.M_processInstances = append(object.M_processInstances, val)

		/** association initialisation **/
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of ProcessInstance **/
func (this *BPMS_XmlFactory) InitProcessInstance(xmlElement *BPMS.XsdProcessInstance, object *BPMS.ProcessInstance) {
	log.Println("Initialize ProcessInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.ProcessInstance%" + Utility.RandomUUID()
	}

	/** ProcessInstance **/
	object.M_id = xmlElement.M_id

	/** Init ref bpmnElementId **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_bpmnElementId) > 0 {
		if _, ok := this.m_object[object.M_id]["bpmnElementId"]; !ok {
			this.m_object[object.M_id]["bpmnElementId"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElementId"] = append(this.m_object[object.M_id]["bpmnElementId"], xmlElement.M_bpmnElementId)
	}

	/** Init activityInstance **/
	object.M_flowNodeInstances = make([]BPMS.FlowNodeInstance, 0)
	for i := 0; i < len(xmlElement.M_flowNodeInstances_0); i++ {
		val := new(BPMS.ActivityInstance)
		this.InitActivityInstance(xmlElement.M_flowNodeInstances_0[i], val)
		object.M_flowNodeInstances = append(object.M_flowNodeInstances, val)

		/** association initialisation **/
		val.SetProcessInstancePtr(object)
	}

	/** Init SubprocessInstance **/
	for i := 0; i < len(xmlElement.M_flowNodeInstances_1); i++ {
		val := new(BPMS.SubprocessInstance)
		this.InitSubprocessInstance(xmlElement.M_flowNodeInstances_1[i], val)
		object.M_flowNodeInstances = append(object.M_flowNodeInstances, val)

		/** association initialisation **/
		val.SetProcessInstancePtr(object)
	}

	/** Init gatewayInstance **/
	for i := 0; i < len(xmlElement.M_flowNodeInstances_2); i++ {
		val := new(BPMS.GatewayInstance)
		this.InitGatewayInstance(xmlElement.M_flowNodeInstances_2[i], val)
		object.M_flowNodeInstances = append(object.M_flowNodeInstances, val)

		/** association initialisation **/
		val.SetProcessInstancePtr(object)
	}

	/** Init eventInstance **/
	for i := 0; i < len(xmlElement.M_flowNodeInstances_3); i++ {
		val := new(BPMS.EventInstance)
		this.InitEventInstance(xmlElement.M_flowNodeInstances_3[i], val)
		object.M_flowNodeInstances = append(object.M_flowNodeInstances, val)

		/** association initialisation **/
		val.SetProcessInstancePtr(object)
	}

	/** Instance **/
	object.M_number = xmlElement.M_number

	/** Instance **/
	object.M_colorNumber = xmlElement.M_colorNumber

	/** Instance **/
	object.M_colorName = xmlElement.M_colorName
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of SubprocessInstance **/
func (this *BPMS_XmlFactory) InitSubprocessInstance(xmlElement *BPMS.XsdSubprocessInstance, object *BPMS.SubprocessInstance) {
	log.Println("Initialize SubprocessInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.SubprocessInstance%" + Utility.RandomUUID()
	}

	/** SubprocessInstance **/
	object.M_id = xmlElement.M_id

	/** Init ref bpmnElementId **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_bpmnElementId) > 0 {
		if _, ok := this.m_object[object.M_id]["bpmnElementId"]; !ok {
			this.m_object[object.M_id]["bpmnElementId"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElementId"] = append(this.m_object[object.M_id]["bpmnElementId"], xmlElement.M_bpmnElementId)
	}

	/** Init ref inputRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_inputRef); i++ {
		if _, ok := this.m_object[object.M_id]["inputRef"]; !ok {
			this.m_object[object.M_id]["inputRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["inputRef"] = append(this.m_object[object.M_id]["inputRef"], xmlElement.M_inputRef[i])
	}

	/** Init ref outputRef **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	for i := 0; i < len(xmlElement.M_outputRef); i++ {
		if _, ok := this.m_object[object.M_id]["outputRef"]; !ok {
			this.m_object[object.M_id]["outputRef"] = make([]string, 0)
		}
		this.m_object[object.M_id]["outputRef"] = append(this.m_object[object.M_id]["outputRef"], xmlElement.M_outputRef[i])
	}

	/** FlowNodeType **/
	if xmlElement.M_flowNodeType == "##AbstractTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_AbstractTask
	} else if xmlElement.M_flowNodeType == "##ServiceTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ServiceTask
	} else if xmlElement.M_flowNodeType == "##UserTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_UserTask
	} else if xmlElement.M_flowNodeType == "##ManualTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ManualTask
	} else if xmlElement.M_flowNodeType == "##BusinessRuleTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_BusinessRuleTask
	} else if xmlElement.M_flowNodeType == "##ScriptTask" {
		object.M_flowNodeType = BPMS.FlowNodeType_ScriptTask
	} else if xmlElement.M_flowNodeType == "##EmbeddedSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_EmbeddedSubprocess
	} else if xmlElement.M_flowNodeType == "##EventSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_EventSubprocess
	} else if xmlElement.M_flowNodeType == "##AdHocSubprocess" {
		object.M_flowNodeType = BPMS.FlowNodeType_AdHocSubprocess
	} else if xmlElement.M_flowNodeType == "##Transaction" {
		object.M_flowNodeType = BPMS.FlowNodeType_Transaction
	} else if xmlElement.M_flowNodeType == "##CallActivity" {
		object.M_flowNodeType = BPMS.FlowNodeType_CallActivity
	} else if xmlElement.M_flowNodeType == "##ParallelGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ParallelGateway
	} else if xmlElement.M_flowNodeType == "##ExclusiveGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ExclusiveGateway
	} else if xmlElement.M_flowNodeType == "##InclusiveGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_InclusiveGateway
	} else if xmlElement.M_flowNodeType == "##EventBasedGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_EventBasedGateway
	} else if xmlElement.M_flowNodeType == "##ComplexGateway" {
		object.M_flowNodeType = BPMS.FlowNodeType_ComplexGateway
	} else if xmlElement.M_flowNodeType == "##StartEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_StartEvent
	} else if xmlElement.M_flowNodeType == "##IntermediateCatchEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_IntermediateCatchEvent
	} else if xmlElement.M_flowNodeType == "##BoundaryEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_BoundaryEvent
	} else if xmlElement.M_flowNodeType == "##EndEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_EndEvent
	} else if xmlElement.M_flowNodeType == "##IntermediateThrowEvent" {
		object.M_flowNodeType = BPMS.FlowNodeType_IntermediateThrowEvent
	}

	/** LifecycleState **/
	if xmlElement.M_lifecycleState == "##Completed" {
		object.M_lifecycleState = BPMS.LifecycleState_Completed
	} else if xmlElement.M_lifecycleState == "##Compensated" {
		object.M_lifecycleState = BPMS.LifecycleState_Compensated
	} else if xmlElement.M_lifecycleState == "##Failed" {
		object.M_lifecycleState = BPMS.LifecycleState_Failed
	} else if xmlElement.M_lifecycleState == "##Terminated" {
		object.M_lifecycleState = BPMS.LifecycleState_Terminated
	} else if xmlElement.M_lifecycleState == "##Ready" {
		object.M_lifecycleState = BPMS.LifecycleState_Ready
	} else if xmlElement.M_lifecycleState == "##Active" {
		object.M_lifecycleState = BPMS.LifecycleState_Active
	} else if xmlElement.M_lifecycleState == "##Completing" {
		object.M_lifecycleState = BPMS.LifecycleState_Completing
	} else if xmlElement.M_lifecycleState == "##Compensating" {
		object.M_lifecycleState = BPMS.LifecycleState_Compensating
	} else if xmlElement.M_lifecycleState == "##Failing" {
		object.M_lifecycleState = BPMS.LifecycleState_Failing
	} else if xmlElement.M_lifecycleState == "##Terminating" {
		object.M_lifecycleState = BPMS.LifecycleState_Terminating
	}

	/** Init activityInstance **/
	object.M_flowNodeInstances = make([]BPMS.FlowNodeInstance, 0)
	for i := 0; i < len(xmlElement.M_flowNodeInstances_0); i++ {
		val := new(BPMS.ActivityInstance)
		this.InitActivityInstance(xmlElement.M_flowNodeInstances_0[i], val)
		object.M_flowNodeInstances = append(object.M_flowNodeInstances, val)

		/** association initialisation **/
	}

	/** Init SubprocessInstance **/
	for i := 0; i < len(xmlElement.M_flowNodeInstances_1); i++ {
		val := new(BPMS.SubprocessInstance)
		this.InitSubprocessInstance(xmlElement.M_flowNodeInstances_1[i], val)
		object.M_flowNodeInstances = append(object.M_flowNodeInstances, val)

		/** association initialisation **/
	}

	/** Init gatewayInstance **/
	for i := 0; i < len(xmlElement.M_flowNodeInstances_2); i++ {
		val := new(BPMS.GatewayInstance)
		this.InitGatewayInstance(xmlElement.M_flowNodeInstances_2[i], val)
		object.M_flowNodeInstances = append(object.M_flowNodeInstances, val)

		/** association initialisation **/
	}

	/** Init eventInstance **/
	for i := 0; i < len(xmlElement.M_flowNodeInstances_3); i++ {
		val := new(BPMS.EventInstance)
		this.InitEventInstance(xmlElement.M_flowNodeInstances_3[i], val)
		object.M_flowNodeInstances = append(object.M_flowNodeInstances, val)

		/** association initialisation **/
	}

	/** Init outputRef **/
	object.M_connectingObjects = make([]*BPMS.ConnectingObject, 0)
	for i := 0; i < len(xmlElement.M_connectingObjects); i++ {
		val := new(BPMS.ConnectingObject)
		this.InitConnectingObject(xmlElement.M_connectingObjects[i], val)
		object.M_connectingObjects = append(object.M_connectingObjects, val)

		/** association initialisation **/
	}

	/** SubprocessType **/
	if xmlElement.M_SubprocessType == "##EmbeddedSubprocess" {
		object.M_SubprocessType = BPMS.SubprocessType_EmbeddedSubprocess
	} else if xmlElement.M_SubprocessType == "##EventSubprocess" {
		object.M_SubprocessType = BPMS.SubprocessType_EventSubprocess
	} else if xmlElement.M_SubprocessType == "##AdHocSubprocess" {
		object.M_SubprocessType = BPMS.SubprocessType_AdHocSubprocess
	} else if xmlElement.M_SubprocessType == "##Transaction" {
		object.M_SubprocessType = BPMS.SubprocessType_Transaction
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of EventDefinitionInstance **/
func (this *BPMS_XmlFactory) InitEventDefinitionInstance(xmlElement *BPMS.XsdEventDefinitionInstance, object *BPMS.EventDefinitionInstance) {
	log.Println("Initialize EventDefinitionInstance")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.EventDefinitionInstance%" + Utility.RandomUUID()
	}

	/** EventDefinitionInstance **/
	object.M_id = xmlElement.M_id

	/** Init ref bpmnElementId **/
	if len(object.M_id) == 0 {
		object.M_id = Utility.RandomUUID()
		this.m_references[object.M_id] = object
	}
	if _, ok := this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id] = make(map[string][]string)
	}
	if len(xmlElement.M_bpmnElementId) > 0 {
		if _, ok := this.m_object[object.M_id]["bpmnElementId"]; !ok {
			this.m_object[object.M_id]["bpmnElementId"] = make([]string, 0)
		}
		this.m_object[object.M_id]["bpmnElementId"] = append(this.m_object[object.M_id]["bpmnElementId"], xmlElement.M_bpmnElementId)
	}

	/** EventDefinitionType **/
	if xmlElement.M_eventDefinitionType == "##MessageEventDefinition" {
		object.M_eventDefinitionType = BPMS.EventDefinitionType_MessageEventDefinition
	} else if xmlElement.M_eventDefinitionType == "##LinkEventDefinition" {
		object.M_eventDefinitionType = BPMS.EventDefinitionType_LinkEventDefinition
	} else if xmlElement.M_eventDefinitionType == "##ErrorEventDefinition" {
		object.M_eventDefinitionType = BPMS.EventDefinitionType_ErrorEventDefinition
	} else if xmlElement.M_eventDefinitionType == "##TerminateEventDefinition" {
		object.M_eventDefinitionType = BPMS.EventDefinitionType_TerminateEventDefinition
	} else if xmlElement.M_eventDefinitionType == "##CompensationEventDefinition" {
		object.M_eventDefinitionType = BPMS.EventDefinitionType_CompensationEventDefinition
	} else if xmlElement.M_eventDefinitionType == "##ConditionalEventDefinition" {
		object.M_eventDefinitionType = BPMS.EventDefinitionType_ConditionalEventDefinition
	} else if xmlElement.M_eventDefinitionType == "##TimerEventDefinition" {
		object.M_eventDefinitionType = BPMS.EventDefinitionType_TimerEventDefinition
	} else if xmlElement.M_eventDefinitionType == "##CancelEventDefinition" {
		object.M_eventDefinitionType = BPMS.EventDefinitionType_CancelEventDefinition
	} else if xmlElement.M_eventDefinitionType == "##EscalationEventDefinition" {
		object.M_eventDefinitionType = BPMS.EventDefinitionType_EscalationEventDefinition
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Runtimes **/
func (this *BPMS_XmlFactory) InitRuntimes(xmlElement *BPMS.XsdRuntimes, object *BPMS.Runtimes) {
	log.Println("Initialize Runtimes")
	if len(object.UUID) == 0 {
		object.UUID = "BPMS.Runtimes%" + Utility.RandomUUID()
	}

	/** Init definitions **/
	object.M_definitions = make([]*BPMS.DefinitionsInstance, 0)
	for i := 0; i < len(xmlElement.M_definitions); i++ {
		val := new(BPMS.DefinitionsInstance)
		this.InitDefinitionsInstance(xmlElement.M_definitions[i], val)
		object.M_definitions = append(object.M_definitions, val)

		/** association initialisation **/
	}

	/** Init exception **/
	object.M_exceptions = make([]*BPMS.Exception, 0)
	for i := 0; i < len(xmlElement.M_exceptions); i++ {
		val := new(BPMS.Exception)
		this.InitException(xmlElement.M_exceptions[i], val)
		object.M_exceptions = append(object.M_exceptions, val)

		/** association initialisation **/
		val.SetRuntimesPtr(object)
	}

	/** Init trigger **/
	object.M_triggers = make([]*BPMS.Trigger, 0)
	for i := 0; i < len(xmlElement.M_triggers); i++ {
		val := new(BPMS.Trigger)
		this.InitTrigger(xmlElement.M_triggers[i], val)
		object.M_triggers = append(object.M_triggers, val)

		/** association initialisation **/
		val.SetRuntimesPtr(object)
	}

	/** Init correlationInfo **/
	object.M_correlationInfos = make([]*BPMS.CorrelationInfo, 0)
	for i := 0; i < len(xmlElement.M_correlationInfos); i++ {
		val := new(BPMS.CorrelationInfo)
		this.InitCorrelationInfo(xmlElement.M_correlationInfos[i], val)
		object.M_correlationInfos = append(object.M_correlationInfos, val)

		/** association initialisation **/
		val.SetRuntimesPtr(object)
	}

	/** Runtimes **/
	object.M_id = xmlElement.M_id

	/** Runtimes **/
	object.M_name = xmlElement.M_name

	/** Runtimes **/
	object.M_version = xmlElement.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of GatewayInstance **/
func (this *BPMS_XmlFactory) SerialyzeGatewayInstance(xmlElement *BPMS.XsdGatewayInstance, object *BPMS.GatewayInstance) {
	if xmlElement == nil {
		return
	}

	/** GatewayInstance **/
	xmlElement.M_id = object.M_id

	/** Serialyze ref bpmnElementId **/
	xmlElement.M_bpmnElementId = object.M_bpmnElementId

	/** Serialyze ref inputRef **/
	xmlElement.M_inputRef = object.M_inputRef

	/** Serialyze ref outputRef **/
	xmlElement.M_outputRef = object.M_outputRef

	/** FlowNodeType **/
	if object.M_flowNodeType == BPMS.FlowNodeType_AbstractTask {
		xmlElement.M_flowNodeType = "##AbstractTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ServiceTask {
		xmlElement.M_flowNodeType = "##ServiceTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_UserTask {
		xmlElement.M_flowNodeType = "##UserTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ManualTask {
		xmlElement.M_flowNodeType = "##ManualTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_BusinessRuleTask {
		xmlElement.M_flowNodeType = "##BusinessRuleTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ScriptTask {
		xmlElement.M_flowNodeType = "##ScriptTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EmbeddedSubprocess {
		xmlElement.M_flowNodeType = "##EmbeddedSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EventSubprocess {
		xmlElement.M_flowNodeType = "##EventSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_AdHocSubprocess {
		xmlElement.M_flowNodeType = "##AdHocSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_Transaction {
		xmlElement.M_flowNodeType = "##Transaction"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_CallActivity {
		xmlElement.M_flowNodeType = "##CallActivity"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ParallelGateway {
		xmlElement.M_flowNodeType = "##ParallelGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ExclusiveGateway {
		xmlElement.M_flowNodeType = "##ExclusiveGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_InclusiveGateway {
		xmlElement.M_flowNodeType = "##InclusiveGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EventBasedGateway {
		xmlElement.M_flowNodeType = "##EventBasedGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ComplexGateway {
		xmlElement.M_flowNodeType = "##ComplexGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_StartEvent {
		xmlElement.M_flowNodeType = "##StartEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_IntermediateCatchEvent {
		xmlElement.M_flowNodeType = "##IntermediateCatchEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_BoundaryEvent {
		xmlElement.M_flowNodeType = "##BoundaryEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EndEvent {
		xmlElement.M_flowNodeType = "##EndEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_IntermediateThrowEvent {
		xmlElement.M_flowNodeType = "##IntermediateThrowEvent"
	}

	/** LifecycleState **/
	if object.M_lifecycleState == BPMS.LifecycleState_Completed {
		xmlElement.M_lifecycleState = "##Completed"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Compensated {
		xmlElement.M_lifecycleState = "##Compensated"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Failed {
		xmlElement.M_lifecycleState = "##Failed"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Terminated {
		xmlElement.M_lifecycleState = "##Terminated"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Ready {
		xmlElement.M_lifecycleState = "##Ready"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Active {
		xmlElement.M_lifecycleState = "##Active"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Completing {
		xmlElement.M_lifecycleState = "##Completing"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Compensating {
		xmlElement.M_lifecycleState = "##Compensating"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Failing {
		xmlElement.M_lifecycleState = "##Failing"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Terminating {
		xmlElement.M_lifecycleState = "##Terminating"
	}

	/** GatewayType **/
	if object.M_gatewayType == BPMS.GatewayType_ParallelGateway {
		xmlElement.M_gatewayType = "##ParallelGateway"
	} else if object.M_gatewayType == BPMS.GatewayType_ExclusiveGateway {
		xmlElement.M_gatewayType = "##ExclusiveGateway"
	} else if object.M_gatewayType == BPMS.GatewayType_InclusiveGateway {
		xmlElement.M_gatewayType = "##InclusiveGateway"
	} else if object.M_gatewayType == BPMS.GatewayType_EventBasedGateway {
		xmlElement.M_gatewayType = "##EventBasedGateway"
	} else if object.M_gatewayType == BPMS.GatewayType_ComplexGateway {
		xmlElement.M_gatewayType = "##ComplexGateway"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of EventDefinitionInstance **/
func (this *BPMS_XmlFactory) SerialyzeEventDefinitionInstance(xmlElement *BPMS.XsdEventDefinitionInstance, object *BPMS.EventDefinitionInstance) {
	if xmlElement == nil {
		return
	}

	/** EventDefinitionInstance **/
	xmlElement.M_id = object.M_id

	/** Serialyze ref bpmnElementId **/
	xmlElement.M_bpmnElementId = object.M_bpmnElementId

	/** EventDefinitionType **/
	if object.M_eventDefinitionType == BPMS.EventDefinitionType_MessageEventDefinition {
		xmlElement.M_eventDefinitionType = "##MessageEventDefinition"
	} else if object.M_eventDefinitionType == BPMS.EventDefinitionType_LinkEventDefinition {
		xmlElement.M_eventDefinitionType = "##LinkEventDefinition"
	} else if object.M_eventDefinitionType == BPMS.EventDefinitionType_ErrorEventDefinition {
		xmlElement.M_eventDefinitionType = "##ErrorEventDefinition"
	} else if object.M_eventDefinitionType == BPMS.EventDefinitionType_TerminateEventDefinition {
		xmlElement.M_eventDefinitionType = "##TerminateEventDefinition"
	} else if object.M_eventDefinitionType == BPMS.EventDefinitionType_CompensationEventDefinition {
		xmlElement.M_eventDefinitionType = "##CompensationEventDefinition"
	} else if object.M_eventDefinitionType == BPMS.EventDefinitionType_ConditionalEventDefinition {
		xmlElement.M_eventDefinitionType = "##ConditionalEventDefinition"
	} else if object.M_eventDefinitionType == BPMS.EventDefinitionType_TimerEventDefinition {
		xmlElement.M_eventDefinitionType = "##TimerEventDefinition"
	} else if object.M_eventDefinitionType == BPMS.EventDefinitionType_CancelEventDefinition {
		xmlElement.M_eventDefinitionType = "##CancelEventDefinition"
	} else if object.M_eventDefinitionType == BPMS.EventDefinitionType_EscalationEventDefinition {
		xmlElement.M_eventDefinitionType = "##EscalationEventDefinition"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of EventInstance **/
func (this *BPMS_XmlFactory) SerialyzeEventInstance(xmlElement *BPMS.XsdEventInstance, object *BPMS.EventInstance) {
	if xmlElement == nil {
		return
	}

	/** EventInstance **/
	xmlElement.M_id = object.M_id

	/** Serialyze ref bpmnElementId **/
	xmlElement.M_bpmnElementId = object.M_bpmnElementId

	/** Serialyze ref inputRef **/
	xmlElement.M_inputRef = object.M_inputRef

	/** Serialyze ref outputRef **/
	xmlElement.M_outputRef = object.M_outputRef

	/** FlowNodeType **/
	if object.M_flowNodeType == BPMS.FlowNodeType_AbstractTask {
		xmlElement.M_flowNodeType = "##AbstractTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ServiceTask {
		xmlElement.M_flowNodeType = "##ServiceTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_UserTask {
		xmlElement.M_flowNodeType = "##UserTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ManualTask {
		xmlElement.M_flowNodeType = "##ManualTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_BusinessRuleTask {
		xmlElement.M_flowNodeType = "##BusinessRuleTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ScriptTask {
		xmlElement.M_flowNodeType = "##ScriptTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EmbeddedSubprocess {
		xmlElement.M_flowNodeType = "##EmbeddedSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EventSubprocess {
		xmlElement.M_flowNodeType = "##EventSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_AdHocSubprocess {
		xmlElement.M_flowNodeType = "##AdHocSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_Transaction {
		xmlElement.M_flowNodeType = "##Transaction"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_CallActivity {
		xmlElement.M_flowNodeType = "##CallActivity"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ParallelGateway {
		xmlElement.M_flowNodeType = "##ParallelGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ExclusiveGateway {
		xmlElement.M_flowNodeType = "##ExclusiveGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_InclusiveGateway {
		xmlElement.M_flowNodeType = "##InclusiveGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EventBasedGateway {
		xmlElement.M_flowNodeType = "##EventBasedGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ComplexGateway {
		xmlElement.M_flowNodeType = "##ComplexGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_StartEvent {
		xmlElement.M_flowNodeType = "##StartEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_IntermediateCatchEvent {
		xmlElement.M_flowNodeType = "##IntermediateCatchEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_BoundaryEvent {
		xmlElement.M_flowNodeType = "##BoundaryEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EndEvent {
		xmlElement.M_flowNodeType = "##EndEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_IntermediateThrowEvent {
		xmlElement.M_flowNodeType = "##IntermediateThrowEvent"
	}

	/** LifecycleState **/
	if object.M_lifecycleState == BPMS.LifecycleState_Completed {
		xmlElement.M_lifecycleState = "##Completed"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Compensated {
		xmlElement.M_lifecycleState = "##Compensated"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Failed {
		xmlElement.M_lifecycleState = "##Failed"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Terminated {
		xmlElement.M_lifecycleState = "##Terminated"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Ready {
		xmlElement.M_lifecycleState = "##Ready"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Active {
		xmlElement.M_lifecycleState = "##Active"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Completing {
		xmlElement.M_lifecycleState = "##Completing"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Compensating {
		xmlElement.M_lifecycleState = "##Compensating"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Failing {
		xmlElement.M_lifecycleState = "##Failing"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Terminating {
		xmlElement.M_lifecycleState = "##Terminating"
	}

	/** Serialyze EventDefinitionInstance **/
	if len(object.M_eventDefintionInstances) > 0 {
		xmlElement.M_eventDefintionInstances = make([]*BPMS.XsdEventDefinitionInstance, 0)
	}

	/** Now I will save the value of eventDefintionInstances **/
	for i := 0; i < len(object.M_eventDefintionInstances); i++ {
		xmlElement.M_eventDefintionInstances = append(xmlElement.M_eventDefintionInstances, new(BPMS.XsdEventDefinitionInstance))
		this.SerialyzeEventDefinitionInstance(xmlElement.M_eventDefintionInstances[i], object.M_eventDefintionInstances[i])
	}

	/** EventType **/
	if object.M_eventType == BPMS.EventType_StartEvent {
		xmlElement.M_eventType = "##StartEvent"
	} else if object.M_eventType == BPMS.EventType_IntermediateCatchEvent {
		xmlElement.M_eventType = "##IntermediateCatchEvent"
	} else if object.M_eventType == BPMS.EventType_BoundaryEvent {
		xmlElement.M_eventType = "##BoundaryEvent"
	} else if object.M_eventType == BPMS.EventType_EndEvent {
		xmlElement.M_eventType = "##EndEvent"
	} else if object.M_eventType == BPMS.EventType_IntermediateThrowEvent {
		xmlElement.M_eventType = "##IntermediateThrowEvent"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ConnectingObject **/
func (this *BPMS_XmlFactory) SerialyzeConnectingObject(xmlElement *BPMS.XsdConnectingObject, object *BPMS.ConnectingObject) {
	if xmlElement == nil {
		return
	}

	/** ConnectingObject **/
	xmlElement.M_id = object.M_id

	/** Serialyze ref bpmnElementId **/
	xmlElement.M_bpmnElementId = object.M_bpmnElementId

	/** Serialyze ref sourceRef **/
	xmlElement.M_sourceRef = &object.M_sourceRef

	/** Serialyze ref targetRef **/
	xmlElement.M_targetRef = &object.M_targetRef

	/** ConnectingObjectType **/
	if object.M_connectingObjectType == BPMS.ConnectingObjectType_SequenceFlow {
		xmlElement.M_connectingObjectType = "##SequenceFlow"
	} else if object.M_connectingObjectType == BPMS.ConnectingObjectType_MessageFlow {
		xmlElement.M_connectingObjectType = "##MessageFlow"
	} else if object.M_connectingObjectType == BPMS.ConnectingObjectType_Association {
		xmlElement.M_connectingObjectType = "##Association"
	} else if object.M_connectingObjectType == BPMS.ConnectingObjectType_DataAssociation {
		xmlElement.M_connectingObjectType = "##DataAssociation"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Runtimes **/
func (this *BPMS_XmlFactory) SerialyzeRuntimes(xmlElement *BPMS.XsdRuntimes, object *BPMS.Runtimes) {
	if xmlElement == nil {
		return
	}

	/** Serialyze DefinitionsInstance **/
	if len(object.M_definitions) > 0 {
		xmlElement.M_definitions = make([]*BPMS.XsdDefinitionsInstance, 0)
	}

	/** Now I will save the value of definitions **/
	for i := 0; i < len(object.M_definitions); i++ {
		xmlElement.M_definitions = append(xmlElement.M_definitions, new(BPMS.XsdDefinitionsInstance))
		this.SerialyzeDefinitionsInstance(xmlElement.M_definitions[i], object.M_definitions[i])
	}

	/** Serialyze Exception **/
	if len(object.M_exceptions) > 0 {
		xmlElement.M_exceptions = make([]*BPMS.XsdException, 0)
	}

	/** Now I will save the value of exceptions **/
	for i := 0; i < len(object.M_exceptions); i++ {
		xmlElement.M_exceptions = append(xmlElement.M_exceptions, new(BPMS.XsdException))
		this.SerialyzeException(xmlElement.M_exceptions[i], object.M_exceptions[i])
	}

	/** Serialyze Trigger **/
	if len(object.M_triggers) > 0 {
		xmlElement.M_triggers = make([]*BPMS.XsdTrigger, 0)
	}

	/** Now I will save the value of triggers **/
	for i := 0; i < len(object.M_triggers); i++ {
		xmlElement.M_triggers = append(xmlElement.M_triggers, new(BPMS.XsdTrigger))
		this.SerialyzeTrigger(xmlElement.M_triggers[i], object.M_triggers[i])
	}

	/** Serialyze CorrelationInfo **/
	if len(object.M_correlationInfos) > 0 {
		xmlElement.M_correlationInfos = make([]*BPMS.XsdCorrelationInfo, 0)
	}

	/** Now I will save the value of correlationInfos **/
	for i := 0; i < len(object.M_correlationInfos); i++ {
		xmlElement.M_correlationInfos = append(xmlElement.M_correlationInfos, new(BPMS.XsdCorrelationInfo))
		this.SerialyzeCorrelationInfo(xmlElement.M_correlationInfos[i], object.M_correlationInfos[i])
	}

	/** Runtimes **/
	xmlElement.M_id = object.M_id

	/** Runtimes **/
	xmlElement.M_name = object.M_name

	/** Runtimes **/
	xmlElement.M_version = object.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of DefinitionsInstance **/
func (this *BPMS_XmlFactory) SerialyzeDefinitionsInstance(xmlElement *BPMS.XsdDefinitionsInstance, object *BPMS.DefinitionsInstance) {
	if xmlElement == nil {
		return
	}

	/** DefinitionsInstance **/
	xmlElement.M_id = object.M_id

	/** Serialyze ref bpmnElementId **/
	xmlElement.M_bpmnElementId = object.M_bpmnElementId

	/** Serialyze ProcessInstance **/
	if len(object.M_processInstances) > 0 {
		xmlElement.M_processInstances = make([]*BPMS.XsdProcessInstance, 0)
	}

	/** Now I will save the value of processInstances **/
	for i := 0; i < len(object.M_processInstances); i++ {
		xmlElement.M_processInstances = append(xmlElement.M_processInstances, new(BPMS.XsdProcessInstance))
		this.SerialyzeProcessInstance(xmlElement.M_processInstances[i], object.M_processInstances[i])
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ProcessInstance **/
func (this *BPMS_XmlFactory) SerialyzeProcessInstance(xmlElement *BPMS.XsdProcessInstance, object *BPMS.ProcessInstance) {
	if xmlElement == nil {
		return
	}

	/** ProcessInstance **/
	xmlElement.M_id = object.M_id

	/** Serialyze ref bpmnElementId **/
	xmlElement.M_bpmnElementId = object.M_bpmnElementId

	/** Serialyze FlowNodeInstance **/
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_0 = make([]*BPMS.XsdActivityInstance, 0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_1 = make([]*BPMS.XsdSubprocessInstance, 0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_2 = make([]*BPMS.XsdGatewayInstance, 0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_3 = make([]*BPMS.XsdEventInstance, 0)
	}

	/** Now I will save the value of flowNodeInstances **/
	for i := 0; i < len(object.M_flowNodeInstances); i++ {
		switch v := object.M_flowNodeInstances[i].(type) {
		case *BPMS.ActivityInstance:
			xmlElement.M_flowNodeInstances_0 = append(xmlElement.M_flowNodeInstances_0, new(BPMS.XsdActivityInstance))
			this.SerialyzeActivityInstance(xmlElement.M_flowNodeInstances_0[len(xmlElement.M_flowNodeInstances_0)-1], v)
		case *BPMS.SubprocessInstance:
			xmlElement.M_flowNodeInstances_1 = append(xmlElement.M_flowNodeInstances_1, new(BPMS.XsdSubprocessInstance))
			this.SerialyzeSubprocessInstance(xmlElement.M_flowNodeInstances_1[len(xmlElement.M_flowNodeInstances_1)-1], v)
		case *BPMS.GatewayInstance:
			xmlElement.M_flowNodeInstances_2 = append(xmlElement.M_flowNodeInstances_2, new(BPMS.XsdGatewayInstance))
			this.SerialyzeGatewayInstance(xmlElement.M_flowNodeInstances_2[len(xmlElement.M_flowNodeInstances_2)-1], v)
		case *BPMS.EventInstance:
			xmlElement.M_flowNodeInstances_3 = append(xmlElement.M_flowNodeInstances_3, new(BPMS.XsdEventInstance))
			this.SerialyzeEventInstance(xmlElement.M_flowNodeInstances_3[len(xmlElement.M_flowNodeInstances_3)-1], v)
		}
	}

	/** Instance **/
	xmlElement.M_number = object.M_number

	/** Instance **/
	xmlElement.M_colorNumber = object.M_colorNumber

	/** Instance **/
	xmlElement.M_colorName = object.M_colorName
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of EventData **/
func (this *BPMS_XmlFactory) SerialyzeEventData(xmlElement *BPMS.XsdEventData, object *BPMS.EventData) {
	if xmlElement == nil {
		return
	}

	/** EventData **/
	xmlElement.M_id = object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of CorrelationInfo **/
func (this *BPMS_XmlFactory) SerialyzeCorrelationInfo(xmlElement *BPMS.XsdCorrelationInfo, object *BPMS.CorrelationInfo) {
	if xmlElement == nil {
		return
	}

	/** CorrelationInfo **/
	xmlElement.M_id = object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ActivityInstance **/
func (this *BPMS_XmlFactory) SerialyzeActivityInstance(xmlElement *BPMS.XsdActivityInstance, object *BPMS.ActivityInstance) {
	if xmlElement == nil {
		return
	}

	/** ActivityInstance **/
	xmlElement.M_id = object.M_id

	/** Serialyze ref bpmnElementId **/
	xmlElement.M_bpmnElementId = object.M_bpmnElementId

	/** Serialyze ref inputRef **/
	xmlElement.M_inputRef = object.M_inputRef

	/** Serialyze ref outputRef **/
	xmlElement.M_outputRef = object.M_outputRef

	/** FlowNodeType **/
	if object.M_flowNodeType == BPMS.FlowNodeType_AbstractTask {
		xmlElement.M_flowNodeType = "##AbstractTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ServiceTask {
		xmlElement.M_flowNodeType = "##ServiceTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_UserTask {
		xmlElement.M_flowNodeType = "##UserTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ManualTask {
		xmlElement.M_flowNodeType = "##ManualTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_BusinessRuleTask {
		xmlElement.M_flowNodeType = "##BusinessRuleTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ScriptTask {
		xmlElement.M_flowNodeType = "##ScriptTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EmbeddedSubprocess {
		xmlElement.M_flowNodeType = "##EmbeddedSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EventSubprocess {
		xmlElement.M_flowNodeType = "##EventSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_AdHocSubprocess {
		xmlElement.M_flowNodeType = "##AdHocSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_Transaction {
		xmlElement.M_flowNodeType = "##Transaction"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_CallActivity {
		xmlElement.M_flowNodeType = "##CallActivity"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ParallelGateway {
		xmlElement.M_flowNodeType = "##ParallelGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ExclusiveGateway {
		xmlElement.M_flowNodeType = "##ExclusiveGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_InclusiveGateway {
		xmlElement.M_flowNodeType = "##InclusiveGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EventBasedGateway {
		xmlElement.M_flowNodeType = "##EventBasedGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ComplexGateway {
		xmlElement.M_flowNodeType = "##ComplexGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_StartEvent {
		xmlElement.M_flowNodeType = "##StartEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_IntermediateCatchEvent {
		xmlElement.M_flowNodeType = "##IntermediateCatchEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_BoundaryEvent {
		xmlElement.M_flowNodeType = "##BoundaryEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EndEvent {
		xmlElement.M_flowNodeType = "##EndEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_IntermediateThrowEvent {
		xmlElement.M_flowNodeType = "##IntermediateThrowEvent"
	}

	/** LifecycleState **/
	if object.M_lifecycleState == BPMS.LifecycleState_Completed {
		xmlElement.M_lifecycleState = "##Completed"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Compensated {
		xmlElement.M_lifecycleState = "##Compensated"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Failed {
		xmlElement.M_lifecycleState = "##Failed"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Terminated {
		xmlElement.M_lifecycleState = "##Terminated"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Ready {
		xmlElement.M_lifecycleState = "##Ready"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Active {
		xmlElement.M_lifecycleState = "##Active"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Completing {
		xmlElement.M_lifecycleState = "##Completing"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Compensating {
		xmlElement.M_lifecycleState = "##Compensating"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Failing {
		xmlElement.M_lifecycleState = "##Failing"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Terminating {
		xmlElement.M_lifecycleState = "##Terminating"
	}

	/** ActivityType **/
	if object.M_activityType == BPMS.ActivityType_AbstractTask {
		xmlElement.M_activityType = "##AbstractTask"
	} else if object.M_activityType == BPMS.ActivityType_ServiceTask {
		xmlElement.M_activityType = "##ServiceTask"
	} else if object.M_activityType == BPMS.ActivityType_UserTask {
		xmlElement.M_activityType = "##UserTask"
	} else if object.M_activityType == BPMS.ActivityType_ManualTask {
		xmlElement.M_activityType = "##ManualTask"
	} else if object.M_activityType == BPMS.ActivityType_BusinessRuleTask {
		xmlElement.M_activityType = "##BusinessRuleTask"
	} else if object.M_activityType == BPMS.ActivityType_ScriptTask {
		xmlElement.M_activityType = "##ScriptTask"
	} else if object.M_activityType == BPMS.ActivityType_EmbeddedSubprocess {
		xmlElement.M_activityType = "##EmbeddedSubprocess"
	} else if object.M_activityType == BPMS.ActivityType_EventSubprocess {
		xmlElement.M_activityType = "##EventSubprocess"
	} else if object.M_activityType == BPMS.ActivityType_AdHocSubprocess {
		xmlElement.M_activityType = "##AdHocSubprocess"
	} else if object.M_activityType == BPMS.ActivityType_Transaction {
		xmlElement.M_activityType = "##Transaction"
	} else if object.M_activityType == BPMS.ActivityType_CallActivity {
		xmlElement.M_activityType = "##CallActivity"
	}

	/** LoopCharacteristicType **/
	if object.M_loopCharacteristicType == BPMS.LoopCharacteristicType_StandardLoopCharacteristics {
		xmlElement.M_loopCharacteristicType = "##StandardLoopCharacteristics"
	} else if object.M_loopCharacteristicType == BPMS.LoopCharacteristicType_MultiInstanceLoopCharacteristics {
		xmlElement.M_loopCharacteristicType = "##MultiInstanceLoopCharacteristics"
	}

	/** MultiInstanceBehaviorType **/
	if object.M_multiInstanceBehaviorType == BPMS.MultiInstanceBehaviorType_None {
		xmlElement.M_multiInstanceBehaviorType = "##None"
	} else if object.M_multiInstanceBehaviorType == BPMS.MultiInstanceBehaviorType_One {
		xmlElement.M_multiInstanceBehaviorType = "##One"
	} else if object.M_multiInstanceBehaviorType == BPMS.MultiInstanceBehaviorType_All {
		xmlElement.M_multiInstanceBehaviorType = "##All"
	} else if object.M_multiInstanceBehaviorType == BPMS.MultiInstanceBehaviorType_Complex {
		xmlElement.M_multiInstanceBehaviorType = "##Complex"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Exception **/
func (this *BPMS_XmlFactory) SerialyzeException(xmlElement *BPMS.XsdException, object *BPMS.Exception) {
	if xmlElement == nil {
		return
	}

	/** ExceptionType **/
	if object.M_exceptionType == BPMS.ExceptionType_GatewayException {
		xmlElement.M_exceptionType = "##GatewayException"
	} else if object.M_exceptionType == BPMS.ExceptionType_NoIORuleException {
		xmlElement.M_exceptionType = "##NoIORuleException"
	} else if object.M_exceptionType == BPMS.ExceptionType_NoAvailableOutputSetException {
		xmlElement.M_exceptionType = "##NoAvailableOutputSetException"
	} else if object.M_exceptionType == BPMS.ExceptionType_NotMatchingIOSpecification {
		xmlElement.M_exceptionType = "##NotMatchingIOSpecification"
	} else if object.M_exceptionType == BPMS.ExceptionType_IllegalStartEventException {
		xmlElement.M_exceptionType = "##IllegalStartEventException"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Trigger **/
func (this *BPMS_XmlFactory) SerialyzeTrigger(xmlElement *BPMS.XsdTrigger, object *BPMS.Trigger) {
	if xmlElement == nil {
		return
	}

	/** Serialyze EventData **/
	if len(object.M_eventDatas) > 0 {
		xmlElement.M_eventDatas = make([]*BPMS.XsdEventData, 0)
	}

	/** Now I will save the value of eventDatas **/
	for i := 0; i < len(object.M_eventDatas); i++ {
		xmlElement.M_eventDatas = append(xmlElement.M_eventDatas, new(BPMS.XsdEventData))
		this.SerialyzeEventData(xmlElement.M_eventDatas[i], object.M_eventDatas[i])
	}

	/** Serialyze ref dataRef **/
	xmlElement.M_dataRef = object.M_dataRef

	/** Serialyze ref sourceRef **/
	xmlElement.M_sourceRef = &object.M_sourceRef

	/** Serialyze ref targetRef **/
	xmlElement.M_targetRef = &object.M_targetRef

	/** Trigger **/
	xmlElement.M_id = object.M_id

	/** Trigger **/
	xmlElement.M_processUUID = object.M_processUUID

	/** Trigger **/
	xmlElement.M_sessionId = object.M_sessionId

	/** EventTriggerType **/
	if object.M_eventTriggerType == BPMS.EventTriggerType_None {
		xmlElement.M_eventTriggerType = "##None"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Timer {
		xmlElement.M_eventTriggerType = "##Timer"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Conditional {
		xmlElement.M_eventTriggerType = "##Conditional"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Message {
		xmlElement.M_eventTriggerType = "##Message"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Signal {
		xmlElement.M_eventTriggerType = "##Signal"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Multiple {
		xmlElement.M_eventTriggerType = "##Multiple"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_ParallelMultiple {
		xmlElement.M_eventTriggerType = "##ParallelMultiple"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Escalation {
		xmlElement.M_eventTriggerType = "##Escalation"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Error {
		xmlElement.M_eventTriggerType = "##Error"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Compensation {
		xmlElement.M_eventTriggerType = "##Compensation"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Terminate {
		xmlElement.M_eventTriggerType = "##Terminate"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Cancel {
		xmlElement.M_eventTriggerType = "##Cancel"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Link {
		xmlElement.M_eventTriggerType = "##Link"
	} else if object.M_eventTriggerType == BPMS.EventTriggerType_Start {
		xmlElement.M_eventTriggerType = "##Start"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of SubprocessInstance **/
func (this *BPMS_XmlFactory) SerialyzeSubprocessInstance(xmlElement *BPMS.XsdSubprocessInstance, object *BPMS.SubprocessInstance) {
	if xmlElement == nil {
		return
	}

	/** SubprocessInstance **/
	xmlElement.M_id = object.M_id

	/** Serialyze ref bpmnElementId **/
	xmlElement.M_bpmnElementId = object.M_bpmnElementId

	/** Serialyze ref inputRef **/
	xmlElement.M_inputRef = object.M_inputRef

	/** Serialyze ref outputRef **/
	xmlElement.M_outputRef = object.M_outputRef

	/** FlowNodeType **/
	if object.M_flowNodeType == BPMS.FlowNodeType_AbstractTask {
		xmlElement.M_flowNodeType = "##AbstractTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ServiceTask {
		xmlElement.M_flowNodeType = "##ServiceTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_UserTask {
		xmlElement.M_flowNodeType = "##UserTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ManualTask {
		xmlElement.M_flowNodeType = "##ManualTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_BusinessRuleTask {
		xmlElement.M_flowNodeType = "##BusinessRuleTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ScriptTask {
		xmlElement.M_flowNodeType = "##ScriptTask"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EmbeddedSubprocess {
		xmlElement.M_flowNodeType = "##EmbeddedSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EventSubprocess {
		xmlElement.M_flowNodeType = "##EventSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_AdHocSubprocess {
		xmlElement.M_flowNodeType = "##AdHocSubprocess"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_Transaction {
		xmlElement.M_flowNodeType = "##Transaction"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_CallActivity {
		xmlElement.M_flowNodeType = "##CallActivity"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ParallelGateway {
		xmlElement.M_flowNodeType = "##ParallelGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ExclusiveGateway {
		xmlElement.M_flowNodeType = "##ExclusiveGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_InclusiveGateway {
		xmlElement.M_flowNodeType = "##InclusiveGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EventBasedGateway {
		xmlElement.M_flowNodeType = "##EventBasedGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_ComplexGateway {
		xmlElement.M_flowNodeType = "##ComplexGateway"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_StartEvent {
		xmlElement.M_flowNodeType = "##StartEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_IntermediateCatchEvent {
		xmlElement.M_flowNodeType = "##IntermediateCatchEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_BoundaryEvent {
		xmlElement.M_flowNodeType = "##BoundaryEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_EndEvent {
		xmlElement.M_flowNodeType = "##EndEvent"
	} else if object.M_flowNodeType == BPMS.FlowNodeType_IntermediateThrowEvent {
		xmlElement.M_flowNodeType = "##IntermediateThrowEvent"
	}

	/** LifecycleState **/
	if object.M_lifecycleState == BPMS.LifecycleState_Completed {
		xmlElement.M_lifecycleState = "##Completed"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Compensated {
		xmlElement.M_lifecycleState = "##Compensated"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Failed {
		xmlElement.M_lifecycleState = "##Failed"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Terminated {
		xmlElement.M_lifecycleState = "##Terminated"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Ready {
		xmlElement.M_lifecycleState = "##Ready"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Active {
		xmlElement.M_lifecycleState = "##Active"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Completing {
		xmlElement.M_lifecycleState = "##Completing"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Compensating {
		xmlElement.M_lifecycleState = "##Compensating"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Failing {
		xmlElement.M_lifecycleState = "##Failing"
	} else if object.M_lifecycleState == BPMS.LifecycleState_Terminating {
		xmlElement.M_lifecycleState = "##Terminating"
	}

	/** Serialyze FlowNodeInstance **/
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_0 = make([]*BPMS.XsdActivityInstance, 0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_1 = make([]*BPMS.XsdSubprocessInstance, 0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_2 = make([]*BPMS.XsdGatewayInstance, 0)
	}
	if len(object.M_flowNodeInstances) > 0 {
		xmlElement.M_flowNodeInstances_3 = make([]*BPMS.XsdEventInstance, 0)
	}

	/** Now I will save the value of flowNodeInstances **/
	for i := 0; i < len(object.M_flowNodeInstances); i++ {
		switch v := object.M_flowNodeInstances[i].(type) {
		case *BPMS.ActivityInstance:
			xmlElement.M_flowNodeInstances_0 = append(xmlElement.M_flowNodeInstances_0, new(BPMS.XsdActivityInstance))
			this.SerialyzeActivityInstance(xmlElement.M_flowNodeInstances_0[len(xmlElement.M_flowNodeInstances_0)-1], v)
		case *BPMS.SubprocessInstance:
			xmlElement.M_flowNodeInstances_1 = append(xmlElement.M_flowNodeInstances_1, new(BPMS.XsdSubprocessInstance))
			this.SerialyzeSubprocessInstance(xmlElement.M_flowNodeInstances_1[len(xmlElement.M_flowNodeInstances_1)-1], v)
		case *BPMS.GatewayInstance:
			xmlElement.M_flowNodeInstances_2 = append(xmlElement.M_flowNodeInstances_2, new(BPMS.XsdGatewayInstance))
			this.SerialyzeGatewayInstance(xmlElement.M_flowNodeInstances_2[len(xmlElement.M_flowNodeInstances_2)-1], v)
		case *BPMS.EventInstance:
			xmlElement.M_flowNodeInstances_3 = append(xmlElement.M_flowNodeInstances_3, new(BPMS.XsdEventInstance))
			this.SerialyzeEventInstance(xmlElement.M_flowNodeInstances_3[len(xmlElement.M_flowNodeInstances_3)-1], v)
		}
	}

	/** Serialyze ConnectingObject **/
	if len(object.M_connectingObjects) > 0 {
		xmlElement.M_connectingObjects = make([]*BPMS.XsdConnectingObject, 0)
	}

	/** Now I will save the value of connectingObjects **/
	for i := 0; i < len(object.M_connectingObjects); i++ {
		xmlElement.M_connectingObjects = append(xmlElement.M_connectingObjects, new(BPMS.XsdConnectingObject))
		this.SerialyzeConnectingObject(xmlElement.M_connectingObjects[i], object.M_connectingObjects[i])
	}

	/** SubprocessType **/
	if object.M_SubprocessType == BPMS.SubprocessType_EmbeddedSubprocess {
		xmlElement.M_SubprocessType = "##EmbeddedSubprocess"
	} else if object.M_SubprocessType == BPMS.SubprocessType_EventSubprocess {
		xmlElement.M_SubprocessType = "##EventSubprocess"
	} else if object.M_SubprocessType == BPMS.SubprocessType_AdHocSubprocess {
		xmlElement.M_SubprocessType = "##AdHocSubprocess"
	} else if object.M_SubprocessType == BPMS.SubprocessType_Transaction {
		xmlElement.M_SubprocessType = "##Transaction"
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}
